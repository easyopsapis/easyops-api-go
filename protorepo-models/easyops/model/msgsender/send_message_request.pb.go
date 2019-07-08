// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: send_message_request.proto

package msgsender

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//发送通知的请求body
type SendMessageRequest struct {
	//
	//发送通知的请求数据
	Data                 *SendMessageRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29d7446efc63e94d, []int{0}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetData() *SendMessageRequestData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "msgsender.SendMessageRequest")
}

func init() { proto.RegisterFile("send_message_request.proto", fileDescriptor_29d7446efc63e94d) }

var fileDescriptor_29d7446efc63e94d = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0xcd, 0x4b,
	0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x2d, 0x4e, 0x07, 0x49, 0xa7, 0x16, 0x49,
	0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0x55, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x29, 0x95, 0x9c,
	0x9e, 0xaf, 0x97, 0x9a, 0x58, 0x5c, 0x99, 0x5f, 0x50, 0xac, 0x97, 0x93, 0x9f, 0x9c, 0x98, 0xa3,
	0x9f, 0x9c, 0x9f, 0x57, 0x52, 0x94, 0x98, 0x5c, 0x52, 0x0c, 0xd1, 0x59, 0x94, 0x5a, 0x90, 0xaf,
	0x9b, 0x9b, 0x9f, 0x92, 0x9a, 0x53, 0xac, 0x0f, 0x55, 0xa8, 0x0f, 0xe6, 0xea, 0xc3, 0xed, 0xd4,
	0xc7, 0xe6, 0xb2, 0xf8, 0x94, 0xc4, 0x92, 0x44, 0x88, 0x25, 0x4a, 0x31, 0x5c, 0x42, 0xc1, 0xa9,
	0x79, 0x29, 0xbe, 0x10, 0x15, 0x41, 0x10, 0x05, 0x42, 0x6e, 0x5c, 0x2c, 0x20, 0x35, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x8a, 0x7a, 0x70, 0xf3, 0xf4, 0x30, 0x15, 0xbb, 0x24, 0x96, 0x24,
	0x3a, 0xf1, 0x7f, 0xba, 0x27, 0xcf, 0x9d, 0x96, 0x5f, 0x94, 0x6b, 0xa5, 0x04, 0xd2, 0xa8, 0x14,
	0x04, 0xd6, 0xef, 0xe4, 0x1a, 0xe5, 0x4c, 0x05, 0x4f, 0x24, 0xb1, 0x81, 0x55, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x0b, 0x80, 0xda, 0x68, 0x01, 0x00, 0x00,
}
