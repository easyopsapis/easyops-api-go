// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: probe.proto

package container

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Pod 健康检查
type Probe struct {
	//
	//探测方式
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" form:"type"`
	//
	//exec
	Exec *Probe_Exec `protobuf:"bytes,2,opt,name=exec,proto3" json:"exec" form:"exec"`
	//
	//httpGet 探测方式
	HttpGet *Probe_HttpGet `protobuf:"bytes,3,opt,name=httpGet,proto3" json:"httpGet" form:"httpGet"`
	//
	//tcpSocket
	TcpSocket *Probe_TcpSocket `protobuf:"bytes,4,opt,name=tcpSocket,proto3" json:"tcpSocket" form:"tcpSocket"`
	//
	//容器启动完成后首次探测时间, 单位s
	InitialDelaySeconds int32 `protobuf:"varint,5,opt,name=initialDelaySeconds,proto3" json:"initialDelaySeconds" form:"initialDelaySeconds"`
	//
	//对容器进行健康检查探测等待响应时间, 默认1s
	TimeoutSeconds int32 `protobuf:"varint,6,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds" form:"timeoutSeconds"`
	//
	//对容器健康检查检查间隔，默认10s一次
	PeriodSeconds int32 `protobuf:"varint,7,opt,name=periodSeconds,proto3" json:"periodSeconds" form:"periodSeconds"`
	//
	//探测成功多少次后，认为成功
	SuccessThreshold int32 `protobuf:"varint,8,opt,name=successThreshold,proto3" json:"successThreshold" form:"successThreshold"`
	//
	//探测失败多少次后，认为失败
	FailureThreshold     int32    `protobuf:"varint,9,opt,name=failureThreshold,proto3" json:"failureThreshold" form:"failureThreshold"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}
func (*Probe) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0}
}
func (m *Probe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe.Unmarshal(m, b)
}
func (m *Probe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe.Marshal(b, m, deterministic)
}
func (m *Probe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe.Merge(m, src)
}
func (m *Probe) XXX_Size() int {
	return xxx_messageInfo_Probe.Size(m)
}
func (m *Probe) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe.DiscardUnknown(m)
}

var xxx_messageInfo_Probe proto.InternalMessageInfo

func (m *Probe) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Probe) GetExec() *Probe_Exec {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *Probe) GetHttpGet() *Probe_HttpGet {
	if m != nil {
		return m.HttpGet
	}
	return nil
}

func (m *Probe) GetTcpSocket() *Probe_TcpSocket {
	if m != nil {
		return m.TcpSocket
	}
	return nil
}

func (m *Probe) GetInitialDelaySeconds() int32 {
	if m != nil {
		return m.InitialDelaySeconds
	}
	return 0
}

func (m *Probe) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *Probe) GetPeriodSeconds() int32 {
	if m != nil {
		return m.PeriodSeconds
	}
	return 0
}

func (m *Probe) GetSuccessThreshold() int32 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

func (m *Probe) GetFailureThreshold() int32 {
	if m != nil {
		return m.FailureThreshold
	}
	return 0
}

type Probe_Exec struct {
	//
	//exec 探测方式，检查命令
	Command              []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command" form:"command"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe_Exec) Reset()         { *m = Probe_Exec{} }
func (m *Probe_Exec) String() string { return proto.CompactTextString(m) }
func (*Probe_Exec) ProtoMessage()    {}
func (*Probe_Exec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0, 0}
}
func (m *Probe_Exec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe_Exec.Unmarshal(m, b)
}
func (m *Probe_Exec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe_Exec.Marshal(b, m, deterministic)
}
func (m *Probe_Exec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe_Exec.Merge(m, src)
}
func (m *Probe_Exec) XXX_Size() int {
	return xxx_messageInfo_Probe_Exec.Size(m)
}
func (m *Probe_Exec) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe_Exec.DiscardUnknown(m)
}

var xxx_messageInfo_Probe_Exec proto.InternalMessageInfo

func (m *Probe_Exec) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

type Probe_HttpGet struct {
	//
	//uri
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" form:"path"`
	//
	//端口
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port" form:"port"`
	//
	//域名
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host" form:"host"`
	//
	//http协议
	Schema string `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema" form:"schema"`
	//
	//http 请求头
	HttpHeaders          []*Probe_HttpGet_HttpHeaders `protobuf:"bytes,5,rep,name=httpHeaders,proto3" json:"httpHeaders" form:"httpHeaders"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Probe_HttpGet) Reset()         { *m = Probe_HttpGet{} }
func (m *Probe_HttpGet) String() string { return proto.CompactTextString(m) }
func (*Probe_HttpGet) ProtoMessage()    {}
func (*Probe_HttpGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0, 1}
}
func (m *Probe_HttpGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe_HttpGet.Unmarshal(m, b)
}
func (m *Probe_HttpGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe_HttpGet.Marshal(b, m, deterministic)
}
func (m *Probe_HttpGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe_HttpGet.Merge(m, src)
}
func (m *Probe_HttpGet) XXX_Size() int {
	return xxx_messageInfo_Probe_HttpGet.Size(m)
}
func (m *Probe_HttpGet) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe_HttpGet.DiscardUnknown(m)
}

var xxx_messageInfo_Probe_HttpGet proto.InternalMessageInfo

func (m *Probe_HttpGet) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Probe_HttpGet) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Probe_HttpGet) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Probe_HttpGet) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Probe_HttpGet) GetHttpHeaders() []*Probe_HttpGet_HttpHeaders {
	if m != nil {
		return m.HttpHeaders
	}
	return nil
}

type Probe_HttpGet_HttpHeaders struct {
	//
	//变量名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//变量值
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe_HttpGet_HttpHeaders) Reset()         { *m = Probe_HttpGet_HttpHeaders{} }
func (m *Probe_HttpGet_HttpHeaders) String() string { return proto.CompactTextString(m) }
func (*Probe_HttpGet_HttpHeaders) ProtoMessage()    {}
func (*Probe_HttpGet_HttpHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0, 1, 0}
}
func (m *Probe_HttpGet_HttpHeaders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe_HttpGet_HttpHeaders.Unmarshal(m, b)
}
func (m *Probe_HttpGet_HttpHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe_HttpGet_HttpHeaders.Marshal(b, m, deterministic)
}
func (m *Probe_HttpGet_HttpHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe_HttpGet_HttpHeaders.Merge(m, src)
}
func (m *Probe_HttpGet_HttpHeaders) XXX_Size() int {
	return xxx_messageInfo_Probe_HttpGet_HttpHeaders.Size(m)
}
func (m *Probe_HttpGet_HttpHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe_HttpGet_HttpHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_Probe_HttpGet_HttpHeaders proto.InternalMessageInfo

func (m *Probe_HttpGet_HttpHeaders) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Probe_HttpGet_HttpHeaders) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Probe_TcpSocket struct {
	//
	//通过tcpSocket探测的端口
	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port" form:"port"`
	//
	//hostname, 默认 pod IP
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host" form:"host"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe_TcpSocket) Reset()         { *m = Probe_TcpSocket{} }
func (m *Probe_TcpSocket) String() string { return proto.CompactTextString(m) }
func (*Probe_TcpSocket) ProtoMessage()    {}
func (*Probe_TcpSocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0, 2}
}
func (m *Probe_TcpSocket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe_TcpSocket.Unmarshal(m, b)
}
func (m *Probe_TcpSocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe_TcpSocket.Marshal(b, m, deterministic)
}
func (m *Probe_TcpSocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe_TcpSocket.Merge(m, src)
}
func (m *Probe_TcpSocket) XXX_Size() int {
	return xxx_messageInfo_Probe_TcpSocket.Size(m)
}
func (m *Probe_TcpSocket) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe_TcpSocket.DiscardUnknown(m)
}

