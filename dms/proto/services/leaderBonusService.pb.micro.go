// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: leaderBonusService.proto

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

// Api Endpoints for LeaderBonusService service

func NewLeaderBonusServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LeaderBonusService service

type LeaderBonusService interface {
	// 查看结算详情
	Get(ctx context.Context, in *LeaderBonus, opts ...client.CallOption) (*LeaderBonusResponse, error)
	// 分页查询结算列表
	Search(ctx context.Context, in *LeaderBonusWhere, opts ...client.CallOption) (*LeaderBonusResponse, error)
	// 结算分红
	Settlement(ctx context.Context, in *LeaderBonusWhere, opts ...client.CallOption) (*LeaderBonusResponse, error)
}

type leaderBonusService struct {
	c    client.Client
	name string
}

func NewLeaderBonusService(name string, c client.Client) LeaderBonusService {
	return &leaderBonusService{
		c:    c,
		name: name,
	}
}

func (c *leaderBonusService) Get(ctx context.Context, in *LeaderBonus, opts ...client.CallOption) (*LeaderBonusResponse, error) {
	req := c.c.NewRequest(c.name, "LeaderBonusService.Get", in)
	out := new(LeaderBonusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderBonusService) Search(ctx context.Context, in *LeaderBonusWhere, opts ...client.CallOption) (*LeaderBonusResponse, error) {
	req := c.c.NewRequest(c.name, "LeaderBonusService.Search", in)
	out := new(LeaderBonusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderBonusService) Settlement(ctx context.Context, in *LeaderBonusWhere, opts ...client.CallOption) (*LeaderBonusResponse, error) {
	req := c.c.NewRequest(c.name, "LeaderBonusService.Settlement", in)
	out := new(LeaderBonusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LeaderBonusService service

type LeaderBonusServiceHandler interface {
	// 查看结算详情
	Get(context.Context, *LeaderBonus, *LeaderBonusResponse) error
	// 分页查询结算列表
	Search(context.Context, *LeaderBonusWhere, *LeaderBonusResponse) error
	// 结算分红
	Settlement(context.Context, *LeaderBonusWhere, *LeaderBonusResponse) error
}

func RegisterLeaderBonusServiceHandler(s server.Server, hdlr LeaderBonusServiceHandler, opts ...server.HandlerOption) error {
	type leaderBonusService interface {
		Get(ctx context.Context, in *LeaderBonus, out *LeaderBonusResponse) error
		Search(ctx context.Context, in *LeaderBonusWhere, out *LeaderBonusResponse) error
		Settlement(ctx context.Context, in *LeaderBonusWhere, out *LeaderBonusResponse) error
	}
	type LeaderBonusService struct {
		leaderBonusService
	}
	h := &leaderBonusServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LeaderBonusService{h}, opts...))
}

type leaderBonusServiceHandler struct {
	LeaderBonusServiceHandler
}

func (h *leaderBonusServiceHandler) Get(ctx context.Context, in *LeaderBonus, out *LeaderBonusResponse) error {
	return h.LeaderBonusServiceHandler.Get(ctx, in, out)
}

func (h *leaderBonusServiceHandler) Search(ctx context.Context, in *LeaderBonusWhere, out *LeaderBonusResponse) error {
	return h.LeaderBonusServiceHandler.Search(ctx, in, out)
}

func (h *leaderBonusServiceHandler) Settlement(ctx context.Context, in *LeaderBonusWhere, out *LeaderBonusResponse) error {
	return h.LeaderBonusServiceHandler.Settlement(ctx, in, out)
}
