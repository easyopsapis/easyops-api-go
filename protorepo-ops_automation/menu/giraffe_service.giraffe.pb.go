// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package menu

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	_ "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	ops_automation "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/ops_automation"
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

// Client is the client API for menu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	CreateMenu(ctx context.Context, in *ops_automation.Menu) (*CreateMenuResponse, error)
	CreateMenuTree(ctx context.Context, in *CreateMenuTreeRequest) (*CreateMenuTreeResponse, error)
	DeleteMenu(ctx context.Context, in *DeleteMenuRequest) (*DeleteMenuResponse, error)
	GetMenu(ctx context.Context, in *GetMenuRequest) (*GetMenuResponse, error)
	GetMenuTree(ctx context.Context, in *types.Empty) (*types.Struct, error)
	ListMenus(ctx context.Context, in *ListMenusRequest) (*ListMenusResponse, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest) (*UpdateMenuResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) CreateMenu(ctx context.Context, in *ops_automation.Menu) (*CreateMenuResponse, error) {
	out := new(CreateMenuResponse)
	err := c.c.Invoke(ctx, _CreateMenuContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateMenuTree(ctx context.Context, in *CreateMenuTreeRequest) (*CreateMenuTreeResponse, error) {
	out := new(CreateMenuTreeResponse)
	err := c.c.Invoke(ctx, _CreateMenuTreeContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteMenu(ctx context.Context, in *DeleteMenuRequest) (*DeleteMenuResponse, error) {
	out := new(DeleteMenuResponse)
	err := c.c.Invoke(ctx, _DeleteMenuContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetMenu(ctx context.Context, in *GetMenuRequest) (*GetMenuResponse, error) {
	out := new(GetMenuResponse)
	err := c.c.Invoke(ctx, _GetMenuContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetMenuTree(ctx context.Context, in *types.Empty) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _GetMenuTreeContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListMenus(ctx context.Context, in *ListMenusRequest) (*ListMenusResponse, error) {
	out := new(ListMenusResponse)
	err := c.c.Invoke(ctx, _ListMenusContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateMenu(ctx context.Context, in *UpdateMenuRequest) (*UpdateMenuResponse, error) {
	out := new(UpdateMenuResponse)
	err := c.c.Invoke(ctx, _UpdateMenuContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for menu service.
type Service interface {
	CreateMenu(context.Context, *ops_automation.Menu) (*CreateMenuResponse, error)
	CreateMenuTree(context.Context, *CreateMenuTreeRequest) (*CreateMenuTreeResponse, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuResponse, error)
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	GetMenuTree(context.Context, *types.Empty) (*types.Struct, error)
	ListMenus(context.Context, *ListMenusRequest) (*ListMenusResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
}

func _CreateMenuEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateMenu(ctx, req.(*ops_automation.Menu))
	}
}

func _CreateMenuTreeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateMenuTree(ctx, req.(*CreateMenuTreeRequest))
	}
}

func _DeleteMenuEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteMenu(ctx, req.(*DeleteMenuRequest))
	}
}

func _GetMenuEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMenu(ctx, req.(*GetMenuRequest))
	}
}

func _GetMenuTreeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMenuTree(ctx, req.(*types.Empty))
	}
}

func _ListMenusEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListMenus(ctx, req.(*ListMenusRequest))
	}
}

func _UpdateMenuEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMenuContract, _CreateMenuEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateMenuTreeContract, _CreateMenuTreeEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteMenuContract, _DeleteMenuEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetMenuContract, _GetMenuEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetMenuTreeContract, _GetMenuTreeEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListMenusContract, _ListMenusEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateMenuContract, _UpdateMenuEndpoint(srv))
}

// API Contract
var _CreateMenuContract = &createMenuContract{}

type createMenuContract struct{}

func (*createMenuContract) ServiceName() string          { return "menu.rpc" }
func (*createMenuContract) MethodName() string           { return "CreateMenu" }
func (*createMenuContract) RequestMessage() interface{}  { return new(ops_automation.Menu) }
func (*createMenuContract) ResponseMessage() interface{} { return new(ops_automation.Menu) }
func (*createMenuContract) ContractName() string         { return "easyops.api.ops_automation.menu.CreateMenu" }
func (*createMenuContract) ContractVersion() string      { return "1.0" }
func (*createMenuContract) Pattern() (string, string)    { return "POST", "/menus" }
func (*createMenuContract) Body() string                 { return "" }

var _CreateMenuTreeContract = &createMenuTreeContract{}

type createMenuTreeContract struct{}

