// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: send_message_for_alert_request.proto

package msgsender

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
//告警发送通知的请求body
type SendMessageForAlertRequest struct {
	//
	//告警数据
	Data *types.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data" form:"data"`
	//
	//数据源
	System               string   `protobuf:"bytes,2,opt,name=system,proto3" json:"system" form:"system"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageForAlertRequest) Reset()         { *m = SendMessageForAlertRequest{} }
func (m *SendMessageForAlertRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageForAlertRequest) ProtoMessage()    {}
func (*SendMessageForAlertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb0c03b6034adc4, []int{0}
}
func (m *SendMessageForAlertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageForAlertRequest.Unmarshal(m, b)
}
func (m *SendMessageForAlertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageForAlertRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageForAlertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageForAlertRequest.Merge(m, src)
}
func (m *SendMessageForAlertRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageForAlertRequest.Size(m)
}
func (m *SendMessageForAlertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageForAlertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageForAlertRequest proto.InternalMessageInfo

func (m *SendMessageForAlertRequest) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SendMessageForAlertRequest) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMessageForAlertRequest)(nil), "msgsender.SendMessageForAlertRequest")
}

func init() {
	proto.RegisterFile("send_message_for_alert_request.proto", fileDescriptor_ceb0c03b6034adc4)
}

var fileDescriptor_ceb0c03b6034adc4 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xa9, 0xc8, 0x60, 0x1d, 0x22, 0xf6, 0xe2, 0x28, 0xc2, 0x46, 0xf1, 0x30, 0x0f, 0x4b,
	0x40, 0x6f, 0xe2, 0xc5, 0x8a, 0xde, 0xbc, 0x74, 0x37, 0x2f, 0x25, 0x6d, 0xbf, 0x46, 0x21, 0xe9,
	0x57, 0xbf, 0x2f, 0x3d, 0xec, 0x07, 0xf8, 0x37, 0xf7, 0x23, 0xf6, 0x0b, 0xa4, 0x49, 0x75, 0xb7,
	0xbc, 0xe1, 0x79, 0x9f, 0x37, 0x24, 0xbe, 0x65, 0xe8, 0x9a, 0xd2, 0x02, 0xb3, 0xd2, 0x50, 0xb6,
	0x48, 0xa5, 0x32, 0x40, 0xae, 0x24, 0xf8, 0x1e, 0x80, 0x9d, 0xe8, 0x09, 0x1d, 0x26, 0x73, 0xcb,
	0x7a, 0x04, 0x81, 0xd2, 0xad, 0xfe, 0x72, 0x9f, 0x43, 0x25, 0x6a, 0xb4, 0x52, 0xa3, 0x46, 0xe9,
	0x89, 0x6a, 0x68, 0x7d, 0xf2, 0xc1, 0x9f, 0x42, 0x33, 0xbd, 0xd1, 0x88, 0xda, 0xc0, 0x89, 0x62,
	0x47, 0x43, 0x3d, 0x79, 0xb3, 0x9f, 0x28, 0x4e, 0x77, 0xd0, 0x35, 0xef, 0x61, 0xff, 0x0d, 0xe9,
	0x79, 0x5c, 0x2f, 0xc2, 0x78, 0xf2, 0x14, 0x9f, 0x37, 0xca, 0xa9, 0x65, 0xb4, 0x8e, 0x36, 0x8b,
	0xfb, 0x6b, 0x11, 0x5c, 0xe2, 0xcf, 0x25, 0x76, 0xde, 0x95, 0x5f, 0x1e, 0x0f, 0xab, 0x45, 0x8b,
	0x64, 0x1f, 0xb3, 0x11, 0xcf, 0x0a, 0xdf, 0x4a, 0xee, 0xe2, 0x19, 0xef, 0xd9, 0x81, 0x5d, 0x9e,
	0xad, 0xa3, 0xcd, 0x3c, 0xbf, 0x3a, 0x1e, 0x56, 0x17, 0x01, 0x0b, 0xf7, 0x59, 0x31, 0x01, 0xf9,
	0xeb, 0xc7, 0x8b, 0x46, 0x01, 0x8a, 0xf7, 0xd8, 0xb3, 0x30, 0x58, 0x2b, 0x23, 0x6b, 0xec, 0x1c,
	0xa9, 0xda, 0x71, 0x78, 0x3b, 0x41, 0x8f, 0x5b, 0x8b, 0x0d, 0x18, 0x96, 0x13, 0x28, 0x7d, 0x94,
	0xff, 0x7f, 0x53, 0xcd, 0x3c, 0xf9, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x23, 0xa3, 0x24, 0x10,
	0x55, 0x01, 0x00, 0x00,
}