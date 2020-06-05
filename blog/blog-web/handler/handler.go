package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	templ "github.com/micro/examples/blog/blog-web/templates"
	postproto "github.com/micro/examples/blog/post/proto/post"
	"github.com/micro/go-micro/v2/client"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	request := h.Client.NewRequest("go.micro.service.post", "PostService.Query", &postproto.QueryRequest{})
	rsp := &postproto.QueryResponse{}
	if err := h.Client.Call(r.Context(), request, rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	postTemplate := templ.Header + templ.IndexBody + templ.Footer
	t, err := template.New("webpage").Parse(postTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(rsp.Posts) == 0 {
		http.Error(w, "Not found", 404)
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

	request := h.Client.NewRequest("go.micro.service.post", "PostService.Query", &postproto.QueryRequest{
		Slug: slug,
	})
	rsp := &postproto.QueryResponse{}
	if err := h.Client.Call(r.Context(), request, rsp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	postTemplate := templ.Header + templ.PostBody + templ.Footer
	t, err := template.New("webpage").Parse(postTemplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(rsp.Posts) == 0 {
		http.Error(w, "Not found", 404)
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

func (h Handler) PostAPI(w http.ResponseWriter, r *http.Request) {

}
