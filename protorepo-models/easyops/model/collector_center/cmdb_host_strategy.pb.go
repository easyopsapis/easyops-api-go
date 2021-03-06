// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmdb_host_strategy.proto

package collector_center

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//cmdb主机查询策略
type CmdbHostStrategy struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//查询策略ID
	QueryId string `protobuf:"bytes,2,opt,name=queryId,proto3" json:"queryId" form:"queryId"`
	//
	//查询者
	Querier string `protobuf:"bytes,3,opt,name=querier,proto3" json:"querier" form:"querier"`
	//
	//查询策略字符串
	QueryString          string   `protobuf:"bytes,4,opt,name=queryString,proto3" json:"queryString" form:"queryString"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdbHostStrategy) Reset()         { *m = CmdbHostStrategy{} }
func (m *CmdbHostStrategy) String() string { return proto.CompactTextString(m) }
func (*CmdbHostStrategy) ProtoMessage()    {}
func (*CmdbHostStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b2b28207ebc29ef, []int{0}
}
func (m *CmdbHostStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdbHostStrategy.Unmarshal(m, b)
}
func (m *CmdbHostStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdbHostStrategy.Marshal(b, m, deterministic)
}
func (m *CmdbHostStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdbHostStrategy.Merge(m, src)
}
func (m *CmdbHostStrategy) XXX_Size() int {
	return xxx_messageInfo_CmdbHostStrategy.Size(m)
}
func (m *CmdbHostStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdbHostStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_CmdbHostStrategy proto.InternalMessageInfo

func (m *CmdbHostStrategy) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *CmdbHostStrategy) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *CmdbHostStrategy) GetQuerier() string {
	if m != nil {
		return m.Querier
	}
	return ""
}

func (m *CmdbHostStrategy) GetQueryString() string {
	if m != nil {
		return m.QueryString
	}
	return ""
}

func init() {
	proto.RegisterType((*CmdbHostStrategy)(nil), "collector_center.CmdbHostStrategy")
}

func init() { proto.RegisterFile("cmdb_host_strategy.proto", fileDescriptor_3b2b28207ebc29ef) }

var fileDescriptor_3b2b28207ebc29ef = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xbf, 0x4e, 0x84, 0x40,
	0x10, 0x06, 0xf0, 0xa0, 0xc6, 0x3f, 0x98, 0xe8, 0x65, 0x4d, 0x0c, 0xb1, 0xc1, 0x50, 0x59, 0x78,
	0x6c, 0x61, 0x63, 0x2c, 0xcf, 0xc6, 0x33, 0x56, 0x5c, 0x67, 0x43, 0xd8, 0x65, 0x6e, 0x0f, 0x03,
	0x0c, 0xce, 0x0e, 0x05, 0x8f, 0x6a, 0xc3, 0x43, 0xf0, 0x04, 0x26, 0x0b, 0x98, 0x8b, 0xb9, 0x6e,
	0x27, 0xdf, 0xef, 0xfb, 0x8a, 0xf5, 0x03, 0x5d, 0xe5, 0x2a, 0xdd, 0xa1, 0xe5, 0xd4, 0x32, 0x65,
	0x0c, 0xa6, 0x8b, 0x1b, 0x42, 0x46, 0xb1, 0xd0, 0x58, 0x96, 0xa0, 0x19, 0x29, 0xd5, 0x50, 0x33,
	0xd0, 0xdd, 0xd2, 0x14, 0xbc, 0x6b, 0x55, 0xac, 0xb1, 0x92, 0x06, 0x0d, 0x4a, 0x07, 0x55, 0xbb,
	0x75, 0x97, 0x3b, 0xdc, 0x6b, 0x1c, 0x88, 0x7e, 0x3c, 0x7f, 0xf1, 0x5a, 0xe5, 0xea, 0x0d, 0x2d,
	0x6f, 0xa6, 0x6d, 0x21, 0xfd, 0x73, 0x54, 0x5f, 0xa0, 0x79, 0x9d, 0x07, 0xde, 0xbd, 0xf7, 0x70,
	0xb1, 0xba, 0x19, 0xfa, 0xf0, 0x7a, 0x8b, 0x54, 0xbd, 0x44, 0x73, 0x12, 0x25, 0x7f, 0x48, 0x3c,
	0xfa, 0x67, 0xdf, 0x2d, 0x50, 0xb7, 0xce, 0x83, 0x23, 0xe7, 0xc5, 0xd0, 0x87, 0x57, 0xa3, 0x9f,
	0x82, 0x28, 0x99, 0xc9, 0xac, 0x0b, 0xa0, 0xe0, 0xf8, 0x90, 0x2e, 0x80, 0x26, 0x5d, 0x00, 0x89,
	0x67, 0xff, 0xd2, 0x15, 0x37, 0x4c, 0x45, 0x6d, 0x82, 0x13, 0xd7, 0xb8, 0x1d, 0xfa, 0x50, 0xec,
	0xed, 0x8f, 0x61, 0x94, 0xec, 0xd3, 0xd5, 0xc7, 0xe7, 0xbb, 0xc1, 0x18, 0x32, 0xdb, 0x61, 0x63,
	0xe3, 0x12, 0x75, 0x56, 0x4a, 0x8d, 0x35, 0x53, 0xa6, 0xd9, 0x8e, 0xff, 0x42, 0xd0, 0xe0, 0xb2,
	0xc2, 0x1c, 0x4a, 0x2b, 0x27, 0x28, 0xdd, 0x29, 0xff, 0x7f, 0xac, 0x3a, 0x75, 0x85, 0xa7, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xc3, 0x6a, 0x74, 0x8d, 0x01, 0x00, 0x00,
}
