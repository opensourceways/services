package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/blog/comments/handler"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.comments"),
	)

	// Register Handler
	srv.Handle(new(handler.Comments))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
