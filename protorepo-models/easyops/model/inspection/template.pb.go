// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: template.proto

package inspection

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
//报告模板
type InspectionTemplate struct {
	//
	//套件ID
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//报告模板ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//报告模板名
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//报告总览模板
	SummaryTemplates *InspectionTemplate_SummaryTemplates `protobuf:"bytes,5,opt,name=summaryTemplates,proto3" json:"summaryTemplates" form:"summaryTemplates"`
	//
	//报告详情模板
	MetricGroups []*InspectionTemplate_MetricGroups `protobuf:"bytes,6,rep,name=metricGroups,proto3" json:"metricGroups" form:"metricGroups"`
	//
	//创建者
	Creator string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime                string   `protobuf:"bytes,8,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTemplate) Reset()         { *m = InspectionTemplate{} }
func (m *InspectionTemplate) String() string { return proto.CompactTextString(m) }
func (*InspectionTemplate) ProtoMessage()    {}
func (*InspectionTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0}
}
func (m *InspectionTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTemplate.Unmarshal(m, b)
}
func (m *InspectionTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTemplate.Marshal(b, m, deterministic)
}
func (m *InspectionTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTemplate.Merge(m, src)
}
func (m *InspectionTemplate) XXX_Size() int {
	return xxx_messageInfo_InspectionTemplate.Size(m)
}
func (m *InspectionTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTemplate proto.InternalMessageInfo

func (m *InspectionTemplate) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *InspectionTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InspectionTemplate) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *InspectionTemplate) GetSummaryTemplates() *InspectionTemplate_SummaryTemplates {
	if m != nil {
		return m.SummaryTemplates
	}
	return nil
}

func (m *InspectionTemplate) GetMetricGroups() []*InspectionTemplate_MetricGroups {
	if m != nil {
		return m.MetricGroups
	}
	return nil
}

func (m *InspectionTemplate) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *InspectionTemplate) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

type InspectionTemplate_SummaryTemplates struct {
	//
	//如何展示指标组
	MetricGroups []*InspectionTemplate_SummaryTemplates_MetricGroups `protobuf:"bytes,1,rep,name=metricGroups,proto3" json:"metricGroups" form:"metricGroups"`
	//
	//如何展示指标
	Metrics              []*InspectionTemplate_SummaryTemplates_Metrics `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics" form:"metrics"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *InspectionTemplate_SummaryTemplates) Reset()         { *m = InspectionTemplate_SummaryTemplates{} }
func (m *InspectionTemplate_SummaryTemplates) String() string { return proto.CompactTextString(m) }
func (*InspectionTemplate_SummaryTemplates) ProtoMessage()    {}
func (*InspectionTemplate_SummaryTemplates) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0, 0}
}
func (m *InspectionTemplate_SummaryTemplates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates.Unmarshal(m, b)
}
func (m *InspectionTemplate_SummaryTemplates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates.Marshal(b, m, deterministic)
}
func (m *InspectionTemplate_SummaryTemplates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates.Merge(m, src)
}
func (m *InspectionTemplate_SummaryTemplates) XXX_Size() int {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates.Size(m)
}
func (m *InspectionTemplate_SummaryTemplates) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTemplate_SummaryTemplates proto.InternalMessageInfo

func (m *InspectionTemplate_SummaryTemplates) GetMetricGroups() []*InspectionTemplate_SummaryTemplates_MetricGroups {
	if m != nil {
		return m.MetricGroups
	}
	return nil
}

