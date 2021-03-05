// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: groupTeamService.proto

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

// Api Endpoints for GroupTeamService service

func NewGroupTeamServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GroupTeamService service

type GroupTeamService interface {
	Create(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
	Search(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
	List(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
	Get(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
	Delete(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
}

type groupTeamService struct {
	c    client.Client
	name string
}

func NewGroupTeamService(name string, c client.Client) GroupTeamService {
	return &groupTeamService{
		c:    c,
		name: name,
	}
}

func (c *groupTeamService) Create(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "GroupTeamService.Create", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupTeamService) Search(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "GroupTeamService.Search", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupTeamService) List(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "GroupTeamService.List", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupTeamService) Get(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "GroupTeamService.Get", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupTeamService) Delete(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "GroupTeamService.Delete", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupTeamService service

type GroupTeamServiceHandler interface {
	Create(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
	Search(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
	List(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
	Get(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
	Delete(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
}

func RegisterGroupTeamServiceHandler(s server.Server, hdlr GroupTeamServiceHandler, opts ...server.HandlerOption) error {
	type groupTeamService interface {
		Create(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
		Search(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
		List(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
		Get(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
		Delete(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
	}
	type GroupTeamService struct {
		groupTeamService
	}
	h := &groupTeamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GroupTeamService{h}, opts...))
}

type groupTeamServiceHandler struct {
	GroupTeamServiceHandler
}

func (h *groupTeamServiceHandler) Create(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.GroupTeamServiceHandler.Create(ctx, in, out)
}

func (h *groupTeamServiceHandler) Search(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.GroupTeamServiceHandler.Search(ctx, in, out)
}

func (h *groupTeamServiceHandler) List(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.GroupTeamServiceHandler.List(ctx, in, out)
}

func (h *groupTeamServiceHandler) Get(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.GroupTeamServiceHandler.Get(ctx, in, out)
}

func (h *groupTeamServiceHandler) Delete(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.GroupTeamServiceHandler.Delete(ctx, in, out)
}

// Api Endpoints for FrontGroupTeamService service

func NewFrontGroupTeamServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontGroupTeamService service

type FrontGroupTeamService interface {
	Get(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error)
}

type frontGroupTeamService struct {
	c    client.Client
	name string
}

func NewFrontGroupTeamService(name string, c client.Client) FrontGroupTeamService {
	return &frontGroupTeamService{
		c:    c,
		name: name,
	}
}

func (c *frontGroupTeamService) Get(ctx context.Context, in *GroupTeamWhere, opts ...client.CallOption) (*GroupTeamResponse, error) {
	req := c.c.NewRequest(c.name, "FrontGroupTeamService.Get", in)
	out := new(GroupTeamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontGroupTeamService service

type FrontGroupTeamServiceHandler interface {
	Get(context.Context, *GroupTeamWhere, *GroupTeamResponse) error
}

func RegisterFrontGroupTeamServiceHandler(s server.Server, hdlr FrontGroupTeamServiceHandler, opts ...server.HandlerOption) error {
	type frontGroupTeamService interface {
		Get(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error
	}
	type FrontGroupTeamService struct {
		frontGroupTeamService
	}
	h := &frontGroupTeamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontGroupTeamService{h}, opts...))
}

type frontGroupTeamServiceHandler struct {
	FrontGroupTeamServiceHandler
}

func (h *frontGroupTeamServiceHandler) Get(ctx context.Context, in *GroupTeamWhere, out *GroupTeamResponse) error {
	return h.FrontGroupTeamServiceHandler.Get(ctx, in, out)
}
