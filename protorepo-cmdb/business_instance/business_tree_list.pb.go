// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: business_tree_list.proto

package business_instance

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//GetBusinessTreeList请求
type GetBusinessTreeListRequest struct {
	//
	//实例Id列表使用逗号分隔
	Ids                  string   `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids" form:"ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusinessTreeListRequest) Reset()         { *m = GetBusinessTreeListRequest{} }
func (m *GetBusinessTreeListRequest) String() string { return proto.CompactTextString(m) }
func (*GetBusinessTreeListRequest) ProtoMessage()    {}
func (*GetBusinessTreeListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_763f8d780905f746, []int{0}
}
func (m *GetBusinessTreeListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusinessTreeListRequest.Unmarshal(m, b)
}
func (m *GetBusinessTreeListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusinessTreeListRequest.Marshal(b, m, deterministic)
}
func (m *GetBusinessTreeListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusinessTreeListRequest.Merge(m, src)
}
func (m *GetBusinessTreeListRequest) XXX_Size() int {
	return xxx_messageInfo_GetBusinessTreeListRequest.Size(m)
}
func (m *GetBusinessTreeListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusinessTreeListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusinessTreeListRequest proto.InternalMessageInfo

func (m *GetBusinessTreeListRequest) GetIds() string {
	if m != nil {
		return m.Ids
	}
	return ""
}

//
//GetBusinessTreeList返回
type GetBusinessTreeListResponse struct {
	//
	//实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//业务名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//第一个为根节点,最后一个为该业务的父节点, 例如 parents=[根节点...,父节点的父节点, 父节点]
	Parents              []*GetBusinessTreeListResponse_Parents `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents" form:"parents"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *GetBusinessTreeListResponse) Reset()         { *m = GetBusinessTreeListResponse{} }
func (m *GetBusinessTreeListResponse) String() string { return proto.CompactTextString(m) }
func (*GetBusinessTreeListResponse) ProtoMessage()    {}
func (*GetBusinessTreeListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_763f8d780905f746, []int{1}
}
func (m *GetBusinessTreeListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusinessTreeListResponse.Unmarshal(m, b)
}
func (m *GetBusinessTreeListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusinessTreeListResponse.Marshal(b, m, deterministic)
}
func (m *GetBusinessTreeListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusinessTreeListResponse.Merge(m, src)
}
func (m *GetBusinessTreeListResponse) XXX_Size() int {
	return xxx_messageInfo_GetBusinessTreeListResponse.Size(m)
}
func (m *GetBusinessTreeListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusinessTreeListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusinessTreeListResponse proto.InternalMessageInfo

func (m *GetBusinessTreeListResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetBusinessTreeListResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBusinessTreeListResponse) GetParents() []*GetBusinessTreeListResponse_Parents {
	if m != nil {
		return m.Parents
	}
	return nil
}

type GetBusinessTreeListResponse_Parents struct {
	//
	//实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//业务名称
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusinessTreeListResponse_Parents) Reset()         { *m = GetBusinessTreeListResponse_Parents{} }
func (m *GetBusinessTreeListResponse_Parents) String() string { return proto.CompactTextString(m) }
func (*GetBusinessTreeListResponse_Parents) ProtoMessage()    {}
func (*GetBusinessTreeListResponse_Parents) Descriptor() ([]byte, []int) {
	return fileDescriptor_763f8d780905f746, []int{1, 0}
}
func (m *GetBusinessTreeListResponse_Parents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusinessTreeListResponse_Parents.Unmarshal(m, b)
}
func (m *GetBusinessTreeListResponse_Parents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusinessTreeListResponse_Parents.Marshal(b, m, deterministic)
}
func (m *GetBusinessTreeListResponse_Parents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusinessTreeListResponse_Parents.Merge(m, src)
}
func (m *GetBusinessTreeListResponse_Parents) XXX_Size() int {
	return xxx_messageInfo_GetBusinessTreeListResponse_Parents.Size(m)
}
func (m *GetBusinessTreeListResponse_Parents) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusinessTreeListResponse_Parents.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusinessTreeListResponse_Parents proto.InternalMessageInfo

func (m *GetBusinessTreeListResponse_Parents) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetBusinessTreeListResponse_Parents) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//
//GetBusinessTreeListApi返回
type GetBusinessTreeListResponseWrapper struct {
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
	Data                 []*GetBusinessTreeListResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetBusinessTreeListResponseWrapper) Reset()         { *m = GetBusinessTreeListResponseWrapper{} }
func (m *GetBusinessTreeListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetBusinessTreeListResponseWrapper) ProtoMessage()    {}
func (*GetBusinessTreeListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_763f8d780905f746, []int{2}
}
func (m *GetBusinessTreeListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusinessTreeListResponseWrapper.Unmarshal(m, b)
}
func (m *GetBusinessTreeListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusinessTreeListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetBusinessTreeListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusinessTreeListResponseWrapper.Merge(m, src)
}
func (m *GetBusinessTreeListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetBusinessTreeListResponseWrapper.Size(m)
}
func (m *GetBusinessTreeListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusinessTreeListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusinessTreeListResponseWrapper proto.InternalMessageInfo

func (m *GetBusinessTreeListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBusinessTreeListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetBusinessTreeListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetBusinessTreeListResponseWrapper) GetData() []*GetBusinessTreeListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBusinessTreeListRequest)(nil), "business_instance.GetBusinessTreeListRequest")
	proto.RegisterType((*GetBusinessTreeListResponse)(nil), "business_instance.GetBusinessTreeListResponse")
	proto.RegisterType((*GetBusinessTreeListResponse_Parents)(nil), "business_instance.GetBusinessTreeListResponse.Parents")
	proto.RegisterType((*GetBusinessTreeListResponseWrapper)(nil), "business_instance.GetBusinessTreeListResponseWrapper")
}

