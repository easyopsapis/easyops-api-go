// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: forgot_password.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//ForgotPassword请求
type ForgotPasswordRequest struct {
	//
	//邮箱
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email" form:"email"`
	//
	//重置密码地址
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url" form:"url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordRequest) Reset()         { *m = ForgotPasswordRequest{} }
func (m *ForgotPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordRequest) ProtoMessage()    {}
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b498ac5f953be7, []int{0}
}
func (m *ForgotPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordRequest.Unmarshal(m, b)
}
func (m *ForgotPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordRequest.Merge(m, src)
}
func (m *ForgotPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordRequest.Size(m)
}
func (m *ForgotPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordRequest proto.InternalMessageInfo

func (m *ForgotPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ForgotPasswordRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

//
//ForgotPasswordApi返回
type ForgotPasswordResponseWrapper struct {
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

func (m *ForgotPasswordResponseWrapper) Reset()         { *m = ForgotPasswordResponseWrapper{} }
func (m *ForgotPasswordResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordResponseWrapper) ProtoMessage()    {}
func (*ForgotPasswordResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b498ac5f953be7, []int{1}
}
func (m *ForgotPasswordResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordResponseWrapper.Unmarshal(m, b)
}
func (m *ForgotPasswordResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordResponseWrapper.Merge(m, src)
}
func (m *ForgotPasswordResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordResponseWrapper.Size(m)
}
func (m *ForgotPasswordResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordResponseWrapper proto.InternalMessageInfo

func (m *ForgotPasswordResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ForgotPasswordResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ForgotPasswordResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ForgotPasswordResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ForgotPasswordRequest)(nil), "user_admin.ForgotPasswordRequest")
	proto.RegisterType((*ForgotPasswordResponseWrapper)(nil), "user_admin.ForgotPasswordResponseWrapper")
}

func init() { proto.RegisterFile("forgot_password.proto", fileDescriptor_b1b498ac5f953be7) }

var fileDescriptor_b1b498ac5f953be7 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x6f, 0x4b, 0xdb, 0x40,
	0x18, 0x27, 0x6b, 0x3b, 0xd8, 0x75, 0x6c, 0xe3, 0xa0, 0x25, 0x74, 0x8c, 0x94, 0xdb, 0x18, 0x2d,
	0x5b, 0x92, 0xd1, 0x81, 0x54, 0x8b, 0xa0, 0x85, 0xfa, 0x5a, 0xf3, 0x46, 0xd0, 0xd6, 0x72, 0x6d,
	0xae, 0x67, 0x30, 0xe9, 0x13, 0x2f, 0x17, 0xeb, 0x1f, 0xfc, 0x32, 0x7e, 0xb0, 0x08, 0xfa, 0x0d,
	0xf2, 0x09, 0x24, 0x97, 0xd4, 0x06, 0xf1, 0xd5, 0xdd, 0xf3, 0xfb, 0xf3, 0x3c, 0x3f, 0x7e, 0xa8,
	0xb1, 0x00, 0xc1, 0x41, 0x4e, 0x43, 0x1a, 0x45, 0x2b, 0x10, 0xae, 0x15, 0x0a, 0x90, 0x80, 0x51,
	0x1c, 0x31, 0x31, 0xa5, 0x6e, 0xe0, 0x2d, 0x5b, 0x26, 0xf7, 0xe4, 0x79, 0x3c, 0xb3, 0xe6, 0x10,
	0xd8, 0x1c, 0x38, 0xd8, 0x4a, 0x32, 0x8b, 0x17, 0x6a, 0x52, 0x83, 0xfa, 0xe5, 0xd6, 0xd6, 0x56,
	0x49, 0x1e, 0xac, 0x3c, 0x79, 0x01, 0x2b, 0x9b, 0x83, 0xa9, 0x48, 0xf3, 0x8a, 0xfa, 0x9e, 0x4b,
	0x25, 0x88, 0xc8, 0x7e, 0xfd, 0x16, 0xbe, 0xef, 0x1c, 0x80, 0xfb, 0x6c, 0xb3, 0x9d, 0x05, 0xa1,
	0xbc, 0xc9, 0x49, 0xf2, 0xa0, 0xa1, 0xc6, 0x81, 0x4a, 0x7a, 0x58, 0x04, 0x75, 0xd8, 0x65, 0xcc,
	0x22, 0x89, 0x39, 0xaa, 0xb1, 0x80, 0x7a, 0xbe, 0xae, 0xb5, 0xb5, 0xce, 0xa7, 0xe1, 0x51, 0x9a,
	0x18, 0x9f, 0x17, 0x20, 0x82, 0x1d, 0xa2, 0x60, 0xf2, 0xf4, 0x68, 0xec, 0xa2, 0xc1, 0x59, 0xe7,
	0x74, 0xdf, 0x3c, 0xa1, 0xe6, 0xed, 0x3f, 0x73, 0x7b, 0x3a, 0x36, 0xc7, 0xd6, 0xa4, 0xfb, 0x67,
	0xbc, 0xf7, 0x1e, 0x68, 0xad, 0xc1, 0xc9, 0x5d, 0xef, 0x6f, 0xff, 0xbe, 0xfb, 0xcb, 0xc9, 0xf7,
	0xe3, 0x36, 0xaa, 0xc4, 0xc2, 0xd7, 0x3f, 0xa8, 0x33, 0x5f, 0xd2, 0xc4, 0x40, 0xf9, 0x99, 0x58,
	0xf8, 0xc4, 0xc9, 0x28, 0xf2, 0xac, 0xa1, 0x1f, 0x6f, 0x43, 0x46, 0x21, 0x2c, 0x23, 0x76, 0x2c,
	0x68, 0x18, 0x32, 0x81, 0x7f, 0xa2, 0xea, 0x1c, 0x5c, 0xa6, 0xb2, 0xd6, 0x86, 0x5f, 0xd3, 0xc4,
	0xa8, 0xe7, 0x4b, 0x32, 0x94, 0x38, 0x8a, 0xc4, 0x7d, 0x54, 0xcf, 0xde, 0xd1, 0x75, 0xe8, 0x53,
	0x6f, 0x59, 0x1c, 0x6c, 0xa6, 0x89, 0x81, 0x37, 0xda, 0x82, 0x24, 0x4e, 0x59, 0x8a, 0x7f, 0xa3,
	0x1a, 0x13, 0x02, 0x84, 0x5e, 0x51, 0x9e, 0x6f, 0xa5, 0x2e, 0x32, 0x98, 0x38, 0x39, 0x8d, 0x07,
	0xa8, 0xea, 0x52, 0x49, 0xf5, 0x6a, 0x5b, 0xeb, 0xd4, 0x7b, 0x4d, 0x2b, 0x6f, 0xde, 0x5a, 0x37,
	0x6f, 0x8d, 0xb2, 0xe6, 0xcb, 0xf1, 0x32, 0x35, 0x71, 0x94, 0x69, 0xf6, 0x51, 0xc9, 0xfe, 0xbf,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xbf, 0x45, 0x65, 0x3a, 0x02, 0x00, 0x00,
}
