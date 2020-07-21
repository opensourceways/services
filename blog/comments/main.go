package main

import (
	"github.com/micro/examples/blog/comments/handler"
	"github.com/micro/examples/blog/comments/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	comments "github.com/micro/examples/blog/comments/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.comments"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	comments.RegisterCommentsHandler(service.Server(), new(handler.Comments))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.comments", service.Server(), new(subscriber.Comments))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
