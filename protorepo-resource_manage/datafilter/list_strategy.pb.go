// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_strategy.proto

package datafilter

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	resource_manage "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/resource_manage"
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
//ListStrategy请求
type ListStrategyRequest struct {
	//
	//query
	Query *types.Struct `protobuf:"bytes,1,opt,name=query,proto3" json:"query" form:"query"`
	//
	//page
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//page_size
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//sort
	Sort                 *types.Struct `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort" form:"sort"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListStrategyRequest) Reset()         { *m = ListStrategyRequest{} }
func (m *ListStrategyRequest) String() string { return proto.CompactTextString(m) }
func (*ListStrategyRequest) ProtoMessage()    {}
func (*ListStrategyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f587e9ce9caed2e, []int{0}
}
func (m *ListStrategyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStrategyRequest.Unmarshal(m, b)
}
func (m *ListStrategyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStrategyRequest.Marshal(b, m, deterministic)
}
func (m *ListStrategyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStrategyRequest.Merge(m, src)
}
func (m *ListStrategyRequest) XXX_Size() int {
	return xxx_messageInfo_ListStrategyRequest.Size(m)
}
func (m *ListStrategyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStrategyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStrategyRequest proto.InternalMessageInfo

func (m *ListStrategyRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ListStrategyRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListStrategyRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListStrategyRequest) GetSort() *types.Struct {
	if m != nil {
		return m.Sort
	}
	return nil
}

//
//ListStrategy返回
type ListStrategyResponse struct {
	//
	//总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total"`
	//
	//page
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//page_size
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//数据列表
	Data                 []*ListStrategyResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListStrategyResponse) Reset()         { *m = ListStrategyResponse{} }
