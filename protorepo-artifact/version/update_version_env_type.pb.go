// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_version_env_type.proto

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
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0x66, 0xd6, 0xc4, 0x35, 0x1d, 0x61, 0xd7, 0x11, 0x24, 0xac, 0xc8, 0x84, 0x5e, 0x91, 0x5c,
	0x7a, 0x5a, 0x57, 0x59, 0x56, 0x2f, 0x62, 0x60, 0x85, 0xbd, 0xe9, 0xf8, 0x77, 0x10, 0x0d, 0x9d,
	0x99, 0xca, 0x38, 0x64, 0x32, 0xd5, 0x76, 0x77, 0xb2, 0x06, 0xd9, 0x07, 0xf0, 0x25, 0x47, 0x10,
	0x7d, 0x81, 0x79, 0x02, 0x49, 0x77, 0x27, 0x1b, 0xc4, 0x8b, 0x07, 0x4f, 0xdd, 0x5f, 0x7d, 0x5f,
	0x55, 0x7d, 0x55, 0x14, 0xb9, 0x33, 0x97, 0x99, 0x30, 0x30, 0x5a, 0x80, 0xd2, 0x05, 0x56, 0x23,
	0xa8, 0x16, 0x23, 0xb3, 0x94, 0x10, 0x4b, 0x85, 0x06, 0xc3, 0x5d, 0x1f, 0x3f, 0x60, 0x79, 0x61,
	0x3e, 0xcd, 0xc7, 0x71, 0x8a, 0x33, 0x9e, 0x63, 0x8e, 0xdc, 0xf2, 0xe3, 0xf9, 0xc4, 0x22, 0x0b,
	0xec, 0xcf, 0xe5, 0x1d, 0xbc, 0xc8, 0x31, 0x06, 0xa1, 0x97, 0x28, 0x75, 0x5c, 0x62, 0x2a, 0x4a,
	0x9e, 0x62, 0x65, 0x94, 0x48, 0x8d, 0x76, 0x99, 0x0a, 0x24, 0xb2, 0x19, 0x66, 0x50, 0x6a, 0xee,
	0x85, 0xdc, 0x42, 0x2e, 0x94, 0x29, 0x26, 0x22, 0x35, 0xdc, 0xb7, 0xf6, 0x15, 0x8f, 0xb7, 0x0c,
	0xcc, 0xce, 0x0b, 0x33, 0xc5, 0x73, 0x9e, 0x23, 0xb3, 0x24, 0x5b, 0x88, 0xb2, 0xc8, 0x84, 0x41,
	0xa5, 0xf9, 0xe6, 0xeb, 0xf2, 0xe8, 0xaf, 0x1d, 0x72, 0xfb, 0x8d, 0x9d, 0xf1, 0xad, 0xab, 0x77,
	0x5a, 0x2d, 0x5e, 0x2f, 0x25, 0x24, 0xf0, 0x79, 0x0e, 0xda, 0x84, 0xdf, 0x02, 0xd2, 0x91, 0x22,
	0x9d, 0x8a, 0x1c, 0xce, 0xb2, 0x5e, 0xd0, 0x0f, 0x06, 0x9d, 0xe1, 0xb4, 0xa9, 0xa3, 0xfd, 0x09,
	0xaa, 0xd9, 0x13, 0xba, 0xa1, 0xe8, 0x8f, 0xef, 0xd1, 0x2b, 0xf2, 0xf2, 0xe3, 0x7b, 0xc1, 0x26,
	0xcf, 0xd8, 0xf3, 0xfb, 0xec, 0xf1, 0x87, 0xaf, 0x27, 0x17, 0xec, 0xe9, 0x36, 0x7e, 0xf4, 0x8f,
	0xf8, 0xc1, 0xd1, 0xc5, 0xdd, 0xe4, 0xb2, 0xbb, 0xf5, 0xe2, 0xa7, 0x3e, 0xcb, 0x7a, 0x3b, 0x7f,
	0x7a, 0xd9, 0x50, 0xff, 0xcf, 0xcb, 0xa6, 0x45, 0x18, 0x93, 0x6b, 0xeb, 0x5b, 0xe8, 0x5d, 0xe9,
	0x07, 0x83, 0xf6, 0xf0, 0x66, 0x53, 0x47, 0x7b, 0xce, 0xc9, 0x9a, 0xa1, 0xc9, 0x2e, 0xb8, 0x75,
	0xd2, 0x9f, 0x01, 0x39, 0xfc, 0xfb, 0x9e, 0xb5, 0xc4, 0x4a, 0xc3, 0x3b, 0x25, 0xa4, 0x04, 0x15,
	0x1e, 0x92, 0x56, 0x8a, 0x19, 0xd8, 0x4d, 0xb7, 0x87, 0x7b, 0x4d, 0x1d, 0x75, 0x5d, 0xcd, 0x55,
	0x94, 0x26, 0x96, 0x0c, 0x4f, 0x48, 0x77, 0xf5, 0x9e, 0x7e, 0x91, 0xa5, 0x28, 0x2a, 0xbf, 0x89,
	0x5b, 0x4d, 0x1d, 0x85, 0x97, 0x5a, 0x4f, 0xd2, 0x64, 0x5b, 0x1a, 0xde, 0x23, 0x6d, 0x50, 0x0a,
	0x95, 0xf5, 0xdc, 0x19, 0xee, 0x37, 0x75, 0x74, 0xdd, 0x7b, 0x5e, 0x85, 0x69, 0xe2, 0xe8, 0xf0,
	0x98, 0xb4, 0x32, 0x61, 0x44, 0xaf, 0xd5, 0x0f, 0x06, 0xdd, 0xa3, 0x1b, 0xf1, 0xfa, 0xea, 0x62,
	0xef, 0x7e, 0xdb, 0xd9, 0x4a, 0x48, 0x13, 0xab, 0x1f, 0x5f, 0xb5, 0x57, 0xf5, 0xf0, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2d, 0x4d, 0xc0, 0x59, 0x38, 0x03, 0x00, 0x00,
}
