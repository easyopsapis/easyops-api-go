// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: action.proto

package easy_command

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//Action规格
type Action struct {
	//
	//创建任务时设置的action名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//action类型，一般不用设置
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	//
	//设置的action动作
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action" form:"action"`
	//
	//设置的action超时时间，单位为秒。设置为0时不超时
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout" form:"timeout"`
	//
	//设置的action参数。见 CmdActionParam FileActionParam
	Param                *types.Struct `protobuf:"bytes,5,opt,name=param,proto3" json:"param" form:"param"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{0}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Action) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Action) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Action) GetParam() *types.Struct {
	if m != nil {
		return m.Param
	}
	return nil
}

func init() {
	proto.RegisterType((*Action)(nil), "easy_command.Action")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor_59885c909ad4dfd3) }

var fileDescriptor_59885c909ad4dfd3 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x68, 0x8b, 0x70, 0xcb, 0x3f, 0x2f, 0x44, 0x15, 0x52, 0x2a, 0xb3, 0x14, 0x89,
	0xda, 0x12, 0x6c, 0x2c, 0x88, 0x0c, 0x88, 0x39, 0x6c, 0x2c, 0xc8, 0x71, 0x5d, 0x13, 0x29, 0xce,
	0x45, 0x8e, 0x33, 0xf4, 0x65, 0xf3, 0x04, 0x4c, 0x79, 0x02, 0x94, 0x73, 0x2a, 0x10, 0x9b, 0xef,
	0xfb, 0x7e, 0x77, 0xf6, 0x77, 0x26, 0x0b, 0xa9, 0x7c, 0x01, 0x15, 0xaf, 0x1d, 0x78, 0xa0, 0x0b,
	0x2d, 0x9b, 0xfd, 0xa7, 0x02, 0x6b, 0x65, 0xb5, 0x5d, 0x6e, 0x4c, 0xe1, 0xbf, 0xda, 0x9c, 0x2b,
	0xb0, 0xc2, 0x80, 0x01, 0x81, 0x50, 0xde, 0xee, 0xb0, 0xc2, 0x02, 0x4f, 0xa1, 0x79, 0x79, 0x63,
	0x00, 0x4c, 0xa9, 0x7f, 0xa9, 0xc6, 0xbb, 0x56, 0xf9, 0xe0, 0xb2, 0xef, 0x88, 0xcc, 0x5e, 0xf0,
	0x2e, 0x7a, 0x4b, 0x26, 0x95, 0xb4, 0x3a, 0x8e, 0x56, 0xd1, 0xfa, 0x34, 0xbd, 0xe8, 0xbb, 0x64,
	0xbe, 0x03, 0x67, 0x9f, 0xd8, 0xa0, 0xb2, 0x0c, 0xcd, 0x01, 0xf2, 0xfb, 0x5a, 0xc7, 0x47, 0xff,
	0xa1, 0x41, 0x65, 0x19, 0x9a, 0xf4, 0x8e, 0xcc, 0xc2, 0xfb, 0xe3, 0x63, 0xc4, 0xae, 0xfa, 0x2e,
	0x39, 0x0b, 0x58, 0xd0, 0x59, 0x36, 0x02, 0xf4, 0x9e, 0x9c, 0xf8, 0xc2, 0x6a, 0x68, 0x7d, 0x3c,
	0x59, 0x45, 0xeb, 0x69, 0x4a, 0xfb, 0x2e, 0x39, 0x1f, 0x47, 0x06, 0x83, 0x65, 0x07, 0x84, 0x3e,
	0x93, 0x69, 0x2d, 0x9d, 0xb4, 0xf1, 0x74, 0x15, 0xad, 0xe7, 0x0f, 0xd7, 0x3c, 0x64, 0xe3, 0x87,
	0x6c, 0xfc, 0x1d, 0xb3, 0xa5, 0x97, 0x7d, 0x97, 0x2c, 0xc2, 0x10, 0xe4, 0x59, 0x16, 0xfa, 0xd2,
	0xb7, 0x8f, 0x57, 0x03, 0x7c, 0x58, 0x27, 0xd4, 0x0d, 0x2f, 0x41, 0xc9, 0x52, 0x28, 0xa8, 0xbc,
	0x93, 0xca, 0x37, 0x61, 0x45, 0x4e, 0xd7, 0xb0, 0xb1, 0xb0, 0xd5, 0x65, 0x23, 0x46, 0x50, 0x60,
	0x29, 0xfe, 0xfe, 0x42, 0x3e, 0x43, 0xf8, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xed, 0x88,
	0x73, 0xaa, 0x01, 0x00, 0x00,
}