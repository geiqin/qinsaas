// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: workflowRecordService.proto

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

// Api Endpoints for WorkflowRecordService service

func NewWorkflowRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WorkflowRecordService service

type WorkflowRecordService interface {
	Create(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Update(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Delete(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Check(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
}

type workflowRecordService struct {
	c    client.Client
	name string
}

func NewWorkflowRecordService(name string, c client.Client) WorkflowRecordService {
	return &workflowRecordService{
		c:    c,
		name: name,
	}
}

func (c *workflowRecordService) Create(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Create", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) Update(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Update", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Search", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) Delete(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Delete", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Get", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) Check(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.Check", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowRecordService) List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "WorkflowRecordService.List", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowRecordService service

type WorkflowRecordServiceHandler interface {
	Create(context.Context, *WorkflowRecord, *WorkflowRecordResponse) error
	Update(context.Context, *WorkflowRecord, *WorkflowRecordResponse) error
	Search(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	Delete(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	Get(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	Check(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	List(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
}

func RegisterWorkflowRecordServiceHandler(s server.Server, hdlr WorkflowRecordServiceHandler, opts ...server.HandlerOption) error {
	type workflowRecordService interface {
		Create(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error
		Update(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error
		Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		Delete(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		Check(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
	}
	type WorkflowRecordService struct {
		workflowRecordService
	}
	h := &workflowRecordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WorkflowRecordService{h}, opts...))
}

type workflowRecordServiceHandler struct {
	WorkflowRecordServiceHandler
}

func (h *workflowRecordServiceHandler) Create(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Create(ctx, in, out)
}

func (h *workflowRecordServiceHandler) Update(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Update(ctx, in, out)
}

func (h *workflowRecordServiceHandler) Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Search(ctx, in, out)
}

func (h *workflowRecordServiceHandler) Delete(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Delete(ctx, in, out)
}

func (h *workflowRecordServiceHandler) Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Get(ctx, in, out)
}

func (h *workflowRecordServiceHandler) Check(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.Check(ctx, in, out)
}

func (h *workflowRecordServiceHandler) List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.WorkflowRecordServiceHandler.List(ctx, in, out)
}

// Api Endpoints for MyWorkflowRecordService service

func NewMyWorkflowRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyWorkflowRecordService service

type MyWorkflowRecordService interface {
	Create(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Update(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
}

type myWorkflowRecordService struct {
	c    client.Client
	name string
}

func NewMyWorkflowRecordService(name string, c client.Client) MyWorkflowRecordService {
	return &myWorkflowRecordService{
		c:    c,
		name: name,
	}
}

func (c *myWorkflowRecordService) Create(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "MyWorkflowRecordService.Create", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myWorkflowRecordService) Update(ctx context.Context, in *WorkflowRecord, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "MyWorkflowRecordService.Update", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myWorkflowRecordService) Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "MyWorkflowRecordService.Get", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myWorkflowRecordService) Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "MyWorkflowRecordService.Search", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myWorkflowRecordService) List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "MyWorkflowRecordService.List", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyWorkflowRecordService service

type MyWorkflowRecordServiceHandler interface {
	Create(context.Context, *WorkflowRecord, *WorkflowRecordResponse) error
	Update(context.Context, *WorkflowRecord, *WorkflowRecordResponse) error
	Get(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	Search(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	List(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
}

func RegisterMyWorkflowRecordServiceHandler(s server.Server, hdlr MyWorkflowRecordServiceHandler, opts ...server.HandlerOption) error {
	type myWorkflowRecordService interface {
		Create(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error
		Update(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error
		Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
	}
	type MyWorkflowRecordService struct {
		myWorkflowRecordService
	}
	h := &myWorkflowRecordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyWorkflowRecordService{h}, opts...))
}

type myWorkflowRecordServiceHandler struct {
	MyWorkflowRecordServiceHandler
}

func (h *myWorkflowRecordServiceHandler) Create(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error {
	return h.MyWorkflowRecordServiceHandler.Create(ctx, in, out)
}

func (h *myWorkflowRecordServiceHandler) Update(ctx context.Context, in *WorkflowRecord, out *WorkflowRecordResponse) error {
	return h.MyWorkflowRecordServiceHandler.Update(ctx, in, out)
}

func (h *myWorkflowRecordServiceHandler) Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.MyWorkflowRecordServiceHandler.Get(ctx, in, out)
}

func (h *myWorkflowRecordServiceHandler) Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.MyWorkflowRecordServiceHandler.Search(ctx, in, out)
}

func (h *myWorkflowRecordServiceHandler) List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.MyWorkflowRecordServiceHandler.List(ctx, in, out)
}

// Api Endpoints for FrontWorkflowRecordService service

func NewFrontWorkflowRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontWorkflowRecordService service

type FrontWorkflowRecordService interface {
	Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
	List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error)
}

type frontWorkflowRecordService struct {
	c    client.Client
	name string
}

func NewFrontWorkflowRecordService(name string, c client.Client) FrontWorkflowRecordService {
	return &frontWorkflowRecordService{
		c:    c,
		name: name,
	}
}

func (c *frontWorkflowRecordService) Get(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "FrontWorkflowRecordService.Get", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontWorkflowRecordService) Search(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "FrontWorkflowRecordService.Search", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontWorkflowRecordService) List(ctx context.Context, in *WorkflowRecordWhere, opts ...client.CallOption) (*WorkflowRecordResponse, error) {
	req := c.c.NewRequest(c.name, "FrontWorkflowRecordService.List", in)
	out := new(WorkflowRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontWorkflowRecordService service

type FrontWorkflowRecordServiceHandler interface {
	Get(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	Search(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
	List(context.Context, *WorkflowRecordWhere, *WorkflowRecordResponse) error
}

func RegisterFrontWorkflowRecordServiceHandler(s server.Server, hdlr FrontWorkflowRecordServiceHandler, opts ...server.HandlerOption) error {
	type frontWorkflowRecordService interface {
		Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
		List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error
	}
	type FrontWorkflowRecordService struct {
		frontWorkflowRecordService
	}
	h := &frontWorkflowRecordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontWorkflowRecordService{h}, opts...))
}

type frontWorkflowRecordServiceHandler struct {
	FrontWorkflowRecordServiceHandler
}

func (h *frontWorkflowRecordServiceHandler) Get(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.FrontWorkflowRecordServiceHandler.Get(ctx, in, out)
}

func (h *frontWorkflowRecordServiceHandler) Search(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.FrontWorkflowRecordServiceHandler.Search(ctx, in, out)
}

func (h *frontWorkflowRecordServiceHandler) List(ctx context.Context, in *WorkflowRecordWhere, out *WorkflowRecordResponse) error {
	return h.FrontWorkflowRecordServiceHandler.List(ctx, in, out)
}
