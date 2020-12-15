// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: menu.proto

package services

import (
	common "geiqin.saas.mts/app/proto/common"
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

// Api Endpoints for MenuService service

func NewMenuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MenuService service

type MenuService interface {
	Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*MenuResponse, error)
	List(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Tree(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	StoreTree(ctx context.Context, in *MenuWhere, opts ...client.CallOption) (*MenuResponse, error)
}

type menuService struct {
	c    client.Client
	name string
}

func NewMenuService(name string, c client.Client) MenuService {
	return &menuService{
		c:    c,
		name: name,
	}
}

func (c *menuService) Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Get", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Search", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) List(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.List", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Tree(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Tree", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) StoreTree(ctx context.Context, in *MenuWhere, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.StoreTree", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MenuService service

type MenuServiceHandler interface {
	Get(context.Context, *Menu, *MenuResponse) error
	Search(context.Context, *common.BaseWhere, *MenuResponse) error
	List(context.Context, *Menu, *MenuResponse) error
	Tree(context.Context, *Menu, *MenuResponse) error
	StoreTree(context.Context, *MenuWhere, *MenuResponse) error
}

func RegisterMenuServiceHandler(s server.Server, hdlr MenuServiceHandler, opts ...server.HandlerOption) error {
	type menuService interface {
		Get(ctx context.Context, in *Menu, out *MenuResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *MenuResponse) error
		List(ctx context.Context, in *Menu, out *MenuResponse) error
		Tree(ctx context.Context, in *Menu, out *MenuResponse) error
		StoreTree(ctx context.Context, in *MenuWhere, out *MenuResponse) error
	}
	type MenuService struct {
		menuService
	}
	h := &menuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MenuService{h}, opts...))
}

type menuServiceHandler struct {
	MenuServiceHandler
}

func (h *menuServiceHandler) Get(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Get(ctx, in, out)
}

func (h *menuServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *MenuResponse) error {
	return h.MenuServiceHandler.Search(ctx, in, out)
}

func (h *menuServiceHandler) List(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.List(ctx, in, out)
}

func (h *menuServiceHandler) Tree(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Tree(ctx, in, out)
}

func (h *menuServiceHandler) StoreTree(ctx context.Context, in *MenuWhere, out *MenuResponse) error {
	return h.MenuServiceHandler.StoreTree(ctx, in, out)
}
