// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execute_strategy.proto

package datafilter

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//ExecuteStrategy请求
type ExecuteStrategyRequest struct {
	//
	//规则实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//是否发送消息
	SendMsg              bool     `protobuf:"varint,2,opt,name=sendMsg,proto3" json:"sendMsg" form:"sendMsg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteStrategyRequest) Reset()         { *m = ExecuteStrategyRequest{} }
func (m *ExecuteStrategyRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStrategyRequest) ProtoMessage()    {}
func (*ExecuteStrategyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1e6113c2a1ee6a1, []int{0}
}
func (m *ExecuteStrategyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStrategyRequest.Unmarshal(m, b)
}
func (m *ExecuteStrategyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStrategyRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStrategyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStrategyRequest.Merge(m, src)
}
func (m *ExecuteStrategyRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStrategyRequest.Size(m)
}
func (m *ExecuteStrategyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStrategyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStrategyRequest proto.InternalMessageInfo

func (m *ExecuteStrategyRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *ExecuteStrategyRequest) GetSendMsg() bool {
	if m != nil {
		return m.SendMsg
	}
	return false
}

//
//ExecuteStrategy返回
type ExecuteStrategyResponse struct {
	//
	//执行id
	ExecId               string   `protobuf:"bytes,1,opt,name=execId,proto3" json:"execId" form:"execId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteStrategyResponse) Reset()         { *m = ExecuteStrategyResponse{} }
func (m *ExecuteStrategyResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStrategyResponse) ProtoMessage()    {}
func (*ExecuteStrategyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1e6113c2a1ee6a1, []int{1}
}
func (m *ExecuteStrategyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStrategyResponse.Unmarshal(m, b)
}
func (m *ExecuteStrategyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStrategyResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStrategyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStrategyResponse.Merge(m, src)
}
func (m *ExecuteStrategyResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStrategyResponse.Size(m)
}
func (m *ExecuteStrategyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStrategyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStrategyResponse proto.InternalMessageInfo

func (m *ExecuteStrategyResponse) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

//
//ExecuteStrategyApi返回
type ExecuteStrategyResponseWrapper struct {
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
	Data                 *ExecuteStrategyResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExecuteStrategyResponseWrapper) Reset()         { *m = ExecuteStrategyResponseWrapper{} }
func (m *ExecuteStrategyResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ExecuteStrategyResponseWrapper) ProtoMessage()    {}
func (*ExecuteStrategyResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1e6113c2a1ee6a1, []int{2}
}
func (m *ExecuteStrategyResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStrategyResponseWrapper.Unmarshal(m, b)
}
func (m *ExecuteStrategyResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStrategyResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ExecuteStrategyResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStrategyResponseWrapper.Merge(m, src)
}
func (m *ExecuteStrategyResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ExecuteStrategyResponseWrapper.Size(m)
}
func (m *ExecuteStrategyResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStrategyResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStrategyResponseWrapper proto.InternalMessageInfo

func (m *ExecuteStrategyResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ExecuteStrategyResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ExecuteStrategyResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ExecuteStrategyResponseWrapper) GetData() *ExecuteStrategyResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteStrategyRequest)(nil), "datafilter.ExecuteStrategyRequest")
	proto.RegisterType((*ExecuteStrategyResponse)(nil), "datafilter.ExecuteStrategyResponse")
	proto.RegisterType((*ExecuteStrategyResponseWrapper)(nil), "datafilter.ExecuteStrategyResponseWrapper")
}

func init() { proto.RegisterFile("execute_strategy.proto", fileDescriptor_c1e6113c2a1ee6a1) }

