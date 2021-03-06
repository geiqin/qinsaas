// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: codeStatsDetailService.proto

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

// Api Endpoints for CodeStatsDetailService service

func NewCodeStatsDetailServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CodeStatsDetailService service

type CodeStatsDetailService interface {
	Get(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error)
	Search(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error)
	List(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error)
}

type codeStatsDetailService struct {
	c    client.Client
	name string
}

func NewCodeStatsDetailService(name string, c client.Client) CodeStatsDetailService {
	return &codeStatsDetailService{
		c:    c,
		name: name,
	}
}

func (c *codeStatsDetailService) Get(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error) {
	req := c.c.NewRequest(c.name, "CodeStatsDetailService.Get", in)
	out := new(CodeStatsDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeStatsDetailService) Search(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error) {
	req := c.c.NewRequest(c.name, "CodeStatsDetailService.Search", in)
	out := new(CodeStatsDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeStatsDetailService) List(ctx context.Context, in *CodeStatsDetailWhere, opts ...client.CallOption) (*CodeStatsDetailResponse, error) {
	req := c.c.NewRequest(c.name, "CodeStatsDetailService.List", in)
	out := new(CodeStatsDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CodeStatsDetailService service

type CodeStatsDetailServiceHandler interface {
	Get(context.Context, *CodeStatsDetailWhere, *CodeStatsDetailResponse) error
	Search(context.Context, *CodeStatsDetailWhere, *CodeStatsDetailResponse) error
	List(context.Context, *CodeStatsDetailWhere, *CodeStatsDetailResponse) error
}

func RegisterCodeStatsDetailServiceHandler(s server.Server, hdlr CodeStatsDetailServiceHandler, opts ...server.HandlerOption) error {
	type codeStatsDetailService interface {
		Get(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error
		Search(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error
		List(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error
	}
	type CodeStatsDetailService struct {
		codeStatsDetailService
	}
	h := &codeStatsDetailServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CodeStatsDetailService{h}, opts...))
}

type codeStatsDetailServiceHandler struct {
	CodeStatsDetailServiceHandler
}

func (h *codeStatsDetailServiceHandler) Get(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error {
	return h.CodeStatsDetailServiceHandler.Get(ctx, in, out)
}

func (h *codeStatsDetailServiceHandler) Search(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error {
	return h.CodeStatsDetailServiceHandler.Search(ctx, in, out)
}

func (h *codeStatsDetailServiceHandler) List(ctx context.Context, in *CodeStatsDetailWhere, out *CodeStatsDetailResponse) error {
	return h.CodeStatsDetailServiceHandler.List(ctx, in, out)
}
