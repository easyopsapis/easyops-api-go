// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: switch.proto

package idcmanager

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
//交换机
type Switch struct {
	//
	//交换机ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//交换机名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//管理IP
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//交换机状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" form:"status"`
	//
	//交换机类型
	Type string `protobuf:"bytes,5,opt,name=type,proto3" json:"type" form:"type"`
	//
	//起始U位
	XStartU int32 `protobuf:"varint,6,opt,name=_startU,json=StartU,proto3" json:"_startU" form:"_startU"`
	//
	//占用U位
	XOccupiedU int32 `protobuf:"varint,7,opt,name=_occupiedU,json=OccupiedU,proto3" json:"_occupiedU" form:"_occupiedU"`
	//
	//备注
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//资产编号
	Propertyid string `protobuf:"bytes,9,opt,name=propertyid,proto3" json:"propertyid" form:"propertyid"`
	//
	//创建时间
	Ctime string `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//创建者
	Creator string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//产品线
	Product string `protobuf:"bytes,12,opt,name=product,proto3" json:"product" form:"product"`
	//
	//客户
	Customer string `protobuf:"bytes,13,opt,name=customer,proto3" json:"customer" form:"customer"`
	//
	//是否单电源
	IsSinglePower        bool     `protobuf:"varint,14,opt,name=isSinglePower,proto3" json:"isSinglePower" form:"isSinglePower"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Switch) Reset()         { *m = Switch{} }
func (m *Switch) String() string { return proto.CompactTextString(m) }
func (*Switch) ProtoMessage()    {}
func (*Switch) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5d7f34d8fcf87fc, []int{0}
}
func (m *Switch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Switch.Unmarshal(m, b)
}
func (m *Switch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Switch.Marshal(b, m, deterministic)
}
func (m *Switch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Switch.Merge(m, src)
}
func (m *Switch) XXX_Size() int {
	return xxx_messageInfo_Switch.Size(m)
}
func (m *Switch) XXX_DiscardUnknown() {
	xxx_messageInfo_Switch.DiscardUnknown(m)
}

var xxx_messageInfo_Switch proto.InternalMessageInfo

func (m *Switch) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Switch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Switch) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Switch) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Switch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Switch) GetXStartU() int32 {
	if m != nil {
		return m.XStartU
	}
	return 0
}

func (m *Switch) GetXOccupiedU() int32 {
	if m != nil {
		return m.XOccupiedU
	}
	return 0
}

func (m *Switch) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Switch) GetPropertyid() string {
	if m != nil {
		return m.Propertyid
	}
	return ""
}

func (m *Switch) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Switch) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Switch) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Switch) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *Switch) GetIsSinglePower() bool {
	if m != nil {
		return m.IsSinglePower
	}
	return false
}

func init() {
	proto.RegisterType((*Switch)(nil), "idcmanager.Switch")
}

func init() { proto.RegisterFile("switch.proto", fileDescriptor_d5d7f34d8fcf87fc) }

