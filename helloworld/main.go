package main

import (
	"github.com/micro/examples/helloworld/handler"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/micro/v2/service"

	pb "github.com/micro/examples/helloworld/proto"
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
