// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package translate_schema

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

// Client is the client API for translate_schema service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	TranslateDataHandler(ctx context.Context, in *TranslateDataHandlerRequest) (*TranslateDataHandlerResponse, error)
	TranslateSchema(ctx context.Context, in *TranslateSchemaRequest) (*TranslateSchemaResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) TranslateDataHandler(ctx context.Context, in *TranslateDataHandlerRequest) (*TranslateDataHandlerResponse, error) {
	out := new(TranslateDataHandlerResponse)
	err := c.c.Invoke(ctx, _TranslateDataHandlerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) TranslateSchema(ctx context.Context, in *TranslateSchemaRequest) (*TranslateSchemaResponse, error) {
	out := new(TranslateSchemaResponse)
	err := c.c.Invoke(ctx, _TranslateSchemaMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for translate_schema service.
type Service interface {
	TranslateDataHandler(context.Context, *TranslateDataHandlerRequest) (*TranslateDataHandlerResponse, error)
	TranslateSchema(context.Context, *TranslateSchemaRequest) (*TranslateSchemaResponse, error)
}

func _TranslateDataHandlerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.TranslateDataHandler(ctx, req.(*TranslateDataHandlerRequest))
	}
}

func _TranslateSchemaEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.TranslateSchema(ctx, req.(*TranslateSchemaRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_TranslateDataHandlerMethodDesc, _TranslateDataHandlerEndpoint(srv))
	s.RegisterUnaryEndpoint(_TranslateSchemaMethodDesc, _TranslateSchemaEndpoint(srv))
}

// Method Description
var _TranslateDataHandlerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.translate.translate_schema.TranslateDataHandler",
		Version: "1.0",
	},
	ServiceName:  "translate_schema.rpc",
	MethodName:   "TranslateDataHandler",
	RequestType:  (*TranslateDataHandlerRequest)(nil),
	ResponseType: (*TranslateDataHandlerResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v2/translate/storm/data",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _TranslateSchemaMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.translate.translate_schema.TranslateSchema",
		Version: "1.0",
	},
	ServiceName:  "translate_schema.rpc",
	MethodName:   "TranslateSchema",
	RequestType:  (*TranslateSchemaRequest)(nil),
	ResponseType: (*TranslateSchemaResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/translate/storm/schema",
		},
		Body:         "",
		ResponseBody: "",
	},
}