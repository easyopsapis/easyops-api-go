// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: history.proto

package scheduler

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//定时任务
type SchedulerHistory struct {
	//
	//history id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//任务名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//执行返回码
	Code int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code" form:"code"`
	//
	//调度结果信息
	Msg string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//任务执行总时间, 单位 s
	TotalTime int32 `protobuf:"varint,5,opt,name=total_time,json=totalTime,proto3" json:"total_time" form:"total_time"`
	//
	//总状态
	TotalStatus string `protobuf:"bytes,6,opt,name=total_status,json=totalStatus,proto3" json:"total_status" form:"total_status"`
	//
	//task id
	TaskId string `protobuf:"bytes,7,opt,name=task_id,json=taskId,proto3" json:"task_id" form:"task_id"`
	//
	//工具/流程/流水线/巡检 任务id
	ToolExecId string `protobuf:"bytes,8,opt,name=tool_exec_id,json=toolExecId,proto3" json:"tool_exec_id" form:"tool_exec_id"`
	//
	//启动时间
	StartTime string `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time" form:"start_time"`
	//
	//结束时间
	EndTime string `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3" json:"end_time" form:"end_time"`
	//
	//创建用户
	User string `protobuf:"bytes,11,opt,name=user,proto3" json:"user" form:"user"`
	//
	//org
	Org int32 `protobuf:"varint,12,opt,name=org,proto3" json:"org" form:"org"`
	//
	//发起定时任务的源 id
	SrcId string `protobuf:"bytes,13,opt,name=src_id,json=srcId,proto3" json:"src_id" form:"src_id"`
	//
	//任务对象类型
	JobType string `protobuf:"bytes,14,opt,name=job_type,json=jobType,proto3" json:"job_type" form:"job_type"`
	//
	//相关参数
	Annotations          *SchedulerHistory_Annotations `protobuf:"bytes,15,opt,name=annotations,proto3" json:"annotations" form:"annotations"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SchedulerHistory) Reset()         { *m = SchedulerHistory{} }
func (m *SchedulerHistory) String() string { return proto.CompactTextString(m) }
func (*SchedulerHistory) ProtoMessage()    {}
func (*SchedulerHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0}
}
func (m *SchedulerHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerHistory.Unmarshal(m, b)
}
func (m *SchedulerHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerHistory.Marshal(b, m, deterministic)
}
func (m *SchedulerHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerHistory.Merge(m, src)
}
func (m *SchedulerHistory) XXX_Size() int {
	return xxx_messageInfo_SchedulerHistory.Size(m)
}
func (m *SchedulerHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerHistory.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerHistory proto.InternalMessageInfo

func (m *SchedulerHistory) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SchedulerHistory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchedulerHistory) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SchedulerHistory) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SchedulerHistory) GetTotalTime() int32 {
	if m != nil {
		return m.TotalTime
	}
	return 0
}

func (m *SchedulerHistory) GetTotalStatus() string {
	if m != nil {
		return m.TotalStatus
	}
	return ""
}

func (m *SchedulerHistory) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SchedulerHistory) GetToolExecId() string {
	if m != nil {
		return m.ToolExecId
	}
	return ""
}

func (m *SchedulerHistory) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *SchedulerHistory) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *SchedulerHistory) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SchedulerHistory) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *SchedulerHistory) GetSrcId() string {
	if m != nil {
		return m.SrcId
	}
	return ""
}

func (m *SchedulerHistory) GetJobType() string {
	if m != nil {
		return m.JobType
	}
	return ""
}

func (m *SchedulerHistory) GetAnnotations() *SchedulerHistory_Annotations {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type SchedulerHistory_Annotations struct {
	//
	//应用ID
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulerHistory_Annotations) Reset()         { *m = SchedulerHistory_Annotations{} }
func (m *SchedulerHistory_Annotations) String() string { return proto.CompactTextString(m) }
func (*SchedulerHistory_Annotations) ProtoMessage()    {}
func (*SchedulerHistory_Annotations) Descriptor() ([]byte, []int) {
	return fileDescriptor_454388b49b309873, []int{0, 0}
}
func (m *SchedulerHistory_Annotations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerHistory_Annotations.Unmarshal(m, b)
}
func (m *SchedulerHistory_Annotations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerHistory_Annotations.Marshal(b, m, deterministic)
}
func (m *SchedulerHistory_Annotations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerHistory_Annotations.Merge(m, src)
}
func (m *SchedulerHistory_Annotations) XXX_Size() int {
	return xxx_messageInfo_SchedulerHistory_Annotations.Size(m)
}
func (m *SchedulerHistory_Annotations) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerHistory_Annotations.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerHistory_Annotations proto.InternalMessageInfo

func (m *SchedulerHistory_Annotations) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func init() {
	proto.RegisterType((*SchedulerHistory)(nil), "scheduler.SchedulerHistory")
	proto.RegisterType((*SchedulerHistory_Annotations)(nil), "scheduler.SchedulerHistory.Annotations")
}

func init() { proto.RegisterFile("history.proto", fileDescriptor_454388b49b309873) }

var fileDescriptor_454388b49b309873 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcb, 0x6e, 0xe3, 0x36,
	0x14, 0x1d, 0xe5, 0x61, 0xc7, 0x74, 0x9e, 0x4a, 0x1f, 0x42, 0x80, 0x42, 0xae, 0xc6, 0x98, 0xca,
	0x9d, 0x52, 0xf2, 0x6b, 0x06, 0x18, 0x6f, 0x8c, 0xb8, 0x98, 0xa2, 0xe9, 0x52, 0x93, 0x55, 0x54,
	0xc5, 0xa0, 0x25, 0x46, 0xd1, 0x8c, 0x65, 0x0a, 0x22, 0xdd, 0x19, 0x37, 0xce, 0x0f, 0x15, 0xe8,
	0x67, 0x14, 0xfd, 0x0b, 0x17, 0xed, 0xa6, 0xab, 0x6e, 0xfc, 0x05, 0x05, 0x2f, 0x6d, 0xcb, 0x28,
	0xda, 0x0f, 0xe8, 0x4a, 0xbc, 0xf7, 0x9c, 0x73, 0x2f, 0x8f, 0x78, 0x49, 0x74, 0x74, 0x9f, 0x70,
	0xc1, 0xf2, 0x99, 0x93, 0xe5, 0x4c, 0x30, 0xbd, 0xc2, 0xc3, 0x7b, 0x1a, 0x4d, 0xc7, 0x34, 0xbf,
	0xc0, 0x71, 0x22, 0xee, 0xa7, 0x23, 0x27, 0x64, 0xa9, 0x1b, 0xb3, 0x98, 0xb9, 0xc0, 0x18, 0x4d,
	0xef, 0x20, 0x82, 0x00, 0x56, 0x4a, 0x79, 0xf1, 0x72, 0x8b, 0x9e, 0xbe, 0x4f, 0xc4, 0x3b, 0xf6,
	0xde, 0x8d, 0x19, 0x06, 0x10, 0xff, 0x40, 0xc6, 0x49, 0x44, 0x04, 0xcb, 0xb9, 0xbb, 0x59, 0x2a,
	0x9d, 0xf5, 0x57, 0x05, 0x9d, 0xbe, 0x59, 0x37, 0xfd, 0x56, 0x6d, 0x46, 0xef, 0xa2, 0x9d, 0x24,
	0x32, 0xb4, 0x9a, 0x66, 0x57, 0x06, 0xf5, 0xe5, 0xc2, 0xac, 0xdc, 0xb1, 0x3c, 0xed, 0x59, 0x49,
	0x64, 0xfd, 0xf1, 0x9b, 0x79, 0x8e, 0xce, 0x6e, 0xfd, 0x26, 0x7e, 0x45, 0xf0, 0xdd, 0x25, 0xfe,
	0x26, 0x78, 0x68, 0x77, 0x1f, 0xeb, 0xde, 0x4e, 0x12, 0xe9, 0x0d, 0xb4, 0x37, 0x21, 0x29, 0x35,
	0x76, 0x40, 0xf7, 0xf1, 0x72, 0x61, 0x56, 0x95, 0x4e, 0x66, 0xa5, 0x72, 0x27, 0x7b, 0xe2, 0x01,
	0x45, 0x7f, 0x8a, 0xf6, 0x42, 0x16, 0x51, 0x63, 0xb7, 0xa6, 0xd9, 0xfb, 0x83, 0x93, 0x82, 0x2a,
	0xb3, 0x96, 0x07, 0xa0, 0x5e, 0x43, 0xbb, 0x29, 0x8f, 0x8d, 0x3d, 0x28, 0x77, 0xbc, 0x5c, 0x98,
	0x48, 0x71, 0x52, 0x1e, 0x5b, 0x9e, 0x84, 0xf4, 0x2e, 0x42, 0x82, 0x09, 0x32, 0x1e, 0x8a, 0x24,
	0xa5, 0xc6, 0x3e, 0x14, 0x93, 0x7d, 0xcf, 0x14, 0xb1, 0xc0, 0x2c, 0xaf, 0x02, 0xc1, 0x75, 0x92,
	0x52, 0xbd, 0x87, 0x0e, 0x15, 0xc2, 0x05, 0x11, 0x53, 0x6e, 0x94, 0xa0, 0xc1, 0xa7, 0xcb, 0x85,
	0x79, 0xbe, 0xad, 0x53, 0xa8, 0xe5, 0x55, 0x21, 0x7c, 0x03, 0x91, 0x3e, 0x40, 0x65, 0x41, 0xf8,
	0xbb, 0x61, 0x12, 0x19, 0x65, 0x90, 0x35, 0x96, 0x0b, 0xf3, 0x78, 0x25, 0x53, 0xc0, 0x7f, 0xfe,
	0xa3, 0x92, 0x24, 0x5c, 0x45, 0x7a, 0x5f, 0xf6, 0x67, 0xe3, 0x21, 0xfd, 0x40, 0x43, 0x59, 0xe8,
	0x00, 0x0a, 0x7d, 0xb6, 0xdd, 0xbf, 0x40, 0xd5, 0x7f, 0x3b, 0xf3, 0x90, 0x4c, 0xbe, 0xfe, 0x40,
	0xc3, 0xab, 0x48, 0xff, 0x53, 0x43, 0x88, 0x0b, 0x92, 0x0b, 0xe5, 0xbb, 0x02, 0xfa, 0x5f, 0xb5,
	0xc2, 0x78, 0x01, 0x4a, 0xf9, 0xcf, 0x1a, 0xfa, 0x49, 0xbb, 0xb5, 0xed, 0x7e, 0xcf, 0x6f, 0xe1,
	0x57, 0x81, 0xdc, 0x57, 0xf0, 0x65, 0xa3, 0x0f, 0xdf, 0x87, 0xee, 0x63, 0x03, 0xdb, 0x2d, 0xbf,
	0x89, 0xdb, 0xc1, 0xbc, 0x09, 0x78, 0x03, 0xdb, 0x1d, 0xbf, 0x89, 0x5b, 0xeb, 0x78, 0xee, 0xb7,
	0x70, 0x5b, 0xa9, 0x1a, 0xfe, 0x75, 0x2d, 0xb0, 0xdb, 0x7e, 0x13, 0x77, 0x82, 0x39, 0x70, 0x54,
	0xba, 0x67, 0xfb, 0x4d, 0xfc, 0x62, 0x1d, 0x14, 0x6b, 0xfb, 0x7b, 0x07, 0xbe, 0xcf, 0x1b, 0x7d,
	0xfb, 0x66, 0xee, 0x3f, 0xc7, 0x81, 0xdd, 0xef, 0xfd, 0x8b, 0x7c, 0x4b, 0xdd, 0xaf, 0x7b, 0x15,
	0xd8, 0x3e, 0x1c, 0xd5, 0xef, 0x1a, 0x3a, 0xa0, 0x93, 0x48, 0xf9, 0x44, 0xe0, 0xf3, 0x17, 0xe9,
	0xf3, 0x44, 0xf9, 0x5c, 0x43, 0xff, 0x47, 0x97, 0x65, 0x3a, 0x89, 0xc0, 0xe3, 0x77, 0x68, 0x6f,
	0xca, 0x69, 0x6e, 0x54, 0xc1, 0xde, 0xcb, 0xe2, 0x2e, 0xc8, 0xac, 0x74, 0xf6, 0x14, 0x7d, 0x7e,
	0xeb, 0x13, 0xfc, 0xe3, 0x25, 0xbe, 0x91, 0x62, 0xdf, 0xd9, 0xac, 0x87, 0x38, 0x78, 0x68, 0x7f,
	0xd5, 0x69, 0x3d, 0xd6, 0x3d, 0xa8, 0xa1, 0x3f, 0x43, 0xbb, 0x2c, 0x8f, 0x8d, 0x43, 0xb8, 0x09,
	0x1f, 0x15, 0x57, 0x86, 0xe5, 0x31, 0x0c, 0xd2, 0xe9, 0x13, 0x4f, 0x12, 0x74, 0x07, 0x95, 0x78,
	0x0e, 0xc3, 0x77, 0xb4, 0x19, 0xfe, 0xa3, 0xd5, 0xec, 0xe4, 0xc5, 0xd8, 0x69, 0xde, 0x3e, 0xcf,
	0xe5, 0xc4, 0x39, 0xe8, 0xe0, 0x2d, 0x1b, 0x0d, 0xc5, 0x2c, 0xa3, 0xc6, 0x31, 0x28, 0xce, 0x8b,
	0x53, 0x58, 0x23, 0x96, 0x57, 0x7e, 0xcb, 0x46, 0xd7, 0xb3, 0x8c, 0xea, 0x04, 0x55, 0xc9, 0x64,
	0xc2, 0x04, 0x11, 0x09, 0x9b, 0x70, 0xe3, 0xa4, 0xa6, 0xd9, 0xd5, 0xf6, 0x17, 0xce, 0xe6, 0x75,
	0x73, 0xfe, 0xf9, 0xe4, 0x38, 0x97, 0x05, 0x7d, 0xf0, 0xc9, 0x72, 0x61, 0xea, 0xaa, 0xf6, 0x56,
	0x15, 0xcb, 0xdb, 0xae, 0x79, 0xf1, 0x02, 0x55, 0xb7, 0x34, 0xfa, 0x33, 0xb4, 0x4f, 0xb2, 0xec,
	0x6a, 0xfd, 0x6a, 0x9d, 0x2e, 0x17, 0xe6, 0xe1, 0xaa, 0x84, 0x4c, 0x5b, 0x9e, 0x82, 0x07, 0xaf,
	0x6f, 0xbe, 0x8e, 0x99, 0x43, 0x09, 0x9f, 0xb1, 0x8c, 0x3b, 0x63, 0x16, 0x92, 0xb1, 0x1b, 0xb2,
	0x89, 0xc8, 0x49, 0x28, 0xb8, 0x7a, 0x63, 0x73, 0x9a, 0x31, 0x9c, 0xb2, 0x88, 0x8e, 0xb9, 0xbb,
	0x22, 0xba, 0x10, 0xba, 0x9b, 0xfd, 0x8f, 0x4a, 0xc0, 0xec, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xb9, 0xbe, 0x11, 0xc0, 0x05, 0x00, 0x00,
}
