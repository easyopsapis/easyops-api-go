// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_version_env_type.proto

package version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//UpdateVersionEnvType请求
type UpdateVersionEnvTypeRequest struct {
	//
	//程序包Id
	PackageId string `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId string `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//版本类型 1 开发, 3 测试, 7 预发布, 15 生产
	EnvType              int32    `protobuf:"varint,3,opt,name=env_type,json=envType,proto3" json:"env_type" form:"env_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVersionEnvTypeRequest) Reset()         { *m = UpdateVersionEnvTypeRequest{} }
func (m *UpdateVersionEnvTypeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVersionEnvTypeRequest) ProtoMessage()    {}
func (*UpdateVersionEnvTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab43a2d06399d67d, []int{0}
}
func (m *UpdateVersionEnvTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVersionEnvTypeRequest.Unmarshal(m, b)
}
func (m *UpdateVersionEnvTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVersionEnvTypeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateVersionEnvTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVersionEnvTypeRequest.Merge(m, src)
}
func (m *UpdateVersionEnvTypeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVersionEnvTypeRequest.Size(m)
}
func (m *UpdateVersionEnvTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVersionEnvTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVersionEnvTypeRequest proto.InternalMessageInfo

func (m *UpdateVersionEnvTypeRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *UpdateVersionEnvTypeRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *UpdateVersionEnvTypeRequest) GetEnvType() int32 {
	if m != nil {
		return m.EnvType
	}
	return 0
}

//
//UpdateVersionEnvTypeApi返回
type UpdateVersionEnvTypeResponseWrapper struct {
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

func (m *UpdateVersionEnvTypeResponseWrapper) Reset()         { *m = UpdateVersionEnvTypeResponseWrapper{} }
func (m *UpdateVersionEnvTypeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateVersionEnvTypeResponseWrapper) ProtoMessage()    {}
func (*UpdateVersionEnvTypeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab43a2d06399d67d, []int{1}
}
func (m *UpdateVersionEnvTypeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateVersionEnvTypeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateVersionEnvTypeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper.Merge(m, src)
}
func (m *UpdateVersionEnvTypeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper.Size(m)
}
func (m *UpdateVersionEnvTypeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVersionEnvTypeResponseWrapper proto.InternalMessageInfo

func (m *UpdateVersionEnvTypeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateVersionEnvTypeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateVersionEnvTypeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateVersionEnvTypeResponseWrapper) GetData() *artifact.Version {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateVersionEnvTypeRequest)(nil), "version.UpdateVersionEnvTypeRequest")
	proto.RegisterType((*UpdateVersionEnvTypeResponseWrapper)(nil), "version.UpdateVersionEnvTypeResponseWrapper")
}

func init() { proto.RegisterFile("update_version_env_type.proto", fileDescriptor_ab43a2d06399d67d) }

var fileDescriptor_ab43a2d06399d67d = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0xda, 0xba, 0x76, 0x2a, 0xec, 0x1a, 0x41, 0xca, 0xaa, 0xa4, 0xcc, 0x8a, 0xf4,
	0x26, 0x19, 0x5d, 0x65, 0x59, 0xbd, 0x11, 0x0b, 0x7b, 0xb1, 0x77, 0x12, 0x56, 0x05, 0x45, 0xcb,
	0x34, 0x39, 0x8d, 0x61, 0xd3, 0x9c, 0x71, 0x66, 0xd2, 0xb5, 0x4a, 0x5f, 0xc0, 0x87, 0x8c, 0x20,
	0x82, 0xf7, 0x79, 0x02, 0xe9, 0xcc, 0xf4, 0x8f, 0xe2, 0x8d, 0x17, 0x5e, 0x25, 0xe7, 0x7c, 0xdf,
	0x99, 0xf3, 0xe3, 0xe3, 0x90, 0xbb, 0x95, 0x48, 0xb9, 0x86, 0xd1, 0x0c, 0xa4, 0xca, 0xb1, 0x1c,
	0x41, 0x39, 0x1b, 0xe9, 0xb9, 0x80, 0x48, 0x48, 0xd4, 0xe8, 0xef, 0xba, 0xfe, 0x41, 0x98, 0xe5,
	0xfa, 0x43, 0x35, 0x8e, 0x12, 0x9c, 0xb2, 0x0c, 0x33, 0x64, 0x46, 0x1f, 0x57, 0x13, 0x53, 0x99,
	0xc2, 0xfc, 0xd9, 0xb9, 0x83, 0x17, 0x19, 0x46, 0xc0, 0xd5, 0x1c, 0x85, 0x8a, 0x0a, 0x4c, 0x78,
	0xc1, 0x12, 0x2c, 0xb5, 0xe4, 0x89, 0x56, 0x76, 0x52, 0x82, 0xc0, 0x70, 0x8a, 0x29, 0x14, 0x8a,
	0x39, 0x23, 0x33, 0x25, 0xe3, 0x52, 0xe7, 0x13, 0x9e, 0x68, 0xe6, 0x56, 0xbb, 0x17, 0x8f, 0xb7,
	0x00, 0xa6, 0x97, 0xb9, 0xbe, 0xc0, 0x4b, 0x96, 0x61, 0x68, 0xc4, 0x70, 0xc6, 0x8b, 0x3c, 0xe5,
	0x1a, 0xa5, 0x62, 0xeb, 0x5f, 0x37, 0x77, 0x27, 0x43, 0xcc, 0x0a, 0xd8, 0xf0, 0x2a, 0x2d, 0xab,
	0x44, 0x5b, 0x95, 0xfe, 0xdc, 0x21, 0xb7, 0x5f, 0x9a, 0x04, 0x5e, 0xd9, 0x6d, 0xa7, 0xe5, 0xec,
	0x7c, 0x2e, 0x20, 0x86, 0x8f, 0x15, 0x28, 0xed, 0x7f, 0xf5, 0x48, 0x47, 0xf0, 0xe4, 0x82, 0x67,
	0x70, 0x96, 0xf6, 0xbc, 0xbe, 0x37, 0xe8, 0x0c, 0x8b, 0xa6, 0x0e, 0xf6, 0x27, 0x28, 0xa7, 0x4f,
	0xe9, 0x5a, 0xa2, 0xdf, 0xbf, 0x05, 0xe7, 0x24, 0x7e, 0xff, 0xf6, 0x19, 0x0f, 0x3f, 0x3f, 0x0f,
	0xdf, 0x3c, 0x08, 0x9f, 0xbc, 0xfb, 0x72, 0xb2, 0x08, 0x7f, 0xab, 0x1f, 0xff, 0x63, 0xfd, 0xf0,
	0x68, 0x71, 0x2f, 0xde, 0xac, 0x37, 0x30, 0x2e, 0x94, 0xb3, 0xb4, 0xb7, 0xf3, 0x27, 0xcc, 0x5a,
	0xfa, 0x8f, 0x30, 0xeb, 0x1d, 0x7e, 0x44, 0xae, 0xad, 0x6e, 0xa5, 0x77, 0xa5, 0xef, 0x0d, 0xda,
	0xc3, 0x9b, 0x4d, 0x1d, 0xec, 0x59, 0x94, 0x95, 0x42, 0xe3, 0x5d, 0xb0, 0x81, 0xd2, 0x1f, 0x1e,
	0x39, 0xfc, 0x7b, 0xd2, 0x4a, 0x60, 0xa9, 0xe0, 0xb5, 0xe4, 0x42, 0x80, 0xf4, 0x0f, 0x49, 0x2b,
	0xc1, 0x14, 0x4c, 0xd6, 0xed, 0xe1, 0x5e, 0x53, 0x07, 0x5d, 0xfb, 0xe6, 0xb2, 0x4b, 0x63, 0x23,
	0xfa, 0x27, 0xa4, 0xbb, 0xfc, 0x9e, 0x7e, 0x12, 0x05, 0xcf, 0x4b, 0x17, 0xc5, 0xad, 0xa6, 0x0e,
	0xfc, 0x8d, 0xd7, 0x89, 0x34, 0xde, 0xb6, 0xfa, 0xf7, 0x49, 0x1b, 0xa4, 0x44, 0x69, 0x98, 0x3b,
	0xc3, 0xfd, 0xa6, 0x0e, 0xae, 0x3b, 0xe6, 0x65, 0x9b, 0xc6, 0x56, 0xf6, 0x8f, 0x49, 0x2b, 0xe5,
	0x9a, 0xf7, 0x5a, 0x7d, 0x6f, 0xd0, 0x3d, 0xba, 0x11, 0xad, 0xae, 0x32, 0x72, 0xf4, 0xdb, 0x64,
	0x4b, 0x23, 0x8d, 0x8d, 0x7f, 0x7c, 0xd5, 0xdc, 0xd5, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x30, 0x7e, 0x69, 0x58, 0x03, 0x00, 0x00,
}