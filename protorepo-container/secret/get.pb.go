// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package secret

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	container "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
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
//Get请求
type GetRequest struct {
	//
	//secret instance id
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//GetApi返回
type GetResponseWrapper struct {
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
	Data                 *container.Secret `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetResponseWrapper) Reset()         { *m = GetResponseWrapper{} }
func (m *GetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetResponseWrapper) ProtoMessage()    {}
func (*GetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseWrapper.Unmarshal(m, b)
}
func (m *GetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseWrapper.Merge(m, src)
}
func (m *GetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetResponseWrapper.Size(m)
}
func (m *GetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseWrapper proto.InternalMessageInfo

func (m *GetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetResponseWrapper) GetData() *container.Secret {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "secret.GetRequest")
	proto.RegisterType((*GetResponseWrapper)(nil), "secret.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x26, 0xda, 0x7b, 0xe1, 0x4e, 0x45, 0xef, 0x9d, 0x85, 0x84, 0x22, 0xa4, 0x8c, 0x22, 0x75,
	0x91, 0x8c, 0x5a, 0x28, 0xea, 0xb2, 0x20, 0xd2, 0x9d, 0xc4, 0x45, 0x17, 0xa2, 0x30, 0x9d, 0x9c,
	0x8e, 0xc1, 0x24, 0x27, 0xce, 0x4c, 0xac, 0x3f, 0xf8, 0x94, 0xee, 0x23, 0xf8, 0x08, 0x79, 0x82,
	0x4b, 0x4f, 0x42, 0x9b, 0xd5, 0x9c, 0x73, 0xbe, 0x9f, 0xf9, 0xf8, 0xd8, 0x95, 0x01, 0x9f, 0xd4,
	0x16, 0x3d, 0xf2, 0x4b, 0x07, 0xda, 0x82, 0x9f, 0xc5, 0x26, 0xf7, 0x5f, 0x9a, 0x5d, 0xa2, 0xb1,
	0x94, 0x06, 0x0d, 0x4a, 0x82, 0x77, 0xcd, 0x9e, 0x36, 0x5a, 0x68, 0xea, 0x65, 0xb3, 0xf7, 0x06,
	0x13, 0x50, 0xee, 0x27, 0xd6, 0x2e, 0x29, 0x50, 0xab, 0x42, 0x6a, 0xac, 0xbc, 0x55, 0xda, 0xbb,
	0x5e, 0x69, 0xa1, 0xc6, 0xb8, 0xc4, 0x0c, 0x0a, 0x27, 0x07, 0xa2, 0xa4, 0x95, 0x88, 0x2a, 0xaf,
	0xc0, 0xca, 0xfe, 0xeb, 0xc1, 0x71, 0x35, 0x0a, 0x50, 0x1e, 0x72, 0xff, 0x15, 0x0f, 0xd2, 0x60,
	0x4c, 0x60, 0xfc, 0x5d, 0x15, 0x79, 0xa6, 0x3c, 0x5a, 0x27, 0x4f, 0xe3, 0xa0, 0x7b, 0x64, 0x10,
	0x4d, 0x01, 0xe7, 0xbc, 0xce, 0xdb, 0x46, 0x0f, 0xae, 0x62, 0xcb, 0xd8, 0x3b, 0xf0, 0x29, 0x7c,
	0x6b, 0xc0, 0x79, 0xbe, 0x61, 0x2c, 0xaf, 0x9c, 0x57, 0x95, 0x86, 0x4d, 0x16, 0x06, 0xf3, 0x60,
	0x71, 0xb5, 0x7e, 0xd6, 0xb5, 0xd1, 0xcd, 0x1e, 0x6d, 0xf9, 0x46, 0x9c, 0x31, 0xf1, 0xff, 0x5f,
	0x74, 0xcd, 0xee, 0x7f, 0xfe, 0xf8, 0x3c, 0x7e, 0xad, 0xe2, 0x5f, 0x9f, 0x7e, 0xbf, 0x58, 0xfe,
	0x79, 0x92, 0x8e, 0xc4, 0xe2, 0x6f, 0xc0, 0x38, 0x39, 0xbb, 0x1a, 0x2b, 0x07, 0x5b, 0xab, 0xea,
	0x1a, 0x2c, 0x7f, 0xcc, 0x26, 0x1a, 0x33, 0x20, 0xef, 0x8b, 0xf5, 0x83, 0xae, 0x8d, 0xa6, 0xbd,
	0xf7, 0xf1, 0x2a, 0x52, 0x02, 0xf9, 0x2b, 0x36, 0x3d, 0xbe, 0x6f, 0x7f, 0xd4, 0x85, 0xca, 0xab,
	0xf0, 0x0e, 0xe5, 0x78, 0xd8, 0xb5, 0x11, 0x3f, 0x73, 0x07, 0x50, 0xa4, 0x63, 0x2a, 0x7f, 0xca,
	0x2e, 0xc0, 0x5a, 0xb4, 0xe1, 0x5d, 0xd2, 0x5c, 0x77, 0x6d, 0x74, 0xaf, 0xd7, 0xd0, 0x59, 0xa4,
	0x3d, 0xcc, 0x57, 0x6c, 0x92, 0x29, 0xaf, 0xc2, 0xc9, 0x3c, 0x58, 0x4c, 0x5f, 0xde, 0x24, 0xa7,
	0xce, 0x93, 0x0f, 0xd4, 0xf9, 0x38, 0xd9, 0x91, 0x28, 0x52, 0xe2, 0xef, 0x2e, 0xa9, 0xb5, 0xe5,
	0x6d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x48, 0x38, 0x9c, 0x21, 0x02, 0x00, 0x00,
}