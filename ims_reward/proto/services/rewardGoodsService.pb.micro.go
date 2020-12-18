// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rewardGoodsService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for RewardGoodsService service

func NewRewardGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RewardGoodsService service

type RewardGoodsService interface {
	Create(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error)
	Delete(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error)
	Get(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardGoodsResponse, error)
}

type rewardGoodsService struct {
	c    client.Client
	name string
}

func NewRewardGoodsService(name string, c client.Client) RewardGoodsService {
	return &rewardGoodsService{
		c:    c,
		name: name,
	}
}

func (c *rewardGoodsService) Create(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardGoodsService.Create", in)
	out := new(RewardGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardGoodsService) Delete(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardGoodsService.Delete", in)
	out := new(RewardGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardGoodsService) Get(ctx context.Context, in *RewardGoods, opts ...client.CallOption) (*RewardGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardGoodsService.Get", in)
	out := new(RewardGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardGoodsService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "RewardGoodsService.Search", in)
	out := new(RewardGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RewardGoodsService service

type RewardGoodsServiceHandler interface {
	Create(context.Context, *RewardGoods, *RewardGoodsResponse) error
	Delete(context.Context, *RewardGoods, *RewardGoodsResponse) error
	Get(context.Context, *RewardGoods, *RewardGoodsResponse) error
	Search(context.Context, *common.BaseWhere, *RewardGoodsResponse) error
}

func RegisterRewardGoodsServiceHandler(s server.Server, hdlr RewardGoodsServiceHandler, opts ...server.HandlerOption) error {
	type rewardGoodsService interface {
		Create(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error
		Delete(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error
		Get(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *RewardGoodsResponse) error
	}
	type RewardGoodsService struct {
		rewardGoodsService
	}
	h := &rewardGoodsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RewardGoodsService{h}, opts...))
}

type rewardGoodsServiceHandler struct {
	RewardGoodsServiceHandler
}

func (h *rewardGoodsServiceHandler) Create(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error {
	return h.RewardGoodsServiceHandler.Create(ctx, in, out)
}

func (h *rewardGoodsServiceHandler) Delete(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error {
	return h.RewardGoodsServiceHandler.Delete(ctx, in, out)
}

func (h *rewardGoodsServiceHandler) Get(ctx context.Context, in *RewardGoods, out *RewardGoodsResponse) error {
	return h.RewardGoodsServiceHandler.Get(ctx, in, out)
}

func (h *rewardGoodsServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *RewardGoodsResponse) error {
	return h.RewardGoodsServiceHandler.Search(ctx, in, out)
}
