// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extra_info.proto

package tool

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
//执行的额外信息
type ExtraInfo struct {
	//
	//工具名称
	ToolName string `protobuf:"bytes,1,opt,name=toolName,proto3" json:"toolName" form:"toolName"`
	//
	//运行模式
	ExecMode string `protobuf:"bytes,2,opt,name=execMode,proto3" json:"execMode" form:"execMode"`
	//
	//输出信息
	Outputs []string `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs" form:"outputs"`
	//
	//请求源地址
	Origin string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin" form:"origin"`
	//
	//触发器名称
	Trigger string `protobuf:"bytes,5,opt,name=trigger,proto3" json:"trigger" form:"trigger"`
	//
	//详细信息
	Detail *ExtraInfo_Detail `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail" form:"detail"`
	//
	//工具Id，服务端自动生成
	ToolId string `protobuf:"bytes,7,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	//
	//执行用户
	ExecUser string `protobuf:"bytes,8,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//工具输入参数
	Inputs               *types.Struct `protobuf:"bytes,9,opt,name=inputs,proto3" json:"inputs" form:"inputs"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExtraInfo) Reset()         { *m = ExtraInfo{} }
func (m *ExtraInfo) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo) ProtoMessage()    {}
func (*ExtraInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0}
}
func (m *ExtraInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo.Unmarshal(m, b)
}
func (m *ExtraInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo.Marshal(b, m, deterministic)
}
func (m *ExtraInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo.Merge(m, src)
}
func (m *ExtraInfo) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo.Size(m)
}
func (m *ExtraInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo proto.InternalMessageInfo

func (m *ExtraInfo) GetToolName() string {
	if m != nil {
		return m.ToolName
	}
	return ""
}

func (m *ExtraInfo) GetExecMode() string {
	if m != nil {
		return m.ExecMode
	}
	return ""
}

func (m *ExtraInfo) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *ExtraInfo) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ExtraInfo) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *ExtraInfo) GetDetail() *ExtraInfo_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *ExtraInfo) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func (m *ExtraInfo) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *ExtraInfo) GetInputs() *types.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type ExtraInfo_Detail struct {
	//
	//回调信息
	Callback *Callback `protobuf:"bytes,1,opt,name=callback,proto3" json:"callback" form:"callback"`
	//
	//工具的输出, 输出为空为[],非空时为object, 结构如下:
	//fields:
	//- name: dimensions
	//type: object[]
	//description: 维度列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//- name: columns
	//type: object[]
	//description: 输出列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//
	ToolOutputs *types.Value `protobuf:"bytes,2,opt,name=toolOutputs,proto3" json:"toolOutputs" form:"toolOutputs"`
	//
	//工具的Env
	ToolEnv *types.Struct `protobuf:"bytes,3,opt,name=toolEnv,proto3" json:"toolEnv" form:"toolEnv"`
	//
	//工具的输出定义
	OutputDefs []*ExtraInfo_Detail_OutputDefs `protobuf:"bytes,4,rep,name=outputDefs,proto3" json:"outputDefs" form:"outputDefs"`
	//
	//工具的表格输出定义
	TableDefs []*ExtraInfo_Detail_TableDefs `protobuf:"bytes,5,rep,name=tableDefs,proto3" json:"tableDefs" form:"tableDefs"`
	//
	//邮件订阅用户
	Subscribers          []string `protobuf:"bytes,6,rep,name=subscribers,proto3" json:"subscribers" form:"subscribers"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraInfo_Detail) Reset()         { *m = ExtraInfo_Detail{} }
func (m *ExtraInfo_Detail) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo_Detail) ProtoMessage()    {}
func (*ExtraInfo_Detail) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0, 0}
}
func (m *ExtraInfo_Detail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo_Detail.Unmarshal(m, b)
}
func (m *ExtraInfo_Detail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo_Detail.Marshal(b, m, deterministic)
}
func (m *ExtraInfo_Detail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo_Detail.Merge(m, src)
}
func (m *ExtraInfo_Detail) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo_Detail.Size(m)
}
func (m *ExtraInfo_Detail) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo_Detail.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo_Detail proto.InternalMessageInfo

func (m *ExtraInfo_Detail) GetCallback() *Callback {
	if m != nil {
		return m.Callback
	}
	return nil
}

