// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_translate_data.proto

package translate

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
//ListTranslateData请求
type ListTranslateDataRequest struct {
	//
	//翻译规则id列表
	RuleIds              []string `protobuf:"bytes,1,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids" form:"rule_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTranslateDataRequest) Reset()         { *m = ListTranslateDataRequest{} }
func (m *ListTranslateDataRequest) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataRequest) ProtoMessage()    {}
func (*ListTranslateDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{0}
}
func (m *ListTranslateDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataRequest.Unmarshal(m, b)
}
func (m *ListTranslateDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataRequest.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataRequest.Merge(m, src)
}
func (m *ListTranslateDataRequest) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataRequest.Size(m)
}
func (m *ListTranslateDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataRequest proto.InternalMessageInfo

func (m *ListTranslateDataRequest) GetRuleIds() []string {
	if m != nil {
		return m.RuleIds
	}
	return nil
}

//
//ListTranslateData返回
type ListTranslateDataResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//分页参数
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//页数
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//返回信息
	Msg string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//总数
	Total int32 `protobuf:"varint,5,opt,name=total,proto3" json:"total" form:"total"`
	//
	//返回数据
	Data                 []*ListTranslateDataResponse_Data `protobuf:"bytes,6,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListTranslateDataResponse) Reset()         { *m = ListTranslateDataResponse{} }
func (m *ListTranslateDataResponse) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataResponse) ProtoMessage()    {}
func (*ListTranslateDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{1}
}
func (m *ListTranslateDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataResponse.Unmarshal(m, b)
}
func (m *ListTranslateDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataResponse.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataResponse.Merge(m, src)
}
func (m *ListTranslateDataResponse) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataResponse.Size(m)
}
func (m *ListTranslateDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataResponse proto.InternalMessageInfo

func (m *ListTranslateDataResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListTranslateDataResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTranslateDataResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTranslateDataResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ListTranslateDataResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListTranslateDataResponse) GetData() []*ListTranslateDataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListTranslateDataResponse_Data struct {
	//
	//匹配key
	MatchKey string `protobuf:"bytes,1,opt,name=matchKey,proto3" json:"matchKey" form:"matchKey"`
	//
	//关联资源
	Resource *ListTranslateDataResponse_Data_Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource" form:"resource"`
	//
	//标签
	Tags                 []*ListTranslateDataResponse_Data_Tags `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags" form:"tags"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ListTranslateDataResponse_Data) Reset()         { *m = ListTranslateDataResponse_Data{} }
func (m *ListTranslateDataResponse_Data) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataResponse_Data) ProtoMessage()    {}
func (*ListTranslateDataResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{1, 0}
}
func (m *ListTranslateDataResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataResponse_Data.Unmarshal(m, b)
}
func (m *ListTranslateDataResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataResponse_Data.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataResponse_Data.Merge(m, src)
}
func (m *ListTranslateDataResponse_Data) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataResponse_Data.Size(m)
}
func (m *ListTranslateDataResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataResponse_Data proto.InternalMessageInfo

func (m *ListTranslateDataResponse_Data) GetMatchKey() string {
	if m != nil {
		return m.MatchKey
	}
	return ""
}

func (m *ListTranslateDataResponse_Data) GetResource() *ListTranslateDataResponse_Data_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ListTranslateDataResponse_Data) GetTags() []*ListTranslateDataResponse_Data_Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ListTranslateDataResponse_Data_Resource struct {
	//
	//实例id列表
	Instances []string `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances" form:"instances"`
	//
	//关联模型Id
	ObjectId             string   `protobuf:"bytes,2,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTranslateDataResponse_Data_Resource) Reset() {
	*m = ListTranslateDataResponse_Data_Resource{}
}
func (m *ListTranslateDataResponse_Data_Resource) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataResponse_Data_Resource) ProtoMessage()    {}
func (*ListTranslateDataResponse_Data_Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{1, 0, 0}
}
func (m *ListTranslateDataResponse_Data_Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Resource.Unmarshal(m, b)
}
func (m *ListTranslateDataResponse_Data_Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Resource.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataResponse_Data_Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataResponse_Data_Resource.Merge(m, src)
}
func (m *ListTranslateDataResponse_Data_Resource) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Resource.Size(m)
}
func (m *ListTranslateDataResponse_Data_Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataResponse_Data_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataResponse_Data_Resource proto.InternalMessageInfo

func (m *ListTranslateDataResponse_Data_Resource) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *ListTranslateDataResponse_Data_Resource) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type ListTranslateDataResponse_Data_Tags struct {
	//
	//值
	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value" form:"value"`
	//
	//键
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key" form:"key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTranslateDataResponse_Data_Tags) Reset()         { *m = ListTranslateDataResponse_Data_Tags{} }
func (m *ListTranslateDataResponse_Data_Tags) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataResponse_Data_Tags) ProtoMessage()    {}
func (*ListTranslateDataResponse_Data_Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{1, 0, 1}
}
func (m *ListTranslateDataResponse_Data_Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Tags.Unmarshal(m, b)
}
func (m *ListTranslateDataResponse_Data_Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Tags.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataResponse_Data_Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataResponse_Data_Tags.Merge(m, src)
}
func (m *ListTranslateDataResponse_Data_Tags) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataResponse_Data_Tags.Size(m)
}
func (m *ListTranslateDataResponse_Data_Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataResponse_Data_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataResponse_Data_Tags proto.InternalMessageInfo

func (m *ListTranslateDataResponse_Data_Tags) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ListTranslateDataResponse_Data_Tags) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

