// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package solution

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

// Client is the client API for solution service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	GetSolution(ctx context.Context, in *GetSolutionRequest) (*GetSolutionResponse, error)
	ListSolutions(ctx context.Context, in *ListSolutionsRequest) (*ListSolutionsResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) GetSolution(ctx context.Context, in *GetSolutionRequest) (*GetSolutionResponse, error) {
	out := new(GetSolutionResponse)
	err := c.c.Invoke(ctx, _GetSolutionMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListSolutions(ctx context.Context, in *ListSolutionsRequest) (*ListSolutionsResponse, error) {
	out := new(ListSolutionsResponse)
	err := c.c.Invoke(ctx, _ListSolutionsMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for solution service.
type Service interface {
	GetSolution(context.Context, *GetSolutionRequest) (*GetSolutionResponse, error)
	ListSolutions(context.Context, *ListSolutionsRequest) (*ListSolutionsResponse, error)
}

func _GetSolutionEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSolution(ctx, req.(*GetSolutionRequest))
	}
}

func _ListSolutionsEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListSolutions(ctx, req.(*ListSolutionsRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_GetSolutionMethodDesc, _GetSolutionEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListSolutionsMethodDesc, _ListSolutionsEndpoint(srv))
}

// Method Description
var _GetSolutionMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.app_store.solution.GetSolution",
		Version: "1.0",
	},
	ServiceName:  "solution.rpc",
	MethodName:   "GetSolution",
	RequestType:  (*GetSolutionRequest)(nil),
	ResponseType: (*GetSolutionResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/app_store/v1/solution/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListSolutionsMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.app_store.solution.ListSolutions",
		Version: "1.0",
	},
	ServiceName:  "solution.rpc",
	MethodName:   "ListSolutions",
	RequestType:  (*ListSolutionsRequest)(nil),
	ResponseType: (*ListSolutionsResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/app_store/v1/search/solution",
		},
		Body:         "",
		ResponseBody: "data",
	},
}