// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package strategy

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	_ "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	easy_flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_flow"
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

// Client is the client API for strategy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *CreateRequest) (*easy_flow.DeployStrategy, error)
	DeleteStrategy(ctx context.Context, in *DeleteStrategyRequest) (*types.Empty, error)
	Get(ctx context.Context, in *GetRequest) (*easy_flow.DeployStrategy, error)
	List(ctx context.Context, in *ListRequest) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateRequest) (*easy_flow.DeployStrategy, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *CreateRequest) (*easy_flow.DeployStrategy, error) {
	out := new(easy_flow.DeployStrategy)
	err := c.c.Invoke(ctx, _CreateContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteStrategy(ctx context.Context, in *DeleteStrategyRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteStrategyContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Get(ctx context.Context, in *GetRequest) (*easy_flow.DeployStrategy, error) {
	out := new(easy_flow.DeployStrategy)
	err := c.c.Invoke(ctx, _GetContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.c.Invoke(ctx, _ListContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Update(ctx context.Context, in *UpdateRequest) (*easy_flow.DeployStrategy, error) {
	out := new(easy_flow.DeployStrategy)
	err := c.c.Invoke(ctx, _UpdateContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for strategy service.
type Service interface {
	Create(context.Context, *CreateRequest) (*easy_flow.DeployStrategy, error)
	DeleteStrategy(context.Context, *DeleteStrategyRequest) (*types.Empty, error)
	Get(context.Context, *GetRequest) (*easy_flow.DeployStrategy, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Update(context.Context, *UpdateRequest) (*easy_flow.DeployStrategy, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
}

func _DeleteStrategyEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteStrategy(ctx, req.(*DeleteStrategyRequest))
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
	s.RegisterUnaryEndpoint(_CreateContract, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteStrategyContract, _DeleteStrategyEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetContract, _GetEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListContract, _ListEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateContract, _UpdateEndpoint(srv))
}

// API Contract
var _CreateContract = &createContract{}

type createContract struct{}

func (*createContract) ServiceName() string          { return "strategy.rpc" }
func (*createContract) MethodName() string           { return "Create" }
func (*createContract) RequestMessage() interface{}  { return new(CreateRequest) }
func (*createContract) ResponseMessage() interface{} { return new(CreateRequest) }
func (*createContract) ContractName() string         { return "easyops.api.easy_flow.strategy.Create" }
func (*createContract) ContractVersion() string      { return "1.0" }
func (*createContract) Pattern() (string, string)    { return "POST", "/deployStrategy" }
func (*createContract) Body() string                 { return "" }

var _DeleteStrategyContract = &deleteStrategyContract{}

type deleteStrategyContract struct{}

func (*deleteStrategyContract) ServiceName() string          { return "strategy.rpc" }
func (*deleteStrategyContract) MethodName() string           { return "DeleteStrategy" }
func (*deleteStrategyContract) RequestMessage() interface{}  { return new(DeleteStrategyRequest) }
func (*deleteStrategyContract) ResponseMessage() interface{} { return new(DeleteStrategyRequest) }
func (*deleteStrategyContract) ContractName() string {
	return "easyops.api.easy_flow.strategy.DeleteStrategy"
}
func (*deleteStrategyContract) ContractVersion() string { return "1.0" }
func (*deleteStrategyContract) Pattern() (string, string) {
	return "DELETE", "/deployStrategy/:strategyID"
}
func (*deleteStrategyContract) Body() string { return "" }

var _GetContract = &getContract{}

type getContract struct{}

func (*getContract) ServiceName() string          { return "strategy.rpc" }
func (*getContract) MethodName() string           { return "Get" }
func (*getContract) RequestMessage() interface{}  { return new(GetRequest) }
func (*getContract) ResponseMessage() interface{} { return new(GetRequest) }
func (*getContract) ContractName() string         { return "easyops.api.easy_flow.strategy.Get" }
func (*getContract) ContractVersion() string      { return "1.0" }
func (*getContract) Pattern() (string, string)    { return "GET", "/deployStrategy/:strategyID" }
func (*getContract) Body() string                 { return "" }

var _ListContract = &listContract{}

type listContract struct{}

func (*listContract) ServiceName() string          { return "strategy.rpc" }
func (*listContract) MethodName() string           { return "List" }
func (*listContract) RequestMessage() interface{}  { return new(ListRequest) }
func (*listContract) ResponseMessage() interface{} { return new(ListRequest) }
func (*listContract) ContractName() string         { return "easyops.api.easy_flow.strategy.List" }
func (*listContract) ContractVersion() string      { return "1.0" }
func (*listContract) Pattern() (string, string)    { return "GET", "/deployStrategy" }
func (*listContract) Body() string                 { return "" }

var _UpdateContract = &updateContract{}

type updateContract struct{}

func (*updateContract) ServiceName() string          { return "strategy.rpc" }
func (*updateContract) MethodName() string           { return "Update" }
func (*updateContract) RequestMessage() interface{}  { return new(UpdateRequest) }
func (*updateContract) ResponseMessage() interface{} { return new(UpdateRequest) }
func (*updateContract) ContractName() string         { return "easyops.api.easy_flow.strategy.Update" }
func (*updateContract) ContractVersion() string      { return "1.0" }
func (*updateContract) Pattern() (string, string)    { return "PUT", "/deployStrategy/:strategyID" }
func (*updateContract) Body() string                 { return "" }