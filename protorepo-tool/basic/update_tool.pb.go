// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_tool.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	tool "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
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
//UpdateToolApi返回
type UpdateToolResponseWrapper struct {
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
	Data                 *tool.Tool `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateToolResponseWrapper) Reset()         { *m = UpdateToolResponseWrapper{} }
func (m *UpdateToolResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateToolResponseWrapper) ProtoMessage()    {}
func (*UpdateToolResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d4588e1eff3bdf4, []int{0}
}
func (m *UpdateToolResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateToolResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateToolResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateToolResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateToolResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateToolResponseWrapper.Merge(m, src)
}
func (m *UpdateToolResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateToolResponseWrapper.Size(m)
}
func (m *UpdateToolResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateToolResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateToolResponseWrapper proto.InternalMessageInfo

func (m *UpdateToolResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateToolResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateToolResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateToolResponseWrapper) GetData() *tool.Tool {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateToolResponseWrapper)(nil), "basic.UpdateToolResponseWrapper")
}

func init() { proto.RegisterFile("update_tool.proto", fileDescriptor_2d4588e1eff3bdf4) }

var fileDescriptor_2d4588e1eff3bdf4 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0xdf, 0xd7, 0x0a, 0x6e, 0x05, 0x35, 0x07, 0x89, 0x45, 0x68, 0x89, 0x20, 0xbd, 0x24,
	0x0b, 0x0a, 0x22, 0x1e, 0x0b, 0x1e, 0x3c, 0x09, 0x41, 0xf1, 0x58, 0x36, 0x9b, 0xed, 0xba, 0xb8,
	0xc9, 0x2c, 0xbb, 0x1b, 0xab, 0xbf, 0xd3, 0x7b, 0x7f, 0x44, 0x7f, 0x81, 0x64, 0xb2, 0xd8, 0xfc,
	0x00, 0x2f, 0xc9, 0xcc, 0xbc, 0xf7, 0x66, 0xdf, 0x1b, 0x72, 0xda, 0x9a, 0x8a, 0x79, 0xb1, 0xf2,
	0x00, 0x3a, 0x37, 0x16, 0x3c, 0xc4, 0xe3, 0x92, 0x39, 0xc5, 0xa7, 0x99, 0x54, 0xfe, 0xad, 0x2d,
	0x73, 0x0e, 0x35, 0x95, 0x20, 0x81, 0x22, 0x5a, 0xb6, 0x6b, 0xec, 0xb0, 0xc1, 0xaa, 0x57, 0x4d,
	0x9f, 0x24, 0xe4, 0x82, 0xb9, 0x2f, 0x30, 0x2e, 0xd7, 0xc0, 0x99, 0xa6, 0x1c, 0x1a, 0x6f, 0x19,
	0xf7, 0xae, 0x57, 0x5a, 0x61, 0x20, 0xab, 0xa1, 0x12, 0xda, 0xd1, 0x40, 0xa4, 0xd8, 0xd2, 0xee,
	0x6d, 0xfc, 0xac, 0x54, 0x63, 0x5a, 0x1f, 0x16, 0x3e, 0xfe, 0xc1, 0xc2, 0xb0, 0xea, 0x76, 0x10,
	0xa5, 0xde, 0x28, 0xff, 0x0e, 0x1b, 0x2a, 0x21, 0x43, 0x30, 0xfb, 0x60, 0x5a, 0x55, 0xcc, 0x83,
	0x75, 0xf4, 0xb7, 0x0c, 0xba, 0x0b, 0x09, 0x20, 0xb5, 0xd8, 0x27, 0x77, 0xde, 0xb6, 0x3c, 0x18,
	0x4c, 0xbf, 0x23, 0x72, 0xfe, 0x82, 0xd7, 0x7b, 0x06, 0xd0, 0x85, 0x70, 0x06, 0x1a, 0x27, 0x5e,
	0x2d, 0x33, 0x46, 0xd8, 0xf8, 0x92, 0x8c, 0x38, 0x54, 0x22, 0x89, 0xe6, 0xd1, 0x62, 0xbc, 0x3c,
	0xde, 0x6d, 0x67, 0x93, 0x35, 0xd8, 0xfa, 0x3e, 0xed, 0xa6, 0x69, 0x81, 0x60, 0x7c, 0x47, 0x26,
	0xdd, 0xff, 0xe1, 0xd3, 0x68, 0xa6, 0x9a, 0xe4, 0xdf, 0x3c, 0x5a, 0x1c, 0x2e, 0xcf, 0x76, 0xdb,
	0x59, 0xbc, 0xe7, 0x06, 0x30, 0x2d, 0x86, 0xd4, 0xf8, 0x8a, 0x8c, 0x85, 0xb5, 0x60, 0x93, 0xff,
	0xa8, 0x39, 0xd9, 0x6d, 0x67, 0x47, 0xbd, 0x06, 0xc7, 0x69, 0xd1, 0xc3, 0x31, 0x25, 0xa3, 0x8a,
	0x79, 0x96, 0x8c, 0xe6, 0xd1, 0x62, 0x72, 0x4d, 0x72, 0xbc, 0x4a, 0xe7, 0x77, 0x68, 0xa9, 0x63,
	0xa4, 0x05, 0x12, 0xcb, 0x03, 0x0c, 0x77, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x80, 0xec, 0x10,
	0x58, 0x19, 0x02, 0x00, 0x00,
}
