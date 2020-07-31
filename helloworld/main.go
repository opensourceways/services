package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/helloworld/handler"

	pb "github.com/micro/services/helloworld/proto"
)

func main() {
	// Register Handler
	pb.RegisterHelloworldHandler(new(handler.Helloworld))

	// Run the service
	service.Run()
}
