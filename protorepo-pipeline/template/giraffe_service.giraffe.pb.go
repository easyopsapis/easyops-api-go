// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package template

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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

// Client is the client API for template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *pipeline.PipelineTemplate) (*pipeline.PipelineTemplate, error)
	DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest) (*types.Empty, error)
	Get(ctx context.Context, in *GetRequest) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateRequest) (*pipeline.PipelineTemplate, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *pipeline.PipelineTemplate) (*pipeline.PipelineTemplate, error) {
	out := new(pipeline.PipelineTemplate)
	err := c.c.Invoke(ctx, _CreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteTemplateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.c.Invoke(ctx, _GetMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.c.Invoke(ctx, _ListMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Update(ctx context.Context, in *UpdateRequest) (*pipeline.PipelineTemplate, error) {
	out := new(pipeline.PipelineTemplate)
	err := c.c.Invoke(ctx, _UpdateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for template service.
type Service interface {
	Create(context.Context, *pipeline.PipelineTemplate) (*pipeline.PipelineTemplate, error)
	DeleteTemplate(context.Context, *DeleteTemplateRequest) (*types.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Update(context.Context, *UpdateRequest) (*pipeline.PipelineTemplate, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*pipeline.PipelineTemplate))
	}
}

func _DeleteTemplateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteTemplate(ctx, req.(*DeleteTemplateRequest))
	}
}

func _GetEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Get(ctx, req.(*GetRequest))
	}
}

func _ListEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx, req.(*ListRequest))
	}
}

func _UpdateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Update(ctx, req.(*UpdateRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMethodDesc, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteTemplateMethodDesc, _DeleteTemplateEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetMethodDesc, _GetEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListMethodDesc, _ListEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateMethodDesc, _UpdateEndpoint(srv))
}

// Method Description
var _CreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.template.Create",
		Version: "1.0",
	},
	ServiceName:  "template.rpc",
	MethodName:   "Create",
	RequestType:  (*pipeline.PipelineTemplate)(nil),
	ResponseType: (*pipeline.PipelineTemplate)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/pipeline/v1/templates",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteTemplateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.template.DeleteTemplate",
		Version: "1.0",
	},
	ServiceName:  "template.rpc",
	MethodName:   "DeleteTemplate",
	RequestType:  (*DeleteTemplateRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/pipeline/v1/templates/:id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.template.Get",
		Version: "1.0",
	},
	ServiceName:  "template.rpc",
	MethodName:   "Get",
	RequestType:  (*GetRequest)(nil),
	ResponseType: (*GetResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/templates/:id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.template.List",
		Version: "1.0",
	},
	ServiceName:  "template.rpc",
	MethodName:   "List",
	RequestType:  (*ListRequest)(nil),
	ResponseType: (*ListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/pipeline/v1/templates/list",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.template.Update",
		Version: "1.0",
	},
	ServiceName:  "template.rpc",
	MethodName:   "Update",
	RequestType:  (*UpdateRequest)(nil),
	ResponseType: (*pipeline.PipelineTemplate)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/pipeline/v1/templates/:id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
