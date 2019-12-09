// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_instance.proto

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
//CreateInstance请求
type CreateInstanceRequest struct {
	//
	//objectId
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//创建实例
	Instance             *types.Struct `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f49522ca8c6de9, []int{0}
}
func (m *CreateInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceRequest.Unmarshal(m, b)
}
func (m *CreateInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceRequest.Marshal(b, m, deterministic)
}
func (m *CreateInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceRequest.Merge(m, src)
}
func (m *CreateInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceRequest.Size(m)
}
func (m *CreateInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceRequest proto.InternalMessageInfo

func (m *CreateInstanceRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *CreateInstanceRequest) GetInstance() *types.Struct {
	if m != nil {
		return m.Instance
	}
	return nil
}

//
//CreateInstanceApi返回
type CreateInstanceResponseWrapper struct {
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

func (m *CreateInstanceResponseWrapper) Reset()         { *m = CreateInstanceResponseWrapper{} }
func (m *CreateInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceResponseWrapper) ProtoMessage()    {}
func (*CreateInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f49522ca8c6de9, []int{1}
}
func (m *CreateInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *CreateInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceResponseWrapper.Merge(m, src)
}
func (m *CreateInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceResponseWrapper.Size(m)
}
func (m *CreateInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceResponseWrapper proto.InternalMessageInfo

func (m *CreateInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateInstanceResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateInstanceRequest)(nil), "instance.CreateInstanceRequest")
	proto.RegisterType((*CreateInstanceResponseWrapper)(nil), "instance.CreateInstanceResponseWrapper")
}

func init() { proto.RegisterFile("create_instance.proto", fileDescriptor_39f49522ca8c6de9) }

var fileDescriptor_39f49522ca8c6de9 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xc9, 0xbd, 0xed, 0xa5, 0x9d, 0x5e, 0xa8, 0x8c, 0x54, 0x43, 0xb1, 0xa4, 0x8c, 0x22,
	0x5d, 0x98, 0xa4, 0x5a, 0x28, 0x2a, 0x6e, 0x8c, 0x08, 0x76, 0x3b, 0x2e, 0x04, 0x8b, 0x96, 0x49,
	0x32, 0x8d, 0xd1, 0x36, 0x13, 0x27, 0x13, 0x2b, 0x8a, 0x8f, 0xe5, 0xeb, 0x44, 0x10, 0x9f, 0x20,
	0x4f, 0x20, 0x9d, 0x34, 0x69, 0x71, 0xe1, 0x6a, 0xce, 0x39, 0xff, 0xff, 0x9f, 0xf9, 0x0e, 0x68,
	0x38, 0x9c, 0x12, 0x41, 0x47, 0x7e, 0x10, 0x09, 0x12, 0x38, 0xd4, 0x08, 0x39, 0x13, 0x0c, 0x56,
	0xf2, 0xbe, 0xa9, 0x7b, 0xbe, 0xb8, 0x8b, 0x6d, 0xc3, 0x61, 0x53, 0xd3, 0x63, 0x1e, 0x33, 0xa5,
	0xc1, 0x8e, 0xc7, 0xb2, 0x93, 0x8d, 0xac, 0xb2, 0x60, 0xb3, 0xbf, 0x62, 0x9f, 0xce, 0x7c, 0xf1,
	0xc0, 0x66, 0xa6, 0xc7, 0x74, 0x29, 0xea, 0x4f, 0x64, 0xe2, 0xbb, 0x44, 0x30, 0x1e, 0x99, 0x45,
	0xb9, 0xc8, 0x6d, 0x79, 0x8c, 0x79, 0x13, 0xba, 0xdc, 0x1e, 0x09, 0x1e, 0x3b, 0x22, 0x53, 0xd1,
	0xbb, 0x02, 0x1a, 0x67, 0x12, 0x74, 0xb0, 0xe0, 0xc2, 0xf4, 0x31, 0xa6, 0x91, 0x80, 0x18, 0x54,
	0x98, 0x7d, 0x4f, 0x1d, 0x31, 0x70, 0x55, 0xa5, 0xad, 0x74, 0xaa, 0x56, 0x3f, 0x4d, 0xb4, 0xfa,
	0x98, 0xf1, 0xe9, 0x31, 0xca, 0x15, 0xf4, 0xf9, 0xa1, 0x69, 0xa0, 0x75, 0x3b, 0x24, 0xfa, 0xcb,
	0xa9, 0x7e, 0x3d, 0xba, 0x19, 0x76, 0xf5, 0xa3, 0xbc, 0x7e, 0xed, 0xee, 0xf5, 0xf6, 0xdf, 0x76,
	0x70, 0xb1, 0x07, 0x5e, 0x80, 0xe2, 0x7c, 0xf5, 0x4f, 0x5b, 0xe9, 0xd4, 0x0e, 0x36, 0x8d, 0x0c,
	0xcf, 0xc8, 0xf1, 0x8c, 0x4b, 0x89, 0x67, 0xad, 0x2f, 0x3f, 0xcb, 0x23, 0x08, 0x17, 0x69, 0xf4,
	0xa5, 0x80, 0xd6, 0x4f, 0xee, 0x28, 0x64, 0x41, 0x44, 0xaf, 0x38, 0x09, 0x43, 0xca, 0xe1, 0x36,
	0x28, 0x39, 0xcc, 0xa5, 0x92, 0xbd, 0x6c, 0xd5, 0xd3, 0x44, 0xab, 0x65, 0xeb, 0xe6, 0x53, 0x84,
	0xa5, 0x08, 0x0f, 0x41, 0x6d, 0xfe, 0x9e, 0x3f, 0x87, 0x13, 0xe2, 0x07, 0x92, 0xa9, 0x6a, 0x6d,
	0xa4, 0x89, 0x06, 0x97, 0xde, 0x85, 0x88, 0xf0, 0xaa, 0x15, 0xee, 0x82, 0x32, 0xe5, 0x9c, 0x71,
	0xf5, 0xaf, 0xcc, 0xac, 0xa5, 0x89, 0xf6, 0x3f, 0xcb, 0xc8, 0x31, 0xc2, 0x99, 0x0c, 0x4f, 0x40,
	0xc9, 0x25, 0x82, 0xa8, 0xa5, 0xdf, 0xcf, 0x5d, 0xe1, 0x9b, 0xdb, 0x11, 0x96, 0x29, 0xfb, 0x9f,
	0xf4, 0xf5, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62, 0xac, 0xec, 0x07, 0x4d, 0x02, 0x00, 0x00,
}
