// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package pipeline

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

// Client is the client API for pipeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *CreateRequest) (*pipeline.Pipeline, error)
	CreateTrigger(ctx context.Context, in *CreateTriggerRequest) (*pipeline.Trigger, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest) (*types.Empty, error)
	DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest) (*types.Empty, error)
	DeleteTriggers(ctx context.Context, in *DeleteTriggersRequest) (*types.Empty, error)
	Execute(ctx context.Context, in *ExecuteRequest) (*ExecuteResponse, error)
	Get(ctx context.Context, in *GetRequest) (*GetResponse, error)
	GetTrigger(ctx context.Context, in *GetTriggerRequest) (*pipeline.Trigger, error)
	GetTriggerDetail(ctx context.Context, in *GetTriggerDetailRequest) (*GetTriggerDetailResponse, error)
	List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	ListTrigger(ctx context.Context, in *ListTriggerRequest) (*ListTriggerResponse, error)
	Update(ctx context.Context, in *UpdateRequest) (*pipeline.Pipeline, error)
	UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest) (*pipeline.Trigger, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *CreateRequest) (*pipeline.Pipeline, error) {
	out := new(pipeline.Pipeline)
	err := c.c.Invoke(ctx, _CreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateTrigger(ctx context.Context, in *CreateTriggerRequest) (*pipeline.Trigger, error) {
	out := new(pipeline.Trigger)
	err := c.c.Invoke(ctx, _CreateTriggerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeletePipeline(ctx context.Context, in *DeletePipelineRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeletePipelineMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteTriggerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTriggers(ctx context.Context, in *DeleteTriggersRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteTriggersMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Execute(ctx context.Context, in *ExecuteRequest) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.c.Invoke(ctx, _ExecuteMethodDesc, in, out)
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

func (c *client) GetTrigger(ctx context.Context, in *GetTriggerRequest) (*pipeline.Trigger, error) {
	out := new(pipeline.Trigger)
	err := c.c.Invoke(ctx, _GetTriggerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetTriggerDetail(ctx context.Context, in *GetTriggerDetailRequest) (*GetTriggerDetailResponse, error) {
	out := new(GetTriggerDetailResponse)
	err := c.c.Invoke(ctx, _GetTriggerDetailMethodDesc, in, out)
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

func (c *client) ListTrigger(ctx context.Context, in *ListTriggerRequest) (*ListTriggerResponse, error) {
	out := new(ListTriggerResponse)
	err := c.c.Invoke(ctx, _ListTriggerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Update(ctx context.Context, in *UpdateRequest) (*pipeline.Pipeline, error) {
	out := new(pipeline.Pipeline)
	err := c.c.Invoke(ctx, _UpdateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest) (*pipeline.Trigger, error) {
	out := new(pipeline.Trigger)
	err := c.c.Invoke(ctx, _UpdateTriggerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for pipeline service.
type Service interface {
	Create(context.Context, *CreateRequest) (*pipeline.Pipeline, error)
	CreateTrigger(context.Context, *CreateTriggerRequest) (*pipeline.Trigger, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*types.Empty, error)
	DeleteTrigger(context.Context, *DeleteTriggerRequest) (*types.Empty, error)
	DeleteTriggers(context.Context, *DeleteTriggersRequest) (*types.Empty, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetTrigger(context.Context, *GetTriggerRequest) (*pipeline.Trigger, error)
	GetTriggerDetail(context.Context, *GetTriggerDetailRequest) (*GetTriggerDetailResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	ListTrigger(context.Context, *ListTriggerRequest) (*ListTriggerResponse, error)
	Update(context.Context, *UpdateRequest) (*pipeline.Pipeline, error)
	UpdateTrigger(context.Context, *UpdateTriggerRequest) (*pipeline.Trigger, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
}

func _CreateTriggerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateTrigger(ctx, req.(*CreateTriggerRequest))
	}
}

func _DeletePipelineEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
}

func _DeleteTriggerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteTrigger(ctx, req.(*DeleteTriggerRequest))
	}
}

func _DeleteTriggersEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteTriggers(ctx, req.(*DeleteTriggersRequest))
	}
}

func _ExecuteEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Execute(ctx, req.(*ExecuteRequest))
	}
}

func _GetEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Get(ctx, req.(*GetRequest))
	}
}

func _GetTriggerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetTrigger(ctx, req.(*GetTriggerRequest))
	}
}

func _GetTriggerDetailEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetTriggerDetail(ctx, req.(*GetTriggerDetailRequest))
	}
}

func _ListEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx, req.(*ListRequest))
	}
}

