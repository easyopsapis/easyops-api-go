// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data_name_list_handler.proto

package data_name

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
//DataNameListHandler请求
type DataNameListHandlerRequest struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//数据表名称
	DataName string `protobuf:"bytes,3,opt,name=data_name,json=dataName,proto3" json:"data_name" form:"data_name"`
	//
	//数据表名称列表
	DataName_In          string   `protobuf:"bytes,4,opt,name=data_name__in,json=dataNameIn,proto3" json:"data_name__in" form:"data_name__in"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataNameListHandlerRequest) Reset()         { *m = DataNameListHandlerRequest{} }
func (m *DataNameListHandlerRequest) String() string { return proto.CompactTextString(m) }
func (*DataNameListHandlerRequest) ProtoMessage()    {}
func (*DataNameListHandlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{0}
}
func (m *DataNameListHandlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerRequest.Unmarshal(m, b)
}
func (m *DataNameListHandlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerRequest.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerRequest.Merge(m, src)
}
func (m *DataNameListHandlerRequest) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerRequest.Size(m)
}
func (m *DataNameListHandlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerRequest proto.InternalMessageInfo

func (m *DataNameListHandlerRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DataNameListHandlerRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *DataNameListHandlerRequest) GetDataName() string {
	if m != nil {
		return m.DataName
	}
	return ""
}

func (m *DataNameListHandlerRequest) GetDataName_In() string {
	if m != nil {
		return m.DataName_In
	}
	return ""
}

//
//DataNameListHandler返回
type DataNameListHandlerResponse struct {
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
	Data                 []*DataNameListHandlerResponse_Data `protobuf:"bytes,6,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *DataNameListHandlerResponse) Reset()         { *m = DataNameListHandlerResponse{} }
