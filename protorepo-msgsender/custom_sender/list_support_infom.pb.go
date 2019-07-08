// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_support_infom.proto

package custom_sender

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	msgsender "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/msgsender"
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
//ListSupportInformApi返回
type ListSupportInformResponseWrapper struct {
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
	Data                 []*msgsender.SupportInform `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListSupportInformResponseWrapper) Reset()         { *m = ListSupportInformResponseWrapper{} }
func (m *ListSupportInformResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListSupportInformResponseWrapper) ProtoMessage()    {}
func (*ListSupportInformResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_980446af05d1a1f3, []int{0}
}
func (m *ListSupportInformResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSupportInformResponseWrapper.Unmarshal(m, b)
}
func (m *ListSupportInformResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSupportInformResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListSupportInformResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSupportInformResponseWrapper.Merge(m, src)
}
func (m *ListSupportInformResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListSupportInformResponseWrapper.Size(m)
}
func (m *ListSupportInformResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSupportInformResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListSupportInformResponseWrapper proto.InternalMessageInfo

func (m *ListSupportInformResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListSupportInformResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListSupportInformResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListSupportInformResponseWrapper) GetData() []*msgsender.SupportInform {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListSupportInformResponseWrapper)(nil), "custom_sender.ListSupportInformResponseWrapper")
}

func init() { proto.RegisterFile("list_support_infom.proto", fileDescriptor_980446af05d1a1f3) }

var fileDescriptor_980446af05d1a1f3 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xe9, 0x7f, 0x77, 0xff, 0x60, 0x56, 0x51, 0x72, 0x90, 0xb0, 0x97, 0x2d, 0x11, 0x64,
	0x2f, 0x9b, 0x80, 0x5e, 0x44, 0xf0, 0xb2, 0xe0, 0x41, 0xf0, 0x54, 0x0f, 0x7a, 0x5b, 0xb2, 0x69,
	0xb6, 0x16, 0xda, 0x4e, 0xc8, 0xa4, 0xa0, 0x5f, 0xb6, 0x67, 0xcf, 0xfd, 0x04, 0xd2, 0xa4, 0x68,
	0x3d, 0x25, 0x6f, 0xde, 0xef, 0x31, 0x6f, 0x08, 0xab, 0x4a, 0xf4, 0x7b, 0x6c, 0xad, 0x05, 0xe7,
	0xf7, 0x65, 0x73, 0x84, 0x5a, 0x58, 0x07, 0x1e, 0xe8, 0x99, 0x6e, 0xd1, 0x43, 0xbd, 0x47, 0xd3,
	0xe4, 0xc6, 0xad, 0xb6, 0x45, 0xe9, 0xdf, 0xdb, 0x83, 0xd0, 0x50, 0xcb, 0x02, 0x0a, 0x90, 0x81,
	0x3a, 0xb4, 0xc7, 0xa0, 0x82, 0x08, 0xbf, 0x98, 0x5e, 0xbd, 0x15, 0x20, 0x8c, 0xc2, 0x4f, 0xb0,
	0x28, 0x2a, 0xd0, 0xaa, 0x92, 0x1a, 0x1a, 0xef, 0x94, 0xf6, 0x18, 0x93, 0xce, 0x58, 0xd8, 0xd6,
	0x90, 0x9b, 0x0a, 0xe5, 0x08, 0xca, 0x20, 0x65, 0x8d, 0x45, 0xdc, 0x29, 0xa7, 0xb5, 0xdc, 0xd8,
	0x8b, 0x7f, 0x25, 0x24, 0x7d, 0x2e, 0xd1, 0xbf, 0x44, 0xf3, 0x29, 0x78, 0x99, 0x41, 0x0b, 0x0d,
	0x9a, 0x57, 0xa7, 0xac, 0x35, 0x8e, 0x5e, 0x91, 0xb9, 0x86, 0xdc, 0xb0, 0x24, 0x4d, 0x36, 0x8b,
	0xdd, 0x79, 0xdf, 0xad, 0x97, 0x03, 0x76, 0xcf, 0x87, 0x29, 0xcf, 0x82, 0x49, 0xef, 0xc8, 0x72,
	0x78, 0x1f, 0x3f, 0x6c, 0xa5, 0xca, 0x86, 0xfd, 0x4b, 0x93, 0xcd, 0xc9, 0xee, 0xb2, 0xef, 0xd6,
	0xf4, 0x97, 0x1d, 0x4d, 0x9e, 0x4d, 0x51, 0x7a, 0x4d, 0x16, 0xc6, 0x39, 0x70, 0x6c, 0x16, 0x32,
	0x17, 0x7d, 0xb7, 0x3e, 0x8d, 0x99, 0x30, 0xe6, 0x59, 0xb4, 0xe9, 0x03, 0x99, 0xe7, 0xca, 0x2b,
	0x36, 0x4f, 0x67, 0x9b, 0xe5, 0x0d, 0x13, 0x3f, 0xa7, 0x89, 0x3f, 0xed, 0xa7, 0x05, 0x07, 0x9e,
	0x67, 0x21, 0x76, 0xf8, 0x1f, 0x2e, 0xbe, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xe5, 0x55,
	0x4f, 0xa5, 0x01, 0x00, 0x00,
}