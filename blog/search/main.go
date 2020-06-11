package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"search/handler"
	"search/subscriber"

	search "search/proto/search"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.search"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	search.RegisterSearchHandler(service.Server(), new(handler.Search))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.search", service.Server(), new(subscriber.Search))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
