// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relation_count_aggregate.proto

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
//RelationCountAggregate请求
type RelationCountAggregateRequest struct {
	//
	//模型对象ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//e.g.: { name: { $like: '%q%' } }, { $or: [{ name: { $like: '%q%' }}] }
	Query *types.Struct `protobuf:"bytes,2,opt,name=query,proto3" json:"query" form:"query"`
	//
	//关系sideId,例如APP的owner
	RelationSideIds []string `protobuf:"bytes,3,rep,name=relation_side_ids,json=relationSideIds,proto3" json:"relation_side_ids" form:"relation_side_ids"`
	//
	//当为 true 时，只搜索与我相关实例
	OnlyMyInstance       bool     `protobuf:"varint,4,opt,name=only_my_instance,json=onlyMyInstance,proto3" json:"only_my_instance" form:"only_my_instance"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationCountAggregateRequest) Reset()         { *m = RelationCountAggregateRequest{} }
func (m *RelationCountAggregateRequest) String() string { return proto.CompactTextString(m) }
func (*RelationCountAggregateRequest) ProtoMessage()    {}
func (*RelationCountAggregateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b9bc328045285bd, []int{0}
}
func (m *RelationCountAggregateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationCountAggregateRequest.Unmarshal(m, b)
}
func (m *RelationCountAggregateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationCountAggregateRequest.Marshal(b, m, deterministic)
}
func (m *RelationCountAggregateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationCountAggregateRequest.Merge(m, src)
}
func (m *RelationCountAggregateRequest) XXX_Size() int {
	return xxx_messageInfo_RelationCountAggregateRequest.Size(m)
}
func (m *RelationCountAggregateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationCountAggregateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelationCountAggregateRequest proto.InternalMessageInfo

func (m *RelationCountAggregateRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *RelationCountAggregateRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *RelationCountAggregateRequest) GetRelationSideIds() []string {
	if m != nil {
		return m.RelationSideIds
	}
	return nil
}

func (m *RelationCountAggregateRequest) GetOnlyMyInstance() bool {
	if m != nil {
		return m.OnlyMyInstance
	}
	return false
}

//
//RelationCountAggregate返回
type RelationCountAggregateResponse struct {
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
	//关系数量统计结果
	Data                 []*RelationCountAggregateResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RelationCountAggregateResponse) Reset()         { *m = RelationCountAggregateResponse{} }
func (m *RelationCountAggregateResponse) String() string { return proto.CompactTextString(m) }
func (*RelationCountAggregateResponse) ProtoMessage()    {}
func (*RelationCountAggregateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b9bc328045285bd, []int{1}
}
func (m *RelationCountAggregateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationCountAggregateResponse.Unmarshal(m, b)
}
func (m *RelationCountAggregateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationCountAggregateResponse.Marshal(b, m, deterministic)
}
func (m *RelationCountAggregateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationCountAggregateResponse.Merge(m, src)
}
func (m *RelationCountAggregateResponse) XXX_Size() int {
	return xxx_messageInfo_RelationCountAggregateResponse.Size(m)
}
func (m *RelationCountAggregateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationCountAggregateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RelationCountAggregateResponse proto.InternalMessageInfo

func (m *RelationCountAggregateResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RelationCountAggregateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RelationCountAggregateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RelationCountAggregateResponse) GetData() []*RelationCountAggregateResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RelationCountAggregateResponse_Data struct {
	//
	//关系sideId
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" form:"key"`
	//
	//关系个数
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationCountAggregateResponse_Data) Reset()         { *m = RelationCountAggregateResponse_Data{} }
func (m *RelationCountAggregateResponse_Data) String() string { return proto.CompactTextString(m) }
func (*RelationCountAggregateResponse_Data) ProtoMessage()    {}
func (*RelationCountAggregateResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b9bc328045285bd, []int{1, 0}
}
func (m *RelationCountAggregateResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationCountAggregateResponse_Data.Unmarshal(m, b)
}
func (m *RelationCountAggregateResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationCountAggregateResponse_Data.Marshal(b, m, deterministic)
}
func (m *RelationCountAggregateResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationCountAggregateResponse_Data.Merge(m, src)
}
func (m *RelationCountAggregateResponse_Data) XXX_Size() int {
	return xxx_messageInfo_RelationCountAggregateResponse_Data.Size(m)
}
func (m *RelationCountAggregateResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationCountAggregateResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RelationCountAggregateResponse_Data proto.InternalMessageInfo

func (m *RelationCountAggregateResponse_Data) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RelationCountAggregateResponse_Data) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

