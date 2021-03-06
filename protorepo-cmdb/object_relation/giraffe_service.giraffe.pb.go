// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package object_relation

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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

// Client is the client API for object_relation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	Create(ctx context.Context, in *CreateRequest) (*cmdb.ObjectRelation, error)
	ObjectRelationSnapshot(ctx context.Context, in *ObjectRelationSnapshotRequest) (*ObjectRelationSnapshotResponse, error)
	GetRelationsByGroupId(ctx context.Context, in *GetRelationsByGroupIdRequest) (*GetRelationsByGroupIdResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) Create(ctx context.Context, in *CreateRequest) (*cmdb.ObjectRelation, error) {
	out := new(cmdb.ObjectRelation)
	err := c.c.Invoke(ctx, _CreateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ObjectRelationSnapshot(ctx context.Context, in *ObjectRelationSnapshotRequest) (*ObjectRelationSnapshotResponse, error) {
	out := new(ObjectRelationSnapshotResponse)
	err := c.c.Invoke(ctx, _ObjectRelationSnapshotMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetRelationsByGroupId(ctx context.Context, in *GetRelationsByGroupIdRequest) (*GetRelationsByGroupIdResponse, error) {
	out := new(GetRelationsByGroupIdResponse)
	err := c.c.Invoke(ctx, _GetRelationsByGroupIdMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for object_relation service.
type Service interface {
	Create(context.Context, *CreateRequest) (*cmdb.ObjectRelation, error)
	ObjectRelationSnapshot(context.Context, *ObjectRelationSnapshotRequest) (*ObjectRelationSnapshotResponse, error)
	GetRelationsByGroupId(context.Context, *GetRelationsByGroupIdRequest) (*GetRelationsByGroupIdResponse, error)
}

func _CreateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
}

func _ObjectRelationSnapshotEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ObjectRelationSnapshot(ctx, req.(*ObjectRelationSnapshotRequest))
	}
}

func _GetRelationsByGroupIdEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetRelationsByGroupId(ctx, req.(*GetRelationsByGroupIdRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_CreateMethodDesc, _CreateEndpoint(srv))
	s.RegisterUnaryEndpoint(_ObjectRelationSnapshotMethodDesc, _ObjectRelationSnapshotEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetRelationsByGroupIdMethodDesc, _GetRelationsByGroupIdEndpoint(srv))
}

// Method Description
var _CreateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.object_relation.Create",
		Version: "1.0",
	},
	ServiceName:  "object_relation.rpc",
	MethodName:   "Create",
	RequestType:  (*CreateRequest)(nil),
	ResponseType: (*cmdb.ObjectRelation)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object_relation",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ObjectRelationSnapshotMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.object_relation.ObjectRelationSnapshot",
		Version: "1.0",
	},
	ServiceName:  "object_relation.rpc",
	MethodName:   "ObjectRelationSnapshot",
	RequestType:  (*ObjectRelationSnapshotRequest)(nil),
	ResponseType: (*ObjectRelationSnapshotResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/history/object_relation/:relation_id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetRelationsByGroupIdMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.object_relation.GetRelationsByGroupId",
		Version: "1.0",
	},
	ServiceName:  "object_relation.rpc",
	MethodName:   "GetRelationsByGroupId",
	RequestType:  (*GetRelationsByGroupIdRequest)(nil),
	ResponseType: (*GetRelationsByGroupIdResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object_relation/object/:object_id/relation_group/:group_id",
		},
		Body:         "",
		ResponseBody: "",
	},
}
