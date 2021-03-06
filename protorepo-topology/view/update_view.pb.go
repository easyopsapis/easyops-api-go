// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_view.proto

package view

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	topology "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
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
//UpdateView请求
type UpdateViewRequest struct {
	//
	//ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//允许可读用户
	ReadAuthorizers []string `protobuf:"bytes,3,rep,name=readAuthorizers,proto3" json:"readAuthorizers" form:"readAuthorizers"`
	//
	//允许可写用户
	WriteAuthorizers []string `protobuf:"bytes,4,rep,name=writeAuthorizers,proto3" json:"writeAuthorizers" form:"writeAuthorizers"`
	//
	//版本
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version" form:"version"`
	//
	//根节点
	RootNode *topology.Node `protobuf:"bytes,6,opt,name=rootNode,proto3" json:"rootNode" form:"rootNode"`
	//
	//节点
	Nodes []*topology.Node `protobuf:"bytes,7,rep,name=nodes,proto3" json:"nodes" form:"nodes"`
	//
	//线
	Links []*topology.Link `protobuf:"bytes,8,rep,name=links,proto3" json:"links" form:"links"`
	//
	//区域
	Areas []*topology.Area `protobuf:"bytes,9,rep,name=areas,proto3" json:"areas" form:"areas"`
	//
	//标注
	Notes []*topology.Note `protobuf:"bytes,10,rep,name=notes,proto3" json:"notes" form:"notes"`
	//
	//差异
	Diff                 *UpdateViewRequest_Diff `protobuf:"bytes,11,opt,name=diff,proto3" json:"diff" form:"diff"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateViewRequest) Reset()         { *m = UpdateViewRequest{} }
func (m *UpdateViewRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateViewRequest) ProtoMessage()    {}
func (*UpdateViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b17f42b4a591c6db, []int{0}
}
func (m *UpdateViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateViewRequest.Unmarshal(m, b)
}
func (m *UpdateViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateViewRequest.Marshal(b, m, deterministic)
}
func (m *UpdateViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateViewRequest.Merge(m, src)
}
func (m *UpdateViewRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateViewRequest.Size(m)
}
func (m *UpdateViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateViewRequest proto.InternalMessageInfo

func (m *UpdateViewRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateViewRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateViewRequest) GetReadAuthorizers() []string {
	if m != nil {
		return m.ReadAuthorizers
	}
	return nil
}

func (m *UpdateViewRequest) GetWriteAuthorizers() []string {
	if m != nil {
		return m.WriteAuthorizers
	}
	return nil
}

func (m *UpdateViewRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *UpdateViewRequest) GetRootNode() *topology.Node {
	if m != nil {
		return m.RootNode
	}
	return nil
}

func (m *UpdateViewRequest) GetNodes() []*topology.Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *UpdateViewRequest) GetLinks() []*topology.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *UpdateViewRequest) GetAreas() []*topology.Area {
	if m != nil {
		return m.Areas
	}
	return nil
}

func (m *UpdateViewRequest) GetNotes() []*topology.Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *UpdateViewRequest) GetDiff() *UpdateViewRequest_Diff {
	if m != nil {
		return m.Diff
	}
	return nil
}

type UpdateViewRequest_Diff struct {
	//
	//新增节点
	AddNodes []*topology.Node `protobuf:"bytes,1,rep,name=addNodes,proto3" json:"addNodes" form:"addNodes"`
	//
	//已删除节点
	RemoveNodes []*topology.Node `protobuf:"bytes,2,rep,name=removeNodes,proto3" json:"removeNodes" form:"removeNodes"`
	//
	//新增线
	AddLinks []*topology.Link `protobuf:"bytes,3,rep,name=addLinks,proto3" json:"addLinks" form:"addLinks"`
	//
	//已删除线
	RemoveLinks          []*topology.Link `protobuf:"bytes,4,rep,name=removeLinks,proto3" json:"removeLinks" form:"removeLinks"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateViewRequest_Diff) Reset()         { *m = UpdateViewRequest_Diff{} }
