// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: wxappRoomGoodsService.proto

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

// Api Endpoints for WxappRoomGoodsService service

func NewWxappRoomGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WxappRoomGoodsService service

type WxappRoomGoodsService interface {
	Create(ctx context.Context, in *WxappRoomGoods, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	ResetAudit(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	Audit(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	Delete(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	Update(ctx context.Context, in *WxappRoomGoods, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	GetGoodsWarehouse(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
	Search(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error)
}

type wxappRoomGoodsService struct {
	c    client.Client
	name string
}

func NewWxappRoomGoodsService(name string, c client.Client) WxappRoomGoodsService {
	return &wxappRoomGoodsService{
		c:    c,
		name: name,
	}
}

func (c *wxappRoomGoodsService) Create(ctx context.Context, in *WxappRoomGoods, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.Create", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) ResetAudit(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.ResetAudit", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) Audit(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.Audit", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) Delete(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.Delete", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) Update(ctx context.Context, in *WxappRoomGoods, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.Update", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) GetGoodsWarehouse(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.GetGoodsWarehouse", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomGoodsService) Search(ctx context.Context, in *WxappRoomGoodsWhere, opts ...client.CallOption) (*WxappRoomGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomGoodsService.Search", in)
	out := new(WxappRoomGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WxappRoomGoodsService service

type WxappRoomGoodsServiceHandler interface {
	Create(context.Context, *WxappRoomGoods, *WxappRoomGoodsResponse) error
	ResetAudit(context.Context, *WxappRoomGoodsWhere, *WxappRoomGoodsResponse) error
	Audit(context.Context, *WxappRoomGoodsWhere, *WxappRoomGoodsResponse) error
	Delete(context.Context, *WxappRoomGoodsWhere, *WxappRoomGoodsResponse) error
	Update(context.Context, *WxappRoomGoods, *WxappRoomGoodsResponse) error
	GetGoodsWarehouse(context.Context, *WxappRoomGoodsWhere, *WxappRoomGoodsResponse) error
	Search(context.Context, *WxappRoomGoodsWhere, *WxappRoomGoodsResponse) error
}

func RegisterWxappRoomGoodsServiceHandler(s server.Server, hdlr WxappRoomGoodsServiceHandler, opts ...server.HandlerOption) error {
	type wxappRoomGoodsService interface {
		Create(ctx context.Context, in *WxappRoomGoods, out *WxappRoomGoodsResponse) error
		ResetAudit(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error
		Audit(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error
		Delete(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error
		Update(ctx context.Context, in *WxappRoomGoods, out *WxappRoomGoodsResponse) error
		GetGoodsWarehouse(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error
		Search(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error
	}
	type WxappRoomGoodsService struct {
		wxappRoomGoodsService
	}
	h := &wxappRoomGoodsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WxappRoomGoodsService{h}, opts...))
}

type wxappRoomGoodsServiceHandler struct {
	WxappRoomGoodsServiceHandler
}

func (h *wxappRoomGoodsServiceHandler) Create(ctx context.Context, in *WxappRoomGoods, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.Create(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) ResetAudit(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.ResetAudit(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) Audit(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.Audit(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) Delete(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.Delete(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) Update(ctx context.Context, in *WxappRoomGoods, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.Update(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) GetGoodsWarehouse(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.GetGoodsWarehouse(ctx, in, out)
}

func (h *wxappRoomGoodsServiceHandler) Search(ctx context.Context, in *WxappRoomGoodsWhere, out *WxappRoomGoodsResponse) error {
	return h.WxappRoomGoodsServiceHandler.Search(ctx, in, out)
}