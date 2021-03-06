// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: install_plugin.proto

package admin_task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//InstallPluginTask请求
type InstallPluginTaskRequest struct {
	//
	//插件部署任务的详情数据
	Targets []*InstallPluginTaskRequest_Targets `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets" form:"targets"`
	//
	//部署类型
	TaskType string `protobuf:"bytes,2,opt,name=taskType,proto3" json:"taskType" form:"taskType"`
	//
	//每批次数量
	BatchNum int32 `protobuf:"varint,3,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//每批时间间隔
	BatchInterval        int32    `protobuf:"varint,4,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskRequest) Reset()         { *m = InstallPluginTaskRequest{} }
func (m *InstallPluginTaskRequest) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskRequest) ProtoMessage()    {}
func (*InstallPluginTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c2f6c8853dd07ff, []int{0}
}
func (m *InstallPluginTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskRequest.Unmarshal(m, b)
}
func (m *InstallPluginTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskRequest.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskRequest.Merge(m, src)
}
func (m *InstallPluginTaskRequest) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskRequest.Size(m)
}
func (m *InstallPluginTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskRequest proto.InternalMessageInfo

func (m *InstallPluginTaskRequest) GetTargets() []*InstallPluginTaskRequest_Targets {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *InstallPluginTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *InstallPluginTaskRequest) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *InstallPluginTaskRequest) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

type InstallPluginTaskRequest_Targets struct {
	//
	//AgentID(uuid)
	AgentId string `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId" form:"agentId"`
	//
	//ip
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//部署路径
	DeployPath []string `protobuf:"bytes,3,rep,name=deployPath,proto3" json:"deployPath" form:"deployPath"`
	//
	//插件包ID
	PluginId string `protobuf:"bytes,4,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//插件包名称
	PluginName string `protobuf:"bytes,5,opt,name=pluginName,proto3" json:"pluginName" form:"pluginName"`
	//
	//插件包版本名称
	PluginVersionName string `protobuf:"bytes,6,opt,name=pluginVersionName,proto3" json:"pluginVersionName" form:"pluginVersionName"`
	//
	//插件包版本ID
	PluginVersionId      string   `protobuf:"bytes,7,opt,name=pluginVersionId,proto3" json:"pluginVersionId" form:"pluginVersionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskRequest_Targets) Reset()         { *m = InstallPluginTaskRequest_Targets{} }
func (m *InstallPluginTaskRequest_Targets) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskRequest_Targets) ProtoMessage()    {}
func (*InstallPluginTaskRequest_Targets) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c2f6c8853dd07ff, []int{0, 0}
}
func (m *InstallPluginTaskRequest_Targets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskRequest_Targets.Unmarshal(m, b)
}
func (m *InstallPluginTaskRequest_Targets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskRequest_Targets.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskRequest_Targets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskRequest_Targets.Merge(m, src)
}
func (m *InstallPluginTaskRequest_Targets) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskRequest_Targets.Size(m)
}
func (m *InstallPluginTaskRequest_Targets) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskRequest_Targets.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskRequest_Targets proto.InternalMessageInfo

func (m *InstallPluginTaskRequest_Targets) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *InstallPluginTaskRequest_Targets) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *InstallPluginTaskRequest_Targets) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *InstallPluginTaskRequest_Targets) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *InstallPluginTaskRequest_Targets) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *InstallPluginTaskRequest_Targets) GetPluginVersionName() string {
	if m != nil {
		return m.PluginVersionName
	}
	return ""
}

func (m *InstallPluginTaskRequest_Targets) GetPluginVersionId() string {
	if m != nil {
		return m.PluginVersionId
	}
	return ""
}

