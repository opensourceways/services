package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"
	pb "github.com/micro/services/blog/comments/proto"
)

type Comments struct{}

// Call is a single request handler called via client.Call or the generated client code
func (c *Comments) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	logger.Info("Received Comments.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (c *Comments) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Comments_StreamStream) error {
	logger.Infof("Received Comments.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (c *Comments) PingPong(ctx context.Context, stream pb.Comments_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
