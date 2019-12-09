// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

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
//Create请求
type CreateRequest struct {
	//
	//小产品所属桌面
	Container *CreateRequest_Container `protobuf:"bytes,1,opt,name=container,proto3" json:"container" form:"container"`
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
	Icons *CreateRequest_Icons `protobuf:"bytes,5,opt,name=icons,proto3" json:"icons" form:"icons"`
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
	ClonedFrom *CreateRequest_ClonedFrom `protobuf:"bytes,14,opt,name=clonedFrom,proto3" json:"clonedFrom" form:"clonedFrom"`
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
	MenuIcon             *CreateRequest_MenuIcon `protobuf:"bytes,22,opt,name=menuIcon,proto3" json:"menuIcon" form:"menuIcon"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetContainer() *CreateRequest_Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *CreateRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateRequest) GetIcons() *CreateRequest_Icons {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *CreateRequest) GetStoryboardJson() string {
	if m != nil {
		return m.StoryboardJson
	}
	return ""
}

func (m *CreateRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreateRequest) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *CreateRequest) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *CreateRequest) GetInstallStatus() string {
	if m != nil {
		return m.InstallStatus
	}
	return ""
}

func (m *CreateRequest) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *CreateRequest) GetInternal() string {
	if m != nil {
		return m.Internal
	}
	return ""
}

func (m *CreateRequest) GetPrivate() string {
	if m != nil {
		return m.Private
	}
	return ""
}

func (m *CreateRequest) GetClonedFrom() *CreateRequest_ClonedFrom {
	if m != nil {
		return m.ClonedFrom
	}
	return nil
}

func (m *CreateRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CreateRequest) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *CreateRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateRequest) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *CreateRequest) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *CreateRequest) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *CreateRequest) GetIconBackground() string {
	if m != nil {
		return m.IconBackground
	}
	return ""
}

func (m *CreateRequest) GetMenuIcon() *CreateRequest_MenuIcon {
	if m != nil {
		return m.MenuIcon
	}
	return nil
}

type CreateRequest_Container struct {
	//
	//桌面的id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_Container) Reset()         { *m = CreateRequest_Container{} }
func (m *CreateRequest_Container) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Container) ProtoMessage()    {}
func (*CreateRequest_Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 0}
}
func (m *CreateRequest_Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_Container.Unmarshal(m, b)
}
func (m *CreateRequest_Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_Container.Marshal(b, m, deterministic)
}
func (m *CreateRequest_Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_Container.Merge(m, src)
}
func (m *CreateRequest_Container) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_Container.Size(m)
}
func (m *CreateRequest_Container) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_Container.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_Container proto.InternalMessageInfo

func (m *CreateRequest_Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateRequest_Icons struct {
	//
	//图标url
	Large                string   `protobuf:"bytes,1,opt,name=large,proto3" json:"large" form:"large"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_Icons) Reset()         { *m = CreateRequest_Icons{} }
func (m *CreateRequest_Icons) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Icons) ProtoMessage()    {}
func (*CreateRequest_Icons) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 1}
}
func (m *CreateRequest_Icons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_Icons.Unmarshal(m, b)
}
func (m *CreateRequest_Icons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_Icons.Marshal(b, m, deterministic)
}
func (m *CreateRequest_Icons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_Icons.Merge(m, src)
}
func (m *CreateRequest_Icons) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_Icons.Size(m)
}
func (m *CreateRequest_Icons) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_Icons.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_Icons proto.InternalMessageInfo

func (m *CreateRequest_Icons) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

type CreateRequest_ClonedFrom struct {
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

func (m *CreateRequest_ClonedFrom) Reset()         { *m = CreateRequest_ClonedFrom{} }
func (m *CreateRequest_ClonedFrom) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_ClonedFrom) ProtoMessage()    {}
func (*CreateRequest_ClonedFrom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 2}
}
func (m *CreateRequest_ClonedFrom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_ClonedFrom.Unmarshal(m, b)
}
func (m *CreateRequest_ClonedFrom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_ClonedFrom.Marshal(b, m, deterministic)
}
func (m *CreateRequest_ClonedFrom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_ClonedFrom.Merge(m, src)
}
func (m *CreateRequest_ClonedFrom) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_ClonedFrom.Size(m)
}
func (m *CreateRequest_ClonedFrom) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_ClonedFrom.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_ClonedFrom proto.InternalMessageInfo

