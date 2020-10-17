// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: statistic/message.proto

package github_com_panovateam_rbac_proto_statistic

import (
	api "api"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SystemSvc service

type SystemSvcService interface {
	GetDPChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error)
	GetDPHighChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error)
	GetDPLowChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error)
}

type systemSvcService struct {
	c    client.Client
	name string
}

func NewSystemSvcService(name string, c client.Client) SystemSvcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "system.statistic"
	}
	return &systemSvcService{
		c:    c,
		name: name,
	}
}

func (c *systemSvcService) GetDPChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetDPChartAPI", in)
	out := new(api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) GetDPHighChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetDPHighChartAPI", in)
	out := new(api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemSvcService) GetDPLowChartAPI(ctx context.Context, in *api.Request, opts ...client.CallOption) (*api.Response, error) {
	req := c.c.NewRequest(c.name, "SystemSvc.GetDPLowChartAPI", in)
	out := new(api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SystemSvc service

type SystemSvcHandler interface {
	GetDPChartAPI(context.Context, *api.Request, *api.Response) error
	GetDPHighChartAPI(context.Context, *api.Request, *api.Response) error
	GetDPLowChartAPI(context.Context, *api.Request, *api.Response) error
}

func RegisterSystemSvcHandler(s server.Server, hdlr SystemSvcHandler, opts ...server.HandlerOption) error {
	type systemSvc interface {
		GetDPChartAPI(ctx context.Context, in *api.Request, out *api.Response) error
		GetDPHighChartAPI(ctx context.Context, in *api.Request, out *api.Response) error
		GetDPLowChartAPI(ctx context.Context, in *api.Request, out *api.Response) error
	}
	type SystemSvc struct {
		systemSvc
	}
	h := &systemSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&SystemSvc{h}, opts...))
}

type systemSvcHandler struct {
	SystemSvcHandler
}

func (h *systemSvcHandler) GetDPChartAPI(ctx context.Context, in *api.Request, out *api.Response) error {
	return h.SystemSvcHandler.GetDPChartAPI(ctx, in, out)
}

func (h *systemSvcHandler) GetDPHighChartAPI(ctx context.Context, in *api.Request, out *api.Response) error {
	return h.SystemSvcHandler.GetDPHighChartAPI(ctx, in, out)
}

func (h *systemSvcHandler) GetDPLowChartAPI(ctx context.Context, in *api.Request, out *api.Response) error {
	return h.SystemSvcHandler.GetDPLowChartAPI(ctx, in, out)
}
