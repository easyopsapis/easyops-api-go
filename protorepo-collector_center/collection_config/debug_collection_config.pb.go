// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: debug_collection_config.proto

package collection_config

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
//DebugCollectionConfig请求
type DebugCollectionConfigRequest struct {
	//
	//目标主机
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host" form:"host"`
	//
	//采集类型ID
	ClazzId string `protobuf:"bytes,2,opt,name=clazzId,proto3" json:"clazzId" form:"clazzId"`
	//
	//是否忽略不符合要求数据
	IgnoreInvalid bool `protobuf:"varint,3,opt,name=ignoreInvalid,proto3" json:"ignoreInvalid" form:"ignoreInvalid"`
	//
	//超时时间
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout" form:"timeout"`
	//
	//脚本
	Script *DebugCollectionConfigRequest_Script `protobuf:"bytes,5,opt,name=script,proto3" json:"script" form:"script"`
	//
	//执行时间范围（24小时）
	TimeRange string `protobuf:"bytes,6,opt,name=timeRange,proto3" json:"timeRange" form:"timeRange"`
	//
	//名称
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" form:"name"`
	//
	//环境变量
	Env []*DebugCollectionConfigRequest_Env `protobuf:"bytes,8,rep,name=env,proto3" json:"env" form:"env"`
	//
	//参数
	Kwargs []*DebugCollectionConfigRequest_Kwargs `protobuf:"bytes,9,rep,name=kwargs,proto3" json:"kwargs" form:"kwargs"`
	//
	//关联模型ID
	ObjectId             string   `protobuf:"bytes,10,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectionConfigRequest) Reset()         { *m = DebugCollectionConfigRequest{} }
func (m *DebugCollectionConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigRequest) ProtoMessage()    {}
func (*DebugCollectionConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{0}
}
func (m *DebugCollectionConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigRequest.Unmarshal(m, b)
}
func (m *DebugCollectionConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigRequest.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigRequest.Merge(m, src)
}
func (m *DebugCollectionConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigRequest.Size(m)
}
func (m *DebugCollectionConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigRequest proto.InternalMessageInfo

func (m *DebugCollectionConfigRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DebugCollectionConfigRequest) GetClazzId() string {
	if m != nil {
		return m.ClazzId
	}
	return ""
}

func (m *DebugCollectionConfigRequest) GetIgnoreInvalid() bool {
	if m != nil {
		return m.IgnoreInvalid
	}
	return false
}

func (m *DebugCollectionConfigRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *DebugCollectionConfigRequest) GetScript() *DebugCollectionConfigRequest_Script {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *DebugCollectionConfigRequest) GetTimeRange() string {
	if m != nil {
		return m.TimeRange
	}
	return ""
}

func (m *DebugCollectionConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DebugCollectionConfigRequest) GetEnv() []*DebugCollectionConfigRequest_Env {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *DebugCollectionConfigRequest) GetKwargs() []*DebugCollectionConfigRequest_Kwargs {
	if m != nil {
		return m.Kwargs
	}
	return nil
}

func (m *DebugCollectionConfigRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type DebugCollectionConfigRequest_Script struct {
	//
	//脚本Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//脚本版本
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version" form:"version"`
	//
	//脚本类型
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type" form:"type"`
	//
	//脚本内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content" form:"content"`
	//
	//脚本名称
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectionConfigRequest_Script) Reset()         { *m = DebugCollectionConfigRequest_Script{} }
func (m *DebugCollectionConfigRequest_Script) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigRequest_Script) ProtoMessage()    {}
func (*DebugCollectionConfigRequest_Script) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{0, 0}
}
func (m *DebugCollectionConfigRequest_Script) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigRequest_Script.Unmarshal(m, b)
}
func (m *DebugCollectionConfigRequest_Script) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigRequest_Script.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigRequest_Script) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigRequest_Script.Merge(m, src)
}
func (m *DebugCollectionConfigRequest_Script) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigRequest_Script.Size(m)
}
func (m *DebugCollectionConfigRequest_Script) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigRequest_Script.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigRequest_Script proto.InternalMessageInfo

