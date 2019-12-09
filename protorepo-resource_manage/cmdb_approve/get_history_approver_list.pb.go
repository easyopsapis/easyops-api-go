// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_history_approver_list.proto

package cmdb_approve

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//GetHistoryApproverList请求
type GetHistoryApproverListRequest struct {
	//
	//查询条件、类似mongodb
	Query                *types.Struct `protobuf:"bytes,1,opt,name=query,proto3" json:"query" form:"query"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetHistoryApproverListRequest) Reset()         { *m = GetHistoryApproverListRequest{} }
func (m *GetHistoryApproverListRequest) String() string { return proto.CompactTextString(m) }
func (*GetHistoryApproverListRequest) ProtoMessage()    {}
func (*GetHistoryApproverListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ac9cc852a572cb, []int{0}
}
func (m *GetHistoryApproverListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryApproverListRequest.Unmarshal(m, b)
}
func (m *GetHistoryApproverListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryApproverListRequest.Marshal(b, m, deterministic)
}
func (m *GetHistoryApproverListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryApproverListRequest.Merge(m, src)
}
func (m *GetHistoryApproverListRequest) XXX_Size() int {
	return xxx_messageInfo_GetHistoryApproverListRequest.Size(m)
}
func (m *GetHistoryApproverListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryApproverListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryApproverListRequest proto.InternalMessageInfo

func (m *GetHistoryApproverListRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

//
//GetHistoryApproverList返回
type GetHistoryApproverListResponse struct {
	//
	//审批者列表(name)
	UserList             []string `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList" form:"userList"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryApproverListResponse) Reset()         { *m = GetHistoryApproverListResponse{} }
func (m *GetHistoryApproverListResponse) String() string { return proto.CompactTextString(m) }
func (*GetHistoryApproverListResponse) ProtoMessage()    {}
func (*GetHistoryApproverListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ac9cc852a572cb, []int{1}
}
func (m *GetHistoryApproverListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryApproverListResponse.Unmarshal(m, b)
}
func (m *GetHistoryApproverListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryApproverListResponse.Marshal(b, m, deterministic)
}
func (m *GetHistoryApproverListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryApproverListResponse.Merge(m, src)
}
func (m *GetHistoryApproverListResponse) XXX_Size() int {
	return xxx_messageInfo_GetHistoryApproverListResponse.Size(m)
}
func (m *GetHistoryApproverListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryApproverListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryApproverListResponse proto.InternalMessageInfo

func (m *GetHistoryApproverListResponse) GetUserList() []string {
	if m != nil {
		return m.UserList
	}
	return nil
}

//
//GetHistoryApproverListApi返回
type GetHistoryApproverListResponseWrapper struct {
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
	Data                 *GetHistoryApproverListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetHistoryApproverListResponseWrapper) Reset()         { *m = GetHistoryApproverListResponseWrapper{} }
func (m *GetHistoryApproverListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetHistoryApproverListResponseWrapper) ProtoMessage()    {}
func (*GetHistoryApproverListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ac9cc852a572cb, []int{2}
}
func (m *GetHistoryApproverListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryApproverListResponseWrapper.Unmarshal(m, b)
}
func (m *GetHistoryApproverListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryApproverListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetHistoryApproverListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryApproverListResponseWrapper.Merge(m, src)
}
func (m *GetHistoryApproverListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetHistoryApproverListResponseWrapper.Size(m)
}
func (m *GetHistoryApproverListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryApproverListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryApproverListResponseWrapper proto.InternalMessageInfo

func (m *GetHistoryApproverListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetHistoryApproverListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetHistoryApproverListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetHistoryApproverListResponseWrapper) GetData() *GetHistoryApproverListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetHistoryApproverListRequest)(nil), "cmdb_approve.GetHistoryApproverListRequest")
	proto.RegisterType((*GetHistoryApproverListResponse)(nil), "cmdb_approve.GetHistoryApproverListResponse")
	proto.RegisterType((*GetHistoryApproverListResponseWrapper)(nil), "cmdb_approve.GetHistoryApproverListResponseWrapper")
}

func init() { proto.RegisterFile("get_history_approver_list.proto", fileDescriptor_20ac9cc852a572cb) }

