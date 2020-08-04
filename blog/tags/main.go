package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/tags/handler"
)

func main() {
	// New Serviceost
	srv := service.New(
		service.Name("go.micro.service.tags"),
		service.Version("latest"),
	)

	// Register Handler
	srv.Handle(new(handler.Tags))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
