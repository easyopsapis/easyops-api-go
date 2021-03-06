// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool_task.proto

package tool

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
//工具执行任务
type ToolTask struct {
	//
	//用户
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" form:"username"`
	//
	//工具输入参数
	Inputs *types.Struct `protobuf:"bytes,2,opt,name=inputs,proto3" json:"inputs" form:"inputs"`
	//
	//额外详细信息
	ExtraDetail *ToolTask_ExtraDetail `protobuf:"bytes,3,opt,name=extraDetail,proto3" json:"extraDetail" form:"extraDetail"`
	//
	//执行状态
	TotalStatus string `protobuf:"bytes,4,opt,name=totalStatus,proto3" json:"totalStatus" form:"totalStatus"`
	//
	//错误信息
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error" form:"error"`
	//
	//执行id
	ExecId string `protobuf:"bytes,6,opt,name=execId,proto3" json:"execId" form:"execId"`
	//
	//工具的Env
	ToolEnv *types.Struct `protobuf:"bytes,7,opt,name=toolEnv,proto3" json:"toolEnv" form:"toolEnv"`
	//
	//工具输出参数
	Outputs *types.Struct `protobuf:"bytes,8,opt,name=outputs,proto3" json:"outputs" form:"outputs"`
	//
	//outViewData
	OutViewData *types.Struct `protobuf:"bytes,9,opt,name=outViewData,proto3" json:"outViewData" form:"outViewData"`
	//
	//agent数据
	AgentData *types.Struct `protobuf:"bytes,10,opt,name=agentData,proto3" json:"agentData" form:"agentData"`
	//
	//agent列表
	Agents []string `protobuf:"bytes,11,rep,name=agents,proto3" json:"agents" form:"agents"`
	//
	//agents启动时间
	StartTime *types.Struct `protobuf:"bytes,12,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//agents状态
	Status *types.Struct `protobuf:"bytes,13,opt,name=status,proto3" json:"status" form:"status"`
	//
	//agents的msg
	Msg *types.Struct `protobuf:"bytes,14,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//agents结束时间
	EndTime *types.Struct `protobuf:"bytes,15,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	//
	//agents退出时间
	ExitStatus *types.Struct `protobuf:"bytes,16,opt,name=exitStatus,proto3" json:"exitStatus" form:"exitStatus"`
	//
	//agents系统状态
	SysStatus *types.Struct `protobuf:"bytes,17,opt,name=sysStatus,proto3" json:"sysStatus" form:"sysStatus"`
	//
	//工具输出数据
	ToolOutputsData []*types.Struct `protobuf:"bytes,18,rep,name=toolOutputsData,proto3" json:"toolOutputsData" form:"toolOutputsData"`
	//
	//表格数据
	TableData *types.Struct `protobuf:"bytes,19,opt,name=tableData,proto3" json:"tableData" form:"tableData"`
	//
	//工具信息
	ToolData *Tool `protobuf:"bytes,20,opt,name=toolData,proto3" json:"toolData" form:"toolData"`
	//
	//分批策略
	BatchStrategy *ToolTask_BatchStrategy `protobuf:"bytes,21,opt,name=batchStrategy,proto3" json:"batchStrategy" form:"batchStrategy"`
	//
	//工具执行的通知信息
	NeedNotify string `protobuf:"bytes,22,opt,name=needNotify,proto3" json:"needNotify" form:"needNotify"`
	//
	//执行用户
	ExecUser string `protobuf:"bytes,23,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//工具版本
	VId string `protobuf:"bytes,24,opt,name=vId,proto3" json:"vId" form:"vId"`
	//
	//工具Id，服务端自动生成
	ToolId string `protobuf:"bytes,25,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	//
	//工具的输出定义
	OutputDefs []*ToolTask_OutputDefs `protobuf:"bytes,26,rep,name=outputDefs,proto3" json:"outputDefs" form:"outputDefs"`
	//
	//工具的表格输出定义
	TableDefs []*ToolTask_TableDefs `protobuf:"bytes,27,rep,name=tableDefs,proto3" json:"tableDefs" form:"tableDefs"`
	//
	//工具的输出, 输出为空为[],非空时为object, 结构如下:
	//fields:
	//- name: dimensions
	//type: object[]
	//description: 维度列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//- name: columns
	//type: object[]
	//description: 输出列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//
	ToolOutputs          *types.Value `protobuf:"bytes,28,opt,name=toolOutputs,proto3" json:"toolOutputs" form:"toolOutputs"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ToolTask) Reset()         { *m = ToolTask{} }
func (m *ToolTask) String() string { return proto.CompactTextString(m) }
func (*ToolTask) ProtoMessage()    {}
func (*ToolTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0}
}
func (m *ToolTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask.Unmarshal(m, b)
}
func (m *ToolTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask.Marshal(b, m, deterministic)
}
func (m *ToolTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask.Merge(m, src)
}
func (m *ToolTask) XXX_Size() int {
	return xxx_messageInfo_ToolTask.Size(m)
}
func (m *ToolTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask proto.InternalMessageInfo

func (m *ToolTask) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ToolTask) GetInputs() *types.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ToolTask) GetExtraDetail() *ToolTask_ExtraDetail {
	if m != nil {
		return m.ExtraDetail
	}
	return nil
}

