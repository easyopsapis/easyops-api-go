// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_metrics_states.proto

package stream

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	metadata_center "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/metadata_center"
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
//ListMetricsStates返回
type ListMetricsStatesResponse struct {
	//
	//指标表列表
	States               []*metadata_center.StreamMetricsSchema `protobuf:"bytes,1,rep,name=states,proto3" json:"states" form:"states"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ListMetricsStatesResponse) Reset()         { *m = ListMetricsStatesResponse{} }
func (m *ListMetricsStatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMetricsStatesResponse) ProtoMessage()    {}
func (*ListMetricsStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ac00bfafecc82f, []int{0}
}
func (m *ListMetricsStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMetricsStatesResponse.Unmarshal(m, b)
}
func (m *ListMetricsStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMetricsStatesResponse.Marshal(b, m, deterministic)
}
func (m *ListMetricsStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMetricsStatesResponse.Merge(m, src)
}
func (m *ListMetricsStatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMetricsStatesResponse.Size(m)
}
func (m *ListMetricsStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMetricsStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMetricsStatesResponse proto.InternalMessageInfo

func (m *ListMetricsStatesResponse) GetStates() []*metadata_center.StreamMetricsSchema {
	if m != nil {
		return m.States
	}
	return nil
}

//
//ListMetricsStatesApi返回
type ListMetricsStatesResponseWrapper struct {
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
	Data                 *ListMetricsStatesResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMetricsStatesResponseWrapper) Reset()         { *m = ListMetricsStatesResponseWrapper{} }
func (m *ListMetricsStatesResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListMetricsStatesResponseWrapper) ProtoMessage()    {}
func (*ListMetricsStatesResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ac00bfafecc82f, []int{1}
}
func (m *ListMetricsStatesResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMetricsStatesResponseWrapper.Unmarshal(m, b)
}
func (m *ListMetricsStatesResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMetricsStatesResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListMetricsStatesResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMetricsStatesResponseWrapper.Merge(m, src)
}
func (m *ListMetricsStatesResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListMetricsStatesResponseWrapper.Size(m)
}
func (m *ListMetricsStatesResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMetricsStatesResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListMetricsStatesResponseWrapper proto.InternalMessageInfo

func (m *ListMetricsStatesResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListMetricsStatesResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListMetricsStatesResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListMetricsStatesResponseWrapper) GetData() *ListMetricsStatesResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMetricsStatesResponse)(nil), "stream.ListMetricsStatesResponse")
	proto.RegisterType((*ListMetricsStatesResponseWrapper)(nil), "stream.ListMetricsStatesResponseWrapper")
}

func init() { proto.RegisterFile("list_metrics_states.proto", fileDescriptor_73ac00bfafecc82f) }

var fileDescriptor_73ac00bfafecc82f = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xf1, 0xe6, 0x0f, 0xac, 0xbc, 0xcb, 0x6e, 0x7d, 0x28, 0x4e, 0x2e, 0x76, 0xd5, 0x52,
	0x7c, 0x89, 0x0c, 0xe9, 0xa5, 0xf4, 0x18, 0x68, 0x4f, 0x2d, 0x05, 0xe7, 0xd0, 0x63, 0x50, 0x94,
	0x89, 0x63, 0xb0, 0x3c, 0x42, 0x52, 0xa0, 0x7d, 0x59, 0xbf, 0x43, 0xfd, 0x04, 0xc5, 0x52, 0x4a,
	0x42, 0x21, 0x27, 0x7b, 0xf4, 0xcd, 0xef, 0x9b, 0x6f, 0x86, 0x4c, 0xea, 0xca, 0xd8, 0x95, 0x04,
	0xab, 0x2b, 0x61, 0x56, 0xc6, 0x72, 0x0b, 0x86, 0x29, 0x8d, 0x16, 0xa3, 0xb1, 0xb1, 0x1a, 0xb8,
	0x9c, 0xce, 0xca, 0xca, 0xee, 0xf6, 0x6b, 0x26, 0x50, 0xe6, 0x25, 0x96, 0x98, 0x3b, 0x79, 0xbd,
	0xdf, 0xba, 0xca, 0x15, 0xee, 0xcf, 0x63, 0x53, 0x28, 0x91, 0x01, 0x37, 0x1f, 0xa8, 0x0c, 0xab,
	0x51, 0xf0, 0x3a, 0x17, 0xd8, 0x58, 0xcd, 0x85, 0x35, 0x9e, 0xd4, 0xa0, 0x70, 0x26, 0x71, 0x03,
	0xb5, 0xc9, 0x0f, 0x8d, 0xb9, 0x2b, 0x73, 0x09, 0x96, 0x6f, 0xb8, 0xe5, 0x2b, 0x01, 0x8d, 0x05,
	0x9d, 0xfb, 0x00, 0xc7, 0x74, 0x62, 0x07, 0x92, 0xfb, 0x31, 0xb4, 0x26, 0x93, 0xe7, 0xca, 0xd8,
	0x17, 0xaf, 0x2d, 0x5d, 0xf0, 0x02, 0x8c, 0xc2, 0xc6, 0x40, 0xf4, 0x4a, 0xc6, 0x7e, 0x95, 0x38,
	0x48, 0x07, 0x59, 0x38, 0xbf, 0x61, 0x3f, 0xac, 0xd9, 0xd2, 0x59, 0x7f, 0xd3, 0xce, 0x78, 0x71,
	0xd1, 0xb5, 0xc9, 0xdf, 0x2d, 0x6a, 0xf9, 0x40, 0x3d, 0x4d, 0x8b, 0x83, 0x0d, 0xfd, 0x0c, 0x48,
	0x7a, 0x76, 0xdc, 0x9b, 0xe6, 0x4a, 0x81, 0x8e, 0xae, 0xc9, 0x50, 0xe0, 0x06, 0xe2, 0x20, 0x0d,
	0xb2, 0xd1, 0xe2, 0x5f, 0xd7, 0x26, 0xa1, 0x77, 0xeb, 0x5f, 0x69, 0xe1, 0xc4, 0xe8, 0x9e, 0x84,
	0xfd, 0xf7, 0xf1, 0x5d, 0xd5, 0xbc, 0x6a, 0xe2, 0x5f, 0x69, 0x90, 0xfd, 0x5e, 0x5c, 0x76, 0x6d,
	0x12, 0x1d, 0x7b, 0x0f, 0x22, 0x2d, 0x4e, 0x5b, 0xa3, 0x5b, 0x32, 0x02, 0xad, 0x51, 0xc7, 0x03,
	0xc7, 0xfc, 0xef, 0xda, 0xe4, 0x8f, 0x67, 0xdc, 0x33, 0x2d, 0xbc, 0x1c, 0x3d, 0x91, 0x61, 0xbf,
	0x69, 0x3c, 0x4c, 0x83, 0x2c, 0x9c, 0x5f, 0x31, 0x7f, 0x45, 0x76, 0x36, 0xfe, 0x69, 0xd2, 0x1e,
	0xa4, 0x85, 0xe3, 0xd7, 0x63, 0x77, 0xe8, 0xbb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x00,
	0x4d, 0x81, 0x23, 0x02, 0x00, 0x00,
}
