// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package instance_path_search

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

// Client is the client API for instance_path_search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	PathSearchV3(ctx context.Context, in *PathSearchV3Request) (*PathSearchV3Response, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) PathSearchV3(ctx context.Context, in *PathSearchV3Request) (*PathSearchV3Response, error) {
	out := new(PathSearchV3Response)
	err := c.c.Invoke(ctx, _PathSearchV3MethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for instance_path_search service.
type Service interface {
	PathSearchV3(context.Context, *PathSearchV3Request) (*PathSearchV3Response, error)
}

func _PathSearchV3Endpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PathSearchV3(ctx, req.(*PathSearchV3Request))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_PathSearchV3MethodDesc, _PathSearchV3Endpoint(srv))
}

// Method Description
var _PathSearchV3MethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance_path_search.PathSearchV3",
		Version: "1.0",
	},
	ServiceName:  "instance_path_search.rpc",
	MethodName:   "PathSearchV3",
	RequestType:  (*PathSearchV3Request)(nil),
	ResponseType: (*PathSearchV3Response)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/v3/path/_search",
		},
		Body:         "",
		ResponseBody: "data",
	},
}