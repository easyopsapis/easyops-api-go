// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_archive_instance.proto

package instance_archive

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
//ListArchiveInstance请求
type ListArchiveInstanceRequest struct {
	//
	//实例所属的模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArchiveInstanceRequest) Reset()         { *m = ListArchiveInstanceRequest{} }
func (m *ListArchiveInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*ListArchiveInstanceRequest) ProtoMessage()    {}
func (*ListArchiveInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a1d8df8dfb01a8, []int{0}
}
func (m *ListArchiveInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArchiveInstanceRequest.Unmarshal(m, b)
}
func (m *ListArchiveInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArchiveInstanceRequest.Marshal(b, m, deterministic)
}
func (m *ListArchiveInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArchiveInstanceRequest.Merge(m, src)
}
func (m *ListArchiveInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_ListArchiveInstanceRequest.Size(m)
}
func (m *ListArchiveInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArchiveInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArchiveInstanceRequest proto.InternalMessageInfo

func (m *ListArchiveInstanceRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ListArchiveInstanceRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListArchiveInstanceRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListArchiveInstance返回
type ListArchiveInstanceResponse struct {
	//
	//instance list
	List []*types.Struct `protobuf:"bytes,1,rep,name=list,proto3" json:"list" form:"list"`
	//
	//实例总数
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total"`
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

func (m *ListArchiveInstanceResponse) Reset()         { *m = ListArchiveInstanceResponse{} }
func (m *ListArchiveInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*ListArchiveInstanceResponse) ProtoMessage()    {}
func (*ListArchiveInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a1d8df8dfb01a8, []int{1}
}
func (m *ListArchiveInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArchiveInstanceResponse.Unmarshal(m, b)
}
func (m *ListArchiveInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArchiveInstanceResponse.Marshal(b, m, deterministic)
}
func (m *ListArchiveInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArchiveInstanceResponse.Merge(m, src)
}
func (m *ListArchiveInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_ListArchiveInstanceResponse.Size(m)
}
func (m *ListArchiveInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArchiveInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArchiveInstanceResponse proto.InternalMessageInfo

func (m *ListArchiveInstanceResponse) GetList() []*types.Struct {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListArchiveInstanceResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListArchiveInstanceResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListArchiveInstanceResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListArchiveInstanceApi返回
type ListArchiveInstanceResponseWrapper struct {
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
	Data                 *ListArchiveInstanceResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListArchiveInstanceResponseWrapper) Reset()         { *m = ListArchiveInstanceResponseWrapper{} }
func (m *ListArchiveInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListArchiveInstanceResponseWrapper) ProtoMessage()    {}
func (*ListArchiveInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a1d8df8dfb01a8, []int{2}
}
func (m *ListArchiveInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArchiveInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *ListArchiveInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArchiveInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListArchiveInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArchiveInstanceResponseWrapper.Merge(m, src)
}
func (m *ListArchiveInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListArchiveInstanceResponseWrapper.Size(m)
}
func (m *ListArchiveInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArchiveInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListArchiveInstanceResponseWrapper proto.InternalMessageInfo

func (m *ListArchiveInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListArchiveInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListArchiveInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListArchiveInstanceResponseWrapper) GetData() *ListArchiveInstanceResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListArchiveInstanceRequest)(nil), "instance_archive.ListArchiveInstanceRequest")
	proto.RegisterType((*ListArchiveInstanceResponse)(nil), "instance_archive.ListArchiveInstanceResponse")
	proto.RegisterType((*ListArchiveInstanceResponseWrapper)(nil), "instance_archive.ListArchiveInstanceResponseWrapper")
}

func init() { proto.RegisterFile("list_archive_instance.proto", fileDescriptor_99a1d8df8dfb01a8) }

var fileDescriptor_99a1d8df8dfb01a8 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x76, 0x9b, 0x8d, 0x74, 0x27, 0x82, 0x61, 0x40, 0x0d, 0xa9, 0xb2, 0x61, 0x14, 0x89, 0xe0,
	0x6e, 0x6a, 0x0b, 0x56, 0xc5, 0x4b, 0x03, 0x1e, 0x0a, 0x9e, 0xa6, 0x82, 0x60, 0xd1, 0x65, 0xb2,
	0x3b, 0xdd, 0x8e, 0x6e, 0x32, 0xeb, 0xcc, 0xa4, 0x95, 0x8a, 0x7f, 0xd3, 0x63, 0x04, 0x2f, 0xde,
	0xf3, 0x07, 0x94, 0x79, 0x93, 0xcd, 0x2e, 0x52, 0x0a, 0x39, 0xed, 0xbc, 0xef, 0xfb, 0xde, 0xc7,
	0xfb, 0xde, 0x63, 0xd1, 0x4e, 0x21, 0xb4, 0x49, 0x98, 0x4a, 0xcf, 0xc4, 0x39, 0x4f, 0xc4, 0x4c,
	0x1b, 0x36, 0x4b, 0x79, 0x5c, 0x2a, 0x69, 0x24, 0xee, 0x56, 0x75, 0x25, 0xe8, 0x47, 0xb9, 0x30,
	0x67, 0xf3, 0x49, 0x9c, 0xca, 0xe9, 0x28, 0x97, 0xb9, 0x1c, 0x81, 0x70, 0x32, 0x3f, 0x85, 0x0a,
	0x0a, 0x78, 0x39, 0x83, 0xfe, 0xf3, 0x86, 0x7c, 0x7a, 0x21, 0xcc, 0x17, 0x79, 0x31, 0xca, 0x65,
	0x04, 0x64, 0x74, 0xce, 0x0a, 0x91, 0x31, 0x23, 0x95, 0x1e, 0xad, 0x9f, 0xab, 0xbe, 0xfb, 0xb9,
	0x94, 0x79, 0xc1, 0x6b, 0x77, 0x6d, 0xd4, 0x3c, 0x35, 0x8e, 0x25, 0x3f, 0x3d, 0xd4, 0x7f, 0x2b,
	0xb4, 0x39, 0x74, 0x43, 0x1d, 0xad, 0x86, 0xa4, 0xfc, 0xeb, 0x9c, 0x6b, 0x83, 0xdf, 0xa1, 0x40,
	0x4e, 0x3e, 0xf3, 0xd4, 0x24, 0x22, 0xeb, 0x79, 0x03, 0x6f, 0x18, 0x8c, 0x0f, 0x96, 0x8b, 0xb0,
	0x7b, 0x2a, 0xd5, 0xf4, 0x15, 0x59, 0x53, 0xe4, 0xf7, 0xaf, 0x30, 0x44, 0x0f, 0x3e, 0x9d, 0xb0,
	0xe8, 0xf2, 0x30, 0xfa, 0x90, 0x7c, 0x3c, 0xd9, 0x8d, 0x5e, 0x56, 0xef, 0xef, 0xbb, 0x4f, 0xf7,
	0x9f, 0xfd, 0x78, 0x44, 0xb7, 0x9d, 0xfc, 0x28, 0xc3, 0x4f, 0x90, 0x5f, 0xb2, 0x9c, 0xf7, 0xb6,
	0x06, 0xde, 0xb0, 0x3d, 0xbe, 0xb3, 0x5c, 0x84, 0x1d, 0x67, 0x68, 0x51, 0xeb, 0xb5, 0xd5, 0xbd,
	0x41, 0x41, 0x82, 0x0f, 0x50, 0x60, 0xbf, 0x89, 0x16, 0x97, 0xbc, 0xd7, 0x02, 0x7d, 0xbf, 0x1e,
	0x60, 0x4d, 0x55, 0x4d, 0xdb, 0x16, 0x39, 0xb6, 0xc0, 0x1f, 0x0f, 0xed, 0x5c, 0x19, 0x4c, 0x97,
	0x72, 0xa6, 0x39, 0x7e, 0x8d, 0x7c, 0x7b, 0xae, 0x9e, 0x37, 0x68, 0x0d, 0x3b, 0x7b, 0xf7, 0x62,
	0xb7, 0xa5, 0xb8, 0xda, 0x52, 0x7c, 0x0c, 0x5b, 0x1a, 0xdf, 0xae, 0x87, 0xb3, 0x72, 0x42, 0xa1,
	0x0b, 0x3f, 0x46, 0x6d, 0x23, 0x0d, 0x2b, 0x56, 0x11, 0xba, 0xcb, 0x45, 0x78, 0xcb, 0xa9, 0x00,
	0x26, 0xd4, 0xd1, 0xeb, 0xa4, 0xad, 0x0d, 0x93, 0xfa, 0x1b, 0x24, 0xfd, 0xeb, 0x21, 0x72, 0x4d,
	0xd2, 0xf7, 0x8a, 0x95, 0x25, 0x57, 0xf8, 0x21, 0xf2, 0x53, 0x99, 0x71, 0xb8, 0x62, 0xbb, 0x99,
	0xcb, 0xa2, 0x84, 0x02, 0x89, 0x5f, 0xa0, 0x8e, 0xfd, 0xbe, 0xf9, 0x56, 0x16, 0x4c, 0xcc, 0x20,
	0x5d, 0x30, 0xbe, 0xbb, 0x5c, 0x84, 0xb8, 0xd6, 0xae, 0x48, 0x42, 0x9b, 0x52, 0xbb, 0x11, 0xae,
	0x94, 0x54, 0x10, 0x35, 0x68, 0x6e, 0x04, 0x60, 0x42, 0x1d, 0x8d, 0x29, 0xf2, 0x33, 0x66, 0x18,
	0x24, 0xec, 0xec, 0x45, 0xf1, 0xff, 0xbf, 0x45, 0x7c, 0x4d, 0x94, 0xe6, 0xd4, 0xd6, 0x84, 0x50,
	0xf0, 0x9a, 0xdc, 0x84, 0xab, 0xed, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xf7, 0x7b, 0x7b,
	0x81, 0x03, 0x00, 0x00,
}