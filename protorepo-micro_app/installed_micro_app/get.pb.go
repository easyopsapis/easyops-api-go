// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package installed_micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	micro_app "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//GetInstalledMicroApp请求
type GetInstalledMicroAppRequest struct {
	//
	//小产品id
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id" form:"app_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstalledMicroAppRequest) Reset()         { *m = GetInstalledMicroAppRequest{} }
func (m *GetInstalledMicroAppRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstalledMicroAppRequest) ProtoMessage()    {}
func (*GetInstalledMicroAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetInstalledMicroAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstalledMicroAppRequest.Unmarshal(m, b)
}
func (m *GetInstalledMicroAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstalledMicroAppRequest.Marshal(b, m, deterministic)
}
func (m *GetInstalledMicroAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstalledMicroAppRequest.Merge(m, src)
}
func (m *GetInstalledMicroAppRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstalledMicroAppRequest.Size(m)
}
func (m *GetInstalledMicroAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstalledMicroAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstalledMicroAppRequest proto.InternalMessageInfo

func (m *GetInstalledMicroAppRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

//
//GetInstalledMicroAppApi返回
type GetInstalledMicroAppResponseWrapper struct {
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
	Data                 *micro_app.InstalledMicroApp `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetInstalledMicroAppResponseWrapper) Reset()         { *m = GetInstalledMicroAppResponseWrapper{} }
func (m *GetInstalledMicroAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetInstalledMicroAppResponseWrapper) ProtoMessage()    {}
func (*GetInstalledMicroAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetInstalledMicroAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstalledMicroAppResponseWrapper.Unmarshal(m, b)
}
func (m *GetInstalledMicroAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstalledMicroAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetInstalledMicroAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstalledMicroAppResponseWrapper.Merge(m, src)
}
func (m *GetInstalledMicroAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetInstalledMicroAppResponseWrapper.Size(m)
}
func (m *GetInstalledMicroAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstalledMicroAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstalledMicroAppResponseWrapper proto.InternalMessageInfo

func (m *GetInstalledMicroAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetInstalledMicroAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetInstalledMicroAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetInstalledMicroAppResponseWrapper) GetData() *micro_app.InstalledMicroApp {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetInstalledMicroAppRequest)(nil), "installed_micro_app.GetInstalledMicroAppRequest")
	proto.RegisterType((*GetInstalledMicroAppResponseWrapper)(nil), "installed_micro_app.GetInstalledMicroAppResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4f, 0x6b, 0xe2, 0x40,
	0x14, 0x27, 0xbb, 0x2a, 0x38, 0xee, 0xb2, 0xbb, 0x11, 0x96, 0xe0, 0x2e, 0x44, 0xc6, 0x65, 0xf1,
	0xb0, 0xc9, 0xb0, 0x0a, 0xa5, 0xf4, 0x52, 0x14, 0x4a, 0xf1, 0xd0, 0x4b, 0x2e, 0x85, 0x16, 0x0d,
	0x63, 0x32, 0xa6, 0xa1, 0x49, 0xde, 0x74, 0x66, 0xac, 0xfd, 0x43, 0xbf, 0x6a, 0x0a, 0xfd, 0x02,
	0x85, 0x7c, 0x82, 0x92, 0x89, 0x55, 0xa1, 0x9e, 0x32, 0xef, 0xfd, 0xfe, 0xe6, 0xa1, 0x66, 0xc4,
	0x94, 0xcb, 0x05, 0x28, 0x30, 0xdb, 0x71, 0x26, 0x15, 0x4d, 0x12, 0x16, 0xfa, 0x69, 0x1c, 0x08,
	0xf0, 0x29, 0xe7, 0x1d, 0x27, 0x8a, 0xd5, 0xd5, 0x72, 0xee, 0x06, 0x90, 0x92, 0x08, 0x22, 0x20,
	0x9a, 0x3b, 0x5f, 0x2e, 0xf4, 0xa4, 0x07, 0xfd, 0xaa, 0x3c, 0x3a, 0xd3, 0x08, 0x5c, 0x46, 0xe5,
	0x3d, 0x70, 0xe9, 0x26, 0x10, 0xd0, 0x84, 0x04, 0x90, 0x29, 0x41, 0x03, 0x25, 0x2b, 0xa5, 0x60,
	0x1c, 0x9c, 0x14, 0x42, 0x96, 0x48, 0xb2, 0x26, 0x12, 0x3d, 0x92, 0x4d, 0x26, 0xd9, 0xd3, 0x63,
	0x6d, 0x7f, 0xb0, 0xd3, 0x26, 0x5d, 0xc5, 0xea, 0x1a, 0x56, 0x24, 0x02, 0x47, 0x83, 0xce, 0x2d,
	0x4d, 0xe2, 0x90, 0x2a, 0x10, 0x92, 0x6c, 0x9e, 0x95, 0x0e, 0xcf, 0xd0, 0xaf, 0x53, 0xa6, 0x26,
	0xef, 0xbe, 0x67, 0xa5, 0xed, 0x88, 0x73, 0x8f, 0xdd, 0x2c, 0x99, 0x54, 0xe6, 0x31, 0x6a, 0x50,
	0xce, 0xfd, 0x38, 0xb4, 0x8c, 0xae, 0xd1, 0x6f, 0x8e, 0xfb, 0x45, 0x6e, 0x7f, 0x5d, 0x80, 0x48,
	0x8f, 0x70, 0xb5, 0xc7, 0x2f, 0xcf, 0x76, 0x1b, 0xfd, 0x98, 0x5d, 0x52, 0xe7, 0x61, 0xe4, 0x5c,
	0xf8, 0xd3, 0xc7, 0xff, 0xff, 0x86, 0x83, 0xa7, 0x3f, 0x5e, 0x9d, 0x72, 0x3e, 0x09, 0xf1, 0xab,
	0x81, 0x7a, 0xfb, 0x03, 0x24, 0x87, 0x4c, 0xb2, 0x73, 0x41, 0x39, 0x67, 0xc2, 0xec, 0xa1, 0x5a,
	0x00, 0x21, 0xd3, 0x31, 0xf5, 0xf1, 0xb7, 0x22, 0xb7, 0x5b, 0x55, 0x4c, 0xb9, 0xc5, 0x9e, 0x06,
	0xcd, 0x43, 0xd4, 0x2a, 0xbf, 0x27, 0x77, 0x3c, 0xa1, 0x71, 0x66, 0x7d, 0xd2, 0x95, 0x7e, 0x16,
	0xb9, 0x6d, 0x6e, 0xb9, 0x6b, 0x10, 0x7b, 0xbb, 0x54, 0xf3, 0x2f, 0xaa, 0x33, 0x21, 0x40, 0x58,
	0x9f, 0xb5, 0xe6, 0x7b, 0x91, 0xdb, 0x5f, 0x2a, 0x8d, 0x5e, 0x63, 0xaf, 0x82, 0xcd, 0x11, 0xaa,
	0x85, 0x54, 0x51, 0xab, 0xd6, 0x35, 0xfa, 0xad, 0xc1, 0x6f, 0x77, 0x7b, 0xe6, 0x0f, 0x7f, 0xb0,
	0x5b, 0xb2, 0xd4, 0x60, 0x4f, 0x4b, 0xe7, 0x0d, 0x7d, 0xd8, 0xe1, 0x5b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x00, 0xf8, 0x3d, 0x40, 0x02, 0x00, 0x00,
}