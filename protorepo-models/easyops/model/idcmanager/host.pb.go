// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: host.proto

package idcmanager

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
//主机
type Host struct {
	//
	//主机ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//主机名称
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname" form:"hostname"`
	//
	//主机IP
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//agent版本
	AgentVersion string `protobuf:"bytes,4,opt,name=agentVersion,proto3" json:"agentVersion" form:"agentVersion"`
	//
	//agent状态
	XAgentStatus string `protobuf:"bytes,5,opt,name=_agentStatus,json=AgentStatus,proto3" json:"_agentStatus" form:"_agentStatus"`
	//
	//agent心跳
	XAgentHeartBeat int32 `protobuf:"varint,6,opt,name=_agentHeartBeat,json=AgentHeartBeat,proto3" json:"_agentHeartBeat" form:"_agentHeartBeat"`
	//
	//物理地址
	XMac string `protobuf:"bytes,7,opt,name=_mac,json=Mac,proto3" json:"_mac" form:"_mac"`
	//
	//主机环境
	XEnvironment string `protobuf:"bytes,8,opt,name=_environment,json=Environment,proto3" json:"_environment" form:"_environment"`
	//
	//起始U位
	XStartU int32 `protobuf:"varint,9,opt,name=_startU,json=StartU,proto3" json:"_startU" form:"_startU"`
	//
	//占用U位
	XOccupiedU int32 `protobuf:"varint,10,opt,name=_occupiedU,json=OccupiedU,proto3" json:"_occupiedU" form:"_occupiedU"`
	//
	//供应商
	Provider string `protobuf:"bytes,11,opt,name=provider,proto3" json:"provider" form:"provider"`
	//
	//备注
	Memo string `protobuf:"bytes,12,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//内存大小
	MemSize int32 `protobuf:"varint,13,opt,name=memSize,proto3" json:"memSize" form:"memSize"`
	//
	//磁盘大小
	DiskSize int32 `protobuf:"varint,14,opt,name=diskSize,proto3" json:"diskSize" form:"diskSize"`
	//
	//操作系统类型
	OsSystem string `protobuf:"bytes,15,opt,name=osSystem,proto3" json:"osSystem" form:"osSystem"`
	//
	//操作系统架构
	OsArchitecture string `protobuf:"bytes,16,opt,name=osArchitecture,proto3" json:"osArchitecture" form:"osArchitecture"`
	//
	//操作系统
	OsVersion string `protobuf:"bytes,17,opt,name=osVersion,proto3" json:"osVersion" form:"osVersion"`
	//
	//操作系统内核发行版本
	OsRelease string `protobuf:"bytes,18,opt,name=osRelease,proto3" json:"osRelease" form:"osRelease"`
	//
	//创建时间
	Ctime string `protobuf:"bytes,19,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//创建人
	Creator string `protobuf:"bytes,20,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//产品线
	Product string `protobuf:"bytes,21,opt,name=product,proto3" json:"product" form:"product"`
	//
	//客户
	Customer string `protobuf:"bytes,22,opt,name=customer,proto3" json:"customer" form:"customer"`
	//
	//是否单电源
	IsSinglePower        bool     `protobuf:"varint,23,opt,name=isSinglePower,proto3" json:"isSinglePower" form:"isSinglePower"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{0}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Host) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Host) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Host) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

func (m *Host) GetXAgentStatus() string {
	if m != nil {
		return m.XAgentStatus
	}
	return ""
}

func (m *Host) GetXAgentHeartBeat() int32 {
	if m != nil {
		return m.XAgentHeartBeat
	}
	return 0
}

func (m *Host) GetXMac() string {
	if m != nil {
		return m.XMac
	}
	return ""
}

func (m *Host) GetXEnvironment() string {
	if m != nil {
		return m.XEnvironment
	}
	return ""
}

func (m *Host) GetXStartU() int32 {
	if m != nil {
		return m.XStartU
	}
	return 0
}

