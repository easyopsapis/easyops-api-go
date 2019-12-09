// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flow_execute_step.proto

package flow

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
//流程执行步骤
type FlowExecuteStep struct {
	//
	//流程执行的taskId
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//所在流程id
	FlowId string `protobuf:"bytes,2,opt,name=flowId,proto3" json:"flowId" form:"flowId"`
	//
	//执行时间信息
	Times *FlowExecuteStep_Times `protobuf:"bytes,3,opt,name=times,proto3" json:"times" form:"times"`
	//
	//步骤执行状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" form:"status"`
	//
	//org
	Org int32 `protobuf:"varint,5,opt,name=org,proto3" json:"org" form:"org"`
	//
	//执行输入定义, 当输入为空时是数组,非空时为map
	InputsDefinition *types.Value `protobuf:"bytes,6,opt,name=inputsDefinition,proto3" json:"inputsDefinition" form:"inputsDefinition"`
	//
	//执行输出定义,当输出定义为空时是数组，非空时是map
	OutputsDefinition *types.Value `protobuf:"bytes,7,opt,name=outputsDefinition,proto3" json:"outputsDefinition" form:"outputsDefinition"`
	//
	//步骤执行的taskId
	ExecId string `protobuf:"bytes,8,opt,name=execId,proto3" json:"execId" form:"execId"`
	//
	//步骤执行退出码
	ExitStatus int32 `protobuf:"varint,9,opt,name=exitStatus,proto3" json:"exitStatus" form:"exitStatus"`
	//
	//agent执行状态如: {"192.168.100.163": "success"}
	AgentStatus *types.Struct `protobuf:"bytes,10,opt,name=agentStatus,proto3" json:"agentStatus" form:"agentStatus"`
	//
	//agent执行输出如 {"192.168.100.163": null}
	AgentOutputs *types.Struct `protobuf:"bytes,11,opt,name=agentOutputs,proto3" json:"agentOutputs" form:"agentOutputs"`
	//
	//agent执行详细信息
	//{
	//"agentData": {
	//"192.168.100.163": {
	//"targetId": "5caee6c67d86a",
	//"tableData": [],
	//"outputs": null,
	//"msg": "Wed Jul 24 14:18:28 CST 2019\nWed Jul 24 14:18:29 CST 2019",
	//"startTime": "2019-07-24T14:18:28+08:00",
	//"updateTime": "2019-07-24T14:18:30+08:00",
	//"endTime": "2019-07-24T14:18:30+08:00",
	//"status": "success",
	//"exitStatus": 0,
	//"sysStatus": "ok",
	//"retryCount": 0,
	//"backoffLimit": 0
	//}
	//}
	//}
	//
	AgentData *types.Struct `protobuf:"bytes,12,opt,name=agentData,proto3" json:"agentData" form:"agentData"`
	//
	//步骤ID
	StepId int32 `protobuf:"varint,13,opt,name=stepId,proto3" json:"stepId" form:"stepId"`
	//
	//引用步骤ID(工具ID或者流程ID，当为人工审批节点是为空)
	ToolId string `protobuf:"bytes,14,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	//
	//流程审核示例
	//"actions":[
	//{
	//"label":"继续",
	//"value":"yes",
	//"terminate":false
	//},
	//{
	//"label":"停止",
	//"value":"no",
	//"terminate":true
	//},
	//]
	//
	Actions []*FlowExecuteStep_Actions `protobuf:"bytes,15,rep,name=actions,proto3" json:"actions" form:"actions"`
	//
	//审批用户或者用户组
	Approvers []string `protobuf:"bytes,16,rep,name=approvers,proto3" json:"approvers" form:"approvers"`
	//
	//邮件通知
	MailToApprovers bool `protobuf:"varint,17,opt,name=mailToApprovers,proto3" json:"mailToApprovers" form:"mailToApprovers"`
	//
	//步骤名
	StepName string `protobuf:"bytes,18,opt,name=stepName,proto3" json:"stepName" form:"stepName"`
	//
	//使用工具版本
	VId string `protobuf:"bytes,19,opt,name=vId,proto3" json:"vId" form:"vId"`
	//
	//执行用户
	ExecUser string `protobuf:"bytes,20,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//失败退出
	FailureExit bool `protobuf:"varint,21,opt,name=failureExit,proto3" json:"failureExit" form:"failureExit"`
	//
	//步骤类型(tool、flow、manual)
	SubType string `protobuf:"bytes,22,opt,name=subType,proto3" json:"subType" form:"subType"`
	//
	//类型
	Type string `protobuf:"bytes,23,opt,name=type,proto3" json:"type" form:"type"`
	//
	//在画布中X坐标
	XX float32 `protobuf:"fixed32,24,opt,name=_x,json=X,proto3" json:"_x" form:"_x"`
	//
	//在画布中Y坐标
	XY float32 `protobuf:"fixed32,25,opt,name=_y,json=Y,proto3" json:"_y" form:"_y"`
	//
	//是否root执行
	Root bool `protobuf:"varint,26,opt,name=root,proto3" json:"root" form:"root"`
	//
	//相对流程起点X坐标
	XDx float32 `protobuf:"fixed32,27,opt,name=_dx,json=Dx,proto3" json:"_dx" form:"_dx"`
	//
	//相对流程启动Y坐标
	XDy float32 `protobuf:"fixed32,28,opt,name=_dy,json=Dy,proto3" json:"_dy" form:"_dy"`
	//
	//样式
	Styles *FlowExecuteStep_Styles `protobuf:"bytes,29,opt,name=styles,proto3" json:"styles" form:"styles"`
	//
	//步骤输出, 有时候是{}, 有时候[]
	Outputs *types.Value `protobuf:"bytes,30,opt,name=outputs,proto3" json:"outputs" form:"outputs"`
	//
	//流程输入
	//{
	//"inputs": {
	//"must_input": "1",
	//"option_input": "000",
	//"@agents": [
	//{
	//"instanceId": "5c6f6cf0d8079",
	//"ip": "192.168.100.162"
	//},
	//{
	//"instanceId": "5c6f6cf0d8075",
	//"ip": "192.168.100.163"
	//}
	//]
	//}
	//}
	//
	Inputs *types.Value `protobuf:"bytes,31,opt,name=inputs,proto3" json:"inputs" form:"inputs"`
	//
	//通知条件 "OnSuccess", "OnFailure", "OnTimeout"
	InformCondition []string `protobuf:"bytes,32,rep,name=informCondition,proto3" json:"informCondition" form:"informCondition"`
	//
	//前置节点信息
	Parents              []*FlowExecuteStep_Parents `protobuf:"bytes,33,rep,name=parents,proto3" json:"parents" form:"parents"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FlowExecuteStep) Reset()         { *m = FlowExecuteStep{} }
func (m *FlowExecuteStep) String() string { return proto.CompactTextString(m) }
func (*FlowExecuteStep) ProtoMessage()    {}
func (*FlowExecuteStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb219dbf7ac088a, []int{0}
}
func (m *FlowExecuteStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowExecuteStep.Unmarshal(m, b)
}
func (m *FlowExecuteStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowExecuteStep.Marshal(b, m, deterministic)
}
func (m *FlowExecuteStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowExecuteStep.Merge(m, src)
}
func (m *FlowExecuteStep) XXX_Size() int {
	return xxx_messageInfo_FlowExecuteStep.Size(m)
}
func (m *FlowExecuteStep) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowExecuteStep.DiscardUnknown(m)
}

var xxx_messageInfo_FlowExecuteStep proto.InternalMessageInfo

func (m *FlowExecuteStep) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *FlowExecuteStep) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FlowExecuteStep) GetTimes() *FlowExecuteStep_Times {
	if m != nil {
		return m.Times
	}
	return nil
}

func (m *FlowExecuteStep) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowExecuteStep) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *FlowExecuteStep) GetInputsDefinition() *types.Value {
	if m != nil {
		return m.InputsDefinition
	}
	return nil
}

func (m *FlowExecuteStep) GetOutputsDefinition() *types.Value {
	if m != nil {
		return m.OutputsDefinition
	}
	return nil
}

func (m *FlowExecuteStep) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

func (m *FlowExecuteStep) GetExitStatus() int32 {
	if m != nil {
		return m.ExitStatus
	}
	return 0
}

func (m *FlowExecuteStep) GetAgentStatus() *types.Struct {
	if m != nil {
		return m.AgentStatus
	}
	return nil
}

func (m *FlowExecuteStep) GetAgentOutputs() *types.Struct {
	if m != nil {
		return m.AgentOutputs
	}
	return nil
}

func (m *FlowExecuteStep) GetAgentData() *types.Struct {
	if m != nil {
		return m.AgentData
	}
	return nil
}

func (m *FlowExecuteStep) GetStepId() int32 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *FlowExecuteStep) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func (m *FlowExecuteStep) GetActions() []*FlowExecuteStep_Actions {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *FlowExecuteStep) GetApprovers() []string {
	if m != nil {
		return m.Approvers
	}
	return nil
}

func (m *FlowExecuteStep) GetMailToApprovers() bool {
	if m != nil {
		return m.MailToApprovers
	}
	return false
}

func (m *FlowExecuteStep) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *FlowExecuteStep) GetVId() string {
	if m != nil {
		return m.VId
	}
	return ""
}

func (m *FlowExecuteStep) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *FlowExecuteStep) GetFailureExit() bool {
	if m != nil {
		return m.FailureExit
	}
	return false
}

func (m *FlowExecuteStep) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *FlowExecuteStep) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FlowExecuteStep) GetXX() float32 {
	if m != nil {
		return m.XX
	}
	return 0
}

func (m *FlowExecuteStep) GetXY() float32 {
	if m != nil {
		return m.XY
	}
	return 0
}

func (m *FlowExecuteStep) GetRoot() bool {
	if m != nil {
		return m.Root
	}
	return false
}

func (m *FlowExecuteStep) GetXDx() float32 {
	if m != nil {
		return m.XDx
	}
	return 0
}

func (m *FlowExecuteStep) GetXDy() float32 {
	if m != nil {
		return m.XDy
	}
	return 0
}

func (m *FlowExecuteStep) GetStyles() *FlowExecuteStep_Styles {
	if m != nil {
		return m.Styles
	}
	return nil
}

func (m *FlowExecuteStep) GetOutputs() *types.Value {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *FlowExecuteStep) GetInputs() *types.Value {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *FlowExecuteStep) GetInformCondition() []string {
	if m != nil {
		return m.InformCondition
	}
	return nil
}

func (m *FlowExecuteStep) GetParents() []*FlowExecuteStep_Parents {
	if m != nil {
		return m.Parents
	}
	return nil
}

type FlowExecuteStep_Times struct {
	//
	//步骤开始执行毫秒时间戳，如1563949106082
	StepStartTime int64 `protobuf:"varint,1,opt,name=stepStartTime,proto3" json:"stepStartTime" form:"stepStartTime"`
	//
	//工具开始执行毫秒时间戳，如1563949106082
	ToolStartTime int64 `protobuf:"varint,2,opt,name=toolStartTime,proto3" json:"toolStartTime" form:"toolStartTime"`
	//
	//工具结束执行毫秒时间戳，如1563949106082
	ToolEndTime int64 `protobuf:"varint,3,opt,name=toolEndTime,proto3" json:"toolEndTime" form:"toolEndTime"`
	//
	//步骤结束执行毫秒时间戳，如1563949106082
	StepEndTime          int64    `protobuf:"varint,4,opt,name=stepEndTime,proto3" json:"stepEndTime" form:"stepEndTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowExecuteStep_Times) Reset()         { *m = FlowExecuteStep_Times{} }
func (m *FlowExecuteStep_Times) String() string { return proto.CompactTextString(m) }
func (*FlowExecuteStep_Times) ProtoMessage()    {}
func (*FlowExecuteStep_Times) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb219dbf7ac088a, []int{0, 0}
}
func (m *FlowExecuteStep_Times) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowExecuteStep_Times.Unmarshal(m, b)
}
func (m *FlowExecuteStep_Times) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowExecuteStep_Times.Marshal(b, m, deterministic)
}
func (m *FlowExecuteStep_Times) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowExecuteStep_Times.Merge(m, src)
}
func (m *FlowExecuteStep_Times) XXX_Size() int {
	return xxx_messageInfo_FlowExecuteStep_Times.Size(m)
}
func (m *FlowExecuteStep_Times) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowExecuteStep_Times.DiscardUnknown(m)
}

