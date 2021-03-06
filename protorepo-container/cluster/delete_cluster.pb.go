// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_cluster.proto

package cluster

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//DeleteCluster请求
type DeleteClusterRequest struct {
	//
	//集群Id
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bcfb8ee9109f429, []int{0}
}
func (m *DeleteClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterRequest.Unmarshal(m, b)
}
func (m *DeleteClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterRequest.Marshal(b, m, deterministic)
}
func (m *DeleteClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterRequest.Merge(m, src)
}
func (m *DeleteClusterRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterRequest.Size(m)
}
func (m *DeleteClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterRequest proto.InternalMessageInfo

func (m *DeleteClusterRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//DeleteClusterApi返回
type DeleteClusterResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteClusterResponseWrapper) Reset()         { *m = DeleteClusterResponseWrapper{} }
func (m *DeleteClusterResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterResponseWrapper) ProtoMessage()    {}
func (*DeleteClusterResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bcfb8ee9109f429, []int{1}
}
func (m *DeleteClusterResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteClusterResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteClusterResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterResponseWrapper.Merge(m, src)
}
func (m *DeleteClusterResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterResponseWrapper.Size(m)
}
func (m *DeleteClusterResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterResponseWrapper proto.InternalMessageInfo

func (m *DeleteClusterResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteClusterResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteClusterResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteClusterResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteClusterRequest)(nil), "cluster.DeleteClusterRequest")
	proto.RegisterType((*DeleteClusterResponseWrapper)(nil), "cluster.DeleteClusterResponseWrapper")
}

func init() { proto.RegisterFile("delete_cluster.proto", fileDescriptor_2bcfb8ee9109f429) }

var fileDescriptor_2bcfb8ee9109f429 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xc9, 0xf7, 0xa5, 0x8a, 0x13, 0x41, 0x1d, 0x6a, 0x09, 0x55, 0x48, 0x19, 0x41, 0xba,
	0x31, 0x05, 0x45, 0x10, 0xdd, 0x55, 0xbb, 0x70, 0xe1, 0x66, 0x36, 0x2e, 0x25, 0x7f, 0x6e, 0x63,
	0x20, 0xc9, 0x1d, 0x27, 0x13, 0xd0, 0x97, 0xcd, 0xc2, 0x47, 0xc8, 0x13, 0x48, 0xee, 0x54, 0x12,
	0x5c, 0x25, 0xf7, 0x9e, 0xdf, 0x39, 0x77, 0x38, 0x6c, 0x9a, 0x42, 0x01, 0x06, 0xde, 0x92, 0xa2,
	0xa9, 0x0d, 0xe8, 0x50, 0x69, 0x34, 0xc8, 0xf7, 0x77, 0xe3, 0xfc, 0x2a, 0xcb, 0xcd, 0x7b, 0x13,
	0x87, 0x09, 0x96, 0xab, 0x0c, 0x33, 0x5c, 0x91, 0x1e, 0x37, 0x5b, 0x9a, 0x68, 0xa0, 0x3f, 0xeb,
	0x9b, 0x9f, 0x65, 0x88, 0x59, 0x01, 0x03, 0x05, 0xa5, 0x32, 0x5f, 0x56, 0x14, 0x2f, 0x6c, 0xfa,
	0x44, 0xc7, 0x1e, 0x6d, 0xb8, 0x84, 0x8f, 0x06, 0x6a, 0xc3, 0x6f, 0x19, 0xcb, 0xab, 0xda, 0x44,
	0x55, 0x02, 0xcf, 0xa9, 0xef, 0x2c, 0x9c, 0xe5, 0xc1, 0xfa, 0xb4, 0x6b, 0x83, 0x93, 0x2d, 0xea,
	0xf2, 0x5e, 0x0c, 0x9a, 0x90, 0x23, 0x50, 0x7c, 0x3b, 0xec, 0xfc, 0x4f, 0x5e, 0xad, 0xb0, 0xaa,
	0xe1, 0x55, 0x47, 0x4a, 0x81, 0xe6, 0x17, 0xcc, 0x4d, 0x30, 0x05, 0x4a, 0x9c, 0xac, 0x8f, 0xba,
	0x36, 0xf0, 0x6c, 0x62, 0xbf, 0x15, 0x92, 0x44, 0x7e, 0xc7, 0xbc, 0xfe, 0xbb, 0xf9, 0x54, 0x45,
	0x94, 0x57, 0xfe, 0x3f, 0xba, 0x3e, 0xeb, 0xda, 0x80, 0x0f, 0xec, 0x4e, 0x14, 0x72, 0x8c, 0xf2,
	0x4b, 0x36, 0x01, 0xad, 0x51, 0xfb, 0xff, 0xc9, 0x73, 0xdc, 0xb5, 0xc1, 0xa1, 0xf5, 0xd0, 0x5a,
	0x48, 0x2b, 0xf3, 0x07, 0xe6, 0xa6, 0x91, 0x89, 0x7c, 0x77, 0xe1, 0x2c, 0xbd, 0xeb, 0x59, 0x68,
	0x2b, 0x0a, 0x7f, 0x2b, 0x0a, 0x37, 0x7d, 0x45, 0xe3, 0xe7, 0xf5, 0xb4, 0x90, 0x64, 0x8a, 0xf7,
	0x08, 0xbb, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb1, 0x67, 0x72, 0xa7, 0x01, 0x00, 0x00,
}
