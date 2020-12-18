// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: kindTagService.proto

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

// Api Endpoints for KindTagService service

func NewKindTagServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for KindTagService service

type KindTagService interface {
	Search(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error)
	List(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error)
	Get(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error)
	Create(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error)
	Update(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error)
	Delete(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error)
}

type kindTagService struct {
	c    client.Client
	name string
}

func NewKindTagService(name string, c client.Client) KindTagService {
	return &kindTagService{
		c:    c,
		name: name,
	}
}

func (c *kindTagService) Search(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.Search", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindTagService) List(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.List", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindTagService) Get(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.Get", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindTagService) Create(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.Create", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindTagService) Update(ctx context.Context, in *KindTag, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.Update", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kindTagService) Delete(ctx context.Context, in *KindTagWhere, opts ...client.CallOption) (*KindTagResponse, error) {
	req := c.c.NewRequest(c.name, "KindTagService.Delete", in)
	out := new(KindTagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KindTagService service

type KindTagServiceHandler interface {
	Search(context.Context, *KindTagWhere, *KindTagResponse) error
	List(context.Context, *KindTagWhere, *KindTagResponse) error
	Get(context.Context, *KindTag, *KindTagResponse) error
	Create(context.Context, *KindTag, *KindTagResponse) error
	Update(context.Context, *KindTag, *KindTagResponse) error
	Delete(context.Context, *KindTagWhere, *KindTagResponse) error
}

func RegisterKindTagServiceHandler(s server.Server, hdlr KindTagServiceHandler, opts ...server.HandlerOption) error {
	type kindTagService interface {
		Search(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error
		List(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error
		Get(ctx context.Context, in *KindTag, out *KindTagResponse) error
		Create(ctx context.Context, in *KindTag, out *KindTagResponse) error
		Update(ctx context.Context, in *KindTag, out *KindTagResponse) error
		Delete(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error
	}
	type KindTagService struct {
		kindTagService
	}
	h := &kindTagServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&KindTagService{h}, opts...))
}

type kindTagServiceHandler struct {
	KindTagServiceHandler
}

func (h *kindTagServiceHandler) Search(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error {
	return h.KindTagServiceHandler.Search(ctx, in, out)
}

func (h *kindTagServiceHandler) List(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error {
	return h.KindTagServiceHandler.List(ctx, in, out)
}

func (h *kindTagServiceHandler) Get(ctx context.Context, in *KindTag, out *KindTagResponse) error {
	return h.KindTagServiceHandler.Get(ctx, in, out)
}

func (h *kindTagServiceHandler) Create(ctx context.Context, in *KindTag, out *KindTagResponse) error {
	return h.KindTagServiceHandler.Create(ctx, in, out)
}

func (h *kindTagServiceHandler) Update(ctx context.Context, in *KindTag, out *KindTagResponse) error {
	return h.KindTagServiceHandler.Update(ctx, in, out)
}

func (h *kindTagServiceHandler) Delete(ctx context.Context, in *KindTagWhere, out *KindTagResponse) error {
	return h.KindTagServiceHandler.Delete(ctx, in, out)
}
