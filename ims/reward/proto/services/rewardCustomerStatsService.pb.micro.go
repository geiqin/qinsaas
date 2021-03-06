// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rewardCustomerStatsService.proto

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

// Api Endpoints for RewardCustomerStatsService service

func NewRewardCustomerStatsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RewardCustomerStatsService service

type RewardCustomerStatsService interface {
	// 查询满减送活动用户统计
	Search(ctx context.Context, in *RewardCustomerStatsWhere, opts ...client.CallOption) (*RewardCustomerStatsResponse, error)
	// 获取满减送活动用户统计
	Get(ctx context.Context, in *RewardCustomerStatsWhere, opts ...client.CallOption) (*RewardCustomerStatsResponse, error)
}

type rewardCustomerStatsService struct {
	c    client.Client
	name string
}

func NewRewardCustomerStatsService(name string, c client.Client) RewardCustomerStatsService {
	return &rewardCustomerStatsService{
		c:    c,
		name: name,
	}
}

func (c *rewardCustomerStatsService) Search(ctx context.Context, in *RewardCustomerStatsWhere, opts ...client.CallOption) (*RewardCustomerStatsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardCustomerStatsService.Search", in)
	out := new(RewardCustomerStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardCustomerStatsService) Get(ctx context.Context, in *RewardCustomerStatsWhere, opts ...client.CallOption) (*RewardCustomerStatsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardCustomerStatsService.Get", in)
	out := new(RewardCustomerStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RewardCustomerStatsService service

type RewardCustomerStatsServiceHandler interface {
	// 查询满减送活动用户统计
	Search(context.Context, *RewardCustomerStatsWhere, *RewardCustomerStatsResponse) error
	// 获取满减送活动用户统计
	Get(context.Context, *RewardCustomerStatsWhere, *RewardCustomerStatsResponse) error
}

func RegisterRewardCustomerStatsServiceHandler(s server.Server, hdlr RewardCustomerStatsServiceHandler, opts ...server.HandlerOption) error {
	type rewardCustomerStatsService interface {
		Search(ctx context.Context, in *RewardCustomerStatsWhere, out *RewardCustomerStatsResponse) error
		Get(ctx context.Context, in *RewardCustomerStatsWhere, out *RewardCustomerStatsResponse) error
	}
	type RewardCustomerStatsService struct {
		rewardCustomerStatsService
	}
	h := &rewardCustomerStatsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RewardCustomerStatsService{h}, opts...))
}

type rewardCustomerStatsServiceHandler struct {
	RewardCustomerStatsServiceHandler
}

func (h *rewardCustomerStatsServiceHandler) Search(ctx context.Context, in *RewardCustomerStatsWhere, out *RewardCustomerStatsResponse) error {
	return h.RewardCustomerStatsServiceHandler.Search(ctx, in, out)
}

func (h *rewardCustomerStatsServiceHandler) Get(ctx context.Context, in *RewardCustomerStatsWhere, out *RewardCustomerStatsResponse) error {
	return h.RewardCustomerStatsServiceHandler.Get(ctx, in, out)
}