func (m *ToolTask) GetTotalStatus() string {
	if m != nil {
		return m.TotalStatus
	}
	return ""
}

func (m *ToolTask) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ToolTask) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

func (m *ToolTask) GetToolEnv() *types.Struct {
	if m != nil {
		return m.ToolEnv
	}
	return nil
}

func (m *ToolTask) GetOutputs() *types.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *ToolTask) GetOutViewData() *types.Struct {
	if m != nil {
		return m.OutViewData
	}
	return nil
}

func (m *ToolTask) GetAgentData() *types.Struct {
	if m != nil {
		return m.AgentData
	}
	return nil
}

func (m *ToolTask) GetAgents() []string {
	if m != nil {
		return m.Agents
	}
	return nil
}

func (m *ToolTask) GetStartTime() *types.Struct {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ToolTask) GetStatus() *types.Struct {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ToolTask) GetMsg() *types.Struct {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ToolTask) GetEndTime() *types.Struct {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ToolTask) GetExitStatus() *types.Struct {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func (m *ToolTask) GetSysStatus() *types.Struct {
	if m != nil {
		return m.SysStatus
	}
	return nil
}

func (m *ToolTask) GetToolOutputsData() []*types.Struct {
	if m != nil {
		return m.ToolOutputsData
	}
	return nil
}

func (m *ToolTask) GetTableData() *types.Struct {
	if m != nil {
		return m.TableData
	}
	return nil
}

func (m *ToolTask) GetToolData() *Tool {
	if m != nil {
		return m.ToolData
	}
	return nil
}

func (m *ToolTask) GetBatchStrategy() *ToolTask_BatchStrategy {
	if m != nil {
		return m.BatchStrategy
	}
	return nil
}

func (m *ToolTask) GetNeedNotify() string {
	if m != nil {
		return m.NeedNotify
	}
	return ""
}

func (m *ToolTask) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *ToolTask) GetVId() string {
	if m != nil {
		return m.VId
	}
	return ""
}

func (m *ToolTask) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func (m *ToolTask) GetOutputDefs() []*ToolTask_OutputDefs {
	if m != nil {
		return m.OutputDefs
	}
	return nil
}

func (m *ToolTask) GetTableDefs() []*ToolTask_TableDefs {
	if m != nil {
		return m.TableDefs
	}
	return nil
}

func (m *ToolTask) GetToolOutputs() *types.Value {
	if m != nil {
		return m.ToolOutputs
	}
	return nil
}

