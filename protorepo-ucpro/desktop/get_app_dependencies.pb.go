// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_app_dependencies.proto

package desktop

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
//GetAppDependencies请求
type GetAppDependenciesRequest struct {
	//
	//小产品id
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppDependenciesRequest) Reset()         { *m = GetAppDependenciesRequest{} }
func (m *GetAppDependenciesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppDependenciesRequest) ProtoMessage()    {}
func (*GetAppDependenciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b25f48986be4092, []int{0}
}
func (m *GetAppDependenciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppDependenciesRequest.Unmarshal(m, b)
}
func (m *GetAppDependenciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppDependenciesRequest.Marshal(b, m, deterministic)
}
func (m *GetAppDependenciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppDependenciesRequest.Merge(m, src)
}
func (m *GetAppDependenciesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppDependenciesRequest.Size(m)
}
func (m *GetAppDependenciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppDependenciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppDependenciesRequest proto.InternalMessageInfo

func (m *GetAppDependenciesRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

//
//GetAppDependencies返回
type GetAppDependenciesResponse struct {
	//
	//依赖信息列表
	Dependencies         []*GetAppDependenciesResponse_Dependencies `protobuf:"bytes,1,rep,name=dependencies,proto3" json:"dependencies" form:"dependencies"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *GetAppDependenciesResponse) Reset()         { *m = GetAppDependenciesResponse{} }
func (m *GetAppDependenciesResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppDependenciesResponse) ProtoMessage()    {}
func (*GetAppDependenciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b25f48986be4092, []int{1}
}
func (m *GetAppDependenciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppDependenciesResponse.Unmarshal(m, b)
}
func (m *GetAppDependenciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppDependenciesResponse.Marshal(b, m, deterministic)
}
func (m *GetAppDependenciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppDependenciesResponse.Merge(m, src)
}
func (m *GetAppDependenciesResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppDependenciesResponse.Size(m)
}
func (m *GetAppDependenciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppDependenciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppDependenciesResponse proto.InternalMessageInfo

func (m *GetAppDependenciesResponse) GetDependencies() []*GetAppDependenciesResponse_Dependencies {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type GetAppDependenciesResponse_Dependencies struct {
	//
	//依赖组件
	PackageName string `protobuf:"bytes,1,opt,name=packageName,proto3" json:"packageName" form:"packageName"`
	//
	//依赖版本
	Constraint string `protobuf:"bytes,2,opt,name=constraint,proto3" json:"constraint" form:"constraint"`
	//
	//当前版本
	CurrentVersion string `protobuf:"bytes,3,opt,name=currentVersion,proto3" json:"currentVersion" form:"currentVersion"`
	//
	//当前版本是否符合依赖
	MeetConstraint       bool     `protobuf:"varint,4,opt,name=meetConstraint,proto3" json:"meetConstraint" form:"meetConstraint"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppDependenciesResponse_Dependencies) Reset() {
	*m = GetAppDependenciesResponse_Dependencies{}
}
func (m *GetAppDependenciesResponse_Dependencies) String() string { return proto.CompactTextString(m) }
func (*GetAppDependenciesResponse_Dependencies) ProtoMessage()    {}
func (*GetAppDependenciesResponse_Dependencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b25f48986be4092, []int{1, 0}
}
func (m *GetAppDependenciesResponse_Dependencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppDependenciesResponse_Dependencies.Unmarshal(m, b)
}
func (m *GetAppDependenciesResponse_Dependencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppDependenciesResponse_Dependencies.Marshal(b, m, deterministic)
}
func (m *GetAppDependenciesResponse_Dependencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppDependenciesResponse_Dependencies.Merge(m, src)
}
func (m *GetAppDependenciesResponse_Dependencies) XXX_Size() int {
	return xxx_messageInfo_GetAppDependenciesResponse_Dependencies.Size(m)
}
func (m *GetAppDependenciesResponse_Dependencies) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppDependenciesResponse_Dependencies.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppDependenciesResponse_Dependencies proto.InternalMessageInfo

func (m *GetAppDependenciesResponse_Dependencies) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *GetAppDependenciesResponse_Dependencies) GetConstraint() string {
	if m != nil {
		return m.Constraint
	}
	return ""
}

func (m *GetAppDependenciesResponse_Dependencies) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *GetAppDependenciesResponse_Dependencies) GetMeetConstraint() bool {
	if m != nil {
		return m.MeetConstraint
	}
	return false
}

