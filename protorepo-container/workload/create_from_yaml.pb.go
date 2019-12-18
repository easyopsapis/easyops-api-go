// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_from_yaml.proto

package workload

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
//CreateFromYaml请求
type CreateFromYamlRequest struct {
	//
	//部署资源组 (resource group) id
	RgId string `protobuf:"bytes,1,opt,name=rgId,proto3" json:"rgId" form:"rgId"`
	//
	//资源原生 yaml 文件定义
	ResourceSpec         string   `protobuf:"bytes,2,opt,name=resourceSpec,proto3" json:"resourceSpec" form:"resourceSpec"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFromYamlRequest) Reset()         { *m = CreateFromYamlRequest{} }
func (m *CreateFromYamlRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFromYamlRequest) ProtoMessage()    {}
func (*CreateFromYamlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b10b3b9fd11a2d4, []int{0}
}
func (m *CreateFromYamlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFromYamlRequest.Unmarshal(m, b)
}
func (m *CreateFromYamlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFromYamlRequest.Marshal(b, m, deterministic)
}
func (m *CreateFromYamlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFromYamlRequest.Merge(m, src)
}
func (m *CreateFromYamlRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFromYamlRequest.Size(m)
}
func (m *CreateFromYamlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFromYamlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFromYamlRequest proto.InternalMessageInfo

func (m *CreateFromYamlRequest) GetRgId() string {
	if m != nil {
		return m.RgId
	}
	return ""
}

func (m *CreateFromYamlRequest) GetResourceSpec() string {
	if m != nil {
		return m.ResourceSpec
	}
	return ""
}

//
//CreateFromYamlApi返回
type CreateFromYamlResponseWrapper struct {
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
	Data                 *container.Workload `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateFromYamlResponseWrapper) Reset()         { *m = CreateFromYamlResponseWrapper{} }
func (m *CreateFromYamlResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateFromYamlResponseWrapper) ProtoMessage()    {}
func (*CreateFromYamlResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b10b3b9fd11a2d4, []int{1}
}
func (m *CreateFromYamlResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFromYamlResponseWrapper.Unmarshal(m, b)
}
func (m *CreateFromYamlResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFromYamlResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateFromYamlResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFromYamlResponseWrapper.Merge(m, src)
}
func (m *CreateFromYamlResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateFromYamlResponseWrapper.Size(m)
}
func (m *CreateFromYamlResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFromYamlResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFromYamlResponseWrapper proto.InternalMessageInfo

func (m *CreateFromYamlResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateFromYamlResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateFromYamlResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateFromYamlResponseWrapper) GetData() *container.Workload {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateFromYamlRequest)(nil), "workload.CreateFromYamlRequest")
	proto.RegisterType((*CreateFromYamlResponseWrapper)(nil), "workload.CreateFromYamlResponseWrapper")
}

func init() { proto.RegisterFile("create_from_yaml.proto", fileDescriptor_4b10b3b9fd11a2d4) }

