// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: set_org_expires.proto

package organization

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
//SetOrgExpiredDate请求
type SetOrgExpiredDateRequest struct {
	//
	//org
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//过期日期
	ExpiredDate          string   `protobuf:"bytes,2,opt,name=expired_date,json=expiredDate,proto3" json:"expired_date" form:"expired_date"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOrgExpiredDateRequest) Reset()         { *m = SetOrgExpiredDateRequest{} }
func (m *SetOrgExpiredDateRequest) String() string { return proto.CompactTextString(m) }
func (*SetOrgExpiredDateRequest) ProtoMessage()    {}
func (*SetOrgExpiredDateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ab01c66f95e22b6, []int{0}
}
func (m *SetOrgExpiredDateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOrgExpiredDateRequest.Unmarshal(m, b)
}
func (m *SetOrgExpiredDateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOrgExpiredDateRequest.Marshal(b, m, deterministic)
}
func (m *SetOrgExpiredDateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOrgExpiredDateRequest.Merge(m, src)
}
func (m *SetOrgExpiredDateRequest) XXX_Size() int {
	return xxx_messageInfo_SetOrgExpiredDateRequest.Size(m)
}
func (m *SetOrgExpiredDateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOrgExpiredDateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOrgExpiredDateRequest proto.InternalMessageInfo

func (m *SetOrgExpiredDateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetOrgExpiredDateRequest) GetExpiredDate() string {
	if m != nil {
		return m.ExpiredDate
	}
	return ""
}

//
//SetOrgExpiredDateApi返回
type SetOrgExpiredDateResponseWrapper struct {
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

func (m *SetOrgExpiredDateResponseWrapper) Reset()         { *m = SetOrgExpiredDateResponseWrapper{} }
func (m *SetOrgExpiredDateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SetOrgExpiredDateResponseWrapper) ProtoMessage()    {}
func (*SetOrgExpiredDateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ab01c66f95e22b6, []int{1}
}
func (m *SetOrgExpiredDateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOrgExpiredDateResponseWrapper.Unmarshal(m, b)
}
func (m *SetOrgExpiredDateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOrgExpiredDateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SetOrgExpiredDateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOrgExpiredDateResponseWrapper.Merge(m, src)
}
func (m *SetOrgExpiredDateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SetOrgExpiredDateResponseWrapper.Size(m)
}
func (m *SetOrgExpiredDateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOrgExpiredDateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SetOrgExpiredDateResponseWrapper proto.InternalMessageInfo

func (m *SetOrgExpiredDateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SetOrgExpiredDateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SetOrgExpiredDateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SetOrgExpiredDateResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SetOrgExpiredDateRequest)(nil), "organization.SetOrgExpiredDateRequest")
	proto.RegisterType((*SetOrgExpiredDateResponseWrapper)(nil), "organization.SetOrgExpiredDateResponseWrapper")
}

func init() { proto.RegisterFile("set_org_expires.proto", fileDescriptor_3ab01c66f95e22b6) }

var fileDescriptor_3ab01c66f95e22b6 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0xc0, 0x95, 0xd2, 0x21, 0xcd, 0x2d, 0x02, 0x05, 0x31, 0x45, 0x43, 0x28, 0x95, 0x41, 0xc8,
	0x39, 0x38, 0x4d, 0xda, 0x31, 0x56, 0xb8, 0x4d, 0xf4, 0x8c, 0x14, 0x0e, 0x1c, 0xa2, 0x50, 0xb9,
	0xb3, 0x67, 0x2c, 0xda, 0x7e, 0xc1, 0x71, 0xd9, 0x18, 0xe1, 0x39, 0x78, 0x18, 0xde, 0x25, 0x48,
	0x1c, 0x78, 0x80, 0x5c, 0xb8, 0xa2, 0x38, 0x65, 0xed, 0xa1, 0x27, 0xdb, 0xdf, 0xef, 0xf7, 0x59,
	0xdf, 0x1f, 0xf4, 0xa8, 0x10, 0x66, 0x06, 0x5a, 0xce, 0xc4, 0x75, 0xae, 0xb4, 0x28, 0xc2, 0x5c,
	0x83, 0x01, 0xb7, 0x0f, 0x5a, 0xb2, 0x95, 0xba, 0x61, 0x46, 0xc1, 0xea, 0x98, 0x4a, 0x65, 0x3e,
	0xae, 0xe7, 0xe1, 0x05, 0x2c, 0x87, 0x12, 0x24, 0x0c, 0xad, 0x34, 0x5f, 0x5f, 0xda, 0x97, 0x7d,
	0xd8, 0x5b, 0x9b, 0x7c, 0x7c, 0xba, 0xa3, 0x2f, 0xaf, 0x94, 0xf9, 0x04, 0x57, 0x43, 0x09, 0xd4,
	0x42, 0xfa, 0x85, 0x2d, 0x14, 0x67, 0x06, 0x74, 0x31, 0xbc, 0xbd, 0x6e, 0xf2, 0x1e, 0x4b, 0x00,
	0xb9, 0x10, 0xdb, 0xdf, 0xc5, 0x32, 0x37, 0x5f, 0x5b, 0x88, 0xff, 0x3a, 0xc8, 0x7b, 0x27, 0xcc,
	0x5b, 0x2d, 0xa7, 0xb6, 0x52, 0xfe, 0x86, 0x19, 0x91, 0x88, 0xcf, 0x6b, 0x51, 0x18, 0xf7, 0x09,
	0xea, 0x28, 0xee, 0x39, 0x03, 0x87, 0x1c, 0x9e, 0xdf, 0xab, 0x2b, 0xff, 0xf0, 0x12, 0xf4, 0xf2,
	0x15, 0x56, 0x1c, 0x27, 0x1d, 0xc5, 0xdd, 0x9f, 0x0e, 0xea, 0xb7, 0xfd, 0xf1, 0x19, 0x67, 0x46,
	0x78, 0x1d, 0x6b, 0xfe, 0x70, 0xea, 0xca, 0x7f, 0xd8, 0xaa, 0xbb, 0x18, 0xff, 0xfe, 0xe5, 0x97,
	0xe8, 0xe6, 0x43, 0x1a, 0xd1, 0x49, 0xf6, 0xed, 0xe4, 0x3b, 0x25, 0x84, 0x44, 0x69, 0x3c, 0x7e,
	0xf1, 0xf2, 0x2c, 0x2b, 0x49, 0x1c, 0x95, 0xf1, 0x28, 0x08, 0x68, 0x13, 0xa2, 0x93, 0xac, 0x4c,
	0x63, 0x3a, 0xca, 0xac, 0x5a, 0x8e, 0xd3, 0x88, 0xc6, 0x59, 0x10, 0x94, 0x24, 0x1a, 0xed, 0xe1,
	0x0d, 0x20, 0x51, 0x7a, 0x72, 0x3a, 0xc9, 0xca, 0x38, 0xde, 0xff, 0x43, 0x14, 0x04, 0xc1, 0xb3,
	0xa4, 0x27, 0xb6, 0x3d, 0xe2, 0x3f, 0x0e, 0x1a, 0xec, 0xe9, 0xbc, 0xc8, 0x61, 0x55, 0x88, 0xf7,
	0x9a, 0xe5, 0xb9, 0xd0, 0xee, 0x53, 0xd4, 0xbd, 0x00, 0x2e, 0xec, 0x0c, 0x0e, 0xce, 0xef, 0xd7,
	0x95, 0xdf, 0x6b, 0x1b, 0x6b, 0xa2, 0x38, 0xb1, 0xd0, 0x3d, 0x43, 0xbd, 0xe6, 0x9c, 0x5e, 0xe7,
	0x0b, 0xa6, 0x56, 0x9b, 0x29, 0x1c, 0xd5, 0x95, 0xef, 0x6e, 0xdd, 0x0d, 0xc4, 0xc9, 0xae, 0xea,
	0x3e, 0x47, 0x07, 0x42, 0x6b, 0xd0, 0xde, 0x1d, 0x9b, 0xf3, 0xa0, 0xae, 0xfc, 0xfe, 0x66, 0x70,
	0x4d, 0x18, 0x27, 0x2d, 0x76, 0x5f, 0xa3, 0x2e, 0x67, 0x86, 0x79, 0xdd, 0x81, 0x43, 0x7a, 0xa3,
	0xa3, 0xb0, 0xdd, 0x68, 0xf8, 0x7f, 0xa3, 0xe1, 0xb4, 0xd9, 0xe8, 0x6e, 0x79, 0x8d, 0x8d, 0x13,
	0x9b, 0x34, 0xbf, 0x6b, 0xb5, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x90, 0xf6, 0x62,
	0x94, 0x02, 0x00, 0x00,
}