//
//GetAppDependenciesApi返回
type GetAppDependenciesResponseWrapper struct {
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
	Data                 *GetAppDependenciesResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetAppDependenciesResponseWrapper) Reset()         { *m = GetAppDependenciesResponseWrapper{} }
func (m *GetAppDependenciesResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetAppDependenciesResponseWrapper) ProtoMessage()    {}
func (*GetAppDependenciesResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b25f48986be4092, []int{2}
}
func (m *GetAppDependenciesResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppDependenciesResponseWrapper.Unmarshal(m, b)
}
func (m *GetAppDependenciesResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppDependenciesResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetAppDependenciesResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppDependenciesResponseWrapper.Merge(m, src)
}
func (m *GetAppDependenciesResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetAppDependenciesResponseWrapper.Size(m)
}
func (m *GetAppDependenciesResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppDependenciesResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppDependenciesResponseWrapper proto.InternalMessageInfo

func (m *GetAppDependenciesResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAppDependenciesResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAppDependenciesResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetAppDependenciesResponseWrapper) GetData() *GetAppDependenciesResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAppDependenciesRequest)(nil), "desktop.GetAppDependenciesRequest")
	proto.RegisterType((*GetAppDependenciesResponse)(nil), "desktop.GetAppDependenciesResponse")
	proto.RegisterType((*GetAppDependenciesResponse_Dependencies)(nil), "desktop.GetAppDependenciesResponse.Dependencies")
	proto.RegisterType((*GetAppDependenciesResponseWrapper)(nil), "desktop.GetAppDependenciesResponseWrapper")
}

func init() { proto.RegisterFile("get_app_dependencies.proto", fileDescriptor_2b25f48986be4092) }

