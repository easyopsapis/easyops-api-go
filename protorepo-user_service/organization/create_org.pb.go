// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_org.proto

package organization

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//CreateOrg请求
type CreateOrgRequest struct {
	//
	//org
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//过期日期
	ExpiredDays int32 `protobuf:"varint,2,opt,name=expired_days,json=expiredDays,proto3" json:"expired_days" form:"expired_days"`
	//
	//名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo                 string   `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrgRequest) Reset()         { *m = CreateOrgRequest{} }
func (m *CreateOrgRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrgRequest) ProtoMessage()    {}
func (*CreateOrgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57863ae8e244ad8f, []int{0}
}
func (m *CreateOrgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrgRequest.Unmarshal(m, b)
}
func (m *CreateOrgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrgRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrgRequest.Merge(m, src)
}
func (m *CreateOrgRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrgRequest.Size(m)
}
func (m *CreateOrgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrgRequest proto.InternalMessageInfo

func (m *CreateOrgRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateOrgRequest) GetExpiredDays() int32 {
	if m != nil {
		return m.ExpiredDays
	}
	return 0
}

func (m *CreateOrgRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrgRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

//
//CreateOrg返回
type CreateOrgResponse struct {
	//
	//org
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//过期时间戳
	Expires int32 `protobuf:"varint,2,opt,name=expires,proto3" json:"expires" form:"expires"`
	//
	//创建日期
	CreateAt string `protobuf:"bytes,3,opt,name=createAt,proto3" json:"createAt" form:"createAt"`
	//
	//是否可用
	Valid bool `protobuf:"varint,4,opt,name=valid,proto3" json:"valid" form:"valid"`
	//
	//更新事件戳
	Ts int32 `protobuf:"varint,5,opt,name=ts,proto3" json:"ts" form:"ts"`
	//
	//名称
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo                 string   `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo" form:"memo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrgResponse) Reset()         { *m = CreateOrgResponse{} }
func (m *CreateOrgResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrgResponse) ProtoMessage()    {}
func (*CreateOrgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57863ae8e244ad8f, []int{1}
}
func (m *CreateOrgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrgResponse.Unmarshal(m, b)
}
func (m *CreateOrgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrgResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrgResponse.Merge(m, src)
}
func (m *CreateOrgResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrgResponse.Size(m)
}
func (m *CreateOrgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrgResponse proto.InternalMessageInfo

func (m *CreateOrgResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateOrgResponse) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *CreateOrgResponse) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *CreateOrgResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *CreateOrgResponse) GetTs() int32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *CreateOrgResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrgResponse) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

//
//CreateOrgApi返回
type CreateOrgResponseWrapper struct {
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
	Data                 *CreateOrgResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateOrgResponseWrapper) Reset()         { *m = CreateOrgResponseWrapper{} }
func (m *CreateOrgResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateOrgResponseWrapper) ProtoMessage()    {}
func (*CreateOrgResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_57863ae8e244ad8f, []int{2}
}
func (m *CreateOrgResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrgResponseWrapper.Unmarshal(m, b)
}
func (m *CreateOrgResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrgResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateOrgResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrgResponseWrapper.Merge(m, src)
}
func (m *CreateOrgResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateOrgResponseWrapper.Size(m)
}
func (m *CreateOrgResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrgResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrgResponseWrapper proto.InternalMessageInfo

func (m *CreateOrgResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateOrgResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateOrgResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateOrgResponseWrapper) GetData() *CreateOrgResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrgRequest)(nil), "organization.CreateOrgRequest")
	proto.RegisterType((*CreateOrgResponse)(nil), "organization.CreateOrgResponse")
	proto.RegisterType((*CreateOrgResponseWrapper)(nil), "organization.CreateOrgResponseWrapper")
}

func init() { proto.RegisterFile("create_org.proto", fileDescriptor_57863ae8e244ad8f) }

var fileDescriptor_57863ae8e244ad8f = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8a, 0x13, 0x31,
	0x18, 0xc7, 0x99, 0xb1, 0xdd, 0xdd, 0x66, 0xaa, 0x5b, 0xb3, 0xa0, 0x83, 0x20, 0x59, 0x22, 0xc8,
	0x1e, 0x74, 0x0a, 0x7a, 0x91, 0xbd, 0x39, 0xae, 0x67, 0x21, 0x17, 0x8f, 0x4b, 0xda, 0xa4, 0x63,
	0xa0, 0x33, 0x19, 0x93, 0x54, 0x5a, 0x1f, 0xc7, 0x97, 0xf0, 0x6d, 0xe6, 0xea, 0x7d, 0x9e, 0x40,
	0xf2, 0x65, 0x6a, 0x43, 0x0b, 0xea, 0x69, 0x26, 0xdf, 0xff, 0xff, 0x85, 0x1f, 0x3f, 0x82, 0x66,
	0x4b, 0x23, 0xb9, 0x93, 0xf7, 0xda, 0x54, 0x45, 0x6b, 0xb4, 0xd3, 0x78, 0xaa, 0x4d, 0xc5, 0x1b,
	0xf5, 0x9d, 0x3b, 0xa5, 0x9b, 0x67, 0xaf, 0x2b, 0xe5, 0xbe, 0x6c, 0x16, 0xc5, 0x52, 0xd7, 0xf3,
	0x4a, 0x57, 0x7a, 0x0e, 0xa5, 0xc5, 0x66, 0x05, 0x27, 0x38, 0xc0, 0x5f, 0x58, 0xa6, 0x3f, 0x13,
	0x34, 0xfb, 0x00, 0x37, 0x7e, 0x32, 0x15, 0x93, 0x5f, 0x37, 0xd2, 0x3a, 0xfc, 0x1c, 0xa5, 0x4a,
	0xe4, 0xc9, 0x75, 0x72, 0x33, 0x2e, 0x1f, 0xf6, 0x1d, 0x99, 0xac, 0xb4, 0xa9, 0x6f, 0xa9, 0x12,
	0x94, 0xa5, 0x4a, 0xe0, 0x5b, 0x34, 0x95, 0xdb, 0x56, 0x19, 0x29, 0xee, 0x05, 0xdf, 0xd9, 0x3c,
	0x85, 0xe2, 0xd3, 0xbe, 0x23, 0x57, 0xa1, 0x18, 0xa7, 0x94, 0x65, 0xc3, 0xf1, 0x8e, 0xef, 0x2c,
	0x7e, 0x81, 0x46, 0x0d, 0xaf, 0x65, 0xfe, 0xe0, 0x3a, 0xb9, 0x99, 0x94, 0x97, 0x7d, 0x47, 0xb2,
	0xb0, 0xe3, 0xa7, 0x94, 0x41, 0xe8, 0x4b, 0xb5, 0xac, 0x75, 0x3e, 0x3a, 0x2e, 0xf9, 0x29, 0x65,
	0x10, 0xd2, 0x1f, 0x29, 0x7a, 0x1c, 0x91, 0xdb, 0x56, 0x37, 0x56, 0xfe, 0x0b, 0xfd, 0x15, 0x3a,
	0x0f, 0x34, 0x7b, 0x6a, 0xdc, 0x77, 0xe4, 0x51, 0x4c, 0x6d, 0x29, 0xdb, 0x57, 0xf0, 0x1c, 0x5d,
	0x04, 0xdb, 0xef, 0xdd, 0x00, 0x7c, 0xd5, 0x77, 0xe4, 0x32, 0xd4, 0xf7, 0x09, 0x65, 0x7f, 0x4a,
	0xf8, 0x25, 0x1a, 0x7f, 0xe3, 0x6b, 0x25, 0x80, 0xfc, 0xa2, 0x9c, 0xf5, 0x1d, 0x99, 0x86, 0x36,
	0x8c, 0x29, 0x0b, 0xb1, 0xa7, 0x74, 0x36, 0x1f, 0x1f, 0x53, 0x3a, 0x4b, 0x59, 0xea, 0x0e, 0x92,
	0xce, 0xfe, 0x47, 0xd2, 0xf9, 0xdf, 0x24, 0xfd, 0x4a, 0x50, 0x7e, 0x22, 0xe9, 0xb3, 0xe1, 0x6d,
	0x2b, 0x8d, 0xbf, 0x61, 0xa9, 0x85, 0x1c, 0x6c, 0x45, 0x37, 0xf8, 0x29, 0x65, 0x10, 0xe2, 0x77,
	0x28, 0xf3, 0xdf, 0x8f, 0xdb, 0x76, 0xcd, 0x55, 0x03, 0xd6, 0x26, 0xe5, 0x93, 0xbe, 0x23, 0xf8,
	0xd0, 0x1d, 0x42, 0xca, 0xe2, 0xaa, 0x97, 0x21, 0x8d, 0xd1, 0x66, 0x50, 0x17, 0xc9, 0x80, 0x31,
	0x65, 0x21, 0xc6, 0x77, 0x68, 0x24, 0xb8, 0xe3, 0xe0, 0x2c, 0x7b, 0x43, 0x8a, 0xf8, 0x39, 0x17,
	0x27, 0xf0, 0x31, 0xa7, 0x5f, 0xa3, 0x0c, 0xb6, 0x17, 0x67, 0xf0, 0x9e, 0xdf, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x08, 0x92, 0x9b, 0x20, 0x03, 0x00, 0x00,
}