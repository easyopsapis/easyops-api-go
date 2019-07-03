// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: action_param_file.proto

package easy_command

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//动作为 "saveFile" 的action的参数
type FileActionParam struct {
	//
	//文件内容 base64 编码字符串
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content" form:"content"`
	//
	//保存到目标机器的绝对路径，如 /tmp/foo/bar.tar.gz
	DstFile              string   `protobuf:"bytes,2,opt,name=dst_file,json=dstFile,proto3" json:"dst_file" form:"dst_file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileActionParam) Reset()         { *m = FileActionParam{} }
func (m *FileActionParam) String() string { return proto.CompactTextString(m) }
func (*FileActionParam) ProtoMessage()    {}
func (*FileActionParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dde12c8920c0207, []int{0}
}
func (m *FileActionParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileActionParam.Unmarshal(m, b)
}
func (m *FileActionParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileActionParam.Marshal(b, m, deterministic)
}
func (m *FileActionParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileActionParam.Merge(m, src)
}
func (m *FileActionParam) XXX_Size() int {
	return xxx_messageInfo_FileActionParam.Size(m)
}
func (m *FileActionParam) XXX_DiscardUnknown() {
	xxx_messageInfo_FileActionParam.DiscardUnknown(m)
}

var xxx_messageInfo_FileActionParam proto.InternalMessageInfo

func (m *FileActionParam) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *FileActionParam) GetDstFile() string {
	if m != nil {
		return m.DstFile
	}
	return ""
}

func init() {
	proto.RegisterType((*FileActionParam)(nil), "easy_command.FileActionParam")
}

func init() { proto.RegisterFile("action_param_file.proto", fileDescriptor_6dde12c8920c0207) }

var fileDescriptor_6dde12c8920c0207 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x18, 0x84, 0x83, 0x07, 0xd1, 0xc6, 0x48, 0xd2, 0x8b, 0x84, 0x0b, 0xa6, 0xf1, 0x50, 0xc0, 0x76,
	0x0f, 0x26, 0x1e, 0x88, 0xd1, 0xc0, 0x01, 0xf5, 0x66, 0x38, 0xc2, 0x0f, 0xb8, 0xb4, 0x4b, 0xdd,
	0xb8, 0xdb, 0xbf, 0xd9, 0x5d, 0x24, 0x6a, 0x1f, 0xcd, 0x67, 0xc1, 0xc4, 0x47, 0xe0, 0x09, 0x4c,
	0xb7, 0x60, 0x48, 0xbc, 0xfd, 0x93, 0xf9, 0x26, 0x99, 0x7f, 0x9c, 0x33, 0x1a, 0x19, 0x8e, 0xe9,
	0x2c, 0xa3, 0x8a, 0xca, 0xd9, 0x82, 0x0b, 0x16, 0x66, 0x0a, 0x0d, 0xba, 0x27, 0x8c, 0xea, 0xf7,
	0x59, 0x84, 0x52, 0xd2, 0x34, 0x6e, 0x04, 0x09, 0x37, 0x2f, 0xcb, 0x79, 0x18, 0xa1, 0x24, 0x09,
	0x26, 0x48, 0x2c, 0x34, 0x5f, 0x2e, 0xac, 0xb2, 0xc2, 0x5e, 0x65, 0xb8, 0x71, 0xbd, 0x87, 0xcb,
	0x15, 0x37, 0xaf, 0xb8, 0x22, 0x09, 0x06, 0xd6, 0x0c, 0xde, 0xa8, 0xe0, 0x31, 0x35, 0xa8, 0x34,
	0xf9, 0x3b, 0xcb, 0x9c, 0xf7, 0x55, 0x71, 0x6a, 0x03, 0x2e, 0x58, 0xcf, 0x96, 0x7a, 0x2a, 0x3a,
	0xb9, 0x97, 0x4e, 0x35, 0xc2, 0xd4, 0xb0, 0xd4, 0xd4, 0x2b, 0xe7, 0x15, 0xff, 0xb8, 0xef, 0x6e,
	0xd6, 0xcd, 0xd3, 0x05, 0x2a, 0xd9, 0xf5, 0xb6, 0x86, 0x37, 0xdc, 0x21, 0xee, 0xa7, 0x73, 0x14,
	0x6b, 0x63, 0x1f, 0xa9, 0x1f, 0x58, 0xfc, 0x79, 0xb3, 0x6e, 0xd6, 0x4a, 0x7c, 0xe7, 0x78, 0x3f,
	0xdf, 0xcd, 0x47, 0xe7, 0x7e, 0xea, 0xfb, 0x40, 0xc6, 0x53, 0x20, 0x5d, 0x68, 0xc3, 0x9d, 0xe7,
	0xdd, 0xdc, 0x42, 0x0e, 0x0a, 0xd2, 0x49, 0xa7, 0xd5, 0x69, 0xe5, 0x3e, 0x90, 0x56, 0x3e, 0xa6,
	0xc1, 0x47, 0x2f, 0x18, 0x4d, 0xba, 0x3e, 0xc0, 0x78, 0x0a, 0xf0, 0x9f, 0x6c, 0x5f, 0x0c, 0xab,
	0xb1, 0x36, 0x45, 0xeb, 0xfe, 0xc3, 0x68, 0x90, 0x60, 0x58, 0x0c, 0x87, 0x99, 0x0e, 0x05, 0x46,
	0x54, 0x90, 0xa2, 0x98, 0xa2, 0x91, 0xd1, 0xe5, 0x64, 0x8a, 0x65, 0x18, 0x48, 0x8c, 0x99, 0xd0,
	0x64, 0x0b, 0x12, 0x2b, 0xc9, 0xfe, 0xde, 0xf3, 0x43, 0x0b, 0x5f, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xa2, 0x61, 0xc9, 0x8e, 0x9f, 0x01, 0x00, 0x00,
}
