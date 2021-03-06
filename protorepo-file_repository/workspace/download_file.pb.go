// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: download_file.proto

package workspace

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
//DownloadFile请求
type DownloadFileRequest struct {
	//
	//要操作的文件、目录路径,上传时path是指存储的目录
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//文件操作类型包括：下载 -> download
	Op string `protobuf:"bytes,3,opt,name=op,proto3" json:"op" form:"op"`
	//
	//包Id
	PackageId            string   `protobuf:"bytes,4,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileRequest) Reset()         { *m = DownloadFileRequest{} }
func (m *DownloadFileRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadFileRequest) ProtoMessage()    {}
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8a1d44449f91bf, []int{0}
}
func (m *DownloadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileRequest.Unmarshal(m, b)
}
func (m *DownloadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileRequest.Marshal(b, m, deterministic)
}
func (m *DownloadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileRequest.Merge(m, src)
}
func (m *DownloadFileRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadFileRequest.Size(m)
}
func (m *DownloadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileRequest proto.InternalMessageInfo

func (m *DownloadFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DownloadFileRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *DownloadFileRequest) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *DownloadFileRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func init() {
	proto.RegisterType((*DownloadFileRequest)(nil), "workspace.DownloadFileRequest")
}

func init() { proto.RegisterFile("download_file.proto", fileDescriptor_1c8a1d44449f91bf) }

var fileDescriptor_1c8a1d44449f91bf = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcd, 0xca, 0xd3, 0x40,
	0x18, 0x85, 0x69, 0xfc, 0x10, 0x33, 0x22, 0x4a, 0xba, 0x09, 0x9f, 0x48, 0x24, 0x74, 0x91, 0xb6,
	0x4c, 0xc7, 0x3f, 0x44, 0x8b, 0x58, 0x2a, 0x52, 0xe9, 0xd2, 0xb8, 0x33, 0xd3, 0xc8, 0x34, 0x99,
	0x4c, 0x43, 0x7e, 0xde, 0x31, 0x4d, 0x0d, 0x6a, 0x7b, 0x01, 0xde, 0x64, 0x04, 0xbd, 0x83, 0x5c,
	0x81, 0x74, 0xd2, 0x3f, 0x70, 0xf5, 0xed, 0xde, 0x33, 0xe7, 0x79, 0xe7, 0xcc, 0x19, 0xd4, 0x0d,
	0xa1, 0xca, 0x53, 0x60, 0xe1, 0x97, 0x28, 0x4e, 0xf9, 0x48, 0x16, 0x50, 0x82, 0xa1, 0x57, 0x50,
	0x24, 0x6b, 0xc9, 0x02, 0x7e, 0x8d, 0x45, 0x5c, 0xae, 0x36, 0xcb, 0x51, 0x00, 0x19, 0x11, 0x20,
	0x80, 0x28, 0x62, 0xb9, 0x89, 0x94, 0x52, 0x42, 0x4d, 0xed, 0xe6, 0xf5, 0xcb, 0x0b, 0x3c, 0xab,
	0xe2, 0x32, 0x81, 0x8a, 0x08, 0xc0, 0xca, 0xc4, 0xdf, 0x58, 0x1a, 0x87, 0xac, 0x84, 0x62, 0x4d,
	0x4e, 0xe3, 0x61, 0xef, 0xa1, 0x00, 0x10, 0x29, 0x3f, 0xdf, 0xce, 0x33, 0x59, 0x7e, 0x6f, 0x4d,
	0xfb, 0xaf, 0x86, 0xba, 0xef, 0x0f, 0xcf, 0x9c, 0xc5, 0x29, 0x77, 0xf9, 0xd7, 0x0d, 0x5f, 0x97,
	0x86, 0x44, 0x57, 0x92, 0x95, 0x2b, 0xb3, 0xf3, 0xb8, 0xe3, 0xe8, 0xef, 0x68, 0x53, 0x5b, 0x77,
	0x23, 0x28, 0xb2, 0xb1, 0xbd, 0x3f, 0xb5, 0xff, 0xfc, 0xb6, 0xe6, 0xe8, 0x83, 0xef, 0x38, 0x94,
	0x78, 0x3e, 0x25, 0x63, 0x3a, 0xa0, 0x13, 0xdb, 0x7e, 0xf3, 0x96, 0x6e, 0x69, 0x41, 0xf3, 0xc5,
	0xb0, 0x3f, 0xec, 0x6f, 0x1d, 0x4a, 0xfa, 0x5b, 0x8f, 0xe1, 0x1f, 0x53, 0xfc, 0x79, 0x31, 0x76,
	0x28, 0xf5, 0x7c, 0x4a, 0xff, 0x27, 0x07, 0x3d, 0x57, 0x25, 0x19, 0x04, 0xdd, 0xe1, 0x79, 0x00,
	0x61, 0x9c, 0x0b, 0x53, 0x53, 0xa9, 0xdd, 0xa6, 0xb6, 0xee, 0xb7, 0xa9, 0x47, 0xc7, 0x76, 0x4f,
	0x90, 0xf1, 0x08, 0x69, 0x20, 0xcd, 0x5b, 0x0a, 0xbd, 0xd7, 0xd4, 0x96, 0xde, 0xa2, 0x20, 0x6d,
	0x57, 0x03, 0x69, 0xfc, 0xea, 0x20, 0x5d, 0xb2, 0x20, 0x61, 0x82, 0xcf, 0x43, 0xf3, 0x4a, 0x61,
	0x49, 0x53, 0x5b, 0x0f, 0x8e, 0x3d, 0x0e, 0xd6, 0xbe, 0xcc, 0x27, 0xf4, 0xd1, 0xf7, 0x18, 0x8e,
	0xa6, 0x78, 0xf6, 0x04, 0xbf, 0x5e, 0xfc, 0x7c, 0xb5, 0xc3, 0x93, 0x4b, 0xfd, 0xe2, 0x86, 0xfa,
	0xe9, 0xb3, 0x5d, 0xcf, 0x3d, 0xa7, 0x2f, 0x6f, 0xab, 0xcf, 0x7e, 0xfe, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xbf, 0xa7, 0x1f, 0x12, 0x02, 0x00, 0x00,
}
