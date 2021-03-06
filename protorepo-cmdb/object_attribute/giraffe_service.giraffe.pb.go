// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package object_attribute

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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

// Client is the client API for object_attribute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *CreateRequest) (*cmdb.ObjectAttr, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *CreateRequest) (*cmdb.ObjectAttr, error) {
	out := new(cmdb.ObjectAttr)
	err := c.c.Invoke(ctx, _CreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for object_attribute service.
type Service interface {
	Create(context.Context, *CreateRequest) (*cmdb.ObjectAttr, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMethodDesc, _CreateEndpoint(srv))
}

// Method Description
var _CreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.object_attribute.Create",
		Version: "1.0",
	},
	ServiceName:  "object_attribute.rpc",
	MethodName:   "Create",
	RequestType:  (*CreateRequest)(nil),
	ResponseType: (*cmdb.ObjectAttr)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/attr",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
