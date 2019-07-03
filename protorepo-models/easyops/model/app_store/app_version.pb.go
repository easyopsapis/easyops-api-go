// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_version.proto

package app_store

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
//小产品版本
type AppVersion struct {
	//
	//版本实例id
	VersionId string `protobuf:"bytes,1,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//小产品名称+版本组合的唯一id
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//版本名称
	VersionName string `protobuf:"bytes,3,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//版本变更记录
	ChangeLog string `protobuf:"bytes,4,opt,name=changeLog,proto3" json:"changeLog" form:"changeLog"`
	//
	//版本发布时间
	ReleaseTime          string   `protobuf:"bytes,5,opt,name=releaseTime,proto3" json:"releaseTime" form:"releaseTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppVersion) Reset()         { *m = AppVersion{} }
func (m *AppVersion) String() string { return proto.CompactTextString(m) }
func (*AppVersion) ProtoMessage()    {}
func (*AppVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f2457a79bb9d2bd, []int{0}
}
func (m *AppVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppVersion.Unmarshal(m, b)
}
func (m *AppVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppVersion.Marshal(b, m, deterministic)
}
func (m *AppVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppVersion.Merge(m, src)
}
func (m *AppVersion) XXX_Size() int {
	return xxx_messageInfo_AppVersion.Size(m)
}
func (m *AppVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_AppVersion.DiscardUnknown(m)
}

var xxx_messageInfo_AppVersion proto.InternalMessageInfo

func (m *AppVersion) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *AppVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppVersion) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *AppVersion) GetChangeLog() string {
	if m != nil {
		return m.ChangeLog
	}
	return ""
}

func (m *AppVersion) GetReleaseTime() string {
	if m != nil {
		return m.ReleaseTime
	}
	return ""
}

func init() {
	proto.RegisterType((*AppVersion)(nil), "app_store.AppVersion")
}

func init() { proto.RegisterFile("app_version.proto", fileDescriptor_5f2457a79bb9d2bd) }

var fileDescriptor_5f2457a79bb9d2bd = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0xd1, 0xda, 0x42, 0x46, 0x68, 0x35, 0x2d, 0x25, 0x78, 0x93, 0x12, 0xa4, 0x08, 0x76,
	0x32, 0x55, 0xa1, 0xb4, 0xde, 0x69, 0x69, 0xa1, 0xd0, 0x7a, 0x11, 0xca, 0x5e, 0xf8, 0x97, 0x31,
	0x19, 0x63, 0xd8, 0x24, 0x27, 0xcc, 0x44, 0x65, 0x15, 0x1f, 0x71, 0x5f, 0xc1, 0x85, 0x7d, 0x04,
	0x9f, 0x60, 0xc9, 0x24, 0x9b, 0x64, 0x6f, 0xc2, 0x9c, 0xf3, 0x7d, 0xdf, 0x8f, 0x2f, 0x1c, 0xd4,
	0xa4, 0x51, 0xb4, 0xda, 0x33, 0x2e, 0x3c, 0x08, 0xcd, 0x88, 0x43, 0x0c, 0xaa, 0x92, 0xac, 0x44,
	0x0c, 0x9c, 0xb5, 0xb0, 0xeb, 0xc5, 0xdb, 0xdd, 0xda, 0xb4, 0x21, 0x20, 0x2e, 0xb8, 0x40, 0xa4,
	0x63, 0xbd, 0xdb, 0xc8, 0x49, 0x0e, 0xf2, 0x95, 0x26, 0x5b, 0xdf, 0x4a, 0xf6, 0xe0, 0xe0, 0xc5,
	0xb7, 0x70, 0x20, 0x2e, 0x60, 0x29, 0xe2, 0x3d, 0xf5, 0x3d, 0x87, 0xc6, 0xc0, 0x05, 0xc9, 0x9f,
	0x69, 0xce, 0xb8, 0xaf, 0x22, 0x34, 0x8a, 0xa2, 0x9b, 0xb4, 0x86, 0xfa, 0x1b, 0x29, 0x59, 0xa3,
	0x3f, 0x8e, 0x56, 0xf9, 0x54, 0xe9, 0x28, 0xe3, 0xce, 0xf5, 0xa2, 0x37, 0x36, 0xc0, 0x83, 0xa1,
	0x91, 0x4b, 0xc6, 0xe3, 0x83, 0xde, 0x40, 0x6f, 0x97, 0xb3, 0xaf, 0xf8, 0x07, 0xc5, 0xc7, 0xc5,
	0xa9, 0x37, 0x38, 0xb7, 0xad, 0x22, 0xaa, 0x0e, 0x51, 0x2d, 0xa4, 0x01, 0xd3, 0xaa, 0x12, 0xf1,
	0xf9, 0x7a, 0xd1, 0xeb, 0x29, 0x22, 0xd9, 0x26, 0xe9, 0xf7, 0xa8, 0xb9, 0x9c, 0x51, 0x7c, 0x1c,
	0xe1, 0xe9, 0x6a, 0x71, 0xea, 0x7d, 0x19, 0xf4, 0xcf, 0x6d, 0x4b, 0x66, 0xd4, 0x7f, 0xa8, 0x9e,
	0x81, 0x26, 0x09, 0xe2, 0x95, 0x44, 0x74, 0xaf, 0x17, 0x5d, 0x7d, 0xd1, 0x62, 0x92, 0x91, 0x9a,
	0xe8, 0xdd, 0x72, 0xee, 0x74, 0xe7, 0xe6, 0xf3, 0xa7, 0x6d, 0x95, 0xf3, 0x6a, 0x1f, 0x29, 0xf6,
	0x96, 0x86, 0x2e, 0xfb, 0x0b, 0xae, 0x56, 0x93, 0xb0, 0x0f, 0xc5, 0x2f, 0xe5, 0x92, 0x61, 0x15,
	0x36, 0xf5, 0x3b, 0xaa, 0x73, 0xe6, 0x33, 0x2a, 0xd8, 0x7f, 0x2f, 0x60, 0xda, 0x6b, 0x99, 0xfa,
	0x58, 0x54, 0x28, 0x89, 0x86, 0x55, 0xb6, 0x8e, 0x7f, 0x4d, 0x7f, 0xba, 0x60, 0x32, 0x2a, 0xee,
	0x20, 0x12, 0xa6, 0x0f, 0x36, 0xf5, 0x89, 0x0d, 0x61, 0xcc, 0xa9, 0x1d, 0x8b, 0xf4, 0x86, 0x9c,
	0x45, 0x80, 0x03, 0x70, 0x98, 0x2f, 0x48, 0x66, 0x24, 0x72, 0x24, 0xf9, 0xf5, 0xd7, 0x6f, 0xa4,
	0x73, 0xf0, 0x14, 0x00, 0x00, 0xff, 0xff, 0xba, 0xa9, 0xb3, 0x02, 0x24, 0x02, 0x00, 0x00,
}
