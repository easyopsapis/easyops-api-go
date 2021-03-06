// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: export_instance_csv.proto

package instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
//ExportInstanceCsv请求
type ExportInstanceCsvRequest struct {
	//
	//模型对象ID
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//请求参数json序列化
	Json                 string   `protobuf:"bytes,2,opt,name=json,proto3" json:"json" form:"json"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportInstanceCsvRequest) Reset()         { *m = ExportInstanceCsvRequest{} }
func (m *ExportInstanceCsvRequest) String() string { return proto.CompactTextString(m) }
func (*ExportInstanceCsvRequest) ProtoMessage()    {}
func (*ExportInstanceCsvRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7d6af5b653859e, []int{0}
}
func (m *ExportInstanceCsvRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportInstanceCsvRequest.Unmarshal(m, b)
}
func (m *ExportInstanceCsvRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportInstanceCsvRequest.Marshal(b, m, deterministic)
}
func (m *ExportInstanceCsvRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportInstanceCsvRequest.Merge(m, src)
}
func (m *ExportInstanceCsvRequest) XXX_Size() int {
	return xxx_messageInfo_ExportInstanceCsvRequest.Size(m)
}
func (m *ExportInstanceCsvRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportInstanceCsvRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportInstanceCsvRequest proto.InternalMessageInfo

func (m *ExportInstanceCsvRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ExportInstanceCsvRequest) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*ExportInstanceCsvRequest)(nil), "instance.ExportInstanceCsvRequest")
}

func init() { proto.RegisterFile("export_instance_csv.proto", fileDescriptor_ac7d6af5b653859e) }

var fileDescriptor_ac7d6af5b653859e = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xad, 0x28, 0xc8,
	0x2f, 0x2a, 0x89, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x8d, 0x4f, 0x2e, 0x2e, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x49, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x15, 0x24, 0x95, 0xa6,
	0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x28, 0x65, 0x86, 0xa4, 0x3c, 0xb7, 0x3c, 0xb3, 0x24,
	0x3b, 0xbf, 0x5c, 0x3f, 0x3d, 0x5f, 0x17, 0x2c, 0xa9, 0x5b, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58,
	0x92, 0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0xf5, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x22,
	0x4c, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0x4d, 0x65, 0xe4, 0x92, 0x70, 0x05, 0xbb,
	0xd5, 0x13, 0xea, 0x2c, 0xe7, 0xe2, 0xb2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x10,
	0x2e, 0xce, 0xfc, 0xa4, 0xac, 0xd4, 0xe4, 0x92, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x4e, 0x27, 0xf3, 0x4f, 0xf7, 0xe4, 0x05, 0xd2, 0xf2, 0x8b, 0x72, 0xad, 0x94, 0xe0, 0x52, 0x4a,
	0x8f, 0xee, 0xcb, 0xcb, 0x73, 0xc9, 0xc6, 0x45, 0x27, 0xea, 0x56, 0x39, 0xea, 0x46, 0xc5, 0xc7,
	0x46, 0x1b, 0xe8, 0x5a, 0xc2, 0xd8, 0xd5, 0x06, 0x3a, 0xc6, 0x86, 0xb5, 0x2a, 0x41, 0x1c, 0x10,
	0xe5, 0x9e, 0x29, 0x42, 0xca, 0x5c, 0x2c, 0x59, 0xc5, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x03, 0xf9,
	0x3f, 0xdd, 0x93, 0xe7, 0x86, 0x18, 0x08, 0x12, 0x55, 0x0a, 0x02, 0x4b, 0x26, 0xb1, 0x81, 0x9d,
	0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x58, 0x86, 0x30, 0x49, 0x01, 0x00, 0x00,
}
