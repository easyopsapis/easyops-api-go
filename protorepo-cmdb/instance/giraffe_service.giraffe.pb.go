// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giraffe_service.proto

package instance

import (
	context "context"
	fmt "fmt"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	go_proto_giraffe "github.com/easyops-cn/go-proto-giraffe"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Client is the client API for instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	AggregateCount(ctx context.Context, in *AggregateCountRequest) (*AggregateCountResponse, error)
	AutoDiscovery(ctx context.Context, in *AutoDiscoveryRequest) (*AutoDiscoveryResponse, error)
	BatchListInstanceQueryStrategy(ctx context.Context, in *BatchListInstanceQueryStrategyRequest) (*BatchListInstanceQueryStrategyResponse, error)
	BatchSearch(ctx context.Context, in *BatchSearchRequest) (*BatchSearchResponse, error)
	CreateInstance(ctx context.Context, in *CreateInstanceRequest) (*types.Struct, error)
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest) (*types.Struct, error)
	DeleteInstanceBatch(ctx context.Context, in *DeleteInstanceBatchRequest) (*DeleteInstanceBatchResponse, error)
	FulltextSearch(ctx context.Context, in *FulltextSearchRequest) (*FulltextSearchResponse, error)
	GetDefaultValueTemplate(ctx context.Context, in *GetDefaultValueTemplateRequest) (*types.Struct, error)
	GetDetail(ctx context.Context, in *GetDetailRequest) (*types.Struct, error)
	GetInstanceQueryStrategy(ctx context.Context, in *GetInstanceQueryStrategyRequest) (*cmdb.InstanceQueryStrategy, error)
	ImportInstance(ctx context.Context, in *ImportInstanceRequest) (*ImportInstanceResponse, error)
	InstanceRelationsCount(ctx context.Context, in *InstanceRelationsCountRequest) (*types.Struct, error)
	InstanceSnapshot(ctx context.Context, in *InstanceSnapshotRequest) (*types.Struct, error)
	ListInstance(ctx context.Context, in *ListInstanceRequest) (*ListInstanceResponse, error)
	ListInstanceQueryStrategy(ctx context.Context, in *ListInstanceQueryStrategyRequest) (*ListInstanceQueryStrategyResponse, error)
	PostSearch(ctx context.Context, in *PostSearchRequest) (*PostSearchResponse, error)
	RelationCountAggregate(ctx context.Context, in *RelationCountAggregateRequest) (*RelationCountAggregateResponse, error)
	RelationMaxCheck(ctx context.Context, in *RelationMaxCheckRequest) (*RelationMaxCheckResponse, error)
	SearchTotal(ctx context.Context, in *SearchTotalRequest) (*SearchTotalResponse, error)
	UpdateInstance(ctx context.Context, in *UpdateInstanceRequest) (*types.Struct, error)
	UpdateInstanceV2(ctx context.Context, in *UpdateInstanceV2Request) (*types.Struct, error)
	UpdatePermissionBatch(ctx context.Context, in *UpdatePermissionBatchRequest) (*UpdatePermissionBatchResponse, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) AggregateCount(ctx context.Context, in *AggregateCountRequest) (*AggregateCountResponse, error) {
	out := new(AggregateCountResponse)
	err := c.c.Invoke(ctx, _AggregateCountMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) AutoDiscovery(ctx context.Context, in *AutoDiscoveryRequest) (*AutoDiscoveryResponse, error) {
	out := new(AutoDiscoveryResponse)
	err := c.c.Invoke(ctx, _AutoDiscoveryMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchListInstanceQueryStrategy(ctx context.Context, in *BatchListInstanceQueryStrategyRequest) (*BatchListInstanceQueryStrategyResponse, error) {
	out := new(BatchListInstanceQueryStrategyResponse)
	err := c.c.Invoke(ctx, _BatchListInstanceQueryStrategyMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchSearch(ctx context.Context, in *BatchSearchRequest) (*BatchSearchResponse, error) {
	out := new(BatchSearchResponse)
	err := c.c.Invoke(ctx, _BatchSearchMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateInstance(ctx context.Context, in *CreateInstanceRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _CreateInstanceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _DeleteInstanceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteInstanceBatch(ctx context.Context, in *DeleteInstanceBatchRequest) (*DeleteInstanceBatchResponse, error) {
	out := new(DeleteInstanceBatchResponse)
	err := c.c.Invoke(ctx, _DeleteInstanceBatchMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) FulltextSearch(ctx context.Context, in *FulltextSearchRequest) (*FulltextSearchResponse, error) {
	out := new(FulltextSearchResponse)
	err := c.c.Invoke(ctx, _FulltextSearchMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetDefaultValueTemplate(ctx context.Context, in *GetDefaultValueTemplateRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _GetDefaultValueTemplateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetDetail(ctx context.Context, in *GetDetailRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _GetDetailMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetInstanceQueryStrategy(ctx context.Context, in *GetInstanceQueryStrategyRequest) (*cmdb.InstanceQueryStrategy, error) {
	out := new(cmdb.InstanceQueryStrategy)
	err := c.c.Invoke(ctx, _GetInstanceQueryStrategyMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ImportInstance(ctx context.Context, in *ImportInstanceRequest) (*ImportInstanceResponse, error) {
	out := new(ImportInstanceResponse)
	err := c.c.Invoke(ctx, _ImportInstanceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) InstanceRelationsCount(ctx context.Context, in *InstanceRelationsCountRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _InstanceRelationsCountMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) InstanceSnapshot(ctx context.Context, in *InstanceSnapshotRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _InstanceSnapshotMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListInstance(ctx context.Context, in *ListInstanceRequest) (*ListInstanceResponse, error) {
	out := new(ListInstanceResponse)
	err := c.c.Invoke(ctx, _ListInstanceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListInstanceQueryStrategy(ctx context.Context, in *ListInstanceQueryStrategyRequest) (*ListInstanceQueryStrategyResponse, error) {
	out := new(ListInstanceQueryStrategyResponse)
	err := c.c.Invoke(ctx, _ListInstanceQueryStrategyMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) PostSearch(ctx context.Context, in *PostSearchRequest) (*PostSearchResponse, error) {
	out := new(PostSearchResponse)
	err := c.c.Invoke(ctx, _PostSearchMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) RelationCountAggregate(ctx context.Context, in *RelationCountAggregateRequest) (*RelationCountAggregateResponse, error) {
	out := new(RelationCountAggregateResponse)
	err := c.c.Invoke(ctx, _RelationCountAggregateMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) RelationMaxCheck(ctx context.Context, in *RelationMaxCheckRequest) (*RelationMaxCheckResponse, error) {
	out := new(RelationMaxCheckResponse)
	err := c.c.Invoke(ctx, _RelationMaxCheckMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SearchTotal(ctx context.Context, in *SearchTotalRequest) (*SearchTotalResponse, error) {
	out := new(SearchTotalResponse)
	err := c.c.Invoke(ctx, _SearchTotalMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateInstance(ctx context.Context, in *UpdateInstanceRequest) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _UpdateInstanceMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateInstanceV2(ctx context.Context, in *UpdateInstanceV2Request) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _UpdateInstanceV2MethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdatePermissionBatch(ctx context.Context, in *UpdatePermissionBatchRequest) (*UpdatePermissionBatchResponse, error) {
	out := new(UpdatePermissionBatchResponse)
	err := c.c.Invoke(ctx, _UpdatePermissionBatchMethodDesc, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for instance service.
type Service interface {
	AggregateCount(context.Context, *AggregateCountRequest) (*AggregateCountResponse, error)
	AutoDiscovery(context.Context, *AutoDiscoveryRequest) (*AutoDiscoveryResponse, error)
	BatchListInstanceQueryStrategy(context.Context, *BatchListInstanceQueryStrategyRequest) (*BatchListInstanceQueryStrategyResponse, error)
	BatchSearch(context.Context, *BatchSearchRequest) (*BatchSearchResponse, error)
	CreateInstance(context.Context, *CreateInstanceRequest) (*types.Struct, error)
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*types.Struct, error)
	DeleteInstanceBatch(context.Context, *DeleteInstanceBatchRequest) (*DeleteInstanceBatchResponse, error)
	FulltextSearch(context.Context, *FulltextSearchRequest) (*FulltextSearchResponse, error)
	GetDefaultValueTemplate(context.Context, *GetDefaultValueTemplateRequest) (*types.Struct, error)
	GetDetail(context.Context, *GetDetailRequest) (*types.Struct, error)
	GetInstanceQueryStrategy(context.Context, *GetInstanceQueryStrategyRequest) (*cmdb.InstanceQueryStrategy, error)
	ImportInstance(context.Context, *ImportInstanceRequest) (*ImportInstanceResponse, error)
	InstanceRelationsCount(context.Context, *InstanceRelationsCountRequest) (*types.Struct, error)
	InstanceSnapshot(context.Context, *InstanceSnapshotRequest) (*types.Struct, error)
	ListInstance(context.Context, *ListInstanceRequest) (*ListInstanceResponse, error)
	ListInstanceQueryStrategy(context.Context, *ListInstanceQueryStrategyRequest) (*ListInstanceQueryStrategyResponse, error)
	PostSearch(context.Context, *PostSearchRequest) (*PostSearchResponse, error)
	RelationCountAggregate(context.Context, *RelationCountAggregateRequest) (*RelationCountAggregateResponse, error)
	RelationMaxCheck(context.Context, *RelationMaxCheckRequest) (*RelationMaxCheckResponse, error)
	SearchTotal(context.Context, *SearchTotalRequest) (*SearchTotalResponse, error)
	UpdateInstance(context.Context, *UpdateInstanceRequest) (*types.Struct, error)
	UpdateInstanceV2(context.Context, *UpdateInstanceV2Request) (*types.Struct, error)
	UpdatePermissionBatch(context.Context, *UpdatePermissionBatchRequest) (*UpdatePermissionBatchResponse, error)
}

func _AggregateCountEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AggregateCount(ctx, req.(*AggregateCountRequest))
	}
}

func _AutoDiscoveryEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AutoDiscovery(ctx, req.(*AutoDiscoveryRequest))
	}
}

func _BatchListInstanceQueryStrategyEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BatchListInstanceQueryStrategy(ctx, req.(*BatchListInstanceQueryStrategyRequest))
	}
}

func _BatchSearchEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BatchSearch(ctx, req.(*BatchSearchRequest))
	}
}

func _CreateInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
}

func _DeleteInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
}

func _DeleteInstanceBatchEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteInstanceBatch(ctx, req.(*DeleteInstanceBatchRequest))
	}
}

func _FulltextSearchEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.FulltextSearch(ctx, req.(*FulltextSearchRequest))
	}
}

func _GetDefaultValueTemplateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDefaultValueTemplate(ctx, req.(*GetDefaultValueTemplateRequest))
	}
}

func _GetDetailEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDetail(ctx, req.(*GetDetailRequest))
	}
}

func _GetInstanceQueryStrategyEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetInstanceQueryStrategy(ctx, req.(*GetInstanceQueryStrategyRequest))
	}
}

func _ImportInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ImportInstance(ctx, req.(*ImportInstanceRequest))
	}
}

func _InstanceRelationsCountEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.InstanceRelationsCount(ctx, req.(*InstanceRelationsCountRequest))
	}
}

func _InstanceSnapshotEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.InstanceSnapshot(ctx, req.(*InstanceSnapshotRequest))
	}
}

func _ListInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListInstance(ctx, req.(*ListInstanceRequest))
	}
}

func _ListInstanceQueryStrategyEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListInstanceQueryStrategy(ctx, req.(*ListInstanceQueryStrategyRequest))
	}
}

func _PostSearchEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PostSearch(ctx, req.(*PostSearchRequest))
	}
}

func _RelationCountAggregateEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RelationCountAggregate(ctx, req.(*RelationCountAggregateRequest))
	}
}

func _RelationMaxCheckEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RelationMaxCheck(ctx, req.(*RelationMaxCheckRequest))
	}
}

func _SearchTotalEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SearchTotal(ctx, req.(*SearchTotalRequest))
	}
}

func _UpdateInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateInstance(ctx, req.(*UpdateInstanceRequest))
	}
}

func _UpdateInstanceV2Endpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateInstanceV2(ctx, req.(*UpdateInstanceV2Request))
	}
}

func _UpdatePermissionBatchEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdatePermissionBatch(ctx, req.(*UpdatePermissionBatchRequest))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_AggregateCountMethodDesc, _AggregateCountEndpoint(srv))
	s.RegisterUnaryEndpoint(_AutoDiscoveryMethodDesc, _AutoDiscoveryEndpoint(srv))
	s.RegisterUnaryEndpoint(_BatchListInstanceQueryStrategyMethodDesc, _BatchListInstanceQueryStrategyEndpoint(srv))
	s.RegisterUnaryEndpoint(_BatchSearchMethodDesc, _BatchSearchEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateInstanceMethodDesc, _CreateInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteInstanceMethodDesc, _DeleteInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteInstanceBatchMethodDesc, _DeleteInstanceBatchEndpoint(srv))
	s.RegisterUnaryEndpoint(_FulltextSearchMethodDesc, _FulltextSearchEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetDefaultValueTemplateMethodDesc, _GetDefaultValueTemplateEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetDetailMethodDesc, _GetDetailEndpoint(srv))
	s.RegisterUnaryEndpoint(_GetInstanceQueryStrategyMethodDesc, _GetInstanceQueryStrategyEndpoint(srv))
	s.RegisterUnaryEndpoint(_ImportInstanceMethodDesc, _ImportInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_InstanceRelationsCountMethodDesc, _InstanceRelationsCountEndpoint(srv))
	s.RegisterUnaryEndpoint(_InstanceSnapshotMethodDesc, _InstanceSnapshotEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListInstanceMethodDesc, _ListInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_ListInstanceQueryStrategyMethodDesc, _ListInstanceQueryStrategyEndpoint(srv))
	s.RegisterUnaryEndpoint(_PostSearchMethodDesc, _PostSearchEndpoint(srv))
	s.RegisterUnaryEndpoint(_RelationCountAggregateMethodDesc, _RelationCountAggregateEndpoint(srv))
	s.RegisterUnaryEndpoint(_RelationMaxCheckMethodDesc, _RelationMaxCheckEndpoint(srv))
	s.RegisterUnaryEndpoint(_SearchTotalMethodDesc, _SearchTotalEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateInstanceMethodDesc, _UpdateInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_UpdateInstanceV2MethodDesc, _UpdateInstanceV2Endpoint(srv))
	s.RegisterUnaryEndpoint(_UpdatePermissionBatchMethodDesc, _UpdatePermissionBatchEndpoint(srv))
}

