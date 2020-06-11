package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	search "search/proto/search"
)

type Search struct{}

func (e *Search) Handle(ctx context.Context, msg *search.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *search.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
