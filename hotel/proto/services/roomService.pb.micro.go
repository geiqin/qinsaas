// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roomService.proto

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

// Api Endpoints for RoomService service

func NewRoomServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoomService service

type RoomService interface {
	Search(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
	List(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
	Delete(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
	Get(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error)
	Create(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error)
	Update(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error)
	SetStatus(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
	SetCleanStatus(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
	ModifySort(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error)
}

type roomService struct {
	c    client.Client
	name string
}

func NewRoomService(name string, c client.Client) RoomService {
	return &roomService{
		c:    c,
		name: name,
	}
}

func (c *roomService) Search(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.Search", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) List(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.List", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) Delete(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.Delete", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) Get(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.Get", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) Create(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.Create", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) Update(ctx context.Context, in *Room, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.Update", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) SetStatus(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.SetStatus", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) SetCleanStatus(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.SetCleanStatus", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) ModifySort(ctx context.Context, in *RoomWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomService.ModifySort", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomService service

type RoomServiceHandler interface {
	Search(context.Context, *RoomWhere, *RoomResponse) error
	List(context.Context, *RoomWhere, *RoomResponse) error
	Delete(context.Context, *RoomWhere, *RoomResponse) error
	Get(context.Context, *Room, *RoomResponse) error
	Create(context.Context, *Room, *RoomResponse) error
	Update(context.Context, *Room, *RoomResponse) error
	SetStatus(context.Context, *RoomWhere, *RoomResponse) error
	SetCleanStatus(context.Context, *RoomWhere, *RoomResponse) error
	ModifySort(context.Context, *RoomWhere, *RoomResponse) error
}

func RegisterRoomServiceHandler(s server.Server, hdlr RoomServiceHandler, opts ...server.HandlerOption) error {
	type roomService interface {
		Search(ctx context.Context, in *RoomWhere, out *RoomResponse) error
		List(ctx context.Context, in *RoomWhere, out *RoomResponse) error
		Delete(ctx context.Context, in *RoomWhere, out *RoomResponse) error
		Get(ctx context.Context, in *Room, out *RoomResponse) error
		Create(ctx context.Context, in *Room, out *RoomResponse) error
		Update(ctx context.Context, in *Room, out *RoomResponse) error
		SetStatus(ctx context.Context, in *RoomWhere, out *RoomResponse) error
		SetCleanStatus(ctx context.Context, in *RoomWhere, out *RoomResponse) error
		ModifySort(ctx context.Context, in *RoomWhere, out *RoomResponse) error
	}
	type RoomService struct {
		roomService
	}
	h := &roomServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomService{h}, opts...))
}

type roomServiceHandler struct {
	RoomServiceHandler
}

func (h *roomServiceHandler) Search(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.Search(ctx, in, out)
}

func (h *roomServiceHandler) List(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.List(ctx, in, out)
}

func (h *roomServiceHandler) Delete(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.Delete(ctx, in, out)
}

func (h *roomServiceHandler) Get(ctx context.Context, in *Room, out *RoomResponse) error {
	return h.RoomServiceHandler.Get(ctx, in, out)
}

func (h *roomServiceHandler) Create(ctx context.Context, in *Room, out *RoomResponse) error {
	return h.RoomServiceHandler.Create(ctx, in, out)
}

func (h *roomServiceHandler) Update(ctx context.Context, in *Room, out *RoomResponse) error {
	return h.RoomServiceHandler.Update(ctx, in, out)
}

func (h *roomServiceHandler) SetStatus(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.SetStatus(ctx, in, out)
}

func (h *roomServiceHandler) SetCleanStatus(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.SetCleanStatus(ctx, in, out)
}

func (h *roomServiceHandler) ModifySort(ctx context.Context, in *RoomWhere, out *RoomResponse) error {
	return h.RoomServiceHandler.ModifySort(ctx, in, out)
}
