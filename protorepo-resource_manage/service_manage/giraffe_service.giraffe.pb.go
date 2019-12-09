// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package service_manage

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
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

// Client is the client API for service_manage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateServiceTask(ctx context.Context, in *CreateServiceTaskRequest) (*CreateServiceTaskResponse, error)
	CreateServiceTopologyTask(ctx context.Context, in *CreateServiceTopologyTaskRequest) (*CreateServiceTopologyTaskResponse, error)
	ListServiceInfo(ctx context.Context, in *ListServiceInfoRequest) (*ListServiceInfoResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateServiceTask(ctx context.Context, in *CreateServiceTaskRequest) (*CreateServiceTaskResponse, error) {
	out := new(CreateServiceTaskResponse)
	err := c.c.Invoke(ctx, _CreateServiceTaskMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateServiceTopologyTask(ctx context.Context, in *CreateServiceTopologyTaskRequest) (*CreateServiceTopologyTaskResponse, error) {
	out := new(CreateServiceTopologyTaskResponse)
	err := c.c.Invoke(ctx, _CreateServiceTopologyTaskMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListServiceInfo(ctx context.Context, in *ListServiceInfoRequest) (*ListServiceInfoResponse, error) {
	out := new(ListServiceInfoResponse)
	err := c.c.Invoke(ctx, _ListServiceInfoMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for service_manage service.
type Service interface {
	CreateServiceTask(context.Context, *CreateServiceTaskRequest) (*CreateServiceTaskResponse, error)
	CreateServiceTopologyTask(context.Context, *CreateServiceTopologyTaskRequest) (*CreateServiceTopologyTaskResponse, error)
	ListServiceInfo(context.Context, *ListServiceInfoRequest) (*ListServiceInfoResponse, error)
}

func _CreateServiceTaskEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateServiceTask(ctx, req.(*CreateServiceTaskRequest))
	}
}

func _CreateServiceTopologyTaskEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateServiceTopologyTask(ctx, req.(*CreateServiceTopologyTaskRequest))
	}
}

func _ListServiceInfoEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListServiceInfo(ctx, req.(*ListServiceInfoRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateServiceTaskMethodDesc, _CreateServiceTaskEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateServiceTopologyTaskMethodDesc, _CreateServiceTopologyTaskEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListServiceInfoMethodDesc, _ListServiceInfoEndpoint(srv))
}

// Method Description
var _CreateServiceTaskMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.resource_manage.service_manage.CreateServiceTask",
		Version: "1.0",
	},
	ServiceName:  "service_manage.rpc",
	MethodName:   "CreateServiceTask",
	RequestType:  (*CreateServiceTaskRequest)(nil),
	ResponseType: (*CreateServiceTaskResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/service/task",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _CreateServiceTopologyTaskMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.resource_manage.service_manage.CreateServiceTopologyTask",
		Version: "2.0",
	},
	ServiceName:  "service_manage.rpc",
	MethodName:   "CreateServiceTopologyTask",
	RequestType:  (*CreateServiceTopologyTaskRequest)(nil),
	ResponseType: (*CreateServiceTopologyTaskResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v2/service/task/topology",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListServiceInfoMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.resource_manage.service_manage.ListServiceInfo",
		Version: "1.0",
	},
	ServiceName:  "service_manage.rpc",
	MethodName:   "ListServiceInfo",
	RequestType:  (*ListServiceInfoRequest)(nil),
	ResponseType: (*ListServiceInfoResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/service/info",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
