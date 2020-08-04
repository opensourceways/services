package handler

import (
	"context"

	"github.com/micro/go-micro/v3/util/log"
	pb "github.com/micro/services/helloworld/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (h *Helloworld) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Log("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
