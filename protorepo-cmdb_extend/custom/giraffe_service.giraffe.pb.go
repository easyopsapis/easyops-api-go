// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package custom

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

// Client is the client API for custom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	GetAllSystemTree(ctx context.Context, in *types.Empty) (*GetAllSystemTreeResponse, error)
	GetComponentArchitecture(ctx context.Context, in *GetComponentArchitectureRequest) (*cmdb_extend.AppDependency, error)
	GetIdcAllRackUnit(ctx context.Context, in *GetIdcAllRackUnitRequest) (*types.Struct, error)
	GetIdcrackUnit(ctx context.Context, in *GetIdcrackUnitRequest) (*cmdb_extend.IdcrackUnitInfo, error)
	GetOneSubSystemTree(ctx context.Context, in *GetOneSubSystemTreeRequest) (*cmdb_extend.SubsystemDependency, error)
	GetOneSystemTree(ctx context.Context, in *GetOneSystemTreeRequest) (*cmdb_extend.SystemDependency, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) GetAllSystemTree(ctx context.Context, in *types.Empty) (*GetAllSystemTreeResponse, error) {
	out := new(GetAllSystemTreeResponse)
	err := c.c.Invoke(ctx, _GetAllSystemTreeMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetComponentArchitecture(ctx context.Context, in *GetComponentArchitectureRequest) (*cmdb_extend.AppDependency, error) {
	out := new(cmdb_extend.AppDependency)
	err := c.c.Invoke(ctx, _GetComponentArchitectureMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetIdcAllRackUnit(ctx context.Context, in *GetIdcAllRackUnitRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _GetIdcAllRackUnitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetIdcrackUnit(ctx context.Context, in *GetIdcrackUnitRequest) (*cmdb_extend.IdcrackUnitInfo, error) {
	out := new(cmdb_extend.IdcrackUnitInfo)
	err := c.c.Invoke(ctx, _GetIdcrackUnitMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetOneSubSystemTree(ctx context.Context, in *GetOneSubSystemTreeRequest) (*cmdb_extend.SubsystemDependency, error) {
	out := new(cmdb_extend.SubsystemDependency)
	err := c.c.Invoke(ctx, _GetOneSubSystemTreeMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetOneSystemTree(ctx context.Context, in *GetOneSystemTreeRequest) (*cmdb_extend.SystemDependency, error) {
	out := new(cmdb_extend.SystemDependency)
	err := c.c.Invoke(ctx, _GetOneSystemTreeMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for custom service.
type Service interface {
	GetAllSystemTree(context.Context, *types.Empty) (*GetAllSystemTreeResponse, error)
	GetComponentArchitecture(context.Context, *GetComponentArchitectureRequest) (*cmdb_extend.AppDependency, error)
	GetIdcAllRackUnit(context.Context, *GetIdcAllRackUnitRequest) (*types.Struct, error)
	GetIdcrackUnit(context.Context, *GetIdcrackUnitRequest) (*cmdb_extend.IdcrackUnitInfo, error)
	GetOneSubSystemTree(context.Context, *GetOneSubSystemTreeRequest) (*cmdb_extend.SubsystemDependency, error)
	GetOneSystemTree(context.Context, *GetOneSystemTreeRequest) (*cmdb_extend.SystemDependency, error)
}

func _GetAllSystemTreeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAllSystemTree(ctx, req.(*types.Empty))
	}
}

func _GetComponentArchitectureEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetComponentArchitecture(ctx, req.(*GetComponentArchitectureRequest))
	}
}

func _GetIdcAllRackUnitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetIdcAllRackUnit(ctx, req.(*GetIdcAllRackUnitRequest))
	}
}

func _GetIdcrackUnitEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetIdcrackUnit(ctx, req.(*GetIdcrackUnitRequest))
	}
}

func _GetOneSubSystemTreeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetOneSubSystemTree(ctx, req.(*GetOneSubSystemTreeRequest))
	}
}

func _GetOneSystemTreeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetOneSystemTree(ctx, req.(*GetOneSystemTreeRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_GetAllSystemTreeMethodDesc, _GetAllSystemTreeEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetComponentArchitectureMethodDesc, _GetComponentArchitectureEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetIdcAllRackUnitMethodDesc, _GetIdcAllRackUnitEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetIdcrackUnitMethodDesc, _GetIdcrackUnitEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetOneSubSystemTreeMethodDesc, _GetOneSubSystemTreeEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetOneSystemTreeMethodDesc, _GetOneSystemTreeEndpoint(srv))
}

// Method Description
var _GetAllSystemTreeMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetAllSystemTree",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetAllSystemTree",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*GetAllSystemTreeResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/application-tree",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetComponentArchitectureMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetComponentArchitecture",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetComponentArchitecture",
	RequestType:  (*GetComponentArchitectureRequest)(nil),
	ResponseType: (*cmdb_extend.AppDependency)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/component-architecture/:abbreviation",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetIdcAllRackUnitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetIdcAllRackUnit",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetIdcAllRackUnit",
	RequestType:  (*GetIdcAllRackUnitRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/_IDC/instance/:instance_id/unit",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetIdcrackUnitMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetIdcrackUnit",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetIdcrackUnit",
	RequestType:  (*GetIdcrackUnitRequest)(nil),
	ResponseType: (*cmdb_extend.IdcrackUnitInfo)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/_IDCRACK/instance/:instance_id/unit",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetOneSubSystemTreeMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetOneSubSystemTree",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetOneSubSystemTree",
	RequestType:  (*GetOneSubSystemTreeRequest)(nil),
	ResponseType: (*cmdb_extend.SubsystemDependency)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/subsystem-architecture/:abbreviation",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetOneSystemTreeMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb_extend.custom.GetOneSystemTree",
		Version: "1.0",
	},
	ServiceName:  "custom.rpc",
	MethodName:   "GetOneSystemTree",
	RequestType:  (*GetOneSystemTreeRequest)(nil),
	ResponseType: (*cmdb_extend.SystemDependency)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/system-architecture/:abbreviation",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
