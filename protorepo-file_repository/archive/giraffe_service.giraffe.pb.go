// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package archive

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	file_repository "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/file_repository"
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

// Client is the client API for archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	DeleteArchive(ctx context.Context, in *DeleteArchiveRequest) (*DeleteArchiveResponse, error)
	DeleteArchiveV2(ctx context.Context, in *DeleteArchiveV2Request) (*DeleteArchiveV2Response, error)
	GetDiffSize(ctx context.Context, in *GetDiffSizeRequest) (*GetDiffSizeResponse, error)
	GetFileInfo(ctx context.Context, in *GetFileInfoRequest) (*GetFileInfoResponse, error)
	Listdir(ctx context.Context, in *ListdirRequest) (*ListdirResponse, error)
	GetSign(ctx context.Context, in *GetSignRequest) (*GetSignResponse, error)
	GetSize(ctx context.Context, in *GetSizeRequest) (*GetSizeResponse, error)
	GetDifference(ctx context.Context, in *GetDifferenceRequest) (*file_repository.Diff, error)
	GetPackageDifference(ctx context.Context, in *GetPackageDifferenceRequest) (*file_repository.Diff, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) DeleteArchive(ctx context.Context, in *DeleteArchiveRequest) (*DeleteArchiveResponse, error) {
	out := new(DeleteArchiveResponse)
	err := c.c.Invoke(ctx, _DeleteArchiveMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteArchiveV2(ctx context.Context, in *DeleteArchiveV2Request) (*DeleteArchiveV2Response, error) {
	out := new(DeleteArchiveV2Response)
	err := c.c.Invoke(ctx, _DeleteArchiveV2MethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetDiffSize(ctx context.Context, in *GetDiffSizeRequest) (*GetDiffSizeResponse, error) {
	out := new(GetDiffSizeResponse)
	err := c.c.Invoke(ctx, _GetDiffSizeMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetFileInfo(ctx context.Context, in *GetFileInfoRequest) (*GetFileInfoResponse, error) {
	out := new(GetFileInfoResponse)
	err := c.c.Invoke(ctx, _GetFileInfoMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Listdir(ctx context.Context, in *ListdirRequest) (*ListdirResponse, error) {
	out := new(ListdirResponse)
	err := c.c.Invoke(ctx, _ListdirMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetSign(ctx context.Context, in *GetSignRequest) (*GetSignResponse, error) {
	out := new(GetSignResponse)
	err := c.c.Invoke(ctx, _GetSignMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetSize(ctx context.Context, in *GetSizeRequest) (*GetSizeResponse, error) {
	out := new(GetSizeResponse)
	err := c.c.Invoke(ctx, _GetSizeMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetDifference(ctx context.Context, in *GetDifferenceRequest) (*file_repository.Diff, error) {
	out := new(file_repository.Diff)
	err := c.c.Invoke(ctx, _GetDifferenceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetPackageDifference(ctx context.Context, in *GetPackageDifferenceRequest) (*file_repository.Diff, error) {
	out := new(file_repository.Diff)
	err := c.c.Invoke(ctx, _GetPackageDifferenceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for archive service.
type Service interface {
	DeleteArchive(context.Context, *DeleteArchiveRequest) (*DeleteArchiveResponse, error)
	DeleteArchiveV2(context.Context, *DeleteArchiveV2Request) (*DeleteArchiveV2Response, error)
	GetDiffSize(context.Context, *GetDiffSizeRequest) (*GetDiffSizeResponse, error)
	GetFileInfo(context.Context, *GetFileInfoRequest) (*GetFileInfoResponse, error)
	Listdir(context.Context, *ListdirRequest) (*ListdirResponse, error)
	GetSign(context.Context, *GetSignRequest) (*GetSignResponse, error)
	GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error)
	GetDifference(context.Context, *GetDifferenceRequest) (*file_repository.Diff, error)
	GetPackageDifference(context.Context, *GetPackageDifferenceRequest) (*file_repository.Diff, error)
}

func _DeleteArchiveEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteArchive(ctx, req.(*DeleteArchiveRequest))
	}
}

func _DeleteArchiveV2Endpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteArchiveV2(ctx, req.(*DeleteArchiveV2Request))
	}
}

func _GetDiffSizeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDiffSize(ctx, req.(*GetDiffSizeRequest))
	}
}

func _GetFileInfoEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetFileInfo(ctx, req.(*GetFileInfoRequest))
	}
}

func _ListdirEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Listdir(ctx, req.(*ListdirRequest))
	}
}

func _GetSignEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSign(ctx, req.(*GetSignRequest))
	}
}

func _GetSizeEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSize(ctx, req.(*GetSizeRequest))
	}
}

func _GetDifferenceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDifference(ctx, req.(*GetDifferenceRequest))
	}
}

func _GetPackageDifferenceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetPackageDifference(ctx, req.(*GetPackageDifferenceRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_DeleteArchiveMethodDesc, _DeleteArchiveEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteArchiveV2MethodDesc, _DeleteArchiveV2Endpoint(srv))
	s.RegisterUnaryEndpoint(_GetDiffSizeMethodDesc, _GetDiffSizeEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetFileInfoMethodDesc, _GetFileInfoEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListdirMethodDesc, _ListdirEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetSignMethodDesc, _GetSignEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetSizeMethodDesc, _GetSizeEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetDifferenceMethodDesc, _GetDifferenceEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetPackageDifferenceMethodDesc, _GetPackageDifferenceEndpoint(srv))
}

// Method Description
var _DeleteArchiveMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.DeleteArchive",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "DeleteArchive",
	RequestType:  (*DeleteArchiveRequest)(nil),
	ResponseType: (*DeleteArchiveResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/archive/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _DeleteArchiveV2MethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.DeleteArchiveV2",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "DeleteArchiveV2",
	RequestType:  (*DeleteArchiveV2Request)(nil),
	ResponseType: (*DeleteArchiveV2Response)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/v2/archive/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetDiffSizeMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetDiffSize",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetDiffSize",
	RequestType:  (*GetDiffSizeRequest)(nil),
	ResponseType: (*GetDiffSizeResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/diff/size/:packageId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetFileInfoMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetFileInfo",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetFileInfo",
	RequestType:  (*GetFileInfoRequest)(nil),
	ResponseType: (*GetFileInfoResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/info/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListdirMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.Listdir",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "Listdir",
	RequestType:  (*ListdirRequest)(nil),
	ResponseType: (*ListdirResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/list/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetSignMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetSign",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetSign",
	RequestType:  (*GetSignRequest)(nil),
	ResponseType: (*GetSignResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/sign/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetSizeMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetSize",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetSize",
	RequestType:  (*GetSizeRequest)(nil),
	ResponseType: (*GetSizeResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/size/:packageId/:versionId",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _GetDifferenceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetDifference",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetDifference",
	RequestType:  (*GetDifferenceRequest)(nil),
	ResponseType: (*file_repository.Diff)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/difference/:packageId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetPackageDifferenceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.file_repository.archive.GetPackageDifference",
		Version: "1.0",
	},
	ServiceName:  "archive.rpc",
	MethodName:   "GetPackageDifference",
	RequestType:  (*GetPackageDifferenceRequest)(nil),
	ResponseType: (*file_repository.Diff)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/archive/difference",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
