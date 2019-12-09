// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: diff.proto

package file_repository

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
//变更信息
type Diff struct {
	//
	//操作类型
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op" form:"op"`
	//
	//文件路径
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file" form:"file"`
	//
	//文件类型
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type" form:"type"`
	//
	//文件权限
	Perm string `protobuf:"bytes,4,opt,name=perm,proto3" json:"perm" form:"perm"`
	//
	//新的检验值
	NewMd5 string `protobuf:"bytes,5,opt,name=newMd5,proto3" json:"newMd5" form:"newMd5"`
	//
	//旧的检验值
	OldMd5 string `protobuf:"bytes,6,opt,name=oldMd5,proto3" json:"oldMd5" form:"oldMd5"`
	//
	//文件原本大小
	OldSize int64 `protobuf:"varint,7,opt,name=oldSize,proto3" json:"oldSize" form:"oldSize"`
	//
	//文件新的大小
	NewSize int64 `protobuf:"varint,8,opt,name=newSize,proto3" json:"newSize" form:"newSize"`
	//
	//参数path的编码格式
	Encoding             string   `protobuf:"bytes,9,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Diff) Reset()         { *m = Diff{} }
func (m *Diff) String() string { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()    {}
func (*Diff) Descriptor() ([]byte, []int) {
	return fileDescriptor_686521effc814b25, []int{0}
}
func (m *Diff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Diff.Unmarshal(m, b)
}
func (m *Diff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Diff.Marshal(b, m, deterministic)
}
func (m *Diff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Diff.Merge(m, src)
}
func (m *Diff) XXX_Size() int {
	return xxx_messageInfo_Diff.Size(m)
}
func (m *Diff) XXX_DiscardUnknown() {
	xxx_messageInfo_Diff.DiscardUnknown(m)
}

var xxx_messageInfo_Diff proto.InternalMessageInfo

func (m *Diff) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *Diff) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Diff) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Diff) GetPerm() string {
	if m != nil {
		return m.Perm
	}
	return ""
}

func (m *Diff) GetNewMd5() string {
	if m != nil {
		return m.NewMd5
	}
	return ""
}

func (m *Diff) GetOldMd5() string {
	if m != nil {
		return m.OldMd5
	}
	return ""
}

func (m *Diff) GetOldSize() int64 {
	if m != nil {
		return m.OldSize
	}
	return 0
}

func (m *Diff) GetNewSize() int64 {
	if m != nil {
		return m.NewSize
	}
	return 0
}

func (m *Diff) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func init() {
	proto.RegisterType((*Diff)(nil), "file_repository.Diff")
}

func init() { proto.RegisterFile("diff.proto", fileDescriptor_686521effc814b25) }

var fileDescriptor_686521effc814b25 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x95, 0xd4, 0x7f, 0xda, 0xcc, 0x4f, 0x1b, 0xe1, 0x6e, 0xac, 0x4a, 0xc8, 0x65, 0xe8,
	0xc2, 0x69, 0x99, 0x4c, 0x44, 0x54, 0x10, 0x11, 0x02, 0x1a, 0x21, 0x50, 0x85, 0xd8, 0x98, 0x5d,
	0x3d, 0x31, 0x72, 0xec, 0xb1, 0xb1, 0xb0, 0x7d, 0x2d, 0xdb, 0x25, 0x4a, 0x9b, 0xee, 0x78, 0x4e,
	0x23, 0xf1, 0x08, 0x7e, 0x02, 0x34, 0x33, 0x49, 0x8a, 0x82, 0xd8, 0xcd, 0x9c, 0xf3, 0x9d, 0x73,
	0xef, 0xe2, 0x22, 0x14, 0xc4, 0x61, 0x38, 0xc8, 0x0b, 0xa8, 0x40, 0xef, 0x85, 0x71, 0xc2, 0xbf,
	0x14, 0x3c, 0x87, 0x32, 0xae, 0xa0, 0x58, 0x1c, 0x91, 0x28, 0xae, 0xbe, 0x5e, 0xcf, 0x06, 0x3e,
	0xa4, 0x34, 0x82, 0x08, 0xa8, 0xe4, 0x66, 0xd7, 0xa1, 0xfc, 0xc9, 0x8f, 0x7c, 0xa9, 0xfc, 0xd1,
	0xf3, 0x3f, 0xf0, 0x74, 0x1e, 0x57, 0xdf, 0x60, 0x4e, 0x23, 0x20, 0xd2, 0x24, 0xdf, 0xbd, 0x24,
	0x0e, 0xbc, 0x0a, 0x8a, 0x92, 0x6e, 0x9e, 0x2a, 0x87, 0x7f, 0x68, 0x48, 0x7b, 0x17, 0x87, 0xa1,
	0xfe, 0x08, 0xb5, 0x21, 0x37, 0x5a, 0xc7, 0x2d, 0xab, 0x3b, 0xd9, 0x6f, 0x6a, 0xb3, 0x1b, 0x42,
	0x91, 0x8e, 0x31, 0xe4, 0xd8, 0x6e, 0x43, 0xae, 0xe7, 0x48, 0x13, 0x1b, 0x1a, 0x6d, 0x09, 0xb0,
	0xa6, 0x36, 0xff, 0x57, 0x80, 0x50, 0xf1, 0xaf, 0x9f, 0xe6, 0x25, 0xfa, 0xe0, 0x5a, 0x16, 0xa3,
	0x8e, 0xcb, 0xe8, 0x98, 0x9d, 0xb2, 0x37, 0x18, 0xbf, 0x7a, 0xcd, 0x96, 0xac, 0x60, 0xd9, 0xf4,
	0xac, 0x7f, 0xd6, 0x5f, 0x5a, 0x8c, 0xf6, 0x97, 0x8e, 0x47, 0x6e, 0x2e, 0xc8, 0xd5, 0x74, 0x6c,
	0x31, 0xe6, 0xb8, 0x8c, 0xfd, 0x4d, 0x9e, 0x9e, 0xd8, 0x72, 0x92, 0xfe, 0x04, 0x69, 0xd5, 0x22,
	0xe7, 0xc6, 0x8e, 0x9c, 0xd8, 0xbb, 0x9f, 0x28, 0x54, 0x6c, 0x4b, 0x53, 0x3f, 0x47, 0x5a, 0xce,
	0x8b, 0xd4, 0xd0, 0x24, 0xf4, 0xf8, 0x1e, 0x12, 0xaa, 0x58, 0xeb, 0x00, 0x3d, 0x70, 0xad, 0xa1,
	0x33, 0x24, 0x2f, 0xa6, 0xb7, 0xa3, 0xbb, 0xbe, 0x2d, 0x71, 0xfd, 0x2d, 0xea, 0x64, 0x7c, 0xfe,
	0x29, 0x38, 0x37, 0xfe, 0x93, 0x41, 0xab, 0xa9, 0xcd, 0x7d, 0x15, 0x54, 0xba, 0x88, 0x1e, 0xa2,
	0x87, 0xae, 0xe3, 0x91, 0xf0, 0x82, 0xbc, 0x1f, 0x92, 0x97, 0xd3, 0xdb, 0xd1, 0xb3, 0xbb, 0x13,
	0x7b, 0x95, 0x13, 0x0d, 0x90, 0x04, 0xa2, 0xa1, 0xb3, 0xdd, 0xa0, 0xf4, 0x7f, 0x37, 0x28, 0x5f,
	0x7f, 0x8a, 0x76, 0x21, 0x09, 0x3e, 0xc7, 0x37, 0xdc, 0xd8, 0x3d, 0x6e, 0x59, 0x3b, 0x13, 0xbd,
	0xa9, 0xcd, 0x83, 0x4d, 0x85, 0x30, 0xb0, 0xbd, 0x46, 0x04, 0x9d, 0xf1, 0xb9, 0xa4, 0xf7, 0xb6,
	0xe9, 0x95, 0x81, 0xed, 0x35, 0xa2, 0x53, 0xb4, 0xc7, 0x33, 0x1f, 0x82, 0x38, 0x8b, 0x8c, 0xae,
	0xdc, 0xef, 0xb0, 0xa9, 0xcd, 0x9e, 0xc2, 0xd7, 0x0e, 0xb6, 0x37, 0xd0, 0xe4, 0xe3, 0xd5, 0x65,
	0x04, 0x03, 0xee, 0x95, 0x0b, 0xc8, 0xcb, 0x41, 0x02, 0xbe, 0x97, 0x50, 0x1f, 0xb2, 0xaa, 0xf0,
	0xfc, 0xaa, 0x54, 0xa7, 0x27, 0x6e, 0x93, 0xa4, 0x10, 0xf0, 0xa4, 0xa4, 0x2b, 0x90, 0xca, 0x2f,
	0xdd, 0x3a, 0xdd, 0x59, 0x47, 0xf2, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x48, 0x63,
	0x01, 0xe0, 0x02, 0x00, 0x00,
}
