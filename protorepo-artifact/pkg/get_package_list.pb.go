// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_package_list.proto

package pkg

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
//GetPackageList返回
type GetPackageListResponse struct {
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
	Platform string `protobuf:"bytes,18,opt,name=platform,proto3" json:"platform" form:"platform"`
	//
	//最新版本信息
	LastVersionInfo *GetPackageListResponse_LastVersionInfo `protobuf:"bytes,19,opt,name=lastVersionInfo,proto3" json:"lastVersionInfo" form:"lastVersionInfo"`
	//
	//包实例数量
	InstanceCount        int32    `protobuf:"varint,20,opt,name=instanceCount,proto3" json:"instanceCount" form:"instanceCount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackageListResponse) Reset()         { *m = GetPackageListResponse{} }
func (m *GetPackageListResponse) String() string { return proto.CompactTextString(m) }
func (*GetPackageListResponse) ProtoMessage()    {}
func (*GetPackageListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2739fedb2fb1db51, []int{0}
}
func (m *GetPackageListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageListResponse.Unmarshal(m, b)
}
func (m *GetPackageListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageListResponse.Marshal(b, m, deterministic)
}
func (m *GetPackageListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageListResponse.Merge(m, src)
}
func (m *GetPackageListResponse) XXX_Size() int {
	return xxx_messageInfo_GetPackageListResponse.Size(m)
}
func (m *GetPackageListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageListResponse proto.InternalMessageInfo

func (m *GetPackageListResponse) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetPackageListResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPackageListResponse) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GetPackageListResponse) GetCId() int32 {
	if m != nil {
		return m.CId
	}
	return 0
}

func (m *GetPackageListResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *GetPackageListResponse) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *GetPackageListResponse) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *GetPackageListResponse) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *GetPackageListResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetPackageListResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetPackageListResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetPackageListResponse) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GetPackageListResponse) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *GetPackageListResponse) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *GetPackageListResponse) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *GetPackageListResponse) GetAuthUsers() string {
	if m != nil {
		return m.AuthUsers
	}
	return ""
}

func (m *GetPackageListResponse) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *GetPackageListResponse) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *GetPackageListResponse) GetLastVersionInfo() *GetPackageListResponse_LastVersionInfo {
	if m != nil {
		return m.LastVersionInfo
	}
	return nil
}

func (m *GetPackageListResponse) GetInstanceCount() int32 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

type GetPackageListResponse_LastVersionInfo struct {
	//
	//创建时间
	Ctime string `protobuf:"bytes,1,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//版本类型 1 开发, 3 测试, 7 预发布, 15 生产
	EnvType int32 `protobuf:"varint,2,opt,name=env_type,json=envType,proto3" json:"env_type" form:"env_type"`
	//
	//包名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//版本Id
	VersionId            string   `protobuf:"bytes,4,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackageListResponse_LastVersionInfo) Reset() {
	*m = GetPackageListResponse_LastVersionInfo{}
}
func (m *GetPackageListResponse_LastVersionInfo) String() string { return proto.CompactTextString(m) }
func (*GetPackageListResponse_LastVersionInfo) ProtoMessage()    {}
func (*GetPackageListResponse_LastVersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2739fedb2fb1db51, []int{0, 0}
}
func (m *GetPackageListResponse_LastVersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageListResponse_LastVersionInfo.Unmarshal(m, b)
}
func (m *GetPackageListResponse_LastVersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageListResponse_LastVersionInfo.Marshal(b, m, deterministic)
}
func (m *GetPackageListResponse_LastVersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageListResponse_LastVersionInfo.Merge(m, src)
}
func (m *GetPackageListResponse_LastVersionInfo) XXX_Size() int {
	return xxx_messageInfo_GetPackageListResponse_LastVersionInfo.Size(m)
}
func (m *GetPackageListResponse_LastVersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageListResponse_LastVersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageListResponse_LastVersionInfo proto.InternalMessageInfo

func (m *GetPackageListResponse_LastVersionInfo) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *GetPackageListResponse_LastVersionInfo) GetEnvType() int32 {
	if m != nil {
		return m.EnvType
	}
	return 0
}

func (m *GetPackageListResponse_LastVersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPackageListResponse_LastVersionInfo) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//GetPackageListApi返回
type GetPackageListResponseWrapper struct {
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
	Data                 *GetPackageListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetPackageListResponseWrapper) Reset()         { *m = GetPackageListResponseWrapper{} }
func (m *GetPackageListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetPackageListResponseWrapper) ProtoMessage()    {}
func (*GetPackageListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2739fedb2fb1db51, []int{1}
}
func (m *GetPackageListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageListResponseWrapper.Unmarshal(m, b)
}
func (m *GetPackageListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetPackageListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageListResponseWrapper.Merge(m, src)
}
func (m *GetPackageListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetPackageListResponseWrapper.Size(m)
}
func (m *GetPackageListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageListResponseWrapper proto.InternalMessageInfo

func (m *GetPackageListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPackageListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetPackageListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetPackageListResponseWrapper) GetData() *GetPackageListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPackageListResponse)(nil), "pkg.GetPackageListResponse")
	proto.RegisterType((*GetPackageListResponse_LastVersionInfo)(nil), "pkg.GetPackageListResponse.LastVersionInfo")
	proto.RegisterType((*GetPackageListResponseWrapper)(nil), "pkg.GetPackageListResponseWrapper")
}

func init() { proto.RegisterFile("get_package_list.proto", fileDescriptor_2739fedb2fb1db51) }

var fileDescriptor_2739fedb2fb1db51 = []byte{
	// 998 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x9b, 0xf4, 0x6f, 0xb2, 0x6d, 0xb3, 0xb3, 0xa5, 0x58, 0x45, 0xc8, 0x95, 0xa9, 0x16,
	0xbb, 0x59, 0x27, 0x69, 0xba, 0xac, 0x76, 0x23, 0x44, 0xd8, 0x22, 0x58, 0x45, 0x2a, 0xd2, 0x6a,
	0x54, 0x40, 0xaa, 0x9d, 0x14, 0xd7, 0x99, 0xba, 0x56, 0x13, 0x8f, 0xb1, 0x9d, 0x76, 0xbb, 0x9b,
	0x5e, 0x72, 0xc3, 0xa3, 0x20, 0xf1, 0x12, 0xbc, 0x04, 0xe2, 0xc6, 0x08, 0xae, 0xb8, 0xf6, 0x0b,
	0x80, 0xe6, 0xd8, 0x89, 0x9d, 0x6c, 0xb5, 0x02, 0x09, 0x2e, 0x20, 0x37, 0x99, 0x39, 0xdf, 0x77,
	0xe6, 0xfc, 0xcc, 0x39, 0x73, 0x8c, 0x36, 0x6d, 0x1a, 0x9e, 0x78, 0xa6, 0x75, 0x61, 0xda, 0xf4,
	0xa4, 0xef, 0x04, 0x61, 0xd5, 0xf3, 0x59, 0xc8, 0x70, 0xc1, 0xbb, 0xb0, 0xb7, 0x34, 0xdb, 0x09,
	0xcf, 0x87, 0xa7, 0x55, 0x8b, 0x0d, 0x6a, 0x36, 0xb3, 0x59, 0x0d, 0xb0, 0xd3, 0xe1, 0x19, 0xec,
	0x60, 0x03, 0xab, 0x44, 0x67, 0xeb, 0x51, 0x8e, 0x3e, 0xb8, 0x72, 0xc2, 0x0b, 0x76, 0x55, 0xb3,
	0x99, 0x06, 0xa0, 0x76, 0x69, 0xf6, 0x9d, 0x9e, 0x19, 0x32, 0x3f, 0xa8, 0x4d, 0x96, 0x89, 0x9e,
	0xfc, 0x6b, 0x19, 0x6d, 0x3e, 0xa3, 0xe1, 0xf3, 0xc4, 0x8b, 0x43, 0x27, 0x08, 0x09, 0x0d, 0x3c,
	0xe6, 0x06, 0x14, 0x7f, 0x27, 0xa0, 0x95, 0xd4, 0xbb, 0x76, 0x4f, 0x14, 0xb6, 0x05, 0x65, 0xe5,
	0xa0, 0x1f, 0x47, 0x52, 0xf9, 0x8c, 0xf9, 0x83, 0xa6, 0x3c, 0x81, 0xe4, 0xdf, 0x7e, 0x91, 0x8e,
	0x10, 0xe9, 0xea, 0x2d, 0x53, 0x7b, 0xf9, 0x54, 0x3b, 0xae, 0x6b, 0x4f, 0x3a, 0xaf, 0x1e, 0xdf,
	0x68, 0x53, 0xfb, 0x87, 0x7f, 0x73, 0xbf, 0xd7, 0xb8, 0xd9, 0x21, 0x99, 0x79, 0x5c, 0x41, 0x45,
	0xd7, 0x1c, 0x50, 0x71, 0x1e, 0xdc, 0x78, 0x3b, 0x8e, 0xa4, 0x52, 0xe2, 0x06, 0x97, 0x72, 0x0f,
	0x8a, 0xde, 0xdc, 0x8b, 0x2a, 0x01, 0x12, 0x7e, 0x0f, 0x15, 0xc3, 0x6b, 0x8f, 0x8a, 0x85, 0x6d,
	0x41, 0x59, 0x38, 0x58, 0xcf, 0xc8, 0x5c, 0x2a, 0x13, 0x00, 0xf1, 0x36, 0x2a, 0x58, 0xed, 0x9e,
	0x58, 0x04, 0xce, 0x5a, 0x1c, 0x49, 0x28, 0xe1, 0x58, 0xed, 0x9e, 0x4c, 0x38, 0x84, 0x55, 0xb4,
	0x18, 0xb0, 0xa1, 0x6f, 0x51, 0x71, 0x01, 0xac, 0xde, 0x8d, 0x23, 0x69, 0x35, 0x21, 0x25, 0x72,
	0x99, 0xa4, 0x04, 0x4e, 0xf5, 0xa9, 0xc7, 0xda, 0x3d, 0x71, 0x71, 0x96, 0x9a, 0xc8, 0x65, 0x92,
	0x12, 0xf0, 0x08, 0x2d, 0xf3, 0xd5, 0x73, 0x33, 0x3c, 0x17, 0x97, 0x80, 0xfc, 0x75, 0x1c, 0x49,
	0xeb, 0x19, 0x99, 0x23, 0x3c, 0xa2, 0x36, 0x7a, 0xd6, 0x55, 0x14, 0xa3, 0xa6, 0x77, 0x8d, 0x5a,
	0xd3, 0xd8, 0x35, 0x5a, 0xb2, 0xfc, 0xe1, 0x47, 0xc6, 0xc8, 0xf0, 0x0d, 0xb7, 0x53, 0x51, 0x2b,
	0xea, 0x48, 0x31, 0x6a, 0xea, 0x48, 0x4f, 0x92, 0xd6, 0x69, 0x2a, 0x86, 0xa1, 0x77, 0x0d, 0xe3,
	0x75, 0xe6, 0xee, 0x0e, 0x99, 0x58, 0xc4, 0x2a, 0x2a, 0x0e, 0xe8, 0x80, 0x89, 0xcb, 0x60, 0xf9,
	0xad, 0x2c, 0x35, 0x5c, 0xca, 0xad, 0xce, 0x7b, 0x73, 0x04, 0x28, 0xf8, 0x33, 0xb4, 0x64, 0xf9,
	0x94, 0xd7, 0x8a, 0xb8, 0x02, 0xec, 0x07, 0x71, 0x24, 0xad, 0xa5, 0x49, 0x4a, 0x00, 0xae, 0xb0,
	0x89, 0x36, 0xba, 0xfa, 0x53, 0xed, 0xd8, 0xd4, 0x5e, 0x9e, 0x68, 0x1d, 0xe3, 0xea, 0xd5, 0xfe,
	0x83, 0x47, 0x0f, 0x6f, 0x76, 0xc8, 0x58, 0x19, 0xdf, 0x47, 0x05, 0xe6, 0xdb, 0x22, 0x82, 0x44,
	0x6f, 0x64, 0x89, 0x66, 0xbe, 0x0d, 0x06, 0xcb, 0x73, 0x84, 0x13, 0x70, 0x0d, 0x2d, 0x5b, 0x66,
	0x48, 0x6d, 0xe6, 0x5f, 0x8b, 0x25, 0x30, 0x78, 0x2f, 0x4b, 0xcc, 0x18, 0x91, 0xc9, 0x84, 0xc4,
	0xaf, 0xd9, 0xb1, 0x98, 0x2b, 0xde, 0x01, 0x72, 0xee, 0x9a, 0xb9, 0x54, 0x26, 0x00, 0xe2, 0xfb,
	0x68, 0x21, 0x08, 0xaf, 0xfb, 0x54, 0x5c, 0x05, 0x56, 0x39, 0x8e, 0xa4, 0x3b, 0xe9, 0x1d, 0x72,
	0xb1, 0x4c, 0x12, 0x18, 0xff, 0x2c, 0xa0, 0x05, 0x2b, 0x74, 0x06, 0x54, 0x5c, 0x03, 0xe2, 0x8f,
	0x42, 0xc6, 0x04, 0x39, 0xf7, 0xf5, 0x07, 0x01, 0x7d, 0x2f, 0x74, 0x15, 0xa5, 0xd5, 0xd4, 0xf7,
	0xb4, 0x27, 0x1d, 0x9d, 0x97, 0xea, 0xae, 0xda, 0xd2, 0xd3, 0x12, 0x56, 0x35, 0x65, 0x4f, 0xaf,
	0x6b, 0x8d, 0xce, 0xa8, 0x0e, 0xb8, 0xaa, 0x29, 0xfb, 0x7a, 0x5d, 0xdb, 0x1b, 0xef, 0x47, 0xfa,
	0x9e, 0xd6, 0x48, 0xb4, 0x54, 0xfd, 0x68, 0xbb, 0xa3, 0x34, 0xf4, 0xba, 0xb6, 0xdf, 0x19, 0x01,
	0x27, 0x11, 0x37, 0x15, 0xbd, 0xae, 0x7d, 0x30, 0xde, 0x64, 0x6b, 0xc5, 0xa8, 0xc2, 0x7f, 0x45,
	0x6d, 0x29, 0xc7, 0x23, 0xbd, 0xa2, 0x75, 0x94, 0x56, 0xf3, 0x16, 0xf5, 0x9c, 0x76, 0x6b, 0x87,
	0x24, 0x11, 0x41, 0x6c, 0x03, 0x88, 0x6d, 0xfd, 0xb5, 0xd8, 0x06, 0xff, 0xd9, 0xd8, 0xc0, 0x73,
	0x7c, 0x88, 0x56, 0xcc, 0x61, 0x78, 0xfe, 0x45, 0x40, 0xfd, 0x40, 0x2c, 0x43, 0x78, 0xd5, 0xec,
	0x91, 0x9a, 0x40, 0x6f, 0xaa, 0xd4, 0xec, 0x00, 0xfc, 0xad, 0x80, 0x4a, 0x8e, 0x1b, 0x84, 0x66,
	0xbf, 0x0f, 0x0d, 0x7a, 0x17, 0x0e, 0xb4, 0xe2, 0x48, 0xc2, 0x69, 0x69, 0x65, 0xe0, 0x3f, 0xdc,
	0xa3, 0x79, 0xbb, 0xbc, 0x17, 0xbc, 0xbe, 0x19, 0x72, 0x9b, 0x22, 0x9e, 0xed, 0x85, 0x31, 0x22,
	0x93, 0x09, 0x09, 0x7f, 0x83, 0xd6, 0xfb, 0x66, 0x10, 0x7e, 0x49, 0xfd, 0xc0, 0x61, 0x6e, 0xdb,
	0x3d, 0x63, 0xe2, 0xbd, 0x6d, 0x41, 0x29, 0x35, 0x2a, 0x55, 0xef, 0xc2, 0xae, 0xde, 0xfe, 0xc4,
	0x57, 0x0f, 0xa7, 0x55, 0x0e, 0xb6, 0xe2, 0x48, 0xda, 0x4c, 0x8c, 0xcc, 0x9c, 0x26, 0x93, 0xd9,
	0xf3, 0xf1, 0xe7, 0x68, 0x15, 0x5c, 0x76, 0x2d, 0xfa, 0x09, 0x1b, 0xba, 0xa1, 0xb8, 0x01, 0x1d,
	0xfe, 0x7e, 0x1c, 0x49, 0x1b, 0xb9, 0x64, 0x8d, 0x61, 0x9e, 0xae, 0x52, 0xf9, 0x8f, 0xf1, 0x4f,
	0x20, 0xd3, 0xda, 0x5b, 0x3f, 0x15, 0xd0, 0xfa, 0x8c, 0x3f, 0xb9, 0xa6, 0x14, 0xfe, 0x77, 0x4d,
	0x59, 0x45, 0xcb, 0xd4, 0xbd, 0x3c, 0x81, 0x41, 0x35, 0x0f, 0x99, 0xcb, 0x5d, 0xf1, 0x18, 0x91,
	0xc9, 0x12, 0x75, 0x2f, 0x8f, 0xf8, 0xbc, 0x1a, 0x4f, 0xc0, 0xc2, 0x5f, 0x99, 0x80, 0x7c, 0x76,
	0x5f, 0xa6, 0x89, 0x4c, 0x66, 0xdc, 0xd4, 0xec, 0x9e, 0x40, 0xff, 0xe2, 0xec, 0xce, 0x6c, 0xfc,
	0x2e, 0xa0, 0x77, 0x6f, 0x2f, 0xc0, 0xaf, 0x7c, 0xd3, 0xf3, 0xa8, 0xcf, 0x5f, 0x72, 0x8b, 0xf5,
	0x92, 0x5b, 0x9e, 0x1a, 0xd8, 0x5c, 0x2a, 0x13, 0x00, 0xf1, 0x63, 0x54, 0xe2, 0xff, 0x9f, 0xbe,
	0xf0, 0xfa, 0xa6, 0xe3, 0xa6, 0x5f, 0x02, 0x9b, 0x59, 0x6b, 0xe6, 0x40, 0x99, 0xe4, 0xa9, 0x7c,
	0x06, 0x50, 0xdf, 0x67, 0x7e, 0x9a, 0xbb, 0xdc, 0x0c, 0x00, 0xb1, 0x4c, 0x12, 0x18, 0x7f, 0x8c,
	0x8a, 0x3d, 0x33, 0x34, 0x21, 0x5f, 0xa5, 0xc6, 0x3b, 0x6f, 0xe8, 0x9c, 0xbc, 0x8f, 0x5c, 0x45,
	0x26, 0xa0, 0x79, 0xba, 0x08, 0x5f, 0x55, 0xfb, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xd1,
	0x72, 0xd4, 0xdb, 0x09, 0x00, 0x00,
}
