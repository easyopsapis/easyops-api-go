// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_users_id_nick.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//ListUsersIdNick请求
type ListUsersIdNickRequest struct {
	//
	//状态
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state" form:"state"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersIdNickRequest) Reset()         { *m = ListUsersIdNickRequest{} }
func (m *ListUsersIdNickRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersIdNickRequest) ProtoMessage()    {}
func (*ListUsersIdNickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebdcb778d4ea3134, []int{0}
}
func (m *ListUsersIdNickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersIdNickRequest.Unmarshal(m, b)
}
func (m *ListUsersIdNickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersIdNickRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersIdNickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersIdNickRequest.Merge(m, src)
}
func (m *ListUsersIdNickRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersIdNickRequest.Size(m)
}
func (m *ListUsersIdNickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersIdNickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersIdNickRequest proto.InternalMessageInfo

func (m *ListUsersIdNickRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

//
//ListUsersIdNickApi返回
type ListUsersIdNickResponseWrapper struct {
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
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListUsersIdNickResponseWrapper) Reset()         { *m = ListUsersIdNickResponseWrapper{} }
func (m *ListUsersIdNickResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListUsersIdNickResponseWrapper) ProtoMessage()    {}
func (*ListUsersIdNickResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebdcb778d4ea3134, []int{1}
}
func (m *ListUsersIdNickResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersIdNickResponseWrapper.Unmarshal(m, b)
}
func (m *ListUsersIdNickResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersIdNickResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListUsersIdNickResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersIdNickResponseWrapper.Merge(m, src)
}
func (m *ListUsersIdNickResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListUsersIdNickResponseWrapper.Size(m)
}
func (m *ListUsersIdNickResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersIdNickResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersIdNickResponseWrapper proto.InternalMessageInfo

func (m *ListUsersIdNickResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListUsersIdNickResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListUsersIdNickResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListUsersIdNickResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListUsersIdNickRequest)(nil), "user_admin.ListUsersIdNickRequest")
	proto.RegisterType((*ListUsersIdNickResponseWrapper)(nil), "user_admin.ListUsersIdNickResponseWrapper")
}

func init() { proto.RegisterFile("list_users_id_nick.proto", fileDescriptor_ebdcb778d4ea3134) }

var fileDescriptor_ebdcb778d4ea3134 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xe9, 0xf7, 0x75, 0x04, 0x53, 0x41, 0xc9, 0x62, 0x2c, 0x83, 0xd8, 0x21, 0x82, 0xcc,
	0xc6, 0x0e, 0xe8, 0x46, 0xc4, 0x85, 0x0c, 0xb8, 0x10, 0xc4, 0x45, 0x44, 0x5c, 0x96, 0xb4, 0xcd,
	0xd4, 0x30, 0x6d, 0x6f, 0x4d, 0x52, 0xf0, 0x69, 0xbb, 0xf3, 0x05, 0xfa, 0x04, 0x92, 0x1b, 0x65,
	0x06, 0x5d, 0x25, 0xf7, 0x9c, 0xdf, 0xfd, 0xc3, 0x21, 0x71, 0xad, 0x8c, 0xcd, 0x7a, 0x23, 0xb5,
	0xc9, 0x54, 0x99, 0xb5, 0xaa, 0xd8, 0xa4, 0x9d, 0x06, 0x0b, 0x94, 0x38, 0x31, 0x13, 0x65, 0xa3,
	0xda, 0xd9, 0x45, 0xa5, 0xec, 0x5b, 0x9f, 0xa7, 0x05, 0x34, 0xcb, 0x0a, 0x2a, 0x58, 0x22, 0x92,
	0xf7, 0x6b, 0xac, 0xb0, 0xc0, 0x9f, 0x6f, 0x9d, 0x9d, 0x54, 0x00, 0x55, 0x2d, 0xb7, 0x94, 0xb1,
	0xba, 0x2f, 0xac, 0x77, 0xd9, 0x1d, 0x99, 0x3e, 0x2a, 0x63, 0x5f, 0xdc, 0xce, 0x87, 0xf2, 0x49,
	0x15, 0x1b, 0x2e, 0xdf, 0x7b, 0x69, 0x2c, 0x3d, 0x27, 0x13, 0x63, 0x85, 0x95, 0x71, 0x30, 0x0f,
	0x16, 0xfb, 0xab, 0xa3, 0x71, 0x48, 0x0e, 0xd6, 0xa0, 0x9b, 0x1b, 0x86, 0x32, 0xe3, 0xde, 0x66,
	0x9f, 0x01, 0x39, 0xfd, 0x33, 0xc2, 0x74, 0xd0, 0x1a, 0xf9, 0xaa, 0x45, 0xd7, 0x49, 0x4d, 0xcf,
	0x48, 0x58, 0x40, 0xe9, 0x27, 0x4d, 0x56, 0x87, 0xe3, 0x90, 0x44, 0x7e, 0x92, 0x53, 0x19, 0x47,
	0x93, 0x5e, 0x93, 0xc8, 0xbd, 0xf7, 0x1f, 0x5d, 0x2d, 0x54, 0x1b, 0xff, 0xc3, 0xad, 0xd3, 0x71,
	0x48, 0xe8, 0x96, 0xfd, 0x36, 0x19, 0xdf, 0x45, 0xdd, 0xa5, 0x52, 0x6b, 0xd0, 0xf1, 0xff, 0xdf,
	0x97, 0xa2, 0xcc, 0xb8, 0xb7, 0xe9, 0x2d, 0x09, 0x4b, 0x61, 0x45, 0x1c, 0xce, 0x83, 0x45, 0x74,
	0x79, 0x9c, 0xfa, 0x60, 0xd2, 0x9f, 0x60, 0xd2, 0x67, 0x0c, 0x66, 0xf7, 0x3e, 0x87, 0x33, 0x8e,
	0x5d, 0xf9, 0x1e, 0x72, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x41, 0x71, 0xf4, 0xa5,
	0x01, 0x00, 0x00,
}