//
//RelationCountAggregateApi返回
type RelationCountAggregateResponseWrapper struct {
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
	Data                 *RelationCountAggregateResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RelationCountAggregateResponseWrapper) Reset()         { *m = RelationCountAggregateResponseWrapper{} }
func (m *RelationCountAggregateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*RelationCountAggregateResponseWrapper) ProtoMessage()    {}
func (*RelationCountAggregateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b9bc328045285bd, []int{2}
}
func (m *RelationCountAggregateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationCountAggregateResponseWrapper.Unmarshal(m, b)
}
func (m *RelationCountAggregateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationCountAggregateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *RelationCountAggregateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationCountAggregateResponseWrapper.Merge(m, src)
}
func (m *RelationCountAggregateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_RelationCountAggregateResponseWrapper.Size(m)
}
func (m *RelationCountAggregateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationCountAggregateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RelationCountAggregateResponseWrapper proto.InternalMessageInfo

func (m *RelationCountAggregateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RelationCountAggregateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *RelationCountAggregateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RelationCountAggregateResponseWrapper) GetData() *RelationCountAggregateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RelationCountAggregateRequest)(nil), "instance.RelationCountAggregateRequest")
	proto.RegisterType((*RelationCountAggregateResponse)(nil), "instance.RelationCountAggregateResponse")
	proto.RegisterType((*RelationCountAggregateResponse_Data)(nil), "instance.RelationCountAggregateResponse.Data")
	proto.RegisterType((*RelationCountAggregateResponseWrapper)(nil), "instance.RelationCountAggregateResponseWrapper")
}

func init() { proto.RegisterFile("relation_count_aggregate.proto", fileDescriptor_3b9bc328045285bd) }

