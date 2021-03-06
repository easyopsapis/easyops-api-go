// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package project

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
//DeleteProject请求
type DeleteProjectRequest struct {
	//
	//关联的project id
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id" form:"project_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectRequest) Reset()         { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectRequest.Unmarshal(m, b)
}
func (m *DeleteProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectRequest.Merge(m, src)
}
func (m *DeleteProjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectRequest.Size(m)
}
func (m *DeleteProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectRequest proto.InternalMessageInfo

func (m *DeleteProjectRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

//
//DeleteProjectApi返回
type DeleteProjectResponseWrapper struct {
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

func (m *DeleteProjectResponseWrapper) Reset()         { *m = DeleteProjectResponseWrapper{} }
func (m *DeleteProjectResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectResponseWrapper) ProtoMessage()    {}
func (*DeleteProjectResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteProjectResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteProjectResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteProjectResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectResponseWrapper.Merge(m, src)
}
func (m *DeleteProjectResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectResponseWrapper.Size(m)
}
func (m *DeleteProjectResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectResponseWrapper proto.InternalMessageInfo

func (m *DeleteProjectResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteProjectResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteProjectResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteProjectResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteProjectRequest)(nil), "project.DeleteProjectRequest")
	proto.RegisterType((*DeleteProjectResponseWrapper)(nil), "project.DeleteProjectResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x5d, 0x4b, 0xfb, 0x30,
	0x14, 0xc6, 0xe9, 0xff, 0xbf, 0x29, 0xcb, 0x86, 0xce, 0x20, 0xa3, 0x4c, 0xa1, 0x23, 0x8a, 0xcc,
	0x8b, 0xb6, 0xea, 0x40, 0x7c, 0xb9, 0x1b, 0x0e, 0xf4, 0x4e, 0x7a, 0xe3, 0x85, 0xf8, 0x92, 0xad,
	0x59, 0xad, 0xb6, 0x3b, 0x31, 0xcd, 0x9c, 0x2f, 0xf8, 0x55, 0x2b, 0xec, 0x23, 0xf4, 0x13, 0x48,
	0x93, 0xea, 0x8a, 0x57, 0x39, 0xe7, 0xfc, 0x9e, 0x27, 0xe7, 0xe1, 0xa0, 0x86, 0xcf, 0x22, 0x26,
	0x99, 0xc3, 0x05, 0x48, 0xc0, 0xcb, 0x5c, 0xc0, 0x23, 0x1b, 0xc9, 0xb6, 0x1d, 0x84, 0xf2, 0x61,
	0x3a, 0x74, 0x46, 0x10, 0xbb, 0x01, 0x04, 0xe0, 0x2a, 0x3e, 0x9c, 0x8e, 0x55, 0xa7, 0x1a, 0x55,
	0x69, 0x5f, 0xfb, 0xb0, 0x24, 0x8f, 0x67, 0xa1, 0x7c, 0x82, 0x99, 0x1b, 0x80, 0xad, 0xa0, 0xfd,
	0x42, 0xa3, 0xd0, 0xa7, 0x12, 0x44, 0xe2, 0xfe, 0x96, 0x85, 0x6f, 0x23, 0x00, 0x08, 0x22, 0xb6,
	0xf8, 0x9d, 0xc5, 0x5c, 0xbe, 0x69, 0x48, 0xee, 0xd1, 0xfa, 0x99, 0x0a, 0x77, 0xa9, 0x43, 0x79,
	0xec, 0x79, 0xca, 0x12, 0x89, 0xcf, 0x11, 0x2a, 0x62, 0xde, 0x85, 0xbe, 0x69, 0x74, 0x8c, 0x6e,
	0xad, 0xbf, 0x9b, 0xa5, 0xd6, 0xda, 0x18, 0x44, 0x7c, 0x42, 0x16, 0x8c, 0xcc, 0xbf, 0xac, 0x26,
	0x5a, 0xb9, 0xbd, 0xde, 0xb3, 0x8f, 0xa9, 0xfd, 0x7e, 0xf3, 0xb1, 0xdf, 0xfb, 0xdc, 0xf6, 0x6a,
	0x85, 0xe0, 0xc2, 0x27, 0x73, 0x03, 0x6d, 0xfe, 0x59, 0x91, 0x70, 0x98, 0x24, 0xec, 0x4a, 0x50,
	0xce, 0x99, 0xc0, 0x5b, 0xa8, 0x32, 0x02, 0x9f, 0xa9, 0x25, 0xd5, 0xfe, 0x6a, 0x96, 0x5a, 0x75,
	0xbd, 0x24, 0x9f, 0x12, 0x4f, 0x41, 0x7c, 0x84, 0xea, 0xf9, 0x3b, 0x78, 0xe5, 0x11, 0x0d, 0x27,
	0xe6, 0x3f, 0x15, 0xa8, 0x95, 0xa5, 0x16, 0x5e, 0x68, 0x0b, 0x48, 0xbc, 0xb2, 0x14, 0xef, 0xa0,
	0x2a, 0x13, 0x02, 0x84, 0xf9, 0x5f, 0x79, 0x9a, 0x59, 0x6a, 0x35, 0xb4, 0x47, 0x8d, 0x89, 0xa7,
	0x31, 0x3e, 0x45, 0x15, 0x9f, 0x4a, 0x6a, 0x56, 0x3a, 0x46, 0xb7, 0x7e, 0xd0, 0x72, 0xf4, 0xd5,
	0x9c, 0x9f, 0xab, 0x39, 0x83, 0xfc, 0x6a, 0xe5, 0x78, 0xb9, 0x9a, 0x78, 0xca, 0x34, 0x5c, 0x52,
	0xb2, 0xde, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x94, 0x6a, 0x09, 0xea, 0x01, 0x00, 0x00,
}
