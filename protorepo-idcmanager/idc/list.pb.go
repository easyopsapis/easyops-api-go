// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package idc

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	idcmanager "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/idcmanager"
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
//ListIDC请求
type ListIDCRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIDCRequest) Reset()         { *m = ListIDCRequest{} }
func (m *ListIDCRequest) String() string { return proto.CompactTextString(m) }
func (*ListIDCRequest) ProtoMessage()    {}
func (*ListIDCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListIDCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIDCRequest.Unmarshal(m, b)
}
func (m *ListIDCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIDCRequest.Marshal(b, m, deterministic)
}
func (m *ListIDCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIDCRequest.Merge(m, src)
}
func (m *ListIDCRequest) XXX_Size() int {
	return xxx_messageInfo_ListIDCRequest.Size(m)
}
func (m *ListIDCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIDCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIDCRequest proto.InternalMessageInfo

func (m *ListIDCRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListIDCRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListIDC返回
type ListIDCResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*idcmanager.IDC `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListIDCResponse) Reset()         { *m = ListIDCResponse{} }
func (m *ListIDCResponse) String() string { return proto.CompactTextString(m) }
func (*ListIDCResponse) ProtoMessage()    {}
func (*ListIDCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListIDCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIDCResponse.Unmarshal(m, b)
}
func (m *ListIDCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIDCResponse.Marshal(b, m, deterministic)
}
func (m *ListIDCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIDCResponse.Merge(m, src)
}
func (m *ListIDCResponse) XXX_Size() int {
	return xxx_messageInfo_ListIDCResponse.Size(m)
}
func (m *ListIDCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIDCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIDCResponse proto.InternalMessageInfo

func (m *ListIDCResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListIDCResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListIDCResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListIDCResponse) GetList() []*idcmanager.IDC {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListIDCApi返回
type ListIDCResponseWrapper struct {
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
	Data                 *ListIDCResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListIDCResponseWrapper) Reset()         { *m = ListIDCResponseWrapper{} }
func (m *ListIDCResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListIDCResponseWrapper) ProtoMessage()    {}
func (*ListIDCResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}
func (m *ListIDCResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIDCResponseWrapper.Unmarshal(m, b)
}
func (m *ListIDCResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIDCResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListIDCResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIDCResponseWrapper.Merge(m, src)
}
func (m *ListIDCResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListIDCResponseWrapper.Size(m)
}
func (m *ListIDCResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIDCResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListIDCResponseWrapper proto.InternalMessageInfo

func (m *ListIDCResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListIDCResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListIDCResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListIDCResponseWrapper) GetData() *ListIDCResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListIDCRequest)(nil), "idc.ListIDCRequest")
	proto.RegisterType((*ListIDCResponse)(nil), "idc.ListIDCResponse")
	proto.RegisterType((*ListIDCResponseWrapper)(nil), "idc.ListIDCResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x26, 0xdb, 0x2c, 0xa2, 0x2e, 0xa2, 0x95, 0x05, 0xab, 0xa8, 0x97, 0x54, 0x46, 0x42, 0xe5,
	0xd0, 0x58, 0x5a, 0x10, 0x7f, 0xc7, 0xee, 0x72, 0x58, 0x09, 0x71, 0x30, 0x07, 0x8e, 0xc8, 0x4d,
	0xbc, 0xc1, 0x22, 0xc9, 0x18, 0xdb, 0x65, 0x81, 0xd7, 0xe4, 0xc6, 0x21, 0x48, 0x3c, 0x42, 0x9e,
	0x00, 0x79, 0xd2, 0xdd, 0x44, 0xfb, 0x00, 0x9c, 0x66, 0xe6, 0xfb, 0xc9, 0xcc, 0x97, 0x84, 0x90,
	0x4a, 0x3b, 0x9f, 0x19, 0x0b, 0x1e, 0xe8, 0x44, 0x17, 0xf9, 0x72, 0x53, 0x6a, 0xff, 0x79, 0xbf,
	0xcb, 0x72, 0xa8, 0x79, 0x09, 0x25, 0x70, 0xe4, 0x76, 0xfb, 0x4b, 0x9c, 0x70, 0xc0, 0xae, 0xf7,
	0x2c, 0xdf, 0x97, 0x90, 0x29, 0xe9, 0x7e, 0x80, 0x71, 0x59, 0x05, 0xb9, 0xac, 0x78, 0x0e, 0x8d,
	0xb7, 0x32, 0xf7, 0xae, 0x77, 0x5a, 0x65, 0x60, 0x53, 0x43, 0xa1, 0x2a, 0xc7, 0x0f, 0x42, 0x8e,
	0x23, 0xd7, 0x45, 0x5e, 0xcb, 0x46, 0x96, 0xca, 0x86, 0xf6, 0xf0, 0xbc, 0x17, 0xa3, 0xf5, 0xf5,
	0x95, 0xf6, 0x5f, 0xe0, 0x8a, 0x97, 0xb0, 0x41, 0x72, 0xf3, 0x4d, 0x56, 0xba, 0x90, 0x1e, 0xac,
	0xe3, 0x37, 0x6d, 0xef, 0x63, 0x9e, 0x3c, 0x78, 0xa7, 0x9d, 0xbf, 0x38, 0x3f, 0x13, 0xea, 0xeb,
	0x5e, 0x39, 0x4f, 0x9f, 0x92, 0xd8, 0xc8, 0x52, 0x25, 0xd1, 0x2a, 0x5a, 0x1f, 0x6f, 0x1f, 0x75,
	0x6d, 0x3a, 0xbb, 0x04, 0x5b, 0xbf, 0x61, 0x01, 0x65, 0x7f, 0xff, 0xa4, 0x47, 0x8b, 0x3b, 0x02,
	0x25, 0xf4, 0x25, 0x99, 0x86, 0xfa, 0xc9, 0xe9, 0x9f, 0x2a, 0x39, 0x42, 0xfd, 0xb2, 0x6b, 0xd3,
	0xc5, 0xa0, 0x47, 0xea, 0xda, 0x74, 0x2f, 0x20, 0x1f, 0x02, 0xf0, 0x2b, 0x22, 0xf3, 0x9b, 0xb5,
	0xce, 0x40, 0xe3, 0xd4, 0xff, 0xd8, 0x4b, 0x9f, 0x90, 0x63, 0x0f, 0x5e, 0x56, 0xc9, 0x04, 0x4d,
	0x8b, 0xae, 0x4d, 0xef, 0xf7, 0x26, 0x84, 0x99, 0xe8, 0x69, 0xfa, 0x9c, 0xc4, 0xe1, 0xfb, 0x26,
	0xf1, 0x6a, 0xb2, 0x9e, 0x9d, 0xce, 0xb3, 0xe1, 0x95, 0x67, 0x17, 0xe7, 0x67, 0xdb, 0xf9, 0x70,
	0x5c, 0x90, 0x31, 0x81, 0x6a, 0xf6, 0x3b, 0x22, 0x27, 0xb7, 0x52, 0x7d, 0xb4, 0xd2, 0x18, 0x65,
	0xe9, 0x63, 0x12, 0xe7, 0x50, 0x5c, 0x87, 0x1b, 0xf9, 0x03, 0xca, 0x04, 0x92, 0xf4, 0x15, 0x99,
	0x85, 0xfa, 0xf6, 0xbb, 0xa9, 0xa4, 0x6e, 0x30, 0xd8, 0x74, 0x7b, 0xd2, 0xb5, 0x29, 0x1d, 0xb4,
	0x07, 0x92, 0x89, 0xb1, 0x34, 0xe4, 0x52, 0xd6, 0x82, 0xc5, 0x5c, 0xd3, 0x71, 0x2e, 0x84, 0x99,
	0xe8, 0x69, 0xfa, 0x9a, 0xc4, 0x85, 0xf4, 0x32, 0x89, 0x57, 0xd1, 0x7a, 0x76, 0xfa, 0x30, 0xe4,
	0xca, 0x6e, 0x5d, 0x3c, 0x3e, 0x2e, 0x68, 0x99, 0x40, 0xcb, 0xee, 0x2e, 0xfe, 0x2f, 0xcf, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xdf, 0xe4, 0xcc, 0xf9, 0x02, 0x00, 0x00,
}
