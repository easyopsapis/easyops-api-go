// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deploy_info.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//部署信息
type DeployInfo struct {
	//
	//应用部署方式
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from" form:"from"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployInfo) Reset()         { *m = DeployInfo{} }
func (m *DeployInfo) String() string { return proto.CompactTextString(m) }
func (*DeployInfo) ProtoMessage()    {}
func (*DeployInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a04ce1d7852630b7, []int{0}
}
func (m *DeployInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployInfo.Unmarshal(m, b)
}
func (m *DeployInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployInfo.Marshal(b, m, deterministic)
}
func (m *DeployInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployInfo.Merge(m, src)
}
func (m *DeployInfo) XXX_Size() int {
	return xxx_messageInfo_DeployInfo.Size(m)
}
func (m *DeployInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeployInfo proto.InternalMessageInfo

func (m *DeployInfo) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployInfo)(nil), "notify.DeployInfo")
}

func init() { proto.RegisterFile("deploy_info.proto", fileDescriptor_a04ce1d7852630b7) }

var fileDescriptor_a04ce1d7852630b7 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0x2d, 0xc8,
	0xc9, 0xaf, 0x8c, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb,
	0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e,
	0x98, 0x05, 0xd1, 0xa6, 0x64, 0xc8, 0xc5, 0xe5, 0x02, 0x36, 0xcb, 0x33, 0x2f, 0x2d, 0x5f, 0x48,
	0x99, 0x8b, 0x25, 0xad, 0x28, 0x3f, 0x57, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xff, 0xd3,
	0x3d, 0x79, 0xee, 0xb4, 0xfc, 0xa2, 0x5c, 0x2b, 0x25, 0x90, 0xa8, 0x52, 0x10, 0x58, 0xd2, 0xc9,
	0x29, 0xca, 0x21, 0x3d, 0x5f, 0x2f, 0x35, 0xb1, 0xb8, 0x32, 0xbf, 0xa0, 0x58, 0x2f, 0x27, 0x3f,
	0x39, 0x31, 0x47, 0x3f, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31, 0xb9, 0xa4, 0x18, 0x62, 0x5d, 0x51,
	0x6a, 0x41, 0xbe, 0x6e, 0x6e, 0x7e, 0x4a, 0x6a, 0x4e, 0xb1, 0x3e, 0x54, 0xa1, 0x3e, 0x98, 0xab,
	0x0f, 0x71, 0x65, 0x12, 0x1b, 0x58, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x93, 0xfd, 0x3c,
	0xd0, 0xc9, 0x00, 0x00, 0x00,
}