var xxx_messageInfo_Probe_TcpSocket proto.InternalMessageInfo

func (m *Probe_TcpSocket) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Probe_TcpSocket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*Probe)(nil), "container.Probe")
	proto.RegisterType((*Probe_Exec)(nil), "container.Probe.Exec")
	proto.RegisterType((*Probe_HttpGet)(nil), "container.Probe.HttpGet")
	proto.RegisterType((*Probe_HttpGet_HttpHeaders)(nil), "container.Probe.HttpGet.HttpHeaders")
	proto.RegisterType((*Probe_TcpSocket)(nil), "container.Probe.TcpSocket")
}

func init() { proto.RegisterFile("probe.proto", fileDescriptor_f8cc8551cf1a5c4f) }

var fileDescriptor_f8cc8551cf1a5c4f = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x4f, 0xdb, 0x3a,
	0x18, 0x56, 0x69, 0x4b, 0x4f, 0xdc, 0x03, 0xf4, 0xf8, 0xc0, 0x39, 0x59, 0x26, 0x2d, 0x95, 0x87,
	0xa6, 0x14, 0xd6, 0x46, 0x62, 0x13, 0x17, 0xd5, 0xc4, 0x44, 0x36, 0x06, 0xbb, 0x99, 0x90, 0xe1,
	0x0a, 0x42, 0x25, 0x37, 0x35, 0x4d, 0x44, 0x52, 0x47, 0x89, 0xcb, 0xc7, 0xc4, 0x05, 0xff, 0x6c,
	0xd7, 0xfb, 0x13, 0x99, 0xb4, 0x9f, 0x90, 0x3f, 0xb0, 0xc9, 0x4e, 0xfa, 0xcd, 0x76, 0xb1, 0xde,
	0x34, 0x7e, 0x9f, 0x0f, 0x3b, 0x6f, 0x9e, 0xd7, 0xa0, 0x1a, 0x46, 0xac, 0x4b, 0x5b, 0x61, 0xc4,
	0x38, 0x83, 0x8a, 0xc3, 0x06, 0x9c, 0x78, 0x03, 0x1a, 0x69, 0xcd, 0xbe, 0xc7, 0xdd, 0x61, 0xb7,
	0xe5, 0xb0, 0xc0, 0xec, 0xb3, 0x3e, 0x33, 0x25, 0xa3, 0x3b, 0xbc, 0x94, 0x2b, 0xb9, 0x90, 0x4f,
	0x99, 0x52, 0xdb, 0x9d, 0xa2, 0x07, 0x37, 0x1e, 0xbf, 0x62, 0x37, 0x66, 0x9f, 0x35, 0x25, 0xd8,
	0xbc, 0x26, 0xbe, 0xd7, 0x23, 0x9c, 0x45, 0xb1, 0x39, 0x7e, 0xcc, 0x74, 0xe8, 0xab, 0x02, 0xca,
	0xc7, 0xe2, 0x04, 0xf0, 0x39, 0x28, 0xf1, 0xbb, 0x90, 0xaa, 0x85, 0x7a, 0xc1, 0x50, 0xac, 0xb5,
	0x34, 0xd1, 0xab, 0x97, 0x2c, 0x0a, 0xda, 0x48, 0x54, 0x11, 0x96, 0x20, 0x6c, 0x83, 0x12, 0xbd,
	0xa5, 0x8e, 0xba, 0x54, 0x2f, 0x18, 0xd5, 0x9d, 0x8d, 0xd6, 0xf8, 0xbc, 0x2d, 0x69, 0xd2, 0x3a,
	0xb8, 0xa5, 0xce, 0xb4, 0x56, 0x90, 0x11, 0x96, 0x1a, 0xf8, 0x01, 0x54, 0x5c, 0xce, 0xc3, 0x43,
	0xca, 0xd5, 0xa2, 0x94, 0xab, 0x0b, 0xf2, 0xa3, 0x0c, 0xb7, 0x60, 0x9a, 0xe8, 0xab, 0x99, 0x43,
	0x2e, 0x41, 0x78, 0x24, 0x86, 0x9f, 0x80, 0xc2, 0x9d, 0xf0, 0x84, 0x39, 0x57, 0x94, 0xab, 0x25,
	0xe9, 0xa4, 0x2d, 0x38, 0x9d, 0x8e, 0x18, 0xd6, 0x7a, 0x9a, 0xe8, 0xb5, 0xfc, 0x4d, 0x46, 0x45,
	0x84, 0x27, 0x16, 0xf0, 0x18, 0xfc, 0xeb, 0x0d, 0x3c, 0xee, 0x11, 0xff, 0x3d, 0xf5, 0xc9, 0xdd,
	0x09, 0x75, 0xd8, 0xa0, 0x17, 0xab, 0xe5, 0x7a, 0xc1, 0x28, 0x5b, 0xcf, 0xd2, 0x44, 0xd7, 0x32,
	0xf5, 0x23, 0x24, 0x84, 0x1f, 0x93, 0xc2, 0x7d, 0xb0, 0xca, 0xbd, 0x80, 0xb2, 0x21, 0x1f, 0x99,
	0x2d, 0x4b, 0xb3, 0x27, 0x69, 0xa2, 0x6f, 0xe4, 0x47, 0x99, 0xc1, 0x11, 0x9e, 0x13, 0xc0, 0x3d,
	0xb0, 0x12, 0xd2, 0xc8, 0x63, 0xbd, 0x91, 0x43, 0x45, 0x3a, 0xa8, 0x69, 0xa2, 0xaf, 0x67, 0x0e,
	0x33, 0x30, 0xc2, 0xb3, 0x74, 0x78, 0x08, 0x6a, 0xf1, 0xd0, 0x71, 0x68, 0x1c, 0x9f, 0xba, 0x11,
	0x8d, 0x5d, 0xe6, 0xf7, 0xd4, 0xbf, 0xa4, 0xc5, 0xd3, 0x34, 0xd1, 0xff, 0xcf, 0x2c, 0xe6, 0x19,
	0x08, 0x2f, 0x88, 0x84, 0xd1, 0x25, 0xf1, 0xfc, 0x61, 0x44, 0x27, 0x46, 0xca, 0xbc, 0xd1, 0x3c,
	0x03, 0xe1, 0x05, 0x91, 0xf6, 0x1a, 0x94, 0x44, 0x3a, 0xe0, 0x4b, 0x50, 0x71, 0x58, 0x10, 0x90,
	0x41, 0x4f, 0x2d, 0xd4, 0x8b, 0x86, 0x32, 0xfd, 0xb1, 0x73, 0x00, 0xe1, 0x11, 0x45, 0xfb, 0x52,
	0x04, 0x95, 0x3c, 0x15, 0x30, 0x04, 0xa5, 0x90, 0x70, 0x37, 0x4f, 0xa8, 0x3d, 0x49, 0x99, 0xa8,
	0xa2, 0xef, 0xdf, 0xf4, 0x8f, 0xe0, 0xb0, 0x63, 0x18, 0xb6, 0x79, 0xde, 0xb1, 0xcd, 0xb6, 0xbd,
	0x65, 0xbf, 0x45, 0xe8, 0xcd, 0x9e, 0x7d, 0x6f, 0x47, 0xf6, 0xe0, 0x62, 0xbb, 0xb1, 0xdd, 0xb8,
	0x37, 0x6c, 0xb3, 0x71, 0x7f, 0x4e, 0x9a, 0x9f, 0xf7, 0x9b, 0x67, 0x17, 0x6d, 0xc3, 0xb6, 0xcf,
	0x3b, 0xb6, 0xbd, 0xc8, 0xdc, 0xda, 0xc4, 0x72, 0x27, 0xb8, 0x0b, 0x4a, 0x21, 0x8b, 0xb8, 0x8c,
	0x7b, 0xd9, 0x42, 0x53, 0x3b, 0xb2, 0x88, 0x8b, 0x1d, 0xd7, 0x6a, 0x3f, 0x46, 0xbf, 0x82, 0xfa,
	0xf0, 0x50, 0xc2, 0x92, 0x2f, 0x66, 0xc9, 0x65, 0x71, 0x96, 0xf3, 0x99, 0x59, 0x12, 0x55, 0x84,
	0x25, 0x08, 0x1b, 0x60, 0x39, 0x76, 0x5c, 0x1a, 0x10, 0x19, 0x62, 0xc5, 0xfa, 0x27, 0x4d, 0xf4,
	0x95, 0xfc, 0xc3, 0xc8, 0x3a, 0xc2, 0x39, 0x01, 0x76, 0x40, 0x55, 0xa4, 0xff, 0x88, 0x92, 0x1e,
	0x8d, 0x44, 0x34, 0x8b, 0x46, 0x75, 0x67, 0xf3, 0x57, 0xe3, 0x23, 0xff, 0x73, 0xae, 0xf5, 0x5f,
	0x9a, 0xe8, 0x70, 0x32, 0x4a, 0x79, 0x19, 0xe1, 0x69, 0x43, 0xed, 0x0c, 0x54, 0xa7, 0x34, 0xe2,
	0xf8, 0x03, 0x12, 0x3c, 0x72, 0x15, 0x88, 0x2a, 0xc2, 0x12, 0x84, 0x2f, 0x40, 0xf9, 0x9a, 0xf8,
	0x43, 0x2a, 0x9b, 0xa3, 0x58, 0xb5, 0x34, 0xd1, 0xff, 0xce, 0x58, 0xb2, 0x8c, 0x70, 0x06, 0x6b,
	0x2e, 0x50, 0xc6, 0xc3, 0x38, 0x6e, 0x68, 0xe1, 0x0f, 0x1b, 0xba, 0xf4, 0x9b, 0x86, 0x5a, 0x07,
	0x67, 0xef, 0xfa, 0xac, 0x45, 0x49, 0x7c, 0xc7, 0xc2, 0xb8, 0xe5, 0x33, 0x87, 0xf8, 0xa6, 0x68,
	0x51, 0x44, 0x1c, 0x1e, 0x67, 0xf7, 0x67, 0x44, 0x43, 0xd6, 0x0c, 0x58, 0x8f, 0xfa, 0xb1, 0x99,
	0x13, 0x4d, 0xb9, 0x34, 0xc7, 0xbd, 0xec, 0x2e, 0x4b, 0xe6, 0xab, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0xbb, 0x0c, 0x3d, 0x9a, 0x05, 0x00, 0x00,
}
