// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_relation_query_strategy_v2.proto

package cmdb_object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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
//ListRelationQueryStrategyV2Request请求
type ListRelationQueryStrategyV2RequestRequest struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//策略类型
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRelationQueryStrategyV2RequestRequest) Reset() {
	*m = ListRelationQueryStrategyV2RequestRequest{}
}
func (m *ListRelationQueryStrategyV2RequestRequest) String() string { return proto.CompactTextString(m) }
func (*ListRelationQueryStrategyV2RequestRequest) ProtoMessage()    {}
func (*ListRelationQueryStrategyV2RequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7341529250187e68, []int{0}
}
func (m *ListRelationQueryStrategyV2RequestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest.Unmarshal(m, b)
}
func (m *ListRelationQueryStrategyV2RequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest.Marshal(b, m, deterministic)
}
func (m *ListRelationQueryStrategyV2RequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest.Merge(m, src)
}
func (m *ListRelationQueryStrategyV2RequestRequest) XXX_Size() int {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest.Size(m)
}
func (m *ListRelationQueryStrategyV2RequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRelationQueryStrategyV2RequestRequest proto.InternalMessageInfo

func (m *ListRelationQueryStrategyV2RequestRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ListRelationQueryStrategyV2RequestRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

//
//ListRelationQueryStrategyV2Request返回
type ListRelationQueryStrategyV2RequestResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//页码
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回信息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//数据
	Data                 []*cmdb.RelationQueryStrategyV2 `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListRelationQueryStrategyV2RequestResponse) Reset() {
	*m = ListRelationQueryStrategyV2RequestResponse{}
}
func (m *ListRelationQueryStrategyV2RequestResponse) String() string {
	return proto.CompactTextString(m)
}
func (*ListRelationQueryStrategyV2RequestResponse) ProtoMessage() {}
func (*ListRelationQueryStrategyV2RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7341529250187e68, []int{1}
}
func (m *ListRelationQueryStrategyV2RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse.Unmarshal(m, b)
}
func (m *ListRelationQueryStrategyV2RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse.Marshal(b, m, deterministic)
}
func (m *ListRelationQueryStrategyV2RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse.Merge(m, src)
}
func (m *ListRelationQueryStrategyV2RequestResponse) XXX_Size() int {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse.Size(m)
}
func (m *ListRelationQueryStrategyV2RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRelationQueryStrategyV2RequestResponse proto.InternalMessageInfo

func (m *ListRelationQueryStrategyV2RequestResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListRelationQueryStrategyV2RequestResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListRelationQueryStrategyV2RequestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ListRelationQueryStrategyV2RequestResponse) GetData() []*cmdb.RelationQueryStrategyV2 {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//ListRelationQueryStrategyV2RequestApi返回
type ListRelationQueryStrategyV2RequestResponseWrapper struct {
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
	Data                 *ListRelationQueryStrategyV2RequestResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ListRelationQueryStrategyV2RequestResponseWrapper) Reset() {
	*m = ListRelationQueryStrategyV2RequestResponseWrapper{}
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*ListRelationQueryStrategyV2RequestResponseWrapper) ProtoMessage() {}
func (*ListRelationQueryStrategyV2RequestResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7341529250187e68, []int{2}
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper.Unmarshal(m, b)
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper.Merge(m, src)
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper.Size(m)
}
func (m *ListRelationQueryStrategyV2RequestResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListRelationQueryStrategyV2RequestResponseWrapper proto.InternalMessageInfo

func (m *ListRelationQueryStrategyV2RequestResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListRelationQueryStrategyV2RequestResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListRelationQueryStrategyV2RequestResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListRelationQueryStrategyV2RequestResponseWrapper) GetData() *ListRelationQueryStrategyV2RequestResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRelationQueryStrategyV2RequestRequest)(nil), "cmdb_object.ListRelationQueryStrategyV2RequestRequest")
	proto.RegisterType((*ListRelationQueryStrategyV2RequestResponse)(nil), "cmdb_object.ListRelationQueryStrategyV2RequestResponse")
	proto.RegisterType((*ListRelationQueryStrategyV2RequestResponseWrapper)(nil), "cmdb_object.ListRelationQueryStrategyV2RequestResponseWrapper")
}

func init() {
	proto.RegisterFile("list_relation_query_strategy_v2.proto", fileDescriptor_7341529250187e68)
}

var fileDescriptor_7341529250187e68 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x9b, 0x14, 0xe8, 0x1a, 0x41, 0xe5, 0x03, 0x8a, 0x22, 0x55, 0xae, 0x96, 0x1f, 0x05,
	0x54, 0x7b, 0xdb, 0x54, 0xa2, 0xc0, 0x8d, 0x4a, 0x1c, 0x90, 0xb8, 0xb0, 0x20, 0x90, 0x28, 0x60,
	0xad, 0xed, 0xad, 0x31, 0xd8, 0x19, 0x77, 0x77, 0xd3, 0x12, 0x10, 0xef, 0xc1, 0x13, 0xf0, 0x58,
	0x46, 0xe2, 0xca, 0xcd, 0x4f, 0x80, 0x76, 0xd7, 0x4d, 0x7d, 0x09, 0xf5, 0x69, 0x67, 0xe6, 0xfb,
	0xe6, 0x9b, 0x6f, 0x26, 0x31, 0xba, 0x5b, 0xe4, 0x52, 0x45, 0x82, 0x17, 0x4c, 0xe5, 0x30, 0x8b,
	0x4e, 0xe6, 0x5c, 0x2c, 0x22, 0xa9, 0x04, 0x53, 0x3c, 0x5b, 0x44, 0xa7, 0xd3, 0xb0, 0x12, 0xa0,
	0xc0, 0x73, 0x93, 0x32, 0x8d, 0x23, 0x88, 0x3f, 0xf3, 0x44, 0x8d, 0x83, 0x2c, 0x57, 0x9f, 0xe6,
	0x71, 0x98, 0x40, 0x49, 0x32, 0xc8, 0x80, 0x18, 0x4e, 0x3c, 0x3f, 0x36, 0x99, 0x49, 0x4c, 0x64,
	0x7b, 0xc7, 0x51, 0x06, 0x21, 0x67, 0x72, 0x01, 0x95, 0x0c, 0x0b, 0x48, 0x58, 0x41, 0x12, 0x98,
	0x29, 0xc1, 0x12, 0x25, 0x6d, 0xa7, 0xe0, 0x15, 0x04, 0x25, 0xa4, 0xbc, 0x90, 0xa4, 0x25, 0x12,
	0x93, 0x12, 0x3d, 0x95, 0x5c, 0x66, 0x6e, 0xfc, 0xb0, 0xe3, 0xa7, 0x3c, 0xcb, 0xd5, 0x17, 0x38,
	0x23, 0x19, 0x04, 0x06, 0x0c, 0x4e, 0x59, 0x91, 0xa7, 0x4c, 0x81, 0x90, 0x64, 0x19, 0xda, 0x3e,
	0xfc, 0xcb, 0x41, 0xf7, 0x5f, 0xe4, 0x52, 0xd1, 0x76, 0xc0, 0x4b, 0xad, 0xff, 0xaa, 0x95, 0x7f,
	0x33, 0xa5, 0xfc, 0x64, 0xce, 0x35, 0x6a, 0x1e, 0xef, 0x35, 0xda, 0xb0, 0xfb, 0x47, 0x79, 0x3a,
	0x72, 0xb6, 0x9d, 0xc9, 0xc6, 0xe1, 0x41, 0x53, 0xfb, 0x9b, 0xc7, 0x20, 0xca, 0x27, 0x78, 0x09,
	0xe1, 0x3f, 0xbf, 0x7d, 0x1f, 0x6d, 0x7d, 0x3c, 0x62, 0xc1, 0xb7, 0xa7, 0xc1, 0xbb, 0xe8, 0xc3,
	0xd1, 0x6e, 0xf0, 0xf8, 0x3c, 0xfe, 0xbe, 0xbb, 0xb3, 0xbf, 0xf7, 0xe3, 0x0e, 0xbd, 0x66, 0xe9,
	0xcf, 0x53, 0xef, 0x36, 0x1a, 0xaa, 0x45, 0xc5, 0x47, 0x6b, 0x46, 0xf0, 0x66, 0x53, 0xfb, 0xae,
	0x15, 0xd4, 0x55, 0x4c, 0x0d, 0x88, 0xff, 0x3a, 0xe8, 0x41, 0x1f, 0xa3, 0xb2, 0x82, 0x99, 0xe4,
	0x5a, 0x33, 0x81, 0x94, 0x1b, 0x93, 0xeb, 0x5d, 0x4d, 0x5d, 0xc5, 0xd4, 0x80, 0xde, 0x3d, 0xb4,
	0xce, 0x85, 0x00, 0xd1, 0x4e, 0xde, 0x6c, 0x6a, 0xff, 0xba, 0x65, 0x99, 0x32, 0xa6, 0x16, 0xf6,
	0x76, 0xd0, 0xd5, 0x92, 0x4b, 0xc9, 0x32, 0x3e, 0x1a, 0x18, 0xa6, 0xd7, 0xd4, 0xfe, 0x0d, 0xcb,
	0x6c, 0x01, 0x4c, 0xcf, 0x29, 0xde, 0x21, 0x1a, 0xa6, 0x4c, 0xb1, 0xd1, 0x70, 0x7b, 0x30, 0x71,
	0xa7, 0x5b, 0xa1, 0xfe, 0x01, 0xc3, 0x15, 0xb6, 0xbb, 0xce, 0x74, 0x13, 0xa6, 0xa6, 0x17, 0xff,
	0x5c, 0x43, 0x7b, 0xfd, 0xb7, 0x7d, 0x2b, 0x58, 0x55, 0x71, 0xd1, 0x6f, 0xe9, 0x47, 0xc8, 0xd5,
	0xef, 0xb3, 0xaf, 0x55, 0xc1, 0xf2, 0x59, 0xbb, 0xfa, 0xad, 0xa6, 0xf6, 0xbd, 0x0b, 0x6e, 0x0b,
	0x62, 0xda, 0xa5, 0x5e, 0x9c, 0x6b, 0xf0, 0xff, 0x73, 0xbd, 0x5f, 0x1e, 0xc0, 0x99, 0xb8, 0xd3,
	0x83, 0xb0, 0xf3, 0xdd, 0x84, 0xfd, 0x97, 0x5a, 0x71, 0x9a, 0xf8, 0x8a, 0xf9, 0xe3, 0xee, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x0c, 0x2c, 0x07, 0xb6, 0x03, 0x00, 0x00,
}
