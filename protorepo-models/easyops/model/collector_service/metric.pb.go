// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metric.proto

package collector_service

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
//采集指标
type CollectorMetric struct {
	//
	//指标名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//指标名
	KeyName string `protobuf:"bytes,2,opt,name=keyName,proto3" json:"keyName" form:"keyName"`
	//
	//标签
	Labels []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels" form:"labels"`
	//
	//指标id
	MetricId string `protobuf:"bytes,4,opt,name=metricId,proto3" json:"metricId" form:"metricId"`
	//
	//指标版本
	Version int32 `protobuf:"varint,5,opt,name=version,proto3" json:"version" form:"version"`
	//
	//指标key
	Key string `protobuf:"bytes,6,opt,name=key,proto3" json:"key" form:"key"`
	//
	//指标类型
	MetricType string `protobuf:"bytes,7,opt,name=metricType,proto3" json:"metricType" form:"metricType"`
	//
	//指标数据类型
	DataType string `protobuf:"bytes,8,opt,name=dataType,proto3" json:"dataType" form:"dataType"`
	//
	//所属Agent
	AgentType string `protobuf:"bytes,9,opt,name=agentType,proto3" json:"agentType" form:"agentType"`
	//
	//维度定义
	TagDefine []*CollectorMetric_TagDefine `protobuf:"bytes,10,rep,name=tagDefine,proto3" json:"tagDefine" form:"tagDefine"`
	//
	//参数定义
	ParamDefine []*CollectorMetric_ParamDefine `protobuf:"bytes,11,rep,name=paramDefine,proto3" json:"paramDefine" form:"paramDefine"`
	//
	//帮助
	Help string `protobuf:"bytes,12,opt,name=help,proto3" json:"help" form:"help"`
	//
	//单位
	Unit string `protobuf:"bytes,13,opt,name=unit,proto3" json:"unit" form:"unit"`
	//
	//采集指标实例版本号
	MetricVersion        int32    `protobuf:"varint,14,opt,name=metricVersion,proto3" json:"metricVersion" form:"metricVersion"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectorMetric) Reset()         { *m = CollectorMetric{} }
func (m *CollectorMetric) String() string { return proto.CompactTextString(m) }
func (*CollectorMetric) ProtoMessage()    {}
func (*CollectorMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_da41641f55bff5df, []int{0}
}
func (m *CollectorMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectorMetric.Unmarshal(m, b)
}
func (m *CollectorMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectorMetric.Marshal(b, m, deterministic)
}
func (m *CollectorMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectorMetric.Merge(m, src)
}
func (m *CollectorMetric) XXX_Size() int {
	return xxx_messageInfo_CollectorMetric.Size(m)
}
func (m *CollectorMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectorMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CollectorMetric proto.InternalMessageInfo

func (m *CollectorMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectorMetric) GetKeyName() string {
	if m != nil {
		return m.KeyName
	}
	return ""
}

func (m *CollectorMetric) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CollectorMetric) GetMetricId() string {
	if m != nil {
		return m.MetricId
	}
	return ""
}

func (m *CollectorMetric) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CollectorMetric) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CollectorMetric) GetMetricType() string {
	if m != nil {
		return m.MetricType
	}
	return ""
}

func (m *CollectorMetric) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *CollectorMetric) GetAgentType() string {
	if m != nil {
		return m.AgentType
	}
	return ""
}

func (m *CollectorMetric) GetTagDefine() []*CollectorMetric_TagDefine {
	if m != nil {
		return m.TagDefine
	}
	return nil
}

func (m *CollectorMetric) GetParamDefine() []*CollectorMetric_ParamDefine {
	if m != nil {
		return m.ParamDefine
	}
	return nil
}

func (m *CollectorMetric) GetHelp() string {
	if m != nil {
		return m.Help
	}
	return ""
}

func (m *CollectorMetric) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *CollectorMetric) GetMetricVersion() int32 {
	if m != nil {
		return m.MetricVersion
	}
	return 0
}

type CollectorMetric_TagDefine struct {
	//
	//名字
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//默认值
	Default string `protobuf:"bytes,2,opt,name=default,proto3" json:"default" form:"default"`
	//
	//是否只读
	ReadOnly             bool     `protobuf:"varint,3,opt,name=readOnly,proto3" json:"readOnly" form:"readOnly"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectorMetric_TagDefine) Reset()         { *m = CollectorMetric_TagDefine{} }