func (m *CreateRequest_ClonedFrom) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateRequest_ClonedFrom) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateRequest_ClonedFrom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest_MenuIcon struct {
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

func (m *CreateRequest_MenuIcon) Reset()         { *m = CreateRequest_MenuIcon{} }
func (m *CreateRequest_MenuIcon) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_MenuIcon) ProtoMessage()    {}
func (*CreateRequest_MenuIcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 3}
}
func (m *CreateRequest_MenuIcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_MenuIcon.Unmarshal(m, b)
}
func (m *CreateRequest_MenuIcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_MenuIcon.Marshal(b, m, deterministic)
}
func (m *CreateRequest_MenuIcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_MenuIcon.Merge(m, src)
}
func (m *CreateRequest_MenuIcon) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_MenuIcon.Size(m)
}
func (m *CreateRequest_MenuIcon) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_MenuIcon.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_MenuIcon proto.InternalMessageInfo

func (m *CreateRequest_MenuIcon) GetLib() string {
	if m != nil {
		return m.Lib
	}
	return ""
}

func (m *CreateRequest_MenuIcon) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest_MenuIcon) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

func (m *CreateRequest_MenuIcon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *CreateRequest_MenuIcon) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *CreateRequest_MenuIcon) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

//
//Create返回
type CreateResponse struct {
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
	Icons *CreateResponse_Icons `protobuf:"bytes,5,opt,name=icons,proto3" json:"icons" form:"icons"`
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
	ClonedFrom *CreateResponse_ClonedFrom `protobuf:"bytes,14,opt,name=clonedFrom,proto3" json:"clonedFrom" form:"clonedFrom"`
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
	MenuIcon             *CreateResponse_MenuIcon `protobuf:"bytes,22,opt,name=menuIcon,proto3" json:"menuIcon" form:"menuIcon"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetContainer() *micro_app.MicroAppContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *CreateResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateResponse) GetIcons() *CreateResponse_Icons {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *CreateResponse) GetStoryboardJson() string {
	if m != nil {
		return m.StoryboardJson
	}
	return ""
}

func (m *CreateResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreateResponse) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *CreateResponse) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *CreateResponse) GetInstallStatus() string {
	if m != nil {
		return m.InstallStatus
	}
	return ""
}

func (m *CreateResponse) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *CreateResponse) GetInternal() string {
	if m != nil {
		return m.Internal
	}
	return ""
}

func (m *CreateResponse) GetPrivate() string {
	if m != nil {
		return m.Private
	}
	return ""
}

func (m *CreateResponse) GetClonedFrom() *CreateResponse_ClonedFrom {
	if m != nil {
		return m.ClonedFrom
	}
	return nil
}

func (m *CreateResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CreateResponse) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *CreateResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateResponse) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *CreateResponse) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *CreateResponse) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *CreateResponse) GetIconBackground() string {
	if m != nil {
		return m.IconBackground
	}
	return ""
}

func (m *CreateResponse) GetMenuIcon() *CreateResponse_MenuIcon {
	if m != nil {
		return m.MenuIcon
	}
	return nil
}

type CreateResponse_Icons struct {
	//
	//图标url
	Large                string   `protobuf:"bytes,1,opt,name=large,proto3" json:"large" form:"large"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse_Icons) Reset()         { *m = CreateResponse_Icons{} }
func (m *CreateResponse_Icons) String() string { return proto.CompactTextString(m) }
func (*CreateResponse_Icons) ProtoMessage()    {}
func (*CreateResponse_Icons) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1, 0}
}
func (m *CreateResponse_Icons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse_Icons.Unmarshal(m, b)
}
func (m *CreateResponse_Icons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse_Icons.Marshal(b, m, deterministic)
}
func (m *CreateResponse_Icons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse_Icons.Merge(m, src)
}
func (m *CreateResponse_Icons) XXX_Size() int {
	return xxx_messageInfo_CreateResponse_Icons.Size(m)
}
func (m *CreateResponse_Icons) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse_Icons.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse_Icons proto.InternalMessageInfo

func (m *CreateResponse_Icons) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

