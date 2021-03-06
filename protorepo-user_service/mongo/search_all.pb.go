// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_all.proto

package mongo

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//SearchAll请求
type SearchAllRequest struct {
	//
	//表名
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection" form:"collection"`
	//
	//查询条件
	Query *types.Struct `protobuf:"bytes,2,opt,name=query,proto3" json:"query" form:"query"`
	//
	//fields
	Fields []string `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields" form:"fields"`
	//
	//sort
	Sort                 []string `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort" form:"sort"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAllRequest) Reset()         { *m = SearchAllRequest{} }
func (m *SearchAllRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAllRequest) ProtoMessage()    {}
func (*SearchAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_389450bae772f5e4, []int{0}
}
func (m *SearchAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllRequest.Unmarshal(m, b)
}
func (m *SearchAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllRequest.Marshal(b, m, deterministic)
}
func (m *SearchAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllRequest.Merge(m, src)
}
func (m *SearchAllRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAllRequest.Size(m)
}
func (m *SearchAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllRequest proto.InternalMessageInfo

func (m *SearchAllRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *SearchAllRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *SearchAllRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchAllRequest) GetSort() []string {
	if m != nil {
		return m.Sort
	}
	return nil
}

//
//SearchAll返回
type SearchAllResponse struct {
	//
	//记录列表
	DocList              []*types.Struct `protobuf:"bytes,1,rep,name=docList,proto3" json:"docList" form:"docList"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SearchAllResponse) Reset()         { *m = SearchAllResponse{} }
func (m *SearchAllResponse) String() string { return proto.CompactTextString(m) }
func (*SearchAllResponse) ProtoMessage()    {}
func (*SearchAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_389450bae772f5e4, []int{1}
}
func (m *SearchAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllResponse.Unmarshal(m, b)
}
func (m *SearchAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllResponse.Marshal(b, m, deterministic)
}
func (m *SearchAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllResponse.Merge(m, src)
}
func (m *SearchAllResponse) XXX_Size() int {
	return xxx_messageInfo_SearchAllResponse.Size(m)
}
func (m *SearchAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllResponse proto.InternalMessageInfo

func (m *SearchAllResponse) GetDocList() []*types.Struct {
	if m != nil {
		return m.DocList
	}
	return nil
}

//
//SearchAllApi返回
type SearchAllResponseWrapper struct {
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
	Data                 *SearchAllResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchAllResponseWrapper) Reset()         { *m = SearchAllResponseWrapper{} }
func (m *SearchAllResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SearchAllResponseWrapper) ProtoMessage()    {}
func (*SearchAllResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_389450bae772f5e4, []int{2}
}
func (m *SearchAllResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAllResponseWrapper.Unmarshal(m, b)
}
func (m *SearchAllResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAllResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SearchAllResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAllResponseWrapper.Merge(m, src)
}
func (m *SearchAllResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SearchAllResponseWrapper.Size(m)
}
func (m *SearchAllResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAllResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAllResponseWrapper proto.InternalMessageInfo

func (m *SearchAllResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchAllResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SearchAllResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SearchAllResponseWrapper) GetData() *SearchAllResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchAllRequest)(nil), "mongo.SearchAllRequest")
	proto.RegisterType((*SearchAllResponse)(nil), "mongo.SearchAllResponse")
	proto.RegisterType((*SearchAllResponseWrapper)(nil), "mongo.SearchAllResponseWrapper")
}

func init() { proto.RegisterFile("search_all.proto", fileDescriptor_389450bae772f5e4) }

