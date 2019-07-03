// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: role.proto

package permission

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
//角色
type Role struct {
	//
	//角色id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//角色所在org
	Org int32 `protobuf:"varint,2,opt,name=org,proto3" json:"org" form:"org"`
	//
	//角色名称
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role" form:"role"`
	//
	//角色下用户列表
	User []string `protobuf:"bytes,4,rep,name=user,proto3" json:"user" form:"user"`
	//
	//角色下权限id列表
	Permission []string `protobuf:"bytes,5,rep,name=permission,proto3" json:"permission" form:"permission"`
	//
	//角色下用户组列表
	UserGroup            []string `protobuf:"bytes,6,rep,name=user_group,json=userGroup,proto3" json:"user_group" form:"user_group"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{0}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *Role) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Role) GetUser() []string {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Role) GetPermission() []string {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Role) GetUserGroup() []string {
	if m != nil {
		return m.UserGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*Role)(nil), "permission.Role")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_48a3ff9f7c9032f8) }

var fileDescriptor_48a3ff9f7c9032f8 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x8f, 0x9a, 0x40,
	0x18, 0x86, 0x2b, 0xa2, 0x09, 0xd3, 0x43, 0x2b, 0xb5, 0x0d, 0xf1, 0x82, 0xa1, 0xa6, 0xb1, 0x89,
	0xc3, 0x34, 0xd5, 0x98, 0xb4, 0x49, 0x0f, 0x92, 0xc6, 0x5e, 0x7a, 0xe2, 0xa8, 0x55, 0x83, 0x80,
	0xd3, 0xc9, 0x82, 0x1f, 0x99, 0x81, 0x35, 0xbb, 0x66, 0xff, 0xe1, 0xfe, 0x06, 0x36, 0xd9, 0x9f,
	0xc0, 0x2f, 0xd8, 0xcc, 0xb0, 0x59, 0xbc, 0xec, 0xed, 0x7b, 0xbf, 0xf7, 0x7d, 0x1f, 0xf8, 0x00,
	0x21, 0x0e, 0x49, 0xec, 0x66, 0x1c, 0x72, 0x30, 0x51, 0x16, 0xf3, 0x94, 0x09, 0xc1, 0xe0, 0x38,
	0xc0, 0x94, 0xe5, 0xff, 0x8b, 0xbd, 0x1b, 0x42, 0x4a, 0x28, 0x50, 0x20, 0x2a, 0xb2, 0x2f, 0x0e,
	0x4a, 0x29, 0xa1, 0xa6, 0xba, 0x3a, 0x98, 0x5f, 0xc4, 0xd3, 0x13, 0xcb, 0xaf, 0xe0, 0x44, 0x28,
	0x60, 0x65, 0xe2, 0xeb, 0x20, 0x61, 0x51, 0x90, 0x03, 0x17, 0xe4, 0x65, 0xac, 0x7b, 0xce, 0xbd,
	0x86, 0x74, 0x1f, 0x92, 0xd8, 0x9c, 0x21, 0x8d, 0x45, 0x56, 0x6b, 0xd8, 0x1a, 0x1b, 0xde, 0xa8,
	0x2a, 0x6d, 0xe3, 0x00, 0x3c, 0xfd, 0xe9, 0xb0, 0xc8, 0x79, 0x7c, 0xb0, 0x3f, 0xa0, 0xde, 0x76,
	0xfd, 0x0d, 0xff, 0x08, 0xf0, 0x61, 0x81, 0x97, 0x9b, 0xf3, 0xf7, 0xd9, 0xdd, 0xc8, 0xd7, 0x58,
	0x64, 0x7e, 0x41, 0x6d, 0xe0, 0xd4, 0xd2, 0x86, 0xad, 0x71, 0xc7, 0xeb, 0x57, 0xa5, 0x8d, 0xea,
	0x1a, 0x70, 0x2a, 0x7b, 0xda, 0xfb, 0x37, 0xbe, 0x0c, 0x98, 0x9f, 0x91, 0x2e, 0xef, 0xb4, 0xda,
	0x8a, 0xff, 0xae, 0x2a, 0xed, 0xb7, 0x75, 0x50, 0x6e, 0x1d, 0x5f, 0x99, 0xe6, 0x2f, 0xa4, 0x17,
	0x22, 0xe6, 0x96, 0x3e, 0x6c, 0x8f, 0x0d, 0xef, 0x6b, 0x13, 0x92, 0x5b, 0x89, 0xfb, 0x84, 0xfa,
	0xdb, 0xf5, 0x02, 0xaf, 0x02, 0x7c, 0xbb, 0xc3, 0x9b, 0x7f, 0xa7, 0xf3, 0x74, 0x32, 0x97, 0x6f,
	0xa2, 0x6a, 0xe6, 0x5f, 0x74, 0xf1, 0xfd, 0xac, 0x8e, 0x82, 0x4c, 0xaa, 0xd2, 0xee, 0xd5, 0x90,
	0xc6, 0x7b, 0xf5, 0xa2, 0x8b, 0xbe, 0x39, 0x43, 0x48, 0x52, 0x77, 0x94, 0x43, 0x91, 0x59, 0x5d,
	0x45, 0xfb, 0xd8, 0xd0, 0x1a, 0xcf, 0xf1, 0x0d, 0x29, 0xfe, 0xc8, 0xd9, 0x5b, 0xae, 0x7e, 0x53,
	0x70, 0xe3, 0x40, 0xdc, 0x40, 0x26, 0xdc, 0x04, 0xc2, 0x20, 0x21, 0x21, 0x1c, 0x73, 0x1e, 0x84,
	0xb9, 0xa8, 0x7f, 0x21, 0x8f, 0x33, 0xc0, 0x29, 0x44, 0x71, 0x22, 0xc8, 0x73, 0x90, 0x28, 0x49,
	0x9a, 0xa7, 0xef, 0xbb, 0x2a, 0x3a, 0x7d, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x78, 0xd2, 0x8b, 0x4c,
	0x1e, 0x02, 0x00, 0x00,
}
