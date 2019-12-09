// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_not_recover_alert_event_list.proto

package alert

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
//GetNotRecoverAlertEventList请求
type GetNotRecoverAlertEventListRequest struct {
	//
	//开始时间, e.g.: -1h 过去1小时, -1d 过去1天, -30m 过去30分钟, 1490967693 unix 时间戳
	St string `protobuf:"bytes,1,opt,name=st,proto3" json:"st" form:"st"`
	//
	//结束时间, e.g.: -1h 过去1小时, -1d 过去1天, -30m 过去30分钟, 1490967693 unix 时间戳。默认为当前时间
	Et string `protobuf:"bytes,2,opt,name=et,proto3" json:"et" form:"et"`
	//
	//搜索关键字
	KeyWord string `protobuf:"bytes,3,opt,name=key_word,json=keyWord,proto3" json:"key_word" form:"key_word"`
	//
	//告警类型
	AlertType string `protobuf:"bytes,4,opt,name=alert_type,json=alertType,proto3" json:"alert_type" form:"alert_type"`
	//
	//告警子类型
	AlertSubType string `protobuf:"bytes,5,opt,name=alert_sub_type,json=alertSubType,proto3" json:"alert_sub_type" form:"alert_sub_type"`
	//
	//告警项
	AlertItem string `protobuf:"bytes,6,opt,name=alert_item,json=alertItem,proto3" json:"alert_item" form:"alert_item"`
	//
	//模型id
	ObjectId string `protobuf:"bytes,7,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//实例id
	InstanceId string `protobuf:"bytes,8,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//页码
	Page int32 `protobuf:"varint,9,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小, [1, 300]
	PageSize int32 `protobuf:"varint,10,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//排序
	XSortby__            string   `protobuf:"bytes,11,opt,name=__sortby__,json=Sortby,proto3" json:"__sortby__" form:"__sortby__"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotRecoverAlertEventListRequest) Reset()         { *m = GetNotRecoverAlertEventListRequest{} }
func (m *GetNotRecoverAlertEventListRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotRecoverAlertEventListRequest) ProtoMessage()    {}
func (*GetNotRecoverAlertEventListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_974076026267aaab, []int{0}
}
func (m *GetNotRecoverAlertEventListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotRecoverAlertEventListRequest.Unmarshal(m, b)
}
func (m *GetNotRecoverAlertEventListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotRecoverAlertEventListRequest.Marshal(b, m, deterministic)
}
func (m *GetNotRecoverAlertEventListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotRecoverAlertEventListRequest.Merge(m, src)
}
func (m *GetNotRecoverAlertEventListRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotRecoverAlertEventListRequest.Size(m)
}
func (m *GetNotRecoverAlertEventListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotRecoverAlertEventListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotRecoverAlertEventListRequest proto.InternalMessageInfo

func (m *GetNotRecoverAlertEventListRequest) GetSt() string {
	if m != nil {
		return m.St
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetEt() string {
	if m != nil {
		return m.Et
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetKeyWord() string {
	if m != nil {
		return m.KeyWord
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetAlertType() string {
	if m != nil {
		return m.AlertType
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetAlertSubType() string {
	if m != nil {
		return m.AlertSubType
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetAlertItem() string {
	if m != nil {
		return m.AlertItem
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetNotRecoverAlertEventListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetNotRecoverAlertEventListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetNotRecoverAlertEventListRequest) GetXSortby__() string {
	if m != nil {
		return m.XSortby__
	}
	return ""
}

//
//GetNotRecoverAlertEventList返回
type GetNotRecoverAlertEventListResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//页码
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,5,opt,name=total,proto3" json:"total" form:"total"`
	//
	//告警事件列表
	Data                 []*monitor.AlertEvent `protobuf:"bytes,6,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetNotRecoverAlertEventListResponse) Reset()         { *m = GetNotRecoverAlertEventListResponse{} }
func (m *GetNotRecoverAlertEventListResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotRecoverAlertEventListResponse) ProtoMessage()    {}
func (*GetNotRecoverAlertEventListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_974076026267aaab, []int{1}
}
func (m *GetNotRecoverAlertEventListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponse.Unmarshal(m, b)
}
func (m *GetNotRecoverAlertEventListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponse.Marshal(b, m, deterministic)
}
func (m *GetNotRecoverAlertEventListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotRecoverAlertEventListResponse.Merge(m, src)
}
func (m *GetNotRecoverAlertEventListResponse) XXX_Size() int {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponse.Size(m)
}
func (m *GetNotRecoverAlertEventListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotRecoverAlertEventListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotRecoverAlertEventListResponse proto.InternalMessageInfo

func (m *GetNotRecoverAlertEventListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetNotRecoverAlertEventListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetNotRecoverAlertEventListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetNotRecoverAlertEventListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetNotRecoverAlertEventListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetNotRecoverAlertEventListResponse) GetData() []*monitor.AlertEvent {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//GetNotRecoverAlertEventListApi返回
type GetNotRecoverAlertEventListResponseWrapper struct {
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
	Data                 *GetNotRecoverAlertEventListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GetNotRecoverAlertEventListResponseWrapper) Reset() {
	*m = GetNotRecoverAlertEventListResponseWrapper{}
}
func (m *GetNotRecoverAlertEventListResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*GetNotRecoverAlertEventListResponseWrapper) ProtoMessage() {}
func (*GetNotRecoverAlertEventListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_974076026267aaab, []int{2}
}
func (m *GetNotRecoverAlertEventListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper.Unmarshal(m, b)
}
func (m *GetNotRecoverAlertEventListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetNotRecoverAlertEventListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper.Merge(m, src)
}
func (m *GetNotRecoverAlertEventListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper.Size(m)
}
func (m *GetNotRecoverAlertEventListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotRecoverAlertEventListResponseWrapper proto.InternalMessageInfo

func (m *GetNotRecoverAlertEventListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetNotRecoverAlertEventListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetNotRecoverAlertEventListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetNotRecoverAlertEventListResponseWrapper) GetData() *GetNotRecoverAlertEventListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetNotRecoverAlertEventListRequest)(nil), "alert.GetNotRecoverAlertEventListRequest")
	proto.RegisterType((*GetNotRecoverAlertEventListResponse)(nil), "alert.GetNotRecoverAlertEventListResponse")
	proto.RegisterType((*GetNotRecoverAlertEventListResponseWrapper)(nil), "alert.GetNotRecoverAlertEventListResponseWrapper")
}

func init() {
	proto.RegisterFile("get_not_recover_alert_event_list.proto", fileDescriptor_974076026267aaab)
}

var fileDescriptor_974076026267aaab = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xef, 0x4b, 0xdc, 0x48,
	0x18, 0xc7, 0xcf, 0xfd, 0xe5, 0xee, 0xac, 0xa7, 0x7b, 0x11, 0x8f, 0x9c, 0x20, 0x59, 0x46, 0x11,
	0x95, 0xdb, 0xc4, 0x1f, 0x87, 0xe7, 0xdd, 0x9b, 0x43, 0x41, 0x8e, 0x85, 0xd2, 0xc2, 0x6c, 0x41,
	0xa8, 0x68, 0xc8, 0x26, 0x8f, 0x69, 0x6a, 0x92, 0x49, 0x67, 0x66, 0xb5, 0xab, 0xf8, 0xae, 0x7f,
	0x51, 0xff, 0xa0, 0x14, 0xfa, 0x27, 0x04, 0xfa, 0xbe, 0x64, 0x26, 0xbb, 0x59, 0xa4, 0x2d, 0xfa,
	0x6a, 0x67, 0x9e, 0xcf, 0xf7, 0xfb, 0x64, 0x9e, 0x2f, 0x33, 0x8b, 0x36, 0x7d, 0x10, 0x76, 0x4c,
	0x85, 0xcd, 0xc0, 0xa5, 0x37, 0xc0, 0x6c, 0x27, 0x04, 0x26, 0x6c, 0xb8, 0x81, 0x58, 0xd8, 0x61,
	0xc0, 0x85, 0x99, 0x30, 0x2a, 0xa8, 0x56, 0x97, 0xf5, 0xd5, 0x9e, 0x1f, 0x88, 0xb7, 0xa3, 0xa1,
	0xe9, 0xd2, 0xc8, 0xf2, 0xa9, 0x4f, 0x2d, 0x49, 0x87, 0xa3, 0x2b, 0xb9, 0x93, 0x1b, 0xb9, 0x52,
	0xae, 0xd5, 0x81, 0x4f, 0x4d, 0x70, 0xf8, 0x98, 0x26, 0xdc, 0x0c, 0xa9, 0xeb, 0x84, 0x96, 0x4b,
	0x63, 0xc1, 0x1c, 0x57, 0x70, 0xe5, 0x64, 0x90, 0xd0, 0x5e, 0x44, 0x3d, 0x08, 0xb9, 0x55, 0x08,
	0x2d, 0xb9, 0xb5, 0x22, 0x1a, 0x07, 0x82, 0x32, 0x6b, 0xe6, 0x3c, 0x45, 0xd3, 0xc3, 0x99, 0x33,
	0x44, 0xb7, 0x81, 0xb8, 0xa6, 0xb7, 0x96, 0x4f, 0x7b, 0x12, 0xf6, 0x6e, 0x9c, 0x30, 0xf0, 0x1c,
	0x41, 0x19, 0xb7, 0xa6, 0x4b, 0xe5, 0xc3, 0x5f, 0x6b, 0x08, 0xff, 0x0f, 0xe2, 0x25, 0x15, 0x44,
	0xcd, 0x7a, 0x9c, 0xb7, 0x3e, 0xcd, 0x3b, 0xbf, 0x08, 0xb8, 0x20, 0xf0, 0x7e, 0x04, 0x5c, 0x68,
	0x6b, 0xa8, 0xc2, 0x85, 0x3e, 0xd7, 0x9d, 0xdb, 0x6a, 0x9d, 0xfc, 0x9a, 0xa5, 0x46, 0xeb, 0x8a,
	0xb2, 0xe8, 0x5f, 0xcc, 0x05, 0x26, 0x15, 0x85, 0x41, 0xe8, 0x95, 0xc7, 0x18, 0x72, 0x0c, 0x42,
	0x33, 0x51, 0xf3, 0x1a, 0xc6, 0xf6, 0x2d, 0x65, 0x9e, 0x5e, 0x95, 0xa2, 0xe5, 0x2c, 0x35, 0x96,
	0x94, 0x68, 0x42, 0x30, 0x99, 0xbf, 0x86, 0xf1, 0x19, 0x65, 0x9e, 0xf6, 0x17, 0x42, 0x6a, 0x42,
	0x31, 0x4e, 0x40, 0xaf, 0x49, 0xc7, 0x4a, 0x96, 0x1a, 0xbf, 0x29, 0x47, 0xc9, 0x30, 0x69, 0xc9,
	0xcd, 0xeb, 0x71, 0x02, 0xda, 0x7f, 0x68, 0x51, 0x11, 0x3e, 0x1a, 0x2a, 0x67, 0x5d, 0x3a, 0xff,
	0xc8, 0x52, 0x63, 0x65, 0xd6, 0x39, 0xe1, 0x98, 0x2c, 0xc8, 0xc2, 0x60, 0x34, 0x94, 0x0d, 0xa6,
	0x9f, 0x0d, 0x04, 0x44, 0x7a, 0xe3, 0xfb, 0x9f, 0xcd, 0xd9, 0xe4, 0xb3, 0x7d, 0x01, 0x91, 0x46,
	0x50, 0x93, 0x0e, 0xdf, 0x81, 0x2b, 0xfa, 0x9e, 0x3e, 0x2f, 0x3d, 0x87, 0xe5, 0x70, 0x13, 0x82,
	0xbf, 0x7c, 0x36, 0x0c, 0xb4, 0x76, 0x79, 0xee, 0xf4, 0xee, 0x8e, 0x7b, 0x6f, 0xec, 0x8b, 0xf3,
	0xdd, 0xde, 0x3f, 0x93, 0xf5, 0xfd, 0xee, 0x9f, 0x07, 0x7b, 0x0f, 0x1b, 0x64, 0xda, 0x47, 0xeb,
	0x23, 0x14, 0xc4, 0x5c, 0x38, 0xb1, 0x0b, 0x7d, 0x4f, 0x6f, 0xca, 0xae, 0xdb, 0xe5, 0x49, 0x4a,
	0x96, 0xf7, 0xed, 0xa0, 0xc5, 0xcb, 0xa2, 0xdd, 0xc5, 0xfd, 0xde, 0xc1, 0xc3, 0x06, 0x99, 0x31,
	0x6b, 0xdb, 0xa8, 0x96, 0x38, 0x3e, 0xe8, 0xad, 0xee, 0xdc, 0x56, 0x5d, 0x8e, 0xd3, 0x56, 0x4d,
	0xf2, 0x6a, 0x6e, 0xaf, 0x74, 0x7e, 0x21, 0x52, 0xa2, 0xfd, 0x8d, 0x5a, 0xf9, 0xaf, 0xcd, 0x83,
	0x3b, 0xd0, 0x91, 0xd4, 0xaf, 0x66, 0xa9, 0xd1, 0x29, 0xf5, 0x12, 0x4d, 0x4c, 0xcd, 0xbc, 0x32,
	0x08, 0xee, 0x40, 0xdb, 0x43, 0xc8, 0xb6, 0x39, 0x65, 0x62, 0x38, 0xb6, 0x6d, 0xbd, 0xfd, 0x38,
	0xb8, 0x92, 0x61, 0xd2, 0x18, 0xc8, 0x25, 0xfe, 0x54, 0x41, 0xeb, 0x3f, 0xbd, 0x77, 0x3c, 0xa1,
	0x31, 0x07, 0x6d, 0x1d, 0xd5, 0x5c, 0xea, 0x81, 0xbc, 0x7a, 0xf5, 0x93, 0xa5, 0xf2, 0xf8, 0x79,
	0x15, 0x13, 0x09, 0xb5, 0x2e, 0xaa, 0x46, 0xdc, 0x2f, 0xee, 0xdf, 0x62, 0x96, 0x1a, 0x48, 0x69,
	0x22, 0xee, 0x63, 0x92, 0xa3, 0x69, 0x0a, 0xd5, 0x67, 0xa6, 0x50, 0x7b, 0x46, 0x0a, 0x9b, 0xa8,
	0x2e, 0xa8, 0x70, 0x42, 0x79, 0xed, 0xea, 0x27, 0x9d, 0x2c, 0x35, 0x16, 0x94, 0x49, 0x96, 0x31,
	0x51, 0x58, 0x3b, 0x42, 0x35, 0xcf, 0x11, 0x8e, 0xde, 0xe8, 0x56, 0xb7, 0xda, 0xfb, 0xcb, 0x66,
	0xf1, 0xa8, 0xcd, 0x32, 0x81, 0xd9, 0x39, 0x73, 0x29, 0x26, 0xd2, 0x81, 0x3f, 0x56, 0xd0, 0xce,
	0x13, 0x42, 0x3b, 0x63, 0x4e, 0x92, 0x00, 0x7b, 0x5a, 0x76, 0x47, 0xa8, 0x9d, 0xff, 0x9e, 0x7e,
	0x48, 0x42, 0x27, 0x88, 0x8b, 0x0c, 0x7f, 0xcf, 0x52, 0x43, 0x2b, 0xb5, 0x05, 0xc4, 0x64, 0x56,
	0x9a, 0xcf, 0x0b, 0x8c, 0x51, 0x56, 0x3c, 0xe9, 0x99, 0x79, 0x65, 0x19, 0x13, 0x85, 0xb5, 0x57,
	0xc5, 0xbc, 0x79, 0x96, 0xed, 0xfd, 0x1d, 0x53, 0x3e, 0x1d, 0xf3, 0x09, 0x73, 0xfc, 0x20, 0x86,
	0x61, 0x43, 0xfe, 0x75, 0x1d, 0x7c, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xdd, 0x23, 0x69, 0xa7,
	0x05, 0x00, 0x00,
}
