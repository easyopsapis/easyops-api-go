// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reset_invitation_code.proto

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
//ResetInvitationCode返回
type ResetInvitationCodeResponse struct {
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

func (m *ResetInvitationCodeResponse) Reset()         { *m = ResetInvitationCodeResponse{} }
func (m *ResetInvitationCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ResetInvitationCodeResponse) ProtoMessage()    {}
func (*ResetInvitationCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc869bd25034577, []int{0}
}
func (m *ResetInvitationCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetInvitationCodeResponse.Unmarshal(m, b)
}
func (m *ResetInvitationCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetInvitationCodeResponse.Marshal(b, m, deterministic)
}
func (m *ResetInvitationCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetInvitationCodeResponse.Merge(m, src)
}
func (m *ResetInvitationCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ResetInvitationCodeResponse.Size(m)
}
func (m *ResetInvitationCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetInvitationCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResetInvitationCodeResponse proto.InternalMessageInfo

func (m *ResetInvitationCodeResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ResetInvitationCodeResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *ResetInvitationCodeResponse) GetExpires() string {
	if m != nil {
		return m.Expires
	}
	return ""
}

func (m *ResetInvitationCodeResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

//
//ResetInvitationCodeApi返回
type ResetInvitationCodeResponseWrapper struct {
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
	Data                 *ResetInvitationCodeResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResetInvitationCodeResponseWrapper) Reset()         { *m = ResetInvitationCodeResponseWrapper{} }
func (m *ResetInvitationCodeResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ResetInvitationCodeResponseWrapper) ProtoMessage()    {}
func (*ResetInvitationCodeResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc869bd25034577, []int{1}
}
func (m *ResetInvitationCodeResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetInvitationCodeResponseWrapper.Unmarshal(m, b)
}
func (m *ResetInvitationCodeResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetInvitationCodeResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ResetInvitationCodeResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetInvitationCodeResponseWrapper.Merge(m, src)
}
func (m *ResetInvitationCodeResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ResetInvitationCodeResponseWrapper.Size(m)
}
func (m *ResetInvitationCodeResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetInvitationCodeResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ResetInvitationCodeResponseWrapper proto.InternalMessageInfo

func (m *ResetInvitationCodeResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResetInvitationCodeResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ResetInvitationCodeResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResetInvitationCodeResponseWrapper) GetData() *ResetInvitationCodeResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ResetInvitationCodeResponse)(nil), "invitation_code.ResetInvitationCodeResponse")
	proto.RegisterType((*ResetInvitationCodeResponseWrapper)(nil), "invitation_code.ResetInvitationCodeResponseWrapper")
}

func init() { proto.RegisterFile("reset_invitation_code.proto", fileDescriptor_acc869bd25034577) }

var fileDescriptor_acc869bd25034577 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xa9, 0xdb, 0x94, 0x65, 0xb2, 0x49, 0x0e, 0x52, 0xdc, 0xa1, 0x25, 0x82, 0xec, 0x30,
	0x3b, 0xd0, 0x8b, 0x78, 0xac, 0x78, 0xf0, 0x68, 0x2e, 0x1e, 0x47, 0xba, 0x66, 0x35, 0xe0, 0xfa,
	0x0f, 0x69, 0x26, 0xfb, 0x6c, 0x7e, 0x98, 0x7e, 0x88, 0x7e, 0x00, 0x91, 0xfc, 0x53, 0xb5, 0x0e,
	0xdc, 0xa9, 0xcd, 0x7b, 0xbf, 0x3c, 0xde, 0x23, 0x64, 0x6a, 0x64, 0x25, 0xed, 0x52, 0x95, 0xef,
	0xca, 0x0a, 0xab, 0xa0, 0x5c, 0xae, 0x20, 0x97, 0x89, 0x36, 0x60, 0x81, 0x4e, 0xf6, 0xe4, 0x8b,
	0xeb, 0x42, 0xd9, 0xd7, 0x6d, 0x96, 0xac, 0x60, 0xb3, 0x28, 0xa0, 0x80, 0x05, 0x72, 0xd9, 0x76,
	0x8d, 0x27, 0x3c, 0xe0, 0x9f, 0xbf, 0xcf, 0x3e, 0x02, 0x32, 0xe5, 0x2e, 0xff, 0xe9, 0x27, 0xe7,
	0x01, 0x72, 0xc9, 0x65, 0xa5, 0xa1, 0xac, 0x24, 0xbd, 0x24, 0x7d, 0x17, 0x1b, 0x06, 0x71, 0x30,
	0x1b, 0xa6, 0x93, 0xa6, 0x8e, 0x46, 0x6b, 0x30, 0x9b, 0x7b, 0xe6, 0x54, 0xc6, 0xd1, 0xa4, 0x31,
	0xe9, 0x81, 0x29, 0xc2, 0xa3, 0x38, 0x98, 0x0d, 0xd2, 0x71, 0x53, 0x47, 0xc4, 0x33, 0x60, 0x0a,
	0xc6, 0x9d, 0x45, 0xe7, 0xe4, 0x44, 0xee, 0xb4, 0x32, 0xb2, 0x0a, 0x7b, 0x98, 0x44, 0x9b, 0x3a,
	0x1a, 0x7b, 0xaa, 0x35, 0x18, 0xff, 0x46, 0xe8, 0x15, 0x19, 0x54, 0x56, 0x58, 0x19, 0xf6, 0x91,
	0x3d, 0x6b, 0xea, 0xe8, 0xd4, 0xb3, 0x28, 0x33, 0xee, 0x6d, 0xf6, 0x19, 0x10, 0x76, 0xa0, 0xfc,
	0x8b, 0x11, 0x5a, 0x4b, 0xf3, 0x67, 0xc3, 0xe0, 0xbf, 0x0d, 0x77, 0x64, 0xe4, 0xbe, 0x8f, 0x3b,
	0xfd, 0x26, 0x54, 0x89, 0x5b, 0x86, 0xe9, 0x79, 0x53, 0x47, 0xf4, 0x97, 0x6d, 0x4d, 0xc6, 0xbb,
	0xa8, 0x6b, 0x2b, 0x8d, 0x01, 0xd3, 0x2e, 0xeb, 0xb4, 0x45, 0x99, 0x71, 0x6f, 0xd3, 0x67, 0xd2,
	0xcf, 0x85, 0x15, 0x38, 0x6a, 0x74, 0x33, 0x4f, 0xf6, 0x1f, 0xf4, 0xc0, 0x92, 0x6e, 0x69, 0x97,
	0xc1, 0x38, 0x46, 0x65, 0xc7, 0xf8, 0x88, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x96,
	0xf3, 0xaa, 0x23, 0x02, 0x00, 0x00,
}
