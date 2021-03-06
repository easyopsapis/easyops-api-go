// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_task.proto

package admin_task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	agent_admin "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/agent_admin"
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
//GetTask请求
type GetTaskRequest struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92ba9aa4fe3ed1c, []int{0}
}
func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetTaskApi返回
type GetTaskResponseWrapper struct {
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
	Data                 *agent_admin.AdminTask `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetTaskResponseWrapper) Reset()         { *m = GetTaskResponseWrapper{} }
func (m *GetTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponseWrapper) ProtoMessage()    {}
func (*GetTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92ba9aa4fe3ed1c, []int{1}
}
func (m *GetTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponseWrapper.Unmarshal(m, b)
}
func (m *GetTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponseWrapper.Merge(m, src)
}
func (m *GetTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponseWrapper.Size(m)
}
func (m *GetTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponseWrapper proto.InternalMessageInfo

func (m *GetTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetTaskResponseWrapper) GetData() *agent_admin.AdminTask {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTaskRequest)(nil), "admin_task.GetTaskRequest")
	proto.RegisterType((*GetTaskResponseWrapper)(nil), "admin_task.GetTaskResponseWrapper")
}

func init() { proto.RegisterFile("get_task.proto", fileDescriptor_e92ba9aa4fe3ed1c) }

var fileDescriptor_e92ba9aa4fe3ed1c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xe9, 0xfe, 0xdb, 0x1f, 0x96, 0xe9, 0x94, 0x5c, 0x8c, 0x32, 0x90, 0x8e, 0x08, 0xb2,
	0x9b, 0x35, 0xa0, 0x37, 0xa2, 0x57, 0x0e, 0xc4, 0xfb, 0x22, 0x7a, 0x39, 0xb2, 0xe6, 0xac, 0x96,
	0xb5, 0x3d, 0x31, 0xc9, 0x40, 0xdf, 0x55, 0xfa, 0x10, 0x7d, 0x02, 0xe9, 0xe9, 0x70, 0xbd, 0x49,
	0x72, 0xce, 0xf7, 0xfd, 0x72, 0xbe, 0x84, 0x4d, 0x33, 0xf0, 0x1b, 0xaf, 0xdc, 0x3e, 0x36, 0x16,
	0x3d, 0x72, 0xa6, 0x74, 0x99, 0x57, 0xd4, 0x99, 0xaf, 0xb2, 0xdc, 0x7f, 0x1c, 0xb6, 0x71, 0x8a,
	0xa5, 0xcc, 0x30, 0x43, 0x49, 0x96, 0xed, 0x61, 0x47, 0x15, 0x15, 0x74, 0xea, 0xd0, 0xf9, 0x5b,
	0x86, 0x31, 0x28, 0xf7, 0x8d, 0xc6, 0xc5, 0x05, 0xa6, 0xaa, 0x90, 0x29, 0x56, 0xde, 0xaa, 0xd4,
	0xbb, 0x8e, 0xb4, 0x60, 0x70, 0x55, 0xa2, 0x86, 0xc2, 0xc9, 0xa3, 0x51, 0x52, 0x29, 0x55, 0x06,
	0x95, 0xdf, 0xd0, 0x68, 0x79, 0x0a, 0xd0, 0xdd, 0x2b, 0x24, 0x9b, 0xbe, 0x80, 0x7f, 0x55, 0x6e,
	0x9f, 0xc0, 0xe7, 0x01, 0x9c, 0xe7, 0x57, 0x6c, 0x90, 0xeb, 0x30, 0x58, 0x04, 0xcb, 0xf1, 0xfa,
	0xbc, 0xa9, 0xa3, 0xf1, 0x0e, 0x6d, 0xf9, 0x20, 0x72, 0x2d, 0x92, 0x41, 0xae, 0xc5, 0x4f, 0xc0,
	0x66, 0x7f, 0x84, 0x33, 0x58, 0x39, 0x78, 0xb7, 0xca, 0x18, 0xb0, 0xfc, 0x9a, 0x0d, 0x53, 0xd4,
	0x40, 0xec, 0x68, 0x7d, 0xd1, 0xd4, 0xd1, 0xa4, 0x63, 0xdb, 0xae, 0x48, 0x48, 0xe4, 0xf7, 0x6c,
	0xd2, 0xee, 0xcf, 0x5f, 0xa6, 0x50, 0x79, 0x15, 0x0e, 0x68, 0xce, 0xac, 0xa9, 0x23, 0x7e, 0xf2,
	0x1e, 0x45, 0x91, 0xf4, 0xad, 0xfc, 0x86, 0x8d, 0xc0, 0x5a, 0xb4, 0xe1, 0x3f, 0x62, 0x2e, 0x9b,
	0x3a, 0x3a, 0xeb, 0x18, 0x6a, 0x8b, 0xa4, 0x93, 0xf9, 0x23, 0x1b, 0x6a, 0xe5, 0x55, 0x38, 0x5c,
	0x04, 0xcb, 0xc9, 0xed, 0x2c, 0xee, 0xbd, 0x3f, 0x7e, 0x6a, 0xd7, 0x36, 0x7b, 0x3f, 0x5e, 0xeb,
	0x16, 0x09, 0x41, 0xdb, 0xff, 0xf4, 0x2d, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x2e,
	0x5d, 0xda, 0xbb, 0x01, 0x00, 0x00,
}
