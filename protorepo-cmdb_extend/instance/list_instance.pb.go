// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_instance.proto

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
//ListInstance请求
type ListInstanceRequest struct {
	//
	//模型Id
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//查询参数
	Query                *types.Struct `protobuf:"bytes,4,opt,name=query,proto3" json:"query" form:"query"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListInstanceRequest) Reset()         { *m = ListInstanceRequest{} }
func (m *ListInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstanceRequest) ProtoMessage()    {}
func (*ListInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f7236325836c5ab, []int{0}
}
func (m *ListInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceRequest.Unmarshal(m, b)
}
func (m *ListInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceRequest.Marshal(b, m, deterministic)
}
func (m *ListInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceRequest.Merge(m, src)
}
func (m *ListInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_ListInstanceRequest.Size(m)
}
func (m *ListInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceRequest proto.InternalMessageInfo

func (m *ListInstanceRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ListInstanceRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListInstanceRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListInstanceRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

//
//ListInstance返回
type ListInstanceResponse struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//实例数据
	List                 []*types.Struct `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListInstanceResponse) Reset()         { *m = ListInstanceResponse{} }
func (m *ListInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstanceResponse) ProtoMessage()    {}
func (*ListInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f7236325836c5ab, []int{1}
}
func (m *ListInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceResponse.Unmarshal(m, b)
}
func (m *ListInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceResponse.Marshal(b, m, deterministic)
}
func (m *ListInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceResponse.Merge(m, src)
}
func (m *ListInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_ListInstanceResponse.Size(m)
}
func (m *ListInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceResponse proto.InternalMessageInfo

func (m *ListInstanceResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListInstanceResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListInstanceResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListInstanceResponse) GetList() []*types.Struct {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListInstanceApi返回
type ListInstanceResponseWrapper struct {
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
	Data                 *ListInstanceResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListInstanceResponseWrapper) Reset()         { *m = ListInstanceResponseWrapper{} }
func (m *ListInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListInstanceResponseWrapper) ProtoMessage()    {}
func (*ListInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f7236325836c5ab, []int{2}
}
func (m *ListInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *ListInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceResponseWrapper.Merge(m, src)
}
func (m *ListInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListInstanceResponseWrapper.Size(m)
}
func (m *ListInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceResponseWrapper proto.InternalMessageInfo

func (m *ListInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListInstanceResponseWrapper) GetData() *ListInstanceResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListInstanceRequest)(nil), "instance.ListInstanceRequest")
	proto.RegisterType((*ListInstanceResponse)(nil), "instance.ListInstanceResponse")
	proto.RegisterType((*ListInstanceResponseWrapper)(nil), "instance.ListInstanceResponseWrapper")
}

func init() { proto.RegisterFile("list_instance.proto", fileDescriptor_1f7236325836c5ab) }

