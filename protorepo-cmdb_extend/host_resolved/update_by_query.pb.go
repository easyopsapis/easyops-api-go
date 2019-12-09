// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_by_query.proto

package host_resolved

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
//UpdateByQuery请求
type UpdateByQueryRequest struct {
	//
	//更新条件
	Query *types.Struct `protobuf:"bytes,1,opt,name=query,proto3" json:"query" form:"query"`
	//
	//文档数据内容
	Update               *types.Struct `protobuf:"bytes,2,opt,name=update,proto3" json:"update" form:"update"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateByQueryRequest) Reset()         { *m = UpdateByQueryRequest{} }
func (m *UpdateByQueryRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateByQueryRequest) ProtoMessage()    {}
func (*UpdateByQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb2c8f8af00a4263, []int{0}
}
func (m *UpdateByQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateByQueryRequest.Unmarshal(m, b)
}
func (m *UpdateByQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateByQueryRequest.Marshal(b, m, deterministic)
}
func (m *UpdateByQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateByQueryRequest.Merge(m, src)
}
func (m *UpdateByQueryRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateByQueryRequest.Size(m)
}
func (m *UpdateByQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateByQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateByQueryRequest proto.InternalMessageInfo

func (m *UpdateByQueryRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *UpdateByQueryRequest) GetUpdate() *types.Struct {
	if m != nil {
		return m.Update
	}
	return nil
}

//
//UpdateByQueryApi返回
type UpdateByQueryResponseWrapper struct {
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

func (m *UpdateByQueryResponseWrapper) Reset()         { *m = UpdateByQueryResponseWrapper{} }
func (m *UpdateByQueryResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateByQueryResponseWrapper) ProtoMessage()    {}
func (*UpdateByQueryResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb2c8f8af00a4263, []int{1}
}
func (m *UpdateByQueryResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateByQueryResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateByQueryResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateByQueryResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateByQueryResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateByQueryResponseWrapper.Merge(m, src)
}
func (m *UpdateByQueryResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateByQueryResponseWrapper.Size(m)
}
func (m *UpdateByQueryResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateByQueryResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateByQueryResponseWrapper proto.InternalMessageInfo

func (m *UpdateByQueryResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateByQueryResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateByQueryResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateByQueryResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateByQueryRequest)(nil), "host_resolved.UpdateByQueryRequest")
	proto.RegisterType((*UpdateByQueryResponseWrapper)(nil), "host_resolved.UpdateByQueryResponseWrapper")
}

func init() { proto.RegisterFile("update_by_query.proto", fileDescriptor_eb2c8f8af00a4263) }

var fileDescriptor_eb2c8f8af00a4263 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xda, 0x16, 0xdc, 0x5a, 0xd4, 0x45, 0x6b, 0xa8, 0x85, 0xc8, 0x0a, 0xe2, 0xc5, 0x14,
	0xf4, 0x22, 0x7a, 0x10, 0x02, 0x7d, 0x00, 0x57, 0xc4, 0x63, 0x49, 0x9a, 0x69, 0x5a, 0x68, 0x3b,
	0xdb, 0xcd, 0xae, 0xd8, 0x67, 0xf1, 0xdd, 0x72, 0xf0, 0x11, 0xf2, 0x04, 0x92, 0xd9, 0x48, 0x4b,
	0x05, 0x4f, 0xc9, 0xcc, 0xf7, 0x33, 0xdf, 0xb7, 0xec, 0xcc, 0xaa, 0x34, 0x36, 0x30, 0x4a, 0xd6,
	0xa3, 0x95, 0x05, 0xbd, 0x0e, 0x95, 0x46, 0x83, 0xbc, 0x33, 0xc5, 0xdc, 0x8c, 0x34, 0xe4, 0x38,
	0xff, 0x80, 0xb4, 0x77, 0x9b, 0xcd, 0xcc, 0xd4, 0x26, 0xe1, 0x18, 0x17, 0x83, 0x0c, 0x33, 0x1c,
	0x10, 0x2b, 0xb1, 0x13, 0x9a, 0x68, 0xa0, 0x3f, 0xa7, 0xee, 0xf5, 0x33, 0xc4, 0x6c, 0x0e, 0x1b,
	0x56, 0x6e, 0xb4, 0x1d, 0x9b, 0x1a, 0xbd, 0xd8, 0x45, 0x61, 0xa1, 0x4c, 0x7d, 0x58, 0x7c, 0x79,
	0xec, 0xf4, 0x8d, 0x22, 0x45, 0xeb, 0x97, 0x2a, 0x90, 0x84, 0x95, 0x85, 0xdc, 0xf0, 0x67, 0xd6,
	0xa4, 0x80, 0xbe, 0x77, 0xe9, 0xdd, 0xb4, 0xef, 0xce, 0x43, 0xe7, 0x12, 0xfe, 0xba, 0x84, 0xaf,
	0x74, 0x23, 0x3a, 0x2e, 0x8b, 0xe0, 0x70, 0x82, 0x7a, 0xf1, 0x28, 0x88, 0x2f, 0xa4, 0xd3, 0xf1,
	0x88, 0xb5, 0x5c, 0x57, 0x7f, 0xef, 0x7f, 0x87, 0x93, 0xb2, 0x08, 0x3a, 0xce, 0xc1, 0x09, 0x84,
	0xac, 0x95, 0xe2, 0xdb, 0x63, 0xfd, 0x9d, 0x74, 0xb9, 0xc2, 0x65, 0x0e, 0xef, 0x3a, 0x56, 0x0a,
	0x34, 0xbf, 0x62, 0x8d, 0x31, 0xa6, 0x40, 0x21, 0x9b, 0xd1, 0x51, 0x59, 0x04, 0x6d, 0xe7, 0x54,
	0x6d, 0x85, 0x24, 0x90, 0x3f, 0xb0, 0x76, 0xf5, 0x1d, 0x7e, 0xaa, 0x79, 0x3c, 0x5b, 0x52, 0x9c,
	0x83, 0xa8, 0x5b, 0x16, 0x01, 0xdf, 0x70, 0x6b, 0x50, 0xc8, 0x6d, 0x2a, 0xbf, 0x66, 0x4d, 0xd0,
	0x1a, 0xb5, 0xbf, 0x4f, 0x9a, 0xad, 0xae, 0xb4, 0x16, 0xd2, 0xc1, 0xfc, 0x89, 0x35, 0xd2, 0xd8,
	0xc4, 0x7e, 0x83, 0x9a, 0x76, 0xff, 0x34, 0x1d, 0x56, 0x2f, 0xbe, 0x1d, 0xaf, 0x62, 0x0b, 0x49,
	0xa2, 0xa4, 0x45, 0xb4, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6c, 0x42, 0xf6, 0x1b,
	0x02, 0x00, 0x00,
}
