// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_list.proto

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
//资源限制
type ResourceList struct {
	//
	//cpu限制, 单位为 milli CPUs
	Cpu string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu" form:"cpu"`
	//
	//memory限制, 单位可以为MiB
	Memory               string   `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory" form:"memory"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceList) Reset()         { *m = ResourceList{} }
func (m *ResourceList) String() string { return proto.CompactTextString(m) }
func (*ResourceList) ProtoMessage()    {}
func (*ResourceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6406a261e4797f6e, []int{0}
}
func (m *ResourceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceList.Unmarshal(m, b)
}
func (m *ResourceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceList.Marshal(b, m, deterministic)
}
func (m *ResourceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceList.Merge(m, src)
}
func (m *ResourceList) XXX_Size() int {
	return xxx_messageInfo_ResourceList.Size(m)
}
func (m *ResourceList) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceList.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceList proto.InternalMessageInfo

func (m *ResourceList) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *ResourceList) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceList)(nil), "container.ResourceList")
}

func init() { proto.RegisterFile("resource_list.proto", fileDescriptor_6406a261e4797f6e) }

var fileDescriptor_6406a261e4797f6e = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x4a, 0xc6, 0x30,
	0x10, 0xc7, 0xf9, 0x14, 0x3e, 0xf8, 0x82, 0x0a, 0xd6, 0xa5, 0xb8, 0xb4, 0x64, 0xd2, 0xa1, 0xc9,
	0xe0, 0xe6, 0x58, 0x71, 0x73, 0xea, 0xa8, 0x83, 0xa4, 0x31, 0x8d, 0x81, 0xa4, 0x17, 0x2e, 0xc9,
	0xd0, 0x97, 0xed, 0x43, 0xf4, 0x09, 0xa4, 0x49, 0x71, 0xbb, 0xdf, 0xdd, 0xef, 0x8e, 0xff, 0x91,
	0x07, 0x54, 0x01, 0x12, 0x4a, 0xf5, 0x6d, 0x4d, 0x88, 0xcc, 0x23, 0x44, 0xa8, 0x2e, 0x12, 0xe6,
	0x28, 0xcc, 0xac, 0xf0, 0xb1, 0xd3, 0x26, 0xfe, 0xa6, 0x91, 0x49, 0x70, 0x5c, 0x83, 0x06, 0x9e,
	0x8d, 0x31, 0x4d, 0x99, 0x32, 0xe4, 0xaa, 0x6c, 0xd2, 0x2f, 0x72, 0x33, 0x1c, 0x07, 0x3f, 0x4c,
	0x88, 0x55, 0x4b, 0xae, 0xa5, 0x4f, 0xf5, 0xa9, 0x3d, 0x3d, 0x5d, 0xfa, 0xbb, 0x6d, 0x6d, 0xc8,
	0x04, 0xe8, 0x5e, 0xa9, 0xf4, 0x89, 0x0e, 0xfb, 0xa8, 0x7a, 0x26, 0x67, 0xa7, 0x1c, 0xe0, 0x52,
	0x5f, 0x65, 0xe9, 0x7e, 0x5b, 0x9b, 0xdb, 0x22, 0x95, 0x3e, 0x1d, 0x0e, 0xa1, 0x7f, 0xff, 0x7c,
	0xd3, 0xc0, 0x94, 0x08, 0x0b, 0xf8, 0xc0, 0x2c, 0x48, 0x61, 0xf9, 0x9e, 0x14, 0x85, 0x8c, 0xa1,
	0x04, 0x43, 0xe5, 0xa1, 0x73, 0xf0, 0xa3, 0x6c, 0xe0, 0x87, 0xc8, 0x33, 0xf2, 0xff, 0x97, 0xc6,
	0x73, 0x36, 0x5f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x51, 0xf4, 0xe0, 0xfb, 0x00, 0x00,
	0x00,
}