// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execute_tool.proto

package execute

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	tool "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
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
//ExecuteTool请求
type ExecuteToolRequest struct {
	//
	//执行时输入参数
	//"inputs": {
	//"@agents": [
	//{
	//"ip": "192.168.100.163",
	//"instanceId": "5caee6c67d86a"
	//}
	//],
	//"ansibleServer": {
	//"ip": "192.168.100.163",
	//"hostInstanceId": "5caee6c67d86a"
	//}
	//},
	//
	Inputs *types.Struct `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs" form:"inputs"`
	//
	//工具执行成功的回调信息
	Callback *tool.Callback `protobuf:"bytes,2,opt,name=callback,proto3" json:"callback" form:"callback"`
	//
	//工具执行的通知信息
	NeedNotify string `protobuf:"bytes,3,opt,name=needNotify,proto3" json:"needNotify" form:"needNotify"`
	//
	//工具执行通知的用户
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" form:"name"`
	//
	//一些拓展字段 目前支持origin=页面访问地址
	Metadata *ExecuteToolRequest_Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata" form:"metadata"`
	//
	//工具Id，服务端自动生成
	ToolId string `protobuf:"bytes,6,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	//
	//工具版本
	VId string `protobuf:"bytes,7,opt,name=vId,proto3" json:"vId" form:"vId"`
	//
	//执行用户
	ExecUser string `protobuf:"bytes,8,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//工具执行的超时时间(0表示不超时，默认为86400秒)
	Timeout int32 `protobuf:"varint,9,opt,name=timeout,proto3" json:"timeout" form:"timeout"`
	//
	//ansible-playbook verbose mode
	Log *ExecuteToolRequest_Log `protobuf:"bytes,10,opt,name=log,proto3" json:"log" form:"log"`
	//
	//Ansible Node机器执行用户
	AnsibleNodeExecUser  string   `protobuf:"bytes,11,opt,name=ansibleNodeExecUser,proto3" json:"ansibleNodeExecUser" form:"ansibleNodeExecUser"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteToolRequest) Reset()         { *m = ExecuteToolRequest{} }
func (m *ExecuteToolRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteToolRequest) ProtoMessage()    {}
func (*ExecuteToolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87121a3358c06b4, []int{0}
}
func (m *ExecuteToolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteToolRequest.Unmarshal(m, b)
}
func (m *ExecuteToolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteToolRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteToolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteToolRequest.Merge(m, src)
}
func (m *ExecuteToolRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteToolRequest.Size(m)
}
func (m *ExecuteToolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteToolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteToolRequest proto.InternalMessageInfo

func (m *ExecuteToolRequest) GetInputs() *types.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ExecuteToolRequest) GetCallback() *tool.Callback {
	if m != nil {
		return m.Callback
	}
	return nil
}

func (m *ExecuteToolRequest) GetNeedNotify() string {
	if m != nil {
		return m.NeedNotify
	}
	return ""
}

func (m *ExecuteToolRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecuteToolRequest) GetMetadata() *ExecuteToolRequest_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ExecuteToolRequest) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func (m *ExecuteToolRequest) GetVId() string {
	if m != nil {
		return m.VId
	}
	return ""
}

func (m *ExecuteToolRequest) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *ExecuteToolRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ExecuteToolRequest) GetLog() *ExecuteToolRequest_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ExecuteToolRequest) GetAnsibleNodeExecUser() string {
	if m != nil {
		return m.AnsibleNodeExecUser
	}
	return ""
}

type ExecuteToolRequest_Metadata struct {
	//
	//页面访问地址
	Origin               string   `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin" form:"origin"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteToolRequest_Metadata) Reset()         { *m = ExecuteToolRequest_Metadata{} }
func (m *ExecuteToolRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*ExecuteToolRequest_Metadata) ProtoMessage()    {}
func (*ExecuteToolRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87121a3358c06b4, []int{0, 0}
}
func (m *ExecuteToolRequest_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteToolRequest_Metadata.Unmarshal(m, b)
}
func (m *ExecuteToolRequest_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteToolRequest_Metadata.Marshal(b, m, deterministic)
}
func (m *ExecuteToolRequest_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteToolRequest_Metadata.Merge(m, src)
}
func (m *ExecuteToolRequest_Metadata) XXX_Size() int {
	return xxx_messageInfo_ExecuteToolRequest_Metadata.Size(m)
}
func (m *ExecuteToolRequest_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteToolRequest_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteToolRequest_Metadata proto.InternalMessageInfo

func (m *ExecuteToolRequest_Metadata) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type ExecuteToolRequest_Log struct {
	//
	//是否启用日志
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled" form:"enabled"`
	//
	//v,vv,vvv for more, vvvv to enable connection debugging
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level" form:"level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteToolRequest_Log) Reset()         { *m = ExecuteToolRequest_Log{} }
func (m *ExecuteToolRequest_Log) String() string { return proto.CompactTextString(m) }
func (*ExecuteToolRequest_Log) ProtoMessage()    {}
func (*ExecuteToolRequest_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87121a3358c06b4, []int{0, 1}
}
func (m *ExecuteToolRequest_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteToolRequest_Log.Unmarshal(m, b)
}
func (m *ExecuteToolRequest_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteToolRequest_Log.Marshal(b, m, deterministic)
}
func (m *ExecuteToolRequest_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteToolRequest_Log.Merge(m, src)
}
func (m *ExecuteToolRequest_Log) XXX_Size() int {
	return xxx_messageInfo_ExecuteToolRequest_Log.Size(m)
}
func (m *ExecuteToolRequest_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteToolRequest_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteToolRequest_Log proto.InternalMessageInfo

func (m *ExecuteToolRequest_Log) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ExecuteToolRequest_Log) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

//
//ExecuteTool返回
type ExecuteToolResponse struct {
	//
	//工具执行id
	ExecId               string   `protobuf:"bytes,1,opt,name=execId,proto3" json:"execId" form:"execId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteToolResponse) Reset()         { *m = ExecuteToolResponse{} }
func (m *ExecuteToolResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteToolResponse) ProtoMessage()    {}
func (*ExecuteToolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87121a3358c06b4, []int{1}
}
func (m *ExecuteToolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteToolResponse.Unmarshal(m, b)
}
func (m *ExecuteToolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteToolResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteToolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteToolResponse.Merge(m, src)
}
func (m *ExecuteToolResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteToolResponse.Size(m)
}
func (m *ExecuteToolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteToolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteToolResponse proto.InternalMessageInfo

func (m *ExecuteToolResponse) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

//
//ExecuteToolApi返回
type ExecuteToolResponseWrapper struct {
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
	Data                 *ExecuteToolResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExecuteToolResponseWrapper) Reset()         { *m = ExecuteToolResponseWrapper{} }
func (m *ExecuteToolResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ExecuteToolResponseWrapper) ProtoMessage()    {}
func (*ExecuteToolResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87121a3358c06b4, []int{2}
}
func (m *ExecuteToolResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteToolResponseWrapper.Unmarshal(m, b)
}
func (m *ExecuteToolResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteToolResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ExecuteToolResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteToolResponseWrapper.Merge(m, src)
}
func (m *ExecuteToolResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ExecuteToolResponseWrapper.Size(m)
}
func (m *ExecuteToolResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteToolResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteToolResponseWrapper proto.InternalMessageInfo

func (m *ExecuteToolResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ExecuteToolResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ExecuteToolResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ExecuteToolResponseWrapper) GetData() *ExecuteToolResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteToolRequest)(nil), "execute.ExecuteToolRequest")
	proto.RegisterType((*ExecuteToolRequest_Metadata)(nil), "execute.ExecuteToolRequest.Metadata")
	proto.RegisterType((*ExecuteToolRequest_Log)(nil), "execute.ExecuteToolRequest.Log")
	proto.RegisterType((*ExecuteToolResponse)(nil), "execute.ExecuteToolResponse")
	proto.RegisterType((*ExecuteToolResponseWrapper)(nil), "execute.ExecuteToolResponseWrapper")
}

func init() { proto.RegisterFile("execute_tool.proto", fileDescriptor_d87121a3358c06b4) }

var fileDescriptor_d87121a3358c06b4 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6e, 0xdc, 0x44,
	0x14, 0xd5, 0x92, 0xcd, 0x66, 0x33, 0x81, 0x94, 0x4c, 0x04, 0x58, 0xab, 0x4a, 0x4e, 0x27, 0x69,
	0xb5, 0x51, 0x63, 0xbb, 0x24, 0xa2, 0xd0, 0x20, 0x94, 0xac, 0x51, 0x90, 0x82, 0xd2, 0x3c, 0x0c,
	0x14, 0xa4, 0x86, 0x10, 0xcd, 0xda, 0xb3, 0xae, 0xd5, 0x59, 0x5f, 0x63, 0xcf, 0x26, 0xb4, 0x49,
	0x79, 0xe5, 0x13, 0xf8, 0x3b, 0x23, 0xf1, 0xc0, 0x07, 0xf8, 0x0b, 0x90, 0x67, 0xc6, 0x6b, 0xa3,
	0x46, 0x55, 0xdf, 0x66, 0xee, 0xb9, 0xe7, 0xce, 0x3d, 0x67, 0xe6, 0x0e, 0xc2, 0xfc, 0x77, 0x1e,
	0xcc, 0x24, 0xbf, 0x90, 0x00, 0xc2, 0x4d, 0x33, 0x90, 0x80, 0x97, 0x4c, 0x6c, 0xe0, 0x44, 0xb1,
	0x7c, 0x31, 0x1b, 0xbb, 0x01, 0x4c, 0xbd, 0x08, 0x22, 0xf0, 0x14, 0x3e, 0x9e, 0x4d, 0xd4, 0x4e,
	0x6d, 0xd4, 0x4a, 0xf3, 0x06, 0x4f, 0x23, 0x70, 0x39, 0xcb, 0x5f, 0x41, 0x9a, 0xbb, 0x02, 0x02,
	0x26, 0xbc, 0x00, 0x12, 0x99, 0xb1, 0x40, 0xe6, 0x9a, 0x99, 0xf1, 0x14, 0x9c, 0x29, 0x84, 0x5c,
	0xe4, 0x9e, 0x49, 0xf4, 0xd4, 0xd6, 0xab, 0x4e, 0xf7, 0x02, 0x26, 0xc4, 0x98, 0x05, 0x2f, 0x4d,
	0xb9, 0xc7, 0xad, 0xd3, 0xa7, 0x57, 0xb1, 0x7c, 0x09, 0x57, 0x5e, 0x04, 0x8e, 0x02, 0x9d, 0x4b,
	0x26, 0xe2, 0x90, 0x49, 0xc8, 0x72, 0x6f, 0xbe, 0x34, 0xbc, 0xbb, 0x11, 0x40, 0x24, 0x78, 0xd3,
	0x6c, 0x2e, 0xb3, 0x59, 0x20, 0x35, 0x4a, 0xfe, 0xea, 0x23, 0x7c, 0xa4, 0xf5, 0xfd, 0x08, 0x20,
	0x28, 0xff, 0x6d, 0xc6, 0x73, 0x89, 0x7d, 0xd4, 0x8b, 0x93, 0x74, 0x26, 0x73, 0xab, 0xb3, 0xd1,
	0x19, 0xae, 0xec, 0x7e, 0xe6, 0xea, 0x2a, 0x6e, 0x5d, 0xc5, 0xfd, 0x41, 0x55, 0xf1, 0xd7, 0xca,
	0xc2, 0xfe, 0x68, 0x02, 0xd9, 0x74, 0x9f, 0x68, 0x02, 0xa1, 0x86, 0x89, 0x0f, 0x50, 0xbf, 0x96,
	0x60, 0x7d, 0xa0, 0xaa, 0xac, 0xba, 0xca, 0xd6, 0x6f, 0x4d, 0xd4, 0x5f, 0x2f, 0x0b, 0xfb, 0x8e,
	0x26, 0xd7, 0x99, 0x84, 0xce, 0x49, 0xf8, 0x0b, 0x84, 0x12, 0xce, 0xc3, 0x53, 0x90, 0xf1, 0xe4,
	0x95, 0xb5, 0xb0, 0xd1, 0x19, 0x2e, 0xfb, 0x9f, 0x94, 0x85, 0xbd, 0xa6, 0x29, 0x0d, 0x46, 0x68,
	0x2b, 0x11, 0x7f, 0x8f, 0xba, 0x09, 0x9b, 0x72, 0xab, 0xab, 0x08, 0x8f, 0xcb, 0xc2, 0x5e, 0x31,
	0x04, 0x36, 0xe5, 0xe4, 0x9f, 0xbf, 0xed, 0x4d, 0x74, 0xef, 0xd7, 0x33, 0xe6, 0xbc, 0x1e, 0x39,
	0xcf, 0x1f, 0x39, 0x4f, 0xce, 0xcf, 0xdc, 0xf9, 0xfa, 0xc2, 0x39, 0xbf, 0xde, 0xdd, 0xd9, 0xfb,
	0xfc, 0xcd, 0x16, 0x55, 0x35, 0xf0, 0x33, 0xd4, 0x9f, 0x72, 0xc9, 0x42, 0x26, 0x99, 0xb5, 0xa8,
	0x34, 0x6c, 0xb9, 0xe6, 0x39, 0xb8, 0x6f, 0xdb, 0xe6, 0x3e, 0x35, 0xb9, 0x6d, 0x65, 0x35, 0x9f,
	0xd0, 0x79, 0x29, 0x7c, 0x88, 0x7a, 0x95, 0x13, 0xc7, 0xa1, 0xd5, 0x53, 0x4d, 0x0e, 0x1b, 0x17,
	0x75, 0xbc, 0x6a, 0x73, 0x1d, 0xad, 0x55, 0x6d, 0x4e, 0x46, 0xce, 0x77, 0x55, 0x9b, 0xd7, 0x7b,
	0xbb, 0x6f, 0xb6, 0xa8, 0xe1, 0xe1, 0x2f, 0xd1, 0xc2, 0xe5, 0x71, 0x68, 0x2d, 0x29, 0xfa, 0xfd,
	0xb2, 0xb0, 0x91, 0xa6, 0x5f, 0xbe, 0x83, 0x5b, 0x31, 0xf0, 0x4f, 0xa8, 0x5f, 0x09, 0x78, 0x96,
	0xf3, 0xcc, 0xea, 0x2b, 0xf6, 0x7e, 0xd3, 0x6b, 0x8d, 0xbc, 0xb7, 0x4b, 0xf3, 0x5a, 0x78, 0x07,
	0x2d, 0xc9, 0x78, 0xca, 0x61, 0x26, 0xad, 0xe5, 0x8d, 0xce, 0x70, 0xd1, 0xc7, 0x65, 0x61, 0xaf,
	0x1a, 0x4d, 0x1a, 0x20, 0xb4, 0x4e, 0xc1, 0x23, 0xb4, 0x20, 0x20, 0xb2, 0x90, 0xb2, 0xd4, 0x7e,
	0x97, 0xa5, 0x27, 0x10, 0xf9, 0xab, 0x8d, 0x3e, 0x01, 0x11, 0xa1, 0x15, 0x17, 0x5f, 0xa1, 0x75,
	0x96, 0xe4, 0xf1, 0x58, 0xf0, 0x53, 0x08, 0xf9, 0x51, 0xad, 0x69, 0x45, 0x69, 0x3a, 0x2a, 0x0b,
	0x7b, 0xa0, 0x19, 0xb7, 0x24, 0xbd, 0xb7, 0xbc, 0xdb, 0x4e, 0x18, 0xfc, 0xd9, 0x41, 0xfd, 0xfa,
	0xa2, 0xf1, 0x35, 0xea, 0x41, 0x16, 0x47, 0x71, 0xa2, 0x06, 0x65, 0xd9, 0x0f, 0x9a, 0x9b, 0xd4,
	0xf1, 0xea, 0xac, 0x53, 0x74, 0x32, 0xfc, 0x65, 0x3c, 0x7c, 0x21, 0x65, 0x9a, 0x1f, 0xdc, 0x4c,
	0x64, 0x7a, 0x33, 0x89, 0x05, 0xdf, 0xde, 0xf7, 0xbc, 0xed, 0x83, 0x33, 0x67, 0xe4, 0x3c, 0x67,
	0xce, 0xeb, 0x47, 0xce, 0x93, 0x87, 0x0f, 0x0e, 0x37, 0xbd, 0xfb, 0x07, 0xdf, 0xfc, 0x71, 0x71,
	0x73, 0x6f, 0x7f, 0xc7, 0xfd, 0xfa, 0xfc, 0xe1, 0x5b, 0x60, 0x85, 0x9d, 0x53, 0x73, 0xe4, 0xe0,
	0x0c, 0x2d, 0x9c, 0x40, 0x54, 0x59, 0xcf, 0x13, 0x36, 0x16, 0x3c, 0x54, 0x4d, 0xf4, 0xdb, 0xd6,
	0x1b, 0x80, 0xd0, 0x3a, 0x05, 0x3f, 0x40, 0x8b, 0x82, 0x5f, 0x72, 0xa1, 0x66, 0x72, 0xd9, 0xff,
	0xb8, 0x2c, 0xec, 0x0f, 0x8d, 0xb7, 0x55, 0x98, 0x50, 0x0d, 0x93, 0x43, 0xb4, 0xfe, 0xbf, 0xeb,
	0xc8, 0x53, 0x48, 0x72, 0x8e, 0xb7, 0x51, 0xaf, 0xba, 0xad, 0xe3, 0xd0, 0x08, 0x6e, 0x7d, 0x00,
	0x3a, 0x4e, 0xa8, 0x49, 0x20, 0xff, 0x76, 0xd0, 0xe0, 0x96, 0x12, 0x3f, 0x67, 0x2c, 0x4d, 0x79,
	0x86, 0x37, 0x51, 0x37, 0x80, 0x90, 0xab, 0x3a, 0x8b, 0xfe, 0x9d, 0x66, 0x4e, 0xab, 0x28, 0xa1,
	0x0a, 0xc4, 0x5f, 0xa1, 0x95, 0x40, 0x99, 0x9f, 0x0a, 0x16, 0x27, 0xa6, 0xe7, 0x4f, 0xcb, 0xc2,
	0xc6, 0x4d, 0xae, 0x01, 0x09, 0x6d, 0xa7, 0x56, 0x3a, 0x79, 0x96, 0x41, 0x66, 0x3e, 0x8e, 0x96,
	0x4e, 0x15, 0x26, 0x54, 0xc3, 0x78, 0x84, 0xba, 0x6a, 0xbc, 0xbb, 0xea, 0x2d, 0xde, 0xbd, 0xfd,
	0x2d, 0xea, 0xce, 0xdb, 0x4d, 0xea, 0x91, 0x56, 0xd4, 0x71, 0x4f, 0xfd, 0x8a, 0x7b, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xba, 0xc8, 0x8a, 0x91, 0x3e, 0x06, 0x00, 0x00,
}