// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fail_instance.proto

package cmdb

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
//失败的实例详情
type FailInstance struct {
	//
	//实例ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//资源模型ID
	XObjectId string `protobuf:"bytes,2,opt,name=_object_id,json=ObjectId,proto3" json:"_object_id" form:"_object_id"`
	//
	//实例名称
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailInstance) Reset()         { *m = FailInstance{} }
func (m *FailInstance) String() string { return proto.CompactTextString(m) }
func (*FailInstance) ProtoMessage()    {}
func (*FailInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f9f4a7e43fa179, []int{0}
}
func (m *FailInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailInstance.Unmarshal(m, b)
}
func (m *FailInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailInstance.Marshal(b, m, deterministic)
}
func (m *FailInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailInstance.Merge(m, src)
}
func (m *FailInstance) XXX_Size() int {
	return xxx_messageInfo_FailInstance.Size(m)
}
func (m *FailInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_FailInstance.DiscardUnknown(m)
}

var xxx_messageInfo_FailInstance proto.InternalMessageInfo

func (m *FailInstance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *FailInstance) GetXObjectId() string {
	if m != nil {
		return m.XObjectId
	}
	return ""
}

func (m *FailInstance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*FailInstance)(nil), "cmdb.FailInstance")
}

func init() { proto.RegisterFile("fail_instance.proto", fileDescriptor_48f9f4a7e43fa179) }

var fileDescriptor_48f9f4a7e43fa179 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xd1, 0x4a, 0x02, 0x41,
	0x14, 0x86, 0xb1, 0x24, 0x6a, 0x8a, 0xca, 0xf5, 0x46, 0xbc, 0x31, 0x36, 0x2f, 0x0a, 0x9c, 0x9d,
	0x4c, 0x08, 0xea, 0x22, 0xca, 0x8b, 0x60, 0xbb, 0x09, 0xbc, 0x54, 0x74, 0x99, 0x9d, 0x19, 0xb7,
	0xa9, 0x99, 0x3d, 0xb2, 0x33, 0x26, 0x29, 0xbe, 0x60, 0x0f, 0x61, 0xd0, 0x23, 0xf8, 0x04, 0xe1,
	0xac, 0xa5, 0x77, 0xe7, 0x3f, 0xff, 0xff, 0x7f, 0x1c, 0x0e, 0x2a, 0x0f, 0xa9, 0x54, 0x91, 0x4c,
	0x8d, 0xa5, 0x29, 0x13, 0xc1, 0x28, 0x03, 0x0b, 0x5e, 0x91, 0x69, 0x1e, 0x57, 0x71, 0x22, 0xed,
	0xeb, 0x38, 0x0e, 0x18, 0x68, 0x92, 0x40, 0x02, 0xc4, 0x99, 0xf1, 0x78, 0xe8, 0x94, 0x13, 0x6e,
	0xca, 0x4b, 0xd5, 0x9b, 0xad, 0xb8, 0x9e, 0x48, 0xfb, 0x0e, 0x13, 0x92, 0x00, 0x76, 0x26, 0xfe,
	0xa0, 0x4a, 0x72, 0x6a, 0x21, 0x33, 0xe4, 0x7f, 0xcc, 0x7b, 0xfe, 0x57, 0x01, 0x1d, 0x3d, 0x51,
	0xa9, 0xc2, 0xf5, 0x0d, 0x5e, 0x88, 0xd0, 0xdf, 0x3d, 0x21, 0xaf, 0x14, 0xce, 0x0a, 0x17, 0x07,
	0xed, 0xcb, 0xe5, 0xa2, 0x56, 0x1a, 0x42, 0xa6, 0xef, 0xfc, 0x8d, 0xe7, 0xff, 0x7c, 0xd7, 0x4e,
	0xd1, 0xf1, 0xa0, 0x77, 0x85, 0x6f, 0x29, 0x9e, 0xf6, 0x67, 0xcd, 0xd6, 0xbc, 0xde, 0xd9, 0x2a,
	0x7b, 0xcf, 0x08, 0x45, 0x10, 0xbf, 0x09, 0x66, 0x23, 0xc9, 0x2b, 0x3b, 0x0e, 0xd5, 0xd8, 0xa0,
	0x36, 0xde, 0x0a, 0x55, 0x46, 0xa5, 0x41, 0x8f, 0xe2, 0xe9, 0x23, 0xee, 0x46, 0xfd, 0x59, 0xb3,
	0xd1, 0xba, 0x9e, 0xd7, 0x3b, 0xfb, 0x2f, 0x2e, 0x12, 0x72, 0xef, 0x1c, 0x15, 0x53, 0xaa, 0x45,
	0x65, 0xd7, 0x51, 0x4e, 0x96, 0x8b, 0xda, 0x61, 0x4e, 0x59, 0x6d, 0xfd, 0x8e, 0x33, 0xdb, 0x0f,
	0xdd, 0xfb, 0x04, 0x02, 0x41, 0xcd, 0x27, 0x8c, 0x4c, 0xa0, 0x80, 0x51, 0x45, 0x18, 0xa4, 0x36,
	0xa3, 0xcc, 0x9a, 0xfc, 0x81, 0x99, 0x18, 0x01, 0xd6, 0xc0, 0x85, 0x32, 0x64, 0x1d, 0x24, 0x4e,
	0x92, 0xd5, 0xd7, 0xe3, 0x3d, 0x17, 0x6a, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xef, 0x17, 0x8f,
	0xc2, 0x99, 0x01, 0x00, 0x00,
}
