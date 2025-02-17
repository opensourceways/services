// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/etas.proto

package etas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ETAs service

func NewETAsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ETAs service

type ETAsService interface {
	Calculate(ctx context.Context, in *Route, opts ...client.CallOption) (*Response, error)
}

type eTAsService struct {
	c    client.Client
	name string
}

func NewETAsService(name string, c client.Client) ETAsService {
	return &eTAsService{
		c:    c,
		name: name,
	}
}

func (c *eTAsService) Calculate(ctx context.Context, in *Route, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ETAs.Calculate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ETAs service

type ETAsHandler interface {
	Calculate(context.Context, *Route, *Response) error
}

func RegisterETAsHandler(s server.Server, hdlr ETAsHandler, opts ...server.HandlerOption) error {
	type eTAs interface {
		Calculate(ctx context.Context, in *Route, out *Response) error
	}
	type ETAs struct {
		eTAs
	}
	h := &eTAsHandler{hdlr}
	return s.Handle(s.NewHandler(&ETAs{h}, opts...))
}

type eTAsHandler struct {
	ETAsHandler
}

func (h *eTAsHandler) Calculate(ctx context.Context, in *Route, out *Response) error {
	return h.ETAsHandler.Calculate(ctx, in, out)
}
