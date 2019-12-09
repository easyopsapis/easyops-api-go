// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_result.proto

package cmdb

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
//ImportResult
type ImportResult struct {
	//
	//模型的实例ID列表
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//模型名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//分类, 点分表示二级分类
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category" form:"category"`
	//
	//说明
	Memo string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//true代表是核心模型，是受保护的模型，不允许被删除
	Protected bool `protobuf:"varint,5,opt,name=protected,proto3" json:"protected" form:"protected"`
	//
	//模型所属小产品，有值则会在界面隐藏该模型
	System string `protobuf:"bytes,6,opt,name=system,proto3" json:"system" form:"system"`
	//
	//状态码
	Code int32 `protobuf:"varint,7,opt,name=code,proto3" json:"code" form:"code"`
	//
	//导入信息
	Message string `protobuf:"bytes,8,opt,name=message,proto3" json:"message" form:"message"`
	//
	//基础信息导入结果
	InfoResult []*ImportStatus `protobuf:"bytes,9,rep,name=info_result,json=infoResult,proto3" json:"info_result" form:"info_result"`
	//
	//关系分组导入结果
	RelationGroupResult []*ImportStatus `protobuf:"bytes,10,rep,name=relation_group_result,json=relationGroupResult,proto3" json:"relation_group_result" form:"relation_group_result"`
	//
	//属性导入结果
	AttrListResult []*ImportStatus `protobuf:"bytes,11,rep,name=attr_list_result,json=attrListResult,proto3" json:"attr_list_result" form:"attr_list_result"`
	//
	//关系导入结果
	RelationListResult []*ImportStatus `protobuf:"bytes,12,rep,name=relation_list_result,json=relationListResult,proto3" json:"relation_list_result" form:"relation_list_result"`
	//
	//是否新建模型
	IsCreate             bool     `protobuf:"varint,13,opt,name=is_create,json=isCreate,proto3" json:"is_create" form:"is_create"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportResult) Reset()         { *m = ImportResult{} }
func (m *ImportResult) String() string { return proto.CompactTextString(m) }
func (*ImportResult) ProtoMessage()    {}
func (*ImportResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafdc51192a18066, []int{0}
}
func (m *ImportResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportResult.Unmarshal(m, b)
}
func (m *ImportResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportResult.Marshal(b, m, deterministic)
}
func (m *ImportResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportResult.Merge(m, src)
}
func (m *ImportResult) XXX_Size() int {
	return xxx_messageInfo_ImportResult.Size(m)
}
func (m *ImportResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportResult.DiscardUnknown(m)
}

var xxx_messageInfo_ImportResult proto.InternalMessageInfo

func (m *ImportResult) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ImportResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImportResult) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ImportResult) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ImportResult) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *ImportResult) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *ImportResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ImportResult) GetInfoResult() []*ImportStatus {
	if m != nil {
		return m.InfoResult
	}
	return nil
}

func (m *ImportResult) GetRelationGroupResult() []*ImportStatus {
	if m != nil {
		return m.RelationGroupResult
	}
	return nil
}

func (m *ImportResult) GetAttrListResult() []*ImportStatus {
	if m != nil {
		return m.AttrListResult
	}
	return nil
}

func (m *ImportResult) GetRelationListResult() []*ImportStatus {
	if m != nil {
		return m.RelationListResult
	}
	return nil
}

func (m *ImportResult) GetIsCreate() bool {
	if m != nil {
		return m.IsCreate
	}
	return false
}

func init() {
	proto.RegisterType((*ImportResult)(nil), "cmdb.ImportResult")
}

func init() { proto.RegisterFile("import_result.proto", fileDescriptor_cafdc51192a18066) }

var fileDescriptor_cafdc51192a18066 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0xb6, 0x75, 0xad, 0xbb, 0x7f, 0xb8, 0x03, 0xa2, 0x0d, 0xe4, 0xca, 0x70, 0x51,
	0xa4, 0x25, 0xd9, 0x1f, 0x69, 0x12, 0x5c, 0x20, 0x28, 0x17, 0x68, 0x82, 0x2b, 0x73, 0xc5, 0x26,
	0xa8, 0xdc, 0xc4, 0xcd, 0x02, 0x49, 0x4f, 0x65, 0xbb, 0x4c, 0x05, 0xf1, 0x64, 0xbc, 0x4b, 0x90,
	0x78, 0x84, 0x3c, 0x01, 0x8a, 0x9d, 0x64, 0xa5, 0xa2, 0x57, 0xdc, 0xd9, 0xdf, 0xf7, 0x9d, 0xdf,
	0x39, 0x96, 0x6d, 0xd4, 0x8d, 0xd3, 0x29, 0x48, 0x3d, 0x94, 0x42, 0xcd, 0x12, 0xed, 0x4d, 0x25,
	0x68, 0xc0, 0xeb, 0x41, 0x1a, 0x8e, 0x0e, 0xdc, 0x28, 0xd6, 0xd7, 0xb3, 0x91, 0x17, 0x40, 0xea,
	0x47, 0x10, 0x81, 0x6f, 0xcc, 0xd1, 0x6c, 0x6c, 0x76, 0x66, 0x63, 0x56, 0xb6, 0xe8, 0x80, 0x45,
	0xe0, 0x09, 0xae, 0xe6, 0x30, 0x55, 0x5e, 0x02, 0x01, 0x4f, 0xfc, 0x00, 0x26, 0x5a, 0xf2, 0x40,
	0x2b, 0x5b, 0x29, 0xc5, 0x14, 0xdc, 0x14, 0x42, 0x91, 0x28, 0xbf, 0x0c, 0xfa, 0x66, 0xeb, 0x17,
	0xed, 0xfc, 0x72, 0x10, 0xa5, 0xb9, 0x9e, 0xa9, 0x92, 0x79, 0xbe, 0x30, 0x42, 0x7a, 0x13, 0xeb,
	0x2f, 0x70, 0xe3, 0x47, 0xe0, 0x1a, 0xd3, 0xfd, 0xca, 0x93, 0x38, 0xe4, 0x1a, 0xa4, 0xf2, 0xeb,
	0xa5, 0xad, 0xa3, 0x3f, 0x9b, 0x68, 0xeb, 0xc2, 0xf0, 0x98, 0x39, 0x17, 0x66, 0xa8, 0x05, 0xa3,
	0xcf, 0x22, 0xd0, 0x17, 0xa1, 0xd3, 0xe8, 0x35, 0xfa, 0xed, 0xc1, 0x79, 0x9e, 0x91, 0xdd, 0x31,
	0xc8, 0xf4, 0x39, 0xad, 0x1c, 0xfa, 0xfb, 0x17, 0x21, 0xe8, 0xd1, 0xa7, 0x2b, 0xee, 0x7e, 0x7b,
	0xe5, 0x5e, 0x0e, 0x3f, 0x5e, 0x1d, 0xbb, 0xcf, 0xaa, 0xf5, 0xf7, 0xe3, 0xa3, 0xb3, 0x93, 0x1f,
	0x4f, 0x58, 0xcd, 0xc1, 0x8f, 0xd1, 0xfa, 0x84, 0xa7, 0xc2, 0xb9, 0x63, 0x78, 0xbb, 0x79, 0x46,
	0x3a, 0x96, 0x57, 0xa8, 0x94, 0x19, 0x13, 0xfb, 0xa8, 0x15, 0x70, 0x2d, 0x22, 0x90, 0x73, 0x67,
	0xcd, 0x04, 0xbb, 0xb7, 0x8d, 0x2b, 0x87, 0xb2, 0x3a, 0x54, 0x50, 0x53, 0x91, 0x82, 0xb3, 0xbe,
	0x4c, 0x2d, 0x54, 0xca, 0x8c, 0x89, 0x4f, 0x51, 0xbb, 0x38, 0xa8, 0x08, 0xb4, 0x08, 0x9d, 0x8d,
	0x5e, 0xa3, 0xdf, 0x1a, 0xec, 0xe7, 0x19, 0xd9, 0xb3, 0xc9, 0xda, 0xa2, 0xec, 0x36, 0x86, 0x9f,
	0xa2, 0xa6, 0x9a, 0x2b, 0x2d, 0x52, 0xa7, 0x69, 0xd0, 0x77, 0xf3, 0x8c, 0x6c, 0xdb, 0x02, 0xab,
	0x53, 0x56, 0x06, 0x8a, 0x19, 0x02, 0x08, 0x85, 0xb3, 0xd9, 0x6b, 0xf4, 0x37, 0x16, 0x67, 0x28,
	0x54, 0xca, 0x8c, 0x89, 0x8f, 0xd0, 0x66, 0x2a, 0x94, 0xe2, 0x91, 0x70, 0x5a, 0x06, 0x88, 0xf3,
	0x8c, 0xec, 0x54, 0xb3, 0x1a, 0x83, 0xb2, 0x2a, 0x82, 0xdf, 0xa2, 0x4e, 0x3c, 0x19, 0x43, 0xf9,
	0xce, 0x9c, 0x76, 0x6f, 0xad, 0xdf, 0x39, 0xc5, 0x5e, 0x71, 0xf3, 0x9e, 0xbd, 0xa9, 0xf7, 0xe6,
	0xe2, 0x07, 0xf7, 0xf3, 0x8c, 0x60, 0x4b, 0x59, 0x28, 0xa0, 0x0c, 0x15, 0xbb, 0xf2, 0x36, 0xaf,
	0xd1, 0x3d, 0x29, 0x12, 0xae, 0x63, 0x98, 0x0c, 0x23, 0x09, 0xb3, 0x69, 0x85, 0x45, 0x2b, 0xb1,
	0xbd, 0x3c, 0x23, 0x0f, 0x2d, 0xf6, 0x9f, 0xa5, 0x94, 0x75, 0x2b, 0xfd, 0x4d, 0x21, 0x97, 0x9d,
	0x3e, 0xa0, 0x3d, 0xae, 0xb5, 0x1c, 0x26, 0xb1, 0xaa, 0xfe, 0x88, 0xd3, 0x59, 0xd9, 0xe4, 0x30,
	0xcf, 0xc8, 0x03, 0xdb, 0x64, 0xb9, 0x8a, 0xb2, 0x9d, 0x42, 0x7a, 0x17, 0xab, 0xea, 0x49, 0x0a,
	0xb4, 0x5f, 0x4f, 0xb2, 0x88, 0xdf, 0x5a, 0x89, 0x27, 0x79, 0x46, 0x0e, 0x97, 0xce, 0xf0, 0x57,
	0x0b, 0x5c, 0xc9, 0x0b, 0x6d, 0x4e, 0x50, 0x3b, 0x56, 0xc3, 0x40, 0x0a, 0xae, 0x85, 0xb3, 0xbd,
	0xfc, 0x54, 0x6a, 0x8b, 0xb2, 0x56, 0xac, 0x5e, 0x9b, 0xe5, 0xe0, 0xe5, 0xe5, 0x8b, 0xff, 0xfb,
	0xcb, 0xa3, 0xa6, 0x09, 0x9d, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x45, 0xb2, 0xa1, 0x5e,
	0x04, 0x00, 0x00,
}
