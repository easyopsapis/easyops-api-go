// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: convert_sql.proto

package dbtask

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
//ConvertSQLToChangeset请求
type ConvertSQLToChangesetRequest struct {
	//
	//待转换的SQL文本body
	ConvertSQL           *ConvertSQLToChangesetRequest_ConvertSQL `protobuf:"bytes,1,opt,name=convertSQL,proto3" json:"convertSQL" form:"convertSQL"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ConvertSQLToChangesetRequest) Reset()         { *m = ConvertSQLToChangesetRequest{} }
func (m *ConvertSQLToChangesetRequest) String() string { return proto.CompactTextString(m) }
func (*ConvertSQLToChangesetRequest) ProtoMessage()    {}
func (*ConvertSQLToChangesetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aa208feffd5b9c8, []int{0}
}
func (m *ConvertSQLToChangesetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertSQLToChangesetRequest.Unmarshal(m, b)
}
func (m *ConvertSQLToChangesetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertSQLToChangesetRequest.Marshal(b, m, deterministic)
}
func (m *ConvertSQLToChangesetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertSQLToChangesetRequest.Merge(m, src)
}
func (m *ConvertSQLToChangesetRequest) XXX_Size() int {
	return xxx_messageInfo_ConvertSQLToChangesetRequest.Size(m)
}
func (m *ConvertSQLToChangesetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertSQLToChangesetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertSQLToChangesetRequest proto.InternalMessageInfo

func (m *ConvertSQLToChangesetRequest) GetConvertSQL() *ConvertSQLToChangesetRequest_ConvertSQL {
	if m != nil {
		return m.ConvertSQL
	}
	return nil
}

type ConvertSQLToChangesetRequest_ConvertSQL struct {
	//
	//变更SQL文本
	SqlText              string   `protobuf:"bytes,1,opt,name=sqlText,proto3" json:"sqlText" form:"sqlText"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertSQLToChangesetRequest_ConvertSQL) Reset() {
	*m = ConvertSQLToChangesetRequest_ConvertSQL{}
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) String() string { return proto.CompactTextString(m) }
func (*ConvertSQLToChangesetRequest_ConvertSQL) ProtoMessage()    {}
func (*ConvertSQLToChangesetRequest_ConvertSQL) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aa208feffd5b9c8, []int{0, 0}
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL.Unmarshal(m, b)
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL.Marshal(b, m, deterministic)
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL.Merge(m, src)
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) XXX_Size() int {
	return xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL.Size(m)
}
func (m *ConvertSQLToChangesetRequest_ConvertSQL) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertSQLToChangesetRequest_ConvertSQL proto.InternalMessageInfo

func (m *ConvertSQLToChangesetRequest_ConvertSQL) GetSqlText() string {
	if m != nil {
		return m.SqlText
	}
	return ""
}

