// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: liquibase_task_callback.proto

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
//LiquibaseTaskCallback请求
type LiquibaseTaskCallbackRequest struct {
	//
	//回调的参数详情
	Command              *easy_command.TaskDetail `protobuf:"bytes,1,opt,name=command,proto3" json:"command" form:"command"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LiquibaseTaskCallbackRequest) Reset()         { *m = LiquibaseTaskCallbackRequest{} }
func (m *LiquibaseTaskCallbackRequest) String() string { return proto.CompactTextString(m) }
func (*LiquibaseTaskCallbackRequest) ProtoMessage()    {}
func (*LiquibaseTaskCallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e8d645c66fc61a, []int{0}
}
func (m *LiquibaseTaskCallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquibaseTaskCallbackRequest.Unmarshal(m, b)
}
func (m *LiquibaseTaskCallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquibaseTaskCallbackRequest.Marshal(b, m, deterministic)
}
func (m *LiquibaseTaskCallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquibaseTaskCallbackRequest.Merge(m, src)
}
func (m *LiquibaseTaskCallbackRequest) XXX_Size() int {
	return xxx_messageInfo_LiquibaseTaskCallbackRequest.Size(m)
}
func (m *LiquibaseTaskCallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquibaseTaskCallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LiquibaseTaskCallbackRequest proto.InternalMessageInfo

func (m *LiquibaseTaskCallbackRequest) GetCommand() *easy_command.TaskDetail {
	if m != nil {
		return m.Command
	}
	return nil
}

//
//LiquibaseTaskCallback返回
type LiquibaseTaskCallbackResponse struct {
	//
	//回调处理结果
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiquibaseTaskCallbackResponse) Reset()         { *m = LiquibaseTaskCallbackResponse{} }
func (m *LiquibaseTaskCallbackResponse) String() string { return proto.CompactTextString(m) }
func (*LiquibaseTaskCallbackResponse) ProtoMessage()    {}
func (*LiquibaseTaskCallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e8d645c66fc61a, []int{1}
}
func (m *LiquibaseTaskCallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquibaseTaskCallbackResponse.Unmarshal(m, b)
}
func (m *LiquibaseTaskCallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquibaseTaskCallbackResponse.Marshal(b, m, deterministic)
}
func (m *LiquibaseTaskCallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquibaseTaskCallbackResponse.Merge(m, src)
}
func (m *LiquibaseTaskCallbackResponse) XXX_Size() int {
	return xxx_messageInfo_LiquibaseTaskCallbackResponse.Size(m)
}
func (m *LiquibaseTaskCallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquibaseTaskCallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LiquibaseTaskCallbackResponse proto.InternalMessageInfo

func (m *LiquibaseTaskCallbackResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//LiquibaseTaskCallbackApi返回
type LiquibaseTaskCallbackResponseWrapper struct {
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
	Data                 *LiquibaseTaskCallbackResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *LiquibaseTaskCallbackResponseWrapper) Reset()         { *m = LiquibaseTaskCallbackResponseWrapper{} }
func (m *LiquibaseTaskCallbackResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*LiquibaseTaskCallbackResponseWrapper) ProtoMessage()    {}
func (*LiquibaseTaskCallbackResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e8d645c66fc61a, []int{2}
}
func (m *LiquibaseTaskCallbackResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper.Unmarshal(m, b)
}
func (m *LiquibaseTaskCallbackResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper.Marshal(b, m, deterministic)
}
func (m *LiquibaseTaskCallbackResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper.Merge(m, src)
}
func (m *LiquibaseTaskCallbackResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper.Size(m)
}
func (m *LiquibaseTaskCallbackResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_LiquibaseTaskCallbackResponseWrapper proto.InternalMessageInfo

func (m *LiquibaseTaskCallbackResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LiquibaseTaskCallbackResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *LiquibaseTaskCallbackResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LiquibaseTaskCallbackResponseWrapper) GetData() *LiquibaseTaskCallbackResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LiquibaseTaskCallbackRequest)(nil), "dbtask.LiquibaseTaskCallbackRequest")
	proto.RegisterType((*LiquibaseTaskCallbackResponse)(nil), "dbtask.LiquibaseTaskCallbackResponse")
	proto.RegisterType((*LiquibaseTaskCallbackResponseWrapper)(nil), "dbtask.LiquibaseTaskCallbackResponseWrapper")
}

func init() { proto.RegisterFile("liquibase_task_callback.proto", fileDescriptor_03e8d645c66fc61a) }

var fileDescriptor_03e8d645c66fc61a = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6f, 0xdb, 0x30,
	0x10, 0xc5, 0xe1, 0xd6, 0x76, 0x51, 0xba, 0x7f, 0x39, 0x14, 0x82, 0x51, 0x43, 0x05, 0xdb, 0x06,
	0xc9, 0x60, 0x0a, 0x48, 0x96, 0x20, 0xa3, 0xf3, 0x67, 0x30, 0x32, 0x09, 0x01, 0x92, 0xcd, 0xa0,
	0x28, 0x5a, 0x11, 0x4c, 0xe9, 0x64, 0x92, 0x02, 0x92, 0x2f, 0xab, 0x0f, 0xa1, 0x2d, 0x5b, 0xa0,
	0xa3, 0x8c, 0x78, 0x89, 0x27, 0xe9, 0xee, 0xdd, 0xef, 0xdd, 0x93, 0x48, 0x32, 0xd3, 0xf9, 0xb6,
	0xce, 0x13, 0x61, 0xd5, 0xca, 0x09, 0xbb, 0x59, 0x49, 0xa1, 0x75, 0x22, 0xe4, 0x86, 0x57, 0x06,
	0x1c, 0xd0, 0x71, 0x9a, 0x74, 0xed, 0xe9, 0x3c, 0xcb, 0xdd, 0x63, 0x9d, 0x70, 0x09, 0x45, 0x94,
	0x41, 0x06, 0x11, 0xca, 0x49, 0xbd, 0xc6, 0x0a, 0x0b, 0x7c, 0xf3, 0xd8, 0xf4, 0x21, 0x03, 0xae,
	0x84, 0x7d, 0x86, 0xca, 0x72, 0x0d, 0x52, 0xe8, 0x48, 0x42, 0xe9, 0x8c, 0x90, 0xce, 0x7a, 0xd2,
	0xa8, 0x0a, 0xe6, 0x05, 0xa4, 0x4a, 0xdb, 0xa8, 0x1f, 0x8c, 0xb0, 0xc4, 0x6a, 0x25, 0xa1, 0x28,
	0x44, 0x99, 0x46, 0x18, 0x29, 0x55, 0x4e, 0xe4, 0xda, 0x3b, 0xb3, 0x35, 0xf9, 0x7d, 0xbb, 0x4b,
	0x7c, 0x27, 0xec, 0xe6, 0xb2, 0xcf, 0x1b, 0xab, 0x6d, 0xad, 0xac, 0xa3, 0x37, 0xe4, 0x53, 0x0f,
	0x07, 0x83, 0x3f, 0x83, 0xe3, 0xc9, 0x69, 0xc0, 0xf7, 0x1d, 0x79, 0xc7, 0x5c, 0xa1, 0xe1, 0x82,
	0xb6, 0x4d, 0xf8, 0x6d, 0x0d, 0xa6, 0xb8, 0x60, 0xbd, 0xca, 0xe2, 0x1d, 0xcc, 0x96, 0x64, 0xf6,
	0xce, 0x1e, 0x5b, 0x41, 0x69, 0x15, 0x3d, 0x21, 0x63, 0xeb, 0x84, 0xab, 0x2d, 0xee, 0xf9, 0xbc,
	0xf8, 0xd9, 0x36, 0xe1, 0x57, 0xef, 0xe6, 0xfb, 0x2c, 0xee, 0x07, 0xd8, 0xcb, 0x80, 0xfc, 0x3b,
	0x68, 0x76, 0x6f, 0x44, 0x55, 0x29, 0x43, 0xff, 0x92, 0xa1, 0x84, 0x54, 0xa1, 0xe3, 0x68, 0xf1,
	0xbd, 0x6d, 0xc2, 0xc9, 0x2e, 0x5f, 0xaa, 0x58, 0x8c, 0x22, 0x3d, 0x27, 0x93, 0xee, 0x79, 0xfd,
	0x54, 0x69, 0x91, 0x97, 0xc1, 0x07, 0xdc, 0xfe, 0xab, 0x6d, 0x42, 0xfa, 0x36, 0xdb, 0x8b, 0x2c,
	0xde, 0x1f, 0xa5, 0x47, 0x64, 0xa4, 0x8c, 0x01, 0x13, 0x7c, 0x44, 0xe6, 0x47, 0xdb, 0x84, 0x5f,
	0x3c, 0x83, 0x6d, 0x16, 0x7b, 0x99, 0x2e, 0xc9, 0x30, 0x15, 0x4e, 0x04, 0x43, 0xfc, 0x81, 0xff,
	0xb9, 0xbf, 0x03, 0xfc, 0xe0, 0x27, 0xec, 0xa7, 0xed, 0x60, 0x16, 0xa3, 0x47, 0x32, 0xc6, 0x63,
	0x3b, 0x7b, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x37, 0x12, 0x7d, 0x68, 0x02, 0x00, 0x00,
}