// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_service_info.proto

package service_manage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	resource_manage "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/resource_manage"
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
//ListServiceInfo请求
type ListServiceInfoRequest struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//节点类型
	NodeType             string   `protobuf:"bytes,2,opt,name=nodeType,proto3" json:"nodeType" form:"nodeType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServiceInfoRequest) Reset()         { *m = ListServiceInfoRequest{} }
func (m *ListServiceInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ListServiceInfoRequest) ProtoMessage()    {}
func (*ListServiceInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_281aa3f72f53560e, []int{0}
}
func (m *ListServiceInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceInfoRequest.Unmarshal(m, b)
}
func (m *ListServiceInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceInfoRequest.Marshal(b, m, deterministic)
}
func (m *ListServiceInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceInfoRequest.Merge(m, src)
}
func (m *ListServiceInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ListServiceInfoRequest.Size(m)
}
func (m *ListServiceInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceInfoRequest proto.InternalMessageInfo

func (m *ListServiceInfoRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ListServiceInfoRequest) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

//
//ListServiceInfo返回
type ListServiceInfoResponse struct {
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
	List                 []*resource_manage.ServiceInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListServiceInfoResponse) Reset()         { *m = ListServiceInfoResponse{} }
func (m *ListServiceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ListServiceInfoResponse) ProtoMessage()    {}
func (*ListServiceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_281aa3f72f53560e, []int{1}
}
func (m *ListServiceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceInfoResponse.Unmarshal(m, b)
}
func (m *ListServiceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceInfoResponse.Marshal(b, m, deterministic)
}
func (m *ListServiceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceInfoResponse.Merge(m, src)
}
func (m *ListServiceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ListServiceInfoResponse.Size(m)
}
func (m *ListServiceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceInfoResponse proto.InternalMessageInfo

func (m *ListServiceInfoResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListServiceInfoResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListServiceInfoResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListServiceInfoResponse) GetList() []*resource_manage.ServiceInfo {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListServiceInfoApi返回
type ListServiceInfoResponseWrapper struct {
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
	Data                 *ListServiceInfoResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListServiceInfoResponseWrapper) Reset()         { *m = ListServiceInfoResponseWrapper{} }
func (m *ListServiceInfoResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListServiceInfoResponseWrapper) ProtoMessage()    {}
func (*ListServiceInfoResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_281aa3f72f53560e, []int{2}
}
func (m *ListServiceInfoResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceInfoResponseWrapper.Unmarshal(m, b)
}
func (m *ListServiceInfoResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceInfoResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListServiceInfoResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceInfoResponseWrapper.Merge(m, src)
}
func (m *ListServiceInfoResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListServiceInfoResponseWrapper.Size(m)
}
func (m *ListServiceInfoResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceInfoResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceInfoResponseWrapper proto.InternalMessageInfo

func (m *ListServiceInfoResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListServiceInfoResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListServiceInfoResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListServiceInfoResponseWrapper) GetData() *ListServiceInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListServiceInfoRequest)(nil), "service_manage.ListServiceInfoRequest")
	proto.RegisterType((*ListServiceInfoResponse)(nil), "service_manage.ListServiceInfoResponse")
	proto.RegisterType((*ListServiceInfoResponseWrapper)(nil), "service_manage.ListServiceInfoResponseWrapper")
}

func init() { proto.RegisterFile("list_service_info.proto", fileDescriptor_281aa3f72f53560e) }

var fileDescriptor_281aa3f72f53560e = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0x5b, 0x8a, 0x5a, 0x17, 0xb1, 0xca, 0x88, 0xad, 0xaa, 0x80, 0x54, 0x06, 0x41, 0x91,
	0x48, 0x32, 0x36, 0x69, 0x7c, 0xdc, 0x56, 0x89, 0xc3, 0xa4, 0x9d, 0x3c, 0x24, 0x24, 0x06, 0x54,
	0x6e, 0xe2, 0x06, 0x43, 0x9a, 0x37, 0xd8, 0xee, 0xc6, 0x8a, 0xf8, 0x29, 0xfc, 0xb5, 0x22, 0x71,
	0xe2, 0x1c, 0x89, 0x3b, 0xb2, 0xd3, 0xb4, 0x81, 0x69, 0x27, 0xdb, 0xef, 0xf3, 0xe1, 0xe7, 0x69,
	0x1d, 0xb4, 0x93, 0x0a, 0xa5, 0x47, 0x8a, 0xcb, 0x33, 0x11, 0xf1, 0x91, 0xc8, 0x26, 0x10, 0xe4,
	0x12, 0x34, 0xe0, 0x9b, 0xd5, 0x6c, 0xca, 0x32, 0x96, 0xf0, 0x9e, 0x9f, 0x08, 0xfd, 0x71, 0x36,
	0x0e, 0x22, 0x98, 0x86, 0x09, 0x24, 0x10, 0x5a, 0xda, 0x78, 0x36, 0xb1, 0x27, 0x7b, 0xb0, 0xbb,
	0x52, 0xde, 0x7b, 0x97, 0x40, 0xc0, 0x99, 0xba, 0x80, 0x5c, 0x05, 0x29, 0x44, 0x2c, 0x0d, 0x23,
	0xc8, 0xb4, 0x64, 0x91, 0x56, 0xa5, 0x52, 0xf2, 0x1c, 0xfc, 0x29, 0xc4, 0x3c, 0x55, 0xe1, 0x92,
	0x18, 0xda, 0x63, 0x28, 0xb9, 0x82, 0x99, 0x5c, 0xdd, 0x1c, 0x5e, 0x0e, 0xd7, 0x3b, 0xa8, 0x85,
	0x99, 0x9e, 0x0b, 0xfd, 0x19, 0xce, 0xc3, 0x04, 0x7c, 0x0b, 0xfa, 0x67, 0x2c, 0x15, 0x31, 0xd3,
	0x20, 0x55, 0xb8, 0xda, 0x96, 0x3a, 0xf2, 0xc3, 0x41, 0xdb, 0xc7, 0x42, 0xe9, 0x93, 0xd2, 0xf2,
	0x28, 0x9b, 0x00, 0xe5, 0x5f, 0x66, 0x5c, 0x69, 0x4c, 0x51, 0x13, 0xc6, 0x9f, 0x78, 0xa4, 0x8f,
	0xe2, 0xae, 0xd3, 0x77, 0x06, 0xad, 0xe1, 0x41, 0xb1, 0xf0, 0xb6, 0x26, 0x20, 0xa7, 0x2f, 0x49,
	0x85, 0x90, 0x5f, 0x3f, 0x3d, 0x0f, 0xdd, 0xfd, 0x70, 0xca, 0xfc, 0xf9, 0xa1, 0xff, 0x76, 0xf4,
	0xfe, 0x74, 0xd7, 0x7f, 0x51, 0xed, 0xbf, 0xed, 0x3e, 0xd9, 0x7f, 0xfa, 0xfd, 0x01, 0x5d, 0xf9,
	0xe0, 0x10, 0x35, 0x33, 0x88, 0xf9, 0xeb, 0x8b, 0x9c, 0x77, 0x37, 0xac, 0xe7, 0xad, 0xb5, 0x67,
	0x85, 0x10, 0xba, 0x22, 0x91, 0xdf, 0x0e, 0xda, 0xb9, 0x94, 0x4f, 0xe5, 0x90, 0x29, 0x8e, 0x1f,
	0x23, 0x37, 0x67, 0x09, 0xb7, 0xe1, 0x1a, 0xc3, 0xdb, 0xc5, 0xc2, 0x6b, 0x97, 0x46, 0x66, 0x6a,
	0x82, 0x6d, 0x74, 0xae, 0x51, 0x4b, 0xc1, 0xcf, 0x50, 0xcb, 0xac, 0x23, 0x25, 0xe6, 0xe5, 0xc5,
	0x8d, 0x61, 0xaf, 0x58, 0x78, 0x9d, 0x35, 0xdf, 0x42, 0x95, 0xa8, 0x69, 0x26, 0x27, 0x62, 0xce,
	0xf1, 0x43, 0xd4, 0xd0, 0xa0, 0x59, 0xda, 0xdd, 0xb4, 0xa2, 0x4e, 0xb1, 0xf0, 0x6e, 0x94, 0x22,
	0x3b, 0x26, 0xb4, 0x84, 0xf1, 0x21, 0x72, 0xcd, 0xbb, 0xe9, 0xba, 0xfd, 0xcd, 0x41, 0x7b, 0xef,
	0x4e, 0xf0, 0xdf, 0x5f, 0x16, 0xd4, 0xf2, 0x0f, 0xb7, 0xd6, 0x49, 0x8d, 0x86, 0x50, 0x2b, 0x25,
	0x7f, 0x1c, 0x74, 0xef, 0x8a, 0xaa, 0x6f, 0x24, 0xcb, 0x73, 0x2e, 0xf1, 0x7d, 0xe4, 0x46, 0x10,
	0x57, 0x8d, 0x6b, 0x3e, 0x66, 0x4a, 0xa8, 0x05, 0xf1, 0x73, 0xd4, 0x36, 0xeb, 0xab, 0xaf, 0x79,
	0xca, 0x44, 0xb6, 0xfc, 0x99, 0xb7, 0x8b, 0x85, 0x87, 0xd7, 0xdc, 0x25, 0x48, 0x68, 0x9d, 0x6a,
	0xca, 0x72, 0x29, 0x41, 0xda, 0xb2, 0xad, 0x7a, 0x59, 0x3b, 0x26, 0xb4, 0x84, 0xf1, 0x31, 0x72,
	0x63, 0xa6, 0x59, 0xd7, 0xed, 0x3b, 0x83, 0xf6, 0xde, 0xa3, 0xe0, 0xdf, 0x0f, 0x23, 0xb8, 0xa2,
	0x44, 0x3d, 0xaf, 0x91, 0x13, 0x6a, 0x5d, 0xc6, 0xd7, 0xed, 0x4b, 0xdc, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xc8, 0x8a, 0xdf, 0x79, 0x03, 0x00, 0x00,
}
