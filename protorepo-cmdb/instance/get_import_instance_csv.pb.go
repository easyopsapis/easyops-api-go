// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_import_instance_csv.proto

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
//GetImportInstanceCsv请求
type GetImportInstanceCsvRequest struct {
	//
	//模型对象ID
	ObjectId             string   `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImportInstanceCsvRequest) Reset()         { *m = GetImportInstanceCsvRequest{} }
func (m *GetImportInstanceCsvRequest) String() string { return proto.CompactTextString(m) }
func (*GetImportInstanceCsvRequest) ProtoMessage()    {}
func (*GetImportInstanceCsvRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc77b48f8c9b27ef, []int{0}
}
func (m *GetImportInstanceCsvRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImportInstanceCsvRequest.Unmarshal(m, b)
}
func (m *GetImportInstanceCsvRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImportInstanceCsvRequest.Marshal(b, m, deterministic)
}
func (m *GetImportInstanceCsvRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImportInstanceCsvRequest.Merge(m, src)
}
func (m *GetImportInstanceCsvRequest) XXX_Size() int {
	return xxx_messageInfo_GetImportInstanceCsvRequest.Size(m)
}
func (m *GetImportInstanceCsvRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImportInstanceCsvRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImportInstanceCsvRequest proto.InternalMessageInfo

func (m *GetImportInstanceCsvRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetImportInstanceCsvRequest)(nil), "instance.GetImportInstanceCsvRequest")
}

func init() { proto.RegisterFile("get_import_instance_csv.proto", fileDescriptor_dc77b48f8c9b27ef) }

var fileDescriptor_dc77b48f8c9b27ef = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4f, 0x2d, 0x89,
	0xcf, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0x89, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x8d,
	0x4f, 0x2e, 0x2e, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x49, 0xe9, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83,
	0x15, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x28, 0x65, 0x86, 0xa4, 0x3c,
	0xb7, 0x3c, 0xb3, 0x24, 0x3b, 0xbf, 0x5c, 0x3f, 0x3d, 0x5f, 0x17, 0x2c, 0xa9, 0x5b, 0x96, 0x98,
	0x93, 0x99, 0x92, 0x58, 0x92, 0x5f, 0x54, 0xac, 0x0f, 0x67, 0x42, 0xf5, 0x49, 0xa7, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0x22, 0x4c, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0x15, 0x73, 0x49,
	0xbb, 0xa7, 0x96, 0x78, 0x82, 0x5d, 0xeb, 0x09, 0x75, 0x98, 0x73, 0x71, 0x59, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x50, 0x08, 0x17, 0x67, 0x7e, 0x52, 0x56, 0x6a, 0x72, 0x49, 0x7c, 0x66,
	0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xf9, 0xa7, 0x7b, 0xf2, 0x02, 0x69, 0xf9, 0x45,
	0xb9, 0x56, 0x4a, 0x70, 0x29, 0xa5, 0x47, 0xf7, 0xe5, 0xe5, 0xb9, 0x64, 0xe3, 0xa2, 0x13, 0x75,
	0xab, 0x1c, 0x75, 0xa3, 0xe2, 0x63, 0xa3, 0x0d, 0x74, 0x2d, 0x61, 0xec, 0x6a, 0x03, 0x1d, 0x63,
	0xc3, 0x5a, 0x95, 0x20, 0x0e, 0x88, 0x72, 0xcf, 0x94, 0x24, 0x36, 0xb0, 0xdd, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x21, 0xc2, 0x16, 0x2a, 0x01, 0x00, 0x00,
}
