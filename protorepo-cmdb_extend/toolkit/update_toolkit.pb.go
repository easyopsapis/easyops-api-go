// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_toolkit.proto

package toolkit

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	cmdb_extend "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
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
//UpdateToolkit请求
type UpdateToolkitRequest struct {
	//
	//模型Id
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//任意门Id
	ToolkitId string `protobuf:"bytes,2,opt,name=toolkitId,proto3" json:"toolkitId" form:"toolkitId"`
	//
	//任意门详情
	Toolkit              *cmdb_extend.Toolkit `protobuf:"bytes,3,opt,name=toolkit,proto3" json:"toolkit" form:"toolkit"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateToolkitRequest) Reset()         { *m = UpdateToolkitRequest{} }
func (m *UpdateToolkitRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateToolkitRequest) ProtoMessage()    {}
func (*UpdateToolkitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30bf0cfe8a755c2, []int{0}
}
func (m *UpdateToolkitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateToolkitRequest.Unmarshal(m, b)
}
func (m *UpdateToolkitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateToolkitRequest.Marshal(b, m, deterministic)
}
func (m *UpdateToolkitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateToolkitRequest.Merge(m, src)
}
func (m *UpdateToolkitRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateToolkitRequest.Size(m)
}
func (m *UpdateToolkitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateToolkitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateToolkitRequest proto.InternalMessageInfo

func (m *UpdateToolkitRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UpdateToolkitRequest) GetToolkitId() string {
	if m != nil {
		return m.ToolkitId
	}
	return ""
}

func (m *UpdateToolkitRequest) GetToolkit() *cmdb_extend.Toolkit {
	if m != nil {
		return m.Toolkit
	}
	return nil
}

//
//UpdateToolkitApi返回
type UpdateToolkitResponseWrapper struct {
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
	Data                 *cmdb_extend.Toolkit `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateToolkitResponseWrapper) Reset()         { *m = UpdateToolkitResponseWrapper{} }
func (m *UpdateToolkitResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateToolkitResponseWrapper) ProtoMessage()    {}
func (*UpdateToolkitResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30bf0cfe8a755c2, []int{1}
}
func (m *UpdateToolkitResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateToolkitResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateToolkitResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateToolkitResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateToolkitResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateToolkitResponseWrapper.Merge(m, src)
}
func (m *UpdateToolkitResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateToolkitResponseWrapper.Size(m)
}
func (m *UpdateToolkitResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateToolkitResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateToolkitResponseWrapper proto.InternalMessageInfo

func (m *UpdateToolkitResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateToolkitResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateToolkitResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateToolkitResponseWrapper) GetData() *cmdb_extend.Toolkit {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateToolkitRequest)(nil), "toolkit.UpdateToolkitRequest")
	proto.RegisterType((*UpdateToolkitResponseWrapper)(nil), "toolkit.UpdateToolkitResponseWrapper")
}

func init() { proto.RegisterFile("update_toolkit.proto", fileDescriptor_a30bf0cfe8a755c2) }

