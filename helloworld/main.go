package main

import (
	"github.com/micro/examples/helloworld/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	helloworld "github.com/micro/examples/helloworld/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
