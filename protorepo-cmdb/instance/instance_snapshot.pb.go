// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance_snapshot.proto

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
//InstanceSnapshot请求
type InstanceSnapshotRequest struct {
	//
	//模型对象ID
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//实例ID
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id" form:"instance_id"`
	//
	//版本号
	XVersion int32 `protobuf:"varint,3,opt,name=_version,json=Version,proto3" json:"_version" form:"_version"`
	//
	//操作时间
	XTs                  int32    `protobuf:"varint,4,opt,name=_ts,json=Ts,proto3" json:"_ts" form:"_ts"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceSnapshotRequest) Reset()         { *m = InstanceSnapshotRequest{} }
func (m *InstanceSnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*InstanceSnapshotRequest) ProtoMessage()    {}
func (*InstanceSnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33661c3f72e1b68c, []int{0}
}
func (m *InstanceSnapshotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceSnapshotRequest.Unmarshal(m, b)
}
func (m *InstanceSnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceSnapshotRequest.Marshal(b, m, deterministic)
}
func (m *InstanceSnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceSnapshotRequest.Merge(m, src)
}
func (m *InstanceSnapshotRequest) XXX_Size() int {
	return xxx_messageInfo_InstanceSnapshotRequest.Size(m)
}
func (m *InstanceSnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceSnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceSnapshotRequest proto.InternalMessageInfo

func (m *InstanceSnapshotRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *InstanceSnapshotRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *InstanceSnapshotRequest) GetXVersion() int32 {
	if m != nil {
		return m.XVersion
	}
	return 0
}

func (m *InstanceSnapshotRequest) GetXTs() int32 {
	if m != nil {
		return m.XTs
	}
	return 0
}

//
//InstanceSnapshotApi返回
type InstanceSnapshotResponseWrapper struct {
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
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InstanceSnapshotResponseWrapper) Reset()         { *m = InstanceSnapshotResponseWrapper{} }
func (m *InstanceSnapshotResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InstanceSnapshotResponseWrapper) ProtoMessage()    {}
func (*InstanceSnapshotResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_33661c3f72e1b68c, []int{1}
}
func (m *InstanceSnapshotResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceSnapshotResponseWrapper.Unmarshal(m, b)
}
func (m *InstanceSnapshotResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceSnapshotResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InstanceSnapshotResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceSnapshotResponseWrapper.Merge(m, src)
}
func (m *InstanceSnapshotResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InstanceSnapshotResponseWrapper.Size(m)
}
func (m *InstanceSnapshotResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceSnapshotResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceSnapshotResponseWrapper proto.InternalMessageInfo

func (m *InstanceSnapshotResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InstanceSnapshotResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InstanceSnapshotResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InstanceSnapshotResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceSnapshotRequest)(nil), "instance.InstanceSnapshotRequest")
	proto.RegisterType((*InstanceSnapshotResponseWrapper)(nil), "instance.InstanceSnapshotResponseWrapper")
}

func init() { proto.RegisterFile("instance_snapshot.proto", fileDescriptor_33661c3f72e1b68c) }

var fileDescriptor_33661c3f72e1b68c = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x71, 0x9a, 0x6c, 0x89, 0x3c, 0xd6, 0xa0, 0x8b, 0x25, 0x94, 0x15, 0x15, 0xad, 0x8c,
	0x32, 0x66, 0xa7, 0x5d, 0x60, 0xff, 0xd8, 0x2e, 0x16, 0xd8, 0x45, 0xd8, 0x9d, 0x5b, 0x36, 0x58,
	0x69, 0x8d, 0x62, 0xab, 0xae, 0xb6, 0xc4, 0xc7, 0x93, 0xe4, 0x76, 0xac, 0xf4, 0x89, 0xf6, 0x4e,
	0x1e, 0xec, 0x62, 0x0f, 0xe0, 0x17, 0x58, 0xb1, 0x64, 0xa7, 0x86, 0xfa, 0xc6, 0x47, 0xfa, 0x7d,
	0xdf, 0xc7, 0xd1, 0xe1, 0xa0, 0x91, 0x48, 0x95, 0x66, 0x69, 0xc4, 0x43, 0x95, 0xb2, 0x4c, 0x9d,
	0x83, 0xf6, 0x33, 0x09, 0x1a, 0x70, 0xbf, 0x01, 0x5b, 0x5e, 0x22, 0xf4, 0x79, 0xbe, 0xf0, 0x23,
	0x58, 0x4d, 0x12, 0x48, 0x60, 0x62, 0x04, 0x8b, 0xfc, 0xcc, 0x9c, 0xcc, 0xc1, 0x54, 0xd6, 0xb8,
	0xf5, 0xb2, 0x25, 0x5f, 0x5d, 0x0a, 0xfd, 0x1d, 0x2e, 0x27, 0x09, 0x78, 0x06, 0x7a, 0x17, 0x6c,
	0x29, 0x62, 0xa6, 0x41, 0xaa, 0xc9, 0xba, 0xac, 0x7d, 0x8f, 0x13, 0x80, 0x64, 0xc9, 0x6f, 0xd3,
	0x95, 0x96, 0x79, 0x54, 0xb7, 0x43, 0x7f, 0x77, 0xd0, 0x68, 0x5e, 0x77, 0x74, 0x58, 0x77, 0x1a,
	0xf0, 0x1f, 0x39, 0x57, 0x1a, 0x1f, 0xa1, 0x01, 0x2c, 0xbe, 0xf1, 0x48, 0x87, 0x22, 0x1e, 0x3b,
	0x3b, 0xce, 0xde, 0x60, 0xf6, 0xaa, 0x2c, 0xc8, 0xf0, 0x0c, 0xe4, 0xea, 0x2d, 0x5d, 0x23, 0xfa,
	0xf7, 0x0f, 0x21, 0x68, 0xfb, 0xf4, 0x98, 0x79, 0xbf, 0x3e, 0x78, 0x5f, 0xc3, 0x93, 0xe3, 0x7d,
	0xef, 0x4d, 0x53, 0x5f, 0xed, 0x3f, 0x9f, 0x1e, 0x5c, 0xef, 0x06, 0x7d, 0x2b, 0x9f, 0xc7, 0xf8,
	0x13, 0x72, 0xd7, 0xb3, 0x11, 0xf1, 0xb8, 0x63, 0x72, 0x9f, 0x95, 0x05, 0xc1, 0x36, 0xb7, 0x05,
	0xab, 0xe4, 0x21, 0x7a, 0x78, 0x5a, 0x07, 0x9e, 0x5c, 0x1d, 0x4c, 0xaf, 0x77, 0x03, 0xd4, 0x28,
	0xe6, 0x31, 0x7e, 0x8f, 0xfa, 0xe1, 0x05, 0x97, 0x4a, 0x40, 0x3a, 0xde, 0xd8, 0x71, 0xf6, 0x7a,
	0x33, 0x5a, 0x16, 0x64, 0xd3, 0x26, 0x35, 0xa4, 0x8a, 0x71, 0x87, 0xff, 0x9b, 0xcf, 0x09, 0xee,
	0x7f, 0xb6, 0x00, 0xfb, 0x68, 0x23, 0xd4, 0x6a, 0xdc, 0x35, 0xce, 0xed, 0xb2, 0x20, 0xa8, 0x76,
	0x6a, 0x75, 0xc7, 0xd4, 0x39, 0x52, 0xf4, 0x9f, 0x83, 0xc8, 0xdd, 0x69, 0xa9, 0x0c, 0x52, 0xc5,
	0xbf, 0x48, 0x96, 0x65, 0x5c, 0xe2, 0x27, 0xa8, 0x1b, 0x41, 0xcc, 0xcd, 0xc0, 0x7a, 0xb3, 0xcd,
	0xb2, 0x20, 0xae, 0x0d, 0xad, 0x6e, 0x69, 0x60, 0x20, 0x7e, 0x8d, 0xdc, 0xea, 0xff, 0xf1, 0x67,
	0xb6, 0x64, 0x22, 0xad, 0x87, 0xf0, 0xe8, 0x76, 0x08, 0x2d, 0x48, 0x83, 0xb6, 0x14, 0x3f, 0x45,
	0x3d, 0x2e, 0x25, 0x48, 0xf3, 0xdc, 0xc1, 0x6c, 0x58, 0x16, 0xe4, 0x81, 0xf5, 0x98, 0x6b, 0x1a,
	0x58, 0x8c, 0xdf, 0xa1, 0x6e, 0xcc, 0x34, 0x33, 0x6f, 0x73, 0x5f, 0x8c, 0x7c, 0xbb, 0x05, 0x7e,
	0xb3, 0x05, 0xfe, 0xa1, 0xd9, 0x82, 0x76, 0x7f, 0x95, 0x9c, 0x06, 0xc6, 0xb5, 0xb8, 0x67, 0x74,
	0xd3, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x42, 0x72, 0x56, 0xc7, 0x02, 0x00, 0x00,
}
