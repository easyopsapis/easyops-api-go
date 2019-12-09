// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: detail_configuration_template.proto

package template

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	collector_center "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/collector_center"
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
//DetailConfigurationTemplate请求
type DetailConfigurationTemplateRequest struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailConfigurationTemplateRequest) Reset()         { *m = DetailConfigurationTemplateRequest{} }
func (m *DetailConfigurationTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*DetailConfigurationTemplateRequest) ProtoMessage()    {}
func (*DetailConfigurationTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeb42ce1ce8a9cf1, []int{0}
}
func (m *DetailConfigurationTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailConfigurationTemplateRequest.Unmarshal(m, b)
}
func (m *DetailConfigurationTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailConfigurationTemplateRequest.Marshal(b, m, deterministic)
}
func (m *DetailConfigurationTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailConfigurationTemplateRequest.Merge(m, src)
}
func (m *DetailConfigurationTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_DetailConfigurationTemplateRequest.Size(m)
}
func (m *DetailConfigurationTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailConfigurationTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailConfigurationTemplateRequest proto.InternalMessageInfo

func (m *DetailConfigurationTemplateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DetailConfigurationTemplateApi返回
type DetailConfigurationTemplateResponseWrapper struct {
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
	Data                 *collector_center.ConfigurationTemplate `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *DetailConfigurationTemplateResponseWrapper) Reset() {
	*m = DetailConfigurationTemplateResponseWrapper{}
}
func (m *DetailConfigurationTemplateResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*DetailConfigurationTemplateResponseWrapper) ProtoMessage() {}
func (*DetailConfigurationTemplateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeb42ce1ce8a9cf1, []int{1}
}
func (m *DetailConfigurationTemplateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailConfigurationTemplateResponseWrapper.Unmarshal(m, b)
}
func (m *DetailConfigurationTemplateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailConfigurationTemplateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DetailConfigurationTemplateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailConfigurationTemplateResponseWrapper.Merge(m, src)
}
func (m *DetailConfigurationTemplateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DetailConfigurationTemplateResponseWrapper.Size(m)
}
func (m *DetailConfigurationTemplateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailConfigurationTemplateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DetailConfigurationTemplateResponseWrapper proto.InternalMessageInfo

func (m *DetailConfigurationTemplateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DetailConfigurationTemplateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DetailConfigurationTemplateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DetailConfigurationTemplateResponseWrapper) GetData() *collector_center.ConfigurationTemplate {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DetailConfigurationTemplateRequest)(nil), "template.DetailConfigurationTemplateRequest")
	proto.RegisterType((*DetailConfigurationTemplateResponseWrapper)(nil), "template.DetailConfigurationTemplateResponseWrapper")
}

func init() {
	proto.RegisterFile("detail_configuration_template.proto", fileDescriptor_aeb42ce1ce8a9cf1)
}

var fileDescriptor_aeb42ce1ce8a9cf1 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0xcf, 0x6b, 0x14, 0x3f,
	0x14, 0x67, 0xf7, 0xbb, 0xfd, 0x62, 0xb3, 0x45, 0x65, 0x0e, 0xb2, 0x2c, 0xc2, 0x94, 0x14, 0xb4,
	0x08, 0x33, 0x91, 0x2a, 0xa2, 0x1e, 0xab, 0xde, 0x3c, 0x0d, 0x4a, 0x0f, 0x52, 0x96, 0x6c, 0xe6,
	0x6d, 0x0c, 0x66, 0xe6, 0xc5, 0x97, 0x37, 0x56, 0x11, 0x0f, 0xfe, 0xa3, 0x2b, 0xf8, 0x27, 0xec,
	0x5f, 0x20, 0x9b, 0x19, 0xed, 0x22, 0xd2, 0x53, 0x4f, 0xc9, 0xcb, 0xe7, 0x57, 0xf2, 0x89, 0x38,
	0xaa, 0x81, 0xb5, 0xf3, 0x0b, 0x83, 0xed, 0xca, 0xd9, 0x8e, 0x34, 0x3b, 0x6c, 0x17, 0x0c, 0x4d,
	0xf0, 0x9a, 0xa1, 0x0c, 0x84, 0x8c, 0xd9, 0x8d, 0xdf, 0xf3, 0xbc, 0xb0, 0x8e, 0xdf, 0x77, 0xcb,
	0xd2, 0x60, 0xa3, 0x2c, 0x5a, 0x54, 0x89, 0xb0, 0xec, 0x56, 0x69, 0x4a, 0x43, 0xda, 0xf5, 0xc2,
	0xf9, 0x99, 0xc5, 0x12, 0x74, 0xfc, 0x82, 0x21, 0x96, 0x1e, 0x8d, 0xf6, 0xca, 0x60, 0xcb, 0xa4,
	0x0d, 0xc7, 0x5e, 0x49, 0x10, 0xb0, 0x68, 0xb0, 0x06, 0x1f, 0xd5, 0x40, 0x54, 0x69, 0x54, 0x06,
	0xbd, 0x07, 0xc3, 0x48, 0x0b, 0x03, 0x2d, 0x03, 0xa9, 0x68, 0xc8, 0x05, 0x1e, 0x8c, 0xcf, 0xaf,
	0xd1, 0x98, 0x35, 0x59, 0xe0, 0x05, 0xe9, 0xd6, 0x0e, 0x0f, 0x9e, 0xdb, 0x6b, 0xb4, 0xbf, 0xaa,
	0xd9, 0xf9, 0x93, 0x9d, 0x3e, 0x9b, 0x0b, 0xc7, 0x1f, 0xf0, 0x42, 0x59, 0x2c, 0x12, 0x58, 0x7c,
	0xd2, 0xde, 0xd5, 0x9a, 0x91, 0xa2, 0xfa, 0xb3, 0x1d, 0x74, 0x77, 0x2d, 0xa2, 0xf5, 0x70, 0x59,
	0x7f, 0x64, 0xea, 0xcc, 0xd0, 0x8e, 0x7c, 0x2b, 0xe4, 0xcb, 0xf4, 0xad, 0x2f, 0x76, 0xb3, 0xdf,
	0x0c, 0xd1, 0x15, 0x7c, 0xec, 0x20, 0x72, 0xa6, 0xc4, 0xd8, 0xd5, 0xb3, 0xd1, 0xe1, 0xe8, 0x78,
	0xff, 0x34, 0xdf, 0xac, 0xf3, 0xfd, 0x15, 0x52, 0xf3, 0x5c, 0xba, 0x5a, 0xfe, 0xfc, 0x91, 0xdf,
	0x14, 0x07, 0xef, 0x1e, 0x16, 0xcf, 0x74, 0xb1, 0x3a, 0xff, 0x7a, 0xf2, 0xf8, 0x5b, 0x35, 0x76,
	0xb5, 0xfc, 0x3e, 0x16, 0x0f, 0xae, 0xf4, 0x8d, 0x01, 0xdb, 0x08, 0x67, 0xa4, 0x43, 0x00, 0xca,
	0x8e, 0xc4, 0xc4, 0x60, 0x0d, 0x29, 0x61, 0xef, 0xf4, 0xd6, 0x66, 0x9d, 0x4f, 0xfb, 0x84, 0xed,
	0xa9, 0xac, 0x12, 0x98, 0x3d, 0x15, 0xd3, 0xed, 0xfa, 0xea, 0x73, 0xf0, 0xda, 0xb5, 0xb3, 0x71,
	0xba, 0xcd, 0x9d, 0xcd, 0x3a, 0xcf, 0x2e, 0xb9, 0x03, 0x28, 0xab, 0x5d, 0x6a, 0x76, 0x4f, 0xec,
	0x01, 0x11, 0xd2, 0xec, 0xbf, 0xa4, 0xb9, 0xbd, 0x59, 0xe7, 0x07, 0xbd, 0x26, 0x1d, 0xcb, 0xaa,
	0x87, 0xb3, 0xd7, 0x62, 0x52, 0x6b, 0xd6, 0xb3, 0xc9, 0xe1, 0xe8, 0x78, 0x7a, 0x72, 0xbf, 0xfc,
	0xfb, 0x83, 0xca, 0x7f, 0x3e, 0x66, 0xf7, 0xbe, 0x5b, 0xb9, 0xac, 0x92, 0xcb, 0xf2, 0xff, 0xd4,
	0xf0, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x74, 0xab, 0xb2, 0x38, 0x03, 0x00, 0x00,
}