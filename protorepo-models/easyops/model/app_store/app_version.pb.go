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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

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
	ReleaseTime string `protobuf:"bytes,5,opt,name=releaseTime,proto3" json:"releaseTime" form:"releaseTime"`
	//
	//依赖包列表
	Dependencies         []*AppVersion_Dependencies `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies" form:"dependencies"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
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

func (m *AppVersion) GetDependencies() []*AppVersion_Dependencies {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type AppVersion_Dependencies struct {
	//
	//依赖组件
	PackageName string `protobuf:"bytes,1,opt,name=packageName,proto3" json:"packageName" form:"packageName"`
	//
	//依赖版本semver约束
	Constraint           string   `protobuf:"bytes,2,opt,name=constraint,proto3" json:"constraint" form:"constraint"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppVersion_Dependencies) Reset()         { *m = AppVersion_Dependencies{} }
func (m *AppVersion_Dependencies) String() string { return proto.CompactTextString(m) }
func (*AppVersion_Dependencies) ProtoMessage()    {}
func (*AppVersion_Dependencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f2457a79bb9d2bd, []int{0, 0}
}
func (m *AppVersion_Dependencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppVersion_Dependencies.Unmarshal(m, b)
}
func (m *AppVersion_Dependencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppVersion_Dependencies.Marshal(b, m, deterministic)
}
func (m *AppVersion_Dependencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppVersion_Dependencies.Merge(m, src)
}
func (m *AppVersion_Dependencies) XXX_Size() int {
	return xxx_messageInfo_AppVersion_Dependencies.Size(m)
}
func (m *AppVersion_Dependencies) XXX_DiscardUnknown() {
	xxx_messageInfo_AppVersion_Dependencies.DiscardUnknown(m)
}

var xxx_messageInfo_AppVersion_Dependencies proto.InternalMessageInfo

func (m *AppVersion_Dependencies) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *AppVersion_Dependencies) GetConstraint() string {
	if m != nil {
		return m.Constraint
	}
	return ""
}

func init() {
	proto.RegisterType((*AppVersion)(nil), "app_store.AppVersion")
	proto.RegisterType((*AppVersion_Dependencies)(nil), "app_store.AppVersion.Dependencies")
}

func init() { proto.RegisterFile("app_version.proto", fileDescriptor_5f2457a79bb9d2bd) }

