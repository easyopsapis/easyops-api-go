// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flow_instance.proto

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
//流程执行结果
type FlowInstance struct {
	//
	//流程步骤详情
	StepList []*FlowExecuteStep `protobuf:"bytes,1,rep,name=stepList,proto3" json:"stepList" form:"stepList"`
	//
	//任务ID
	TaskId string `protobuf:"bytes,2,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//自动填充参数
	InstanceMap []*types.Struct `protobuf:"bytes,3,rep,name=instanceMap,proto3" json:"instanceMap" form:"instanceMap"`
	//
	//执行输出
	Outputs *types.Value `protobuf:"bytes,4,opt,name=outputs,proto3" json:"outputs" form:"outputs"`
	//
	//执行中步骤,如: {2: {startTime: 1563958843992, status: "executing", name: 2}}
	RunningSteps *types.Value `protobuf:"bytes,5,opt,name=runningSteps,proto3" json:"runningSteps" form:"runningSteps"`
	//
	//是否发送notify通知
	NeedNotify bool `protobuf:"varint,6,opt,name=needNotify,proto3" json:"needNotify" form:"needNotify"`
	//
	//开始时间戳毫秒
	StartTime int64 `protobuf:"varint,7,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//结束时间戳毫秒
	EndTime int64 `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	//
	//当前时间戳毫秒
	CurrentTime int64 `protobuf:"varint,9,opt,name=currentTime,proto3" json:"currentTime" form:"currentTime"`
	//
	//流程执行结果总状态("wait"-初始态等待执行, "executing"-执行中, "success"-成功, "failed"-失败, "paused"-人工确认, "interrupted"-暂停, "terminate"-终止)
	TotalStatus string `protobuf:"bytes,10,opt,name=totalStatus,proto3" json:"totalStatus" form:"totalStatus"`
	//
	//流程执行结果
	Message string `protobuf:"bytes,11,opt,name=message,proto3" json:"message" form:"message"`
	//
	//任务数量
	TaskCounter int32 `protobuf:"varint,12,opt,name=taskCounter,proto3" json:"taskCounter" form:"taskCounter"`
	//
	//流程输出数据
	FlowOutputsData *types.Value `protobuf:"bytes,13,opt,name=flowOutputsData,proto3" json:"flowOutputsData" form:"flowOutputsData"`
	//
	//表格数据
	TableData *types.Value `protobuf:"bytes,14,opt,name=tableData,proto3" json:"tableData" form:"tableData"`
	//
	//标准输出
	StandardOutputs *types.Value `protobuf:"bytes,15,opt,name=standardOutputs,proto3" json:"standardOutputs" form:"standardOutputs"`
	//
	//agent数据
	AgentData *types.Value `protobuf:"bytes,16,opt,name=agentData,proto3" json:"agentData" form:"agentData"`
	//
	//流程ID
	FlowId string `protobuf:"bytes,17,opt,name=flowId,proto3" json:"flowId" form:"flowId"`
	//
	//版本号(时间戳-秒)
	Version int32 `protobuf:"varint,18,opt,name=version,proto3" json:"version" form:"version"`
	//
	//流程输入
	//{
	//"flowInputs": {
	//"@agents": [
	//{
	//"instanceId": "5c6f6cf0d8079",
	//"ip": "192.168.100.162"
	//},
	//{
	//"instanceId": "5c6f6cf0d8075",
	//"ip": "192.168.100.163"
	//}
	//],
	//"must_input": "1",
	//"option_input": "000"
	//}
	//}
	//
	FlowInputs *types.Value `protobuf:"bytes,19,opt,name=flowInputs,proto3" json:"flowInputs" form:"flowInputs"`
	//
	//流程环境变量 { "youcash":"1234",
	//}
	FlowEnv *types.Value `protobuf:"bytes,20,opt,name=flowEnv,proto3" json:"flowEnv" form:"flowEnv"`
	//
	//流程环境类型
	Metadata *FlowInstance_Metadata `protobuf:"bytes,21,opt,name=metadata,proto3" json:"metadata" form:"metadata"`
	//
	//流程名称
	Name string `protobuf:"bytes,22,opt,name=name,proto3" json:"name" form:"name"`
	//
	//org
	Org int32 `protobuf:"varint,23,opt,name=org,proto3" json:"org" form:"org"`
	//
	//流程输出设置
	FlowOutputs []*FlowInstance_FlowOutputs `protobuf:"bytes,24,rep,name=flowOutputs,proto3" json:"flowOutputs" form:"flowOutputs"`
	//
	//输出定义
	OutputDefs []*FlowInstance_OutputDefs `protobuf:"bytes,25,rep,name=outputDefs,proto3" json:"outputDefs" form:"outputDefs"`
	//
	//表格输出定义
	TableDefs []*FlowInstance_TableDefs `protobuf:"bytes,26,rep,name=tableDefs,proto3" json:"tableDefs" form:"tableDefs"`
	//
	//流程创建者
	Creator string `protobuf:"bytes,27,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//流程类别如流水线、流程、流水线模板等
	Category string `protobuf:"bytes,28,opt,name=category,proto3" json:"category" form:"category"`
	//
	//修改时间
	UpdateTime string `protobuf:"bytes,29,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//创建时间
	CreateTime           string   `protobuf:"bytes,30,opt,name=createTime,proto3" json:"createTime" form:"createTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance) Reset()         { *m = FlowInstance{} }
func (m *FlowInstance) String() string { return proto.CompactTextString(m) }
func (*FlowInstance) ProtoMessage()    {}
func (*FlowInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0}
}
func (m *FlowInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance.Unmarshal(m, b)
}
func (m *FlowInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance.Marshal(b, m, deterministic)
}
func (m *FlowInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance.Merge(m, src)
}
func (m *FlowInstance) XXX_Size() int {
	return xxx_messageInfo_FlowInstance.Size(m)
}
func (m *FlowInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance proto.InternalMessageInfo

func (m *FlowInstance) GetStepList() []*FlowExecuteStep {
	if m != nil {
		return m.StepList
	}
	return nil
}

func (m *FlowInstance) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *FlowInstance) GetInstanceMap() []*types.Struct {
	if m != nil {
		return m.InstanceMap
	}
	return nil
}

