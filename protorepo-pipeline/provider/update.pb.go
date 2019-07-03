// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

package provider

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//UpdateApi返回
type UpdateResponseWrapper struct {
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
	Data                 *pipeline.Provider `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateResponseWrapper) Reset()         { *m = UpdateResponseWrapper{} }
func (m *UpdateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseWrapper) ProtoMessage()    {}
func (*UpdateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{0}
}
func (m *UpdateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseWrapper.Merge(m, src)
}
func (m *UpdateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseWrapper.Size(m)
}
func (m *UpdateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseWrapper proto.InternalMessageInfo

func (m *UpdateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateResponseWrapper) GetData() *pipeline.Provider {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateResponseWrapper)(nil), "provider.UpdateResponseWrapper")
}

func init() { proto.RegisterFile("update.proto", fileDescriptor_3f0fa214029f1c21) }

var fileDescriptor_3f0fa214029f1c21 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xb6, 0xa2, 0x69, 0x41, 0x59, 0x50, 0x42, 0x2f, 0x2d, 0x11, 0xa4, 0x97, 0x64,
	0x41, 0x41, 0xc5, 0x63, 0xc1, 0xbb, 0x06, 0xc4, 0xf3, 0x36, 0x3b, 0x8d, 0x8b, 0x9b, 0xcc, 0x32,
	0xbb, 0x6d, 0xf5, 0x49, 0xbd, 0xf5, 0x21, 0xfa, 0x04, 0xd2, 0x49, 0xaa, 0x3d, 0x65, 0xfe, 0xf9,
	0xff, 0x2f, 0x3b, 0x33, 0xf1, 0x70, 0xe9, 0xb4, 0x0a, 0x90, 0x3b, 0xc2, 0x80, 0xe2, 0xd4, 0x11,
	0xae, 0x8c, 0x06, 0x1a, 0x65, 0x95, 0x09, 0x1f, 0xcb, 0x79, 0x5e, 0x62, 0x2d, 0x2b, 0xac, 0x50,
	0x72, 0x60, 0xbe, 0x5c, 0xb0, 0x62, 0xc1, 0x55, 0x0b, 0x8e, 0x5e, 0x2b, 0xcc, 0x41, 0xf9, 0x6f,
	0x74, 0x3e, 0xb7, 0x58, 0x2a, 0x2b, 0x4b, 0x6c, 0x02, 0xa9, 0x32, 0xf8, 0x96, 0x24, 0x70, 0x98,
	0xd5, 0xa8, 0xc1, 0x7a, 0xd9, 0x05, 0x25, 0x4b, 0xe9, 0x8c, 0x03, 0x6b, 0x1a, 0x90, 0xfb, 0xb7,
	0xbb, 0x5f, 0xde, 0x1f, 0x4c, 0x50, 0xaf, 0x4d, 0xf8, 0xc4, 0xb5, 0xac, 0x30, 0x63, 0x33, 0x5b,
	0x29, 0x6b, 0xb4, 0x0a, 0x48, 0x5e, 0xfe, 0x95, 0x2d, 0x97, 0xfe, 0x44, 0xf1, 0xe5, 0x1b, 0x2f,
	0x55, 0x80, 0x77, 0xd8, 0x78, 0x78, 0x27, 0xe5, 0x1c, 0x90, 0xb8, 0x8e, 0x7b, 0x25, 0x6a, 0x48,
	0xa2, 0x49, 0x34, 0xed, 0xcf, 0xce, 0xb7, 0x9b, 0xf1, 0x60, 0x81, 0x54, 0x3f, 0xa5, 0xbb, 0x6e,
	0x5a, 0xb0, 0x29, 0x1e, 0xe3, 0xc1, 0xee, 0xfb, 0xfc, 0xe5, 0xac, 0x32, 0x4d, 0x72, 0x34, 0x89,
	0xa6, 0x67, 0xb3, 0xab, 0xed, 0x66, 0x2c, 0xfe, 0xb3, 0x9d, 0x99, 0x16, 0x87, 0x51, 0x71, 0x13,
	0xf7, 0x81, 0x08, 0x29, 0x39, 0x66, 0xe6, 0x62, 0xbb, 0x19, 0x0f, 0x5b, 0x86, 0xdb, 0x69, 0xd1,
	0xda, 0xe2, 0x21, 0xee, 0x69, 0x15, 0x54, 0xd2, 0x9b, 0x44, 0xd3, 0xc1, 0xad, 0xc8, 0xf7, 0x07,
	0xc8, 0x5f, 0xba, 0x03, 0x1c, 0x8e, 0xb6, 0x4b, 0xa6, 0x05, 0x03, 0xf3, 0x13, 0x5e, 0xf0, 0xee,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x58, 0xc2, 0xaa, 0x6c, 0xb4, 0x01, 0x00, 0x00,
}