func init() { proto.RegisterFile("business_tree_list.proto", fileDescriptor_763f8d780905f746) }

var fileDescriptor_763f8d780905f746 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x69, 0xd3, 0xfe, 0xe5, 0xbf, 0x95, 0x6a, 0x07, 0x94, 0x50, 0x17, 0x29, 0x23, 0x48,
	0x37, 0xa6, 0xa0, 0x28, 0xe2, 0xc2, 0x45, 0x41, 0x44, 0x70, 0x21, 0xa3, 0xe0, 0xb2, 0x24, 0xcd,
	0x6d, 0x1d, 0x68, 0x67, 0xe2, 0xcc, 0x04, 0x5c, 0xfa, 0x74, 0x3e, 0x46, 0x1e, 0x22, 0x4f, 0x20,
	0x99, 0x4c, 0x6b, 0x40, 0x29, 0xb8, 0x70, 0x95, 0xcc, 0x3d, 0xe7, 0x7c, 0x73, 0x72, 0x09, 0xf8,
	0x71, 0xa6, 0xb9, 0x40, 0xad, 0xa7, 0x46, 0x21, 0x4e, 0x97, 0x5c, 0x9b, 0x30, 0x55, 0xd2, 0x48,
	0xd2, 0xdf, 0x28, 0x5c, 0x68, 0x13, 0x89, 0x19, 0x0e, 0x4e, 0x16, 0xdc, 0xbc, 0x64, 0x71, 0x38,
	0x93, 0xab, 0xf1, 0x42, 0x2e, 0xe4, 0xd8, 0x3a, 0xe3, 0x6c, 0x6e, 0x4f, 0xf6, 0x60, 0xdf, 0x2a,
	0x02, 0xbd, 0x86, 0xc1, 0x2d, 0x9a, 0x89, 0xc3, 0x3c, 0x29, 0xc4, 0x7b, 0xae, 0x0d, 0xc3, 0xd7,
	0x0c, 0xb5, 0x21, 0x43, 0xf0, 0x78, 0xa2, 0xfd, 0xc6, 0xb0, 0x31, 0xfa, 0x3f, 0xe9, 0x15, 0x79,
	0x00, 0x73, 0xa9, 0x56, 0x57, 0x94, 0x27, 0x9a, 0xb2, 0x52, 0xa2, 0x1f, 0x4d, 0x38, 0xfc, 0x11,
	0xa0, 0x53, 0x29, 0x34, 0x92, 0x73, 0x80, 0x75, 0xb5, 0xbb, 0xc4, 0x81, 0xf6, 0x8b, 0x3c, 0xe8,
	0x3b, 0xd0, 0x46, 0xa3, 0xac, 0x66, 0x24, 0x47, 0xd0, 0x12, 0xd1, 0x0a, 0xfd, 0xa6, 0x0d, 0xec,
	0x16, 0x79, 0xd0, 0xad, 0x02, 0xe5, 0x94, 0x32, 0x2b, 0x92, 0x04, 0x3a, 0x69, 0xa4, 0x50, 0x18,
	0xed, 0x7b, 0x43, 0x6f, 0xd4, 0x3d, 0xbd, 0x08, 0xbf, 0xed, 0x23, 0xdc, 0x52, 0x2e, 0x7c, 0xa8,
	0xd2, 0x13, 0x52, 0xe4, 0x41, 0xaf, 0xe2, 0x3b, 0x20, 0x65, 0x6b, 0xf4, 0x00, 0xa1, 0xe3, 0x7c,
	0x7f, 0xf9, 0x31, 0xf4, 0xbd, 0x09, 0x74, 0x4b, 0xd7, 0x67, 0x15, 0xa5, 0x29, 0xaa, 0x92, 0x35,
	0x93, 0x09, 0xda, 0xcb, 0xdb, 0x75, 0x56, 0x39, 0xa5, 0xcc, 0x8a, 0xe4, 0x12, 0xba, 0xe5, 0xf3,
	0xe6, 0x2d, 0x5d, 0x46, 0x5c, 0xb8, 0x7b, 0x0f, 0x8a, 0x3c, 0x20, 0x5f, 0x5e, 0x27, 0x52, 0x56,
	0xb7, 0x92, 0x63, 0x68, 0xa3, 0x52, 0x52, 0xf9, 0x9e, 0xcd, 0xec, 0x15, 0x79, 0xb0, 0x53, 0x65,
	0xec, 0x98, 0xb2, 0x4a, 0x26, 0x8f, 0xd0, 0x4a, 0x22, 0x13, 0xf9, 0x2d, 0xbb, 0xf7, 0xf0, 0x77,
	0x7b, 0xaf, 0xd7, 0x2e, 0x29, 0x94, 0x59, 0x58, 0xfc, 0xcf, 0xfe, 0x92, 0x67, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0xf2, 0x50, 0xf8, 0xf0, 0x02, 0x00, 0x00,
}