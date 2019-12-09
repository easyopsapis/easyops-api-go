// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: micro_app_container.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//小产品桌面
type MicroAppContainer struct {
	//
	//cmdb唯一标识, 业务逻辑请使用id作为唯一标识
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//小产品桌面名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//小产品桌面id
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id" form:"id"`
	//
	//桌面类型
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type" form:"type"`
	//
	//桌面显示顺序
	Order                int32    `protobuf:"varint,5,opt,name=order,proto3" json:"order" form:"order"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicroAppContainer) Reset()         { *m = MicroAppContainer{} }
func (m *MicroAppContainer) String() string { return proto.CompactTextString(m) }
func (*MicroAppContainer) ProtoMessage()    {}
func (*MicroAppContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8ca2cd86606eeb, []int{0}
}
func (m *MicroAppContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicroAppContainer.Unmarshal(m, b)
}
func (m *MicroAppContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicroAppContainer.Marshal(b, m, deterministic)
}
func (m *MicroAppContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroAppContainer.Merge(m, src)
}
func (m *MicroAppContainer) XXX_Size() int {
	return xxx_messageInfo_MicroAppContainer.Size(m)
}
func (m *MicroAppContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroAppContainer.DiscardUnknown(m)
}

var xxx_messageInfo_MicroAppContainer proto.InternalMessageInfo

func (m *MicroAppContainer) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *MicroAppContainer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MicroAppContainer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MicroAppContainer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MicroAppContainer) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func init() {
	proto.RegisterType((*MicroAppContainer)(nil), "micro_app.MicroAppContainer")
}

func init() { proto.RegisterFile("micro_app_container.proto", fileDescriptor_ee8ca2cd86606eeb) }

var fileDescriptor_ee8ca2cd86606eeb = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x49, 0x5f, 0xfb, 0xa0, 0xfb, 0x9e, 0xda, 0xe6, 0x14, 0x0b, 0x92, 0xb2, 0x8a, 0xd4,
	0x43, 0xb2, 0x4a, 0x41, 0xd0, 0x9b, 0x2d, 0x1e, 0x7a, 0xf0, 0x92, 0xa3, 0xa2, 0x65, 0x9b, 0x6c,
	0xe3, 0x62, 0x92, 0x59, 0x36, 0x5b, 0x4b, 0x15, 0xbf, 0x6a, 0x04, 0xbf, 0x80, 0x90, 0x4f, 0x20,
	0x99, 0xd4, 0x5a, 0xbc, 0xcd, 0xcc, 0xff, 0xf7, 0xff, 0xcf, 0xb2, 0x43, 0xf6, 0x53, 0x19, 0x6a,
	0x98, 0x72, 0xa5, 0xa6, 0x21, 0x64, 0x86, 0xcb, 0x4c, 0x68, 0x5f, 0x69, 0x30, 0x60, 0xb7, 0x37,
	0x52, 0xcf, 0x8b, 0xa5, 0x79, 0x5c, 0xcc, 0xfc, 0x10, 0x52, 0x16, 0x43, 0x0c, 0x0c, 0x89, 0xd9,
	0x62, 0x8e, 0x1d, 0x36, 0x58, 0xd5, 0xce, 0xde, 0xf9, 0x16, 0x9e, 0x2e, 0xa5, 0x79, 0x82, 0x25,
	0x8b, 0xc1, 0x43, 0xd1, 0x7b, 0xe6, 0x89, 0x8c, 0xb8, 0x01, 0x9d, 0xb3, 0x4d, 0x59, 0xfb, 0xe8,
	0xa7, 0x45, 0xba, 0x37, 0xd5, 0xd2, 0x2b, 0xa5, 0xc6, 0xdf, 0xaf, 0xb1, 0x27, 0x84, 0xc8, 0x2c,
	0x37, 0x3c, 0x0b, 0xc5, 0x24, 0x72, 0xac, 0xbe, 0x35, 0x68, 0x8f, 0x4e, 0xca, 0xc2, 0xed, 0xce,
	0x41, 0xa7, 0x97, 0xf4, 0x47, 0xa3, 0x1f, 0xef, 0x6e, 0x87, 0xec, 0x3e, 0xdc, 0x9d, 0x7a, 0x17,
	0xdc, 0x7b, 0xb9, 0x7f, 0x3d, 0x1b, 0xbe, 0x1d, 0x05, 0x5b, 0x66, 0xfb, 0x90, 0x34, 0x33, 0x9e,
	0x0a, 0xa7, 0x81, 0x21, 0x7b, 0x65, 0xe1, 0xfe, 0xab, 0x43, 0xaa, 0x29, 0x0d, 0x50, 0xb4, 0x0f,
	0x48, 0x43, 0x46, 0xce, 0x1f, 0x44, 0x76, 0xca, 0xc2, 0x6d, 0xaf, 0xf7, 0x44, 0x34, 0x68, 0x48,
	0xcc, 0x30, 0x2b, 0x25, 0x9c, 0xe6, 0xef, 0x8c, 0x6a, 0x4a, 0x03, 0x14, 0xed, 0x63, 0xd2, 0x02,
	0x1d, 0x09, 0xed, 0xb4, 0xfa, 0xd6, 0xa0, 0x35, 0xea, 0x94, 0x85, 0xfb, 0xbf, 0xa6, 0x70, 0x4c,
	0x83, 0x5a, 0x1e, 0x5d, 0xdf, 0x8e, 0x63, 0xf0, 0x05, 0xcf, 0x57, 0xa0, 0x72, 0x3f, 0x81, 0x90,
	0x27, 0xac, 0xba, 0x83, 0xe6, 0xa1, 0xc9, 0xeb, 0x5f, 0xd6, 0x42, 0x81, 0x97, 0x42, 0x24, 0x92,
	0x9c, 0xad, 0x41, 0x86, 0x2d, 0xdb, 0xdc, 0x67, 0xf6, 0x17, 0xc9, 0xe1, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x22, 0xb5, 0x2a, 0x4b, 0xce, 0x01, 0x00, 0x00,
}