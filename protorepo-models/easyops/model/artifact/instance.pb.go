// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance.proto

package artifact

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//程序包实例
type Instance struct {
	//
	//所属应用Id
	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//所属集群Id
	ClusterId string `protobuf:"bytes,2,opt,name=clusterId,proto3" json:"clusterId" form:"clusterId"`
	//
	//创建者
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime string `protobuf:"bytes,4,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//设备Id
	DeviceId string `protobuf:"bytes,5,opt,name=deviceId,proto3" json:"deviceId" form:"deviceId"`
	//
	//设备IP
	DeviceIp string `protobuf:"bytes,6,opt,name=deviceIp,proto3" json:"deviceIp" form:"deviceIp"`
	//
	//安装路径
	InstallPath string `protobuf:"bytes,7,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	//
	//安装时间
	InstallTime string `protobuf:"bytes,8,opt,name=installTime,proto3" json:"installTime" form:"installTime"`
	//
	//实例Id
	InstanceId string `protobuf:"bytes,9,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//最近一次操作时间
	LastOperation string `protobuf:"bytes,10,opt,name=lastOperation,proto3" json:"lastOperation" form:"lastOperation"`
	//
	//修改时间
	Mtime string `protobuf:"bytes,11,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//组织Id
	Org int32 `protobuf:"varint,12,opt,name=org,proto3" json:"org" form:"org"`
	//
	//所属程序包Id
	PackageId string `protobuf:"bytes,13,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//所属包名称
	PackageName string `protobuf:"bytes,14,opt,name=packageName,proto3" json:"packageName" form:"packageName"`
	//
	//包类型
	PackageType string `protobuf:"bytes,15,opt,name=packageType,proto3" json:"packageType" form:"packageType"`
	//
	//更新时间
	UpdateTime string `protobuf:"bytes,16,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//所属版本Id
	VersionId string `protobuf:"bytes,17,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//所属版本名
	VersionName          string   `protobuf:"bytes,18,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22322185b2070b, []int{0}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Instance) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Instance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Instance) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Instance) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Instance) GetDeviceIp() string {
	if m != nil {
		return m.DeviceIp
	}
	return ""
}

func (m *Instance) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *Instance) GetInstallTime() string {
	if m != nil {
		return m.InstallTime
	}
	return ""
}

func (m *Instance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Instance) GetLastOperation() string {
	if m != nil {
		return m.LastOperation
	}
	return ""
}

func (m *Instance) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *Instance) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *Instance) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *Instance) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Instance) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *Instance) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Instance) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *Instance) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func init() {
	proto.RegisterType((*Instance)(nil), "artifact.Instance")
}

func init() { proto.RegisterFile("instance.proto", fileDescriptor_fd22322185b2070b) }

