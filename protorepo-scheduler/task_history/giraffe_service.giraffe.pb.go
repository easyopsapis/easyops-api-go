// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package task_history

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

// Client is the client API for task_history service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateHistory(ctx context.Context, in *CreateHistoryRequest) (*CreateHistoryResponse, error)
	ListHistory(ctx context.Context, in *ListHistoryRequest) (*ListHistoryResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateHistory(ctx context.Context, in *CreateHistoryRequest) (*CreateHistoryResponse, error) {
	out := new(CreateHistoryResponse)
	err := c.c.Invoke(ctx, _CreateHistoryMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListHistory(ctx context.Context, in *ListHistoryRequest) (*ListHistoryResponse, error) {
	out := new(ListHistoryResponse)
	err := c.c.Invoke(ctx, _ListHistoryMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for task_history service.
type Service interface {
	CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)
	ListHistory(context.Context, *ListHistoryRequest) (*ListHistoryResponse, error)
}

func _CreateHistoryEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateHistory(ctx, req.(*CreateHistoryRequest))
	}
}

func _ListHistoryEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListHistory(ctx, req.(*ListHistoryRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateHistoryMethodDesc, _CreateHistoryEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListHistoryMethodDesc, _ListHistoryEndpoint(srv))
}

// Method Description
var _CreateHistoryMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.scheduler.task_history.CreateHistory",
		Version: "1.0",
	},
	ServiceName:  "task_history.rpc",
	MethodName:   "CreateHistory",
	RequestType:  (*CreateHistoryRequest)(nil),
	ResponseType: (*CreateHistoryResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/scheduler/history",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _ListHistoryMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.scheduler.task_history.ListHistory",
		Version: "1.0",
	},
	ServiceName:  "task_history.rpc",
	MethodName:   "ListHistory",
	RequestType:  (*ListHistoryRequest)(nil),
	ResponseType: (*ListHistoryResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/scheduler/history",
		},
		Body:         "",
		ResponseBody: "",
	},
}
