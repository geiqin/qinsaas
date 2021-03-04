// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: metaClassService.proto

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

// Api Endpoints for MetaClassService service

func NewMetaClassServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MetaClassService service

type MetaClassService interface {
	Create(ctx context.Context, in *MetaClass, opts ...client.CallOption) (*MetaClassResponse, error)
	Update(ctx context.Context, in *MetaClass, opts ...client.CallOption) (*MetaClassResponse, error)
	Search(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error)
	List(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error)
	Delete(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error)
	Get(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error)
}

type metaClassService struct {
	c    client.Client
	name string
}

func NewMetaClassService(name string, c client.Client) MetaClassService {
	return &metaClassService{
		c:    c,
		name: name,
	}
}

func (c *metaClassService) Create(ctx context.Context, in *MetaClass, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.Create", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClassService) Update(ctx context.Context, in *MetaClass, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.Update", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClassService) Search(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.Search", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClassService) List(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.List", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClassService) Delete(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.Delete", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClassService) Get(ctx context.Context, in *MetaClassWhere, opts ...client.CallOption) (*MetaClassResponse, error) {
	req := c.c.NewRequest(c.name, "MetaClassService.Get", in)
	out := new(MetaClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetaClassService service

type MetaClassServiceHandler interface {
	Create(context.Context, *MetaClass, *MetaClassResponse) error
	Update(context.Context, *MetaClass, *MetaClassResponse) error
	Search(context.Context, *MetaClassWhere, *MetaClassResponse) error
	List(context.Context, *MetaClassWhere, *MetaClassResponse) error
	Delete(context.Context, *MetaClassWhere, *MetaClassResponse) error
	Get(context.Context, *MetaClassWhere, *MetaClassResponse) error
}

func RegisterMetaClassServiceHandler(s server.Server, hdlr MetaClassServiceHandler, opts ...server.HandlerOption) error {
	type metaClassService interface {
		Create(ctx context.Context, in *MetaClass, out *MetaClassResponse) error
		Update(ctx context.Context, in *MetaClass, out *MetaClassResponse) error
		Search(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error
		List(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error
		Delete(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error
		Get(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error
	}
	type MetaClassService struct {
		metaClassService
	}
	h := &metaClassServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MetaClassService{h}, opts...))
}

type metaClassServiceHandler struct {
	MetaClassServiceHandler
}

func (h *metaClassServiceHandler) Create(ctx context.Context, in *MetaClass, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.Create(ctx, in, out)
}

func (h *metaClassServiceHandler) Update(ctx context.Context, in *MetaClass, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.Update(ctx, in, out)
}

func (h *metaClassServiceHandler) Search(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.Search(ctx, in, out)
}

func (h *metaClassServiceHandler) List(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.List(ctx, in, out)
}

func (h *metaClassServiceHandler) Delete(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.Delete(ctx, in, out)
}

func (h *metaClassServiceHandler) Get(ctx context.Context, in *MetaClassWhere, out *MetaClassResponse) error {
	return h.MetaClassServiceHandler.Get(ctx, in, out)
}