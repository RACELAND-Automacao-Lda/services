// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tunnel.proto

package tunnel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for Tunnel service

func NewTunnelEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Tunnel service

type TunnelService interface {
	Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error)
}

type tunnelService struct {
	c    client.Client
	name string
}

func NewTunnelService(name string, c client.Client) TunnelService {
	return &tunnelService{
		c:    c,
		name: name,
	}
}

func (c *tunnelService) Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error) {
	req := c.c.NewRequest(c.name, "Tunnel.Send", in)
	out := new(SendResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tunnel service

type TunnelHandler interface {
	Send(context.Context, *SendRequest, *SendResponse) error
}

func RegisterTunnelHandler(s server.Server, hdlr TunnelHandler, opts ...server.HandlerOption) error {
	type tunnel interface {
		Send(ctx context.Context, in *SendRequest, out *SendResponse) error
	}
	type Tunnel struct {
		tunnel
	}
	h := &tunnelHandler{hdlr}
	return s.Handle(s.NewHandler(&Tunnel{h}, opts...))
}

type tunnelHandler struct {
	TunnelHandler
}

func (h *tunnelHandler) Send(ctx context.Context, in *SendRequest, out *SendResponse) error {
	return h.TunnelHandler.Send(ctx, in, out)
}