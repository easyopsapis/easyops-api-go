// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: set.proto

package instance_relation

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//Set请求
type SetRequest struct {
	//
	//模型Id
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//关系sideId
	RelationSideId string `protobuf:"bytes,2,opt,name=relationSideId,proto3" json:"relationSideId" form:"relationSideId"`
	//
	//模型的实例ID列表
	InstanceIds []string `protobuf:"bytes,3,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids" form:"instance_ids"`
	//
	//关联的实例ID列表
	RelatedInstanceIds   []string `protobuf:"bytes,4,rep,name=related_instance_ids,json=relatedInstanceIds,proto3" json:"related_instance_ids" form:"related_instance_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d650fd95c5da449, []int{0}
}
func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *SetRequest) GetRelationSideId() string {
	if m != nil {
		return m.RelationSideId
	}
	return ""
}

func (m *SetRequest) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

func (m *SetRequest) GetRelatedInstanceIds() []string {
	if m != nil {
		return m.RelatedInstanceIds
	}
	return nil
}

//
//SetApi返回
type SetResponseWrapper struct {
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

func (m *SetResponseWrapper) Reset()         { *m = SetResponseWrapper{} }
func (m *SetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SetResponseWrapper) ProtoMessage()    {}
func (*SetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d650fd95c5da449, []int{1}
}
func (m *SetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponseWrapper.Unmarshal(m, b)
}
func (m *SetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponseWrapper.Merge(m, src)
}
func (m *SetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SetResponseWrapper.Size(m)
}
func (m *SetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponseWrapper proto.InternalMessageInfo

func (m *SetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SetResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SetRequest)(nil), "instance_relation.SetRequest")
	proto.RegisterType((*SetResponseWrapper)(nil), "instance_relation.SetResponseWrapper")
}

func init() { proto.RegisterFile("set.proto", fileDescriptor_2d650fd95c5da449) }

var fileDescriptor_2d650fd95c5da449 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0xb5, 0x43, 0xd4, 0x9d, 0xb6, 0x61, 0x60, 0x8a, 0x36, 0x21, 0x57, 0x61, 0x42,
	0x95, 0x20, 0xe9, 0x46, 0xa5, 0x89, 0x5f, 0x17, 0x2a, 0xed, 0xd0, 0x0b, 0x87, 0xec, 0x80, 0xc4,
	0xb4, 0x55, 0x6e, 0xfd, 0x16, 0x0c, 0x6d, 0x5e, 0xb0, 0x5d, 0x06, 0x4c, 0xfb, 0x3f, 0x39, 0x05,
	0x89, 0x3f, 0x21, 0x77, 0x24, 0x14, 0x3b, 0x69, 0xb3, 0x89, 0x03, 0xa7, 0x3e, 0xfb, 0xfb, 0x7d,
	0x9f, 0xaf, 0xfb, 0xf2, 0x48, 0x5b, 0x83, 0x09, 0x53, 0x85, 0x06, 0xe9, 0x3d, 0x99, 0x68, 0xc3,
	0x93, 0x29, 0x8c, 0x15, 0xcc, 0xb8, 0x91, 0x98, 0xec, 0x06, 0xb1, 0x34, 0x1f, 0x17, 0x93, 0x70,
	0x8a, 0xf3, 0x7e, 0x8c, 0x31, 0xf6, 0xad, 0x73, 0xb2, 0xb8, 0xb0, 0x27, 0x7b, 0xb0, 0x95, 0x23,
	0xec, 0x1e, 0xd5, 0xec, 0xf3, 0x4b, 0x69, 0x3e, 0xe3, 0x65, 0x3f, 0xc6, 0xc0, 0x8a, 0xc1, 0x57,
	0x3e, 0x93, 0x82, 0x1b, 0x54, 0xba, 0xbf, 0x2c, 0xcb, 0xbe, 0xbd, 0x18, 0x31, 0x9e, 0xc1, 0x8a,
	0x0e, 0xf3, 0xd4, 0x7c, 0x77, 0xa2, 0xff, 0x67, 0x8d, 0x90, 0x13, 0x30, 0x11, 0x7c, 0x59, 0x80,
	0x36, 0x34, 0x22, 0x77, 0x71, 0xf2, 0x09, 0xa6, 0x66, 0x24, 0xbc, 0x46, 0xb7, 0xd1, 0x6b, 0x0f,
	0x8f, 0xf2, 0x8c, 0x6d, 0x5d, 0xa0, 0x9a, 0xbf, 0xf2, 0x2b, 0xc5, 0xff, 0xfd, 0x8b, 0x31, 0xf2,
	0xe8, 0xfc, 0x94, 0x07, 0x3f, 0xde, 0x06, 0x1f, 0xc6, 0x67, 0xa7, 0x07, 0xc1, 0xcb, 0xaa, 0xbe,
	0x3a, 0x78, 0x36, 0x38, 0xbc, 0xde, 0x8f, 0x96, 0x1c, 0x2a, 0xc8, 0x66, 0xf5, 0x97, 0x4f, 0xa4,
	0x80, 0x91, 0xf0, 0xd6, 0x2c, 0xf9, 0x4d, 0x9e, 0xb1, 0x87, 0x8e, 0x7c, 0x53, 0xff, 0x2f, 0xfe,
	0x2d, 0x26, 0x7d, 0x47, 0x36, 0x96, 0x13, 0x96, 0x42, 0x7b, 0xcd, 0x6e, 0xb3, 0xd7, 0x1e, 0x3e,
	0xcd, 0x33, 0x76, 0xdf, 0x65, 0xd4, 0xd5, 0x22, 0x61, 0x9b, 0x6c, 0x9e, 0x97, 0xe0, 0xb3, 0xab,
	0xc3, 0xc1, 0xf5, 0x7e, 0xd4, 0xa9, 0x2c, 0x23, 0xa1, 0x29, 0x90, 0x07, 0x36, 0x01, 0xc4, 0xf8,
	0x06, 0xb7, 0x65, 0xb9, 0x83, 0x3c, 0x63, 0x7b, 0xb5, 0xb7, 0xdf, 0x72, 0xfd, 0x9b, 0x4f, 0x4b,
	0xeb, 0x68, 0x15, 0xe3, 0xff, 0x6c, 0x10, 0x6a, 0xe7, 0xaf, 0x53, 0x4c, 0x34, 0xbc, 0x57, 0x3c,
	0x4d, 0x41, 0xd1, 0xc7, 0xa4, 0x35, 0x45, 0x01, 0xf6, 0x1b, 0xac, 0x0f, 0xb7, 0xf2, 0x8c, 0x75,
	0x5c, 0x5a, 0x71, 0xeb, 0x47, 0x56, 0xa4, 0x2f, 0x48, 0xa7, 0xf8, 0x3d, 0xfe, 0x96, 0xce, 0xb8,
	0x4c, 0xca, 0xa9, 0xee, 0xe4, 0x19, 0xa3, 0x2b, 0x6f, 0x29, 0xfa, 0x51, 0xdd, 0x4a, 0x9f, 0x90,
	0x75, 0x50, 0x0a, 0x95, 0xd7, 0xb4, 0x3d, 0xdb, 0x79, 0xc6, 0x36, 0x5c, 0x8f, 0xbd, 0xf6, 0x23,
	0x27, 0xd3, 0xd7, 0xa4, 0x25, 0xb8, 0xe1, 0x5e, 0xab, 0xdb, 0xe8, 0x75, 0x9e, 0xef, 0x84, 0x6e,
	0x93, 0xc2, 0x6a, 0x93, 0xc2, 0xe3, 0x62, 0x93, 0xea, 0xcf, 0x2b, 0xdc, 0x7e, 0x64, 0x9b, 0x26,
	0x77, 0xac, 0x6d, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xfe, 0x6f, 0x67, 0x05, 0x03, 0x00,
	0x00,
}
