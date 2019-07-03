// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_detail.proto

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
//GetDetail请求
type GetDetailRequest struct {
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

func (m *GetDetailRequest) Reset()         { *m = GetDetailRequest{} }
func (m *GetDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailRequest) ProtoMessage()    {}
func (*GetDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0076c603ebb1acf6, []int{0}
}
func (m *GetDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailRequest.Unmarshal(m, b)
}
func (m *GetDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailRequest.Merge(m, src)
}
func (m *GetDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetDetailRequest.Size(m)
}
func (m *GetDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailRequest proto.InternalMessageInfo

func (m *GetDetailRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *GetDetailRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//GetDetailApi返回
type GetDetailResponseWrapper struct {
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

func (m *GetDetailResponseWrapper) Reset()         { *m = GetDetailResponseWrapper{} }
func (m *GetDetailResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetDetailResponseWrapper) ProtoMessage()    {}
func (*GetDetailResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_0076c603ebb1acf6, []int{1}
}
func (m *GetDetailResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailResponseWrapper.Unmarshal(m, b)
}
func (m *GetDetailResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetDetailResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailResponseWrapper.Merge(m, src)
}
func (m *GetDetailResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetDetailResponseWrapper.Size(m)
}
func (m *GetDetailResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailResponseWrapper proto.InternalMessageInfo

func (m *GetDetailResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetDetailResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetDetailResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetDetailResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDetailRequest)(nil), "instance.GetDetailRequest")
	proto.RegisterType((*GetDetailResponseWrapper)(nil), "instance.GetDetailResponseWrapper")
}

func init() { proto.RegisterFile("get_detail.proto", fileDescriptor_0076c603ebb1acf6) }

var fileDescriptor_0076c603ebb1acf6 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xdd, 0x6a, 0xe2, 0x40,
	0x1c, 0xc5, 0xc9, 0xae, 0x2e, 0xee, 0xb8, 0xac, 0x71, 0x16, 0x76, 0x83, 0x2c, 0x44, 0x66, 0x65,
	0xb1, 0xa5, 0x49, 0xaa, 0x42, 0x69, 0x4b, 0x6f, 0x2a, 0xfd, 0xc0, 0xdb, 0xf4, 0xa2, 0x50, 0x51,
	0x99, 0x24, 0x63, 0x9a, 0x36, 0x66, 0xd2, 0xc9, 0xa4, 0x16, 0xc5, 0xc7, 0xe9, 0x6b, 0x45, 0xe8,
	0x23, 0xe4, 0x09, 0x4a, 0x26, 0x7e, 0xe4, 0x2a, 0x33, 0xff, 0x73, 0x7e, 0x27, 0x67, 0xfe, 0x40,
	0x76, 0x09, 0x9f, 0x38, 0x84, 0x63, 0xcf, 0xd7, 0x43, 0x46, 0x39, 0x85, 0x15, 0x2f, 0x88, 0x38,
	0x0e, 0x6c, 0xd2, 0xd0, 0x5c, 0x8f, 0x3f, 0xc6, 0x96, 0x6e, 0xd3, 0x99, 0xe1, 0x52, 0x97, 0x1a,
	0xc2, 0x60, 0xc5, 0x53, 0x71, 0x13, 0x17, 0x71, 0xca, 0xc1, 0xc6, 0x49, 0xc1, 0x3e, 0x9b, 0x7b,
	0xfc, 0x99, 0xce, 0x0d, 0x97, 0x6a, 0x42, 0xd4, 0x5e, 0xb1, 0xef, 0x39, 0x98, 0x53, 0x16, 0x19,
	0xbb, 0xe3, 0x86, 0xfb, 0xeb, 0x52, 0xea, 0xfa, 0x64, 0x9f, 0x1e, 0x71, 0x16, 0xdb, 0x3c, 0x57,
	0xd1, 0xbb, 0x04, 0xe4, 0x5b, 0xc2, 0xaf, 0x44, 0x45, 0x93, 0xbc, 0xc4, 0x24, 0xe2, 0xf0, 0x06,
	0x54, 0xa8, 0xf5, 0x44, 0x6c, 0x3e, 0x70, 0x14, 0xa9, 0x29, 0xb5, 0xbf, 0xf7, 0x0f, 0xd3, 0x44,
	0xad, 0x4d, 0x29, 0x9b, 0x9d, 0xa3, 0xad, 0x82, 0x3e, 0xd6, 0xea, 0x2f, 0x50, 0x1f, 0x0f, 0xb1,
	0xb6, 0xb8, 0xd4, 0x1e, 0x26, 0xa3, 0x65, 0xe7, 0xa8, 0xd7, 0x5d, 0xb5, 0xcc, 0x1d, 0x0b, 0x07,
	0x00, 0x6c, 0x5f, 0x3b, 0x70, 0x94, 0x2f, 0x22, 0xe9, 0x20, 0x4d, 0xd4, 0x7a, 0x9e, 0xb4, 0xd7,
	0xb2, 0x2c, 0x19, 0xfc, 0x1c, 0x0f, 0x8f, 0xb5, 0x33, 0xac, 0x2d, 0x46, 0xcb, 0x4e, 0x6f, 0xd5,
	0x32, 0x0b, 0x30, 0x5a, 0x4b, 0x40, 0x29, 0xf4, 0x8c, 0x42, 0x1a, 0x44, 0xe4, 0x9e, 0xe1, 0x30,
	0x24, 0x0c, 0xfe, 0x03, 0x25, 0x9b, 0x3a, 0x44, 0x74, 0x2d, 0xf7, 0x6b, 0x69, 0xa2, 0x56, 0xf3,
	0x3f, 0x64, 0x53, 0x64, 0x0a, 0x11, 0x9e, 0x82, 0x6a, 0xf6, 0xbd, 0x7e, 0x0b, 0x7d, 0xec, 0x05,
	0x9b, 0x36, 0xbf, 0xd3, 0x44, 0x85, 0x7b, 0xef, 0x46, 0x44, 0x66, 0xd1, 0x0a, 0xff, 0x83, 0x32,
	0x61, 0x8c, 0x32, 0xe5, 0xab, 0x60, 0xe4, 0x34, 0x51, 0x7f, 0xe4, 0x8c, 0x18, 0x23, 0x33, 0x97,
	0xe1, 0x05, 0x28, 0x39, 0x98, 0x63, 0xa5, 0xd4, 0x94, 0xda, 0xd5, 0xee, 0x1f, 0x3d, 0x5f, 0xbc,
	0xbe, 0x5d, 0xbc, 0x7e, 0x27, 0x16, 0x5f, 0xec, 0x97, 0xd9, 0x91, 0x29, 0x28, 0xeb, 0x9b, 0xf0,
	0xf5, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x45, 0xf3, 0x26, 0x33, 0x02, 0x00, 0x00,
}