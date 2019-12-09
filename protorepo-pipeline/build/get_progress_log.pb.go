// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_progress_log.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//GetProgressLog请求
type GetProgressLogRequest struct {
	//
	//实时id，服务端自动生成
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProgressLogRequest) Reset()         { *m = GetProgressLogRequest{} }
func (m *GetProgressLogRequest) String() string { return proto.CompactTextString(m) }
func (*GetProgressLogRequest) ProtoMessage()    {}
func (*GetProgressLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd47666fdf2fc0ad, []int{0}
}
func (m *GetProgressLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProgressLogRequest.Unmarshal(m, b)
}
func (m *GetProgressLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProgressLogRequest.Marshal(b, m, deterministic)
}
func (m *GetProgressLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProgressLogRequest.Merge(m, src)
}
func (m *GetProgressLogRequest) XXX_Size() int {
	return xxx_messageInfo_GetProgressLogRequest.Size(m)
}
func (m *GetProgressLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProgressLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProgressLogRequest proto.InternalMessageInfo

func (m *GetProgressLogRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//GetProgressLogApi返回
type GetProgressLogResponseWrapper struct {
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
	Data                 *pipeline.ProgressLog `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetProgressLogResponseWrapper) Reset()         { *m = GetProgressLogResponseWrapper{} }
func (m *GetProgressLogResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetProgressLogResponseWrapper) ProtoMessage()    {}
func (*GetProgressLogResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd47666fdf2fc0ad, []int{1}
}
func (m *GetProgressLogResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProgressLogResponseWrapper.Unmarshal(m, b)
}
func (m *GetProgressLogResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProgressLogResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetProgressLogResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProgressLogResponseWrapper.Merge(m, src)
}
func (m *GetProgressLogResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetProgressLogResponseWrapper.Size(m)
}
func (m *GetProgressLogResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProgressLogResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetProgressLogResponseWrapper proto.InternalMessageInfo

func (m *GetProgressLogResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetProgressLogResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetProgressLogResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetProgressLogResponseWrapper) GetData() *pipeline.ProgressLog {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetProgressLogRequest)(nil), "build.GetProgressLogRequest")
	proto.RegisterType((*GetProgressLogResponseWrapper)(nil), "build.GetProgressLogResponseWrapper")
}

func init() { proto.RegisterFile("get_progress_log.proto", fileDescriptor_cd47666fdf2fc0ad) }

var fileDescriptor_cd47666fdf2fc0ad = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xe9, 0xdc, 0x84, 0x65, 0x8a, 0x52, 0xd8, 0x28, 0x83, 0xd1, 0x11, 0x41, 0xf6, 0xb2,
	0x06, 0x14, 0x44, 0xf6, 0x38, 0x10, 0x5f, 0x7c, 0x90, 0x82, 0xf8, 0x38, 0xd2, 0xe6, 0x2e, 0x06,
	0xb2, 0xde, 0x98, 0xa4, 0xa0, 0x7f, 0xb6, 0x0f, 0xfe, 0x84, 0xfe, 0x02, 0x59, 0x3a, 0x71, 0xe8,
	0x53, 0x7b, 0xee, 0x39, 0xdf, 0xcd, 0xe1, 0x92, 0x89, 0x04, 0xbf, 0x31, 0x16, 0xa5, 0x05, 0xe7,
	0x36, 0x1a, 0x65, 0x66, 0x2c, 0x7a, 0x8c, 0x07, 0x45, 0xad, 0xb4, 0x98, 0x2e, 0xa5, 0xf2, 0x6f,
	0x75, 0x91, 0x95, 0xb8, 0x63, 0x12, 0x25, 0xb2, 0xe0, 0x16, 0xf5, 0x36, 0xa8, 0x20, 0xc2, 0x5f,
	0x47, 0x4d, 0x5f, 0x24, 0x66, 0xc0, 0xdd, 0x27, 0x1a, 0x97, 0x69, 0x2c, 0xb9, 0x66, 0x25, 0x56,
	0xde, 0xf2, 0xd2, 0xbb, 0x8e, 0xb4, 0x60, 0x70, 0xb9, 0x43, 0x01, 0xda, 0xb1, 0x43, 0x90, 0x05,
	0xc9, 0x8c, 0x32, 0xa0, 0x55, 0x05, 0xec, 0x7f, 0x19, 0x7a, 0x47, 0xc6, 0x8f, 0xe0, 0x9f, 0x0f,
	0xc6, 0x13, 0xca, 0x1c, 0xde, 0x6b, 0x70, 0x3e, 0x9e, 0x91, 0x9e, 0x12, 0x49, 0x34, 0x8f, 0x16,
	0xc3, 0xf5, 0x79, 0xdb, 0xa4, 0xc3, 0x2d, 0xda, 0xdd, 0x8a, 0x2a, 0x41, 0xf3, 0x9e, 0x12, 0xf4,
	0x2b, 0x22, 0xb3, 0xbf, 0xa0, 0x33, 0x58, 0x39, 0x78, 0xb5, 0xdc, 0x18, 0xb0, 0xf1, 0x15, 0xe9,
	0x97, 0x28, 0x20, 0xac, 0x18, 0xac, 0x2f, 0xda, 0x26, 0x1d, 0x75, 0x2b, 0xf6, 0x53, 0x9a, 0x07,
	0x33, 0xbe, 0x27, 0xa3, 0xfd, 0xf7, 0xe1, 0xc3, 0x68, 0xae, 0xaa, 0xa4, 0x17, 0x9e, 0x9b, 0xb4,
	0x4d, 0x1a, 0xff, 0x66, 0x0f, 0x26, 0xcd, 0x8f, 0xa3, 0xf1, 0x35, 0x19, 0x80, 0xb5, 0x68, 0x93,
	0x93, 0xc0, 0x5c, 0xb6, 0x4d, 0x7a, 0xd6, 0x31, 0x61, 0x4c, 0xf3, 0xce, 0x8e, 0x57, 0xa4, 0x2f,
	0xb8, 0xe7, 0x49, 0x7f, 0x1e, 0x2d, 0x46, 0x37, 0xe3, 0xec, 0xe7, 0x18, 0xd9, 0x51, 0xf5, 0xe3,
	0x76, 0xfb, 0x30, 0xcd, 0x03, 0x53, 0x9c, 0x86, 0x1b, 0xdd, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xd3, 0xd9, 0x9e, 0x2d, 0xca, 0x01, 0x00, 0x00,
}
