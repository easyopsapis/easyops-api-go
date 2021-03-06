// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice_config.proto

package legacy

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//NameServiceConfig请求
type NameServiceConfigRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameServiceConfigRequest) Reset()         { *m = NameServiceConfigRequest{} }
func (m *NameServiceConfigRequest) String() string { return proto.CompactTextString(m) }
func (*NameServiceConfigRequest) ProtoMessage()    {}
func (*NameServiceConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeffacffb8202a10, []int{0}
}
func (m *NameServiceConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServiceConfigRequest.Unmarshal(m, b)
}
func (m *NameServiceConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServiceConfigRequest.Marshal(b, m, deterministic)
}
func (m *NameServiceConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServiceConfigRequest.Merge(m, src)
}
func (m *NameServiceConfigRequest) XXX_Size() int {
	return xxx_messageInfo_NameServiceConfigRequest.Size(m)
}
func (m *NameServiceConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServiceConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameServiceConfigRequest proto.InternalMessageInfo

func (m *NameServiceConfigRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *NameServiceConfigRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//NameServiceConfig返回
type NameServiceConfigResponse struct {
	//
	//名字列表
	Data []*NameServiceConfigResponse_Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data" form:"data"`
	//
	//实例总数
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total"`
	//
	//页码
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//code
	Code int32 `protobuf:"varint,5,opt,name=code,proto3" json:"code" form:"code"`
	//
	//msg
	Msg                  string   `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg" form:"msg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameServiceConfigResponse) Reset()         { *m = NameServiceConfigResponse{} }
func (m *NameServiceConfigResponse) String() string { return proto.CompactTextString(m) }
func (*NameServiceConfigResponse) ProtoMessage()    {}
func (*NameServiceConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeffacffb8202a10, []int{1}
}
func (m *NameServiceConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServiceConfigResponse.Unmarshal(m, b)
}
func (m *NameServiceConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServiceConfigResponse.Marshal(b, m, deterministic)
}
func (m *NameServiceConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServiceConfigResponse.Merge(m, src)
}
func (m *NameServiceConfigResponse) XXX_Size() int {
	return xxx_messageInfo_NameServiceConfigResponse.Size(m)
}
func (m *NameServiceConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServiceConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameServiceConfigResponse proto.InternalMessageInfo

func (m *NameServiceConfigResponse) GetData() []*NameServiceConfigResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *NameServiceConfigResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *NameServiceConfigResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *NameServiceConfigResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *NameServiceConfigResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *NameServiceConfigResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type NameServiceConfigResponse_Data struct {
	//
	//服务名
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name" form:"service_name"`
	//
	//服务节点列表
	Hosts                []*NameServiceConfigResponse_Data_Hosts `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts" form:"hosts"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *NameServiceConfigResponse_Data) Reset()         { *m = NameServiceConfigResponse_Data{} }
func (m *NameServiceConfigResponse_Data) String() string { return proto.CompactTextString(m) }
func (*NameServiceConfigResponse_Data) ProtoMessage()    {}
func (*NameServiceConfigResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeffacffb8202a10, []int{1, 0}
}
func (m *NameServiceConfigResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServiceConfigResponse_Data.Unmarshal(m, b)
}
func (m *NameServiceConfigResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServiceConfigResponse_Data.Marshal(b, m, deterministic)
}
func (m *NameServiceConfigResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServiceConfigResponse_Data.Merge(m, src)
}
func (m *NameServiceConfigResponse_Data) XXX_Size() int {
	return xxx_messageInfo_NameServiceConfigResponse_Data.Size(m)
}
func (m *NameServiceConfigResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServiceConfigResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_NameServiceConfigResponse_Data proto.InternalMessageInfo

func (m *NameServiceConfigResponse_Data) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *NameServiceConfigResponse_Data) GetHosts() []*NameServiceConfigResponse_Data_Hosts {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type NameServiceConfigResponse_Data_Hosts struct {
	//
	//ip
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//tag
	Tag []string `protobuf:"bytes,2,rep,name=tag,proto3" json:"tag" form:"tag"`
	//
	//port
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port" form:"port"`
	//
	//weight
	Weight               int32    `protobuf:"varint,4,opt,name=weight,proto3" json:"weight" form:"weight"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameServiceConfigResponse_Data_Hosts) Reset()         { *m = NameServiceConfigResponse_Data_Hosts{} }
func (m *NameServiceConfigResponse_Data_Hosts) String() string { return proto.CompactTextString(m) }
func (*NameServiceConfigResponse_Data_Hosts) ProtoMessage()    {}
func (*NameServiceConfigResponse_Data_Hosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeffacffb8202a10, []int{1, 0, 0}
}
func (m *NameServiceConfigResponse_Data_Hosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServiceConfigResponse_Data_Hosts.Unmarshal(m, b)
}
func (m *NameServiceConfigResponse_Data_Hosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServiceConfigResponse_Data_Hosts.Marshal(b, m, deterministic)
}
func (m *NameServiceConfigResponse_Data_Hosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServiceConfigResponse_Data_Hosts.Merge(m, src)
}
func (m *NameServiceConfigResponse_Data_Hosts) XXX_Size() int {
	return xxx_messageInfo_NameServiceConfigResponse_Data_Hosts.Size(m)
}
func (m *NameServiceConfigResponse_Data_Hosts) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServiceConfigResponse_Data_Hosts.DiscardUnknown(m)
}

var xxx_messageInfo_NameServiceConfigResponse_Data_Hosts proto.InternalMessageInfo

func (m *NameServiceConfigResponse_Data_Hosts) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *NameServiceConfigResponse_Data_Hosts) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *NameServiceConfigResponse_Data_Hosts) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NameServiceConfigResponse_Data_Hosts) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

//
//NameServiceConfigApi返回
type NameServiceConfigResponseWrapper struct {
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
	Data                 *NameServiceConfigResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NameServiceConfigResponseWrapper) Reset()         { *m = NameServiceConfigResponseWrapper{} }
func (m *NameServiceConfigResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*NameServiceConfigResponseWrapper) ProtoMessage()    {}
func (*NameServiceConfigResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeffacffb8202a10, []int{2}
}
func (m *NameServiceConfigResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServiceConfigResponseWrapper.Unmarshal(m, b)
}
func (m *NameServiceConfigResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServiceConfigResponseWrapper.Marshal(b, m, deterministic)
}
func (m *NameServiceConfigResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServiceConfigResponseWrapper.Merge(m, src)
}
func (m *NameServiceConfigResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_NameServiceConfigResponseWrapper.Size(m)
}
func (m *NameServiceConfigResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServiceConfigResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_NameServiceConfigResponseWrapper proto.InternalMessageInfo

func (m *NameServiceConfigResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *NameServiceConfigResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *NameServiceConfigResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *NameServiceConfigResponseWrapper) GetData() *NameServiceConfigResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*NameServiceConfigRequest)(nil), "legacy.NameServiceConfigRequest")
	proto.RegisterType((*NameServiceConfigResponse)(nil), "legacy.NameServiceConfigResponse")
	proto.RegisterType((*NameServiceConfigResponse_Data)(nil), "legacy.NameServiceConfigResponse.Data")
	proto.RegisterType((*NameServiceConfigResponse_Data_Hosts)(nil), "legacy.NameServiceConfigResponse.Data.Hosts")
	proto.RegisterType((*NameServiceConfigResponseWrapper)(nil), "legacy.NameServiceConfigResponseWrapper")
}

func init() { proto.RegisterFile("nameservice_config.proto", fileDescriptor_eeffacffb8202a10) }

var fileDescriptor_eeffacffb8202a10 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x36, 0x69, 0x5a, 0xcc, 0x74, 0x57, 0xeb, 0x88, 0x1a, 0x0b, 0x92, 0x3a, 0xc2, 0xd2, 0x05,
	0x9b, 0xc2, 0x0a, 0x2a, 0x7b, 0x59, 0x7f, 0x10, 0x04, 0x2f, 0x66, 0x05, 0x2f, 0x97, 0x69, 0x3b,
	0x9d, 0x0e, 0x36, 0x9d, 0x98, 0x99, 0x6e, 0xd5, 0x0b, 0x9f, 0xc3, 0xf7, 0xf0, 0x7d, 0x22, 0xf8,
	0x06, 0x9b, 0x27, 0x90, 0x39, 0x93, 0x6d, 0x23, 0x58, 0xdd, 0xbd, 0x4a, 0xe6, 0x7c, 0xdf, 0x37,
	0x39, 0xdf, 0x77, 0x4e, 0x50, 0xb4, 0x64, 0x29, 0xd7, 0x3c, 0x3f, 0x93, 0x13, 0x7e, 0x3a, 0x51,
	0xcb, 0x99, 0x14, 0x49, 0x96, 0x2b, 0xa3, 0x70, 0x6b, 0xc1, 0x05, 0x9b, 0x7c, 0xe9, 0x0e, 0x84,
	0x34, 0xf3, 0xd5, 0x38, 0x99, 0xa8, 0x74, 0x28, 0x94, 0x50, 0x43, 0x80, 0xc7, 0xab, 0x19, 0x9c,
	0xe0, 0x00, 0x6f, 0x4e, 0xd6, 0x7d, 0x5a, 0xa3, 0xa7, 0x6b, 0x69, 0x3e, 0xaa, 0xf5, 0x50, 0xa8,
	0x01, 0x80, 0x83, 0x33, 0xb6, 0x90, 0x53, 0x66, 0x54, 0xae, 0x87, 0x9b, 0x57, 0xa7, 0x23, 0xdf,
	0x50, 0xf4, 0x8e, 0xa5, 0xfc, 0xc4, 0xb5, 0xf2, 0x02, 0x3a, 0xa1, 0xfc, 0xd3, 0x8a, 0x6b, 0x83,
	0x0f, 0x51, 0x90, 0x31, 0xc1, 0x23, 0xaf, 0xe7, 0xf5, 0x9b, 0xa3, 0x3b, 0x65, 0x11, 0xb7, 0x67,
	0x2a, 0x4f, 0x8f, 0x89, 0xad, 0x92, 0x5f, 0x3f, 0x63, 0xbf, 0x73, 0x8d, 0x02, 0x05, 0x3f, 0x43,
	0xa1, 0x7d, 0x9e, 0x6a, 0xf9, 0x95, 0x47, 0x3e, 0xf0, 0xbb, 0x65, 0x11, 0x77, 0xb6, 0x7c, 0x80,
	0x2e, 0x44, 0xd7, 0x6d, 0xe5, 0xc4, 0x16, 0xce, 0x03, 0x74, 0xff, 0x2f, 0x0d, 0xe8, 0x4c, 0x2d,
	0x35, 0xc7, 0x6f, 0x51, 0x30, 0x65, 0x86, 0x45, 0x5e, 0xaf, 0xd1, 0x6f, 0x1f, 0x1d, 0x24, 0x2e,
	0x9b, 0x64, 0xa7, 0x20, 0x79, 0xc9, 0x0c, 0x1b, 0xdd, 0xdc, 0x76, 0x6a, 0xd5, 0x84, 0xc2, 0x25,
	0xf8, 0x00, 0x35, 0x8d, 0x32, 0x6c, 0x51, 0xf5, 0xd7, 0x29, 0x8b, 0x78, 0xcf, 0xb1, 0xa0, 0x4c,
	0xa8, 0x83, 0x37, 0xb6, 0x1b, 0x57, 0xb4, 0x1d, 0x5c, 0xde, 0x36, 0x7e, 0x84, 0x82, 0x89, 0x9a,
	0xf2, 0xa8, 0x09, 0x9a, 0x5a, 0xc3, 0xb6, 0x4a, 0x28, 0x80, 0xb8, 0x87, 0x1a, 0xa9, 0x16, 0x51,
	0xab, 0xe7, 0xf5, 0xc3, 0xd1, 0x8d, 0xb2, 0x88, 0x91, 0xe3, 0xa4, 0x5a, 0x10, 0x6a, 0xa1, 0xee,
	0x0f, 0x1f, 0x05, 0xd6, 0x32, 0x3e, 0x46, 0x7b, 0x17, 0xdb, 0x64, 0x37, 0x0b, 0x46, 0x16, 0x8e,
	0xee, 0x95, 0x45, 0x7c, 0xdb, 0x69, 0xea, 0x28, 0xa1, 0xed, 0xea, 0x68, 0x83, 0xc4, 0xef, 0x51,
	0x73, 0xae, 0xb4, 0xd1, 0x91, 0x0f, 0x29, 0x3f, 0xbe, 0x5c, 0xca, 0xc9, 0x1b, 0xab, 0xa9, 0xa7,
	0x08, 0x97, 0x10, 0xea, 0x2e, 0xeb, 0x7e, 0xf7, 0x50, 0x13, 0x28, 0xf8, 0x01, 0xf2, 0x65, 0x56,
	0x75, 0xb4, 0x5f, 0x16, 0x71, 0xe8, 0xe8, 0x32, 0x23, 0xd4, 0x97, 0x99, 0x75, 0x69, 0x98, 0x80,
	0x8f, 0xff, 0xe1, 0xd2, 0x30, 0xeb, 0xd2, 0x30, 0x61, 0xc3, 0xca, 0x54, 0x6e, 0xaa, 0x81, 0xd4,
	0xc2, 0xb2, 0x55, 0x42, 0x01, 0xc4, 0x87, 0xa8, 0xb5, 0xe6, 0x52, 0xcc, 0x4d, 0x35, 0x87, 0x5b,
	0x65, 0x11, 0xef, 0x3b, 0x9a, 0xab, 0x13, 0x5a, 0x11, 0xc8, 0xb9, 0x87, 0x7a, 0x3b, 0xcd, 0x7d,
	0xc8, 0x59, 0x96, 0xf1, 0x7c, 0x33, 0x21, 0xef, 0x5f, 0x13, 0x7a, 0x8e, 0xda, 0xf6, 0xf9, 0xea,
	0x73, 0xb6, 0x60, 0x72, 0x09, 0x8b, 0x15, 0x8e, 0xee, 0x96, 0x45, 0x8c, 0xb7, 0xdc, 0x0a, 0x24,
	0xb4, 0x4e, 0xb5, 0xcb, 0xc8, 0xf3, 0x5c, 0xe5, 0x60, 0x2a, 0xac, 0xc7, 0x08, 0x65, 0x42, 0x1d,
	0x8c, 0x5f, 0x57, 0x7f, 0x80, 0x35, 0xd5, 0x3e, 0x7a, 0xf8, 0xdf, 0xd9, 0xec, 0x58, 0xfe, 0x71,
	0x0b, 0x7e, 0xf7, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x95, 0x43, 0x08, 0x79, 0x04,
	0x00, 0x00,
}
