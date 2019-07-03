// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_container.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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
//ListContainer请求
type ListContainerRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListContainerRequest) Reset()         { *m = ListContainerRequest{} }
func (m *ListContainerRequest) String() string { return proto.CompactTextString(m) }
func (*ListContainerRequest) ProtoMessage()    {}
func (*ListContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8293facfc6c99bf7, []int{0}
}
func (m *ListContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContainerRequest.Unmarshal(m, b)
}
func (m *ListContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContainerRequest.Marshal(b, m, deterministic)
}
func (m *ListContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContainerRequest.Merge(m, src)
}
func (m *ListContainerRequest) XXX_Size() int {
	return xxx_messageInfo_ListContainerRequest.Size(m)
}
func (m *ListContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListContainerRequest proto.InternalMessageInfo

func (m *ListContainerRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListContainerRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListContainer返回
type ListContainerResponse struct {
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
	List                 []*topology.Container `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListContainerResponse) Reset()         { *m = ListContainerResponse{} }
func (m *ListContainerResponse) String() string { return proto.CompactTextString(m) }
func (*ListContainerResponse) ProtoMessage()    {}
func (*ListContainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8293facfc6c99bf7, []int{1}
}
func (m *ListContainerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContainerResponse.Unmarshal(m, b)
}
func (m *ListContainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContainerResponse.Marshal(b, m, deterministic)
}
func (m *ListContainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContainerResponse.Merge(m, src)
}
func (m *ListContainerResponse) XXX_Size() int {
	return xxx_messageInfo_ListContainerResponse.Size(m)
}
func (m *ListContainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListContainerResponse proto.InternalMessageInfo

func (m *ListContainerResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListContainerResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListContainerResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListContainerResponse) GetList() []*topology.Container {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListContainerApi返回
type ListContainerResponseWrapper struct {
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
	Data                 *ListContainerResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListContainerResponseWrapper) Reset()         { *m = ListContainerResponseWrapper{} }
func (m *ListContainerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListContainerResponseWrapper) ProtoMessage()    {}
func (*ListContainerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8293facfc6c99bf7, []int{2}
}
func (m *ListContainerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContainerResponseWrapper.Unmarshal(m, b)
}
func (m *ListContainerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContainerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListContainerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContainerResponseWrapper.Merge(m, src)
}
func (m *ListContainerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListContainerResponseWrapper.Size(m)
}
func (m *ListContainerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContainerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListContainerResponseWrapper proto.InternalMessageInfo

func (m *ListContainerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListContainerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListContainerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListContainerResponseWrapper) GetData() *ListContainerResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListContainerRequest)(nil), "container.ListContainerRequest")
	proto.RegisterType((*ListContainerResponse)(nil), "container.ListContainerResponse")
	proto.RegisterType((*ListContainerResponseWrapper)(nil), "container.ListContainerResponseWrapper")
}

func init() { proto.RegisterFile("list_container.proto", fileDescriptor_8293facfc6c99bf7) }

var fileDescriptor_8293facfc6c99bf7 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0xdb, 0x2d, 0x22, 0x0e, 0x52, 0x23, 0xd3, 0xa2, 0x25, 0x42, 0xda, 0xc8, 0x48, 0x28,
	0x1c, 0xb2, 0x96, 0x8a, 0x80, 0x8a, 0x63, 0x50, 0x6f, 0x9c, 0xcc, 0x81, 0x63, 0xe5, 0x6c, 0xdc,
	0xc5, 0xc2, 0xc9, 0x18, 0xdb, 0xa1, 0x2d, 0xff, 0xca, 0x22, 0x71, 0xe6, 0xb4, 0x5f, 0x80, 0x3c,
	0xdb, 0x64, 0x57, 0x25, 0x17, 0x4e, 0xf6, 0xcc, 0x7b, 0x6f, 0xe6, 0x3d, 0xd9, 0xe4, 0xc4, 0x68,
	0x1f, 0x2e, 0x4b, 0x58, 0x07, 0xa9, 0xd7, 0xca, 0x15, 0xd6, 0x41, 0x00, 0x3a, 0xd8, 0x35, 0xc6,
	0xb3, 0x4a, 0x87, 0x2f, 0x9b, 0x45, 0x51, 0xc2, 0x8a, 0x57, 0x50, 0x01, 0x47, 0xc6, 0x62, 0x73,
	0x85, 0x15, 0x16, 0x78, 0x6b, 0x95, 0x63, 0x51, 0x41, 0xa1, 0xa4, 0xbf, 0x05, 0xeb, 0x0b, 0x03,
	0xa5, 0x34, 0x3c, 0x8e, 0x72, 0xb2, 0x0c, 0xbe, 0x55, 0x3a, 0x65, 0x61, 0xb6, 0x82, 0xa5, 0x32,
	0x9e, 0xdf, 0x11, 0x39, 0x96, 0x3c, 0x80, 0x05, 0x03, 0xd5, 0x2d, 0xbf, 0xe7, 0x66, 0xfc, 0xb6,
	0x67, 0x61, 0x75, 0xad, 0xc3, 0x57, 0xb8, 0xe6, 0x15, 0xcc, 0x10, 0x9c, 0x7d, 0x97, 0x46, 0x2f,
	0x65, 0x00, 0xe7, 0xf9, 0xee, 0xda, 0xea, 0xd8, 0x0d, 0x39, 0xf9, 0xa8, 0x7d, 0xf8, 0xb0, 0x1d,
	0x27, 0xd4, 0xb7, 0x8d, 0xf2, 0x81, 0xbe, 0x22, 0xa9, 0x95, 0x95, 0xca, 0x92, 0x49, 0x32, 0x3d,
	0x9a, 0x9f, 0x36, 0x75, 0x3e, 0xbc, 0x02, 0xb7, 0x7a, 0xcf, 0x62, 0x97, 0xfd, 0xfe, 0x95, 0x1f,
	0x8c, 0x1e, 0x08, 0xa4, 0xd0, 0x37, 0xe4, 0x51, 0x3c, 0x3f, 0xe9, 0x1f, 0x2a, 0x3b, 0x40, 0xfa,
	0xb3, 0xa6, 0xce, 0x8f, 0x3b, 0x7a, 0x44, 0xb6, 0x92, 0x1d, 0x95, 0xfd, 0x4c, 0xc8, 0xe9, 0xbd,
	0xd5, 0xde, 0xc2, 0xda, 0xab, 0xff, 0xd9, 0xfd, 0x8e, 0x0c, 0xe2, 0x79, 0xe9, 0xbb, 0xe5, 0xe3,
	0xa6, 0xce, 0x47, 0x1d, 0x1f, 0xa1, 0x7f, 0xb6, 0xd3, 0x97, 0xe4, 0x28, 0x40, 0x90, 0x26, 0x3b,
	0x44, 0xd1, 0xa8, 0xa9, 0xf3, 0xc7, 0xad, 0x08, 0xdb, 0x4c, 0xb4, 0x30, 0x3d, 0x27, 0x69, 0x7c,
	0xfd, 0x2c, 0x9d, 0x1c, 0x4e, 0x87, 0x67, 0x4f, 0x8a, 0xed, 0x03, 0x14, 0x3b, 0xdb, 0xf3, 0xe3,
	0xce, 0x60, 0xa4, 0x32, 0x81, 0x0a, 0xf6, 0x27, 0x21, 0xcf, 0xf7, 0xe6, 0xfb, 0xec, 0xa4, 0xb5,
	0xca, 0xd1, 0x17, 0x24, 0x2d, 0x61, 0xb9, 0x8d, 0xd9, 0x9b, 0x12, 0xbb, 0x4c, 0x20, 0x48, 0xcf,
	0xc9, 0x30, 0x9e, 0x17, 0x37, 0xd6, 0x48, 0xbd, 0xc6, 0x88, 0x83, 0xf9, 0xd3, 0xa6, 0xce, 0x69,
	0xc7, 0xbd, 0x03, 0x99, 0xe8, 0x53, 0x63, 0x42, 0xe5, 0x1c, 0x38, 0x4c, 0x38, 0xe8, 0x27, 0xc4,
	0x36, 0x13, 0x2d, 0x4c, 0x2f, 0x48, 0xba, 0x94, 0x41, 0x66, 0xe9, 0x24, 0x99, 0x0e, 0xcf, 0x26,
	0x45, 0xf7, 0xb3, 0xf6, 0xba, 0xef, 0x1b, 0x8d, 0x3a, 0x26, 0x50, 0xbe, 0x78, 0x88, 0xff, 0xe9,
	0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x1e, 0x22, 0xe8, 0x2d, 0x03, 0x00, 0x00,
}
