// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: init_package_permission.proto

package pkg

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//InitPackagePermissionApi返回
type InitPackagePermissionResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InitPackagePermissionResponseWrapper) Reset()         { *m = InitPackagePermissionResponseWrapper{} }
func (m *InitPackagePermissionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InitPackagePermissionResponseWrapper) ProtoMessage()    {}
func (*InitPackagePermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37133193ae4c3e6, []int{0}
}
func (m *InitPackagePermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitPackagePermissionResponseWrapper.Unmarshal(m, b)
}
func (m *InitPackagePermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitPackagePermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InitPackagePermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitPackagePermissionResponseWrapper.Merge(m, src)
}
func (m *InitPackagePermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InitPackagePermissionResponseWrapper.Size(m)
}
func (m *InitPackagePermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InitPackagePermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InitPackagePermissionResponseWrapper proto.InternalMessageInfo

func (m *InitPackagePermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InitPackagePermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InitPackagePermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InitPackagePermissionResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InitPackagePermissionResponseWrapper)(nil), "pkg.InitPackagePermissionResponseWrapper")
}

func init() { proto.RegisterFile("init_package_permission.proto", fileDescriptor_a37133193ae4c3e6) }

var fileDescriptor_a37133193ae4c3e6 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xeb, 0x04, 0x53, 0x41, 0xc9, 0x61, 0x94, 0x89, 0xb4, 0x44, 0x91, 0x5e, 0xcc,
	0x40, 0x2f, 0xa2, 0xb7, 0xc1, 0x0e, 0xde, 0x46, 0x2e, 0x1e, 0x47, 0xda, 0x65, 0x31, 0x6c, 0xcd,
	0x0b, 0x69, 0x06, 0xfa, 0x65, 0xfb, 0x01, 0x3c, 0xf6, 0x13, 0x48, 0x13, 0xdd, 0x76, 0x6a, 0xdf,
	0xfb, 0xfd, 0xfe, 0xbc, 0xfc, 0xf1, 0xad, 0x36, 0xda, 0xaf, 0xac, 0xa8, 0xb7, 0x42, 0xc9, 0x95,
	0x95, 0xae, 0xd1, 0x6d, 0xab, 0xc1, 0x30, 0xeb, 0xc0, 0x03, 0x19, 0xd9, 0xad, 0x9a, 0x3e, 0x2a,
	0xed, 0x3f, 0xf7, 0x15, 0xab, 0xa1, 0x99, 0x29, 0x50, 0x30, 0x0b, 0xac, 0xda, 0x6f, 0xc2, 0x14,
	0x86, 0xf0, 0x17, 0x33, 0xd3, 0x1b, 0x05, 0xa0, 0x76, 0xf2, 0x68, 0xc9, 0xc6, 0xfa, 0xef, 0x08,
	0xe9, 0x0f, 0xc2, 0xf7, 0xef, 0x46, 0xfb, 0x65, 0xbc, 0xb8, 0x3c, 0x1c, 0xe4, 0xb2, 0xb5, 0x60,
	0x5a, 0xf9, 0xe1, 0x84, 0xb5, 0xd2, 0x91, 0x3b, 0x9c, 0xd4, 0xb0, 0x96, 0x19, 0x2a, 0x50, 0x39,
	0x9e, 0x5f, 0xf5, 0x5d, 0x9e, 0x6e, 0xc0, 0x35, 0xaf, 0x74, 0xd8, 0x52, 0x1e, 0x20, 0x79, 0xc1,
	0xe9, 0xf0, 0x5d, 0x7c, 0xd9, 0x9d, 0xd0, 0x26, 0x3b, 0x2b, 0x50, 0x79, 0x31, 0x9f, 0xf4, 0x5d,
	0x4e, 0x8e, 0xee, 0x1f, 0xa4, 0xfc, 0x54, 0x25, 0x0f, 0x78, 0x2c, 0x9d, 0x03, 0x97, 0x8d, 0x42,
	0xe6, 0xba, 0xef, 0xf2, 0xcb, 0x98, 0x09, 0x6b, 0xca, 0x23, 0x26, 0x6f, 0x38, 0x59, 0x0b, 0x2f,
	0xb2, 0xa4, 0x40, 0x65, 0xfa, 0x34, 0x61, 0xb1, 0x1b, 0xfb, 0xef, 0xc6, 0x16, 0x43, 0xb7, 0xd3,
	0xe7, 0x0d, 0x36, 0xe5, 0x21, 0x54, 0x9d, 0x07, 0xed, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x23,
	0x99, 0x7a, 0x82, 0x65, 0x01, 0x00, 0x00,
}