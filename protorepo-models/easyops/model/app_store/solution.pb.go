// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: solution.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//解决方案
type Solution struct {
	//
	//实例id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//markdown格式的内容
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" form:"content"`
	//
	//轮播图片
	Thumbnail string `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail" form:"thumbnail"`
	//
	//简介
	Brief string `protobuf:"bytes,5,opt,name=brief,proto3" json:"brief" form:"brief"`
	//
	//图标
	Icon                 *Solution_Icon `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon" form:"icon"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Solution) Reset()         { *m = Solution{} }
func (m *Solution) String() string { return proto.CompactTextString(m) }
func (*Solution) ProtoMessage()    {}
func (*Solution) Descriptor() ([]byte, []int) {
	return fileDescriptor_070ae05a8a5c55b6, []int{0}
}
func (m *Solution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solution.Unmarshal(m, b)
}
func (m *Solution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solution.Marshal(b, m, deterministic)
}
func (m *Solution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solution.Merge(m, src)
}
func (m *Solution) XXX_Size() int {
	return xxx_messageInfo_Solution.Size(m)
}
func (m *Solution) XXX_DiscardUnknown() {
	xxx_messageInfo_Solution.DiscardUnknown(m)
}

var xxx_messageInfo_Solution proto.InternalMessageInfo

func (m *Solution) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Solution) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Solution) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Solution) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *Solution) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

func (m *Solution) GetIcon() *Solution_Icon {
	if m != nil {
		return m.Icon
	}
	return nil
}

type Solution_Icon struct {
	//
	//图标库
	Lib string `protobuf:"bytes,1,opt,name=lib,proto3" json:"lib" form:"lib"`
	//
	//图标url
	Icon                 string   `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon" form:"icon"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Solution_Icon) Reset()         { *m = Solution_Icon{} }
func (m *Solution_Icon) String() string { return proto.CompactTextString(m) }
func (*Solution_Icon) ProtoMessage()    {}
func (*Solution_Icon) Descriptor() ([]byte, []int) {
	return fileDescriptor_070ae05a8a5c55b6, []int{0, 0}
}
func (m *Solution_Icon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solution_Icon.Unmarshal(m, b)
}
func (m *Solution_Icon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solution_Icon.Marshal(b, m, deterministic)
}
func (m *Solution_Icon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solution_Icon.Merge(m, src)
}
func (m *Solution_Icon) XXX_Size() int {
	return xxx_messageInfo_Solution_Icon.Size(m)
}
func (m *Solution_Icon) XXX_DiscardUnknown() {
	xxx_messageInfo_Solution_Icon.DiscardUnknown(m)
}

var xxx_messageInfo_Solution_Icon proto.InternalMessageInfo

func (m *Solution_Icon) GetLib() string {
	if m != nil {
		return m.Lib
	}
	return ""
}

func (m *Solution_Icon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func init() {
	proto.RegisterType((*Solution)(nil), "app_store.Solution")
	proto.RegisterType((*Solution_Icon)(nil), "app_store.Solution.Icon")
}

func init() { proto.RegisterFile("solution.proto", fileDescriptor_070ae05a8a5c55b6) }

var fileDescriptor_070ae05a8a5c55b6 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0x55, 0xda, 0x0d, 0xea, 0xa1, 0x52, 0x2c, 0x2e, 0xa2, 0xde, 0xa4, 0x32, 0x13, 0x4a,
	0x37, 0x12, 0xc3, 0x26, 0x21, 0x51, 0xf1, 0x47, 0x0c, 0x21, 0xd4, 0xdb, 0x70, 0xb7, 0xba, 0x45,
	0x4e, 0x9a, 0x66, 0x16, 0x89, 0x4f, 0x94, 0xb8, 0x4c, 0x8c, 0xf2, 0x22, 0x3c, 0x5c, 0x86, 0xf6,
	0x08, 0x79, 0x02, 0x14, 0x3b, 0xcb, 0x2a, 0xb8, 0xf3, 0xf1, 0xf7, 0x9d, 0x93, 0x93, 0x5f, 0x82,
	0x06, 0x05, 0x24, 0x1b, 0x25, 0x40, 0x7a, 0x59, 0x0e, 0x0a, 0x70, 0x9f, 0x67, 0xd9, 0xd7, 0x42,
	0x41, 0x1e, 0x8d, 0xdc, 0x58, 0xa8, 0x8b, 0x4d, 0xe0, 0x85, 0x90, 0xd2, 0x18, 0x62, 0xa0, 0xda,
	0x08, 0x36, 0x6b, 0x5d, 0xe9, 0x42, 0x9f, 0x4c, 0xe7, 0xe8, 0xd5, 0x8e, 0x9e, 0x5e, 0x0a, 0xf5,
	0x0d, 0x2e, 0x69, 0x0c, 0xae, 0x86, 0xee, 0x77, 0x9e, 0x88, 0x15, 0x57, 0x90, 0x17, 0xb4, 0x3d,
	0x9a, 0x3e, 0xf2, 0xa7, 0x8b, 0x1e, 0x7c, 0x69, 0x96, 0xc0, 0x33, 0x84, 0x84, 0x2c, 0x14, 0x97,
	0x61, 0x34, 0x5b, 0x59, 0x9d, 0x71, 0xc7, 0xe9, 0x9f, 0x4d, 0xaa, 0xd2, 0x7e, 0xbc, 0x86, 0x3c,
	0x9d, 0x92, 0x3b, 0x46, 0x6e, 0xae, 0xed, 0x21, 0x1a, 0x2c, 0xe7, 0x2f, 0xdc, 0xd7, 0xdc, 0xbd,
	0x5a, 0xfc, 0x7c, 0x79, 0xfa, 0xeb, 0xd0, 0xdf, 0x69, 0xc6, 0x4f, 0x51, 0x4f, 0xf2, 0x34, 0xb2,
	0xee, 0xe9, 0x21, 0x8f, 0xaa, 0xd2, 0x3e, 0x30, 0x43, 0xea, 0x5b, 0xe2, 0x6b, 0x88, 0x9f, 0xa3,
	0xfb, 0x21, 0x48, 0x15, 0x49, 0x65, 0x75, 0xb5, 0x87, 0xab, 0xd2, 0x1e, 0x18, 0xaf, 0x01, 0xc4,
	0xbf, 0x55, 0xf0, 0x09, 0xea, 0xab, 0x8b, 0x4d, 0x1a, 0x48, 0x2e, 0x12, 0xab, 0xa7, 0xfd, 0x27,
	0x55, 0x69, 0x0f, 0x8d, 0xdf, 0x22, 0xe2, 0xdf, 0x69, 0xf8, 0x19, 0xda, 0x0b, 0x72, 0x11, 0xad,
	0xad, 0x3d, 0xed, 0x0f, 0xab, 0xd2, 0x7e, 0x68, 0x7c, 0x7d, 0x4d, 0x7c, 0x83, 0xf1, 0x5b, 0xd4,
	0x13, 0x21, 0x48, 0x6b, 0x7f, 0xdc, 0x71, 0x0e, 0x4e, 0x2c, 0xaf, 0xfd, 0x0e, 0xde, 0x6d, 0x38,
	0xde, 0x2c, 0x04, 0xb9, 0xfb, 0x22, 0xb5, 0x4f, 0x7c, 0xdd, 0x36, 0xfa, 0xdd, 0x41, 0xbd, 0x9a,
	0xe3, 0x31, 0xea, 0x26, 0x22, 0x68, 0xa2, 0x1b, 0x54, 0xa5, 0x8d, 0x8c, 0x9c, 0x88, 0x80, 0xf8,
	0x35, 0xc2, 0x59, 0xf3, 0x24, 0x13, 0x0c, 0xfb, 0x67, 0xde, 0xcd, 0xb5, 0x3d, 0x43, 0x9f, 0x97,
	0x8e, 0xc3, 0xe8, 0x7c, 0xc9, 0xe8, 0x94, 0x1d, 0xb1, 0xf7, 0x84, 0xbc, 0x79, 0xc7, 0xb6, 0x2c,
	0x67, 0x72, 0x71, 0x3c, 0x39, 0x9e, 0x6c, 0x1d, 0x46, 0x27, 0xdb, 0x39, 0x77, 0xaf, 0x3e, 0xb8,
	0xe7, 0x8b, 0xa9, 0xc3, 0xd8, 0x7c, 0xc9, 0xd8, 0xff, 0xe6, 0xd1, 0xa1, 0x59, 0xee, 0xec, 0xd3,
	0xf9, 0xc7, 0x18, 0xbc, 0x88, 0x17, 0x3f, 0x20, 0x2b, 0xbc, 0x04, 0x42, 0x9e, 0xd0, 0x3a, 0xd3,
	0x9c, 0x87, 0xaa, 0x30, 0xbf, 0x55, 0x1e, 0x65, 0xe0, 0xa6, 0xb0, 0x8a, 0x92, 0x82, 0x36, 0x22,
	0xd5, 0x25, 0x6d, 0x83, 0x08, 0xf6, 0xb5, 0x79, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x35,
	0xa1, 0x9e, 0xb4, 0x02, 0x00, 0x00,
}
