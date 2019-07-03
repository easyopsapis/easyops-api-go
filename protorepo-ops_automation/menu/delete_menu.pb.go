// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_menu.proto

package menu

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//DeleteMenu请求
type DeleteMenuRequest struct {
	//
	//菜单id
	MenusId              string   `protobuf:"bytes,1,opt,name=menusId,proto3" json:"menusId" form:"menusId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMenuRequest) Reset()         { *m = DeleteMenuRequest{} }
func (m *DeleteMenuRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMenuRequest) ProtoMessage()    {}
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72bfa346086d8ba9, []int{0}
}
func (m *DeleteMenuRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMenuRequest.Unmarshal(m, b)
}
func (m *DeleteMenuRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMenuRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMenuRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMenuRequest.Merge(m, src)
}
func (m *DeleteMenuRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMenuRequest.Size(m)
}
func (m *DeleteMenuRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMenuRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMenuRequest proto.InternalMessageInfo

func (m *DeleteMenuRequest) GetMenusId() string {
	if m != nil {
		return m.MenusId
	}
	return ""
}

//
//DeleteMenu返回
type DeleteMenuResponse struct {
	//
	//菜单id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMenuResponse) Reset()         { *m = DeleteMenuResponse{} }
func (m *DeleteMenuResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMenuResponse) ProtoMessage()    {}
func (*DeleteMenuResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72bfa346086d8ba9, []int{1}
}
func (m *DeleteMenuResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMenuResponse.Unmarshal(m, b)
}
func (m *DeleteMenuResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMenuResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMenuResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMenuResponse.Merge(m, src)
}
func (m *DeleteMenuResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMenuResponse.Size(m)
}
func (m *DeleteMenuResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMenuResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMenuResponse proto.InternalMessageInfo

func (m *DeleteMenuResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DeleteMenuApi返回
type DeleteMenuResponseWrapper struct {
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
	Data                 *DeleteMenuResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeleteMenuResponseWrapper) Reset()         { *m = DeleteMenuResponseWrapper{} }
func (m *DeleteMenuResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DeleteMenuResponseWrapper) ProtoMessage()    {}
func (*DeleteMenuResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_72bfa346086d8ba9, []int{2}
}
func (m *DeleteMenuResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMenuResponseWrapper.Unmarshal(m, b)
}
func (m *DeleteMenuResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMenuResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DeleteMenuResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMenuResponseWrapper.Merge(m, src)
}
func (m *DeleteMenuResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DeleteMenuResponseWrapper.Size(m)
}
func (m *DeleteMenuResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMenuResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMenuResponseWrapper proto.InternalMessageInfo

func (m *DeleteMenuResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteMenuResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DeleteMenuResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DeleteMenuResponseWrapper) GetData() *DeleteMenuResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteMenuRequest)(nil), "menu.DeleteMenuRequest")
	proto.RegisterType((*DeleteMenuResponse)(nil), "menu.DeleteMenuResponse")
	proto.RegisterType((*DeleteMenuResponseWrapper)(nil), "menu.DeleteMenuResponseWrapper")
}

func init() { proto.RegisterFile("delete_menu.proto", fileDescriptor_72bfa346086d8ba9) }

var fileDescriptor_72bfa346086d8ba9 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x4d, 0x95, 0x6e, 0xfd, 0xd7, 0x3d, 0x48, 0x2c, 0x48, 0xca, 0x0a, 0xd2, 0x4b,
	0x1b, 0xb0, 0x22, 0x22, 0x78, 0x29, 0x7a, 0xe8, 0xc1, 0xcb, 0x5e, 0x3c, 0x4a, 0xda, 0xdd, 0xc6,
	0xc5, 0x24, 0x13, 0x37, 0x1b, 0xdb, 0xa7, 0x8d, 0xe8, 0x23, 0xe4, 0x09, 0x24, 0xb3, 0xad, 0x06,
	0x7a, 0xda, 0x99, 0xf9, 0xbe, 0xdf, 0xec, 0xc7, 0x2e, 0xe9, 0x09, 0x19, 0x4b, 0x23, 0x5f, 0x13,
	0x99, 0x16, 0xe3, 0x4c, 0x83, 0x01, 0xea, 0xd6, 0x75, 0x7f, 0x14, 0x29, 0xf3, 0x56, 0xcc, 0xc7,
	0x0b, 0x48, 0x82, 0x08, 0x22, 0x08, 0x50, 0x9c, 0x17, 0x4b, 0xec, 0xb0, 0xc1, 0xca, 0x42, 0xfd,
	0xdb, 0x86, 0x3d, 0x59, 0x29, 0xf3, 0x0e, 0xab, 0x20, 0x82, 0x11, 0x8a, 0xa3, 0xcf, 0x30, 0x56,
	0x22, 0x34, 0xa0, 0xf3, 0xe0, 0xaf, 0xb4, 0x1c, 0x9b, 0x91, 0xde, 0x23, 0x26, 0x78, 0x96, 0x69,
	0xc1, 0xe5, 0x47, 0x21, 0x73, 0x43, 0x6f, 0xc8, 0x41, 0x9d, 0x21, 0x9f, 0x09, 0xcf, 0x19, 0x38,
	0xc3, 0xce, 0xb4, 0x5f, 0x95, 0xfe, 0xf1, 0x12, 0x74, 0x72, 0xcf, 0x36, 0x02, 0xfb, 0xf9, 0xf2,
	0xdd, 0xcc, 0x59, 0x33, 0xbe, 0xb5, 0xb2, 0x09, 0xa1, 0xcd, 0x55, 0x79, 0x06, 0x69, 0x2e, 0xe9,
	0x05, 0x69, 0xa9, 0xed, 0x9a, 0xa3, 0xaa, 0xf4, 0x3b, 0x76, 0x8d, 0x12, 0x8c, 0xb7, 0x94, 0x60,
	0xdf, 0x0e, 0x39, 0xdf, 0xa5, 0x5e, 0x74, 0x98, 0x65, 0x52, 0xd3, 0x4b, 0xe2, 0x2e, 0x40, 0x48,
	0xc4, 0xdb, 0xd3, 0x93, 0xaa, 0xf4, 0xbb, 0x16, 0xaf, 0xa7, 0x8c, 0xa3, 0x48, 0xef, 0x48, 0xb7,
	0x3e, 0x9f, 0xd6, 0x59, 0x1c, 0xaa, 0xd4, 0x6b, 0xe1, 0x55, 0x67, 0x55, 0xe9, 0xd3, 0x7f, 0xef,
	0x46, 0x64, 0xbc, 0x69, 0xa5, 0x57, 0xa4, 0x2d, 0xb5, 0x06, 0xed, 0xed, 0x21, 0x73, 0x5a, 0x95,
	0xfe, 0xa1, 0x65, 0x70, 0xcc, 0xb8, 0x95, 0xe9, 0x03, 0x71, 0x45, 0x68, 0x42, 0xcf, 0x1d, 0x38,
	0xc3, 0xee, 0xb5, 0x37, 0xc6, 0xcf, 0xda, 0x4d, 0xdd, 0x0c, 0x58, 0xfb, 0x19, 0x47, 0x6c, 0xbe,
	0x8f, 0x4f, 0x3d, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x31, 0xcd, 0x9e, 0xdb, 0xec, 0x01, 0x00,
	0x00,
}
