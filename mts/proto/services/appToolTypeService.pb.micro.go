// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appToolTypeService.proto

package services

import (
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for AppToolTypeService service

func NewAppToolTypeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppToolTypeService service

type AppToolTypeService interface {
	Create(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error)
	Update(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error)
	Delete(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error)
	Get(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error)
	List(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error)
	Search(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error)
	ChangeSort(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error)
	Tree(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error)
}

type appToolTypeService struct {
	c    client.Client
	name string
}

func NewAppToolTypeService(name string, c client.Client) AppToolTypeService {
	return &appToolTypeService{
		c:    c,
		name: name,
	}
}

func (c *appToolTypeService) Create(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Create", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) Update(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Update", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) Delete(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Delete", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) Get(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Get", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) List(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.List", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) Search(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Search", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) ChangeSort(ctx context.Context, in *AppToolType, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.ChangeSort", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appToolTypeService) Tree(ctx context.Context, in *AppToolTypeWhere, opts ...client.CallOption) (*AppToolTypeResponse, error) {
	req := c.c.NewRequest(c.name, "AppToolTypeService.Tree", in)
	out := new(AppToolTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppToolTypeService service

type AppToolTypeServiceHandler interface {
	Create(context.Context, *AppToolType, *AppToolTypeResponse) error
	Update(context.Context, *AppToolType, *AppToolTypeResponse) error
	Delete(context.Context, *AppToolTypeWhere, *AppToolTypeResponse) error
	Get(context.Context, *AppToolType, *AppToolTypeResponse) error
	List(context.Context, *AppToolTypeWhere, *AppToolTypeResponse) error
	Search(context.Context, *AppToolTypeWhere, *AppToolTypeResponse) error
	ChangeSort(context.Context, *AppToolType, *AppToolTypeResponse) error
	Tree(context.Context, *AppToolTypeWhere, *AppToolTypeResponse) error
}

func RegisterAppToolTypeServiceHandler(s server.Server, hdlr AppToolTypeServiceHandler, opts ...server.HandlerOption) error {
	type appToolTypeService interface {
		Create(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error
		Update(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error
		Delete(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error
		Get(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error
		List(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error
		Search(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error
		ChangeSort(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error
		Tree(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error
	}
	type AppToolTypeService struct {
		appToolTypeService
	}
	h := &appToolTypeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AppToolTypeService{h}, opts...))
}

type appToolTypeServiceHandler struct {
	AppToolTypeServiceHandler
}

func (h *appToolTypeServiceHandler) Create(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Create(ctx, in, out)
}

func (h *appToolTypeServiceHandler) Update(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Update(ctx, in, out)
}

func (h *appToolTypeServiceHandler) Delete(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Delete(ctx, in, out)
}

func (h *appToolTypeServiceHandler) Get(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Get(ctx, in, out)
}

func (h *appToolTypeServiceHandler) List(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.List(ctx, in, out)
}

func (h *appToolTypeServiceHandler) Search(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Search(ctx, in, out)
}

func (h *appToolTypeServiceHandler) ChangeSort(ctx context.Context, in *AppToolType, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.ChangeSort(ctx, in, out)
}

func (h *appToolTypeServiceHandler) Tree(ctx context.Context, in *AppToolTypeWhere, out *AppToolTypeResponse) error {
	return h.AppToolTypeServiceHandler.Tree(ctx, in, out)
}
