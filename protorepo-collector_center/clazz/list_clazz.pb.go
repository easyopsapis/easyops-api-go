// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_clazz.proto

package clazz

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
//ListClazz请求
type ListClazzRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//是否启禁用
	Disabled int32 `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled" form:"disabled"`
	//
	//是否全部
	IsAll                int32    `protobuf:"varint,4,opt,name=isAll,proto3" json:"isAll" form:"isAll"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClazzRequest) Reset()         { *m = ListClazzRequest{} }
func (m *ListClazzRequest) String() string { return proto.CompactTextString(m) }
func (*ListClazzRequest) ProtoMessage()    {}
func (*ListClazzRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7055e34279888799, []int{0}
}
func (m *ListClazzRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClazzRequest.Unmarshal(m, b)
}
func (m *ListClazzRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClazzRequest.Marshal(b, m, deterministic)
}
func (m *ListClazzRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClazzRequest.Merge(m, src)
}
func (m *ListClazzRequest) XXX_Size() int {
	return xxx_messageInfo_ListClazzRequest.Size(m)
}
func (m *ListClazzRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClazzRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListClazzRequest proto.InternalMessageInfo

func (m *ListClazzRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListClazzRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListClazzRequest) GetDisabled() int32 {
	if m != nil {
		return m.Disabled
	}
	return 0
}

func (m *ListClazzRequest) GetIsAll() int32 {
	if m != nil {
		return m.IsAll
	}
	return 0
}

//
//ListClazz返回
type ListClazzResponse struct {
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
	List                 []*ListClazzResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListClazzResponse) Reset()         { *m = ListClazzResponse{} }
func (m *ListClazzResponse) String() string { return proto.CompactTextString(m) }
func (*ListClazzResponse) ProtoMessage()    {}
func (*ListClazzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7055e34279888799, []int{1}
}
func (m *ListClazzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClazzResponse.Unmarshal(m, b)
}
func (m *ListClazzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClazzResponse.Marshal(b, m, deterministic)
}
func (m *ListClazzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClazzResponse.Merge(m, src)
}
func (m *ListClazzResponse) XXX_Size() int {
	return xxx_messageInfo_ListClazzResponse.Size(m)
}
func (m *ListClazzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClazzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListClazzResponse proto.InternalMessageInfo

func (m *ListClazzResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListClazzResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListClazzResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListClazzResponse) GetList() []*ListClazzResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListClazzResponse_List struct {
	//
	//ORG
	Org int32 `protobuf:"varint,1,opt,name=org,proto3" json:"org" form:"org"`
	//
	//Id
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id" form:"Id"`
	//
	//名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name" form:"Name"`
	//
	//数据ID
	DataId string `protobuf:"bytes,4,opt,name=dataId,proto3" json:"dataId" form:"dataId"`
	//
	//启禁用
	Disabled bool `protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled" form:"disabled"`
	//
	//执行方法
	Fun string `protobuf:"bytes,6,opt,name=fun,proto3" json:"fun" form:"fun"`
	//
	//是否要求脚本
	RequiredScript bool `protobuf:"varint,7,opt,name=requiredScript,proto3" json:"requiredScript" form:"requiredScript"`
	//
	//数据要求字段
	RequiredFields []string `protobuf:"bytes,8,rep,name=requiredFields,proto3" json:"requiredFields" form:"requiredFields"`
	//
	//创建者
	Creator string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//修改者
	Modifer string `protobuf:"bytes,10,opt,name=modifer,proto3" json:"modifer" form:"modifer"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,11,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime                int32    `protobuf:"varint,12,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClazzResponse_List) Reset()         { *m = ListClazzResponse_List{} }
func (m *ListClazzResponse_List) String() string { return proto.CompactTextString(m) }
func (*ListClazzResponse_List) ProtoMessage()    {}
func (*ListClazzResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_7055e34279888799, []int{1, 0}
}
func (m *ListClazzResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClazzResponse_List.Unmarshal(m, b)
}
func (m *ListClazzResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClazzResponse_List.Marshal(b, m, deterministic)
}
func (m *ListClazzResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClazzResponse_List.Merge(m, src)
}
func (m *ListClazzResponse_List) XXX_Size() int {
	return xxx_messageInfo_ListClazzResponse_List.Size(m)
}
func (m *ListClazzResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClazzResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListClazzResponse_List proto.InternalMessageInfo

func (m *ListClazzResponse_List) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *ListClazzResponse_List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListClazzResponse_List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListClazzResponse_List) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *ListClazzResponse_List) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ListClazzResponse_List) GetFun() string {
	if m != nil {
		return m.Fun
	}
	return ""
}

func (m *ListClazzResponse_List) GetRequiredScript() bool {
	if m != nil {
		return m.RequiredScript
	}
	return false
}

func (m *ListClazzResponse_List) GetRequiredFields() []string {
	if m != nil {
		return m.RequiredFields
	}
	return nil
}

func (m *ListClazzResponse_List) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ListClazzResponse_List) GetModifer() string {
	if m != nil {
		return m.Modifer
	}
	return ""
}

