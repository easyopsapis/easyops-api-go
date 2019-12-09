// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modify_pipeline.proto

package pipeline

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//ModifyPipeline请求
type ModifyPipelineRequest struct {
	//
	//应用Id
	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//流水线Id
	FlowId string `protobuf:"bytes,2,opt,name=flowId,proto3" json:"flowId" form:"flowId"`
	//
	//流水线信息
	Pipeline             *types.Struct `protobuf:"bytes,3,opt,name=pipeline,proto3" json:"pipeline" form:"pipeline"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ModifyPipelineRequest) Reset()         { *m = ModifyPipelineRequest{} }
func (m *ModifyPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyPipelineRequest) ProtoMessage()    {}
func (*ModifyPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4fa2a5ba77828b1, []int{0}
}
func (m *ModifyPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyPipelineRequest.Unmarshal(m, b)
}
func (m *ModifyPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyPipelineRequest.Marshal(b, m, deterministic)
}
func (m *ModifyPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyPipelineRequest.Merge(m, src)
}
func (m *ModifyPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyPipelineRequest.Size(m)
}
func (m *ModifyPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyPipelineRequest proto.InternalMessageInfo

func (m *ModifyPipelineRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ModifyPipelineRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ModifyPipelineRequest) GetPipeline() *types.Struct {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

//
//ModifyPipelineApi返回
type ModifyPipelineResponseWrapper struct {
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
	Data                 *cmdb_extend.AppPipeline `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ModifyPipelineResponseWrapper) Reset()         { *m = ModifyPipelineResponseWrapper{} }
func (m *ModifyPipelineResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ModifyPipelineResponseWrapper) ProtoMessage()    {}
func (*ModifyPipelineResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4fa2a5ba77828b1, []int{1}
}
func (m *ModifyPipelineResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyPipelineResponseWrapper.Unmarshal(m, b)
}
func (m *ModifyPipelineResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyPipelineResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ModifyPipelineResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyPipelineResponseWrapper.Merge(m, src)
}
func (m *ModifyPipelineResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ModifyPipelineResponseWrapper.Size(m)
}
func (m *ModifyPipelineResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyPipelineResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyPipelineResponseWrapper proto.InternalMessageInfo

func (m *ModifyPipelineResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ModifyPipelineResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ModifyPipelineResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ModifyPipelineResponseWrapper) GetData() *cmdb_extend.AppPipeline {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ModifyPipelineRequest)(nil), "pipeline.ModifyPipelineRequest")
	proto.RegisterType((*ModifyPipelineResponseWrapper)(nil), "pipeline.ModifyPipelineResponseWrapper")
}

func init() { proto.RegisterFile("modify_pipeline.proto", fileDescriptor_c4fa2a5ba77828b1) }

