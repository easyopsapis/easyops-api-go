// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: send_message.proto

package custom_sender

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	msgsender "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/msgsender"
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
//SendMessageApi返回
type SendMessageResponseWrapper struct {
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
	Data                 *msgsender.SendMessageResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SendMessageResponseWrapper) Reset()         { *m = SendMessageResponseWrapper{} }
func (m *SendMessageResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponseWrapper) ProtoMessage()    {}
func (*SendMessageResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_be92a216bce57c56, []int{0}
}
func (m *SendMessageResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponseWrapper.Unmarshal(m, b)
}
func (m *SendMessageResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SendMessageResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponseWrapper.Merge(m, src)
}
func (m *SendMessageResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponseWrapper.Size(m)
}
func (m *SendMessageResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponseWrapper proto.InternalMessageInfo

func (m *SendMessageResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SendMessageResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SendMessageResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SendMessageResponseWrapper) GetData() *msgsender.SendMessageResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMessageResponseWrapper)(nil), "custom_sender.SendMessageResponseWrapper")
}

func init() { proto.RegisterFile("send_message.proto", fileDescriptor_be92a216bce57c56) }

var fileDescriptor_be92a216bce57c56 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xc9, 0xff, 0xdf, 0x0a, 0xa6, 0x8a, 0xb2, 0x07, 0x09, 0x41, 0x6c, 0x89, 0x20, 0xbd,
	0x34, 0x0b, 0x7a, 0x11, 0x8f, 0x15, 0x8f, 0x5e, 0xe2, 0xc1, 0x63, 0xd8, 0x6e, 0xa6, 0x6b, 0x21,
	0xc9, 0xac, 0x33, 0x1b, 0xd0, 0x2f, 0xdb, 0xb3, 0xe7, 0x7e, 0x02, 0xc9, 0x6e, 0xb0, 0x3d, 0x78,
	0xf5, 0xb4, 0x3b, 0xf3, 0xde, 0x8f, 0x99, 0x37, 0xb1, 0x60, 0x68, 0xab, 0xb2, 0x01, 0x66, 0x65,
	0x20, 0xb7, 0x84, 0x0e, 0xc5, 0xa9, 0xee, 0xd8, 0x61, 0x53, 0xf6, 0x12, 0x50, 0xba, 0x30, 0x1b,
	0xf7, 0xd6, 0xad, 0x72, 0x8d, 0x8d, 0x34, 0x68, 0x50, 0x7a, 0xd7, 0xaa, 0x5b, 0xfb, 0xca, 0x17,
	0xfe, 0x17, 0xe8, 0x54, 0x1b, 0xcc, 0x41, 0xf1, 0x27, 0x5a, 0xce, 0x6b, 0xd4, 0xaa, 0x96, 0x1a,
	0x5b, 0x47, 0x4a, 0x3b, 0x0e, 0x24, 0x81, 0xc5, 0x45, 0x83, 0x15, 0xd4, 0x2c, 0x07, 0xa3, 0xf4,
	0xa5, 0x6c, 0xd8, 0x84, 0x99, 0xf2, 0x70, 0xab, 0x92, 0xe0, 0xbd, 0x03, 0x76, 0x65, 0xa5, 0x9c,
	0x1a, 0x86, 0x94, 0x7f, 0x30, 0x84, 0x2d, 0xb6, 0x3c, 0xdc, 0x20, 0xbd, 0x34, 0x88, 0xa6, 0x86,
	0x7d, 0x56, 0x76, 0xd4, 0x69, 0x17, 0xd4, 0xec, 0x2b, 0x8a, 0xd3, 0x17, 0x68, 0xab, 0xe7, 0x00,
	0x17, 0x03, 0xfb, 0x4a, 0xca, 0x5a, 0x20, 0x71, 0x1d, 0x8f, 0x34, 0x56, 0x90, 0x44, 0xb3, 0x68,
	0x3e, 0x5e, 0x9e, 0xed, 0xb6, 0xd3, 0xc9, 0x1a, 0xa9, 0x79, 0xc8, 0xfa, 0x6e, 0x56, 0x78, 0x51,
	0xdc, 0xc7, 0x93, 0xfe, 0x7d, 0xfa, 0xb0, 0xb5, 0xda, 0xb4, 0xc9, 0xbf, 0x59, 0x34, 0x3f, 0x5e,
	0x5e, 0xec, 0xb6, 0x53, 0xb1, 0xf7, 0x0e, 0x62, 0x56, 0x1c, 0x5a, 0xc5, 0x4d, 0x3c, 0x06, 0x22,
	0xa4, 0xe4, 0xbf, 0x67, 0xce, 0x77, 0xdb, 0xe9, 0x49, 0x60, 0x7c, 0x3b, 0x2b, 0x82, 0x2c, 0x1e,
	0xe3, 0x51, 0x7f, 0xb2, 0x64, 0x34, 0x8b, 0xe6, 0x93, 0xdb, 0xab, 0xfc, 0x27, 0x79, 0xfe, 0xcb,
	0xee, 0x87, 0x6b, 0xf6, 0x54, 0x56, 0x78, 0x78, 0x75, 0xe4, 0x13, 0xdf, 0x7d, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x1f, 0x39, 0x73, 0x29, 0x02, 0x00, 0x00,
}
