// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: paygate.proto

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

// Api Endpoints for PaygateService service

func NewPaygateServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PaygateService service

type PaygateService interface {
	Create(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Update(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Delete(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Lock(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Unlock(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Get(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*PaygateResponse, error)
}

type paygateService struct {
	c    client.Client
	name string
}

func NewPaygateService(name string, c client.Client) PaygateService {
	return &paygateService{
		c:    c,
		name: name,
	}
}

func (c *paygateService) Create(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Create", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Update(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Update", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Delete(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Delete", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Lock(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Lock", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Unlock(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Unlock", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Get(ctx context.Context, in *Paygate, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Get", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paygateService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*PaygateResponse, error) {
	req := c.c.NewRequest(c.name, "PaygateService.Search", in)
	out := new(PaygateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaygateService service

type PaygateServiceHandler interface {
	Create(context.Context, *Paygate, *PaygateResponse) error
	Update(context.Context, *Paygate, *PaygateResponse) error
	Delete(context.Context, *Paygate, *PaygateResponse) error
	Lock(context.Context, *Paygate, *PaygateResponse) error
	Unlock(context.Context, *Paygate, *PaygateResponse) error
	Get(context.Context, *Paygate, *PaygateResponse) error
	Search(context.Context, *common.BaseWhere, *PaygateResponse) error
}

func RegisterPaygateServiceHandler(s server.Server, hdlr PaygateServiceHandler, opts ...server.HandlerOption) error {
	type paygateService interface {
		Create(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Update(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Delete(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Lock(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Unlock(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Get(ctx context.Context, in *Paygate, out *PaygateResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *PaygateResponse) error
	}
	type PaygateService struct {
		paygateService
	}
	h := &paygateServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaygateService{h}, opts...))
}

type paygateServiceHandler struct {
	PaygateServiceHandler
}

func (h *paygateServiceHandler) Create(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Create(ctx, in, out)
}

func (h *paygateServiceHandler) Update(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Update(ctx, in, out)
}

func (h *paygateServiceHandler) Delete(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Delete(ctx, in, out)
}

func (h *paygateServiceHandler) Lock(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Lock(ctx, in, out)
}

func (h *paygateServiceHandler) Unlock(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Unlock(ctx, in, out)
}

func (h *paygateServiceHandler) Get(ctx context.Context, in *Paygate, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Get(ctx, in, out)
}

func (h *paygateServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *PaygateResponse) error {
	return h.PaygateServiceHandler.Search(ctx, in, out)
}