var fileDescriptor_c4fa2a5ba77828b1 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x86, 0x89, 0xce, 0x2e, 0xda, 0xa3, 0xac, 0xb6, 0xac, 0x86, 0x41, 0xc9, 0xd2, 0x82, 0xac,
	0x87, 0xed, 0x80, 0x5e, 0x44, 0xf0, 0xe0, 0x80, 0xe0, 0x1e, 0x04, 0x69, 0x0f, 0x7a, 0x1b, 0x3a,
	0xe9, 0x4e, 0x0c, 0x24, 0xa9, 0xb6, 0xbb, 0x83, 0x33, 0xef, 0xe5, 0xf3, 0xe4, 0xe6, 0x0b, 0xe4,
	0x09, 0x24, 0xd5, 0xc9, 0xcc, 0x38, 0xa7, 0xa4, 0xea, 0xff, 0xff, 0xca, 0x57, 0x15, 0x72, 0xd9,
	0x80, 0xaa, 0x8a, 0xdd, 0xc6, 0x54, 0x46, 0xd7, 0x55, 0xab, 0xb9, 0xb1, 0xe0, 0x81, 0xde, 0x9b,
	0xeb, 0xd5, 0x4d, 0x59, 0xf9, 0x9f, 0x5d, 0xc6, 0x73, 0x68, 0xd2, 0x12, 0x4a, 0x48, 0xd1, 0x90,
	0x75, 0x05, 0x56, 0x58, 0xe0, 0x5b, 0x08, 0xae, 0x7e, 0x94, 0xc0, 0xb5, 0x74, 0x3b, 0x30, 0x8e,
	0xd7, 0x90, 0xcb, 0x3a, 0xcd, 0xa1, 0xf5, 0x56, 0xe6, 0xde, 0x85, 0xa4, 0xd5, 0x06, 0x6e, 0x1a,
	0x50, 0xba, 0x76, 0xe9, 0x64, 0x4c, 0xb1, 0x4c, 0xf3, 0x46, 0x65, 0x1b, 0xbd, 0xf5, 0xba, 0x55,
	0xa9, 0x34, 0xe6, 0x04, 0x69, 0xf5, 0xbc, 0x04, 0x28, 0x6b, 0x7d, 0xf8, 0xbe, 0xf3, 0xb6, 0xcb,
	0x7d, 0x50, 0xd9, 0x9f, 0x88, 0x5c, 0x7e, 0xc1, 0x55, 0xbe, 0x4e, 0x31, 0xa1, 0x7f, 0x75, 0xda,
	0x79, 0xfa, 0x8a, 0x9c, 0x49, 0x63, 0x6e, 0x55, 0x1c, 0x5d, 0x45, 0xd7, 0xf7, 0xd7, 0x8f, 0x86,
	0x3e, 0x79, 0x50, 0x80, 0x6d, 0xde, 0x33, 0x6c, 0x33, 0x11, 0x64, 0xfa, 0x9a, 0x9c, 0x17, 0x35,
	0xfc, 0xbe, 0x55, 0xf1, 0x1d, 0x34, 0x3e, 0x1e, 0xfa, 0xe4, 0x61, 0x30, 0x86, 0x3e, 0x13, 0x93,
	0x81, 0x7e, 0x26, 0xfb, 0xfb, 0xc4, 0x77, 0xaf, 0xa2, 0xeb, 0xe5, 0x9b, 0x67, 0x3c, 0xd0, 0xf1,
	0x99, 0x8e, 0x7f, 0x43, 0xba, 0xf5, 0x93, 0xa1, 0x4f, 0x2e, 0xc2, 0x94, 0x39, 0xc2, 0xc4, 0x3e,
	0xcd, 0xfe, 0x46, 0xe4, 0xc5, 0x29, 0xb6, 0x33, 0xd0, 0x3a, 0xfd, 0xdd, 0x4a, 0x63, 0xb4, 0xa5,
	0x2f, 0xc9, 0x22, 0x07, 0xa5, 0x91, 0xfe, 0x6c, 0x7d, 0x31, 0xf4, 0xc9, 0x32, 0x8c, 0x1b, 0xbb,
	0x4c, 0xa0, 0x48, 0xdf, 0x91, 0xe5, 0xf8, 0xfc, 0xb4, 0x35, 0xb5, 0xac, 0xda, 0x69, 0x81, 0xa7,
	0x43, 0x9f, 0xd0, 0x83, 0x77, 0x12, 0x99, 0x38, 0xb6, 0x8e, 0xd7, 0xd1, 0xd6, 0x82, 0xc5, 0x3d,
	0xfe, 0xbb, 0x0e, 0xb6, 0x99, 0x08, 0x32, 0xfd, 0x40, 0x16, 0x4a, 0x7a, 0x19, 0x2f, 0x70, 0xdd,
	0x98, 0x1f, 0xfd, 0x2c, 0xfe, 0xd1, 0x98, 0x99, 0xfe, 0x18, 0x70, 0xf4, 0x33, 0x81, 0xb1, 0xec,
	0x1c, 0xef, 0xf2, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x59, 0xb7, 0x57, 0x6f, 0x02,
	0x00, 0x00,
}