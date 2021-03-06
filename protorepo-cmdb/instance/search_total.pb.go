// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_total.proto

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
//SearchTotal请求
type SearchTotalRequest struct {
	//
	//模型对象ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//查询条件, 与[实例搜索接口]的query字段说明一致
	Query *types.Struct `protobuf:"bytes,2,opt,name=query,proto3" json:"query" form:"query"`
	//
	//按照权限过滤[权限列表]
	Permission []string `protobuf:"bytes,3,rep,name=permission,proto3" json:"permission" form:"permission"`
	//
	//是否只获取与自己有关的那部分数据, 默认为false
	OnlyMyInstance       bool     `protobuf:"varint,4,opt,name=only_my_instance,json=onlyMyInstance,proto3" json:"only_my_instance" form:"only_my_instance"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTotalRequest) Reset()         { *m = SearchTotalRequest{} }
func (m *SearchTotalRequest) String() string { return proto.CompactTextString(m) }
func (*SearchTotalRequest) ProtoMessage()    {}
func (*SearchTotalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd5e64f17424e57, []int{0}
}
func (m *SearchTotalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTotalRequest.Unmarshal(m, b)
}
func (m *SearchTotalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTotalRequest.Marshal(b, m, deterministic)
}
func (m *SearchTotalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTotalRequest.Merge(m, src)
}
func (m *SearchTotalRequest) XXX_Size() int {
	return xxx_messageInfo_SearchTotalRequest.Size(m)
}
func (m *SearchTotalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTotalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTotalRequest proto.InternalMessageInfo

func (m *SearchTotalRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *SearchTotalRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *SearchTotalRequest) GetPermission() []string {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *SearchTotalRequest) GetOnlyMyInstance() bool {
	if m != nil {
		return m.OnlyMyInstance
	}
	return false
}

//
//SearchTotal返回
type SearchTotalResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//实例总数
	Data                 int32    `protobuf:"varint,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTotalResponse) Reset()         { *m = SearchTotalResponse{} }
func (m *SearchTotalResponse) String() string { return proto.CompactTextString(m) }
func (*SearchTotalResponse) ProtoMessage()    {}
func (*SearchTotalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd5e64f17424e57, []int{1}
}
func (m *SearchTotalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTotalResponse.Unmarshal(m, b)
}
func (m *SearchTotalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTotalResponse.Marshal(b, m, deterministic)
}
func (m *SearchTotalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTotalResponse.Merge(m, src)
}
func (m *SearchTotalResponse) XXX_Size() int {
	return xxx_messageInfo_SearchTotalResponse.Size(m)
}
func (m *SearchTotalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTotalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTotalResponse proto.InternalMessageInfo

func (m *SearchTotalResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchTotalResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SearchTotalResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SearchTotalResponse) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

//
//SearchTotalApi返回
type SearchTotalResponseWrapper struct {
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
	Data                 *SearchTotalResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchTotalResponseWrapper) Reset()         { *m = SearchTotalResponseWrapper{} }
func (m *SearchTotalResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SearchTotalResponseWrapper) ProtoMessage()    {}
func (*SearchTotalResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd5e64f17424e57, []int{2}
}
func (m *SearchTotalResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTotalResponseWrapper.Unmarshal(m, b)
}
func (m *SearchTotalResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTotalResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SearchTotalResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTotalResponseWrapper.Merge(m, src)
}
func (m *SearchTotalResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SearchTotalResponseWrapper.Size(m)
}
func (m *SearchTotalResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTotalResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTotalResponseWrapper proto.InternalMessageInfo

func (m *SearchTotalResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchTotalResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SearchTotalResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SearchTotalResponseWrapper) GetData() *SearchTotalResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchTotalRequest)(nil), "instance.SearchTotalRequest")
	proto.RegisterType((*SearchTotalResponse)(nil), "instance.SearchTotalResponse")
	proto.RegisterType((*SearchTotalResponseWrapper)(nil), "instance.SearchTotalResponseWrapper")
}

func init() { proto.RegisterFile("search_total.proto", fileDescriptor_6bd5e64f17424e57) }

