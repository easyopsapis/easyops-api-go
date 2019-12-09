// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_agent.proto

package agent

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
//GetAgent请求
type GetAgentRequest struct {
	//
	//实例唯一标识
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentRequest) Reset()         { *m = GetAgentRequest{} }
func (m *GetAgentRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentRequest) ProtoMessage()    {}
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_622294a87105ff31, []int{0}
}
func (m *GetAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentRequest.Unmarshal(m, b)
}
func (m *GetAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentRequest.Marshal(b, m, deterministic)
}
func (m *GetAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentRequest.Merge(m, src)
}
func (m *GetAgentRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentRequest.Size(m)
}
func (m *GetAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentRequest proto.InternalMessageInfo

func (m *GetAgentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetAgentApi返回
type GetAgentResponseWrapper struct {
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
	Data                 *agent_admin.Agent `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAgentResponseWrapper) Reset()         { *m = GetAgentResponseWrapper{} }
func (m *GetAgentResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetAgentResponseWrapper) ProtoMessage()    {}
func (*GetAgentResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_622294a87105ff31, []int{1}
}
func (m *GetAgentResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentResponseWrapper.Unmarshal(m, b)
}
func (m *GetAgentResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetAgentResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentResponseWrapper.Merge(m, src)
}
func (m *GetAgentResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetAgentResponseWrapper.Size(m)
}
func (m *GetAgentResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentResponseWrapper proto.InternalMessageInfo

func (m *GetAgentResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAgentResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAgentResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetAgentResponseWrapper) GetData() *agent_admin.Agent {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAgentRequest)(nil), "agent.GetAgentRequest")
	proto.RegisterType((*GetAgentResponseWrapper)(nil), "agent.GetAgentResponseWrapper")
}

func init() { proto.RegisterFile("get_agent.proto", fileDescriptor_622294a87105ff31) }

var fileDescriptor_622294a87105ff31 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0xfe, 0xdb, 0x1f, 0x96, 0x29, 0x93, 0x1c, 0xb4, 0x0c, 0xa4, 0x23, 0x82, 0xec,
	0xb2, 0x46, 0xf4, 0xa0, 0x78, 0x73, 0x20, 0x9e, 0xed, 0xc5, 0xe3, 0xc8, 0x9a, 0x77, 0x35, 0xd0,
	0xf6, 0x8d, 0x49, 0x06, 0xfa, 0x51, 0xbd, 0xf4, 0x43, 0xf4, 0x13, 0x48, 0xdf, 0x0e, 0xd6, 0x53,
	0xf2, 0xbc, 0xcf, 0xf3, 0xe3, 0x7d, 0x12, 0x36, 0x2f, 0x20, 0x6c, 0x55, 0x01, 0x75, 0x48, 0xad,
	0xc3, 0x80, 0x7c, 0x42, 0x62, 0xb1, 0x2e, 0x4c, 0xf8, 0x3c, 0xec, 0xd2, 0x1c, 0x2b, 0x59, 0x60,
	0x81, 0x92, 0xdc, 0xdd, 0x61, 0x4f, 0x8a, 0x04, 0xdd, 0x7a, 0x6a, 0xf1, 0x5e, 0x60, 0x0a, 0xca,
	0xff, 0xa0, 0xf5, 0x69, 0x89, 0xb9, 0x2a, 0x65, 0x8e, 0x75, 0x70, 0x2a, 0x0f, 0xbe, 0x27, 0x1d,
	0x58, 0x5c, 0x57, 0xa8, 0xa1, 0xf4, 0xf2, 0x18, 0x94, 0x24, 0x25, 0xed, 0xdb, 0x2a, 0x5d, 0x99,
	0x5a, 0x0e, 0x8a, 0x88, 0x3b, 0x36, 0x7f, 0x83, 0xf0, 0xd2, 0x4d, 0x32, 0xf8, 0x3a, 0x80, 0x0f,
	0xfc, 0x9a, 0x8d, 0x8c, 0x8e, 0xa3, 0x65, 0xb4, 0x9a, 0x6e, 0xce, 0xdb, 0x26, 0x99, 0xee, 0xd1,
	0x55, 0xcf, 0xc2, 0x68, 0x91, 0x8d, 0x8c, 0x16, 0xbf, 0x11, 0xbb, 0x3a, 0x21, 0xde, 0x62, 0xed,
	0xe1, 0xc3, 0x29, 0x6b, 0xc1, 0xf1, 0x1b, 0x36, 0xce, 0x51, 0x03, 0xc1, 0x93, 0xcd, 0xbc, 0x6d,
	0x92, 0x59, 0x0f, 0x77, 0x53, 0x91, 0x91, 0xc9, 0x9f, 0xd8, 0xac, 0x3b, 0x5f, 0xbf, 0x6d, 0xa9,
	0x4c, 0x1d, 0x8f, 0x68, 0xd1, 0x65, 0xdb, 0x24, 0xfc, 0x94, 0x3d, 0x9a, 0x22, 0x1b, 0x46, 0xf9,
	0x2d, 0x9b, 0x80, 0x73, 0xe8, 0xe2, 0x7f, 0xc4, 0x5c, 0xb4, 0x4d, 0x72, 0xd6, 0x33, 0x34, 0x16,
	0x59, 0x6f, 0xf3, 0x47, 0x36, 0xd6, 0x2a, 0xa8, 0x78, 0xbc, 0x8c, 0x56, 0xb3, 0x7b, 0x9e, 0x0e,
	0x1e, 0x9f, 0x52, 0xef, 0x61, 0xb5, 0x2e, 0x29, 0x32, 0x02, 0x76, 0xff, 0xe9, 0x53, 0x1e, 0xfe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x7f, 0x21, 0xeb, 0xb0, 0x01, 0x00, 0x00,
}