// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_metrics_states.proto

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
//指标表状态数据
type StreamMetricsStates struct {
	//
	//org
	Org int32 `protobuf:"varint,1,opt,name=org,proto3" json:"org" form:"org"`
	//
	//状态数据变更命令
	//upsert：插入/更新指定的状态数据
	//delete：删除指定的状态数据
	//retain：仅保留指定的状态数据
	//
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command" form:"command"`
	//
	//状态数据列表
	Payload              []*StreamMetricsSchema `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload" form:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamMetricsStates) Reset()         { *m = StreamMetricsStates{} }
func (m *StreamMetricsStates) String() string { return proto.CompactTextString(m) }
func (*StreamMetricsStates) ProtoMessage()    {}
func (*StreamMetricsStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a3cd221fdf3550, []int{0}
}
func (m *StreamMetricsStates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMetricsStates.Unmarshal(m, b)
}
func (m *StreamMetricsStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMetricsStates.Marshal(b, m, deterministic)
}
func (m *StreamMetricsStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetricsStates.Merge(m, src)
}
func (m *StreamMetricsStates) XXX_Size() int {
	return xxx_messageInfo_StreamMetricsStates.Size(m)
}
func (m *StreamMetricsStates) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetricsStates.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetricsStates proto.InternalMessageInfo

func (m *StreamMetricsStates) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *StreamMetricsStates) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *StreamMetricsStates) GetPayload() []*StreamMetricsSchema {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamMetricsStates)(nil), "metadata_center.StreamMetricsStates")
}

func init() { proto.RegisterFile("stream_metrics_states.proto", fileDescriptor_d7a3cd221fdf3550) }

var fileDescriptor_d7a3cd221fdf3550 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x6a, 0xfb, 0x30,
	0x10, 0xc6, 0xf1, 0x3f, 0xfc, 0x1b, 0xaa, 0x42, 0x0a, 0xee, 0x62, 0xd2, 0x21, 0xc6, 0x74, 0xf0,
	0xd0, 0x48, 0xd0, 0x6e, 0x1d, 0xbd, 0x95, 0xd2, 0xc5, 0xd9, 0xba, 0x98, 0xb3, 0x7c, 0x51, 0x02,
	0x96, 0xcf, 0x48, 0x97, 0x21, 0x8f, 0xd6, 0x97, 0xc9, 0x43, 0xe4, 0x09, 0x4a, 0x65, 0x07, 0x5a,
	0xcf, 0xdd, 0xf4, 0xf1, 0xfd, 0xbe, 0xbb, 0x4f, 0x27, 0xee, 0x3d, 0x3b, 0x04, 0x5b, 0x59, 0x64,
	0xb7, 0xd7, 0xbe, 0xf2, 0x0c, 0x8c, 0x5e, 0xf6, 0x8e, 0x98, 0xe2, 0x5b, 0x8b, 0x0c, 0x0d, 0x30,
	0x54, 0x1a, 0x3b, 0x46, 0xb7, 0x5c, 0x9b, 0x3d, 0xef, 0x0e, 0xb5, 0xd4, 0x64, 0x95, 0x21, 0x43,
	0x2a, 0x70, 0xf5, 0x61, 0x1b, 0x54, 0x10, 0xe1, 0x35, 0xe4, 0x97, 0x68, 0x48, 0x22, 0xf8, 0x23,
	0xf5, 0x5e, 0xb6, 0xa4, 0xa1, 0x55, 0x9a, 0x3a, 0x76, 0xa0, 0xd9, 0x0f, 0x49, 0x87, 0x3d, 0xad,
	0x2d, 0x35, 0xd8, 0x7a, 0x35, 0x82, 0x2a, 0x48, 0x35, 0xd9, 0xac, 0xa6, 0x35, 0xf5, 0x0e, 0x2d,
	0x0c, 0x6b, 0xb2, 0xcf, 0x48, 0xdc, 0x6d, 0x82, 0xff, 0x3e, 0xd8, 0x9b, 0xf0, 0x89, 0x38, 0x15,
	0x33, 0x72, 0x26, 0x89, 0xd2, 0x28, 0xff, 0x5f, 0x2c, 0xce, 0xa7, 0x95, 0xd8, 0x92, 0xb3, 0x2f,
	0x19, 0x39, 0x93, 0x95, 0xdf, 0x56, 0xfc, 0x28, 0xe6, 0x9a, 0xac, 0x85, 0xae, 0x49, 0xfe, 0xa5,
	0x51, 0x7e, 0x5d, 0xc4, 0xe7, 0xd3, 0x6a, 0x31, 0x50, 0xa3, 0x91, 0x95, 0x17, 0x24, 0x2e, 0xc5,
	0xbc, 0x87, 0x63, 0x4b, 0xd0, 0x24, 0xb3, 0x74, 0x96, 0xdf, 0x3c, 0x3d, 0xc8, 0x49, 0x4d, 0xf9,
	0xbb, 0x46, 0x28, 0xf9, 0x73, 0xe6, 0x18, 0xcf, 0xca, 0xcb, 0xa0, 0xe2, 0xed, 0xe3, 0xf5, 0xcf,
	0x8e, 0x54, 0x5f, 0x05, 0xfe, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xcb, 0xe2, 0x23, 0xd5,
	0x01, 0x00, 0x00,
}
