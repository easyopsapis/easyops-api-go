// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_tool.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	tool "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
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
//ListTool请求
type ListToolRequest struct {
	//
	//是否显示具体信息
	Detail bool `protobuf:"varint,1,opt,name=detail,proto3" json:"detail" form:"detail"`
	//
	//是否显示插件
	Plugin bool `protobuf:"varint,2,opt,name=plugin,proto3" json:"plugin" form:"plugin"`
	//
	//工具的分类
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category" form:"category"`
	//
	//用户对工具具有的权限; 用","分割(read=查看, update=更新, delete=删除, execute=执行, execResult=执行结果查看, rootExec=Root执行)
	Permissions string `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions" form:"permissions"`
	//
	//仅选择有生产版本的工具 true | false. 默认是false
	OnlyProduction bool `protobuf:"varint,5,opt,name=onlyProduction,proto3" json:"onlyProduction" form:"onlyProduction"`
	//
	//显示listVisible=false的工具; true=显示 | false=不显示. 默认是false
	ShowInvisible bool `protobuf:"varint,6,opt,name=showInvisible,proto3" json:"showInvisible" form:"showInvisible"`
	//
	//标签，用","分隔，标签只支持字符串
	Tags                 string   `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags" form:"tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListToolRequest) Reset()         { *m = ListToolRequest{} }
func (m *ListToolRequest) String() string { return proto.CompactTextString(m) }
func (*ListToolRequest) ProtoMessage()    {}
func (*ListToolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a700652ad1db5a3, []int{0}
}
func (m *ListToolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListToolRequest.Unmarshal(m, b)
}
func (m *ListToolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListToolRequest.Marshal(b, m, deterministic)
}
func (m *ListToolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListToolRequest.Merge(m, src)
}
func (m *ListToolRequest) XXX_Size() int {
	return xxx_messageInfo_ListToolRequest.Size(m)
}
func (m *ListToolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListToolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListToolRequest proto.InternalMessageInfo

func (m *ListToolRequest) GetDetail() bool {
	if m != nil {
		return m.Detail
	}
	return false
}

func (m *ListToolRequest) GetPlugin() bool {
	if m != nil {
		return m.Plugin
	}
	return false
}

func (m *ListToolRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ListToolRequest) GetPermissions() string {
	if m != nil {
		return m.Permissions
	}
	return ""
}

func (m *ListToolRequest) GetOnlyProduction() bool {
	if m != nil {
		return m.OnlyProduction
	}
	return false
}

func (m *ListToolRequest) GetShowInvisible() bool {
	if m != nil {
		return m.ShowInvisible
	}
	return false
}

func (m *ListToolRequest) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

//
//ListTool返回
type ListToolResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*tool.Tool `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListToolResponse) Reset()         { *m = ListToolResponse{} }
func (m *ListToolResponse) String() string { return proto.CompactTextString(m) }
func (*ListToolResponse) ProtoMessage()    {}
func (*ListToolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a700652ad1db5a3, []int{1}
}
func (m *ListToolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListToolResponse.Unmarshal(m, b)
}
func (m *ListToolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListToolResponse.Marshal(b, m, deterministic)
}
func (m *ListToolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListToolResponse.Merge(m, src)
}
func (m *ListToolResponse) XXX_Size() int {
	return xxx_messageInfo_ListToolResponse.Size(m)
}
func (m *ListToolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListToolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListToolResponse proto.InternalMessageInfo

func (m *ListToolResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListToolResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListToolResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListToolResponse) GetList() []*tool.Tool {
	if m != nil {
		return m.List
	}
	return nil
}

//
//ListToolApi返回
type ListToolResponseWrapper struct {
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
	Data                 *ListToolResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListToolResponseWrapper) Reset()         { *m = ListToolResponseWrapper{} }
func (m *ListToolResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListToolResponseWrapper) ProtoMessage()    {}
func (*ListToolResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a700652ad1db5a3, []int{2}
}
func (m *ListToolResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListToolResponseWrapper.Unmarshal(m, b)
}
func (m *ListToolResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListToolResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListToolResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListToolResponseWrapper.Merge(m, src)
}
func (m *ListToolResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListToolResponseWrapper.Size(m)
}
func (m *ListToolResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListToolResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListToolResponseWrapper proto.InternalMessageInfo

func (m *ListToolResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListToolResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListToolResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListToolResponseWrapper) GetData() *ListToolResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListToolRequest)(nil), "basic.ListToolRequest")
	proto.RegisterType((*ListToolResponse)(nil), "basic.ListToolResponse")
	proto.RegisterType((*ListToolResponseWrapper)(nil), "basic.ListToolResponseWrapper")
}