//
//ListTranslateDataApi返回
type ListTranslateDataResponseWrapper struct {
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
	Data                 *ListTranslateDataResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListTranslateDataResponseWrapper) Reset()         { *m = ListTranslateDataResponseWrapper{} }
func (m *ListTranslateDataResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListTranslateDataResponseWrapper) ProtoMessage()    {}
func (*ListTranslateDataResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445ac8065905fc2, []int{2}
}
func (m *ListTranslateDataResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTranslateDataResponseWrapper.Unmarshal(m, b)
}
func (m *ListTranslateDataResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTranslateDataResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListTranslateDataResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTranslateDataResponseWrapper.Merge(m, src)
}
func (m *ListTranslateDataResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListTranslateDataResponseWrapper.Size(m)
}
func (m *ListTranslateDataResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTranslateDataResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListTranslateDataResponseWrapper proto.InternalMessageInfo

func (m *ListTranslateDataResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListTranslateDataResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListTranslateDataResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListTranslateDataResponseWrapper) GetData() *ListTranslateDataResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListTranslateDataRequest)(nil), "translate.ListTranslateDataRequest")
	proto.RegisterType((*ListTranslateDataResponse)(nil), "translate.ListTranslateDataResponse")
	proto.RegisterType((*ListTranslateDataResponse_Data)(nil), "translate.ListTranslateDataResponse.Data")
	proto.RegisterType((*ListTranslateDataResponse_Data_Resource)(nil), "translate.ListTranslateDataResponse.Data.Resource")
	proto.RegisterType((*ListTranslateDataResponse_Data_Tags)(nil), "translate.ListTranslateDataResponse.Data.Tags")
	proto.RegisterType((*ListTranslateDataResponseWrapper)(nil), "translate.ListTranslateDataResponseWrapper")
}

func init() { proto.RegisterFile("list_translate_data.proto", fileDescriptor_5445ac8065905fc2) }

