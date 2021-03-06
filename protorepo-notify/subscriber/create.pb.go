// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package subscriber

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/notify"
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
//CreateSubscriberApi返回
type CreateSubscriberResponseWrapper struct {
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

func (m *CreateSubscriberResponseWrapper) Reset()         { *m = CreateSubscriberResponseWrapper{} }
func (m *CreateSubscriberResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateSubscriberResponseWrapper) ProtoMessage()    {}
func (*CreateSubscriberResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateSubscriberResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSubscriberResponseWrapper.Unmarshal(m, b)
}
func (m *CreateSubscriberResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSubscriberResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateSubscriberResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSubscriberResponseWrapper.Merge(m, src)
}
func (m *CreateSubscriberResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateSubscriberResponseWrapper.Size(m)
}
func (m *CreateSubscriberResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSubscriberResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSubscriberResponseWrapper proto.InternalMessageInfo

func (m *CreateSubscriberResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateSubscriberResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateSubscriberResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateSubscriberResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSubscriberResponseWrapper)(nil), "subscriber.CreateSubscriberResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xa9, 0x6e, 0x82, 0xd9, 0x40, 0xe9, 0xc5, 0x28, 0xf3, 0xa2, 0xa3, 0x82, 0xec, 0xa6,
	0x09, 0x28, 0x88, 0xe8, 0xdd, 0x64, 0x2f, 0x50, 0x11, 0x2f, 0x25, 0x4d, 0xd3, 0x18, 0x6c, 0x7b,
	0xc2, 0x49, 0xe6, 0xdc, 0xcb, 0xee, 0xce, 0x17, 0xd8, 0x13, 0x48, 0xd3, 0xfd, 0xe9, 0x55, 0xcf,
	0x39, 0xdf, 0xf7, 0x6b, 0xbe, 0x8f, 0x8c, 0x05, 0x4a, 0xee, 0x24, 0x35, 0x08, 0x0e, 0x42, 0x62,
	0x57, 0xb9, 0x15, 0xa8, 0x73, 0x89, 0xd3, 0x54, 0x69, 0xf7, 0xb5, 0xca, 0xa9, 0x80, 0x9a, 0x29,
	0x50, 0xc0, 0xbc, 0x25, 0x5f, 0x95, 0x7e, 0xf3, 0x8b, 0x9f, 0x3a, 0x74, 0xfa, 0xae, 0x80, 0x4a,
	0x6e, 0x37, 0x60, 0x2c, 0xad, 0x40, 0xf0, 0x8a, 0x09, 0x68, 0x1c, 0x72, 0xe1, 0x6c, 0x47, 0xa2,
	0x34, 0x90, 0xd6, 0x50, 0xc8, 0xca, 0xb2, 0xbd, 0x91, 0xf9, 0x95, 0x35, 0xe0, 0x74, 0xb9, 0x61,
	0xc7, 0xb7, 0x3f, 0x75, 0x53, 0x1e, 0x7e, 0xfb, 0xd8, 0x4b, 0x51, 0xaf, 0xb5, 0xfb, 0x86, 0x35,
	0x53, 0x90, 0x7a, 0x31, 0xfd, 0xe1, 0x95, 0x2e, 0xb8, 0x03, 0xb4, 0xec, 0x38, 0xee, 0xb9, 0x1b,
	0x05, 0xa0, 0x2a, 0x79, 0x0a, 0x2d, 0x6b, 0xe3, 0x36, 0x9d, 0x98, 0xfc, 0x05, 0x24, 0x7e, 0xf5,
	0xbd, 0xdf, 0x8e, 0x7d, 0x33, 0x69, 0x0d, 0x34, 0x56, 0x7e, 0x20, 0x37, 0x46, 0x62, 0x78, 0x4b,
	0x06, 0x02, 0x0a, 0x19, 0x05, 0xb3, 0x60, 0x3e, 0x5c, 0x5c, 0xed, 0xb6, 0xf1, 0xa8, 0x04, 0xac,
	0x9f, 0x93, 0xf6, 0x9a, 0x64, 0x5e, 0x0c, 0x9f, 0xc8, 0xa8, 0xfd, 0x2e, 0x7f, 0x4d, 0xc5, 0x75,
	0x13, 0x9d, 0xcd, 0x82, 0xf9, 0xe5, 0x62, 0xb2, 0xdb, 0xc6, 0xe1, 0xc9, 0xbb, 0x17, 0x93, 0xac,
	0x6f, 0x0d, 0xef, 0xc8, 0x50, 0x22, 0x02, 0x46, 0xe7, 0x9e, 0xb9, 0xde, 0x6d, 0xe3, 0x71, 0xc7,
	0xf8, 0x73, 0x92, 0x75, 0x72, 0xf8, 0x42, 0x06, 0x05, 0x77, 0x3c, 0x1a, 0xcc, 0x82, 0xf9, 0xe8,
	0x7e, 0x42, 0xbb, 0x5a, 0xf4, 0x50, 0x8b, 0x2e, 0xdb, 0x5a, 0xfd, 0x78, 0xad, 0x3b, 0xc9, 0x3c,
	0x94, 0x5f, 0x78, 0xdb, 0xc3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x3b, 0xce, 0x36, 0xe5,
	0x01, 0x00, 0x00,
}
