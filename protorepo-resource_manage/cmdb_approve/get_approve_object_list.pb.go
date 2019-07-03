// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_approve_object_list.proto

package cmdb_approve

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//GetApproveObjectList请求
type GetApproveObjectListRequest struct {
	//
	//查询条件、类似mongodb
	Query *types.Struct `protobuf:"bytes,1,opt,name=query,proto3" json:"query" form:"query"`
	//
	//获取待审批的或我提交的记录
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetApproveObjectListRequest) Reset()         { *m = GetApproveObjectListRequest{} }
func (m *GetApproveObjectListRequest) String() string { return proto.CompactTextString(m) }
func (*GetApproveObjectListRequest) ProtoMessage()    {}
func (*GetApproveObjectListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b271186ec7493141, []int{0}
}
func (m *GetApproveObjectListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApproveObjectListRequest.Unmarshal(m, b)
}
func (m *GetApproveObjectListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApproveObjectListRequest.Marshal(b, m, deterministic)
}
func (m *GetApproveObjectListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApproveObjectListRequest.Merge(m, src)
}
func (m *GetApproveObjectListRequest) XXX_Size() int {
	return xxx_messageInfo_GetApproveObjectListRequest.Size(m)
}
func (m *GetApproveObjectListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApproveObjectListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetApproveObjectListRequest proto.InternalMessageInfo

func (m *GetApproveObjectListRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *GetApproveObjectListRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

//
//GetApproveObjectList返回
type GetApproveObjectListResponse struct {
	//
	//模型id列表
	ObjectIdList         []string `protobuf:"bytes,1,rep,name=objectIdList,proto3" json:"objectIdList" form:"objectIdList"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetApproveObjectListResponse) Reset()         { *m = GetApproveObjectListResponse{} }
func (m *GetApproveObjectListResponse) String() string { return proto.CompactTextString(m) }
func (*GetApproveObjectListResponse) ProtoMessage()    {}
func (*GetApproveObjectListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b271186ec7493141, []int{1}
}
func (m *GetApproveObjectListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApproveObjectListResponse.Unmarshal(m, b)
}
func (m *GetApproveObjectListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApproveObjectListResponse.Marshal(b, m, deterministic)
}
func (m *GetApproveObjectListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApproveObjectListResponse.Merge(m, src)
}
func (m *GetApproveObjectListResponse) XXX_Size() int {
	return xxx_messageInfo_GetApproveObjectListResponse.Size(m)
}
func (m *GetApproveObjectListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApproveObjectListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetApproveObjectListResponse proto.InternalMessageInfo

func (m *GetApproveObjectListResponse) GetObjectIdList() []string {
	if m != nil {
		return m.ObjectIdList
	}
	return nil
}

//
//GetApproveObjectListApi返回
type GetApproveObjectListResponseWrapper struct {
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
	Data                 *GetApproveObjectListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetApproveObjectListResponseWrapper) Reset()         { *m = GetApproveObjectListResponseWrapper{} }
func (m *GetApproveObjectListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetApproveObjectListResponseWrapper) ProtoMessage()    {}
func (*GetApproveObjectListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b271186ec7493141, []int{2}
}
func (m *GetApproveObjectListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApproveObjectListResponseWrapper.Unmarshal(m, b)
}
func (m *GetApproveObjectListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApproveObjectListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetApproveObjectListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApproveObjectListResponseWrapper.Merge(m, src)
}
func (m *GetApproveObjectListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetApproveObjectListResponseWrapper.Size(m)
}
func (m *GetApproveObjectListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApproveObjectListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetApproveObjectListResponseWrapper proto.InternalMessageInfo

func (m *GetApproveObjectListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetApproveObjectListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetApproveObjectListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetApproveObjectListResponseWrapper) GetData() *GetApproveObjectListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetApproveObjectListRequest)(nil), "cmdb_approve.GetApproveObjectListRequest")
	proto.RegisterType((*GetApproveObjectListResponse)(nil), "cmdb_approve.GetApproveObjectListResponse")
	proto.RegisterType((*GetApproveObjectListResponseWrapper)(nil), "cmdb_approve.GetApproveObjectListResponseWrapper")
}

func init() { proto.RegisterFile("get_approve_object_list.proto", fileDescriptor_b271186ec7493141) }

var fileDescriptor_b271186ec7493141 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x25, 0x76, 0x57, 0xe8, 0xec, 0x82, 0x3a, 0x05, 0x0d, 0x6b, 0x25, 0xcb, 0x54, 0x64, 0x11,
	0x33, 0xc1, 0x16, 0x44, 0x7c, 0x91, 0x2e, 0x88, 0x08, 0x42, 0x61, 0x7c, 0x10, 0x14, 0x0d, 0x93,
	0x64, 0x1a, 0xa3, 0x49, 0xef, 0x74, 0x32, 0x69, 0xad, 0xe2, 0x93, 0xff, 0x19, 0xc1, 0x4f, 0xc8,
	0xbb, 0x20, 0xb9, 0xb3, 0x75, 0x23, 0x94, 0x7d, 0xca, 0xcc, 0x3d, 0xe7, 0x9e, 0x9c, 0x73, 0xef,
	0x90, 0x7b, 0xb9, 0xb2, 0xb1, 0xd4, 0xda, 0xc0, 0x99, 0x8a, 0x21, 0xf9, 0xac, 0x52, 0x1b, 0x97,
	0x45, 0x6d, 0xb9, 0x36, 0x60, 0x81, 0x4e, 0xd3, 0x2a, 0x4b, 0x2e, 0xf1, 0x59, 0x98, 0x17, 0xf6,
	0x53, 0x93, 0xf0, 0x14, 0xaa, 0x28, 0x87, 0x1c, 0x22, 0x24, 0x25, 0xcd, 0x31, 0xde, 0xf0, 0x82,
	0x27, 0xd7, 0x3c, 0xdb, 0xcd, 0x01, 0xf2, 0x52, 0xad, 0x59, 0xb5, 0x35, 0x4d, 0xba, 0x92, 0x9e,
	0x3d, 0x19, 0x88, 0x55, 0xe7, 0x85, 0xfd, 0x02, 0xe7, 0x51, 0x0e, 0x21, 0x82, 0xe1, 0x99, 0x2c,
	0x8b, 0x4c, 0x5a, 0x30, 0x75, 0xf4, 0xef, 0xe8, 0xfa, 0xd8, 0x4f, 0x8f, 0xdc, 0x7d, 0xa9, 0xec,
	0xa1, 0xf3, 0x74, 0x84, 0x96, 0x5f, 0x17, 0xb5, 0x15, 0xea, 0xb4, 0x51, 0xb5, 0xa5, 0xcf, 0xc9,
	0xf8, 0xb4, 0x51, 0xe6, 0xc2, 0xf7, 0xe6, 0xde, 0x62, 0xb2, 0x7f, 0x87, 0x3b, 0x17, 0xfc, 0xd2,
	0x05, 0x7f, 0x83, 0x2e, 0x96, 0x37, 0xbb, 0x36, 0x98, 0x1e, 0x83, 0xa9, 0x9e, 0x31, 0xe4, 0x33,
	0xe1, 0xfa, 0xe8, 0x1e, 0x19, 0xd9, 0x0b, 0xad, 0xfc, 0x6b, 0x73, 0x6f, 0xb1, 0xbd, 0xbc, 0xd1,
	0xb5, 0xc1, 0xc4, 0xd1, 0xfa, 0x2a, 0x13, 0x08, 0x32, 0x43, 0x76, 0xaf, 0x36, 0x51, 0x6b, 0x38,
	0xa9, 0x15, 0x15, 0x64, 0xea, 0xa6, 0xf9, 0x2a, 0xeb, 0xeb, 0xbe, 0x37, 0xdf, 0x5a, 0x6c, 0x2f,
	0x79, 0xd7, 0x06, 0x3b, 0x4e, 0x6c, 0x88, 0xb2, 0xdf, 0xbf, 0x82, 0x1d, 0x72, 0xeb, 0xe3, 0x7b,
	0x19, 0x7e, 0x3b, 0x0c, 0xdf, 0xc5, 0x1f, 0xbe, 0x3f, 0x7e, 0x74, 0xb0, 0xff, 0xe3, 0xbe, 0xf8,
	0x4f, 0x83, 0xfd, 0xf1, 0xc8, 0xde, 0xa6, 0x9f, 0xbe, 0x35, 0x52, 0x6b, 0x65, 0xfa, 0x00, 0x29,
	0x64, 0x0a, 0x07, 0x30, 0x1e, 0x06, 0xe8, 0xab, 0x4c, 0x20, 0x48, 0x9f, 0x92, 0x49, 0xff, 0x7d,
	0xf1, 0x55, 0x97, 0xb2, 0x38, 0x59, 0x85, 0xbd, 0xdd, 0xb5, 0x01, 0x5d, 0x73, 0x57, 0x20, 0x13,
	0x43, 0x2a, 0x7d, 0x40, 0xc6, 0xca, 0x18, 0x30, 0xfe, 0x16, 0xf6, 0x0c, 0xe6, 0x88, 0x65, 0x26,
	0x1c, 0x4c, 0x8f, 0xc8, 0x28, 0x93, 0x56, 0xfa, 0x23, 0xdc, 0xc3, 0x43, 0x3e, 0x7c, 0x4a, 0x7c,
	0x53, 0x8e, 0xa1, 0xe5, 0x5e, 0x81, 0x09, 0x14, 0x4a, 0xae, 0xe3, 0x0a, 0x0f, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x94, 0x85, 0x32, 0x7a, 0xb4, 0x02, 0x00, 0x00,
}