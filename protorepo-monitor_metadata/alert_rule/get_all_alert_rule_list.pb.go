// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_all_alert_rule_list.proto

package alert_rule

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	alert_manager "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/alert_manager"
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
//GetAllAlertRuleList返回
type GetAllAlertRuleListResponse struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//返回码字符串表示
	CodeExplain string `protobuf:"bytes,2,opt,name=codeExplain,proto3" json:"codeExplain" form:"codeExplain"`
	//
	//返回码字符串表示
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//data
	Data                 []*alert_manager.AlertRule `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetAllAlertRuleListResponse) Reset()         { *m = GetAllAlertRuleListResponse{} }
func (m *GetAllAlertRuleListResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllAlertRuleListResponse) ProtoMessage()    {}
func (*GetAllAlertRuleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7a64a9aad9f01da, []int{0}
}
func (m *GetAllAlertRuleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllAlertRuleListResponse.Unmarshal(m, b)
}
func (m *GetAllAlertRuleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllAlertRuleListResponse.Marshal(b, m, deterministic)
}
func (m *GetAllAlertRuleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllAlertRuleListResponse.Merge(m, src)
}
func (m *GetAllAlertRuleListResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllAlertRuleListResponse.Size(m)
}
func (m *GetAllAlertRuleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllAlertRuleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllAlertRuleListResponse proto.InternalMessageInfo

func (m *GetAllAlertRuleListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAllAlertRuleListResponse) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAllAlertRuleListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetAllAlertRuleListResponse) GetData() []*alert_manager.AlertRule {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//GetAllAlertRuleListApi返回
type GetAllAlertRuleListResponseWrapper struct {
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
	Data                 *GetAllAlertRuleListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetAllAlertRuleListResponseWrapper) Reset()         { *m = GetAllAlertRuleListResponseWrapper{} }
func (m *GetAllAlertRuleListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetAllAlertRuleListResponseWrapper) ProtoMessage()    {}
func (*GetAllAlertRuleListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7a64a9aad9f01da, []int{1}
}
func (m *GetAllAlertRuleListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllAlertRuleListResponseWrapper.Unmarshal(m, b)
}
func (m *GetAllAlertRuleListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllAlertRuleListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetAllAlertRuleListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllAlertRuleListResponseWrapper.Merge(m, src)
}
func (m *GetAllAlertRuleListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetAllAlertRuleListResponseWrapper.Size(m)
}
func (m *GetAllAlertRuleListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllAlertRuleListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllAlertRuleListResponseWrapper proto.InternalMessageInfo

func (m *GetAllAlertRuleListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAllAlertRuleListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAllAlertRuleListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetAllAlertRuleListResponseWrapper) GetData() *GetAllAlertRuleListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllAlertRuleListResponse)(nil), "alert_rule.GetAllAlertRuleListResponse")
	proto.RegisterType((*GetAllAlertRuleListResponseWrapper)(nil), "alert_rule.GetAllAlertRuleListResponseWrapper")
}

func init() { proto.RegisterFile("get_all_alert_rule_list.proto", fileDescriptor_b7a64a9aad9f01da) }

var fileDescriptor_b7a64a9aad9f01da = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xb1, 0x4f, 0xe3, 0x30,
	0x14, 0xc6, 0x95, 0x6b, 0x7b, 0xa7, 0x73, 0x4e, 0x07, 0xf2, 0x80, 0xa2, 0x22, 0x94, 0xca, 0x48,
	0x90, 0x81, 0x26, 0x52, 0x59, 0x10, 0x12, 0x43, 0x2b, 0x21, 0x96, 0x4e, 0x59, 0x60, 0x8b, 0xdc,
	0xf4, 0x35, 0x44, 0x72, 0xf2, 0x22, 0xdb, 0x91, 0xe0, 0x9f, 0xcd, 0xce, 0x9a, 0x8d, 0x0d, 0xc5,
	0x0e, 0x6d, 0x19, 0xe8, 0xc6, 0x64, 0xbf, 0xf7, 0x7d, 0x9f, 0xfd, 0x7e, 0x7a, 0xe4, 0x2c, 0x03,
	0x9d, 0x70, 0x21, 0x12, 0x2e, 0x40, 0xea, 0x44, 0xd6, 0x02, 0x12, 0x91, 0x2b, 0x1d, 0x56, 0x12,
	0x35, 0x52, 0xb2, 0x6b, 0x8f, 0xa7, 0x59, 0xae, 0x9f, 0xeb, 0x55, 0x98, 0x62, 0x11, 0x65, 0x98,
	0x61, 0x64, 0x2c, 0xab, 0x7a, 0x63, 0x2a, 0x53, 0x98, 0x9b, 0x8d, 0x8e, 0x9f, 0x32, 0x0c, 0x81,
	0xab, 0x57, 0xac, 0x54, 0x28, 0x30, 0xe5, 0x22, 0x4a, 0xb1, 0xd4, 0x92, 0xa7, 0x5a, 0xd9, 0xa4,
	0x84, 0x0a, 0xa7, 0x05, 0xae, 0x41, 0xa8, 0xa8, 0x37, 0x46, 0xa6, 0x8c, 0xec, 0xa7, 0x05, 0x2f,
	0x79, 0x06, 0x32, 0xda, 0x8d, 0x60, 0x5f, 0x66, 0x6f, 0x0e, 0x39, 0x7d, 0x00, 0x3d, 0x17, 0x62,
	0xde, 0x49, 0x71, 0x2d, 0x60, 0x99, 0x2b, 0x1d, 0x83, 0xaa, 0xb0, 0x54, 0x40, 0xcf, 0xc9, 0x30,
	0xc5, 0x35, 0x78, 0xce, 0xc4, 0x09, 0x46, 0x8b, 0xa3, 0xb6, 0xf1, 0xdd, 0x0d, 0xca, 0xe2, 0x96,
	0x75, 0x5d, 0x16, 0x1b, 0x91, 0xde, 0x10, 0xb7, 0x3b, 0xef, 0x5f, 0x2a, 0xc1, 0xf3, 0xd2, 0xfb,
	0x35, 0x71, 0x82, 0xbf, 0x8b, 0x93, 0xb6, 0xf1, 0xe9, 0xce, 0xdb, 0x8b, 0x2c, 0xde, 0xb7, 0xd2,
	0x2b, 0xf2, 0xa7, 0x00, 0xa5, 0x78, 0x06, 0xde, 0xc0, 0xa4, 0x68, 0xdb, 0xf8, 0xff, 0x6d, 0xaa,
	0x17, 0x58, 0xfc, 0x69, 0xa1, 0x77, 0x64, 0xb8, 0xe6, 0x9a, 0x7b, 0xc3, 0xc9, 0x20, 0x70, 0x67,
	0x5e, 0xf8, 0x85, 0x2d, 0xdc, 0x02, 0xec, 0x8f, 0xd9, 0xf9, 0x59, 0x6c, 0x62, 0xec, 0xdd, 0x21,
	0xec, 0x00, 0xeb, 0xa3, 0xe4, 0x55, 0x05, 0xf2, 0xa7, 0x91, 0x2f, 0xc8, 0x08, 0xa4, 0x44, 0xd9,
	0x03, 0x1f, 0xb7, 0x8d, 0xff, 0xcf, 0x66, 0x4c, 0x9b, 0xc5, 0x56, 0xa6, 0xcb, 0x2d, 0xac, 0x13,
	0xb8, 0xb3, 0xcb, 0x70, 0x6f, 0x75, 0x07, 0x20, 0xbe, 0x61, 0x5f, 0xfd, 0x36, 0xeb, 0xbe, 0xfe,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xcf, 0x26, 0xbb, 0xa4, 0x02, 0x00, 0x00,
}