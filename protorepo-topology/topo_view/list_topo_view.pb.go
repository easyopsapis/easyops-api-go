// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_topo_view.proto

package topo_view

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//ListTopoView请求
type ListTopoViewRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//查询条件
	Query                *types.Struct `protobuf:"bytes,3,opt,name=query,proto3" json:"query" form:"query"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListTopoViewRequest) Reset()         { *m = ListTopoViewRequest{} }
func (m *ListTopoViewRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopoViewRequest) ProtoMessage()    {}
func (*ListTopoViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c71fad39afb4873, []int{0}
}
func (m *ListTopoViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopoViewRequest.Unmarshal(m, b)
}
func (m *ListTopoViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopoViewRequest.Marshal(b, m, deterministic)
}
func (m *ListTopoViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopoViewRequest.Merge(m, src)
}
func (m *ListTopoViewRequest) XXX_Size() int {
	return xxx_messageInfo_ListTopoViewRequest.Size(m)
}
func (m *ListTopoViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopoViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopoViewRequest proto.InternalMessageInfo

func (m *ListTopoViewRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTopoViewRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTopoViewRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

//
//ListTopoView返回
type ListTopoViewResponse struct {
	//
	//视图列表
	List []*topology.TopoView `protobuf:"bytes,1,rep,name=list,proto3" json:"list" form:"list"`
	//
	//视图总数
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total"`
	//
	//页码
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopoViewResponse) Reset()         { *m = ListTopoViewResponse{} }
func (m *ListTopoViewResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopoViewResponse) ProtoMessage()    {}
func (*ListTopoViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c71fad39afb4873, []int{1}
}
func (m *ListTopoViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopoViewResponse.Unmarshal(m, b)
}
func (m *ListTopoViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopoViewResponse.Marshal(b, m, deterministic)
}
func (m *ListTopoViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopoViewResponse.Merge(m, src)
}
func (m *ListTopoViewResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopoViewResponse.Size(m)
}
func (m *ListTopoViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopoViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopoViewResponse proto.InternalMessageInfo

func (m *ListTopoViewResponse) GetList() []*topology.TopoView {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListTopoViewResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListTopoViewResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTopoViewResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListTopoViewApi返回
type ListTopoViewResponseWrapper struct {
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
	Data                 *ListTopoViewResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListTopoViewResponseWrapper) Reset()         { *m = ListTopoViewResponseWrapper{} }
func (m *ListTopoViewResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListTopoViewResponseWrapper) ProtoMessage()    {}
func (*ListTopoViewResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c71fad39afb4873, []int{2}
}
func (m *ListTopoViewResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopoViewResponseWrapper.Unmarshal(m, b)
}
func (m *ListTopoViewResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopoViewResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListTopoViewResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopoViewResponseWrapper.Merge(m, src)
}
func (m *ListTopoViewResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListTopoViewResponseWrapper.Size(m)
}
func (m *ListTopoViewResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopoViewResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopoViewResponseWrapper proto.InternalMessageInfo

func (m *ListTopoViewResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListTopoViewResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListTopoViewResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListTopoViewResponseWrapper) GetData() *ListTopoViewResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListTopoViewRequest)(nil), "topo_view.ListTopoViewRequest")
	proto.RegisterType((*ListTopoViewResponse)(nil), "topo_view.ListTopoViewResponse")
	proto.RegisterType((*ListTopoViewResponseWrapper)(nil), "topo_view.ListTopoViewResponseWrapper")
}

func init() { proto.RegisterFile("list_topo_view.proto", fileDescriptor_0c71fad39afb4873) }

