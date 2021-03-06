// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package organization

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ = io.EOF
var _ context.Context
var _ giraffe_micro.Client
var _ go_proto_giraffe.Contract

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = giraffe_micro.SupportPackageIsVersion4 // please upgrade the giraffe_micro package

// Client is the client API for organization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateOrg(ctx context.Context, in *CreateOrgRequest) (*CreateOrgResponse, error)
	DisableOrg(ctx context.Context, in *DisableOrgRequest) (*types.Empty, error)
	EnableOrg(ctx context.Context, in *EnableOrgRequest) (*types.Empty, error)
	ListOrg(ctx context.Context, in *types.Empty) (*ListOrgResponse, error)
	SetOrgExpiredDate(ctx context.Context, in *SetOrgExpiredDateRequest) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateOrg(ctx context.Context, in *CreateOrgRequest) (*CreateOrgResponse, error) {
	out := new(CreateOrgResponse)
	err := c.c.Invoke(ctx, _CreateOrgMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DisableOrg(ctx context.Context, in *DisableOrgRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DisableOrgMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) EnableOrg(ctx context.Context, in *EnableOrgRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _EnableOrgMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListOrg(ctx context.Context, in *types.Empty) (*ListOrgResponse, error) {
	out := new(ListOrgResponse)
	err := c.c.Invoke(ctx, _ListOrgMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SetOrgExpiredDate(ctx context.Context, in *SetOrgExpiredDateRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _SetOrgExpiredDateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for organization service.
type Service interface {
	CreateOrg(context.Context, *CreateOrgRequest) (*CreateOrgResponse, error)
	DisableOrg(context.Context, *DisableOrgRequest) (*types.Empty, error)
	EnableOrg(context.Context, *EnableOrgRequest) (*types.Empty, error)
	ListOrg(context.Context, *types.Empty) (*ListOrgResponse, error)
	SetOrgExpiredDate(context.Context, *SetOrgExpiredDateRequest) (*types.Empty, error)
}

func _CreateOrgEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateOrg(ctx, req.(*CreateOrgRequest))
	}
}

func _DisableOrgEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DisableOrg(ctx, req.(*DisableOrgRequest))
	}
}

func _EnableOrgEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.EnableOrg(ctx, req.(*EnableOrgRequest))
	}
}

func _ListOrgEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListOrg(ctx, req.(*types.Empty))
	}
}

func _SetOrgExpiredDateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SetOrgExpiredDate(ctx, req.(*SetOrgExpiredDateRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateOrgMethodDesc, _CreateOrgEndpoint(srv))
	s.RegisterUnaryEndpoint(_DisableOrgMethodDesc, _DisableOrgEndpoint(srv))
	s.RegisterUnaryEndpoint(_EnableOrgMethodDesc, _EnableOrgEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListOrgMethodDesc, _ListOrgEndpoint(srv))
	s.RegisterUnaryEndpoint(_SetOrgExpiredDateMethodDesc, _SetOrgExpiredDateEndpoint(srv))
}

// Method Description
var _CreateOrgMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.organization.CreateOrg",
		Version: "1.0",
	},
	ServiceName:  "organization.rpc",
	MethodName:   "CreateOrg",
	RequestType:  (*CreateOrgRequest)(nil),
	ResponseType: (*CreateOrgResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/org",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DisableOrgMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.organization.DisableOrg",
		Version: "1.0",
	},
	ServiceName:  "organization.rpc",
	MethodName:   "DisableOrg",
	RequestType:  (*DisableOrgRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/v1/org/:id/disable",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _EnableOrgMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.organization.EnableOrg",
		Version: "1.0",
	},
	ServiceName:  "organization.rpc",
	MethodName:   "EnableOrg",
	RequestType:  (*EnableOrgRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/v1/org/:id/enable",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListOrgMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.organization.ListOrg",
		Version: "1.0",
	},
	ServiceName:  "organization.rpc",
	MethodName:   "ListOrg",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*ListOrgResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/org/list",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _SetOrgExpiredDateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.organization.SetOrgExpiredDate",
		Version: "1.0",
	},
	ServiceName:  "organization.rpc",
	MethodName:   "SetOrgExpiredDate",
	RequestType:  (*SetOrgExpiredDateRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/v1/org/:id/_expired_date",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
