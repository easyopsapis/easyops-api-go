// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: view.proto

package topology

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//视图
type View struct {
	//
	//ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//创建者
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//修改者
	Modifier string `protobuf:"bytes,4,opt,name=modifier,proto3" json:"modifier" form:"modifier"`
	//
	//允许可读用户
	ReadAuthorizers []string `protobuf:"bytes,5,rep,name=readAuthorizers,proto3" json:"readAuthorizers" form:"readAuthorizers"`
	//
	//允许可写用户
	WriteAuthorizers []string `protobuf:"bytes,6,rep,name=writeAuthorizers,proto3" json:"writeAuthorizers" form:"writeAuthorizers"`
	//
	//版本
	Version string `protobuf:"bytes,7,opt,name=version,proto3" json:"version" form:"version"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,8,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime int32 `protobuf:"varint,9,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//根节点
	RootNode *Node `protobuf:"bytes,10,opt,name=rootNode,proto3" json:"rootNode" form:"rootNode"`
	//
	//节点
	Nodes []*Node `protobuf:"bytes,11,rep,name=nodes,proto3" json:"nodes" form:"nodes"`
	//
	//线
	Links []*Link `protobuf:"bytes,12,rep,name=links,proto3" json:"links" form:"links"`
	//
	//区域
	Areas []*Area `protobuf:"bytes,13,rep,name=areas,proto3" json:"areas" form:"areas"`
	//
	//标注
	Notes []*Note `protobuf:"bytes,14,rep,name=notes,proto3" json:"notes" form:"notes"`
	//
	//差异
	Diff                 *View_Diff `protobuf:"bytes,15,opt,name=diff,proto3" json:"diff" form:"diff"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *View) Reset()         { *m = View{} }
func (m *View) String() string { return proto.CompactTextString(m) }
func (*View) ProtoMessage()    {}
func (*View) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c1b2aca93c333f, []int{0}
}
func (m *View) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_View.Unmarshal(m, b)
}
func (m *View) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_View.Marshal(b, m, deterministic)
}
func (m *View) XXX_Merge(src proto.Message) {
	xxx_messageInfo_View.Merge(m, src)
}
func (m *View) XXX_Size() int {
	return xxx_messageInfo_View.Size(m)
}
func (m *View) XXX_DiscardUnknown() {
	xxx_messageInfo_View.DiscardUnknown(m)
}

var xxx_messageInfo_View proto.InternalMessageInfo

func (m *View) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *View) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *View) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *View) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

func (m *View) GetReadAuthorizers() []string {
	if m != nil {
		return m.ReadAuthorizers
	}
	return nil
}

func (m *View) GetWriteAuthorizers() []string {
	if m != nil {
		return m.WriteAuthorizers
	}
	return nil
}

func (m *View) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *View) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *View) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *View) GetRootNode() *Node {
	if m != nil {
		return m.RootNode
	}
	return nil
}

func (m *View) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *View) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *View) GetAreas() []*Area {
	if m != nil {
		return m.Areas
	}
	return nil
}

func (m *View) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *View) GetDiff() *View_Diff {
	if m != nil {
		return m.Diff
	}
	return nil
}

type View_Diff struct {
	//
	//新增节点
	AddNodes []*Node `protobuf:"bytes,1,rep,name=addNodes,proto3" json:"addNodes" form:"addNodes"`
	//
	//已删除节点
	RemoveNodes []*Node `protobuf:"bytes,2,rep,name=removeNodes,proto3" json:"removeNodes" form:"removeNodes"`
	//
	//新增线
	AddLinks []*Link `protobuf:"bytes,3,rep,name=addLinks,proto3" json:"addLinks" form:"addLinks"`
	//
	//已删除线
	RemoveLinks          []*Link  `protobuf:"bytes,4,rep,name=removeLinks,proto3" json:"removeLinks" form:"removeLinks"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *View_Diff) Reset()         { *m = View_Diff{} }
func (m *View_Diff) String() string { return proto.CompactTextString(m) }
func (*View_Diff) ProtoMessage()    {}
func (*View_Diff) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c1b2aca93c333f, []int{0, 0}
}
func (m *View_Diff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_View_Diff.Unmarshal(m, b)
}
func (m *View_Diff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_View_Diff.Marshal(b, m, deterministic)
}
func (m *View_Diff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_View_Diff.Merge(m, src)
}
func (m *View_Diff) XXX_Size() int {
	return xxx_messageInfo_View_Diff.Size(m)
}
func (m *View_Diff) XXX_DiscardUnknown() {
	xxx_messageInfo_View_Diff.DiscardUnknown(m)
}

var xxx_messageInfo_View_Diff proto.InternalMessageInfo

func (m *View_Diff) GetAddNodes() []*Node {
	if m != nil {
		return m.AddNodes
	}
	return nil
}

