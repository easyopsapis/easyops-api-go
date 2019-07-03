// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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
//Create请求
type CreateRequest struct {
	//
	//流水线id
	PipelineId string `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id" form:"pipeline_id"`
	//
	//构建任务
	Build                *pipeline.Build `protobuf:"bytes,2,opt,name=build,proto3" json:"build" form:"build"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

func (m *CreateRequest) GetBuild() *pipeline.Build {
	if m != nil {
		return m.Build
	}
	return nil
}

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
	Data                 *pipeline.Build `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
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

func (m *CreateResponseWrapper) GetData() *pipeline.Build {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "build.CreateRequest")
	proto.RegisterType((*CreateResponseWrapper)(nil), "build.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0x26, 0x3a, 0x23, 0x6c, 0x67, 0x75, 0x97, 0x06, 0x25, 0xec, 0x25, 0x4b, 0x2b, 0xb2, 0x08,
	0x49, 0xab, 0x2b, 0xfe, 0x1d, 0x23, 0x1e, 0x44, 0x10, 0x0c, 0x88, 0x07, 0xd1, 0xa1, 0x93, 0xf4,
	0xc4, 0xc6, 0x24, 0xd5, 0x76, 0x77, 0x1c, 0x7f, 0xf0, 0x31, 0x7c, 0x33, 0xcf, 0x23, 0xf8, 0x08,
	0xf3, 0x04, 0x92, 0xea, 0x44, 0x73, 0xf0, 0x38, 0xa7, 0x74, 0xd5, 0xf7, 0x53, 0x5f, 0x55, 0xc8,
	0x61, 0x69, 0xa4, 0x70, 0x32, 0xd5, 0x06, 0x1c, 0xd0, 0x65, 0xd1, 0xab, 0xa6, 0x3a, 0x49, 0x6a,
	0xe5, 0xde, 0xf7, 0x45, 0x5a, 0x42, 0xcb, 0x6b, 0xa8, 0x81, 0x23, 0x5a, 0xf4, 0x6b, 0xac, 0xb0,
	0xc0, 0x97, 0x57, 0x9d, 0xbc, 0xa8, 0x21, 0x95, 0xc2, 0x7e, 0x01, 0x6d, 0xd3, 0x06, 0x4a, 0xd1,
	0xf0, 0x12, 0x3a, 0x67, 0x44, 0xe9, 0xac, 0x57, 0x1a, 0xa9, 0x21, 0x69, 0xa1, 0x92, 0x8d, 0xe5,
	0x23, 0x91, 0x63, 0xc9, 0xb5, 0xd2, 0xb2, 0x51, 0x9d, 0xe4, 0x38, 0x78, 0xf4, 0x7b, 0xb9, 0x07,
	0xbf, 0x5a, 0xb9, 0x55, 0x2b, 0x9d, 0x18, 0x2d, 0x5f, 0xed, 0x2b, 0xe2, 0xca, 0x3a, 0xe1, 0x7a,
	0x3b, 0xda, 0xde, 0x9f, 0x1d, 0xaa, 0xdd, 0x28, 0xf7, 0x01, 0x36, 0xbc, 0x86, 0x04, 0xc1, 0xe4,
	0x93, 0x68, 0x54, 0x25, 0x1c, 0x18, 0xcb, 0xff, 0x3e, 0xbd, 0x8e, 0xfd, 0x08, 0xc8, 0xe5, 0x27,
	0x78, 0xf8, 0x5c, 0x7e, 0xec, 0xa5, 0x75, 0xf4, 0x39, 0x09, 0xa7, 0x41, 0x2b, 0x55, 0x45, 0xc1,
	0x69, 0x70, 0x76, 0x90, 0xdd, 0xda, 0x6d, 0x63, 0xba, 0x06, 0xd3, 0x3e, 0x66, 0x33, 0x90, 0xfd,
	0xfe, 0x15, 0x1f, 0x93, 0x2b, 0xef, 0xde, 0xdc, 0x4e, 0x1e, 0x89, 0xe4, 0xeb, 0xdb, 0x6f, 0x77,
	0xce, 0xbf, 0xdf, 0xc8, 0xc9, 0xc4, 0x78, 0x56, 0xd1, 0x07, 0xc4, 0xff, 0xc8, 0xe8, 0xc2, 0x69,
	0x70, 0x16, 0xde, 0x3d, 0x4a, 0x27, 0x2c, 0xcd, 0x86, 0x76, 0x76, 0xbc, 0xdb, 0xc6, 0x87, 0xde,
	0x17, 0x79, 0x2c, 0xf7, 0x7c, 0xf6, 0x33, 0x20, 0x57, 0xa7, 0x5c, 0x56, 0x43, 0x67, 0xe5, 0x6b,
	0x23, 0xb4, 0x96, 0x86, 0x5e, 0x27, 0x8b, 0x12, 0x2a, 0x89, 0xc1, 0x96, 0xd9, 0xd1, 0x6e, 0x1b,
	0x87, 0xde, 0x60, 0xe8, 0xb2, 0x1c, 0x41, 0xfa, 0x90, 0x84, 0xc3, 0xf7, 0xe9, 0x67, 0xdd, 0x08,
	0xd5, 0xe1, 0xf4, 0x83, 0xec, 0xda, 0xbf, 0x25, 0x66, 0x20, 0xcb, 0xe7, 0x54, 0x7a, 0x93, 0x2c,
	0xa5, 0x31, 0x60, 0xa2, 0x8b, 0xa8, 0x99, 0x05, 0xc4, 0x36, 0xcb, 0x3d, 0x4c, 0xef, 0x91, 0x45,
	0x25, 0x9c, 0x88, 0x16, 0xff, 0x5f, 0x6c, 0x96, 0x6b, 0xa0, 0xb1, 0x1c, 0xd9, 0xc5, 0x25, 0xbc,
	0xfa, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb0, 0x64, 0xc3, 0xed, 0x02, 0x00, 0x00,
}
