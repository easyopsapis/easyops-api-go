// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_tool_global_config.proto

package basic

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
//GetToolGlobalConfig返回
type GetToolGlobalConfigResponse struct {
	//
	//系统ORG
	Org int32 `protobuf:"varint,1,opt,name=org,proto3" json:"org" form:"org"`
	//
	//配置类型
	ConfigType string `protobuf:"bytes,2,opt,name=ConfigType,proto3" json:"ConfigType" form:"ConfigType"`
	//
	//审批并同意
	CheckAndAgree string `protobuf:"bytes,3,opt,name=CheckAndAgree,proto3" json:"CheckAndAgree" form:"CheckAndAgree"`
	//
	//审批人列表
	CheckManagers        []string `protobuf:"bytes,4,rep,name=CheckManagers,proto3" json:"CheckManagers" form:"CheckManagers"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetToolGlobalConfigResponse) Reset()         { *m = GetToolGlobalConfigResponse{} }
func (m *GetToolGlobalConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetToolGlobalConfigResponse) ProtoMessage()    {}
func (*GetToolGlobalConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b31ee3dd6aacea8, []int{0}
}
func (m *GetToolGlobalConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetToolGlobalConfigResponse.Unmarshal(m, b)
}
func (m *GetToolGlobalConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetToolGlobalConfigResponse.Marshal(b, m, deterministic)
}
func (m *GetToolGlobalConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetToolGlobalConfigResponse.Merge(m, src)
}
func (m *GetToolGlobalConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetToolGlobalConfigResponse.Size(m)
}
func (m *GetToolGlobalConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetToolGlobalConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetToolGlobalConfigResponse proto.InternalMessageInfo

func (m *GetToolGlobalConfigResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetToolGlobalConfigResponse) GetConfigType() string {
	if m != nil {
		return m.ConfigType
	}
	return ""
}

func (m *GetToolGlobalConfigResponse) GetCheckAndAgree() string {
	if m != nil {
		return m.CheckAndAgree
	}
	return ""
}

func (m *GetToolGlobalConfigResponse) GetCheckManagers() []string {
	if m != nil {
		return m.CheckManagers
	}
	return nil
}

//
//GetToolGlobalConfigApi返回
type GetToolGlobalConfigResponseWrapper struct {
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
	Data                 *GetToolGlobalConfigResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetToolGlobalConfigResponseWrapper) Reset()         { *m = GetToolGlobalConfigResponseWrapper{} }
func (m *GetToolGlobalConfigResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetToolGlobalConfigResponseWrapper) ProtoMessage()    {}
func (*GetToolGlobalConfigResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b31ee3dd6aacea8, []int{1}
}
func (m *GetToolGlobalConfigResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetToolGlobalConfigResponseWrapper.Unmarshal(m, b)
}
func (m *GetToolGlobalConfigResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetToolGlobalConfigResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetToolGlobalConfigResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetToolGlobalConfigResponseWrapper.Merge(m, src)
}
func (m *GetToolGlobalConfigResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetToolGlobalConfigResponseWrapper.Size(m)
}
func (m *GetToolGlobalConfigResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetToolGlobalConfigResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetToolGlobalConfigResponseWrapper proto.InternalMessageInfo

func (m *GetToolGlobalConfigResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetToolGlobalConfigResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetToolGlobalConfigResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetToolGlobalConfigResponseWrapper) GetData() *GetToolGlobalConfigResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetToolGlobalConfigResponse)(nil), "basic.GetToolGlobalConfigResponse")
	proto.RegisterType((*GetToolGlobalConfigResponseWrapper)(nil), "basic.GetToolGlobalConfigResponseWrapper")
}

func init() { proto.RegisterFile("get_tool_global_config.proto", fileDescriptor_2b31ee3dd6aacea8) }

var fileDescriptor_2b31ee3dd6aacea8 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xe7, 0xc4, 0x19, 0x4c, 0xd9, 0x58, 0x26, 0xb2, 0x61, 0xb2, 0x81, 0x8d, 0x06, 0x21,
	0x97, 0xd8, 0xb0, 0xb1, 0x51, 0x7a, 0x28, 0x24, 0xa1, 0xe4, 0xd4, 0x8b, 0x08, 0xf4, 0x18, 0x64,
	0x5b, 0x51, 0x4c, 0x1c, 0x3f, 0x23, 0x2b, 0x4d, 0xfb, 0x65, 0x5d, 0xe8, 0x47, 0x30, 0x3d, 0xf6,
	0x50, 0x2c, 0x3b, 0x89, 0x0b, 0x25, 0x27, 0xeb, 0xe9, 0xf7, 0xff, 0x89, 0xa7, 0x67, 0xa1, 0x5f,
	0x82, 0xab, 0xa5, 0x02, 0x88, 0x97, 0x22, 0x06, 0x9f, 0xc5, 0xcb, 0x00, 0x92, 0x55, 0x24, 0xdc,
	0x54, 0x82, 0x02, 0xdc, 0xf1, 0x59, 0x16, 0x05, 0x83, 0xb1, 0x88, 0xd4, 0x7a, 0xe7, 0xbb, 0x01,
	0x6c, 0x3d, 0x01, 0x02, 0x3c, 0x4d, 0xfd, 0xdd, 0x4a, 0x57, 0xba, 0xd0, 0xab, 0xca, 0x1a, 0xfc,
	0x6f, 0xc4, 0xb7, 0xfb, 0x48, 0x6d, 0x60, 0xef, 0x09, 0x18, 0x6b, 0x38, 0xbe, 0x63, 0x71, 0x14,
	0x32, 0x05, 0x32, 0xf3, 0x8e, 0xcb, 0xca, 0x23, 0x2f, 0x06, 0xfa, 0x39, 0xe7, 0x6a, 0x01, 0x10,
	0xcf, 0x75, 0x33, 0x33, 0xdd, 0x0b, 0xe5, 0x59, 0x0a, 0x49, 0xc6, 0xf1, 0x10, 0xb5, 0x41, 0x0a,
	0xcb, 0x70, 0x8c, 0x51, 0x67, 0xda, 0x2f, 0x72, 0x1b, 0xad, 0x40, 0x6e, 0x2f, 0x09, 0x48, 0x41,
	0x9e, 0x1e, 0xed, 0x56, 0xef, 0x03, 0x2d, 0x03, 0xf8, 0x1f, 0x42, 0x95, 0xb9, 0x78, 0x48, 0xb9,
	0xd5, 0x72, 0x8c, 0xd1, 0xa7, 0xe9, 0xf7, 0x22, 0xb7, 0xbf, 0x55, 0xf1, 0x13, 0x23, 0xb4, 0x11,
	0xc4, 0x57, 0xe8, 0xcb, 0x6c, 0xcd, 0x83, 0xcd, 0x24, 0x09, 0x27, 0x42, 0x72, 0x6e, 0xb5, 0xb5,
	0x69, 0x15, 0xb9, 0xdd, 0xaf, 0xcd, 0x26, 0x26, 0xf4, 0x6d, 0xfc, 0xe8, 0xdf, 0xb0, 0x84, 0x09,
	0x2e, 0x33, 0xcb, 0x74, 0xda, 0xef, 0xf8, 0x07, 0x7c, 0xf0, 0x8f, 0xf5, 0xb3, 0x81, 0xc8, 0x99,
	0xeb, 0xdf, 0x4a, 0x96, 0xa6, 0x5c, 0xe2, 0xdf, 0xc8, 0x0c, 0x20, 0xe4, 0xf5, 0x18, 0xbe, 0x16,
	0xb9, 0xdd, 0xad, 0x4e, 0x2f, 0x77, 0x09, 0xd5, 0x10, 0x5f, 0xa0, 0x6e, 0xf9, 0xbd, 0xbe, 0x4f,
	0x63, 0x16, 0x25, 0xf5, 0x0c, 0x7e, 0x14, 0xb9, 0x8d, 0x4f, 0xd9, 0x1a, 0x12, 0xda, 0x8c, 0xe2,
	0x21, 0xea, 0x70, 0x29, 0x41, 0xd6, 0xb7, 0xef, 0x15, 0xb9, 0xfd, 0xb9, 0x72, 0xf4, 0x36, 0xa1,
	0x15, 0xc6, 0x73, 0x64, 0x86, 0x4c, 0x31, 0xcb, 0x74, 0x8c, 0x51, 0xf7, 0x0f, 0x71, 0xf5, 0x4b,
	0x71, 0xcf, 0xf4, 0xdf, 0x6c, 0xb5, 0x34, 0x09, 0xd5, 0x07, 0xf8, 0x1f, 0xf5, 0xcf, 0xff, 0xfb,
	0x1a, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x32, 0xb0, 0x12, 0x8a, 0x02, 0x00, 0x00,
}
