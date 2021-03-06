// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter_strategy_instance.proto

package resource_manage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//策略执行实例
type FilterStrategyInstance struct {
	//
	//策略执行实例Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//所属策略instanceId
	StrategyId string `protobuf:"bytes,2,opt,name=strategyId,proto3" json:"strategyId" form:"strategyId"`
	//
	//开始时间
	Ctime string `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//结束时间
	Etime string `protobuf:"bytes,4,opt,name=etime,proto3" json:"etime" form:"etime"`
	//
	//实例总数
	Total int32 `protobuf:"varint,5,opt,name=total,proto3" json:"total" form:"total"`
	//
	//不合规实例数
	FailCount int32 `protobuf:"varint,6,opt,name=failCount,proto3" json:"failCount" form:"failCount"`
	//
	//状态
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status" form:"status"`
	//
	//status为failed时有此信息
	Message              string   `protobuf:"bytes,8,opt,name=message,proto3" json:"message" form:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterStrategyInstance) Reset()         { *m = FilterStrategyInstance{} }
func (m *FilterStrategyInstance) String() string { return proto.CompactTextString(m) }
func (*FilterStrategyInstance) ProtoMessage()    {}
func (*FilterStrategyInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b64305366ab0e74, []int{0}
}
func (m *FilterStrategyInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStrategyInstance.Unmarshal(m, b)
}
func (m *FilterStrategyInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStrategyInstance.Marshal(b, m, deterministic)
}
func (m *FilterStrategyInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStrategyInstance.Merge(m, src)
}
func (m *FilterStrategyInstance) XXX_Size() int {
	return xxx_messageInfo_FilterStrategyInstance.Size(m)
}
func (m *FilterStrategyInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStrategyInstance.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStrategyInstance proto.InternalMessageInfo

func (m *FilterStrategyInstance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FilterStrategyInstance) GetStrategyId() string {
	if m != nil {
		return m.StrategyId
	}
	return ""
}

func (m *FilterStrategyInstance) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *FilterStrategyInstance) GetEtime() string {
	if m != nil {
		return m.Etime
	}
	return ""
}

func (m *FilterStrategyInstance) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *FilterStrategyInstance) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *FilterStrategyInstance) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FilterStrategyInstance) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FilterStrategyInstance)(nil), "resource_manage.FilterStrategyInstance")
}

func init() { proto.RegisterFile("filter_strategy_instance.proto", fileDescriptor_8b64305366ab0e74) }

var fileDescriptor_8b64305366ab0e74 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd1, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x06, 0x60, 0xac, 0xd6, 0x76, 0xbd, 0xd4, 0xad, 0xbd, 0xb4, 0x45, 0x14, 0x5a, 0x99, 0x3d,
	0x04, 0x07, 0x62, 0x09, 0x12, 0x72, 0xc9, 0xd1, 0x81, 0x80, 0xc9, 0x4d, 0xb9, 0xe5, 0x62, 0xd6,
	0xd2, 0x48, 0x59, 0x90, 0x34, 0x66, 0x77, 0x75, 0xf0, 0xcb, 0xea, 0x21, 0x94, 0x17, 0x08, 0x9e,
	0x95, 0x62, 0xe3, 0x9b, 0x66, 0xfe, 0x6f, 0x98, 0x15, 0xc3, 0xfe, 0x67, 0xaa, 0xb0, 0xa0, 0xb7,
	0xc6, 0x6a, 0x69, 0x21, 0x3f, 0x6c, 0x55, 0x65, 0xac, 0xac, 0x12, 0x08, 0xf7, 0x1a, 0x2d, 0xf2,
	0x9f, 0x1a, 0x0c, 0xd6, 0x3a, 0x81, 0x6d, 0x29, 0x2b, 0x99, 0xc3, 0xdf, 0x55, 0xae, 0xec, 0x5b,
	0xbd, 0x0b, 0x13, 0x2c, 0xa3, 0x1c, 0x73, 0x8c, 0xc8, 0xed, 0xea, 0x8c, 0x2a, 0x2a, 0xe8, 0xcb,
	0xcd, 0x8b, 0x77, 0x8f, 0xfd, 0x79, 0xa2, 0x15, 0x2f, 0xdd, 0x86, 0x4d, 0xb7, 0x80, 0xff, 0x63,
	0x9e, 0x4a, 0xfd, 0xc1, 0x62, 0xb0, 0x9c, 0xac, 0xa7, 0x6d, 0x13, 0x4c, 0x32, 0xd4, 0xe5, 0x83,
	0x50, 0xa9, 0x88, 0x3d, 0x95, 0xf2, 0x7b, 0xc6, 0xfa, 0x47, 0x6d, 0x52, 0xdf, 0x23, 0xf6, 0xbb,
	0x6d, 0x82, 0xb9, 0x63, 0xa7, 0x4c, 0xc4, 0x67, 0x90, 0x5f, 0xb1, 0x61, 0x62, 0x55, 0x09, 0xfe,
	0x17, 0x9a, 0x98, 0xb5, 0x4d, 0xf0, 0xdd, 0x4d, 0x50, 0x5b, 0xc4, 0x2e, 0x3e, 0x3a, 0x20, 0xf7,
	0xf5, 0xd2, 0x41, 0xe7, 0xa0, 0x77, 0x16, 0xad, 0x2c, 0xfc, 0xe1, 0x62, 0xb0, 0x1c, 0x9e, 0x3b,
	0x6a, 0x8b, 0xd8, 0xc5, 0xfc, 0x96, 0x4d, 0x32, 0xa9, 0x8a, 0x47, 0xac, 0x2b, 0xeb, 0x8f, 0xc8,
	0xfe, 0x6a, 0x9b, 0x60, 0xe6, 0xec, 0x67, 0x24, 0xe2, 0x13, 0xe3, 0xd7, 0x6c, 0x64, 0xac, 0xb4,
	0xb5, 0xf1, 0xc7, 0xf4, 0x88, 0x79, 0xdb, 0x04, 0xd3, 0xfe, 0xf7, 0x8e, 0x7d, 0x11, 0x77, 0x80,
	0xdf, 0xb0, 0x71, 0x09, 0xc6, 0xc8, 0x1c, 0xfc, 0x6f, 0x64, 0x79, 0xdb, 0x04, 0x3f, 0x9c, 0xed,
	0x02, 0x11, 0xf7, 0x64, 0xfd, 0xfc, 0xba, 0xc9, 0x31, 0x04, 0x69, 0x0e, 0xb8, 0x37, 0x61, 0x81,
	0x89, 0x2c, 0xa2, 0x04, 0x2b, 0xab, 0x65, 0x62, 0x8d, 0xbb, 0x98, 0x86, 0x3d, 0xae, 0x4a, 0x4c,
	0xa1, 0x30, 0x51, 0x07, 0x23, 0x2a, 0xa3, 0x8b, 0x8b, 0xef, 0x46, 0xe4, 0xef, 0x3e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x53, 0xd9, 0x60, 0x2b, 0x02, 0x00, 0x00,
}
