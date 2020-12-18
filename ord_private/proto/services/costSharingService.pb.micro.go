// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: costSharingService.proto

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

// Api Endpoints for CostSharingService service

func NewCostSharingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CostSharingService service

type CostSharingService interface {
	// 计算订单明细分摊后的总金额
	Calculate(ctx context.Context, in *CostSharingWhere, opts ...client.CallOption) (*CostSharingResponse, error)
}

type costSharingService struct {
	c    client.Client
	name string
}

func NewCostSharingService(name string, c client.Client) CostSharingService {
	return &costSharingService{
		c:    c,
		name: name,
	}
}

func (c *costSharingService) Calculate(ctx context.Context, in *CostSharingWhere, opts ...client.CallOption) (*CostSharingResponse, error) {
	req := c.c.NewRequest(c.name, "CostSharingService.Calculate", in)
	out := new(CostSharingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CostSharingService service

type CostSharingServiceHandler interface {
	// 计算订单明细分摊后的总金额
	Calculate(context.Context, *CostSharingWhere, *CostSharingResponse) error
}

func RegisterCostSharingServiceHandler(s server.Server, hdlr CostSharingServiceHandler, opts ...server.HandlerOption) error {
	type costSharingService interface {
		Calculate(ctx context.Context, in *CostSharingWhere, out *CostSharingResponse) error
	}
	type CostSharingService struct {
		costSharingService
	}
	h := &costSharingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CostSharingService{h}, opts...))
}

type costSharingServiceHandler struct {
	CostSharingServiceHandler
}

func (h *costSharingServiceHandler) Calculate(ctx context.Context, in *CostSharingWhere, out *CostSharingResponse) error {
	return h.CostSharingServiceHandler.Calculate(ctx, in, out)
}
