// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: translate_rule.proto

package monitor

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
//翻译规则
type TransalteRule struct {
	//
	//类别
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category" form:"category"`
	//
	//翻译规则名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//翻译字段
	TranslateFields []*TransalteRule_TranslateFields `protobuf:"bytes,3,rep,name=translate_fields,json=translateFields,proto3" json:"translate_fields" form:"translate_fields"`
	//
	//关联模型id
	ObjectId string `protobuf:"bytes,4,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//匹配规则
	MatchFields []*TransalteRule_MatchFields `protobuf:"bytes,5,rep,name=match_fields,json=matchFields,proto3" json:"match_fields" form:"match_fields"`
	//
	//类型
	Type string `protobuf:"bytes,6,opt,name=type,proto3" json:"type" form:"type"`
	//
	//数据源
	DataNames []string `protobuf:"bytes,7,rep,name=data_names,json=dataNames,proto3" json:"data_names" form:"data_names"`
	//
	//数据获取方法
	Fetcher string `protobuf:"bytes,8,opt,name=fetcher,proto3" json:"fetcher" form:"fetcher"`
	//
	//是否受保护
	Protected bool `protobuf:"varint,9,opt,name=protected,proto3" json:"protected" form:"protected"`
	//
	//是否启用
	Disabled bool `protobuf:"varint,10,opt,name=disabled,proto3" json:"disabled" form:"disabled"`
	//
	//org
	Org int32 `protobuf:"varint,11,opt,name=org,proto3" json:"org" form:"org"`
	//
	//翻译规则ID
	XId                  string   `protobuf:"bytes,12,opt,name=_id,json=Id,proto3" json:"_id" form:"_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransalteRule) Reset()         { *m = TransalteRule{} }
func (m *TransalteRule) String() string { return proto.CompactTextString(m) }
func (*TransalteRule) ProtoMessage()    {}
func (*TransalteRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6edd93aa3b1cdb5f, []int{0}
}
func (m *TransalteRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransalteRule.Unmarshal(m, b)
}
func (m *TransalteRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransalteRule.Marshal(b, m, deterministic)
}
func (m *TransalteRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransalteRule.Merge(m, src)
}
func (m *TransalteRule) XXX_Size() int {
	return xxx_messageInfo_TransalteRule.Size(m)
}
func (m *TransalteRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransalteRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransalteRule proto.InternalMessageInfo

func (m *TransalteRule) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *TransalteRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransalteRule) GetTranslateFields() []*TransalteRule_TranslateFields {
	if m != nil {
		return m.TranslateFields
	}
	return nil
}

func (m *TransalteRule) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *TransalteRule) GetMatchFields() []*TransalteRule_MatchFields {
	if m != nil {
		return m.MatchFields
	}
	return nil
}

func (m *TransalteRule) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TransalteRule) GetDataNames() []string {
	if m != nil {
		return m.DataNames
	}
	return nil
}

func (m *TransalteRule) GetFetcher() string {
	if m != nil {
		return m.Fetcher
	}
	return ""
}

func (m *TransalteRule) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *TransalteRule) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *TransalteRule) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *TransalteRule) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

