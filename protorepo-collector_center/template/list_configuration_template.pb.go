// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_configuration_template.proto

package template

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	collector_center "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/collector_center"
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
//ListConfigurationTemplate请求
type ListConfigurationTemplateRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//是否全部
	IsAll                bool     `protobuf:"varint,3,opt,name=isAll,proto3" json:"isAll" form:"isAll"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConfigurationTemplateRequest) Reset()         { *m = ListConfigurationTemplateRequest{} }
func (m *ListConfigurationTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationTemplateRequest) ProtoMessage()    {}
func (*ListConfigurationTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_470f400a0bd60e6d, []int{0}
}
func (m *ListConfigurationTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationTemplateRequest.Unmarshal(m, b)
}
func (m *ListConfigurationTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationTemplateRequest.Marshal(b, m, deterministic)
}
func (m *ListConfigurationTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationTemplateRequest.Merge(m, src)
}
func (m *ListConfigurationTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationTemplateRequest.Size(m)
}
func (m *ListConfigurationTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationTemplateRequest proto.InternalMessageInfo

func (m *ListConfigurationTemplateRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListConfigurationTemplateRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConfigurationTemplateRequest) GetIsAll() bool {
	if m != nil {
		return m.IsAll
	}
	return false
}

//
//ListConfigurationTemplate返回
type ListConfigurationTemplateResponse struct {
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
	List                 []*collector_center.ConfigurationTemplate `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ListConfigurationTemplateResponse) Reset()         { *m = ListConfigurationTemplateResponse{} }
func (m *ListConfigurationTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationTemplateResponse) ProtoMessage()    {}
func (*ListConfigurationTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_470f400a0bd60e6d, []int{1}
}
func (m *ListConfigurationTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationTemplateResponse.Unmarshal(m, b)
}
func (m *ListConfigurationTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationTemplateResponse.Marshal(b, m, deterministic)
}
func (m *ListConfigurationTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationTemplateResponse.Merge(m, src)
}
func (m *ListConfigurationTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationTemplateResponse.Size(m)
}
func (m *ListConfigurationTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationTemplateResponse proto.InternalMessageInfo

func (m *ListConfigurationTemplateResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListConfigurationTemplateResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConfigurationTemplateResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListConfigurationTemplateResponse) GetList() []*collector_center.ConfigurationTemplate {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListConfigurationTemplateApi返回
type ListConfigurationTemplateResponseWrapper struct {
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
	Data                 *ListConfigurationTemplateResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ListConfigurationTemplateResponseWrapper) Reset() {
	*m = ListConfigurationTemplateResponseWrapper{}
}
func (m *ListConfigurationTemplateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationTemplateResponseWrapper) ProtoMessage()    {}
func (*ListConfigurationTemplateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_470f400a0bd60e6d, []int{2}
}
func (m *ListConfigurationTemplateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationTemplateResponseWrapper.Unmarshal(m, b)
}
func (m *ListConfigurationTemplateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationTemplateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListConfigurationTemplateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationTemplateResponseWrapper.Merge(m, src)
}
func (m *ListConfigurationTemplateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationTemplateResponseWrapper.Size(m)
}
func (m *ListConfigurationTemplateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationTemplateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationTemplateResponseWrapper proto.InternalMessageInfo

func (m *ListConfigurationTemplateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListConfigurationTemplateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListConfigurationTemplateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListConfigurationTemplateResponseWrapper) GetData() *ListConfigurationTemplateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListConfigurationTemplateRequest)(nil), "template.ListConfigurationTemplateRequest")
	proto.RegisterType((*ListConfigurationTemplateResponse)(nil), "template.ListConfigurationTemplateResponse")
	proto.RegisterType((*ListConfigurationTemplateResponseWrapper)(nil), "template.ListConfigurationTemplateResponseWrapper")
}

func init() { proto.RegisterFile("list_configuration_template.proto", fileDescriptor_470f400a0bd60e6d) }

