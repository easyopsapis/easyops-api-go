// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

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
//Update请求
type UpdateRequest struct {
	//
	//版本Id
	VersionId string `protobuf:"bytes,1,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//包名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//版本备注
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//基础Docker镜像Id
	BaseImageId          string   `protobuf:"bytes,4,opt,name=baseImageId,proto3" json:"baseImageId" form:"baseImageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{0}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *UpdateRequest) GetBaseImageId() string {
	if m != nil {
		return m.BaseImageId
	}
	return ""
}

//
//UpdateApi返回
type UpdateResponseWrapper struct {
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

func (m *UpdateResponseWrapper) Reset()         { *m = UpdateResponseWrapper{} }
func (m *UpdateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseWrapper) ProtoMessage()    {}
func (*UpdateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{1}
}
func (m *UpdateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseWrapper.Merge(m, src)
}
func (m *UpdateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseWrapper.Size(m)
}
func (m *UpdateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseWrapper proto.InternalMessageInfo

func (m *UpdateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateResponseWrapper) GetData() *artifact.Version {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "version.UpdateRequest")
	proto.RegisterType((*UpdateResponseWrapper)(nil), "version.UpdateResponseWrapper")
}

func init() { proto.RegisterFile("update.proto", fileDescriptor_3f0fa214029f1c21) }

var fileDescriptor_3f0fa214029f1c21 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6a, 0xd4, 0x40,
	0x18, 0xed, 0xae, 0xa9, 0xd2, 0xd9, 0x8a, 0x75, 0xa0, 0xba, 0x14, 0x21, 0x32, 0x8a, 0x28, 0x92,
	0x8c, 0x56, 0x29, 0xd5, 0x1b, 0xb1, 0xe0, 0xc5, 0xde, 0xc9, 0xe0, 0x0f, 0x28, 0x0a, 0xb3, 0xc9,
	0xb7, 0x31, 0x98, 0xe4, 0x1b, 0x67, 0x26, 0x6d, 0x55, 0xfa, 0x02, 0x3e, 0x8c, 0x8f, 0xe3, 0xe5,
	0x0a, 0x3e, 0xc2, 0x3e, 0x81, 0xe4, 0x4b, 0xba, 0x89, 0x97, 0x5e, 0x65, 0xce, 0x77, 0xce, 0xc9,
	0x39, 0xdf, 0x30, 0x6c, 0xbb, 0x36, 0xa9, 0xf6, 0x10, 0x1b, 0x8b, 0x1e, 0xf9, 0xa5, 0x63, 0xb0,
	0x2e, 0xc7, 0x6a, 0x2f, 0xca, 0x72, 0xff, 0xa9, 0x9e, 0xc7, 0x09, 0x96, 0x32, 0xc3, 0x0c, 0x25,
	0xf1, 0xf3, 0x7a, 0x41, 0x88, 0x00, 0x9d, 0x5a, 0xdf, 0xde, 0xcb, 0x0c, 0x63, 0xd0, 0xee, 0x2b,
	0x1a, 0x17, 0x17, 0x98, 0xe8, 0x42, 0x26, 0x58, 0x79, 0xab, 0x13, 0xef, 0x5a, 0xa7, 0x05, 0x83,
	0x51, 0x89, 0x29, 0x14, 0x4e, 0x76, 0x42, 0x49, 0x50, 0x6a, 0xeb, 0xf3, 0x85, 0x4e, 0xbc, 0xec,
	0xa2, 0xbb, 0x3f, 0x1e, 0x0c, 0x0a, 0x94, 0x27, 0xb9, 0xff, 0x8c, 0x27, 0x32, 0xc3, 0x88, 0xc8,
	0xe8, 0x58, 0x17, 0x79, 0xaa, 0x3d, 0x5a, 0x27, 0xd7, 0xc7, 0xce, 0x77, 0x23, 0x43, 0xcc, 0x0a,
	0xe8, 0xfb, 0x3a, 0x6f, 0xeb, 0xc4, 0xb7, 0xac, 0xf8, 0x39, 0x66, 0x97, 0x5f, 0xd3, 0xc2, 0x0a,
	0xbe, 0xd4, 0xe0, 0x3c, 0xff, 0x31, 0x62, 0x5b, 0x5d, 0xf2, 0x2c, 0x9d, 0x8e, 0x6e, 0x8e, 0xee,
	0x6e, 0x1d, 0x15, 0xab, 0x65, 0xb8, 0xb3, 0x40, 0x5b, 0x3e, 0x15, 0x6b, 0x4a, 0xfc, 0xf9, 0x1d,
	0xbe, 0x62, 0xea, 0xe3, 0xfb, 0x67, 0x3a, 0xfa, 0xf6, 0x3c, 0x7a, 0xf7, 0x20, 0x7a, 0xf2, 0xe1,
	0xfb, 0xe1, 0x59, 0xf4, 0x0f, 0x7e, 0xfc, 0x9f, 0xf8, 0xe1, 0xfe, 0xd9, 0x6d, 0xd5, 0xc7, 0xf3,
	0xfb, 0x2c, 0xa8, 0x74, 0x09, 0xd3, 0x31, 0xd5, 0xb8, 0xbe, 0x5a, 0x86, 0x93, 0xb6, 0x46, 0x33,
	0x6d, 0x1a, 0x04, 0x66, 0xe3, 0x34, 0x56, 0x24, 0xe2, 0xf7, 0x58, 0x50, 0x42, 0x89, 0xd3, 0x0b,
	0x24, 0xde, 0xed, 0xc5, 0xcd, 0xb4, 0x11, 0x8f, 0xcd, 0x86, 0x22, 0x09, 0x3f, 0x64, 0x93, 0xb9,
	0x76, 0x30, 0x2b, 0x75, 0x06, 0xb3, 0x74, 0x1a, 0x90, 0xe3, 0xda, 0x6a, 0x19, 0xf2, 0xd6, 0x31,
	0x20, 0x85, 0x1a, 0x4a, 0xc5, 0xaf, 0x11, 0xdb, 0x3d, 0xbf, 0x30, 0x67, 0xb0, 0x72, 0xf0, 0xd6,
	0x6a, 0x63, 0xc0, 0xf2, 0x5b, 0x2c, 0x48, 0x30, 0x05, 0xba, 0xb2, 0xcd, 0xa3, 0x2b, 0x7d, 0x7c,
	0x33, 0x15, 0x8a, 0xc8, 0x26, 0xb8, 0xf9, 0xbe, 0x38, 0x35, 0x85, 0xce, 0xab, 0x6e, 0xaf, 0x41,
	0xf0, 0x80, 0x14, 0x6a, 0x28, 0xe5, 0x77, 0xd8, 0x26, 0x58, 0x8b, 0xb6, 0x5b, 0x6f, 0x67, 0xb5,
	0x0c, 0xb7, 0x5b, 0x0f, 0x8d, 0x85, 0x6a, 0x69, 0x7e, 0xc0, 0x82, 0x54, 0x7b, 0x4d, 0x3b, 0x4d,
	0xf6, 0xaf, 0xc6, 0xe7, 0xcf, 0x29, 0x7e, 0xd3, 0xde, 0xea, 0xb0, 0x59, 0x23, 0x14, 0x8a, 0xf4,
	0xf3, 0x8b, 0xf4, 0x20, 0x1e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x29, 0xbd, 0xb8, 0x00,
	0x03, 0x00, 0x00,
}