var xxx_messageInfo_FlowExecuteStep_Times proto.InternalMessageInfo

func (m *FlowExecuteStep_Times) GetStepStartTime() int64 {
	if m != nil {
		return m.StepStartTime
	}
	return 0
}

func (m *FlowExecuteStep_Times) GetToolStartTime() int64 {
	if m != nil {
		return m.ToolStartTime
	}
	return 0
}

func (m *FlowExecuteStep_Times) GetToolEndTime() int64 {
	if m != nil {
		return m.ToolEndTime
	}
	return 0
}

func (m *FlowExecuteStep_Times) GetStepEndTime() int64 {
	if m != nil {
		return m.StepEndTime
	}
	return 0
}

type FlowExecuteStep_Actions struct {
	//
	//标题
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label" form:"label"`
	//
	//值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	//
	//是否中断
	Terminate            bool     `protobuf:"varint,3,opt,name=terminate,proto3" json:"terminate" form:"terminate"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowExecuteStep_Actions) Reset()         { *m = FlowExecuteStep_Actions{} }
func (m *FlowExecuteStep_Actions) String() string { return proto.CompactTextString(m) }
func (*FlowExecuteStep_Actions) ProtoMessage()    {}
func (*FlowExecuteStep_Actions) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb219dbf7ac088a, []int{0, 1}
}
func (m *FlowExecuteStep_Actions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowExecuteStep_Actions.Unmarshal(m, b)
}
func (m *FlowExecuteStep_Actions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowExecuteStep_Actions.Marshal(b, m, deterministic)
}
func (m *FlowExecuteStep_Actions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowExecuteStep_Actions.Merge(m, src)
}
func (m *FlowExecuteStep_Actions) XXX_Size() int {
	return xxx_messageInfo_FlowExecuteStep_Actions.Size(m)
}
func (m *FlowExecuteStep_Actions) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowExecuteStep_Actions.DiscardUnknown(m)
}

var xxx_messageInfo_FlowExecuteStep_Actions proto.InternalMessageInfo

func (m *FlowExecuteStep_Actions) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *FlowExecuteStep_Actions) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FlowExecuteStep_Actions) GetTerminate() bool {
	if m != nil {
		return m.Terminate
	}
	return false
}

type FlowExecuteStep_Styles struct {
	//
	//宽
	Width                float32  `protobuf:"fixed32,1,opt,name=width,proto3" json:"width" form:"width"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowExecuteStep_Styles) Reset()         { *m = FlowExecuteStep_Styles{} }
