// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_instance.proto

package instance

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
//DeleteInstance请求
type DeleteInstanceRequest struct {
	//
	//模型对象ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例ID
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf20306375d260d, []int{0}
}
func (m *DeleteInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceRequest.Unmarshal(m, b)
}
func (m *DeleteInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceRequest.Merge(m, src)
}
func (m *DeleteInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceRequest.Size(m)
}
func (m *DeleteInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceRequest proto.InternalMessageInfo

func (m *DeleteInstanceRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *DeleteInstanceRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//DeleteInstanceApi返回
type DeleteInstanceResponseWrapper struct {
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

func (m *DeleteInstanceResponseWrapper) Reset()         { *m = DeleteInstanceResponseWrapper{} }
func (m *DeleteInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceResponseWrapper) ProtoMessage()    {}
func (*DeleteInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf20306375d260d, []int{1}
}
func (m *DeleteInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceResponseWrapper.Merge(m, src)
}
func (m *DeleteInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceResponseWrapper.Size(m)
}
func (m *DeleteInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceResponseWrapper proto.InternalMessageInfo

func (m *DeleteInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteInstanceResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteInstanceRequest)(nil), "instance.DeleteInstanceRequest")
	proto.RegisterType((*DeleteInstanceResponseWrapper)(nil), "instance.DeleteInstanceResponseWrapper")
}

func init() { proto.RegisterFile("delete_instance.proto", fileDescriptor_1bf20306375d260d) }

var fileDescriptor_1bf20306375d260d = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0xda, 0x4a, 0x9d, 0x8a, 0xad, 0x03, 0xd5, 0x50, 0x2c, 0x29, 0x63, 0x91, 0x0a, 0x26,
	0x69, 0x2d, 0x14, 0x15, 0x6f, 0x2c, 0x7a, 0x91, 0xdb, 0xf1, 0x42, 0xb0, 0xb4, 0x65, 0x92, 0x4c,
	0x63, 0x34, 0xcd, 0x64, 0x27, 0x93, 0xed, 0xb2, 0xa5, 0xaf, 0xb5, 0x8f, 0x93, 0x85, 0x65, 0x9f,
	0x20, 0x4f, 0xb0, 0x64, 0x92, 0xb4, 0x61, 0xaf, 0x72, 0xce, 0xf9, 0x7e, 0xf2, 0x0d, 0x1f, 0xe8,
	0xb9, 0x34, 0xa0, 0x82, 0x6e, 0xfc, 0x30, 0x16, 0x24, 0x74, 0xa8, 0x11, 0x71, 0x26, 0x18, 0x6c,
	0x55, 0x7b, 0x5f, 0xf7, 0x7c, 0xf1, 0x37, 0xb1, 0x0d, 0x87, 0xed, 0x4c, 0x8f, 0x79, 0xcc, 0x94,
	0x04, 0x3b, 0xd9, 0xca, 0x4d, 0x2e, 0x72, 0x2a, 0x84, 0xfd, 0x79, 0x8d, 0xbe, 0xdb, 0xfb, 0xe2,
	0x3f, 0xdb, 0x9b, 0x1e, 0xd3, 0x25, 0xa8, 0x5f, 0x92, 0xc0, 0x77, 0x89, 0x60, 0x3c, 0x36, 0x4f,
	0x63, 0xa9, 0x7b, 0xeb, 0x31, 0xe6, 0x05, 0xf4, 0xec, 0x1e, 0x0b, 0x9e, 0x38, 0xa2, 0x40, 0xd1,
	0x8d, 0x02, 0x7a, 0x3f, 0x64, 0x50, 0xab, 0xcc, 0x85, 0xe9, 0x45, 0x42, 0x63, 0x01, 0x31, 0x68,
	0x31, 0xfb, 0x1f, 0x75, 0x84, 0xe5, 0xaa, 0xca, 0x50, 0x19, 0x3f, 0x5f, 0xcc, 0xb3, 0x54, 0xeb,
	0x6c, 0x19, 0xdf, 0x7d, 0x45, 0x15, 0x82, 0xee, 0x6e, 0x35, 0x0d, 0x0c, 0xd6, 0x4b, 0xa2, 0x5f,
	0x7f, 0xd7, 0xff, 0x6c, 0x56, 0xcb, 0x89, 0xfe, 0xa5, 0x9a, 0x0f, 0x93, 0x8f, 0xb3, 0xe9, 0x71,
	0x84, 0x4f, 0x3e, 0xd0, 0x02, 0xa0, 0x7a, 0xbe, 0xe5, 0xaa, 0x4f, 0xa4, 0xeb, 0x87, 0x2c, 0xd5,
	0x5e, 0x15, 0xae, 0x67, 0x2c, 0xf7, 0xed, 0x82, 0x97, 0xeb, 0xd2, 0x6e, 0x75, 0x98, 0xce, 0x8e,
	0x23, 0x5c, 0x13, 0xa3, 0x7b, 0x05, 0x0c, 0x1e, 0x07, 0x8f, 0x23, 0x16, 0xc6, 0xf4, 0x37, 0x27,
	0x51, 0x44, 0x39, 0x7c, 0x07, 0x1a, 0x0e, 0x73, 0xa9, 0x0c, 0xdf, 0x5c, 0x74, 0xb2, 0x54, 0x6b,
	0x17, 0xbf, 0xc9, 0xaf, 0x08, 0x4b, 0x10, 0x7e, 0x06, 0xed, 0xfc, 0xfb, 0xf3, 0x2a, 0x0a, 0x88,
	0x1f, 0x96, 0x91, 0x5e, 0x67, 0xa9, 0x06, 0xcf, 0xdc, 0x12, 0x44, 0xb8, 0x4e, 0x85, 0xef, 0x41,
	0x93, 0x72, 0xce, 0xb8, 0xfa, 0x54, 0x6a, 0xba, 0x59, 0xaa, 0xbd, 0x28, 0x34, 0xf2, 0x8c, 0x70,
	0x01, 0xc3, 0x6f, 0xa0, 0xe1, 0x12, 0x41, 0xd4, 0xc6, 0x50, 0x19, 0xb7, 0x3f, 0xbd, 0x31, 0x8a,
	0x3a, 0x8c, 0xaa, 0x0e, 0xe3, 0x97, 0xac, 0xa3, 0x9e, 0x2f, 0xa7, 0x23, 0x2c, 0x55, 0xf6, 0x33,
	0xc9, 0x9b, 0x3d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x83, 0x76, 0x40, 0xa1, 0x4e, 0x02, 0x00, 0x00,
}
