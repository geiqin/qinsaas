// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: reviewService.proto

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

// Api Endpoints for ReviewService service

func NewReviewServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ReviewService service

type ReviewService interface {
	Create(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error)
	Delete(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error)
	Get(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ReviewResponse, error)
}

type reviewService struct {
	c    client.Client
	name string
}

func NewReviewService(name string, c client.Client) ReviewService {
	return &reviewService{
		c:    c,
		name: name,
	}
}

func (c *reviewService) Create(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.Create", in)
	out := new(ReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) Delete(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.Delete", in)
	out := new(ReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) Get(ctx context.Context, in *Review, opts ...client.CallOption) (*ReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.Get", in)
	out := new(ReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*ReviewResponse, error) {
	req := c.c.NewRequest(c.name, "ReviewService.Search", in)
	out := new(ReviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReviewService service

type ReviewServiceHandler interface {
	Create(context.Context, *Review, *ReviewResponse) error
	Delete(context.Context, *Review, *ReviewResponse) error
	Get(context.Context, *Review, *ReviewResponse) error
	Search(context.Context, *common.BaseWhere, *ReviewResponse) error
}

func RegisterReviewServiceHandler(s server.Server, hdlr ReviewServiceHandler, opts ...server.HandlerOption) error {
	type reviewService interface {
		Create(ctx context.Context, in *Review, out *ReviewResponse) error
		Delete(ctx context.Context, in *Review, out *ReviewResponse) error
		Get(ctx context.Context, in *Review, out *ReviewResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *ReviewResponse) error
	}
	type ReviewService struct {
		reviewService
	}
	h := &reviewServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ReviewService{h}, opts...))
}

type reviewServiceHandler struct {
	ReviewServiceHandler
}

func (h *reviewServiceHandler) Create(ctx context.Context, in *Review, out *ReviewResponse) error {
	return h.ReviewServiceHandler.Create(ctx, in, out)
}

func (h *reviewServiceHandler) Delete(ctx context.Context, in *Review, out *ReviewResponse) error {
	return h.ReviewServiceHandler.Delete(ctx, in, out)
}

func (h *reviewServiceHandler) Get(ctx context.Context, in *Review, out *ReviewResponse) error {
	return h.ReviewServiceHandler.Get(ctx, in, out)
}

func (h *reviewServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *ReviewResponse) error {
	return h.ReviewServiceHandler.Search(ctx, in, out)
}
