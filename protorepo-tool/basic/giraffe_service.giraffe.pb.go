// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package basic

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	tool "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
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

// Client is the client API for basic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	BatchGetToolDetail(ctx context.Context, in *BatchGetToolDetailRequest) (*BatchGetToolDetailResponse, error)
	BatchUpdateToolPermission(ctx context.Context, in *BatchUpdateToolPermissionRequest) (*BatchUpdateToolPermissionResponse, error)
	CreateTool(ctx context.Context, in *tool.Tool) (*CreateToolResponse, error)
	DeleteTool(ctx context.Context, in *DeleteToolRequest) (*types.Empty, error)
	DisableTool(ctx context.Context, in *DisableToolRequest) (*DisableToolResponse, error)
	GetReferenceToolObject(ctx context.Context, in *GetReferenceToolObjectRequest) (*GetReferenceToolObjectResponse, error)
	GetReferenceToolObjectByVid(ctx context.Context, in *GetReferenceToolObjectByVidRequest) (*GetReferenceToolObjectByVidResponse, error)
	GetTool(ctx context.Context, in *GetToolRequest) (*tool.Tool, error)
	GetToolCategories(ctx context.Context, in *types.Empty) (*GetToolCategoriesResponse, error)
	GetToolGlobalConfig(ctx context.Context, in *types.Empty) (*GetToolGlobalConfigResponse, error)
	ListTool(ctx context.Context, in *ListToolRequest) (*ListToolResponse, error)
	ToolApproval(ctx context.Context, in *ToolApprovalRequest) (*ToolApprovalResponse, error)
	UpdateTool(ctx context.Context, in *tool.Tool) (*tool.Tool, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) BatchGetToolDetail(ctx context.Context, in *BatchGetToolDetailRequest) (*BatchGetToolDetailResponse, error) {
	out := new(BatchGetToolDetailResponse)
	err := c.c.Invoke(ctx, _BatchGetToolDetailMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchUpdateToolPermission(ctx context.Context, in *BatchUpdateToolPermissionRequest) (*BatchUpdateToolPermissionResponse, error) {
	out := new(BatchUpdateToolPermissionResponse)
	err := c.c.Invoke(ctx, _BatchUpdateToolPermissionMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateTool(ctx context.Context, in *tool.Tool) (*CreateToolResponse, error) {
	out := new(CreateToolResponse)
	err := c.c.Invoke(ctx, _CreateToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTool(ctx context.Context, in *DeleteToolRequest) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DisableTool(ctx context.Context, in *DisableToolRequest) (*DisableToolResponse, error) {
	out := new(DisableToolResponse)
	err := c.c.Invoke(ctx, _DisableToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetReferenceToolObject(ctx context.Context, in *GetReferenceToolObjectRequest) (*GetReferenceToolObjectResponse, error) {
	out := new(GetReferenceToolObjectResponse)
	err := c.c.Invoke(ctx, _GetReferenceToolObjectMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetReferenceToolObjectByVid(ctx context.Context, in *GetReferenceToolObjectByVidRequest) (*GetReferenceToolObjectByVidResponse, error) {
	out := new(GetReferenceToolObjectByVidResponse)
	err := c.c.Invoke(ctx, _GetReferenceToolObjectByVidMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetTool(ctx context.Context, in *GetToolRequest) (*tool.Tool, error) {
	out := new(tool.Tool)
	err := c.c.Invoke(ctx, _GetToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetToolCategories(ctx context.Context, in *types.Empty) (*GetToolCategoriesResponse, error) {
	out := new(GetToolCategoriesResponse)
	err := c.c.Invoke(ctx, _GetToolCategoriesMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetToolGlobalConfig(ctx context.Context, in *types.Empty) (*GetToolGlobalConfigResponse, error) {
	out := new(GetToolGlobalConfigResponse)
	err := c.c.Invoke(ctx, _GetToolGlobalConfigMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListTool(ctx context.Context, in *ListToolRequest) (*ListToolResponse, error) {
	out := new(ListToolResponse)
	err := c.c.Invoke(ctx, _ListToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ToolApproval(ctx context.Context, in *ToolApprovalRequest) (*ToolApprovalResponse, error) {
	out := new(ToolApprovalResponse)
	err := c.c.Invoke(ctx, _ToolApprovalMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateTool(ctx context.Context, in *tool.Tool) (*tool.Tool, error) {
	out := new(tool.Tool)
	err := c.c.Invoke(ctx, _UpdateToolMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for basic service.
type Service interface {
	BatchGetToolDetail(context.Context, *BatchGetToolDetailRequest) (*BatchGetToolDetailResponse, error)
	BatchUpdateToolPermission(context.Context, *BatchUpdateToolPermissionRequest) (*BatchUpdateToolPermissionResponse, error)
	CreateTool(context.Context, *tool.Tool) (*CreateToolResponse, error)
	DeleteTool(context.Context, *DeleteToolRequest) (*types.Empty, error)
	DisableTool(context.Context, *DisableToolRequest) (*DisableToolResponse, error)
	GetReferenceToolObject(context.Context, *GetReferenceToolObjectRequest) (*GetReferenceToolObjectResponse, error)
	GetReferenceToolObjectByVid(context.Context, *GetReferenceToolObjectByVidRequest) (*GetReferenceToolObjectByVidResponse, error)
	GetTool(context.Context, *GetToolRequest) (*tool.Tool, error)
	GetToolCategories(context.Context, *types.Empty) (*GetToolCategoriesResponse, error)
	GetToolGlobalConfig(context.Context, *types.Empty) (*GetToolGlobalConfigResponse, error)
	ListTool(context.Context, *ListToolRequest) (*ListToolResponse, error)
	ToolApproval(context.Context, *ToolApprovalRequest) (*ToolApprovalResponse, error)
	UpdateTool(context.Context, *tool.Tool) (*tool.Tool, error)
}

func _BatchGetToolDetailEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BatchGetToolDetail(ctx, req.(*BatchGetToolDetailRequest))
	}
}

func _BatchUpdateToolPermissionEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BatchUpdateToolPermission(ctx, req.(*BatchUpdateToolPermissionRequest))
	}
}

func _CreateToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateTool(ctx, req.(*tool.Tool))
	}
}

func _DeleteToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteTool(ctx, req.(*DeleteToolRequest))
	}
}

func _DisableToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DisableTool(ctx, req.(*DisableToolRequest))
	}
}

func _GetReferenceToolObjectEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetReferenceToolObject(ctx, req.(*GetReferenceToolObjectRequest))
	}
}

func _GetReferenceToolObjectByVidEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetReferenceToolObjectByVid(ctx, req.(*GetReferenceToolObjectByVidRequest))
	}
}

func _GetToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetTool(ctx, req.(*GetToolRequest))
	}
}

func _GetToolCategoriesEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetToolCategories(ctx, req.(*types.Empty))
	}
}

func _GetToolGlobalConfigEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetToolGlobalConfig(ctx, req.(*types.Empty))
	}
}

func _ListToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListTool(ctx, req.(*ListToolRequest))
	}
}

func _ToolApprovalEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ToolApproval(ctx, req.(*ToolApprovalRequest))
	}
}

