// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_version_with_sign.proto

package version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	artifact "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/artifact"
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
//CreateVersionWithSign请求
type CreateVersionWithSignRequest struct {
	//
	//程序包Id
	PackageId string `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId string `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//包名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//版本备注
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//包文件的md5检验值
	Sign string `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign" form:"sign"`
	//
	//版本的进程及启停、安装、升级脚本, package.conf.yaml
	Conf string `protobuf:"bytes,6,opt,name=conf,proto3" json:"conf" form:"conf"`
	//
	//版本类型 1 开发, 3 测试, 7 预发布, 15 生产
	EnvType int32 `protobuf:"varint,7,opt,name=env_type,json=envType,proto3" json:"env_type" form:"env_type"`
	//
	//基础Docker镜像Id
	BaseImageId string `protobuf:"bytes,8,opt,name=baseImageId,proto3" json:"baseImageId" form:"baseImageId"`
	//
	//工作区基础版本Id
	WorkspaceBaseId string `protobuf:"bytes,9,opt,name=workspaceBaseId,proto3" json:"workspaceBaseId" form:"workspaceBaseId"`
	//
	//包存储源信息,如"{"host":"deployrepo.easyops-only.com","ensName":"logic.deploy.repo.archive"}"
	Source               string   `protobuf:"bytes,10,opt,name=source,proto3" json:"source" form:"source"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVersionWithSignRequest) Reset()         { *m = CreateVersionWithSignRequest{} }
func (m *CreateVersionWithSignRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVersionWithSignRequest) ProtoMessage()    {}
func (*CreateVersionWithSignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a5186d574c97c7, []int{0}
}
func (m *CreateVersionWithSignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVersionWithSignRequest.Unmarshal(m, b)
}
func (m *CreateVersionWithSignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVersionWithSignRequest.Marshal(b, m, deterministic)
}
func (m *CreateVersionWithSignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVersionWithSignRequest.Merge(m, src)
}
func (m *CreateVersionWithSignRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVersionWithSignRequest.Size(m)
}
func (m *CreateVersionWithSignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVersionWithSignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVersionWithSignRequest proto.InternalMessageInfo

func (m *CreateVersionWithSignRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetEnvType() int32 {
	if m != nil {
		return m.EnvType
	}
	return 0
}

func (m *CreateVersionWithSignRequest) GetBaseImageId() string {
	if m != nil {
		return m.BaseImageId
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetWorkspaceBaseId() string {
	if m != nil {
		return m.WorkspaceBaseId
	}
	return ""
}

func (m *CreateVersionWithSignRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

//
//CreateVersionWithSignApi返回
type CreateVersionWithSignResponseWrapper struct {
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
	Data                 *artifact.Version `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateVersionWithSignResponseWrapper) Reset()         { *m = CreateVersionWithSignResponseWrapper{} }
func (m *CreateVersionWithSignResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateVersionWithSignResponseWrapper) ProtoMessage()    {}
func (*CreateVersionWithSignResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a5186d574c97c7, []int{1}
}
func (m *CreateVersionWithSignResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVersionWithSignResponseWrapper.Unmarshal(m, b)
}
func (m *CreateVersionWithSignResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVersionWithSignResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateVersionWithSignResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVersionWithSignResponseWrapper.Merge(m, src)
}
func (m *CreateVersionWithSignResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateVersionWithSignResponseWrapper.Size(m)
}
func (m *CreateVersionWithSignResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVersionWithSignResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVersionWithSignResponseWrapper proto.InternalMessageInfo

func (m *CreateVersionWithSignResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateVersionWithSignResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateVersionWithSignResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateVersionWithSignResponseWrapper) GetData() *artifact.Version {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateVersionWithSignRequest)(nil), "version.CreateVersionWithSignRequest")
	proto.RegisterType((*CreateVersionWithSignResponseWrapper)(nil), "version.CreateVersionWithSignResponseWrapper")
}

func init() { proto.RegisterFile("create_version_with_sign.proto", fileDescriptor_e7a5186d574c97c7) }

