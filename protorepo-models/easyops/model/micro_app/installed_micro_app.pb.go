// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: installed_micro_app.proto

package micro_app

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
//小产品
type InstalledMicroApp struct {
	//
	//小产品名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//小产品id
	AppId string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//图标url
	Icons *InstalledMicroApp_Icons `protobuf:"bytes,3,opt,name=icons,proto3" json:"icons" form:"icons"`
	//
	//storyboard内容
	Storyboard string `protobuf:"bytes,4,opt,name=storyboard,proto3" json:"storyboard" form:"storyboard"`
	//
	//标签
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags" form:"tags"`
	//
	//当前版本
	CurrentVersion string `protobuf:"bytes,6,opt,name=currentVersion,proto3" json:"currentVersion" form:"currentVersion"`
	//
	//当前版本
	InstallStatus        string   `protobuf:"bytes,7,opt,name=installStatus,proto3" json:"installStatus" form:"installStatus"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstalledMicroApp) Reset()         { *m = InstalledMicroApp{} }
func (m *InstalledMicroApp) String() string { return proto.CompactTextString(m) }
func (*InstalledMicroApp) ProtoMessage()    {}
func (*InstalledMicroApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e99c71557ae3ac8, []int{0}
}
func (m *InstalledMicroApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstalledMicroApp.Unmarshal(m, b)
}
func (m *InstalledMicroApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstalledMicroApp.Marshal(b, m, deterministic)
}
func (m *InstalledMicroApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstalledMicroApp.Merge(m, src)
}
func (m *InstalledMicroApp) XXX_Size() int {
	return xxx_messageInfo_InstalledMicroApp.Size(m)
}
func (m *InstalledMicroApp) XXX_DiscardUnknown() {
	xxx_messageInfo_InstalledMicroApp.DiscardUnknown(m)
}

var xxx_messageInfo_InstalledMicroApp proto.InternalMessageInfo

func (m *InstalledMicroApp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstalledMicroApp) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *InstalledMicroApp) GetIcons() *InstalledMicroApp_Icons {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *InstalledMicroApp) GetStoryboard() string {
	if m != nil {
		return m.Storyboard
	}
	return ""
}

func (m *InstalledMicroApp) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *InstalledMicroApp) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *InstalledMicroApp) GetInstallStatus() string {
	if m != nil {
		return m.InstallStatus
	}
	return ""
}

type InstalledMicroApp_Icons struct {
	//
	//图标url
	Large                string   `protobuf:"bytes,1,opt,name=large,proto3" json:"large" form:"large"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstalledMicroApp_Icons) Reset()         { *m = InstalledMicroApp_Icons{} }
func (m *InstalledMicroApp_Icons) String() string { return proto.CompactTextString(m) }
func (*InstalledMicroApp_Icons) ProtoMessage()    {}
func (*InstalledMicroApp_Icons) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e99c71557ae3ac8, []int{0, 0}
}
func (m *InstalledMicroApp_Icons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstalledMicroApp_Icons.Unmarshal(m, b)
}
func (m *InstalledMicroApp_Icons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstalledMicroApp_Icons.Marshal(b, m, deterministic)
}
func (m *InstalledMicroApp_Icons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstalledMicroApp_Icons.Merge(m, src)
}
func (m *InstalledMicroApp_Icons) XXX_Size() int {
	return xxx_messageInfo_InstalledMicroApp_Icons.Size(m)
}
func (m *InstalledMicroApp_Icons) XXX_DiscardUnknown() {
	xxx_messageInfo_InstalledMicroApp_Icons.DiscardUnknown(m)
}

var xxx_messageInfo_InstalledMicroApp_Icons proto.InternalMessageInfo

func (m *InstalledMicroApp_Icons) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

func init() {
	proto.RegisterType((*InstalledMicroApp)(nil), "micro_app.InstalledMicroApp")
	proto.RegisterType((*InstalledMicroApp_Icons)(nil), "micro_app.InstalledMicroApp.Icons")
}

func init() { proto.RegisterFile("installed_micro_app.proto", fileDescriptor_7e99c71557ae3ac8) }

