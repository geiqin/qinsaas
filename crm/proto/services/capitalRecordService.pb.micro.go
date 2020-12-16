// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: capitalRecordService.proto

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

// Api Endpoints for CapitalRecordService service

func NewCapitalRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CapitalRecordService service

type CapitalRecordService interface {
	//增加余额（增加）
	Income(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error)
	//消费余额(支出)
	Expend(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error)
	//获得余额记录信息
	Get(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error)
	//查询余额记录信息
	Search(ctx context.Context, in *CapitalRecordWhere, opts ...client.CallOption) (*CapitalRecordResponse, error)
}

type capitalRecordService struct {
	c    client.Client
	name string
}

func NewCapitalRecordService(name string, c client.Client) CapitalRecordService {
	return &capitalRecordService{
		c:    c,
		name: name,
	}
}

func (c *capitalRecordService) Income(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error) {
	req := c.c.NewRequest(c.name, "CapitalRecordService.Income", in)
	out := new(CapitalRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capitalRecordService) Expend(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error) {
	req := c.c.NewRequest(c.name, "CapitalRecordService.Expend", in)
	out := new(CapitalRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capitalRecordService) Get(ctx context.Context, in *CapitalRecord, opts ...client.CallOption) (*CapitalRecordResponse, error) {
	req := c.c.NewRequest(c.name, "CapitalRecordService.Get", in)
	out := new(CapitalRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capitalRecordService) Search(ctx context.Context, in *CapitalRecordWhere, opts ...client.CallOption) (*CapitalRecordResponse, error) {
	req := c.c.NewRequest(c.name, "CapitalRecordService.Search", in)
	out := new(CapitalRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CapitalRecordService service

type CapitalRecordServiceHandler interface {
	//增加余额（增加）
	Income(context.Context, *CapitalRecord, *CapitalRecordResponse) error
	//消费余额(支出)
	Expend(context.Context, *CapitalRecord, *CapitalRecordResponse) error
	//获得余额记录信息
	Get(context.Context, *CapitalRecord, *CapitalRecordResponse) error
	//查询余额记录信息
	Search(context.Context, *CapitalRecordWhere, *CapitalRecordResponse) error
}

func RegisterCapitalRecordServiceHandler(s server.Server, hdlr CapitalRecordServiceHandler, opts ...server.HandlerOption) error {
	type capitalRecordService interface {
		Income(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error
		Expend(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error
		Get(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error
		Search(ctx context.Context, in *CapitalRecordWhere, out *CapitalRecordResponse) error
	}
	type CapitalRecordService struct {
		capitalRecordService
	}
	h := &capitalRecordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CapitalRecordService{h}, opts...))
}

type capitalRecordServiceHandler struct {
	CapitalRecordServiceHandler
}

func (h *capitalRecordServiceHandler) Income(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error {
	return h.CapitalRecordServiceHandler.Income(ctx, in, out)
}

func (h *capitalRecordServiceHandler) Expend(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error {
	return h.CapitalRecordServiceHandler.Expend(ctx, in, out)
}

func (h *capitalRecordServiceHandler) Get(ctx context.Context, in *CapitalRecord, out *CapitalRecordResponse) error {
	return h.CapitalRecordServiceHandler.Get(ctx, in, out)
}

func (h *capitalRecordServiceHandler) Search(ctx context.Context, in *CapitalRecordWhere, out *CapitalRecordResponse) error {
	return h.CapitalRecordServiceHandler.Search(ctx, in, out)
}
