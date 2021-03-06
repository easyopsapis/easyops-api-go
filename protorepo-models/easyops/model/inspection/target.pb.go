// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: target.proto

package inspection

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
//巡检对象
type InspectionTarget struct {
	//
	//巡检对象实例ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//由唯一键的值拼接得到的实例名
	//如: sid(pid xxx)
	//
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//唯一键的值
	Keys []*InspectionTarget_Keys `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys" form:"keys"`
	//
	//评分(精确到小数点后2位)
	Score float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score" form:"score"`
	//
	//unexecuted:脚本未执行; failed:返回码大于0; normal:返回码等于0,没有异常指标; abnormal:返回码等于0,有异常指标
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTarget) Reset()         { *m = InspectionTarget{} }
func (m *InspectionTarget) String() string { return proto.CompactTextString(m) }
func (*InspectionTarget) ProtoMessage()    {}
func (*InspectionTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{0}
}
func (m *InspectionTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTarget.Unmarshal(m, b)
}
func (m *InspectionTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTarget.Marshal(b, m, deterministic)
}
func (m *InspectionTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTarget.Merge(m, src)
}
func (m *InspectionTarget) XXX_Size() int {
	return xxx_messageInfo_InspectionTarget.Size(m)
}
func (m *InspectionTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTarget.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTarget proto.InternalMessageInfo

func (m *InspectionTarget) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *InspectionTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InspectionTarget) GetKeys() []*InspectionTarget_Keys {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *InspectionTarget) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *InspectionTarget) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type InspectionTarget_Keys struct {
	//
	//唯一键的id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//唯一键的值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	//
	//唯一键的名字
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectionTarget_Keys) Reset()         { *m = InspectionTarget_Keys{} }
func (m *InspectionTarget_Keys) String() string { return proto.CompactTextString(m) }
func (*InspectionTarget_Keys) ProtoMessage()    {}
func (*InspectionTarget_Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{0, 0}
}
func (m *InspectionTarget_Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectionTarget_Keys.Unmarshal(m, b)
}
func (m *InspectionTarget_Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectionTarget_Keys.Marshal(b, m, deterministic)
}
func (m *InspectionTarget_Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectionTarget_Keys.Merge(m, src)
}
func (m *InspectionTarget_Keys) XXX_Size() int {
	return xxx_messageInfo_InspectionTarget_Keys.Size(m)
}
func (m *InspectionTarget_Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectionTarget_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_InspectionTarget_Keys proto.InternalMessageInfo

func (m *InspectionTarget_Keys) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InspectionTarget_Keys) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *InspectionTarget_Keys) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*InspectionTarget)(nil), "inspection.InspectionTarget")
	proto.RegisterType((*InspectionTarget_Keys)(nil), "inspection.InspectionTarget.Keys")
}

func init() { proto.RegisterFile("target.proto", fileDescriptor_468528a86129e532) }

