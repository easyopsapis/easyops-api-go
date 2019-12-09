// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: callback_task_v1.proto

package admin_task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	easy_command "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_command"
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
//CallbackAdminTaskV1请求
type CallbackAdminTaskV1Request struct {
	//
	//任务ID
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//任务执行状态
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status" form:"status"`
	//
	//最后一个出错的目标机器action的退出码
	Code int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code" form:"code"`
	//
	//任务的各个目标机器的执行结果详情
	TargetsLog []*easy_command.TargetLog `protobuf:"bytes,4,rep,name=targetsLog,proto3" json:"targetsLog" form:"targetsLog"`
	//
	//任务名称
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackAdminTaskV1Request) Reset()         { *m = CallbackAdminTaskV1Request{} }
func (m *CallbackAdminTaskV1Request) String() string { return proto.CompactTextString(m) }
func (*CallbackAdminTaskV1Request) ProtoMessage()    {}
func (*CallbackAdminTaskV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fbbebf2ae504d4, []int{0}
}
func (m *CallbackAdminTaskV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackAdminTaskV1Request.Unmarshal(m, b)
}
func (m *CallbackAdminTaskV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackAdminTaskV1Request.Marshal(b, m, deterministic)
}
func (m *CallbackAdminTaskV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackAdminTaskV1Request.Merge(m, src)
}
func (m *CallbackAdminTaskV1Request) XXX_Size() int {
	return xxx_messageInfo_CallbackAdminTaskV1Request.Size(m)
}
func (m *CallbackAdminTaskV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackAdminTaskV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackAdminTaskV1Request proto.InternalMessageInfo

func (m *CallbackAdminTaskV1Request) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *CallbackAdminTaskV1Request) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CallbackAdminTaskV1Request) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CallbackAdminTaskV1Request) GetTargetsLog() []*easy_command.TargetLog {
	if m != nil {
		return m.TargetsLog
	}
	return nil
}

