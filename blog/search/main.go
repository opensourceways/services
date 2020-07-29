package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/micro/services/blog/search/handler"
	search "github.com/micro/services/blog/search/proto/search"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.search"),
		service.Version("latest"),
	)

	// Register Handler
	search.RegisterSearchServiceHandler(srv.Server(), new(handler.Search))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
