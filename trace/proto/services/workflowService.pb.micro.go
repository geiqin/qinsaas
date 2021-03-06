// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: workflowService.proto

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

// Api Endpoints for WorkflowService service

func NewWorkflowServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WorkflowService service

type WorkflowService interface {
	Create(ctx context.Context, in *Workflow, opts ...client.CallOption) (*WorkflowResponse, error)
	Update(ctx context.Context, in *Workflow, opts ...client.CallOption) (*WorkflowResponse, error)
	Search(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error)
	Delete(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error)
	Get(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error)
	List(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error)
}

type workflowService struct {
	c    client.Client
	name string
}

func NewWorkflowService(name string, c client.Client) WorkflowService {
	return &workflowService{
		c:    c,
		name: name,
	}
}

func (c *workflowService) Create(ctx context.Context, in *Workflow, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.Create", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowService) Update(ctx context.Context, in *Workflow, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.Update", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowService) Search(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.Search", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowService) Delete(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.Delete", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowService) Get(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.Get", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowService) List(ctx context.Context, in *WorkflowWhere, opts ...client.CallOption) (*WorkflowResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowService.List", in)
	out := new(WorkflowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowService service

type WorkflowServiceHandler interface {
	Create(context.Context, *Workflow, *WorkflowResponse) error
	Update(context.Context, *Workflow, *WorkflowResponse) error
	Search(context.Context, *WorkflowWhere, *WorkflowResponse) error
	Delete(context.Context, *WorkflowWhere, *WorkflowResponse) error
	Get(context.Context, *WorkflowWhere, *WorkflowResponse) error
	List(context.Context, *WorkflowWhere, *WorkflowResponse) error
}

func RegisterWorkflowServiceHandler(s server.Server, hdlr WorkflowServiceHandler, opts ...server.HandlerOption) error {
	type workflowService interface {
		Create(ctx context.Context, in *Workflow, out *WorkflowResponse) error
		Update(ctx context.Context, in *Workflow, out *WorkflowResponse) error
		Search(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error
		Delete(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error
		Get(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error
		List(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error
	}
	type WorkflowService struct {
		workflowService
	}
	h := &workflowServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WorkflowService{h}, opts...))
}

type workflowServiceHandler struct {
	WorkflowServiceHandler
}

func (h *workflowServiceHandler) Create(ctx context.Context, in *Workflow, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.Create(ctx, in, out)
}

func (h *workflowServiceHandler) Update(ctx context.Context, in *Workflow, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.Update(ctx, in, out)
}

func (h *workflowServiceHandler) Search(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.Search(ctx, in, out)
}

func (h *workflowServiceHandler) Delete(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.Delete(ctx, in, out)
}

func (h *workflowServiceHandler) Get(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.Get(ctx, in, out)
}

func (h *workflowServiceHandler) List(ctx context.Context, in *WorkflowWhere, out *WorkflowResponse) error {
	return h.WorkflowServiceHandler.List(ctx, in, out)
}