func (m *ExtraInfo_Detail) GetToolOutputs() *types.Value {
	if m != nil {
		return m.ToolOutputs
	}
	return nil
}

func (m *ExtraInfo_Detail) GetToolEnv() *types.Struct {
	if m != nil {
		return m.ToolEnv
	}
	return nil
}

func (m *ExtraInfo_Detail) GetOutputDefs() []*ExtraInfo_Detail_OutputDefs {
	if m != nil {
		return m.OutputDefs
	}
	return nil
}

func (m *ExtraInfo_Detail) GetTableDefs() []*ExtraInfo_Detail_TableDefs {
	if m != nil {
		return m.TableDefs
	}
	return nil
}

func (m *ExtraInfo_Detail) GetSubscribers() []string {
	if m != nil {
		return m.Subscribers
	}
	return nil
}

type ExtraInfo_Detail_OutputDefs struct {
	//
	//输出参数ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出参数标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraInfo_Detail_OutputDefs) Reset()         { *m = ExtraInfo_Detail_OutputDefs{} }
func (m *ExtraInfo_Detail_OutputDefs) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo_Detail_OutputDefs) ProtoMessage()    {}
func (*ExtraInfo_Detail_OutputDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0, 0, 0}
}
func (m *ExtraInfo_Detail_OutputDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo_Detail_OutputDefs.Unmarshal(m, b)
}
func (m *ExtraInfo_Detail_OutputDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo_Detail_OutputDefs.Marshal(b, m, deterministic)
}
func (m *ExtraInfo_Detail_OutputDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo_Detail_OutputDefs.Merge(m, src)
}
func (m *ExtraInfo_Detail_OutputDefs) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo_Detail_OutputDefs.Size(m)
}
func (m *ExtraInfo_Detail_OutputDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo_Detail_OutputDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo_Detail_OutputDefs proto.InternalMessageInfo

func (m *ExtraInfo_Detail_OutputDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExtraInfo_Detail_OutputDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ExtraInfo_Detail_TableDefs struct {
	//
	//输出表格id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出表格标题
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列
	Dimensions []*ExtraInfo_Detail_TableDefs_Dimensions `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//输出列
	Columns              []*ExtraInfo_Detail_TableDefs_Columns `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ExtraInfo_Detail_TableDefs) Reset()         { *m = ExtraInfo_Detail_TableDefs{} }
func (m *ExtraInfo_Detail_TableDefs) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo_Detail_TableDefs) ProtoMessage()    {}
func (*ExtraInfo_Detail_TableDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0, 0, 1}
}
func (m *ExtraInfo_Detail_TableDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs.Unmarshal(m, b)
}
func (m *ExtraInfo_Detail_TableDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs.Marshal(b, m, deterministic)
}
func (m *ExtraInfo_Detail_TableDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs.Merge(m, src)
}
func (m *ExtraInfo_Detail_TableDefs) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs.Size(m)
}
func (m *ExtraInfo_Detail_TableDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo_Detail_TableDefs proto.InternalMessageInfo

func (m *ExtraInfo_Detail_TableDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExtraInfo_Detail_TableDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtraInfo_Detail_TableDefs) GetDimensions() []*ExtraInfo_Detail_TableDefs_Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *ExtraInfo_Detail_TableDefs) GetColumns() []*ExtraInfo_Detail_TableDefs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ExtraInfo_Detail_TableDefs_Dimensions struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraInfo_Detail_TableDefs_Dimensions) Reset()         { *m = ExtraInfo_Detail_TableDefs_Dimensions{} }
func (m *ExtraInfo_Detail_TableDefs_Dimensions) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo_Detail_TableDefs_Dimensions) ProtoMessage()    {}
func (*ExtraInfo_Detail_TableDefs_Dimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0, 0, 1, 0}
}
func (m *ExtraInfo_Detail_TableDefs_Dimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions.Unmarshal(m, b)
}
func (m *ExtraInfo_Detail_TableDefs_Dimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions.Marshal(b, m, deterministic)
}
func (m *ExtraInfo_Detail_TableDefs_Dimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions.Merge(m, src)
}
func (m *ExtraInfo_Detail_TableDefs_Dimensions) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions.Size(m)
}
func (m *ExtraInfo_Detail_TableDefs_Dimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo_Detail_TableDefs_Dimensions proto.InternalMessageInfo

func (m *ExtraInfo_Detail_TableDefs_Dimensions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExtraInfo_Detail_TableDefs_Dimensions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ExtraInfo_Detail_TableDefs_Columns struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraInfo_Detail_TableDefs_Columns) Reset()         { *m = ExtraInfo_Detail_TableDefs_Columns{} }
func (m *ExtraInfo_Detail_TableDefs_Columns) String() string { return proto.CompactTextString(m) }
func (*ExtraInfo_Detail_TableDefs_Columns) ProtoMessage()    {}
func (*ExtraInfo_Detail_TableDefs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b508dafc886a9a, []int{0, 0, 1, 1}
}
func (m *ExtraInfo_Detail_TableDefs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns.Unmarshal(m, b)
}
func (m *ExtraInfo_Detail_TableDefs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns.Marshal(b, m, deterministic)
}
func (m *ExtraInfo_Detail_TableDefs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns.Merge(m, src)
}
func (m *ExtraInfo_Detail_TableDefs_Columns) XXX_Size() int {
	return xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns.Size(m)
}
func (m *ExtraInfo_Detail_TableDefs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraInfo_Detail_TableDefs_Columns proto.InternalMessageInfo

func (m *ExtraInfo_Detail_TableDefs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExtraInfo_Detail_TableDefs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtraInfo)(nil), "tool.ExtraInfo")
	proto.RegisterType((*ExtraInfo_Detail)(nil), "tool.ExtraInfo.Detail")
	proto.RegisterType((*ExtraInfo_Detail_OutputDefs)(nil), "tool.ExtraInfo.Detail.OutputDefs")
	proto.RegisterType((*ExtraInfo_Detail_TableDefs)(nil), "tool.ExtraInfo.Detail.TableDefs")
	proto.RegisterType((*ExtraInfo_Detail_TableDefs_Dimensions)(nil), "tool.ExtraInfo.Detail.TableDefs.Dimensions")
	proto.RegisterType((*ExtraInfo_Detail_TableDefs_Columns)(nil), "tool.ExtraInfo.Detail.TableDefs.Columns")
}

