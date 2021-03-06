// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package pipeline

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
//DeletePipeline请求
type DeletePipelineRequest struct {
	//
	//关联的project id
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id" form:"project_id"`
	//
	//要删除的pipeline id
	PipelineId           string   `protobuf:"bytes,2,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id" form:"pipeline_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePipelineRequest) Reset()         { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()    {}
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeletePipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePipelineRequest.Unmarshal(m, b)
}
func (m *DeletePipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePipelineRequest.Marshal(b, m, deterministic)
}
func (m *DeletePipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePipelineRequest.Merge(m, src)
}
func (m *DeletePipelineRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePipelineRequest.Size(m)
}
func (m *DeletePipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePipelineRequest proto.InternalMessageInfo

func (m *DeletePipelineRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *DeletePipelineRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

//
//DeletePipelineApi返回
type DeletePipelineResponseWrapper struct {
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

func (m *DeletePipelineResponseWrapper) Reset()         { *m = DeletePipelineResponseWrapper{} }
func (m *DeletePipelineResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineResponseWrapper) ProtoMessage()    {}
func (*DeletePipelineResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeletePipelineResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePipelineResponseWrapper.Unmarshal(m, b)
}
func (m *DeletePipelineResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePipelineResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeletePipelineResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePipelineResponseWrapper.Merge(m, src)
}
func (m *DeletePipelineResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeletePipelineResponseWrapper.Size(m)
}
func (m *DeletePipelineResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePipelineResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePipelineResponseWrapper proto.InternalMessageInfo

func (m *DeletePipelineResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeletePipelineResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeletePipelineResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeletePipelineResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeletePipelineRequest)(nil), "pipeline.DeletePipelineRequest")
	proto.RegisterType((*DeletePipelineResponseWrapper)(nil), "pipeline.DeletePipelineResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdb, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xa9, 0x6e, 0xe2, 0xd2, 0xa1, 0x33, 0xe0, 0x18, 0x13, 0xe9, 0x88, 0x22, 0x53, 0x68,
	0xab, 0x0e, 0xc4, 0xc3, 0xdd, 0x70, 0xe0, 0xf0, 0x46, 0x7a, 0xe3, 0x85, 0xa8, 0x74, 0x6b, 0x56,
	0xab, 0xed, 0xbe, 0x98, 0x66, 0xce, 0x03, 0xbe, 0x92, 0x8f, 0x54, 0x41, 0xdf, 0xa0, 0x4f, 0x20,
	0x4d, 0x3b, 0x5b, 0xc4, 0xab, 0x26, 0xdf, 0xff, 0xd0, 0x5f, 0xf8, 0x50, 0xd5, 0xa1, 0x3e, 0x15,
	0xd4, 0x60, 0x1c, 0x04, 0xe0, 0x45, 0xe6, 0x31, 0xea, 0x7b, 0x63, 0xda, 0xd4, 0x5d, 0x4f, 0xdc,
	0x4d, 0x06, 0xc6, 0x10, 0x02, 0xd3, 0x05, 0x17, 0x4c, 0x69, 0x18, 0x4c, 0x46, 0xf2, 0x26, 0x2f,
	0xf2, 0x94, 0x06, 0x9b, 0x07, 0x05, 0x7b, 0x30, 0xf5, 0xc4, 0x03, 0x4c, 0x4d, 0x17, 0x74, 0x29,
	0xea, 0x4f, 0xb6, 0xef, 0x39, 0xb6, 0x00, 0x1e, 0x9a, 0xbf, 0xc7, 0x2c, 0xb7, 0xe6, 0x02, 0xb8,
	0x3e, 0xcd, 0xdb, 0x69, 0xc0, 0xc4, 0x4b, 0x2a, 0x92, 0x0f, 0x05, 0xad, 0x9e, 0x4a, 0xbc, 0x8b,
	0x0c, 0xcb, 0xa2, 0x8f, 0x13, 0x1a, 0x0a, 0x7c, 0x86, 0x10, 0xe3, 0x70, 0x4f, 0x87, 0xe2, 0xd6,
	0x73, 0x1a, 0x4a, 0x4b, 0x69, 0x57, 0xba, 0xdb, 0x71, 0xa4, 0xad, 0x8c, 0x80, 0x07, 0xc7, 0x24,
	0xd7, 0xc8, 0xd7, 0xa7, 0x56, 0x43, 0x4b, 0x37, 0x57, 0xbb, 0xfa, 0x91, 0xad, 0xbf, 0x5e, 0xbf,
	0xed, 0x75, 0xde, 0x37, 0xad, 0x4a, 0x66, 0xe8, 0x3b, 0xf8, 0x1c, 0xa9, 0xb3, 0x37, 0x27, 0x55,
	0x73, 0xb2, 0x6a, 0x27, 0x8e, 0x34, 0x9c, 0x55, 0xe5, 0xe2, 0xff, 0x5d, 0x68, 0xe6, 0xe8, 0x3b,
	0xe4, 0x5b, 0x41, 0xeb, 0x7f, 0x81, 0x43, 0x06, 0xe3, 0x90, 0x5e, 0x72, 0x9b, 0x31, 0xca, 0xf1,
	0x06, 0x2a, 0x0d, 0xc1, 0xa1, 0x12, 0xb9, 0xdc, 0x5d, 0x8e, 0x23, 0x4d, 0x4d, 0xff, 0x93, 0x4c,
	0x89, 0x25, 0x45, 0x7c, 0x88, 0xd4, 0xe4, 0xdb, 0x7b, 0x66, 0xbe, 0xed, 0x8d, 0x33, 0xa6, 0x7a,
	0xce, 0x54, 0x10, 0x89, 0x55, 0xb4, 0xe2, 0x2d, 0x54, 0xa6, 0x9c, 0x03, 0x6f, 0xcc, 0xcb, 0x4c,
	0x2d, 0x8e, 0xb4, 0x6a, 0x9a, 0x91, 0x63, 0x62, 0xa5, 0x32, 0x3e, 0x41, 0x25, 0xc7, 0x16, 0x76,
	0xa3, 0xd4, 0x52, 0xda, 0xea, 0x7e, 0xdd, 0x48, 0xb7, 0x60, 0xcc, 0xb6, 0x60, 0xf4, 0x92, 0x2d,
	0x14, 0xf1, 0x12, 0x37, 0xb1, 0x64, 0x68, 0xb0, 0x20, 0x6d, 0x9d, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x0e, 0xce, 0xd2, 0x3b, 0x02, 0x00, 0x00,
}
