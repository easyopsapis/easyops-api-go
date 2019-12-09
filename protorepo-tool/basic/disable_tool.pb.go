// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: disable_tool.proto

package basic

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
//DisableTool请求
type DisableToolRequest struct {
	//
	//禁用工具开关
	Disable string `protobuf:"bytes,1,opt,name=disable,proto3" json:"disable" form:"disable"`
	//
	//工具Id，服务端自动生成
	ToolId               string   `protobuf:"bytes,2,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisableToolRequest) Reset()         { *m = DisableToolRequest{} }
func (m *DisableToolRequest) String() string { return proto.CompactTextString(m) }
func (*DisableToolRequest) ProtoMessage()    {}
func (*DisableToolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0085f802157aaeb0, []int{0}
}
func (m *DisableToolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisableToolRequest.Unmarshal(m, b)
}
func (m *DisableToolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisableToolRequest.Marshal(b, m, deterministic)
}
func (m *DisableToolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableToolRequest.Merge(m, src)
}
func (m *DisableToolRequest) XXX_Size() int {
	return xxx_messageInfo_DisableToolRequest.Size(m)
}
func (m *DisableToolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableToolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisableToolRequest proto.InternalMessageInfo

func (m *DisableToolRequest) GetDisable() string {
	if m != nil {
		return m.Disable
	}
	return ""
}

func (m *DisableToolRequest) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

//
//DisableTool返回
type DisableToolResponse struct {
	//
	//工具Id，服务端自动生成
	ToolId               string   `protobuf:"bytes,1,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisableToolResponse) Reset()         { *m = DisableToolResponse{} }
func (m *DisableToolResponse) String() string { return proto.CompactTextString(m) }
func (*DisableToolResponse) ProtoMessage()    {}
func (*DisableToolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0085f802157aaeb0, []int{1}
}
func (m *DisableToolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisableToolResponse.Unmarshal(m, b)
}
func (m *DisableToolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisableToolResponse.Marshal(b, m, deterministic)
}
func (m *DisableToolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableToolResponse.Merge(m, src)
}
func (m *DisableToolResponse) XXX_Size() int {
	return xxx_messageInfo_DisableToolResponse.Size(m)
}
func (m *DisableToolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableToolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisableToolResponse proto.InternalMessageInfo

func (m *DisableToolResponse) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

//
//DisableToolApi返回
type DisableToolResponseWrapper struct {
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
	Data                 *DisableToolResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DisableToolResponseWrapper) Reset()         { *m = DisableToolResponseWrapper{} }
func (m *DisableToolResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DisableToolResponseWrapper) ProtoMessage()    {}
func (*DisableToolResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_0085f802157aaeb0, []int{2}
}
func (m *DisableToolResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisableToolResponseWrapper.Unmarshal(m, b)
}
func (m *DisableToolResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisableToolResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DisableToolResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableToolResponseWrapper.Merge(m, src)
}
func (m *DisableToolResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DisableToolResponseWrapper.Size(m)
}
func (m *DisableToolResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableToolResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DisableToolResponseWrapper proto.InternalMessageInfo

func (m *DisableToolResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DisableToolResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DisableToolResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DisableToolResponseWrapper) GetData() *DisableToolResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DisableToolRequest)(nil), "basic.DisableToolRequest")
	proto.RegisterType((*DisableToolResponse)(nil), "basic.DisableToolResponse")
	proto.RegisterType((*DisableToolResponseWrapper)(nil), "basic.DisableToolResponseWrapper")
}

func init() { proto.RegisterFile("disable_tool.proto", fileDescriptor_0085f802157aaeb0) }