var fileDescriptor_0c71fad39afb4873 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x8b, 0x13, 0x31,
	0x18, 0x75, 0xb6, 0xad, 0xd8, 0x54, 0xb0, 0xc4, 0x55, 0x4b, 0x15, 0xa6, 0x44, 0x90, 0x7a, 0x68,
	0x06, 0x2a, 0xa8, 0x78, 0x11, 0x8a, 0xde, 0x3c, 0x65, 0x45, 0x8f, 0x25, 0x9d, 0x66, 0xc7, 0xe0,
	0xb4, 0x5f, 0x36, 0xc9, 0x6c, 0xed, 0xfe, 0x34, 0xff, 0x8a, 0x50, 0xc1, 0xa3, 0xc7, 0xfe, 0x02,
	0xc9, 0x37, 0x9d, 0xce, 0x28, 0x7b, 0x70, 0x4f, 0x49, 0xbe, 0xf7, 0xde, 0xe4, 0xbd, 0x97, 0x21,
	0xa7, 0xb9, 0x76, 0x7e, 0xee, 0xc1, 0xc0, 0xfc, 0x52, 0xab, 0x0d, 0x37, 0x16, 0x3c, 0xd0, 0xee,
	0x71, 0x30, 0x9c, 0x64, 0xda, 0x7f, 0x29, 0x16, 0x3c, 0x85, 0x55, 0x92, 0x41, 0x06, 0x09, 0x32,
	0x16, 0xc5, 0x39, 0x9e, 0xf0, 0x80, 0xbb, 0x52, 0x39, 0x14, 0x19, 0x70, 0x25, 0xdd, 0x16, 0x8c,
	0xe3, 0x39, 0xa4, 0x32, 0x4f, 0x52, 0x58, 0x7b, 0x2b, 0x53, 0xef, 0x4a, 0xa5, 0x55, 0x06, 0x26,
	0x2b, 0x58, 0xaa, 0xdc, 0x25, 0x07, 0x62, 0x82, 0xc7, 0x24, 0xdc, 0x99, 0x43, 0xb6, 0x4d, 0xfe,
	0x71, 0x33, 0x7c, 0xd9, 0xb0, 0xb0, 0xda, 0x68, 0xff, 0x15, 0x36, 0x49, 0x06, 0x13, 0x04, 0x27,
	0x97, 0x32, 0xd7, 0x4b, 0xe9, 0xc1, 0xba, 0xe4, 0xb8, 0x3d, 0xe8, 0x9e, 0x64, 0x00, 0x59, 0xae,
	0x6a, 0xc7, 0xce, 0xdb, 0x22, 0xf5, 0x25, 0xca, 0xbe, 0x47, 0xe4, 0xfe, 0x07, 0xed, 0xfc, 0x47,
	0x30, 0xf0, 0x49, 0xab, 0x8d, 0x50, 0x17, 0x85, 0x72, 0x9e, 0x3e, 0x27, 0x6d, 0x23, 0x33, 0x35,
	0x88, 0x46, 0xd1, 0xb8, 0x33, 0x7b, 0xb0, 0xdf, 0xc5, 0xbd, 0x73, 0xb0, 0xab, 0x37, 0x2c, 0x4c,
	0xd9, 0xaf, 0x9f, 0xf1, 0x49, 0xff, 0x96, 0x40, 0x0a, 0x7d, 0x45, 0xba, 0x61, 0x9d, 0x3b, 0x7d,
	0xa5, 0x06, 0x27, 0xc8, 0x1f, 0xee, 0x77, 0x71, 0xbf, 0xe6, 0x23, 0x54, 0x89, 0xee, 0x84, 0xc9,
	0x99, 0xbe, 0x52, 0xf4, 0x2d, 0xe9, 0x5c, 0x14, 0xca, 0x6e, 0x07, 0xad, 0x51, 0x34, 0xee, 0x4d,
	0x1f, 0xf1, 0xd2, 0x29, 0xaf, 0x9c, 0xf2, 0x33, 0x74, 0x3a, 0xeb, 0xef, 0x77, 0xf1, 0xdd, 0xf2,
	0x6b, 0xc8, 0x67, 0xa2, 0xd4, 0xb1, 0x1f, 0x11, 0x39, 0xfd, 0xdb, 0xbc, 0x33, 0xb0, 0x76, 0xc1,
	0x52, 0x3b, 0xbc, 0xe8, 0x20, 0x1a, 0xb5, 0xc6, 0xbd, 0x29, 0xe5, 0x55, 0xa9, 0xbc, 0x62, 0xce,
	0xee, 0xd5, 0x89, 0x02, 0x93, 0x09, 0x14, 0xd0, 0x67, 0xa4, 0xe3, 0xc1, 0xcb, 0xfc, 0x90, 0xa3,
	0x71, 0x33, 0x8e, 0x99, 0x28, 0xe1, 0x63, 0x3d, 0xad, 0x1b, 0xd6, 0xd3, 0xfe, 0xff, 0x7a, 0xd8,
	0xef, 0x88, 0x3c, 0xbe, 0x2e, 0xdd, 0x67, 0x2b, 0x8d, 0x51, 0x96, 0x3e, 0x25, 0xed, 0x14, 0x96,
	0xd5, 0x13, 0x35, 0x02, 0x85, 0x29, 0x13, 0x08, 0xd2, 0xd7, 0xa4, 0x17, 0xd6, 0xf7, 0xdf, 0x4c,
	0x2e, 0xf5, 0x1a, 0x63, 0x75, 0x67, 0x0f, 0xf7, 0xbb, 0x98, 0xd6, 0xdc, 0x03, 0xc8, 0x44, 0x93,
	0x1a, 0xaa, 0x50, 0xd6, 0x82, 0xc5, 0x8c, 0xdd, 0x66, 0x15, 0x38, 0x66, 0xa2, 0x84, 0xe9, 0x3b,
	0xd2, 0x5e, 0x4a, 0x2f, 0x31, 0x5a, 0x6f, 0x1a, 0xf3, 0xfa, 0xbf, 0xbd, 0xce, 0x7c, 0xd3, 0x67,
	0x90, 0x31, 0x81, 0xea, 0xc5, 0x6d, 0x7c, 0xf4, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0x78, 0xf9, 0x76, 0x8a, 0x03, 0x00, 0x00,
}