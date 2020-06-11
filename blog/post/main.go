package main

import (
	"github.com/micro/examples/blog/post/handler"
	"github.com/micro/examples/blog/post/subscriber"

	"github.com/micro/go-micro/v2"

	log "github.com/micro/go-micro/v2/logger"

	post "github.com/micro/examples/blog/post/proto/post"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.post"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	post.RegisterPostServiceHandler(service.Server(), &handler.PostService{
		Store:  service.Options().Store,
		Client: service.Client(),
	})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.post", service.Server(), new(subscriber.Post))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
