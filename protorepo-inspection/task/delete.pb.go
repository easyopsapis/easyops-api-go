// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package task

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
//DeleteTask请求
type DeleteTaskRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//巡检任务ID
	InspectionTaskId     string   `protobuf:"bytes,2,opt,name=inspectionTaskId,proto3" json:"inspectionTaskId" form:"inspectionTaskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *DeleteTaskRequest) GetInspectionTaskId() string {
	if m != nil {
		return m.InspectionTaskId
	}
	return ""
}

//
//DeleteTaskApi返回
type DeleteTaskResponseWrapper struct {
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

func (m *DeleteTaskResponseWrapper) Reset()         { *m = DeleteTaskResponseWrapper{} }
func (m *DeleteTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponseWrapper) ProtoMessage()    {}
func (*DeleteTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponseWrapper.Merge(m, src)
}
func (m *DeleteTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponseWrapper.Size(m)
}
func (m *DeleteTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponseWrapper proto.InternalMessageInfo

func (m *DeleteTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteTaskResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteTaskRequest)(nil), "task.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponseWrapper)(nil), "task.DeleteTaskResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdd, 0x6a, 0xdb, 0x30,
	0x1c, 0xc5, 0xf1, 0xe6, 0x8c, 0x4d, 0x09, 0x24, 0xd3, 0x45, 0xe6, 0x25, 0x0c, 0x07, 0x6d, 0x8c,
	0x5c, 0xcc, 0x76, 0xd6, 0x40, 0xe8, 0xc7, 0x55, 0x43, 0x43, 0xc9, 0xad, 0x28, 0x14, 0x1a, 0xda,
	0xa2, 0xc4, 0x8a, 0x6b, 0x62, 0x5b, 0xaa, 0x2c, 0x37, 0xfd, 0xa0, 0xef, 0xd4, 0x27, 0x72, 0xa0,
	0x8f, 0xe0, 0x27, 0x28, 0x96, 0xf3, 0x61, 0xda, 0x2b, 0xff, 0xf5, 0xff, 0x9d, 0x73, 0x7c, 0xe0,
	0x0f, 0x6a, 0x2e, 0x0d, 0xa8, 0xa4, 0x36, 0x17, 0x4c, 0x32, 0xa8, 0x4b, 0x12, 0x2f, 0x5a, 0x96,
	0xe7, 0xcb, 0x9b, 0x64, 0x6a, 0xcf, 0x58, 0xe8, 0x78, 0xcc, 0x63, 0x8e, 0x82, 0xd3, 0x64, 0xae,
	0x5e, 0xea, 0xa1, 0xa6, 0xc2, 0xd4, 0x1a, 0x94, 0xe4, 0xe1, 0xd2, 0x97, 0x0b, 0xb6, 0x74, 0x3c,
	0x66, 0x29, 0x68, 0xdd, 0x91, 0xc0, 0x77, 0x89, 0x64, 0x22, 0x76, 0xb6, 0xe3, 0xda, 0xd7, 0xf6,
	0x18, 0xf3, 0x02, 0xba, 0x4b, 0xa7, 0x21, 0x97, 0x0f, 0x05, 0x44, 0x2f, 0x1a, 0xf8, 0x7e, 0xa2,
	0xaa, 0x9d, 0x91, 0x78, 0x81, 0xe9, 0x6d, 0x42, 0x63, 0x09, 0x31, 0xf8, 0xca, 0x83, 0xc4, 0xf3,
	0xa3, 0xb1, 0x6b, 0x68, 0x1d, 0xad, 0xfb, 0x6d, 0x38, 0xc8, 0x52, 0xb3, 0x3e, 0x67, 0x22, 0x3c,
	0x44, 0x1b, 0x82, 0x5e, 0x57, 0xa6, 0x09, 0x7e, 0x5d, 0x4d, 0x88, 0xf5, 0x78, 0x6c, 0x5d, 0x5c,
	0x5f, 0x4e, 0x7a, 0xd6, 0xc1, 0x66, 0x7e, 0xea, 0xfd, 0xeb, 0xff, 0x7f, 0xfe, 0x83, 0xb7, 0x39,
	0xf0, 0x14, 0x34, 0xfc, 0x28, 0xe6, 0x74, 0x26, 0x7d, 0x16, 0xe5, 0x3f, 0x1b, 0xbb, 0xc6, 0x27,
	0x95, 0xdd, 0xce, 0x52, 0xf3, 0x47, 0x91, 0xfd, 0x5e, 0x81, 0xf0, 0x07, 0x13, 0x5a, 0x69, 0xe0,
	0x67, 0xb9, 0x72, 0xcc, 0x59, 0x14, 0xd3, 0x73, 0x41, 0x38, 0xa7, 0x02, 0xfe, 0x06, 0xfa, 0x8c,
	0xb9, 0x54, 0xd5, 0xae, 0x0c, 0xeb, 0x59, 0x6a, 0x56, 0x8b, 0xe8, 0x7c, 0x8b, 0xb0, 0x82, 0x70,
	0x1f, 0x54, 0xf3, 0xef, 0xe8, 0x9e, 0x07, 0xc4, 0x8f, 0xd6, 0x35, 0x9a, 0x59, 0x6a, 0xc2, 0x9d,
	0x76, 0x0d, 0x11, 0x2e, 0x4b, 0xe1, 0x5f, 0x50, 0xa1, 0x42, 0x30, 0x61, 0x7c, 0x56, 0x9e, 0x46,
	0x96, 0x9a, 0xb5, 0xc2, 0xa3, 0xd6, 0x08, 0x17, 0x18, 0x1e, 0x01, 0xdd, 0x25, 0x92, 0x18, 0x7a,
	0x47, 0xeb, 0x56, 0xf7, 0x9a, 0x76, 0x71, 0x03, 0x7b, 0x73, 0x03, 0x7b, 0x94, 0xdf, 0xa0, 0x5c,
	0x2f, 0x57, 0x23, 0xac, 0x4c, 0xd3, 0x2f, 0x4a, 0xd6, 0x7f, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xad, 0x48, 0xeb, 0x35, 0x02, 0x00, 0x00,
}
