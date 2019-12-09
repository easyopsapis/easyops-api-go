// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package cluster

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
	//集群Id
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
	Data                 *container.Cluster `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *GetResponseWrapper) GetData() *container.Cluster {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "cluster.GetRequest")
	proto.RegisterType((*GetResponseWrapper)(nil), "cluster.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xc9, 0xfb, 0xb6, 0xef, 0x4b, 0xb7, 0x82, 0xba, 0xa0, 0x84, 0x5e, 0x52, 0x56, 0x90,
	0x5e, 0x92, 0x05, 0xc5, 0x3f, 0x78, 0x6c, 0x11, 0xf1, 0x68, 0x2e, 0x9e, 0xb7, 0xc9, 0x34, 0x06,
	0x93, 0x4c, 0xdc, 0xdd, 0x58, 0xfd, 0x9e, 0x9e, 0xf3, 0x21, 0xf2, 0x09, 0xa4, 0xb3, 0xa1, 0xcd,
	0x69, 0x77, 0xe6, 0x79, 0x7e, 0xcf, 0x0c, 0xc3, 0x26, 0x19, 0xd8, 0xa8, 0xd6, 0x68, 0x91, 0xff,
	0x4f, 0x8a, 0xc6, 0x58, 0xd0, 0xb3, 0x30, 0xcb, 0xed, 0x5b, 0xb3, 0x8e, 0x12, 0x2c, 0x65, 0x86,
	0x19, 0x4a, 0xd2, 0xd7, 0xcd, 0x86, 0x2a, 0x2a, 0xe8, 0xe7, 0xb8, 0xd9, 0x4b, 0x86, 0x11, 0x28,
	0xf3, 0x8d, 0xb5, 0x89, 0x0a, 0x4c, 0x54, 0x21, 0x13, 0xac, 0xac, 0x56, 0x89, 0x35, 0x8e, 0xd4,
	0x50, 0x63, 0x58, 0x62, 0x0a, 0x85, 0x91, 0xbd, 0x51, 0x52, 0x49, 0x46, 0x95, 0x57, 0xa0, 0x65,
	0x3f, 0xbb, 0x8f, 0xbc, 0x1d, 0x6c, 0x50, 0x6e, 0x73, 0xfb, 0x8e, 0x5b, 0x99, 0x61, 0x48, 0x62,
	0xf8, 0xa9, 0x8a, 0x3c, 0x55, 0x16, 0xb5, 0x91, 0xfb, 0xaf, 0xe3, 0xc4, 0x8a, 0xb1, 0x27, 0xb0,
	0x31, 0x7c, 0x34, 0x60, 0x2c, 0xbf, 0x61, 0x2c, 0xaf, 0x8c, 0x55, 0x55, 0x02, 0xcf, 0xa9, 0xef,
	0xcd, 0xbd, 0xc5, 0x64, 0x79, 0xd6, 0xb5, 0xc1, 0xe9, 0x06, 0x75, 0xf9, 0x20, 0x0e, 0x9a, 0x88,
	0x07, 0x46, 0xf1, 0xe3, 0x31, 0x4e, 0x29, 0xa6, 0xc6, 0xca, 0xc0, 0xab, 0x56, 0x75, 0x0d, 0x9a,
	0x5f, 0xb0, 0x51, 0x82, 0x29, 0x50, 0xce, 0x78, 0x79, 0xdc, 0xb5, 0xc1, 0xd4, 0xe5, 0xec, 0xba,
	0x22, 0x26, 0x91, 0xdf, 0xb3, 0xe9, 0xee, 0x7d, 0xfc, 0xaa, 0x0b, 0x95, 0x57, 0xfe, 0x1f, 0x9a,
	0x79, 0xde, 0xb5, 0x01, 0x3f, 0x78, 0x7b, 0x51, 0xc4, 0x43, 0x2b, 0xbf, 0x64, 0x63, 0xd0, 0x1a,
	0xb5, 0xff, 0x97, 0x98, 0x93, 0xae, 0x0d, 0x8e, 0x1c, 0x43, 0x6d, 0x11, 0x3b, 0x99, 0xdf, 0xb1,
	0x51, 0xaa, 0xac, 0xf2, 0x47, 0x73, 0x6f, 0x31, 0xbd, 0xe2, 0xd1, 0xfe, 0x84, 0xd1, 0xca, 0x9d,
	0x70, 0xb8, 0xda, 0xce, 0x29, 0x62, 0x02, 0xd6, 0xff, 0xe8, 0x44, 0xd7, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x80, 0x77, 0x05, 0x88, 0xf2, 0x01, 0x00, 0x00,
}