// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: condition.proto

package inspection

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//阈值判定条件
type InspectionCondition struct {
	//
	//比较器的类型
	ComparatorType string `protobuf:"bytes,1,opt,name=comparatorType,proto3" json:"comparatorType" form:"comparatorType"`
	//
	//告警等级。0、5、10 分别对应 notice、warning、emergency 三个等级
	Level int32 `protobuf:"varint,2,opt,name=level,proto3" json:"level" form:"level"`
	//
	//一个 level 内最大值。仅当 type == int 时有效
	MaxValue float32 `protobuf:"fixed32,3,opt,name=maxValue,proto3" json:"maxValue" form:"maxValue"`
	//
	//一个 level 内最小值。仅当 type == int 时有效
	MinValue float32 `protobuf:"fixed32,4,opt,name=minValue,proto3" json:"minValue" form:"minValue"`
	//
	//阈值。仅当 type == string 时有效
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionCondition) Reset()         { *m = InspectionCondition{} }
func (m *InspectionCondition) String() string { return proto.CompactTextString(m) }
func (*InspectionCondition) ProtoMessage()    {}
func (*InspectionCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_02a70b272fae186f, []int{0}
}
func (m *InspectionCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionCondition.Unmarshal(m, b)
}
func (m *InspectionCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionCondition.Marshal(b, m, deterministic)
}
func (m *InspectionCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionCondition.Merge(m, src)
}
func (m *InspectionCondition) XXX_Size() int {
	return xxx_messageInfo_InspectionCondition.Size(m)
}
func (m *InspectionCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionCondition.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionCondition proto.InternalMessageInfo

func (m *InspectionCondition) GetComparatorType() string {
	if m != nil {
		return m.ComparatorType
	}
	return ""
}

func (m *InspectionCondition) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *InspectionCondition) GetMaxValue() float32 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *InspectionCondition) GetMinValue() float32 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *InspectionCondition) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*InspectionCondition)(nil), "inspection.InspectionCondition")
}

func init() { proto.RegisterFile("condition.proto", fileDescriptor_02a70b272fae186f) }

var fileDescriptor_02a70b272fae186f = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4e, 0x84, 0x30,
	0x10, 0x87, 0x03, 0x8a, 0xd1, 0xc6, 0xb8, 0x86, 0x8d, 0x09, 0x7a, 0x81, 0x70, 0x30, 0x5c, 0x96,
	0x1e, 0xbc, 0x79, 0x13, 0x8d, 0x89, 0x57, 0x62, 0x3c, 0x78, 0x2b, 0xdd, 0x2e, 0x92, 0x14, 0xa6,
	0x29, 0x85, 0xb8, 0xaf, 0xe0, 0x43, 0xf2, 0x10, 0x3c, 0x81, 0x61, 0xca, 0xae, 0x7f, 0xf6, 0x36,
	0x33, 0xbf, 0xef, 0x6b, 0xa6, 0x2d, 0x59, 0x70, 0x68, 0xd6, 0x95, 0xa9, 0xa0, 0x49, 0x95, 0x06,
	0x03, 0x3e, 0xa9, 0x9a, 0x56, 0x09, 0x3e, 0x4d, 0x6e, 0x56, 0x65, 0x65, 0x3e, 0xba, 0x22, 0xe5,
	0x50, 0xd3, 0x12, 0x4a, 0xa0, 0x88, 0x14, 0xdd, 0x06, 0x3b, 0x6c, 0xb0, 0xb2, 0x6a, 0xfc, 0xe5,
	0x92, 0xe5, 0xcb, 0xde, 0x7e, 0xdc, 0x1d, 0xec, 0x3f, 0x90, 0x0b, 0x0e, 0xb5, 0x62, 0x9a, 0x19,
	0xd0, 0xaf, 0x5b, 0x25, 0x02, 0x27, 0x72, 0x92, 0xb3, 0xec, 0x7a, 0x1c, 0xc2, 0xab, 0x0d, 0xe8,
	0xfa, 0x3e, 0xfe, 0x9b, 0xc7, 0xf9, 0x3f, 0xc1, 0xbf, 0x25, 0x9e, 0x14, 0xbd, 0x90, 0x81, 0x1b,
	0x39, 0x89, 0x97, 0x5d, 0x8e, 0x43, 0x78, 0x6e, 0x4d, 0x1c, 0xc7, 0xb9, 0x8d, 0x7d, 0x4a, 0x4e,
	0x6b, 0xf6, 0xf9, 0xc6, 0x64, 0x27, 0x82, 0xa3, 0xc8, 0x49, 0xdc, 0x6c, 0x39, 0x0e, 0xe1, 0xc2,
	0xa2, 0xbb, 0x24, 0xce, 0xf7, 0x10, 0x0a, 0x55, 0x63, 0x85, 0xe3, 0x03, 0x61, 0x4e, 0x26, 0x61,
	0x2e, 0xa7, 0x4d, 0x7a, 0xa4, 0x3d, 0xbc, 0xc3, 0xaf, 0x4d, 0x7a, 0x8b, 0xda, 0x38, 0x7b, 0x7e,
	0x7f, 0x2a, 0x21, 0x15, 0xac, 0xdd, 0x82, 0x6a, 0x53, 0x09, 0x9c, 0x49, 0xca, 0xa1, 0x31, 0x9a,
	0x71, 0xd3, 0xda, 0x87, 0xd4, 0x42, 0xc1, 0xaa, 0x86, 0xb5, 0x90, 0x2d, 0x9d, 0x41, 0x8a, 0x2d,
	0xfd, 0xf9, 0x83, 0xe2, 0x04, 0xd1, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x2d, 0x42,
	0x10, 0xa9, 0x01, 0x00, 0x00,
}