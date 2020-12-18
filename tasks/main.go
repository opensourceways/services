package main

import (
	"tasks/handler"
	pb "tasks/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tasks"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTasksHandler(srv.Server(), new(handler.Tasks))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
