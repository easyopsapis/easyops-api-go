// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: link.proto

package topology

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
//边（线）
type Link struct {
	//
	//源点
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source" form:"source"`
	//
	//目标点
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target" form:"target"`
	//
	//线样式
	Style                *LinkStyle `protobuf:"bytes,3,opt,name=style,proto3" json:"style" form:"style"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ee656911eb8a56a, []int{0}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Link) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Link) GetStyle() *LinkStyle {
	if m != nil {
		return m.Style
	}
	return nil
}

func init() {
	proto.RegisterType((*Link)(nil), "topology.Link")
}

func init() { proto.RegisterFile("link.proto", fileDescriptor_2ee656911eb8a56a) }

var fileDescriptor_2ee656911eb8a56a = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0x15, 0xfe, 0x54, 0x60, 0x40, 0x82, 0xb0, 0x44, 0x5d, 0x5a, 0x79, 0x2a, 0x43, 0x6d,
	0x09, 0x36, 0xd8, 0x22, 0xc6, 0x4e, 0x61, 0x63, 0x4b, 0x8c, 0x6b, 0xa2, 0x3a, 0xb9, 0xc8, 0xfe,
	0x32, 0xe4, 0x51, 0x78, 0xb9, 0x3e, 0x44, 0x9f, 0x00, 0xc5, 0x4e, 0x16, 0x56, 0x36, 0xdf, 0xe9,
	0x77, 0xbe, 0xfb, 0x18, 0xb3, 0x75, 0x7b, 0x10, 0x9d, 0x03, 0x21, 0xbd, 0x22, 0x74, 0xb0, 0x30,
	0xc3, 0x72, 0x6b, 0x6a, 0xfa, 0xee, 0x2b, 0xa1, 0xd0, 0x48, 0x03, 0x03, 0x19, 0x80, 0xaa, 0xdf,
	0x07, 0x15, 0x44, 0x78, 0xc5, 0xe0, 0xb2, 0x30, 0x10, 0xba, 0xf4, 0x03, 0x3a, 0x2f, 0x2c, 0x54,
	0x69, 0xa5, 0x42, 0x4b, 0xae, 0x54, 0xe4, 0x63, 0xd2, 0xe9, 0x0e, 0xdb, 0x06, 0x5f, 0xda, 0x7a,
	0x39, 0x81, 0x32, 0x48, 0x39, 0x57, 0xca, 0x71, 0xc7, 0x07, 0x0d, 0x56, 0xc7, 0x3f, 0xf9, 0x4f,
	0xc2, 0x2e, 0x76, 0x75, 0x7b, 0x48, 0x9f, 0xd8, 0xc2, 0xa3, 0x77, 0x4a, 0x67, 0xc9, 0x3a, 0xd9,
	0x5c, 0xe7, 0x0f, 0xa7, 0xe3, 0xea, 0x6e, 0x0f, 0xd7, 0xbc, 0xf2, 0xe8, 0xf3, 0x62, 0x02, 0x46,
	0x94, 0x4a, 0x67, 0x34, 0x65, 0x67, 0x7f, 0xd1, 0xe8, 0xf3, 0x62, 0x02, 0xd2, 0x37, 0x76, 0xe9,
	0xc7, 0xb6, 0xec, 0x7c, 0x9d, 0x6c, 0x6e, 0x9e, 0x1f, 0xc5, 0x3c, 0x44, 0xec, 0xe6, 0x21, 0xf9,
	0xfd, 0xe9, 0xb8, 0xba, 0x9d, 0x9a, 0x46, 0x83, 0x17, 0x31, 0x93, 0xbf, 0x7f, 0xe6, 0xff, 0xbf,
	0xb8, 0x5a, 0x04, 0xf0, 0xe5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x85, 0x69, 0x91, 0x3a, 0x83, 0x01,
	0x00, 0x00,
}
