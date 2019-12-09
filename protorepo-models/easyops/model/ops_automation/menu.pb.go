// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: menu.proto

package ops_automation

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
//创建巡检菜单
type Menu struct {
	//
	//菜单名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//菜单图标
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//菜单分类
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category" form:"category"`
	//
	//访问白名单 [easyops, test_uesr]
	Visitors []string `protobuf:"bytes,4,rep,name=visitors,proto3" json:"visitors" form:"visitors"`
	//
	//管理白名单 [easyops, test_uesr]
	Managers []string `protobuf:"bytes,5,rep,name=managers,proto3" json:"managers" form:"managers"`
	//
	//菜单Id
	Id                   string   `protobuf:"bytes,6,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Menu) Reset()         { *m = Menu{} }
func (m *Menu) String() string { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()    {}
func (*Menu) Descriptor() ([]byte, []int) {
	return fileDescriptor_0357f39af0c26546, []int{0}
}
func (m *Menu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menu.Unmarshal(m, b)
}
func (m *Menu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menu.Marshal(b, m, deterministic)
}
func (m *Menu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menu.Merge(m, src)
}
func (m *Menu) XXX_Size() int {
	return xxx_messageInfo_Menu.Size(m)
}
func (m *Menu) XXX_DiscardUnknown() {
	xxx_messageInfo_Menu.DiscardUnknown(m)
}

var xxx_messageInfo_Menu proto.InternalMessageInfo

func (m *Menu) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Menu) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Menu) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Menu) GetVisitors() []string {
	if m != nil {
		return m.Visitors
	}
	return nil
}

func (m *Menu) GetManagers() []string {
	if m != nil {
		return m.Managers
	}
	return nil
}

func (m *Menu) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Menu)(nil), "ops_automation.Menu")
}

func init() { proto.RegisterFile("menu.proto", fileDescriptor_0357f39af0c26546) }

var fileDescriptor_0357f39af0c26546 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0xf1, 0x4f, 0x45, 0x47, 0xa8, 0x90, 0x53, 0x10, 0x4a, 0x6c, 0xec, 0xc1, 0x43, 0x27,
	0x93, 0x56, 0x28, 0xe8, 0xa1, 0xb4, 0x9e, 0x4a, 0xa1, 0x97, 0x1c, 0x7a, 0x50, 0x5c, 0x19, 0x93,
	0x31, 0x3b, 0x6c, 0x26, 0x6f, 0xc8, 0x4c, 0x14, 0x57, 0xf6, 0xb2, 0x1f, 0xd4, 0x85, 0xfd, 0x08,
	0x7e, 0x81, 0x5d, 0x32, 0x31, 0xd9, 0x5d, 0x4f, 0x9b, 0xd3, 0xf3, 0xbe, 0xcf, 0xef, 0x79, 0x48,
	0x32, 0x83, 0x90, 0x60, 0x71, 0xe6, 0x24, 0x29, 0x28, 0x30, 0x3e, 0x42, 0x22, 0x57, 0x34, 0x53,
	0x20, 0xa8, 0xe2, 0x10, 0xf7, 0x71, 0xc8, 0xd5, 0x75, 0xb6, 0x76, 0x7c, 0x10, 0x24, 0x84, 0x10,
	0x88, 0xc6, 0xd6, 0xd9, 0x46, 0x4f, 0x7a, 0xd0, 0xaa, 0x88, 0xf7, 0x7f, 0xbc, 0xc2, 0xc5, 0x8e,
	0xab, 0x1b, 0xd8, 0x91, 0x10, 0xb0, 0x36, 0xf1, 0x96, 0x46, 0x3c, 0xa0, 0x0a, 0x52, 0x49, 0x2a,
	0x59, 0xe4, 0xec, 0xfb, 0x06, 0x6a, 0xfe, 0x63, 0x71, 0x66, 0x0c, 0x51, 0x33, 0xa6, 0x82, 0x99,
	0xb5, 0x41, 0x6d, 0xd4, 0x99, 0xf5, 0x4e, 0x47, 0xab, 0xbb, 0x81, 0x54, 0x4c, 0xed, 0x7c, 0x6b,
	0x7b, 0xda, 0xcc, 0x21, 0xee, 0x43, 0x6c, 0xd6, 0x2f, 0xa1, 0x7c, 0x6b, 0x7b, 0xda, 0x34, 0x7e,
	0xa2, 0xb6, 0x4f, 0x15, 0x0b, 0x21, 0xdd, 0x9b, 0x0d, 0x0d, 0xda, 0xa7, 0xa3, 0xd5, 0x2b, 0xc0,
	0xd2, 0xb1, 0x1f, 0x1f, 0xac, 0x6e, 0xf2, 0x54, 0x3e, 0x35, 0xaf, 0xca, 0x18, 0xff, 0x51, 0x7b,
	0xcb, 0x25, 0xcf, 0x5f, 0xd7, 0x6c, 0x0e, 0x1a, 0xa3, 0xce, 0x6c, 0xfa, 0x92, 0x2f, 0x9d, 0x3c,
	0x3f, 0x44, 0x9f, 0xaf, 0x16, 0x14, 0xdf, 0xfe, 0xc6, 0x73, 0x17, 0x4f, 0x96, 0x0b, 0xa7, 0xd2,
	0x2b, 0xbc, 0x3c, 0x7c, 0xff, 0x3a, 0xfe, 0x76, 0xf7, 0xc5, 0xab, 0xba, 0xf2, 0x5e, 0x41, 0x63,
	0x1a, 0xb2, 0x54, 0x9a, 0x1f, 0x2e, 0x7b, 0x4b, 0xe7, 0xfd, 0xbd, 0x65, 0xc2, 0xf8, 0x85, 0xea,
	0x3c, 0x30, 0x5b, 0xfa, 0x4b, 0xdd, 0xd3, 0xd1, 0xea, 0x9c, 0x7f, 0x49, 0x90, 0x77, 0x59, 0xe8,
	0x53, 0xd9, 0xb5, 0x5a, 0x2e, 0x5c, 0x3c, 0x29, 0xf5, 0xc1, 0x2d, 0x7a, 0xea, 0x3c, 0x98, 0xfd,
	0x9d, 0xff, 0x09, 0xc1, 0x61, 0x54, 0xee, 0x21, 0x91, 0x4e, 0x04, 0x3e, 0x8d, 0x88, 0x0f, 0xb1,
	0x4a, 0xa9, 0xaf, 0x64, 0x71, 0xf0, 0x29, 0x4b, 0x00, 0x0b, 0x08, 0x58, 0x24, 0xc9, 0x19, 0x24,
	0x7a, 0x24, 0x6f, 0xef, 0xcd, 0xba, 0xa5, 0xf1, 0xf1, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18,
	0x08, 0xdf, 0x84, 0x5c, 0x02, 0x00, 0x00,
}
