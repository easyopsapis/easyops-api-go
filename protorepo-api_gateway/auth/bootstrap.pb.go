// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bootstrap.proto

package auth

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	auth "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/auth"
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
//Bootstrap请求
type BootstrapRequest struct {
	//
	//是否需要检查登录态
	CheckLogin           bool     `protobuf:"varint,1,opt,name=check_login,json=checkLogin,proto3" json:"check_login" form:"check_login"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootstrapRequest) Reset()         { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()    {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3cc050d1b41a9e, []int{0}
}
func (m *BootstrapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapRequest.Unmarshal(m, b)
}
func (m *BootstrapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapRequest.Marshal(b, m, deterministic)
}
func (m *BootstrapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapRequest.Merge(m, src)
}
func (m *BootstrapRequest) XXX_Size() int {
	return xxx_messageInfo_BootstrapRequest.Size(m)
}
func (m *BootstrapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapRequest proto.InternalMessageInfo

func (m *BootstrapRequest) GetCheckLogin() bool {
	if m != nil {
		return m.CheckLogin
	}
	return false
}

//
//Bootstrap返回
type BootstrapResponse struct {
	//
	//${BRICK_NEXT}/packages/brick-container/conf/navbar.json 的内容
	Navbar *types.Struct `protobuf:"bytes,1,opt,name=navbar,proto3" json:"navbar" form:"navbar"`
	//
	//${BRICK_NEXT}/packages/brick-container/conf/storyboards/ *.json 的内容列表
	Storyboards []*auth.StoryBoard `protobuf:"bytes,2,rep,name=storyboards,proto3" json:"storyboards" form:"storyboards"`
	//
	//${BRICK_NEXT}/bricks/ * /dist/ 的内容
	BrickPackages []*types.Struct `protobuf:"bytes,3,rep,name=brickPackages,proto3" json:"brickPackages" form:"brickPackages"`
	//
	//console特性配置
	Settings *auth.Settings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings" form:"settings"`
	//
	//桌面列表
	Desktops             []*auth.Desktop `protobuf:"bytes,5,rep,name=desktops,proto3" json:"desktops" form:"desktops"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BootstrapResponse) Reset()         { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()    {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3cc050d1b41a9e, []int{1}
}
func (m *BootstrapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapResponse.Unmarshal(m, b)
}
func (m *BootstrapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapResponse.Marshal(b, m, deterministic)
}
func (m *BootstrapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapResponse.Merge(m, src)
}
func (m *BootstrapResponse) XXX_Size() int {
	return xxx_messageInfo_BootstrapResponse.Size(m)
}
func (m *BootstrapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapResponse proto.InternalMessageInfo

func (m *BootstrapResponse) GetNavbar() *types.Struct {
	if m != nil {
		return m.Navbar
	}
	return nil
}

func (m *BootstrapResponse) GetStoryboards() []*auth.StoryBoard {
	if m != nil {
		return m.Storyboards
	}
	return nil
}

func (m *BootstrapResponse) GetBrickPackages() []*types.Struct {
	if m != nil {
		return m.BrickPackages
	}
	return nil
}

func (m *BootstrapResponse) GetSettings() *auth.Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *BootstrapResponse) GetDesktops() []*auth.Desktop {
	if m != nil {
		return m.Desktops
	}
	return nil
}

//
//BootstrapApi返回
type BootstrapResponseWrapper struct {
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
	Data                 *BootstrapResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BootstrapResponseWrapper) Reset()         { *m = BootstrapResponseWrapper{} }
func (m *BootstrapResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*BootstrapResponseWrapper) ProtoMessage()    {}
func (*BootstrapResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3cc050d1b41a9e, []int{2}
}
func (m *BootstrapResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapResponseWrapper.Unmarshal(m, b)
}
func (m *BootstrapResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapResponseWrapper.Marshal(b, m, deterministic)
}
func (m *BootstrapResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapResponseWrapper.Merge(m, src)
}
func (m *BootstrapResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_BootstrapResponseWrapper.Size(m)
}
func (m *BootstrapResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapResponseWrapper proto.InternalMessageInfo

func (m *BootstrapResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BootstrapResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *BootstrapResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BootstrapResponseWrapper) GetData() *BootstrapResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BootstrapRequest)(nil), "auth.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "auth.BootstrapResponse")
	proto.RegisterType((*BootstrapResponseWrapper)(nil), "auth.BootstrapResponseWrapper")
}

func init() { proto.RegisterFile("bootstrap.proto", fileDescriptor_5e3cc050d1b41a9e) }

var fileDescriptor_5e3cc050d1b41a9e = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0xa5, 0xff, 0x96, 0xfe, 0x26, 0xbf, 0xda, 0xee, 0x28, 0x6e, 0x28, 0x42, 0xca, 0x08, 0xd2,
	0x97, 0x4d, 0x60, 0x7d, 0x50, 0x44, 0x14, 0x82, 0xbe, 0xe8, 0x8a, 0x32, 0xfb, 0xb0, 0x8f, 0x32,
	0x99, 0xcc, 0xa6, 0xa1, 0x69, 0x6e, 0x9c, 0x99, 0x88, 0xfb, 0x65, 0xe3, 0x77, 0xc8, 0xbb, 0x20,
	0x99, 0x99, 0x6e, 0x53, 0x16, 0x9f, 0xf6, 0x29, 0xb9, 0xf7, 0xdc, 0x73, 0xe6, 0xdc, 0x7b, 0xd0,
	0x3c, 0x01, 0xd0, 0x4a, 0x4b, 0x56, 0x85, 0x95, 0x04, 0x0d, 0x78, 0xcc, 0x6a, 0xbd, 0x59, 0x9e,
	0x67, 0xb9, 0xde, 0xd4, 0x49, 0xc8, 0x61, 0x17, 0x65, 0x90, 0x41, 0x64, 0xc0, 0xa4, 0xbe, 0x31,
	0x95, 0x29, 0xcc, 0x9f, 0x25, 0x2d, 0xbf, 0x66, 0x10, 0x0a, 0xa6, 0x6e, 0xa1, 0x52, 0x61, 0x01,
	0x9c, 0x15, 0x11, 0x87, 0x52, 0x4b, 0xc6, 0xb5, 0xb2, 0x4c, 0x29, 0x2a, 0x38, 0xdf, 0x41, 0x2a,
	0x0a, 0x15, 0xb9, 0xc1, 0xc8, 0x94, 0x51, 0xf7, 0x5c, 0xa4, 0x34, 0xc8, 0xdb, 0x04, 0x98, 0x4c,
	0x9d, 0xe0, 0x97, 0x87, 0x0a, 0x0a, 0xad, 0xf3, 0x32, 0x53, 0x4e, 0xee, 0xf2, 0x81, 0x72, 0xa9,
	0x50, 0x5b, 0x0d, 0xee, 0x44, 0xcb, 0x67, 0x19, 0x40, 0x56, 0x88, 0xc3, 0x4d, 0x94, 0x96, 0x35,
	0xd7, 0x16, 0x25, 0x9f, 0xd1, 0x22, 0xde, 0xdf, 0x94, 0x8a, 0x1f, 0xb5, 0x50, 0x1a, 0xbf, 0x42,
	0x1e, 0xdf, 0x08, 0xbe, 0xfd, 0x5e, 0x40, 0x96, 0x97, 0xfe, 0x60, 0x35, 0x58, 0x4f, 0xe3, 0xa7,
	0x6d, 0x13, 0xe0, 0x1b, 0x90, 0xbb, 0x37, 0xa4, 0x07, 0x12, 0x8a, 0x4c, 0x75, 0x69, 0x8a, 0x3f,
	0x43, 0x74, 0xda, 0x53, 0x53, 0x15, 0x94, 0x4a, 0xe0, 0x18, 0x9d, 0x94, 0xec, 0x67, 0xc2, 0xa4,
	0x51, 0xf2, 0x2e, 0xce, 0x42, 0xeb, 0x28, 0xdc, 0x3b, 0x0a, 0xaf, 0x8c, 0xa3, 0xf8, 0xb4, 0x6d,
	0x82, 0x99, 0x7d, 0xc2, 0x12, 0x08, 0x75, 0x4c, 0xfc, 0x09, 0x79, 0x87, 0xab, 0x2b, 0x7f, 0xb8,
	0x1a, 0xad, 0xbd, 0x8b, 0x45, 0xd8, 0xad, 0x1b, 0x5e, 0x75, 0x40, 0xdc, 0x01, 0x7d, 0x93, 0xbd,
	0x71, 0x42, 0xfb, 0x64, 0x7c, 0x8d, 0x66, 0x89, 0xcc, 0xf9, 0xf6, 0x1b, 0xe3, 0x5b, 0x96, 0x09,
	0xe5, 0x8f, 0x8c, 0xda, 0x3f, 0x6d, 0xf9, 0x6d, 0x13, 0x3c, 0xb1, 0xa2, 0x47, 0x3c, 0x42, 0x8f,
	0x75, 0xf0, 0x7b, 0x34, 0xdd, 0x27, 0xe9, 0x8f, 0xcd, 0xaa, 0x8f, 0x9c, 0x43, 0xd7, 0x8d, 0x1f,
	0xb7, 0x4d, 0x30, 0x77, 0xfe, 0x5c, 0x8f, 0xd0, 0x3b, 0x12, 0x7e, 0x87, 0xa6, 0x2e, 0x3b, 0xe5,
	0x4f, 0x8c, 0xa9, 0x99, 0x15, 0xf8, 0x60, 0xbb, 0x7d, 0xfe, 0x7e, 0x90, 0xd0, 0x3b, 0x0e, 0xf9,
	0x3d, 0x40, 0xfe, 0xbd, 0xfb, 0x5f, 0x4b, 0x56, 0x55, 0x42, 0xe2, 0xe7, 0x68, 0xcc, 0x21, 0x15,
	0x26, 0x84, 0x49, 0x3c, 0x6f, 0x9b, 0xc0, 0x73, 0x71, 0x42, 0x2a, 0x08, 0x35, 0x20, 0x7e, 0x8d,
	0xbc, 0xee, 0xfb, 0xf1, 0x57, 0x55, 0xb0, 0xbc, 0xf4, 0x87, 0xab, 0xc1, 0xfa, 0xbf, 0xa3, 0xe8,
	0x0f, 0x20, 0xa1, 0xfd, 0x51, 0xfc, 0x02, 0x4d, 0x84, 0x94, 0x20, 0xfd, 0x91, 0xe1, 0x2c, 0xda,
	0x26, 0xf8, 0xdf, 0x72, 0x4c, 0x9b, 0x50, 0x0b, 0xe3, 0xb7, 0x68, 0x9c, 0x32, 0xcd, 0xdc, 0x81,
	0xce, 0xec, 0x7e, 0xf7, 0x4c, 0xf7, 0xfd, 0x75, 0xe3, 0x84, 0x1a, 0x56, 0x72, 0x62, 0xc2, 0x79,
	0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x98, 0xcd, 0x03, 0x09, 0x04, 0x00, 0x00,
}
