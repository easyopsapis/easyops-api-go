// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_translate_states.proto

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
//翻译状态数据
type StreamTranslateStates struct {
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
	Payload              []*StreamTranslatePackage `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload" form:"payload"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StreamTranslateStates) Reset()         { *m = StreamTranslateStates{} }
func (m *StreamTranslateStates) String() string { return proto.CompactTextString(m) }
func (*StreamTranslateStates) ProtoMessage()    {}
func (*StreamTranslateStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6ade5770edfc51, []int{0}
}
func (m *StreamTranslateStates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTranslateStates.Unmarshal(m, b)
}
func (m *StreamTranslateStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTranslateStates.Marshal(b, m, deterministic)
}
func (m *StreamTranslateStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTranslateStates.Merge(m, src)
}
func (m *StreamTranslateStates) XXX_Size() int {
	return xxx_messageInfo_StreamTranslateStates.Size(m)
}
func (m *StreamTranslateStates) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTranslateStates.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTranslateStates proto.InternalMessageInfo

func (m *StreamTranslateStates) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *StreamTranslateStates) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *StreamTranslateStates) GetPayload() []*StreamTranslatePackage {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTranslateStates)(nil), "metadata_center.StreamTranslateStates")
}

func init() { proto.RegisterFile("stream_translate_states.proto", fileDescriptor_ca6ade5770edfc51) }

var fileDescriptor_ca6ade5770edfc51 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x41, 0x8b, 0x2b, 0x54, 0x08, 0x08, 0xa1, 0x20, 0x0d, 0xb9, 0x98, 0x83, 0xdd,
	0x05, 0xbd, 0x79, 0xcc, 0x4d, 0xbc, 0x48, 0xaa, 0x17, 0x2f, 0x61, 0xb2, 0xd9, 0x6e, 0xc5, 0x6c,
	0x26, 0xec, 0x4e, 0x0f, 0x7d, 0x39, 0x1f, 0xa5, 0x0f, 0xd1, 0x27, 0x10, 0x37, 0x09, 0x48, 0xbc,
	0xf6, 0xb6, 0x3f, 0xf3, 0xcd, 0xff, 0xff, 0x3b, 0xec, 0xd6, 0x91, 0x55, 0x60, 0x4a, 0xb2, 0xd0,
	0xba, 0x06, 0x48, 0x95, 0x8e, 0x80, 0x94, 0xe3, 0x9d, 0x45, 0xc2, 0xe8, 0xda, 0x28, 0x82, 0x1a,
	0x08, 0x4a, 0xa9, 0x5a, 0x52, 0x76, 0xb1, 0xd2, 0x9f, 0xb4, 0xdd, 0x55, 0x5c, 0xa2, 0x11, 0x1a,
	0x35, 0x0a, 0xcf, 0x55, 0xbb, 0x8d, 0x57, 0x5e, 0xf8, 0x57, 0xbf, 0xbf, 0xd8, 0x6a, 0xe4, 0x0a,
	0xdc, 0x1e, 0x3b, 0xc7, 0x1b, 0x94, 0xd0, 0x08, 0x89, 0x2d, 0x59, 0x90, 0xe4, 0xfa, 0x4d, 0xab,
	0x3a, 0x5c, 0x19, 0xac, 0x55, 0xe3, 0xc4, 0x00, 0x0a, 0x2f, 0xc5, 0x24, 0x59, 0xfc, 0x2b, 0xda,
	0x81, 0xfc, 0x02, 0xad, 0xfa, 0xa4, 0xf4, 0x3b, 0x60, 0x37, 0x6b, 0x8f, 0xbc, 0x8d, 0xc4, 0xda,
	0xff, 0x24, 0x4a, 0x58, 0x88, 0x56, 0xc7, 0x41, 0x12, 0x64, 0xe7, 0xf9, 0xfc, 0x78, 0x58, 0xb2,
	0x0d, 0x5a, 0xf3, 0x94, 0xa2, 0xd5, 0x69, 0xf1, 0x3b, 0x8a, 0xee, 0xd9, 0x4c, 0xa2, 0x31, 0xd0,
	0xd6, 0xf1, 0x59, 0x12, 0x64, 0x97, 0x79, 0x74, 0x3c, 0x2c, 0xe7, 0x3d, 0x35, 0x0c, 0xd2, 0x62,
	0x44, 0xa2, 0x77, 0x36, 0xeb, 0x60, 0xdf, 0x20, 0xd4, 0x71, 0x98, 0x84, 0xd9, 0xd5, 0xc3, 0x1d,
	0x9f, 0x74, 0xe5, 0x93, 0x22, 0xaf, 0x7d, 0xd3, 0xbf, 0xb6, 0x83, 0x43, 0x5a, 0x8c, 0x5e, 0xf9,
	0xcb, 0xc7, 0xf3, 0xc9, 0x8e, 0x55, 0x5d, 0x78, 0xfe, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x08, 0x20, 0xf2, 0xdf, 0x01, 0x00, 0x00,
}