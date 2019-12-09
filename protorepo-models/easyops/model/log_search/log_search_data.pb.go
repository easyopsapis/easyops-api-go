// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log_search_data.proto

package log_search

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
//日志搜索返回数据
type LogSearchData struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//搜索主机的ip
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//错误提示
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg" form:"msg"`
	//
	//搜索结果
	Data                 *LogSearchData_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogSearchData) Reset()         { *m = LogSearchData{} }
func (m *LogSearchData) String() string { return proto.CompactTextString(m) }
func (*LogSearchData) ProtoMessage()    {}
func (*LogSearchData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23113a33471c2cc, []int{0}
}
func (m *LogSearchData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchData.Unmarshal(m, b)
}
func (m *LogSearchData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchData.Marshal(b, m, deterministic)
}
func (m *LogSearchData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchData.Merge(m, src)
}
func (m *LogSearchData) XXX_Size() int {
	return xxx_messageInfo_LogSearchData.Size(m)
}
func (m *LogSearchData) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchData.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchData proto.InternalMessageInfo

func (m *LogSearchData) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogSearchData) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *LogSearchData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LogSearchData) GetData() *LogSearchData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type LogSearchData_Data struct {
	//
	//最大查询行数
	MaxQueryLines int32 `protobuf:"varint,1,opt,name=max_query_lines,json=maxQueryLines,proto3" json:"max_query_lines" form:"max_query_lines"`
	//
	//最大返回行数
	MaxReturnLines int32 `protobuf:"varint,2,opt,name=max_return_lines,json=maxReturnLines,proto3" json:"max_return_lines" form:"max_return_lines"`
	//
	//日志返回（可同时搜索多个日志路径，所以这里是个列表）
	Content              []*LogSearchData_Data_Content `protobuf:"bytes,3,rep,name=content,proto3" json:"content" form:"content"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LogSearchData_Data) Reset()         { *m = LogSearchData_Data{} }
func (m *LogSearchData_Data) String() string { return proto.CompactTextString(m) }
func (*LogSearchData_Data) ProtoMessage()    {}
func (*LogSearchData_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23113a33471c2cc, []int{0, 0}
}
func (m *LogSearchData_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchData_Data.Unmarshal(m, b)
}
func (m *LogSearchData_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchData_Data.Marshal(b, m, deterministic)
}
func (m *LogSearchData_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchData_Data.Merge(m, src)
}
func (m *LogSearchData_Data) XXX_Size() int {
	return xxx_messageInfo_LogSearchData_Data.Size(m)
}
func (m *LogSearchData_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchData_Data.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchData_Data proto.InternalMessageInfo

func (m *LogSearchData_Data) GetMaxQueryLines() int32 {
	if m != nil {
		return m.MaxQueryLines
	}
	return 0
}

func (m *LogSearchData_Data) GetMaxReturnLines() int32 {
	if m != nil {
		return m.MaxReturnLines
	}
	return 0
}

func (m *LogSearchData_Data) GetContent() []*LogSearchData_Data_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

type LogSearchData_Data_Content struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//日志路径
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path" form:"file_path"`
	//
	//日志内容
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text" form:"text"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSearchData_Data_Content) Reset()         { *m = LogSearchData_Data_Content{} }
func (m *LogSearchData_Data_Content) String() string { return proto.CompactTextString(m) }
func (*LogSearchData_Data_Content) ProtoMessage()    {}
func (*LogSearchData_Data_Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23113a33471c2cc, []int{0, 0, 0}
}
func (m *LogSearchData_Data_Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchData_Data_Content.Unmarshal(m, b)
}
func (m *LogSearchData_Data_Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchData_Data_Content.Marshal(b, m, deterministic)
}
func (m *LogSearchData_Data_Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchData_Data_Content.Merge(m, src)
}
func (m *LogSearchData_Data_Content) XXX_Size() int {
	return xxx_messageInfo_LogSearchData_Data_Content.Size(m)
}
func (m *LogSearchData_Data_Content) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchData_Data_Content.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchData_Data_Content proto.InternalMessageInfo

func (m *LogSearchData_Data_Content) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogSearchData_Data_Content) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *LogSearchData_Data_Content) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*LogSearchData)(nil), "log_search.LogSearchData")
	proto.RegisterType((*LogSearchData_Data)(nil), "log_search.LogSearchData.Data")
	proto.RegisterType((*LogSearchData_Data_Content)(nil), "log_search.LogSearchData.Data.Content")
}

func init() { proto.RegisterFile("log_search_data.proto", fileDescriptor_c23113a33471c2cc) }

