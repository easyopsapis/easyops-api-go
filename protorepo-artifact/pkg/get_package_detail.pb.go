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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xce, 0xc6, 0xce, 0xbf, 0x71, 0x93, 0xb8, 0xd3, 0xfc, 0xd2, 0x51, 0xa4, 0x9f, 0x36, 0x1a,
	0xa2, 0xb2, 0x4e, 0x3a, 0xb6, 0xe3, 0x94, 0xaa, 0xb5, 0x10, 0x56, 0xc3, 0x9f, 0xca, 0x52, 0x91,
	0xaa, 0x51, 0x00, 0x29, 0xbb, 0x76, 0xd8, 0xac, 0x27, 0x1b, 0x2b, 0xf6, 0xce, 0xb2, 0xbb, 0x4e,
	0x9a, 0xd6, 0xb9, 0xe4, 0x86, 0x07, 0xe1, 0x02, 0x89, 0x97, 0xe0, 0x25, 0x10, 0x37, 0x46, 0x42,
	0xe2, 0x8e, 0x2b, 0xbf, 0x00, 0x68, 0xce, 0xda, 0xde, 0x8d, 0x43, 0x2b, 0x90, 0x40, 0x02, 0x7c,
	0xe3, 0x99, 0xf3, 0x7d, 0x67, 0xcf, 0x77, 0xce, 0xcc, 0xd9, 0xb3, 0x88, 0xb8, 0x22, 0x3a, 0xf2,
	0x6d, 0xe7, 0xcc, 0x76, 0xc5, 0x51, 0x4b, 0x44, 0x76, 0xbb, 0x53, 0xf4, 0x03, 0x19, 0x49, 0x9c,
	0xf1, 0xcf, 0xdc, 0x0d, 0xe6, 0xb6, 0xa3, 0xd3, 0xde, 0x71, 0xd1, 0x91, 0xdd, 0x92, 0x2b, 0x5d,
	0x59, 0x02, 0xec, 0xb8, 0x77, 0x02, 0x3b, 0xd8, 0xc0, 0x2a, 0xf6, 0xd9, 0x78, 0x98, 0xa2, 0x77,
	0x2f, 0xda, 0xd1, 0x99, 0xbc, 0x28, 0xb9, 0x92, 0x01, 0xc8, 0xce, 0xed, 0x4e, 0xbb, 0x65, 0x47,
	0x32, 0x08, 0x4b, 0x93, 0x65, 0xec, 0x47, 0xbf, 0xd6, 0xd0, 0xdd, 0xa7, 0x22, 0x7a, 0x1e, 0xeb,
	0xf8, 0x00, 0x64, 0x70, 0xf1, 0x45, 0x4f, 0x84, 0x11, 0xfe, 0x4a, 0x43, 0x4b, 0x23, 0x81, 0xf5,
	0x16, 0xd1, 0x36, 0x35, 0x63, 0x69, 0xbf, 0x33, 0x1c, 0xe8, 0xf9, 0x13, 0x19, 0x74, 0xab, 0x74,
	0x02, 0xd1, 0x9f, 0x7e, 0xd4, 0x0f, 0x10, 0x6f, 0x9a, 0x35, 0x9b, 0xbd, 0x7c, 0xc2, 0x0e, 0xcb,
	0xec, 0x71, 0xe3, 0xd5, 0xa3, 0x2b, 0x76, 0x6d, 0xff, 0xe0, 0x4f, 0xee, 0x77, 0x2b, 0x57, 0x5b,
	0x3c, 0x09, 0x4f, 0x7f, 0xce, 0x23, 0x72, 0x53, 0x68, 0xe8, 0x4b, 0x2f, 0x14, 0xff, 0x28, 0xa5,
	0x78, 0x07, 0x65, 0x3d, 0xbb, 0x2b, 0xc8, 0x2c, 0xc8, 0xb8, 0x3b, 0x1c, 0xe8, 0xb9, 0x58, 0x86,
	0xb2, 0x2a, 0x05, 0x59, 0x7f, 0xe6, 0x45, 0x91, 0x03, 0x09, 0xbf, 0x85, 0xb2, 0xd1, 0xa5, 0x2f,
	0x48, 0x66, 0x53, 0x33, 0xe6, 0xf6, 0x57, 0x13, 0xb2, 0xb2, 0x52, 0x0e, 0x20, 0xde, 0x44, 0x19,
	0xa7, 0xde, 0x22, 0x59, 0xe0, 0xac, 0x0c, 0x07, 0x3a, 0x8a, 0x39, 0x4e, 0xbd, 0x45, 0xb9, 0x82,
	0x70, 0x01, 0xcd, 0x87, 0xb2, 0x17, 0x38, 0x82, 0xcc, 0x41, 0xd4, 0xdb, 0xc3, 0x81, 0xbe, 0x1c,
	0x93, 0x62, 0x3b, 0xe5, 0x23, 0x82, 0xa2, 0x06, 0xc2, 0x97, 0xf5, 0x16, 0x99, 0x9f, 0xa6, 0xc6,
	0x76, 0xca, 0x47, 0x04, 0xdc, 0x47, 0x8b, 0x6a, 0xf5, 0xdc, 0x8e, 0x4e, 0xc9, 0x02, 0x90, 0x3f,
	0x1f, 0x0e, 0xf4, 0xd5, 0x84, 0xac, 0x10, 0x95, 0x51, 0x1d, 0x3d, 0x6d, 0x1a, 0x86, 0x55, 0x32,
	0x9b, 0x56, 0xa9, 0x6a, 0x6d, 0x5b, 0x35, 0x4a, 0xdf, 0x7d, 0xcf, 0xea, 0x5b, 0x81, 0xe5, 0x35,
	0x76, 0x0a, 0x3b, 0x85, 0xbe, 0x61, 0x95, 0x0a, 0x7d, 0x33, 0x2e, 0x5a, 0xa3, 0x6a, 0x58, 0x96,
	0xd9, 0xb4, 0xac, 0x9b, 0xcc, 0xed, 0x2d, 0x3e, 0x89, 0x88, 0x0b, 0x28, 0xdb, 0x15, 0x5d, 0x49,
	0x16, 0x21, 0xf2, 0xff, 0x92, 0xd2, 0x28, 0xab, 0x8a, 0x3a, 0xeb, 0xcf, 0x70, 0xa0, 0xe0, 0x8f,
	0xd0, 0x82, 0x13, 0x08, 0x75, 0xad, 0xc9, 0x12, 0xb0, 0xef, 0x0f, 0x07, 0xfa, 0xca, 0xa8, 0x48,
	0x31, 0xa0, 0x1c, 0xd6, 0xd1, 0x5a, 0xd3, 0x7c, 0xc2, 0x0e, 0x6d, 0xf6, 0xf2, 0x88, 0x35, 0xac,
	0x8b, 0x57, 0x7b, 0xf7, 0x1f, 0x3e, 0xb8, 0xda, 0xe2, 0x63, 0x67, 0x7c, 0x0f, 0x65, 0x64, 0xe0,
	0x12, 0x04, 0x85, 0x5e, 0x4b, 0x0a, 0x2d, 0x03, 0x17, 0x02, 0xe6, 0x67, 0xb8, 0x22, 0xe0, 0x12,
	0x5a, 0x74, 0xec, 0x48, 0xb8, 0x32, 0xb8, 0x24, 0x39, 0x08, 0x78, 0x27, 0x29, 0xcc, 0x18, 0xa1,
	0x7c, 0x42, 0x52, 0xc7, 0xdc, 0x76, 0xa4, 0x47, 0x6e, 0x01, 0x39, 0x75, 0xcc, 0xca, 0x4a, 0x39,
	0x80, 0xf8, 0x1e, 0x9a, 0x0b, 0xa3, 0xcb, 0x8e, 0x20, 0xcb, 0xc0, 0xca, 0x0f, 0x07, 0xfa, 0xad,
	0xd1, 0x19, 0x2a, 0x33, 0xe5, 0x31, 0x8c, 0x7f, 0xd0, 0xd0, 0x9c, 0x13, 0xb5, 0xbb, 0x82, 0xac,
	0x00, 0xf1, 0x3b, 0x2d, 0x61, 0x82, 0x5d, 0x69, 0xfd, 0x56, 0x43, 0xdf, 0x68, 0x4d, 0xc3, 0xa8,
	0x55, 0xcd, 0x5d, 0xf6, 0xb8, 0x61, 0xaa, 0xab, 0xba, 0x5d, 0xa8, 0x99, 0xa3, 0x2b, 0x5c, 0x60,
	0xc6, 0xae, 0x59, 0x66, 0x95, 0x46, 0xbf, 0x0c, 0x78, 0x81, 0x19, 0x7b, 0x66, 0x99, 0xed, 0x8e,
	0xf7, 0x7d, 0x73, 0x97, 0x55, 0x62, 0xaf, 0x82, 0x79, 0xb0, 0xd9, 0x30, 0x2a, 0x66, 0x99, 0xed,
	0x35, 0xfa, 0xc0, 0x89, 0xcd, 0x55, 0xc3, 0x2c, 0xb3, 0x77, 0xc6, 0x9b, 0x64, 0x6d, 0x58, 0x45,
	0xf8, 0xdf, 0x29, 0xd4, 0x8c, 0xc3, 0xbe, 0xb9, 0xc3, 0x1a, 0x46, 0xad, 0xfa, 0x3b, 0xee, 0x29,
	0xef, 0xda, 0x16, 0x8f, 0x33, 0x82, 0xdc, 0xba, 0x90, 0xdb, 0xea, 0x8d, 0xdc, 0xba, 0xff, 0xda,
	0xdc, 0x40, 0x39, 0x7e, 0x86, 0x96, 0xec, 0x5e, 0x74, 0xfa, 0x49, 0x28, 0x82, 0x90, 0xe4, 0x21,
	0xbd, 0x62, 0xf2, 0x92, 0x9a, 0x40, 0x6f, 0xba, 0xa9, 0xc9, 0x03, 0xf0, 0x97, 0x1a, 0xca, 0xb5,
	0xbd, 0x30, 0xb2, 0x3b, 0x1d, 0x68, 0xd0, 0xdb, 0xf0, 0x40, 0x67, 0x38, 0xd0, 0xf1, 0xe8, 0x6a,
	0x25, 0xe0, 0x5f, 0xdc, 0xa3, 0xe9, 0xb8, 0xaa, 0x17, 0xfc, 0x8e, 0x1d, 0xa9, 0x98, 0x04, 0x4f,
	0xf7, 0xc2, 0x18, 0xa1, 0x7c, 0x42, 0xc2, 0x21, 0x5a, 0xed, 0xd8, 0x61, 0xf4, 0xa9, 0x08, 0xc2,
	0xb6, 0xf4, 0xea, 0xde, 0x89, 0x24, 0x77, 0x36, 0x35, 0x23, 0x57, 0x61, 0x45, 0xff, 0xcc, 0x2d,
	0xbe, 0xee, 0x25, 0x5f, 0x7c, 0x76, 0xdd, 0x69, 0x7f, 0x63, 0x38, 0xd0, 0xd7, 0xe3, 0x30, 0x53,
	0xcf, 0xa3, 0x7c, 0x3a, 0x02, 0xfe, 0x18, 0x2d, 0x83, 0x68, 0xcf, 0x11, 0xef, 0xcb, 0x9e, 0x17,
	0x91, 0x35, 0xe8, 0xf1, 0xb7, 0x87, 0x03, 0x7d, 0x2d, 0x55, 0xae, 0x31, 0xac, 0x0a, 0x96, 0xcb,
	0xff, 0x3a, 0xfe, 0x69, 0xfc, 0xba, 0xf7, 0xc6, 0xf7, 0x19, 0xb4, 0x3a, 0xa5, 0x27, 0xd5, 0x96,
	0xda, 0x7f, 0xae, 0x2d, 0x8b, 0x68, 0x51, 0x78, 0xe7, 0x47, 0x30, 0xaa, 0x66, 0xa1, 0x72, 0xa9,
	0x43, 0x1e, 0x23, 0x94, 0x2f, 0x08, 0xef, 0xfc, 0x40, 0x4d, 0xac, 0xf1, 0x0c, 0xcc, 0xfc, 0x91,
	0x19, 0xa8, 0xa6, 0xf7, 0xf9, 0xa8, 0x90, 0xf1, 0x94, 0xbb, 0x36, 0xbd, 0x27, 0xd0, 0xdf, 0x38,
	0xbd, 0x93, 0x18, 0xbf, 0x68, 0x48, 0x7f, 0xdd, 0x15, 0xfc, 0x2c, 0xb0, 0x7d, 0x5f, 0x04, 0xea,
	0x6d, 0xee, 0xc8, 0x56, 0x7c, 0xce, 0xd7, 0x86, 0xb6, 0xb2, 0x52, 0x0e, 0x20, 0x7e, 0x84, 0x72,
	0xea, 0xff, 0xc3, 0x17, 0x7e, 0xc7, 0x6e, 0x7b, 0xa3, 0xaf, 0x81, 0xf5, 0xa4, 0x3d, 0x53, 0x20,
	0xe5, 0x69, 0xaa, 0x9a, 0x03, 0x22, 0x08, 0x64, 0x30, 0xaa, 0x5e, 0x6a, 0x0e, 0x80, 0x99, 0xf2,
	0x18, 0xc6, 0xfb, 0x28, 0xdb, 0xb2, 0x23, 0x1b, 0x2a, 0x96, 0xab, 0xfc, 0xff, 0x8d, 0xdd, 0x93,
	0x56, 0xa9, 0x9c, 0x28, 0x07, 0xdf, 0xe3, 0x79, 0xf8, 0x0c, 0xdc, 0xfb, 0x2d, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0x7e, 0xd4, 0x79, 0x8e, 0x0a, 0x00, 0x00,
}
