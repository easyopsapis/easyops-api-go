// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package dbservice

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

// Client is the client API for dbservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	BatchUpdateDBServiceOwner(ctx context.Context, in *BatchUpdateDBServiceOwnerRequest) (*types.Empty, error)
	CreateDBService(ctx context.Context, in *CreateDBServiceRequest) (*CreateDBServiceResponse, error)
	DeleteDBService(ctx context.Context, in *DeleteDBServiceRequest) (*types.Empty, error)
	GetDBService(ctx context.Context, in *GetDBServiceRequest) (*GetDBServiceResponse, error)
	ListDBService(ctx context.Context, in *ListDBServiceRequest) (*ListDBServiceResponse, error)
	ListenOrgRegister(ctx context.Context, in *ListenOrgRegisterRequest) (*types.Empty, error)
	UpdateDBService(ctx context.Context, in *UpdateDBServiceRequest) (*UpdateDBServiceResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) BatchUpdateDBServiceOwner(ctx context.Context, in *BatchUpdateDBServiceOwnerRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _BatchUpdateDBServiceOwnerMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateDBService(ctx context.Context, in *CreateDBServiceRequest) (*CreateDBServiceResponse, error) {
	out := new(CreateDBServiceResponse)
	err := c.c.Invoke(ctx, _CreateDBServiceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteDBService(ctx context.Context, in *DeleteDBServiceRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteDBServiceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetDBService(ctx context.Context, in *GetDBServiceRequest) (*GetDBServiceResponse, error) {
	out := new(GetDBServiceResponse)
	err := c.c.Invoke(ctx, _GetDBServiceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListDBService(ctx context.Context, in *ListDBServiceRequest) (*ListDBServiceResponse, error) {
	out := new(ListDBServiceResponse)
	err := c.c.Invoke(ctx, _ListDBServiceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListenOrgRegister(ctx context.Context, in *ListenOrgRegisterRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _ListenOrgRegisterMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDBService(ctx context.Context, in *UpdateDBServiceRequest) (*UpdateDBServiceResponse, error) {
	out := new(UpdateDBServiceResponse)
	err := c.c.Invoke(ctx, _UpdateDBServiceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for dbservice service.
type Service interface {
	BatchUpdateDBServiceOwner(context.Context, *BatchUpdateDBServiceOwnerRequest) (*types.Empty, error)
	CreateDBService(context.Context, *CreateDBServiceRequest) (*CreateDBServiceResponse, error)
	DeleteDBService(context.Context, *DeleteDBServiceRequest) (*types.Empty, error)
	GetDBService(context.Context, *GetDBServiceRequest) (*GetDBServiceResponse, error)
	ListDBService(context.Context, *ListDBServiceRequest) (*ListDBServiceResponse, error)
	ListenOrgRegister(context.Context, *ListenOrgRegisterRequest) (*types.Empty, error)
	UpdateDBService(context.Context, *UpdateDBServiceRequest) (*UpdateDBServiceResponse, error)
}

func _BatchUpdateDBServiceOwnerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BatchUpdateDBServiceOwner(ctx, req.(*BatchUpdateDBServiceOwnerRequest))
	}
}

func _CreateDBServiceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateDBService(ctx, req.(*CreateDBServiceRequest))
	}
}

func _DeleteDBServiceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteDBService(ctx, req.(*DeleteDBServiceRequest))
	}
}

func _GetDBServiceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDBService(ctx, req.(*GetDBServiceRequest))
	}
}

func _ListDBServiceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListDBService(ctx, req.(*ListDBServiceRequest))
	}
}

func _ListenOrgRegisterEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListenOrgRegister(ctx, req.(*ListenOrgRegisterRequest))
	}
}

func _UpdateDBServiceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateDBService(ctx, req.(*UpdateDBServiceRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_BatchUpdateDBServiceOwnerMethodDesc, _BatchUpdateDBServiceOwnerEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateDBServiceMethodDesc, _CreateDBServiceEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteDBServiceMethodDesc, _DeleteDBServiceEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetDBServiceMethodDesc, _GetDBServiceEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListDBServiceMethodDesc, _ListDBServiceEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListenOrgRegisterMethodDesc, _ListenOrgRegisterEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateDBServiceMethodDesc, _UpdateDBServiceEndpoint(srv))
}

// Method Description
var _BatchUpdateDBServiceOwnerMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.BatchUpdateDBServiceOwner",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "BatchUpdateDBServiceOwner",
	RequestType:  (*BatchUpdateDBServiceOwnerRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/v1/dbservices-owners",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _CreateDBServiceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.CreateDBService",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "CreateDBService",
	RequestType:  (*CreateDBServiceRequest)(nil),
	ResponseType: (*CreateDBServiceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/dbservices",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteDBServiceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.DeleteDBService",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "DeleteDBService",
	RequestType:  (*DeleteDBServiceRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/api/v1/dbservices/:dbServiceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetDBServiceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.GetDBService",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "GetDBService",
	RequestType:  (*GetDBServiceRequest)(nil),
	ResponseType: (*GetDBServiceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/dbservices/:dbServiceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListDBServiceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.ListDBService",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "ListDBService",
	RequestType:  (*ListDBServiceRequest)(nil),
	ResponseType: (*ListDBServiceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/dbservices",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListenOrgRegisterMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.ListenOrgRegister",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "ListenOrgRegister",
	RequestType:  (*ListenOrgRegisterRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/org/register",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateDBServiceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.database_delivery.dbservice.UpdateDBService",
		Version: "1.0",
	},
	ServiceName:  "dbservice.rpc",
	MethodName:   "UpdateDBService",
	RequestType:  (*UpdateDBServiceRequest)(nil),
	ResponseType: (*UpdateDBServiceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/api/v1/dbservices/:dbServiceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
