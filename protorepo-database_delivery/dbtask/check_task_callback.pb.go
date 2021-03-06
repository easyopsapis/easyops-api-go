// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: check_task_callback.proto

package dbtask

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	easy_command "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_command"
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
//CheckTaskCallback请求
type CheckTaskCallbackRequest struct {
	//
	//回调的参数详情
	Command              *easy_command.TaskDetail `protobuf:"bytes,1,opt,name=command,proto3" json:"command" form:"command"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CheckTaskCallbackRequest) Reset()         { *m = CheckTaskCallbackRequest{} }
func (m *CheckTaskCallbackRequest) String() string { return proto.CompactTextString(m) }
func (*CheckTaskCallbackRequest) ProtoMessage()    {}
func (*CheckTaskCallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_328fa6081d272f04, []int{0}
}
func (m *CheckTaskCallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTaskCallbackRequest.Unmarshal(m, b)
}
func (m *CheckTaskCallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTaskCallbackRequest.Marshal(b, m, deterministic)
}
func (m *CheckTaskCallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTaskCallbackRequest.Merge(m, src)
}
func (m *CheckTaskCallbackRequest) XXX_Size() int {
	return xxx_messageInfo_CheckTaskCallbackRequest.Size(m)
}
func (m *CheckTaskCallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTaskCallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTaskCallbackRequest proto.InternalMessageInfo

func (m *CheckTaskCallbackRequest) GetCommand() *easy_command.TaskDetail {
	if m != nil {
		return m.Command
	}
	return nil
}

//
//CheckTaskCallback返回
type CheckTaskCallbackResponse struct {
	//
	//回调处理结果
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckTaskCallbackResponse) Reset()         { *m = CheckTaskCallbackResponse{} }
func (m *CheckTaskCallbackResponse) String() string { return proto.CompactTextString(m) }
func (*CheckTaskCallbackResponse) ProtoMessage()    {}
func (*CheckTaskCallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_328fa6081d272f04, []int{1}
}
func (m *CheckTaskCallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTaskCallbackResponse.Unmarshal(m, b)
}
func (m *CheckTaskCallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTaskCallbackResponse.Marshal(b, m, deterministic)
}
func (m *CheckTaskCallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTaskCallbackResponse.Merge(m, src)
}
func (m *CheckTaskCallbackResponse) XXX_Size() int {
	return xxx_messageInfo_CheckTaskCallbackResponse.Size(m)
}
func (m *CheckTaskCallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTaskCallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTaskCallbackResponse proto.InternalMessageInfo

func (m *CheckTaskCallbackResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//CheckTaskCallbackApi返回
type CheckTaskCallbackResponseWrapper struct {
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
	Data                 *CheckTaskCallbackResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CheckTaskCallbackResponseWrapper) Reset()         { *m = CheckTaskCallbackResponseWrapper{} }
func (m *CheckTaskCallbackResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CheckTaskCallbackResponseWrapper) ProtoMessage()    {}
func (*CheckTaskCallbackResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_328fa6081d272f04, []int{2}
}
func (m *CheckTaskCallbackResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTaskCallbackResponseWrapper.Unmarshal(m, b)
}
func (m *CheckTaskCallbackResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTaskCallbackResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CheckTaskCallbackResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTaskCallbackResponseWrapper.Merge(m, src)
}
func (m *CheckTaskCallbackResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CheckTaskCallbackResponseWrapper.Size(m)
}
func (m *CheckTaskCallbackResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTaskCallbackResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTaskCallbackResponseWrapper proto.InternalMessageInfo

func (m *CheckTaskCallbackResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CheckTaskCallbackResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CheckTaskCallbackResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CheckTaskCallbackResponseWrapper) GetData() *CheckTaskCallbackResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckTaskCallbackRequest)(nil), "dbtask.CheckTaskCallbackRequest")
	proto.RegisterType((*CheckTaskCallbackResponse)(nil), "dbtask.CheckTaskCallbackResponse")
	proto.RegisterType((*CheckTaskCallbackResponseWrapper)(nil), "dbtask.CheckTaskCallbackResponseWrapper")
}

func init() { proto.RegisterFile("check_task_callback.proto", fileDescriptor_328fa6081d272f04) }

var fileDescriptor_328fa6081d272f04 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0xd5, 0x7b, 0xdb, 0x5e, 0x5d, 0x97, 0x5f, 0x2f, 0x50, 0xda, 0x4d, 0x8a, 0x91, 0x10,
	0x2c, 0xea, 0x48, 0xb0, 0x41, 0x2c, 0x5b, 0xe8, 0x03, 0x44, 0x48, 0xb0, 0xab, 0x1c, 0xc7, 0x4d,
	0xab, 0xfc, 0x4c, 0xb0, 0x1d, 0x09, 0x5e, 0x36, 0xef, 0x40, 0x9e, 0x00, 0x65, 0x9c, 0x8a, 0x2e,
	0xe8, 0x2a, 0x99, 0x39, 0xf3, 0x9d, 0x39, 0xf2, 0x90, 0xb1, 0xdc, 0x28, 0x99, 0xae, 0xac, 0x30,
	0xe9, 0x4a, 0x8a, 0x2c, 0x8b, 0x84, 0x4c, 0x79, 0xa9, 0xc1, 0x02, 0x1d, 0xc6, 0x51, 0xdb, 0x9e,
	0xcc, 0x92, 0xad, 0xdd, 0x54, 0x11, 0x97, 0x90, 0x07, 0x09, 0x24, 0x10, 0xa0, 0x1c, 0x55, 0x6b,
	0xac, 0xb0, 0xc0, 0x3f, 0x87, 0x4d, 0xde, 0x12, 0xe0, 0x4a, 0x98, 0x4f, 0x28, 0x0d, 0xcf, 0x40,
	0x8a, 0x2c, 0x90, 0x50, 0x58, 0x2d, 0xa4, 0x35, 0x8e, 0xd4, 0xaa, 0x84, 0x59, 0x0e, 0xb1, 0xca,
	0x4c, 0xd0, 0x0d, 0x06, 0x58, 0x62, 0xb5, 0x92, 0x90, 0xe7, 0xa2, 0x88, 0x03, 0x8c, 0x14, 0x2b,
	0x2b, 0xb6, 0x99, 0x73, 0x66, 0x11, 0xf1, 0x16, 0x6d, 0xda, 0x17, 0x61, 0xd2, 0x45, 0x97, 0x35,
	0x54, 0xef, 0x95, 0x32, 0x96, 0x2e, 0xc9, 0xbf, 0x0e, 0xf4, 0x7a, 0xd3, 0xde, 0xcd, 0xe8, 0xce,
	0xe3, 0xfb, 0x6e, 0xbc, 0x65, 0x9e, 0xd0, 0x6c, 0x4e, 0x9b, 0xda, 0x3f, 0x59, 0x83, 0xce, 0x1f,
	0x59, 0xa7, 0xb2, 0x70, 0x07, 0xb3, 0x25, 0x19, 0xff, 0xb2, 0xc3, 0x94, 0x50, 0x18, 0x45, 0x6f,
	0xc9, 0xd0, 0x58, 0x61, 0x2b, 0x83, 0x3b, 0xfe, 0xcf, 0xcf, 0x9b, 0xda, 0x3f, 0x76, 0x4e, 0xae,
	0xcf, 0xc2, 0x6e, 0x80, 0x7d, 0xf5, 0xc8, 0xf4, 0xa0, 0xd1, 0xab, 0x16, 0x65, 0xa9, 0x34, 0xbd,
	0x22, 0x7d, 0x09, 0xb1, 0x42, 0xb7, 0xc1, 0xfc, 0xb4, 0xa9, 0xfd, 0xd1, 0x2e, 0x57, 0xac, 0x58,
	0x88, 0x22, 0x7d, 0x20, 0xa3, 0xf6, 0xfb, 0xfc, 0x51, 0x66, 0x62, 0x5b, 0x78, 0x7f, 0x70, 0xf3,
	0x45, 0x53, 0xfb, 0xf4, 0x67, 0xb6, 0x13, 0x59, 0xb8, 0x3f, 0x4a, 0xaf, 0xc9, 0x40, 0x69, 0x0d,
	0xda, 0xfb, 0x8b, 0xcc, 0x59, 0x53, 0xfb, 0x47, 0x8e, 0xc1, 0x36, 0x0b, 0x9d, 0x4c, 0x97, 0xa4,
	0x1f, 0x0b, 0x2b, 0xbc, 0x3e, 0x3e, 0xdc, 0x25, 0x77, 0x77, 0xe7, 0x07, 0xe3, 0xef, 0x27, 0x6d,
	0x41, 0x16, 0x22, 0x1f, 0x0d, 0xf1, 0x4c, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x46,
	0x5b, 0x47, 0x54, 0x02, 0x00, 0x00,
}
