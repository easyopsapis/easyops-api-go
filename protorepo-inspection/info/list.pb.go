// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package info

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inspection "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/inspection"
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
//ListInspectionInfo请求
type ListInspectionInfoRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页大小
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListInspectionInfoRequest) Reset()         { *m = ListInspectionInfoRequest{} }
func (m *ListInspectionInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ListInspectionInfoRequest) ProtoMessage()    {}
func (*ListInspectionInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListInspectionInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInspectionInfoRequest.Unmarshal(m, b)
}
func (m *ListInspectionInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInspectionInfoRequest.Marshal(b, m, deterministic)
}
func (m *ListInspectionInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInspectionInfoRequest.Merge(m, src)
}
func (m *ListInspectionInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ListInspectionInfoRequest.Size(m)
}
func (m *ListInspectionInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInspectionInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInspectionInfoRequest proto.InternalMessageInfo

func (m *ListInspectionInfoRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListInspectionInfoRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListInspectionInfo返回
type ListInspectionInfoResponse struct {
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
	List                 []*inspection.InspectionInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListInspectionInfoResponse) Reset()         { *m = ListInspectionInfoResponse{} }
func (m *ListInspectionInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ListInspectionInfoResponse) ProtoMessage()    {}
func (*ListInspectionInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListInspectionInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInspectionInfoResponse.Unmarshal(m, b)
}
func (m *ListInspectionInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInspectionInfoResponse.Marshal(b, m, deterministic)
}
func (m *ListInspectionInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInspectionInfoResponse.Merge(m, src)
}
func (m *ListInspectionInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ListInspectionInfoResponse.Size(m)
}
func (m *ListInspectionInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInspectionInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInspectionInfoResponse proto.InternalMessageInfo

func (m *ListInspectionInfoResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListInspectionInfoResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListInspectionInfoResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListInspectionInfoResponse) GetList() []*inspection.InspectionInfo {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListInspectionInfoApi返回
type ListInspectionInfoResponseWrapper struct {
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
	Data                 *ListInspectionInfoResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListInspectionInfoResponseWrapper) Reset()         { *m = ListInspectionInfoResponseWrapper{} }
func (m *ListInspectionInfoResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListInspectionInfoResponseWrapper) ProtoMessage()    {}
func (*ListInspectionInfoResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}
func (m *ListInspectionInfoResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInspectionInfoResponseWrapper.Unmarshal(m, b)
}
func (m *ListInspectionInfoResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInspectionInfoResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListInspectionInfoResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInspectionInfoResponseWrapper.Merge(m, src)
}
func (m *ListInspectionInfoResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListInspectionInfoResponseWrapper.Size(m)
}
func (m *ListInspectionInfoResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInspectionInfoResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListInspectionInfoResponseWrapper proto.InternalMessageInfo

func (m *ListInspectionInfoResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListInspectionInfoResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListInspectionInfoResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListInspectionInfoResponseWrapper) GetData() *ListInspectionInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListInspectionInfoRequest)(nil), "info.ListInspectionInfoRequest")
	proto.RegisterType((*ListInspectionInfoResponse)(nil), "info.ListInspectionInfoResponse")
	proto.RegisterType((*ListInspectionInfoResponseWrapper)(nil), "info.ListInspectionInfoResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x8b, 0x13, 0x31,
	0x14, 0x76, 0x76, 0x67, 0xc5, 0xa6, 0xc2, 0x96, 0x80, 0x32, 0x3b, 0x97, 0xa9, 0x11, 0xa4, 0x1e,
	0x3a, 0x81, 0x15, 0x7f, 0xe0, 0x45, 0x28, 0xec, 0x61, 0x41, 0x10, 0xe2, 0xc1, 0xa3, 0xa4, 0xd3,
	0x74, 0x0c, 0x4e, 0xe7, 0xc5, 0x24, 0x75, 0x55, 0xf0, 0x5f, 0x1d, 0xc1, 0x9b, 0x47, 0xe7, 0x2f,
	0x90, 0xbc, 0xe9, 0x76, 0x46, 0xb4, 0x87, 0x3d, 0x25, 0xef, 0x7d, 0xdf, 0xf7, 0xde, 0xf7, 0x91,
	0x10, 0x52, 0x69, 0xe7, 0x73, 0x63, 0xc1, 0x03, 0x8d, 0x75, 0xbd, 0x86, 0x74, 0x5e, 0x6a, 0xff,
	0x61, 0xbb, 0xcc, 0x0b, 0xd8, 0xf0, 0x12, 0x4a, 0xe0, 0x08, 0x2e, 0xb7, 0x6b, 0xac, 0xb0, 0xc0,
	0x5b, 0x27, 0x4a, 0xdf, 0x94, 0x90, 0x2b, 0xe9, 0xbe, 0x82, 0x71, 0x79, 0x05, 0x85, 0xac, 0x78,
	0x01, 0xb5, 0xb7, 0xb2, 0xf0, 0xae, 0x53, 0x5a, 0x65, 0x60, 0xbe, 0x81, 0x95, 0xaa, 0x1c, 0xdf,
	0x11, 0x39, 0x96, 0x5c, 0xd7, 0xce, 0xa8, 0xc2, 0x6b, 0xa8, 0x79, 0xd8, 0xbc, 0x1b, 0xf8, 0x6c,
	0xb0, 0x7f, 0x73, 0xa5, 0xfd, 0x47, 0xb8, 0xe2, 0x25, 0xcc, 0x11, 0x9c, 0x7f, 0x96, 0x95, 0x5e,
	0x49, 0x0f, 0xd6, 0xf1, 0xfd, 0xb5, 0xd3, 0xb1, 0xef, 0xe4, 0xec, 0xb5, 0x76, 0xfe, 0x72, 0x3f,
	0xf4, 0xb2, 0x5e, 0x83, 0x50, 0x9f, 0xb6, 0xca, 0x79, 0xfa, 0x98, 0xc4, 0x46, 0x96, 0x2a, 0x89,
	0xa6, 0xd1, 0xec, 0x64, 0x71, 0xaf, 0x6d, 0xb2, 0xf1, 0x1a, 0xec, 0xe6, 0x25, 0x0b, 0x5d, 0xf6,
	0xf3, 0x47, 0x76, 0x34, 0xb9, 0x25, 0x90, 0x42, 0x9f, 0x92, 0x3b, 0xe1, 0x7c, 0xab, 0xbf, 0xa9,
	0xe4, 0x08, 0xe9, 0x67, 0x6d, 0x93, 0x9d, 0xf6, 0xf4, 0x80, 0x5c, 0x4b, 0xf6, 0x54, 0xf6, 0x2b,
	0x22, 0xe9, 0xff, 0xf6, 0x3b, 0x03, 0xb5, 0x53, 0x37, 0x31, 0xf0, 0x9c, 0x8c, 0xc2, 0xf9, 0xde,
	0xf5, 0x0e, 0xd2, 0xb6, 0xc9, 0x26, 0x3d, 0x1f, 0xa1, 0x7f, 0x2c, 0xd0, 0x47, 0xe4, 0xc4, 0x83,
	0x97, 0x55, 0x72, 0x8c, 0xa2, 0x49, 0xdb, 0x64, 0x77, 0x3b, 0x11, 0xb6, 0x99, 0xe8, 0x60, 0xfa,
	0x8a, 0xc4, 0xe1, 0xd5, 0x93, 0x78, 0x7a, 0x3c, 0x1b, 0x9f, 0xa7, 0x79, 0xff, 0x0e, 0xf9, 0xdf,
	0xee, 0x17, 0xa7, 0xbd, 0xcf, 0xa0, 0x60, 0x02, 0x85, 0xec, 0x77, 0x44, 0x1e, 0x1c, 0xce, 0xfa,
	0xce, 0x4a, 0x63, 0x94, 0xa5, 0x0f, 0x49, 0x5c, 0xc0, 0xea, 0x3a, 0xf2, 0x60, 0x54, 0xe8, 0x32,
	0x81, 0x20, 0x7d, 0x41, 0xc6, 0xe1, 0xbc, 0xf8, 0x62, 0x2a, 0xa9, 0x6b, 0x8c, 0x3b, 0x5a, 0xdc,
	0x6f, 0x9b, 0x8c, 0xf6, 0xdc, 0x1d, 0xc8, 0xc4, 0x90, 0x1a, 0xd2, 0x2a, 0x6b, 0xc1, 0x62, 0xda,
	0xd1, 0x30, 0x2d, 0xb6, 0x99, 0xe8, 0x60, 0x7a, 0x41, 0xe2, 0x95, 0xf4, 0x32, 0x89, 0xa7, 0xd1,
	0x6c, 0x7c, 0x3e, 0xcd, 0xf1, 0xab, 0x1d, 0x76, 0x3f, 0x34, 0x1a, 0x74, 0x4c, 0xa0, 0x7c, 0x79,
	0x1b, 0x7f, 0xd9, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc6, 0xc4, 0xb7, 0x31, 0x03,
	0x00, 0x00,
}
