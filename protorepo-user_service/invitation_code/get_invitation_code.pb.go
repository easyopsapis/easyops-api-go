// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_invitation_code.proto

package invitation_code

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
//GetInvitationCode返回
type GetInvitationCodeResponse struct {
	//
	//邀请码
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//org
	Org int32 `protobuf:"varint,2,opt,name=org,proto3" json:"org" form:"org"`
	//
	//过期时间
	Expires string `protobuf:"bytes,3,opt,name=expires,proto3" json:"expires" form:"expires"`
	//
	//是否可用
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state" form:"state"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInvitationCodeResponse) Reset()         { *m = GetInvitationCodeResponse{} }
func (m *GetInvitationCodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetInvitationCodeResponse) ProtoMessage()    {}
func (*GetInvitationCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09754aa85b05004, []int{0}
}
func (m *GetInvitationCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvitationCodeResponse.Unmarshal(m, b)
}
func (m *GetInvitationCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvitationCodeResponse.Marshal(b, m, deterministic)
}
func (m *GetInvitationCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvitationCodeResponse.Merge(m, src)
}
func (m *GetInvitationCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetInvitationCodeResponse.Size(m)
}
func (m *GetInvitationCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvitationCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvitationCodeResponse proto.InternalMessageInfo

func (m *GetInvitationCodeResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetInvitationCodeResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetInvitationCodeResponse) GetExpires() string {
	if m != nil {
		return m.Expires
	}
	return ""
}

func (m *GetInvitationCodeResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

//
//GetInvitationCodeApi返回
type GetInvitationCodeResponseWrapper struct {
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
	Data                 *GetInvitationCodeResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetInvitationCodeResponseWrapper) Reset()         { *m = GetInvitationCodeResponseWrapper{} }
func (m *GetInvitationCodeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetInvitationCodeResponseWrapper) ProtoMessage()    {}
func (*GetInvitationCodeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09754aa85b05004, []int{1}
}
func (m *GetInvitationCodeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvitationCodeResponseWrapper.Unmarshal(m, b)
}
func (m *GetInvitationCodeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvitationCodeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetInvitationCodeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvitationCodeResponseWrapper.Merge(m, src)
}
func (m *GetInvitationCodeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetInvitationCodeResponseWrapper.Size(m)
}
func (m *GetInvitationCodeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvitationCodeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvitationCodeResponseWrapper proto.InternalMessageInfo

func (m *GetInvitationCodeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetInvitationCodeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetInvitationCodeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetInvitationCodeResponseWrapper) GetData() *GetInvitationCodeResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetInvitationCodeResponse)(nil), "invitation_code.GetInvitationCodeResponse")
	proto.RegisterType((*GetInvitationCodeResponseWrapper)(nil), "invitation_code.GetInvitationCodeResponseWrapper")
}

func init() { proto.RegisterFile("get_invitation_code.proto", fileDescriptor_b09754aa85b05004) }

var fileDescriptor_b09754aa85b05004 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xa9, 0xdb, 0x94, 0x65, 0xb2, 0x49, 0x0e, 0xd2, 0x79, 0x69, 0x89, 0x20, 0x43, 0x74,
	0x03, 0xbd, 0x88, 0xc7, 0x89, 0x88, 0x27, 0x21, 0x17, 0x8f, 0x23, 0x5b, 0xb3, 0x18, 0x70, 0xfd,
	0x87, 0x34, 0x93, 0x7d, 0x34, 0x3f, 0x4d, 0x3f, 0x44, 0x8f, 0x9e, 0x24, 0xff, 0x54, 0xad, 0x83,
	0x9e, 0xda, 0xbc, 0xf7, 0xcb, 0xe3, 0x3d, 0x42, 0xc6, 0x4a, 0xba, 0x85, 0xce, 0x3f, 0xb4, 0x13,
	0x4e, 0x43, 0xbe, 0x58, 0x41, 0x26, 0xa7, 0xc6, 0x82, 0x03, 0x3a, 0xda, 0x93, 0xcf, 0xae, 0x95,
	0x76, 0x6f, 0xdb, 0xe5, 0x74, 0x05, 0x9b, 0x99, 0x02, 0x05, 0x33, 0xe4, 0x96, 0xdb, 0x35, 0x9e,
	0xf0, 0x80, 0x7f, 0xe1, 0x3e, 0xfb, 0x8c, 0xc8, 0xf8, 0x49, 0xba, 0xe7, 0xdf, 0x94, 0x07, 0xc8,
	0x24, 0x97, 0x85, 0x81, 0xbc, 0x90, 0xf4, 0x9c, 0x74, 0x7d, 0x68, 0x1c, 0xa5, 0xd1, 0xa4, 0x3f,
	0x1f, 0x55, 0x65, 0x32, 0x58, 0x83, 0xdd, 0xdc, 0x33, 0xaf, 0x32, 0x8e, 0x26, 0x4d, 0x49, 0x07,
	0xac, 0x8a, 0x0f, 0xd2, 0x68, 0xd2, 0x9b, 0x0f, 0xab, 0x32, 0x21, 0x81, 0x01, 0xab, 0x18, 0xf7,
	0x16, 0xbd, 0x22, 0x47, 0x72, 0x67, 0xb4, 0x95, 0x45, 0xdc, 0xc1, 0x24, 0x5a, 0x95, 0xc9, 0x30,
	0x50, 0xb5, 0xc1, 0xf8, 0x0f, 0x42, 0x2f, 0x48, 0xaf, 0x70, 0xc2, 0xc9, 0xb8, 0x8b, 0xec, 0x49,
	0x55, 0x26, 0xc7, 0x81, 0x45, 0x99, 0xf1, 0x60, 0xb3, 0xaf, 0x88, 0xa4, 0xad, 0xd5, 0x5f, 0xad,
	0x30, 0x46, 0xda, 0x7f, 0x0b, 0x7a, 0x6d, 0x0b, 0xee, 0xc8, 0xc0, 0x7f, 0x1f, 0x77, 0xe6, 0x5d,
	0xe8, 0x1c, 0x97, 0xf4, 0xe7, 0xa7, 0x55, 0x99, 0xd0, 0x3f, 0xb6, 0x36, 0x19, 0x6f, 0xa2, 0xbe,
	0xab, 0xb4, 0x16, 0x6c, 0xbd, 0xab, 0xd1, 0x15, 0x65, 0xc6, 0x83, 0x4d, 0x5f, 0x48, 0x37, 0x13,
	0x4e, 0xe0, 0xa4, 0xc1, 0xcd, 0xe5, 0x74, 0xff, 0x31, 0x5b, 0x77, 0x34, 0x2b, 0xfb, 0x04, 0xc6,
	0x31, 0x68, 0x79, 0x88, 0xcf, 0x77, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x62, 0x71, 0x47, 0xb4,
	0x1b, 0x02, 0x00, 0x00,
}
