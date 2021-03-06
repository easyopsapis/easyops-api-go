// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_plugin_v1.proto

package plugin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	agent_admin "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/agent_admin"
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
//GetPluginV1请求
type GetPluginV1Request struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginV1Request) Reset()         { *m = GetPluginV1Request{} }
func (m *GetPluginV1Request) String() string { return proto.CompactTextString(m) }
func (*GetPluginV1Request) ProtoMessage()    {}
func (*GetPluginV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15dd050b79b8b1b, []int{0}
}
func (m *GetPluginV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginV1Request.Unmarshal(m, b)
}
func (m *GetPluginV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginV1Request.Marshal(b, m, deterministic)
}
func (m *GetPluginV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginV1Request.Merge(m, src)
}
func (m *GetPluginV1Request) XXX_Size() int {
	return xxx_messageInfo_GetPluginV1Request.Size(m)
}
func (m *GetPluginV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginV1Request proto.InternalMessageInfo

func (m *GetPluginV1Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetPluginV1返回
type GetPluginV1Response struct {
	//
	//插件包的最新版本信息
	LastestVersion []*agent_admin.PluginVersion `protobuf:"bytes,1,rep,name=lastestVersion,proto3" json:"lastestVersion" form:"lastestVersion"`
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

func (m *GetPluginV1Response) Reset()         { *m = GetPluginV1Response{} }
func (m *GetPluginV1Response) String() string { return proto.CompactTextString(m) }
func (*GetPluginV1Response) ProtoMessage()    {}
func (*GetPluginV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15dd050b79b8b1b, []int{1}
}
func (m *GetPluginV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginV1Response.Unmarshal(m, b)
}
func (m *GetPluginV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginV1Response.Marshal(b, m, deterministic)
}
func (m *GetPluginV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginV1Response.Merge(m, src)
}
func (m *GetPluginV1Response) XXX_Size() int {
	return xxx_messageInfo_GetPluginV1Response.Size(m)
}
func (m *GetPluginV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginV1Response proto.InternalMessageInfo

func (m *GetPluginV1Response) GetLastestVersion() []*agent_admin.PluginVersion {
	if m != nil {
		return m.LastestVersion
	}
	return nil
}

func (m *GetPluginV1Response) GetDeployedCount() int32 {
	if m != nil {
		return m.DeployedCount
	}
	return 0
}

func (m *GetPluginV1Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPluginV1Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginV1Response) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *GetPluginV1Response) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *GetPluginV1Response) GetRepoPackageId() string {
	if m != nil {
		return m.RepoPackageId
	}
	return ""
}

func (m *GetPluginV1Response) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetPluginV1Response) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *GetPluginV1Response) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *GetPluginV1Response) GetIsLocked() bool {
	if m != nil {
		return m.IsLocked
	}
	return false
}

//
//GetPluginV1Api返回
type GetPluginV1ResponseWrapper struct {
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
	Data                 *GetPluginV1Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetPluginV1ResponseWrapper) Reset()         { *m = GetPluginV1ResponseWrapper{} }
func (m *GetPluginV1ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetPluginV1ResponseWrapper) ProtoMessage()    {}
func (*GetPluginV1ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15dd050b79b8b1b, []int{2}
}
func (m *GetPluginV1ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginV1ResponseWrapper.Unmarshal(m, b)
}
func (m *GetPluginV1ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginV1ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetPluginV1ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginV1ResponseWrapper.Merge(m, src)
}
func (m *GetPluginV1ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetPluginV1ResponseWrapper.Size(m)
}
func (m *GetPluginV1ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginV1ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginV1ResponseWrapper proto.InternalMessageInfo

func (m *GetPluginV1ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetPluginV1ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetPluginV1ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetPluginV1ResponseWrapper) GetData() *GetPluginV1Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPluginV1Request)(nil), "plugin.GetPluginV1Request")
	proto.RegisterType((*GetPluginV1Response)(nil), "plugin.GetPluginV1Response")
	proto.RegisterType((*GetPluginV1ResponseWrapper)(nil), "plugin.GetPluginV1ResponseWrapper")
}

