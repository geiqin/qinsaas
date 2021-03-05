// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: groupGoodsStatsService.proto

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

// Api Endpoints for GroupGoodsStatsService service

func NewGroupGoodsStatsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GroupGoodsStatsService service

type GroupGoodsStatsService interface {
	// 查询活动统计
	Search(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error)
	// 获取活动统计
	Get(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error)
	// 获取活动统计列表
	List(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error)
}

type groupGoodsStatsService struct {
	c    client.Client
	name string
}

func NewGroupGoodsStatsService(name string, c client.Client) GroupGoodsStatsService {
	return &groupGoodsStatsService{
		c:    c,
		name: name,
	}
}

func (c *groupGoodsStatsService) Search(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "GroupGoodsStatsService.Search", in)
	out := new(GroupGoodsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupGoodsStatsService) Get(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "GroupGoodsStatsService.Get", in)
	out := new(GroupGoodsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupGoodsStatsService) List(ctx context.Context, in *GroupStatsWhere, opts ...client.CallOption) (*GroupGoodsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "GroupGoodsStatsService.List", in)
	out := new(GroupGoodsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupGoodsStatsService service

type GroupGoodsStatsServiceHandler interface {
	// 查询活动统计
	Search(context.Context, *GroupStatsWhere, *GroupGoodsStatsResponse) error
	// 获取活动统计
	Get(context.Context, *GroupStatsWhere, *GroupGoodsStatsResponse) error
	// 获取活动统计列表
	List(context.Context, *GroupStatsWhere, *GroupGoodsStatsResponse) error
}

func RegisterGroupGoodsStatsServiceHandler(s server.Server, hdlr GroupGoodsStatsServiceHandler, opts ...server.HandlerOption) error {
	type groupGoodsStatsService interface {
		Search(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error
		Get(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error
		List(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error
	}
	type GroupGoodsStatsService struct {
		groupGoodsStatsService
	}
	h := &groupGoodsStatsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GroupGoodsStatsService{h}, opts...))
}

type groupGoodsStatsServiceHandler struct {
	GroupGoodsStatsServiceHandler
}

func (h *groupGoodsStatsServiceHandler) Search(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error {
	return h.GroupGoodsStatsServiceHandler.Search(ctx, in, out)
}

func (h *groupGoodsStatsServiceHandler) Get(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error {
	return h.GroupGoodsStatsServiceHandler.Get(ctx, in, out)
}

func (h *groupGoodsStatsServiceHandler) List(ctx context.Context, in *GroupStatsWhere, out *GroupGoodsStatsResponse) error {
	return h.GroupGoodsStatsServiceHandler.List(ctx, in, out)
}