var fileDescriptor_2b25f48986be4092 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4b, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x21, 0x3f, 0xfa, 0xa0, 0x8c, 0x3e, 0x58, 0xb8, 0x95, 0xb5, 0x91, 0x4b, 0x03, 0x85,
	0x37, 0x95, 0x0b, 0x17, 0x05, 0x8a, 0xee, 0x6c, 0xb7, 0x68, 0xbb, 0xc9, 0x82, 0x8b, 0x64, 0x69,
	0xd0, 0xd2, 0x58, 0x11, 0x1c, 0x91, 0x0c, 0x45, 0x01, 0xb9, 0x40, 0x4e, 0x90, 0xfb, 0xe9, 0x10,
	0x42, 0x0e, 0x10, 0x88, 0x52, 0x60, 0xda, 0x88, 0x81, 0xac, 0x24, 0xf2, 0x9f, 0xef, 0xe7, 0xcc,
	0x8f, 0x41, 0x7e, 0x02, 0x7a, 0xcd, 0xa4, 0x5c, 0xc7, 0x20, 0x81, 0xc7, 0xc0, 0xa3, 0x14, 0xf2,
	0x50, 0x2a, 0xa1, 0x05, 0x7e, 0x19, 0x43, 0xbe, 0xd3, 0x42, 0xfa, 0x5f, 0x93, 0x54, 0x5f, 0x16,
	0x9b, 0x30, 0x12, 0xd9, 0x2c, 0x11, 0x89, 0x98, 0x19, 0x7d, 0x53, 0x6c, 0xcd, 0xc9, 0x1c, 0xcc,
	0x5f, 0xc3, 0x91, 0x15, 0x1a, 0xfd, 0x05, 0xbd, 0x90, 0xf2, 0xb7, 0xe5, 0x49, 0xe1, 0xba, 0x80,
	0x5c, 0xe3, 0x2f, 0xa8, 0xcf, 0xa4, 0xfc, 0x1f, 0x7b, 0xce, 0xd8, 0x99, 0xbe, 0x5e, 0xbe, 0xab,
	0xca, 0x60, 0xb0, 0x15, 0x2a, 0xfb, 0x45, 0xcc, 0x35, 0xa1, 0x8d, 0x4c, 0xee, 0xba, 0xc8, 0x7f,
	0xca, 0x25, 0x97, 0x82, 0xe7, 0x80, 0x33, 0x34, 0xb0, 0x3b, 0xf6, 0x9c, 0x71, 0x77, 0xea, 0xce,
	0xbf, 0x85, 0x6d, 0xcb, 0xe1, 0x69, 0x34, 0xb4, 0x2f, 0x97, 0x9f, 0xaa, 0x32, 0xf8, 0xd0, 0xbc,
	0x6f, 0xfb, 0x11, 0x7a, 0x60, 0xef, 0xdf, 0x76, 0xd0, 0xc0, 0xe6, 0xf0, 0x4f, 0xe4, 0x4a, 0x16,
	0xed, 0x58, 0x02, 0x67, 0x2c, 0x83, 0x76, 0x98, 0x8f, 0x55, 0x19, 0xe0, 0xc6, 0xcc, 0x12, 0x09,
	0xb5, 0x4b, 0xf1, 0x0f, 0x84, 0x22, 0xc1, 0x73, 0xad, 0x58, 0xca, 0xb5, 0xd7, 0x31, 0xe0, 0xb0,
	0x2a, 0x83, 0xf7, 0x0d, 0xb8, 0xd7, 0x08, 0xb5, 0x0a, 0xf1, 0x02, 0xbd, 0x89, 0x0a, 0xa5, 0x80,
	0xeb, 0x73, 0x50, 0x79, 0x2a, 0xb8, 0xd7, 0x35, 0xe8, 0xa8, 0x2a, 0x83, 0x61, 0x8b, 0x1e, 0xe8,
	0x84, 0x1e, 0x01, 0xb5, 0x45, 0x06, 0xa0, 0x57, 0xfb, 0xd7, 0x7b, 0x63, 0x67, 0xfa, 0xca, 0xb6,
	0x38, 0xd4, 0x09, 0x3d, 0x02, 0xc8, 0xbd, 0x83, 0x3e, 0x9f, 0x8e, 0xf6, 0x42, 0x31, 0x29, 0x41,
	0xe1, 0x09, 0xea, 0x45, 0x22, 0x6e, 0x52, 0xe9, 0x2f, 0xdf, 0x56, 0x65, 0xe0, 0x3e, 0x0e, 0x17,
	0x03, 0xa1, 0x46, 0xac, 0x13, 0xac, 0xbf, 0x7f, 0x6e, 0xe4, 0x15, 0x4b, 0x79, 0x1b, 0x84, 0x95,
	0xa0, 0x25, 0x12, 0x6a, 0x97, 0xd6, 0x2b, 0x04, 0x4a, 0x09, 0xd5, 0x26, 0x60, 0xad, 0x90, 0xb9,
	0x26, 0xb4, 0x91, 0xf1, 0x3f, 0xd4, 0x8b, 0x99, 0x66, 0x66, 0x4a, 0x77, 0x3e, 0x79, 0xc6, 0x6e,
	0xd8, 0xbd, 0xd6, 0x28, 0xa1, 0xc6, 0x61, 0xf3, 0xc2, 0x2c, 0xf6, 0xf7, 0x87, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x31, 0x9e, 0x80, 0x2e, 0x03, 0x00, 0x00,
}
