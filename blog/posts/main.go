package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/posts/handler"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.posts"),
		service.Version("latest"),
	)

	// Register Handler
	srv.Handle(&handler.Posts{
		Client: srv.Client(),
	})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
