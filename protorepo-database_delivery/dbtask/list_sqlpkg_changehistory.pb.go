// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_sqlpkg_changehistory.proto

package dbtask

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	database_delivery "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/database_delivery"
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
//ListSQLPackageChangeHistory请求
type ListSQLPackageChangeHistoryRequest struct {
	//
	//数据库服务ID
	PkgId string `protobuf:"bytes,1,opt,name=pkgId,proto3" json:"pkgId" form:"pkgId"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页面大小
	PageSize             int32    `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSQLPackageChangeHistoryRequest) Reset()         { *m = ListSQLPackageChangeHistoryRequest{} }
func (m *ListSQLPackageChangeHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*ListSQLPackageChangeHistoryRequest) ProtoMessage()    {}
func (*ListSQLPackageChangeHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_434bfa92ed7fcc95, []int{0}
}
func (m *ListSQLPackageChangeHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSQLPackageChangeHistoryRequest.Unmarshal(m, b)
}
func (m *ListSQLPackageChangeHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSQLPackageChangeHistoryRequest.Marshal(b, m, deterministic)
}
func (m *ListSQLPackageChangeHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSQLPackageChangeHistoryRequest.Merge(m, src)
}
func (m *ListSQLPackageChangeHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_ListSQLPackageChangeHistoryRequest.Size(m)
}
func (m *ListSQLPackageChangeHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSQLPackageChangeHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSQLPackageChangeHistoryRequest proto.InternalMessageInfo

func (m *ListSQLPackageChangeHistoryRequest) GetPkgId() string {
	if m != nil {
		return m.PkgId
	}
	return ""
}

func (m *ListSQLPackageChangeHistoryRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListSQLPackageChangeHistoryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListSQLPackageChangeHistory返回
type ListSQLPackageChangeHistoryResponse struct {
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
	List                 []*database_delivery.SQLPackageDBTaskChangeHistory `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *ListSQLPackageChangeHistoryResponse) Reset()         { *m = ListSQLPackageChangeHistoryResponse{} }
func (m *ListSQLPackageChangeHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*ListSQLPackageChangeHistoryResponse) ProtoMessage()    {}
func (*ListSQLPackageChangeHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_434bfa92ed7fcc95, []int{1}
}
func (m *ListSQLPackageChangeHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponse.Unmarshal(m, b)
}
func (m *ListSQLPackageChangeHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponse.Marshal(b, m, deterministic)
}
func (m *ListSQLPackageChangeHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSQLPackageChangeHistoryResponse.Merge(m, src)
}
func (m *ListSQLPackageChangeHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponse.Size(m)
}
func (m *ListSQLPackageChangeHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSQLPackageChangeHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSQLPackageChangeHistoryResponse proto.InternalMessageInfo

func (m *ListSQLPackageChangeHistoryResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListSQLPackageChangeHistoryResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSQLPackageChangeHistoryResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListSQLPackageChangeHistoryResponse) GetList() []*database_delivery.SQLPackageDBTaskChangeHistory {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListSQLPackageChangeHistoryApi返回
type ListSQLPackageChangeHistoryResponseWrapper struct {
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
	Data                 *ListSQLPackageChangeHistoryResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListSQLPackageChangeHistoryResponseWrapper) Reset() {
	*m = ListSQLPackageChangeHistoryResponseWrapper{}
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*ListSQLPackageChangeHistoryResponseWrapper) ProtoMessage() {}
func (*ListSQLPackageChangeHistoryResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_434bfa92ed7fcc95, []int{2}
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper.Unmarshal(m, b)
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper.Merge(m, src)
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper.Size(m)
}
func (m *ListSQLPackageChangeHistoryResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListSQLPackageChangeHistoryResponseWrapper proto.InternalMessageInfo

func (m *ListSQLPackageChangeHistoryResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListSQLPackageChangeHistoryResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListSQLPackageChangeHistoryResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListSQLPackageChangeHistoryResponseWrapper) GetData() *ListSQLPackageChangeHistoryResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListSQLPackageChangeHistoryRequest)(nil), "dbtask.ListSQLPackageChangeHistoryRequest")
	proto.RegisterType((*ListSQLPackageChangeHistoryResponse)(nil), "dbtask.ListSQLPackageChangeHistoryResponse")
	proto.RegisterType((*ListSQLPackageChangeHistoryResponseWrapper)(nil), "dbtask.ListSQLPackageChangeHistoryResponseWrapper")
}

func init() { proto.RegisterFile("list_sqlpkg_changehistory.proto", fileDescriptor_434bfa92ed7fcc95) }

var fileDescriptor_434bfa92ed7fcc95 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x25, 0x5d, 0x36, 0x51, 0x17, 0xb1, 0xca, 0x12, 0xa8, 0xf4, 0x26, 0x95, 0x37, 0xa1, 0x02,
	0x4a, 0x32, 0x36, 0xf1, 0xb7, 0xcb, 0x02, 0x12, 0x48, 0xbb, 0x18, 0x19, 0x88, 0x0b, 0x04, 0x95,
	0x9b, 0x78, 0x6e, 0x94, 0xb4, 0x9f, 0x67, 0xbb, 0x1b, 0x1d, 0xe2, 0x12, 0x89, 0xe7, 0xe2, 0x41,
	0x8a, 0xc4, 0x23, 0xf4, 0x09, 0x90, 0xed, 0xae, 0xe9, 0x34, 0x09, 0x76, 0xe5, 0xf8, 0x3b, 0xe7,
	0x7c, 0x3a, 0xe7, 0xc4, 0x28, 0x28, 0x73, 0xa5, 0xfb, 0xea, 0xa4, 0x14, 0x05, 0xef, 0xa7, 0x43,
	0x3a, 0xe6, 0x6c, 0x98, 0x2b, 0x0d, 0x72, 0x1a, 0x09, 0x09, 0x1a, 0xf0, 0x46, 0x36, 0xd0, 0x54,
	0x15, 0xed, 0x90, 0xe7, 0x7a, 0x38, 0x19, 0x44, 0x29, 0x8c, 0x62, 0x0e, 0x1c, 0x62, 0x0b, 0x0f,
	0x26, 0xc7, 0xf6, 0x66, 0x2f, 0xf6, 0xcb, 0xc9, 0xda, 0x9c, 0x43, 0xc4, 0xa8, 0x9a, 0x82, 0x50,
	0x51, 0x09, 0x29, 0x2d, 0xe3, 0x14, 0xc6, 0x5a, 0xd2, 0x54, 0x2b, 0xa7, 0x94, 0x4c, 0x40, 0x38,
	0x82, 0x8c, 0x95, 0x2a, 0x5e, 0x10, 0x63, 0x7b, 0x8d, 0x33, 0xaa, 0xe9, 0x80, 0x2a, 0xd6, 0xcf,
	0x58, 0x99, 0x9f, 0x32, 0x39, 0x8d, 0x17, 0xf6, 0x9c, 0x93, 0xfe, 0x25, 0x7f, 0xed, 0xa7, 0x2b,
	0xbe, 0x46, 0x67, 0xb9, 0x2e, 0xe0, 0x2c, 0xe6, 0x10, 0x5a, 0x30, 0x3c, 0xa5, 0x65, 0x9e, 0x51,
	0x0d, 0x52, 0xc5, 0xcb, 0x4f, 0xa7, 0x23, 0xbf, 0x3c, 0x44, 0x0e, 0x72, 0xa5, 0x8f, 0xde, 0x1d,
	0x1c, 0xd2, 0xb4, 0xa0, 0x9c, 0xbd, 0xb4, 0xe9, 0xdf, 0xb8, 0xed, 0x09, 0x3b, 0x99, 0x30, 0xa5,
	0xf1, 0x3e, 0x5a, 0x17, 0x05, 0x7f, 0x9b, 0xb5, 0xbc, 0x8e, 0xd7, 0xad, 0xf7, 0xb6, 0xe7, 0xb3,
	0xe0, 0xd6, 0x31, 0xc8, 0xd1, 0x3e, 0xb1, 0x63, 0xf2, 0xe7, 0x77, 0xd0, 0x44, 0xb7, 0xbf, 0x7c,
	0xda, 0x09, 0x5f, 0xd0, 0xf0, 0xfc, 0xf3, 0xb7, 0xc7, 0x7b, 0xdf, 0xb7, 0x13, 0x27, 0xc1, 0x0f,
	0x90, 0x2f, 0x28, 0x67, 0xad, 0x5a, 0xc7, 0xeb, 0xae, 0xf7, 0xee, 0xcc, 0x67, 0x41, 0x63, 0x21,
	0xa5, 0x9c, 0x19, 0x65, 0xad, 0x79, 0x23, 0xb1, 0x14, 0xfc, 0x04, 0xdd, 0x34, 0xe7, 0x51, 0x7e,
	0xce, 0x5a, 0x6b, 0x96, 0x7e, 0x6f, 0x3e, 0x0b, 0x36, 0x2b, 0xba, 0x41, 0x2e, 0x24, 0x4b, 0x2a,
	0xf9, 0x59, 0x43, 0x5b, 0xff, 0x0c, 0xa1, 0x04, 0x8c, 0x15, 0x5b, 0x3a, 0xf1, 0xfe, 0xef, 0xe4,
	0x19, 0xaa, 0x9b, 0xb3, 0xaf, 0x8c, 0x15, 0xe7, 0xbc, 0x3d, 0x9f, 0x05, 0xcd, 0x8a, 0x6f, 0xa1,
	0x2b, 0x5e, 0xf0, 0x7d, 0xb4, 0xae, 0x41, 0xd3, 0x72, 0xe1, 0xbf, 0x59, 0x35, 0x65, 0xc7, 0x24,
	0x71, 0x30, 0xfe, 0x80, 0x7c, 0xf3, 0xe6, 0x5a, 0x7e, 0x67, 0xad, 0xdb, 0xd8, 0xdd, 0x89, 0xae,
	0xfc, 0xee, 0xa8, 0x4a, 0xf3, 0xaa, 0xf7, 0x9e, 0xaa, 0xe2, 0x52, 0xa6, 0xde, 0x66, 0xe5, 0xde,
	0xec, 0x21, 0x89, 0x5d, 0x47, 0x7e, 0xd4, 0xd0, 0xc3, 0x6b, 0x54, 0xf1, 0x51, 0x52, 0x21, 0x98,
	0xc4, 0x5b, 0xc8, 0x4f, 0x21, 0xbb, 0x68, 0x64, 0x65, 0xa7, 0x99, 0x92, 0xc4, 0x82, 0xf8, 0x39,
	0x6a, 0x98, 0xf3, 0xf5, 0x57, 0x51, 0xd2, 0x7c, 0x6c, 0xdb, 0xa8, 0xf7, 0xee, 0xce, 0x67, 0x01,
	0xae, 0xb8, 0x0b, 0x90, 0x24, 0xab, 0x54, 0x53, 0x06, 0x93, 0x12, 0xa4, 0x2d, 0xa3, 0xbe, 0x5a,
	0x86, 0x1d, 0x93, 0xc4, 0xc1, 0xf8, 0x10, 0xf9, 0x26, 0x7f, 0xcb, 0xef, 0x78, 0xdd, 0xc6, 0xee,
	0xa3, 0xc8, 0x3d, 0xf1, 0xe8, 0x1a, 0x41, 0x56, 0x3d, 0x9b, 0x15, 0x24, 0xb1, 0x9b, 0x06, 0x1b,
	0xf6, 0x79, 0xef, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x14, 0x25, 0x68, 0x56, 0xd9, 0x03, 0x00,
	0x00,
}