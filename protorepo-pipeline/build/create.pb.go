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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Create请求
type CreateRequest struct {
	//
	//项目id
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id" form:"project_id"`
	//
	//流水线id
	PipelineId string `protobuf:"bytes,2,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id" form:"pipeline_id"`
	//
	//构建任务
	Build                *pipeline.Build `protobuf:"bytes,3,opt,name=build,proto3" json:"build" form:"build"`
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

func (m *CreateRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

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
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0xa9, 0xce, 0x08, 0x93, 0x59, 0xdd, 0x35, 0xa0, 0x0c, 0x7b, 0xd3, 0x25, 0x8a, 0xac,
	0x42, 0x1b, 0x75, 0xc5, 0x7f, 0x97, 0x23, 0x82, 0x8b, 0x20, 0x58, 0x10, 0x2f, 0x44, 0x87, 0x4c,
	0x93, 0xad, 0xd1, 0x76, 0x4e, 0x4c, 0x52, 0xd7, 0x3f, 0xf8, 0x96, 0x5e, 0xcf, 0x82, 0x8f, 0xd0,
	0x27, 0x90, 0x9e, 0xb4, 0x3b, 0xbd, 0xd8, 0xcb, 0xb9, 0x6a, 0xce, 0xf9, 0xbe, 0xf3, 0xeb, 0xf9,
	0xd2, 0x92, 0x9d, 0xdc, 0x2a, 0xe1, 0x55, 0x6a, 0x2c, 0x78, 0xa0, 0xe3, 0x65, 0xad, 0x4b, 0xb9,
	0x9f, 0x14, 0xda, 0x7f, 0xae, 0x97, 0x69, 0x0e, 0x15, 0x2f, 0xa0, 0x00, 0x8e, 0xea, 0xb2, 0x3e,
	0xc1, 0x0a, 0x0b, 0x3c, 0x85, 0xa9, 0xfd, 0x37, 0x05, 0xa4, 0x4a, 0xb8, 0x9f, 0x60, 0x5c, 0x5a,
	0x42, 0x2e, 0x4a, 0x9e, 0xc3, 0xca, 0x5b, 0x91, 0x7b, 0x17, 0x26, 0xad, 0x32, 0x90, 0x54, 0x20,
	0x55, 0xe9, 0x78, 0x67, 0xe4, 0x58, 0x72, 0xa3, 0x8d, 0x2a, 0xf5, 0x4a, 0x71, 0x7c, 0x71, 0xc7,
	0x7b, 0xbb, 0x05, 0x5e, 0xa1, 0xfd, 0xa2, 0x52, 0x5e, 0x74, 0xc8, 0x77, 0xdb, 0x5a, 0x71, 0xe1,
	0xbc, 0xf0, 0xb5, 0xeb, 0xb0, 0x8f, 0x07, 0x17, 0x55, 0x9d, 0x6a, 0xff, 0x15, 0x4e, 0x79, 0x01,
	0x09, 0x8a, 0xc9, 0x77, 0x51, 0x6a, 0x29, 0x3c, 0x58, 0xc7, 0xcf, 0x8f, 0x61, 0x8e, 0x9d, 0x45,
	0xe4, 0xea, 0x0b, 0xbc, 0xf8, 0x4c, 0x7d, 0xab, 0x95, 0xf3, 0xf4, 0x15, 0x21, 0xc6, 0xc2, 0x17,
	0x95, 0xfb, 0x85, 0x96, 0xb3, 0xe8, 0x20, 0x3a, 0x9c, 0xcc, 0xef, 0x36, 0xeb, 0xf8, 0xfa, 0x09,
	0xd8, 0xea, 0x39, 0xdb, 0x68, 0xec, 0xdf, 0x59, 0xbc, 0x47, 0xae, 0x7d, 0xfa, 0x70, 0x3f, 0x79,
	0x26, 0x92, 0x5f, 0x1f, 0x7f, 0x3f, 0x38, 0xfa, 0x73, 0x3b, 0x9b, 0x74, 0x86, 0x63, 0x49, 0x5f,
	0x93, 0x69, 0xbf, 0x72, 0x8b, 0xba, 0x84, 0xa8, 0x7b, 0xcd, 0x3a, 0xa6, 0x1d, 0x6a, 0x23, 0x5e,
	0xcc, 0x22, 0xbd, 0xe3, 0x58, 0xd2, 0x27, 0x24, 0xfc, 0x12, 0xb3, 0xcb, 0x07, 0xd1, 0xe1, 0xf4,
	0xe1, 0x6e, 0xda, 0x6b, 0xe9, 0xbc, 0x6d, 0xcf, 0xf7, 0x9a, 0x75, 0xbc, 0x13, 0xb8, 0xe8, 0x63,
	0x59, 0xf0, 0xb3, 0xbf, 0x11, 0xb9, 0xd1, 0x27, 0x74, 0x06, 0x56, 0x4e, 0xbd, 0xb7, 0xc2, 0x18,
	0x65, 0xe9, 0x2d, 0x32, 0xca, 0x41, 0x2a, 0xcc, 0x38, 0x9e, 0xef, 0x36, 0xeb, 0x78, 0x1a, 0x00,
	0x6d, 0x97, 0x65, 0x28, 0xd2, 0xa7, 0x64, 0xda, 0x3e, 0x5f, 0xfe, 0x30, 0xa5, 0xd0, 0xab, 0x2e,
	0xc4, 0xcd, 0x4d, 0x88, 0x81, 0xc8, 0xb2, 0xa1, 0x95, 0xde, 0x21, 0x63, 0x65, 0x2d, 0x58, 0xdc,
	0x78, 0x32, 0x5c, 0x10, 0xdb, 0x2c, 0x0b, 0x32, 0x7d, 0x44, 0x46, 0x52, 0x78, 0x31, 0x1b, 0x5d,
	0x1c, 0x6c, 0xb0, 0x57, 0x6b, 0x63, 0x19, 0xba, 0x97, 0x57, 0xf0, 0xfb, 0x1d, 0xfd, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x43, 0x9a, 0x34, 0x37, 0x03, 0x00, 0x00,
}