func (m *Host) GetXOccupiedU() int32 {
	if m != nil {
		return m.XOccupiedU
	}
	return 0
}

func (m *Host) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Host) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Host) GetMemSize() int32 {
	if m != nil {
		return m.MemSize
	}
	return 0
}

func (m *Host) GetDiskSize() int32 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

func (m *Host) GetOsSystem() string {
	if m != nil {
		return m.OsSystem
	}
	return ""
}

func (m *Host) GetOsArchitecture() string {
	if m != nil {
		return m.OsArchitecture
	}
	return ""
}

func (m *Host) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *Host) GetOsRelease() string {
	if m != nil {
		return m.OsRelease
	}
	return ""
}

func (m *Host) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Host) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Host) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Host) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *Host) GetIsSinglePower() bool {
	if m != nil {
		return m.IsSinglePower
	}
	return false
}

func init() {
	proto.RegisterType((*Host)(nil), "idcmanager.Host")
}

func init() { proto.RegisterFile("host.proto", fileDescriptor_85e40b83b4d50a8d) }

var fileDescriptor_85e40b83b4d50a8d = []byte{
	// 1053 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x4f, 0x4f, 0xe3, 0xc6,
	0x1f, 0xc6, 0x95, 0x5d, 0x16, 0xc8, 0xc0, 0x02, 0x6b, 0x60, 0x77, 0x7e, 0x5c, 0xcc, 0xcf, 0x4b,
	0xa5, 0x31, 0xe0, 0xfc, 0x87, 0x15, 0xa9, 0xd4, 0x28, 0xb4, 0x5d, 0xed, 0x1e, 0xaa, 0x56, 0x93,
	0x6e, 0x0f, 0x24, 0x06, 0x19, 0x7b, 0x08, 0xd6, 0xc6, 0x9e, 0xc8, 0x33, 0x01, 0x2d, 0x89, 0x5f,
	0x43, 0xa5, 0x9e, 0x7b, 0xec, 0xa9, 0x52, 0x8f, 0x7d, 0x03, 0x7d, 0x17, 0xbd, 0xb8, 0x52, 0x5f,
	0x82, 0x5f, 0x41, 0xe5, 0xb1, 0x9d, 0x98, 0x40, 0xda, 0x1e, 0xb6, 0x9c, 0xf2, 0xfd, 0xce, 0xf3,
	0x3c, 0x33, 0x9f, 0x91, 0x66, 0x32, 0x06, 0xe0, 0x92, 0x32, 0x5e, 0xe8, 0x7b, 0x94, 0x53, 0x09,
	0xd8, 0x96, 0xe9, 0x18, 0xae, 0xd1, 0x25, 0xde, 0x96, 0xd6, 0xb5, 0xf9, 0xe5, 0xe0, 0xbc, 0x60,
	0x52, 0xa7, 0xd8, 0xa5, 0x5d, 0x5a, 0x14, 0x96, 0xf3, 0xc1, 0x85, 0xe8, 0x44, 0x23, 0xaa, 0x38,
	0xba, 0x75, 0x98, 0xb1, 0x3b, 0xd7, 0x36, 0x7f, 0x4f, 0xaf, 0x8b, 0x5d, 0xaa, 0x09, 0x51, 0xbb,
	0x32, 0x7a, 0xb6, 0x65, 0x70, 0xea, 0xb1, 0xe2, 0xb8, 0x8c, 0x73, 0xca, 0x8f, 0x1b, 0x60, 0xee,
	0x0d, 0x65, 0x5c, 0x7a, 0x0b, 0x80, 0xed, 0x32, 0x6e, 0xb8, 0x26, 0x79, 0x6b, 0xc1, 0xdc, 0x76,
	0x0e, 0xe5, 0x8f, 0xd5, 0x30, 0x90, 0x9f, 0x5d, 0x50, 0xcf, 0xa9, 0x2b, 0x13, 0x4d, 0xf9, 0xf3,
	0x0f, 0x79, 0x0d, 0xac, 0x9c, 0xb6, 0x4b, 0xda, 0x91, 0xa1, 0xdd, 0xe8, 0xc3, 0x72, 0xd5, 0xdf,
	0xc1, 0x99, 0xb0, 0x54, 0x04, 0x8b, 0xd1, 0xa6, 0x5c, 0xc3, 0x21, 0xf0, 0x91, 0x98, 0x68, 0x3d,
	0x0c, 0xe4, 0xd5, 0x78, 0xa2, 0x54, 0x51, 0xf0, 0xd8, 0x24, 0xfd, 0x9a, 0x07, 0x8f, 0xec, 0x3e,
	0x7c, 0x2c, 0xbc, 0x3f, 0xe5, 0xc3, 0x40, 0xce, 0x27, 0xab, 0xf6, 0xa3, 0xd5, 0x7e, 0xc8, 0x83,
	0xef, 0xf3, 0xa7, 0x08, 0x55, 0xd0, 0x41, 0xbb, 0xa4, 0x1d, 0xe8, 0xc3, 0xb2, 0x3f, 0x6a, 0x97,
	0xb4, 0x9a, 0xde, 0xb1, 0x86, 0x65, 0x5f, 0x8d, 0xea, 0xb2, 0xde, 0x88, 0x9a, 0xfd, 0x8a, 0xaf,
	0xa2, 0x4e, 0xe1, 0x5f, 0x3a, 0xd5, 0x61, 0xd5, 0x57, 0x47, 0x1d, 0xb6, 0x8b, 0x10, 0x8a, 0x76,
	0xd3, 0xd4, 0x5e, 0x1b, 0xda, 0x85, 0x3e, 0x2c, 0xef, 0xd7, 0xfc, 0xba, 0x3a, 0x7c, 0xe5, 0xdf,
	0x19, 0x1d, 0xd5, 0x55, 0x75, 0x74, 0xaf, 0xf9, 0xd0, 0x47, 0xf5, 0x3b, 0x6e, 0x84, 0x2a, 0x31,
	0xc7, 0xa8, 0x92, 0x50, 0x8c, 0xca, 0x1d, 0xab, 0x63, 0x8d, 0xda, 0x65, 0xed, 0x28, 0xe2, 0x88,
	0x61, 0xff, 0xc1, 0x13, 0x63, 0xce, 0x5c, 0xf9, 0xc0, 0x47, 0xe8, 0xee, 0xda, 0x6a, 0xbc, 0xc5,
	0x51, 0xfd, 0x41, 0x18, 0x6a, 0x33, 0x19, 0xa2, 0xd8, 0x7d, 0x52, 0xe3, 0x63, 0x82, 0xfd, 0x0d,
	0x59, 0x75, 0x26, 0x59, 0x6d, 0x06, 0xd9, 0xb0, 0xb4, 0x5f, 0xf1, 0x1f, 0x88, 0xae, 0x32, 0x93,
	0xee, 0x60, 0x36, 0x5d, 0xf5, 0xa1, 0xe8, 0xca, 0x33, 0xe9, 0x0e, 0x67, 0xd3, 0xd5, 0xfe, 0x0b,
	0xba, 0xfa, 0x2c, 0x90, 0x57, 0xb3, 0x41, 0x0e, 0x3e, 0x3e, 0x88, 0x8a, 0x3e, 0x29, 0xec, 0xa9,
	0x8d, 0x0e, 0xdb, 0xdd, 0xc1, 0x8f, 0xec, 0xbe, 0xf4, 0x29, 0x58, 0x36, 0xba, 0xc4, 0xe5, 0xdf,
	0x11, 0x8f, 0xd9, 0xd4, 0x85, 0x73, 0xe2, 0x0f, 0xec, 0x45, 0x18, 0xc8, 0xeb, 0xf1, 0xff, 0x57,
	0x56, 0x55, 0xf0, 0x2d, 0xb3, 0x54, 0x07, 0xcb, 0x67, 0x62, 0xa0, 0xc5, 0x0d, 0x3e, 0x60, 0xf0,
	0xc9, 0x74, 0x38, 0xab, 0x2a, 0x78, 0xa9, 0x39, 0xe9, 0xa4, 0xcf, 0xc1, 0x6a, 0xac, 0xbe, 0x21,
	0x86, 0xc7, 0x8f, 0x89, 0xc1, 0xe1, 0xfc, 0x76, 0x0e, 0x3d, 0x39, 0xde, 0x0a, 0x03, 0xf9, 0x79,
	0x36, 0x3e, 0x36, 0x28, 0x78, 0xa5, 0x79, 0x6b, 0x40, 0x52, 0xc0, 0xdc, 0x99, 0x63, 0x98, 0x70,
	0x41, 0x2c, 0xbc, 0x1a, 0x06, 0xf2, 0x52, 0x92, 0x74, 0x0c, 0x53, 0xc1, 0x8f, 0xbf, 0x32, 0x4c,
	0x01, 0x49, 0xdc, 0x2b, 0xdb, 0xa3, 0xae, 0x43, 0x5c, 0x0e, 0x17, 0xef, 0x40, 0x66, 0x54, 0x05,
	0x2f, 0x7d, 0x39, 0xe9, 0xa4, 0x3d, 0xb0, 0x70, 0xc6, 0xb8, 0xe1, 0xf1, 0x77, 0x30, 0x2f, 0xe0,
	0xa4, 0x30, 0x90, 0x57, 0x92, 0x58, 0x2c, 0x28, 0x78, 0xbe, 0x25, 0x0a, 0xa9, 0x06, 0xc0, 0x19,
	0x35, 0xcd, 0x41, 0xdf, 0x26, 0xd6, 0x3b, 0x08, 0x84, 0x7f, 0x73, 0xf2, 0xfc, 0x4c, 0x34, 0x05,
	0xe7, 0xbf, 0x4e, 0xeb, 0xe8, 0xa5, 0xe9, 0x7b, 0xf4, 0xca, 0xb6, 0x88, 0x07, 0x97, 0xa6, 0x5f,
	0x9a, 0x54, 0x51, 0xf0, 0xd8, 0x24, 0xbd, 0x04, 0x73, 0x0e, 0x71, 0x28, 0x5c, 0x9e, 0xde, 0x73,
	0x34, 0xaa, 0x60, 0x21, 0x4a, 0xfb, 0x60, 0xc1, 0x21, 0x4e, 0xcb, 0xbe, 0x21, 0xf0, 0xe9, 0x34,
	0x78, 0x22, 0x28, 0x38, 0xb5, 0x44, 0x0c, 0x96, 0xcd, 0xde, 0x0b, 0xfb, 0x8a, 0xb0, 0x67, 0x18,
	0x52, 0x45, 0xc1, 0x63, 0x53, 0x14, 0xa0, 0xac, 0xf5, 0x81, 0x71, 0xe2, 0xc0, 0xd5, 0x69, 0xe8,
	0x54, 0x51, 0xf0, 0xd8, 0x24, 0x35, 0xc1, 0x0a, 0x65, 0x4d, 0xcf, 0xbc, 0xb4, 0x39, 0x31, 0xf9,
	0xc0, 0x23, 0x70, 0x4d, 0xc4, 0xfe, 0x17, 0x06, 0xf2, 0x66, 0x1a, 0xcb, 0xea, 0x0a, 0x9e, 0x0a,
	0x48, 0x15, 0x90, 0xa7, 0x2c, 0x3d, 0xa6, 0xcf, 0x44, 0x7a, 0x23, 0x0c, 0xe4, 0xb5, 0x34, 0x3d,
	0x3e, 0xa3, 0x13, 0x5b, 0x9c, 0xc1, 0xa4, 0x47, 0x0c, 0x46, 0xa0, 0x74, 0x37, 0x93, 0x48, 0x22,
	0x93, 0xd4, 0xd2, 0xef, 0x39, 0xf0, 0xc4, 0xe4, 0xb6, 0x43, 0xe0, 0xba, 0x08, 0xfc, 0x96, 0x0b,
	0x03, 0x79, 0x39, 0x4e, 0x88, 0xf1, 0xe8, 0x3d, 0xff, 0x25, 0x07, 0x7e, 0xce, 0x9d, 0x22, 0xd4,
	0xa8, 0x8b, 0x3b, 0x16, 0x5d, 0x5e, 0x7d, 0x57, 0x6d, 0x88, 0xdf, 0x61, 0xcd, 0x57, 0x35, 0x54,
	0x6e, 0x97, 0xb4, 0x8a, 0x3e, 0x2a, 0x09, 0x5d, 0xd5, 0x50, 0x55, 0x3c, 0xdc, 0x49, 0x1f, 0xdd,
	0xcc, 0x4a, 0x9c, 0x52, 0xdb, 0xdf, 0x6e, 0xeb, 0x28, 0xba, 0xb8, 0x55, 0x3d, 0x7e, 0xdc, 0xe3,
	0xe1, 0x3a, 0x12, 0x57, 0x3a, 0x69, 0x26, 0x35, 0xea, 0x14, 0xc4, 0xef, 0x9e, 0xda, 0x40, 0x27,
	0xa3, 0xf6, 0x9e, 0xa6, 0xa3, 0x46, 0xfd, 0x9e, 0x78, 0x26, 0xdd, 0xd8, 0xc1, 0xf1, 0x8e, 0xa4,
	0x16, 0x58, 0x30, 0x3d, 0x12, 0x7d, 0x3b, 0xc1, 0x0d, 0xb1, 0xb9, 0xa3, 0xc9, 0xb1, 0x48, 0x84,
	0x68, 0x77, 0x2f, 0xc1, 0xff, 0x4f, 0xdb, 0x86, 0x76, 0xd3, 0xd4, 0x4e, 0xa2, 0x09, 0xda, 0x85,
	0x71, 0x7d, 0xa6, 0xe9, 0xc3, 0xca, 0x7e, 0xb5, 0xec, 0xef, 0xe0, 0x74, 0xa6, 0xe8, 0xac, 0xf5,
	0x3d, 0x6a, 0x0d, 0x4c, 0x0e, 0x37, 0xc5, 0xa4, 0x99, 0xb3, 0x96, 0x08, 0x0a, 0x4e, 0x2d, 0xd1,
	0xd1, 0x31, 0x07, 0x8c, 0x53, 0x87, 0x78, 0xf0, 0xf9, 0xf4, 0xd1, 0x49, 0x15, 0x05, 0x8f, 0x4d,
	0xd2, 0x67, 0xe0, 0xa9, 0xcd, 0x5a, 0xb6, 0xdb, 0xed, 0x91, 0x6f, 0xe8, 0x35, 0xf1, 0xe0, 0x8b,
	0xed, 0x1c, 0x5a, 0x3c, 0x86, 0x61, 0x20, 0x6f, 0x24, 0x9f, 0x58, 0x59, 0x59, 0xc1, 0xb7, 0xed,
	0xc7, 0xaf, 0x4f, 0xbe, 0xe8, 0xd2, 0x02, 0x31, 0xd8, 0x07, 0xda, 0x67, 0x85, 0x1e, 0x35, 0x8d,
	0x5e, 0xd1, 0xa4, 0x2e, 0xf7, 0x0c, 0x93, 0xb3, 0xf8, 0x93, 0xd4, 0x23, 0x7d, 0xaa, 0x39, 0xd4,
	0x22, 0x3d, 0x56, 0x4c, 0x8c, 0x45, 0xd1, 0x16, 0x27, 0x5f, 0xb3, 0xe7, 0xf3, 0xc2, 0x5a, 0xfd,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x27, 0x2d, 0xdd, 0xd8, 0xee, 0x0a, 0x00, 0x00,
}