func (m *CollectorMetric_TagDefine) String() string { return proto.CompactTextString(m) }
func (*CollectorMetric_TagDefine) ProtoMessage()    {}
func (*CollectorMetric_TagDefine) Descriptor() ([]byte, []int) {
	return fileDescriptor_da41641f55bff5df, []int{0, 0}
}
func (m *CollectorMetric_TagDefine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectorMetric_TagDefine.Unmarshal(m, b)
}
func (m *CollectorMetric_TagDefine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectorMetric_TagDefine.Marshal(b, m, deterministic)
}
func (m *CollectorMetric_TagDefine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectorMetric_TagDefine.Merge(m, src)
}
func (m *CollectorMetric_TagDefine) XXX_Size() int {
	return xxx_messageInfo_CollectorMetric_TagDefine.Size(m)
}
func (m *CollectorMetric_TagDefine) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectorMetric_TagDefine.DiscardUnknown(m)
}

var xxx_messageInfo_CollectorMetric_TagDefine proto.InternalMessageInfo

func (m *CollectorMetric_TagDefine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectorMetric_TagDefine) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *CollectorMetric_TagDefine) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

type CollectorMetric_ParamDefine struct {
	//
	//名字
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//值类型
	ValueType string `protobuf:"bytes,2,opt,name=valueType,proto3" json:"valueType" form:"valueType"`
	//
	//默认值
	Default string `protobuf:"bytes,3,opt,name=default,proto3" json:"default" form:"default"`
	//
	//是否只读
	ReadOnly             bool     `protobuf:"varint,4,opt,name=readOnly,proto3" json:"readOnly" form:"readOnly"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectorMetric_ParamDefine) Reset()         { *m = CollectorMetric_ParamDefine{} }
func (m *CollectorMetric_ParamDefine) String() string { return proto.CompactTextString(m) }
func (*CollectorMetric_ParamDefine) ProtoMessage()    {}
func (*CollectorMetric_ParamDefine) Descriptor() ([]byte, []int) {
	return fileDescriptor_da41641f55bff5df, []int{0, 1}
}
func (m *CollectorMetric_ParamDefine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectorMetric_ParamDefine.Unmarshal(m, b)
}
func (m *CollectorMetric_ParamDefine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectorMetric_ParamDefine.Marshal(b, m, deterministic)
}
func (m *CollectorMetric_ParamDefine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectorMetric_ParamDefine.Merge(m, src)
}
func (m *CollectorMetric_ParamDefine) XXX_Size() int {
	return xxx_messageInfo_CollectorMetric_ParamDefine.Size(m)
}
func (m *CollectorMetric_ParamDefine) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectorMetric_ParamDefine.DiscardUnknown(m)
}

var xxx_messageInfo_CollectorMetric_ParamDefine proto.InternalMessageInfo

func (m *CollectorMetric_ParamDefine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectorMetric_ParamDefine) GetValueType() string {
	if m != nil {
		return m.ValueType
	}
	return ""
}

func (m *CollectorMetric_ParamDefine) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *CollectorMetric_ParamDefine) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func init() {
	proto.RegisterType((*CollectorMetric)(nil), "collector_service.CollectorMetric")
	proto.RegisterType((*CollectorMetric_TagDefine)(nil), "collector_service.CollectorMetric.TagDefine")
	proto.RegisterType((*CollectorMetric_ParamDefine)(nil), "collector_service.CollectorMetric.ParamDefine")
}

func init() { proto.RegisterFile("metric.proto", fileDescriptor_da41641f55bff5df) }

var fileDescriptor_da41641f55bff5df = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x65, 0x7f, 0x1a, 0x77, 0xeb, 0x36, 0x33, 0x50, 0xd4, 0x9b, 0x44, 0x06, 0x89,
	0x22, 0x2d, 0x09, 0x6c, 0x02, 0x09, 0x2e, 0xb8, 0xe8, 0xb8, 0x41, 0x68, 0x30, 0x59, 0xd3, 0x2e,
	0x40, 0xfc, 0x71, 0x13, 0x37, 0x8b, 0xe6, 0xc4, 0x51, 0xe2, 0x76, 0x2a, 0x88, 0x67, 0xe0, 0xa9,
	0x78, 0x8d, 0x20, 0xf1, 0x04, 0x28, 0x4f, 0x80, 0x62, 0x27, 0x69, 0xd6, 0x5d, 0x50, 0xee, 0x6a,
	0x7f, 0xbf, 0xef, 0x1c, 0x1f, 0xfb, 0x4b, 0xc1, 0x76, 0x44, 0x45, 0x1a, 0x7a, 0x4e, 0x92, 0x72,
	0xc1, 0xe1, 0xbe, 0xc7, 0x19, 0xa3, 0x9e, 0xe0, 0xe9, 0xe7, 0x8c, 0xa6, 0xb3, 0xd0, 0xa3, 0x03,
	0x3b, 0x08, 0xc5, 0xe5, 0x74, 0xec, 0x78, 0x3c, 0x72, 0x03, 0x1e, 0x70, 0x57, 0x92, 0xe3, 0xe9,
	0x44, 0xae, 0xe4, 0x42, 0xfe, 0x52, 0x15, 0x06, 0xcf, 0x5a, 0x78, 0x74, 0x1d, 0x8a, 0x2b, 0x7e,
	0xed, 0x06, 0xdc, 0x96, 0xa2, 0x3d, 0x23, 0x2c, 0xf4, 0x89, 0xe0, 0x69, 0xe6, 0x36, 0x3f, 0x95,
	0x0f, 0xfd, 0xe9, 0x82, 0xdd, 0x93, 0xba, 0xf9, 0xa9, 0x3c, 0x13, 0xbc, 0x0f, 0xd6, 0x63, 0x12,
	0x51, 0xa3, 0x63, 0x75, 0x86, 0xfa, 0x68, 0xb7, 0xc8, 0xcd, 0xde, 0x84, 0xa7, 0xd1, 0x0b, 0x54,
	0xee, 0x22, 0x2c, 0x45, 0x78, 0x08, 0xb6, 0xae, 0xe8, 0xfc, 0x6d, 0xc9, 0xad, 0x49, 0x0e, 0x16,
	0xb9, 0xd9, 0x57, 0x5c, 0x25, 0x20, 0x5c, 0x23, 0xf0, 0x11, 0xd8, 0x64, 0x64, 0x4c, 0x59, 0x66,
	0x68, 0x96, 0x36, 0xd4, 0x47, 0xfb, 0x45, 0x6e, 0xee, 0x28, 0x58, 0xed, 0x23, 0x5c, 0x01, 0xf0,
	0x04, 0x74, 0xd5, 0xdd, 0xbc, 0xf6, 0x8d, 0x75, 0x59, 0xf9, 0x61, 0x91, 0x9b, 0xbb, 0x0a, 0xae,
	0x15, 0xf4, 0xfb, 0x97, 0xb9, 0x07, 0xfa, 0x9f, 0x3e, 0x3c, 0xb6, 0x9f, 0x13, 0xfb, 0xeb, 0xc7,
	0x6f, 0x4f, 0x8e, 0xbf, 0x3f, 0xc0, 0x8d, 0xb1, 0x3c, 0xdd, 0x8c, 0xa6, 0x59, 0xc8, 0x63, 0x63,
	0xc3, 0xea, 0x0c, 0x37, 0xda, 0xa7, 0xab, 0x04, 0x84, 0x6b, 0x04, 0x5a, 0x40, 0xbb, 0xa2, 0x73,
	0x63, 0x53, 0x76, 0xeb, 0x17, 0xb9, 0x09, 0x9a, 0x39, 0x10, 0x2e, 0x25, 0xf8, 0x14, 0x00, 0x55,
	0xfb, 0x7c, 0x9e, 0x50, 0x63, 0x4b, 0x82, 0x77, 0x8b, 0xdc, 0xdc, 0x6f, 0x1f, 0xab, 0xd4, 0x10,
	0x6e, 0x81, 0xd0, 0x05, 0x5d, 0x9f, 0x08, 0x22, 0x4d, 0x5d, 0x69, 0xba, 0xb3, 0x98, 0xa5, 0x56,
	0x10, 0x6e, 0x20, 0x78, 0x04, 0x74, 0x12, 0xd0, 0x58, 0x48, 0x87, 0x2e, 0x1d, 0x07, 0x45, 0x6e,
	0xee, 0x29, 0x47, 0x23, 0x21, 0xbc, 0xc0, 0xe0, 0x17, 0xa0, 0x0b, 0x12, 0xbc, 0xa2, 0x93, 0x30,
	0xa6, 0x06, 0xb0, 0xb4, 0x61, 0xef, 0xe8, 0xd0, 0xb9, 0x15, 0x28, 0x67, 0xe9, 0x95, 0x9d, 0xf3,
	0xda, 0xd3, 0xee, 0xd0, 0x14, 0x42, 0x78, 0x51, 0x14, 0x5e, 0x82, 0x5e, 0x42, 0x52, 0x12, 0x55,
	0x3d, 0x7a, 0xb2, 0x87, 0xb3, 0x42, 0x8f, 0xb3, 0x85, 0x6b, 0x74, 0xaf, 0xc8, 0x4d, 0xa8, 0xba,
	0xb4, 0x8a, 0x21, 0xdc, 0x2e, 0x5d, 0x46, 0xef, 0x92, 0xb2, 0xc4, 0xd8, 0x5e, 0x8e, 0x5e, 0xb9,
	0x8b, 0xb0, 0x14, 0x4b, 0x68, 0x1a, 0x87, 0xc2, 0xd8, 0x59, 0x86, 0xca, 0x5d, 0x84, 0xa5, 0x08,
	0x5f, 0x82, 0x1d, 0xf5, 0x10, 0x17, 0x55, 0x0e, 0xfa, 0x32, 0x07, 0x46, 0x91, 0x9b, 0x07, 0xed,
	0x47, 0xbb, 0xa8, 0xd3, 0x70, 0x13, 0x1f, 0xfc, 0xe8, 0x00, 0xbd, 0xb9, 0xa2, 0x95, 0x3f, 0x09,
	0x9f, 0x4e, 0xc8, 0x94, 0x89, 0xdb, 0x9f, 0x44, 0x25, 0x20, 0x5c, 0x23, 0x65, 0x36, 0x52, 0x4a,
	0xfc, 0x77, 0x31, 0x9b, 0x1b, 0x9a, 0xd5, 0x19, 0x76, 0xdb, 0xd9, 0xa8, 0x15, 0x84, 0x1b, 0x68,
	0xf0, 0xb3, 0x03, 0x7a, 0x67, 0x37, 0xef, 0xea, 0xdf, 0x67, 0x3a, 0x02, 0xfa, 0x8c, 0xb0, 0x29,
	0x95, 0x81, 0x5a, 0x5b, 0x0e, 0x54, 0x23, 0x21, 0xbc, 0xc0, 0xda, 0x73, 0x68, 0xff, 0x37, 0xc7,
	0xfa, 0x0a, 0x73, 0x8c, 0x4e, 0xdf, 0xbf, 0x09, 0xb8, 0x43, 0x49, 0x36, 0xe7, 0x49, 0xe6, 0x30,
	0xee, 0x11, 0xe6, 0x7a, 0x3c, 0x16, 0x29, 0xf1, 0x44, 0xa6, 0xfe, 0xe6, 0x52, 0x9a, 0x70, 0x3b,
	0xe2, 0x3e, 0x65, 0x99, 0x5b, 0x81, 0xae, 0x5c, 0xba, 0xb7, 0x32, 0x37, 0xde, 0x94, 0x8e, 0xe3,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xe4, 0x54, 0x6f, 0x52, 0x05, 0x00, 0x00,
}