var fileDescriptor_6bd5e64f17424e57 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x75, 0x85, 0xd6, 0x45, 0xdb, 0x30, 0x82, 0x45, 0x85, 0x29, 0x95, 0x41, 0xa8,
	0x87, 0x25, 0x1d, 0x9b, 0x98, 0x80, 0x0b, 0x22, 0xd2, 0x0e, 0x3b, 0x70, 0xf1, 0x90, 0x90, 0x98,
	0xa0, 0x72, 0x53, 0x2f, 0x0b, 0x24, 0x79, 0x99, 0xed, 0x30, 0x0a, 0xe2, 0xf3, 0xf0, 0x09, 0xf8,
	0x3a, 0x41, 0xe2, 0xc2, 0x3d, 0x9f, 0x00, 0xc5, 0x4e, 0xd6, 0x0c, 0x0d, 0xc4, 0xa9, 0xf6, 0xfb,
	0xff, 0xfc, 0xde, 0xff, 0xff, 0x1a, 0x84, 0x25, 0x67, 0x22, 0x38, 0x9d, 0x2a, 0x50, 0x2c, 0xf6,
	0x32, 0x01, 0x0a, 0x70, 0x2f, 0x4a, 0xa5, 0x62, 0x69, 0xc0, 0x87, 0x6e, 0x18, 0xa9, 0xd3, 0x7c,
	0xe6, 0x05, 0x90, 0x4c, 0x42, 0x08, 0x61, 0xa2, 0x81, 0x59, 0x7e, 0xa2, 0x6f, 0xfa, 0xa2, 0x4f,
	0xe6, 0xe1, 0x70, 0xbf, 0x85, 0x27, 0xe7, 0x91, 0xfa, 0x00, 0xe7, 0x93, 0x10, 0x5c, 0x2d, 0xba,
	0x1f, 0x59, 0x1c, 0xcd, 0x99, 0x02, 0x21, 0x27, 0x17, 0xc7, 0xfa, 0xdd, 0xbd, 0x10, 0x20, 0x8c,
	0xf9, 0xb2, 0xbb, 0x54, 0x22, 0x0f, 0x94, 0x51, 0xc9, 0xb7, 0x15, 0x84, 0x8f, 0xb4, 0xcb, 0x57,
	0x95, 0x49, 0xca, 0xcf, 0x72, 0x2e, 0x15, 0xa6, 0xa8, 0x07, 0xb3, 0xf7, 0x3c, 0x50, 0x87, 0x73,
	0xdb, 0x1a, 0x59, 0xe3, 0xbe, 0xbf, 0x5f, 0x16, 0xce, 0xfa, 0x09, 0x88, 0xe4, 0x19, 0x69, 0x14,
	0xf2, 0xf3, 0x87, 0xe3, 0xa0, 0xad, 0x77, 0xc7, 0xcc, 0xfd, 0xfc, 0xc2, 0x7d, 0x33, 0x7d, 0x7b,
	0xbc, 0xe3, 0x3e, 0x6d, 0xce, 0x5f, 0x76, 0xb6, 0xf7, 0x1e, 0x7d, 0x7d, 0x40, 0x2f, 0xfa, 0xe0,
	0xe7, 0xa8, 0x7b, 0x96, 0x73, 0xb1, 0xb0, 0x57, 0x46, 0xd6, 0x78, 0xb0, 0xbb, 0xe9, 0x19, 0x63,
	0x5e, 0x63, 0xcc, 0x3b, 0xd2, 0xc6, 0xfc, 0x8d, 0xb2, 0x70, 0x6e, 0x98, 0x49, 0x9a, 0x27, 0xd4,
	0xbc, 0xc3, 0x8f, 0x11, 0xca, 0xb8, 0x48, 0x22, 0x29, 0x23, 0x48, 0xed, 0xce, 0xa8, 0x33, 0xee,
	0xfb, 0xb7, 0xcb, 0xc2, 0xb9, 0x69, 0xe0, 0xa5, 0x46, 0x68, 0x0b, 0xc4, 0x07, 0x68, 0x03, 0xd2,
	0x78, 0x31, 0x4d, 0x16, 0xd3, 0x66, 0xf7, 0xf6, 0xea, 0xc8, 0x1a, 0xf7, 0xfc, 0xbb, 0x65, 0xe1,
	0x6c, 0xd6, 0x99, 0xfe, 0x20, 0x08, 0x5d, 0xab, 0x4a, 0x2f, 0x17, 0x87, 0x4d, 0xe1, 0xbb, 0x85,
	0x6e, 0x5d, 0xda, 0x94, 0xcc, 0x20, 0x95, 0x1c, 0xdf, 0x47, 0xab, 0x01, 0xcc, 0xb9, 0x5e, 0x53,
	0xd7, 0x5f, 0x2f, 0x0b, 0x67, 0x60, 0x5a, 0x56, 0x55, 0x42, 0xb5, 0x88, 0x1f, 0xa2, 0x2e, 0x17,
	0x02, 0x84, 0xce, 0xde, 0x6f, 0x47, 0xd4, 0x65, 0x42, 0x8d, 0x8c, 0xb7, 0xd1, 0xf5, 0x84, 0x4b,
	0xc9, 0x42, 0x6e, 0x77, 0x34, 0x89, 0xcb, 0xc2, 0x59, 0x33, 0x64, 0x2d, 0x10, 0xda, 0x20, 0xd5,
	0xe8, 0x39, 0x53, 0x4c, 0xa7, 0xb9, 0x34, 0xba, 0xaa, 0x12, 0xaa, 0x45, 0xf2, 0xcb, 0x42, 0xc3,
	0x2b, 0x7c, 0xbf, 0x16, 0x2c, 0xcb, 0xb8, 0xf8, 0x3f, 0xfb, 0x4f, 0xd0, 0xa0, 0xfa, 0x3d, 0xf8,
	0x94, 0xc5, 0x2c, 0x4a, 0xeb, 0x10, 0x77, 0xca, 0xc2, 0xc1, 0x4b, 0xb6, 0x16, 0x09, 0x6d, 0xa3,
	0xcb, 0xe0, 0x9d, 0x7f, 0x07, 0xf7, 0x5b, 0x51, 0x06, 0xbb, 0x5b, 0x5e, 0xf3, 0x3f, 0x78, 0x57,
	0x58, 0xff, 0x4b, 0xd2, 0xd9, 0x35, 0xfd, 0x25, 0xed, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0x7e, 0xba, 0x47, 0x77, 0x03, 0x00, 0x00,
}
