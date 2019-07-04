// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alter_password.proto

package user_admin

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
//AlterPassword请求
type AlterPasswordRequest struct {
	//
	//用户名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//密码
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password" form:"password"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlterPasswordRequest) Reset()         { *m = AlterPasswordRequest{} }
func (m *AlterPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*AlterPasswordRequest) ProtoMessage()    {}
func (*AlterPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cde7fbc89cf2f0f, []int{0}
}
func (m *AlterPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlterPasswordRequest.Unmarshal(m, b)
}
func (m *AlterPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlterPasswordRequest.Marshal(b, m, deterministic)
}
func (m *AlterPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlterPasswordRequest.Merge(m, src)
}
func (m *AlterPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_AlterPasswordRequest.Size(m)
}
func (m *AlterPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlterPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlterPasswordRequest proto.InternalMessageInfo

func (m *AlterPasswordRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlterPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

//
//AlterPasswordApi返回
type AlterPasswordResponseWrapper struct {
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

func (m *AlterPasswordResponseWrapper) Reset()         { *m = AlterPasswordResponseWrapper{} }
func (m *AlterPasswordResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AlterPasswordResponseWrapper) ProtoMessage()    {}
func (*AlterPasswordResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cde7fbc89cf2f0f, []int{1}
}
func (m *AlterPasswordResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlterPasswordResponseWrapper.Unmarshal(m, b)
}
func (m *AlterPasswordResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlterPasswordResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AlterPasswordResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlterPasswordResponseWrapper.Merge(m, src)
}
func (m *AlterPasswordResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AlterPasswordResponseWrapper.Size(m)
}
func (m *AlterPasswordResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AlterPasswordResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AlterPasswordResponseWrapper proto.InternalMessageInfo

func (m *AlterPasswordResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AlterPasswordResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AlterPasswordResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AlterPasswordResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AlterPasswordRequest)(nil), "user_admin.AlterPasswordRequest")
	proto.RegisterType((*AlterPasswordResponseWrapper)(nil), "user_admin.AlterPasswordResponseWrapper")
}

func init() { proto.RegisterFile("alter_password.proto", fileDescriptor_4cde7fbc89cf2f0f) }

var fileDescriptor_4cde7fbc89cf2f0f = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x89, 0xa6, 0xa2, 0x53, 0xa1, 0x65, 0x2c, 0x25, 0x54, 0x21, 0x65, 0xbc, 0xd0, 0x85,
	0x49, 0x44, 0x41, 0xbc, 0xac, 0x5a, 0xe8, 0x5e, 0xb2, 0x11, 0xbc, 0x95, 0x69, 0x33, 0x8d, 0xc1,
	0x24, 0x27, 0x4e, 0x26, 0xc6, 0x0b, 0x6e, 0x7c, 0xd0, 0x08, 0x7d, 0x84, 0x3c, 0x81, 0x64, 0x92,
	0x5e, 0x70, 0x35, 0xe7, 0x9c, 0xff, 0xff, 0x66, 0xfe, 0x39, 0xa8, 0x45, 0x7d, 0xc1, 0xf8, 0x28,
	0xa2, 0x71, 0x9c, 0x02, 0x77, 0xcc, 0x88, 0x83, 0x00, 0x8c, 0x92, 0x98, 0xf1, 0x11, 0x75, 0x02,
	0x2f, 0xec, 0x18, 0xae, 0x27, 0x9e, 0x93, 0xb1, 0x39, 0x81, 0xc0, 0x72, 0xc1, 0x05, 0x4b, 0x5a,
	0xc6, 0xc9, 0x54, 0x76, 0xb2, 0x91, 0x55, 0x89, 0x76, 0xce, 0x57, 0xec, 0x41, 0xea, 0x89, 0x17,
	0x48, 0x2d, 0x17, 0x0c, 0x29, 0x1a, 0x6f, 0xd4, 0xf7, 0x1c, 0x2a, 0x80, 0xc7, 0xd6, 0xa2, 0xac,
	0xb8, 0x5d, 0x17, 0xc0, 0xf5, 0xd9, 0xf2, 0x76, 0x16, 0x44, 0xe2, 0xa3, 0x14, 0xc9, 0x8f, 0x82,
	0x5a, 0xfd, 0x22, 0xe8, 0x4d, 0x95, 0xd3, 0x66, 0xaf, 0x09, 0x8b, 0x05, 0xbe, 0x44, 0x6a, 0x48,
	0x03, 0xa6, 0x29, 0x5d, 0xa5, 0xb7, 0x35, 0x38, 0xcc, 0x33, 0xbd, 0x3e, 0x05, 0x1e, 0x5c, 0x91,
	0x62, 0x4a, 0x66, 0xbf, 0x3a, 0x46, 0xcd, 0xa7, 0x7b, 0x6a, 0x7c, 0xf6, 0x8d, 0xbb, 0xc7, 0x87,
	0xf4, 0xeb, 0xe4, 0xf8, 0xfb, 0xc0, 0x96, 0x08, 0xb6, 0xd0, 0xe6, 0xfc, 0xd7, 0xda, 0x9a, 0xc4,
	0x77, 0xf2, 0x4c, 0x6f, 0x94, 0xf8, 0x5c, 0x21, 0xf6, 0xc2, 0x44, 0x66, 0x0a, 0xda, 0xfb, 0x17,
	0x22, 0x8e, 0x20, 0x8c, 0xd9, 0x2d, 0xa7, 0x51, 0xc4, 0x38, 0xde, 0x47, 0xea, 0x04, 0x9c, 0x32,
	0x4c, 0x6d, 0xd0, 0x58, 0x86, 0x29, 0xa6, 0xc4, 0x96, 0x22, 0xbe, 0x40, 0xf5, 0xe2, 0x1c, 0xbe,
	0x47, 0x3e, 0xf5, 0xc2, 0xea, 0xe5, 0x76, 0x9e, 0xe9, 0x78, 0xe9, 0xad, 0x44, 0x62, 0xaf, 0x5a,
	0xf1, 0x11, 0xaa, 0x31, 0xce, 0x81, 0x6b, 0xeb, 0x92, 0x69, 0xe6, 0x99, 0xbe, 0x5d, 0x32, 0x72,
	0x4c, 0xec, 0x52, 0xc6, 0xd7, 0x48, 0x75, 0xa8, 0xa0, 0x9a, 0xda, 0x55, 0x7a, 0xf5, 0xd3, 0xb6,
	0x59, 0x2e, 0xd6, 0x9c, 0x2f, 0xd6, 0x1c, 0x16, 0x8b, 0x5d, 0x8d, 0x57, 0xb8, 0x89, 0x2d, 0xa1,
	0xf1, 0x86, 0xb4, 0x9d, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x5f, 0x55, 0xe5, 0x18, 0x02,
	0x00, 0x00,
}
