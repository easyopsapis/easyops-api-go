// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_one_subsystem_tree.proto

package custom

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	cmdb_extend "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
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
//GetOneSubSystemTree请求
type GetOneSubSystemTreeRequest struct {
	//
	//系统简称
	Abbreviation         string   `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation" form:"abbreviation"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneSubSystemTreeRequest) Reset()         { *m = GetOneSubSystemTreeRequest{} }
func (m *GetOneSubSystemTreeRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneSubSystemTreeRequest) ProtoMessage()    {}
func (*GetOneSubSystemTreeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee58ba75402e1b1, []int{0}
}
func (m *GetOneSubSystemTreeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneSubSystemTreeRequest.Unmarshal(m, b)
}
func (m *GetOneSubSystemTreeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneSubSystemTreeRequest.Marshal(b, m, deterministic)
}
func (m *GetOneSubSystemTreeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneSubSystemTreeRequest.Merge(m, src)
}
func (m *GetOneSubSystemTreeRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneSubSystemTreeRequest.Size(m)
}
func (m *GetOneSubSystemTreeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneSubSystemTreeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneSubSystemTreeRequest proto.InternalMessageInfo

func (m *GetOneSubSystemTreeRequest) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

//
//GetOneSubSystemTreeApi返回
type GetOneSubSystemTreeResponseWrapper struct {
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
	Data                 *cmdb_extend.SubsystemDependency `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetOneSubSystemTreeResponseWrapper) Reset()         { *m = GetOneSubSystemTreeResponseWrapper{} }
func (m *GetOneSubSystemTreeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetOneSubSystemTreeResponseWrapper) ProtoMessage()    {}
func (*GetOneSubSystemTreeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee58ba75402e1b1, []int{1}
}
func (m *GetOneSubSystemTreeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneSubSystemTreeResponseWrapper.Unmarshal(m, b)
}
func (m *GetOneSubSystemTreeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneSubSystemTreeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetOneSubSystemTreeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneSubSystemTreeResponseWrapper.Merge(m, src)
}
func (m *GetOneSubSystemTreeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetOneSubSystemTreeResponseWrapper.Size(m)
}
func (m *GetOneSubSystemTreeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneSubSystemTreeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneSubSystemTreeResponseWrapper proto.InternalMessageInfo

func (m *GetOneSubSystemTreeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetOneSubSystemTreeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetOneSubSystemTreeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetOneSubSystemTreeResponseWrapper) GetData() *cmdb_extend.SubsystemDependency {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOneSubSystemTreeRequest)(nil), "custom.GetOneSubSystemTreeRequest")
	proto.RegisterType((*GetOneSubSystemTreeResponseWrapper)(nil), "custom.GetOneSubSystemTreeResponseWrapper")
}

func init() { proto.RegisterFile("get_one_subsystem_tree.proto", fileDescriptor_4ee58ba75402e1b1) }

