// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_archive_v2.proto

package archive

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
//DeleteArchiveV2请求
type DeleteArchiveV2Request struct {
	//
	//是否彻底删除文件（true|false）
	DeleteFile string `protobuf:"bytes,1,opt,name=deleteFile,proto3" json:"deleteFile" form:"deleteFile"`
	//
	//是否最后一个版本
	LastVersion string `protobuf:"bytes,2,opt,name=lastVersion,proto3" json:"lastVersion" form:"lastVersion"`
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

func (m *DeleteArchiveV2Request) Reset()         { *m = DeleteArchiveV2Request{} }
func (m *DeleteArchiveV2Request) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveV2Request) ProtoMessage()    {}
func (*DeleteArchiveV2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ad2cef03265bff, []int{0}
}
func (m *DeleteArchiveV2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveV2Request.Unmarshal(m, b)
}
func (m *DeleteArchiveV2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveV2Request.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveV2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveV2Request.Merge(m, src)
}
func (m *DeleteArchiveV2Request) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveV2Request.Size(m)
}
func (m *DeleteArchiveV2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveV2Request.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveV2Request proto.InternalMessageInfo

func (m *DeleteArchiveV2Request) GetDeleteFile() string {
	if m != nil {
		return m.DeleteFile
	}
	return ""
}

func (m *DeleteArchiveV2Request) GetLastVersion() string {
	if m != nil {
		return m.LastVersion
	}
	return ""
}

func (m *DeleteArchiveV2Request) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *DeleteArchiveV2Request) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//DeleteArchiveV2返回
type DeleteArchiveV2Response struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回信息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//返回数据,成功删除的版本Id
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArchiveV2Response) Reset()         { *m = DeleteArchiveV2Response{} }
func (m *DeleteArchiveV2Response) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveV2Response) ProtoMessage()    {}
func (*DeleteArchiveV2Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ad2cef03265bff, []int{1}
}
func (m *DeleteArchiveV2Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveV2Response.Unmarshal(m, b)
}
func (m *DeleteArchiveV2Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveV2Response.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveV2Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveV2Response.Merge(m, src)
}
func (m *DeleteArchiveV2Response) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveV2Response.Size(m)
}
func (m *DeleteArchiveV2Response) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveV2Response.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveV2Response proto.InternalMessageInfo

func (m *DeleteArchiveV2Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteArchiveV2Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteArchiveV2Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteArchiveV2Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//
//DeleteArchiveV2Api返回
type DeleteArchiveV2ResponseWrapper struct {
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
	Data                 *DeleteArchiveV2Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DeleteArchiveV2ResponseWrapper) Reset()         { *m = DeleteArchiveV2ResponseWrapper{} }
func (m *DeleteArchiveV2ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveV2ResponseWrapper) ProtoMessage()    {}
func (*DeleteArchiveV2ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ad2cef03265bff, []int{2}
}
func (m *DeleteArchiveV2ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveV2ResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteArchiveV2ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveV2ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveV2ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveV2ResponseWrapper.Merge(m, src)
}
func (m *DeleteArchiveV2ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveV2ResponseWrapper.Size(m)
}
func (m *DeleteArchiveV2ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveV2ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveV2ResponseWrapper proto.InternalMessageInfo

func (m *DeleteArchiveV2ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteArchiveV2ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteArchiveV2ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteArchiveV2ResponseWrapper) GetData() *DeleteArchiveV2Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteArchiveV2Request)(nil), "archive.DeleteArchiveV2Request")
	proto.RegisterType((*DeleteArchiveV2Response)(nil), "archive.DeleteArchiveV2Response")
	proto.RegisterType((*DeleteArchiveV2ResponseWrapper)(nil), "archive.DeleteArchiveV2ResponseWrapper")
}

func init() { proto.RegisterFile("delete_archive_v2.proto", fileDescriptor_a7ad2cef03265bff) }

