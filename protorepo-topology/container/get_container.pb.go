// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_container.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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
//GetContainer请求
type GetContainerRequest struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContainerRequest) Reset()         { *m = GetContainerRequest{} }
func (m *GetContainerRequest) String() string { return proto.CompactTextString(m) }
func (*GetContainerRequest) ProtoMessage()    {}
func (*GetContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b221bc53fae6eb11, []int{0}
}
func (m *GetContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerRequest.Unmarshal(m, b)
}
func (m *GetContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerRequest.Marshal(b, m, deterministic)
}
func (m *GetContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerRequest.Merge(m, src)
}
func (m *GetContainerRequest) XXX_Size() int {
	return xxx_messageInfo_GetContainerRequest.Size(m)
}
func (m *GetContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerRequest proto.InternalMessageInfo

func (m *GetContainerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetContainerApi返回
type GetContainerResponseWrapper struct {
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
	Data                 *topology.Container `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetContainerResponseWrapper) Reset()         { *m = GetContainerResponseWrapper{} }
func (m *GetContainerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetContainerResponseWrapper) ProtoMessage()    {}
func (*GetContainerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b221bc53fae6eb11, []int{1}
}
func (m *GetContainerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerResponseWrapper.Unmarshal(m, b)
}
func (m *GetContainerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetContainerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerResponseWrapper.Merge(m, src)
}
func (m *GetContainerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetContainerResponseWrapper.Size(m)
}
func (m *GetContainerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerResponseWrapper proto.InternalMessageInfo

func (m *GetContainerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetContainerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetContainerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetContainerResponseWrapper) GetData() *topology.Container {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetContainerRequest)(nil), "container.GetContainerRequest")
	proto.RegisterType((*GetContainerResponseWrapper)(nil), "container.GetContainerResponseWrapper")
}

func init() { proto.RegisterFile("get_container.proto", fileDescriptor_b221bc53fae6eb11) }

var fileDescriptor_b221bc53fae6eb11 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xff, 0xed, 0x1f, 0xba, 0x55, 0x94, 0x2d, 0x48, 0xa8, 0x48, 0xca, 0x0a, 0xd2,
	0x4b, 0x13, 0x50, 0x0f, 0xc5, 0x63, 0x45, 0x3c, 0xbb, 0x17, 0x8f, 0xb2, 0x4d, 0xa6, 0x71, 0x21,
	0xcd, 0xac, 0xbb, 0x5b, 0xb0, 0x2f, 0x9b, 0x87, 0xc8, 0x13, 0x48, 0x26, 0x69, 0xad, 0x9e, 0x3d,
	0xed, 0xce, 0x37, 0xdf, 0xef, 0xdb, 0xd9, 0x61, 0xe3, 0x1c, 0xfc, 0x5b, 0x8a, 0xa5, 0x57, 0xba,
	0x04, 0x1b, 0x1b, 0x8b, 0x1e, 0xf9, 0xf0, 0x20, 0x4c, 0xe6, 0xb9, 0xf6, 0xef, 0xdb, 0x55, 0x9c,
	0xe2, 0x26, 0xc9, 0x31, 0xc7, 0x84, 0x1c, 0xab, 0xed, 0x9a, 0x2a, 0x2a, 0xe8, 0xd6, 0x92, 0x93,
	0x97, 0x1c, 0x63, 0x50, 0x6e, 0x87, 0xc6, 0xc5, 0x05, 0xa6, 0xaa, 0x48, 0x9a, 0x28, 0xab, 0x52,
	0xef, 0x5a, 0xd2, 0x82, 0xc1, 0xf9, 0x06, 0x33, 0x28, 0x5c, 0xd2, 0x19, 0x13, 0x2a, 0x13, 0x8f,
	0x06, 0x0b, 0xcc, 0x77, 0x8d, 0xcf, 0x80, 0xf5, 0xbb, 0x2e, 0x52, 0xfe, 0x41, 0xe4, 0xaf, 0x0f,
	0x8a, 0x7b, 0x36, 0x7e, 0x06, 0xff, 0xb8, 0x57, 0x25, 0x7c, 0x6c, 0xc1, 0x79, 0x7e, 0xc5, 0x7a,
	0x3a, 0x0b, 0x83, 0x69, 0x30, 0x1b, 0x2e, 0x4f, 0xeb, 0x2a, 0x1a, 0xae, 0xd1, 0x6e, 0x1e, 0x84,
	0xce, 0x84, 0xec, 0xe9, 0x4c, 0x54, 0x01, 0xbb, 0xfc, 0x89, 0x39, 0x83, 0xa5, 0x83, 0x57, 0xab,
	0x8c, 0x01, 0xcb, 0xaf, 0x59, 0x3f, 0xc5, 0x0c, 0x28, 0x60, 0xb0, 0x3c, 0xab, 0xab, 0x68, 0xd4,
	0x06, 0x34, 0xaa, 0x90, 0xd4, 0xe4, 0x0b, 0x36, 0x6a, 0xce, 0xa7, 0x4f, 0x53, 0x28, 0x5d, 0x86,
	0x3d, 0x7a, 0xec, 0xa2, 0xae, 0x22, 0xfe, 0xed, 0xed, 0x9a, 0x42, 0x1e, 0x5b, 0xf9, 0x0d, 0x1b,
	0x80, 0xb5, 0x68, 0xc3, 0x7f, 0xc4, 0x9c, 0xd7, 0x55, 0x74, 0xd2, 0x32, 0x24, 0x0b, 0xd9, 0xb6,
	0xf9, 0x82, 0xf5, 0x33, 0xe5, 0x55, 0xd8, 0x9f, 0x06, 0xb3, 0xd1, 0xed, 0x38, 0xde, 0x6f, 0x21,
	0x3e, 0x0c, 0x7e, 0x3c, 0x5b, 0x63, 0x15, 0x92, 0x88, 0xd5, 0x7f, 0xda, 0xce, 0xdd, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0x7b, 0x24, 0xc7, 0x15, 0x02, 0x00, 0x00,
}