var fileDescriptor_470f400a0bd60e6d = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x37, 0xdd, 0x54, 0xda, 0xa9, 0xb0, 0x25, 0xa0, 0xc4, 0x5e, 0x92, 0x1d, 0x41, 0x23, 0xd2,
	0x04, 0x56, 0xfc, 0x83, 0x37, 0x2b, 0xde, 0xf6, 0x20, 0x51, 0xf0, 0x58, 0xa6, 0xe9, 0x34, 0x0e,
	0x4e, 0xf2, 0xe2, 0xcc, 0xd4, 0x55, 0xaf, 0x7e, 0x1d, 0x3f, 0x53, 0x04, 0x3f, 0x42, 0xee, 0x82,
	0xcc, 0x4b, 0xdb, 0x04, 0x59, 0xb6, 0xec, 0x69, 0xf2, 0xde, 0xef, 0x0f, 0xbf, 0xf7, 0xf2, 0xc8,
	0x99, 0x14, 0xda, 0x2c, 0x33, 0x28, 0x37, 0x22, 0xdf, 0x2a, 0x66, 0x04, 0x94, 0x4b, 0xc3, 0x8b,
	0x4a, 0x32, 0xc3, 0xe3, 0x4a, 0x81, 0x01, 0x6f, 0xb4, 0xaf, 0x67, 0xf3, 0x5c, 0x98, 0x4f, 0xdb,
	0x55, 0x9c, 0x41, 0x91, 0xe4, 0x90, 0x43, 0x82, 0x84, 0xd5, 0x76, 0x83, 0x15, 0x16, 0xf8, 0xd5,
	0x0a, 0x67, 0x79, 0x0e, 0x31, 0x67, 0xfa, 0x3b, 0x54, 0x3a, 0x96, 0x90, 0x31, 0x99, 0x64, 0x50,
	0x1a, 0xc5, 0x32, 0xa3, 0x5b, 0xa5, 0xe2, 0x15, 0xcc, 0x0b, 0x58, 0x73, 0xa9, 0x93, 0x1d, 0x31,
	0xc1, 0x32, 0xc9, 0x40, 0x4a, 0x9e, 0x19, 0x50, 0xcb, 0x8c, 0x97, 0x86, 0xab, 0xe4, 0xba, 0x84,
	0xb3, 0xe7, 0xbd, 0x5c, 0xc5, 0xa5, 0x30, 0x9f, 0xe1, 0x32, 0xc9, 0x61, 0x8e, 0xe0, 0xfc, 0x2b,
	0x93, 0x62, 0xcd, 0x0c, 0x28, 0x9d, 0x1c, 0x3e, 0x5b, 0x1d, 0xfd, 0xe5, 0x90, 0xf0, 0x42, 0x68,
	0xf3, 0xa6, 0x6f, 0xfe, 0x61, 0xe7, 0x9d, 0xf2, 0x2f, 0x5b, 0xae, 0x8d, 0xf7, 0x98, 0xb8, 0x15,
	0xcb, 0xb9, 0xef, 0x84, 0x4e, 0x34, 0x5c, 0xdc, 0x6d, 0xea, 0x60, 0xb2, 0x01, 0x55, 0xbc, 0xa2,
	0xb6, 0x4b, 0xff, 0xfc, 0x0e, 0x06, 0xd3, 0x5b, 0x29, 0x52, 0xbc, 0x67, 0x64, 0x64, 0xdf, 0xf7,
	0xe2, 0x07, 0xf7, 0x07, 0x48, 0xbf, 0xdf, 0xd4, 0xc1, 0x69, 0x47, 0xb7, 0xc8, 0x5e, 0x72, 0xa0,
	0x7a, 0x0f, 0xc9, 0x50, 0xe8, 0xd7, 0x52, 0xfa, 0x27, 0xa1, 0x13, 0x8d, 0x16, 0xd3, 0xa6, 0x0e,
	0xee, 0xb4, 0x1a, 0x6c, 0xd3, 0xb4, 0x85, 0xe9, 0x5f, 0x87, 0x9c, 0x5d, 0x13, 0x57, 0x57, 0x50,
	0x6a, 0x7e, 0x93, 0xbc, 0x2f, 0xc8, 0xd8, 0xbe, 0x4b, 0xdd, 0x05, 0x9e, 0x35, 0x75, 0x30, 0xed,
	0xf8, 0x08, 0x5d, 0x99, 0xd8, 0x80, 0x61, 0x6d, 0xe2, 0x61, 0x3f, 0x31, 0xb6, 0x69, 0xda, 0xc2,
	0xde, 0x05, 0x71, 0xed, 0x7d, 0xf9, 0x6e, 0x78, 0x12, 0x4d, 0xce, 0x1f, 0xc5, 0xff, 0xff, 0xd6,
	0xf8, 0xca, 0x51, 0x16, 0xa7, 0x5d, 0x68, 0x2b, 0xa7, 0x29, 0xba, 0xd0, 0x9f, 0x03, 0x12, 0x1d,
	0x9d, 0xff, 0xa3, 0x62, 0x55, 0xc5, 0x95, 0xf7, 0x80, 0xb8, 0x19, 0xac, 0xf7, 0x6b, 0xe8, 0x39,
	0xda, 0x2e, 0x4d, 0x11, 0xf4, 0x5e, 0x92, 0x89, 0x7d, 0xdf, 0x7e, 0xab, 0x24, 0x13, 0x25, 0xae,
	0x60, 0xbc, 0xb8, 0xd7, 0xd4, 0x81, 0xd7, 0x71, 0x77, 0x20, 0x4d, 0xfb, 0x54, 0xbb, 0x01, 0xae,
	0x14, 0x28, 0xdc, 0xc0, 0xb8, 0xbf, 0x01, 0x6c, 0xd3, 0xb4, 0x85, 0xbd, 0x77, 0xc4, 0x5d, 0x33,
	0xc3, 0x7c, 0x37, 0x74, 0xa2, 0xc9, 0xf9, 0x93, 0xf8, 0x70, 0xb9, 0x47, 0x07, 0xe9, 0x67, 0xb6,
	0x16, 0x34, 0x45, 0xa7, 0xd5, 0x6d, 0xbc, 0xdd, 0xa7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27,
	0xe4, 0x3b, 0x98, 0xba, 0x03, 0x00, 0x00,
}