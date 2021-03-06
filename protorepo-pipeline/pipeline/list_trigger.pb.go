// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_trigger.proto

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
//ListTrigger请求
type ListTriggerRequest struct {
	//
	//按pipeline id过滤
	PipelineId           string   `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id" form:"pipeline_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTriggerRequest) Reset()         { *m = ListTriggerRequest{} }
func (m *ListTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*ListTriggerRequest) ProtoMessage()    {}
func (*ListTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f09fec2ebdef0e78, []int{0}
}
func (m *ListTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggerRequest.Unmarshal(m, b)
}
func (m *ListTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggerRequest.Marshal(b, m, deterministic)
}
func (m *ListTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggerRequest.Merge(m, src)
}
func (m *ListTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_ListTriggerRequest.Size(m)
}
func (m *ListTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggerRequest proto.InternalMessageInfo

func (m *ListTriggerRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

//
//ListTrigger返回
type ListTriggerResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*pipeline.Trigger `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListTriggerResponse) Reset()         { *m = ListTriggerResponse{} }
func (m *ListTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*ListTriggerResponse) ProtoMessage()    {}
func (*ListTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f09fec2ebdef0e78, []int{1}
}
func (m *ListTriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggerResponse.Unmarshal(m, b)
}
func (m *ListTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggerResponse.Marshal(b, m, deterministic)
}
func (m *ListTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggerResponse.Merge(m, src)
}
func (m *ListTriggerResponse) XXX_Size() int {
	return xxx_messageInfo_ListTriggerResponse.Size(m)
}
func (m *ListTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggerResponse proto.InternalMessageInfo

func (m *ListTriggerResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTriggerResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTriggerResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListTriggerResponse) GetList() []*pipeline.Trigger {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListTriggerApi返回
type ListTriggerResponseWrapper struct {
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
	Data                 *ListTriggerResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListTriggerResponseWrapper) Reset()         { *m = ListTriggerResponseWrapper{} }
func (m *ListTriggerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListTriggerResponseWrapper) ProtoMessage()    {}
func (*ListTriggerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f09fec2ebdef0e78, []int{2}
}
func (m *ListTriggerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggerResponseWrapper.Unmarshal(m, b)
}
func (m *ListTriggerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListTriggerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggerResponseWrapper.Merge(m, src)
}
func (m *ListTriggerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListTriggerResponseWrapper.Size(m)
}
func (m *ListTriggerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggerResponseWrapper proto.InternalMessageInfo

func (m *ListTriggerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListTriggerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListTriggerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListTriggerResponseWrapper) GetData() *ListTriggerResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListTriggerRequest)(nil), "pipeline.ListTriggerRequest")
	proto.RegisterType((*ListTriggerResponse)(nil), "pipeline.ListTriggerResponse")
	proto.RegisterType((*ListTriggerResponseWrapper)(nil), "pipeline.ListTriggerResponseWrapper")
}

func init() { proto.RegisterFile("list_trigger.proto", fileDescriptor_f09fec2ebdef0e78) }

