// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: download_workspace_file.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//FileDownload请求
type FileDownloadRequest struct {
	//
	//只查看指定文件内容差异
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//程序包Id
	PackageId            string   `protobuf:"bytes,3,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileDownloadRequest) Reset()         { *m = FileDownloadRequest{} }
func (m *FileDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*FileDownloadRequest) ProtoMessage()    {}
func (*FileDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecfe951deb46bd06, []int{0}
}
func (m *FileDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDownloadRequest.Unmarshal(m, b)
}
func (m *FileDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDownloadRequest.Marshal(b, m, deterministic)
}
func (m *FileDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDownloadRequest.Merge(m, src)
}
func (m *FileDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_FileDownloadRequest.Size(m)
}
func (m *FileDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileDownloadRequest proto.InternalMessageInfo

func (m *FileDownloadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileDownloadRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *FileDownloadRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func init() {
	proto.RegisterType((*FileDownloadRequest)(nil), "workspace.FileDownloadRequest")
}

func init() { proto.RegisterFile("download_workspace_file.proto", fileDescriptor_ecfe951deb46bd06) }

var fileDescriptor_ecfe951deb46bd06 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xc9, 0x2f, 0xcf,
	0xcb, 0xc9, 0x4f, 0x4c, 0x89, 0x2f, 0xcf, 0x2f, 0xca, 0x2e, 0x2e, 0x48, 0x4c, 0x4e, 0x8d, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x8b, 0x4a, 0xe9, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83,
	0x55, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x29, 0x65, 0x86, 0xa4, 0x3c,
	0xb7, 0x3c, 0xb3, 0x24, 0x3b, 0xbf, 0x5c, 0x3f, 0x3d, 0x5f, 0x17, 0x2c, 0xa9, 0x5b, 0x96, 0x98,
	0x93, 0x99, 0x92, 0x58, 0x92, 0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0xf5, 0x49, 0xa7, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0x22, 0x4c, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0xcd, 0x60, 0xe2,
	0x12, 0x76, 0xcb, 0xcc, 0x49, 0x75, 0x81, 0x3a, 0x3a, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0xc8, 0x99, 0x8b, 0xa5, 0x20, 0xb1, 0x24, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xff,
	0xd3, 0x3d, 0x79, 0xee, 0xb4, 0xfc, 0xa2, 0x5c, 0x2b, 0x25, 0x90, 0xa8, 0xd2, 0xa3, 0xfb, 0xf2,
	0x32, 0x5c, 0x52, 0x71, 0x31, 0xfa, 0x1a, 0x1a, 0xd1, 0x71, 0x31, 0xfa, 0xb1, 0xda, 0x9a, 0x1a,
	0x31, 0xfa, 0x50, 0x96, 0x96, 0xa6, 0xbd, 0x4a, 0x10, 0x58, 0xb3, 0x90, 0x3e, 0x17, 0x47, 0x6a,
	0x5e, 0x72, 0x7e, 0x4a, 0x66, 0x5e, 0xba, 0x04, 0x13, 0xd8, 0x20, 0xe1, 0x4f, 0xf7, 0xe4, 0xf9,
	0x21, 0x06, 0xc1, 0x64, 0x94, 0x82, 0xe0, 0x8a, 0x84, 0xba, 0x18, 0xb9, 0x38, 0x0b, 0x12, 0x93,
	0xb3, 0x13, 0xd3, 0x53, 0x3d, 0x53, 0x24, 0x98, 0xc1, 0x5a, 0x72, 0x3e, 0xdd, 0x93, 0x17, 0x80,
	0xd9, 0x0d, 0x95, 0x02, 0x39, 0x20, 0x84, 0x2b, 0x28, 0x2e, 0xda, 0x3e, 0x51, 0xb7, 0xca, 0x51,
	0x37, 0xca, 0x40, 0xd7, 0x32, 0xb6, 0xda, 0xa2, 0x56, 0x17, 0x85, 0x6f, 0x42, 0x22, 0xdf, 0xd0,
	0xa8, 0x56, 0x25, 0x08, 0x61, 0x7d, 0x12, 0x1b, 0x38, 0x84, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x3c, 0xe0, 0xc3, 0xd1, 0x01, 0x00, 0x00,
}