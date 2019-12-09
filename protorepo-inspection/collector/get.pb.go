// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package collector

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
//GetCollector请求
type GetCollectorRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//采集脚本Id
	CollectorId          string   `protobuf:"bytes,2,opt,name=collectorId,proto3" json:"collectorId" form:"collectorId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCollectorRequest) Reset()         { *m = GetCollectorRequest{} }
func (m *GetCollectorRequest) String() string { return proto.CompactTextString(m) }
func (*GetCollectorRequest) ProtoMessage()    {}
func (*GetCollectorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetCollectorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollectorRequest.Unmarshal(m, b)
}
func (m *GetCollectorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollectorRequest.Marshal(b, m, deterministic)
}
func (m *GetCollectorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollectorRequest.Merge(m, src)
}
func (m *GetCollectorRequest) XXX_Size() int {
	return xxx_messageInfo_GetCollectorRequest.Size(m)
}
func (m *GetCollectorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollectorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollectorRequest proto.InternalMessageInfo

func (m *GetCollectorRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *GetCollectorRequest) GetCollectorId() string {
	if m != nil {
		return m.CollectorId
	}
	return ""
}

//
//GetCollectorApi返回
type GetCollectorResponseWrapper struct {
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
	Data                 *inspection.InspectionCollector `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetCollectorResponseWrapper) Reset()         { *m = GetCollectorResponseWrapper{} }
