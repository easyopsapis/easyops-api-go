// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: up_insert_global_variables.proto

package pkg

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//UpInsertGlobalVariables请求
type UpInsertGlobalVariablesRequest struct {
	//
	//用户全局变量
	Var                  *UpInsertGlobalVariablesRequest_Var `protobuf:"bytes,1,opt,name=var,proto3" json:"var" form:"var"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *UpInsertGlobalVariablesRequest) Reset()         { *m = UpInsertGlobalVariablesRequest{} }
func (m *UpInsertGlobalVariablesRequest) String() string { return proto.CompactTextString(m) }
func (*UpInsertGlobalVariablesRequest) ProtoMessage()    {}
func (*UpInsertGlobalVariablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_092cbcb724330116, []int{0}
}
func (m *UpInsertGlobalVariablesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest.Unmarshal(m, b)
}
func (m *UpInsertGlobalVariablesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest.Marshal(b, m, deterministic)
}
func (m *UpInsertGlobalVariablesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpInsertGlobalVariablesRequest.Merge(m, src)
}
func (m *UpInsertGlobalVariablesRequest) XXX_Size() int {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest.Size(m)
}
func (m *UpInsertGlobalVariablesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpInsertGlobalVariablesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpInsertGlobalVariablesRequest proto.InternalMessageInfo

func (m *UpInsertGlobalVariablesRequest) GetVar() *UpInsertGlobalVariablesRequest_Var {
	if m != nil {
		return m.Var
	}
	return nil
}

type UpInsertGlobalVariablesRequest_Var struct {
	//
	//变量key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" form:"key"`
	//
	//变量Value
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	//
	//innerVarsRepeat
	InnerVarsRepeat bool `protobuf:"varint,3,opt,name=innerVarsRepeat,proto3" json:"innerVarsRepeat" form:"innerVarsRepeat"`
	//
	//diyVarsRepeat
	DiyVarsRepeat        bool     `protobuf:"varint,4,opt,name=diyVarsRepeat,proto3" json:"diyVarsRepeat" form:"diyVarsRepeat"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpInsertGlobalVariablesRequest_Var) Reset()         { *m = UpInsertGlobalVariablesRequest_Var{} }
func (m *UpInsertGlobalVariablesRequest_Var) String() string { return proto.CompactTextString(m) }
func (*UpInsertGlobalVariablesRequest_Var) ProtoMessage()    {}
func (*UpInsertGlobalVariablesRequest_Var) Descriptor() ([]byte, []int) {
	return fileDescriptor_092cbcb724330116, []int{0, 0}
}
func (m *UpInsertGlobalVariablesRequest_Var) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest_Var.Unmarshal(m, b)
}
func (m *UpInsertGlobalVariablesRequest_Var) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest_Var.Marshal(b, m, deterministic)
}
func (m *UpInsertGlobalVariablesRequest_Var) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpInsertGlobalVariablesRequest_Var.Merge(m, src)
}
func (m *UpInsertGlobalVariablesRequest_Var) XXX_Size() int {
	return xxx_messageInfo_UpInsertGlobalVariablesRequest_Var.Size(m)
}
func (m *UpInsertGlobalVariablesRequest_Var) XXX_DiscardUnknown() {
	xxx_messageInfo_UpInsertGlobalVariablesRequest_Var.DiscardUnknown(m)
}

var xxx_messageInfo_UpInsertGlobalVariablesRequest_Var proto.InternalMessageInfo

func (m *UpInsertGlobalVariablesRequest_Var) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpInsertGlobalVariablesRequest_Var) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *UpInsertGlobalVariablesRequest_Var) GetInnerVarsRepeat() bool {
	if m != nil {
		return m.InnerVarsRepeat
	}
	return false
}

func (m *UpInsertGlobalVariablesRequest_Var) GetDiyVarsRepeat() bool {
	if m != nil {
		return m.DiyVarsRepeat
	}
	return false
}

//
//UpInsertGlobalVariablesApi返回
type UpInsertGlobalVariablesResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpInsertGlobalVariablesResponseWrapper) Reset() {
	*m = UpInsertGlobalVariablesResponseWrapper{}
}
func (m *UpInsertGlobalVariablesResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpInsertGlobalVariablesResponseWrapper) ProtoMessage()    {}
func (*UpInsertGlobalVariablesResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_092cbcb724330116, []int{1}
}
func (m *UpInsertGlobalVariablesResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper.Unmarshal(m, b)
}
func (m *UpInsertGlobalVariablesResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpInsertGlobalVariablesResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper.Merge(m, src)
}
func (m *UpInsertGlobalVariablesResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper.Size(m)
}
func (m *UpInsertGlobalVariablesResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpInsertGlobalVariablesResponseWrapper proto.InternalMessageInfo

func (m *UpInsertGlobalVariablesResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpInsertGlobalVariablesResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpInsertGlobalVariablesResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpInsertGlobalVariablesResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpInsertGlobalVariablesRequest)(nil), "pkg.UpInsertGlobalVariablesRequest")
	proto.RegisterType((*UpInsertGlobalVariablesRequest_Var)(nil), "pkg.UpInsertGlobalVariablesRequest.Var")
	proto.RegisterType((*UpInsertGlobalVariablesResponseWrapper)(nil), "pkg.UpInsertGlobalVariablesResponseWrapper")
}

func init() { proto.RegisterFile("up_insert_global_variables.proto", fileDescriptor_092cbcb724330116) }

var fileDescriptor_092cbcb724330116 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0xe9, 0x76, 0x56, 0x34, 0x55, 0x57, 0x82, 0x0c, 0xa5, 0x82, 0x2d, 0x11, 0xd6, 0xbd,
	0x98, 0x85, 0xf5, 0x22, 0x0a, 0x1e, 0x06, 0x17, 0xd9, 0x6b, 0xc0, 0x7a, 0x1c, 0xd2, 0x9d, 0x77,
	0x63, 0x69, 0xa7, 0x89, 0xe9, 0x1f, 0xec, 0xe7, 0xf3, 0x33, 0x78, 0xed, 0x27, 0xf0, 0xd4, 0x4f,
	0x20, 0x7d, 0xb3, 0xc3, 0x74, 0x06, 0xdc, 0xd3, 0x4c, 0xf2, 0xfc, 0x9e, 0xe7, 0x6d, 0x1e, 0x5e,
	0x92, 0xb4, 0x66, 0x9d, 0x57, 0x35, 0xd8, 0x66, 0xad, 0x4a, 0x9d, 0xc9, 0x72, 0xdd, 0x49, 0x9b,
	0xcb, 0xac, 0x84, 0x9a, 0x1b, 0xab, 0x1b, 0x4d, 0x7d, 0x53, 0xa8, 0xe8, 0x9d, 0xca, 0x9b, 0x1f,
	0x6d, 0xc6, 0x6f, 0xf5, 0xf6, 0x52, 0x69, 0xa5, 0x2f, 0x51, 0xcb, 0xda, 0x3b, 0x3c, 0xe1, 0x01,
	0xff, 0x39, 0x4f, 0xf4, 0x4a, 0x69, 0xad, 0x4a, 0xd8, 0x53, 0xb0, 0x35, 0x4d, 0xef, 0x44, 0xf6,
	0xfb, 0x84, 0xbc, 0xfe, 0x66, 0x6e, 0x70, 0xe8, 0x57, 0x9c, 0x99, 0xee, 0x46, 0x0a, 0xf8, 0xd9,
	0x42, 0xdd, 0xd0, 0x1b, 0xe2, 0x77, 0xd2, 0x86, 0x5e, 0xe2, 0x5d, 0x04, 0x57, 0x6f, 0xb9, 0x29,
	0x14, 0x7f, 0xd8, 0xc1, 0x53, 0x69, 0x57, 0xcf, 0xc7, 0x21, 0x26, 0x77, 0xda, 0x6e, 0x3f, 0xb2,
	0x4e, 0x5a, 0x26, 0xa6, 0x8c, 0xe8, 0x8f, 0x47, 0xfc, 0x54, 0x5a, 0x9a, 0x10, 0xbf, 0x80, 0x1e,
	0x23, 0x9f, 0xcc, 0xc9, 0x02, 0x7a, 0x26, 0x26, 0x89, 0x9e, 0x93, 0xd3, 0x4e, 0x96, 0x2d, 0x84,
	0x27, 0xc8, 0xbc, 0x18, 0x87, 0xf8, 0xe9, 0x2e, 0xad, 0x6c, 0x81, 0x09, 0x27, 0xd3, 0x2f, 0xe4,
	0x2c, 0xaf, 0x2a, 0xb0, 0xa9, 0xb4, 0xb5, 0x00, 0x03, 0xb2, 0x09, 0xfd, 0xc4, 0xbb, 0x78, 0xbc,
	0x8a, 0xc6, 0x21, 0x5e, 0x3a, 0xc7, 0x11, 0xc0, 0xc4, 0xb1, 0x85, 0x7e, 0x26, 0xcf, 0x36, 0x79,
	0x3f, 0xcb, 0x58, 0x60, 0x46, 0x38, 0x0e, 0xf1, 0x4b, 0x97, 0x71, 0x20, 0x33, 0x71, 0x88, 0xb3,
	0xbf, 0x1e, 0x39, 0xff, 0x6f, 0x27, 0xb5, 0xd1, 0x55, 0x0d, 0xdf, 0xad, 0x34, 0x06, 0x2c, 0x7d,
	0x43, 0x16, 0xb7, 0x7a, 0x03, 0xf8, 0xf6, 0xd3, 0xd5, 0xd9, 0x38, 0xc4, 0x81, 0x9b, 0x30, 0xdd,
	0x32, 0x81, 0x22, 0xfd, 0x40, 0x82, 0xe9, 0xf7, 0xfa, 0x97, 0x29, 0x65, 0x5e, 0xdd, 0x77, 0xb0,
	0x1c, 0x87, 0x98, 0xee, 0xd9, 0x7b, 0x91, 0x89, 0x39, 0x3a, 0xf5, 0x06, 0xd6, 0x6a, 0x8b, 0x2d,
	0x1c, 0xf4, 0x86, 0xd7, 0x4c, 0x38, 0x99, 0x7e, 0x22, 0x8b, 0x8d, 0x6c, 0x24, 0x3e, 0x34, 0xb8,
	0x5a, 0x72, 0xb7, 0x23, 0x7c, 0xb7, 0x23, 0xfc, 0x7a, 0xda, 0x91, 0xf9, 0xe7, 0x4d, 0x34, 0x13,
	0x68, 0xca, 0x1e, 0x21, 0xf6, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x29, 0xe5, 0x19,
	0xb0, 0x02, 0x00, 0x00,
}