var fileDescriptor_468528a86129e532 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x26, 0xb6, 0x53, 0xa8, 0xda, 0xac, 0xa9, 0xae, 0x4c, 0x60, 0x38, 0x53, 0xcb, 0x70, 0x28,
	0x96, 0xb6, 0x15, 0x06, 0xdd, 0xcd, 0x20, 0x8c, 0x42, 0xd8, 0x9d, 0xd9, 0xd5, 0x96, 0x15, 0x14,
	0x5b, 0xf5, 0x4c, 0x6d, 0x1f, 0x23, 0xc9, 0x2d, 0xd9, 0xd8, 0xdb, 0xec, 0xb9, 0x3c, 0xd8, 0x23,
	0xf8, 0x09, 0x86, 0x65, 0x27, 0x0e, 0xbb, 0xd8, 0xdd, 0x39, 0xdf, 0xcf, 0xf1, 0xe7, 0x0f, 0xa1,
	0x53, 0xcd, 0x65, 0x22, 0x34, 0x2d, 0x25, 0x68, 0xc0, 0x28, 0x2d, 0x54, 0x29, 0x22, 0x9d, 0x42,
	0x31, 0x0b, 0x92, 0x54, 0x7f, 0xab, 0x36, 0x34, 0x82, 0x9c, 0x25, 0x90, 0x00, 0x33, 0x92, 0x4d,
	0x75, 0x6f, 0x36, 0xb3, 0x98, 0xa9, 0xb3, 0xce, 0xde, 0x1e, 0xc8, 0xf3, 0xa7, 0x54, 0x3f, 0xc0,
	0x13, 0x4b, 0x20, 0x30, 0x64, 0xf0, 0xc8, 0xb3, 0x34, 0xe6, 0x1a, 0xa4, 0x62, 0xfb, 0xb1, 0xf3,
	0x91, 0x5f, 0x36, 0x9a, 0xae, 0xf6, 0x5f, 0xfd, 0x64, 0xd2, 0xe0, 0x15, 0x6a, 0x93, 0x68, 0x5e,
	0x44, 0x62, 0x15, 0xbb, 0xa3, 0xf9, 0xc8, 0x3f, 0x5e, 0x2e, 0x9a, 0xda, 0x3b, 0xbf, 0x07, 0x99,
	0xbf, 0x23, 0x03, 0x47, 0xfe, 0xfc, 0xf6, 0xa6, 0xe8, 0xd9, 0xdd, 0x97, 0x57, 0xc1, 0x0d, 0x0f,
	0xbe, 0x7f, 0xfd, 0xf1, 0xfa, 0xfa, 0xe7, 0x65, 0x78, 0x60, 0xc6, 0x17, 0xc8, 0x29, 0x78, 0x2e,
	0x5c, 0xcb, 0x1c, 0x39, 0x6b, 0x6a, 0xef, 0xa4, 0x3b, 0xd2, 0xa2, 0x24, 0x34, 0x24, 0xbe, 0x45,
	0xce, 0x83, 0xd8, 0x2a, 0xd7, 0x9e, 0xdb, 0xfe, 0xc9, 0x9b, 0x17, 0x74, 0xa8, 0x81, 0xfe, 0x9b,
	0x8d, 0x7e, 0x14, 0x5b, 0x75, 0x78, 0xa7, 0x35, 0x92, 0xd0, 0xf8, 0xf1, 0x0d, 0x1a, 0xab, 0x08,
	0xa4, 0x70, 0x9d, 0xf9, 0xc8, 0xb7, 0x96, 0x17, 0x4d, 0xed, 0x9d, 0x76, 0x2a, 0x03, 0xb7, 0x69,
	0xcf, 0xd0, 0xe4, 0x6e, 0x1d, 0x5f, 0xf9, 0x6b, 0xba, 0x8e, 0xaf, 0x16, 0xef, 0x2f, 0xc3, 0xce,
	0x81, 0x17, 0xe8, 0x48, 0x69, 0xae, 0x2b, 0xe5, 0x8e, 0x4d, 0xd2, 0xf3, 0xa6, 0xf6, 0x26, 0xbd,
	0xd7, 0xe0, 0x24, 0xec, 0x05, 0x33, 0x89, 0x9c, 0x36, 0x04, 0x7e, 0x8e, 0xac, 0x74, 0xd7, 0xce,
	0xa4, 0xa9, 0xbd, 0xe3, 0xbe, 0x9d, 0x98, 0x84, 0x56, 0x1a, 0xe3, 0x97, 0x68, 0xfc, 0xc8, 0xb3,
	0x6a, 0xf7, 0xeb, 0xd3, 0x21, 0x8c, 0x81, 0x49, 0xd8, 0xd1, 0xfb, 0x86, 0xec, 0xff, 0x34, 0xb4,
	0xbc, 0xfd, 0xfc, 0x21, 0x01, 0x2a, 0xb8, 0xda, 0x42, 0xa9, 0x68, 0x06, 0x11, 0xcf, 0x58, 0x04,
	0x85, 0x96, 0x3c, 0xd2, 0xaa, 0x7b, 0x1a, 0x52, 0x94, 0x10, 0xe4, 0x10, 0x8b, 0x4c, 0xb1, 0x5e,
	0xc8, 0xcc, 0xca, 0x86, 0x3a, 0x37, 0x47, 0x46, 0x7a, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x69,
	0xd9, 0xe0, 0x14, 0x78, 0x02, 0x00, 0x00,
}
