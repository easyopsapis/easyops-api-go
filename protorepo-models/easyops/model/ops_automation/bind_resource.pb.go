// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bind_resource.proto

package ops_automation

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//绑定资源
type BindResource struct {
	//
	//绑定类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//资源ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//最新版本ID
	VId string `protobuf:"bytes,3,opt,name=vId,proto3" json:"vId" form:"vId"`
	//
	//工具的输入,固定@agent。{@agents:[{'ip':xxxx, 'instanceId':xxxx}], ....}
	DefaultInputs        *types.Struct `protobuf:"bytes,4,opt,name=defaultInputs,proto3" json:"defaultInputs" form:"defaultInputs"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BindResource) Reset()         { *m = BindResource{} }
func (m *BindResource) String() string { return proto.CompactTextString(m) }
func (*BindResource) ProtoMessage()    {}
func (*BindResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1547b9ee722bb261, []int{0}
}
func (m *BindResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindResource.Unmarshal(m, b)
}
func (m *BindResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindResource.Marshal(b, m, deterministic)
}
func (m *BindResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindResource.Merge(m, src)
}
func (m *BindResource) XXX_Size() int {
	return xxx_messageInfo_BindResource.Size(m)
}
func (m *BindResource) XXX_DiscardUnknown() {
	xxx_messageInfo_BindResource.DiscardUnknown(m)
}

var xxx_messageInfo_BindResource proto.InternalMessageInfo

func (m *BindResource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BindResource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BindResource) GetVId() string {
	if m != nil {
		return m.VId
	}
	return ""
}

func (m *BindResource) GetDefaultInputs() *types.Struct {
	if m != nil {
		return m.DefaultInputs
	}
	return nil
}

func init() {
	proto.RegisterType((*BindResource)(nil), "ops_automation.BindResource")
}

func init() { proto.RegisterFile("bind_resource.proto", fileDescriptor_1547b9ee722bb261) }

var fileDescriptor_1547b9ee722bb261 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x49, 0x5b, 0x2e, 0x34, 0xbd, 0xbd, 0x17, 0xa2, 0x62, 0x28, 0x42, 0x4a, 0x44, 0xe8,
	0x26, 0x19, 0x54, 0x70, 0xe1, 0x32, 0x2b, 0xeb, 0x32, 0x2e, 0x04, 0x37, 0x65, 0x92, 0x99, 0x8e,
	0x83, 0x49, 0xbe, 0x61, 0xfe, 0xb4, 0xf6, 0x65, 0x23, 0xb8, 0xf0, 0x01, 0xf2, 0x02, 0x4a, 0x67,
	0x5a, 0xb5, 0x59, 0x7d, 0x5f, 0xce, 0xef, 0x1c, 0xe6, 0xcc, 0xf8, 0x47, 0x05, 0x6f, 0xc8, 0x42,
	0x52, 0x05, 0x46, 0x96, 0x34, 0x15, 0x12, 0x34, 0x04, 0xff, 0x40, 0xa8, 0x05, 0x36, 0x1a, 0x6a,
	0xac, 0x39, 0x34, 0x93, 0x84, 0x71, 0xfd, 0x6c, 0x8a, 0xb4, 0x84, 0x1a, 0x31, 0x60, 0x80, 0x2c,
	0x56, 0x98, 0xa5, 0xdd, 0xec, 0x62, 0x27, 0x67, 0x9f, 0xdc, 0xfc, 0xc2, 0xeb, 0x35, 0xd7, 0x2f,
	0xb0, 0x46, 0x0c, 0x12, 0x2b, 0x26, 0x2b, 0x5c, 0x71, 0x82, 0x35, 0x48, 0x85, 0xbe, 0xc7, 0x9d,
	0xef, 0x8c, 0x01, 0xb0, 0x8a, 0xfe, 0xa4, 0x2b, 0x2d, 0x4d, 0xa9, 0x9d, 0x1a, 0x7f, 0x78, 0xfe,
	0xdf, 0x8c, 0x37, 0x24, 0xdf, 0x9d, 0x35, 0x38, 0xf7, 0x07, 0x7a, 0x23, 0x68, 0xe8, 0x4d, 0xbd,
	0xd9, 0x30, 0xfb, 0xdf, 0xb5, 0xd1, 0x68, 0x09, 0xb2, 0xbe, 0x8d, 0xb7, 0x7f, 0xe3, 0xdc, 0x8a,
	0xc1, 0x85, 0xdf, 0xe3, 0x24, 0xec, 0x59, 0xe4, 0xa4, 0x6b, 0xa3, 0xa1, 0x43, 0x38, 0x89, 0xdf,
	0xdf, 0xa2, 0x81, 0xf0, 0x5e, 0xe3, 0xbc, 0xc7, 0x49, 0x70, 0xe9, 0xf7, 0x57, 0x73, 0x12, 0xf6,
	0x2d, 0x17, 0x75, 0x6d, 0xe4, 0x3b, 0x6e, 0x35, 0xb7, 0xe0, 0x58, 0x7c, 0xee, 0xbf, 0xad, 0x63,
	0xcb, 0x06, 0x8f, 0xfe, 0x98, 0xd0, 0x25, 0x36, 0x95, 0x9e, 0x37, 0xc2, 0x68, 0x15, 0x0e, 0xa6,
	0xde, 0x6c, 0x74, 0x75, 0x9a, 0xba, 0x16, 0xe9, 0xbe, 0x45, 0xfa, 0x60, 0x5b, 0x64, 0x61, 0xd7,
	0x46, 0xc7, 0x2e, 0xf5, 0xc0, 0x17, 0xe7, 0x87, 0x39, 0xd9, 0xfd, 0xd3, 0x1d, 0x83, 0x94, 0x62,
	0xb5, 0x01, 0xa1, 0xd2, 0x0a, 0x4a, 0x5c, 0xa1, 0x12, 0x1a, 0x2d, 0x71, 0xa9, 0x95, 0xbb, 0x1c,
	0x49, 0x05, 0x24, 0x35, 0x10, 0x5a, 0x29, 0xb4, 0x03, 0x91, 0x5d, 0xd1, 0xe1, 0xcb, 0x15, 0x7f,
	0x2c, 0x7e, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x38, 0x22, 0x7e, 0xe7, 0x01, 0x00, 0x00,
}