var fileDescriptor_a7ad2cef03265bff = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xd6, 0x8e, 0x69, 0x0e, 0x82, 0x61, 0x89, 0x2d, 0xda, 0x05, 0x99, 0x0c, 0x42, 0x5c,
	0x90, 0x04, 0xc2, 0x8f, 0x0a, 0x37, 0x68, 0x13, 0x9b, 0xb4, 0x4b, 0x8c, 0x34, 0x2e, 0x10, 0x54,
	0x6e, 0xe2, 0xa4, 0x51, 0x93, 0x3a, 0xd8, 0x6e, 0x8a, 0x40, 0x7d, 0x00, 0xee, 0x78, 0x05, 0x5e,
	0x2c, 0x48, 0x3c, 0x00, 0x17, 0x79, 0x02, 0x54, 0x3b, 0x34, 0x16, 0x50, 0x04, 0x17, 0xbd, 0x8a,
	0xbf, 0xef, 0x9c, 0xa3, 0xef, 0xf8, 0x7c, 0x31, 0x38, 0x88, 0x69, 0x4e, 0x25, 0x1d, 0x12, 0x1e,
	0x8d, 0xb3, 0x8a, 0x0e, 0xab, 0xd0, 0x2f, 0x39, 0x93, 0x0c, 0xee, 0xb4, 0x9d, 0x43, 0x2f, 0xcd,
	0xe4, 0x78, 0x36, 0xf2, 0x23, 0x56, 0x04, 0x29, 0x4b, 0x59, 0xa0, 0xf0, 0xd1, 0x2c, 0x51, 0x95,
	0x2a, 0xd4, 0x49, 0xeb, 0x0e, 0x1f, 0x1b, 0xf4, 0x62, 0x9e, 0xc9, 0x09, 0x9b, 0x07, 0x29, 0xf3,
	0x14, 0xe8, 0x55, 0x24, 0xcf, 0x62, 0x22, 0x19, 0x17, 0xc1, 0xea, 0xa8, 0x75, 0xe8, 0x4b, 0x0f,
	0xec, 0x3f, 0x57, 0x5e, 0x8e, 0xf5, 0xe0, 0x8b, 0x10, 0xd3, 0x77, 0x33, 0x2a, 0x24, 0x7c, 0x04,
	0x80, 0x76, 0x79, 0x96, 0xe5, 0xd4, 0xb1, 0x8e, 0xac, 0x3b, 0xbb, 0x27, 0xd7, 0x9b, 0xda, 0xbd,
	0x96, 0x30, 0x5e, 0x3c, 0x45, 0x1d, 0x86, 0xb0, 0x41, 0x84, 0x03, 0x60, 0xe7, 0x44, 0xc8, 0x0b,
	0xca, 0x45, 0xc6, 0xa6, 0xce, 0x96, 0xd2, 0xed, 0x37, 0xb5, 0x0b, 0xb5, 0xce, 0x00, 0x11, 0x36,
	0xa9, 0xf0, 0x93, 0x05, 0x76, 0x4b, 0x12, 0x4d, 0x48, 0x4a, 0xcf, 0x63, 0xa7, 0xa7, 0x84, 0x93,
	0xa6, 0x76, 0xf7, 0xb4, 0x70, 0x05, 0xa1, 0x6f, 0x5f, 0xdd, 0x97, 0xe0, 0xc5, 0xdb, 0xd7, 0xc4,
	0x4b, 0x8e, 0xbd, 0xb3, 0x7b, 0xde, 0x93, 0x37, 0x1f, 0x07, 0x0b, 0xef, 0x99, 0x59, 0x3f, 0xfc,
	0xcf, 0xfa, 0x7e, 0xb8, 0xb8, 0x85, 0xbb, 0xe9, 0xca, 0x4b, 0xa5, 0x7d, 0x9d, 0xc7, 0x4e, 0xff,
	0x57, 0x2f, 0x2b, 0x68, 0x73, 0x5e, 0xba, 0x11, 0x9f, 0xb7, 0xc0, 0xc1, 0x6f, 0x3b, 0x12, 0x25,
	0x9b, 0x0a, 0x0a, 0x6f, 0x82, 0x7e, 0xc4, 0x62, 0xbd, 0x9e, 0xed, 0x93, 0xab, 0x4d, 0xed, 0xda,
	0xda, 0xe1, 0xb2, 0x8b, 0xb0, 0x02, 0xe1, 0x6d, 0xb0, 0x4d, 0x39, 0x67, 0xbc, 0x5d, 0xc6, 0x5e,
	0x53, 0xbb, 0x97, 0x35, 0x4b, 0xb5, 0x11, 0xd6, 0x30, 0xbc, 0x0b, 0x76, 0x0a, 0x2a, 0x04, 0x49,
	0x69, 0x9b, 0x3e, 0x6c, 0x6a, 0xf7, 0x8a, 0x66, 0xb6, 0x00, 0xc2, 0x3f, 0x29, 0xf0, 0x03, 0xe8,
	0xc7, 0x44, 0x92, 0x36, 0x9c, 0xa4, 0x1b, 0xbd, 0xec, 0x6e, 0x2c, 0x17, 0x35, 0x13, 0x7d, 0xb7,
	0xc0, 0x8d, 0x35, 0x91, 0xbc, 0xe2, 0xa4, 0x2c, 0x29, 0xff, 0xb7, 0x64, 0x06, 0xc0, 0x5e, 0x7e,
	0x4f, 0xdf, 0x97, 0x39, 0xc9, 0xfe, 0xf0, 0xb3, 0x1a, 0x20, 0xc2, 0x26, 0xb5, 0xcb, 0xb4, 0xf7,
	0xf7, 0x4c, 0x4f, 0x8d, 0x94, 0xec, 0xf0, 0xc8, 0x6f, 0xdf, 0xb7, 0xbf, 0xc6, 0xbd, 0x69, 0x54,
	0xdd, 0x54, 0x5f, 0x78, 0x74, 0x49, 0x3d, 0xd7, 0x07, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0x4a, 0xba, 0x9a, 0x39, 0x04, 0x00, 0x00,
}