var fileDescriptor_c23113a33471c2cc = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xc7, 0x95, 0xcb, 0xf7, 0x95, 0x4c, 0xd5, 0x8b, 0x22, 0x2e, 0x51, 0x90, 0x70, 0x64, 0x2e,
	0x1a, 0x97, 0xda, 0x8e, 0x9d, 0xa4, 0x55, 0xb3, 0xa9, 0x48, 0x4b, 0x57, 0x5d, 0x80, 0x61, 0x45,
	0x12, 0xa2, 0x69, 0xe2, 0x38, 0x16, 0x76, 0xc6, 0xd8, 0x13, 0x9a, 0x2a, 0xf6, 0x96, 0x2d, 0x12,
	0xcf, 0xc0, 0x2b, 0xf0, 0x0a, 0x3c, 0x86, 0x91, 0x78, 0x02, 0xe4, 0x27, 0x40, 0x33, 0x76, 0x9b,
	0x54, 0xad, 0x4b, 0x17, 0xa5, 0x9b, 0x68, 0xce, 0x39, 0xff, 0x73, 0xe6, 0x37, 0xb1, 0xf4, 0x3f,
	0xe0, 0x9e, 0x85, 0x8d, 0x9e, 0xa7, 0x23, 0xb7, 0x3f, 0xea, 0x0d, 0x10, 0x41, 0x92, 0xe3, 0x62,
	0x82, 0x8b, 0x60, 0x9e, 0x2e, 0x8b, 0x86, 0x49, 0x46, 0x93, 0x23, 0xa9, 0x8f, 0x6d, 0xd9, 0xc0,
	0x06, 0x96, 0x99, 0xe4, 0x68, 0x32, 0x64, 0x11, 0x0b, 0xd8, 0x29, 0x6e, 0x2d, 0x6f, 0x2d, 0xc8,
	0xed, 0x63, 0x93, 0x7c, 0xc0, 0xc7, 0xb2, 0x81, 0x45, 0x56, 0x14, 0x3f, 0x21, 0xcb, 0x1c, 0x20,
	0x82, 0x5d, 0x4f, 0x3e, 0x3b, 0xc6, 0x7d, 0xfc, 0x8f, 0x15, 0xb0, 0x72, 0x88, 0x8d, 0x37, 0xec,
	0xd2, 0x7d, 0x44, 0x50, 0xf1, 0x31, 0xc8, 0xf7, 0xf1, 0x40, 0x2f, 0x65, 0x2a, 0x19, 0xf8, 0x5f,
	0x6b, 0x2d, 0x0a, 0xb9, 0xe5, 0x21, 0x76, 0xed, 0x26, 0x4f, 0xb3, 0xbc, 0xc6, 0x8a, 0xc5, 0xef,
	0x05, 0x90, 0x35, 0x9d, 0x52, 0xb6, 0x92, 0x81, 0x85, 0xd6, 0xb7, 0x42, 0x14, 0x72, 0x85, 0x58,
	0x64, 0x3a, 0xfc, 0xaf, 0x9f, 0xdc, 0xd7, 0x02, 0xf8, 0x52, 0x78, 0x0f, 0xa1, 0x0a, 0x1b, 0xed,
	0xaa, 0xd8, 0xe8, 0xce, 0x94, 0xc0, 0x6f, 0x57, 0xc5, 0x7a, 0xb7, 0x33, 0x98, 0x29, 0x81, 0x40,
	0xcf, 0x4a, 0x77, 0x97, 0x06, 0x9b, 0x6a, 0x20, 0xc0, 0x8e, 0x74, 0x4d, 0xa5, 0x30, 0xab, 0x05,
	0x82, 0xdf, 0xf1, 0x36, 0x20, 0x84, 0xed, 0xaa, 0xb8, 0xf3, 0x42, 0x3c, 0x40, 0xe2, 0xb0, 0x3b,
	0x53, 0x36, 0xeb, 0x41, 0x53, 0x98, 0x6d, 0x07, 0x17, 0xb2, 0x7e, 0x53, 0x10, 0xfc, 0x4b, 0xc5,
	0x5b, 0x01, 0x6c, 0x5e, 0x50, 0x43, 0xa8, 0xc6, 0x1c, 0xbe, 0x9a, 0x50, 0xf8, 0x4a, 0x67, 0xd0,
	0x19, 0xf8, 0x6d, 0x45, 0xdc, 0xa1, 0x1c, 0x31, 0xec, 0x5f, 0x34, 0x31, 0x66, 0xea, 0xcd, 0x8d,
	0x00, 0xc2, 0x8b, 0x77, 0x0b, 0xf1, 0x13, 0xfd, 0xe6, 0xad, 0x30, 0xd4, 0x53, 0x19, 0x68, 0xdb,
	0x65, 0xa5, 0xdd, 0x9b, 0x04, 0xbb, 0x82, 0xac, 0x96, 0x4a, 0x56, 0x4f, 0x21, 0x9b, 0x55, 0x37,
	0xd5, 0xe0, 0x96, 0xe8, 0xd4, 0x54, 0xba, 0x46, 0x3a, 0x5d, 0xed, 0xb6, 0xe8, 0x94, 0x54, 0xba,
	0xad, 0x74, 0xba, 0xfa, 0xbf, 0xa0, 0x6b, 0xa6, 0x81, 0x6c, 0xa7, 0x83, 0x34, 0x6e, 0x1e, 0x44,
	0x80, 0x4f, 0xa5, 0xe7, 0xc2, 0x6e, 0xc7, 0xdb, 0x78, 0xa2, 0x65, 0x4d, 0xa7, 0x58, 0x01, 0x39,
	0xdb, 0x33, 0x4a, 0x39, 0xe6, 0x5b, 0xab, 0x51, 0xc8, 0x81, 0xd8, 0xb6, 0x6c, 0xcf, 0xe0, 0x35,
	0x5a, 0x2a, 0xee, 0x81, 0x3c, 0x75, 0xe4, 0x52, 0xbe, 0x92, 0x81, 0xcb, 0xea, 0x23, 0x69, 0x6e,
	0xc9, 0xd2, 0x39, 0x9f, 0x94, 0xe8, 0xcf, 0xa2, 0x3d, 0xd2, 0x2e, 0x5e, 0x63, 0xcd, 0xe5, 0xdf,
	0x59, 0x90, 0x67, 0x66, 0xda, 0x02, 0x6b, 0x36, 0x9a, 0xf6, 0x3e, 0x4e, 0x74, 0xf7, 0xa4, 0x67,
	0x99, 0x63, 0xdd, 0x4b, 0x7c, 0xb5, 0x1c, 0x85, 0xdc, 0xfd, 0xe4, 0xee, 0xf3, 0x02, 0x5e, 0x5b,
	0xb1, 0xd1, 0xf4, 0x35, 0x4d, 0x1c, 0xd2, 0xb8, 0xf8, 0x12, 0xac, 0x53, 0x89, 0xab, 0x93, 0x89,
	0x3b, 0x4e, 0x86, 0x64, 0xd9, 0x90, 0x87, 0x51, 0xc8, 0x3d, 0x98, 0x0f, 0x59, 0x54, 0xf0, 0xda,
	0xaa, 0x8d, 0xa6, 0x1a, 0xcb, 0xc4, 0x63, 0xde, 0x82, 0xa5, 0x3e, 0x1e, 0x13, 0x7d, 0x4c, 0x4a,
	0xb9, 0x4a, 0x0e, 0x2e, 0xab, 0xcf, 0xae, 0x7e, 0x9b, 0xb4, 0x17, 0xab, 0x5b, 0xc5, 0x28, 0xe4,
	0x56, 0x4f, 0x57, 0x00, 0x4b, 0xf1, 0xda, 0xe9, 0xa8, 0xf2, 0xe7, 0x0c, 0x58, 0x4a, 0x84, 0xd7,
	0xdb, 0x1c, 0x0a, 0x28, 0x0c, 0x4d, 0x4b, 0xef, 0x39, 0x88, 0x8c, 0x92, 0xfd, 0x71, 0x37, 0x0a,
	0xb9, 0xf5, 0x58, 0x79, 0x56, 0xe2, 0xb5, 0x3b, 0xf4, 0xfc, 0x0a, 0x91, 0x11, 0x9d, 0x4b, 0xf4,
	0x29, 0x49, 0xbe, 0xda, 0xc2, 0x5c, 0x9a, 0xe5, 0x35, 0x56, 0x6c, 0x1d, 0xbc, 0xdb, 0x37, 0xb0,
	0xa4, 0x23, 0xef, 0x04, 0x3b, 0x9e, 0x64, 0xe1, 0x3e, 0xb2, 0x64, 0x0a, 0xe9, 0xa2, 0x3e, 0xf1,
	0xe2, 0xe5, 0xe9, 0xea, 0x0e, 0x16, 0x6d, 0x3c, 0xd0, 0x2d, 0x4f, 0x4e, 0x84, 0x32, 0x0b, 0xe5,
	0xf9, 0x1f, 0x71, 0xf4, 0x3f, 0x93, 0xd6, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x18, 0x83, 0xa3,
	0x2d, 0xa3, 0x07, 0x00, 0x00,
}
