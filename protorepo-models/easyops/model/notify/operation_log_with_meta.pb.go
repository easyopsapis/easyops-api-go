// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operation_log_with_meta.proto

package notify

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//带有系统名称和topic的operation log
type OperationLogWithMeta struct {
	//
	//系统名称
	System string `protobuf:"bytes,1,opt,name=system,proto3" json:"system" form:"system"`
	//
	//topic
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic" form:"topic"`
	//
	//OperationLog
	Data                 *OperationLog `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OperationLogWithMeta) Reset()         { *m = OperationLogWithMeta{} }
func (m *OperationLogWithMeta) String() string { return proto.CompactTextString(m) }
func (*OperationLogWithMeta) ProtoMessage()    {}
func (*OperationLogWithMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_c385e111beb967ff, []int{0}
}
func (m *OperationLogWithMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationLogWithMeta.Unmarshal(m, b)
}
func (m *OperationLogWithMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationLogWithMeta.Marshal(b, m, deterministic)
}
func (m *OperationLogWithMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationLogWithMeta.Merge(m, src)
}
func (m *OperationLogWithMeta) XXX_Size() int {
	return xxx_messageInfo_OperationLogWithMeta.Size(m)
}
func (m *OperationLogWithMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationLogWithMeta.DiscardUnknown(m)
}

var xxx_messageInfo_OperationLogWithMeta proto.InternalMessageInfo

func (m *OperationLogWithMeta) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *OperationLogWithMeta) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *OperationLogWithMeta) GetData() *OperationLog {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationLogWithMeta)(nil), "notify.OperationLogWithMeta")
}

func init() { proto.RegisterFile("operation_log_with_meta.proto", fileDescriptor_c385e111beb967ff) }

var fileDescriptor_c385e111beb967ff = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xa9, 0x3f, 0x03, 0x76, 0x14, 0xb5, 0xcc, 0xa2, 0x0c, 0x48, 0x87, 0x2e, 0x64, 0x5c,
	0x4c, 0x02, 0xba, 0xd2, 0x95, 0x74, 0xad, 0x08, 0x45, 0x10, 0xdc, 0x94, 0x34, 0x93, 0xa6, 0x81,
	0xa6, 0x37, 0x24, 0x77, 0x90, 0xbe, 0x8f, 0xcf, 0xd5, 0x87, 0xe8, 0x13, 0x88, 0xc9, 0x08, 0xba,
	0x76, 0x97, 0x73, 0xee, 0x77, 0x72, 0xcf, 0x8d, 0xaf, 0xc0, 0x08, 0xcb, 0x50, 0x41, 0x5f, 0x75,
	0x20, 0xab, 0x0f, 0x85, 0x6d, 0xa5, 0x05, 0x32, 0x62, 0x2c, 0x20, 0x24, 0xb3, 0x1e, 0x50, 0x35,
	0xc3, 0x72, 0x23, 0x15, 0xb6, 0xbb, 0x9a, 0x70, 0xd0, 0x54, 0x82, 0x04, 0xea, 0xc7, 0xf5, 0xae,
	0xf1, 0xca, 0x0b, 0xff, 0x0a, 0xb1, 0xe5, 0xab, 0x04, 0x22, 0x98, 0x1b, 0xc0, 0x38, 0xd2, 0x01,
	0x67, 0x1d, 0xe5, 0xd0, 0xa3, 0x65, 0x1c, 0x5d, 0x48, 0x5a, 0x61, 0x60, 0xa3, 0x61, 0x2b, 0x3a,
	0x47, 0xf7, 0x20, 0xf5, 0x92, 0x86, 0x85, 0xf4, 0x4f, 0xad, 0xf0, 0x6b, 0xfe, 0x19, 0xc5, 0x8b,
	0x97, 0x1f, 0xff, 0x09, 0xe4, 0x9b, 0xc2, 0xf6, 0x59, 0x20, 0x4b, 0x6e, 0xe2, 0x99, 0x1b, 0x1c,
	0x0a, 0x9d, 0x46, 0xab, 0x68, 0x7d, 0x52, 0x5c, 0x4e, 0x63, 0x76, 0xd6, 0x80, 0xd5, 0x0f, 0x79,
	0xf0, 0xf3, 0x72, 0x0f, 0x24, 0xd7, 0xf1, 0x31, 0x82, 0x51, 0x3c, 0x3d, 0xf0, 0xe4, 0xc5, 0x34,
	0x66, 0xa7, 0x81, 0xf4, 0x76, 0x5e, 0x86, 0x71, 0x72, 0x1f, 0x1f, 0x6d, 0x19, 0xb2, 0xf4, 0x70,
	0x15, 0xad, 0xe7, 0xb7, 0x0b, 0x12, 0x6a, 0x91, 0xdf, 0xeb, 0x8b, 0xf3, 0x69, 0xcc, 0xe6, 0x21,
	0xfc, 0xcd, 0xe6, 0xa5, 0x8f, 0x14, 0xc5, 0xfb, 0xe3, 0x7f, 0xcf, 0xaf, 0x67, 0x1e, 0xbb, 0xfb,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x84, 0x86, 0x44, 0x9f, 0x01, 0x00, 0x00,
}