func (m *View_Diff) GetRemoveNodes() []*Node {
	if m != nil {
		return m.RemoveNodes
	}
	return nil
}

func (m *View_Diff) GetAddLinks() []*Link {
	if m != nil {
		return m.AddLinks
	}
	return nil
}

func (m *View_Diff) GetRemoveLinks() []*Link {
	if m != nil {
		return m.RemoveLinks
	}
	return nil
}

func init() {
	proto.RegisterType((*View)(nil), "topology.View")
	proto.RegisterType((*View_Diff)(nil), "topology.View.Diff")
}

func init() { proto.RegisterFile("view.proto", fileDescriptor_10c1b2aca93c333f) }

var fileDescriptor_10c1b2aca93c333f = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6a, 0xdb, 0x3c,
	0x18, 0xc7, 0xc9, 0x57, 0x9b, 0x28, 0x6d, 0x53, 0xd4, 0x97, 0xbe, 0x22, 0x63, 0x38, 0x68, 0x63,
	0xe4, 0xa0, 0x8e, 0x61, 0x1d, 0x65, 0xec, 0x64, 0x34, 0x94, 0x7d, 0x40, 0xd7, 0x03, 0x1d, 0xec,
	0xa0, 0xfb, 0xc2, 0x8d, 0xe5, 0x54, 0x34, 0xce, 0x13, 0x64, 0x35, 0xa1, 0x1d, 0xbb, 0x93, 0x5d,
	0x9b, 0x07, 0x63, 0x57, 0xe0, 0x2b, 0x18, 0x92, 0xec, 0xc4, 0x73, 0xc3, 0x4e, 0xd6, 0xa3, 0xd8,
	0x7a, 0x7e, 0xbf, 0xbf, 0x1f, 0x3d, 0x48, 0x41, 0x68, 0x2e, 0xf8, 0x62, 0x30, 0x93, 0xa0, 0x00,
	0x37, 0x15, 0xcc, 0x60, 0x02, 0xe3, 0x9b, 0xae, 0x3b, 0x16, 0xea, 0xf2, 0xfa, 0x62, 0x30, 0x82,
	0xc8, 0x1b, 0xc3, 0x18, 0x3c, 0x03, 0x5c, 0x5c, 0x87, 0xe6, 0xcd, 0xbc, 0x98, 0x27, 0x2b, 0x76,
	0xdf, 0x8d, 0x61, 0xc0, 0xfd, 0xf8, 0x06, 0x66, 0xf1, 0x60, 0x02, 0x23, 0x7f, 0xe2, 0x8d, 0x60,
	0xaa, 0xa4, 0x3f, 0x52, 0xb1, 0x35, 0x25, 0x9f, 0x81, 0x1b, 0x41, 0xc0, 0x27, 0xb1, 0x97, 0x81,
	0x9e, 0x79, 0xf5, 0xf2, 0x4f, 0x7a, 0x53, 0x08, 0xf8, 0x3d, 0xc6, 0x4d, 0xc4, 0xf4, 0xea, 0x1e,
	0xe3, 0x7c, 0xc9, 0xfd, 0x7b, 0xdd, 0xac, 0xca, 0x37, 0x7b, 0x54, 0x18, 0x75, 0xb4, 0x10, 0xea,
	0x0a, 0x16, 0xde, 0x18, 0x5c, 0x53, 0x74, 0xe7, 0xfe, 0x44, 0x04, 0xbe, 0x02, 0x19, 0x7b, 0xcb,
	0x47, 0xeb, 0xd1, 0x5f, 0x4d, 0x54, 0x7f, 0x2f, 0xf8, 0x02, 0x3f, 0x44, 0x55, 0x11, 0x90, 0x4a,
	0xaf, 0xd2, 0x6f, 0x0d, 0xb7, 0xd3, 0xc4, 0x69, 0x85, 0x20, 0xa3, 0x17, 0x54, 0x04, 0x94, 0x55,
	0x45, 0x80, 0x1f, 0xa1, 0xfa, 0xd4, 0x8f, 0x38, 0xa9, 0x1a, 0xa0, 0x93, 0x26, 0x4e, 0xdb, 0x02,
	0x7a, 0x95, 0x32, 0x53, 0xc4, 0xaf, 0xd0, 0xe6, 0x48, 0x72, 0x9d, 0x4e, 0x6a, 0x86, 0x3b, 0x48,
	0x13, 0x67, 0xc7, 0x72, 0x59, 0x81, 0xfe, 0xfc, 0xe1, 0xec, 0xa3, 0xff, 0x3e, 0x7f, 0x38, 0x76,
	0xcf, 0x7d, 0xf7, 0xf6, 0x8b, 0xfb, 0xe9, 0xe3, 0xe2, 0xeb, 0xe1, 0xc1, 0xd1, 0xb3, 0x6f, 0x8f,
	0x59, 0x2e, 0xe3, 0xb7, 0xa8, 0x19, 0x41, 0x20, 0x42, 0xc1, 0x25, 0xa9, 0x9b, 0x20, 0x37, 0x4d,
	0x9c, 0x8e, 0x0d, 0xca, 0x2b, 0x7f, 0x4b, 0x5a, 0xea, 0xf8, 0x04, 0x75, 0x24, 0xf7, 0x83, 0xe3,
	0x6b, 0x75, 0x09, 0x52, 0xdc, 0x72, 0x19, 0x93, 0x46, 0xaf, 0xd6, 0x6f, 0x0d, 0xbb, 0x69, 0xe2,
	0xec, 0xdb, 0xc4, 0x12, 0x40, 0x59, 0x59, 0xc1, 0xaf, 0xd1, 0xee, 0x42, 0x0a, 0xc5, 0x8b, 0x31,
	0x1b, 0x26, 0xe6, 0x41, 0x9a, 0x38, 0xff, 0xdb, 0x98, 0x32, 0x41, 0xd9, 0x1d, 0x09, 0x1f, 0xa0,
	0xcd, 0x39, 0x97, 0xb1, 0x80, 0x29, 0xd9, 0x34, 0x1b, 0xc3, 0xab, 0x09, 0x65, 0x05, 0xca, 0x72,
	0x04, 0x3f, 0x41, 0x8d, 0x91, 0x12, 0x11, 0x27, 0xcd, 0x5e, 0xa5, 0xdf, 0x18, 0xee, 0xa6, 0x89,
	0xb3, 0x95, 0x4d, 0x53, 0x2f, 0x53, 0x66, 0xcb, 0x9a, 0x8b, 0x0c, 0xd7, 0x2a, 0x73, 0x51, 0xc6,
	0x99, 0x5f, 0xfc, 0x12, 0x35, 0x25, 0x80, 0x3a, 0x83, 0x80, 0x13, 0xd4, 0xab, 0xf4, 0xdb, 0x4f,
	0x77, 0x06, 0xf9, 0x61, 0x1a, 0xe8, 0xd5, 0xe1, 0xde, 0x6a, 0xce, 0x39, 0x49, 0xd9, 0x52, 0xc2,
	0x47, 0xa8, 0xa1, 0x2f, 0x58, 0x4c, 0xda, 0xbd, 0xda, 0x1a, 0xbb, 0xf0, 0x61, 0x83, 0x51, 0x66,
	0x71, 0xed, 0xe9, 0x9b, 0x14, 0x93, 0xad, 0xb2, 0x77, 0x2a, 0xa6, 0x57, 0x45, 0xcf, 0x60, 0x94,
	0x59, 0x5c, 0x7b, 0xfa, 0xca, 0xc4, 0x64, 0xbb, 0xec, 0x1d, 0x4b, 0xee, 0x17, 0x3d, 0x83, 0x51,
	0x66, 0x71, 0xdb, 0xa7, 0xe2, 0x31, 0xd9, 0xb9, 0xdb, 0xa7, 0x2a, 0xf5, 0xa9, 0xb2, 0x3e, 0x15,
	0x8f, 0xf1, 0x73, 0x54, 0x0f, 0x44, 0x18, 0x92, 0x8e, 0x19, 0xce, 0xde, 0x4a, 0xd3, 0x57, 0x64,
	0x70, 0x22, 0xc2, 0xb0, 0x78, 0xf4, 0x35, 0x4a, 0x99, 0x31, 0xba, 0xdf, 0xab, 0xa8, 0xae, 0xeb,
	0x7a, 0xc6, 0x7e, 0x10, 0x9c, 0x99, 0x29, 0x55, 0xd6, 0x4e, 0xa9, 0x30, 0xe3, 0x9c, 0xa4, 0x6c,
	0x29, 0xe1, 0x37, 0xa8, 0x2d, 0x79, 0x04, 0x73, 0x6e, 0x33, 0xaa, 0x6b, 0x33, 0xf6, 0xd3, 0xc4,
	0xc1, 0xf9, 0xe9, 0x5d, 0xc2, 0x94, 0x15, 0xd5, 0xac, 0x95, 0x53, 0x33, 0xf8, 0xda, 0xda, 0xc1,
	0xff, 0xd9, 0xca, 0xa9, 0x9d, 0xfd, 0x52, 0x5a, 0xb5, 0x62, 0x33, 0xea, 0x6b, 0x33, 0xee, 0xb4,
	0x92, 0xc5, 0x14, 0xd5, 0xe1, 0xc9, 0xf9, 0xf0, 0xdf, 0xff, 0xef, 0x2e, 0x36, 0x0c, 0x78, 0xf8,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x31, 0x51, 0xc9, 0x6e, 0x06, 0x00, 0x00,
}