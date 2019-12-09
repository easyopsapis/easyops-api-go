// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_workspace_status.proto

package workspace

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Status请求
type StatusRequest struct {
	//
	//是否要获取文件内容差异
	DiffFile string `protobuf:"bytes,1,opt,name=diff_file,json=diffFile,proto3" json:"diff_file" form:"diff_file"`
	//
	//只查看指定文件内容差异
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path" form:"path"`
	//
	//参数path的编码格式
	Encoding string `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding" form:"encoding"`
	//
	//程序包Id
	PackageId            string   `protobuf:"bytes,4,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b7751fcb3aa3f8, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetDiffFile() string {
	if m != nil {
		return m.DiffFile
	}
	return ""
}

func (m *StatusRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *StatusRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *StatusRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//StatusApi返回
type StatusResponseWrapper struct {
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
	Data                 []*file_repository.Diff `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StatusResponseWrapper) Reset()         { *m = StatusResponseWrapper{} }
func (m *StatusResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*StatusResponseWrapper) ProtoMessage()    {}
func (*StatusResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b7751fcb3aa3f8, []int{1}
}
func (m *StatusResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponseWrapper.Unmarshal(m, b)
}
func (m *StatusResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponseWrapper.Marshal(b, m, deterministic)
}
func (m *StatusResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponseWrapper.Merge(m, src)
}
func (m *StatusResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_StatusResponseWrapper.Size(m)
}
func (m *StatusResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponseWrapper proto.InternalMessageInfo

func (m *StatusResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *StatusResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *StatusResponseWrapper) GetData() []*file_repository.Diff {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "workspace.StatusRequest")
	proto.RegisterType((*StatusResponseWrapper)(nil), "workspace.StatusResponseWrapper")
}

func init() { proto.RegisterFile("get_workspace_status.proto", fileDescriptor_12b7751fcb3aa3f8) }

