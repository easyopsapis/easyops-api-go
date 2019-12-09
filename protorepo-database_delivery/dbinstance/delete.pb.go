// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package dbinstance

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
//DeleteDBInstance请求
type DeleteDBInstanceRequest struct {
	//
	//数据库实例ID
	DbInstanceId string `protobuf:"bytes,1,opt,name=dbInstanceId,proto3" json:"dbInstanceId" form:"dbInstanceId"`
	//
	//数据库服务ID
	DbServiceId          string   `protobuf:"bytes,2,opt,name=dbServiceId,proto3" json:"dbServiceId" form:"dbServiceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDBInstanceRequest) Reset()         { *m = DeleteDBInstanceRequest{} }
func (m *DeleteDBInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDBInstanceRequest) ProtoMessage()    {}
func (*DeleteDBInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}
func (m *DeleteDBInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDBInstanceRequest.Unmarshal(m, b)
}
func (m *DeleteDBInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDBInstanceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDBInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDBInstanceRequest.Merge(m, src)
}
func (m *DeleteDBInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDBInstanceRequest.Size(m)
}
func (m *DeleteDBInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDBInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDBInstanceRequest proto.InternalMessageInfo

func (m *DeleteDBInstanceRequest) GetDbInstanceId() string {
	if m != nil {
		return m.DbInstanceId
	}
	return ""
}

func (m *DeleteDBInstanceRequest) GetDbServiceId() string {
	if m != nil {
		return m.DbServiceId
	}
	return ""
}

//
//DeleteDBInstanceApi返回
type DeleteDBInstanceResponseWrapper struct {
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

func (m *DeleteDBInstanceResponseWrapper) Reset()         { *m = DeleteDBInstanceResponseWrapper{} }
func (m *DeleteDBInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteDBInstanceResponseWrapper) ProtoMessage()    {}
func (*DeleteDBInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}
func (m *DeleteDBInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDBInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteDBInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDBInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteDBInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDBInstanceResponseWrapper.Merge(m, src)
}
func (m *DeleteDBInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteDBInstanceResponseWrapper.Size(m)
}
func (m *DeleteDBInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDBInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDBInstanceResponseWrapper proto.InternalMessageInfo

func (m *DeleteDBInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteDBInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteDBInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteDBInstanceResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteDBInstanceRequest)(nil), "dbinstance.DeleteDBInstanceRequest")
	proto.RegisterType((*DeleteDBInstanceResponseWrapper)(nil), "dbinstance.DeleteDBInstanceResponseWrapper")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcb, 0x4a, 0xeb, 0x40,
	0x18, 0xc7, 0xc9, 0x39, 0xed, 0x81, 0x33, 0x29, 0xe7, 0x94, 0x11, 0x6a, 0xa9, 0x8b, 0x94, 0x51,
	0xa4, 0x28, 0x49, 0xd4, 0x82, 0x78, 0xd9, 0x95, 0x76, 0x51, 0x10, 0x84, 0xb8, 0x70, 0x21, 0x0a,
	0x93, 0xce, 0x34, 0x06, 0x93, 0x4c, 0x9c, 0x4c, 0x5b, 0x2f, 0xf8, 0x5e, 0x3e, 0x4d, 0x04, 0xc1,
	0x17, 0xc8, 0x13, 0x48, 0xbe, 0xf4, 0x12, 0xd1, 0x55, 0xe6, 0xfb, 0xfe, 0x97, 0xfc, 0xe0, 0x43,
	0x35, 0xc6, 0x03, 0xae, 0xb8, 0x15, 0x4b, 0xa1, 0x04, 0x46, 0xcc, 0xf5, 0xa3, 0x44, 0xd1, 0x68,
	0xc4, 0x5b, 0xa6, 0xe7, 0xab, 0xdb, 0x89, 0x6b, 0x8d, 0x44, 0x68, 0x7b, 0xc2, 0x13, 0x36, 0x58,
	0xdc, 0xc9, 0x18, 0x26, 0x18, 0xe0, 0x55, 0x44, 0x5b, 0x87, 0x25, 0x7b, 0x38, 0xf3, 0xd5, 0x9d,
	0x98, 0xd9, 0x9e, 0x30, 0x41, 0x34, 0xa7, 0x34, 0xf0, 0x19, 0x55, 0x42, 0x26, 0xf6, 0xf2, 0x39,
	0xcf, 0x6d, 0x78, 0x42, 0x78, 0x01, 0x5f, 0xb5, 0xf3, 0x30, 0x56, 0x8f, 0x85, 0x48, 0x5e, 0x35,
	0xb4, 0xde, 0x07, 0xc0, 0x7e, 0x6f, 0x38, 0x07, 0x73, 0xf8, 0xfd, 0x84, 0x27, 0x0a, 0x9f, 0xa3,
	0x1a, 0x73, 0x17, 0xcb, 0x21, 0x6b, 0x6a, 0x6d, 0xad, 0xf3, 0xb7, 0xb7, 0x9b, 0xa5, 0xc6, 0xda,
	0x58, 0xc8, 0xf0, 0x84, 0x94, 0x55, 0xf2, 0xfe, 0x66, 0xd4, 0xd1, 0xbf, 0x9b, 0xab, 0x3d, 0xf3,
	0x98, 0x9a, 0x4f, 0xd7, 0xcf, 0xfb, 0xdd, 0x97, 0x2d, 0xe7, 0x4b, 0x01, 0x3e, 0x43, 0x3a, 0x73,
	0x2f, 0xb8, 0x9c, 0xfa, 0xd0, 0xf7, 0x0b, 0xfa, 0x76, 0xb2, 0xd4, 0xc0, 0x8b, 0xbe, 0xa5, 0xf8,
	0x73, 0x5d, 0x39, 0x4e, 0x3e, 0x34, 0x64, 0x7c, 0x47, 0x4f, 0x62, 0x11, 0x25, 0xfc, 0x52, 0xd2,
	0x38, 0xe6, 0x12, 0x6f, 0xa2, 0xca, 0x48, 0x30, 0x0e, 0xe8, 0xd5, 0xde, 0xff, 0x2c, 0x35, 0xf4,
	0xe2, 0x57, 0xf9, 0x96, 0x38, 0x20, 0xe2, 0x23, 0xa4, 0xe7, 0xdf, 0xc1, 0x43, 0x1c, 0x50, 0x3f,
	0x9a, 0x63, 0x35, 0x56, 0x58, 0x25, 0x91, 0x38, 0x65, 0x2b, 0xde, 0x46, 0x55, 0x2e, 0xa5, 0x90,
	0xcd, 0xdf, 0x90, 0xa9, 0x67, 0xa9, 0x51, 0x2b, 0x32, 0xb0, 0x26, 0x4e, 0x21, 0xe3, 0x53, 0x54,
	0x61, 0x54, 0xd1, 0x66, 0xa5, 0xad, 0x75, 0xf4, 0x83, 0x86, 0x55, 0x5c, 0xc4, 0x5a, 0x5c, 0xc4,
	0x1a, 0xe4, 0x17, 0x29, 0xe3, 0xe5, 0x6e, 0xe2, 0x40, 0xc8, 0xfd, 0x03, 0xb6, 0xee, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0x0f, 0x38, 0x51, 0x49, 0x02, 0x00, 0x00,
}