var fileDescriptor_f09fec2ebdef0e78 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xb6, 0xbb, 0x59, 0xd9, 0x4e, 0x44, 0xeb, 0x88, 0x12, 0x02, 0x92, 0x32, 0x8a, 0x54, 0x21,
	0x89, 0xee, 0xc2, 0xfa, 0xe3, 0x18, 0xf0, 0x20, 0x7a, 0x90, 0x28, 0x78, 0x10, 0x2d, 0xd3, 0x64,
	0x76, 0x1c, 0x4c, 0xfb, 0xc6, 0x99, 0xa9, 0xab, 0x15, 0xff, 0x50, 0x2f, 0x15, 0xbc, 0x78, 0xef,
	0x5f, 0x20, 0xf3, 0xd2, 0xd8, 0x2c, 0xec, 0x69, 0xe6, 0xbd, 0xef, 0xfb, 0xde, 0xfb, 0xbe, 0x64,
	0x08, 0x6d, 0x94, 0x75, 0x53, 0x67, 0x94, 0x94, 0xc2, 0x64, 0xda, 0x80, 0x03, 0x7a, 0xa8, 0x95,
	0x16, 0x8d, 0x5a, 0x88, 0x38, 0x95, 0xca, 0x7d, 0x5a, 0xce, 0xb2, 0x0a, 0xe6, 0xb9, 0x04, 0x09,
	0x39, 0x12, 0x66, 0xcb, 0x53, 0xac, 0xb0, 0xc0, 0x5b, 0x2b, 0x8c, 0x5f, 0x4b, 0xc8, 0x04, 0xb7,
	0xdf, 0x41, 0xdb, 0xac, 0x81, 0x8a, 0x37, 0x79, 0x05, 0x0b, 0x67, 0x78, 0xe5, 0x6c, 0xab, 0x34,
	0x42, 0x43, 0x3a, 0x87, 0x5a, 0x34, 0x36, 0xdf, 0x12, 0x73, 0x2c, 0xf3, 0x6e, 0x65, 0x7e, 0xce,
	0x4a, 0x7c, 0xd2, 0x33, 0x30, 0x3f, 0x53, 0xee, 0x33, 0x9c, 0xe5, 0x12, 0x52, 0x04, 0xd3, 0xaf,
	0xbc, 0x51, 0x35, 0x77, 0x60, 0x6c, 0xfe, 0xff, 0xda, 0xea, 0x18, 0x27, 0xf4, 0x95, 0xb2, 0xee,
	0x6d, 0x3b, 0xac, 0x14, 0x5f, 0x96, 0xc2, 0x3a, 0xfa, 0x92, 0x84, 0xdd, 0x9e, 0xa9, 0xaa, 0xa3,
	0xc1, 0x78, 0x30, 0x19, 0x16, 0x0f, 0x36, 0xeb, 0x84, 0x9e, 0x82, 0x99, 0x3f, 0x63, 0x3d, 0x90,
	0xfd, 0xf9, 0x9d, 0x8c, 0xc8, 0xd5, 0x8f, 0xef, 0x1f, 0xa6, 0x4f, 0x79, 0xba, 0xfa, 0xf0, 0xe3,
	0xd1, 0xf1, 0xcf, 0xbb, 0x25, 0xe9, 0x18, 0x2f, 0x6a, 0xf6, 0x6b, 0x40, 0x6e, 0x9c, 0xdb, 0x61,
	0x35, 0x2c, 0xac, 0xa0, 0xf7, 0x49, 0xa0, 0xb9, 0x14, 0x38, 0xfd, 0xa0, 0xb8, 0xb9, 0x59, 0x27,
	0xe1, 0x76, 0x3a, 0x97, 0xc2, 0x8f, 0xdd, 0x1b, 0x5d, 0x2a, 0x91, 0x42, 0x1f, 0x93, 0xa1, 0x3f,
	0xa7, 0x56, 0xad, 0x44, 0xb4, 0x87, 0xfc, 0x78, 0xb3, 0x4e, 0x46, 0x3b, 0x3e, 0x42, 0x9d, 0xe8,
	0xd0, 0x77, 0xde, 0xa8, 0x95, 0xa0, 0xf7, 0xc8, 0x81, 0x03, 0xc7, 0x9b, 0x68, 0x1f, 0x45, 0xa3,
	0xcd, 0x3a, 0xb9, 0xd2, 0x8a, 0xb0, 0xcd, 0xca, 0x16, 0xa6, 0x27, 0x24, 0xf0, 0xff, 0x37, 0x0a,
	0xc6, 0xfb, 0x93, 0xf0, 0xe8, 0x7a, 0xd6, 0xd9, 0xcf, 0xb6, 0xa6, 0x8b, 0x6b, 0x3b, 0x7b, 0x9e,
	0xc8, 0x4a, 0xe4, 0xb3, 0xbf, 0x03, 0x12, 0x5f, 0x90, 0xed, 0x9d, 0xe1, 0x5a, 0x0b, 0x43, 0xef,
	0x90, 0xa0, 0x82, 0xba, 0x8b, 0xd8, 0x9b, 0xe1, 0xbb, 0xac, 0x44, 0x90, 0x3e, 0x21, 0xa1, 0x3f,
	0x9f, 0x7f, 0xd3, 0x0d, 0x57, 0x0b, 0x8c, 0x37, 0x2c, 0x6e, 0xed, 0x3e, 0x76, 0x0f, 0x64, 0x65,
	0x9f, 0xea, 0xd3, 0x09, 0x63, 0xc0, 0x60, 0xba, 0x61, 0x3f, 0x1d, 0xb6, 0x59, 0xd9, 0xc2, 0xb4,
	0x20, 0x41, 0xcd, 0x1d, 0x8f, 0x82, 0xf1, 0x60, 0x12, 0x1e, 0xdd, 0xde, 0xa5, 0xbb, 0xc0, 0x7a,
	0xdf, 0xa5, 0x17, 0xb1, 0x12, 0xb5, 0xb3, 0xcb, 0xf8, 0x5e, 0x8e, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x66, 0x9b, 0x9b, 0xa2, 0x08, 0x03, 0x00, 0x00,
}
