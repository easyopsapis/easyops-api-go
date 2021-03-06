// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_delete_versions.proto

package version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//BatchDeleteVersions请求
type BatchDeleteVersionsRequest struct {
	//
	//版本ids，以逗号分隔
	VersionIds string `protobuf:"bytes,1,opt,name=versionIds,proto3" json:"versionIds" form:"versionIds"`
	//
	//是否彻底删除文件,默认为false
	DeleteFile bool `protobuf:"varint,2,opt,name=deleteFile,proto3" json:"deleteFile" form:"deleteFile"`
	//
	//程序包Id
	PackageId            string   `protobuf:"bytes,3,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchDeleteVersionsRequest) Reset()         { *m = BatchDeleteVersionsRequest{} }
func (m *BatchDeleteVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*BatchDeleteVersionsRequest) ProtoMessage()    {}
func (*BatchDeleteVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c07b0381ee8a942, []int{0}
}
func (m *BatchDeleteVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchDeleteVersionsRequest.Unmarshal(m, b)
}
func (m *BatchDeleteVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchDeleteVersionsRequest.Marshal(b, m, deterministic)
}
func (m *BatchDeleteVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchDeleteVersionsRequest.Merge(m, src)
}
func (m *BatchDeleteVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_BatchDeleteVersionsRequest.Size(m)
}
func (m *BatchDeleteVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchDeleteVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchDeleteVersionsRequest proto.InternalMessageInfo

func (m *BatchDeleteVersionsRequest) GetVersionIds() string {
	if m != nil {
		return m.VersionIds
	}
	return ""
}

func (m *BatchDeleteVersionsRequest) GetDeleteFile() bool {
	if m != nil {
		return m.DeleteFile
	}
	return false
}

func (m *BatchDeleteVersionsRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//BatchDeleteVersionsApi返回
type BatchDeleteVersionsResponseWrapper struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//返回码解释
	CodeExplain string `protobuf:"bytes,2,opt,name=codeExplain,proto3" json:"codeExplain" form:"codeExplain"`
	//
	//错误详情
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回数据
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BatchDeleteVersionsResponseWrapper) Reset()         { *m = BatchDeleteVersionsResponseWrapper{} }
func (m *BatchDeleteVersionsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*BatchDeleteVersionsResponseWrapper) ProtoMessage()    {}
func (*BatchDeleteVersionsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c07b0381ee8a942, []int{1}
}
func (m *BatchDeleteVersionsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchDeleteVersionsResponseWrapper.Unmarshal(m, b)
}
func (m *BatchDeleteVersionsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchDeleteVersionsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *BatchDeleteVersionsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchDeleteVersionsResponseWrapper.Merge(m, src)
}
func (m *BatchDeleteVersionsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_BatchDeleteVersionsResponseWrapper.Size(m)
}
func (m *BatchDeleteVersionsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchDeleteVersionsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BatchDeleteVersionsResponseWrapper proto.InternalMessageInfo

func (m *BatchDeleteVersionsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchDeleteVersionsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *BatchDeleteVersionsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BatchDeleteVersionsResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchDeleteVersionsRequest)(nil), "version.BatchDeleteVersionsRequest")
	proto.RegisterType((*BatchDeleteVersionsResponseWrapper)(nil), "version.BatchDeleteVersionsResponseWrapper")
}

func init() { proto.RegisterFile("batch_delete_versions.proto", fileDescriptor_4c07b0381ee8a942) }

var fileDescriptor_4c07b0381ee8a942 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6a, 0xd4, 0x40,
	0x18, 0x25, 0x75, 0xab, 0xee, 0xac, 0x60, 0x1d, 0xb0, 0x84, 0x2d, 0x98, 0x30, 0x16, 0x89, 0x3f,
	0x93, 0xd4, 0xfa, 0x43, 0x57, 0x2f, 0xa4, 0xc1, 0x16, 0x7a, 0xe9, 0x08, 0x8a, 0x96, 0x76, 0x99,
	0x24, 0xb3, 0xd9, 0xb0, 0xc9, 0x4e, 0x9c, 0xcc, 0xee, 0xfa, 0xc3, 0x3e, 0x80, 0xaf, 0xe4, 0xc3,
	0x44, 0xf0, 0xca, 0xeb, 0x3c, 0x81, 0x64, 0x26, 0x6e, 0x72, 0xe1, 0x8d, 0x57, 0x99, 0xf3, 0x9d,
	0x73, 0xbe, 0x1c, 0x0e, 0x1f, 0xd8, 0x0b, 0xa8, 0x0c, 0xa7, 0xe3, 0x88, 0xa5, 0x4c, 0xb2, 0xf1,
	0x92, 0x89, 0x22, 0xe1, 0xf3, 0xc2, 0xcd, 0x05, 0x97, 0x1c, 0x5e, 0x6b, 0xf0, 0x10, 0xc7, 0x89,
	0x9c, 0x2e, 0x02, 0x37, 0xe4, 0x99, 0x17, 0xf3, 0x98, 0x7b, 0x8a, 0x0f, 0x16, 0x13, 0x85, 0x14,
	0x50, 0x2f, 0xed, 0x1b, 0x3e, 0xef, 0xc8, 0xb3, 0x55, 0x22, 0x67, 0x7c, 0xe5, 0xc5, 0x1c, 0x2b,
	0x12, 0x2f, 0x69, 0x9a, 0x44, 0x54, 0x72, 0x51, 0x78, 0x9b, 0x67, 0xe3, 0xdb, 0x8b, 0x39, 0x8f,
	0x53, 0xd6, 0x6e, 0x67, 0x59, 0x2e, 0xbf, 0x68, 0x12, 0xfd, 0xd8, 0x02, 0x43, 0xbf, 0x0e, 0xfb,
	0x5a, 0x65, 0x7d, 0xd7, 0x44, 0x25, 0xec, 0xd3, 0x82, 0x15, 0x12, 0x7e, 0x00, 0xa0, 0x49, 0x7b,
	0x16, 0x15, 0xa6, 0x61, 0x1b, 0x4e, 0xdf, 0x1f, 0x55, 0xa5, 0x75, 0x6b, 0xc2, 0x45, 0xf6, 0x02,
	0xb5, 0x1c, 0xfa, 0xf5, 0xd3, 0xb2, 0xc1, 0x9d, 0xcb, 0xf3, 0x03, 0x3c, 0xa2, 0xf8, 0xeb, 0x31,
	0xfe, 0x78, 0xf1, 0xd0, 0x79, 0xd4, 0x45, 0xf7, 0x1f, 0xec, 0x93, 0xce, 0x32, 0xf8, 0x0c, 0x00,
	0xdd, 0xcf, 0x69, 0x92, 0x32, 0x73, 0xcb, 0x36, 0x9c, 0xeb, 0xfe, 0xed, 0x76, 0x75, 0xcb, 0x21,
	0xd2, 0x11, 0xc2, 0xef, 0x06, 0xe8, 0xe7, 0x34, 0x9c, 0xd1, 0x98, 0x9d, 0x45, 0xe6, 0x15, 0x95,
	0x68, 0x56, 0x95, 0xd6, 0x8e, 0xb6, 0x6d, 0xa8, 0x3a, 0xd0, 0x5b, 0xf0, 0xe6, 0xf2, 0x9c, 0xe2,
	0xc9, 0x31, 0x3e, 0x3d, 0xc0, 0xa3, 0x8b, 0x6f, 0x47, 0x6b, 0xfc, 0xaa, 0x8b, 0x9f, 0xfe, 0x27,
	0x7e, 0x7c, 0xb8, 0xde, 0x27, 0xed, 0xdf, 0xd1, 0x6f, 0x03, 0xa0, 0x7f, 0x96, 0x57, 0xe4, 0x7c,
	0x5e, 0xb0, 0xf7, 0x82, 0xe6, 0x39, 0x13, 0xf0, 0x2e, 0xe8, 0x85, 0x3c, 0x62, 0xaa, 0xbe, 0x6d,
	0xff, 0x66, 0x55, 0x5a, 0x03, 0x1d, 0xb6, 0x9e, 0x22, 0xa2, 0x48, 0x78, 0x04, 0x06, 0xf5, 0xf7,
	0xe4, 0x73, 0x9e, 0xd2, 0x64, 0xae, 0xfa, 0xe8, 0xfb, 0xbb, 0x55, 0x69, 0xc1, 0x56, 0xdb, 0x90,
	0x88, 0x74, 0xa5, 0xf0, 0x1e, 0xd8, 0x66, 0x42, 0x70, 0xd1, 0x94, 0xb1, 0x53, 0x95, 0xd6, 0x0d,
	0xed, 0x51, 0x63, 0x44, 0x34, 0x0d, 0x5f, 0x82, 0x5e, 0x44, 0x25, 0x35, 0x7b, 0xb6, 0xe1, 0x0c,
	0x0e, 0x77, 0x5d, 0x7d, 0x16, 0xee, 0xdf, 0xb3, 0x70, 0x4f, 0xea, 0xb3, 0xe8, 0xc6, 0xab, 0xd5,
	0x88, 0x28, 0x53, 0x70, 0x55, 0xc9, 0x9e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x04, 0x92, 0x32,
	0x7a, 0xda, 0x02, 0x00, 0x00,
}
