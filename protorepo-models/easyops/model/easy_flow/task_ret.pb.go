// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task_ret.proto

package easy_flow

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
//部署任务返回
type TaskRet struct {
	//
	//任务Id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//姓名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//组Id
	GroupId string `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId" form:"groupId"`
	//
	//应用Id
	AppId string `protobuf:"bytes,4,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//分批数量
	BatchNum int32 `protobuf:"varint,5,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//分批间隔
	BatchInterval int32 `protobuf:"varint,6,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	//
	//是否暂停失败
	FailedStop string `protobuf:"bytes,7,opt,name=failedStop,proto3" json:"failedStop" form:"failedStop"`
	//
	//包Id
	PackageId string `protobuf:"bytes,8,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//类型
	Type string `protobuf:"bytes,9,opt,name=type,proto3" json:"type" form:"type"`
	//
	//操作
	Operation string `protobuf:"bytes,10,opt,name=operation,proto3" json:"operation" form:"operation"`
	//
	//客户id
	Org int32 `protobuf:"varint,11,opt,name=org,proto3" json:"org" form:"org"`
	//
	//操作人
	Operator string `protobuf:"bytes,12,opt,name=operator,proto3" json:"operator" form:"operator"`
	//
	//状态
	Status string `protobuf:"bytes,13,opt,name=status,proto3" json:"status" form:"status"`
	//
	//错误码
	Code int32 `protobuf:"varint,14,opt,name=code,proto3" json:"code" form:"code"`
	//
	//使用次数
	UsedTime int32 `protobuf:"varint,15,opt,name=usedTime,proto3" json:"usedTime" form:"usedTime"`
	//
	//开始时间
	StartTime string `protobuf:"bytes,16,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//结束时间
	EndTime string `protobuf:"bytes,17,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	//
	//额外信息
	ExtraInfo *types.Struct `protobuf:"bytes,18,opt,name=extraInfo,proto3" json:"extraInfo" form:"extraInfo"`
	//
	//目标设备日志
	TargetsLog           []*TaskRet_TargetsLog `protobuf:"bytes,19,rep,name=targetsLog,proto3" json:"targetsLog" form:"targetsLog"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskRet) Reset()         { *m = TaskRet{} }
func (m *TaskRet) String() string { return proto.CompactTextString(m) }
func (*TaskRet) ProtoMessage()    {}
func (*TaskRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee839931d343e4e, []int{0}
}
func (m *TaskRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRet.Unmarshal(m, b)
}
func (m *TaskRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRet.Marshal(b, m, deterministic)
}
func (m *TaskRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRet.Merge(m, src)
}
func (m *TaskRet) XXX_Size() int {
	return xxx_messageInfo_TaskRet.Size(m)
}
func (m *TaskRet) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRet.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRet proto.InternalMessageInfo

func (m *TaskRet) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskRet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskRet) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *TaskRet) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *TaskRet) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *TaskRet) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

func (m *TaskRet) GetFailedStop() string {
	if m != nil {
		return m.FailedStop
	}
	return ""
}

func (m *TaskRet) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *TaskRet) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskRet) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *TaskRet) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *TaskRet) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TaskRet) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRet) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TaskRet) GetUsedTime() int32 {
	if m != nil {
		return m.UsedTime
	}
	return 0
}

func (m *TaskRet) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *TaskRet) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *TaskRet) GetExtraInfo() *types.Struct {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

func (m *TaskRet) GetTargetsLog() []*TaskRet_TargetsLog {
	if m != nil {
		return m.TargetsLog
	}
	return nil
}

type TaskRet_TargetsLog struct {
	//
	//目标设备Id
	TargetId string `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId" form:"targetId"`
	//
	//目标设备名称
	TargetName string `protobuf:"bytes,2,opt,name=targetName,proto3" json:"targetName" form:"targetName"`
	//
	//状态
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status" form:"status"`
	//
	//系统状态
	SysStatus string `protobuf:"bytes,4,opt,name=sysStatus,proto3" json:"sysStatus" form:"sysStatus"`
	//
	//错误码
	Code int32 `protobuf:"varint,5,opt,name=code,proto3" json:"code" form:"code"`
	//
	//消息
	Msg string `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//进度
	Progress *types.Struct `protobuf:"bytes,7,opt,name=progress,proto3" json:"progress" form:"progress"`
	//
	//执行日志
	ActionsLog           []*TaskRet_TargetsLog_ActionsLog `protobuf:"bytes,8,rep,name=actionsLog,proto3" json:"actionsLog" form:"actionsLog"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *TaskRet_TargetsLog) Reset()         { *m = TaskRet_TargetsLog{} }
