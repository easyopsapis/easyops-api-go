// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_job_export.proto

package job_export

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
//GetJobExport请求
type GetJobExportRequest struct {
	//
	//menu id
	MenuId               string   `protobuf:"bytes,1,opt,name=menuId,proto3" json:"menuId" form:"menuId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobExportRequest) Reset()         { *m = GetJobExportRequest{} }
func (m *GetJobExportRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobExportRequest) ProtoMessage()    {}
func (*GetJobExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c5f43868cd5850, []int{0}
}
func (m *GetJobExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobExportRequest.Unmarshal(m, b)
}
func (m *GetJobExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobExportRequest.Marshal(b, m, deterministic)
}
func (m *GetJobExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobExportRequest.Merge(m, src)
}
func (m *GetJobExportRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobExportRequest.Size(m)
}
func (m *GetJobExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobExportRequest proto.InternalMessageInfo

func (m *GetJobExportRequest) GetMenuId() string {
	if m != nil {
		return m.MenuId
	}
	return ""
}

//
//GetJobExport返回
type GetJobExportResponse struct {
	//
	//changelog的 version，只能为 int
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version" form:"version"`
	//
	//创建者
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//导出时间 2019-07-01 20:34:45.54878400
	Time string `protobuf:"bytes,3,opt,name=time,proto3" json:"time" form:"time"`
	//
	//导出内容
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content" form:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobExportResponse) Reset()         { *m = GetJobExportResponse{} }
func (m *GetJobExportResponse) String() string { return proto.CompactTextString(m) }
func (*GetJobExportResponse) ProtoMessage()    {}
func (*GetJobExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c5f43868cd5850, []int{1}
}
func (m *GetJobExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobExportResponse.Unmarshal(m, b)
}
func (m *GetJobExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobExportResponse.Marshal(b, m, deterministic)
}
func (m *GetJobExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobExportResponse.Merge(m, src)
}
func (m *GetJobExportResponse) XXX_Size() int {
	return xxx_messageInfo_GetJobExportResponse.Size(m)
}
func (m *GetJobExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobExportResponse proto.InternalMessageInfo

func (m *GetJobExportResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetJobExportResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetJobExportResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *GetJobExportResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//
//GetJobExportApi返回
type GetJobExportResponseWrapper struct {
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
	Data                 *GetJobExportResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetJobExportResponseWrapper) Reset()         { *m = GetJobExportResponseWrapper{} }
func (m *GetJobExportResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetJobExportResponseWrapper) ProtoMessage()    {}
func (*GetJobExportResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c5f43868cd5850, []int{2}
}
func (m *GetJobExportResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobExportResponseWrapper.Unmarshal(m, b)
}
func (m *GetJobExportResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobExportResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetJobExportResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobExportResponseWrapper.Merge(m, src)
}
func (m *GetJobExportResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetJobExportResponseWrapper.Size(m)
}
func (m *GetJobExportResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobExportResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobExportResponseWrapper proto.InternalMessageInfo

func (m *GetJobExportResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetJobExportResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetJobExportResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetJobExportResponseWrapper) GetData() *GetJobExportResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetJobExportRequest)(nil), "job_export.GetJobExportRequest")
	proto.RegisterType((*GetJobExportResponse)(nil), "job_export.GetJobExportResponse")
	proto.RegisterType((*GetJobExportResponseWrapper)(nil), "job_export.GetJobExportResponseWrapper")
}

func init() { proto.RegisterFile("get_job_export.proto", fileDescriptor_65c5f43868cd5850) }

var fileDescriptor_65c5f43868cd5850 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4b, 0xab, 0xd3, 0x40,
	0x14, 0x26, 0x9a, 0x7b, 0x2f, 0x4e, 0xbd, 0x3e, 0xc6, 0x72, 0x09, 0x75, 0x91, 0x3a, 0x8a, 0x28,
	0x34, 0x0d, 0xd8, 0x52, 0xc4, 0x95, 0x16, 0xaa, 0xe8, 0x72, 0x36, 0x82, 0xaf, 0x92, 0xc7, 0x34,
	0x46, 0x9b, 0x9c, 0x38, 0x99, 0xb4, 0x45, 0xf1, 0xaf, 0x46, 0x70, 0xeb, 0x46, 0xf2, 0x07, 0x94,
	0x9c, 0x4c, 0x4c, 0x84, 0x66, 0x93, 0x99, 0xf3, 0x3d, 0xe6, 0xfb, 0xe0, 0x90, 0x61, 0x24, 0xd4,
	0xfa, 0x13, 0xf8, 0x6b, 0x71, 0xc8, 0x40, 0xaa, 0x69, 0x26, 0x41, 0x01, 0x25, 0xdd, 0x64, 0xe4,
	0x44, 0xb1, 0xfa, 0x58, 0xf8, 0xd3, 0x00, 0x12, 0x37, 0x82, 0x08, 0x5c, 0xa4, 0xf8, 0xc5, 0x06,
	0x6f, 0x78, 0xc1, 0x53, 0x23, 0x1d, 0x2d, 0x7a, 0xf4, 0x64, 0x1f, 0xab, 0xcf, 0xb0, 0x77, 0x23,
	0x70, 0x10, 0x74, 0x76, 0xde, 0x36, 0x0e, 0x3d, 0x05, 0x32, 0x77, 0xff, 0x1d, 0x1b, 0x1d, 0x7b,
	0x4a, 0x6e, 0xbd, 0x10, 0xea, 0x15, 0xf8, 0x2b, 0x7c, 0x96, 0x8b, 0x2f, 0x85, 0xc8, 0x15, 0x7d,
	0x48, 0x4e, 0x13, 0x91, 0x16, 0x2f, 0x43, 0xcb, 0x18, 0x1b, 0x0f, 0xae, 0x2c, 0x6f, 0x56, 0xa5,
	0x7d, 0xbe, 0x01, 0x99, 0x3c, 0x61, 0xcd, 0x9c, 0x71, 0x4d, 0x60, 0xbf, 0x0d, 0x32, 0xfc, 0xdf,
	0x22, 0xcf, 0x20, 0xcd, 0x05, 0x9d, 0x90, 0xb3, 0x9d, 0x90, 0x79, 0x0c, 0x29, 0x9a, 0x9c, 0x2c,
	0x69, 0x55, 0xda, 0xd7, 0x1a, 0x13, 0x0d, 0x30, 0xde, 0x52, 0xe8, 0x73, 0x72, 0x16, 0x48, 0x51,
	0x27, 0xb3, 0x2e, 0xe1, 0x93, 0x93, 0x8e, 0xad, 0x01, 0xf6, 0xf3, 0x87, 0x7d, 0x41, 0x86, 0x1f,
	0xde, 0x3e, 0x73, 0xde, 0x78, 0xce, 0xd7, 0xb5, 0xf3, 0xfe, 0xdd, 0xfe, 0xdb, 0x6c, 0xb2, 0x98,
	0x7f, 0xbf, 0xc7, 0x5b, 0x31, 0x9d, 0x13, 0x53, 0xc5, 0x89, 0xb0, 0x2e, 0xa3, 0xc9, 0xb8, 0x2a,
	0xed, 0x41, 0x63, 0x52, 0x4f, 0x6b, 0x87, 0xf3, 0xec, 0x4f, 0xfb, 0x19, 0x87, 0x3b, 0x1c, 0xd9,
	0x75, 0xd6, 0x00, 0x52, 0x25, 0x52, 0x65, 0x99, 0x28, 0xec, 0x65, 0xd5, 0x00, 0xe3, 0x2d, 0x85,
	0xfd, 0x32, 0xc8, 0xed, 0x63, 0x95, 0x5f, 0x4b, 0x2f, 0xcb, 0x84, 0xa4, 0x77, 0x89, 0x19, 0x40,
	0x28, 0x74, 0xed, 0xeb, 0x5d, 0x86, 0x7a, 0xca, 0x38, 0x82, 0xf4, 0x31, 0x19, 0xd4, 0xff, 0xd5,
	0x21, 0xdb, 0x7a, 0x71, 0xaa, 0x4b, 0x5f, 0x54, 0xa5, 0x4d, 0x3b, 0xae, 0x06, 0x19, 0xef, 0x53,
	0xe9, 0x7d, 0x72, 0x22, 0xa4, 0x04, 0xa9, 0x3b, 0xde, 0xa8, 0x4a, 0xfb, 0x6a, 0xa3, 0xc1, 0x31,
	0xe3, 0x0d, 0x4c, 0x57, 0xc4, 0x0c, 0x3d, 0xe5, 0x61, 0xa3, 0xc1, 0xa3, 0xf1, 0xb4, 0xb7, 0x6f,
	0xc7, 0xd2, 0xf7, 0x83, 0xd6, 0x3a, 0xc6, 0x51, 0xee, 0x9f, 0xe2, 0xa6, 0xcc, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x94, 0xdf, 0x88, 0xd0, 0xb4, 0x02, 0x00, 0x00,
}
