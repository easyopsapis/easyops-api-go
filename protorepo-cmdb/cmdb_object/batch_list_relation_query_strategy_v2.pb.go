// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_list_relation_query_strategy_v2.proto

package cmdb_object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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
//BatchListRelationQueryStrategyV2Request请求
type BatchListRelationQueryStrategyV2RequestRequest struct {
	//
	//多个模型ID，以逗号分隔
	ObjectIds string `protobuf:"bytes,1,opt,name=object_ids,json=objectIds,proto3" json:"object_ids" form:"object_ids"`
	//
	//策略类型
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchListRelationQueryStrategyV2RequestRequest) Reset() {
	*m = BatchListRelationQueryStrategyV2RequestRequest{}
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) String() string {
	return proto.CompactTextString(m)
}
func (*BatchListRelationQueryStrategyV2RequestRequest) ProtoMessage() {}
func (*BatchListRelationQueryStrategyV2RequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07dc2fcefa871665, []int{0}
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest.Unmarshal(m, b)
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest.Marshal(b, m, deterministic)
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest.Merge(m, src)
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) XXX_Size() int {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest.Size(m)
}
func (m *BatchListRelationQueryStrategyV2RequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchListRelationQueryStrategyV2RequestRequest proto.InternalMessageInfo

func (m *BatchListRelationQueryStrategyV2RequestRequest) GetObjectIds() string {
	if m != nil {
		return m.ObjectIds
	}
	return ""
}

func (m *BatchListRelationQueryStrategyV2RequestRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

//
//BatchListRelationQueryStrategyV2Request返回
type BatchListRelationQueryStrategyV2RequestResponse struct {
	//
	//各模型策略数据
	StrategyData         []*BatchListRelationQueryStrategyV2RequestResponse_StrategyData `protobuf:"bytes,1,rep,name=strategy_data,json=strategyData,proto3" json:"strategy_data" form:"strategy_data"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *BatchListRelationQueryStrategyV2RequestResponse) Reset() {
	*m = BatchListRelationQueryStrategyV2RequestResponse{}
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) String() string {
	return proto.CompactTextString(m)
}
func (*BatchListRelationQueryStrategyV2RequestResponse) ProtoMessage() {}
func (*BatchListRelationQueryStrategyV2RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07dc2fcefa871665, []int{1}
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse.Unmarshal(m, b)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse.Marshal(b, m, deterministic)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse.Merge(m, src)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) XXX_Size() int {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse.Size(m)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse proto.InternalMessageInfo

func (m *BatchListRelationQueryStrategyV2RequestResponse) GetStrategyData() []*BatchListRelationQueryStrategyV2RequestResponse_StrategyData {
	if m != nil {
		return m.StrategyData
	}
	return nil
}

type BatchListRelationQueryStrategyV2RequestResponse_StrategyData struct {
	//
	//模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//策略列表
	StrategyList         []*cmdb.RelationQueryStrategyV2 `protobuf:"bytes,2,rep,name=strategy_list,json=strategyList,proto3" json:"strategy_list" form:"strategy_list"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) Reset() {
	*m = BatchListRelationQueryStrategyV2RequestResponse_StrategyData{}
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) String() string {
	return proto.CompactTextString(m)
}
func (*BatchListRelationQueryStrategyV2RequestResponse_StrategyData) ProtoMessage() {}
func (*BatchListRelationQueryStrategyV2RequestResponse_StrategyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_07dc2fcefa871665, []int{1, 0}
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData.Unmarshal(m, b)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData.Marshal(b, m, deterministic)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData.Merge(m, src)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) XXX_Size() int {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData.Size(m)
}
func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData.DiscardUnknown(m)
}

var xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponse_StrategyData proto.InternalMessageInfo

func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *BatchListRelationQueryStrategyV2RequestResponse_StrategyData) GetStrategyList() []*cmdb.RelationQueryStrategyV2 {
	if m != nil {
		return m.StrategyList
	}
	return nil
}

