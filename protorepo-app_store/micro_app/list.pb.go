// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	app_store "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/app_store"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//ListAppStoreMicroApp请求
type ListAppStoreMicroAppRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//查询条件
	Query *ListAppStoreMicroAppRequest_Query `protobuf:"bytes,3,opt,name=query,proto3" json:"query" form:"query"`
	//
	//条件查询：按照字段正/倒序 :1表示升序, -1表示降序
	Sort                 *ListAppStoreMicroAppRequest_Sort `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort" form:"sort"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListAppStoreMicroAppRequest) Reset()         { *m = ListAppStoreMicroAppRequest{} }
func (m *ListAppStoreMicroAppRequest) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppRequest) ProtoMessage()    {}
func (*ListAppStoreMicroAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListAppStoreMicroAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppRequest.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppRequest.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppRequest.Merge(m, src)
}
func (m *ListAppStoreMicroAppRequest) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppRequest.Size(m)
}
func (m *ListAppStoreMicroAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppRequest proto.InternalMessageInfo

func (m *ListAppStoreMicroAppRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListAppStoreMicroAppRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAppStoreMicroAppRequest) GetQuery() *ListAppStoreMicroAppRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ListAppStoreMicroAppRequest) GetSort() *ListAppStoreMicroAppRequest_Sort {
	if m != nil {
		return m.Sort
	}
	return nil
}

type ListAppStoreMicroAppRequest_Query struct {
	//
	//按name搜索,%xx%:模糊匹配,xx%:前缀匹配
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//按id列表查询
	Id                   []string `protobuf:"bytes,2,rep,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppStoreMicroAppRequest_Query) Reset()         { *m = ListAppStoreMicroAppRequest_Query{} }
func (m *ListAppStoreMicroAppRequest_Query) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppRequest_Query) ProtoMessage()    {}
func (*ListAppStoreMicroAppRequest_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0, 0}
}
func (m *ListAppStoreMicroAppRequest_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Query.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppRequest_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Query.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppRequest_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppRequest_Query.Merge(m, src)
}
func (m *ListAppStoreMicroAppRequest_Query) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Query.Size(m)
}
func (m *ListAppStoreMicroAppRequest_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppRequest_Query.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppRequest_Query proto.InternalMessageInfo

func (m *ListAppStoreMicroAppRequest_Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListAppStoreMicroAppRequest_Query) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListAppStoreMicroAppRequest_Sort struct {
	//
	//按name排序，1为升序，-1为降序
	Name                 int32    `protobuf:"varint,1,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppStoreMicroAppRequest_Sort) Reset()         { *m = ListAppStoreMicroAppRequest_Sort{} }
func (m *ListAppStoreMicroAppRequest_Sort) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppRequest_Sort) ProtoMessage()    {}
func (*ListAppStoreMicroAppRequest_Sort) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0, 1}
}
func (m *ListAppStoreMicroAppRequest_Sort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Sort.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppRequest_Sort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Sort.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppRequest_Sort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppRequest_Sort.Merge(m, src)
}
func (m *ListAppStoreMicroAppRequest_Sort) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppRequest_Sort.Size(m)
}
func (m *ListAppStoreMicroAppRequest_Sort) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppRequest_Sort.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppRequest_Sort proto.InternalMessageInfo

func (m *ListAppStoreMicroAppRequest_Sort) GetName() int32 {
	if m != nil {
		return m.Name
	}
	return 0
}

//
//ListAppStoreMicroApp返回
type ListAppStoreMicroAppResponse struct {
	//
	//返回总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total"`
	//
	//页数
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//该页大小
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//数据列表
	List                 []*ListAppStoreMicroAppResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListAppStoreMicroAppResponse) Reset()         { *m = ListAppStoreMicroAppResponse{} }
func (m *ListAppStoreMicroAppResponse) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppResponse) ProtoMessage()    {}
func (*ListAppStoreMicroAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListAppStoreMicroAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppResponse.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppResponse.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppResponse.Merge(m, src)
}
func (m *ListAppStoreMicroAppResponse) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppResponse.Size(m)
}
func (m *ListAppStoreMicroAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppResponse proto.InternalMessageInfo

func (m *ListAppStoreMicroAppResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListAppStoreMicroAppResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListAppStoreMicroAppResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAppStoreMicroAppResponse) GetList() []*ListAppStoreMicroAppResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListAppStoreMicroAppResponse_List struct {
	//
	//当前版本
	CurrentVersion *app_store.AppVersion `protobuf:"bytes,1,opt,name=currentVersion,proto3" json:"currentVersion" form:"currentVersion"`
	//
	//小产品名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//小产品id
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id" form:"id"`
	//
	//图标url
	Icon string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//小产品概要介绍
	Intro string `protobuf:"bytes,5,opt,name=intro,proto3" json:"intro" form:"intro"`
	//
	//功能预览
	Preview []string `protobuf:"bytes,6,rep,name=preview,proto3" json:"preview" form:"preview"`
	//
	//功能详细介绍
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description" form:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppStoreMicroAppResponse_List) Reset()         { *m = ListAppStoreMicroAppResponse_List{} }
func (m *ListAppStoreMicroAppResponse_List) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppResponse_List) ProtoMessage()    {}
func (*ListAppStoreMicroAppResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0}
}
func (m *ListAppStoreMicroAppResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppResponse_List.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppResponse_List.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppResponse_List.Merge(m, src)
}
func (m *ListAppStoreMicroAppResponse_List) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppResponse_List.Size(m)
}
func (m *ListAppStoreMicroAppResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppResponse_List proto.InternalMessageInfo

func (m *ListAppStoreMicroAppResponse_List) GetCurrentVersion() *app_store.AppVersion {
	if m != nil {
		return m.CurrentVersion
	}
	return nil
}

func (m *ListAppStoreMicroAppResponse_List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListAppStoreMicroAppResponse_List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListAppStoreMicroAppResponse_List) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ListAppStoreMicroAppResponse_List) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func (m *ListAppStoreMicroAppResponse_List) GetPreview() []string {
	if m != nil {
		return m.Preview
	}
	return nil
}

