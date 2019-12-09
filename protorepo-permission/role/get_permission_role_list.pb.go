// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_permission_role_list.proto

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
//GetPermissionRoleList请求
type GetPermissionRoleListRequest struct {
	//
	//用户名
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user" form:"user"`
	//
	//角色名
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role" form:"role"`
	//
	//分页
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//sdk限制字段使用
	XXX_RestFieldMask    []string `protobuf:"bytes,5,rep,name=XXX_RestFieldMask,json=XXXRestFieldMask,proto3" json:"XXX_RestFieldMask" form:"XXX_RestFieldMask"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPermissionRoleListRequest) Reset()         { *m = GetPermissionRoleListRequest{} }
func (m *GetPermissionRoleListRequest) String() string { return proto.CompactTextString(m) }
func (*GetPermissionRoleListRequest) ProtoMessage()    {}
func (*GetPermissionRoleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1ea9958bee5f86, []int{0}
}
func (m *GetPermissionRoleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionRoleListRequest.Unmarshal(m, b)
}
func (m *GetPermissionRoleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionRoleListRequest.Marshal(b, m, deterministic)
}
func (m *GetPermissionRoleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionRoleListRequest.Merge(m, src)
}
func (m *GetPermissionRoleListRequest) XXX_Size() int {
	return xxx_messageInfo_GetPermissionRoleListRequest.Size(m)
}
func (m *GetPermissionRoleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionRoleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionRoleListRequest proto.InternalMessageInfo

func (m *GetPermissionRoleListRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GetPermissionRoleListRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GetPermissionRoleListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetPermissionRoleListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetPermissionRoleListRequest) GetXXX_RestFieldMask() []string {
	if m != nil {
		return m.XXX_RestFieldMask
	}
	return nil
}

//
//GetPermissionRoleList返回
type GetPermissionRoleListResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//角色列表
	Data                 []*GetPermissionRoleListResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GetPermissionRoleListResponse) Reset()         { *m = GetPermissionRoleListResponse{} }
func (m *GetPermissionRoleListResponse) String() string { return proto.CompactTextString(m) }
func (*GetPermissionRoleListResponse) ProtoMessage()    {}
func (*GetPermissionRoleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1ea9958bee5f86, []int{1}
}
func (m *GetPermissionRoleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionRoleListResponse.Unmarshal(m, b)
}
func (m *GetPermissionRoleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionRoleListResponse.Marshal(b, m, deterministic)
}
func (m *GetPermissionRoleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionRoleListResponse.Merge(m, src)
}
func (m *GetPermissionRoleListResponse) XXX_Size() int {
	return xxx_messageInfo_GetPermissionRoleListResponse.Size(m)
}
func (m *GetPermissionRoleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionRoleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionRoleListResponse proto.InternalMessageInfo

func (m *GetPermissionRoleListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPermissionRoleListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetPermissionRoleListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetPermissionRoleListResponse) GetData() []*GetPermissionRoleListResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetPermissionRoleListResponse_Data struct {
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

func (m *GetPermissionRoleListResponse_Data) Reset()         { *m = GetPermissionRoleListResponse_Data{} }
func (m *GetPermissionRoleListResponse_Data) String() string { return proto.CompactTextString(m) }
func (*GetPermissionRoleListResponse_Data) ProtoMessage()    {}
func (*GetPermissionRoleListResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1ea9958bee5f86, []int{1, 0}
}
func (m *GetPermissionRoleListResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionRoleListResponse_Data.Unmarshal(m, b)
}
func (m *GetPermissionRoleListResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionRoleListResponse_Data.Marshal(b, m, deterministic)
}
func (m *GetPermissionRoleListResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionRoleListResponse_Data.Merge(m, src)
}
func (m *GetPermissionRoleListResponse_Data) XXX_Size() int {
	return xxx_messageInfo_GetPermissionRoleListResponse_Data.Size(m)
}
func (m *GetPermissionRoleListResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionRoleListResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionRoleListResponse_Data proto.InternalMessageInfo

func (m *GetPermissionRoleListResponse_Data) GetForbiddenMenu() []string {
	if m != nil {
		return m.ForbiddenMenu
	}
	return nil
}

func (m *GetPermissionRoleListResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPermissionRoleListResponse_Data) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetPermissionRoleListResponse_Data) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GetPermissionRoleListResponse_Data) GetUser() []string {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetPermissionRoleListResponse_Data) GetPermission() []string {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *GetPermissionRoleListResponse_Data) GetUserGroup() []string {
	if m != nil {
		return m.UserGroup
	}
	return nil
}

//
//GetPermissionRoleListApi返回
type GetPermissionRoleListResponseWrapper struct {
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
	Data                 *GetPermissionRoleListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetPermissionRoleListResponseWrapper) Reset()         { *m = GetPermissionRoleListResponseWrapper{} }
func (m *GetPermissionRoleListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetPermissionRoleListResponseWrapper) ProtoMessage()    {}
func (*GetPermissionRoleListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1ea9958bee5f86, []int{2}
}
func (m *GetPermissionRoleListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionRoleListResponseWrapper.Unmarshal(m, b)
}
func (m *GetPermissionRoleListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionRoleListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetPermissionRoleListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionRoleListResponseWrapper.Merge(m, src)
}
func (m *GetPermissionRoleListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetPermissionRoleListResponseWrapper.Size(m)
}
func (m *GetPermissionRoleListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionRoleListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionRoleListResponseWrapper proto.InternalMessageInfo

func (m *GetPermissionRoleListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPermissionRoleListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetPermissionRoleListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetPermissionRoleListResponseWrapper) GetData() *GetPermissionRoleListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPermissionRoleListRequest)(nil), "role.GetPermissionRoleListRequest")
	proto.RegisterType((*GetPermissionRoleListResponse)(nil), "role.GetPermissionRoleListResponse")
	proto.RegisterType((*GetPermissionRoleListResponse_Data)(nil), "role.GetPermissionRoleListResponse.Data")
	proto.RegisterType((*GetPermissionRoleListResponseWrapper)(nil), "role.GetPermissionRoleListResponseWrapper")
}

func init() { proto.RegisterFile("get_permission_role_list.proto", fileDescriptor_fc1ea9958bee5f86) }

var fileDescriptor_fc1ea9958bee5f86 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0xa5, 0x6d, 0x3a, 0xa8, 0xcb, 0x4f, 0x67, 0x18, 0x0a, 0xd5, 0x20, 0x25, 0xab, 0xa6, 0x22,
	0x2d, 0xed, 0xe8, 0xa6, 0xc1, 0xb8, 0x62, 0x15, 0x6c, 0x80, 0x36, 0x09, 0x99, 0x0b, 0x2a, 0xa6,
	0x2d, 0x72, 0x17, 0x37, 0x58, 0x6b, 0xeb, 0x60, 0x3b, 0x0c, 0x6d, 0xda, 0x63, 0x70, 0xc5, 0xbb,
	0xf0, 0x28, 0x41, 0xe2, 0x11, 0x72, 0xc3, 0x2d, 0xb2, 0xd3, 0x25, 0x9d, 0xd0, 0x0a, 0x13, 0x57,
	0xb1, 0xbf, 0xf3, 0x9d, 0x53, 0x9f, 0xcf, 0xc7, 0x05, 0x0f, 0x7c, 0x22, 0xdd, 0x80, 0xf0, 0x21,
	0x15, 0x82, 0xb2, 0x91, 0xcb, 0xd9, 0x80, 0xb8, 0x03, 0x2a, 0x64, 0x33, 0xe0, 0x4c, 0x32, 0x68,
	0xa8, 0x42, 0xd5, 0xf1, 0xa9, 0xfc, 0x18, 0xf6, 0x9a, 0x07, 0x6c, 0xd8, 0xf2, 0x99, 0xcf, 0x5a,
	0x1a, 0xec, 0x85, 0x7d, 0xbd, 0xd3, 0x1b, 0xbd, 0x4a, 0x48, 0xd5, 0xb5, 0x89, 0xf6, 0xe1, 0x11,
	0x95, 0x87, 0xec, 0xa8, 0xe5, 0x33, 0x47, 0x83, 0xce, 0x67, 0x3c, 0xa0, 0x1e, 0x96, 0x8c, 0x8b,
	0x56, 0xba, 0x4c, 0x78, 0xf6, 0xf7, 0x3c, 0x98, 0xdf, 0x22, 0xf2, 0x6d, 0x7a, 0x1c, 0xc4, 0x06,
	0x64, 0x9b, 0x0a, 0x89, 0xc8, 0xa7, 0x90, 0x08, 0x09, 0xdf, 0x00, 0x23, 0x14, 0x84, 0x9b, 0xb9,
	0x5a, 0xae, 0x51, 0xea, 0xac, 0xc5, 0x91, 0x55, 0xee, 0x33, 0x3e, 0x7c, 0x66, 0xab, 0xaa, 0xfd,
	0xf3, 0x87, 0xb5, 0x00, 0x1e, 0xee, 0xef, 0x62, 0xe7, 0x78, 0xc3, 0xf9, 0xb0, 0xec, 0xac, 0xef,
	0xed, 0x36, 0xd3, 0xb5, 0xeb, 0xec, 0x9d, 0xb4, 0x97, 0x56, 0x1e, 0x9f, 0xd6, 0x91, 0xd6, 0x80,
	0x0b, 0x40, 0x7b, 0x33, 0xf3, 0x5a, 0xeb, 0x56, 0xa6, 0xa5, 0xaa, 0x36, 0xd2, 0x20, 0x7c, 0x04,
	0x8c, 0x00, 0xfb, 0xc4, 0x2c, 0xd4, 0x72, 0x8d, 0x62, 0x67, 0x2e, 0x6b, 0x52, 0x55, 0xf5, 0x83,
	0xf9, 0xca, 0x15, 0xa4, 0x5b, 0xe0, 0x13, 0x50, 0x52, 0x5f, 0x57, 0xd0, 0x63, 0x62, 0x1a, 0xba,
	0xbf, 0x1a, 0x47, 0x56, 0x25, 0xeb, 0xd7, 0xd0, 0x19, 0xe9, 0x9a, 0xaa, 0xbc, 0xa3, 0xc7, 0x04,
	0xbe, 0x06, 0xb3, 0xdd, 0x6e, 0xd7, 0x45, 0x44, 0xc8, 0x4d, 0x4a, 0x06, 0xde, 0x0e, 0x16, 0x87,
	0x66, 0xb1, 0x56, 0x68, 0x94, 0x3a, 0xf3, 0x71, 0x64, 0x99, 0x89, 0xc0, 0x1f, 0x2d, 0x36, 0xaa,
	0x74, 0xbb, 0xdd, 0xf3, 0xa5, 0xaf, 0x45, 0x70, 0xff, 0x82, 0x01, 0x8a, 0x80, 0x8d, 0x04, 0x51,
	0xae, 0x0f, 0x98, 0x47, 0xf4, 0x04, 0x8b, 0x93, 0xae, 0x55, 0xd5, 0x46, 0x1a, 0x4c, 0x5d, 0xe7,
	0x2f, 0xe9, 0xba, 0x70, 0x09, 0xd7, 0x3b, 0xc0, 0xf0, 0xb0, 0xc4, 0xa6, 0x51, 0x2b, 0x34, 0xca,
	0xed, 0x46, 0x53, 0x8d, 0xbb, 0x39, 0xf5, 0xec, 0xcd, 0x17, 0x58, 0xe2, 0xc9, 0x23, 0x2b, 0xbe,
	0x8d, 0xb4, 0x4c, 0xf5, 0x5b, 0x01, 0x18, 0x0a, 0x87, 0xcf, 0xc1, 0xcd, 0x3e, 0xe3, 0x3d, 0xea,
	0x79, 0x64, 0xe4, 0x0e, 0xc9, 0x28, 0x34, 0x73, 0x7a, 0x94, 0xf7, 0xe2, 0xc8, 0x9a, 0x4b, 0x78,
	0xe7, 0x71, 0x1b, 0xdd, 0x48, 0x0b, 0x3b, 0x64, 0x14, 0xc2, 0x55, 0x90, 0xa7, 0xde, 0x38, 0x16,
	0xf5, 0x38, 0xb2, 0x4a, 0x09, 0x8b, 0x7a, 0xca, 0xc4, 0x6d, 0x30, 0xbb, 0xbf, 0xbb, 0xec, 0xac,
	0x63, 0xa7, 0xbf, 0xe1, 0x6c, 0xee, 0x9d, 0xb4, 0x57, 0x4f, 0xeb, 0x28, 0x4f, 0x3d, 0xb8, 0x08,
	0x0a, 0x8c, 0xfb, 0xe3, 0x11, 0xdc, 0x89, 0x23, 0x0b, 0x24, 0x34, 0xc6, 0xfd, 0x33, 0xf3, 0xaa,
	0x21, 0x8d, 0x9d, 0x31, 0x2d, 0x76, 0x67, 0x39, 0x4f, 0x52, 0xf0, 0x7f, 0x39, 0xdf, 0x06, 0x20,
	0x7b, 0xdf, 0xe6, 0x8c, 0x56, 0x5c, 0x8a, 0x23, 0x6b, 0x76, 0x7c, 0x45, 0x29, 0x76, 0xa1, 0xbd,
	0x09, 0x3e, 0x5c, 0x05, 0x40, 0xa9, 0xba, 0x3e, 0x67, 0x61, 0x60, 0x5e, 0xd5, 0x6a, 0x73, 0x99,
	0x5a, 0x86, 0xd9, 0xa8, 0xa4, 0x36, 0x5b, 0x7a, 0xfd, 0x2b, 0x07, 0xea, 0x53, 0xef, 0xf6, 0x3d,
	0xc7, 0x41, 0x90, 0x3c, 0xca, 0xbf, 0xc7, 0xf3, 0x29, 0x28, 0xab, 0xef, 0xcb, 0x2f, 0xc1, 0x00,
	0xd3, 0xd1, 0xf8, 0xa6, 0xee, 0xc6, 0x91, 0x05, 0xb3, 0xde, 0x31, 0x68, 0xa3, 0xc9, 0x56, 0xb8,
	0x08, 0x8a, 0x84, 0x73, 0xc6, 0xf5, 0x35, 0x95, 0x3a, 0x95, 0x38, 0xb2, 0xae, 0x27, 0x1c, 0x5d,
	0xb6, 0x51, 0x02, 0xc3, 0x57, 0x69, 0x38, 0x73, 0x8d, 0x72, 0x7b, 0xe1, 0x1f, 0xc2, 0x79, 0x41,
	0x2e, 0x7b, 0x33, 0xfa, 0x9f, 0x6d, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xb0, 0x0d,
	0xb2, 0x68, 0x05, 0x00, 0x00,
}