func (m *ListClazzResponse_List) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *ListClazzResponse_List) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

//
//ListClazzApi返回
type ListClazzResponseWrapper struct {
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
	Data                 *ListClazzResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListClazzResponseWrapper) Reset()         { *m = ListClazzResponseWrapper{} }
func (m *ListClazzResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListClazzResponseWrapper) ProtoMessage()    {}
func (*ListClazzResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_7055e34279888799, []int{2}
}
func (m *ListClazzResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClazzResponseWrapper.Unmarshal(m, b)
}
func (m *ListClazzResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClazzResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListClazzResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClazzResponseWrapper.Merge(m, src)
}
func (m *ListClazzResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListClazzResponseWrapper.Size(m)
}
func (m *ListClazzResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClazzResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListClazzResponseWrapper proto.InternalMessageInfo

func (m *ListClazzResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListClazzResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListClazzResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListClazzResponseWrapper) GetData() *ListClazzResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListClazzRequest)(nil), "clazz.ListClazzRequest")
	proto.RegisterType((*ListClazzResponse)(nil), "clazz.ListClazzResponse")
	proto.RegisterType((*ListClazzResponse_List)(nil), "clazz.ListClazzResponse.List")
	proto.RegisterType((*ListClazzResponseWrapper)(nil), "clazz.ListClazzResponseWrapper")
}

func init() { proto.RegisterFile("list_clazz.proto", fileDescriptor_7055e34279888799) }

