// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: relationshipService.proto

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

// Api Endpoints for RelationshipService service

func NewRelationshipServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RelationshipService service

type RelationshipService interface {
	//更改关系
	Change(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error)
	//删除关联
	Delete(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error)
	//查看关系
	Get(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error)
	//查询关系记录
	Search(ctx context.Context, in *RelationshipWhere, opts ...client.CallOption) (*RelationshipResponse, error)
}

type relationshipService struct {
	c    client.Client
	name string
}

func NewRelationshipService(name string, c client.Client) RelationshipService {
	return &relationshipService{
		c:    c,
		name: name,
	}
}

func (c *relationshipService) Change(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error) {
	req := c.c.NewRequest(c.name, "RelationshipService.Change", in)
	out := new(RelationshipResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipService) Delete(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error) {
	req := c.c.NewRequest(c.name, "RelationshipService.Delete", in)
	out := new(RelationshipResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipService) Get(ctx context.Context, in *Relationship, opts ...client.CallOption) (*RelationshipResponse, error) {
	req := c.c.NewRequest(c.name, "RelationshipService.Get", in)
	out := new(RelationshipResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipService) Search(ctx context.Context, in *RelationshipWhere, opts ...client.CallOption) (*RelationshipResponse, error) {
	req := c.c.NewRequest(c.name, "RelationshipService.Search", in)
	out := new(RelationshipResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RelationshipService service

type RelationshipServiceHandler interface {
	//更改关系
	Change(context.Context, *Relationship, *RelationshipResponse) error
	//删除关联
	Delete(context.Context, *Relationship, *RelationshipResponse) error
	//查看关系
	Get(context.Context, *Relationship, *RelationshipResponse) error
	//查询关系记录
	Search(context.Context, *RelationshipWhere, *RelationshipResponse) error
}

func RegisterRelationshipServiceHandler(s server.Server, hdlr RelationshipServiceHandler, opts ...server.HandlerOption) error {
	type relationshipService interface {
		Change(ctx context.Context, in *Relationship, out *RelationshipResponse) error
		Delete(ctx context.Context, in *Relationship, out *RelationshipResponse) error
		Get(ctx context.Context, in *Relationship, out *RelationshipResponse) error
		Search(ctx context.Context, in *RelationshipWhere, out *RelationshipResponse) error
	}
	type RelationshipService struct {
		relationshipService
	}
	h := &relationshipServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RelationshipService{h}, opts...))
}

type relationshipServiceHandler struct {
	RelationshipServiceHandler
}

func (h *relationshipServiceHandler) Change(ctx context.Context, in *Relationship, out *RelationshipResponse) error {
	return h.RelationshipServiceHandler.Change(ctx, in, out)
}

func (h *relationshipServiceHandler) Delete(ctx context.Context, in *Relationship, out *RelationshipResponse) error {
	return h.RelationshipServiceHandler.Delete(ctx, in, out)
}

func (h *relationshipServiceHandler) Get(ctx context.Context, in *Relationship, out *RelationshipResponse) error {
	return h.RelationshipServiceHandler.Get(ctx, in, out)
}

func (h *relationshipServiceHandler) Search(ctx context.Context, in *RelationshipWhere, out *RelationshipResponse) error {
	return h.RelationshipServiceHandler.Search(ctx, in, out)
}
