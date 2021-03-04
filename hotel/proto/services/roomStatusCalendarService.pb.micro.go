// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roomStatusCalendarService.proto

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

// Api Endpoints for RoomStatusCalendarService service

func NewRoomStatusCalendarServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoomStatusCalendarService service

type RoomStatusCalendarService interface {
	List(ctx context.Context, in *RoomStatusCalendarWhere, opts ...client.CallOption) (*RoomStatusCalendarResponse, error)
	Today(ctx context.Context, in *RoomStatusCalendarWhere, opts ...client.CallOption) (*RoomStatusCalendarResponse, error)
}

type roomStatusCalendarService struct {
	c    client.Client
	name string
}

func NewRoomStatusCalendarService(name string, c client.Client) RoomStatusCalendarService {
	return &roomStatusCalendarService{
		c:    c,
		name: name,
	}
}

func (c *roomStatusCalendarService) List(ctx context.Context, in *RoomStatusCalendarWhere, opts ...client.CallOption) (*RoomStatusCalendarResponse, error) {
	req := c.c.NewRequest(c.name, "RoomStatusCalendarService.List", in)
	out := new(RoomStatusCalendarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomStatusCalendarService) Today(ctx context.Context, in *RoomStatusCalendarWhere, opts ...client.CallOption) (*RoomStatusCalendarResponse, error) {
	req := c.c.NewRequest(c.name, "RoomStatusCalendarService.Today", in)
	out := new(RoomStatusCalendarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomStatusCalendarService service

type RoomStatusCalendarServiceHandler interface {
	List(context.Context, *RoomStatusCalendarWhere, *RoomStatusCalendarResponse) error
	Today(context.Context, *RoomStatusCalendarWhere, *RoomStatusCalendarResponse) error
}

func RegisterRoomStatusCalendarServiceHandler(s server.Server, hdlr RoomStatusCalendarServiceHandler, opts ...server.HandlerOption) error {
	type roomStatusCalendarService interface {
		List(ctx context.Context, in *RoomStatusCalendarWhere, out *RoomStatusCalendarResponse) error
		Today(ctx context.Context, in *RoomStatusCalendarWhere, out *RoomStatusCalendarResponse) error
	}
	type RoomStatusCalendarService struct {
		roomStatusCalendarService
	}
	h := &roomStatusCalendarServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomStatusCalendarService{h}, opts...))
}

type roomStatusCalendarServiceHandler struct {
	RoomStatusCalendarServiceHandler
}

func (h *roomStatusCalendarServiceHandler) List(ctx context.Context, in *RoomStatusCalendarWhere, out *RoomStatusCalendarResponse) error {
	return h.RoomStatusCalendarServiceHandler.List(ctx, in, out)
}

func (h *roomStatusCalendarServiceHandler) Today(ctx context.Context, in *RoomStatusCalendarWhere, out *RoomStatusCalendarResponse) error {
	return h.RoomStatusCalendarServiceHandler.Today(ctx, in, out)
}