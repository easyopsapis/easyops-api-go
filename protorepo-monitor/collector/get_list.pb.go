// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_list.proto

package collector

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//GetList请求
type GetListRequest struct {
	//
	//指定查询的表
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table" form:"table"`
	//
	//指定查询的指标名
	XSelect__ []string `protobuf:"bytes,2,rep,name=__select__,json=Select,proto3" json:"__select__" form:"__select__"`
	//
	//指定分组字段
	XGroupby__ string `protobuf:"bytes,3,opt,name=__groupby__,json=Groupby,proto3" json:"__groupby__" form:"__groupby__"`
	//
	//开始时间
	St string `protobuf:"bytes,4,opt,name=st,proto3" json:"st" form:"st"`
	//
	//结束时间
	Et string `protobuf:"bytes,5,opt,name=et,proto3" json:"et" form:"et"`
	//
	//时间范围
	Timerange string `protobuf:"bytes,6,opt,name=timerange,proto3" json:"timerange" form:"timerange"`
	//
	//页码
	Page int32 `protobuf:"varint,7,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize             int32    `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ee61abb72e5a22, []int{0}
}
func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *GetListRequest) GetXSelect__() []string {
	if m != nil {
		return m.XSelect__
	}
	return nil
}

func (m *GetListRequest) GetXGroupby__() string {
	if m != nil {
		return m.XGroupby__
	}
	return ""
}

func (m *GetListRequest) GetSt() string {
	if m != nil {
		return m.St
	}
	return ""
}

func (m *GetListRequest) GetEt() string {
	if m != nil {
		return m.Et
	}
	return ""
}

func (m *GetListRequest) GetTimerange() string {
	if m != nil {
		return m.Timerange
	}
	return ""
}

