// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_apikey.proto

package apikey

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
//DeleteApiKey请求
type DeleteApiKeyRequest struct {
	//
	//access_key
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key" form:"access_key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteApiKeyRequest) Reset()         { *m = DeleteApiKeyRequest{} }
func (m *DeleteApiKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteApiKeyRequest) ProtoMessage()    {}
func (*DeleteApiKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8149896ac671a73, []int{0}
}
func (m *DeleteApiKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteApiKeyRequest.Unmarshal(m, b)
}
func (m *DeleteApiKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteApiKeyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteApiKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteApiKeyRequest.Merge(m, src)
}
func (m *DeleteApiKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteApiKeyRequest.Size(m)
}
func (m *DeleteApiKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteApiKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteApiKeyRequest proto.InternalMessageInfo

func (m *DeleteApiKeyRequest) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

//
//DeleteApiKeyApi返回
type DeleteApiKeyResponseWrapper struct {
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

func (m *DeleteApiKeyResponseWrapper) Reset()         { *m = DeleteApiKeyResponseWrapper{} }
func (m *DeleteApiKeyResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteApiKeyResponseWrapper) ProtoMessage()    {}
func (*DeleteApiKeyResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8149896ac671a73, []int{1}
}
func (m *DeleteApiKeyResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteApiKeyResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteApiKeyResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteApiKeyResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteApiKeyResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteApiKeyResponseWrapper.Merge(m, src)
}
func (m *DeleteApiKeyResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteApiKeyResponseWrapper.Size(m)
}
func (m *DeleteApiKeyResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteApiKeyResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteApiKeyResponseWrapper proto.InternalMessageInfo

func (m *DeleteApiKeyResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteApiKeyResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteApiKeyResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteApiKeyResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteApiKeyRequest)(nil), "apikey.DeleteApiKeyRequest")
	proto.RegisterType((*DeleteApiKeyResponseWrapper)(nil), "apikey.DeleteApiKeyResponseWrapper")
}

func init() { proto.RegisterFile("delete_apikey.proto", fileDescriptor_b8149896ac671a73) }

var fileDescriptor_b8149896ac671a73 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0xa7, 0xba, 0x0d, 0x96, 0x0a, 0x6a, 0x86, 0xa3, 0x6c, 0x87, 0x8e, 0x08, 0xb2, 0x8b, 0x1d,
	0xa8, 0x07, 0xd1, 0x93, 0xc3, 0x9d, 0x7a, 0xcb, 0xc5, 0xe3, 0x48, 0xdb, 0xb7, 0x5a, 0xd6, 0x2e,
	0x31, 0x4d, 0xc1, 0x7e, 0xd9, 0x82, 0x5f, 0xa1, 0x9f, 0x40, 0xfa, 0x32, 0xe9, 0x3c, 0x25, 0xbf,
	0xf7, 0xfb, 0xf3, 0x1e, 0x3f, 0x32, 0x49, 0x20, 0x07, 0x03, 0x5b, 0xa1, 0xb2, 0x3d, 0xd4, 0x81,
	0xd2, 0xd2, 0x48, 0x3a, 0xb2, 0x68, 0x76, 0x9f, 0x66, 0xe6, 0xb3, 0x8a, 0x82, 0x58, 0x16, 0xab,
	0x54, 0xa6, 0x72, 0x85, 0x74, 0x54, 0xed, 0x10, 0x21, 0xc0, 0x9f, 0xb5, 0xcd, 0xe6, 0xa9, 0x94,
	0x69, 0x0e, 0xbd, 0x0a, 0x0a, 0x65, 0x8e, 0x99, 0x2c, 0x24, 0x93, 0x77, 0x5c, 0xf5, 0xa6, 0xb2,
	0x10, 0x6a, 0x0e, 0x5f, 0x15, 0x94, 0x86, 0x3e, 0x11, 0x22, 0xe2, 0x18, 0xca, 0x72, 0xbb, 0x87,
	0xda, 0x73, 0x16, 0xce, 0x72, 0xbc, 0xbe, 0x69, 0x1b, 0xff, 0x7a, 0x27, 0x75, 0xf1, 0xc2, 0x7a,
	0x8e, 0xf1, 0xb1, 0x05, 0x21, 0xd4, 0xec, 0xc7, 0x21, 0xf3, 0xff, 0x69, 0xa5, 0x92, 0x87, 0x12,
	0x3e, 0xb4, 0x50, 0x0a, 0x34, 0xbd, 0x25, 0x83, 0x58, 0x26, 0x80, 0x79, 0xc3, 0xf5, 0x65, 0xdb,
	0xf8, 0xae, 0xcd, 0xeb, 0xa6, 0x8c, 0x23, 0x49, 0x9f, 0x89, 0xdb, 0xbd, 0x9b, 0x6f, 0x95, 0x8b,
	0xec, 0xe0, 0x9d, 0xe1, 0xee, 0x69, 0xdb, 0xf8, 0xb4, 0xd7, 0x1e, 0x49, 0xc6, 0x4f, 0xa5, 0xf4,
	0x8e, 0x0c, 0x41, 0x6b, 0xa9, 0xbd, 0x73, 0xf4, 0x5c, 0xb5, 0x8d, 0x7f, 0x61, 0x3d, 0x38, 0x66,
	0xdc, 0xd2, 0xf4, 0x95, 0x0c, 0x12, 0x61, 0x84, 0x37, 0x58, 0x38, 0x4b, 0xf7, 0x61, 0x1a, 0xd8,
	0x7e, 0x82, 0xbf, 0x7e, 0x82, 0x4d, 0xd7, 0xcf, 0xe9, 0x79, 0x9d, 0x9a, 0x71, 0x34, 0x45, 0x23,
	0x94, 0x3d, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x64, 0xbf, 0x39, 0xa2, 0x01, 0x00, 0x00,
}
