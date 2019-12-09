// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_instance_apply_permission.proto

package cmdb_approve

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/notify"
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
//CreateInstanceApplyPermissionApi返回
type CreateInstanceApplyPermissionResponseWrapper struct {
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

func (m *CreateInstanceApplyPermissionResponseWrapper) Reset() {
	*m = CreateInstanceApplyPermissionResponseWrapper{}
}
func (m *CreateInstanceApplyPermissionResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*CreateInstanceApplyPermissionResponseWrapper) ProtoMessage() {}
func (*CreateInstanceApplyPermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa968746a1e2a95, []int{0}
}
func (m *CreateInstanceApplyPermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper.Unmarshal(m, b)
}
func (m *CreateInstanceApplyPermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateInstanceApplyPermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper.Merge(m, src)
}
func (m *CreateInstanceApplyPermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper.Size(m)
}
func (m *CreateInstanceApplyPermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceApplyPermissionResponseWrapper proto.InternalMessageInfo

func (m *CreateInstanceApplyPermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateInstanceApplyPermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateInstanceApplyPermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateInstanceApplyPermissionResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateInstanceApplyPermissionResponseWrapper)(nil), "cmdb_approve.CreateInstanceApplyPermissionResponseWrapper")
}

func init() {
	proto.RegisterFile("create_instance_apply_permission.proto", fileDescriptor_dfa968746a1e2a95)
}

var fileDescriptor_dfa968746a1e2a95 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x6a, 0xeb, 0x30,
	0x14, 0x86, 0xf1, 0xbd, 0xc9, 0x85, 0xeb, 0x04, 0x5a, 0x3c, 0x04, 0x93, 0x0e, 0x09, 0x2e, 0x84,
	0x0c, 0x8d, 0x04, 0xed, 0x52, 0xda, 0xa9, 0x29, 0x19, 0xba, 0x15, 0x53, 0xe8, 0x68, 0x64, 0xf9,
	0x44, 0x15, 0xc8, 0x3a, 0x42, 0x52, 0x4a, 0xfd, 0xb2, 0x79, 0x08, 0x3f, 0x41, 0xb1, 0xec, 0x36,
	0x99, 0xec, 0x73, 0xfe, 0xef, 0x97, 0xf4, 0xc5, 0x2b, 0x6e, 0x81, 0x79, 0x28, 0xa4, 0x76, 0x9e,
	0x69, 0x0e, 0x05, 0x33, 0x46, 0x35, 0x85, 0x01, 0x5b, 0x4b, 0xe7, 0x24, 0x6a, 0x62, 0x2c, 0x7a,
	0x4c, 0xa6, 0xbc, 0xae, 0xca, 0x2e, 0xb4, 0xf8, 0x09, 0xf3, 0x8d, 0x90, 0xfe, 0xe3, 0x50, 0x12,
	0x8e, 0x35, 0x15, 0x28, 0x90, 0x06, 0xa8, 0x3c, 0xec, 0xc3, 0x14, 0x86, 0xf0, 0xd7, 0x97, 0xe7,
	0x6f, 0x02, 0x09, 0x30, 0xd7, 0xa0, 0x71, 0x44, 0x21, 0x67, 0x8a, 0x72, 0xd4, 0xde, 0x32, 0xee,
	0x5d, 0xdf, 0xb4, 0x60, 0x70, 0x53, 0x63, 0x05, 0xca, 0xd1, 0x01, 0xa4, 0x61, 0xa4, 0x1a, 0xbd,
	0xdc, 0x37, 0x14, 0x0d, 0x58, 0xe6, 0x25, 0xea, 0x42, 0xa1, 0x18, 0x4e, 0xbd, 0x12, 0x88, 0x42,
	0xc1, 0xe9, 0x6e, 0xa8, 0x8d, 0x6f, 0xfa, 0x30, 0x6b, 0xa3, 0xf8, 0xe6, 0x39, 0xa8, 0xbd, 0x0c,
	0x66, 0x4f, 0x9d, 0xd8, 0xeb, 0xaf, 0x57, 0x0e, 0xce, 0xa0, 0x76, 0xf0, 0x6e, 0x99, 0x31, 0x60,
	0x93, 0xeb, 0x78, 0xc4, 0xb1, 0x82, 0x34, 0x5a, 0x46, 0xeb, 0xf1, 0xf6, 0xa2, 0x3d, 0x2e, 0x26,
	0x7b, 0xb4, 0xf5, 0x43, 0xd6, 0x6d, 0xb3, 0x3c, 0x84, 0xc9, 0x7d, 0x3c, 0xe9, 0xbe, 0xbb, 0x2f,
	0xa3, 0x98, 0xd4, 0xe9, 0x9f, 0x65, 0xb4, 0xfe, 0xbf, 0x9d, 0xb5, 0xc7, 0x45, 0x72, 0x62, 0x87,
	0x30, 0xcb, 0xcf, 0xd1, 0x64, 0x15, 0x8f, 0xc1, 0x5a, 0xb4, 0xe9, 0xdf, 0xd0, 0xb9, 0x6c, 0x8f,
	0x8b, 0x69, 0xdf, 0x09, 0xeb, 0x2c, 0xef, 0xe3, 0xe4, 0x31, 0x1e, 0x55, 0xcc, 0xb3, 0x74, 0xb4,
	0x8c, 0xd6, 0x93, 0xdb, 0x19, 0xe9, 0x1d, 0xc9, 0x8f, 0x23, 0xd9, 0x75, 0x8e, 0xe7, 0xcf, 0xeb,
	0xe8, 0x2c, 0x0f, 0xa5, 0xf2, 0x5f, 0xc0, 0xee, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0xf5,
	0x0c, 0x92, 0xd5, 0x01, 0x00, 0x00,
}
