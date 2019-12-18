// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter_strategy.proto

package resource_manage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	console "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/console"
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
//过滤策略
type FilterStrategy struct {
	//
	//策略Id
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//策略名称
	StrategyName string `protobuf:"bytes,2,opt,name=strategyName,proto3" json:"strategyName" form:"strategyName"`
	//
	//所属模型Id
	StrategyObjectId string `protobuf:"bytes,3,opt,name=strategyObjectId,proto3" json:"strategyObjectId" form:"strategyObjectId"`
	//
	//实例查询策略
	Query *console.CmdbQueryStrategy `protobuf:"bytes,4,opt,name=query,proto3" json:"query" form:"query"`
	//
	//筛选规则, or 逻辑
	Filter []*FilterConditionGroup `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter" form:"filter"`
	//
	//定时策略
	Crontab string `protobuf:"bytes,6,opt,name=crontab,proto3" json:"crontab" form:"crontab"`
	//
	//创建时间
	Ctime string `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime string `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//创建者
	Creator string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//修改者
	Modifier string `protobuf:"bytes,10,opt,name=modifier,proto3" json:"modifier" form:"modifier"`
	//
	//下次执行时间
	NextExecTime string `protobuf:"bytes,11,opt,name=nextExecTime,proto3" json:"nextExecTime" form:"nextExecTime"`
	//
	//定时策略
	Enable bool `protobuf:"varint,12,opt,name=enable,proto3" json:"enable" form:"enable"`
	//
	//更新白名单
	UpdateAuthorizers []string `protobuf:"bytes,13,rep,name=updateAuthorizers,proto3" json:"updateAuthorizers" form:"updateAuthorizers"`
	//
	//读取白名单
	ReadAuthorizers []string `protobuf:"bytes,14,rep,name=readAuthorizers,proto3" json:"readAuthorizers" form:"readAuthorizers"`
	//
	//通知人/组
	NotifyUsers []string `protobuf:"bytes,15,rep,name=notifyUsers,proto3" json:"notifyUsers" form:"notifyUsers"`
	//
	//通知方式
	NotifyMethods []string `protobuf:"bytes,16,rep,name=notifyMethods,proto3" json:"notifyMethods" form:"notifyMethods"`
	//
	//org
	Org                  int32    `protobuf:"varint,17,opt,name=org,proto3" json:"org" form:"org"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterStrategy) Reset()         { *m = FilterStrategy{} }
func (m *FilterStrategy) String() string { return proto.CompactTextString(m) }
func (*FilterStrategy) ProtoMessage()    {}
func (*FilterStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_f543ff0a595bb6c3, []int{0}
}
func (m *FilterStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStrategy.Unmarshal(m, b)
}
func (m *FilterStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStrategy.Marshal(b, m, deterministic)
}
func (m *FilterStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStrategy.Merge(m, src)
}
func (m *FilterStrategy) XXX_Size() int {
	return xxx_messageInfo_FilterStrategy.Size(m)
}
func (m *FilterStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStrategy proto.InternalMessageInfo

func (m *FilterStrategy) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *FilterStrategy) GetStrategyName() string {
	if m != nil {
		return m.StrategyName
	}
	return ""
}

func (m *FilterStrategy) GetStrategyObjectId() string {
	if m != nil {
		return m.StrategyObjectId
	}
	return ""
}

func (m *FilterStrategy) GetQuery() *console.CmdbQueryStrategy {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *FilterStrategy) GetFilter() []*FilterConditionGroup {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *FilterStrategy) GetCrontab() string {
	if m != nil {
		return m.Crontab
	}
	return ""
}

func (m *FilterStrategy) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *FilterStrategy) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *FilterStrategy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FilterStrategy) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

func (m *FilterStrategy) GetNextExecTime() string {
	if m != nil {
		return m.NextExecTime
	}
	return ""
}

func (m *FilterStrategy) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *FilterStrategy) GetUpdateAuthorizers() []string {
	if m != nil {
		return m.UpdateAuthorizers
	}
	return nil
}

func (m *FilterStrategy) GetReadAuthorizers() []string {
	if m != nil {
		return m.ReadAuthorizers
	}
	return nil
}

func (m *FilterStrategy) GetNotifyUsers() []string {
	if m != nil {
		return m.NotifyUsers
	}
	return nil
}

func (m *FilterStrategy) GetNotifyMethods() []string {
	if m != nil {
		return m.NotifyMethods
	}
	return nil
}

func (m *FilterStrategy) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterStrategy)(nil), "resource_manage.FilterStrategy")
}

func init() { proto.RegisterFile("filter_strategy.proto", fileDescriptor_f543ff0a595bb6c3) }

var fileDescriptor_f543ff0a595bb6c3 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x6b, 0xdb, 0x3a,
	0x18, 0xc6, 0xc9, 0xc9, 0x49, 0xda, 0x2a, 0x6d, 0xd2, 0xa8, 0xa7, 0xad, 0xc8, 0x39, 0x60, 0x23,
	0x38, 0x23, 0x83, 0xd5, 0x86, 0x8e, 0xc1, 0xd8, 0x60, 0x30, 0x77, 0x5b, 0xe9, 0xc6, 0xfe, 0x69,
	0xdb, 0xcd, 0x76, 0x11, 0x64, 0x5b, 0x71, 0x3d, 0x62, 0xbf, 0x99, 0x2c, 0x43, 0xb3, 0x0f, 0xba,
	0x4b, 0x7f, 0x08, 0x7f, 0x82, 0x61, 0xc9, 0x59, 0x1d, 0xe7, 0x6e, 0xec, 0x2a, 0x7e, 0xfd, 0xfc,
	0x9e, 0x47, 0x7a, 0x1f, 0x43, 0xd0, 0xf1, 0x3c, 0x5e, 0x28, 0x21, 0x67, 0x99, 0x92, 0x5c, 0x89,
	0x68, 0xe5, 0x2c, 0x25, 0x28, 0xc0, 0x23, 0x29, 0x32, 0xc8, 0x65, 0x20, 0x66, 0x09, 0x4f, 0x79,
	0x24, 0x26, 0x67, 0x51, 0xac, 0xae, 0x73, 0xdf, 0x09, 0x20, 0x71, 0x23, 0x88, 0xc0, 0xd5, 0x9c,
	0x9f, 0xcf, 0xf5, 0xa4, 0x07, 0xfd, 0x64, 0xfc, 0x93, 0x2f, 0x11, 0x38, 0x82, 0x67, 0x2b, 0x58,
	0x66, 0xce, 0x02, 0x02, 0xbe, 0x70, 0x03, 0x48, 0x95, 0xe4, 0x81, 0xca, 0x8c, 0x53, 0x8a, 0x25,
	0x9c, 0x25, 0x10, 0x8a, 0x45, 0xe6, 0xd6, 0xa0, 0xab, 0xc7, 0x0a, 0xcc, 0x60, 0x21, 0xdc, 0x20,
	0x09, 0xfd, 0xd9, 0xb7, 0x5c, 0xc8, 0x55, 0xeb, 0x72, 0x93, 0xf9, 0xef, 0x87, 0xb7, 0xd6, 0x72,
	0xeb, 0xed, 0x03, 0x48, 0xc3, 0x58, 0xc5, 0x90, 0xce, 0x22, 0x09, 0xf9, 0xd2, 0x9c, 0x43, 0x7f,
	0xec, 0xa0, 0xe1, 0x0b, 0x0d, 0x7c, 0xa8, 0x2f, 0x80, 0x1f, 0x20, 0x14, 0xa7, 0x99, 0xe2, 0x69,
	0x20, 0xae, 0x42, 0xd2, 0xb1, 0x3b, 0xd3, 0x3d, 0xef, 0xb8, 0x2c, 0xac, 0xf1, 0x1c, 0x64, 0xf2,
	0x88, 0xde, 0x6a, 0x94, 0x35, 0x40, 0xfc, 0x18, 0xed, 0xaf, 0x77, 0x78, 0xc3, 0x13, 0x41, 0xfe,
	0xd2, 0xc6, 0xd3, 0xb2, 0xb0, 0x8e, 0x8c, 0xb1, 0xa9, 0x52, 0xb6, 0x01, 0xe3, 0x4b, 0x74, 0xb8,
	0x9e, 0xdf, 0xfa, 0x5f, 0x45, 0xa0, 0xae, 0x42, 0xd2, 0xd5, 0x01, 0xff, 0x96, 0x85, 0x75, 0xba,
	0x19, 0xb0, 0x26, 0x28, 0xdb, 0x32, 0x61, 0x0f, 0xf5, 0x74, 0x9f, 0xe4, 0x6f, 0xbb, 0x33, 0x1d,
	0x9c, 0x4f, 0x9c, 0xba, 0x6a, 0xe7, 0x22, 0x09, 0xfd, 0xf7, 0x95, 0xb2, 0xde, 0xd3, 0x3b, 0x2c,
	0x0b, 0x6b, 0xdf, 0x24, 0x6b, 0x0b, 0x65, 0xc6, 0x8a, 0xdf, 0xa1, 0xbe, 0xe9, 0x8c, 0xf4, 0xec,
	0xee, 0x74, 0x70, 0xfe, 0xbf, 0xd3, 0xaa, 0xd4, 0x31, 0x8d, 0x5d, 0xac, 0x1b, 0xbd, 0xac, 0x0a,
	0xf5, 0xc6, 0x65, 0x61, 0x1d, 0x98, 0x3c, 0x63, 0xa7, 0xac, 0xce, 0xc1, 0xf7, 0xd0, 0x4e, 0x20,
	0x21, 0x55, 0xdc, 0x27, 0x7d, 0xbd, 0x15, 0x2e, 0x0b, 0x6b, 0x68, 0xd8, 0x5a, 0xa0, 0x6c, 0x8d,
	0xe0, 0x3b, 0xa8, 0x17, 0xa8, 0x38, 0x11, 0x64, 0x47, 0xb3, 0x8d, 0x7b, 0xea, 0xd7, 0x94, 0x19,
	0xb9, 0xe2, 0x12, 0xcd, 0xed, 0xb6, 0xb9, 0xa4, 0xe6, 0xf4, 0xaf, 0x39, 0x5d, 0x70, 0x05, 0x92,
	0xec, 0x6d, 0x9f, 0xae, 0x05, 0x7d, 0xba, 0x7e, 0xc2, 0x2e, 0xda, 0x4d, 0x20, 0x8c, 0xe7, 0xb1,
	0x90, 0x04, 0x69, 0xfc, 0xa8, 0x2c, 0xac, 0x51, 0x1d, 0x5c, 0x2b, 0x94, 0xfd, 0x82, 0xaa, 0x0f,
	0x9f, 0x8a, 0x1b, 0xf5, 0xfc, 0x46, 0x04, 0x1f, 0xab, 0xdb, 0x0c, 0xda, 0x1f, 0xbe, 0xa9, 0x52,
	0xb6, 0x01, 0xe3, 0xbb, 0xa8, 0x2f, 0x52, 0xee, 0x2f, 0x04, 0xd9, 0xb7, 0x3b, 0xd3, 0xdd, 0x66,
	0x89, 0xe6, 0x3d, 0x65, 0x35, 0x80, 0x5f, 0xa2, 0x71, 0xbe, 0x0c, 0xb9, 0x12, 0x4f, 0x73, 0x75,
	0x0d, 0x32, 0xfe, 0x2e, 0x64, 0x46, 0x0e, 0xec, 0xee, 0x74, 0xcf, 0xfb, 0xaf, 0x2c, 0x2c, 0x62,
	0x5c, 0x5b, 0x08, 0x65, 0xdb, 0x36, 0xfc, 0x0c, 0x8d, 0xa4, 0xe0, 0x61, 0x33, 0x69, 0xa8, 0x93,
	0x26, 0x65, 0x61, 0x9d, 0x98, 0xa4, 0x16, 0x40, 0x59, 0xdb, 0x82, 0x1f, 0xa2, 0x41, 0x0a, 0x2a,
	0x9e, 0xaf, 0x3e, 0x65, 0x55, 0xc2, 0x48, 0x27, 0x9c, 0x94, 0x85, 0x85, 0xeb, 0xc5, 0x6f, 0x45,
	0xca, 0x9a, 0x28, 0x7e, 0x82, 0x0e, 0xcc, 0xf8, 0x5a, 0xa8, 0x6b, 0x08, 0x33, 0x72, 0xa8, 0xbd,
	0xa4, 0x2c, 0xac, 0x7f, 0x9a, 0xde, 0x5a, 0xa6, 0x6c, 0x13, 0xc7, 0x36, 0xea, 0x82, 0x8c, 0xc8,
	0xd8, 0xee, 0x4c, 0x7b, 0xde, 0xb0, 0x2c, 0x2c, 0x64, 0x5c, 0x20, 0x23, 0xca, 0x2a, 0xc9, 0x7b,
	0xf5, 0xf9, 0xea, 0x8f, 0xfd, 0x85, 0xf8, 0x7d, 0xcd, 0xdf, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0xdb, 0x55, 0x44, 0x4a, 0x05, 0x00, 0x00,
}