//
//ConvertSQLToChangeset返回
type ConvertSQLToChangesetResponse struct {
	//
	//变更集名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//变更集作者
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author" form:"author"`
	//
	//变更集其它属性
	OtherAttr string `protobuf:"bytes,3,opt,name=otherAttr,proto3" json:"otherAttr" form:"otherAttr"`
	//
	//变更SQL
	UpdateSql string `protobuf:"bytes,4,opt,name=updateSql,proto3" json:"updateSql" form:"updateSql"`
	//
	//回退SQL
	RollbackSql          string   `protobuf:"bytes,5,opt,name=rollbackSql,proto3" json:"rollbackSql" form:"rollbackSql"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertSQLToChangesetResponse) Reset()         { *m = ConvertSQLToChangesetResponse{} }
func (m *ConvertSQLToChangesetResponse) String() string { return proto.CompactTextString(m) }
func (*ConvertSQLToChangesetResponse) ProtoMessage()    {}
func (*ConvertSQLToChangesetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aa208feffd5b9c8, []int{1}
}
func (m *ConvertSQLToChangesetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertSQLToChangesetResponse.Unmarshal(m, b)
}
func (m *ConvertSQLToChangesetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertSQLToChangesetResponse.Marshal(b, m, deterministic)
}
func (m *ConvertSQLToChangesetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertSQLToChangesetResponse.Merge(m, src)
}
func (m *ConvertSQLToChangesetResponse) XXX_Size() int {
	return xxx_messageInfo_ConvertSQLToChangesetResponse.Size(m)
}
func (m *ConvertSQLToChangesetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertSQLToChangesetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertSQLToChangesetResponse proto.InternalMessageInfo

func (m *ConvertSQLToChangesetResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConvertSQLToChangesetResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ConvertSQLToChangesetResponse) GetOtherAttr() string {
	if m != nil {
		return m.OtherAttr
	}
	return ""
}

func (m *ConvertSQLToChangesetResponse) GetUpdateSql() string {
	if m != nil {
		return m.UpdateSql
	}
	return ""
}

func (m *ConvertSQLToChangesetResponse) GetRollbackSql() string {
	if m != nil {
		return m.RollbackSql
	}
	return ""
}

//
//ConvertSQLToChangesetApi返回
type ConvertSQLToChangesetResponseWrapper struct {
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
	Data                 []*ConvertSQLToChangesetResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ConvertSQLToChangesetResponseWrapper) Reset()         { *m = ConvertSQLToChangesetResponseWrapper{} }
func (m *ConvertSQLToChangesetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ConvertSQLToChangesetResponseWrapper) ProtoMessage()    {}
func (*ConvertSQLToChangesetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aa208feffd5b9c8, []int{2}
}
func (m *ConvertSQLToChangesetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertSQLToChangesetResponseWrapper.Unmarshal(m, b)
}
func (m *ConvertSQLToChangesetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertSQLToChangesetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ConvertSQLToChangesetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertSQLToChangesetResponseWrapper.Merge(m, src)
}
func (m *ConvertSQLToChangesetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ConvertSQLToChangesetResponseWrapper.Size(m)
}
func (m *ConvertSQLToChangesetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertSQLToChangesetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertSQLToChangesetResponseWrapper proto.InternalMessageInfo

func (m *ConvertSQLToChangesetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConvertSQLToChangesetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ConvertSQLToChangesetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ConvertSQLToChangesetResponseWrapper) GetData() []*ConvertSQLToChangesetResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ConvertSQLToChangesetRequest)(nil), "dbtask.ConvertSQLToChangesetRequest")
	proto.RegisterType((*ConvertSQLToChangesetRequest_ConvertSQL)(nil), "dbtask.ConvertSQLToChangesetRequest.ConvertSQL")
	proto.RegisterType((*ConvertSQLToChangesetResponse)(nil), "dbtask.ConvertSQLToChangesetResponse")
	proto.RegisterType((*ConvertSQLToChangesetResponseWrapper)(nil), "dbtask.ConvertSQLToChangesetResponseWrapper")
}

func init() { proto.RegisterFile("convert_sql.proto", fileDescriptor_6aa208feffd5b9c8) }

var fileDescriptor_6aa208feffd5b9c8 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd5, 0xad, 0x2d, 0x9a, 0xcb, 0x9f, 0xcd, 0x02, 0x14, 0x55, 0xa0, 0x0c, 0x6f, 0xa0,
	0x1d, 0x96, 0x04, 0x3a, 0x09, 0x8d, 0xdd, 0xb6, 0x89, 0x0b, 0xe2, 0x00, 0xde, 0x24, 0x24, 0xa6,
	0x31, 0x39, 0x89, 0x9b, 0x56, 0x4d, 0xf3, 0xa6, 0x8e, 0xb3, 0x4d, 0x20, 0xbe, 0x17, 0x67, 0x3e,
	0x48, 0x90, 0xf8, 0x08, 0xb9, 0x71, 0x43, 0xb6, 0xbb, 0xd4, 0x42, 0xa2, 0x3b, 0xd5, 0x7e, 0x7f,
	0xcf, 0xf3, 0xf4, 0xcd, 0xa3, 0x04, 0x6d, 0x44, 0x90, 0x5d, 0x72, 0x21, 0x2f, 0x8a, 0x59, 0xea,
	0xe7, 0x02, 0x24, 0xe0, 0x6e, 0x1c, 0x4a, 0x56, 0x4c, 0xfa, 0x5e, 0x32, 0x96, 0xa3, 0x32, 0xf4,
	0x23, 0x98, 0x06, 0x09, 0x24, 0x10, 0x68, 0x1c, 0x96, 0x43, 0x7d, 0xd3, 0x17, 0x7d, 0x32, 0xb6,
	0xfe, 0x6b, 0x4b, 0x3e, 0xbd, 0x1a, 0xcb, 0x09, 0x5c, 0x05, 0x09, 0x78, 0x1a, 0x7a, 0x97, 0x2c,
	0x1d, 0xc7, 0x4c, 0x82, 0x28, 0x82, 0xe6, 0x68, 0x7c, 0xe4, 0x67, 0x0b, 0x3d, 0x39, 0x36, 0x4b,
	0x9c, 0x7c, 0x7c, 0x7f, 0x0a, 0xc7, 0x23, 0x96, 0x25, 0xbc, 0xe0, 0x92, 0xf2, 0x59, 0xc9, 0x0b,
	0x89, 0x87, 0x08, 0x45, 0x0d, 0x77, 0x5a, 0x9b, 0xad, 0x9d, 0xde, 0x20, 0xf0, 0xcd, 0x92, 0xfe,
	0x32, 0xa7, 0x05, 0x8f, 0x1e, 0xd5, 0x95, 0xbb, 0x31, 0x04, 0x31, 0x3d, 0x20, 0x8b, 0x30, 0x42,
	0xad, 0xe4, 0xfe, 0x01, 0x42, 0x0b, 0x03, 0xde, 0x45, 0x77, 0x8a, 0x59, 0x7a, 0xca, 0xaf, 0xa5,
	0xfe, 0xcb, 0xb5, 0x23, 0x5c, 0x57, 0xee, 0x7d, 0x93, 0x30, 0x07, 0x84, 0xde, 0x48, 0xc8, 0x8f,
	0x15, 0xf4, 0xf4, 0x3f, 0xab, 0x14, 0x39, 0x64, 0x05, 0xc7, 0x5b, 0xa8, 0x9d, 0xb1, 0x29, 0x9f,
	0x87, 0x3d, 0xa8, 0x2b, 0xb7, 0x67, 0xc2, 0xd4, 0x94, 0x50, 0x0d, 0xf1, 0x07, 0xd4, 0x65, 0xa5,
	0x1c, 0x81, 0x70, 0x56, 0xb4, 0x6c, 0xbf, 0xae, 0xdc, 0x7b, 0x46, 0x66, 0xe6, 0xe4, 0xf7, 0x2f,
	0x77, 0x0b, 0x3d, 0xfb, 0x72, 0xc6, 0xbc, 0xaf, 0x87, 0xde, 0xe7, 0x97, 0xde, 0x9b, 0xf3, 0x33,
	0xbf, 0x39, 0x5f, 0x78, 0xe7, 0xdf, 0x06, 0xbb, 0x7b, 0xaf, 0xbe, 0x6f, 0xd3, 0x79, 0x0e, 0x1e,
	0xa0, 0x35, 0x90, 0x23, 0x2e, 0x0e, 0xa5, 0x14, 0xce, 0xaa, 0x0e, 0x7d, 0x58, 0x57, 0xee, 0xba,
	0x09, 0x6d, 0x10, 0xa1, 0x0b, 0x99, 0xf2, 0x94, 0x79, 0xcc, 0x24, 0x3f, 0x99, 0xa5, 0x4e, 0xfb,
	0x5f, 0x4f, 0x83, 0x08, 0x5d, 0xc8, 0xf0, 0x3e, 0xea, 0x09, 0x48, 0xd3, 0x90, 0x45, 0x13, 0xe5,
	0xea, 0x68, 0xd7, 0xe3, 0xba, 0x72, 0xb1, 0x71, 0x59, 0x90, 0x50, 0x5b, 0x4a, 0xfe, 0xb4, 0xd0,
	0xf6, 0xd2, 0xea, 0x3e, 0x09, 0x96, 0xe7, 0x5c, 0xa8, 0x06, 0x23, 0x88, 0x4d, 0x83, 0x1d, 0xbb,
	0x41, 0x35, 0x25, 0x54, 0x43, 0xb5, 0x87, 0xfa, 0x7d, 0x7b, 0x9d, 0xa7, 0x6c, 0x9c, 0xcd, 0x6b,
	0xb4, 0xf6, 0xb0, 0x20, 0xa1, 0xb6, 0x14, 0xbf, 0x40, 0x1d, 0x2e, 0x04, 0xdc, 0xb4, 0xb4, 0x5e,
	0x57, 0xee, 0x5d, 0xe3, 0xd1, 0x63, 0x42, 0x0d, 0xc6, 0xef, 0x50, 0x3b, 0x66, 0x92, 0x39, 0xed,
	0xcd, 0xd5, 0x9d, 0xde, 0xe0, 0xf9, 0x2d, 0x2f, 0xa2, 0x79, 0x04, 0x7b, 0x5b, 0x65, 0x26, 0x54,
	0x67, 0x84, 0x5d, 0xfd, 0x09, 0xec, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x35, 0xfc, 0xf8,
	0x86, 0x03, 0x00, 0x00,
}