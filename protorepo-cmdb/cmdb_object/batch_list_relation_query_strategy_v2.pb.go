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
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" form:"type"`
	//
	//策略名称(模糊匹配)
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
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

func (m *BatchListRelationQueryStrategyV2RequestRequest) GetName() string {
	if m != nil {
		return m.Name
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
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x25, 0xdb, 0xad, 0xb8, 0x93, 0x15, 0xeb, 0x50, 0x25, 0x2c, 0x48, 0x96, 0x11, 0x64, 0x41,
	0x9a, 0xe0, 0x2a, 0x22, 0xe2, 0xd3, 0xa2, 0x0f, 0x05, 0x5f, 0x1c, 0x41, 0x5f, 0x84, 0x30, 0x49,
	0xa6, 0x69, 0x24, 0xc9, 0x4d, 0x67, 0x66, 0xc5, 0xfd, 0x03, 0xff, 0x42, 0xf0, 0x0b, 0xfc, 0xa2,
	0x7c, 0x44, 0xbe, 0x40, 0xe6, 0x4e, 0xda, 0xc6, 0xa2, 0x68, 0xf1, 0x69, 0x72, 0xef, 0x39, 0x77,
	0xe6, 0x9e, 0x7b, 0x6e, 0xc8, 0xa3, 0x54, 0x98, 0xec, 0x34, 0xa9, 0x4a, 0x6d, 0x12, 0x25, 0x2b,
	0x61, 0x4a, 0x68, 0x92, 0xb3, 0xad, 0x54, 0xbb, 0x44, 0x1b, 0x25, 0x8c, 0x2c, 0x76, 0xc9, 0xe7,
	0x75, 0xd4, 0x2a, 0x30, 0x40, 0xfd, 0xac, 0xce, 0xd3, 0x04, 0xd2, 0x4f, 0x32, 0x33, 0x8b, 0xa3,
	0xa2, 0x34, 0xa7, 0xdb, 0x34, 0xca, 0xa0, 0x8e, 0x0b, 0x28, 0x20, 0x46, 0x4e, 0xba, 0x3d, 0xc1,
	0x08, 0x03, 0xfc, 0x72, 0xb5, 0x8b, 0xa4, 0x80, 0x48, 0x0a, 0xbd, 0x83, 0x56, 0x47, 0x15, 0x64,
	0xa2, 0x8a, 0x33, 0x68, 0x8c, 0x12, 0x99, 0xd1, 0xae, 0x52, 0xc9, 0x16, 0x8e, 0x6a, 0xc8, 0x65,
	0xa5, 0xe3, 0x81, 0x18, 0x63, 0x18, 0xdb, 0x57, 0xe3, 0xbf, 0x35, 0xc7, 0x7e, 0x78, 0x24, 0xda,
	0x58, 0x31, 0x6f, 0x4a, 0x6d, 0xf8, 0xc0, 0x7e, 0x6b, 0xc9, 0xef, 0x06, 0xee, 0xfb, 0x35, 0x97,
	0x67, 0x5b, 0x69, 0x51, 0x3c, 0xe8, 0x53, 0x42, 0x9c, 0x98, 0xa4, 0xcc, 0x75, 0xe0, 0x2d, 0xbd,
	0xd5, 0x6c, 0x73, 0xb7, 0xef, 0xc2, 0x3b, 0x27, 0xa0, 0xea, 0x17, 0xec, 0x12, 0x63, 0x7c, 0xe6,
	0x82, 0xe3, 0x5c, 0xd3, 0x07, 0x64, 0x6a, 0x76, 0xad, 0x0c, 0x26, 0xc8, 0xbf, 0xdd, 0x77, 0xa1,
	0xef, 0xf8, 0x36, 0xcb, 0x38, 0x82, 0x96, 0xd4, 0x88, 0x5a, 0x06, 0x7b, 0x57, 0x49, 0x36, 0xcb,
	0x38, 0x82, 0xac, 0x9b, 0x90, 0xf8, 0x9f, 0x5b, 0xd6, 0x2d, 0x34, 0x5a, 0xd2, 0xaf, 0x1e, 0xb9,
	0x75, 0x21, 0x3e, 0x17, 0x46, 0x04, 0xde, 0x72, 0x6f, 0xe5, 0xaf, 0x8f, 0xa3, 0x91, 0x39, 0xd1,
	0x35, 0x6f, 0x8d, 0xce, 0x91, 0x57, 0xc2, 0x88, 0x4d, 0xd0, 0x77, 0xe1, 0xa1, 0xeb, 0xf6, 0x97,
	0x97, 0x18, 0x9f, 0xeb, 0x11, 0x6f, 0xf1, 0xcd, 0x23, 0xf3, 0x71, 0x21, 0x7d, 0x4c, 0x66, 0x17,
	0x33, 0x1b, 0xc6, 0x79, 0xd8, 0x77, 0xe1, 0xc1, 0x95, 0x71, 0x32, 0x7e, 0xf3, 0x7c, 0x9a, 0xf4,
	0xe3, 0x48, 0x8d, 0x5d, 0xc2, 0x60, 0x82, 0x6a, 0xee, 0xa3, 0x9a, 0xe8, 0x0f, 0xdd, 0xff, 0xb6,
	0x43, 0x5b, 0x3d, 0xea, 0xd0, 0x8a, 0x67, 0xdf, 0x27, 0xe4, 0xd9, 0x35, 0x47, 0xf1, 0x41, 0x89,
	0xb6, 0x95, 0xca, 0x1a, 0x98, 0x41, 0x2e, 0x51, 0xc6, 0xfe, 0xd8, 0x40, 0x9b, 0x65, 0x1c, 0x41,
	0xfa, 0x9c, 0xf8, 0xf6, 0x7c, 0xfd, 0xa5, 0xad, 0x44, 0xd9, 0x0c, 0x1b, 0x71, 0xaf, 0xef, 0x42,
	0x7a, 0xc9, 0x1d, 0x40, 0xc6, 0xc7, 0x54, 0xfa, 0x90, 0xec, 0x4b, 0xa5, 0x40, 0x0d, 0x0b, 0x72,
	0xd0, 0x77, 0xe1, 0xdc, 0xd5, 0x60, 0x9a, 0x71, 0x07, 0x53, 0x41, 0xa6, 0x68, 0xf2, 0x74, 0xe9,
	0xad, 0xfc, 0xf5, 0xcb, 0xff, 0x31, 0x79, 0x2c, 0xc2, 0xd9, 0x89, 0x57, 0xa7, 0x37, 0xf0, 0xff,
	0x79, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xf6, 0xc9, 0xd8, 0x0b, 0x04, 0x00, 0x00,
}
