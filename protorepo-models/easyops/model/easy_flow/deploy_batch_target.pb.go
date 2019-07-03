// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deploy_batch_target.proto

package easy_flow

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
//分批目标
type DeployBatchTarget struct {
	//
	//目标Id
	TargetId string `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId" form:"targetId"`
	//
	//目标名称
	TargetName string `protobuf:"bytes,2,opt,name=targetName,proto3" json:"targetName" form:"targetName"`
	//
	//集群Id
	ClusterId string `protobuf:"bytes,3,opt,name=clusterId,proto3" json:"clusterId" form:"clusterId"`
	//
	//参数列表
	ActionParams         []*DeployBatchTarget_ActionParams `protobuf:"bytes,4,rep,name=actionParams,proto3" json:"actionParams" form:"actionParams"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DeployBatchTarget) Reset()         { *m = DeployBatchTarget{} }
func (m *DeployBatchTarget) String() string { return proto.CompactTextString(m) }
func (*DeployBatchTarget) ProtoMessage()    {}
func (*DeployBatchTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_8958653d41a9d260, []int{0}
}
func (m *DeployBatchTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployBatchTarget.Unmarshal(m, b)
}
func (m *DeployBatchTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployBatchTarget.Marshal(b, m, deterministic)
}
func (m *DeployBatchTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployBatchTarget.Merge(m, src)
}
func (m *DeployBatchTarget) XXX_Size() int {
	return xxx_messageInfo_DeployBatchTarget.Size(m)
}
func (m *DeployBatchTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployBatchTarget.DiscardUnknown(m)
}

var xxx_messageInfo_DeployBatchTarget proto.InternalMessageInfo

func (m *DeployBatchTarget) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *DeployBatchTarget) GetTargetName() string {
	if m != nil {
		return m.TargetName
	}
	return ""
}

func (m *DeployBatchTarget) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeployBatchTarget) GetActionParams() []*DeployBatchTarget_ActionParams {
	if m != nil {
		return m.ActionParams
	}
	return nil
}

type DeployBatchTarget_ActionParams struct {
	//
	//参数
	Param                string   `protobuf:"bytes,1,opt,name=param,proto3" json:"param" form:"param"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployBatchTarget_ActionParams) Reset()         { *m = DeployBatchTarget_ActionParams{} }
func (m *DeployBatchTarget_ActionParams) String() string { return proto.CompactTextString(m) }
func (*DeployBatchTarget_ActionParams) ProtoMessage()    {}
func (*DeployBatchTarget_ActionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8958653d41a9d260, []int{0, 0}
}
func (m *DeployBatchTarget_ActionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployBatchTarget_ActionParams.Unmarshal(m, b)
}
func (m *DeployBatchTarget_ActionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployBatchTarget_ActionParams.Marshal(b, m, deterministic)
}
func (m *DeployBatchTarget_ActionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployBatchTarget_ActionParams.Merge(m, src)
}
func (m *DeployBatchTarget_ActionParams) XXX_Size() int {
	return xxx_messageInfo_DeployBatchTarget_ActionParams.Size(m)
}
func (m *DeployBatchTarget_ActionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployBatchTarget_ActionParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeployBatchTarget_ActionParams proto.InternalMessageInfo

func (m *DeployBatchTarget_ActionParams) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployBatchTarget)(nil), "easy_flow.DeployBatchTarget")
	proto.RegisterType((*DeployBatchTarget_ActionParams)(nil), "easy_flow.DeployBatchTarget.ActionParams")
}

func init() { proto.RegisterFile("deploy_batch_target.proto", fileDescriptor_8958653d41a9d260) }

