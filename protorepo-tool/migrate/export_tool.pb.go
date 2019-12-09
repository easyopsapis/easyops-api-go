// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: export_tool.proto

package migrate

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
//ExportTool请求
type ExportToolRequest struct {
	//
	//版本Id
	VersionId string `protobuf:"bytes,1,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//兼容版本策略 latest=最新; 2.32=兼容2.32的版本
	Compatibility string `protobuf:"bytes,2,opt,name=compatibility,proto3" json:"compatibility" form:"compatibility"`
	//
	//工具Id，服务端自动生成
	ToolId               string   `protobuf:"bytes,3,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToolRequest) Reset()         { *m = ExportToolRequest{} }
func (m *ExportToolRequest) String() string { return proto.CompactTextString(m) }
func (*ExportToolRequest) ProtoMessage()    {}
func (*ExportToolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9124619b6a6924eb, []int{0}
}
func (m *ExportToolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToolRequest.Unmarshal(m, b)
}
func (m *ExportToolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToolRequest.Marshal(b, m, deterministic)
}
func (m *ExportToolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToolRequest.Merge(m, src)
}
func (m *ExportToolRequest) XXX_Size() int {
	return xxx_messageInfo_ExportToolRequest.Size(m)
}
func (m *ExportToolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToolRequest proto.InternalMessageInfo

func (m *ExportToolRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ExportToolRequest) GetCompatibility() string {
	if m != nil {
		return m.Compatibility
	}
	return ""
}

func (m *ExportToolRequest) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func init() {
	proto.RegisterType((*ExportToolRequest)(nil), "migrate.ExportToolRequest")
}

func init() { proto.RegisterFile("export_tool.proto", fileDescriptor_9124619b6a6924eb) }

var fileDescriptor_9124619b6a6924eb = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x99, 0xc2, 0x64, 0x81, 0x81, 0xad, 0x1e, 0xca, 0x3c, 0x54, 0x8a, 0x87, 0x81, 0x64,
	0x15, 0x07, 0x82, 0x1e, 0x44, 0x05, 0x85, 0x5e, 0x8b, 0x37, 0x51, 0x49, 0xd7, 0x34, 0x06, 0x13,
	0x5f, 0x4d, 0x5f, 0x37, 0x87, 0xf8, 0x55, 0xab, 0xf8, 0x11, 0xfa, 0x09, 0xc4, 0x64, 0x6c, 0xee,
	0xe0, 0xed, 0xbd, 0xfc, 0xde, 0xef, 0xcf, 0xcb, 0x23, 0x1e, 0x7f, 0x2b, 0xc1, 0xe0, 0x23, 0x02,
	0xa8, 0x51, 0x69, 0x00, 0xc1, 0xdf, 0xd2, 0x52, 0x18, 0x86, 0x7c, 0x40, 0x85, 0xc4, 0xa7, 0x3a,
	0x1b, 0x4d, 0x40, 0xc7, 0x02, 0x04, 0xc4, 0x96, 0x67, 0x75, 0x61, 0x3b, 0xdb, 0xd8, 0xca, 0x79,
	0x83, 0x93, 0x3f, 0xe3, 0x7a, 0x26, 0xf1, 0x19, 0x66, 0xb1, 0x00, 0x6a, 0x21, 0x9d, 0x32, 0x25,
	0x73, 0x86, 0x60, 0xaa, 0x78, 0x59, 0x2e, 0xbc, 0x3d, 0x01, 0x20, 0x14, 0x5f, 0xa5, 0x73, 0x5d,
	0xe2, 0xdc, 0xc1, 0xe8, 0xab, 0x43, 0xbc, 0x6b, 0xbb, 0xe2, 0x2d, 0x80, 0x4a, 0xf9, 0x6b, 0xcd,
	0x2b, 0xf4, 0x13, 0xd2, 0x9b, 0x72, 0x53, 0x49, 0x78, 0x49, 0xf2, 0xa0, 0xb3, 0xdf, 0x19, 0xf6,
	0xae, 0x0e, 0xdb, 0x26, 0xdc, 0x2e, 0xc0, 0xe8, 0xb3, 0x68, 0x89, 0xa2, 0xef, 0xcf, 0x70, 0x87,
	0x78, 0x0f, 0x77, 0x8c, 0x16, 0x97, 0xf4, 0xe6, 0x88, 0x9e, 0xde, 0xbf, 0x8f, 0x8f, 0x3f, 0x0e,
	0xd2, 0x95, 0xed, 0x9f, 0x93, 0xfe, 0x04, 0x74, 0xc9, 0x50, 0x66, 0x52, 0x49, 0x9c, 0x07, 0x1b,
	0x36, 0x2e, 0x68, 0x9b, 0x70, 0xd7, 0xc5, 0xad, 0xe1, 0x28, 0x5d, 0x1f, 0xf7, 0x2f, 0x48, 0xf7,
	0xf7, 0x76, 0x49, 0x1e, 0x6c, 0x5a, 0x71, 0xd8, 0x36, 0x61, 0xdf, 0x89, 0xee, 0xfd, 0xdf, 0x25,
	0x16, 0x5e, 0xd6, 0xb5, 0x3f, 0x1d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x36, 0x19, 0x49, 0x34,
	0x8b, 0x01, 0x00, 0x00,
}