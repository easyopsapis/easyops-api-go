// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_alert_rule_list.proto

package alert_rule

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	monitor "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/monitor"
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
//GetAlertRuleList请求
type GetAlertRuleListRequest struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//规则id, 多个用逗号分隔符
	XId_In               string   `protobuf:"bytes,3,opt,name=_id__in,json=IdIn,proto3" json:"_id__in" form:"_id__in"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAlertRuleListRequest) Reset()         { *m = GetAlertRuleListRequest{} }
func (m *GetAlertRuleListRequest) String() string { return proto.CompactTextString(m) }
func (*GetAlertRuleListRequest) ProtoMessage()    {}
func (*GetAlertRuleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33199339d2ce1d3e, []int{0}
}
func (m *GetAlertRuleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertRuleListRequest.Unmarshal(m, b)
}
func (m *GetAlertRuleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertRuleListRequest.Marshal(b, m, deterministic)
}
func (m *GetAlertRuleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertRuleListRequest.Merge(m, src)
}
func (m *GetAlertRuleListRequest) XXX_Size() int {
	return xxx_messageInfo_GetAlertRuleListRequest.Size(m)
}
func (m *GetAlertRuleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertRuleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertRuleListRequest proto.InternalMessageInfo

func (m *GetAlertRuleListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAlertRuleListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetAlertRuleListRequest) GetXId_In() string {
	if m != nil {
		return m.XId_In
	}
	return ""
}

//
//GetAlertRuleList返回
type GetAlertRuleListResponse struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//页数
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//总数
	Total int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total" form:"total"`
	//
	//msg
	Msg string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//data
	Data                 []*monitor.AlertRule `protobuf:"bytes,6,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetAlertRuleListResponse) Reset()         { *m = GetAlertRuleListResponse{} }
func (m *GetAlertRuleListResponse) String() string { return proto.CompactTextString(m) }
func (*GetAlertRuleListResponse) ProtoMessage()    {}
func (*GetAlertRuleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33199339d2ce1d3e, []int{1}
}
func (m *GetAlertRuleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertRuleListResponse.Unmarshal(m, b)
}
func (m *GetAlertRuleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertRuleListResponse.Marshal(b, m, deterministic)
}
func (m *GetAlertRuleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertRuleListResponse.Merge(m, src)
}
func (m *GetAlertRuleListResponse) XXX_Size() int {
	return xxx_messageInfo_GetAlertRuleListResponse.Size(m)
}
func (m *GetAlertRuleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertRuleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertRuleListResponse proto.InternalMessageInfo

func (m *GetAlertRuleListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAlertRuleListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetAlertRuleListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAlertRuleListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetAlertRuleListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetAlertRuleListResponse) GetData() []*monitor.AlertRule {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//GetAlertRuleListApi返回
type GetAlertRuleListResponseWrapper struct {
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
	Data                 *GetAlertRuleListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetAlertRuleListResponseWrapper) Reset()         { *m = GetAlertRuleListResponseWrapper{} }
func (m *GetAlertRuleListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetAlertRuleListResponseWrapper) ProtoMessage()    {}
func (*GetAlertRuleListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_33199339d2ce1d3e, []int{2}
}
func (m *GetAlertRuleListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertRuleListResponseWrapper.Unmarshal(m, b)
}
func (m *GetAlertRuleListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertRuleListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetAlertRuleListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertRuleListResponseWrapper.Merge(m, src)
}
func (m *GetAlertRuleListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetAlertRuleListResponseWrapper.Size(m)
}
func (m *GetAlertRuleListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertRuleListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertRuleListResponseWrapper proto.InternalMessageInfo

func (m *GetAlertRuleListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAlertRuleListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetAlertRuleListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetAlertRuleListResponseWrapper) GetData() *GetAlertRuleListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAlertRuleListRequest)(nil), "alert_rule.GetAlertRuleListRequest")
	proto.RegisterType((*GetAlertRuleListResponse)(nil), "alert_rule.GetAlertRuleListResponse")
	proto.RegisterType((*GetAlertRuleListResponseWrapper)(nil), "alert_rule.GetAlertRuleListResponseWrapper")
}

func init() { proto.RegisterFile("get_alert_rule_list.proto", fileDescriptor_33199339d2ce1d3e) }

var fileDescriptor_33199339d2ce1d3e = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0x8a, 0xd4, 0x30,
	0x18, 0xb7, 0x3b, 0x9d, 0xd5, 0xc9, 0xc8, 0x3a, 0x04, 0xd4, 0x3a, 0x97, 0x0e, 0x51, 0x64, 0xf6,
	0x30, 0x2d, 0xac, 0xe0, 0x8a, 0x37, 0x07, 0x44, 0x06, 0x3c, 0xc5, 0x83, 0xc7, 0x92, 0x69, 0xb3,
	0x35, 0x98, 0xf6, 0xab, 0x49, 0xea, 0xea, 0x3e, 0x90, 0xf8, 0x54, 0x15, 0x7c, 0x84, 0xfa, 0x02,
	0x92, 0xb4, 0xbb, 0x2d, 0xc8, 0xa0, 0xe0, 0x29, 0xf9, 0x7e, 0x7f, 0x92, 0xef, 0xfb, 0x25, 0xe8,
	0x51, 0xce, 0x4d, 0xc2, 0x24, 0x57, 0x26, 0x51, 0xb5, 0xe4, 0x89, 0x14, 0xda, 0x44, 0x95, 0x02,
	0x03, 0x18, 0x0d, 0xf0, 0x72, 0x93, 0x0b, 0xf3, 0xa1, 0xde, 0x47, 0x29, 0x14, 0x71, 0x0e, 0x39,
	0xc4, 0x4e, 0xb2, 0xaf, 0x2f, 0x5c, 0xe5, 0x0a, 0xb7, 0xeb, 0xac, 0x4b, 0x9a, 0x43, 0xc4, 0x99,
	0xfe, 0x0a, 0x95, 0x8e, 0x24, 0xa4, 0x4c, 0xc6, 0x29, 0x94, 0x46, 0xb1, 0xd4, 0xe8, 0xce, 0xa9,
	0x78, 0x05, 0x9b, 0x02, 0x32, 0x2e, 0x75, 0xdc, 0x0b, 0x63, 0x57, 0xc6, 0x05, 0x94, 0xc2, 0x80,
	0x8a, 0x87, 0xcb, 0xfb, 0x33, 0x9f, 0x8f, 0x5a, 0x28, 0x2e, 0x85, 0xf9, 0x08, 0x97, 0x71, 0x0e,
	0x1b, 0x47, 0x6e, 0x3e, 0x33, 0x29, 0x32, 0x66, 0x40, 0xe9, 0xf8, 0x66, 0xdb, 0xf9, 0xc8, 0x77,
	0x0f, 0x3d, 0x7c, 0xc3, 0xcd, 0x2b, 0x7b, 0x1e, 0xad, 0x25, 0x7f, 0x2b, 0xb4, 0xa1, 0xfc, 0x53,
	0xcd, 0xb5, 0xc1, 0xa7, 0xc8, 0xaf, 0x58, 0xce, 0x03, 0x6f, 0xe5, 0xad, 0xa7, 0xdb, 0xfb, 0x6d,
	0x13, 0xce, 0x2f, 0x40, 0x15, 0x2f, 0x89, 0x45, 0xc9, 0xcf, 0x1f, 0xe1, 0xd1, 0xe2, 0x16, 0x75,
	0x12, 0x7c, 0x8e, 0x66, 0x76, 0x4d, 0xb4, 0xb8, 0xe2, 0xc1, 0x91, 0xd3, 0x2f, 0xdb, 0x26, 0x5c,
	0x0c, 0x7a, 0x47, 0x5d, 0x9b, 0xee, 0x58, 0xe4, 0x9d, 0xb8, 0xe2, 0xf8, 0x14, 0xdd, 0x4e, 0x44,
	0x96, 0x24, 0xa2, 0x0c, 0x26, 0x2b, 0x6f, 0x3d, 0xdb, 0xe2, 0xb6, 0x09, 0x4f, 0x3a, 0x5b, 0x4f,
	0x10, 0xea, 0xef, 0xb2, 0x5d, 0x49, 0xbe, 0x1d, 0xa1, 0xe0, 0xcf, 0x56, 0x75, 0x05, 0xa5, 0xe6,
	0xf8, 0x31, 0xf2, 0x53, 0xc8, 0xae, 0x7b, 0xbd, 0x37, 0xf4, 0x6a, 0x51, 0x42, 0x1d, 0xf9, 0x3f,
	0x5d, 0x76, 0x49, 0x4c, 0xfe, 0x9e, 0xc4, 0x53, 0x34, 0x35, 0x60, 0x98, 0x0c, 0x7c, 0xa7, 0x5d,
	0xb4, 0x4d, 0x78, 0xb7, 0xd3, 0x3a, 0x98, 0xd0, 0x8e, 0xc6, 0x2b, 0x34, 0x29, 0x74, 0x1e, 0x4c,
	0xdd, 0xd0, 0x27, 0x6d, 0x13, 0xa2, 0x4e, 0x55, 0xe8, 0x9c, 0x50, 0x4b, 0xe1, 0x73, 0xe4, 0x67,
	0xcc, 0xb0, 0xe0, 0x78, 0x35, 0x59, 0xcf, 0xcf, 0x70, 0xd4, 0xbf, 0x7d, 0x74, 0x13, 0xc0, 0x78,
	0x4c, 0xab, 0x24, 0xd4, 0x19, 0xc8, 0x2f, 0x0f, 0x85, 0x87, 0x82, 0x7a, 0xaf, 0x58, 0x55, 0x71,
	0xf5, 0x6f, 0x79, 0xbd, 0x40, 0x73, 0xbb, 0xbe, 0xfe, 0x52, 0x49, 0x26, 0x4a, 0x97, 0xd8, 0x6c,
	0xfb, 0xa0, 0x6d, 0x42, 0x3c, 0x68, 0x7b, 0x92, 0xd0, 0xb1, 0xd4, 0xa6, 0xc0, 0x95, 0x02, 0xd5,
	0x3f, 0xea, 0x28, 0x05, 0x07, 0x13, 0xda, 0xd1, 0x78, 0xd7, 0xcf, 0x68, 0xc3, 0x9a, 0x9f, 0x3d,
	0x89, 0x46, 0xff, 0xfa, 0xd0, 0x04, 0x07, 0xa6, 0xde, 0x1f, 0xbb, 0x0f, 0xfd, 0xec, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x14, 0x69, 0x1e, 0xc1, 0xb4, 0x03, 0x00, 0x00,
}