// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cluster_delete_package.proto

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
//ClusterDeletePackage请求
type ClusterDeletePackageRequest struct {
	//
	//应用实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//程序包Id列表，用;分隔
	PackageIds string `protobuf:"bytes,2,opt,name=packageIds,proto3" json:"packageIds" form:"packageIds"`
	//
	//安装路径
	InstallPath          string   `protobuf:"bytes,3,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterDeletePackageRequest) Reset()         { *m = ClusterDeletePackageRequest{} }
func (m *ClusterDeletePackageRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterDeletePackageRequest) ProtoMessage()    {}
func (*ClusterDeletePackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2083440dd90276b, []int{0}
}
func (m *ClusterDeletePackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDeletePackageRequest.Unmarshal(m, b)
}
func (m *ClusterDeletePackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDeletePackageRequest.Marshal(b, m, deterministic)
}
func (m *ClusterDeletePackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDeletePackageRequest.Merge(m, src)
}
func (m *ClusterDeletePackageRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterDeletePackageRequest.Size(m)
}
func (m *ClusterDeletePackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDeletePackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDeletePackageRequest proto.InternalMessageInfo

func (m *ClusterDeletePackageRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *ClusterDeletePackageRequest) GetPackageIds() string {
	if m != nil {
		return m.PackageIds
	}
	return ""
}

func (m *ClusterDeletePackageRequest) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

//
//ClusterDeletePackageApi返回
type ClusterDeletePackageResponseWrapper struct {
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

func (m *ClusterDeletePackageResponseWrapper) Reset()         { *m = ClusterDeletePackageResponseWrapper{} }
func (m *ClusterDeletePackageResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ClusterDeletePackageResponseWrapper) ProtoMessage()    {}
func (*ClusterDeletePackageResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2083440dd90276b, []int{1}
}
func (m *ClusterDeletePackageResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDeletePackageResponseWrapper.Unmarshal(m, b)
}
func (m *ClusterDeletePackageResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDeletePackageResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ClusterDeletePackageResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDeletePackageResponseWrapper.Merge(m, src)
}
func (m *ClusterDeletePackageResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ClusterDeletePackageResponseWrapper.Size(m)
}
func (m *ClusterDeletePackageResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDeletePackageResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDeletePackageResponseWrapper proto.InternalMessageInfo

func (m *ClusterDeletePackageResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ClusterDeletePackageResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ClusterDeletePackageResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ClusterDeletePackageResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterDeletePackageRequest)(nil), "instance.ClusterDeletePackageRequest")
	proto.RegisterType((*ClusterDeletePackageResponseWrapper)(nil), "instance.ClusterDeletePackageResponseWrapper")
}

func init() { proto.RegisterFile("cluster_delete_package.proto", fileDescriptor_a2083440dd90276b) }

var fileDescriptor_a2083440dd90276b = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6f, 0x9b, 0x30,
	0x18, 0xc7, 0xc5, 0x96, 0x4c, 0x9b, 0x33, 0x6d, 0x99, 0xa5, 0x45, 0x28, 0x99, 0x44, 0xe4, 0x4c,
	0x53, 0x76, 0x00, 0xb6, 0x45, 0x9b, 0xb6, 0xf5, 0x96, 0x36, 0x87, 0xdc, 0x22, 0x5f, 0x7a, 0xa8,
	0xda, 0xc8, 0x01, 0x87, 0xa0, 0x00, 0xa6, 0xc6, 0x34, 0x7d, 0x51, 0x3f, 0x68, 0x2f, 0x54, 0xea,
	0xad, 0x57, 0x3e, 0x41, 0x85, 0x4d, 0x8a, 0x0f, 0x3d, 0x61, 0xfb, 0xff, 0xa2, 0xdf, 0xc3, 0x03,
	0xbe, 0x78, 0x51, 0x9e, 0x09, 0xca, 0x97, 0x3e, 0x8d, 0xa8, 0xa0, 0xcb, 0x94, 0x78, 0x5b, 0x12,
	0x50, 0x27, 0xe5, 0x4c, 0x30, 0xf8, 0x36, 0x4c, 0x32, 0x41, 0x12, 0x8f, 0xf6, 0xed, 0x20, 0x14,
	0x9b, 0x7c, 0xe5, 0x78, 0x2c, 0x76, 0x03, 0x16, 0x30, 0x57, 0x1a, 0x56, 0xf9, 0x5a, 0xde, 0xe4,
	0x45, 0x9e, 0x54, 0xb0, 0xff, 0x47, 0xb3, 0xc7, 0xbb, 0x50, 0x6c, 0xd9, 0xce, 0x0d, 0x98, 0x2d,
	0x45, 0xfb, 0x82, 0x44, 0xa1, 0x4f, 0x04, 0xe3, 0x99, 0xfb, 0x7c, 0xac, 0x73, 0x83, 0x80, 0xb1,
	0x20, 0xa2, 0x4d, 0x3b, 0x8d, 0x53, 0x71, 0xa5, 0x44, 0x74, 0x67, 0x80, 0xc1, 0xa1, 0xc2, 0x3d,
	0x92, 0xb4, 0x0b, 0x05, 0x8b, 0xe9, 0x79, 0x4e, 0x33, 0x01, 0xe7, 0x00, 0xec, 0x79, 0xe7, 0xbe,
	0x69, 0x0c, 0x8d, 0xf1, 0xbb, 0xe9, 0xf7, 0xb2, 0xb0, 0x3e, 0xad, 0x19, 0x8f, 0xff, 0xa3, 0x46,
	0x43, 0x0f, 0xf7, 0x56, 0x17, 0x7c, 0x38, 0x3b, 0xf9, 0x61, 0xff, 0x23, 0xf6, 0xf5, 0xe9, 0xcd,
	0xcf, 0xc9, 0xed, 0x57, 0xac, 0x85, 0xe1, 0x6f, 0x00, 0xea, 0x3f, 0x31, 0xf7, 0x33, 0xf3, 0x95,
	0xac, 0xfa, 0xdc, 0x54, 0x35, 0x1a, 0xc2, 0x9a, 0x11, 0xfe, 0x05, 0x1d, 0x59, 0x12, 0x45, 0x0b,
	0x22, 0x36, 0xe6, 0x6b, 0x99, 0xeb, 0x95, 0x85, 0x05, 0x35, 0x04, 0x25, 0x22, 0xac, 0x5b, 0xd1,
	0xa3, 0x01, 0x46, 0x2f, 0xcf, 0x96, 0xa5, 0x2c, 0xc9, 0xe8, 0x31, 0x27, 0x69, 0x4a, 0x39, 0x1c,
	0x81, 0x96, 0xc7, 0x7c, 0x2a, 0xa7, 0x6b, 0x4f, 0x3f, 0x96, 0x85, 0xd5, 0x51, 0xd5, 0xd5, 0x2b,
	0xc2, 0x52, 0xac, 0x30, 0xaa, 0xef, 0xec, 0x32, 0x8d, 0x48, 0x98, 0xd4, 0xf8, 0x1a, 0x86, 0x26,
	0x22, 0xac, 0x5b, 0xe1, 0x37, 0xd0, 0xa6, 0x9c, 0x33, 0x5e, 0xa3, 0x77, 0xcb, 0xc2, 0x7a, 0xaf,
	0x32, 0xf2, 0x19, 0x61, 0x25, 0xc3, 0x03, 0xd0, 0xf2, 0x89, 0x20, 0x66, 0x6b, 0x68, 0x8c, 0x3b,
	0xbf, 0x7a, 0x8e, 0x5a, 0x9b, 0xb3, 0x5f, 0x9b, 0x33, 0xab, 0xd6, 0xa6, 0xe3, 0x55, 0x6e, 0x84,
	0x65, 0x68, 0xf5, 0x46, 0xda, 0x26, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xec, 0xb7, 0xb7,
	0x7c, 0x02, 0x00, 0x00,
}
