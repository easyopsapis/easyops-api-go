// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_package_detail.proto

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
//GetPackageDetail请求
type GetPackageDetailRequest struct {
	//
	//包Id
	PackageId            string   `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackageDetailRequest) Reset()         { *m = GetPackageDetailRequest{} }
func (m *GetPackageDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetPackageDetailRequest) ProtoMessage()    {}
func (*GetPackageDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbf2193aabb1406, []int{0}
}
func (m *GetPackageDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageDetailRequest.Unmarshal(m, b)
}
func (m *GetPackageDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetPackageDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageDetailRequest.Merge(m, src)
}
func (m *GetPackageDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetPackageDetailRequest.Size(m)
}
func (m *GetPackageDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageDetailRequest proto.InternalMessageInfo

func (m *GetPackageDetailRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//GetPackageDetail返回
type GetPackageDetailResponse struct {
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
	LastVersionInfo *GetPackageDetailResponse_LastVersionInfo `protobuf:"bytes,19,opt,name=lastVersionInfo,proto3" json:"lastVersionInfo" form:"lastVersionInfo"`
	//
	//包实例数量
	InstanceCount        int32    `protobuf:"varint,20,opt,name=instanceCount,proto3" json:"instanceCount" form:"instanceCount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackageDetailResponse) Reset()         { *m = GetPackageDetailResponse{} }
func (m *GetPackageDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetPackageDetailResponse) ProtoMessage()    {}
func (*GetPackageDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbf2193aabb1406, []int{1}
}
func (m *GetPackageDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageDetailResponse.Unmarshal(m, b)
}
func (m *GetPackageDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetPackageDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageDetailResponse.Merge(m, src)
}
func (m *GetPackageDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetPackageDetailResponse.Size(m)
}
func (m *GetPackageDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageDetailResponse proto.InternalMessageInfo

func (m *GetPackageDetailResponse) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetPackageDetailResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPackageDetailResponse) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GetPackageDetailResponse) GetCId() int32 {
	if m != nil {
		return m.CId
	}
	return 0
}

func (m *GetPackageDetailResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *GetPackageDetailResponse) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *GetPackageDetailResponse) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *GetPackageDetailResponse) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *GetPackageDetailResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetPackageDetailResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetPackageDetailResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetPackageDetailResponse) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *GetPackageDetailResponse) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *GetPackageDetailResponse) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *GetPackageDetailResponse) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *GetPackageDetailResponse) GetAuthUsers() string {
	if m != nil {
		return m.AuthUsers
	}
	return ""
}

func (m *GetPackageDetailResponse) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *GetPackageDetailResponse) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *GetPackageDetailResponse) GetLastVersionInfo() *GetPackageDetailResponse_LastVersionInfo {
	if m != nil {
		return m.LastVersionInfo
	}
	return nil
}

func (m *GetPackageDetailResponse) GetInstanceCount() int32 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

