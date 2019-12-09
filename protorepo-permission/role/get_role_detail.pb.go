// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_role_detail.proto

package role

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//GetRoleDetail请求
type GetRoleDetailRequest struct {
	//
	//角色Id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoleDetailRequest) Reset()         { *m = GetRoleDetailRequest{} }
func (m *GetRoleDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoleDetailRequest) ProtoMessage()    {}
func (*GetRoleDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0f022be40ca18f, []int{0}
}
func (m *GetRoleDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleDetailRequest.Unmarshal(m, b)
}
func (m *GetRoleDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetRoleDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleDetailRequest.Merge(m, src)
}
func (m *GetRoleDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoleDetailRequest.Size(m)
}
func (m *GetRoleDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleDetailRequest proto.InternalMessageInfo

func (m *GetRoleDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetRoleDetail返回
type GetRoleDetailResponse struct {
	//
	//角色禁用菜单列表
	ForbiddenMenu []string `protobuf:"bytes,1,rep,name=forbidden_menu,json=forbiddenMenu,proto3" json:"forbidden_menu" form:"forbidden_menu"`
	//
	//角色id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//角色所在org
	Org int32 `protobuf:"varint,3,opt,name=org,proto3" json:"org" form:"org"`
	//
	//角色名称
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role" form:"role"`
	//
	//角色下用户列表
	User []string `protobuf:"bytes,5,rep,name=user,proto3" json:"user" form:"user"`
	//
	//角色下权限id列表
	Permission []string `protobuf:"bytes,6,rep,name=permission,proto3" json:"permission" form:"permission"`
	//
	//角色下用户组列表
	UserGroup            []string `protobuf:"bytes,7,rep,name=user_group,json=userGroup,proto3" json:"user_group" form:"user_group"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoleDetailResponse) Reset()         { *m = GetRoleDetailResponse{} }
func (m *GetRoleDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetRoleDetailResponse) ProtoMessage()    {}
func (*GetRoleDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0f022be40ca18f, []int{1}
}
func (m *GetRoleDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleDetailResponse.Unmarshal(m, b)
}
func (m *GetRoleDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetRoleDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleDetailResponse.Merge(m, src)
}
func (m *GetRoleDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetRoleDetailResponse.Size(m)
}
func (m *GetRoleDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleDetailResponse proto.InternalMessageInfo

func (m *GetRoleDetailResponse) GetForbiddenMenu() []string {
	if m != nil {
		return m.ForbiddenMenu
	}
	return nil
}

func (m *GetRoleDetailResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRoleDetailResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetRoleDetailResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GetRoleDetailResponse) GetUser() []string {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetRoleDetailResponse) GetPermission() []string {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *GetRoleDetailResponse) GetUserGroup() []string {
	if m != nil {
		return m.UserGroup
	}
	return nil
}

//
//GetRoleDetailApi返回
type GetRoleDetailResponseWrapper struct {
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
	Data                 *GetRoleDetailResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetRoleDetailResponseWrapper) Reset()         { *m = GetRoleDetailResponseWrapper{} }
func (m *GetRoleDetailResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetRoleDetailResponseWrapper) ProtoMessage()    {}
func (*GetRoleDetailResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0f022be40ca18f, []int{2}
}
func (m *GetRoleDetailResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleDetailResponseWrapper.Unmarshal(m, b)
}
func (m *GetRoleDetailResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleDetailResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetRoleDetailResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleDetailResponseWrapper.Merge(m, src)
}
func (m *GetRoleDetailResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetRoleDetailResponseWrapper.Size(m)
}
func (m *GetRoleDetailResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleDetailResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleDetailResponseWrapper proto.InternalMessageInfo

func (m *GetRoleDetailResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetRoleDetailResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetRoleDetailResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetRoleDetailResponseWrapper) GetData() *GetRoleDetailResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRoleDetailRequest)(nil), "role.GetRoleDetailRequest")
	proto.RegisterType((*GetRoleDetailResponse)(nil), "role.GetRoleDetailResponse")
	proto.RegisterType((*GetRoleDetailResponseWrapper)(nil), "role.GetRoleDetailResponseWrapper")
}

func init() { proto.RegisterFile("get_role_detail.proto", fileDescriptor_7a0f022be40ca18f) }

var fileDescriptor_7a0f022be40ca18f = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd3, 0x3e,
	0x18, 0xff, 0xf7, 0x6d, 0x7f, 0xc5, 0xe5, 0x65, 0x33, 0x2b, 0x0a, 0x03, 0x29, 0xc5, 0xab, 0xa6,
	0x1e, 0xe6, 0x74, 0x74, 0xd5, 0xc4, 0x38, 0x6d, 0x15, 0x30, 0x09, 0x8d, 0x8b, 0x2f, 0x48, 0x54,
	0x5d, 0x94, 0x2e, 0x6e, 0xb0, 0x48, 0xe2, 0xe0, 0x24, 0x0c, 0x31, 0xed, 0x1b, 0xf1, 0x99, 0x82,
	0xc4, 0x85, 0x7b, 0x3e, 0x01, 0xf2, 0x93, 0xaa, 0x09, 0x68, 0x1c, 0x38, 0xe5, 0xf1, 0xef, 0x4d,
	0xbf, 0xf8, 0x31, 0xea, 0xf9, 0x3c, 0x75, 0x94, 0x0c, 0xb8, 0xe3, 0xf1, 0xd4, 0x15, 0x81, 0x1d,
	0x2b, 0x99, 0x4a, 0xdc, 0xd6, 0xd0, 0x0e, 0xf5, 0x45, 0xfa, 0x21, 0x5b, 0xd8, 0x97, 0x32, 0x1c,
	0xf9, 0xd2, 0x97, 0x23, 0x20, 0x17, 0xd9, 0x12, 0x4e, 0x70, 0x80, 0xa9, 0x34, 0xed, 0x1c, 0xd5,
	0xe4, 0xe1, 0x95, 0x48, 0x3f, 0xca, 0xab, 0x91, 0x2f, 0x29, 0x90, 0xf4, 0xb3, 0x1b, 0x08, 0xcf,
	0x4d, 0xa5, 0x4a, 0x46, 0xeb, 0xb1, 0xf4, 0x91, 0x73, 0xb4, 0x7d, 0xc6, 0x53, 0x26, 0x03, 0xfe,
	0x12, 0x3a, 0x30, 0xfe, 0x29, 0xe3, 0x49, 0x8a, 0x27, 0xa8, 0x29, 0x3c, 0xb3, 0xd1, 0x6f, 0x0c,
	0x8d, 0xe9, 0xa0, 0xc8, 0x2d, 0x63, 0x29, 0x55, 0xf8, 0x82, 0x08, 0x8f, 0xfc, 0xf8, 0x6e, 0x3d,
	0x40, 0x5b, 0x17, 0xb3, 0x03, 0x7a, 0xec, 0xd2, 0xe5, 0x29, 0x7d, 0x3d, 0xbf, 0x1e, 0x4f, 0x6e,
	0x06, 0xac, 0x29, 0x3c, 0xf2, 0xad, 0x85, 0x7a, 0x7f, 0xc4, 0x25, 0xb1, 0x8c, 0x12, 0x8e, 0x4f,
	0xd0, 0xbd, 0xa5, 0x54, 0x0b, 0xe1, 0x79, 0x3c, 0x72, 0x42, 0x1e, 0x65, 0x66, 0xa3, 0xdf, 0x1a,
	0x1a, 0xd3, 0x47, 0x45, 0x6e, 0xf5, 0xca, 0xec, 0xdf, 0x79, 0xc2, 0xee, 0xae, 0x81, 0xb7, 0x3c,
	0xca, 0x56, 0x8d, 0x9a, 0xff, 0xd6, 0x08, 0xef, 0xa1, 0x96, 0x54, 0xbe, 0xd9, 0xea, 0x37, 0x86,
	0x9d, 0xe9, 0x76, 0x91, 0x5b, 0xa8, 0xb4, 0x49, 0xe5, 0x6b, 0x5f, 0x73, 0xf3, 0x3f, 0xa6, 0x05,
	0x78, 0x17, 0xc1, 0xb5, 0x9b, 0x6d, 0xc8, 0xbf, 0x5f, 0xe4, 0x56, 0xb7, 0x14, 0x6a, 0x94, 0x30,
	0x20, 0xf1, 0x1b, 0xd4, 0xce, 0x12, 0xae, 0xcc, 0x0e, 0x54, 0x3f, 0xaa, 0x44, 0x1a, 0xd5, 0x71,
	0xbb, 0xe8, 0xe9, 0xc5, 0xcc, 0xa5, 0x5f, 0x4f, 0xe9, 0xfb, 0x03, 0x7a, 0x3c, 0x9f, 0xd9, 0xeb,
	0xd9, 0xa1, 0xf3, 0xeb, 0xf1, 0xfe, 0xe1, 0xb3, 0x9b, 0x01, 0x83, 0x0c, 0x7c, 0x8e, 0x50, 0xcc,
	0x55, 0x28, 0x92, 0x44, 0xc8, 0xc8, 0xdc, 0x80, 0xc4, 0xfd, 0x22, 0xb7, 0xb6, 0xca, 0xc4, 0x8a,
	0xfb, 0xeb, 0xef, 0xd5, 0xfc, 0x78, 0x82, 0x90, 0x4e, 0x75, 0x7c, 0x25, 0xb3, 0xd8, 0xfc, 0x1f,
	0xd2, 0x7a, 0x55, 0x5a, 0xc5, 0x11, 0x66, 0xe8, 0xc3, 0x19, 0xcc, 0x3f, 0x1b, 0xe8, 0xc9, 0xad,
	0xeb, 0x7a, 0xa7, 0xdc, 0x38, 0xe6, 0x4a, 0xdf, 0xca, 0xa5, 0xf4, 0x38, 0xbc, 0x83, 0x4e, 0xfd,
	0x56, 0x34, 0x4a, 0x18, 0x90, 0xf8, 0x39, 0xea, 0xea, 0xef, 0xab, 0x2f, 0x71, 0xe0, 0x8a, 0x68,
	0xb5, 0xa1, 0x87, 0x45, 0x6e, 0xe1, 0x4a, 0xbb, 0x22, 0x09, 0xab, 0x4b, 0xf1, 0x1e, 0xea, 0x70,
	0xa5, 0xa4, 0x82, 0xf5, 0x18, 0xd3, 0xcd, 0x22, 0xb7, 0xee, 0x94, 0x1e, 0x80, 0x09, 0x2b, 0x69,
	0x7c, 0x82, 0xda, 0x9e, 0x9b, 0xba, 0xb0, 0x9c, 0xee, 0xf8, 0xb1, 0xad, 0x97, 0x61, 0xdf, 0x5a,
	0xbc, 0xde, 0x51, 0x5b, 0x08, 0x03, 0xe7, 0x62, 0x03, 0x5e, 0xfb, 0xe1, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0xf3, 0x9e, 0xf1, 0x73, 0x03, 0x00, 0x00,
}