func (m *CallbackAdminTaskV1Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//
//CallbackAdminTaskV1返回
type CallbackAdminTaskV1Response struct {
	//
	//返回接受状态
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackAdminTaskV1Response) Reset()         { *m = CallbackAdminTaskV1Response{} }
func (m *CallbackAdminTaskV1Response) String() string { return proto.CompactTextString(m) }
func (*CallbackAdminTaskV1Response) ProtoMessage()    {}
func (*CallbackAdminTaskV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fbbebf2ae504d4, []int{1}
}
func (m *CallbackAdminTaskV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackAdminTaskV1Response.Unmarshal(m, b)
}
func (m *CallbackAdminTaskV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackAdminTaskV1Response.Marshal(b, m, deterministic)
}
func (m *CallbackAdminTaskV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackAdminTaskV1Response.Merge(m, src)
}
func (m *CallbackAdminTaskV1Response) XXX_Size() int {
	return xxx_messageInfo_CallbackAdminTaskV1Response.Size(m)
}
func (m *CallbackAdminTaskV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackAdminTaskV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackAdminTaskV1Response proto.InternalMessageInfo

func (m *CallbackAdminTaskV1Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//CallbackAdminTaskV1Api返回
type CallbackAdminTaskV1ResponseWrapper struct {
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
	Data                 *CallbackAdminTaskV1Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CallbackAdminTaskV1ResponseWrapper) Reset()         { *m = CallbackAdminTaskV1ResponseWrapper{} }
func (m *CallbackAdminTaskV1ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CallbackAdminTaskV1ResponseWrapper) ProtoMessage()    {}
func (*CallbackAdminTaskV1ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fbbebf2ae504d4, []int{2}
}
func (m *CallbackAdminTaskV1ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper.Unmarshal(m, b)
}
func (m *CallbackAdminTaskV1ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CallbackAdminTaskV1ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper.Merge(m, src)
}
func (m *CallbackAdminTaskV1ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper.Size(m)
}
func (m *CallbackAdminTaskV1ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackAdminTaskV1ResponseWrapper proto.InternalMessageInfo

func (m *CallbackAdminTaskV1ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CallbackAdminTaskV1ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CallbackAdminTaskV1ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CallbackAdminTaskV1ResponseWrapper) GetData() *CallbackAdminTaskV1Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CallbackAdminTaskV1Request)(nil), "admin_task.CallbackAdminTaskV1Request")
	proto.RegisterType((*CallbackAdminTaskV1Response)(nil), "admin_task.CallbackAdminTaskV1Response")
	proto.RegisterType((*CallbackAdminTaskV1ResponseWrapper)(nil), "admin_task.CallbackAdminTaskV1ResponseWrapper")
}

func init() { proto.RegisterFile("callback_task_v1.proto", fileDescriptor_91fbbebf2ae504d4) }

var fileDescriptor_91fbbebf2ae504d4 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x25, 0xbb, 0x33, 0x0b, 0xdb, 0xa3, 0xe8, 0x06, 0x5c, 0xc3, 0x78, 0xc8, 0xd0, 0x82, 0x8e,
	0x87, 0xed, 0xb0, 0xeb, 0x45, 0xbc, 0x39, 0x22, 0x28, 0x0c, 0x1e, 0x9a, 0xc5, 0x3d, 0x0e, 0x95,
	0x4e, 0x6f, 0x3b, 0x4c, 0x92, 0x8a, 0xdd, 0x1d, 0xd1, 0x7f, 0xf0, 0x1b, 0xf3, 0x11, 0xb9, 0x79,
	0x93, 0xee, 0xce, 0x3a, 0x41, 0x64, 0xf6, 0x94, 0xd4, 0xab, 0xf7, 0xaa, 0x5f, 0x3d, 0x8a, 0x9c,
	0x0b, 0x28, 0xcb, 0x1c, 0xc4, 0x6e, 0x63, 0xc1, 0xec, 0x36, 0xdf, 0x2f, 0x59, 0xa3, 0xd1, 0x62,
	0x4c, 0xa0, 0xa8, 0xb6, 0xb5, 0x07, 0xe7, 0x17, 0x6a, 0x6b, 0xbf, 0xb6, 0x39, 0x13, 0x58, 0x65,
	0x0a, 0x15, 0x66, 0x9e, 0x92, 0xb7, 0xb7, 0xbe, 0xf2, 0x85, 0xff, 0x0b, 0xd2, 0xf9, 0x8d, 0x42,
	0x26, 0xc1, 0xfc, 0xc4, 0xc6, 0xb0, 0x12, 0x05, 0x94, 0x99, 0xc0, 0xda, 0x6a, 0x10, 0xd6, 0x04,
	0xa5, 0x96, 0x0d, 0x5e, 0x54, 0x58, 0xc8, 0xd2, 0x64, 0x03, 0x31, 0xf3, 0xa5, 0xaf, 0x36, 0x02,
	0xab, 0x0a, 0xea, 0x22, 0xb3, 0xa0, 0x95, 0xb4, 0x9b, 0x12, 0x55, 0x18, 0x4c, 0x7f, 0x1d, 0x91,
	0xf9, 0xfb, 0xc1, 0xee, 0x3b, 0x67, 0xef, 0x1a, 0xcc, 0xee, 0xcb, 0x25, 0x97, 0xdf, 0x5a, 0x69,
	0x6c, 0xfc, 0x8a, 0x9c, 0x38, 0xbb, 0x9f, 0x8a, 0x24, 0x5a, 0x44, 0xcb, 0xd3, 0xd5, 0x59, 0xdf,
	0xa5, 0x0f, 0x6f, 0x51, 0x57, 0x6f, 0x69, 0xc0, 0x29, 0x1f, 0x08, 0x8e, 0x6a, 0x2c, 0xd8, 0xd6,
	0x24, 0x47, 0xff, 0x52, 0x03, 0x4e, 0xf9, 0x40, 0x88, 0x9f, 0x93, 0x89, 0xc0, 0x42, 0x26, 0xc7,
	0x8b, 0x68, 0x39, 0x5d, 0x3d, 0xea, 0xbb, 0x74, 0x16, 0x88, 0x0e, 0xa5, 0xdc, 0x37, 0xe3, 0xcf,
	0x84, 0x04, 0xb7, 0x66, 0x8d, 0x2a, 0x99, 0x2c, 0x8e, 0x97, 0xb3, 0xab, 0xa7, 0x6c, 0xbc, 0x0d,
	0xbb, 0xf6, 0xfd, 0x35, 0xaa, 0xd5, 0x93, 0xbe, 0x4b, 0xcf, 0xee, 0x7c, 0xdd, 0x89, 0x28, 0x1f,
	0x4d, 0x70, 0x8f, 0xd6, 0x50, 0xc9, 0x64, 0xea, 0xdd, 0x8d, 0x1e, 0x75, 0x28, 0xe5, 0xbe, 0x49,
	0x3f, 0x92, 0x67, 0xff, 0x4d, 0xc3, 0x34, 0x58, 0x1b, 0x39, 0xda, 0x31, 0xba, 0x67, 0x47, 0xfa,
	0x3b, 0x22, 0xf4, 0xc0, 0xa8, 0x1b, 0x0d, 0x4d, 0x23, 0xf5, 0xdf, 0x28, 0xa2, 0x43, 0x51, 0xbc,
	0x21, 0x33, 0xf7, 0xfd, 0xf0, 0xa3, 0x29, 0x61, 0x5b, 0x0f, 0xf9, 0x9e, 0xf7, 0x5d, 0x1a, 0xef,
	0xb9, 0x43, 0x93, 0xf2, 0x31, 0x35, 0x7e, 0x41, 0xa6, 0x52, 0x6b, 0xd4, 0x3e, 0xea, 0xd3, 0xd5,
	0xe3, 0xbe, 0x4b, 0x1f, 0x04, 0x8d, 0x87, 0x29, 0x0f, 0xed, 0x78, 0x4d, 0x26, 0x05, 0x58, 0x48,
	0x26, 0x8b, 0x68, 0x39, 0xbb, 0x7a, 0xc9, 0xf6, 0x97, 0xca, 0x0e, 0x2c, 0x31, 0xf6, 0xeb, 0xe4,
	0x94, 0xfb, 0x29, 0xf9, 0x89, 0xbf, 0xad, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x2c,
	0x96, 0x1c, 0x09, 0x03, 0x00, 0x00,
}
