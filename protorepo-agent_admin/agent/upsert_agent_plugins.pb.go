// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upsert_agent_plugins.proto

package agent

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//UpsertAgentPlugins请求
type UpsertAgentPluginsRequest struct {
	//
	//数据上报的信息
	Vals *UpsertAgentPluginsRequest_Vals `protobuf:"bytes,1,opt,name=vals,proto3" json:"vals" form:"vals"`
	//
	//数据上报维度
	Dims                 *types.Struct `protobuf:"bytes,2,opt,name=dims,proto3" json:"dims" form:"dims"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpsertAgentPluginsRequest) Reset()         { *m = UpsertAgentPluginsRequest{} }
func (m *UpsertAgentPluginsRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertAgentPluginsRequest) ProtoMessage()    {}
func (*UpsertAgentPluginsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ddfd6abcf203d6, []int{0}
}
func (m *UpsertAgentPluginsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAgentPluginsRequest.Unmarshal(m, b)
}
func (m *UpsertAgentPluginsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAgentPluginsRequest.Marshal(b, m, deterministic)
}
func (m *UpsertAgentPluginsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAgentPluginsRequest.Merge(m, src)
}
func (m *UpsertAgentPluginsRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertAgentPluginsRequest.Size(m)
}
func (m *UpsertAgentPluginsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAgentPluginsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAgentPluginsRequest proto.InternalMessageInfo

func (m *UpsertAgentPluginsRequest) GetVals() *UpsertAgentPluginsRequest_Vals {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *UpsertAgentPluginsRequest) GetDims() *types.Struct {
	if m != nil {
		return m.Dims
	}
	return nil
}

type UpsertAgentPluginsRequest_Vals struct {
	//
	//数据上报的信息
	Plugins              []*UpsertAgentPluginsRequest_Vals_Plugins `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins" form:"plugins"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *UpsertAgentPluginsRequest_Vals) Reset()         { *m = UpsertAgentPluginsRequest_Vals{} }
func (m *UpsertAgentPluginsRequest_Vals) String() string { return proto.CompactTextString(m) }
func (*UpsertAgentPluginsRequest_Vals) ProtoMessage()    {}
func (*UpsertAgentPluginsRequest_Vals) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ddfd6abcf203d6, []int{0, 0}
}
func (m *UpsertAgentPluginsRequest_Vals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals.Unmarshal(m, b)
}
func (m *UpsertAgentPluginsRequest_Vals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals.Marshal(b, m, deterministic)
}
func (m *UpsertAgentPluginsRequest_Vals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAgentPluginsRequest_Vals.Merge(m, src)
}
func (m *UpsertAgentPluginsRequest_Vals) XXX_Size() int {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals.Size(m)
}
func (m *UpsertAgentPluginsRequest_Vals) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAgentPluginsRequest_Vals.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAgentPluginsRequest_Vals proto.InternalMessageInfo

func (m *UpsertAgentPluginsRequest_Vals) GetPlugins() []*UpsertAgentPluginsRequest_Vals_Plugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

type UpsertAgentPluginsRequest_Vals_Plugins struct {
	//
	//实际部署路径
	DeployPath string `protobuf:"bytes,1,opt,name=deployPath,proto3" json:"deployPath" form:"deployPath"`
	//
	//AgentUUID
	AgentId string `protobuf:"bytes,2,opt,name=agentId,proto3" json:"agentId" form:"agentId"`
	//
	//部署的插件版本名称
	AgentPluginId string `protobuf:"bytes,3,opt,name=agentPluginId,proto3" json:"agentPluginId" form:"agentPluginId"`
	//
	//插件的版本ID
	VersionId string `protobuf:"bytes,4,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//AgentIp
	Ip                   string   `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip" form:"ip"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertAgentPluginsRequest_Vals_Plugins) Reset() {
	*m = UpsertAgentPluginsRequest_Vals_Plugins{}
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) String() string { return proto.CompactTextString(m) }
func (*UpsertAgentPluginsRequest_Vals_Plugins) ProtoMessage()    {}
func (*UpsertAgentPluginsRequest_Vals_Plugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ddfd6abcf203d6, []int{0, 0, 0}
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins.Unmarshal(m, b)
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins.Marshal(b, m, deterministic)
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins.Merge(m, src)
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) XXX_Size() int {
	return xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins.Size(m)
}
func (m *UpsertAgentPluginsRequest_Vals_Plugins) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAgentPluginsRequest_Vals_Plugins proto.InternalMessageInfo

func (m *UpsertAgentPluginsRequest_Vals_Plugins) GetDeployPath() string {
	if m != nil {
		return m.DeployPath
	}
	return ""
}

func (m *UpsertAgentPluginsRequest_Vals_Plugins) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *UpsertAgentPluginsRequest_Vals_Plugins) GetAgentPluginId() string {
	if m != nil {
		return m.AgentPluginId
	}
	return ""
}

