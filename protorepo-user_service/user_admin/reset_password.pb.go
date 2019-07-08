// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reset_password.proto

package user_admin

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
//ResetPassword请求
type ResetPasswordRequest struct {
	//
	//用户名
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" form:"username"`
	//
	//时间戳
	Ts int32 `protobuf:"varint,2,opt,name=ts,proto3" json:"ts" form:"ts"`
	//
	//新密码
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" form:"password"`
	//
	//重置密码key
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key" form:"key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordRequest) Reset()         { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordRequest) ProtoMessage()    {}
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d992df8d703869fe, []int{0}
}
func (m *ResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordRequest.Unmarshal(m, b)
}
func (m *ResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordRequest.Merge(m, src)
}
func (m *ResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordRequest.Size(m)
}
func (m *ResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordRequest proto.InternalMessageInfo

func (m *ResetPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ResetPasswordRequest) GetTs() int32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *ResetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

//
//ResetPasswordApi返回
type ResetPasswordResponseWrapper struct {
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

func (m *ResetPasswordResponseWrapper) Reset()         { *m = ResetPasswordResponseWrapper{} }
func (m *ResetPasswordResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordResponseWrapper) ProtoMessage()    {}
func (*ResetPasswordResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_d992df8d703869fe, []int{1}
}
func (m *ResetPasswordResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordResponseWrapper.Unmarshal(m, b)
}
func (m *ResetPasswordResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ResetPasswordResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordResponseWrapper.Merge(m, src)
}
func (m *ResetPasswordResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordResponseWrapper.Size(m)
}
func (m *ResetPasswordResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordResponseWrapper proto.InternalMessageInfo

func (m *ResetPasswordResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResetPasswordResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ResetPasswordResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResetPasswordResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ResetPasswordRequest)(nil), "user_admin.ResetPasswordRequest")
	proto.RegisterType((*ResetPasswordResponseWrapper)(nil), "user_admin.ResetPasswordResponseWrapper")
}

func init() { proto.RegisterFile("reset_password.proto", fileDescriptor_d992df8d703869fe) }

var fileDescriptor_d992df8d703869fe = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0x86, 0x33, 0xfc, 0x7c, 0xf9, 0x28, 0x2a, 0xa6, 0x12, 0x32, 0x41, 0xcd, 0x90, 0x9a, 0x18,
	0x36, 0x0e, 0x89, 0x6e, 0x8c, 0xee, 0x26, 0x61, 0x6f, 0xba, 0x71, 0x49, 0x0a, 0x73, 0x18, 0x09,
	0xcc, 0xb4, 0xb6, 0x9d, 0x28, 0x97, 0xe6, 0xcd, 0xcc, 0xc2, 0x4b, 0xe8, 0x15, 0x98, 0xb6, 0x0c,
	0x10, 0x56, 0xed, 0x39, 0xef, 0x73, 0x7e, 0xde, 0x1c, 0xd4, 0x97, 0xa0, 0x40, 0xcf, 0x04, 0x53,
	0xea, 0x8b, 0xcb, 0x34, 0x16, 0x92, 0x6b, 0x8e, 0x51, 0xa9, 0x40, 0xce, 0x58, 0x9a, 0xaf, 0x8a,
	0xe1, 0x43, 0xb6, 0xd2, 0x1f, 0xe5, 0x3c, 0x5e, 0xf0, 0x7c, 0x92, 0xf1, 0x8c, 0x4f, 0x1c, 0x32,
	0x2f, 0x97, 0x2e, 0x72, 0x81, 0xfb, 0xf9, 0xd2, 0xe1, 0x75, 0xc6, 0x79, 0xb6, 0x81, 0x03, 0x05,
	0xb9, 0xd0, 0x5b, 0x2f, 0x92, 0x9f, 0x00, 0xf5, 0xa9, 0x1d, 0xf8, 0xb6, 0x9b, 0x47, 0xe1, 0xb3,
	0x04, 0xa5, 0xf1, 0x04, 0xfd, 0xb7, 0x23, 0x0b, 0x96, 0x43, 0x18, 0x8c, 0x82, 0x71, 0x27, 0xb9,
	0x32, 0x55, 0xd4, 0x5b, 0x72, 0x99, 0xbf, 0x90, 0x5a, 0x21, 0x74, 0x0f, 0xe1, 0x5b, 0xd4, 0xd0,
	0x2a, 0x6c, 0x8c, 0x82, 0x71, 0x3b, 0x39, 0x37, 0x55, 0xd4, 0xf1, 0xa8, 0x56, 0x84, 0x36, 0xb4,
	0xb2, 0xfd, 0x6a, 0x4b, 0x61, 0xf3, 0xb4, 0x5f, 0xad, 0x10, 0xba, 0x87, 0xf0, 0x08, 0x35, 0xd7,
	0xb0, 0x0d, 0x5b, 0x8e, 0xbd, 0x30, 0x55, 0x84, 0x3c, 0xbb, 0x86, 0x2d, 0xa1, 0x56, 0x22, 0xbf,
	0x01, 0xba, 0x39, 0xd9, 0x5d, 0x09, 0x5e, 0x28, 0x78, 0x97, 0x4c, 0x08, 0x90, 0xf8, 0x0e, 0xb5,
	0x16, 0x3c, 0xf5, 0xfb, 0xb7, 0x93, 0x9e, 0xa9, 0xa2, 0xae, 0xef, 0x61, 0xb3, 0x84, 0x3a, 0x11,
	0x3f, 0xa3, 0xae, 0x7d, 0xa7, 0xdf, 0x62, 0xc3, 0x56, 0x85, 0x33, 0xd0, 0x49, 0x06, 0xa6, 0x8a,
	0xf0, 0x81, 0xdd, 0x89, 0x84, 0x1e, 0xa3, 0xf8, 0x1e, 0xb5, 0x41, 0x4a, 0x2e, 0x77, 0x7e, 0x2e,
	0x4d, 0x15, 0x9d, 0xf9, 0x1a, 0x97, 0x26, 0xd4, 0xcb, 0xf8, 0x15, 0xb5, 0x52, 0xa6, 0x99, 0xb3,
	0xd2, 0x7d, 0x1c, 0xc4, 0xfe, 0x1e, 0x71, 0x7d, 0x8f, 0x78, 0x6a, 0xef, 0x71, 0xbc, 0x9e, 0xa5,
	0x09, 0x75, 0x45, 0xf3, 0x7f, 0x0e, 0x7b, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x20, 0x84, 0xc0,
	0x9a, 0x17, 0x02, 0x00, 0x00,
}
