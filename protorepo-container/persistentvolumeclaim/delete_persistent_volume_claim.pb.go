// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_persistent_volume_claim.proto

package persistentvolumeclaim

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
//DeletePersistentVolumeClaims请求
type DeletePersistentVolumeClaimsRequest struct {
	//
	//要删除的实例 id
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePersistentVolumeClaimsRequest) Reset()         { *m = DeletePersistentVolumeClaimsRequest{} }
func (m *DeletePersistentVolumeClaimsRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePersistentVolumeClaimsRequest) ProtoMessage()    {}
func (*DeletePersistentVolumeClaimsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f347078992114f93, []int{0}
}
func (m *DeletePersistentVolumeClaimsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePersistentVolumeClaimsRequest.Unmarshal(m, b)
}
func (m *DeletePersistentVolumeClaimsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePersistentVolumeClaimsRequest.Marshal(b, m, deterministic)
}
func (m *DeletePersistentVolumeClaimsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePersistentVolumeClaimsRequest.Merge(m, src)
}
func (m *DeletePersistentVolumeClaimsRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePersistentVolumeClaimsRequest.Size(m)
}
func (m *DeletePersistentVolumeClaimsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePersistentVolumeClaimsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePersistentVolumeClaimsRequest proto.InternalMessageInfo

func (m *DeletePersistentVolumeClaimsRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//DeletePersistentVolumeClaimsApi返回
type DeletePersistentVolumeClaimsResponseWrapper struct {
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

func (m *DeletePersistentVolumeClaimsResponseWrapper) Reset() {
	*m = DeletePersistentVolumeClaimsResponseWrapper{}
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*DeletePersistentVolumeClaimsResponseWrapper) ProtoMessage() {}
func (*DeletePersistentVolumeClaimsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f347078992114f93, []int{1}
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper.Unmarshal(m, b)
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper.Merge(m, src)
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper.Size(m)
}
func (m *DeletePersistentVolumeClaimsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePersistentVolumeClaimsResponseWrapper proto.InternalMessageInfo

func (m *DeletePersistentVolumeClaimsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeletePersistentVolumeClaimsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeletePersistentVolumeClaimsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeletePersistentVolumeClaimsResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeletePersistentVolumeClaimsRequest)(nil), "persistentvolumeclaim.DeletePersistentVolumeClaimsRequest")
	proto.RegisterType((*DeletePersistentVolumeClaimsResponseWrapper)(nil), "persistentvolumeclaim.DeletePersistentVolumeClaimsResponseWrapper")
}

func init() {
	proto.RegisterFile("delete_persistent_volume_claim.proto", fileDescriptor_f347078992114f93)
}

var fileDescriptor_f347078992114f93 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4e, 0xa3, 0x40,
	0x14, 0xc6, 0xc3, 0x6e, 0xbb, 0xc9, 0x4e, 0x37, 0xbb, 0x5d, 0x92, 0x6d, 0x48, 0xf7, 0x82, 0x66,
	0xda, 0x98, 0x1a, 0x03, 0xa8, 0x4d, 0x8c, 0x7f, 0xee, 0xaa, 0xbd, 0xe8, 0x9d, 0xe1, 0x42, 0x2f,
	0x8c, 0x36, 0x53, 0x98, 0x22, 0x11, 0x98, 0x71, 0x66, 0x68, 0xfd, 0x13, 0x5f, 0x15, 0x13, 0xdf,
	0x40, 0x9e, 0xc0, 0x70, 0xa8, 0x85, 0x2b, 0xaf, 0x98, 0x73, 0x7e, 0xdf, 0x77, 0xce, 0xc7, 0x41,
	0x03, 0x9f, 0x46, 0x54, 0xd1, 0x19, 0xa7, 0x42, 0x86, 0x52, 0xd1, 0x44, 0xcd, 0x96, 0x2c, 0x4a,
	0x63, 0x3a, 0xf3, 0x22, 0x12, 0xc6, 0x36, 0x17, 0x4c, 0x31, 0xfd, 0x5f, 0x85, 0x4b, 0x0a, 0xb0,
	0x6b, 0x05, 0xa1, 0xba, 0x4d, 0xe7, 0xb6, 0xc7, 0x62, 0x27, 0x60, 0x01, 0x73, 0x40, 0x3d, 0x4f,
	0x17, 0x50, 0x41, 0x01, 0xaf, 0x72, 0x4a, 0xf7, 0xa0, 0x26, 0x8f, 0x57, 0xa1, 0xba, 0x63, 0x2b,
	0x27, 0x60, 0x16, 0x40, 0x6b, 0x49, 0xa2, 0xd0, 0x27, 0x8a, 0x09, 0xe9, 0x6c, 0x9e, 0x6b, 0xdf,
	0xff, 0x80, 0xb1, 0x20, 0xa2, 0xd5, 0x74, 0x1a, 0x73, 0xf5, 0x58, 0x42, 0xcc, 0x51, 0xff, 0x0c,
	0x7e, 0xe1, 0x7c, 0x13, 0xf1, 0x02, 0x22, 0x9e, 0x16, 0x11, 0xa5, 0x4b, 0xef, 0x53, 0x2a, 0x95,
	0x3e, 0x45, 0x28, 0x4c, 0xa4, 0x22, 0x89, 0x47, 0xa7, 0xbe, 0xa1, 0xf5, 0xb4, 0xe1, 0xcf, 0xf1,
	0x76, 0x9e, 0x99, 0x7f, 0x17, 0x4c, 0xc4, 0xc7, 0xb8, 0x62, 0xf8, 0xed, 0xd5, 0x6c, 0xa3, 0xdf,
	0x37, 0x57, 0xbb, 0xd6, 0x11, 0xb1, 0x9e, 0xae, 0x9f, 0xf7, 0x46, 0x2f, 0x03, 0xb7, 0x66, 0xc6,
	0xef, 0x1a, 0xda, 0xf9, 0x7a, 0xa5, 0xe4, 0x2c, 0x91, 0xf4, 0x52, 0x10, 0xce, 0xa9, 0xd0, 0xfb,
	0xa8, 0xe1, 0x31, 0x9f, 0xc2, 0xd2, 0xe6, 0xf8, 0x4f, 0x9e, 0x99, 0xad, 0x72, 0x69, 0xd1, 0xc5,
	0x2e, 0x40, 0xfd, 0x10, 0xb5, 0x8a, 0xef, 0xe4, 0x81, 0x47, 0x24, 0x4c, 0x8c, 0x6f, 0x10, 0xb0,
	0x93, 0x67, 0xa6, 0x5e, 0x69, 0xd7, 0x10, 0xbb, 0x75, 0xa9, 0xbe, 0x85, 0x9a, 0x54, 0x08, 0x26,
	0x8c, 0xef, 0xe0, 0x69, 0xe7, 0x99, 0xf9, 0xab, 0xf4, 0x40, 0x1b, 0xbb, 0x25, 0xd6, 0x4f, 0x50,
	0xc3, 0x27, 0x8a, 0x18, 0x8d, 0x9e, 0x36, 0x6c, 0xed, 0x77, 0xec, 0xf2, 0xa8, 0xf6, 0xe7, 0x51,
	0xed, 0x49, 0x71, 0xd4, 0x7a, 0xbc, 0x42, 0x8d, 0x5d, 0x30, 0xcd, 0x7f, 0x80, 0x6c, 0xf4, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x94, 0x8b, 0x8d, 0x5b, 0x2f, 0x02, 0x00, 0x00,
}
