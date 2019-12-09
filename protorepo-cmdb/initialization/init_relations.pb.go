// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: init_relations.proto

package initialization

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
//InitRelations请求
type InitRelationsRequest struct {
	//
	//org
	Org                  int32    `protobuf:"varint,1,opt,name=org,proto3" json:"org" form:"org"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitRelationsRequest) Reset()         { *m = InitRelationsRequest{} }
func (m *InitRelationsRequest) String() string { return proto.CompactTextString(m) }
func (*InitRelationsRequest) ProtoMessage()    {}
func (*InitRelationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_893171aa83c99585, []int{0}
}
func (m *InitRelationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRelationsRequest.Unmarshal(m, b)
}
func (m *InitRelationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRelationsRequest.Marshal(b, m, deterministic)
}
func (m *InitRelationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRelationsRequest.Merge(m, src)
}
func (m *InitRelationsRequest) XXX_Size() int {
	return xxx_messageInfo_InitRelationsRequest.Size(m)
}
func (m *InitRelationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRelationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRelationsRequest proto.InternalMessageInfo

func (m *InitRelationsRequest) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

//
//InitRelationsApi返回
type InitRelationsResponseWrapper struct {
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

func (m *InitRelationsResponseWrapper) Reset()         { *m = InitRelationsResponseWrapper{} }
func (m *InitRelationsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InitRelationsResponseWrapper) ProtoMessage()    {}
func (*InitRelationsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_893171aa83c99585, []int{1}
}
func (m *InitRelationsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRelationsResponseWrapper.Unmarshal(m, b)
}
func (m *InitRelationsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRelationsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InitRelationsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRelationsResponseWrapper.Merge(m, src)
}
func (m *InitRelationsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InitRelationsResponseWrapper.Size(m)
}
func (m *InitRelationsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRelationsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InitRelationsResponseWrapper proto.InternalMessageInfo

func (m *InitRelationsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InitRelationsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InitRelationsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InitRelationsResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InitRelationsRequest)(nil), "initialization.InitRelationsRequest")
	proto.RegisterType((*InitRelationsResponseWrapper)(nil), "initialization.InitRelationsResponseWrapper")
}

func init() { proto.RegisterFile("init_relations.proto", fileDescriptor_893171aa83c99585) }

var fileDescriptor_893171aa83c99585 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xa9, 0xeb, 0x04, 0x53, 0x99, 0x12, 0xc6, 0x28, 0x53, 0x68, 0x89, 0x20, 0xbd, 0xd8,
	0x81, 0x5e, 0x86, 0xde, 0x06, 0x3b, 0x78, 0xcd, 0xc5, 0xa3, 0xa4, 0x5b, 0x16, 0x03, 0x6d, 0xfe,
	0x31, 0x4d, 0x41, 0xfd, 0xb0, 0x3d, 0xf8, 0x11, 0xfa, 0x09, 0x24, 0xc9, 0xc6, 0x86, 0xa7, 0xe4,
	0xff, 0xde, 0xef, 0xfd, 0x13, 0x1e, 0x9a, 0x4a, 0x25, 0xed, 0xbb, 0xe1, 0x35, 0xb3, 0x12, 0x54,
	0x5b, 0x6a, 0x03, 0x16, 0xf0, 0xc4, 0xa9, 0x92, 0xd5, 0xf2, 0xc7, 0xcb, 0xf3, 0x07, 0x21, 0xed,
	0x47, 0x57, 0x95, 0x1b, 0x68, 0x16, 0x02, 0x04, 0x2c, 0x3c, 0x56, 0x75, 0x3b, 0x3f, 0xf9, 0xc1,
	0xdf, 0x42, 0x7c, 0x7e, 0x23, 0x00, 0x44, 0xcd, 0x8f, 0x14, 0x6f, 0xb4, 0xfd, 0x0e, 0x26, 0x59,
	0xa2, 0xe9, 0xab, 0x92, 0x96, 0x1e, 0x9e, 0xa4, 0xfc, 0xb3, 0xe3, 0xad, 0xc5, 0x39, 0x1a, 0x81,
	0x11, 0x69, 0x94, 0x47, 0xc5, 0x78, 0x35, 0x19, 0xfa, 0x0c, 0xed, 0xc0, 0x34, 0xcf, 0x04, 0x8c,
	0x20, 0xd4, 0x59, 0xe4, 0x37, 0x42, 0xb7, 0xff, 0xa2, 0xad, 0x06, 0xd5, 0xf2, 0x37, 0xc3, 0xb4,
	0xe6, 0x06, 0xdf, 0xa1, 0x78, 0x03, 0x5b, 0xbe, 0xdf, 0x71, 0x35, 0xf4, 0x59, 0x12, 0x76, 0x38,
	0x95, 0x50, 0x6f, 0xe2, 0x25, 0x4a, 0xdc, 0xb9, 0xfe, 0xd2, 0x35, 0x93, 0x2a, 0x3d, 0xcb, 0xa3,
	0xe2, 0x62, 0x35, 0x1b, 0xfa, 0x0c, 0x1f, 0xd9, 0xbd, 0x49, 0xe8, 0x29, 0x8a, 0xef, 0xd1, 0x98,
	0x1b, 0x03, 0x26, 0x1d, 0xf9, 0xcc, 0xf5, 0xd0, 0x67, 0x97, 0x21, 0xe3, 0x65, 0x42, 0x83, 0x8d,
	0x5f, 0x50, 0xbc, 0x65, 0x96, 0xa5, 0x71, 0x1e, 0x15, 0xc9, 0xe3, 0xac, 0x0c, 0x6d, 0x94, 0x87,
	0x36, 0xca, 0xb5, 0x6b, 0xe3, 0xf4, 0x7b, 0x8e, 0x26, 0xd4, 0x87, 0xaa, 0x73, 0x8f, 0x3d, 0xfd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x38, 0x2a, 0x3a, 0x99, 0x01, 0x00, 0x00,
}