func (m *TaskRet_TargetsLog) String() string { return proto.CompactTextString(m) }
func (*TaskRet_TargetsLog) ProtoMessage()    {}
func (*TaskRet_TargetsLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee839931d343e4e, []int{0, 0}
}
func (m *TaskRet_TargetsLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRet_TargetsLog.Unmarshal(m, b)
}
func (m *TaskRet_TargetsLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRet_TargetsLog.Marshal(b, m, deterministic)
}
func (m *TaskRet_TargetsLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRet_TargetsLog.Merge(m, src)
}
func (m *TaskRet_TargetsLog) XXX_Size() int {
	return xxx_messageInfo_TaskRet_TargetsLog.Size(m)
}
func (m *TaskRet_TargetsLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRet_TargetsLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRet_TargetsLog proto.InternalMessageInfo

func (m *TaskRet_TargetsLog) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *TaskRet_TargetsLog) GetTargetName() string {
	if m != nil {
		return m.TargetName
	}
	return ""
}

func (m *TaskRet_TargetsLog) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRet_TargetsLog) GetSysStatus() string {
	if m != nil {
		return m.SysStatus
	}
	return ""
}

func (m *TaskRet_TargetsLog) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TaskRet_TargetsLog) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TaskRet_TargetsLog) GetProgress() *types.Struct {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *TaskRet_TargetsLog) GetActionsLog() []*TaskRet_TargetsLog_ActionsLog {
	if m != nil {
		return m.ActionsLog
	}
	return nil
}