var fileDescriptor_12b7751fcb3aa3f8 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xb7, 0x0e, 0xad, 0x2e, 0x68, 0x93, 0x61, 0xa8, 0xaa, 0x90, 0x32, 0x99, 0x09, 0x75,
	0xa0, 0xc4, 0x6c, 0x20, 0x34, 0xf6, 0x32, 0x31, 0x60, 0x12, 0x8f, 0x64, 0x48, 0x3c, 0xac, 0x6b,
	0xe5, 0x26, 0x4e, 0x66, 0x35, 0xcd, 0x35, 0xb6, 0x4b, 0x99, 0xd0, 0x3e, 0x80, 0xff, 0xe2, 0x2b,
	0x78, 0x08, 0x12, 0x9f, 0x90, 0x2f, 0x40, 0x76, 0xb2, 0xb4, 0xe2, 0x8d, 0xa7, 0xdc, 0x73, 0xcf,
	0x39, 0xd7, 0x37, 0x47, 0x17, 0xf5, 0x53, 0x6e, 0xc6, 0x0b, 0x50, 0x53, 0x2d, 0x59, 0xc4, 0xc7,
	0xda, 0x30, 0x33, 0xd7, 0x81, 0x54, 0x60, 0x00, 0x77, 0x9a, 0x7e, 0xdf, 0x4f, 0x85, 0xb9, 0x9a,
	0x4f, 0x82, 0x08, 0x66, 0x34, 0x85, 0x14, 0xa8, 0x53, 0x4c, 0xe6, 0x89, 0x43, 0x0e, 0xb8, 0xaa,
	0x72, 0xf6, 0x3f, 0xa5, 0x10, 0x70, 0xa6, 0xaf, 0x41, 0xea, 0x20, 0x83, 0x88, 0x65, 0x34, 0x82,
	0xdc, 0x28, 0x16, 0x19, 0x5d, 0x39, 0x15, 0x97, 0xe0, 0xcf, 0x20, 0xe6, 0x99, 0xa6, 0xb5, 0x90,
	0x3a, 0x48, 0x13, 0x91, 0xf1, 0xb1, 0xa5, 0xb5, 0x30, 0xa0, 0xae, 0x69, 0x2c, 0x92, 0xa4, 0x9e,
	0xfa, 0x6a, 0x65, 0x89, 0xd9, 0x42, 0x98, 0x29, 0x2c, 0x68, 0x0a, 0xbe, 0x23, 0xfd, 0xaf, 0x2c,
	0x13, 0x31, 0x33, 0xa0, 0x34, 0x6d, 0xca, 0xca, 0x47, 0x7e, 0xae, 0xa1, 0x7b, 0xe7, 0xee, 0xc7,
	0x42, 0xfe, 0x65, 0xce, 0xb5, 0xc1, 0x07, 0xa8, 0x63, 0xe7, 0x8e, 0xed, 0x63, 0xbd, 0xd6, 0x6e,
	0x6b, 0xd0, 0x39, 0x7d, 0x50, 0x16, 0xde, 0x76, 0x02, 0x6a, 0x76, 0x4c, 0x1a, 0x8a, 0x84, 0x9b,
	0xb6, 0x3e, 0x13, 0x19, 0xc7, 0x6f, 0x51, 0x5b, 0x32, 0x73, 0xd5, 0x5b, 0x73, 0x6a, 0x5a, 0x16,
	0x5e, 0xb7, 0x52, 0xdb, 0x2e, 0xf9, 0xf3, 0xdb, 0x7b, 0x84, 0xfa, 0xa3, 0x21, 0x1d, 0x0c, 0x2e,
	0x46, 0x43, 0x7a, 0xf9, 0x6c, 0x7f, 0x30, 0xa4, 0x75, 0xf5, 0x74, 0xff, 0x64, 0x2f, 0x74, 0x66,
	0x4c, 0xd1, 0x26, 0xcf, 0x23, 0x88, 0x45, 0x9e, 0xf6, 0xd6, 0xdd, 0xa0, 0xfb, 0x65, 0xe1, 0x6d,
	0x55, 0x83, 0x6e, 0x19, 0x12, 0x36, 0x22, 0xfc, 0xa3, 0x85, 0x3a, 0x92, 0x45, 0x53, 0x96, 0xf2,
	0x0f, 0x71, 0xaf, 0xed, 0x2c, 0xd3, 0xe5, 0xa6, 0x0d, 0x65, 0x17, 0x38, 0x47, 0x1f, 0x47, 0x17,
	0xcc, 0x4f, 0xde, 0xf8, 0x67, 0xcf, 0xfd, 0xd7, 0x97, 0xdf, 0x8f, 0x6e, 0xfc, 0x93, 0x55, 0xfc,
	0xf2, 0x3f, 0xf1, 0xc1, 0xe1, 0xcd, 0x5e, 0xb8, 0x7c, 0x9d, 0xfc, 0x6a, 0xa1, 0x9d, 0xdb, 0x18,
	0xb5, 0x84, 0x5c, 0xf3, 0xcf, 0x8a, 0x49, 0xc9, 0x15, 0x7e, 0x8c, 0xda, 0x11, 0xc4, 0x55, 0x92,
	0x1b, 0xa7, 0x5b, 0xcb, 0x6c, 0x6c, 0x97, 0x84, 0x8e, 0xc4, 0x47, 0xa8, 0x6b, 0xbf, 0xef, 0xbf,
	0xc9, 0x8c, 0x89, 0xbc, 0xce, 0xf1, 0x61, 0x59, 0x78, 0x78, 0xa9, 0xad, 0x49, 0x12, 0xae, 0x4a,
	0xf1, 0x13, 0xb4, 0xc1, 0x95, 0x02, 0x55, 0x47, 0xb6, 0x5d, 0x16, 0xde, 0xdd, 0x3a, 0x32, 0xdb,
	0x26, 0x61, 0x45, 0xe3, 0x63, 0xd4, 0x8e, 0x99, 0x61, 0xbd, 0xf6, 0xee, 0xfa, 0xa0, 0x7b, 0xb8,
	0x13, 0xfc, 0x73, 0x4a, 0xc1, 0x3b, 0x91, 0x24, 0xab, 0xdb, 0x59, 0x31, 0x09, 0x9d, 0x67, 0x72,
	0xc7, 0x9d, 0xca, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x43, 0x07, 0xe1, 0x10, 0x03,
	0x00, 0x00,
}
