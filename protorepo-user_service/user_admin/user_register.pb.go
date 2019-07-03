// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user_register.proto

package user_admin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//UserRegister请求
type UserRegisterRequest struct {
	//
	//用户名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password" form:"password"`
	//
	//邮箱
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email" form:"email"`
	//
	//org
	Org                  int32    `protobuf:"varint,4,opt,name=org,proto3" json:"org" form:"org"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a746fe9a0db45e37, []int{0}
}
func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (m *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(m, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

func (m *UserRegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRegisterRequest) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

//
//UserRegister返回
type UserRegisterResponse struct {
	//
	//用户名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//邮箱
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email" form:"email"`
	//
	//org
	Org                  int32    `protobuf:"varint,3,opt,name=org,proto3" json:"org" form:"org"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterResponse) Reset()         { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()    {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a746fe9a0db45e37, []int{1}
}
func (m *UserRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResponse.Unmarshal(m, b)
}
func (m *UserRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResponse.Marshal(b, m, deterministic)
}
func (m *UserRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResponse.Merge(m, src)
}
func (m *UserRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResponse.Size(m)
}
func (m *UserRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResponse proto.InternalMessageInfo

func (m *UserRegisterResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRegisterResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRegisterResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

//
//UserRegisterApi返回
type UserRegisterResponseWrapper struct {
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
	Data                 *UserRegisterResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserRegisterResponseWrapper) Reset()         { *m = UserRegisterResponseWrapper{} }
func (m *UserRegisterResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResponseWrapper) ProtoMessage()    {}
func (*UserRegisterResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a746fe9a0db45e37, []int{2}
}
func (m *UserRegisterResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResponseWrapper.Unmarshal(m, b)
}
func (m *UserRegisterResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UserRegisterResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResponseWrapper.Merge(m, src)
}
func (m *UserRegisterResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResponseWrapper.Size(m)
}
func (m *UserRegisterResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResponseWrapper proto.InternalMessageInfo

func (m *UserRegisterResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserRegisterResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UserRegisterResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UserRegisterResponseWrapper) GetData() *UserRegisterResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "user_admin.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "user_admin.UserRegisterResponse")
	proto.RegisterType((*UserRegisterResponseWrapper)(nil), "user_admin.UserRegisterResponseWrapper")
}

func init() { proto.RegisterFile("user_register.proto", fileDescriptor_a746fe9a0db45e37) }

var fileDescriptor_a746fe9a0db45e37 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x61, 0x6b, 0xd3, 0x40,
	0x18, 0xc7, 0x49, 0xdb, 0x88, 0x3e, 0x15, 0x37, 0x6e, 0x22, 0x75, 0xbe, 0x48, 0x39, 0xa7, 0x74,
	0xb8, 0x4b, 0xc7, 0x44, 0xd9, 0x14, 0xc1, 0x15, 0xf6, 0x01, 0x3c, 0x10, 0x61, 0x6b, 0x37, 0xae,
	0xcb, 0xed, 0x0c, 0x36, 0xbd, 0x78, 0x97, 0x18, 0xd9, 0xd8, 0x97, 0xf0, 0xf3, 0x49, 0x04, 0xdf,
	0xfa, 0x2e, 0x9f, 0x40, 0xee, 0x2e, 0x6b, 0x23, 0xd4, 0x57, 0x82, 0xaf, 0x72, 0xf7, 0xff, 0x3f,
	0xbf, 0xcb, 0xf3, 0x7f, 0x2e, 0x81, 0x8d, 0x5c, 0x73, 0x75, 0xa6, 0xb8, 0x88, 0x75, 0xc6, 0x55,
	0x98, 0x2a, 0x99, 0x49, 0x04, 0x56, 0x64, 0x51, 0x12, 0xcf, 0x37, 0x89, 0x88, 0xb3, 0x8f, 0xf9,
	0x34, 0x3c, 0x97, 0xc9, 0x50, 0x48, 0x21, 0x87, 0xb6, 0x64, 0x9a, 0x5f, 0xd8, 0x9d, 0xdd, 0xd8,
	0x95, 0x43, 0x37, 0x5f, 0x36, 0xca, 0x93, 0x22, 0xce, 0x3e, 0xc9, 0x62, 0x28, 0x24, 0xb1, 0x26,
	0xf9, 0xc2, 0x66, 0x71, 0xc4, 0x32, 0xa9, 0xf4, 0x70, 0xb1, 0x74, 0x1c, 0xfe, 0xd6, 0x82, 0x8d,
	0xf7, 0x9a, 0x2b, 0x5a, 0x77, 0x42, 0xf9, 0xe7, 0x9c, 0xeb, 0x0c, 0x1d, 0x40, 0x67, 0xce, 0x12,
	0xde, 0xf3, 0xfa, 0xde, 0xe0, 0xce, 0xe8, 0x49, 0x55, 0x06, 0xdd, 0x0b, 0xa9, 0x92, 0x57, 0xd8,
	0xa8, 0xf8, 0xe7, 0x8f, 0x00, 0xc1, 0xfa, 0xe9, 0x09, 0x23, 0x97, 0x87, 0xe4, 0x78, 0x32, 0x2e,
	0xae, 0x76, 0x77, 0xae, 0xb7, 0xa8, 0x45, 0xd0, 0x0b, 0xb8, 0x9d, 0x32, 0xad, 0x0b, 0xa9, 0xa2,
	0x5e, 0xcb, 0xe2, 0x0f, 0xab, 0x32, 0x58, 0x73, 0xf8, 0x8d, 0x63, 0x8e, 0x68, 0xa5, 0x3e, 0x5d,
	0x94, 0x22, 0x01, 0x3e, 0x4f, 0x58, 0x3c, 0xeb, 0xb5, 0x2d, 0xf3, 0xae, 0x2a, 0x83, 0xbb, 0x8e,
	0xb1, 0xb2, 0x01, 0xde, 0xc0, 0xeb, 0xd3, 0xc1, 0xc9, 0x21, 0x39, 0x66, 0xe4, 0x72, 0x97, 0x1c,
	0x9c, 0x8d, 0xc9, 0x38, 0x9c, 0x6c, 0x3f, 0x1b, 0xbf, 0x5d, 0x25, 0x86, 0x37, 0xe2, 0xe4, 0x6a,
	0x6f, 0x67, 0xff, 0x7a, 0x7b, 0x8b, 0xba, 0xf3, 0x51, 0x1f, 0xda, 0x52, 0x89, 0x5e, 0xa7, 0xef,
	0x0d, 0xfc, 0xd1, 0xbd, 0xaa, 0x0c, 0xc0, 0xbd, 0x46, 0x2a, 0x81, 0xa9, 0xb1, 0xf0, 0x77, 0x0f,
	0xee, 0xff, 0x39, 0x14, 0x9d, 0xca, 0xb9, 0xe6, 0xff, 0x32, 0x95, 0x45, 0xbc, 0xd6, 0xff, 0x89,
	0xd7, 0xfe, 0x7b, 0xbc, 0x5f, 0x1e, 0x3c, 0x5a, 0x15, 0xef, 0x83, 0x62, 0x69, 0xca, 0x15, 0x7a,
	0x0c, 0x9d, 0x73, 0x19, 0xb9, 0x94, 0xfe, 0x68, 0x6d, 0x99, 0xd2, 0xa8, 0x98, 0x5a, 0x13, 0xed,
	0x43, 0xd7, 0x3c, 0x8f, 0xbe, 0xa6, 0x33, 0x16, 0xcf, 0xeb, 0x54, 0x0f, 0xaa, 0x32, 0x40, 0xcb,
	0xda, 0xda, 0xc4, 0xb4, 0x59, 0x8a, 0x9e, 0x82, 0xcf, 0x95, 0x92, 0xaa, 0xbe, 0xe8, 0xf5, 0xc6,
	0x24, 0x8c, 0x8c, 0xa9, 0xb3, 0xd1, 0x11, 0x74, 0x22, 0x96, 0x31, 0x7b, 0x51, 0xdd, 0xbd, 0x7e,
	0xb8, 0xfc, 0x39, 0xc2, 0x55, 0xdd, 0x37, 0x1b, 0x35, 0x1c, 0xa6, 0x16, 0x9f, 0xde, 0xb2, 0x1f,
	0xfa, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x2b, 0x21, 0xa8, 0x72, 0x03, 0x00, 0x00,
}