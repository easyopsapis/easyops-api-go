// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_stage_status.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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
//CreateStageStatus请求
type CreateStageStatusRequest struct {
	//
	//构建任务id
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id" form:"build_id"`
	//
	//stage状态
	Status               *pipeline.StageStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateStageStatusRequest) Reset()         { *m = CreateStageStatusRequest{} }
func (m *CreateStageStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStageStatusRequest) ProtoMessage()    {}
func (*CreateStageStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_231ea8441a8ca949, []int{0}
}
func (m *CreateStageStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStageStatusRequest.Unmarshal(m, b)
}
func (m *CreateStageStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStageStatusRequest.Marshal(b, m, deterministic)
}
func (m *CreateStageStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStageStatusRequest.Merge(m, src)
}
func (m *CreateStageStatusRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStageStatusRequest.Size(m)
}
func (m *CreateStageStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStageStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStageStatusRequest proto.InternalMessageInfo

func (m *CreateStageStatusRequest) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *CreateStageStatusRequest) GetStatus() *pipeline.StageStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

//
//CreateStageStatusApi返回
type CreateStageStatusResponseWrapper struct {
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
	Data                 *pipeline.StageStatus `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateStageStatusResponseWrapper) Reset()         { *m = CreateStageStatusResponseWrapper{} }
func (m *CreateStageStatusResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateStageStatusResponseWrapper) ProtoMessage()    {}
func (*CreateStageStatusResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_231ea8441a8ca949, []int{1}
}
func (m *CreateStageStatusResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStageStatusResponseWrapper.Unmarshal(m, b)
}
func (m *CreateStageStatusResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStageStatusResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateStageStatusResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStageStatusResponseWrapper.Merge(m, src)
}
func (m *CreateStageStatusResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateStageStatusResponseWrapper.Size(m)
}
func (m *CreateStageStatusResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStageStatusResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStageStatusResponseWrapper proto.InternalMessageInfo

func (m *CreateStageStatusResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateStageStatusResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateStageStatusResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateStageStatusResponseWrapper) GetData() *pipeline.StageStatus {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateStageStatusRequest)(nil), "build.CreateStageStatusRequest")
	proto.RegisterType((*CreateStageStatusResponseWrapper)(nil), "build.CreateStageStatusResponseWrapper")
}

func init() { proto.RegisterFile("create_stage_status.proto", fileDescriptor_231ea8441a8ca949) }

var fileDescriptor_231ea8441a8ca949 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xdd, 0x8a, 0xd4, 0x30,
	0x14, 0xa6, 0x3a, 0xb3, 0xba, 0x19, 0x75, 0xd7, 0x80, 0x52, 0xe7, 0xa6, 0x43, 0x14, 0xdd, 0x9b,
	0x36, 0xea, 0x82, 0xe8, 0x5e, 0x49, 0xc5, 0x0b, 0x6f, 0xbb, 0x88, 0x17, 0xa2, 0x4b, 0xda, 0x66,
	0x6a, 0xb0, 0xed, 0x89, 0x49, 0xea, 0xf8, 0x83, 0xef, 0xe2, 0x93, 0x55, 0x10, 0x7c, 0x81, 0x3e,
	0x81, 0xf4, 0x64, 0x46, 0x0b, 0xca, 0xde, 0xb4, 0xe7, 0x9c, 0xef, 0xa7, 0xdf, 0x47, 0xc9, 0xad,
	0xc2, 0x48, 0xe1, 0xe4, 0x99, 0x75, 0xa2, 0xc2, 0xa7, 0xeb, 0x6c, 0xa2, 0x0d, 0x38, 0xa0, 0xf3,
	0xbc, 0x53, 0x75, 0xb9, 0x8c, 0x2b, 0xe5, 0xde, 0x75, 0x79, 0x52, 0x40, 0xc3, 0x2b, 0xa8, 0x80,
	0x23, 0x9a, 0x77, 0x6b, 0xdc, 0x70, 0xc1, 0xc9, 0xab, 0x96, 0x2f, 0x2b, 0x48, 0xa4, 0xb0, 0x9f,
	0x41, 0xdb, 0xa4, 0x86, 0x42, 0xd4, 0xbc, 0x80, 0xd6, 0x19, 0x51, 0x38, 0xeb, 0x95, 0x46, 0x6a,
	0x88, 0x1b, 0x28, 0x65, 0x6d, 0xf9, 0x96, 0xc8, 0x71, 0xe5, 0x5a, 0x69, 0x59, 0xab, 0x56, 0xf2,
	0x7f, 0xc3, 0x2c, 0x1f, 0x4d, 0x52, 0x34, 0x1b, 0xe5, 0xde, 0xc3, 0x86, 0x57, 0x10, 0x23, 0x18,
	0x7f, 0x14, 0xb5, 0x2a, 0x85, 0x03, 0x63, 0xf9, 0x9f, 0xd1, 0xeb, 0xd8, 0xf7, 0x80, 0x84, 0xcf,
	0xb0, 0xe2, 0xe9, 0x68, 0x7a, 0x8a, 0x9e, 0x99, 0xfc, 0xd0, 0x49, 0xeb, 0x68, 0x4a, 0x2e, 0x63,
	0xc7, 0x33, 0x55, 0x86, 0xc1, 0x2a, 0x38, 0xda, 0x4f, 0xef, 0x0d, 0x7d, 0x74, 0xb0, 0x06, 0xd3,
	0x9c, 0xb0, 0x1d, 0xc2, 0x7e, 0xfe, 0x88, 0x0e, 0xc9, 0xb5, 0xb7, 0xaf, 0xef, 0xc7, 0x4f, 0x44,
	0xfc, 0xe5, 0xcd, 0xd7, 0x07, 0xc7, 0xdf, 0xee, 0x64, 0x97, 0x10, 0x7e, 0x51, 0xd2, 0xa7, 0x64,
	0xcf, 0x07, 0x0d, 0x2f, 0xac, 0x82, 0xa3, 0xc5, 0xc3, 0x1b, 0xc9, 0xae, 0x46, 0x32, 0xf9, 0x62,
	0x7a, 0x7d, 0xe8, 0xa3, 0xab, 0xde, 0xd8, 0xd3, 0x59, 0xb6, 0xd5, 0xb1, 0x5f, 0x01, 0x59, 0xfd,
	0x27, 0xa2, 0xd5, 0xd0, 0x5a, 0xf9, 0xca, 0x08, 0xad, 0xa5, 0xa1, 0xb7, 0xc9, 0xac, 0x80, 0x52,
	0x62, 0xcc, 0x79, 0x7a, 0x30, 0xf4, 0xd1, 0xc2, 0xbb, 0x8d, 0x57, 0x96, 0x21, 0x48, 0x1f, 0x93,
	0xc5, 0xf8, 0x7e, 0xfe, 0x49, 0xd7, 0x42, 0xb5, 0x18, 0x68, 0x3f, 0xbd, 0x39, 0xf4, 0x11, 0xfd,
	0xcb, 0xdd, 0x82, 0x2c, 0x9b, 0x52, 0xe9, 0x5d, 0x32, 0x97, 0xc6, 0x80, 0x09, 0x2f, 0xa2, 0xe6,
	0x70, 0xe8, 0xa3, 0x2b, 0x5e, 0x83, 0x67, 0x96, 0x79, 0x98, 0x9e, 0x90, 0x59, 0x29, 0x9c, 0x08,
	0x67, 0xe7, 0x75, 0x9d, 0xa4, 0x1b, 0xc9, 0x2c, 0x43, 0x4d, 0xbe, 0x87, 0x7f, 0xe4, 0xf8, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x92, 0xab, 0x27, 0x73, 0x02, 0x00, 0x00,
}
