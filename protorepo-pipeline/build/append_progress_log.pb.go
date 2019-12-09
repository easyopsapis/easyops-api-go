// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: append_progress_log.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//AppendProgressLogApi返回
type AppendProgressLogResponseWrapper struct {
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
	Data                 *pipeline.ProgressLog `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AppendProgressLogResponseWrapper) Reset()         { *m = AppendProgressLogResponseWrapper{} }
func (m *AppendProgressLogResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AppendProgressLogResponseWrapper) ProtoMessage()    {}
func (*AppendProgressLogResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f74b379381186f, []int{0}
}
func (m *AppendProgressLogResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendProgressLogResponseWrapper.Unmarshal(m, b)
}
func (m *AppendProgressLogResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendProgressLogResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AppendProgressLogResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendProgressLogResponseWrapper.Merge(m, src)
}
func (m *AppendProgressLogResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AppendProgressLogResponseWrapper.Size(m)
}
func (m *AppendProgressLogResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendProgressLogResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AppendProgressLogResponseWrapper proto.InternalMessageInfo

func (m *AppendProgressLogResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AppendProgressLogResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AppendProgressLogResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AppendProgressLogResponseWrapper) GetData() *pipeline.ProgressLog {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AppendProgressLogResponseWrapper)(nil), "build.AppendProgressLogResponseWrapper")
}

func init() { proto.RegisterFile("append_progress_log.proto", fileDescriptor_a6f74b379381186f) }

var fileDescriptor_a6f74b379381186f = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xff, 0xcf, 0x08, 0xa6, 0x82, 0x12, 0x50, 0xea, 0x6c, 0x5a, 0x22, 0x48, 0x37,
	0x93, 0x80, 0x6e, 0x64, 0x76, 0x0e, 0xb8, 0x73, 0x21, 0x05, 0x71, 0x39, 0xa4, 0x6d, 0x26, 0x16,
	0xd2, 0xde, 0x90, 0xb4, 0xa0, 0x2f, 0xdb, 0x9d, 0x2f, 0xd0, 0x27, 0x90, 0xde, 0x56, 0x2c, 0xb8,
	0x4a, 0xce, 0x3d, 0xe7, 0x83, 0x73, 0xc8, 0xb5, 0xb4, 0x56, 0x35, 0xe5, 0xc1, 0x3a, 0xd0, 0x4e,
	0x79, 0x7f, 0x30, 0xa0, 0xb9, 0x75, 0xd0, 0x02, 0x5d, 0xe7, 0x5d, 0x65, 0xca, 0xcd, 0x56, 0x57,
	0xed, 0x7b, 0x97, 0xf3, 0x02, 0x6a, 0xa1, 0x41, 0x83, 0x40, 0x37, 0xef, 0x8e, 0xa8, 0x50, 0xe0,
	0x6f, 0xa2, 0x36, 0xaf, 0x1a, 0xb8, 0x92, 0xfe, 0x13, 0xac, 0xe7, 0x06, 0x0a, 0x69, 0x44, 0x01,
	0x4d, 0xeb, 0x64, 0xd1, 0xfa, 0x89, 0x74, 0xca, 0xc2, 0xb6, 0x86, 0x52, 0x19, 0x2f, 0xe6, 0xa0,
	0x40, 0x29, 0x6c, 0x65, 0x95, 0xa9, 0x1a, 0x25, 0xfe, 0x96, 0x61, 0x5f, 0x01, 0x49, 0x1e, 0xb1,
	0xea, 0xcb, 0x6c, 0x3e, 0x83, 0xce, 0x94, 0xb7, 0xd0, 0x78, 0xf5, 0xe6, 0xc6, 0x15, 0x8e, 0xde,
	0x90, 0x55, 0x01, 0xa5, 0x8a, 0x82, 0x24, 0x48, 0xd7, 0xfb, 0xf3, 0xa1, 0x8f, 0xc3, 0x23, 0xb8,
	0x7a, 0xc7, 0xc6, 0x2b, 0xcb, 0xd0, 0xa4, 0x0f, 0x24, 0x1c, 0xdf, 0xa7, 0x0f, 0x6b, 0x64, 0xd5,
	0x44, 0xff, 0x92, 0x20, 0x3d, 0xdd, 0x5f, 0x0d, 0x7d, 0x4c, 0x7f, 0xb3, 0xb3, 0xc9, 0xb2, 0x65,
	0x94, 0xde, 0x92, 0xb5, 0x72, 0x0e, 0x5c, 0xf4, 0x1f, 0x99, 0x8b, 0xa1, 0x8f, 0xcf, 0x26, 0x06,
	0xcf, 0x2c, 0x9b, 0x6c, 0xba, 0x23, 0xab, 0x52, 0xb6, 0x32, 0x5a, 0x25, 0x41, 0x1a, 0xde, 0x5d,
	0xf2, 0x9f, 0x5d, 0x7c, 0x51, 0x7d, 0xd9, 0x6e, 0x0c, 0xb3, 0x0c, 0x99, 0xfc, 0x04, 0xe7, 0xde,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x27, 0x5a, 0xd3, 0x44, 0x98, 0x01, 0x00, 0x00,
}
