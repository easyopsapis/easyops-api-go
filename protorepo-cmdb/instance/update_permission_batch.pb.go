// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_permission_batch.proto

package instance

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
//UpdatePermissionBatch请求
type UpdatePermissionBatchRequest struct {
	//
	//资源模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例ID列表
	Ids []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids" form:"ids"`
	//
	//权限字段
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field" form:"field"`
	//
	//updateAthorizers, deleteAuthorizers, readAuthorizers, operateAuthorizers中的一个或多个
	Fields []string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields" form:"fields"`
	//
	//修改操作的类型(overwrite, append, pull)(在枚举之外的操作类型会报错)
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method" form:"method"`
	//
	//用户,用户组列表
	List                 []string `protobuf:"bytes,6,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePermissionBatchRequest) Reset()         { *m = UpdatePermissionBatchRequest{} }
func (m *UpdatePermissionBatchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePermissionBatchRequest) ProtoMessage()    {}
func (*UpdatePermissionBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69548149f90027e6, []int{0}
}
func (m *UpdatePermissionBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePermissionBatchRequest.Unmarshal(m, b)
}
func (m *UpdatePermissionBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePermissionBatchRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePermissionBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePermissionBatchRequest.Merge(m, src)
}
func (m *UpdatePermissionBatchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePermissionBatchRequest.Size(m)
}
func (m *UpdatePermissionBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePermissionBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePermissionBatchRequest proto.InternalMessageInfo

func (m *UpdatePermissionBatchRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UpdatePermissionBatchRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *UpdatePermissionBatchRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *UpdatePermissionBatchRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *UpdatePermissionBatchRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UpdatePermissionBatchRequest) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

//
//UpdatePermissionBatch返回
type UpdatePermissionBatchResponse struct {
	//
	//没有权限更新的实例列表
	NoAuthorizedList     []*types.Struct `protobuf:"bytes,1,rep,name=noAuthorizedList,proto3" json:"noAuthorizedList" form:"noAuthorizedList"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdatePermissionBatchResponse) Reset()         { *m = UpdatePermissionBatchResponse{} }
func (m *UpdatePermissionBatchResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePermissionBatchResponse) ProtoMessage()    {}
func (*UpdatePermissionBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69548149f90027e6, []int{1}
}
func (m *UpdatePermissionBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePermissionBatchResponse.Unmarshal(m, b)
}
func (m *UpdatePermissionBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePermissionBatchResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePermissionBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePermissionBatchResponse.Merge(m, src)
}
func (m *UpdatePermissionBatchResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePermissionBatchResponse.Size(m)
}
func (m *UpdatePermissionBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePermissionBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePermissionBatchResponse proto.InternalMessageInfo

func (m *UpdatePermissionBatchResponse) GetNoAuthorizedList() []*types.Struct {
	if m != nil {
		return m.NoAuthorizedList
	}
	return nil
}

//
//UpdatePermissionBatchApi返回
type UpdatePermissionBatchResponseWrapper struct {
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
	Data                 *UpdatePermissionBatchResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UpdatePermissionBatchResponseWrapper) Reset()         { *m = UpdatePermissionBatchResponseWrapper{} }
func (m *UpdatePermissionBatchResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdatePermissionBatchResponseWrapper) ProtoMessage()    {}
func (*UpdatePermissionBatchResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_69548149f90027e6, []int{2}
}
func (m *UpdatePermissionBatchResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePermissionBatchResponseWrapper.Unmarshal(m, b)
}
func (m *UpdatePermissionBatchResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePermissionBatchResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdatePermissionBatchResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePermissionBatchResponseWrapper.Merge(m, src)
}
func (m *UpdatePermissionBatchResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdatePermissionBatchResponseWrapper.Size(m)
}
func (m *UpdatePermissionBatchResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePermissionBatchResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePermissionBatchResponseWrapper proto.InternalMessageInfo

func (m *UpdatePermissionBatchResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdatePermissionBatchResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdatePermissionBatchResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdatePermissionBatchResponseWrapper) GetData() *UpdatePermissionBatchResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdatePermissionBatchRequest)(nil), "instance.UpdatePermissionBatchRequest")
	proto.RegisterType((*UpdatePermissionBatchResponse)(nil), "instance.UpdatePermissionBatchResponse")
	proto.RegisterType((*UpdatePermissionBatchResponseWrapper)(nil), "instance.UpdatePermissionBatchResponseWrapper")
}

func init() { proto.RegisterFile("update_permission_batch.proto", fileDescriptor_69548149f90027e6) }

