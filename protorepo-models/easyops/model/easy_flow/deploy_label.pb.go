// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deploy_label.proto

package easy_flow

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
//标签
type DeployLabel struct {
	//
	//来源
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from" form:"from"`
	//
	//策略名称
	StrategyName         string   `protobuf:"bytes,2,opt,name=strategyName,proto3" json:"strategyName" form:"strategyName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployLabel) Reset()         { *m = DeployLabel{} }
func (m *DeployLabel) String() string { return proto.CompactTextString(m) }
func (*DeployLabel) ProtoMessage()    {}
func (*DeployLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_df3befb5e7ede0bf, []int{0}
}
func (m *DeployLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployLabel.Unmarshal(m, b)
}
func (m *DeployLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployLabel.Marshal(b, m, deterministic)
}
func (m *DeployLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployLabel.Merge(m, src)
}
func (m *DeployLabel) XXX_Size() int {
	return xxx_messageInfo_DeployLabel.Size(m)
}
func (m *DeployLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployLabel.DiscardUnknown(m)
}

var xxx_messageInfo_DeployLabel proto.InternalMessageInfo

func (m *DeployLabel) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *DeployLabel) GetStrategyName() string {
	if m != nil {
		return m.StrategyName
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployLabel)(nil), "easy_flow.DeployLabel")
}

func init() { proto.RegisterFile("deploy_label.proto", fileDescriptor_df3befb5e7ede0bf) }

var fileDescriptor_df3befb5e7ede0bf = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4b, 0x43, 0x31,
	0x14, 0x85, 0xa9, 0x88, 0xd0, 0x54, 0x10, 0xe2, 0x60, 0x71, 0xa9, 0xc4, 0xc5, 0xa5, 0xc9, 0xe0,
	0xa6, 0x5b, 0xd5, 0x4d, 0x1c, 0x3a, 0xba, 0x94, 0x24, 0x4d, 0xa2, 0x90, 0x78, 0x42, 0x92, 0x22,
	0xef, 0xcf, 0xbe, 0x1f, 0xf1, 0x7e, 0x81, 0xbc, 0xfb, 0x44, 0xec, 0x76, 0xcf, 0x3d, 0xdf, 0x3d,
	0x9c, 0xcb, 0xf8, 0xde, 0xe5, 0x88, 0x6e, 0x17, 0xb5, 0x71, 0x51, 0xe6, 0x82, 0x06, 0x3e, 0x77,
	0xba, 0x76, 0x3b, 0x1f, 0xf1, 0x7d, 0xbd, 0x0e, 0x9f, 0xed, 0xe3, 0x60, 0xa4, 0x45, 0x52, 0x01,
	0x01, 0x8a, 0x08, 0x73, 0xf0, 0xa4, 0x48, 0xd0, 0x34, 0x5d, 0x0a, 0xb0, 0xc5, 0x33, 0xe5, 0xbd,
	0x8e, 0x71, 0xfc, 0x96, 0x9d, 0xfa, 0x82, 0xb4, 0x9c, 0xdd, 0xcc, 0xee, 0xe6, 0x9b, 0x8b, 0xa1,
	0x5f, 0x2d, 0x3c, 0x4a, 0x7a, 0x10, 0xe3, 0x56, 0x6c, 0xc9, 0xe4, 0x8f, 0xec, 0xbc, 0xb6, 0xa2,
	0x9b, 0x0b, 0xdd, 0x9b, 0x4e, 0x6e, 0x79, 0x42, 0xf0, 0xd5, 0xd0, 0xaf, 0x2e, 0x27, 0xf8, 0xbf,
	0x2b, 0xb6, 0x47, 0xf0, 0xe6, 0xe5, 0xfd, 0x29, 0x40, 0x8e, 0x7d, 0x91, 0xab, 0x8c, 0xb0, 0x3a,
	0x2a, 0x8b, 0xaf, 0x56, 0xb4, 0x6d, 0x75, 0x2a, 0x5b, 0x5c, 0xc6, 0x3a, 0x61, 0xef, 0x62, 0x55,
	0xbf, 0xa0, 0x22, 0xa9, 0xfe, 0xde, 0x34, 0x67, 0x44, 0xde, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x85, 0x6b, 0x5c, 0x0e, 0x01, 0x00, 0x00,
}