type ToolTask_ExtraDetail struct {
	//
	//工具执行成功的回调信息
	Callback *Callback `protobuf:"bytes,1,opt,name=callback,proto3" json:"callback" form:"callback"`
	//
	//工具的Env
	ToolEnv *types.Struct `protobuf:"bytes,2,opt,name=toolEnv,proto3" json:"toolEnv" form:"toolEnv"`
	//
	//工具的输出, 输出为空为[],非空时为object, 结构如下:
	//fields:
	//- name: dimensions
	//type: object[]
	//description: 维度列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//- name: columns
	//type: object[]
	//description: 输出列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//
	ToolOutputs *types.Value `protobuf:"bytes,3,opt,name=toolOutputs,proto3" json:"toolOutputs" form:"toolOutputs"`
	//
	//工具的表格输出定义
	TableDefs []*ToolTask_ExtraDetail_TableDefs `protobuf:"bytes,4,rep,name=tableDefs,proto3" json:"tableDefs" form:"tableDefs"`
	//
	//工具的输出定义
	OutputDefs           []*ToolTask_ExtraDetail_OutputDefs `protobuf:"bytes,5,rep,name=outputDefs,proto3" json:"outputDefs" form:"outputDefs"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ToolTask_ExtraDetail) Reset()         { *m = ToolTask_ExtraDetail{} }
func (m *ToolTask_ExtraDetail) String() string { return proto.CompactTextString(m) }
func (*ToolTask_ExtraDetail) ProtoMessage()    {}
func (*ToolTask_ExtraDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 0}
}
func (m *ToolTask_ExtraDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_ExtraDetail.Unmarshal(m, b)
}
func (m *ToolTask_ExtraDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_ExtraDetail.Marshal(b, m, deterministic)
}
func (m *ToolTask_ExtraDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_ExtraDetail.Merge(m, src)
}
func (m *ToolTask_ExtraDetail) XXX_Size() int {
	return xxx_messageInfo_ToolTask_ExtraDetail.Size(m)
}
func (m *ToolTask_ExtraDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_ExtraDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_ExtraDetail proto.InternalMessageInfo

func (m *ToolTask_ExtraDetail) GetCallback() *Callback {
	if m != nil {
		return m.Callback
	}
	return nil
}

func (m *ToolTask_ExtraDetail) GetToolEnv() *types.Struct {
	if m != nil {
		return m.ToolEnv
	}
	return nil
}

func (m *ToolTask_ExtraDetail) GetToolOutputs() *types.Value {
	if m != nil {
		return m.ToolOutputs
	}
	return nil
}

func (m *ToolTask_ExtraDetail) GetTableDefs() []*ToolTask_ExtraDetail_TableDefs {
	if m != nil {
		return m.TableDefs
	}
	return nil
}

func (m *ToolTask_ExtraDetail) GetOutputDefs() []*ToolTask_ExtraDetail_OutputDefs {
	if m != nil {
		return m.OutputDefs
	}
	return nil
}

type ToolTask_ExtraDetail_TableDefs struct {
	//
	//输出表格id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出表格标题
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列
	Dimensions []*ToolTask_ExtraDetail_TableDefs_Dimensions `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//输出列
	Columns              []*ToolTask_ExtraDetail_TableDefs_Columns `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ToolTask_ExtraDetail_TableDefs) Reset()         { *m = ToolTask_ExtraDetail_TableDefs{} }
func (m *ToolTask_ExtraDetail_TableDefs) String() string { return proto.CompactTextString(m) }
func (*ToolTask_ExtraDetail_TableDefs) ProtoMessage()    {}
func (*ToolTask_ExtraDetail_TableDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 0, 0}
}
func (m *ToolTask_ExtraDetail_TableDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs.Unmarshal(m, b)
}
func (m *ToolTask_ExtraDetail_TableDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs.Marshal(b, m, deterministic)
}
func (m *ToolTask_ExtraDetail_TableDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs.Merge(m, src)
}
func (m *ToolTask_ExtraDetail_TableDefs) XXX_Size() int {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs.Size(m)
}
func (m *ToolTask_ExtraDetail_TableDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_ExtraDetail_TableDefs proto.InternalMessageInfo

func (m *ToolTask_ExtraDetail_TableDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_ExtraDetail_TableDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToolTask_ExtraDetail_TableDefs) GetDimensions() []*ToolTask_ExtraDetail_TableDefs_Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *ToolTask_ExtraDetail_TableDefs) GetColumns() []*ToolTask_ExtraDetail_TableDefs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ToolTask_ExtraDetail_TableDefs_Dimensions struct {
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

func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) Reset() {
	*m = ToolTask_ExtraDetail_TableDefs_Dimensions{}
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) String() string { return proto.CompactTextString(m) }
func (*ToolTask_ExtraDetail_TableDefs_Dimensions) ProtoMessage()    {}
func (*ToolTask_ExtraDetail_TableDefs_Dimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 0, 0, 0}
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions.Unmarshal(m, b)
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions.Marshal(b, m, deterministic)
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions.Merge(m, src)
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) XXX_Size() int {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions.Size(m)
}
func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Dimensions proto.InternalMessageInfo

func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_ExtraDetail_TableDefs_Dimensions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ToolTask_ExtraDetail_TableDefs_Columns struct {
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

func (m *ToolTask_ExtraDetail_TableDefs_Columns) Reset() {
	*m = ToolTask_ExtraDetail_TableDefs_Columns{}
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) String() string { return proto.CompactTextString(m) }
func (*ToolTask_ExtraDetail_TableDefs_Columns) ProtoMessage()    {}
func (*ToolTask_ExtraDetail_TableDefs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 0, 0, 1}
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns.Unmarshal(m, b)
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns.Marshal(b, m, deterministic)
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns.Merge(m, src)
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) XXX_Size() int {
	return xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns.Size(m)
}
func (m *ToolTask_ExtraDetail_TableDefs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_ExtraDetail_TableDefs_Columns proto.InternalMessageInfo

func (m *ToolTask_ExtraDetail_TableDefs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_ExtraDetail_TableDefs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ToolTask_ExtraDetail_OutputDefs struct {
	//
	//输出参数ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出参数标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolTask_ExtraDetail_OutputDefs) Reset()         { *m = ToolTask_ExtraDetail_OutputDefs{} }
func (m *ToolTask_ExtraDetail_OutputDefs) String() string { return proto.CompactTextString(m) }
func (*ToolTask_ExtraDetail_OutputDefs) ProtoMessage()    {}
func (*ToolTask_ExtraDetail_OutputDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 0, 1}
}
func (m *ToolTask_ExtraDetail_OutputDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs.Unmarshal(m, b)
}
func (m *ToolTask_ExtraDetail_OutputDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs.Marshal(b, m, deterministic)
}
func (m *ToolTask_ExtraDetail_OutputDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs.Merge(m, src)
}
func (m *ToolTask_ExtraDetail_OutputDefs) XXX_Size() int {
	return xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs.Size(m)
}
func (m *ToolTask_ExtraDetail_OutputDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_ExtraDetail_OutputDefs proto.InternalMessageInfo

func (m *ToolTask_ExtraDetail_OutputDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_ExtraDetail_OutputDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ToolTask_BatchStrategy struct {
	//
	//自动分批:每批次部署机器数;手动分批:部署总共批次
	BatchNum int32 `protobuf:"varint,1,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//部署间隔时长(s)
	BatchInterval int32 `protobuf:"varint,2,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	//
	//是否失败暂停
	FailedStop           string   `protobuf:"bytes,3,opt,name=failedStop,proto3" json:"failedStop" form:"failedStop"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolTask_BatchStrategy) Reset()         { *m = ToolTask_BatchStrategy{} }
func (m *ToolTask_BatchStrategy) String() string { return proto.CompactTextString(m) }
func (*ToolTask_BatchStrategy) ProtoMessage()    {}
func (*ToolTask_BatchStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 1}
}
func (m *ToolTask_BatchStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_BatchStrategy.Unmarshal(m, b)
}
func (m *ToolTask_BatchStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_BatchStrategy.Marshal(b, m, deterministic)
}
func (m *ToolTask_BatchStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_BatchStrategy.Merge(m, src)
}
func (m *ToolTask_BatchStrategy) XXX_Size() int {
	return xxx_messageInfo_ToolTask_BatchStrategy.Size(m)
}
func (m *ToolTask_BatchStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_BatchStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_BatchStrategy proto.InternalMessageInfo

func (m *ToolTask_BatchStrategy) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *ToolTask_BatchStrategy) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

