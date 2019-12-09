// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_idc_all_rack_unit.proto

package custom

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//GetIdcAllRackUnit请求
type GetIdcAllRackUnitRequest struct {
	//
	//实例ID
	InstanceId           string   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id" form:"instance_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIdcAllRackUnitRequest) Reset()         { *m = GetIdcAllRackUnitRequest{} }
func (m *GetIdcAllRackUnitRequest) String() string { return proto.CompactTextString(m) }
func (*GetIdcAllRackUnitRequest) ProtoMessage()    {}
func (*GetIdcAllRackUnitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86210784bfcd87d5, []int{0}
}
func (m *GetIdcAllRackUnitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIdcAllRackUnitRequest.Unmarshal(m, b)
}
func (m *GetIdcAllRackUnitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIdcAllRackUnitRequest.Marshal(b, m, deterministic)
}
func (m *GetIdcAllRackUnitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIdcAllRackUnitRequest.Merge(m, src)
}
func (m *GetIdcAllRackUnitRequest) XXX_Size() int {
	return xxx_messageInfo_GetIdcAllRackUnitRequest.Size(m)
}
func (m *GetIdcAllRackUnitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIdcAllRackUnitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIdcAllRackUnitRequest proto.InternalMessageInfo

func (m *GetIdcAllRackUnitRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//GetIdcAllRackUnitApi返回
type GetIdcAllRackUnitResponseWrapper struct {
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
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetIdcAllRackUnitResponseWrapper) Reset()         { *m = GetIdcAllRackUnitResponseWrapper{} }
func (m *GetIdcAllRackUnitResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetIdcAllRackUnitResponseWrapper) ProtoMessage()    {}
func (*GetIdcAllRackUnitResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_86210784bfcd87d5, []int{1}
}
func (m *GetIdcAllRackUnitResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIdcAllRackUnitResponseWrapper.Unmarshal(m, b)
}
func (m *GetIdcAllRackUnitResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIdcAllRackUnitResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetIdcAllRackUnitResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIdcAllRackUnitResponseWrapper.Merge(m, src)
}
func (m *GetIdcAllRackUnitResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetIdcAllRackUnitResponseWrapper.Size(m)
}
func (m *GetIdcAllRackUnitResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIdcAllRackUnitResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetIdcAllRackUnitResponseWrapper proto.InternalMessageInfo

func (m *GetIdcAllRackUnitResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetIdcAllRackUnitResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetIdcAllRackUnitResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetIdcAllRackUnitResponseWrapper) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetIdcAllRackUnitRequest)(nil), "custom.GetIdcAllRackUnitRequest")
	proto.RegisterType((*GetIdcAllRackUnitResponseWrapper)(nil), "custom.GetIdcAllRackUnitResponseWrapper")
}

func init() { proto.RegisterFile("get_idc_all_rack_unit.proto", fileDescriptor_86210784bfcd87d5) }

var fileDescriptor_86210784bfcd87d5 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xc9, 0xff, 0xdf, 0x16, 0xdc, 0x88, 0x96, 0x1c, 0x34, 0x54, 0x21, 0x65, 0x15, 0x29,
	0x42, 0x12, 0xb5, 0x20, 0x2a, 0x5e, 0x2c, 0x88, 0x14, 0x6f, 0x11, 0xf1, 0x20, 0x1a, 0xb6, 0x9b,
	0xed, 0xba, 0x34, 0xc9, 0xc4, 0xcd, 0xc6, 0x8a, 0xe2, 0x57, 0xad, 0xe0, 0xc5, 0x7b, 0x3f, 0x81,
	0x74, 0xd3, 0xda, 0x1c, 0x3c, 0xed, 0xcc, 0xfc, 0xde, 0xdb, 0x79, 0x0c, 0xda, 0xe2, 0x4c, 0x85,
	0x22, 0xa2, 0x21, 0x89, 0xe3, 0x50, 0x12, 0x3a, 0x0a, 0x8b, 0x54, 0x28, 0x2f, 0x93, 0xa0, 0xc0,
	0x6a, 0xd0, 0x22, 0x57, 0x90, 0xb4, 0x5c, 0x2e, 0xd4, 0x53, 0x31, 0xf0, 0x28, 0x24, 0x3e, 0x07,
	0x0e, 0xbe, 0xc6, 0x83, 0x62, 0xa8, 0x3b, 0xdd, 0xe8, 0xaa, 0xb4, 0xb5, 0x8e, 0x2b, 0xf2, 0x64,
	0x2c, 0xd4, 0x08, 0xc6, 0x3e, 0x07, 0x57, 0x43, 0xf7, 0x85, 0xc4, 0x22, 0x22, 0x0a, 0x64, 0xee,
	0xff, 0x96, 0x73, 0xdf, 0x36, 0x07, 0xe0, 0x31, 0x5b, 0xfe, 0x9e, 0x2b, 0x59, 0xd0, 0x79, 0x18,
	0xcc, 0x91, 0x7d, 0xc5, 0x54, 0x3f, 0xa2, 0x17, 0x71, 0x1c, 0x10, 0x3a, 0xba, 0x4d, 0x85, 0x0a,
	0xd8, 0x73, 0xc1, 0x72, 0x65, 0x5d, 0x23, 0x53, 0xa4, 0xb9, 0x22, 0x29, 0x65, 0xa1, 0x88, 0x6c,
	0xa3, 0x6d, 0x74, 0x56, 0x7a, 0xfb, 0xd3, 0x89, 0x63, 0x0d, 0x41, 0x26, 0x67, 0xb8, 0x02, 0xf1,
	0xd7, 0xa7, 0xd3, 0x44, 0x6b, 0x8f, 0xf7, 0x07, 0xee, 0x29, 0x71, 0xdf, 0x1e, 0xde, 0x0f, 0xbb,
	0x1f, 0xbb, 0x01, 0x5a, 0x28, 0xfa, 0x11, 0xfe, 0x36, 0x50, 0xfb, 0x8f, 0x4d, 0x79, 0x06, 0x69,
	0xce, 0xee, 0x24, 0xc9, 0x32, 0x26, 0xad, 0x1d, 0x54, 0xa3, 0x10, 0x31, 0xbd, 0xaa, 0xde, 0x5b,
	0x9f, 0x4e, 0x1c, 0xb3, 0x5c, 0x35, 0x9b, 0xe2, 0x40, 0x43, 0xeb, 0x04, 0x99, 0xb3, 0xf7, 0xf2,
	0x35, 0x8b, 0x89, 0x48, 0xed, 0x7f, 0x3a, 0xd6, 0xc6, 0x32, 0x56, 0x05, 0xe2, 0xa0, 0x2a, 0xb5,
	0xf6, 0x50, 0x9d, 0x49, 0x09, 0xd2, 0xfe, 0xaf, 0x3d, 0xcd, 0xe9, 0xc4, 0x59, 0x2d, 0x3d, 0x7a,
	0x8c, 0x83, 0x12, 0x5b, 0xe7, 0xa8, 0x16, 0x11, 0x45, 0xec, 0x5a, 0xdb, 0xe8, 0x98, 0x47, 0x9b,
	0x5e, 0x79, 0x41, 0x6f, 0x71, 0x41, 0xef, 0x46, 0x5f, 0xb0, 0x9a, 0x6f, 0x26, 0xc7, 0x81, 0x76,
	0x0d, 0x1a, 0x5a, 0xd7, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x10, 0xd8, 0x01, 0x05, 0x02,
	0x00, 0x00,
}
