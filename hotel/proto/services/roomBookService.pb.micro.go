// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roomBookService.proto

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

// Api Endpoints for RoomBookService service

func NewRoomBookServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoomBookService service

type RoomBookService interface {
	Create(ctx context.Context, in *RoomBook, opts ...client.CallOption) (*RoomBookResponse, error)
	Update(ctx context.Context, in *RoomBook, opts ...client.CallOption) (*RoomBookResponse, error)
	// 获取预订详情
	Get(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*RoomBookResponse, error)
	// 获取可预定的房型房号
	GetAllRoomTypeAndAvailableRoom(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*AllRoomTypeAndAvailableRoomResponse, error)
	// 获取可入住|排房的房号列表(入住、排房时选择房号)
	GetAvailableRoom(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*RoomResponse, error)
}

type roomBookService struct {
	c    client.Client
	name string
}

func NewRoomBookService(name string, c client.Client) RoomBookService {
	return &roomBookService{
		c:    c,
		name: name,
	}
}

func (c *roomBookService) Create(ctx context.Context, in *RoomBook, opts ...client.CallOption) (*RoomBookResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookService.Create", in)
	out := new(RoomBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBookService) Update(ctx context.Context, in *RoomBook, opts ...client.CallOption) (*RoomBookResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookService.Update", in)
	out := new(RoomBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBookService) Get(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*RoomBookResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookService.Get", in)
	out := new(RoomBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBookService) GetAllRoomTypeAndAvailableRoom(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*AllRoomTypeAndAvailableRoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookService.GetAllRoomTypeAndAvailableRoom", in)
	out := new(AllRoomTypeAndAvailableRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomBookService) GetAvailableRoom(ctx context.Context, in *RoomBookWhere, opts ...client.CallOption) (*RoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomBookService.GetAvailableRoom", in)
	out := new(RoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomBookService service

type RoomBookServiceHandler interface {
	Create(context.Context, *RoomBook, *RoomBookResponse) error
	Update(context.Context, *RoomBook, *RoomBookResponse) error
	// 获取预订详情
	Get(context.Context, *RoomBookWhere, *RoomBookResponse) error
	// 获取可预定的房型房号
	GetAllRoomTypeAndAvailableRoom(context.Context, *RoomBookWhere, *AllRoomTypeAndAvailableRoomResponse) error
	// 获取可入住|排房的房号列表(入住、排房时选择房号)
	GetAvailableRoom(context.Context, *RoomBookWhere, *RoomResponse) error
}

func RegisterRoomBookServiceHandler(s server.Server, hdlr RoomBookServiceHandler, opts ...server.HandlerOption) error {
	type roomBookService interface {
		Create(ctx context.Context, in *RoomBook, out *RoomBookResponse) error
		Update(ctx context.Context, in *RoomBook, out *RoomBookResponse) error
		Get(ctx context.Context, in *RoomBookWhere, out *RoomBookResponse) error
		GetAllRoomTypeAndAvailableRoom(ctx context.Context, in *RoomBookWhere, out *AllRoomTypeAndAvailableRoomResponse) error
		GetAvailableRoom(ctx context.Context, in *RoomBookWhere, out *RoomResponse) error
	}
	type RoomBookService struct {
		roomBookService
	}
	h := &roomBookServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomBookService{h}, opts...))
}

type roomBookServiceHandler struct {
	RoomBookServiceHandler
}

func (h *roomBookServiceHandler) Create(ctx context.Context, in *RoomBook, out *RoomBookResponse) error {
	return h.RoomBookServiceHandler.Create(ctx, in, out)
}

func (h *roomBookServiceHandler) Update(ctx context.Context, in *RoomBook, out *RoomBookResponse) error {
	return h.RoomBookServiceHandler.Update(ctx, in, out)
}

func (h *roomBookServiceHandler) Get(ctx context.Context, in *RoomBookWhere, out *RoomBookResponse) error {
	return h.RoomBookServiceHandler.Get(ctx, in, out)
}

func (h *roomBookServiceHandler) GetAllRoomTypeAndAvailableRoom(ctx context.Context, in *RoomBookWhere, out *AllRoomTypeAndAvailableRoomResponse) error {
	return h.RoomBookServiceHandler.GetAllRoomTypeAndAvailableRoom(ctx, in, out)
}

func (h *roomBookServiceHandler) GetAvailableRoom(ctx context.Context, in *RoomBookWhere, out *RoomResponse) error {
	return h.RoomBookServiceHandler.GetAvailableRoom(ctx, in, out)
}
