// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: debug.proto

package collector

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
//DebugCollector请求
type DebugCollectorRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//执行目标实例Id
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target" form:"target"`
	//
	//脚本入参
	Args []*DebugCollectorRequest_Args `protobuf:"bytes,3,rep,name=args,proto3" json:"args" form:"args"`
	//
	//脚本类型
	Script string `protobuf:"bytes,4,opt,name=script,proto3" json:"script" form:"script"`
	//
	//脚本内容
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content" form:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorRequest) Reset()         { *m = DebugCollectorRequest{} }
func (m *DebugCollectorRequest) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorRequest) ProtoMessage()    {}
func (*DebugCollectorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0}
}
func (m *DebugCollectorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorRequest.Unmarshal(m, b)
}
func (m *DebugCollectorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorRequest.Marshal(b, m, deterministic)
}
func (m *DebugCollectorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorRequest.Merge(m, src)
}
func (m *DebugCollectorRequest) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorRequest.Size(m)
}
func (m *DebugCollectorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorRequest proto.InternalMessageInfo

func (m *DebugCollectorRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *DebugCollectorRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *DebugCollectorRequest) GetArgs() []*DebugCollectorRequest_Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *DebugCollectorRequest) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *DebugCollectorRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DebugCollectorRequest_Args struct {
	//
	//脚本入参值(当入参为自定义时需要)
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value" form:"value"`
	//
	//脚本入参key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key" form:"key"`
	//
	//参数的来源
	//attr_id: CMDB模型的属性
	//custom: 自定义
	//
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source" form:"source"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorRequest_Args) Reset()         { *m = DebugCollectorRequest_Args{} }
func (m *DebugCollectorRequest_Args) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorRequest_Args) ProtoMessage()    {}
func (*DebugCollectorRequest_Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0, 0}
}
func (m *DebugCollectorRequest_Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorRequest_Args.Unmarshal(m, b)
}
func (m *DebugCollectorRequest_Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorRequest_Args.Marshal(b, m, deterministic)
}
func (m *DebugCollectorRequest_Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorRequest_Args.Merge(m, src)
}
func (m *DebugCollectorRequest_Args) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorRequest_Args.Size(m)
}
func (m *DebugCollectorRequest_Args) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorRequest_Args.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorRequest_Args proto.InternalMessageInfo

func (m *DebugCollectorRequest_Args) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DebugCollectorRequest_Args) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DebugCollectorRequest_Args) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

