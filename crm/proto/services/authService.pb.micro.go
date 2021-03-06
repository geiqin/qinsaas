// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authService.proto

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

// Api Endpoints for AuthService service

func NewAuthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthService service

type AuthService interface {
	//客户登录
	Login(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*common.TokenResponse, error)
	//修改密码
	UpdatePwd(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*common.TokenResponse, error)
	//获得当前客户
	Info(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CustomerResponse, error)
	//安全退出
	Logout(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*AuthResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Login(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*common.TokenResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.Login", in)
	out := new(common.TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdatePwd(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*common.TokenResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.UpdatePwd", in)
	out := new(common.TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Info(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.Info", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Logout(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.Logout", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	//客户登录
	Login(context.Context, *AuthRequest, *common.TokenResponse) error
	//修改密码
	UpdatePwd(context.Context, *AuthRequest, *common.TokenResponse) error
	//获得当前客户
	Info(context.Context, *common.Empty, *CustomerResponse) error
	//安全退出
	Logout(context.Context, *common.Empty, *AuthResponse) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) error {
	type authService interface {
		Login(ctx context.Context, in *AuthRequest, out *common.TokenResponse) error
		UpdatePwd(ctx context.Context, in *AuthRequest, out *common.TokenResponse) error
		Info(ctx context.Context, in *common.Empty, out *CustomerResponse) error
		Logout(ctx context.Context, in *common.Empty, out *AuthResponse) error
	}
	type AuthService struct {
		authService
	}
	h := &authServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthService{h}, opts...))
}

type authServiceHandler struct {
	AuthServiceHandler
}

func (h *authServiceHandler) Login(ctx context.Context, in *AuthRequest, out *common.TokenResponse) error {
	return h.AuthServiceHandler.Login(ctx, in, out)
}

func (h *authServiceHandler) UpdatePwd(ctx context.Context, in *AuthRequest, out *common.TokenResponse) error {
	return h.AuthServiceHandler.UpdatePwd(ctx, in, out)
}

func (h *authServiceHandler) Info(ctx context.Context, in *common.Empty, out *CustomerResponse) error {
	return h.AuthServiceHandler.Info(ctx, in, out)
}

func (h *authServiceHandler) Logout(ctx context.Context, in *common.Empty, out *AuthResponse) error {
	return h.AuthServiceHandler.Logout(ctx, in, out)
}
