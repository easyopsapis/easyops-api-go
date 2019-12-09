// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_toolkit_status.proto

package toolkit

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
//UpdateToolkitStatus请求
type UpdateToolkitStatusRequest struct {
	//
	//模型Id
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//状态
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled" form:"enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateToolkitStatusRequest) Reset()         { *m = UpdateToolkitStatusRequest{} }
func (m *UpdateToolkitStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateToolkitStatusRequest) ProtoMessage()    {}
func (*UpdateToolkitStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71aaf5e49c8b6018, []int{0}
}
func (m *UpdateToolkitStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateToolkitStatusRequest.Unmarshal(m, b)
}
func (m *UpdateToolkitStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateToolkitStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateToolkitStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateToolkitStatusRequest.Merge(m, src)
}
func (m *UpdateToolkitStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateToolkitStatusRequest.Size(m)
}
func (m *UpdateToolkitStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateToolkitStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateToolkitStatusRequest proto.InternalMessageInfo

func (m *UpdateToolkitStatusRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UpdateToolkitStatusRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

//
//UpdateToolkitStatusApi返回
type UpdateToolkitStatusResponseWrapper struct {
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

func (m *UpdateToolkitStatusResponseWrapper) Reset()         { *m = UpdateToolkitStatusResponseWrapper{} }
func (m *UpdateToolkitStatusResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateToolkitStatusResponseWrapper) ProtoMessage()    {}
func (*UpdateToolkitStatusResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_71aaf5e49c8b6018, []int{1}
}
func (m *UpdateToolkitStatusResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateToolkitStatusResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateToolkitStatusResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateToolkitStatusResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateToolkitStatusResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateToolkitStatusResponseWrapper.Merge(m, src)
}
func (m *UpdateToolkitStatusResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateToolkitStatusResponseWrapper.Size(m)
}
func (m *UpdateToolkitStatusResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateToolkitStatusResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateToolkitStatusResponseWrapper proto.InternalMessageInfo

func (m *UpdateToolkitStatusResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateToolkitStatusResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateToolkitStatusResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateToolkitStatusResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateToolkitStatusRequest)(nil), "toolkit.UpdateToolkitStatusRequest")
	proto.RegisterType((*UpdateToolkitStatusResponseWrapper)(nil), "toolkit.UpdateToolkitStatusResponseWrapper")
}

func init() { proto.RegisterFile("update_toolkit_status.proto", fileDescriptor_71aaf5e49c8b6018) }

var fileDescriptor_71aaf5e49c8b6018 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xb6, 0xb6, 0xdd, 0x8a, 0x95, 0x15, 0x4a, 0x68, 0x0f, 0x29, 0x2b, 0x48, 0x0f,
	0x9a, 0x82, 0x5e, 0x44, 0x6f, 0x81, 0x1e, 0xbc, 0xae, 0x8a, 0xc7, 0xb2, 0x69, 0xb6, 0xb1, 0x9a,
	0x64, 0xd6, 0xcd, 0x06, 0x14, 0xdf, 0x35, 0x67, 0xcf, 0x79, 0x02, 0xc9, 0x6c, 0x62, 0x7b, 0xf0,
	0x94, 0xec, 0x7c, 0xdf, 0xcc, 0x3f, 0x0c, 0x99, 0x16, 0x2a, 0x12, 0x46, 0xae, 0x0c, 0x40, 0xf2,
	0xbe, 0x35, 0xab, 0xdc, 0x08, 0x53, 0xe4, 0xbe, 0xd2, 0x60, 0x80, 0xf6, 0x9a, 0xea, 0xe4, 0x2a,
	0xde, 0x9a, 0xd7, 0x22, 0xf4, 0xd7, 0x90, 0x2e, 0x62, 0x88, 0x61, 0x81, 0x3c, 0x2c, 0x36, 0xf8,
	0xc2, 0x07, 0xfe, 0xd9, 0xbe, 0xc9, 0x34, 0x06, 0x88, 0x13, 0xb9, 0xb3, 0x64, 0xaa, 0xcc, 0x97,
	0x85, 0xec, 0x9b, 0x4c, 0x9e, 0x31, 0xf3, 0xc9, 0x0e, 0x7f, 0xc4, 0x44, 0x2e, 0x3f, 0x0a, 0x99,
	0x1b, 0xba, 0x20, 0x7d, 0x08, 0xdf, 0xe4, 0xda, 0x3c, 0x44, 0xae, 0x33, 0x73, 0xe6, 0x83, 0xe0,
	0xac, 0x2a, 0xbd, 0xd1, 0x06, 0x74, 0x7a, 0xc7, 0x5a, 0xc2, 0xf8, 0x9f, 0x44, 0x2f, 0x49, 0x4f,
	0x66, 0x22, 0x4c, 0x64, 0xe4, 0x1e, 0xcc, 0x9c, 0x79, 0x3f, 0xa0, 0x55, 0xe9, 0x9d, 0x58, 0xbf,
	0x01, 0x8c, 0xb7, 0x0a, 0xfb, 0x71, 0x08, 0xfb, 0x37, 0x3d, 0x57, 0x90, 0xe5, 0xf2, 0x45, 0x0b,
	0xa5, 0xa4, 0xa6, 0xe7, 0xa4, 0xb3, 0x86, 0x48, 0xe2, 0x06, 0xdd, 0x60, 0x54, 0x95, 0xde, 0xd0,
	0x4e, 0xac, 0xab, 0x8c, 0x23, 0xa4, 0xb7, 0x64, 0x58, 0x7f, 0x97, 0x9f, 0x2a, 0x11, 0xdb, 0x0c,
	0xd3, 0x07, 0xc1, 0xb8, 0x2a, 0x3d, 0xba, 0x73, 0x1b, 0xc8, 0xf8, 0xbe, 0x4a, 0x2f, 0x48, 0x57,
	0x6a, 0x0d, 0xda, 0x3d, 0xc4, 0x9e, 0xd3, 0xaa, 0xf4, 0x8e, 0x9b, 0x8d, 0xeb, 0x32, 0xe3, 0x16,
	0xd3, 0x7b, 0xd2, 0x89, 0x84, 0x11, 0x6e, 0x67, 0xe6, 0xcc, 0x87, 0xd7, 0x63, 0xdf, 0x9e, 0xd5,
	0x6f, 0xcf, 0xea, 0x2f, 0xeb, 0xb3, 0xee, 0xaf, 0x57, 0xdb, 0x8c, 0x63, 0x53, 0x78, 0x84, 0xda,
	0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x80, 0xfa, 0x9e, 0xe2, 0x01, 0x00, 0x00,
}
