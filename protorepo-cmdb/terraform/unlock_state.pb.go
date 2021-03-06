// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unlock_state.proto

package terraform

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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
//UnLockState请求
type UnLockStateRequest struct {
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
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user" form:"user"`
	//
	//锁信息
	Body                 *cmdb.TFLockInfo `protobuf:"bytes,5,opt,name=body,proto3" json:"body" form:"body"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UnLockStateRequest) Reset()         { *m = UnLockStateRequest{} }
func (m *UnLockStateRequest) String() string { return proto.CompactTextString(m) }
func (*UnLockStateRequest) ProtoMessage()    {}
func (*UnLockStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48ae9a7415f914c, []int{0}
}
func (m *UnLockStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnLockStateRequest.Unmarshal(m, b)
}
func (m *UnLockStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnLockStateRequest.Marshal(b, m, deterministic)
}
func (m *UnLockStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnLockStateRequest.Merge(m, src)
}
func (m *UnLockStateRequest) XXX_Size() int {
	return xxx_messageInfo_UnLockStateRequest.Size(m)
}
func (m *UnLockStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnLockStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnLockStateRequest proto.InternalMessageInfo

func (m *UnLockStateRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UnLockStateRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *UnLockStateRequest) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *UnLockStateRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UnLockStateRequest) GetBody() *cmdb.TFLockInfo {
	if m != nil {
		return m.Body
	}
	return nil
}

//
//UnLockStateApi返回
type UnLockStateResponseWrapper struct {
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

func (m *UnLockStateResponseWrapper) Reset()         { *m = UnLockStateResponseWrapper{} }
func (m *UnLockStateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UnLockStateResponseWrapper) ProtoMessage()    {}
func (*UnLockStateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48ae9a7415f914c, []int{1}
}
func (m *UnLockStateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnLockStateResponseWrapper.Unmarshal(m, b)
}
func (m *UnLockStateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnLockStateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UnLockStateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnLockStateResponseWrapper.Merge(m, src)
}
func (m *UnLockStateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UnLockStateResponseWrapper.Size(m)
}
func (m *UnLockStateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UnLockStateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UnLockStateResponseWrapper proto.InternalMessageInfo

func (m *UnLockStateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UnLockStateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UnLockStateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UnLockStateResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UnLockStateRequest)(nil), "terraform.UnLockStateRequest")
	proto.RegisterType((*UnLockStateResponseWrapper)(nil), "terraform.UnLockStateResponseWrapper")
}

func init() { proto.RegisterFile("unlock_state.proto", fileDescriptor_d48ae9a7415f914c) }

var fileDescriptor_d48ae9a7415f914c = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x6f, 0x6f, 0xd2, 0x40,
	0x18, 0x17, 0x06, 0x46, 0x0e, 0xe3, 0xf0, 0x62, 0x16, 0x82, 0x31, 0xc5, 0x6e, 0x59, 0x30, 0xb1,
	0x2d, 0x83, 0xb8, 0x38, 0x7d, 0x25, 0xc9, 0x4c, 0x30, 0xbe, 0xb1, 0x6a, 0x4c, 0x24, 0x8c, 0x5c,
	0xaf, 0x47, 0xad, 0xb4, 0x7d, 0xea, 0xdd, 0xe1, 0x44, 0xb2, 0xaf, 0x5a, 0x8d, 0x1f, 0x81, 0x4f,
	0x60, 0xee, 0x5a, 0xa0, 0x7b, 0xd5, 0xe7, 0xee, 0xf7, 0x27, 0xbf, 0xfb, 0xf5, 0x41, 0x78, 0x99,
	0x44, 0x40, 0x17, 0x33, 0x21, 0x89, 0x64, 0x76, 0xca, 0x41, 0x02, 0x6e, 0x48, 0xc6, 0x39, 0x99,
	0x03, 0x8f, 0x3b, 0x56, 0x10, 0xca, 0x6f, 0x4b, 0xcf, 0xa6, 0x10, 0x3b, 0x01, 0x04, 0xe0, 0x68,
	0x86, 0xb7, 0x9c, 0xeb, 0x93, 0x3e, 0xe8, 0x29, 0x57, 0x76, 0x3e, 0x04, 0x60, 0x33, 0x22, 0x56,
	0x90, 0x0a, 0x3b, 0x02, 0x4a, 0x22, 0x87, 0x42, 0x22, 0x39, 0xa1, 0x52, 0xe4, 0x4a, 0xce, 0x52,
	0xb0, 0x62, 0xf0, 0x59, 0x24, 0x9c, 0x82, 0xe8, 0xe8, 0xa3, 0x43, 0x63, 0xdf, 0x73, 0xe4, 0x7c,
	0xa6, 0xd3, 0x84, 0xc9, 0x7c, 0x6b, 0x79, 0x5e, 0x4a, 0x10, 0x5f, 0x87, 0x72, 0x01, 0xd7, 0x4e,
	0x00, 0x96, 0x06, 0xad, 0x9f, 0x24, 0x0a, 0x7d, 0x22, 0x81, 0x0b, 0x67, 0x37, 0x16, 0xba, 0xc7,
	0x01, 0x40, 0x10, 0xb1, 0x7d, 0x60, 0x16, 0xa7, 0x72, 0x95, 0x83, 0x66, 0x56, 0x45, 0xf8, 0x73,
	0xf2, 0x1e, 0xe8, 0xe2, 0xa3, 0x7a, 0xb7, 0xcb, 0x7e, 0x2c, 0x99, 0x90, 0xd8, 0x45, 0xf7, 0xc0,
	0xfb, 0xce, 0xa8, 0x1c, 0xfb, 0xed, 0x4a, 0xb7, 0xd2, 0x6b, 0x8c, 0xce, 0x37, 0x99, 0x71, 0xa8,
	0x9a, 0x78, 0x65, 0x6e, 0x11, 0xf3, 0xdf, 0x1f, 0xc3, 0x40, 0x4f, 0xae, 0x26, 0xc4, 0xfa, 0xfd,
	0xc6, 0xfa, 0x3a, 0x9b, 0x4e, 0xfa, 0xd6, 0xc5, 0x76, 0x5e, 0xf7, 0x9f, 0x0f, 0xcf, 0x6e, 0x4e,
	0xdc, 0x9d, 0x0f, 0x1e, 0x23, 0x14, 0x26, 0x42, 0x92, 0x84, 0xb2, 0xb1, 0xdf, 0xae, 0x6a, 0xd7,
	0x67, 0x9b, 0xcc, 0x78, 0x98, 0xbb, 0xee, 0x31, 0xe5, 0xdb, 0x42, 0x0f, 0xae, 0x0a, 0xbb, 0xe9,
	0xfa, 0x6c, 0x78, 0x73, 0xe2, 0x96, 0xc4, 0xf8, 0x14, 0x1d, 0x00, 0x0f, 0xda, 0x07, 0xdd, 0x4a,
	0xaf, 0x3e, 0x7a, 0xb4, 0xc9, 0x0c, 0x54, 0x24, 0xe3, 0x81, 0x12, 0x57, 0x5b, 0x77, 0x5c, 0x45,
	0xc0, 0xef, 0x50, 0x6d, 0x29, 0x18, 0x6f, 0xd7, 0x76, 0x4f, 0x68, 0xe6, 0x44, 0x75, 0xab, 0x98,
	0xc7, 0xe8, 0xe9, 0x36, 0x7e, 0xdf, 0xba, 0x98, 0x4e, 0xec, 0xdd, 0x3c, 0xb3, 0xa6, 0xeb, 0x41,
	0xfe, 0x04, 0xed, 0x81, 0x5f, 0xa0, 0x9a, 0x07, 0xfe, 0xaa, 0x5d, 0xef, 0x56, 0x7a, 0xcd, 0x41,
	0xcb, 0x56, 0xbf, 0xc9, 0xfe, 0xf4, 0x56, 0x55, 0x37, 0x4e, 0xe6, 0x30, 0x3a, 0xdc, 0xbb, 0x2b,
	0x9e, 0xe9, 0x6a, 0xba, 0xf9, 0xb7, 0x82, 0x3a, 0xb7, 0x0a, 0x16, 0x29, 0x24, 0x82, 0x7d, 0xe1,
	0x24, 0x4d, 0x19, 0xc7, 0xc7, 0xa8, 0x46, 0xc1, 0x67, 0xba, 0xe4, 0x7a, 0xd9, 0x43, 0xdd, 0x9a,
	0xae, 0x06, 0xf1, 0x4b, 0xd4, 0x54, 0xdf, 0xcb, 0x5f, 0x69, 0x44, 0xc2, 0xa4, 0xa8, 0xee, 0x68,
	0x93, 0x19, 0x78, 0xcf, 0x2d, 0x40, 0xd3, 0x2d, 0x53, 0xf1, 0x29, 0xaa, 0x33, 0xce, 0x81, 0xeb,
	0xaa, 0x1a, 0xa3, 0xd6, 0x26, 0x33, 0xee, 0xe7, 0x1a, 0x7d, 0x6d, 0xba, 0x39, 0x8c, 0x5f, 0xa3,
	0x9a, 0x4f, 0x24, 0xd1, 0x45, 0x35, 0x07, 0x47, 0x76, 0xbe, 0x32, 0xf6, 0x76, 0x65, 0xec, 0x4b,
	0xb5, 0x32, 0xe5, 0x78, 0x8a, 0x6d, 0xba, 0x5a, 0xe4, 0xdd, 0xd5, 0xb4, 0xe1, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0x63, 0xe9, 0x06, 0x42, 0x03, 0x00, 0x00,
}
