// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: wxappRoomSubAnchorService.proto

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

// Api Endpoints for WxappRoomSubAnchorService service

func NewWxappRoomSubAnchorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WxappRoomSubAnchorService service

type WxappRoomSubAnchorService interface {
	Create(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error)
	Update(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error)
	Delete(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error)
	Get(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error)
}

type wxappRoomSubAnchorService struct {
	c    client.Client
	name string
}

func NewWxappRoomSubAnchorService(name string, c client.Client) WxappRoomSubAnchorService {
	return &wxappRoomSubAnchorService{
		c:    c,
		name: name,
	}
}

func (c *wxappRoomSubAnchorService) Create(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomSubAnchorService.Create", in)
	out := new(WxappRoomSubAnchorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomSubAnchorService) Update(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomSubAnchorService.Update", in)
	out := new(WxappRoomSubAnchorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomSubAnchorService) Delete(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomSubAnchorService.Delete", in)
	out := new(WxappRoomSubAnchorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxappRoomSubAnchorService) Get(ctx context.Context, in *WxappRoomSubAnchorWhere, opts ...client.CallOption) (*WxappRoomSubAnchorResponse, error) {
	req := c.c.NewRequest(c.name, "WxappRoomSubAnchorService.Get", in)
	out := new(WxappRoomSubAnchorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WxappRoomSubAnchorService service

type WxappRoomSubAnchorServiceHandler interface {
	Create(context.Context, *WxappRoomSubAnchorWhere, *WxappRoomSubAnchorResponse) error
	Update(context.Context, *WxappRoomSubAnchorWhere, *WxappRoomSubAnchorResponse) error
	Delete(context.Context, *WxappRoomSubAnchorWhere, *WxappRoomSubAnchorResponse) error
	Get(context.Context, *WxappRoomSubAnchorWhere, *WxappRoomSubAnchorResponse) error
}

func RegisterWxappRoomSubAnchorServiceHandler(s server.Server, hdlr WxappRoomSubAnchorServiceHandler, opts ...server.HandlerOption) error {
	type wxappRoomSubAnchorService interface {
		Create(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error
		Update(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error
		Delete(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error
		Get(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error
	}
	type WxappRoomSubAnchorService struct {
		wxappRoomSubAnchorService
	}
	h := &wxappRoomSubAnchorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WxappRoomSubAnchorService{h}, opts...))
}

type wxappRoomSubAnchorServiceHandler struct {
	WxappRoomSubAnchorServiceHandler
}

func (h *wxappRoomSubAnchorServiceHandler) Create(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error {
	return h.WxappRoomSubAnchorServiceHandler.Create(ctx, in, out)
}

func (h *wxappRoomSubAnchorServiceHandler) Update(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error {
	return h.WxappRoomSubAnchorServiceHandler.Update(ctx, in, out)
}

func (h *wxappRoomSubAnchorServiceHandler) Delete(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error {
	return h.WxappRoomSubAnchorServiceHandler.Delete(ctx, in, out)
}

func (h *wxappRoomSubAnchorServiceHandler) Get(ctx context.Context, in *WxappRoomSubAnchorWhere, out *WxappRoomSubAnchorResponse) error {
	return h.WxappRoomSubAnchorServiceHandler.Get(ctx, in, out)
}
