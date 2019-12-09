// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package installed_micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	micro_app "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
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
//ListMicroApp请求
type ListMicroAppRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页大小
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMicroAppRequest) Reset()         { *m = ListMicroAppRequest{} }
func (m *ListMicroAppRequest) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppRequest) ProtoMessage()    {}
func (*ListMicroAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListMicroAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppRequest.Unmarshal(m, b)
}
func (m *ListMicroAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppRequest.Marshal(b, m, deterministic)
}
func (m *ListMicroAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppRequest.Merge(m, src)
}
func (m *ListMicroAppRequest) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppRequest.Size(m)
}
func (m *ListMicroAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppRequest proto.InternalMessageInfo

func (m *ListMicroAppRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListMicroAppRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListMicroApp返回
type ListMicroAppResponse struct {
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
	List                 []*ListMicroAppResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListMicroAppResponse) Reset()         { *m = ListMicroAppResponse{} }
func (m *ListMicroAppResponse) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponse) ProtoMessage()    {}
func (*ListMicroAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListMicroAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponse.Unmarshal(m, b)
}
func (m *ListMicroAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponse.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponse.Merge(m, src)
}
func (m *ListMicroAppResponse) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponse.Size(m)
}
func (m *ListMicroAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponse proto.InternalMessageInfo

func (m *ListMicroAppResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListMicroAppResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListMicroAppResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListMicroAppResponse) GetList() []*ListMicroAppResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListMicroAppResponse_List struct {
	//
	//小产品关联的桌面信息
	Container *micro_app.MicroAppContainer `protobuf:"bytes,1,opt,name=container,proto3" json:"container" form:"container"`
	//
	//cmdb唯一标识, 业务逻辑请使用appId作为唯一标识
	InstanceId string `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//小产品名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//小产品id
	AppId string `protobuf:"bytes,4,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//图标url
	Icons *ListMicroAppResponse_List_Icons `protobuf:"bytes,5,opt,name=icons,proto3" json:"icons" form:"icons"`
	//
	//小产品配置
	StoryboardJson string `protobuf:"bytes,6,opt,name=storyboardJson,proto3" json:"storyboardJson" form:"storyboardJson"`
	//
	//标签
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags" form:"tags"`
	//
	//NA程序包当前版本
	CurrentVersion string `protobuf:"bytes,8,opt,name=currentVersion,proto3" json:"currentVersion" form:"currentVersion"`
	//
	//小产品在应用商店的版本号
	AppVersion string `protobuf:"bytes,9,opt,name=appVersion,proto3" json:"appVersion" form:"appVersion"`
	//
	//安装状态， ok-成功, running-正在安装
	InstallStatus string `protobuf:"bytes,10,opt,name=installStatus,proto3" json:"installStatus" form:"installStatus"`
	//
	//小产品首页
	Homepage string `protobuf:"bytes,11,opt,name=homepage,proto3" json:"homepage" form:"homepage"`
	//
	//表示该应用是内部的，不出现在 launchpad 和 app store 中
	Internal string `protobuf:"bytes,12,opt,name=internal,proto3" json:"internal" form:"internal"`
	//
	//私有安装应用, true or false
	Private string `protobuf:"bytes,13,opt,name=private,proto3" json:"private" form:"private"`
	//
	//克隆对象
	ClonedFrom *ListMicroAppResponse_List_ClonedFrom `protobuf:"bytes,14,opt,name=clonedFrom,proto3" json:"clonedFrom" form:"clonedFrom"`
	//
	//作者
	Owner string `protobuf:"bytes,15,opt,name=owner,proto3" json:"owner" form:"owner"`
	//
	//readme
	Readme string `protobuf:"bytes,16,opt,name=readme,proto3" json:"readme" form:"readme"`
	//
	//状态
	Status string `protobuf:"bytes,17,opt,name=status,proto3" json:"status" form:"status"`
	//
	//创建时间
	Ctime string `protobuf:"bytes,18,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime string `protobuf:"bytes,19,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//关联程序包名称
	PkgName string `protobuf:"bytes,20,opt,name=pkgName,proto3" json:"pkgName" form:"pkgName"`
	//
	//图标背景
	IconBackground string `protobuf:"bytes,21,opt,name=iconBackground,proto3" json:"iconBackground" form:"iconBackground"`
	//
	//菜单中显示的图标
	MenuIcon             *ListMicroAppResponse_List_MenuIcon `protobuf:"bytes,22,opt,name=menuIcon,proto3" json:"menuIcon" form:"menuIcon"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListMicroAppResponse_List) Reset()         { *m = ListMicroAppResponse_List{} }
func (m *ListMicroAppResponse_List) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponse_List) ProtoMessage()    {}
func (*ListMicroAppResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0}
}
func (m *ListMicroAppResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponse_List.Unmarshal(m, b)
}
func (m *ListMicroAppResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponse_List.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponse_List.Merge(m, src)
}
func (m *ListMicroAppResponse_List) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponse_List.Size(m)
}
func (m *ListMicroAppResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponse_List proto.InternalMessageInfo

func (m *ListMicroAppResponse_List) GetContainer() *micro_app.MicroAppContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *ListMicroAppResponse_List) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetIcons() *ListMicroAppResponse_List_Icons {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *ListMicroAppResponse_List) GetStoryboardJson() string {
	if m != nil {
		return m.StoryboardJson
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ListMicroAppResponse_List) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetInstallStatus() string {
	if m != nil {
		return m.InstallStatus
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetInternal() string {
	if m != nil {
		return m.Internal
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetPrivate() string {
	if m != nil {
		return m.Private
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetClonedFrom() *ListMicroAppResponse_List_ClonedFrom {
	if m != nil {
		return m.ClonedFrom
	}
	return nil
}

func (m *ListMicroAppResponse_List) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetIconBackground() string {
	if m != nil {
		return m.IconBackground
	}
	return ""
}

func (m *ListMicroAppResponse_List) GetMenuIcon() *ListMicroAppResponse_List_MenuIcon {
	if m != nil {
		return m.MenuIcon
	}
	return nil
}

type ListMicroAppResponse_List_Icons struct {
	//
	//图标url
	Large                string   `protobuf:"bytes,1,opt,name=large,proto3" json:"large" form:"large"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMicroAppResponse_List_Icons) Reset()         { *m = ListMicroAppResponse_List_Icons{} }
func (m *ListMicroAppResponse_List_Icons) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponse_List_Icons) ProtoMessage()    {}
func (*ListMicroAppResponse_List_Icons) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0, 0}
}
func (m *ListMicroAppResponse_List_Icons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponse_List_Icons.Unmarshal(m, b)
}
func (m *ListMicroAppResponse_List_Icons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponse_List_Icons.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponse_List_Icons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponse_List_Icons.Merge(m, src)
}
func (m *ListMicroAppResponse_List_Icons) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponse_List_Icons.Size(m)
}
func (m *ListMicroAppResponse_List_Icons) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponse_List_Icons.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponse_List_Icons proto.InternalMessageInfo

