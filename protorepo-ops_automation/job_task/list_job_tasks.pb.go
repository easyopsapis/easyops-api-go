// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_job_tasks.proto

package job_task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	ops_automation "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/ops_automation"
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
//ListJobTasks请求
type ListJobTasksRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页大小
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//job id
	JobId string `protobuf:"bytes,3,opt,name=jobId,proto3" json:"jobId" form:"jobId"`
	//
	//执行触发器
	Trigger string `protobuf:"bytes,4,opt,name=trigger,proto3" json:"trigger" form:"trigger"`
	//
	//easy command id
	TaskId string `protobuf:"bytes,5,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//工具类型
	TaskType string `protobuf:"bytes,6,opt,name=taskType,proto3" json:"taskType" form:"taskType"`
	//
	//工具执行用户
	ExecUser string `protobuf:"bytes,7,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//工具结果状态
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobTasksRequest) Reset()         { *m = ListJobTasksRequest{} }
func (m *ListJobTasksRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobTasksRequest) ProtoMessage()    {}
func (*ListJobTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc584d8dcea5461, []int{0}
}
func (m *ListJobTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobTasksRequest.Unmarshal(m, b)
}
func (m *ListJobTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobTasksRequest.Marshal(b, m, deterministic)
}
func (m *ListJobTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobTasksRequest.Merge(m, src)
}
func (m *ListJobTasksRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobTasksRequest.Size(m)
}
func (m *ListJobTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobTasksRequest proto.InternalMessageInfo

func (m *ListJobTasksRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListJobTasksRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListJobTasksRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *ListJobTasksRequest) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *ListJobTasksRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ListJobTasksRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *ListJobTasksRequest) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *ListJobTasksRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//ListJobTasks返回
type ListJobTasksResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*ops_automation.JobTasks `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListJobTasksResponse) Reset()         { *m = ListJobTasksResponse{} }
func (m *ListJobTasksResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobTasksResponse) ProtoMessage()    {}
func (*ListJobTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc584d8dcea5461, []int{1}
}
func (m *ListJobTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobTasksResponse.Unmarshal(m, b)
}
func (m *ListJobTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobTasksResponse.Marshal(b, m, deterministic)
}
func (m *ListJobTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobTasksResponse.Merge(m, src)
}
func (m *ListJobTasksResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobTasksResponse.Size(m)
}
func (m *ListJobTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobTasksResponse proto.InternalMessageInfo

func (m *ListJobTasksResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListJobTasksResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListJobTasksResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListJobTasksResponse) GetList() []*ops_automation.JobTasks {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListJobTasksApi返回
type ListJobTasksResponseWrapper struct {
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
	Data                 *ListJobTasksResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListJobTasksResponseWrapper) Reset()         { *m = ListJobTasksResponseWrapper{} }
func (m *ListJobTasksResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListJobTasksResponseWrapper) ProtoMessage()    {}
func (*ListJobTasksResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cc584d8dcea5461, []int{2}
}
func (m *ListJobTasksResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobTasksResponseWrapper.Unmarshal(m, b)
}
func (m *ListJobTasksResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobTasksResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListJobTasksResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobTasksResponseWrapper.Merge(m, src)
}
func (m *ListJobTasksResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListJobTasksResponseWrapper.Size(m)
}
func (m *ListJobTasksResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobTasksResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobTasksResponseWrapper proto.InternalMessageInfo

func (m *ListJobTasksResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListJobTasksResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListJobTasksResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListJobTasksResponseWrapper) GetData() *ListJobTasksResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListJobTasksRequest)(nil), "job_task.ListJobTasksRequest")
	proto.RegisterType((*ListJobTasksResponse)(nil), "job_task.ListJobTasksResponse")
	proto.RegisterType((*ListJobTasksResponseWrapper)(nil), "job_task.ListJobTasksResponseWrapper")
}

func init() { proto.RegisterFile("list_job_tasks.proto", fileDescriptor_1cc584d8dcea5461) }

var fileDescriptor_1cc584d8dcea5461 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd7, 0xad, 0xed, 0x3a, 0x17, 0xd8, 0xf0, 0xc6, 0x08, 0x45, 0x22, 0x93, 0x41, 0x68,
	0x93, 0x96, 0x44, 0xda, 0x60, 0x20, 0x24, 0x0e, 0x14, 0x71, 0x18, 0xe2, 0x64, 0x86, 0x40, 0xfc,
	0xab, 0xdc, 0xd6, 0x0b, 0xd9, 0xd2, 0xbe, 0xc1, 0x76, 0xd9, 0x1f, 0xc4, 0xc7, 0xe1, 0x6b, 0x05,
	0xc1, 0x8d, 0x6b, 0x3e, 0x01, 0xf2, 0x9b, 0xa6, 0xe9, 0xa0, 0x17, 0x4e, 0xb6, 0xdf, 0xe7, 0xf7,
	0x38, 0x8f, 0xf3, 0xda, 0x64, 0x2d, 0x8e, 0xb4, 0xe9, 0x1c, 0x41, 0xb7, 0x63, 0x84, 0x3e, 0xd6,
	0x7e, 0xa2, 0xc0, 0x00, 0x6d, 0x14, 0x85, 0x96, 0x17, 0x46, 0xe6, 0xd3, 0xa8, 0xeb, 0xf7, 0x60,
	0x10, 0x84, 0x10, 0x42, 0x80, 0x40, 0x77, 0x74, 0x88, 0x2b, 0x5c, 0xe0, 0x2c, 0x37, 0xb6, 0xde,
	0x84, 0xe0, 0x4b, 0xa1, 0xcf, 0x20, 0xd1, 0x7e, 0x0c, 0x3d, 0x11, 0x07, 0x3d, 0x18, 0x1a, 0x25,
	0x7a, 0x46, 0xe7, 0x4e, 0x25, 0x13, 0xf0, 0x06, 0xd0, 0x97, 0xb1, 0x0e, 0xc6, 0x60, 0x80, 0xcb,
	0x00, 0x12, 0xdd, 0x11, 0x23, 0x03, 0x03, 0x61, 0x22, 0x18, 0x06, 0x7f, 0x45, 0x6a, 0xed, 0x4d,
	0x05, 0x19, 0x9c, 0x44, 0xe6, 0x18, 0x4e, 0x82, 0x10, 0x3c, 0x14, 0xbd, 0x2f, 0x22, 0x8e, 0xfa,
	0xc2, 0x80, 0xd2, 0xc1, 0x64, 0x9a, 0xfb, 0xd8, 0xf7, 0x05, 0xb2, 0xfa, 0x22, 0xd2, 0xe6, 0x39,
	0x74, 0x0f, 0xec, 0x76, 0x5c, 0x7e, 0x1e, 0x49, 0x6d, 0xe8, 0x16, 0xa9, 0x26, 0x22, 0x94, 0x4e,
	0x65, 0xa3, 0xb2, 0x59, 0x6b, 0x5f, 0xcb, 0x52, 0xb7, 0x79, 0x08, 0x6a, 0xf0, 0x88, 0xd9, 0x2a,
	0xfb, 0xf5, 0xc3, 0x9d, 0x5f, 0x99, 0xe3, 0x88, 0xd0, 0xfb, 0xa4, 0x61, 0xc7, 0x97, 0xd1, 0xb9,
	0x74, 0xe6, 0x11, 0xbf, 0x91, 0xa5, 0xee, 0x72, 0x89, 0x5b, 0xa5, 0xb0, 0x4c, 0x50, 0x7a, 0x97,
	0xd4, 0x8e, 0xa0, 0xbb, 0xdf, 0x77, 0x16, 0x36, 0x2a, 0x9b, 0x4b, 0xed, 0x95, 0x2c, 0x75, 0x2f,
	0xe5, 0x1e, 0x2c, 0x33, 0x9e, 0xcb, 0x74, 0x9b, 0x2c, 0x1a, 0x15, 0x85, 0xa1, 0x54, 0x4e, 0x15,
	0x49, 0x9a, 0xa5, 0xee, 0x95, 0x9c, 0x1c, 0x0b, 0x8c, 0x17, 0x08, 0x0d, 0x48, 0xdd, 0xfe, 0x96,
	0xfd, 0xbe, 0x53, 0x43, 0xf8, 0x7a, 0x96, 0xba, 0x97, 0xc7, 0x30, 0xd6, 0x31, 0x48, 0x32, 0xc7,
	0xc7, 0x18, 0x0d, 0x48, 0xc3, 0xce, 0x0e, 0xce, 0x12, 0xe9, 0xd4, 0xd1, 0xb2, 0x5a, 0xa6, 0x2f,
	0x14, 0xc6, 0x27, 0x10, 0xdd, 0x27, 0x0d, 0x79, 0x2a, 0x7b, 0xaf, 0xb4, 0x54, 0xce, 0x22, 0x1a,
	0xbc, 0xd2, 0x50, 0x28, 0xf6, 0x2b, 0xeb, 0x64, 0xed, 0xe3, 0xbb, 0x27, 0xde, 0x5b, 0xe1, 0x9d,
	0x77, 0xbc, 0x0f, 0xef, 0x4f, 0xbe, 0xee, 0x6e, 0xef, 0xdd, 0xfb, 0x76, 0x87, 0x4f, 0xec, 0x74,
	0x8b, 0xd4, 0xb5, 0x11, 0x66, 0xa4, 0x9d, 0x06, 0x6e, 0x74, 0xb5, 0x0c, 0x9b, 0xd7, 0x19, 0x1f,
	0x03, 0xec, 0x67, 0x85, 0xac, 0x5d, 0xec, 0x93, 0x4e, 0x60, 0xa8, 0xe5, 0xff, 0x34, 0xea, 0x01,
	0x59, 0xb2, 0x63, 0x47, 0x97, 0x9d, 0x6a, 0x65, 0xa9, 0xbb, 0x52, 0xf2, 0x28, 0xcd, 0x6c, 0x95,
	0x01, 0x23, 0x62, 0x6c, 0x55, 0x6d, 0xba, 0x55, 0x58, 0x66, 0x3c, 0x97, 0xe9, 0x63, 0x52, 0xb5,
	0xef, 0xc5, 0xa9, 0x6e, 0x2c, 0x6c, 0x36, 0x77, 0x1c, 0xff, 0xe2, 0x9d, 0xf5, 0x8b, 0xec, 0xed,
	0xe5, 0x32, 0xa5, 0xe5, 0x19, 0x47, 0x1b, 0xfb, 0x5d, 0x21, 0x37, 0x67, 0x9d, 0xf1, 0xb5, 0x12,
	0x49, 0x22, 0x15, 0xbd, 0x4d, 0xaa, 0x3d, 0xe8, 0x17, 0x47, 0x9d, 0xda, 0xc4, 0x56, 0x19, 0x47,
	0x91, 0x3e, 0x24, 0x4d, 0x3b, 0x3e, 0x3b, 0x4d, 0x62, 0x11, 0x0d, 0xf1, 0x98, 0x4b, 0xed, 0xf5,
	0x2c, 0x75, 0x69, 0xc9, 0x8e, 0x45, 0xc6, 0xa7, 0x51, 0x7b, 0x4a, 0xa9, 0x14, 0xa8, 0x7f, 0x2f,
	0x24, 0x96, 0x19, 0xcf, 0x65, 0xfa, 0x94, 0x54, 0xfb, 0xc2, 0x08, 0xbc, 0x8d, 0xcd, 0x9d, 0x5b,
	0x7e, 0xf1, 0x14, 0xfd, 0x59, 0xd9, 0xa7, 0x63, 0x5a, 0x17, 0xe3, 0x68, 0xee, 0xd6, 0xf1, 0xf9,
	0xed, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xff, 0x2e, 0x04, 0xf8, 0x61, 0x04, 0x00, 0x00,
}