// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package namespace

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	container "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Create请求
type CreateRequest struct {
	//
	//k8s 集群 id
	ClusterId string `protobuf:"bytes,1,opt,name=clusterId,proto3" json:"clusterId" form:"clusterId"`
	//
	//命名空间
	Namespace *container.Namespace `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" form:"namespace"`
	//
	//namespace 资源限制
	//{
	//"count/deployments.apps": "100",   #  无状态工作负载
	//"count/statefulsets.apps": "100",  #  有状态工作负载
	//"count/daemonsets.apps": "100",    #  守护进程
	//"count/jobs.batch": "100",     # 任务
	//"count/cronjobs.batch": "100", # 定时任务
	//"count/pods": "500",      # Pod
	//"count/services": "100",  # 服务
	//"count/ingresses.extensions": "100", # 应用路由
	//"persistentvolumeclaims": "100", # 存储卷
	//"count/secrets": "100",    # 密钥
	//"count/configmaps": "100", # 配置
	//"limits.cpu": "10",        # CPU 限额
	//"limits.memory": "20Gi"    # 内存限额
	//"requests.cpu": "1",       # CPU 预留
	//"requests.memory": "100Mi" # 内存预留
	//}
	//
	NamespaceQuota *types.Struct `protobuf:"bytes,3,opt,name=namespaceQuota,proto3" json:"namespaceQuota" form:"namespaceQuota"`
	//
	//最大资源限制和初始资源请求
	ContainerDefaultLimit *container.ResourceRequirements `protobuf:"bytes,4,opt,name=containerDefaultLimit,proto3" json:"containerDefaultLimit" form:"containerDefaultLimit"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateRequest) GetNamespace() *container.Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *CreateRequest) GetNamespaceQuota() *types.Struct {
	if m != nil {
		return m.NamespaceQuota
	}
	return nil
}

func (m *CreateRequest) GetContainerDefaultLimit() *container.ResourceRequirements {
	if m != nil {
		return m.ContainerDefaultLimit
	}
	return nil
}

//
//Create返回
type CreateResponse struct {
	//
	//命名空间
	Namespace *container.Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace" form:"namespace"`
	//
	//namespace 资源限制
	//{
	//"limits.cpu": "2",
	//"limits.memory": "1Gi"
	//"configmaps": 50,
	//"persistentvolumeclaims": 50,
	//"pods": 50,
	//"replicationcontrollers": 50,
	//"resourcequotas": 50,
	//"secrets": 50,
	//"services": 50,
	//"service.loadbalancers": 50,
	//"service.nodeports": 50
	//}
	//
	NamespaceQuota *types.Struct `protobuf:"bytes,2,opt,name=namespaceQuota,proto3" json:"namespaceQuota" form:"namespaceQuota"`
	//
	//最大资源限制和初始资源请求
	ContainerDefaultLimit *container.ResourceRequirements `protobuf:"bytes,3,opt,name=containerDefaultLimit,proto3" json:"containerDefaultLimit" form:"containerDefaultLimit"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetNamespace() *container.Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *CreateResponse) GetNamespaceQuota() *types.Struct {
	if m != nil {
		return m.NamespaceQuota
	}
	return nil
}

func (m *CreateResponse) GetContainerDefaultLimit() *container.ResourceRequirements {
	if m != nil {
		return m.ContainerDefaultLimit
	}
	return nil
}

//
//CreateApi返回
type CreateResponseWrapper struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//返回码解释
	CodeExplain string `protobuf:"bytes,2,opt,name=codeExplain,proto3" json:"codeExplain" form:"codeExplain"`
	//
	//错误详情
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回数据
	Data                 *CreateResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{2}
}
func (m *CreateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseWrapper.Unmarshal(m, b)
}
func (m *CreateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseWrapper.Merge(m, src)
}
func (m *CreateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateResponseWrapper.Size(m)
}
func (m *CreateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseWrapper proto.InternalMessageInfo

func (m *CreateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponseWrapper) GetData() *CreateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "namespace.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "namespace.CreateResponse")
	proto.RegisterType((*CreateResponseWrapper)(nil), "namespace.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xd2, 0x14, 0x29, 0x9b, 0x12, 0x22, 0xab, 0x81, 0xb4, 0xaa, 0xe4, 0x68, 0x41, 0x28,
	0x17, 0xdb, 0x40, 0x25, 0x04, 0x1c, 0x38, 0x98, 0x0f, 0x81, 0x84, 0x90, 0xd8, 0x1e, 0x90, 0x40,
	0x50, 0x6d, 0xd6, 0x13, 0x63, 0x61, 0x7b, 0xdc, 0xdd, 0x35, 0xe1, 0x43, 0xfc, 0x26, 0xfe, 0x91,
	0x91, 0xf8, 0x09, 0x3e, 0x70, 0xae, 0xba, 0x76, 0xec, 0xa4, 0xea, 0xb1, 0x3d, 0x65, 0x27, 0x6f,
	0xde, 0x9b, 0x79, 0x6f, 0xa2, 0x90, 0x1d, 0x21, 0x81, 0x6b, 0x70, 0x33, 0x89, 0x1a, 0xad, 0x7e,
	0xca, 0x13, 0x50, 0x19, 0x17, 0xb0, 0xef, 0x84, 0x91, 0xfe, 0x92, 0xcf, 0x5d, 0x81, 0x89, 0x17,
	0x62, 0x88, 0x9e, 0xe9, 0x98, 0xe7, 0x0b, 0x53, 0x99, 0xc2, 0xbc, 0x2a, 0xe6, 0xfe, 0x51, 0x88,
	0x2e, 0x70, 0xf5, 0x03, 0x33, 0xe5, 0xc6, 0x28, 0x78, 0xec, 0x09, 0x4c, 0xb5, 0xe4, 0x42, 0xab,
	0x8a, 0x29, 0x21, 0x43, 0x27, 0xc1, 0x00, 0x62, 0xe5, 0xd5, 0x8d, 0x9e, 0x29, 0x4d, 0x23, 0x8f,
	0x52, 0x90, 0x5e, 0x33, 0xbd, 0x16, 0x3d, 0xbe, 0x0c, 0x51, 0x09, 0x0a, 0x73, 0x29, 0xe0, 0x58,
	0xc2, 0x49, 0x1e, 0x49, 0x48, 0x20, 0xd5, 0xaa, 0x1e, 0xf0, 0x70, 0xcd, 0x64, 0xb2, 0x8c, 0xf4,
	0x57, 0x5c, 0x7a, 0x21, 0x3a, 0x06, 0x74, 0xbe, 0xf1, 0x38, 0x0a, 0xb8, 0x46, 0xa9, 0xbc, 0xe6,
	0x59, 0xf3, 0x0e, 0x42, 0xc4, 0x30, 0x86, 0x36, 0x13, 0xa5, 0x65, 0x2e, 0x74, 0x85, 0xd2, 0xff,
	0x5d, 0x72, 0xfd, 0x99, 0x89, 0x95, 0xc1, 0x49, 0x0e, 0x4a, 0x5b, 0x2f, 0x49, 0x5f, 0xc4, 0xb9,
	0xd2, 0x20, 0x5f, 0x07, 0x93, 0xce, 0xb4, 0x33, 0xeb, 0xfb, 0xb3, 0xb2, 0xb0, 0x47, 0x0b, 0x94,
	0xc9, 0x13, 0xda, 0x40, 0xf4, 0xdf, 0x5f, 0x7b, 0x44, 0x86, 0x9f, 0x3f, 0xde, 0x73, 0x1e, 0x73,
	0xe7, 0xe7, 0xa7, 0x5f, 0xf7, 0x0f, 0x7f, 0xdf, 0x61, 0x2d, 0xd5, 0x7a, 0x45, 0xda, 0x0b, 0x4d,
	0xba, 0xd3, 0xce, 0x6c, 0xf0, 0x60, 0xd7, 0x6d, 0xac, 0xba, 0x6f, 0x57, 0x98, 0xbf, 0xdb, 0xaa,
	0x37, 0x04, 0xca, 0x5a, 0xb2, 0xf5, 0x81, 0x0c, 0x9b, 0xe2, 0x5d, 0x8e, 0x9a, 0x4f, 0xb6, 0x8c,
	0xdc, 0x2d, 0xb7, 0xb2, 0xe6, 0xae, 0xac, 0xb9, 0x47, 0xc6, 0x9a, 0xbf, 0x57, 0x16, 0xf6, 0xf8,
	0x9c, 0xa2, 0x21, 0x52, 0x76, 0x4e, 0xc9, 0x5a, 0x92, 0x71, 0xb3, 0xd3, 0x73, 0x58, 0xf0, 0x3c,
	0xd6, 0x6f, 0xa2, 0x24, 0xd2, 0x93, 0x9e, 0x19, 0x61, 0xaf, 0x6d, 0xcc, 0xea, 0xe3, 0xb0, 0xb5,
	0xdb, 0xf8, 0xd3, 0xb2, 0xb0, 0x0f, 0xea, 0x68, 0x2e, 0xd2, 0xa1, 0xec, 0x62, 0x7d, 0xfa, 0xa7,
	0x4b, 0x86, 0xab, 0xe0, 0x55, 0x86, 0xa9, 0x82, 0xcd, 0xc4, 0x3a, 0x97, 0x9b, 0x58, 0xf7, 0xea,
	0x13, 0xdb, 0xba, 0xe2, 0xc4, 0x8a, 0x0e, 0x19, 0x6f, 0x26, 0xf6, 0x5e, 0xf2, 0x2c, 0x03, 0x69,
	0xdd, 0x26, 0x3d, 0x81, 0x41, 0x95, 0xd9, 0xb6, 0x7f, 0xa3, 0x2c, 0xec, 0xc1, 0x6a, 0x40, 0x00,
	0x94, 0x19, 0xd0, 0x7a, 0x44, 0x06, 0x67, 0x9f, 0x2f, 0xbe, 0x67, 0x31, 0x8f, 0x52, 0x13, 0x48,
	0xdf, 0xbf, 0x59, 0x16, 0xb6, 0xd5, 0xf6, 0xd6, 0x20, 0x65, 0xeb, 0xad, 0xd6, 0x5d, 0xb2, 0x0d,
	0x52, 0xa2, 0x34, 0x0e, 0xfb, 0xfe, 0xa8, 0x2c, 0xec, 0x9d, 0x8a, 0x63, 0xbe, 0xa6, 0xac, 0x82,
	0xad, 0xa7, 0xa4, 0x17, 0x70, 0xcd, 0xeb, 0x9f, 0xce, 0x9e, 0xdb, 0xfe, 0x45, 0x6c, 0xae, 0xbd,
	0xbe, 0xe1, 0x19, 0x81, 0x32, 0xc3, 0x9b, 0x5f, 0x33, 0x57, 0x39, 0x3c, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x77, 0x62, 0x98, 0xe8, 0x04, 0x00, 0x00,
}