func (m *ListStrategyResponse) String() string { return proto.CompactTextString(m) }
func (*ListStrategyResponse) ProtoMessage()    {}
func (*ListStrategyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f587e9ce9caed2e, []int{1}
}
func (m *ListStrategyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStrategyResponse.Unmarshal(m, b)
}
func (m *ListStrategyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStrategyResponse.Marshal(b, m, deterministic)
}
func (m *ListStrategyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStrategyResponse.Merge(m, src)
}
func (m *ListStrategyResponse) XXX_Size() int {
	return xxx_messageInfo_ListStrategyResponse.Size(m)
}
func (m *ListStrategyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStrategyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListStrategyResponse proto.InternalMessageInfo

func (m *ListStrategyResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListStrategyResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListStrategyResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListStrategyResponse) GetData() []*ListStrategyResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListStrategyResponse_Data struct {
	//
	//策略详情
	Strategy *resource_manage.FilterStrategy `protobuf:"bytes,1,opt,name=strategy,proto3" json:"strategy" form:"strategy"`
	//
	//策略最后执行详情
	Instance             *resource_manage.FilterStrategyInstance `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ListStrategyResponse_Data) Reset()         { *m = ListStrategyResponse_Data{} }
func (m *ListStrategyResponse_Data) String() string { return proto.CompactTextString(m) }
func (*ListStrategyResponse_Data) ProtoMessage()    {}
func (*ListStrategyResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f587e9ce9caed2e, []int{1, 0}
}
func (m *ListStrategyResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStrategyResponse_Data.Unmarshal(m, b)
}
func (m *ListStrategyResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStrategyResponse_Data.Marshal(b, m, deterministic)
}
func (m *ListStrategyResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStrategyResponse_Data.Merge(m, src)
}
func (m *ListStrategyResponse_Data) XXX_Size() int {
	return xxx_messageInfo_ListStrategyResponse_Data.Size(m)
}
func (m *ListStrategyResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStrategyResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_ListStrategyResponse_Data proto.InternalMessageInfo

func (m *ListStrategyResponse_Data) GetStrategy() *resource_manage.FilterStrategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *ListStrategyResponse_Data) GetInstance() *resource_manage.FilterStrategyInstance {
	if m != nil {
		return m.Instance
	}
	return nil
}

//
//ListStrategyApi返回
type ListStrategyResponseWrapper struct {
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
	Data                 *ListStrategyResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListStrategyResponseWrapper) Reset()         { *m = ListStrategyResponseWrapper{} }
func (m *ListStrategyResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListStrategyResponseWrapper) ProtoMessage()    {}
func (*ListStrategyResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f587e9ce9caed2e, []int{2}
}
func (m *ListStrategyResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStrategyResponseWrapper.Unmarshal(m, b)
}
func (m *ListStrategyResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStrategyResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListStrategyResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStrategyResponseWrapper.Merge(m, src)
}
func (m *ListStrategyResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListStrategyResponseWrapper.Size(m)
}
func (m *ListStrategyResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStrategyResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListStrategyResponseWrapper proto.InternalMessageInfo

func (m *ListStrategyResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListStrategyResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListStrategyResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListStrategyResponseWrapper) GetData() *ListStrategyResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListStrategyRequest)(nil), "datafilter.ListStrategyRequest")
	proto.RegisterType((*ListStrategyResponse)(nil), "datafilter.ListStrategyResponse")
	proto.RegisterType((*ListStrategyResponse_Data)(nil), "datafilter.ListStrategyResponse.Data")
	proto.RegisterType((*ListStrategyResponseWrapper)(nil), "datafilter.ListStrategyResponseWrapper")
}

func init() { proto.RegisterFile("list_strategy.proto", fileDescriptor_3f587e9ce9caed2e) }

var fileDescriptor_3f587e9ce9caed2e = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x41, 0x8f, 0xd3, 0x3c,
	0x10, 0xfd, 0xb2, 0x9b, 0x7e, 0xda, 0xba, 0x48, 0xad, 0xdc, 0x05, 0xaa, 0x82, 0x94, 0xca, 0x08,
	0x28, 0x87, 0x3a, 0xa2, 0x48, 0x80, 0x10, 0x12, 0x52, 0xc5, 0x22, 0x81, 0x38, 0xa0, 0xf4, 0x00,
	0xb7, 0xca, 0x4d, 0xdd, 0x6c, 0x44, 0x9a, 0xc9, 0xda, 0x0e, 0x4b, 0xf7, 0x8f, 0xf0, 0x3f, 0xf8,
	0x41, 0x45, 0xe2, 0xca, 0xad, 0x27, 0x8e, 0xc8, 0x8e, 0xd3, 0x66, 0x4b, 0x61, 0x4f, 0x9c, 0x12,
	0xcf, 0xbc, 0x79, 0x33, 0x6f, 0xde, 0xa0, 0x76, 0x12, 0x4b, 0x35, 0x91, 0x4a, 0x30, 0xc5, 0xa3,
	0x25, 0xcd, 0x04, 0x28, 0xc0, 0x68, 0xc6, 0x14, 0x9b, 0xc7, 0x89, 0xe2, 0xa2, 0x3b, 0x88, 0x62,
	0x75, 0x9a, 0x4f, 0x69, 0x08, 0x0b, 0x3f, 0x82, 0x08, 0x7c, 0x03, 0x99, 0xe6, 0x73, 0xf3, 0x32,
	0x0f, 0xf3, 0x57, 0x94, 0x76, 0x27, 0x11, 0x50, 0xce, 0xe4, 0x12, 0x32, 0x49, 0x13, 0x08, 0x59,
	0xe2, 0x87, 0x90, 0x2a, 0xc1, 0x42, 0x25, 0x8b, 0x4a, 0xc1, 0x33, 0x18, 0x2c, 0x60, 0xc6, 0x13,
	0xe9, 0x5b, 0xa0, 0x6f, 0x9e, 0xbe, 0xe0, 0x12, 0x72, 0x11, 0xf2, 0xc9, 0x82, 0xa5, 0x2c, 0xe2,
	0x7e, 0x31, 0xc0, 0xce, 0x6c, 0xdd, 0xd3, 0x7f, 0xd6, 0x60, 0x12, 0xa7, 0x52, 0xb1, 0x34, 0xe4,
	0xb6, 0xd3, 0xe3, 0x8a, 0xf2, 0xc5, 0x79, 0xac, 0x3e, 0xc2, 0xb9, 0x1f, 0xc1, 0xc0, 0x24, 0x07,
	0x9f, 0x58, 0x12, 0xcf, 0x98, 0x02, 0x21, 0xfd, 0xcd, 0xaf, 0xad, 0xbb, 0x1d, 0x01, 0x44, 0x09,
	0xdf, 0x2e, 0x4a, 0x2a, 0x91, 0x87, 0xaa, 0xc8, 0x92, 0x9f, 0x0e, 0x6a, 0xbf, 0x8d, 0xa5, 0x1a,
	0xdb, 0xae, 0x01, 0x3f, 0xcb, 0xb9, 0x54, 0xf8, 0x05, 0xaa, 0x9d, 0xe5, 0x5c, 0x2c, 0x3b, 0x4e,
	0xcf, 0xe9, 0x37, 0x86, 0x37, 0x69, 0xc1, 0x42, 0x4b, 0x16, 0x3a, 0x36, 0x2c, 0xa3, 0xd6, 0x7a,
	0xe5, 0x5d, 0x9b, 0x83, 0x58, 0x3c, 0x23, 0x06, 0x4f, 0x82, 0xa2, 0x0e, 0x3f, 0x40, 0x6e, 0xc6,
	0x22, 0xde, 0x39, 0xe8, 0x39, 0xfd, 0xda, 0xe8, 0xfa, 0x7a, 0xe5, 0x35, 0x0a, 0x98, 0x8e, 0x92,
	0xef, 0xdf, 0xbc, 0x83, 0xd6, 0x7f, 0x81, 0x81, 0xe0, 0x27, 0xa8, 0xae, 0xbf, 0x13, 0x19, 0x5f,
	0xf0, 0xce, 0xa1, 0xc1, 0x77, 0xd7, 0x2b, 0xaf, 0xb5, 0xc5, 0x9b, 0x54, 0x59, 0x74, 0xa4, 0x23,
	0xe3, 0xf8, 0x82, 0xe3, 0xe7, 0xc8, 0x95, 0x20, 0x54, 0xc7, 0xfd, 0xfb, 0x8c, 0xcd, 0x6d, 0x73,
	0x0d, 0x27, 0x81, 0xa9, 0x22, 0x5f, 0x0e, 0xd1, 0xf1, 0x65, 0xe9, 0x32, 0x83, 0x54, 0x72, 0x7c,
	0x0f, 0xd5, 0x14, 0x28, 0x96, 0x18, 0xed, 0xb5, 0xaa, 0x44, 0x13, 0x26, 0x41, 0x91, 0xc6, 0x77,
	0x2e, 0x49, 0x6c, 0xee, 0x48, 0xb4, 0xe2, 0x1e, 0xfe, 0x2e, 0xee, 0x78, 0x9f, 0xb8, 0x8a, 0xac,
	0x37, 0xc8, 0xd5, 0x17, 0xdf, 0x71, 0x7b, 0x87, 0xfd, 0xc6, 0xf0, 0x2e, 0xdd, 0x9e, 0x3f, 0xdd,
	0x37, 0x2f, 0x7d, 0xc9, 0x14, 0xab, 0xb6, 0xd7, 0x78, 0x12, 0x18, 0x8e, 0xee, 0x57, 0x07, 0xb9,
	0x3a, 0x8f, 0xdf, 0xa1, 0xa3, 0xf2, 0xb2, 0xac, 0xa7, 0x1e, 0xdd, 0xb9, 0x40, 0xfa, 0xca, 0x34,
	0x29, 0xf9, 0x47, 0xed, 0xf5, 0xca, 0x6b, 0xda, 0xbd, 0xd9, 0x18, 0x09, 0x36, 0x2c, 0xf8, 0x03,
	0x3a, 0x2a, 0x4f, 0xd4, 0xac, 0xa0, 0x31, 0xbc, 0x7f, 0x05, 0xe3, 0x6b, 0x0b, 0xaf, 0x32, 0x97,
	0x14, 0x24, 0xd8, 0xb0, 0x91, 0x1f, 0x0e, 0xba, 0xb5, 0x4f, 0xe9, 0x7b, 0xc1, 0xb2, 0x8c, 0x0b,
	0xbd, 0xf8, 0x10, 0x66, 0xdc, 0xfa, 0x53, 0x51, 0xae, 0xa3, 0x24, 0x30, 0x49, 0xfc, 0x14, 0x35,
	0xf4, 0xf7, 0xe4, 0x73, 0x96, 0xb0, 0x38, 0x35, 0x13, 0xd6, 0x47, 0x37, 0xd6, 0x2b, 0x0f, 0x6f,
	0xb1, 0x36, 0x49, 0x82, 0x2a, 0x54, 0xfb, 0xcf, 0x85, 0x00, 0x61, 0xec, 0xaa, 0x57, 0xfd, 0x37,
	0x61, 0x12, 0x14, 0x69, 0x7c, 0xb2, 0xf1, 0x49, 0x8b, 0xef, 0x5d, 0xe5, 0xd3, 0x1f, 0x2c, 0x9a,
	0xfe, 0x6f, 0xee, 0xf5, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x0f, 0xa3, 0x14, 0xfc,
	0x04, 0x00, 0x00,
}