//
//InstallPluginTask返回
type InstallPluginTaskResponse struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskResponse) Reset()         { *m = InstallPluginTaskResponse{} }
func (m *InstallPluginTaskResponse) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskResponse) ProtoMessage()    {}
func (*InstallPluginTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c2f6c8853dd07ff, []int{1}
}
func (m *InstallPluginTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskResponse.Unmarshal(m, b)
}
func (m *InstallPluginTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskResponse.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskResponse.Merge(m, src)
}
func (m *InstallPluginTaskResponse) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskResponse.Size(m)
}
func (m *InstallPluginTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskResponse proto.InternalMessageInfo

func (m *InstallPluginTaskResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//InstallPluginTaskApi返回
type InstallPluginTaskResponseWrapper struct {
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
	Data                 *InstallPluginTaskResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *InstallPluginTaskResponseWrapper) Reset()         { *m = InstallPluginTaskResponseWrapper{} }
func (m *InstallPluginTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskResponseWrapper) ProtoMessage()    {}
func (*InstallPluginTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c2f6c8853dd07ff, []int{2}
}
func (m *InstallPluginTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskResponseWrapper.Unmarshal(m, b)
}
func (m *InstallPluginTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskResponseWrapper.Merge(m, src)
}
func (m *InstallPluginTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskResponseWrapper.Size(m)
}
func (m *InstallPluginTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskResponseWrapper proto.InternalMessageInfo

func (m *InstallPluginTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InstallPluginTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InstallPluginTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InstallPluginTaskResponseWrapper) GetData() *InstallPluginTaskResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallPluginTaskRequest)(nil), "admin_task.InstallPluginTaskRequest")
	proto.RegisterType((*InstallPluginTaskRequest_Targets)(nil), "admin_task.InstallPluginTaskRequest.Targets")
	proto.RegisterType((*InstallPluginTaskResponse)(nil), "admin_task.InstallPluginTaskResponse")
	proto.RegisterType((*InstallPluginTaskResponseWrapper)(nil), "admin_task.InstallPluginTaskResponseWrapper")
}

func init() { proto.RegisterFile("install_plugin.proto", fileDescriptor_7c2f6c8853dd07ff) }

