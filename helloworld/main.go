package main

import (
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/micro/v2/service"
	"github.com/micro/services/helloworld/handler"

	pb "github.com/micro/services/helloworld/proto"
)

func main() {
	// New Service
	helloworld := service.New(
		service.Name("helloworld"),
	)

	// Initialise service
	helloworld.Init()

	// Register Handler
	pb.RegisterHelloworldHandler(helloworld.Server(), new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}
