// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: download_archive.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//DownloadArchive请求
type DownloadArchiveRequest struct {
	//
	//单独下载指定的文件
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//程序包Id
	PackageId string `protobuf:"bytes,3,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId            string   `protobuf:"bytes,4,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadArchiveRequest) Reset()         { *m = DownloadArchiveRequest{} }
func (m *DownloadArchiveRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadArchiveRequest) ProtoMessage()    {}
func (*DownloadArchiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a7f71e1611f77c, []int{0}
}
func (m *DownloadArchiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadArchiveRequest.Unmarshal(m, b)
}
func (m *DownloadArchiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadArchiveRequest.Marshal(b, m, deterministic)
}
func (m *DownloadArchiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadArchiveRequest.Merge(m, src)
}
func (m *DownloadArchiveRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadArchiveRequest.Size(m)
}
func (m *DownloadArchiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadArchiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadArchiveRequest proto.InternalMessageInfo

func (m *DownloadArchiveRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DownloadArchiveRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *DownloadArchiveRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *DownloadArchiveRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func init() {
	proto.RegisterType((*DownloadArchiveRequest)(nil), "archive.DownloadArchiveRequest")
}

func init() { proto.RegisterFile("download_archive.proto", fileDescriptor_03a7f71e1611f77c) }

var fileDescriptor_03a7f71e1611f77c = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xc9, 0x2f, 0xcf,
	0xcb, 0xc9, 0x4f, 0x4c, 0x89, 0x4f, 0x2c, 0x4a, 0xce, 0xc8, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xf2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98,
	0x03, 0x66, 0x41, 0xf4, 0x49, 0x99, 0x21, 0x29, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7,
	0x4f, 0xcf, 0xd7, 0x05, 0x4b, 0xea, 0x96, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x15,
	0xeb, 0xc3, 0x99, 0x50, 0x7d, 0xd2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0x08, 0xd3, 0x53, 0x73,
	0x0b, 0x4a, 0x2a, 0x21, 0x92, 0x4a, 0x57, 0x99, 0xb9, 0xc4, 0x5c, 0xa0, 0xee, 0x74, 0x84, 0xb8,
	0x2b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xa8, 0x80, 0x8b, 0xa5, 0x20, 0xb1, 0x24, 0x43,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x29, 0xe6, 0xd3, 0x3d, 0x79, 0xee, 0xb4, 0xfc, 0xa2, 0x5c,
	0x2b, 0x25, 0x90, 0xa8, 0xd2, 0xa3, 0xfb, 0xf2, 0x9e, 0x5c, 0xee, 0x71, 0x1a, 0x1a, 0x31, 0xfa,
	0xd1, 0x71, 0x31, 0xfa, 0x56, 0x31, 0x5a, 0x31, 0xf6, 0x4a, 0x4a, 0x36, 0x76, 0x31, 0x35, 0x31,
	0x45, 0x31, 0x79, 0xb1, 0xda, 0x9a, 0xda, 0x9a, 0x35, 0x1a, 0x31, 0xfa, 0x9a, 0x35, 0xd1, 0x89,
	0xba, 0x55, 0x8e, 0xba, 0x51, 0xb1, 0x56, 0x1a, 0x31, 0x31, 0xd1, 0x71, 0x31, 0x31, 0x98, 0x2a,
	0xb5, 0x54, 0x82, 0xc0, 0x36, 0x09, 0xe9, 0x73, 0x71, 0xa4, 0xe6, 0x25, 0xe7, 0xa7, 0x64, 0xe6,
	0xa5, 0x4b, 0x30, 0x81, 0x6d, 0x15, 0xfe, 0x74, 0x4f, 0x9e, 0x1f, 0x62, 0x2b, 0x4c, 0x46, 0x29,
	0x08, 0xae, 0x48, 0xa8, 0x8b, 0x91, 0x8b, 0xb3, 0x20, 0x31, 0x39, 0x3b, 0x31, 0x3d, 0xd5, 0x33,
	0x45, 0x82, 0x19, 0xac, 0x25, 0xe7, 0xd3, 0x3d, 0x79, 0x01, 0x98, 0x43, 0xa1, 0x52, 0x20, 0xd7,
	0x86, 0x70, 0x05, 0xc5, 0x45, 0xdb, 0x43, 0x1c, 0x63, 0xa0, 0x6b, 0x19, 0x5b, 0x6d, 0x51, 0xab,
	0x8b, 0xc2, 0x37, 0x21, 0x91, 0x6f, 0x68, 0x54, 0xab, 0x12, 0x84, 0xb0, 0x1e, 0xec, 0x98, 0xb2,
	0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0xcf, 0x14, 0x09, 0x16, 0x74, 0xc7, 0xc0, 0xa5, 0x68, 0xe8,
	0x18, 0xb8, 0x1d, 0x49, 0x6c, 0xe0, 0xe8, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xa4,
	0x4f, 0x97, 0x85, 0x02, 0x00, 0x00,
}