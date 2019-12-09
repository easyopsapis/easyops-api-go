// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: action_param_custom.proto

package easy_command

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
//一个action在一个target上的差异化参数
type ActionParamCustom struct {
	//
	//对动作为 "runCmd" 的action，为要运行的脚本的差异化内容，会拼在该action的 CmdActionParam.cmd 之前
	Param                string   `protobuf:"bytes,1,opt,name=param,proto3" json:"param" form:"param"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionParamCustom) Reset()         { *m = ActionParamCustom{} }
func (m *ActionParamCustom) String() string { return proto.CompactTextString(m) }
func (*ActionParamCustom) ProtoMessage()    {}
func (*ActionParamCustom) Descriptor() ([]byte, []int) {
	return fileDescriptor_2da62ea5c5df43cb, []int{0}
}
func (m *ActionParamCustom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionParamCustom.Unmarshal(m, b)
}
func (m *ActionParamCustom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionParamCustom.Marshal(b, m, deterministic)
}
func (m *ActionParamCustom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionParamCustom.Merge(m, src)
}
func (m *ActionParamCustom) XXX_Size() int {
	return xxx_messageInfo_ActionParamCustom.Size(m)
}
func (m *ActionParamCustom) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionParamCustom.DiscardUnknown(m)
}

var xxx_messageInfo_ActionParamCustom proto.InternalMessageInfo

func (m *ActionParamCustom) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionParamCustom)(nil), "easy_command.ActionParamCustom")
}

func init() { proto.RegisterFile("action_param_custom.proto", fileDescriptor_2da62ea5c5df43cb) }

var fileDescriptor_2da62ea5c5df43cb = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4c, 0x2e, 0xc9,
	0xcc, 0xcf, 0x8b, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x8d, 0x4f, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x4d, 0x2c, 0xae, 0x8c, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b,
	0xa2, 0x59, 0xc9, 0x9a, 0x4b, 0xd0, 0x11, 0x6c, 0x72, 0x00, 0xc8, 0x60, 0x67, 0xb0, 0xb9, 0x42,
	0x6a, 0x5c, 0xac, 0x60, 0x7b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x04, 0x3e, 0xdd, 0x93,
	0xe7, 0x49, 0xcb, 0x2f, 0xca, 0xb5, 0x52, 0x02, 0x0b, 0x2b, 0x05, 0x41, 0xa4, 0x9d, 0x3c, 0xa2,
	0xdc, 0xd2, 0xf3, 0xf5, 0x40, 0xd6, 0xe7, 0x17, 0x14, 0xeb, 0xe5, 0xe4, 0x27, 0x27, 0xe6, 0xe8,
	0x27, 0xe7, 0xe7, 0x95, 0x14, 0x25, 0x26, 0x97, 0x14, 0x43, 0x2c, 0x2e, 0x4a, 0x2d, 0xc8, 0xd7,
	0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x87, 0x2a, 0xd4, 0x07, 0x73, 0xf5, 0x91, 0x5d, 0x9d,
	0xc4, 0x06, 0x56, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x57, 0xfb, 0xcd, 0x50, 0xe7, 0x00,
	0x00, 0x00,
}