//
//DebugCollector返回
type DebugCollectorResponse struct {
	//
	//脚本执行状态
	//ok: 正常执行且结果通过调试,
	//failed: 正常执行但结果未通过调试,
	//unknown: 执行失败或解析脚本输出失败
	//
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	//
	//执行日志
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//指标组列表
	MetricGroups         []*DebugCollectorResponse_MetricGroups `protobuf:"bytes,3,rep,name=metric_groups,json=metricGroups,proto3" json:"metric_groups" form:"metric_groups"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *DebugCollectorResponse) Reset()         { *m = DebugCollectorResponse{} }
func (m *DebugCollectorResponse) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorResponse) ProtoMessage()    {}
func (*DebugCollectorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1}
}
func (m *DebugCollectorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse.Unmarshal(m, b)
}
func (m *DebugCollectorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse.Merge(m, src)
}
func (m *DebugCollectorResponse) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse.Size(m)
}
func (m *DebugCollectorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse proto.InternalMessageInfo

func (m *DebugCollectorResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DebugCollectorResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DebugCollectorResponse) GetMetricGroups() []*DebugCollectorResponse_MetricGroups {
	if m != nil {
		return m.MetricGroups
	}
	return nil
}

type DebugCollectorResponse_MetricGroups struct {
	//
	//指标组状态
	//ok: 通过调试
	//undefined: 指标组未定义
	//missing: 未采集指标组
	//dims_not_unique: 维度不能唯一确定数据
	//metrics_not_ok: dims或vals不合法(未定义,类型错误,未采集到或重复)
	//
	//如果同时具有多种状态只返回最重要的状态
	//1:metrics_not_ok和dims_not_unique只返回dims_not_unique
	//2:undefined和dims_not_unique只返回undefined
	//
	MetricGroupStatus string `protobuf:"bytes,1,opt,name=metric_group_status,json=metricGroupStatus,proto3" json:"metric_group_status" form:"metric_group_status"`
	//
	//dim的状态列表
	DimStatus []*DebugCollectorResponse_MetricGroups_DimStatus `protobuf:"bytes,2,rep,name=dim_status,json=dimStatus,proto3" json:"dim_status" form:"dim_status"`
	//
	//val的状态列表
	ValStatus []*DebugCollectorResponse_MetricGroups_ValStatus `protobuf:"bytes,3,rep,name=val_status,json=valStatus,proto3" json:"val_status" form:"val_status"`
	//
	//指标组的值的列表
	List []*DebugCollectorResponse_MetricGroups_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	//
	//指标组id
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id" form:"id"`
	//
	//指标组名
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" form:"name"`
	//
	//两级分类，用.分割，(如：主机状态.基本配置 )
	Category             string   `protobuf:"bytes,7,opt,name=category,proto3" json:"category" form:"category"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups) Reset()         { *m = DebugCollectorResponse_MetricGroups{} }
func (m *DebugCollectorResponse_MetricGroups) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorResponse_MetricGroups) ProtoMessage()    {}
func (*DebugCollectorResponse_MetricGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0}
}
func (m *DebugCollectorResponse_MetricGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups) GetMetricGroupStatus() string {
	if m != nil {
		return m.MetricGroupStatus
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups) GetDimStatus() []*DebugCollectorResponse_MetricGroups_DimStatus {
	if m != nil {
		return m.DimStatus
	}
	return nil
}

func (m *DebugCollectorResponse_MetricGroups) GetValStatus() []*DebugCollectorResponse_MetricGroups_ValStatus {
	if m != nil {
		return m.ValStatus
	}
	return nil
}

func (m *DebugCollectorResponse_MetricGroups) GetList() []*DebugCollectorResponse_MetricGroups_List {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *DebugCollectorResponse_MetricGroups) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type DebugCollectorResponse_MetricGroups_DimStatus struct {
	//
	//dim的状态
	//ok: 通过调试
	//undefined: 未定义
	//mismatched_type: 类型错误
	//missing: 未采集到
	//duplicated: 重复
	//
	//如果同时具有多种状态只返回最重要的状态
	//1:mismatched_type和duplicated只返回mismatched_type
	//2:undefined和duplicated只返回undefined
	//
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	//
	//维度id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//维度名称
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups_DimStatus) Reset() {
	*m = DebugCollectorResponse_MetricGroups_DimStatus{}
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) String() string {
	return proto.CompactTextString(m)
}
func (*DebugCollectorResponse_MetricGroups_DimStatus) ProtoMessage() {}
func (*DebugCollectorResponse_MetricGroups_DimStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0, 0}
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups_DimStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups_DimStatus proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups_DimStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_DimStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_DimStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DebugCollectorResponse_MetricGroups_ValStatus struct {
	//
	//val的状态
	//ok: 通过调试
	//undefined: 未定义
	//mismatched_type: 类型错误
	//missing: 未采集到
	//duplicated: 重复
	//
	//如果同时具有多种状态只返回最重要的状态
	//1:mismatched_type和duplicated只返回mismatched_type
	//2:undefined和duplicated只返回undefined
	//
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	//
	//指标id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//指标名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//指标类型
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type" form:"type"`
	//
	//指标单位
	Unit                 string   `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit" form:"unit"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups_ValStatus) Reset() {
	*m = DebugCollectorResponse_MetricGroups_ValStatus{}
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) String() string {
	return proto.CompactTextString(m)
}
func (*DebugCollectorResponse_MetricGroups_ValStatus) ProtoMessage() {}
func (*DebugCollectorResponse_MetricGroups_ValStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0, 1}
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups_ValStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups_ValStatus proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups_ValStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_ValStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_ValStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_ValStatus) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_ValStatus) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type DebugCollectorResponse_MetricGroups_List struct {
	//
	//一条数据的维度列表
	Dims []*DebugCollectorResponse_MetricGroups_List_Dims `protobuf:"bytes,1,rep,name=dims,proto3" json:"dims" form:"dims"`
	//
	//一条数据的val列表
	Vals                 []*DebugCollectorResponse_MetricGroups_List_Vals `protobuf:"bytes,2,rep,name=vals,proto3" json:"vals" form:"vals"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups_List) Reset() {
	*m = DebugCollectorResponse_MetricGroups_List{}
}
func (m *DebugCollectorResponse_MetricGroups_List) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorResponse_MetricGroups_List) ProtoMessage()    {}
func (*DebugCollectorResponse_MetricGroups_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0, 2}
}
func (m *DebugCollectorResponse_MetricGroups_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups_List) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups_List) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups_List proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups_List) GetDims() []*DebugCollectorResponse_MetricGroups_List_Dims {
	if m != nil {
		return m.Dims
	}
	return nil
}

