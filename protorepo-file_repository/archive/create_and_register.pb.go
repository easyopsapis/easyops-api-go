// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_and_register.proto

package archive

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//CreateAndRegister返回
type CreateAndRegisterResponse struct {
	//
	//版本Id
	VersionId string `protobuf:"bytes,1,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//包文件的md5检验值
	Sign *types.Struct `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign" form:"sign"`
	//
	//版本的进程及启停、安装、升级脚本, package.conf.yaml
	Conf                 string   `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf" form:"conf"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAndRegisterResponse) Reset()         { *m = CreateAndRegisterResponse{} }
func (m *CreateAndRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAndRegisterResponse) ProtoMessage()    {}
func (*CreateAndRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c64912abdcb14fc9, []int{0}
}
func (m *CreateAndRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAndRegisterResponse.Unmarshal(m, b)
}
func (m *CreateAndRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAndRegisterResponse.Marshal(b, m, deterministic)
}
func (m *CreateAndRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAndRegisterResponse.Merge(m, src)
}
func (m *CreateAndRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAndRegisterResponse.Size(m)
}
func (m *CreateAndRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAndRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAndRegisterResponse proto.InternalMessageInfo

func (m *CreateAndRegisterResponse) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *CreateAndRegisterResponse) GetSign() *types.Struct {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *CreateAndRegisterResponse) GetConf() string {
	if m != nil {
		return m.Conf
	}
	return ""
}

//
//CreateAndRegisterApi返回
type CreateAndRegisterResponseWrapper struct {
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
	Data                 *CreateAndRegisterResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CreateAndRegisterResponseWrapper) Reset()         { *m = CreateAndRegisterResponseWrapper{} }
func (m *CreateAndRegisterResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateAndRegisterResponseWrapper) ProtoMessage()    {}
func (*CreateAndRegisterResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c64912abdcb14fc9, []int{1}
}
func (m *CreateAndRegisterResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAndRegisterResponseWrapper.Unmarshal(m, b)
}
func (m *CreateAndRegisterResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAndRegisterResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateAndRegisterResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAndRegisterResponseWrapper.Merge(m, src)
}
func (m *CreateAndRegisterResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateAndRegisterResponseWrapper.Size(m)
}
func (m *CreateAndRegisterResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAndRegisterResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAndRegisterResponseWrapper proto.InternalMessageInfo

func (m *CreateAndRegisterResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateAndRegisterResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateAndRegisterResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateAndRegisterResponseWrapper) GetData() *CreateAndRegisterResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAndRegisterResponse)(nil), "archive.CreateAndRegisterResponse")
	proto.RegisterType((*CreateAndRegisterResponseWrapper)(nil), "archive.CreateAndRegisterResponseWrapper")
}

func init() { proto.RegisterFile("create_and_register.proto", fileDescriptor_c64912abdcb14fc9) }

var fileDescriptor_c64912abdcb14fc9 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdf, 0xaa, 0xd3, 0x30,
	0x18, 0xa7, 0xc7, 0x1e, 0x65, 0x99, 0xe0, 0x21, 0x17, 0xda, 0x73, 0x10, 0x3a, 0xa2, 0xc8, 0x6e,
	0xda, 0xea, 0x14, 0x99, 0x22, 0xc8, 0x26, 0x22, 0xde, 0x46, 0x41, 0x50, 0x74, 0x64, 0x4d, 0x96,
	0x05, 0xbb, 0xa6, 0xa4, 0xd9, 0x26, 0xca, 0x5e, 0x40, 0xdf, 0xb1, 0x82, 0x8f, 0xd0, 0x27, 0x90,
	0x7e, 0xa9, 0x5b, 0x77, 0xb1, 0x0b, 0xaf, 0x9a, 0x2f, 0xbf, 0x3f, 0xfd, 0x7d, 0x3f, 0x82, 0x2e,
	0x53, 0x23, 0x98, 0x15, 0x33, 0x96, 0xf3, 0x99, 0x11, 0x52, 0x95, 0x56, 0x98, 0xb8, 0x30, 0xda,
	0x6a, 0x7c, 0x83, 0x99, 0x74, 0xa9, 0x36, 0xe2, 0x2a, 0x92, 0xca, 0x2e, 0xd7, 0xf3, 0x38, 0xd5,
	0xab, 0x44, 0x6a, 0xa9, 0x13, 0xc0, 0xe7, 0xeb, 0x05, 0x4c, 0x30, 0xc0, 0xc9, 0xe9, 0xae, 0x9e,
	0x76, 0xe8, 0xab, 0xad, 0xb2, 0x5f, 0xf5, 0x36, 0x91, 0x3a, 0x02, 0x30, 0xda, 0xb0, 0x4c, 0x71,
	0x66, 0xb5, 0x29, 0x93, 0xfd, 0xb1, 0xd5, 0xdd, 0x95, 0x5a, 0xcb, 0x4c, 0x1c, 0xdc, 0x4b, 0x6b,
	0xd6, 0xa9, 0x75, 0x28, 0xf9, 0x75, 0x86, 0x2e, 0x5f, 0x41, 0xd6, 0x49, 0xce, 0x69, 0x9b, 0x94,
	0x8a, 0xb2, 0xd0, 0x79, 0x29, 0xf0, 0x4f, 0x0f, 0xf5, 0x36, 0xc2, 0x94, 0x4a, 0xe7, 0x6f, 0x79,
	0xe0, 0x0d, 0xbc, 0x61, 0x6f, 0x9a, 0xd5, 0x55, 0x78, 0xb1, 0xd0, 0x66, 0xf5, 0x9c, 0xec, 0x21,
	0xf2, 0xe7, 0x77, 0xf8, 0x1e, 0xd1, 0x2f, 0x9f, 0x5e, 0xb2, 0xe8, 0xfb, 0x24, 0xfa, 0xf8, 0x30,
	0x7a, 0xf6, 0xf9, 0xc7, 0x78, 0x17, 0x1d, 0xcd, 0x4f, 0xfe, 0x73, 0x7e, 0x34, 0xda, 0xdd, 0xa7,
	0x87, 0xdf, 0xe3, 0x17, 0xc8, 0x2f, 0x95, 0xcc, 0x83, 0xb3, 0x81, 0x37, 0xec, 0x8f, 0xee, 0xc4,
	0x6e, 0xaf, 0xf8, 0xdf, 0x5e, 0xf1, 0x3b, 0xd8, 0x6b, 0x7a, 0xab, 0xae, 0xc2, 0xbe, 0xcb, 0xd7,
	0xd0, 0x09, 0x05, 0x15, 0xbe, 0x87, 0xfc, 0x54, 0xe7, 0x8b, 0xe0, 0x1a, 0x2c, 0xd1, 0x21, 0x35,
	0xb7, 0x84, 0x02, 0x48, 0x6a, 0x0f, 0x0d, 0x4e, 0xb6, 0xf1, 0xc1, 0xb0, 0xa2, 0x10, 0xc6, 0x39,
	0x71, 0x01, 0x75, 0x9c, 0x1f, 0x3b, 0x71, 0x01, 0x4e, 0x5c, 0xe0, 0x31, 0xea, 0x37, 0xdf, 0xd7,
	0xdf, 0x8a, 0x8c, 0x29, 0x97, 0xb9, 0x37, 0xbd, 0x5d, 0x57, 0x21, 0x3e, 0x70, 0x5b, 0x90, 0xd0,
	0x2e, 0x15, 0x3f, 0x40, 0xe7, 0xc2, 0x18, 0x6d, 0xda, 0xa4, 0x17, 0x75, 0x15, 0xde, 0x74, 0x1a,
	0xb8, 0x26, 0xd4, 0xc1, 0xf8, 0x0d, 0xf2, 0x39, 0xb3, 0x2c, 0xf0, 0xa1, 0x0e, 0x12, 0xb7, 0xcf,
	0x2a, 0x3e, 0x99, 0xbf, 0x1b, 0xb5, 0x51, 0x12, 0x0a, 0x06, 0xf3, 0xeb, 0xd0, 0xe0, 0xe3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xd1, 0xbd, 0x7e, 0xb4, 0x02, 0x00, 0x00,
}
