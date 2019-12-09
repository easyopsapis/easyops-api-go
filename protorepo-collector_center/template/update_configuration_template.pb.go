// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_configuration_template.proto

package template

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	collector_center "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/collector_center"
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
//UpdateConfigurationTemplate请求
type UpdateConfigurationTemplateRequest struct {
	//
	//ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//参数
	Kwargs *types.Value `protobuf:"bytes,3,opt,name=kwargs,proto3" json:"kwargs" form:"kwargs"`
	//
	//超时时间
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout" form:"timeout"`
	//
	//环境变量
	Env *types.Value `protobuf:"bytes,5,opt,name=env,proto3" json:"env" form:"env"`
	//
	//是否禁用
	Disabled bool `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled" form:"disabled"`
	//
	//标签
	Labels []string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels" form:"labels"`
	//
	//脚本
	Script *collector_center.Script `protobuf:"bytes,8,opt,name=script,proto3" json:"script" form:"script"`
	//
	//是否忽略不符合要求数据
	IgnoreInvalid bool `protobuf:"varint,9,opt,name=ignoreInvalid,proto3" json:"ignoreInvalid" form:"ignoreInvalid"`
	//
	//目标动态范围
	TargetRange *collector_center.TargetRange `protobuf:"bytes,10,opt,name=targetRange,proto3" json:"targetRange" form:"targetRange"`
	//
	//采集间隔
	Interval int32 `protobuf:"varint,11,opt,name=interval,proto3" json:"interval" form:"interval"`
	//
	//缓存过期次数（时间）
	CacheTtl int32 `protobuf:"varint,12,opt,name=cacheTtl,proto3" json:"cacheTtl" form:"cacheTtl"`
	//
	//执行时间范围（24小时）
	TimeRange string `protobuf:"bytes,13,opt,name=timeRange,proto3" json:"timeRange" form:"timeRange"`
	//
	//采集类型名称
	ClazzName            string   `protobuf:"bytes,14,opt,name=clazzName,proto3" json:"clazzName" form:"clazzName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigurationTemplateRequest) Reset()         { *m = UpdateConfigurationTemplateRequest{} }
func (m *UpdateConfigurationTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigurationTemplateRequest) ProtoMessage()    {}
func (*UpdateConfigurationTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aace828b690b56b, []int{0}
}
func (m *UpdateConfigurationTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationTemplateRequest.Unmarshal(m, b)
}
func (m *UpdateConfigurationTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationTemplateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationTemplateRequest.Merge(m, src)
}
func (m *UpdateConfigurationTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationTemplateRequest.Size(m)
}
func (m *UpdateConfigurationTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationTemplateRequest proto.InternalMessageInfo

func (m *UpdateConfigurationTemplateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateConfigurationTemplateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateConfigurationTemplateRequest) GetKwargs() *types.Value {
	if m != nil {
		return m.Kwargs
	}
	return nil
}

func (m *UpdateConfigurationTemplateRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *UpdateConfigurationTemplateRequest) GetEnv() *types.Value {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *UpdateConfigurationTemplateRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *UpdateConfigurationTemplateRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateConfigurationTemplateRequest) GetScript() *collector_center.Script {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *UpdateConfigurationTemplateRequest) GetIgnoreInvalid() bool {
	if m != nil {
		return m.IgnoreInvalid
	}
	return false
}

func (m *UpdateConfigurationTemplateRequest) GetTargetRange() *collector_center.TargetRange {
	if m != nil {
		return m.TargetRange
	}
	return nil
}

func (m *UpdateConfigurationTemplateRequest) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *UpdateConfigurationTemplateRequest) GetCacheTtl() int32 {
	if m != nil {
		return m.CacheTtl
	}
	return 0
}

func (m *UpdateConfigurationTemplateRequest) GetTimeRange() string {
	if m != nil {
		return m.TimeRange
	}
	return ""
}

func (m *UpdateConfigurationTemplateRequest) GetClazzName() string {
	if m != nil {
		return m.ClazzName
	}
	return ""
}

//
//UpdateConfigurationTemplate返回
type UpdateConfigurationTemplateResponse struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigurationTemplateResponse) Reset()         { *m = UpdateConfigurationTemplateResponse{} }
func (m *UpdateConfigurationTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigurationTemplateResponse) ProtoMessage()    {}
func (*UpdateConfigurationTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aace828b690b56b, []int{1}
}
func (m *UpdateConfigurationTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationTemplateResponse.Unmarshal(m, b)
}
func (m *UpdateConfigurationTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationTemplateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationTemplateResponse.Merge(m, src)
}
func (m *UpdateConfigurationTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationTemplateResponse.Size(m)
}
func (m *UpdateConfigurationTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationTemplateResponse proto.InternalMessageInfo

func (m *UpdateConfigurationTemplateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//UpdateConfigurationTemplateApi返回
type UpdateConfigurationTemplateResponseWrapper struct {
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
	Data                 *UpdateConfigurationTemplateResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *UpdateConfigurationTemplateResponseWrapper) Reset() {
	*m = UpdateConfigurationTemplateResponseWrapper{}
}
func (m *UpdateConfigurationTemplateResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*UpdateConfigurationTemplateResponseWrapper) ProtoMessage() {}
func (*UpdateConfigurationTemplateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aace828b690b56b, []int{2}
}
func (m *UpdateConfigurationTemplateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateConfigurationTemplateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationTemplateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper.Merge(m, src)
}
func (m *UpdateConfigurationTemplateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper.Size(m)
}
func (m *UpdateConfigurationTemplateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationTemplateResponseWrapper proto.InternalMessageInfo

func (m *UpdateConfigurationTemplateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateConfigurationTemplateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateConfigurationTemplateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateConfigurationTemplateResponseWrapper) GetData() *UpdateConfigurationTemplateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateConfigurationTemplateRequest)(nil), "template.UpdateConfigurationTemplateRequest")
	proto.RegisterType((*UpdateConfigurationTemplateResponse)(nil), "template.UpdateConfigurationTemplateResponse")
	proto.RegisterType((*UpdateConfigurationTemplateResponseWrapper)(nil), "template.UpdateConfigurationTemplateResponseWrapper")
}

func init() {
	proto.RegisterFile("update_configuration_template.proto", fileDescriptor_0aace828b690b56b)
}

var fileDescriptor_0aace828b690b56b = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x4f, 0xe3, 0x38,
	0x18, 0xa6, 0xa5, 0x2d, 0xad, 0x0b, 0x85, 0xf5, 0xb2, 0x28, 0x8b, 0x76, 0x95, 0xca, 0xd5, 0xae,
	0xca, 0x6a, 0x93, 0xac, 0xba, 0x1f, 0x62, 0x91, 0x76, 0xa5, 0x29, 0x9a, 0xc3, 0x5c, 0xe6, 0xe0,
	0x61, 0xe0, 0x80, 0x50, 0xe5, 0x26, 0x6e, 0x88, 0x70, 0xe3, 0x8c, 0xe3, 0x94, 0x19, 0x10, 0xe7,
	0xf9, 0x97, 0x41, 0xe2, 0x27, 0xe4, 0x17, 0x8c, 0x62, 0xa7, 0x4d, 0x66, 0x90, 0x10, 0x23, 0xcd,
	0xa9, 0xb6, 0x9f, 0x8f, 0xbc, 0x7e, 0xfd, 0xbc, 0x05, 0x83, 0x24, 0xf2, 0x88, 0xa4, 0x13, 0x97,
	0x87, 0xb3, 0xc0, 0x4f, 0x04, 0x91, 0x01, 0x0f, 0x27, 0x92, 0xce, 0x23, 0x46, 0x24, 0xb5, 0x23,
	0xc1, 0x25, 0x87, 0xed, 0xe5, 0x7e, 0xdf, 0xf2, 0x03, 0x79, 0x99, 0x4c, 0x6d, 0x97, 0xcf, 0x1d,
	0x9f, 0xfb, 0xdc, 0x51, 0x84, 0x69, 0x32, 0x53, 0x3b, 0xb5, 0x51, 0x2b, 0x2d, 0xdc, 0x3f, 0xf3,
	0xb9, 0x4d, 0x49, 0xfc, 0x81, 0x47, 0xb1, 0xcd, 0xb8, 0x4b, 0x98, 0xe3, 0xf2, 0x50, 0x0a, 0xe2,
	0xca, 0x58, 0x2b, 0x05, 0x8d, 0xb8, 0x35, 0xe7, 0x1e, 0x65, 0xb1, 0x53, 0x10, 0x1d, 0xb5, 0x75,
	0x5c, 0xce, 0x18, 0x75, 0x25, 0x17, 0x13, 0x97, 0x86, 0x92, 0x0a, 0x27, 0x76, 0x45, 0x10, 0xc9,
	0xc2, 0xf8, 0xe2, 0x1b, 0x1a, 0x4b, 0x22, 0x7c, 0x2a, 0x27, 0x82, 0x84, 0x7e, 0x71, 0xe1, 0xfd,
	0x7f, 0x2a, 0xd7, 0x9c, 0x5f, 0x07, 0xf2, 0x8a, 0x5f, 0x3b, 0x3e, 0xb7, 0x14, 0x68, 0x2d, 0x08,
	0x0b, 0x3c, 0x22, 0xb9, 0x88, 0x9d, 0xd5, 0xb2, 0xd0, 0xfd, 0xe4, 0x73, 0xee, 0x33, 0x5a, 0x76,
	0x25, 0x96, 0x22, 0x71, 0x8b, 0xa2, 0x51, 0xda, 0x02, 0xe8, 0xad, 0x6a, 0xf7, 0x71, 0xb5, 0xdb,
	0x27, 0x45, 0x73, 0x31, 0x7d, 0x97, 0xd0, 0x58, 0x42, 0x07, 0xd4, 0x03, 0xcf, 0xa8, 0xf5, 0x6b,
	0xc3, 0xce, 0xd8, 0xcc, 0x52, 0xb3, 0x33, 0xe3, 0x62, 0x7e, 0x84, 0x02, 0x0f, 0x3d, 0xdc, 0x9b,
	0x3d, 0xb0, 0x79, 0xfe, 0x87, 0xf5, 0x2f, 0xb1, 0x66, 0x17, 0xb7, 0xa3, 0xbf, 0xee, 0x70, 0x3d,
	0xf0, 0xe0, 0x01, 0x68, 0x84, 0x64, 0x4e, 0x8d, 0xba, 0x92, 0xfc, 0x90, 0xa5, 0x66, 0x57, 0x4b,
	0xf2, 0xd3, 0x5c, 0x54, 0x8f, 0xd6, 0xb0, 0xa2, 0xc0, 0x17, 0xa0, 0x75, 0x75, 0x4d, 0x84, 0x1f,
	0x1b, 0xeb, 0xfd, 0xda, 0xb0, 0x3b, 0xda, 0xb3, 0x75, 0xc5, 0xf6, 0xb2, 0x62, 0xfb, 0x94, 0xb0,
	0x84, 0x8e, 0xbf, 0xcb, 0x52, 0x73, 0x4b, 0x9b, 0x68, 0x3e, 0xc2, 0x85, 0x10, 0xfe, 0x0e, 0x36,
	0x64, 0x30, 0xa7, 0x3c, 0x91, 0x46, 0xa3, 0x5f, 0x1b, 0x36, 0xc7, 0x30, 0x4b, 0xcd, 0x9e, 0xe6,
	0x16, 0x00, 0xc2, 0x4b, 0x0a, 0x3c, 0x04, 0xeb, 0x34, 0x5c, 0x18, 0xcd, 0x27, 0xbf, 0xd6, 0xcb,
	0x52, 0x13, 0x68, 0x07, 0x1a, 0x2e, 0x10, 0xce, 0x25, 0xd0, 0x01, 0x6d, 0x2f, 0x88, 0xc9, 0x94,
	0x51, 0xcf, 0x68, 0xf5, 0x6b, 0xc3, 0xf6, 0xf8, 0xfb, 0x2c, 0x35, 0xb7, 0x35, 0x6d, 0x89, 0x20,
	0xbc, 0x22, 0xc1, 0x03, 0xd0, 0x62, 0x64, 0x4a, 0x59, 0x6c, 0x6c, 0xf4, 0xd7, 0x87, 0x9d, 0xea,
	0x1d, 0xf4, 0x39, 0xc2, 0x05, 0x01, 0x1e, 0x83, 0x96, 0x8e, 0x93, 0xd1, 0x56, 0x85, 0x19, 0xf6,
	0x97, 0xa9, 0xb0, 0xdf, 0x28, 0xbc, 0x6a, 0xa2, 0x15, 0x08, 0x17, 0x52, 0xf8, 0x3f, 0xd8, 0x0a,
	0xfc, 0x90, 0x0b, 0xfa, 0x2a, 0x54, 0x39, 0x30, 0x3a, 0xaa, 0x4a, 0x23, 0x4b, 0xcd, 0xdd, 0xe2,
	0xc9, 0xaa, 0x30, 0xc2, 0x9f, 0xd3, 0xe1, 0x19, 0xe8, 0xea, 0xe8, 0xe1, 0x3c, 0x79, 0x06, 0x50,
	0x95, 0xfc, 0xfc, 0xb8, 0x92, 0x93, 0x92, 0x34, 0xde, 0xcb, 0x52, 0x13, 0x16, 0xbd, 0x2e, 0x8f,
	0x11, 0xae, 0x3a, 0xc1, 0xbf, 0x41, 0x3b, 0xc8, 0x95, 0x0b, 0xc2, 0x8c, 0xae, 0x7a, 0xa2, 0x1f,
	0xcb, 0xce, 0x2d, 0x11, 0x95, 0x8b, 0x9d, 0x35, 0xbc, 0xa2, 0xe6, 0x0d, 0x77, 0x89, 0x7b, 0x49,
	0x4f, 0x24, 0x33, 0x36, 0x95, 0xac, 0xd2, 0xf0, 0x25, 0x82, 0xf0, 0x8a, 0x04, 0xcf, 0x41, 0x27,
	0x7f, 0x66, 0x5d, 0xfe, 0x96, 0x0a, 0xdf, 0x7f, 0x59, 0x6a, 0xee, 0x94, 0x59, 0xd0, 0xd5, 0x3d,
	0xdc, 0x9b, 0xbf, 0x80, 0x41, 0x1e, 0xdb, 0x8b, 0xdb, 0xd1, 0xdd, 0xd1, 0x72, 0x61, 0x3d, 0x3a,
	0xc1, 0xa5, 0x1f, 0x1c, 0x81, 0x8e, 0xcb, 0xc8, 0xcd, 0xcd, 0xeb, 0x3c, 0xd9, 0x3d, 0x65, 0xbe,
	0x5b, 0x9a, 0xaf, 0x20, 0x84, 0x4b, 0x1a, 0x3a, 0x05, 0x83, 0x27, 0xe7, 0x2b, 0x8e, 0x78, 0x18,
	0xd3, 0xaf, 0x1e, 0x30, 0xf4, 0xb1, 0x0e, 0x7e, 0x7b, 0x86, 0xf1, 0x99, 0x20, 0x51, 0x44, 0x05,
	0x1c, 0x80, 0x86, 0xcb, 0x3d, 0xaa, 0xbe, 0xd0, 0x1c, 0x6f, 0x97, 0xf3, 0x98, 0x9f, 0x22, 0xac,
	0x40, 0x78, 0x08, 0xba, 0xf9, 0xef, 0xcb, 0xf7, 0x11, 0x23, 0x41, 0x58, 0xcc, 0x6e, 0xe5, 0x79,
	0x2b, 0x20, 0xc2, 0x55, 0x2a, 0xfc, 0x15, 0x34, 0xa9, 0x10, 0x5c, 0xa8, 0x11, 0xee, 0x8c, 0x77,
	0xb2, 0xd4, 0xdc, 0x2c, 0x86, 0x27, 0x3f, 0x46, 0x58, 0xc3, 0x10, 0x83, 0x86, 0x47, 0x24, 0x51,
	0x53, 0xda, 0x1d, 0x59, 0xf6, 0xea, 0x4f, 0xfd, 0x19, 0x57, 0xa9, 0x56, 0x9d, 0x9b, 0x20, 0xac,
	0xbc, 0xa6, 0x2d, 0x35, 0xb9, 0x7f, 0x7e, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xce, 0x1a, 0x96, 0xdd,
	0x37, 0x06, 0x00, 0x00,
}