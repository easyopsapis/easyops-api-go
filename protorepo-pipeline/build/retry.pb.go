// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: retry.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//Retry请求
type RetryRequest struct {
	//
	//构建任务id
	BuildId              string   `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id" form:"build_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryRequest) Reset()         { *m = RetryRequest{} }
func (m *RetryRequest) String() string { return proto.CompactTextString(m) }
func (*RetryRequest) ProtoMessage()    {}
func (*RetryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fd052cddd5591c, []int{0}
}
func (m *RetryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryRequest.Unmarshal(m, b)
}
func (m *RetryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryRequest.Marshal(b, m, deterministic)
}
func (m *RetryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryRequest.Merge(m, src)
}
func (m *RetryRequest) XXX_Size() int {
	return xxx_messageInfo_RetryRequest.Size(m)
}
func (m *RetryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetryRequest proto.InternalMessageInfo

func (m *RetryRequest) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

//
//Retry返回
type RetryResponse struct {
	//
	//build id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryResponse) Reset()         { *m = RetryResponse{} }
func (m *RetryResponse) String() string { return proto.CompactTextString(m) }
func (*RetryResponse) ProtoMessage()    {}
func (*RetryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fd052cddd5591c, []int{1}
}
func (m *RetryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryResponse.Unmarshal(m, b)
}
func (m *RetryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryResponse.Marshal(b, m, deterministic)
}
func (m *RetryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryResponse.Merge(m, src)
}
func (m *RetryResponse) XXX_Size() int {
	return xxx_messageInfo_RetryResponse.Size(m)
}
func (m *RetryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetryResponse proto.InternalMessageInfo

func (m *RetryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//RetryApi返回
type RetryResponseWrapper struct {
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
	Data                 *RetryResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RetryResponseWrapper) Reset()         { *m = RetryResponseWrapper{} }
func (m *RetryResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*RetryResponseWrapper) ProtoMessage()    {}
func (*RetryResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fd052cddd5591c, []int{2}
}
func (m *RetryResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryResponseWrapper.Unmarshal(m, b)
}
func (m *RetryResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryResponseWrapper.Marshal(b, m, deterministic)
}
func (m *RetryResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryResponseWrapper.Merge(m, src)
}
func (m *RetryResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_RetryResponseWrapper.Size(m)
}
func (m *RetryResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RetryResponseWrapper proto.InternalMessageInfo

func (m *RetryResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RetryResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *RetryResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RetryResponseWrapper) GetData() *RetryResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RetryRequest)(nil), "build.RetryRequest")
	proto.RegisterType((*RetryResponse)(nil), "build.RetryResponse")
	proto.RegisterType((*RetryResponseWrapper)(nil), "build.RetryResponseWrapper")
}

func init() { proto.RegisterFile("retry.proto", fileDescriptor_a7fd052cddd5591c) }

var fileDescriptor_a7fd052cddd5591c = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0x04, 0x95, 0x2d, 0x0a, 0xd9, 0x10, 0xd3, 0x70, 0x29, 0xae, 0x46, 0xb9, 0xb4,
	0x15, 0x49, 0x8c, 0x78, 0x6c, 0xe2, 0xc1, 0xeb, 0x5e, 0x3c, 0x18, 0x35, 0x2d, 0xbb, 0xd4, 0x8d,
	0xc0, 0xd6, 0xed, 0x56, 0xfc, 0x13, 0xdf, 0xd3, 0x53, 0x4d, 0x7c, 0x84, 0x3e, 0x81, 0xe9, 0x14,
	0x11, 0x0e, 0x9e, 0x3a, 0x33, 0xdf, 0xf7, 0xfd, 0x66, 0xba, 0xc8, 0x54, 0x5c, 0xab, 0x57, 0x37,
	0x56, 0x52, 0x4b, 0x5c, 0x0b, 0x53, 0x31, 0x61, 0x1d, 0x27, 0x12, 0xfa, 0x21, 0x0d, 0xdd, 0x91,
	0x9c, 0x7a, 0x91, 0x8c, 0xa4, 0x07, 0x6a, 0x98, 0x8e, 0xa1, 0x83, 0x06, 0xaa, 0x32, 0xd5, 0x39,
	0x5b, 0xb1, 0x4f, 0xe7, 0x42, 0x3f, 0xca, 0xb9, 0x17, 0x49, 0x07, 0x44, 0xe7, 0x39, 0x98, 0x08,
	0x16, 0x68, 0xa9, 0x12, 0x6f, 0x59, 0x96, 0x39, 0x42, 0x51, 0x83, 0x16, 0xcb, 0x29, 0x7f, 0x4a,
	0x79, 0xa2, 0xb1, 0x8f, 0xb6, 0x61, 0xff, 0xbd, 0x60, 0x96, 0xd1, 0x35, 0x7a, 0x75, 0xff, 0x38,
	0xcf, 0xec, 0xe6, 0x58, 0xaa, 0xe9, 0x05, 0xf9, 0x55, 0xc8, 0xf7, 0x97, 0xdd, 0x42, 0xbb, 0x77,
	0x37, 0x27, 0xce, 0x30, 0x70, 0xde, 0x6e, 0xdf, 0xfb, 0x83, 0x8f, 0x43, 0xba, 0x05, 0xf2, 0x15,
	0x23, 0x3e, 0xda, 0x59, 0x30, 0x93, 0x58, 0xce, 0x12, 0x8e, 0xfb, 0xa8, 0xb2, 0xc4, 0xed, 0xe7,
	0x99, 0x5d, 0x2f, 0x71, 0xff, 0x81, 0x2a, 0x82, 0x91, 0x4f, 0x03, 0xb5, 0xd7, 0x20, 0xd7, 0x2a,
	0x88, 0x63, 0xae, 0xf0, 0x01, 0xaa, 0x8e, 0x24, 0xe3, 0x40, 0xab, 0xf9, 0xcd, 0x3c, 0xb3, 0xcd,
	0x92, 0x56, 0x4c, 0x09, 0x05, 0x11, 0x9f, 0x23, 0xb3, 0xf8, 0x5e, 0xbe, 0xc4, 0x93, 0x40, 0xcc,
	0xac, 0x0a, 0x6c, 0xde, 0xcb, 0x33, 0x1b, 0xff, 0x79, 0x17, 0x22, 0xa1, 0xab, 0x56, 0x7c, 0x84,
	0x6a, 0x5c, 0x29, 0xa9, 0xac, 0x0d, 0xc8, 0xb4, 0xf2, 0xcc, 0x6e, 0x94, 0x19, 0x18, 0x13, 0x5a,
	0xca, 0x78, 0x88, 0xaa, 0x2c, 0xd0, 0x81, 0x55, 0xed, 0x1a, 0x3d, 0xf3, 0xb4, 0xed, 0xc2, 0xbf,
	0xbb, 0x6b, 0x17, 0xaf, 0x1e, 0x57, 0x78, 0x09, 0x85, 0x48, 0xb8, 0x09, 0x2f, 0x3f, 0xf8, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xc5, 0x4c, 0x8e, 0xf6, 0x01, 0x00, 0x00,
}