// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_alert_rule.proto

package alert_rule

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/monitor"
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
//CreateAlertRule返回
type CreateAlertRuleResponse struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//data
	Data                 *CreateAlertRuleResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CreateAlertRuleResponse) Reset()         { *m = CreateAlertRuleResponse{} }
func (m *CreateAlertRuleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAlertRuleResponse) ProtoMessage()    {}
func (*CreateAlertRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ee26232dd15311d, []int{0}
}
func (m *CreateAlertRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAlertRuleResponse.Unmarshal(m, b)
}
func (m *CreateAlertRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAlertRuleResponse.Marshal(b, m, deterministic)
}
func (m *CreateAlertRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAlertRuleResponse.Merge(m, src)
}
func (m *CreateAlertRuleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAlertRuleResponse.Size(m)
}
func (m *CreateAlertRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAlertRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAlertRuleResponse proto.InternalMessageInfo

func (m *CreateAlertRuleResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateAlertRuleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CreateAlertRuleResponse) GetData() *CreateAlertRuleResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateAlertRuleResponse_Data struct {
	//
	//告警规则id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAlertRuleResponse_Data) Reset()         { *m = CreateAlertRuleResponse_Data{} }
func (m *CreateAlertRuleResponse_Data) String() string { return proto.CompactTextString(m) }
func (*CreateAlertRuleResponse_Data) ProtoMessage()    {}
func (*CreateAlertRuleResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ee26232dd15311d, []int{0, 0}
}
func (m *CreateAlertRuleResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAlertRuleResponse_Data.Unmarshal(m, b)
}
func (m *CreateAlertRuleResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAlertRuleResponse_Data.Marshal(b, m, deterministic)
}
func (m *CreateAlertRuleResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAlertRuleResponse_Data.Merge(m, src)
}
func (m *CreateAlertRuleResponse_Data) XXX_Size() int {
	return xxx_messageInfo_CreateAlertRuleResponse_Data.Size(m)
}
func (m *CreateAlertRuleResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAlertRuleResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAlertRuleResponse_Data proto.InternalMessageInfo

func (m *CreateAlertRuleResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//CreateAlertRuleApi返回
type CreateAlertRuleResponseWrapper struct {
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
	Data                 *CreateAlertRuleResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateAlertRuleResponseWrapper) Reset()         { *m = CreateAlertRuleResponseWrapper{} }
func (m *CreateAlertRuleResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateAlertRuleResponseWrapper) ProtoMessage()    {}
func (*CreateAlertRuleResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ee26232dd15311d, []int{1}
}
func (m *CreateAlertRuleResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAlertRuleResponseWrapper.Unmarshal(m, b)
}
func (m *CreateAlertRuleResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAlertRuleResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateAlertRuleResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAlertRuleResponseWrapper.Merge(m, src)
}
func (m *CreateAlertRuleResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateAlertRuleResponseWrapper.Size(m)
}
func (m *CreateAlertRuleResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAlertRuleResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAlertRuleResponseWrapper proto.InternalMessageInfo

func (m *CreateAlertRuleResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateAlertRuleResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateAlertRuleResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateAlertRuleResponseWrapper) GetData() *CreateAlertRuleResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAlertRuleResponse)(nil), "alert_rule.CreateAlertRuleResponse")
	proto.RegisterType((*CreateAlertRuleResponse_Data)(nil), "alert_rule.CreateAlertRuleResponse.Data")
	proto.RegisterType((*CreateAlertRuleResponseWrapper)(nil), "alert_rule.CreateAlertRuleResponseWrapper")
}

func init() { proto.RegisterFile("create_alert_rule.proto", fileDescriptor_3ee26232dd15311d) }

var fileDescriptor_3ee26232dd15311d = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0xa6, 0xb3, 0xb3, 0xc2, 0x64, 0xc4, 0x1f, 0x15, 0xdc, 0x61, 0x10, 0x3b, 0x64, 0x17, 0x99,
	0x4b, 0x1b, 0x59, 0x17, 0x51, 0xf1, 0xb2, 0xe3, 0x0f, 0xbc, 0x78, 0xc9, 0x45, 0x41, 0x74, 0xc9,
	0x24, 0x99, 0x18, 0x4c, 0xfb, 0x4a, 0x92, 0xba, 0x8a, 0xf8, 0xaf, 0x56, 0xf0, 0xee, 0xc1, 0xfe,
	0x05, 0xd2, 0xa4, 0x38, 0x73, 0x19, 0xd8, 0x53, 0xf3, 0xde, 0xf7, 0x7d, 0xef, 0x7b, 0xef, 0x2b,
	0x3a, 0xe2, 0x56, 0x32, 0x2f, 0x2f, 0x98, 0x91, 0xd6, 0x5f, 0xd8, 0xc6, 0xc8, 0xa2, 0xb6, 0xe0,
	0x21, 0x45, 0xdb, 0xce, 0x3c, 0x57, 0xda, 0x7f, 0x6e, 0xd6, 0x05, 0x87, 0x92, 0x28, 0x50, 0x40,
	0x02, 0x65, 0xdd, 0x6c, 0x42, 0x15, 0x8a, 0xf0, 0x8a, 0xd2, 0xf9, 0x7b, 0x05, 0x85, 0x64, 0xee,
	0x3b, 0xd4, 0xae, 0x30, 0xc0, 0x99, 0x21, 0x1c, 0x2a, 0x6f, 0x19, 0xf7, 0x2e, 0x2a, 0xad, 0xac,
	0x21, 0x2f, 0x41, 0x48, 0xe3, 0xc8, 0x40, 0x24, 0xa1, 0x24, 0x25, 0x54, 0xda, 0x83, 0x25, 0xd1,
	0x9c, 0x43, 0x25, 0xb4, 0xd7, 0x50, 0xb9, 0x61, 0xf2, 0xe3, 0x9d, 0x45, 0xca, 0x4b, 0xed, 0xbf,
	0xc0, 0x25, 0x51, 0x90, 0x07, 0x30, 0xff, 0xca, 0x8c, 0x16, 0xcc, 0x83, 0x75, 0xe4, 0xff, 0x73,
	0xd0, 0xdd, 0x53, 0x00, 0xca, 0xc8, 0xed, 0xde, 0xce, 0xdb, 0x86, 0xfb, 0x88, 0xe2, 0x3f, 0x09,
	0x3a, 0x7a, 0x11, 0x62, 0x38, 0xef, 0x6d, 0x69, 0x63, 0x24, 0x95, 0xae, 0x86, 0xca, 0xc9, 0xf4,
	0x18, 0x8d, 0x39, 0x08, 0x39, 0x4b, 0x16, 0xc9, 0xf2, 0x70, 0x75, 0xb3, 0x6b, 0xb3, 0xe9, 0x06,
	0x6c, 0xf9, 0x0c, 0xf7, 0x5d, 0x4c, 0x03, 0x98, 0x2e, 0xd0, 0x41, 0xe9, 0xd4, 0x6c, 0xb4, 0x48,
	0x96, 0x93, 0xd5, 0x8d, 0xae, 0xcd, 0x50, 0xe4, 0x94, 0x4e, 0x61, 0xda, 0x43, 0xe9, 0x5b, 0x34,
	0x16, 0xcc, 0xb3, 0xd9, 0xc1, 0x22, 0x59, 0x4e, 0x4f, 0x97, 0xc5, 0x4e, 0xdc, 0x7b, 0x9c, 0x8b,
	0x97, 0xcc, 0xb3, 0x5d, 0xc3, 0x5e, 0x8f, 0x69, 0x18, 0x33, 0x7f, 0x8e, 0xc6, 0x3d, 0x9c, 0x9e,
	0xa1, 0x91, 0x16, 0x61, 0xb7, 0xc9, 0xea, 0xa4, 0x6b, 0xb3, 0x49, 0xa4, 0x6a, 0x81, 0x7f, 0xff,
	0xca, 0xee, 0xa0, 0xdb, 0x9f, 0x3e, 0x3c, 0xcc, 0x9f, 0xb2, 0x7c, 0x73, 0x9e, 0xbf, 0xfe, 0xf8,
	0xe3, 0xf4, 0xec, 0xe7, 0x09, 0x1d, 0x69, 0x81, 0xff, 0x26, 0xe8, 0xfe, 0x1e, 0xd7, 0x77, 0x96,
	0xd5, 0xb5, 0xb4, 0x57, 0x3b, 0xfb, 0x09, 0x9a, 0xf6, 0xdf, 0x57, 0xdf, 0x6a, 0xc3, 0x74, 0x35,
	0x9c, 0x7f, 0xb7, 0x6b, 0xb3, 0x74, 0xcb, 0x1d, 0x40, 0x4c, 0x77, 0xa9, 0xe9, 0x03, 0x74, 0x28,
	0xad, 0x05, 0x1b, 0xf2, 0x98, 0xac, 0x6e, 0x75, 0x6d, 0x76, 0x3d, 0x6a, 0x42, 0x1b, 0xd3, 0x08,
	0xa7, 0x6f, 0x86, 0xd8, 0xc6, 0x21, 0xb6, 0xe3, 0x2b, 0xc4, 0xb6, 0x27, 0xb1, 0xf5, 0xb5, 0xf0,
	0xab, 0x1f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x84, 0x8f, 0x7d, 0x9a, 0xf0, 0x02, 0x00, 0x00,
}
