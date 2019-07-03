// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_difference.proto

package archive

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	file_repository "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/file_repository"
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
//GetDifference请求
type GetDifferenceRequest struct {
	//
	//源版本Id
	VerFrom string `protobuf:"bytes,1,opt,name=ver_from,json=verFrom,proto3" json:"ver_from" form:"ver_from"`
	//
	//目的版本Id
	VerTo string `protobuf:"bytes,2,opt,name=ver_to,json=verTo,proto3" json:"ver_to" form:"ver_to"`
	//
	//是否要获取文件内容差异（true,false）
	DiffFile string `protobuf:"bytes,3,opt,name=diff_file,json=diffFile,proto3" json:"diff_file" form:"diff_file"`
	//
	//只查看指定的文件内容差异
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,5,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//程序包Id
	PackageId            string   `protobuf:"bytes,6,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDifferenceRequest) Reset()         { *m = GetDifferenceRequest{} }
func (m *GetDifferenceRequest) String() string { return proto.CompactTextString(m) }
func (*GetDifferenceRequest) ProtoMessage()    {}
func (*GetDifferenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af03b77173a80cb, []int{0}
}
func (m *GetDifferenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDifferenceRequest.Unmarshal(m, b)
}
func (m *GetDifferenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDifferenceRequest.Marshal(b, m, deterministic)
}
func (m *GetDifferenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDifferenceRequest.Merge(m, src)
}
func (m *GetDifferenceRequest) XXX_Size() int {
	return xxx_messageInfo_GetDifferenceRequest.Size(m)
}
func (m *GetDifferenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDifferenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDifferenceRequest proto.InternalMessageInfo

func (m *GetDifferenceRequest) GetVerFrom() string {
	if m != nil {
		return m.VerFrom
	}
	return ""
}

func (m *GetDifferenceRequest) GetVerTo() string {
	if m != nil {
		return m.VerTo
	}
	return ""
}

func (m *GetDifferenceRequest) GetDiffFile() string {
	if m != nil {
		return m.DiffFile
	}
	return ""
}

func (m *GetDifferenceRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetDifferenceRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *GetDifferenceRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//GetDifferenceApi返回
type GetDifferenceResponseWrapper struct {
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
	Data                 *file_repository.Diff `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDifferenceResponseWrapper) Reset()         { *m = GetDifferenceResponseWrapper{} }
func (m *GetDifferenceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetDifferenceResponseWrapper) ProtoMessage()    {}
func (*GetDifferenceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af03b77173a80cb, []int{1}
}
func (m *GetDifferenceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDifferenceResponseWrapper.Unmarshal(m, b)
}
func (m *GetDifferenceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDifferenceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetDifferenceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDifferenceResponseWrapper.Merge(m, src)
}
func (m *GetDifferenceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetDifferenceResponseWrapper.Size(m)
}
func (m *GetDifferenceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDifferenceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetDifferenceResponseWrapper proto.InternalMessageInfo

func (m *GetDifferenceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetDifferenceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetDifferenceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetDifferenceResponseWrapper) GetData() *file_repository.Diff {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDifferenceRequest)(nil), "archive.GetDifferenceRequest")
	proto.RegisterType((*GetDifferenceResponseWrapper)(nil), "archive.GetDifferenceResponseWrapper")
}

func init() { proto.RegisterFile("get_difference.proto", fileDescriptor_6af03b77173a80cb) }