var fileDescriptor_4ee58ba75402e1b1 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0xdd, 0xaa, 0xd3, 0x40,
	0x10, 0x26, 0xda, 0x73, 0xc0, 0xed, 0x01, 0x25, 0x82, 0x86, 0x22, 0xa4, 0xac, 0x20, 0xe7, 0x26,
	0x59, 0x50, 0x10, 0xd1, 0xbb, 0x83, 0xc5, 0x4b, 0x21, 0x15, 0x44, 0x6f, 0xe2, 0xfe, 0x4c, 0x63,
	0x30, 0xd9, 0x59, 0x77, 0x37, 0xfd, 0x79, 0xd9, 0x3c, 0x44, 0x9e, 0x40, 0xba, 0xa9, 0x6d, 0x44,
	0x2f, 0xbd, 0xda, 0x99, 0xf9, 0x7e, 0x76, 0xe6, 0x23, 0xcf, 0x2a, 0xf0, 0x25, 0x6a, 0x28, 0x5d,
	0x27, 0xdc, 0xc1, 0x79, 0x68, 0x4b, 0x6f, 0x01, 0x72, 0x63, 0xd1, 0x63, 0x7c, 0x2d, 0x3b, 0xe7,
	0xb1, 0x5d, 0x64, 0x55, 0xed, 0xbf, 0x77, 0x22, 0x97, 0xd8, 0xb2, 0x0a, 0x2b, 0x64, 0x01, 0x16,
	0xdd, 0x26, 0x74, 0xa1, 0x09, 0xd5, 0x28, 0x5b, 0x7c, 0xad, 0x30, 0x07, 0xee, 0x0e, 0x68, 0x5c,
	0xde, 0xa0, 0xe4, 0x0d, 0x93, 0xa8, 0xbd, 0xe5, 0xd2, 0xbb, 0x51, 0x69, 0xc1, 0x60, 0xd6, 0xa2,
	0x82, 0xc6, 0xb1, 0x13, 0x91, 0x85, 0x96, 0xc9, 0x56, 0x89, 0x12, 0xf6, 0x1e, 0xb4, 0x62, 0xdc,
	0x98, 0x52, 0x81, 0x01, 0xad, 0x40, 0xcb, 0xc3, 0xc9, 0xfb, 0xdb, 0xff, 0xf1, 0xbe, 0x9c, 0xfb,
	0xd7, 0x0f, 0xaf, 0x27, 0xc7, 0xb6, 0xbb, 0xda, 0xff, 0xc0, 0x1d, 0xab, 0x30, 0x0b, 0x60, 0xb6,
	0xe5, 0x4d, 0xad, 0xb8, 0x47, 0xeb, 0xd8, 0xb9, 0x1c, 0x75, 0xf4, 0x0b, 0x59, 0x7c, 0x00, 0xff,
	0x51, 0xc3, 0xba, 0x13, 0xeb, 0xe0, 0xfd, 0xc9, 0x02, 0x14, 0xf0, 0xb3, 0x03, 0xe7, 0xe3, 0x77,
	0xe4, 0x86, 0x0b, 0x61, 0x61, 0x5b, 0x73, 0x5f, 0xa3, 0x4e, 0xa2, 0x65, 0x74, 0xfb, 0xe0, 0xee,
	0xe9, 0xd0, 0xa7, 0x8f, 0x37, 0x68, 0xdb, 0xb7, 0x74, 0x8a, 0xd2, 0xe2, 0x0f, 0x32, 0x1d, 0x22,
	0x42, 0xff, 0xe9, 0xed, 0x0c, 0x6a, 0x07, 0x9f, 0x2d, 0x37, 0x06, 0x6c, 0xfc, 0x9c, 0xcc, 0x24,
	0x2a, 0x08, 0xde, 0x57, 0x77, 0x0f, 0x87, 0x3e, 0x9d, 0x8f, 0xde, 0xc7, 0x29, 0x2d, 0x02, 0x18,
	0xbf, 0x21, 0xf3, 0xe3, 0xbb, 0xda, 0x9b, 0x86, 0xd7, 0x3a, 0xb9, 0x17, 0xf6, 0x78, 0x32, 0xf4,
	0x69, 0x7c, 0xe1, 0x9e, 0x40, 0x5a, 0x4c, 0xa9, 0xf1, 0x0b, 0x72, 0x05, 0xd6, 0xa2, 0x4d, 0xee,
	0x07, 0xcd, 0xa3, 0xa1, 0x4f, 0x6f, 0x46, 0x4d, 0x18, 0xd3, 0x62, 0x84, 0xe3, 0x15, 0x99, 0x29,
	0xee, 0x79, 0x32, 0x5b, 0x46, 0xb7, 0xf3, 0x97, 0xcb, 0x7c, 0x92, 0x7b, 0xbe, 0xfe, 0x9d, 0xfb,
	0xfb, 0x73, 0xec, 0xd3, 0x45, 0x8f, 0x3a, 0x5a, 0x04, 0xb9, 0xb8, 0x0e, 0xb1, 0xbe, 0xfa, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x3c, 0x9b, 0x9c, 0x4e, 0xa3, 0x02, 0x00, 0x00,
}
