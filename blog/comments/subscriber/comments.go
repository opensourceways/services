package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v3/logger"
	pb "github.com/micro/services/blog/comments/proto"
)

type Comments struct{}

func (e *Comments) Handle(ctx context.Context, msg *pb.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *pb.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
