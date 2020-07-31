package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"
	pb "github.com/micro/services/blog/search/proto"
)

func New() pb.SearchHandler {
	return new(search)
}

type search struct{}

func (s *search) Index(ctx context.Context, req *pb.IndexRequest, rsp *pb.IndexResponse) error {
	logger.Info("Received Search.Index request")
	return nil
}

func (s *search) Search(ctx context.Context, req *pb.SearchRequest, rsp *pb.SearchResponse) error {
	logger.Info("Received Search.Search request")
	return nil
}
