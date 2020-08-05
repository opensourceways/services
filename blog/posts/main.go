package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/micro/services/blog/posts/handler"
	pb "github.com/micro/services/blog/posts/proto/posts"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("comments"),
	)

	// Register Handler
	pb.RegisterPostsHandler(srv.Server(), handler.New(srv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