//
//BatchListRelationQueryStrategyV2RequestApi返回
type BatchListRelationQueryStrategyV2RequestResponseWrapper struct {
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
	Data                 *BatchListRelationQueryStrategyV2RequestResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) Reset() {
	*m = BatchListRelationQueryStrategyV2RequestResponseWrapper{}
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) String() string {
	return proto.CompactTextString(m)
}
func (*BatchListRelationQueryStrategyV2RequestResponseWrapper) ProtoMessage() {}
func (*BatchListRelationQueryStrategyV2RequestResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_07dc2fcefa871665, []int{2}
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper.Unmarshal(m, b)
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper.Marshal(b, m, deterministic)
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper.Merge(m, src)
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper.Size(m)
}
func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BatchListRelationQueryStrategyV2RequestResponseWrapper proto.InternalMessageInfo

func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BatchListRelationQueryStrategyV2RequestResponseWrapper) GetData() *BatchListRelationQueryStrategyV2RequestResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchListRelationQueryStrategyV2RequestRequest)(nil), "cmdb_object.BatchListRelationQueryStrategyV2RequestRequest")
	proto.RegisterType((*BatchListRelationQueryStrategyV2RequestResponse)(nil), "cmdb_object.BatchListRelationQueryStrategyV2RequestResponse")
	proto.RegisterType((*BatchListRelationQueryStrategyV2RequestResponse_StrategyData)(nil), "cmdb_object.BatchListRelationQueryStrategyV2RequestResponse.StrategyData")
	proto.RegisterType((*BatchListRelationQueryStrategyV2RequestResponseWrapper)(nil), "cmdb_object.BatchListRelationQueryStrategyV2RequestResponseWrapper")
}

func init() {
	proto.RegisterFile("batch_list_relation_query_strategy_v2.proto", fileDescriptor_07dc2fcefa871665)
}

