// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: managementService.proto

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

// Api Endpoints for ManagementService service

func NewManagementServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ManagementService service

type ManagementService interface {
	Add(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error)
	Update(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error)
	Remove(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error)
	Get(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error)
	Flag(ctx context.Context, in *Management, opts ...client.CallOption) (*FlagResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ManagementResponse, error)
}

type managementService struct {
	c    client.Client
	name string
}

func NewManagementService(name string, c client.Client) ManagementService {
	return &managementService{
		c:    c,
		name: name,
	}
}

func (c *managementService) Add(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Add", in)
	out := new(ManagementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementService) Update(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Update", in)
	out := new(ManagementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementService) Remove(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Remove", in)
	out := new(ManagementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementService) Get(ctx context.Context, in *Management, opts ...client.CallOption) (*ManagementResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Get", in)
	out := new(ManagementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementService) Flag(ctx context.Context, in *Management, opts ...client.CallOption) (*FlagResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Flag", in)
	out := new(FlagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ManagementResponse, error) {
	req := c.c.NewRequest(c.name, "ManagementService.Search", in)
	out := new(ManagementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagementService service

type ManagementServiceHandler interface {
	Add(context.Context, *Management, *ManagementResponse) error
	Update(context.Context, *Management, *ManagementResponse) error
	Remove(context.Context, *Management, *ManagementResponse) error
	Get(context.Context, *Management, *ManagementResponse) error
	Flag(context.Context, *Management, *FlagResponse) error
	Search(context.Context, *common.BaseWhere, *ManagementResponse) error
}

func RegisterManagementServiceHandler(s server.Server, hdlr ManagementServiceHandler, opts ...server.HandlerOption) error {
	type managementService interface {
		Add(ctx context.Context, in *Management, out *ManagementResponse) error
		Update(ctx context.Context, in *Management, out *ManagementResponse) error
		Remove(ctx context.Context, in *Management, out *ManagementResponse) error
		Get(ctx context.Context, in *Management, out *ManagementResponse) error
		Flag(ctx context.Context, in *Management, out *FlagResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *ManagementResponse) error
	}
	type ManagementService struct {
		managementService
	}
	h := &managementServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ManagementService{h}, opts...))
}

type managementServiceHandler struct {
	ManagementServiceHandler
}

func (h *managementServiceHandler) Add(ctx context.Context, in *Management, out *ManagementResponse) error {
	return h.ManagementServiceHandler.Add(ctx, in, out)
}

func (h *managementServiceHandler) Update(ctx context.Context, in *Management, out *ManagementResponse) error {
	return h.ManagementServiceHandler.Update(ctx, in, out)
}

func (h *managementServiceHandler) Remove(ctx context.Context, in *Management, out *ManagementResponse) error {
	return h.ManagementServiceHandler.Remove(ctx, in, out)
}

func (h *managementServiceHandler) Get(ctx context.Context, in *Management, out *ManagementResponse) error {
	return h.ManagementServiceHandler.Get(ctx, in, out)
}

func (h *managementServiceHandler) Flag(ctx context.Context, in *Management, out *FlagResponse) error {
	return h.ManagementServiceHandler.Flag(ctx, in, out)
}

func (h *managementServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *ManagementResponse) error {
	return h.ManagementServiceHandler.Search(ctx, in, out)
}
