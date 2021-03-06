// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_configuration.proto

package configuration

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
//UpdateConfiguration请求
type UpdateConfigurationRequest struct {
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
	ClazzName string `protobuf:"bytes,14,opt,name=clazzName,proto3" json:"clazzName" form:"clazzName"`
	//
	//关联模型ID
	ObjectId string `protobuf:"bytes,15,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//关联模型实例ID宏
	InstanceIdMacro      string   `protobuf:"bytes,16,opt,name=instanceIdMacro,proto3" json:"instanceIdMacro" form:"instanceIdMacro"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigurationRequest) Reset()         { *m = UpdateConfigurationRequest{} }
func (m *UpdateConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigurationRequest) ProtoMessage()    {}
func (*UpdateConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5454cc811fd73736, []int{0}
}
func (m *UpdateConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationRequest.Unmarshal(m, b)
}
func (m *UpdateConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationRequest.Merge(m, src)
}
func (m *UpdateConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationRequest.Size(m)
}
func (m *UpdateConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationRequest proto.InternalMessageInfo

func (m *UpdateConfigurationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateConfigurationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateConfigurationRequest) GetKwargs() *types.Value {
	if m != nil {
		return m.Kwargs
	}
	return nil
}

func (m *UpdateConfigurationRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *UpdateConfigurationRequest) GetEnv() *types.Value {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *UpdateConfigurationRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *UpdateConfigurationRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateConfigurationRequest) GetScript() *collector_center.Script {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *UpdateConfigurationRequest) GetIgnoreInvalid() bool {
	if m != nil {
		return m.IgnoreInvalid
	}
	return false
}

func (m *UpdateConfigurationRequest) GetTargetRange() *collector_center.TargetRange {
	if m != nil {
		return m.TargetRange
	}
	return nil
}

func (m *UpdateConfigurationRequest) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *UpdateConfigurationRequest) GetCacheTtl() int32 {
	if m != nil {
		return m.CacheTtl
	}
	return 0
}

func (m *UpdateConfigurationRequest) GetTimeRange() string {
	if m != nil {
		return m.TimeRange
	}
	return ""
}

func (m *UpdateConfigurationRequest) GetClazzName() string {
	if m != nil {
		return m.ClazzName
	}
	return ""
}

func (m *UpdateConfigurationRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *UpdateConfigurationRequest) GetInstanceIdMacro() string {
	if m != nil {
		return m.InstanceIdMacro
	}
	return ""
}

