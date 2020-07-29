package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/posts/handler"
	posts "github.com/micro/services/blog/posts/proto/posts"
	"github.com/micro/services/blog/posts/subscriber"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.posts"),
		service.Version("latest"),
	)

	// Register Handler
	posts.RegisterPostsHandler(srv.Server(), &handler.Posts{
		Client: srv.Client(),
	})

	// Register Struct as Subscriber
	service.RegisterSubscriber("go.micro.service.posts", new(subscriber.Post))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
