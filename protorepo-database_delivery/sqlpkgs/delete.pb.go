// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package sqlpkgs

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
//DeleteSQLPackage请求
type DeleteSQLPackageRequest struct {
	//
	//SQL包ID
	PkgId                string   `protobuf:"bytes,1,opt,name=pkgId,proto3" json:"pkgId" form:"pkgId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSQLPackageRequest) Reset()         { *m = DeleteSQLPackageRequest{} }
func (m *DeleteSQLPackageRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSQLPackageRequest) ProtoMessage()    {}
func (*DeleteSQLPackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteSQLPackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSQLPackageRequest.Unmarshal(m, b)
}
func (m *DeleteSQLPackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSQLPackageRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSQLPackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSQLPackageRequest.Merge(m, src)
}
func (m *DeleteSQLPackageRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSQLPackageRequest.Size(m)
}
func (m *DeleteSQLPackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSQLPackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSQLPackageRequest proto.InternalMessageInfo

func (m *DeleteSQLPackageRequest) GetPkgId() string {
	if m != nil {
		return m.PkgId
	}
	return ""
}

//
//DeleteSQLPackageApi返回
type DeleteSQLPackageResponseWrapper struct {
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

func (m *DeleteSQLPackageResponseWrapper) Reset()         { *m = DeleteSQLPackageResponseWrapper{} }
func (m *DeleteSQLPackageResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteSQLPackageResponseWrapper) ProtoMessage()    {}
func (*DeleteSQLPackageResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteSQLPackageResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSQLPackageResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteSQLPackageResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSQLPackageResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteSQLPackageResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSQLPackageResponseWrapper.Merge(m, src)
}
func (m *DeleteSQLPackageResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteSQLPackageResponseWrapper.Size(m)
}
func (m *DeleteSQLPackageResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSQLPackageResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSQLPackageResponseWrapper proto.InternalMessageInfo

func (m *DeleteSQLPackageResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteSQLPackageResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteSQLPackageResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteSQLPackageResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteSQLPackageRequest)(nil), "sqlpkgs.DeleteSQLPackageRequest")
	proto.RegisterType((*DeleteSQLPackageResponseWrapper)(nil), "sqlpkgs.DeleteSQLPackageResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0xc9, 0x7b, 0xed, 0x7b, 0xbc, 0x69, 0x79, 0x96, 0x59, 0xd4, 0x50, 0x17, 0x29, 0x63,
	0x91, 0x6e, 0x92, 0xa8, 0x05, 0xd1, 0xba, 0x2b, 0x76, 0x21, 0xb8, 0xd0, 0x88, 0xb8, 0x10, 0x85,
	0x69, 0x33, 0x1d, 0x43, 0x92, 0xde, 0xe9, 0x64, 0x6a, 0xfd, 0x83, 0x5f, 0xb5, 0x82, 0xe0, 0x17,
	0xe8, 0x27, 0x90, 0xdc, 0x54, 0x5b, 0x70, 0x35, 0x73, 0xef, 0xef, 0x9c, 0xc3, 0xe1, 0x92, 0x6a,
	0x28, 0x12, 0x61, 0x84, 0xa7, 0x34, 0x18, 0xa0, 0x7f, 0xb3, 0x49, 0xa2, 0x62, 0x99, 0x35, 0x5c,
	0x19, 0x99, 0xfb, 0xe9, 0xc0, 0x1b, 0x42, 0xea, 0x4b, 0x90, 0xe0, 0x23, 0x1f, 0x4c, 0x47, 0x38,
	0xe1, 0x80, 0xbf, 0xc2, 0xd7, 0x38, 0x58, 0x93, 0xa7, 0xb3, 0xc8, 0xc4, 0x30, 0xf3, 0x25, 0xb8,
	0x08, 0xdd, 0x07, 0x9e, 0x44, 0x21, 0x37, 0xa0, 0x33, 0xff, 0xfb, 0xbb, 0xf4, 0x6d, 0x49, 0x00,
	0x99, 0x88, 0x55, 0xba, 0x48, 0x95, 0x79, 0x2a, 0x20, 0xbb, 0x22, 0x9b, 0x27, 0x58, 0xee, 0xf2,
	0xe2, 0xec, 0x9c, 0x0f, 0x63, 0x2e, 0x45, 0x20, 0x26, 0x53, 0x91, 0x19, 0xda, 0x25, 0x65, 0x15,
	0xcb, 0xd3, 0xd0, 0xb6, 0x9a, 0x56, 0xfb, 0x5f, 0xaf, 0xb5, 0x98, 0x3b, 0xd5, 0x11, 0xe8, 0xb4,
	0xcb, 0x70, 0xcd, 0xde, 0xdf, 0x9c, 0x1a, 0xf9, 0x7f, 0x77, 0xb3, 0xeb, 0x1e, 0x71, 0xf7, 0xf9,
	0xf6, 0x65, 0xaf, 0xf3, 0xda, 0x0a, 0x0a, 0x0b, 0xfb, 0xb0, 0x88, 0xf3, 0x33, 0x37, 0x53, 0x30,
	0xce, 0xc4, 0xb5, 0xe6, 0x4a, 0x09, 0x4d, 0xb7, 0x49, 0x69, 0x08, 0xa1, 0xc0, 0xf8, 0x72, 0x6f,
	0x63, 0x31, 0x77, 0x2a, 0x45, 0x7c, 0xbe, 0x65, 0x01, 0x42, 0x7a, 0x48, 0x2a, 0xf9, 0xdb, 0x7f,
	0x54, 0x09, 0x8f, 0xc6, 0xf6, 0x2f, 0xac, 0x52, 0x5f, 0xcc, 0x1d, 0xba, 0xd2, 0x2e, 0x21, 0x0b,
	0xd6, 0xa5, 0x74, 0x87, 0x94, 0x85, 0xd6, 0xa0, 0xed, 0xdf, 0xe8, 0xa9, 0xad, 0xea, 0xe3, 0x9a,
	0x05, 0x05, 0xa6, 0xc7, 0xa4, 0x14, 0x72, 0xc3, 0xed, 0x52, 0xd3, 0x6a, 0x57, 0xf6, 0xeb, 0x5e,
	0x71, 0x2d, 0xef, 0xeb, 0x5a, 0x5e, 0x3f, 0xbf, 0xd6, 0x7a, 0xbd, 0x5c, 0xcd, 0x02, 0x34, 0x0d,
	0xfe, 0xa0, 0xac, 0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x1d, 0xe7, 0x16, 0xe2, 0x01, 0x00,
	0x00,
}