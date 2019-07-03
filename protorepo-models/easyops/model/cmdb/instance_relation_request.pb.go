// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance_relation_request.proto

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
//InstanceRelationRequest
type InstanceRelationRequest struct {
	//
	//模型的实例ID列表
	InstanceIds []string `protobuf:"bytes,1,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids" form:"instance_ids"`
	//
	//关联的实例ID列表
	RelatedInstanceIds   []string `protobuf:"bytes,2,rep,name=related_instance_ids,json=relatedInstanceIds,proto3" json:"related_instance_ids" form:"related_instance_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceRelationRequest) Reset()         { *m = InstanceRelationRequest{} }
func (m *InstanceRelationRequest) String() string { return proto.CompactTextString(m) }
func (*InstanceRelationRequest) ProtoMessage()    {}
func (*InstanceRelationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9cba36a62ea33d9, []int{0}
}
func (m *InstanceRelationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceRelationRequest.Unmarshal(m, b)
}
func (m *InstanceRelationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceRelationRequest.Marshal(b, m, deterministic)
}
func (m *InstanceRelationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceRelationRequest.Merge(m, src)
}
func (m *InstanceRelationRequest) XXX_Size() int {
	return xxx_messageInfo_InstanceRelationRequest.Size(m)
}
func (m *InstanceRelationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceRelationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceRelationRequest proto.InternalMessageInfo

func (m *InstanceRelationRequest) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

func (m *InstanceRelationRequest) GetRelatedInstanceIds() []string {
	if m != nil {
		return m.RelatedInstanceIds
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceRelationRequest)(nil), "cmdb.InstanceRelationRequest")
}

func init() { proto.RegisterFile("instance_relation_request.proto", fileDescriptor_f9cba36a62ea33d9) }

var fileDescriptor_f9cba36a62ea33d9 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x19, 0x15, 0xc1, 0x2a, 0x22, 0x55, 0x50, 0xc6, 0x45, 0x65, 0x70, 0x21, 0x48, 0x1a,
	0xa5, 0x20, 0xe8, 0x42, 0x64, 0x76, 0xdd, 0xb8, 0xe8, 0x52, 0xd1, 0x92, 0x26, 0x99, 0x18, 0x4c,
	0xfa, 0x6a, 0x92, 0x3a, 0xa8, 0xf8, 0x8f, 0x7e, 0xc1, 0x08, 0x7e, 0xc2, 0x7c, 0x81, 0x98, 0xd4,
	0x71, 0x84, 0xd9, 0xbd, 0xcb, 0x3d, 0xef, 0xe6, 0xe6, 0x45, 0x89, 0xac, 0xad, 0x23, 0x35, 0xe5,
	0xa5, 0xe1, 0x8a, 0x38, 0x09, 0x75, 0x69, 0xf8, 0x53, 0xcb, 0xad, 0x4b, 0x1b, 0x03, 0x0e, 0xe2,
	0x15, 0xaa, 0x59, 0xd5, 0x47, 0x42, 0xba, 0x87, 0xb6, 0x4a, 0x29, 0x68, 0x2c, 0x40, 0x00, 0xf6,
	0x66, 0xd5, 0x8e, 0xbc, 0xf2, 0xc2, 0x4f, 0x61, 0xa9, 0x7f, 0x36, 0x87, 0xeb, 0xb1, 0x74, 0x8f,
	0x30, 0xc6, 0x02, 0x90, 0x37, 0xd1, 0x33, 0x51, 0x92, 0x11, 0x07, 0xc6, 0xe2, 0xd9, 0x18, 0xf6,
	0x06, 0x1f, 0xbd, 0x68, 0x37, 0xef, 0x0a, 0x15, 0x5d, 0x9f, 0x22, 0xd4, 0x89, 0xaf, 0xa3, 0x8d,
	0x59, 0x57, 0xc9, 0xec, 0x5e, 0xef, 0x60, 0xf9, 0x68, 0x6d, 0x78, 0x3c, 0x9d, 0x24, 0xdb, 0x23,
	0x30, 0xfa, 0x62, 0x30, 0xef, 0x0e, 0xbe, 0x3e, 0x93, 0xad, 0x68, 0xf3, 0xfe, 0xf6, 0x04, 0x9d,
	0x13, 0xf4, 0x7a, 0xf7, 0x76, 0x9a, 0xbd, 0x1f, 0x16, 0xeb, 0xbf, 0x48, 0xce, 0x6c, 0xcc, 0xa3,
	0x1d, 0xff, 0x65, 0xce, 0xca, 0x7f, 0xb9, 0x4b, 0x3e, 0x37, 0x9b, 0x4e, 0x92, 0xfd, 0x90, 0xbb,
	0x88, 0x5a, 0x9c, 0x1f, 0x77, 0x68, 0xfe, 0xf7, 0xcc, 0xf0, 0xea, 0xe6, 0x52, 0x40, 0xca, 0x89,
	0x7d, 0x81, 0xc6, 0xa6, 0x0a, 0x28, 0x51, 0x98, 0x42, 0xed, 0x0c, 0xa1, 0xce, 0x86, 0x33, 0x1a,
	0xde, 0x00, 0xd2, 0xc0, 0xb8, 0xb2, 0xb8, 0x03, 0xb1, 0x97, 0xf8, 0xe7, 0xf6, 0xd5, 0xaa, 0x87,
	0xb2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x02, 0x7c, 0x62, 0xab, 0x01, 0x00, 0x00,
}