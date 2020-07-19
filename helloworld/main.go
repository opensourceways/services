package main

import (
	"github.com/micro/examples/helloworld/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/logger"

	pb "github.com/micro/examples/helloworld/proto"
)

func main() {
	// New Service
	helloworld := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Initialise service
	helloworld.Init()

	// Register Handler
	pb.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}