var fileDescriptor_8958653d41a9d260 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcf, 0x8b, 0xd3, 0x40,
	0x18, 0x25, 0xb6, 0x8a, 0x1d, 0x0b, 0xb6, 0x11, 0x31, 0xf6, 0x92, 0x12, 0x44, 0xea, 0x61, 0x32,
	0x5a, 0xb5, 0xa8, 0x17, 0x69, 0xd4, 0x43, 0x2f, 0x22, 0xa1, 0xa7, 0x8a, 0x96, 0xc9, 0xcf, 0x86,
	0x4e, 0xfa, 0x85, 0xc9, 0xd4, 0x52, 0x97, 0x1e, 0xf6, 0xb4, 0xb0, 0x7f, 0x64, 0x16, 0xf6, 0x4f,
	0xc8, 0x5f, 0xb0, 0x74, 0xa6, 0x9b, 0x4d, 0xd9, 0xd3, 0x1e, 0xf6, 0xf6, 0x5e, 0xde, 0x7b, 0xdf,
	0xbc, 0x8f, 0x2f, 0xe8, 0x65, 0x10, 0x66, 0x0c, 0xb6, 0x73, 0x8f, 0x0a, 0x7f, 0x31, 0x17, 0x94,
	0xc7, 0xa1, 0xb0, 0x33, 0x0e, 0x02, 0xf4, 0x56, 0x48, 0xf3, 0xed, 0x3c, 0x62, 0xb0, 0xe9, 0xe1,
	0x38, 0x11, 0x8b, 0xb5, 0x67, 0xfb, 0x90, 0x92, 0x18, 0x62, 0x20, 0xd2, 0xe1, 0xad, 0x23, 0xc9,
	0x24, 0x91, 0x48, 0x25, 0x7b, 0xa3, 0x9a, 0x3d, 0xdd, 0x24, 0x62, 0x09, 0x1b, 0x12, 0x03, 0x96,
	0x22, 0xfe, 0x47, 0x59, 0x12, 0x50, 0x01, 0x3c, 0x27, 0x15, 0x54, 0x39, 0xeb, 0xb4, 0x89, 0xba,
	0xdf, 0x65, 0x1f, 0x67, 0x5f, 0x67, 0x2a, 0xdb, 0xe8, 0x67, 0x1a, 0x7a, 0xac, 0x8a, 0x4d, 0x02,
	0x43, 0xeb, 0x6b, 0x83, 0x96, 0xb3, 0x2c, 0x0b, 0xf3, 0x69, 0x04, 0x3c, 0xfd, 0x62, 0x5d, 0x2b,
	0xd6, 0xe5, 0x85, 0x39, 0x45, 0xee, 0xdf, 0xdf, 0x5f, 0x29, 0xfe, 0x3f, 0xc6, 0xb3, 0xb7, 0xf8,
	0xf3, 0x9f, 0x93, 0x4f, 0x3b, 0x7c, 0xc4, 0x3f, 0xdc, 0x91, 0xbf, 0x1b, 0xee, 0x5e, 0xb9, 0xd5,
	0xe3, 0xfa, 0x47, 0x84, 0x14, 0xfe, 0x49, 0xd3, 0xd0, 0x78, 0x20, 0xab, 0x3c, 0x2f, 0x0b, 0xb3,
	0x5b, 0xaf, 0xb2, 0xd7, 0x2c, 0xb7, 0x66, 0xd4, 0xcf, 0x35, 0xd4, 0xf2, 0xd9, 0x3a, 0x17, 0x21,
	0x9f, 0x04, 0x46, 0x43, 0xc6, 0x58, 0x59, 0x98, 0x1d, 0x15, 0xab, 0xa4, 0xfb, 0x5b, 0xe1, 0xe6,
	0x79, 0x3d, 0x42, 0x6d, 0xea, 0x8b, 0x04, 0x56, 0xbf, 0x28, 0xa7, 0x69, 0x6e, 0x34, 0xfb, 0x8d,
	0xc1, 0x93, 0xe1, 0x1b, 0xbb, 0x3a, 0xb6, 0x7d, 0xeb, 0x02, 0xf6, 0xb8, 0x16, 0x70, 0x5e, 0x94,
	0x85, 0xf9, 0x4c, 0x35, 0xaf, 0x0f, 0xb2, 0xdc, 0xa3, 0xb9, 0xbd, 0x11, 0x6a, 0xd7, 0x63, 0xfa,
	0x6b, 0xf4, 0x30, 0xdb, 0xa3, 0xc3, 0x05, 0x3b, 0x65, 0x61, 0xb6, 0xd5, 0x14, 0xf9, 0xd9, 0x72,
	0x95, 0xec, 0xfc, 0x98, 0x7d, 0x8b, 0x41, 0xb6, 0x81, 0x2c, 0xb7, 0x19, 0xf8, 0x94, 0x11, 0x1f,
	0x56, 0x82, 0x53, 0x5f, 0xe4, 0xea, 0xbf, 0xe3, 0x61, 0x06, 0x38, 0x85, 0x20, 0x64, 0x39, 0x39,
	0x18, 0x89, 0xa4, 0xa4, 0x5a, 0xc2, 0x7b, 0x24, 0x9d, 0xef, 0xaf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xbb, 0xa0, 0x24, 0xe0, 0x02, 0x00, 0x00,
}
