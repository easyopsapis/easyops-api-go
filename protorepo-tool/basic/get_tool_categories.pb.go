// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_tool_categories.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//GetToolCategories返回
type GetToolCategoriesResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//工具分类
	Data                 []string `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetToolCategoriesResponse) Reset()         { *m = GetToolCategoriesResponse{} }
func (m *GetToolCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GetToolCategoriesResponse) ProtoMessage()    {}
func (*GetToolCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f2486a8cd0ffa7, []int{0}
}
func (m *GetToolCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetToolCategoriesResponse.Unmarshal(m, b)
}
func (m *GetToolCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetToolCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GetToolCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetToolCategoriesResponse.Merge(m, src)
}
func (m *GetToolCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GetToolCategoriesResponse.Size(m)
}
func (m *GetToolCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetToolCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetToolCategoriesResponse proto.InternalMessageInfo

func (m *GetToolCategoriesResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetToolCategoriesResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetToolCategoriesResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetToolCategoriesResponse) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//GetToolCategoriesApi返回
type GetToolCategoriesResponseWrapper struct {
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
	Data                 *GetToolCategoriesResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetToolCategoriesResponseWrapper) Reset()         { *m = GetToolCategoriesResponseWrapper{} }
func (m *GetToolCategoriesResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetToolCategoriesResponseWrapper) ProtoMessage()    {}
func (*GetToolCategoriesResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f2486a8cd0ffa7, []int{1}
}
func (m *GetToolCategoriesResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetToolCategoriesResponseWrapper.Unmarshal(m, b)
}
func (m *GetToolCategoriesResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetToolCategoriesResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetToolCategoriesResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetToolCategoriesResponseWrapper.Merge(m, src)
}
func (m *GetToolCategoriesResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetToolCategoriesResponseWrapper.Size(m)
}
func (m *GetToolCategoriesResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetToolCategoriesResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetToolCategoriesResponseWrapper proto.InternalMessageInfo

func (m *GetToolCategoriesResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetToolCategoriesResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetToolCategoriesResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetToolCategoriesResponseWrapper) GetData() *GetToolCategoriesResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetToolCategoriesResponse)(nil), "basic.GetToolCategoriesResponse")
	proto.RegisterType((*GetToolCategoriesResponseWrapper)(nil), "basic.GetToolCategoriesResponseWrapper")
}

func init() { proto.RegisterFile("get_tool_categories.proto", fileDescriptor_f9f2486a8cd0ffa7) }

var fileDescriptor_f9f2486a8cd0ffa7 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xfc, 0x30,
	0x10, 0xc6, 0xc9, 0xbf, 0xdb, 0xbf, 0x6c, 0x2a, 0x2a, 0x39, 0x48, 0xd7, 0x4b, 0x4b, 0x04, 0xe9,
	0x41, 0xbb, 0xa0, 0x17, 0xf1, 0xb8, 0xb2, 0x78, 0x0f, 0x82, 0xc7, 0x25, 0xed, 0xce, 0xc6, 0x42,
	0xbb, 0x53, 0x92, 0x2c, 0xf8, 0x6e, 0xbe, 0x4b, 0x9f, 0x41, 0xfa, 0x04, 0xd2, 0xb4, 0xee, 0xee,
	0xa5, 0xe2, 0x29, 0x99, 0xef, 0xfb, 0x92, 0x99, 0x1f, 0x43, 0x67, 0x0a, 0xec, 0xca, 0x22, 0x96,
	0xab, 0x5c, 0x5a, 0x50, 0xa8, 0x0b, 0x30, 0x69, 0xad, 0xd1, 0x22, 0xf3, 0x33, 0x69, 0x8a, 0xfc,
	0xea, 0x4e, 0x15, 0xf6, 0x7d, 0x97, 0xa5, 0x39, 0x56, 0x73, 0x85, 0x0a, 0xe7, 0xce, 0xcd, 0x76,
	0x1b, 0x57, 0xb9, 0xc2, 0xdd, 0xfa, 0x57, 0xfc, 0x93, 0xd0, 0xd9, 0x0b, 0xd8, 0x57, 0xc4, 0xf2,
	0x79, 0xff, 0xa3, 0x00, 0x53, 0xe3, 0xd6, 0x00, 0xbb, 0xa6, 0x93, 0x1c, 0xd7, 0x10, 0x92, 0x98,
	0x24, 0xfe, 0xe2, 0xbc, 0x6d, 0xa2, 0x60, 0x83, 0xba, 0x7a, 0xe2, 0x9d, 0xca, 0x85, 0x33, 0xd9,
	0x0d, 0xf5, 0x41, 0x6b, 0xd4, 0xe1, 0xbf, 0x98, 0x24, 0xd3, 0xc5, 0x45, 0xdb, 0x44, 0xa7, 0x7d,
	0xca, 0xc9, 0x5c, 0xf4, 0x36, 0xbb, 0xa5, 0x27, 0x15, 0x18, 0x23, 0x15, 0x84, 0x9e, 0x4b, 0xb2,
	0xb6, 0x89, 0xce, 0xfa, 0xe4, 0x60, 0x70, 0xf1, 0x13, 0xe9, 0x5a, 0xaf, 0xa5, 0x95, 0xe1, 0x24,
	0xf6, 0x92, 0xe9, 0x71, 0xeb, 0x4e, 0xe5, 0xc2, 0x99, 0xfc, 0x8b, 0xd0, 0x78, 0x74, 0xfa, 0x37,
	0x2d, 0xeb, 0x1a, 0xf4, 0xdf, 0x20, 0x1e, 0x69, 0xd0, 0x9d, 0xcb, 0x8f, 0xba, 0x94, 0xc5, 0x76,
	0x40, 0xb9, 0x6c, 0x9b, 0x88, 0x1d, 0xb2, 0x83, 0xc9, 0xc5, 0x71, 0xf4, 0x80, 0xef, 0xfd, 0x8e,
	0xbf, 0xdc, 0x03, 0x91, 0x24, 0xb8, 0x8f, 0x53, 0xb7, 0xae, 0x74, 0x74, 0xfa, 0x11, 0xe4, 0xec,
	0xbf, 0xdb, 0xdb, 0xc3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x8f, 0xd5, 0xa5, 0x0a, 0x02,
	0x00, 0x00,
}
