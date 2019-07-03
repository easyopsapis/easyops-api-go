// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: running_tasks.proto

package desktop

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//RunningTasks返回
type RunningTasksResponse struct {
	//
	//appId
	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//任务Id
	TaskId string `protobuf:"bytes,2,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//任务类型， install-安装， uninstall-卸载
	TaskType             string   `protobuf:"bytes,3,opt,name=taskType,proto3" json:"taskType" form:"taskType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunningTasksResponse) Reset()         { *m = RunningTasksResponse{} }
func (m *RunningTasksResponse) String() string { return proto.CompactTextString(m) }
func (*RunningTasksResponse) ProtoMessage()    {}
func (*RunningTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd1664e6bebdef2d, []int{0}
}
func (m *RunningTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunningTasksResponse.Unmarshal(m, b)
}
func (m *RunningTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunningTasksResponse.Marshal(b, m, deterministic)
}
func (m *RunningTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunningTasksResponse.Merge(m, src)
}
func (m *RunningTasksResponse) XXX_Size() int {
	return xxx_messageInfo_RunningTasksResponse.Size(m)
}
func (m *RunningTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunningTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunningTasksResponse proto.InternalMessageInfo

func (m *RunningTasksResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RunningTasksResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RunningTasksResponse) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

//
//RunningTasksApi返回
type RunningTasksResponseWrapper struct {
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
	Data                 []*RunningTasksResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RunningTasksResponseWrapper) Reset()         { *m = RunningTasksResponseWrapper{} }
func (m *RunningTasksResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*RunningTasksResponseWrapper) ProtoMessage()    {}
func (*RunningTasksResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd1664e6bebdef2d, []int{1}
}
func (m *RunningTasksResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunningTasksResponseWrapper.Unmarshal(m, b)
}
func (m *RunningTasksResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunningTasksResponseWrapper.Marshal(b, m, deterministic)
}
func (m *RunningTasksResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunningTasksResponseWrapper.Merge(m, src)
}
func (m *RunningTasksResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_RunningTasksResponseWrapper.Size(m)
}
func (m *RunningTasksResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RunningTasksResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RunningTasksResponseWrapper proto.InternalMessageInfo

func (m *RunningTasksResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RunningTasksResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *RunningTasksResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RunningTasksResponseWrapper) GetData() []*RunningTasksResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RunningTasksResponse)(nil), "desktop.RunningTasksResponse")
	proto.RegisterType((*RunningTasksResponseWrapper)(nil), "desktop.RunningTasksResponseWrapper")
}

func init() { proto.RegisterFile("running_tasks.proto", fileDescriptor_dd1664e6bebdef2d) }

var fileDescriptor_dd1664e6bebdef2d = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0xfd, 0xa3, 0x6e, 0x94, 0xea, 0x56, 0x24, 0x28, 0x92, 0xb2, 0x82, 0xd4, 0x83,
	0x09, 0xe8, 0x45, 0x3c, 0x06, 0x3c, 0xf4, 0xba, 0x14, 0x3c, 0xca, 0xa6, 0xd9, 0xc6, 0x50, 0x9b,
	0x59, 0x36, 0x09, 0xe8, 0xab, 0xf8, 0x70, 0x39, 0x7b, 0xce, 0x13, 0xc8, 0xce, 0x2e, 0x36, 0x87,
	0x9e, 0x76, 0xe7, 0xfb, 0x7e, 0x33, 0x7c, 0xcc, 0x90, 0xa9, 0x6e, 0xca, 0xb2, 0x28, 0xf3, 0xf7,
	0x5a, 0x54, 0x9b, 0x2a, 0x52, 0x1a, 0x6a, 0xa0, 0x87, 0x99, 0xac, 0x36, 0x35, 0xa8, 0xab, 0x87,
	0xbc, 0xa8, 0x3f, 0x9a, 0x34, 0x5a, 0xc1, 0x36, 0xce, 0x21, 0x87, 0x18, 0xfd, 0xb4, 0x59, 0x63,
	0x85, 0x05, 0xfe, 0x6c, 0x1f, 0xfb, 0xf1, 0xc8, 0x05, 0xb7, 0xf3, 0x96, 0x66, 0x1c, 0x97, 0x95,
	0x82, 0xb2, 0x92, 0xf4, 0x8e, 0x8c, 0x84, 0x52, 0x8b, 0x2c, 0xf0, 0x66, 0xde, 0xfc, 0x38, 0x39,
	0xeb, 0xda, 0xf0, 0x64, 0x0d, 0x7a, 0xfb, 0xc2, 0x50, 0x66, 0xdc, 0xda, 0xf4, 0x9e, 0x8c, 0x4d,
	0x8e, 0x45, 0x16, 0x1c, 0x20, 0x78, 0xde, 0xb5, 0xe1, 0xa9, 0x05, 0xad, 0xce, 0xb8, 0x03, 0x68,
	0x4c, 0x8e, 0xcc, 0x6f, 0xf9, 0xad, 0x64, 0x30, 0x40, 0x78, 0xda, 0xb5, 0xe1, 0x64, 0x07, 0x1b,
	0x87, 0xf1, 0x7f, 0x88, 0xfd, 0x7a, 0xe4, 0x7a, 0x5f, 0xb8, 0x37, 0x2d, 0x94, 0x92, 0x9a, 0xde,
	0x92, 0xe1, 0x0a, 0x32, 0x89, 0x11, 0x47, 0xc9, 0xa4, 0x6b, 0x43, 0xdf, 0x0e, 0x33, 0x2a, 0xe3,
	0x68, 0xd2, 0x67, 0xe2, 0x9b, 0xf7, 0xf5, 0x4b, 0x7d, 0x8a, 0xa2, 0x74, 0x29, 0x2f, 0xbb, 0x36,
	0xa4, 0x3b, 0xd6, 0x99, 0x8c, 0xf7, 0x51, 0xb3, 0x02, 0xa9, 0x35, 0x68, 0x17, 0xb6, 0xb7, 0x02,
	0x94, 0x19, 0xb7, 0x36, 0x4d, 0xc8, 0x30, 0x13, 0xb5, 0x08, 0x86, 0xb3, 0xc1, 0xdc, 0x7f, 0xbc,
	0x89, 0xdc, 0x29, 0xa2, 0x7d, 0xd1, 0xfb, 0x29, 0x4d, 0x13, 0xe3, 0xd8, 0x9b, 0x8e, 0xf1, 0x1c,
	0x4f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x52, 0x78, 0x44, 0xdd, 0x01, 0x00, 0x00,
}