func (m *ToolTask_BatchStrategy) GetFailedStop() string {
	if m != nil {
		return m.FailedStop
	}
	return ""
}

type ToolTask_OutputDefs struct {
	//
	//输出参数ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出参数标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolTask_OutputDefs) Reset()         { *m = ToolTask_OutputDefs{} }
func (m *ToolTask_OutputDefs) String() string { return proto.CompactTextString(m) }
func (*ToolTask_OutputDefs) ProtoMessage()    {}
func (*ToolTask_OutputDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 2}
}
func (m *ToolTask_OutputDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_OutputDefs.Unmarshal(m, b)
}
func (m *ToolTask_OutputDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_OutputDefs.Marshal(b, m, deterministic)
}
func (m *ToolTask_OutputDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_OutputDefs.Merge(m, src)
}
func (m *ToolTask_OutputDefs) XXX_Size() int {
	return xxx_messageInfo_ToolTask_OutputDefs.Size(m)
}
func (m *ToolTask_OutputDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_OutputDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_OutputDefs proto.InternalMessageInfo

func (m *ToolTask_OutputDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_OutputDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ToolTask_TableDefs struct {
	//
	//输出表格id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出表格标题
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列
	Dimensions []*ToolTask_TableDefs_Dimensions `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//输出列
	Columns              []*ToolTask_TableDefs_Columns `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ToolTask_TableDefs) Reset()         { *m = ToolTask_TableDefs{} }
func (m *ToolTask_TableDefs) String() string { return proto.CompactTextString(m) }
func (*ToolTask_TableDefs) ProtoMessage()    {}
func (*ToolTask_TableDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 3}
}
func (m *ToolTask_TableDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_TableDefs.Unmarshal(m, b)
}
func (m *ToolTask_TableDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_TableDefs.Marshal(b, m, deterministic)
}
func (m *ToolTask_TableDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_TableDefs.Merge(m, src)
}
func (m *ToolTask_TableDefs) XXX_Size() int {
	return xxx_messageInfo_ToolTask_TableDefs.Size(m)
}
func (m *ToolTask_TableDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_TableDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_TableDefs proto.InternalMessageInfo

func (m *ToolTask_TableDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_TableDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToolTask_TableDefs) GetDimensions() []*ToolTask_TableDefs_Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *ToolTask_TableDefs) GetColumns() []*ToolTask_TableDefs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ToolTask_TableDefs_Dimensions struct {
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

func (m *ToolTask_TableDefs_Dimensions) Reset()         { *m = ToolTask_TableDefs_Dimensions{} }
func (m *ToolTask_TableDefs_Dimensions) String() string { return proto.CompactTextString(m) }
func (*ToolTask_TableDefs_Dimensions) ProtoMessage()    {}
func (*ToolTask_TableDefs_Dimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 3, 0}
}
func (m *ToolTask_TableDefs_Dimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_TableDefs_Dimensions.Unmarshal(m, b)
}
func (m *ToolTask_TableDefs_Dimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_TableDefs_Dimensions.Marshal(b, m, deterministic)
}
func (m *ToolTask_TableDefs_Dimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_TableDefs_Dimensions.Merge(m, src)
}
func (m *ToolTask_TableDefs_Dimensions) XXX_Size() int {
	return xxx_messageInfo_ToolTask_TableDefs_Dimensions.Size(m)
}
func (m *ToolTask_TableDefs_Dimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_TableDefs_Dimensions.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_TableDefs_Dimensions proto.InternalMessageInfo

func (m *ToolTask_TableDefs_Dimensions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_TableDefs_Dimensions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ToolTask_TableDefs_Columns struct {
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

func (m *ToolTask_TableDefs_Columns) Reset()         { *m = ToolTask_TableDefs_Columns{} }
func (m *ToolTask_TableDefs_Columns) String() string { return proto.CompactTextString(m) }
func (*ToolTask_TableDefs_Columns) ProtoMessage()    {}
func (*ToolTask_TableDefs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f86bed534fe5bc, []int{0, 3, 1}
}
func (m *ToolTask_TableDefs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTask_TableDefs_Columns.Unmarshal(m, b)
}
func (m *ToolTask_TableDefs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTask_TableDefs_Columns.Marshal(b, m, deterministic)
}
func (m *ToolTask_TableDefs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTask_TableDefs_Columns.Merge(m, src)
}
func (m *ToolTask_TableDefs_Columns) XXX_Size() int {
	return xxx_messageInfo_ToolTask_TableDefs_Columns.Size(m)
}
func (m *ToolTask_TableDefs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTask_TableDefs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTask_TableDefs_Columns proto.InternalMessageInfo

func (m *ToolTask_TableDefs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToolTask_TableDefs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ToolTask)(nil), "tool.ToolTask")
	proto.RegisterType((*ToolTask_ExtraDetail)(nil), "tool.ToolTask.ExtraDetail")
	proto.RegisterType((*ToolTask_ExtraDetail_TableDefs)(nil), "tool.ToolTask.ExtraDetail.TableDefs")
	proto.RegisterType((*ToolTask_ExtraDetail_TableDefs_Dimensions)(nil), "tool.ToolTask.ExtraDetail.TableDefs.Dimensions")
	proto.RegisterType((*ToolTask_ExtraDetail_TableDefs_Columns)(nil), "tool.ToolTask.ExtraDetail.TableDefs.Columns")
	proto.RegisterType((*ToolTask_ExtraDetail_OutputDefs)(nil), "tool.ToolTask.ExtraDetail.OutputDefs")
	proto.RegisterType((*ToolTask_BatchStrategy)(nil), "tool.ToolTask.BatchStrategy")
	proto.RegisterType((*ToolTask_OutputDefs)(nil), "tool.ToolTask.OutputDefs")
	proto.RegisterType((*ToolTask_TableDefs)(nil), "tool.ToolTask.TableDefs")
	proto.RegisterType((*ToolTask_TableDefs_Dimensions)(nil), "tool.ToolTask.TableDefs.Dimensions")
	proto.RegisterType((*ToolTask_TableDefs_Columns)(nil), "tool.ToolTask.TableDefs.Columns")
}

func init() { proto.RegisterFile("tool_task.proto", fileDescriptor_c0f86bed534fe5bc) }

var fileDescriptor_c0f86bed534fe5bc = []byte{
	// 1428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x58, 0xdf, 0x4e, 0x1b, 0xc7,
	0x17, 0x16, 0x98, 0x7f, 0x1e, 0x07, 0x08, 0x93, 0x84, 0x6c, 0xfc, 0xcb, 0x4f, 0xa6, 0x13, 0xa8,
	0xd6, 0x29, 0x6b, 0x63, 0x03, 0x49, 0x71, 0xa5, 0x90, 0x38, 0x24, 0x12, 0x6a, 0x43, 0xe8, 0x40,
	0xa3, 0x0a, 0xec, 0x44, 0x8b, 0x77, 0xec, 0xac, 0xb2, 0xf6, 0xa0, 0xdd, 0x31, 0x09, 0xb5, 0xfd,
	0x0c, 0x95, 0x7a, 0xdd, 0x07, 0xc8, 0x03, 0x54, 0xea, 0x75, 0xa5, 0x3e, 0x87, 0x2b, 0xf5, 0x11,
	0xfc, 0x04, 0xd5, 0xcc, 0xfe, 0x1b, 0x4c, 0xd6, 0x75, 0x04, 0xc9, 0x45, 0x6f, 0xac, 0x9d, 0x3d,
	0xe7, 0xfb, 0xe6, 0x9b, 0xb3, 0x67, 0xe6, 0x9c, 0x31, 0x98, 0x65, 0x94, 0x5a, 0xaf, 0x98, 0xee,
	0xbc, 0xc9, 0x1c, 0xdb, 0x94, 0x51, 0x38, 0xc6, 0x5f, 0x24, 0xb5, 0x9a, 0xc9, 0x5e, 0x37, 0x8f,
	0x32, 0x15, 0x5a, 0xcf, 0xd6, 0x68, 0x8d, 0x66, 0x85, 0xf1, 0xa8, 0x59, 0x15, 0x23, 0x31, 0x10,
	0x4f, 0x2e, 0x28, 0xf9, 0xac, 0x46, 0x33, 0x44, 0x77, 0x4e, 0xe9, 0xb1, 0x93, 0xb1, 0x68, 0x45,
	0xb7, 0xb2, 0x15, 0xda, 0x60, 0xb6, 0x5e, 0x61, 0x8e, 0x8b, 0xb4, 0xc9, 0x31, 0xd5, 0xea, 0xd4,
	0x20, 0x96, 0x93, 0xf5, 0x1c, 0xb3, 0x62, 0x98, 0xe5, 0xd3, 0x65, 0x2b, 0xba, 0x65, 0x1d, 0xe9,
	0x15, 0x4f, 0x43, 0x72, 0xfb, 0x82, 0x74, 0xfc, 0xc7, 0xa3, 0xba, 0x27, 0x2d, 0xa4, 0xfe, 0xd6,
	0x64, 0x6f, 0xe8, 0xdb, 0x6c, 0x8d, 0x6a, 0xc2, 0xa8, 0x9d, 0xe8, 0x96, 0x69, 0xe8, 0x8c, 0xda,
	0x4e, 0x36, 0x78, 0xf4, 0x70, 0xb7, 0x6b, 0x94, 0xd6, 0x2c, 0x12, 0xae, 0xdb, 0x61, 0x76, 0xb3,
	0xc2, 0x5c, 0x2b, 0xfa, 0x73, 0x09, 0x4c, 0xed, 0x53, 0x6a, 0xed, 0xeb, 0xce, 0x1b, 0xf8, 0x02,
	0x4c, 0x35, 0x1d, 0x62, 0x37, 0xf4, 0x3a, 0x51, 0x46, 0x16, 0x46, 0xd4, 0x78, 0xb1, 0xd0, 0xeb,
	0xa6, 0x66, 0xab, 0xd4, 0xae, 0x17, 0x90, 0x6f, 0x41, 0x7f, 0xff, 0x95, 0xba, 0x03, 0xbe, 0x78,
	0x79, 0xa8, 0x6b, 0x3f, 0x3d, 0xd2, 0x0e, 0x56, 0xb4, 0x8d, 0xf2, 0x61, 0x26, 0x78, 0x7e, 0xa5,
	0x95, 0x5b, 0xf9, 0xe5, 0xd5, 0x5c, 0x67, 0x11, 0x07, 0x5c, 0xb0, 0x08, 0x26, 0xcc, 0xc6, 0x71,
	0x93, 0x39, 0xca, 0xe8, 0xc2, 0x88, 0x9a, 0xc8, 0xdf, 0xcc, 0xb8, 0x9a, 0x32, 0xbe, 0xa6, 0xcc,
	0x9e, 0xd0, 0x54, 0x9c, 0xeb, 0x75, 0x53, 0xd3, 0xee, 0x74, 0x2e, 0x00, 0x61, 0x0f, 0x09, 0xf7,
	0x41, 0x82, 0xbc, 0x63, 0xb6, 0xbe, 0x45, 0x98, 0x6e, 0x5a, 0x4a, 0x4c, 0x10, 0x25, 0x33, 0x22,
	0x40, 0xfe, 0x02, 0x32, 0x4f, 0x42, 0x8f, 0xe2, 0x7c, 0xaf, 0x9b, 0x82, 0x2e, 0x97, 0x04, 0x44,
	0x58, 0xa6, 0x81, 0x5f, 0x83, 0x04, 0xa3, 0x4c, 0xb7, 0xf6, 0x98, 0xce, 0x9a, 0x8e, 0x32, 0x26,
	0x16, 0x2d, 0x21, 0x25, 0x23, 0xc2, 0xb2, 0x2b, 0xfc, 0x12, 0x8c, 0x13, 0xdb, 0xa6, 0xb6, 0x32,
	0x2e, 0x30, 0x57, 0x7b, 0xdd, 0xd4, 0x15, 0x6f, 0x36, 0xfe, 0x1a, 0x61, 0xd7, 0x0c, 0xd3, 0x60,
	0x82, 0xbc, 0x23, 0x95, 0x6d, 0x43, 0x99, 0x10, 0x8e, 0xd2, 0x12, 0xdd, 0xf7, 0x08, 0x7b, 0x0e,
	0xf0, 0x09, 0x98, 0xe4, 0xcb, 0x79, 0xd2, 0x38, 0x51, 0x26, 0x07, 0xc7, 0x09, 0xf6, 0xba, 0xa9,
	0x19, 0x5f, 0xa1, 0x40, 0x20, 0xec, 0x63, 0x39, 0x0d, 0x6d, 0x32, 0x11, 0xee, 0xa9, 0xa1, 0x69,
	0x3c, 0x04, 0xc2, 0x3e, 0x16, 0x7e, 0x0f, 0x12, 0xb4, 0xc9, 0x5e, 0x98, 0xe4, 0xed, 0x96, 0xce,
	0x74, 0x25, 0x3e, 0x98, 0x4a, 0x8a, 0x99, 0x84, 0x42, 0x58, 0xe6, 0x80, 0xdf, 0x82, 0xb8, 0x5e,
	0x23, 0x0d, 0x26, 0x08, 0xc1, 0x60, 0xc2, 0xeb, 0xbd, 0x6e, 0xea, 0xaa, 0x4b, 0x18, 0x60, 0x10,
	0x0e, 0xf1, 0xf0, 0x8f, 0x38, 0x98, 0x10, 0x23, 0x47, 0x49, 0x2c, 0xc4, 0xd4, 0x78, 0xf1, 0x7d,
	0x3c, 0x0c, 0xad, 0x6b, 0xe0, 0xa9, 0xfa, 0x4b, 0x1c, 0xfc, 0x1c, 0x7f, 0xa9, 0xaa, 0x79, 0x75,
	0xfd, 0x70, 0x45, 0x5b, 0x2f, 0xb7, 0x72, 0x9d, 0xf6, 0xe1, 0x8a, 0xb6, 0x56, 0x2e, 0x19, 0xad,
	0x5c, 0x27, 0xcd, 0x9f, 0x73, 0xe5, 0x4d, 0x3e, 0x58, 0xce, 0x77, 0xd2, 0x6a, 0x29, 0x33, 0xa4,
	0x67, 0xba, 0xb5, 0xda, 0x49, 0xb7, 0x4b, 0xce, 0x5d, 0x55, 0x55, 0x0f, 0x57, 0xb4, 0x8d, 0x47,
	0xda, 0x53, 0x5d, 0xab, 0x96, 0x5b, 0xb9, 0xe5, 0xb5, 0x4e, 0x21, 0xdd, 0xba, 0xdf, 0x39, 0xf7,
	0xb6, 0x5d, 0x48, 0xa7, 0xdb, 0x1f, 0x74, 0xbe, 0xd7, 0x51, 0x0b, 0xe7, 0xbc, 0x55, 0x35, 0xef,
	0xea, 0x68, 0xe7, 0x3d, 0x15, 0xed, 0x5c, 0xc9, 0x28, 0x19, 0xed, 0xc3, 0x9c, 0xb6, 0xc1, 0x75,
	0xb8, 0x62, 0xff, 0xc5, 0xc7, 0x95, 0x19, 0x39, 0xf3, 0x7a, 0x47, 0x55, 0xcf, 0xcf, 0x9d, 0x76,
	0x97, 0xd8, 0x2e, 0x7c, 0x16, 0x0d, 0x6b, 0x91, 0x1a, 0x38, 0xec, 0x43, 0xa6, 0xcd, 0xcb, 0x14,
	0x36, 0x40, 0xd9, 0x6a, 0xa4, 0xb2, 0xb5, 0x08, 0x65, 0xad, 0x95, 0xe5, 0x7c, 0xe7, 0x33, 0xa9,
	0xcb, 0x47, 0xaa, 0x5b, 0x8f, 0x56, 0xb7, 0xfa, 0xb9, 0xd4, 0xe5, 0x22, 0xd5, 0xdd, 0x8b, 0x56,
	0xb7, 0xf6, 0x29, 0xd4, 0x15, 0xa2, 0x84, 0xdc, 0x8f, 0x16, 0xb2, 0x7e, 0xf9, 0x42, 0xd2, 0xea,
	0x52, 0xe6, 0xab, 0xf4, 0x66, 0xc9, 0xb9, 0xbb, 0x88, 0xbd, 0x83, 0x8b, 0x1f, 0x88, 0x0e, 0xd3,
	0x6d, 0xb6, 0x6f, 0xd6, 0x89, 0x72, 0x65, 0xe8, 0x03, 0x31, 0xc0, 0x20, 0x1c, 0xe2, 0x79, 0x95,
	0x75, 0xdc, 0x32, 0x36, 0x3d, 0x74, 0x95, 0x75, 0xbc, 0xd2, 0xe6, 0x21, 0xe1, 0x06, 0x88, 0xd5,
	0x9d, 0x9a, 0x32, 0x33, 0x98, 0x60, 0xa6, 0xd7, 0x4d, 0x01, 0x97, 0xa0, 0xee, 0xd4, 0x10, 0xe6,
	0x18, 0x5e, 0x76, 0x48, 0xc3, 0x10, 0x2b, 0x99, 0x1d, 0xba, 0xec, 0x78, 0x08, 0x84, 0x7d, 0x2c,
	0xdc, 0x01, 0x80, 0xbc, 0x33, 0x99, 0x57, 0x90, 0xaf, 0x0e, 0x66, 0xba, 0xd1, 0xeb, 0xa6, 0xe6,
	0xfc, 0x62, 0xea, 0x83, 0x10, 0x96, 0x18, 0x44, 0x88, 0x4f, 0x1d, 0x8f, 0x6e, 0x6e, 0xf8, 0x10,
	0xfb, 0x18, 0x1e, 0x62, 0xff, 0x19, 0x96, 0xdd, 0x2e, 0xf3, 0xb9, 0x5b, 0x22, 0x45, 0x19, 0x83,
	0x0b, 0xb1, 0x41, 0x94, 0xc9, 0x5e, 0x37, 0x35, 0x1f, 0x56, 0x6a, 0x09, 0x89, 0x70, 0x3f, 0x17,
	0xd7, 0xca, 0xf4, 0x23, 0x8b, 0x08, 0xe2, 0x6b, 0x43, 0x6b, 0x0d, 0x30, 0x08, 0x87, 0x78, 0xf8,
	0x0d, 0x98, 0xe2, 0xfc, 0x82, 0xeb, 0xba, 0xe0, 0x02, 0x61, 0xb7, 0x54, 0xbc, 0x16, 0x36, 0x76,
	0xbe, 0x17, 0xc2, 0x01, 0x00, 0x96, 0xc0, 0xf4, 0x91, 0xce, 0x2a, 0xaf, 0xf7, 0x98, 0xad, 0x33,
	0x52, 0x3b, 0x55, 0x6e, 0x08, 0x86, 0xdb, 0x7d, 0xfd, 0x56, 0x51, 0xf6, 0x29, 0x2a, 0xbd, 0x6e,
	0xea, 0xba, 0xcb, 0x79, 0x06, 0x8c, 0xf0, 0x59, 0x32, 0xb8, 0x0e, 0x40, 0x83, 0x10, 0x63, 0x87,
	0x32, 0xb3, 0x7a, 0xaa, 0xcc, 0x8b, 0xbe, 0x48, 0xfa, 0x94, 0xa1, 0x0d, 0x61, 0xc9, 0x91, 0xb7,
	0xa7, 0xbc, 0x53, 0xfa, 0xc1, 0x21, 0xb6, 0x72, 0xb3, 0xbf, 0x3d, 0xf5, 0x2d, 0xc3, 0xb7, 0xa7,
	0x3e, 0x02, 0xde, 0x07, 0xb1, 0x93, 0x6d, 0x43, 0x51, 0x04, 0xe5, 0x52, 0x98, 0xdb, 0x27, 0xdb,
	0x06, 0x67, 0xbb, 0x06, 0xe6, 0x38, 0x5b, 0xf5, 0x91, 0xf6, 0x94, 0xb3, 0xb5, 0x56, 0xf3, 0x9d,
	0x45, 0xcc, 0x11, 0xf0, 0x21, 0x98, 0xe0, 0xf1, 0xd8, 0x36, 0x94, 0x5b, 0x02, 0xab, 0x86, 0x1b,
	0xcb, 0x7d, 0x1f, 0x09, 0xf7, 0x70, 0x70, 0x17, 0x00, 0xb7, 0xdf, 0xda, 0x22, 0x55, 0x47, 0x49,
	0x8a, 0x5c, 0xba, 0xd5, 0x17, 0xe4, 0xe7, 0x81, 0x83, 0x1c, 0xa4, 0x10, 0x86, 0xb0, 0xc4, 0x01,
	0xbf, 0xf3, 0x73, 0x88, 0x13, 0xfe, 0x4f, 0x10, 0x2a, 0x7d, 0x84, 0xfb, 0xbe, 0xfd, 0x7c, 0x12,
	0x09, 0xba, 0x90, 0x00, 0xee, 0xf2, 0xfe, 0x38, 0x48, 0x52, 0xe5, 0xb6, 0xc8, 0x82, 0xf9, 0x73,
	0x39, 0xf9, 0x42, 0xb7, 0x9a, 0xe4, 0x6c, 0xdf, 0x1c, 0x80, 0x44, 0xdf, 0x1c, 0x8c, 0x92, 0xef,
	0x27, 0x41, 0x42, 0x6a, 0xd3, 0xe1, 0x26, 0x98, 0xf2, 0xef, 0x4c, 0xe2, 0xce, 0x91, 0xc8, 0xcf,
	0xb8, 0x72, 0x1f, 0x7b, 0x6f, 0xe5, 0x54, 0xf5, 0x3d, 0x11, 0x0e, 0x40, 0x72, 0xd7, 0x3c, 0x7a,
	0x81, 0xae, 0xb9, 0x6f, 0xa5, 0xb1, 0x0b, 0xaf, 0x14, 0xfe, 0x28, 0x7f, 0x89, 0x31, 0xf1, 0x25,
	0x16, 0xa3, 0xef, 0x2b, 0x1f, 0xf3, 0x55, 0x4a, 0x67, 0xb2, 0x66, 0x5c, 0x50, 0x2f, 0x0d, 0xa0,
	0xfe, 0xa8, 0x0c, 0x4a, 0xfe, 0x16, 0x03, 0xf1, 0x40, 0x0c, 0xfc, 0x3f, 0x18, 0x35, 0x0d, 0xef,
	0x36, 0x38, 0xdd, 0xeb, 0xa6, 0xe2, 0xde, 0xf5, 0xcc, 0x40, 0x78, 0xd4, 0x34, 0xe0, 0x1d, 0x30,
	0x26, 0xae, 0x8b, 0xa3, 0xc2, 0x61, 0xb6, 0xd7, 0x4d, 0x25, 0xbc, 0x4d, 0xcc, 0xaf, 0x8a, 0x58,
	0x18, 0x61, 0x15, 0x00, 0xc3, 0xac, 0x93, 0x86, 0x63, 0xd2, 0x06, 0x0f, 0x2d, 0xd7, 0x9b, 0x1d,
	0x26, 0x14, 0x99, 0xad, 0x00, 0x26, 0x2b, 0x0f, 0xc9, 0x10, 0x96, 0x98, 0xe1, 0x01, 0x98, 0xac,
	0x50, 0xab, 0x59, 0x6f, 0xf8, 0xf1, 0x5e, 0x1e, 0x6a, 0x92, 0xc7, 0x2e, 0x46, 0xce, 0x0f, 0x8f,
	0x06, 0x61, 0x9f, 0x30, 0xb9, 0x0b, 0x40, 0x28, 0xe6, 0x32, 0xa2, 0x92, 0x7c, 0x06, 0x26, 0xbd,
	0x99, 0x2f, 0x85, 0x6e, 0x17, 0x80, 0xf0, 0x3b, 0x5f, 0x0a, 0xe3, 0xef, 0x23, 0x60, 0xfa, 0xcc,
	0x09, 0x0f, 0xb3, 0x60, 0x4a, 0x9c, 0xe4, 0x3b, 0xcd, 0xba, 0xe0, 0x1e, 0x97, 0x37, 0xa7, 0x6f,
	0x41, 0x38, 0x70, 0x82, 0x0f, 0xbc, 0x3a, 0xb2, 0xdd, 0x60, 0xc4, 0x3e, 0xd1, 0x2d, 0x31, 0xe1,
	0xf8, 0xb9, 0x4a, 0xe1, 0x9b, 0xfd, 0x4a, 0xe1, 0x8f, 0x79, 0xa5, 0xa8, 0xea, 0xa6, 0x45, 0x8c,
	0x3d, 0x46, 0x8f, 0xc5, 0xa6, 0x3c, 0x53, 0x29, 0x42, 0x1b, 0xc2, 0x92, 0xe3, 0x27, 0x88, 0xc5,
	0xaf, 0x97, 0xbe, 0x29, 0x0e, 0x3e, 0xb0, 0x29, 0xee, 0x44, 0x9d, 0xd4, 0x1f, 0xb9, 0x11, 0x76,
	0xfa, 0x37, 0xc2, 0x42, 0x24, 0xf1, 0x7f, 0x22, 0xf9, 0x8b, 0x0f, 0x0f, 0x1e, 0x5c, 0xec, 0x9f,
	0xb6, 0xa3, 0x09, 0xe1, 0xb4, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xa7, 0x0d, 0xe0,
	0x47, 0x14, 0x00, 0x00,
}
