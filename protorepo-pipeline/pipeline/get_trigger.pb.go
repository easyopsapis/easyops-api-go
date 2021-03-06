// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_trigger.proto

package pipeline

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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
//GetTrigger请求
type GetTriggerRequest struct {
	//
	//trigger id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTriggerRequest) Reset()         { *m = GetTriggerRequest{} }
func (m *GetTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*GetTriggerRequest) ProtoMessage()    {}
func (*GetTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd941d02f2f3811, []int{0}
}
func (m *GetTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTriggerRequest.Unmarshal(m, b)
}
func (m *GetTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTriggerRequest.Marshal(b, m, deterministic)
}
func (m *GetTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTriggerRequest.Merge(m, src)
}
func (m *GetTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_GetTriggerRequest.Size(m)
}
func (m *GetTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTriggerRequest proto.InternalMessageInfo

func (m *GetTriggerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetTriggerApi返回
type GetTriggerResponseWrapper struct {
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
	Data                 *pipeline.Trigger `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetTriggerResponseWrapper) Reset()         { *m = GetTriggerResponseWrapper{} }
func (m *GetTriggerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetTriggerResponseWrapper) ProtoMessage()    {}
func (*GetTriggerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd941d02f2f3811, []int{1}
}
func (m *GetTriggerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTriggerResponseWrapper.Unmarshal(m, b)
}
func (m *GetTriggerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTriggerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetTriggerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTriggerResponseWrapper.Merge(m, src)
}
func (m *GetTriggerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetTriggerResponseWrapper.Size(m)
}
func (m *GetTriggerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTriggerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetTriggerResponseWrapper proto.InternalMessageInfo

func (m *GetTriggerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetTriggerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetTriggerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetTriggerResponseWrapper) GetData() *pipeline.Trigger {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTriggerRequest)(nil), "pipeline.GetTriggerRequest")
	proto.RegisterType((*GetTriggerResponseWrapper)(nil), "pipeline.GetTriggerResponseWrapper")
}

func init() { proto.RegisterFile("get_trigger.proto", fileDescriptor_9dd941d02f2f3811) }

var fileDescriptor_9dd941d02f2f3811 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0xcb, 0x4e, 0xe3, 0x30,
	0x14, 0x55, 0x3a, 0xed, 0x68, 0xea, 0x8e, 0x66, 0xda, 0x2c, 0x46, 0x99, 0x6e, 0xd2, 0xf1, 0x20,
	0xd4, 0x4d, 0x62, 0x4a, 0xa5, 0x0a, 0x58, 0x56, 0x02, 0xb6, 0x28, 0x42, 0x62, 0x81, 0x00, 0xb9,
	0x89, 0x6b, 0x2c, 0x92, 0x5c, 0x63, 0xbb, 0x94, 0x87, 0xf8, 0x51, 0x16, 0x41, 0xe2, 0x13, 0xf2,
	0x05, 0xa8, 0x4e, 0x0b, 0x61, 0x65, 0x9f, 0x7b, 0x1e, 0xf7, 0xe8, 0xa2, 0x1e, 0x67, 0xe6, 0xca,
	0x28, 0xc1, 0x39, 0x53, 0xa1, 0x54, 0x60, 0xc0, 0xfd, 0x21, 0x85, 0x64, 0xa9, 0xc8, 0x59, 0x3f,
	0xe0, 0xc2, 0x5c, 0x2f, 0x66, 0x61, 0x0c, 0x19, 0xe1, 0xc0, 0x81, 0x58, 0xc1, 0x6c, 0x31, 0xb7,
	0xc8, 0x02, 0xfb, 0xab, 0x8c, 0xfd, 0x13, 0x0e, 0x21, 0xa3, 0xfa, 0x01, 0xa4, 0x0e, 0x53, 0x88,
	0x69, 0x4a, 0x62, 0xc8, 0x8d, 0xa2, 0xb1, 0xd1, 0x95, 0x53, 0x31, 0x09, 0x41, 0x06, 0x09, 0x4b,
	0x35, 0x59, 0x0b, 0x89, 0x85, 0x64, 0xb3, 0x92, 0x7c, 0xa9, 0xd2, 0x9f, 0xd4, 0x0a, 0x64, 0x4b,
	0x61, 0x6e, 0x60, 0x49, 0x38, 0x04, 0x96, 0x0c, 0xee, 0x68, 0x2a, 0x12, 0x6a, 0x40, 0x69, 0xf2,
	0xf1, 0xad, 0x7c, 0xf8, 0x08, 0xf5, 0x8e, 0x99, 0x39, 0xad, 0xb2, 0x22, 0x76, 0xbb, 0x60, 0xda,
	0xb8, 0x23, 0xd4, 0x10, 0x89, 0xe7, 0x0c, 0x9c, 0x61, 0x7b, 0xfa, 0xaf, 0x2c, 0xfc, 0xf6, 0x1c,
	0x54, 0x76, 0x80, 0x45, 0x82, 0xdf, 0x5e, 0xfd, 0x2e, 0xfa, 0x75, 0x79, 0xbe, 0x13, 0xec, 0xd3,
	0xe0, 0xf1, 0xe2, 0x69, 0x34, 0x7e, 0xde, 0x8a, 0x1a, 0x22, 0xc1, 0x2f, 0x0e, 0xfa, 0x5b, 0x0f,
	0xd2, 0x12, 0x72, 0xcd, 0xce, 0x14, 0x95, 0x92, 0x29, 0xf7, 0x3f, 0x6a, 0xc6, 0x90, 0x30, 0x1b,
	0xd9, 0x9a, 0xfe, 0x2e, 0x0b, 0xbf, 0x53, 0x45, 0xae, 0xa6, 0x38, 0xb2, 0xa4, 0xbb, 0x87, 0x3a,
	0xab, 0xf7, 0xf0, 0x5e, 0xa6, 0x54, 0xe4, 0x5e, 0xc3, 0xae, 0xff, 0x53, 0x16, 0xbe, 0xfb, 0xa9,
	0x5d, 0x93, 0x38, 0xaa, 0x4b, 0xdd, 0x6d, 0xd4, 0x62, 0x4a, 0x81, 0xf2, 0xbe, 0x59, 0x4f, 0xb7,
	0x2c, 0xfc, 0x9f, 0x95, 0xc7, 0x8e, 0x71, 0x54, 0xd1, 0xee, 0x04, 0x35, 0x13, 0x6a, 0xa8, 0xd7,
	0x1c, 0x38, 0xc3, 0xce, 0x6e, 0x2f, 0xdc, 0xdc, 0x32, 0x5c, 0xd7, 0xae, 0x37, 0x5b, 0x09, 0x71,
	0x64, 0xf5, 0xb3, 0xef, 0xf6, 0x56, 0xe3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x33, 0x33,
	0x7b, 0x03, 0x02, 0x00, 0x00,
}
