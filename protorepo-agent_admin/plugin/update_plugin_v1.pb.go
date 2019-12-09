// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_plugin_v1.proto

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
//UpdatePluginV1请求
type UpdatePluginV1Request struct {
	//
	//ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//部署路径
	DeployPath           []string `protobuf:"bytes,4,rep,name=deployPath,proto3" json:"deployPath" form:"deployPath"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePluginV1Request) Reset()         { *m = UpdatePluginV1Request{} }
func (m *UpdatePluginV1Request) String() string { return proto.CompactTextString(m) }
func (*UpdatePluginV1Request) ProtoMessage()    {}
func (*UpdatePluginV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d7b60a1edb22541, []int{0}
}
func (m *UpdatePluginV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePluginV1Request.Unmarshal(m, b)
}
func (m *UpdatePluginV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePluginV1Request.Marshal(b, m, deterministic)
}
func (m *UpdatePluginV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePluginV1Request.Merge(m, src)
}
func (m *UpdatePluginV1Request) XXX_Size() int {
	return xxx_messageInfo_UpdatePluginV1Request.Size(m)
}
func (m *UpdatePluginV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePluginV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePluginV1Request proto.InternalMessageInfo

func (m *UpdatePluginV1Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePluginV1Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePluginV1Request) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *UpdatePluginV1Request) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

//
//UpdatePluginV1返回
type UpdatePluginV1Response struct {
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

func (m *UpdatePluginV1Response) Reset()         { *m = UpdatePluginV1Response{} }
func (m *UpdatePluginV1Response) String() string { return proto.CompactTextString(m) }
func (*UpdatePluginV1Response) ProtoMessage()    {}
func (*UpdatePluginV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d7b60a1edb22541, []int{1}
}
func (m *UpdatePluginV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePluginV1Response.Unmarshal(m, b)
}
func (m *UpdatePluginV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePluginV1Response.Marshal(b, m, deterministic)
}
func (m *UpdatePluginV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePluginV1Response.Merge(m, src)
}
func (m *UpdatePluginV1Response) XXX_Size() int {
	return xxx_messageInfo_UpdatePluginV1Response.Size(m)
}
func (m *UpdatePluginV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePluginV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePluginV1Response proto.InternalMessageInfo

func (m *UpdatePluginV1Response) GetLastestVersion() []*agent_admin.PluginVersion {
	if m != nil {
		return m.LastestVersion
	}
	return nil
}

func (m *UpdatePluginV1Response) GetDeployedCount() int32 {
	if m != nil {
		return m.DeployedCount
	}
	return 0
}

func (m *UpdatePluginV1Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePluginV1Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePluginV1Response) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *UpdatePluginV1Response) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *UpdatePluginV1Response) GetRepoPackageId() string {
	if m != nil {
		return m.RepoPackageId
	}
	return ""
}

func (m *UpdatePluginV1Response) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *UpdatePluginV1Response) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *UpdatePluginV1Response) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *UpdatePluginV1Response) GetIsLocked() bool {
	if m != nil {
		return m.IsLocked
	}
	return false
}

//
//UpdatePluginV1Api返回
type UpdatePluginV1ResponseWrapper struct {
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
	Data                 *UpdatePluginV1Response `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdatePluginV1ResponseWrapper) Reset()         { *m = UpdatePluginV1ResponseWrapper{} }
func (m *UpdatePluginV1ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdatePluginV1ResponseWrapper) ProtoMessage()    {}
func (*UpdatePluginV1ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d7b60a1edb22541, []int{2}
}
func (m *UpdatePluginV1ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePluginV1ResponseWrapper.Unmarshal(m, b)
}
func (m *UpdatePluginV1ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePluginV1ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdatePluginV1ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePluginV1ResponseWrapper.Merge(m, src)
}
func (m *UpdatePluginV1ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdatePluginV1ResponseWrapper.Size(m)
}
func (m *UpdatePluginV1ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePluginV1ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePluginV1ResponseWrapper proto.InternalMessageInfo

func (m *UpdatePluginV1ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdatePluginV1ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdatePluginV1ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdatePluginV1ResponseWrapper) GetData() *UpdatePluginV1Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdatePluginV1Request)(nil), "plugin.UpdatePluginV1Request")
	proto.RegisterType((*UpdatePluginV1Response)(nil), "plugin.UpdatePluginV1Response")
	proto.RegisterType((*UpdatePluginV1ResponseWrapper)(nil), "plugin.UpdatePluginV1ResponseWrapper")
}

