// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_archive.proto

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
//DeleteArchive请求
type DeleteArchiveRequest struct {
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

func (m *DeleteArchiveRequest) Reset()         { *m = DeleteArchiveRequest{} }
func (m *DeleteArchiveRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveRequest) ProtoMessage()    {}
func (*DeleteArchiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56973be536a39dc3, []int{0}
}
func (m *DeleteArchiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveRequest.Unmarshal(m, b)
}
func (m *DeleteArchiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveRequest.Merge(m, src)
}
func (m *DeleteArchiveRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveRequest.Size(m)
}
func (m *DeleteArchiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveRequest proto.InternalMessageInfo

func (m *DeleteArchiveRequest) GetDeleteFile() string {
	if m != nil {
		return m.DeleteFile
	}
	return ""
}

func (m *DeleteArchiveRequest) GetLastVersion() string {
	if m != nil {
		return m.LastVersion
	}
	return ""
}

func (m *DeleteArchiveRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *DeleteArchiveRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//DeleteArchive返回
type DeleteArchiveResponse struct {
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

func (m *DeleteArchiveResponse) Reset()         { *m = DeleteArchiveResponse{} }
func (m *DeleteArchiveResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveResponse) ProtoMessage()    {}
func (*DeleteArchiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56973be536a39dc3, []int{1}
}
func (m *DeleteArchiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveResponse.Unmarshal(m, b)
}
func (m *DeleteArchiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveResponse.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveResponse.Merge(m, src)
}
func (m *DeleteArchiveResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveResponse.Size(m)
}
func (m *DeleteArchiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveResponse proto.InternalMessageInfo

func (m *DeleteArchiveResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteArchiveResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteArchiveResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteArchiveResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//
//DeleteArchiveApi返回
type DeleteArchiveResponseWrapper struct {
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
	Data                 *DeleteArchiveResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeleteArchiveResponseWrapper) Reset()         { *m = DeleteArchiveResponseWrapper{} }
func (m *DeleteArchiveResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteArchiveResponseWrapper) ProtoMessage()    {}
func (*DeleteArchiveResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_56973be536a39dc3, []int{2}
}
func (m *DeleteArchiveResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArchiveResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteArchiveResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArchiveResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteArchiveResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchiveResponseWrapper.Merge(m, src)
}
func (m *DeleteArchiveResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteArchiveResponseWrapper.Size(m)
}
func (m *DeleteArchiveResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchiveResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchiveResponseWrapper proto.InternalMessageInfo

func (m *DeleteArchiveResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteArchiveResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteArchiveResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteArchiveResponseWrapper) GetData() *DeleteArchiveResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteArchiveRequest)(nil), "archive.DeleteArchiveRequest")
	proto.RegisterType((*DeleteArchiveResponse)(nil), "archive.DeleteArchiveResponse")
	proto.RegisterType((*DeleteArchiveResponseWrapper)(nil), "archive.DeleteArchiveResponseWrapper")
}

func init() { proto.RegisterFile("delete_archive.proto", fileDescriptor_56973be536a39dc3) }

var fileDescriptor_56973be536a39dc3 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0x25, 0xb7, 0xbd, 0x5e, 0xee, 0x44, 0xf4, 0x3a, 0xdc, 0x4a, 0x28, 0x62, 0x64, 0x14, 0x71,
	0x61, 0x1a, 0xad, 0x3f, 0x54, 0x37, 0xd2, 0xaa, 0x85, 0x2e, 0x1d, 0x41, 0x17, 0xa2, 0x32, 0x4d,
	0x26, 0x69, 0x68, 0xd2, 0x89, 0x33, 0xd3, 0x56, 0x94, 0x3e, 0x80, 0x1b, 0x5f, 0xc1, 0x37, 0x8b,
	0xe0, 0xd2, 0x65, 0x9e, 0x40, 0x3a, 0x93, 0x36, 0x43, 0x29, 0xa2, 0x8b, 0xae, 0x32, 0xdf, 0x77,
	0xce, 0xe1, 0x3b, 0x73, 0xbe, 0x0c, 0x38, 0x0f, 0x69, 0x4a, 0x25, 0xfd, 0x48, 0x78, 0x30, 0x49,
	0x16, 0xb4, 0x93, 0x73, 0x26, 0x19, 0x3c, 0xa9, 0xca, 0xb6, 0x17, 0x27, 0x72, 0x32, 0x1f, 0x77,
	0x02, 0x96, 0xf9, 0x31, 0x8b, 0x99, 0xaf, 0xf0, 0xf1, 0x3c, 0x52, 0x95, 0x2a, 0xd4, 0x49, 0xeb,
	0xda, 0x8f, 0x0d, 0x7a, 0xb6, 0x4c, 0xe4, 0x94, 0x2d, 0xfd, 0x98, 0x79, 0x0a, 0xf4, 0x16, 0x24,
	0x4d, 0x42, 0x22, 0x19, 0x17, 0xfe, 0xf6, 0xa8, 0x75, 0xe8, 0x47, 0x03, 0x9c, 0xbf, 0x50, 0x46,
	0xfa, 0x7a, 0x30, 0xa6, 0x9f, 0xe6, 0x54, 0x48, 0xf8, 0x08, 0x00, 0x6d, 0x70, 0x98, 0xa4, 0xd4,
	0xb1, 0x6e, 0x58, 0x77, 0x4e, 0x07, 0xad, 0xb2, 0x70, 0xaf, 0x44, 0x8c, 0x67, 0x4f, 0x51, 0x8d,
	0x21, 0x6c, 0x10, 0x61, 0x0f, 0xd8, 0x29, 0x11, 0xf2, 0x0d, 0xe5, 0x22, 0x61, 0x33, 0xe7, 0x48,
	0xe9, 0xae, 0x96, 0x85, 0x0b, 0xb5, 0xce, 0x00, 0x11, 0x36, 0xa9, 0xf0, 0x9b, 0x05, 0x4e, 0x73,
	0x12, 0x4c, 0x49, 0x4c, 0x47, 0xa1, 0xd3, 0x50, 0xc2, 0x69, 0x59, 0xb8, 0x67, 0x5a, 0xb8, 0x85,
	0xd0, 0xaf, 0x9f, 0xee, 0x6b, 0xf0, 0xea, 0xc3, 0x3b, 0xe2, 0x45, 0x7d, 0x6f, 0x78, 0xcf, 0x7b,
	0xf2, 0xfe, 0x6b, 0x6f, 0xe5, 0x3d, 0x33, 0xeb, 0x87, 0xff, 0x59, 0xdf, 0xef, 0xae, 0x6e, 0xe1,
	0x7a, 0xba, 0xf2, 0xb2, 0xd0, 0xbe, 0x46, 0xa1, 0xd3, 0xdc, 0xf5, 0xb2, 0x85, 0x0e, 0xe7, 0xa5,
	0x1e, 0xf1, 0xfd, 0x08, 0xb4, 0x76, 0x36, 0x24, 0x72, 0x36, 0x13, 0x14, 0xde, 0x04, 0xcd, 0x80,
	0x85, 0x7a, 0x39, 0xc7, 0x83, 0xcb, 0x65, 0xe1, 0xda, 0xda, 0xdf, 0xba, 0x8b, 0xb0, 0x02, 0xe1,
	0x6d, 0x70, 0x4c, 0x39, 0x67, 0xbc, 0x5a, 0xc5, 0x59, 0x59, 0xb8, 0x17, 0x35, 0x4b, 0xb5, 0x11,
	0xd6, 0x30, 0xbc, 0x0b, 0x4e, 0x32, 0x2a, 0x04, 0x89, 0x69, 0x95, 0x3d, 0x2c, 0x0b, 0xf7, 0x92,
	0x66, 0x56, 0x00, 0xc2, 0x1b, 0x0a, 0xfc, 0x02, 0x9a, 0x21, 0x91, 0xa4, 0x8a, 0x26, 0xaa, 0x47,
	0xaf, 0xbb, 0x07, 0x4b, 0x45, 0xcd, 0x44, 0xbf, 0x2d, 0x70, 0x6d, 0x6f, 0x20, 0x6f, 0x39, 0xc9,
	0x73, 0xca, 0xff, 0x2d, 0x97, 0x1e, 0xb0, 0xd7, 0xdf, 0x97, 0x9f, 0xf3, 0x94, 0x24, 0x7b, 0x7e,
	0x54, 0x03, 0x44, 0xd8, 0xa4, 0xd6, 0x89, 0x36, 0xfe, 0x9e, 0xe8, 0x73, 0x23, 0x23, 0xbb, 0x7b,
	0xbd, 0xb3, 0x79, 0xe8, 0x7b, 0xbd, 0x9b, 0x36, 0xd5, 0x2d, 0xf5, 0x65, 0xc7, 0x17, 0xd4, 0x33,
	0x7d, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x63, 0x90, 0x67, 0xd5, 0x2e, 0x04, 0x00, 0x00,
}
