// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: set_default.proto

package provider

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
//SetDefaultProvider请求
type SetDefaultProviderRequest struct {
	//
	//实例id，服务端自动生成
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDefaultProviderRequest) Reset()         { *m = SetDefaultProviderRequest{} }
func (m *SetDefaultProviderRequest) String() string { return proto.CompactTextString(m) }
func (*SetDefaultProviderRequest) ProtoMessage()    {}
func (*SetDefaultProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_163b734f77243279, []int{0}
}
func (m *SetDefaultProviderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefaultProviderRequest.Unmarshal(m, b)
}
func (m *SetDefaultProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefaultProviderRequest.Marshal(b, m, deterministic)
}
func (m *SetDefaultProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefaultProviderRequest.Merge(m, src)
}
func (m *SetDefaultProviderRequest) XXX_Size() int {
	return xxx_messageInfo_SetDefaultProviderRequest.Size(m)
}
func (m *SetDefaultProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefaultProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefaultProviderRequest proto.InternalMessageInfo

func (m *SetDefaultProviderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//SetDefaultProviderApi返回
type SetDefaultProviderResponseWrapper struct {
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

func (m *SetDefaultProviderResponseWrapper) Reset()         { *m = SetDefaultProviderResponseWrapper{} }
func (m *SetDefaultProviderResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SetDefaultProviderResponseWrapper) ProtoMessage()    {}
func (*SetDefaultProviderResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_163b734f77243279, []int{1}
}
func (m *SetDefaultProviderResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefaultProviderResponseWrapper.Unmarshal(m, b)
}
func (m *SetDefaultProviderResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefaultProviderResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SetDefaultProviderResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefaultProviderResponseWrapper.Merge(m, src)
}
func (m *SetDefaultProviderResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SetDefaultProviderResponseWrapper.Size(m)
}
func (m *SetDefaultProviderResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefaultProviderResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefaultProviderResponseWrapper proto.InternalMessageInfo

func (m *SetDefaultProviderResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SetDefaultProviderResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SetDefaultProviderResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SetDefaultProviderResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SetDefaultProviderRequest)(nil), "provider.SetDefaultProviderRequest")
	proto.RegisterType((*SetDefaultProviderResponseWrapper)(nil), "provider.SetDefaultProviderResponseWrapper")
}

func init() { proto.RegisterFile("set_default.proto", fileDescriptor_163b734f77243279) }

var fileDescriptor_163b734f77243279 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x53, 0x1e, 0xbc, 0x3c, 0x86, 0x97, 0xf7, 0xb0, 0x0b, 0x52, 0x71, 0x51, 0x18, 0x8d,
	0x61, 0xd3, 0x56, 0x24, 0x31, 0xfe, 0xd9, 0x11, 0xd9, 0x1a, 0x53, 0x17, 0x2e, 0x8c, 0x9a, 0x81,
	0x19, 0xea, 0xc4, 0x96, 0x3b, 0x4e, 0xa7, 0xe0, 0x9f, 0xf8, 0x55, 0x6b, 0xe2, 0xc6, 0x7d, 0x3f,
	0x81, 0xe9, 0x0c, 0x0a, 0x0b, 0x57, 0x33, 0xf7, 0x9e, 0xdf, 0x39, 0x39, 0xb9, 0x68, 0x23, 0x65,
	0xea, 0x96, 0xb2, 0x29, 0xc9, 0x62, 0xe5, 0x0b, 0x09, 0x0a, 0xec, 0x3f, 0x42, 0xc2, 0x9c, 0x53,
	0x26, 0xdb, 0x5e, 0xc4, 0xd5, 0x5d, 0x36, 0xf6, 0x27, 0x90, 0x04, 0x11, 0x44, 0x10, 0x68, 0x60,
	0x9c, 0x4d, 0xf5, 0xa4, 0x07, 0xfd, 0x33, 0xc6, 0xf6, 0xc1, 0x1a, 0x9e, 0x2c, 0xb8, 0xba, 0x87,
	0x45, 0x10, 0x81, 0xa7, 0x45, 0x6f, 0x4e, 0x62, 0x4e, 0x89, 0x02, 0x99, 0x06, 0xdf, 0xdf, 0xa5,
	0x6f, 0x2b, 0x02, 0x88, 0x62, 0xb6, 0x4a, 0x67, 0x89, 0x50, 0x4f, 0x46, 0xc4, 0x67, 0x68, 0xf3,
	0x82, 0xa9, 0x53, 0xd3, 0xf0, 0x7c, 0xd9, 0x2c, 0x64, 0x0f, 0x19, 0x4b, 0x95, 0xdd, 0x47, 0x15,
	0x4e, 0x1d, 0xab, 0x63, 0xf5, 0xea, 0xc3, 0x6e, 0x91, 0xbb, 0xf5, 0x29, 0xc8, 0xe4, 0x18, 0x73,
	0x8a, 0xdf, 0xdf, 0xdc, 0x26, 0xfa, 0x77, 0x73, 0xb5, 0xe7, 0x1d, 0x11, 0xef, 0xf9, 0xfa, 0xa5,
	0x3f, 0x78, 0xdd, 0x09, 0x2b, 0x9c, 0xe2, 0x0f, 0x0b, 0x75, 0x7f, 0x0a, 0x4c, 0x05, 0xcc, 0x52,
	0x76, 0x29, 0x89, 0x10, 0x4c, 0xda, 0xdb, 0xa8, 0x3a, 0x01, 0xca, 0x74, 0x74, 0x6d, 0xf8, 0xbf,
	0xc8, 0xdd, 0x86, 0x89, 0x2e, 0xb7, 0x38, 0xd4, 0xa2, 0x7d, 0x88, 0x1a, 0xe5, 0x3b, 0x7a, 0x14,
	0x31, 0xe1, 0x33, 0xa7, 0xa2, 0x6b, 0xb4, 0x8a, 0xdc, 0xb5, 0x57, 0xec, 0x52, 0xc4, 0xe1, 0x3a,
	0x6a, 0xef, 0xa2, 0x1a, 0x93, 0x12, 0xa4, 0xf3, 0x4b, 0x7b, 0x9a, 0x45, 0xee, 0xfe, 0x35, 0x1e,
	0xbd, 0xc6, 0xa1, 0x91, 0xed, 0x13, 0x54, 0xa5, 0x44, 0x11, 0xa7, 0xda, 0xb1, 0x7a, 0x8d, 0xfd,
	0x96, 0x6f, 0x0e, 0xe5, 0x7f, 0x1d, 0xca, 0x1f, 0x95, 0x87, 0x5a, 0xaf, 0x57, 0xd2, 0x38, 0xd4,
	0xa6, 0xf1, 0x6f, 0x8d, 0x0d, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x35, 0xa3, 0x9f, 0xe3,
	0x01, 0x00, 0x00,
}
