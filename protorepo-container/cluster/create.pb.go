// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package cluster

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	container "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
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
//CreateApi返回
type CreateResponseWrapper struct {
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
	Data                 *container.Cluster `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseWrapper.Unmarshal(m, b)
}
func (m *CreateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseWrapper.Merge(m, src)
}
func (m *CreateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateResponseWrapper.Size(m)
}
func (m *CreateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseWrapper proto.InternalMessageInfo

func (m *CreateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponseWrapper) GetData() *container.Cluster {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateResponseWrapper)(nil), "cluster.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xb6, 0x8a, 0x69, 0x41, 0x59, 0x50, 0x42, 0x2f, 0x2d, 0x11, 0xa4, 0x97, 0x64,
	0x41, 0x41, 0xc5, 0x63, 0x8b, 0x0f, 0x60, 0x2e, 0x9e, 0xb7, 0x9b, 0x69, 0x0c, 0x6e, 0x32, 0xcb,
	0xec, 0xd6, 0xea, 0x93, 0x7a, 0xeb, 0x43, 0xf4, 0x09, 0xa4, 0xb3, 0xa1, 0xe6, 0x94, 0xf9, 0xe7,
	0xff, 0xbf, 0xec, 0xcc, 0xc4, 0x63, 0x4d, 0xa0, 0x3c, 0xe4, 0x96, 0xd0, 0xa3, 0x38, 0xd7, 0x66,
	0xe3, 0x3c, 0xd0, 0x24, 0xab, 0x6a, 0xff, 0xb1, 0x59, 0xe5, 0x1a, 0x1b, 0x59, 0x61, 0x85, 0x92,
	0xfd, 0xd5, 0x66, 0xcd, 0x8a, 0x05, 0x57, 0x81, 0x9b, 0xbc, 0x55, 0x98, 0x83, 0x72, 0x3f, 0x68,
	0x5d, 0x6e, 0x50, 0x2b, 0x23, 0x35, 0xb6, 0x9e, 0x94, 0xf6, 0x2e, 0x90, 0x04, 0x16, 0xb3, 0x06,
	0x4b, 0x30, 0x4e, 0x76, 0x41, 0xc9, 0x92, 0x83, 0xaa, 0x6e, 0x81, 0x64, 0xf7, 0x76, 0xf7, 0xcb,
	0xc7, 0xde, 0x04, 0xcd, 0xb6, 0xf6, 0x9f, 0xb8, 0x95, 0x15, 0x66, 0x6c, 0x66, 0x5f, 0xca, 0xd4,
	0xa5, 0xf2, 0x48, 0x4e, 0x1e, 0xcb, 0xc0, 0xa5, 0xbf, 0x51, 0x7c, 0xbd, 0xe4, 0x9d, 0x0a, 0x70,
	0x16, 0x5b, 0x07, 0xef, 0xa4, 0xac, 0x05, 0x12, 0xb7, 0xf1, 0x40, 0x63, 0x09, 0x49, 0x34, 0x8b,
	0xe6, 0xc3, 0xc5, 0xe5, 0x7e, 0x37, 0x1d, 0xad, 0x91, 0x9a, 0x97, 0xf4, 0xd0, 0x4d, 0x0b, 0x36,
	0xc5, 0x73, 0x3c, 0x3a, 0x7c, 0x5f, 0xbf, 0xad, 0x51, 0x75, 0x9b, 0x9c, 0xcc, 0xa2, 0xf9, 0xc5,
	0xe2, 0x66, 0xbf, 0x9b, 0x8a, 0xff, 0x6c, 0x67, 0xa6, 0x45, 0x3f, 0x2a, 0xee, 0xe2, 0x21, 0x10,
	0x21, 0x25, 0xa7, 0xcc, 0x5c, 0xed, 0x77, 0xd3, 0x71, 0x60, 0xb8, 0x9d, 0x16, 0xc1, 0x16, 0x4f,
	0xf1, 0xa0, 0x54, 0x5e, 0x25, 0x83, 0x59, 0x34, 0x1f, 0xdd, 0x8b, 0xfc, 0x78, 0x80, 0x7c, 0x19,
	0x0e, 0xd0, 0x1f, 0xed, 0x90, 0x4c, 0x0b, 0x06, 0x56, 0x67, 0xbc, 0xe0, 0xc3, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0xd6, 0xfb, 0xca, 0xb3, 0x01, 0x00, 0x00,
}
