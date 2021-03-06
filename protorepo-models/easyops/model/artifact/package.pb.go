// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: package.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//包模型
type Package struct {
	//
	//包Id
	PackageId string `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//包名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//版本类型 1 程序包,  2 配置包,  4 文件包
	Type int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type" form:"type"`
	//
	//包分类
	CId int32 `protobuf:"varint,4,opt,name=cId,proto3" json:"cId" form:"cId"`
	//
	//包文件源
	Source string `protobuf:"bytes,5,opt,name=source,proto3" json:"source" form:"source"`
	//
	//repoId
	RepoId string `protobuf:"bytes,6,opt,name=repoId,proto3" json:"repoId" form:"repoId"`
	//
	//repoPath
	RepoPath string `protobuf:"bytes,7,opt,name=repoPath,proto3" json:"repoPath" form:"repoPath"`
	//
	//备注说明
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//创建者
	Creator string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//客户唯一编号
	Org int32 `protobuf:"varint,10,opt,name=org,proto3" json:"org" form:"org"`
	//
	//包分类标签
	Category string `protobuf:"bytes,11,opt,name=category,proto3" json:"category" form:"category"`
	//
	//包图标
	Icon string `protobuf:"bytes,12,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//包图标样式(颜色)
	Style string `protobuf:"bytes,13,opt,name=style,proto3" json:"style" form:"style"`
	//
	//ctime
	Ctime string `protobuf:"bytes,14,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//mtime
	Mtime string `protobuf:"bytes,15,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//authUsers
	AuthUsers string `protobuf:"bytes,16,opt,name=authUsers,proto3" json:"authUsers" form:"authUsers"`
	//
	//安装路径
	InstallPath string `protobuf:"bytes,17,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	//
	//平台
	Platform             string   `protobuf:"bytes,18,opt,name=platform,proto3" json:"platform" form:"platform"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{0}
}
func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Package) GetCId() int32 {
	if m != nil {
		return m.CId
	}
	return 0
}

func (m *Package) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Package) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *Package) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *Package) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Package) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Package) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *Package) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Package) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Package) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *Package) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Package) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *Package) GetAuthUsers() string {
	if m != nil {
		return m.AuthUsers
	}
	return ""
}

func (m *Package) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *Package) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*Package)(nil), "artifact.Package")
}

func init() { proto.RegisterFile("package.proto", fileDescriptor_ae8103ff0e06fb71) }

var fileDescriptor_ae8103ff0e06fb71 = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xcb, 0x8e, 0xe3, 0x44,
	0x14, 0x86, 0xdb, 0xd3, 0xe9, 0x5b, 0xf5, 0xf4, 0x65, 0x0a, 0x10, 0xa5, 0xd9, 0xb8, 0xf1, 0x44,
	0x23, 0xbb, 0x43, 0xc5, 0xb9, 0x00, 0x62, 0x02, 0x22, 0x9a, 0x08, 0x81, 0xb2, 0x1b, 0x3c, 0xb0,
	0xa0, 0x6d, 0x07, 0xaa, 0x1d, 0xc7, 0x1d, 0xb5, 0x9d, 0xb2, 0xec, 0x0a, 0x43, 0x18, 0xf7, 0x92,
	0x05, 0xaf, 0x82, 0xc4, 0x4b, 0xf0, 0x02, 0xac, 0xd9, 0x18, 0x89, 0x47, 0xf0, 0x13, 0xa0, 0x3a,
	0x76, 0xe2, 0x68, 0x86, 0x05, 0x48, 0x6c, 0x58, 0xa5, 0xce, 0xf9, 0xbf, 0x53, 0xe7, 0x52, 0x55,
	0x31, 0x3a, 0x89, 0x99, 0x77, 0xcb, 0x02, 0xbf, 0x1d, 0x27, 0x5c, 0x70, 0x7c, 0xc8, 0x12, 0x31,
	0x9f, 0x31, 0x4f, 0x3c, 0xa4, 0xc1, 0x5c, 0xdc, 0x2c, 0xaf, 0xdb, 0x1e, 0x8f, 0xcc, 0x80, 0x07,
	0xdc, 0x04, 0xe0, 0x7a, 0x39, 0x03, 0x0b, 0x0c, 0x58, 0x95, 0x81, 0x0f, 0x3f, 0xd8, 0xc2, 0xa3,
	0x17, 0x73, 0x71, 0xcb, 0x5f, 0x98, 0x01, 0xa7, 0x20, 0xd2, 0xef, 0x58, 0x38, 0x9f, 0x32, 0xc1,
	0x93, 0xd4, 0xdc, 0x2c, 0xcb, 0x38, 0xed, 0xb7, 0x63, 0x74, 0xf0, 0xac, 0x2c, 0x01, 0xff, 0xa4,
	0xa0, 0xa3, 0xaa, 0x9c, 0xf1, 0x94, 0x28, 0x17, 0x8a, 0x7e, 0x34, 0xba, 0x2d, 0x72, 0xf5, 0x7c,
	0xc6, 0x93, 0x68, 0xa0, 0x6d, 0x24, 0xed, 0xcf, 0x3f, 0xd4, 0xe7, 0xe8, 0x8b, 0x89, 0xcd, 0xe8,
	0xec, 0x29, 0xfd, 0xac, 0x43, 0x9f, 0xb8, 0x2f, 0x3f, 0xbc, 0xa3, 0xc3, 0x6d, 0xfb, 0xbd, 0x7f,
	0x69, 0x77, 0x7b, 0x77, 0x4d, 0xab, 0xce, 0x8e, 0x5b, 0xa8, 0xb1, 0x60, 0x91, 0x4f, 0xee, 0x41,
	0x15, 0x6f, 0x17, 0xb9, 0x7a, 0x5c, 0x56, 0x21, 0xbd, 0xb2, 0x80, 0x46, 0xbc, 0xf3, 0x7d, 0xdb,
	0x02, 0x08, 0x3f, 0x42, 0x0d, 0xb1, 0x8a, 0x7d, 0xb2, 0x7b, 0xa1, 0xe8, 0x7b, 0xa3, 0xb3, 0x1a,
	0x96, 0x5e, 0xcd, 0x02, 0x11, 0x5f, 0xa0, 0x5d, 0x6f, 0x3c, 0x25, 0x0d, 0x60, 0x4e, 0x8b, 0x5c,
	0x45, 0x25, 0xe3, 0x8d, 0xa7, 0x9a, 0x25, 0x25, 0x6c, 0xa0, 0xfd, 0x94, 0x2f, 0x13, 0xcf, 0x27,
	0x7b, 0x90, 0xf5, 0x41, 0x91, 0xab, 0x27, 0x25, 0x54, 0xfa, 0x35, 0xab, 0x02, 0x24, 0x9a, 0xf8,
	0x31, 0x1f, 0x4f, 0xc9, 0xfe, 0xab, 0x68, 0xe9, 0xd7, 0xac, 0x0a, 0xc0, 0x19, 0x3a, 0x94, 0xab,
	0x67, 0x4c, 0xdc, 0x90, 0x03, 0x80, 0xbf, 0x2d, 0x72, 0xf5, 0xac, 0x86, 0xa5, 0x22, 0x3b, 0x1a,
	0xa3, 0xcf, 0x27, 0xba, 0xee, 0x98, 0xf6, 0xc4, 0x31, 0x07, 0xce, 0xa5, 0x33, 0xd4, 0xb4, 0x8f,
	0x3f, 0x71, 0x32, 0x27, 0x71, 0x16, 0x6e, 0xcb, 0x68, 0x19, 0x99, 0xee, 0x98, 0x46, 0x66, 0x33,
	0xfa, 0xc3, 0x53, 0x7a, 0xe5, 0x0e, 0x74, 0xc7, 0xb1, 0x27, 0x8e, 0xf3, 0x3a, 0x79, 0xd9, 0xb4,
	0x36, 0x19, 0xb1, 0x81, 0x1a, 0x91, 0x1f, 0x71, 0x72, 0x08, 0x99, 0xdf, 0xaa, 0x47, 0x23, 0xbd,
	0x32, 0xeb, 0xbd, 0x78, 0xc7, 0x02, 0x04, 0x3f, 0x47, 0x07, 0x5e, 0xe2, 0xcb, 0xbb, 0x41, 0x8e,
	0x80, 0x7e, 0x52, 0xe4, 0xea, 0x69, 0x35, 0xa4, 0x52, 0x90, 0x01, 0x8f, 0xd0, 0x3b, 0x93, 0xaa,
	0x08, 0x79, 0x72, 0x76, 0x7b, 0xb3, 0xfe, 0x86, 0xba, 0x2f, 0x7b, 0xef, 0xf6, 0xbb, 0x77, 0x4d,
	0x6b, 0xbd, 0x13, 0x7e, 0x8c, 0x76, 0x79, 0x12, 0x10, 0x04, 0x53, 0x7f, 0xb3, 0x9e, 0x3a, 0x4f,
	0x02, 0xc8, 0x7e, 0xbe, 0x63, 0x49, 0x00, 0x9b, 0xe8, 0xd0, 0x63, 0xc2, 0x0f, 0x78, 0xb2, 0x22,
	0xc7, 0x90, 0xfd, 0x8d, 0x7a, 0x4a, 0x6b, 0x45, 0xb3, 0x36, 0x90, 0x3c, 0xf3, 0xb9, 0xc7, 0x17,
	0xe4, 0x3e, 0xc0, 0x5b, 0x67, 0x2e, 0xbd, 0x9a, 0x05, 0x22, 0x7e, 0x8c, 0xf6, 0x52, 0xb1, 0x0a,
	0x7d, 0x72, 0x02, 0xd4, 0x79, 0x91, 0xab, 0xf7, 0xab, 0x03, 0x95, 0x6e, 0xcd, 0x2a, 0x65, 0xfc,
	0xbb, 0x82, 0xf6, 0x3c, 0x31, 0x8f, 0x7c, 0x72, 0x0a, 0xe0, 0xaf, 0x4a, 0x4d, 0x82, 0x5f, 0xd6,
	0xfa, 0x8b, 0x82, 0x7e, 0x56, 0x26, 0xba, 0x3e, 0x1c, 0xd8, 0x5d, 0xd9, 0xb8, 0xec, 0xfe, 0xd2,
	0x18, 0xda, 0xd5, 0x7d, 0x36, 0xa8, 0xde, 0xb5, 0x3b, 0xb4, 0xe7, 0x66, 0x1d, 0xd0, 0x0d, 0xaa,
	0xf7, 0xed, 0x0e, 0xed, 0xae, 0xed, 0xcc, 0xee, 0xd2, 0x5e, 0x19, 0x65, 0xd8, 0x5f, 0x5e, 0xb8,
	0x7a, 0xcf, 0xee, 0xd0, 0xbe, 0x9b, 0x01, 0x53, 0xba, 0x07, 0xba, 0xdd, 0xa1, 0xef, 0xaf, 0x8d,
	0x7a, 0xad, 0x3b, 0x6d, 0xf8, 0x6d, 0x19, 0x43, 0xfd, 0x2a, 0xb3, 0x5b, 0xd4, 0xd5, 0x87, 0x83,
	0xbf, 0x09, 0xdf, 0x8a, 0x1e, 0x36, 0xad, 0xb2, 0x23, 0xe8, 0x2d, 0x82, 0xde, 0xce, 0x5e, 0xeb,
	0x2d, 0xfa, 0xdf, 0xf6, 0x06, 0x95, 0xe3, 0xaf, 0xd1, 0x11, 0x5b, 0x8a, 0x9b, 0xaf, 0x52, 0x3f,
	0x49, 0xc9, 0x39, 0xb4, 0xf7, 0x51, 0xfd, 0x87, 0xb5, 0x91, 0xfe, 0xf1, 0xb5, 0xad, 0x77, 0xc3,
	0x3f, 0x2a, 0xe8, 0x78, 0xbe, 0x48, 0x05, 0x0b, 0x43, 0x78, 0xba, 0x0f, 0x60, 0x77, 0xaf, 0xc8,
	0x55, 0x5c, 0xdd, 0xb3, 0x5a, 0xfc, 0x8f, 0x5f, 0xef, 0x76, 0x5e, 0xf9, 0x30, 0xe2, 0x90, 0x09,
	0x99, 0x93, 0xe0, 0x57, 0x1f, 0xc6, 0x5a, 0xd1, 0xac, 0x0d, 0x34, 0xfa, 0xf4, 0x6a, 0x14, 0xf0,
	0xb6, 0xcf, 0xd2, 0x15, 0x8f, 0xd3, 0x76, 0xc8, 0x3d, 0x16, 0x9a, 0x1e, 0x5f, 0x88, 0x84, 0x79,
	0x22, 0x2d, 0xbf, 0x22, 0xf2, 0xdf, 0x81, 0x46, 0x7c, 0xea, 0x87, 0xa9, 0x59, 0x81, 0x26, 0x98,
	0xe6, 0xfa, 0xf3, 0x73, 0xbd, 0x0f, 0x60, 0xff, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xc8,
	0x11, 0x51, 0xa0, 0x06, 0x00, 0x00,
}
