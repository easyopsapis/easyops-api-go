// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregate_micro_app.proto

package object_micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//AggregateMicroApp请求
type AggregateMicroAppRequest struct {
	//
	//查询条件
	Query                *AggregateMicroAppRequest_Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query" form:"query"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AggregateMicroAppRequest) Reset()         { *m = AggregateMicroAppRequest{} }
func (m *AggregateMicroAppRequest) String() string { return proto.CompactTextString(m) }
func (*AggregateMicroAppRequest) ProtoMessage()    {}
func (*AggregateMicroAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1de8679a28a1eb5, []int{0}
}
func (m *AggregateMicroAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateMicroAppRequest.Unmarshal(m, b)
}
func (m *AggregateMicroAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateMicroAppRequest.Marshal(b, m, deterministic)
}
func (m *AggregateMicroAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateMicroAppRequest.Merge(m, src)
}
func (m *AggregateMicroAppRequest) XXX_Size() int {
	return xxx_messageInfo_AggregateMicroAppRequest.Size(m)
}
func (m *AggregateMicroAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateMicroAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateMicroAppRequest proto.InternalMessageInfo

func (m *AggregateMicroAppRequest) GetQuery() *AggregateMicroAppRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type AggregateMicroAppRequest_Query struct {
	//
	//按objectId列表查询
	ObjectId             []string `protobuf:"bytes,1,rep,name=objectId,proto3" json:"objectId" form:"objectId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregateMicroAppRequest_Query) Reset()         { *m = AggregateMicroAppRequest_Query{} }
func (m *AggregateMicroAppRequest_Query) String() string { return proto.CompactTextString(m) }
func (*AggregateMicroAppRequest_Query) ProtoMessage()    {}
func (*AggregateMicroAppRequest_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1de8679a28a1eb5, []int{0, 0}
}
func (m *AggregateMicroAppRequest_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateMicroAppRequest_Query.Unmarshal(m, b)
}
func (m *AggregateMicroAppRequest_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateMicroAppRequest_Query.Marshal(b, m, deterministic)
}
func (m *AggregateMicroAppRequest_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateMicroAppRequest_Query.Merge(m, src)
}
func (m *AggregateMicroAppRequest_Query) XXX_Size() int {
	return xxx_messageInfo_AggregateMicroAppRequest_Query.Size(m)
}
func (m *AggregateMicroAppRequest_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateMicroAppRequest_Query.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateMicroAppRequest_Query proto.InternalMessageInfo

func (m *AggregateMicroAppRequest_Query) GetObjectId() []string {
	if m != nil {
		return m.ObjectId
	}
	return nil
}

//
//AggregateMicroApp返回
type AggregateMicroAppResponse struct {
	//
	//数据列表
	List                 []*AggregateMicroAppResponse_List `protobuf:"bytes,1,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AggregateMicroAppResponse) Reset()         { *m = AggregateMicroAppResponse{} }
func (m *AggregateMicroAppResponse) String() string { return proto.CompactTextString(m) }
func (*AggregateMicroAppResponse) ProtoMessage()    {}
func (*AggregateMicroAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1de8679a28a1eb5, []int{1}
}
func (m *AggregateMicroAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateMicroAppResponse.Unmarshal(m, b)
}
func (m *AggregateMicroAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateMicroAppResponse.Marshal(b, m, deterministic)
}
func (m *AggregateMicroAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateMicroAppResponse.Merge(m, src)
}
func (m *AggregateMicroAppResponse) XXX_Size() int {
	return xxx_messageInfo_AggregateMicroAppResponse.Size(m)
}
func (m *AggregateMicroAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateMicroAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateMicroAppResponse proto.InternalMessageInfo

func (m *AggregateMicroAppResponse) GetList() []*AggregateMicroAppResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type AggregateMicroAppResponse_List struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//关联小产品列表
	MicroAppList         []*micro_app.ObjectMicroApp `protobuf:"bytes,2,rep,name=microAppList,proto3" json:"microAppList" form:"microAppList"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AggregateMicroAppResponse_List) Reset()         { *m = AggregateMicroAppResponse_List{} }
func (m *AggregateMicroAppResponse_List) String() string { return proto.CompactTextString(m) }
func (*AggregateMicroAppResponse_List) ProtoMessage()    {}
func (*AggregateMicroAppResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1de8679a28a1eb5, []int{1, 0}
}
func (m *AggregateMicroAppResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateMicroAppResponse_List.Unmarshal(m, b)
}
func (m *AggregateMicroAppResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateMicroAppResponse_List.Marshal(b, m, deterministic)
}
func (m *AggregateMicroAppResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateMicroAppResponse_List.Merge(m, src)
}
func (m *AggregateMicroAppResponse_List) XXX_Size() int {
	return xxx_messageInfo_AggregateMicroAppResponse_List.Size(m)
}
func (m *AggregateMicroAppResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateMicroAppResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateMicroAppResponse_List proto.InternalMessageInfo

func (m *AggregateMicroAppResponse_List) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AggregateMicroAppResponse_List) GetMicroAppList() []*micro_app.ObjectMicroApp {
	if m != nil {
		return m.MicroAppList
	}
	return nil
}

//
//AggregateMicroAppApi返回
type AggregateMicroAppResponseWrapper struct {
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
	Data                 *AggregateMicroAppResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AggregateMicroAppResponseWrapper) Reset()         { *m = AggregateMicroAppResponseWrapper{} }
func (m *AggregateMicroAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AggregateMicroAppResponseWrapper) ProtoMessage()    {}
func (*AggregateMicroAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1de8679a28a1eb5, []int{2}
}
func (m *AggregateMicroAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateMicroAppResponseWrapper.Unmarshal(m, b)
}
func (m *AggregateMicroAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateMicroAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AggregateMicroAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateMicroAppResponseWrapper.Merge(m, src)
}
func (m *AggregateMicroAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AggregateMicroAppResponseWrapper.Size(m)
}
func (m *AggregateMicroAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateMicroAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateMicroAppResponseWrapper proto.InternalMessageInfo

func (m *AggregateMicroAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AggregateMicroAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AggregateMicroAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AggregateMicroAppResponseWrapper) GetData() *AggregateMicroAppResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AggregateMicroAppRequest)(nil), "object_micro_app.AggregateMicroAppRequest")
	proto.RegisterType((*AggregateMicroAppRequest_Query)(nil), "object_micro_app.AggregateMicroAppRequest.Query")
	proto.RegisterType((*AggregateMicroAppResponse)(nil), "object_micro_app.AggregateMicroAppResponse")
	proto.RegisterType((*AggregateMicroAppResponse_List)(nil), "object_micro_app.AggregateMicroAppResponse.List")
	proto.RegisterType((*AggregateMicroAppResponseWrapper)(nil), "object_micro_app.AggregateMicroAppResponseWrapper")
}

func init() { proto.RegisterFile("aggregate_micro_app.proto", fileDescriptor_e1de8679a28a1eb5) }

var fileDescriptor_e1de8679a28a1eb5 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x99, 0xed, 0xac, 0x68, 0xb6, 0xd0, 0x92, 0x82, 0xce, 0xee, 0x65, 0x96, 0x08, 0x52,
	0x90, 0x26, 0xa2, 0x97, 0xe2, 0xad, 0x0b, 0x1e, 0x04, 0x45, 0x0d, 0xf8, 0x07, 0x2f, 0x25, 0x3b,
	0x9b, 0x8e, 0x23, 0x33, 0xfb, 0xa6, 0x49, 0x16, 0xec, 0xdd, 0xbb, 0xdf, 0xc4, 0x8f, 0x34, 0x1f,
	0x62, 0xae, 0x5e, 0x24, 0x6f, 0x76, 0xbb, 0x63, 0x75, 0x51, 0x4f, 0x33, 0xc9, 0xf3, 0x3c, 0xbf,
	0xf7, 0x09, 0x09, 0x19, 0xab, 0xb2, 0xb4, 0xba, 0x54, 0x5e, 0x9f, 0x37, 0x55, 0x61, 0xe1, 0x5c,
	0x19, 0xc3, 0x8d, 0x05, 0x0f, 0xf4, 0x10, 0xe6, 0x9f, 0x75, 0xe1, 0xb7, 0xfb, 0x93, 0x93, 0xb2,
	0xf2, 0x9f, 0x56, 0x73, 0x5e, 0x40, 0x23, 0x4a, 0x28, 0x41, 0xa0, 0x71, 0xbe, 0xba, 0xc0, 0x15,
	0x2e, 0xf0, 0x2f, 0x02, 0x26, 0x1f, 0x4b, 0xe0, 0x5a, 0xb9, 0x2b, 0x30, 0x8e, 0xd7, 0x50, 0xa8,
	0x5a, 0x14, 0xb0, 0xf4, 0x56, 0x15, 0xde, 0xc5, 0xa4, 0xd5, 0x06, 0x4e, 0x1a, 0x58, 0xe8, 0xda,
	0x89, 0xb5, 0x51, 0xe0, 0x52, 0x5c, 0xcf, 0x14, 0x37, 0x4b, 0x44, 0x36, 0xfb, 0x9e, 0x90, 0xec,
	0x6c, 0x53, 0xfd, 0x65, 0x10, 0xcf, 0x8c, 0x91, 0xfa, 0x72, 0xa5, 0x9d, 0xa7, 0x1f, 0xc8, 0xf0,
	0x72, 0xa5, 0xed, 0x55, 0x96, 0x4c, 0x93, 0xe3, 0xd1, 0xe3, 0x47, 0xfc, 0x37, 0xc8, 0xae, 0x28,
	0x7f, 0x13, 0x72, 0xb3, 0xc3, 0xae, 0xcd, 0xf7, 0x2f, 0xc0, 0x36, 0x4f, 0x19, 0x82, 0x98, 0x8c,
	0xc0, 0xc9, 0x29, 0x19, 0xa2, 0x83, 0x0a, 0x72, 0x3b, 0x42, 0x9f, 0x2f, 0xb2, 0x64, 0xba, 0x77,
	0x7c, 0x67, 0x76, 0xd4, 0xb5, 0xf9, 0x41, 0xcc, 0x6c, 0x14, 0x26, 0xaf, 0x4d, 0xec, 0xeb, 0x80,
	0x8c, 0xff, 0x30, 0xd5, 0x19, 0x58, 0x3a, 0x4d, 0xdf, 0x92, 0xb4, 0xae, 0x9c, 0x47, 0xd4, 0xbf,
	0x16, 0x8e, 0x51, 0xfe, 0xa2, 0x72, 0x7e, 0x76, 0xd0, 0xb5, 0xf9, 0x28, 0x0e, 0x0f, 0x1c, 0x26,
	0x11, 0x37, 0xf9, 0x96, 0x90, 0x34, 0xe8, 0x37, 0xea, 0x26, 0x7f, 0xad, 0x4b, 0xdf, 0x91, 0xfd,
	0x66, 0x3d, 0x29, 0x00, 0xb2, 0x01, 0x16, 0x1b, 0xf3, 0x6d, 0xa3, 0x57, 0x68, 0xdd, 0xd4, 0x99,
	0xdd, 0xeb, 0xda, 0xfc, 0x28, 0xf2, 0xfa, 0x41, 0x26, 0x7f, 0xe1, 0xb0, 0x1f, 0x09, 0x99, 0xee,
	0x3c, 0xcb, 0x7b, 0xab, 0x8c, 0xd1, 0x96, 0xde, 0x27, 0x69, 0x01, 0x0b, 0x8d, 0x4d, 0x87, 0xfd,
	0xb3, 0x85, 0x5d, 0x26, 0x51, 0xa4, 0xa7, 0x64, 0x14, 0xbe, 0xcf, 0xbe, 0x98, 0x5a, 0x55, 0xcb,
	0x6c, 0x80, 0xa7, 0xba, 0xdb, 0xb5, 0x39, 0xdd, 0x7a, 0xd7, 0x22, 0x93, 0x7d, 0x2b, 0x7d, 0x40,
	0x86, 0xda, 0x5a, 0xb0, 0xd9, 0x1e, 0x66, 0x7a, 0x97, 0x8d, 0xdb, 0x4c, 0x46, 0x99, 0xbe, 0x26,
	0xe9, 0x42, 0x79, 0x95, 0xa5, 0xf8, 0x8a, 0x1e, 0xfe, 0xc7, 0xa5, 0xf4, 0x3b, 0x07, 0x04, 0x93,
	0x48, 0x9a, 0xdf, 0xc2, 0xc7, 0xfb, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x6a, 0xdd,
	0xa3, 0x76, 0x03, 0x00, 0x00,
}