var fileDescriptor_e7a5186d574c97c7 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xd7, 0x91, 0x76, 0x9b, 0x07, 0xda, 0x66, 0xc4, 0x88, 0x26, 0x44, 0x26, 0x33, 0xa1,
	0x4d, 0x28, 0x09, 0x0c, 0x34, 0x0d, 0x6e, 0x10, 0x43, 0x20, 0xed, 0x0e, 0x3c, 0xc4, 0x2e, 0x10,
	0x54, 0x6e, 0xe2, 0x66, 0x51, 0x9b, 0xd8, 0xb3, 0xdd, 0x76, 0x13, 0xda, 0x03, 0x20, 0x1e, 0x81,
	0x77, 0x0b, 0x12, 0x12, 0x2f, 0x90, 0x27, 0x40, 0x3e, 0xc9, 0xba, 0x50, 0x89, 0x0b, 0x2e, 0x76,
	0xd5, 0x9c, 0xf3, 0x7f, 0xc7, 0xfe, 0xed, 0x73, 0x5c, 0x74, 0x3f, 0x52, 0x9c, 0x19, 0xde, 0x1d,
	0x73, 0xa5, 0x53, 0x91, 0x77, 0x27, 0xa9, 0x39, 0xe9, 0xea, 0x34, 0xc9, 0x03, 0xa9, 0x84, 0x11,
	0x78, 0xa1, 0x16, 0x36, 0xfc, 0x24, 0x35, 0x27, 0xa3, 0x5e, 0x10, 0x89, 0x2c, 0x4c, 0x44, 0x22,
	0x42, 0xd0, 0x7b, 0xa3, 0x3e, 0x44, 0x10, 0xc0, 0x57, 0x55, 0xb7, 0xf1, 0x2e, 0x11, 0x01, 0x67,
	0xfa, 0x5c, 0x48, 0x1d, 0x0c, 0x45, 0xc4, 0x86, 0x61, 0x24, 0x72, 0xa3, 0x58, 0x64, 0x74, 0x55,
	0xa9, 0xb8, 0x14, 0x7e, 0x26, 0x62, 0x3e, 0xd4, 0x61, 0x0d, 0x86, 0x10, 0x86, 0x4c, 0x99, 0xb4,
	0xcf, 0x22, 0x13, 0xd6, 0x5b, 0xd7, 0x2b, 0xee, 0x35, 0x0c, 0x64, 0x93, 0xd4, 0x0c, 0xc4, 0x24,
	0x4c, 0x84, 0x0f, 0xa2, 0x3f, 0x66, 0xc3, 0x34, 0x66, 0x46, 0x28, 0x1d, 0x4e, 0x3f, 0xab, 0x3a,
	0xf2, 0xbd, 0x83, 0xee, 0xbd, 0x86, 0x43, 0x7e, 0xac, 0xd6, 0x3b, 0x4e, 0xcd, 0xc9, 0x51, 0x9a,
	0xe4, 0x94, 0x9f, 0x8e, 0xb8, 0x36, 0xf8, 0x5b, 0x0b, 0x2d, 0x49, 0x16, 0x0d, 0x58, 0xc2, 0x0f,
	0x63, 0xb7, 0xb5, 0xd9, 0xda, 0x5e, 0x3a, 0x18, 0x94, 0x85, 0xb7, 0xda, 0x17, 0x2a, 0x7b, 0x41,
	0xa6, 0x12, 0xf9, 0xf5, 0xd3, 0x3b, 0x42, 0xef, 0xbf, 0x7c, 0x62, 0x7e, 0xff, 0x95, 0xff, 0xf6,
	0xb1, 0xff, 0xfc, 0xf3, 0xd7, 0xfd, 0x0b, 0xff, 0x65, 0x33, 0x7e, 0xf6, 0x9f, 0xf1, 0x93, 0xdd,
	0x8b, 0x2d, 0x7a, 0xb5, 0x3b, 0x78, 0xa9, 0x8f, 0x7d, 0x18, 0xbb, 0xf3, 0xb3, 0x5e, 0xa6, 0xd2,
	0xf5, 0x79, 0x99, 0x6e, 0x81, 0x1f, 0x21, 0x27, 0x67, 0x19, 0x77, 0x6f, 0x80, 0x8b, 0xbb, 0x65,
	0xe1, 0x2d, 0x57, 0x2e, 0x6c, 0xd6, 0x1a, 0x70, 0xe4, 0xdc, 0x59, 0x40, 0x01, 0xc2, 0x3b, 0xc8,
	0xc9, 0x78, 0x26, 0x5c, 0x07, 0xe0, 0x3b, 0x57, 0xb0, 0xcd, 0x5a, 0x78, 0x5e, 0xce, 0x51, 0x40,
	0xf0, 0x03, 0xe4, 0xd8, 0x01, 0x73, 0xdb, 0x80, 0xae, 0x5c, 0xa1, 0x36, 0x4b, 0x28, 0x88, 0x16,
	0x8a, 0x44, 0xde, 0x77, 0x3b, 0xb3, 0x90, 0xcd, 0x12, 0x0a, 0x22, 0x0e, 0xd0, 0x22, 0xcf, 0xc7,
	0x5d, 0x73, 0x2e, 0xb9, 0xbb, 0xb0, 0xd9, 0xda, 0x6e, 0x1f, 0xdc, 0x2e, 0x0b, 0x6f, 0xa5, 0x02,
	0x2f, 0x15, 0x42, 0x17, 0x78, 0x3e, 0xfe, 0x70, 0x2e, 0x39, 0xde, 0x47, 0xcb, 0x3d, 0xa6, 0xf9,
	0x61, 0x56, 0xb5, 0x7a, 0x11, 0xd6, 0x5e, 0x2f, 0x0b, 0x0f, 0x57, 0x25, 0x0d, 0x91, 0xd0, 0x26,
	0x8a, 0x7f, 0xb4, 0xd0, 0xca, 0x44, 0xa8, 0x81, 0x96, 0x2c, 0xe2, 0x07, 0x56, 0x88, 0xdd, 0x25,
	0x28, 0x3f, 0x2d, 0x0b, 0x6f, 0xbd, 0x2a, 0x9f, 0x01, 0xae, 0xad, 0x47, 0xb3, 0x4e, 0xf0, 0x0e,
	0xea, 0x68, 0x31, 0x52, 0x11, 0x77, 0x11, 0x78, 0x5a, 0x2b, 0x0b, 0xef, 0x56, 0x7d, 0xa7, 0x90,
	0x27, 0xb4, 0x06, 0xc8, 0xef, 0x16, 0xda, 0xfa, 0xc7, 0x6b, 0xd0, 0x52, 0xe4, 0x9a, 0x1f, 0x2b,
	0x26, 0x25, 0x57, 0x55, 0x03, 0x62, 0x0e, 0xef, 0xa1, 0xfd, 0x77, 0x03, 0x62, 0x0e, 0x0d, 0x88,
	0xe1, 0x42, 0xed, 0xef, 0x9b, 0x33, 0x39, 0x64, 0x69, 0x5e, 0xcf, 0x6b, 0xe3, 0x42, 0x1b, 0x22,
	0xa1, 0x4d, 0x14, 0x3f, 0x44, 0x6d, 0xae, 0x94, 0x50, 0xf5, 0x74, 0xad, 0x96, 0x85, 0x77, 0xb3,
	0xee, 0x9b, 0x4d, 0x13, 0x5a, 0xc9, 0x78, 0x0f, 0x39, 0x31, 0x33, 0x0c, 0xe6, 0x6a, 0x79, 0x77,
	0x2d, 0xb8, 0xfc, 0x73, 0x08, 0x6a, 0xfb, 0x4d, 0x67, 0x16, 0x24, 0x14, 0xf8, 0x5e, 0x07, 0x1e,
	0xff, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0xf4, 0x32, 0x73, 0xe0, 0x04, 0x00, 0x00,
}