var fileDescriptor_0085f802157aaeb0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xc9, 0xf7, 0xb5, 0x15, 0xa7, 0xfe, 0x9d, 0x82, 0x84, 0x6c, 0x52, 0x46, 0x91, 0x2e,
	0x4c, 0x22, 0x2d, 0x88, 0xba, 0x51, 0x8b, 0x0a, 0x6e, 0x07, 0xa1, 0x0b, 0x51, 0x99, 0x34, 0x69,
	0x0c, 0xa6, 0xbd, 0x71, 0x66, 0x6a, 0x05, 0x71, 0xe7, 0x73, 0x46, 0x10, 0x9f, 0x20, 0x4f, 0x20,
	0x99, 0x89, 0x1a, 0xa1, 0x2b, 0x57, 0x99, 0xb9, 0xe7, 0xfc, 0x0e, 0xe7, 0x92, 0x41, 0x38, 0x88,
	0x05, 0xf3, 0x93, 0xf0, 0x56, 0x02, 0x24, 0x6e, 0xca, 0x41, 0x02, 0xae, 0xfb, 0x4c, 0xc4, 0x43,
	0xcb, 0x89, 0x62, 0x79, 0x37, 0xf5, 0xdd, 0x21, 0x8c, 0xbd, 0x08, 0x22, 0xf0, 0x94, 0xea, 0x4f,
	0x47, 0xea, 0xa6, 0x2e, 0xea, 0xa4, 0x29, 0x6b, 0xaf, 0x62, 0x1f, 0xcf, 0x62, 0x79, 0x0f, 0x33,
	0x2f, 0x02, 0x47, 0x89, 0xce, 0x23, 0x4b, 0xe2, 0x80, 0x49, 0xe0, 0xc2, 0xfb, 0x3e, 0x6a, 0x8e,
	0xbc, 0x1a, 0x08, 0x9f, 0xea, 0x12, 0x97, 0x00, 0x09, 0x0d, 0x1f, 0xa6, 0xa1, 0x90, 0x78, 0x07,
	0x2d, 0x94, 0xd5, 0x4c, 0xa3, 0x6d, 0x74, 0x16, 0xfb, 0x38, 0xcf, 0xec, 0x95, 0x11, 0xf0, 0xf1,
	0x21, 0x29, 0x05, 0x42, 0xbf, 0x2c, 0xf8, 0x18, 0x35, 0x8a, 0x05, 0x2e, 0x02, 0xf3, 0x9f, 0x32,
	0x77, 0xf2, 0xcc, 0x5e, 0xd6, 0x66, 0x3d, 0x27, 0xef, 0x6f, 0x76, 0x0b, 0xad, 0xdf, 0x5c, 0x31,
	0x67, 0x74, 0xe2, 0x9c, 0xef, 0x3a, 0x07, 0xd7, 0xcf, 0xbd, 0xee, 0xcb, 0x16, 0x2d, 0x39, 0x32,
	0x40, 0xad, 0x5f, 0x2d, 0x44, 0x0a, 0x13, 0x51, 0x0d, 0x36, 0xfe, 0x18, 0xfc, 0x61, 0x20, 0x6b,
	0x4e, 0xf2, 0x80, 0xb3, 0x34, 0x0d, 0x39, 0xde, 0x44, 0xb5, 0x21, 0x04, 0x7a, 0xc9, 0x7a, 0x7f,
	0x35, 0xcf, 0xec, 0xa6, 0x8e, 0x2f, 0xa6, 0x84, 0x2a, 0x11, 0xef, 0xa3, 0x66, 0xf1, 0x3d, 0x7b,
	0x4a, 0x13, 0x16, 0x4f, 0xca, 0x1d, 0x37, 0xf2, 0xcc, 0xc6, 0x3f, 0xde, 0x52, 0x24, 0xb4, 0x6a,
	0xc5, 0xdb, 0xa8, 0x1e, 0x72, 0x0e, 0xdc, 0xfc, 0xaf, 0x98, 0xb5, 0x3c, 0xb3, 0x97, 0x34, 0xa3,
	0xc6, 0x84, 0x6a, 0x19, 0x1f, 0xa1, 0x5a, 0xc0, 0x24, 0x33, 0x6b, 0x6d, 0xa3, 0xd3, 0xec, 0x5a,
	0xae, 0x7a, 0x02, 0xee, 0x9c, 0xde, 0xd5, 0x8a, 0x05, 0x41, 0xa8, 0x02, 0xfd, 0x86, 0xfa, 0x9b,
	0xbd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xa9, 0x47, 0x90, 0x51, 0x02, 0x00, 0x00,
}