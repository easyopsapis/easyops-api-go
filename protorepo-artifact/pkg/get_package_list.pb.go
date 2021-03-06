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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

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
	// 1001 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x5f, 0x6f, 0xdb, 0x54,
	0x14, 0xaf, 0x9b, 0xf4, 0xdf, 0xcd, 0xda, 0x66, 0x77, 0xa5, 0x58, 0x45, 0xc8, 0xc5, 0xab, 0x86,
	0xdd, 0xe0, 0xfc, 0x2b, 0xa0, 0x2d, 0x20, 0xc2, 0x8a, 0x60, 0x8a, 0x34, 0xa4, 0x71, 0x37, 0x40,
	0xcc, 0x4e, 0x8a, 0xeb, 0xdc, 0xba, 0x56, 0x12, 0x5f, 0x63, 0x3b, 0xdd, 0xca, 0xd2, 0x47, 0x1e,
	0xf8, 0x2a, 0x48, 0x7c, 0x09, 0x3e, 0x04, 0x12, 0x2f, 0x46, 0xe2, 0x89, 0x67, 0xbf, 0xf1, 0x04,
	0xba, 0xc7, 0x4e, 0xec, 0x66, 0x15, 0x1a, 0x12, 0x3c, 0xb0, 0xbc, 0xe4, 0xde, 0xf3, 0xfb, 0x9d,
	0x7b, 0xfe, 0xdd, 0x73, 0x8f, 0xd1, 0xb6, 0x4d, 0xc3, 0x23, 0xcf, 0xb4, 0x06, 0xa6, 0x4d, 0x8f,
	0x86, 0x4e, 0x10, 0x56, 0x3d, 0x9f, 0x85, 0x0c, 0x17, 0xbc, 0x81, 0xbd, 0xa3, 0xd9, 0x4e, 0x78,
	0x3a, 0x3e, 0xae, 0x5a, 0x6c, 0x54, 0xb3, 0x99, 0xcd, 0x6a, 0x80, 0x1d, 0x8f, 0x4f, 0x60, 0x07,
	0x1b, 0x58, 0x25, 0x3a, 0x3b, 0xef, 0xe6, 0xe8, 0xa3, 0x27, 0x4e, 0x38, 0x60, 0x4f, 0x6a, 0x36,
	0xd3, 0x00, 0xd4, 0xce, 0xcc, 0xa1, 0xd3, 0x37, 0x43, 0xe6, 0x07, 0xb5, 0xd9, 0x32, 0xd1, 0x93,
	0xff, 0x28, 0xa3, 0xed, 0x7b, 0x34, 0x7c, 0x90, 0x78, 0x71, 0xdf, 0x09, 0x42, 0x42, 0x03, 0x8f,
	0xb9, 0x01, 0xc5, 0xdf, 0x0b, 0x68, 0x2d, 0xf5, 0xae, 0xd3, 0x17, 0x85, 0x5d, 0x41, 0x59, 0x3b,
	0x1c, 0xc4, 0x91, 0x54, 0x3e, 0x61, 0xfe, 0xa8, 0x25, 0xcf, 0x20, 0xf9, 0xb7, 0x5f, 0xa5, 0x87,
	0xe8, 0xb3, 0x9e, 0x6e, 0x6a, 0x27, 0x77, 0xb5, 0x4f, 0xea, 0xda, 0x9d, 0xee, 0xb3, 0xdb, 0x17,
	0x5a, 0x3b, 0xbf, 0x7f, 0xfb, 0x1f, 0xee, 0x1b, 0xcd, 0x8b, 0x3d, 0x92, 0x59, 0xc7, 0x15, 0x54,
	0x74, 0xcd, 0x11, 0x15, 0x17, 0xc1, 0x8b, 0x57, 0xe3, 0x48, 0x2a, 0x25, 0x5e, 0x70, 0x29, 0x77,
	0xa0, 0xe8, 0x2d, 0x3c, 0xad, 0x12, 0x20, 0xe1, 0x9b, 0xa8, 0x18, 0x9e, 0x7b, 0x54, 0x2c, 0xec,
	0x0a, 0xca, 0xd2, 0xe1, 0x66, 0x46, 0xe6, 0x52, 0x99, 0x00, 0x88, 0x77, 0x51, 0xc1, 0xea, 0xf4,
	0xc5, 0x22, 0x70, 0x36, 0xe2, 0x48, 0x42, 0x09, 0xc7, 0xea, 0xf4, 0x65, 0xc2, 0x21, 0xac, 0xa2,
	0xe5, 0x80, 0x8d, 0x7d, 0x8b, 0x8a, 0x4b, 0x60, 0xf5, 0x7a, 0x1c, 0x49, 0xeb, 0x09, 0x29, 0x91,
	0xcb, 0x24, 0x25, 0x70, 0xaa, 0x4f, 0x3d, 0xd6, 0xe9, 0x8b, 0xcb, 0xf3, 0xd4, 0x44, 0x2e, 0x93,
	0x94, 0x80, 0x27, 0x68, 0x95, 0xaf, 0x1e, 0x98, 0xe1, 0xa9, 0xb8, 0x02, 0xe4, 0xaf, 0xe3, 0x48,
	0xda, 0xcc, 0xc8, 0x1c, 0xe1, 0x11, 0x75, 0xd0, 0xbd, 0x9e, 0xa2, 0x18, 0x35, 0xbd, 0x67, 0xd4,
	0x5a, 0xc6, 0xbe, 0xd1, 0x96, 0xe5, 0xf7, 0x3f, 0x30, 0x26, 0x86, 0x6f, 0xb8, 0xdd, 0x8a, 0x5a,
	0x51, 0x27, 0x8a, 0x51, 0x53, 0x27, 0xba, 0xa9, 0x7d, 0x7b, 0x57, 0x7b, 0xdc, 0x6d, 0x29, 0x86,
	0xa1, 0xf7, 0x0c, 0xe3, 0x79, 0xe6, 0xfe, 0x1e, 0x99, 0x59, 0xc4, 0x2a, 0x2a, 0x8e, 0xe8, 0x88,
	0x89, 0xab, 0x60, 0xf9, 0x95, 0x2c, 0x35, 0x5c, 0xca, 0xad, 0x2e, 0x7a, 0x0b, 0x04, 0x28, 0xf8,
	0x21, 0x5a, 0xb1, 0x7c, 0xca, 0xaf, 0x8a, 0xb8, 0x06, 0xec, 0x3b, 0x71, 0x24, 0x6d, 0xa4, 0x49,
	0x4a, 0x00, 0xae, 0x70, 0x13, 0xbd, 0xd1, 0x4b, 0x9d, 0xe0, 0x95, 0xd3, 0xab, 0xb3, 0xf5, 0x91,
	0xd6, 0x7d, 0xd6, 0x7c, 0xeb, 0xa0, 0x71, 0xb1, 0x47, 0xa6, 0x27, 0xe1, 0x5b, 0xa8, 0xc0, 0x7c,
	0x5b, 0x44, 0x90, 0xf5, 0xad, 0x2c, 0xeb, 0xcc, 0xb7, 0xc1, 0x7a, 0x79, 0x81, 0x70, 0x02, 0xae,
	0xa1, 0x55, 0xcb, 0x0c, 0xa9, 0xcd, 0xfc, 0x73, 0xb1, 0x04, 0xd6, 0x6f, 0x64, 0x59, 0x9a, 0x22,
	0x32, 0x99, 0x91, 0x78, 0xcd, 0x1d, 0x8b, 0xb9, 0xe2, 0x35, 0x20, 0xe7, 0x6a, 0xce, 0xa5, 0x32,
	0x01, 0x10, 0xdf, 0x42, 0x4b, 0x41, 0x78, 0x3e, 0xa4, 0xe2, 0x3a, 0xb0, 0xca, 0x71, 0x24, 0x5d,
	0x4b, 0x0b, 0xca, 0xc5, 0x32, 0x49, 0x60, 0xfc, 0x8b, 0x80, 0x96, 0xac, 0xd0, 0x19, 0x51, 0x71,
	0x03, 0x88, 0x3f, 0x09, 0x19, 0x13, 0xe4, 0xdc, 0xd7, 0x1f, 0x05, 0xf4, 0x83, 0xd0, 0x53, 0x94,
	0x76, 0x4b, 0x6f, 0xf0, 0xc0, 0x79, 0xf4, 0xfb, 0x6a, 0x5b, 0x4f, 0xef, 0xb3, 0xaa, 0x29, 0x0d,
	0xbd, 0xae, 0x35, 0xbb, 0x93, 0x3a, 0xe0, 0xaa, 0xa6, 0x1c, 0xe8, 0x75, 0xad, 0x31, 0xdd, 0x4f,
	0xf4, 0x86, 0xd6, 0x4c, 0xb4, 0x54, 0xfd, 0xd1, 0x6e, 0x57, 0x69, 0xea, 0x75, 0xed, 0xa0, 0x3b,
	0x01, 0x4e, 0x22, 0x6e, 0x29, 0x7a, 0x5d, 0x7b, 0x67, 0xba, 0xc9, 0xd6, 0x8a, 0x51, 0x85, 0xff,
	0x8a, 0xda, 0x56, 0x1e, 0x4f, 0xf4, 0x8a, 0xd6, 0x55, 0xda, 0xad, 0x2b, 0xd4, 0x73, 0xda, 0xed,
	0x3d, 0x92, 0x44, 0x04, 0xb1, 0x8d, 0x20, 0xb6, 0xcd, 0xe7, 0x62, 0x1b, 0xfd, 0x6f, 0x63, 0x03,
	0xcf, 0xf1, 0x57, 0x68, 0xcd, 0x1c, 0x87, 0xa7, 0x9f, 0x07, 0xd4, 0x0f, 0xc4, 0x32, 0x84, 0xf7,
	0x5e, 0xf6, 0x60, 0xcd, 0xa0, 0x17, 0xbe, 0xb6, 0xd9, 0x69, 0xf8, 0x3b, 0x01, 0x95, 0x1c, 0x37,
	0x08, 0xcd, 0xe1, 0x10, 0x5a, 0xf7, 0x3a, 0x9c, 0x6e, 0xc5, 0x91, 0x84, 0xd3, 0x7b, 0x96, 0x81,
	0xff, 0x72, 0xf7, 0xe6, 0xed, 0xf2, 0xc6, 0xf0, 0x86, 0x66, 0xc8, 0x6d, 0x8a, 0x78, 0xbe, 0x31,
	0xa6, 0x88, 0x4c, 0x66, 0x24, 0xfc, 0x0d, 0xda, 0x1c, 0x9a, 0x41, 0xf8, 0x05, 0xf5, 0x03, 0x87,
	0xb9, 0x1d, 0xf7, 0x84, 0x89, 0x37, 0x76, 0x05, 0xa5, 0xd4, 0xac, 0x54, 0xbd, 0x81, 0x5d, 0xbd,
	0xfa, 0xed, 0xaf, 0xde, 0xbf, 0xac, 0x72, 0xb8, 0x13, 0x47, 0xd2, 0x76, 0x62, 0x64, 0xee, 0x34,
	0x99, 0xcc, 0x9f, 0x8f, 0x3f, 0x45, 0xeb, 0xe0, 0xb2, 0x6b, 0xd1, 0x8f, 0xd8, 0xd8, 0x0d, 0xc5,
	0x2d, 0x68, 0xf7, 0x37, 0xe3, 0x48, 0xda, 0xca, 0x25, 0x6b, 0x0a, 0xf3, 0x74, 0x95, 0xca, 0x7f,
	0x4e, 0x7f, 0x02, 0xb9, 0xac, 0xbd, 0xf3, 0x73, 0x01, 0x6d, 0xce, 0xf9, 0x93, 0xeb, 0x50, 0xe1,
	0xa5, 0xeb, 0xd0, 0x2a, 0x5a, 0xa5, 0xee, 0xd9, 0x11, 0x8c, 0xb0, 0x45, 0xc8, 0x5c, 0xae, 0xc4,
	0x53, 0x44, 0x26, 0x2b, 0xd4, 0x3d, 0x7b, 0xc4, 0x27, 0xd9, 0x74, 0x36, 0x16, 0x5e, 0x64, 0x36,
	0xf2, 0xa1, 0x7e, 0x96, 0x26, 0x32, 0x99, 0x7e, 0x97, 0x86, 0xfa, 0x0c, 0xfa, 0xef, 0x86, 0x7a,
	0x66, 0xe2, 0x77, 0x01, 0xbd, 0x7e, 0xf5, 0xfd, 0xfb, 0xd2, 0x37, 0x3d, 0x8f, 0xfa, 0xfc, 0x55,
	0xb7, 0x58, 0x3f, 0x29, 0xf2, 0xa5, 0x49, 0xce, 0xa5, 0x32, 0x01, 0x10, 0xdf, 0x46, 0x25, 0xfe,
	0xff, 0xf1, 0x53, 0x6f, 0x68, 0x3a, 0x6e, 0xfa, 0x89, 0xb0, 0x9d, 0x75, 0x66, 0x0e, 0x94, 0x49,
	0x9e, 0xca, 0xe7, 0x01, 0xf5, 0x7d, 0xe6, 0xa7, 0xa9, 0xcb, 0xcd, 0x03, 0x10, 0xcb, 0x24, 0x81,
	0xf1, 0x87, 0xa8, 0xd8, 0x37, 0x43, 0x13, 0xd2, 0x55, 0x6a, 0xbe, 0xf6, 0x37, 0x8d, 0x93, 0xf7,
	0x91, 0xab, 0xc8, 0x04, 0x34, 0x8f, 0x97, 0xe1, 0x6b, 0xeb, 0xe0, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x3c, 0x19, 0xc1, 0xf3, 0x09, 0x00, 0x00,
}
