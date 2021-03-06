// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_metrics_schema.proto

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
//指标表
type StreamMetricsSchema struct {
	//
	//id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//org
	Org int32 `protobuf:"varint,2,opt,name=org,proto3" json:"org" form:"org"`
	//
	//指标表版本
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version" form:"version"`
	//
	//指标表名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列表
	Dimensions []*StreamMetricsSchemaField `protobuf:"bytes,5,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//指标列表
	Metrics              []*StreamMetricsSchemaField `protobuf:"bytes,6,rep,name=metrics,proto3" json:"metrics" form:"metrics"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StreamMetricsSchema) Reset()         { *m = StreamMetricsSchema{} }
func (m *StreamMetricsSchema) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsSchema) ProtoMessage()    {}
func (*StreamMetricsSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_761de2744074de43, []int{0}
}
func (m *StreamMetricsSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsSchema.Unmarshal(m, b)
}
func (m *StreamMetricsSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsSchema.Marshal(b, m, deterministic)
}
func (m *StreamMetricsSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsSchema.Merge(m, src)
}
func (m *StreamMetricsSchema) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsSchema.Size(m)
}
func (m *StreamMetricsSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsSchema.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsSchema proto.InternalMessageInfo

func (m *StreamMetricsSchema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamMetricsSchema) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *StreamMetricsSchema) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StreamMetricsSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamMetricsSchema) GetDimensions() []*StreamMetricsSchemaField {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *StreamMetricsSchema) GetMetrics() []*StreamMetricsSchemaField {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsSchema)(nil), "metadata_center.StreamMetricsSchema")
}

func init() { proto.RegisterFile("stream_metrics_schema.proto", fileDescriptor_761de2744074de43) }

var fileDescriptor_761de2744074de43 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0xd9, 0x6c, 0xff, 0xd0, 0x29, 0xb4, 0x74, 0x8a, 0x10, 0xea, 0x21, 0x61, 0xec, 0x61,
	0x05, 0x93, 0x91, 0x5a, 0x04, 0xbd, 0x99, 0x43, 0x41, 0xc4, 0x4b, 0x7a, 0x10, 0x14, 0x8d, 0xb3,
	0x99, 0x49, 0x3a, 0x34, 0x93, 0x77, 0x99, 0x99, 0x76, 0x11, 0xf1, 0xab, 0x46, 0xf0, 0x23, 0xe4,
	0xea, 0x45, 0x32, 0x93, 0xd5, 0x65, 0xd9, 0x8b, 0xd0, 0xdb, 0xbc, 0x79, 0x9e, 0xe7, 0xf7, 0xfe,
	0x21, 0xe8, 0xb1, 0xb1, 0x5a, 0x30, 0x55, 0x28, 0x61, 0xb5, 0x2c, 0x4d, 0x61, 0xca, 0x1b, 0xa1,
	0x58, 0xba, 0xd0, 0x60, 0x01, 0x1f, 0x2b, 0x61, 0x19, 0x67, 0x96, 0x15, 0xa5, 0x68, 0xad, 0xd0,
	0x67, 0x49, 0x2d, 0xed, 0xcd, 0xdd, 0x3c, 0x2d, 0x41, 0xd1, 0x1a, 0x6a, 0xa0, 0xce, 0x37, 0xbf,
	0xab, 0x5c, 0xe5, 0x0a, 0xf7, 0xf2, 0xf9, 0xb3, 0xdb, 0x1a, 0x52, 0xc1, 0xcc, 0x37, 0x58, 0x98,
	0xb4, 0x81, 0x92, 0x35, 0xb4, 0x84, 0xd6, 0x6a, 0x56, 0x5a, 0xe3, 0x93, 0x5a, 0x2c, 0x20, 0x51,
	0xc0, 0x45, 0x63, 0xe8, 0x68, 0xa4, 0xae, 0xa4, 0x1b, 0x9d, 0xe9, 0xd6, 0x31, 0x8b, 0x4a, 0x8a,
	0x86, 0x8f, 0xcd, 0x5e, 0xae, 0xcd, 0xa6, 0x96, 0xd2, 0xde, 0xc2, 0x92, 0xd6, 0x90, 0x38, 0x31,
	0xb9, 0x67, 0x8d, 0xe4, 0xcc, 0x82, 0x36, 0xf4, 0xef, 0xd3, 0xe7, 0xc8, 0xef, 0x00, 0x9d, 0x5e,
	0x3b, 0xfa, 0x7b, 0x0f, 0xbf, 0x76, 0x6c, 0x7c, 0x89, 0x02, 0xc9, 0xc3, 0x49, 0x3c, 0x99, 0x1d,
	0x64, 0xe7, 0x7d, 0x17, 0x1d, 0x54, 0xa0, 0xd5, 0x6b, 0x22, 0x39, 0xf9, 0xf5, 0x33, 0x3a, 0x45,
	0x27, 0x5f, 0x3e, 0x3d, 0x4f, 0x5e, 0xb1, 0xa4, 0x7a, 0x93, 0x5c, 0x7d, 0xfe, 0x7e, 0x71, 0xf9,
	0xe3, 0x3c, 0x0f, 0x24, 0xc7, 0x31, 0x9a, 0x82, 0xae, 0xc3, 0x20, 0x9e, 0xcc, 0x76, 0xb3, 0xa3,
	0xbe, 0x8b, 0x90, 0x8f, 0x81, 0xae, 0x49, 0x3e, 0x48, 0xf8, 0x19, 0xda, 0xbf, 0x17, 0xda, 0x48,
	0x68, 0xc3, 0xa9, 0x73, 0xe1, 0xbe, 0x8b, 0x8e, 0xbc, 0x6b, 0x14, 0x48, 0xbe, 0xb2, 0xe0, 0x27,
	0x68, 0xa7, 0x65, 0x4a, 0x84, 0x3b, 0x6e, 0x8e, 0xe3, 0xbe, 0x8b, 0x0e, 0xbd, 0x75, 0xf8, 0x4a,
	0x72, 0x27, 0xe2, 0xaf, 0x08, 0x71, 0xa9, 0x44, 0x3b, 0x24, 0x4c, 0xb8, 0x1b, 0x4f, 0x67, 0x87,
	0x17, 0x4f, 0xd3, 0x8d, 0x13, 0xa6, 0x5b, 0x96, 0xbc, 0x1a, 0xee, 0x97, 0x3d, 0xea, 0xbb, 0xe8,
	0xc4, 0x53, 0xff, 0x61, 0x48, 0xbe, 0xc6, 0xc4, 0x1f, 0xd0, 0xfe, 0x78, 0xfa, 0x70, 0xef, 0x7f,
	0xf1, 0x6b, 0xfb, 0x8d, 0x0c, 0x92, 0xaf, 0x68, 0xd9, 0xbb, 0x8f, 0x6f, 0x1f, 0xec, 0x27, 0x99,
	0xef, 0x39, 0xff, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xd0, 0xc4, 0xdc, 0xd5, 0x02,
	0x00, 0x00,
}
