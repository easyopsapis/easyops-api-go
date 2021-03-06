// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fetch_origin_view.proto

package view

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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
//FetchOriginView请求
type FetchOriginViewRequest struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例ID
	InstanceId string `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//相关关系字段
	Relations []*topology.Relation `protobuf:"bytes,3,rep,name=relations,proto3" json:"relations" form:"relations"`
	//
	//数据源
	DataSource string `protobuf:"bytes,4,opt,name=dataSource,proto3" json:"dataSource" form:"dataSource"`
	//
	//隐藏关系
	HideLinks            []*topology.HideLink `protobuf:"bytes,5,rep,name=hideLinks,proto3" json:"hideLinks" form:"hideLinks"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FetchOriginViewRequest) Reset()         { *m = FetchOriginViewRequest{} }
func (m *FetchOriginViewRequest) String() string { return proto.CompactTextString(m) }
func (*FetchOriginViewRequest) ProtoMessage()    {}
func (*FetchOriginViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7782f8e8def56600, []int{0}
}
func (m *FetchOriginViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOriginViewRequest.Unmarshal(m, b)
}
func (m *FetchOriginViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOriginViewRequest.Marshal(b, m, deterministic)
}
func (m *FetchOriginViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOriginViewRequest.Merge(m, src)
}
func (m *FetchOriginViewRequest) XXX_Size() int {
	return xxx_messageInfo_FetchOriginViewRequest.Size(m)
}
func (m *FetchOriginViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOriginViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOriginViewRequest proto.InternalMessageInfo

func (m *FetchOriginViewRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *FetchOriginViewRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *FetchOriginViewRequest) GetRelations() []*topology.Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

func (m *FetchOriginViewRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

func (m *FetchOriginViewRequest) GetHideLinks() []*topology.HideLink {
	if m != nil {
		return m.HideLinks
	}
	return nil
}

//
//FetchOriginViewApi返回
type FetchOriginViewResponseWrapper struct {
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
	Data                 *topology.View `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchOriginViewResponseWrapper) Reset()         { *m = FetchOriginViewResponseWrapper{} }
func (m *FetchOriginViewResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*FetchOriginViewResponseWrapper) ProtoMessage()    {}
func (*FetchOriginViewResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7782f8e8def56600, []int{1}
}
func (m *FetchOriginViewResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOriginViewResponseWrapper.Unmarshal(m, b)
}
func (m *FetchOriginViewResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOriginViewResponseWrapper.Marshal(b, m, deterministic)
}
func (m *FetchOriginViewResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOriginViewResponseWrapper.Merge(m, src)
}
func (m *FetchOriginViewResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_FetchOriginViewResponseWrapper.Size(m)
}
func (m *FetchOriginViewResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOriginViewResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOriginViewResponseWrapper proto.InternalMessageInfo

func (m *FetchOriginViewResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FetchOriginViewResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *FetchOriginViewResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FetchOriginViewResponseWrapper) GetData() *topology.View {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchOriginViewRequest)(nil), "view.FetchOriginViewRequest")
	proto.RegisterType((*FetchOriginViewResponseWrapper)(nil), "view.FetchOriginViewResponseWrapper")
}

func init() { proto.RegisterFile("fetch_origin_view.proto", fileDescriptor_7782f8e8def56600) }

var fileDescriptor_7782f8e8def56600 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd7, 0x16, 0x31, 0x17, 0xb1, 0x61, 0x60, 0x44, 0xbd, 0x20, 0x93, 0x91, 0xd0, 0x6e,
	0x1a, 0x4b, 0x9b, 0x40, 0x88, 0xcb, 0x4a, 0x4c, 0x4c, 0x02, 0x21, 0x8c, 0x04, 0x97, 0x93, 0x9b,
	0x9c, 0xa6, 0x66, 0x69, 0x4e, 0xb0, 0xdd, 0x95, 0x3d, 0x21, 0x2f, 0x81, 0xf2, 0x10, 0x79, 0x02,
	0x64, 0x27, 0x6d, 0xa2, 0x71, 0x9b, 0xab, 0xf8, 0xe4, 0xfb, 0x39, 0x3f, 0xfa, 0xc8, 0x8b, 0x25,
	0xd8, 0x78, 0x75, 0x8d, 0x5a, 0xa5, 0x2a, 0xbf, 0xbe, 0x55, 0xb0, 0x8d, 0x0a, 0x8d, 0x16, 0xe9,
	0xc8, 0xbd, 0xa7, 0xb3, 0x54, 0xd9, 0xd5, 0x66, 0x11, 0xc5, 0xb8, 0xe6, 0x29, 0xa6, 0xc8, 0x3d,
	0xb8, 0xd8, 0x2c, 0x7d, 0xe5, 0x0b, 0xff, 0xaa, 0x45, 0xd3, 0xaf, 0x29, 0x46, 0x20, 0xcd, 0x1d,
	0x16, 0x26, 0xca, 0x30, 0x96, 0x19, 0x8f, 0x31, 0xb7, 0x5a, 0xc6, 0xd6, 0xd4, 0x4a, 0x0d, 0x05,
	0xce, 0xd6, 0x98, 0x40, 0x66, 0x78, 0x43, 0xe4, 0xbe, 0xe4, 0x16, 0x0b, 0xcc, 0x30, 0xbd, 0xe3,
	0x1a, 0x32, 0x69, 0x15, 0xe6, 0x3d, 0x5a, 0xae, 0x54, 0x02, 0x9f, 0x54, 0x7e, 0xd3, 0x58, 0x7e,
	0xee, 0xc1, 0x32, 0xc7, 0x04, 0x7a, 0xb4, 0xcb, 0xfa, 0x9d, 0x4e, 0x6a, 0x90, 0xbd, 0x2e, 0x6b,
	0xfb, 0x5c, 0xb6, 0x4d, 0xd9, 0xf4, 0x6d, 0x27, 0x5f, 0xeb, 0xad, 0xb2, 0x37, 0xb8, 0xe5, 0x29,
	0xce, 0x3c, 0x38, 0xbb, 0x95, 0x99, 0x4a, 0xa4, 0x45, 0x6d, 0xf8, 0xfe, 0x59, 0xeb, 0xd8, 0x9f,
	0x03, 0x72, 0x72, 0xe9, 0x92, 0xfb, 0xc5, 0x07, 0xf7, 0xbb, 0x82, 0xad, 0x80, 0x5f, 0x1b, 0x30,
	0x96, 0x72, 0xf2, 0x10, 0x17, 0x3f, 0x21, 0xb6, 0x57, 0x49, 0x30, 0x38, 0x1d, 0x9c, 0x1d, 0xce,
	0x9f, 0x56, 0x65, 0x78, 0xb4, 0x44, 0xbd, 0x7e, 0xcf, 0x76, 0x08, 0x13, 0x7b, 0x12, 0x7d, 0x43,
	0x88, 0xca, 0x8d, 0x95, 0x79, 0x0c, 0x57, 0x49, 0x70, 0xe0, 0x25, 0xcf, 0xab, 0x32, 0x7c, 0x52,
	0x4b, 0x5a, 0x8c, 0x89, 0x0e, 0x91, 0x5e, 0x92, 0xc3, 0x5d, 0x54, 0x4d, 0x30, 0x3c, 0x1d, 0x9e,
	0x4d, 0xce, 0x69, 0xb4, 0xdb, 0x31, 0x12, 0x0d, 0x34, 0x7f, 0x56, 0x95, 0xe1, 0x71, 0xed, 0xb4,
	0xa7, 0x33, 0xd1, 0x4a, 0x5d, 0xfb, 0x44, 0x5a, 0xf9, 0x0d, 0x37, 0x3a, 0x86, 0x60, 0x74, 0xbf,
	0x7d, 0x8b, 0x31, 0xd1, 0x21, 0xba, 0xf6, 0xbb, 0x58, 0x9b, 0x60, 0x7c, 0xbf, 0xfd, 0xc7, 0x06,
	0xea, 0xb6, 0xdf, 0xd3, 0x99, 0x68, 0xa5, 0xec, 0xef, 0x80, 0xbc, 0xfc, 0xef, 0x92, 0xa6, 0xc0,
	0xdc, 0xc0, 0x0f, 0x2d, 0x8b, 0x02, 0x34, 0x7d, 0x45, 0x46, 0x31, 0x26, 0xe0, 0xaf, 0x39, 0x9e,
	0x1f, 0x55, 0x65, 0x38, 0xa9, 0x1d, 0xdd, 0x5f, 0x26, 0x3c, 0x48, 0xdf, 0x91, 0x89, 0xfb, 0x7e,
	0xf8, 0x5d, 0x64, 0x52, 0xe5, 0xcd, 0x19, 0x4f, 0xaa, 0x32, 0xa4, 0x2d, 0xb7, 0x01, 0x99, 0xe8,
	0x52, 0xe9, 0x6b, 0x32, 0x06, 0xad, 0x51, 0x07, 0x43, 0xaf, 0x39, 0xae, 0xca, 0xf0, 0x51, 0xad,
	0xf1, 0xbf, 0x99, 0xa8, 0x61, 0x7a, 0x41, 0x46, 0x6e, 0x7f, 0x7f, 0xa2, 0xc9, 0xf9, 0xe3, 0x76,
	0x59, 0x37, 0x73, 0x77, 0x2c, 0xc7, 0x62, 0xc2, 0x93, 0x17, 0x0f, 0x7c, 0x5e, 0x2e, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x54, 0xf2, 0xc6, 0x35, 0xe8, 0x04, 0x00, 0x00,
}
