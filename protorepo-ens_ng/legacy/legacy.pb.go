// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.easyops.local/contracts/protorepo-ens_ng/legacy/legacy.proto

package legacy

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Route struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f73d4f7e957093, []int{0}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Route) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Route) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Name struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f73d4f7e957093, []int{1}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Routes struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Routes) Reset()         { *m = Routes{} }
func (m *Routes) String() string { return proto.CompactTextString(m) }
func (*Routes) ProtoMessage()    {}
func (*Routes) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f73d4f7e957093, []int{2}
}
func (m *Routes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Routes.Unmarshal(m, b)
}
func (m *Routes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Routes.Marshal(b, m, deterministic)
}
func (m *Routes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Routes.Merge(m, src)
}
func (m *Routes) XXX_Size() int {
	return xxx_messageInfo_Routes.Size(m)
}
func (m *Routes) XXX_DiscardUnknown() {
	xxx_messageInfo_Routes.DiscardUnknown(m)
}

var xxx_messageInfo_Routes proto.InternalMessageInfo

func (m *Routes) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Status struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f73d4f7e957093, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Route)(nil), "legacy.Route")
	proto.RegisterType((*Name)(nil), "legacy.Name")
	proto.RegisterType((*Routes)(nil), "legacy.Routes")
	proto.RegisterType((*Status)(nil), "legacy.Status")
}

func init() {
	proto.RegisterFile("github.com/easyopsapis/easyops-api-go/protorepo-ens_ng/legacy/legacy.proto", fileDescriptor_60f73d4f7e957093)
}

var fileDescriptor_60f73d4f7e957093 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x25, 0xee, 0x1a, 0x74, 0xd6, 0x2a, 0xe6, 0x54, 0xf6, 0x54, 0x0a, 0x42, 0x2f, 0xdb, 0x4a,
	0xfd, 0x00, 0x11, 0x05, 0x6f, 0x1e, 0xb2, 0x1f, 0x20, 0x31, 0x0c, 0x75, 0x21, 0x76, 0x4a, 0x32,
	0x2b, 0xf4, 0x77, 0xfc, 0x52, 0x49, 0xda, 0x22, 0x7a, 0xf3, 0x94, 0x37, 0xef, 0xbd, 0xbc, 0xbc,
	0x0c, 0xdc, 0x77, 0x54, 0xa3, 0x09, 0x23, 0x0d, 0xa1, 0x76, 0x64, 0x8d, 0x6b, 0x2c, 0xf5, 0xec,
	0x8d, 0xe5, 0xd0, 0x0c, 0x9e, 0x98, 0x3c, 0x0e, 0xb4, 0xc3, 0x3e, 0xbc, 0xf6, 0x5d, 0xe3, 0xb0,
	0x33, 0x76, 0x9c, 0x8f, 0x3a, 0xc9, 0x4a, 0x4e, 0x53, 0xf9, 0x08, 0xa7, 0x9a, 0x8e, 0x8c, 0x4a,
	0xc1, 0xba, 0x37, 0x1f, 0x98, 0x8b, 0x42, 0x54, 0xe7, 0x3a, 0xe1, 0xc8, 0xbd, 0x53, 0xe0, 0xfc,
	0x64, 0xe2, 0x22, 0x8e, 0xdc, 0x40, 0x9e, 0xf3, 0x55, 0x21, 0xaa, 0x4c, 0x27, 0x5c, 0x6e, 0x61,
	0xfd, 0x32, 0xfb, 0xff, 0x66, 0x94, 0x0d, 0xc8, 0xf4, 0x40, 0x50, 0x37, 0x20, 0x7d, 0x42, 0xb9,
	0x28, 0x56, 0xd5, 0xa6, 0xcd, 0xea, 0xb9, 0x51, 0xd2, 0xf5, 0x2c, 0x96, 0x67, 0x20, 0xf7, 0x6c,
	0xf8, 0x18, 0xda, 0x2f, 0x01, 0x9b, 0x98, 0xbb, 0x47, 0xff, 0x79, 0xb0, 0xa8, 0x6e, 0xe1, 0x4a,
	0x63, 0x77, 0x08, 0x8c, 0x7e, 0xa1, 0x7e, 0x67, 0x6c, 0x2f, 0x97, 0x71, 0x4a, 0x50, 0x2d, 0x5c,
	0x3f, 0xa1, 0xff, 0xdf, 0x9d, 0x1d, 0x64, 0xcf, 0xc8, 0x0f, 0xce, 0x2d, 0xfe, 0x8b, 0xc5, 0x10,
	0xbb, 0xfc, 0xd8, 0xa7, 0x5f, 0xbd, 0xc9, 0xb4, 0xcf, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0xd0, 0xaa, 0xa1, 0x92, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NameServiceClient is the client API for NameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameServiceClient interface {
	RegisterService(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Status, error)
	DeregisterService(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Status, error)
	GetAllService(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Routes, error)
}

type nameServiceClient struct {
	cc *grpc.ClientConn
}

func NewNameServiceClient(cc *grpc.ClientConn) NameServiceClient {
	return &nameServiceClient{cc}
}

func (c *nameServiceClient) RegisterService(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/legacy.NameService/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameServiceClient) DeregisterService(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/legacy.NameService/DeregisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameServiceClient) GetAllService(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Routes, error) {
	out := new(Routes)
	err := c.cc.Invoke(ctx, "/legacy.NameService/GetAllService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameServiceServer is the server API for NameService service.
type NameServiceServer interface {
	RegisterService(context.Context, *Route) (*Status, error)
	DeregisterService(context.Context, *Route) (*Status, error)
	GetAllService(context.Context, *Name) (*Routes, error)
}

// UnimplementedNameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNameServiceServer struct {
}

func (*UnimplementedNameServiceServer) RegisterService(ctx context.Context, req *Route) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (*UnimplementedNameServiceServer) DeregisterService(ctx context.Context, req *Route) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterService not implemented")
}
func (*UnimplementedNameServiceServer) GetAllService(ctx context.Context, req *Name) (*Routes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllService not implemented")
}

func RegisterNameServiceServer(s *grpc.Server, srv NameServiceServer) {
	s.RegisterService(&_NameService_serviceDesc, srv)
}

func _NameService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/legacy.NameService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).RegisterService(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameService_DeregisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).DeregisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/legacy.NameService/DeregisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).DeregisterService(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameService_GetAllService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).GetAllService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/legacy.NameService/GetAllService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).GetAllService(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _NameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "legacy.NameService",
	HandlerType: (*NameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _NameService_RegisterService_Handler,
		},
		{
			MethodName: "DeregisterService",
			Handler:    _NameService_DeregisterService_Handler,
		},
		{
			MethodName: "GetAllService",
			Handler:    _NameService_GetAllService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/easyopsapis/easyops-api-go/protorepo-ens_ng/legacy/legacy.proto",
}
