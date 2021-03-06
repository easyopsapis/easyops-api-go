// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_metric_schema.proto

package metadata_center

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
//指标项
type StreamMetricSchema struct {
	//
	//id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//org
	Org int32 `protobuf:"varint,2,opt,name=org,proto3" json:"org" form:"org"`
	//
	//指标版本
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version" form:"version"`
	//
	//指标统计类型
	MetricType string `protobuf:"bytes,4,opt,name=metricType,proto3" json:"metricType" form:"metricType"`
	//
	//指标数据类型
	DataType             string   `protobuf:"bytes,5,opt,name=dataType,proto3" json:"dataType" form:"dataType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMetricSchema) Reset()         { *m = StreamMetricSchema{} }
func (m *StreamMetricSchema) String() string { return proto.CompactTextString(m) }
func (*StreamMetricSchema) ProtoMessage()    {}
func (*StreamMetricSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_e908f2546dace26c, []int{0}
}
func (m *StreamMetricSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricSchema.Unmarshal(m, b)
}
func (m *StreamMetricSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricSchema.Marshal(b, m, deterministic)
}
func (m *StreamMetricSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricSchema.Merge(m, src)
}
func (m *StreamMetricSchema) XXX_Size() int {
	return xxx_messageInfo_StreamMetricSchema.Size(m)
}
func (m *StreamMetricSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricSchema.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricSchema proto.InternalMessageInfo

func (m *StreamMetricSchema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamMetricSchema) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *StreamMetricSchema) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StreamMetricSchema) GetMetricType() string {
	if m != nil {
		return m.MetricType
	}
	return ""
}

func (m *StreamMetricSchema) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamMetricSchema)(nil), "metadata_center.StreamMetricSchema")
}

func init() { proto.RegisterFile("stream_metric_schema.proto", fileDescriptor_e908f2546dace26c) }

var fileDescriptor_e908f2546dace26c = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0x32, 0x41,
	0x1c, 0xc6, 0x71, 0x7d, 0x7d, 0xdf, 0xd7, 0x39, 0x28, 0x8e, 0x04, 0x8b, 0x97, 0x95, 0xc5, 0x83,
	0x87, 0x76, 0x27, 0xca, 0x82, 0xba, 0xe5, 0x21, 0x88, 0xe8, 0xb2, 0x76, 0x2a, 0x4a, 0xc6, 0xdd,
	0x71, 0x1d, 0xda, 0xf1, 0xbf, 0xcc, 0x8c, 0x8a, 0x44, 0x97, 0x3e, 0xe8, 0x06, 0x7d, 0x84, 0xfd,
	0x04, 0xe1, 0x8c, 0x9a, 0x74, 0x9b, 0x3f, 0xcf, 0xef, 0x07, 0x0f, 0xcf, 0xa0, 0x8e, 0xd2, 0x92,
	0x51, 0x31, 0x16, 0x4c, 0x4b, 0x1e, 0x8f, 0x55, 0x3c, 0x63, 0x82, 0x86, 0xb9, 0x04, 0x0d, 0xb8,
	0x29, 0x98, 0xa6, 0x09, 0xd5, 0x74, 0x1c, 0xb3, 0xb9, 0x66, 0xb2, 0x13, 0xa4, 0x5c, 0xcf, 0x16,
	0x93, 0x30, 0x06, 0x41, 0x52, 0x48, 0x81, 0x18, 0x6e, 0xb2, 0x98, 0x9a, 0xcb, 0x1c, 0xe6, 0x65,
	0xfd, 0xce, 0xc5, 0x01, 0x2e, 0x56, 0x5c, 0xbf, 0xc2, 0x8a, 0xa4, 0x10, 0x98, 0x30, 0x58, 0xd2,
	0x8c, 0x27, 0x54, 0x83, 0x54, 0x64, 0xff, 0xb4, 0x9e, 0xff, 0xe1, 0x20, 0x3c, 0x32, 0xb5, 0xee,
	0x4d, 0xab, 0x91, 0x29, 0x85, 0x07, 0xc8, 0xe1, 0x89, 0x5b, 0xe9, 0x56, 0xfa, 0xf5, 0x61, 0xaf,
	0x2c, 0xbc, 0xfa, 0x14, 0xa4, 0xb8, 0xf2, 0x79, 0xe2, 0x7f, 0x7d, 0x7a, 0x6d, 0xd4, 0x7a, 0x79,
	0x3a, 0x09, 0x2e, 0x69, 0x30, 0xbd, 0x0e, 0x6e, 0x9e, 0xdf, 0x4e, 0x07, 0xef, 0xbd, 0xc8, 0xe1,
	0x09, 0xee, 0xa2, 0x2a, 0xc8, 0xd4, 0x75, 0xba, 0x95, 0x7e, 0x6d, 0xd8, 0x28, 0x0b, 0x0f, 0x59,
	0x0d, 0x64, 0xea, 0x47, 0x9b, 0x08, 0x1f, 0xa3, 0x7f, 0x4b, 0x26, 0x15, 0x87, 0xb9, 0x5b, 0x35,
	0x14, 0x2e, 0x0b, 0xaf, 0x61, 0xa9, 0x6d, 0xe0, 0x47, 0x3b, 0x04, 0x9f, 0x23, 0x64, 0xb7, 0x7a,
	0x58, 0xe7, 0xcc, 0xfd, 0x63, 0xda, 0x1c, 0x95, 0x85, 0xd7, 0xb2, 0xc2, 0x4f, 0xe6, 0x47, 0x07,
	0x20, 0x26, 0xe8, 0xff, 0x66, 0x49, 0x23, 0xd5, 0x8c, 0xd4, 0x2e, 0x0b, 0xaf, 0x69, 0xa5, 0x5d,
	0xe2, 0x47, 0x7b, 0x68, 0x78, 0xf7, 0x78, 0x9b, 0x42, 0xc8, 0xa8, 0x5a, 0x43, 0xae, 0xc2, 0x0c,
	0x62, 0x9a, 0x91, 0x18, 0xe6, 0x5a, 0xd2, 0x58, 0x2b, 0x3b, 0xbc, 0x64, 0x39, 0x04, 0x02, 0x12,
	0x96, 0x29, 0xb2, 0x05, 0x89, 0x39, 0xc9, 0xaf, 0x8f, 0x9b, 0xfc, 0x35, 0xfc, 0xd9, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x57, 0x0b, 0xbb, 0xee, 0x01, 0x00, 0x00,
}