// Method Description
var _AggregateCountMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.AggregateCount",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "AggregateCount",
	RequestType:  (*AggregateCountRequest)(nil),
	ResponseType: (*AggregateCountResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:objectId/instance/aggregate/count/:attrId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _AutoDiscoveryMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.AutoDiscovery",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "AutoDiscovery",
	RequestType:  (*AutoDiscoveryRequest)(nil),
	ResponseType: (*AutoDiscoveryResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_import-json",
		},
		Body:         "body",
		ResponseBody: "",
	},
}

var _BatchListInstanceQueryStrategyMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.BatchListInstanceQueryStrategy",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "BatchListInstanceQueryStrategy",
	RequestType:  (*BatchListInstanceQueryStrategyRequest)(nil),
	ResponseType: (*BatchListInstanceQueryStrategyResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "batch/object/query/strategy",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _BatchSearchMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.BatchSearch",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "BatchSearch",
	RequestType:  (*BatchSearchRequest)(nil),
	ResponseType: (*BatchSearchResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/batch/object/instance/_search",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _CreateInstanceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.CreateInstance",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "CreateInstance",
	RequestType:  (*CreateInstanceRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/v2/object/:objectId/instance",
		},
		Body:         "instance",
		ResponseBody: "data",
	},
}

var _DeleteInstanceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.DeleteInstance",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "DeleteInstance",
	RequestType:  (*DeleteInstanceRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/object/:objectId/instance/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _DeleteInstanceBatchMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.DeleteInstanceBatch",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "DeleteInstanceBatch",
	RequestType:  (*DeleteInstanceBatchRequest)(nil),
	ResponseType: (*DeleteInstanceBatchResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Delete{
			Delete: "/object/:objectId/instance/_batch",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _FulltextSearchMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.FulltextSearch",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "FulltextSearch",
	RequestType:  (*FulltextSearchRequest)(nil),
	ResponseType: (*FulltextSearchResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/fulltext/_search",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetDefaultValueTemplateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.GetDefaultValueTemplate",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "GetDefaultValueTemplate",
	RequestType:  (*GetDefaultValueTemplateRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:objectId/instance_default_value_template",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetDetailMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.GetDetail",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "GetDetail",
	RequestType:  (*GetDetailRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:objectId/instance/:instanceId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _GetInstanceQueryStrategyMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.GetInstanceQueryStrategy",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "GetInstanceQueryStrategy",
	RequestType:  (*GetInstanceQueryStrategyRequest)(nil),
	ResponseType: (*cmdb.InstanceQueryStrategy)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:objectId/query/strategy/:strategyId",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ImportInstanceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.ImportInstance",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "ImportInstance",
	RequestType:  (*ImportInstanceRequest)(nil),
	ResponseType: (*ImportInstanceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_import",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _InstanceRelationsCountMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.InstanceRelationsCount",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "InstanceRelationsCount",
	RequestType:  (*InstanceRelationsCountRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:object_id/instance/:instance_id/_relation_count",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _InstanceSnapshotMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.InstanceSnapshot",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "InstanceSnapshot",
	RequestType:  (*InstanceSnapshotRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/history/object/:object_id/instance/:instance_id",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListInstanceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.ListInstance",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "ListInstance",
	RequestType:  (*ListInstanceRequest)(nil),
	ResponseType: (*ListInstanceResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:object_id/instance",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _ListInstanceQueryStrategyMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.ListInstanceQueryStrategy",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "ListInstanceQueryStrategy",
	RequestType:  (*ListInstanceQueryStrategyRequest)(nil),
	ResponseType: (*ListInstanceQueryStrategyResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Get{
			Get: "/object/:object_id/query/strategy",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _PostSearchMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.PostSearch",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "PostSearch",
	RequestType:  (*PostSearchRequest)(nil),
	ResponseType: (*PostSearchResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_search",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _RelationCountAggregateMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.RelationCountAggregate",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "RelationCountAggregate",
	RequestType:  (*RelationCountAggregateRequest)(nil),
	ResponseType: (*RelationCountAggregateResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_relation_count_aggregate",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _RelationMaxCheckMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.RelationMaxCheck",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "RelationMaxCheck",
	RequestType:  (*RelationMaxCheckRequest)(nil),
	ResponseType: (*RelationMaxCheckResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_search_relation_max_check",
		},
		Body:         "",
		ResponseBody: "data",
	},
}

var _SearchTotalMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.SearchTotal",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "SearchTotal",
	RequestType:  (*SearchTotalRequest)(nil),
	ResponseType: (*SearchTotalResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Post{
			Post: "/object/:objectId/instance/_search_total",
		},
		Body:         "",
		ResponseBody: "",
	},
}

var _UpdateInstanceMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.UpdateInstance",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "UpdateInstance",
	RequestType:  (*UpdateInstanceRequest)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/object/:objectId/instance/:instanceId",
		},
		Body:         "instance",
		ResponseBody: "data",
	},
}

var _UpdateInstanceV2MethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.UpdateInstanceV2",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "UpdateInstanceV2",
	RequestType:  (*UpdateInstanceV2Request)(nil),
	ResponseType: (*types.Struct)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/v2/object/:objectId/instance/:instanceId",
		},
		Body:         "instance",
		ResponseBody: "data",
	},
}

var _UpdatePermissionBatchMethodDesc = &giraffe_micro.MethodDesc{
	Contract: &go_proto_giraffe.Contract{
		Name:    "easyops.api.cmdb.instance.UpdatePermissionBatch",
		Version: "1.0",
	},
	ServiceName:  "instance.rpc",
	MethodName:   "UpdatePermissionBatch",
	RequestType:  (*UpdatePermissionBatchRequest)(nil),
	ResponseType: (*UpdatePermissionBatchResponse)(nil),
	HttpRule: &go_proto_giraffe.HttpRule{
		Pattern: &go_proto_giraffe.HttpRule_Put{
			Put: "/permission/:objectId/instances/_batch",
		},
		Body:         "",
		ResponseBody: "data",
	},
}
