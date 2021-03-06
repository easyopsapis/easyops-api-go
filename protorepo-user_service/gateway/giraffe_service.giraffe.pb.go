// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package gateway

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

// Client is the client API for gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	ListApiKey(ctx context.Context, in *ListApiKeyRequest) (*ListApiKeyResponse, error)
	UserRegister(ctx context.Context, in *UserRegisterRequest) (*UserRegisterResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) ListApiKey(ctx context.Context, in *ListApiKeyRequest) (*ListApiKeyResponse, error) {
	out := new(ListApiKeyResponse)
	err := c.c.Invoke(ctx, _ListApiKeyMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UserRegister(ctx context.Context, in *UserRegisterRequest) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.c.Invoke(ctx, _UserRegisterMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for gateway service.
type Service interface {
	ListApiKey(context.Context, *ListApiKeyRequest) (*ListApiKeyResponse, error)
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
}

func _ListApiKeyEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListApiKey(ctx, req.(*ListApiKeyRequest))
	}
}

func _UserRegisterEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UserRegister(ctx, req.(*UserRegisterRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_ListApiKeyMethodDesc, _ListApiKeyEndpoint(srv))
	s.RegisterUnaryEndpoint(_UserRegisterMethodDesc, _UserRegisterEndpoint(srv))
}

// Method Description
var _ListApiKeyMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.gateway.ListApiKey",
		Version: "1.0",
	},
	ServiceName:  "gateway.rpc",
	MethodName:   "ListApiKey",
	RequestType:  (*ListApiKeyRequest)(nil),
	ResponseType: (*ListApiKeyResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/user_service/v1/apikey",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UserRegisterMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.gateway.UserRegister",
		Version: "1.0",
	},
	ServiceName:  "gateway.rpc",
	MethodName:   "UserRegister",
	RequestType:  (*UserRegisterRequest)(nil),
	ResponseType: (*UserRegisterResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/user_service/v1/users/register",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
