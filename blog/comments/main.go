package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/micro/services/blog/comments/handler"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("comments"),
	)

	// Register Handler
	srv.Handle(handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