func (m *InspectionTemplate_SummaryTemplates) GetMetrics() []*InspectionTemplate_SummaryTemplates_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type InspectionTemplate_SummaryTemplates_MetricGroups struct {
	//
	//排序
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index" form:"index"`
	//
	//高度
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height" form:"height"`
	//
	//宽度
	Width int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width" form:"width"`
	//
	//图表类型(Form 表格)
	DisplayType string `protobuf:"bytes,4,opt,name=displayType,proto3" json:"displayType" form:"displayType"`
	//
	//是否转换行列(仅当图表类型为表格时有效)
	Transposed bool `protobuf:"varint,5,opt,name=transposed,proto3" json:"transposed" form:"transposed"`
	//
	//指标组id
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id" form:"id"`
	//
	//指标组名
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" form:"name"`
	//
	//两级分类，用.分割，(如：主机状态.基本配置 )
	Category             string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category" form:"category"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) Reset() {
	*m = InspectionTemplate_SummaryTemplates_MetricGroups{}
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) String() string {
	return proto.CompactTextString(m)
}
func (*InspectionTemplate_SummaryTemplates_MetricGroups) ProtoMessage() {}
func (*InspectionTemplate_SummaryTemplates_MetricGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0, 0, 0}
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups.Unmarshal(m, b)
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups.Marshal(b, m, deterministic)
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups.Merge(m, src)
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) XXX_Size() int {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups.Size(m)
}
func (m *InspectionTemplate_SummaryTemplates_MetricGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTemplate_SummaryTemplates_MetricGroups proto.InternalMessageInfo

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetTransposed() bool {
	if m != nil {
		return m.Transposed
	}
	return false
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_MetricGroups) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type InspectionTemplate_SummaryTemplates_Metrics struct {
	//
	//指标组Id
	MetricGroupId string `protobuf:"bytes,1,opt,name=metricGroupId,proto3" json:"metricGroupId" form:"metricGroupId"`
	//
	//排序
	Index int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index" form:"index"`
	//
	//高度
	Height int32 `protobuf:"varint,3,opt,name=height,proto3" json:"height" form:"height"`
	//
	//宽度
	Width int32 `protobuf:"varint,4,opt,name=width,proto3" json:"width" form:"width"`
	//
	//折线图横坐标ID(仅当图表类型为折线图时有效)
	AbscissaId string `protobuf:"bytes,5,opt,name=abscissaId,proto3" json:"abscissaId" form:"abscissaId"`
	//
	//折线图横坐标名(仅当图表类型为折线图时有效)
	AbscissaName string `protobuf:"bytes,6,opt,name=abscissaName,proto3" json:"abscissaName" form:"abscissaName"`
	//
	//图表类型(BasicInfo 基本信息/LineChart 折线图/BarChart 柱状图/DoughnutChart 环形图/Card 卡片)
	DisplayType string `protobuf:"bytes,7,opt,name=displayType,proto3" json:"displayType" form:"displayType"`
	//
	//指标id
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id" form:"id"`
	//
	//指标名称
	Name                 string   `protobuf:"bytes,9,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) Reset() {
	*m = InspectionTemplate_SummaryTemplates_Metrics{}
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) String() string {
	return proto.CompactTextString(m)
}
func (*InspectionTemplate_SummaryTemplates_Metrics) ProtoMessage() {}
func (*InspectionTemplate_SummaryTemplates_Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0, 0, 1}
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics.Unmarshal(m, b)
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics.Marshal(b, m, deterministic)
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics.Merge(m, src)
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) XXX_Size() int {
	return xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics.Size(m)
}
func (m *InspectionTemplate_SummaryTemplates_Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTemplate_SummaryTemplates_Metrics proto.InternalMessageInfo

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetMetricGroupId() string {
	if m != nil {
		return m.MetricGroupId
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetAbscissaId() string {
	if m != nil {
		return m.AbscissaId
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetAbscissaName() string {
	if m != nil {
		return m.AbscissaName
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionTemplate_SummaryTemplates_Metrics) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type InspectionTemplate_MetricGroups struct {
	//
	//排序
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index" form:"index"`
	//
	//图表类型(LineChart 折线图/BarChart 柱状图/Form 表格)
	DisplayType string `protobuf:"bytes,2,opt,name=displayType,proto3" json:"displayType" form:"displayType"`
	//
	//是否转换行列(仅当图表类型为表格时有效)
	Transposed bool `protobuf:"varint,3,opt,name=transposed,proto3" json:"transposed" form:"transposed"`
	//
	//折线图横坐标ID(仅当图表类型为折线图时有效)
	AbscissaId string `protobuf:"bytes,4,opt,name=abscissaId,proto3" json:"abscissaId" form:"abscissaId"`
	//
	//折线图横坐标名(仅当图表类型为折线图时有效)
	AbscissaName string `protobuf:"bytes,5,opt,name=abscissaName,proto3" json:"abscissaName" form:"abscissaName"`
	//
	//指标组id
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id" form:"id"`
	//
	//指标组名
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" form:"name"`
	//
	//两级分类，用.分割，(如：主机状态.基本配置 )
	Category             string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category" form:"category"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTemplate_MetricGroups) Reset()         { *m = InspectionTemplate_MetricGroups{} }
func (m *InspectionTemplate_MetricGroups) String() string { return proto.CompactTextString(m) }
func (*InspectionTemplate_MetricGroups) ProtoMessage()    {}
func (*InspectionTemplate_MetricGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0, 1}
}
func (m *InspectionTemplate_MetricGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTemplate_MetricGroups.Unmarshal(m, b)
}
func (m *InspectionTemplate_MetricGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTemplate_MetricGroups.Marshal(b, m, deterministic)
}
func (m *InspectionTemplate_MetricGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTemplate_MetricGroups.Merge(m, src)
}
func (m *InspectionTemplate_MetricGroups) XXX_Size() int {
	return xxx_messageInfo_InspectionTemplate_MetricGroups.Size(m)
}
func (m *InspectionTemplate_MetricGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTemplate_MetricGroups.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTemplate_MetricGroups proto.InternalMessageInfo

func (m *InspectionTemplate_MetricGroups) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *InspectionTemplate_MetricGroups) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func (m *InspectionTemplate_MetricGroups) GetTransposed() bool {
	if m != nil {
		return m.Transposed
	}
	return false
}

func (m *InspectionTemplate_MetricGroups) GetAbscissaId() string {
	if m != nil {
		return m.AbscissaId
	}
	return ""
}

func (m *InspectionTemplate_MetricGroups) GetAbscissaName() string {
	if m != nil {
		return m.AbscissaName
	}
	return ""
}

func (m *InspectionTemplate_MetricGroups) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionTemplate_MetricGroups) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InspectionTemplate_MetricGroups) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func init() {
	proto.RegisterType((*InspectionTemplate)(nil), "inspection.InspectionTemplate")
	proto.RegisterType((*InspectionTemplate_SummaryTemplates)(nil), "inspection.InspectionTemplate.SummaryTemplates")
	proto.RegisterType((*InspectionTemplate_SummaryTemplates_MetricGroups)(nil), "inspection.InspectionTemplate.SummaryTemplates.MetricGroups")
	proto.RegisterType((*InspectionTemplate_SummaryTemplates_Metrics)(nil), "inspection.InspectionTemplate.SummaryTemplates.Metrics")
	proto.RegisterType((*InspectionTemplate_MetricGroups)(nil), "inspection.InspectionTemplate.MetricGroups")
}

