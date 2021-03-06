// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package custom_sender

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	msgsender "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/msgsender"
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

// Client is the client API for custom_sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	ListSupportInform(ctx context.Context, in *types.Empty) (*ListSupportInformResponse, error)
	SendMessage(ctx context.Context, in *msgsender.SendMessageRequest) (*SendMessageResponse, error)
	SendMessageForAlert(ctx context.Context, in *msgsender.SendMessageForAlertRequest) (*SendMessageForAlertResponse, error)
	SendMessageWithAppendix(ctx context.Context, in *msgsender.SendMessageWithAppendixRequest) (*SendMessageWithAppendixResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) ListSupportInform(ctx context.Context, in *types.Empty) (*ListSupportInformResponse, error) {
	out := new(ListSupportInformResponse)
	err := c.c.Invoke(ctx, _ListSupportInformMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SendMessage(ctx context.Context, in *msgsender.SendMessageRequest) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.c.Invoke(ctx, _SendMessageMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SendMessageForAlert(ctx context.Context, in *msgsender.SendMessageForAlertRequest) (*SendMessageForAlertResponse, error) {
	out := new(SendMessageForAlertResponse)
	err := c.c.Invoke(ctx, _SendMessageForAlertMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SendMessageWithAppendix(ctx context.Context, in *msgsender.SendMessageWithAppendixRequest) (*SendMessageWithAppendixResponse, error) {
	out := new(SendMessageWithAppendixResponse)
	err := c.c.Invoke(ctx, _SendMessageWithAppendixMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for custom_sender service.
type Service interface {
	ListSupportInform(context.Context, *types.Empty) (*ListSupportInformResponse, error)
	SendMessage(context.Context, *msgsender.SendMessageRequest) (*SendMessageResponse, error)
	SendMessageForAlert(context.Context, *msgsender.SendMessageForAlertRequest) (*SendMessageForAlertResponse, error)
	SendMessageWithAppendix(context.Context, *msgsender.SendMessageWithAppendixRequest) (*SendMessageWithAppendixResponse, error)
}

func _ListSupportInformEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListSupportInform(ctx, req.(*types.Empty))
	}
}

func _SendMessageEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SendMessage(ctx, req.(*msgsender.SendMessageRequest))
	}
}

func _SendMessageForAlertEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SendMessageForAlert(ctx, req.(*msgsender.SendMessageForAlertRequest))
	}
}

func _SendMessageWithAppendixEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SendMessageWithAppendix(ctx, req.(*msgsender.SendMessageWithAppendixRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_ListSupportInformMethodDesc, _ListSupportInformEndpoint(srv))
	s.RegisterUnaryEndpoint(_SendMessageMethodDesc, _SendMessageEndpoint(srv))
	s.RegisterUnaryEndpoint(_SendMessageForAlertMethodDesc, _SendMessageForAlertEndpoint(srv))
	s.RegisterUnaryEndpoint(_SendMessageWithAppendixMethodDesc, _SendMessageWithAppendixEndpoint(srv))
}

// Method Description
var _ListSupportInformMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.msgsender.custom_sender.ListSupportInform",
		Version: "1.0",
	},
	ServiceName:  "custom_sender.rpc",
	MethodName:   "ListSupportInform",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*ListSupportInformResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/message_sender/method",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _SendMessageMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.msgsender.custom_sender.SendMessage",
		Version: "1.0",
	},
	ServiceName:  "custom_sender.rpc",
	MethodName:   "SendMessage",
	RequestType:  (*msgsender.SendMessageRequest)(nil),
	ResponseType: (*SendMessageResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/message_sender/send_message",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _SendMessageForAlertMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.msgsender.custom_sender.SendMessageForAlert",
		Version: "1.0",
	},
	ServiceName:  "custom_sender.rpc",
	MethodName:   "SendMessageForAlert",
	RequestType:  (*msgsender.SendMessageForAlertRequest)(nil),
	ResponseType: (*SendMessageForAlertResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/alert_adapter/receive",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _SendMessageWithAppendixMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.msgsender.custom_sender.SendMessageWithAppendix",
		Version: "1.0",
	},
	ServiceName:  "custom_sender.rpc",
	MethodName:   "SendMessageWithAppendix",
	RequestType:  (*msgsender.SendMessageWithAppendixRequest)(nil),
	ResponseType: (*SendMessageWithAppendixResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/message_sender/mail_with_appendix",
		},
		Body:         "",
		ResponseBody: "",
	},
}
