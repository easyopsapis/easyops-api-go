// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archive_instance_list_all.proto

package cmdb_service_ctrl

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
//ArchiveInstanceListAll请求
type ArchiveInstanceListAllRequest struct {
	//
	//实例所属的模型ID
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveInstanceListAllRequest) Reset()         { *m = ArchiveInstanceListAllRequest{} }
func (m *ArchiveInstanceListAllRequest) String() string { return proto.CompactTextString(m) }
func (*ArchiveInstanceListAllRequest) ProtoMessage()    {}
func (*ArchiveInstanceListAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e5197c2f5835a, []int{0}
}
func (m *ArchiveInstanceListAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveInstanceListAllRequest.Unmarshal(m, b)
}
func (m *ArchiveInstanceListAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveInstanceListAllRequest.Marshal(b, m, deterministic)
}
func (m *ArchiveInstanceListAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveInstanceListAllRequest.Merge(m, src)
}
func (m *ArchiveInstanceListAllRequest) XXX_Size() int {
	return xxx_messageInfo_ArchiveInstanceListAllRequest.Size(m)
}
func (m *ArchiveInstanceListAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveInstanceListAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveInstanceListAllRequest proto.InternalMessageInfo

func (m *ArchiveInstanceListAllRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

//
//ArchiveInstanceListAll返回
type ArchiveInstanceListAllResponse struct {
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
	//实例列表
	Data                 []*types.Struct `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ArchiveInstanceListAllResponse) Reset()         { *m = ArchiveInstanceListAllResponse{} }
func (m *ArchiveInstanceListAllResponse) String() string { return proto.CompactTextString(m) }
func (*ArchiveInstanceListAllResponse) ProtoMessage()    {}
func (*ArchiveInstanceListAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e5197c2f5835a, []int{1}
}
func (m *ArchiveInstanceListAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveInstanceListAllResponse.Unmarshal(m, b)
}
func (m *ArchiveInstanceListAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveInstanceListAllResponse.Marshal(b, m, deterministic)
}
func (m *ArchiveInstanceListAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveInstanceListAllResponse.Merge(m, src)
}
func (m *ArchiveInstanceListAllResponse) XXX_Size() int {
	return xxx_messageInfo_ArchiveInstanceListAllResponse.Size(m)
}
func (m *ArchiveInstanceListAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveInstanceListAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveInstanceListAllResponse proto.InternalMessageInfo

func (m *ArchiveInstanceListAllResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ArchiveInstanceListAllResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ArchiveInstanceListAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ArchiveInstanceListAllResponse) GetData() []*types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//ArchiveInstanceListAllApi返回
type ArchiveInstanceListAllResponseWrapper struct {
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
	Data                 *ArchiveInstanceListAllResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ArchiveInstanceListAllResponseWrapper) Reset()         { *m = ArchiveInstanceListAllResponseWrapper{} }
func (m *ArchiveInstanceListAllResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ArchiveInstanceListAllResponseWrapper) ProtoMessage()    {}
func (*ArchiveInstanceListAllResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e5197c2f5835a, []int{2}
}
func (m *ArchiveInstanceListAllResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveInstanceListAllResponseWrapper.Unmarshal(m, b)
}
func (m *ArchiveInstanceListAllResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveInstanceListAllResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ArchiveInstanceListAllResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveInstanceListAllResponseWrapper.Merge(m, src)
}
func (m *ArchiveInstanceListAllResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ArchiveInstanceListAllResponseWrapper.Size(m)
}
func (m *ArchiveInstanceListAllResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveInstanceListAllResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveInstanceListAllResponseWrapper proto.InternalMessageInfo

func (m *ArchiveInstanceListAllResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ArchiveInstanceListAllResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ArchiveInstanceListAllResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ArchiveInstanceListAllResponseWrapper) GetData() *ArchiveInstanceListAllResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ArchiveInstanceListAllRequest)(nil), "cmdb_service_ctrl.ArchiveInstanceListAllRequest")
	proto.RegisterType((*ArchiveInstanceListAllResponse)(nil), "cmdb_service_ctrl.ArchiveInstanceListAllResponse")
	proto.RegisterType((*ArchiveInstanceListAllResponseWrapper)(nil), "cmdb_service_ctrl.ArchiveInstanceListAllResponseWrapper")
}

func init() { proto.RegisterFile("archive_instance_list_all.proto", fileDescriptor_319e5197c2f5835a) }

var fileDescriptor_319e5197c2f5835a = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8b, 0xd4, 0x30,
	0x18, 0xc7, 0x99, 0x9d, 0x59, 0x5f, 0x32, 0xe2, 0x6a, 0x0e, 0x5a, 0x16, 0xd7, 0x2c, 0xf1, 0x85,
	0x39, 0x6c, 0xdb, 0x7d, 0x81, 0x45, 0xc5, 0xcb, 0x2c, 0x78, 0x58, 0xf0, 0x14, 0x41, 0xc1, 0x45,
	0x4b, 0x9a, 0x66, 0xbb, 0xd1, 0xb4, 0xa9, 0x49, 0x3a, 0x2b, 0x8a, 0x57, 0xbf, 0xa3, 0x97, 0x0a,
	0x7e, 0x84, 0x7e, 0x02, 0x99, 0xa4, 0x9d, 0x19, 0x90, 0x1d, 0x3c, 0xf5, 0x49, 0xfe, 0xbf, 0x3c,
	0xcf, 0xf3, 0xff, 0x53, 0x80, 0xa8, 0x66, 0x17, 0x62, 0xc6, 0x13, 0x51, 0x1a, 0x4b, 0x4b, 0xc6,
	0x13, 0x29, 0x8c, 0x4d, 0xa8, 0x94, 0x51, 0xa5, 0x95, 0x55, 0xf0, 0x2e, 0x2b, 0xb2, 0x34, 0x31,
	0x5c, 0xcf, 0x04, 0xe3, 0x09, 0xb3, 0x5a, 0x6e, 0x87, 0xb9, 0xb0, 0x17, 0x75, 0x1a, 0x31, 0x55,
	0xc4, 0xb9, 0xca, 0x55, 0xec, 0xc8, 0xb4, 0x3e, 0x77, 0x27, 0x77, 0x70, 0x95, 0xef, 0xb0, 0x7d,
	0xbc, 0x82, 0x17, 0x97, 0xc2, 0x7e, 0x56, 0x97, 0x71, 0xae, 0x42, 0x27, 0x86, 0x33, 0x2a, 0x45,
	0x46, 0xad, 0xd2, 0x26, 0x5e, 0x94, 0xdd, 0xbb, 0x07, 0xb9, 0x52, 0xb9, 0xe4, 0xcb, 0xee, 0xc6,
	0xea, 0x9a, 0x59, 0xaf, 0x62, 0x03, 0x76, 0xa6, 0x7e, 0xf5, 0xd3, 0x6e, 0xf3, 0xd7, 0xc2, 0xd8,
	0xa9, 0x94, 0x84, 0x7f, 0xa9, 0xb9, 0xb1, 0x90, 0x80, 0x1b, 0x2a, 0xfd, 0xc4, 0x99, 0x3d, 0xcd,
	0x82, 0xc1, 0xee, 0x60, 0x72, 0xf3, 0xe4, 0xb8, 0x6d, 0xd0, 0xd6, 0xb9, 0xd2, 0xc5, 0x0b, 0xdc,
	0x2b, 0xf8, 0xcf, 0x6f, 0x84, 0xc0, 0xce, 0xc7, 0x33, 0x1a, 0x7e, 0x9b, 0x86, 0xef, 0x93, 0x0f,
	0x67, 0xfb, 0xe1, 0xf3, 0xbe, 0xfe, 0xbe, 0xbf, 0x77, 0x74, 0xf0, 0xe3, 0x31, 0x59, 0xf4, 0xc1,
	0xbf, 0x06, 0xe0, 0xe1, 0x55, 0x53, 0x4d, 0xa5, 0x4a, 0xc3, 0xe1, 0x23, 0x30, 0x62, 0x2a, 0xe3,
	0x6e, 0xe4, 0xe6, 0xc9, 0x56, 0xdb, 0xa0, 0xb1, 0x1f, 0x39, 0xbf, 0xc5, 0xc4, 0x89, 0xf0, 0x29,
	0xd8, 0xe4, 0x5a, 0x2b, 0x1d, 0x6c, 0xb8, 0xc5, 0xee, 0xb4, 0x0d, 0xba, 0xe5, 0x29, 0x77, 0x8d,
	0x89, 0x97, 0xe1, 0x1e, 0xb8, 0x5e, 0x70, 0x63, 0x68, 0xce, 0x83, 0xa1, 0x23, 0x61, 0xdb, 0xa0,
	0xdb, 0x9e, 0xec, 0x04, 0x4c, 0x7a, 0x04, 0xbe, 0x04, 0xa3, 0x8c, 0x5a, 0x1a, 0x8c, 0x76, 0x87,
	0x93, 0xf1, 0xe1, 0xfd, 0xc8, 0xe7, 0x17, 0xf5, 0xf9, 0x45, 0x6f, 0x5c, 0x7e, 0xab, 0x3b, 0xcd,
	0x71, 0x4c, 0xdc, 0x2b, 0xfc, 0x73, 0x03, 0x3c, 0x59, 0xef, 0xed, 0x9d, 0xa6, 0x55, 0xc5, 0xf5,
	0xff, 0x59, 0x7c, 0x06, 0xc6, 0xf3, 0xef, 0xab, 0xaf, 0x95, 0xa4, 0xa2, 0xec, 0x8c, 0xde, 0x6b,
	0x1b, 0x04, 0x97, 0x6c, 0x27, 0x62, 0xb2, 0x8a, 0x2e, 0xc3, 0x19, 0xae, 0x0f, 0xe7, 0xed, 0xc2,
	0xee, 0x60, 0x32, 0x3e, 0x3c, 0x88, 0xfe, 0xf9, 0x51, 0xa3, 0xf5, 0x76, 0xae, 0x08, 0x22, 0xbd,
	0xe6, 0x02, 0x3b, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xad, 0xd8, 0x0c, 0x1b, 0x03, 0x00,
	0x00,
}