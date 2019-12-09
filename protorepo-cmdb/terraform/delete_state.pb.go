// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_state.proto

package terraform

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
//DeleteState请求
type DeleteStateRequest struct {
	//
	//模型Id
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例Id
	InstanceId string `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//ORG
	Org int32 `protobuf:"varint,3,opt,name=org,proto3" json:"org" form:"org"`
	//
	//操作用户,默认defaultUser
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user" form:"user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStateRequest) Reset()         { *m = DeleteStateRequest{} }
func (m *DeleteStateRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteStateRequest) ProtoMessage()    {}
func (*DeleteStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56de079644f6296, []int{0}
}
func (m *DeleteStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStateRequest.Unmarshal(m, b)
}
func (m *DeleteStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStateRequest.Marshal(b, m, deterministic)
}
func (m *DeleteStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStateRequest.Merge(m, src)
}
func (m *DeleteStateRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteStateRequest.Size(m)
}
func (m *DeleteStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStateRequest proto.InternalMessageInfo

func (m *DeleteStateRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *DeleteStateRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *DeleteStateRequest) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *DeleteStateRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

//
//DeleteStateApi返回
type DeleteStateResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteStateResponseWrapper) Reset()         { *m = DeleteStateResponseWrapper{} }
func (m *DeleteStateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteStateResponseWrapper) ProtoMessage()    {}
func (*DeleteStateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56de079644f6296, []int{1}
}
func (m *DeleteStateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStateResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteStateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteStateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStateResponseWrapper.Merge(m, src)
}
func (m *DeleteStateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteStateResponseWrapper.Size(m)
}
func (m *DeleteStateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStateResponseWrapper proto.InternalMessageInfo

func (m *DeleteStateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteStateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteStateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteStateResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteStateRequest)(nil), "terraform.DeleteStateRequest")
	proto.RegisterType((*DeleteStateResponseWrapper)(nil), "terraform.DeleteStateResponseWrapper")
}

func init() { proto.RegisterFile("delete_state.proto", fileDescriptor_a56de079644f6296) }

var fileDescriptor_a56de079644f6296 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x5f, 0x6f, 0x12, 0x41,
	0x10, 0x17, 0x0a, 0x46, 0x06, 0x63, 0x71, 0x63, 0x1a, 0x82, 0x31, 0x57, 0xaf, 0x4d, 0x53, 0x13,
	0xef, 0xa0, 0x25, 0x69, 0xac, 0x3e, 0x49, 0xec, 0x03, 0x3e, 0xae, 0x0f, 0x26, 0x12, 0x4a, 0x16,
	0x6e, 0xba, 0x9e, 0x02, 0x73, 0xee, 0x2d, 0x56, 0x25, 0x7c, 0x1e, 0xbf, 0xd5, 0x69, 0xfc, 0x08,
	0xf7, 0x09, 0xcc, 0xee, 0x1e, 0x7f, 0xfa, 0x74, 0xb3, 0xf3, 0xfb, 0x73, 0xbf, 0x99, 0x01, 0x16,
	0xe1, 0x14, 0x35, 0x8e, 0x52, 0x2d, 0x34, 0x86, 0x89, 0x22, 0x4d, 0xac, 0xa6, 0x51, 0x29, 0x71,
	0x43, 0x6a, 0xd6, 0x0a, 0x64, 0xac, 0x3f, 0x2f, 0xc6, 0xe1, 0x84, 0x66, 0x6d, 0x49, 0x92, 0xda,
	0x96, 0x31, 0x5e, 0xdc, 0xd8, 0x97, 0x7d, 0xd8, 0xca, 0x29, 0x5b, 0x17, 0x3b, 0xf4, 0xd9, 0x6d,
	0xac, 0xbf, 0xd2, 0x6d, 0x5b, 0x52, 0x60, 0xc1, 0xe0, 0xbb, 0x98, 0xc6, 0x91, 0xd0, 0xa4, 0xd2,
	0xf6, 0xa6, 0x2c, 0x74, 0x4f, 0x25, 0x91, 0x9c, 0xe2, 0xd6, 0x1d, 0x67, 0x89, 0xfe, 0xe9, 0x40,
	0xff, 0x77, 0x19, 0xd8, 0x3b, 0x9b, 0xf2, 0x83, 0x09, 0xc9, 0xf1, 0xdb, 0x02, 0x53, 0xcd, 0x38,
	0x3c, 0xa0, 0xf1, 0x17, 0x9c, 0xe8, 0x7e, 0xd4, 0x2c, 0x1d, 0x96, 0x4e, 0x6b, 0xbd, 0x8b, 0x3c,
	0xf3, 0xf6, 0x4d, 0xec, 0xd7, 0xfe, 0x1a, 0xf1, 0xff, 0xfd, 0xf1, 0x3c, 0x78, 0x76, 0x3d, 0x10,
	0xc1, 0xaf, 0xb7, 0xc1, 0xa7, 0xd1, 0x70, 0xd0, 0x09, 0x2e, 0xd7, 0xf5, 0xb2, 0xf3, 0xb2, 0x7b,
	0xb6, 0x3a, 0xe6, 0x1b, 0x1f, 0xd6, 0x07, 0x88, 0xe7, 0xa9, 0x16, 0xf3, 0x09, 0xf6, 0xa3, 0x66,
	0xd9, 0xba, 0xbe, 0xc8, 0x33, 0xef, 0xb1, 0x73, 0xdd, 0x62, 0xc6, 0xb7, 0x01, 0x8f, 0xae, 0x0b,
	0xbb, 0xe1, 0xf2, 0xac, 0xbb, 0x3a, 0xe6, 0x3b, 0x62, 0x76, 0x02, 0x7b, 0xa4, 0x64, 0x73, 0xef,
	0xb0, 0x74, 0x5a, 0xed, 0x3d, 0xc9, 0x33, 0x0f, 0x8a, 0x64, 0x4a, 0x1a, 0x71, 0xb9, 0x71, 0x8f,
	0x1b, 0x02, 0x7b, 0x0f, 0x95, 0x45, 0x8a, 0xaa, 0x59, 0xd9, 0x8c, 0x50, 0x77, 0x44, 0xd3, 0x35,
	0xcc, 0x23, 0x78, 0xbe, 0x8e, 0xdf, 0x09, 0x2e, 0x87, 0x83, 0x70, 0x53, 0x8f, 0x82, 0xe1, 0xf2,
	0xdc, 0x8d, 0x60, 0x3d, 0xfc, 0xbf, 0x25, 0x68, 0xdd, 0xd9, 0x54, 0x9a, 0xd0, 0x3c, 0xc5, 0x8f,
	0x4a, 0x24, 0x09, 0x2a, 0x76, 0x04, 0x95, 0x09, 0x45, 0x68, 0xb7, 0x55, 0xed, 0xed, 0x6f, 0x7f,
	0x65, 0xba, 0x3e, 0xb7, 0x20, 0x7b, 0x05, 0x75, 0xf3, 0xbd, 0xfa, 0x91, 0x4c, 0x45, 0x3c, 0x2f,
	0x76, 0x70, 0x90, 0x67, 0x1e, 0xdb, 0x72, 0x0b, 0xd0, 0xe7, 0xbb, 0x54, 0x76, 0x02, 0x55, 0x54,
	0x8a, 0x94, 0x9d, 0xb9, 0xd6, 0x6b, 0xe4, 0x99, 0xf7, 0xd0, 0x69, 0x6c, 0xdb, 0xe7, 0x0e, 0x66,
	0x6f, 0xa0, 0x12, 0x09, 0x2d, 0xec, 0xc4, 0xf5, 0xf3, 0x83, 0xd0, 0xdd, 0x3e, 0x5c, 0xdf, 0x3e,
	0xbc, 0x32, 0xb7, 0xdf, 0x8d, 0x67, 0xd8, 0x3e, 0xb7, 0xa2, 0xf1, 0x7d, 0x4b, 0xeb, 0xfe, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x0c, 0x45, 0x23, 0xb8, 0x02, 0x00, 0x00,
}
