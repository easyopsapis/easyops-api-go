// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_update_package_permission.proto

package pkg

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//BatchUpdatePackagePermission请求
type BatchUpdatePackagePermissionRequest struct {
	//
	//需要变动的权限
	Permissions []string `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions" form:"permissions"`
	//
	//需要权限变动的用户
	Authorizers []string `protobuf:"bytes,2,rep,name=authorizers,proto3" json:"authorizers" form:"authorizers"`
	//
	//变动的方法
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method" form:"method"`
	//
	//变动的包id
	PackageIds           []string `protobuf:"bytes,4,rep,name=packageIds,proto3" json:"packageIds" form:"packageIds"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchUpdatePackagePermissionRequest) Reset()         { *m = BatchUpdatePackagePermissionRequest{} }
func (m *BatchUpdatePackagePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*BatchUpdatePackagePermissionRequest) ProtoMessage()    {}
func (*BatchUpdatePackagePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8e4008e6164ed3, []int{0}
}
func (m *BatchUpdatePackagePermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdatePackagePermissionRequest.Unmarshal(m, b)
}
func (m *BatchUpdatePackagePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdatePackagePermissionRequest.Marshal(b, m, deterministic)
}
func (m *BatchUpdatePackagePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdatePackagePermissionRequest.Merge(m, src)
}
func (m *BatchUpdatePackagePermissionRequest) XXX_Size() int {
	return xxx_messageInfo_BatchUpdatePackagePermissionRequest.Size(m)
}
func (m *BatchUpdatePackagePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdatePackagePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdatePackagePermissionRequest proto.InternalMessageInfo

func (m *BatchUpdatePackagePermissionRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *BatchUpdatePackagePermissionRequest) GetAuthorizers() []string {
	if m != nil {
		return m.Authorizers
	}
	return nil
}

func (m *BatchUpdatePackagePermissionRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *BatchUpdatePackagePermissionRequest) GetPackageIds() []string {
	if m != nil {
		return m.PackageIds
	}
	return nil
}

//
//BatchUpdatePackagePermission返回
type BatchUpdatePackagePermissionResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//返回信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	//
	//包Id
	PackageId            string   `protobuf:"bytes,3,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchUpdatePackagePermissionResponse) Reset()         { *m = BatchUpdatePackagePermissionResponse{} }
func (m *BatchUpdatePackagePermissionResponse) String() string { return proto.CompactTextString(m) }
func (*BatchUpdatePackagePermissionResponse) ProtoMessage()    {}
func (*BatchUpdatePackagePermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8e4008e6164ed3, []int{1}
}
func (m *BatchUpdatePackagePermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponse.Unmarshal(m, b)
}
func (m *BatchUpdatePackagePermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponse.Marshal(b, m, deterministic)
}
func (m *BatchUpdatePackagePermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdatePackagePermissionResponse.Merge(m, src)
}
func (m *BatchUpdatePackagePermissionResponse) XXX_Size() int {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponse.Size(m)
}
func (m *BatchUpdatePackagePermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdatePackagePermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdatePackagePermissionResponse proto.InternalMessageInfo

func (m *BatchUpdatePackagePermissionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchUpdatePackagePermissionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BatchUpdatePackagePermissionResponse) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

//
//BatchUpdatePackagePermissionApi返回
type BatchUpdatePackagePermissionResponseWrapper struct {
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
	Data                 []*BatchUpdatePackagePermissionResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *BatchUpdatePackagePermissionResponseWrapper) Reset() {
	*m = BatchUpdatePackagePermissionResponseWrapper{}
}
func (m *BatchUpdatePackagePermissionResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*BatchUpdatePackagePermissionResponseWrapper) ProtoMessage() {}
func (*BatchUpdatePackagePermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8e4008e6164ed3, []int{2}
}
func (m *BatchUpdatePackagePermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper.Unmarshal(m, b)
}
func (m *BatchUpdatePackagePermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *BatchUpdatePackagePermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper.Merge(m, src)
}
func (m *BatchUpdatePackagePermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper.Size(m)
}
func (m *BatchUpdatePackagePermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdatePackagePermissionResponseWrapper proto.InternalMessageInfo

func (m *BatchUpdatePackagePermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchUpdatePackagePermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *BatchUpdatePackagePermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BatchUpdatePackagePermissionResponseWrapper) GetData() []*BatchUpdatePackagePermissionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchUpdatePackagePermissionRequest)(nil), "pkg.BatchUpdatePackagePermissionRequest")
	proto.RegisterType((*BatchUpdatePackagePermissionResponse)(nil), "pkg.BatchUpdatePackagePermissionResponse")
	proto.RegisterType((*BatchUpdatePackagePermissionResponseWrapper)(nil), "pkg.BatchUpdatePackagePermissionResponseWrapper")
}

