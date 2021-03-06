// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_service_task.proto

package service_manage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//CreateServiceTask请求
type CreateServiceTaskRequest struct {
	//
	//任务名称（vnode：虚拟节点采集，topology：服务拓扑发现）
	Task string `protobuf:"bytes,1,opt,name=task,proto3" json:"task" form:"task"`
	//
	//执行命令（check：检查任务执行状态， execute：立即执行任务）
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command" form:"command"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceTaskRequest) Reset()         { *m = CreateServiceTaskRequest{} }
func (m *CreateServiceTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceTaskRequest) ProtoMessage()    {}
func (*CreateServiceTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36766c5bb8738500, []int{0}
}
func (m *CreateServiceTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceTaskRequest.Unmarshal(m, b)
}
func (m *CreateServiceTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceTaskRequest.Merge(m, src)
}
func (m *CreateServiceTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceTaskRequest.Size(m)
}
func (m *CreateServiceTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceTaskRequest proto.InternalMessageInfo

func (m *CreateServiceTaskRequest) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *CreateServiceTaskRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

//
//CreateServiceTask返回
type CreateServiceTaskResponse struct {
	//
	//任务名称
	Task string `protobuf:"bytes,1,opt,name=task,proto3" json:"task" form:"task"`
	//
	//任务运行状态
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceTaskResponse) Reset()         { *m = CreateServiceTaskResponse{} }
func (m *CreateServiceTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceTaskResponse) ProtoMessage()    {}
func (*CreateServiceTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36766c5bb8738500, []int{1}
}
func (m *CreateServiceTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceTaskResponse.Unmarshal(m, b)
}
func (m *CreateServiceTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceTaskResponse.Merge(m, src)
}
func (m *CreateServiceTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceTaskResponse.Size(m)
}
func (m *CreateServiceTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceTaskResponse proto.InternalMessageInfo

func (m *CreateServiceTaskResponse) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *CreateServiceTaskResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//CreateServiceTaskApi返回
type CreateServiceTaskResponseWrapper struct {
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
	Data                 *CreateServiceTaskResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CreateServiceTaskResponseWrapper) Reset()         { *m = CreateServiceTaskResponseWrapper{} }
func (m *CreateServiceTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateServiceTaskResponseWrapper) ProtoMessage()    {}
func (*CreateServiceTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_36766c5bb8738500, []int{2}
}
func (m *CreateServiceTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceTaskResponseWrapper.Unmarshal(m, b)
}
func (m *CreateServiceTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateServiceTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceTaskResponseWrapper.Merge(m, src)
}
func (m *CreateServiceTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateServiceTaskResponseWrapper.Size(m)
}
func (m *CreateServiceTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceTaskResponseWrapper proto.InternalMessageInfo

func (m *CreateServiceTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateServiceTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateServiceTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateServiceTaskResponseWrapper) GetData() *CreateServiceTaskResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceTaskRequest)(nil), "service_manage.CreateServiceTaskRequest")
	proto.RegisterType((*CreateServiceTaskResponse)(nil), "service_manage.CreateServiceTaskResponse")
	proto.RegisterType((*CreateServiceTaskResponseWrapper)(nil), "service_manage.CreateServiceTaskResponseWrapper")
}

func init() { proto.RegisterFile("create_service_task.proto", fileDescriptor_36766c5bb8738500) }

var fileDescriptor_36766c5bb8738500 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xe9, 0xfb, 0x6e, 0x13, 0x53, 0x9d, 0x9a, 0x83, 0x74, 0x5e, 0x5a, 0x22, 0xc8, 0x06,
	0xda, 0x81, 0x5e, 0xc4, 0xa3, 0xe2, 0xd5, 0x43, 0x14, 0x3c, 0x8e, 0xb4, 0xcd, 0x6a, 0xa9, 0x69,
	0x6a, 0x92, 0x8a, 0x9f, 0xb6, 0x1f, 0xa2, 0x37, 0x6f, 0x92, 0x27, 0x1d, 0x56, 0x70, 0xe0, 0xa9,
	0xcd, 0xff, 0xff, 0x7b, 0x9e, 0x5f, 0x4b, 0xd0, 0x2c, 0x55, 0x9c, 0x19, 0xbe, 0xd2, 0x5c, 0xbd,
	0x17, 0x29, 0x5f, 0x19, 0xa6, 0xcb, 0xb8, 0x56, 0xd2, 0x48, 0x3c, 0xdd, 0x64, 0x82, 0x55, 0x2c,
	0xe7, 0x27, 0x17, 0x79, 0x61, 0x5e, 0x9a, 0x24, 0x4e, 0xa5, 0x58, 0xe6, 0x32, 0x97, 0x4b, 0xc0,
	0x92, 0x66, 0x0d, 0x27, 0x38, 0xc0, 0x9b, 0x1b, 0x27, 0x02, 0x05, 0x77, 0xb0, 0xfb, 0xd1, 0xad,
	0x79, 0x62, 0xba, 0xa4, 0xfc, 0xad, 0xe1, 0xda, 0xe0, 0x53, 0x34, 0xb2, 0xa2, 0xc0, 0x8b, 0xbc,
	0xf9, 0xee, 0xed, 0x41, 0xd7, 0x86, 0xfe, 0x5a, 0x2a, 0x71, 0x43, 0x6c, 0x4a, 0x28, 0x94, 0xf8,
	0x1c, 0xed, 0xa4, 0x52, 0x08, 0x56, 0x65, 0xc1, 0x3f, 0xe0, 0x70, 0xd7, 0x86, 0x53, 0xc7, 0xf5,
	0x05, 0xa1, 0x1b, 0x84, 0x94, 0x68, 0xf6, 0x8b, 0x4e, 0xd7, 0xb2, 0xd2, 0xfc, 0x6f, 0xbe, 0x05,
	0x9a, 0x68, 0xc3, 0x4c, 0xa3, 0x7b, 0xdd, 0x51, 0xd7, 0x86, 0xfb, 0x0e, 0x73, 0x39, 0xa1, 0x3d,
	0x40, 0x3e, 0x3d, 0x14, 0x6d, 0xb5, 0x3d, 0x2b, 0x56, 0xd7, 0x5c, 0x59, 0x69, 0x2a, 0x33, 0x0e,
	0xd2, 0xf1, 0x50, 0x6a, 0x53, 0x42, 0xa1, 0xc4, 0xd7, 0xc8, 0xb7, 0xcf, 0xfb, 0x8f, 0xfa, 0x95,
	0x15, 0x55, 0x6f, 0x3e, 0xee, 0xda, 0x10, 0x7f, 0xb3, 0x7d, 0x49, 0xe8, 0x10, 0xc5, 0x67, 0x68,
	0xcc, 0x95, 0x92, 0x2a, 0xf8, 0x0f, 0x33, 0x87, 0x5d, 0x1b, 0xee, 0xb9, 0x19, 0x88, 0x09, 0x75,
	0x35, 0x7e, 0x40, 0xa3, 0x8c, 0x19, 0x16, 0x8c, 0x22, 0x6f, 0xee, 0x5f, 0x2e, 0xe2, 0x9f, 0xb7,
	0x1a, 0x6f, 0xfd, 0x8d, 0xe1, 0x17, 0xdb, 0x05, 0x84, 0xc2, 0x9e, 0x64, 0x02, 0xd7, 0x7b, 0xf5,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x5d, 0xe6, 0x0a, 0x3a, 0x02, 0x00, 0x00,
}
