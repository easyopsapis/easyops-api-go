// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package instance_graph

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	_ "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
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

// Client is the client API for instance_graph service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	TraverseGraph(ctx context.Context, in *TraverseGraphRequest) (*TraverseGraphResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) TraverseGraph(ctx context.Context, in *TraverseGraphRequest) (*TraverseGraphResponse, error) {
	out := new(TraverseGraphResponse)
	err := c.c.Invoke(ctx, _TraverseGraphContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for instance_graph service.
type Service interface {
	TraverseGraph(context.Context, *TraverseGraphRequest) (*TraverseGraphResponse, error)
}

func _TraverseGraphEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.TraverseGraph(ctx, req.(*TraverseGraphRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_TraverseGraphContract, _TraverseGraphEndpoint(srv))
}

// API Contract
var _TraverseGraphContract = &traverseGraphContract{}

type traverseGraphContract struct{}

func (*traverseGraphContract) ServiceName() string          { return "instance_graph.rpc" }
func (*traverseGraphContract) MethodName() string           { return "TraverseGraph" }
func (*traverseGraphContract) RequestMessage() interface{}  { return new(TraverseGraphRequest) }
func (*traverseGraphContract) ResponseMessage() interface{} { return new(TraverseGraphRequest) }
func (*traverseGraphContract) ContractName() string {
	return "easyops.api.cmdb.instance_graph.TraverseGraph"
}
func (*traverseGraphContract) ContractVersion() string   { return "1.0" }
func (*traverseGraphContract) Pattern() (string, string) { return "POST", "/instance/traverse" }
func (*traverseGraphContract) Body() string              { return "" }