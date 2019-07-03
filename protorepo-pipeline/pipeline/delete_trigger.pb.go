// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_trigger.proto

package pipeline

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//DeleteTrigger请求
type DeleteTriggerRequest struct {
	//
	//trigger id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTriggerRequest) Reset()         { *m = DeleteTriggerRequest{} }
func (m *DeleteTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTriggerRequest) ProtoMessage()    {}
func (*DeleteTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b0848ecd5389995, []int{0}
}
func (m *DeleteTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTriggerRequest.Unmarshal(m, b)
}
func (m *DeleteTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTriggerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTriggerRequest.Merge(m, src)
}
func (m *DeleteTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTriggerRequest.Size(m)
}
func (m *DeleteTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTriggerRequest proto.InternalMessageInfo

func (m *DeleteTriggerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DeleteTriggerApi返回
type DeleteTriggerResponseWrapper struct {
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

func (m *DeleteTriggerResponseWrapper) Reset()         { *m = DeleteTriggerResponseWrapper{} }
func (m *DeleteTriggerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteTriggerResponseWrapper) ProtoMessage()    {}
func (*DeleteTriggerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b0848ecd5389995, []int{1}
}
func (m *DeleteTriggerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTriggerResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteTriggerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTriggerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteTriggerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTriggerResponseWrapper.Merge(m, src)
}
func (m *DeleteTriggerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteTriggerResponseWrapper.Size(m)
}
func (m *DeleteTriggerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTriggerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTriggerResponseWrapper proto.InternalMessageInfo

func (m *DeleteTriggerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteTriggerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteTriggerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteTriggerResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteTriggerRequest)(nil), "pipeline.DeleteTriggerRequest")
	proto.RegisterType((*DeleteTriggerResponseWrapper)(nil), "pipeline.DeleteTriggerResponseWrapper")
}

func init() { proto.RegisterFile("delete_trigger.proto", fileDescriptor_2b0848ecd5389995) }

var fileDescriptor_2b0848ecd5389995 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0x5f, 0xfb, 0x78, 0x9d, 0x3e, 0xb4, 0x84, 0x52, 0x42, 0x15, 0x52, 0xa3, 0x48,
	0x37, 0x49, 0xac, 0x05, 0xf1, 0xcf, 0xae, 0xd8, 0x85, 0xdb, 0x20, 0xb8, 0x10, 0x95, 0x69, 0x73,
	0x3b, 0x0e, 0x26, 0xbd, 0xe3, 0x64, 0x6a, 0xfd, 0x83, 0x5f, 0x35, 0x82, 0x1f, 0x21, 0x9f, 0x40,
	0x32, 0x53, 0x6d, 0x71, 0x35, 0x73, 0xef, 0xf9, 0x9d, 0xc3, 0xe1, 0x92, 0x56, 0x0c, 0x09, 0x28,
	0xb8, 0x53, 0x92, 0x33, 0x06, 0x32, 0x10, 0x12, 0x15, 0xda, 0xff, 0x04, 0x17, 0x90, 0xf0, 0x19,
	0x74, 0x7c, 0xc6, 0xd5, 0xfd, 0x7c, 0x1c, 0x4c, 0x30, 0x0d, 0x19, 0x32, 0x0c, 0x35, 0x30, 0x9e,
	0x4f, 0xf5, 0xa4, 0x07, 0xfd, 0x33, 0xc6, 0xce, 0xd1, 0x1a, 0x9e, 0x2e, 0xb8, 0x7a, 0xc0, 0x45,
	0xc8, 0xd0, 0xd7, 0xa2, 0xff, 0x44, 0x13, 0x1e, 0x53, 0x85, 0x32, 0x0b, 0x7f, 0xbe, 0x4b, 0xdf,
	0x16, 0x43, 0x64, 0x09, 0xac, 0xd2, 0x21, 0x15, 0xea, 0xc5, 0x88, 0xde, 0x05, 0x69, 0x9d, 0xeb,
	0x96, 0x97, 0xa6, 0x64, 0x04, 0x8f, 0x73, 0xc8, 0x94, 0xdd, 0x27, 0x15, 0x1e, 0x3b, 0x56, 0xd7,
	0xea, 0xd5, 0x87, 0x3b, 0x45, 0xee, 0xd6, 0xa7, 0x28, 0xd3, 0x53, 0x8f, 0xc7, 0xde, 0xe7, 0x87,
	0xdb, 0x24, 0x1b, 0xb7, 0xd7, 0x07, 0xfe, 0x09, 0xf5, 0x5f, 0x6f, 0xde, 0xfa, 0x83, 0xf7, 0xbd,
	0xa8, 0x52, 0x0a, 0x16, 0xd9, 0xfe, 0x95, 0x95, 0x09, 0x9c, 0x65, 0x70, 0x25, 0xa9, 0x10, 0x20,
	0xed, 0x5d, 0x52, 0x9d, 0x60, 0x0c, 0x3a, 0xb5, 0x36, 0xdc, 0x2c, 0x72, 0xb7, 0x61, 0x52, 0xcb,
	0xad, 0x17, 0x69, 0xd1, 0x3e, 0x26, 0x8d, 0xf2, 0x1d, 0x3d, 0x8b, 0x84, 0xf2, 0x99, 0x53, 0xd1,
	0x0d, 0xda, 0x45, 0xee, 0xda, 0x2b, 0x76, 0x29, 0x7a, 0xd1, 0x3a, 0x6a, 0xef, 0x93, 0x1a, 0x48,
	0x89, 0xd2, 0xf9, 0xa3, 0x3d, 0xcd, 0x22, 0x77, 0xff, 0x1b, 0x8f, 0x5e, 0x7b, 0x91, 0x91, 0xed,
	0x33, 0x52, 0x8d, 0xa9, 0xa2, 0x4e, 0xb5, 0x6b, 0xf5, 0x1a, 0x87, 0xed, 0xc0, 0x9c, 0x27, 0xf8,
	0x3e, 0x4f, 0x30, 0x2a, 0xcf, 0xb3, 0x5e, 0xaf, 0xa4, 0xbd, 0x48, 0x9b, 0xc6, 0x7f, 0x35, 0x36,
	0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x18, 0x6a, 0x56, 0xdc, 0x01, 0x00, 0x00,
}