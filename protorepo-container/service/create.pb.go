// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package service

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
//Create请求
type CreateRequest struct {
	//
	//部署资源组 (resource group) id
	RgId string `protobuf:"bytes,1,opt,name=rgId,proto3" json:"rgId" form:"rgId"`
	//
	//命名空间 id
	NamespaceId string `protobuf:"bytes,2,opt,name=namespaceId,proto3" json:"namespaceId" form:"namespaceId"`
	//
	//关联的 workload id
	WorkloadId string `protobuf:"bytes,3,opt,name=workloadId,proto3" json:"workloadId" form:"workloadId"`
	//
	//Service 信息
	Service              *container.Service `protobuf:"bytes,4,opt,name=service,proto3" json:"service" form:"service"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetRgId() string {
	if m != nil {
		return m.RgId
	}
	return ""
}

func (m *CreateRequest) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *CreateRequest) GetWorkloadId() string {
	if m != nil {
		return m.WorkloadId
	}
	return ""
}

func (m *CreateRequest) GetService() *container.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

//
//CreateApi返回
type CreateResponseWrapper struct {
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
	Data                 *container.Service `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseWrapper.Unmarshal(m, b)
}
func (m *CreateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseWrapper.Merge(m, src)
}
func (m *CreateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateResponseWrapper.Size(m)
}
func (m *CreateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseWrapper proto.InternalMessageInfo

func (m *CreateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponseWrapper) GetData() *container.Service {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "service.CreateRequest")
	proto.RegisterType((*CreateResponseWrapper)(nil), "service.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x8a, 0xd4, 0x40,
	0x14, 0x85, 0xe9, 0xb1, 0x47, 0xb1, 0x7a, 0x1c, 0xc7, 0x02, 0xa5, 0x19, 0x84, 0x0c, 0x35, 0x22,
	0xa3, 0x90, 0x44, 0x1d, 0xf0, 0x6f, 0x25, 0x2d, 0x2e, 0x1a, 0xdc, 0x18, 0x41, 0x17, 0xa2, 0x52,
	0x5d, 0xb9, 0x13, 0xc3, 0x24, 0xb9, 0x65, 0x55, 0xf5, 0xb4, 0x3f, 0xf8, 0x34, 0x3e, 0x94, 0xbb,
	0x08, 0x3e, 0x42, 0x9e, 0x40, 0x72, 0x2b, 0x3d, 0xc9, 0xa2, 0x71, 0xe5, 0x2a, 0x55, 0xf7, 0x9c,
	0xef, 0x24, 0x75, 0x52, 0x6c, 0x47, 0x19, 0x90, 0x0e, 0x22, 0x6d, 0xd0, 0x21, 0xbf, 0x64, 0xc1,
	0x9c, 0xe5, 0x0a, 0xf6, 0xc3, 0x2c, 0x77, 0x9f, 0x96, 0x8b, 0x48, 0x61, 0x19, 0x67, 0x98, 0x61,
	0x4c, 0xfa, 0x62, 0x79, 0x42, 0x3b, 0xda, 0xd0, 0xca, 0x73, 0xfb, 0xaf, 0x32, 0x8c, 0x40, 0xda,
	0xaf, 0xa8, 0x6d, 0x54, 0xa0, 0x92, 0x45, 0xac, 0xb0, 0x72, 0x46, 0x2a, 0x67, 0x3d, 0x69, 0x40,
	0x63, 0x58, 0x62, 0x0a, 0x85, 0x8d, 0x3b, 0x63, 0x4c, 0x5b, 0x32, 0xca, 0xbc, 0x02, 0x13, 0x77,
	0xef, 0xee, 0x22, 0xdf, 0xfc, 0xc7, 0xc8, 0x8f, 0x56, 0x83, 0xea, 0x72, 0x1f, 0x0e, 0x4e, 0x56,
	0xae, 0x72, 0x77, 0x8a, 0xab, 0x38, 0xc3, 0x90, 0xc4, 0xf0, 0x4c, 0x16, 0x79, 0x2a, 0x1d, 0x1a,
	0x1b, 0x9f, 0x2f, 0x3b, 0xee, 0x66, 0x86, 0x98, 0x15, 0xd0, 0x17, 0x61, 0x9d, 0x59, 0x2a, 0xe7,
	0x55, 0xf1, 0x73, 0x8b, 0x5d, 0x79, 0x4e, 0x4d, 0x26, 0xf0, 0x79, 0x09, 0xd6, 0xf1, 0x47, 0x6c,
	0x6c, 0xb2, 0x79, 0x3a, 0x1d, 0x1d, 0x8c, 0x8e, 0x2e, 0xcf, 0x0e, 0x9b, 0x3a, 0x98, 0x9c, 0xa0,
	0x29, 0x9f, 0x8a, 0x76, 0x2a, 0xfe, 0xfc, 0x0e, 0xf6, 0xd8, 0xee, 0x87, 0x77, 0xf7, 0xc2, 0x27,
	0x32, 0xfc, 0xf6, 0xfe, 0xfb, 0xfd, 0xe3, 0x1f, 0xb7, 0x12, 0x02, 0xf8, 0x4b, 0x36, 0xa9, 0x64,
	0x09, 0x56, 0x4b, 0x05, 0xf3, 0x74, 0xba, 0x45, 0xfc, 0xdd, 0xa6, 0x0e, 0xb8, 0xe7, 0x07, 0xe2,
	0xe6, 0x98, 0x21, 0xce, 0xe7, 0x8c, 0xad, 0xd0, 0x9c, 0x16, 0x28, 0xd3, 0x79, 0x3a, 0xbd, 0x40,
	0x61, 0x77, 0x9a, 0x3a, 0xb8, 0xe6, 0xc3, 0x7a, 0x6d, 0x73, 0xd6, 0x00, 0xe6, 0xcf, 0xd8, 0xfa,
	0x7a, 0x4c, 0xc7, 0x07, 0xa3, 0xa3, 0xc9, 0x03, 0x1e, 0x9d, 0x37, 0x1d, 0xbd, 0xf6, 0xca, 0x8c,
	0x37, 0x75, 0xb0, 0xeb, 0xb3, 0x3b, 0xb3, 0x48, 0xd6, 0x98, 0xf8, 0x35, 0x62, 0xd7, 0xd7, 0x2d,
	0x59, 0x8d, 0x95, 0x85, 0xb7, 0x46, 0x6a, 0x0d, 0x86, 0x1f, 0xb2, 0xb1, 0xc2, 0x14, 0xa8, 0xad,
	0xed, 0xd9, 0xd5, 0xbe, 0xad, 0x76, 0x2a, 0x12, 0x12, 0xf9, 0x63, 0x36, 0x69, 0x9f, 0x2f, 0xbe,
	0xe8, 0x42, 0xe6, 0x55, 0xd7, 0xcc, 0x8d, 0xbe, 0x99, 0x81, 0x28, 0x92, 0xa1, 0x95, 0xdf, 0x66,
	0xdb, 0x60, 0x0c, 0x9a, 0xae, 0x80, 0xbd, 0xa6, 0x0e, 0x76, 0x3c, 0x43, 0x63, 0x91, 0x78, 0xb9,
	0xfd, 0x69, 0xa9, 0x74, 0xf2, 0x1f, 0xe7, 0x1b, 0x7c, 0x5a, 0xeb, 0x14, 0x09, 0x01, 0x8b, 0x8b,
	0x74, 0x0d, 0x8e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x05, 0x93, 0x22, 0x96, 0x4f, 0x03, 0x00,
	0x00,
}