type GetPackageDetailResponse_LastVersionInfo struct {
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

func (m *GetPackageDetailResponse_LastVersionInfo) Reset() {
	*m = GetPackageDetailResponse_LastVersionInfo{}
}
func (m *GetPackageDetailResponse_LastVersionInfo) String() string { return proto.CompactTextString(m) }
func (*GetPackageDetailResponse_LastVersionInfo) ProtoMessage()    {}
func (*GetPackageDetailResponse_LastVersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbf2193aabb1406, []int{1, 0}
}
func (m *GetPackageDetailResponse_LastVersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo.Unmarshal(m, b)
}
func (m *GetPackageDetailResponse_LastVersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo.Marshal(b, m, deterministic)
}
func (m *GetPackageDetailResponse_LastVersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo.Merge(m, src)
}
func (m *GetPackageDetailResponse_LastVersionInfo) XXX_Size() int {
	return xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo.Size(m)
}
func (m *GetPackageDetailResponse_LastVersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageDetailResponse_LastVersionInfo proto.InternalMessageInfo

func (m *GetPackageDetailResponse_LastVersionInfo) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *GetPackageDetailResponse_LastVersionInfo) GetEnvType() int32 {
	if m != nil {
		return m.EnvType
	}
	return 0
}

func (m *GetPackageDetailResponse_LastVersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPackageDetailResponse_LastVersionInfo) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//GetPackageDetailApi返回
type GetPackageDetailResponseWrapper struct {
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
	Data                 *GetPackageDetailResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetPackageDetailResponseWrapper) Reset()         { *m = GetPackageDetailResponseWrapper{} }
func (m *GetPackageDetailResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetPackageDetailResponseWrapper) ProtoMessage()    {}
func (*GetPackageDetailResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbf2193aabb1406, []int{2}
}
func (m *GetPackageDetailResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageDetailResponseWrapper.Unmarshal(m, b)
}
func (m *GetPackageDetailResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageDetailResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetPackageDetailResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageDetailResponseWrapper.Merge(m, src)
}
func (m *GetPackageDetailResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetPackageDetailResponseWrapper.Size(m)
}
func (m *GetPackageDetailResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageDetailResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageDetailResponseWrapper proto.InternalMessageInfo

func (m *GetPackageDetailResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPackageDetailResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetPackageDetailResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetPackageDetailResponseWrapper) GetData() *GetPackageDetailResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPackageDetailRequest)(nil), "pkg.GetPackageDetailRequest")
	proto.RegisterType((*GetPackageDetailResponse)(nil), "pkg.GetPackageDetailResponse")
	proto.RegisterType((*GetPackageDetailResponse_LastVersionInfo)(nil), "pkg.GetPackageDetailResponse.LastVersionInfo")
	proto.RegisterType((*GetPackageDetailResponseWrapper)(nil), "pkg.GetPackageDetailResponseWrapper")
}

func init() { proto.RegisterFile("get_package_detail.proto", fileDescriptor_1fbf2193aabb1406) }

