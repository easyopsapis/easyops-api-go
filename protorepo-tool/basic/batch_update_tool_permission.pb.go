// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_update_tool_permission.proto

package basic

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
//BatchUpdateToolPermission请求
type BatchUpdateToolPermissionRequest struct {
	//
	//批量更新方式, 可选值: append(追加), overwrite(重置), pull(移除)
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method" form:"method"`
	//
	//权限列表, 可选值: deleteAuthorizers(删除), updateAuthorizers(更新), readAuthorizers(读取或查看),executeAuthorizers,rootExecuteAuthorizers,readExecutionResultAuthorizers
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions" form:"permissions"`
	//
	//用户或用户组列表
	Authorizers []string `protobuf:"bytes,3,rep,name=authorizers,proto3" json:"authorizers" form:"authorizers"`
	//
	//工具ID列表
	ToolIds              []string `protobuf:"bytes,4,rep,name=toolIds,proto3" json:"toolIds" form:"toolIds"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchUpdateToolPermissionRequest) Reset()         { *m = BatchUpdateToolPermissionRequest{} }
func (m *BatchUpdateToolPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*BatchUpdateToolPermissionRequest) ProtoMessage()    {}
func (*BatchUpdateToolPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b50f8a323b6f559, []int{0}
}
func (m *BatchUpdateToolPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdateToolPermissionRequest.Unmarshal(m, b)
}
func (m *BatchUpdateToolPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdateToolPermissionRequest.Marshal(b, m, deterministic)
}
func (m *BatchUpdateToolPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdateToolPermissionRequest.Merge(m, src)
}
func (m *BatchUpdateToolPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_BatchUpdateToolPermissionRequest.Size(m)
}
func (m *BatchUpdateToolPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdateToolPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdateToolPermissionRequest proto.InternalMessageInfo

func (m *BatchUpdateToolPermissionRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *BatchUpdateToolPermissionRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *BatchUpdateToolPermissionRequest) GetAuthorizers() []string {
	if m != nil {
		return m.Authorizers
	}
	return nil
}

func (m *BatchUpdateToolPermissionRequest) GetToolIds() []string {
	if m != nil {
		return m.ToolIds
	}
	return nil
}

//
//BatchUpdateToolPermission返回
type BatchUpdateToolPermissionResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//工具修改详情,返回结果如下 { "code": 0, "error": "成功: Success", "message": "Success", "data": { "999dab4e776c6378bd8c15ff5730cbd7": { "code": 0, "message": "success", "name": "cccc" } } }
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BatchUpdateToolPermissionResponse) Reset()         { *m = BatchUpdateToolPermissionResponse{} }
func (m *BatchUpdateToolPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*BatchUpdateToolPermissionResponse) ProtoMessage()    {}
func (*BatchUpdateToolPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b50f8a323b6f559, []int{1}
}
func (m *BatchUpdateToolPermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdateToolPermissionResponse.Unmarshal(m, b)
}
func (m *BatchUpdateToolPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdateToolPermissionResponse.Marshal(b, m, deterministic)
}
func (m *BatchUpdateToolPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdateToolPermissionResponse.Merge(m, src)
}
func (m *BatchUpdateToolPermissionResponse) XXX_Size() int {
	return xxx_messageInfo_BatchUpdateToolPermissionResponse.Size(m)
}
func (m *BatchUpdateToolPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdateToolPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdateToolPermissionResponse proto.InternalMessageInfo

func (m *BatchUpdateToolPermissionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchUpdateToolPermissionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BatchUpdateToolPermissionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BatchUpdateToolPermissionResponse) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//BatchUpdateToolPermissionApi返回
type BatchUpdateToolPermissionResponseWrapper struct {
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
	Data                 *BatchUpdateToolPermissionResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BatchUpdateToolPermissionResponseWrapper) Reset() {
	*m = BatchUpdateToolPermissionResponseWrapper{}
}
func (m *BatchUpdateToolPermissionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*BatchUpdateToolPermissionResponseWrapper) ProtoMessage()    {}
func (*BatchUpdateToolPermissionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b50f8a323b6f559, []int{2}
}
func (m *BatchUpdateToolPermissionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper.Unmarshal(m, b)
}
func (m *BatchUpdateToolPermissionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *BatchUpdateToolPermissionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper.Merge(m, src)
}
func (m *BatchUpdateToolPermissionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper.Size(m)
}
func (m *BatchUpdateToolPermissionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdateToolPermissionResponseWrapper proto.InternalMessageInfo

func (m *BatchUpdateToolPermissionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchUpdateToolPermissionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *BatchUpdateToolPermissionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BatchUpdateToolPermissionResponseWrapper) GetData() *BatchUpdateToolPermissionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchUpdateToolPermissionRequest)(nil), "basic.BatchUpdateToolPermissionRequest")
	proto.RegisterType((*BatchUpdateToolPermissionResponse)(nil), "basic.BatchUpdateToolPermissionResponse")
	proto.RegisterType((*BatchUpdateToolPermissionResponseWrapper)(nil), "basic.BatchUpdateToolPermissionResponseWrapper")
}

func init() { proto.RegisterFile("batch_update_tool_permission.proto", fileDescriptor_8b50f8a323b6f559) }

var fileDescriptor_8b50f8a323b6f559 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x8a, 0xd4, 0x30,
	0x18, 0x65, 0xfe, 0x76, 0xd9, 0x8c, 0x7f, 0x1b, 0x41, 0xcb, 0x22, 0x74, 0x8c, 0x22, 0x5d, 0xb0,
	0x1d, 0xd9, 0x05, 0x59, 0xc5, 0x1b, 0x47, 0x14, 0xbc, 0x10, 0x24, 0x2a, 0x5e, 0x88, 0x0e, 0x69,
	0x9b, 0xe9, 0x04, 0xdb, 0xf9, 0x6a, 0x92, 0xba, 0xa2, 0x08, 0x5e, 0xf9, 0x94, 0x52, 0xc1, 0x47,
	0xe8, 0x13, 0x48, 0x93, 0xce, 0x4e, 0x14, 0xdc, 0xf5, 0xaa, 0xc9, 0x77, 0xce, 0x69, 0xcf, 0x39,
	0x1f, 0x45, 0x24, 0x66, 0x3a, 0x59, 0xce, 0xab, 0x32, 0x65, 0x9a, 0xcf, 0x35, 0x40, 0x3e, 0x2f,
	0xb9, 0x2c, 0x84, 0x52, 0x02, 0x56, 0x51, 0x29, 0x41, 0x03, 0x1e, 0xc5, 0x4c, 0x89, 0x64, 0x2f,
	0xcc, 0x84, 0x5e, 0x56, 0x71, 0x94, 0x40, 0x31, 0xcd, 0x20, 0x83, 0xa9, 0x41, 0xe3, 0x6a, 0x61,
	0x6e, 0xe6, 0x62, 0x4e, 0x56, 0xb5, 0x77, 0xd7, 0xa1, 0x17, 0xc7, 0x42, 0xbf, 0x87, 0xe3, 0x69,
	0x06, 0xa1, 0x01, 0xc3, 0x8f, 0x2c, 0x17, 0x29, 0xd3, 0x20, 0xd5, 0xf4, 0xe4, 0xd8, 0xe9, 0xae,
	0x65, 0x00, 0x59, 0xce, 0x37, 0x6f, 0x57, 0x5a, 0x56, 0x89, 0xb6, 0x28, 0xf9, 0xde, 0x47, 0x93,
	0x59, 0x6b, 0xf9, 0x95, 0x71, 0xfc, 0x12, 0x20, 0x7f, 0x7e, 0xe2, 0x97, 0xf2, 0x0f, 0x15, 0x57,
	0x1a, 0xef, 0xa3, 0xad, 0x82, 0xeb, 0x25, 0xa4, 0x5e, 0x6f, 0xd2, 0x0b, 0x76, 0x66, 0xbb, 0x4d,
	0xed, 0x9f, 0x5f, 0x80, 0x2c, 0xee, 0x13, 0x3b, 0x27, 0xb4, 0x23, 0xe0, 0x23, 0x34, 0xde, 0xe4,
	0x55, 0x5e, 0x7f, 0x32, 0x08, 0x76, 0x66, 0x57, 0x9a, 0xda, 0xc7, 0x96, 0xef, 0x80, 0x84, 0xba,
	0xd4, 0x56, 0xc9, 0x2a, 0xbd, 0x04, 0x29, 0x3e, 0x73, 0xa9, 0xbc, 0xc1, 0xdf, 0x4a, 0x07, 0x24,
	0xd4, 0xa5, 0xe2, 0x47, 0x68, 0xbb, 0x2d, 0xfa, 0x69, 0xaa, 0xbc, 0xa1, 0x51, 0xed, 0x37, 0xb5,
	0x7f, 0xc1, 0xaa, 0x3a, 0x80, 0xfc, 0xfa, 0xe9, 0x5f, 0x46, 0xbb, 0xef, 0xde, 0xb0, 0x70, 0xf1,
	0x30, 0x7c, 0x72, 0x27, 0xbc, 0xf7, 0xf6, 0xcb, 0xe1, 0xc1, 0xd7, 0x9b, 0x74, 0xad, 0x24, 0x3f,
	0x7a, 0xe8, 0xfa, 0x29, 0x45, 0xa8, 0x12, 0x56, 0x8a, 0xe3, 0x1b, 0x68, 0x98, 0x40, 0xca, 0x4d,
	0x0f, 0xa3, 0xd9, 0xc5, 0xa6, 0xf6, 0xc7, 0xf6, 0x3b, 0xed, 0x94, 0x50, 0x03, 0xe2, 0x5b, 0x68,
	0xc4, 0xa5, 0x04, 0xe9, 0xf5, 0x4d, 0x5b, 0x97, 0x9a, 0xda, 0x3f, 0x67, 0x59, 0x66, 0x4c, 0xa8,
	0x85, 0xf1, 0x6d, 0xb4, 0x5d, 0x70, 0xa5, 0x58, 0xc6, 0xbd, 0x81, 0x61, 0xe2, 0x8d, 0xef, 0x0e,
	0x20, 0x74, 0x4d, 0xc1, 0x0f, 0xd0, 0x30, 0x65, 0x9a, 0x79, 0xc3, 0x49, 0x2f, 0x18, 0x1f, 0x5c,
	0x8d, 0xec, 0x5a, 0xa3, 0xf5, 0x5a, 0xa3, 0x17, 0x66, 0xad, 0xae, 0xa7, 0x96, 0x4e, 0xa8, 0x51,
	0x91, 0x6f, 0x7d, 0x14, 0x9c, 0x19, 0xef, 0xb5, 0x64, 0x65, 0xc9, 0xe5, 0xff, 0xa5, 0x3c, 0x42,
	0xe3, 0xf6, 0xf9, 0xf8, 0x53, 0x99, 0x33, 0xb1, 0xea, 0xb2, 0x3a, 0xfb, 0x72, 0x40, 0x42, 0x5d,
	0xea, 0xa6, 0x9f, 0xc1, 0xe9, 0xfd, 0x3c, 0xfb, 0x23, 0x71, 0x10, 0x99, 0xdf, 0x26, 0x3a, 0x33,
	0xc5, 0x3f, 0x2a, 0x88, 0xb7, 0x4c, 0x55, 0x87, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x7a,
	0xe2, 0x68, 0xa3, 0x03, 0x00, 0x00,
}
