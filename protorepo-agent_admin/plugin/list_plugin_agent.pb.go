// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_plugin_agent.proto

package plugin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	agent_admin "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/agent_admin"
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
//ListPluginAgents请求
type ListPluginAgentsRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//ID
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginAgentsRequest) Reset()         { *m = ListPluginAgentsRequest{} }
func (m *ListPluginAgentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPluginAgentsRequest) ProtoMessage()    {}
func (*ListPluginAgentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1db29d0e8cb1da20, []int{0}
}
func (m *ListPluginAgentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginAgentsRequest.Unmarshal(m, b)
}
func (m *ListPluginAgentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginAgentsRequest.Marshal(b, m, deterministic)
}
func (m *ListPluginAgentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginAgentsRequest.Merge(m, src)
}
func (m *ListPluginAgentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPluginAgentsRequest.Size(m)
}
func (m *ListPluginAgentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginAgentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginAgentsRequest proto.InternalMessageInfo

func (m *ListPluginAgentsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginAgentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPluginAgentsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//ListPluginAgents返回
type ListPluginAgentsResponse struct {
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
	List                 []*agent_admin.Agent `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListPluginAgentsResponse) Reset()         { *m = ListPluginAgentsResponse{} }
func (m *ListPluginAgentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPluginAgentsResponse) ProtoMessage()    {}
func (*ListPluginAgentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1db29d0e8cb1da20, []int{1}
}
func (m *ListPluginAgentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginAgentsResponse.Unmarshal(m, b)
}
func (m *ListPluginAgentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginAgentsResponse.Marshal(b, m, deterministic)
}
func (m *ListPluginAgentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginAgentsResponse.Merge(m, src)
}
func (m *ListPluginAgentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPluginAgentsResponse.Size(m)
}
func (m *ListPluginAgentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginAgentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginAgentsResponse proto.InternalMessageInfo

func (m *ListPluginAgentsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginAgentsResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPluginAgentsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListPluginAgentsResponse) GetList() []*agent_admin.Agent {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListPluginAgentsApi返回
type ListPluginAgentsResponseWrapper struct {
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
	Data                 *ListPluginAgentsResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListPluginAgentsResponseWrapper) Reset()         { *m = ListPluginAgentsResponseWrapper{} }
func (m *ListPluginAgentsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListPluginAgentsResponseWrapper) ProtoMessage()    {}
func (*ListPluginAgentsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1db29d0e8cb1da20, []int{2}
}
func (m *ListPluginAgentsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginAgentsResponseWrapper.Unmarshal(m, b)
}
func (m *ListPluginAgentsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginAgentsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListPluginAgentsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginAgentsResponseWrapper.Merge(m, src)
}
func (m *ListPluginAgentsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListPluginAgentsResponseWrapper.Size(m)
}
func (m *ListPluginAgentsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginAgentsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginAgentsResponseWrapper proto.InternalMessageInfo

func (m *ListPluginAgentsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListPluginAgentsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListPluginAgentsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListPluginAgentsResponseWrapper) GetData() *ListPluginAgentsResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPluginAgentsRequest)(nil), "plugin.ListPluginAgentsRequest")
	proto.RegisterType((*ListPluginAgentsResponse)(nil), "plugin.ListPluginAgentsResponse")
	proto.RegisterType((*ListPluginAgentsResponseWrapper)(nil), "plugin.ListPluginAgentsResponseWrapper")
}

func init() { proto.RegisterFile("list_plugin_agent.proto", fileDescriptor_1db29d0e8cb1da20) }

var fileDescriptor_1db29d0e8cb1da20 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xb5, 0xb3, 0x9d, 0xc5, 0xc9, 0x28, 0x3b, 0x04, 0x74, 0xeb, 0x80, 0xb4, 0x44, 0x90, 0xf1,
	0x61, 0x5a, 0x58, 0xd1, 0x15, 0xdf, 0x2c, 0xec, 0x9b, 0x0f, 0x5a, 0x1f, 0x7c, 0x1c, 0x32, 0x6d,
	0xb6, 0x06, 0xdb, 0xde, 0x9a, 0x64, 0x5c, 0xf5, 0x53, 0xfc, 0xb8, 0x2e, 0xf8, 0x03, 0x42, 0xbf,
	0x40, 0x72, 0x33, 0xbb, 0x2d, 0x2c, 0xf3, 0xe0, 0xd3, 0xcd, 0xbd, 0xe7, 0x9c, 0xe4, 0x1c, 0x72,
	0xc9, 0x69, 0x25, 0xb5, 0xd9, 0xb4, 0xd5, 0xae, 0x94, 0xcd, 0x86, 0x97, 0xa2, 0x31, 0x71, 0xab,
	0xc0, 0x00, 0x3d, 0x76, 0xb3, 0xe5, 0xba, 0x94, 0xe6, 0xcb, 0x6e, 0x1b, 0xe7, 0x50, 0x27, 0x25,
	0x94, 0x90, 0x20, 0xbc, 0xdd, 0x5d, 0x62, 0x87, 0x0d, 0x9e, 0x9c, 0x6c, 0xf9, 0xb1, 0x84, 0x58,
	0x70, 0xfd, 0x13, 0x5a, 0x1d, 0x57, 0x90, 0xf3, 0x2a, 0xc9, 0xa1, 0x31, 0x8a, 0xe7, 0x46, 0x3b,
	0xa5, 0x12, 0x2d, 0xac, 0x6b, 0x28, 0x44, 0xa5, 0x93, 0x3d, 0x31, 0xc1, 0x36, 0xc1, 0xd7, 0x37,
	0xbc, 0xa8, 0x65, 0x93, 0x8c, 0x9c, 0x2c, 0x5f, 0x8f, 0x1c, 0xd4, 0x57, 0xd2, 0x7c, 0x85, 0xab,
	0xa4, 0x84, 0x35, 0x82, 0xeb, 0xef, 0xbc, 0x92, 0x05, 0x37, 0xa0, 0x74, 0x72, 0x7b, 0x74, 0x3a,
	0xf6, 0xdb, 0x23, 0xa7, 0xef, 0xa5, 0x36, 0x1f, 0x30, 0xc8, 0x3b, 0x7b, 0xa3, 0xce, 0xc4, 0xb7,
	0x9d, 0xd0, 0x86, 0xbe, 0x20, 0x7e, 0xcb, 0x4b, 0x11, 0x78, 0x91, 0xb7, 0x9a, 0xa6, 0x8f, 0xfa,
	0x2e, 0x9c, 0x5f, 0x82, 0xaa, 0xdf, 0x32, 0x3b, 0x65, 0x7f, 0xae, 0xc3, 0xc9, 0xe2, 0x5e, 0x86,
	0x14, 0xfa, 0x8a, 0xdc, 0xb7, 0xf5, 0x93, 0xfc, 0x25, 0x82, 0x09, 0xd2, 0x9f, 0xf4, 0x5d, 0x78,
	0x32, 0xd0, 0x2d, 0x72, 0x23, 0xb9, 0xa5, 0xd2, 0xa7, 0x64, 0x22, 0x8b, 0xe0, 0x28, 0xf2, 0x56,
	0xb3, 0xf4, 0x61, 0xdf, 0x85, 0x33, 0x27, 0x90, 0x05, 0xcb, 0x26, 0xb2, 0x60, 0xd7, 0x1e, 0x09,
	0xee, 0x9a, 0xd3, 0x2d, 0x34, 0x5a, 0xfc, 0x8f, 0xbb, 0x73, 0x32, 0xb3, 0x75, 0xa3, 0x07, 0x7b,
	0xcb, 0xbe, 0x0b, 0x17, 0x03, 0x1f, 0xa1, 0xbb, 0xfe, 0x9e, 0x93, 0xa9, 0x01, 0xc3, 0x2b, 0xb4,
	0x38, 0x4d, 0x17, 0x7d, 0x17, 0x3e, 0x70, 0x22, 0x1c, 0xb3, 0xcc, 0xc1, 0xf4, 0x9c, 0xf8, 0x76,
	0x45, 0x02, 0x3f, 0x3a, 0x5a, 0xcd, 0xcf, 0x68, 0x3c, 0xfa, 0xa5, 0x18, 0x6d, 0xa7, 0x27, 0x83,
	0x3f, 0xcb, 0x64, 0x19, 0x0a, 0xd8, 0x5f, 0x8f, 0x84, 0x87, 0x12, 0x7e, 0x56, 0xbc, 0x6d, 0x85,
	0xa2, 0xcf, 0x88, 0x9f, 0x43, 0x71, 0x13, 0x74, 0x74, 0x91, 0x9d, 0xb2, 0x0c, 0x41, 0xfa, 0x86,
	0xcc, 0x6d, 0xbd, 0xf8, 0xd1, 0x56, 0x5c, 0x36, 0x18, 0x72, 0x96, 0x3e, 0xee, 0xbb, 0x90, 0x0e,
	0xdc, 0x3d, 0xc8, 0xb2, 0x31, 0xd5, 0x66, 0x14, 0x4a, 0x81, 0xda, 0x7f, 0xc3, 0x28, 0x23, 0x8e,
	0x59, 0xe6, 0x60, 0x7a, 0x41, 0xfc, 0x82, 0x1b, 0x1e, 0xf8, 0x91, 0xb7, 0x9a, 0x9f, 0x45, 0xb1,
	0x5b, 0xfd, 0xf8, 0x90, 0xfb, 0xb1, 0x51, 0xab, 0x63, 0x19, 0xca, 0xb7, 0xc7, 0xb8, 0x77, 0x2f,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x81, 0x81, 0x97, 0x2f, 0x54, 0x03, 0x00, 0x00,
}
