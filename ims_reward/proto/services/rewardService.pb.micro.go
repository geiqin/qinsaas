// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rewardService.proto

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

// Api Endpoints for RewardService service

func NewRewardServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RewardService service

type RewardService interface {
	//创建满减送活动
	Create(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error)
	// 编辑满减送活动
	Update(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error)
	// 删除满减活动
	Delete(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error)
	//终止满减送活动
	Cancel(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error)
	// 获取活动详情
	Get(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error)
	//分页查询活动列表
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardResponse, error)
	//获取活动列表
	List(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardResponse, error)
}

type rewardService struct {
	c    client.Client
	name string
}

func NewRewardService(name string, c client.Client) RewardService {
	return &rewardService{
		c:    c,
		name: name,
	}
}

func (c *rewardService) Create(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Create", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) Update(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Update", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) Delete(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Delete", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) Cancel(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Cancel", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) Get(ctx context.Context, in *Reward, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Get", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.Search", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardService) List(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RewardResponse, error) {
	req := c.c.NewRequest(c.name, "RewardService.List", in)
	out := new(RewardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RewardService service

type RewardServiceHandler interface {
	//创建满减送活动
	Create(context.Context, *Reward, *RewardResponse) error
	// 编辑满减送活动
	Update(context.Context, *Reward, *RewardResponse) error
	// 删除满减活动
	Delete(context.Context, *Reward, *RewardResponse) error
	//终止满减送活动
	Cancel(context.Context, *Reward, *RewardResponse) error
	// 获取活动详情
	Get(context.Context, *Reward, *RewardResponse) error
	//分页查询活动列表
	Search(context.Context, *common.BaseWhere, *RewardResponse) error
	//获取活动列表
	List(context.Context, *common.BaseWhere, *RewardResponse) error
}

func RegisterRewardServiceHandler(s server.Server, hdlr RewardServiceHandler, opts ...server.HandlerOption) error {
	type rewardService interface {
		Create(ctx context.Context, in *Reward, out *RewardResponse) error
		Update(ctx context.Context, in *Reward, out *RewardResponse) error
		Delete(ctx context.Context, in *Reward, out *RewardResponse) error
		Cancel(ctx context.Context, in *Reward, out *RewardResponse) error
		Get(ctx context.Context, in *Reward, out *RewardResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *RewardResponse) error
		List(ctx context.Context, in *common.BaseWhere, out *RewardResponse) error
	}
	type RewardService struct {
		rewardService
	}
	h := &rewardServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RewardService{h}, opts...))
}

type rewardServiceHandler struct {
	RewardServiceHandler
}

func (h *rewardServiceHandler) Create(ctx context.Context, in *Reward, out *RewardResponse) error {
	return h.RewardServiceHandler.Create(ctx, in, out)
}

func (h *rewardServiceHandler) Update(ctx context.Context, in *Reward, out *RewardResponse) error {
	return h.RewardServiceHandler.Update(ctx, in, out)
}

func (h *rewardServiceHandler) Delete(ctx context.Context, in *Reward, out *RewardResponse) error {
	return h.RewardServiceHandler.Delete(ctx, in, out)
}

func (h *rewardServiceHandler) Cancel(ctx context.Context, in *Reward, out *RewardResponse) error {
	return h.RewardServiceHandler.Cancel(ctx, in, out)
}

func (h *rewardServiceHandler) Get(ctx context.Context, in *Reward, out *RewardResponse) error {
	return h.RewardServiceHandler.Get(ctx, in, out)
}

func (h *rewardServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *RewardResponse) error {
	return h.RewardServiceHandler.Search(ctx, in, out)
}

func (h *rewardServiceHandler) List(ctx context.Context, in *common.BaseWhere, out *RewardResponse) error {
	return h.RewardServiceHandler.List(ctx, in, out)
}
