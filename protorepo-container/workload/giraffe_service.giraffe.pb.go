// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package workload

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	container "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
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

// Client is the client API for workload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *CreateRequest) (*container.Workload, error)
	CreateFromYaml(ctx context.Context, in *CreateFromYamlRequest) (*container.Workload, error)
	DeleteWorkload(ctx context.Context, in *DeleteWorkloadRequest) (*types.Empty, error)
	Get(ctx context.Context, in *GetRequest) (*GetResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest) (*container.Workload, error)
	GetSummary(ctx context.Context, in *GetSummaryRequest) (*GetSummaryResponse, error)
	List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateRequest) (*container.Workload, error)
	UpdateResourceSpec(ctx context.Context, in *UpdateResourceSpecRequest) (*container.Workload, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *CreateRequest) (*container.Workload, error) {
	out := new(container.Workload)
	err := c.c.Invoke(ctx, _CreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateFromYaml(ctx context.Context, in *CreateFromYamlRequest) (*container.Workload, error) {
	out := new(container.Workload)
	err := c.c.Invoke(ctx, _CreateFromYamlMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteWorkload(ctx context.Context, in *DeleteWorkloadRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteWorkloadMethodDesc, in, out)
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

func (c *client) GetStatus(ctx context.Context, in *GetStatusRequest) (*container.Workload, error) {
	out := new(container.Workload)
	err := c.c.Invoke(ctx, _GetStatusMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetSummary(ctx context.Context, in *GetSummaryRequest) (*GetSummaryResponse, error) {
	out := new(GetSummaryResponse)
	err := c.c.Invoke(ctx, _GetSummaryMethodDesc, in, out)
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

func (c *client) Update(ctx context.Context, in *UpdateRequest) (*container.Workload, error) {
	out := new(container.Workload)
	err := c.c.Invoke(ctx, _UpdateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateResourceSpec(ctx context.Context, in *UpdateResourceSpecRequest) (*container.Workload, error) {
	out := new(container.Workload)
	err := c.c.Invoke(ctx, _UpdateResourceSpecMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for workload service.
type Service interface {
	Create(context.Context, *CreateRequest) (*container.Workload, error)
	CreateFromYaml(context.Context, *CreateFromYamlRequest) (*container.Workload, error)
	DeleteWorkload(context.Context, *DeleteWorkloadRequest) (*types.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*container.Workload, error)
	GetSummary(context.Context, *GetSummaryRequest) (*GetSummaryResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Update(context.Context, *UpdateRequest) (*container.Workload, error)
	UpdateResourceSpec(context.Context, *UpdateResourceSpecRequest) (*container.Workload, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
}

func _CreateFromYamlEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateFromYaml(ctx, req.(*CreateFromYamlRequest))
	}
}

func _DeleteWorkloadEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteWorkload(ctx, req.(*DeleteWorkloadRequest))
	}
}

func _GetEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Get(ctx, req.(*GetRequest))
	}
}

func _GetStatusEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetStatus(ctx, req.(*GetStatusRequest))
	}
}

func _GetSummaryEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSummary(ctx, req.(*GetSummaryRequest))
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

func _UpdateResourceSpecEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateResourceSpec(ctx, req.(*UpdateResourceSpecRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMethodDesc, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateFromYamlMethodDesc, _CreateFromYamlEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteWorkloadMethodDesc, _DeleteWorkloadEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetMethodDesc, _GetEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetStatusMethodDesc, _GetStatusEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetSummaryMethodDesc, _GetSummaryEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListMethodDesc, _ListEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateMethodDesc, _UpdateEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateResourceSpecMethodDesc, _UpdateResourceSpecEndpoint(srv))
}

// Method Description
var _CreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.Create",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "Create",
	RequestType:  (*CreateRequest)(nil),
	ResponseType: (*container.Workload)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/container/v1/workloads",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _CreateFromYamlMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.CreateFromYaml",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "CreateFromYaml",
	RequestType:  (*CreateFromYamlRequest)(nil),
	ResponseType: (*container.Workload)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/container/v1/workloads/yaml",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteWorkloadMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.DeleteWorkload",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "DeleteWorkload",
	RequestType:  (*DeleteWorkloadRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/container/v1/workloads/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.Get",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "Get",
	RequestType:  (*GetRequest)(nil),
	ResponseType: (*GetResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/container/v1/workloads/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetStatusMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.GetStatus",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "GetStatus",
	RequestType:  (*GetStatusRequest)(nil),
	ResponseType: (*container.Workload)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/container/v1/workloads/:instanceId/status",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetSummaryMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.GetSummary",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "GetSummary",
	RequestType:  (*GetSummaryRequest)(nil),
	ResponseType: (*GetSummaryResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/container/v1/workloads/:instanceId/summary",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.List",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "List",
	RequestType:  (*ListRequest)(nil),
	ResponseType: (*ListResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/container/v1/workloads",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.Update",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "Update",
	RequestType:  (*UpdateRequest)(nil),
	ResponseType: (*container.Workload)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/container/v1/workloads/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateResourceSpecMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.container.workload.UpdateResourceSpec",
		Version: "1.0",
	},
	ServiceName:  "workload.rpc",
	MethodName:   "UpdateResourceSpec",
	RequestType:  (*UpdateResourceSpecRequest)(nil),
	ResponseType: (*container.Workload)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/container/v1/workloads/:instanceId/yaml",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