func _UpdateToolEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateTool(ctx, req.(*tool.Tool))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_BatchGetToolDetailMethodDesc, _BatchGetToolDetailEndpoint(srv))
	s.RegisterUnaryEndpoint(_BatchUpdateToolPermissionMethodDesc, _BatchUpdateToolPermissionEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateToolMethodDesc, _CreateToolEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteToolMethodDesc, _DeleteToolEndpoint(srv))
	s.RegisterUnaryEndpoint(_DisableToolMethodDesc, _DisableToolEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetReferenceToolObjectMethodDesc, _GetReferenceToolObjectEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetReferenceToolObjectByVidMethodDesc, _GetReferenceToolObjectByVidEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetToolMethodDesc, _GetToolEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetToolCategoriesMethodDesc, _GetToolCategoriesEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetToolGlobalConfigMethodDesc, _GetToolGlobalConfigEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListToolMethodDesc, _ListToolEndpoint(srv))
	s.RegisterUnaryEndpoint(_ToolApprovalMethodDesc, _ToolApprovalEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateToolMethodDesc, _UpdateToolEndpoint(srv))
}

// Method Description
var _BatchGetToolDetailMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.BatchGetToolDetail",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "BatchGetToolDetail",
	RequestType:  (*BatchGetToolDetailRequest)(nil),
	ResponseType: (*BatchGetToolDetailResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/tools/batch/:toolIds",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _BatchUpdateToolPermissionMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.BatchUpdateToolPermission",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "BatchUpdateToolPermission",
	RequestType:  (*BatchUpdateToolPermissionRequest)(nil),
	ResponseType: (*BatchUpdateToolPermissionResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/tools/batch/permission",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _CreateToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.CreateTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "CreateTool",
	RequestType:  (*tool.Tool)(nil),
	ResponseType: (*CreateToolResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/tools",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.DeleteTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "DeleteTool",
	RequestType:  (*DeleteToolRequest)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/tools/:toolId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DisableToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.DisableTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "DisableTool",
	RequestType:  (*DisableToolRequest)(nil),
	ResponseType: (*DisableToolResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/tools/:toolId/disable",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetReferenceToolObjectMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.GetReferenceToolObject",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "GetReferenceToolObject",
	RequestType:  (*GetReferenceToolObjectRequest)(nil),
	ResponseType: (*GetReferenceToolObjectResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/referenceToolObject/:toolId",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetReferenceToolObjectByVidMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.GetReferenceToolObjectByVid",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "GetReferenceToolObjectByVid",
	RequestType:  (*GetReferenceToolObjectByVidRequest)(nil),
	ResponseType: (*GetReferenceToolObjectByVidResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/referenceToolObject/:toolId//version/:vId",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.GetTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "GetTool",
	RequestType:  (*GetToolRequest)(nil),
	ResponseType: (*tool.Tool)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/tools/:toolId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetToolCategoriesMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.GetToolCategories",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "GetToolCategories",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*GetToolCategoriesResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/tool_categories",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetToolGlobalConfigMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.GetToolGlobalConfig",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "GetToolGlobalConfig",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*GetToolGlobalConfigResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/toolsGlobalConfig",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.ListTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "ListTool",
	RequestType:  (*ListToolRequest)(nil),
	ResponseType: (*ListToolResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/tools",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ToolApprovalMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.ToolApproval",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "ToolApproval",
	RequestType:  (*ToolApprovalRequest)(nil),
	ResponseType: (*ToolApprovalResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/tools/:toolId/changeEnvType",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UpdateToolMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.tool.basic.UpdateTool",
		Version: "1.0",
	},
	ServiceName:  "basic.rpc",
	MethodName:   "UpdateTool",
	RequestType:  (*tool.Tool)(nil),
	ResponseType: (*tool.Tool)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/tools/:toolId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
