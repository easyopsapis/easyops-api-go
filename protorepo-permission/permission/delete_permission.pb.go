// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_permission.proto

package permission

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//DeletePermission请求
type DeletePermissionRequest struct {
	//
	//需要鉴权的动作, 全局唯一
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action" form:"action"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePermissionRequest) Reset()         { *m = DeletePermissionRequest{} }
func (m *DeletePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionRequest) ProtoMessage()    {}
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffc91ce64ee8bb4, []int{0}
}
func (m *DeletePermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionRequest.Unmarshal(m, b)
}
func (m *DeletePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionRequest.Marshal(b, m, deterministic)
}
func (m *DeletePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionRequest.Merge(m, src)
}
func (m *DeletePermissionRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionRequest.Size(m)
}
func (m *DeletePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionRequest proto.InternalMessageInfo

func (m *DeletePermissionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

//
//DeletePermission返回
type DeletePermissionResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//结果信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//数据
	Data                 *DeletePermissionResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DeletePermissionResponse) Reset()         { *m = DeletePermissionResponse{} }
func (m *DeletePermissionResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionResponse) ProtoMessage()    {}
func (*DeletePermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffc91ce64ee8bb4, []int{1}
}
func (m *DeletePermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionResponse.Unmarshal(m, b)
}
func (m *DeletePermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionResponse.Marshal(b, m, deterministic)
}
func (m *DeletePermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionResponse.Merge(m, src)
}
func (m *DeletePermissionResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionResponse.Size(m)
}
func (m *DeletePermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionResponse proto.InternalMessageInfo

func (m *DeletePermissionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeletePermissionResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DeletePermissionResponse) GetData() *DeletePermissionResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeletePermissionResponse_Data struct {
	//
	//删除权限数量
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count" form:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePermissionResponse_Data) Reset()         { *m = DeletePermissionResponse_Data{} }
func (m *DeletePermissionResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionResponse_Data) ProtoMessage()    {}
func (*DeletePermissionResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffc91ce64ee8bb4, []int{1, 0}
}
func (m *DeletePermissionResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionResponse_Data.Unmarshal(m, b)
}
func (m *DeletePermissionResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionResponse_Data.Marshal(b, m, deterministic)
}
func (m *DeletePermissionResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionResponse_Data.Merge(m, src)
}
func (m *DeletePermissionResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionResponse_Data.Size(m)
}
func (m *DeletePermissionResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionResponse_Data proto.InternalMessageInfo

func (m *DeletePermissionResponse_Data) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

//
//DeletePermissionApi返回
type DeletePermissionResponseWrapper struct {
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
	Data                 *DeletePermissionResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DeletePermissionResponseWrapper) Reset()         { *m = DeletePermissionResponseWrapper{} }
func (m *DeletePermissionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionResponseWrapper) ProtoMessage()    {}
func (*DeletePermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ffc91ce64ee8bb4, []int{2}
}
func (m *DeletePermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionResponseWrapper.Unmarshal(m, b)
}
func (m *DeletePermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeletePermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionResponseWrapper.Merge(m, src)
}
func (m *DeletePermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionResponseWrapper.Size(m)
}
func (m *DeletePermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionResponseWrapper proto.InternalMessageInfo

func (m *DeletePermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeletePermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeletePermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeletePermissionResponseWrapper) GetData() *DeletePermissionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeletePermissionRequest)(nil), "permission.DeletePermissionRequest")
	proto.RegisterType((*DeletePermissionResponse)(nil), "permission.DeletePermissionResponse")
	proto.RegisterType((*DeletePermissionResponse_Data)(nil), "permission.DeletePermissionResponse.Data")
	proto.RegisterType((*DeletePermissionResponseWrapper)(nil), "permission.DeletePermissionResponseWrapper")
}

func init() { proto.RegisterFile("delete_permission.proto", fileDescriptor_2ffc91ce64ee8bb4) }

var fileDescriptor_2ffc91ce64ee8bb4 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4a, 0xfb, 0x40,
	0x18, 0x24, 0xbf, 0xa6, 0x85, 0x7e, 0xfd, 0xf9, 0x6f, 0x0f, 0x36, 0xf4, 0x92, 0xb2, 0x8a, 0xb4,
	0x07, 0x53, 0xd0, 0x8b, 0x78, 0x2c, 0xf5, 0xe0, 0x45, 0x64, 0x2f, 0x1e, 0x65, 0x9b, 0x6e, 0x63,
	0xa0, 0xc9, 0x17, 0x77, 0x37, 0xe0, 0xd3, 0x06, 0x7c, 0x85, 0xf8, 0x02, 0x92, 0x2f, 0xd1, 0x06,
	0x24, 0xd0, 0x53, 0xb2, 0x33, 0x93, 0x99, 0xd9, 0x21, 0x30, 0xde, 0xa8, 0x9d, 0xb2, 0xea, 0x35,
	0x53, 0x3a, 0x89, 0x8d, 0x89, 0x31, 0x0d, 0x32, 0x8d, 0x16, 0x19, 0xec, 0x91, 0xc9, 0x75, 0x14,
	0xdb, 0xb7, 0x7c, 0x1d, 0x84, 0x98, 0x2c, 0x22, 0x8c, 0x70, 0x41, 0x92, 0x75, 0xbe, 0xa5, 0x13,
	0x1d, 0xe8, 0xad, 0xfe, 0x94, 0xaf, 0x60, 0xbc, 0x22, 0xd7, 0xe7, 0x5f, 0x0b, 0xa1, 0xde, 0x73,
	0x65, 0x2c, 0x9b, 0xc3, 0x40, 0x86, 0x36, 0xc6, 0xd4, 0x73, 0xa6, 0xce, 0x6c, 0xb8, 0x3c, 0x2b,
	0x0b, 0xff, 0x68, 0x8b, 0x3a, 0xb9, 0xe7, 0x35, 0xce, 0x45, 0x23, 0xe0, 0x9f, 0x0e, 0x78, 0x7f,
	0x6d, 0x4c, 0x86, 0xa9, 0x51, 0xec, 0x02, 0xdc, 0x10, 0x37, 0x8a, 0x5c, 0xfa, 0xcb, 0x93, 0xb2,
	0xf0, 0x47, 0xb5, 0x4b, 0x85, 0x72, 0x41, 0x24, 0x9b, 0x42, 0x2f, 0x31, 0x91, 0xf7, 0x8f, 0x92,
	0x8e, 0xcb, 0xc2, 0x87, 0x5a, 0x93, 0x98, 0x88, 0x8b, 0x8a, 0x62, 0x4f, 0xe0, 0x6e, 0xa4, 0x95,
	0x5e, 0x6f, 0xea, 0xcc, 0x46, 0x37, 0xf3, 0xa0, 0xb5, 0x42, 0x57, 0x74, 0xb0, 0x92, 0x56, 0xb6,
	0x13, 0x2b, 0x03, 0x2e, 0xc8, 0x67, 0x12, 0x80, 0x5b, 0xd1, 0xec, 0x0a, 0xfa, 0x21, 0xe6, 0xa9,
	0x6d, 0xfa, 0x9d, 0x96, 0x85, 0xff, 0xff, 0xa7, 0x5f, 0x9e, 0x5a, 0x2e, 0x6a, 0x9a, 0x7f, 0x39,
	0xe0, 0x77, 0x05, 0xbd, 0x68, 0x99, 0x65, 0x4a, 0x1f, 0x76, 0xd5, 0x3b, 0x18, 0x55, 0xcf, 0x87,
	0x8f, 0x6c, 0x27, 0xe3, 0xb4, 0xb9, 0xf2, 0x79, 0x59, 0xf8, 0x6c, 0xaf, 0x6d, 0x48, 0x2e, 0xda,
	0xd2, 0xaa, 0xaa, 0xd2, 0x1a, 0x35, 0x6d, 0x30, 0x6c, 0x57, 0x25, 0x98, 0x8b, 0x9a, 0x66, 0x8f,
	0xcd, 0x54, 0x2e, 0x4d, 0x75, 0x79, 0xc8, 0x54, 0x1d, 0x2b, 0xad, 0x07, 0xf4, 0x9b, 0xdc, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc1, 0xe3, 0x09, 0x7c, 0x02, 0x00, 0x00,
}