var fileDescriptor_fd22322185b2070b = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x97, 0xdd, 0x6a, 0xe3, 0x46,
	0x14, 0xc7, 0xab, 0xdd, 0x66, 0x37, 0x9e, 0xec, 0x66, 0x13, 0x11, 0x16, 0x91, 0x1b, 0x07, 0xe3,
	0x96, 0x51, 0x9c, 0xb1, 0x2c, 0xd9, 0xce, 0x6e, 0x44, 0x59, 0x91, 0x50, 0xb6, 0xf8, 0xa2, 0x1f,
	0x08, 0x5f, 0x45, 0x23, 0x97, 0x89, 0xa4, 0x78, 0x4d, 0x6c, 0x4b, 0x48, 0x4a, 0xc2, 0xc6, 0xd2,
	0x65, 0x29, 0x14, 0x4a, 0xa1, 0x8f, 0xb1, 0xd0, 0x97, 0xe8, 0xcd, 0xf6, 0x15, 0x7a, 0x51, 0x17,
	0x7a, 0x55, 0xe8, 0x9d, 0x9f, 0xa0, 0x68, 0xc6, 0xb6, 0x94, 0x75, 0xd4, 0x0f, 0x58, 0x07, 0x72,
	0x63, 0xeb, 0xcc, 0x39, 0x67, 0xe6, 0xf7, 0xd7, 0x91, 0x66, 0x8e, 0xc0, 0x7a, 0x6f, 0x18, 0x84,
	0x64, 0x68, 0x39, 0x55, 0xcf, 0x77, 0x43, 0x97, 0x5f, 0x25, 0x7e, 0xd8, 0x3b, 0x25, 0x56, 0xb8,
	0x8d, 0xba, 0xbd, 0xf0, 0xd5, 0xf9, 0x49, 0xd5, 0x72, 0x07, 0x52, 0xd7, 0xed, 0xba, 0x12, 0x0d,
	0x38, 0x39, 0x3f, 0xa5, 0x16, 0x35, 0xe8, 0x15, 0x4b, 0xdc, 0xde, 0xcf, 0x84, 0x0f, 0x2e, 0x7b,
	0xe1, 0x99, 0x7b, 0x29, 0x75, 0x5d, 0x44, 0x9d, 0xe8, 0x82, 0xf4, 0x7b, 0x36, 0x09, 0x5d, 0x3f,
	0x90, 0xe6, 0x97, 0x2c, 0xaf, 0xf4, 0xdb, 0x36, 0x58, 0x6d, 0x4d, 0x19, 0x78, 0x15, 0xac, 0x10,
	0xcf, 0x6b, 0xd9, 0x02, 0xb7, 0xc3, 0xc1, 0xc2, 0x51, 0x79, 0x32, 0x2e, 0x3e, 0x3a, 0x75, 0xfd,
	0x81, 0x5a, 0xa2, 0xc3, 0xa5, 0x3f, 0x7e, 0x2f, 0x6e, 0x80, 0xf5, 0x8e, 0x51, 0x43, 0x07, 0x04,
	0x5d, 0x99, 0x23, 0xb9, 0x1e, 0x97, 0x75, 0x96, 0xc2, 0x7f, 0xc7, 0x81, 0x82, 0xd5, 0x3f, 0x0f,
	0x42, 0xc7, 0x6f, 0xd9, 0xc2, 0x3d, 0x3a, 0x41, 0x7f, 0x32, 0x2e, 0x6e, 0xb0, 0x09, 0xe6, 0xae,
	0x64, 0x92, 0x36, 0xd0, 0x3b, 0x86, 0x46, 0xd0, 0xd5, 0x21, 0x3a, 0xae, 0xa1, 0x03, 0x73, 0xf4,
	0x3c, 0x46, 0xd7, 0xec, 0xc6, 0xff, 0xb4, 0x65, 0x25, 0x2e, 0xeb, 0xe9, 0xf2, 0xfc, 0x4b, 0xf0,
	0xd0, 0xf2, 0x9d, 0x44, 0xa6, 0x70, 0x9f, 0x92, 0xec, 0x4d, 0xc6, 0xc5, 0xf5, 0x29, 0x09, 0x73,
	0x24, 0x1c, 0x4f, 0xc1, 0x56, 0xc7, 0x38, 0x44, 0xc7, 0x04, 0x5d, 0x7d, 0x8d, 0x4c, 0x7c, 0x39,
	0xaa, 0xef, 0xed, 0x37, 0xe2, 0xb2, 0x3e, 0x4b, 0xe6, 0x7f, 0xe5, 0xc0, 0x8a, 0x15, 0xf6, 0x06,
	0x8e, 0xf0, 0x21, 0x9d, 0xe6, 0x67, 0x2e, 0xbd, 0x25, 0x74, 0x3c, 0x99, 0xe5, 0x27, 0x0e, 0xbc,
	0xe1, 0x3a, 0x10, 0x6a, 0xaa, 0x21, 0xa3, 0x03, 0x33, 0xb9, 0x3b, 0xe6, 0xae, 0xa8, 0x19, 0x53,
	0x52, 0x11, 0x41, 0xd9, 0xa8, 0x21, 0xc5, 0x8c, 0x6a, 0xd4, 0x2f, 0x22, 0x58, 0x37, 0x6a, 0x48,
	0x9e, 0xd9, 0x91, 0x21, 0x23, 0x85, 0x65, 0x89, 0x46, 0x7b, 0xc7, 0x84, 0x8a, 0x51, 0x43, 0x75,
	0x33, 0xa2, 0x31, 0x6c, 0x58, 0x85, 0x46, 0x0d, 0x35, 0x67, 0x46, 0x7a, 0x0d, 0x71, 0x95, 0xfe,
	0x57, 0x44, 0x0d, 0x1e, 0x47, 0x46, 0x05, 0x99, 0x50, 0x53, 0x6f, 0x48, 0xcf, 0x64, 0x6b, 0x65,
	0x9d, 0x29, 0xe2, 0xbf, 0xe5, 0xc0, 0xaa, 0xed, 0x5c, 0xf4, 0x2c, 0xa7, 0x65, 0x0b, 0x2b, 0x54,
	0xde, 0xd9, 0x64, 0x5c, 0x7c, 0xc2, 0xd4, 0xcd, 0x3c, 0xcb, 0x2b, 0xd7, 0x7c, 0x71, 0xfe, 0x6d,
	0x61, 0x4e, 0xe2, 0x09, 0x0f, 0x28, 0xc9, 0x9b, 0xc2, 0x02, 0x8a, 0x97, 0xa0, 0xfc, 0x58, 0x00,
	0x3f, 0x14, 0x3a, 0x10, 0x2a, 0xb0, 0x49, 0x45, 0x8d, 0xe4, 0x38, 0x51, 0xda, 0x30, 0xb1, 0x3d,
	0x92, 0x63, 0x91, 0xa9, 0xd6, 0x12, 0x63, 0x4f, 0x89, 0x45, 0x88, 0xab, 0xff, 0x31, 0x52, 0x1c,
	0xd5, 0x63, 0x31, 0xc2, 0xc1, 0x2e, 0x84, 0xc9, 0xbd, 0x3e, 0x38, 0x44, 0x2f, 0x09, 0x3a, 0x35,
	0x47, 0xf2, 0x5e, 0x23, 0x56, 0xc5, 0xd1, 0xb3, 0x78, 0x61, 0x34, 0x52, 0x45, 0x31, 0xba, 0x31,
	0x78, 0x3f, 0x86, 0xea, 0x42, 0x34, 0x84, 0x0a, 0xe3, 0x88, 0x94, 0x29, 0x45, 0x24, 0x63, 0x1b,
	0xdb, 0x11, 0x7d, 0x06, 0x34, 0x6c, 0x33, 0xd8, 0x7f, 0x89, 0x61, 0x98, 0xb9, 0x2b, 0x37, 0x63,
	0x08, 0x17, 0xd7, 0x16, 0x99, 0xc4, 0x48, 0xbd, 0x15, 0x86, 0x46, 0x2e, 0x43, 0x92, 0x76, 0x93,
	0x4b, 0x7b, 0x9f, 0x60, 0xff, 0x40, 0x56, 0xcf, 0x25, 0x6b, 0xe4, 0x90, 0x8d, 0x6a, 0x7b, 0x4a,
	0x7c, 0x4b, 0x74, 0x4a, 0x2e, 0x5d, 0x33, 0x9f, 0xae, 0x7e, 0x5b, 0x74, 0x72, 0x2e, 0xdd, 0x7e,
	0x3e, 0x5d, 0x63, 0x19, 0x74, 0x6a, 0x1e, 0xc8, 0xb3, 0x7c, 0x90, 0xe6, 0xfb, 0x07, 0x11, 0xe1,
	0x47, 0xd5, 0x8a, 0xa8, 0xe1, 0x60, 0x37, 0xdd, 0xc9, 0x3c, 0xfe, 0x1b, 0x0e, 0xac, 0xd1, 0x13,
	0xbd, 0xdf, 0xff, 0x8a, 0x84, 0xaf, 0x84, 0x87, 0x74, 0x33, 0xb3, 0x26, 0xe3, 0x22, 0xcf, 0xf6,
	0xb2, 0x8c, 0x33, 0xd9, 0xce, 0x5a, 0xe0, 0xb3, 0x0e, 0x84, 0x58, 0x32, 0x3a, 0x58, 0x52, 0xf1,
	0x2e, 0xd6, 0x4a, 0xa5, 0x4f, 0x5e, 0xe0, 0x08, 0xfb, 0x78, 0x68, 0x56, 0xc4, 0x8a, 0x18, 0x41,
	0x2c, 0x89, 0x91, 0xc1, 0xb6, 0x4e, 0x53, 0x85, 0x18, 0x1b, 0x1d, 0x8c, 0x17, 0x23, 0x77, 0xcb,
	0x7a, 0x76, 0x5d, 0xfe, 0xaf, 0x94, 0xa3, 0x9d, 0x9c, 0x5e, 0xab, 0x94, 0xe3, 0x17, 0x6e, 0x01,
	0xa4, 0x7d, 0x47, 0xcf, 0xb0, 0xac, 0x3a, 0xfe, 0x7b, 0x0e, 0x80, 0x59, 0x1f, 0xd5, 0xb2, 0x85,
	0x02, 0x15, 0x3b, 0x98, 0x8c, 0x8b, 0x9b, 0x19, 0xad, 0xc3, 0xe5, 0x9e, 0x66, 0x19, 0x00, 0xfe,
	0x05, 0x78, 0xdc, 0x27, 0x41, 0xf8, 0xa5, 0xe7, 0xf8, 0x24, 0xec, 0xb9, 0x43, 0x01, 0x50, 0x22,
	0x61, 0x32, 0x2e, 0x6e, 0x31, 0xa2, 0x6b, 0xee, 0x92, 0x7e, 0x3d, 0x9c, 0x76, 0x1d, 0x03, 0xda,
	0x75, 0xac, 0x2d, 0x74, 0x1d, 0x83, 0x3b, 0xdb, 0x75, 0x50, 0x72, 0xfe, 0x63, 0x70, 0xdf, 0xf5,
	0xbb, 0xc2, 0xa3, 0x1d, 0x0e, 0xae, 0x1c, 0x6d, 0x4d, 0xc6, 0x45, 0xc0, 0x74, 0xb9, 0x7e, 0x37,
	0x51, 0x75, 0x6f, 0xe3, 0x03, 0x3d, 0x09, 0xa0, 0xed, 0xa4, 0x47, 0xac, 0x33, 0xd2, 0x4d, 0x4a,
	0xfa, 0xf8, 0xdd, 0x76, 0x72, 0xee, 0x5a, 0x62, 0x3b, 0x39, 0x5f, 0x83, 0x7f, 0x0e, 0xd6, 0xa6,
	0xc6, 0x17, 0x64, 0xe0, 0x08, 0xeb, 0x94, 0xe6, 0x69, 0xfa, 0x32, 0x65, 0x9c, 0x25, 0x3d, 0x1b,
	0x9a, 0xc9, 0x6c, 0xbf, 0xf6, 0x1c, 0xe1, 0x49, 0x4e, 0x66, 0xe2, 0x4c, 0x33, 0x13, 0x8b, 0xff,
	0x93, 0x03, 0xe0, 0xdc, 0xb3, 0x49, 0xe8, 0xd0, 0x37, 0x78, 0x83, 0x66, 0xbe, 0xe5, 0xd2, 0xa7,
	0x3a, 0x75, 0xde, 0xc5, 0xc7, 0x21, 0xa3, 0x8d, 0xd6, 0xfa, 0xc2, 0xf1, 0x83, 0x9e, 0x3b, 0x6c,
	0xd9, 0xc2, 0xe6, 0xbb, 0xb5, 0x9e, 0xbb, 0x96, 0x58, 0xeb, 0xf9, 0x1a, 0xfc, 0xe7, 0x60, 0x6d,
	0x6a, 0xd0, 0x5a, 0xf3, 0x94, 0xa6, 0x92, 0x56, 0x2c, 0xe3, 0x4c, 0x78, 0x36, 0xc1, 0x93, 0x0e,
	0xb6, 0x2b, 0xb8, 0x3a, 0xfb, 0x29, 0xeb, 0xd9, 0xfc, 0xa3, 0x4f, 0x8f, 0x8f, 0xba, 0x6e, 0xd5,
	0x21, 0xc1, 0x6b, 0xd7, 0x0b, 0xaa, 0x7d, 0xd7, 0x22, 0x7d, 0xc9, 0x72, 0x87, 0xa1, 0x4f, 0xac,
	0x30, 0x60, 0xdf, 0x74, 0xbe, 0xe3, 0xb9, 0x68, 0xe0, 0xda, 0x4e, 0x3f, 0x90, 0xa6, 0x81, 0x12,
	0x35, 0xa5, 0xd9, 0xc7, 0xe0, 0xc9, 0x03, 0x1a, 0x58, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0x74, 0x8d, 0x91, 0x2f, 0x0e, 0x00, 0x00,
}
