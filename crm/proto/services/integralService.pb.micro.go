// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: integralService.proto

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

// Api Endpoints for IntegralService service

func NewIntegralServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IntegralService service

type IntegralService interface {
	Create(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error)
	Delete(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error)
	Get(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*IntegralResponse, error)
}

type integralService struct {
	c    client.Client
	name string
}

func NewIntegralService(name string, c client.Client) IntegralService {
	return &integralService{
		c:    c,
		name: name,
	}
}

func (c *integralService) Create(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Create", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) Delete(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Delete", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) Get(ctx context.Context, in *IntegralResponse, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Get", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Search", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IntegralService service

type IntegralServiceHandler interface {
	Create(context.Context, *IntegralResponse, *IntegralResponse) error
	Delete(context.Context, *IntegralResponse, *IntegralResponse) error
	Get(context.Context, *IntegralResponse, *IntegralResponse) error
	Search(context.Context, *common.BaseWhere, *IntegralResponse) error
}

func RegisterIntegralServiceHandler(s server.Server, hdlr IntegralServiceHandler, opts ...server.HandlerOption) error {
	type integralService interface {
		Create(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error
		Delete(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error
		Get(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *IntegralResponse) error
	}
	type IntegralService struct {
		integralService
	}
	h := &integralServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IntegralService{h}, opts...))
}

type integralServiceHandler struct {
	IntegralServiceHandler
}

func (h *integralServiceHandler) Create(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Create(ctx, in, out)
}

func (h *integralServiceHandler) Delete(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Delete(ctx, in, out)
}

func (h *integralServiceHandler) Get(ctx context.Context, in *IntegralResponse, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Get(ctx, in, out)
}

func (h *integralServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Search(ctx, in, out)
}