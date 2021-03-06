// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

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
//UpdateInspectionInfo请求
type UpdateInspectionInfoRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//套件名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//备注
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//套件索引
	Index                int32    `protobuf:"varint,4,opt,name=index,proto3" json:"index" form:"index"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInspectionInfoRequest) Reset()         { *m = UpdateInspectionInfoRequest{} }
func (m *UpdateInspectionInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInspectionInfoRequest) ProtoMessage()    {}
func (*UpdateInspectionInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{0}
}
func (m *UpdateInspectionInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInspectionInfoRequest.Unmarshal(m, b)
}
func (m *UpdateInspectionInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInspectionInfoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInspectionInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInspectionInfoRequest.Merge(m, src)
}
func (m *UpdateInspectionInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInspectionInfoRequest.Size(m)
}
func (m *UpdateInspectionInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInspectionInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInspectionInfoRequest proto.InternalMessageInfo

func (m *UpdateInspectionInfoRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *UpdateInspectionInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateInspectionInfoRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *UpdateInspectionInfoRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

//
//UpdateInspectionInfoApi返回
type UpdateInspectionInfoResponseWrapper struct {
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
	Data                 *inspection.InspectionInfo `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UpdateInspectionInfoResponseWrapper) Reset()         { *m = UpdateInspectionInfoResponseWrapper{} }
func (m *UpdateInspectionInfoResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateInspectionInfoResponseWrapper) ProtoMessage()    {}
func (*UpdateInspectionInfoResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{1}
}
func (m *UpdateInspectionInfoResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInspectionInfoResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateInspectionInfoResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInspectionInfoResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateInspectionInfoResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInspectionInfoResponseWrapper.Merge(m, src)
}
func (m *UpdateInspectionInfoResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateInspectionInfoResponseWrapper.Size(m)
}
func (m *UpdateInspectionInfoResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInspectionInfoResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInspectionInfoResponseWrapper proto.InternalMessageInfo

func (m *UpdateInspectionInfoResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateInspectionInfoResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateInspectionInfoResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateInspectionInfoResponseWrapper) GetData() *inspection.InspectionInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateInspectionInfoRequest)(nil), "info.UpdateInspectionInfoRequest")
	proto.RegisterType((*UpdateInspectionInfoResponseWrapper)(nil), "info.UpdateInspectionInfoResponseWrapper")
}

func init() { proto.RegisterFile("update.proto", fileDescriptor_3f0fa214029f1c21) }