func init() { proto.RegisterFile("template.proto", fileDescriptor_b1b68e1b5f001c74) }

var fileDescriptor_b1b68e1b5f001c74 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0x25, 0x49, 0xd3, 0xa4, 0xd3, 0x76, 0xb7, 0x9d, 0x02, 0x6b, 0x05, 0x21, 0x57, 0xa6, 0x42,
	0x4e, 0xbb, 0x63, 0xa7, 0x69, 0x0b, 0x6c, 0x41, 0x44, 0x44, 0xa8, 0xa8, 0x0f, 0xf0, 0x30, 0xf4,
	0x69, 0xbd, 0xa9, 0xe4, 0xda, 0x53, 0x67, 0x44, 0x9c, 0xb1, 0x3c, 0x0e, 0xdd, 0x68, 0xbb, 0x3f,
	0xc1, 0x4f, 0x20, 0x21, 0xc1, 0x47, 0xec, 0x07, 0xf0, 0xce, 0x8b, 0x91, 0xf8, 0x04, 0x7f, 0x01,
	0xf2, 0x8c, 0x1d, 0x3b, 0x29, 0xda, 0x6d, 0xf3, 0xc4, 0x93, 0xe7, 0xde, 0x7b, 0xce, 0xdc, 0x99,
	0x39, 0xf7, 0x5e, 0x19, 0x3c, 0x8a, 0x88, 0x1f, 0x8c, 0xec, 0x88, 0x18, 0x41, 0xc8, 0x22, 0x06,
	0x01, 0x1d, 0xf3, 0x80, 0x38, 0x11, 0x65, 0xe3, 0x16, 0xf2, 0x68, 0x34, 0x9c, 0x5c, 0x19, 0x0e,
	0xf3, 0x4d, 0x8f, 0x79, 0xcc, 0x14, 0x90, 0xab, 0xc9, 0xb5, 0xb0, 0x84, 0x21, 0x56, 0x92, 0xda,
	0xfa, 0xac, 0x04, 0xf7, 0x6f, 0x68, 0xf4, 0x13, 0xbb, 0x31, 0x3d, 0x86, 0x44, 0x10, 0xfd, 0x6c,
	0x8f, 0xa8, 0x6b, 0x47, 0x2c, 0xe4, 0xe6, 0x6c, 0x29, 0x79, 0xda, 0x1f, 0xdb, 0x00, 0x9e, 0xcf,
	0xb2, 0x5e, 0x64, 0xe7, 0x81, 0x26, 0x68, 0x06, 0xa3, 0x89, 0x47, 0xc7, 0xe7, 0xae, 0x52, 0xd9,
	0xad, 0xe8, 0x6b, 0xfd, 0x9d, 0x24, 0x56, 0x1f, 0x5f, 0xb3, 0xd0, 0x3f, 0xd5, 0xf2, 0x88, 0x86,
	0x67, 0x20, 0x78, 0x0c, 0xaa, 0xd4, 0x55, 0xaa, 0x02, 0xba, 0x97, 0xc4, 0xea, 0x9a, 0x84, 0x52,
	0x57, 0xfb, 0xe7, 0x6f, 0x75, 0x07, 0x6c, 0x5f, 0x5a, 0x1d, 0xf4, 0xcc, 0x46, 0xd7, 0xdf, 0xa0,
	0xb3, 0xc1, 0xab, 0xee, 0xf1, 0xeb, 0x3d, 0x5c, 0xa5, 0x2e, 0xfc, 0x04, 0xac, 0x8c, 0x6d, 0x9f,
	0x28, 0x35, 0xc1, 0x7b, 0x9c, 0xc4, 0xea, 0xba, 0xe4, 0xa5, 0x5e, 0x0d, 0x8b, 0x60, 0x0a, 0xf2,
	0x89, 0xcf, 0x94, 0x95, 0x45, 0x50, 0xea, 0xd5, 0xb0, 0x08, 0xc2, 0x97, 0x60, 0x8b, 0x4f, 0x7c,
	0xdf, 0x0e, 0xa7, 0xf9, 0x1d, 0xb8, 0x52, 0xdf, 0xad, 0xe8, 0xeb, 0x5d, 0xd3, 0x28, 0x5e, 0xd5,
	0xb8, 0x7b, 0x55, 0xe3, 0xc7, 0x05, 0x5a, 0xff, 0xa3, 0x24, 0x56, 0x9f, 0xc8, 0x0c, 0x8b, 0x5b,
	0x6a, 0xf8, 0x4e, 0x16, 0x38, 0x04, 0x1b, 0x3e, 0x89, 0x42, 0xea, 0x7c, 0x17, 0xb2, 0x49, 0xc0,
	0x95, 0xd5, 0xdd, 0x9a, 0xbe, 0xde, 0x3d, 0x78, 0x47, 0xd6, 0xef, 0x4b, 0x94, 0xfe, 0x93, 0x24,
	0x56, 0x77, 0xf2, 0x3b, 0x15, 0x7e, 0x0d, 0xcf, 0xed, 0x0c, 0x9f, 0x82, 0x86, 0x13, 0x92, 0x54,
	0x3c, 0xa5, 0x21, 0xde, 0x02, 0x26, 0xb1, 0xfa, 0x48, 0xf2, 0xb2, 0x80, 0x86, 0x73, 0x08, 0xfc,
	0xab, 0x02, 0xea, 0x4e, 0x44, 0x7d, 0xa2, 0x34, 0x05, 0xf8, 0x4d, 0x25, 0x89, 0xd5, 0x8d, 0x0c,
	0x9d, 0xfa, 0x53, 0x65, 0x7e, 0xaf, 0x80, 0xdf, 0x2a, 0x97, 0xba, 0xde, 0x3b, 0xb5, 0x0e, 0xd1,
	0xb3, 0x41, 0x2a, 0xd2, 0x60, 0xbf, 0xdd, 0x13, 0xdf, 0x57, 0xc7, 0xaf, 0xdb, 0x48, 0x3f, 0xb4,
	0x3a, 0xa8, 0x3b, 0xb8, 0xed, 0x88, 0x78, 0x1b, 0xe9, 0x47, 0x56, 0x07, 0x1d, 0xe6, 0xf6, 0xad,
	0x75, 0x88, 0xba, 0x92, 0xd5, 0xb6, 0x2e, 0x76, 0x07, 0x7a, 0xd7, 0xea, 0xa0, 0xa3, 0xc1, 0xad,
	0xc0, 0x48, 0xf7, 0xa9, 0x6e, 0x75, 0xd0, 0x49, 0x6e, 0x14, 0x6b, 0xfd, 0x85, 0x21, 0xbe, 0x07,
	0xed, 0x9e, 0xfe, 0xfc, 0xd6, 0x3a, 0x40, 0x03, 0xbd, 0x77, 0xfa, 0x1f, 0xf4, 0x12, 0xbb, 0xb7,
	0x87, 0xe5, 0x8d, 0x5a, 0x7f, 0x36, 0xc1, 0xd6, 0xa2, 0x6e, 0x70, 0xba, 0x20, 0x44, 0x45, 0x08,
	0xf1, 0xd5, 0x03, 0xe5, 0x5f, 0x4a, 0x19, 0x02, 0x1a, 0xd2, 0xe6, 0x4a, 0x55, 0x64, 0xfd, 0x7c,
	0xb9, 0xac, 0xbc, 0x2c, 0x69, 0xb6, 0xa3, 0x86, 0xf3, 0xbd, 0x5b, 0xbf, 0xd4, 0xc0, 0x46, 0xf9,
	0x78, 0xf0, 0x53, 0x50, 0xa7, 0x63, 0x97, 0xbc, 0x14, 0x3d, 0x5a, 0xef, 0x6f, 0x15, 0x0a, 0x0b,
	0xb7, 0x86, 0x65, 0x18, 0x9a, 0x60, 0x75, 0x48, 0xa8, 0x37, 0x8c, 0x44, 0x87, 0xd6, 0xc5, 0xb5,
	0x36, 0x25, 0x50, 0xfa, 0xd3, 0x5a, 0xa8, 0x6e, 0xbd, 0x87, 0x33, 0x18, 0x7c, 0x0a, 0xea, 0x37,
	0xd4, 0x8d, 0x86, 0xa2, 0x33, 0xeb, 0xfd, 0x0f, 0x8b, 0x8d, 0x85, 0x3b, 0x87, 0x4b, 0x10, 0xfc,
	0x02, 0xac, 0xbb, 0x94, 0x07, 0x23, 0x7b, 0x7a, 0x31, 0x0d, 0x48, 0xd6, 0xa8, 0x29, 0x07, 0x4a,
	0x4e, 0x29, 0xa8, 0xe1, 0x32, 0x14, 0x9e, 0x00, 0x10, 0x85, 0xf6, 0x98, 0x07, 0x8c, 0x13, 0x57,
	0x34, 0x6c, 0xb3, 0xff, 0x41, 0x12, 0xab, 0xdb, 0x92, 0x58, 0xc4, 0x34, 0x5c, 0x02, 0xc2, 0x8f,
	0xc5, 0xb4, 0x59, 0x15, 0x79, 0x36, 0xe7, 0xa6, 0xcd, 0xdc, 0x58, 0x69, 0xbc, 0x6d, 0xac, 0x9c,
	0x81, 0xa6, 0x63, 0x47, 0xc4, 0x63, 0xe1, 0x34, 0xeb, 0x90, 0xfd, 0x62, 0xc4, 0xe5, 0x91, 0x7c,
	0x7a, 0xe9, 0xd6, 0xa5, 0x31, 0xd8, 0x7f, 0x61, 0xb4, 0xf7, 0xc5, 0x62, 0x0f, 0xcf, 0xb8, 0xad,
	0x37, 0x35, 0xd0, 0xc8, 0xd4, 0x83, 0x5f, 0x83, 0xcd, 0x52, 0x5d, 0xcc, 0x66, 0xa7, 0x92, 0xc4,
	0xea, 0xfb, 0x77, 0xaa, 0x28, 0x1d, 0xa0, 0xf3, 0xf0, 0x42, 0xcf, 0xea, 0x7d, 0xf5, 0xac, 0x3d,
	0x50, 0xcf, 0x95, 0xfb, 0xe8, 0x79, 0x02, 0x80, 0x7d, 0xc5, 0x1d, 0xca, 0xb9, 0x7d, 0x2e, 0x55,
	0x59, 0x2b, 0xab, 0x52, 0xc4, 0x34, 0x5c, 0x02, 0xc2, 0x2f, 0xc1, 0x46, 0x6e, 0xfd, 0x90, 0x3e,
	0xbf, 0xd4, 0xa7, 0xd4, 0x42, 0xe5, 0xa8, 0x86, 0xe7, 0xc0, 0x8b, 0x35, 0xd4, 0xb8, 0x7f, 0x0d,
	0xc9, 0x62, 0x68, 0xbe, 0xab, 0x18, 0xd6, 0xde, 0x52, 0x0c, 0xad, 0x5f, 0x97, 0xed, 0xac, 0x85,
	0x63, 0x57, 0x97, 0x2d, 0xfd, 0xda, 0x7d, 0x4b, 0x7f, 0x5e, 0x9b, 0x95, 0x65, 0xb5, 0xa9, 0x3f,
	0x44, 0x9b, 0xff, 0x51, 0xbb, 0xf5, 0xcf, 0x9e, 0x7f, 0xeb, 0x31, 0x83, 0xd8, 0x7c, 0xca, 0x02,
	0x6e, 0x8c, 0x98, 0x63, 0x8f, 0x4c, 0x87, 0x8d, 0xa3, 0xd0, 0x76, 0x22, 0x2e, 0x7f, 0x92, 0x42,
	0x12, 0x30, 0xe4, 0x33, 0x97, 0x8c, 0xb8, 0x99, 0x01, 0x4d, 0x61, 0x9a, 0xc5, 0x50, 0xbe, 0x5a,
	0x15, 0xd0, 0xa3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0x1c, 0xbc, 0xef, 0x84, 0x09, 0x00,
	0x00,
}