func (m *ListAppStoreMicroAppResponse_List) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

//
//ListAppStoreMicroAppApi返回
type ListAppStoreMicroAppResponseWrapper struct {
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
	Data                 *ListAppStoreMicroAppResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ListAppStoreMicroAppResponseWrapper) Reset()         { *m = ListAppStoreMicroAppResponseWrapper{} }
func (m *ListAppStoreMicroAppResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListAppStoreMicroAppResponseWrapper) ProtoMessage()    {}
func (*ListAppStoreMicroAppResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}
func (m *ListAppStoreMicroAppResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppStoreMicroAppResponseWrapper.Unmarshal(m, b)
}
func (m *ListAppStoreMicroAppResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppStoreMicroAppResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListAppStoreMicroAppResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppStoreMicroAppResponseWrapper.Merge(m, src)
}
func (m *ListAppStoreMicroAppResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListAppStoreMicroAppResponseWrapper.Size(m)
}
func (m *ListAppStoreMicroAppResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppStoreMicroAppResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppStoreMicroAppResponseWrapper proto.InternalMessageInfo

func (m *ListAppStoreMicroAppResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListAppStoreMicroAppResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListAppStoreMicroAppResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListAppStoreMicroAppResponseWrapper) GetData() *ListAppStoreMicroAppResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAppStoreMicroAppRequest)(nil), "micro_app.ListAppStoreMicroAppRequest")
	proto.RegisterType((*ListAppStoreMicroAppRequest_Query)(nil), "micro_app.ListAppStoreMicroAppRequest.Query")
	proto.RegisterType((*ListAppStoreMicroAppRequest_Sort)(nil), "micro_app.ListAppStoreMicroAppRequest.Sort")
	proto.RegisterType((*ListAppStoreMicroAppResponse)(nil), "micro_app.ListAppStoreMicroAppResponse")
	proto.RegisterType((*ListAppStoreMicroAppResponse_List)(nil), "micro_app.ListAppStoreMicroAppResponse.List")
	proto.RegisterType((*ListAppStoreMicroAppResponseWrapper)(nil), "micro_app.ListAppStoreMicroAppResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x4f, 0x13, 0x4b,
	0x14, 0xbe, 0xdd, 0x6e, 0xe1, 0x76, 0x7a, 0x2f, 0xe2, 0x1a, 0xcc, 0x5a, 0x35, 0xdb, 0x0c, 0x04,
	0x31, 0xd0, 0xdd, 0x08, 0x46, 0x0d, 0x6f, 0x6d, 0xe2, 0x93, 0x98, 0xc8, 0xe0, 0xaf, 0x68, 0xb4,
	0x59, 0x76, 0x87, 0x3a, 0xb1, 0xdd, 0x19, 0x66, 0xa6, 0x20, 0x18, 0x9f, 0xf9, 0x43, 0xfc, 0xbf,
	0x4a, 0xc2, 0x9f, 0xd0, 0x47, 0x9f, 0xcc, 0x9c, 0x59, 0xca, 0x96, 0x00, 0x81, 0xa7, 0xdd, 0x39,
	0xdf, 0x77, 0xbe, 0x39, 0xe7, 0x7c, 0x67, 0x17, 0xa1, 0x1e, 0x53, 0x3a, 0x14, 0x92, 0x6b, 0xee,
	0x55, 0xfb, 0x2c, 0x91, 0xbc, 0x13, 0x0b, 0x51, 0x6f, 0x76, 0x99, 0xfe, 0x36, 0xd8, 0x0e, 0x13,
	0xde, 0x8f, 0xba, 0xbc, 0xcb, 0x23, 0x60, 0x6c, 0x0f, 0x76, 0xe0, 0x04, 0x07, 0x78, 0xb3, 0x99,
	0xf5, 0x77, 0x5d, 0x1e, 0xd2, 0x58, 0x1d, 0x70, 0xa1, 0xc2, 0x1e, 0x4f, 0xe2, 0x5e, 0x94, 0xf0,
	0x4c, 0xcb, 0x38, 0xd1, 0xca, 0x66, 0x4a, 0x2a, 0x78, 0xb3, 0xcf, 0x53, 0xda, 0x53, 0x51, 0x4e,
	0x8c, 0xe0, 0x18, 0xc5, 0x42, 0x74, 0x94, 0x21, 0xc0, 0xdb, 0x1e, 0x95, 0x8a, 0xf1, 0x2c, 0x97,
	0x7d, 0x56, 0xa8, 0xa2, 0xbf, 0xcf, 0xf4, 0x77, 0xbe, 0x1f, 0x75, 0x79, 0x13, 0xc0, 0xe6, 0x5e,
	0xdc, 0x63, 0x69, 0xac, 0xb9, 0x54, 0xd1, 0xf8, 0xd5, 0xe6, 0xe1, 0xdf, 0x65, 0x74, 0x7f, 0x83,
	0x29, 0xdd, 0x12, 0x62, 0xcb, 0x48, 0xbf, 0x36, 0x7d, 0xb5, 0x84, 0x20, 0x74, 0x77, 0x40, 0x95,
	0xf6, 0x1e, 0x23, 0x57, 0xc4, 0x5d, 0xea, 0x97, 0x1a, 0xa5, 0xa5, 0x4a, 0x7b, 0x6e, 0x34, 0x0c,
	0x6a, 0x3b, 0x5c, 0xf6, 0xd7, 0xb1, 0x89, 0xe2, 0x93, 0xe3, 0xc0, 0x99, 0xfd, 0x87, 0x00, 0xc5,
	0x7b, 0x8e, 0xaa, 0xe6, 0xd9, 0x51, 0xec, 0x90, 0xfa, 0x0e, 0xf0, 0xeb, 0xa3, 0x61, 0x30, 0x7b,
	0xc6, 0x07, 0xe8, 0x34, 0xe9, 0x5f, 0x13, 0xd9, 0x62, 0x87, 0xd4, 0x7b, 0x8b, 0x2a, 0xbb, 0x03,
	0x2a, 0x0f, 0xfc, 0x72, 0xa3, 0xb4, 0x54, 0x5b, 0x5d, 0x09, 0xc7, 0xc3, 0x0d, 0xaf, 0x28, 0x2d,
	0xdc, 0x34, 0x39, 0xed, 0xd9, 0xd1, 0x30, 0xf8, 0xcf, 0x5e, 0x01, 0x22, 0x98, 0x58, 0x31, 0xef,
	0x0d, 0x72, 0x15, 0x97, 0xda, 0x77, 0x41, 0x74, 0xf9, 0x9a, 0xa2, 0x5b, 0x5c, 0xea, 0xf6, 0xad,
	0xb3, 0x36, 0x8d, 0x04, 0x26, 0xa0, 0x54, 0x7f, 0x85, 0x2a, 0x70, 0xa7, 0x37, 0x8f, 0xdc, 0x2c,
	0xee, 0xdb, 0xa1, 0x54, 0x8b, 0x6c, 0x13, 0xc5, 0x04, 0x40, 0xef, 0x21, 0x72, 0x58, 0xea, 0x3b,
	0x8d, 0xf2, 0x52, 0xb5, 0xfd, 0xff, 0x68, 0x18, 0x54, 0x2d, 0x85, 0xa5, 0x98, 0x38, 0x2c, 0xad,
	0x2f, 0x23, 0xd7, 0xdc, 0x35, 0xa1, 0x55, 0xb9, 0x44, 0x0b, 0x1f, 0x55, 0xd0, 0x83, 0x8b, 0xab,
	0x56, 0x82, 0x67, 0x8a, 0x7a, 0x8b, 0xa8, 0xa2, 0xb9, 0x8e, 0x7b, 0xb9, 0x4c, 0x61, 0x28, 0x10,
	0xc6, 0xc4, 0xc2, 0x63, 0x3b, 0x9d, 0x1b, 0xda, 0x59, 0xbe, 0x81, 0x9d, 0x9b, 0xc8, 0x35, 0x5f,
	0x8a, 0xef, 0x36, 0xca, 0xd7, 0x72, 0xd3, 0xb6, 0x00, 0x60, 0xb1, 0x7f, 0xa3, 0x81, 0x09, 0x48,
	0xd5, 0x8f, 0xca, 0xc8, 0x35, 0xb8, 0xf7, 0x11, 0xcd, 0x24, 0x03, 0x29, 0x69, 0xa6, 0xdf, 0xdb,
	0xf5, 0x87, 0x86, 0x6b, 0xab, 0x73, 0xe1, 0xf8, 0xe3, 0x08, 0x5b, 0x42, 0xe4, 0x60, 0xfb, 0xde,
	0x68, 0x18, 0xcc, 0x59, 0xb9, 0xc9, 0x34, 0x4c, 0xce, 0xe9, 0x78, 0xeb, 0xb9, 0x0f, 0x0e, 0x78,
	0xba, 0x78, 0xce, 0x87, 0x93, 0xe3, 0xe0, 0x0e, 0xba, 0xfd, 0xf5, 0x73, 0xdc, 0x3c, 0x6c, 0x35,
	0x3f, 0x75, 0xbe, 0xfc, 0x7c, 0xb2, 0xb2, 0xb6, 0xfa, 0x6b, 0x21, 0xb7, 0xfa, 0x29, 0x58, 0x5d,
	0x86, 0xcc, 0x85, 0x09, 0xab, 0x2f, 0xcb, 0x73, 0x58, 0x6a, 0x9c, 0x67, 0x09, 0xcf, 0x60, 0x41,
	0x27, 0xb6, 0xc8, 0x44, 0x31, 0x01, 0xd0, 0x18, 0xcb, 0x32, 0x2d, 0xb9, 0x5f, 0x01, 0x56, 0xc1,
	0x58, 0x08, 0x63, 0x62, 0x61, 0x6f, 0x05, 0x4d, 0x0b, 0x49, 0xf7, 0x18, 0xdd, 0xf7, 0xa7, 0x60,
	0xe5, 0xbc, 0xd1, 0x30, 0x98, 0xc9, 0xbd, 0xb2, 0x00, 0x26, 0xa7, 0x14, 0xef, 0x05, 0xaa, 0xa5,
	0x54, 0x25, 0x92, 0x09, 0x6d, 0x66, 0x38, 0x0d, 0xda, 0x77, 0x47, 0xc3, 0xc0, 0xb3, 0x19, 0x05,
	0x10, 0x93, 0x22, 0x15, 0xff, 0x29, 0xa1, 0xf9, 0xab, 0x6c, 0xfc, 0x20, 0x63, 0x21, 0xa8, 0x34,
	0xcd, 0x25, 0x3c, 0xbd, 0x60, 0xad, 0x4d, 0x14, 0x13, 0x00, 0x4d, 0x19, 0xe6, 0xf9, 0xf2, 0x87,
	0xe8, 0xc5, 0x2c, 0xcb, 0x47, 0x5f, 0x28, 0xa3, 0x00, 0x62, 0x52, 0xa4, 0x9a, 0xb1, 0x50, 0x29,
	0xb9, 0xcc, 0x87, 0x5e, 0x18, 0x0b, 0x84, 0x31, 0xb1, 0xb0, 0xb7, 0x81, 0xdc, 0x34, 0xd6, 0x71,
	0xfe, 0x13, 0x78, 0x74, 0xcd, 0x5d, 0x2c, 0xd6, 0x6b, 0xd2, 0x31, 0x01, 0x95, 0xed, 0x29, 0xf8,
	0x67, 0xae, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x69, 0x09, 0xb0, 0x7b, 0x0a, 0x06, 0x00, 0x00,
}
