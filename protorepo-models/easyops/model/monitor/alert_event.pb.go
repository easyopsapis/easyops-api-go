// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alert_event.proto

package monitor

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
//告警事件 Model
type AlertEvent struct {
	//
	//记录 ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//告警 ID
	AlertId string `protobuf:"bytes,2,opt,name=alert_id,json=alertId,proto3" json:"alert_id" form:"alert_id"`
	//
	//告警规则 ID
	RuleId string `protobuf:"bytes,3,opt,name=rule_id,json=ruleId,proto3" json:"rule_id" form:"rule_id"`
	//
	//告警是否已恢复
	IsRecover bool `protobuf:"varint,4,opt,name=is_recover,json=isRecover,proto3" json:"is_recover" form:"is_recover"`
	//
	//告警是否发送成功
	SendSucc bool `protobuf:"varint,5,opt,name=send_succ,json=sendSucc,proto3" json:"send_succ" form:"send_succ"`
	//
	//告警主题
	Subject string `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject" form:"subject"`
	//
	//告警内容
	Content string `protobuf:"bytes,7,opt,name=content,proto3" json:"content" form:"content"`
	//
	//告警来源（system：easyops平台）
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source" form:"source"`
	//
	//告警状态: 0-告警已发送，1-告警被收敛，2-告警恢复，3-告警被屏蔽
	Status int32 `protobuf:"varint,9,opt,name=status,proto3" json:"status" form:"status"`
	//
	//是否发送邮件成功（0：未发送/失败，1：成功）
	SendDetail int32 `protobuf:"varint,10,opt,name=send_detail,json=sendDetail,proto3" json:"send_detail" form:"send_detail"`
	//
	//告警恢复类型（auto：自动恢复，manual：手动恢复）
	RecoverType string `protobuf:"bytes,11,opt,name=recover_type,json=recoverType,proto3" json:"recover_type" form:"recover_type"`
	//
	//org
	Org int32 `protobuf:"varint,12,opt,name=org,proto3" json:"org" form:"org"`
	//
	//告警对象
	Target string `protobuf:"bytes,13,opt,name=target,proto3" json:"target" form:"target"`
	//
	//告警等级
	Level int32 `protobuf:"varint,14,opt,name=level,proto3" json:"level" form:"level"`
	//
	//告警值
	Value int32 `protobuf:"varint,15,opt,name=value,proto3" json:"value" form:"value"`
	//
	//告警持续时间
	AlertDuration float32 `protobuf:"fixed32,16,opt,name=alert_duration,json=alertDuration,proto3" json:"alert_duration" form:"alert_duration"`
	//
	//告警开始时间
	AlertBeginTime int32 `protobuf:"varint,17,opt,name=alert_begin_time,json=alertBeginTime,proto3" json:"alert_begin_time" form:"alert_begin_time"`
	//
	//告警结束时间
	AlertEndTime int32 `protobuf:"varint,18,opt,name=alert_end_time,json=alertEndTime,proto3" json:"alert_end_time" form:"alert_end_time"`
	//
	//告警时间
	Time int32 `protobuf:"varint,19,opt,name=time,proto3" json:"time" form:"time"`
	//
	//告警开始时间
	StartTime int32 `protobuf:"varint,20,opt,name=start_time,json=startTime,proto3" json:"start_time" form:"start_time"`
	//
	//插入 mongodb 时间戳
	InsertTime int32 `protobuf:"varint,21,opt,name=insert_time,json=insertTime,proto3" json:"insert_time" form:"insert_time"`
	//
	//告警接收人
	AlertReceivers []*AlertEvent_AlertReceivers `protobuf:"bytes,22,rep,name=alert_receivers,json=alertReceivers,proto3" json:"alert_receivers" form:"alert_receivers"`
	//
	//告警维度
	AlertDims []*AlertEvent_AlertDims `protobuf:"bytes,23,rep,name=alert_dims,json=alertDims,proto3" json:"alert_dims" form:"alert_dims"`
	//
	//告警发生后的动作，可以是通知或者执行工具流程
	Actions []*AlertEvent_Actions `protobuf:"bytes,24,rep,name=actions,proto3" json:"actions" form:"actions"`
	//
	//告警条件
	AlertConditions *AlertConditions `protobuf:"bytes,25,opt,name=alert_conditions,json=alertConditions,proto3" json:"alert_conditions" form:"alert_conditions"`
	//
	//模型id
	ObjectId string `protobuf:"bytes,26,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例id
	InstanceId           string   `protobuf:"bytes,27,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertEvent) Reset()         { *m = AlertEvent{} }
func (m *AlertEvent) String() string { return proto.CompactTextString(m) }
func (*AlertEvent) ProtoMessage()    {}
func (*AlertEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0}
}
func (m *AlertEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent.Unmarshal(m, b)
}
func (m *AlertEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent.Marshal(b, m, deterministic)
}
func (m *AlertEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent.Merge(m, src)
}
func (m *AlertEvent) XXX_Size() int {
	return xxx_messageInfo_AlertEvent.Size(m)
}
func (m *AlertEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent proto.InternalMessageInfo

func (m *AlertEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AlertEvent) GetAlertId() string {
	if m != nil {
		return m.AlertId
	}
	return ""
}

func (m *AlertEvent) GetRuleId() string {
	if m != nil {
		return m.RuleId
	}
	return ""
}

func (m *AlertEvent) GetIsRecover() bool {
	if m != nil {
		return m.IsRecover
	}
	return false
}

func (m *AlertEvent) GetSendSucc() bool {
	if m != nil {
		return m.SendSucc
	}
	return false
}

func (m *AlertEvent) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *AlertEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *AlertEvent) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AlertEvent) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AlertEvent) GetSendDetail() int32 {
	if m != nil {
		return m.SendDetail
	}
	return 0
}

