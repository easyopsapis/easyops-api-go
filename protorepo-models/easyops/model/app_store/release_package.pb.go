// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: release_package.proto

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
//发布的包信息
type ReleasePackage struct {
	//
	//包名称称+版本组合的唯一id
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//包名
	PackageName string `protobuf:"bytes,2,opt,name=packageName,proto3" json:"packageName" form:"packageName"`
	//
	//版本名
	VersionName string `protobuf:"bytes,3,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//包Id
	PackageId string `protobuf:"bytes,4,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId string `protobuf:"bytes,5,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//下载地址
	DownloadUrl string `protobuf:"bytes,6,opt,name=downloadUrl,proto3" json:"downloadUrl" form:"downloadUrl"`
	//
	//程序包所属分组
	Group                string   `protobuf:"bytes,7,opt,name=group,proto3" json:"group" form:"group"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleasePackage) Reset()         { *m = ReleasePackage{} }
func (m *ReleasePackage) String() string { return proto.CompactTextString(m) }
func (*ReleasePackage) ProtoMessage()    {}
func (*ReleasePackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_26165800255565f1, []int{0}
}
func (m *ReleasePackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasePackage.Unmarshal(m, b)
}
func (m *ReleasePackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasePackage.Marshal(b, m, deterministic)
}
func (m *ReleasePackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasePackage.Merge(m, src)
}
func (m *ReleasePackage) XXX_Size() int {
	return xxx_messageInfo_ReleasePackage.Size(m)
}
func (m *ReleasePackage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasePackage.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasePackage proto.InternalMessageInfo

func (m *ReleasePackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReleasePackage) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ReleasePackage) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *ReleasePackage) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *ReleasePackage) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ReleasePackage) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *ReleasePackage) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func init() {
	proto.RegisterType((*ReleasePackage)(nil), "app_store.ReleasePackage")
}

func init() { proto.RegisterFile("release_package.proto", fileDescriptor_26165800255565f1) }

var fileDescriptor_26165800255565f1 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x4a, 0xe3, 0x40,
	0x14, 0x86, 0xe9, 0x6e, 0xdb, 0xa5, 0xd3, 0x65, 0x29, 0x81, 0x5d, 0x42, 0x61, 0x89, 0x04, 0x11,
	0x2f, 0x4c, 0xa6, 0x5a, 0x50, 0xeb, 0x9d, 0x15, 0x85, 0xde, 0x88, 0x04, 0x7a, 0xd3, 0xa2, 0x65,
	0x9a, 0x4c, 0x63, 0x68, 0x92, 0x13, 0x66, 0xd2, 0x16, 0x15, 0x5f, 0xc6, 0x07, 0x8b, 0xe0, 0x23,
	0xe4, 0x09, 0x24, 0x33, 0x31, 0x8d, 0xbd, 0xca, 0xdd, 0x39, 0xf3, 0xff, 0xff, 0x77, 0x4e, 0xc2,
	0x41, 0x7f, 0x19, 0xf5, 0x29, 0xe1, 0x74, 0x16, 0x11, 0x7b, 0x49, 0x5c, 0x6a, 0x46, 0x0c, 0x62,
	0x50, 0x5a, 0x24, 0x8a, 0x66, 0x3c, 0x06, 0x46, 0xbb, 0x86, 0xeb, 0xc5, 0x8f, 0xab, 0xb9, 0x69,
	0x43, 0x80, 0x5d, 0x70, 0x01, 0x0b, 0xc7, 0x7c, 0xb5, 0x10, 0x9d, 0x68, 0x44, 0x25, 0x93, 0xdd,
	0xd3, 0x92, 0x3d, 0xd8, 0x78, 0xf1, 0x12, 0x36, 0xd8, 0x05, 0x43, 0x88, 0xc6, 0x9a, 0xf8, 0x9e,
	0x43, 0x62, 0x60, 0x1c, 0x17, 0xa5, 0xcc, 0xe9, 0x6f, 0x75, 0xf4, 0xc7, 0x92, 0xbb, 0xdc, 0xc9,
	0x55, 0x94, 0x1b, 0x54, 0x0f, 0x49, 0x40, 0xd5, 0xda, 0x5e, 0xed, 0xb0, 0x35, 0x3c, 0x49, 0x13,
	0xad, 0xbd, 0x00, 0x16, 0x5c, 0xe8, 0xd9, 0xab, 0xfe, 0xf1, 0xae, 0x69, 0xe8, 0xff, 0xc3, 0x94,
	0x18, 0xcf, 0x97, 0xc6, 0x64, 0x76, 0x3f, 0xed, 0x19, 0x83, 0xaf, 0xfa, 0xa5, 0x77, 0xd4, 0x3f,
	0x7e, 0xdd, 0xb7, 0x44, 0x5e, 0x99, 0xa2, 0x76, 0xfe, 0x75, 0xb7, 0x19, 0xee, 0x87, 0xc0, 0x0d,
	0xd2, 0x44, 0x53, 0x24, 0xae, 0x24, 0x56, 0xa2, 0x96, 0x69, 0x19, 0x7c, 0x4d, 0x19, 0xf7, 0x20,
	0x14, 0xf0, 0x9f, 0xbb, 0xf0, 0x92, 0x58, 0x0d, 0x5e, 0x0a, 0x28, 0x63, 0xd4, 0xca, 0x67, 0x8d,
	0x1c, 0xb5, 0x2e, 0xd0, 0x67, 0x69, 0xa2, 0x75, 0xbe, 0xed, 0x3d, 0x72, 0x2a, 0x81, 0xb7, 0xa4,
	0x0c, 0x9b, 0x4f, 0x19, 0x39, 0x6a, 0x63, 0x17, 0x5b, 0x48, 0xd5, 0xb0, 0x85, 0x5d, 0x39, 0x47,
	0x6d, 0x07, 0x36, 0xa1, 0x0f, 0xc4, 0x19, 0x33, 0x5f, 0x6d, 0x0a, 0xf0, 0xbf, 0xed, 0xaf, 0x28,
	0x89, 0xba, 0x55, 0xb6, 0x2a, 0x07, 0xa8, 0xe1, 0x32, 0x58, 0x45, 0xea, 0x2f, 0x91, 0xe9, 0xa4,
	0x89, 0xf6, 0x5b, 0x66, 0xc4, 0xb3, 0x6e, 0x49, 0x79, 0x78, 0x3d, 0xb9, 0x72, 0xc1, 0xa4, 0x84,
	0x3f, 0x41, 0xc4, 0x4d, 0x1f, 0x6c, 0xe2, 0x63, 0x1b, 0xc2, 0x98, 0x11, 0x3b, 0xe6, 0xf2, 0x30,
	0x19, 0x8d, 0xc0, 0x08, 0xc0, 0xa1, 0x3e, 0xc7, 0xb9, 0x11, 0x8b, 0x16, 0x17, 0x27, 0x3d, 0x6f,
	0x0a, 0x67, 0xff, 0x33, 0x00, 0x00, 0xff, 0xff, 0x34, 0x98, 0x40, 0xb2, 0xfd, 0x02, 0x00, 0x00,
}
