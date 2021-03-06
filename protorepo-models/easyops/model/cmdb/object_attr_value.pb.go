// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: object_attr_value.proto

package cmdb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//模型属性类型
type ObjectAttrValue struct {
	//
	//数据类型, FK与FKs已废弃
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//type是enum时是枚举值数组, struct和structs没有regex, 其他的type时是正则表达式
	Regex *types.Value `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex" form:"regex"`
	//
	//默认值类型( value: 具体的值, function:可执行函数, auto-increment-id:自增ID, series-number:定长序列号)
	DefaultType string `protobuf:"bytes,3,opt,name=default_type,json=defaultType,proto3" json:"default_type" form:"default_type"`
	//
	//默认值
	Default *types.Value `protobuf:"bytes,4,opt,name=default,proto3" json:"default" form:"default"`
	//
	//结构体字段定义: 当type 为 struct 和 structs 时候为必填
	StructDefine []*ObjectAttrValueStruct `protobuf:"bytes,5,rep,name=struct_define,json=structDefine,proto3" json:"struct_define" form:"struct_define"`
	//
	//字符串模式定义: 多行文本和普通字符串
	Mode string `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode" form:"mode"`
	//
	//自增ID和流水号的前缀: 当default_type 是series-number和 auto-increment-id时候为必填
	Prefix string `protobuf:"bytes,7,opt,name=prefix,proto3" json:"prefix" form:"prefix"`
	//
	//自增ID和流水号的起始值
	StartValue int32 `protobuf:"varint,8,opt,name=start_value,json=startValue,proto3" json:"start_value" form:"start_value"`
	//
	//流水号的固定长度
	SeriesNumberLength   int32    `protobuf:"varint,9,opt,name=series_number_length,json=seriesNumberLength,proto3" json:"series_number_length" form:"series_number_length"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectAttrValue) Reset()         { *m = ObjectAttrValue{} }
func (m *ObjectAttrValue) String() string { return proto.CompactTextString(m) }
func (*ObjectAttrValue) ProtoMessage()    {}
func (*ObjectAttrValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9c6d258de04327, []int{0}
}
func (m *ObjectAttrValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectAttrValue.Unmarshal(m, b)
}
func (m *ObjectAttrValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectAttrValue.Marshal(b, m, deterministic)
}
func (m *ObjectAttrValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectAttrValue.Merge(m, src)
}
func (m *ObjectAttrValue) XXX_Size() int {
	return xxx_messageInfo_ObjectAttrValue.Size(m)
}
func (m *ObjectAttrValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectAttrValue.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectAttrValue proto.InternalMessageInfo

func (m *ObjectAttrValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ObjectAttrValue) GetRegex() *types.Value {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *ObjectAttrValue) GetDefaultType() string {
	if m != nil {
		return m.DefaultType
	}
	return ""
}

func (m *ObjectAttrValue) GetDefault() *types.Value {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *ObjectAttrValue) GetStructDefine() []*ObjectAttrValueStruct {
	if m != nil {
		return m.StructDefine
	}
	return nil
}

func (m *ObjectAttrValue) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ObjectAttrValue) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ObjectAttrValue) GetStartValue() int32 {
	if m != nil {
		return m.StartValue
	}
	return 0
}

func (m *ObjectAttrValue) GetSeriesNumberLength() int32 {
	if m != nil {
		return m.SeriesNumberLength
	}
	return 0
}

func init() {
	proto.RegisterType((*ObjectAttrValue)(nil), "cmdb.ObjectAttrValue")
}

func init() { proto.RegisterFile("object_attr_value.proto", fileDescriptor_2a9c6d258de04327) }

var fileDescriptor_2a9c6d258de04327 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x67, 0xed, 0xee, 0xd6, 0x4e, 0xb6, 0x56, 0xc7, 0xd2, 0x0e, 0xad, 0x30, 0xcb, 0x78, 0x59,
	0x0f, 0x9d, 0x40, 0x3d, 0x08, 0x3d, 0x14, 0x0d, 0x3d, 0x8a, 0xe2, 0x28, 0x1e, 0x0a, 0x12, 0x92,
	0xec, 0x4b, 0xba, 0x92, 0xec, 0x84, 0xc9, 0x44, 0xda, 0x6f, 0xe5, 0x27, 0xca, 0x87, 0xc8, 0x27,
	0x90, 0x79, 0x93, 0xd5, 0xb6, 0x8a, 0x97, 0xde, 0xf2, 0x7b, 0xbf, 0x7f, 0xc3, 0x7b, 0x21, 0x87,
	0x3a, 0xfd, 0x0e, 0x99, 0x8d, 0x13, 0x6b, 0x4d, 0xfc, 0x23, 0x29, 0x5b, 0x90, 0xb5, 0xd1, 0x56,
	0xd3, 0x71, 0x56, 0x2d, 0xd3, 0xa3, 0x93, 0x62, 0x65, 0xaf, 0xda, 0x54, 0x66, 0xba, 0x0a, 0x0b,
	0x5d, 0xe8, 0x10, 0xc9, 0xb4, 0xcd, 0x11, 0x21, 0xc0, 0x2f, 0x6f, 0x3a, 0xfa, 0x56, 0x68, 0x09,
	0x49, 0x73, 0xa3, 0xeb, 0x46, 0x96, 0x3a, 0x4b, 0xca, 0x30, 0xd3, 0x6b, 0x6b, 0x92, 0xcc, 0x36,
	0xde, 0x69, 0xa0, 0xd6, 0x27, 0x95, 0x5e, 0x42, 0xd9, 0x84, 0x83, 0x30, 0x44, 0x18, 0xba, 0xba,
	0xf0, 0xaf, 0xc7, 0xc4, 0x8d, 0x35, 0x6d, 0x66, 0x87, 0xf8, 0x17, 0x85, 0xd6, 0x45, 0x09, 0x7f,
	0x1e, 0x71, 0x9b, 0x15, 0x3f, 0xc7, 0x64, 0xef, 0x23, 0x06, 0xbc, 0xb3, 0xd6, 0x7c, 0x75, 0x76,
	0xfa, 0x92, 0x8c, 0xed, 0x4d, 0x0d, 0x6c, 0x34, 0x1f, 0x2d, 0x76, 0xa2, 0xbd, 0xbe, 0xe3, 0x41,
	0xae, 0x4d, 0x75, 0x26, 0xdc, 0x54, 0x28, 0x24, 0xe9, 0x39, 0x99, 0x18, 0x28, 0xe0, 0x9a, 0x3d,
	0x9a, 0x8f, 0x16, 0xc1, 0xe9, 0x81, 0xf4, 0x35, 0x72, 0x53, 0x23, 0x31, 0x2b, 0x7a, 0xda, 0x77,
	0x7c, 0xe6, 0xdd, 0x28, 0x17, 0xca, 0xdb, 0xe8, 0x19, 0x99, 0x2d, 0x21, 0x4f, 0xda, 0xd2, 0xc6,
	0x58, 0xb6, 0x85, 0x65, 0x87, 0x7d, 0xc7, 0x9f, 0x7b, 0xf9, 0x6d, 0x56, 0xa8, 0x60, 0x80, 0x5f,
	0x5c, 0xf7, 0x05, 0xd9, 0x1e, 0x20, 0x1b, 0xff, 0xb7, 0x9d, 0xf6, 0x1d, 0x7f, 0x72, 0x27, 0x4e,
	0xa8, 0x8d, 0x95, 0x5e, 0x92, 0x5d, 0xbf, 0x8a, 0x78, 0x09, 0xf9, 0x6a, 0x0d, 0x6c, 0x32, 0xdf,
	0x5a, 0x04, 0xa7, 0xc7, 0xd2, 0x6d, 0x55, 0xde, 0x5b, 0xca, 0x67, 0x54, 0x46, 0xac, 0xef, 0xf8,
	0xbe, 0x0f, 0xbc, 0xe3, 0x15, 0x6a, 0xe6, 0xf1, 0x05, 0x42, 0xb7, 0x42, 0x77, 0x21, 0x36, 0xbd,
	0xbf, 0x42, 0x37, 0x15, 0x0a, 0x49, 0xfa, 0x8a, 0x4c, 0x6b, 0x03, 0xf9, 0xea, 0x9a, 0x6d, 0xa3,
	0xec, 0x59, 0xdf, 0xf1, 0x5d, 0x2f, 0xf3, 0x73, 0xa1, 0x06, 0x01, 0x7d, 0x43, 0x82, 0xc6, 0x26,
	0xc6, 0xfa, 0x03, 0xb3, 0xc7, 0xf3, 0xd1, 0x62, 0x12, 0x1d, 0xf4, 0x1d, 0xa7, 0x9b, 0xc7, 0xfc,
	0x26, 0x85, 0x22, 0x88, 0xfc, 0x2d, 0x3f, 0x91, 0xfd, 0x06, 0xcc, 0x0a, 0x9a, 0x78, 0xdd, 0x56,
	0x29, 0x98, 0xb8, 0x84, 0x75, 0x61, 0xaf, 0xd8, 0x0e, 0x26, 0xf0, 0xbe, 0xe3, 0xc7, 0x43, 0xc2,
	0x3f, 0x54, 0x42, 0x51, 0x3f, 0xfe, 0x80, 0xd3, 0xf7, 0x38, 0x8c, 0xde, 0x5e, 0x9e, 0x3f, 0xec,
	0x8f, 0x4d, 0xa7, 0x28, 0x7a, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x09, 0x2c, 0xa1, 0x40, 0x48,
	0x03, 0x00, 0x00,
}