func init() { proto.RegisterFile("extra_info.proto", fileDescriptor_62b508dafc886a9a) }

var fileDescriptor_62b508dafc886a9a = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x4e, 0xdc, 0x46,
	0x18, 0x15, 0xbb, 0x8b, 0x97, 0x1d, 0x8b, 0xbf, 0xa1, 0xa5, 0xd6, 0xaa, 0x95, 0x97, 0x81, 0x8b,
	0xad, 0x8a, 0xed, 0x76, 0x91, 0xaa, 0x96, 0x8b, 0x16, 0x16, 0xa8, 0xc4, 0x05, 0x05, 0x4d, 0x29,
	0xaa, 0x40, 0x14, 0xf9, 0x67, 0xd6, 0xb1, 0xf0, 0x7a, 0x56, 0xfe, 0x01, 0x12, 0xc4, 0x1b, 0xe5,
	0x55, 0xf2, 0x0a, 0x8e, 0x14, 0x29, 0x2f, 0xe0, 0x27, 0x88, 0xe6, 0xc7, 0x5e, 0x07, 0x25, 0xe1,
	0x22, 0xdc, 0xcd, 0xcc, 0x77, 0xce, 0xf9, 0x66, 0xce, 0x9c, 0xb1, 0xc1, 0x12, 0xb9, 0x4b, 0x63,
	0xfb, 0x2a, 0x88, 0x46, 0xd4, 0x9c, 0xc4, 0x34, 0xa5, 0xb0, 0x95, 0x52, 0x1a, 0x76, 0x0d, 0x3f,
	0x48, 0x5f, 0x64, 0x8e, 0xe9, 0xd2, 0xb1, 0xe5, 0x53, 0x9f, 0x5a, 0xbc, 0xe8, 0x64, 0x23, 0x3e,
	0xe3, 0x13, 0x3e, 0x12, 0xa4, 0xee, 0x91, 0x4f, 0x4d, 0x62, 0x27, 0x2f, 0xe9, 0x24, 0x31, 0x43,
	0xea, 0xda, 0xa1, 0xe5, 0xd2, 0x28, 0x8d, 0x6d, 0x37, 0x4d, 0x04, 0x33, 0x26, 0x13, 0x6a, 0x8c,
	0xa9, 0x47, 0xc2, 0xc4, 0x92, 0x40, 0x8b, 0x4f, 0x2d, 0xd6, 0xce, 0x72, 0xed, 0x30, 0x74, 0x6c,
	0xf7, 0x5a, 0xca, 0xfd, 0x5a, 0xeb, 0x3e, 0xbe, 0x0d, 0xd2, 0x6b, 0x7a, 0x6b, 0xf9, 0xd4, 0xe0,
	0x45, 0xe3, 0xc6, 0x0e, 0x03, 0xcf, 0x4e, 0x69, 0x9c, 0x58, 0xd5, 0x50, 0xf2, 0xbe, 0xf7, 0x29,
	0xf5, 0x43, 0x32, 0xdd, 0x6c, 0x92, 0xc6, 0x99, 0x9b, 0x8a, 0x2a, 0x7a, 0xaf, 0x82, 0xce, 0x01,
	0x3b, 0xee, 0x61, 0x34, 0xa2, 0xd0, 0x02, 0x73, 0xac, 0xf5, 0xdf, 0xf6, 0x98, 0x68, 0x33, 0xbd,
	0x99, 0x7e, 0x67, 0xb8, 0x52, 0xe4, 0xfa, 0xe2, 0x88, 0xc6, 0xe3, 0x6d, 0x54, 0x56, 0x10, 0xae,
	0x40, 0x8c, 0x40, 0xee, 0x88, 0x7b, 0x44, 0x3d, 0xa2, 0x35, 0x1e, 0x13, 0xca, 0x0a, 0xc2, 0x15,
	0x08, 0x6e, 0x82, 0x36, 0xcd, 0xd2, 0x49, 0x96, 0x26, 0x5a, 0xb3, 0xd7, 0xec, 0x77, 0x86, 0xb0,
	0xc8, 0xf5, 0x05, 0x81, 0x97, 0x05, 0x84, 0x4b, 0x08, 0xfc, 0x11, 0x28, 0x34, 0x0e, 0xfc, 0x20,
	0xd2, 0x5a, 0x5c, 0x7c, 0xb9, 0xc8, 0xf5, 0x79, 0x09, 0xe6, 0xeb, 0x08, 0x4b, 0x00, 0x13, 0x4e,
	0xe3, 0xc0, 0xf7, 0x49, 0xac, 0xcd, 0x72, 0x6c, 0x4d, 0x58, 0x16, 0x10, 0x2e, 0x21, 0x70, 0x17,
	0x28, 0x1e, 0x49, 0xed, 0x20, 0xd4, 0x94, 0xde, 0x4c, 0x5f, 0x1d, 0xac, 0x9a, 0xec, 0x48, 0x66,
	0xe5, 0x84, 0xb9, 0xcf, 0xab, 0xf5, 0x86, 0x02, 0x8f, 0xb0, 0x24, 0xc2, 0x1d, 0xa0, 0x30, 0xce,
	0xa1, 0xa7, 0xb5, 0x79, 0xbf, 0xfe, 0x14, 0x2a, 0xd6, 0xd1, 0xbb, 0xb7, 0xfa, 0x0a, 0x58, 0xfe,
	0xff, 0xc2, 0x36, 0x46, 0xbb, 0xc6, 0x5f, 0x3f, 0x1b, 0xbf, 0x5f, 0xde, 0x6f, 0x0d, 0x1e, 0x36,
	0xb0, 0xe4, 0xc1, 0x33, 0x61, 0xde, 0xbf, 0x09, 0x89, 0xb5, 0x39, 0xae, 0xb1, 0xfd, 0xb1, 0x79,
	0xac, 0xc2, 0x54, 0xd6, 0xc1, 0x1a, 0x53, 0x79, 0xb5, 0x6b, 0x9c, 0x33, 0x95, 0x0b, 0xb3, 0x1a,
	0x5f, 0x19, 0x97, 0xf7, 0x83, 0xcd, 0xad, 0x5f, 0x1e, 0x36, 0x70, 0xa5, 0x05, 0x87, 0x40, 0x09,
	0x22, 0x6e, 0x71, 0x87, 0x1f, 0xee, 0x3b, 0x53, 0x44, 0xc0, 0x2c, 0x23, 0x60, 0xfe, 0xc3, 0x23,
	0x50, 0x3f, 0x9d, 0x20, 0x20, 0x2c, 0x99, 0xdd, 0x37, 0x6d, 0xa0, 0x08, 0x0f, 0xe0, 0x9f, 0x60,
	0xae, 0x8c, 0x22, 0x0f, 0x85, 0x3a, 0x58, 0x10, 0x6e, 0xed, 0xc9, 0xd5, 0xfa, 0x9d, 0x97, 0x48,
	0x84, 0x2b, 0x12, 0x3c, 0x01, 0x2a, 0xc3, 0x1f, 0xcb, 0x7b, 0x6f, 0x48, 0xc7, 0x1f, 0x6f, 0xea,
	0xcc, 0x0e, 0x33, 0x32, 0x5c, 0x2d, 0x72, 0x1d, 0x4e, 0x6d, 0x3c, 0x2e, 0x33, 0x51, 0x97, 0x80,
	0x07, 0xa0, 0xcd, 0xa6, 0x07, 0xd1, 0x8d, 0xd6, 0xfc, 0xf2, 0x11, 0xeb, 0x29, 0x10, 0x0c, 0x96,
	0x02, 0x31, 0x82, 0xff, 0x01, 0x20, 0x92, 0xb6, 0x4f, 0x46, 0x89, 0xd6, 0xea, 0x35, 0xfb, 0xea,
	0x60, 0xed, 0xd3, 0x49, 0x30, 0x8f, 0x2b, 0xe0, 0xf0, 0xdb, 0x22, 0xd7, 0x97, 0xeb, 0x91, 0x65,
	0xab, 0x08, 0xd7, 0xb4, 0xe0, 0x29, 0xe8, 0xa4, 0xb6, 0x13, 0x12, 0x2e, 0x3c, 0xcb, 0x85, 0x7b,
	0x9f, 0x11, 0x3e, 0x2d, 0x71, 0xc3, 0x6f, 0x8a, 0x5c, 0x5f, 0x92, 0x7b, 0x2d, 0x17, 0x11, 0x9e,
	0x0a, 0xc1, 0xdf, 0x80, 0x9a, 0x64, 0x4e, 0xe2, 0xc6, 0x81, 0x43, 0xe2, 0x44, 0x53, 0xf8, 0x03,
	0xaa, 0x19, 0x56, 0x2b, 0x22, 0x5c, 0x87, 0x76, 0x4f, 0x00, 0x98, 0x1e, 0x00, 0xfe, 0x00, 0x1a,
	0x81, 0x27, 0x1f, 0xf8, 0x7c, 0x91, 0xeb, 0x1d, 0x99, 0x01, 0x0f, 0xe1, 0x46, 0xe0, 0xc1, 0x75,
	0xd0, 0x8a, 0xd8, 0x17, 0x40, 0x3c, 0xe8, 0xc5, 0x22, 0xd7, 0x55, 0x01, 0x88, 0xf8, 0xeb, 0xe7,
	0xc5, 0xee, 0xeb, 0x26, 0xe8, 0x54, 0x5b, 0x7f, 0x0e, 0x45, 0xe8, 0x00, 0xe0, 0x05, 0x63, 0x12,
	0x25, 0x01, 0x8d, 0xc4, 0xd7, 0x41, 0x1d, 0xfc, 0xf4, 0x94, 0x69, 0xe6, 0x7e, 0x45, 0xa9, 0xdf,
	0xcb, 0x54, 0x08, 0xe1, 0x9a, 0x2a, 0x3c, 0x03, 0x6d, 0x97, 0x86, 0xd9, 0x38, 0x2a, 0xaf, 0xbb,
	0xff, 0x64, 0x83, 0x3d, 0x81, 0xaf, 0x27, 0x49, 0x4a, 0x20, 0x5c, 0x8a, 0x31, 0x7f, 0xa7, 0x1b,
	0x79, 0x16, 0x7f, 0x8f, 0x40, 0x5b, 0x76, 0x7e, 0x0e, 0xb9, 0xe1, 0xce, 0xf9, 0x1f, 0x5f, 0xf7,
	0x3b, 0x72, 0x14, 0x0e, 0xda, 0xfa, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x2b, 0x4e, 0x93, 0x1e,
	0x07, 0x00, 0x00,
}