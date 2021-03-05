// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: buyingService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for BuyingService service

func NewBuyingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BuyingService service

type BuyingService interface {
	// 拼团确认/提交订单时调用
	Calculate(ctx context.Context, in *BuyingReq, opts ...client.CallOption) (*BuyingResponse, error)
	// 拼团订单创建后调用
	AfterCreateOrder(ctx context.Context, in *BuyingReq, opts ...client.CallOption) (*BuyingResponse, error)
}

type buyingService struct {
	c    client.Client
	name string
}

func NewBuyingService(name string, c client.Client) BuyingService {
	return &buyingService{
		c:    c,
		name: name,
	}
}

func (c *buyingService) Calculate(ctx context.Context, in *BuyingReq, opts ...client.CallOption) (*BuyingResponse, error) {
	req := c.c.NewRequest(c.name, "BuyingService.Calculate", in)
	out := new(BuyingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyingService) AfterCreateOrder(ctx context.Context, in *BuyingReq, opts ...client.CallOption) (*BuyingResponse, error) {
	req := c.c.NewRequest(c.name, "BuyingService.AfterCreateOrder", in)
	out := new(BuyingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BuyingService service

type BuyingServiceHandler interface {
	// 拼团确认/提交订单时调用
	Calculate(context.Context, *BuyingReq, *BuyingResponse) error
	// 拼团订单创建后调用
	AfterCreateOrder(context.Context, *BuyingReq, *BuyingResponse) error
}

func RegisterBuyingServiceHandler(s server.Server, hdlr BuyingServiceHandler, opts ...server.HandlerOption) error {
	type buyingService interface {
		Calculate(ctx context.Context, in *BuyingReq, out *BuyingResponse) error
		AfterCreateOrder(ctx context.Context, in *BuyingReq, out *BuyingResponse) error
	}
	type BuyingService struct {
		buyingService
	}
	h := &buyingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BuyingService{h}, opts...))
}

type buyingServiceHandler struct {
	BuyingServiceHandler
}

func (h *buyingServiceHandler) Calculate(ctx context.Context, in *BuyingReq, out *BuyingResponse) error {
	return h.BuyingServiceHandler.Calculate(ctx, in, out)
}

func (h *buyingServiceHandler) AfterCreateOrder(ctx context.Context, in *BuyingReq, out *BuyingResponse) error {
	return h.BuyingServiceHandler.AfterCreateOrder(ctx, in, out)
}
