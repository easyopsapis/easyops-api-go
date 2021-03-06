// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: namespace.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//Namespace 资源定义
type Namespace struct {
	//
	//namespace id，服务端自动生成
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//namespace 类型, 只能是 Namespace
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind" form:"kind"`
	//
	//namespace 全称，命名规则 clusterId:kind:resourceName, 创建之后不能修改
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//名称
	ResourceName string `protobuf:"bytes,4,opt,name=resourceName,proto3" json:"resourceName" form:"resourceName"`
	//
	//存放 resourceQuota 资源定义
	ResourceSpec string `protobuf:"bytes,5,opt,name=resourceSpec,proto3" json:"resourceSpec" form:"resourceSpec"`
	//
	//namespace 状态
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status" form:"status"`
	//
	//namespace 描述
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description" form:"description"`
	//
	//是否是系统命名空间
	System bool `protobuf:"varint,8,opt,name=system,proto3" json:"system" form:"system"`
	//
	//创建者, 服务端自动生成
	Creator              string   `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator" form:"creator"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{0}
}
func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Namespace) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Namespace) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Namespace) GetResourceSpec() string {
	if m != nil {
		return m.ResourceSpec
	}
	return ""
}

func (m *Namespace) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Namespace) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Namespace) GetSystem() bool {
	if m != nil {
		return m.System
	}
	return false
}

func (m *Namespace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Namespace)(nil), "container.Namespace")
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor_ecb1e126f615f5dd) }

var fileDescriptor_ecb1e126f615f5dd = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x55, 0x28, 0xdd, 0xea, 0xc1, 0xc6, 0x8c, 0x04, 0xd6, 0x5e, 0xb2, 0x65, 0x7d, 0x68,
	0x11, 0x4e, 0x36, 0x2a, 0x21, 0x0a, 0x48, 0x88, 0x21, 0x1e, 0xf6, 0xc2, 0x43, 0xf6, 0xb6, 0xaa,
	0x03, 0xd7, 0xf1, 0x42, 0xb4, 0x26, 0x8e, 0x6c, 0x97, 0x69, 0x9b, 0xf6, 0x6d, 0xf8, 0x5c, 0x41,
	0xe2, 0x23, 0xe4, 0x13, 0x20, 0x9f, 0xd3, 0x36, 0x85, 0xb7, 0xf3, 0xfd, 0x7f, 0x77, 0xff, 0xd3,
	0xf9, 0xd0, 0x4e, 0xce, 0x32, 0xa1, 0x0b, 0xc6, 0x45, 0x50, 0x28, 0x69, 0x24, 0xee, 0x72, 0x99,
	0x1b, 0x96, 0xe6, 0x42, 0xed, 0xd1, 0x24, 0x35, 0x3f, 0xe6, 0xd3, 0x80, 0xcb, 0x2c, 0x4c, 0x64,
	0x22, 0x43, 0x20, 0xa6, 0xf3, 0x4b, 0x78, 0xc1, 0x03, 0x22, 0x57, 0xb9, 0xf7, 0xa6, 0x81, 0x67,
	0xd7, 0xa9, 0xb9, 0x92, 0xd7, 0x61, 0x22, 0x29, 0x88, 0xf4, 0x27, 0x9b, 0xa5, 0x31, 0x33, 0x52,
	0xe9, 0x70, 0x19, 0xba, 0x3a, 0xff, 0x57, 0x1b, 0x75, 0xbf, 0x2e, 0xa6, 0xc0, 0xa7, 0x08, 0xa5,
	0xb9, 0x36, 0x2c, 0xe7, 0xe2, 0x34, 0x26, 0xad, 0xfd, 0x56, 0xbf, 0x7b, 0x32, 0xa8, 0x4a, 0x6f,
	0xf7, 0x52, 0xaa, 0xec, 0x9d, 0xbf, 0xd2, 0xfc, 0x3f, 0xbf, 0xbd, 0xa7, 0x68, 0xfb, 0x62, 0x7c,
	0x44, 0x47, 0x8c, 0xde, 0x4e, 0xee, 0x8e, 0x87, 0xf7, 0xbd, 0xa8, 0x51, 0x8c, 0x0f, 0x51, 0xfb,
	0x2a, 0xcd, 0x63, 0xf2, 0x00, 0x9a, 0xec, 0x54, 0xa5, 0xb7, 0xe5, 0x9a, 0xd8, 0xac, 0x1f, 0x81,
	0x68, 0x21, 0xbb, 0x02, 0xf2, 0xf0, 0x5f, 0xc8, 0x66, 0xfd, 0x08, 0x44, 0xfc, 0x1d, 0x3d, 0x56,
	0x42, 0xcb, 0xb9, 0xe2, 0xc2, 0x4e, 0x4a, 0xda, 0x00, 0x7f, 0xa8, 0x4a, 0xef, 0x99, 0x83, 0x9b,
	0xaa, 0x1d, 0xec, 0x00, 0x79, 0x17, 0x63, 0x46, 0x6f, 0x8f, 0xe8, 0x68, 0xd2, 0x1f, 0xd3, 0x3a,
	0x7a, 0xb9, 0x48, 0x0d, 0x3e, 0xf6, 0xa2, 0xb5, 0x8e, 0xf8, 0xfd, 0xca, 0xe1, 0xac, 0x10, 0x9c,
	0x3c, 0x02, 0x87, 0x17, 0xff, 0x3b, 0x58, 0xd5, 0x8f, 0xd6, 0x60, 0x3c, 0x40, 0x1d, 0x6d, 0x98,
	0x99, 0x6b, 0xd2, 0x81, 0xb2, 0xdd, 0xaa, 0xf4, 0x9e, 0xb8, 0x32, 0x97, 0xf7, 0xa3, 0x1a, 0xc0,
	0x6f, 0xd1, 0x56, 0x2c, 0x34, 0x57, 0x69, 0x61, 0x52, 0x99, 0x93, 0x0d, 0xe0, 0x9f, 0x57, 0xa5,
	0x87, 0x1d, 0xdf, 0x10, 0xfd, 0xa8, 0x89, 0x82, 0xc9, 0x8d, 0x36, 0x22, 0x23, 0x9b, 0xfb, 0xad,
	0xfe, 0xe6, 0x9a, 0x09, 0xe4, 0xad, 0x09, 0x04, 0xf8, 0x0c, 0x6d, 0x70, 0x25, 0xec, 0x17, 0x93,
	0x2e, 0x18, 0x8c, 0xaa, 0xd2, 0xdb, 0x76, 0x6c, 0x2d, 0xd8, 0x25, 0x1d, 0xa2, 0x03, 0x58, 0xd2,
	0x27, 0x7a, 0x6e, 0x97, 0x32, 0x0e, 0x96, 0xf1, 0x37, 0x3a, 0xb9, 0x7b, 0xfd, 0x6a, 0x78, 0x7c,
	0xdf, 0x8b, 0x16, 0x9d, 0x4e, 0xbe, 0x9c, 0x7f, 0x4e, 0x64, 0x20, 0x98, 0xbe, 0x91, 0x85, 0x0e,
	0x66, 0x92, 0xb3, 0x59, 0x68, 0x6f, 0x55, 0x31, 0x6e, 0xb4, 0x3b, 0x4d, 0x25, 0x0a, 0x49, 0x33,
	0x19, 0x8b, 0x99, 0x0e, 0x6b, 0x30, 0x84, 0x67, 0xb8, 0x3c, 0xea, 0x69, 0x07, 0xc8, 0xe1, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x26, 0x39, 0x80, 0xf9, 0x02, 0x00, 0x00,
}
