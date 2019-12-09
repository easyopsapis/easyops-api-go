// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_requirements.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//ResourceRequirements 资源定义
type ResourceRequirements struct {
	//
	//最大资源限制
	Limits *ResourceList `protobuf:"bytes,1,opt,name=limits,proto3" json:"limits" form:"limits"`
	//
	//容器启动的初始资源请求
	Requests             *ResourceList `protobuf:"bytes,2,opt,name=requests,proto3" json:"requests" form:"requests"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_e065d4cf4a9d0704, []int{0}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (m *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(m, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetLimits() *ResourceList {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *ResourceRequirements) GetRequests() *ResourceList {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceRequirements)(nil), "container.ResourceRequirements")
}

func init() { proto.RegisterFile("resource_requirements.proto", fileDescriptor_e065d4cf4a9d0704) }

var fileDescriptor_e065d4cf4a9d0704 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x89, 0xc5, 0xa1, 0x11, 0x11, 0xa3, 0xe0, 0x71, 0x16, 0x27, 0xa9, 0x6c, 0x6e, 0x17,
	0xb4, 0xb3, 0x8c, 0x08, 0x16, 0x56, 0x69, 0x04, 0x1b, 0x49, 0xd6, 0xb9, 0xb8, 0xb0, 0x9b, 0x89,
	0x33, 0x93, 0xc2, 0xf7, 0xf1, 0xb9, 0xee, 0x21, 0xee, 0x09, 0xe4, 0x76, 0x73, 0x8b, 0x95, 0x95,
	0xdd, 0xfe, 0xcb, 0xff, 0x7d, 0x33, 0x4c, 0x7e, 0x45, 0xc0, 0x38, 0x92, 0x81, 0x37, 0x82, 0xcf,
	0xd1, 0x12, 0x78, 0xe8, 0x85, 0xd5, 0x40, 0x28, 0x58, 0x1c, 0x19, 0xec, 0xa5, 0xb1, 0x3d, 0xd0,
	0x62, 0xd5, 0x59, 0xf9, 0x18, 0x5b, 0x65, 0xd0, 0xeb, 0x0e, 0x3b, 0xd4, 0xa1, 0xd1, 0x8e, 0xeb,
	0x90, 0x42, 0x08, 0xaf, 0x48, 0x2e, 0x5e, 0x3a, 0x54, 0xd0, 0xf0, 0x17, 0x0e, 0xac, 0x1c, 0x9a,
	0xc6, 0xe9, 0x9d, 0x8a, 0x1a, 0x23, 0x1c, 0x49, 0x82, 0x01, 0x57, 0x1e, 0xdf, 0xc1, 0xb1, 0x9e,
	0x8a, 0x3a, 0x44, 0x9d, 0x66, 0xea, 0xb4, 0x9a, 0xb3, 0x2c, 0x51, 0x5c, 0x7e, 0x67, 0xf9, 0x45,
	0x3d, 0xfd, 0xd7, 0xbf, 0x36, 0x2e, 0xaa, 0x7c, 0xe6, 0xac, 0xb7, 0xc2, 0xf3, 0xec, 0x3a, 0xbb,
	0x39, 0xbe, 0xbd, 0x54, 0x49, 0xa4, 0xf6, 0xc0, 0xb3, 0x65, 0xa9, 0xce, 0xb6, 0x9b, 0xe5, 0xc9,
	0x1a, 0xc9, 0xdf, 0x97, 0x11, 0x28, 0xeb, 0x89, 0x2c, 0x9e, 0xf2, 0xc3, 0xdd, 0x15, 0x80, 0x85,
	0xe7, 0x07, 0x7f, 0x5b, 0xce, 0xb7, 0x9b, 0xe5, 0x69, 0xb4, 0xec, 0x91, 0xb2, 0x4e, 0x74, 0xf5,
	0xf8, 0xfa, 0xf0, 0x0f, 0x17, 0x68, 0x67, 0xa1, 0x79, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0x13, 0x0d, 0x0c, 0xa6, 0x01, 0x00, 0x00,
}