func (m *GetListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//GetList返回
type GetListResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//数据
	Data []*types.Struct `protobuf:"bytes,3,rep,name=data,proto3" json:"data" form:"data"`
	//
	//页码
	Page int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total                int32    `protobuf:"varint,6,opt,name=total,proto3" json:"total" form:"total"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListResponse) Reset()         { *m = GetListResponse{} }
func (m *GetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetListResponse) ProtoMessage()    {}
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ee61abb72e5a22, []int{1}
}
func (m *GetListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResponse.Unmarshal(m, b)
}
func (m *GetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResponse.Marshal(b, m, deterministic)
}
func (m *GetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResponse.Merge(m, src)
}
func (m *GetListResponse) XXX_Size() int {
	return xxx_messageInfo_GetListResponse.Size(m)
}
func (m *GetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResponse proto.InternalMessageInfo

func (m *GetListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetListResponse) GetData() []*types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

//
//GetListApi返回
type GetListResponseWrapper struct {
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
	Data                 *GetListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetListResponseWrapper) Reset()         { *m = GetListResponseWrapper{} }
func (m *GetListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetListResponseWrapper) ProtoMessage()    {}
func (*GetListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ee61abb72e5a22, []int{2}
}
func (m *GetListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResponseWrapper.Unmarshal(m, b)
}
func (m *GetListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResponseWrapper.Merge(m, src)
}
func (m *GetListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetListResponseWrapper.Size(m)
}
func (m *GetListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResponseWrapper proto.InternalMessageInfo

func (m *GetListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetListResponseWrapper) GetData() *GetListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetListRequest)(nil), "collector.GetListRequest")
	proto.RegisterType((*GetListResponse)(nil), "collector.GetListResponse")
	proto.RegisterType((*GetListResponseWrapper)(nil), "collector.GetListResponseWrapper")
}

func init() { proto.RegisterFile("get_list.proto", fileDescriptor_62ee61abb72e5a22) }

var fileDescriptor_62ee61abb72e5a22 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xb6, 0x69, 0xd3, 0xdd, 0x4c, 0xb5, 0x5b, 0x07, 0x5d, 0x43, 0x51, 0x52, 0x46, 0x90, 0x7a,
	0x68, 0x8a, 0x55, 0x54, 0x44, 0x10, 0x0a, 0xb2, 0x17, 0x4f, 0xd3, 0x83, 0xc7, 0x90, 0xb6, 0xb3,
	0x63, 0x30, 0xe9, 0xc4, 0x99, 0xa9, 0xab, 0xfe, 0x1c, 0x7f, 0x58, 0x14, 0x7f, 0x42, 0x2e, 0x5e,
	0x65, 0xde, 0xa4, 0x49, 0x2c, 0x08, 0xeb, 0x69, 0x5e, 0xde, 0xf7, 0xbd, 0xc9, 0xfb, 0xbe, 0x7c,
	0x41, 0x43, 0xce, 0x74, 0x94, 0x26, 0x4a, 0x87, 0xb9, 0x14, 0x5a, 0x60, 0x6f, 0x23, 0xd2, 0x94,
	0x6d, 0xb4, 0x90, 0xe3, 0x19, 0x4f, 0xf4, 0x87, 0xfd, 0x3a, 0xdc, 0x88, 0x6c, 0xce, 0x05, 0x17,
	0x73, 0x60, 0xac, 0xf7, 0x97, 0xf0, 0x04, 0x0f, 0x50, 0xd9, 0xc9, 0xf1, 0xf3, 0x16, 0x3d, 0xbb,
	0x4a, 0xf4, 0x47, 0x71, 0x35, 0xe7, 0x62, 0x06, 0xe0, 0xec, 0x73, 0x9c, 0x26, 0xdb, 0x58, 0x0b,
	0xa9, 0xe6, 0x75, 0x59, 0xcd, 0xdd, 0xe7, 0x42, 0xf0, 0x94, 0x35, 0xb7, 0x2b, 0x2d, 0xf7, 0x9b,
	0x6a, 0x1f, 0xf2, 0xdb, 0x41, 0xc3, 0x0b, 0xa6, 0xdf, 0x25, 0x4a, 0x53, 0xf6, 0x69, 0xcf, 0x94,
	0xc6, 0x8f, 0x90, 0xab, 0xe3, 0x75, 0xca, 0xfc, 0xce, 0xa4, 0x33, 0xf5, 0x96, 0xa3, 0xb2, 0x08,
	0x6e, 0x5e, 0x0a, 0x99, 0xbd, 0x22, 0xd0, 0x26, 0xd4, 0xc2, 0xf8, 0x09, 0x42, 0x51, 0xa4, 0x98,
	0x51, 0x13, 0x45, 0xbe, 0x33, 0xe9, 0x4e, 0xbd, 0xe5, 0xdd, 0xb2, 0x08, 0x6e, 0x5b, 0x72, 0x83,
	0x11, 0xda, 0x5f, 0x41, 0x89, 0x9f, 0xa1, 0x41, 0x14, 0x71, 0x29, 0xf6, 0xf9, 0xfa, 0x6b, 0x14,
	0xf9, 0x5d, 0x78, 0xc1, 0x79, 0x59, 0x04, 0xf8, 0x30, 0x53, 0x83, 0x84, 0x9e, 0x5c, 0xd8, 0x1a,
	0x3f, 0x40, 0x8e, 0xd2, 0x7e, 0x0f, 0xc8, 0xb7, 0xca, 0x22, 0xf0, 0x2c, 0x59, 0x69, 0x42, 0x1d,
	0xa5, 0x0d, 0xcc, 0xb4, 0xef, 0x1e, 0xc3, 0xcc, 0xc0, 0x4c, 0xe3, 0x05, 0xf2, 0x74, 0x92, 0x31,
	0x19, 0xef, 0x38, 0xf3, 0xfb, 0xc0, 0xba, 0x53, 0x16, 0xc1, 0xa8, 0x92, 0x74, 0x80, 0x08, 0x6d,
	0x68, 0xf8, 0x31, 0xea, 0xe5, 0x31, 0x67, 0xfe, 0xc9, 0xa4, 0x33, 0x75, 0x41, 0xd4, 0xc0, 0xd2,
	0x4d, 0x97, 0xfc, 0xfa, 0x11, 0x38, 0xa3, 0x1b, 0x14, 0x28, 0xf8, 0x05, 0xf2, 0xcc, 0x19, 0xa9,
	0xe4, 0x1b, 0xf3, 0x4f, 0x81, 0x3f, 0x6e, 0xae, 0xaf, 0xa1, 0xc3, 0xd0, 0xa9, 0xe9, 0xac, 0x4c,
	0xe3, 0xbb, 0x83, 0xce, 0x6a, 0xe7, 0x55, 0x2e, 0x76, 0x8a, 0xe1, 0x87, 0xa8, 0xb7, 0x11, 0x5b,
	0xeb, 0xbc, 0xbb, 0x3c, 0x6b, 0xde, 0x6b, 0xba, 0x84, 0x02, 0x88, 0x27, 0xa8, 0x9b, 0x29, 0xee,
	0x3b, 0x20, 0x65, 0x58, 0x16, 0x01, 0xb2, 0x9c, 0x4c, 0x71, 0x42, 0x0d, 0x84, 0x5f, 0xa3, 0xde,
	0x36, 0xd6, 0xb1, 0xdf, 0x9d, 0x74, 0xa7, 0x83, 0xc5, 0xbd, 0xd0, 0x26, 0x20, 0x3c, 0x24, 0x20,
	0x5c, 0x41, 0x02, 0xda, 0xf7, 0x1b, 0x3a, 0xa1, 0x30, 0x55, 0x8b, 0xef, 0xfd, 0xa7, 0x78, 0xf7,
	0xfa, 0xe2, 0x21, 0x63, 0x42, 0xc7, 0x29, 0x7c, 0x10, 0xf7, 0xaf, 0x8c, 0x99, 0xb6, 0xc9, 0x18,
	0x9c, 0x3f, 0x3b, 0xe8, 0xfc, 0xc8, 0xa4, 0xf7, 0x32, 0xce, 0x73, 0x26, 0xaf, 0xe7, 0xd5, 0x4b,
	0x34, 0x30, 0xe7, 0xdb, 0x2f, 0x79, 0x1a, 0x27, 0xbb, 0xca, 0xb3, 0x56, 0xe0, 0x5a, 0x20, 0xa1,
	0x6d, 0xaa, 0xd9, 0x90, 0x49, 0x29, 0x64, 0x15, 0xd2, 0xd6, 0x86, 0xd0, 0x26, 0xd4, 0xc2, 0xf8,
	0x4d, 0xe5, 0xb5, 0x71, 0x6b, 0xb0, 0x18, 0x87, 0xf5, 0xff, 0x1d, 0x1e, 0xed, 0xfd, 0x0f, 0xbb,
	0xd7, 0x7d, 0xf8, 0x2c, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x87, 0x79, 0xe4, 0x2a,
	0x04, 0x00, 0x00,
}