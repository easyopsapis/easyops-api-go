// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_metrics_schema_field.proto

package metadata_center

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
//指标表字段
type StreamMetricsSchemaField struct {
	//
	//字段名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//字段类型
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricsSchemaField) Reset()         { *m = StreamMetricsSchemaField{} }
func (m *StreamMetricsSchemaField) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsSchemaField) ProtoMessage()    {}
func (*StreamMetricsSchemaField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f11b50e2286f71, []int{0}
}
func (m *StreamMetricsSchemaField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsSchemaField.Unmarshal(m, b)
}
func (m *StreamMetricsSchemaField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsSchemaField.Marshal(b, m, deterministic)
}
func (m *StreamMetricsSchemaField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsSchemaField.Merge(m, src)
}
func (m *StreamMetricsSchemaField) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsSchemaField.Size(m)
}
func (m *StreamMetricsSchemaField) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsSchemaField.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsSchemaField proto.InternalMessageInfo

func (m *StreamMetricsSchemaField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamMetricsSchemaField) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamMetricsSchemaField)(nil), "metadata_center.StreamMetricsSchemaField")
}

func init() { proto.RegisterFile("stream_metrics_schema_field.proto", fileDescriptor_e4f11b50e2286f71) }

var fileDescriptor_e4f11b50e2286f71 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0x84, 0x90, 0x08, 0x43, 0xa5, 0x4e, 0x11, 0x0b, 0x10, 0x16, 0x96, 0xc6, 0x03,
	0x1b, 0x63, 0x07, 0x24, 0x84, 0x58, 0xda, 0x8d, 0xc5, 0xba, 0xd8, 0x17, 0x37, 0x92, 0x2f, 0x67,
	0xd9, 0xd7, 0xa1, 0x2f, 0xcb, 0x43, 0xf0, 0x04, 0xc8, 0x17, 0xa6, 0x6e, 0xfe, 0xff, 0xff, 0xb3,
	0x3e, 0xbb, 0x79, 0x2a, 0x92, 0x11, 0xc8, 0x12, 0x4a, 0x9e, 0x5c, 0xb1, 0xc5, 0x1d, 0x91, 0xc0,
	0x8e, 0x13, 0x46, 0xdf, 0xa7, 0xcc, 0xc2, 0x9b, 0x35, 0xa1, 0x80, 0x07, 0x01, 0xeb, 0x70, 0x16,
	0xcc, 0xf7, 0xdb, 0x30, 0xc9, 0xf1, 0x34, 0xf4, 0x8e, 0xc9, 0x04, 0x0e, 0x6c, 0x94, 0x1b, 0x4e,
	0xa3, 0x26, 0x0d, 0x7a, 0x5a, 0xee, 0x77, 0xbe, 0x69, 0x0f, 0x2a, 0xf9, 0x5a, 0x1c, 0x07, 0x55,
	0xbc, 0x57, 0xc3, 0xe6, 0xb9, 0xb9, 0x9e, 0x81, 0xb0, 0x5d, 0x3d, 0xae, 0x5e, 0x6e, 0x77, 0xeb,
	0xdf, 0x9f, 0x87, 0xbb, 0x91, 0x33, 0xbd, 0x75, 0xb5, 0xed, 0xf6, 0x3a, 0x56, 0x48, 0xce, 0x09,
	0xdb, 0xab, 0x4b, 0xa8, 0xb6, 0xdd, 0x5e, 0xc7, 0xdd, 0xe7, 0xf7, 0x47, 0xe0, 0x1e, 0xa1, 0x9c,
	0x39, 0x95, 0x3e, 0xb2, 0x83, 0x68, 0x1c, 0xcf, 0x92, 0xc1, 0x49, 0x59, 0x5e, 0x98, 0x31, 0xf1,
	0x96, 0xd8, 0x63, 0x2c, 0xe6, 0x1f, 0x34, 0x1a, 0xcd, 0xc5, 0x0f, 0x87, 0x1b, 0xe5, 0x5f, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xa4, 0x3d, 0x0b, 0x1e, 0x01, 0x00, 0x00,
}