func (m *GetCollectorResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetCollectorResponseWrapper) ProtoMessage()    {}
func (*GetCollectorResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetCollectorResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollectorResponseWrapper.Unmarshal(m, b)
}
func (m *GetCollectorResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollectorResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetCollectorResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollectorResponseWrapper.Merge(m, src)
}
func (m *GetCollectorResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetCollectorResponseWrapper.Size(m)
}
func (m *GetCollectorResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollectorResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollectorResponseWrapper proto.InternalMessageInfo

func (m *GetCollectorResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetCollectorResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetCollectorResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetCollectorResponseWrapper) GetData() *inspection.InspectionCollector {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCollectorRequest)(nil), "collector.GetCollectorRequest")
	proto.RegisterType((*GetCollectorResponseWrapper)(nil), "collector.GetCollectorResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x5f, 0x6b, 0x14, 0x31,
	0x14, 0xc5, 0x19, 0xdd, 0x8a, 0x9b, 0x15, 0xaa, 0x29, 0xc8, 0x52, 0x91, 0x29, 0xb1, 0x48, 0x1f,
	0xcc, 0xa4, 0xb6, 0x52, 0xd4, 0xb7, 0xd6, 0x7f, 0xec, 0x8b, 0x42, 0x10, 0x04, 0x4b, 0x2b, 0xd9,
	0x24, 0x1b, 0x07, 0xb3, 0x73, 0x63, 0x92, 0xb5, 0x6a, 0xd9, 0x2f, 0xe6, 0x87, 0x59, 0xc1, 0x47,
	0x1f, 0xf7, 0x13, 0xc8, 0x66, 0x76, 0xc6, 0xf1, 0xdd, 0xb7, 0x7b, 0xe7, 0xdc, 0xf3, 0x9b, 0x73,
	0x08, 0xea, 0x1b, 0x1d, 0x0b, 0xe7, 0x21, 0x02, 0xee, 0x4b, 0xb0, 0x56, 0xcb, 0x08, 0x7e, 0x9b,
	0x9a, 0x32, 0x7e, 0x9c, 0x8d, 0x0b, 0x09, 0x53, 0x66, 0xc0, 0x00, 0x4b, 0x17, 0xe3, 0xd9, 0x24,
	0x6d, 0x69, 0x49, 0x53, 0xed, 0xdc, 0x7e, 0x6d, 0xa0, 0xd0, 0x22, 0x7c, 0x03, 0x17, 0x0a, 0x0b,
	0x52, 0x58, 0x26, 0xa1, 0x8a, 0x5e, 0xc8, 0x18, 0x6a, 0xa7, 0xd7, 0x0e, 0xe8, 0x14, 0x94, 0xb6,
	0x81, 0xad, 0x0f, 0x59, 0x5a, 0x59, 0x59, 0x05, 0xa7, 0x65, 0x2c, 0xa1, 0x62, 0xc2, 0x9b, 0x35,
	0xef, 0xed, 0x7f, 0xe1, 0xb5, 0x75, 0xd6, 0xd4, 0xa3, 0x4e, 0xa9, 0xe9, 0x45, 0x19, 0x3f, 0xc1,
	0x05, 0x33, 0x40, 0x93, 0x48, 0xbf, 0x08, 0x5b, 0x2a, 0x11, 0xc1, 0x07, 0xd6, 0x8e, 0xb5, 0x8f,
	0xfc, 0xc8, 0xd0, 0xd6, 0x2b, 0x1d, 0x9f, 0x35, 0x38, 0xae, 0x3f, 0xcf, 0x74, 0x88, 0x98, 0xa3,
	0xeb, 0xce, 0xce, 0x4c, 0x59, 0x8d, 0xd4, 0x30, 0xdb, 0xc9, 0xf6, 0xfa, 0x27, 0x47, 0xcb, 0x45,
	0xbe, 0x39, 0x01, 0x3f, 0x7d, 0x4a, 0x1a, 0x85, 0xfc, 0xfa, 0x99, 0xe7, 0xe8, 0xee, 0xf9, 0xa9,
	0xa0, 0xdf, 0x8f, 0xe9, 0xfb, 0x0f, 0x67, 0xa7, 0xfb, 0xf4, 0x49, 0x33, 0x5f, 0xee, 0x3f, 0x38,
	0x7c, 0x38, 0xdf, 0xe5, 0x2d, 0x07, 0xbf, 0x41, 0x83, 0x36, 0xf6, 0x48, 0x0d, 0xaf, 0x24, 0x2c,
	0x5d, 0x2e, 0x72, 0x5c, 0x63, 0x3b, 0xe2, 0x8a, 0xbc, 0x85, 0x6e, 0x9d, 0xd7, 0xc0, 0xc9, 0x31,
	0x7d, 0x79, 0x76, 0x79, 0xf0, 0x68, 0xbe, 0xcb, 0xbb, 0x04, 0xf2, 0x3b, 0x43, 0x77, 0xfe, 0x0d,
	0x1f, 0x1c, 0x54, 0x41, 0xbf, 0xf3, 0xc2, 0x39, 0xed, 0xf1, 0x3d, 0xd4, 0x93, 0xa0, 0x74, 0x2a,
	0xb0, 0x71, 0xb2, 0xb9, 0x5c, 0xe4, 0x83, 0xe6, 0x4f, 0x4a, 0x13, 0x9e, 0x44, 0xfc, 0x78, 0x95,
	0x4a, 0xe9, 0x17, 0x5f, 0x9d, 0x15, 0x65, 0xb5, 0x4e, 0x75, 0xbb, 0x9b, 0xaa, 0x15, 0x09, 0xef,
	0x9e, 0xe2, 0xfb, 0x68, 0x43, 0x7b, 0x0f, 0x7e, 0x78, 0x35, 0x79, 0x6e, 0x2e, 0x17, 0xf9, 0x8d,
	0xda, 0x93, 0x3e, 0x13, 0x5e, 0xcb, 0xf8, 0x39, 0xea, 0x29, 0x11, 0xc5, 0xb0, 0xb7, 0x93, 0xed,
	0x0d, 0x0e, 0xf2, 0xe2, 0xef, 0x33, 0x16, 0xa3, 0x76, 0x6c, 0x4b, 0x74, 0x73, 0xae, 0x6c, 0x84,
	0x27, 0xf7, 0xf8, 0x5a, 0x7a, 0xb0, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x06, 0x29,
	0xe9, 0xd5, 0x02, 0x00, 0x00,
}