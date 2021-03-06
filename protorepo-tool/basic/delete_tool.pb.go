// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_tool.proto

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
//DeleteTool请求
type DeleteToolRequest struct {
	//
	//强制删除工具，当工具被流程引用时 "true"|"false"
	Force string `protobuf:"bytes,1,opt,name=force,proto3" json:"force" form:"force"`
	//
	//版本Id
	VersionId string `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//工具Id，服务端自动生成
	ToolId               string   `protobuf:"bytes,3,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteToolRequest) Reset()         { *m = DeleteToolRequest{} }
func (m *DeleteToolRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteToolRequest) ProtoMessage()    {}
func (*DeleteToolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ec6c172fcf1da1a, []int{0}
}
func (m *DeleteToolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteToolRequest.Unmarshal(m, b)
}
func (m *DeleteToolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteToolRequest.Marshal(b, m, deterministic)
}
func (m *DeleteToolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteToolRequest.Merge(m, src)
}
func (m *DeleteToolRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteToolRequest.Size(m)
}
func (m *DeleteToolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteToolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteToolRequest proto.InternalMessageInfo

func (m *DeleteToolRequest) GetForce() string {
	if m != nil {
		return m.Force
	}
	return ""
}

func (m *DeleteToolRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *DeleteToolRequest) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

//
//DeleteToolApi返回
type DeleteToolResponseWrapper struct {
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

func (m *DeleteToolResponseWrapper) Reset()         { *m = DeleteToolResponseWrapper{} }
func (m *DeleteToolResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteToolResponseWrapper) ProtoMessage()    {}
func (*DeleteToolResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ec6c172fcf1da1a, []int{1}
}
func (m *DeleteToolResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteToolResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteToolResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteToolResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteToolResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteToolResponseWrapper.Merge(m, src)
}
func (m *DeleteToolResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteToolResponseWrapper.Size(m)
}
func (m *DeleteToolResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteToolResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteToolResponseWrapper proto.InternalMessageInfo

func (m *DeleteToolResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteToolResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteToolResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteToolResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteToolRequest)(nil), "basic.DeleteToolRequest")
	proto.RegisterType((*DeleteToolResponseWrapper)(nil), "basic.DeleteToolResponseWrapper")
}

func init() { proto.RegisterFile("delete_tool.proto", fileDescriptor_0ec6c172fcf1da1a) }

var fileDescriptor_0ec6c172fcf1da1a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xdb, 0x4a, 0xe3, 0x40,
	0x18, 0x26, 0xbb, 0x6d, 0xa1, 0xd3, 0x5d, 0xb6, 0x9d, 0x85, 0x92, 0xed, 0x5e, 0xa4, 0xcc, 0x2e,
	0x4b, 0x61, 0x49, 0x22, 0x2d, 0x88, 0x87, 0x1b, 0x2d, 0x56, 0xe8, 0x6d, 0x10, 0xbc, 0x10, 0x95,
	0x49, 0x32, 0x89, 0xc1, 0xa4, 0x7f, 0x9c, 0x4c, 0x5b, 0x45, 0x7c, 0x3f, 0x9f, 0x22, 0x05, 0x1f,
	0x21, 0x4f, 0x20, 0x99, 0x89, 0x6d, 0x6e, 0xbc, 0x9a, 0xf9, 0xbf, 0x13, 0xff, 0x01, 0xf5, 0x7c,
	0x16, 0x33, 0xc1, 0x6e, 0x05, 0x40, 0x6c, 0xa5, 0x1c, 0x04, 0xe0, 0xa6, 0x4b, 0xb3, 0xc8, 0x1b,
	0x98, 0x61, 0x24, 0xee, 0x96, 0xae, 0xe5, 0x41, 0x62, 0x87, 0x10, 0x82, 0x2d, 0x59, 0x77, 0x19,
	0xc8, 0x4a, 0x16, 0xf2, 0xa7, 0x5c, 0x83, 0xfd, 0x9a, 0x3c, 0x59, 0x47, 0xe2, 0x1e, 0xd6, 0x76,
	0x08, 0xa6, 0x24, 0xcd, 0x15, 0x8d, 0x23, 0x9f, 0x0a, 0xe0, 0x99, 0xbd, 0xfd, 0x56, 0xbe, 0xdf,
	0x21, 0x40, 0x18, 0xb3, 0x5d, 0x3a, 0x4b, 0x52, 0xf1, 0xa4, 0x48, 0xf2, 0xaa, 0xa1, 0xde, 0x99,
	0x6c, 0xf0, 0x02, 0x20, 0x76, 0xd8, 0xc3, 0x92, 0x65, 0x02, 0xff, 0x43, 0xcd, 0x00, 0xb8, 0xc7,
	0x74, 0x6d, 0xa8, 0x8d, 0xda, 0xd3, 0x6e, 0x91, 0x1b, 0xdf, 0x02, 0xe0, 0xc9, 0x11, 0x91, 0x30,
	0x71, 0x14, 0x8d, 0xe7, 0xa8, 0xbd, 0x62, 0x3c, 0x8b, 0x60, 0x31, 0xf7, 0xf5, 0x2f, 0x52, 0xfb,
	0xbf, 0xc8, 0x8d, 0xae, 0xd2, 0x6e, 0x29, 0xf2, 0xb6, 0x31, 0x7e, 0xa2, 0xde, 0xcd, 0x15, 0x35,
	0x83, 0x53, 0xf3, 0x7c, 0xcf, 0x3c, 0xbc, 0x7e, 0x9e, 0x8c, 0x5f, 0xfe, 0x3a, 0x3b, 0x37, 0x3e,
	0x41, 0xad, 0x72, 0x43, 0x73, 0x5f, 0xff, 0x2a, 0x73, 0x46, 0x45, 0x6e, 0x7c, 0x57, 0x39, 0x0a,
	0xff, 0x34, 0xa4, 0xf2, 0x91, 0x8d, 0x86, 0x7e, 0xd5, 0x47, 0xc9, 0x52, 0x58, 0x64, 0xec, 0x92,
	0xd3, 0x34, 0x65, 0x1c, 0xff, 0x41, 0x0d, 0x0f, 0x7c, 0x35, 0x51, 0x73, 0xfa, 0xa3, 0xc8, 0x8d,
	0x8e, 0x4a, 0x2f, 0x51, 0xe2, 0x48, 0x12, 0x1f, 0xa0, 0x4e, 0xf9, 0xce, 0x1e, 0xd3, 0x98, 0x46,
	0x8b, 0x6a, 0xa2, 0x7e, 0x91, 0x1b, 0x78, 0xa7, 0xad, 0x48, 0xe2, 0xd4, 0xa5, 0xe5, 0xc6, 0x18,
	0xe7, 0xc0, 0xab, 0xee, 0x6b, 0x1b, 0x93, 0x30, 0x71, 0x14, 0x8d, 0x8f, 0x51, 0xc3, 0xa7, 0x82,
	0xea, 0x8d, 0xa1, 0x36, 0xea, 0x8c, 0xfb, 0x96, 0xba, 0x8d, 0xf5, 0x71, 0x1b, 0x6b, 0x56, 0xde,
	0xa6, 0xde, 0x5e, 0xa9, 0x26, 0x8e, 0x34, 0xb9, 0x2d, 0x29, 0x9b, 0xbc, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0xe8, 0x95, 0x09, 0x53, 0x02, 0x00, 0x00,
}
