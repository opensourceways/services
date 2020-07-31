package main

import (
	"github.com/micro/micro/v3/service"

	"github.com/micro/services/blog/tags/handler"
	pb "github.com/micro/services/blog/tags/proto"
)

func main() {
	// Register Handler
	pb.RegisterTagsHandler(handler.New())

	// Run service
	service.Run()
}
