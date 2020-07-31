package main

import (
	"github.com/micro/micro/v3/service"

	"github.com/micro/services/blog/posts/handler"
	pb "github.com/micro/services/blog/posts/proto/posts"
)

func main() {
	// Register Handler
	pb.RegisterPostsHandler(handler.New())

	// Run service
	service.Run()
}