func (m *FlowInstance) GetOutputs() *types.Value {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *FlowInstance) GetRunningSteps() *types.Value {
	if m != nil {
		return m.RunningSteps
	}
	return nil
}

func (m *FlowInstance) GetNeedNotify() bool {
	if m != nil {
		return m.NeedNotify
	}
	return false
}

func (m *FlowInstance) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *FlowInstance) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *FlowInstance) GetCurrentTime() int64 {
	if m != nil {
		return m.CurrentTime
	}
	return 0
}

func (m *FlowInstance) GetTotalStatus() string {
	if m != nil {
		return m.TotalStatus
	}
	return ""
}

func (m *FlowInstance) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FlowInstance) GetTaskCounter() int32 {
	if m != nil {
		return m.TaskCounter
	}
	return 0
}

func (m *FlowInstance) GetFlowOutputsData() *types.Value {
	if m != nil {
		return m.FlowOutputsData
	}
	return nil
}

func (m *FlowInstance) GetTableData() *types.Value {
	if m != nil {
		return m.TableData
	}
	return nil
}

func (m *FlowInstance) GetStandardOutputs() *types.Value {
	if m != nil {
		return m.StandardOutputs
	}
	return nil
}

func (m *FlowInstance) GetAgentData() *types.Value {
	if m != nil {
		return m.AgentData
	}
	return nil
}

func (m *FlowInstance) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FlowInstance) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FlowInstance) GetFlowInputs() *types.Value {
	if m != nil {
		return m.FlowInputs
	}
	return nil
}

func (m *FlowInstance) GetFlowEnv() *types.Value {
	if m != nil {
		return m.FlowEnv
	}
	return nil
}

func (m *FlowInstance) GetMetadata() *FlowInstance_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FlowInstance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlowInstance) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *FlowInstance) GetFlowOutputs() []*FlowInstance_FlowOutputs {
	if m != nil {
		return m.FlowOutputs
	}
	return nil
}

func (m *FlowInstance) GetOutputDefs() []*FlowInstance_OutputDefs {
	if m != nil {
		return m.OutputDefs
	}
	return nil
}

func (m *FlowInstance) GetTableDefs() []*FlowInstance_TableDefs {
	if m != nil {
		return m.TableDefs
	}
	return nil
}