func (m *ListMicroAppResponse_List_Icons) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

type ListMicroAppResponse_List_ClonedFrom struct {
	//
	//克隆的appId
	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//克隆的版本号
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version" form:"version"`
	//
	//克隆的app名称
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMicroAppResponse_List_ClonedFrom) Reset()         { *m = ListMicroAppResponse_List_ClonedFrom{} }
func (m *ListMicroAppResponse_List_ClonedFrom) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponse_List_ClonedFrom) ProtoMessage()    {}
func (*ListMicroAppResponse_List_ClonedFrom) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0, 1}
}
func (m *ListMicroAppResponse_List_ClonedFrom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom.Unmarshal(m, b)
}
func (m *ListMicroAppResponse_List_ClonedFrom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponse_List_ClonedFrom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom.Merge(m, src)
}
func (m *ListMicroAppResponse_List_ClonedFrom) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom.Size(m)
}
func (m *ListMicroAppResponse_List_ClonedFrom) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponse_List_ClonedFrom proto.InternalMessageInfo

func (m *ListMicroAppResponse_List_ClonedFrom) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListMicroAppResponse_List_ClonedFrom) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ListMicroAppResponse_List_ClonedFrom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListMicroAppResponse_List_MenuIcon struct {
	//
	//图标库
	Lib string `protobuf:"bytes,1,opt,name=lib,proto3" json:"lib" form:"lib"`
	//
	//Antd 图标类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	//
	//Antd 图标主题
	Theme string `protobuf:"bytes,3,opt,name=theme,proto3" json:"theme" form:"theme"`
	//
	//FA 或 EasyOps 图标
	Icon string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//FA 图标前缀
	Prefix string `protobuf:"bytes,5,opt,name=prefix,proto3" json:"prefix" form:"prefix"`
	//
	//EasyOps 图标分类
	Category             string   `protobuf:"bytes,6,opt,name=category,proto3" json:"category" form:"category"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMicroAppResponse_List_MenuIcon) Reset()         { *m = ListMicroAppResponse_List_MenuIcon{} }
func (m *ListMicroAppResponse_List_MenuIcon) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponse_List_MenuIcon) ProtoMessage()    {}
func (*ListMicroAppResponse_List_MenuIcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0, 2}
}
func (m *ListMicroAppResponse_List_MenuIcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponse_List_MenuIcon.Unmarshal(m, b)
}
func (m *ListMicroAppResponse_List_MenuIcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponse_List_MenuIcon.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponse_List_MenuIcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponse_List_MenuIcon.Merge(m, src)
}
func (m *ListMicroAppResponse_List_MenuIcon) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponse_List_MenuIcon.Size(m)
}
func (m *ListMicroAppResponse_List_MenuIcon) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponse_List_MenuIcon.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponse_List_MenuIcon proto.InternalMessageInfo

func (m *ListMicroAppResponse_List_MenuIcon) GetLib() string {
	if m != nil {
		return m.Lib
	}
	return ""
}

func (m *ListMicroAppResponse_List_MenuIcon) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ListMicroAppResponse_List_MenuIcon) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

func (m *ListMicroAppResponse_List_MenuIcon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ListMicroAppResponse_List_MenuIcon) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListMicroAppResponse_List_MenuIcon) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

//
//ListMicroAppApi返回
type ListMicroAppResponseWrapper struct {
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
	Data                 *ListMicroAppResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListMicroAppResponseWrapper) Reset()         { *m = ListMicroAppResponseWrapper{} }
func (m *ListMicroAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListMicroAppResponseWrapper) ProtoMessage()    {}
func (*ListMicroAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}
func (m *ListMicroAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMicroAppResponseWrapper.Unmarshal(m, b)
}
func (m *ListMicroAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMicroAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListMicroAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMicroAppResponseWrapper.Merge(m, src)
}
func (m *ListMicroAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListMicroAppResponseWrapper.Size(m)
}
func (m *ListMicroAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMicroAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListMicroAppResponseWrapper proto.InternalMessageInfo

func (m *ListMicroAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListMicroAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListMicroAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListMicroAppResponseWrapper) GetData() *ListMicroAppResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMicroAppRequest)(nil), "installed_micro_app.ListMicroAppRequest")
	proto.RegisterType((*ListMicroAppResponse)(nil), "installed_micro_app.ListMicroAppResponse")
	proto.RegisterType((*ListMicroAppResponse_List)(nil), "installed_micro_app.ListMicroAppResponse.List")
	proto.RegisterType((*ListMicroAppResponse_List_Icons)(nil), "installed_micro_app.ListMicroAppResponse.List.Icons")
	proto.RegisterType((*ListMicroAppResponse_List_ClonedFrom)(nil), "installed_micro_app.ListMicroAppResponse.List.ClonedFrom")
	proto.RegisterType((*ListMicroAppResponse_List_MenuIcon)(nil), "installed_micro_app.ListMicroAppResponse.List.MenuIcon")
	proto.RegisterType((*ListMicroAppResponseWrapper)(nil), "installed_micro_app.ListMicroAppResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 1292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5f, 0x6f, 0xdc, 0x44,
	0x10, 0xe7, 0x9a, 0xbb, 0x36, 0xd9, 0xb4, 0xf9, 0xb3, 0x49, 0x2b, 0x73, 0x20, 0x5d, 0xd8, 0x46,
	0xe0, 0x4b, 0x6a, 0xdf, 0xe5, 0x52, 0x28, 0x8d, 0x50, 0xa3, 0xa4, 0x02, 0x94, 0x8a, 0x56, 0x68,
	0x8b, 0x8a, 0x54, 0xfb, 0x2e, 0x6c, 0x7c, 0xdb, 0x8b, 0x55, 0xdb, 0x6b, 0xd6, 0xbe, 0xb4, 0xd7,
	0x5e, 0xbf, 0x09, 0x9f, 0x00, 0x89, 0x47, 0x9e, 0x78, 0xe3, 0x5b, 0xf0, 0x72, 0x48, 0x7c, 0x84,
	0x13, 0x4f, 0x3c, 0xa1, 0x1d, 0xdb, 0x67, 0x5f, 0x12, 0xaa, 0x56, 0x42, 0x48, 0xbc, 0x78, 0xd7,
	0xf3, 0x9b, 0xf9, 0xed, 0xcc, 0xec, 0xec, 0xec, 0x22, 0xe4, 0xb9, 0x51, 0x6c, 0x86, 0x52, 0xc4,
	0x02, 0xaf, 0xb8, 0x41, 0x14, 0x33, 0xcf, 0xe3, 0xdd, 0x43, 0xdf, 0x75, 0xa4, 0x38, 0x64, 0x61,
	0x58, 0x35, 0x7a, 0x6e, 0x7c, 0xdc, 0x3f, 0x32, 0x1d, 0xe1, 0x37, 0x7a, 0xa2, 0x27, 0x1a, 0xa0,
	0x7b, 0xd4, 0x7f, 0x02, 0x7f, 0xf0, 0x03, 0xb3, 0x84, 0xa3, 0xda, 0xee, 0x09, 0x93, 0xb3, 0x68,
	0x20, 0xc2, 0xc8, 0xf4, 0x84, 0xc3, 0xbc, 0x86, 0x23, 0x82, 0x58, 0x32, 0x27, 0x8e, 0x12, 0x4b,
	0xc9, 0x43, 0x61, 0xf8, 0xa2, 0xcb, 0xbd, 0xa8, 0x91, 0x2a, 0x36, 0xe0, 0xb7, 0x31, 0x59, 0x33,
	0x9f, 0x1d, 0x2a, 0x63, 0xe6, 0x06, 0x5c, 0xa6, 0xf4, 0x9f, 0x14, 0xbc, 0xf1, 0x9f, 0xb9, 0xf1,
	0x53, 0xf1, 0xac, 0xd1, 0x13, 0x06, 0x80, 0xc6, 0x09, 0xf3, 0xdc, 0x2e, 0x8b, 0x85, 0x8c, 0x1a,
	0x93, 0x69, 0x62, 0x47, 0x06, 0x68, 0xe5, 0x2b, 0x37, 0x8a, 0xef, 0x2b, 0xe2, 0xbd, 0x30, 0xa4,
	0xfc, 0xfb, 0x3e, 0x8f, 0x62, 0x5c, 0x47, 0xe5, 0x90, 0xf5, 0xb8, 0x56, 0x5a, 0x2b, 0xe9, 0x95,
	0xfd, 0xab, 0xe3, 0x51, 0x6d, 0xfe, 0x89, 0x90, 0xfe, 0x0e, 0x51, 0x52, 0xf2, 0xc7, 0xef, 0xb5,
	0x0b, 0x4b, 0xef, 0x50, 0x50, 0xc1, 0xb7, 0xd0, 0x9c, 0x1a, 0x0f, 0x23, 0xf7, 0x05, 0xd7, 0x2e,
	0x80, 0x7e, 0x75, 0x3c, 0xaa, 0x2d, 0xe5, 0xfa, 0x00, 0x65, 0x46, 0xb3, 0x4a, 0xf2, 0x50, 0x09,
	0x7e, 0x5e, 0x45, 0xab, 0xd3, 0x6b, 0x47, 0xa1, 0x08, 0x22, 0xfe, 0x5f, 0x2c, 0x8e, 0x3f, 0x44,
	0x95, 0x58, 0xc4, 0xcc, 0xd3, 0x66, 0xc0, 0x68, 0x69, 0x3c, 0xaa, 0x5d, 0x4e, 0x8c, 0x40, 0x4c,
	0x68, 0x02, 0xe3, 0x87, 0xa8, 0xac, 0x0a, 0x41, 0x2b, 0xaf, 0xcd, 0xe8, 0xf3, 0x2d, 0xd3, 0x3c,
	0xa7, 0x12, 0xcc, 0xf3, 0x82, 0x00, 0xe1, 0xfe, 0x62, 0xee, 0xbb, 0x62, 0x21, 0x14, 0xc8, 0xaa,
	0xbf, 0x60, 0x54, 0x56, 0x38, 0xfe, 0x1a, 0xcd, 0x4d, 0x36, 0x12, 0xc2, 0x9d, 0x6f, 0xbd, 0x6f,
	0xe6, 0xc4, 0x19, 0xe9, 0xdd, 0x4c, 0x67, 0x7f, 0x35, 0x0f, 0x6e, 0x62, 0x48, 0x68, 0x4e, 0x82,
	0x0f, 0x10, 0x02, 0x17, 0x03, 0x87, 0x1f, 0x74, 0x21, 0x23, 0x73, 0xfb, 0xf5, 0xf1, 0xa8, 0xb6,
	0x9c, 0x18, 0xe5, 0x98, 0x4a, 0xc9, 0x12, 0x5a, 0xe8, 0x58, 0x4d, 0xe3, 0x36, 0x33, 0x5e, 0xb4,
	0x5f, 0x6e, 0x6d, 0xbf, 0x5a, 0xa7, 0x05, 0x63, 0x7c, 0x1d, 0x95, 0x03, 0xe6, 0x73, 0xc8, 0xd0,
	0x5c, 0x31, 0x14, 0x25, 0x25, 0x14, 0x40, 0x95, 0x47, 0x16, 0x86, 0x07, 0x5d, 0xad, 0x0c, 0x5a,
	0x85, 0x3c, 0x82, 0x98, 0xd0, 0x04, 0xc6, 0x36, 0xaa, 0xb8, 0x8e, 0x08, 0x22, 0xad, 0x02, 0x51,
	0xde, 0x7c, 0xbb, 0x44, 0x9a, 0x07, 0xca, 0xb6, 0xc8, 0x0e, 0x64, 0x84, 0x26, 0xa4, 0x78, 0x0f,
	0x2d, 0x44, 0xb1, 0x90, 0x83, 0x23, 0xc1, 0x64, 0xf7, 0x5e, 0x24, 0x02, 0xed, 0x22, 0xb8, 0xf3,
	0xee, 0x78, 0x54, 0xbb, 0x9a, 0x18, 0x4c, 0xe3, 0x84, 0x9e, 0x32, 0x50, 0xd1, 0xc6, 0xac, 0x17,
	0x69, 0x97, 0xd6, 0x66, 0xa6, 0xa3, 0x55, 0x52, 0x42, 0x01, 0xc4, 0x8f, 0xd0, 0x82, 0xd3, 0x97,
	0x92, 0x07, 0xf1, 0x23, 0x2e, 0x23, 0x57, 0x04, 0xda, 0x2c, 0xac, 0x63, 0xe6, 0xeb, 0x4c, 0xe3,
	0x2a, 0xcb, 0xcb, 0x68, 0xb1, 0x63, 0x77, 0x37, 0x6d, 0x33, 0xfb, 0xac, 0xd3, 0x53, 0x2c, 0xf8,
	0x1e, 0x42, 0x2c, 0x0c, 0x33, 0xce, 0x39, 0xe0, 0xdc, 0xc8, 0x77, 0x2d, 0xc7, 0xfe, 0x81, 0xaf,
	0x60, 0x8d, 0xef, 0xa0, 0x2b, 0x69, 0x6e, 0x1f, 0xc6, 0x2c, 0xee, 0x47, 0x1a, 0x02, 0x3a, 0x6d,
	0x3c, 0xaa, 0xad, 0x16, 0x8a, 0x20, 0x83, 0x09, 0x9d, 0x56, 0xc7, 0x43, 0x34, 0x7b, 0x2c, 0x7c,
	0x0e, 0x27, 0x70, 0x1e, 0x4c, 0xbf, 0x1b, 0x8f, 0x6a, 0x8b, 0x89, 0x69, 0x86, 0x28, 0x3f, 0x0e,
	0xd0, 0x97, 0x1d, 0x5d, 0xb7, 0x1b, 0x56, 0xc7, 0x6e, 0xec, 0xd8, 0x1b, 0xf6, 0x2e, 0x21, 0x9f,
	0xdd, 0xb1, 0x87, 0xb6, 0xb4, 0x83, 0xf6, 0x66, 0x7d, 0xb3, 0x3e, 0xd4, 0xed, 0x46, 0x7d, 0x68,
	0x31, 0xe3, 0xc5, 0x9e, 0xf1, 0xb8, 0xbd, 0xa3, 0xdb, 0xb6, 0xd5, 0xb1, 0xed, 0xb3, 0x9a, 0x1b,
	0xeb, 0x74, 0xb2, 0x22, 0x6e, 0xa0, 0x59, 0x37, 0x88, 0xb9, 0x0c, 0x98, 0xa7, 0x5d, 0x86, 0xd5,
	0x57, 0xf2, 0xd5, 0x33, 0x84, 0xd0, 0x89, 0x12, 0xbe, 0x81, 0x2e, 0x85, 0xd2, 0x3d, 0x61, 0x31,
	0xd7, 0xae, 0x80, 0x3e, 0x1e, 0x8f, 0x6a, 0x0b, 0xe9, 0xf9, 0x4f, 0x00, 0x42, 0x33, 0x15, 0x1c,
	0x22, 0xe4, 0x78, 0x22, 0xe0, 0xdd, 0x2f, 0xa4, 0xf0, 0xb5, 0x05, 0xa8, 0xc5, 0xdb, 0x6f, 0x59,
	0x8b, 0x77, 0x27, 0x04, 0xd0, 0x9b, 0xd2, 0x3d, 0xca, 0x69, 0x09, 0x2d, 0xac, 0x81, 0xef, 0xa3,
	0x8a, 0x78, 0xa6, 0x8e, 0xf7, 0x22, 0x78, 0x77, 0x2b, 0x2f, 0x61, 0x10, 0xab, 0x44, 0x5e, 0x47,
	0x1f, 0x74, 0xd2, 0x34, 0x35, 0x8d, 0xdb, 0x6d, 0xcb, 0x9c, 0xcc, 0x0f, 0x8d, 0xf6, 0xcb, 0xd6,
	0x8d, 0xed, 0xad, 0x57, 0xeb, 0x34, 0x61, 0xc1, 0x75, 0x74, 0x51, 0x72, 0xd6, 0xf5, 0xb9, 0xb6,
	0x04, 0x7c, 0xcb, 0xe3, 0x51, 0xed, 0x4a, 0xc2, 0x97, 0xc8, 0x09, 0x4d, 0x15, 0x94, 0x6a, 0x94,
	0x54, 0xc0, 0xf2, 0x69, 0xd5, 0x28, 0xdd, 0xfa, 0x54, 0x01, 0xff, 0x56, 0x42, 0x15, 0x27, 0x76,
	0x7d, 0xae, 0x61, 0x50, 0xfd, 0xb5, 0x94, 0xbb, 0x09, 0x72, 0xe5, 0xe6, 0x4f, 0x25, 0xf4, 0x63,
	0xa9, 0xa3, 0xeb, 0xbb, 0x3b, 0xd6, 0x96, 0x72, 0x53, 0xf9, 0xba, 0x51, 0xdf, 0x85, 0xf1, 0xe5,
	0xcd, 0x57, 0x75, 0x43, 0xdf, 0xb2, 0x9a, 0x46, 0xab, 0x3d, 0x6c, 0x02, 0x5e, 0x37, 0xf4, 0x6d,
	0xab, 0x69, 0x6c, 0x65, 0xff, 0x43, 0x6b, 0xcb, 0x68, 0x25, 0x56, 0x75, 0xeb, 0x9b, 0xb5, 0xb6,
	0xde, 0xb2, 0x9a, 0xc6, 0x76, 0x7b, 0x08, 0x3a, 0x89, 0x78, 0x47, 0xb7, 0x9a, 0xc6, 0xc7, 0xd9,
	0x4f, 0x3e, 0xd7, 0x6d, 0x13, 0xc6, 0xcd, 0xfa, 0xae, 0xfe, 0x78, 0x68, 0x6d, 0x1a, 0x6d, 0x7d,
	0x77, 0xe7, 0x1c, 0xf3, 0x82, 0xf5, 0xee, 0x3a, 0x4d, 0x22, 0x82, 0xd8, 0x7c, 0x88, 0x6d, 0xe5,
	0x4c, 0x6c, 0xfe, 0xff, 0x36, 0x36, 0xf0, 0x1c, 0x8a, 0xff, 0x69, 0xef, 0x81, 0xea, 0xd2, 0xab,
	0x67, 0x8a, 0x3f, 0x01, 0x54, 0xf1, 0x27, 0x33, 0xd5, 0x25, 0x55, 0xbb, 0xdc, 0x67, 0xce, 0xd3,
	0x9e, 0x14, 0xfd, 0xa0, 0xab, 0x5d, 0x3d, 0xdd, 0x25, 0xa7, 0x71, 0x42, 0x4f, 0x19, 0xe0, 0x63,
	0x34, 0xeb, 0xf3, 0xa0, 0xaf, 0xda, 0xb1, 0x76, 0x0d, 0x4e, 0xcf, 0xad, 0xb7, 0x3c, 0x3d, 0xf7,
	0x53, 0xf3, 0xe2, 0xb9, 0xce, 0x28, 0x09, 0x9d, 0xb0, 0x57, 0x87, 0xa8, 0x02, 0x4d, 0x1f, 0x47,
	0xa8, 0xe2, 0x31, 0x99, 0x3e, 0x07, 0xe6, 0xf6, 0xdb, 0xf9, 0xee, 0x81, 0xf8, 0x5f, 0xee, 0x44,
	0xc9, 0x5a, 0xd5, 0x1f, 0x4a, 0x08, 0xe5, 0xe7, 0x3c, 0xbf, 0xe5, 0x4a, 0xaf, 0xbf, 0xe5, 0xf6,
	0xd0, 0xa5, 0x93, 0xb4, 0x89, 0x27, 0x57, 0xef, 0x47, 0xf9, 0x7e, 0x9c, 0xbc, 0xb6, 0x83, 0x67,
	0x76, 0x6f, 0x74, 0xeb, 0x56, 0xff, 0x2a, 0xa1, 0xd9, 0x2c, 0x91, 0x78, 0x0d, 0xcd, 0x78, 0xee,
	0x51, 0xea, 0xda, 0xc2, 0x78, 0x54, 0x43, 0xd9, 0x8b, 0xe3, 0x88, 0x50, 0x05, 0xc1, 0xdd, 0x36,
	0x08, 0x79, 0xea, 0x53, 0xf1, 0x6e, 0x1b, 0x84, 0x8a, 0x53, 0x0d, 0xf0, 0x22, 0x3a, 0xe6, 0x93,
	0x95, 0x8b, 0x2f, 0x22, 0x25, 0x56, 0x2f, 0x22, 0x35, 0x2a, 0x32, 0x55, 0x14, 0xe9, 0x85, 0x5f,
	0x20, 0x73, 0x61, 0x07, 0x01, 0x54, 0xbd, 0x27, 0x94, 0xfc, 0x89, 0xfb, 0x1c, 0xee, 0xfb, 0xa9,
	0xde, 0x93, 0xc8, 0x09, 0x4d, 0x15, 0x54, 0xc7, 0x77, 0x58, 0xcc, 0x7b, 0x42, 0x0e, 0xd2, 0x5b,
	0xbb, 0x50, 0x19, 0x19, 0x42, 0xe8, 0x44, 0x89, 0xfc, 0x59, 0x42, 0xef, 0x9d, 0x57, 0x5f, 0xdf,
	0x4a, 0x16, 0x86, 0x5c, 0x2a, 0x07, 0x1d, 0xd1, 0xcd, 0x9e, 0x8f, 0x05, 0x07, 0x95, 0x94, 0x50,
	0x00, 0xf1, 0xa7, 0x68, 0x5e, 0x8d, 0x9f, 0x3f, 0x0f, 0x3d, 0xe6, 0x66, 0xbb, 0x75, 0x6d, 0x3c,
	0xaa, 0xe1, 0x5c, 0x37, 0x05, 0x09, 0x2d, 0xaa, 0xaa, 0x3c, 0x71, 0x29, 0x85, 0x3c, 0x9b, 0x27,
	0x10, 0x13, 0x9a, 0xc0, 0xf8, 0x01, 0x2a, 0x77, 0x59, 0xcc, 0x20, 0x4f, 0xf3, 0xad, 0xfa, 0x1b,
	0x1f, 0x93, 0xa2, 0xc7, 0x8a, 0x80, 0x50, 0xe0, 0x39, 0xba, 0x08, 0x0f, 0xf6, 0xed, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xf7, 0xf2, 0xef, 0x3f, 0x99, 0x0c, 0x00, 0x00,
}