var fileDescriptor_6af03b77173a80cb = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0xce, 0xea, 0x2e, 0x3f, 0x83, 0x06, 0x32, 0xa2, 0xd9, 0x10, 0x93, 0x9a, 0x91, 0x98, 0x05,
	0xd2, 0x8e, 0xa0, 0x31, 0xb8, 0x31, 0x12, 0x89, 0x42, 0xb8, 0x6d, 0x48, 0x4c, 0xdc, 0x61, 0x37,
	0x43, 0x3b, 0x2d, 0x0d, 0x6d, 0x4f, 0x9d, 0x0e, 0x8b, 0x28, 0x24, 0xde, 0xe8, 0x85, 0x0f, 0x59,
	0xa3, 0x8f, 0xd0, 0x27, 0x30, 0x33, 0x2d, 0x5b, 0xc4, 0x2b, 0x2f, 0xb8, 0xea, 0x39, 0xe7, 0xfb,
	0x4e, 0xcf, 0x77, 0x7e, 0x06, 0x2d, 0x86, 0x42, 0x8d, 0xfc, 0x28, 0x08, 0x84, 0x14, 0xa9, 0x27,
	0x9c, 0x4c, 0x82, 0x02, 0x3c, 0xcd, 0xa5, 0x77, 0x14, 0x8d, 0xc5, 0x92, 0x1d, 0x46, 0xea, 0xe8,
	0xe4, 0xd0, 0xf1, 0x20, 0xa1, 0x21, 0x84, 0x40, 0x0d, 0x7e, 0x78, 0x12, 0x18, 0xcf, 0x38, 0xc6,
	0xaa, 0xf2, 0x96, 0xf6, 0x43, 0x70, 0x04, 0xcf, 0xcf, 0x20, 0xcb, 0x9d, 0x18, 0x3c, 0x1e, 0x53,
	0x0f, 0x52, 0x25, 0xb9, 0xa7, 0xf2, 0x2a, 0x53, 0x8a, 0x0c, 0xec, 0x04, 0x7c, 0x11, 0xe7, 0xb4,
	0x26, 0x52, 0xe3, 0xd2, 0x20, 0x8a, 0xc5, 0x48, 0xc3, 0x79, 0xa4, 0x40, 0x9e, 0x51, 0xad, 0xa9,
	0xfe, 0xeb, 0x8b, 0x2b, 0x22, 0x92, 0xd3, 0x48, 0x1d, 0xc3, 0x29, 0x0d, 0xc1, 0x36, 0xa0, 0x3d,
	0xe6, 0x71, 0xe4, 0x73, 0x05, 0x32, 0xa7, 0x13, 0xb3, 0xca, 0x23, 0xdf, 0x3a, 0x68, 0x71, 0x57,
	0xa8, 0xb7, 0x93, 0xee, 0x5c, 0xf1, 0xf1, 0x44, 0xe4, 0x0a, 0x7f, 0x6f, 0xa1, 0x99, 0xb1, 0x90,
	0xa3, 0x40, 0x42, 0xd2, 0x6d, 0x3d, 0x6a, 0xf5, 0x66, 0xb7, 0x8f, 0xcb, 0xc2, 0x9a, 0x0f, 0x40,
	0x26, 0x7d, 0x72, 0x89, 0x90, 0xdf, 0x3f, 0xad, 0x7d, 0xe4, 0x0e, 0x07, 0x5b, 0xdc, 0xfe, 0xfc,
	0xc6, 0xfe, 0xf0, 0xd4, 0x7e, 0x79, 0xf0, 0x65, 0xf3, 0xc2, 0xfe, 0xcb, 0x7f, 0xfe, 0x9f, 0xfe,
	0xfa, 0xc6, 0xc5, 0xb2, 0x3b, 0x3d, 0x16, 0x72, 0x47, 0x42, 0x82, 0xbf, 0xb6, 0xd0, 0x94, 0x2e,
	0xa7, 0xa0, 0x7b, 0xcb, 0xc8, 0x88, 0xca, 0xc2, 0xba, 0xdb, 0xc8, 0x50, 0x70, 0x73, 0x22, 0x3a,
	0x63, 0x21, 0xf7, 0x01, 0xaf, 0xa3, 0x59, 0x3d, 0xea, 0x91, 0x9e, 0x7f, 0xf7, 0xb6, 0x11, 0xb1,
	0x58, 0x16, 0xd6, 0x42, 0x25, 0x62, 0x02, 0x11, 0x77, 0x46, 0xdb, 0x3b, 0x51, 0x2c, 0x70, 0x86,
	0xda, 0x19, 0x57, 0x47, 0xdd, 0xb6, 0x61, 0xb3, 0xb2, 0xb0, 0xe6, 0x2a, 0xb6, 0x8e, 0x6a, 0xc1,
	0x7b, 0x68, 0x77, 0xd8, 0xeb, 0x31, 0x3a, 0x18, 0x32, 0xda, 0x67, 0xab, 0x6c, 0x8b, 0x90, 0x57,
	0xaf, 0xd9, 0x39, 0x93, 0x2c, 0x3d, 0x58, 0x5b, 0x59, 0x5b, 0x39, 0xef, 0x31, 0xba, 0x72, 0x3e,
	0xa8, 0x14, 0x1d, 0xf4, 0x7b, 0x8c, 0x0d, 0x86, 0x8c, 0xfd, 0xcb, 0x5c, 0x5d, 0x76, 0x4d, 0x25,
	0x4c, 0xd1, 0x8c, 0x48, 0x3d, 0xf0, 0xa3, 0x34, 0xec, 0x76, 0x4c, 0xd5, 0x7b, 0xcd, 0xbe, 0x2e,
	0x11, 0xe2, 0x4e, 0x48, 0xf8, 0x47, 0x0b, 0xcd, 0x66, 0xdc, 0x3b, 0xe6, 0xa1, 0xd8, 0xf3, 0xbb,
	0x53, 0x26, 0x25, 0x6e, 0xda, 0x9a, 0x40, 0x37, 0x37, 0xde, 0xa6, 0x3c, 0xf9, 0xd5, 0x42, 0x0f,
	0xaf, 0xdd, 0x61, 0x9e, 0x41, 0x9a, 0x8b, 0xf7, 0x92, 0x67, 0x99, 0x90, 0xf8, 0x31, 0x6a, 0x7b,
	0xe0, 0x0b, 0x73, 0x8a, 0x9d, 0xed, 0xf9, 0x66, 0xa0, 0x3a, 0x4a, 0x5c, 0x03, 0xe2, 0x4d, 0x34,
	0xa7, 0xbf, 0xef, 0x3e, 0x65, 0x31, 0x8f, 0xd2, 0xfa, 0x5e, 0x1e, 0x94, 0x85, 0x85, 0x1b, 0x6e,
	0x0d, 0x12, 0xf7, 0x2a, 0x15, 0x3f, 0x41, 0x1d, 0x21, 0x25, 0xc8, 0x7a, 0xbd, 0x0b, 0x65, 0x61,
	0xdd, 0xa9, 0x47, 0xa7, 0xc3, 0xc4, 0xad, 0x60, 0xdc, 0x47, 0x6d, 0x9f, 0x2b, 0x6e, 0xf6, 0x3a,
	0xb7, 0x71, 0xdf, 0xb9, 0xf6, 0x24, 0x1d, 0xdd, 0xc0, 0x55, 0x75, 0x9a, 0x4c, 0x5c, 0x93, 0x73,
	0x38, 0x65, 0x9e, 0xdc, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xf8, 0xf1, 0x78, 0x50,
	0x04, 0x00, 0x00,
}