package subscriber

import (
	"context"

	pb "github.com/micro/examples/blog/comments/proto"
	log "github.com/micro/go-micro/v2/logger"
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