func (m *AlertEvent) GetRecoverType() string {
	if m != nil {
		return m.RecoverType
	}
	return ""
}

func (m *AlertEvent) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *AlertEvent) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *AlertEvent) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *AlertEvent) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *AlertEvent) GetAlertDuration() float32 {
	if m != nil {
		return m.AlertDuration
	}
	return 0
}

func (m *AlertEvent) GetAlertBeginTime() int32 {
	if m != nil {
		return m.AlertBeginTime
	}
	return 0
}

func (m *AlertEvent) GetAlertEndTime() int32 {
	if m != nil {
		return m.AlertEndTime
	}
	return 0
}

func (m *AlertEvent) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *AlertEvent) GetStartTime() int32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AlertEvent) GetInsertTime() int32 {
	if m != nil {
		return m.InsertTime
	}
	return 0
}

func (m *AlertEvent) GetAlertReceivers() []*AlertEvent_AlertReceivers {
	if m != nil {
		return m.AlertReceivers
	}
	return nil
}

func (m *AlertEvent) GetAlertDims() []*AlertEvent_AlertDims {
	if m != nil {
		return m.AlertDims
	}
	return nil
}

func (m *AlertEvent) GetActions() []*AlertEvent_Actions {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *AlertEvent) GetAlertConditions() *AlertConditions {
	if m != nil {
		return m.AlertConditions
	}
	return nil
}

func (m *AlertEvent) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AlertEvent) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type AlertEvent_AlertReceivers struct {
	//
	//接收人名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//通知方式, e.g.: email
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method" form:"method"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertEvent_AlertReceivers) Reset()         { *m = AlertEvent_AlertReceivers{} }
func (m *AlertEvent_AlertReceivers) String() string { return proto.CompactTextString(m) }
func (*AlertEvent_AlertReceivers) ProtoMessage()    {}
func (*AlertEvent_AlertReceivers) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0, 0}
}
func (m *AlertEvent_AlertReceivers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent_AlertReceivers.Unmarshal(m, b)
}
func (m *AlertEvent_AlertReceivers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent_AlertReceivers.Marshal(b, m, deterministic)
}
func (m *AlertEvent_AlertReceivers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent_AlertReceivers.Merge(m, src)
}
func (m *AlertEvent_AlertReceivers) XXX_Size() int {
	return xxx_messageInfo_AlertEvent_AlertReceivers.Size(m)
}
func (m *AlertEvent_AlertReceivers) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent_AlertReceivers.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent_AlertReceivers proto.InternalMessageInfo