func _ListTriggerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListTrigger(ctx, req.(*ListTriggerRequest))
	}
}

func _UpdateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Update(ctx, req.(*UpdateRequest))
	}
}

func _UpdateTriggerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateTrigger(ctx, req.(*UpdateTriggerRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMethodDesc, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateTriggerMethodDesc, _CreateTriggerEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeletePipelineMethodDesc, _DeletePipelineEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteTriggerMethodDesc, _DeleteTriggerEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteTriggersMethodDesc, _DeleteTriggersEndpoint(srv))
	s.RegisterUnaryEndpoint(_ExecuteMethodDesc, _ExecuteEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetMethodDesc, _GetEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetTriggerMethodDesc, _GetTriggerEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetTriggerDetailMethodDesc, _GetTriggerDetailEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListMethodDesc, _ListEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListTriggerMethodDesc, _ListTriggerEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateMethodDesc, _UpdateEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateTriggerMethodDesc, _UpdateTriggerEndpoint(srv))
}

// Method Description
var _CreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.Create",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "Create",
	RequestType:  (*CreateRequest)(nil),
	ResponseType: (*pipeline.Pipeline)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/pipeline/v1/projects/:project_id/pipelines",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _CreateTriggerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.CreateTrigger",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "CreateTrigger",
	RequestType:  (*CreateTriggerRequest)(nil),
	ResponseType: (*pipeline.Trigger)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/pipeline/v1/triggers",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeletePipelineMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.DeletePipeline",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "DeletePipeline",
	RequestType:  (*DeletePipelineRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/pipeline/v1/projects/:project_id/pipelines/:pipeline_id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteTriggerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.DeleteTrigger",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "DeleteTrigger",
	RequestType:  (*DeleteTriggerRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/pipeline/v1/triggers/:id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteTriggersMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.DeleteTriggers",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "DeleteTriggers",
	RequestType:  (*DeleteTriggersRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/pipeline/v1/triggers",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ExecuteMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.Execute",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "Execute",
	RequestType:  (*ExecuteRequest)(nil),
	ResponseType: (*ExecuteResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/pipeline/v1/projects/:project_id/pipelines/:pipeline_id/execute",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.Get",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "Get",
	RequestType:  (*GetRequest)(nil),
	ResponseType: (*GetResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/projects/:project_id/pipelines/:pipeline_id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetTriggerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.GetTrigger",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "GetTrigger",
	RequestType:  (*GetTriggerRequest)(nil),
	ResponseType: (*pipeline.Trigger)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/triggers/:id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetTriggerDetailMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.GetTriggerDetail",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "GetTriggerDetail",
	RequestType:  (*GetTriggerDetailRequest)(nil),
	ResponseType: (*GetTriggerDetailResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/triggers/:id/detail",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.List",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "List",
	RequestType:  (*ListRequest)(nil),
	ResponseType: (*ListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/projects/:project_id/pipelines",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListTriggerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.ListTrigger",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "ListTrigger",
	RequestType:  (*ListTriggerRequest)(nil),
	ResponseType: (*ListTriggerResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/pipeline/v1/triggers",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.Update",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "Update",
	RequestType:  (*UpdateRequest)(nil),
	ResponseType: (*pipeline.Pipeline)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/pipeline/v1/projects/:project_id/pipelines/:pipeline_id",
		},
		Body:         "pipeline",
		ResponseBody: "data",
	},
}

var _UpdateTriggerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.pipeline.pipeline.UpdateTrigger",
		Version: "1.0",
	},
	ServiceName:  "pipeline.rpc",
	MethodName:   "UpdateTrigger",
	RequestType:  (*UpdateTriggerRequest)(nil),
	ResponseType: (*pipeline.Trigger)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/pipeline/v1/triggers/:id",
		},
		Body:         "trigger",
		ResponseBody: "data",
	},
}