var fileDescriptor_20ac9cc852a572cb = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0xc6, 0x95, 0xfe, 0x59, 0x6d, 0xdd, 0x4a, 0x5d, 0x79, 0xa5, 0xdd, 0xa8, 0x02, 0x52, 0x19,
	0x81, 0x7a, 0x00, 0x57, 0x82, 0x0b, 0xe2, 0x82, 0x88, 0x84, 0xe0, 0xc0, 0xa5, 0xe6, 0xc0, 0xb1,
	0x24, 0xa9, 0x9b, 0x46, 0x6a, 0x3b, 0xae, 0xed, 0x20, 0x7a, 0xe4, 0x45, 0xf3, 0x10, 0x79, 0x02,
	0xd4, 0x71, 0x4b, 0x73, 0xa1, 0xa7, 0x64, 0xe6, 0xfb, 0xe6, 0x37, 0xe3, 0x19, 0x12, 0xa4, 0xd2,
	0x8e, 0x67, 0x99, 0xb1, 0xa0, 0xd7, 0xe3, 0x48, 0x29, 0x0d, 0xef, 0x52, 0x8f, 0xe7, 0x99, 0xb1,
	0x5c, 0x69, 0xb0, 0x40, 0x3b, 0xc9, 0x62, 0x12, 0xef, 0x94, 0xde, 0x65, 0x9a, 0xd9, 0x59, 0x1e,
	0xf3, 0x04, 0x16, 0xc3, 0x14, 0x52, 0x18, 0xa2, 0x29, 0xce, 0xa7, 0x18, 0x61, 0x80, 0x7f, 0xae,
	0xb8, 0x77, 0x94, 0x02, 0xa4, 0x73, 0xb9, 0x77, 0x19, 0xab, 0xf3, 0x64, 0x8b, 0x66, 0x6f, 0xe4,
	0xf8, 0x51, 0xda, 0x27, 0xd7, 0xfc, 0x7e, 0xdb, 0xfb, 0x39, 0x33, 0x56, 0xc8, 0x55, 0x2e, 0x8d,
	0xa5, 0x77, 0xa4, 0xb9, 0xca, 0xa5, 0x5e, 0xfb, 0x5e, 0xdf, 0x1b, 0xb4, 0xaf, 0xfe, 0x73, 0x87,
	0xe3, 0x3b, 0x1c, 0x7f, 0x41, 0x5c, 0xf8, 0xa7, 0x2c, 0x82, 0xce, 0x14, 0xf4, 0xe2, 0x96, 0xa1,
	0x9f, 0x09, 0x57, 0xc7, 0x46, 0xe4, 0xe4, 0xa7, 0x0e, 0x46, 0xc1, 0xd2, 0x48, 0x3a, 0x24, 0xbf,
	0x73, 0xe3, 0x72, 0xbe, 0xd7, 0xaf, 0x0f, 0x5a, 0xe1, 0xdf, 0xb2, 0x08, 0xba, 0x0e, 0xb6, 0x53,
	0x98, 0xf8, 0x36, 0xb1, 0xcf, 0x1a, 0x39, 0x3b, 0xcc, 0x7c, 0xd5, 0x91, 0x52, 0x52, 0xd3, 0x53,
	0xd2, 0x48, 0x60, 0x22, 0x71, 0xf8, 0x66, 0xd8, 0x2d, 0x8b, 0xa0, 0xed, 0xb0, 0x9b, 0x2c, 0x13,
	0x28, 0xd2, 0x1b, 0xd2, 0xde, 0x7c, 0x1f, 0x3e, 0xd4, 0x3c, 0xca, 0x96, 0x7e, 0xad, 0xef, 0x0d,
	0x5a, 0xe1, 0xbf, 0xb2, 0x08, 0xe8, 0xde, 0xbb, 0x15, 0x99, 0xa8, 0x5a, 0xe9, 0x39, 0x69, 0x4a,
	0xad, 0x41, 0xfb, 0x75, 0xac, 0xa9, 0xec, 0x00, 0xd3, 0x4c, 0x38, 0x99, 0x8e, 0x48, 0x63, 0x12,
	0xd9, 0xc8, 0x6f, 0xe0, 0x0e, 0x2f, 0x78, 0xf5, 0x9e, 0xfc, 0xf0, 0x4b, 0xaa, 0x43, 0x6f, 0x18,
	0x4c, 0x20, 0x2a, 0xfe, 0x85, 0x07, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xd2, 0xd3,
	0x53, 0x3d, 0x02, 0x00, 0x00,
}
