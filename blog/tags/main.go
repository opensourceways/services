package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/tags/handler"
	"github.com/micro/services/blog/tags/subscriber"

	tags "github.com/micro/services/blog/tags/proto/tags"
)

func main() {
	// New Serviceost
	srv := service.New(
		service.Name("go.micro.service.tags"),
		service.Version("latest"),
	)

	// Register Handler
	tags.RegisterTagsHandler(srv.Server(), new(handler.Tags))

	// Register Struct as Subscriber
	service.RegisterSubscriber("go.micro.service.tags", new(subscriber.Tags))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
