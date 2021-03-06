// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_plugin_v1.proto

package plugin

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
//ListPluginV1请求
type ListPluginV1Request struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginV1Request) Reset()         { *m = ListPluginV1Request{} }
func (m *ListPluginV1Request) String() string { return proto.CompactTextString(m) }
func (*ListPluginV1Request) ProtoMessage()    {}
func (*ListPluginV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8518ae18de81f85, []int{0}
}
func (m *ListPluginV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginV1Request.Unmarshal(m, b)
}
func (m *ListPluginV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginV1Request.Marshal(b, m, deterministic)
}
func (m *ListPluginV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginV1Request.Merge(m, src)
}
func (m *ListPluginV1Request) XXX_Size() int {
	return xxx_messageInfo_ListPluginV1Request.Size(m)
}
func (m *ListPluginV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginV1Request proto.InternalMessageInfo

func (m *ListPluginV1Request) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginV1Request) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListPluginV1返回
type ListPluginV1Response struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*ListPluginV1Response_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListPluginV1Response) Reset()         { *m = ListPluginV1Response{} }
func (m *ListPluginV1Response) String() string { return proto.CompactTextString(m) }
func (*ListPluginV1Response) ProtoMessage()    {}
func (*ListPluginV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8518ae18de81f85, []int{1}
}
func (m *ListPluginV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginV1Response.Unmarshal(m, b)
}
func (m *ListPluginV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginV1Response.Marshal(b, m, deterministic)
}
func (m *ListPluginV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginV1Response.Merge(m, src)
}
func (m *ListPluginV1Response) XXX_Size() int {
	return xxx_messageInfo_ListPluginV1Response.Size(m)
}
func (m *ListPluginV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginV1Response proto.InternalMessageInfo

func (m *ListPluginV1Response) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginV1Response) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPluginV1Response) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListPluginV1Response) GetList() []*ListPluginV1Response_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListPluginV1Response_List struct {
	//
	//插件包的最新版本信息
	LastestVersion *ListPluginV1Response_List_LastestVersion `protobuf:"bytes,1,opt,name=lastestVersion,proto3" json:"lastestVersion" form:"lastestVersion"`
	//
	//部署的实例个数
	DeployedCount int32 `protobuf:"varint,2,opt,name=deployedCount,proto3" json:"deployedCount" form:"deployedCount"`
	//
	//ID
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" form:"name"`
	//
	//部署路径
	DeployPath []string `protobuf:"bytes,5,rep,name=deployPath,proto3" json:"deployPath" form:"deployPath"`
	//
	//备注
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//仓库对应包ID
	RepoPackageId string `protobuf:"bytes,7,opt,name=repoPackageId,proto3" json:"repoPackageId" form:"repoPackageId"`
	//
	//创建者
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,9,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime int32 `protobuf:"varint,10,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//是否锁定
	IsLocked             bool     `protobuf:"varint,11,opt,name=isLocked,proto3" json:"isLocked" form:"isLocked"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginV1Response_List) Reset()         { *m = ListPluginV1Response_List{} }
func (m *ListPluginV1Response_List) String() string { return proto.CompactTextString(m) }
func (*ListPluginV1Response_List) ProtoMessage()    {}
func (*ListPluginV1Response_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8518ae18de81f85, []int{1, 0}
}
func (m *ListPluginV1Response_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginV1Response_List.Unmarshal(m, b)
}
func (m *ListPluginV1Response_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginV1Response_List.Marshal(b, m, deterministic)
}
func (m *ListPluginV1Response_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginV1Response_List.Merge(m, src)
}
func (m *ListPluginV1Response_List) XXX_Size() int {
	return xxx_messageInfo_ListPluginV1Response_List.Size(m)
}
func (m *ListPluginV1Response_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginV1Response_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginV1Response_List proto.InternalMessageInfo

func (m *ListPluginV1Response_List) GetLastestVersion() *ListPluginV1Response_List_LastestVersion {
	if m != nil {
		return m.LastestVersion
	}
	return nil
}

func (m *ListPluginV1Response_List) GetDeployedCount() int32 {
	if m != nil {
		return m.DeployedCount
	}
	return 0
}

func (m *ListPluginV1Response_List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListPluginV1Response_List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListPluginV1Response_List) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *ListPluginV1Response_List) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ListPluginV1Response_List) GetRepoPackageId() string {
	if m != nil {
		return m.RepoPackageId
	}
	return ""
}

