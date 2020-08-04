package subscriber

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"
	tags "github.com/micro/services/blog/tags/proto/tags"
)

type Tags struct{}

func (e *Tags) Handle(ctx context.Context, msg *tags.IncreaseCountRequest) error {
	log.Info("Handler Received message")
	return nil
}

func Handler(ctx context.Context, msg *tags.IncreaseCountRequest) error {
	log.Info("Function Received message: ")
	return nil
}
