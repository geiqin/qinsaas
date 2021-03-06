// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rewardStatsService.proto

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

// Api Endpoints for RewardStatsService service

func NewRewardStatsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RewardStatsService service

type RewardStatsService interface {
	// 查询满减送活动汇总统计
	Search(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error)
	// 获取满减送活动汇总统计
	Get(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error)
	// 重置满减送活动统计
	Reset(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error)
}

type rewardStatsService struct {
	c    client.Client
	name string
}

func NewRewardStatsService(name string, c client.Client) RewardStatsService {
	return &rewardStatsService{
		c:    c,
		name: name,
	}
}

func (c *rewardStatsService) Search(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardStatsService.Search", in)
	out := new(RewardStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardStatsService) Get(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardStatsService.Get", in)
	out := new(RewardStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardStatsService) Reset(ctx context.Context, in *RewardStatsWhere, opts ...client.CallOption) (*RewardStatsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardStatsService.Reset", in)
	out := new(RewardStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RewardStatsService service

type RewardStatsServiceHandler interface {
	// 查询满减送活动汇总统计
	Search(context.Context, *RewardStatsWhere, *RewardStatsResponse) error
	// 获取满减送活动汇总统计
	Get(context.Context, *RewardStatsWhere, *RewardStatsResponse) error
	// 重置满减送活动统计
	Reset(context.Context, *RewardStatsWhere, *RewardStatsResponse) error
}

func RegisterRewardStatsServiceHandler(s server.Server, hdlr RewardStatsServiceHandler, opts ...server.HandlerOption) error {
	type rewardStatsService interface {
		Search(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error
		Get(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error
		Reset(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error
	}
	type RewardStatsService struct {
		rewardStatsService
	}
	h := &rewardStatsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RewardStatsService{h}, opts...))
}

type rewardStatsServiceHandler struct {
	RewardStatsServiceHandler
}

func (h *rewardStatsServiceHandler) Search(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error {
	return h.RewardStatsServiceHandler.Search(ctx, in, out)
}

func (h *rewardStatsServiceHandler) Get(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error {
	return h.RewardStatsServiceHandler.Get(ctx, in, out)
}

func (h *rewardStatsServiceHandler) Reset(ctx context.Context, in *RewardStatsWhere, out *RewardStatsResponse) error {
	return h.RewardStatsServiceHandler.Reset(ctx, in, out)
}
