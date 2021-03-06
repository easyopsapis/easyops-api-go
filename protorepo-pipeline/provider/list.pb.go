// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package provider

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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
//List返回
type ListResponse struct {
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
	List                 []*pipeline.Provider `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListResponse) GetList() []*pipeline.Provider {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListApi返回
type ListResponseWrapper struct {
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
	Data                 *ListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponseWrapper) Reset()         { *m = ListResponseWrapper{} }
func (m *ListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListResponseWrapper) ProtoMessage()    {}
func (*ListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponseWrapper.Unmarshal(m, b)
}
func (m *ListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponseWrapper.Merge(m, src)
}
func (m *ListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListResponseWrapper.Size(m)
}
func (m *ListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponseWrapper proto.InternalMessageInfo

func (m *ListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListResponseWrapper) GetData() *ListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListResponse)(nil), "provider.ListResponse")
	proto.RegisterType((*ListResponseWrapper)(nil), "provider.ListResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0x96, 0xa1, 0xd5, 0x9d, 0x44, 0x65, 0xc4, 0x14, 0xf5, 0x92, 0xca, 0x48, 0xa8,
	0x1c, 0x1a, 0x4b, 0x43, 0x62, 0x08, 0x6e, 0x95, 0xb8, 0x71, 0x00, 0x73, 0xe0, 0x88, 0xdc, 0xc4,
	0x0b, 0x16, 0x6e, 0x9e, 0x65, 0x7b, 0x1b, 0xec, 0xdf, 0xe4, 0xc4, 0x25, 0x48, 0xfc, 0x09, 0xf9,
	0x0b, 0x26, 0xbf, 0x24, 0x6d, 0x4e, 0xf6, 0xfb, 0xbe, 0xef, 0xd7, 0xbe, 0x2f, 0x26, 0xc4, 0x68,
	0x1f, 0x0a, 0xeb, 0x20, 0x00, 0x3d, 0xb7, 0x0e, 0xee, 0x74, 0xa5, 0xdc, 0x72, 0x53, 0xeb, 0xf0,
	0xe3, 0x76, 0x57, 0x94, 0xb0, 0xe7, 0x35, 0xd4, 0xc0, 0x31, 0xb0, 0xbb, 0xbd, 0xc1, 0x09, 0x07,
	0xbc, 0xf5, 0xe0, 0xf2, 0x4b, 0x0d, 0x85, 0x92, 0xfe, 0x37, 0x58, 0x5f, 0x18, 0x28, 0xa5, 0xe1,
	0x25, 0x34, 0xc1, 0xc9, 0x32, 0xf8, 0x9e, 0x74, 0xca, 0xc2, 0x66, 0x0f, 0x95, 0x32, 0x9e, 0x0f,
	0x41, 0x8e, 0x23, 0xb7, 0xda, 0x2a, 0xa3, 0x1b, 0xc5, 0xc7, 0xff, 0x1e, 0x7e, 0xf2, 0xed, 0x64,
	0x83, 0xfd, 0xbd, 0x0e, 0x3f, 0xe1, 0x9e, 0xd7, 0xb0, 0x41, 0x73, 0x73, 0x27, 0x8d, 0xae, 0x64,
	0x00, 0xe7, 0xf9, 0xe1, 0xda, 0x73, 0xec, 0x4f, 0x42, 0x2e, 0x3e, 0x69, 0x1f, 0x84, 0xf2, 0x16,
	0x1a, 0xaf, 0xe8, 0x6b, 0x92, 0x5a, 0x59, 0xab, 0x2c, 0x59, 0x25, 0xeb, 0xb3, 0xed, 0x8b, 0xae,
	0xcd, 0xe7, 0x37, 0xe0, 0xf6, 0xef, 0x59, 0x54, 0xd9, 0xff, 0x7f, 0xf9, 0xc9, 0xe2, 0x89, 0xc0,
	0x08, 0xbd, 0x26, 0xb3, 0x78, 0x7e, 0xf7, 0xfa, 0x41, 0x65, 0x27, 0x98, 0x5f, 0x76, 0x6d, 0xbe,
	0x38, 0xe6, 0xd1, 0x1a, 0xa1, 0xf3, 0xa8, 0x7c, 0xd5, 0x0f, 0x8a, 0xbe, 0x22, 0x67, 0x01, 0x82,
	0x34, 0xd9, 0x29, 0x42, 0x8b, 0xae, 0xcd, 0x2f, 0x7a, 0x08, 0x65, 0x26, 0x7a, 0x9b, 0x5e, 0x93,
	0x34, 0x7e, 0xee, 0x2c, 0x5d, 0x9d, 0xae, 0xe7, 0x57, 0xb4, 0x18, 0xcb, 0x17, 0x9f, 0x87, 0xf2,
	0xdb, 0x67, 0xc7, 0xfd, 0x62, 0x92, 0x09, 0x04, 0xd8, 0xdf, 0x84, 0x3c, 0x9f, 0xb6, 0xfa, 0xe6,
	0xa4, 0xb5, 0xca, 0xd1, 0x97, 0x24, 0x2d, 0xa1, 0x1a, 0xcb, 0x4d, 0xe0, 0xa8, 0x32, 0x81, 0x26,
	0x7d, 0x47, 0xe6, 0xf1, 0xfc, 0xf8, 0xcb, 0x1a, 0xa9, 0x1b, 0x2c, 0x36, 0xdb, 0x5e, 0x76, 0x6d,
	0x4e, 0x8f, 0xd9, 0xc1, 0x64, 0x62, 0x1a, 0x8d, 0xbd, 0x94, 0x73, 0xe0, 0xb0, 0xd7, 0x6c, 0xda,
	0x0b, 0x65, 0x26, 0x7a, 0x9b, 0x7e, 0x20, 0x69, 0x25, 0x83, 0xcc, 0xd2, 0x55, 0xb2, 0x9e, 0x5f,
	0x5d, 0x16, 0x87, 0xb7, 0x9c, 0xee, 0x3c, 0x5d, 0x2f, 0xa6, 0x99, 0x40, 0x68, 0xf7, 0x14, 0x1f,
	0xee, 0xcd, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x1f, 0x98, 0x65, 0x8a, 0x02, 0x00, 0x00,
}
