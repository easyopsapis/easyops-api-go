// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service_port.proto

package container

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
//Service 需要暴露的端口列表
type ServicePort struct {
	//
	//端口名称,
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//服务监听端口号
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port" form:"port"`
	//
	//端口协议, 默认 TCP
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol" form:"protocol"`
	//
	//需要转发到后端 Pod 的端口或者 Pod 名称, 类型为 int/string
	TargetPort *types.Value `protobuf:"bytes,4,opt,name=targetPort,proto3" json:"targetPort" form:"targetPort"`
	//
	//当 type=NodePort/LoadBalancer 时，指定映射到物理机的端口号
	NodePort             int32    `protobuf:"varint,5,opt,name=nodePort,proto3" json:"nodePort" form:"nodePort"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicePort) Reset()         { *m = ServicePort{} }
func (m *ServicePort) String() string { return proto.CompactTextString(m) }
func (*ServicePort) ProtoMessage()    {}
func (*ServicePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_759dbc9b710ea1c7, []int{0}
}
func (m *ServicePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicePort.Unmarshal(m, b)
}
func (m *ServicePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicePort.Marshal(b, m, deterministic)
}
func (m *ServicePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicePort.Merge(m, src)
}
func (m *ServicePort) XXX_Size() int {
	return xxx_messageInfo_ServicePort.Size(m)
}
func (m *ServicePort) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicePort.DiscardUnknown(m)
}

var xxx_messageInfo_ServicePort proto.InternalMessageInfo

func (m *ServicePort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServicePort) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServicePort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *ServicePort) GetTargetPort() *types.Value {
	if m != nil {
		return m.TargetPort
	}
	return nil
}

func (m *ServicePort) GetNodePort() int32 {
	if m != nil {
		return m.NodePort
	}
	return 0
}

func init() {
	proto.RegisterType((*ServicePort)(nil), "container.ServicePort")
}

func init() { proto.RegisterFile("service_port.proto", fileDescriptor_759dbc9b710ea1c7) }

var fileDescriptor_759dbc9b710ea1c7 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x7d, 0xe9, 0xa3, 0x9d, 0x2e, 0xfa, 0x5e, 0x44, 0x09, 0x45, 0x48, 0x89, 0x20, 0xdd,
	0x64, 0x06, 0x14, 0xba, 0x70, 0x19, 0x71, 0x29, 0x48, 0x04, 0x17, 0x6e, 0x64, 0x3a, 0x9d, 0x8e,
	0xc1, 0x49, 0x6e, 0x98, 0x99, 0xb4, 0xb8, 0xeb, 0x17, 0xf9, 0x49, 0x15, 0xfc, 0x84, 0xfe, 0x80,
	0x92, 0x49, 0x9a, 0x74, 0x61, 0x56, 0x73, 0xef, 0x39, 0xe7, 0xde, 0x73, 0x72, 0x91, 0xa7, 0xb9,
	0x5a, 0xa7, 0x8c, 0xbf, 0x14, 0xa0, 0x0c, 0x2e, 0x14, 0x18, 0xf0, 0x86, 0x0c, 0x72, 0x43, 0xd3,
	0x9c, 0xab, 0x49, 0x24, 0x52, 0xf3, 0x5a, 0x2e, 0x30, 0x83, 0x8c, 0x08, 0x10, 0x40, 0x2c, 0x63,
	0x51, 0xae, 0x6c, 0x65, 0x0b, 0xfb, 0xaa, 0x95, 0x93, 0xf9, 0x11, 0x3d, 0xdb, 0xa4, 0xe6, 0x0d,
	0x36, 0x44, 0x40, 0x64, 0xc1, 0x68, 0x4d, 0x65, 0xba, 0xa4, 0x06, 0x94, 0x26, 0xed, 0xb3, 0xd1,
	0x9d, 0x0b, 0x00, 0x21, 0x79, 0x37, 0x5d, 0x1b, 0x55, 0xb2, 0xc6, 0x4f, 0xf8, 0xd1, 0x43, 0xa3,
	0xc7, 0xda, 0xe6, 0x03, 0x28, 0xe3, 0x5d, 0x20, 0x37, 0xa7, 0x19, 0xf7, 0x9d, 0xa9, 0x33, 0x1b,
	0xc6, 0xe3, 0xfd, 0x2e, 0x18, 0xad, 0x40, 0x65, 0x37, 0x61, 0xd5, 0x0d, 0x13, 0x0b, 0x7a, 0x73,
	0xe4, 0x56, 0x91, 0xfc, 0xde, 0xd4, 0x99, 0xf5, 0xe3, 0xb0, 0x23, 0x55, 0xdd, 0xf0, 0xeb, 0x33,
	0x18, 0xff, 0xfb, 0x3e, 0x7c, 0x8e, 0xbf, 0xdd, 0xba, 0x89, 0xe5, 0x7b, 0x04, 0x0d, 0xec, 0x56,
	0x06, 0xd2, 0xff, 0x63, 0x17, 0x9c, 0xec, 0x77, 0xc1, 0xb8, 0xd1, 0x36, 0x48, 0x98, 0xb4, 0x24,
	0xef, 0x1e, 0x21, 0x43, 0x95, 0xe0, 0xa6, 0xf2, 0xe6, 0xbb, 0x53, 0x67, 0x36, 0xba, 0x3a, 0xc3,
	0x75, 0x20, 0x7c, 0x08, 0x84, 0x9f, 0xa8, 0x2c, 0x79, 0x7c, 0xba, 0xdf, 0x05, 0xff, 0xeb, 0x51,
	0x9d, 0x26, 0x4c, 0x8e, 0x06, 0x78, 0x31, 0x1a, 0xe4, 0xb0, 0xb4, 0x41, 0xfd, 0xbe, 0xf5, 0x7e,
	0xd9, 0xed, 0x3f, 0x20, 0xbf, 0xfa, 0x6f, 0x75, 0xf1, 0xdd, 0xf3, 0xad, 0x00, 0xcc, 0xa9, 0x7e,
	0x87, 0x42, 0x63, 0x09, 0x8c, 0x4a, 0x52, 0xdd, 0x54, 0x51, 0x66, 0x74, 0xfd, 0x93, 0x15, 0x2f,
	0x20, 0xca, 0x60, 0xc9, 0xa5, 0x26, 0x0d, 0x91, 0xd8, 0x92, 0xb4, 0xc7, 0x5f, 0xfc, 0xb5, 0xcc,
	0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa0, 0xf3, 0x97, 0x24, 0x02, 0x00, 0x00,
}
