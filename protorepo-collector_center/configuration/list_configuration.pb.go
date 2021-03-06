// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_configuration.proto

package configuration

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	collector_center "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/collector_center"
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
//ListConfiguration请求
type ListConfigurationRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//是否全部
	IsAll                bool     `protobuf:"varint,3,opt,name=isAll,proto3" json:"isAll" form:"isAll"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConfigurationRequest) Reset()         { *m = ListConfigurationRequest{} }
func (m *ListConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationRequest) ProtoMessage()    {}
func (*ListConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_669564e7f48e8437, []int{0}
}
func (m *ListConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationRequest.Unmarshal(m, b)
}
func (m *ListConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationRequest.Marshal(b, m, deterministic)
}
func (m *ListConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationRequest.Merge(m, src)
}
func (m *ListConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationRequest.Size(m)
}
func (m *ListConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationRequest proto.InternalMessageInfo

func (m *ListConfigurationRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListConfigurationRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConfigurationRequest) GetIsAll() bool {
	if m != nil {
		return m.IsAll
	}
	return false
}

//
//ListConfiguration返回
type ListConfigurationResponse struct {
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
	List                 []*collector_center.Configuration `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListConfigurationResponse) Reset()         { *m = ListConfigurationResponse{} }
func (m *ListConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationResponse) ProtoMessage()    {}
func (*ListConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_669564e7f48e8437, []int{1}
}
func (m *ListConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationResponse.Unmarshal(m, b)
}
func (m *ListConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationResponse.Marshal(b, m, deterministic)
}
func (m *ListConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationResponse.Merge(m, src)
}
func (m *ListConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationResponse.Size(m)
}
func (m *ListConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationResponse proto.InternalMessageInfo

func (m *ListConfigurationResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListConfigurationResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConfigurationResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListConfigurationResponse) GetList() []*collector_center.Configuration {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListConfigurationApi返回
type ListConfigurationResponseWrapper struct {
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
	Data                 *ListConfigurationResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListConfigurationResponseWrapper) Reset()         { *m = ListConfigurationResponseWrapper{} }
func (m *ListConfigurationResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListConfigurationResponseWrapper) ProtoMessage()    {}
func (*ListConfigurationResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_669564e7f48e8437, []int{2}
}
func (m *ListConfigurationResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigurationResponseWrapper.Unmarshal(m, b)
}
func (m *ListConfigurationResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigurationResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListConfigurationResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigurationResponseWrapper.Merge(m, src)
}
func (m *ListConfigurationResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListConfigurationResponseWrapper.Size(m)
}
func (m *ListConfigurationResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigurationResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigurationResponseWrapper proto.InternalMessageInfo

func (m *ListConfigurationResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListConfigurationResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListConfigurationResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListConfigurationResponseWrapper) GetData() *ListConfigurationResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListConfigurationRequest)(nil), "configuration.ListConfigurationRequest")
	proto.RegisterType((*ListConfigurationResponse)(nil), "configuration.ListConfigurationResponse")
	proto.RegisterType((*ListConfigurationResponseWrapper)(nil), "configuration.ListConfigurationResponseWrapper")
}

func init() { proto.RegisterFile("list_configuration.proto", fileDescriptor_669564e7f48e8437) }

