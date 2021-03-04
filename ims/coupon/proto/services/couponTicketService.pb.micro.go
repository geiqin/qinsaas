// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: couponTicketService.proto

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

// Api Endpoints for CouponTicketService service

func NewCouponTicketServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CouponTicketService service

type CouponTicketService interface {
	//核销优惠劵凭证
	Use(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error)
	//获取优惠劵凭证信息
	Get(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
	//查询优惠劵凭证
	Search(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
	//获取优惠劵凭证列表
	List(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
}

type couponTicketService struct {
	c    client.Client
	name string
}

func NewCouponTicketService(name string, c client.Client) CouponTicketService {
	return &couponTicketService{
		c:    c,
		name: name,
	}
}

func (c *couponTicketService) Use(ctx context.Context, in *CouponTicket, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Use", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) Get(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Get", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) Search(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.Search", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponTicketService) List(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "CouponTicketService.List", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CouponTicketService service

type CouponTicketServiceHandler interface {
	//核销优惠劵凭证
	Use(context.Context, *CouponTicket, *CouponTicketResponse) error
	//获取优惠劵凭证信息
	Get(context.Context, *CouponTicketWhere, *CouponTicketResponse) error
	//查询优惠劵凭证
	Search(context.Context, *CouponTicketWhere, *CouponTicketResponse) error
	//获取优惠劵凭证列表
	List(context.Context, *CouponTicketWhere, *CouponTicketResponse) error
}

func RegisterCouponTicketServiceHandler(s server.Server, hdlr CouponTicketServiceHandler, opts ...server.HandlerOption) error {
	type couponTicketService interface {
		Use(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error
		Get(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error
		Search(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error
		List(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error
	}
	type CouponTicketService struct {
		couponTicketService
	}
	h := &couponTicketServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CouponTicketService{h}, opts...))
}

type couponTicketServiceHandler struct {
	CouponTicketServiceHandler
}

func (h *couponTicketServiceHandler) Use(ctx context.Context, in *CouponTicket, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Use(ctx, in, out)
}

func (h *couponTicketServiceHandler) Get(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Get(ctx, in, out)
}

func (h *couponTicketServiceHandler) Search(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.Search(ctx, in, out)
}

func (h *couponTicketServiceHandler) List(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error {
	return h.CouponTicketServiceHandler.List(ctx, in, out)
}

// Api Endpoints for MyCouponTicketService service

func NewMyCouponTicketServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyCouponTicketService service

type MyCouponTicketService interface {
	//查询优惠劵凭证
	Search(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
	//获取可使用的优惠劵凭证列表（如：下单选择优惠券）
	UsableList(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error)
}

type myCouponTicketService struct {
	c    client.Client
	name string
}

func NewMyCouponTicketService(name string, c client.Client) MyCouponTicketService {
	return &myCouponTicketService{
		c:    c,
		name: name,
	}
}

func (c *myCouponTicketService) Search(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "MyCouponTicketService.Search", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myCouponTicketService) UsableList(ctx context.Context, in *CouponTicketWhere, opts ...client.CallOption) (*CouponTicketResponse, error) {
	req := c.c.NewRequest(c.name, "MyCouponTicketService.UsableList", in)
	out := new(CouponTicketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyCouponTicketService service

type MyCouponTicketServiceHandler interface {
	//查询优惠劵凭证
	Search(context.Context, *CouponTicketWhere, *CouponTicketResponse) error
	//获取可使用的优惠劵凭证列表（如：下单选择优惠券）
	UsableList(context.Context, *CouponTicketWhere, *CouponTicketResponse) error
}

func RegisterMyCouponTicketServiceHandler(s server.Server, hdlr MyCouponTicketServiceHandler, opts ...server.HandlerOption) error {
	type myCouponTicketService interface {
		Search(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error
		UsableList(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error
	}
	type MyCouponTicketService struct {
		myCouponTicketService
	}
	h := &myCouponTicketServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyCouponTicketService{h}, opts...))
}

type myCouponTicketServiceHandler struct {
	MyCouponTicketServiceHandler
}

func (h *myCouponTicketServiceHandler) Search(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error {
	return h.MyCouponTicketServiceHandler.Search(ctx, in, out)
}

func (h *myCouponTicketServiceHandler) UsableList(ctx context.Context, in *CouponTicketWhere, out *CouponTicketResponse) error {
	return h.MyCouponTicketServiceHandler.UsableList(ctx, in, out)
}