func (m *DebugCollectorResponse_MetricGroups_List) GetVals() []*DebugCollectorResponse_MetricGroups_List_Vals {
	if m != nil {
		return m.Vals
	}
	return nil
}

type DebugCollectorResponse_MetricGroups_List_Dims struct {
	//
	//维度的值
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value" form:"value"`
	//
	//维度id
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups_List_Dims) Reset() {
	*m = DebugCollectorResponse_MetricGroups_List_Dims{}
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) String() string {
	return proto.CompactTextString(m)
}
func (*DebugCollectorResponse_MetricGroups_List_Dims) ProtoMessage() {}
func (*DebugCollectorResponse_MetricGroups_List_Dims) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0, 2, 0}
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups_List_Dims) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Dims proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups_List_Dims) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_List_Dims) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DebugCollectorResponse_MetricGroups_List_Vals struct {
	//
	//val的值
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value" form:"value"`
	//
	//指标id
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectorResponse_MetricGroups_List_Vals) Reset() {
	*m = DebugCollectorResponse_MetricGroups_List_Vals{}
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) String() string {
	return proto.CompactTextString(m)
}
func (*DebugCollectorResponse_MetricGroups_List_Vals) ProtoMessage() {}
func (*DebugCollectorResponse_MetricGroups_List_Vals) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1, 0, 2, 1}
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals.Unmarshal(m, b)
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals.Merge(m, src)
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals.Size(m)
}
func (m *DebugCollectorResponse_MetricGroups_List_Vals) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponse_MetricGroups_List_Vals proto.InternalMessageInfo

func (m *DebugCollectorResponse_MetricGroups_List_Vals) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DebugCollectorResponse_MetricGroups_List_Vals) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DebugCollectorApi返回
type DebugCollectorResponseWrapper struct {
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
	Data                 *DebugCollectorResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DebugCollectorResponseWrapper) Reset()         { *m = DebugCollectorResponseWrapper{} }
func (m *DebugCollectorResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DebugCollectorResponseWrapper) ProtoMessage()    {}
func (*DebugCollectorResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{2}
}
func (m *DebugCollectorResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectorResponseWrapper.Unmarshal(m, b)
}
func (m *DebugCollectorResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectorResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DebugCollectorResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectorResponseWrapper.Merge(m, src)
}
func (m *DebugCollectorResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DebugCollectorResponseWrapper.Size(m)
}
func (m *DebugCollectorResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectorResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectorResponseWrapper proto.InternalMessageInfo

func (m *DebugCollectorResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DebugCollectorResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DebugCollectorResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DebugCollectorResponseWrapper) GetData() *DebugCollectorResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DebugCollectorRequest)(nil), "collector.DebugCollectorRequest")
	proto.RegisterType((*DebugCollectorRequest_Args)(nil), "collector.DebugCollectorRequest.Args")
	proto.RegisterType((*DebugCollectorResponse)(nil), "collector.DebugCollectorResponse")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups)(nil), "collector.DebugCollectorResponse.MetricGroups")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups_DimStatus)(nil), "collector.DebugCollectorResponse.MetricGroups.DimStatus")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups_ValStatus)(nil), "collector.DebugCollectorResponse.MetricGroups.ValStatus")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups_List)(nil), "collector.DebugCollectorResponse.MetricGroups.List")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups_List_Dims)(nil), "collector.DebugCollectorResponse.MetricGroups.List.Dims")
	proto.RegisterType((*DebugCollectorResponse_MetricGroups_List_Vals)(nil), "collector.DebugCollectorResponse.MetricGroups.List.Vals")
	proto.RegisterType((*DebugCollectorResponseWrapper)(nil), "collector.DebugCollectorResponseWrapper")
}