func (m *ListPluginV1Response_List) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ListPluginV1Response_List) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *ListPluginV1Response_List) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *ListPluginV1Response_List) GetIsLocked() bool {
	if m != nil {
		return m.IsLocked
	}
	return false
}

type ListPluginV1Response_List_LastestVersion struct {
	//
	//ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//仓库版本ID
	RepoVersionId string `protobuf:"bytes,2,opt,name=repoVersionId,proto3" json:"repoVersionId" form:"repoVersionId"`
	//
	//版本名称
	VersionName string `protobuf:"bytes,3,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//备注
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//所属插件ID
	PluginId string `protobuf:"bytes,5,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//创建者
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,7,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime                int32    `protobuf:"varint,8,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginV1Response_List_LastestVersion) Reset() {
	*m = ListPluginV1Response_List_LastestVersion{}
}
func (m *ListPluginV1Response_List_LastestVersion) String() string { return proto.CompactTextString(m) }
func (*ListPluginV1Response_List_LastestVersion) ProtoMessage()    {}
func (*ListPluginV1Response_List_LastestVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8518ae18de81f85, []int{1, 0, 0}
}
func (m *ListPluginV1Response_List_LastestVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginV1Response_List_LastestVersion.Unmarshal(m, b)
}
func (m *ListPluginV1Response_List_LastestVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginV1Response_List_LastestVersion.Marshal(b, m, deterministic)
}
func (m *ListPluginV1Response_List_LastestVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginV1Response_List_LastestVersion.Merge(m, src)
}
func (m *ListPluginV1Response_List_LastestVersion) XXX_Size() int {
	return xxx_messageInfo_ListPluginV1Response_List_LastestVersion.Size(m)
}
func (m *ListPluginV1Response_List_LastestVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginV1Response_List_LastestVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginV1Response_List_LastestVersion proto.InternalMessageInfo

func (m *ListPluginV1Response_List_LastestVersion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetRepoVersionId() string {
	if m != nil {
		return m.RepoVersionId
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ListPluginV1Response_List_LastestVersion) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *ListPluginV1Response_List_LastestVersion) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

//
//ListPluginV1Api返回
type ListPluginV1ResponseWrapper struct {
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
	Data                 *ListPluginV1Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListPluginV1ResponseWrapper) Reset()         { *m = ListPluginV1ResponseWrapper{} }
func (m *ListPluginV1ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListPluginV1ResponseWrapper) ProtoMessage()    {}
func (*ListPluginV1ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8518ae18de81f85, []int{2}
}
func (m *ListPluginV1ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginV1ResponseWrapper.Unmarshal(m, b)
}
func (m *ListPluginV1ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginV1ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListPluginV1ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginV1ResponseWrapper.Merge(m, src)
}
func (m *ListPluginV1ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListPluginV1ResponseWrapper.Size(m)
}
func (m *ListPluginV1ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginV1ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginV1ResponseWrapper proto.InternalMessageInfo

func (m *ListPluginV1ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListPluginV1ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListPluginV1ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListPluginV1ResponseWrapper) GetData() *ListPluginV1Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPluginV1Request)(nil), "plugin.ListPluginV1Request")
	proto.RegisterType((*ListPluginV1Response)(nil), "plugin.ListPluginV1Response")
	proto.RegisterType((*ListPluginV1Response_List)(nil), "plugin.ListPluginV1Response.List")
	proto.RegisterType((*ListPluginV1Response_List_LastestVersion)(nil), "plugin.ListPluginV1Response.List.LastestVersion")
	proto.RegisterType((*ListPluginV1ResponseWrapper)(nil), "plugin.ListPluginV1ResponseWrapper")
}

func init() { proto.RegisterFile("list_plugin_v1.proto", fileDescriptor_f8518ae18de81f85) }

var fileDescriptor_f8518ae18de81f85 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0xb5, 0x4b, 0xff, 0xb9, 0xbf, 0xfd, 0xc1, 0xdb, 0x50, 0x28, 0xa0, 0x14, 0x23, 0xa1,
	0x22, 0xb1, 0x96, 0x0d, 0x0d, 0x10, 0x0f, 0x48, 0x14, 0x81, 0x34, 0x69, 0x42, 0x93, 0x91, 0xc6,
	0xe3, 0x94, 0x35, 0x5e, 0x66, 0x2d, 0xa9, 0x43, 0xe2, 0x6e, 0xc0, 0x17, 0xe3, 0x3b, 0xf0, 0x21,
	0x82, 0xc4, 0x0b, 0xef, 0x79, 0x47, 0x42, 0xbe, 0x4e, 0x53, 0xa7, 0x1b, 0x85, 0x3d, 0xb9, 0xbe,
	0xe7, 0x1c, 0xdf, 0xeb, 0x7b, 0xae, 0x53, 0xb4, 0x11, 0xf0, 0x44, 0x1e, 0x45, 0xc1, 0xc4, 0xe7,
	0xe3, 0xa3, 0xf3, 0xed, 0x7e, 0x14, 0x0b, 0x29, 0x70, 0x5d, 0x07, 0x3a, 0x5b, 0x3e, 0x97, 0xa7,
	0x93, 0xe3, 0xfe, 0x48, 0x84, 0x03, 0x5f, 0xf8, 0x62, 0x00, 0xf0, 0xf1, 0xe4, 0x04, 0x76, 0xb0,
	0x81, 0x5f, 0x5a, 0xd6, 0x79, 0x6a, 0xd0, 0xc3, 0x0b, 0x2e, 0xcf, 0xc4, 0xc5, 0xc0, 0x17, 0x5b,
	0x00, 0x6e, 0x9d, 0xbb, 0x01, 0xf7, 0x5c, 0x29, 0xe2, 0x64, 0x50, 0xfc, 0xd4, 0x3a, 0x72, 0x81,
	0xd6, 0xf7, 0x79, 0x22, 0x0f, 0x20, 0xe9, 0xe1, 0x36, 0x65, 0x1f, 0x27, 0x2c, 0x91, 0xf8, 0x21,
	0xb2, 0x22, 0xd7, 0x67, 0x76, 0xa5, 0x5b, 0xe9, 0xd5, 0x86, 0x9b, 0x59, 0xea, 0xb4, 0x4f, 0x44,
	0x1c, 0xbe, 0x20, 0x2a, 0x4a, 0x7e, 0x7c, 0x77, 0xaa, 0x6b, 0xff, 0x51, 0xa0, 0xe0, 0x5d, 0xd4,
	0x54, 0xeb, 0x7b, 0xfe, 0x85, 0xd9, 0x55, 0xa0, 0xdf, 0xca, 0x52, 0x67, 0x75, 0x46, 0x57, 0xc8,
	0x54, 0x52, 0x50, 0xc9, 0xd7, 0x16, 0xda, 0x28, 0x67, 0x4e, 0x22, 0x31, 0x4e, 0xd8, 0x75, 0x52,
	0x3f, 0x43, 0x2d, 0xb5, 0x1e, 0x25, 0xb3, 0xdc, 0x9d, 0x2c, 0x75, 0xd6, 0x66, 0x7c, 0x80, 0x2e,
	0x25, 0xc7, 0x0f, 0x50, 0x4d, 0x0a, 0xe9, 0x06, 0xf6, 0x12, 0x88, 0xd6, 0xb2, 0xd4, 0xf9, 0x5f,
	0x8b, 0x20, 0x4c, 0xa8, 0x86, 0xf1, 0x5b, 0x64, 0x29, 0x93, 0x6c, 0xab, 0xbb, 0xd4, 0x6b, 0xef,
	0xdc, 0xeb, 0x6b, 0x6f, 0xfa, 0x57, 0xd5, 0x0d, 0xc1, 0xe1, 0xea, 0xac, 0x5c, 0x25, 0x24, 0x14,
	0xf4, 0x9d, 0x6f, 0x0d, 0x64, 0x29, 0x1c, 0x27, 0x68, 0x25, 0x70, 0x13, 0xc9, 0x12, 0x79, 0xc8,
	0xe2, 0x84, 0x8b, 0x31, 0x5c, 0xb3, 0xbd, 0xf3, 0xf8, 0xaf, 0x47, 0xf7, 0xf7, 0x4b, 0x3a, 0x68,
	0xf2, 0x66, 0x9e, 0xa9, 0x84, 0x10, 0x3a, 0x97, 0x02, 0xbf, 0x44, 0xcb, 0x1e, 0x8b, 0x02, 0xf1,
	0x99, 0x79, 0xaf, 0xc5, 0x64, 0x2c, 0xf3, 0x56, 0xd9, 0x59, 0xea, 0x6c, 0xe8, 0x13, 0x4a, 0x30,
	0xa1, 0x65, 0x3a, 0xbe, 0x8b, 0xaa, 0xdc, 0x83, 0x56, 0xb5, 0x86, 0xcb, 0x59, 0xea, 0xb4, 0xb4,
	0x88, 0x7b, 0x84, 0x56, 0xb9, 0x87, 0xef, 0x23, 0x6b, 0xec, 0x86, 0xcc, 0xb6, 0x80, 0x60, 0x74,
	0x40, 0x45, 0x09, 0x05, 0x10, 0xef, 0x22, 0xa4, 0x0f, 0x3d, 0x70, 0xe5, 0xa9, 0x5d, 0xeb, 0x2e,
	0xf5, 0x5a, 0xe0, 0xed, 0x0d, 0xb3, 0x00, 0x85, 0x11, 0x6a, 0x10, 0xd5, 0xd9, 0x21, 0x0b, 0x85,
	0x5d, 0x9f, 0x3f, 0x5b, 0x45, 0x09, 0x05, 0x50, 0xdd, 0x2f, 0x66, 0x91, 0x38, 0x70, 0x47, 0x67,
	0xae, 0xcf, 0xf6, 0x3c, 0xbb, 0x01, 0x6c, 0xe3, 0x7e, 0x25, 0x98, 0xd0, 0x32, 0x1d, 0x3f, 0x42,
	0x8d, 0x51, 0xcc, 0xd4, 0xa3, 0xb0, 0x9b, 0xa0, 0xc4, 0x59, 0xea, 0xac, 0x68, 0x65, 0x0e, 0x10,
	0x3a, 0xa5, 0xa8, 0xd9, 0x19, 0x49, 0x1e, 0x32, 0xbb, 0x35, 0x3f, 0x3b, 0x10, 0x26, 0x54, 0xc3,
	0x8a, 0x17, 0x02, 0x0f, 0xcd, 0xf3, 0xc2, 0x9c, 0x07, 0x2b, 0x1e, 0xa0, 0x26, 0x4f, 0xf6, 0xc5,
	0xe8, 0x8c, 0x79, 0x76, 0xbb, 0x5b, 0xe9, 0x35, 0x87, 0xeb, 0xb3, 0xf7, 0x33, 0x45, 0x08, 0x2d,
	0x48, 0x9d, 0x5f, 0x55, 0xb4, 0x52, 0x1e, 0x86, 0xdc, 0xa1, 0xca, 0x9f, 0x1c, 0xca, 0x1b, 0x94,
	0xb3, 0xf7, 0x3c, 0x18, 0x80, 0x4b, 0x0d, 0x2a, 0xe0, 0xbc, 0x41, 0xc5, 0x1e, 0x3f, 0x47, 0xed,
	0x73, 0xbd, 0x79, 0xa7, 0x8c, 0xd6, 0x93, 0x70, 0x33, 0x4b, 0x1d, 0xac, 0xd5, 0x06, 0x48, 0xa8,
	0x49, 0x2d, 0xfc, 0xb3, 0x16, 0xf9, 0x37, 0x40, 0x4d, 0x3d, 0xfd, 0x7b, 0x9e, 0x5d, 0x03, 0xa2,
	0xd1, 0x81, 0x29, 0x42, 0x68, 0x41, 0x32, 0x0d, 0xab, 0x5f, 0xc3, 0xb0, 0xc6, 0x3f, 0x1a, 0xd6,
	0x5c, 0x68, 0x18, 0xf9, 0x59, 0x41, 0xb7, 0xaf, 0x7a, 0xa6, 0x1f, 0x62, 0x37, 0x8a, 0x58, 0xac,
	0xee, 0x3c, 0x12, 0xde, 0xf4, 0x03, 0x66, 0xdc, 0x59, 0x45, 0x09, 0x05, 0x50, 0xb5, 0x54, 0xad,
	0x6f, 0x3e, 0x45, 0x81, 0xcb, 0xc7, 0xb9, 0x21, 0x46, 0x4b, 0x0d, 0x90, 0x50, 0x93, 0xaa, 0xca,
	0x64, 0x71, 0x2c, 0xe2, 0xdc, 0x06, 0xa3, 0x4c, 0x08, 0x13, 0xaa, 0x61, 0xfc, 0x0a, 0x59, 0x9e,
	0x2b, 0x5d, 0x68, 0x7d, 0x7b, 0xe7, 0xce, 0xa2, 0x0f, 0x8c, 0x59, 0xa4, 0xd2, 0x10, 0x0a, 0xd2,
	0xe3, 0x3a, 0xfc, 0x47, 0x3c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6b, 0x6a, 0x63, 0xaa,
	0x06, 0x00, 0x00,
}
