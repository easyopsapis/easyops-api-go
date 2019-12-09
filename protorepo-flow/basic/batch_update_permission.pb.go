// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_update_permission.proto

package basic

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
//UpdateBatchPermission请求
type UpdateBatchPermissionRequest struct {
	//
	//批量更新方式, 可选值: append(追加), overwrite(重置), pull(移除)
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method" form:"method"`
	//
	//权限列表,可选值deleteAuthorizers(删除), updateAuthorizers(更新), readAuthorizers(读取或查看),executeAuthorizers(执行)
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions" form:"permissions"`
	//
	//流程ID
	FlowIds []string `protobuf:"bytes,3,rep,name=flowIds,proto3" json:"flowIds" form:"flowIds"`
	//
	//用户、用户组列表
	Authorizers          []string `protobuf:"bytes,4,rep,name=authorizers,proto3" json:"authorizers" form:"authorizers"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBatchPermissionRequest) Reset()         { *m = UpdateBatchPermissionRequest{} }
func (m *UpdateBatchPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchPermissionRequest) ProtoMessage()    {}
func (*UpdateBatchPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6dcef781171176, []int{0}
}
func (m *UpdateBatchPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchPermissionRequest.Unmarshal(m, b)
}
func (m *UpdateBatchPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchPermissionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBatchPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchPermissionRequest.Merge(m, src)
}
func (m *UpdateBatchPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchPermissionRequest.Size(m)
}
func (m *UpdateBatchPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchPermissionRequest proto.InternalMessageInfo

func (m *UpdateBatchPermissionRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UpdateBatchPermissionRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *UpdateBatchPermissionRequest) GetFlowIds() []string {
	if m != nil {
		return m.FlowIds
	}
	return nil
}

func (m *UpdateBatchPermissionRequest) GetAuthorizers() []string {
	if m != nil {
		return m.Authorizers
	}
	return nil
}

//
//UpdateBatchPermissionApi返回
type UpdateBatchPermissionResponseWrapper struct {
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

func (m *UpdateBatchPermissionResponseWrapper) Reset()         { *m = UpdateBatchPermissionResponseWrapper{} }
func (m *UpdateBatchPermissionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchPermissionResponseWrapper) ProtoMessage()    {}
func (*UpdateBatchPermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef6dcef781171176, []int{1}
}
func (m *UpdateBatchPermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchPermissionResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateBatchPermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchPermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateBatchPermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchPermissionResponseWrapper.Merge(m, src)
}
func (m *UpdateBatchPermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchPermissionResponseWrapper.Size(m)
}
func (m *UpdateBatchPermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchPermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchPermissionResponseWrapper proto.InternalMessageInfo

func (m *UpdateBatchPermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateBatchPermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateBatchPermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateBatchPermissionResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateBatchPermissionRequest)(nil), "basic.UpdateBatchPermissionRequest")
	proto.RegisterType((*UpdateBatchPermissionResponseWrapper)(nil), "basic.UpdateBatchPermissionResponseWrapper")
}

func init() { proto.RegisterFile("batch_update_permission.proto", fileDescriptor_ef6dcef781171176) }

var fileDescriptor_ef6dcef781171176 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x30,
	0x18, 0x26, 0x5f, 0x1d, 0x55, 0xf6, 0x55, 0x0f, 0x36, 0x53, 0x3a, 0x1c, 0xb4, 0x30, 0x1c, 0x98,
	0x9c, 0x36, 0x19, 0x65, 0x2d, 0xbb, 0x34, 0x63, 0x83, 0xdd, 0x8a, 0xc7, 0x18, 0xac, 0xa4, 0x41,
	0xb6, 0x15, 0xc7, 0x2c, 0xce, 0xeb, 0x49, 0xf2, 0x32, 0x92, 0xe5, 0x9f, 0xed, 0xb7, 0x78, 0xb0,
	0xeb, 0x6e, 0xfe, 0x05, 0xc3, 0x52, 0xda, 0xe8, 0xd2, 0x93, 0xa5, 0xf7, 0xf9, 0xd0, 0xe3, 0x87,
	0x17, 0x3d, 0x0f, 0xa8, 0x0c, 0x67, 0x93, 0x3c, 0x8b, 0xa8, 0x64, 0x93, 0x8c, 0xf1, 0x34, 0x11,
	0x22, 0x81, 0x85, 0x97, 0x71, 0x90, 0x60, 0xb5, 0x02, 0x2a, 0x92, 0xf0, 0x90, 0xc4, 0x89, 0x9c,
	0xe5, 0x81, 0x17, 0x42, 0xda, 0x8f, 0x21, 0x86, 0xbe, 0x42, 0x83, 0x7c, 0xaa, 0x6e, 0xea, 0xa2,
	0x4e, 0x5a, 0x75, 0x78, 0x6a, 0xd0, 0xd3, 0x65, 0x22, 0xbf, 0xc1, 0xb2, 0x1f, 0x03, 0x51, 0x20,
	0xf9, 0x41, 0xe7, 0x49, 0x44, 0x25, 0x70, 0xd1, 0xbf, 0x3d, 0x6e, 0x75, 0x47, 0x31, 0x40, 0x3c,
	0x67, 0x3b, 0x77, 0x21, 0x79, 0x1e, 0x4a, 0x8d, 0xe2, 0xdf, 0x75, 0x74, 0xf4, 0x59, 0xe5, 0x1c,
	0x55, 0x99, 0x2f, 0x6f, 0xb3, 0xfa, 0xec, 0x7b, 0xce, 0x84, 0xb4, 0x7a, 0x68, 0x2f, 0x65, 0x72,
	0x06, 0x91, 0x5d, 0xeb, 0xd4, 0xdc, 0xfd, 0xd1, 0x41, 0x59, 0x38, 0x0f, 0xa6, 0xc0, 0xd3, 0x73,
	0xac, 0xe7, 0xd8, 0xdf, 0x12, 0xac, 0x37, 0xa8, 0xbd, 0xfb, 0x57, 0x61, 0xd7, 0x3b, 0x0d, 0x77,
	0x7f, 0xf4, 0xb4, 0x2c, 0x1c, 0x4b, 0xf3, 0x0d, 0x10, 0xfb, 0x26, 0xd5, 0x7a, 0x87, 0xee, 0x4d,
	0xe7, 0xb0, 0xfc, 0x18, 0x09, 0xbb, 0xa1, 0x54, 0xbd, 0xb2, 0x70, 0x1e, 0x6a, 0xd5, 0x16, 0xc0,
	0x7f, 0xff, 0x38, 0x4f, 0xd0, 0xc1, 0xf5, 0x15, 0x25, 0xd3, 0x0b, 0xf2, 0xe1, 0x98, 0x9c, 0x8d,
	0xd7, 0xc3, 0xc1, 0xa6, 0xeb, 0xdf, 0x28, 0x2d, 0x8e, 0xda, 0x34, 0x97, 0x33, 0xe0, 0xc9, 0x8a,
	0x71, 0x61, 0x37, 0x95, 0xd1, 0xe5, 0xee, 0x79, 0x03, 0xac, 0xcc, 0x4e, 0xd1, 0x6b, 0xb7, 0x72,
	0x5b, 0x5d, 0x90, 0xaf, 0x95, 0xdb, 0xee, 0x38, 0x21, 0xe3, 0xf5, 0xe0, 0xd5, 0xf0, 0x64, 0xd3,
	0xeb, 0xfe, 0x72, 0xaf, 0xcf, 0xaf, 0x8e, 0xc9, 0x19, 0x25, 0xab, 0xf1, 0xfa, 0x64, 0xb8, 0xe9,
	0x75, 0x7d, 0xf3, 0x11, 0xfc, 0xaf, 0x86, 0xba, 0x77, 0xd4, 0x27, 0x32, 0x58, 0x08, 0xf6, 0x85,
	0xd3, 0x2c, 0x63, 0xdc, 0x7a, 0x81, 0x9a, 0x21, 0x44, 0x4c, 0x95, 0xd8, 0x1a, 0x3d, 0x2a, 0x0b,
	0xa7, 0xad, 0x53, 0x55, 0x53, 0xec, 0x2b, 0xb0, 0x2a, 0xb0, 0xfa, 0xbe, 0xff, 0x99, 0xcd, 0x69,
	0xb2, 0xb0, 0xeb, 0xaa, 0x70, 0xa3, 0x40, 0x03, 0xc4, 0xbe, 0x49, 0xb5, 0x5e, 0xa2, 0x16, 0xe3,
	0x1c, 0xb8, 0xdd, 0x50, 0x9a, 0xc7, 0x65, 0xe1, 0xdc, 0xd7, 0x1a, 0x35, 0xc6, 0xbe, 0x86, 0xad,
	0xb7, 0xa8, 0x19, 0x51, 0x49, 0xed, 0x66, 0xa7, 0xe6, 0xb6, 0x07, 0xcf, 0x3c, 0xbd, 0x1b, 0xde,
	0xcd, 0x6e, 0x78, 0x9f, 0xd4, 0x6e, 0x98, 0xf9, 0x2a, 0x3a, 0xf6, 0x95, 0x2a, 0xd8, 0x53, 0xbc,
	0xe1, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xc6, 0xdc, 0xbd, 0xe0, 0x02, 0x00, 0x00,
}