func init() {
	proto.RegisterFile("batch_update_package_permission.proto", fileDescriptor_7a8e4008e6164ed3)
}

var fileDescriptor_7a8e4008e6164ed3 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x95, 0xae, 0x1b, 0xaa, 0xc3, 0xcb, 0x66, 0xa1, 0x29, 0xda, 0x25, 0x95, 0x57, 0x50,
	0x27, 0xe6, 0x76, 0x6c, 0xd3, 0x54, 0xb8, 0xa0, 0x55, 0xe2, 0xc0, 0x05, 0x4d, 0x16, 0x08, 0x69,
	0x63, 0xab, 0xdc, 0xc6, 0x4b, 0xa3, 0x36, 0xb5, 0xb1, 0x1d, 0x8a, 0x3a, 0xf5, 0x86, 0x38, 0x20,
	0xf1, 0x15, 0x83, 0xc4, 0x07, 0xe0, 0x90, 0x4f, 0x80, 0xe2, 0x64, 0xad, 0xe1, 0x80, 0xca, 0x81,
	0x53, 0xfc, 0x3c, 0xff, 0xe7, 0xe5, 0x9f, 0x9f, 0x6c, 0xf0, 0xa8, 0x4f, 0xf5, 0x60, 0xd8, 0x4b,
	0x44, 0x40, 0x35, 0xeb, 0x09, 0x3a, 0x18, 0xd1, 0x90, 0xf5, 0x04, 0x93, 0x71, 0xa4, 0x54, 0xc4,
	0x27, 0x2d, 0x21, 0xb9, 0xe6, 0x70, 0x4d, 0x8c, 0xc2, 0x1d, 0x1c, 0x46, 0x7a, 0x98, 0xf4, 0x5b,
	0x03, 0x1e, 0xb7, 0x43, 0x1e, 0xf2, 0xb6, 0xd1, 0xfa, 0xc9, 0xb5, 0x89, 0x4c, 0x60, 0x4e, 0x45,
	0xcf, 0xce, 0x89, 0x55, 0x1e, 0x4f, 0x23, 0x3d, 0xe2, 0xd3, 0x76, 0xc8, 0xb1, 0x11, 0xf1, 0x47,
	0x3a, 0x8e, 0x02, 0xaa, 0xb9, 0x54, 0xed, 0xc5, 0xb1, 0xe8, 0x43, 0x3f, 0x2b, 0x60, 0xb7, 0x9b,
	0xbb, 0x7a, 0x6b, 0x4c, 0x9d, 0x15, 0x9e, 0xce, 0x16, 0x96, 0x08, 0xfb, 0x90, 0x30, 0xa5, 0x61,
	0x07, 0xb8, 0x4b, 0x9f, 0xca, 0x73, 0xea, 0x6b, 0xcd, 0x5a, 0x77, 0x3b, 0x4b, 0x7d, 0x78, 0xcd,
	0x65, 0xfc, 0x1c, 0x59, 0x22, 0x22, 0x76, 0x29, 0x24, 0xc0, 0xa5, 0x89, 0x1e, 0x72, 0x19, 0xcd,
	0x98, 0x54, 0x5e, 0xc5, 0x74, 0x1e, 0x2c, 0x3b, 0x2d, 0x11, 0xfd, 0xf8, 0xee, 0x6f, 0x83, 0x87,
	0x57, 0x17, 0xa7, 0xf8, 0x9c, 0xe2, 0x59, 0x0f, 0x5f, 0xbe, 0x9f, 0xde, 0x1c, 0xed, 0x9f, 0x1c,
	0xcf, 0x1b, 0xc4, 0x1e, 0x02, 0xf7, 0xc0, 0x46, 0xcc, 0xf4, 0x90, 0x07, 0xde, 0x5a, 0xdd, 0x69,
	0xd6, 0xba, 0x5b, 0x59, 0xea, 0xdf, 0x2b, 0xc6, 0x15, 0x79, 0x44, 0xca, 0x02, 0xf8, 0xcd, 0x01,
	0xa0, 0x24, 0xfd, 0x2a, 0x50, 0x5e, 0xd5, 0xac, 0x8f, 0xb3, 0xd4, 0xdf, 0x2a, 0x8d, 0x2f, 0xb4,
	0x7c, 0xfb, 0x1b, 0x40, 0xae, 0x2e, 0x5e, 0x50, 0x3c, 0x3b, 0xc5, 0xe7, 0x07, 0xf8, 0xd9, 0xe5,
	0x4d, 0x67, 0x8e, 0x7f, 0x8b, 0x8f, 0xff, 0x31, 0x7e, 0x7a, 0x38, 0x6f, 0x10, 0xcb, 0x00, 0xfa,
	0x52, 0x01, 0x8d, 0xbf, 0x03, 0x57, 0x82, 0x4f, 0x14, 0x83, 0xbb, 0xa0, 0x3a, 0xe0, 0x01, 0xf3,
	0x9c, 0xba, 0xd3, 0x5c, 0xef, 0x3e, 0xc8, 0x52, 0xdf, 0x2d, 0x1c, 0xe7, 0x59, 0x44, 0x8c, 0x08,
	0xf7, 0xc1, 0x9d, 0x98, 0x29, 0x45, 0x43, 0xe6, 0x55, 0x0c, 0x09, 0x98, 0xa5, 0xfe, 0xfd, 0x5b,
	0x12, 0x46, 0x40, 0xe4, 0xb6, 0x04, 0x7e, 0x75, 0x40, 0x6d, 0x61, 0xa5, 0x44, 0x37, 0xce, 0x52,
	0x7f, 0xf3, 0x0f, 0x14, 0xff, 0x8f, 0xc4, 0x72, 0x3d, 0xfa, 0x5c, 0x01, 0x4f, 0x56, 0x01, 0xf1,
	0x4e, 0x52, 0x21, 0x98, 0x5c, 0x8d, 0x47, 0x07, 0xb8, 0xf9, 0xf7, 0xe5, 0x27, 0x31, 0xa6, 0xd1,
	0xa4, 0x64, 0x62, 0x5d, 0x53, 0x4b, 0x44, 0xc4, 0x2e, 0x85, 0x8f, 0xc1, 0x3a, 0x93, 0x92, 0xcb,
	0x12, 0xcb, 0x66, 0x96, 0xfa, 0x77, 0x8b, 0x1e, 0x93, 0x46, 0xa4, 0x90, 0xe1, 0x6b, 0x50, 0x0d,
	0xa8, 0xa6, 0xe6, 0x22, 0xb9, 0x87, 0x7b, 0x2d, 0x31, 0x0a, 0x5b, 0xab, 0xfc, 0x86, 0xed, 0x38,
	0x1f, 0x80, 0x88, 0x99, 0xd3, 0xdf, 0x30, 0xef, 0xf0, 0xe8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x6e, 0x06, 0x15, 0x1c, 0x04, 0x00, 0x00,
}
