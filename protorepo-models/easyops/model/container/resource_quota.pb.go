// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_quota.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
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
//ResourceQuota 资源定义
type ResourceQuota struct {
	//
	//id，服务端自动生成
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//ResourceQuota 类型
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind" form:"kind"`
	//
	//resourcequota 全称，命名规则 clusterId:kind:name, 创建之后不能修改
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//名称
	ResourceName string `protobuf:"bytes,4,opt,name=resourceName,proto3" json:"resourceName" form:"resourceName"`
	//
	//quota 资源定义文件
	ResourceSpec string `protobuf:"bytes,5,opt,name=resourceSpec,proto3" json:"resourceSpec" form:"resourceSpec"`
	//
	//命名空间, 创建之后不能修改
	Namespace string `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace" form:"namespace"`
	//
	//资源限制
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
	Hard *types.Struct `protobuf:"bytes,7,opt,name=hard,proto3" json:"hard" form:"hard"`
	//
	//已经使用的资源
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
	Used *types.Struct `protobuf:"bytes,8,opt,name=used,proto3" json:"used" form:"used"`
	//
	//创建者, 服务端自动生成
	Creator              string   `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator" form:"creator"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceQuota) Reset()         { *m = ResourceQuota{} }
func (m *ResourceQuota) String() string { return proto.CompactTextString(m) }
func (*ResourceQuota) ProtoMessage()    {}
func (*ResourceQuota) Descriptor() ([]byte, []int) {
	return fileDescriptor_f89d404064b740ce, []int{0}
}
func (m *ResourceQuota) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceQuota.Unmarshal(m, b)
}
func (m *ResourceQuota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceQuota.Marshal(b, m, deterministic)
}
func (m *ResourceQuota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceQuota.Merge(m, src)
}
func (m *ResourceQuota) XXX_Size() int {
	return xxx_messageInfo_ResourceQuota.Size(m)
}
func (m *ResourceQuota) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceQuota.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceQuota proto.InternalMessageInfo

func (m *ResourceQuota) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *ResourceQuota) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ResourceQuota) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceQuota) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ResourceQuota) GetResourceSpec() string {
	if m != nil {
		return m.ResourceSpec
	}
	return ""
}

func (m *ResourceQuota) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ResourceQuota) GetHard() *types.Struct {
	if m != nil {
		return m.Hard
	}
	return nil
}

func (m *ResourceQuota) GetUsed() *types.Struct {
	if m != nil {
		return m.Used
	}
	return nil
}

func (m *ResourceQuota) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceQuota)(nil), "container.ResourceQuota")
}

func init() { proto.RegisterFile("resource_quota.proto", fileDescriptor_f89d404064b740ce) }

var fileDescriptor_f89d404064b740ce = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x55, 0x28, 0x1b, 0xf5, 0x60, 0x8c, 0x30, 0x69, 0xd1, 0x84, 0x94, 0x2d, 0xeb, 0x43,
	0x87, 0x70, 0xb2, 0xb5, 0x12, 0x52, 0x61, 0x12, 0x62, 0x88, 0x87, 0xbd, 0x20, 0xe1, 0xbe, 0xad,
	0xea, 0x86, 0xeb, 0x78, 0x59, 0xb4, 0x36, 0x57, 0x6c, 0x87, 0x89, 0x4d, 0xfb, 0x7e, 0x7c, 0x8a,
	0x20, 0xf1, 0x11, 0xf2, 0x09, 0x90, 0x2f, 0x69, 0xd3, 0xc1, 0x0b, 0x6f, 0x67, 0xff, 0x7f, 0xff,
	0xf3, 0xe5, 0xee, 0x42, 0x36, 0x95, 0xd4, 0x90, 0x29, 0x21, 0xcf, 0xbf, 0x65, 0x60, 0x78, 0x30,
	0x53, 0x60, 0xc0, 0x69, 0x09, 0x48, 0x0d, 0x4f, 0x52, 0xa9, 0xb6, 0x69, 0x9c, 0x98, 0xcb, 0x6c,
	0x1c, 0x08, 0x98, 0x86, 0x31, 0xc4, 0x10, 0x22, 0x31, 0xce, 0x2e, 0xf0, 0x84, 0x07, 0x8c, 0x4a,
	0xe7, 0xf6, 0x9b, 0x25, 0x7c, 0x7a, 0x9d, 0x98, 0x2b, 0xb8, 0x0e, 0x63, 0xa0, 0x28, 0xd2, 0xef,
	0x7c, 0x92, 0x44, 0xdc, 0x80, 0xd2, 0xe1, 0x22, 0xac, 0x7c, 0x2f, 0x63, 0x80, 0x78, 0x22, 0xeb,
	0xec, 0xda, 0xa8, 0x4c, 0x98, 0x52, 0xf5, 0x7f, 0x36, 0xc9, 0x53, 0x56, 0x15, 0xfa, 0xc5, 0xd6,
	0xe9, 0x9c, 0x10, 0x92, 0xa4, 0xda, 0xf0, 0x54, 0xc8, 0x93, 0xc8, 0x6d, 0xec, 0x34, 0x3a, 0xad,
	0xe3, 0xfd, 0x22, 0xf7, 0x9e, 0x5f, 0x80, 0x9a, 0xbe, 0xf5, 0x6b, 0xcd, 0xff, 0xfd, 0xcb, 0xdb,
	0x20, 0xeb, 0x67, 0xc3, 0x03, 0xda, 0xe7, 0xf4, 0x66, 0x74, 0x7b, 0xd8, 0xbb, 0x6b, 0xb3, 0x25,
	0xb3, 0xb3, 0x47, 0x9a, 0x57, 0x49, 0x1a, 0xb9, 0x0f, 0x30, 0xc9, 0xb3, 0x22, 0xf7, 0xd6, 0xca,
	0x24, 0xf6, 0xd6, 0x67, 0x28, 0x5a, 0x28, 0xe5, 0x53, 0xe9, 0x3e, 0xfc, 0x1b, 0xb2, 0xb7, 0x3e,
	0x43, 0xd1, 0xf9, 0x4a, 0x9e, 0xcc, 0xdb, 0xf9, 0xd9, 0xc2, 0x4d, 0x84, 0x8f, 0x8a, 0xdc, 0x7b,
	0x51, 0xc2, 0xcb, 0xaa, 0x2d, 0x6c, 0x97, 0x78, 0x67, 0x43, 0x4e, 0x6f, 0x0e, 0x68, 0x7f, 0xd4,
	0x19, 0xd2, 0x2a, 0x7a, 0x35, 0xbf, 0xda, 0x7f, 0xdf, 0x66, 0xf7, 0x32, 0x3a, 0xef, 0xea, 0x17,
	0x06, 0x33, 0x29, 0xdc, 0x47, 0xf8, 0xc2, 0xd6, 0xbf, 0x2f, 0x58, 0xd5, 0x67, 0xf7, 0x60, 0xa7,
	0x4b, 0x5a, 0xb6, 0x4c, 0x3d, 0xe3, 0x42, 0xba, 0x2b, 0xe8, 0xdc, 0x2c, 0x72, 0x6f, 0xa3, 0xfe,
	0x10, 0x94, 0x7c, 0x56, 0x63, 0xce, 0x11, 0x69, 0x5e, 0x72, 0x15, 0xb9, 0xab, 0x3b, 0x8d, 0xce,
	0x5a, 0x77, 0x2b, 0x28, 0xc7, 0x14, 0xcc, 0xc7, 0x14, 0x0c, 0x70, 0x4c, 0xcb, 0x0d, 0xb1, 0xb8,
	0xcf, 0xd0, 0x65, 0xdd, 0x99, 0x96, 0x91, 0xfb, 0xf8, 0xbf, 0xdd, 0x16, 0xf7, 0x19, 0xba, 0x9c,
	0x01, 0x59, 0x15, 0x4a, 0xda, 0x25, 0x71, 0x5b, 0x58, 0x6d, 0xbf, 0xc8, 0xbd, 0xf5, 0x92, 0xab,
	0x04, 0xdb, 0xc4, 0x3d, 0xb2, 0x8b, 0x4d, 0xfc, 0x40, 0x4f, 0x6d, 0xd3, 0x86, 0xc1, 0x22, 0x3e,
	0xa7, 0xa3, 0xdb, 0xee, 0xeb, 0xde, 0xe1, 0x5d, 0x9b, 0xcd, 0x33, 0x1d, 0x7f, 0x3a, 0xfd, 0x18,
	0x43, 0x20, 0xb9, 0xfe, 0x01, 0x33, 0x1d, 0x4c, 0x40, 0xf0, 0x49, 0x68, 0xb7, 0x5d, 0x71, 0x61,
	0x74, 0xb9, 0x7e, 0x4a, 0xce, 0x80, 0x4e, 0x21, 0x92, 0x13, 0x1d, 0x56, 0x60, 0x88, 0xc7, 0x70,
	0xf1, 0x5b, 0x8c, 0x57, 0x90, 0xec, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xf1, 0x3a, 0xbd,
	0x40, 0x03, 0x00, 0x00,
}