func (m *AlertEvent_AlertReceivers) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertEvent_AlertReceivers) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type AlertEvent_AlertDims struct {
	//
	//告警维度前台展示名字
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//告警维度值
	Value                *types.Struct `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AlertEvent_AlertDims) Reset()         { *m = AlertEvent_AlertDims{} }
func (m *AlertEvent_AlertDims) String() string { return proto.CompactTextString(m) }
func (*AlertEvent_AlertDims) ProtoMessage()    {}
func (*AlertEvent_AlertDims) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0, 1}
}
func (m *AlertEvent_AlertDims) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent_AlertDims.Unmarshal(m, b)
}
func (m *AlertEvent_AlertDims) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent_AlertDims.Marshal(b, m, deterministic)
}
func (m *AlertEvent_AlertDims) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent_AlertDims.Merge(m, src)
}
func (m *AlertEvent_AlertDims) XXX_Size() int {
	return xxx_messageInfo_AlertEvent_AlertDims.Size(m)
}
func (m *AlertEvent_AlertDims) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent_AlertDims.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent_AlertDims proto.InternalMessageInfo

func (m *AlertEvent_AlertDims) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertEvent_AlertDims) GetValue() *types.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

type AlertEvent_Actions struct {
	//
	//动作触发条件
	Condition *AlertEvent_Actions_Condition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition" form:"condition"`
	//
	//动作类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	//
	//告警动作执行状态(0:未执行, 1:已执行)
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status" form:"status"`
	//
	//告警动作是否升级
	Upgrade bool `protobuf:"varint,4,opt,name=upgrade,proto3" json:"upgrade" form:"upgrade"`
	//
	//告警动作是否执行
	Run bool `protobuf:"varint,5,opt,name=run,proto3" json:"run" form:"run"`
	//
	//告警通知方法
	Method []string `protobuf:"bytes,6,rep,name=method,proto3" json:"method" form:"method"`
	//
	//告警通知人列表
	Receivers []string `protobuf:"bytes,7,rep,name=receivers,proto3" json:"receivers" form:"receivers"`
	//
	//告警通知用户组列表
	ReceiverUserGroups []string `protobuf:"bytes,8,rep,name=receiver_user_groups,json=receiverUserGroups,proto3" json:"receiver_user_groups" form:"receiver_user_groups"`
	//
	//告警通知运维负责人列表
	ReceiverOwners       []*AlertEvent_Actions_ReceiverOwners `protobuf:"bytes,9,rep,name=receiver_owners,json=receiverOwners,proto3" json:"receiver_owners" form:"receiver_owners"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AlertEvent_Actions) Reset()         { *m = AlertEvent_Actions{} }
func (m *AlertEvent_Actions) String() string { return proto.CompactTextString(m) }
func (*AlertEvent_Actions) ProtoMessage()    {}
func (*AlertEvent_Actions) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0, 2}
}
func (m *AlertEvent_Actions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent_Actions.Unmarshal(m, b)
}
func (m *AlertEvent_Actions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent_Actions.Marshal(b, m, deterministic)
}
func (m *AlertEvent_Actions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent_Actions.Merge(m, src)
}
func (m *AlertEvent_Actions) XXX_Size() int {
	return xxx_messageInfo_AlertEvent_Actions.Size(m)
}
func (m *AlertEvent_Actions) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent_Actions.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent_Actions proto.InternalMessageInfo

func (m *AlertEvent_Actions) GetCondition() *AlertEvent_Actions_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *AlertEvent_Actions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlertEvent_Actions) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AlertEvent_Actions) GetUpgrade() bool {
	if m != nil {
		return m.Upgrade
	}
	return false
}

func (m *AlertEvent_Actions) GetRun() bool {
	if m != nil {
		return m.Run
	}
	return false
}

func (m *AlertEvent_Actions) GetMethod() []string {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *AlertEvent_Actions) GetReceivers() []string {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *AlertEvent_Actions) GetReceiverUserGroups() []string {
	if m != nil {
		return m.ReceiverUserGroups
	}
	return nil
}

func (m *AlertEvent_Actions) GetReceiverOwners() []*AlertEvent_Actions_ReceiverOwners {
	if m != nil {
		return m.ReceiverOwners
	}
	return nil
}

type AlertEvent_Actions_Condition struct {
	//
	//告警持续时间多久后升级，单位：分钟
	LastingFor int32 `protobuf:"varint,1,opt,name=lasting_for,json=lastingFor,proto3" json:"lasting_for" form:"lasting_for"`
	//
	//告警等级大于多少升级
	Level                int32    `protobuf:"varint,2,opt,name=level,proto3" json:"level" form:"level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertEvent_Actions_Condition) Reset()         { *m = AlertEvent_Actions_Condition{} }
func (m *AlertEvent_Actions_Condition) String() string { return proto.CompactTextString(m) }
func (*AlertEvent_Actions_Condition) ProtoMessage()    {}
func (*AlertEvent_Actions_Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0, 2, 0}
}
func (m *AlertEvent_Actions_Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent_Actions_Condition.Unmarshal(m, b)
}
func (m *AlertEvent_Actions_Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent_Actions_Condition.Marshal(b, m, deterministic)
}
func (m *AlertEvent_Actions_Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent_Actions_Condition.Merge(m, src)
}
func (m *AlertEvent_Actions_Condition) XXX_Size() int {
	return xxx_messageInfo_AlertEvent_Actions_Condition.Size(m)
}
func (m *AlertEvent_Actions_Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent_Actions_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent_Actions_Condition proto.InternalMessageInfo

func (m *AlertEvent_Actions_Condition) GetLastingFor() int32 {
	if m != nil {
		return m.LastingFor
	}
	return 0
}

func (m *AlertEvent_Actions_Condition) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type AlertEvent_Actions_ReceiverOwners struct {
	//
	//描述
	Translate string `protobuf:"bytes,1,opt,name=translate,proto3" json:"translate" form:"translate"`
	//
	//模型 ID
	ObjectId string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//模型属性 ID
	ObjectAttrId         string   `protobuf:"bytes,3,opt,name=object_attr_id,json=objectAttrId,proto3" json:"object_attr_id" form:"object_attr_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertEvent_Actions_ReceiverOwners) Reset()         { *m = AlertEvent_Actions_ReceiverOwners{} }
func (m *AlertEvent_Actions_ReceiverOwners) String() string { return proto.CompactTextString(m) }
func (*AlertEvent_Actions_ReceiverOwners) ProtoMessage()    {}
func (*AlertEvent_Actions_ReceiverOwners) Descriptor() ([]byte, []int) {
	return fileDescriptor_783e30eadc53c1a4, []int{0, 2, 1}
}
func (m *AlertEvent_Actions_ReceiverOwners) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertEvent_Actions_ReceiverOwners.Unmarshal(m, b)
}
func (m *AlertEvent_Actions_ReceiverOwners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertEvent_Actions_ReceiverOwners.Marshal(b, m, deterministic)
}
func (m *AlertEvent_Actions_ReceiverOwners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertEvent_Actions_ReceiverOwners.Merge(m, src)
}
func (m *AlertEvent_Actions_ReceiverOwners) XXX_Size() int {
	return xxx_messageInfo_AlertEvent_Actions_ReceiverOwners.Size(m)
}
func (m *AlertEvent_Actions_ReceiverOwners) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertEvent_Actions_ReceiverOwners.DiscardUnknown(m)
}

var xxx_messageInfo_AlertEvent_Actions_ReceiverOwners proto.InternalMessageInfo

func (m *AlertEvent_Actions_ReceiverOwners) GetTranslate() string {
	if m != nil {
		return m.Translate
	}
	return ""
}

func (m *AlertEvent_Actions_ReceiverOwners) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AlertEvent_Actions_ReceiverOwners) GetObjectAttrId() string {
	if m != nil {
		return m.ObjectAttrId
	}
	return ""
}

func init() {
	proto.RegisterType((*AlertEvent)(nil), "monitor.AlertEvent")
	proto.RegisterType((*AlertEvent_AlertReceivers)(nil), "monitor.AlertEvent.AlertReceivers")
	proto.RegisterType((*AlertEvent_AlertDims)(nil), "monitor.AlertEvent.AlertDims")
	proto.RegisterType((*AlertEvent_Actions)(nil), "monitor.AlertEvent.Actions")
	proto.RegisterType((*AlertEvent_Actions_Condition)(nil), "monitor.AlertEvent.Actions.Condition")
	proto.RegisterType((*AlertEvent_Actions_ReceiverOwners)(nil), "monitor.AlertEvent.Actions.ReceiverOwners")
}

func init() { proto.RegisterFile("alert_event.proto", fileDescriptor_783e30eadc53c1a4) }

var fileDescriptor_783e30eadc53c1a4 = []byte{
	// 1240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xe1, 0x72, 0xd4, 0x36,
	0x10, 0xee, 0x5d, 0x20, 0x17, 0xeb, 0x92, 0x4b, 0x22, 0x02, 0x31, 0x97, 0x32, 0xbe, 0x51, 0x29,
	0x13, 0x5a, 0xee, 0x0e, 0x42, 0x07, 0xa6, 0xfc, 0xa1, 0x09, 0xd0, 0xce, 0xfd, 0xea, 0x54, 0xd0,
	0x99, 0x16, 0x06, 0x0e, 0xc7, 0x16, 0xc6, 0xad, 0xcf, 0xba, 0xca, 0x72, 0x98, 0x94, 0xe1, 0xbd,
	0xfa, 0x0e, 0xfd, 0xd3, 0x27, 0x70, 0x67, 0xfa, 0x08, 0x7e, 0x82, 0x8e, 0x56, 0xb2, 0xad, 0x3b,
	0x4a, 0xda, 0x99, 0xfe, 0xd3, 0xee, 0xf7, 0xed, 0x6a, 0xa5, 0x5d, 0xed, 0x0a, 0x6d, 0xfb, 0x09,
	0x13, 0x72, 0xca, 0x4e, 0x58, 0x2a, 0x47, 0x73, 0xc1, 0x25, 0xc7, 0x9d, 0x19, 0x4f, 0x63, 0xc9,
	0x45, 0x7f, 0x18, 0xc5, 0xf2, 0x75, 0x7e, 0x3c, 0x0a, 0xf8, 0x6c, 0x1c, 0xf1, 0x88, 0x8f, 0x01,
	0x3f, 0xce, 0x5f, 0x81, 0x04, 0x02, 0xac, 0xb4, 0x5d, 0xff, 0x87, 0x88, 0x8f, 0x98, 0x9f, 0x9d,
	0xf2, 0x79, 0x36, 0x4a, 0x78, 0xe0, 0x27, 0xe3, 0x80, 0xa7, 0x52, 0xf8, 0x81, 0xcc, 0xb4, 0xa5,
	0x60, 0x73, 0x3e, 0x9c, 0xf1, 0x90, 0x25, 0xd9, 0xd8, 0x10, 0xc7, 0x20, 0x8e, 0xcd, 0x8e, 0x63,
	0x1d, 0x4c, 0xc0, 0xd3, 0x30, 0x96, 0x31, 0x4f, 0x33, 0xe3, 0xf9, 0x8e, 0x15, 0xc8, 0xec, 0x4d,
	0x2c, 0x7f, 0xe6, 0x6f, 0xc6, 0x11, 0x1f, 0x02, 0x38, 0x3c, 0xf1, 0x93, 0x38, 0xf4, 0x25, 0x17,
	0xd9, 0xb8, 0x5e, 0x1a, 0xbb, 0x8f, 0x23, 0xce, 0xa3, 0x84, 0x35, 0x71, 0x67, 0x52, 0xe4, 0x81,
	0x39, 0x27, 0xf9, 0x7d, 0x07, 0xa1, 0x43, 0xb5, 0xe1, 0x23, 0x75, 0x78, 0x7c, 0x05, 0xb5, 0xe3,
	0xd0, 0x6d, 0x0d, 0x5a, 0xfb, 0xce, 0xd1, 0x46, 0x59, 0x78, 0xce, 0x2b, 0x2e, 0x66, 0xf7, 0x48,
	0x1c, 0x12, 0xda, 0x8e, 0x43, 0x3c, 0x42, 0x6b, 0x3a, 0xba, 0x38, 0x74, 0xdb, 0x40, 0xba, 0x50,
	0x16, 0xde, 0xa6, 0x26, 0x55, 0x08, 0xa1, 0x1d, 0x58, 0x4e, 0x42, 0xfc, 0x39, 0xea, 0x88, 0x3c,
	0x61, 0x8a, 0xbe, 0x02, 0x74, 0x5c, 0x16, 0x5e, 0x4f, 0xd3, 0x0d, 0x40, 0xe8, 0xaa, 0x5a, 0x4d,
	0x42, 0xfc, 0x05, 0x42, 0x71, 0x36, 0x15, 0x2c, 0xe0, 0x27, 0x4c, 0xb8, 0xe7, 0x06, 0xad, 0xfd,
	0xb5, 0xa3, 0x8b, 0x65, 0xe1, 0x6d, 0x9b, 0x18, 0x6a, 0x8c, 0x50, 0x27, 0xce, 0xa8, 0x5e, 0xe3,
	0x5b, 0xc8, 0xc9, 0x58, 0x1a, 0x4e, 0xb3, 0x3c, 0x08, 0xdc, 0xf3, 0x60, 0xb4, 0x53, 0x16, 0xde,
	0x96, 0x36, 0xaa, 0x21, 0x42, 0xd7, 0xd4, 0xfa, 0x71, 0x1e, 0x04, 0xf8, 0x06, 0xea, 0x64, 0xf9,
	0xf1, 0x4f, 0x2c, 0x90, 0xee, 0xea, 0x72, 0x54, 0x06, 0x20, 0xb4, 0xa2, 0x28, 0xb6, 0x4a, 0x21,
	0x4b, 0xa5, 0xdb, 0x59, 0x66, 0x1b, 0x80, 0xd0, 0x8a, 0x82, 0xaf, 0xa3, 0xd5, 0x8c, 0xe7, 0x22,
	0x60, 0xee, 0x1a, 0x90, 0xb7, 0xcb, 0xc2, 0xdb, 0x30, 0xae, 0x41, 0x4f, 0xa8, 0x21, 0x00, 0x55,
	0xfa, 0x32, 0xcf, 0x5c, 0x67, 0xd0, 0xda, 0x3f, 0xbf, 0x40, 0x05, 0xbd, 0xa2, 0xc2, 0x02, 0xdf,
	0x45, 0x5d, 0x38, 0x49, 0xc8, 0xa4, 0x1f, 0x27, 0x2e, 0x02, 0xfe, 0xa5, 0xb2, 0xf0, 0xb0, 0x75,
	0x4c, 0x0d, 0x12, 0x8a, 0x94, 0xf4, 0x10, 0x04, 0x7c, 0x0f, 0xad, 0x9b, 0x4b, 0x9b, 0xca, 0xd3,
	0x39, 0x73, 0xbb, 0x10, 0xd4, 0x6e, 0x59, 0x78, 0x17, 0x4c, 0x16, 0x2c, 0x94, 0xd0, 0xae, 0x11,
	0x9f, 0x9c, 0xce, 0x19, 0xbe, 0x86, 0x56, 0xb8, 0x88, 0xdc, 0x75, 0xd8, 0x4c, 0xdd, 0x29, 0xd2,
	0x26, 0x5c, 0x44, 0xe4, 0xaf, 0x3f, 0xbd, 0xf6, 0xd6, 0x47, 0x54, 0x11, 0xd4, 0x39, 0xa4, 0x2f,
	0x22, 0x26, 0xdd, 0x8d, 0xe5, 0x23, 0x6b, 0x3d, 0xa1, 0x86, 0x80, 0xaf, 0xa1, 0xf3, 0x09, 0x3b,
	0x61, 0x89, 0xdb, 0x03, 0xa7, 0x5b, 0x65, 0xe1, 0xad, 0x6b, 0x26, 0xa8, 0x09, 0xd5, 0xb0, 0xe2,
	0x9d, 0xf8, 0x49, 0xce, 0xdc, 0xcd, 0x65, 0x1e, 0xa8, 0x09, 0xd5, 0x30, 0xfe, 0x0a, 0xf5, 0x74,
	0xd5, 0x85, 0xb9, 0xf0, 0xd5, 0x63, 0x71, 0xb7, 0x06, 0xad, 0xfd, 0xf6, 0xd1, 0xe5, 0xb2, 0xf0,
	0x2e, 0xda, 0x55, 0x59, 0xe1, 0x84, 0x6e, 0x80, 0xe2, 0xa1, 0x91, 0xf1, 0x23, 0xb4, 0xa5, 0x19,
	0xc7, 0x2c, 0x8a, 0xd3, 0xa9, 0x8c, 0x67, 0xcc, 0xdd, 0x86, 0x4d, 0xf7, 0xca, 0xc2, 0xdb, 0xb5,
	0x7d, 0x34, 0x0c, 0x42, 0xf5, 0xb6, 0x47, 0x4a, 0xf3, 0x24, 0x9e, 0x31, 0x7c, 0xbf, 0x0a, 0x44,
	0x25, 0x02, 0x9c, 0x60, 0x70, 0xf2, 0x5e, 0x20, 0x15, 0x4e, 0xe8, 0x3a, 0x28, 0x1e, 0xa5, 0x21,
	0x38, 0xf8, 0x04, 0x9d, 0x03, 0xb3, 0x0b, 0x60, 0xb6, 0x59, 0x16, 0x5e, 0xd7, 0x5c, 0x21, 0x90,
	0x01, 0x54, 0x2f, 0x24, 0x93, 0xbe, 0x90, 0x7a, 0x87, 0x1d, 0xa0, 0x5a, 0x2f, 0xa4, 0xc1, 0x08,
	0x75, 0x40, 0x00, 0xd7, 0x77, 0x51, 0x37, 0x4e, 0x33, 0x56, 0x99, 0x5d, 0x5c, 0x2e, 0x1e, 0x0b,
	0x24, 0x14, 0x69, 0x09, 0x0c, 0x23, 0xb4, 0xa9, 0x83, 0x16, 0x2c, 0x60, 0xf1, 0x09, 0x13, 0x99,
	0x7b, 0x69, 0xb0, 0xb2, 0xdf, 0x3d, 0x20, 0x23, 0xd3, 0xab, 0x46, 0x4d, 0xeb, 0xd0, 0x4b, 0x5a,
	0x31, 0x8f, 0xfa, 0x65, 0xe1, 0x5d, 0xb2, 0x4f, 0x5e, 0x3b, 0xa9, 0x6e, 0xaf, 0xe6, 0xe2, 0xc7,
	0x08, 0x99, 0x34, 0xc5, 0xb3, 0xcc, 0xdd, 0x85, 0x3d, 0xae, 0x7c, 0x70, 0x8f, 0x87, 0xf1, 0x2c,
	0xb3, 0x8f, 0xdd, 0x98, 0x12, 0xea, 0xf8, 0x15, 0x03, 0x4f, 0x50, 0xc7, 0x0f, 0xa0, 0x81, 0xba,
	0x2e, 0x78, 0xdc, 0xfb, 0x47, 0x8f, 0x9a, 0x62, 0x3f, 0x6a, 0x63, 0xa5, 0xda, 0x98, 0x5e, 0xe1,
	0x97, 0x55, 0x91, 0x34, 0x4d, 0xd9, 0xbd, 0x3c, 0x68, 0xed, 0x77, 0x0f, 0xdc, 0x45, 0x9f, 0x0f,
	0x6a, 0xfc, 0xfd, 0xf2, 0x69, 0x6c, 0x09, 0xd5, 0xf7, 0xda, 0xb0, 0x31, 0x45, 0x6b, 0x1c, 0xda,
	0xcd, 0x24, 0x74, 0xfb, 0xf0, 0x8a, 0xee, 0x34, 0x8d, 0xb5, 0x42, 0xd4, 0xab, 0xf3, 0xd0, 0x95,
	0x17, 0xcf, 0xfc, 0xe1, 0xaf, 0x87, 0xc3, 0xa7, 0xd3, 0xe7, 0xcf, 0x6e, 0x0e, 0xbf, 0xac, 0xd6,
	0x6f, 0x6f, 0xde, 0xb8, 0x7d, 0xeb, 0xdd, 0x55, 0x5a, 0xfb, 0xc1, 0x13, 0xa4, 0x92, 0x29, 0xfd,
	0x34, 0x60, 0x93, 0xd0, 0xdd, 0x03, 0xaf, 0xd7, 0xad, 0x7e, 0x5a, 0x63, 0xca, 0xef, 0x16, 0xea,
	0xbd, 0x30, 0xee, 0x9e, 0xbf, 0xbd, 0x75, 0xfb, 0xdd, 0x55, 0x6a, 0x19, 0xf7, 0x5f, 0xa2, 0xde,
	0x62, 0x7a, 0x55, 0xbd, 0xa6, 0xfe, 0x8c, 0x99, 0x51, 0x61, 0xd5, 0xab, 0xd2, 0x12, 0x0a, 0xa0,
	0xea, 0x0c, 0x33, 0x26, 0x5f, 0xf3, 0x6a, 0x58, 0x58, 0x9d, 0x41, 0xeb, 0x09, 0x35, 0x84, 0xfe,
	0x2f, 0xc8, 0xa9, 0x93, 0xfb, 0xdf, 0x9c, 0xdf, 0xaf, 0x7a, 0x44, 0x1b, 0x32, 0xb1, 0x3b, 0xd2,
	0x73, 0x6e, 0x54, 0xcd, 0xb9, 0xd1, 0x63, 0x98, 0x73, 0x1f, 0x6c, 0x1e, 0xfd, 0x3f, 0x56, 0x51,
	0xc7, 0xa4, 0x1f, 0xff, 0x88, 0x9c, 0x3a, 0x3f, 0xb0, 0x6d, 0xf7, 0xe0, 0xd3, 0x33, 0xca, 0x65,
	0x54, 0xa7, 0xce, 0x1e, 0x36, 0xb5, 0x07, 0x42, 0x1b, 0x6f, 0xf0, 0xb2, 0x55, 0xeb, 0x6d, 0x2f,
	0x1f, 0x46, 0xb7, 0x5c, 0x00, 0xad, 0x59, 0xb0, 0xf2, 0x6f, 0xb3, 0xe0, 0x06, 0xea, 0xe4, 0xf3,
	0x48, 0xf8, 0x21, 0x33, 0x33, 0xd2, 0x2a, 0x5d, 0x03, 0x10, 0x5a, 0x51, 0xf0, 0x00, 0xad, 0x88,
	0x3c, 0x35, 0x83, 0xb1, 0xd7, 0x34, 0x71, 0x91, 0xa7, 0x84, 0x2a, 0xc8, 0x4a, 0xd2, 0xea, 0x60,
	0xe5, 0xcc, 0x24, 0xe1, 0x03, 0xe4, 0x34, 0xad, 0xa0, 0x03, 0x6c, 0xeb, 0xf8, 0xd6, 0x03, 0x6f,
	0x68, 0xf8, 0x3b, 0xb4, 0x53, 0x09, 0xd3, 0x3c, 0x63, 0x62, 0x1a, 0x09, 0x9e, 0xcf, 0x33, 0x77,
	0x0d, 0xcc, 0xbd, 0xb2, 0xf0, 0xf6, 0x16, 0xcd, 0x6d, 0x16, 0xa1, 0xb8, 0x52, 0x7f, 0x9f, 0x31,
	0xf1, 0x0d, 0x28, 0x31, 0x47, 0x9b, 0x35, 0x99, 0xbf, 0x49, 0x55, 0x30, 0x0e, 0xbc, 0xf0, 0xcf,
	0xce, 0x4a, 0x59, 0x55, 0xbb, 0xdf, 0x82, 0x85, 0xdd, 0x9f, 0x96, 0x9c, 0x11, 0xda, 0x13, 0x0b,
	0xdc, 0x7e, 0x82, 0x9c, 0x3a, 0xe1, 0xaa, 0x9d, 0x26, 0x7e, 0x26, 0xe3, 0x34, 0x9a, 0xbe, 0xe2,
	0x02, 0x8a, 0x65, 0xa1, 0x9d, 0x5a, 0x20, 0xa1, 0xc8, 0x48, 0x5f, 0x73, 0xd1, 0x0c, 0xbf, 0xf6,
	0x99, 0xc3, 0xaf, 0xff, 0x5b, 0x0b, 0xf5, 0x16, 0x83, 0x55, 0x17, 0x2f, 0x85, 0x9f, 0x66, 0x89,
	0x2f, 0xab, 0x57, 0x61, 0x5d, 0x7c, 0x0d, 0x11, 0xda, 0xd0, 0xd4, 0xc7, 0x48, 0xb7, 0x82, 0xe6,
	0xb3, 0x66, 0xd9, 0xd4, 0x10, 0xb1, 0x3a, 0xc6, 0x7d, 0xd4, 0x33, 0x7a, 0x5f, 0x4a, 0xd1, 0xfc,
	0xda, 0xac, 0x29, 0xb6, 0x88, 0x13, 0xba, 0xae, 0x15, 0x87, 0x52, 0x8a, 0x49, 0x78, 0xf4, 0xe0,
	0xe9, 0xe1, 0xff, 0xfe, 0xff, 0x1e, 0xaf, 0x02, 0xef, 0xf6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x11, 0xf8, 0x82, 0x08, 0x96, 0x0b, 0x00, 0x00,
}