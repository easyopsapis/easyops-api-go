// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: micro_app.proto

package app_store

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//小产品包
type AppStoreMicroApp struct {
	//
	//小产品名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//小产品id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//图标url
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//小产品概要介绍
	Intro string `protobuf:"bytes,4,opt,name=intro,proto3" json:"intro" form:"intro"`
	//
	//功能预览
	Preview []string `protobuf:"bytes,5,rep,name=preview,proto3" json:"preview" form:"preview"`
	//
	//功能详细介绍
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description" form:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppStoreMicroApp) Reset()         { *m = AppStoreMicroApp{} }
func (m *AppStoreMicroApp) String() string { return proto.CompactTextString(m) }
func (*AppStoreMicroApp) ProtoMessage()    {}
func (*AppStoreMicroApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c3c3ea0841e904e, []int{0}
}
func (m *AppStoreMicroApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppStoreMicroApp.Unmarshal(m, b)
}
func (m *AppStoreMicroApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppStoreMicroApp.Marshal(b, m, deterministic)
}
func (m *AppStoreMicroApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppStoreMicroApp.Merge(m, src)
}
func (m *AppStoreMicroApp) XXX_Size() int {
	return xxx_messageInfo_AppStoreMicroApp.Size(m)
}
func (m *AppStoreMicroApp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppStoreMicroApp.DiscardUnknown(m)
}

var xxx_messageInfo_AppStoreMicroApp proto.InternalMessageInfo

func (m *AppStoreMicroApp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppStoreMicroApp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AppStoreMicroApp) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *AppStoreMicroApp) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func (m *AppStoreMicroApp) GetPreview() []string {
	if m != nil {
		return m.Preview
	}
	return nil
}

func (m *AppStoreMicroApp) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AppStoreMicroApp)(nil), "app_store.AppStoreMicroApp")
}

func init() { proto.RegisterFile("micro_app.proto", fileDescriptor_2c3c3ea0841e904e) }

var fileDescriptor_2c3c3ea0841e904e = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0x87, 0x89, 0xf3, 0x67, 0xb1, 0xb2, 0x6c, 0xb2, 0x5a, 0x58, 0x4c, 0x2e, 0x0e, 0xde, 0x10,
	0x72, 0x88, 0x2d, 0xb6, 0x29, 0xa5, 0xe4, 0x96, 0x94, 0x1e, 0x7b, 0x71, 0x6f, 0x29, 0x6d, 0x50,
	0x6c, 0xc5, 0x15, 0xb5, 0x3d, 0x42, 0x56, 0x12, 0xda, 0xd2, 0x07, 0xeb, 0xcb, 0xb8, 0xd0, 0x47,
	0xf0, 0x13, 0x14, 0xcb, 0x69, 0x9a, 0x4b, 0x6f, 0x33, 0xf3, 0xfb, 0x3e, 0x8d, 0x60, 0x50, 0x27,
	0xe1, 0x81, 0x84, 0x25, 0x15, 0xc2, 0x13, 0x12, 0x14, 0x60, 0x93, 0x0a, 0xb1, 0xcc, 0x14, 0x48,
	0xd6, 0x73, 0x23, 0xae, 0xee, 0x37, 0x2b, 0x2f, 0x80, 0x84, 0x44, 0x10, 0x01, 0xd1, 0xc4, 0x6a,
	0xb3, 0xd6, 0x9d, 0x6e, 0x74, 0x55, 0x99, 0xbd, 0xb3, 0x23, 0x3c, 0xd9, 0x71, 0xf5, 0x00, 0x3b,
	0x12, 0x81, 0xab, 0x43, 0x77, 0x4b, 0x63, 0x1e, 0x52, 0x05, 0x32, 0x23, 0x87, 0xb2, 0xf2, 0x9c,
	0x57, 0x03, 0x75, 0x67, 0x42, 0x5c, 0x97, 0x3b, 0xaf, 0xca, 0xdf, 0xcc, 0x84, 0xc0, 0x53, 0xd4,
	0x48, 0x69, 0xc2, 0xac, 0x5a, 0xbf, 0x36, 0x32, 0xe7, 0xc3, 0x22, 0xb7, 0xdb, 0x6b, 0x90, 0xc9,
	0xd4, 0x29, 0xa7, 0xce, 0xfb, 0x9b, 0xfd, 0x07, 0xfd, 0xbe, 0xbb, 0xa1, 0xee, 0xd3, 0xcc, 0x5d,
	0x2c, 0x6f, 0x9f, 0xff, 0x8f, 0x27, 0x27, 0x2f, 0x03, 0x5f, 0x3b, 0xf8, 0x14, 0x19, 0x3c, 0xb4,
	0x0c, 0x6d, 0x0e, 0x8a, 0xdc, 0x36, 0x2b, 0x93, 0x87, 0xdf, 0x7a, 0x06, 0x0f, 0xf1, 0x3f, 0xd4,
	0xe0, 0x01, 0xa4, 0x56, 0x5d, 0x7b, 0x9d, 0xaf, 0x8d, 0xe5, 0xd4, 0xf1, 0x75, 0x88, 0x87, 0xa8,
	0xc9, 0x53, 0x25, 0xc1, 0x6a, 0x68, 0xaa, 0x5b, 0xe4, 0xf6, 0xcf, 0x3d, 0x55, 0x8e, 0x1d, 0xbf,
	0x8a, 0xf1, 0x18, 0xfd, 0x10, 0x92, 0x6d, 0x39, 0xdb, 0x59, 0xcd, 0x7e, 0x7d, 0x64, 0xce, 0x71,
	0x91, 0xdb, 0xbf, 0x2a, 0x72, 0x1f, 0x38, 0xfe, 0x27, 0x82, 0xcf, 0x51, 0x3b, 0x64, 0x59, 0x20,
	0xb9, 0x50, 0x1c, 0x52, 0xab, 0xa5, 0xdf, 0xfe, 0x5b, 0xe4, 0x36, 0xae, 0x8c, 0xa3, 0xd0, 0xf1,
	0x8f, 0xd1, 0xf9, 0xe5, 0xe2, 0x22, 0x02, 0x8f, 0xd1, 0xec, 0x11, 0x44, 0xe6, 0xc5, 0x10, 0xd0,
	0x98, 0x04, 0x90, 0x2a, 0x49, 0x03, 0x95, 0x55, 0xf7, 0x92, 0x4c, 0x80, 0x9b, 0x40, 0xc8, 0xe2,
	0x8c, 0xec, 0x41, 0xa2, 0x5b, 0x72, 0xb8, 0xf4, 0xaa, 0xa5, 0xc9, 0xc9, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x73, 0x16, 0xa7, 0x0e, 0x02, 0x00, 0x00,
}