type TaskRet_TargetsLog_ActionsLog struct {
	//
	//执行名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	//
	//动作
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action" form:"action"`
	//
	//状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" form:"status"`
	//
	//错误码
	Code int32 `protobuf:"varint,5,opt,name=code,proto3" json:"code" form:"code"`
	//
	//额外信息
	ExtInfo *types.Struct `protobuf:"bytes,6,opt,name=extInfo,proto3" json:"extInfo" form:"extInfo"`
	//
	//消息
	Msg string `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//进度
	Progress *types.Struct `protobuf:"bytes,8,opt,name=progress,proto3" json:"progress" form:"progress"`
	//
	//使用次数
	UsedTime int32 `protobuf:"varint,9,opt,name=usedTime,proto3" json:"usedTime" form:"usedTime"`
	//
	//开始时间
	StartTime string `protobuf:"bytes,10,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//结束时间
	EndTime              string   `protobuf:"bytes,11,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRet_TargetsLog_ActionsLog) Reset()         { *m = TaskRet_TargetsLog_ActionsLog{} }
func (m *TaskRet_TargetsLog_ActionsLog) String() string { return proto.CompactTextString(m) }
func (*TaskRet_TargetsLog_ActionsLog) ProtoMessage()    {}
func (*TaskRet_TargetsLog_ActionsLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee839931d343e4e, []int{0, 0, 0}
}
func (m *TaskRet_TargetsLog_ActionsLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRet_TargetsLog_ActionsLog.Unmarshal(m, b)
}
func (m *TaskRet_TargetsLog_ActionsLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRet_TargetsLog_ActionsLog.Marshal(b, m, deterministic)
}
func (m *TaskRet_TargetsLog_ActionsLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRet_TargetsLog_ActionsLog.Merge(m, src)
}
func (m *TaskRet_TargetsLog_ActionsLog) XXX_Size() int {
	return xxx_messageInfo_TaskRet_TargetsLog_ActionsLog.Size(m)
}
func (m *TaskRet_TargetsLog_ActionsLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRet_TargetsLog_ActionsLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRet_TargetsLog_ActionsLog proto.InternalMessageInfo

func (m *TaskRet_TargetsLog_ActionsLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TaskRet_TargetsLog_ActionsLog) GetExtInfo() *types.Struct {
	if m != nil {
		return m.ExtInfo
	}
	return nil
}

func (m *TaskRet_TargetsLog_ActionsLog) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetProgress() *types.Struct {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *TaskRet_TargetsLog_ActionsLog) GetUsedTime() int32 {
	if m != nil {
		return m.UsedTime
	}
	return 0
}

func (m *TaskRet_TargetsLog_ActionsLog) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *TaskRet_TargetsLog_ActionsLog) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskRet)(nil), "easy_flow.TaskRet")
	proto.RegisterType((*TaskRet_TargetsLog)(nil), "easy_flow.TaskRet.TargetsLog")
	proto.RegisterType((*TaskRet_TargetsLog_ActionsLog)(nil), "easy_flow.TaskRet.TargetsLog.ActionsLog")
}

func init() { proto.RegisterFile("task_ret.proto", fileDescriptor_0ee839931d343e4e) }

var fileDescriptor_0ee839931d343e4e = []byte{
	// 1058 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xdb, 0x6e, 0x1b, 0xc5,
	0x1b, 0xff, 0x6f, 0x0e, 0xb6, 0x77, 0xdc, 0x9c, 0xa6, 0xfd, 0xab, 0xab, 0x08, 0x64, 0x33, 0x8d,
	0xd0, 0x9a, 0x32, 0xeb, 0x43, 0x28, 0xa2, 0xbe, 0xc0, 0x4a, 0x50, 0x11, 0x11, 0x28, 0x12, 0x13,
	0x8b, 0x8b, 0x2c, 0x4e, 0x34, 0xb1, 0xc7, 0x5b, 0xe3, 0xc3, 0xac, 0x76, 0xc7, 0x4d, 0xd3, 0x24,
	0x57, 0x48, 0x48, 0x3c, 0x05, 0xf7, 0x48, 0xdc, 0xf3, 0x00, 0x88, 0xc7, 0x70, 0x25, 0x24, 0x5e,
	0xc0, 0x2f, 0x00, 0x9a, 0x99, 0x3d, 0x91, 0x22, 0xd3, 0x0a, 0xf5, 0x02, 0xb8, 0xf2, 0xcc, 0xfc,
	0x7e, 0xbf, 0xd9, 0xef, 0xfb, 0xe6, 0x3b, 0x18, 0xac, 0x0b, 0x1a, 0x0e, 0x4f, 0x03, 0x26, 0x1c,
	0x3f, 0xe0, 0x82, 0x43, 0x93, 0xd1, 0xf0, 0xe2, 0xb4, 0x3f, 0xe2, 0xe7, 0xdb, 0xd8, 0x1b, 0x88,
	0xc7, 0xd3, 0x33, 0xa7, 0xcb, 0xc7, 0x55, 0x8f, 0x7b, 0xbc, 0xaa, 0x18, 0x67, 0xd3, 0xbe, 0xda,
	0xa9, 0x8d, 0x5a, 0x69, 0xe5, 0xf6, 0xfb, 0x19, 0xfa, 0xf8, 0x7c, 0x20, 0x86, 0xfc, 0xbc, 0xea,
	0x71, 0xac, 0x40, 0xfc, 0x84, 0x8e, 0x06, 0x3d, 0x2a, 0x78, 0x10, 0x56, 0x93, 0x65, 0xa4, 0x7b,
	0xc3, 0xe3, 0xdc, 0x1b, 0xb1, 0xf4, 0xf6, 0x50, 0x04, 0xd3, 0x6e, 0x64, 0x0f, 0xfa, 0xf1, 0x2e,
	0xc8, 0xb7, 0x69, 0x38, 0x24, 0x4c, 0xc0, 0x0a, 0xc8, 0x49, 0x6b, 0x0f, 0x7a, 0x96, 0x51, 0x36,
	0x6c, 0x73, 0x7f, 0x6b, 0x3e, 0x2b, 0xad, 0xf5, 0x79, 0x30, 0x6e, 0x22, 0x7d, 0x8e, 0x48, 0x44,
	0x80, 0xf7, 0xc0, 0xca, 0x84, 0x8e, 0x99, 0xb5, 0xa4, 0x88, 0x1b, 0xf3, 0x59, 0xa9, 0xa8, 0x89,
	0xf2, 0x14, 0x11, 0x05, 0xc2, 0xaf, 0x0d, 0x90, 0xf7, 0x02, 0x3e, 0xf5, 0x0f, 0x7a, 0xd6, 0xb2,
	0x22, 0x0e, 0xe6, 0xb3, 0xd2, 0xba, 0x26, 0x46, 0x00, 0xfa, 0xe5, 0x79, 0xe9, 0x08, 0x7c, 0x7e,
	0xe2, 0x52, 0xdc, 0xdf, 0xc3, 0x1f, 0xd7, 0xf0, 0xc3, 0xce, 0xe5, 0x07, 0xd7, 0xb8, 0x95, 0xdd,
	0xbf, 0xf7, 0x8a, 0xfb, 0x7a, 0xe3, 0x7a, 0x87, 0xc4, 0x5f, 0x86, 0x4d, 0xb0, 0x4a, 0x7d, 0x69,
	0xc2, 0x8a, 0x32, 0x61, 0x67, 0x3e, 0x2b, 0xdd, 0xd2, 0x26, 0xa8, 0x63, 0x69, 0xc0, 0x26, 0x58,
	0x3f, 0x71, 0x6b, 0xf8, 0x21, 0xc5, 0xcf, 0x3a, 0x97, 0xf5, 0xdd, 0xeb, 0x1d, 0xa2, 0x25, 0xb0,
	0x0a, 0x0a, 0x67, 0x54, 0x74, 0x1f, 0x1f, 0x4e, 0xc7, 0xd6, 0x6a, 0xd9, 0xb0, 0x57, 0xf7, 0x6f,
	0xcf, 0x67, 0xa5, 0x0d, 0x2d, 0x8f, 0x11, 0x44, 0x12, 0x12, 0xfc, 0x10, 0xac, 0xa9, 0xf5, 0xc1,
	0x44, 0xb0, 0xe0, 0x09, 0x1d, 0x59, 0x39, 0xa5, 0xb2, 0xe6, 0xb3, 0xd2, 0x9d, 0x8c, 0x2a, 0x86,
	0x11, 0xf9, 0x23, 0x1d, 0x3e, 0x00, 0xa0, 0x4f, 0x07, 0x23, 0xd6, 0x3b, 0x12, 0xdc, 0xb7, 0xf2,
	0xca, 0xe2, 0xff, 0xcf, 0x67, 0xa5, 0x2d, 0x2d, 0x4e, 0x31, 0x44, 0x32, 0x44, 0xf8, 0xad, 0x01,
	0x4c, 0x9f, 0x76, 0x87, 0xd4, 0x63, 0x07, 0x3d, 0xab, 0xa0, 0x64, 0xc3, 0xf9, 0xac, 0xb4, 0xa9,
	0x65, 0x09, 0xf4, 0xda, 0xa2, 0x9d, 0x7e, 0x5d, 0xa6, 0x86, 0xb8, 0xf0, 0x99, 0x65, 0xde, 0x4c,
	0x0d, 0x79, 0x8a, 0x88, 0x02, 0x61, 0x03, 0x98, 0xdc, 0x67, 0x01, 0x15, 0x03, 0x3e, 0xb1, 0x80,
	0x62, 0xde, 0x49, 0xed, 0x4d, 0x20, 0x44, 0x52, 0x1a, 0x7c, 0x1b, 0x2c, 0xf3, 0xc0, 0xb3, 0x8a,
	0x2a, 0xa2, 0x92, 0x0d, 0x22, 0x76, 0xe0, 0x49, 0xbf, 0x96, 0x36, 0xff, 0x47, 0x24, 0x01, 0x7e,
	0x01, 0x0a, 0x5a, 0xc4, 0x03, 0xeb, 0x96, 0xba, 0xba, 0x99, 0x3e, 0x5a, 0x8c, 0x48, 0xc5, 0x3d,
	0xf0, 0x96, 0x8c, 0xc4, 0xb3, 0x3d, 0x7c, 0x2c, 0x3d, 0x71, 0x9d, 0x64, 0x7d, 0x8a, 0x3b, 0x97,
	0x8d, 0x77, 0x77, 0xeb, 0xd7, 0x3b, 0x24, 0xb9, 0x4b, 0x96, 0x47, 0x28, 0xa8, 0x98, 0x86, 0xd6,
	0xda, 0xcd, 0xf2, 0xd0, 0xe7, 0x88, 0x44, 0x04, 0x19, 0x83, 0x2e, 0xef, 0x31, 0x6b, 0x5d, 0xd9,
	0x9a, 0x89, 0x81, 0x3c, 0x45, 0x44, 0x81, 0x32, 0xb9, 0xa6, 0x21, 0xeb, 0xb5, 0x07, 0x63, 0x66,
	0x6d, 0xdc, 0x4c, 0xae, 0x18, 0x41, 0x24, 0x21, 0xc1, 0x5f, 0x0d, 0x60, 0x86, 0x82, 0x06, 0x42,
	0x49, 0x36, 0x95, 0x11, 0x3f, 0x1b, 0x69, 0xd8, 0x12, 0x4c, 0x3a, 0xf7, 0x83, 0x01, 0xbe, 0x37,
	0x4e, 0x6c, 0xbb, 0xd5, 0x74, 0xeb, 0xd2, 0x39, 0xe9, 0xe1, 0x3b, 0x95, 0x96, 0x1b, 0xbd, 0x61,
	0x05, 0xdb, 0x75, 0xb7, 0x86, 0x1b, 0x9d, 0xab, 0x9a, 0xc2, 0x2b, 0xd8, 0xde, 0x75, 0x6b, 0xb8,
	0x1e, 0xef, 0xaf, 0xdc, 0x3a, 0x6e, 0x68, 0x55, 0xc5, 0x6d, 0x97, 0x3b, 0x76, 0xc3, 0xad, 0xe1,
	0xdd, 0xce, 0x95, 0xe2, 0xe8, 0xe3, 0xa6, 0xed, 0xd6, 0xf0, 0x83, 0x78, 0x93, 0xae, 0xed, 0x2f,
	0x1d, 0xf5, 0x7b, 0xbf, 0xd2, 0xb2, 0x8f, 0xaf, 0xdc, 0xfb, 0xb8, 0x63, 0xb7, 0x9a, 0x7f, 0x22,
	0xcf, 0xa8, 0x5b, 0x3b, 0x24, 0xf5, 0x0c, 0x3e, 0x37, 0x40, 0x9e, 0x4d, 0x74, 0x60, 0xb6, 0x94,
	0x97, 0x3f, 0x19, 0x69, 0xe3, 0x88, 0x90, 0x7f, 0xa2, 0x8f, 0xb1, 0x57, 0xf0, 0x53, 0x60, 0xb2,
	0xa7, 0x22, 0xa0, 0x07, 0x93, 0x3e, 0xb7, 0x60, 0xd9, 0xb0, 0x8b, 0x8d, 0xbb, 0x8e, 0xee, 0xd3,
	0x4e, 0xdc, 0xa7, 0x9d, 0x23, 0xd5, 0xa7, 0xb3, 0x75, 0x91, 0x68, 0x10, 0x49, 0xf5, 0xb0, 0x0d,
	0x80, 0xa0, 0x81, 0xc7, 0x44, 0xf8, 0x19, 0xf7, 0xac, 0xdb, 0xe5, 0x65, 0xbb, 0xd8, 0x78, 0xd3,
	0x49, 0xe6, 0x8c, 0x13, 0xb5, 0x77, 0xa7, 0x9d, 0x90, 0xb2, 0x2d, 0x25, 0x95, 0x22, 0x92, 0xb9,
	0x67, 0xfb, 0xbb, 0x22, 0x00, 0xa9, 0x02, 0x7e, 0x63, 0x80, 0x82, 0x46, 0x93, 0xf1, 0xf0, 0x55,
	0x9a, 0xad, 0x31, 0xf2, 0xda, 0xfa, 0x4b, 0xf2, 0x6d, 0xd9, 0x21, 0xf5, 0xfa, 0x30, 0x9d, 0x3f,
	0x2f, 0xb8, 0x73, 0xa8, 0xa6, 0x50, 0x86, 0x98, 0x29, 0xde, 0xe5, 0xbf, 0x2a, 0xde, 0x06, 0x30,
	0xc3, 0x8b, 0xf0, 0x48, 0xb3, 0x57, 0x6e, 0xf6, 0xa6, 0x04, 0x42, 0x24, 0xa5, 0x25, 0x05, 0xbf,
	0xba, 0xa8, 0xe0, 0xcb, 0x60, 0x79, 0x1c, 0x7a, 0x6a, 0x24, 0x98, 0xfb, 0xeb, 0x69, 0x03, 0x1b,
	0x87, 0x1e, 0x22, 0x12, 0x82, 0x9f, 0x80, 0x82, 0x1f, 0x70, 0x2f, 0x60, 0x61, 0xa8, 0x9a, 0xff,
	0x82, 0xb4, 0xc8, 0xf4, 0x8a, 0x58, 0x82, 0x48, 0xa2, 0x86, 0xa7, 0x00, 0xd0, 0xae, 0x6c, 0x9b,
	0x2a, 0x29, 0x0a, 0x2a, 0x29, 0xec, 0x85, 0x49, 0xe1, 0xec, 0x25, 0xfc, 0x6c, 0x40, 0xd3, 0x5b,
	0x10, 0xc9, 0x5c, 0xb9, 0xfd, 0x5b, 0x0e, 0x80, 0x54, 0x91, 0xfc, 0x21, 0x30, 0x16, 0xfd, 0x21,
	0x88, 0x47, 0xc3, 0xd2, 0xa2, 0xd1, 0x50, 0x01, 0x39, 0xfd, 0x99, 0x17, 0x5f, 0x4a, 0x9f, 0x23,
	0x12, 0x11, 0x32, 0x8f, 0xba, 0xf2, 0xb2, 0x1d, 0x79, 0xe1, 0x03, 0x3d, 0x02, 0x79, 0xf6, 0x54,
	0xa8, 0xa2, 0xcc, 0x2d, 0x8e, 0x3e, 0xcc, 0xf4, 0x23, 0xad, 0x40, 0x24, 0xd6, 0xc6, 0xef, 0x9c,
	0x7f, 0xb9, 0x77, 0x2e, 0xfc, 0xad, 0x77, 0xce, 0x0e, 0x11, 0xf3, 0xd5, 0x87, 0x08, 0xf8, 0x4f,
	0x0c, 0x91, 0xe2, 0xbf, 0x71, 0x88, 0xec, 0x3f, 0x3a, 0xfe, 0xc8, 0xe3, 0xaa, 0xa4, 0xb9, 0x1f,
	0x3a, 0x23, 0xde, 0xa5, 0xa3, 0x6a, 0x97, 0x4f, 0x44, 0x40, 0xbb, 0x22, 0xd4, 0x7f, 0xf7, 0x03,
	0xe6, 0x73, 0x3c, 0xe6, 0x3d, 0x36, 0x0a, 0xab, 0x11, 0xb1, 0xaa, 0xb6, 0xd5, 0xa4, 0x13, 0x9c,
	0xe5, 0x14, 0x73, 0xf7, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x17, 0xdf, 0x68, 0xaa, 0x0c,
	0x00, 0x00,
}
