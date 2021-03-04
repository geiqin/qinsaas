// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cronService.proto

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

// Api Endpoints for CronService service

func NewCronServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CronService service

type CronService interface {
	// 订单超时自动关闭（下单后超过1天未支付）
	OrderClose(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CronServiceResponse, error)
	// 订单超时自动收货（发货后超过7天未收货）
	OrderReceipt(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CronServiceResponse, error)
}

type cronService struct {
	c    client.Client
	name string
}

func NewCronService(name string, c client.Client) CronService {
	return &cronService{
		c:    c,
		name: name,
	}
}

func (c *cronService) OrderClose(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CronServiceResponse, error) {
	req := c.c.NewRequest(c.name, "CronService.OrderClose", in)
	out := new(CronServiceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronService) OrderReceipt(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CronServiceResponse, error) {
	req := c.c.NewRequest(c.name, "CronService.OrderReceipt", in)
	out := new(CronServiceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CronService service

type CronServiceHandler interface {
	// 订单超时自动关闭（下单后超过1天未支付）
	OrderClose(context.Context, *common.Empty, *CronServiceResponse) error
	// 订单超时自动收货（发货后超过7天未收货）
	OrderReceipt(context.Context, *common.Empty, *CronServiceResponse) error
}

func RegisterCronServiceHandler(s server.Server, hdlr CronServiceHandler, opts ...server.HandlerOption) error {
	type cronService interface {
		OrderClose(ctx context.Context, in *common.Empty, out *CronServiceResponse) error
		OrderReceipt(ctx context.Context, in *common.Empty, out *CronServiceResponse) error
	}
	type CronService struct {
		cronService
	}
	h := &cronServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CronService{h}, opts...))
}

type cronServiceHandler struct {
	CronServiceHandler
}

func (h *cronServiceHandler) OrderClose(ctx context.Context, in *common.Empty, out *CronServiceResponse) error {
	return h.CronServiceHandler.OrderClose(ctx, in, out)
}

func (h *cronServiceHandler) OrderReceipt(ctx context.Context, in *common.Empty, out *CronServiceResponse) error {
	return h.CronServiceHandler.OrderReceipt(ctx, in, out)
}