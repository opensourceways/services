package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"comment/handler"
	"comment/subscriber"

	comment "comment/proto/comment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.comment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	comment.RegisterCommentHandler(service.Server(), new(handler.Comment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.comment", service.Server(), new(subscriber.Comment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
