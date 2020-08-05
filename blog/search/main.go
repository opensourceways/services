package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/blog/search/handler"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.search"),
	)

	// Register Handler
	srv.Handle(new(handler.Search))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
