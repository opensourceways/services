package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/comments/handler"
	"github.com/micro/services/blog/comments/subscriber"

	comments "github.com/micro/services/blog/comments/proto"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.comments"),
		service.Version("latest"),
	)

	// Register Handler
	comments.RegisterCommentsHandler(srv.Server(), new(handler.Comments))

	// Register Struct as Subscriber
	service.RegisterSubscriber("go.micro.service.comments", new(subscriber.Comments))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
