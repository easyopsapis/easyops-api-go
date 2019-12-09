// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_flow_result.proto

package execute

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/flow"
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
//GetFlowResult请求
type GetFlowResultRequest struct {
	//
	//流程执行任务ID
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlowResultRequest) Reset()         { *m = GetFlowResultRequest{} }
func (m *GetFlowResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlowResultRequest) ProtoMessage()    {}
func (*GetFlowResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e734a9aea6d6542, []int{0}
}
func (m *GetFlowResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowResultRequest.Unmarshal(m, b)
}
func (m *GetFlowResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowResultRequest.Marshal(b, m, deterministic)
}
func (m *GetFlowResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowResultRequest.Merge(m, src)
}
func (m *GetFlowResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlowResultRequest.Size(m)
}
func (m *GetFlowResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowResultRequest proto.InternalMessageInfo

func (m *GetFlowResultRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

//
//GetFlowResultApi返回
type GetFlowResultResponseWrapper struct {
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
	Data                 *flow.FlowInstance `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFlowResultResponseWrapper) Reset()         { *m = GetFlowResultResponseWrapper{} }
func (m *GetFlowResultResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetFlowResultResponseWrapper) ProtoMessage()    {}
func (*GetFlowResultResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e734a9aea6d6542, []int{1}
}
func (m *GetFlowResultResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowResultResponseWrapper.Unmarshal(m, b)
}
func (m *GetFlowResultResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowResultResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetFlowResultResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowResultResponseWrapper.Merge(m, src)
}
func (m *GetFlowResultResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetFlowResultResponseWrapper.Size(m)
}
func (m *GetFlowResultResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowResultResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowResultResponseWrapper proto.InternalMessageInfo

func (m *GetFlowResultResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetFlowResultResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetFlowResultResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetFlowResultResponseWrapper) GetData() *flow.FlowInstance {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFlowResultRequest)(nil), "execute.GetFlowResultRequest")
	proto.RegisterType((*GetFlowResultResponseWrapper)(nil), "execute.GetFlowResultResponseWrapper")
}

func init() { proto.RegisterFile("get_flow_result.proto", fileDescriptor_5e734a9aea6d6542) }

var fileDescriptor_5e734a9aea6d6542 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0x5f, 0x8b, 0xd4, 0x30,
	0x10, 0x67, 0x75, 0xef, 0xc4, 0xac, 0xa2, 0x06, 0x95, 0xe5, 0x38, 0xd8, 0x23, 0x82, 0x9c, 0x0f,
	0xdb, 0x80, 0x82, 0x8a, 0x6f, 0x1e, 0xa8, 0xdc, 0x6b, 0x1e, 0xf4, 0xb1, 0x64, 0xd3, 0xd9, 0x58,
	0x2e, 0xed, 0xc4, 0x64, 0x6a, 0xcf, 0x2f, 0xdb, 0x0f, 0xd1, 0x4f, 0x20, 0x4d, 0x83, 0xbb, 0xf8,
	0x7c, 0x2f, 0xed, 0xcc, 0xfc, 0xfe, 0xcc, 0xe4, 0xc7, 0x5e, 0x58, 0xa0, 0x72, 0xef, 0xb0, 0x2f,
	0x03, 0xc4, 0xce, 0x51, 0xe1, 0x03, 0x12, 0xf2, 0x07, 0x70, 0x0b, 0xa6, 0x23, 0x38, 0xdb, 0xda,
	0x9a, 0x7e, 0x76, 0xbb, 0xc2, 0x60, 0x23, 0x2d, 0x5a, 0x94, 0x09, 0xdf, 0x75, 0xfb, 0xd4, 0xa5,
	0x26, 0x55, 0xb3, 0xee, 0xec, 0xbb, 0xc5, 0x02, 0x74, 0xfc, 0x83, 0x3e, 0x16, 0x0e, 0x8d, 0x76,
	0xd2, 0x60, 0x4b, 0x41, 0x1b, 0x8a, 0xb3, 0x32, 0x80, 0xc7, 0x6d, 0x83, 0x15, 0xb8, 0x28, 0x33,
	0x51, 0xa6, 0x56, 0x4e, 0x47, 0xa4, 0x4f, 0x99, 0x77, 0x97, 0x91, 0xc0, 0x67, 0x5f, 0x75, 0x17,
	0xbe, 0x75, 0x1b, 0x49, 0xb7, 0x06, 0xb2, 0xe7, 0xfb, 0xa3, 0xa7, 0x35, 0x7d, 0x4d, 0x37, 0xd8,
	0x4b, 0x8b, 0xdb, 0x04, 0x6e, 0x7f, 0x6b, 0x57, 0x57, 0x9a, 0x30, 0x44, 0xf9, 0xaf, 0xcc, 0xba,
	0x73, 0x8b, 0x68, 0x1d, 0x1c, 0x92, 0x88, 0x14, 0x3a, 0x93, 0x93, 0x13, 0x9f, 0xd9, 0xf3, 0x6f,
	0x40, 0x5f, 0x1d, 0xf6, 0x2a, 0x05, 0xaa, 0xe0, 0x57, 0x07, 0x91, 0xf8, 0x1b, 0x76, 0x4a, 0x3a,
	0xde, 0x5c, 0x57, 0xeb, 0xc5, 0xc5, 0xe2, 0xf2, 0xe1, 0xd5, 0xb3, 0x71, 0xd8, 0x3c, 0xde, 0x63,
	0x68, 0x3e, 0x89, 0x79, 0x2e, 0x54, 0x26, 0x88, 0x61, 0xc1, 0xce, 0xff, 0xf3, 0x88, 0x1e, 0xdb,
	0x08, 0x3f, 0x82, 0xf6, 0x1e, 0x02, 0x7f, 0xc5, 0x96, 0x06, 0x2b, 0x48, 0x4e, 0x27, 0x57, 0x4f,
	0xc6, 0x61, 0xb3, 0x9a, 0x9d, 0xa6, 0xa9, 0x50, 0x09, 0xe4, 0x1f, 0xd9, 0x6a, 0xfa, 0x7f, 0xb9,
	0xf5, 0x4e, 0xd7, 0xed, 0xfa, 0x5e, 0xda, 0xfa, 0x72, 0x1c, 0x36, 0xfc, 0xc0, 0xcd, 0xa0, 0x50,
	0xc7, 0x54, 0xfe, 0x9a, 0x9d, 0x40, 0x08, 0x18, 0xd6, 0xf7, 0x93, 0xe6, 0xe9, 0x38, 0x6c, 0x1e,
	0xcd, 0x9a, 0x34, 0x16, 0x6a, 0x86, 0xf9, 0x07, 0xb6, 0xac, 0x34, 0xe9, 0xf5, 0xf2, 0x62, 0x71,
	0xb9, 0x7a, 0xcb, 0x8b, 0x29, 0xe4, 0x62, 0xba, 0xfa, 0x3a, 0x07, 0x7d, 0x7c, 0xda, 0xc4, 0x14,
	0x2a, 0x09, 0x76, 0xa7, 0x29, 0xaa, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x28, 0x80,
	0x92, 0x7d, 0x02, 0x00, 0x00,
}
