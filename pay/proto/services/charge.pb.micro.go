// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: charge.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
	fmt "fmt"
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

// Api Endpoints for ChargeService service

func NewChargeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ChargeService service

type ChargeService interface {
	//删除支付凭证
	Delete(ctx context.Context, in *Charge, opts ...client.CallOption) (*ChargeResponse, error)
	//获得支付凭证
	Get(ctx context.Context, in *Charge, opts ...client.CallOption) (*ChargeResponse, error)
	//查询支付凭证
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ChargeResponse, error)
}

type chargeService struct {
	c    client.Client
	name string
}

func NewChargeService(name string, c client.Client) ChargeService {
	return &chargeService{
		c:    c,
		name: name,
	}
}

func (c *chargeService) Delete(ctx context.Context, in *Charge, opts ...client.CallOption) (*ChargeResponse, error) {
	req := c.c.NewRequest(c.name, "ChargeService.Delete", in)
	out := new(ChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chargeService) Get(ctx context.Context, in *Charge, opts ...client.CallOption) (*ChargeResponse, error) {
	req := c.c.NewRequest(c.name, "ChargeService.Get", in)
	out := new(ChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chargeService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ChargeResponse, error) {
	req := c.c.NewRequest(c.name, "ChargeService.Search", in)
	out := new(ChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChargeService service

type ChargeServiceHandler interface {
	//删除支付凭证
	Delete(context.Context, *Charge, *ChargeResponse) error
	//获得支付凭证
	Get(context.Context, *Charge, *ChargeResponse) error
	//查询支付凭证
	Search(context.Context, *common.BaseWhere, *ChargeResponse) error
}

func RegisterChargeServiceHandler(s server.Server, hdlr ChargeServiceHandler, opts ...server.HandlerOption) error {
	type chargeService interface {
		Delete(ctx context.Context, in *Charge, out *ChargeResponse) error
		Get(ctx context.Context, in *Charge, out *ChargeResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *ChargeResponse) error
	}
	type ChargeService struct {
		chargeService
	}
	h := &chargeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ChargeService{h}, opts...))
}

type chargeServiceHandler struct {
	ChargeServiceHandler
}

func (h *chargeServiceHandler) Delete(ctx context.Context, in *Charge, out *ChargeResponse) error {
	return h.ChargeServiceHandler.Delete(ctx, in, out)
}

func (h *chargeServiceHandler) Get(ctx context.Context, in *Charge, out *ChargeResponse) error {
	return h.ChargeServiceHandler.Get(ctx, in, out)
}

func (h *chargeServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *ChargeResponse) error {
	return h.ChargeServiceHandler.Search(ctx, in, out)
}