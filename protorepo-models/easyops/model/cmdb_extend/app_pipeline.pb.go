// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_pipeline.proto

package cmdb_extend

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//流水线
type AppPipeline struct {
	//
	//名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//分类
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category" form:"category"`
	//
	//流水线Id
	FlowId string `protobuf:"bytes,3,opt,name=flowId,proto3" json:"flowId" form:"flowId"`
	//
	//流水线版本
	FlowVersion int32 `protobuf:"varint,4,opt,name=flowVersion,proto3" json:"flowVersion" form:"flowVersion"`
	//
	//模板Id
	TemplateId string `protobuf:"bytes,5,opt,name=templateId,proto3" json:"templateId" form:"templateId"`
	//
	//模板版本
	TemplateVersion int32 `protobuf:"varint,6,opt,name=templateVersion,proto3" json:"templateVersion" form:"templateVersion"`
	//
	//通知人
	Subscribers []string `protobuf:"bytes,7,rep,name=subscribers,proto3" json:"subscribers" form:"subscribers"`
	//
	//通知规则(WARNING/ALL)
	SubscribedChannel string `protobuf:"bytes,8,opt,name=subscribedChannel,proto3" json:"subscribedChannel" form:"subscribedChannel"`
	//
	//钩子触发规则, {"gitlabHooks":{"enabled":true},"subversionHooks":{"enabled":false}}
	Rules *types.Struct `protobuf:"bytes,9,opt,name=rules,proto3" json:"rules" form:"rules"`
	//
	//元数据, {"type":"development","desc":"xxx","appId":"5cc18b157f868"}
	Metadata             *types.Struct `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata" form:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AppPipeline) Reset()         { *m = AppPipeline{} }
func (m *AppPipeline) String() string { return proto.CompactTextString(m) }
func (*AppPipeline) ProtoMessage()    {}
func (*AppPipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_883216473cdf2e20, []int{0}
}
func (m *AppPipeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPipeline.Unmarshal(m, b)
}
func (m *AppPipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPipeline.Marshal(b, m, deterministic)
}
func (m *AppPipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPipeline.Merge(m, src)
}
func (m *AppPipeline) XXX_Size() int {
	return xxx_messageInfo_AppPipeline.Size(m)
}
func (m *AppPipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPipeline.DiscardUnknown(m)
}

var xxx_messageInfo_AppPipeline proto.InternalMessageInfo

func (m *AppPipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppPipeline) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *AppPipeline) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *AppPipeline) GetFlowVersion() int32 {
	if m != nil {
		return m.FlowVersion
	}
	return 0
}

func (m *AppPipeline) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *AppPipeline) GetTemplateVersion() int32 {
	if m != nil {
		return m.TemplateVersion
	}
	return 0
}

func (m *AppPipeline) GetSubscribers() []string {
	if m != nil {
		return m.Subscribers
	}
	return nil
}

func (m *AppPipeline) GetSubscribedChannel() string {
	if m != nil {
		return m.SubscribedChannel
	}
	return ""
}

func (m *AppPipeline) GetRules() *types.Struct {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *AppPipeline) GetMetadata() *types.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*AppPipeline)(nil), "cmdb_extend.AppPipeline")
}

func init() { proto.RegisterFile("app_pipeline.proto", fileDescriptor_883216473cdf2e20) }

var fileDescriptor_883216473cdf2e20 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0xa9, 0xdb, 0xd6, 0x76, 0xa2, 0xd4, 0x8e, 0xb8, 0x0e, 0x65, 0x21, 0x65, 0xbc, 0xd4,
	0xc3, 0x26, 0xa0, 0x08, 0xe2, 0x45, 0xac, 0x8a, 0xae, 0x27, 0x89, 0xe0, 0xc1, 0xcb, 0x32, 0x49,
	0x5e, 0xb3, 0x81, 0x49, 0xde, 0x30, 0x33, 0x41, 0xf7, 0x1f, 0xf5, 0x98, 0x3f, 0x22, 0x7f, 0x81,
	0x74, 0xa6, 0x69, 0xe3, 0x0a, 0xde, 0xe6, 0x9b, 0xef, 0xf7, 0xcd, 0xf7, 0xe0, 0x0d, 0xa1, 0x42,
	0xa9, 0x6b, 0x55, 0x2a, 0x90, 0x65, 0x0d, 0x91, 0xd2, 0x68, 0x91, 0x06, 0x59, 0x95, 0xa7, 0xd7,
	0xf0, 0xcb, 0x42, 0x9d, 0xaf, 0x2e, 0x8b, 0xd2, 0xde, 0x34, 0x69, 0x94, 0x61, 0x15, 0x17, 0x58,
	0x60, 0xec, 0x98, 0xb4, 0xd9, 0x39, 0xe5, 0x84, 0x3b, 0xf9, 0xec, 0xea, 0xa2, 0x40, 0x2c, 0x24,
	0x9c, 0x28, 0x63, 0x75, 0x93, 0x59, 0xef, 0xf2, 0xdf, 0x63, 0x12, 0xbc, 0x53, 0xea, 0xeb, 0xa1,
	0x8f, 0x3e, 0x23, 0xe3, 0x5a, 0x54, 0xc0, 0x46, 0xeb, 0xd1, 0x66, 0xbe, 0x5d, 0x74, 0x6d, 0x18,
	0xec, 0x50, 0x57, 0x6f, 0xf8, 0xfe, 0x96, 0x27, 0xce, 0xa4, 0x31, 0x99, 0x65, 0xc2, 0x42, 0x81,
	0xfa, 0x96, 0xdd, 0x73, 0xe0, 0xe3, 0xae, 0x0d, 0x17, 0x1e, 0xec, 0x1d, 0x9e, 0x1c, 0x21, 0xfa,
	0x9c, 0x4c, 0x77, 0x12, 0x7f, 0x5e, 0xe5, 0xec, 0xcc, 0xe1, 0xcb, 0xae, 0x0d, 0x1f, 0x7a, 0xdc,
	0xdf, 0xf3, 0xe4, 0x00, 0xd0, 0xd7, 0x24, 0xd8, 0x9f, 0xbe, 0x83, 0x36, 0x25, 0xd6, 0x6c, 0xbc,
	0x1e, 0x6d, 0x26, 0xdb, 0xf3, 0xae, 0x0d, 0xe9, 0x89, 0x3f, 0x98, 0x3c, 0x19, 0xa2, 0xf4, 0x15,
	0x21, 0x16, 0x2a, 0x25, 0x85, 0x85, 0xab, 0x9c, 0x4d, 0x5c, 0xd1, 0x93, 0xae, 0x0d, 0x97, 0x3e,
	0x78, 0xf2, 0x78, 0x32, 0x00, 0xe9, 0x07, 0xb2, 0xe8, 0x55, 0x5f, 0x3a, 0x75, 0xa5, 0xab, 0xae,
	0x0d, 0xcf, 0xff, 0xce, 0x1e, 0x8b, 0xef, 0x46, 0xf6, 0x63, 0x9b, 0x26, 0x35, 0x99, 0x2e, 0x53,
	0xd0, 0x86, 0xdd, 0x5f, 0x9f, 0x6d, 0xe6, 0xc3, 0xb1, 0x07, 0x26, 0x4f, 0x86, 0x28, 0xfd, 0x42,
	0x96, 0x47, 0x99, 0xbf, 0xbf, 0x11, 0x75, 0x0d, 0x92, 0xcd, 0xdc, 0xf4, 0x17, 0x5d, 0x1b, 0xb2,
	0x3b, 0xf9, 0x1e, 0xe1, 0xc9, 0xbf, 0x31, 0xfa, 0x96, 0x4c, 0x74, 0x23, 0xc1, 0xb0, 0xf9, 0x7a,
	0xb4, 0x09, 0x5e, 0x3c, 0x8d, 0xfc, 0xee, 0xa3, 0x7e, 0xf7, 0xd1, 0x37, 0xb7, 0xfb, 0xed, 0xa3,
	0xae, 0x0d, 0x1f, 0xf8, 0x87, 0x1d, 0xcf, 0x13, 0x9f, 0xa3, 0x9f, 0xc9, 0xac, 0x02, 0x2b, 0x72,
	0x61, 0x05, 0x23, 0xff, 0x7f, 0x63, 0xb0, 0xf2, 0x3e, 0xc2, 0x93, 0x63, 0x7a, 0xfb, 0xe9, 0xc7,
	0xc7, 0x02, 0x23, 0x10, 0xe6, 0x16, 0x95, 0x89, 0x24, 0x66, 0x42, 0xc6, 0x19, 0xd6, 0x56, 0x8b,
	0xcc, 0x1a, 0xff, 0x19, 0x35, 0x28, 0xbc, 0xac, 0x30, 0x07, 0x69, 0xe2, 0x03, 0x18, 0x3b, 0x19,
	0x0f, 0xbe, 0x7b, 0x3a, 0x75, 0xec, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xb6, 0x10,
	0x71, 0x18, 0x03, 0x00, 0x00,
}