// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_plugin_version.proto

package plugin_version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//CreatePluginVersion请求
type CreatePluginVersionRequest struct {
	//
	//插件Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//版本名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo                 string   `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePluginVersionRequest) Reset()         { *m = CreatePluginVersionRequest{} }
func (m *CreatePluginVersionRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePluginVersionRequest) ProtoMessage()    {}
func (*CreatePluginVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3599c9a676583d, []int{0}
}
func (m *CreatePluginVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePluginVersionRequest.Unmarshal(m, b)
}
func (m *CreatePluginVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePluginVersionRequest.Marshal(b, m, deterministic)
}
func (m *CreatePluginVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePluginVersionRequest.Merge(m, src)
}
func (m *CreatePluginVersionRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePluginVersionRequest.Size(m)
}
func (m *CreatePluginVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePluginVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePluginVersionRequest proto.InternalMessageInfo

func (m *CreatePluginVersionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreatePluginVersionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePluginVersionRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

//
//CreatePluginVersion返回
type CreatePluginVersionResponse struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePluginVersionResponse) Reset()         { *m = CreatePluginVersionResponse{} }
func (m *CreatePluginVersionResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePluginVersionResponse) ProtoMessage()    {}
func (*CreatePluginVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3599c9a676583d, []int{1}
}
func (m *CreatePluginVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePluginVersionResponse.Unmarshal(m, b)
}
func (m *CreatePluginVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePluginVersionResponse.Marshal(b, m, deterministic)
}
func (m *CreatePluginVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePluginVersionResponse.Merge(m, src)
}
func (m *CreatePluginVersionResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePluginVersionResponse.Size(m)
}
func (m *CreatePluginVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePluginVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePluginVersionResponse proto.InternalMessageInfo

func (m *CreatePluginVersionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//CreatePluginVersionApi返回
type CreatePluginVersionResponseWrapper struct {
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
	Data                 *CreatePluginVersionResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreatePluginVersionResponseWrapper) Reset()         { *m = CreatePluginVersionResponseWrapper{} }
func (m *CreatePluginVersionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreatePluginVersionResponseWrapper) ProtoMessage()    {}
func (*CreatePluginVersionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3599c9a676583d, []int{2}
}
func (m *CreatePluginVersionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePluginVersionResponseWrapper.Unmarshal(m, b)
}
func (m *CreatePluginVersionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePluginVersionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreatePluginVersionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePluginVersionResponseWrapper.Merge(m, src)
}
func (m *CreatePluginVersionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreatePluginVersionResponseWrapper.Size(m)
}
func (m *CreatePluginVersionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePluginVersionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePluginVersionResponseWrapper proto.InternalMessageInfo

func (m *CreatePluginVersionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreatePluginVersionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreatePluginVersionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreatePluginVersionResponseWrapper) GetData() *CreatePluginVersionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePluginVersionRequest)(nil), "plugin_version.CreatePluginVersionRequest")
	proto.RegisterType((*CreatePluginVersionResponse)(nil), "plugin_version.CreatePluginVersionResponse")
	proto.RegisterType((*CreatePluginVersionResponseWrapper)(nil), "plugin_version.CreatePluginVersionResponseWrapper")
}

func init() { proto.RegisterFile("create_plugin_version.proto", fileDescriptor_9e3599c9a676583d) }

var fileDescriptor_9e3599c9a676583d = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x4c, 0x85, 0x6e, 0xfc, 0xc7, 0x1e, 0x24, 0xb4, 0x48, 0xca, 0x08, 0x52, 0x10,
	0x53, 0xd0, 0x8b, 0x88, 0x27, 0xc5, 0x7b, 0xd9, 0x83, 0x1e, 0xcb, 0x26, 0xd9, 0xc6, 0x85, 0x26,
	0xbb, 0x6e, 0x12, 0xf1, 0x0b, 0xf8, 0x35, 0xf3, 0x21, 0x72, 0x17, 0x64, 0x27, 0x01, 0x63, 0x91,
	0x7a, 0x4a, 0xf6, 0xbd, 0xdf, 0x3c, 0xde, 0x30, 0x64, 0x9a, 0x18, 0xc1, 0x2b, 0xb1, 0xd2, 0x9b,
	0x3a, 0x93, 0xc5, 0xea, 0x5d, 0x98, 0x52, 0xaa, 0x22, 0xd2, 0x46, 0x55, 0x8a, 0x1e, 0xfd, 0x56,
	0x27, 0x57, 0x99, 0xac, 0x5e, 0xeb, 0x38, 0x4a, 0x54, 0xbe, 0xc8, 0x54, 0xa6, 0x16, 0x88, 0xc5,
	0xf5, 0x1a, 0x5f, 0xf8, 0xc0, 0xbf, 0x6e, 0x1c, 0x3e, 0x1d, 0x32, 0x79, 0xc4, 0xf8, 0x25, 0xe6,
	0x3c, 0x77, 0x31, 0x4c, 0xbc, 0xd5, 0xa2, 0xac, 0xe8, 0x19, 0x71, 0x65, 0x1a, 0x38, 0x33, 0x67,
	0x3e, 0x7e, 0x38, 0x6c, 0x9b, 0x70, 0xbc, 0x56, 0x26, 0xbf, 0x03, 0x99, 0x02, 0x73, 0x65, 0x4a,
	0xcf, 0x89, 0x57, 0xf0, 0x5c, 0x04, 0x2e, 0x02, 0xc7, 0x6d, 0x13, 0xfa, 0x1d, 0x60, 0x55, 0x60,
	0x68, 0x5a, 0x28, 0x17, 0xb9, 0x0a, 0xf6, 0xb6, 0x21, 0xab, 0x02, 0x43, 0x13, 0xee, 0xc9, 0xf4,
	0xcf, 0x1a, 0xa5, 0x56, 0x45, 0x29, 0xfe, 0xe9, 0x01, 0x5f, 0x0e, 0x81, 0x1d, 0xe3, 0x2f, 0x86,
	0x6b, 0x2d, 0x8c, 0x6d, 0x92, 0xa8, 0x54, 0x60, 0xce, 0x68, 0xd8, 0xc4, 0xaa, 0xc0, 0xd0, 0xa4,
	0xb7, 0xc4, 0xb7, 0xdf, 0xa7, 0x0f, 0xbd, 0xe1, 0xb2, 0xe8, 0x57, 0x3b, 0x6d, 0x9b, 0x90, 0xfe,
	0xb0, 0xbd, 0x09, 0x6c, 0x88, 0xd2, 0x0b, 0x32, 0x12, 0xc6, 0x28, 0xd3, 0x6f, 0x7a, 0xd2, 0x36,
	0xe1, 0x41, 0x37, 0x83, 0x32, 0xb0, 0xce, 0xa6, 0x4b, 0xe2, 0xa5, 0xbc, 0xe2, 0x81, 0x37, 0x73,
	0xe6, 0xfe, 0xf5, 0x65, 0xb4, 0x75, 0xd7, 0x1d, 0x8b, 0x0c, 0x3b, 0xdb, 0x08, 0x60, 0x98, 0x14,
	0xef, 0xe3, 0x31, 0x6f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7b, 0xbc, 0x5e, 0x2a, 0x02,
	0x00, 0x00,
}