func (*createMenuTreeContract) ServiceName() string          { return "menu.rpc" }
func (*createMenuTreeContract) MethodName() string           { return "CreateMenuTree" }
func (*createMenuTreeContract) RequestMessage() interface{}  { return new(CreateMenuTreeRequest) }
func (*createMenuTreeContract) ResponseMessage() interface{} { return new(CreateMenuTreeRequest) }
func (*createMenuTreeContract) ContractName() string {
	return "easyops.api.ops_automation.menu.CreateMenuTree"
}
func (*createMenuTreeContract) ContractVersion() string   { return "1.0" }
func (*createMenuTreeContract) Pattern() (string, string) { return "POST", "/menuTree" }
func (*createMenuTreeContract) Body() string              { return "" }

var _DeleteMenuContract = &deleteMenuContract{}

type deleteMenuContract struct{}

func (*deleteMenuContract) ServiceName() string          { return "menu.rpc" }
func (*deleteMenuContract) MethodName() string           { return "DeleteMenu" }
func (*deleteMenuContract) RequestMessage() interface{}  { return new(DeleteMenuRequest) }
func (*deleteMenuContract) ResponseMessage() interface{} { return new(DeleteMenuRequest) }
func (*deleteMenuContract) ContractName() string         { return "easyops.api.ops_automation.menu.DeleteMenu" }
func (*deleteMenuContract) ContractVersion() string      { return "1.0" }
func (*deleteMenuContract) Pattern() (string, string)    { return "DELETE", "/menus/:menusId" }
func (*deleteMenuContract) Body() string                 { return "" }

var _GetMenuContract = &getMenuContract{}

type getMenuContract struct{}

func (*getMenuContract) ServiceName() string          { return "menu.rpc" }
func (*getMenuContract) MethodName() string           { return "GetMenu" }
func (*getMenuContract) RequestMessage() interface{}  { return new(GetMenuRequest) }
func (*getMenuContract) ResponseMessage() interface{} { return new(GetMenuRequest) }
func (*getMenuContract) ContractName() string         { return "easyops.api.ops_automation.menu.GetMenu" }
func (*getMenuContract) ContractVersion() string      { return "1.0" }
func (*getMenuContract) Pattern() (string, string)    { return "GET", "/menus/:menusId" }
func (*getMenuContract) Body() string                 { return "" }

var _GetMenuTreeContract = &getMenuTreeContract{}

type getMenuTreeContract struct{}

func (*getMenuTreeContract) ServiceName() string          { return "menu.rpc" }
func (*getMenuTreeContract) MethodName() string           { return "GetMenuTree" }
func (*getMenuTreeContract) RequestMessage() interface{}  { return new(types.Empty) }
func (*getMenuTreeContract) ResponseMessage() interface{} { return new(types.Empty) }
func (*getMenuTreeContract) ContractName() string {
	return "easyops.api.ops_automation.menu.GetMenuTree"
}
func (*getMenuTreeContract) ContractVersion() string   { return "1.0" }
func (*getMenuTreeContract) Pattern() (string, string) { return "GET", "/menuTree" }
func (*getMenuTreeContract) Body() string              { return "" }

var _ListMenusContract = &listMenusContract{}

type listMenusContract struct{}

func (*listMenusContract) ServiceName() string          { return "menu.rpc" }
func (*listMenusContract) MethodName() string           { return "ListMenus" }
func (*listMenusContract) RequestMessage() interface{}  { return new(ListMenusRequest) }
func (*listMenusContract) ResponseMessage() interface{} { return new(ListMenusRequest) }
func (*listMenusContract) ContractName() string         { return "easyops.api.ops_automation.menu.ListMenus" }
func (*listMenusContract) ContractVersion() string      { return "1.0" }
func (*listMenusContract) Pattern() (string, string)    { return "GET", "/menus" }
func (*listMenusContract) Body() string                 { return "" }

var _UpdateMenuContract = &updateMenuContract{}

type updateMenuContract struct{}

func (*updateMenuContract) ServiceName() string          { return "menu.rpc" }
func (*updateMenuContract) MethodName() string           { return "UpdateMenu" }
func (*updateMenuContract) RequestMessage() interface{}  { return new(UpdateMenuRequest) }
func (*updateMenuContract) ResponseMessage() interface{} { return new(UpdateMenuRequest) }
func (*updateMenuContract) ContractName() string         { return "easyops.api.ops_automation.menu.UpdateMenu" }
func (*updateMenuContract) ContractVersion() string      { return "1.0" }
func (*updateMenuContract) Pattern() (string, string)    { return "PUT", "/menus/:menusId" }
func (*updateMenuContract) Body() string                 { return "" }