//
//UpdateConfiguration返回
type UpdateConfigurationResponse struct {
	//
	//ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigurationResponse) Reset()         { *m = UpdateConfigurationResponse{} }
func (m *UpdateConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigurationResponse) ProtoMessage()    {}
func (*UpdateConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5454cc811fd73736, []int{1}
}
func (m *UpdateConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationResponse.Unmarshal(m, b)
}
func (m *UpdateConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationResponse.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationResponse.Merge(m, src)
}
func (m *UpdateConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationResponse.Size(m)
}
func (m *UpdateConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationResponse proto.InternalMessageInfo

func (m *UpdateConfigurationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//UpdateConfigurationApi返回
type UpdateConfigurationResponseWrapper struct {
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
	Data                 *UpdateConfigurationResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UpdateConfigurationResponseWrapper) Reset()         { *m = UpdateConfigurationResponseWrapper{} }
func (m *UpdateConfigurationResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigurationResponseWrapper) ProtoMessage()    {}
func (*UpdateConfigurationResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5454cc811fd73736, []int{2}
}
func (m *UpdateConfigurationResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigurationResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateConfigurationResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigurationResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateConfigurationResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigurationResponseWrapper.Merge(m, src)
}
func (m *UpdateConfigurationResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigurationResponseWrapper.Size(m)
}
func (m *UpdateConfigurationResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigurationResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigurationResponseWrapper proto.InternalMessageInfo

func (m *UpdateConfigurationResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateConfigurationResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateConfigurationResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateConfigurationResponseWrapper) GetData() *UpdateConfigurationResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateConfigurationRequest)(nil), "configuration.UpdateConfigurationRequest")
	proto.RegisterType((*UpdateConfigurationResponse)(nil), "configuration.UpdateConfigurationResponse")
	proto.RegisterType((*UpdateConfigurationResponseWrapper)(nil), "configuration.UpdateConfigurationResponseWrapper")
}

func init() { proto.RegisterFile("update_configuration.proto", fileDescriptor_5454cc811fd73736) }

var fileDescriptor_5454cc811fd73736 = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xe3, 0x34,
	0x14, 0xde, 0x76, 0xda, 0x4e, 0xeb, 0x4e, 0x7f, 0x30, 0xcb, 0xc8, 0x14, 0x50, 0x2a, 0xaf, 0x40,
	0x5d, 0x44, 0x13, 0x54, 0x7e, 0xb4, 0xac, 0x04, 0x12, 0x5d, 0xb8, 0x98, 0x0b, 0x16, 0xc9, 0x2c,
	0xcc, 0xc5, 0x6a, 0x34, 0x72, 0x1d, 0x37, 0x1b, 0x36, 0x8d, 0x83, 0xe3, 0x74, 0x60, 0x47, 0xf3,
	0x88, 0xbc, 0x42, 0x46, 0x9a, 0x47, 0xc8, 0x35, 0x17, 0x28, 0x76, 0xd2, 0xa4, 0x33, 0x68, 0x24,
	0xa4, 0xbd, 0xaa, 0x7d, 0xbe, 0x9f, 0x9e, 0x73, 0x7c, 0xec, 0x80, 0x49, 0x12, 0xb9, 0x54, 0xf1,
	0x73, 0x26, 0xc2, 0xb5, 0xef, 0x25, 0x92, 0x2a, 0x5f, 0x84, 0x76, 0x24, 0x85, 0x12, 0x70, 0xb0,
	0x17, 0x9c, 0xcc, 0x3d, 0x5f, 0xbd, 0x4a, 0x56, 0x36, 0x13, 0x1b, 0xc7, 0x13, 0x9e, 0x70, 0x34,
	0x6b, 0x95, 0xac, 0xf5, 0x4e, 0x6f, 0xf4, 0xca, 0xa8, 0x27, 0xa7, 0x9e, 0xb0, 0x39, 0x8d, 0xff,
	0x12, 0x51, 0x6c, 0x07, 0x82, 0xd1, 0xc0, 0x61, 0x22, 0x54, 0x92, 0x32, 0x15, 0x1b, 0xa5, 0xe4,
	0x91, 0x98, 0x6f, 0x84, 0xcb, 0x83, 0xd8, 0x29, 0x88, 0x8e, 0xde, 0x3a, 0x4c, 0x04, 0x01, 0x67,
	0x4a, 0xc8, 0x73, 0xc6, 0x43, 0xc5, 0xa5, 0x13, 0x33, 0xe9, 0x47, 0xaa, 0x30, 0x3e, 0x7b, 0x8b,
	0xc6, 0x8a, 0x4a, 0x8f, 0xab, 0x73, 0x49, 0x43, 0x8f, 0x17, 0xf6, 0x5f, 0xd7, 0xca, 0xdc, 0x5c,
	0xf8, 0xea, 0xb5, 0xb8, 0x70, 0x3c, 0x31, 0xd7, 0xe0, 0x7c, 0x4b, 0x03, 0xdf, 0xa5, 0x4a, 0xc8,
	0xd8, 0xd9, 0x2d, 0x0b, 0xdd, 0x87, 0x9e, 0x10, 0x5e, 0xc0, 0xab, 0xae, 0xc4, 0x4a, 0x26, 0xac,
	0x48, 0x1a, 0xff, 0x7d, 0x08, 0x26, 0xbf, 0xea, 0x56, 0x3f, 0xab, 0x37, 0x95, 0xf0, 0x3f, 0x12,
	0x1e, 0x2b, 0xe8, 0x80, 0xa6, 0xef, 0xa2, 0xc6, 0xb4, 0x31, 0xeb, 0x2d, 0xad, 0x2c, 0xb5, 0x7a,
	0x6b, 0x21, 0x37, 0x4f, 0xb1, 0xef, 0xe2, 0x9b, 0x6b, 0x6b, 0x08, 0x8e, 0x5e, 0x7e, 0x3e, 0xff,
	0x86, 0xce, 0xd7, 0x67, 0x97, 0x8b, 0x2f, 0xaf, 0x48, 0xd3, 0x77, 0xe1, 0x63, 0xd0, 0x0a, 0xe9,
	0x86, 0xa3, 0xa6, 0x96, 0xbc, 0x97, 0xa5, 0x56, 0xdf, 0x48, 0xf2, 0x68, 0x2e, 0x6a, 0x46, 0x0f,
	0x88, 0xa6, 0xc0, 0xef, 0x41, 0xe7, 0xf5, 0x05, 0x95, 0x5e, 0x8c, 0x0e, 0xa6, 0x8d, 0x59, 0x7f,
	0x71, 0x6c, 0x9b, 0x4c, 0xed, 0x32, 0x53, 0xfb, 0x37, 0x1a, 0x24, 0x7c, 0xf9, 0x4e, 0x96, 0x5a,
	0x03, 0x63, 0x62, 0xf8, 0x98, 0x14, 0x42, 0xf8, 0x19, 0x38, 0x54, 0xfe, 0x86, 0x8b, 0x44, 0xa1,
	0xd6, 0xb4, 0x31, 0x6b, 0x2f, 0x61, 0x96, 0x5a, 0x43, 0xc3, 0x2d, 0x00, 0x4c, 0x4a, 0x0a, 0x7c,
	0x02, 0x0e, 0x78, 0xb8, 0x45, 0xed, 0x7b, 0xff, 0x6d, 0x98, 0xa5, 0x16, 0x30, 0x0e, 0x3c, 0xdc,
	0x62, 0x92, 0x4b, 0xa0, 0x03, 0xba, 0xae, 0x1f, 0xd3, 0x55, 0xc0, 0x5d, 0xd4, 0x99, 0x36, 0x66,
	0xdd, 0xe5, 0xbb, 0x59, 0x6a, 0x8d, 0x0c, 0xad, 0x44, 0x30, 0xd9, 0x91, 0xe0, 0x63, 0xd0, 0x09,
	0xe8, 0x8a, 0x07, 0x31, 0x3a, 0x9c, 0x1e, 0xcc, 0x7a, 0xf5, 0x1a, 0x4c, 0x1c, 0x93, 0x82, 0x00,
	0x9f, 0x81, 0x8e, 0x19, 0x23, 0xd4, 0xd5, 0x89, 0x21, 0xfb, 0xf6, 0x34, 0xd8, 0xbf, 0x68, 0xbc,
	0x6e, 0x62, 0x14, 0x98, 0x14, 0x52, 0xf8, 0x1d, 0x18, 0xf8, 0x5e, 0x28, 0x24, 0x3f, 0x09, 0xf5,
	0xf9, 0xa3, 0x9e, 0xce, 0x12, 0x65, 0xa9, 0xf5, 0xb0, 0x38, 0xb2, 0x3a, 0x8c, 0xc9, 0x3e, 0x1d,
	0x9e, 0x82, 0xbe, 0x19, 0x39, 0x92, 0x4f, 0x1c, 0x02, 0x3a, 0x93, 0x8f, 0xee, 0x66, 0xf2, 0xa2,
	0x22, 0x2d, 0x8f, 0xb3, 0xd4, 0x82, 0x45, 0xaf, 0xab, 0x30, 0x26, 0x75, 0x27, 0xf8, 0x15, 0xe8,
	0xfa, 0xb9, 0x72, 0x4b, 0x03, 0xd4, 0xd7, 0x47, 0xf4, 0x7e, 0xd5, 0xb9, 0x12, 0xd1, 0x73, 0x31,
	0x7e, 0x40, 0x76, 0xd4, 0xbc, 0xe1, 0x8c, 0xb2, 0x57, 0xfc, 0x85, 0x0a, 0xd0, 0x91, 0x96, 0xd5,
	0x1a, 0x5e, 0x22, 0x98, 0xec, 0x48, 0xf0, 0x25, 0xe8, 0xe5, 0xc7, 0x6c, 0xd2, 0x1f, 0xe8, 0xe1,
	0xfb, 0x36, 0x4b, 0xad, 0x71, 0x35, 0x0b, 0x26, 0xbb, 0x9b, 0x6b, 0xeb, 0x63, 0xf0, 0x28, 0x1f,
	0xdb, 0xb3, 0xcb, 0xc5, 0xd5, 0xd3, 0x72, 0x31, 0xbf, 0x13, 0x21, 0x95, 0x1f, 0x5c, 0x80, 0x1e,
	0x0b, 0xe8, 0x9b, 0x37, 0xcf, 0xf3, 0xc9, 0x1e, 0x6a, 0xf3, 0x87, 0x95, 0xf9, 0x0e, 0xc2, 0xa4,
	0xa2, 0xe5, 0x15, 0x88, 0xd5, 0xef, 0x9c, 0xa9, 0x13, 0x17, 0x8d, 0xb4, 0xa4, 0x56, 0x41, 0x89,
	0x60, 0xb2, 0x23, 0xc1, 0x1f, 0xc0, 0xc8, 0x0f, 0x63, 0x45, 0x43, 0xc6, 0x4f, 0xdc, 0x9f, 0x28,
	0x93, 0x02, 0x8d, 0xb5, 0x6e, 0x92, 0xa5, 0xd6, 0x71, 0xd9, 0xb0, 0x3d, 0x02, 0x26, 0xb7, 0x25,
	0xf8, 0x39, 0xf8, 0xe0, 0x3f, 0xaf, 0x73, 0x1c, 0x89, 0x30, 0xe6, 0xff, 0xfb, 0x3e, 0xe3, 0x7f,
	0x1a, 0x00, 0xdf, 0x63, 0x78, 0x2a, 0x69, 0x14, 0x71, 0x09, 0x1f, 0x81, 0x16, 0x13, 0x2e, 0xd7,
	0xce, 0xed, 0xe5, 0xa8, 0xba, 0xf6, 0x79, 0x14, 0x13, 0x0d, 0xc2, 0x27, 0xa0, 0x9f, 0xff, 0xfe,
	0xf8, 0x67, 0x14, 0x50, 0x3f, 0x2c, 0x9e, 0x88, 0xda, 0x14, 0xd5, 0x40, 0x4c, 0xea, 0x54, 0xf8,
	0x09, 0x68, 0x73, 0x29, 0x85, 0xd4, 0x2f, 0x45, 0x6f, 0x39, 0xce, 0x52, 0xeb, 0xa8, 0xb8, 0xa3,
	0x79, 0x18, 0x13, 0x03, 0xc3, 0x9f, 0x41, 0xcb, 0xa5, 0x8a, 0xea, 0xc7, 0xa0, 0xbf, 0xf8, 0xd4,
	0xde, 0xff, 0x7a, 0xdc, 0x53, 0x47, 0x3d, 0xe5, 0xdc, 0x01, 0x13, 0x6d, 0xb4, 0xea, 0xe8, 0xd7,
	0xe1, 0x8b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0x79, 0x1b, 0xdf, 0x8f, 0x06, 0x00, 0x00,
}
