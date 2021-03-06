// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregate_total_v2.proto

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
//AggregateCountV2请求
type AggregateCountV2Request struct {
	//
	//URL中资源模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//URL中模型属性ID, 支持结构体，用.分隔结构体名称和属性名称，如cpu.brand
	AttrId string `protobuf:"bytes,2,opt,name=attrId,proto3" json:"attrId" form:"attrId"`
	//
	//top n, 默认300, 为0则获取所有, top n+1以后的属性值将会聚合追加到最后,属性值为'-'
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top" form:"top"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregateCountV2Request) Reset()         { *m = AggregateCountV2Request{} }
func (m *AggregateCountV2Request) String() string { return proto.CompactTextString(m) }
func (*AggregateCountV2Request) ProtoMessage()    {}
func (*AggregateCountV2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b9ab8cbd02166ed, []int{0}
}
func (m *AggregateCountV2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountV2Request.Unmarshal(m, b)
}
func (m *AggregateCountV2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountV2Request.Marshal(b, m, deterministic)
}
func (m *AggregateCountV2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountV2Request.Merge(m, src)
}
func (m *AggregateCountV2Request) XXX_Size() int {
	return xxx_messageInfo_AggregateCountV2Request.Size(m)
}
func (m *AggregateCountV2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountV2Request.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountV2Request proto.InternalMessageInfo

func (m *AggregateCountV2Request) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AggregateCountV2Request) GetAttrId() string {
	if m != nil {
		return m.AttrId
	}
	return ""
}

func (m *AggregateCountV2Request) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

//
//AggregateCountV2返回
type AggregateCountV2Response struct {
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

func (m *AggregateCountV2Response) Reset()         { *m = AggregateCountV2Response{} }
func (m *AggregateCountV2Response) String() string { return proto.CompactTextString(m) }
func (*AggregateCountV2Response) ProtoMessage()    {}
func (*AggregateCountV2Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b9ab8cbd02166ed, []int{1}
}
func (m *AggregateCountV2Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountV2Response.Unmarshal(m, b)
}
func (m *AggregateCountV2Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountV2Response.Marshal(b, m, deterministic)
}
func (m *AggregateCountV2Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountV2Response.Merge(m, src)
}
func (m *AggregateCountV2Response) XXX_Size() int {
	return xxx_messageInfo_AggregateCountV2Response.Size(m)
}
func (m *AggregateCountV2Response) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountV2Response.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountV2Response proto.InternalMessageInfo

func (m *AggregateCountV2Response) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AggregateCountV2Response) GetAttr() *types.Struct {
	if m != nil {
		return m.Attr
	}
	return nil
}

//
//AggregateCountV2Api返回
type AggregateCountV2ResponseWrapper struct {
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
	Data                 []*AggregateCountV2Response `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AggregateCountV2ResponseWrapper) Reset()         { *m = AggregateCountV2ResponseWrapper{} }
func (m *AggregateCountV2ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AggregateCountV2ResponseWrapper) ProtoMessage()    {}
func (*AggregateCountV2ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b9ab8cbd02166ed, []int{2}
}
func (m *AggregateCountV2ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateCountV2ResponseWrapper.Unmarshal(m, b)
}
func (m *AggregateCountV2ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateCountV2ResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AggregateCountV2ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateCountV2ResponseWrapper.Merge(m, src)
}
func (m *AggregateCountV2ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AggregateCountV2ResponseWrapper.Size(m)
}
func (m *AggregateCountV2ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateCountV2ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateCountV2ResponseWrapper proto.InternalMessageInfo

func (m *AggregateCountV2ResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AggregateCountV2ResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AggregateCountV2ResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AggregateCountV2ResponseWrapper) GetData() []*AggregateCountV2Response {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AggregateCountV2Request)(nil), "instance.AggregateCountV2Request")
	proto.RegisterType((*AggregateCountV2Response)(nil), "instance.AggregateCountV2Response")
	proto.RegisterType((*AggregateCountV2ResponseWrapper)(nil), "instance.AggregateCountV2ResponseWrapper")
}

func init() { proto.RegisterFile("aggregate_total_v2.proto", fileDescriptor_7b9ab8cbd02166ed) }

var fileDescriptor_7b9ab8cbd02166ed = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x89, 0x6d, 0x8f, 0xbb, 0xad, 0x7a, 0xb2, 0x0f, 0x5e, 0x38, 0x94, 0x2d, 0xab, 0x48,
	0x1f, 0x4c, 0x7a, 0xf6, 0xe4, 0x50, 0xf1, 0xe5, 0x2a, 0x22, 0x07, 0x3e, 0xad, 0xa0, 0xe0, 0xa1,
	0x65, 0x9b, 0xec, 0xad, 0xd1, 0x36, 0x13, 0x37, 0x93, 0x3b, 0x51, 0x04, 0xbf, 0xa8, 0x11, 0xfc,
	0x06, 0xe6, 0x13, 0x48, 0x76, 0x93, 0x6b, 0x40, 0x0a, 0x3e, 0x75, 0x76, 0xfe, 0x33, 0xbf, 0xfe,
	0xe7, 0x4f, 0x88, 0x2f, 0xb5, 0x36, 0x4a, 0x4b, 0x54, 0x73, 0x04, 0x94, 0xcb, 0xf9, 0xf9, 0x34,
	0xcc, 0x0c, 0x20, 0xd0, 0xed, 0x24, 0xcd, 0x51, 0xa6, 0x91, 0xda, 0x0f, 0x74, 0x82, 0x1f, 0x8a,
	0x45, 0x18, 0xc1, 0x6a, 0xa2, 0x41, 0xc3, 0xc4, 0x0e, 0x2c, 0x8a, 0x33, 0xfb, 0xb2, 0x0f, 0x5b,
	0xb9, 0xc5, 0xfd, 0xa3, 0xce, 0xf8, 0xea, 0x22, 0xc1, 0x4f, 0x70, 0x31, 0xd1, 0x10, 0x58, 0x31,
	0x38, 0x97, 0xcb, 0x24, 0x96, 0x08, 0x26, 0x9f, 0x5c, 0x96, 0xcd, 0xde, 0x2d, 0x0d, 0xa0, 0x97,
	0x6a, 0x4d, 0xcf, 0xd1, 0x14, 0x11, 0x3a, 0x95, 0xff, 0xf4, 0xc8, 0xde, 0x71, 0xeb, 0xf5, 0x19,
	0x14, 0x29, 0xbe, 0x9e, 0x0a, 0xf5, 0xb9, 0x50, 0x39, 0x52, 0x41, 0xb6, 0x61, 0xf1, 0x51, 0x45,
	0x78, 0x12, 0xfb, 0xde, 0xc8, 0x1b, 0xef, 0xcc, 0x8e, 0xaa, 0x92, 0xed, 0x9e, 0x81, 0x59, 0x3d,
	0xe1, 0xad, 0xc2, 0x7f, 0xff, 0x62, 0x8c, 0xdc, 0x7e, 0x7f, 0x2a, 0x83, 0xaf, 0xc7, 0xc1, 0xdb,
	0xf9, 0xbb, 0xd3, 0x83, 0xe0, 0x71, 0x5b, 0x7f, 0x3b, 0xb8, 0x7f, 0xf8, 0xe0, 0xfb, 0x5d, 0x71,
	0xc9, 0xa1, 0x2f, 0xc9, 0x96, 0x44, 0x34, 0x27, 0xb1, 0x7f, 0xc5, 0x12, 0x1f, 0x56, 0x25, 0xbb,
	0xe6, 0x88, 0xae, 0xff, 0x5f, 0xbc, 0x86, 0x41, 0x47, 0xa4, 0x87, 0x90, 0xf9, 0xbd, 0x91, 0x37,
	0x1e, 0xcc, 0xae, 0x57, 0x25, 0x23, 0x0e, 0x85, 0x90, 0x71, 0x51, 0x4b, 0xfc, 0x87, 0x47, 0xfc,
	0x7f, 0xef, 0xcb, 0x33, 0x48, 0x73, 0x45, 0xef, 0x91, 0x41, 0x54, 0xb7, 0xec, 0x75, 0x83, 0xd9,
	0x8d, 0xaa, 0x64, 0x57, 0x1d, 0xc0, 0xb6, 0xb9, 0x70, 0x32, 0x7d, 0x4a, 0xfa, 0xf5, 0x1f, 0x5a,
	0xcb, 0xc3, 0xe9, 0x5e, 0xe8, 0x12, 0x0d, 0xdb, 0x44, 0xc3, 0x57, 0x36, 0xd1, 0xd9, 0x6e, 0x55,
	0xb2, 0xe1, 0xfa, 0x16, 0x2e, 0xec, 0x16, 0xff, 0xe3, 0x11, 0xb6, 0xc9, 0xc2, 0x1b, 0x23, 0xb3,
	0x4c, 0x19, 0x7a, 0x87, 0xf4, 0x23, 0x88, 0x55, 0x63, 0xa4, 0x03, 0xaa, 0xbb, 0x5c, 0x58, 0x91,
	0x3e, 0x22, 0xc3, 0xfa, 0xf7, 0xf9, 0x97, 0x6c, 0x29, 0x93, 0xb4, 0x09, 0xf0, 0x66, 0x55, 0x32,
	0xba, 0x9e, 0x6d, 0x44, 0x2e, 0xba, 0xa3, 0xf5, 0xa1, 0xca, 0x18, 0x30, 0x36, 0xa9, 0x9d, 0xee,
	0xa1, 0xb6, 0xcd, 0x85, 0x93, 0xe9, 0x0b, 0xd2, 0x8f, 0x25, 0x4a, 0xbf, 0x3f, 0xea, 0x8d, 0x87,
	0x53, 0x1e, 0xb6, 0xdf, 0x6a, 0xb8, 0xc9, 0x7f, 0xd7, 0x6a, 0xbd, 0xc9, 0x85, 0x05, 0x2c, 0xb6,
	0x6c, 0x36, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x73, 0x94, 0xf6, 0x08, 0x03, 0x00,
	0x00,
}