var fileDescriptor_3b9bc328045285bd = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0x9b, 0x16, 0x5a, 0x17, 0xad, 0xc3, 0x07, 0x16, 0x15, 0xb6, 0x54, 0xe6, 0x87, 0x72,
	0x58, 0xd2, 0xfd, 0x90, 0x26, 0xe0, 0x32, 0xad, 0x63, 0x87, 0x1e, 0x76, 0xf1, 0x0e, 0x48, 0x4c,
	0x10, 0xb9, 0x89, 0x17, 0x42, 0xd3, 0xb8, 0xb3, 0x9d, 0x8d, 0x80, 0xf8, 0x57, 0x83, 0x84, 0x38,
	0x70, 0xce, 0x1d, 0x09, 0xc5, 0x49, 0xda, 0x8a, 0x69, 0x55, 0x4f, 0xb5, 0xbf, 0xf7, 0xde, 0xd7,
	0xf7, 0xbd, 0x2f, 0x06, 0x3b, 0x9c, 0x86, 0x44, 0x06, 0x2c, 0x72, 0x5c, 0x16, 0x47, 0xd2, 0x21,
	0xbe, 0xcf, 0xa9, 0x4f, 0x24, 0xb5, 0x67, 0x9c, 0x49, 0x06, 0x5b, 0x41, 0x24, 0x24, 0x89, 0x5c,
	0xda, 0xb3, 0xfc, 0x40, 0x7e, 0x8e, 0xc7, 0xb6, 0xcb, 0xa6, 0x03, 0x9f, 0xf9, 0x6c, 0xa0, 0x08,
	0xe3, 0xf8, 0x4a, 0xdd, 0xd4, 0x45, 0x9d, 0x0a, 0x61, 0xef, 0x68, 0x89, 0x3e, 0xbd, 0x0d, 0xe4,
	0x84, 0xdd, 0x0e, 0x7c, 0x66, 0x29, 0xd0, 0xba, 0x21, 0x61, 0xe0, 0x11, 0xc9, 0xb8, 0x18, 0xcc,
	0x8f, 0xa5, 0xee, 0x99, 0xcf, 0x98, 0x1f, 0xd2, 0x45, 0x77, 0x21, 0x79, 0xec, 0xca, 0x02, 0x45,
	0x7f, 0xea, 0x60, 0x1b, 0x97, 0x8e, 0x4f, 0x73, 0xc3, 0x27, 0x95, 0x5f, 0x4c, 0xaf, 0x63, 0x2a,
	0x24, 0xc4, 0xa0, 0xc5, 0xc6, 0x5f, 0xa8, 0x2b, 0x47, 0x9e, 0x5e, 0xeb, 0xd7, 0xcc, 0xf6, 0xf0,
	0x28, 0x4b, 0x8d, 0xee, 0x15, 0xe3, 0xd3, 0xb7, 0xa8, 0x42, 0xd0, 0xaf, 0x9f, 0x86, 0x01, 0xb6,
	0x3f, 0x5d, 0x12, 0xeb, 0xdb, 0x89, 0xf5, 0xc1, 0xf9, 0x78, 0xb9, 0x67, 0xbd, 0xa9, 0xce, 0xdf,
	0xf7, 0x76, 0x0f, 0xf7, 0x7f, 0xbc, 0xc0, 0xf3, 0x3e, 0xf0, 0x18, 0x34, 0xaf, 0x63, 0xca, 0x13,
	0xbd, 0xde, 0xaf, 0x99, 0x9d, 0x83, 0x2d, 0xbb, 0xf0, 0x68, 0x57, 0x1e, 0xed, 0x0b, 0xe5, 0x71,
	0xb8, 0x99, 0xa5, 0xc6, 0xa3, 0xe2, 0x9f, 0x14, 0x1f, 0xe1, 0x42, 0x07, 0x27, 0xe0, 0xf1, 0x3c,
	0x67, 0x11, 0x78, 0xd4, 0x09, 0x3c, 0xa1, 0x6b, 0x7d, 0xcd, 0x6c, 0x0f, 0x8f, 0xb3, 0xd4, 0xd0,
	0x0b, 0xcd, 0x1d, 0xca, 0x5a, 0x36, 0xbb, 0x95, 0xec, 0x22, 0xf0, 0xe8, 0xc8, 0x13, 0xf0, 0x0c,
	0x6c, 0xb2, 0x28, 0x4c, 0x9c, 0x69, 0xe2, 0x54, 0xcb, 0xd3, 0x1b, 0xfd, 0x9a, 0xd9, 0x1a, 0x3e,
	0xcd, 0x52, 0x63, 0xab, 0x4c, 0xe2, 0x3f, 0x06, 0xc2, 0x1b, 0x79, 0xe9, 0x3c, 0x19, 0x55, 0x85,
	0xdf, 0x75, 0xb0, 0x73, 0x5f, 0xd4, 0x62, 0xc6, 0x22, 0x41, 0xe1, 0x73, 0xd0, 0x70, 0x99, 0x47,
	0x55, 0xce, 0xcd, 0x61, 0x37, 0x4b, 0x8d, 0x4e, 0xd1, 0x3d, 0xaf, 0x22, 0xac, 0x40, 0xf8, 0x0a,
	0x34, 0x29, 0xe7, 0x8c, 0xab, 0xf0, 0xda, 0xcb, 0x19, 0xa9, 0x32, 0xc2, 0x05, 0x0c, 0x77, 0xc1,
	0xc3, 0x29, 0x15, 0x82, 0xf8, 0x54, 0xd7, 0x14, 0x13, 0x66, 0xa9, 0xb1, 0x51, 0x30, 0x4b, 0x00,
	0xe1, 0x8a, 0x02, 0x31, 0x68, 0x78, 0x44, 0x12, 0xbd, 0xd1, 0xd7, 0xcc, 0xce, 0x81, 0x65, 0x57,
	0x73, 0xd8, 0xab, 0x2d, 0xdb, 0xef, 0x88, 0x24, 0xcb, 0x4e, 0xf3, 0x26, 0x08, 0xab, 0x5e, 0x3d,
	0x01, 0x1a, 0x39, 0x0c, 0x4f, 0x81, 0x36, 0xa1, 0x49, 0xf9, 0xf5, 0xec, 0x67, 0xa9, 0x01, 0x0a,
	0xee, 0x84, 0x26, 0x6b, 0x6d, 0x24, 0x57, 0xe7, 0x63, 0xdf, 0x90, 0x30, 0xa6, 0x6a, 0xec, 0xe6,
	0xf2, 0xd8, 0xaa, 0x8c, 0x70, 0x01, 0xa3, 0xbf, 0x35, 0xf0, 0x72, 0xb5, 0xe7, 0xf7, 0x9c, 0xcc,
	0x66, 0x94, 0xaf, 0x97, 0xf6, 0x6b, 0xd0, 0xc9, 0x7f, 0xcf, 0xbe, 0xce, 0x42, 0x12, 0x44, 0x65,
	0xe6, 0x4f, 0xb2, 0xd4, 0x80, 0x0b, 0x6e, 0x09, 0x22, 0xbc, 0x4c, 0x5d, 0xec, 0x49, 0x5b, 0xbd,
	0xa7, 0xf3, 0x79, 0xf2, 0xf9, 0x5b, 0x30, 0xd7, 0x4d, 0xfe, 0x9e, 0xd0, 0xc7, 0x0f, 0xd4, 0x23,
	0x3a, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x27, 0xec, 0xd6, 0xe2, 0x89, 0x04, 0x00, 0x00,
}