func (m *UpsertAgentPluginsRequest_Vals_Plugins) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *UpsertAgentPluginsRequest_Vals_Plugins) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

//
//UpsertAgentPlugins返回
type UpsertAgentPluginsResponse struct {
	//
	//信息
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message" form:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertAgentPluginsResponse) Reset()         { *m = UpsertAgentPluginsResponse{} }
func (m *UpsertAgentPluginsResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertAgentPluginsResponse) ProtoMessage()    {}
func (*UpsertAgentPluginsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ddfd6abcf203d6, []int{1}
}
func (m *UpsertAgentPluginsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAgentPluginsResponse.Unmarshal(m, b)
}
func (m *UpsertAgentPluginsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAgentPluginsResponse.Marshal(b, m, deterministic)
}
func (m *UpsertAgentPluginsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAgentPluginsResponse.Merge(m, src)
}
func (m *UpsertAgentPluginsResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertAgentPluginsResponse.Size(m)
}
func (m *UpsertAgentPluginsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAgentPluginsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAgentPluginsResponse proto.InternalMessageInfo

func (m *UpsertAgentPluginsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//
//UpsertAgentPluginsApi返回
type UpsertAgentPluginsResponseWrapper struct {
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
	Data                 *UpsertAgentPluginsResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UpsertAgentPluginsResponseWrapper) Reset()         { *m = UpsertAgentPluginsResponseWrapper{} }
func (m *UpsertAgentPluginsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpsertAgentPluginsResponseWrapper) ProtoMessage()    {}
func (*UpsertAgentPluginsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ddfd6abcf203d6, []int{2}
}
func (m *UpsertAgentPluginsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertAgentPluginsResponseWrapper.Unmarshal(m, b)
}
func (m *UpsertAgentPluginsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertAgentPluginsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpsertAgentPluginsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertAgentPluginsResponseWrapper.Merge(m, src)
}
func (m *UpsertAgentPluginsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpsertAgentPluginsResponseWrapper.Size(m)
}
func (m *UpsertAgentPluginsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertAgentPluginsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertAgentPluginsResponseWrapper proto.InternalMessageInfo

func (m *UpsertAgentPluginsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpsertAgentPluginsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpsertAgentPluginsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpsertAgentPluginsResponseWrapper) GetData() *UpsertAgentPluginsResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpsertAgentPluginsRequest)(nil), "agent.UpsertAgentPluginsRequest")
	proto.RegisterType((*UpsertAgentPluginsRequest_Vals)(nil), "agent.UpsertAgentPluginsRequest.Vals")
	proto.RegisterType((*UpsertAgentPluginsRequest_Vals_Plugins)(nil), "agent.UpsertAgentPluginsRequest.Vals.Plugins")
	proto.RegisterType((*UpsertAgentPluginsResponse)(nil), "agent.UpsertAgentPluginsResponse")
	proto.RegisterType((*UpsertAgentPluginsResponseWrapper)(nil), "agent.UpsertAgentPluginsResponseWrapper")
}

func init() { proto.RegisterFile("upsert_agent_plugins.proto", fileDescriptor_79ddfd6abcf203d6) }

var fileDescriptor_79ddfd6abcf203d6 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x67, 0xd3, 0xac, 0xcb, 0xbe, 0x58, 0xad, 0x43, 0xd5, 0x18, 0x94, 0xb4, 0x23, 0x4a, 0x0f,
	0x36, 0x85, 0x15, 0x41, 0x44, 0x04, 0x17, 0x14, 0xda, 0x53, 0x19, 0x51, 0x0f, 0x1e, 0xca, 0xec,
	0x66, 0x9a, 0x06, 0xb2, 0x3b, 0xe3, 0xcc, 0xa4, 0xe8, 0xd1, 0x4f, 0xe0, 0x37, 0xcc, 0xd5, 0x7b,
	0x3e, 0x81, 0xe4, 0xcd, 0xc4, 0x4d, 0xc1, 0x8a, 0xa7, 0x64, 0xde, 0xef, 0xcf, 0xfc, 0xe6, 0xc7,
	0x83, 0xa4, 0x56, 0x46, 0x68, 0x7b, 0xc6, 0x0b, 0xb1, 0xb6, 0x67, 0xaa, 0xaa, 0x8b, 0x72, 0x6d,
	0x32, 0xa5, 0xa5, 0x95, 0x64, 0x8c, 0xc3, 0xe4, 0xb0, 0x28, 0xed, 0x45, 0xbd, 0xc8, 0x96, 0x72,
	0x75, 0x54, 0xc8, 0x42, 0x1e, 0x21, 0xba, 0xa8, 0xcf, 0xf1, 0x84, 0x07, 0xfc, 0x73, 0xaa, 0xe4,
	0x61, 0x21, 0x65, 0x51, 0x89, 0x0d, 0xcb, 0x58, 0x5d, 0x2f, 0xad, 0x43, 0xe9, 0xcf, 0x10, 0x1e,
	0x7c, 0xc4, 0x2b, 0xdf, 0x76, 0xe6, 0xa7, 0xee, 0x42, 0x26, 0xbe, 0xd6, 0xc2, 0x58, 0x72, 0x02,
	0xe1, 0x25, 0xaf, 0x4c, 0x3c, 0xda, 0x1b, 0x1d, 0x44, 0xb3, 0x27, 0x19, 0x06, 0xc8, 0xae, 0xe5,
	0x67, 0x9f, 0x78, 0x65, 0xe6, 0xb7, 0xdb, 0x26, 0x8d, 0xce, 0xa5, 0x5e, 0xbd, 0xa2, 0x9d, 0x98,
	0x32, 0xf4, 0x20, 0xaf, 0x21, 0xcc, 0xcb, 0x95, 0x89, 0x03, 0xf4, 0xba, 0x9f, 0xb9, 0x58, 0x59,
	0x1f, 0x2b, 0xfb, 0x80, 0xb1, 0x86, 0xea, 0x8e, 0x4e, 0x19, 0xaa, 0x92, 0x5f, 0x01, 0x84, 0x9d,
	0x3b, 0xf9, 0x02, 0x13, 0xdf, 0x4a, 0x3c, 0xda, 0xdb, 0x3a, 0x88, 0x66, 0x87, 0xff, 0x95, 0x2a,
	0xf3, 0xb3, 0x39, 0x69, 0x9b, 0xf4, 0x96, 0xf3, 0xf7, 0x3e, 0x94, 0xf5, 0x8e, 0xc9, 0x8f, 0x00,
	0x26, 0x9e, 0x48, 0x5e, 0x00, 0xe4, 0x42, 0x55, 0xf2, 0xfb, 0x29, 0xb7, 0x17, 0xd8, 0xc0, 0x74,
	0x7e, 0xb7, 0x6d, 0xd2, 0x3b, 0x3e, 0xdc, 0x1f, 0x8c, 0xb2, 0x01, 0x91, 0x3c, 0x83, 0x09, 0xe6,
	0x39, 0xce, 0xf1, 0xa5, 0xd3, 0xe1, 0x85, 0x1e, 0xa0, 0xac, 0xa7, 0x90, 0x37, 0xb0, 0xcd, 0x37,
	0x89, 0x8f, 0xf3, 0x78, 0x0b, 0x35, 0x71, 0xdb, 0xa4, 0xbb, 0x03, 0x4d, 0x0f, 0x53, 0x76, 0x95,
	0x4e, 0x66, 0x30, 0xbd, 0x14, 0xda, 0x94, 0xb2, 0xd3, 0x86, 0xa8, 0xdd, 0x6d, 0x9b, 0x74, 0xc7,
	0xd7, 0xdf, 0x43, 0x94, 0x6d, 0x68, 0xe4, 0x11, 0x04, 0xa5, 0x8a, 0xc7, 0x48, 0xde, 0x6e, 0x9b,
	0x74, 0xea, 0xc8, 0xa5, 0xa2, 0x2c, 0x28, 0x15, 0x3d, 0x81, 0xe4, 0x6f, 0x55, 0x1a, 0x25, 0xd7,
	0x46, 0x74, 0xcf, 0x5b, 0x09, 0x63, 0x78, 0x21, 0x7c, 0x25, 0x83, 0xe7, 0x79, 0x80, 0xb2, 0x9e,
	0x42, 0xdb, 0x11, 0xec, 0x5f, 0x6f, 0xf6, 0x59, 0x73, 0xa5, 0x84, 0x26, 0x8f, 0x21, 0x5c, 0xca,
	0xdc, 0x19, 0x8e, 0x87, 0x0b, 0xd0, 0x4d, 0x29, 0x43, 0x90, 0xbc, 0x84, 0xa8, 0xfb, 0xbe, 0xfb,
	0xa6, 0x2a, 0x5e, 0xae, 0x7d, 0xb7, 0xf7, 0xda, 0x26, 0x25, 0x1b, 0xae, 0x07, 0x29, 0x1b, 0x52,
	0xc9, 0x53, 0x18, 0x0b, 0xad, 0xa5, 0xf6, 0xdd, 0xee, 0xb4, 0x4d, 0x7a, 0xd3, 0x69, 0x70, 0x4c,
	0x99, 0x83, 0xc9, 0x7b, 0x08, 0x73, 0x6e, 0x39, 0xd6, 0x18, 0xcd, 0xf6, 0xff, 0xb1, 0x56, 0x2e,
	0xfe, 0x95, 0x55, 0xe5, 0x96, 0x77, 0xab, 0xca, 0x2d, 0x5f, 0xdc, 0xc0, 0x95, 0x7e, 0xfe, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xe2, 0x2b, 0xf9, 0x96, 0xcb, 0x03, 0x00, 0x00,
}
