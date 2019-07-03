// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package installed_micro_app

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	_ "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	micro_app "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
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

// Client is the client API for installed_micro_app service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *micro_app.InstalledMicroApp) (*micro_app.InstalledMicroApp, error)
	DeleteMicroApp(ctx context.Context, in *DeleteMicroAppRequest) (*types.Empty, error)
	GetInstalledMicroApp(ctx context.Context, in *GetInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error)
	ListMicroApp(ctx context.Context, in *ListMicroAppRequest) (*ListMicroAppResponse, error)
	UpdateInstalledMicroApp(ctx context.Context, in *UpdateInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *micro_app.InstalledMicroApp) (*micro_app.InstalledMicroApp, error) {
	out := new(micro_app.InstalledMicroApp)
	err := c.c.Invoke(ctx, _CreateContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteMicroApp(ctx context.Context, in *DeleteMicroAppRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteMicroAppContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetInstalledMicroApp(ctx context.Context, in *GetInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error) {
	out := new(micro_app.InstalledMicroApp)
	err := c.c.Invoke(ctx, _GetInstalledMicroAppContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListMicroApp(ctx context.Context, in *ListMicroAppRequest) (*ListMicroAppResponse, error) {
	out := new(ListMicroAppResponse)
	err := c.c.Invoke(ctx, _ListMicroAppContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateInstalledMicroApp(ctx context.Context, in *UpdateInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error) {
	out := new(micro_app.InstalledMicroApp)
	err := c.c.Invoke(ctx, _UpdateInstalledMicroAppContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for installed_micro_app service.
type Service interface {
	Create(context.Context, *micro_app.InstalledMicroApp) (*micro_app.InstalledMicroApp, error)
	DeleteMicroApp(context.Context, *DeleteMicroAppRequest) (*types.Empty, error)
	GetInstalledMicroApp(context.Context, *GetInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error)
	ListMicroApp(context.Context, *ListMicroAppRequest) (*ListMicroAppResponse, error)
	UpdateInstalledMicroApp(context.Context, *UpdateInstalledMicroAppRequest) (*micro_app.InstalledMicroApp, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*micro_app.InstalledMicroApp))
	}
}

func _DeleteMicroAppEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteMicroApp(ctx, req.(*DeleteMicroAppRequest))
	}
}

func _GetInstalledMicroAppEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetInstalledMicroApp(ctx, req.(*GetInstalledMicroAppRequest))
	}
}

func _ListMicroAppEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListMicroApp(ctx, req.(*ListMicroAppRequest))
	}
}

func _UpdateInstalledMicroAppEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateInstalledMicroApp(ctx, req.(*UpdateInstalledMicroAppRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateContract, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteMicroAppContract, _DeleteMicroAppEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetInstalledMicroAppContract, _GetInstalledMicroAppEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListMicroAppContract, _ListMicroAppEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateInstalledMicroAppContract, _UpdateInstalledMicroAppEndpoint(srv))
}

// API Contract
var _CreateContract = &createContract{}

type createContract struct{}

func (*createContract) ServiceName() string          { return "installed_micro_app.rpc" }
func (*createContract) MethodName() string           { return "Create" }
func (*createContract) RequestMessage() interface{}  { return new(micro_app.InstalledMicroApp) }
func (*createContract) ResponseMessage() interface{} { return new(micro_app.InstalledMicroApp) }
func (*createContract) ContractName() string {
	return "easyops.api.micro_app.installed_micro_app.Create"
}
func (*createContract) ContractVersion() string { return "1.0" }
func (*createContract) Pattern() (string, string) {
	return "POST", "/api/micro_app/v1/installed_micro_app"
}
func (*createContract) Body() string { return "" }

var _DeleteMicroAppContract = &deleteMicroAppContract{}

type deleteMicroAppContract struct{}

func (*deleteMicroAppContract) ServiceName() string          { return "installed_micro_app.rpc" }
func (*deleteMicroAppContract) MethodName() string           { return "DeleteMicroApp" }
func (*deleteMicroAppContract) RequestMessage() interface{}  { return new(DeleteMicroAppRequest) }
func (*deleteMicroAppContract) ResponseMessage() interface{} { return new(DeleteMicroAppRequest) }
func (*deleteMicroAppContract) ContractName() string {
	return "easyops.api.micro_app.installed_micro_app.DeleteMicroApp"
}
func (*deleteMicroAppContract) ContractVersion() string { return "1.0" }
func (*deleteMicroAppContract) Pattern() (string, string) {
	return "DELETE", "/api/micro_app/v1/installed_micro_app/:app_id"
}
func (*deleteMicroAppContract) Body() string { return "" }

var _GetInstalledMicroAppContract = &getInstalledMicroAppContract{}

type getInstalledMicroAppContract struct{}

func (*getInstalledMicroAppContract) ServiceName() string { return "installed_micro_app.rpc" }
func (*getInstalledMicroAppContract) MethodName() string  { return "GetInstalledMicroApp" }
func (*getInstalledMicroAppContract) RequestMessage() interface{} {
	return new(GetInstalledMicroAppRequest)
}
func (*getInstalledMicroAppContract) ResponseMessage() interface{} {
	return new(GetInstalledMicroAppRequest)
}
func (*getInstalledMicroAppContract) ContractName() string {
	return "easyops.api.micro_app.installed_micro_app.GetInstalledMicroApp"
}
func (*getInstalledMicroAppContract) ContractVersion() string { return "1.0" }
func (*getInstalledMicroAppContract) Pattern() (string, string) {
	return "GET", "/api/micro_app/v1/installed_micro_app/:app_id"
}
func (*getInstalledMicroAppContract) Body() string { return "" }

var _ListMicroAppContract = &listMicroAppContract{}

type listMicroAppContract struct{}

func (*listMicroAppContract) ServiceName() string          { return "installed_micro_app.rpc" }
func (*listMicroAppContract) MethodName() string           { return "ListMicroApp" }
func (*listMicroAppContract) RequestMessage() interface{}  { return new(ListMicroAppRequest) }
func (*listMicroAppContract) ResponseMessage() interface{} { return new(ListMicroAppRequest) }
func (*listMicroAppContract) ContractName() string {
	return "easyops.api.micro_app.installed_micro_app.ListMicroApp"
}
func (*listMicroAppContract) ContractVersion() string { return "1.0" }
func (*listMicroAppContract) Pattern() (string, string) {
	return "GET", "/api/micro_app/v1/installed_micro_app"
}
func (*listMicroAppContract) Body() string { return "" }

var _UpdateInstalledMicroAppContract = &updateInstalledMicroAppContract{}

type updateInstalledMicroAppContract struct{}

func (*updateInstalledMicroAppContract) ServiceName() string { return "installed_micro_app.rpc" }
func (*updateInstalledMicroAppContract) MethodName() string  { return "UpdateInstalledMicroApp" }
func (*updateInstalledMicroAppContract) RequestMessage() interface{} {
	return new(UpdateInstalledMicroAppRequest)
}
func (*updateInstalledMicroAppContract) ResponseMessage() interface{} {
	return new(UpdateInstalledMicroAppRequest)
}
func (*updateInstalledMicroAppContract) ContractName() string {
	return "easyops.api.micro_app.installed_micro_app.UpdateInstalledMicroApp"
}
func (*updateInstalledMicroAppContract) ContractVersion() string { return "1.0" }
func (*updateInstalledMicroAppContract) Pattern() (string, string) {
	return "PUT", "/api/micro_app/v1/installed_micro_app/:app_id"
}
func (*updateInstalledMicroAppContract) Body() string { return "" }