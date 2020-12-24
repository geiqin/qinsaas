// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roomTypeService.proto

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

// Api Endpoints for RoomTypeService service

func NewRoomTypeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoomTypeService service

type RoomTypeService interface {
	Search(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	List(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	Delete(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	Get(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error)
	Create(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error)
	Update(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error)
	ModifySort(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	PricePlanList(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
}

type roomTypeService struct {
	c    client.Client
	name string
}

func NewRoomTypeService(name string, c client.Client) RoomTypeService {
	return &roomTypeService{
		c:    c,
		name: name,
	}
}

func (c *roomTypeService) Search(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.Search", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) List(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.List", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) Delete(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.Delete", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) Get(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.Get", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) Create(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.Create", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) Update(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.Update", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) ModifySort(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.ModifySort", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeService) PricePlanList(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "RoomTypeService.PricePlanList", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomTypeService service

type RoomTypeServiceHandler interface {
	Search(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	List(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	Delete(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	Get(context.Context, *RoomType, *RoomTypeResponse) error
	Create(context.Context, *RoomType, *RoomTypeResponse) error
	Update(context.Context, *RoomType, *RoomTypeResponse) error
	ModifySort(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	PricePlanList(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
}

func RegisterRoomTypeServiceHandler(s server.Server, hdlr RoomTypeServiceHandler, opts ...server.HandlerOption) error {
	type roomTypeService interface {
		Search(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		List(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		Delete(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		Get(ctx context.Context, in *RoomType, out *RoomTypeResponse) error
		Create(ctx context.Context, in *RoomType, out *RoomTypeResponse) error
		Update(ctx context.Context, in *RoomType, out *RoomTypeResponse) error
		ModifySort(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		PricePlanList(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
	}
	type RoomTypeService struct {
		roomTypeService
	}
	h := &roomTypeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomTypeService{h}, opts...))
}

type roomTypeServiceHandler struct {
	RoomTypeServiceHandler
}

func (h *roomTypeServiceHandler) Search(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.Search(ctx, in, out)
}

func (h *roomTypeServiceHandler) List(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.List(ctx, in, out)
}

func (h *roomTypeServiceHandler) Delete(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.Delete(ctx, in, out)
}

func (h *roomTypeServiceHandler) Get(ctx context.Context, in *RoomType, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.Get(ctx, in, out)
}

func (h *roomTypeServiceHandler) Create(ctx context.Context, in *RoomType, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.Create(ctx, in, out)
}

func (h *roomTypeServiceHandler) Update(ctx context.Context, in *RoomType, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.Update(ctx, in, out)
}

func (h *roomTypeServiceHandler) ModifySort(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.ModifySort(ctx, in, out)
}

func (h *roomTypeServiceHandler) PricePlanList(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.RoomTypeServiceHandler.PricePlanList(ctx, in, out)
}

// Api Endpoints for FrontRoomTypeService service

func NewFrontRoomTypeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontRoomTypeService service

type FrontRoomTypeService interface {
	Search(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	List(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error)
	Get(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error)
}

type frontRoomTypeService struct {
	c    client.Client
	name string
}

func NewFrontRoomTypeService(name string, c client.Client) FrontRoomTypeService {
	return &frontRoomTypeService{
		c:    c,
		name: name,
	}
}

func (c *frontRoomTypeService) Search(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "FrontRoomTypeService.Search", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontRoomTypeService) List(ctx context.Context, in *RoomTypeWhere, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "FrontRoomTypeService.List", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontRoomTypeService) Get(ctx context.Context, in *RoomType, opts ...client.CallOption) (*RoomTypeResponse, error) {
	req := c.c.NewRequest(c.name, "FrontRoomTypeService.Get", in)
	out := new(RoomTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontRoomTypeService service

type FrontRoomTypeServiceHandler interface {
	Search(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	List(context.Context, *RoomTypeWhere, *RoomTypeResponse) error
	Get(context.Context, *RoomType, *RoomTypeResponse) error
}

func RegisterFrontRoomTypeServiceHandler(s server.Server, hdlr FrontRoomTypeServiceHandler, opts ...server.HandlerOption) error {
	type frontRoomTypeService interface {
		Search(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		List(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error
		Get(ctx context.Context, in *RoomType, out *RoomTypeResponse) error
	}
	type FrontRoomTypeService struct {
		frontRoomTypeService
	}
	h := &frontRoomTypeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontRoomTypeService{h}, opts...))
}

type frontRoomTypeServiceHandler struct {
	FrontRoomTypeServiceHandler
}

func (h *frontRoomTypeServiceHandler) Search(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.FrontRoomTypeServiceHandler.Search(ctx, in, out)
}

func (h *frontRoomTypeServiceHandler) List(ctx context.Context, in *RoomTypeWhere, out *RoomTypeResponse) error {
	return h.FrontRoomTypeServiceHandler.List(ctx, in, out)
}

func (h *frontRoomTypeServiceHandler) Get(ctx context.Context, in *RoomType, out *RoomTypeResponse) error {
	return h.FrontRoomTypeServiceHandler.Get(ctx, in, out)
}