var fileDescriptor_69548149f90027e6 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xfa, 0xa7, 0xcd, 0x05, 0xd6, 0x05, 0x89, 0x45, 0x63, 0x53, 0x2a, 0x53, 0x41, 0x41,
	0x24, 0x65, 0x2d, 0x42, 0xc0, 0xdd, 0x2a, 0x81, 0x84, 0xb4, 0x0b, 0x64, 0x84, 0x90, 0x98, 0xb6,
	0xe2, 0x26, 0x6e, 0x6a, 0x68, 0xe3, 0x60, 0x3b, 0x0c, 0x75, 0xda, 0x05, 0x4f, 0xc6, 0x9b, 0x04,
	0x89, 0x47, 0xc8, 0x2d, 0x37, 0xc8, 0x27, 0x6d, 0x17, 0x31, 0xb1, 0xab, 0xd8, 0xe7, 0xfb, 0xf1,
	0x39, 0xe7, 0x53, 0xd0, 0x7e, 0x9a, 0x84, 0x54, 0xb3, 0x51, 0xc2, 0xe4, 0x9c, 0x2b, 0xc5, 0x45,
	0x3c, 0x1a, 0x53, 0x1d, 0x4c, 0xfd, 0x44, 0x0a, 0x2d, 0xec, 0x0d, 0x1e, 0x2b, 0x4d, 0xe3, 0x80,
	0xed, 0x7a, 0x11, 0xd7, 0xd3, 0x74, 0xec, 0x07, 0x62, 0xde, 0x8b, 0x44, 0x24, 0x7a, 0x40, 0x18,
	0xa7, 0x13, 0xb8, 0xc1, 0x05, 0x4e, 0x85, 0x70, 0xf7, 0x59, 0x89, 0x3e, 0x3f, 0xe3, 0xfa, 0x8b,
	0x38, 0xeb, 0x45, 0xc2, 0x03, 0xd0, 0xfb, 0x46, 0x67, 0x3c, 0xa4, 0x5a, 0x48, 0xd5, 0x5b, 0x1f,
	0x97, 0xba, 0xbd, 0x48, 0x88, 0x68, 0xc6, 0x2e, 0xdd, 0x95, 0x96, 0x69, 0xa0, 0x0b, 0x14, 0xff,
	0xac, 0xa0, 0xbd, 0xf7, 0xd0, 0xf0, 0xdb, 0x75, 0xbf, 0x43, 0xd3, 0x2e, 0x61, 0x5f, 0x53, 0xa6,
	0xb4, 0xfd, 0x1a, 0x6d, 0x88, 0xf1, 0x67, 0x16, 0xe8, 0x37, 0xa1, 0x63, 0xb5, 0xad, 0xee, 0xe6,
	0xf0, 0x51, 0x9e, 0xb9, 0x5b, 0x13, 0x21, 0xe7, 0x2f, 0xf1, 0x0a, 0xc1, 0xbf, 0x7f, 0xb9, 0xb7,
	0xd1, 0xf6, 0xe9, 0x31, 0xf5, 0x16, 0x87, 0xde, 0xc7, 0xd1, 0xc9, 0xf9, 0xc1, 0xe3, 0x41, 0xff,
	0xa2, 0x43, 0xd6, 0x5a, 0xfb, 0x29, 0xaa, 0xf2, 0x50, 0x39, 0x95, 0x76, 0xb5, 0xbb, 0x39, 0xc4,
	0x79, 0xe6, 0xa2, 0xc2, 0x82, 0x87, 0xca, 0xa8, 0x5b, 0xe8, 0xd6, 0xe9, 0xf1, 0x13, 0xef, 0x05,
	0xf5, 0x16, 0x27, 0xe7, 0x07, 0x83, 0x8b, 0x0e, 0x31, 0x74, 0xfb, 0x3e, 0xaa, 0x4f, 0x38, 0x9b,
	0x85, 0x4e, 0x15, 0x9e, 0x6e, 0xe5, 0x99, 0x7b, 0xa3, 0xd0, 0x41, 0x19, 0x93, 0x02, 0xb6, 0x1f,
	0xa2, 0x06, 0x1c, 0x94, 0x53, 0x83, 0x07, 0xb6, 0xf3, 0xcc, 0xbd, 0x59, 0x22, 0x2a, 0x4c, 0x96,
	0x04, 0x43, 0x9d, 0x33, 0x3d, 0x15, 0xa1, 0x53, 0x07, 0xcf, 0x12, 0xb5, 0xa8, 0x63, 0xb2, 0x24,
	0xd8, 0xf7, 0x50, 0x6d, 0xc6, 0x95, 0x76, 0x1a, 0xe0, 0xb9, 0x95, 0x67, 0x6e, 0xb3, 0x20, 0x9a,
	0x2a, 0x26, 0x00, 0xe2, 0x1f, 0x16, 0xda, 0xff, 0xcf, 0x06, 0x55, 0x22, 0x62, 0xc5, 0xec, 0x4f,
	0xa8, 0x15, 0x8b, 0xc3, 0x54, 0x4f, 0x85, 0xe4, 0x0b, 0x16, 0x1e, 0x19, 0x4b, 0xab, 0x5d, 0xed,
	0x36, 0xfb, 0x3b, 0x7e, 0x11, 0x8e, 0xbf, 0x0a, 0xc7, 0x7f, 0x07, 0xe1, 0x0c, 0xef, 0xe6, 0x99,
	0xbb, 0x53, 0xbc, 0xf5, 0xaf, 0x14, 0x93, 0x2b, 0x6e, 0xf8, 0x8f, 0x85, 0x3a, 0xd7, 0xf6, 0xf0,
	0x41, 0xd2, 0x24, 0x61, 0xd2, 0x4c, 0x14, 0x88, 0x90, 0x41, 0x92, 0xf5, 0xf2, 0x44, 0xa6, 0x8a,
	0x09, 0x80, 0xf6, 0x73, 0xd4, 0x34, 0xdf, 0x57, 0xdf, 0x93, 0x19, 0xe5, 0xb1, 0x53, 0x81, 0x35,
	0xdd, 0xc9, 0x33, 0xd7, 0xbe, 0xe4, 0x2e, 0x41, 0x4c, 0xca, 0x54, 0x13, 0x17, 0x93, 0x52, 0xc8,
	0xab, 0x71, 0x41, 0x19, 0x93, 0x02, 0xb6, 0x8f, 0x50, 0x2d, 0xa4, 0x9a, 0x3a, 0xb5, 0xb6, 0xd5,
	0x6d, 0xf6, 0x1f, 0xf8, 0xab, 0x7f, 0xc2, 0xbf, 0x76, 0x88, 0x72, 0xbf, 0x46, 0x8e, 0x09, 0xb8,
	0x8c, 0x1b, 0xb0, 0xbd, 0xc1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x99, 0x69, 0x1a, 0x7a,
	0x03, 0x00, 0x00,
}