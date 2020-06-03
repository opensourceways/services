package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"blog-web/handler"
	"blog-web/subscriber"

	blog "blog-web/proto/blog"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.blog"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	blog.RegisterBlogHandler(service.Server(), new(handler.Blog))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.blog", service.Server(), new(subscriber.Blog))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
