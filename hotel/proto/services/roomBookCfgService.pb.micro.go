// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roomBookCfgService.proto

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

// Api Endpoints for RoomBookCfgService service

func NewRoomBookCfgServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoomBookCfgService service

type RoomBookCfgService interface {
	Get(ctx context.Context, in *RoomBookCfg, opts ...client.CallOption) (*RoomBookCfgResponse, error)
	Save(ctx context.Context, in *RoomBookCfg, opts ...client.CallOption) (*RoomBookCfgResponse, error)
}

type roomBookCfgService struct {
	c    client.Client
	name string
}

func NewRoomBookCfgService(name string, c client.Client) RoomBookCfgService {
	return &roomBookCfgService{
		c:    c,
		name: name,
	}
}

func (c *roomBookCfgService) Get(ctx context.Context, in *RoomBookCfg, opts ...client.CallOption) (*RoomBookCfgResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookCfgService.Get", in)
	out := new(RoomBookCfgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBookCfgService) Save(ctx context.Context, in *RoomBookCfg, opts ...client.CallOption) (*RoomBookCfgResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookCfgService.Save", in)
	out := new(RoomBookCfgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomBookCfgService service

type RoomBookCfgServiceHandler interface {
	Get(context.Context, *RoomBookCfg, *RoomBookCfgResponse) error
	Save(context.Context, *RoomBookCfg, *RoomBookCfgResponse) error
}

func RegisterRoomBookCfgServiceHandler(s server.Server, hdlr RoomBookCfgServiceHandler, opts ...server.HandlerOption) error {
	type roomBookCfgService interface {
		Get(ctx context.Context, in *RoomBookCfg, out *RoomBookCfgResponse) error
		Save(ctx context.Context, in *RoomBookCfg, out *RoomBookCfgResponse) error
	}
	type RoomBookCfgService struct {
		roomBookCfgService
	}
	h := &roomBookCfgServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomBookCfgService{h}, opts...))
}

type roomBookCfgServiceHandler struct {
	RoomBookCfgServiceHandler
}

func (h *roomBookCfgServiceHandler) Get(ctx context.Context, in *RoomBookCfg, out *RoomBookCfgResponse) error {
	return h.RoomBookCfgServiceHandler.Get(ctx, in, out)
}

func (h *roomBookCfgServiceHandler) Save(ctx context.Context, in *RoomBookCfg, out *RoomBookCfgResponse) error {
	return h.RoomBookCfgServiceHandler.Save(ctx, in, out)
}