func (m *DebugCollectionConfigRequest_Script) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Script) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Script) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Script) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Script) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DebugCollectionConfigRequest_Env struct {
	//
	//键
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" form:"key"`
	//
	//值
	Value                *types.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DebugCollectionConfigRequest_Env) Reset()         { *m = DebugCollectionConfigRequest_Env{} }
func (m *DebugCollectionConfigRequest_Env) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigRequest_Env) ProtoMessage()    {}
func (*DebugCollectionConfigRequest_Env) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{0, 1}
}
func (m *DebugCollectionConfigRequest_Env) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigRequest_Env.Unmarshal(m, b)
}
func (m *DebugCollectionConfigRequest_Env) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigRequest_Env.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigRequest_Env) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigRequest_Env.Merge(m, src)
}
func (m *DebugCollectionConfigRequest_Env) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigRequest_Env.Size(m)
}
func (m *DebugCollectionConfigRequest_Env) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigRequest_Env.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigRequest_Env proto.InternalMessageInfo

func (m *DebugCollectionConfigRequest_Env) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Env) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type DebugCollectionConfigRequest_Kwargs struct {
	//
	//键
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" form:"key"`
	//
	//值
	Value                *types.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DebugCollectionConfigRequest_Kwargs) Reset()         { *m = DebugCollectionConfigRequest_Kwargs{} }
func (m *DebugCollectionConfigRequest_Kwargs) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigRequest_Kwargs) ProtoMessage()    {}
func (*DebugCollectionConfigRequest_Kwargs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{0, 2}
}
func (m *DebugCollectionConfigRequest_Kwargs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigRequest_Kwargs.Unmarshal(m, b)
}
func (m *DebugCollectionConfigRequest_Kwargs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigRequest_Kwargs.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigRequest_Kwargs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigRequest_Kwargs.Merge(m, src)
}
func (m *DebugCollectionConfigRequest_Kwargs) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigRequest_Kwargs.Size(m)
}
func (m *DebugCollectionConfigRequest_Kwargs) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigRequest_Kwargs.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigRequest_Kwargs proto.InternalMessageInfo

func (m *DebugCollectionConfigRequest_Kwargs) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DebugCollectionConfigRequest_Kwargs) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

//
//DebugCollectionConfig返回
type DebugCollectionConfigResponse struct {
	//
	//调试任务ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugCollectionConfigResponse) Reset()         { *m = DebugCollectionConfigResponse{} }
func (m *DebugCollectionConfigResponse) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigResponse) ProtoMessage()    {}
func (*DebugCollectionConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{1}
}
func (m *DebugCollectionConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigResponse.Unmarshal(m, b)
}
func (m *DebugCollectionConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigResponse.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigResponse.Merge(m, src)
}
func (m *DebugCollectionConfigResponse) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigResponse.Size(m)
}
func (m *DebugCollectionConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigResponse proto.InternalMessageInfo

func (m *DebugCollectionConfigResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//DebugCollectionConfigApi返回
type DebugCollectionConfigResponseWrapper struct {
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
	Data                 *DebugCollectionConfigResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DebugCollectionConfigResponseWrapper) Reset()         { *m = DebugCollectionConfigResponseWrapper{} }
func (m *DebugCollectionConfigResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*DebugCollectionConfigResponseWrapper) ProtoMessage()    {}
func (*DebugCollectionConfigResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_d629e42c1956a2a3, []int{2}
}
func (m *DebugCollectionConfigResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugCollectionConfigResponseWrapper.Unmarshal(m, b)
}
func (m *DebugCollectionConfigResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugCollectionConfigResponseWrapper.Marshal(b, m, deterministic)
}
func (m *DebugCollectionConfigResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugCollectionConfigResponseWrapper.Merge(m, src)
}
func (m *DebugCollectionConfigResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_DebugCollectionConfigResponseWrapper.Size(m)
}
func (m *DebugCollectionConfigResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugCollectionConfigResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_DebugCollectionConfigResponseWrapper proto.InternalMessageInfo

func (m *DebugCollectionConfigResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DebugCollectionConfigResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *DebugCollectionConfigResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *DebugCollectionConfigResponseWrapper) GetData() *DebugCollectionConfigResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DebugCollectionConfigRequest)(nil), "collection_config.DebugCollectionConfigRequest")
	proto.RegisterType((*DebugCollectionConfigRequest_Script)(nil), "collection_config.DebugCollectionConfigRequest.Script")
	proto.RegisterType((*DebugCollectionConfigRequest_Env)(nil), "collection_config.DebugCollectionConfigRequest.Env")
	proto.RegisterType((*DebugCollectionConfigRequest_Kwargs)(nil), "collection_config.DebugCollectionConfigRequest.Kwargs")
	proto.RegisterType((*DebugCollectionConfigResponse)(nil), "collection_config.DebugCollectionConfigResponse")
	proto.RegisterType((*DebugCollectionConfigResponseWrapper)(nil), "collection_config.DebugCollectionConfigResponseWrapper")
}