func init() { proto.RegisterFile("update_plugin_v1.proto", fileDescriptor_7d7b60a1edb22541) }

var fileDescriptor_7d7b60a1edb22541 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0xb5, 0xdd, 0x5a, 0x97, 0x6e, 0x60, 0x68, 0x15, 0x2a, 0x8d, 0x54, 0x46, 0x42,
	0xbd, 0x60, 0x89, 0x36, 0x84, 0x84, 0xb8, 0xd8, 0x45, 0x27, 0x2e, 0x90, 0xb8, 0xa8, 0x2c, 0x31,
	0x24, 0x24, 0x54, 0xb9, 0xb1, 0x97, 0x45, 0x4b, 0xe2, 0xe0, 0x38, 0x88, 0x3d, 0x19, 0x6f, 0xc0,
	0x63, 0xe4, 0x05, 0xb8, 0xcb, 0x13, 0xa0, 0x1c, 0x7b, 0x5b, 0x5a, 0x4d, 0x83, 0xab, 0xd4, 0xe7,
	0xff, 0xfe, 0xf4, 0xf8, 0xcf, 0x39, 0x68, 0x52, 0xe6, 0x9c, 0x69, 0xb1, 0xca, 0x93, 0x32, 0x8a,
	0xb3, 0xd5, 0x8f, 0x63, 0x3f, 0x57, 0x52, 0x4b, 0xbc, 0x6b, 0x0a, 0xd3, 0xa3, 0x28, 0xd6, 0x97,
	0xe5, 0xda, 0x0f, 0x65, 0x1a, 0x44, 0x32, 0x92, 0x01, 0xc8, 0xeb, 0xf2, 0x02, 0x4e, 0x70, 0x80,
	0x5f, 0xc6, 0x36, 0xfd, 0x1a, 0x49, 0x5f, 0xb0, 0xe2, 0x5a, 0xe6, 0x85, 0x9f, 0xc8, 0x90, 0x25,
	0x41, 0x28, 0x33, 0xad, 0x58, 0xa8, 0x0b, 0xe3, 0x54, 0x22, 0x97, 0x47, 0xa9, 0xe4, 0x22, 0x29,
	0x02, 0x0b, 0x06, 0x70, 0x0c, 0x58, 0x24, 0x32, 0xbd, 0x62, 0x3c, 0x8d, 0xb3, 0xe0, 0xa6, 0x1b,
	0xa1, 0x8a, 0x58, 0x66, 0xe6, 0xdd, 0xe4, 0x97, 0x83, 0xc6, 0x9f, 0xa1, 0xdb, 0x25, 0xc8, 0xe7,
	0xc7, 0x54, 0x7c, 0x2f, 0x45, 0xa1, 0xf1, 0x21, 0xda, 0x89, 0xb9, 0xeb, 0xcc, 0x9c, 0xf9, 0x60,
	0x31, 0xaa, 0x2b, 0x6f, 0x70, 0x21, 0x55, 0xfa, 0x9e, 0xc4, 0x9c, 0xd0, 0x9d, 0x98, 0xe3, 0x97,
	0xa8, 0x9b, 0xb1, 0x54, 0xb8, 0x3b, 0x00, 0x1c, 0xd4, 0x95, 0x37, 0x34, 0x40, 0x53, 0x25, 0x14,
	0xc4, 0x06, 0x4a, 0x45, 0x2a, 0xdd, 0xce, 0x36, 0xd4, 0x54, 0x09, 0x05, 0x11, 0xbf, 0x45, 0x88,
	0x8b, 0x3c, 0x91, 0xd7, 0x4b, 0xa6, 0x2f, 0xdd, 0xee, 0xac, 0x33, 0x1f, 0x2c, 0xc6, 0x75, 0xe5,
	0x3d, 0x31, 0xe8, 0x9d, 0x46, 0x68, 0x0b, 0x24, 0xbf, 0xbb, 0x68, 0xb2, 0xdd, 0x79, 0x91, 0xcb,
	0xac, 0x10, 0xf8, 0x1b, 0xda, 0x4f, 0x58, 0xa1, 0x45, 0xa1, 0xcf, 0xcd, 0x65, 0x5d, 0x67, 0xd6,
	0x99, 0x0f, 0x4f, 0xa6, 0x7e, 0x2b, 0x0f, 0xdf, 0xda, 0x0c, 0xb1, 0x78, 0x5e, 0x57, 0xde, 0xd8,
	0xfc, 0xe3, 0xa6, 0x97, 0xd0, 0xad, 0x97, 0xe1, 0x53, 0x34, 0x32, 0x7d, 0x08, 0x7e, 0x26, 0xcb,
	0x4c, 0x43, 0x06, 0xbd, 0x85, 0x5b, 0x57, 0xde, 0xb3, 0x76, 0xcf, 0x56, 0x26, 0x74, 0x13, 0xb7,
	0xc9, 0x76, 0xfe, 0x95, 0x6c, 0xf7, 0xa1, 0x64, 0x37, 0x43, 0xeb, 0xfd, 0x67, 0x68, 0xb7, 0x1f,
	0x64, 0xf7, 0xa1, 0x0f, 0x72, 0x8a, 0x46, 0xcd, 0x44, 0x2d, 0x59, 0x78, 0xc5, 0x22, 0xf1, 0x91,
	0xbb, 0x7b, 0x40, 0xb7, 0xee, 0xb7, 0x21, 0x13, 0xba, 0x89, 0xe3, 0xd7, 0x68, 0x2f, 0x54, 0x82,
	0x69, 0xa9, 0xdc, 0x3e, 0x38, 0x71, 0x5d, 0x79, 0xfb, 0xc6, 0x69, 0x05, 0x42, 0x6f, 0x10, 0xfc,
	0x0a, 0xf5, 0x42, 0x1d, 0xa7, 0xc2, 0x1d, 0x40, 0x8a, 0x8f, 0xeb, 0xca, 0x7b, 0x64, 0xd9, 0xa6,
	0x4c, 0xa8, 0x91, 0x1b, 0x2e, 0x05, 0x0e, 0x6d, 0x73, 0xa9, 0xe5, 0xe0, 0x89, 0x03, 0xd4, 0x8f,
	0x8b, 0x4f, 0x32, 0xbc, 0x12, 0xdc, 0x1d, 0xce, 0x9c, 0x79, 0x7f, 0xf1, 0xb4, 0xae, 0xbc, 0x03,
	0x9b, 0xb1, 0x55, 0x08, 0xbd, 0x85, 0xc8, 0x1f, 0x07, 0x1d, 0xde, 0x3f, 0x48, 0x5f, 0x14, 0xcb,
	0x73, 0xa1, 0x9a, 0xd4, 0x42, 0xc9, 0x05, 0x2c, 0x43, 0xaf, 0x9d, 0x5a, 0x53, 0x25, 0x14, 0x44,
	0xfc, 0x0e, 0x0d, 0x9b, 0xe7, 0x87, 0x9f, 0x79, 0xc2, 0xe2, 0xcc, 0xee, 0xc5, 0xa4, 0xae, 0x3c,
	0x7c, 0xc7, 0x5a, 0x91, 0xd0, 0x36, 0xda, 0xdc, 0x4c, 0x28, 0x25, 0x95, 0x1d, 0x89, 0xd6, 0xcd,
	0xa0, 0x4c, 0xa8, 0x91, 0xf1, 0x19, 0xea, 0x72, 0xa6, 0x19, 0x0c, 0xc6, 0xf0, 0xe4, 0x85, 0x6f,
	0x16, 0xda, 0xbf, 0xbf, 0xf7, 0x76, 0x9b, 0x8d, 0x8b, 0x50, 0x30, 0xaf, 0x77, 0x61, 0xef, 0xdf,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x90, 0x93, 0x84, 0xa4, 0x04, 0x00, 0x00,
}
