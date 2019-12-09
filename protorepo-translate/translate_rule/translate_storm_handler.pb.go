// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: translate_storm_handler.proto

package translate_rule

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//TranslateStormHandler请求
type TranslateStormHandlerRequest struct {
	//
	//数据表名称
	DataName             string   `protobuf:"bytes,1,opt,name=data_name,json=dataName,proto3" json:"data_name" form:"data_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateStormHandlerRequest) Reset()         { *m = TranslateStormHandlerRequest{} }
func (m *TranslateStormHandlerRequest) String() string { return proto.CompactTextString(m) }
func (*TranslateStormHandlerRequest) ProtoMessage()    {}
func (*TranslateStormHandlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{0}
}
func (m *TranslateStormHandlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerRequest.Unmarshal(m, b)
}
func (m *TranslateStormHandlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerRequest.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerRequest.Merge(m, src)
}
func (m *TranslateStormHandlerRequest) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerRequest.Size(m)
}
func (m *TranslateStormHandlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerRequest proto.InternalMessageInfo

func (m *TranslateStormHandlerRequest) GetDataName() string {
	if m != nil {
		return m.DataName
	}
	return ""
}

//
//TranslateStormHandler返回
type TranslateStormHandlerResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//分页参数
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//返回信息
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//总数
	Total int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total" form:"total"`
	//
	//页数
	Page int32 `protobuf:"varint,5,opt,name=page,proto3" json:"page" form:"page"`
	//
	//返回数据
	Data                 *TranslateStormHandlerResponse_Data `protobuf:"bytes,6,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *TranslateStormHandlerResponse) Reset()         { *m = TranslateStormHandlerResponse{} }
func (m *TranslateStormHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*TranslateStormHandlerResponse) ProtoMessage()    {}
func (*TranslateStormHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{1}
}
func (m *TranslateStormHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerResponse.Unmarshal(m, b)
}
func (m *TranslateStormHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerResponse.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerResponse.Merge(m, src)
}
func (m *TranslateStormHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerResponse.Size(m)
}
func (m *TranslateStormHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerResponse proto.InternalMessageInfo

func (m *TranslateStormHandlerResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TranslateStormHandlerResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *TranslateStormHandlerResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TranslateStormHandlerResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *TranslateStormHandlerResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *TranslateStormHandlerResponse) GetData() *TranslateStormHandlerResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type TranslateStormHandlerResponse_Data struct {
	//
	//匹配规则
	MatchRules *TranslateStormHandlerResponse_Data_MatchRules `protobuf:"bytes,1,opt,name=match_rules,json=matchRules,proto3" json:"match_rules" form:"match_rules"`
	//
	//规则id列表
	RuleIds              []string `protobuf:"bytes,2,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids" form:"rule_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateStormHandlerResponse_Data) Reset()         { *m = TranslateStormHandlerResponse_Data{} }
func (m *TranslateStormHandlerResponse_Data) String() string { return proto.CompactTextString(m) }
func (*TranslateStormHandlerResponse_Data) ProtoMessage()    {}
func (*TranslateStormHandlerResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{1, 0}
}
func (m *TranslateStormHandlerResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data.Unmarshal(m, b)
}
func (m *TranslateStormHandlerResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerResponse_Data.Merge(m, src)
}
func (m *TranslateStormHandlerResponse_Data) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data.Size(m)
}
func (m *TranslateStormHandlerResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerResponse_Data proto.InternalMessageInfo

func (m *TranslateStormHandlerResponse_Data) GetMatchRules() *TranslateStormHandlerResponse_Data_MatchRules {
	if m != nil {
		return m.MatchRules
	}
	return nil
}

func (m *TranslateStormHandlerResponse_Data) GetRuleIds() []string {
	if m != nil {
		return m.RuleIds
	}
	return nil
}

type TranslateStormHandlerResponse_Data_MatchRules struct {
	//
	//匹配字段分隔符
	SepFieldValue string `protobuf:"bytes,1,opt,name=sep_field_value,json=sepFieldValue,proto3" json:"sep_field_value" form:"sep_field_value"`
	//
	//命名空间分隔符
	SepNameSpace string `protobuf:"bytes,2,opt,name=sep_name_space,json=sepNameSpace,proto3" json:"sep_name_space" form:"sep_name_space"`
	//
	//匹配字段
	MatchFields          []*TranslateStormHandlerResponse_Data_MatchRules_MatchFields `protobuf:"bytes,3,rep,name=match_fields,json=matchFields,proto3" json:"match_fields" form:"match_fields"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *TranslateStormHandlerResponse_Data_MatchRules) Reset() {
	*m = TranslateStormHandlerResponse_Data_MatchRules{}
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) String() string {
	return proto.CompactTextString(m)
}
func (*TranslateStormHandlerResponse_Data_MatchRules) ProtoMessage() {}
func (*TranslateStormHandlerResponse_Data_MatchRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{1, 0, 0}
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules.Unmarshal(m, b)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules.Merge(m, src)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules.Size(m)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules proto.InternalMessageInfo

func (m *TranslateStormHandlerResponse_Data_MatchRules) GetSepFieldValue() string {
	if m != nil {
		return m.SepFieldValue
	}
	return ""
}

func (m *TranslateStormHandlerResponse_Data_MatchRules) GetSepNameSpace() string {
	if m != nil {
		return m.SepNameSpace
	}
	return ""
}

func (m *TranslateStormHandlerResponse_Data_MatchRules) GetMatchFields() []*TranslateStormHandlerResponse_Data_MatchRules_MatchFields {
	if m != nil {
		return m.MatchFields
	}
	return nil
}

type TranslateStormHandlerResponse_Data_MatchRules_MatchFields struct {
	//
	//字段列表
	Fields []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields" form:"fields"`
	//
	//规则id
	RuleId string `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id" form:"rule_id"`
	//
	//关联模型id
	ObjectId             string   `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) Reset() {
	*m = TranslateStormHandlerResponse_Data_MatchRules_MatchFields{}
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) String() string {
	return proto.CompactTextString(m)
}
func (*TranslateStormHandlerResponse_Data_MatchRules_MatchFields) ProtoMessage() {}
func (*TranslateStormHandlerResponse_Data_MatchRules_MatchFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{1, 0, 0, 0}
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields.Unmarshal(m, b)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields.Merge(m, src)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields.Size(m)
}
func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerResponse_Data_MatchRules_MatchFields proto.InternalMessageInfo

func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) GetRuleId() string {
	if m != nil {
		return m.RuleId
	}
	return ""
}

func (m *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

//
//TranslateStormHandlerApi返回
type TranslateStormHandlerResponseWrapper struct {
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
	Data                 *TranslateStormHandlerResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TranslateStormHandlerResponseWrapper) Reset()         { *m = TranslateStormHandlerResponseWrapper{} }
func (m *TranslateStormHandlerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*TranslateStormHandlerResponseWrapper) ProtoMessage()    {}
func (*TranslateStormHandlerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e91050f46a7a30, []int{2}
}
func (m *TranslateStormHandlerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateStormHandlerResponseWrapper.Unmarshal(m, b)
}
func (m *TranslateStormHandlerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateStormHandlerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *TranslateStormHandlerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateStormHandlerResponseWrapper.Merge(m, src)
}
func (m *TranslateStormHandlerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_TranslateStormHandlerResponseWrapper.Size(m)
}
func (m *TranslateStormHandlerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateStormHandlerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateStormHandlerResponseWrapper proto.InternalMessageInfo

func (m *TranslateStormHandlerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TranslateStormHandlerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *TranslateStormHandlerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TranslateStormHandlerResponseWrapper) GetData() *TranslateStormHandlerResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TranslateStormHandlerRequest)(nil), "translate_rule.TranslateStormHandlerRequest")
	proto.RegisterType((*TranslateStormHandlerResponse)(nil), "translate_rule.TranslateStormHandlerResponse")
	proto.RegisterType((*TranslateStormHandlerResponse_Data)(nil), "translate_rule.TranslateStormHandlerResponse.Data")
	proto.RegisterType((*TranslateStormHandlerResponse_Data_MatchRules)(nil), "translate_rule.TranslateStormHandlerResponse.Data.MatchRules")
	proto.RegisterType((*TranslateStormHandlerResponse_Data_MatchRules_MatchFields)(nil), "translate_rule.TranslateStormHandlerResponse.Data.MatchRules.MatchFields")
	proto.RegisterType((*TranslateStormHandlerResponseWrapper)(nil), "translate_rule.TranslateStormHandlerResponseWrapper")
}

func init() { proto.RegisterFile("translate_storm_handler.proto", fileDescriptor_f6e91050f46a7a30) }

var fileDescriptor_f6e91050f46a7a30 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x6e, 0x7e, 0x9b, 0x4c, 0xda, 0x34, 0x4e, 0xb5, 0x5d, 0x83, 0x65, 0xc3, 0x58, 0x24, 0x62,
	0x93, 0xb4, 0x29, 0x58, 0x15, 0x44, 0x0c, 0x2a, 0xf6, 0x42, 0xc1, 0x69, 0xb1, 0x60, 0xd1, 0x65,
	0x92, 0x9d, 0xa6, 0xab, 0xbb, 0x99, 0x75, 0x67, 0xd2, 0x4a, 0xc5, 0x0b, 0x2f, 0x04, 0x5f, 0x49,
	0x7c, 0x10, 0xf1, 0x66, 0x05, 0x1f, 0x61, 0x9f, 0x40, 0x66, 0x66, 0x93, 0xdd, 0x0a, 0x2d, 0x16,
	0xaf, 0x72, 0xf6, 0x7c, 0xdf, 0xf9, 0xce, 0x99, 0xf3, 0x13, 0xb0, 0x22, 0x02, 0x32, 0xe2, 0x2e,
	0x11, 0xd4, 0xe2, 0x82, 0x05, 0x9e, 0x75, 0x48, 0x46, 0xb6, 0x4b, 0x83, 0xb6, 0x1f, 0x30, 0xc1,
	0x60, 0x35, 0x81, 0x83, 0xb1, 0x4b, 0xeb, 0xad, 0xa1, 0x23, 0x0e, 0xc7, 0xfd, 0xf6, 0x80, 0x79,
	0x9d, 0x21, 0x1b, 0xb2, 0x8e, 0xa2, 0xf5, 0xc7, 0x07, 0xea, 0x4b, 0x7d, 0x28, 0x4b, 0x87, 0xd7,
	0x6f, 0xa7, 0xe8, 0xde, 0xb1, 0x23, 0xde, 0xb1, 0xe3, 0xce, 0x90, 0xb5, 0x14, 0xd8, 0x3a, 0x22,
	0xae, 0x63, 0x13, 0xc1, 0x02, 0xde, 0x99, 0x9a, 0x3a, 0x0e, 0xbd, 0x00, 0xd7, 0x76, 0x27, 0x89,
	0x77, 0x64, 0x59, 0x4f, 0x75, 0x55, 0x98, 0xbe, 0x1f, 0x53, 0x2e, 0xe0, 0x06, 0x28, 0xdb, 0x44,
	0x10, 0x6b, 0x44, 0x3c, 0x6a, 0x64, 0x1a, 0x99, 0x66, 0xb9, 0x77, 0x39, 0x0a, 0xcd, 0xda, 0x01,
	0x0b, 0xbc, 0x7b, 0x68, 0x0a, 0x21, 0x5c, 0x92, 0xf6, 0x73, 0x69, 0xfe, 0x9c, 0x05, 0x2b, 0x67,
	0x68, 0x72, 0x9f, 0x8d, 0x38, 0x85, 0xd7, 0x41, 0x7e, 0xc0, 0x6c, 0xad, 0x57, 0xe8, 0x2d, 0x44,
	0xa1, 0x59, 0xd1, 0x7a, 0xd2, 0x8b, 0xb0, 0x02, 0xe1, 0x16, 0x28, 0xfb, 0x64, 0x48, 0x2d, 0xee,
	0x9c, 0x50, 0x23, 0xab, 0x98, 0xf5, 0x24, 0xf3, 0x14, 0x42, 0xbf, 0x7f, 0x99, 0xd9, 0xda, 0x0c,
	0x2e, 0x49, 0xcf, 0x8e, 0x73, 0x42, 0x61, 0x03, 0xe4, 0x3c, 0x3e, 0x34, 0x72, 0xaa, 0xd8, 0x6a,
	0x14, 0x9a, 0x40, 0x87, 0x78, 0x7c, 0x88, 0xb0, 0x84, 0xe0, 0x0d, 0x50, 0x10, 0x4c, 0x10, 0xd7,
	0xc8, 0x2b, 0xd9, 0x5a, 0x14, 0x9a, 0x73, 0x9a, 0xa3, 0xdc, 0x08, 0x6b, 0x18, 0xae, 0x81, 0xbc,
	0x54, 0x35, 0x0a, 0x8a, 0x66, 0x24, 0x75, 0x4a, 0xaf, 0x4c, 0x5c, 0xa8, 0xcd, 0x18, 0xdf, 0x96,
	0xb1, 0x62, 0xc1, 0x3d, 0x90, 0x97, 0x3d, 0x30, 0x8a, 0x8d, 0x4c, 0xb3, 0xd2, 0xed, 0xb6, 0x4f,
	0x0f, 0xb4, 0x7d, 0x6e, 0x4b, 0xda, 0x8f, 0x88, 0x20, 0xe9, 0x4e, 0x48, 0x25, 0x84, 0x95, 0x60,
	0xfd, 0x6b, 0x01, 0xe4, 0x25, 0x0e, 0x8f, 0x40, 0xc5, 0x23, 0x62, 0x70, 0xa8, 0x04, 0xb9, 0x6a,
	0x5f, 0xa5, 0x7b, 0xff, 0xe2, 0x89, 0xda, 0xcf, 0xa4, 0x0a, 0x96, 0x22, 0xbd, 0xa5, 0x28, 0x34,
	0x61, 0xdc, 0xa0, 0x44, 0x1b, 0x61, 0xe0, 0x4d, 0x39, 0xb0, 0x0d, 0x4a, 0xd2, 0x6b, 0x39, 0x36,
	0x37, 0xb2, 0x8d, 0x5c, 0xb3, 0xdc, 0x5b, 0x8c, 0x42, 0x73, 0x41, 0x47, 0x4d, 0x10, 0x84, 0x67,
	0xa5, 0xb9, 0x6d, 0xf3, 0xfa, 0x8f, 0x1c, 0x00, 0x49, 0x0a, 0xd8, 0x03, 0x0b, 0x9c, 0xfa, 0xd6,
	0x81, 0x43, 0x5d, 0xdb, 0x3a, 0x22, 0xee, 0x78, 0xb2, 0x49, 0x72, 0x9e, 0x4b, 0x5a, 0xe5, 0x2f,
	0x02, 0xc2, 0xf3, 0x9c, 0xfa, 0x4f, 0xa4, 0xe3, 0xa5, 0xfc, 0x86, 0x0f, 0x40, 0x55, 0x52, 0xe4,
	0xae, 0x59, 0xdc, 0x27, 0x03, 0xbd, 0x12, 0xe5, 0xde, 0xd5, 0x28, 0x34, 0xaf, 0x24, 0x12, 0x09,
	0x8e, 0xf0, 0x1c, 0xa7, 0xbe, 0x5c, 0xc8, 0x1d, 0xf9, 0x09, 0xbf, 0x64, 0xc0, 0x9c, 0x7e, 0xa0,
	0x4a, 0xc3, 0x8d, 0x5c, 0x23, 0xd7, 0xac, 0x74, 0xb7, 0xff, 0xab, 0x7b, 0xda, 0x54, 0x65, 0xf2,
	0xde, 0x72, 0x14, 0x9a, 0x8b, 0xe9, 0x4e, 0xea, 0x44, 0x08, 0xeb, 0xa1, 0x69, 0x56, 0xfd, 0x7b,
	0x06, 0x54, 0x52, 0x51, 0xf0, 0x26, 0x28, 0xc6, 0x05, 0x65, 0x54, 0x67, 0x2f, 0x45, 0xa1, 0x39,
	0xaf, 0x55, 0x26, 0xf1, 0x31, 0x01, 0xde, 0x02, 0xb3, 0x71, 0xb3, 0xe3, 0xc7, 0xc3, 0x28, 0x34,
	0xab, 0xa7, 0xa6, 0x80, 0x70, 0x51, 0x0f, 0x01, 0xee, 0x82, 0x32, 0xeb, 0xbf, 0xa5, 0x03, 0x21,
	0xe9, 0xfa, 0x16, 0xb6, 0x92, 0xf3, 0x99, 0x42, 0x72, 0x8b, 0x4d, 0xb0, 0xf2, 0x66, 0x9f, 0xb4,
	0x4e, 0x1e, 0xb6, 0x5e, 0x59, 0xaf, 0xf7, 0xd7, 0x5b, 0x77, 0x27, 0xf6, 0xc7, 0xf5, 0xb5, 0xcd,
	0x8d, 0x4f, 0xab, 0xb8, 0xa4, 0xe9, 0xdb, 0x36, 0xfa, 0x9c, 0x05, 0xab, 0xe7, 0x76, 0x68, 0x2f,
	0x20, 0xbe, 0x4f, 0x83, 0x7f, 0x3b, 0xf1, 0x3b, 0xa0, 0x22, 0x7f, 0x1f, 0x7f, 0xf0, 0x5d, 0xe2,
	0x8c, 0xe2, 0x47, 0xa5, 0x16, 0x32, 0x05, 0x22, 0x9c, 0xa6, 0xca, 0x0b, 0xa6, 0x41, 0xc0, 0x82,
	0xf8, 0x65, 0xa9, 0x0b, 0x56, 0x6e, 0x84, 0x35, 0x0c, 0x71, 0x7c, 0x93, 0x79, 0x75, 0x2a, 0xad,
	0x0b, 0x0d, 0xfb, 0x8c, 0x73, 0xec, 0x17, 0xd5, 0x3f, 0xe7, 0xe6, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x46, 0xf3, 0xae, 0xd1, 0x05, 0x00, 0x00,
}
