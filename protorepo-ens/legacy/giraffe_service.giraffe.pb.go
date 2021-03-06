// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package legacy

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

// Client is the client API for legacy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	NameServiceConfig(ctx context.Context, in *NameServiceConfigRequest) (*NameServiceConfigResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) NameServiceConfig(ctx context.Context, in *NameServiceConfigRequest) (*NameServiceConfigResponse, error) {
	out := new(NameServiceConfigResponse)
	err := c.c.Invoke(ctx, _NameServiceConfigMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for legacy service.
type Service interface {
	NameServiceConfig(context.Context, *NameServiceConfigRequest) (*NameServiceConfigResponse, error)
}

func _NameServiceConfigEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NameServiceConfig(ctx, req.(*NameServiceConfigRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_NameServiceConfigMethodDesc, _NameServiceConfigEndpoint(srv))
}

// Method Description
var _NameServiceConfigMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.ens.legacy.NameServiceConfig",
		Version: "1.0",
	},
	ServiceName:  "legacy.rpc",
	MethodName:   "NameServiceConfig",
	RequestType:  (*NameServiceConfigRequest)(nil),
	ResponseType: (*NameServiceConfigResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/name_service/config",
		},
		Body:         "",
		ResponseBody: "",
	},
}
