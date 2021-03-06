// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package initialization

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

// Client is the client API for initialization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	InitDatabase(ctx context.Context, in *InitDatabaseRequest) (*types.Empty, error)
	InitObjects(ctx context.Context, in *InitObjectsRequest) (*types.Empty, error)
	InitRelations(ctx context.Context, in *InitRelationsRequest) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) InitDatabase(ctx context.Context, in *InitDatabaseRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _InitDatabaseMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) InitObjects(ctx context.Context, in *InitObjectsRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _InitObjectsMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) InitRelations(ctx context.Context, in *InitRelationsRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _InitRelationsMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for initialization service.
type Service interface {
	InitDatabase(context.Context, *InitDatabaseRequest) (*types.Empty, error)
	InitObjects(context.Context, *InitObjectsRequest) (*types.Empty, error)
	InitRelations(context.Context, *InitRelationsRequest) (*types.Empty, error)
}

func _InitDatabaseEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.InitDatabase(ctx, req.(*InitDatabaseRequest))
	}
}

func _InitObjectsEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.InitObjects(ctx, req.(*InitObjectsRequest))
	}
}

func _InitRelationsEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.InitRelations(ctx, req.(*InitRelationsRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_InitDatabaseMethodDesc, _InitDatabaseEndpoint(srv))
	s.RegisterUnaryEndpoint(_InitObjectsMethodDesc, _InitObjectsEndpoint(srv))
	s.RegisterUnaryEndpoint(_InitRelationsMethodDesc, _InitRelationsEndpoint(srv))
}

// Method Description
var _InitDatabaseMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.initialization.InitDatabase",
		Version: "1.0",
	},
	ServiceName:  "initialization.rpc",
	MethodName:   "InitDatabase",
	RequestType:  (*InitDatabaseRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/initialization/database",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _InitObjectsMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.initialization.InitObjects",
		Version: "1.0",
	},
	ServiceName:  "initialization.rpc",
	MethodName:   "InitObjects",
	RequestType:  (*InitObjectsRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/initialization/objects",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _InitRelationsMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.initialization.InitRelations",
		Version: "1.0",
	},
	ServiceName:  "initialization.rpc",
	MethodName:   "InitRelations",
	RequestType:  (*InitRelationsRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/initialization/relations",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
