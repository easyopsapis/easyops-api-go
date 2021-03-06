// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: add_str_to_set.proto

package instance

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
//AddStrToSet请求
type AddStrToSetRequest struct {
	//
	//objectId
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//属性key(仅支持arr类型属性)
	AttrId string `protobuf:"bytes,2,opt,name=attr_id,json=attrId,proto3" json:"attr_id" form:"attr_id"`
	//
	//与实例搜索query一致,将对search到的所有实例操作
	Query *types.Struct `protobuf:"bytes,3,opt,name=query,proto3" json:"query" form:"query"`
	//
	//append的值
	Values               []string `protobuf:"bytes,4,rep,name=values,proto3" json:"values" form:"values"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStrToSetRequest) Reset()         { *m = AddStrToSetRequest{} }
func (m *AddStrToSetRequest) String() string { return proto.CompactTextString(m) }
func (*AddStrToSetRequest) ProtoMessage()    {}
func (*AddStrToSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e71ff9307d0edb0, []int{0}
}
func (m *AddStrToSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStrToSetRequest.Unmarshal(m, b)
}
func (m *AddStrToSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStrToSetRequest.Marshal(b, m, deterministic)
}
func (m *AddStrToSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStrToSetRequest.Merge(m, src)
}
func (m *AddStrToSetRequest) XXX_Size() int {
	return xxx_messageInfo_AddStrToSetRequest.Size(m)
}
func (m *AddStrToSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStrToSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddStrToSetRequest proto.InternalMessageInfo

func (m *AddStrToSetRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AddStrToSetRequest) GetAttrId() string {
	if m != nil {
		return m.AttrId
	}
	return ""
}

func (m *AddStrToSetRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *AddStrToSetRequest) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

//
//AddStrToSet返回
type AddStrToSetResponse struct {
	//
	//实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//属性值
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values" form:"values"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStrToSetResponse) Reset()         { *m = AddStrToSetResponse{} }
func (m *AddStrToSetResponse) String() string { return proto.CompactTextString(m) }
func (*AddStrToSetResponse) ProtoMessage()    {}
func (*AddStrToSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e71ff9307d0edb0, []int{1}
}
func (m *AddStrToSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStrToSetResponse.Unmarshal(m, b)
}
func (m *AddStrToSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStrToSetResponse.Marshal(b, m, deterministic)
}
func (m *AddStrToSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStrToSetResponse.Merge(m, src)
}
func (m *AddStrToSetResponse) XXX_Size() int {
	return xxx_messageInfo_AddStrToSetResponse.Size(m)
}
func (m *AddStrToSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStrToSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddStrToSetResponse proto.InternalMessageInfo

func (m *AddStrToSetResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *AddStrToSetResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

//
//AddStrToSetApi返回
type AddStrToSetResponseWrapper struct {
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
	Data                 []*AddStrToSetResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AddStrToSetResponseWrapper) Reset()         { *m = AddStrToSetResponseWrapper{} }
func (m *AddStrToSetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AddStrToSetResponseWrapper) ProtoMessage()    {}
func (*AddStrToSetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e71ff9307d0edb0, []int{2}
}
func (m *AddStrToSetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStrToSetResponseWrapper.Unmarshal(m, b)
}
func (m *AddStrToSetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStrToSetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AddStrToSetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStrToSetResponseWrapper.Merge(m, src)
}
func (m *AddStrToSetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AddStrToSetResponseWrapper.Size(m)
}
func (m *AddStrToSetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStrToSetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AddStrToSetResponseWrapper proto.InternalMessageInfo

func (m *AddStrToSetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AddStrToSetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AddStrToSetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AddStrToSetResponseWrapper) GetData() []*AddStrToSetResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AddStrToSetRequest)(nil), "instance.AddStrToSetRequest")
	proto.RegisterType((*AddStrToSetResponse)(nil), "instance.AddStrToSetResponse")
	proto.RegisterType((*AddStrToSetResponseWrapper)(nil), "instance.AddStrToSetResponseWrapper")
}

func init() { proto.RegisterFile("add_str_to_set.proto", fileDescriptor_8e71ff9307d0edb0) }

var fileDescriptor_8e71ff9307d0edb0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0x56, 0xbb, 0xb6, 0xac, 0x2e, 0xec, 0x8f, 0x41, 0x50, 0x55, 0x4c, 0xae, 0xcc, 0x84, 0x3a,
	0x89, 0xa4, 0xdb, 0x2a, 0x26, 0xe0, 0x82, 0x16, 0x89, 0x43, 0x2f, 0x1c, 0x5c, 0x24, 0x24, 0xa6,
	0x2d, 0x72, 0x63, 0x2f, 0x04, 0xda, 0x3a, 0xb3, 0x9d, 0x0d, 0x98, 0xf6, 0x04, 0xbc, 0x06, 0xcf,
	0x15, 0x24, 0x2e, 0xdc, 0xf3, 0x04, 0x28, 0x76, 0xd2, 0x06, 0x69, 0x12, 0x9c, 0x62, 0xfb, 0xfb,
	0x93, 0xef, 0xfb, 0xfd, 0xc0, 0x03, 0xca, 0x98, 0xaf, 0xb4, 0xf4, 0xb5, 0xf0, 0x15, 0xd7, 0x6e,
	0x2c, 0x85, 0x16, 0x70, 0x3d, 0x5a, 0x28, 0x4d, 0x17, 0x01, 0xef, 0x39, 0x61, 0xa4, 0x3f, 0x26,
	0x53, 0x37, 0x10, 0xf3, 0x61, 0x28, 0x42, 0x31, 0x34, 0x84, 0x69, 0x72, 0x6e, 0x6e, 0xe6, 0x62,
	0x4e, 0x56, 0xd8, 0x3b, 0xaa, 0xd0, 0xe7, 0x57, 0x91, 0xfe, 0x2c, 0xae, 0x86, 0xa1, 0x70, 0x0c,
	0xe8, 0x5c, 0xd2, 0x59, 0xc4, 0xa8, 0x16, 0x52, 0x0d, 0x97, 0xc7, 0x42, 0xf7, 0x38, 0x14, 0x22,
	0x9c, 0xf1, 0x95, 0xbb, 0xd2, 0x32, 0x09, 0x8a, 0x38, 0xf8, 0x47, 0x1d, 0xc0, 0x63, 0xc6, 0x26,
	0x5a, 0xbe, 0x13, 0x13, 0xae, 0x09, 0xbf, 0x48, 0xb8, 0xd2, 0x90, 0x80, 0x75, 0x31, 0xfd, 0xc4,
	0x03, 0x3d, 0x66, 0xdd, 0x5a, 0xbf, 0x36, 0x68, 0x7b, 0x47, 0x59, 0x8a, 0x36, 0xcf, 0x85, 0x9c,
	0xbf, 0xc2, 0x25, 0x82, 0x7f, 0xfd, 0x44, 0x08, 0xec, 0x9c, 0x9d, 0x50, 0xe7, 0xdb, 0xb1, 0xf3,
	0xc1, 0x3f, 0x3d, 0xd9, 0x77, 0x5e, 0x96, 0xe7, 0xeb, 0xfd, 0x67, 0xa3, 0x83, 0x9b, 0x5d, 0xb2,
	0xf4, 0x81, 0x6f, 0xc1, 0x1d, 0xaa, 0xb5, 0xf4, 0x23, 0xd6, 0xad, 0x1b, 0xcb, 0xe7, 0x59, 0x8a,
	0x36, 0xac, 0x65, 0x01, 0xfc, 0x97, 0x63, 0x2b, 0x27, 0x8f, 0x19, 0x7c, 0x0d, 0x9a, 0x17, 0x09,
	0x97, 0x5f, 0xbb, 0x6b, 0xfd, 0xda, 0xa0, 0x73, 0xf8, 0xc8, 0xb5, 0x45, 0xdd, 0xb2, 0xa8, 0x3b,
	0x31, 0x45, 0xbd, 0xad, 0x2c, 0x45, 0x77, 0xed, 0x6f, 0x0c, 0x1f, 0x13, 0xab, 0x83, 0x7b, 0xa0,
	0x75, 0x49, 0x67, 0x09, 0x57, 0xdd, 0x46, 0x7f, 0x6d, 0xd0, 0xf6, 0xb6, 0xb3, 0x14, 0xdd, 0xb3,
	0x44, 0xfb, 0x8e, 0x49, 0x41, 0xc0, 0xdf, 0x6b, 0xe0, 0xfe, 0x5f, 0x63, 0x52, 0xb1, 0x58, 0x28,
	0x0e, 0xc7, 0x00, 0x94, 0xfb, 0x5c, 0x4e, 0x6a, 0x2f, 0x4b, 0xd1, 0xb6, 0xb5, 0x59, 0x61, 0x79,
	0xb3, 0x2d, 0xb0, 0x71, 0x56, 0x14, 0x3a, 0xbd, 0x3e, 0x18, 0xdd, 0xec, 0x92, 0x8a, 0xb8, 0x92,
	0xa6, 0xfe, 0xaf, 0x34, 0xbf, 0x6b, 0xa0, 0x77, 0x4b, 0x9a, 0xf7, 0x92, 0xc6, 0x31, 0x97, 0xf0,
	0x09, 0x68, 0x04, 0x82, 0x71, 0x13, 0xa7, 0xe9, 0x6d, 0x66, 0x29, 0xea, 0x58, 0x9f, 0xfc, 0x15,
	0x13, 0x03, 0xc2, 0x17, 0xa0, 0x93, 0x7f, 0xdf, 0x7c, 0x89, 0x67, 0x34, 0x5a, 0x14, 0x1b, 0x79,
	0x98, 0xa5, 0x08, 0xae, 0xb8, 0x05, 0x88, 0x49, 0x95, 0x0a, 0x9f, 0x82, 0x26, 0x97, 0x52, 0x48,
	0x33, 0xf7, 0x76, 0x75, 0xbc, 0xe6, 0x19, 0x13, 0x0b, 0x43, 0x0f, 0x34, 0x18, 0xd5, 0xd4, 0x0c,
	0xb7, 0x73, 0xb8, 0xe3, 0x96, 0x5d, 0xdd, 0x5b, 0xa2, 0x57, 0x53, 0xe6, 0x22, 0x4c, 0x8c, 0x76,
	0xda, 0x32, 0xcb, 0x1c, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x3e, 0x19, 0x51, 0x4c, 0x03,
	0x00, 0x00,
}