var fileDescriptor_1f7236325836c5ab = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x37, 0x95, 0x76, 0x2a, 0xac, 0x4c, 0x45, 0xc3, 0xfa, 0x27, 0x65, 0x14, 0xe9,
	0x85, 0x49, 0x56, 0x17, 0x16, 0x15, 0x41, 0xac, 0x78, 0xb1, 0xe0, 0xd5, 0xec, 0x85, 0xe0, 0xa2,
	0xcb, 0x34, 0x9d, 0x8d, 0xd1, 0xb4, 0x27, 0x3b, 0x99, 0xb8, 0xba, 0x22, 0xf8, 0x8c, 0x3e, 0x40,
	0x05, 0xef, 0xbc, 0xed, 0x13, 0xc8, 0x9c, 0x49, 0xd2, 0x08, 0x65, 0xd9, 0xab, 0x99, 0x39, 0xdf,
	0x77, 0x4e, 0xbe, 0xdf, 0x24, 0x21, 0xc3, 0x2c, 0x2d, 0xf4, 0x71, 0xba, 0x28, 0xb4, 0x58, 0xc4,
	0x32, 0xcc, 0x15, 0x68, 0xa0, 0xbd, 0xfa, 0xbc, 0x13, 0x24, 0xa9, 0xfe, 0x58, 0x4e, 0xc3, 0x18,
	0xe6, 0x51, 0x02, 0x09, 0x44, 0x68, 0x98, 0x96, 0x27, 0x78, 0xc2, 0x03, 0xee, 0x6c, 0xe3, 0xce,
	0x7e, 0xcb, 0x3e, 0x3f, 0x4b, 0xf5, 0x67, 0x38, 0x8b, 0x12, 0x08, 0x50, 0x0c, 0xbe, 0x88, 0x2c,
	0x9d, 0x09, 0x0d, 0xaa, 0x88, 0x9a, 0x6d, 0xd5, 0x77, 0x3b, 0x01, 0x48, 0x32, 0xb9, 0x9e, 0x5e,
	0x68, 0x55, 0xc6, 0xda, 0xaa, 0xec, 0x67, 0x87, 0x0c, 0xdf, 0xa4, 0x85, 0x3e, 0xa8, 0x52, 0x71,
	0x79, 0x5a, 0xca, 0x42, 0x53, 0x4e, 0x7a, 0x30, 0xfd, 0x24, 0x63, 0x7d, 0x30, 0xf3, 0x9c, 0x91,
	0x33, 0xee, 0x4f, 0xf6, 0x57, 0x4b, 0x7f, 0xfb, 0x04, 0xd4, 0xfc, 0x19, 0xab, 0x15, 0xf6, 0xe7,
	0xb7, 0xef, 0x93, 0x3b, 0x1f, 0x8e, 0x44, 0x70, 0xfe, 0x32, 0x78, 0x77, 0xfc, 0xfe, 0x68, 0x37,
	0x78, 0x5a, 0xef, 0xbf, 0xef, 0x3e, 0xdc, 0x7b, 0xf4, 0xe3, 0x3e, 0x6f, 0xe6, 0xd0, 0x7b, 0xc4,
	0xcd, 0x45, 0x22, 0xbd, 0xce, 0xc8, 0x19, 0x77, 0x27, 0xdb, 0xab, 0xa5, 0x3f, 0xb0, 0xf3, 0x4c,
	0x95, 0x71, 0x14, 0x69, 0x44, 0x7a, 0x66, 0x3d, 0x4c, 0xcf, 0xa5, 0xb7, 0x85, 0xc6, 0xe1, 0xfa,
	0xc1, 0xb5, 0xc2, 0x78, 0x63, 0xa2, 0x2f, 0x48, 0xf7, 0xb4, 0x94, 0xea, 0x9b, 0xe7, 0x8e, 0x9c,
	0xf1, 0xe0, 0xf1, 0xcd, 0xd0, 0xf2, 0x86, 0x35, 0x6f, 0x78, 0x88, 0xbc, 0x93, 0x6b, 0xab, 0xa5,
	0x7f, 0xd5, 0x8e, 0x41, 0x3f, 0xe3, 0xb6, 0x8f, 0xfd, 0x72, 0xc8, 0xf5, 0xff, 0xaf, 0xa0, 0xc8,
	0x61, 0x51, 0xc8, 0x26, 0xaf, 0x73, 0xd9, 0xbc, 0x9d, 0xcb, 0xe4, 0x7d, 0x40, 0xba, 0x1a, 0xb4,
	0xc8, 0x2a, 0xba, 0x56, 0x2c, 0x2c, 0x33, 0x6e, 0x65, 0xfa, 0x9c, 0xb8, 0xe6, 0xfb, 0xf1, 0xdc,
	0xd1, 0xd6, 0x45, 0x58, 0xad, 0x58, 0xc6, 0xce, 0x38, 0x76, 0xb1, 0xbf, 0x0e, 0xb9, 0xb5, 0x09,
	0xea, 0xad, 0x12, 0x79, 0x2e, 0x95, 0x61, 0x8b, 0x61, 0xb6, 0x81, 0xcd, 0x54, 0x19, 0x47, 0x91,
	0x3e, 0x21, 0x03, 0xb3, 0xbe, 0xfe, 0x9a, 0x67, 0x22, 0x5d, 0x20, 0x5e, 0x7f, 0x72, 0x63, 0xb5,
	0xf4, 0xe9, 0xda, 0x5b, 0x89, 0x8c, 0xb7, 0xad, 0x06, 0x52, 0x2a, 0x05, 0x0a, 0x21, 0xfb, 0x6d,
	0x48, 0x2c, 0x33, 0x6e, 0x65, 0xfa, 0x8a, 0xb8, 0x33, 0xa1, 0x45, 0xf5, 0xee, 0xee, 0x86, 0xcd,
	0xcf, 0xb2, 0x29, 0x7b, 0x3b, 0xa6, 0xe9, 0x62, 0x1c, 0x9b, 0xa7, 0x57, 0xf0, 0x4e, 0xf6, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x14, 0x23, 0xca, 0x70, 0x03, 0x00, 0x00,
}
