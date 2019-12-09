// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package org

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

// Client is the client API for org service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	ImportModels(ctx context.Context, in *types.Empty) (*types.Empty, error)
	RegisterOrg(ctx context.Context, in *types.Empty) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) ImportModels(ctx context.Context, in *types.Empty) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _ImportModelsMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) RegisterOrg(ctx context.Context, in *types.Empty) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _RegisterOrgMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for org service.
type Service interface {
	ImportModels(context.Context, *types.Empty) (*types.Empty, error)
	RegisterOrg(context.Context, *types.Empty) (*types.Empty, error)
}

func _ImportModelsEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ImportModels(ctx, req.(*types.Empty))
	}
}

func _RegisterOrgEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RegisterOrg(ctx, req.(*types.Empty))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_ImportModelsMethodDesc, _ImportModelsEndpoint(srv))
	s.RegisterUnaryEndpoint(_RegisterOrgMethodDesc, _RegisterOrgEndpoint(srv))
}

// Method Description
var _ImportModelsMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.ucpro.org.ImportModels",
		Version: "1.0",
	},
	ServiceName:  "org.rpc",
	MethodName:   "ImportModels",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/org/import-models",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _RegisterOrgMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.ucpro.org.RegisterOrg",
		Version: "1.0",
	},
	ServiceName:  "org.rpc",
	MethodName:   "RegisterOrg",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/org/register",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
