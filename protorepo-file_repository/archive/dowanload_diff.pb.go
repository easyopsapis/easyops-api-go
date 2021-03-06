// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dowanload_diff.proto

package archive

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
//DownloadDiff请求
type DownloadDiffRequest struct {
	//
	//版本差异列表
	Diff string `protobuf:"bytes,1,opt,name=diff,proto3" json:"diff" form:"diff"`
	//
	//包Id
	PackageId            string   `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadDiffRequest) Reset()         { *m = DownloadDiffRequest{} }
func (m *DownloadDiffRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadDiffRequest) ProtoMessage()    {}
func (*DownloadDiffRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14dad918b73859e5, []int{0}
}
func (m *DownloadDiffRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadDiffRequest.Unmarshal(m, b)
}
func (m *DownloadDiffRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadDiffRequest.Marshal(b, m, deterministic)
}
func (m *DownloadDiffRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadDiffRequest.Merge(m, src)
}
func (m *DownloadDiffRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadDiffRequest.Size(m)
}
func (m *DownloadDiffRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadDiffRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadDiffRequest proto.InternalMessageInfo

func (m *DownloadDiffRequest) GetDiff() string {
	if m != nil {
		return m.Diff
	}
	return ""
}

func (m *DownloadDiffRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func init() {
	proto.RegisterType((*DownloadDiffRequest)(nil), "archive.DownloadDiffRequest")
}

func init() { proto.RegisterFile("dowanload_diff.proto", fileDescriptor_14dad918b73859e5) }

var fileDescriptor_14dad918b73859e5 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xc9, 0x2f, 0x4f,
	0xcc, 0xcb, 0xc9, 0x4f, 0x4c, 0x89, 0x4f, 0xc9, 0x4c, 0x4b, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0x2c, 0x4a, 0xce, 0xc8, 0x2c, 0x4b, 0x95, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x27, 0x95, 0xa6,
	0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x27, 0x65, 0x86, 0xa4, 0x3c, 0xb7, 0x3c, 0xb3, 0x24,
	0x3b, 0xbf, 0x5c, 0x3f, 0x3d, 0x5f, 0x17, 0x2c, 0xa9, 0x5b, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58,
	0x92, 0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0xf5, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x22,
	0x4c, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0x1d, 0x63, 0xe4, 0x12, 0x76, 0xc9, 0x2f,
	0x07, 0x3b, 0xd2, 0x25, 0x33, 0x2d, 0x2d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x99,
	0x8b, 0x05, 0xe4, 0x64, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xfe, 0x4f, 0xf7, 0xe4, 0xb9,
	0xd3, 0xf2, 0x8b, 0x72, 0xad, 0x94, 0x40, 0xa2, 0x4a, 0x41, 0x60, 0x49, 0xa1, 0x4e, 0x46, 0x2e,
	0xce, 0x82, 0xc4, 0xe4, 0xec, 0xc4, 0xf4, 0x54, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0xd2, 0xec, 0x4f,
	0xf7, 0xe4, 0x05, 0x20, 0x4a, 0xe1, 0x52, 0x4a, 0x8f, 0xee, 0xcb, 0x07, 0x73, 0x05, 0xc6, 0x45,
	0x27, 0xea, 0xa6, 0x39, 0xea, 0xba, 0x19, 0xe8, 0x5a, 0xc6, 0x56, 0x5b, 0xd4, 0xea, 0xda, 0x23,
	0xf3, 0x4d, 0x48, 0xe4, 0x1b, 0x1a, 0xd5, 0xaa, 0x04, 0x21, 0x6c, 0x4f, 0x62, 0x03, 0xfb, 0xc7,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x92, 0x6f, 0x8e, 0xe0, 0x74, 0x01, 0x00, 0x00,
}
