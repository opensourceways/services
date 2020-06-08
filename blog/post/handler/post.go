package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	tagProto "github.com/micro/examples/blog/tag/proto/tag"

	"github.com/gosimple/slug"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"

	post "github.com/micro/examples/blog/post/proto/post"
)

const (
	tagType         = "post-tag"
	slugPrefix      = "slug"
	timeStampPrefix = "timestamp"
)

type Post struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	Content         string   `json:"content"`
	CreateTimestamp int64    `json:"create_timestamp"`
	UpdateTimestamp int64    `json:"update_timestamp"`
	TagNames        []string `json:"tagNames"`
}

type PostService struct {
	Store  store.Store
	Client client.Client
}

func (t *PostService) Post(ctx context.Context, req *post.PostRequest, rsp *post.PostResponse) error {
	if len(req.Post.Id) == 0 || len(req.Post.Title) == 0 || len(req.Post.Content) == 0 {
		return errors.New("ID, title or content is missing")
	}

	// read by parent ID + slug, the record is identical in boths places anyway
	records, err := t.Store.Read(req.Post.Slug)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	postSlug := slug.Make(req.Post.Title)
	// If no existing record is found, create a new one
	if len(records) == 0 {
		post := &Post{
			ID:              req.Post.Id,
			Title:           req.Post.Title,
			Content:         req.Post.Content,
			TagNames:        req.Post.TagNames,
			Slug:            postSlug,
			CreateTimestamp: time.Now().Unix(),
		}
		return t.savePost(ctx, nil, post)
	}
	record := records[0]
	oldPost := &Post{}
	err = json.Unmarshal(record.Value, oldPost)
	if err != nil {
		return err
	}
	post := &Post{
		ID:              req.Post.Id,
		Title:           req.Post.Title,
		Slug:            postSlug,
		TagNames:        req.Post.TagNames,
		CreateTimestamp: oldPost.CreateTimestamp,
		UpdateTimestamp: time.Now().Unix(),
	}
	return t.savePost(ctx, oldPost, post)
}

func (t *PostService) savePost(ctx context.Context, oldPost, post *Post) error {
	bytes, err := json.Marshal(post)
	if err != nil {
		return err
	}

	err = t.Store.Write(&store.Record{
		Key:   fmt.Sprintf("%v:%v", slugPrefix, post.Slug),
		Value: bytes,
	})
	if err != nil {
		return err
	}
	err = t.Store.Write(&store.Record{
		Key:   fmt.Sprintf("%v:%v", timeStampPrefix, math.MaxInt64-post.CreateTimestamp),
		Value: bytes,
	})
	if err != nil {
		return err
	}
	if oldPost == nil {
		tagClient := tagProto.NewTagService("go.micro.service.tag", t.Client)
		for _, tagName := range post.TagNames {
			_, err := tagClient.IncreaseCount(ctx, &tagProto.IncreaseCountRequest{
				ParentID: post.ID,
				Type:     tagType,
				Title:    tagName,
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
	return t.diffTags(ctx, oldPost.ID, oldPost.TagNames, post.TagNames)
}

func (t *PostService) diffTags(ctx context.Context, parentID string, oldTagNames, newTagNames []string) error {
	oldTags := map[string]struct{}{}
	for _, v := range oldTagNames {
		oldTags[v] = struct{}{}
	}
	newTags := map[string]struct{}{}
	for _, v := range newTagNames {
		newTags[v] = struct{}{}
	}
	tagClient := tagProto.NewTagService("go.micro.service.tag", t.Client)
	for i := range oldTags {
		_, stillThere := newTags[i]
		if !stillThere {
			tagClient.DecreaseCount(ctx, &tagProto.DecreaseCountRequest{
				ParentID: parentID,
				Type:     tagType,
				Title:    i,
			})
		}
	}
	for i := range newTags {
		_, newlyAdded := oldTags[i]
		if newlyAdded {
			tagClient.IncreaseCount(ctx, &tagProto.IncreaseCountRequest{
				ParentID: parentID,
				Type:     tagType,
				Title:    i,
			})
		}
	}
	return nil
}

func (t *PostService) Query(ctx context.Context, req *post.QueryRequest, rsp *post.QueryResponse) error {
	var records []*store.Record
	var err error
	if len(req.Slug) > 0 {
		key := fmt.Sprintf("%v:%v", slugPrefix, req.Slug)
		log.Infof("Reading post by slug: %v", req.Slug)
		records, err = t.Store.Read(key, store.ReadPrefix())
	} else {
		key := fmt.Sprintf("%v:", timeStampPrefix)
		var limit uint
		limit = 20
		if req.Limit > 0 {
			limit = uint(req.Limit)
		}
		log.Infof("Listing posts, offset: %v, limit: %v", req.Offset, limit)
		records, err = t.Store.Read(key, store.ReadPrefix(),
			store.ReadOffset(uint(req.Offset)),
			store.ReadLimit(limit))
	}

	if err != nil {
		return err
	}
	rsp.Posts = make([]*post.Post, len(records))
	for i, record := range records {
		postRecord := &Post{}
		err := json.Unmarshal(record.Value, postRecord)
		if err != nil {
			return err
		}
		rsp.Posts[i] = &post.Post{
			Id:       postRecord.ID,
			Title:    postRecord.Title,
			Slug:     postRecord.Slug,
			Content:  postRecord.Content,
			TagNames: postRecord.TagNames,
		}
	}
	return nil
}

func (t *PostService) Delete(ctx context.Context, req *post.DeleteRequest, rsp *post.DeleteResponse) error {
	log.Info("Received Post.Delete request")
	return t.Store.Delete(fmt.Sprintf("%v:%v", slugPrefix, req.Slug))
}
