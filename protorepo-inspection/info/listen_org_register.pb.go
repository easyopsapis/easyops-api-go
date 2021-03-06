// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: listen_org_register.proto

package info

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//ListenOrgRegisterApi返回
type ListenOrgRegisterResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListenOrgRegisterResponseWrapper) Reset()         { *m = ListenOrgRegisterResponseWrapper{} }
func (m *ListenOrgRegisterResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListenOrgRegisterResponseWrapper) ProtoMessage()    {}
func (*ListenOrgRegisterResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_94f45d717a0f5061, []int{0}
}
func (m *ListenOrgRegisterResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenOrgRegisterResponseWrapper.Unmarshal(m, b)
}
func (m *ListenOrgRegisterResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenOrgRegisterResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListenOrgRegisterResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenOrgRegisterResponseWrapper.Merge(m, src)
}
func (m *ListenOrgRegisterResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListenOrgRegisterResponseWrapper.Size(m)
}
func (m *ListenOrgRegisterResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenOrgRegisterResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListenOrgRegisterResponseWrapper proto.InternalMessageInfo

func (m *ListenOrgRegisterResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListenOrgRegisterResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListenOrgRegisterResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListenOrgRegisterResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenOrgRegisterResponseWrapper)(nil), "info.ListenOrgRegisterResponseWrapper")
}

func init() { proto.RegisterFile("listen_org_register.proto", fileDescriptor_94f45d717a0f5061) }

var fileDescriptor_94f45d717a0f5061 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xa6, 0x82, 0x1b, 0x41, 0xd9, 0x43, 0x89, 0xf5, 0x90, 0xb0, 0x82, 0xe4, 0xe2,
	0x16, 0xf4, 0x22, 0x7a, 0x2b, 0xf4, 0x26, 0x08, 0x7b, 0xf1, 0x58, 0x92, 0x76, 0xb3, 0x06, 0x92,
	0xcc, 0x32, 0xd9, 0x82, 0xfe, 0xd9, 0x1c, 0xfd, 0x01, 0xf9, 0x05, 0xb2, 0xb3, 0x95, 0xf6, 0xb4,
	0xfb, 0xe6, 0x7d, 0x8f, 0x99, 0xc7, 0x6e, 0xdb, 0x66, 0x70, 0xba, 0xdf, 0x00, 0x9a, 0x0d, 0x6a,
	0xe3, 0x05, 0x4a, 0x8b, 0xe0, 0x80, 0xc7, 0x4d, 0x5f, 0xc3, 0xe2, 0xd1, 0x34, 0xee, 0x6b, 0x5f,
	0xc9, 0x2d, 0x74, 0x4b, 0x03, 0x06, 0x96, 0x64, 0x56, 0xfb, 0x9a, 0x14, 0x09, 0xfa, 0x85, 0xd0,
	0xe2, 0xce, 0x00, 0x98, 0x56, 0x1f, 0x29, 0xdd, 0x59, 0xf7, 0x13, 0x4c, 0xf1, 0x1b, 0xb1, 0xfc,
	0x9d, 0xf6, 0x7d, 0xa0, 0x51, 0x87, 0x6d, 0x4a, 0x0f, 0x16, 0xfa, 0x41, 0x7f, 0x62, 0x69, 0xad,
	0x46, 0x7e, 0xcf, 0xe2, 0x2d, 0xec, 0x74, 0x1a, 0xe5, 0x51, 0x31, 0x5b, 0x5d, 0x4f, 0x63, 0x96,
	0xd4, 0x80, 0xdd, 0xab, 0xf0, 0x53, 0xa1, 0xc8, 0xe4, 0x2f, 0x2c, 0xf1, 0xef, 0xfa, 0xdb, 0xb6,
	0x65, 0xd3, 0xa7, 0x67, 0x79, 0x54, 0x5c, 0xae, 0xe6, 0xd3, 0x98, 0xf1, 0x23, 0x7b, 0x30, 0x85,
	0x3a, 0x45, 0xf9, 0x03, 0x9b, 0x69, 0x44, 0xc0, 0xf4, 0x9c, 0x32, 0x37, 0xd3, 0x98, 0x5d, 0x85,
	0x0c, 0x8d, 0x85, 0x0a, 0x36, 0x7f, 0x63, 0xf1, 0xae, 0x74, 0x65, 0x1a, 0xe7, 0x51, 0x91, 0x3c,
	0xcd, 0x65, 0xe8, 0x25, 0xff, 0x7b, 0xc9, 0xb5, 0xef, 0x75, 0x7a, 0x9e, 0xa7, 0x85, 0xa2, 0x50,
	0x75, 0x41, 0xd8, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb2, 0x41, 0xf5, 0x5e, 0x01,
	0x00, 0x00,
}
