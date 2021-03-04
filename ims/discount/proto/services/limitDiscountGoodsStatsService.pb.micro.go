// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: limitDiscountGoodsStatsService.proto

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

// Api Endpoints for LimitDiscountGoodsStatsService service

func NewLimitDiscountGoodsStatsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LimitDiscountGoodsStatsService service

type LimitDiscountGoodsStatsService interface {
	// 查询限时折扣活动用户购买的商品统计
	Search(ctx context.Context, in *LimitDiscountGoodsStatsWhere, opts ...client.CallOption) (*LimitDiscountGoodsStatsResponse, error)
	// 获取限时折扣活动用户购买的商品统计
	Get(ctx context.Context, in *LimitDiscountGoodsStatsWhere, opts ...client.CallOption) (*LimitDiscountGoodsStatsResponse, error)
}

type limitDiscountGoodsStatsService struct {
	c    client.Client
	name string
}

func NewLimitDiscountGoodsStatsService(name string, c client.Client) LimitDiscountGoodsStatsService {
	return &limitDiscountGoodsStatsService{
		c:    c,
		name: name,
	}
}

func (c *limitDiscountGoodsStatsService) Search(ctx context.Context, in *LimitDiscountGoodsStatsWhere, opts ...client.CallOption) (*LimitDiscountGoodsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountGoodsStatsService.Search", in)
	out := new(LimitDiscountGoodsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitDiscountGoodsStatsService) Get(ctx context.Context, in *LimitDiscountGoodsStatsWhere, opts ...client.CallOption) (*LimitDiscountGoodsStatsResponse, error) {
	req := c.c.NewRequest(c.name, "LimitDiscountGoodsStatsService.Get", in)
	out := new(LimitDiscountGoodsStatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LimitDiscountGoodsStatsService service

type LimitDiscountGoodsStatsServiceHandler interface {
	// 查询限时折扣活动用户购买的商品统计
	Search(context.Context, *LimitDiscountGoodsStatsWhere, *LimitDiscountGoodsStatsResponse) error
	// 获取限时折扣活动用户购买的商品统计
	Get(context.Context, *LimitDiscountGoodsStatsWhere, *LimitDiscountGoodsStatsResponse) error
}

func RegisterLimitDiscountGoodsStatsServiceHandler(s server.Server, hdlr LimitDiscountGoodsStatsServiceHandler, opts ...server.HandlerOption) error {
	type limitDiscountGoodsStatsService interface {
		Search(ctx context.Context, in *LimitDiscountGoodsStatsWhere, out *LimitDiscountGoodsStatsResponse) error
		Get(ctx context.Context, in *LimitDiscountGoodsStatsWhere, out *LimitDiscountGoodsStatsResponse) error
	}
	type LimitDiscountGoodsStatsService struct {
		limitDiscountGoodsStatsService
	}
	h := &limitDiscountGoodsStatsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LimitDiscountGoodsStatsService{h}, opts...))
}

type limitDiscountGoodsStatsServiceHandler struct {
	LimitDiscountGoodsStatsServiceHandler
}

func (h *limitDiscountGoodsStatsServiceHandler) Search(ctx context.Context, in *LimitDiscountGoodsStatsWhere, out *LimitDiscountGoodsStatsResponse) error {
	return h.LimitDiscountGoodsStatsServiceHandler.Search(ctx, in, out)
}

func (h *limitDiscountGoodsStatsServiceHandler) Get(ctx context.Context, in *LimitDiscountGoodsStatsWhere, out *LimitDiscountGoodsStatsResponse) error {
	return h.LimitDiscountGoodsStatsServiceHandler.Get(ctx, in, out)
}