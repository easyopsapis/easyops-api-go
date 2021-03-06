// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_task.proto

package task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
//CreateTask返回
type CreateTaskResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//结果信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//返回数据
	Data                 *CreateTaskResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f57290545750a0b, []int{0}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateTaskResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CreateTaskResponse) GetData() *CreateTaskResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateTaskResponse_Data struct {
	//
	//新建立的task_id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse_Data) Reset()         { *m = CreateTaskResponse_Data{} }
func (m *CreateTaskResponse_Data) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse_Data) ProtoMessage()    {}
func (*CreateTaskResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f57290545750a0b, []int{0, 0}
}
func (m *CreateTaskResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse_Data.Unmarshal(m, b)
}
func (m *CreateTaskResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse_Data.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse_Data.Merge(m, src)
}
func (m *CreateTaskResponse_Data) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse_Data.Size(m)
}
func (m *CreateTaskResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse_Data proto.InternalMessageInfo

func (m *CreateTaskResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//CreateTaskApi返回
type CreateTaskResponseWrapper struct {
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
	Data                 *CreateTaskResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateTaskResponseWrapper) Reset()         { *m = CreateTaskResponseWrapper{} }
func (m *CreateTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponseWrapper) ProtoMessage()    {}
func (*CreateTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f57290545750a0b, []int{1}
}
func (m *CreateTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponseWrapper.Unmarshal(m, b)
}
func (m *CreateTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponseWrapper.Merge(m, src)
}
func (m *CreateTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponseWrapper.Size(m)
}
func (m *CreateTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponseWrapper proto.InternalMessageInfo

func (m *CreateTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateTaskResponseWrapper) GetData() *CreateTaskResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskResponse)(nil), "task.CreateTaskResponse")
	proto.RegisterType((*CreateTaskResponse_Data)(nil), "task.CreateTaskResponse.Data")
	proto.RegisterType((*CreateTaskResponseWrapper)(nil), "task.CreateTaskResponseWrapper")
}

func init() { proto.RegisterFile("create_task.proto", fileDescriptor_5f57290545750a0b) }

var fileDescriptor_5f57290545750a0b = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x4a, 0xf3, 0x30,
	0x18, 0xc6, 0xe9, 0xd6, 0x7d, 0xd0, 0xec, 0x43, 0x5d, 0x04, 0xa9, 0x43, 0x69, 0x89, 0x43, 0x76,
	0xd2, 0x56, 0xe6, 0x10, 0x15, 0x3d, 0xb0, 0xfe, 0xb9, 0x80, 0x22, 0x78, 0x20, 0x2a, 0x59, 0x9b,
	0xc5, 0xb2, 0x75, 0x29, 0x49, 0xe6, 0x04, 0xf1, 0x52, 0xad, 0xe8, 0x25, 0xf4, 0x0a, 0xa4, 0x69,
	0x71, 0x83, 0x31, 0xf0, 0xa8, 0x6f, 0xf2, 0x7b, 0x9e, 0xbe, 0xef, 0xfb, 0x04, 0xb4, 0x42, 0x4e,
	0xb0, 0x24, 0x4f, 0x12, 0x8b, 0x91, 0x9b, 0x72, 0x26, 0x19, 0xd4, 0x8b, 0xba, 0xed, 0xd0, 0x58,
	0x3e, 0x4f, 0x07, 0x6e, 0xc8, 0x12, 0x8f, 0x32, 0xca, 0x3c, 0x05, 0x07, 0xd3, 0xa1, 0x3a, 0xa9,
	0x83, 0xaa, 0x4a, 0x53, 0xfb, 0x68, 0x41, 0x9e, 0xcc, 0x62, 0x39, 0x62, 0x33, 0x8f, 0x32, 0x47,
	0x41, 0xe7, 0x05, 0x8f, 0xe3, 0x08, 0x4b, 0xc6, 0x85, 0xf7, 0x5b, 0x56, 0xbe, 0x1d, 0xca, 0x18,
	0x1d, 0x93, 0xf9, 0xdf, 0x85, 0xe4, 0xd3, 0x50, 0x96, 0x14, 0x7d, 0x68, 0x00, 0x5e, 0xaa, 0x01,
	0x6f, 0xb1, 0x18, 0x05, 0x44, 0xa4, 0x6c, 0x22, 0x08, 0xdc, 0x03, 0x7a, 0xc8, 0x22, 0x62, 0x6a,
	0xb6, 0xd6, 0x6d, 0xf8, 0xeb, 0x79, 0x66, 0x35, 0x87, 0x8c, 0x27, 0xa7, 0xa8, 0xb8, 0x45, 0x81,
	0x82, 0xd0, 0x06, 0xf5, 0x44, 0x50, 0xb3, 0x66, 0x6b, 0x5d, 0xc3, 0x5f, 0xcb, 0x33, 0x0b, 0x94,
	0x9a, 0x44, 0x50, 0x14, 0x14, 0x08, 0xfa, 0x40, 0x8f, 0xb0, 0xc4, 0x66, 0xdd, 0xd6, 0xba, 0xcd,
	0xde, 0xae, 0xab, 0x32, 0x58, 0x6e, 0xe7, 0x5e, 0x61, 0x89, 0x17, 0xbb, 0x14, 0x26, 0x14, 0x28,
	0x6f, 0xfb, 0x0c, 0xe8, 0x05, 0x86, 0x7d, 0x50, 0x8b, 0x23, 0x35, 0x90, 0xe1, 0x77, 0xf2, 0xcc,
	0x32, 0x4a, 0x69, 0x1c, 0xa1, 0xef, 0x4f, 0x6b, 0x13, 0xb4, 0x1e, 0xef, 0x0f, 0x9c, 0x13, 0xec,
	0x0c, 0x2f, 0x9c, 0x9b, 0x87, 0xb7, 0x5e, 0xff, 0xbd, 0x13, 0xd4, 0xe2, 0x08, 0x7d, 0x69, 0x60,
	0x7b, 0xb9, 0xe1, 0x1d, 0xc7, 0x69, 0x4a, 0xf8, 0xdf, 0xd6, 0x3c, 0x06, 0xcd, 0xe2, 0x7b, 0xfd,
	0x9a, 0x8e, 0x71, 0x3c, 0xa9, 0xd6, 0xdd, 0xca, 0x33, 0x0b, 0xce, 0xb5, 0x15, 0x44, 0xc1, 0xa2,
	0x14, 0xee, 0x83, 0x06, 0xe1, 0x9c, 0x71, 0xb5, 0xbf, 0xe1, 0x6f, 0xe4, 0x99, 0xf5, 0xbf, 0xf4,
	0xa8, 0x6b, 0x14, 0x94, 0x18, 0x9e, 0x57, 0x31, 0xe9, 0x2a, 0x26, 0x73, 0x55, 0x4c, 0x2b, 0x12,
	0x1a, 0xfc, 0x53, 0x4f, 0x79, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xbc, 0x9f, 0x61, 0x6a,
	0x02, 0x00, 0x00,
}
