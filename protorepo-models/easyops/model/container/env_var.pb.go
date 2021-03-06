// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: env_var.proto

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
//环境变量
type EnvVar struct {
	//
	//变量名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//变量值
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvVar) Reset()         { *m = EnvVar{} }
func (m *EnvVar) String() string { return proto.CompactTextString(m) }
func (*EnvVar) ProtoMessage()    {}
func (*EnvVar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b797cd497a4454, []int{0}
}
func (m *EnvVar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvVar.Unmarshal(m, b)
}
func (m *EnvVar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvVar.Marshal(b, m, deterministic)
}
func (m *EnvVar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvVar.Merge(m, src)
}
func (m *EnvVar) XXX_Size() int {
	return xxx_messageInfo_EnvVar.Size(m)
}
func (m *EnvVar) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvVar.DiscardUnknown(m)
}

var xxx_messageInfo_EnvVar proto.InternalMessageInfo

func (m *EnvVar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EnvVar)(nil), "container.EnvVar")
}

func init() { proto.RegisterFile("env_var.proto", fileDescriptor_d1b797cd497a4454) }

var fileDescriptor_d1b797cd497a4454 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcd, 0x2b, 0x8b,
	0x2f, 0x4b, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0xcf, 0x2b, 0x49,
	0xcc, 0xcc, 0x4b, 0x2d, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x48, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30,
	0x0b, 0xa2, 0x53, 0x29, 0x94, 0x8b, 0xcd, 0x35, 0xaf, 0x2c, 0x2c, 0xb1, 0x48, 0x48, 0x99, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xff, 0xd3, 0x3d, 0x79,
	0xee, 0xb4, 0xfc, 0xa2, 0x5c, 0x2b, 0x25, 0x90, 0xa8, 0x52, 0x10, 0x58, 0x52, 0x48, 0x8d, 0x8b,
	0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0xac, 0x4a, 0xe0, 0xd3, 0x3d, 0x79, 0x1e, 0x88,
	0x2a, 0xb0, 0xb0, 0x52, 0x10, 0x44, 0xda, 0xc9, 0x35, 0xca, 0x39, 0x3d, 0x5f, 0x2f, 0x35, 0xb1,
	0xb8, 0x32, 0xbf, 0xa0, 0x58, 0x2f, 0x27, 0x3f, 0x39, 0x31, 0x47, 0x1f, 0xe4, 0xc6, 0xa2, 0xc4,
	0xe4, 0x92, 0x62, 0x88, 0x93, 0x8a, 0x52, 0x0b, 0xf2, 0x75, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x8a,
	0xf5, 0xa1, 0x0a, 0xf5, 0xc1, 0x5c, 0x7d, 0xb8, 0x67, 0x92, 0xd8, 0xc0, 0x2a, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x0a, 0x9d, 0xc7, 0xef, 0x00, 0x00, 0x00,
}
