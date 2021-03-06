// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_file.proto

package workspace

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//UpdateFile请求
type UpdateFileRequest struct {
	//
	//rename的新文件名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//要操作的文件、目录路径,上传时path是指存储的目录
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//文件操作类型包括：上传 -> upload，删除 -> delete,重命名 -> rename,修改权限 -> chmod,创建目录 -> mkdir
	Op string `protobuf:"bytes,4,opt,name=op,proto3" json:"op" form:"op"`
	//
	//chmod、mkdir新文件权限（0755, 0644)
	Perm string `protobuf:"bytes,5,opt,name=perm,proto3" json:"perm" form:"perm"`
	//
	//包Id
	PackageId            string   `protobuf:"bytes,6,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFileRequest) Reset()         { *m = UpdateFileRequest{} }
func (m *UpdateFileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFileRequest) ProtoMessage()    {}
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_284bb8860c133e17, []int{0}
}
func (m *UpdateFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFileRequest.Unmarshal(m, b)
}
func (m *UpdateFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFileRequest.Merge(m, src)
}
func (m *UpdateFileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFileRequest.Size(m)
}
func (m *UpdateFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFileRequest proto.InternalMessageInfo

func (m *UpdateFileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UpdateFileRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *UpdateFileRequest) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *UpdateFileRequest) GetPerm() string {
	if m != nil {
		return m.Perm
	}
	return ""
}

func (m *UpdateFileRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//UpdateFileApi返回
type UpdateFileResponseWrapper struct {
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

func (m *UpdateFileResponseWrapper) Reset()         { *m = UpdateFileResponseWrapper{} }
func (m *UpdateFileResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateFileResponseWrapper) ProtoMessage()    {}
func (*UpdateFileResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_284bb8860c133e17, []int{1}
}
func (m *UpdateFileResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFileResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateFileResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFileResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateFileResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFileResponseWrapper.Merge(m, src)
}
func (m *UpdateFileResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateFileResponseWrapper.Size(m)
}
func (m *UpdateFileResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFileResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFileResponseWrapper proto.InternalMessageInfo

func (m *UpdateFileResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateFileResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateFileResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateFileResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateFileRequest)(nil), "workspace.UpdateFileRequest")
	proto.RegisterType((*UpdateFileResponseWrapper)(nil), "workspace.UpdateFileResponseWrapper")
}

func init() { proto.RegisterFile("update_file.proto", fileDescriptor_284bb8860c133e17) }

var fileDescriptor_284bb8860c133e17 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xef, 0x6e, 0xd3, 0x3c,
	0x14, 0xc6, 0xd5, 0xae, 0x9d, 0xde, 0xba, 0x7b, 0xd9, 0x66, 0xa4, 0x29, 0x0c, 0xa1, 0x80, 0x99,
	0x50, 0xba, 0xc9, 0x4d, 0xd9, 0xf8, 0x33, 0x0a, 0x62, 0x1a, 0xd2, 0x86, 0xf6, 0x11, 0x23, 0x84,
	0x44, 0xdd, 0x22, 0xb7, 0x71, 0xb3, 0xa8, 0x4d, 0x6c, 0xdc, 0x94, 0x02, 0x6b, 0x2f, 0x80, 0x8b,
	0xe2, 0x56, 0x32, 0x89, 0x4b, 0xc8, 0x15, 0x20, 0x3b, 0xe9, 0x1a, 0x89, 0x4f, 0x7c, 0xb2, 0xcf,
	0x79, 0x7e, 0xc7, 0xc7, 0x47, 0xcf, 0x01, 0xdb, 0x53, 0xe9, 0xb1, 0x98, 0x7f, 0x1e, 0x06, 0x63,
	0xde, 0x94, 0x4a, 0xc4, 0x02, 0xd6, 0x66, 0x42, 0x8d, 0x26, 0x92, 0x0d, 0xf8, 0x2e, 0xf6, 0x83,
	0xf8, 0x72, 0xda, 0x6f, 0x0e, 0x44, 0xe8, 0xfa, 0xc2, 0x17, 0xae, 0x21, 0xfa, 0xd3, 0xa1, 0x89,
	0x4c, 0x60, 0x6e, 0x59, 0xe5, 0xee, 0xb3, 0x02, 0x1e, 0xce, 0x82, 0x78, 0x24, 0x66, 0xae, 0x2f,
	0xb0, 0x11, 0xf1, 0x57, 0x36, 0x0e, 0x3c, 0x16, 0x0b, 0x35, 0x71, 0x6f, 0xae, 0x79, 0xdd, 0x5d,
	0x5f, 0x08, 0x7f, 0xcc, 0x57, 0xaf, 0xf3, 0x50, 0xc6, 0xdf, 0x33, 0x11, 0xfd, 0x5a, 0x03, 0xdb,
	0x1f, 0xcc, 0x27, 0xcf, 0x83, 0x31, 0x27, 0xfc, 0xcb, 0x94, 0x4f, 0x62, 0xf8, 0x10, 0x54, 0x22,
	0x16, 0x72, 0xab, 0x74, 0xbf, 0xe4, 0xd4, 0xde, 0x6c, 0xa6, 0x89, 0x5d, 0x1f, 0x0a, 0x15, 0xb6,
	0x91, 0xce, 0x22, 0x62, 0x44, 0x28, 0x41, 0x45, 0xb2, 0xf8, 0xd2, 0x2a, 0x1b, 0x88, 0xae, 0x20,
	0x9d, 0x45, 0xbf, 0xaf, 0xed, 0x0b, 0xf0, 0xb6, 0xe7, 0x38, 0xd4, 0xed, 0xf4, 0xa8, 0xdb, 0xa6,
	0xfb, 0xf4, 0x04, 0xa1, 0x57, 0xaf, 0xe9, 0x9c, 0x2a, 0x1a, 0x75, 0x0f, 0x1a, 0x07, 0x8d, 0xb9,
	0x43, 0xdd, 0xc6, 0xbc, 0xc3, 0xf0, 0x8f, 0x53, 0xfc, 0xa9, 0xdb, 0x76, 0x28, 0xed, 0xf4, 0x28,
	0xfd, 0x9b, 0xdc, 0xdf, 0x23, 0xa6, 0x13, 0x74, 0xc1, 0x7f, 0x3c, 0x1a, 0x08, 0x2f, 0x88, 0x7c,
	0x6b, 0xcd, 0x74, 0xbd, 0x9d, 0x26, 0xf6, 0x66, 0xd6, 0x75, 0xa9, 0x20, 0x72, 0x03, 0xc1, 0x7b,
	0xa0, 0x2c, 0xa4, 0x55, 0x31, 0xe8, 0xff, 0x69, 0x62, 0xd7, 0x32, 0x54, 0x48, 0x44, 0xca, 0x42,
	0xc2, 0xa7, 0xa0, 0x22, 0xb9, 0x0a, 0xad, 0xaa, 0x01, 0x1e, 0x14, 0x26, 0xe0, 0x2a, 0xd4, 0x13,
	0xdc, 0x02, 0x1b, 0x3d, 0xa7, 0xd5, 0x69, 0xe1, 0xe7, 0xdd, 0xab, 0xa3, 0x45, 0x83, 0x18, 0x1c,
	0xfe, 0x2c, 0x81, 0x9a, 0x64, 0x83, 0x11, 0xf3, 0xf9, 0x85, 0x67, 0xad, 0x9b, 0xe2, 0x51, 0x9a,
	0xd8, 0x5b, 0xcb, 0xf1, 0x73, 0x49, 0xbf, 0xf0, 0x1e, 0xbc, 0xeb, 0x75, 0x18, 0x1e, 0x9e, 0xe2,
	0xf3, 0x16, 0x7e, 0xd1, 0xbd, 0x3a, 0x5e, 0xe0, 0x93, 0x62, 0xfc, 0xe4, 0x1f, 0xe3, 0xc7, 0x87,
	0x8b, 0x3d, 0xb2, 0xea, 0x8e, 0xae, 0x4b, 0xe0, 0x4e, 0xd1, 0xbf, 0x89, 0x14, 0xd1, 0x84, 0x7f,
	0x54, 0x4c, 0x4a, 0xae, 0xb4, 0x8f, 0x03, 0xe1, 0x65, 0x3e, 0x56, 0x8b, 0x3e, 0xea, 0x2c, 0x22,
	0x46, 0x84, 0xc7, 0xa0, 0xae, 0xcf, 0xb3, 0x6f, 0x72, 0xcc, 0x82, 0x28, 0xb7, 0x73, 0x27, 0x4d,
	0x6c, 0xb8, 0x62, 0x73, 0x11, 0x91, 0x22, 0x0a, 0x1f, 0x81, 0x2a, 0x57, 0x4a, 0xa8, 0xdc, 0x8c,
	0xad, 0x34, 0xb1, 0x37, 0x72, 0x33, 0x74, 0x1a, 0x91, 0x4c, 0x86, 0x2f, 0x41, 0xc5, 0x63, 0x31,
	0x33, 0x46, 0xd4, 0x0f, 0x77, 0x9a, 0xd9, 0x42, 0x36, 0x97, 0x0b, 0xd9, 0x3c, 0xd3, 0x0b, 0x59,
	0xfc, 0x9e, 0xa6, 0x11, 0x31, 0x45, 0xfd, 0x75, 0x83, 0x1d, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x55, 0x38, 0x6e, 0x13, 0x4c, 0x03, 0x00, 0x00,
}
