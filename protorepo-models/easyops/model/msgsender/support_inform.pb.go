// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: support_inform.proto

package msgsender

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//支持的通知方式详情
type SupportInform struct {
	//
	//关联的cmdb用户对象字段
	ColOfCmdbUserObject string `protobuf:"bytes,1,opt,name=col_of_cmdb_user_object,json=colOfCmdbUserObject,proto3" json:"col_of_cmdb_user_object" form:"col_of_cmdb_user_object"`
	//
	//是否启用
	Enable string `protobuf:"bytes,2,opt,name=enable,proto3" json:"enable" form:"enable"`
	//
	//通知插件名称
	PluginName string `protobuf:"bytes,3,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name" form:"plugin_name"`
	//
	//通知方式描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description"`
	//
	//通知方式名称
	InformType           string   `protobuf:"bytes,5,opt,name=inform_type,json=informType,proto3" json:"inform_type" form:"inform_type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportInform) Reset()         { *m = SupportInform{} }
func (m *SupportInform) String() string { return proto.CompactTextString(m) }
func (*SupportInform) ProtoMessage()    {}
func (*SupportInform) Descriptor() ([]byte, []int) {
	return fileDescriptor_8608723efc9942d1, []int{0}
}
func (m *SupportInform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportInform.Unmarshal(m, b)
}
func (m *SupportInform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportInform.Marshal(b, m, deterministic)
}
func (m *SupportInform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportInform.Merge(m, src)
}
func (m *SupportInform) XXX_Size() int {
	return xxx_messageInfo_SupportInform.Size(m)
}
func (m *SupportInform) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportInform.DiscardUnknown(m)
}

var xxx_messageInfo_SupportInform proto.InternalMessageInfo

func (m *SupportInform) GetColOfCmdbUserObject() string {
	if m != nil {
		return m.ColOfCmdbUserObject
	}
	return ""
}

func (m *SupportInform) GetEnable() string {
	if m != nil {
		return m.Enable
	}
	return ""
}

func (m *SupportInform) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SupportInform) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SupportInform) GetInformType() string {
	if m != nil {
		return m.InformType
	}
	return ""
}

func init() {
	proto.RegisterType((*SupportInform)(nil), "msgsender.SupportInform")
}

func init() { proto.RegisterFile("support_inform.proto", fileDescriptor_8608723efc9942d1) }

var fileDescriptor_8608723efc9942d1 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x69, 0xff, 0xdf, 0x42, 0x53, 0xba, 0x70, 0x14, 0x1d, 0x5c, 0x58, 0xc9, 0x4a, 0x17,
	0x9d, 0x59, 0xb8, 0x50, 0x5c, 0xb6, 0xb8, 0x70, 0x63, 0x61, 0x54, 0x10, 0x37, 0x61, 0x92, 0xb9,
	0x33, 0x8e, 0x24, 0xb9, 0x21, 0xc9, 0x2c, 0xfa, 0x76, 0x3e, 0x49, 0x1f, 0xa2, 0x4f, 0x20, 0x4d,
	0x8a, 0x0c, 0x05, 0x77, 0x39, 0xf7, 0x7c, 0xe7, 0x04, 0xee, 0x25, 0xa7, 0xae, 0x33, 0x06, 0xad,
	0x67, 0xad, 0xae, 0xd1, 0xaa, 0xcc, 0x58, 0xf4, 0x98, 0x8c, 0x95, 0x6b, 0x1c, 0xe8, 0x0a, 0xec,
	0xc5, 0xbc, 0x69, 0xfd, 0x67, 0xc7, 0x33, 0x81, 0x2a, 0x6f, 0xb0, 0xc1, 0x3c, 0x10, 0xbc, 0xab,
	0x83, 0x0a, 0x22, 0xbc, 0x62, 0x92, 0x7e, 0x0f, 0xc9, 0xf4, 0x25, 0x56, 0x3e, 0x85, 0xc6, 0xe4,
	0x9d, 0x9c, 0x0b, 0x94, 0x0c, 0x6b, 0x26, 0x54, 0xc5, 0x59, 0xe7, 0xc0, 0x32, 0xe4, 0x5f, 0x20,
	0x7c, 0x3a, 0xb8, 0x1a, 0x5c, 0x8f, 0x17, 0x74, 0xbb, 0x99, 0x5d, 0xee, 0xd0, 0x07, 0xfa, 0x07,
	0x48, 0x8b, 0x13, 0x81, 0x72, 0x55, 0x2f, 0x55, 0xc5, 0xdf, 0x1c, 0xd8, 0x55, 0x98, 0x26, 0x37,
	0x64, 0x04, 0xba, 0xe4, 0x12, 0xd2, 0x61, 0x28, 0x3a, 0xde, 0x6e, 0x66, 0xd3, 0x58, 0x14, 0xe7,
	0xb4, 0xd8, 0x03, 0xc9, 0x1d, 0x99, 0x18, 0xd9, 0x35, 0xad, 0x66, 0xba, 0x54, 0x90, 0xfe, 0x0b,
	0xfc, 0xd9, 0x76, 0x33, 0x4b, 0x22, 0xdf, 0x33, 0x69, 0x41, 0xa2, 0x7a, 0x2e, 0x15, 0x24, 0xf7,
	0x64, 0x52, 0x81, 0x13, 0xb6, 0x35, 0xbe, 0x45, 0x9d, 0xfe, 0x3f, 0x0c, 0xf6, 0x4c, 0x5a, 0xf4,
	0xd1, 0xdd, 0x97, 0x71, 0xa7, 0xcc, 0xaf, 0x0d, 0xa4, 0x47, 0x87, 0xc9, 0x9e, 0x49, 0x0b, 0x12,
	0xd5, 0xeb, 0xda, 0xc0, 0xe2, 0xf1, 0x63, 0xd9, 0x60, 0x06, 0xa5, 0x5b, 0xa3, 0x71, 0x99, 0x44,
	0x51, 0xca, 0x5c, 0xa0, 0xf6, 0xb6, 0x14, 0xde, 0xc5, 0xf5, 0x5b, 0x30, 0x38, 0x57, 0x58, 0x81,
	0x74, 0xf9, 0x1e, 0xcc, 0x83, 0xcc, 0x7f, 0x0f, 0xc7, 0x47, 0x81, 0xbc, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0xb4, 0xef, 0xa1, 0xe2, 0x01, 0x00, 0x00,
}