func init() { proto.RegisterFile("get_plugin_v1.proto", fileDescriptor_c15dd050b79b8b1b) }

var fileDescriptor_c15dd050b79b8b1b = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x95, 0xad, 0xed, 0x5a, 0xf7, 0xdb, 0xf6, 0xe1, 0x32, 0x14, 0x8a, 0x50, 0x2a, 0x23,
	0xa1, 0x5e, 0xb0, 0x44, 0xdb, 0x84, 0x84, 0xb8, 0x40, 0x28, 0x08, 0x21, 0x24, 0x2e, 0x2a, 0x5f,
	0x0c, 0x09, 0x09, 0x55, 0x6e, 0x7c, 0x96, 0x45, 0x4b, 0xe2, 0x60, 0xbb, 0x88, 0x3d, 0x1f, 0xef,
	0x91, 0x3b, 0x5e, 0x20, 0x4f, 0x80, 0x6c, 0x67, 0x90, 0x54, 0xe3, 0xca, 0xf5, 0xf9, 0xff, 0xfe,
	0xa7, 0xe7, 0x9c, 0x1c, 0xa3, 0x59, 0x0a, 0x7a, 0x5d, 0xe5, 0xdb, 0x34, 0x2b, 0xd7, 0xdf, 0xcf,
	0xc2, 0x4a, 0x0a, 0x2d, 0xf0, 0xc8, 0x05, 0xe6, 0xa7, 0x69, 0xa6, 0xaf, 0xb7, 0x9b, 0x30, 0x11,
	0x45, 0x94, 0x8a, 0x54, 0x44, 0x56, 0xde, 0x6c, 0xaf, 0xec, 0xcd, 0x5e, 0xec, 0x2f, 0x67, 0x9b,
	0x7f, 0x49, 0x45, 0x08, 0x4c, 0xdd, 0x8a, 0x4a, 0x85, 0xb9, 0x48, 0x58, 0x1e, 0x25, 0xa2, 0xd4,
	0x92, 0x25, 0x5a, 0x39, 0xa7, 0x84, 0x4a, 0x9c, 0x16, 0x82, 0x43, 0xae, 0xa2, 0x16, 0x8c, 0xec,
	0x35, 0x62, 0x29, 0x94, 0x7a, 0xcd, 0x78, 0x91, 0x95, 0xd1, 0x5d, 0x35, 0x20, 0x55, 0x26, 0x4a,
	0x97, 0x9b, 0x5c, 0x20, 0xfc, 0x01, 0xf4, 0xca, 0x4a, 0x97, 0x67, 0x14, 0xbe, 0x6d, 0x41, 0x69,
	0xfc, 0x14, 0xed, 0x65, 0xdc, 0xf7, 0x16, 0xde, 0x72, 0x12, 0x1f, 0x36, 0x75, 0x30, 0xb9, 0x12,
	0xb2, 0x78, 0x4d, 0x32, 0x4e, 0xe8, 0x5e, 0xc6, 0xc9, 0xcf, 0x01, 0x9a, 0xf5, 0x5c, 0xaa, 0x12,
	0xa5, 0x02, 0xfc, 0x15, 0x1d, 0xe5, 0x4c, 0x69, 0x50, 0xfa, 0xd2, 0xfd, 0x89, 0xef, 0x2d, 0xf6,
	0x97, 0xd3, 0xf3, 0x79, 0xd8, 0xa9, 0x23, 0x6c, 0x6d, 0x8e, 0x88, 0x1f, 0x37, 0x75, 0x70, 0xe2,
	0xd2, 0xf7, 0xbd, 0x84, 0xee, 0x24, 0xc3, 0x6f, 0xd0, 0x21, 0x87, 0x2a, 0x17, 0xb7, 0xc0, 0xdf,
	0x89, 0x6d, 0xa9, 0xfd, 0xbd, 0x85, 0xb7, 0x1c, 0xc6, 0x7e, 0x53, 0x07, 0x0f, 0x5d, 0x86, 0x9e,
	0x4c, 0x68, 0x1f, 0x6f, 0xbb, 0xda, 0xff, 0x47, 0x57, 0xf8, 0x19, 0x1a, 0x94, 0xac, 0x00, 0x7f,
	0x60, 0x81, 0xe3, 0xa6, 0x0e, 0xa6, 0x0e, 0x30, 0x51, 0x42, 0xad, 0x88, 0x5f, 0x22, 0xe4, 0x92,
	0xae, 0x98, 0xbe, 0xf6, 0x87, 0x8b, 0xfd, 0xe5, 0x24, 0x3e, 0x69, 0xea, 0xe0, 0x41, 0xb7, 0x00,
	0xa3, 0x11, 0xda, 0x01, 0x4d, 0xee, 0x02, 0x0a, 0xe1, 0x8f, 0x76, 0x73, 0x9b, 0x28, 0xa1, 0x56,
	0x34, 0xfd, 0x99, 0x2f, 0xb9, 0x62, 0xc9, 0x0d, 0x4b, 0xe1, 0x23, 0xf7, 0x0f, 0x2c, 0xdd, 0xe9,
	0xaf, 0x27, 0x13, 0xda, 0xc7, 0xf1, 0x0b, 0x74, 0x90, 0x48, 0x60, 0x5a, 0x48, 0x7f, 0x6c, 0x9d,
	0xb8, 0xa9, 0x83, 0x23, 0xe7, 0x6c, 0x05, 0x42, 0xef, 0x10, 0xfc, 0x1c, 0x0d, 0x13, 0x9d, 0x15,
	0xe0, 0x4f, 0xec, 0x14, 0xff, 0x6f, 0xea, 0xe0, 0xbf, 0x96, 0x35, 0x61, 0x42, 0x9d, 0x6c, 0xb8,
	0xc2, 0x72, 0x68, 0x97, 0x2b, 0x5a, 0xce, 0x9e, 0x38, 0x42, 0xe3, 0x4c, 0x7d, 0x12, 0xc9, 0x0d,
	0x70, 0x7f, 0xba, 0xf0, 0x96, 0xe3, 0x78, 0xd6, 0xd4, 0xc1, 0x71, 0x3b, 0xe3, 0x56, 0x21, 0xf4,
	0x0f, 0x44, 0x7e, 0x79, 0x68, 0x7e, 0xcf, 0x16, 0x7d, 0x96, 0xac, 0xaa, 0x40, 0x9a, 0x91, 0x25,
	0x82, 0x83, 0xdd, 0xc2, 0x61, 0x77, 0x64, 0x26, 0x4a, 0xa8, 0x15, 0xf1, 0x2b, 0x34, 0x35, 0xe7,
	0xfb, 0x1f, 0x55, 0xce, 0xb2, 0xd2, 0x2e, 0xc4, 0x24, 0x7e, 0xd4, 0xd4, 0x01, 0xfe, 0xcb, 0xb6,
	0x22, 0xa1, 0x5d, 0xd4, 0xb4, 0x05, 0x52, 0x0a, 0xd9, 0xee, 0x43, 0xa7, 0x2d, 0x1b, 0x26, 0xd4,
	0xc9, 0xf8, 0x2d, 0x1a, 0x70, 0xa6, 0x99, 0xdd, 0x8a, 0xe9, 0xf9, 0x93, 0xd0, 0xbd, 0xa2, 0xf0,
	0x9e, 0xc2, 0xbb, 0x35, 0x1a, 0x0b, 0xa1, 0xd6, 0xb9, 0x19, 0xd9, 0x97, 0x76, 0xf1, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0xc0, 0x92, 0x40, 0x13, 0x04, 0x00, 0x00,
}
