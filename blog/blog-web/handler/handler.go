package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	templ "github.com/micro/examples/blog/blog-web/templates"
	postproto "github.com/micro/examples/blog/post/proto/post"
	"github.com/micro/go-micro/v2/client"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json

	// we want to augment the response
	response := map[string]interface{}{
		"msg": "hi",
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received post req")
	pastPostFragments := strings.Split(r.URL.Path, "post/")
	slug := strings.Split(pastPostFragments[0], "/")[0]

	request := h.Client.NewRequest("go.micro.service.post", "PostService.Query", &postproto.QueryRequest{
		Slug: slug,
	})
	rsp := &postproto.QueryResponse{}
	if err := h.Client.Call(r.Context(), request, rsp); err != nil {
		fmt.Println(err)
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