var fileDescriptor_5445ac8065905fc2 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x6d, 0x9a, 0x4d, 0x9b, 0x4c, 0xa4, 0x0d, 0x23, 0xca, 0x36, 0x50, 0x36, 0x8c, 0x45, 0x52,
	0x70, 0x37, 0xfd, 0x80, 0xa2, 0x7d, 0x33, 0xa8, 0x50, 0x15, 0x91, 0x69, 0x41, 0xb0, 0xb4, 0x61,
	0xb2, 0x99, 0x6e, 0xd7, 0x6e, 0x32, 0xeb, 0xcc, 0x6c, 0x6b, 0x53, 0xfa, 0x77, 0x7c, 0xf6, 0x17,
	0xad, 0xe0, 0x9b, 0xe0, 0xd3, 0xe2, 0x0f, 0x90, 0x99, 0xfd, 0xc8, 0x16, 0x2c, 0xe4, 0x29, 0x33,
	0xf7, 0x9c, 0x7b, 0xee, 0xdc, 0x7b, 0xcf, 0x06, 0xac, 0x05, 0xbe, 0x90, 0x03, 0xc9, 0xc9, 0x44,
	0x04, 0x44, 0xd2, 0xc1, 0x88, 0x48, 0xe2, 0x84, 0x9c, 0x49, 0x06, 0x1b, 0x45, 0xb4, 0x6d, 0x7b,
	0xbe, 0x3c, 0x8f, 0x86, 0x8e, 0xcb, 0xc6, 0x3d, 0x8f, 0x79, 0xac, 0xa7, 0x19, 0xc3, 0xe8, 0x4c,
	0xdf, 0xf4, 0x45, 0x9f, 0xd2, 0xcc, 0xf6, 0x5e, 0x89, 0x3e, 0xbe, 0xf2, 0xe5, 0x05, 0xbb, 0xea,
	0x79, 0xcc, 0xd6, 0xa0, 0x7d, 0x49, 0x02, 0x7f, 0x44, 0x24, 0xe3, 0xa2, 0x57, 0x1c, 0xd3, 0x3c,
	0xf4, 0x16, 0x98, 0xef, 0x7d, 0x21, 0x8f, 0xf2, 0xba, 0xaf, 0x88, 0x24, 0x98, 0x7e, 0x8d, 0xa8,
	0x90, 0xd0, 0x01, 0x75, 0x1e, 0x05, 0x74, 0xe0, 0x8f, 0x84, 0x59, 0xe9, 0x54, 0xbb, 0x8d, 0xfe,
	0xc3, 0x24, 0xb6, 0x56, 0xcf, 0x18, 0x1f, 0xef, 0xa3, 0x1c, 0x41, 0x78, 0x59, 0x1d, 0x0f, 0x46,
	0x02, 0xfd, 0x58, 0x02, 0x6b, 0xff, 0x11, 0x13, 0x21, 0x9b, 0x08, 0x0a, 0x9f, 0x00, 0xc3, 0x65,
	0x23, 0x6a, 0x56, 0x3a, 0x95, 0x6e, 0xad, 0xbf, 0x9a, 0xc4, 0x56, 0x33, 0x55, 0x52, 0x51, 0x84,
	0x35, 0x08, 0xf7, 0x41, 0x23, 0x24, 0x1e, 0x1d, 0x08, 0x7f, 0x4a, 0xcd, 0x45, 0xcd, 0x5c, 0x4f,
	0x62, 0xab, 0x95, 0x32, 0x0b, 0x08, 0xfd, 0xfa, 0x69, 0xd5, 0x5a, 0x0b, 0xe6, 0xef, 0x65, 0x5c,
	0x57, 0xc1, 0x43, 0x7f, 0x4a, 0xe1, 0x26, 0x30, 0xd4, 0xd9, 0xac, 0xea, 0xb4, 0x47, 0xb3, 0x02,
	0x2a, 0xaa, 0x32, 0x16, 0x5b, 0x0b, 0x58, 0x53, 0x60, 0x07, 0x54, 0xc7, 0xc2, 0x33, 0x8d, 0x4e,
	0xa5, 0xdb, 0xe8, 0xaf, 0x24, 0xb1, 0x05, 0x52, 0xe6, 0x58, 0x78, 0x08, 0x2b, 0x08, 0x3e, 0x05,
	0x35, 0xc9, 0x24, 0x09, 0xcc, 0x9a, 0x56, 0x6b, 0x25, 0xb1, 0xf5, 0x20, 0xe5, 0xe8, 0x30, 0xc2,
	0x29, 0x0c, 0x3f, 0x00, 0x43, 0xed, 0xcf, 0x5c, 0xea, 0x54, 0xbb, 0xcd, 0x9d, 0x4d, 0xa7, 0x58,
	0xa0, 0x73, 0xef, 0x24, 0x1c, 0x75, 0x29, 0x0f, 0x40, 0x09, 0x20, 0xac, 0x75, 0xda, 0x7f, 0xaa,
	0xc0, 0x50, 0x38, 0xec, 0x81, 0xfa, 0x98, 0x48, 0xf7, 0xfc, 0x1d, 0xbd, 0xd6, 0x23, 0xbb, 0x33,
	0xfc, 0x1c, 0x41, 0xb8, 0x20, 0x41, 0x17, 0xd4, 0x39, 0x15, 0x2c, 0xe2, 0x6e, 0x3a, 0xb9, 0xe6,
	0xce, 0xce, 0xdc, 0xaf, 0x71, 0x70, 0x96, 0x79, 0x67, 0xc3, 0x59, 0x0c, 0xe1, 0x42, 0x18, 0x1e,
	0x02, 0x43, 0x12, 0x4f, 0x98, 0x55, 0xdd, 0xae, 0x33, 0x7f, 0x81, 0x23, 0xe2, 0x89, 0x72, 0xcf,
	0x4a, 0x05, 0x61, 0x2d, 0xd6, 0xfe, 0x5e, 0x01, 0xf5, 0xfc, 0x01, 0xf0, 0x0d, 0x68, 0xf8, 0x13,
	0x21, 0xc9, 0xc4, 0xa5, 0xb9, 0xeb, 0xba, 0x33, 0x07, 0x14, 0x90, 0xda, 0x67, 0x0b, 0xac, 0x9c,
	0x1e, 0x6f, 0xd9, 0x2f, 0x88, 0x3d, 0x3d, 0xb9, 0xd9, 0xde, 0xbd, 0xdd, 0xc0, 0xb3, 0x54, 0x88,
	0x41, 0x9d, 0x0d, 0xbf, 0x50, 0x57, 0x1e, 0x8c, 0xf4, 0x38, 0x1a, 0xfd, 0xbd, 0x59, 0x6b, 0x39,
	0xa2, 0x54, 0x2c, 0xb0, 0x7e, 0x7a, 0x4c, 0xec, 0xe9, 0x4b, 0xfb, 0xf3, 0xe0, 0x24, 0x93, 0xd3,
	0xe7, 0x9b, 0xad, 0x67, 0xbb, 0xdb, 0xb7, 0x1b, 0xb8, 0xd0, 0x69, 0x7f, 0x04, 0x86, 0xea, 0x43,
	0x99, 0xe3, 0x92, 0x04, 0x11, 0xcd, 0xde, 0x57, 0x32, 0x87, 0x0e, 0x23, 0x9c, 0xc2, 0xca, 0x66,
	0x17, 0xf4, 0x3a, 0x2b, 0x5f, 0xb2, 0xd9, 0x85, 0xda, 0x9c, 0x82, 0xd0, 0xdf, 0x0a, 0xe8, 0xdc,
	0x3b, 0xb9, 0x4f, 0x9c, 0x84, 0x21, 0xe5, 0xf3, 0x7d, 0x39, 0xcf, 0x41, 0x53, 0xfd, 0xbe, 0xfe,
	0x16, 0x06, 0xc4, 0x9f, 0x64, 0x35, 0x1f, 0x27, 0xb1, 0x05, 0x67, 0xdc, 0x0c, 0x44, 0xb8, 0x4c,
	0x55, 0xdd, 0x50, 0xce, 0x19, 0xd7, 0x1f, 0xce, 0x9d, 0x6e, 0x74, 0x18, 0xe1, 0x14, 0x86, 0x07,
	0x99, 0xd5, 0x0d, 0x6d, 0xae, 0x8d, 0x79, 0x76, 0x7f, 0x8f, 0xcb, 0x87, 0x4b, 0xfa, 0xcf, 0x67,
	0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x5a, 0xdd, 0xa3, 0x0b, 0x05, 0x00, 0x00,
}