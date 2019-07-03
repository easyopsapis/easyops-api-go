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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x61, 0x6b, 0xda, 0x40,
	0x1c, 0xc6, 0xc9, 0xa6, 0xc3, 0x9d, 0x63, 0xea, 0x0d, 0xb7, 0x20, 0x1b, 0x91, 0x9b, 0x0c, 0x37,
	0x96, 0x64, 0x2a, 0x8c, 0x6d, 0xec, 0xcd, 0x64, 0x1b, 0xe4, 0x6d, 0xfa, 0xa2, 0x50, 0x51, 0xb9,
	0x24, 0x67, 0x9a, 0x36, 0xe6, 0xd2, 0xcb, 0xa5, 0x16, 0xc5, 0x2f, 0xd4, 0x0f, 0x95, 0x42, 0xe9,
	0x27, 0xc8, 0x27, 0x28, 0xb9, 0x24, 0x1a, 0xfa, 0x2a, 0xf7, 0xbf, 0xe7, 0xf9, 0x3d, 0x79, 0x8e,
	0x3f, 0xe8, 0x3a, 0xc4, 0x27, 0x9c, 0x2c, 0xbd, 0x20, 0xe2, 0x38, 0xb0, 0x89, 0x16, 0x32, 0xca,
	0x29, 0x6c, 0x94, 0x73, 0x4f, 0x75, 0x3d, 0x7e, 0x1e, 0x5b, 0x9a, 0x4d, 0xd7, 0xba, 0x4b, 0x5d,
	0xaa, 0x0b, 0x83, 0x15, 0xaf, 0xc4, 0x24, 0x06, 0x71, 0xca, 0xc1, 0xde, 0xf7, 0x8a, 0x7d, 0xbd,
	0xf1, 0xf8, 0x25, 0xdd, 0xe8, 0x2e, 0x55, 0x85, 0xa8, 0x5e, 0x63, 0xdf, 0x73, 0x30, 0xa7, 0x2c,
	0xd2, 0x0f, 0xc7, 0x82, 0x7b, 0xef, 0x52, 0xea, 0xfa, 0xe4, 0x98, 0x1e, 0x71, 0x16, 0xdb, 0x3c,
	0x57, 0xd1, 0xad, 0x04, 0xba, 0x7f, 0x45, 0x51, 0xa3, 0xe8, 0x65, 0x92, 0xab, 0x98, 0x44, 0x1c,
	0xfe, 0x07, 0x0d, 0x6a, 0x5d, 0x10, 0x9b, 0x1b, 0x8e, 0x2c, 0xf5, 0xa5, 0xe1, 0xcb, 0xe9, 0x97,
	0x34, 0x51, 0x5a, 0x2b, 0xca, 0xd6, 0xbf, 0x50, 0xa9, 0xa0, 0xfb, 0x3b, 0xe5, 0x0d, 0xe8, 0x2c,
	0x66, 0x58, 0xdd, 0xfe, 0x51, 0xcf, 0x96, 0xf3, 0xdd, 0xe8, 0xeb, 0x64, 0xbc, 0x1f, 0x98, 0x07,
	0x16, 0x1a, 0x00, 0x94, 0x4f, 0x36, 0x1c, 0xf9, 0x99, 0x48, 0xfa, 0x9c, 0x26, 0x4a, 0x27, 0x4f,
	0x3a, 0x6a, 0x59, 0x56, 0x1b, 0xbc, 0x5e, 0xcc, 0xbe, 0xa9, 0x3f, 0xb1, 0xba, 0x9d, 0xef, 0x46,
	0x93, 0xfd, 0xc0, 0xac, 0xc0, 0xe8, 0x41, 0x02, 0x1f, 0x9e, 0x96, 0x8d, 0x42, 0x1a, 0x44, 0xe4,
	0x94, 0xe1, 0x30, 0x24, 0x0c, 0x7e, 0x04, 0x35, 0x9b, 0x3a, 0x44, 0x14, 0xae, 0x4f, 0x5b, 0x69,
	0xa2, 0x34, 0xf3, 0xdf, 0x64, 0xb7, 0xc8, 0x14, 0x22, 0xfc, 0x01, 0x9a, 0xd9, 0xf7, 0xdf, 0x4d,
	0xe8, 0x63, 0x2f, 0x28, 0x2a, 0xbd, 0x4d, 0x13, 0x05, 0x1e, 0xbd, 0x85, 0x88, 0xcc, 0xaa, 0x15,
	0x7e, 0x02, 0x75, 0xc2, 0x18, 0x65, 0xf2, 0x73, 0xc1, 0xb4, 0xd3, 0x44, 0x79, 0x95, 0x33, 0xe2,
	0x1a, 0x99, 0xb9, 0x0c, 0x7f, 0x83, 0x9a, 0x83, 0x39, 0x96, 0x6b, 0x7d, 0x69, 0xd8, 0x1c, 0xbf,
	0xd3, 0xf2, 0x15, 0x68, 0xe5, 0x0a, 0xb4, 0x13, 0xb1, 0x82, 0x6a, 0xbf, 0xcc, 0x8e, 0x4c, 0x41,
	0x59, 0x2f, 0x84, 0x6f, 0xf2, 0x18, 0x00, 0x00, 0xff, 0xff, 0x97, 0x33, 0x9c, 0x5d, 0x42, 0x02,
	0x00, 0x00,
}
