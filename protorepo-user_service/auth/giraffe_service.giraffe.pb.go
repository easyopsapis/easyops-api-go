// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package auth

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

// Client is the client API for auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	GetCaptcha(ctx context.Context, in *types.Empty) (*types.Empty, error)
	UserAuth(ctx context.Context, in *UserAuthRequest) (*UserAuthResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) GetCaptcha(ctx context.Context, in *types.Empty) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _GetCaptchaMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UserAuth(ctx context.Context, in *UserAuthRequest) (*UserAuthResponse, error) {
	out := new(UserAuthResponse)
	err := c.c.Invoke(ctx, _UserAuthMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for auth service.
type Service interface {
	GetCaptcha(context.Context, *types.Empty) (*types.Empty, error)
	UserAuth(context.Context, *UserAuthRequest) (*UserAuthResponse, error)
}

func _GetCaptchaEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetCaptcha(ctx, req.(*types.Empty))
	}
}

func _UserAuthEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UserAuth(ctx, req.(*UserAuthRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_GetCaptchaMethodDesc, _GetCaptchaEndpoint(srv))
	s.RegisterUnaryEndpoint(_UserAuthMethodDesc, _UserAuthEndpoint(srv))
}

// Method Description
var _GetCaptchaMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.auth.GetCaptcha",
		Version: "1.0",
	},
	ServiceName:  "auth.rpc",
	MethodName:   "GetCaptcha",
	RequestType:  (*types.Empty)(nil),
	ResponseType: (*types.Empty)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/api/v1/users/captcha",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _UserAuthMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.user_service.auth.UserAuth",
		Version: "1.0",
	},
	ServiceName:  "auth.rpc",
	MethodName:   "UserAuth",
	RequestType:  (*UserAuthRequest)(nil),
	ResponseType: (*UserAuthResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/api/v1/users/auth",
		},
		Body:         "",
		ResponseBody: "data",
	},
}