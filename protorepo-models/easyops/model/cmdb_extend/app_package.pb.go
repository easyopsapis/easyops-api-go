// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_package.proto

package cmdb_extend

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
//应用程序包
type AppPackage struct {
	//
	//包名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//包ID
	PackageId string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//安装路径
	InstallPath string `protobuf:"bytes,3,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	//
	//是否主程序包
	IsMaster string `protobuf:"bytes,4,opt,name=isMaster,proto3" json:"isMaster" form:"isMaster"`
	//
	//系统类型
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform" form:"platform"`
	//
	//升级前停止
	PreStop int32 `protobuf:"varint,6,opt,name=preStop,proto3" json:"preStop" form:"preStop"`
	//
	//强制升级
	ForceUpdate int32 `protobuf:"varint,7,opt,name=forceUpdate,proto3" json:"forceUpdate" form:"forceUpdate"`
	//
	//升级后重启
	PostRestart int32 `protobuf:"varint,8,opt,name=postRestart,proto3" json:"postRestart" form:"postRestart"`
	//
	//安装完成自动启动
	AutoStart int32 `protobuf:"varint,9,opt,name=autoStart,proto3" json:"autoStart" form:"autoStart"`
	//
	//目标版本
	TargetVersion string `protobuf:"bytes,10,opt,name=targetVersion,proto3" json:"targetVersion" form:"targetVersion"`
	//
	//发布后检查
	UserCheck int32 `protobuf:"varint,11,opt,name=userCheck,proto3" json:"userCheck" form:"userCheck"`
	//
	//全量升级
	FullUpdate           int32    `protobuf:"varint,12,opt,name=fullUpdate,proto3" json:"fullUpdate" form:"fullUpdate"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPackage) Reset()         { *m = AppPackage{} }
func (m *AppPackage) String() string { return proto.CompactTextString(m) }
func (*AppPackage) ProtoMessage()    {}
func (*AppPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd06b6ee71df995, []int{0}
}
func (m *AppPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPackage.Unmarshal(m, b)
}
func (m *AppPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPackage.Marshal(b, m, deterministic)
}
func (m *AppPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPackage.Merge(m, src)
}
func (m *AppPackage) XXX_Size() int {
	return xxx_messageInfo_AppPackage.Size(m)
}
func (m *AppPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPackage.DiscardUnknown(m)
}

var xxx_messageInfo_AppPackage proto.InternalMessageInfo

func (m *AppPackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppPackage) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *AppPackage) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *AppPackage) GetIsMaster() string {
	if m != nil {
		return m.IsMaster
	}
	return ""
}

func (m *AppPackage) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *AppPackage) GetPreStop() int32 {
	if m != nil {
		return m.PreStop
	}
	return 0
}

func (m *AppPackage) GetForceUpdate() int32 {
	if m != nil {
		return m.ForceUpdate
	}
	return 0
}

func (m *AppPackage) GetPostRestart() int32 {
	if m != nil {
		return m.PostRestart
	}
	return 0
}

func (m *AppPackage) GetAutoStart() int32 {
	if m != nil {
		return m.AutoStart
	}
	return 0
}

func (m *AppPackage) GetTargetVersion() string {
	if m != nil {
		return m.TargetVersion
	}
	return ""
}

func (m *AppPackage) GetUserCheck() int32 {
	if m != nil {
		return m.UserCheck
	}
	return 0
}

func (m *AppPackage) GetFullUpdate() int32 {
	if m != nil {
		return m.FullUpdate
	}
	return 0
}

func init() {
	proto.RegisterType((*AppPackage)(nil), "cmdb_extend.AppPackage")
}

func init() { proto.RegisterFile("app_package.proto", fileDescriptor_3dd06b6ee71df995) }

var fileDescriptor_3dd06b6ee71df995 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x55, 0x58, 0xb7, 0xd5, 0x05, 0xc6, 0xcc, 0x40, 0x16, 0x97, 0x4c, 0xe6, 0xb2, 0x03,
	0x6b, 0x24, 0x10, 0x12, 0xe2, 0x80, 0x44, 0x11, 0x42, 0x1c, 0x90, 0x26, 0x4f, 0x70, 0xe0, 0x32,
	0xb9, 0xc9, 0x6b, 0x5a, 0xcd, 0xc9, 0xb3, 0xec, 0x17, 0x09, 0xae, 0xfc, 0xd0, 0xfc, 0x88, 0xfc,
	0x82, 0x29, 0x4e, 0x96, 0x7a, 0xbd, 0xc5, 0xf9, 0xbe, 0x2f, 0x8e, 0x9f, 0xcc, 0x4e, 0xb5, 0xb5,
	0x37, 0x56, 0x67, 0xb7, 0xba, 0x80, 0x85, 0x75, 0x48, 0xc8, 0xe7, 0x59, 0x99, 0xaf, 0x6e, 0xe0,
	0x2f, 0x41, 0x95, 0xbf, 0xbe, 0x2c, 0xb6, 0xb4, 0xa9, 0x57, 0x8b, 0x0c, 0xcb, 0xb4, 0xc0, 0x02,
	0xd3, 0xe0, 0xac, 0xea, 0x75, 0x58, 0x85, 0x45, 0x78, 0xea, 0x5b, 0xf9, 0x7f, 0xca, 0xd8, 0x17,
	0x6b, 0xaf, 0xfa, 0x0f, 0xf2, 0x37, 0xec, 0xa0, 0xd2, 0x25, 0x88, 0xc9, 0xf9, 0xe4, 0x62, 0xb6,
	0x3c, 0x69, 0x9b, 0x64, 0xbe, 0x46, 0x57, 0x7e, 0x92, 0xdd, 0x5b, 0xa9, 0x02, 0xe4, 0xef, 0xd8,
	0x6c, 0xf8, 0x81, 0x1f, 0xb9, 0x78, 0x14, 0xcc, 0xb3, 0xb6, 0x49, 0x9e, 0xf7, 0xe6, 0x88, 0xa4,
	0xda, 0x69, 0xfc, 0x23, 0x9b, 0x6f, 0x2b, 0x4f, 0xda, 0x98, 0x2b, 0x4d, 0x1b, 0xf1, 0x38, 0x54,
	0xaf, 0xda, 0x26, 0xe1, 0x7d, 0x15, 0x41, 0xa9, 0x62, 0x95, 0xa7, 0xec, 0x78, 0xeb, 0x7f, 0x6a,
	0x4f, 0xe0, 0xc4, 0x41, 0xc8, 0x5e, 0xb4, 0x4d, 0x72, 0x32, 0x64, 0x03, 0x91, 0x6a, 0x94, 0xba,
	0xc0, 0x1a, 0x4d, 0x9d, 0x20, 0xa6, 0xfb, 0xc1, 0x3d, 0x91, 0x6a, 0x94, 0xf8, 0x5b, 0x76, 0x64,
	0x1d, 0x5c, 0x13, 0x5a, 0x71, 0x78, 0x3e, 0xb9, 0x98, 0x2e, 0x79, 0xdb, 0x24, 0xcf, 0x06, 0xbf,
	0x07, 0x52, 0xdd, 0x2b, 0xdd, 0x49, 0xd6, 0xe8, 0x32, 0xf8, 0x65, 0x73, 0x4d, 0x20, 0x8e, 0x42,
	0x11, 0x9d, 0x24, 0x82, 0x52, 0xc5, 0x6a, 0x57, 0x5a, 0xf4, 0xa4, 0xc0, 0x93, 0x76, 0x24, 0x8e,
	0xf7, 0xcb, 0x08, 0x4a, 0x15, 0xab, 0xdd, 0xc4, 0x75, 0x4d, 0x78, 0x1d, 0xba, 0x59, 0xe8, 0xa2,
	0x89, 0x8f, 0x48, 0xaa, 0x9d, 0xc6, 0x3f, 0xb3, 0xa7, 0xa4, 0x5d, 0x01, 0xf4, 0x1b, 0x9c, 0xdf,
	0x62, 0x25, 0x58, 0x98, 0x85, 0x68, 0x9b, 0xe4, 0xac, 0xef, 0x1e, 0x60, 0xa9, 0x1e, 0xea, 0xdd,
	0x9e, 0xb5, 0x07, 0xf7, 0x75, 0x03, 0xd9, 0xad, 0x98, 0xef, 0xef, 0x39, 0x22, 0xa9, 0x76, 0x1a,
	0xff, 0xc0, 0xd8, 0xba, 0x36, 0x66, 0x18, 0xcd, 0x93, 0x10, 0xbd, 0x6c, 0x9b, 0xe4, 0x74, 0x18,
	0xcd, 0xc8, 0xa4, 0x8a, 0xc4, 0xe5, 0xf7, 0x3f, 0xdf, 0x0a, 0x5c, 0x80, 0xf6, 0xff, 0xd0, 0xfa,
	0x85, 0xc1, 0x4c, 0x9b, 0x34, 0xc3, 0x8a, 0x9c, 0xce, 0xc8, 0xf7, 0x17, 0xd8, 0x81, 0xc5, 0xcb,
	0x12, 0x73, 0x30, 0x3e, 0x1d, 0xc4, 0x34, 0x2c, 0xd3, 0xe8, 0xf2, 0xaf, 0x0e, 0x83, 0xfb, 0xfe,
	0x2e, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xb2, 0x27, 0x2d, 0x25, 0x03, 0x00, 0x00,
}
