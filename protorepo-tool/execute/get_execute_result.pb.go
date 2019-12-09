// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_execute_result.proto

package execute

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
//GetExecuteResult请求
type GetExecuteResultRequest struct {
	//
	//工具执行的任务Id
	ExecId string `protobuf:"bytes,1,opt,name=execId,proto3" json:"execId" form:"execId"`
	//
	//精简模式，不展示执行结果的日志信息（包含 日志、输出、表格输出）,true|false(缺省)
	Brief string `protobuf:"bytes,2,opt,name=brief,proto3" json:"brief" form:"brief"`
	//
	//taskId列表；用";"隔开
	TargetIds string `protobuf:"bytes,3,opt,name=targetIds,proto3" json:"targetIds" form:"targetIds"`
	//
	//所属流程的taskId
	TaskId string `protobuf:"bytes,4,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//在所属流程中的步骤Id
	StepId               string   `protobuf:"bytes,5,opt,name=stepId,proto3" json:"stepId" form:"stepId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExecuteResultRequest) Reset()         { *m = GetExecuteResultRequest{} }
func (m *GetExecuteResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetExecuteResultRequest) ProtoMessage()    {}
func (*GetExecuteResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aa9f0a1953f8941, []int{0}
}
func (m *GetExecuteResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecuteResultRequest.Unmarshal(m, b)
}
func (m *GetExecuteResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecuteResultRequest.Marshal(b, m, deterministic)
}
func (m *GetExecuteResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecuteResultRequest.Merge(m, src)
}
func (m *GetExecuteResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetExecuteResultRequest.Size(m)
}
func (m *GetExecuteResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecuteResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecuteResultRequest proto.InternalMessageInfo

func (m *GetExecuteResultRequest) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

func (m *GetExecuteResultRequest) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

func (m *GetExecuteResultRequest) GetTargetIds() string {
	if m != nil {
		return m.TargetIds
	}
	return ""
}

func (m *GetExecuteResultRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *GetExecuteResultRequest) GetStepId() string {
	if m != nil {
		return m.StepId
	}
	return ""
}

//
//GetExecuteResultApi返回
type GetExecuteResultResponseWrapper struct {
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
	Data                 *tool.ToolTask `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetExecuteResultResponseWrapper) Reset()         { *m = GetExecuteResultResponseWrapper{} }
func (m *GetExecuteResultResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetExecuteResultResponseWrapper) ProtoMessage()    {}
func (*GetExecuteResultResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aa9f0a1953f8941, []int{1}
}
func (m *GetExecuteResultResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecuteResultResponseWrapper.Unmarshal(m, b)
}
func (m *GetExecuteResultResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecuteResultResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetExecuteResultResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecuteResultResponseWrapper.Merge(m, src)
}
func (m *GetExecuteResultResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetExecuteResultResponseWrapper.Size(m)
}
func (m *GetExecuteResultResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecuteResultResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecuteResultResponseWrapper proto.InternalMessageInfo

func (m *GetExecuteResultResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetExecuteResultResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetExecuteResultResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetExecuteResultResponseWrapper) GetData() *tool.ToolTask {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetExecuteResultRequest)(nil), "execute.GetExecuteResultRequest")
	proto.RegisterType((*GetExecuteResultResponseWrapper)(nil), "execute.GetExecuteResultResponseWrapper")
}

func init() { proto.RegisterFile("get_execute_result.proto", fileDescriptor_7aa9f0a1953f8941) }

var fileDescriptor_7aa9f0a1953f8941 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xe9, 0xda, 0x56, 0x76, 0xaa, 0x6b, 0x0d, 0xa2, 0x61, 0x11, 0x22, 0xe3, 0x1f, 0xf4,
	0x22, 0x89, 0x6e, 0x41, 0x54, 0xf0, 0x66, 0x61, 0x91, 0x5c, 0xe8, 0xc5, 0xb0, 0xe0, 0x85, 0x68,
	0x99, 0x24, 0xa7, 0x63, 0xe8, 0xb4, 0x27, 0xce, 0x4c, 0xdc, 0x55, 0xf1, 0x05, 0x7c, 0x47, 0x23,
	0xf8, 0x08, 0x79, 0x02, 0xc9, 0xcc, 0xd8, 0x16, 0xbd, 0xd4, 0x9b, 0x76, 0xe6, 0xfb, 0xbe, 0x73,
	0x38, 0xf3, 0x3b, 0x21, 0xa1, 0x00, 0x33, 0x87, 0x73, 0x28, 0x1a, 0x03, 0x73, 0x05, 0xba, 0x91,
	0x26, 0xa9, 0x15, 0x1a, 0x0c, 0x2e, 0x7a, 0xf5, 0x30, 0x16, 0x95, 0x79, 0xdf, 0xe4, 0x49, 0x81,
	0xab, 0x54, 0xa0, 0xc0, 0xd4, 0xfa, 0x79, 0xb3, 0xb0, 0x37, 0x7b, 0xb1, 0x27, 0x57, 0x77, 0xf8,
	0x52, 0x60, 0x02, 0x5c, 0x7f, 0xc2, 0x5a, 0x27, 0x12, 0x0b, 0x2e, 0xd3, 0x02, 0xd7, 0x46, 0xf1,
	0xc2, 0x68, 0x57, 0xa9, 0xa0, 0xc6, 0x78, 0x85, 0x25, 0x48, 0x9d, 0xfa, 0x60, 0x6a, 0xaf, 0xa9,
	0x41, 0x94, 0x69, 0xc1, 0xa5, 0xcc, 0x79, 0xb1, 0xf4, 0xed, 0xb2, 0x7f, 0x6c, 0xd7, 0xff, 0xf8,
	0x56, 0xaf, 0xfe, 0x43, 0xab, 0xb9, 0xe1, 0xfa, 0xf7, 0x68, 0x8f, 0x77, 0xc0, 0xac, 0xce, 0x2a,
	0xb3, 0xc4, 0xb3, 0x54, 0x60, 0x6c, 0xcd, 0xf8, 0x23, 0x97, 0x55, 0xc9, 0x0d, 0x2a, 0x9d, 0x6e,
	0x8e, 0xbe, 0xee, 0xa6, 0x40, 0x14, 0x12, 0xb6, 0x1c, 0xb5, 0x51, 0x4d, 0xe1, 0xb9, 0xd3, 0x6f,
	0x7b, 0xe4, 0xc6, 0x0b, 0x30, 0x27, 0x8e, 0x3e, 0xb3, 0x2b, 0x61, 0xf0, 0xa1, 0x01, 0x6d, 0x82,
	0x07, 0x64, 0xdc, 0x6f, 0x25, 0x2b, 0xc3, 0xc1, 0xad, 0xc1, 0xfd, 0xfd, 0xe3, 0xab, 0x5d, 0x1b,
	0x5d, 0x5e, 0xa0, 0x5a, 0x3d, 0xa3, 0x4e, 0xa7, 0xcc, 0x07, 0x82, 0x7b, 0x64, 0x94, 0xab, 0x0a,
	0x16, 0xe1, 0x9e, 0x4d, 0x4e, 0xbb, 0x36, 0xba, 0xe4, 0x92, 0x56, 0xa6, 0xcc, 0xd9, 0xc1, 0x11,
	0xd9, 0x37, 0x5c, 0x09, 0x30, 0x59, 0xa9, 0xc3, 0x0b, 0x36, 0x7b, 0xad, 0x6b, 0xa3, 0xa9, 0xcb,
	0x6e, 0x2c, 0xca, 0xb6, 0xb1, 0xe0, 0x39, 0x19, 0xf7, 0x18, 0xb2, 0x32, 0x1c, 0xda, 0x82, 0xbb,
	0xdb, 0x31, 0x9c, 0x4e, 0x7f, 0xfe, 0x88, 0xa6, 0xe4, 0xe0, 0xdd, 0x9b, 0x87, 0xf1, 0x53, 0x1e,
	0x7f, 0x7e, 0xfb, 0xe5, 0xd1, 0xec, 0xeb, 0x1d, 0xe6, 0x8b, 0xfa, 0x57, 0x68, 0x03, 0x75, 0x56,
	0x86, 0xa3, 0x3f, 0x5f, 0xe1, 0x74, 0xca, 0x7c, 0x80, 0x7e, 0x1f, 0x90, 0xe8, 0x6f, 0x18, 0xba,
	0xc6, 0xb5, 0x86, 0xd7, 0x8a, 0xd7, 0x35, 0xa8, 0xe0, 0x36, 0x19, 0x16, 0x58, 0x82, 0x45, 0x32,
	0x3a, 0xbe, 0xd2, 0xb5, 0xd1, 0xc4, 0x35, 0xeb, 0x55, 0xca, 0xac, 0x19, 0x3c, 0x21, 0x93, 0xfe,
	0xff, 0xe4, 0xbc, 0x96, 0xbc, 0x5a, 0x7b, 0x28, 0xd7, 0xbb, 0x36, 0x0a, 0xb6, 0x59, 0x6f, 0x52,
	0xb6, 0x1b, 0xed, 0x41, 0x82, 0x52, 0xa8, 0x3c, 0x9c, 0x1d, 0x90, 0x56, 0xa6, 0xcc, 0xd9, 0xc1,
	0x8c, 0x0c, 0x4b, 0x6e, 0xb8, 0x45, 0x32, 0x39, 0x3a, 0x48, 0xec, 0x87, 0x77, 0x8a, 0x28, 0x4f,
	0xb9, 0x5e, 0xee, 0x8e, 0xd5, 0xa7, 0x28, 0xb3, 0xe1, 0x7c, 0x6c, 0x77, 0x3e, 0xfb, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x53, 0x10, 0x9e, 0x87, 0x03, 0x00, 0x00,
}
