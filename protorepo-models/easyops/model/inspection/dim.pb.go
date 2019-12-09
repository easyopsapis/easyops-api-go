// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dim.proto

package inspection

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
//维度信息
type InspectionDim struct {
	//
	//维度id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//维度名称
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionDim) Reset()         { *m = InspectionDim{} }
func (m *InspectionDim) String() string { return proto.CompactTextString(m) }
func (*InspectionDim) ProtoMessage()    {}
func (*InspectionDim) Descriptor() ([]byte, []int) {
	return fileDescriptor_10ae577f632ffe6e, []int{0}
}
func (m *InspectionDim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionDim.Unmarshal(m, b)
}
func (m *InspectionDim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionDim.Marshal(b, m, deterministic)
}
func (m *InspectionDim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionDim.Merge(m, src)
}
func (m *InspectionDim) XXX_Size() int {
	return xxx_messageInfo_InspectionDim.Size(m)
}
func (m *InspectionDim) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionDim.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionDim proto.InternalMessageInfo

func (m *InspectionDim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionDim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*InspectionDim)(nil), "inspection.InspectionDim")
}

func init() { proto.RegisterFile("dim.proto", fileDescriptor_10ae577f632ffe6e) }

var fileDescriptor_10ae577f632ffe6e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x86, 0xf9, 0x8a, 0x08, 0x8d, 0x7c, 0x08, 0x9d, 0x3e, 0x04, 0xa9, 0xc4, 0xc5, 0xa5, 0xc9,
	0xe0, 0xe6, 0x58, 0x8a, 0xe0, 0x5a, 0x37, 0xb7, 0x34, 0x49, 0xe3, 0x41, 0x93, 0x0b, 0x49, 0x3a,
	0xf8, 0x67, 0xfb, 0x23, 0xfa, 0x0b, 0xc4, 0xab, 0xe8, 0x76, 0xcf, 0xdd, 0xf3, 0xc2, 0xbd, 0xac,
	0x36, 0xe0, 0x45, 0x4c, 0x58, 0xb0, 0x61, 0x10, 0x72, 0xb4, 0xba, 0x00, 0x86, 0xbb, 0xce, 0x41,
	0xf9, 0x5c, 0x27, 0xa1, 0xd1, 0x4b, 0x87, 0x0e, 0x25, 0x29, 0xd3, 0x3a, 0x13, 0x11, 0xd0, 0x74,
	0x44, 0xf9, 0x3b, 0x3b, 0xbf, 0xfd, 0x85, 0x07, 0xf0, 0xcd, 0x3d, 0xab, 0xc0, 0x5c, 0x4e, 0x0f,
	0xa7, 0xa7, 0xba, 0x3f, 0xef, 0x5b, 0x5b, 0xcf, 0x98, 0xfc, 0x0b, 0x07, 0xc3, 0xc7, 0x0a, 0x4c,
	0xf3, 0xc8, 0xae, 0x82, 0xf2, 0xf6, 0x52, 0x91, 0x70, 0xbb, 0x6f, 0xed, 0xcd, 0x21, 0xfc, 0x6c,
	0xf9, 0x48, 0xc7, 0xfe, 0xf5, 0x63, 0x70, 0x28, 0xac, 0xca, 0x5f, 0x18, 0xb3, 0x58, 0x50, 0xab,
	0x45, 0x6a, 0x0c, 0x25, 0x29, 0x5d, 0xf2, 0xf1, 0x50, 0xb2, 0x11, 0x3b, 0x8f, 0xc6, 0x2e, 0x59,
	0xfe, 0x8a, 0x92, 0x50, 0xfe, 0x77, 0x99, 0xae, 0x49, 0x7d, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x51, 0x48, 0xcd, 0xa3, 0xeb, 0x00, 0x00, 0x00,
}