func init() { proto.RegisterFile("debug.proto", fileDescriptor_8d9d361be58531fb) }

var fileDescriptor_8d9d361be58531fb = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xe1, 0x6e, 0xdc, 0x44,
	0x10, 0xc7, 0x95, 0x3b, 0x27, 0xed, 0xed, 0x25, 0x69, 0xb3, 0xa5, 0x95, 0x65, 0x29, 0x38, 0x2c,
	0x29, 0x4a, 0xa3, 0x9e, 0xd3, 0x36, 0x52, 0x55, 0x90, 0xf8, 0xd0, 0x50, 0x8a, 0x40, 0x94, 0x0f,
	0x8b, 0x54, 0x10, 0x21, 0x89, 0x36, 0xf6, 0xd6, 0x58, 0xb1, 0x6f, 0xcd, 0x7a, 0x7d, 0x25, 0x8d,
	0xfa, 0x34, 0xbc, 0x0a, 0x1f, 0xfb, 0x0c, 0x46, 0xea, 0x03, 0xf0, 0xc1, 0x4f, 0x80, 0x66, 0xd6,
	0xbe, 0x73, 0x48, 0x95, 0x92, 0x08, 0xf1, 0xe9, 0xd6, 0x33, 0xff, 0xf9, 0xcd, 0xcc, 0xce, 0xee,
	0xea, 0xc8, 0x30, 0x92, 0x87, 0x65, 0x1c, 0xe4, 0x5a, 0x19, 0x45, 0x07, 0xa1, 0x4a, 0x53, 0x19,
	0x1a, 0xa5, 0xbd, 0x51, 0x9c, 0x98, 0x5f, 0xca, 0xc3, 0x20, 0x54, 0xd9, 0x56, 0xac, 0x62, 0xb5,
	0x85, 0x8a, 0xc3, 0xf2, 0x05, 0x7e, 0xe1, 0x07, 0xae, 0x6c, 0xa4, 0xf7, 0xb0, 0x23, 0xcf, 0x5e,
	0x26, 0xe6, 0x48, 0xbd, 0xdc, 0x8a, 0xd5, 0x08, 0x9d, 0xa3, 0x89, 0x48, 0x93, 0x48, 0x18, 0xa5,
	0x8b, 0xad, 0xe9, 0xd2, 0xc6, 0xb1, 0x37, 0x7d, 0x72, 0xf3, 0x09, 0x54, 0xf0, 0x45, 0x9b, 0x99,
	0xcb, 0x5f, 0x4b, 0x59, 0x18, 0xca, 0xc9, 0xd5, 0x3c, 0x2d, 0xe3, 0x64, 0xfc, 0x75, 0xe4, 0xce,
	0xad, 0xcd, 0x6d, 0x0c, 0x76, 0x1e, 0xd6, 0x95, 0x7f, 0xed, 0x85, 0xd2, 0xd9, 0x67, 0xac, 0xf5,
	0xb0, 0xb7, 0x7f, 0xfa, 0x3e, 0x59, 0xdd, 0xdf, 0x15, 0xa3, 0x57, 0x8f, 0x47, 0x3f, 0x1d, 0xec,
	0xed, 0xde, 0x1b, 0x7d, 0xda, 0xae, 0x4f, 0xee, 0xdd, 0xdd, 0xbe, 0xff, 0x7a, 0x9d, 0x4f, 0x39,
	0xf4, 0x73, 0xb2, 0x60, 0x84, 0x8e, 0xa5, 0x71, 0x7b, 0x48, 0xbc, 0x5d, 0x57, 0xfe, 0x92, 0x25,
	0x5a, 0x3b, 0xf0, 0xae, 0x93, 0xe5, 0xfd, 0x06, 0xb3, 0x77, 0x72, 0x7f, 0xfb, 0xf5, 0x3a, 0x6f,
	0x82, 0xe8, 0x37, 0xc4, 0x11, 0x3a, 0x2e, 0xdc, 0xfe, 0x5a, 0x7f, 0x63, 0xf8, 0xe0, 0x76, 0x30,
	0xdd, 0xad, 0xe0, 0x9d, 0x2d, 0x04, 0x8f, 0x75, 0x5c, 0xec, 0x5c, 0xab, 0x2b, 0x7f, 0x68, 0x73,
	0x40, 0x30, 0xe3, 0xc8, 0xa0, 0x77, 0xc8, 0x42, 0x11, 0xea, 0x24, 0x37, 0xae, 0x83, 0xa5, 0xac,
	0xcc, 0x4a, 0xb1, 0x76, 0xc6, 0x1b, 0x01, 0xbd, 0x4b, 0xae, 0x84, 0x6a, 0x6c, 0xe4, 0xd8, 0xb8,
	0xf3, 0xa8, 0xa5, 0x75, 0xe5, 0x2f, 0x5b, 0x6d, 0xe3, 0x60, 0xbc, 0x95, 0x78, 0x27, 0xc4, 0x81,
	0xbc, 0xf4, 0x13, 0x32, 0x3f, 0x11, 0x69, 0x29, 0x9b, 0xcd, 0xbb, 0x5e, 0x57, 0xfe, 0xa2, 0x8d,
	0x41, 0x33, 0xe3, 0xd6, 0x4d, 0xd7, 0x48, 0xff, 0x48, 0x1e, 0x37, 0x1b, 0xb2, 0x5c, 0x57, 0x3e,
	0xb1, 0xaa, 0x23, 0x79, 0xcc, 0x38, 0xb8, 0xb0, 0x54, 0x55, 0xea, 0x50, 0xba, 0xfd, 0x33, 0xa5,
	0xa2, 0x1d, 0x4a, 0xb5, 0x8b, 0xdf, 0x87, 0xe4, 0xd6, 0x3f, 0xf7, 0xa2, 0xc8, 0xd5, 0xb8, 0x90,
	0x48, 0x31, 0xc2, 0x94, 0x45, 0x53, 0x50, 0x97, 0x82, 0x76, 0xa0, 0xe0, 0x02, 0x4a, 0xca, 0x8a,
	0xf8, 0x6c, 0x49, 0x59, 0x11, 0x33, 0x0e, 0x2e, 0x9a, 0x91, 0xa5, 0x4c, 0x1a, 0x9d, 0x84, 0x07,
	0xb1, 0x56, 0x65, 0xde, 0x8e, 0x24, 0x38, 0x67, 0x24, 0xb6, 0x8c, 0xe0, 0x19, 0x86, 0x7d, 0x85,
	0x51, 0x3b, 0x6e, 0x5d, 0xf9, 0x1f, 0x34, 0xec, 0x2e, 0x8e, 0xf1, 0xc5, 0xac, 0xa3, 0xf3, 0xde,
	0x0c, 0xc8, 0x62, 0x37, 0x90, 0x7e, 0x47, 0x6e, 0x74, 0x03, 0x0e, 0x4e, 0x75, 0xf6, 0x61, 0x5d,
	0xf9, 0xde, 0x59, 0xea, 0x41, 0xdb, 0xe6, 0x4a, 0x87, 0xfd, 0xbd, 0xed, 0x78, 0x4c, 0x48, 0x94,
	0x64, 0x2d, 0xa6, 0x87, 0xcd, 0x3c, 0xba, 0x58, 0x33, 0xc1, 0x93, 0x24, 0xb3, 0xb4, 0x9d, 0x9b,
	0x75, 0xe5, 0xaf, 0xd8, 0x02, 0x66, 0x54, 0xc6, 0x07, 0x51, 0xab, 0x80, 0x7c, 0x13, 0x91, 0xb6,
	0xf9, 0xfa, 0x97, 0xca, 0xf7, 0x5c, 0xa4, 0x67, 0xf3, 0xcd, 0xa8, 0x8c, 0x0f, 0x26, 0xad, 0x82,
	0xfe, 0x48, 0x9c, 0x34, 0x29, 0xe0, 0xac, 0x43, 0xa6, 0xed, 0x0b, 0x66, 0xfa, 0x36, 0x29, 0x4c,
	0xf7, 0x1e, 0x01, 0x8a, 0x71, 0x24, 0xd2, 0x55, 0xd2, 0x4b, 0xa2, 0xe6, 0x5e, 0x2c, 0xd5, 0x95,
	0x3f, 0xb0, 0x92, 0x24, 0x62, 0xbc, 0x97, 0x44, 0xf4, 0x63, 0xe2, 0x8c, 0x45, 0x26, 0xdd, 0x05,
	0x14, 0x74, 0x18, 0x60, 0x65, 0x1c, 0x9d, 0xf4, 0x29, 0xb9, 0x1a, 0x0a, 0x23, 0x63, 0xa5, 0x8f,
	0xdd, 0x2b, 0x28, 0xdc, 0x9c, 0x3d, 0x35, 0xad, 0x07, 0x9e, 0x86, 0x1b, 0x64, 0x65, 0x7f, 0x63,
	0x77, 0x3f, 0xd8, 0xdb, 0xfc, 0x39, 0xb8, 0xb3, 0x89, 0x8b, 0x75, 0x3e, 0x8d, 0xf5, 0x5e, 0x91,
	0xc1, 0x74, 0x08, 0x17, 0x39, 0xef, 0xb6, 0x87, 0xde, 0xfb, 0x7a, 0xe8, 0x9f, 0xd3, 0x83, 0xf7,
	0xc7, 0x1c, 0x19, 0x4c, 0x27, 0xf2, 0x3f, 0x27, 0x07, 0x91, 0x39, 0xce, 0x65, 0xf3, 0x94, 0x75,
	0x44, 0x60, 0x65, 0x1c, 0x9d, 0x20, 0x2a, 0xc7, 0x49, 0xfb, 0x86, 0x75, 0x44, 0x60, 0x65, 0x1c,
	0x9d, 0xde, 0xdb, 0x1e, 0x71, 0x60, 0xdc, 0x74, 0x8f, 0x38, 0x51, 0x92, 0x41, 0xfd, 0x97, 0x39,
	0x9b, 0x80, 0x80, 0x0b, 0x71, 0xea, 0xf9, 0x05, 0x1e, 0xe3, 0x88, 0x05, 0xfc, 0x44, 0xa4, 0x97,
	0xbd, 0x6a, 0x88, 0x7f, 0x2e, 0xd2, 0x53, 0x78, 0xe0, 0x31, 0x8e, 0x58, 0xef, 0x19, 0x71, 0x20,
	0xfb, 0xbf, 0x7e, 0x84, 0xcf, 0x1f, 0x02, 0xe0, 0x20, 0xdb, 0x7f, 0x84, 0x63, 0x7f, 0xcd, 0x91,
	0xd5, 0x77, 0xb7, 0xf9, 0x83, 0x16, 0x79, 0x2e, 0x35, 0xcc, 0x2a, 0x54, 0x91, 0xcd, 0x33, 0xdf,
	0x6d, 0x12, 0xac, 0x8c, 0xa3, 0x93, 0x3e, 0x22, 0x43, 0xf8, 0xfd, 0xf2, 0xb7, 0x3c, 0x15, 0xc9,
	0xb8, 0x49, 0x77, 0xab, 0xae, 0x7c, 0x3a, 0xd3, 0x36, 0x4e, 0xc6, 0xbb, 0x52, 0xe8, 0x43, 0x6a,
	0xad, 0x74, 0x73, 0xaa, 0x3a, 0x7d, 0xa0, 0x99, 0x71, 0xeb, 0xa6, 0x4f, 0x89, 0x13, 0x09, 0x23,
	0xf0, 0x5c, 0x0d, 0x1f, 0x7c, 0xf4, 0xde, 0x29, 0x9d, 0x9a, 0xb6, 0x30, 0x02, 0xa6, 0x2d, 0x8c,
	0x38, 0x5c, 0xc0, 0x3f, 0x1b, 0xdb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x99, 0x75, 0x19, 0x36,
	0xed, 0x08, 0x00, 0x00,
}
