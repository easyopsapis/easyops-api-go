// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package alert

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

// Client is the client API for alert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	GetAlertCount(ctx context.Context, in *GetAlertCountRequest) (*GetAlertCountResponse, error)
	GetAlertEventDetail(ctx context.Context, in *GetAlertEventDetailRequest) (*GetAlertEventDetailResponse, error)
	GetAlertEventList(ctx context.Context, in *GetAlertEventListRequest) (*GetAlertEventListResponse, error)
	GetNotRecoverAlertEventList(ctx context.Context, in *GetNotRecoverAlertEventListRequest) (*GetNotRecoverAlertEventListResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) GetAlertCount(ctx context.Context, in *GetAlertCountRequest) (*GetAlertCountResponse, error) {
	out := new(GetAlertCountResponse)
	err := c.c.Invoke(ctx, _GetAlertCountMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetAlertEventDetail(ctx context.Context, in *GetAlertEventDetailRequest) (*GetAlertEventDetailResponse, error) {
	out := new(GetAlertEventDetailResponse)
	err := c.c.Invoke(ctx, _GetAlertEventDetailMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetAlertEventList(ctx context.Context, in *GetAlertEventListRequest) (*GetAlertEventListResponse, error) {
	out := new(GetAlertEventListResponse)
	err := c.c.Invoke(ctx, _GetAlertEventListMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetNotRecoverAlertEventList(ctx context.Context, in *GetNotRecoverAlertEventListRequest) (*GetNotRecoverAlertEventListResponse, error) {
	out := new(GetNotRecoverAlertEventListResponse)
	err := c.c.Invoke(ctx, _GetNotRecoverAlertEventListMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for alert service.
type Service interface {
	GetAlertCount(context.Context, *GetAlertCountRequest) (*GetAlertCountResponse, error)
	GetAlertEventDetail(context.Context, *GetAlertEventDetailRequest) (*GetAlertEventDetailResponse, error)
	GetAlertEventList(context.Context, *GetAlertEventListRequest) (*GetAlertEventListResponse, error)
	GetNotRecoverAlertEventList(context.Context, *GetNotRecoverAlertEventListRequest) (*GetNotRecoverAlertEventListResponse, error)
}

func _GetAlertCountEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAlertCount(ctx, req.(*GetAlertCountRequest))
	}
}

func _GetAlertEventDetailEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAlertEventDetail(ctx, req.(*GetAlertEventDetailRequest))
	}
}

func _GetAlertEventListEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAlertEventList(ctx, req.(*GetAlertEventListRequest))
	}
}

func _GetNotRecoverAlertEventListEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetNotRecoverAlertEventList(ctx, req.(*GetNotRecoverAlertEventListRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_GetAlertCountMethodDesc, _GetAlertCountEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetAlertEventDetailMethodDesc, _GetAlertEventDetailEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetAlertEventListMethodDesc, _GetAlertEventListEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetNotRecoverAlertEventListMethodDesc, _GetNotRecoverAlertEventListEndpoint(srv))
}

// Method Description
var _GetAlertCountMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.monitor.alert.GetAlertCount",
		Version: "1.0",
	},
	ServiceName:  "alert.rpc",
	MethodName:   "GetAlertCount",
	RequestType:  (*GetAlertCountRequest)(nil),
	ResponseType: (*GetAlertCountResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/alert/count",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetAlertEventDetailMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.monitor.alert.GetAlertEventDetail",
		Version: "1.0",
	},
	ServiceName:  "alert.rpc",
	MethodName:   "GetAlertEventDetail",
	RequestType:  (*GetAlertEventDetailRequest)(nil),
	ResponseType: (*GetAlertEventDetailResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/alert",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetAlertEventListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.monitor.alert.GetAlertEventList",
		Version: "1.0",
	},
	ServiceName:  "alert.rpc",
	MethodName:   "GetAlertEventList",
	RequestType:  (*GetAlertEventListRequest)(nil),
	ResponseType: (*GetAlertEventListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/alert",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetNotRecoverAlertEventListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.monitor.alert.GetNotRecoverAlertEventList",
		Version: "1.0",
	},
	ServiceName:  "alert.rpc",
	MethodName:   "GetNotRecoverAlertEventList",
	RequestType:  (*GetNotRecoverAlertEventListRequest)(nil),
	ResponseType: (*GetNotRecoverAlertEventListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/alert/alert_status_not_recover",
		},
		Body:         "",
		ResponseBody: "",
	},
}