var fileDescriptor_7c2f6c8853dd07ff = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xd5, 0xb5, 0x5d, 0xa9, 0xcb, 0x28, 0x33, 0x63, 0x32, 0x15, 0x28, 0x95, 0x11, 0xa8,
	0x87, 0xd1, 0x4a, 0x43, 0x48, 0x68, 0x07, 0x0e, 0x15, 0x1c, 0xba, 0xc3, 0x34, 0x59, 0x15, 0x48,
	0x5c, 0x26, 0xb7, 0xf1, 0xd2, 0x68, 0x69, 0x6c, 0x12, 0x07, 0xb1, 0x27, 0xe0, 0x3d, 0x78, 0x30,
	0x3f, 0x84, 0x25, 0xee, 0x28, 0x9f, 0x93, 0xd6, 0x85, 0x4d, 0x9c, 0x62, 0xfb, 0xff, 0xff, 0xfd,
	0x9d, 0xef, 0xfb, 0x12, 0x74, 0x14, 0xa7, 0xb9, 0xe6, 0x49, 0x72, 0xa5, 0x92, 0x22, 0x8a, 0xd3,
	0xb1, 0xca, 0xa4, 0x96, 0x18, 0xf1, 0x70, 0x1d, 0xa7, 0x57, 0x9a, 0xe7, 0x37, 0x83, 0x37, 0x51,
	0xac, 0x57, 0xc5, 0x62, 0xbc, 0x94, 0xeb, 0x49, 0x24, 0x23, 0x39, 0x01, 0xcb, 0xa2, 0xb8, 0x86,
	0x1d, 0x6c, 0x60, 0xe5, 0x50, 0xfa, 0xab, 0x8d, 0xc8, 0xcc, 0x65, 0x5e, 0x42, 0xe4, 0x9c, 0xe7,
	0x37, 0x4c, 0x7c, 0x2b, 0x44, 0xae, 0xf1, 0x57, 0xd4, 0xd1, 0x3c, 0x8b, 0x84, 0xce, 0x49, 0x63,
	0xd8, 0x1c, 0xf5, 0x4e, 0x4f, 0xc6, 0xdb, 0x9b, 0xc6, 0xf7, 0x61, 0xe3, 0xb9, 0x63, 0xa6, 0xd8,
	0x9a, 0xe0, 0xd1, 0xb5, 0xcc, 0xd6, 0x67, 0xb4, 0x8a, 0xa1, 0xac, 0x0e, 0xc4, 0x13, 0xf4, 0xa0,
	0x4c, 0x99, 0xdf, 0x2a, 0x41, 0xf6, 0x86, 0x8d, 0x51, 0x77, 0xfa, 0xc4, 0x9a, 0xa0, 0x5f, 0xdb,
	0x9d, 0x42, 0xd9, 0xc6, 0x54, 0x02, 0x0b, 0xae, 0x97, 0xab, 0x8b, 0x62, 0x4d, 0x9a, 0xc3, 0xc6,
	0xa8, 0xed, 0x03, 0xb5, 0x42, 0xd9, 0xc6, 0x84, 0x3f, 0xa0, 0x03, 0x58, 0xcf, 0x52, 0x2d, 0xb2,
	0xef, 0x3c, 0x21, 0x2d, 0xa0, 0x88, 0x35, 0xc1, 0x91, 0x47, 0xd5, 0x32, 0x65, 0xbb, 0xf6, 0xc1,
	0xcf, 0x26, 0xea, 0x54, 0xa5, 0xe0, 0x13, 0xd4, 0xe1, 0x91, 0x48, 0xf5, 0x2c, 0x24, 0x0d, 0x78,
	0x59, 0xaf, 0xb6, 0x4a, 0xa0, 0xac, 0xb6, 0xe0, 0x17, 0x68, 0x2f, 0x56, 0x55, 0x55, 0x07, 0xd6,
	0x04, 0x5d, 0x67, 0x8c, 0x15, 0x65, 0x7b, 0xb1, 0xc2, 0xef, 0x10, 0x0a, 0x85, 0x4a, 0xe4, 0xed,
	0x25, 0xd7, 0x2b, 0xd2, 0x1c, 0x36, 0x47, 0xdd, 0xe9, 0x53, 0x6b, 0x82, 0x43, 0x67, 0xdb, 0x6a,
	0x94, 0x79, 0xc6, 0xb2, 0x01, 0x6e, 0xea, 0xb3, 0x10, 0x4a, 0xd9, 0xe9, 0x58, 0xad, 0x50, 0xb6,
	0x31, 0x95, 0xf7, 0xb8, 0xf5, 0x05, 0x5f, 0x0b, 0xd2, 0x06, 0xc4, 0xbb, 0x67, 0xab, 0x51, 0xe6,
	0x19, 0xf1, 0x39, 0x3a, 0x74, 0xbb, 0xcf, 0x22, 0xcb, 0x63, 0xe9, 0xe8, 0x7d, 0xa0, 0x9f, 0x5b,
	0x13, 0x10, 0x9f, 0xf6, 0x2c, 0x94, 0xfd, 0x8b, 0xe1, 0x8f, 0xa8, 0xbf, 0x73, 0x38, 0x0b, 0x49,
	0x07, 0x92, 0x06, 0xd6, 0x04, 0xc7, 0x77, 0x24, 0x95, 0x15, 0xfc, 0x8d, 0xd0, 0x33, 0xf4, 0xec,
	0x8e, 0x8f, 0x2d, 0x57, 0x32, 0xcd, 0x05, 0x34, 0xbb, 0x9e, 0x8a, 0xdf, 0xec, 0xb0, 0x6c, 0x76,
	0x48, 0x7f, 0x37, 0xd0, 0xf0, 0x5e, 0xf8, 0x4b, 0xc6, 0x95, 0x12, 0x19, 0x7e, 0x89, 0x5a, 0x4b,
	0x19, 0x0a, 0x48, 0x69, 0x4f, 0xfb, 0xd6, 0x04, 0x3d, 0x97, 0x52, 0x9e, 0x52, 0x06, 0x22, 0x7e,
	0x8f, 0x7a, 0xe5, 0xf3, 0xd3, 0x0f, 0x95, 0xf0, 0x38, 0xad, 0xc6, 0x7b, 0x6c, 0x4d, 0x80, 0xb7,
	0xde, 0x4a, 0xa4, 0xcc, 0xb7, 0xe2, 0xd7, 0xa8, 0x2d, 0xb2, 0x4c, 0x66, 0xf0, 0xdd, 0x76, 0xa7,
	0x8f, 0xad, 0x09, 0x1e, 0x3a, 0x06, 0x8e, 0x29, 0x73, 0x32, 0x3e, 0x47, 0xad, 0x90, 0x6b, 0x0e,
	0xd3, 0xed, 0x9d, 0xbe, 0xfa, 0xcf, 0xcf, 0xe6, 0x4a, 0xf0, 0xdf, 0xb6, 0x84, 0x29, 0x83, 0x8c,
	0xc5, 0x3e, 0xfc, 0xdf, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xae, 0x50, 0xe7, 0xc2, 0x32,
	0x04, 0x00, 0x00,
}