func (m *FlowExecuteStep_Styles) String() string { return proto.CompactTextString(m) }
func (*FlowExecuteStep_Styles) ProtoMessage()    {}
func (*FlowExecuteStep_Styles) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb219dbf7ac088a, []int{0, 2}
}
func (m *FlowExecuteStep_Styles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowExecuteStep_Styles.Unmarshal(m, b)
}
func (m *FlowExecuteStep_Styles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowExecuteStep_Styles.Marshal(b, m, deterministic)
}
func (m *FlowExecuteStep_Styles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowExecuteStep_Styles.Merge(m, src)
}
func (m *FlowExecuteStep_Styles) XXX_Size() int {
	return xxx_messageInfo_FlowExecuteStep_Styles.Size(m)
}
func (m *FlowExecuteStep_Styles) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowExecuteStep_Styles.DiscardUnknown(m)
}

var xxx_messageInfo_FlowExecuteStep_Styles proto.InternalMessageInfo

func (m *FlowExecuteStep_Styles) GetWidth() float32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type FlowExecuteStep_Parents struct {
	//
	//步骤名称
	StepName string `protobuf:"bytes,1,opt,name=stepName,proto3" json:"stepName" form:"stepName"`
	//
	//步骤序号
	StepId int32 `protobuf:"varint,2,opt,name=stepId,proto3" json:"stepId" form:"stepId"`
	//
	//连线起点位置
	SourcePoint string `protobuf:"bytes,3,opt,name=sourcePoint,proto3" json:"sourcePoint" form:"sourcePoint"`
	//
	//连线终点位置
	TargetPoint          string   `protobuf:"bytes,4,opt,name=targetPoint,proto3" json:"targetPoint" form:"targetPoint"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowExecuteStep_Parents) Reset()         { *m = FlowExecuteStep_Parents{} }
func (m *FlowExecuteStep_Parents) String() string { return proto.CompactTextString(m) }
func (*FlowExecuteStep_Parents) ProtoMessage()    {}
func (*FlowExecuteStep_Parents) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb219dbf7ac088a, []int{0, 3}
}
func (m *FlowExecuteStep_Parents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowExecuteStep_Parents.Unmarshal(m, b)
}
func (m *FlowExecuteStep_Parents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowExecuteStep_Parents.Marshal(b, m, deterministic)
}
func (m *FlowExecuteStep_Parents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowExecuteStep_Parents.Merge(m, src)
}
func (m *FlowExecuteStep_Parents) XXX_Size() int {
	return xxx_messageInfo_FlowExecuteStep_Parents.Size(m)
}
func (m *FlowExecuteStep_Parents) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowExecuteStep_Parents.DiscardUnknown(m)
}

var xxx_messageInfo_FlowExecuteStep_Parents proto.InternalMessageInfo

func (m *FlowExecuteStep_Parents) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *FlowExecuteStep_Parents) GetStepId() int32 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *FlowExecuteStep_Parents) GetSourcePoint() string {
	if m != nil {
		return m.SourcePoint
	}
	return ""
}

func (m *FlowExecuteStep_Parents) GetTargetPoint() string {
	if m != nil {
		return m.TargetPoint
	}
	return ""
}

func init() {
	proto.RegisterType((*FlowExecuteStep)(nil), "flow.FlowExecuteStep")
	proto.RegisterType((*FlowExecuteStep_Times)(nil), "flow.FlowExecuteStep.Times")
	proto.RegisterType((*FlowExecuteStep_Actions)(nil), "flow.FlowExecuteStep.Actions")
	proto.RegisterType((*FlowExecuteStep_Styles)(nil), "flow.FlowExecuteStep.Styles")
	proto.RegisterType((*FlowExecuteStep_Parents)(nil), "flow.FlowExecuteStep.Parents")
}

func init() { proto.RegisterFile("flow_execute_step.proto", fileDescriptor_9bb219dbf7ac088a) }

var fileDescriptor_9bb219dbf7ac088a = []byte{
	// 1236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x6f, 0x6f, 0x13, 0xc7,
	0x13, 0xfe, 0x9d, 0x9d, 0x7f, 0xde, 0x10, 0x92, 0x6c, 0x42, 0xb2, 0x3f, 0x13, 0x7a, 0x66, 0x89,
	0x90, 0x91, 0xb0, 0x0d, 0x49, 0x8b, 0x0a, 0x2f, 0x10, 0x31, 0x81, 0x2a, 0x42, 0xa2, 0x74, 0x93,
	0xa2, 0x16, 0x04, 0xd6, 0xc5, 0xb7, 0x31, 0x27, 0xce, 0x5e, 0xeb, 0x6e, 0x9d, 0xd8, 0xd0, 0x7c,
	0x84, 0x7e, 0xc0, 0xbe, 0xb9, 0x4a, 0xfd, 0x08, 0xf7, 0xb6, 0xaa, 0x54, 0xcd, 0xee, 0x9d, 0x6f,
	0x7d, 0xa6, 0x6e, 0xdf, 0xdd, 0xce, 0xf3, 0x3c, 0xb3, 0xb3, 0xb3, 0x33, 0xb3, 0x87, 0xb6, 0xcf,
	0x7c, 0x71, 0xd1, 0xe2, 0x43, 0xde, 0x1e, 0x48, 0xde, 0x0a, 0x25, 0xef, 0xd7, 0xfb, 0x81, 0x90,
	0x02, 0xcf, 0x01, 0x50, 0xae, 0x75, 0x3c, 0xf9, 0x61, 0x70, 0x5a, 0x6f, 0x8b, 0x6e, 0xa3, 0x23,
	0x3a, 0xa2, 0xa1, 0xc0, 0xd3, 0xc1, 0x99, 0x5a, 0xa9, 0x85, 0xfa, 0xd2, 0xa2, 0xf2, 0x03, 0x83,
	0xde, 0xbd, 0xf0, 0xe4, 0x47, 0x71, 0xd1, 0xe8, 0x88, 0x9a, 0x02, 0x6b, 0xe7, 0x8e, 0xef, 0xb9,
	0x8e, 0x14, 0x41, 0xd8, 0x18, 0x7f, 0x26, 0xba, 0x9d, 0x8e, 0x10, 0x1d, 0x9f, 0x67, 0xde, 0x43,
	0x19, 0x0c, 0xda, 0x52, 0xa3, 0xf4, 0xaf, 0x4d, 0xb4, 0xfa, 0xdc, 0x17, 0x17, 0xcf, 0x74, 0x94,
	0xc7, 0x92, 0xf7, 0xf1, 0x13, 0xb4, 0x20, 0x9d, 0xf0, 0xe3, 0x91, 0x4b, 0xac, 0x8a, 0x55, 0x2d,
	0x35, 0xab, 0x71, 0x64, 0xaf, 0x9c, 0x89, 0xa0, 0xfb, 0x88, 0x6a, 0x3b, 0xfd, 0xe3, 0x77, 0x7b,
	0x03, 0xad, 0xbf, 0x7f, 0xeb, 0xd4, 0xce, 0x0e, 0x6a, 0xcf, 0xef, 0xd5, 0x1e, 0xbe, 0xfb, 0xbc,
	0xbf, 0x77, 0xb9, 0xcb, 0x12, 0x1d, 0x78, 0x80, 0x23, 0x1e, 0xb9, 0xa4, 0x90, 0xf7, 0xa0, 0xed,
	0xff, 0xec, 0x41, 0xe3, 0xf8, 0x29, 0x9a, 0x97, 0x5e, 0x97, 0x87, 0xa4, 0x58, 0xb1, 0xaa, 0xcb,
	0x7b, 0xd7, 0xeb, 0x60, 0xaf, 0xe7, 0x22, 0xad, 0x9f, 0x00, 0xa5, 0xb9, 0x16, 0x47, 0xf6, 0x95,
	0x24, 0x3e, 0x30, 0x50, 0xa6, 0xb5, 0xf8, 0x0e, 0x5a, 0x08, 0xa5, 0x23, 0x07, 0x21, 0x99, 0x53,
	0x61, 0xac, 0x67, 0x61, 0x68, 0x3b, 0x65, 0x09, 0x01, 0xdf, 0x46, 0x45, 0x11, 0x74, 0xc8, 0x7c,
	0xc5, 0xaa, 0xce, 0x37, 0x37, 0xe3, 0xc8, 0x46, 0x9a, 0x27, 0x82, 0x0e, 0xc4, 0x5a, 0x58, 0xfb,
	0x1f, 0x03, 0x02, 0x6e, 0xa1, 0x35, 0xaf, 0xd7, 0x1f, 0xc8, 0xf0, 0x90, 0x9f, 0x79, 0x3d, 0x4f,
	0x7a, 0xa2, 0x47, 0x16, 0x54, 0x88, 0x5b, 0x75, 0x9d, 0xe8, 0x7a, 0x9a, 0xe8, 0xfa, 0x6b, 0xc7,
	0x1f, 0xf0, 0xe6, 0xf5, 0x38, 0xb2, 0xb7, 0xb5, 0xb3, 0xbc, 0x92, 0xb2, 0x29, 0x67, 0xf8, 0x14,
	0xad, 0x8b, 0x81, 0xcc, 0xed, 0xb0, 0x38, 0x73, 0x87, 0x9d, 0x38, 0xb2, 0x49, 0x12, 0x6e, 0x5e,
	0x4a, 0xd9, 0xb4, 0x3b, 0xc8, 0x0b, 0x54, 0xe5, 0x91, 0x4b, 0x96, 0xf2, 0x79, 0xd1, 0x76, 0xca,
	0x12, 0x02, 0xfe, 0x06, 0x21, 0x3e, 0xf4, 0xe4, 0xb1, 0x4e, 0x63, 0x49, 0xa5, 0xe7, 0x5a, 0x1c,
	0xd9, 0xeb, 0x29, 0x3d, 0xc5, 0x28, 0x33, 0x88, 0xf8, 0x07, 0xb4, 0xec, 0x74, 0x78, 0x2f, 0xd5,
	0x21, 0x15, 0xff, 0xf6, 0x54, 0xfc, 0xc7, 0xaa, 0x14, 0x9b, 0x5b, 0x71, 0x64, 0x63, 0xed, 0xd0,
	0x50, 0x51, 0x66, 0xfa, 0xc0, 0x27, 0xe8, 0x8a, 0x5a, 0x7e, 0xaf, 0x8f, 0x43, 0x96, 0x67, 0xfb,
	0xdc, 0x8e, 0x23, 0x7b, 0xc3, 0xf0, 0x99, 0xc8, 0x28, 0x9b, 0xf0, 0x82, 0x5f, 0xa0, 0x92, 0x5a,
	0x1f, 0x3a, 0xd2, 0x21, 0x57, 0x66, 0xbb, 0x84, 0xb2, 0x58, 0x33, 0x5c, 0x82, 0x86, 0xb2, 0x4c,
	0xaf, 0xeb, 0x8d, 0xf7, 0x8f, 0x5c, 0xb2, 0xa2, 0x12, 0x35, 0x51, 0x6f, 0x60, 0x57, 0xf5, 0x06,
	0x1f, 0x40, 0x95, 0x42, 0xf8, 0x47, 0x2e, 0xb9, 0x9a, 0xbf, 0x02, 0x6d, 0xa7, 0x2c, 0x21, 0xe0,
	0x17, 0x68, 0xd1, 0x69, 0xc3, 0xbd, 0x85, 0x64, 0xb5, 0x52, 0xac, 0x2e, 0xef, 0xdd, 0xf8, 0x72,
	0x33, 0x1c, 0x68, 0x52, 0x13, 0xc7, 0x91, 0x7d, 0x35, 0x09, 0x53, 0x9b, 0x28, 0x4b, 0x3d, 0x60,
	0x1f, 0x95, 0x9c, 0x7e, 0x3f, 0x10, 0xe7, 0x3c, 0x08, 0xc9, 0x5a, 0xa5, 0x58, 0x2d, 0x35, 0x5f,
	0x1a, 0xc7, 0x4a, 0x21, 0xa8, 0xf9, 0x07, 0xe8, 0xeb, 0x2a, 0x34, 0xe8, 0xa7, 0x83, 0xda, 0x1b,
	0x68, 0xd0, 0xec, 0xb3, 0x55, 0x7b, 0xf7, 0x79, 0xef, 0xee, 0xfe, 0xfd, 0xcb, 0x3b, 0xbb, 0xbf,
	0x54, 0xdf, 0x3f, 0x7a, 0x7b, 0xaf, 0xf6, 0xd0, 0xa9, 0x7d, 0x7a, 0xf7, 0xf9, 0xfe, 0xfe, 0xe5,
	0x9d, 0x5d, 0x96, 0x6d, 0x80, 0x0f, 0xd1, 0x6a, 0xd7, 0xf1, 0xfc, 0x13, 0x71, 0x30, 0xde, 0x73,
	0xbd, 0x62, 0x55, 0x97, 0x9a, 0xe5, 0x38, 0xb2, 0xb7, 0xf4, 0x9e, 0x39, 0x02, 0x65, 0x79, 0x09,
	0x6e, 0xa0, 0x25, 0xc8, 0xda, 0x4b, 0xa7, 0xcb, 0x09, 0x56, 0xd9, 0xda, 0x88, 0x23, 0x7b, 0x35,
	0x4b, 0x2c, 0x20, 0x94, 0x8d, 0x49, 0xb8, 0x82, 0x8a, 0xe7, 0x47, 0x2e, 0xd9, 0x50, 0xdc, 0xab,
	0x59, 0x33, 0x9f, 0x43, 0x5a, 0x01, 0xc2, 0xaf, 0xd1, 0x12, 0x14, 0xf8, 0x8f, 0x21, 0x0f, 0xc8,
	0xa6, 0xa2, 0x3d, 0xca, 0x5c, 0xa6, 0x08, 0x24, 0xe1, 0x16, 0xba, 0x39, 0x91, 0x83, 0xfa, 0x74,
	0x12, 0x76, 0xd9, 0xd8, 0x17, 0xfe, 0x16, 0x2d, 0x9f, 0x39, 0x9e, 0x3f, 0x08, 0xf8, 0xb3, 0xa1,
	0x27, 0xc9, 0x35, 0x75, 0x58, 0xa3, 0xbc, 0x0d, 0x90, 0x32, 0x93, 0x8a, 0xef, 0xa2, 0xc5, 0x70,
	0x70, 0x7a, 0x32, 0xea, 0x73, 0xb2, 0xa5, 0x02, 0x32, 0xae, 0x31, 0x01, 0x28, 0x4b, 0x29, 0xf8,
	0x16, 0x9a, 0x93, 0x40, 0xdd, 0x56, 0xd4, 0xd5, 0x38, 0xb2, 0x97, 0x93, 0xe2, 0x51, 0x3c, 0x05,
	0xe2, 0x1d, 0x54, 0x68, 0x0d, 0x09, 0xa9, 0x58, 0xd5, 0x42, 0x73, 0x25, 0x8e, 0xec, 0x92, 0xa6,
	0xb4, 0x86, 0x94, 0x59, 0x3f, 0x29, 0x74, 0x44, 0xfe, 0x3f, 0x85, 0x8e, 0x28, 0xb3, 0x7e, 0x86,
	0x0d, 0x02, 0x21, 0x24, 0x29, 0xab, 0x13, 0x18, 0x1b, 0x80, 0x95, 0x32, 0x05, 0x62, 0x1b, 0x15,
	0x5b, 0xee, 0x90, 0x5c, 0x57, 0x3e, 0x8c, 0x3c, 0xb7, 0xdc, 0x21, 0x65, 0x85, 0xc3, 0xa1, 0x26,
	0x8c, 0xc8, 0xce, 0x34, 0x61, 0x04, 0x84, 0x11, 0xfe, 0x0e, 0x3a, 0x66, 0xe4, 0xf3, 0x90, 0xdc,
	0x50, 0xbd, 0xb7, 0xf3, 0xe5, 0xd2, 0x3e, 0x56, 0x9c, 0xc9, 0x7e, 0x02, 0x8b, 0xea, 0x27, 0xf8,
	0xc0, 0x87, 0x68, 0x31, 0x99, 0x73, 0xe4, 0xab, 0x99, 0xc3, 0xd2, 0x48, 0xab, 0x48, 0x47, 0x42,
	0x2a, 0xc5, 0x07, 0x68, 0x41, 0x0f, 0x64, 0x62, 0xcf, 0x74, 0x62, 0x04, 0xa2, 0xf9, 0x94, 0x25,
	0x42, 0x28, 0x79, 0xaf, 0x07, 0xd0, 0x53, 0xd1, 0x73, 0xf5, 0xf4, 0xae, 0xa8, 0x36, 0x33, 0x4a,
	0x3e, 0x47, 0xa0, 0x2c, 0x2f, 0x81, 0x9e, 0xef, 0x3b, 0x01, 0xef, 0xc9, 0x90, 0xdc, 0x9c, 0xd5,
	0xf3, 0xaf, 0x34, 0xc9, 0x3c, 0x55, 0xa2, 0xa3, 0x2c, 0xf5, 0x50, 0xfe, 0xd3, 0x42, 0xf3, 0xea,
	0xa5, 0xc4, 0x8f, 0xd1, 0x0a, 0x34, 0xc9, 0xb1, 0x74, 0x02, 0x09, 0x16, 0xf5, 0xc0, 0x17, 0x9b,
	0x24, 0x8e, 0xec, 0xcd, 0xac, 0x9d, 0xc6, 0x30, 0x65, 0x93, 0x74, 0xd0, 0xc3, 0x50, 0xca, 0xf4,
	0x85, 0xbc, 0x7e, 0x02, 0xa6, 0x6c, 0x92, 0x0e, 0xed, 0x01, 0x86, 0x67, 0x3d, 0x57, 0xa9, 0x8b,
	0x4a, 0x6d, 0xb4, 0x87, 0x01, 0x52, 0x66, 0x52, 0x41, 0x09, 0xa1, 0xa4, 0xca, 0xb9, 0xbc, 0xd2,
	0x00, 0x29, 0x33, 0xa9, 0xe5, 0x5f, 0x2d, 0xb4, 0x98, 0x8c, 0x46, 0x7c, 0x1b, 0xcd, 0xfb, 0xce,
	0x29, 0xf7, 0x93, 0x1f, 0x1b, 0xe3, 0xc7, 0x41, 0x99, 0x29, 0xd3, 0x30, 0xf0, 0xce, 0xe1, 0xa2,
	0x93, 0xdf, 0x17, 0x83, 0xa7, 0xcc, 0x94, 0x69, 0x18, 0xef, 0xa1, 0x92, 0xe4, 0x41, 0xd7, 0xeb,
	0x39, 0x52, 0x9f, 0x66, 0xc9, 0x7c, 0x24, 0xc6, 0x10, 0x65, 0x19, 0xad, 0x7c, 0x0f, 0x2d, 0xe8,
	0x72, 0x86, 0x5d, 0x2e, 0x3c, 0x57, 0x7e, 0x50, 0xd1, 0x14, 0xcc, 0x5d, 0x94, 0x99, 0x32, 0x0d,
	0x97, 0x7f, 0xb3, 0xd0, 0x62, 0x72, 0xd1, 0x13, 0xb3, 0xd0, 0xfa, 0x2f, 0xb3, 0x30, 0x7b, 0x93,
	0x0a, 0xff, 0xf6, 0x26, 0x41, 0x8e, 0xc5, 0x20, 0x68, 0xf3, 0x57, 0xc2, 0xeb, 0x49, 0x75, 0x9e,
	0xd2, 0x44, 0x8e, 0x33, 0x10, 0x72, 0x9c, 0xad, 0xd4, 0xbd, 0x3a, 0x41, 0x87, 0x4b, 0xad, 0x9c,
	0xcb, 0x2b, 0x0d, 0x10, 0xee, 0x35, 0x5b, 0x35, 0x9f, 0xbc, 0x79, 0xdc, 0x11, 0x75, 0xee, 0x84,
	0x23, 0xd1, 0x0f, 0xeb, 0xbe, 0x68, 0x3b, 0x7e, 0xa3, 0x2d, 0x7a, 0x32, 0x70, 0xda, 0x32, 0xd4,
	0xff, 0xac, 0x01, 0xef, 0x8b, 0x5a, 0x57, 0xb8, 0xdc, 0x0f, 0x1b, 0x09, 0xb1, 0xa1, 0x96, 0x0d,
	0x68, 0x89, 0xd3, 0x05, 0x45, 0xda, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x89, 0xa4, 0x9e, 0x54,
	0x6e, 0x0b, 0x00, 0x00,
}