func (m *UpdateViewRequest_Diff) String() string { return proto.CompactTextString(m) }
func (*UpdateViewRequest_Diff) ProtoMessage()    {}
func (*UpdateViewRequest_Diff) Descriptor() ([]byte, []int) {
	return fileDescriptor_b17f42b4a591c6db, []int{0, 0}
}
func (m *UpdateViewRequest_Diff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateViewRequest_Diff.Unmarshal(m, b)
}
func (m *UpdateViewRequest_Diff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateViewRequest_Diff.Marshal(b, m, deterministic)
}
func (m *UpdateViewRequest_Diff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateViewRequest_Diff.Merge(m, src)
}
func (m *UpdateViewRequest_Diff) XXX_Size() int {
	return xxx_messageInfo_UpdateViewRequest_Diff.Size(m)
}
func (m *UpdateViewRequest_Diff) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateViewRequest_Diff.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateViewRequest_Diff proto.InternalMessageInfo

func (m *UpdateViewRequest_Diff) GetAddNodes() []*topology.Node {
	if m != nil {
		return m.AddNodes
	}
	return nil
}

func (m *UpdateViewRequest_Diff) GetRemoveNodes() []*topology.Node {
	if m != nil {
		return m.RemoveNodes
	}
	return nil
}

func (m *UpdateViewRequest_Diff) GetAddLinks() []*topology.Link {
	if m != nil {
		return m.AddLinks
	}
	return nil
}

func (m *UpdateViewRequest_Diff) GetRemoveLinks() []*topology.Link {
	if m != nil {
		return m.RemoveLinks
	}
	return nil
}

//
//UpdateView返回
type UpdateViewResponse struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateViewResponse) Reset()         { *m = UpdateViewResponse{} }
func (m *UpdateViewResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateViewResponse) ProtoMessage()    {}
func (*UpdateViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b17f42b4a591c6db, []int{1}
}
func (m *UpdateViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateViewResponse.Unmarshal(m, b)
}
func (m *UpdateViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateViewResponse.Marshal(b, m, deterministic)
}
func (m *UpdateViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateViewResponse.Merge(m, src)
}
func (m *UpdateViewResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateViewResponse.Size(m)
}
func (m *UpdateViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateViewResponse proto.InternalMessageInfo

func (m *UpdateViewResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//UpdateViewApi返回
type UpdateViewResponseWrapper struct {
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
	Data                 *UpdateViewResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateViewResponseWrapper) Reset()         { *m = UpdateViewResponseWrapper{} }
func (m *UpdateViewResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateViewResponseWrapper) ProtoMessage()    {}
func (*UpdateViewResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b17f42b4a591c6db, []int{2}
}
func (m *UpdateViewResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateViewResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateViewResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateViewResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateViewResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateViewResponseWrapper.Merge(m, src)
}
func (m *UpdateViewResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateViewResponseWrapper.Size(m)
}
func (m *UpdateViewResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateViewResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateViewResponseWrapper proto.InternalMessageInfo

func (m *UpdateViewResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateViewResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateViewResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateViewResponseWrapper) GetData() *UpdateViewResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateViewRequest)(nil), "view.UpdateViewRequest")
	proto.RegisterType((*UpdateViewRequest_Diff)(nil), "view.UpdateViewRequest.Diff")
	proto.RegisterType((*UpdateViewResponse)(nil), "view.UpdateViewResponse")
	proto.RegisterType((*UpdateViewResponseWrapper)(nil), "view.UpdateViewResponseWrapper")
}

func init() { proto.RegisterFile("update_view.proto", fileDescriptor_b17f42b4a591c6db) }

