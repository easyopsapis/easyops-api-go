// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mail_info.proto

package ops_automation

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//邮件通知配置详情
type MailInfo struct {
	//
	//通知哪些人员
	Notifiers []string `protobuf:"bytes,1,rep,name=notifiers,proto3" json:"notifiers" form:"notifiers"`
	//
	//通知方式
	Strategy             string   `protobuf:"bytes,2,opt,name=strategy,proto3" json:"strategy" form:"strategy"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MailInfo) Reset()         { *m = MailInfo{} }
func (m *MailInfo) String() string { return proto.CompactTextString(m) }
func (*MailInfo) ProtoMessage()    {}
func (*MailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a3a4b11b464364, []int{0}
}
func (m *MailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailInfo.Unmarshal(m, b)
}
func (m *MailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailInfo.Marshal(b, m, deterministic)
}
func (m *MailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailInfo.Merge(m, src)
}
func (m *MailInfo) XXX_Size() int {
	return xxx_messageInfo_MailInfo.Size(m)
}
func (m *MailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MailInfo proto.InternalMessageInfo

func (m *MailInfo) GetNotifiers() []string {
	if m != nil {
		return m.Notifiers
	}
	return nil
}

func (m *MailInfo) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func init() {
	proto.RegisterType((*MailInfo)(nil), "ops_automation.MailInfo")
}

func init() { proto.RegisterFile("mail_info.proto", fileDescriptor_49a3a4b11b464364) }

var fileDescriptor_49a3a4b11b464364 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x18, 0x84, 0x89, 0x82, 0x34, 0x39, 0x58, 0x89, 0x97, 0xd2, 0x4b, 0x6a, 0xf4, 0xd0, 0x83, 0x9b,
	0x55, 0x0b, 0x82, 0x7a, 0x32, 0x27, 0x15, 0xbc, 0xf4, 0x66, 0x4b, 0x0d, 0x9b, 0x34, 0x59, 0x17,
	0x77, 0xf3, 0x87, 0xdd, 0x3f, 0x96, 0x2a, 0xbe, 0x86, 0x8f, 0x57, 0xc1, 0x47, 0xe8, 0x13, 0x08,
	0x1b, 0x8d, 0xf6, 0x36, 0xc3, 0x7c, 0x33, 0x87, 0xf1, 0xba, 0x8a, 0x09, 0x99, 0x88, 0xb2, 0x80,
	0xa8, 0xd2, 0x80, 0xe0, 0xef, 0x42, 0x65, 0x12, 0x56, 0x23, 0x28, 0x86, 0x02, 0xca, 0x3e, 0xe1,
	0x02, 0x9f, 0xea, 0x34, 0xca, 0x40, 0x51, 0x0e, 0x1c, 0xa8, 0xc5, 0xd2, 0xba, 0xb0, 0xce, 0x1a,
	0xab, 0x9a, 0x7a, 0xff, 0xfc, 0x1f, 0xae, 0x16, 0x02, 0x9f, 0x61, 0x41, 0x39, 0x10, 0x1b, 0x92,
	0x17, 0x26, 0xc5, 0x9c, 0x21, 0x68, 0x43, 0x5b, 0xd9, 0xf4, 0xc2, 0x0f, 0xc7, 0xeb, 0xdc, 0x33,
	0x21, 0x6f, 0xcb, 0x02, 0xfc, 0x07, 0xcf, 0x2d, 0x01, 0x45, 0x21, 0x72, 0x6d, 0x7a, 0xce, 0x60,
	0x7b, 0xe8, 0xc6, 0x57, 0xeb, 0x55, 0xb0, 0x57, 0x80, 0x56, 0x97, 0x61, 0x1b, 0x85, 0x5f, 0x9f,
	0xc1, 0xa1, 0x77, 0xf0, 0x38, 0x65, 0xe4, 0xf5, 0x9a, 0x4c, 0x4e, 0xc8, 0xc5, 0x6c, 0x1a, 0xb5,
	0x3a, 0x21, 0xb3, 0xb7, 0xb3, 0xe3, 0xd1, 0xe9, 0xfb, 0xd1, 0xf8, 0x6f, 0xcd, 0xa7, 0x5e, 0xc7,
	0xa0, 0x66, 0x98, 0xf3, 0x65, 0x6f, 0x6b, 0xe0, 0x0c, 0xdd, 0x78, 0x7f, 0xbd, 0x0a, 0xba, 0xcd,
	0xf2, 0x6f, 0x12, 0x8e, 0x5b, 0x28, 0xbe, 0x9b, 0xdc, 0x70, 0x88, 0x72, 0x66, 0x96, 0x50, 0x99,
	0x48, 0x42, 0xc6, 0x24, 0xcd, 0xa0, 0x44, 0xcd, 0x32, 0x34, 0xcd, 0x19, 0x3a, 0xaf, 0x80, 0x28,
	0x98, 0xe7, 0xd2, 0xd0, 0x1f, 0x90, 0x5a, 0x4b, 0x37, 0xbf, 0x4c, 0x77, 0x2c, 0x3e, 0xfa, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x68, 0x0f, 0x01, 0x75, 0x01, 0x00, 0x00,
}