var fileDescriptor_7055e34279888799 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0x9b, 0x76, 0x8d, 0xbb, 0x8f, 0xce, 0x68, 0xc8, 0x54, 0x9a, 0x52, 0x19, 0x69,
	0xea, 0x24, 0xb6, 0x4a, 0x43, 0x7c, 0x08, 0x89, 0x8b, 0x15, 0x81, 0x54, 0x09, 0x71, 0xe1, 0x5d,
	0x70, 0x39, 0xa5, 0xb5, 0x1b, 0x2c, 0x92, 0x3a, 0x73, 0x5c, 0x86, 0x76, 0xc9, 0x73, 0xf0, 0x4a,
	0xbc, 0x42, 0x40, 0x3c, 0x42, 0x9e, 0x00, 0xf9, 0x38, 0x5d, 0xb3, 0x16, 0x84, 0xb8, 0x72, 0x72,
	0xfe, 0xbf, 0xff, 0xc9, 0x39, 0x3d, 0xc7, 0x45, 0xdd, 0x58, 0x66, 0xe6, 0x72, 0x1a, 0x87, 0x37,
	0x37, 0xa7, 0xa9, 0x56, 0x46, 0xe1, 0x26, 0xbc, 0xf4, 0x4e, 0x22, 0x69, 0x3e, 0x2e, 0x26, 0xa7,
	0x53, 0x95, 0x0c, 0x23, 0x15, 0xa9, 0x21, 0xa8, 0x93, 0xc5, 0x0c, 0xde, 0xe0, 0x05, 0x9e, 0x9c,
	0xab, 0xf7, 0xac, 0x82, 0x27, 0xd7, 0xd2, 0x7c, 0x52, 0xd7, 0xc3, 0x48, 0x9d, 0x80, 0x78, 0xf2,
	0x39, 0x8c, 0x25, 0x0f, 0x8d, 0xd2, 0xd9, 0xf0, 0xf6, 0xd1, 0xf9, 0xe8, 0xf7, 0x1a, 0xea, 0xbe,
	0x93, 0x99, 0x79, 0x6d, 0x3f, 0xca, 0xc4, 0xd5, 0x42, 0x64, 0x06, 0x1f, 0x23, 0x2f, 0x0d, 0x23,
	0x41, 0x6a, 0xfd, 0xda, 0xa0, 0x39, 0x3a, 0x28, 0xf2, 0xa0, 0x33, 0x53, 0x3a, 0x79, 0x49, 0x6d,
	0x94, 0xfe, 0xfa, 0x11, 0xd4, 0xbb, 0xf7, 0x18, 0x20, 0xf8, 0x29, 0x6a, 0xdb, 0xf3, 0x42, 0xde,
	0x08, 0x52, 0x07, 0xfc, 0x61, 0x91, 0x07, 0x7b, 0x2b, 0xdc, 0x2a, 0x4b, 0xcb, 0x2d, 0x8a, 0x87,
	0xa8, 0xcd, 0x65, 0x16, 0x4e, 0x62, 0xc1, 0x49, 0x03, 0x6c, 0xf7, 0x57, 0xb6, 0xa5, 0x42, 0xd9,
	0x2d, 0x84, 0x8f, 0x50, 0x53, 0x66, 0xe7, 0x71, 0x4c, 0x3c, 0xa0, 0xbb, 0x45, 0x1e, 0x6c, 0x3b,
	0x1a, 0xc2, 0x94, 0x39, 0x99, 0x7e, 0x6d, 0xa1, 0xfd, 0x4a, 0x3f, 0x59, 0xaa, 0xe6, 0x99, 0xf8,
	0x9f, 0x86, 0x9e, 0x23, 0xdf, 0x9e, 0x97, 0xd9, 0xaa, 0xa3, 0x5e, 0x91, 0x07, 0xdd, 0x15, 0x0f,
	0xd2, 0x66, 0x4b, 0x47, 0xa8, 0x69, 0x94, 0x09, 0xe3, 0xb2, 0x9f, 0x4a, 0x85, 0x10, 0xa6, 0xcc,
	0xc9, 0x78, 0x84, 0x3c, 0x3b, 0x73, 0xe2, 0xf5, 0x1b, 0x83, 0xce, 0xd9, 0xe1, 0xa9, 0x9b, 0xfd,
	0x46, 0xcd, 0x10, 0x19, 0xed, 0xad, 0x4a, 0xb5, 0x26, 0xca, 0xc0, 0xdb, 0xfb, 0xe6, 0x21, 0xcf,
	0xea, 0xb8, 0x8f, 0x1a, 0x4a, 0x47, 0x65, 0x5f, 0xbb, 0x45, 0x1e, 0x20, 0x07, 0x2b, 0x1d, 0x51,
	0x66, 0x25, 0x7c, 0x88, 0xea, 0x63, 0x0e, 0x8d, 0xf8, 0xa3, 0x9d, 0x22, 0x0f, 0x7c, 0x07, 0x8c,
	0x39, 0x65, 0xf5, 0x31, 0xc7, 0x8f, 0x90, 0xf7, 0x3e, 0x4c, 0x04, 0x14, 0xed, 0x57, 0x3f, 0x67,
	0xa3, 0x94, 0x81, 0x88, 0x8f, 0x51, 0x8b, 0x87, 0x26, 0x1c, 0x73, 0xf8, 0xf5, 0xfd, 0xd1, 0x7e,
	0x91, 0x07, 0x3b, 0xe5, 0xac, 0x20, 0x4e, 0x59, 0x09, 0xdc, 0x19, 0x6c, 0xb3, 0x5f, 0x1b, 0xb4,
	0xff, 0x35, 0xd8, 0x3e, 0x6a, 0xcc, 0x16, 0x73, 0xd2, 0x82, 0xc4, 0x95, 0x0e, 0x66, 0x8b, 0x39,
	0x65, 0x56, 0xc2, 0xe7, 0x68, 0x57, 0x8b, 0xab, 0x85, 0xd4, 0x82, 0x5f, 0x4c, 0xb5, 0x4c, 0x0d,
	0xd9, 0x82, 0xc4, 0x76, 0xd1, 0x0e, 0x1c, 0x7c, 0x57, 0xa7, 0x6c, 0xcd, 0x50, 0x4d, 0xf1, 0x56,
	0x8a, 0x98, 0x67, 0xa4, 0xdd, 0x6f, 0x0c, 0xfc, 0x3f, 0xa5, 0x70, 0x7a, 0x25, 0x85, 0x0b, 0xe0,
	0xc7, 0x68, 0x6b, 0xaa, 0x85, 0xbd, 0x39, 0xc4, 0x87, 0x5a, 0x71, 0x91, 0x07, 0xbb, 0xce, 0x5b,
	0x0a, 0x94, 0x2d, 0x11, 0x4b, 0x27, 0x8a, 0xcb, 0x99, 0xd0, 0x04, 0xad, 0xd3, 0xa5, 0x40, 0xd9,
	0x12, 0xb1, 0xab, 0x33, 0x35, 0x32, 0x11, 0xa4, 0xb3, 0xbe, 0x3a, 0x10, 0xa6, 0xcc, 0xc9, 0x96,
	0x4b, 0x80, 0xdb, 0x5e, 0xe7, 0x92, 0x92, 0x73, 0xe7, 0xcf, 0x1a, 0x22, 0x1b, 0x0b, 0xf5, 0x41,
	0x87, 0x69, 0x2a, 0xb4, 0x9d, 0xf8, 0x54, 0xf1, 0xe5, 0x5d, 0xa8, 0x4c, 0xdc, 0x46, 0x29, 0x03,
	0x11, 0xbf, 0x40, 0x1d, 0x7b, 0xbe, 0xf9, 0x92, 0xc6, 0xa1, 0x9c, 0x97, 0xeb, 0xf3, 0xa0, 0xc8,
	0x03, 0xbc, 0x62, 0x4b, 0x91, 0xb2, 0x2a, 0x6a, 0x6b, 0x14, 0x5a, 0x2b, 0x5d, 0x6e, 0x54, 0xa5,
	0x46, 0x08, 0x53, 0xe6, 0x64, 0xfc, 0x0a, 0x79, 0x76, 0x65, 0x60, 0xa3, 0x3a, 0x67, 0xe4, 0x6f,
	0xd7, 0xa0, 0x5a, 0xa0, 0xe5, 0x29, 0x03, 0xdb, 0xa4, 0x05, 0x7f, 0x5f, 0x4f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x9a, 0xcd, 0x3d, 0xed, 0x40, 0x05, 0x00, 0x00,
}
