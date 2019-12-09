// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_app_in_system.proto

package business_instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//SearchAppInSystem请求
type SearchAppInSystemRequest struct {
	//
	//系统的实例ID
	SystemInstanceId string `protobuf:"bytes,1,opt,name=systemInstanceId,proto3" json:"systemInstanceId" form:"systemInstanceId"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//查询条件,与[实例搜索接口]的query字段说明一致
	Query *types.Struct `protobuf:"bytes,4,opt,name=query,proto3" json:"query" form:"query"`
	//
	//过滤的字段列表, 留空代表返回所有字段(true: 表示指定获取字段, false: 表示指定不获取的字段)(支持关系数据的二级jsonPath格式的指定字段如clusters.name)
	Fields *types.Struct `protobuf:"bytes,5,opt,name=fields,proto3" json:"fields" form:"fields"`
	//
	//按字段排序, 留空默认按照实例ID降序排序(1表示升序, -1表示降序)
	Sort *types.Struct `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort" form:"sort"`
	//
	//按照权限过滤(通用实例都有`read`, `update`, `delete`权限控制, 主机实例在通用实例权限基础上有额外的`operate`权限, 应用实例在通用实例权限基础上有额外的`developClusterOperate`, `testClusterOperate`, `prereleaseClusterOperate`, `productionClusterOperate`权限)
	Permission []string `protobuf:"bytes,7,rep,name=permission,proto3" json:"permission" form:"permission"`
	//
	//对于关联的实例数据是否只获取relation_view中指定的属性, 这个字段为true时, 会覆盖fields字段中指定的二级字段设置
	OnlyRelationView bool `protobuf:"varint,8,opt,name=only_relation_view,json=onlyRelationView,proto3" json:"only_relation_view" form:"only_relation_view"`
	//
	//是否只获取与自己有关的那部分数据, 默认为false
	OnlyMyInstance bool `protobuf:"varint,9,opt,name=only_my_instance,json=onlyMyInstance,proto3" json:"only_my_instance" form:"only_my_instance"`
	//
	//是否包含子系统以及子系统的子系统的应用
	IncludeSubSystems    bool     `protobuf:"varint,10,opt,name=include_sub_systems,json=includeSubSystems,proto3" json:"include_sub_systems" form:"include_sub_systems"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAppInSystemRequest) Reset()         { *m = SearchAppInSystemRequest{} }
func (m *SearchAppInSystemRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAppInSystemRequest) ProtoMessage()    {}
func (*SearchAppInSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22482794e0518a21, []int{0}
}
func (m *SearchAppInSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAppInSystemRequest.Unmarshal(m, b)
}
func (m *SearchAppInSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAppInSystemRequest.Marshal(b, m, deterministic)
}
func (m *SearchAppInSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAppInSystemRequest.Merge(m, src)
}
func (m *SearchAppInSystemRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAppInSystemRequest.Size(m)
}
func (m *SearchAppInSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAppInSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAppInSystemRequest proto.InternalMessageInfo

func (m *SearchAppInSystemRequest) GetSystemInstanceId() string {
	if m != nil {
		return m.SystemInstanceId
	}
	return ""
}

func (m *SearchAppInSystemRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchAppInSystemRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchAppInSystemRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *SearchAppInSystemRequest) GetFields() *types.Struct {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchAppInSystemRequest) GetSort() *types.Struct {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *SearchAppInSystemRequest) GetPermission() []string {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *SearchAppInSystemRequest) GetOnlyRelationView() bool {
	if m != nil {
		return m.OnlyRelationView
	}
	return false
}

func (m *SearchAppInSystemRequest) GetOnlyMyInstance() bool {
	if m != nil {
		return m.OnlyMyInstance
	}
	return false
}

func (m *SearchAppInSystemRequest) GetIncludeSubSystems() bool {
	if m != nil {
		return m.IncludeSubSystems
	}
	return false
}

//
//SearchAppInSystem返回
type SearchAppInSystemResponse struct {
	//
	//total
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total"`
	//
	//page
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//page_size
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//实例列表
	List                 []*types.Struct `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SearchAppInSystemResponse) Reset()         { *m = SearchAppInSystemResponse{} }
func (m *SearchAppInSystemResponse) String() string { return proto.CompactTextString(m) }
func (*SearchAppInSystemResponse) ProtoMessage()    {}
func (*SearchAppInSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22482794e0518a21, []int{1}
}
func (m *SearchAppInSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAppInSystemResponse.Unmarshal(m, b)
}
func (m *SearchAppInSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAppInSystemResponse.Marshal(b, m, deterministic)
}
func (m *SearchAppInSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAppInSystemResponse.Merge(m, src)
}
func (m *SearchAppInSystemResponse) XXX_Size() int {
	return xxx_messageInfo_SearchAppInSystemResponse.Size(m)
}
func (m *SearchAppInSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAppInSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAppInSystemResponse proto.InternalMessageInfo

func (m *SearchAppInSystemResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SearchAppInSystemResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchAppInSystemResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchAppInSystemResponse) GetList() []*types.Struct {
	if m != nil {
		return m.List
	}
	return nil
}

//
//SearchAppInSystemApi返回
type SearchAppInSystemResponseWrapper struct {
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
	Data                 *SearchAppInSystemResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SearchAppInSystemResponseWrapper) Reset()         { *m = SearchAppInSystemResponseWrapper{} }
func (m *SearchAppInSystemResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*SearchAppInSystemResponseWrapper) ProtoMessage()    {}
func (*SearchAppInSystemResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_22482794e0518a21, []int{2}
}
func (m *SearchAppInSystemResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAppInSystemResponseWrapper.Unmarshal(m, b)
}
func (m *SearchAppInSystemResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAppInSystemResponseWrapper.Marshal(b, m, deterministic)
}
func (m *SearchAppInSystemResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAppInSystemResponseWrapper.Merge(m, src)
}
func (m *SearchAppInSystemResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_SearchAppInSystemResponseWrapper.Size(m)
}
func (m *SearchAppInSystemResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAppInSystemResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAppInSystemResponseWrapper proto.InternalMessageInfo

func (m *SearchAppInSystemResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchAppInSystemResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *SearchAppInSystemResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SearchAppInSystemResponseWrapper) GetData() *SearchAppInSystemResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchAppInSystemRequest)(nil), "business_instance.SearchAppInSystemRequest")
	proto.RegisterType((*SearchAppInSystemResponse)(nil), "business_instance.SearchAppInSystemResponse")
	proto.RegisterType((*SearchAppInSystemResponseWrapper)(nil), "business_instance.SearchAppInSystemResponseWrapper")
}

func init() { proto.RegisterFile("search_app_in_system.proto", fileDescriptor_22482794e0518a21) }

var fileDescriptor_22482794e0518a21 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdf, 0x6b, 0xd3, 0x40,
	0x1c, 0x37, 0x5b, 0x3a, 0xd7, 0x9b, 0xce, 0xf6, 0x44, 0x97, 0xd5, 0x1f, 0x09, 0xa7, 0x48, 0x05,
	0x9b, 0xaa, 0xc3, 0x9f, 0x0c, 0xc4, 0xc2, 0x1e, 0x86, 0x28, 0x78, 0x05, 0x7d, 0x18, 0x1a, 0xd2,
	0xf6, 0x96, 0x1d, 0xa6, 0xb9, 0xec, 0xee, 0xb2, 0xda, 0x89, 0x7f, 0x98, 0xff, 0x4c, 0x04, 0xc1,
	0x7f, 0x20, 0xcf, 0x3e, 0x48, 0xbe, 0xc9, 0xd6, 0xb0, 0x6e, 0xc3, 0x27, 0x9f, 0x92, 0x7c, 0x7e,
	0x71, 0xdf, 0x4f, 0xbe, 0x09, 0x6a, 0x29, 0xe6, 0xcb, 0xe1, 0x9e, 0xe7, 0xc7, 0xb1, 0xc7, 0x23,
	0x4f, 0x4d, 0x95, 0x66, 0x63, 0x37, 0x96, 0x42, 0x0b, 0xdc, 0x1c, 0x24, 0x8a, 0x47, 0x4c, 0x29,
	0x8f, 0x47, 0x4a, 0xfb, 0xd1, 0x90, 0xb5, 0x3a, 0x01, 0xd7, 0x7b, 0xc9, 0xc0, 0x1d, 0x8a, 0x71,
	0x37, 0x10, 0x81, 0xe8, 0x82, 0x72, 0x90, 0xec, 0xc2, 0x13, 0x3c, 0xc0, 0x5d, 0x91, 0xd0, 0x7a,
	0x5a, 0x91, 0x8f, 0x27, 0x5c, 0x7f, 0x11, 0x93, 0x6e, 0x20, 0x3a, 0x40, 0x76, 0x0e, 0xfc, 0x90,
	0x8f, 0x7c, 0x2d, 0xa4, 0xea, 0x1e, 0xdf, 0x96, 0xbe, 0x9b, 0x81, 0x10, 0x41, 0xc8, 0x66, 0xe9,
	0x4a, 0xcb, 0x64, 0xa8, 0x0b, 0x96, 0xfc, 0xa8, 0x21, 0xab, 0x0f, 0xc7, 0x7e, 0x1d, 0xc7, 0xdb,
	0x51, 0x1f, 0xce, 0x4c, 0xd9, 0x7e, 0xc2, 0x94, 0xc6, 0x3b, 0xa8, 0x51, 0x0c, 0xb1, 0x5d, 0x9e,
	0x79, 0x7b, 0x64, 0x19, 0x8e, 0xd1, 0xae, 0xf7, 0xba, 0x59, 0x6a, 0xaf, 0xed, 0x0a, 0x39, 0x7e,
	0x49, 0x4e, 0x2a, 0xc8, 0xaf, 0x9f, 0x76, 0x03, 0xad, 0x7e, 0xde, 0x79, 0xd8, 0x79, 0xe1, 0x77,
	0x0e, 0x3f, 0x7d, 0x7b, 0xb4, 0xf1, 0xfd, 0x2e, 0x9d, 0x0b, 0xc2, 0xf7, 0x91, 0x19, 0xfb, 0x01,
	0xb3, 0x16, 0x1c, 0xa3, 0x5d, 0xeb, 0x5d, 0xcb, 0x52, 0x7b, 0xa5, 0x08, 0xcc, 0xd1, 0x3c, 0x64,
	0xa1, 0x71, 0x81, 0x82, 0x04, 0x3f, 0x43, 0xf5, 0xfc, 0xea, 0x29, 0x7e, 0xc8, 0xac, 0x45, 0xd0,
	0xb7, 0xb2, 0xd4, 0x6e, 0xcc, 0xf4, 0x40, 0x1d, 0x99, 0x96, 0x73, 0xa4, 0xcf, 0x0f, 0x19, 0x7e,
	0x85, 0x6a, 0xfb, 0x09, 0x93, 0x53, 0xcb, 0x74, 0x8c, 0xf6, 0xca, 0xe3, 0x35, 0xb7, 0xe8, 0xc2,
	0x3d, 0xea, 0xc2, 0xed, 0x43, 0x17, 0xbd, 0x46, 0x96, 0xda, 0x97, 0x8a, 0x34, 0xd0, 0x13, 0x5a,
	0xf8, 0x70, 0x0f, 0x2d, 0xed, 0x72, 0x16, 0x8e, 0x94, 0x55, 0x3b, 0x3f, 0xa1, 0x99, 0xa5, 0xf6,
	0xe5, 0x22, 0xa1, 0x30, 0x10, 0x5a, 0x3a, 0xf1, 0x26, 0x32, 0x95, 0x90, 0xda, 0x5a, 0x3a, 0x3f,
	0xe1, 0xca, 0xac, 0x81, 0x5c, 0x4e, 0x28, 0xb8, 0xf0, 0x13, 0x84, 0x62, 0x26, 0xc7, 0x5c, 0x29,
	0x2e, 0x22, 0xeb, 0xa2, 0xb3, 0xd8, 0xae, 0x43, 0x59, 0xcd, 0x72, 0xf8, 0x63, 0x8e, 0xd0, 0x8a,
	0x10, 0xbf, 0x41, 0x58, 0x44, 0xe1, 0xd4, 0x93, 0x2c, 0xf4, 0x35, 0x17, 0x91, 0x77, 0xc0, 0xd9,
	0xc4, 0x5a, 0x76, 0x8c, 0xf6, 0x72, 0xef, 0x56, 0x96, 0xda, 0xeb, 0x85, 0x7d, 0x5e, 0x43, 0x68,
	0x23, 0x07, 0x69, 0x89, 0x7d, 0xe0, 0x6c, 0x82, 0xb7, 0x10, 0x60, 0xde, 0x78, 0x7a, 0xbc, 0xbd,
	0x56, 0x1d, 0xa2, 0x6e, 0xcc, 0xf6, 0xe0, 0xa4, 0x82, 0xd0, 0xd5, 0x1c, 0x7a, 0x3b, 0x3d, 0x7a,
	0xe7, 0xf8, 0x1d, 0xba, 0xca, 0xa3, 0x61, 0x98, 0x8c, 0x98, 0xa7, 0x92, 0x41, 0xf9, 0x7d, 0x28,
	0x0b, 0x41, 0xd2, 0xed, 0x2c, 0xb5, 0x5b, 0x45, 0xd2, 0x29, 0x22, 0x42, 0x9b, 0x25, 0xda, 0x4f,
	0x06, 0xfd, 0x12, 0xfb, 0x6d, 0xa0, 0xf5, 0x53, 0x76, 0x57, 0xc5, 0x22, 0x52, 0x0c, 0xdf, 0x43,
	0x35, 0x2d, 0xb4, 0x1f, 0xc2, 0xc6, 0xd6, 0xaa, 0xaf, 0x18, 0x60, 0x42, 0x0b, 0xfa, 0xbf, 0xec,
	0xe1, 0x26, 0x32, 0x43, 0xae, 0xb4, 0x65, 0x3a, 0x8b, 0xff, 0xb8, 0x02, 0xb9, 0x9c, 0x50, 0x70,
	0x91, 0x3f, 0x06, 0x72, 0xce, 0x9c, 0xf3, 0xa3, 0xf4, 0xe3, 0x98, 0x49, 0x7c, 0x07, 0x99, 0x43,
	0x31, 0x62, 0xe5, 0xb4, 0x95, 0xa4, 0x1c, 0x25, 0x14, 0x48, 0xfc, 0x1c, 0xad, 0xe4, 0xd7, 0xad,
	0xaf, 0x71, 0xe8, 0xf3, 0x08, 0x46, 0xae, 0xf7, 0xae, 0x67, 0xa9, 0x8d, 0x67, 0xda, 0x92, 0x24,
	0xb4, 0x2a, 0xcd, 0xdb, 0x64, 0x52, 0x0a, 0x09, 0x63, 0xd7, 0xab, 0x6d, 0x02, 0x4c, 0x68, 0x41,
	0xe3, 0xf7, 0xc8, 0x1c, 0xf9, 0xda, 0x2f, 0x3f, 0xb8, 0x07, 0xee, 0xdc, 0x6f, 0xcf, 0x3d, 0x73,
	0x92, 0xea, 0xa1, 0xf3, 0x0c, 0x42, 0x21, 0x6a, 0xb0, 0x04, 0x35, 0x6d, 0xfc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x3d, 0x6b, 0x16, 0x5f, 0x05, 0x00, 0x00,
}
