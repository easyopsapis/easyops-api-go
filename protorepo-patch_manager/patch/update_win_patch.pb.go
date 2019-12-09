// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_win_patch.proto

package patch

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
//UpdateWinPatch请求
type UpdateWinPatchRequest struct {
	//
	//补丁id
	PatchId string `protobuf:"bytes,1,opt,name=patchId,proto3" json:"patchId" form:"patchId"`
	//
	//补丁编号
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//发布时间
	ReleaseTime string `protobuf:"bytes,3,opt,name=releaseTime,proto3" json:"releaseTime" form:"releaseTime"`
	//
	//适用系统
	OsSystem string `protobuf:"bytes,4,opt,name=osSystem,proto3" json:"osSystem" form:"osSystem"`
	//
	//是否需要重启
	RequireReboot bool `protobuf:"varint,5,opt,name=requireReboot,proto3" json:"requireReboot" form:"requireReboot"`
	//
	//MSRC编号
	Msrc string `protobuf:"bytes,6,opt,name=msrc,proto3" json:"msrc" form:"msrc"`
	//
	//操作系统架构
	OsArchitecture string `protobuf:"bytes,7,opt,name=osArchitecture,proto3" json:"osArchitecture" form:"osArchitecture"`
	//
	//补丁大小,单位KB
	Size_ int32 `protobuf:"varint,8,opt,name=size,proto3" json:"size" form:"size"`
	//
	//说明
	Desc string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc" form:"desc"`
	//
	//补丁下载链接
	Url                  string   `protobuf:"bytes,10,opt,name=url,proto3" json:"url" form:"url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWinPatchRequest) Reset()         { *m = UpdateWinPatchRequest{} }
func (m *UpdateWinPatchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWinPatchRequest) ProtoMessage()    {}
func (*UpdateWinPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4abd90c3b15ed6, []int{0}
}
func (m *UpdateWinPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWinPatchRequest.Unmarshal(m, b)
}
func (m *UpdateWinPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWinPatchRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWinPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWinPatchRequest.Merge(m, src)
}
func (m *UpdateWinPatchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWinPatchRequest.Size(m)
}
func (m *UpdateWinPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWinPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWinPatchRequest proto.InternalMessageInfo

func (m *UpdateWinPatchRequest) GetPatchId() string {
	if m != nil {
		return m.PatchId
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetReleaseTime() string {
	if m != nil {
		return m.ReleaseTime
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetOsSystem() string {
	if m != nil {
		return m.OsSystem
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetRequireReboot() bool {
	if m != nil {
		return m.RequireReboot
	}
	return false
}

func (m *UpdateWinPatchRequest) GetMsrc() string {
	if m != nil {
		return m.Msrc
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetOsArchitecture() string {
	if m != nil {
		return m.OsArchitecture
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetSize_() int32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *UpdateWinPatchRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *UpdateWinPatchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

//
//UpdateWinPatch返回
type UpdateWinPatchResponse struct {
	//
	//补丁id
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateWinPatchResponse) Reset()         { *m = UpdateWinPatchResponse{} }
func (m *UpdateWinPatchResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateWinPatchResponse) ProtoMessage()    {}
func (*UpdateWinPatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4abd90c3b15ed6, []int{1}
}
func (m *UpdateWinPatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWinPatchResponse.Unmarshal(m, b)
}
func (m *UpdateWinPatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWinPatchResponse.Marshal(b, m, deterministic)
}
func (m *UpdateWinPatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWinPatchResponse.Merge(m, src)
}
func (m *UpdateWinPatchResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateWinPatchResponse.Size(m)
}
func (m *UpdateWinPatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWinPatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWinPatchResponse proto.InternalMessageInfo

func (m *UpdateWinPatchResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//UpdateWinPatchApi返回
type UpdateWinPatchResponseWrapper struct {
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
	Data                 *UpdateWinPatchResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateWinPatchResponseWrapper) Reset()         { *m = UpdateWinPatchResponseWrapper{} }
func (m *UpdateWinPatchResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateWinPatchResponseWrapper) ProtoMessage()    {}
func (*UpdateWinPatchResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4abd90c3b15ed6, []int{2}
}
func (m *UpdateWinPatchResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWinPatchResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateWinPatchResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWinPatchResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateWinPatchResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWinPatchResponseWrapper.Merge(m, src)
}
func (m *UpdateWinPatchResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateWinPatchResponseWrapper.Size(m)
}
func (m *UpdateWinPatchResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWinPatchResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWinPatchResponseWrapper proto.InternalMessageInfo

func (m *UpdateWinPatchResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateWinPatchResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateWinPatchResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateWinPatchResponseWrapper) GetData() *UpdateWinPatchResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateWinPatchRequest)(nil), "patch.UpdateWinPatchRequest")
	proto.RegisterType((*UpdateWinPatchResponse)(nil), "patch.UpdateWinPatchResponse")
	proto.RegisterType((*UpdateWinPatchResponseWrapper)(nil), "patch.UpdateWinPatchResponseWrapper")
}

func init() { proto.RegisterFile("update_win_patch.proto", fileDescriptor_fe4abd90c3b15ed6) }

var fileDescriptor_fe4abd90c3b15ed6 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x95, 0xdf, 0xd6, 0xfd, 0x71, 0x7f, 0x74, 0xc3, 0xb0, 0x29, 0x4c, 0x9a, 0x52, 0x99,
	0x69, 0x4a, 0x99, 0xdc, 0xbf, 0x0c, 0xb1, 0x5e, 0x50, 0x36, 0x89, 0x0b, 0xae, 0x40, 0x66, 0x68,
	0x12, 0x21, 0x9b, 0xdc, 0xd4, 0xeb, 0x22, 0x9a, 0x38, 0x73, 0x1c, 0x06, 0x5b, 0xf7, 0x42, 0x48,
	0x5c, 0xf0, 0x16, 0x3c, 0x45, 0x90, 0xb8, 0x84, 0xbb, 0x3c, 0x01, 0xb2, 0xb3, 0x6e, 0x59, 0xb5,
	0x2b, 0xfb, 0x9c, 0xcf, 0xf7, 0xeb, 0x73, 0xec, 0x9c, 0x16, 0xac, 0x26, 0xd1, 0x80, 0x4a, 0x76,
	0x74, 0xe6, 0x87, 0x47, 0x11, 0x95, 0xde, 0x49, 0x3d, 0x12, 0x5c, 0x72, 0x58, 0xd2, 0xc1, 0x1a,
	0x1e, 0xfa, 0xf2, 0x24, 0xe9, 0xd7, 0x3d, 0x1e, 0x34, 0x86, 0x7c, 0xc8, 0x1b, 0x9a, 0xf6, 0x93,
	0x63, 0x1d, 0xe9, 0x40, 0xef, 0x72, 0xd7, 0xda, 0xb3, 0x82, 0x3c, 0x38, 0xf3, 0xe5, 0x27, 0x7e,
	0xd6, 0x18, 0x72, 0xac, 0x21, 0xfe, 0x4c, 0x47, 0xfe, 0x80, 0x4a, 0x2e, 0xe2, 0xc6, 0xf5, 0x36,
	0xf7, 0xa1, 0x1f, 0x25, 0xb0, 0xf2, 0x5e, 0x37, 0x72, 0xe0, 0x87, 0x6f, 0x55, 0x65, 0xc2, 0x4e,
	0x13, 0x16, 0x4b, 0xf8, 0x12, 0xcc, 0xeb, 0x4e, 0x5e, 0x0f, 0x4c, 0xa3, 0x6a, 0xd8, 0x8b, 0x7b,
	0x9b, 0x59, 0x6a, 0x55, 0x8e, 0xb9, 0x08, 0xba, 0xe8, 0x0a, 0xa0, 0xdf, 0xbf, 0xac, 0x65, 0x50,
	0x39, 0x74, 0x9a, 0x78, 0x87, 0xe2, 0x73, 0xf7, 0xa2, 0xd5, 0xb9, 0xdc, 0x20, 0x13, 0x1b, 0x7c,
	0x0c, 0x66, 0x43, 0x1a, 0x30, 0xf3, 0x3f, 0x6d, 0x5f, 0xca, 0x52, 0xab, 0x9c, 0xdb, 0x55, 0x16,
	0x11, 0x0d, 0xe1, 0x5f, 0x03, 0x94, 0x05, 0x1b, 0x31, 0x1a, 0xb3, 0x7d, 0x3f, 0x60, 0xe6, 0x8c,
	0x16, 0xff, 0x34, 0xb2, 0xd4, 0x82, 0xb9, 0xba, 0x40, 0x55, 0xc1, 0xef, 0x06, 0xf8, 0x66, 0x1c,
	0xda, 0x76, 0xaf, 0xeb, 0xb4, 0xf0, 0x8e, 0xab, 0x6a, 0xbb, 0x4f, 0x6a, 0x3d, 0xbd, 0x5e, 0x3c,
	0xbd, 0xac, 0x61, 0xbb, 0xe5, 0x34, 0x71, 0xdb, 0x1d, 0x37, 0x35, 0xaf, 0x61, 0xbb, 0xe3, 0x34,
	0x71, 0x6b, 0x12, 0x8f, 0x9d, 0x16, 0x6e, 0xe7, 0xae, 0x9a, 0xb3, 0x5f, 0x75, 0xed, 0xb6, 0xd3,
	0xc4, 0x1d, 0x77, 0xac, 0x35, 0x79, 0xba, 0x6b, 0x3b, 0x4d, 0xbc, 0x3d, 0x09, 0x6e, 0xf6, 0xf6,
	0xc7, 0xba, 0x5e, 0xb7, 0x6a, 0x3d, 0xfb, 0xc3, 0xd8, 0xd9, 0xc2, 0xae, 0xdd, 0xeb, 0xde, 0x61,
	0x2f, 0xb8, 0x7b, 0x1b, 0xa4, 0x78, 0x3b, 0xd8, 0x00, 0x0b, 0x3c, 0x7e, 0xf7, 0x35, 0x96, 0x2c,
	0x30, 0x67, 0xf5, 0x4d, 0x1f, 0x64, 0xa9, 0xb5, 0x94, 0x5f, 0x74, 0x42, 0x10, 0xb9, 0x16, 0xc1,
	0x17, 0xe0, 0x9e, 0x60, 0xa7, 0x89, 0x2f, 0x18, 0x61, 0x7d, 0xce, 0xa5, 0x59, 0xaa, 0x1a, 0xf6,
	0xc2, 0x9e, 0x99, 0xa5, 0xd6, 0xc3, 0xc9, 0xf3, 0x14, 0x30, 0x22, 0xb7, 0xe5, 0xea, 0x1b, 0x04,
	0xb1, 0xf0, 0xcc, 0xb9, 0xe9, 0x6f, 0xa0, 0xb2, 0x88, 0x68, 0x08, 0x77, 0x41, 0x85, 0xc7, 0xbb,
	0xc2, 0x3b, 0xf1, 0x25, 0xf3, 0x64, 0x22, 0x98, 0x39, 0xaf, 0xe5, 0x8f, 0xb2, 0xd4, 0x5a, 0x99,
	0xf4, 0x56, 0xe4, 0x88, 0x4c, 0x19, 0x54, 0x9d, 0xd8, 0x3f, 0x67, 0xe6, 0x42, 0xd5, 0xb0, 0x4b,
	0xc5, 0x3a, 0x2a, 0x8b, 0x88, 0x86, 0x4a, 0x34, 0x60, 0xb1, 0x67, 0x2e, 0x4e, 0x37, 0xa3, 0xb2,
	0x88, 0x68, 0x08, 0xab, 0x60, 0x26, 0x11, 0x23, 0x13, 0x68, 0x4d, 0x25, 0x4b, 0x2d, 0x90, 0x6b,
	0x12, 0x31, 0x42, 0x44, 0x21, 0xf4, 0x06, 0xac, 0x4e, 0x8f, 0x6c, 0x1c, 0xf1, 0x30, 0x66, 0x70,
	0x1b, 0x00, 0x3f, 0x8c, 0x25, 0x0d, 0x3d, 0x76, 0x3d, 0xb6, 0x2b, 0x59, 0x6a, 0xdd, 0xcf, 0x8f,
	0xb8, 0x61, 0x88, 0x14, 0x84, 0xe8, 0x8f, 0x01, 0xd6, 0xef, 0x3e, 0xf1, 0x40, 0xd0, 0x28, 0x62,
	0x42, 0x75, 0xee, 0xf1, 0x01, 0xd3, 0x47, 0xde, 0xba, 0x9e, 0xca, 0x22, 0xa2, 0x21, 0x7c, 0x0e,
	0xca, 0x6a, 0x7d, 0xf5, 0x25, 0x1a, 0x51, 0x3f, 0xbc, 0x1a, 0xfb, 0xd5, 0x9b, 0x41, 0x2e, 0x40,
	0x44, 0x8a, 0x52, 0xb8, 0x09, 0x4a, 0x4c, 0x08, 0x2e, 0xae, 0xa6, 0x7f, 0x39, 0x4b, 0xad, 0xff,
	0x73, 0x8f, 0x4e, 0x23, 0x92, 0x63, 0xb8, 0x07, 0x66, 0x07, 0x54, 0x52, 0x3d, 0x3a, 0xe5, 0xf6,
	0x7a, 0x3d, 0xff, 0xdf, 0xb8, 0xbb, 0xf5, 0x5b, 0xef, 0x4b, 0x25, 0x55, 0xef, 0x4b, 0x25, 0xed,
	0xcf, 0xe9, 0x1f, 0x7e, 0xe7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xff, 0x99, 0xe2, 0x80,
	0x04, 0x00, 0x00,
}
