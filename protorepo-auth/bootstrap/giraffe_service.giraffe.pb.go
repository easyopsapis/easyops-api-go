// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package bootstrap

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

// Client is the client API for bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Bootstrap(ctx context.Context, in *BootstrapRequest) (*BootstrapResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Bootstrap(ctx context.Context, in *BootstrapRequest) (*BootstrapResponse, error) {
	out := new(BootstrapResponse)
	err := c.c.Invoke(ctx, _BootstrapMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for bootstrap service.
type Service interface {
	Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error)
}

func _BootstrapEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Bootstrap(ctx, req.(*BootstrapRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_BootstrapMethodDesc, _BootstrapEndpoint(srv))
}

// Method Description
var _BootstrapMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.auth.bootstrap.Bootstrap",
		Version: "1.0",
	},
	ServiceName:  "bootstrap.rpc",
	MethodName:   "Bootstrap",
	RequestType:  (*BootstrapRequest)(nil),
	ResponseType: (*BootstrapResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/auth/bootstrap",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