var fileDescriptor_4b10b3b9fd11a2d4 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x48, 0x11, 0x6c, 0x2a, 0xa8, 0x5c, 0x51, 0xa2, 0x0a, 0xe4, 0xca, 0x45, 0xa8,
	0x17, 0xdb, 0x40, 0x25, 0x28, 0x70, 0x2b, 0x02, 0x89, 0x1b, 0x72, 0x0f, 0x15, 0xaa, 0x8a, 0xd9,
	0xac, 0x27, 0xc6, 0x74, 0xd7, 0x63, 0x66, 0xd7, 0x0d, 0x01, 0xf1, 0x08, 0xbc, 0x62, 0x40, 0x3c,
	0x42, 0x9e, 0x00, 0x65, 0xbd, 0xc1, 0x81, 0xb3, 0x4f, 0xde, 0x99, 0xff, 0xff, 0x3f, 0xcd, 0xda,
	0x63, 0xb6, 0x23, 0x08, 0xb8, 0x81, 0x6c, 0x42, 0xa8, 0xb2, 0x19, 0x57, 0x32, 0xae, 0x09, 0x0d,
	0xfa, 0xd7, 0xa7, 0x48, 0x17, 0x12, 0x79, 0xbe, 0x1b, 0x15, 0xa5, 0xf9, 0xd8, 0x8c, 0x63, 0x81,
	0x2a, 0x29, 0xb0, 0xc0, 0xc4, 0x1a, 0xc6, 0xcd, 0xc4, 0x56, 0xb6, 0xb0, 0xa7, 0x36, 0xb8, 0x7b,
	0x52, 0x60, 0x0c, 0x5c, 0xcf, 0xb0, 0xd6, 0xb1, 0x44, 0xc1, 0x65, 0x22, 0xb0, 0x32, 0xc4, 0x85,
	0xd1, 0x6d, 0x92, 0xa0, 0xc6, 0x48, 0x61, 0x0e, 0x52, 0x27, 0xce, 0x98, 0xd8, 0xd2, 0x1a, 0x79,
	0x59, 0x01, 0x75, 0x27, 0x07, 0x7d, 0xdb, 0x07, 0xf4, 0x12, 0x65, 0xa3, 0xc0, 0x11, 0xcf, 0xfb,
	0x20, 0xe6, 0x50, 0x4b, 0x9c, 0x29, 0xa8, 0x4c, 0xa6, 0x0d, 0x71, 0x03, 0xc5, 0xcc, 0xe1, 0x3f,
	0xf4, 0x81, 0xb7, 0xc1, 0x0c, 0xc7, 0x9f, 0x40, 0x98, 0x8c, 0x60, 0x02, 0x04, 0x95, 0x58, 0x5d,
	0xe0, 0xac, 0xf7, 0x0b, 0x70, 0xd3, 0x68, 0x07, 0x4f, 0xfb, 0x80, 0xaf, 0x36, 0xc8, 0x31, 0x9f,
	0xac, 0xed, 0x91, 0x9a, 0x96, 0xe6, 0x02, 0xa7, 0x49, 0x81, 0x91, 0x15, 0xa3, 0x4b, 0x2e, 0xcb,
	0x9c, 0x1b, 0x24, 0x9d, 0xfc, 0x3d, 0xba, 0xdc, 0xdd, 0x02, 0xb1, 0x90, 0xd0, 0xad, 0x9d, 0x36,
	0xd4, 0x08, 0xd3, 0xaa, 0xe1, 0x0f, 0x8f, 0xdd, 0x7e, 0x69, 0x57, 0xf8, 0x35, 0xa1, 0x7a, 0xc7,
	0x95, 0x4c, 0xe1, 0x73, 0x03, 0xda, 0xf8, 0x4f, 0xd9, 0x80, 0x8a, 0x37, 0xf9, 0xc8, 0xdb, 0xf3,
	0x0e, 0x6e, 0x1c, 0xef, 0x2f, 0xe6, 0xc1, 0x70, 0x82, 0xa4, 0x9e, 0x87, 0xcb, 0x6e, 0xf8, 0xfb,
	0x67, 0xb0, 0xc5, 0x6e, 0xbe, 0x3f, 0x7b, 0x18, 0x3d, 0xe3, 0xd1, 0xd7, 0xf3, 0x6f, 0x8f, 0x0e,
	0xbf, 0xdf, 0x4f, 0x6d, 0xc0, 0x7f, 0xc1, 0x36, 0x09, 0x34, 0x36, 0x24, 0xe0, 0xa4, 0x06, 0x31,
	0xba, 0x62, 0x01, 0x77, 0x16, 0xf3, 0x60, 0xdb, 0x01, 0xd6, 0xd4, 0x30, 0xfd, 0xc7, 0x1c, 0xfe,
	0xf2, 0xd8, 0xbd, 0xff, 0xe7, 0xd1, 0x35, 0x56, 0x1a, 0x4e, 0x89, 0xd7, 0x35, 0x90, 0xbf, 0xcf,
	0x06, 0x02, 0x73, 0xb0, 0x73, 0x6d, 0x1c, 0xdf, 0xea, 0xe6, 0x5a, 0x76, 0xc3, 0xd4, 0x8a, 0xfe,
	0x11, 0x1b, 0x2e, 0x9f, 0xaf, 0xbe, 0xd4, 0x92, 0x97, 0x95, 0x1b, 0x61, 0x67, 0x31, 0x0f, 0xfc,
	0xce, 0xeb, 0xc4, 0x30, 0x5d, 0xb7, 0xfa, 0x0f, 0xd8, 0x06, 0x10, 0x21, 0x8d, 0xae, 0xda, 0xcc,
	0xd6, 0x62, 0x1e, 0x6c, 0xb6, 0x19, 0xdb, 0x0e, 0xd3, 0x56, 0xf6, 0x8f, 0xd8, 0x20, 0xe7, 0x86,
	0x8f, 0x06, 0x7b, 0xde, 0xc1, 0xf0, 0xf1, 0x76, 0xdc, 0xfd, 0x72, 0xa7, 0xee, 0xbb, 0xad, 0xcf,
	0xb6, 0xb4, 0x86, 0xa9, 0x4d, 0x8c, 0xaf, 0xd9, 0x37, 0x7f, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0xcd, 0xe9, 0x22, 0x3b, 0x04, 0x00, 0x00,
}