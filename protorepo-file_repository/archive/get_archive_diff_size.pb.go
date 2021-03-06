// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_archive_diff_size.proto

package archive

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//GetDiffSize请求
type GetDiffSizeRequest struct {
	//
	//版本Id
	VerFrom string `protobuf:"bytes,1,opt,name=ver_from,json=verFrom,proto3" json:"ver_from" form:"ver_from"`
	//
	//版本Id
	VerTo string `protobuf:"bytes,2,opt,name=ver_to,json=verTo,proto3" json:"ver_to" form:"ver_to"`
	//
	//程序包Id
	PackageId            string   `protobuf:"bytes,3,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffSizeRequest) Reset()         { *m = GetDiffSizeRequest{} }
func (m *GetDiffSizeRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiffSizeRequest) ProtoMessage()    {}
func (*GetDiffSizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3df405500be3cf, []int{0}
}
func (m *GetDiffSizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffSizeRequest.Unmarshal(m, b)
}
func (m *GetDiffSizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffSizeRequest.Marshal(b, m, deterministic)
}
func (m *GetDiffSizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffSizeRequest.Merge(m, src)
}
func (m *GetDiffSizeRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiffSizeRequest.Size(m)
}
func (m *GetDiffSizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffSizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffSizeRequest proto.InternalMessageInfo

func (m *GetDiffSizeRequest) GetVerFrom() string {
	if m != nil {
		return m.VerFrom
	}
	return ""
}

func (m *GetDiffSizeRequest) GetVerTo() string {
	if m != nil {
		return m.VerTo
	}
	return ""
}

func (m *GetDiffSizeRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//GetDiffSize返回
type GetDiffSizeResponse struct {
	//
	//文件大小
	Size_ int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size" form:"size"`
	//
	//差异大小
	DiffSize             int64    `protobuf:"varint,2,opt,name=diffSize,proto3" json:"diffSize" form:"diffSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiffSizeResponse) Reset()         { *m = GetDiffSizeResponse{} }
func (m *GetDiffSizeResponse) String() string { return proto.CompactTextString(m) }
func (*GetDiffSizeResponse) ProtoMessage()    {}
func (*GetDiffSizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3df405500be3cf, []int{1}
}
func (m *GetDiffSizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffSizeResponse.Unmarshal(m, b)
}
func (m *GetDiffSizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffSizeResponse.Marshal(b, m, deterministic)
}
func (m *GetDiffSizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffSizeResponse.Merge(m, src)
}
func (m *GetDiffSizeResponse) XXX_Size() int {
	return xxx_messageInfo_GetDiffSizeResponse.Size(m)
}
func (m *GetDiffSizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffSizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffSizeResponse proto.InternalMessageInfo

func (m *GetDiffSizeResponse) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *GetDiffSizeResponse) GetDiffSize() int64 {
	if m != nil {
		return m.DiffSize
	}
	return 0
}

//
//GetDiffSizeApi返回
type GetDiffSizeResponseWrapper struct {
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
	Data                 *GetDiffSizeResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDiffSizeResponseWrapper) Reset()         { *m = GetDiffSizeResponseWrapper{} }
func (m *GetDiffSizeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetDiffSizeResponseWrapper) ProtoMessage()    {}
func (*GetDiffSizeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3df405500be3cf, []int{2}
}
func (m *GetDiffSizeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiffSizeResponseWrapper.Unmarshal(m, b)
}
func (m *GetDiffSizeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiffSizeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetDiffSizeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiffSizeResponseWrapper.Merge(m, src)
}
func (m *GetDiffSizeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetDiffSizeResponseWrapper.Size(m)
}
func (m *GetDiffSizeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiffSizeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiffSizeResponseWrapper proto.InternalMessageInfo

func (m *GetDiffSizeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetDiffSizeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetDiffSizeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetDiffSizeResponseWrapper) GetData() *GetDiffSizeResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDiffSizeRequest)(nil), "archive.GetDiffSizeRequest")
	proto.RegisterType((*GetDiffSizeResponse)(nil), "archive.GetDiffSizeResponse")
	proto.RegisterType((*GetDiffSizeResponseWrapper)(nil), "archive.GetDiffSizeResponseWrapper")
}

func init() { proto.RegisterFile("get_archive_diff_size.proto", fileDescriptor_ed3df405500be3cf) }

var fileDescriptor_ed3df405500be3cf = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x89, 0xbb, 0xdb, 0x3f, 0xb3, 0x4a, 0xcb, 0x14, 0x24, 0xac, 0x42, 0xca, 0x28, 0xd2,
	0x4b, 0x12, 0xad, 0x22, 0xd5, 0x8b, 0x74, 0xd1, 0x8a, 0x47, 0xa7, 0x82, 0x07, 0xd1, 0x30, 0x9b,
	0xcc, 0x64, 0xc7, 0x34, 0x7d, 0xe3, 0x64, 0x36, 0x95, 0x4a, 0xf1, 0xe4, 0xc1, 0x8f, 0xe2, 0x97,
	0x8a, 0xe0, 0xc1, 0x0f, 0x90, 0x4f, 0x20, 0x99, 0xa4, 0x9b, 0x08, 0xbd, 0xf4, 0xb0, 0xa7, 0xcc,
	0xfb, 0xfe, 0xde, 0x87, 0xe7, 0x61, 0xde, 0x0c, 0xba, 0x13, 0x73, 0x1d, 0x30, 0x15, 0xce, 0x65,
	0xc1, 0x83, 0x48, 0x0a, 0x11, 0xe4, 0xf2, 0x9c, 0x7b, 0x99, 0x02, 0x0d, 0x78, 0xbd, 0x05, 0x13,
	0x37, 0x96, 0x7a, 0xbe, 0x98, 0x79, 0x21, 0xa4, 0x7e, 0x0c, 0x31, 0xf8, 0x86, 0xcf, 0x16, 0xc2,
	0x54, 0xa6, 0x30, 0xa7, 0x46, 0x37, 0x79, 0xda, 0x1b, 0x4f, 0xcf, 0xa4, 0x4e, 0xe0, 0xcc, 0x8f,
	0xc1, 0x35, 0xd0, 0x2d, 0xd8, 0x89, 0x8c, 0x98, 0x06, 0x95, 0xfb, 0xcb, 0x63, 0xa3, 0x23, 0xbf,
	0x06, 0x08, 0xbf, 0xe6, 0xfa, 0xa5, 0x14, 0xe2, 0x58, 0x9e, 0x73, 0xca, 0xbf, 0x2c, 0x78, 0xae,
	0xf1, 0x0f, 0x0b, 0x6d, 0x14, 0x5c, 0x05, 0x42, 0x41, 0x6a, 0x5b, 0xbb, 0xd6, 0xde, 0xe6, 0xf4,
	0x73, 0x55, 0x3a, 0x5b, 0x02, 0x54, 0xfa, 0x9c, 0x5c, 0x12, 0xf2, 0xe7, 0xb7, 0x73, 0x8c, 0xde,
	0x7e, 0xfa, 0xc0, 0x5c, 0x71, 0xe8, 0x1e, 0x3d, 0x74, 0x9f, 0x7d, 0xfc, 0x76, 0x70, 0xe1, 0xbe,
	0xe8, 0xd7, 0x4f, 0xae, 0x59, 0x3f, 0xda, 0xbf, 0xb8, 0x4f, 0xd7, 0x0b, 0xae, 0x8e, 0x14, 0xa4,
	0xf8, 0x3b, 0x5a, 0xab, 0xcd, 0x34, 0xd8, 0x37, 0x4c, 0x88, 0x79, 0x55, 0x3a, 0xb7, 0xba, 0x10,
	0x1a, 0x56, 0x16, 0x61, 0x54, 0x70, 0xf5, 0x0e, 0xf0, 0x4f, 0x0b, 0x6d, 0x66, 0x2c, 0x4c, 0x58,
	0xcc, 0xdf, 0x44, 0xf6, 0xc0, 0x84, 0x48, 0xaa, 0xd2, 0xd9, 0x6e, 0x42, 0x2c, 0xd1, 0xca, 0x72,
	0x74, 0xee, 0x24, 0x41, 0x3b, 0xff, 0xad, 0x2a, 0xcf, 0xe0, 0x34, 0xe7, 0xf8, 0x1e, 0x1a, 0xd6,
	0x3f, 0x90, 0x59, 0xd3, 0x60, 0xba, 0x55, 0x95, 0xce, 0xb8, 0x09, 0x57, 0x77, 0x09, 0x35, 0x10,
	0xfb, 0x68, 0x23, 0x6a, 0x85, 0xe6, 0x2a, 0x07, 0xd3, 0x9d, 0x6e, 0x9f, 0x97, 0x84, 0xd0, 0xe5,
	0x10, 0xf9, 0x6b, 0xa1, 0xc9, 0x15, 0x6e, 0xef, 0x15, 0xcb, 0x32, 0xae, 0x6a, 0xd3, 0x10, 0xa2,
	0xc6, 0x74, 0xd4, 0x37, 0xad, 0xbb, 0x84, 0x1a, 0x88, 0x0f, 0xd0, 0xb8, 0xfe, 0xbe, 0xfa, 0x9a,
	0x9d, 0x30, 0x79, 0xda, 0xae, 0xf0, 0x76, 0x55, 0x3a, 0xb8, 0x9b, 0x6d, 0x21, 0xa1, 0xfd, 0x51,
	0xfc, 0x00, 0x8d, 0xb8, 0x52, 0xa0, 0xda, 0x1b, 0xdf, 0xae, 0x4a, 0xe7, 0x66, 0xa3, 0x31, 0x6d,
	0x42, 0x1b, 0x8c, 0x0f, 0xd1, 0x30, 0x62, 0x9a, 0xd9, 0xc3, 0x5d, 0x6b, 0x6f, 0xbc, 0x7f, 0xd7,
	0x6b, 0x5f, 0x8f, 0x77, 0x45, 0xf2, 0x7e, 0xc8, 0x5a, 0x43, 0xa8, 0x91, 0xce, 0xd6, 0xcc, 0x43,
	0x78, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x28, 0xe2, 0x69, 0x96, 0x97, 0x03, 0x00, 0x00,
}
