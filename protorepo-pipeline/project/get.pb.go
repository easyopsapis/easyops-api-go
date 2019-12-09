// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package project

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
//Get请求
type GetRequest struct {
	//
	//代码项目id
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id" form:"project_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

//
//Get返回
type GetResponse struct {
	//
	//代码项目
	Project *pipeline.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project" form:"project"`
	//
	//project关联的provider信息
	Provider *pipeline.Provider `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider" form:"provider"`
	//
	//project关联的pipeline信息
	Pipeline             []*pipeline.Pipeline `protobuf:"bytes,3,rep,name=pipeline,proto3" json:"pipeline" form:"pipeline"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetProject() *pipeline.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *GetResponse) GetProvider() *pipeline.Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *GetResponse) GetPipeline() []*pipeline.Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

//
//GetApi返回
type GetResponseWrapper struct {
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
	Data                 *GetResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponseWrapper) Reset()         { *m = GetResponseWrapper{} }
func (m *GetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetResponseWrapper) ProtoMessage()    {}
func (*GetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseWrapper.Unmarshal(m, b)
}
func (m *GetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseWrapper.Merge(m, src)
}
func (m *GetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetResponseWrapper.Size(m)
}
func (m *GetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseWrapper proto.InternalMessageInfo

func (m *GetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetResponseWrapper) GetData() *GetResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "project.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "project.GetResponse")
	proto.RegisterType((*GetResponseWrapper)(nil), "project.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x55, 0xda, 0x05, 0xea, 0xa2, 0xfd, 0x30, 0x08, 0x45, 0x7b, 0xc9, 0xca, 0x20, 0x04,
	0x87, 0xc4, 0xc0, 0x4a, 0x88, 0xe5, 0x82, 0x14, 0x84, 0x80, 0xdb, 0xe2, 0x03, 0x1c, 0x10, 0x20,
	0x37, 0xf1, 0x06, 0x43, 0xda, 0x31, 0xb6, 0xbb, 0xcb, 0x87, 0x78, 0x4d, 0xc4, 0x29, 0x48, 0x3c,
	0x42, 0x9e, 0x00, 0x75, 0xe2, 0x74, 0xd3, 0x3d, 0xf7, 0x94, 0x99, 0xf9, 0xcf, 0xff, 0x67, 0xcf,
	0xc4, 0x64, 0x5c, 0x2a, 0x9f, 0x1a, 0x0b, 0x1e, 0xe8, 0x15, 0x63, 0xe1, 0xb3, 0xca, 0xfd, 0x7e,
	0x52, 0x6a, 0xff, 0x69, 0x31, 0x4d, 0x73, 0x98, 0xf1, 0x12, 0x4a, 0xe0, 0xa8, 0x4f, 0x17, 0x27,
	0x98, 0x61, 0x82, 0x51, 0xeb, 0xdb, 0x3f, 0x2e, 0x21, 0x55, 0xd2, 0x7d, 0x07, 0xe3, 0xd2, 0x0a,
	0x72, 0x59, 0xf1, 0x1c, 0xe6, 0xde, 0xca, 0xdc, 0xbb, 0xd6, 0x69, 0x95, 0x81, 0x64, 0x06, 0x85,
	0xaa, 0x1c, 0x0f, 0x8d, 0x1c, 0x53, 0x6e, 0xb4, 0x51, 0x95, 0x9e, 0x2b, 0x1e, 0x8e, 0x0e, 0xc4,
	0xd7, 0x9b, 0x21, 0x9e, 0xea, 0x42, 0xd9, 0x4d, 0x22, 0x43, 0x10, 0x90, 0x8f, 0x7a, 0x6b, 0x9a,
	0x9d, 0x69, 0xff, 0x05, 0xce, 0x78, 0x09, 0x09, 0x8a, 0xc9, 0xa9, 0xac, 0x74, 0x21, 0x3d, 0x58,
	0xc7, 0x57, 0x61, 0xeb, 0x63, 0x6f, 0x08, 0x79, 0xa1, 0xbc, 0x50, 0x5f, 0x17, 0xca, 0x79, 0xfa,
	0x92, 0x90, 0x30, 0xfc, 0x47, 0x5d, 0x44, 0x83, 0x83, 0xc1, 0xdd, 0x71, 0x76, 0xaf, 0xa9, 0xe3,
	0xbd, 0x13, 0xb0, 0xb3, 0x27, 0xec, 0x5c, 0x63, 0xff, 0xfe, 0xc6, 0xbb, 0x64, 0xfb, 0xc3, 0xbb,
	0xfb, 0xc9, 0x91, 0x4c, 0x7e, 0xbc, 0xff, 0xf9, 0xe0, 0xf0, 0xd7, 0x6d, 0x31, 0x0e, 0x0d, 0xaf,
	0x0a, 0xf6, 0x67, 0x40, 0x26, 0x08, 0x76, 0x06, 0xe6, 0x4e, 0xd1, 0xa7, 0xa4, 0xfb, 0xa3, 0x88,
	0x9d, 0x3c, 0xdc, 0x4b, 0x57, 0x13, 0x1c, 0xb7, 0x42, 0x46, 0x9b, 0x3a, 0xde, 0x5e, 0x3b, 0x89,
	0x89, 0xce, 0x45, 0x9f, 0x91, 0xab, 0xdd, 0x16, 0xa3, 0x4b, 0x48, 0xa0, 0x6b, 0x04, 0x54, 0xb2,
	0xeb, 0x4d, 0x1d, 0xef, 0xac, 0x10, 0x58, 0x63, 0x62, 0x65, 0x44, 0x48, 0xf0, 0x44, 0xc3, 0x83,
	0xe1, 0x05, 0x48, 0x08, 0xd6, 0x20, 0xa1, 0xb6, 0x84, 0x74, 0xe1, 0xef, 0x01, 0xa1, 0xbd, 0xd1,
	0xde, 0x5a, 0x69, 0x8c, 0xb2, 0xf4, 0x16, 0x19, 0xe5, 0x50, 0x28, 0x1c, 0x6f, 0x2b, 0xdb, 0x69,
	0xea, 0x78, 0xd2, 0x32, 0x96, 0x55, 0x26, 0x50, 0xa4, 0x8f, 0xc9, 0x64, 0xf9, 0x7d, 0xfe, 0xcd,
	0x54, 0x52, 0xcf, 0x71, 0x90, 0x71, 0x76, 0xb3, 0xa9, 0x63, 0x7a, 0xde, 0x1b, 0x44, 0x26, 0xfa,
	0xad, 0xf4, 0x0e, 0xd9, 0x52, 0xd6, 0x82, 0x8d, 0x86, 0xe8, 0xd9, 0x6d, 0xea, 0xf8, 0x5a, 0xeb,
	0xc1, 0x32, 0x13, 0xad, 0x4c, 0x8f, 0xc8, 0xa8, 0x90, 0x5e, 0x46, 0x23, 0xdc, 0xd1, 0x8d, 0xb4,
	0x7b, 0xcc, 0xbd, 0x1b, 0xf7, 0x2f, 0xb7, 0xec, 0x65, 0x02, 0x2d, 0xd3, 0xcb, 0xf8, 0x24, 0x0e,
	0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xdd, 0x0d, 0x73, 0x87, 0x03, 0x00, 0x00,
}