var fileDescriptor_07dc2fcefa871665 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0xdd, 0xae, 0xd8, 0x49, 0xc5, 0x75, 0x58, 0x25, 0x14, 0x24, 0x65, 0x04, 0x29, 0xc8,
	0x4e, 0xb0, 0x8a, 0x88, 0xf8, 0x54, 0xf4, 0x61, 0xc1, 0x17, 0x47, 0xd0, 0x17, 0x21, 0x4c, 0x92,
	0xd9, 0x6c, 0x24, 0xed, 0xcd, 0xce, 0x4c, 0xc5, 0xfe, 0x81, 0xe0, 0x47, 0x08, 0x7e, 0x58, 0x3e,
	0x22, 0x5f, 0x20, 0x73, 0x27, 0xbb, 0x1b, 0x45, 0xd1, 0x65, 0x9f, 0xa6, 0xf7, 0x9e, 0x73, 0x67,
	0xee, 0x39, 0xa7, 0x21, 0x8f, 0x32, 0x69, 0xf3, 0xd3, 0xb4, 0xae, 0x8c, 0x4d, 0xb5, 0xaa, 0xa5,
	0xad, 0x60, 0x93, 0x9e, 0x6d, 0x95, 0xde, 0xa5, 0xc6, 0x6a, 0x69, 0x55, 0xb9, 0x4b, 0x3f, 0x2f,
	0x79, 0xa3, 0xc1, 0x02, 0x0d, 0xf3, 0x75, 0x91, 0xa5, 0x90, 0x7d, 0x52, 0xb9, 0x9d, 0x1d, 0x95,
	0x95, 0x3d, 0xdd, 0x66, 0x3c, 0x87, 0x75, 0x52, 0x42, 0x09, 0x09, 0x72, 0xb2, 0xed, 0x09, 0x56,
	0x58, 0xe0, 0x2f, 0x3f, 0x3b, 0x4b, 0x4b, 0xe0, 0x4a, 0x9a, 0x1d, 0x34, 0x86, 0xd7, 0x90, 0xcb,
	0x3a, 0xc9, 0x61, 0x63, 0xb5, 0xcc, 0xad, 0xf1, 0x93, 0x5a, 0x35, 0x70, 0xb4, 0x86, 0x42, 0xd5,
	0x26, 0xe9, 0x89, 0x09, 0x96, 0x89, 0x7b, 0x35, 0xf9, 0xd7, 0x72, 0xec, 0x5b, 0x40, 0xf8, 0xca,
	0x89, 0x79, 0x53, 0x19, 0x2b, 0x7a, 0xf6, 0x5b, 0x47, 0x7e, 0xd7, 0x73, 0xdf, 0x2f, 0x85, 0x3a,
	0xdb, 0x2a, 0x87, 0xe2, 0x41, 0x9f, 0x12, 0xe2, 0xc5, 0xa4, 0x55, 0x61, 0xa2, 0x60, 0x1e, 0x2c,
	0x26, 0xab, 0xbb, 0x5d, 0x1b, 0xdf, 0x39, 0x01, 0xbd, 0x7e, 0xc1, 0x2e, 0x31, 0x26, 0x26, 0xbe,
	0x38, 0x2e, 0x0c, 0x7d, 0x40, 0xc6, 0x76, 0xd7, 0xa8, 0x68, 0x84, 0xfc, 0xdb, 0x5d, 0x1b, 0x87,
	0x9e, 0xef, 0xba, 0x4c, 0x20, 0xc8, 0xda, 0x11, 0x49, 0xfe, 0x7b, 0x1b, 0xd3, 0xc0, 0xc6, 0x28,
	0xfa, 0x35, 0x20, 0xb7, 0x2e, 0x74, 0x15, 0xd2, 0xca, 0x28, 0x98, 0xef, 0x2d, 0xc2, 0xe5, 0x31,
	0x1f, 0xf8, 0xce, 0xaf, 0x78, 0x2b, 0x3f, 0x47, 0x5e, 0x49, 0x2b, 0x57, 0x51, 0xd7, 0xc6, 0x87,
	0x7e, 0xdb, 0x5f, 0x5e, 0x62, 0x62, 0x6a, 0x06, 0xbc, 0xd9, 0xf7, 0x80, 0x4c, 0x87, 0x83, 0xf4,
	0x31, 0x99, 0x5c, 0xd8, 0xd1, 0x3b, 0x75, 0xd8, 0xb5, 0xf1, 0xc1, 0x6f, 0x4e, 0x31, 0x71, 0xf3,
	0xdc, 0x28, 0xfa, 0x71, 0xa0, 0xc6, 0xfd, 0xbf, 0xa2, 0x11, 0xaa, 0xb9, 0x8f, 0x6a, 0xf8, 0x5f,
	0xb6, 0xff, 0xe3, 0x86, 0x6e, 0x7a, 0xb0, 0xa1, 0x13, 0xcf, 0x7e, 0x8c, 0xc8, 0xb3, 0x2b, 0x5a,
	0xf1, 0x41, 0xcb, 0xa6, 0x51, 0xda, 0x05, 0x98, 0x43, 0xa1, 0x50, 0xc6, 0xfe, 0x30, 0x40, 0xd7,
	0x65, 0x02, 0x41, 0xfa, 0x9c, 0x84, 0xee, 0x7c, 0xfd, 0xa5, 0xa9, 0x65, 0xb5, 0xe9, 0xc3, 0xbe,
	0xd7, 0xb5, 0x31, 0xbd, 0xe4, 0xf6, 0x20, 0x13, 0x43, 0x2a, 0x7d, 0x48, 0xf6, 0x95, 0xd6, 0xa0,
	0xa3, 0x3d, 0x9c, 0x39, 0xe8, 0xda, 0x78, 0xea, 0x67, 0xb0, 0xcd, 0x84, 0x87, 0xa9, 0x24, 0x63,
	0x0c, 0x79, 0x3c, 0x0f, 0x16, 0xe1, 0xf2, 0xe5, 0x75, 0x42, 0x1e, 0x8a, 0xf0, 0x71, 0xe2, 0xd5,
	0xd9, 0x0d, 0xfc, 0x34, 0x9e, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x89, 0x1a, 0x5f, 0xe6,
	0x03, 0x00, 0x00,
}
