package main

import (
	"github.com/micro/micro/v3/service"

	"github.com/micro/services/blog/comments/handler"
	pb "github.com/micro/services/blog/comments/proto"
)

func main() {
	// Register Handler
	pb.RegisterCommentsHandler(handler.New())

	// Run service
	service.Run()
}
