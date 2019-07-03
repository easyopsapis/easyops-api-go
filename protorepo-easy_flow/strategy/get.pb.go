// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package strategy

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
	easy_flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_flow"
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
//Get请求
type GetRequest struct {
	//
	//策略Id
	StrategyID           string   `protobuf:"bytes,1,opt,name=strategyID,proto3" json:"strategyID" form:"strategyID"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetStrategyID() string {
	if m != nil {
		return m.StrategyID
	}
	return ""
}

//
//GetApi返回
type GetResponseWrapper struct {
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
	Data                 *easy_flow.DeployStrategy `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetResponseWrapper) Reset()         { *m = GetResponseWrapper{} }
func (m *GetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetResponseWrapper) ProtoMessage()    {}
func (*GetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseWrapper.Unmarshal(m, b)
}
func (m *GetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseWrapper.Merge(m, src)
}
func (m *GetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetResponseWrapper.Size(m)
}
func (m *GetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseWrapper proto.InternalMessageInfo

func (m *GetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetResponseWrapper) GetData() *easy_flow.DeployStrategy {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "strategy.GetRequest")
	proto.RegisterType((*GetResponseWrapper)(nil), "strategy.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x86, 0xa9, 0xce, 0x8a, 0x9b, 0x11, 0xd4, 0x5c, 0xc8, 0x38, 0x37, 0x5d, 0xa2, 0xc8, 0xde,
	0xa4, 0xd1, 0x55, 0x64, 0xf5, 0x42, 0x71, 0x59, 0x11, 0x2f, 0xad, 0x8a, 0xa8, 0xac, 0x43, 0x9a,
	0x66, 0xba, 0xc5, 0x74, 0x4e, 0x4c, 0x32, 0x8e, 0xa3, 0xec, 0x23, 0xf8, 0x06, 0xbe, 0x99, 0x50,
	0xc1, 0x47, 0xe8, 0x13, 0xc8, 0x9c, 0x76, 0x66, 0xeb, 0xa5, 0xb2, 0x57, 0xcd, 0x9f, 0xff, 0xfc,
	0x87, 0x2f, 0xfd, 0xc9, 0x76, 0xa1, 0x43, 0x62, 0x1d, 0x04, 0xa0, 0x17, 0x7d, 0x70, 0x32, 0xe8,
	0x62, 0x39, 0xe6, 0x45, 0x19, 0x8e, 0xe7, 0x59, 0xa2, 0xa0, 0x12, 0x05, 0x14, 0x20, 0x70, 0x20,
	0x9b, 0x4f, 0x51, 0xa1, 0xc0, 0x53, 0x1b, 0x1c, 0x1f, 0x15, 0x90, 0x68, 0xe9, 0x97, 0x60, 0x7d,
	0x62, 0x40, 0x49, 0x23, 0x14, 0xcc, 0x82, 0x93, 0x2a, 0xf8, 0x36, 0xe9, 0xb4, 0x05, 0x5e, 0x41,
	0xae, 0x8d, 0x17, 0xdd, 0xa0, 0x40, 0x89, 0x6a, 0x32, 0x35, 0xb0, 0x10, 0xb9, 0xb6, 0x06, 0x96,
	0x93, 0x4c, 0x06, 0x75, 0x3c, 0x09, 0xd2, 0x6d, 0xb8, 0xc6, 0x2f, 0xfe, 0x7f, 0xbd, 0xaa, 0xf2,
	0x4c, 0x28, 0x33, 0xf7, 0x41, 0xbb, 0x49, 0x39, 0x9b, 0xae, 0x89, 0x5f, 0x9f, 0x05, 0x71, 0x0b,
	0xd9, 0x5f, 0xfb, 0xf6, 0x0c, 0x7f, 0xc4, 0xba, 0x8c, 0x6e, 0xf5, 0xfd, 0x5e, 0x25, 0xd5, 0xa2,
	0x0c, 0x1f, 0x61, 0x21, 0x0a, 0xe0, 0x68, 0xf2, 0xcf, 0xd2, 0x94, 0xb9, 0x0c, 0xe0, 0xbc, 0xd8,
	0x1c, 0xdb, 0x1c, 0xfb, 0x11, 0x11, 0xf2, 0x4c, 0x87, 0x54, 0x7f, 0x9a, 0x6b, 0x1f, 0xe8, 0xf7,
	0x88, 0x90, 0xf5, 0xe6, 0xe7, 0x87, 0xa3, 0x68, 0x27, 0xda, 0xdd, 0x3e, 0xa8, 0x9a, 0x3a, 0xbe,
	0x3a, 0x05, 0x57, 0x3d, 0x64, 0xa7, 0x1e, 0xfb, 0xfd, 0x2b, 0x7e, 0x45, 0xd2, 0x0f, 0xef, 0x1f,
	0x4b, 0xfe, 0xf5, 0x09, 0x7f, 0x77, 0x9b, 0x3f, 0x38, 0xfa, 0xb6, 0x7f, 0xc2, 0xff, 0xd2, 0xf7,
	0xfe, 0x51, 0xdf, 0xd9, 0x3b, 0xb9, 0x99, 0xf6, 0x00, 0xd8, 0xcf, 0x88, 0x50, 0xc4, 0xf3, 0x16,
	0x66, 0x5e, 0xbf, 0x71, 0xd2, 0x5a, 0xed, 0xe8, 0x0d, 0x32, 0x50, 0x90, 0x6b, 0xe4, 0xdb, 0x3a,
	0xb8, 0xdc, 0xd4, 0xf1, 0xb0, 0xe5, 0x5b, 0xdd, 0xb2, 0x14, 0x4d, 0xba, 0x4f, 0x86, 0xab, 0xef,
	0xd3, 0x2f, 0xd6, 0xc8, 0x72, 0x36, 0x3a, 0x87, 0x6f, 0xb9, 0xd6, 0xd4, 0x31, 0x3d, 0x9d, 0xed,
	0x4c, 0x96, 0xf6, 0x47, 0xe9, 0x2d, 0xb2, 0xa5, 0x9d, 0x03, 0x37, 0x3a, 0x8f, 0x99, 0x2b, 0x4d,
	0x1d, 0x5f, 0x6a, 0x33, 0x78, 0xcd, 0xd2, 0xd6, 0xa6, 0x8f, 0xc8, 0x20, 0x97, 0x41, 0x8e, 0x06,
	0x3b, 0xd1, 0xee, 0x70, 0xef, 0x7a, 0xb2, 0x29, 0x29, 0x39, 0xc4, 0x92, 0x5e, 0x76, 0x0f, 0xe9,
	0x13, 0xae, 0x02, 0x2c, 0xc5, 0x5c, 0x76, 0x01, 0x3b, 0xb8, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x56, 0x5a, 0xc4, 0x65, 0x03, 0x00, 0x00,
}
