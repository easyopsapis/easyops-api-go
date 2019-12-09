// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregate_total.proto

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
//AggregateCount请求
type AggregateCountRequest struct {
	//
	//URL中资源模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//URL中模型属性ID, 支持结构体，用.分隔结构体名称和属性名称，如cpu.brand
	AttrId string `protobuf:"bytes,2,opt,name=attrId,proto3" json:"attrId" form:"attrId"`
	//
	//页码
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregateCountRequest) Reset()         { *m = AggregateCountRequest{} }
func (m *AggregateCountRequest) String() string { return proto.CompactTextString(m) }
func (*AggregateCountRequest) ProtoMessage()    {}
func (*AggregateCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff94236a6a091f55, []int{0}
}
func (m *AggregateCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountRequest.Unmarshal(m, b)
}
func (m *AggregateCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountRequest.Marshal(b, m, deterministic)
}
func (m *AggregateCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountRequest.Merge(m, src)
}
func (m *AggregateCountRequest) XXX_Size() int {
	return xxx_messageInfo_AggregateCountRequest.Size(m)
}
func (m *AggregateCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountRequest proto.InternalMessageInfo

func (m *AggregateCountRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AggregateCountRequest) GetAttrId() string {
	if m != nil {
		return m.AttrId
	}
	return ""
}

func (m *AggregateCountRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AggregateCountRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//AggregateCount返回
type AggregateCountResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*AggregateCountResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AggregateCountResponse) Reset()         { *m = AggregateCountResponse{} }
func (m *AggregateCountResponse) String() string { return proto.CompactTextString(m) }
func (*AggregateCountResponse) ProtoMessage()    {}
func (*AggregateCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff94236a6a091f55, []int{1}
}
func (m *AggregateCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountResponse.Unmarshal(m, b)
}
func (m *AggregateCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountResponse.Marshal(b, m, deterministic)
}
func (m *AggregateCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountResponse.Merge(m, src)
}
func (m *AggregateCountResponse) XXX_Size() int {
	return xxx_messageInfo_AggregateCountResponse.Size(m)
}
func (m *AggregateCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountResponse proto.InternalMessageInfo

func (m *AggregateCountResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AggregateCountResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *AggregateCountResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *AggregateCountResponse) GetList() []*AggregateCountResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type AggregateCountResponse_List struct {
	//
	//个数
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count" form:"count"`
	//
	//属性值
	Attr                 *types.Struct `protobuf:"bytes,2,opt,name=attr,proto3" json:"attr" form:"attr"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AggregateCountResponse_List) Reset()         { *m = AggregateCountResponse_List{} }
func (m *AggregateCountResponse_List) String() string { return proto.CompactTextString(m) }
func (*AggregateCountResponse_List) ProtoMessage()    {}
func (*AggregateCountResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff94236a6a091f55, []int{1, 0}
}
func (m *AggregateCountResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountResponse_List.Unmarshal(m, b)
}
func (m *AggregateCountResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountResponse_List.Marshal(b, m, deterministic)
}
func (m *AggregateCountResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountResponse_List.Merge(m, src)
}
func (m *AggregateCountResponse_List) XXX_Size() int {
	return xxx_messageInfo_AggregateCountResponse_List.Size(m)
}
func (m *AggregateCountResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountResponse_List proto.InternalMessageInfo

func (m *AggregateCountResponse_List) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AggregateCountResponse_List) GetAttr() *types.Struct {
	if m != nil {
		return m.Attr
	}
	return nil
}

//
//AggregateCountApi返回
type AggregateCountResponseWrapper struct {
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
	Data                 *AggregateCountResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AggregateCountResponseWrapper) Reset()         { *m = AggregateCountResponseWrapper{} }
func (m *AggregateCountResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AggregateCountResponseWrapper) ProtoMessage()    {}
func (*AggregateCountResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff94236a6a091f55, []int{2}
}
func (m *AggregateCountResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountResponseWrapper.Unmarshal(m, b)
}
func (m *AggregateCountResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AggregateCountResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountResponseWrapper.Merge(m, src)
}
func (m *AggregateCountResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AggregateCountResponseWrapper.Size(m)
}
func (m *AggregateCountResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountResponseWrapper proto.InternalMessageInfo

func (m *AggregateCountResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AggregateCountResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AggregateCountResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AggregateCountResponseWrapper) GetData() *AggregateCountResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AggregateCountRequest)(nil), "instance.AggregateCountRequest")
	proto.RegisterType((*AggregateCountResponse)(nil), "instance.AggregateCountResponse")
	proto.RegisterType((*AggregateCountResponse_List)(nil), "instance.AggregateCountResponse.List")
	proto.RegisterType((*AggregateCountResponseWrapper)(nil), "instance.AggregateCountResponseWrapper")
}

func init() { proto.RegisterFile("aggregate_total.proto", fileDescriptor_ff94236a6a091f55) }

var fileDescriptor_ff94236a6a091f55 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x6f, 0xd3, 0x30,
	0x1c, 0xa7, 0x5d, 0x36, 0xb5, 0x2e, 0x68, 0x95, 0xa5, 0x8e, 0xaa, 0x62, 0x4a, 0x65, 0x1e, 0x2a,
	0x12, 0x49, 0x47, 0x87, 0xc6, 0x43, 0x5c, 0x56, 0xb4, 0xc3, 0xd0, 0x4e, 0xde, 0x01, 0x89, 0x09,
	0x2a, 0x37, 0xf5, 0x4c, 0x20, 0x8d, 0x83, 0xe3, 0x30, 0x34, 0xc4, 0x47, 0xe1, 0xab, 0x05, 0x89,
	0x3b, 0x97, 0x7c, 0x02, 0xe4, 0xbf, 0x93, 0x36, 0x42, 0x93, 0x18, 0xa7, 0x3a, 0xbf, 0xd7, 0xff,
	0xd1, 0x3f, 0xea, 0x31, 0x21, 0x14, 0x17, 0x4c, 0xf3, 0x99, 0x96, 0x9a, 0x45, 0x7e, 0xa2, 0xa4,
	0x96, 0xb8, 0x15, 0xc6, 0xa9, 0x66, 0x71, 0xc0, 0x07, 0x9e, 0x08, 0xf5, 0x87, 0x6c, 0xee, 0x07,
	0x72, 0x39, 0x16, 0x52, 0xc8, 0x31, 0x08, 0xe6, 0xd9, 0x39, 0x7c, 0xc1, 0x07, 0xbc, 0xac, 0x71,
	0x70, 0x50, 0x93, 0x2f, 0x2f, 0x42, 0xfd, 0x49, 0x5e, 0x8c, 0x85, 0xf4, 0x80, 0xf4, 0xbe, 0xb0,
	0x28, 0x5c, 0x30, 0x2d, 0x55, 0x3a, 0x5e, 0x3d, 0x4b, 0xdf, 0x1d, 0x21, 0xa5, 0x88, 0xf8, 0x3a,
	0x3d, 0xd5, 0x2a, 0x0b, 0xb4, 0x65, 0xc9, 0x8f, 0x26, 0xea, 0x1d, 0x56, 0x8d, 0xbe, 0x92, 0x59,
	0xac, 0x29, 0xff, 0x9c, 0xf1, 0x54, 0x63, 0x8a, 0x5a, 0x72, 0xfe, 0x91, 0x07, 0xfa, 0x78, 0xd1,
	0x6f, 0x0c, 0x1b, 0xa3, 0xf6, 0xf4, 0xa0, 0xc8, 0xdd, 0xed, 0x73, 0xa9, 0x96, 0x2f, 0x48, 0xc5,
	0x90, 0x5f, 0x3f, 0x5d, 0x17, 0xed, 0xbe, 0x3f, 0x63, 0xde, 0xe5, 0xa1, 0xf7, 0x76, 0xf6, 0xee,
	0x6c, 0xcf, 0x7b, 0x5e, 0xbd, 0xbf, 0xed, 0x3d, 0xda, 0x7f, 0xfc, 0xfd, 0x1e, 0x5d, 0xe5, 0xe0,
	0x13, 0xb4, 0xc5, 0xb4, 0x56, 0xc7, 0x8b, 0x7e, 0x13, 0x12, 0x9f, 0x14, 0xb9, 0x7b, 0xcb, 0x26,
	0x5a, 0xfc, 0x5a, 0x79, 0x65, 0x06, 0x7e, 0x88, 0x9c, 0x84, 0x09, 0xde, 0xdf, 0x18, 0x36, 0x46,
	0x9b, 0xd3, 0x5e, 0x91, 0xbb, 0x1d, 0x9b, 0x65, 0x50, 0x93, 0xd4, 0xec, 0xde, 0xa0, 0x20, 0xc1,
	0x4f, 0x51, 0xdb, 0xfc, 0xce, 0xd2, 0xf0, 0x92, 0xf7, 0x1d, 0xd0, 0x0f, 0x8a, 0xdc, 0xed, 0xae,
	0xf5, 0x40, 0x55, 0xa6, 0x96, 0x41, 0x4e, 0x0d, 0x90, 0x37, 0xd1, 0xce, 0xdf, 0xfb, 0x49, 0x13,
	0x19, 0xa7, 0x7c, 0x55, 0xbe, 0xf1, 0x9f, 0xe5, 0x9b, 0xd7, 0x2f, 0x8f, 0x1f, 0xa0, 0x4d, 0x38,
	0x9e, 0x72, 0xc6, 0x6e, 0x91, 0xbb, 0x37, 0xad, 0x09, 0x60, 0x42, 0x2d, 0x8d, 0x5f, 0x23, 0x27,
	0x0a, 0x53, 0xdd, 0x77, 0x86, 0x1b, 0xa3, 0xce, 0xe4, 0xbe, 0x5f, 0x1d, 0x99, 0x7f, 0x75, 0xef,
	0xfe, 0x49, 0x98, 0xea, 0xe9, 0xf6, 0xba, 0x65, 0x63, 0x26, 0x14, 0x32, 0x06, 0x11, 0x72, 0x0c,
	0x6d, 0x6a, 0x07, 0xc6, 0x54, 0x0e, 0x58, 0xab, 0x0d, 0x30, 0xa1, 0x96, 0xc6, 0x2f, 0x91, 0x63,
	0xfe, 0x10, 0x98, 0xab, 0x33, 0xb9, 0xed, 0xdb, 0x7b, 0xf3, 0xab, 0x7b, 0xf3, 0x4f, 0xe1, 0xde,
	0xea, 0xd5, 0x8c, 0x9c, 0x50, 0x70, 0x91, 0xdf, 0x0d, 0xb4, 0x7b, 0x75, 0x93, 0x6f, 0x14, 0x4b,
	0x12, 0xae, 0xf0, 0x5d, 0xe4, 0x04, 0x72, 0x51, 0xed, 0xb9, 0x16, 0x63, 0x50, 0x42, 0x81, 0xc4,
	0xcf, 0x50, 0xc7, 0xfc, 0x1e, 0x7d, 0x4d, 0x22, 0x16, 0xc6, 0xe5, 0x79, 0xed, 0x14, 0xb9, 0x8b,
	0xd7, 0xda, 0x92, 0x24, 0xb4, 0x2e, 0x35, 0x63, 0x72, 0xa5, 0xa4, 0x82, 0x15, 0xb7, 0xeb, 0x63,
	0x02, 0x4c, 0xa8, 0xa5, 0xf1, 0x11, 0x72, 0x16, 0x4c, 0x33, 0xb8, 0x9e, 0xce, 0x64, 0xf8, 0xaf,
	0x15, 0xd7, 0x1b, 0x35, 0x3e, 0x42, 0xc1, 0x3e, 0xdf, 0x82, 0xbd, 0xec, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x3c, 0x6a, 0x52, 0x1f, 0x04, 0x00, 0x00,
}
