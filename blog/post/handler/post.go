package handler

import (
	"context"
	"encoding/json"
	"errors"

	tagProto "github.com/micro/examples/blog/tag/proto/tag"

	"github.com/gosimple/slug"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"

	post "github.com/micro/examples/blog/post/proto/post"
)

const tagType = "post-tag"

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
			ID:       req.Post.Id,
			Title:    req.Post.Title,
			Content:  req.Post.Content,
			TagNames: req.Post.TagNames,
			Slug:     postSlug,
		}
		bytes, err := json.Marshal(post)
		if err != nil {
			return err
		}

		err = t.Store.Write(&store.Record{
			Key:   post.Slug,
			Value: bytes,
		})
		if err != nil {
			return err
		}
		tagClient := tagProto.NewTagService("go.micro.service.tag", t.Client)
		for _, tagName := range req.Post.TagNames {
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
	record := records[0]
	oldPost := &Post{}
	err = json.Unmarshal(record.Value, oldPost)
	if err != nil {
		return err
	}
	post := &Post{
		ID:       req.Post.Id,
		Title:    req.Post.Title,
		Slug:     postSlug,
		TagNames: req.Post.TagNames,
	}
	bytes, err := json.Marshal(post)
	if err != nil {
		return err
	}
	err = t.Store.Write(&store.Record{
		Key:   post.Slug,
		Value: bytes,
	})
	return t.diffTags(ctx, post.ID, oldPost.TagNames, post.TagNames)
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
	log.Info("Received Post.Query request")

	key := ""
	if len(req.Slug) > 0 {
		key = req.Slug
	} else {
		return errors.New("List params required")
	}

	records, err := t.Store.Read(key, store.ReadPrefix())
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
	log.Info("Received Post.Call request")
	return nil
}
