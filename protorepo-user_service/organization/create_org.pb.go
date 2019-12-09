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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

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
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//版本类型
	Edition              string   `protobuf:"bytes,5,opt,name=edition,proto3" json:"edition" form:"edition"`
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

func (m *CreateOrgRequest) GetEdition() string {
	if m != nil {
		return m.Edition
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
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xd0, 0x6e, 0xab, 0x53, 0x58, 0xf1, 0x24, 0x88, 0x90, 0x90, 0xa7, 0x87, 0x84,
	0x76, 0x80, 0x54, 0x82, 0x0b, 0xda, 0x8d, 0x30, 0xce, 0x48, 0xbe, 0x70, 0x9c, 0xdc, 0xda, 0x0b,
	0x96, 0x96, 0x38, 0x38, 0x2e, 0x5a, 0xf9, 0x38, 0x7c, 0xb0, 0xdc, 0x10, 0xf7, 0x7c, 0x02, 0xe4,
	0xe7, 0x94, 0x1a, 0x2a, 0x01, 0xa7, 0xd8, 0xef, 0xff, 0x7b, 0xf6, 0xfb, 0xff, 0x63, 0xb2, 0x58,
	0x5b, 0x25, 0x9c, 0xba, 0x36, 0xb6, 0x2a, 0x5a, 0x6b, 0x9c, 0xa1, 0x73, 0x63, 0x2b, 0xd1, 0xe8,
	0xaf, 0xc2, 0x69, 0xd3, 0x3c, 0x79, 0x59, 0x69, 0xf7, 0x69, 0xb3, 0x2a, 0xd6, 0xa6, 0x5e, 0x56,
	0xa6, 0x32, 0x4b, 0x84, 0x56, 0x9b, 0x1b, 0xdc, 0xe1, 0x06, 0x57, 0xa1, 0x19, 0xbe, 0x27, 0x64,
	0xf1, 0x0e, 0x4f, 0xfc, 0x60, 0x2b, 0xae, 0x3e, 0x6f, 0x54, 0xe7, 0xe8, 0x53, 0x92, 0x6a, 0x99,
	0x27, 0xe7, 0xc9, 0xc5, 0xb4, 0xbc, 0x3f, 0xf4, 0x6c, 0x76, 0x63, 0x6c, 0x7d, 0x09, 0x5a, 0x02,
	0x4f, 0xb5, 0xa4, 0x97, 0x64, 0xae, 0xee, 0x5a, 0x6d, 0x95, 0xbc, 0x96, 0x62, 0xdb, 0xe5, 0x29,
	0x82, 0x8f, 0x87, 0x9e, 0x9d, 0x05, 0x30, 0x56, 0x81, 0x67, 0xe3, 0xf6, 0x4a, 0x6c, 0x3b, 0xfa,
	0x8c, 0x4c, 0x1a, 0x51, 0xab, 0xfc, 0xde, 0x79, 0x72, 0x31, 0x2b, 0x4f, 0x87, 0x9e, 0x65, 0xa1,
	0xc7, 0x57, 0x81, 0xa3, 0xe8, 0xa1, 0x5a, 0xd5, 0x26, 0x9f, 0xfc, 0x09, 0xf9, 0x2a, 0x70, 0x14,
	0xe9, 0x0b, 0x72, 0xac, 0xa4, 0xf6, 0x9e, 0xf3, 0x29, 0x72, 0x74, 0xe8, 0xd9, 0x83, 0x71, 0x80,
	0x20, 0x00, 0xdf, 0x21, 0xf0, 0x2d, 0x25, 0x0f, 0x23, 0x9f, 0x5d, 0x6b, 0x9a, 0x4e, 0xfd, 0xcb,
	0xa8, 0xbf, 0x02, 0x67, 0xdf, 0x79, 0x8c, 0xaf, 0x08, 0x82, 0xbf, 0x22, 0xac, 0xe8, 0x92, 0x9c,
	0x84, 0x7f, 0xf3, 0xd6, 0x8d, 0xf6, 0xce, 0x86, 0x9e, 0x9d, 0x06, 0x7c, 0xa7, 0x00, 0xff, 0x05,
	0xd1, 0xe7, 0x64, 0xfa, 0x45, 0xdc, 0x6a, 0x89, 0x3e, 0x4f, 0xca, 0xc5, 0xd0, 0xb3, 0x79, 0xa0,
	0xb1, 0x0c, 0x3c, 0xc8, 0x7e, 0x4a, 0xd7, 0xa1, 0xc9, 0xdf, 0xa6, 0x74, 0x1d, 0xf0, 0xd4, 0xed,
	0x23, 0x3d, 0xfa, 0x9f, 0x48, 0x8f, 0xff, 0x12, 0x29, 0xfc, 0x48, 0x48, 0x7e, 0x10, 0xd2, 0x47,
	0x2b, 0xda, 0x56, 0x59, 0x7f, 0xc2, 0xda, 0x48, 0x35, 0xa6, 0x15, 0x9d, 0xe0, 0xab, 0xc0, 0x51,
	0xa4, 0x6f, 0x48, 0xe6, 0xbf, 0xef, 0xef, 0xda, 0x5b, 0xa1, 0x1b, 0x4c, 0x6d, 0x56, 0x3e, 0x1a,
	0x7a, 0x46, 0xf7, 0xec, 0x28, 0x02, 0x8f, 0x51, 0x1f, 0x86, 0xb2, 0xd6, 0xd8, 0x31, 0xba, 0x28,
	0x0c, 0x2c, 0x03, 0x0f, 0x32, 0xbd, 0x22, 0x13, 0x29, 0x9c, 0xc0, 0xcc, 0xb2, 0x57, 0xac, 0x88,
	0x1f, 0x7f, 0x71, 0x30, 0x7c, 0x3c, 0xa7, 0x6f, 0x03, 0x8e, 0xdd, 0xab, 0x23, 0x7c, 0xfd, 0xaf,
	0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xe2, 0x30, 0xec, 0x4e, 0x03, 0x00, 0x00,
}
