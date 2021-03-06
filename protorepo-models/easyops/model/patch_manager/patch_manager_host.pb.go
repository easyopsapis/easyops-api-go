// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: patch_manager_host.proto

package patch_manager

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
//主机详情
type PatchManagerHost struct {
	//
	//实例 ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//主机名
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname" form:"hostname"`
	//
	//主机 IP
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//agent 状态
	XAgentStatus string `protobuf:"bytes,4,opt,name=_agentStatus,json=AgentStatus,proto3" json:"_agentStatus" form:"_agentStatus"`
	//
	//主机环境
	XEnvironment string `protobuf:"bytes,5,opt,name=_environment,json=Environment,proto3" json:"_environment" form:"_environment"`
	//
	//操作系统类型（如:Windows）
	OsSystem string `protobuf:"bytes,6,opt,name=osSystem,proto3" json:"osSystem" form:"osSystem"`
	//
	//操作系统架构
	OsArchitecture string `protobuf:"bytes,7,opt,name=osArchitecture,proto3" json:"osArchitecture" form:"osArchitecture"`
	//
	//操作系统版本（如:Windows 10）
	OsVersion string `protobuf:"bytes,8,opt,name=osVersion,proto3" json:"osVersion" form:"osVersion"`
	//
	//操作系统内核发行版本
	OsRelease string `protobuf:"bytes,9,opt,name=osRelease,proto3" json:"osRelease" form:"osRelease"`
	//
	//是否需要重启
	RequireReboot bool `protobuf:"varint,10,opt,name=requireReboot,proto3" json:"requireReboot" form:"requireReboot"`
	//
	//最近一次启动时间
	LastStartTime int32 `protobuf:"varint,11,opt,name=lastStartTime,proto3" json:"lastStartTime" form:"lastStartTime"`
	//
	//最近一次安装补丁时间
	LastInstallPatchTime int32 `protobuf:"varint,12,opt,name=lastInstallPatchTime,proto3" json:"lastInstallPatchTime" form:"lastInstallPatchTime"`
	//
	//已安装的补丁信息
	InstalledPatch       []*PatchManagerHost_InstalledPatch `protobuf:"bytes,13,rep,name=installedPatch,proto3" json:"installedPatch" form:"installedPatch"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *PatchManagerHost) Reset()         { *m = PatchManagerHost{} }
func (m *PatchManagerHost) String() string { return proto.CompactTextString(m) }
func (*PatchManagerHost) ProtoMessage()    {}
func (*PatchManagerHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_662f3875e1022beb, []int{0}
}
func (m *PatchManagerHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchManagerHost.Unmarshal(m, b)
}
func (m *PatchManagerHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchManagerHost.Marshal(b, m, deterministic)
}
func (m *PatchManagerHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchManagerHost.Merge(m, src)
}
func (m *PatchManagerHost) XXX_Size() int {
	return xxx_messageInfo_PatchManagerHost.Size(m)
}
func (m *PatchManagerHost) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchManagerHost.DiscardUnknown(m)
}

var xxx_messageInfo_PatchManagerHost proto.InternalMessageInfo

func (m *PatchManagerHost) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *PatchManagerHost) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *PatchManagerHost) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *PatchManagerHost) GetXAgentStatus() string {
	if m != nil {
		return m.XAgentStatus
	}
	return ""
}

func (m *PatchManagerHost) GetXEnvironment() string {
	if m != nil {
		return m.XEnvironment
	}
	return ""
}

func (m *PatchManagerHost) GetOsSystem() string {
	if m != nil {
		return m.OsSystem
	}
	return ""
}

func (m *PatchManagerHost) GetOsArchitecture() string {
	if m != nil {
		return m.OsArchitecture
	}
	return ""
}

func (m *PatchManagerHost) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *PatchManagerHost) GetOsRelease() string {
	if m != nil {
		return m.OsRelease
	}
	return ""
}

func (m *PatchManagerHost) GetRequireReboot() bool {
	if m != nil {
		return m.RequireReboot
	}
	return false
}

func (m *PatchManagerHost) GetLastStartTime() int32 {
	if m != nil {
		return m.LastStartTime
	}
	return 0
}

func (m *PatchManagerHost) GetLastInstallPatchTime() int32 {
	if m != nil {
		return m.LastInstallPatchTime
	}
	return 0
}

func (m *PatchManagerHost) GetInstalledPatch() []*PatchManagerHost_InstalledPatch {
	if m != nil {
		return m.InstalledPatch
	}
	return nil
}

type PatchManagerHost_InstalledPatch struct {
	//
	//补丁编号
	ArticleId string `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId" form:"articleId"`
	//
	//补丁安装时间
	InstalledTime        string   `protobuf:"bytes,2,opt,name=installedTime,proto3" json:"installedTime" form:"installedTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchManagerHost_InstalledPatch) Reset()         { *m = PatchManagerHost_InstalledPatch{} }
func (m *PatchManagerHost_InstalledPatch) String() string { return proto.CompactTextString(m) }
func (*PatchManagerHost_InstalledPatch) ProtoMessage()    {}
func (*PatchManagerHost_InstalledPatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_662f3875e1022beb, []int{0, 0}
}
func (m *PatchManagerHost_InstalledPatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchManagerHost_InstalledPatch.Unmarshal(m, b)
}
func (m *PatchManagerHost_InstalledPatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchManagerHost_InstalledPatch.Marshal(b, m, deterministic)
}
func (m *PatchManagerHost_InstalledPatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchManagerHost_InstalledPatch.Merge(m, src)
}
func (m *PatchManagerHost_InstalledPatch) XXX_Size() int {
	return xxx_messageInfo_PatchManagerHost_InstalledPatch.Size(m)
}
func (m *PatchManagerHost_InstalledPatch) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchManagerHost_InstalledPatch.DiscardUnknown(m)
}

var xxx_messageInfo_PatchManagerHost_InstalledPatch proto.InternalMessageInfo

func (m *PatchManagerHost_InstalledPatch) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

func (m *PatchManagerHost_InstalledPatch) GetInstalledTime() string {
	if m != nil {
		return m.InstalledTime
	}
	return ""
}

func init() {
	proto.RegisterType((*PatchManagerHost)(nil), "patch_manager.PatchManagerHost")
	proto.RegisterType((*PatchManagerHost_InstalledPatch)(nil), "patch_manager.PatchManagerHost.InstalledPatch")
}

func init() { proto.RegisterFile("patch_manager_host.proto", fileDescriptor_662f3875e1022beb) }

var fileDescriptor_662f3875e1022beb = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x4d, 0x6f, 0xe3, 0x44,
	0x1c, 0xc6, 0x95, 0x2e, 0xbb, 0xd4, 0xd3, 0x6d, 0x29, 0xde, 0x22, 0x4c, 0x39, 0xb8, 0xb2, 0x16,
	0x69, 0xbc, 0xd4, 0x76, 0xed, 0x24, 0x5d, 0x6d, 0x0e, 0x54, 0xad, 0xc4, 0x4b, 0x0e, 0x48, 0x68,
	0x8a, 0x38, 0xd0, 0x64, 0xab, 0x89, 0x33, 0x4d, 0x2c, 0x6c, 0x4f, 0xd6, 0x33, 0xd9, 0xd5, 0x12,
	0xfb, 0xc6, 0x1d, 0x89, 0xcf, 0xc0, 0x57, 0xe0, 0xeb, 0x18, 0x89, 0x1b, 0x57, 0x7f, 0x02, 0xe4,
	0x71, 0x5e, 0xec, 0xa4, 0x06, 0x0e, 0xdd, 0xde, 0x66, 0xe6, 0x79, 0x9e, 0xff, 0xff, 0xe7, 0x19,
	0x69, 0xc6, 0x40, 0x99, 0x60, 0xee, 0x8e, 0xaf, 0x03, 0x1c, 0xe2, 0x11, 0x89, 0xae, 0xc7, 0x94,
	0x71, 0x73, 0x12, 0x51, 0x4e, 0xe5, 0xdd, 0x8a, 0x72, 0x68, 0x8c, 0x3c, 0x3e, 0x9e, 0x0e, 0x4c,
	0x97, 0x06, 0xd6, 0x88, 0x8e, 0xa8, 0x25, 0x5c, 0x83, 0xe9, 0x8d, 0x98, 0x89, 0x89, 0x18, 0x15,
	0xe9, 0xc3, 0xd3, 0x92, 0x3d, 0x78, 0xe3, 0xf1, 0x9f, 0xe8, 0x1b, 0x6b, 0x44, 0x0d, 0x21, 0x1a,
	0xaf, 0xb1, 0xef, 0x0d, 0x31, 0xa7, 0x11, 0xb3, 0x96, 0xc3, 0x22, 0xa7, 0xfd, 0xbd, 0x0f, 0xf6,
	0xbf, 0xcb, 0x1b, 0x7f, 0x5b, 0xf4, 0xfd, 0x86, 0x32, 0x2e, 0x77, 0x01, 0xf0, 0x42, 0xc6, 0x71,
	0xe8, 0x92, 0xee, 0x50, 0x69, 0x1c, 0x35, 0xa0, 0x74, 0xa1, 0x67, 0xa9, 0xfa, 0xe1, 0x0d, 0x8d,
	0x82, 0x8e, 0xb6, 0xd2, 0xb4, 0xbf, 0xfe, 0x54, 0xf7, 0xc1, 0xde, 0xcb, 0xab, 0x13, 0xe3, 0x05,
	0x36, 0x7e, 0xee, 0xcf, 0xec, 0x66, 0xf2, 0x14, 0x95, 0xc2, 0xb2, 0x05, 0xb6, 0xf3, 0x6f, 0x0c,
	0x71, 0x40, 0x94, 0x2d, 0x51, 0xe8, 0x49, 0x96, 0xaa, 0x1f, 0x14, 0x85, 0x16, 0x8a, 0x86, 0x96,
	0x26, 0xf9, 0x0f, 0x09, 0x6c, 0x79, 0x13, 0xe5, 0x81, 0xf0, 0xfe, 0x2e, 0x65, 0xa9, 0x2a, 0xcd,
	0xbb, 0x4e, 0xf2, 0x6e, 0xbf, 0x49, 0xe0, 0x57, 0xe9, 0x25, 0x84, 0x0e, 0x6c, 0x5f, 0x9d, 0x18,
	0xed, 0xfe, 0xcc, 0x4e, 0xe2, 0xab, 0x13, 0xa3, 0xd5, 0xef, 0x0d, 0x67, 0x76, 0xa2, 0xe7, 0x63,
	0xbb, 0x7f, 0x96, 0x4f, 0x8e, 0x9d, 0x44, 0x87, 0x3d, 0xf3, 0x7f, 0x3a, 0xf5, 0x59, 0x33, 0xd1,
	0xe3, 0x1e, 0x7b, 0x06, 0x21, 0xcc, 0xbf, 0xe6, 0xdc, 0xf8, 0x0a, 0x1b, 0x37, 0xfd, 0x99, 0x7d,
	0xdc, 0x4a, 0x3a, 0xfa, 0xec, 0x79, 0xb2, 0xb1, 0x1a, 0x77, 0x74, 0x3d, 0xbe, 0xd5, 0x7c, 0x9a,
	0xc0, 0xce, 0x86, 0x1b, 0x42, 0xa7, 0xe0, 0x88, 0x9d, 0x39, 0x45, 0x6c, 0xf7, 0x86, 0xbd, 0x61,
	0x7c, 0x65, 0x1b, 0x2f, 0x72, 0x8e, 0x02, 0xf6, 0x3f, 0x3c, 0x05, 0x66, 0x6d, 0xe7, 0x76, 0x02,
	0xe1, 0x66, 0x6f, 0xbd, 0xf8, 0xc4, 0xb8, 0x73, 0x2f, 0x0c, 0xad, 0x5a, 0x86, 0x3c, 0x76, 0x9b,
	0x74, 0x76, 0x97, 0x60, 0xff, 0x42, 0xd6, 0xac, 0x25, 0x6b, 0xd5, 0x90, 0xcd, 0x4e, 0x8e, 0x9d,
	0xe4, 0x9e, 0xe8, 0x9c, 0x5a, 0xba, 0x76, 0x3d, 0x5d, 0xf3, 0xbe, 0xe8, 0xec, 0x5a, 0xba, 0xd3,
	0x7a, 0xba, 0xd6, 0xbb, 0xa0, 0xeb, 0xd4, 0x81, 0x3c, 0xaf, 0x07, 0x69, 0xdf, 0x3d, 0x88, 0x0e,
	0x3f, 0x33, 0x3f, 0xd7, 0xcf, 0x7a, 0xec, 0xd9, 0x53, 0xb4, 0xe5, 0x4d, 0xe4, 0x0e, 0x78, 0x7c,
	0x8d, 0x47, 0x24, 0xe4, 0x97, 0x1c, 0xf3, 0x29, 0x53, 0xde, 0x13, 0x17, 0xd8, 0xc7, 0x59, 0xaa,
	0x3e, 0x29, 0xee, 0xaf, 0xb2, 0xaa, 0xa1, 0x9d, 0xf3, 0xd5, 0x4c, 0x64, 0x49, 0xf8, 0xda, 0x8b,
	0x68, 0x18, 0x90, 0x90, 0x2b, 0x0f, 0x37, 0xb2, 0x25, 0x55, 0x43, 0x3b, 0x5f, 0xae, 0x66, 0xf9,
	0x05, 0x4b, 0xd9, 0xe5, 0x5b, 0xc6, 0x49, 0xa0, 0x3c, 0x5a, 0xbf, 0x60, 0x17, 0x8a, 0x86, 0x96,
	0x26, 0xf9, 0x1c, 0xec, 0x51, 0x76, 0x1e, 0xb9, 0x63, 0x8f, 0x13, 0x97, 0x4f, 0x23, 0xa2, 0xbc,
	0x2f, 0x62, 0x9f, 0x64, 0xa9, 0xfa, 0xd1, 0x22, 0x56, 0xd6, 0x35, 0xb4, 0x16, 0x90, 0x1d, 0x20,
	0x51, 0xf6, 0x03, 0x89, 0x98, 0x47, 0x43, 0x65, 0x5b, 0xa4, 0x0f, 0xb2, 0x54, 0xdd, 0x5f, 0xa4,
	0xe7, 0x92, 0x86, 0x56, 0xb6, 0x22, 0x83, 0x88, 0x4f, 0x30, 0x23, 0x8a, 0xb4, 0x99, 0x99, 0x4b,
	0x22, 0x33, 0x1f, 0xcb, 0x5f, 0x80, 0xdd, 0x88, 0xbc, 0x9a, 0x7a, 0x11, 0x41, 0x64, 0x40, 0x29,
	0x57, 0xc0, 0x51, 0x03, 0x6e, 0x5f, 0x28, 0x59, 0xaa, 0x1e, 0x14, 0xb9, 0x8a, 0xac, 0xa1, 0xaa,
	0x3d, 0xcf, 0xfb, 0x98, 0xe5, 0xbb, 0x1c, 0xf1, 0xef, 0xbd, 0x80, 0x28, 0x3b, 0x47, 0x0d, 0xf8,
	0xb0, 0x9c, 0xaf, 0xc8, 0x1a, 0xaa, 0xda, 0xe5, 0x4b, 0x70, 0x90, 0x2f, 0x74, 0xf3, 0xe7, 0xcc,
	0xf7, 0xc5, 0x33, 0x29, 0xca, 0x3c, 0x16, 0x65, 0xd4, 0x2c, 0x55, 0x3f, 0x5d, 0x95, 0x59, 0x77,
	0x69, 0xe8, 0xd6, 0xb0, 0xfc, 0x0a, 0xec, 0x79, 0xc5, 0x1a, 0x19, 0x8a, 0x55, 0x65, 0xf7, 0xe8,
	0x01, 0xdc, 0x71, 0x4c, 0xb3, 0xf2, 0x03, 0x60, 0xae, 0xbf, 0xca, 0x66, 0xb7, 0x92, 0x2a, 0x9f,
	0x57, 0xb5, 0x9e, 0x86, 0xd6, 0x1a, 0x1c, 0xfe, 0xd2, 0x00, 0x7b, 0xd5, 0x74, 0x7e, 0x1c, 0x38,
	0xe2, 0x9e, 0xeb, 0xaf, 0x5e, 0xf8, 0xd2, 0x71, 0x2c, 0x25, 0x0d, 0xad, 0x6c, 0xf9, 0x76, 0x2e,
	0x0b, 0x8b, 0x7d, 0x28, 0x1e, 0xf4, 0xd2, 0x76, 0x56, 0x64, 0x0d, 0x55, 0xed, 0x17, 0xdd, 0x1f,
	0xbf, 0x1e, 0x51, 0x93, 0x60, 0xf6, 0x96, 0x4e, 0x98, 0xe9, 0x53, 0x17, 0xfb, 0x96, 0x4b, 0x43,
	0x1e, 0x61, 0x97, 0xb3, 0xe2, 0xff, 0x26, 0x22, 0x13, 0x6a, 0x04, 0x74, 0x48, 0x7c, 0x66, 0xcd,
	0x8d, 0x96, 0x98, 0x5a, 0x95, 0xcd, 0x19, 0x3c, 0x12, 0xee, 0xe6, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xde, 0x6a, 0x92, 0xa6, 0x4f, 0x09, 0x00, 0x00,
}
