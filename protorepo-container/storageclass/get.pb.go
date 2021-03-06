// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package storageclass

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	container "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
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
//Get请求
type GetRequest struct {
	//
	//StorageClass ID
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//Get返回
type GetResponse struct {
	//
	//关联的集群
	Cluster *container.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster" form:"cluster"`
	//
	//StorageClass
	StorageClass         *container.StorageClass `protobuf:"bytes,2,opt,name=storageClass,proto3" json:"storageClass" form:"storageClass"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetCluster() *container.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *GetResponse) GetStorageClass() *container.StorageClass {
	if m != nil {
		return m.StorageClass
	}
	return nil
}

//
//GetApi返回
type GetResponseWrapper struct {
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
	Data                 *GetResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponseWrapper) Reset()         { *m = GetResponseWrapper{} }
func (m *GetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetResponseWrapper) ProtoMessage()    {}
func (*GetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseWrapper.Unmarshal(m, b)
}
func (m *GetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseWrapper.Merge(m, src)
}
func (m *GetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetResponseWrapper.Size(m)
}
func (m *GetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseWrapper proto.InternalMessageInfo

func (m *GetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetResponseWrapper) GetData() *GetResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "storageclass.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "storageclass.GetResponse")
	proto.RegisterType((*GetResponseWrapper)(nil), "storageclass.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x66, 0x75, 0xab, 0xec, 0x6c, 0xa9, 0x75, 0x04, 0xbb, 0xf6, 0x26, 0x32, 0x8a, 0xe8, 0x45,
	0x12, 0xb5, 0x20, 0xea, 0x85, 0x48, 0x8a, 0x48, 0x2f, 0x1d, 0x85, 0xbd, 0x10, 0x95, 0xd9, 0xc9,
	0x69, 0x0c, 0x66, 0x73, 0xe2, 0xcc, 0xac, 0xf5, 0x07, 0x1f, 0xc7, 0xa7, 0x12, 0x22, 0xf8, 0x08,
	0x79, 0x02, 0xd9, 0x33, 0xd3, 0xee, 0x78, 0xdf, 0xab, 0xcc, 0xc9, 0xf7, 0x93, 0xef, 0xcb, 0x19,
	0x36, 0xa9, 0xc0, 0x65, 0x9d, 0x41, 0x87, 0x7c, 0xdb, 0x3a, 0x34, 0xaa, 0x02, 0xdd, 0x28, 0x6b,
	0xf7, 0xd3, 0xaa, 0x76, 0x1f, 0x57, 0x8b, 0x4c, 0xe3, 0x32, 0xaf, 0xb0, 0xc2, 0x9c, 0x48, 0x8b,
	0xd5, 0x31, 0x4d, 0x34, 0xd0, 0xc9, 0x8b, 0xf7, 0x5f, 0x55, 0x98, 0x81, 0xb2, 0xdf, 0xb0, 0xb3,
	0x59, 0x83, 0x5a, 0x35, 0xb9, 0xc6, 0xd6, 0x19, 0xa5, 0x9d, 0xf5, 0x4a, 0x03, 0x1d, 0xa6, 0x4b,
	0x2c, 0xa1, 0xb1, 0x79, 0x20, 0xe6, 0x34, 0x12, 0x51, 0xd5, 0x2d, 0x98, 0x5c, 0x37, 0x2b, 0xeb,
	0xc0, 0x04, 0xcb, 0xf9, 0x79, 0x58, 0x86, 0x4e, 0x1f, 0xa8, 0x54, 0x30, 0x7e, 0x14, 0x55, 0x5b,
	0x9e, 0xd4, 0xee, 0x13, 0x9e, 0xe4, 0x15, 0xa6, 0x04, 0xa6, 0x5f, 0x54, 0x53, 0x97, 0xca, 0xa1,
	0xb1, 0xf9, 0xd9, 0xd1, 0xeb, 0xc4, 0x9c, 0xb1, 0x97, 0xe0, 0x24, 0x7c, 0x5e, 0x81, 0x75, 0xfc,
	0x88, 0xb1, 0xba, 0xb5, 0x4e, 0xb5, 0x1a, 0x8e, 0xca, 0xd9, 0xe8, 0xe6, 0xe8, 0xee, 0xa4, 0xb8,
	0x37, 0xf4, 0xc9, 0xd5, 0x63, 0x34, 0xcb, 0xa7, 0x62, 0x83, 0x89, 0xbf, 0x7f, 0x92, 0x5d, 0xb6,
	0xf3, 0xfe, 0xed, 0xfd, 0xf4, 0x89, 0x4a, 0xbf, 0xbf, 0xfb, 0xf1, 0xe0, 0xe0, 0xe7, 0x6d, 0x19,
	0x89, 0xc5, 0xaf, 0x11, 0x9b, 0x92, 0xb3, 0xed, 0xb0, 0xb5, 0xc0, 0x9f, 0xb3, 0xcb, 0xe1, 0x57,
	0x90, 0xef, 0xf4, 0x21, 0xcf, 0xce, 0x1a, 0x65, 0x87, 0x1e, 0x29, 0xf8, 0xd0, 0x27, 0x3b, 0xfe,
	0x5b, 0x81, 0x2c, 0xe4, 0xa9, 0x8c, 0xbf, 0x61, 0xa7, 0xdb, 0x3c, 0x5c, 0x17, 0x9f, 0x5d, 0x20,
	0x9b, 0xbd, 0xc8, 0xe6, 0x75, 0x04, 0x17, 0x7b, 0x43, 0x9f, 0x5c, 0xf3, 0x5e, 0xb1, 0x4c, 0xc8,
	0xff, 0x5c, 0xc4, 0xef, 0x11, 0xe3, 0x51, 0xce, 0xb9, 0x51, 0x5d, 0x07, 0x86, 0xdf, 0x62, 0x63,
	0x8d, 0x25, 0x50, 0xd6, 0xad, 0xe2, 0xca, 0xd0, 0x27, 0xd3, 0x90, 0x0b, 0x4b, 0x10, 0x92, 0x40,
	0xfe, 0x98, 0x4d, 0xd7, 0xcf, 0x17, 0x5f, 0xbb, 0x46, 0xd5, 0x2d, 0x05, 0x9a, 0x14, 0xd7, 0x87,
	0x3e, 0xe1, 0x1b, 0x6e, 0x00, 0x85, 0x8c, 0xa9, 0xfc, 0x0e, 0xdb, 0x02, 0x63, 0xd0, 0xcc, 0x2e,
	0x92, 0x66, 0x77, 0xe8, 0x93, 0x6d, 0xaf, 0xa1, 0xd7, 0x42, 0x7a, 0x98, 0x3f, 0x63, 0xe3, 0x52,
	0x39, 0x35, 0x1b, 0x53, 0xd7, 0x1b, 0x59, 0x7c, 0x9d, 0xb3, 0x28, 0x76, 0x9c, 0x70, 0x2d, 0x10,
	0x92, 0x74, 0x8b, 0x4b, 0xb4, 0xe5, 0x83, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x9b, 0xd8,
	0xbd, 0x13, 0x03, 0x00, 0x00,
}
