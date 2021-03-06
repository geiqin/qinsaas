// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: deliveryService.proto

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

// Api Endpoints for DeliveryCfgService service

func NewDeliveryCfgServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeliveryCfgService service

type DeliveryCfgService interface {
	// 启用同城配送功能
	Open(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error)
	// 关闭同城配送功能
	Close(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error)
	// 检查同城配送功能是否开启
	IsOpen(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error)
}

type deliveryCfgService struct {
	c    client.Client
	name string
}

func NewDeliveryCfgService(name string, c client.Client) DeliveryCfgService {
	return &deliveryCfgService{
		c:    c,
		name: name,
	}
}

func (c *deliveryCfgService) Open(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryCfgService.Open", in)
	out := new(SettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryCfgService) Close(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryCfgService.Close", in)
	out := new(SettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryCfgService) IsOpen(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*SettingResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryCfgService.IsOpen", in)
	out := new(SettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeliveryCfgService service

type DeliveryCfgServiceHandler interface {
	// 启用同城配送功能
	Open(context.Context, *common.Empty, *SettingResponse) error
	// 关闭同城配送功能
	Close(context.Context, *common.Empty, *SettingResponse) error
	// 检查同城配送功能是否开启
	IsOpen(context.Context, *common.Empty, *SettingResponse) error
}

func RegisterDeliveryCfgServiceHandler(s server.Server, hdlr DeliveryCfgServiceHandler, opts ...server.HandlerOption) error {
	type deliveryCfgService interface {
		Open(ctx context.Context, in *common.Empty, out *SettingResponse) error
		Close(ctx context.Context, in *common.Empty, out *SettingResponse) error
		IsOpen(ctx context.Context, in *common.Empty, out *SettingResponse) error
	}
	type DeliveryCfgService struct {
		deliveryCfgService
	}
	h := &deliveryCfgServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeliveryCfgService{h}, opts...))
}

type deliveryCfgServiceHandler struct {
	DeliveryCfgServiceHandler
}

func (h *deliveryCfgServiceHandler) Open(ctx context.Context, in *common.Empty, out *SettingResponse) error {
	return h.DeliveryCfgServiceHandler.Open(ctx, in, out)
}

func (h *deliveryCfgServiceHandler) Close(ctx context.Context, in *common.Empty, out *SettingResponse) error {
	return h.DeliveryCfgServiceHandler.Close(ctx, in, out)
}

func (h *deliveryCfgServiceHandler) IsOpen(ctx context.Context, in *common.Empty, out *SettingResponse) error {
	return h.DeliveryCfgServiceHandler.IsOpen(ctx, in, out)
}

// Api Endpoints for DeliveryService service

func NewDeliveryServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeliveryService service

type DeliveryService interface {
	Get(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error)
	Update(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error)
	// 获取同城配送时间
	GetTimes(ctx context.Context, in *DeliveryWhere, opts ...client.CallOption) (*DateListResponse, error)
}

type deliveryService struct {
	c    client.Client
	name string
}

func NewDeliveryService(name string, c client.Client) DeliveryService {
	return &deliveryService{
		c:    c,
		name: name,
	}
}

func (c *deliveryService) Get(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryService.Get", in)
	out := new(DeliveryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryService) Update(ctx context.Context, in *Delivery, opts ...client.CallOption) (*DeliveryResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryService.Update", in)
	out := new(DeliveryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryService) GetTimes(ctx context.Context, in *DeliveryWhere, opts ...client.CallOption) (*DateListResponse, error) {
	req := c.c.NewRequest(c.name, "DeliveryService.GetTimes", in)
	out := new(DateListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeliveryService service

type DeliveryServiceHandler interface {
	Get(context.Context, *Delivery, *DeliveryResponse) error
	Update(context.Context, *Delivery, *DeliveryResponse) error
	// 获取同城配送时间
	GetTimes(context.Context, *DeliveryWhere, *DateListResponse) error
}

func RegisterDeliveryServiceHandler(s server.Server, hdlr DeliveryServiceHandler, opts ...server.HandlerOption) error {
	type deliveryService interface {
		Get(ctx context.Context, in *Delivery, out *DeliveryResponse) error
		Update(ctx context.Context, in *Delivery, out *DeliveryResponse) error
		GetTimes(ctx context.Context, in *DeliveryWhere, out *DateListResponse) error
	}
	type DeliveryService struct {
		deliveryService
	}
	h := &deliveryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeliveryService{h}, opts...))
}

type deliveryServiceHandler struct {
	DeliveryServiceHandler
}

func (h *deliveryServiceHandler) Get(ctx context.Context, in *Delivery, out *DeliveryResponse) error {
	return h.DeliveryServiceHandler.Get(ctx, in, out)
}

func (h *deliveryServiceHandler) Update(ctx context.Context, in *Delivery, out *DeliveryResponse) error {
	return h.DeliveryServiceHandler.Update(ctx, in, out)
}

func (h *deliveryServiceHandler) GetTimes(ctx context.Context, in *DeliveryWhere, out *DateListResponse) error {
	return h.DeliveryServiceHandler.GetTimes(ctx, in, out)
}
