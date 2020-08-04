package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/helloworld/handler"
)

func main() {
	// New Service
	helloworld := service.New(
		service.Name("helloworld"),
	)

	// Register Handler
	helloworld.Handle(new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}
