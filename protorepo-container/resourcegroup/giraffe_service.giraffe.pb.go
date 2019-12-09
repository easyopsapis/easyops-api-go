// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package resourcegroup

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

// Client is the client API for resourcegroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	DeleteResourceGroup(ctx context.Context, in *DeleteResourceGroupRequest) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) DeleteResourceGroup(ctx context.Context, in *DeleteResourceGroupRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteResourceGroupMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for resourcegroup service.
type Service interface {
	DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest) (*types.Empty, error)
}

func _DeleteResourceGroupEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteResourceGroup(ctx, req.(*DeleteResourceGroupRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_DeleteResourceGroupMethodDesc, _DeleteResourceGroupEndpoint(srv))
}

// Method Description
var _DeleteResourceGroupMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.resourcegroup.DeleteResourceGroup",
		Version: "1.0",
	},
	ServiceName:  "resourcegroup.rpc",
	MethodName:   "DeleteResourceGroup",
	RequestType:  (*DeleteResourceGroupRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/container/v1/resourcegroups/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}