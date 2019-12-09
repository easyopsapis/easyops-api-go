// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log_search.proto

package log_search

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	log_search "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/log_search"
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
//LogSearch请求
type LogSearchRequest struct {
	//
	//日志路径
	LogFileName string `protobuf:"bytes,1,opt,name=log_file_name,json=logFileName,proto3" json:"log_file_name" form:"log_file_name"`
	//
	//是否允许模糊匹配日志路径
	EnabledFilePattern bool `protobuf:"varint,2,opt,name=enabled_file_pattern,json=enabledFilePattern,proto3" json:"enabled_file_pattern" form:"enabled_file_pattern"`
	//
	//搜索关键字。多个用空格分隔，and逻辑
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords" form:"keywords"`
	//
	//最大返回行数。如需修改最大查询行数, 以逗号分隔, 如 10,1000 表示一次最多查询1000行, 返回匹配到关键字的10行
	MaxReturnLines string `protobuf:"bytes,4,opt,name=max_return_lines,json=maxReturnLines,proto3" json:"max_return_lines" form:"max_return_lines"`
	//
	//关键字前n行
	BeforeLines int32 `protobuf:"varint,5,opt,name=before_lines,json=beforeLines,proto3" json:"before_lines" form:"before_lines"`
	//
	//关键字后n行
	AfterLines int32 `protobuf:"varint,6,opt,name=after_lines,json=afterLines,proto3" json:"after_lines" form:"after_lines"`
	//
	//搜索主机IP列表
	Agents               []string `protobuf:"bytes,7,rep,name=agents,proto3" json:"agents" form:"agents"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSearchRequest) Reset()         { *m = LogSearchRequest{} }
func (m *LogSearchRequest) String() string { return proto.CompactTextString(m) }
func (*LogSearchRequest) ProtoMessage()    {}
func (*LogSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e24305d63fb477, []int{0}
}
func (m *LogSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchRequest.Unmarshal(m, b)
}
func (m *LogSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchRequest.Marshal(b, m, deterministic)
}
func (m *LogSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchRequest.Merge(m, src)
}
func (m *LogSearchRequest) XXX_Size() int {
	return xxx_messageInfo_LogSearchRequest.Size(m)
}
func (m *LogSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchRequest proto.InternalMessageInfo

func (m *LogSearchRequest) GetLogFileName() string {
	if m != nil {
		return m.LogFileName
	}
	return ""
}

func (m *LogSearchRequest) GetEnabledFilePattern() bool {
	if m != nil {
		return m.EnabledFilePattern
	}
	return false
}

func (m *LogSearchRequest) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *LogSearchRequest) GetMaxReturnLines() string {
	if m != nil {
		return m.MaxReturnLines
	}
	return ""
}

func (m *LogSearchRequest) GetBeforeLines() int32 {
	if m != nil {
		return m.BeforeLines
	}
	return 0
}

func (m *LogSearchRequest) GetAfterLines() int32 {
	if m != nil {
		return m.AfterLines
	}
	return 0
}

func (m *LogSearchRequest) GetAgents() []string {
	if m != nil {
		return m.Agents
	}
	return nil
}

//
//LogSearch返回
type LogSearchResponse struct {
	//
	//code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//返回码字符串表示
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	//
	//data
	Data                 []*log_search.LogSearchData `protobuf:"bytes,3,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LogSearchResponse) Reset()         { *m = LogSearchResponse{} }
func (m *LogSearchResponse) String() string { return proto.CompactTextString(m) }
func (*LogSearchResponse) ProtoMessage()    {}
func (*LogSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e24305d63fb477, []int{1}
}
func (m *LogSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchResponse.Unmarshal(m, b)
}
func (m *LogSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchResponse.Marshal(b, m, deterministic)
}
func (m *LogSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchResponse.Merge(m, src)
}
func (m *LogSearchResponse) XXX_Size() int {
	return xxx_messageInfo_LogSearchResponse.Size(m)
}
func (m *LogSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchResponse proto.InternalMessageInfo

func (m *LogSearchResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogSearchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogSearchResponse) GetData() []*log_search.LogSearchData {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//LogSearchApi返回
type LogSearchResponseWrapper struct {
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
	Data                 *LogSearchResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LogSearchResponseWrapper) Reset()         { *m = LogSearchResponseWrapper{} }
func (m *LogSearchResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*LogSearchResponseWrapper) ProtoMessage()    {}
func (*LogSearchResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e24305d63fb477, []int{2}
}
func (m *LogSearchResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSearchResponseWrapper.Unmarshal(m, b)
}
func (m *LogSearchResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSearchResponseWrapper.Marshal(b, m, deterministic)
}
func (m *LogSearchResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSearchResponseWrapper.Merge(m, src)
}
func (m *LogSearchResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_LogSearchResponseWrapper.Size(m)
}
func (m *LogSearchResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSearchResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_LogSearchResponseWrapper proto.InternalMessageInfo

func (m *LogSearchResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogSearchResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *LogSearchResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LogSearchResponseWrapper) GetData() *LogSearchResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LogSearchRequest)(nil), "log_search.LogSearchRequest")
	proto.RegisterType((*LogSearchResponse)(nil), "log_search.LogSearchResponse")
	proto.RegisterType((*LogSearchResponseWrapper)(nil), "log_search.LogSearchResponseWrapper")
}

func init() { proto.RegisterFile("log_search.proto", fileDescriptor_38e24305d63fb477) }

var fileDescriptor_38e24305d63fb477 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xdf, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x15, 0xf6, 0x4f, 0x9b, 0xc9, 0xb6, 0x5d, 0xdc, 0x8a, 0x9a, 0x56, 0xc8, 0xd1, 0xf0,
	0x47, 0xe3, 0xb2, 0x8e, 0x13, 0x27, 0xd9, 0xa5, 0x11, 0x62, 0x45, 0x44, 0x7b, 0x55, 0x21, 0x30,
	0x17, 0x48, 0x24, 0x21, 0x9a, 0xc4, 0x13, 0x37, 0xaa, 0xed, 0x31, 0x33, 0x0e, 0xbb, 0x55, 0xec,
	0x67, 0x40, 0xe2, 0x21, 0x10, 0xaf, 0xc1, 0x3b, 0x70, 0x6d, 0x24, 0x24, 0x5e, 0xc0, 0x4f, 0x80,
	0x3c, 0xe3, 0xdd, 0x1a, 0x12, 0x43, 0x2f, 0xb6, 0x7b, 0x65, 0xcf, 0xf9, 0xbe, 0xef, 0x9c, 0xdf,
	0xdc, 0xcc, 0x01, 0x87, 0x1e, 0x75, 0xa7, 0x9c, 0x60, 0x36, 0x7f, 0xde, 0x0a, 0x19, 0x8d, 0xa8,
	0x02, 0x5e, 0x55, 0x1e, 0x18, 0xee, 0x32, 0x7a, 0xbe, 0x9a, 0xb5, 0xe6, 0xd4, 0x37, 0x5d, 0xea,
	0x52, 0x53, 0x58, 0x66, 0xab, 0x85, 0x38, 0x89, 0x83, 0xf8, 0x93, 0xd1, 0x07, 0xdf, 0xb9, 0xb4,
	0x45, 0x30, 0x7f, 0x49, 0x43, 0xde, 0xf2, 0xe8, 0x1c, 0x7b, 0xe6, 0x9c, 0x06, 0x11, 0xc3, 0xf3,
	0x88, 0xcb, 0x24, 0x23, 0x21, 0x35, 0x7c, 0xea, 0x10, 0x8f, 0x9b, 0x85, 0xd1, 0x14, 0x47, 0xf3,
	0xd5, 0xd0, 0xd2, 0xef, 0xd4, 0xc1, 0x11, 0x2e, 0x7a, 0x1f, 0x97, 0x50, 0xfc, 0xb3, 0x65, 0xf4,
	0x82, 0x9e, 0x99, 0x2e, 0x35, 0x84, 0x68, 0xfc, 0x88, 0xbd, 0xa5, 0x83, 0x23, 0xca, 0xb8, 0x79,
	0xf9, 0x2b, 0x73, 0xf0, 0xf7, 0x03, 0x70, 0xf8, 0x8c, 0xba, 0xdf, 0x88, 0x86, 0x36, 0xf9, 0x61,
	0x45, 0x78, 0xa4, 0x7c, 0x0a, 0x6e, 0xe5, 0x53, 0x16, 0x4b, 0x8f, 0x4c, 0x03, 0xec, 0x13, 0xb5,
	0xd6, 0xac, 0xa1, 0xfa, 0x50, 0xcd, 0x52, 0xed, 0xde, 0x82, 0x32, 0x7f, 0x00, 0xff, 0x21, 0x43,
	0xbb, 0xe1, 0x51, 0xf7, 0xe9, 0xd2, 0x23, 0x5f, 0x62, 0x9f, 0x28, 0x5f, 0x83, 0x7b, 0x24, 0xc0,
	0x33, 0x8f, 0x38, 0xd2, 0x12, 0xe2, 0x28, 0x22, 0x2c, 0x50, 0xdf, 0x6a, 0xd6, 0xd0, 0xcd, 0xa1,
	0x96, 0xa5, 0xda, 0x43, 0xd9, 0x64, 0x9b, 0x0b, 0xda, 0x4a, 0x51, 0xce, 0xfb, 0x7d, 0x25, 0x8b,
	0x8a, 0x09, 0x6e, 0xbe, 0x20, 0x2f, 0xcf, 0x28, 0x73, 0xb8, 0xba, 0x23, 0x58, 0xee, 0x66, 0xa9,
	0x76, 0x47, 0xb6, 0xb9, 0x50, 0xa0, 0x7d, 0x69, 0x52, 0x9e, 0x80, 0x43, 0x1f, 0x9f, 0x4f, 0x19,
	0x89, 0x56, 0x2c, 0x98, 0x7a, 0xcb, 0x80, 0x70, 0x75, 0x57, 0x04, 0x1f, 0x66, 0xa9, 0x76, 0x5f,
	0x06, 0xff, 0xed, 0x80, 0xf6, 0x6d, 0x1f, 0x9f, 0xdb, 0xa2, 0xf2, 0x2c, 0x2f, 0x28, 0x03, 0x70,
	0x30, 0x23, 0x0b, 0xca, 0x48, 0xd1, 0x62, 0xaf, 0x59, 0x43, 0x7b, 0xc3, 0xfb, 0x59, 0xaa, 0xdd,
	0x95, 0x2d, 0xca, 0x2a, 0xb4, 0x1b, 0xf2, 0x28, 0xb3, 0x27, 0xa0, 0x81, 0x17, 0x11, 0x61, 0x45,
	0x74, 0x5f, 0x44, 0xdf, 0xc9, 0x52, 0x4d, 0x91, 0xd1, 0x92, 0x08, 0x6d, 0x20, 0x4e, 0x32, 0xf8,
	0x5b, 0x1d, 0xec, 0x63, 0x97, 0x04, 0x11, 0x57, 0x6f, 0x34, 0x77, 0x50, 0x7d, 0xf8, 0x6b, 0x3d,
	0x4b, 0xb5, 0x5b, 0x45, 0x4a, 0x08, 0xf0, 0xcf, 0x3f, 0xb4, 0x9f, 0xeb, 0xe0, 0xa7, 0xfa, 0xf7,
	0x08, 0x59, 0xa8, 0x3f, 0x6a, 0x1b, 0xfd, 0xc9, 0xba, 0x93, 0xc4, 0xa3, 0xb6, 0xd1, 0x9b, 0x8c,
	0x9d, 0x75, 0x27, 0xd1, 0xf3, 0xff, 0xce, 0xe4, 0x34, 0x3f, 0x1c, 0x59, 0x89, 0x8e, 0xc6, 0xad,
	0xd7, 0x74, 0xea, 0xeb, 0x6e, 0xa2, 0xc7, 0x63, 0xfe, 0x08, 0x21, 0x34, 0x6a, 0x1b, 0x8f, 0x3f,
	0x37, 0x9e, 0x62, 0x63, 0x31, 0x59, 0x77, 0x8e, 0x7a, 0xc9, 0x40, 0x5f, 0x9f, 0x24, 0x1b, 0xd5,
	0x78, 0xa0, 0xeb, 0xf1, 0x56, 0xf3, 0x71, 0x82, 0x06, 0x1b, 0x6e, 0x84, 0x2c, 0xc9, 0x11, 0x5b,
	0x05, 0x45, 0xdc, 0x19, 0x3b, 0x63, 0x27, 0x1e, 0x75, 0x8c, 0xc7, 0x39, 0x87, 0x84, 0xfd, 0x1f,
	0x8f, 0xc4, 0xac, 0x9c, 0xdc, 0x4f, 0x10, 0xda, 0x9c, 0xad, 0xcb, 0x2b, 0xc6, 0x83, 0x6b, 0x61,
	0xe8, 0x55, 0x32, 0xe4, 0xb1, 0x6d, 0xd2, 0xe9, 0x55, 0x82, 0xfd, 0x07, 0x59, 0xb7, 0x92, 0xac,
	0x57, 0x41, 0xb6, 0x6e, 0x1f, 0x59, 0xc9, 0x35, 0xd1, 0x59, 0x95, 0x74, 0xfd, 0x6a, 0xba, 0xee,
	0x75, 0xd1, 0x75, 0x2a, 0xe9, 0x8e, 0xab, 0xe9, 0x7a, 0x6f, 0x82, 0x6e, 0x50, 0x05, 0x72, 0x52,
	0x0d, 0xd2, 0xbf, 0x7a, 0x10, 0x1d, 0x7d, 0xd8, 0xfa, 0x58, 0x3f, 0x1d, 0xf3, 0x47, 0x1f, 0xd8,
	0xc5, 0xc3, 0x05, 0x7f, 0xa9, 0x81, 0xb7, 0x4b, 0x6b, 0x85, 0x87, 0x34, 0xe0, 0x44, 0x79, 0x1f,
	0xec, 0xce, 0xa9, 0x23, 0xd7, 0xc9, 0xde, 0xf0, 0x4e, 0x96, 0x6a, 0x0d, 0xf9, 0xaa, 0xe5, 0x55,
	0x68, 0x0b, 0x51, 0x39, 0x02, 0x37, 0x7c, 0xc2, 0x39, 0x76, 0x89, 0xd8, 0x18, 0xf5, 0xa1, 0x92,
	0xa5, 0xda, 0xed, 0xe2, 0xc5, 0x96, 0x02, 0xb4, 0x2f, 0x2c, 0xca, 0x67, 0x60, 0x37, 0xdf, 0x82,
	0xea, 0x4e, 0x73, 0x07, 0x35, 0xac, 0x77, 0x5b, 0xa5, 0x7d, 0x7d, 0x39, 0xff, 0x0b, 0x1c, 0xe1,
	0xf2, 0xb4, 0x3c, 0x00, 0x6d, 0x91, 0x83, 0x7f, 0xd5, 0x80, 0xba, 0x01, 0xfa, 0x2d, 0xc3, 0x61,
	0x48, 0xd8, 0xeb, 0xf1, 0x7e, 0x02, 0x1a, 0xf9, 0xf7, 0xc9, 0x79, 0xe8, 0xe1, 0x65, 0x50, 0x30,
	0x97, 0xde, 0xf9, 0x92, 0x08, 0xed, 0xb2, 0x55, 0xf9, 0x08, 0xec, 0x11, 0xc6, 0x28, 0x2b, 0x56,
	0xda, 0x61, 0x96, 0x6a, 0x07, 0xc5, 0x66, 0xcc, 0xcb, 0xd0, 0x96, 0xb2, 0x32, 0x2c, 0xee, 0x98,
	0x2f, 0xb0, 0x86, 0xf5, 0xde, 0xd6, 0x3b, 0x5e, 0xa0, 0x57, 0xdc, 0x73, 0xb6, 0x2f, 0xd6, 0x7d,
	0xf7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xbb, 0x52, 0xed, 0xd1, 0x08, 0x00, 0x00,
}