var fileDescriptor_7e99c71557ae3ac8 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x55, 0xd6, 0x0e, 0xd5, 0x03, 0x46, 0x03, 0x93, 0x42, 0x6f, 0x52, 0x99, 0x0a, 0xd2,
	0x95, 0xc4, 0x62, 0x13, 0x5c, 0x54, 0x30, 0xb4, 0x22, 0x84, 0x8a, 0xc4, 0x4d, 0x90, 0x76, 0xb1,
	0xba, 0xad, 0xdc, 0x24, 0x0b, 0x11, 0x49, 0x4e, 0x64, 0xbb, 0x4c, 0x83, 0xf2, 0x86, 0x3c, 0x43,
	0x90, 0x78, 0x84, 0x3c, 0x01, 0x8a, 0xd3, 0x36, 0x1d, 0xd3, 0x6e, 0xa2, 0x9c, 0x73, 0x7e, 0xff,
	0xf3, 0x69, 0xf4, 0x24, 0x4c, 0x84, 0x64, 0x51, 0xe4, 0x7b, 0xb3, 0x38, 0x74, 0x39, 0xcc, 0x58,
	0x9a, 0xda, 0x29, 0x07, 0x09, 0x5a, 0x73, 0xe3, 0x68, 0x5b, 0x41, 0x28, 0xbf, 0x2e, 0xe6, 0xb6,
	0x0b, 0x31, 0x09, 0x20, 0x00, 0xa2, 0x88, 0xf9, 0xe2, 0x42, 0x59, 0xca, 0x50, 0x7f, 0xa5, 0xb2,
	0xfd, 0x7a, 0x0b, 0x8f, 0x2f, 0x43, 0xf9, 0x0d, 0x2e, 0x49, 0x00, 0x96, 0x0a, 0x5a, 0xdf, 0x59,
	0x14, 0x7a, 0x4c, 0x02, 0x17, 0x64, 0xf3, 0x5b, 0xea, 0xf0, 0xef, 0x3a, 0x6a, 0x8d, 0xd6, 0xfd,
	0x7c, 0x2e, 0xaa, 0x9f, 0xa6, 0xa9, 0x36, 0x40, 0xf5, 0x84, 0xc5, 0xbe, 0x5e, 0xeb, 0xd4, 0xcc,
	0xe6, 0xf0, 0x59, 0x9e, 0x19, 0x7b, 0x17, 0xc0, 0xe3, 0x01, 0x2e, 0xbc, 0xf8, 0xef, 0x1f, 0xe3,
	0x11, 0x6a, 0x4d, 0xc7, 0xcc, 0xfa, 0x71, 0x6a, 0x9d, 0xcf, 0x26, 0x3f, 0x5f, 0xbe, 0x38, 0x3e,
	0xfa, 0xd5, 0x75, 0x94, 0x46, 0x7b, 0x8b, 0x1a, 0x2c, 0x4d, 0x47, 0x9e, 0x7e, 0x47, 0x89, 0x9f,
	0xe7, 0x99, 0x71, 0xaf, 0x14, 0x2b, 0xf7, 0xad, 0xea, 0x52, 0xa5, 0x7d, 0x42, 0x8d, 0xd0, 0x85,
	0x44, 0xe8, 0x3b, 0x9d, 0x9a, 0xb9, 0x77, 0x84, 0xed, 0x6a, 0x47, 0x37, 0xfa, 0xb4, 0x47, 0x05,
	0x39, 0x7c, 0x58, 0x95, 0x50, 0x52, 0xec, 0x94, 0x29, 0xb4, 0x57, 0x08, 0x09, 0x09, 0xfc, 0x6a,
	0x0e, 0x8c, 0x7b, 0x7a, 0x5d, 0xf5, 0x73, 0x90, 0x67, 0x46, 0xab, 0x84, 0xab, 0x18, 0x76, 0xb6,
	0x40, 0xed, 0x29, 0xaa, 0x4b, 0x16, 0x08, 0xbd, 0xd1, 0xd9, 0x31, 0x9b, 0xc3, 0xfd, 0x6a, 0xfa,
	0xc2, 0x8b, 0x1d, 0x15, 0xd4, 0xce, 0xd0, 0x03, 0x77, 0xc1, 0xb9, 0x9f, 0xc8, 0x33, 0x9f, 0x8b,
	0x10, 0x12, 0x7d, 0x57, 0xe5, 0xb7, 0xf3, 0xcc, 0x38, 0x28, 0xf1, 0xeb, 0xf1, 0x62, 0xf0, 0x16,
	0xda, 0x9f, 0x52, 0xaf, 0x4f, 0xed, 0xf5, 0xa7, 0xeb, 0xfc, 0x97, 0x45, 0x3b, 0x41, 0xf7, 0x57,
	0xef, 0xe3, 0x8b, 0x64, 0x72, 0x21, 0xf4, 0xbb, 0x2a, 0xad, 0x9e, 0x67, 0xc6, 0xe3, 0xd5, 0x8c,
	0xdb, 0x61, 0xec, 0x5c, 0xc7, 0xdb, 0x4b, 0xd4, 0x50, 0x5b, 0xd1, 0x04, 0x6a, 0x44, 0x8c, 0x07,
	0xeb, 0x23, 0x4e, 0xaa, 0x25, 0x29, 0x77, 0xd1, 0xce, 0x08, 0x7d, 0x9c, 0x9a, 0x26, 0x25, 0xe3,
	0x29, 0x25, 0x03, 0x7a, 0x48, 0xdf, 0x61, 0xfc, 0xe6, 0x84, 0x2e, 0x29, 0xa7, 0xc9, 0xa4, 0xdf,
	0xeb, 0xf7, 0x96, 0x26, 0x25, 0xbd, 0xe5, 0xea, 0x52, 0x93, 0x81, 0x49, 0xe9, 0x78, 0x4a, 0xe9,
	0x4d, 0xf2, 0xb0, 0xeb, 0x94, 0xb5, 0x86, 0x1f, 0xce, 0xdf, 0x07, 0x60, 0xfb, 0x4c, 0x5c, 0x41,
	0x2a, 0xec, 0x08, 0x5c, 0x16, 0x11, 0x17, 0x12, 0xc9, 0x99, 0x2b, 0x45, 0xf9, 0x84, 0xb9, 0x9f,
	0x82, 0x15, 0x83, 0xe7, 0x47, 0x82, 0xac, 0x40, 0xa2, 0x4c, 0xb2, 0xb9, 0xf4, 0x7c, 0x57, 0x91,
	0xc7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x29, 0xb9, 0xc5, 0x2b, 0x03, 0x00, 0x00,
}