func init() { proto.RegisterFile("debug_collection_config.proto", fileDescriptor_d629e42c1956a2a3) }

var fileDescriptor_d629e42c1956a2a3 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xc7, 0x4f, 0x3e, 0x21, 0x1b, 0x3e, 0xf7, 0x9c, 0x83, 0xac, 0x08, 0xe4, 0x68, 0xe1, 0x1c,
	0xa5, 0x2a, 0x71, 0x50, 0xa8, 0x50, 0x8b, 0xd4, 0xa8, 0x0a, 0xe5, 0x02, 0xf5, 0x6e, 0x5b, 0xda,
	0x0b, 0x84, 0x90, 0x63, 0x6f, 0x8c, 0x49, 0xe2, 0x75, 0xed, 0x4d, 0x28, 0x20, 0xee, 0xfa, 0x1a,
	0x7d, 0xa4, 0xbe, 0x82, 0x2b, 0xf5, 0x11, 0xfc, 0x04, 0xd5, 0xce, 0xda, 0xc4, 0x14, 0x15, 0xca,
	0x45, 0xaf, 0xd8, 0xcc, 0x7f, 0xe6, 0xb7, 0x3b, 0x33, 0x7f, 0x8c, 0xd6, 0x6c, 0xd6, 0x1b, 0x3b,
	0x27, 0x16, 0x1f, 0x0e, 0x99, 0x25, 0x5c, 0xee, 0x9d, 0x58, 0xdc, 0xeb, 0xbb, 0x8e, 0xe1, 0x07,
	0x5c, 0x70, 0xbc, 0x7c, 0x47, 0xa8, 0x35, 0x1d, 0x57, 0x9c, 0x8e, 0x7b, 0x86, 0xc5, 0x47, 0x2d,
	0x87, 0x3b, 0xbc, 0x05, 0x99, 0xbd, 0x71, 0x1f, 0x7e, 0xc1, 0x0f, 0x38, 0x29, 0x42, 0x6d, 0x27,
	0x93, 0x3e, 0x3a, 0x77, 0xc5, 0x80, 0x9f, 0xb7, 0x1c, 0xde, 0x04, 0xb1, 0x39, 0x31, 0x87, 0xae,
	0x6d, 0x0a, 0x1e, 0x84, 0xad, 0x9b, 0x63, 0x52, 0xb7, 0xea, 0x70, 0xee, 0x0c, 0xd9, 0x94, 0x1e,
	0x8a, 0x60, 0x6c, 0x09, 0xa5, 0x92, 0x2f, 0x15, 0xb4, 0xfa, 0x5a, 0xbe, 0x7c, 0xef, 0xe6, 0x7d,
	0x7b, 0xf0, 0x3c, 0xca, 0x3e, 0x8e, 0x59, 0x28, 0xf0, 0x3b, 0x54, 0x3c, 0xe5, 0xa1, 0xd0, 0x72,
	0xf5, 0x5c, 0xa3, 0xd2, 0x7d, 0x15, 0x47, 0x7a, 0xb5, 0xcf, 0x83, 0xd1, 0x2e, 0x91, 0x51, 0xf2,
	0xfd, 0x9b, 0xde, 0x44, 0x4f, 0x8f, 0xb6, 0x9a, 0x2f, 0x8e, 0xaf, 0xb6, 0x36, 0xb7, 0xaf, 0x8d,
	0x87, 0x8e, 0x14, 0x68, 0xb8, 0x83, 0x66, 0xac, 0xa1, 0x79, 0x79, 0x79, 0x60, 0x6b, 0x79, 0x00,
	0x6f, 0xc4, 0x91, 0xbe, 0xa0, 0xc0, 0x89, 0x20, 0xd9, 0x0b, 0x68, 0x4e, 0x96, 0x9a, 0xcd, 0xfe,
	0xf1, 0x55, 0xfb, 0xd9, 0x35, 0x4d, 0x8b, 0x70, 0x07, 0xcd, 0xbb, 0x8e, 0xc7, 0x03, 0x76, 0xe0,
	0x41, 0xbf, 0x5a, 0xa1, 0x9e, 0x6b, 0xcc, 0x76, 0xb5, 0x38, 0xd2, 0xff, 0x51, 0x94, 0x5b, 0x32,
	0xa1, 0xb7, 0xd3, 0xf1, 0x26, 0x9a, 0x11, 0xee, 0x88, 0xf1, 0xb1, 0xd0, 0x8a, 0xf5, 0x5c, 0xa3,
	0xd4, 0xc5, 0xd3, 0xfb, 0x13, 0x81, 0xd0, 0x34, 0x05, 0x9b, 0xa8, 0x1c, 0x5a, 0x81, 0xeb, 0x0b,
	0xad, 0x54, 0xcf, 0x35, 0xaa, 0xed, 0x1d, 0xe3, 0xee, 0x9a, 0xef, 0x1b, 0xa2, 0xf1, 0x16, 0xaa,
	0xbb, 0xcb, 0x71, 0xa4, 0xcf, 0xab, 0x4b, 0x14, 0x8f, 0xd0, 0x04, 0x8c, 0x8f, 0x50, 0x45, 0xde,
	0x46, 0x4d, 0xcf, 0x61, 0x5a, 0x19, 0x46, 0xf2, 0x32, 0x8e, 0xf4, 0xa5, 0xe9, 0x93, 0x40, 0x92,
	0x43, 0xf9, 0x0f, 0xad, 0xab, 0x79, 0xb6, 0xaf, 0x77, 0xd3, 0x43, 0xf3, 0x4e, 0x84, 0x4e, 0x79,
	0xf8, 0x09, 0x2a, 0x7a, 0xe6, 0x88, 0x69, 0x33, 0xc0, 0xfd, 0x77, 0xba, 0x43, 0x19, 0x95, 0xc8,
	0xbc, 0xff, 0x17, 0x85, 0x14, 0x7c, 0x88, 0x0a, 0xcc, 0x9b, 0x68, 0xb3, 0xf5, 0x42, 0xa3, 0xda,
	0xde, 0x7e, 0x6c, 0x9f, 0xfb, 0xde, 0xa4, 0xbb, 0x10, 0x47, 0x3a, 0x52, 0x78, 0xe6, 0x4d, 0x08,
	0x95, 0x3c, 0x39, 0xc1, 0xc1, 0xb9, 0x19, 0x38, 0xa1, 0x56, 0x01, 0xf2, 0xa3, 0x27, 0xf8, 0x06,
	0xaa, 0xb3, 0x13, 0x54, 0x3c, 0x42, 0x13, 0x30, 0x6e, 0xa1, 0x59, 0xde, 0x3b, 0x63, 0x96, 0x38,
	0xb0, 0x35, 0x04, 0x8d, 0xfe, 0x1d, 0x47, 0xfa, 0xa2, 0x4a, 0x4e, 0x15, 0x42, 0x6f, 0x92, 0x6a,
	0x5f, 0x73, 0xa8, 0xac, 0x16, 0x83, 0xd7, 0x50, 0xde, 0xb5, 0x13, 0x8b, 0xcf, 0xc7, 0x91, 0x5e,
	0x49, 0x3c, 0x64, 0x13, 0x9a, 0x57, 0x6e, 0x99, 0xb0, 0x20, 0x74, 0xb9, 0x97, 0xb8, 0x35, 0xe3,
	0x96, 0x44, 0x20, 0x34, 0x4d, 0xc1, 0xeb, 0xa8, 0x28, 0x2e, 0x7c, 0x06, 0x96, 0xac, 0x74, 0x17,
	0xa7, 0xd3, 0x96, 0x51, 0x42, 0x41, 0x94, 0x48, 0x8b, 0x7b, 0x82, 0x79, 0xca, 0x80, 0xb7, 0x90,
	0x89, 0x40, 0x68, 0x9a, 0x22, 0x91, 0xb0, 0xc0, 0xd2, 0xcf, 0x48, 0x58, 0xa0, 0x5a, 0x5d, 0xcd,
	0x41, 0x85, 0x7d, 0x6f, 0x82, 0xeb, 0xa8, 0x30, 0x60, 0x17, 0x49, 0x33, 0x99, 0x65, 0x0c, 0xd8,
	0x05, 0xa1, 0x52, 0xc2, 0x1d, 0x54, 0x9a, 0x98, 0xc3, 0x31, 0x83, 0x66, 0xaa, 0xed, 0x15, 0x43,
	0x7d, 0x21, 0x8c, 0xf4, 0x0b, 0x61, 0xbc, 0x97, 0x6a, 0x77, 0x29, 0x8e, 0xf4, 0xb9, 0xa4, 0x49,
	0x19, 0x20, 0x54, 0x95, 0xd5, 0xce, 0x50, 0x59, 0xad, 0xe3, 0xcf, 0xdf, 0x45, 0x3a, 0x68, 0xed,
	0x17, 0xbe, 0x08, 0x7d, 0xee, 0x85, 0xec, 0x81, 0xd5, 0x91, 0xcf, 0x79, 0xb4, 0x71, 0x2f, 0xe0,
	0x43, 0x60, 0xfa, 0x3e, 0x0b, 0xe4, 0x88, 0x2d, 0x6e, 0x33, 0x20, 0x95, 0xb2, 0x23, 0x96, 0x51,
	0x42, 0x41, 0xc4, 0xcf, 0x51, 0x55, 0xfe, 0xdd, 0xff, 0xe4, 0x0f, 0x4d, 0x37, 0x35, 0xc3, 0x4a,
	0x1c, 0xe9, 0x78, 0x9a, 0x9b, 0x88, 0x84, 0x66, 0x53, 0xf1, 0xff, 0xa8, 0xc4, 0x82, 0x80, 0x07,
	0x89, 0x2b, 0x32, 0xfd, 0x42, 0x98, 0x50, 0x25, 0xe3, 0x43, 0x54, 0xb4, 0x4d, 0x61, 0x82, 0x29,
	0xaa, 0xed, 0xad, 0xdf, 0xff, 0x37, 0x51, 0xdd, 0x64, 0x1f, 0x2e, 0x39, 0x84, 0x02, 0xae, 0x57,
	0x86, 0x79, 0x6f, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0xd2, 0x45, 0xe4, 0xa6, 0x06, 0x00,
	0x00,
}