func (m *FlowInstance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FlowInstance) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *FlowInstance) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *FlowInstance) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type FlowInstance_Metadata struct {
	//
	//环境
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//描述
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc" form:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance_Metadata) Reset()         { *m = FlowInstance_Metadata{} }
func (m *FlowInstance_Metadata) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_Metadata) ProtoMessage()    {}
func (*FlowInstance_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 0}
}
func (m *FlowInstance_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_Metadata.Unmarshal(m, b)
}
func (m *FlowInstance_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_Metadata.Marshal(b, m, deterministic)
}
func (m *FlowInstance_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_Metadata.Merge(m, src)
}
func (m *FlowInstance_Metadata) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_Metadata.Size(m)
}
func (m *FlowInstance_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_Metadata proto.InternalMessageInfo

func (m *FlowInstance_Metadata) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FlowInstance_Metadata) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type FlowInstance_FlowOutputs struct {
	//
	//流程输出参数设置
	Columns              []*FlowInstance_FlowOutputs_Columns `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FlowInstance_FlowOutputs) Reset()         { *m = FlowInstance_FlowOutputs{} }
func (m *FlowInstance_FlowOutputs) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_FlowOutputs) ProtoMessage()    {}
func (*FlowInstance_FlowOutputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 1}
}
func (m *FlowInstance_FlowOutputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_FlowOutputs.Unmarshal(m, b)
}
func (m *FlowInstance_FlowOutputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_FlowOutputs.Marshal(b, m, deterministic)
}
func (m *FlowInstance_FlowOutputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_FlowOutputs.Merge(m, src)
}
func (m *FlowInstance_FlowOutputs) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_FlowOutputs.Size(m)
}
func (m *FlowInstance_FlowOutputs) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_FlowOutputs.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_FlowOutputs proto.InternalMessageInfo

func (m *FlowInstance_FlowOutputs) GetColumns() []*FlowInstance_FlowOutputs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type FlowInstance_FlowOutputs_Columns struct {
	//
	//参数类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//参数ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//参数标题
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance_FlowOutputs_Columns) Reset()         { *m = FlowInstance_FlowOutputs_Columns{} }
func (m *FlowInstance_FlowOutputs_Columns) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_FlowOutputs_Columns) ProtoMessage()    {}
func (*FlowInstance_FlowOutputs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 1, 0}
}
func (m *FlowInstance_FlowOutputs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_FlowOutputs_Columns.Unmarshal(m, b)
}
func (m *FlowInstance_FlowOutputs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_FlowOutputs_Columns.Marshal(b, m, deterministic)
}
func (m *FlowInstance_FlowOutputs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_FlowOutputs_Columns.Merge(m, src)
}
func (m *FlowInstance_FlowOutputs_Columns) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_FlowOutputs_Columns.Size(m)
}
func (m *FlowInstance_FlowOutputs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_FlowOutputs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_FlowOutputs_Columns proto.InternalMessageInfo

func (m *FlowInstance_FlowOutputs_Columns) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FlowInstance_FlowOutputs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowInstance_FlowOutputs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FlowInstance_OutputDefs struct {
	//
	//参数类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//参数ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//参数标题
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance_OutputDefs) Reset()         { *m = FlowInstance_OutputDefs{} }
func (m *FlowInstance_OutputDefs) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_OutputDefs) ProtoMessage()    {}
func (*FlowInstance_OutputDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 2}
}
func (m *FlowInstance_OutputDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_OutputDefs.Unmarshal(m, b)
}
func (m *FlowInstance_OutputDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_OutputDefs.Marshal(b, m, deterministic)
}
func (m *FlowInstance_OutputDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_OutputDefs.Merge(m, src)
}
func (m *FlowInstance_OutputDefs) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_OutputDefs.Size(m)
}
func (m *FlowInstance_OutputDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_OutputDefs.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_OutputDefs proto.InternalMessageInfo

func (m *FlowInstance_OutputDefs) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FlowInstance_OutputDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowInstance_OutputDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FlowInstance_TableDefs struct {
	//
	//输出表格id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出表格标题
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列
	Dimensions []*FlowInstance_TableDefs_Dimensions `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//输出列
	Columns              []*FlowInstance_TableDefs_Columns `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FlowInstance_TableDefs) Reset()         { *m = FlowInstance_TableDefs{} }
func (m *FlowInstance_TableDefs) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_TableDefs) ProtoMessage()    {}
func (*FlowInstance_TableDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 3}
}
func (m *FlowInstance_TableDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_TableDefs.Unmarshal(m, b)
}
func (m *FlowInstance_TableDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_TableDefs.Marshal(b, m, deterministic)
}
func (m *FlowInstance_TableDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_TableDefs.Merge(m, src)
}
func (m *FlowInstance_TableDefs) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_TableDefs.Size(m)
}
func (m *FlowInstance_TableDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_TableDefs.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_TableDefs proto.InternalMessageInfo

func (m *FlowInstance_TableDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowInstance_TableDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlowInstance_TableDefs) GetDimensions() []*FlowInstance_TableDefs_Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *FlowInstance_TableDefs) GetColumns() []*FlowInstance_TableDefs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type FlowInstance_TableDefs_Dimensions struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance_TableDefs_Dimensions) Reset()         { *m = FlowInstance_TableDefs_Dimensions{} }
func (m *FlowInstance_TableDefs_Dimensions) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_TableDefs_Dimensions) ProtoMessage()    {}
func (*FlowInstance_TableDefs_Dimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 3, 0}
}
func (m *FlowInstance_TableDefs_Dimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_TableDefs_Dimensions.Unmarshal(m, b)
}
func (m *FlowInstance_TableDefs_Dimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_TableDefs_Dimensions.Marshal(b, m, deterministic)
}
func (m *FlowInstance_TableDefs_Dimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_TableDefs_Dimensions.Merge(m, src)
}
func (m *FlowInstance_TableDefs_Dimensions) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_TableDefs_Dimensions.Size(m)
}
func (m *FlowInstance_TableDefs_Dimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_TableDefs_Dimensions.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_TableDefs_Dimensions proto.InternalMessageInfo

func (m *FlowInstance_TableDefs_Dimensions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowInstance_TableDefs_Dimensions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FlowInstance_TableDefs_Columns struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowInstance_TableDefs_Columns) Reset()         { *m = FlowInstance_TableDefs_Columns{} }
func (m *FlowInstance_TableDefs_Columns) String() string { return proto.CompactTextString(m) }
func (*FlowInstance_TableDefs_Columns) ProtoMessage()    {}
func (*FlowInstance_TableDefs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_4330abd55e2ce5f3, []int{0, 3, 1}
}
func (m *FlowInstance_TableDefs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowInstance_TableDefs_Columns.Unmarshal(m, b)
}
func (m *FlowInstance_TableDefs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowInstance_TableDefs_Columns.Marshal(b, m, deterministic)
}
func (m *FlowInstance_TableDefs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowInstance_TableDefs_Columns.Merge(m, src)
}
func (m *FlowInstance_TableDefs_Columns) XXX_Size() int {
	return xxx_messageInfo_FlowInstance_TableDefs_Columns.Size(m)
}
func (m *FlowInstance_TableDefs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowInstance_TableDefs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_FlowInstance_TableDefs_Columns proto.InternalMessageInfo

func (m *FlowInstance_TableDefs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FlowInstance_TableDefs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*FlowInstance)(nil), "flow.FlowInstance")
	proto.RegisterType((*FlowInstance_Metadata)(nil), "flow.FlowInstance.Metadata")
	proto.RegisterType((*FlowInstance_FlowOutputs)(nil), "flow.FlowInstance.FlowOutputs")
	proto.RegisterType((*FlowInstance_FlowOutputs_Columns)(nil), "flow.FlowInstance.FlowOutputs.Columns")
	proto.RegisterType((*FlowInstance_OutputDefs)(nil), "flow.FlowInstance.OutputDefs")
	proto.RegisterType((*FlowInstance_TableDefs)(nil), "flow.FlowInstance.TableDefs")
	proto.RegisterType((*FlowInstance_TableDefs_Dimensions)(nil), "flow.FlowInstance.TableDefs.Dimensions")
	proto.RegisterType((*FlowInstance_TableDefs_Columns)(nil), "flow.FlowInstance.TableDefs.Columns")
}

func init() { proto.RegisterFile("flow_instance.proto", fileDescriptor_4330abd55e2ce5f3) }

var fileDescriptor_4330abd55e2ce5f3 = []byte{
	// 1262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdf, 0x6e, 0xdb, 0xb6,
	0x17, 0xfe, 0x29, 0x49, 0xf3, 0x87, 0x6e, 0x9b, 0x86, 0x69, 0x53, 0xd5, 0xfd, 0x63, 0xff, 0xd8,
	0xa0, 0x53, 0xd6, 0x4a, 0x76, 0x9c, 0x75, 0x58, 0x73, 0xb1, 0xb4, 0x6e, 0x1a, 0xa0, 0x43, 0xb3,
	0x76, 0x4c, 0x50, 0x0c, 0x75, 0xdd, 0x80, 0x91, 0x68, 0x4d, 0xa8, 0x2d, 0x1a, 0x12, 0xd5, 0x2c,
	0x6b, 0x7a, 0xb7, 0x47, 0xd9, 0xd5, 0x80, 0x3d, 0xc7, 0x1e, 0x61, 0x77, 0x1e, 0xb0, 0xab, 0x5d,
	0xfb, 0x09, 0x06, 0x52, 0x94, 0xc9, 0xa6, 0x81, 0x17, 0x60, 0xc5, 0x80, 0xdd, 0xc4, 0x22, 0xcf,
	0xf7, 0x7d, 0xfc, 0x78, 0x0e, 0x79, 0xa4, 0x80, 0xc5, 0x4e, 0x97, 0x1d, 0xec, 0x45, 0x71, 0xca,
	0x49, 0xec, 0x53, 0xaf, 0x9f, 0x30, 0xce, 0xe0, 0x94, 0x98, 0x2c, 0xbb, 0x61, 0xc4, 0xbf, 0xcb,
	0xf6, 0x3d, 0x9f, 0xf5, 0x6a, 0x21, 0x0b, 0x59, 0x4d, 0x06, 0xf7, 0xb3, 0x8e, 0x1c, 0xc9, 0x81,
	0x7c, 0xca, 0x49, 0xe5, 0xe7, 0x21, 0xf3, 0x28, 0x49, 0x0f, 0x59, 0x3f, 0xf5, 0xba, 0xcc, 0x27,
	0xdd, 0x9a, 0xcf, 0x62, 0x9e, 0x10, 0x9f, 0xa7, 0x39, 0x33, 0xa1, 0x7d, 0xe6, 0xf6, 0x58, 0x40,
	0xbb, 0x69, 0x4d, 0x01, 0x6b, 0x72, 0x58, 0x13, 0xcb, 0xc9, 0x3f, 0x7b, 0xf4, 0x7b, 0xea, 0x67,
	0x9c, 0xee, 0xa5, 0x9c, 0xf6, 0x95, 0xee, 0xe7, 0x86, 0x8d, 0xde, 0x41, 0xc4, 0x5f, 0xb3, 0x83,
	0x5a, 0xc8, 0x5c, 0x19, 0x74, 0xdf, 0x90, 0x6e, 0x14, 0x10, 0xce, 0x92, 0xb4, 0x36, 0x7a, 0x54,
	0xbc, 0x6b, 0x21, 0x63, 0x61, 0x97, 0x6a, 0xd7, 0x29, 0x4f, 0x32, 0x9f, 0xe7, 0x51, 0xf4, 0xe3,
	0x15, 0x70, 0x76, 0xab, 0xcb, 0x0e, 0x1e, 0xab, 0x9d, 0xc3, 0x2d, 0x30, 0x2b, 0x16, 0x7d, 0x12,
	0xa5, 0xdc, 0xb6, 0xaa, 0x93, 0x4e, 0xa9, 0x71, 0xc9, 0x13, 0x96, 0x3c, 0x81, 0x7a, 0x94, 0xdb,
	0xda, 0xe1, 0xb4, 0xdf, 0x5c, 0x1c, 0x0e, 0x2a, 0xf3, 0x1d, 0x96, 0xf4, 0xd6, 0x51, 0x41, 0x40,
	0x78, 0xc4, 0x85, 0x2b, 0x60, 0x9a, 0x93, 0xf4, 0xf5, 0xe3, 0xc0, 0x9e, 0xa8, 0x5a, 0xce, 0x5c,
	0x73, 0x61, 0x38, 0xa8, 0x9c, 0xcb, 0xe1, 0xf9, 0x3c, 0xc2, 0x0a, 0x00, 0xbf, 0x01, 0xa5, 0x22,
	0xf1, 0xdb, 0xa4, 0x6f, 0x4f, 0xca, 0x55, 0x2f, 0x7b, 0xb9, 0x6f, 0xaf, 0xf0, 0xed, 0xed, 0x48,
	0xdf, 0xcd, 0xa5, 0xe1, 0xa0, 0x02, 0x73, 0x21, 0x83, 0x85, 0xb0, 0xa9, 0x01, 0x37, 0xc1, 0x0c,
	0xcb, 0x78, 0x3f, 0xe3, 0xa9, 0x3d, 0x55, 0xb5, 0x9c, 0x52, 0x63, 0xe9, 0x03, 0xb9, 0xe7, 0xa4,
	0x9b, 0xd1, 0x26, 0x1c, 0x0e, 0x2a, 0xe7, 0x73, 0x35, 0x45, 0x40, 0xb8, 0xa0, 0xc2, 0x1d, 0x70,
	0x36, 0xc9, 0xe2, 0x38, 0x8a, 0x43, 0xb1, 0xe3, 0xd4, 0x3e, 0x33, 0x56, 0xea, 0xf2, 0x70, 0x50,
	0x59, 0xcc, 0xa5, 0x4c, 0x16, 0xc2, 0xef, 0x89, 0xc0, 0xbb, 0x00, 0xc4, 0x94, 0x06, 0x5f, 0x33,
	0x1e, 0x75, 0x0e, 0xed, 0xe9, 0xaa, 0xe5, 0xcc, 0x36, 0x2f, 0x0d, 0x07, 0x95, 0x85, 0x9c, 0xaa,
	0x63, 0x08, 0x1b, 0x40, 0xd8, 0x00, 0x73, 0x29, 0x27, 0x09, 0xdf, 0x8d, 0x7a, 0xd4, 0x9e, 0xa9,
	0x5a, 0xce, 0x64, 0xf3, 0xe2, 0x70, 0x50, 0xb9, 0x50, 0x54, 0x40, 0x85, 0x10, 0xd6, 0x30, 0x78,
	0x07, 0xcc, 0xd0, 0x38, 0x90, 0x8c, 0x59, 0xc9, 0x30, 0x76, 0xab, 0x02, 0x08, 0x17, 0x10, 0xf8,
	0x05, 0x28, 0xf9, 0x59, 0x92, 0xd0, 0x38, 0x5f, 0x63, 0x4e, 0x32, 0x8c, 0x6c, 0x1b, 0x41, 0x84,
	0x4d, 0xa8, 0x60, 0x72, 0xc6, 0x49, 0x77, 0x87, 0x13, 0x9e, 0xa5, 0x36, 0x90, 0x05, 0x37, 0x98,
	0x46, 0x10, 0x61, 0x13, 0x2a, 0x1c, 0xf6, 0x68, 0x9a, 0x92, 0x90, 0xda, 0x25, 0xc9, 0x32, 0x1c,
	0xaa, 0x00, 0xc2, 0x05, 0x44, 0xae, 0x43, 0xd2, 0xd7, 0x0f, 0x59, 0x16, 0x73, 0x9a, 0xd8, 0x67,
	0xab, 0x96, 0x73, 0xe6, 0xbd, 0x75, 0x74, 0x50, 0xac, 0xa3, 0x47, 0xf0, 0x25, 0x98, 0x17, 0x87,
	0xf8, 0x69, 0x5e, 0xd8, 0x4d, 0xc2, 0x89, 0x7d, 0x6e, 0x6c, 0x31, 0xcb, 0xc3, 0x41, 0x65, 0x29,
	0x57, 0x3d, 0x46, 0x44, 0xf8, 0xb8, 0x14, 0xfc, 0x0a, 0xcc, 0x71, 0xb2, 0xdf, 0xa5, 0x52, 0xf7,
	0xfc, 0x58, 0x5d, 0xa3, 0x66, 0x23, 0x0a, 0xc2, 0x9a, 0x2e, 0x9c, 0x8a, 0x63, 0x1c, 0x90, 0x24,
	0x50, 0x4b, 0xd8, 0xf3, 0xa7, 0x75, 0x7a, 0x8c, 0x88, 0xf0, 0x71, 0x29, 0xe1, 0x94, 0x84, 0x34,
	0xe6, 0xd2, 0xe9, 0x85, 0xd3, 0x3a, 0x1d, 0x51, 0x10, 0xd6, 0x74, 0x78, 0x1f, 0x4c, 0x8b, 0x44,
	0x3c, 0x0e, 0xec, 0x05, 0x59, 0x3a, 0x47, 0xdf, 0xf0, 0x7c, 0x1e, 0xfd, 0xf1, 0x7b, 0x65, 0x11,
	0x2c, 0xbc, 0x6a, 0x11, 0xb7, 0xf3, 0xc0, 0xdd, 0xaa, 0xbb, 0xf7, 0xda, 0x6f, 0xd7, 0x1a, 0xef,
	0x96, 0xb1, 0xe2, 0x89, 0xea, 0xbf, 0xa1, 0x49, 0x1a, 0xb1, 0xd8, 0x86, 0xb2, 0x96, 0x46, 0xf5,
	0x55, 0x00, 0xe1, 0x02, 0x02, 0xb7, 0x01, 0x90, 0xbc, 0x58, 0x26, 0x65, 0x71, 0xac, 0x79, 0xe3,
	0x42, 0x69, 0x0e, 0xc2, 0x86, 0x80, 0x68, 0x11, 0x62, 0xf4, 0x28, 0x7e, 0x63, 0x5f, 0x3c, 0x6d,
	0x8b, 0x50, 0x04, 0x84, 0x0b, 0x2a, 0x7c, 0x02, 0x66, 0x7b, 0x94, 0x93, 0x40, 0xe4, 0xf3, 0x92,
	0x94, 0xb9, 0xaa, 0xdb, 0x65, 0xd1, 0x54, 0xbd, 0x6d, 0x05, 0x31, 0x9b, 0x66, 0x41, 0x43, 0x78,
	0xa4, 0x00, 0x6f, 0x82, 0xa9, 0x98, 0xf4, 0xa8, 0xbd, 0x24, 0x13, 0x3a, 0x3f, 0x1c, 0x54, 0x4a,
	0xaa, 0x2b, 0x10, 0x71, 0xe9, 0x64, 0x10, 0xde, 0x02, 0x93, 0x2c, 0x09, 0xed, 0xcb, 0x32, 0x63,
	0xa2, 0x4a, 0x40, 0xf5, 0xaf, 0x24, 0x14, 0x19, 0x9f, 0xb8, 0xf0, 0x3f, 0x2c, 0x00, 0xf0, 0x5b,
	0x50, 0x32, 0x0e, 0xaa, 0x6d, 0xcb, 0xb6, 0x7a, 0xe3, 0x04, 0x77, 0x5b, 0x1a, 0x65, 0xde, 0x26,
	0x83, 0x8c, 0xb0, 0x29, 0x05, 0x77, 0x01, 0xc8, 0x5b, 0xe4, 0x26, 0xed, 0xa4, 0xf6, 0x15, 0x29,
	0x7c, 0xfd, 0x04, 0xe1, 0xa7, 0x23, 0x90, 0x59, 0x10, 0x4d, 0x45, 0xd8, 0xd0, 0x81, 0xcf, 0x8a,
	0x5b, 0x24, 0x44, 0xcb, 0x52, 0xf4, 0xda, 0x09, 0xa2, 0xbb, 0x05, 0xe6, 0xc3, 0xbb, 0x24, 0x25,
	0xb5, 0x08, 0xdc, 0x01, 0x33, 0x7e, 0x42, 0xc5, 0xbb, 0xd0, 0xbe, 0x2a, 0x33, 0x7a, 0x4f, 0x97,
	0x52, 0x05, 0x44, 0xc6, 0x6e, 0x82, 0xff, 0x8b, 0x33, 0xfa, 0xc3, 0x03, 0xf7, 0x85, 0x38, 0xa3,
	0x2d, 0x6f, 0xf4, 0xbc, 0xe7, 0xb6, 0xdf, 0x36, 0xee, 0xac, 0xad, 0xbe, 0x5b, 0xc6, 0x85, 0x12,
	0xac, 0x81, 0x59, 0x9f, 0x70, 0x1a, 0xb2, 0xe4, 0xd0, 0xbe, 0x26, 0x55, 0x8d, 0xa2, 0x16, 0x11,
	0x84, 0x47, 0x20, 0xf8, 0xa7, 0x05, 0x40, 0xd6, 0x0f, 0x08, 0xa7, 0xb2, 0xaf, 0x5e, 0x97, 0x9c,
	0x5f, 0x2d, 0x9d, 0x10, 0x1d, 0x14, 0x6e, 0x7e, 0xb1, 0xc0, 0xcf, 0xd6, 0x2b, 0xc7, 0xd9, 0x58,
	0x6f, 0xad, 0x0a, 0x37, 0xc2, 0xd2, 0xa7, 0x2b, 0x1b, 0xf2, 0xf7, 0xed, 0x67, 0xef, 0x56, 0x5c,
	0x67, 0xb5, 0x55, 0x77, 0x1b, 0xed, 0xa3, 0xba, 0x8c, 0xaf, 0xb8, 0xce, 0x5a, 0xab, 0xee, 0xae,
	0x16, 0xe3, 0xa3, 0xd6, 0xaa, 0xdb, 0xc8, 0x59, 0x2b, 0xad, 0xdd, 0x6a, 0xdb, 0x69, 0xb4, 0xea,
	0xee, 0x5a, 0xfb, 0x48, 0x62, 0xf2, 0xe9, 0x75, 0xa7, 0x55, 0x77, 0xef, 0x16, 0x03, 0xfd, 0xec,
	0xbc, 0xf4, 0xe4, 0xef, 0xed, 0x95, 0x0d, 0xe7, 0xc5, 0x51, 0xeb, 0xb6, 0xdb, 0x76, 0x36, 0xd6,
	0x4f, 0xa0, 0x1b, 0xec, 0x8d, 0x65, 0x6c, 0xec, 0x4d, 0x6e, 0x55, 0xe6, 0x29, 0xdf, 0xea, 0x8d,
	0x0f, 0xb6, 0xaa, 0x83, 0xff, 0xc9, 0xad, 0x6a, 0xfb, 0xe5, 0x5d, 0x30, 0xbb, 0x6d, 0x5c, 0x5b,
	0x7e, 0xd8, 0xa7, 0xb6, 0x75, 0xfc, 0xda, 0x8a, 0x59, 0x84, 0x65, 0x50, 0x80, 0x02, 0x9a, 0xfa,
	0xea, 0x73, 0xc8, 0x00, 0x89, 0x59, 0x84, 0x65, 0xb0, 0xfc, 0x9b, 0x05, 0x4a, 0x5b, 0xef, 0xdd,
	0xb4, 0x19, 0x9f, 0x75, 0xb3, 0x5e, 0x9c, 0xaa, 0x8f, 0xb1, 0x5b, 0xe3, 0xef, 0xaf, 0xf7, 0x30,
	0x47, 0x9b, 0x4d, 0x4b, 0x09, 0x20, 0x5c, 0x48, 0x95, 0x13, 0x30, 0xa3, 0x70, 0xa7, 0xb3, 0x7e,
	0x1d, 0x4c, 0x44, 0xc5, 0x77, 0xdc, 0xb9, 0xe1, 0xa0, 0x32, 0xa7, 0x3e, 0xbf, 0x02, 0x84, 0x27,
	0xa2, 0x60, 0xd4, 0xb5, 0x26, 0xc7, 0x74, 0xad, 0x72, 0x06, 0x80, 0x6e, 0x07, 0xff, 0xde, 0xb2,
	0x3f, 0x4d, 0x82, 0xb9, 0x51, 0xc7, 0x50, 0x8a, 0xd6, 0xdf, 0x29, 0x4e, 0x8c, 0x6b, 0xbf, 0xaf,
	0x00, 0x08, 0xa2, 0x1e, 0x8d, 0xc5, 0x3b, 0x29, 0x55, 0x1f, 0xab, 0x9f, 0x8c, 0xeb, 0x53, 0xde,
	0xe6, 0x08, 0x6e, 0xb6, 0x41, 0x2d, 0x82, 0xb0, 0xa1, 0x08, 0xb1, 0x2e, 0xf9, 0x94, 0x14, 0x5f,
	0x1e, 0x2b, 0x7e, 0xaa, 0x82, 0x3f, 0x03, 0x40, 0x9b, 0xf8, 0x18, 0x59, 0x28, 0x6f, 0xeb, 0x23,
	0xf4, 0x11, 0xe4, 0x9a, 0xf7, 0x5f, 0x7c, 0xf9, 0xcf, 0xfe, 0x6d, 0xda, 0x9f, 0x96, 0xa0, 0xb5,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xe8, 0xdb, 0xc3, 0xc9, 0x0d, 0x00, 0x00,
}
