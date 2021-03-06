// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auto_discovery.proto

package instance

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
//AutoDiscovery请求
type AutoDiscoveryRequest struct {
	//
	//资源模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//请求体
	Body                 []*AutoDiscoveryRequest_Body `protobuf:"bytes,2,rep,name=body,proto3" json:"body" form:"body"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AutoDiscoveryRequest) Reset()         { *m = AutoDiscoveryRequest{} }
func (m *AutoDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*AutoDiscoveryRequest) ProtoMessage()    {}
func (*AutoDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff7f7f760d7e1a84, []int{0}
}
func (m *AutoDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoDiscoveryRequest.Unmarshal(m, b)
}
func (m *AutoDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoDiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *AutoDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoDiscoveryRequest.Merge(m, src)
}
func (m *AutoDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_AutoDiscoveryRequest.Size(m)
}
func (m *AutoDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AutoDiscoveryRequest proto.InternalMessageInfo

func (m *AutoDiscoveryRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AutoDiscoveryRequest) GetBody() []*AutoDiscoveryRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type AutoDiscoveryRequest_Body struct {
	//
	//筛选器
	Filter *types.Struct `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter" form:"filter"`
	//
	//更新数据
	Update *types.Struct `protobuf:"bytes,2,opt,name=update,proto3" json:"update" form:"update"`
	//
	//存在即更新
	Upsert               bool     `protobuf:"varint,3,opt,name=upsert,proto3" json:"upsert" form:"upsert"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoDiscoveryRequest_Body) Reset()         { *m = AutoDiscoveryRequest_Body{} }
func (m *AutoDiscoveryRequest_Body) String() string { return proto.CompactTextString(m) }
func (*AutoDiscoveryRequest_Body) ProtoMessage()    {}
func (*AutoDiscoveryRequest_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff7f7f760d7e1a84, []int{0, 0}
}
func (m *AutoDiscoveryRequest_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoDiscoveryRequest_Body.Unmarshal(m, b)
}
func (m *AutoDiscoveryRequest_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoDiscoveryRequest_Body.Marshal(b, m, deterministic)
}
func (m *AutoDiscoveryRequest_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoDiscoveryRequest_Body.Merge(m, src)
}
func (m *AutoDiscoveryRequest_Body) XXX_Size() int {
	return xxx_messageInfo_AutoDiscoveryRequest_Body.Size(m)
}
func (m *AutoDiscoveryRequest_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoDiscoveryRequest_Body.DiscardUnknown(m)
}

var xxx_messageInfo_AutoDiscoveryRequest_Body proto.InternalMessageInfo

func (m *AutoDiscoveryRequest_Body) GetFilter() *types.Struct {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *AutoDiscoveryRequest_Body) GetUpdate() *types.Struct {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *AutoDiscoveryRequest_Body) GetUpsert() bool {
	if m != nil {
		return m.Upsert
	}
	return false
}

//
//AutoDiscovery返回
type AutoDiscoveryResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//data
	Data                 []*AutoDiscoveryResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AutoDiscoveryResponse) Reset()         { *m = AutoDiscoveryResponse{} }
func (m *AutoDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*AutoDiscoveryResponse) ProtoMessage()    {}
func (*AutoDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff7f7f760d7e1a84, []int{1}
}
func (m *AutoDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoDiscoveryResponse.Unmarshal(m, b)
}
func (m *AutoDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *AutoDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoDiscoveryResponse.Merge(m, src)
}
func (m *AutoDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_AutoDiscoveryResponse.Size(m)
}
func (m *AutoDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AutoDiscoveryResponse proto.InternalMessageInfo

func (m *AutoDiscoveryResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AutoDiscoveryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AutoDiscoveryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AutoDiscoveryResponse) GetData() []*AutoDiscoveryResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type AutoDiscoveryResponse_Data struct {
	//
	//实例Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//状态码
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code" form:"code"`
	//
	//是否匹配到数据
	Matched bool `protobuf:"varint,3,opt,name=matched,proto3" json:"matched" form:"matched"`
	//
	//是否有效
	Effected             bool     `protobuf:"varint,4,opt,name=effected,proto3" json:"effected" form:"effected"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoDiscoveryResponse_Data) Reset()         { *m = AutoDiscoveryResponse_Data{} }
func (m *AutoDiscoveryResponse_Data) String() string { return proto.CompactTextString(m) }
func (*AutoDiscoveryResponse_Data) ProtoMessage()    {}
func (*AutoDiscoveryResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff7f7f760d7e1a84, []int{1, 0}
}
func (m *AutoDiscoveryResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoDiscoveryResponse_Data.Unmarshal(m, b)
}
func (m *AutoDiscoveryResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoDiscoveryResponse_Data.Marshal(b, m, deterministic)
}
func (m *AutoDiscoveryResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoDiscoveryResponse_Data.Merge(m, src)
}
func (m *AutoDiscoveryResponse_Data) XXX_Size() int {
	return xxx_messageInfo_AutoDiscoveryResponse_Data.Size(m)
}
func (m *AutoDiscoveryResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoDiscoveryResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_AutoDiscoveryResponse_Data proto.InternalMessageInfo

func (m *AutoDiscoveryResponse_Data) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *AutoDiscoveryResponse_Data) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AutoDiscoveryResponse_Data) GetMatched() bool {
	if m != nil {
		return m.Matched
	}
	return false
}

func (m *AutoDiscoveryResponse_Data) GetEffected() bool {
	if m != nil {
		return m.Effected
	}
	return false
}

//
//AutoDiscoveryApi返回
type AutoDiscoveryResponseWrapper struct {
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
	Data                 *AutoDiscoveryResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AutoDiscoveryResponseWrapper) Reset()         { *m = AutoDiscoveryResponseWrapper{} }
func (m *AutoDiscoveryResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*AutoDiscoveryResponseWrapper) ProtoMessage()    {}
func (*AutoDiscoveryResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff7f7f760d7e1a84, []int{2}
}
func (m *AutoDiscoveryResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoDiscoveryResponseWrapper.Unmarshal(m, b)
}
func (m *AutoDiscoveryResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoDiscoveryResponseWrapper.Marshal(b, m, deterministic)
}
func (m *AutoDiscoveryResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoDiscoveryResponseWrapper.Merge(m, src)
}
func (m *AutoDiscoveryResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_AutoDiscoveryResponseWrapper.Size(m)
}
func (m *AutoDiscoveryResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoDiscoveryResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_AutoDiscoveryResponseWrapper proto.InternalMessageInfo

func (m *AutoDiscoveryResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AutoDiscoveryResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *AutoDiscoveryResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *AutoDiscoveryResponseWrapper) GetData() *AutoDiscoveryResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoDiscoveryRequest)(nil), "instance.AutoDiscoveryRequest")
	proto.RegisterType((*AutoDiscoveryRequest_Body)(nil), "instance.AutoDiscoveryRequest.Body")
	proto.RegisterType((*AutoDiscoveryResponse)(nil), "instance.AutoDiscoveryResponse")
	proto.RegisterType((*AutoDiscoveryResponse_Data)(nil), "instance.AutoDiscoveryResponse.Data")
	proto.RegisterType((*AutoDiscoveryResponseWrapper)(nil), "instance.AutoDiscoveryResponseWrapper")
}

func init() { proto.RegisterFile("auto_discovery.proto", fileDescriptor_ff7f7f760d7e1a84) }

var fileDescriptor_ff7f7f760d7e1a84 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd3, 0x3e,
	0x1c, 0x56, 0xda, 0x6c, 0xff, 0xce, 0xfd, 0xb3, 0x17, 0x33, 0x20, 0xaa, 0x86, 0x32, 0x79, 0x13,
	0xda, 0xa4, 0x25, 0xdd, 0x8b, 0x34, 0x01, 0xb7, 0x45, 0x45, 0xa2, 0x57, 0x73, 0x40, 0x62, 0xda,
	0x26, 0x37, 0x71, 0xb3, 0x40, 0x5b, 0x07, 0xc7, 0xd9, 0x28, 0xd3, 0x3e, 0x16, 0x07, 0x3e, 0x08,
	0xd7, 0x20, 0x71, 0xe5, 0x82, 0xf2, 0x09, 0x90, 0xed, 0x24, 0x0d, 0x68, 0x1b, 0x3b, 0xd5, 0xf1,
	0xf3, 0xd2, 0xdf, 0xe3, 0xe7, 0x07, 0x56, 0x49, 0x2a, 0xd8, 0x59, 0x10, 0x25, 0x3e, 0xbb, 0xa0,
	0x7c, 0xea, 0xc6, 0x9c, 0x09, 0x06, 0x5b, 0xd1, 0x24, 0x11, 0x64, 0xe2, 0xd3, 0x8e, 0x13, 0x46,
	0xe2, 0x3c, 0x1d, 0xb8, 0x3e, 0x1b, 0x77, 0x43, 0x16, 0xb2, 0xae, 0x22, 0x0c, 0xd2, 0xa1, 0xfa,
	0x52, 0x1f, 0xea, 0xa4, 0x85, 0x9d, 0xc3, 0x1a, 0x7d, 0x7c, 0x19, 0x89, 0x0f, 0xec, 0xb2, 0x1b,
	0x32, 0x47, 0x81, 0xce, 0x05, 0x19, 0x45, 0x01, 0x11, 0x8c, 0x27, 0xdd, 0xea, 0x58, 0xe8, 0xd6,
	0x42, 0xc6, 0xc2, 0x11, 0x9d, 0xb9, 0x27, 0x82, 0xa7, 0xbe, 0xd0, 0x28, 0xfa, 0xd5, 0x00, 0xab,
	0x47, 0xa9, 0x60, 0xbd, 0x72, 0x4c, 0x4c, 0x3f, 0xa6, 0x34, 0x11, 0x10, 0x83, 0x16, 0x1b, 0xbc,
	0xa7, 0xbe, 0xe8, 0x07, 0x96, 0xb1, 0x6e, 0x6c, 0x2d, 0x78, 0x87, 0x79, 0x66, 0x2f, 0x0d, 0x19,
	0x1f, 0xbf, 0x44, 0x25, 0x82, 0x7e, 0x7c, 0xb7, 0x6d, 0xf0, 0xf4, 0xf4, 0x98, 0x38, 0x9f, 0x8f,
	0x9c, 0x77, 0x67, 0x27, 0xc7, 0xbb, 0xce, 0x8b, 0xf2, 0x7c, 0xb5, 0xbb, 0x73, 0xb0, 0x77, 0xbd,
	0x89, 0x2b, 0x1f, 0xf8, 0x1a, 0x98, 0x03, 0x16, 0x4c, 0xad, 0xc6, 0x7a, 0x73, 0xab, 0xbd, 0xbf,
	0xe1, 0x96, 0x4f, 0xe1, 0xde, 0x34, 0x81, 0xeb, 0xb1, 0x60, 0xea, 0x2d, 0xe5, 0x99, 0xdd, 0xd6,
	0x7f, 0x2a, 0xa5, 0x08, 0x2b, 0x87, 0xce, 0x57, 0x03, 0x98, 0x12, 0x87, 0x1e, 0x98, 0x1f, 0x46,
	0x23, 0x41, 0xb9, 0x1a, 0xb2, 0xbd, 0xff, 0xc4, 0xd5, 0x71, 0xdd, 0x32, 0xae, 0xfb, 0x46, 0xc5,
	0xf5, 0x56, 0xf2, 0xcc, 0x7e, 0xa0, 0x8d, 0xb4, 0x00, 0xe1, 0x42, 0x29, 0x3d, 0xd2, 0x38, 0x20,
	0x82, 0x5a, 0x8d, 0x7b, 0x7b, 0x68, 0x01, 0xc2, 0x85, 0x12, 0x6e, 0x4b, 0x8f, 0x84, 0x72, 0x61,
	0x35, 0xd7, 0x8d, 0xad, 0xd6, 0x9f, 0x54, 0x79, 0xaf, 0xa8, 0xea, 0xf0, 0xa5, 0x09, 0x1e, 0xfd,
	0x15, 0x38, 0x89, 0xd9, 0x24, 0xa1, 0x70, 0x03, 0x98, 0x3e, 0x0b, 0xa8, 0x8a, 0x32, 0x57, 0x8f,
	0x2e, 0x6f, 0x11, 0x56, 0x20, 0x7c, 0x06, 0xe6, 0x28, 0xe7, 0x8c, 0xab, 0x61, 0x17, 0xbc, 0xe5,
	0x3c, 0xb3, 0xff, 0xd7, 0x2c, 0x75, 0x8d, 0xb0, 0x86, 0xe1, 0x0e, 0xf8, 0x6f, 0x4c, 0x93, 0x84,
	0x84, 0x54, 0x8d, 0xb4, 0xe0, 0xc1, 0x3c, 0xb3, 0x17, 0x35, 0xb3, 0x00, 0x10, 0x2e, 0x29, 0xb0,
	0x0f, 0xcc, 0x80, 0x08, 0x62, 0x99, 0xaa, 0x9a, 0xcd, 0x5b, 0xab, 0xd1, 0x93, 0xba, 0x3d, 0x22,
	0x48, 0x7d, 0x40, 0xa9, 0x45, 0x58, 0x59, 0x74, 0xbe, 0x19, 0xc0, 0x94, 0x38, 0xec, 0x03, 0x50,
	0xda, 0x54, 0x4b, 0xb4, 0x9d, 0x67, 0xf6, 0x8a, 0xd6, 0xcc, 0x30, 0xb9, 0x46, 0xcb, 0x60, 0xf1,
	0xb4, 0xd8, 0x9e, 0x93, 0xab, 0xbd, 0x83, 0xeb, 0x4d, 0x5c, 0x13, 0x57, 0x2f, 0xd3, 0xb8, 0xeb,
	0x65, 0x64, 0x62, 0x22, 0xfc, 0x73, 0x1a, 0x14, 0x25, 0xd4, 0x13, 0x6b, 0x40, 0x26, 0xd6, 0x27,
	0xd8, 0x05, 0x2d, 0x3a, 0x1c, 0x52, 0x5f, 0xd0, 0xc0, 0x32, 0x15, 0xfd, 0xe1, 0x6c, 0xc1, 0x4b,
	0x04, 0xe1, 0x8a, 0x84, 0x7e, 0x1a, 0x60, 0xed, 0xc6, 0xd7, 0x78, 0xcb, 0x49, 0x1c, 0x53, 0x7e,
	0xbf, 0xfa, 0x9e, 0x83, 0xb6, 0xfc, 0x7d, 0xf5, 0x29, 0x1e, 0x91, 0x68, 0x52, 0x94, 0xf8, 0x38,
	0xcf, 0x6c, 0x38, 0xe3, 0x16, 0x20, 0xc2, 0x75, 0xea, 0xac, 0xf8, 0xe6, 0xdd, 0xc5, 0xf7, 0xaa,
	0x2a, 0xe5, 0x32, 0xdb, 0xff, 0xa8, 0xf2, 0x96, 0x16, 0x07, 0xf3, 0x6a, 0xf9, 0x0f, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x43, 0xf0, 0x28, 0xc6, 0x04, 0x00, 0x00,
}
