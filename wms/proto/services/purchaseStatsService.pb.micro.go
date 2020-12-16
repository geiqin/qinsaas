// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: purchaseStatsService.proto

package services

import (
	fmt "fmt"
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

// Api Endpoints for PurchaseInStatsDayService service

func NewPurchaseInStatsDayServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PurchaseInStatsDayService service

type PurchaseInStatsDayService interface {
	Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error)
	List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error)
}

type purchaseInStatsDayService struct {
	c    client.Client
	name string
}

func NewPurchaseInStatsDayService(name string, c client.Client) PurchaseInStatsDayService {
	return &purchaseInStatsDayService{
		c:    c,
		name: name,
	}
}

func (c *purchaseInStatsDayService) Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseInStatsDayService.Search", in)
	out := new(StatsDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseInStatsDayService) List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseInStatsDayService.List", in)
	out := new(StatsDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PurchaseInStatsDayService service

type PurchaseInStatsDayServiceHandler interface {
	Search(context.Context, *StatsWhere, *StatsDayResponse) error
	List(context.Context, *StatsWhere, *StatsDayResponse) error
}

func RegisterPurchaseInStatsDayServiceHandler(s server.Server, hdlr PurchaseInStatsDayServiceHandler, opts ...server.HandlerOption) error {
	type purchaseInStatsDayService interface {
		Search(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error
		List(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error
	}
	type PurchaseInStatsDayService struct {
		purchaseInStatsDayService
	}
	h := &purchaseInStatsDayServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PurchaseInStatsDayService{h}, opts...))
}

type purchaseInStatsDayServiceHandler struct {
	PurchaseInStatsDayServiceHandler
}

func (h *purchaseInStatsDayServiceHandler) Search(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error {
	return h.PurchaseInStatsDayServiceHandler.Search(ctx, in, out)
}

func (h *purchaseInStatsDayServiceHandler) List(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error {
	return h.PurchaseInStatsDayServiceHandler.List(ctx, in, out)
}

// Api Endpoints for PurchaseInStatsMonthService service

func NewPurchaseInStatsMonthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PurchaseInStatsMonthService service

type PurchaseInStatsMonthService interface {
	Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error)
	List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error)
}

type purchaseInStatsMonthService struct {
	c    client.Client
	name string
}

func NewPurchaseInStatsMonthService(name string, c client.Client) PurchaseInStatsMonthService {
	return &purchaseInStatsMonthService{
		c:    c,
		name: name,
	}
}

func (c *purchaseInStatsMonthService) Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseInStatsMonthService.Search", in)
	out := new(StatsMonthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseInStatsMonthService) List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseInStatsMonthService.List", in)
	out := new(StatsMonthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PurchaseInStatsMonthService service

type PurchaseInStatsMonthServiceHandler interface {
	Search(context.Context, *StatsWhere, *StatsMonthResponse) error
	List(context.Context, *StatsWhere, *StatsMonthResponse) error
}

func RegisterPurchaseInStatsMonthServiceHandler(s server.Server, hdlr PurchaseInStatsMonthServiceHandler, opts ...server.HandlerOption) error {
	type purchaseInStatsMonthService interface {
		Search(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error
		List(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error
	}
	type PurchaseInStatsMonthService struct {
		purchaseInStatsMonthService
	}
	h := &purchaseInStatsMonthServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PurchaseInStatsMonthService{h}, opts...))
}

type purchaseInStatsMonthServiceHandler struct {
	PurchaseInStatsMonthServiceHandler
}

func (h *purchaseInStatsMonthServiceHandler) Search(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error {
	return h.PurchaseInStatsMonthServiceHandler.Search(ctx, in, out)
}

func (h *purchaseInStatsMonthServiceHandler) List(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error {
	return h.PurchaseInStatsMonthServiceHandler.List(ctx, in, out)
}

// Api Endpoints for PurchaseOutStatsDayService service

func NewPurchaseOutStatsDayServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PurchaseOutStatsDayService service

type PurchaseOutStatsDayService interface {
	Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error)
	List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error)
}

type purchaseOutStatsDayService struct {
	c    client.Client
	name string
}

func NewPurchaseOutStatsDayService(name string, c client.Client) PurchaseOutStatsDayService {
	return &purchaseOutStatsDayService{
		c:    c,
		name: name,
	}
}

func (c *purchaseOutStatsDayService) Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseOutStatsDayService.Search", in)
	out := new(StatsDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseOutStatsDayService) List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsDayResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseOutStatsDayService.List", in)
	out := new(StatsDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PurchaseOutStatsDayService service

type PurchaseOutStatsDayServiceHandler interface {
	Search(context.Context, *StatsWhere, *StatsDayResponse) error
	List(context.Context, *StatsWhere, *StatsDayResponse) error
}

func RegisterPurchaseOutStatsDayServiceHandler(s server.Server, hdlr PurchaseOutStatsDayServiceHandler, opts ...server.HandlerOption) error {
	type purchaseOutStatsDayService interface {
		Search(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error
		List(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error
	}
	type PurchaseOutStatsDayService struct {
		purchaseOutStatsDayService
	}
	h := &purchaseOutStatsDayServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PurchaseOutStatsDayService{h}, opts...))
}

type purchaseOutStatsDayServiceHandler struct {
	PurchaseOutStatsDayServiceHandler
}

func (h *purchaseOutStatsDayServiceHandler) Search(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error {
	return h.PurchaseOutStatsDayServiceHandler.Search(ctx, in, out)
}

func (h *purchaseOutStatsDayServiceHandler) List(ctx context.Context, in *StatsWhere, out *StatsDayResponse) error {
	return h.PurchaseOutStatsDayServiceHandler.List(ctx, in, out)
}

// Api Endpoints for PurchaseOutStatsMonthService service

func NewPurchaseOutStatsMonthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PurchaseOutStatsMonthService service

type PurchaseOutStatsMonthService interface {
	Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error)
	List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error)
}

type purchaseOutStatsMonthService struct {
	c    client.Client
	name string
}

func NewPurchaseOutStatsMonthService(name string, c client.Client) PurchaseOutStatsMonthService {
	return &purchaseOutStatsMonthService{
		c:    c,
		name: name,
	}
}

func (c *purchaseOutStatsMonthService) Search(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseOutStatsMonthService.Search", in)
	out := new(StatsMonthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseOutStatsMonthService) List(ctx context.Context, in *StatsWhere, opts ...client.CallOption) (*StatsMonthResponse, error) {
	req := c.c.NewRequest(c.name, "PurchaseOutStatsMonthService.List", in)
	out := new(StatsMonthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PurchaseOutStatsMonthService service

type PurchaseOutStatsMonthServiceHandler interface {
	Search(context.Context, *StatsWhere, *StatsMonthResponse) error
	List(context.Context, *StatsWhere, *StatsMonthResponse) error
}

func RegisterPurchaseOutStatsMonthServiceHandler(s server.Server, hdlr PurchaseOutStatsMonthServiceHandler, opts ...server.HandlerOption) error {
	type purchaseOutStatsMonthService interface {
		Search(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error
		List(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error
	}
	type PurchaseOutStatsMonthService struct {
		purchaseOutStatsMonthService
	}
	h := &purchaseOutStatsMonthServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PurchaseOutStatsMonthService{h}, opts...))
}

type purchaseOutStatsMonthServiceHandler struct {
	PurchaseOutStatsMonthServiceHandler
}

func (h *purchaseOutStatsMonthServiceHandler) Search(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error {
	return h.PurchaseOutStatsMonthServiceHandler.Search(ctx, in, out)
}

func (h *purchaseOutStatsMonthServiceHandler) List(ctx context.Context, in *StatsWhere, out *StatsMonthResponse) error {
	return h.PurchaseOutStatsMonthServiceHandler.List(ctx, in, out)
}