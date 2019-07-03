// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: register.proto

package org

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//RegisterApi返回
type RegisterResponseWrapper struct {
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

func (m *RegisterResponseWrapper) Reset()         { *m = RegisterResponseWrapper{} }
func (m *RegisterResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*RegisterResponseWrapper) ProtoMessage()    {}
func (*RegisterResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}
func (m *RegisterResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponseWrapper.Unmarshal(m, b)
}
func (m *RegisterResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponseWrapper.Marshal(b, m, deterministic)
}
func (m *RegisterResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponseWrapper.Merge(m, src)
}
func (m *RegisterResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_RegisterResponseWrapper.Size(m)
}
func (m *RegisterResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponseWrapper proto.InternalMessageInfo

func (m *RegisterResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *RegisterResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RegisterResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterResponseWrapper)(nil), "org.RegisterResponseWrapper")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xeb, 0x04, 0x53, 0x51, 0xc9, 0x61, 0x96, 0x79, 0x68, 0x89, 0x20, 0xbd, 0x98,
	0x81, 0x5e, 0x44, 0x6f, 0x83, 0x7d, 0x81, 0x5c, 0x3c, 0xb7, 0xdb, 0x5b, 0x2c, 0xac, 0x7b, 0xe1,
	0x35, 0x03, 0xfd, 0xb0, 0xf6, 0x43, 0xf4, 0x13, 0x48, 0x5f, 0x26, 0xdb, 0x29, 0x79, 0xef, 0xf7,
	0xfb, 0x93, 0xfc, 0xe5, 0x0d, 0x81, 0x6b, 0xba, 0x00, 0x64, 0x3c, 0x61, 0x40, 0x35, 0x41, 0x72,
	0xf3, 0x67, 0xd7, 0x84, 0xaf, 0x43, 0x6d, 0xd6, 0xd8, 0x2e, 0x1c, 0x3a, 0x5c, 0x30, 0xab, 0x0f,
	0x5b, 0x9e, 0x78, 0xe0, 0x5b, 0xcc, 0xcc, 0x1f, 0x1c, 0xa2, 0xdb, 0xc1, 0xc9, 0x82, 0xd6, 0x87,
	0x9f, 0x08, 0xf5, 0xaf, 0x90, 0xf7, 0xf6, 0xf8, 0x86, 0x85, 0xce, 0xe3, 0xbe, 0x83, 0x4f, 0xaa,
	0xbc, 0x07, 0x52, 0x8f, 0x32, 0x59, 0xe3, 0x06, 0x32, 0x51, 0x88, 0x72, 0xba, 0xbc, 0x1d, 0xfa,
	0x3c, 0xdd, 0x22, 0xb5, 0xef, 0x7a, 0xdc, 0x6a, 0xcb, 0x50, 0xbd, 0xc9, 0x74, 0x3c, 0x57, 0xdf,
	0x7e, 0x57, 0x35, 0xfb, 0xec, 0xa2, 0x10, 0xe5, 0xd5, 0x72, 0x36, 0xf4, 0xb9, 0x3a, 0xb9, 0x47,
	0xa8, 0xed, 0xb9, 0xaa, 0x9e, 0xe4, 0x14, 0x88, 0x90, 0xb2, 0x09, 0x67, 0xee, 0x86, 0x3e, 0xbf,
	0x8e, 0x19, 0x5e, 0x6b, 0x1b, 0xb1, 0xfa, 0x90, 0xc9, 0xa6, 0x0a, 0x55, 0x96, 0x14, 0xa2, 0x4c,
	0x5f, 0x66, 0x26, 0xd6, 0x31, 0xff, 0x75, 0xcc, 0x6a, 0xac, 0x73, 0xfe, 0xbd, 0xd1, 0xd6, 0x96,
	0x43, 0xf5, 0x25, 0x6b, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x9a, 0x37, 0x7c, 0x49,
	0x01, 0x00, 0x00,
}
