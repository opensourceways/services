package main

import (
	"github.com/micro/examples/blog/tag/handler"
	"github.com/micro/examples/blog/tag/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	tag "github.com/micro/examples/blog/tag/proto/tag"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.tag"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	tag.RegisterTagServiceHandler(service.Server(), &handler.TagService{
		Store: service.Options().Store,
	})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.tag", service.Server(), new(subscriber.Tag))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
