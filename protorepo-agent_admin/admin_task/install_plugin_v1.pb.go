// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: install_plugin_v1.proto

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
//InstallPluginTaskV1请求
type InstallPluginTaskV1Request struct {
	//
	//插件部署任务的详情数据
	DeployInfoList []*InstallPluginTaskV1Request_DeployInfoList `protobuf:"bytes,1,rep,name=deployInfoList,proto3" json:"deployInfoList" form:"deployInfoList"`
	//
	//部署类型
	DeployType string `protobuf:"bytes,2,opt,name=deployType,proto3" json:"deployType" form:"deployType"`
	//
	//每批次数量
	PerBatchNum int32 `protobuf:"varint,3,opt,name=perBatchNum,proto3" json:"perBatchNum" form:"perBatchNum"`
	//
	//每批时间间隔
	PerBatchInterval     int32    `protobuf:"varint,4,opt,name=perBatchInterval,proto3" json:"perBatchInterval" form:"perBatchInterval"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskV1Request) Reset()         { *m = InstallPluginTaskV1Request{} }
func (m *InstallPluginTaskV1Request) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskV1Request) ProtoMessage()    {}
func (*InstallPluginTaskV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7022741c94db14f, []int{0}
}
func (m *InstallPluginTaskV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskV1Request.Unmarshal(m, b)
}
func (m *InstallPluginTaskV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskV1Request.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskV1Request.Merge(m, src)
}
func (m *InstallPluginTaskV1Request) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskV1Request.Size(m)
}
func (m *InstallPluginTaskV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskV1Request proto.InternalMessageInfo

func (m *InstallPluginTaskV1Request) GetDeployInfoList() []*InstallPluginTaskV1Request_DeployInfoList {
	if m != nil {
		return m.DeployInfoList
	}
	return nil
}

func (m *InstallPluginTaskV1Request) GetDeployType() string {
	if m != nil {
		return m.DeployType
	}
	return ""
}

func (m *InstallPluginTaskV1Request) GetPerBatchNum() int32 {
	if m != nil {
		return m.PerBatchNum
	}
	return 0
}

func (m *InstallPluginTaskV1Request) GetPerBatchInterval() int32 {
	if m != nil {
		return m.PerBatchInterval
	}
	return 0
}

type InstallPluginTaskV1Request_DeployInfoList struct {
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
	AgentPluginId string `protobuf:"bytes,4,opt,name=agentPluginId,proto3" json:"agentPluginId" form:"agentPluginId"`
	//
	//插件包名称
	AgentPluginName string `protobuf:"bytes,5,opt,name=agentPluginName,proto3" json:"agentPluginName" form:"agentPluginName"`
	//
	//插件包部署前名称
	PreVersionName string `protobuf:"bytes,6,opt,name=preVersionName,proto3" json:"preVersionName" form:"preVersionName"`
	//
	//插件包版本名称
	PluginVersionName string `protobuf:"bytes,7,opt,name=pluginVersionName,proto3" json:"pluginVersionName" form:"pluginVersionName"`
	//
	//插件包版本仓库版本ID
	RepoVersionId        string   `protobuf:"bytes,8,opt,name=repoVersionId,proto3" json:"repoVersionId" form:"repoVersionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskV1Request_DeployInfoList) Reset() {
	*m = InstallPluginTaskV1Request_DeployInfoList{}
}
func (m *InstallPluginTaskV1Request_DeployInfoList) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskV1Request_DeployInfoList) ProtoMessage()    {}
func (*InstallPluginTaskV1Request_DeployInfoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7022741c94db14f, []int{0, 0}
}
func (m *InstallPluginTaskV1Request_DeployInfoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList.Unmarshal(m, b)
}
func (m *InstallPluginTaskV1Request_DeployInfoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskV1Request_DeployInfoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList.Merge(m, src)
}
func (m *InstallPluginTaskV1Request_DeployInfoList) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList.Size(m)
}
func (m *InstallPluginTaskV1Request_DeployInfoList) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskV1Request_DeployInfoList proto.InternalMessageInfo

