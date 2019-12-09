// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execute.proto

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
//ExecuteFlow请求
type ExecuteFlowRequest struct {
	//
	//流程执行时输入参数
	//"flowInputs": {
	//"@agents": [
	//{
	//"ip": "192.168.100.163",
	//"instanceId": "5caee6c67d86a"
	//}
	//],
	//"test1": "test params",
	//"ddddd": "5caee6c67d86a",
	//"test2": "develop_1"
	//},
	//
	FlowInputs *types.Struct `protobuf:"bytes,1,opt,name=flowInputs,proto3" json:"flowInputs" form:"flowInputs"`
	//
	//自动填充参数
	InstanceMap *types.Struct `protobuf:"bytes,2,opt,name=instanceMap,proto3" json:"instanceMap" form:"instanceMap"`
	//
	//是否发送notify通知
	NeedNotify bool `protobuf:"varint,3,opt,name=needNotify,proto3" json:"needNotify" form:"needNotify"`
	//
	//父流程任务ID
	ParentTaskId string `protobuf:"bytes,4,opt,name=parentTaskId,proto3" json:"parentTaskId" form:"parentTaskId"`
	//
	//流程执行成功后的回调,有中间态回调,(部署在使用)
	Callback *tool.Callback `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback" form:"callback"`
	//
	//流程全部执行完回调,有别于callback字段
	EndCallback *tool.Callback `protobuf:"bytes,6,opt,name=endCallback,proto3" json:"endCallback" form:"endCallback"`
	//
	//流程ID
	FlowId string `protobuf:"bytes,7,opt,name=flowId,proto3" json:"flowId" form:"flowId"`
	//
	//版本号(时间戳-秒)
	Version int32 `protobuf:"varint,8,opt,name=version,proto3" json:"version" form:"version"`
	//
	//流程名称
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name" form:"name"`
	//
	//流程环境变量 { "youcash":"1234",
	//}
	FlowEnv *types.Value `protobuf:"bytes,10,opt,name=flowEnv,proto3" json:"flowEnv" form:"flowEnv"`
	//
	//流程环境类型
	Metadata *ExecuteFlowRequest_Metadata `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata" form:"metadata"`
	//
	//流程标签
	Tags                 []string `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags" form:"tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteFlowRequest) Reset()         { *m = ExecuteFlowRequest{} }
func (m *ExecuteFlowRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteFlowRequest) ProtoMessage()    {}
func (*ExecuteFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{0}
}
func (m *ExecuteFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteFlowRequest.Unmarshal(m, b)
}
func (m *ExecuteFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteFlowRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteFlowRequest.Merge(m, src)
}
func (m *ExecuteFlowRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteFlowRequest.Size(m)
}
func (m *ExecuteFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteFlowRequest proto.InternalMessageInfo

func (m *ExecuteFlowRequest) GetFlowInputs() *types.Struct {
	if m != nil {
		return m.FlowInputs
	}
	return nil
}

func (m *ExecuteFlowRequest) GetInstanceMap() *types.Struct {
	if m != nil {
		return m.InstanceMap
	}
	return nil
}

func (m *ExecuteFlowRequest) GetNeedNotify() bool {
	if m != nil {
		return m.NeedNotify
	}
	return false
}

func (m *ExecuteFlowRequest) GetParentTaskId() string {
	if m != nil {
		return m.ParentTaskId
	}
	return ""
}

func (m *ExecuteFlowRequest) GetCallback() *tool.Callback {
	if m != nil {
		return m.Callback
	}
	return nil
}

func (m *ExecuteFlowRequest) GetEndCallback() *tool.Callback {
	if m != nil {
		return m.EndCallback
	}
	return nil
}

func (m *ExecuteFlowRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ExecuteFlowRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ExecuteFlowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecuteFlowRequest) GetFlowEnv() *types.Value {
	if m != nil {
		return m.FlowEnv
	}
	return nil
}

func (m *ExecuteFlowRequest) GetMetadata() *ExecuteFlowRequest_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ExecuteFlowRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ExecuteFlowRequest_Metadata struct {
	//
	//环境
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//描述
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc" form:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteFlowRequest_Metadata) Reset()         { *m = ExecuteFlowRequest_Metadata{} }
func (m *ExecuteFlowRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*ExecuteFlowRequest_Metadata) ProtoMessage()    {}
func (*ExecuteFlowRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{0, 0}
}
func (m *ExecuteFlowRequest_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteFlowRequest_Metadata.Unmarshal(m, b)
}
func (m *ExecuteFlowRequest_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteFlowRequest_Metadata.Marshal(b, m, deterministic)
}
func (m *ExecuteFlowRequest_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteFlowRequest_Metadata.Merge(m, src)
}
func (m *ExecuteFlowRequest_Metadata) XXX_Size() int {
	return xxx_messageInfo_ExecuteFlowRequest_Metadata.Size(m)
}
func (m *ExecuteFlowRequest_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteFlowRequest_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteFlowRequest_Metadata proto.InternalMessageInfo

func (m *ExecuteFlowRequest_Metadata) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ExecuteFlowRequest_Metadata) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

//
//ExecuteFlow返回
type ExecuteFlowResponse struct {
	//
	//流程任务ID
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteFlowResponse) Reset()         { *m = ExecuteFlowResponse{} }
func (m *ExecuteFlowResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteFlowResponse) ProtoMessage()    {}
func (*ExecuteFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{1}
}
func (m *ExecuteFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteFlowResponse.Unmarshal(m, b)
}
func (m *ExecuteFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteFlowResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteFlowResponse.Merge(m, src)
}
func (m *ExecuteFlowResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteFlowResponse.Size(m)
}
func (m *ExecuteFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteFlowResponse proto.InternalMessageInfo

func (m *ExecuteFlowResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

//
//ExecuteFlowApi返回
type ExecuteFlowResponseWrapper struct {
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
	Data                 *ExecuteFlowResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExecuteFlowResponseWrapper) Reset()         { *m = ExecuteFlowResponseWrapper{} }
func (m *ExecuteFlowResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ExecuteFlowResponseWrapper) ProtoMessage()    {}
func (*ExecuteFlowResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_58179d2e1720ec81, []int{2}
}
func (m *ExecuteFlowResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteFlowResponseWrapper.Unmarshal(m, b)
}
func (m *ExecuteFlowResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteFlowResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ExecuteFlowResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteFlowResponseWrapper.Merge(m, src)
}
func (m *ExecuteFlowResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ExecuteFlowResponseWrapper.Size(m)
}
func (m *ExecuteFlowResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteFlowResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteFlowResponseWrapper proto.InternalMessageInfo

func (m *ExecuteFlowResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ExecuteFlowResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ExecuteFlowResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ExecuteFlowResponseWrapper) GetData() *ExecuteFlowResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteFlowRequest)(nil), "execute.ExecuteFlowRequest")
	proto.RegisterType((*ExecuteFlowRequest_Metadata)(nil), "execute.ExecuteFlowRequest.Metadata")
	proto.RegisterType((*ExecuteFlowResponse)(nil), "execute.ExecuteFlowResponse")
	proto.RegisterType((*ExecuteFlowResponseWrapper)(nil), "execute.ExecuteFlowResponseWrapper")
}

func init() { proto.RegisterFile("execute.proto", fileDescriptor_58179d2e1720ec81) }

var fileDescriptor_58179d2e1720ec81 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdb, 0x6e, 0xd3, 0x30,
	0x18, 0x56, 0x59, 0x4f, 0x73, 0x77, 0x60, 0xae, 0xd8, 0xa2, 0x6a, 0x52, 0x2a, 0x33, 0xa1, 0x22,
	0xd1, 0x04, 0x6d, 0x02, 0x71, 0xb8, 0x60, 0x2b, 0x6c, 0x62, 0x17, 0x9d, 0x84, 0x19, 0x70, 0x81,
	0x40, 0x72, 0x13, 0x37, 0x54, 0x4b, 0xe3, 0x10, 0x3b, 0xeb, 0x2a, 0xc4, 0xab, 0x06, 0x09, 0x24,
	0x1e, 0x20, 0x4f, 0x80, 0x62, 0x3b, 0xab, 0xab, 0x4d, 0x5c, 0x25, 0xf6, 0x77, 0xf0, 0xff, 0xfb,
	0xfb, 0x0d, 0xd6, 0xe9, 0x15, 0xf5, 0x52, 0x41, 0x9d, 0x38, 0x61, 0x82, 0xc1, 0x86, 0x5e, 0x76,
	0xfa, 0xc1, 0x44, 0x7c, 0x4b, 0x47, 0x8e, 0xc7, 0xa6, 0x6e, 0xc0, 0x02, 0xe6, 0x4a, 0x7c, 0x94,
	0x8e, 0xe5, 0x4a, 0x2e, 0xe4, 0x9f, 0xd2, 0x75, 0x86, 0x01, 0x73, 0x28, 0xe1, 0x73, 0x16, 0x73,
	0x27, 0x64, 0x1e, 0x09, 0x5d, 0x8f, 0x45, 0x22, 0x21, 0x9e, 0xe0, 0x4a, 0x99, 0xd0, 0x98, 0xf5,
	0xa7, 0xcc, 0xa7, 0x21, 0x77, 0x35, 0xd1, 0x95, 0x4b, 0x57, 0x30, 0x16, 0xba, 0x1e, 0x09, 0xc3,
	0x11, 0xf1, 0x2e, 0xb4, 0xdd, 0x53, 0xe3, 0xf4, 0xe9, 0x6c, 0x22, 0x2e, 0xd8, 0xcc, 0x0d, 0x58,
	0x5f, 0x82, 0xfd, 0x4b, 0x12, 0x4e, 0x7c, 0x22, 0x58, 0xc2, 0xdd, 0xeb, 0x5f, 0xad, 0xdb, 0x0d,
	0x18, 0x0b, 0x42, 0xba, 0x28, 0x96, 0x8b, 0x24, 0xf5, 0x84, 0x42, 0xd1, 0x9f, 0x3a, 0x80, 0xc7,
	0xaa, 0xbf, 0x93, 0x90, 0xcd, 0x30, 0xfd, 0x9e, 0x52, 0x2e, 0xe0, 0x19, 0x00, 0xe3, 0x90, 0xcd,
	0x4e, 0xa3, 0x38, 0x15, 0xdc, 0xaa, 0x74, 0x2b, 0xbd, 0xd6, 0xfe, 0x8e, 0xa3, 0x9c, 0x9c, 0xd2,
	0xc9, 0x79, 0x2f, 0x9d, 0x06, 0xf7, 0xf2, 0xcc, 0xde, 0x1a, 0xb3, 0x64, 0xfa, 0x02, 0x2d, 0x44,
	0x08, 0x1b, 0x0e, 0xf0, 0x1d, 0x68, 0x4d, 0x22, 0x2e, 0x48, 0xe4, 0xd1, 0x21, 0x89, 0xad, 0x3b,
	0xff, 0x37, 0xdc, 0xce, 0x33, 0x1b, 0x2a, 0x43, 0x43, 0x85, 0xb0, 0xe9, 0x01, 0x9f, 0x00, 0x10,
	0x51, 0xea, 0x9f, 0x31, 0x31, 0x19, 0xcf, 0xad, 0x95, 0x6e, 0xa5, 0xd7, 0x34, 0x2b, 0x59, 0x60,
	0x08, 0x1b, 0x44, 0xf8, 0x12, 0xac, 0xc5, 0x24, 0xa1, 0x91, 0x38, 0x27, 0xfc, 0xe2, 0xd4, 0xb7,
	0xaa, 0xdd, 0x4a, 0x6f, 0x75, 0xb0, 0x93, 0x67, 0x76, 0x5b, 0x09, 0x4d, 0x14, 0xe1, 0x25, 0x32,
	0x7c, 0x05, 0x9a, 0x65, 0x2a, 0x56, 0x4d, 0xf6, 0xb0, 0xe1, 0x14, 0x59, 0x39, 0xaf, 0xf5, 0xee,
	0xa0, 0x9d, 0x67, 0xf6, 0xa6, 0x32, 0x2a, 0x99, 0x08, 0x5f, 0x8b, 0xe0, 0x5b, 0xd0, 0xa2, 0x91,
	0x5f, 0xb2, 0xad, 0xfa, 0xad, 0x1e, 0x46, 0xfb, 0x06, 0x19, 0x61, 0x53, 0x0a, 0x0f, 0x41, 0x5d,
	0xde, 0xaf, 0x6f, 0x35, 0x64, 0x07, 0xbd, 0x3c, 0xb3, 0xd7, 0x8d, 0x10, 0x7c, 0xf4, 0xfb, 0x97,
	0xdd, 0x06, 0x5b, 0x5f, 0x3f, 0x93, 0xfe, 0xf8, 0xa8, 0x7f, 0xf2, 0xb8, 0xff, 0xfc, 0xcb, 0x8f,
	0x83, 0xfd, 0x9f, 0x7b, 0x58, 0xeb, 0xe0, 0x23, 0xd0, 0xb8, 0xa4, 0x09, 0x9f, 0xb0, 0xc8, 0x6a,
	0x76, 0x2b, 0xbd, 0xda, 0x00, 0xe6, 0x99, 0xbd, 0xa1, 0x2c, 0x34, 0x80, 0x70, 0x49, 0x81, 0xf7,
	0x41, 0x35, 0x22, 0x53, 0x6a, 0xad, 0xca, 0xd3, 0x36, 0xf3, 0xcc, 0x6e, 0xe9, 0x8b, 0x26, 0x53,
	0x8a, 0xb0, 0x04, 0xe1, 0x1b, 0xd0, 0x28, 0xcc, 0x8f, 0xa3, 0x4b, 0x0b, 0xc8, 0xd6, 0xb6, 0x6f,
	0x44, 0xfc, 0x91, 0x84, 0x29, 0x35, 0x8f, 0xd2, 0x02, 0x84, 0x4b, 0x29, 0xfc, 0x00, 0x9a, 0x53,
	0x2a, 0x88, 0x4f, 0x04, 0xb1, 0x5a, 0xd2, 0x66, 0xcf, 0x29, 0x9f, 0xe4, 0xcd, 0x59, 0x75, 0x86,
	0x9a, 0x6b, 0xde, 0x7d, 0xa9, 0x47, 0xf8, 0xda, 0xaa, 0xe8, 0x40, 0x90, 0x80, 0x5b, 0x6b, 0xdd,
	0x95, 0xe5, 0x0e, 0x8a, 0x5d, 0x84, 0x25, 0xd8, 0x39, 0x07, 0xcd, 0xa1, 0x29, 0x98, 0xc7, 0x54,
	0x8e, 0xff, 0xb2, 0x60, 0x1e, 0x17, 0x2d, 0x17, 0x9f, 0x82, 0xe4, 0x53, 0xee, 0xc9, 0x91, 0x5e,
	0x22, 0x15, 0xbb, 0x08, 0x4b, 0x10, 0x1d, 0x82, 0xf6, 0x52, 0xe1, 0x3c, 0x66, 0x11, 0xa7, 0xf0,
	0x21, 0xa8, 0x0b, 0x35, 0x85, 0xea, 0x88, 0xad, 0x45, 0x86, 0x42, 0xcf, 0x9f, 0x26, 0xa0, 0xbf,
	0x15, 0xd0, 0xb9, 0xc5, 0xe2, 0x53, 0x42, 0xe2, 0x98, 0x26, 0x45, 0x15, 0x1e, 0xf3, 0x55, 0xa9,
	0x35, 0xb3, 0x8a, 0x62, 0x17, 0x61, 0x09, 0xc2, 0x67, 0xa0, 0x55, 0x7c, 0x8f, 0xaf, 0xe2, 0x90,
	0x4c, 0x22, 0x5d, 0xb1, 0x31, 0x6c, 0x06, 0x88, 0xb0, 0x49, 0x85, 0x0f, 0x40, 0x8d, 0x26, 0x09,
	0x4b, 0xe4, 0x33, 0x5b, 0x1d, 0xdc, 0xcd, 0x33, 0x7b, 0x4d, 0x0f, 0x68, 0xb1, 0x8d, 0xb0, 0x82,
	0xe1, 0x11, 0xa8, 0xca, 0xd4, 0xaa, 0x32, 0xb5, 0xdd, 0xdb, 0x53, 0x53, 0x95, 0x2f, 0x5d, 0x95,
	0x4c, 0x4a, 0x4a, 0x47, 0x75, 0x39, 0x29, 0x07, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcd,
	0x50, 0x17, 0x85, 0x05, 0x00, 0x00,
}