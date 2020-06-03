package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	tag "github.com/micro/examples/blog/tag/proto/tag"
)

type Tag struct{}

func (e *Tag) Handle(ctx context.Context, msg *tag.IncreaseCountRequest) error {
	log.Info("Handler Received message")
	return nil
}

func Handler(ctx context.Context, msg *tag.IncreaseCountRequest) error {
	log.Info("Function Received message: ")
	return nil
}
