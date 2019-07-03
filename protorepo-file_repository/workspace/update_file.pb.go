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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

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
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0xae, 0x9d, 0xa8, 0x3b, 0xd8, 0x66, 0xa4, 0x29, 0x0c, 0xa1, 0x80, 0x99, 0x50,
	0xba, 0x29, 0x4d, 0xd9, 0xf8, 0x33, 0x0a, 0x62, 0x02, 0x69, 0xa0, 0xdd, 0x5a, 0x20, 0x24, 0xea,
	0x16, 0xb9, 0x8d, 0x9b, 0x45, 0x4d, 0x62, 0xe3, 0xa4, 0x14, 0x58, 0xfb, 0x02, 0x3c, 0x14, 0x8f,
	0x92, 0x49, 0x3c, 0x42, 0x9e, 0x00, 0xd9, 0x49, 0xd7, 0x20, 0xae, 0x76, 0x65, 0x9f, 0xf3, 0xfd,
	0x8e, 0x8f, 0x8f, 0xbe, 0x03, 0xb6, 0xa7, 0xc2, 0xa5, 0x09, 0xfb, 0x32, 0xf6, 0x03, 0xd6, 0x16,
	0x92, 0x27, 0x1c, 0x36, 0x66, 0x5c, 0x4e, 0x62, 0x41, 0x47, 0x6c, 0xd7, 0xf6, 0xfc, 0xe4, 0x7c,
	0x3a, 0x6c, 0x8f, 0x78, 0xe8, 0x78, 0xdc, 0xe3, 0x8e, 0x26, 0x86, 0xd3, 0xb1, 0x8e, 0x74, 0xa0,
	0x6f, 0x79, 0xe5, 0xee, 0xb3, 0x12, 0x1e, 0xce, 0xfc, 0x64, 0xc2, 0x67, 0x8e, 0xc7, 0x6d, 0x2d,
	0xda, 0xdf, 0x68, 0xe0, 0xbb, 0x34, 0xe1, 0x32, 0x76, 0xae, 0xae, 0x45, 0xdd, 0x5d, 0x8f, 0x73,
	0x2f, 0x60, 0xab, 0xd7, 0x59, 0x28, 0x92, 0x1f, 0xb9, 0x88, 0x7e, 0xaf, 0x81, 0xed, 0x8f, 0xfa,
	0x93, 0xef, 0xfc, 0x80, 0x61, 0xf6, 0x75, 0xca, 0xe2, 0x04, 0x3e, 0x04, 0xb5, 0x88, 0x86, 0xcc,
	0xa8, 0xdc, 0xaf, 0x58, 0x8d, 0xb7, 0x9b, 0x59, 0x6a, 0x36, 0xc7, 0x5c, 0x86, 0x5d, 0xa4, 0xb2,
	0x08, 0x6b, 0x11, 0x0a, 0x50, 0x13, 0x34, 0x39, 0x37, 0xaa, 0x1a, 0x22, 0x2b, 0x48, 0x65, 0xd1,
	0x9f, 0x4b, 0xf3, 0x0c, 0xbc, 0x1f, 0x58, 0x16, 0x71, 0x7a, 0x03, 0xe2, 0x74, 0xc9, 0x3e, 0x39,
	0x41, 0xe8, 0xd5, 0x6b, 0x32, 0x27, 0x92, 0x44, 0xfd, 0x83, 0xd6, 0x41, 0x6b, 0x6e, 0x11, 0xa7,
	0x35, 0xef, 0x51, 0xfb, 0xe7, 0x1b, 0xfb, 0x73, 0xbf, 0x6b, 0x11, 0xd2, 0x1b, 0x10, 0xf2, 0x3f,
	0xb9, 0xbf, 0x87, 0x75, 0x27, 0xe8, 0x80, 0x1b, 0x2c, 0x1a, 0x71, 0xd7, 0x8f, 0x3c, 0x63, 0x4d,
	0x77, 0xbd, 0x9d, 0xa5, 0xe6, 0x66, 0xde, 0x75, 0xa9, 0x20, 0x7c, 0x05, 0xc1, 0x7b, 0xa0, 0xca,
	0x85, 0x51, 0xd3, 0xe8, 0xcd, 0x2c, 0x35, 0x1b, 0x39, 0xca, 0x05, 0xc2, 0x55, 0x2e, 0xe0, 0x53,
	0x50, 0x13, 0x4c, 0x86, 0x46, 0x5d, 0x03, 0x0f, 0x4a, 0x13, 0x30, 0x19, 0xaa, 0x09, 0x6e, 0x81,
	0x8d, 0x81, 0xd5, 0xe9, 0x75, 0xec, 0xe7, 0xfd, 0x8b, 0xa3, 0x45, 0x0b, 0x6b, 0x1c, 0xfe, 0xaa,
	0x80, 0x86, 0xa0, 0xa3, 0x09, 0xf5, 0xd8, 0x99, 0x6b, 0xac, 0xeb, 0xe2, 0x20, 0x4b, 0xcd, 0xad,
	0xe5, 0xf8, 0x85, 0xa4, 0x5e, 0xf8, 0x00, 0xf0, 0xa0, 0x77, 0x92, 0x8f, 0xd8, 0xb1, 0x5f, 0xf4,
	0x2f, 0x8e, 0x17, 0xf6, 0x3f, 0xf1, 0x93, 0x6b, 0xc6, 0x8f, 0x0f, 0x17, 0x7b, 0x78, 0xd5, 0x1e,
	0x5d, 0x56, 0xc0, 0x9d, 0xb2, 0x81, 0xb1, 0xe0, 0x51, 0xcc, 0x3e, 0x49, 0x2a, 0x04, 0x93, 0xca,
	0xc8, 0x11, 0x77, 0x73, 0x23, 0xeb, 0x65, 0x23, 0x55, 0x16, 0x61, 0x2d, 0xc2, 0x63, 0xd0, 0x54,
	0xe7, 0xe9, 0x77, 0x11, 0x50, 0x3f, 0x2a, 0xfc, 0xdc, 0xc9, 0x52, 0x13, 0xae, 0xd8, 0x42, 0x44,
	0xb8, 0x8c, 0xc2, 0x47, 0xa0, 0xce, 0xa4, 0xe4, 0xb2, 0x70, 0x63, 0x2b, 0x4b, 0xcd, 0x8d, 0xc2,
	0x0d, 0x95, 0x46, 0x38, 0x97, 0xe1, 0x4b, 0x50, 0x73, 0x69, 0x42, 0xb5, 0x13, 0xcd, 0xc3, 0x9d,
	0x76, 0xbe, 0x91, 0xed, 0xe5, 0x46, 0xb6, 0x4f, 0xd5, 0x46, 0x96, 0xbf, 0xa7, 0x68, 0x84, 0x75,
	0xd1, 0x70, 0x5d, 0x63, 0x47, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc7, 0xf1, 0x53, 0x4d,
	0x03, 0x00, 0x00,
}
