// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderConfigService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/microkit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for OrderConfigService service

func NewOrderConfigServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderConfigService service

type OrderConfigService interface {
	Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*OrderConfigResponse, error)
	Save(ctx context.Context, in *OrderConfig, opts ...client.CallOption) (*OrderConfigResponse, error)
}

type orderConfigService struct {
	c    client.Client
	name string
}

func NewOrderConfigService(name string, c client.Client) OrderConfigService {
	return &orderConfigService{
		c:    c,
		name: name,
	}
}

func (c *orderConfigService) Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*OrderConfigResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfigService.Get", in)
	out := new(OrderConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderConfigService) Save(ctx context.Context, in *OrderConfig, opts ...client.CallOption) (*OrderConfigResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfigService.Save", in)
	out := new(OrderConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderConfigService service

type OrderConfigServiceHandler interface {
	Get(context.Context, *common.Empty, *OrderConfigResponse) error
	Save(context.Context, *OrderConfig, *OrderConfigResponse) error
}

func RegisterOrderConfigServiceHandler(s server.Server, hdlr OrderConfigServiceHandler, opts ...server.HandlerOption) error {
	type orderConfigService interface {
		Get(ctx context.Context, in *common.Empty, out *OrderConfigResponse) error
		Save(ctx context.Context, in *OrderConfig, out *OrderConfigResponse) error
	}
	type OrderConfigService struct {
		orderConfigService
	}
	h := &orderConfigServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderConfigService{h}, opts...))
}

type orderConfigServiceHandler struct {
	OrderConfigServiceHandler
}

func (h *orderConfigServiceHandler) Get(ctx context.Context, in *common.Empty, out *OrderConfigResponse) error {
	return h.OrderConfigServiceHandler.Get(ctx, in, out)
}

func (h *orderConfigServiceHandler) Save(ctx context.Context, in *OrderConfig, out *OrderConfigResponse) error {
	return h.OrderConfigServiceHandler.Save(ctx, in, out)
}
