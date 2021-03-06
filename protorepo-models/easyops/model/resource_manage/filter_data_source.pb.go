// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter_data_source.proto

package resource_manage

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
//数据来源
type FilterDataSource struct {
	//
	//类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//静态值
	Const *FilterDataSource_Const `protobuf:"bytes,2,opt,name=const,proto3" json:"const" form:"const"`
	//
	//映射,来源于实例的某个字段
	Mapping *FilterDataSource_Mapping `protobuf:"bytes,3,opt,name=mapping,proto3" json:"mapping" form:"mapping"`
	//
	//表达式, 返回值 number
	Expression *FilterDataSource_Expression `protobuf:"bytes,4,opt,name=expression,proto3" json:"expression" form:"expression"`
	//
	//时间, 返回值 时间戳(秒)
	Time *FilterDataSource_Time `protobuf:"bytes,5,opt,name=time,proto3" json:"time" form:"time"`
	//
	//值格式化函数, len 对数组, 字典, 字符串求长度;timeStamp 为将cmdb时间格式字符串 转为时间戳秒
	Format               string   `protobuf:"bytes,6,opt,name=format,proto3" json:"format" form:"format"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterDataSource) Reset()         { *m = FilterDataSource{} }
func (m *FilterDataSource) String() string { return proto.CompactTextString(m) }
func (*FilterDataSource) ProtoMessage()    {}
func (*FilterDataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2749531f2cd65f, []int{0}
}
func (m *FilterDataSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDataSource.Unmarshal(m, b)
}
func (m *FilterDataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDataSource.Marshal(b, m, deterministic)
}
func (m *FilterDataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDataSource.Merge(m, src)
}
func (m *FilterDataSource) XXX_Size() int {
	return xxx_messageInfo_FilterDataSource.Size(m)
}
func (m *FilterDataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDataSource.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDataSource proto.InternalMessageInfo

func (m *FilterDataSource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FilterDataSource) GetConst() *FilterDataSource_Const {
	if m != nil {
		return m.Const
	}
	return nil
}

func (m *FilterDataSource) GetMapping() *FilterDataSource_Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

func (m *FilterDataSource) GetExpression() *FilterDataSource_Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (m *FilterDataSource) GetTime() *FilterDataSource_Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *FilterDataSource) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type FilterDataSource_Const struct {
	//
	//值类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//值
	Value                *types.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FilterDataSource_Const) Reset()         { *m = FilterDataSource_Const{} }
func (m *FilterDataSource_Const) String() string { return proto.CompactTextString(m) }
func (*FilterDataSource_Const) ProtoMessage()    {}
func (*FilterDataSource_Const) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2749531f2cd65f, []int{0, 0}
}
func (m *FilterDataSource_Const) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDataSource_Const.Unmarshal(m, b)
}
func (m *FilterDataSource_Const) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDataSource_Const.Marshal(b, m, deterministic)
}
func (m *FilterDataSource_Const) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDataSource_Const.Merge(m, src)
}
func (m *FilterDataSource_Const) XXX_Size() int {
	return xxx_messageInfo_FilterDataSource_Const.Size(m)
}
func (m *FilterDataSource_Const) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDataSource_Const.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDataSource_Const proto.InternalMessageInfo

func (m *FilterDataSource_Const) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FilterDataSource_Const) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type FilterDataSource_Mapping struct {
	//
	//值路径, 点分法 owner.0.name
	KeyPath              string   `protobuf:"bytes,1,opt,name=key_path,json=keyPath,proto3" json:"key_path" form:"key_path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterDataSource_Mapping) Reset()         { *m = FilterDataSource_Mapping{} }
func (m *FilterDataSource_Mapping) String() string { return proto.CompactTextString(m) }
func (*FilterDataSource_Mapping) ProtoMessage()    {}
func (*FilterDataSource_Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2749531f2cd65f, []int{0, 1}
}
func (m *FilterDataSource_Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDataSource_Mapping.Unmarshal(m, b)
}
func (m *FilterDataSource_Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDataSource_Mapping.Marshal(b, m, deterministic)
}
func (m *FilterDataSource_Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDataSource_Mapping.Merge(m, src)
}
func (m *FilterDataSource_Mapping) XXX_Size() int {
	return xxx_messageInfo_FilterDataSource_Mapping.Size(m)
}
func (m *FilterDataSource_Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDataSource_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDataSource_Mapping proto.InternalMessageInfo

func (m *FilterDataSource_Mapping) GetKeyPath() string {
	if m != nil {
		return m.KeyPath
	}
	return ""
}

type FilterDataSource_Expression struct {
	//
	//运算符
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator" form:"operator"`
	//
	//左边数据
	Left *FilterDataSource `protobuf:"bytes,2,opt,name=left,proto3" json:"left" form:"left"`
	//
	//右边数据
	Right                *FilterDataSource `protobuf:"bytes,3,opt,name=right,proto3" json:"right" form:"right"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FilterDataSource_Expression) Reset()         { *m = FilterDataSource_Expression{} }
func (m *FilterDataSource_Expression) String() string { return proto.CompactTextString(m) }
func (*FilterDataSource_Expression) ProtoMessage()    {}
func (*FilterDataSource_Expression) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2749531f2cd65f, []int{0, 2}
}
func (m *FilterDataSource_Expression) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDataSource_Expression.Unmarshal(m, b)
}
func (m *FilterDataSource_Expression) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDataSource_Expression.Marshal(b, m, deterministic)
}
func (m *FilterDataSource_Expression) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDataSource_Expression.Merge(m, src)
}
func (m *FilterDataSource_Expression) XXX_Size() int {
	return xxx_messageInfo_FilterDataSource_Expression.Size(m)
}
func (m *FilterDataSource_Expression) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDataSource_Expression.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDataSource_Expression proto.InternalMessageInfo

func (m *FilterDataSource_Expression) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *FilterDataSource_Expression) GetLeft() *FilterDataSource {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *FilterDataSource_Expression) GetRight() *FilterDataSource {
	if m != nil {
		return m.Right
	}
	return nil
}

type FilterDataSource_Time struct {
	//
	//类型, static 指定时间, now 当前时间
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//type 为 static时需要填此属性
	Static string `protobuf:"bytes,2,opt,name=static,proto3" json:"static" form:"static"`
	//
	//时间偏移单位
	Unit string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit" form:"unit"`
	//
	//时间偏移量
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset" form:"offset"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterDataSource_Time) Reset()         { *m = FilterDataSource_Time{} }
func (m *FilterDataSource_Time) String() string { return proto.CompactTextString(m) }
func (*FilterDataSource_Time) ProtoMessage()    {}
func (*FilterDataSource_Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab2749531f2cd65f, []int{0, 3}
}
func (m *FilterDataSource_Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterDataSource_Time.Unmarshal(m, b)
}
func (m *FilterDataSource_Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterDataSource_Time.Marshal(b, m, deterministic)
}
func (m *FilterDataSource_Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterDataSource_Time.Merge(m, src)
}
func (m *FilterDataSource_Time) XXX_Size() int {
	return xxx_messageInfo_FilterDataSource_Time.Size(m)
}
func (m *FilterDataSource_Time) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterDataSource_Time.DiscardUnknown(m)
}

var xxx_messageInfo_FilterDataSource_Time proto.InternalMessageInfo

func (m *FilterDataSource_Time) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FilterDataSource_Time) GetStatic() string {
	if m != nil {
		return m.Static
	}
	return ""
}

func (m *FilterDataSource_Time) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *FilterDataSource_Time) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterDataSource)(nil), "resource_manage.FilterDataSource")
	proto.RegisterType((*FilterDataSource_Const)(nil), "resource_manage.FilterDataSource.Const")
	proto.RegisterType((*FilterDataSource_Mapping)(nil), "resource_manage.FilterDataSource.Mapping")
	proto.RegisterType((*FilterDataSource_Expression)(nil), "resource_manage.FilterDataSource.Expression")
	proto.RegisterType((*FilterDataSource_Time)(nil), "resource_manage.FilterDataSource.Time")
}

