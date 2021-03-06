// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql_package_version.proto

package database_delivery

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
//SQL包版本
type SQLPackageVersion struct {
	//
	//版本ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//仓库版本ID
	RepoVersionId string `protobuf:"bytes,3,opt,name=repoVersionId,proto3" json:"repoVersionId" form:"repoVersionId"`
	//
	//程序包版本ID
	AppVersionId string `protobuf:"bytes,4,opt,name=appVersionId,proto3" json:"appVersionId" form:"appVersionId"`
	//
	//版本名称
	VersionName string `protobuf:"bytes,5,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//备注
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//创建者
	Creator string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime int64 `protobuf:"varint,8,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime                int64    `protobuf:"varint,9,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SQLPackageVersion) Reset()         { *m = SQLPackageVersion{} }
func (m *SQLPackageVersion) String() string { return proto.CompactTextString(m) }
func (*SQLPackageVersion) ProtoMessage()    {}
func (*SQLPackageVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa1764aae9c11, []int{0}
}
func (m *SQLPackageVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQLPackageVersion.Unmarshal(m, b)
}
func (m *SQLPackageVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQLPackageVersion.Marshal(b, m, deterministic)
}
func (m *SQLPackageVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQLPackageVersion.Merge(m, src)
}
func (m *SQLPackageVersion) XXX_Size() int {
	return xxx_messageInfo_SQLPackageVersion.Size(m)
}
func (m *SQLPackageVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SQLPackageVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SQLPackageVersion proto.InternalMessageInfo

func (m *SQLPackageVersion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SQLPackageVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SQLPackageVersion) GetRepoVersionId() string {
	if m != nil {
		return m.RepoVersionId
	}
	return ""
}

func (m *SQLPackageVersion) GetAppVersionId() string {
	if m != nil {
		return m.AppVersionId
	}
	return ""
}

func (m *SQLPackageVersion) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *SQLPackageVersion) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *SQLPackageVersion) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SQLPackageVersion) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *SQLPackageVersion) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func init() {
	proto.RegisterType((*SQLPackageVersion)(nil), "database_delivery.SQLPackageVersion")
}

func init() { proto.RegisterFile("sql_package_version.proto", fileDescriptor_955aa1764aae9c11) }

