// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_workload.proto

package workload

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
//DeleteWorkload请求
type DeleteWorkloadRequest struct {
	//
	//要更新的 workload id
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteWorkloadRequest) Reset()         { *m = DeleteWorkloadRequest{} }
func (m *DeleteWorkloadRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWorkloadRequest) ProtoMessage()    {}
func (*DeleteWorkloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c91213e07aa9ad5, []int{0}
}
func (m *DeleteWorkloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWorkloadRequest.Unmarshal(m, b)
}
func (m *DeleteWorkloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWorkloadRequest.Marshal(b, m, deterministic)
}
func (m *DeleteWorkloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWorkloadRequest.Merge(m, src)
}
func (m *DeleteWorkloadRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWorkloadRequest.Size(m)
}
func (m *DeleteWorkloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWorkloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWorkloadRequest proto.InternalMessageInfo

func (m *DeleteWorkloadRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//DeleteWorkloadApi返回
type DeleteWorkloadResponseWrapper struct {
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

func (m *DeleteWorkloadResponseWrapper) Reset()         { *m = DeleteWorkloadResponseWrapper{} }
func (m *DeleteWorkloadResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteWorkloadResponseWrapper) ProtoMessage()    {}
func (*DeleteWorkloadResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c91213e07aa9ad5, []int{1}
}
func (m *DeleteWorkloadResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWorkloadResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteWorkloadResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWorkloadResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteWorkloadResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWorkloadResponseWrapper.Merge(m, src)
}
func (m *DeleteWorkloadResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteWorkloadResponseWrapper.Size(m)
}
func (m *DeleteWorkloadResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWorkloadResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWorkloadResponseWrapper proto.InternalMessageInfo

func (m *DeleteWorkloadResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteWorkloadResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteWorkloadResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteWorkloadResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteWorkloadRequest)(nil), "workload.DeleteWorkloadRequest")
	proto.RegisterType((*DeleteWorkloadResponseWrapper)(nil), "workload.DeleteWorkloadResponseWrapper")
}

func init() { proto.RegisterFile("delete_workload.proto", fileDescriptor_1c91213e07aa9ad5) }

var fileDescriptor_1c91213e07aa9ad5 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xc9, 0xfb, 0xb6, 0x2f, 0xaf, 0x5b, 0xd1, 0xba, 0xd0, 0x12, 0x2a, 0x92, 0xb2, 0x8a,
	0xd4, 0x43, 0x12, 0xb5, 0x20, 0xfe, 0xb9, 0x15, 0x7b, 0xe8, 0x35, 0x97, 0x1e, 0x44, 0x65, 0xd3,
	0x6c, 0x63, 0x68, 0x92, 0x89, 0x9b, 0xad, 0xf5, 0x0f, 0x7e, 0xd5, 0x08, 0xfa, 0x0d, 0xf2, 0x09,
	0x24, 0x93, 0xd6, 0x06, 0x4f, 0x3b, 0x33, 0xcf, 0xf3, 0xdb, 0x79, 0x18, 0xd2, 0xf2, 0x44, 0x28,
	0x94, 0xb8, 0x5f, 0x80, 0x9c, 0x85, 0xc0, 0x3d, 0x2b, 0x91, 0xa0, 0x80, 0xfe, 0x5f, 0xf5, 0x1d,
	0xd3, 0x0f, 0xd4, 0xc3, 0xdc, 0xb5, 0x26, 0x10, 0xd9, 0x3e, 0xf8, 0x60, 0xa3, 0xc1, 0x9d, 0x4f,
	0xb1, 0xc3, 0x06, 0xab, 0x12, 0xec, 0x9c, 0x55, 0xec, 0xd1, 0x22, 0x50, 0x33, 0x58, 0xd8, 0x3e,
	0x98, 0x28, 0x9a, 0x4f, 0x3c, 0x0c, 0x3c, 0xae, 0x40, 0xa6, 0xf6, 0x4f, 0xb9, 0xe4, 0x76, 0x7d,
	0x00, 0x3f, 0x14, 0xeb, 0xdf, 0x45, 0x94, 0xa8, 0x97, 0x52, 0x64, 0x2e, 0x69, 0x5d, 0x63, 0xcc,
	0xf1, 0x32, 0x95, 0x23, 0x1e, 0xe7, 0x22, 0x55, 0x74, 0x44, 0x48, 0x10, 0xa7, 0x8a, 0xc7, 0x13,
	0x31, 0xf2, 0x74, 0xad, 0xab, 0xf5, 0x36, 0x06, 0x47, 0x79, 0x66, 0xec, 0x4c, 0x41, 0x46, 0x97,
	0x6c, 0xad, 0xb1, 0xcf, 0x0f, 0xa3, 0x49, 0xb6, 0xee, 0x6e, 0x8e, 0xcd, 0x0b, 0x6e, 0xbe, 0xde,
	0xbe, 0x9d, 0xf4, 0xdf, 0x0f, 0x9c, 0x0a, 0xcc, 0xbe, 0x34, 0xb2, 0xf7, 0x7b, 0x49, 0x9a, 0x40,
	0x9c, 0x8a, 0xb1, 0xe4, 0x49, 0x22, 0x24, 0xdd, 0x27, 0xb5, 0x09, 0x78, 0x02, 0xd7, 0xd4, 0x07,
	0xdb, 0x79, 0x66, 0x34, 0xca, 0x35, 0xc5, 0x94, 0x39, 0x28, 0xd2, 0x73, 0xd2, 0x28, 0xde, 0xe1,
	0x73, 0x12, 0xf2, 0x20, 0xd6, 0xff, 0x60, 0xa4, 0x76, 0x9e, 0x19, 0x74, 0xed, 0x5d, 0x8a, 0xcc,
	0xa9, 0x5a, 0xe9, 0x21, 0xa9, 0x0b, 0x29, 0x41, 0xea, 0x7f, 0x91, 0x69, 0xe6, 0x99, 0xb1, 0x59,
	0x32, 0x38, 0x66, 0x4e, 0x29, 0xd3, 0x2b, 0x52, 0xf3, 0xb8, 0xe2, 0x7a, 0xad, 0xab, 0xf5, 0x1a,
	0xa7, 0x6d, 0xab, 0x3c, 0x9c, 0xb5, 0x3a, 0x9c, 0x35, 0x2c, 0x0e, 0x57, 0x8d, 0x57, 0xb8, 0x99,
	0x83, 0x90, 0xfb, 0x0f, 0x6d, 0xfd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x53, 0x5a, 0xe5,
	0xf7, 0x01, 0x00, 0x00,
}
