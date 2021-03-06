// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_plugin.proto

package plugin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//GetPlugin请求
type GetPluginRequest struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginRequest) Reset()         { *m = GetPluginRequest{} }
func (m *GetPluginRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginRequest) ProtoMessage()    {}
func (*GetPluginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d67018434fae2e2c, []int{0}
}
func (m *GetPluginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginRequest.Unmarshal(m, b)
}
func (m *GetPluginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginRequest.Merge(m, src)
}
func (m *GetPluginRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginRequest.Size(m)
}
func (m *GetPluginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginRequest proto.InternalMessageInfo

func (m *GetPluginRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetPluginApi返回
type GetPluginResponseWrapper struct {
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
	Data                 *agent_admin.Plugin `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetPluginResponseWrapper) Reset()         { *m = GetPluginResponseWrapper{} }
func (m *GetPluginResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetPluginResponseWrapper) ProtoMessage()    {}
func (*GetPluginResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_d67018434fae2e2c, []int{1}
}
func (m *GetPluginResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginResponseWrapper.Unmarshal(m, b)
}
func (m *GetPluginResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetPluginResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginResponseWrapper.Merge(m, src)
}
func (m *GetPluginResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetPluginResponseWrapper.Size(m)
}
func (m *GetPluginResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginResponseWrapper proto.InternalMessageInfo

func (m *GetPluginResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPluginResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetPluginResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetPluginResponseWrapper) GetData() *agent_admin.Plugin {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPluginRequest)(nil), "plugin.GetPluginRequest")
	proto.RegisterType((*GetPluginResponseWrapper)(nil), "plugin.GetPluginResponseWrapper")
}

func init() { proto.RegisterFile("get_plugin.proto", fileDescriptor_d67018434fae2e2c) }

var fileDescriptor_d67018434fae2e2c = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xdc, 0x06, 0xcb, 0x14, 0x47, 0x05, 0x29, 0x03, 0xe9, 0x88, 0x20, 0xbb, 0xac,
	0x41, 0xbd, 0x0c, 0x8f, 0x03, 0xf1, 0x2a, 0xb9, 0x78, 0x1c, 0x59, 0x93, 0xc5, 0x40, 0xdb, 0x37,
	0x26, 0x29, 0xe8, 0x67, 0x15, 0xfa, 0x21, 0xfa, 0x09, 0x64, 0x6f, 0x06, 0xf6, 0xd4, 0x3e, 0x7f,
	0x7e, 0xe4, 0x49, 0xc8, 0x42, 0xab, 0xb0, 0xb7, 0x55, 0xab, 0x4d, 0x53, 0x58, 0x07, 0x01, 0xd2,
	0x69, 0x54, 0xcb, 0x8d, 0x36, 0xe1, 0xb3, 0x3d, 0x14, 0x25, 0xd4, 0x4c, 0x83, 0x06, 0x86, 0xf1,
	0xa1, 0x3d, 0xa2, 0x42, 0x81, 0x7f, 0x11, 0x5b, 0x72, 0x0d, 0x85, 0x12, 0xfe, 0x07, 0xac, 0x2f,
	0x2a, 0x28, 0x45, 0xc5, 0x4a, 0x68, 0x82, 0x13, 0x65, 0xf0, 0x91, 0x74, 0xca, 0xc2, 0xa6, 0x06,
	0xa9, 0x2a, 0xcf, 0xce, 0x45, 0x86, 0x92, 0x09, 0xad, 0x9a, 0xb0, 0x17, 0xb2, 0x36, 0x0d, 0x1b,
	0x4e, 0xa1, 0x8f, 0x64, 0xf1, 0xa6, 0xc2, 0x3b, 0x5a, 0x5c, 0x7d, 0xb5, 0xca, 0x87, 0xf4, 0x8e,
	0x8c, 0x8c, 0xcc, 0x92, 0x55, 0xb2, 0x9e, 0xed, 0xae, 0xfa, 0x2e, 0x9f, 0x1d, 0xc1, 0xd5, 0x2f,
	0xd4, 0x48, 0xca, 0x47, 0x46, 0xd2, 0xdf, 0x84, 0x64, 0x03, 0xc6, 0x5b, 0x68, 0xbc, 0xfa, 0x70,
	0xc2, 0x5a, 0xe5, 0xd2, 0x7b, 0x32, 0x2e, 0x41, 0x2a, 0xa4, 0x27, 0xbb, 0xeb, 0xbe, 0xcb, 0xe7,
	0x91, 0x3e, 0xb9, 0x94, 0x63, 0x98, 0x6e, 0xc9, 0xfc, 0xf4, 0x7d, 0xfd, 0xb6, 0x95, 0x30, 0x4d,
	0x36, 0xc2, 0x93, 0x6e, 0xfb, 0x2e, 0x4f, 0xff, 0xbb, 0xe7, 0x90, 0xf2, 0x61, 0x35, 0x7d, 0x20,
	0x13, 0xe5, 0x1c, 0xb8, 0xec, 0x02, 0x99, 0x45, 0xdf, 0xe5, 0x97, 0x91, 0x41, 0x9b, 0xf2, 0x18,
	0xa7, 0x5b, 0x32, 0x96, 0x22, 0x88, 0x6c, 0xbc, 0x4a, 0xd6, 0xf3, 0xa7, 0x9b, 0x62, 0x70, 0xff,
	0x22, 0x0e, 0x1f, 0x6e, 0x3b, 0x55, 0x29, 0x47, 0xe2, 0x30, 0xc5, 0x77, 0x79, 0xfe, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x12, 0x75, 0x17, 0xb6, 0x01, 0x00, 0x00,
}