var fileDescriptor_955aa1764aae9c11 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x95, 0xa4, 0x1f, 0x64, 0x5b, 0x4a, 0xbb, 0x20, 0x30, 0xbd, 0xb8, 0xdd, 0x56, 0xa8,
	0x07, 0xd6, 0x6e, 0x88, 0x84, 0x08, 0x48, 0x48, 0xe4, 0x86, 0xf8, 0x10, 0xb8, 0x12, 0x87, 0x46,
	0x69, 0xb4, 0xb1, 0xb7, 0x66, 0x55, 0x6f, 0xc6, 0xd8, 0x6e, 0xaa, 0xb6, 0xea, 0x5f, 0x35, 0x12,
	0x67, 0x4e, 0xfe, 0x05, 0x68, 0x67, 0x4d, 0x71, 0xca, 0xc9, 0x33, 0xf3, 0x3e, 0xef, 0x78, 0x76,
	0x34, 0xe4, 0x69, 0xfe, 0x23, 0x99, 0xa4, 0x22, 0x3c, 0x13, 0xb1, 0x9c, 0xcc, 0x65, 0x96, 0x2b,
	0x98, 0x79, 0x69, 0x06, 0x05, 0xd0, 0xad, 0x48, 0x14, 0x62, 0x2a, 0x72, 0x39, 0x89, 0x64, 0xa2,
	0xe6, 0x32, 0xbb, 0xdc, 0xe6, 0xb1, 0x2a, 0xbe, 0x9f, 0x4f, 0xbd, 0x10, 0xb4, 0x1f, 0x43, 0x0c,
	0x3e, 0x92, 0xd3, 0xf3, 0x53, 0xcc, 0x30, 0xc1, 0xc8, 0x76, 0xd8, 0x7e, 0xd9, 0xc0, 0xf5, 0x85,
	0x2a, 0xce, 0xe0, 0xc2, 0x8f, 0x81, 0xa3, 0xc8, 0xe7, 0x22, 0x51, 0x91, 0x28, 0x20, 0xcb, 0xfd,
	0xdb, 0xd0, 0xfa, 0xd8, 0xef, 0x0e, 0xd9, 0x3a, 0xfa, 0xfa, 0xf1, 0x8b, 0x1d, 0xeb, 0x9b, 0x9d,
	0x8a, 0xf6, 0x48, 0x5b, 0x45, 0x4e, 0x6b, 0xa7, 0x75, 0xd0, 0x1d, 0xee, 0x56, 0xa5, 0xdb, 0x3d,
	0x85, 0x4c, 0xbf, 0x66, 0x2a, 0x62, 0xbf, 0x7e, 0xba, 0x9b, 0x64, 0xe3, 0x64, 0x74, 0xc8, 0x07,
	0x82, 0x5f, 0x8d, 0xaf, 0x7b, 0xfd, 0x9b, 0xfd, 0xa0, 0xad, 0x22, 0xba, 0x47, 0x96, 0x66, 0x42,
	0x4b, 0xa7, 0x8d, 0xa6, 0x07, 0x55, 0xe9, 0xae, 0x59, 0x93, 0xa9, 0xb2, 0x00, 0x45, 0xfa, 0x96,
	0xdc, 0xcf, 0x64, 0x0a, 0xf5, 0x6f, 0xde, 0x47, 0x4e, 0x07, 0x69, 0xa7, 0x2a, 0xdd, 0x47, 0x96,
	0x5e, 0x90, 0x59, 0xb0, 0x88, 0xd3, 0x37, 0x64, 0x5d, 0xa4, 0xe9, 0x3f, 0xfb, 0x12, 0xda, 0x9f,
	0x54, 0xa5, 0xfb, 0xd0, 0xda, 0x9b, 0x2a, 0x0b, 0x16, 0x60, 0xfa, 0x8a, 0xac, 0xd5, 0x5b, 0xff,
	0x6c, 0x06, 0x5d, 0x46, 0xef, 0xe3, 0xaa, 0x74, 0xa9, 0xf5, 0x36, 0x44, 0x16, 0x34, 0x51, 0xf3,
	0x36, 0x2d, 0x35, 0x38, 0x2b, 0x77, 0xdf, 0x66, 0xaa, 0x2c, 0x40, 0x91, 0x1e, 0x91, 0xd5, 0x30,
	0x93, 0x66, 0xb5, 0xce, 0x2a, 0x72, 0x83, 0xaa, 0x74, 0x37, 0x2c, 0x57, 0x0b, 0x66, 0x7b, 0x7b,
	0x64, 0xf7, 0x64, 0x24, 0xf8, 0xd5, 0x3b, 0x7e, 0x7c, 0xc8, 0x07, 0xe3, 0x91, 0x77, 0x1b, 0x4f,
	0xf8, 0xf8, 0xfa, 0xc5, 0xf3, 0x7e, 0xef, 0x66, 0x3f, 0xf8, 0xdb, 0x89, 0x3e, 0x23, 0xcb, 0x61,
	0xa1, 0xb4, 0x74, 0xee, 0xed, 0xb4, 0x0e, 0x3a, 0xc3, 0xcd, 0xaa, 0x74, 0xd7, 0xeb, 0x96, 0xa6,
	0xcc, 0x02, 0x2b, 0x1b, 0x4e, 0x23, 0xd7, 0xbd, 0xcb, 0xe9, 0x9a, 0xc3, 0xef, 0xf0, 0xd3, 0xf1,
	0x87, 0x18, 0x3c, 0x29, 0xf2, 0x4b, 0x48, 0x73, 0x2f, 0x81, 0x50, 0x24, 0x7e, 0x08, 0xb3, 0x22,
	0x13, 0x61, 0x91, 0xdb, 0x13, 0x33, 0x1b, 0xe7, 0x1a, 0x22, 0x99, 0xe4, 0x7e, 0x0d, 0xfa, 0x98,
	0xfa, 0xff, 0x1d, 0xe9, 0x74, 0x05, 0x1d, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x78,
	0xe3, 0x58, 0xdb, 0x02, 0x00, 0x00,
}