func init() { proto.RegisterFile("filter_data_source.proto", fileDescriptor_ab2749531f2cd65f) }

var fileDescriptor_ab2749531f2cd65f = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0x68, 0xda, 0xe1, 0x01, 0xdd, 0x8c, 0x40, 0x51, 0x84, 0x94, 0x61, 0x24, 0x58,
	0x25, 0xe6, 0x48, 0x70, 0x05, 0x17, 0x5c, 0x14, 0x98, 0x34, 0x4d, 0x08, 0x64, 0x10, 0x48, 0xdc,
	0x54, 0x6e, 0xe6, 0xa4, 0xd6, 0x92, 0x38, 0xb2, 0x1d, 0x44, 0x5f, 0x88, 0xe7, 0xe1, 0x09, 0x72,
	0xc3, 0x1b, 0xe4, 0x09, 0x50, 0x8e, 0xd3, 0xf5, 0x0f, 0x17, 0xed, 0x5d, 0x7d, 0xce, 0xf7, 0xfd,
	0xec, 0x7e, 0xe7, 0x28, 0xc8, 0x4f, 0x64, 0x66, 0x85, 0x9e, 0x5e, 0x71, 0xcb, 0xa7, 0x46, 0x55,
	0x3a, 0x16, 0xb4, 0xd4, 0xca, 0x2a, 0x3c, 0xd2, 0xc2, 0x9d, 0xa7, 0x39, 0x2f, 0x78, 0x2a, 0x82,
	0xb3, 0x54, 0xda, 0x79, 0x35, 0xa3, 0xb1, 0xca, 0xa3, 0x54, 0xa5, 0x2a, 0x02, 0xdd, 0xac, 0x4a,
	0xe0, 0x04, 0x07, 0xf8, 0xe5, 0xfc, 0xc1, 0xe3, 0x54, 0xa9, 0x34, 0x13, 0x2b, 0x95, 0xb1, 0xba,
	0x8a, 0xad, 0xeb, 0x92, 0xbf, 0x43, 0x74, 0x74, 0x0e, 0x57, 0xbf, 0xe7, 0x96, 0x7f, 0x81, 0x8b,
	0xf0, 0x53, 0xd4, 0xb7, 0x8b, 0x52, 0xf8, 0xbd, 0x93, 0xde, 0xe9, 0x9d, 0xc9, 0xa8, 0xa9, 0xc3,
	0xc3, 0x44, 0xe9, 0xfc, 0x0d, 0x69, 0xab, 0x84, 0x41, 0x13, 0x7f, 0x42, 0x5e, 0xac, 0x0a, 0x63,
	0xfd, 0x5b, 0x27, 0xbd, 0xd3, 0xc3, 0x97, 0xcf, 0xe9, 0xd6, 0x3b, 0xe9, 0x36, 0x96, 0xbe, 0x6b,
	0xe5, 0x93, 0xa3, 0xa6, 0x0e, 0xef, 0x3a, 0x1c, 0xf8, 0x09, 0x73, 0x1c, 0xfc, 0x1d, 0x0d, 0x73,
	0x5e, 0x96, 0xb2, 0x48, 0xfd, 0xdb, 0x80, 0x1c, 0xef, 0x46, 0x7e, 0x74, 0x86, 0x09, 0x6e, 0xea,
	0xf0, 0xbe, 0x83, 0x76, 0x0c, 0xc2, 0x96, 0x34, 0x1c, 0x23, 0x24, 0x7e, 0x95, 0x5a, 0x18, 0x23,
	0x55, 0xe1, 0xf7, 0x81, 0xfd, 0x62, 0x37, 0xfb, 0xc3, 0x8d, 0x67, 0xf2, 0xb0, 0xa9, 0xc3, 0x63,
	0x87, 0x5f, 0x91, 0x08, 0x5b, 0xc3, 0xe2, 0x4b, 0xd4, 0xb7, 0x32, 0x17, 0xbe, 0x07, 0xf8, 0x67,
	0xbb, 0xf1, 0x5f, 0x65, 0x2e, 0x36, 0xb2, 0x95, 0x39, 0x64, 0x2b, 0x73, 0x81, 0xc7, 0x68, 0xd0,
	0x16, 0xb9, 0xf5, 0x07, 0x30, 0x82, 0xe3, 0xa6, 0x0e, 0xef, 0x39, 0x99, 0xab, 0x13, 0xd6, 0x09,
	0x82, 0x0c, 0x79, 0x90, 0xeb, 0x7e, 0x43, 0x7b, 0x8b, 0xbc, 0x9f, 0x3c, 0xab, 0x44, 0x37, 0xb4,
	0x47, 0xd4, 0x2d, 0x07, 0x5d, 0x2e, 0x07, 0xfd, 0xd6, 0x76, 0xd7, 0x67, 0x04, 0x72, 0xc2, 0x9c,
	0x2d, 0x78, 0x8d, 0x86, 0x5d, 0xe4, 0x98, 0xa2, 0x83, 0x6b, 0xb1, 0x98, 0x96, 0xdc, 0xce, 0xbb,
	0x3b, 0x1f, 0x34, 0x75, 0x38, 0x72, 0xae, 0x65, 0x87, 0xb0, 0xe1, 0xb5, 0x58, 0x7c, 0xe6, 0x76,
	0x1e, 0xfc, 0xe9, 0x21, 0xb4, 0x8a, 0x14, 0x47, 0xe8, 0x40, 0x95, 0x42, 0x73, 0xab, 0xf4, 0xff,
	0xf6, 0x65, 0x87, 0xb0, 0x1b, 0x11, 0x3e, 0x47, 0xfd, 0x4c, 0x24, 0xcb, 0x75, 0x7b, 0xb2, 0x33,
	0xe0, 0xf5, 0x08, 0x5a, 0x23, 0x61, 0xe0, 0xc7, 0x17, 0xc8, 0xd3, 0x32, 0x9d, 0xdb, 0x6e, 0xc9,
	0xf6, 0x00, 0xad, 0xa5, 0x01, 0x4e, 0xc2, 0x1c, 0x21, 0xf8, 0xdd, 0x43, 0xfd, 0x76, 0x8c, 0xfb,
	0x65, 0x3f, 0x46, 0x03, 0x63, 0xb9, 0x95, 0x31, 0xfc, 0x85, 0x8d, 0xa1, 0xba, 0x3a, 0x61, 0x9d,
	0xa0, 0xe5, 0x55, 0x85, 0x74, 0x4f, 0xdc, 0xe0, 0xb5, 0x55, 0xc2, 0xa0, 0xd9, 0xf2, 0x54, 0x92,
	0x18, 0x61, 0x61, 0xa5, 0xbd, 0x75, 0x9e, 0xab, 0x13, 0xd6, 0x09, 0x26, 0x97, 0x3f, 0x2e, 0x52,
	0x45, 0x05, 0x37, 0x0b, 0x55, 0x1a, 0x9a, 0xa9, 0x98, 0x67, 0x51, 0xac, 0x0a, 0xab, 0x79, 0x6c,
	0x8d, 0xfb, 0x32, 0x68, 0x51, 0xaa, 0xb3, 0x5c, 0x5d, 0x89, 0xcc, 0x44, 0x9d, 0x30, 0x82, 0x63,
	0xb4, 0x95, 0xcf, 0x6c, 0x00, 0xfa, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x8e, 0xd4,
	0xa9, 0xb3, 0x04, 0x00, 0x00,
}
