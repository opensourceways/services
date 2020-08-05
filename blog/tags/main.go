package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/blog/tags/handler"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.tags"),
	)

	// Register Handler
	srv.Handle(new(handler.Tags))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
