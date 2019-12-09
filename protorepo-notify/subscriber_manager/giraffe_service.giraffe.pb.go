// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package subscriber_manager

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

// Client is the client API for subscriber_manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	PubSubscriberCreate(ctx context.Context, in *PubSubscriberCreateRequest) (*PubSubscriberCreateResponse, error)
	PubSubscriberDelete(ctx context.Context, in *PubSubscriberDeleteRequest) (*types.Empty, error)
	PubSubscriberList(ctx context.Context, in *PubSubscriberListRequest) (*PubSubscriberListResponse, error)
	PubSubscriberUpdate(ctx context.Context, in *PubSubscriberUpdateRequest) (*PubSubscriberUpdateResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) PubSubscriberCreate(ctx context.Context, in *PubSubscriberCreateRequest) (*PubSubscriberCreateResponse, error) {
	out := new(PubSubscriberCreateResponse)
	err := c.c.Invoke(ctx, _PubSubscriberCreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) PubSubscriberDelete(ctx context.Context, in *PubSubscriberDeleteRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _PubSubscriberDeleteMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) PubSubscriberList(ctx context.Context, in *PubSubscriberListRequest) (*PubSubscriberListResponse, error) {
	out := new(PubSubscriberListResponse)
	err := c.c.Invoke(ctx, _PubSubscriberListMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) PubSubscriberUpdate(ctx context.Context, in *PubSubscriberUpdateRequest) (*PubSubscriberUpdateResponse, error) {
	out := new(PubSubscriberUpdateResponse)
	err := c.c.Invoke(ctx, _PubSubscriberUpdateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for subscriber_manager service.
type Service interface {
	PubSubscriberCreate(context.Context, *PubSubscriberCreateRequest) (*PubSubscriberCreateResponse, error)
	PubSubscriberDelete(context.Context, *PubSubscriberDeleteRequest) (*types.Empty, error)
	PubSubscriberList(context.Context, *PubSubscriberListRequest) (*PubSubscriberListResponse, error)
	PubSubscriberUpdate(context.Context, *PubSubscriberUpdateRequest) (*PubSubscriberUpdateResponse, error)
}

func _PubSubscriberCreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PubSubscriberCreate(ctx, req.(*PubSubscriberCreateRequest))
	}
}

func _PubSubscriberDeleteEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PubSubscriberDelete(ctx, req.(*PubSubscriberDeleteRequest))
	}
}

func _PubSubscriberListEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PubSubscriberList(ctx, req.(*PubSubscriberListRequest))
	}
}

func _PubSubscriberUpdateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PubSubscriberUpdate(ctx, req.(*PubSubscriberUpdateRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_PubSubscriberCreateMethodDesc, _PubSubscriberCreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_PubSubscriberDeleteMethodDesc, _PubSubscriberDeleteEndpoint(srv))
	s.RegisterUnaryEndpoint(_PubSubscriberListMethodDesc, _PubSubscriberListEndpoint(srv))
	s.RegisterUnaryEndpoint(_PubSubscriberUpdateMethodDesc, _PubSubscriberUpdateEndpoint(srv))
}

// Method Description
var _PubSubscriberCreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.notify.subscriber_manager.PubSubscriberCreate",
		Version: "1.0",
	},
	ServiceName:  "subscriber_manager.rpc",
	MethodName:   "PubSubscriberCreate",
	RequestType:  (*PubSubscriberCreateRequest)(nil),
	ResponseType: (*PubSubscriberCreateResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/subscriber_manager/v1/subscribers",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _PubSubscriberDeleteMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.notify.subscriber_manager.PubSubscriberDelete",
		Version: "1.0",
	},
	ServiceName:  "subscriber_manager.rpc",
	MethodName:   "PubSubscriberDelete",
	RequestType:  (*PubSubscriberDeleteRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/subscriber_manager/v1/subscribers/:subscriberId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _PubSubscriberListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.notify.subscriber_manager.PubSubscriberList",
		Version: "1.0",
	},
	ServiceName:  "subscriber_manager.rpc",
	MethodName:   "PubSubscriberList",
	RequestType:  (*PubSubscriberListRequest)(nil),
	ResponseType: (*PubSubscriberListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/subscriber_manager/v1/subscribers",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _PubSubscriberUpdateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.notify.subscriber_manager.PubSubscriberUpdate",
		Version: "1.0",
	},
	ServiceName:  "subscriber_manager.rpc",
	MethodName:   "PubSubscriberUpdate",
	RequestType:  (*PubSubscriberUpdateRequest)(nil),
	ResponseType: (*PubSubscriberUpdateResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/subscriber_manager/v1/subscribers/:subscriberId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}