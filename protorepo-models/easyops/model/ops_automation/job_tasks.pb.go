// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: job_tasks.proto

package ops_automation

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
//任务详情
type JobTasks struct {
	//
	//jobTask id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//job id
	JobId string `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId" form:"jobId"`
	//
	//任务名称
	JobName string `protobuf:"bytes,3,opt,name=jobName,proto3" json:"jobName" form:"jobName"`
	//
	//菜单名称
	MenuName string `protobuf:"bytes,4,opt,name=menuName,proto3" json:"menuName" form:"menuName"`
	//
	//exec id/ task id
	ExecId string `protobuf:"bytes,5,opt,name=execId,proto3" json:"execId" form:"execId"`
	//
	//工具类型
	ResourceType string `protobuf:"bytes,6,opt,name=resourceType,proto3" json:"resourceType" form:"resourceType"`
	//
	//tool/flow id
	ResourceId string `protobuf:"bytes,7,opt,name=resourceId,proto3" json:"resourceId" form:"resourceId"`
	//
	//工具版本ID
	ResourceVId string `protobuf:"bytes,8,opt,name=resourceVId,proto3" json:"resourceVId" form:"resourceVId"`
	//
	//工具版本名称
	ResourceVName string `protobuf:"bytes,9,opt,name=resourceVName,proto3" json:"resourceVName" form:"resourceVName"`
	//
	//执行触发器
	Trigger string `protobuf:"bytes,10,opt,name=trigger,proto3" json:"trigger" form:"trigger"`
	//
	//工具执行用户
	ExecUser string `protobuf:"bytes,11,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//执行主机列表
	Hosts []string `protobuf:"bytes,12,rep,name=hosts,proto3" json:"hosts" form:"hosts"`
	//
	//工具结果状态。成功、失败、执行中、初始态等待执行、人工确认、暂停、终止
	Status string `protobuf:"bytes,13,opt,name=status,proto3" json:"status" form:"status"`
	//
	//通知邮件
	Mail *MailInfo `protobuf:"bytes,14,opt,name=mail,proto3" json:"mail" form:"mail"`
	//
	//成功率。success agents / all agents
	SuccessRate float32 `protobuf:"fixed32,15,opt,name=successRate,proto3" json:"successRate" form:"successRate"`
	//
	//错误返回
	Error string `protobuf:"bytes,16,opt,name=error,proto3" json:"error" form:"error"`
	//
	//作业创建时间 2019-07-02 11:38:59
	CreateTime string `protobuf:"bytes,17,opt,name=createTime,proto3" json:"createTime" form:"createTime"`
	//
	//作业创建时间 2019-07-02 11:38:59
	UpdateTime string `protobuf:"bytes,18,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//作业创建者
	Creator string `protobuf:"bytes,19,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//org
	Org                  int32    `protobuf:"varint,20,opt,name=org,proto3" json:"org" form:"org"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobTasks) Reset()         { *m = JobTasks{} }
func (m *JobTasks) String() string { return proto.CompactTextString(m) }
func (*JobTasks) ProtoMessage()    {}
func (*JobTasks) Descriptor() ([]byte, []int) {
	return fileDescriptor_578db8647d352164, []int{0}
}
func (m *JobTasks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobTasks.Unmarshal(m, b)
}
func (m *JobTasks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobTasks.Marshal(b, m, deterministic)
}
func (m *JobTasks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobTasks.Merge(m, src)
}
func (m *JobTasks) XXX_Size() int {
	return xxx_messageInfo_JobTasks.Size(m)
}
func (m *JobTasks) XXX_DiscardUnknown() {
	xxx_messageInfo_JobTasks.DiscardUnknown(m)
}

var xxx_messageInfo_JobTasks proto.InternalMessageInfo

func (m *JobTasks) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JobTasks) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *JobTasks) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *JobTasks) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *JobTasks) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

func (m *JobTasks) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *JobTasks) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *JobTasks) GetResourceVId() string {
	if m != nil {
		return m.ResourceVId
	}
	return ""
}

func (m *JobTasks) GetResourceVName() string {
	if m != nil {
		return m.ResourceVName
	}
	return ""
}

func (m *JobTasks) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *JobTasks) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *JobTasks) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *JobTasks) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *JobTasks) GetMail() *MailInfo {
	if m != nil {
		return m.Mail
	}
	return nil
}

func (m *JobTasks) GetSuccessRate() float32 {
	if m != nil {
		return m.SuccessRate
	}
	return 0
}

func (m *JobTasks) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *JobTasks) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *JobTasks) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *JobTasks) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *JobTasks) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func init() {
	proto.RegisterType((*JobTasks)(nil), "ops_automation.JobTasks")
}

func init() { proto.RegisterFile("job_tasks.proto", fileDescriptor_578db8647d352164) }