var fileDescriptor_1fbf2193aabb1406 = []byte{
	// 1017 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x5f, 0x6f, 0xdb, 0x54,
	0x14, 0xaf, 0x9b, 0xf4, 0xdf, 0xcd, 0xda, 0x66, 0x77, 0x65, 0xbb, 0xaa, 0x84, 0x5c, 0xee, 0xaa,
	0xe1, 0xb4, 0x38, 0xff, 0x0a, 0x68, 0x0b, 0x88, 0x68, 0xe5, 0xcf, 0x14, 0x09, 0xa4, 0x71, 0x37,
	0x40, 0xcc, 0x4e, 0x8a, 0xeb, 0xdc, 0xba, 0x51, 0x13, 0x5f, 0x63, 0x3b, 0xdd, 0xca, 0xd2, 0x47,
	0x1e, 0xf8, 0x20, 0xbc, 0x20, 0xf1, 0x25, 0xf8, 0x10, 0x48, 0xbc, 0x04, 0x89, 0x67, 0x9e, 0xfc,
	0x8e, 0x40, 0xf7, 0xd8, 0x89, 0xdd, 0x94, 0xa1, 0x21, 0x81, 0x04, 0xcb, 0x4b, 0xee, 0x3d, 0xbf,
	0xdf, 0xf1, 0xf9, 0x9d, 0x73, 0xef, 0xf1, 0x31, 0x22, 0x0e, 0x0f, 0x0f, 0x3c, 0xcb, 0x3e, 0xb1,
	0x1c, 0x7e, 0xd0, 0xe5, 0xa1, 0xd5, 0xeb, 0x97, 0x3d, 0x5f, 0x84, 0x02, 0xe7, 0xbc, 0x13, 0x67,
	0x53, 0x77, 0x7a, 0xe1, 0xf1, 0xf0, 0xb0, 0x6c, 0x8b, 0x41, 0xc5, 0x11, 0x8e, 0xa8, 0x00, 0x76,
	0x38, 0x3c, 0x82, 0x1d, 0x6c, 0x60, 0x15, 0xfb, 0x6c, 0xbe, 0x99, 0xa1, 0x0f, 0x1e, 0xf7, 0xc2,
	0x13, 0xf1, 0xb8, 0xe2, 0x08, 0x1d, 0x40, 0xfd, 0xd4, 0xea, 0xf7, 0xba, 0x56, 0x28, 0xfc, 0xa0,
	0x32, 0x5d, 0xc6, 0x7e, 0xf4, 0x5b, 0x05, 0xdd, 0xb8, 0xc7, 0xc3, 0xfb, 0xb1, 0x8e, 0xf7, 0x40,
	0x06, 0xe3, 0x5f, 0x0e, 0x79, 0x10, 0xe2, 0x6f, 0x14, 0xb4, 0x92, 0x08, 0x6c, 0x75, 0x89, 0xb2,
	0xa5, 0x68, 0x2b, 0xfb, 0x27, 0xd1, 0x58, 0x2d, 0x1e, 0x09, 0x7f, 0xd0, 0xa0, 0x53, 0x88, 0xfe,
	0xf2, 0xb3, 0xfa, 0x00, 0x7d, 0xdc, 0x31, 0x2c, 0xfd, 0xe8, 0xae, 0xfe, 0x41, 0x55, 0xbf, 0xd3,
	0x7e, 0x7a, 0xfb, 0x5c, 0x6f, 0x66, 0xf7, 0xaf, 0xff, 0xcd, 0x7d, 0xad, 0x7e, 0xbe, 0xcd, 0xd2,
	0xe8, 0xf4, 0xb7, 0x22, 0x22, 0x97, 0x75, 0x06, 0x9e, 0x70, 0x03, 0xfe, 0x5f, 0x12, 0x8a, 0x77,
	0x51, 0xde, 0xb5, 0x06, 0x9c, 0xcc, 0x83, 0x8a, 0x1b, 0xd1, 0x58, 0x2d, 0xc4, 0x2a, 0xa4, 0x55,
	0x0a, 0xc8, 0x7b, 0x73, 0x4f, 0xca, 0x0c, 0x48, 0xf8, 0x26, 0xca, 0x87, 0x67, 0x1e, 0x27, 0xb9,
	0x2d, 0x45, 0x5b, 0xd8, 0x5f, 0x4f, 0xc9, 0xd2, 0x4a, 0x19, 0x80, 0x78, 0x0b, 0xe5, 0xec, 0x56,
	0x97, 0xe4, 0x81, 0xb3, 0x16, 0x8d, 0x55, 0x14, 0x73, 0xec, 0x56, 0x97, 0x32, 0x09, 0xe1, 0x12,
	0x5a, 0x0c, 0xc4, 0xd0, 0xb7, 0x39, 0x59, 0x80, 0xa8, 0x57, 0xa3, 0xb1, 0xba, 0x1a, 0x93, 0x62,
	0x3b, 0x65, 0x09, 0x41, 0x52, 0x7d, 0xee, 0x89, 0x56, 0x97, 0x2c, 0xce, 0x52, 0x63, 0x3b, 0x65,
	0x09, 0x01, 0x8f, 0xd0, 0xb2, 0x5c, 0xdd, 0xb7, 0xc2, 0x63, 0xb2, 0x04, 0xe4, 0x2f, 0xa2, 0xb1,
	0xba, 0x9e, 0x92, 0x25, 0x22, 0x33, 0x6a, 0xa1, 0x7b, 0x1d, 0x4d, 0x33, 0x2b, 0x46, 0xc7, 0xac,
	0x34, 0xcc, 0x1d, 0xb3, 0x49, 0xe9, 0xdb, 0xef, 0x98, 0x23, 0xd3, 0x37, 0xdd, 0xf6, 0x6e, 0x69,
	0xb7, 0x34, 0xd2, 0xcc, 0x4a, 0x69, 0x64, 0x58, 0xfa, 0x57, 0x77, 0xf5, 0x47, 0xed, 0x86, 0x66,
	0x9a, 0x46, 0xc7, 0x34, 0x2f, 0x33, 0x77, 0xb6, 0xd9, 0x34, 0x22, 0x2e, 0xa1, 0xfc, 0x80, 0x0f,
	0x04, 0x59, 0x86, 0xc8, 0x2f, 0xa5, 0xa5, 0x91, 0x56, 0x19, 0x75, 0xde, 0x9b, 0x63, 0x40, 0xc1,
	0x0f, 0xd0, 0x92, 0xed, 0x73, 0x79, 0xa9, 0xc9, 0x0a, 0xb0, 0xef, 0x44, 0x63, 0x75, 0x2d, 0x29,
	0x52, 0x0c, 0x48, 0x87, 0x9b, 0xe8, 0x95, 0x4e, 0x22, 0x42, 0x9e, 0x9c, 0x51, 0x9e, 0xae, 0x0f,
	0xf4, 0xf6, 0xd3, 0xfa, 0x6b, 0x7b, 0xb5, 0xf3, 0x6d, 0x36, 0x79, 0x12, 0xbe, 0x85, 0x72, 0xc2,
	0x77, 0x08, 0x82, 0xaa, 0x6f, 0xa4, 0x55, 0x17, 0xbe, 0x03, 0xd1, 0x8b, 0x73, 0x4c, 0x12, 0x70,
	0x05, 0x2d, 0xdb, 0x56, 0xc8, 0x1d, 0xe1, 0x9f, 0x91, 0x02, 0x44, 0xbf, 0x96, 0x56, 0x69, 0x82,
	0x50, 0x36, 0x25, 0xc9, 0x33, 0xef, 0xd9, 0xc2, 0x25, 0x57, 0x80, 0x9c, 0x39, 0x73, 0x69, 0xa5,
	0x0c, 0x40, 0x7c, 0x0b, 0x2d, 0x04, 0xe1, 0x59, 0x9f, 0x93, 0x55, 0x60, 0x15, 0xa3, 0xb1, 0x7a,
	0x25, 0x39, 0x50, 0x69, 0xa6, 0x2c, 0x86, 0xf1, 0x4f, 0x0a, 0x5a, 0xb0, 0xc3, 0xde, 0x80, 0x93,
	0x35, 0x20, 0xfe, 0xa0, 0xa4, 0x4c, 0xb0, 0x4b, 0xad, 0xdf, 0x2b, 0xe8, 0x3b, 0xa5, 0xa3, 0x69,
	0xcd, 0x86, 0x51, 0x93, 0x89, 0xcb, 0xec, 0x77, 0x4a, 0x4d, 0x23, 0xb9, 0xcf, 0x25, 0x5d, 0xab,
	0x19, 0x55, 0xbd, 0xde, 0x1e, 0x55, 0x01, 0x2f, 0xe9, 0xda, 0x9e, 0x51, 0xd5, 0x6b, 0x93, 0xfd,
	0xc8, 0xa8, 0xe9, 0xf5, 0xd8, 0xab, 0x64, 0x3c, 0xdc, 0x6a, 0x6b, 0x75, 0xa3, 0xaa, 0xef, 0xb5,
	0x47, 0xc0, 0x89, 0xcd, 0x0d, 0xcd, 0xa8, 0xea, 0x6f, 0x4c, 0x36, 0xe9, 0x5a, 0x33, 0xcb, 0xf0,
	0xbf, 0x5b, 0x6a, 0x6a, 0x8f, 0x46, 0xc6, 0xae, 0xde, 0xd6, 0x9a, 0x8d, 0x3f, 0x71, 0xcf, 0x78,
	0x37, 0xb7, 0x59, 0x9c, 0x11, 0xe4, 0x36, 0x80, 0xdc, 0xd6, 0x2f, 0xe5, 0x36, 0xf8, 0xdf, 0xe6,
	0x06, 0xca, 0xf1, 0xe7, 0x68, 0xc5, 0x1a, 0x86, 0xc7, 0x9f, 0x04, 0xdc, 0x0f, 0x48, 0x11, 0xd2,
	0x7b, 0x2b, 0x7d, 0x61, 0x4d, 0xa1, 0xe7, 0xbe, 0xb6, 0xe9, 0xd3, 0xf0, 0xd7, 0x0a, 0x2a, 0xf4,
	0xdc, 0x20, 0xb4, 0xfa, 0x7d, 0x68, 0xdd, 0xab, 0xf0, 0x74, 0x3b, 0x1a, 0xab, 0x38, 0xb9, 0x67,
	0x29, 0xf8, 0x0f, 0x77, 0x6f, 0x36, 0xae, 0x6c, 0x0c, 0xaf, 0x6f, 0x85, 0x32, 0x26, 0xc1, 0xb3,
	0x8d, 0x31, 0x41, 0x28, 0x9b, 0x92, 0x70, 0x80, 0xd6, 0xfb, 0x56, 0x10, 0x7e, 0xca, 0xfd, 0xa0,
	0x27, 0xdc, 0x96, 0x7b, 0x24, 0xc8, 0xb5, 0x2d, 0x45, 0x2b, 0xd4, 0xf5, 0xb2, 0x77, 0xe2, 0x94,
	0x9f, 0xf5, 0xf6, 0x2f, 0x7f, 0x78, 0xd1, 0x69, 0x7f, 0x33, 0x1a, 0xab, 0xd7, 0xe3, 0x30, 0x33,
	0xcf, 0xa3, 0x6c, 0x36, 0x02, 0xfe, 0x08, 0xad, 0x82, 0x68, 0xd7, 0xe6, 0xef, 0x8a, 0xa1, 0x1b,
	0x92, 0x0d, 0x68, 0xf8, 0x57, 0xa3, 0xb1, 0xba, 0x91, 0x29, 0xd7, 0x04, 0x96, 0x05, 0x2b, 0x14,
	0x7f, 0x9f, 0xfc, 0x14, 0x76, 0xd1, 0x7b, 0xf3, 0xc7, 0x1c, 0x5a, 0x9f, 0xd1, 0x93, 0xe9, 0x51,
	0xe5, 0x85, 0xeb, 0xd1, 0x32, 0x5a, 0xe6, 0xee, 0xe9, 0x01, 0x0c, 0xb1, 0x79, 0xa8, 0x5c, 0xe6,
	0x90, 0x27, 0x08, 0x65, 0x4b, 0xdc, 0x3d, 0x7d, 0x28, 0x67, 0xd9, 0x64, 0x3a, 0xe6, 0x9e, 0x67,
	0x3a, 0xca, 0xb1, 0x7e, 0x9a, 0x14, 0x32, 0x9e, 0x7f, 0x17, 0xc6, 0xfa, 0x14, 0xfa, 0xf7, 0xc6,
	0x7a, 0x1a, 0xe2, 0x57, 0x05, 0xa9, 0xcf, 0xba, 0x81, 0x9f, 0xf9, 0x96, 0xe7, 0x71, 0x5f, 0xbe,
	0xd9, 0x6d, 0xd1, 0x8d, 0x8f, 0xf9, 0xc2, 0x34, 0x97, 0x56, 0xca, 0x00, 0xc4, 0xb7, 0x51, 0x41,
	0xfe, 0xbf, 0xff, 0xc4, 0xeb, 0x5b, 0x3d, 0x37, 0xf9, 0x4c, 0xb8, 0x9e, 0x76, 0x67, 0x06, 0xa4,
	0x2c, 0x4b, 0x95, 0x33, 0x81, 0xfb, 0xbe, 0xf0, 0x93, 0xe2, 0x65, 0x66, 0x02, 0x98, 0x29, 0x8b,
	0x61, 0xbc, 0x8f, 0xf2, 0x5d, 0x2b, 0xb4, 0xa0, 0x60, 0x85, 0xfa, 0xcb, 0x7f, 0xd9, 0x3c, 0x59,
	0x95, 0xd2, 0x89, 0x32, 0xf0, 0x3d, 0x5c, 0x84, 0xaf, 0xc3, 0xbd, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0xad, 0x3a, 0x7c, 0xa5, 0x0a, 0x00, 0x00,
}
