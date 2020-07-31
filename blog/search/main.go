package main

import (
	"github.com/micro/micro/v3/service"

	"github.com/micro/services/blog/search/handler"
	pb "github.com/micro/services/blog/search/proto"
)

func main() {
	// Register Handler
	pb.RegisterSearchHandler(handler.New())

	// Run service
	service.Run()
}