func init() { proto.RegisterFile("list_tool.proto", fileDescriptor_5a700652ad1db5a3) }

var fileDescriptor_5a700652ad1db5a3 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xc1, 0x6e, 0xd4, 0x3c,
	0x10, 0xc7, 0xbf, 0x6d, 0x93, 0x7e, 0xad, 0x97, 0xb2, 0x8b, 0xa1, 0x34, 0xec, 0x25, 0x95, 0x91,
	0x50, 0x7b, 0xd8, 0x44, 0x2a, 0x12, 0x20, 0x84, 0x90, 0x58, 0x89, 0x43, 0x25, 0x0e, 0xc8, 0x20,
	0x71, 0xac, 0xbc, 0x89, 0x9b, 0x5a, 0x78, 0x33, 0xc1, 0xf6, 0xb6, 0xb4, 0x2f, 0xc9, 0x1b, 0x04,
	0x89, 0x47, 0xc8, 0x91, 0x13, 0xf2, 0x24, 0xed, 0x66, 0xf7, 0x94, 0x78, 0xfe, 0xbf, 0xbf, 0x3d,
	0xe3, 0x19, 0x93, 0x91, 0x56, 0xd6, 0x9d, 0x3b, 0x00, 0x9d, 0x54, 0x06, 0x1c, 0xd0, 0x70, 0x2e,
	0xac, 0xca, 0x26, 0xd3, 0x42, 0xb9, 0xcb, 0xe5, 0x3c, 0xc9, 0x60, 0x91, 0x16, 0x50, 0x40, 0x8a,
	0xea, 0x7c, 0x79, 0x81, 0x2b, 0x5c, 0xe0, 0x5f, 0xeb, 0x9a, 0x9c, 0x15, 0x90, 0x48, 0x61, 0x6f,
	0xa0, 0xb2, 0x89, 0x86, 0x4c, 0xe8, 0x34, 0x83, 0xd2, 0x19, 0x91, 0x39, 0xdb, 0x3a, 0x8d, 0xac,
	0x60, 0xba, 0x80, 0x5c, 0x6a, 0x9b, 0x76, 0x60, 0x8a, 0xcb, 0xd4, 0x9f, 0x9d, 0xae, 0x12, 0x98,
	0xbc, 0xea, 0x9d, 0xbc, 0xb8, 0x56, 0xee, 0x3b, 0x5c, 0xa7, 0x05, 0x4c, 0x51, 0x9c, 0x5e, 0x09,
	0xad, 0x72, 0xe1, 0xc0, 0xd8, 0xf4, 0xfe, 0xb7, 0xf5, 0xb1, 0xbf, 0x5b, 0x64, 0xf4, 0x49, 0x59,
	0xf7, 0x15, 0x40, 0x73, 0xf9, 0x63, 0x29, 0xad, 0xa3, 0x27, 0x64, 0x27, 0x97, 0x4e, 0x28, 0x1d,
	0x0d, 0x8e, 0x06, 0xc7, 0xbb, 0xb3, 0x47, 0x4d, 0x1d, 0xef, 0x5f, 0x80, 0x59, 0xbc, 0x65, 0x6d,
	0x9c, 0xf1, 0x0e, 0xf0, 0x68, 0xa5, 0x97, 0x85, 0x2a, 0xa3, 0xad, 0x4d, 0xb4, 0x8d, 0x33, 0xde,
	0x01, 0x34, 0x25, 0xbb, 0x99, 0x70, 0xb2, 0x00, 0x73, 0x13, 0x6d, 0x1f, 0x0d, 0x8e, 0xf7, 0x66,
	0x8f, 0x9b, 0x3a, 0x1e, 0xb5, 0xf0, 0x9d, 0xc2, 0xf8, 0x3d, 0x44, 0xdf, 0x90, 0x61, 0x25, 0xcd,
	0x42, 0x59, 0xab, 0xa0, 0xb4, 0x51, 0x80, 0x9e, 0xa7, 0x4d, 0x1d, 0xd3, 0xee, 0x80, 0x95, 0xc8,
	0x78, 0x1f, 0xa5, 0x1f, 0xc8, 0x43, 0x28, 0xf5, 0xcd, 0x67, 0x03, 0xf9, 0x32, 0x73, 0x0a, 0xca,
	0x28, 0xc4, 0xec, 0x9e, 0x35, 0x75, 0x7c, 0xd0, 0x9a, 0xd7, 0x75, 0xc6, 0x37, 0x0c, 0xf4, 0x3d,
	0xd9, 0xb7, 0x97, 0x70, 0x7d, 0x56, 0x5e, 0x29, 0xab, 0xe6, 0x5a, 0x46, 0x3b, 0xb8, 0x43, 0xd4,
	0xd4, 0xf1, 0x93, 0x76, 0x87, 0x35, 0x99, 0xf1, 0x75, 0x9c, 0x3e, 0x27, 0x81, 0x13, 0x85, 0x8d,
	0xfe, 0xc7, 0xac, 0x47, 0x4d, 0x1d, 0x0f, 0x5b, 0x9b, 0x8f, 0x32, 0x8e, 0x22, 0xfb, 0x35, 0x20,
	0xe3, 0xd5, 0xe5, 0xdb, 0x0a, 0x4a, 0x2b, 0xe9, 0x09, 0x09, 0x2a, 0x51, 0x48, 0xbc, 0xfb, 0x70,
	0x76, 0xb0, 0x72, 0xfa, 0x28, 0xfb, 0xf3, 0x3b, 0xde, 0x1a, 0xff, 0xc7, 0x11, 0xa1, 0xaf, 0xc9,
	0x9e, 0xff, 0x9e, 0x5b, 0x75, 0x2b, 0xb1, 0x01, 0xe1, 0x6c, 0xd2, 0xd4, 0xf1, 0x78, 0xc5, 0xa3,
	0x74, 0x67, 0xda, 0xf5, 0x91, 0x2f, 0xea, 0x56, 0xd2, 0x17, 0x24, 0x74, 0xe0, 0x84, 0xc6, 0x46,
	0x84, 0xb3, 0x71, 0x53, 0xc7, 0x0f, 0xba, 0xf4, 0x7c, 0x98, 0xf1, 0x56, 0xa6, 0x29, 0x09, 0xfc,
	0xa4, 0x47, 0xc1, 0xd1, 0xf6, 0xf1, 0xf0, 0x94, 0x24, 0x38, 0x70, 0x3e, 0xdb, 0x7e, 0x45, 0x9e,
	0x60, 0x1c, 0x41, 0x56, 0x0f, 0xc8, 0xe1, 0x66, 0x45, 0xdf, 0x8c, 0xa8, 0x2a, 0x69, 0xfc, 0x95,
	0x64, 0x90, 0xdf, 0x15, 0xd6, 0xdb, 0xc0, 0x47, 0x19, 0x47, 0xd1, 0x37, 0xdd, 0x7f, 0x3f, 0xfe,
	0xac, 0xb4, 0xe8, 0xa6, 0x6a, 0xad, 0xe9, 0x3d, 0x91, 0xf1, 0x3e, 0xea, 0x6b, 0x92, 0xc6, 0x80,
	0xe9, 0x86, 0xab, 0x57, 0x13, 0x86, 0x19, 0x6f, 0x65, 0xfa, 0x8e, 0x04, 0xb9, 0x70, 0x02, 0xe7,
	0x69, 0x78, 0x7a, 0x98, 0xe0, 0xcb, 0x4d, 0x36, 0x93, 0xee, 0xe7, 0xe7, 0x71, 0xc6, 0xd1, 0x35,
	0xdf, 0xc1, 0x67, 0xf3, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xca, 0x06, 0x2f, 0x02,
	0x04, 0x00, 0x00,
}