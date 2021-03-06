// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_all_system_tree.proto

package custom

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	cmdb_extend "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
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
//GetAllSystemTree返回
type GetAllSystemTreeResponse struct {
	//
	//业务树
	Systems              []*cmdb_extend.SystemDependency `protobuf:"bytes,1,rep,name=systems,proto3" json:"systems" form:"systems"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetAllSystemTreeResponse) Reset()         { *m = GetAllSystemTreeResponse{} }
func (m *GetAllSystemTreeResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSystemTreeResponse) ProtoMessage()    {}
func (*GetAllSystemTreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e779998732ac3ecb, []int{0}
}
func (m *GetAllSystemTreeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllSystemTreeResponse.Unmarshal(m, b)
}
func (m *GetAllSystemTreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllSystemTreeResponse.Marshal(b, m, deterministic)
}
func (m *GetAllSystemTreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllSystemTreeResponse.Merge(m, src)
}
func (m *GetAllSystemTreeResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllSystemTreeResponse.Size(m)
}
func (m *GetAllSystemTreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllSystemTreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllSystemTreeResponse proto.InternalMessageInfo

func (m *GetAllSystemTreeResponse) GetSystems() []*cmdb_extend.SystemDependency {
	if m != nil {
		return m.Systems
	}
	return nil
}

//
//GetAllSystemTreeApi返回
type GetAllSystemTreeResponseWrapper struct {
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
	Data                 *GetAllSystemTreeResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetAllSystemTreeResponseWrapper) Reset()         { *m = GetAllSystemTreeResponseWrapper{} }
func (m *GetAllSystemTreeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetAllSystemTreeResponseWrapper) ProtoMessage()    {}
func (*GetAllSystemTreeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e779998732ac3ecb, []int{1}
}
func (m *GetAllSystemTreeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllSystemTreeResponseWrapper.Unmarshal(m, b)
}
func (m *GetAllSystemTreeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllSystemTreeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetAllSystemTreeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllSystemTreeResponseWrapper.Merge(m, src)
}
func (m *GetAllSystemTreeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetAllSystemTreeResponseWrapper.Size(m)
}
func (m *GetAllSystemTreeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllSystemTreeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllSystemTreeResponseWrapper proto.InternalMessageInfo

func (m *GetAllSystemTreeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAllSystemTreeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAllSystemTreeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetAllSystemTreeResponseWrapper) GetData() *GetAllSystemTreeResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllSystemTreeResponse)(nil), "custom.GetAllSystemTreeResponse")
	proto.RegisterType((*GetAllSystemTreeResponseWrapper)(nil), "custom.GetAllSystemTreeResponseWrapper")
}

func init() { proto.RegisterFile("get_all_system_tree.proto", fileDescriptor_e779998732ac3ecb) }

var fileDescriptor_e779998732ac3ecb = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0x71, 0xf3, 0x53, 0x2a, 0x97, 0xb6, 0x68, 0x51, 0xdc, 0x40, 0xb1, 0x51, 0xa1, 0x78,
	0x13, 0x19, 0xd2, 0x4d, 0xe9, 0xae, 0xa1, 0xa1, 0x8b, 0xee, 0xdc, 0x42, 0x57, 0x83, 0x91, 0xe5,
	0x1b, 0x4f, 0x40, 0xf6, 0x15, 0x92, 0x02, 0xc9, 0xcb, 0xfa, 0x15, 0x06, 0xfc, 0x04, 0x43, 0x24,
	0xcf, 0x4c, 0x36, 0x59, 0xd9, 0xf7, 0x9e, 0xf3, 0x1d, 0x1d, 0x24, 0xf2, 0xa9, 0x05, 0x57, 0x09,
	0xa5, 0x2a, 0x7b, 0xb6, 0x0e, 0xba, 0xca, 0x19, 0x00, 0xae, 0x0d, 0x3a, 0xa4, 0x4b, 0x79, 0xb4,
	0x0e, 0xbb, 0xd5, 0xba, 0x3d, 0xb8, 0xfb, 0x63, 0xcd, 0x25, 0x76, 0x45, 0x8b, 0x2d, 0x16, 0x5e,
	0xae, 0x8f, 0x7b, 0x3f, 0xf9, 0xc1, 0xff, 0x05, 0x6c, 0x75, 0xd7, 0x22, 0x07, 0x61, 0xcf, 0xa8,
	0x2d, 0x57, 0x28, 0x85, 0x2a, 0x24, 0xf6, 0xce, 0x08, 0xe9, 0x6c, 0x20, 0x0d, 0x68, 0x5c, 0x77,
	0xd8, 0x80, 0xb2, 0xc5, 0x64, 0x2c, 0xfc, 0x58, 0xc8, 0xae, 0xa9, 0x2b, 0x38, 0x39, 0xe8, 0x9b,
	0x62, 0xea, 0xd3, 0x80, 0x86, 0xbe, 0x81, 0x5e, 0x9e, 0x43, 0x3c, 0x6b, 0x49, 0xf2, 0x1b, 0xdc,
	0x4f, 0xa5, 0xfe, 0x7a, 0xc3, 0x3f, 0x03, 0x50, 0x82, 0xd5, 0xd8, 0x5b, 0xa0, 0x7f, 0xc8, 0xeb,
	0x80, 0xd9, 0x24, 0xca, 0x66, 0x79, 0xbc, 0xf9, 0xcc, 0xaf, 0x22, 0x79, 0x20, 0x7e, 0x3d, 0x27,
	0x6e, 0xe9, 0x38, 0xa4, 0xef, 0xf6, 0x68, 0xba, 0x1f, 0x6c, 0xe2, 0x58, 0xf9, 0x94, 0xc0, 0x1e,
	0x22, 0x92, 0xde, 0x3a, 0xe9, 0xbf, 0x11, 0x5a, 0x83, 0xa1, 0x5f, 0xc8, 0x5c, 0x62, 0x03, 0x49,
	0x94, 0x45, 0xf9, 0x62, 0xfb, 0x7e, 0x1c, 0xd2, 0x38, 0xc4, 0x5d, 0xb6, 0xac, 0xf4, 0x22, 0xfd,
	0x4e, 0xe2, 0xcb, 0x77, 0x77, 0xd2, 0x4a, 0x1c, 0xfa, 0xe4, 0x55, 0x16, 0xe5, 0x6f, 0xb6, 0x1f,
	0xc7, 0x21, 0xa5, 0x2f, 0xde, 0x49, 0x64, 0xe5, 0xb5, 0x95, 0x7e, 0x25, 0x0b, 0x30, 0x06, 0x4d,
	0x32, 0xf3, 0xcc, 0x87, 0x71, 0x48, 0xdf, 0x06, 0xc6, 0xaf, 0x59, 0x19, 0x64, 0xba, 0x23, 0xf3,
	0x46, 0x38, 0x91, 0xcc, 0xb3, 0x28, 0x8f, 0x37, 0x19, 0x0f, 0x0f, 0xc7, 0x6f, 0xb5, 0xbf, 0x2e,
	0x7a, 0xe1, 0x58, 0xe9, 0xf1, 0x7a, 0xe9, 0x6f, 0xf8, 0xdb, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0x2e, 0x68, 0xb6, 0x14, 0x02, 0x00, 0x00,
}
