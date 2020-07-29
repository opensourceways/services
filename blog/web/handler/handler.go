package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/lithammer/shortuuid/v3"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service/client"
	postsproto "github.com/micro/services/blog/posts/proto/posts"
	templ "github.com/micro/services/blog/web/templates"
)

type Handler struct{}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	request := client.DefaultClient.NewRequest("go.micro.service.posts", "Posts.Query", &postsproto.QueryRequest{})
	rsp := &postsproto.QueryResponse{}
	if err := client.DefaultClient.Call(r.Context(), request, rsp); err != nil {
		fmt.Println("err", err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("Serving index")

	postTemplate := templ.Header + templ.IndexBody + templ.Footer
	t, err := template.New("webpage").Parse(postTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := map[string]interface{}{
		"Posts": rsp.Posts,
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, vars)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, buf.String())
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	pastPostFragments := strings.Split(r.URL.Path, "post/")
	slug := strings.Split(pastPostFragments[1], "/")[0]
	log.Infof("Getting post by slug: %v, for path: %v", slug, r.URL.Path)

	request := client.DefaultClient.NewRequest("go.micro.service.posts", "Posts.Query", &postsproto.QueryRequest{
		Slug: slug,
	})
	rsp := &postsproto.QueryResponse{}
	if err := client.DefaultClient.Call(r.Context(), request, rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(rsp.Posts) == 0 {
		http.Error(w, "Not found", 404)
		return
	}

	postTemplate := templ.Header + templ.PostBody + templ.Footer
	t, err := template.New("webpage").Parse(postTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := map[string]interface{}{
		"Post": rsp.Posts[0],
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, vars)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, buf.String())
}

func (h Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	newPostTemplate := templ.Header + templ.NewPostBody + templ.Footer
	t, err := template.New("webpage").Parse(newPostTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := map[string]interface{}{}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, vars)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, buf.String())
}

func (h Handler) EditPost(w http.ResponseWriter, r *http.Request) {
	pastPostFragments := strings.Split(r.URL.Path, "edit/")
	slug := strings.Split(pastPostFragments[1], "/")[0]

	request := client.DefaultClient.NewRequest("go.micro.service.posts", "Posts.Query", &postsproto.QueryRequest{
		Slug: slug,
	})
	rsp := &postsproto.QueryResponse{}
	if err := client.DefaultClient.Call(r.Context(), request, rsp); err != nil {
		fmt.Println("err", err)
		http.Error(w, err.Error(), 500)
		return
	}
	if len(rsp.Posts) == 0 {
		http.Error(w, "Not found", 404)
		return
	}

	editPostTemplate := templ.Header + templ.EditPostBody + templ.Footer
	t, err := template.New("webpage").Parse(editPostTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := map[string]interface{}{
		"Post":     rsp.Posts[0],
		"TagNames": strings.Join(rsp.Posts[0].TagNames, ", "),
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, vars)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, buf.String())
}

func (h Handler) PostAPI(w http.ResponseWriter, r *http.Request) {
	log.Infof("New post request")
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	tagNames := []string{}
	if len(r.PostFormValue("tagNames")) > 0 {
		for _, tagName := range strings.Split(r.PostFormValue("tagNames"), ",") {
			tagNames = append(tagNames, strings.TrimSpace(tagName))
		}
	}
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	id := r.PostFormValue("id")
	if len(id) == 0 {
		id = shortuuid.New()
	}

	log.Infof("Creating post with id and title %v: %v", title)
	request := client.DefaultClient.NewRequest("go.micro.service.posts", "Posts.Post", &postsproto.PostRequest{
		Post: &postsproto.Post{
			Id:       id,
			Title:    title,
			Content:  content,
			TagNames: tagNames,
		},
	})
	rsp := &postsproto.PostResponse{}
	if err := client.DefaultClient.Call(r.Context(), request, rsp); err != nil {
		fmt.Println("Error creating post: ", err)
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", 301)
}
