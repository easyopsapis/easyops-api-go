// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: insert.proto

package mongo

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Insert请求
type InsertRequest struct {
	//
	//表名
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection" form:"collection"`
	//
	//待插入的记录列表
	DocList              []*types.Struct `protobuf:"bytes,2,rep,name=docList,proto3" json:"docList" form:"docList"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InsertRequest) Reset()         { *m = InsertRequest{} }
func (m *InsertRequest) String() string { return proto.CompactTextString(m) }
func (*InsertRequest) ProtoMessage()    {}
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60fb7aaea2101ef9, []int{0}
}
func (m *InsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRequest.Unmarshal(m, b)
}
func (m *InsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRequest.Marshal(b, m, deterministic)
}
func (m *InsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRequest.Merge(m, src)
}
func (m *InsertRequest) XXX_Size() int {
	return xxx_messageInfo_InsertRequest.Size(m)
}
func (m *InsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRequest proto.InternalMessageInfo

func (m *InsertRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *InsertRequest) GetDocList() []*types.Struct {
	if m != nil {
		return m.DocList
	}
	return nil
}

//
//InsertApi返回
type InsertResponseWrapper struct {
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

func (m *InsertResponseWrapper) Reset()         { *m = InsertResponseWrapper{} }
func (m *InsertResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InsertResponseWrapper) ProtoMessage()    {}
func (*InsertResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_60fb7aaea2101ef9, []int{1}
}
func (m *InsertResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertResponseWrapper.Unmarshal(m, b)
}
func (m *InsertResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InsertResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertResponseWrapper.Merge(m, src)
}
func (m *InsertResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InsertResponseWrapper.Size(m)
}
func (m *InsertResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InsertResponseWrapper proto.InternalMessageInfo

func (m *InsertResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InsertResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InsertResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InsertResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InsertRequest)(nil), "mongo.InsertRequest")
	proto.RegisterType((*InsertResponseWrapper)(nil), "mongo.InsertResponseWrapper")
}

func init() { proto.RegisterFile("insert.proto", fileDescriptor_60fb7aaea2101ef9) }

var fileDescriptor_60fb7aaea2101ef9 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x3e, 0x14, 0xd3, 0xf9, 0x15, 0xd8, 0x2c, 0x53, 0xe8, 0x88, 0x20, 0xbb, 0xb1,
	0x83, 0x89, 0x20, 0x7a, 0x37, 0xd8, 0x85, 0xe0, 0x55, 0xbc, 0xf0, 0xba, 0x6b, 0xb3, 0x58, 0x68,
	0x7b, 0x62, 0x92, 0x82, 0x3e, 0x84, 0x4f, 0x28, 0xf4, 0x21, 0xfa, 0x04, 0xb2, 0x93, 0x95, 0x0d,
	0xbd, 0x6a, 0xd3, 0xdf, 0xff, 0x7f, 0xce, 0xaf, 0x21, 0x83, 0xac, 0x34, 0x42, 0xdb, 0x48, 0x69,
	0xb0, 0x40, 0xfb, 0x05, 0x94, 0x12, 0xc6, 0xb7, 0x32, 0xb3, 0xef, 0xd5, 0x2a, 0x4a, 0xa0, 0x98,
	0x49, 0x90, 0x30, 0x43, 0xba, 0xaa, 0xd6, 0x78, 0xc2, 0x03, 0xbe, 0xb9, 0xd6, 0xf8, 0x4a, 0x02,
	0xc8, 0x5c, 0xec, 0x52, 0xc6, 0xea, 0x2a, 0xd9, 0xce, 0x1c, 0x5f, 0xfe, 0xa5, 0xa2, 0x50, 0xf6,
	0xcb, 0x41, 0xf6, 0xed, 0x91, 0xe3, 0x67, 0x34, 0xe0, 0xe2, 0xa3, 0x12, 0xc6, 0xd2, 0x7b, 0x42,
	0x12, 0xc8, 0x73, 0x91, 0xd8, 0x0c, 0xca, 0xc0, 0x9b, 0x78, 0xd3, 0xa3, 0xc5, 0xb0, 0xa9, 0xc3,
	0xf3, 0x35, 0xe8, 0xe2, 0x91, 0xed, 0x18, 0xe3, 0x7b, 0x41, 0xba, 0x24, 0x87, 0x29, 0x24, 0x2f,
	0x99, 0xb1, 0x41, 0x67, 0xd2, 0x9d, 0xfa, 0xf3, 0x8b, 0xc8, 0xed, 0x8d, 0xda, 0xbd, 0xd1, 0x2b,
	0x5a, 0x2d, 0x68, 0x53, 0x87, 0x27, 0x6e, 0xd8, 0xb6, 0xc1, 0x78, 0xdb, 0x65, 0x3f, 0x1e, 0x19,
	0xb6, 0x3e, 0x46, 0x41, 0x69, 0xc4, 0x9b, 0x8e, 0x95, 0x12, 0x9a, 0x5e, 0x93, 0x5e, 0x02, 0xa9,
	0x40, 0xa3, 0xfe, 0xe2, 0xb4, 0xa9, 0x43, 0xbf, 0x35, 0x4a, 0x05, 0xe3, 0x08, 0xe9, 0x03, 0xf1,
	0x37, 0xcf, 0xe5, 0xa7, 0xca, 0xe3, 0xac, 0x0c, 0x3a, 0x68, 0x3f, 0x6a, 0xea, 0x90, 0xee, 0xb2,
	0x5b, 0xc8, 0xf8, 0x7e, 0x94, 0xde, 0x90, 0xbe, 0xd0, 0x1a, 0x74, 0xd0, 0xc5, 0xce, 0x59, 0x53,
	0x87, 0x03, 0xd7, 0xc1, 0xcf, 0x8c, 0x3b, 0x4c, 0x9f, 0x48, 0x2f, 0x8d, 0x6d, 0x1c, 0xf4, 0x26,
	0xde, 0xd4, 0x9f, 0x8f, 0xfe, 0xfd, 0xe4, 0x72, 0x73, 0xb9, 0xfb, 0x7a, 0x9b, 0x34, 0xe3, 0x58,
	0x5a, 0x1d, 0x60, 0xec, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x62, 0x10, 0xe5, 0xf5, 0x01,
	0x00, 0x00,
}