var fileDescriptor_669564e7f48e8437 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3d, 0x8f, 0xd3, 0x4c,
	0x10, 0x7e, 0x7d, 0xe7, 0xbc, 0xba, 0x6c, 0x40, 0x17, 0x59, 0x02, 0xf9, 0xd2, 0xd8, 0x5a, 0x24,
	0x64, 0x8a, 0xd8, 0xd2, 0x21, 0x3e, 0x44, 0x47, 0x80, 0x0e, 0x9a, 0xa5, 0xa0, 0x23, 0xda, 0x6c,
	0x36, 0x66, 0xc5, 0xc6, 0x63, 0x76, 0x37, 0x1c, 0xf0, 0x7b, 0xf8, 0x5d, 0x46, 0xa2, 0xa7, 0x71,
	0x45, 0x89, 0x76, 0x7c, 0x77, 0x76, 0x50, 0xae, 0xa0, 0x5a, 0xcf, 0xf3, 0x31, 0x7a, 0x1e, 0x79,
	0x48, 0xac, 0x95, 0x75, 0x4b, 0x01, 0xd5, 0x46, 0x95, 0x3b, 0xc3, 0x9d, 0x82, 0x2a, 0xaf, 0x0d,
	0x38, 0x88, 0x6e, 0xef, 0x81, 0xb3, 0x79, 0xa9, 0xdc, 0x87, 0xdd, 0x2a, 0x17, 0xb0, 0x2d, 0x4a,
	0x28, 0xa1, 0x40, 0xd5, 0x6a, 0xb7, 0xc1, 0x09, 0x07, 0xfc, 0xea, 0xdc, 0xb3, 0xf7, 0x25, 0xe4,
	0x92, 0xdb, 0xaf, 0x50, 0xdb, 0x5c, 0x83, 0xe0, 0xba, 0x10, 0x50, 0x39, 0xc3, 0x85, 0xb3, 0x9d,
	0xd3, 0xc8, 0x1a, 0xe6, 0x5b, 0x58, 0x4b, 0x6d, 0x8b, 0x4b, 0x61, 0x81, 0x63, 0x21, 0x40, 0x6b,
	0x29, 0x1c, 0x98, 0xa5, 0x90, 0x95, 0x93, 0xa6, 0x38, 0x90, 0x6e, 0xf6, 0x78, 0x10, 0x67, 0x7b,
	0xa1, 0xdc, 0x47, 0xb8, 0x28, 0x4a, 0x98, 0x23, 0x39, 0xff, 0xcc, 0xb5, 0x5a, 0x73, 0x07, 0xc6,
	0x16, 0xd7, 0x9f, 0x9d, 0x8f, 0x7e, 0x0f, 0x48, 0xfc, 0x5a, 0x59, 0xf7, 0x62, 0xb8, 0x93, 0xc9,
	0x4f, 0x3b, 0x69, 0x5d, 0xf4, 0x80, 0x84, 0x35, 0x2f, 0x65, 0x1c, 0xa4, 0x41, 0x36, 0x5a, 0xdc,
	0x69, 0x9b, 0x64, 0xb2, 0x01, 0xb3, 0x7d, 0x46, 0x3d, 0x4a, 0x7f, 0xfe, 0x48, 0x8e, 0xa6, 0xff,
	0x31, 0x94, 0x44, 0x8f, 0xc8, 0x89, 0x7f, 0xdf, 0xaa, 0x6f, 0x32, 0x3e, 0x42, 0xf9, 0x59, 0xdb,
	0x24, 0xa7, 0xbd, 0xdc, 0x33, 0x57, 0x96, 0x6b, 0x69, 0x74, 0x9f, 0x8c, 0x94, 0x7d, 0xae, 0x75,
	0x7c, 0x9c, 0x06, 0xd9, 0xc9, 0x62, 0xda, 0x36, 0xc9, 0xad, 0xce, 0x83, 0x30, 0x65, 0x1d, 0x4d,
	0x7f, 0x05, 0xe4, 0xec, 0x40, 0x4c, 0x5b, 0x43, 0x65, 0xe5, 0xbf, 0xe4, 0x7c, 0x42, 0xc6, 0xfe,
	0x5d, 0xda, 0x3e, 0xe8, 0xac, 0x6d, 0x92, 0x69, 0xaf, 0x47, 0xea, 0x60, 0x52, 0x07, 0x8e, 0x77,
	0x49, 0x47, 0xc3, 0xa4, 0x08, 0x53, 0xd6, 0xd1, 0xd1, 0x4b, 0x12, 0xfa, 0x13, 0x8a, 0xc3, 0xf4,
	0x38, 0x9b, 0x9c, 0x27, 0xf9, 0xdf, 0x7f, 0x2f, 0xdf, 0xab, 0xb0, 0x38, 0xed, 0xc3, 0x7a, 0x1b,
	0x65, 0xe8, 0xa6, 0xbf, 0x03, 0x92, 0xde, 0xd8, 0xf7, 0x9d, 0xe1, 0x75, 0x2d, 0x4d, 0x74, 0x8f,
	0x84, 0x02, 0xd6, 0x57, 0xb5, 0x07, 0x9b, 0x3c, 0x4a, 0x19, 0x92, 0xd1, 0x53, 0x32, 0xf1, 0xef,
	0xab, 0x2f, 0xb5, 0xe6, 0xaa, 0xc2, 0xca, 0xe3, 0xc5, 0xdd, 0xb6, 0x49, 0xa2, 0x5e, 0x7b, 0x49,
	0x52, 0x36, 0x94, 0xfa, 0xc6, 0xd2, 0x18, 0x30, 0xd8, 0x78, 0x3c, 0x6c, 0x8c, 0x30, 0x65, 0x1d,
	0x1d, 0xbd, 0x21, 0xe1, 0x9a, 0x3b, 0x1e, 0x87, 0x69, 0x90, 0x4d, 0xce, 0xb3, 0x7c, 0xff, 0x3c,
	0x6f, 0x6c, 0x31, 0x0c, 0xec, 0xfd, 0x94, 0xe1, 0x9a, 0xd5, 0xff, 0x78, 0x98, 0x0f, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x80, 0x54, 0x4d, 0x6d, 0x8a, 0x03, 0x00, 0x00,
}