var fileDescriptor_389450bae772f5e4 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x89, 0x49, 0x2a, 0x9d, 0xfa, 0xa7, 0x1d, 0x50, 0x87, 0x22, 0x24, 0x8c, 0x20, 0x75,
	0x61, 0x0a, 0x8a, 0x20, 0x82, 0x88, 0x81, 0xee, 0x5c, 0x4d, 0x17, 0x82, 0x1b, 0x49, 0x93, 0x69,
	0x1a, 0x98, 0xe6, 0xa4, 0x33, 0x13, 0xd0, 0x57, 0xbd, 0x8b, 0x70, 0x9f, 0x21, 0x4f, 0x70, 0xc9,
	0x99, 0x96, 0x06, 0x0a, 0x77, 0x95, 0xcc, 0xf9, 0x7e, 0xdf, 0x39, 0x5f, 0x4e, 0x86, 0xcc, 0x8d,
	0xcc, 0x74, 0x7e, 0xf8, 0x9b, 0x29, 0x95, 0x34, 0x1a, 0x2c, 0xd0, 0xf0, 0x08, 0x75, 0x09, 0xcb,
	0x8f, 0x65, 0x65, 0x0f, 0xed, 0x2e, 0xc9, 0xe1, 0xb8, 0x2e, 0xa1, 0x84, 0x35, 0xaa, 0xbb, 0x76,
	0x8f, 0x27, 0x3c, 0xe0, 0x9b, 0x73, 0x2d, 0xdf, 0x96, 0x00, 0xa5, 0x92, 0x57, 0xca, 0x58, 0xdd,
	0xe6, 0xd6, 0xa9, 0xfc, 0xce, 0x23, 0xf3, 0x2d, 0x0e, 0xfa, 0xa9, 0x94, 0x90, 0xa7, 0x56, 0x1a,
	0x4b, 0xbf, 0x10, 0x92, 0x83, 0x52, 0x32, 0xb7, 0x15, 0xd4, 0xcc, 0x8b, 0xbd, 0xd5, 0x34, 0x7d,
	0xd5, 0x77, 0xd1, 0x62, 0x0f, 0xfa, 0xf8, 0x8d, 0x5f, 0x35, 0x2e, 0x46, 0x20, 0xfd, 0x41, 0xc2,
	0x53, 0x2b, 0xf5, 0x7f, 0xf6, 0x24, 0xf6, 0x56, 0xb3, 0x4f, 0x6f, 0x12, 0x37, 0x39, 0xb9, 0x4c,
	0x4e, 0xb6, 0x38, 0x39, 0x9d, 0xf7, 0x5d, 0xf4, 0xcc, 0xb5, 0x42, 0x9e, 0x0b, 0xe7, 0xa3, 0x1f,
	0xc8, 0x64, 0x5f, 0x49, 0x55, 0x18, 0xe6, 0xc7, 0xfe, 0x6a, 0x9a, 0x2e, 0xfa, 0x2e, 0x7a, 0xee,
	0x40, 0x57, 0xe7, 0xe2, 0x0c, 0xd0, 0x77, 0x24, 0x30, 0xa0, 0x2d, 0x0b, 0x10, 0x7c, 0xd9, 0x77,
	0xd1, 0xcc, 0x81, 0x43, 0x95, 0x0b, 0x14, 0xf9, 0x1f, 0xb2, 0x18, 0x7d, 0x9b, 0x69, 0xa0, 0x36,
	0x92, 0x6e, 0xc8, 0xd3, 0x02, 0xf2, 0x5f, 0x95, 0xb1, 0xcc, 0x8b, 0xfd, 0xc7, 0x72, 0xd2, 0xbe,
	0x8b, 0x5e, 0xb8, 0xae, 0x67, 0x07, 0x17, 0x17, 0x2f, 0xbf, 0xf7, 0x08, 0xbb, 0x69, 0xfe, 0x5b,
	0x67, 0x4d, 0x23, 0xf5, 0x90, 0x2e, 0x87, 0x42, 0xe2, 0xea, 0xc2, 0x71, 0xba, 0xa1, 0xca, 0x05,
	0x8a, 0xf4, 0x2b, 0x99, 0x0d, 0xcf, 0xcd, 0xbf, 0x46, 0x65, 0x55, 0x8d, 0x4b, 0x9b, 0xa6, 0xaf,
	0xfb, 0x2e, 0xa2, 0x57, 0xf6, 0x2c, 0x72, 0x31, 0x46, 0xe9, 0x7b, 0x12, 0x4a, 0xad, 0x41, 0x33,
	0x1f, 0x3d, 0xa3, 0x7d, 0x62, 0x99, 0x0b, 0x27, 0xd3, 0xef, 0x24, 0x28, 0x32, 0x9b, 0xb1, 0x00,
	0xff, 0x07, 0x4b, 0xf0, 0xfe, 0x24, 0x37, 0xa9, 0xc7, 0x01, 0x07, 0x9e, 0x0b, 0xb4, 0xed, 0x26,
	0xb8, 0x90, 0xcf, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x02, 0xc1, 0x4c, 0x8a, 0x02, 0x00,
	0x00,
}
