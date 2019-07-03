// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package container

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	_ "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = giraffe_micro.SupportPackageIsVersion3 // please upgrade the giraffe_micro package

// Client is the client API for container service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateContainer(ctx context.Context, in *CreateContainerRequest) (*CreateContainerResponse, error)
	DeleteContainer(ctx context.Context, in *DeleteContainerRequest) (*DeleteContainerResponse, error)
	GetContainer(ctx context.Context, in *GetContainerRequest) (*topology.Container, error)
	ListContainer(ctx context.Context, in *ListContainerRequest) (*ListContainerResponse, error)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest) (*UpdateContainerResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateContainer(ctx context.Context, in *CreateContainerRequest) (*CreateContainerResponse, error) {
	out := new(CreateContainerResponse)
	err := c.c.Invoke(ctx, _CreateContainerContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteContainer(ctx context.Context, in *DeleteContainerRequest) (*DeleteContainerResponse, error) {
	out := new(DeleteContainerResponse)
	err := c.c.Invoke(ctx, _DeleteContainerContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetContainer(ctx context.Context, in *GetContainerRequest) (*topology.Container, error) {
	out := new(topology.Container)
	err := c.c.Invoke(ctx, _GetContainerContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListContainer(ctx context.Context, in *ListContainerRequest) (*ListContainerResponse, error) {
	out := new(ListContainerResponse)
	err := c.c.Invoke(ctx, _ListContainerContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateContainer(ctx context.Context, in *UpdateContainerRequest) (*UpdateContainerResponse, error) {
	out := new(UpdateContainerResponse)
	err := c.c.Invoke(ctx, _UpdateContainerContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for container service.
type Service interface {
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
	DeleteContainer(context.Context, *DeleteContainerRequest) (*DeleteContainerResponse, error)
	GetContainer(context.Context, *GetContainerRequest) (*topology.Container, error)
	ListContainer(context.Context, *ListContainerRequest) (*ListContainerResponse, error)
	UpdateContainer(context.Context, *UpdateContainerRequest) (*UpdateContainerResponse, error)
}

func _CreateContainerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateContainer(ctx, req.(*CreateContainerRequest))
	}
}

func _DeleteContainerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteContainer(ctx, req.(*DeleteContainerRequest))
	}
}

func _GetContainerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetContainer(ctx, req.(*GetContainerRequest))
	}
}

func _ListContainerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListContainer(ctx, req.(*ListContainerRequest))
	}
}

func _UpdateContainerEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateContainer(ctx, req.(*UpdateContainerRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateContainerContract, _CreateContainerEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteContainerContract, _DeleteContainerEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetContainerContract, _GetContainerEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListContainerContract, _ListContainerEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateContainerContract, _UpdateContainerEndpoint(srv))
}

// API Contract
var _CreateContainerContract = &createContainerContract{}

type createContainerContract struct{}

func (*createContainerContract) ServiceName() string          { return "container.rpc" }
func (*createContainerContract) MethodName() string           { return "CreateContainer" }
func (*createContainerContract) RequestMessage() interface{}  { return new(CreateContainerRequest) }
func (*createContainerContract) ResponseMessage() interface{} { return new(CreateContainerRequest) }
func (*createContainerContract) ContractName() string {
	return "easyops.api.topology.container.CreateContainer"
}
func (*createContainerContract) ContractVersion() string   { return "1.0" }
func (*createContainerContract) Pattern() (string, string) { return "POST", "/api/v1/container" }
func (*createContainerContract) Body() string              { return "" }

var _DeleteContainerContract = &deleteContainerContract{}

type deleteContainerContract struct{}

func (*deleteContainerContract) ServiceName() string          { return "container.rpc" }
func (*deleteContainerContract) MethodName() string           { return "DeleteContainer" }
func (*deleteContainerContract) RequestMessage() interface{}  { return new(DeleteContainerRequest) }
func (*deleteContainerContract) ResponseMessage() interface{} { return new(DeleteContainerRequest) }
func (*deleteContainerContract) ContractName() string {
	return "easyops.api.topology.container.DeleteContainer"
}
func (*deleteContainerContract) ContractVersion() string   { return "1.0" }
func (*deleteContainerContract) Pattern() (string, string) { return "DELETE", "/api/v1/container/:id" }
func (*deleteContainerContract) Body() string              { return "" }

var _GetContainerContract = &getContainerContract{}

type getContainerContract struct{}

func (*getContainerContract) ServiceName() string          { return "container.rpc" }
func (*getContainerContract) MethodName() string           { return "GetContainer" }
func (*getContainerContract) RequestMessage() interface{}  { return new(GetContainerRequest) }
func (*getContainerContract) ResponseMessage() interface{} { return new(GetContainerRequest) }
func (*getContainerContract) ContractName() string {
	return "easyops.api.topology.container.GetContainer"
}
func (*getContainerContract) ContractVersion() string   { return "1.0" }
func (*getContainerContract) Pattern() (string, string) { return "GET", "/api/v1/container/:id" }
func (*getContainerContract) Body() string              { return "" }

var _ListContainerContract = &listContainerContract{}

type listContainerContract struct{}

func (*listContainerContract) ServiceName() string          { return "container.rpc" }
func (*listContainerContract) MethodName() string           { return "ListContainer" }
func (*listContainerContract) RequestMessage() interface{}  { return new(ListContainerRequest) }
func (*listContainerContract) ResponseMessage() interface{} { return new(ListContainerRequest) }
func (*listContainerContract) ContractName() string {
	return "easyops.api.topology.container.ListContainer"
}
func (*listContainerContract) ContractVersion() string   { return "1.0" }
func (*listContainerContract) Pattern() (string, string) { return "GET", "/api/v1/container" }
func (*listContainerContract) Body() string              { return "" }

var _UpdateContainerContract = &updateContainerContract{}

type updateContainerContract struct{}

func (*updateContainerContract) ServiceName() string          { return "container.rpc" }
func (*updateContainerContract) MethodName() string           { return "UpdateContainer" }
func (*updateContainerContract) RequestMessage() interface{}  { return new(UpdateContainerRequest) }
func (*updateContainerContract) ResponseMessage() interface{} { return new(UpdateContainerRequest) }
func (*updateContainerContract) ContractName() string {
	return "easyops.api.topology.container.UpdateContainer"
}
func (*updateContainerContract) ContractVersion() string   { return "1.0" }
func (*updateContainerContract) Pattern() (string, string) { return "PUT", "/api/v1/container/:id" }
func (*updateContainerContract) Body() string              { return "" }
