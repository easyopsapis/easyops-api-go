// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package template

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
//DeleteTemplate请求
type DeleteTemplateRequest struct {
	//
	//要删除的模版id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTemplateRequest) Reset()         { *m = DeleteTemplateRequest{} }
func (m *DeleteTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTemplateRequest) ProtoMessage()    {}
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTemplateRequest.Unmarshal(m, b)
}
func (m *DeleteTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTemplateRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTemplateRequest.Merge(m, src)
}
func (m *DeleteTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTemplateRequest.Size(m)
}
func (m *DeleteTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTemplateRequest proto.InternalMessageInfo

func (m *DeleteTemplateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DeleteTemplateApi返回
type DeleteTemplateResponseWrapper struct {
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

func (m *DeleteTemplateResponseWrapper) Reset()         { *m = DeleteTemplateResponseWrapper{} }
func (m *DeleteTemplateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteTemplateResponseWrapper) ProtoMessage()    {}
func (*DeleteTemplateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteTemplateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTemplateResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteTemplateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTemplateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteTemplateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTemplateResponseWrapper.Merge(m, src)
}
func (m *DeleteTemplateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteTemplateResponseWrapper.Size(m)
}
func (m *DeleteTemplateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTemplateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTemplateResponseWrapper proto.InternalMessageInfo

func (m *DeleteTemplateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteTemplateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteTemplateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteTemplateResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteTemplateRequest)(nil), "template.DeleteTemplateRequest")
	proto.RegisterType((*DeleteTemplateResponseWrapper)(nil), "template.DeleteTemplateResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xe9, 0xdc, 0xc4, 0x65, 0x43, 0x47, 0xc0, 0x51, 0x26, 0xd2, 0x19, 0x45, 0xf6, 0xd2,
	0xd6, 0x39, 0x10, 0xff, 0xbc, 0x0d, 0xf7, 0xe2, 0x63, 0x11, 0x7c, 0x10, 0x85, 0x6c, 0xc9, 0x6a,
	0xb0, 0xdd, 0x8d, 0x69, 0xe6, 0xfc, 0x83, 0x5f, 0xb5, 0x82, 0x7e, 0x83, 0x7e, 0x02, 0x69, 0x3a,
	0x5d, 0xf1, 0x29, 0xb9, 0xf7, 0xfc, 0xce, 0xe1, 0x70, 0x51, 0x93, 0xf1, 0x88, 0x6b, 0xee, 0x49,
	0x05, 0x1a, 0xf0, 0x86, 0xe6, 0xb1, 0x8c, 0xa8, 0xe6, 0x1d, 0x37, 0x14, 0xfa, 0x61, 0x3e, 0xf6,
	0x26, 0x10, 0xfb, 0x21, 0x84, 0xe0, 0x1b, 0x60, 0x3c, 0x9f, 0x9a, 0xc9, 0x0c, 0xe6, 0x57, 0x18,
	0x3b, 0x27, 0x25, 0x3c, 0x5e, 0x08, 0xfd, 0x08, 0x0b, 0x3f, 0x04, 0xd7, 0x88, 0xee, 0x33, 0x8d,
	0x04, 0xa3, 0x1a, 0x54, 0xe2, 0xff, 0x7d, 0x97, 0xbe, 0x9d, 0x10, 0x20, 0x8c, 0xf8, 0x2a, 0x9d,
	0xc7, 0x52, 0xbf, 0x16, 0x22, 0xb9, 0x42, 0xdb, 0x97, 0xa6, 0xdd, 0xf5, 0xb2, 0x55, 0xc0, 0x9f,
	0xe6, 0x3c, 0xd1, 0xb8, 0x8f, 0x2a, 0x82, 0xd9, 0x56, 0xd7, 0xea, 0xd5, 0x87, 0x7b, 0x59, 0xea,
	0xd4, 0xa7, 0xa0, 0xe2, 0x73, 0x22, 0x18, 0xf9, 0xfa, 0x74, 0x5a, 0x68, 0xf3, 0xfe, 0xf6, 0xc8,
	0x3d, 0xa3, 0xee, 0xdb, 0xdd, 0x7b, 0x7f, 0xf0, 0x71, 0x10, 0x54, 0x04, 0x23, 0xdf, 0x16, 0xda,
	0xfd, 0x1f, 0x96, 0x48, 0x98, 0x25, 0xfc, 0x46, 0x51, 0x29, 0xb9, 0xc2, 0xfb, 0xa8, 0x3a, 0x01,
	0xc6, 0x4d, 0x6c, 0x6d, 0xb8, 0x95, 0xa5, 0x4e, 0xa3, 0x88, 0xcd, 0xb7, 0x24, 0x30, 0x22, 0x3e,
	0x45, 0x8d, 0xfc, 0x1d, 0xbd, 0xc8, 0x88, 0x8a, 0x99, 0x5d, 0x31, 0x15, 0xda, 0x59, 0xea, 0xe0,
	0x15, 0xbb, 0x14, 0x49, 0x50, 0x46, 0xf1, 0x21, 0xaa, 0x71, 0xa5, 0x40, 0xd9, 0x6b, 0xc6, 0xd3,
	0xca, 0x52, 0xa7, 0x59, 0x78, 0xcc, 0x9a, 0x04, 0x85, 0x8c, 0x2f, 0x50, 0x95, 0x51, 0x4d, 0xed,
	0x6a, 0xd7, 0xea, 0x35, 0x8e, 0xdb, 0x5e, 0x71, 0x20, 0xef, 0xf7, 0x40, 0xde, 0x28, 0x3f, 0x50,
	0xb9, 0x5e, 0x4e, 0x93, 0xc0, 0x98, 0xc6, 0xeb, 0x06, 0x1b, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0xfd, 0xdb, 0x82, 0xd6, 0x01, 0x00, 0x00,
}