func (m *DataNameListHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*DataNameListHandlerResponse) ProtoMessage()    {}
func (*DataNameListHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{1}
}
func (m *DataNameListHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerResponse.Unmarshal(m, b)
}
func (m *DataNameListHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerResponse.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerResponse.Merge(m, src)
}
func (m *DataNameListHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerResponse.Size(m)
}
func (m *DataNameListHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerResponse proto.InternalMessageInfo

func (m *DataNameListHandlerResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DataNameListHandlerResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *DataNameListHandlerResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DataNameListHandlerResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *DataNameListHandlerResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DataNameListHandlerResponse) GetData() []*DataNameListHandlerResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataNameListHandlerResponse_Data struct {
	//
	//org
	Org int32 `protobuf:"varint,1,opt,name=org,proto3" json:"org" form:"org"`
	//
	//创建用户
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//修改用户
	Modifier string `protobuf:"bytes,3,opt,name=modifier,proto3" json:"modifier" form:"modifier"`
	//
	//数据表描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description"`
	//
	//数据保存策略
	RetentionPolicy string `protobuf:"bytes,5,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy" form:"retention_policy"`
	//
	//数据表类型
	DataType string `protobuf:"bytes,6,opt,name=data_type,json=dataType,proto3" json:"data_type" form:"data_type"`
	//
	//数据表名称
	DataName string `protobuf:"bytes,7,opt,name=data_name,json=dataName,proto3" json:"data_name" form:"data_name"`
	//
	//是否自定义
	Custom bool `protobuf:"varint,8,opt,name=custom,proto3" json:"custom" form:"custom"`
	//
	//版本号
	Version float32 `protobuf:"fixed32,9,opt,name=version,proto3" json:"version" form:"version"`
	//
	//数据表id
	XId string `protobuf:"bytes,10,opt,name=_id,json=Id,proto3" json:"_id" form:"_id"`
	//
	//数据表列表
	Formats              []*DataNameListHandlerResponse_Data_Formats `protobuf:"bytes,11,rep,name=formats,proto3" json:"formats" form:"formats"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *DataNameListHandlerResponse_Data) Reset()         { *m = DataNameListHandlerResponse_Data{} }
func (m *DataNameListHandlerResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DataNameListHandlerResponse_Data) ProtoMessage()    {}
func (*DataNameListHandlerResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{1, 0}
}
func (m *DataNameListHandlerResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerResponse_Data.Unmarshal(m, b)
}
func (m *DataNameListHandlerResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerResponse_Data.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerResponse_Data.Merge(m, src)
}
func (m *DataNameListHandlerResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerResponse_Data.Size(m)
}
func (m *DataNameListHandlerResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerResponse_Data proto.InternalMessageInfo

func (m *DataNameListHandlerResponse_Data) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *DataNameListHandlerResponse_Data) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetRetentionPolicy() string {
	if m != nil {
		return m.RetentionPolicy
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetDataName() string {
	if m != nil {
		return m.DataName
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetCustom() bool {
	if m != nil {
		return m.Custom
	}
	return false
}

func (m *DataNameListHandlerResponse_Data) GetVersion() float32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *DataNameListHandlerResponse_Data) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data) GetFormats() []*DataNameListHandlerResponse_Data_Formats {
	if m != nil {
		return m.Formats
	}
	return nil
}

type DataNameListHandlerResponse_Data_Formats struct {
	//
	//配置id
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id" form:"config_id"`
	//
	//字段定义
	Fields []*DataNameListHandlerResponse_Data_Formats_Fields `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields" form:"fields"`
	//
	//数据表名
	Table                string   `protobuf:"bytes,3,opt,name=table,proto3" json:"table" form:"table"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataNameListHandlerResponse_Data_Formats) Reset() {
	*m = DataNameListHandlerResponse_Data_Formats{}
}
func (m *DataNameListHandlerResponse_Data_Formats) String() string { return proto.CompactTextString(m) }
func (*DataNameListHandlerResponse_Data_Formats) ProtoMessage()    {}
func (*DataNameListHandlerResponse_Data_Formats) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{1, 0, 0}
}
func (m *DataNameListHandlerResponse_Data_Formats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats.Unmarshal(m, b)
}
func (m *DataNameListHandlerResponse_Data_Formats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerResponse_Data_Formats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerResponse_Data_Formats.Merge(m, src)
}
func (m *DataNameListHandlerResponse_Data_Formats) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats.Size(m)
}
func (m *DataNameListHandlerResponse_Data_Formats) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerResponse_Data_Formats.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerResponse_Data_Formats proto.InternalMessageInfo

func (m *DataNameListHandlerResponse_Data_Formats) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data_Formats) GetFields() []*DataNameListHandlerResponse_Data_Formats_Fields {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *DataNameListHandlerResponse_Data_Formats) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

type DataNameListHandlerResponse_Data_Formats_Fields struct {
	//
	//字段类型
	FieldType string `protobuf:"bytes,1,opt,name=field_type,json=fieldType,proto3" json:"field_type" form:"field_type"`
	//
	//字段描述
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" form:"description"`
	//
	//字段名称
	FieldName string `protobuf:"bytes,3,opt,name=field_name,json=fieldName,proto3" json:"field_name" form:"field_name"`
	//
	//字段值类型
	FieldValueType string `protobuf:"bytes,4,opt,name=field_value_type,json=fieldValueType,proto3" json:"field_value_type" form:"field_value_type"`
	//
	//单位
	Unit                 string   `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit" form:"unit"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataNameListHandlerResponse_Data_Formats_Fields) Reset() {
	*m = DataNameListHandlerResponse_Data_Formats_Fields{}
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) String() string {
	return proto.CompactTextString(m)
}
func (*DataNameListHandlerResponse_Data_Formats_Fields) ProtoMessage() {}
func (*DataNameListHandlerResponse_Data_Formats_Fields) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{1, 0, 0, 0}
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields.Unmarshal(m, b)
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields.Merge(m, src)
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields.Size(m)
}
func (m *DataNameListHandlerResponse_Data_Formats_Fields) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerResponse_Data_Formats_Fields proto.InternalMessageInfo

func (m *DataNameListHandlerResponse_Data_Formats_Fields) GetFieldType() string {
	if m != nil {
		return m.FieldType
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data_Formats_Fields) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data_Formats_Fields) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data_Formats_Fields) GetFieldValueType() string {
	if m != nil {
		return m.FieldValueType
	}
	return ""
}

func (m *DataNameListHandlerResponse_Data_Formats_Fields) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

//
//DataNameListHandlerApi返回
type DataNameListHandlerResponseWrapper struct {
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
	Data                 *DataNameListHandlerResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DataNameListHandlerResponseWrapper) Reset()         { *m = DataNameListHandlerResponseWrapper{} }
func (m *DataNameListHandlerResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DataNameListHandlerResponseWrapper) ProtoMessage()    {}
func (*DataNameListHandlerResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f9dfe0aee89d34, []int{2}
}
func (m *DataNameListHandlerResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNameListHandlerResponseWrapper.Unmarshal(m, b)
}
func (m *DataNameListHandlerResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNameListHandlerResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DataNameListHandlerResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNameListHandlerResponseWrapper.Merge(m, src)
}
func (m *DataNameListHandlerResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DataNameListHandlerResponseWrapper.Size(m)
}
func (m *DataNameListHandlerResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNameListHandlerResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DataNameListHandlerResponseWrapper proto.InternalMessageInfo

func (m *DataNameListHandlerResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DataNameListHandlerResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DataNameListHandlerResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DataNameListHandlerResponseWrapper) GetData() *DataNameListHandlerResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DataNameListHandlerRequest)(nil), "data_name.DataNameListHandlerRequest")
	proto.RegisterType((*DataNameListHandlerResponse)(nil), "data_name.DataNameListHandlerResponse")
	proto.RegisterType((*DataNameListHandlerResponse_Data)(nil), "data_name.DataNameListHandlerResponse.Data")
	proto.RegisterType((*DataNameListHandlerResponse_Data_Formats)(nil), "data_name.DataNameListHandlerResponse.Data.Formats")
	proto.RegisterType((*DataNameListHandlerResponse_Data_Formats_Fields)(nil), "data_name.DataNameListHandlerResponse.Data.Formats.Fields")
	proto.RegisterType((*DataNameListHandlerResponseWrapper)(nil), "data_name.DataNameListHandlerResponseWrapper")
}

func init() { proto.RegisterFile("data_name_list_handler.proto", fileDescriptor_95f9dfe0aee89d34) }

var fileDescriptor_95f9dfe0aee89d34 = []byte{
	// 834 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x8f, 0xdb, 0x44,
	0x14, 0xad, 0xf3, 0x9d, 0x1b, 0xba, 0x9b, 0x0e, 0x05, 0xac, 0x14, 0xe4, 0x68, 0x2a, 0x55, 0xa9,
	0x60, 0x13, 0xd1, 0x22, 0x54, 0x2d, 0x6f, 0x11, 0x5d, 0xb1, 0x02, 0xa1, 0x6a, 0x40, 0xf0, 0x68,
	0xcd, 0xc6, 0x13, 0x77, 0x84, 0xed, 0x31, 0xe3, 0xc9, 0x96, 0xed, 0x4f, 0xe1, 0x8f, 0xf0, 0x6f,
	0x8c, 0xc4, 0x13, 0xbc, 0x9a, 0x17, 0xde, 0x40, 0xf3, 0x11, 0xc7, 0x72, 0x97, 0xb2, 0xc0, 0x53,
	0x3c, 0xf7, 0x9c, 0x7b, 0xe7, 0xde, 0x33, 0x67, 0x26, 0xf0, 0x6e, 0x44, 0x15, 0x0d, 0x33, 0x9a,
	0xb2, 0x30, 0xe1, 0x85, 0x0a, 0x9f, 0xd3, 0x2c, 0x4a, 0x98, 0x5c, 0xe6, 0x52, 0x28, 0x81, 0xc6,
	0x35, 0x3a, 0x3b, 0x89, 0xb9, 0x7a, 0xbe, 0xbb, 0x58, 0x6e, 0x44, 0xba, 0x8a, 0x45, 0x2c, 0x56,
	0x86, 0x71, 0xb1, 0xdb, 0x9a, 0x95, 0x59, 0x98, 0x2f, 0x9b, 0x39, 0xfb, 0xb8, 0x41, 0x4f, 0x5f,
	0x70, 0xf5, 0x9d, 0x78, 0xb1, 0x8a, 0xc5, 0x89, 0x01, 0x4f, 0x2e, 0x69, 0xc2, 0x23, 0xaa, 0x84,
	0x2c, 0x56, 0xf5, 0xa7, 0xcd, 0xc3, 0xbf, 0x7b, 0x30, 0xfb, 0x94, 0x2a, 0xfa, 0x25, 0x4d, 0xd9,
	0x17, 0xbc, 0x50, 0x9f, 0xd9, 0x7e, 0x08, 0xfb, 0x7e, 0xc7, 0x0a, 0x85, 0x1e, 0x42, 0x2f, 0xa7,
	0x31, 0xf3, 0xbd, 0xb9, 0xb7, 0xe8, 0xaf, 0xdf, 0xaa, 0xca, 0x60, 0xb2, 0x15, 0x32, 0x3d, 0xc5,
	0x3a, 0x8a, 0x7f, 0xf9, 0x39, 0xe8, 0x4c, 0x6f, 0x11, 0x43, 0x41, 0xa7, 0x30, 0xd6, 0xbf, 0x61,
	0xc1, 0x5f, 0x32, 0xbf, 0x63, 0xf8, 0xef, 0x55, 0x65, 0x30, 0x3d, 0xf0, 0x0d, 0xa4, 0x93, 0xfa,
	0xd3, 0x5b, 0xfe, 0x6f, 0x43, 0x32, 0xd2, 0xc1, 0xaf, 0xf8, 0x4b, 0x86, 0x3e, 0x84, 0xc3, 0xe4,
	0x7e, 0x77, 0xee, 0x2d, 0xc6, 0xeb, 0xbb, 0x87, 0xdc, 0x1a, 0xc2, 0x64, 0x14, 0xb9, 0x5e, 0xd1,
	0x27, 0x70, 0xfb, 0x20, 0x65, 0xc8, 0x33, 0xbf, 0x67, 0xd2, 0xfc, 0xaa, 0x0c, 0xee, 0xb6, 0xd2,
	0x34, 0x8c, 0x09, 0xec, 0x53, 0xcf, 0x33, 0xfc, 0x27, 0xc0, 0xbd, 0x6b, 0xa7, 0x2e, 0x72, 0x91,
	0x15, 0x0c, 0xdd, 0x87, 0xde, 0x46, 0x44, 0xfb, 0xb1, 0x8f, 0x0f, 0x63, 0xeb, 0x28, 0x26, 0x06,
	0xfc, 0x5f, 0x03, 0xcf, 0xa1, 0x9b, 0x16, 0xb1, 0x1b, 0xf5, 0xa8, 0x2a, 0x03, 0xb0, 0x59, 0x69,
	0x11, 0x63, 0xa2, 0x21, 0xf4, 0x00, 0xfa, 0x4a, 0x28, 0x9a, 0x98, 0xb9, 0xfa, 0xeb, 0x69, 0x55,
	0x06, 0x6f, 0x58, 0x8e, 0x09, 0x63, 0x62, 0xe1, 0xfa, 0x84, 0xfa, 0xff, 0x7c, 0x42, 0xcf, 0xa0,
	0xa7, 0x35, 0xf0, 0x07, 0xf3, 0xee, 0x62, 0xf2, 0xe8, 0xfd, 0x65, 0x2d, 0xd0, 0xf2, 0x35, 0x5a,
	0x18, 0xac, 0x29, 0x81, 0xce, 0xc2, 0xc4, 0x54, 0x9a, 0xfd, 0x34, 0x82, 0x9e, 0xc6, 0xf5, 0x3c,
	0x42, 0xc6, 0x4e, 0xaf, 0xc6, 0x3c, 0x42, 0xea, 0x79, 0x84, 0x8c, 0xd1, 0x07, 0x30, 0xdc, 0x48,
	0xa6, 0x9d, 0x67, 0xb4, 0x1a, 0xaf, 0x51, 0x55, 0x06, 0x47, 0x4e, 0x55, 0x0b, 0x60, 0xb2, 0xa7,
	0xa0, 0x15, 0x8c, 0x52, 0x11, 0xf1, 0x2d, 0x67, 0xd2, 0x89, 0xf4, 0x66, 0x55, 0x06, 0xc7, 0x4e,
	0x24, 0x87, 0x60, 0x52, 0x93, 0xd0, 0x13, 0x98, 0x44, 0xac, 0xd8, 0x48, 0x9e, 0x2b, 0x2e, 0xf6,
	0x66, 0x78, 0xbb, 0x2a, 0x03, 0xe4, 0xba, 0x3e, 0x80, 0x98, 0x34, 0xa9, 0xe8, 0x0c, 0xa6, 0x92,
	0x29, 0x96, 0xe9, 0x45, 0x98, 0x8b, 0x84, 0x6f, 0xae, 0x8c, 0x98, 0xe3, 0xf5, 0xbd, 0xaa, 0x0c,
	0xde, 0xb1, 0xe9, 0x6d, 0x06, 0x26, 0xc7, 0x75, 0xe8, 0x99, 0x89, 0xd4, 0x1e, 0x56, 0x57, 0x39,
	0xf3, 0x07, 0xd7, 0x7a, 0x58, 0x43, 0xce, 0xc3, 0x5f, 0x5f, 0xe5, 0x2d, 0xdb, 0x0f, 0x6f, 0x64,
	0xfb, 0x87, 0x30, 0xd8, 0xec, 0x0a, 0x25, 0x52, 0x7f, 0x34, 0xf7, 0x16, 0xa3, 0xf5, 0x9d, 0xaa,
	0x0c, 0x6e, 0x3b, 0x15, 0x4d, 0x1c, 0x13, 0x47, 0xd0, 0x8a, 0x5f, 0x32, 0x59, 0x68, 0x39, 0xc6,
	0x73, 0x6f, 0xd1, 0x69, 0x2a, 0xee, 0x00, 0x4c, 0xf6, 0x14, 0x14, 0x40, 0x37, 0xe4, 0x91, 0x0f,
	0x6d, 0x47, 0x86, 0x3c, 0xc2, 0xa4, 0x73, 0x1e, 0x21, 0x0a, 0x43, 0x1d, 0xa1, 0xaa, 0xf0, 0x27,
	0xc6, 0x40, 0x8f, 0xff, 0x85, 0x81, 0x96, 0x67, 0x36, 0xb5, 0xd9, 0x83, 0xab, 0x86, 0xc9, 0xbe,
	0xee, 0xec, 0xd7, 0x2e, 0x0c, 0x1d, 0x51, 0x6b, 0xb3, 0x11, 0xd9, 0x96, 0xc7, 0xba, 0x2b, 0xaf,
	0xad, 0x4d, 0x0d, 0x61, 0x32, 0xb2, 0xdf, 0xe7, 0x11, 0x62, 0x30, 0xd8, 0x72, 0x96, 0x44, 0x85,
	0xdf, 0x31, 0x0d, 0x9e, 0xfe, 0x87, 0x06, 0x97, 0x67, 0xa6, 0x42, 0x53, 0x57, 0x5b, 0x13, 0x13,
	0x57, 0xdc, 0xdc, 0x4c, 0x7a, 0x91, 0xec, 0x1f, 0xaa, 0xe6, 0xcd, 0xd4, 0x61, 0x7d, 0x33, 0xf5,
	0xef, 0xec, 0xc7, 0x0e, 0x0c, 0x6c, 0x35, 0xf4, 0x11, 0x80, 0x49, 0xb6, 0xe6, 0xb0, 0xd3, 0xe8,
	0xab, 0x7a, 0xa7, 0xb1, 0x83, 0x73, 0xc7, 0xd8, 0x2c, 0x8c, 0x3d, 0x5a, 0x9e, 0xee, 0xdc, 0xdc,
	0xd3, 0xf5, 0x7e, 0x8d, 0x07, 0xf5, 0x95, 0xfd, 0xac, 0xb5, 0xec, 0x7e, 0xc6, 0x5b, 0x4f, 0x61,
	0x6a, 0x91, 0x4b, 0x9a, 0xec, 0x98, 0xed, 0xb5, 0xd7, 0xbe, 0x09, 0x6d, 0x06, 0x26, 0x47, 0x26,
	0xf4, 0x8d, 0x8e, 0x98, 0xb6, 0xef, 0x43, 0x6f, 0x97, 0x71, 0xe5, 0x2e, 0x51, 0xe3, 0xe5, 0xd0,
	0x51, 0x4c, 0x0c, 0x88, 0xff, 0xf0, 0x00, 0xbf, 0xe6, 0x4c, 0xbe, 0x95, 0x34, 0xcf, 0x99, 0xbc,
	0xd9, 0x43, 0xfc, 0x04, 0x26, 0xfa, 0xf7, 0xe9, 0x0f, 0x79, 0x42, 0xf9, 0x35, 0x3a, 0x35, 0x40,
	0x4c, 0x9a, 0x54, 0x7d, 0x94, 0x4c, 0x4a, 0x21, 0x5f, 0x3d, 0x4a, 0x13, 0xc6, 0xc4, 0xc2, 0xe8,
	0x73, 0xf7, 0x72, 0x6a, 0x35, 0x26, 0x8f, 0x1e, 0xdc, 0xcc, 0x57, 0x7f, 0xf3, 0x68, 0x5e, 0x0c,
	0xcc, 0x3f, 0xef, 0xe3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x16, 0xca, 0xa2, 0x0b, 0x08,
	0x00, 0x00,
}
