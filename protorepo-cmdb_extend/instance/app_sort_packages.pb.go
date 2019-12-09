// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_sort_packages.proto

package instance

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
//AppSortPackages请求
type AppSortPackagesRequest struct {
	//
	//应用实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//程序包信息
	Packages             []*AppSortPackagesRequest_Packages `protobuf:"bytes,2,rep,name=packages,proto3" json:"packages" form:"packages"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *AppSortPackagesRequest) Reset()         { *m = AppSortPackagesRequest{} }
func (m *AppSortPackagesRequest) String() string { return proto.CompactTextString(m) }
func (*AppSortPackagesRequest) ProtoMessage()    {}
func (*AppSortPackagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d53ccfec2ccb04e, []int{0}
}
func (m *AppSortPackagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSortPackagesRequest.Unmarshal(m, b)
}
func (m *AppSortPackagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSortPackagesRequest.Marshal(b, m, deterministic)
}
func (m *AppSortPackagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSortPackagesRequest.Merge(m, src)
}
func (m *AppSortPackagesRequest) XXX_Size() int {
	return xxx_messageInfo_AppSortPackagesRequest.Size(m)
}
func (m *AppSortPackagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSortPackagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppSortPackagesRequest proto.InternalMessageInfo

func (m *AppSortPackagesRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *AppSortPackagesRequest) GetPackages() []*AppSortPackagesRequest_Packages {
	if m != nil {
		return m.Packages
	}
	return nil
}

type AppSortPackagesRequest_Packages struct {
	//
	//包ID
	PackageId string `protobuf:"bytes,1,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//安装路径
	InstallPath          string   `protobuf:"bytes,2,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppSortPackagesRequest_Packages) Reset()         { *m = AppSortPackagesRequest_Packages{} }
func (m *AppSortPackagesRequest_Packages) String() string { return proto.CompactTextString(m) }
func (*AppSortPackagesRequest_Packages) ProtoMessage()    {}
func (*AppSortPackagesRequest_Packages) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d53ccfec2ccb04e, []int{0, 0}
}
func (m *AppSortPackagesRequest_Packages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSortPackagesRequest_Packages.Unmarshal(m, b)
}
func (m *AppSortPackagesRequest_Packages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSortPackagesRequest_Packages.Marshal(b, m, deterministic)
}
func (m *AppSortPackagesRequest_Packages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSortPackagesRequest_Packages.Merge(m, src)
}
func (m *AppSortPackagesRequest_Packages) XXX_Size() int {
	return xxx_messageInfo_AppSortPackagesRequest_Packages.Size(m)
}
func (m *AppSortPackagesRequest_Packages) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSortPackagesRequest_Packages.DiscardUnknown(m)
}

var xxx_messageInfo_AppSortPackagesRequest_Packages proto.InternalMessageInfo

func (m *AppSortPackagesRequest_Packages) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *AppSortPackagesRequest_Packages) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

//
//AppSortPackagesApi返回
type AppSortPackagesResponseWrapper struct {
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

func (m *AppSortPackagesResponseWrapper) Reset()         { *m = AppSortPackagesResponseWrapper{} }
func (m *AppSortPackagesResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AppSortPackagesResponseWrapper) ProtoMessage()    {}
func (*AppSortPackagesResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d53ccfec2ccb04e, []int{1}
}
func (m *AppSortPackagesResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSortPackagesResponseWrapper.Unmarshal(m, b)
}
func (m *AppSortPackagesResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSortPackagesResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AppSortPackagesResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSortPackagesResponseWrapper.Merge(m, src)
}
func (m *AppSortPackagesResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AppSortPackagesResponseWrapper.Size(m)
}
func (m *AppSortPackagesResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSortPackagesResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AppSortPackagesResponseWrapper proto.InternalMessageInfo

func (m *AppSortPackagesResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AppSortPackagesResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AppSortPackagesResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AppSortPackagesResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AppSortPackagesRequest)(nil), "instance.AppSortPackagesRequest")
	proto.RegisterType((*AppSortPackagesRequest_Packages)(nil), "instance.AppSortPackagesRequest.Packages")
	proto.RegisterType((*AppSortPackagesResponseWrapper)(nil), "instance.AppSortPackagesResponseWrapper")
}

func init() { proto.RegisterFile("app_sort_packages.proto", fileDescriptor_8d53ccfec2ccb04e) }

var fileDescriptor_8d53ccfec2ccb04e = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xae, 0x43, 0x9d, 0x83, 0x58, 0x31, 0xa8, 0x44, 0x45, 0x22, 0x95, 0x41, 0xa8,
	0xbb, 0x88, 0x03, 0x9d, 0x84, 0xf8, 0x73, 0x45, 0xa5, 0x5d, 0xf4, 0x6e, 0x32, 0x17, 0x48, 0x20,
	0x98, 0xdc, 0xc4, 0xcb, 0xa2, 0x25, 0x39, 0xc6, 0x71, 0xd9, 0x00, 0xf1, 0x60, 0xbc, 0x4c, 0x90,
	0x10, 0x4f, 0x90, 0x27, 0x40, 0xb1, 0x9b, 0x36, 0xa8, 0xbb, 0x8a, 0xcf, 0xf9, 0xbe, 0xef, 0xf8,
	0x17, 0xeb, 0xa0, 0x07, 0x5c, 0xca, 0xb3, 0x12, 0x94, 0x3e, 0x93, 0x3c, 0xba, 0xe4, 0x89, 0x28,
	0xa9, 0x54, 0xa0, 0x01, 0x0f, 0xd2, 0xa2, 0xd4, 0xbc, 0x88, 0xc4, 0x38, 0x48, 0x52, 0x7d, 0xb1,
	0x5a, 0xd2, 0x08, 0xf2, 0x30, 0x81, 0x04, 0x42, 0x63, 0x58, 0xae, 0xce, 0x4d, 0x65, 0x0a, 0x73,
	0xb2, 0xc1, 0xf1, 0x8b, 0x8e, 0x3d, 0xbf, 0x4a, 0xf5, 0x25, 0x5c, 0x85, 0x09, 0x04, 0x46, 0x0c,
	0xbe, 0xf2, 0x2c, 0x8d, 0xb9, 0x06, 0x55, 0x86, 0x9b, 0xe3, 0x3a, 0xf7, 0x30, 0x01, 0x48, 0x32,
	0xb1, 0x9d, 0x2e, 0x72, 0xa9, 0xbf, 0x59, 0x91, 0xfc, 0xea, 0xa1, 0xd1, 0x5b, 0x29, 0xdf, 0x81,
	0xd2, 0xa7, 0x6b, 0x4e, 0x26, 0xbe, 0xac, 0x44, 0xa9, 0xf1, 0x02, 0xa1, 0x16, 0x75, 0x11, 0x7b,
	0xce, 0xc4, 0x99, 0x1e, 0xcc, 0x8f, 0xea, 0xca, 0xbf, 0x7b, 0x0e, 0x2a, 0x7f, 0x4d, 0xb6, 0x1a,
	0xf9, 0xf3, 0xdb, 0x1f, 0xa2, 0x3b, 0x9f, 0x3f, 0x3e, 0x0b, 0x5e, 0xf1, 0xe0, 0xfb, 0xa7, 0x1f,
	0xcf, 0x8f, 0x7f, 0x3e, 0x61, 0x9d, 0x30, 0xfe, 0x80, 0x06, 0xed, 0x2b, 0x78, 0xbd, 0xc9, 0xde,
	0xd4, 0x9d, 0x1d, 0xd1, 0x56, 0xa6, 0x37, 0x5f, 0x4f, 0xdb, 0x7a, 0x7e, 0xaf, 0xae, 0xfc, 0x43,
	0x7b, 0x67, 0x3b, 0x84, 0xb0, 0xcd, 0xbc, 0xf1, 0x35, 0x1a, 0xb4, 0x56, 0x3c, 0x43, 0x07, 0xeb,
	0xfe, 0x86, 0xf8, 0x7e, 0x5d, 0xf9, 0xc3, 0xff, 0xd2, 0x8b, 0x98, 0xb0, 0xad, 0x0d, 0xbf, 0x44,
	0xae, 0x41, 0xc9, 0xb2, 0x53, 0xae, 0x2f, 0xbc, 0x9e, 0x49, 0x8d, 0xea, 0xca, 0xc7, 0x9d, 0xff,
	0xb4, 0x22, 0x61, 0x5d, 0x2b, 0xf9, 0xeb, 0xa0, 0x47, 0x3b, 0xf0, 0xa5, 0x84, 0xa2, 0x14, 0xef,
	0x15, 0x97, 0x52, 0x28, 0xfc, 0x18, 0xf5, 0x23, 0x88, 0x85, 0x61, 0xd9, 0x9f, 0x1f, 0xd6, 0x95,
	0xef, 0xda, 0xa9, 0x4d, 0x97, 0x30, 0x23, 0x36, 0x04, 0xcd, 0xf7, 0xe4, 0x5a, 0x66, 0x3c, 0x2d,
	0x76, 0x09, 0x3a, 0x22, 0x61, 0x5d, 0x2b, 0x7e, 0x8a, 0xf6, 0x85, 0x52, 0xa0, 0xbc, 0x3d, 0x93,
	0x19, 0xd6, 0x95, 0x7f, 0xdb, 0x66, 0x4c, 0x9b, 0x30, 0x2b, 0xe3, 0x37, 0xa8, 0x1f, 0x73, 0xcd,
	0xbd, 0xfe, 0xc4, 0x99, 0xba, 0xb3, 0x11, 0xb5, 0x1b, 0x41, 0xdb, 0x8d, 0xa0, 0x27, 0xcd, 0x46,
	0x74, 0xf1, 0x1a, 0x37, 0x61, 0x26, 0xb4, 0xbc, 0x65, 0x6c, 0xc7, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x83, 0x5c, 0x3a, 0xd2, 0x02, 0x00, 0x00,
}
