// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storyboard.proto

package api_gateway

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	micro_app "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
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
//初始化配置
type StoryBoard struct {
	//
	//小产品基本信息
	App *micro_app.InstalledMicroAppBootstrap `protobuf:"bytes,1,opt,name=app,proto3" json:"app" form:"app"`
	//
	//dependsAll
	DependsAll bool `protobuf:"varint,2,opt,name=dependsAll,proto3" json:"dependsAll" form:"dependsAll"`
	//
	//路由配置
	Routes               []*types.Struct `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes" form:"routes"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoryBoard) Reset()         { *m = StoryBoard{} }
func (m *StoryBoard) String() string { return proto.CompactTextString(m) }
func (*StoryBoard) ProtoMessage()    {}
func (*StoryBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_c61fc9bbd5ce8280, []int{0}
}
func (m *StoryBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryBoard.Unmarshal(m, b)
}
func (m *StoryBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryBoard.Marshal(b, m, deterministic)
}
func (m *StoryBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryBoard.Merge(m, src)
}
func (m *StoryBoard) XXX_Size() int {
	return xxx_messageInfo_StoryBoard.Size(m)
}
func (m *StoryBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryBoard.DiscardUnknown(m)
}

var xxx_messageInfo_StoryBoard proto.InternalMessageInfo

func (m *StoryBoard) GetApp() *micro_app.InstalledMicroAppBootstrap {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *StoryBoard) GetDependsAll() bool {
	if m != nil {
		return m.DependsAll
	}
	return false
}

func (m *StoryBoard) GetRoutes() []*types.Struct {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*StoryBoard)(nil), "api_gateway.StoryBoard")
}

func init() { proto.RegisterFile("storyboard.proto", fileDescriptor_c61fc9bbd5ce8280) }

var fileDescriptor_c61fc9bbd5ce8280 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xcf, 0x6a, 0x32, 0x31,
	0x10, 0x67, 0x3f, 0x41, 0x3e, 0x22, 0x2d, 0x75, 0xa1, 0x54, 0xa4, 0xa0, 0x2c, 0x14, 0xbc, 0x98,
	0x80, 0xa5, 0x97, 0xde, 0x0c, 0x14, 0xe9, 0xa1, 0x17, 0xbd, 0xf5, 0xb2, 0x64, 0xb3, 0x31, 0x5d,
	0xc8, 0x3a, 0x43, 0x92, 0xa5, 0xf8, 0x96, 0x7d, 0x02, 0x1f, 0xc2, 0x27, 0x28, 0x9b, 0xb8, 0xea,
	0xb9, 0xb7, 0xfd, 0xcd, 0xfc, 0xfe, 0xcd, 0x86, 0xdc, 0x39, 0x0f, 0x76, 0x5f, 0x80, 0xb0, 0x25,
	0x45, 0x0b, 0x1e, 0xd2, 0x81, 0xc0, 0x2a, 0xd7, 0xc2, 0xab, 0x6f, 0xb1, 0x1f, 0xcf, 0x75, 0xe5,
	0xbf, 0x9a, 0x82, 0x4a, 0xa8, 0x99, 0x06, 0x0d, 0x2c, 0x70, 0x8a, 0x66, 0x1b, 0x50, 0x00, 0xe1,
	0x2b, 0x6a, 0xc7, 0x5a, 0x03, 0x55, 0xc2, 0xed, 0x01, 0x1d, 0x35, 0x20, 0x85, 0x61, 0x12, 0x76,
	0xde, 0x0a, 0xe9, 0x5d, 0x54, 0x5a, 0x85, 0x30, 0xaf, 0xa1, 0x54, 0xc6, 0xb1, 0x13, 0x91, 0x05,
	0xc8, 0xea, 0x4a, 0x5a, 0xc8, 0x05, 0x22, 0xab, 0x76, 0xce, 0x0b, 0x63, 0x54, 0x99, 0x9f, 0x67,
	0x79, 0x01, 0xe0, 0x9d, 0xb7, 0x02, 0x4f, 0x41, 0x8f, 0x1a, 0x40, 0x1b, 0x75, 0xa9, 0xe3, 0xbc,
	0x6d, 0xa4, 0x8f, 0xdb, 0xec, 0x27, 0x21, 0x64, 0xd3, 0xde, 0xc5, 0xdb, 0xbb, 0xd2, 0x15, 0xe9,
	0x09, 0xc4, 0x51, 0x32, 0x4d, 0x66, 0x83, 0xc5, 0x13, 0x3d, 0xbb, 0xd2, 0xf7, 0x2e, 0xe9, 0xa3,
	0x1d, 0x2d, 0x11, 0x79, 0x17, 0xc3, 0x6f, 0x8f, 0x87, 0x09, 0xd9, 0x82, 0xad, 0x5f, 0x33, 0x81,
	0x98, 0xad, 0x5b, 0x87, 0xf4, 0x85, 0x90, 0x52, 0xa1, 0xda, 0x95, 0x6e, 0x69, 0xcc, 0xe8, 0xdf,
	0x34, 0x99, 0xfd, 0xe7, 0xf7, 0xc7, 0xc3, 0x64, 0x18, 0x89, 0x97, 0x5d, 0xb6, 0xbe, 0x22, 0xa6,
	0x9c, 0xf4, 0x2d, 0x34, 0x5e, 0xb9, 0x51, 0x6f, 0xda, 0x9b, 0x0d, 0x16, 0x0f, 0x34, 0xb6, 0xa7,
	0x5d, 0x7b, 0xba, 0x09, 0xed, 0xf9, 0xf0, 0x78, 0x98, 0xdc, 0x44, 0xaf, 0x28, 0xc8, 0xd6, 0x27,
	0x25, 0x5f, 0x7d, 0xbe, 0xfd, 0xfd, 0xdf, 0x5e, 0xbd, 0x68, 0xd1, 0x0f, 0xdc, 0xe7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x31, 0x69, 0xf9, 0xf9, 0x01, 0x00, 0x00,
}