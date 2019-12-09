// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.proto

package scheduler

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
//定时任务
type SchedulerTask struct {
	//
	//scheduler task id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//定时任务名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//任务对象类型
	JobType string `protobuf:"bytes,3,opt,name=job_type,json=jobType,proto3" json:"job_type" form:"job_type"`
	//
	//定时任务执行周期类型
	TaskType string `protobuf:"bytes,4,opt,name=task_type,json=taskType,proto3" json:"task_type" form:"task_type"`
	//
	//任务执行方式，目前只使用到了http
	CmdType string `protobuf:"bytes,5,opt,name=cmd_type,json=cmdType,proto3" json:"cmd_type" form:"cmd_type"`
	//
	//定时时间：x/5 x x x x， 2018-09-04 15:56:23
	TaskScheduler string `protobuf:"bytes,6,opt,name=task_scheduler,json=taskScheduler,proto3" json:"task_scheduler" form:"task_scheduler"`
	//
	//相关参数
	Annotations *SchedulerTask_Annotations `protobuf:"bytes,7,opt,name=annotations,proto3" json:"annotations" form:"annotations"`
	//
	//发起定时任务的源 id
	SrcId string `protobuf:"bytes,8,opt,name=src_id,json=srcId,proto3" json:"src_id" form:"src_id"`
	//
	//任务的具体配置, http或脚本
	CmdConfig *SchedulerTask_CmdConfig `protobuf:"bytes,9,opt,name=cmd_config,json=cmdConfig,proto3" json:"cmd_config" form:"cmd_config"`
	//
	//任务创建者来源
	Assignner string `protobuf:"bytes,10,opt,name=assignner,proto3" json:"assignner" form:"assignner"`
	//
	//0 表示不禁用, 1 表示禁用
	Disable int32 `protobuf:"varint,11,opt,name=disable,proto3" json:"disable" form:"disable"`
	//
	//0 在页面可见，1 不可见
	Invisible int32 `protobuf:"varint,12,opt,name=invisible,proto3" json:"invisible" form:"invisible"`
	//
	//org
	Org int32 `protobuf:"varint,13,opt,name=org,proto3" json:"org" form:"org"`
	//
	//创建用户
	User string `protobuf:"bytes,14,opt,name=user,proto3" json:"user" form:"user"`
	//
	//操作权限
	OperateAuthorizers []string `protobuf:"bytes,15,rep,name=operateAuthorizers,proto3" json:"operateAuthorizers" form:"operateAuthorizers"`
	//
	//删除权限
	DeleteAuthorizers []string `protobuf:"bytes,16,rep,name=deleteAuthorizers,proto3" json:"deleteAuthorizers" form:"deleteAuthorizers"`
	//
	//更新权限
	UpdateAuthorizers []string `protobuf:"bytes,17,rep,name=updateAuthorizers,proto3" json:"updateAuthorizers" form:"updateAuthorizers"`
	//
	//创建时间
	CreateTime string `protobuf:"bytes,18,opt,name=create_time,json=createTime,proto3" json:"create_time" form:"create_time"`
	//
	//更新时间
	UpdateTime string `protobuf:"bytes,19,opt,name=update_time,json=updateTime,proto3" json:"update_time" form:"update_time"`
	//
	//job id
	JobId string `protobuf:"bytes,20,opt,name=job_id,json=jobId,proto3" json:"job_id" form:"job_id"`
	//
	//运行状态
	Status int32 `protobuf:"varint,21,opt,name=status,proto3" json:"status" form:"status"`
	//
	//回调
	Callback *SchedulerTask_Callback `protobuf:"bytes,22,opt,name=callback,proto3" json:"callback" form:"callback"`
	//
	//调度结果信息。{u'message': u'Success', u'code': 0, u'data': {u'execId': u'201809_task155123648b640b916475'}, u'error': u'\u6210\u529f\u3002Success'}
	Error                string   `protobuf:"bytes,23,opt,name=error,proto3" json:"error" form:"error"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulerTask) Reset()         { *m = SchedulerTask{} }
func (m *SchedulerTask) String() string { return proto.CompactTextString(m) }
func (*SchedulerTask) ProtoMessage()    {}
func (*SchedulerTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}
func (m *SchedulerTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerTask.Unmarshal(m, b)
}
func (m *SchedulerTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerTask.Marshal(b, m, deterministic)
}
func (m *SchedulerTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerTask.Merge(m, src)
}
func (m *SchedulerTask) XXX_Size() int {
	return xxx_messageInfo_SchedulerTask.Size(m)
}
func (m *SchedulerTask) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerTask.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerTask proto.InternalMessageInfo

func (m *SchedulerTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SchedulerTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchedulerTask) GetJobType() string {
	if m != nil {
		return m.JobType
	}
	return ""
}

func (m *SchedulerTask) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *SchedulerTask) GetCmdType() string {
	if m != nil {
		return m.CmdType
	}
	return ""
}

func (m *SchedulerTask) GetTaskScheduler() string {
	if m != nil {
		return m.TaskScheduler
	}
	return ""
}

func (m *SchedulerTask) GetAnnotations() *SchedulerTask_Annotations {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *SchedulerTask) GetSrcId() string {
	if m != nil {
		return m.SrcId
	}
	return ""
}

func (m *SchedulerTask) GetCmdConfig() *SchedulerTask_CmdConfig {
	if m != nil {
		return m.CmdConfig
	}
	return nil
}

func (m *SchedulerTask) GetAssignner() string {
	if m != nil {
		return m.Assignner
	}
	return ""
}

func (m *SchedulerTask) GetDisable() int32 {
	if m != nil {
		return m.Disable
	}
	return 0
}

func (m *SchedulerTask) GetInvisible() int32 {
	if m != nil {
		return m.Invisible
	}
	return 0
}

func (m *SchedulerTask) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *SchedulerTask) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SchedulerTask) GetOperateAuthorizers() []string {
	if m != nil {
		return m.OperateAuthorizers
	}
	return nil
}

func (m *SchedulerTask) GetDeleteAuthorizers() []string {
	if m != nil {
		return m.DeleteAuthorizers
	}
	return nil
}

func (m *SchedulerTask) GetUpdateAuthorizers() []string {
	if m != nil {
		return m.UpdateAuthorizers
	}
	return nil
}

func (m *SchedulerTask) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *SchedulerTask) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *SchedulerTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *SchedulerTask) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SchedulerTask) GetCallback() *SchedulerTask_Callback {
	if m != nil {
		return m.Callback
	}
	return nil
}

func (m *SchedulerTask) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SchedulerTask_Annotations struct {
	//
	//应用ID
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulerTask_Annotations) Reset()         { *m = SchedulerTask_Annotations{} }
func (m *SchedulerTask_Annotations) String() string { return proto.CompactTextString(m) }
func (*SchedulerTask_Annotations) ProtoMessage()    {}
func (*SchedulerTask_Annotations) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0, 0}
}
func (m *SchedulerTask_Annotations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerTask_Annotations.Unmarshal(m, b)
}
func (m *SchedulerTask_Annotations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerTask_Annotations.Marshal(b, m, deterministic)
}
func (m *SchedulerTask_Annotations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerTask_Annotations.Merge(m, src)
}
func (m *SchedulerTask_Annotations) XXX_Size() int {
	return xxx_messageInfo_SchedulerTask_Annotations.Size(m)
}
func (m *SchedulerTask_Annotations) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerTask_Annotations.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerTask_Annotations proto.InternalMessageInfo

func (m *SchedulerTask_Annotations) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type SchedulerTask_CmdConfig struct {
	//
	//uri
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url" form:"url"`
	//
	//名字服务
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name" form:"service_name"`
	//
	//端口
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port" form:"port"`
	//
	//名字服务
	Host string `protobuf:"bytes,4,opt,name=host,proto3" json:"host" form:"host"`
	//
	//http 请求
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method" form:"method"`
	//
	//请求头
	Headers *types.Struct `protobuf:"bytes,6,opt,name=headers,proto3" json:"headers" form:"headers"`
	//
	//参数
	Params               *types.Struct `protobuf:"bytes,7,opt,name=params,proto3" json:"params" form:"params"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SchedulerTask_CmdConfig) Reset()         { *m = SchedulerTask_CmdConfig{} }
func (m *SchedulerTask_CmdConfig) String() string { return proto.CompactTextString(m) }
func (*SchedulerTask_CmdConfig) ProtoMessage()    {}
func (*SchedulerTask_CmdConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0, 1}
}
func (m *SchedulerTask_CmdConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerTask_CmdConfig.Unmarshal(m, b)
}
func (m *SchedulerTask_CmdConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerTask_CmdConfig.Marshal(b, m, deterministic)
}
func (m *SchedulerTask_CmdConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerTask_CmdConfig.Merge(m, src)
}
func (m *SchedulerTask_CmdConfig) XXX_Size() int {
	return xxx_messageInfo_SchedulerTask_CmdConfig.Size(m)
}
func (m *SchedulerTask_CmdConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerTask_CmdConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerTask_CmdConfig proto.InternalMessageInfo

func (m *SchedulerTask_CmdConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SchedulerTask_CmdConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *SchedulerTask_CmdConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SchedulerTask_CmdConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SchedulerTask_CmdConfig) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SchedulerTask_CmdConfig) GetHeaders() *types.Struct {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *SchedulerTask_CmdConfig) GetParams() *types.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

type SchedulerTask_Callback struct {
	//
	//url
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url" form:"url"`
	//
	//header host
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host" form:"host"`
	//
	//名字服务
	EnsName              string   `protobuf:"bytes,3,opt,name=ensName,proto3" json:"ensName" form:"ensName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulerTask_Callback) Reset()         { *m = SchedulerTask_Callback{} }
func (m *SchedulerTask_Callback) String() string { return proto.CompactTextString(m) }
func (*SchedulerTask_Callback) ProtoMessage()    {}
func (*SchedulerTask_Callback) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0, 2}
}
func (m *SchedulerTask_Callback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulerTask_Callback.Unmarshal(m, b)
}
func (m *SchedulerTask_Callback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulerTask_Callback.Marshal(b, m, deterministic)
}
func (m *SchedulerTask_Callback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulerTask_Callback.Merge(m, src)
}
func (m *SchedulerTask_Callback) XXX_Size() int {
	return xxx_messageInfo_SchedulerTask_Callback.Size(m)
}
func (m *SchedulerTask_Callback) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulerTask_Callback.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulerTask_Callback proto.InternalMessageInfo

func (m *SchedulerTask_Callback) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SchedulerTask_Callback) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SchedulerTask_Callback) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func init() {
	proto.RegisterType((*SchedulerTask)(nil), "scheduler.SchedulerTask")
	proto.RegisterType((*SchedulerTask_Annotations)(nil), "scheduler.SchedulerTask.Annotations")
	proto.RegisterType((*SchedulerTask_CmdConfig)(nil), "scheduler.SchedulerTask.CmdConfig")
	proto.RegisterType((*SchedulerTask_Callback)(nil), "scheduler.SchedulerTask.Callback")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 1214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcf, 0x6f, 0xdb, 0xb6,
	0x17, 0xaf, 0xf3, 0xc3, 0x89, 0xe9, 0xe6, 0x17, 0xd3, 0x1f, 0xfa, 0x06, 0x5f, 0x40, 0xa9, 0x9a,
	0x15, 0xf2, 0x12, 0xc9, 0x8e, 0xd3, 0x16, 0xa8, 0x87, 0xc1, 0x8b, 0xb3, 0x16, 0xe8, 0x30, 0xf4,
	0xc0, 0xe6, 0x30, 0xc4, 0x75, 0x32, 0x5a, 0x62, 0x6c, 0x35, 0x92, 0xa9, 0x91, 0x74, 0x8b, 0xb6,
	0xee, 0xd0, 0xbf, 0x67, 0xc0, 0xfe, 0x8f, 0x5d, 0x77, 0xd9, 0xd1, 0x05, 0x7a, 0xdc, 0xd1, 0xe7,
	0x01, 0x1b, 0x48, 0x4a, 0xb2, 0x92, 0xb4, 0x43, 0xaf, 0xcb, 0x25, 0xe4, 0xfb, 0x7c, 0x3e, 0x7c,
	0x8f, 0xef, 0x59, 0xef, 0x11, 0x00, 0x81, 0xf9, 0x99, 0x1b, 0x33, 0x2a, 0x28, 0x2c, 0x71, 0xaf,
	0x4f, 0xfc, 0x61, 0x48, 0xd8, 0x86, 0xd3, 0x0b, 0x44, 0x7f, 0xd8, 0x75, 0x3d, 0x1a, 0x55, 0x7b,
	0xb4, 0x47, 0xab, 0x8a, 0xd1, 0x1d, 0x9e, 0xaa, 0x9d, 0xda, 0xa8, 0x95, 0x56, 0x6e, 0xdc, 0xcf,
	0xd1, 0xa3, 0x97, 0x81, 0x38, 0xa3, 0x2f, 0xab, 0x3d, 0xea, 0x28, 0xd0, 0x79, 0x81, 0xc3, 0xc0,
	0xc7, 0x82, 0x32, 0x5e, 0xcd, 0x96, 0x89, 0xee, 0xff, 0x3d, 0x4a, 0x7b, 0x21, 0x99, 0x9e, 0xce,
	0x05, 0x1b, 0x7a, 0x42, 0xa3, 0xd6, 0x1f, 0xeb, 0x60, 0xe9, 0x69, 0x1a, 0xd2, 0x21, 0xe6, 0x67,
	0xf0, 0x2e, 0x98, 0x09, 0x7c, 0xa3, 0xb0, 0x59, 0xb0, 0x4b, 0xad, 0xad, 0xc9, 0xd8, 0x2c, 0x9d,
	0x52, 0x16, 0x35, 0xac, 0xc0, 0xb7, 0x3e, 0xbc, 0x37, 0xd7, 0xc1, 0xda, 0x71, 0xbb, 0xe6, 0x3c,
	0xc0, 0xce, 0xe9, 0xbe, 0xf3, 0xa8, 0xf3, 0xa6, 0x7e, 0xf7, 0xed, 0x16, 0x9a, 0x09, 0x7c, 0x58,
	0x01, 0x73, 0x03, 0x1c, 0x11, 0x63, 0x46, 0xe9, 0xae, 0x4f, 0xc6, 0x66, 0x59, 0xeb, 0xa4, 0x55,
	0x2a, 0x67, 0xe2, 0x2b, 0x48, 0x51, 0xa0, 0x0b, 0x16, 0x9f, 0xd3, 0xee, 0x89, 0x78, 0x15, 0x13,
	0x63, 0x56, 0xd1, 0xd7, 0x27, 0x63, 0x73, 0x45, 0xd3, 0x53, 0xc4, 0x42, 0x0b, 0xcf, 0x69, 0xf7,
	0xf0, 0x55, 0x4c, 0xe0, 0x2e, 0x28, 0xc9, 0x04, 0x6a, 0xc1, 0x9c, 0x12, 0x5c, 0x9b, 0x8c, 0xcd,
	0x55, 0x2d, 0xc8, 0x20, 0x0b, 0x2d, 0xca, 0xb5, 0x92, 0xb8, 0x60, 0xd1, 0x8b, 0x7c, 0xad, 0x98,
	0xbf, 0xe8, 0x22, 0x45, 0x2c, 0xb4, 0xe0, 0x45, 0xbe, 0xe2, 0x3f, 0x02, 0xcb, 0xea, 0x9c, 0xac,
	0x38, 0x46, 0x51, 0xa9, 0xcc, 0xc9, 0xd8, 0xbc, 0x9e, 0xf3, 0x93, 0xe1, 0xe9, 0x8d, 0x96, 0xa4,
	0x39, 0xcb, 0x1f, 0x3c, 0x06, 0x65, 0x3c, 0x18, 0x50, 0x81, 0x45, 0x40, 0x07, 0xdc, 0x58, 0xd8,
	0x2c, 0xd8, 0xe5, 0xfa, 0x96, 0x9b, 0xc9, 0xdc, 0x73, 0xa9, 0x76, 0xf7, 0xa7, 0xdc, 0xd6, 0x8d,
	0xc9, 0xd8, 0x84, 0xda, 0x55, 0xee, 0x08, 0x0b, 0xe5, 0x0f, 0x84, 0x2e, 0x28, 0x72, 0xe6, 0x9d,
	0x04, 0xbe, 0xb1, 0xa8, 0xe2, 0xbb, 0x39, 0x19, 0x9b, 0x4b, 0x5a, 0xa4, 0xed, 0x3a, 0xae, 0x02,
	0x9a, 0xe7, 0xcc, 0x7b, 0xec, 0xc3, 0x1f, 0x00, 0x90, 0xb7, 0xf5, 0xe8, 0xe0, 0x34, 0xe8, 0x19,
	0x25, 0x15, 0x8e, 0xf5, 0xc9, 0x70, 0x0e, 0x22, 0xff, 0x40, 0x31, 0x55, 0xfd, 0xd6, 0xa6, 0xd9,
	0xd2, 0x7a, 0x0b, 0x95, 0xbc, 0x94, 0x01, 0xeb, 0xa0, 0x84, 0x39, 0x0f, 0x7a, 0x83, 0x01, 0x61,
	0x06, 0xb8, 0x58, 0x94, 0x0c, 0xb2, 0xd0, 0x94, 0x06, 0x77, 0xc0, 0x82, 0x1f, 0x70, 0xdc, 0x0d,
	0x89, 0x51, 0xde, 0x2c, 0xd8, 0xf3, 0x2d, 0x38, 0x19, 0x9b, 0xcb, 0x5a, 0x91, 0x00, 0x16, 0x4a,
	0x29, 0xd2, 0x43, 0x30, 0x78, 0x11, 0xf0, 0x40, 0xf2, 0xaf, 0x2a, 0x7e, 0xce, 0x43, 0x06, 0x59,
	0x68, 0x4a, 0x83, 0x77, 0xc0, 0x2c, 0x65, 0x3d, 0x63, 0x29, 0x63, 0x03, 0xcd, 0xa6, 0xac, 0xa7,
	0x32, 0xb3, 0x7a, 0x05, 0x49, 0x02, 0xfc, 0x0e, 0xcc, 0x0d, 0x39, 0x61, 0xc6, 0xb2, 0x0a, 0xfc,
	0xfe, 0xf4, 0xd7, 0x2a, 0xad, 0x92, 0x79, 0x1b, 0xdc, 0x3a, 0x6e, 0x63, 0xe7, 0xf5, 0xbe, 0x73,
	0x54, 0x73, 0x1e, 0x74, 0xda, 0x6e, 0xb6, 0x3e, 0x71, 0x3a, 0x6f, 0xea, 0x3b, 0x7b, 0xbb, 0x6f,
	0xb7, 0x90, 0x3a, 0x03, 0x0a, 0x00, 0x69, 0x4c, 0x18, 0x16, 0x64, 0x7f, 0x28, 0xfa, 0x94, 0x05,
	0xaf, 0x09, 0xe3, 0xc6, 0xca, 0xe6, 0xac, 0x5d, 0x6a, 0x7d, 0x3b, 0x19, 0x9b, 0xff, 0x4b, 0x42,
	0xb8, 0xc4, 0xf9, 0x6c, 0x3f, 0x1f, 0x39, 0x1f, 0xfe, 0x04, 0xd6, 0x7c, 0x12, 0x92, 0xf3, 0x4e,
	0x57, 0x95, 0xd3, 0x83, 0xc9, 0xd8, 0x34, 0x92, 0xac, 0x5e, 0xa4, 0x7c, 0xb6, 0xcf, 0xcb, 0xa7,
	0x4b, 0x97, 0xc3, 0xd8, 0xbf, 0x70, 0xcf, 0xb5, 0x8b, 0x2e, 0x2f, 0x51, 0x3e, 0xdf, 0xe5, 0x25,
	0x29, 0xfc, 0xb3, 0x00, 0xca, 0x1e, 0x23, 0x58, 0x90, 0x13, 0x11, 0x44, 0xc4, 0x80, 0xaa, 0x5e,
	0xbf, 0x15, 0xa6, 0xdf, 0x4a, 0x0e, 0x95, 0x8e, 0x7e, 0x2d, 0x80, 0x5f, 0x0a, 0xc7, 0xb6, 0xdd,
	0x6c, 0xb4, 0x77, 0xa5, 0x23, 0xe9, 0xed, 0xcb, 0x4a, 0x53, 0xfd, 0x7f, 0x73, 0xf7, 0x6d, 0xc5,
	0xb1, 0x77, 0xdb, 0x35, 0xa7, 0xde, 0x19, 0xd5, 0x14, 0x5e, 0x71, 0xec, 0xbd, 0x76, 0xcd, 0xd9,
	0x4d, 0xf7, 0xa3, 0xf6, 0xae, 0x53, 0xd7, 0xaa, 0x4a, 0xfb, 0x70, 0xb3, 0x63, 0xd7, 0xdb, 0x35,
	0x67, 0xaf, 0x33, 0x52, 0x1c, 0x6d, 0x6e, 0xd8, 0xed, 0x9a, 0x73, 0x2f, 0xdd, 0x4c, 0xd7, 0xf6,
	0x33, 0x57, 0xfd, 0xdf, 0xae, 0x34, 0xed, 0xa3, 0x51, 0x7b, 0xdb, 0xe9, 0xd8, 0xcd, 0xc6, 0x47,
	0xe4, 0x39, 0x75, 0x73, 0x0b, 0x01, 0x1d, 0xff, 0x61, 0x10, 0x11, 0x75, 0x59, 0x9d, 0x02, 0x7d,
	0xd9, 0xf5, 0x4b, 0x97, 0xcd, 0xa1, 0xff, 0xc9, 0xcb, 0xea, 0xf8, 0xd5, 0x65, 0x5d, 0x50, 0x94,
	0xad, 0x3e, 0xf0, 0x8d, 0x6b, 0x17, 0x3b, 0x99, 0xb6, 0x67, 0x9d, 0xec, 0x39, 0xed, 0x3e, 0x96,
	0xf3, 0xa5, 0xc8, 0x05, 0x16, 0x43, 0x6e, 0x5c, 0x57, 0x1f, 0xf7, 0x5a, 0xae, 0xf3, 0x29, 0xbb,
	0x85, 0x12, 0x02, 0x44, 0x60, 0xd1, 0xc3, 0x61, 0xd8, 0xc5, 0xde, 0x99, 0x71, 0x43, 0xb5, 0xbc,
	0x5b, 0x9f, 0x6e, 0x79, 0x09, 0xf1, 0xdc, 0x7c, 0x48, 0x6c, 0x16, 0xca, 0xce, 0x81, 0x77, 0xc0,
	0x3c, 0x61, 0x8c, 0x32, 0xe3, 0xa6, 0x8a, 0x76, 0x75, 0x32, 0x36, 0xaf, 0x6a, 0xb6, 0x32, 0x5b,
	0x48, 0xc3, 0x1b, 0xf7, 0x40, 0x39, 0xd7, 0xd4, 0xa5, 0x0c, 0xc7, 0xf1, 0xe3, 0x74, 0x9c, 0xe6,
	0x64, 0xca, 0x6c, 0x21, 0x0d, 0x6f, 0xfc, 0x35, 0x03, 0x4a, 0x59, 0xf7, 0x85, 0x9b, 0x60, 0x76,
	0xc8, 0xc2, 0x44, 0xb3, 0x3c, 0xed, 0x62, 0x43, 0x16, 0x5a, 0x48, 0x42, 0xb0, 0x01, 0xae, 0x72,
	0xc2, 0x5e, 0x04, 0x1e, 0x39, 0xc9, 0x4d, 0x5d, 0x99, 0xc3, 0xf5, 0x24, 0x27, 0x39, 0xd4, 0x42,
	0xe5, 0x64, 0xfb, 0x44, 0x8e, 0xdf, 0xfb, 0x60, 0x2e, 0xa6, 0x4c, 0xa8, 0xd1, 0x3b, 0xdf, 0xb2,
	0xa6, 0xbd, 0x4f, 0x5a, 0x65, 0xd6, 0x57, 0x56, 0xff, 0x4e, 0xff, 0x0a, 0xc6, 0xbb, 0x77, 0x73,
	0x48, 0xf1, 0xe1, 0x6d, 0x30, 0xd7, 0xa7, 0x5c, 0x24, 0x13, 0x78, 0x65, 0xaa, 0x93, 0x56, 0x0b,
	0x29, 0x50, 0x96, 0x29, 0x22, 0xa2, 0x4f, 0xfd, 0x64, 0xec, 0xe6, 0xca, 0xa4, 0xed, 0x16, 0x4a,
	0x08, 0xf0, 0x21, 0x58, 0xe8, 0x13, 0xec, 0xcb, 0x26, 0x52, 0x54, 0x55, 0xba, 0xe9, 0xea, 0x97,
	0x8a, 0x9b, 0xbe, 0x54, 0xdc, 0xa7, 0xea, 0xa5, 0x92, 0x1f, 0x13, 0x89, 0xc2, 0x42, 0xa9, 0x16,
	0xb6, 0x40, 0x31, 0xc6, 0x0c, 0x47, 0xe9, 0xb4, 0xfd, 0xe4, 0x29, 0xb9, 0x50, 0xb4, 0xc0, 0x42,
	0x89, 0x72, 0xe3, 0xf7, 0x02, 0x58, 0x4c, 0x7f, 0x09, 0x90, 0xe5, 0xb3, 0xff, 0xe3, 0xf9, 0xec,
	0x7f, 0x78, 0x6f, 0x3e, 0x01, 0xdf, 0xdb, 0xcf, 0xba, 0x76, 0x5f, 0x88, 0x98, 0x37, 0x47, 0xa7,
	0x22, 0x1e, 0x9d, 0x06, 0x21, 0xa9, 0x34, 0xaa, 0xd5, 0x4a, 0xb3, 0xed, 0xec, 0x3b, 0x47, 0xd8,
	0x79, 0x5d, 0x73, 0x1e, 0x6c, 0xdf, 0xf9, 0xe6, 0x76, 0xf5, 0x8b, 0xe6, 0xd7, 0x3f, 0x9f, 0x8c,
	0x6e, 0x35, 0x76, 0xdc, 0xaf, 0x3a, 0xdb, 0x97, 0x40, 0x89, 0x75, 0x74, 0x3d, 0xd3, 0xdc, 0xce,
	0xfc, 0x5b, 0x6e, 0x77, 0xc0, 0x02, 0x19, 0x70, 0x59, 0xc3, 0xe4, 0xd9, 0x94, 0xcb, 0x4b, 0x02,
	0x58, 0x28, 0xa5, 0xb4, 0x1e, 0x1e, 0x1d, 0xf4, 0xa8, 0x4b, 0x30, 0x7f, 0x45, 0x63, 0xee, 0x86,
	0xd4, 0xc3, 0x61, 0xd5, 0xa3, 0x03, 0xc1, 0xb0, 0x27, 0xb8, 0x7e, 0x0c, 0x32, 0x12, 0x53, 0x27,
	0xa2, 0x3e, 0x09, 0x79, 0x35, 0x21, 0x56, 0xd5, 0xb6, 0x9a, 0x7d, 0x2e, 0xdd, 0xa2, 0x62, 0xee,
	0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x70, 0x74, 0x3b, 0xc4, 0x0a, 0x00, 0x00,
}