var fileDescriptor_c1e6113c2a1ee6a1 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x6e, 0xda, 0x30,
	0x18, 0x57, 0x36, 0x60, 0xc3, 0x6c, 0x0c, 0x7c, 0x60, 0x11, 0x87, 0x05, 0x99, 0x69, 0x02, 0x69,
	0x49, 0xb6, 0x21, 0x4d, 0xdb, 0x8e, 0xd1, 0x90, 0xc6, 0x61, 0x17, 0xf7, 0xd0, 0x43, 0xd5, 0x56,
	0x26, 0x31, 0x69, 0x54, 0x88, 0x53, 0xdb, 0x29, 0xb4, 0x55, 0x1f, 0xa2, 0x2f, 0x98, 0x4a, 0x7d,
	0x83, 0xe6, 0x09, 0xaa, 0xd8, 0x29, 0x44, 0xaa, 0x38, 0xf9, 0xfb, 0xbe, 0xdf, 0x1f, 0xfd, 0xec,
	0xcf, 0xa0, 0x47, 0x37, 0xd4, 0x4f, 0x25, 0x3d, 0x15, 0x92, 0x13, 0x49, 0xc3, 0x2b, 0x27, 0xe1,
	0x4c, 0x32, 0x08, 0x02, 0x22, 0xc9, 0x22, 0x5a, 0x4a, 0xca, 0xfb, 0x76, 0x18, 0xc9, 0xb3, 0x74,
	0xee, 0xf8, 0x6c, 0xe5, 0x86, 0x2c, 0x64, 0xae, 0xa2, 0xcc, 0xd3, 0x85, 0xea, 0x54, 0xa3, 0x2a,
	0x2d, 0xed, 0xff, 0xac, 0xd0, 0x57, 0xeb, 0x48, 0x9e, 0xb3, 0xb5, 0x1b, 0x32, 0x5b, 0x81, 0xf6,
	0x25, 0x59, 0x46, 0x01, 0x91, 0x8c, 0x0b, 0x77, 0x5b, 0x6a, 0x1d, 0xba, 0x33, 0x40, 0x6f, 0xaa,
	0xd3, 0x1c, 0x94, 0x61, 0x30, 0xbd, 0x48, 0xa9, 0x90, 0x70, 0x06, 0x40, 0x14, 0x0b, 0x49, 0x62,
	0x9f, 0xce, 0x02, 0xd3, 0x18, 0x18, 0xa3, 0xa6, 0x37, 0xce, 0x33, 0xab, 0xbb, 0x60, 0x7c, 0xf5,
	0x07, 0xed, 0x30, 0xf4, 0x70, 0x6f, 0x75, 0x40, 0xfb, 0xe4, 0xe8, 0x9b, 0xfd, 0x9b, 0xd8, 0xd7,
	0xc7, 0x37, 0xdf, 0x27, 0xb7, 0x9f, 0x71, 0x45, 0x0c, 0xbf, 0x82, 0x37, 0x82, 0xc6, 0xc1, 0x7f,
	0x11, 0x9a, 0xaf, 0x06, 0xc6, 0xe8, 0xad, 0x07, 0xf3, 0xcc, 0x6a, 0x6b, 0x9f, 0x12, 0x40, 0xf8,
	0x99, 0x82, 0xfe, 0x82, 0x8f, 0x2f, 0x22, 0x89, 0x84, 0xc5, 0x82, 0xc2, 0x31, 0x68, 0x14, 0x6f,
	0xb7, 0xcd, 0xd3, 0xcd, 0x33, 0xeb, 0xbd, 0xf6, 0xd1, 0x73, 0x84, 0x4b, 0x02, 0x7a, 0x34, 0xc0,
	0xa7, 0x3d, 0x36, 0x87, 0x9c, 0x24, 0x09, 0xe5, 0x70, 0x08, 0x6a, 0x3e, 0x0b, 0xa8, 0xf2, 0xaa,
	0x7b, 0x1f, 0xf2, 0xcc, 0x6a, 0x69, 0xaf, 0x62, 0x8a, 0xb0, 0x02, 0xe1, 0x2f, 0xd0, 0x2a, 0xce,
	0xe9, 0x26, 0x59, 0x92, 0x28, 0x56, 0xf9, 0x9b, 0x5e, 0x2f, 0xcf, 0x2c, 0xb8, 0xe3, 0x96, 0x20,
	0xc2, 0x55, 0x2a, 0xfc, 0x02, 0xea, 0x94, 0x73, 0xc6, 0xcd, 0xd7, 0x4a, 0xd3, 0xc9, 0x33, 0xeb,
	0x5d, 0x99, 0xb5, 0x18, 0x23, 0xac, 0x61, 0xf8, 0x0f, 0xd4, 0x8a, 0xc5, 0x9b, 0xb5, 0x81, 0x31,
	0x6a, 0xfd, 0x18, 0x3a, 0xbb, 0x5f, 0xe0, 0xec, 0xb9, 0x40, 0x35, 0x6b, 0xc1, 0x46, 0x58, 0x39,
	0xcc, 0x1b, 0x6a, 0xa9, 0x93, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x29, 0x9e, 0xbd, 0x61,
	0x02, 0x00, 0x00,
}
