// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: push_message.proto

package datafilter

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
//PushMessageApi返回
type PushMessageResponseWrapper struct {
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

func (m *PushMessageResponseWrapper) Reset()         { *m = PushMessageResponseWrapper{} }
func (m *PushMessageResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*PushMessageResponseWrapper) ProtoMessage()    {}
func (*PushMessageResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2677d527632bbd7f, []int{0}
}
func (m *PushMessageResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMessageResponseWrapper.Unmarshal(m, b)
}
func (m *PushMessageResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMessageResponseWrapper.Marshal(b, m, deterministic)
}
func (m *PushMessageResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMessageResponseWrapper.Merge(m, src)
}
func (m *PushMessageResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_PushMessageResponseWrapper.Size(m)
}
func (m *PushMessageResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMessageResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_PushMessageResponseWrapper proto.InternalMessageInfo

func (m *PushMessageResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PushMessageResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *PushMessageResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PushMessageResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PushMessageResponseWrapper)(nil), "datafilter.PushMessageResponseWrapper")
}

func init() { proto.RegisterFile("push_message.proto", fileDescriptor_2677d527632bbd7f) }

var fileDescriptor_2677d527632bbd7f = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0x6e, 0x82, 0xa9, 0xa0, 0xe4, 0x30, 0x4a, 0x3d, 0xb4, 0x44, 0x90, 0x5e, 0xcc,
	0x40, 0x2f, 0xa2, 0xb7, 0xc1, 0x8e, 0x82, 0xe4, 0xe2, 0x51, 0xd2, 0xed, 0x35, 0x2d, 0xb4, 0x7b,
	0x21, 0x49, 0x41, 0xbf, 0x6c, 0xf1, 0x33, 0xf4, 0x13, 0x48, 0x12, 0x65, 0x3b, 0x25, 0xef, 0xfd,
	0x7f, 0x7f, 0x92, 0x1f, 0xa1, 0x7a, 0xb4, 0xed, 0xe7, 0x00, 0xd6, 0x4a, 0x05, 0x5c, 0x1b, 0x74,
	0x48, 0xc9, 0x5e, 0x3a, 0xd9, 0x74, 0xbd, 0x03, 0x93, 0x3f, 0xa8, 0xce, 0xb5, 0x63, 0xcd, 0x77,
	0x38, 0xac, 0x15, 0x2a, 0x5c, 0x07, 0xa4, 0x1e, 0x9b, 0x30, 0x85, 0x21, 0xdc, 0x62, 0x35, 0xbf,
	0x55, 0x88, 0xaa, 0x87, 0x23, 0x05, 0x83, 0x76, 0xdf, 0x31, 0x64, 0x3f, 0x09, 0xc9, 0xdf, 0x47,
	0xdb, 0xbe, 0xc5, 0xd7, 0x04, 0x58, 0x8d, 0x07, 0x0b, 0x1f, 0x46, 0x6a, 0x0d, 0x86, 0xde, 0x91,
	0xc5, 0x0e, 0xf7, 0x90, 0x25, 0x65, 0x52, 0x2d, 0x37, 0xd7, 0xf3, 0x54, 0xa4, 0x0d, 0x9a, 0xe1,
	0x85, 0xf9, 0x2d, 0x13, 0x21, 0xa4, 0xcf, 0x24, 0xf5, 0xe7, 0xf6, 0x4b, 0xf7, 0xb2, 0x3b, 0x64,
	0x67, 0x65, 0x52, 0x5d, 0x6e, 0x56, 0xf3, 0x54, 0xd0, 0x23, 0xfb, 0x17, 0x32, 0x71, 0x8a, 0xd2,
	0x7b, 0xb2, 0x04, 0x63, 0xd0, 0x64, 0xe7, 0xa1, 0x73, 0x33, 0x4f, 0xc5, 0x55, 0xec, 0x84, 0x35,
	0x13, 0x31, 0xa6, 0xaf, 0x64, 0xe1, 0xfd, 0xb3, 0x45, 0x99, 0x54, 0xe9, 0xe3, 0x8a, 0x47, 0x23,
	0xfe, 0x6f, 0xc4, 0xb7, 0xde, 0xe8, 0xf4, 0x7b, 0x9e, 0x66, 0x22, 0x94, 0xea, 0x8b, 0x80, 0x3d,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x77, 0x6a, 0xe7, 0x31, 0x57, 0x01, 0x00, 0x00,
}
