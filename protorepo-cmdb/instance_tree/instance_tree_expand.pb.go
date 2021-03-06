// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance_tree_expand.proto

package instance_tree

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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
//InstanceTreeExpand请求
type InstanceTreeExpandRequest struct {
	//
	//树定义
	Tree *cmdb.InstanceTreeRootNode `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree" form:"tree"`
	//
	//全局忽略与父级无关联实例,默认false
	IgnoreSingle bool `protobuf:"varint,2,opt,name=ignore_single,json=ignoreSingle,proto3" json:"ignore_single" form:"ignore_single"`
	//
	//模型Id,不传则展开首层
	ObjectId string `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//实例Id
	InstanceId           string   `protobuf:"bytes,4,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceTreeExpandRequest) Reset()         { *m = InstanceTreeExpandRequest{} }
func (m *InstanceTreeExpandRequest) String() string { return proto.CompactTextString(m) }
func (*InstanceTreeExpandRequest) ProtoMessage()    {}
func (*InstanceTreeExpandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd93a24077f5b3, []int{0}
}
func (m *InstanceTreeExpandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceTreeExpandRequest.Unmarshal(m, b)
}
func (m *InstanceTreeExpandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceTreeExpandRequest.Marshal(b, m, deterministic)
}
func (m *InstanceTreeExpandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceTreeExpandRequest.Merge(m, src)
}
func (m *InstanceTreeExpandRequest) XXX_Size() int {
	return xxx_messageInfo_InstanceTreeExpandRequest.Size(m)
}
func (m *InstanceTreeExpandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceTreeExpandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceTreeExpandRequest proto.InternalMessageInfo

func (m *InstanceTreeExpandRequest) GetTree() *cmdb.InstanceTreeRootNode {
	if m != nil {
		return m.Tree
	}
	return nil
}

func (m *InstanceTreeExpandRequest) GetIgnoreSingle() bool {
	if m != nil {
		return m.IgnoreSingle
	}
	return false
}

func (m *InstanceTreeExpandRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *InstanceTreeExpandRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//InstanceTreeExpandApi返回
type InstanceTreeExpandResponseWrapper struct {
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
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InstanceTreeExpandResponseWrapper) Reset()         { *m = InstanceTreeExpandResponseWrapper{} }
func (m *InstanceTreeExpandResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InstanceTreeExpandResponseWrapper) ProtoMessage()    {}
func (*InstanceTreeExpandResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd93a24077f5b3, []int{1}
}
func (m *InstanceTreeExpandResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceTreeExpandResponseWrapper.Unmarshal(m, b)
}
func (m *InstanceTreeExpandResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceTreeExpandResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InstanceTreeExpandResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceTreeExpandResponseWrapper.Merge(m, src)
}
func (m *InstanceTreeExpandResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InstanceTreeExpandResponseWrapper.Size(m)
}
func (m *InstanceTreeExpandResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceTreeExpandResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceTreeExpandResponseWrapper proto.InternalMessageInfo

func (m *InstanceTreeExpandResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InstanceTreeExpandResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InstanceTreeExpandResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InstanceTreeExpandResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceTreeExpandRequest)(nil), "instance_tree.InstanceTreeExpandRequest")
	proto.RegisterType((*InstanceTreeExpandResponseWrapper)(nil), "instance_tree.InstanceTreeExpandResponseWrapper")
}

func init() { proto.RegisterFile("instance_tree_expand.proto", fileDescriptor_9ccd93a24077f5b3) }

var fileDescriptor_9ccd93a24077f5b3 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x54, 0x4a, 0x8a, 0x9a, 0x4d, 0x0b, 0xc1, 0x42, 0x10, 0x22, 0x90, 0x83, 0xa9, 0x50, 0x90,
	0xb0, 0xdd, 0x36, 0x12, 0x5f, 0x02, 0x21, 0x22, 0xf5, 0x90, 0x0b, 0x87, 0x6d, 0x25, 0x24, 0x4a,
	0x6b, 0x6d, 0xbc, 0xaf, 0xc6, 0x60, 0xfb, 0x99, 0xdd, 0x0d, 0x2d, 0x54, 0xfd, 0x67, 0xfc, 0x16,
	0x23, 0x71, 0xe2, 0xec, 0x5f, 0x80, 0xbc, 0xeb, 0xa4, 0x8e, 0xc4, 0x29, 0x6f, 0x76, 0x66, 0x5e,
	0xc6, 0xa3, 0x47, 0x06, 0x71, 0x26, 0x15, 0xcb, 0x42, 0x08, 0x94, 0x00, 0x08, 0xe0, 0x3c, 0x67,
	0x19, 0xf7, 0x72, 0x81, 0x0a, 0xad, 0xad, 0x15, 0x6e, 0xe0, 0x46, 0xb1, 0xfa, 0x3c, 0x9f, 0x79,
	0x21, 0xa6, 0x7e, 0x84, 0x11, 0xfa, 0x5a, 0x35, 0x9b, 0x9f, 0x6a, 0xa4, 0x81, 0x9e, 0x8c, 0x7b,
	0xf0, 0x29, 0x42, 0x0f, 0x98, 0xfc, 0x81, 0xb9, 0xf4, 0x12, 0x0c, 0x59, 0xe2, 0x87, 0x98, 0x29,
	0xc1, 0x42, 0x25, 0x8d, 0x53, 0x40, 0x8e, 0x6e, 0x8a, 0x1c, 0x12, 0xe9, 0xd7, 0x42, 0x5f, 0x43,
	0x3f, 0x4c, 0xf9, 0xcc, 0x5f, 0x0d, 0x26, 0x10, 0x55, 0x90, 0x21, 0x87, 0x7a, 0xfb, 0xb3, 0x46,
	0x98, 0xf4, 0x2c, 0x56, 0x5f, 0xf1, 0xcc, 0x8f, 0xd0, 0xd5, 0xa4, 0xfb, 0x9d, 0x25, 0x31, 0x67,
	0x0a, 0x85, 0xf4, 0x97, 0x63, 0xed, 0xbb, 0x1f, 0x21, 0x46, 0x09, 0x5c, 0x65, 0x97, 0x4a, 0xcc,
	0x43, 0x65, 0x58, 0xe7, 0xd7, 0x1a, 0xb9, 0x37, 0xad, 0xff, 0xf7, 0x50, 0x00, 0xec, 0xeb, 0x3a,
	0x28, 0x7c, 0x9b, 0x83, 0x54, 0xd6, 0x5b, 0xd2, 0xae, 0xb2, 0xf4, 0x5b, 0xc3, 0xd6, 0xa8, 0xbb,
	0x37, 0xf0, 0xaa, 0x98, 0x5e, 0x53, 0x4e, 0x11, 0xd5, 0x7b, 0xe4, 0x30, 0xb9, 0x59, 0x16, 0x76,
	0xf7, 0x14, 0x45, 0xfa, 0xca, 0xa9, 0x1c, 0x0e, 0xd5, 0x46, 0xeb, 0x0d, 0xd9, 0x8a, 0xa3, 0x0c,
	0x05, 0x04, 0x32, 0xce, 0xa2, 0x04, 0xfa, 0x6b, 0xc3, 0xd6, 0x68, 0x63, 0xd2, 0x2f, 0x0b, 0xfb,
	0xb6, 0x51, 0xaf, 0xd0, 0x0e, 0xdd, 0x34, 0xf8, 0x40, 0x43, 0xeb, 0x90, 0x74, 0x70, 0xf6, 0x05,
	0x42, 0x15, 0xc4, 0xbc, 0x7f, 0x6d, 0xd8, 0x1a, 0x75, 0x26, 0xcf, 0xcb, 0xc2, 0xee, 0x19, 0xeb,
	0x92, 0x72, 0xfe, 0xfc, 0xb6, 0x6d, 0xf2, 0xe0, 0xe4, 0x88, 0xb9, 0x3f, 0xdf, 0xb9, 0x1f, 0x83,
	0xe3, 0xa3, 0x1d, 0xf7, 0xe5, 0x62, 0xbe, 0xd8, 0x79, 0x3a, 0xde, 0xbd, 0xdc, 0xa6, 0x1b, 0x46,
	0x3e, 0xe5, 0xd6, 0x94, 0x90, 0x45, 0xd5, 0x53, 0xde, 0x6f, 0xeb, 0xb5, 0x4f, 0xca, 0xc2, 0xbe,
	0x55, 0x27, 0x5a, 0x72, 0xd5, 0xde, 0x1e, 0xb9, 0x71, 0x52, 0xaf, 0x3b, 0xbe, 0xd8, 0x1d, 0x5f,
	0x6e, 0xd3, 0x86, 0xd9, 0xf9, 0xdb, 0x22, 0x0f, 0xff, 0x57, 0x9f, 0xcc, 0x31, 0x93, 0xf0, 0x41,
	0xb0, 0x3c, 0x07, 0x61, 0x3d, 0x22, 0xed, 0x10, 0xb9, 0xa9, 0x71, 0xbd, 0x59, 0x55, 0xf5, 0xea,
	0x50, 0x4d, 0x5a, 0x2f, 0x48, 0xb7, 0xfa, 0xdd, 0x3f, 0xcf, 0x13, 0x16, 0x67, 0xba, 0xa8, 0xce,
	0xe4, 0x4e, 0x59, 0xd8, 0xd6, 0x95, 0xb6, 0x26, 0x1d, 0xda, 0x94, 0x5a, 0x8f, 0xc9, 0x3a, 0x08,
	0x81, 0xa2, 0x6e, 0xa8, 0x57, 0x16, 0xf6, 0xa6, 0xf1, 0xe8, 0x67, 0x87, 0x1a, 0xda, 0x7a, 0x4d,
	0xda, 0x9c, 0x29, 0xa6, 0xbf, 0xb8, 0xbb, 0x77, 0xd7, 0x33, 0x87, 0xe1, 0x2d, 0x0e, 0xc3, 0x3b,
	0xd0, 0x87, 0xd1, 0xcc, 0x57, 0xc9, 0x1d, 0xaa, 0x5d, 0xb3, 0xeb, 0x5a, 0x37, 0xfe, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0xb8, 0x8e, 0x40, 0x40, 0x03, 0x00, 0x00,
}
