// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package subscriber_manager

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
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
//PubSubscriberDelete请求
type PubSubscriberDeleteRequest struct {
	//
	//订阅id
	SubscriberId         string   `protobuf:"bytes,1,opt,name=subscriberId,proto3" json:"subscriberId" form:"subscriberId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubSubscriberDeleteRequest) Reset()         { *m = PubSubscriberDeleteRequest{} }
func (m *PubSubscriberDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*PubSubscriberDeleteRequest) ProtoMessage()    {}
func (*PubSubscriberDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *PubSubscriberDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubscriberDeleteRequest.Unmarshal(m, b)
}
func (m *PubSubscriberDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubscriberDeleteRequest.Marshal(b, m, deterministic)
}
func (m *PubSubscriberDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubscriberDeleteRequest.Merge(m, src)
}
func (m *PubSubscriberDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_PubSubscriberDeleteRequest.Size(m)
}
func (m *PubSubscriberDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubscriberDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubscriberDeleteRequest proto.InternalMessageInfo

func (m *PubSubscriberDeleteRequest) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

//
//PubSubscriberDeleteApi返回
type PubSubscriberDeleteResponseWrapper struct {
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

func (m *PubSubscriberDeleteResponseWrapper) Reset()         { *m = PubSubscriberDeleteResponseWrapper{} }
func (m *PubSubscriberDeleteResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*PubSubscriberDeleteResponseWrapper) ProtoMessage()    {}
func (*PubSubscriberDeleteResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *PubSubscriberDeleteResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubscriberDeleteResponseWrapper.Unmarshal(m, b)
}
func (m *PubSubscriberDeleteResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubscriberDeleteResponseWrapper.Marshal(b, m, deterministic)
}
func (m *PubSubscriberDeleteResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubscriberDeleteResponseWrapper.Merge(m, src)
}
func (m *PubSubscriberDeleteResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_PubSubscriberDeleteResponseWrapper.Size(m)
}
func (m *PubSubscriberDeleteResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubscriberDeleteResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubscriberDeleteResponseWrapper proto.InternalMessageInfo

func (m *PubSubscriberDeleteResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PubSubscriberDeleteResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *PubSubscriberDeleteResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PubSubscriberDeleteResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PubSubscriberDeleteRequest)(nil), "subscriber_manager.PubSubscriberDeleteRequest")
	proto.RegisterType((*PubSubscriberDeleteResponseWrapper)(nil), "subscriber_manager.PubSubscriberDeleteResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xb6, 0x82, 0xdb, 0xa2, 0x65, 0x85, 0x12, 0xea, 0x21, 0x65, 0x15, 0x29, 0x48,
	0x12, 0xb5, 0x20, 0xfe, 0xb9, 0x15, 0x7b, 0xf0, 0xa4, 0xc4, 0x83, 0x07, 0x51, 0xd9, 0x34, 0xdb,
	0x18, 0x4c, 0x3a, 0x71, 0xb3, 0xb1, 0xfe, 0xc1, 0xaf, 0x1a, 0xc1, 0x93, 0xe7, 0x7c, 0x02, 0xc9,
	0xa4, 0xb6, 0x11, 0x3c, 0xed, 0xce, 0xbc, 0xdf, 0x9b, 0x79, 0x0c, 0x69, 0x7a, 0x22, 0x14, 0x4a,
	0x58, 0xb1, 0x04, 0x05, 0x94, 0x26, 0xa9, 0x9b, 0x8c, 0x64, 0xe0, 0x0a, 0x79, 0x1f, 0xf1, 0x09,
	0xf7, 0x85, 0xec, 0x98, 0x7e, 0xa0, 0x1e, 0x52, 0xd7, 0x1a, 0x41, 0x64, 0xfb, 0xe0, 0x83, 0x8d,
	0xa8, 0x9b, 0x8e, 0xb1, 0xc2, 0x02, 0x7f, 0xe5, 0x88, 0xce, 0x61, 0x05, 0x8f, 0xa6, 0x81, 0x7a,
	0x84, 0xa9, 0xed, 0x83, 0x89, 0xa2, 0xf9, 0xcc, 0xc3, 0xc0, 0xe3, 0x0a, 0x64, 0x62, 0xcf, 0xbf,
	0x33, 0xdf, 0xa6, 0x0f, 0xe0, 0x87, 0x62, 0x31, 0x5d, 0x44, 0xb1, 0x7a, 0x2d, 0x45, 0x16, 0x91,
	0xce, 0x65, 0xea, 0x5e, 0xcd, 0xc3, 0x9d, 0x61, 0x68, 0x47, 0x3c, 0xa5, 0x22, 0x51, 0xf4, 0x82,
	0x34, 0x17, 0xb9, 0xcf, 0x3d, 0x5d, 0xeb, 0x6a, 0xbd, 0xd5, 0xc1, 0x6e, 0x9e, 0x19, 0x1b, 0x63,
	0x90, 0xd1, 0x09, 0xab, 0xaa, 0xec, 0xeb, 0xd3, 0x68, 0x91, 0xb5, 0xbb, 0x9b, 0x3d, 0xf3, 0x98,
	0x9b, 0x6f, 0xb7, 0xef, 0xfb, 0xfd, 0x8f, 0x6d, 0xe7, 0xcf, 0x00, 0xf6, 0xad, 0x11, 0xf6, 0xef,
	0xbe, 0x24, 0x86, 0x49, 0x22, 0xae, 0x25, 0x8f, 0x63, 0x21, 0xe9, 0x16, 0xa9, 0x8d, 0xc0, 0x13,
	0xb8, 0xaf, 0x3e, 0x58, 0xcf, 0x33, 0xa3, 0x51, 0xee, 0x2b, 0xba, 0xcc, 0x41, 0x91, 0x1e, 0x91,
	0x46, 0xf1, 0x0e, 0x5f, 0xe2, 0x90, 0x07, 0x13, 0x7d, 0x09, 0xb3, 0xb5, 0xf3, 0xcc, 0xa0, 0x0b,
	0x76, 0x26, 0x32, 0xa7, 0x8a, 0xd2, 0x1d, 0x52, 0x17, 0x52, 0x82, 0xd4, 0x97, 0xd1, 0xd3, 0xca,
	0x33, 0xa3, 0x59, 0x7a, 0xb0, 0xcd, 0x9c, 0x52, 0xa6, 0xa7, 0xa4, 0xe6, 0x71, 0xc5, 0xf5, 0x5a,
	0x57, 0xeb, 0x35, 0x0e, 0xda, 0x56, 0x79, 0x48, 0xeb, 0xf7, 0x90, 0xd6, 0xb0, 0x38, 0x64, 0x35,
	0x5e, 0x41, 0x33, 0x07, 0x4d, 0xee, 0x0a, 0x62, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77,
	0xda, 0x11, 0x87, 0x08, 0x02, 0x00, 0x00,
}