func (m *InstallPluginTaskV1Request_DeployInfoList) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetAgentPluginId() string {
	if m != nil {
		return m.AgentPluginId
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetAgentPluginName() string {
	if m != nil {
		return m.AgentPluginName
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetPreVersionName() string {
	if m != nil {
		return m.PreVersionName
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetPluginVersionName() string {
	if m != nil {
		return m.PluginVersionName
	}
	return ""
}

func (m *InstallPluginTaskV1Request_DeployInfoList) GetRepoVersionId() string {
	if m != nil {
		return m.RepoVersionId
	}
	return ""
}

//
//InstallPluginTaskV1返回
type InstallPluginTaskV1Response struct {
	//
	//命令任务ID
	CmdTaskId            string   `protobuf:"bytes,1,opt,name=cmdTaskId,proto3" json:"cmdTaskId" form:"cmdTaskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallPluginTaskV1Response) Reset()         { *m = InstallPluginTaskV1Response{} }
func (m *InstallPluginTaskV1Response) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskV1Response) ProtoMessage()    {}
func (*InstallPluginTaskV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7022741c94db14f, []int{1}
}
func (m *InstallPluginTaskV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskV1Response.Unmarshal(m, b)
}
func (m *InstallPluginTaskV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskV1Response.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskV1Response.Merge(m, src)
}
func (m *InstallPluginTaskV1Response) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskV1Response.Size(m)
}
func (m *InstallPluginTaskV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskV1Response proto.InternalMessageInfo

func (m *InstallPluginTaskV1Response) GetCmdTaskId() string {
	if m != nil {
		return m.CmdTaskId
	}
	return ""
}

//
//InstallPluginTaskV1Api返回
type InstallPluginTaskV1ResponseWrapper struct {
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
	Data                 *InstallPluginTaskV1Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *InstallPluginTaskV1ResponseWrapper) Reset()         { *m = InstallPluginTaskV1ResponseWrapper{} }
func (m *InstallPluginTaskV1ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InstallPluginTaskV1ResponseWrapper) ProtoMessage()    {}
func (*InstallPluginTaskV1ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7022741c94db14f, []int{2}
}
func (m *InstallPluginTaskV1ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallPluginTaskV1ResponseWrapper.Unmarshal(m, b)
}
func (m *InstallPluginTaskV1ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallPluginTaskV1ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InstallPluginTaskV1ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallPluginTaskV1ResponseWrapper.Merge(m, src)
}
func (m *InstallPluginTaskV1ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InstallPluginTaskV1ResponseWrapper.Size(m)
}
func (m *InstallPluginTaskV1ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallPluginTaskV1ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InstallPluginTaskV1ResponseWrapper proto.InternalMessageInfo

func (m *InstallPluginTaskV1ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InstallPluginTaskV1ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InstallPluginTaskV1ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InstallPluginTaskV1ResponseWrapper) GetData() *InstallPluginTaskV1Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallPluginTaskV1Request)(nil), "admin_task.InstallPluginTaskV1Request")
	proto.RegisterType((*InstallPluginTaskV1Request_DeployInfoList)(nil), "admin_task.InstallPluginTaskV1Request.DeployInfoList")
	proto.RegisterType((*InstallPluginTaskV1Response)(nil), "admin_task.InstallPluginTaskV1Response")
	proto.RegisterType((*InstallPluginTaskV1ResponseWrapper)(nil), "admin_task.InstallPluginTaskV1ResponseWrapper")
}

func init() { proto.RegisterFile("install_plugin_v1.proto", fileDescriptor_c7022741c94db14f) }

var fileDescriptor_c7022741c94db14f = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd5, 0x75, 0xdd, 0xa8, 0xcb, 0xba, 0xcd, 0xda, 0x9f, 0xd0, 0x81, 0x5c, 0x19, 0x09,
	0x76, 0x80, 0x4e, 0x1b, 0x9a, 0x84, 0x38, 0x20, 0x11, 0x0d, 0xa1, 0xa0, 0x69, 0x1a, 0xd6, 0x34,
	0x8e, 0x93, 0xdb, 0xb8, 0x6d, 0xb4, 0x24, 0x36, 0x4e, 0x3a, 0x6d, 0xdf, 0x8d, 0xcf, 0xe2, 0x3b,
	0x57, 0xdf, 0xb8, 0xa1, 0xd8, 0x69, 0xeb, 0xb4, 0x03, 0x4e, 0xad, 0xdf, 0xe7, 0x79, 0x7e, 0xce,
	0xfb, 0xc6, 0x0e, 0xd8, 0x8f, 0xd2, 0x2c, 0xa7, 0x71, 0x7c, 0x23, 0xe2, 0xc9, 0x28, 0x4a, 0x6f,
	0xee, 0x8e, 0x7b, 0x42, 0xf2, 0x9c, 0x43, 0x40, 0xc3, 0x24, 0x4a, 0x6f, 0x72, 0x9a, 0xdd, 0x76,
	0xde, 0x8e, 0xa2, 0x7c, 0x3c, 0xe9, 0xf7, 0x06, 0x3c, 0x39, 0x1a, 0xf1, 0x11, 0x3f, 0x32, 0x96,
	0xfe, 0x64, 0x68, 0x56, 0x66, 0x61, 0xfe, 0xd9, 0x28, 0xfe, 0xb9, 0x06, 0x3a, 0x81, 0xc5, 0x5e,
	0x1a, 0xea, 0x15, 0xcd, 0x6e, 0xaf, 0x8f, 0x09, 0xfb, 0x31, 0x61, 0x59, 0x0e, 0xef, 0x41, 0x3b,
	0x64, 0x22, 0xe6, 0x0f, 0x41, 0x3a, 0xe4, 0xe7, 0x51, 0x96, 0x7b, 0xb5, 0x6e, 0xfd, 0xb0, 0x75,
	0x72, 0xda, 0x9b, 0x6f, 0xd9, 0xfb, 0x7b, 0xbe, 0x77, 0x56, 0x09, 0xfb, 0xcf, 0xb4, 0x42, 0xbb,
	0x43, 0x2e, 0x93, 0x0f, 0xb8, 0x8a, 0xc5, 0x64, 0x61, 0x1f, 0x78, 0x0a, 0x80, 0xad, 0x5c, 0x3d,
	0x08, 0xe6, 0xad, 0x74, 0x6b, 0x87, 0x4d, 0x7f, 0x57, 0x2b, 0xb4, 0xed, 0xc6, 0x0b, 0x0d, 0x13,
	0xc7, 0x08, 0xdf, 0x83, 0x96, 0x60, 0xd2, 0xa7, 0xf9, 0x60, 0x7c, 0x31, 0x49, 0xbc, 0x7a, 0xb7,
	0x76, 0xd8, 0xf0, 0xf7, 0xb4, 0x42, 0xd0, 0xe6, 0x1c, 0x11, 0x13, 0xd7, 0x0a, 0xbf, 0x80, 0xad,
	0xe9, 0x32, 0x48, 0x73, 0x26, 0xef, 0x68, 0xec, 0xad, 0x9a, 0xf8, 0x81, 0x56, 0x68, 0xbf, 0x1a,
	0x9f, 0x3a, 0x30, 0x59, 0x0a, 0x75, 0x7e, 0xd5, 0x41, 0xbb, 0xda, 0x37, 0x7c, 0x03, 0xd6, 0xe9,
	0x88, 0xa5, 0x79, 0x10, 0x7a, 0x35, 0xd3, 0x09, 0xd4, 0x0a, 0xb5, 0x2d, 0xb2, 0x14, 0x30, 0x99,
	0x5a, 0xe0, 0x0b, 0xb0, 0x12, 0x89, 0xb2, 0xe5, 0x0d, 0xad, 0x50, 0xd3, 0x1a, 0x23, 0x81, 0xc9,
	0x4a, 0x24, 0xe6, 0x93, 0xb9, 0xa4, 0xf9, 0xd8, 0xab, 0x77, 0xeb, 0x8f, 0x4d, 0xa6, 0xd0, 0x66,
	0x93, 0x29, 0x16, 0xf0, 0x23, 0xd8, 0x30, 0x1b, 0xd8, 0xd7, 0x14, 0x84, 0xa6, 0xb9, 0xa6, 0xef,
	0x69, 0x85, 0x76, 0x9c, 0x27, 0x99, 0xca, 0x98, 0x54, 0xed, 0xf0, 0x0c, 0x6c, 0x3a, 0x85, 0x0b,
	0x9a, 0x30, 0xaf, 0x61, 0x08, 0x1d, 0xad, 0xd0, 0xde, 0x12, 0xa1, 0x30, 0x60, 0xb2, 0x18, 0x81,
	0x9f, 0x40, 0x5b, 0x48, 0x76, 0xcd, 0x64, 0x16, 0x71, 0x0b, 0x59, 0x33, 0x10, 0xe7, 0x64, 0x54,
	0x75, 0x4c, 0x16, 0x02, 0xf0, 0x2b, 0xd8, 0xb6, 0x17, 0xc0, 0xa5, 0xac, 0x1b, 0xca, 0x73, 0xad,
	0x90, 0x57, 0x52, 0x16, 0x2d, 0x98, 0x2c, 0xc7, 0x8a, 0xa1, 0x48, 0x26, 0x78, 0x59, 0x0a, 0x42,
	0xef, 0xc9, 0xe2, 0x50, 0x2a, 0x32, 0x26, 0x55, 0x3b, 0xfe, 0x06, 0x0e, 0x1e, 0x3d, 0xfd, 0x99,
	0xe0, 0x69, 0xc6, 0xe0, 0x09, 0x68, 0x0e, 0x92, 0xb0, 0x28, 0xce, 0xde, 0xfc, 0x8e, 0x56, 0x68,
	0xcb, 0xa2, 0x67, 0x12, 0x26, 0x73, 0x1b, 0xfe, 0x5d, 0x03, 0xf8, 0x1f, 0xcc, 0xef, 0x92, 0x0a,
	0xc1, 0x24, 0x7c, 0x09, 0x56, 0x07, 0x3c, 0x64, 0x86, 0xda, 0xf0, 0x37, 0xb5, 0x42, 0xad, 0x92,
	0xca, 0x43, 0x86, 0x89, 0x11, 0x8b, 0xdb, 0x50, 0xfc, 0x7e, 0xbe, 0x17, 0x31, 0x8d, 0xd2, 0xf2,
	0x48, 0x39, 0xb7, 0xc1, 0x11, 0x31, 0x71, 0xad, 0xf0, 0x15, 0x68, 0x30, 0x29, 0xb9, 0x34, 0x37,
	0xa8, 0xe9, 0x6f, 0x69, 0x85, 0x9e, 0xda, 0x8c, 0x29, 0x63, 0x62, 0x65, 0x78, 0x0e, 0x56, 0x43,
	0x9a, 0x53, 0x73, 0x98, 0x5a, 0x27, 0xaf, 0xff, 0xfb, 0x59, 0xb0, 0x4d, 0xb8, 0xcf, 0x5b, 0xc4,
	0x31, 0x31, 0x94, 0xfe, 0x9a, 0xf9, 0x28, 0xbd, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x06, 0xbd,
	0xcf, 0xcc, 0xea, 0x04, 0x00, 0x00,
}