type CreateResponse_ClonedFrom struct {
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

func (m *CreateResponse_ClonedFrom) Reset()         { *m = CreateResponse_ClonedFrom{} }
func (m *CreateResponse_ClonedFrom) String() string { return proto.CompactTextString(m) }
func (*CreateResponse_ClonedFrom) ProtoMessage()    {}
func (*CreateResponse_ClonedFrom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1, 1}
}
func (m *CreateResponse_ClonedFrom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse_ClonedFrom.Unmarshal(m, b)
}
func (m *CreateResponse_ClonedFrom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse_ClonedFrom.Marshal(b, m, deterministic)
}
func (m *CreateResponse_ClonedFrom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse_ClonedFrom.Merge(m, src)
}
func (m *CreateResponse_ClonedFrom) XXX_Size() int {
	return xxx_messageInfo_CreateResponse_ClonedFrom.Size(m)
}
func (m *CreateResponse_ClonedFrom) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse_ClonedFrom.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse_ClonedFrom proto.InternalMessageInfo

func (m *CreateResponse_ClonedFrom) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateResponse_ClonedFrom) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateResponse_ClonedFrom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateResponse_MenuIcon struct {
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

func (m *CreateResponse_MenuIcon) Reset()         { *m = CreateResponse_MenuIcon{} }
func (m *CreateResponse_MenuIcon) String() string { return proto.CompactTextString(m) }
func (*CreateResponse_MenuIcon) ProtoMessage()    {}
func (*CreateResponse_MenuIcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1, 2}
}
func (m *CreateResponse_MenuIcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse_MenuIcon.Unmarshal(m, b)
}
func (m *CreateResponse_MenuIcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse_MenuIcon.Marshal(b, m, deterministic)
}
func (m *CreateResponse_MenuIcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse_MenuIcon.Merge(m, src)
}
func (m *CreateResponse_MenuIcon) XXX_Size() int {
	return xxx_messageInfo_CreateResponse_MenuIcon.Size(m)
}
func (m *CreateResponse_MenuIcon) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse_MenuIcon.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse_MenuIcon proto.InternalMessageInfo

func (m *CreateResponse_MenuIcon) GetLib() string {
	if m != nil {
		return m.Lib
	}
	return ""
}

func (m *CreateResponse_MenuIcon) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateResponse_MenuIcon) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

func (m *CreateResponse_MenuIcon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *CreateResponse_MenuIcon) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *CreateResponse_MenuIcon) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