var fileDescriptor_a30bf0cfe8a755c2 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x15, 0x68, 0x81, 0xba, 0x88, 0x56, 0x26, 0x42, 0x51, 0x55, 0x29, 0x95, 0x91, 0x50,
	0x37, 0x8d, 0xa5, 0xb2, 0x01, 0x96, 0x91, 0x58, 0x74, 0x6b, 0x40, 0x2c, 0x2b, 0xc7, 0x76, 0x43,
	0x21, 0xc9, 0x35, 0xb6, 0x23, 0x75, 0x9e, 0x6e, 0xde, 0x24, 0x9a, 0x67, 0xc8, 0x13, 0x8c, 0xea,
	0xa4, 0x7f, 0xb3, 0x98, 0x55, 0x72, 0x7d, 0xce, 0xe7, 0x7b, 0x4e, 0x82, 0xc2, 0x5a, 0x4b, 0xee,
	0xd4, 0xd6, 0x01, 0x14, 0xff, 0xf6, 0x2e, 0xd1, 0x06, 0x1c, 0xe0, 0xd7, 0xfd, 0x38, 0x5b, 0xe5,
	0x7b, 0xf7, 0xa7, 0xce, 0x12, 0x01, 0x25, 0xcd, 0x21, 0x07, 0xea, 0xf5, 0xac, 0xde, 0xf9, 0xc9,
	0x0f, 0xfe, 0xad, 0xe3, 0x66, 0x3f, 0x72, 0x48, 0x14, 0xb7, 0x77, 0xa0, 0x6d, 0x52, 0x80, 0xe0,
	0x05, 0x15, 0x50, 0x39, 0xc3, 0x85, 0xb3, 0x1d, 0x69, 0x94, 0x86, 0x55, 0x09, 0x52, 0x15, 0x96,
	0xf6, 0x46, 0xea, 0x47, 0x2a, 0x4a, 0x99, 0x6d, 0xd5, 0xc1, 0xa9, 0x4a, 0xd2, 0x9b, 0x30, 0xb3,
	0x79, 0x0e, 0x90, 0x17, 0xea, 0xb2, 0xda, 0x3a, 0x53, 0x8b, 0x5e, 0x25, 0xf7, 0x01, 0x0a, 0x7f,
	0xf9, 0x0e, 0x3f, 0x3b, 0x8a, 0xa9, 0xff, 0xb5, 0xb2, 0x0e, 0x53, 0xf4, 0x06, 0xb2, 0xbf, 0x4a,
	0xb8, 0x8d, 0x8c, 0x82, 0x45, 0xb0, 0x1c, 0xa5, 0xef, 0xdb, 0x26, 0x9e, 0xec, 0xc0, 0x94, 0xdf,
	0xc8, 0x49, 0x21, 0xec, 0x6c, 0xc2, 0x6b, 0x34, 0xea, 0x17, 0x6f, 0x64, 0xf4, 0xc2, 0x13, 0x61,
	0xdb, 0xc4, 0xd3, 0x8e, 0x38, 0x4b, 0x84, 0x5d, 0x6c, 0x38, 0x45, 0xa7, 0x4f, 0x15, 0xbd, 0x5c,
	0x04, 0xcb, 0xf1, 0x3a, 0x4c, 0xae, 0x8a, 0x24, 0x7d, 0xa4, 0x14, 0xb7, 0x4d, 0xfc, 0xee, 0xe6,
	0x1e, 0xc2, 0x4e, 0x20, 0x79, 0x08, 0xd0, 0xfc, 0x49, 0x03, 0xab, 0xa1, 0xb2, 0xea, 0xb7, 0xe1,
	0x5a, 0x2b, 0x83, 0x3f, 0xa2, 0x81, 0x00, 0xa9, 0x7c, 0x8b, 0x61, 0x3a, 0x69, 0x9b, 0x78, 0xdc,
	0xdd, 0x75, 0x3c, 0x25, 0xcc, 0x8b, 0xf8, 0x0b, 0x1a, 0x1f, 0x9f, 0xdf, 0x0f, 0xba, 0xe0, 0xfb,
	0xaa, 0xcf, 0xff, 0xa1, 0x6d, 0x62, 0x7c, 0xf1, 0xf6, 0x22, 0x61, 0xd7, 0x56, 0xfc, 0x09, 0x0d,
	0x95, 0x31, 0x60, 0x7c, 0x83, 0x51, 0x3a, 0x6d, 0x9b, 0xf8, 0x6d, 0xc7, 0xf8, 0x63, 0xc2, 0x3a,
	0x19, 0x7f, 0x45, 0x03, 0xc9, 0x1d, 0x8f, 0x06, 0xcf, 0x14, 0xbd, 0x0a, 0x77, 0xf4, 0x12, 0xe6,
	0x91, 0xec, 0x95, 0xff, 0x57, 0x9f, 0x1f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x07, 0x39, 0xd1, 0x7c,
	0x6e, 0x02, 0x00, 0x00,
}