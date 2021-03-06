// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_agent_plugins.proto

package agent

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
//ListAgentPlugins请求
type ListAgentPluginsRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//实例唯一标识
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAgentPluginsRequest) Reset()         { *m = ListAgentPluginsRequest{} }
func (m *ListAgentPluginsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAgentPluginsRequest) ProtoMessage()    {}
func (*ListAgentPluginsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f499f73ffa9682, []int{0}
}
func (m *ListAgentPluginsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentPluginsRequest.Unmarshal(m, b)
}
func (m *ListAgentPluginsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentPluginsRequest.Marshal(b, m, deterministic)
}
func (m *ListAgentPluginsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentPluginsRequest.Merge(m, src)
}
func (m *ListAgentPluginsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAgentPluginsRequest.Size(m)
}
func (m *ListAgentPluginsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentPluginsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentPluginsRequest proto.InternalMessageInfo

func (m *ListAgentPluginsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListAgentPluginsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAgentPluginsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//ListAgentPlugins返回
type ListAgentPluginsResponse struct {
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
	List                 []*agent_admin.PluginInstance `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ListAgentPluginsResponse) Reset()         { *m = ListAgentPluginsResponse{} }
func (m *ListAgentPluginsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAgentPluginsResponse) ProtoMessage()    {}
func (*ListAgentPluginsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f499f73ffa9682, []int{1}
}
func (m *ListAgentPluginsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentPluginsResponse.Unmarshal(m, b)
}
func (m *ListAgentPluginsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentPluginsResponse.Marshal(b, m, deterministic)
}
func (m *ListAgentPluginsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentPluginsResponse.Merge(m, src)
}
func (m *ListAgentPluginsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAgentPluginsResponse.Size(m)
}
func (m *ListAgentPluginsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentPluginsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentPluginsResponse proto.InternalMessageInfo

func (m *ListAgentPluginsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListAgentPluginsResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAgentPluginsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListAgentPluginsResponse) GetList() []*agent_admin.PluginInstance {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListAgentPluginsApi返回
type ListAgentPluginsResponseWrapper struct {
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
	Data                 *ListAgentPluginsResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListAgentPluginsResponseWrapper) Reset()         { *m = ListAgentPluginsResponseWrapper{} }
func (m *ListAgentPluginsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListAgentPluginsResponseWrapper) ProtoMessage()    {}
func (*ListAgentPluginsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f499f73ffa9682, []int{2}
}
func (m *ListAgentPluginsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentPluginsResponseWrapper.Unmarshal(m, b)
}
func (m *ListAgentPluginsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentPluginsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListAgentPluginsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentPluginsResponseWrapper.Merge(m, src)
}
func (m *ListAgentPluginsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListAgentPluginsResponseWrapper.Size(m)
}
func (m *ListAgentPluginsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentPluginsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentPluginsResponseWrapper proto.InternalMessageInfo

func (m *ListAgentPluginsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListAgentPluginsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListAgentPluginsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListAgentPluginsResponseWrapper) GetData() *ListAgentPluginsResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAgentPluginsRequest)(nil), "agent.ListAgentPluginsRequest")
	proto.RegisterType((*ListAgentPluginsResponse)(nil), "agent.ListAgentPluginsResponse")
	proto.RegisterType((*ListAgentPluginsResponseWrapper)(nil), "agent.ListAgentPluginsResponseWrapper")
}

func init() { proto.RegisterFile("list_agent_plugins.proto", fileDescriptor_a5f499f73ffa9682) }

var fileDescriptor_a5f499f73ffa9682 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0xed, 0x6c, 0x47, 0x9c, 0x8c, 0xb2, 0x43, 0x41, 0xad, 0x23, 0xd2, 0x21, 0x82, 0x8c,
	0x87, 0x69, 0x61, 0xc5, 0x3f, 0x78, 0xd2, 0xa2, 0x07, 0xc1, 0x83, 0xd4, 0x83, 0x07, 0x0f, 0x43,
	0xa6, 0xcd, 0xd6, 0x60, 0xdb, 0x37, 0x26, 0x19, 0x57, 0xfd, 0x28, 0x7e, 0xb8, 0x0a, 0x9e, 0x3c,
	0x79, 0xe8, 0x27, 0x90, 0xbc, 0xe9, 0x6e, 0x0b, 0xcb, 0x1c, 0x3c, 0xa5, 0x79, 0x9f, 0xe7, 0x49,
	0x7f, 0x0f, 0x09, 0x09, 0x2b, 0xa1, 0xcd, 0x96, 0x95, 0xbc, 0x31, 0x5b, 0x59, 0xed, 0x4b, 0xd1,
	0xe8, 0x58, 0x2a, 0x30, 0x10, 0x4c, 0x71, 0xb8, 0xdc, 0x94, 0xc2, 0x7c, 0xda, 0xef, 0xe2, 0x1c,
	0xea, 0xa4, 0x84, 0x12, 0x12, 0x54, 0x77, 0xfb, 0x53, 0xdc, 0xe1, 0x06, 0xbf, 0x5c, 0x6a, 0xf9,
	0xb1, 0x84, 0x98, 0x33, 0xfd, 0x1d, 0xa4, 0x8e, 0x2b, 0xc8, 0x59, 0x95, 0xe4, 0xd0, 0x18, 0xc5,
	0x72, 0xa3, 0x5d, 0x52, 0x71, 0x09, 0x9b, 0x1a, 0x0a, 0x5e, 0xe9, 0xa4, 0x37, 0x26, 0xb8, 0x4d,
	0x1c, 0x04, 0x2b, 0x6a, 0xd1, 0x24, 0x0e, 0x65, 0x2b, 0x1a, 0x6d, 0x58, 0x93, 0xf3, 0xfe, 0xf0,
	0x27, 0x23, 0x96, 0xfa, 0x4c, 0x98, 0xcf, 0x70, 0x96, 0x94, 0xb0, 0x41, 0x71, 0xf3, 0x95, 0x55,
	0xa2, 0x60, 0x06, 0x94, 0x4e, 0x2e, 0x3e, 0x5d, 0x8e, 0xfe, 0xf4, 0xc8, 0xed, 0xb7, 0x42, 0x9b,
	0x97, 0xf6, 0x0f, 0xef, 0x5c, 0xcb, 0x8c, 0x7f, 0xd9, 0x73, 0x6d, 0x82, 0x87, 0xc4, 0x97, 0xac,
	0xe4, 0xa1, 0xb7, 0xf2, 0xd6, 0xd3, 0xf4, 0x66, 0xd7, 0x46, 0xf3, 0x53, 0x50, 0xf5, 0x73, 0x6a,
	0xa7, 0xf4, 0xf7, 0xaf, 0x68, 0xb2, 0xb8, 0x92, 0xa1, 0x25, 0x78, 0x4c, 0xae, 0xd9, 0xf5, 0xbd,
	0xf8, 0xc1, 0xc3, 0x09, 0xda, 0xef, 0x74, 0x6d, 0x74, 0x3c, 0xd8, 0xad, 0x72, 0x1e, 0xb9, 0xb0,
	0x06, 0xf7, 0xc8, 0x44, 0x14, 0xe1, 0xd1, 0xca, 0x5b, 0xcf, 0xd2, 0x1b, 0x5d, 0x1b, 0xcd, 0x5c,
	0x40, 0x14, 0x34, 0x9b, 0x88, 0x82, 0xfe, 0xf1, 0x48, 0x78, 0x19, 0x4e, 0x4b, 0x68, 0x34, 0xff,
	0x1f, 0xba, 0xa7, 0x64, 0x66, 0xd7, 0xad, 0x1e, 0xf0, 0x96, 0x5d, 0x1b, 0x2d, 0x06, 0x3f, 0x4a,
	0x97, 0xf9, 0x1e, 0x90, 0xa9, 0x01, 0xc3, 0x2a, 0x44, 0x9c, 0xa6, 0x8b, 0xae, 0x8d, 0xae, 0xbb,
	0x10, 0x8e, 0x69, 0xe6, 0xe4, 0xe0, 0x05, 0xf1, 0xed, 0x63, 0x09, 0xfd, 0xd5, 0xd1, 0x7a, 0x7e,
	0x72, 0x37, 0x1e, 0xdd, 0x57, 0xec, 0xb8, 0xdf, 0xf4, 0xd7, 0x95, 0x1e, 0x0f, 0xa0, 0x36, 0x42,
	0x33, 0x4c, 0xd2, 0xbf, 0x1e, 0x89, 0x0e, 0x55, 0xfd, 0xa0, 0x98, 0x94, 0x5c, 0x05, 0xf7, 0x89,
	0x9f, 0x43, 0x71, 0xde, 0x78, 0x74, 0x90, 0x9d, 0xd2, 0x0c, 0xc5, 0xe0, 0x19, 0x99, 0xdb, 0xf5,
	0xf5, 0x37, 0x59, 0x31, 0xd1, 0x60, 0xdb, 0x59, 0x7a, 0xab, 0x6b, 0xa3, 0x60, 0xf0, 0xf6, 0x22,
	0xcd, 0xc6, 0x56, 0x5b, 0x96, 0x2b, 0x05, 0xaa, 0xbf, 0x8f, 0x51, 0x59, 0x1c, 0xd3, 0xcc, 0xc9,
	0xc1, 0x2b, 0xe2, 0x17, 0xcc, 0xb0, 0xd0, 0x5f, 0x79, 0xeb, 0xf9, 0x49, 0xe4, 0xca, 0xc6, 0x87,
	0xe0, 0xc7, 0x9c, 0x36, 0x46, 0x33, 0x4c, 0xef, 0xae, 0xe2, 0xfb, 0x7b, 0xf4, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0x85, 0x80, 0xb6, 0x66, 0x03, 0x00, 0x00,
}
