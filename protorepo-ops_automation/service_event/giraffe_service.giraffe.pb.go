// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package service_event

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

// Client is the client API for service_event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateServiceEvent(ctx context.Context, in *CreateServiceEventRequest) (*CreateServiceEventResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateServiceEvent(ctx context.Context, in *CreateServiceEventRequest) (*CreateServiceEventResponse, error) {
	out := new(CreateServiceEventResponse)
	err := c.c.Invoke(ctx, _CreateServiceEventContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for service_event service.
type Service interface {
	CreateServiceEvent(context.Context, *CreateServiceEventRequest) (*CreateServiceEventResponse, error)
}

func _CreateServiceEventEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateServiceEvent(ctx, req.(*CreateServiceEventRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateServiceEventContract, _CreateServiceEventEndpoint(srv))
}

// API Contract
var _CreateServiceEventContract = &createServiceEventContract{}

type createServiceEventContract struct{}

func (*createServiceEventContract) ServiceName() string         { return "service_event.rpc" }
func (*createServiceEventContract) MethodName() string          { return "CreateServiceEvent" }
func (*createServiceEventContract) RequestMessage() interface{} { return new(CreateServiceEventRequest) }
func (*createServiceEventContract) ResponseMessage() interface{} {
	return new(CreateServiceEventRequest)
}
func (*createServiceEventContract) ContractName() string {
	return "easyops.api.ops_automation.service_event.CreateServiceEvent"
}
func (*createServiceEventContract) ContractVersion() string   { return "1.0" }
func (*createServiceEventContract) Pattern() (string, string) { return "POST", "/service/event" }
func (*createServiceEventContract) Body() string              { return "" }