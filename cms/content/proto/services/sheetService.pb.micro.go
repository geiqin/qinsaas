// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sheetService.proto

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

// Api Endpoints for SheetService service

func NewSheetServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SheetService service

type SheetService interface {
	Create(ctx context.Context, in *Sheet, opts ...client.CallOption) (*SheetResponse, error)
	Update(ctx context.Context, in *Sheet, opts ...client.CallOption) (*SheetResponse, error)
	Delete(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error)
	Get(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error)
	Search(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error)
	List(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error)
}

type sheetService struct {
	c    client.Client
	name string
}

func NewSheetService(name string, c client.Client) SheetService {
	return &sheetService{
		c:    c,
		name: name,
	}
}

func (c *sheetService) Create(ctx context.Context, in *Sheet, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.Create", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) Update(ctx context.Context, in *Sheet, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.Update", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) Delete(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.Delete", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) Get(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.Get", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) Search(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.Search", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) List(ctx context.Context, in *SheetWhere, opts ...client.CallOption) (*SheetResponse, error) {
	req := c.c.NewRequest(c.name, "SheetService.List", in)
	out := new(SheetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SheetService service

type SheetServiceHandler interface {
	Create(context.Context, *Sheet, *SheetResponse) error
	Update(context.Context, *Sheet, *SheetResponse) error
	Delete(context.Context, *SheetWhere, *SheetResponse) error
	Get(context.Context, *SheetWhere, *SheetResponse) error
	Search(context.Context, *SheetWhere, *SheetResponse) error
	List(context.Context, *SheetWhere, *SheetResponse) error
}

func RegisterSheetServiceHandler(s server.Server, hdlr SheetServiceHandler, opts ...server.HandlerOption) error {
	type sheetService interface {
		Create(ctx context.Context, in *Sheet, out *SheetResponse) error
		Update(ctx context.Context, in *Sheet, out *SheetResponse) error
		Delete(ctx context.Context, in *SheetWhere, out *SheetResponse) error
		Get(ctx context.Context, in *SheetWhere, out *SheetResponse) error
		Search(ctx context.Context, in *SheetWhere, out *SheetResponse) error
		List(ctx context.Context, in *SheetWhere, out *SheetResponse) error
	}
	type SheetService struct {
		sheetService
	}
	h := &sheetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SheetService{h}, opts...))
}

type sheetServiceHandler struct {
	SheetServiceHandler
}

func (h *sheetServiceHandler) Create(ctx context.Context, in *Sheet, out *SheetResponse) error {
	return h.SheetServiceHandler.Create(ctx, in, out)
}

func (h *sheetServiceHandler) Update(ctx context.Context, in *Sheet, out *SheetResponse) error {
	return h.SheetServiceHandler.Update(ctx, in, out)
}

func (h *sheetServiceHandler) Delete(ctx context.Context, in *SheetWhere, out *SheetResponse) error {
	return h.SheetServiceHandler.Delete(ctx, in, out)
}

func (h *sheetServiceHandler) Get(ctx context.Context, in *SheetWhere, out *SheetResponse) error {
	return h.SheetServiceHandler.Get(ctx, in, out)
}

func (h *sheetServiceHandler) Search(ctx context.Context, in *SheetWhere, out *SheetResponse) error {
	return h.SheetServiceHandler.Search(ctx, in, out)
}

func (h *sheetServiceHandler) List(ctx context.Context, in *SheetWhere, out *SheetResponse) error {
	return h.SheetServiceHandler.List(ctx, in, out)
}