//
//CreateApi返回
type CreateResponseWrapper struct {
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
	Data                 *CreateResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{2}
}
func (m *CreateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseWrapper.Unmarshal(m, b)
}
func (m *CreateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseWrapper.Merge(m, src)
}
func (m *CreateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateResponseWrapper.Size(m)
}
func (m *CreateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseWrapper proto.InternalMessageInfo

func (m *CreateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponseWrapper) GetData() *CreateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "installed_micro_app.CreateRequest")
	proto.RegisterType((*CreateRequest_Container)(nil), "installed_micro_app.CreateRequest.Container")
	proto.RegisterType((*CreateRequest_Icons)(nil), "installed_micro_app.CreateRequest.Icons")
	proto.RegisterType((*CreateRequest_ClonedFrom)(nil), "installed_micro_app.CreateRequest.ClonedFrom")
	proto.RegisterType((*CreateRequest_MenuIcon)(nil), "installed_micro_app.CreateRequest.MenuIcon")
	proto.RegisterType((*CreateResponse)(nil), "installed_micro_app.CreateResponse")
	proto.RegisterType((*CreateResponse_Icons)(nil), "installed_micro_app.CreateResponse.Icons")
	proto.RegisterType((*CreateResponse_ClonedFrom)(nil), "installed_micro_app.CreateResponse.ClonedFrom")
	proto.RegisterType((*CreateResponse_MenuIcon)(nil), "installed_micro_app.CreateResponse.MenuIcon")
	proto.RegisterType((*CreateResponseWrapper)(nil), "installed_micro_app.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 1292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0x92, 0x28, 0xb1, 0xd7, 0xb1, 0x6c, 0x6f, 0xe2, 0x80, 0x35, 0x5a, 0xc8, 0xdd, 0x18,
	0x2d, 0x65, 0x9b, 0x92, 0x6c, 0xf7, 0x2f, 0x46, 0x11, 0xc3, 0x0e, 0xfa, 0xe3, 0x00, 0x29, 0xd2,
	0x6d, 0x91, 0x02, 0x21, 0x25, 0x67, 0x4d, 0x6e, 0x14, 0x22, 0x22, 0x97, 0x5d, 0x52, 0x4e, 0x9c,
	0x28, 0x6f, 0xd2, 0x63, 0x4f, 0x05, 0xfa, 0x12, 0x7d, 0x8b, 0x02, 0x85, 0x0a, 0xf4, 0xd4, 0xb3,
	0x8e, 0x3d, 0x15, 0x3b, 0x4b, 0x8a, 0x94, 0xed, 0x06, 0x3a, 0xe4, 0xd2, 0xd6, 0x17, 0x6b, 0x39,
	0xf3, 0xcd, 0xb7, 0x33, 0xbb, 0xb3, 0x33, 0x03, 0xa3, 0xab, 0xae, 0xe4, 0x2c, 0xe1, 0xf5, 0x48,
	0x8a, 0x44, 0xe0, 0x6b, 0x7e, 0x18, 0x27, 0xac, 0xdb, 0xe5, 0xde, 0x41, 0xe0, 0xbb, 0x52, 0x1c,
	0xb0, 0x28, 0x5a, 0xb2, 0x3a, 0x7e, 0xf2, 0xa4, 0x77, 0x58, 0x77, 0x45, 0xd0, 0xe8, 0x88, 0x8e,
	0x68, 0x00, 0xf6, 0xb0, 0xf7, 0x18, 0xbe, 0xe0, 0x03, 0x56, 0x9a, 0x63, 0xa9, 0xd5, 0x11, 0x75,
	0xce, 0xe2, 0x63, 0x11, 0xc5, 0xf5, 0xae, 0x70, 0x59, 0xb7, 0xe1, 0x8a, 0x30, 0x91, 0xcc, 0x4d,
	0x62, 0x6d, 0x29, 0x79, 0x24, 0xac, 0x40, 0x78, 0xbc, 0x1b, 0x37, 0x52, 0x60, 0x03, 0x3e, 0x1b,
	0xa3, 0x3d, 0xf3, 0xd5, 0x81, 0x32, 0x66, 0x7e, 0xc8, 0x65, 0x4a, 0xff, 0x51, 0xc1, 0x9b, 0xe0,
	0x99, 0x9f, 0x3c, 0x15, 0xcf, 0x1a, 0x1d, 0x61, 0x81, 0xd2, 0x3a, 0x62, 0x5d, 0xdf, 0x63, 0x89,
	0x90, 0x71, 0x63, 0xb4, 0xd4, 0x76, 0xe4, 0x37, 0x8c, 0x66, 0xef, 0x40, 0xac, 0x94, 0x7f, 0xdf,
	0xe3, 0x71, 0x82, 0x1f, 0xa1, 0xe9, 0x11, 0xb9, 0x51, 0x5a, 0x2e, 0x99, 0x33, 0x9b, 0xeb, 0xf5,
	0x33, 0x0e, 0xa0, 0x3e, 0x66, 0x56, 0xbf, 0x93, 0xd9, 0xec, 0x5d, 0x1f, 0x0e, 0xaa, 0xf3, 0x8f,
	0x85, 0x0c, 0xb6, 0xc9, 0x88, 0x88, 0xd0, 0x9c, 0x14, 0xef, 0x23, 0x04, 0x7c, 0xa1, 0xcb, 0xf7,
	0x3d, 0xe3, 0xc2, 0x72, 0xc9, 0x9c, 0xde, 0xab, 0x0d, 0x07, 0xd5, 0x05, 0x6d, 0x94, 0xeb, 0xc8,
	0x1f, 0xbf, 0x57, 0xe7, 0x51, 0xa5, 0x6d, 0x37, 0xad, 0x5b, 0xcc, 0x7a, 0xd1, 0x7a, 0xb9, 0xb1,
	0xf5, 0x6a, 0x85, 0x16, 0x8c, 0xf1, 0x4d, 0x74, 0x29, 0x64, 0x01, 0x37, 0x2e, 0x02, 0xc9, 0xdc,
	0x70, 0x50, 0x9d, 0xd1, 0x24, 0x4a, 0x4a, 0x28, 0x28, 0xf1, 0x7b, 0xa8, 0xcc, 0xa2, 0x68, 0xdf,
	0x33, 0x2e, 0x01, 0x6a, 0x7e, 0x38, 0xa8, 0x5e, 0xd5, 0x28, 0x10, 0x13, 0xaa, 0xd5, 0xf8, 0x3e,
	0x2a, 0xfb, 0xae, 0x08, 0x63, 0xa3, 0x0c, 0x51, 0x9b, 0x13, 0x44, 0xbd, 0xaf, 0xf0, 0x45, 0x46,
	0x20, 0x20, 0x54, 0x13, 0xe1, 0x5d, 0x54, 0x89, 0x13, 0x21, 0x8f, 0x0f, 0x05, 0x93, 0xde, 0xdd,
	0x58, 0x84, 0xc6, 0x65, 0x70, 0xe1, 0xad, 0xe1, 0xa0, 0xba, 0xa8, 0x0d, 0xc6, 0xf5, 0x84, 0x9e,
	0x30, 0x50, 0x11, 0x26, 0xac, 0x13, 0x1b, 0x57, 0x96, 0x2f, 0x8e, 0x47, 0xa8, 0xa4, 0x84, 0x82,
	0x12, 0x3f, 0x40, 0x15, 0xb7, 0x27, 0x25, 0x0f, 0x93, 0x07, 0x5c, 0xc6, 0xbe, 0x08, 0x8d, 0x29,
	0xd8, 0xa7, 0x9e, 0xef, 0x33, 0xae, 0x57, 0x27, 0xbb, 0x80, 0xe6, 0xda, 0x8e, 0xb7, 0xe6, 0xd4,
	0xb3, 0x3f, 0x2b, 0xf4, 0x04, 0x0b, 0xbe, 0x8b, 0x10, 0x8b, 0xa2, 0x8c, 0x73, 0x1a, 0x38, 0x57,
	0xf3, 0x9b, 0xca, 0x75, 0xff, 0xc0, 0x57, 0xb0, 0xc6, 0xb7, 0xd1, 0x6c, 0x7a, 0x9e, 0xdf, 0x24,
	0x2c, 0xe9, 0xc5, 0x06, 0x02, 0x3a, 0x63, 0x38, 0xa8, 0x5e, 0x2f, 0x5c, 0x7c, 0xa6, 0x26, 0x74,
	0x1c, 0x8e, 0xfb, 0x68, 0xea, 0x89, 0x08, 0x78, 0xc4, 0x3a, 0xdc, 0x98, 0x01, 0xd3, 0x47, 0xc3,
	0x41, 0x75, 0x4e, 0x9b, 0x66, 0x1a, 0xe5, 0xc7, 0x3e, 0xfa, 0xa2, 0x6d, 0x9a, 0x4e, 0xc3, 0x6e,
	0x3b, 0x8d, 0x6d, 0x67, 0xd5, 0xd9, 0x21, 0xe4, 0xd3, 0xdb, 0x4e, 0xdf, 0x91, 0x4e, 0xd8, 0x5a,
	0xab, 0xad, 0xd5, 0xfa, 0xa6, 0xd3, 0xa8, 0xf5, 0x6d, 0x66, 0xbd, 0xd8, 0xb5, 0x1e, 0xb6, 0xb6,
	0x4d, 0xc7, 0xb1, 0xdb, 0x8e, 0x73, 0x1a, 0xb9, 0xba, 0x42, 0x47, 0x3b, 0xe2, 0x06, 0x9a, 0xf2,
	0xc3, 0x84, 0xcb, 0x90, 0x75, 0x8d, 0xab, 0xb0, 0xfb, 0xb5, 0x7c, 0xf7, 0x4c, 0x43, 0xe8, 0x08,
	0x84, 0xd7, 0xd1, 0x95, 0x48, 0xfa, 0x47, 0x2c, 0xe1, 0xc6, 0x2c, 0xe0, 0xf1, 0x70, 0x50, 0xad,
	0x68, 0x7c, 0xaa, 0x20, 0x34, 0x83, 0x60, 0x0f, 0x21, 0xb7, 0x2b, 0x42, 0xee, 0x7d, 0x2e, 0x45,
	0x60, 0x54, 0x20, 0xff, 0xac, 0x49, 0x5e, 0xdd, 0xc8, 0x68, 0x6f, 0x31, 0xbf, 0x97, 0x9c, 0x8a,
	0xd0, 0x02, 0x2f, 0xbe, 0x87, 0xca, 0xe2, 0x99, 0x7a, 0xd6, 0x73, 0xe0, 0xd1, 0xc7, 0x79, 0xda,
	0x82, 0x58, 0x1d, 0xde, 0x4d, 0xf4, 0x6e, 0x3b, 0x3d, 0x9a, 0xa6, 0x75, 0xab, 0x65, 0xd7, 0x47,
	0xeb, 0x03, 0xab, 0xf5, 0x72, 0x73, 0x7d, 0x6b, 0xe3, 0xd5, 0x0a, 0xd5, 0x2c, 0xb8, 0x86, 0x2e,
	0x4b, 0xce, 0xbc, 0x80, 0x1b, 0xf3, 0xc0, 0xb7, 0x30, 0x1c, 0x54, 0x67, 0x35, 0x9f, 0x96, 0x13,
	0x9a, 0x02, 0x14, 0x34, 0xd6, 0xb7, 0xbe, 0x70, 0x12, 0x1a, 0xa7, 0xd7, 0x9d, 0x02, 0xf0, 0xaf,
	0x25, 0x54, 0x76, 0x13, 0x3f, 0xe0, 0x06, 0x06, 0xe8, 0x2f, 0xa5, 0xdc, 0x4d, 0x90, 0x2b, 0x37,
	0x7f, 0x2e, 0xa1, 0x9f, 0x4a, 0x6d, 0xd3, 0xdc, 0xd9, 0xb6, 0x37, 0x94, 0x9b, 0xca, 0xd7, 0xd5,
	0xda, 0x0e, 0xfc, 0xbe, 0xfc, 0xe0, 0x55, 0xcd, 0x32, 0x37, 0xec, 0xa6, 0xb5, 0xd9, 0xea, 0x37,
	0x41, 0x5f, 0xb3, 0xcc, 0x2d, 0xbb, 0x69, 0x6d, 0x64, 0xdf, 0x7d, 0x7b, 0xc3, 0xda, 0xd4, 0x56,
	0x35, 0xfb, 0xdb, 0xe5, 0x96, 0xb9, 0x69, 0x37, 0xad, 0xad, 0x56, 0x1f, 0x30, 0x5a, 0xbc, 0x6d,
	0xda, 0x4d, 0xeb, 0xc3, 0xec, 0x23, 0x5f, 0x9b, 0x4e, 0x1d, 0x7e, 0xd7, 0x6a, 0x3b, 0xe6, 0xc3,
	0xbe, 0xbd, 0x66, 0xb5, 0xcc, 0x9d, 0xed, 0x33, 0xcc, 0x0b, 0xd6, 0x3b, 0x2b, 0x54, 0x47, 0x04,
	0xb1, 0x05, 0x10, 0xdb, 0xb5, 0x53, 0xb1, 0x05, 0xff, 0xda, 0xd8, 0xc0, 0x73, 0x48, 0xf8, 0xa7,
	0x9d, 0xaf, 0x54, 0x35, 0xbe, 0x7e, 0x2a, 0xe1, 0xb5, 0x42, 0x25, 0xbc, 0x5e, 0xa9, 0xca, 0xa8,
	0x4a, 0xe4, 0x1e, 0x73, 0x9f, 0x76, 0xa4, 0xe8, 0x85, 0x9e, 0xb1, 0x78, 0xb2, 0x32, 0x8e, 0xeb,
	0x09, 0x3d, 0x61, 0x80, 0x1d, 0x34, 0x15, 0xf0, 0xb0, 0xa7, 0x4a, 0xb0, 0x71, 0x03, 0x5e, 0xcc,
	0xda, 0x04, 0x2f, 0xe6, 0x5e, 0x6a, 0x52, 0x7c, 0xbf, 0x19, 0x0d, 0xa1, 0x23, 0xc6, 0xa5, 0x55,
	0x34, 0x3d, 0x6a, 0x69, 0xf8, 0x1d, 0x74, 0xc1, 0xf7, 0xa0, 0x19, 0x4e, 0xef, 0xcd, 0x0e, 0x07,
	0xd5, 0xe9, 0xd4, 0x43, 0x8f, 0xd0, 0x0b, 0xbe, 0xb7, 0xd4, 0x47, 0x65, 0x68, 0x04, 0x38, 0x46,
	0xe5, 0x2e, 0x93, 0x1d, 0x9e, 0x42, 0x5b, 0xf9, 0xed, 0x82, 0xf8, 0x0d, 0x57, 0x27, 0xbd, 0xd7,
	0xd2, 0x0f, 0x25, 0x84, 0xf2, 0x3a, 0x90, 0x77, 0xbb, 0xd2, 0xeb, 0xbb, 0xdd, 0x2e, 0xba, 0x72,
	0x94, 0x16, 0x76, 0xdd, 0x82, 0xdf, 0xcf, 0xef, 0xeb, 0xe8, 0xb5, 0x55, 0x3d, 0xb3, 0x9b, 0xa8,
	0xfb, 0x2e, 0xfd, 0x55, 0x42, 0x53, 0xd9, 0xa1, 0xe3, 0x65, 0x74, 0xb1, 0xeb, 0x1f, 0xa6, 0xae,
	0x55, 0x86, 0x83, 0x2a, 0x4a, 0x8f, 0xc7, 0x3f, 0x24, 0x54, 0xa9, 0xa0, 0xdf, 0x1d, 0x47, 0x3c,
	0xf5, 0xa9, 0xd8, 0xef, 0x8e, 0x23, 0xc5, 0xa9, 0x7e, 0x54, 0x8c, 0xc9, 0x13, 0x3e, 0xda, 0xb9,
	0x10, 0x23, 0x88, 0x09, 0xd5, 0x6a, 0x45, 0xa6, 0x92, 0x26, 0x6d, 0xfc, 0x05, 0x32, 0x1f, 0x6e,
	0x1b, 0x94, 0xaa, 0x36, 0x45, 0x92, 0x3f, 0xf6, 0x9f, 0x43, 0xdf, 0x1f, 0xab, 0x4d, 0x5a, 0x4e,
	0x68, 0x0a, 0x50, 0x5d, 0xc0, 0x65, 0x09, 0xef, 0x08, 0x79, 0x9c, 0x76, 0xf2, 0x42, 0x16, 0x65,
	0x1a, 0x42, 0x47, 0x20, 0xf2, 0x23, 0x46, 0x95, 0x2c, 0xff, 0xe2, 0x48, 0x84, 0x31, 0xc7, 0xf7,
	0x4f, 0xcf, 0x57, 0x6f, 0xd7, 0xf3, 0x6c, 0xbd, 0xa7, 0x56, 0xbb, 0x51, 0xf4, 0x1f, 0x98, 0xa7,
	0xbe, 0x1e, 0x9f, 0xa7, 0x6a, 0xaf, 0x7d, 0x9d, 0xfa, 0x74, 0xce, 0x07, 0xaa, 0xf3, 0x81, 0xea,
	0x8d, 0x0c, 0x54, 0xfc, 0x8c, 0x81, 0xaa, 0x3e, 0x49, 0x02, 0x9e, 0x4f, 0x54, 0xe7, 0x13, 0xd5,
	0xff, 0x75, 0xa2, 0x6a, 0x9d, 0x9a, 0xa8, 0xd6, 0x27, 0x79, 0x32, 0x93, 0x8e, 0x54, 0xe7, 0x63,
	0xd2, 0xf9, 0x98, 0x74, 0xf6, 0x98, 0xf4, 0x67, 0x09, 0x2d, 0x8e, 0x27, 0xd5, 0x77, 0x92, 0x45,
	0x11, 0x97, 0xca, 0x35, 0x57, 0x78, 0x3a, 0x53, 0xca, 0x45, 0xd7, 0x94, 0x94, 0x50, 0x50, 0xe2,
	0x4f, 0xd0, 0x8c, 0xfa, 0xfd, 0xec, 0x79, 0xd4, 0x65, 0x7e, 0x76, 0x4f, 0x37, 0x86, 0x83, 0x2a,
	0xce, 0xb1, 0xa9, 0x92, 0xd0, 0x22, 0x54, 0x9d, 0x10, 0x97, 0x52, 0xc8, 0xd3, 0x27, 0x04, 0x62,
	0x42, 0xb5, 0x1a, 0x7f, 0x89, 0x2e, 0x79, 0x2c, 0x61, 0x70, 0x42, 0x33, 0x9b, 0x37, 0x27, 0x78,
	0x15, 0x45, 0x5f, 0x95, 0x29, 0xa1, 0xc0, 0x70, 0x78, 0x19, 0xfe, 0xef, 0xb6, 0xf5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x08, 0xdb, 0x3d, 0x51, 0x62, 0x14, 0x00, 0x00,
}