type TransalteRule_TranslateFields struct {
	//
	//输出名称
	StreamKey string `protobuf:"bytes,1,opt,name=stream_key,json=streamKey,proto3" json:"stream_key" form:"stream_key"`
	//
	//资源字段
	ObjectKey            string   `protobuf:"bytes,2,opt,name=object_key,json=objectKey,proto3" json:"object_key" form:"object_key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransalteRule_TranslateFields) Reset()         { *m = TransalteRule_TranslateFields{} }
func (m *TransalteRule_TranslateFields) String() string { return proto.CompactTextString(m) }
func (*TransalteRule_TranslateFields) ProtoMessage()    {}
func (*TransalteRule_TranslateFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_6edd93aa3b1cdb5f, []int{0, 0}
}
func (m *TransalteRule_TranslateFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransalteRule_TranslateFields.Unmarshal(m, b)
}
func (m *TransalteRule_TranslateFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransalteRule_TranslateFields.Marshal(b, m, deterministic)
}
func (m *TransalteRule_TranslateFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransalteRule_TranslateFields.Merge(m, src)
}
func (m *TransalteRule_TranslateFields) XXX_Size() int {
	return xxx_messageInfo_TransalteRule_TranslateFields.Size(m)
}
func (m *TransalteRule_TranslateFields) XXX_DiscardUnknown() {
	xxx_messageInfo_TransalteRule_TranslateFields.DiscardUnknown(m)
}

var xxx_messageInfo_TransalteRule_TranslateFields proto.InternalMessageInfo

func (m *TransalteRule_TranslateFields) GetStreamKey() string {
	if m != nil {
		return m.StreamKey
	}
	return ""
}

func (m *TransalteRule_TranslateFields) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

type TransalteRule_MatchFields struct {
	//
	//上报维度
	StreamKey string `protobuf:"bytes,1,opt,name=stream_key,json=streamKey,proto3" json:"stream_key" form:"stream_key"`
	//
	//资源字段
	ObjectKey            string   `protobuf:"bytes,2,opt,name=object_key,json=objectKey,proto3" json:"object_key" form:"object_key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransalteRule_MatchFields) Reset()         { *m = TransalteRule_MatchFields{} }
func (m *TransalteRule_MatchFields) String() string { return proto.CompactTextString(m) }
func (*TransalteRule_MatchFields) ProtoMessage()    {}
func (*TransalteRule_MatchFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_6edd93aa3b1cdb5f, []int{0, 1}
}
func (m *TransalteRule_MatchFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransalteRule_MatchFields.Unmarshal(m, b)
}
func (m *TransalteRule_MatchFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransalteRule_MatchFields.Marshal(b, m, deterministic)
}
func (m *TransalteRule_MatchFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransalteRule_MatchFields.Merge(m, src)
}
func (m *TransalteRule_MatchFields) XXX_Size() int {
	return xxx_messageInfo_TransalteRule_MatchFields.Size(m)
}
func (m *TransalteRule_MatchFields) XXX_DiscardUnknown() {
	xxx_messageInfo_TransalteRule_MatchFields.DiscardUnknown(m)
}

var xxx_messageInfo_TransalteRule_MatchFields proto.InternalMessageInfo

func (m *TransalteRule_MatchFields) GetStreamKey() string {
	if m != nil {
		return m.StreamKey
	}
	return ""
}

func (m *TransalteRule_MatchFields) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

func init() {
	proto.RegisterType((*TransalteRule)(nil), "monitor.TransalteRule")
	proto.RegisterType((*TransalteRule_TranslateFields)(nil), "monitor.TransalteRule.TranslateFields")
	proto.RegisterType((*TransalteRule_MatchFields)(nil), "monitor.TransalteRule.MatchFields")
}

func init() { proto.RegisterFile("translate_rule.proto", fileDescriptor_6edd93aa3b1cdb5f) }

var fileDescriptor_6edd93aa3b1cdb5f = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xd1, 0x6a, 0xdb, 0x3c,
	0x14, 0xc7, 0x49, 0xd3, 0x36, 0xb1, 0xd2, 0xaf, 0xe9, 0xa7, 0x76, 0xd4, 0x64, 0x0c, 0x1b, 0x6d,
	0x8c, 0x5c, 0xd4, 0x76, 0xd7, 0x8e, 0x8d, 0xed, 0xae, 0x1e, 0x0c, 0x4a, 0xd9, 0x2e, 0x44, 0xaf,
	0x5a, 0xd6, 0xa0, 0xd8, 0x8a, 0xeb, 0xd5, 0x8e, 0x82, 0xac, 0xac, 0x78, 0xa3, 0xaf, 0xb5, 0xc7,
	0xf1, 0x60, 0x8f, 0xe0, 0x27, 0x18, 0x92, 0xec, 0xd8, 0x0d, 0xec, 0x76, 0x77, 0x3a, 0xfe, 0xff,
	0xce, 0xd1, 0x39, 0x47, 0x7f, 0x0c, 0x0e, 0x04, 0x27, 0xf3, 0x2c, 0x21, 0x82, 0x4e, 0xf8, 0x32,
	0xa1, 0xee, 0x82, 0x33, 0xc1, 0x60, 0x2f, 0x65, 0xf3, 0x58, 0x30, 0x3e, 0x72, 0xa2, 0x58, 0xdc,
	0x2e, 0xa7, 0x6e, 0xc0, 0x52, 0x2f, 0x62, 0x11, 0xf3, 0x94, 0x3e, 0x5d, 0xce, 0x54, 0xa4, 0x02,
	0x75, 0xd2, 0x79, 0xa3, 0x37, 0x2d, 0x3c, 0xbd, 0x8f, 0xc5, 0x1d, 0xbb, 0xf7, 0x22, 0xe6, 0x28,
	0xd1, 0xf9, 0x46, 0x92, 0x38, 0x24, 0x82, 0xf1, 0xcc, 0x5b, 0x1d, 0x75, 0x1e, 0xfa, 0xd9, 0x03,
	0xff, 0x5d, 0xca, 0x46, 0x48, 0x22, 0x28, 0x5e, 0x26, 0x14, 0x7a, 0xa0, 0x1f, 0x10, 0x41, 0x23,
	0xc6, 0x73, 0xb3, 0x63, 0x77, 0xc6, 0x86, 0xbf, 0x5f, 0x16, 0xd6, 0x70, 0xc6, 0x78, 0xfa, 0x1e,
	0xd5, 0x0a, 0xc2, 0x2b, 0x08, 0x3e, 0x07, 0x9b, 0x73, 0x92, 0x52, 0x73, 0x43, 0xc1, 0xc3, 0xb2,
	0xb0, 0x06, 0x1a, 0x96, 0x5f, 0x11, 0x56, 0x22, 0x9c, 0x83, 0xbd, 0x66, 0xde, 0x59, 0x4c, 0x93,
	0x30, 0x33, 0xbb, 0x76, 0x77, 0x3c, 0x38, 0x79, 0xe9, 0x56, 0x23, 0xbb, 0x8f, 0xfa, 0xd0, 0x91,
	0xc4, 0x3f, 0x2a, 0xda, 0x7f, 0x5a, 0x16, 0xd6, 0xa1, 0x2e, 0xbc, 0x5e, 0x09, 0xe1, 0xa1, 0x78,
	0x4c, 0xc3, 0x4b, 0x60, 0xb0, 0xe9, 0x57, 0x1a, 0x88, 0x49, 0x1c, 0x9a, 0x9b, 0xaa, 0xb3, 0xb7,
	0x65, 0x61, 0xed, 0xe9, 0x02, 0x2b, 0x09, 0xfd, 0xfe, 0x65, 0x59, 0xe0, 0xd9, 0xcd, 0x35, 0x71,
	0xbe, 0x9f, 0x39, 0x57, 0x93, 0x2f, 0xd7, 0xc7, 0xce, 0xbb, 0xfa, 0xfc, 0xe3, 0xf8, 0xe8, 0xf4,
	0xd5, 0xc3, 0x0b, 0xdc, 0xd7, 0xf8, 0x79, 0x08, 0x6f, 0xc0, 0x4e, 0x4a, 0x44, 0x70, 0x5b, 0x4f,
	0xb0, 0xa5, 0x26, 0x40, 0x7f, 0x99, 0xe0, 0x93, 0x44, 0xab, 0xee, 0x0f, 0xcb, 0xc2, 0xda, 0xd7,
	0x97, 0xb7, 0x2b, 0x20, 0x3c, 0x48, 0x1b, 0x4a, 0xae, 0x52, 0xe4, 0x0b, 0x6a, 0x6e, 0xaf, 0xaf,
	0x52, 0x7e, 0x45, 0x58, 0x89, 0xf0, 0x35, 0x00, 0x21, 0x11, 0x64, 0x22, 0xf7, 0x9a, 0x99, 0x3d,
	0xbb, 0x3b, 0x36, 0xfc, 0x27, 0x65, 0x61, 0xfd, 0xaf, 0xd1, 0x46, 0x43, 0xd8, 0x90, 0xc1, 0x67,
	0x79, 0x86, 0x47, 0xa0, 0x37, 0xa3, 0x22, 0xb8, 0xa5, 0xdc, 0xec, 0xab, 0xea, 0xb0, 0x2c, 0xac,
	0x5d, 0x9d, 0x52, 0x09, 0x08, 0xd7, 0x08, 0x3c, 0x01, 0x86, 0xf4, 0x07, 0x0d, 0x04, 0x0d, 0x4d,
	0xc3, 0xee, 0x8c, 0xfb, 0xfe, 0x41, 0xb3, 0xbe, 0x95, 0x84, 0x70, 0x83, 0x49, 0xe3, 0x84, 0x71,
	0x46, 0xa6, 0x09, 0x0d, 0x4d, 0xa0, 0x52, 0x5a, 0xc6, 0xa9, 0x15, 0x84, 0x57, 0x10, 0xb4, 0x41,
	0x97, 0xf1, 0xc8, 0x1c, 0xd8, 0x9d, 0xf1, 0x96, 0xbf, 0x5b, 0x16, 0x16, 0xa8, 0x5e, 0x87, 0x47,
	0x08, 0x4b, 0x09, 0x5a, 0xa0, 0x2b, 0xdf, 0x6f, 0x47, 0x35, 0xdc, 0x22, 0xe4, 0xcb, 0xe1, 0x8d,
	0xf3, 0x70, 0xf4, 0x00, 0x86, 0x6b, 0x3e, 0x91, 0xeb, 0xc9, 0x04, 0xa7, 0x24, 0x9d, 0xdc, 0xd1,
	0xda, 0xc1, 0xad, 0xf5, 0x34, 0x1a, 0xc2, 0x86, 0x0e, 0x2e, 0x68, 0x2e, 0xb3, 0x2a, 0x53, 0xc8,
	0xac, 0x8d, 0xf5, 0xac, 0x46, 0x43, 0xb8, 0x32, 0xd6, 0x05, 0xcd, 0x47, 0x39, 0x18, 0xb4, 0x1e,
	0xf9, 0x5f, 0x5e, 0xed, 0x7f, 0xb8, 0x3a, 0x8b, 0x98, 0x4b, 0x49, 0x96, 0xb3, 0x45, 0xe6, 0x26,
	0x2c, 0x20, 0x89, 0x17, 0xb0, 0xb9, 0xe0, 0x24, 0x10, 0x99, 0xfe, 0x59, 0x70, 0xba, 0x60, 0x4e,
	0xca, 0x42, 0x9a, 0x64, 0x5e, 0x05, 0x7a, 0x2a, 0xf4, 0x2a, 0xbf, 0x4e, 0xb7, 0x15, 0x77, 0xfa,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x77, 0xe5, 0x48, 0x27, 0x8c, 0x04, 0x00, 0x00,
}