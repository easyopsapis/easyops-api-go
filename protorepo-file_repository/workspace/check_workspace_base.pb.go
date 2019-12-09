// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: check_workspace_base.proto

package workspace

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
//CheckWorkspaceBase请求
type CheckWorkspaceBaseRequest struct {
	//
	//程序包Id
	PackageId string `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//工作区基础版本Id
	WorkspaceBaseId      string   `protobuf:"bytes,2,opt,name=workspaceBaseId,proto3" json:"workspaceBaseId" form:"workspaceBaseId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckWorkspaceBaseRequest) Reset()         { *m = CheckWorkspaceBaseRequest{} }
func (m *CheckWorkspaceBaseRequest) String() string { return proto.CompactTextString(m) }
func (*CheckWorkspaceBaseRequest) ProtoMessage()    {}
func (*CheckWorkspaceBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d34e0784273dc25, []int{0}
}
func (m *CheckWorkspaceBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckWorkspaceBaseRequest.Unmarshal(m, b)
}
func (m *CheckWorkspaceBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckWorkspaceBaseRequest.Marshal(b, m, deterministic)
}
func (m *CheckWorkspaceBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckWorkspaceBaseRequest.Merge(m, src)
}
func (m *CheckWorkspaceBaseRequest) XXX_Size() int {
	return xxx_messageInfo_CheckWorkspaceBaseRequest.Size(m)
}
func (m *CheckWorkspaceBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckWorkspaceBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckWorkspaceBaseRequest proto.InternalMessageInfo

func (m *CheckWorkspaceBaseRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *CheckWorkspaceBaseRequest) GetWorkspaceBaseId() string {
	if m != nil {
		return m.WorkspaceBaseId
	}
	return ""
}

//
//CheckWorkspaceBase返回
type CheckWorkspaceBaseResponse struct {
	//
	//成功时，返回工作区关联版本ID
	WorkspaceBaseId      string   `protobuf:"bytes,1,opt,name=workspaceBaseId,proto3" json:"workspaceBaseId" form:"workspaceBaseId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckWorkspaceBaseResponse) Reset()         { *m = CheckWorkspaceBaseResponse{} }
func (m *CheckWorkspaceBaseResponse) String() string { return proto.CompactTextString(m) }
func (*CheckWorkspaceBaseResponse) ProtoMessage()    {}
func (*CheckWorkspaceBaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d34e0784273dc25, []int{1}
}
func (m *CheckWorkspaceBaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckWorkspaceBaseResponse.Unmarshal(m, b)
}
func (m *CheckWorkspaceBaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckWorkspaceBaseResponse.Marshal(b, m, deterministic)
}
func (m *CheckWorkspaceBaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckWorkspaceBaseResponse.Merge(m, src)
}
func (m *CheckWorkspaceBaseResponse) XXX_Size() int {
	return xxx_messageInfo_CheckWorkspaceBaseResponse.Size(m)
}
func (m *CheckWorkspaceBaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckWorkspaceBaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckWorkspaceBaseResponse proto.InternalMessageInfo

func (m *CheckWorkspaceBaseResponse) GetWorkspaceBaseId() string {
	if m != nil {
		return m.WorkspaceBaseId
	}
	return ""
}

//
//CheckWorkspaceBaseApi返回
type CheckWorkspaceBaseResponseWrapper struct {
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
	Data                 *CheckWorkspaceBaseResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CheckWorkspaceBaseResponseWrapper) Reset()         { *m = CheckWorkspaceBaseResponseWrapper{} }
func (m *CheckWorkspaceBaseResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CheckWorkspaceBaseResponseWrapper) ProtoMessage()    {}
func (*CheckWorkspaceBaseResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d34e0784273dc25, []int{2}
}
func (m *CheckWorkspaceBaseResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckWorkspaceBaseResponseWrapper.Unmarshal(m, b)
}
func (m *CheckWorkspaceBaseResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckWorkspaceBaseResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CheckWorkspaceBaseResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckWorkspaceBaseResponseWrapper.Merge(m, src)
}
func (m *CheckWorkspaceBaseResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CheckWorkspaceBaseResponseWrapper.Size(m)
}
func (m *CheckWorkspaceBaseResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckWorkspaceBaseResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CheckWorkspaceBaseResponseWrapper proto.InternalMessageInfo

func (m *CheckWorkspaceBaseResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CheckWorkspaceBaseResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CheckWorkspaceBaseResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CheckWorkspaceBaseResponseWrapper) GetData() *CheckWorkspaceBaseResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckWorkspaceBaseRequest)(nil), "workspace.CheckWorkspaceBaseRequest")
	proto.RegisterType((*CheckWorkspaceBaseResponse)(nil), "workspace.CheckWorkspaceBaseResponse")
	proto.RegisterType((*CheckWorkspaceBaseResponseWrapper)(nil), "workspace.CheckWorkspaceBaseResponseWrapper")
}

func init() { proto.RegisterFile("check_workspace_base.proto", fileDescriptor_2d34e0784273dc25) }

var fileDescriptor_2d34e0784273dc25 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0x48, 0x4d,
	0xce, 0x8e, 0x2f, 0xcf, 0x2f, 0xca, 0x2e, 0x2e, 0x48, 0x4c, 0x4e, 0x8d, 0x4f, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x8b, 0x4a, 0xe9, 0xa6, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x55, 0x24, 0x95,
	0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x29, 0x65, 0x86, 0xa4, 0x3c, 0xb7, 0x3c, 0xb3,
	0x24, 0x3b, 0xbf, 0x5c, 0x3f, 0x3d, 0x5f, 0x17, 0x2c, 0xa9, 0x5b, 0x96, 0x98, 0x93, 0x99, 0x92,
	0x58, 0x92, 0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0xf4, 0x29, 0x1d, 0x65, 0xe2, 0x92, 0x74, 0x06,
	0x39, 0x28, 0x1c, 0x66, 0xb3, 0x53, 0x62, 0x71, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x50, 0x27, 0x23, 0x17, 0x67, 0x41, 0x62, 0x72, 0x76, 0x62, 0x7a, 0xaa, 0x67, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0xa7, 0x53, 0xf6, 0xa7, 0x7b, 0xf2, 0x02, 0x69, 0xf9, 0x45, 0xb9, 0x56, 0x4a,
	0x70, 0x29, 0xa5, 0x47, 0xf7, 0xe5, 0x83, 0xb9, 0x02, 0xe3, 0xa2, 0x13, 0x75, 0xd3, 0x1c, 0x75,
	0xdd, 0x0c, 0x74, 0x2d, 0x63, 0xab, 0x2d, 0x6a, 0x75, 0xed, 0x91, 0xf9, 0x26, 0x24, 0xf2, 0x0d,
	0x8d, 0x6a, 0x55, 0x82, 0x10, 0xb6, 0x0b, 0xcd, 0x66, 0xe4, 0xe2, 0x2f, 0x47, 0x76, 0xa4, 0x67,
	0x8a, 0x04, 0x13, 0xd8, 0x45, 0x85, 0x9f, 0xee, 0xc9, 0x8b, 0x41, 0x5c, 0x84, 0xa6, 0x80, 0x66,
	0xee, 0x42, 0x77, 0x89, 0xd2, 0x2e, 0x46, 0x2e, 0x29, 0x6c, 0xe1, 0x58, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x8a, 0xd5, 0xf1, 0x8c, 0x83, 0xc6, 0xf1, 0xdf, 0x18, 0xb9, 0x14, 0x71, 0x3b, 0x3e, 0xbc,
	0x28, 0xb1, 0xa0, 0x20, 0xb5, 0x48, 0x48, 0x99, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x15, 0xec, 0x6e,
	0x56, 0x27, 0xfe, 0x4f, 0xf7, 0xe4, 0xb9, 0x21, 0xee, 0x06, 0x89, 0x2a, 0x05, 0x81, 0x25, 0x85,
	0x2c, 0xb8, 0xb8, 0x41, 0xb4, 0x6b, 0x45, 0x41, 0x4e, 0x62, 0x66, 0x1e, 0x34, 0x82, 0xc4, 0x3e,
	0xdd, 0x93, 0x17, 0x42, 0xa8, 0x85, 0x4a, 0x2a, 0x05, 0x21, 0x2b, 0x15, 0x52, 0xe3, 0x62, 0x4d,
	0x2d, 0x2a, 0xca, 0x2f, 0x92, 0x60, 0x06, 0xeb, 0x11, 0xf8, 0x74, 0x4f, 0x9e, 0x07, 0xa2, 0x07,
	0x2c, 0xac, 0x14, 0x04, 0x91, 0x16, 0xf2, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x51,
	0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd5, 0x83, 0x7b, 0x46, 0x0f, 0xb7, 0x17, 0x90, 0x5d, 0x0b, 0xd2,
	0xac, 0x14, 0x04, 0x36, 0x23, 0x89, 0x0d, 0x9c, 0x09, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x16, 0xa0, 0x7e, 0x85, 0x94, 0x03, 0x00, 0x00,
}
