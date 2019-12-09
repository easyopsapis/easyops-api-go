// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: install_app.proto

package desktop

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
//InstallApp请求
type InstallAppRequest struct {
	//
	//小产品id
	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//小产品版本号
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version" form:"version"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallAppRequest) Reset()         { *m = InstallAppRequest{} }
func (m *InstallAppRequest) String() string { return proto.CompactTextString(m) }
func (*InstallAppRequest) ProtoMessage()    {}
func (*InstallAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb64c1c8f7ea7714, []int{0}
}
func (m *InstallAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallAppRequest.Unmarshal(m, b)
}
func (m *InstallAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallAppRequest.Marshal(b, m, deterministic)
}
func (m *InstallAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallAppRequest.Merge(m, src)
}
func (m *InstallAppRequest) XXX_Size() int {
	return xxx_messageInfo_InstallAppRequest.Size(m)
}
func (m *InstallAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstallAppRequest proto.InternalMessageInfo

func (m *InstallAppRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *InstallAppRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

//
//InstallApp返回
type InstallAppResponse struct {
	//
	//任务id
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallAppResponse) Reset()         { *m = InstallAppResponse{} }
func (m *InstallAppResponse) String() string { return proto.CompactTextString(m) }
func (*InstallAppResponse) ProtoMessage()    {}
func (*InstallAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb64c1c8f7ea7714, []int{1}
}
func (m *InstallAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallAppResponse.Unmarshal(m, b)
}
func (m *InstallAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallAppResponse.Marshal(b, m, deterministic)
}
func (m *InstallAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallAppResponse.Merge(m, src)
}
func (m *InstallAppResponse) XXX_Size() int {
	return xxx_messageInfo_InstallAppResponse.Size(m)
}
func (m *InstallAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstallAppResponse proto.InternalMessageInfo

func (m *InstallAppResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

//
//InstallAppApi返回
type InstallAppResponseWrapper struct {
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
	Data                 *InstallAppResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InstallAppResponseWrapper) Reset()         { *m = InstallAppResponseWrapper{} }
func (m *InstallAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*InstallAppResponseWrapper) ProtoMessage()    {}
func (*InstallAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb64c1c8f7ea7714, []int{2}
}
func (m *InstallAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallAppResponseWrapper.Unmarshal(m, b)
}
func (m *InstallAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *InstallAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallAppResponseWrapper.Merge(m, src)
}
func (m *InstallAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_InstallAppResponseWrapper.Size(m)
}
func (m *InstallAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_InstallAppResponseWrapper proto.InternalMessageInfo

func (m *InstallAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InstallAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *InstallAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *InstallAppResponseWrapper) GetData() *InstallAppResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallAppRequest)(nil), "desktop.InstallAppRequest")
	proto.RegisterType((*InstallAppResponse)(nil), "desktop.InstallAppResponse")
	proto.RegisterType((*InstallAppResponseWrapper)(nil), "desktop.InstallAppResponseWrapper")
}

func init() { proto.RegisterFile("install_app.proto", fileDescriptor_bb64c1c8f7ea7714) }

var fileDescriptor_bb64c1c8f7ea7714 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4d, 0x4f, 0xdb, 0x30,
	0x18, 0x56, 0xb6, 0x7e, 0x68, 0xee, 0xb6, 0x2e, 0x3e, 0x4c, 0x59, 0x77, 0x48, 0x65, 0x10, 0x14,
	0xa1, 0x26, 0x12, 0x48, 0x08, 0x71, 0x81, 0x56, 0xe2, 0xd0, 0xab, 0x2f, 0x5c, 0x2a, 0x90, 0xdb,
	0xb8, 0x21, 0x6a, 0xda, 0xd7, 0xd8, 0x6e, 0xcb, 0x89, 0x9f, 0x1a, 0x24, 0xc4, 0x2f, 0xc8, 0x2f,
	0x40, 0xb1, 0xd3, 0x12, 0xd1, 0x4b, 0x62, 0x3f, 0x5f, 0x79, 0xf2, 0xbe, 0xc8, 0x4d, 0x96, 0x4a,
	0xb3, 0x34, 0x7d, 0x60, 0x42, 0x04, 0x42, 0x82, 0x06, 0xdc, 0x8c, 0xb8, 0x9a, 0x6b, 0x10, 0x9d,
	0x7e, 0x9c, 0xe8, 0xc7, 0xd5, 0x24, 0x98, 0xc2, 0x22, 0x8c, 0x21, 0x86, 0xd0, 0xf0, 0x93, 0xd5,
	0xcc, 0xdc, 0xcc, 0xc5, 0x9c, 0xac, 0xaf, 0x73, 0x51, 0x91, 0x2f, 0x36, 0x89, 0x9e, 0xc3, 0x26,
	0x8c, 0xa1, 0x6f, 0xc8, 0xfe, 0x9a, 0xa5, 0x49, 0xc4, 0x34, 0x48, 0x15, 0xee, 0x8e, 0xd6, 0x47,
	0x5e, 0x90, 0x3b, 0xb2, 0x25, 0x06, 0x42, 0x50, 0xfe, 0xb4, 0xe2, 0x4a, 0xe3, 0x23, 0x54, 0x67,
	0x42, 0x8c, 0x22, 0xcf, 0xe9, 0x3a, 0xbd, 0x1f, 0xc3, 0x3f, 0x79, 0xe6, 0xff, 0x9c, 0x81, 0x5c,
	0x5c, 0x11, 0x03, 0x13, 0x6a, 0x69, 0x3c, 0x40, 0xcd, 0x35, 0x97, 0x2a, 0x81, 0xa5, 0xf7, 0xcd,
	0x28, 0x8f, 0xf3, 0xcc, 0xff, 0x6d, 0x95, 0x25, 0x41, 0xde, 0x5e, 0x7d, 0x17, 0xb5, 0xef, 0xc7,
	0xd1, 0xe9, 0x38, 0xd8, 0x3e, 0x0e, 0xe9, 0xd6, 0x47, 0xae, 0x11, 0xae, 0x7e, 0x5f, 0x09, 0x58,
	0x2a, 0x8e, 0x4f, 0x50, 0x43, 0x33, 0x35, 0xdf, 0x35, 0x70, 0xf3, 0xcc, 0xff, 0x65, 0x73, 0x2d,
	0x4e, 0x68, 0x29, 0x20, 0xef, 0x0e, 0xfa, 0xb7, 0x9f, 0x70, 0x27, 0x99, 0x10, 0x5c, 0xe2, 0x03,
	0x54, 0x9b, 0x42, 0xc4, 0x4d, 0x4c, 0x7d, 0xd8, 0xce, 0x33, 0xbf, 0x65, 0x63, 0x0a, 0x94, 0x50,
	0x43, 0xe2, 0x4b, 0xd4, 0x2a, 0xde, 0xb7, 0xcf, 0x22, 0x65, 0xc9, 0xf6, 0x57, 0xfe, 0xe6, 0x99,
	0x8f, 0x3f, 0xb5, 0x25, 0x49, 0x68, 0x55, 0x5a, 0x0c, 0x8a, 0x4b, 0x09, 0xd2, 0xfb, 0xfe, 0x75,
	0x50, 0x06, 0x26, 0xd4, 0xd2, 0xf8, 0x06, 0xd5, 0x22, 0xa6, 0x99, 0x57, 0xeb, 0x3a, 0xbd, 0xd6,
	0xd9, 0xff, 0xa0, 0x5c, 0x72, 0xb0, 0x5f, 0xbc, 0xda, 0xb1, 0xb0, 0x10, 0x6a, 0x9c, 0x93, 0x86,
	0x59, 0xd7, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0xcd, 0xbc, 0xc1, 0x33, 0x02, 0x00,
	0x00,
}
