// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fetch_cmdb_business_view.proto

package view

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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
//FetchCmdbBusinessView请求
type FetchCmdbBusinessViewRequest struct {
	//
	//实例ID
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchCmdbBusinessViewRequest) Reset()         { *m = FetchCmdbBusinessViewRequest{} }
func (m *FetchCmdbBusinessViewRequest) String() string { return proto.CompactTextString(m) }
func (*FetchCmdbBusinessViewRequest) ProtoMessage()    {}
func (*FetchCmdbBusinessViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c18f95b0ca5819f0, []int{0}
}
func (m *FetchCmdbBusinessViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchCmdbBusinessViewRequest.Unmarshal(m, b)
}
func (m *FetchCmdbBusinessViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchCmdbBusinessViewRequest.Marshal(b, m, deterministic)
}
func (m *FetchCmdbBusinessViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchCmdbBusinessViewRequest.Merge(m, src)
}
func (m *FetchCmdbBusinessViewRequest) XXX_Size() int {
	return xxx_messageInfo_FetchCmdbBusinessViewRequest.Size(m)
}
func (m *FetchCmdbBusinessViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchCmdbBusinessViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchCmdbBusinessViewRequest proto.InternalMessageInfo

func (m *FetchCmdbBusinessViewRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//FetchCmdbBusinessViewApi返回
type FetchCmdbBusinessViewResponseWrapper struct {
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
	Data                 *topology.View `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchCmdbBusinessViewResponseWrapper) Reset()         { *m = FetchCmdbBusinessViewResponseWrapper{} }
func (m *FetchCmdbBusinessViewResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*FetchCmdbBusinessViewResponseWrapper) ProtoMessage()    {}
func (*FetchCmdbBusinessViewResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c18f95b0ca5819f0, []int{1}
}
func (m *FetchCmdbBusinessViewResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper.Unmarshal(m, b)
}
func (m *FetchCmdbBusinessViewResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper.Marshal(b, m, deterministic)
}
func (m *FetchCmdbBusinessViewResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper.Merge(m, src)
}
func (m *FetchCmdbBusinessViewResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper.Size(m)
}
func (m *FetchCmdbBusinessViewResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_FetchCmdbBusinessViewResponseWrapper proto.InternalMessageInfo

func (m *FetchCmdbBusinessViewResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FetchCmdbBusinessViewResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *FetchCmdbBusinessViewResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FetchCmdbBusinessViewResponseWrapper) GetData() *topology.View {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchCmdbBusinessViewRequest)(nil), "view.FetchCmdbBusinessViewRequest")
	proto.RegisterType((*FetchCmdbBusinessViewResponseWrapper)(nil), "view.FetchCmdbBusinessViewResponseWrapper")
}

func init() { proto.RegisterFile("fetch_cmdb_business_view.proto", fileDescriptor_c18f95b0ca5819f0) }

var fileDescriptor_c18f95b0ca5819f0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x8b, 0xd4, 0x30,
	0x14, 0xa7, 0xda, 0x15, 0xcc, 0x88, 0x7f, 0x0a, 0x4a, 0x19, 0xc4, 0x2e, 0x51, 0x64, 0x2f, 0x6d,
	0xc0, 0x45, 0x11, 0x8f, 0x23, 0x0a, 0x1e, 0xbc, 0x14, 0xd4, 0xe3, 0x90, 0x26, 0x6f, 0xba, 0x61,
	0xdb, 0xbc, 0x98, 0xa4, 0x5b, 0xf7, 0xcb, 0xf6, 0xe0, 0x47, 0xe8, 0x27, 0x90, 0xa6, 0xa3, 0xd3,
	0x8b, 0xb7, 0x9e, 0xf2, 0x5e, 0x7e, 0x7f, 0x78, 0xff, 0xc8, 0x8b, 0x03, 0x78, 0x71, 0xb5, 0x17,
	0xad, 0xac, 0xf6, 0x55, 0xe7, 0x94, 0x06, 0xe7, 0xf6, 0x37, 0x0a, 0xfa, 0xc2, 0x58, 0xf4, 0x98,
	0xc4, 0x53, 0xbc, 0xcd, 0x6b, 0xe5, 0xaf, 0xba, 0xaa, 0x10, 0xd8, 0xb2, 0x1a, 0x6b, 0x64, 0x01,
	0xac, 0xba, 0x43, 0xc8, 0x42, 0x12, 0xa2, 0x59, 0xb4, 0xfd, 0x5a, 0x63, 0x01, 0xdc, 0xdd, 0xa2,
	0x71, 0x45, 0x83, 0x82, 0x37, 0x4c, 0xa0, 0xf6, 0x96, 0x0b, 0xef, 0x66, 0xa5, 0x05, 0x83, 0x79,
	0x8b, 0x12, 0x1a, 0xc7, 0x8e, 0x44, 0x16, 0x52, 0xe6, 0xd1, 0x60, 0x83, 0xf5, 0x2d, 0xd3, 0x28,
	0x61, 0x45, 0xbb, 0x46, 0xe9, 0xeb, 0x15, 0xed, 0xb8, 0x05, 0xbe, 0x6a, 0xb3, 0x7e, 0xcd, 0x66,
	0x4f, 0xfb, 0xdb, 0xbe, 0x5b, 0x6c, 0xae, 0xed, 0x95, 0xbf, 0xc6, 0x9e, 0xd5, 0x98, 0x07, 0x30,
	0xbf, 0xe1, 0x8d, 0x92, 0xdc, 0xa3, 0x75, 0xec, 0x5f, 0x38, 0xeb, 0xe8, 0x37, 0xf2, 0xfc, 0xf3,
	0x74, 0x19, 0x1f, 0x5b, 0x59, 0xed, 0x8e, 0x77, 0xf1, 0x5d, 0x41, 0x5f, 0xc2, 0xcf, 0x0e, 0x9c,
	0x4f, 0xde, 0x12, 0xa2, 0xb4, 0xf3, 0x5c, 0x0b, 0xf8, 0x22, 0xd3, 0xe8, 0x3c, 0xba, 0xb8, 0xbf,
	0x7b, 0x3a, 0x0e, 0xd9, 0x93, 0x03, 0xda, 0xf6, 0x03, 0x3d, 0x61, 0xb4, 0x5c, 0x10, 0xe9, 0xef,
	0x88, 0xbc, 0xfa, 0x8f, 0xaf, 0x33, 0xa8, 0x1d, 0xfc, 0xb0, 0xdc, 0x18, 0xb0, 0xc9, 0x4b, 0x12,
	0x0b, 0x94, 0x10, 0x9c, 0xcf, 0x76, 0x8f, 0xc6, 0x21, 0xdb, 0xcc, 0xce, 0xd3, 0x2f, 0x2d, 0x03,
	0x98, 0xbc, 0x27, 0x9b, 0xe9, 0xfd, 0xf4, 0xcb, 0x34, 0x5c, 0xe9, 0xf4, 0x4e, 0xa8, 0xe2, 0xd9,
	0x38, 0x64, 0xc9, 0x89, 0x7b, 0x04, 0x69, 0xb9, 0xa4, 0x26, 0xaf, 0xc9, 0x19, 0x58, 0x8b, 0x36,
	0xbd, 0x1b, 0x34, 0x8f, 0xc7, 0x21, 0x7b, 0x30, 0x6b, 0xc2, 0x37, 0x2d, 0x67, 0x38, 0xb9, 0x24,
	0xb1, 0xe4, 0x9e, 0xa7, 0xf1, 0x79, 0x74, 0xb1, 0x79, 0xf3, 0xb0, 0xf8, 0x3b, 0xe2, 0x62, 0xaa,
	0x79, 0x59, 0xd6, 0xc4, 0xa2, 0x65, 0x20, 0x57, 0xf7, 0xc2, 0x08, 0x2f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0x23, 0x53, 0xb6, 0x5c, 0x03, 0x00, 0x00,
}