var fileDescriptor_3f0fa214029f1c21 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0xa6, 0x3a, 0x23, 0x9a, 0x59, 0xd8, 0xa5, 0x07, 0x19, 0x46, 0xa4, 0x4b, 0x56, 0x64, 0x0f,
	0xb6, 0x59, 0x5d, 0x5c, 0xd4, 0x8b, 0xb8, 0xe0, 0x61, 0x4e, 0x42, 0x40, 0x04, 0x17, 0x95, 0x4c,
	0x9b, 0xa9, 0xc1, 0x36, 0x2f, 0x26, 0xa9, 0x3b, 0x2a, 0xfe, 0xab, 0x23, 0x78, 0xf7, 0xd2, 0x7f,
	0x40, 0xc9, 0x6b, 0x67, 0xa6, 0xca, 0xf6, 0xd2, 0xbc, 0xf7, 0xfd, 0xc8, 0xf7, 0xd1, 0x92, 0xbd,
	0xc6, 0x14, 0xc2, 0xcb, 0xcc, 0x58, 0xf0, 0x10, 0x8f, 0x94, 0x5e, 0xc2, 0x2c, 0x2d, 0x95, 0xff,
	0xd8, 0x2c, 0xb2, 0x1c, 0x6a, 0x56, 0x42, 0x09, 0x0c, 0xc1, 0x45, 0xb3, 0xc4, 0x09, 0x07, 0x3c,
	0x75, 0xa2, 0xd9, 0xab, 0x12, 0x32, 0x29, 0xdc, 0x57, 0x30, 0x2e, 0xab, 0x20, 0x17, 0x15, 0xcb,
	0x41, 0x7b, 0x2b, 0x72, 0xef, 0x3a, 0xa5, 0x95, 0x06, 0xd2, 0x1a, 0x0a, 0x59, 0x39, 0xd6, 0x13,
	0x19, 0x8e, 0x4c, 0x69, 0x67, 0x64, 0xee, 0x15, 0x68, 0x16, 0x6e, 0xee, 0x0d, 0xcf, 0x06, 0xf7,
	0xd7, 0x97, 0xca, 0x7f, 0x82, 0x4b, 0x56, 0x42, 0x8a, 0x60, 0xfa, 0x45, 0x54, 0xaa, 0x10, 0x1e,
	0xac, 0x63, 0xdb, 0x63, 0xa7, 0xa3, 0x6d, 0x44, 0xee, 0xbc, 0xc6, 0x3a, 0xf3, 0xad, 0xef, 0x5c,
	0x2f, 0x81, 0xcb, 0xcf, 0x8d, 0x74, 0x3e, 0xe6, 0xe4, 0xa6, 0xa9, 0x9a, 0x52, 0xe9, 0x79, 0x31,
	0x8d, 0x0e, 0xa3, 0xe3, 0x5b, 0xe7, 0x67, 0xed, 0x3a, 0xd9, 0x5f, 0x82, 0xad, 0x9f, 0xd1, 0x0d,
	0x42, 0x7f, 0xfd, 0x4c, 0x12, 0x72, 0xf7, 0xfd, 0x85, 0x48, 0xbf, 0xbd, 0x48, 0xdf, 0x7e, 0x78,
	0x77, 0x71, 0x92, 0x3e, 0xdd, 0x9c, 0xbf, 0x9f, 0x3c, 0x38, 0x7d, 0xf8, 0xe3, 0x1e, 0xdf, 0xfa,
	0xc4, 0x47, 0x64, 0xa4, 0x45, 0x2d, 0xa7, 0xd7, 0xd0, 0x6f, 0xbf, 0x5d, 0x27, 0x93, 0xce, 0x2f,
	0x6c, 0x29, 0x47, 0x30, 0x90, 0x6a, 0x59, 0xc3, 0xf4, 0xfa, 0xff, 0xa4, 0xb0, 0xa5, 0x1c, 0xc1,
	0xf8, 0x31, 0x19, 0x2b, 0x5d, 0xc8, 0xd5, 0x74, 0x74, 0x18, 0x1d, 0x8f, 0xcf, 0x93, 0x76, 0x9d,
	0xec, 0x75, 0x2c, 0x5c, 0x87, 0x5c, 0x93, 0x83, 0x3f, 0x9b, 0x27, 0xe2, 0x1d, 0x9b, 0xfe, 0x8e,
	0xc8, 0xd1, 0xd5, 0xa5, 0x9d, 0x01, 0xed, 0xe4, 0x1b, 0x2b, 0x8c, 0x91, 0x36, 0x64, 0xc8, 0xa1,
	0x90, 0x58, 0x7c, 0x3c, 0xcc, 0x10, 0xb6, 0x94, 0x23, 0x18, 0x3f, 0x21, 0x93, 0xf0, 0x7e, 0xb9,
	0x32, 0x95, 0x50, 0xba, 0x2f, 0x75, 0xbb, 0x5d, 0x27, 0xf1, 0x8e, 0xdb, 0x83, 0x94, 0x0f, 0xa9,
	0xf1, 0x7d, 0x32, 0x96, 0xd6, 0x82, 0xed, 0x3b, 0x1e, 0xec, 0xd2, 0xe3, 0x9a, 0xf2, 0x0e, 0x8e,
	0x9f, 0x93, 0x51, 0x21, 0xbc, 0xc0, 0x92, 0x93, 0x47, 0xb3, 0x6c, 0xf7, 0x07, 0x64, 0xff, 0xe6,
	0x1f, 0x46, 0x0c, 0x0a, 0xca, 0x51, 0xb8, 0xb8, 0x81, 0xdf, 0xfa, 0xf4, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x0f, 0xc0, 0x85, 0xb9, 0x02, 0x00, 0x00,
}
