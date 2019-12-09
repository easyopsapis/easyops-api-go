// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_task_detail.proto

package task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//DelTask请求
type DelTaskRequest struct {
	//
	//任务id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//任务名称
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelTaskRequest) Reset()         { *m = DelTaskRequest{} }
func (m *DelTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DelTaskRequest) ProtoMessage()    {}
func (*DelTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_23af3e098a717073, []int{0}
}
func (m *DelTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelTaskRequest.Unmarshal(m, b)
}
func (m *DelTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelTaskRequest.Marshal(b, m, deterministic)
}
func (m *DelTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelTaskRequest.Merge(m, src)
}
func (m *DelTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DelTaskRequest.Size(m)
}
func (m *DelTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelTaskRequest proto.InternalMessageInfo

func (m *DelTaskRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DelTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//
//DelTaskApi返回
type DelTaskResponseWrapper struct {
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

func (m *DelTaskResponseWrapper) Reset()         { *m = DelTaskResponseWrapper{} }
func (m *DelTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DelTaskResponseWrapper) ProtoMessage()    {}
func (*DelTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_23af3e098a717073, []int{1}
}
func (m *DelTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelTaskResponseWrapper.Unmarshal(m, b)
}
func (m *DelTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DelTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelTaskResponseWrapper.Merge(m, src)
}
func (m *DelTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DelTaskResponseWrapper.Size(m)
}
func (m *DelTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DelTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DelTaskResponseWrapper proto.InternalMessageInfo

func (m *DelTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DelTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DelTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DelTaskResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DelTaskRequest)(nil), "task.DelTaskRequest")
	proto.RegisterType((*DelTaskResponseWrapper)(nil), "task.DelTaskResponseWrapper")
}

func init() { proto.RegisterFile("delete_task_detail.proto", fileDescriptor_23af3e098a717073) }

var fileDescriptor_23af3e098a717073 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xc9, 0xf7, 0xa5, 0x05, 0x27, 0xfe, 0x9d, 0x45, 0x08, 0x75, 0x91, 0x32, 0x82, 0xd4,
	0x85, 0x29, 0xe8, 0x46, 0x74, 0x27, 0x76, 0xe1, 0x76, 0x10, 0x5c, 0xd6, 0x49, 0x73, 0x1b, 0x43,
	0x93, 0xcc, 0x38, 0x99, 0x80, 0xbe, 0xab, 0xe4, 0x21, 0xf2, 0x04, 0x32, 0x77, 0xa2, 0xcd, 0x2a,
	0xb9, 0xe7, 0xfc, 0x0e, 0x67, 0x38, 0x24, 0xca, 0xa0, 0x04, 0x03, 0x6b, 0x23, 0x9a, 0xdd, 0x3a,
	0x03, 0x23, 0x8a, 0x32, 0x51, 0x5a, 0x1a, 0x49, 0x7d, 0x2b, 0xcd, 0xae, 0xf3, 0xc2, 0xbc, 0xb7,
	0x69, 0xb2, 0x91, 0xd5, 0x32, 0x97, 0xb9, 0x5c, 0xa2, 0x99, 0xb6, 0x5b, 0xbc, 0xf0, 0xc0, 0x3f,
	0x17, 0x9a, 0x9d, 0xe7, 0x52, 0xe6, 0x25, 0xec, 0x29, 0xa8, 0x94, 0xf9, 0x72, 0x26, 0x7b, 0x23,
	0xc7, 0x4f, 0x50, 0xbe, 0x88, 0x66, 0xc7, 0xe1, 0xa3, 0x85, 0xc6, 0xd0, 0x2b, 0x32, 0xb5, 0x2d,
	0xcf, 0x59, 0xe4, 0xcd, 0xbd, 0xc5, 0xc1, 0xe3, 0x59, 0xdf, 0xc5, 0x47, 0x5b, 0xa9, 0xab, 0x7b,
	0xe6, 0x74, 0xc6, 0x07, 0x80, 0x5e, 0x10, 0xbf, 0x16, 0x15, 0x44, 0xff, 0x10, 0x3c, 0xe9, 0xbb,
	0x38, 0x70, 0xa0, 0x55, 0x19, 0x47, 0x93, 0x7d, 0x7b, 0x24, 0xfc, 0xab, 0x68, 0x94, 0xac, 0x1b,
	0x78, 0xd5, 0x42, 0x29, 0xd0, 0x36, 0xbf, 0x91, 0x19, 0x60, 0xd1, 0x64, 0x9c, 0xb7, 0x2a, 0xe3,
	0x68, 0xd2, 0x3b, 0x12, 0xd8, 0xef, 0xea, 0x53, 0x95, 0xa2, 0xa8, 0x87, 0xae, 0xb0, 0xef, 0x62,
	0xba, 0x67, 0x07, 0x93, 0xf1, 0x31, 0x4a, 0x2f, 0xc9, 0x04, 0xb4, 0x96, 0x3a, 0xfa, 0x8f, 0x99,
	0xd3, 0xbe, 0x8b, 0x0f, 0x5d, 0x06, 0x65, 0xc6, 0x9d, 0x4d, 0x1f, 0x88, 0x9f, 0x09, 0x23, 0x22,
	0x7f, 0xee, 0x2d, 0x82, 0x9b, 0x30, 0x71, 0x7b, 0x25, 0xbf, 0x7b, 0x25, 0x2b, 0xbb, 0xd7, 0xf8,
	0x79, 0x96, 0x66, 0x1c, 0x43, 0xe9, 0x14, 0xb1, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xe0, 0xc4, 0xdf, 0xb5, 0x01, 0x00, 0x00,
}