var fileDescriptor_578db8647d352164 = []byte{
	// 1032 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0x6e, 0xd2, 0xed, 0x9f, 0x4c, 0xff, 0x4f, 0xfb, 0xfb, 0x61, 0xad, 0x90, 0x1c, 0x4c, 0x59,
	0x8d, 0xdb, 0x8e, 0x93, 0x38, 0x49, 0x77, 0x6b, 0x04, 0x51, 0x7b, 0xb1, 0xa2, 0x2b, 0xc1, 0x85,
	0x29, 0x15, 0x6a, 0xe2, 0x56, 0x4e, 0xec, 0x64, 0xdd, 0x4d, 0x32, 0x91, 0xed, 0xb0, 0xbb, 0xc4,
	0x96, 0x78, 0x03, 0x24, 0x5e, 0x02, 0x09, 0x89, 0x57, 0xe0, 0x96, 0x27, 0xe0, 0xd6, 0x48, 0x5c,
	0x71, 0xed, 0x27, 0x40, 0x33, 0x93, 0xbf, 0x6d, 0x0d, 0x7b, 0x51, 0x2a, 0x71, 0x95, 0x39, 0xe7,
	0xfb, 0xbe, 0x99, 0xef, 0x9c, 0x89, 0xe6, 0x18, 0x6c, 0x5c, 0x93, 0xfa, 0x95, 0x6f, 0x7a, 0xaf,
	0x3c, 0xa5, 0xe7, 0x12, 0x9f, 0xc0, 0x75, 0xd2, 0xf3, 0xae, 0xcc, 0xbe, 0x4f, 0x3a, 0xa6, 0xef,
	0x90, 0xee, 0x63, 0xdc, 0x72, 0xfc, 0x97, 0xfd, 0xba, 0xd2, 0x20, 0x9d, 0x5c, 0x8b, 0xb4, 0x48,
	0x8e, 0xd1, 0xea, 0xfd, 0x26, 0x8b, 0x58, 0xc0, 0x56, 0x5c, 0xfe, 0xf8, 0xeb, 0x16, 0x51, 0x6c,
	0xd3, 0x7b, 0x4b, 0x7a, 0x9e, 0xd2, 0x26, 0x0d, 0xb3, 0x9d, 0x6b, 0x90, 0xae, 0xef, 0x9a, 0x0d,
	0xdf, 0xe3, 0x4a, 0xd7, 0xee, 0x11, 0xdc, 0x21, 0x96, 0xdd, 0xf6, 0x72, 0x43, 0x62, 0x8e, 0x85,
	0xb9, 0xd9, 0x83, 0x73, 0x1d, 0xd3, 0x69, 0x5f, 0x39, 0xdd, 0xe6, 0x68, 0xe7, 0xc3, 0x29, 0x23,
	0x9d, 0xd7, 0x8e, 0xff, 0x8a, 0xbc, 0xce, 0xb5, 0x08, 0x66, 0x20, 0xfe, 0xc6, 0x6c, 0x3b, 0x96,
	0xe9, 0x13, 0xd7, 0xcb, 0x8d, 0x97, 0x5c, 0x27, 0xfd, 0xb6, 0x03, 0x96, 0x5f, 0x90, 0xfa, 0x19,
	0xad, 0x11, 0x96, 0x40, 0xda, 0xb1, 0x84, 0x54, 0x36, 0x85, 0x32, 0x27, 0xbb, 0x71, 0x24, 0x66,
	0x9a, 0xc4, 0xed, 0x68, 0x92, 0x63, 0x49, 0x7f, 0xfc, 0x2e, 0x6e, 0x83, 0xad, 0xcb, 0x6a, 0x1e,
	0x1f, 0x99, 0xb8, 0x79, 0x8c, 0x9f, 0x1b, 0x03, 0xb5, 0x14, 0xee, 0xea, 0x69, 0xc7, 0x82, 0x4f,
	0xc0, 0xc2, 0x35, 0xa9, 0x9f, 0x5a, 0x42, 0x9a, 0x09, 0x37, 0xe3, 0x48, 0x5c, 0xe5, 0x42, 0x96,
	0x96, 0x74, 0x0e, 0x43, 0x15, 0x2c, 0x5d, 0x93, 0xfa, 0x17, 0x66, 0xc7, 0x16, 0xe6, 0x19, 0x53,
	0x88, 0x23, 0x71, 0x7d, 0xcc, 0xa4, 0x00, 0x3d, 0x27, 0xdd, 0x9b, 0xd3, 0x47, 0x44, 0xf8, 0x0c,
	0x2c, 0x77, 0xec, 0x6e, 0x9f, 0x89, 0x1e, 0x31, 0xd1, 0xfb, 0x71, 0x24, 0x6e, 0x70, 0xd1, 0x08,
	0xa1, 0xaa, 0x85, 0xde, 0xdc, 0x9b, 0xef, 0xd2, 0xfa, 0x98, 0x0d, 0x65, 0xb0, 0x68, 0xbf, 0xb1,
	0x1b, 0xa7, 0x96, 0xb0, 0xc0, 0x74, 0x5b, 0x71, 0x24, 0xae, 0x71, 0x1d, 0xcf, 0x4b, 0xfa, 0x90,
	0x00, 0x3f, 0x06, 0xab, 0xae, 0xed, 0x91, 0xbe, 0xdb, 0xb0, 0xcf, 0xde, 0xf6, 0x6c, 0x61, 0x91,
	0x09, 0xde, 0x8b, 0x23, 0x71, 0x9b, 0x0b, 0xa6, 0x51, 0x49, 0x9f, 0x21, 0xc3, 0x32, 0x00, 0xa3,
	0xf8, 0xd4, 0x12, 0x96, 0x98, 0xf4, 0x7f, 0x71, 0x24, 0x6e, 0xcd, 0x4a, 0xe9, 0x79, 0x53, 0x44,
	0xf8, 0x0c, 0xac, 0x8c, 0xa2, 0xf3, 0x53, 0x4b, 0x58, 0x66, 0xba, 0xff, 0xc7, 0x91, 0x08, 0x67,
	0x75, 0xe7, 0x54, 0x38, 0x4d, 0x85, 0x9f, 0x82, 0xb5, 0x71, 0xc8, 0xfa, 0x92, 0x19, 0x37, 0x73,
	0xe7, 0x86, 0x96, 0x35, 0x47, 0x9f, 0xa5, 0xc3, 0x03, 0xb0, 0xe4, 0xbb, 0x4e, 0xab, 0x65, 0xbb,
	0x02, 0x60, 0x4a, 0x38, 0xb9, 0x86, 0x21, 0x20, 0xe9, 0x23, 0x0a, 0x3c, 0x07, 0xcb, 0xb4, 0x4b,
	0x5f, 0x79, 0xb6, 0x2b, 0xac, 0x30, 0xba, 0x36, 0xb9, 0x80, 0x11, 0x42, 0x2f, 0xe0, 0x43, 0xf0,
	0xc1, 0x65, 0xd5, 0xc4, 0xdf, 0x1e, 0xe3, 0x8b, 0x3c, 0x3e, 0x32, 0xaa, 0xca, 0x78, 0x7d, 0x85,
	0x8d, 0x81, 0x7a, 0x50, 0x2c, 0x84, 0xbb, 0xfa, 0x78, 0x2f, 0xf8, 0x4b, 0x06, 0x2c, 0xbc, 0x24,
	0x9e, 0xef, 0x09, 0xab, 0xd9, 0x79, 0x94, 0x39, 0xf9, 0x31, 0x33, 0xf9, 0xdb, 0xb0, 0x3c, 0xdd,
	0xf3, 0x87, 0x0c, 0xf8, 0x3e, 0x73, 0x89, 0x90, 0x8a, 0xca, 0xd5, 0x3c, 0x2e, 0x1b, 0x83, 0x42,
	0x18, 0x54, 0xf3, 0xb8, 0x64, 0xd4, 0xac, 0x41, 0x21, 0x94, 0xe9, 0xba, 0x60, 0x54, 0x68, 0x70,
	0xa0, 0x86, 0x32, 0xaa, 0x29, 0xef, 0xc8, 0x94, 0x07, 0xc5, 0x50, 0x0e, 0x6a, 0xde, 0x1e, 0x42,
	0x88, 0xfe, 0xa5, 0x8f, 0xf1, 0x73, 0x13, 0x37, 0x8d, 0x41, 0xe1, 0xa0, 0x14, 0x6a, 0xf2, 0xe0,
	0x69, 0x78, 0x2b, 0x1b, 0x68, 0xb2, 0x1c, 0xdc, 0x49, 0x3e, 0x0c, 0x91, 0x76, 0x8b, 0x8d, 0x90,
	0xca, 0x7d, 0x04, 0xea, 0xd0, 0x45, 0x50, 0xa8, 0x59, 0x35, 0x2b, 0xa8, 0x16, 0xf0, 0x11, 0xf5,
	0xc1, 0xcd, 0xfe, 0x03, 0x87, 0xdb, 0x4c, 0x3c, 0xb9, 0x1c, 0x22, 0x74, 0xfb, 0x6c, 0x99, 0x97,
	0x18, 0x68, 0x0f, 0xe2, 0xa1, 0x94, 0xe8, 0x81, 0xca, 0xee, 0x82, 0x2a, 0xf7, 0x69, 0xec, 0x6f,
	0x9c, 0x15, 0x13, 0x9d, 0x95, 0x12, 0x9c, 0x0d, 0xf2, 0x07, 0x6a, 0xf8, 0x40, 0xee, 0xd4, 0x44,
	0x77, 0xe5, 0x64, 0x77, 0xc5, 0x87, 0x72, 0x57, 0x48, 0x74, 0x77, 0x98, 0xec, 0xae, 0xf4, 0x6f,
	0xb8, 0xd3, 0x92, 0x8c, 0x3c, 0x4d, 0x36, 0x52, 0xbe, 0x7f, 0x23, 0x32, 0xfa, 0x48, 0xd9, 0x97,
	0x2b, 0x35, 0x6f, 0x6f, 0x57, 0xe7, 0xcf, 0x16, 0x9d, 0x2f, 0x9e, 0x6f, 0xfa, 0x7d, 0x4f, 0x58,
	0xbb, 0x39, 0x5f, 0x78, 0x5e, 0xd2, 0x87, 0x04, 0xf8, 0x09, 0x78, 0x44, 0xc7, 0xb5, 0xb0, 0x9e,
	0x4d, 0xa1, 0x15, 0x55, 0x50, 0x66, 0x47, 0xb9, 0xf2, 0xb9, 0xe9, 0xb4, 0x4f, 0xbb, 0x4d, 0x72,
	0xb2, 0x11, 0x47, 0xe2, 0xca, 0x70, 0xb4, 0x99, 0x4e, 0x5b, 0xd2, 0x99, 0x8c, 0x8e, 0x0a, 0xaf,
	0xdf, 0x68, 0xd8, 0x9e, 0xa7, 0x9b, 0xbe, 0x2d, 0x6c, 0x64, 0x53, 0x28, 0x3d, 0x3d, 0x2a, 0xa6,
	0x40, 0x49, 0x9f, 0xa6, 0xd2, 0xc9, 0x6c, 0xbb, 0x2e, 0x71, 0x85, 0xcd, 0x9b, 0x93, 0x99, 0xa5,
	0x25, 0x9d, 0xc3, 0xf0, 0xcf, 0x14, 0x00, 0x0d, 0xd7, 0x36, 0x7d, 0xfb, 0xcc, 0xe9, 0xd8, 0xc2,
	0x16, 0x63, 0xff, 0x9a, 0x9a, 0x4c, 0xb1, 0x09, 0x48, 0x9f, 0xe5, 0x9f, 0x53, 0xe0, 0xa7, 0xd4,
	0x25, 0x42, 0x15, 0x8d, 0xb5, 0x8a, 0xde, 0x81, 0xb1, 0x27, 0x57, 0xd8, 0xef, 0xa0, 0x14, 0xca,
	0x18, 0x15, 0xaa, 0x79, 0xac, 0x1a, 0x41, 0x9e, 0xe1, 0x32, 0x46, 0x45, 0xf6, 0xfe, 0x0e, 0x63,
	0xda, 0x60, 0x95, 0xab, 0xe4, 0xea, 0x59, 0xd6, 0x40, 0xb4, 0xff, 0x45, 0x83, 0xbf, 0xd1, 0x3c,
	0xad, 0x21, 0x76, 0x33, 0xc3, 0x60, 0xb2, 0x46, 0x35, 0x85, 0xfd, 0xee, 0xcb, 0x15, 0x74, 0x11,
	0x54, 0xf7, 0xb1, 0x81, 0x2a, 0xda, 0x1d, 0xf2, 0x29, 0x75, 0x65, 0x57, 0x9f, 0xaa, 0x8d, 0x95,
	0xda, 0xef, 0x59, 0xa3, 0x52, 0xe1, 0xad, 0x52, 0x27, 0xe0, 0x7f, 0xb2, 0xd4, 0x89, 0x7d, 0xf8,
	0x25, 0x58, 0x62, 0x85, 0x13, 0x57, 0xd8, 0x66, 0x65, 0x1e, 0x4d, 0x06, 0xfd, 0x10, 0x78, 0xe7,
	0xc1, 0x3d, 0xda, 0x09, 0x3e, 0x01, 0xf3, 0xc4, 0x6d, 0x09, 0x3b, 0xd9, 0x14, 0x5a, 0x38, 0xd9,
	0x89, 0x23, 0x11, 0xf0, 0x0d, 0x89, 0xdb, 0x62, 0x1f, 0x6f, 0x9b, 0x73, 0x3a, 0x25, 0x9c, 0xbc,
	0xb8, 0xf8, 0xec, 0xbe, 0xbe, 0x75, 0xeb, 0x8b, 0x8c, 0x5e, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0xd2, 0x84, 0x51, 0x8e, 0x0b, 0x00, 0x00,
}
