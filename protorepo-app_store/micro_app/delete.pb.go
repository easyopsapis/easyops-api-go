// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package micro_app

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
//DeleteAppStoreMicroApp请求
type DeleteAppStoreMicroAppRequest struct {
	//
	//小产品id
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id" form:"app_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAppStoreMicroAppRequest) Reset()         { *m = DeleteAppStoreMicroAppRequest{} }
func (m *DeleteAppStoreMicroAppRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAppStoreMicroAppRequest) ProtoMessage()    {}
func (*DeleteAppStoreMicroAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteAppStoreMicroAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAppStoreMicroAppRequest.Unmarshal(m, b)
}
func (m *DeleteAppStoreMicroAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAppStoreMicroAppRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAppStoreMicroAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAppStoreMicroAppRequest.Merge(m, src)
}
func (m *DeleteAppStoreMicroAppRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAppStoreMicroAppRequest.Size(m)
}
func (m *DeleteAppStoreMicroAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAppStoreMicroAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAppStoreMicroAppRequest proto.InternalMessageInfo

func (m *DeleteAppStoreMicroAppRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

//
//DeleteAppStoreMicroAppApi返回
type DeleteAppStoreMicroAppResponseWrapper struct {
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

func (m *DeleteAppStoreMicroAppResponseWrapper) Reset()         { *m = DeleteAppStoreMicroAppResponseWrapper{} }
func (m *DeleteAppStoreMicroAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteAppStoreMicroAppResponseWrapper) ProtoMessage()    {}
func (*DeleteAppStoreMicroAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteAppStoreMicroAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteAppStoreMicroAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteAppStoreMicroAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper.Merge(m, src)
}
func (m *DeleteAppStoreMicroAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper.Size(m)
}
func (m *DeleteAppStoreMicroAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAppStoreMicroAppResponseWrapper proto.InternalMessageInfo

func (m *DeleteAppStoreMicroAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteAppStoreMicroAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteAppStoreMicroAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteAppStoreMicroAppResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteAppStoreMicroAppRequest)(nil), "micro_app.DeleteAppStoreMicroAppRequest")
	proto.RegisterType((*DeleteAppStoreMicroAppResponseWrapper)(nil), "micro_app.DeleteAppStoreMicroAppResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcb, 0x4f, 0xea, 0x40,
	0x14, 0xc6, 0xd3, 0x7b, 0x81, 0x84, 0x81, 0x9b, 0x7b, 0xd3, 0x05, 0x69, 0xb8, 0x21, 0x25, 0xe3,
	0x23, 0x2c, 0x6c, 0x8b, 0x62, 0x8c, 0x8f, 0x15, 0x44, 0x16, 0xc6, 0xb8, 0xa9, 0x0b, 0x13, 0x89,
	0x92, 0x81, 0x0e, 0xb5, 0xb1, 0xe5, 0x1c, 0xa7, 0x83, 0xf8, 0x88, 0xff, 0x6a, 0x4d, 0xdc, 0xba,
	0xeb, 0x5f, 0x60, 0x3a, 0x05, 0x61, 0xe3, 0x6a, 0xce, 0x39, 0xdf, 0xef, 0x9b, 0xf3, 0xe5, 0x90,
	0xaa, 0xc7, 0x43, 0x2e, 0xb9, 0x8d, 0x02, 0x24, 0xe8, 0xe5, 0x28, 0x18, 0x0b, 0x18, 0x32, 0xc4,
	0xba, 0xe5, 0x07, 0xf2, 0x6e, 0x36, 0xb2, 0xc7, 0x10, 0x39, 0x3e, 0xf8, 0xe0, 0x28, 0x62, 0x34,
	0x9b, 0xa8, 0x4e, 0x35, 0xaa, 0xca, 0x9d, 0xf5, 0x83, 0x35, 0x3c, 0x9a, 0x07, 0xf2, 0x1e, 0xe6,
	0x8e, 0x0f, 0x96, 0x12, 0xad, 0x47, 0x16, 0x06, 0x1e, 0x93, 0x20, 0x62, 0xe7, 0xbb, 0x5c, 0xf8,
	0xfe, 0xfb, 0x00, 0x7e, 0xc8, 0x57, 0xbf, 0xf3, 0x08, 0xe5, 0x73, 0x2e, 0xd2, 0x90, 0x34, 0x4e,
	0x55, 0xbc, 0x2e, 0xe2, 0xa5, 0x04, 0xc1, 0x2f, 0xb2, 0x78, 0x5d, 0x44, 0x97, 0x3f, 0xcc, 0x78,
	0x2c, 0xf5, 0x73, 0x52, 0x62, 0x88, 0xc3, 0xc0, 0x33, 0xb4, 0xa6, 0xd6, 0x2a, 0xf7, 0xf6, 0xd3,
	0xc4, 0xfc, 0x33, 0x01, 0x11, 0x1d, 0xd3, 0x7c, 0x4e, 0x3f, 0xde, 0x4d, 0x93, 0x34, 0x6e, 0x07,
	0xcc, 0x7a, 0xe9, 0x5a, 0xd7, 0xc3, 0x9b, 0x41, 0xdb, 0x3a, 0x5a, 0xd6, 0xaf, 0xed, 0x9d, 0xce,
	0xee, 0xdb, 0xa6, 0x5b, 0x64, 0x88, 0x67, 0x1e, 0xfd, 0xd4, 0xc8, 0xd6, 0x4f, 0xeb, 0x62, 0x84,
	0x69, 0xcc, 0xaf, 0x04, 0x43, 0xe4, 0x42, 0xdf, 0x20, 0x85, 0x31, 0x78, 0x5c, 0x2d, 0x2d, 0xf6,
	0xfe, 0xa6, 0x89, 0x59, 0xc9, 0x97, 0x66, 0x53, 0xea, 0x2a, 0x51, 0x3f, 0x24, 0x95, 0xec, 0xed,
	0x3f, 0x61, 0xc8, 0x82, 0xa9, 0xf1, 0x4b, 0x05, 0xac, 0xa5, 0x89, 0xa9, 0xaf, 0xd8, 0x85, 0x48,
	0xdd, 0x75, 0x54, 0xdf, 0x26, 0x45, 0x2e, 0x04, 0x08, 0xe3, 0xb7, 0xf2, 0xfc, 0x4b, 0x13, 0xb3,
	0x9a, 0x7b, 0xd4, 0x98, 0xba, 0xb9, 0xac, 0x9f, 0x90, 0x82, 0xc7, 0x24, 0x33, 0x0a, 0x4d, 0xad,
	0x55, 0xd9, 0xab, 0xd9, 0xf9, 0x29, 0xed, 0xe5, 0x29, 0xed, 0x7e, 0x76, 0xca, 0xf5, 0x78, 0x19,
	0x4d, 0x5d, 0x65, 0x1a, 0x95, 0x14, 0xd6, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xc4, 0x5b,
	0xa6, 0x01, 0x02, 0x00, 0x00,
}