var fileDescriptor_b17f42b4a591c6db = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x95, 0x64, 0xd2, 0x36, 0x0e, 0xf4, 0xc7, 0x48, 0xc5, 0x04, 0xd0, 0x44, 0x46, 0x42,
	0x5d, 0xd0, 0x19, 0x89, 0x4a, 0x08, 0x21, 0x21, 0xd4, 0xaa, 0x08, 0x16, 0xa5, 0x0b, 0x4b, 0xc0,
	0x12, 0xb9, 0x19, 0x27, 0xb5, 0x9a, 0xcc, 0x1d, 0x6c, 0xa7, 0xa5, 0xbc, 0x08, 0x1b, 0x9e, 0x6d,
	0x24, 0x5e, 0x61, 0x9e, 0x00, 0xd9, 0x9e, 0x04, 0x37, 0x89, 0xc4, 0x26, 0xab, 0x8c, 0x7d, 0xcf,
	0x77, 0x75, 0x7c, 0x7d, 0x1c, 0xb4, 0x37, 0x2d, 0x32, 0x6e, 0xc4, 0xb7, 0x6b, 0x29, 0x6e, 0x92,
	0x42, 0x81, 0x01, 0x1c, 0xd9, 0xef, 0xde, 0xe1, 0x48, 0x9a, 0xcb, 0xe9, 0x45, 0x32, 0x80, 0x49,
	0x3a, 0x82, 0x11, 0xa4, 0xae, 0x78, 0x31, 0x1d, 0xba, 0x95, 0x5b, 0xb8, 0x2f, 0x0f, 0xf5, 0x3e,
	0x8d, 0x20, 0x11, 0x5c, 0xdf, 0x42, 0xa1, 0x93, 0x31, 0x0c, 0xf8, 0x38, 0x1d, 0x40, 0x6e, 0x14,
	0x1f, 0x18, 0xed, 0x49, 0x25, 0x0a, 0x38, 0x9c, 0x40, 0x26, 0xc6, 0x3a, 0xad, 0x85, 0xa9, 0x5b,
	0xa6, 0x06, 0x0a, 0x18, 0xc3, 0xe8, 0x36, 0xcd, 0x21, 0x13, 0x6b, 0x6c, 0x37, 0x96, 0xf9, 0xd5,
	0x1a, 0xdb, 0x71, 0x25, 0xf8, 0x5a, 0x0f, 0x6b, 0xea, 0xc3, 0xd2, 0x5f, 0x9b, 0x68, 0xef, 0xb3,
	0xbb, 0x86, 0x2f, 0x52, 0xdc, 0x30, 0xf1, 0x7d, 0x2a, 0xb4, 0xc1, 0x4f, 0x51, 0x53, 0x66, 0xa4,
	0xd1, 0x6f, 0x1c, 0x74, 0x4e, 0xee, 0x57, 0x65, 0xdc, 0x19, 0x82, 0x9a, 0xbc, 0xa1, 0x32, 0xa3,
	0xac, 0x29, 0x33, 0xfc, 0x0c, 0x45, 0x39, 0x9f, 0x08, 0xd2, 0x74, 0x82, 0x9d, 0xaa, 0x8c, 0xbb,
	0x5e, 0x60, 0x77, 0x29, 0x73, 0x45, 0x7c, 0x8a, 0x76, 0x94, 0xe0, 0xd9, 0xf1, 0xd4, 0x5c, 0x82,
	0x92, 0x3f, 0x85, 0xd2, 0xa4, 0xd5, 0x6f, 0x1d, 0x74, 0x4e, 0x7a, 0x55, 0x19, 0xef, 0x7b, 0xfd,
	0x82, 0x80, 0xb2, 0x45, 0x04, 0x7f, 0x40, 0xbb, 0x37, 0x4a, 0x1a, 0x11, 0xb6, 0x89, 0x5c, 0x9b,
	0xc7, 0x55, 0x19, 0x3f, 0xf4, 0x6d, 0x16, 0x15, 0x94, 0x2d, 0x41, 0xf8, 0x05, 0xda, 0xbc, 0x16,
	0x4a, 0x4b, 0xc8, 0x49, 0xdb, 0xd9, 0xc6, 0x55, 0x19, 0x6f, 0x7b, 0xbe, 0x2e, 0x50, 0x36, 0x93,
	0xe0, 0x77, 0x68, 0x4b, 0x01, 0x98, 0x73, 0xc8, 0x04, 0xd9, 0xe8, 0x37, 0x0e, 0xba, 0x2f, 0xb7,
	0x93, 0xd9, 0xf8, 0x12, 0xbb, 0x7b, 0xf2, 0xa0, 0x2a, 0xe3, 0x9d, 0xfa, 0x14, 0xb5, 0x92, 0xb2,
	0x39, 0x84, 0x5f, 0xa1, 0xb6, 0x8d, 0x94, 0x26, 0x9b, 0xfd, 0xd6, 0x0a, 0x7a, 0xb7, 0x2a, 0xe3,
	0x7b, 0xf5, 0xcc, 0xac, 0x8c, 0x32, 0x2f, 0xb7, 0x9c, 0xcd, 0x8e, 0x26, 0x5b, 0x8b, 0xdc, 0x99,
	0xcc, 0xaf, 0x42, 0xce, 0xc9, 0x28, 0xf3, 0x72, 0xcb, 0xd9, 0x90, 0x68, 0xd2, 0x59, 0xe4, 0x8e,
	0x95, 0xe0, 0x21, 0xe7, 0x64, 0x94, 0x79, 0xb9, 0xf7, 0x69, 0x84, 0x26, 0x68, 0xd9, 0xa7, 0x59,
	0xf0, 0x69, 0x6a, 0x9f, 0x46, 0x68, 0x7c, 0x8c, 0xa2, 0x4c, 0x0e, 0x87, 0xa4, 0xeb, 0x86, 0xf3,
	0x24, 0x71, 0x6f, 0x78, 0x29, 0x48, 0xc9, 0xa9, 0x1c, 0x0e, 0xc3, 0x80, 0x58, 0x86, 0x32, 0x87,
	0xf6, 0x7e, 0x37, 0x51, 0x64, 0xeb, 0x76, 0xd8, 0x3c, 0xcb, 0xce, 0xdd, 0xb8, 0x1a, 0x2b, 0xc7,
	0x15, 0x0c, 0x7b, 0xa6, 0xa4, 0x6c, 0x0e, 0xe1, 0x8f, 0xa8, 0xab, 0xc4, 0x04, 0xae, 0x85, 0xef,
	0xd1, 0x5c, 0xd9, 0x63, 0xbf, 0x2a, 0x63, 0x3c, 0x8b, 0xdd, 0x5c, 0x4c, 0x59, 0x88, 0xd6, 0x56,
	0xce, 0xdc, 0x0d, 0xb4, 0x56, 0xde, 0xc0, 0x5d, 0x2b, 0x67, 0xfe, 0x12, 0xe6, 0xd0, 0x3f, 0x2b,
	0xbe, 0x47, 0xb4, 0xb2, 0xc7, 0x92, 0x95, 0xba, 0x4d, 0x88, 0xd2, 0x23, 0x84, 0xc3, 0x79, 0xea,
	0x02, 0x72, 0x2d, 0xfe, 0xf3, 0x32, 0xe9, 0x9f, 0x06, 0x7a, 0xb4, 0x4c, 0x7d, 0x55, 0xbc, 0x28,
	0x84, 0xb2, 0xef, 0x76, 0x60, 0x13, 0x6d, 0xf1, 0x76, 0x78, 0x2d, 0x03, 0x97, 0x5e, 0x57, 0xc4,
	0xaf, 0x51, 0xd7, 0xfe, 0xbe, 0xff, 0x51, 0x8c, 0xb9, 0xcc, 0xeb, 0x37, 0x1e, 0x38, 0x0e, 0x8a,
	0x94, 0x85, 0x52, 0xfc, 0x1c, 0xb5, 0x85, 0x52, 0xa0, 0x48, 0xcb, 0x31, 0x41, 0x76, 0xdc, 0x36,
	0x65, 0xbe, 0x8c, 0xdf, 0xa2, 0x28, 0xe3, 0x86, 0x93, 0xc8, 0x65, 0x87, 0x2c, 0x67, 0xc7, 0xbb,
	0xbe, 0x93, 0x1b, 0x6e, 0xb8, 0xcd, 0x0d, 0x37, 0xfc, 0x62, 0xc3, 0xfd, 0x73, 0x1d, 0xfd, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x97, 0x65, 0x4c, 0x3f, 0x06, 0x00, 0x00,
}
