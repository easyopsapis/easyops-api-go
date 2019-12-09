// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package toolkit

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	cmdb_extend "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
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

// Client is the client API for toolkit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateToolkit(ctx context.Context, in *CreateToolkitRequest) (*cmdb_extend.Toolkit, error)
	DeleteToolkit(ctx context.Context, in *DeleteToolkitRequest) (*types.Empty, error)
	GetToolkit(ctx context.Context, in *GetToolkitRequest) (*cmdb_extend.Toolkit, error)
	GetToolkitStatus(ctx context.Context, in *GetToolkitStatusRequest) (*GetToolkitStatusResponse, error)
	ListToolkit(ctx context.Context, in *ListToolkitRequest) (*ListToolkitResponse, error)
	UpdateToolkit(ctx context.Context, in *UpdateToolkitRequest) (*cmdb_extend.Toolkit, error)
	UpdateToolkitStatus(ctx context.Context, in *UpdateToolkitStatusRequest) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateToolkit(ctx context.Context, in *CreateToolkitRequest) (*cmdb_extend.Toolkit, error) {
	out := new(cmdb_extend.Toolkit)
	err := c.c.Invoke(ctx, _CreateToolkitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteToolkit(ctx context.Context, in *DeleteToolkitRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteToolkitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetToolkit(ctx context.Context, in *GetToolkitRequest) (*cmdb_extend.Toolkit, error) {
	out := new(cmdb_extend.Toolkit)
	err := c.c.Invoke(ctx, _GetToolkitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetToolkitStatus(ctx context.Context, in *GetToolkitStatusRequest) (*GetToolkitStatusResponse, error) {
	out := new(GetToolkitStatusResponse)
	err := c.c.Invoke(ctx, _GetToolkitStatusMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListToolkit(ctx context.Context, in *ListToolkitRequest) (*ListToolkitResponse, error) {
	out := new(ListToolkitResponse)
	err := c.c.Invoke(ctx, _ListToolkitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateToolkit(ctx context.Context, in *UpdateToolkitRequest) (*cmdb_extend.Toolkit, error) {
	out := new(cmdb_extend.Toolkit)
	err := c.c.Invoke(ctx, _UpdateToolkitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateToolkitStatus(ctx context.Context, in *UpdateToolkitStatusRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _UpdateToolkitStatusMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for toolkit service.
type Service interface {
	CreateToolkit(context.Context, *CreateToolkitRequest) (*cmdb_extend.Toolkit, error)
	DeleteToolkit(context.Context, *DeleteToolkitRequest) (*types.Empty, error)
	GetToolkit(context.Context, *GetToolkitRequest) (*cmdb_extend.Toolkit, error)
	GetToolkitStatus(context.Context, *GetToolkitStatusRequest) (*GetToolkitStatusResponse, error)
	ListToolkit(context.Context, *ListToolkitRequest) (*ListToolkitResponse, error)
	UpdateToolkit(context.Context, *UpdateToolkitRequest) (*cmdb_extend.Toolkit, error)
	UpdateToolkitStatus(context.Context, *UpdateToolkitStatusRequest) (*types.Empty, error)
}

func _CreateToolkitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateToolkit(ctx, req.(*CreateToolkitRequest))
	}
}

func _DeleteToolkitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteToolkit(ctx, req.(*DeleteToolkitRequest))
	}
}

func _GetToolkitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetToolkit(ctx, req.(*GetToolkitRequest))
	}
}

func _GetToolkitStatusEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetToolkitStatus(ctx, req.(*GetToolkitStatusRequest))
	}
}

func _ListToolkitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListToolkit(ctx, req.(*ListToolkitRequest))
	}
}

func _UpdateToolkitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateToolkit(ctx, req.(*UpdateToolkitRequest))
	}
}

func _UpdateToolkitStatusEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateToolkitStatus(ctx, req.(*UpdateToolkitStatusRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateToolkitMethodDesc, _CreateToolkitEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteToolkitMethodDesc, _DeleteToolkitEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetToolkitMethodDesc, _GetToolkitEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetToolkitStatusMethodDesc, _GetToolkitStatusEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListToolkitMethodDesc, _ListToolkitEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateToolkitMethodDesc, _UpdateToolkitEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateToolkitStatusMethodDesc, _UpdateToolkitStatusEndpoint(srv))
}

// Method Description
var _CreateToolkitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.CreateToolkit",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "CreateToolkit",
	RequestType:  (*CreateToolkitRequest)(nil),
	ResponseType: (*cmdb_extend.Toolkit)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/toolkit/:objectId",
		},
		Body:         "toolkit",
		ResponseBody: "data",
	},
}

var _DeleteToolkitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.DeleteToolkit",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "DeleteToolkit",
	RequestType:  (*DeleteToolkitRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/toolkit/entity/:objectId/:toolkitId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetToolkitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.GetToolkit",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "GetToolkit",
	RequestType:  (*GetToolkitRequest)(nil),
	ResponseType: (*cmdb_extend.Toolkit)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/toolkit/entity/:objectId/:toolkitId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetToolkitStatusMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.GetToolkitStatus",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "GetToolkitStatus",
	RequestType:  (*GetToolkitStatusRequest)(nil),
	ResponseType: (*GetToolkitStatusResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/toolkit/setting/:objectId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListToolkitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.ListToolkit",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "ListToolkit",
	RequestType:  (*ListToolkitRequest)(nil),
	ResponseType: (*ListToolkitResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/toolkit/tools/:objectId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateToolkitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.UpdateToolkit",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "UpdateToolkit",
	RequestType:  (*UpdateToolkitRequest)(nil),
	ResponseType: (*cmdb_extend.Toolkit)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/toolkit/entity/:objectId/:toolkitId",
		},
		Body:         "toolkit",
		ResponseBody: "data",
	},
}

var _UpdateToolkitStatusMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.toolkit.UpdateToolkitStatus",
		Version: "1.0",
	},
	ServiceName:  "toolkit.rpc",
	MethodName:   "UpdateToolkitStatus",
	RequestType:  (*UpdateToolkitStatusRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/toolkit/setting/:objectId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