var fileDescriptor_d5d7f34d8fcf87fc = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xc1, 0x6e, 0xa3, 0x46,
	0x18, 0xc7, 0xe5, 0x74, 0xe3, 0xc4, 0xb3, 0x49, 0x76, 0x97, 0xb6, 0x12, 0xda, 0x0b, 0x29, 0x9b,
	0x4a, 0x33, 0x89, 0xc7, 0x18, 0x6c, 0x67, 0x15, 0x0e, 0xb5, 0x36, 0xaa, 0x56, 0xda, 0x53, 0x2b,
	0xd2, 0xbd, 0x04, 0x93, 0x88, 0x00, 0xf1, 0xa2, 0x1a, 0x0f, 0x82, 0x71, 0xad, 0x2c, 0xf0, 0x0c,
	0x2b, 0xf5, 0x19, 0x7a, 0xaa, 0xd4, 0x63, 0x5f, 0xa0, 0x6f, 0xd1, 0x8b, 0x2b, 0xf5, 0x11, 0x78,
	0x82, 0x6a, 0x66, 0x48, 0xec, 0x6e, 0x42, 0xdb, 0x43, 0x9a, 0x13, 0xdf, 0x7c, 0xff, 0xff, 0x37,
	0xf3, 0x1b, 0x3e, 0xe9, 0x03, 0xb0, 0x95, 0xce, 0x43, 0xea, 0xbd, 0xeb, 0xc4, 0x09, 0xa1, 0x44,
	0x02, 0xa1, 0xef, 0x45, 0xee, 0xd4, 0x1d, 0x07, 0xc9, 0x73, 0x3c, 0x0e, 0xe9, 0xbb, 0xd9, 0x45,
	0xc7, 0x23, 0x91, 0x36, 0x26, 0x63, 0xa2, 0x71, 0xcb, 0xc5, 0xec, 0x92, 0xaf, 0xf8, 0x82, 0x47,
	0xa2, 0xf4, 0xf9, 0xe1, 0x8a, 0x3d, 0x9a, 0x87, 0xf4, 0x7b, 0x32, 0xd7, 0xc6, 0x04, 0x73, 0x11,
	0xff, 0xe0, 0x4e, 0x42, 0xdf, 0xa5, 0x24, 0x49, 0xb5, 0x9b, 0x50, 0xd4, 0xa9, 0x1f, 0x9e, 0x82,
	0xe6, 0x09, 0x67, 0x90, 0xde, 0x00, 0x10, 0x4e, 0x53, 0xea, 0x4e, 0xbd, 0xe0, 0x8d, 0x2f, 0x37,
	0x76, 0x1b, 0xb0, 0x75, 0x8c, 0xca, 0x85, 0xf2, 0xec, 0x92, 0x24, 0x91, 0xa9, 0x2e, 0x35, 0xf5,
	0xcf, 0x3f, 0x94, 0xa7, 0x60, 0xe7, 0xcc, 0xee, 0xe2, 0x23, 0x17, 0xbf, 0x77, 0x32, 0xbd, 0x57,
	0xec, 0x59, 0x2b, 0xc5, 0xd2, 0x0b, 0xf0, 0x68, 0xea, 0x46, 0x81, 0xbc, 0xc6, 0x37, 0x79, 0x52,
	0x2e, 0x94, 0xc7, 0x62, 0x13, 0x96, 0x55, 0x2d, 0x2e, 0x4a, 0xbf, 0xb6, 0xc0, 0x5a, 0x18, 0xcb,
	0x9f, 0x70, 0xcf, 0x4f, 0xad, 0x72, 0xa1, 0xb4, 0xaa, 0x93, 0x62, 0x76, 0xc2, 0x8f, 0x2d, 0xf0,
	0xa1, 0x75, 0x06, 0xa1, 0x01, 0x07, 0x76, 0x17, 0x0f, 0x9c, 0x4c, 0x2f, 0x72, 0xbb, 0x8b, 0xfb,
	0xce, 0xc8, 0xcf, 0xf4, 0x02, 0xb1, 0x58, 0x77, 0x86, 0x6c, 0xd1, 0x36, 0x0a, 0x04, 0x47, 0x9d,
	0xff, 0xe8, 0x44, 0x59, 0xaf, 0x40, 0xf9, 0x28, 0xdd, 0x87, 0x10, 0xb2, 0x1b, 0xbc, 0xc2, 0xaf,
	0x5d, 0x7c, 0xe9, 0x64, 0x7a, 0xbb, 0x5f, 0x98, 0x28, 0x7b, 0x59, 0xdc, 0xca, 0xe6, 0x26, 0x42,
	0xf9, 0x9d, 0xe6, 0xc3, 0x02, 0x9a, 0xb7, 0xdc, 0x10, 0x1a, 0x82, 0x23, 0x37, 0x2a, 0x8a, 0x5c,
	0x1f, 0xf9, 0x23, 0x3f, 0xb7, 0x75, 0x7c, 0xc4, 0x38, 0x04, 0xec, 0xbf, 0x78, 0x04, 0x66, 0xed,
	0xc9, 0x83, 0x02, 0xc2, 0xdb, 0x67, 0x23, 0x71, 0xc5, 0xdc, 0x7c, 0x10, 0x86, 0x7e, 0x2d, 0x03,
	0x2b, 0xbb, 0x4b, 0x1a, 0xde, 0x27, 0xd8, 0x3f, 0x90, 0xf5, 0x6a, 0xc9, 0xfa, 0x35, 0x64, 0x59,
	0xb7, 0x6d, 0x14, 0x0f, 0x44, 0x67, 0xd4, 0xd2, 0x0d, 0xea, 0xe9, 0x7a, 0x0f, 0x45, 0xa7, 0xd7,
	0xd2, 0x1d, 0xd6, 0xd3, 0xf5, 0xff, 0x0f, 0x3a, 0xb3, 0x0e, 0xe4, 0x65, 0x3d, 0xc8, 0xe0, 0xfe,
	0x41, 0x10, 0xfc, 0xb2, 0x73, 0x80, 0x86, 0xa3, 0x74, 0x7f, 0xcf, 0x5a, 0x0b, 0x63, 0x09, 0x81,
	0x66, 0x4a, 0x5d, 0x3a, 0x4b, 0xe5, 0x47, 0x7c, 0x74, 0x3d, 0x2b, 0x17, 0xca, 0xb6, 0x98, 0x5c,
	0x22, 0xaf, 0x5a, 0x95, 0x81, 0xcd, 0x41, 0x7a, 0x15, 0x07, 0xf2, 0xfa, 0xc7, 0x73, 0x90, 0x65,
	0x55, 0x8b, 0x8b, 0xd2, 0x01, 0xd8, 0x38, 0x4f, 0xa9, 0x9b, 0xd0, 0xb7, 0x72, 0x73, 0xb7, 0x01,
	0xd7, 0x8f, 0xa5, 0x72, 0xa1, 0xec, 0x08, 0x5f, 0x25, 0xa8, 0x56, 0xf3, 0x84, 0x07, 0x52, 0x1f,
	0x80, 0x73, 0xe2, 0x79, 0xb3, 0x38, 0x0c, 0xfc, 0xb7, 0xf2, 0x06, 0xf7, 0x7f, 0xbe, 0x1c, 0xd2,
	0x4b, 0x4d, 0xb5, 0x5a, 0xdf, 0x5c, 0xc7, 0x8c, 0x23, 0x0a, 0x22, 0x22, 0x6f, 0x7e, 0xcc, 0xc1,
	0xb2, 0xaa, 0xc5, 0x45, 0x69, 0x00, 0x40, 0x9c, 0x90, 0x38, 0x48, 0xe8, 0x55, 0xe8, 0xcb, 0x2d,
	0x6e, 0x5d, 0xd9, 0x7a, 0xa9, 0xa9, 0xd6, 0x8a, 0x51, 0xfa, 0xbd, 0x01, 0xd6, 0x3d, 0x1a, 0x46,
	0x81, 0x0c, 0x78, 0xc9, 0x6f, 0x8d, 0x72, 0xa1, 0x6c, 0x89, 0x1a, 0x9e, 0x67, 0xc3, 0xfc, 0x97,
	0x06, 0xf8, 0xb9, 0x71, 0x06, 0xe1, 0xd0, 0xe4, 0x2f, 0x98, 0x75, 0xce, 0xd9, 0x47, 0x43, 0xfe,
	0xcc, 0xfa, 0x05, 0xc2, 0x50, 0xb7, 0xbb, 0xd8, 0x70, 0xf2, 0x2e, 0xd7, 0x11, 0x86, 0x3d, 0x3e,
	0xb5, 0xab, 0x35, 0x6b, 0x8b, 0x21, 0xaa, 0x90, 0xfd, 0xdd, 0xae, 0x03, 0x59, 0xd7, 0x7a, 0x8e,
	0x98, 0xec, 0x22, 0x6d, 0x42, 0xde, 0xcf, 0x6a, 0xb1, 0x8c, 0xe1, 0xa8, 0xc3, 0x9f, 0x07, 0x68,
	0x08, 0x4f, 0x73, 0xfb, 0x00, 0x3b, 0x70, 0x68, 0xde, 0x51, 0xbe, 0x52, 0x3d, 0xdc, 0xb3, 0xc4,
	0x8d, 0xa4, 0x13, 0xb0, 0xe1, 0x25, 0x01, 0xfb, 0x5c, 0xca, 0x8f, 0xf9, 0xe5, 0x8e, 0x96, 0xad,
	0xa9, 0x04, 0x76, 0xbb, 0x17, 0xe0, 0x8b, 0x33, 0xdb, 0xc5, 0xef, 0x5f, 0xe1, 0x53, 0xb6, 0x81,
	0xdd, 0xb9, 0x89, 0xcf, 0xb1, 0x93, 0x19, 0xed, 0x9e, 0x5e, 0xec, 0x59, 0xd7, 0x3b, 0x49, 0x6d,
	0xb0, 0x11, 0x27, 0xc4, 0x9f, 0x79, 0x54, 0xde, 0xe2, 0x9b, 0xae, 0xf4, 0xbb, 0x12, 0x54, 0xeb,
	0xda, 0x22, 0x69, 0x60, 0xd3, 0x9b, 0xa5, 0x94, 0x44, 0x41, 0x22, 0x6f, 0x73, 0xfb, 0xa7, 0xe5,
	0x42, 0x79, 0x52, 0x31, 0x54, 0x8a, 0x6a, 0xdd, 0x98, 0xa4, 0xaf, 0xc0, 0x76, 0x98, 0x9e, 0x84,
	0xd3, 0xf1, 0x24, 0xf8, 0x96, 0xcc, 0x83, 0x44, 0xde, 0xd9, 0x6d, 0xc0, 0xcd, 0x63, 0xb9, 0x5c,
	0x28, 0x9f, 0x55, 0xdf, 0xd7, 0x55, 0x59, 0xb5, 0xfe, 0x6e, 0x3f, 0x7e, 0x7d, 0xfa, 0xf5, 0x98,
	0x74, 0x02, 0x37, 0xbd, 0x22, 0x71, 0xda, 0x99, 0x10, 0xcf, 0x9d, 0x68, 0x1e, 0x99, 0xd2, 0xc4,
	0xf5, 0x68, 0x2a, 0xfe, 0x42, 0x92, 0x20, 0x26, 0x38, 0x22, 0x7e, 0x30, 0x49, 0xb5, 0xca, 0xa8,
	0xf1, 0xa5, 0xb6, 0xfc, 0x81, 0xb9, 0x68, 0x72, 0x6b, 0xef, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x49, 0x69, 0xa7, 0xe3, 0x08, 0x00, 0x00,
}