var fileDescriptor_5f2457a79bb9d2bd = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0xc9, 0x9a, 0x15, 0x22, 0x97, 0xad, 0xd1, 0xfe, 0x99, 0xc0, 0x70, 0x11, 0xb9, 0x08,
	0x74, 0xb6, 0xdb, 0x84, 0x8d, 0x6d, 0x77, 0xcd, 0xb6, 0xc2, 0x60, 0xeb, 0x85, 0x19, 0xbb, 0xe8,
	0x9f, 0x04, 0xc5, 0x56, 0x5d, 0x53, 0xdb, 0x9f, 0x90, 0xd4, 0x96, 0xb5, 0x94, 0x3d, 0xd8, 0xde,
	0xc5, 0x83, 0x3d, 0x82, 0x9f, 0x60, 0x58, 0xf6, 0x6c, 0x65, 0x37, 0x41, 0x9f, 0xce, 0x39, 0x3f,
	0x4e, 0x3e, 0x19, 0x0d, 0x29, 0xe7, 0xcb, 0x6b, 0x26, 0x64, 0x02, 0xb9, 0xc7, 0x05, 0x28, 0xc0,
	0x83, 0xea, 0x4a, 0x2a, 0x10, 0x6c, 0xe4, 0xc6, 0x89, 0xba, 0xb8, 0x5a, 0x79, 0x21, 0x64, 0x7e,
	0x0c, 0x31, 0xf8, 0xda, 0xb1, 0xba, 0x3a, 0xd7, 0x93, 0x1e, 0xf4, 0xa9, 0x4e, 0x8e, 0xde, 0x18,
	0xf6, 0xec, 0x26, 0x51, 0x97, 0x70, 0xe3, 0xc7, 0xe0, 0x6a, 0xd1, 0xbd, 0xa6, 0x69, 0x12, 0x51,
	0x05, 0x42, 0xfa, 0xed, 0xb1, 0xce, 0x91, 0x5f, 0x7d, 0x84, 0x0e, 0x38, 0xff, 0x5e, 0xd7, 0xc0,
	0x87, 0x68, 0xd0, 0x34, 0xfa, 0x1c, 0xd9, 0xbd, 0x9d, 0xde, 0x64, 0x30, 0x9f, 0x94, 0x85, 0xb3,
	0x7d, 0x0e, 0x22, 0x7b, 0x4f, 0x5a, 0x89, 0xfc, 0xf9, 0xed, 0x6c, 0xa3, 0x47, 0x8b, 0x93, 0x3d,
	0xf7, 0x1d, 0x75, 0x6f, 0xcf, 0xee, 0xf6, 0x67, 0xf7, 0xe3, 0xa0, 0x8b, 0xe2, 0x43, 0xd4, 0xcf,
	0x69, 0xc6, 0xec, 0x07, 0x1a, 0x31, 0x2d, 0x0b, 0xc7, 0xaa, 0x11, 0xd5, 0x6d, 0x95, 0x76, 0xd0,
	0xcb, 0xc5, 0x09, 0x75, 0x6f, 0x0f, 0xdc, 0xe3, 0xe5, 0x59, 0x83, 0xd1, 0xe7, 0xbb, 0xbd, 0x57,
	0xb3, 0xfd, 0xfb, 0x71, 0xa0, 0xf3, 0xf8, 0x2b, 0xb2, 0x1a, 0xe8, 0x51, 0x85, 0xdb, 0xd0, 0xb8,
	0xdd, 0xb2, 0x70, 0xf0, 0x5a, 0xa3, 0xa3, 0x86, 0x3a, 0x44, 0x8f, 0x17, 0xa7, 0xd1, 0xee, 0xa9,
	0xf7, 0xef, 0x67, 0x1c, 0x98, 0x79, 0x3c, 0x45, 0x83, 0xf0, 0x82, 0xe6, 0x31, 0xfb, 0x02, 0xb1,
	0xdd, 0xd7, 0xb0, 0xa7, 0xdd, 0xdf, 0x6b, 0x25, 0x12, 0x74, 0x36, 0xfc, 0x16, 0x59, 0x82, 0xa5,
	0x8c, 0x4a, 0xf6, 0x2d, 0xc9, 0x98, 0xfd, 0x50, 0xa7, 0x9e, 0x77, 0x15, 0x0c, 0x91, 0x04, 0xa6,
	0x15, 0x2f, 0xd1, 0x56, 0xc4, 0x38, 0xcb, 0x23, 0x96, 0x87, 0x09, 0x93, 0xf6, 0xe6, 0xce, 0xc6,
	0xc4, 0x9a, 0x12, 0xaf, 0x7d, 0x64, 0xaf, 0xdb, 0xbc, 0xf7, 0xd1, 0x70, 0xce, 0x5f, 0x94, 0x85,
	0xf3, 0xa4, 0xc6, 0x9b, 0x04, 0x12, 0xac, 0x01, 0x47, 0x3f, 0xd1, 0x96, 0x19, 0xab, 0xaa, 0x72,
	0x1a, 0x5e, 0xd2, 0x98, 0xe9, 0x6d, 0xf5, 0xfe, 0xaf, 0x6a, 0x88, 0x24, 0x30, 0xad, 0xf8, 0x35,
	0x42, 0x21, 0xe4, 0x52, 0x09, 0x9a, 0xe4, 0xaa, 0x79, 0xb5, 0x67, 0x65, 0xe1, 0x0c, 0x9b, 0xcd,
	0xb4, 0x1a, 0x09, 0x0c, 0xe3, 0xfc, 0xd3, 0xf1, 0x87, 0x18, 0x3c, 0x46, 0xe5, 0x0f, 0xe0, 0xd2,
	0x4b, 0x21, 0xa4, 0xa9, 0x1f, 0x42, 0xae, 0x04, 0x0d, 0x95, 0xac, 0xbf, 0x58, 0xc1, 0x38, 0xb8,
	0x19, 0x44, 0x2c, 0x95, 0x7e, 0x63, 0xf4, 0xf5, 0xe8, 0xb7, 0x6b, 0x58, 0x6d, 0x6a, 0xe7, 0xec,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x82, 0x93, 0x8a, 0x12, 0x03, 0x00, 0x00,
}
