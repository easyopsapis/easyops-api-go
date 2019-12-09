// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package hh

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/flow"
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
//TestTest请求
type TestTestRequest struct {
	//
	//版本号(时间戳-秒)
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version" form:"version"`
	//
	//流程标签
	Tags                 []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags" form:"tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestTestRequest) Reset()         { *m = TestTestRequest{} }
func (m *TestTestRequest) String() string { return proto.CompactTextString(m) }
func (*TestTestRequest) ProtoMessage()    {}
func (*TestTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *TestTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTestRequest.Unmarshal(m, b)
}
func (m *TestTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTestRequest.Marshal(b, m, deterministic)
}
func (m *TestTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTestRequest.Merge(m, src)
}
func (m *TestTestRequest) XXX_Size() int {
	return xxx_messageInfo_TestTestRequest.Size(m)
}
func (m *TestTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestTestRequest proto.InternalMessageInfo

func (m *TestTestRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TestTestRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

//
//TestTestApi返回
type TestTestResponseWrapper struct {
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
	Data                 *flow.Flow `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TestTestResponseWrapper) Reset()         { *m = TestTestResponseWrapper{} }
func (m *TestTestResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*TestTestResponseWrapper) ProtoMessage()    {}
func (*TestTestResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *TestTestResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTestResponseWrapper.Unmarshal(m, b)
}
func (m *TestTestResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTestResponseWrapper.Marshal(b, m, deterministic)
}
func (m *TestTestResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTestResponseWrapper.Merge(m, src)
}
func (m *TestTestResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_TestTestResponseWrapper.Size(m)
}
func (m *TestTestResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTestResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TestTestResponseWrapper proto.InternalMessageInfo

func (m *TestTestResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TestTestResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *TestTestResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TestTestResponseWrapper) GetData() *flow.Flow {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TestTestRequest)(nil), "hh.TestTestRequest")
	proto.RegisterType((*TestTestResponseWrapper)(nil), "hh.TestTestResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xdf, 0x8a, 0x9c, 0x30,
	0x14, 0xc6, 0x71, 0xfe, 0xb4, 0x4c, 0x66, 0xe8, 0x94, 0x5c, 0xb4, 0x32, 0x14, 0x94, 0x14, 0x8a,
	0x17, 0xd5, 0x40, 0x0b, 0xa5, 0xf4, 0x72, 0xa0, 0x85, 0xde, 0xec, 0x85, 0x2c, 0xec, 0xe5, 0x92,
	0xd1, 0x4c, 0x46, 0x36, 0x7a, 0xb2, 0x49, 0x1c, 0x77, 0x1f, 0x73, 0x5f, 0xc0, 0x87, 0xf0, 0x09,
	0x16, 0xa3, 0xbb, 0xe3, 0x03, 0xec, 0x85, 0x9a, 0x73, 0x7e, 0xdf, 0xf7, 0xc5, 0x73, 0xd0, 0x26,
	0xd3, 0x9c, 0x59, 0x9e, 0x28, 0x0d, 0x16, 0xf0, 0xec, 0x74, 0xda, 0xc5, 0xa2, 0xb0, 0xa7, 0xfa,
	0x90, 0x64, 0x50, 0x52, 0x01, 0x02, 0xa8, 0x43, 0x87, 0xfa, 0xe8, 0x2a, 0x57, 0xb8, 0xd3, 0x60,
	0xd9, 0x5d, 0x09, 0x48, 0x38, 0x33, 0x8f, 0xa0, 0x4c, 0x22, 0x21, 0x63, 0x92, 0x66, 0x50, 0x59,
	0xcd, 0x32, 0x6b, 0x06, 0xa7, 0xe6, 0x0a, 0xe2, 0x12, 0x72, 0x2e, 0x0d, 0x1d, 0x85, 0xd4, 0x95,
	0xf4, 0x28, 0xa1, 0x71, 0xaf, 0x5b, 0x63, 0xb9, 0x1a, 0xf3, 0xfe, 0xbf, 0x41, 0xde, 0x18, 0xf5,
	0x6b, 0x32, 0x49, 0xd9, 0x14, 0xf6, 0x0e, 0x1a, 0x2a, 0x20, 0x76, 0x30, 0x3e, 0x33, 0x59, 0xe4,
	0xcc, 0x82, 0x36, 0xf4, 0xf5, 0x38, 0xfa, 0xbe, 0x08, 0x00, 0x21, 0xf9, 0x65, 0x70, 0x63, 0x75,
	0x9d, 0xd9, 0x81, 0x92, 0x1c, 0x6d, 0xaf, 0xb9, 0xb1, 0xfd, 0x93, 0xf2, 0xfb, 0x9a, 0x1b, 0x8b,
	0xbf, 0xa3, 0xf7, 0x67, 0xae, 0x4d, 0x01, 0x95, 0xef, 0x85, 0x5e, 0xb4, 0xdc, 0xe3, 0xae, 0x0d,
	0x3e, 0x1c, 0x41, 0x97, 0x7f, 0xc8, 0x08, 0x48, 0xfa, 0x22, 0xc1, 0x5f, 0xd1, 0xc2, 0x32, 0x61,
	0xfc, 0x59, 0x38, 0x8f, 0x56, 0xfb, 0x6d, 0xd7, 0x06, 0xeb, 0x41, 0xda, 0x77, 0x49, 0xea, 0x20,
	0x79, 0xf2, 0xd0, 0xe7, 0xcb, 0x35, 0x46, 0x41, 0x65, 0xf8, 0x8d, 0x66, 0x4a, 0x71, 0xdd, 0x07,
	0x64, 0x90, 0xf3, 0xf1, 0xae, 0x49, 0x40, 0xdf, 0x25, 0xa9, 0x83, 0xf8, 0x37, 0x5a, 0xf7, 0xdf,
	0xbf, 0x0f, 0x4a, 0xb2, 0xa2, 0xf2, 0x67, 0xa1, 0x17, 0xad, 0xf6, 0x9f, 0xba, 0x36, 0xc0, 0x17,
	0xed, 0x08, 0x49, 0x3a, 0x95, 0xe2, 0x6f, 0x68, 0xc9, 0xb5, 0x06, 0xed, 0xcf, 0x9d, 0xe7, 0x63,
	0xd7, 0x06, 0x9b, 0xc1, 0xe3, 0xda, 0x24, 0x1d, 0x30, 0xa6, 0x68, 0x91, 0x33, 0xcb, 0xfc, 0x45,
	0xe8, 0x45, 0xeb, 0x1f, 0x28, 0x71, 0x9b, 0xff, 0x27, 0xa1, 0x99, 0xfe, 0x52, 0xaf, 0x20, 0xa9,
	0x13, 0x1e, 0xde, 0xb9, 0x05, 0xfe, 0x7c, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x20, 0x83, 0x6b,
	0x74, 0x02, 0x00, 0x00,
}
