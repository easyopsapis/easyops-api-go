// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: object_import.proto

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
//导入模型的message
type ObjectImport struct {
	//
	//ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//图标
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//分类, 点分表示二级分类
	Category string `protobuf:"bytes,4,opt,name=category,proto3" json:"category" form:"category"`
	//
	//说明
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//属性列表
	AttrList []*ObjectImport_AttrList `protobuf:"bytes,6,rep,name=attrList,proto3" json:"attrList" form:"attrList"`
	//
	//关系分组列表
	RelationGroups []*ObjectImport_RelationGroups `protobuf:"bytes,7,rep,name=relation_groups,json=relationGroups,proto3" json:"relation_groups" form:"relation_groups"`
	//
	//关系列表
	RelationList         []*ObjectImport_RelationList `protobuf:"bytes,8,rep,name=relation_list,json=relationList,proto3" json:"relation_list" form:"relation_list"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ObjectImport) Reset()         { *m = ObjectImport{} }
func (m *ObjectImport) String() string { return proto.CompactTextString(m) }
func (*ObjectImport) ProtoMessage()    {}
func (*ObjectImport) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0255ab5812aa029, []int{0}
}
func (m *ObjectImport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectImport.Unmarshal(m, b)
}
func (m *ObjectImport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectImport.Marshal(b, m, deterministic)
}
func (m *ObjectImport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectImport.Merge(m, src)
}
func (m *ObjectImport) XXX_Size() int {
	return xxx_messageInfo_ObjectImport.Size(m)
}
func (m *ObjectImport) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectImport.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectImport proto.InternalMessageInfo

func (m *ObjectImport) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ObjectImport) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectImport) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ObjectImport) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ObjectImport) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ObjectImport) GetAttrList() []*ObjectImport_AttrList {
	if m != nil {
		return m.AttrList
	}
	return nil
}

func (m *ObjectImport) GetRelationGroups() []*ObjectImport_RelationGroups {
	if m != nil {
		return m.RelationGroups
	}
	return nil
}

func (m *ObjectImport) GetRelationList() []*ObjectImport_RelationList {
	if m != nil {
		return m.RelationList
	}
	return nil
}

type ObjectImport_AttrList struct {
	//
	//属性ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//属性名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//'true' 唯一属性
	Unique string `protobuf:"bytes,3,opt,name=unique,proto3" json:"unique" form:"unique"`
	//
	//'true' 只读属性，'false' 可写属性
	Readonly string `protobuf:"bytes,4,opt,name=readonly,proto3" json:"readonly" form:"readonly"`
	//
	//'true' 必填属性，'false' 非必填属性
	Required string `protobuf:"bytes,5,opt,name=required,proto3" json:"required" form:"required"`
	//
	//属性分类
	Tag []string `protobuf:"bytes,6,rep,name=tag,proto3" json:"tag" form:"tag"`
	//
	//属性的描述
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description" form:"description"`
	//
	//属性的提示
	Tips string `protobuf:"bytes,8,opt,name=tips,proto3" json:"tips" form:"tips"`
	//
	//属性值类型，不同类型有不同字段
	Value                *ObjectAttrValue `protobuf:"bytes,9,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ObjectImport_AttrList) Reset()         { *m = ObjectImport_AttrList{} }
func (m *ObjectImport_AttrList) String() string { return proto.CompactTextString(m) }
func (*ObjectImport_AttrList) ProtoMessage()    {}
func (*ObjectImport_AttrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0255ab5812aa029, []int{0, 0}
}
func (m *ObjectImport_AttrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectImport_AttrList.Unmarshal(m, b)
}
func (m *ObjectImport_AttrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectImport_AttrList.Marshal(b, m, deterministic)
}
func (m *ObjectImport_AttrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectImport_AttrList.Merge(m, src)
}
func (m *ObjectImport_AttrList) XXX_Size() int {
	return xxx_messageInfo_ObjectImport_AttrList.Size(m)
}
func (m *ObjectImport_AttrList) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectImport_AttrList.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectImport_AttrList proto.InternalMessageInfo

func (m *ObjectImport_AttrList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectImport_AttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectImport_AttrList) GetUnique() string {
	if m != nil {
		return m.Unique
	}
	return ""
}

func (m *ObjectImport_AttrList) GetReadonly() string {
	if m != nil {
		return m.Readonly
	}
	return ""
}

func (m *ObjectImport_AttrList) GetRequired() string {
	if m != nil {
		return m.Required
	}
	return ""
}

func (m *ObjectImport_AttrList) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *ObjectImport_AttrList) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ObjectImport_AttrList) GetTips() string {
	if m != nil {
		return m.Tips
	}
	return ""
}

func (m *ObjectImport_AttrList) GetValue() *ObjectAttrValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ObjectImport_RelationGroups struct {
	//
	//分组ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//分组名称
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectImport_RelationGroups) Reset()         { *m = ObjectImport_RelationGroups{} }
func (m *ObjectImport_RelationGroups) String() string { return proto.CompactTextString(m) }
func (*ObjectImport_RelationGroups) ProtoMessage()    {}
func (*ObjectImport_RelationGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0255ab5812aa029, []int{0, 1}
}
func (m *ObjectImport_RelationGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectImport_RelationGroups.Unmarshal(m, b)
}
func (m *ObjectImport_RelationGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectImport_RelationGroups.Marshal(b, m, deterministic)
}
func (m *ObjectImport_RelationGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectImport_RelationGroups.Merge(m, src)
}
func (m *ObjectImport_RelationGroups) XXX_Size() int {
	return xxx_messageInfo_ObjectImport_RelationGroups.Size(m)
}
func (m *ObjectImport_RelationGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectImport_RelationGroups.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectImport_RelationGroups proto.InternalMessageInfo

func (m *ObjectImport_RelationGroups) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectImport_RelationGroups) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ObjectImport_RelationList struct {
	//
	//关系名称, 不是必填, 而且在引入了 left_description 和 right_description 之后这个字段几乎处于废弃状态
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//关系左端的模型ID
	LeftObjectId string `protobuf:"bytes,2,opt,name=left_object_id,json=leftObjectId,proto3" json:"left_object_id" form:"left_object_id"`
	//
	//关系左端模型中表达右端模型实例的别名字段: 比如应用的负责人需要在应用的实例中表达出一个字段; 对已有模型添加关系时这个ID需要加下划线前缀避免冲突
	LeftId string `protobuf:"bytes,3,opt,name=left_id,json=leftId,proto3" json:"left_id" form:"left_id"`
	//
	//是与 left_id 相反的含义, 但仅用于前端展示 (p.s.: 关系左端模型的描述)
	LeftDescription string `protobuf:"bytes,4,opt,name=left_description,json=leftDescription,proto3" json:"left_description" form:"left_description"`
	//
	//关系左端的资源模型实例至少包含多少数量的关系: 目前来说这个字段是过度设计的字段, 一般填0就好了
	LeftMin int32 `protobuf:"varint,5,opt,name=left_min,json=leftMin,proto3" json:"left_min" form:"left_min"`
	//
	//关系左端的资源模型实例最多包含多少数量的关系: 有数据库级别的索引支持, 一般情况填1 或者 -1, -1表示无限多
	LeftMax int32 `protobuf:"varint,6,opt,name=left_max,json=leftMax,proto3" json:"left_max" form:"left_max"`
	//
	//关系在左端模型的哪些分组里
	LeftGroups []string `protobuf:"bytes,7,rep,name=left_groups,json=leftGroups,proto3" json:"left_groups" form:"left_groups"`
	//
	//关系左端标签
	LeftTags []string `protobuf:"bytes,8,rep,name=left_tags,json=leftTags,proto3" json:"left_tags" form:"left_tags"`
	//
	//关系右端的模型ID
	RightObjectId string `protobuf:"bytes,9,opt,name=right_object_id,json=rightObjectId,proto3" json:"right_object_id" form:"right_object_id"`
	//
	//关系右端模型中表达左端模型实例的别名字段: 比如应用的负责人需要在应用的实例中表达出一个字段; 对已有模型添加关系时这个ID需要加下划线前缀避免冲突
	RightId string `protobuf:"bytes,10,opt,name=right_id,json=rightId,proto3" json:"right_id" form:"right_id"`
	//
	//是与 right_id 相反的含义, 但仅用于前端展示 (p.s.: 关系右端模型的描述)
	RightDescription string `protobuf:"bytes,11,opt,name=right_description,json=rightDescription,proto3" json:"right_description" form:"right_description"`
	//
	//关系右端的资源模型实例至少包含多少数量的关系: 目前来说这个字段是过度设计的字段, 一般填 0 就好了
	RightMin int32 `protobuf:"varint,12,opt,name=right_min,json=rightMin,proto3" json:"right_min" form:"right_min"`
	//
	//关系右端的资源模型实例最多包含多少数量的关系: 有数据库级别的索引支持, 一般情况填 1 或者 -1, -1表示无限多
	RightMax int32 `protobuf:"varint,13,opt,name=right_max,json=rightMax,proto3" json:"right_max" form:"right_max"`
	//
	//关系在右端模型的哪些分组里
	RightGroups []string `protobuf:"bytes,14,rep,name=right_groups,json=rightGroups,proto3" json:"right_groups" form:"right_groups"`
	//
	//关系右端标签
	RightTags            []string `protobuf:"bytes,15,rep,name=right_tags,json=rightTags,proto3" json:"right_tags" form:"right_tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectImport_RelationList) Reset()         { *m = ObjectImport_RelationList{} }
func (m *ObjectImport_RelationList) String() string { return proto.CompactTextString(m) }
func (*ObjectImport_RelationList) ProtoMessage()    {}
func (*ObjectImport_RelationList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0255ab5812aa029, []int{0, 2}
}
func (m *ObjectImport_RelationList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectImport_RelationList.Unmarshal(m, b)
}
func (m *ObjectImport_RelationList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectImport_RelationList.Marshal(b, m, deterministic)
}
func (m *ObjectImport_RelationList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectImport_RelationList.Merge(m, src)
}
func (m *ObjectImport_RelationList) XXX_Size() int {
	return xxx_messageInfo_ObjectImport_RelationList.Size(m)
}
func (m *ObjectImport_RelationList) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectImport_RelationList.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectImport_RelationList proto.InternalMessageInfo

func (m *ObjectImport_RelationList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectImport_RelationList) GetLeftObjectId() string {
	if m != nil {
		return m.LeftObjectId
	}
	return ""
}

func (m *ObjectImport_RelationList) GetLeftId() string {
	if m != nil {
		return m.LeftId
	}
	return ""
}

func (m *ObjectImport_RelationList) GetLeftDescription() string {
	if m != nil {
		return m.LeftDescription
	}
	return ""
}

func (m *ObjectImport_RelationList) GetLeftMin() int32 {
	if m != nil {
		return m.LeftMin
	}
	return 0
}

func (m *ObjectImport_RelationList) GetLeftMax() int32 {
	if m != nil {
		return m.LeftMax
	}
	return 0
}

func (m *ObjectImport_RelationList) GetLeftGroups() []string {
	if m != nil {
		return m.LeftGroups
	}
	return nil
}

func (m *ObjectImport_RelationList) GetLeftTags() []string {
	if m != nil {
		return m.LeftTags
	}
	return nil
}

func (m *ObjectImport_RelationList) GetRightObjectId() string {
	if m != nil {
		return m.RightObjectId
	}
	return ""
}

func (m *ObjectImport_RelationList) GetRightId() string {
	if m != nil {
		return m.RightId
	}
	return ""
}

func (m *ObjectImport_RelationList) GetRightDescription() string {
	if m != nil {
		return m.RightDescription
	}
	return ""
}

func (m *ObjectImport_RelationList) GetRightMin() int32 {
	if m != nil {
		return m.RightMin
	}
	return 0
}

func (m *ObjectImport_RelationList) GetRightMax() int32 {
	if m != nil {
		return m.RightMax
	}
	return 0
}

func (m *ObjectImport_RelationList) GetRightGroups() []string {
	if m != nil {
		return m.RightGroups
	}
	return nil
}

func (m *ObjectImport_RelationList) GetRightTags() []string {
	if m != nil {
		return m.RightTags
	}
	return nil
}

func init() {
	proto.RegisterType((*ObjectImport)(nil), "cmdb.ObjectImport")
	proto.RegisterType((*ObjectImport_AttrList)(nil), "cmdb.ObjectImport.AttrList")
	proto.RegisterType((*ObjectImport_RelationGroups)(nil), "cmdb.ObjectImport.RelationGroups")
	proto.RegisterType((*ObjectImport_RelationList)(nil), "cmdb.ObjectImport.RelationList")
}

func init() { proto.RegisterFile("object_import.proto", fileDescriptor_d0255ab5812aa029) }

var fileDescriptor_d0255ab5812aa029 = []byte{
	// 942 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xed, 0x8e, 0xdb, 0x44,
	0x14, 0x65, 0xbf, 0xf2, 0x71, 0x93, 0x4d, 0x76, 0x67, 0xbb, 0xad, 0x49, 0x41, 0x5e, 0xa6, 0xfc,
	0xd8, 0x4a, 0x24, 0xde, 0x6d, 0xa1, 0x40, 0x25, 0xa0, 0x8d, 0x10, 0x28, 0x52, 0xa1, 0x62, 0x84,
	0xfa, 0xa3, 0x15, 0x0d, 0x13, 0xdb, 0xeb, 0x0e, 0xd8, 0x99, 0xd4, 0x9e, 0xb4, 0x29, 0x88, 0x97,
	0xe0, 0x99, 0x78, 0x0c, 0x64, 0x24, 0x1e, 0xc1, 0x4f, 0x80, 0xe6, 0x8e, 0x9d, 0x9d, 0x44, 0x0b,
	0xda, 0x15, 0xff, 0x66, 0xee, 0x39, 0xe7, 0xce, 0xc7, 0x39, 0x9e, 0x04, 0x0e, 0xe4, 0xe4, 0xa7,
	0xd0, 0x57, 0x63, 0x91, 0xcc, 0x64, 0xaa, 0x06, 0xb3, 0x54, 0x2a, 0x49, 0xb6, 0xfd, 0x24, 0x98,
	0xf4, 0xfa, 0x91, 0x50, 0x2f, 0xe6, 0x93, 0x81, 0x2f, 0x13, 0x2f, 0x92, 0x91, 0xf4, 0x10, 0x9c,
	0xcc, 0xcf, 0x70, 0x86, 0x13, 0x1c, 0x19, 0x51, 0xef, 0x49, 0x24, 0x07, 0x21, 0xcf, 0xde, 0xc8,
	0x59, 0x36, 0x88, 0xa5, 0xcf, 0x63, 0xcf, 0x97, 0x53, 0x95, 0x72, 0x5f, 0x65, 0x46, 0x99, 0x86,
	0x33, 0xd9, 0x4f, 0x64, 0x10, 0xc6, 0x99, 0x57, 0x12, 0x3d, 0x9c, 0x7a, 0x7a, 0x39, 0xaf, 0xdc,
	0x08, 0x57, 0x2a, 0x1d, 0xbf, 0xe2, 0xf1, 0x3c, 0x2c, 0xfb, 0xde, 0xb3, 0xb6, 0x91, 0xbc, 0x16,
	0xea, 0x67, 0xf9, 0xda, 0x8b, 0x64, 0x1f, 0xc1, 0xfe, 0x2b, 0x1e, 0x8b, 0x80, 0x2b, 0x99, 0x66,
	0xde, 0x72, 0x68, 0x74, 0xf4, 0xcf, 0x2e, 0xb4, 0x1f, 0x63, 0xcf, 0x11, 0x9e, 0x8d, 0x30, 0x68,
	0x98, 0x35, 0x46, 0x81, 0xb3, 0x71, 0xb4, 0x71, 0xdc, 0x1c, 0xde, 0x2b, 0x72, 0xb7, 0x7b, 0x26,
	0xd3, 0xe4, 0x3e, 0xad, 0x10, 0xfa, 0xf7, 0x5f, 0xae, 0x0b, 0xef, 0x3e, 0x7f, 0xc6, 0xfb, 0xbf,
	0x3c, 0xec, 0x3f, 0x1d, 0xff, 0xf0, 0xec, 0xa4, 0xff, 0x69, 0x35, 0xfe, 0xf5, 0xe4, 0x83, 0xbb,
	0xa7, 0xbf, 0xbd, 0xcf, 0x96, 0x7d, 0xc8, 0x6d, 0xd8, 0x9e, 0xf2, 0x24, 0x74, 0x36, 0xb1, 0xdf,
	0x61, 0x91, 0xbb, 0x2d, 0xd3, 0x4f, 0x57, 0x75, 0xaf, 0xcd, 0xd9, 0x5b, 0x0c, 0x29, 0xe4, 0x16,
	0x6c, 0x0b, 0x5f, 0x4e, 0x9d, 0x2d, 0xa4, 0x76, 0xcf, 0xa9, 0xba, 0x4a, 0x19, 0x82, 0xc4, 0x83,
	0x86, 0xcf, 0x55, 0x18, 0xc9, 0xf4, 0x8d, 0xb3, 0x8d, 0xc4, 0x83, 0xf3, 0x3d, 0x56, 0x08, 0x65,
	0x4b, 0x92, 0xee, 0x9a, 0x84, 0x89, 0x74, 0x76, 0xd6, 0xbb, 0xea, 0x2a, 0x65, 0x08, 0x92, 0x47,
	0xd0, 0xd0, 0xd7, 0xfa, 0x48, 0x64, 0xca, 0xa9, 0x1d, 0x6d, 0x1d, 0xb7, 0xee, 0xdc, 0x1c, 0xe8,
	0x3b, 0x1f, 0xd8, 0xf7, 0x33, 0x78, 0x58, 0x52, 0xec, 0x25, 0x2b, 0x19, 0x65, 0xcb, 0x0e, 0xe4,
	0x0c, 0xba, 0x69, 0x18, 0x73, 0x25, 0xe4, 0x74, 0x1c, 0xa5, 0x72, 0x3e, 0xcb, 0x9c, 0x3a, 0x36,
	0x7d, 0xef, 0x82, 0xa6, 0xac, 0x64, 0x7e, 0x8d, 0xc4, 0x61, 0xaf, 0xc8, 0xdd, 0xeb, 0xa6, 0xf5,
	0x5a, 0x0f, 0xca, 0x3a, 0xe9, 0x0a, 0x97, 0x3c, 0x87, 0xdd, 0x25, 0x27, 0xd6, 0x5b, 0x6f, 0xe0,
	0x2a, 0xee, 0x7f, 0xac, 0x82, 0xdb, 0x77, 0x8a, 0xdc, 0xbd, 0xb6, 0xb6, 0x46, 0x8c, 0x67, 0x68,
	0xa7, 0x16, 0xaf, 0xf7, 0xc7, 0x16, 0x34, 0xaa, 0x33, 0x93, 0x07, 0xb0, 0x29, 0xaa, 0x58, 0x9c,
	0x14, 0xb9, 0xdb, 0x2c, 0xbd, 0xb9, 0x5c, 0x20, 0x36, 0xc5, 0x95, 0xa2, 0x70, 0x1b, 0x6a, 0xf3,
	0xa9, 0x78, 0x39, 0x0f, 0xcb, 0x30, 0xec, 0x17, 0xb9, 0xbb, 0x6b, 0xc8, 0xa6, 0x4e, 0x59, 0x49,
	0xd0, 0x81, 0x48, 0x43, 0x1e, 0xc8, 0x69, 0x7c, 0x41, 0x20, 0x2a, 0x84, 0xb2, 0x25, 0xc9, 0x08,
	0x5e, 0xce, 0x45, 0x1a, 0x06, 0x65, 0x28, 0x56, 0x04, 0x06, 0x41, 0x81, 0x19, 0x92, 0x23, 0xd8,
	0x52, 0x3c, 0xc2, 0x5c, 0x34, 0x87, 0x9d, 0x22, 0x77, 0xc1, 0x70, 0x15, 0x8f, 0x28, 0xd3, 0x10,
	0xf9, 0x04, 0x5a, 0x41, 0x98, 0xf9, 0xa9, 0x98, 0xe9, 0xbb, 0x73, 0xea, 0xd8, 0xf5, 0x7a, 0x91,
	0xbb, 0xc4, 0x30, 0x2d, 0x90, 0x32, 0x9b, 0xaa, 0xd3, 0xa9, 0xc4, 0x2c, 0x73, 0x1a, 0xeb, 0xe9,
	0xd4, 0x55, 0xca, 0x10, 0x24, 0x9f, 0xc1, 0x0e, 0x7e, 0xef, 0x4e, 0xf3, 0x68, 0xe3, 0xb8, 0x75,
	0xe7, 0xd0, 0xf6, 0x57, 0xfb, 0xf3, 0x44, 0x83, 0xc3, 0xbd, 0x22, 0x77, 0xdb, 0x46, 0x8c, 0x6c,
	0xca, 0x8c, 0xaa, 0xf7, 0x23, 0x74, 0x56, 0x43, 0x46, 0x6e, 0x59, 0x5e, 0x1e, 0xac, 0x7b, 0xa9,
	0x5d, 0xb8, 0x9a, 0x5d, 0xbd, 0xdf, 0xeb, 0xd0, 0xb6, 0x13, 0xa6, 0x8f, 0x85, 0xda, 0x8d, 0xf5,
	0x63, 0xa1, 0xb6, 0x34, 0xf9, 0x0b, 0xe8, 0xc4, 0xe1, 0x99, 0x1a, 0x57, 0x0f, 0x6c, 0x50, 0x2e,
	0xf5, 0x76, 0x91, 0xbb, 0x87, 0x86, 0xbe, 0x8a, 0x53, 0xd6, 0xd6, 0x85, 0xc7, 0xd5, 0xdb, 0xf2,
	0x2d, 0xd4, 0x91, 0x20, 0x82, 0x32, 0x26, 0x1f, 0x15, 0xb9, 0xdb, 0xb1, 0x94, 0x97, 0x0c, 0x67,
	0x4d, 0x93, 0x47, 0x01, 0xf9, 0x0a, 0xf6, 0x50, 0x66, 0x7b, 0x69, 0x22, 0x75, 0xb3, 0xc8, 0xdd,
	0x1b, 0x56, 0xe3, 0x15, 0x43, 0xbb, 0xba, 0xf4, 0xa5, 0x65, 0xea, 0x00, 0x1a, 0xc8, 0x4a, 0xc4,
	0x14, 0x13, 0xb6, 0x63, 0x27, 0xac, 0x42, 0x28, 0xc3, 0xcd, 0x7f, 0x23, 0x2c, 0x3e, 0x5f, 0x38,
	0xb5, 0x8b, 0xf9, 0x7c, 0x51, 0xf1, 0xf9, 0x82, 0x7c, 0x0c, 0x2d, 0xac, 0x5a, 0x6f, 0xcb, 0x4a,
	0xdc, 0x2c, 0x90, 0x32, 0xd0, 0xb3, 0xd2, 0xf7, 0x53, 0x68, 0x22, 0xa6, 0x78, 0x94, 0xe1, 0x63,
	0xd1, 0x1c, 0x5e, 0x2b, 0x72, 0x77, 0xcf, 0x92, 0x69, 0x88, 0x32, 0xdc, 0xcf, 0xf7, 0x3c, 0xca,
	0xc8, 0x10, 0xba, 0xa9, 0x88, 0x5e, 0xd8, 0x2e, 0x35, 0xf1, 0x4a, 0xec, 0x87, 0x6a, 0x95, 0x40,
	0xd9, 0x2e, 0x56, 0x96, 0x3e, 0x7d, 0x07, 0x0d, 0x43, 0x11, 0x81, 0x03, 0xeb, 0xbf, 0x2b, 0x15,
	0x72, 0x29, 0xa7, 0xea, 0xc8, 0x1e, 0x05, 0x64, 0x04, 0xfb, 0x46, 0x68, 0x7b, 0xd5, 0xc2, 0xde,
	0xef, 0x14, 0xb9, 0xeb, 0xd8, 0xbd, 0x57, 0xcc, 0xda, 0xc3, 0x9a, 0xed, 0xd6, 0x29, 0x34, 0x0d,
	0x4f, 0xdb, 0xd5, 0xc6, 0xeb, 0xb7, 0x2e, 0x65, 0x09, 0xe9, 0x17, 0x41, 0x8f, 0xb5, 0x61, 0xe7,
	0x12, 0xbe, 0x70, 0x76, 0xff, 0x45, 0xa2, 0x2d, 0x2b, 0x25, 0x7c, 0x41, 0xee, 0x43, 0xdb, 0xd4,
	0x4b, 0xd3, 0x3a, 0x78, 0xfb, 0x37, 0x8a, 0xdc, 0x3d, 0xb0, 0x55, 0x95, 0x6b, 0x2d, 0x9c, 0x96,
	0xb6, 0x7d, 0x08, 0x60, 0x50, 0xf4, 0xad, 0x8b, 0x4a, 0xfd, 0x3d, 0xee, 0xdb, 0x4a, 0x63, 0x9c,
	0xd9, 0x97, 0x76, 0x6e, 0xf8, 0xe0, 0xe9, 0xe7, 0xff, 0xef, 0x0f, 0xc7, 0xa4, 0x86, 0xa4, 0xbb,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x47, 0x05, 0xf3, 0x0e, 0x03, 0x09, 0x00, 0x00,
}
