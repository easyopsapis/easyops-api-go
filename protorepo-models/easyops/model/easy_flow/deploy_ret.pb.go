// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deploy_ret.proto

package easy_flow

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//部署返回
type DeployRet struct {
	//
	//任务Id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//应用Id
	AppId string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId" form:"appId"`
	//
	//app名字
	AppName string `protobuf:"bytes,3,opt,name=appName,proto3" json:"appName" form:"appName"`
	//
	//操作人
	Operator string `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator" form:"operator"`
	//
	//客户id
	Org int32 `protobuf:"varint,5,opt,name=org,proto3" json:"org" form:"org"`
	//
	//目标列表
	TargetList []*TargetResult `protobuf:"bytes,6,rep,name=targetList,proto3" json:"targetList" form:"targetList"`
	//
	//分批数量
	BatchNum int32 `protobuf:"varint,7,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//分批间隔
	BatchInterval int32 `protobuf:"varint,8,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	//
	//是否暂停失败
	FailedStop string `protobuf:"bytes,9,opt,name=failedStop,proto3" json:"failedStop" form:"failedStop"`
	//
	//状态
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status" form:"status"`
	//
	//错误码
	Code int32 `protobuf:"varint,11,opt,name=code,proto3" json:"code" form:"code"`
	//
	//使用次数
	UsedTime int32 `protobuf:"varint,12,opt,name=usedTime,proto3" json:"usedTime" form:"usedTime"`
	//
	//开始时间
	StartTime string `protobuf:"bytes,13,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//结束时间
	EndTime string `protobuf:"bytes,14,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	//
	//集群Id
	ClusterId string `protobuf:"bytes,15,opt,name=clusterId,proto3" json:"clusterId" form:"clusterId"`
	//
	//配置包Id
	ConfigPackageId string `protobuf:"bytes,16,opt,name=configPackageId,proto3" json:"configPackageId" form:"configPackageId"`
	//
	//标签
	Labels               *DeployLabel `protobuf:"bytes,17,opt,name=labels,proto3" json:"labels" form:"labels"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeployRet) Reset()         { *m = DeployRet{} }
func (m *DeployRet) String() string { return proto.CompactTextString(m) }
func (*DeployRet) ProtoMessage()    {}
func (*DeployRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe57fbc8dd4a84e, []int{0}
}
func (m *DeployRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployRet.Unmarshal(m, b)
}
func (m *DeployRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployRet.Marshal(b, m, deterministic)
}
func (m *DeployRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployRet.Merge(m, src)
}
func (m *DeployRet) XXX_Size() int {
	return xxx_messageInfo_DeployRet.Size(m)
}
func (m *DeployRet) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployRet.DiscardUnknown(m)
}

var xxx_messageInfo_DeployRet proto.InternalMessageInfo

func (m *DeployRet) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DeployRet) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeployRet) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *DeployRet) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *DeployRet) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *DeployRet) GetTargetList() []*TargetResult {
	if m != nil {
		return m.TargetList
	}
	return nil
}

func (m *DeployRet) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *DeployRet) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

func (m *DeployRet) GetFailedStop() string {
	if m != nil {
		return m.FailedStop
	}
	return ""
}

func (m *DeployRet) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeployRet) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeployRet) GetUsedTime() int32 {
	if m != nil {
		return m.UsedTime
	}
	return 0
}

func (m *DeployRet) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *DeployRet) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *DeployRet) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeployRet) GetConfigPackageId() string {
	if m != nil {
		return m.ConfigPackageId
	}
	return ""
}

func (m *DeployRet) GetLabels() *DeployLabel {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*DeployRet)(nil), "easy_flow.DeployRet")
}

func init() { proto.RegisterFile("deploy_ret.proto", fileDescriptor_5fe57fbc8dd4a84e) }

var fileDescriptor_5fe57fbc8dd4a84e = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xdb, 0xed, 0x4f, 0x26, 0xdb, 0x36, 0x1d, 0x4a, 0x19, 0xf5, 0xc6, 0x91, 0x89, 0x90,
	0xc3, 0xee, 0xc4, 0xf9, 0xe9, 0xae, 0xd8, 0x5c, 0x10, 0xb5, 0xc0, 0x45, 0xa4, 0x55, 0x85, 0x86,
	0x0a, 0xa4, 0x7a, 0xd3, 0x6a, 0x62, 0x4f, 0xbc, 0x51, 0x9d, 0x8e, 0x65, 0x4f, 0xb6, 0xda, 0x36,
	0x79, 0x01, 0x9e, 0x81, 0x27, 0x40, 0xe2, 0x2d, 0x80, 0xc7, 0xf0, 0x4a, 0x48, 0xbc, 0x80, 0x9f,
	0x00, 0x79, 0xc6, 0x76, 0xdc, 0x8a, 0x1b, 0xa4, 0xe5, 0x82, 0x2b, 0xcf, 0x39, 0xdf, 0xf7, 0x9d,
	0x39, 0xdf, 0x1c, 0x7b, 0x0c, 0x6a, 0x2e, 0x0b, 0x7c, 0xfe, 0xee, 0x32, 0x64, 0xa2, 0x15, 0x84,
	0x5c, 0x70, 0x58, 0x61, 0x34, 0x7a, 0x77, 0x39, 0xf1, 0xf9, 0xcd, 0x21, 0xf6, 0xa6, 0xe2, 0xcd,
	0x7c, 0xdc, 0x72, 0xf8, 0xcc, 0xf2, 0xb8, 0xc7, 0x2d, 0xc9, 0x18, 0xcf, 0x27, 0x32, 0x92, 0x81,
	0x5c, 0x29, 0xe5, 0xe1, 0x8f, 0x1e, 0x6f, 0xa5, 0x62, 0x1e, 0x44, 0x2d, 0x9f, 0x3b, 0xd4, 0xb7,
	0x1c, 0x7e, 0x2d, 0x42, 0xea, 0x88, 0x48, 0x29, 0x43, 0x16, 0x70, 0x3c, 0xe3, 0x2e, 0xf3, 0x23,
	0x2b, 0x23, 0x5a, 0x32, 0xb4, 0x8a, 0x3d, 0x2d, 0x41, 0x43, 0x8f, 0x89, 0xcb, 0x90, 0x45, 0x73,
	0x3f, 0x6b, 0xe9, 0xf0, 0x87, 0x0f, 0x51, 0x38, 0x33, 0xea, 0xd3, 0x31, 0xf3, 0xb3, 0xba, 0x2f,
	0x4a, 0xfe, 0x66, 0x37, 0x53, 0x71, 0xc5, 0x6f, 0x2c, 0x8f, 0x63, 0x09, 0xe2, 0xb7, 0xd4, 0x9f,
	0xba, 0x54, 0xf0, 0x30, 0xb2, 0x8a, 0xa5, 0xd2, 0x19, 0xbf, 0x57, 0x41, 0xe5, 0x1b, 0x59, 0x8e,
	0x30, 0x01, 0x9b, 0x60, 0x43, 0xd0, 0xe8, 0x6a, 0xe8, 0x22, 0xad, 0xae, 0x99, 0x95, 0x93, 0xbd,
	0x24, 0xd6, 0xb7, 0x27, 0x3c, 0x9c, 0xf5, 0x0d, 0x95, 0x37, 0x48, 0x46, 0x80, 0x7d, 0xb0, 0x4e,
	0x83, 0x60, 0xe8, 0xa2, 0x47, 0x92, 0xd9, 0x48, 0x62, 0xfd, 0x89, 0x62, 0xca, 0xb4, 0xf1, 0xe7,
	0x7b, 0xbd, 0x06, 0x76, 0x2e, 0xec, 0x36, 0x7e, 0x49, 0xf1, 0xed, 0xe8, 0xae, 0xd3, 0x5b, 0x36,
	0x88, 0x92, 0xc0, 0x67, 0x60, 0x93, 0x06, 0xc1, 0x29, 0x9d, 0x31, 0xb4, 0x26, 0xd5, 0x30, 0x89,
	0xf5, 0x9d, 0x42, 0x9d, 0x02, 0x06, 0xc9, 0x29, 0x70, 0x08, 0xb6, 0x78, 0xc0, 0xc2, 0xb4, 0x69,
	0xf4, 0x58, 0xd2, 0x71, 0x12, 0xeb, 0xbb, 0x8a, 0x9e, 0x23, 0xe9, 0x7e, 0x07, 0x60, 0xff, 0xc2,
	0x3e, 0xc6, 0xe7, 0x14, 0xdf, 0x5e, 0xe2, 0xd1, 0xeb, 0x9b, 0xbb, 0xde, 0xb3, 0x17, 0x47, 0xcb,
	0x06, 0x29, 0xe4, 0xf0, 0x73, 0xb0, 0xc6, 0x43, 0x0f, 0xad, 0xd7, 0x35, 0x73, 0xfd, 0x64, 0x3f,
	0x89, 0x75, 0x90, 0x55, 0x09, 0xbd, 0xb4, 0xc0, 0xa3, 0xda, 0x47, 0x24, 0x25, 0xc0, 0x53, 0x00,
	0xd4, 0xf0, 0x5e, 0x4d, 0x23, 0x81, 0x36, 0xea, 0x6b, 0x66, 0xb5, 0xfb, 0x69, 0xab, 0x18, 0x40,
	0xeb, 0x4c, 0x82, 0x44, 0x0e, 0xf6, 0xe4, 0x93, 0x24, 0xd6, 0xf7, 0xf2, 0x43, 0xca, 0x45, 0x06,
	0x29, 0x55, 0x80, 0x16, 0xd8, 0x1a, 0x53, 0xe1, 0xbc, 0x39, 0x9d, 0xcf, 0xd0, 0xa6, 0xdc, 0xfc,
	0xe3, 0x95, 0x85, 0x1c, 0x31, 0x48, 0x41, 0x82, 0x5f, 0x81, 0x6d, 0xb9, 0x1e, 0x5e, 0x0b, 0x16,
	0xbe, 0xa5, 0x3e, 0xda, 0x92, 0x2a, 0x94, 0xc4, 0xfa, 0x7e, 0x49, 0x95, 0xc3, 0x06, 0xb9, 0x4f,
	0x87, 0xcf, 0x01, 0x98, 0xd0, 0xa9, 0xcf, 0xdc, 0xef, 0x05, 0x0f, 0x50, 0x45, 0x9e, 0x5a, 0xa9,
	0xcf, 0x15, 0x66, 0x90, 0x12, 0x31, 0x9d, 0x7f, 0x24, 0xa8, 0x98, 0x47, 0x08, 0x3c, 0x9c, 0xbf,
	0xca, 0x1b, 0x24, 0x23, 0xc0, 0xcf, 0xc0, 0x63, 0x87, 0xbb, 0x0c, 0x55, 0x65, 0x63, 0xbb, 0x49,
	0xac, 0x57, 0x15, 0x31, 0xcd, 0x1a, 0x44, 0x82, 0xa9, 0xef, 0x79, 0xc4, 0xdc, 0xb3, 0xe9, 0x8c,
	0xa1, 0x27, 0x0f, 0x7d, 0xe7, 0x88, 0x41, 0x0a, 0x12, 0xfc, 0x4b, 0x03, 0x95, 0x48, 0xd0, 0x50,
	0x48, 0xc9, 0xb6, 0x6c, 0xe2, 0x0f, 0x2d, 0x89, 0xf5, 0x5a, 0xd1, 0x85, 0xc2, 0xd2, 0x71, 0xfd,
	0xaa, 0x81, 0x5f, 0xb4, 0x0b, 0xd3, 0x1c, 0xf4, 0xed, 0x0e, 0x7e, 0x39, 0x4a, 0x5f, 0xb5, 0xd1,
	0x17, 0xcd, 0x81, 0x7c, 0xde, 0x1d, 0x2d, 0x9b, 0xd8, 0xec, 0xd8, 0x6d, 0xdc, 0x1d, 0x2d, 0xda,
	0x12, 0x6f, 0x62, 0xb3, 0x67, 0xb7, 0x71, 0x27, 0x8f, 0x17, 0x76, 0x07, 0x77, 0x95, 0xaa, 0x69,
	0x9f, 0xd5, 0x47, 0x66, 0xd7, 0x6e, 0xe3, 0xde, 0x68, 0x21, 0x39, 0x2a, 0xdd, 0x37, 0xed, 0x36,
	0x7e, 0x9e, 0x07, 0xab, 0xb5, 0xf9, 0xba, 0x25, 0x9f, 0x4f, 0x9b, 0x03, 0xf3, 0x7c, 0x61, 0x3f,
	0xc5, 0x23, 0x73, 0xd0, 0xff, 0x07, 0x79, 0x49, 0x3d, 0x68, 0x90, 0x95, 0x33, 0xf8, 0x5e, 0x03,
	0x9b, 0xec, 0x5a, 0x1d, 0xcc, 0x8e, 0x74, 0xf9, 0x9b, 0xb6, 0xfa, 0x06, 0x32, 0xe4, 0xff, 0xe8,
	0x31, 0x77, 0x05, 0x7f, 0xd2, 0x40, 0xc5, 0xf1, 0xe7, 0x91, 0x60, 0xe1, 0xd0, 0x45, 0xbb, 0xd2,
	0xa3, 0xbf, 0x1a, 0x64, 0x01, 0xa5, 0x26, 0xcf, 0x00, 0xb9, 0xb0, 0x07, 0x14, 0xdf, 0x1e, 0xe3,
	0x73, 0x69, 0xec, 0xcb, 0x25, 0xbe, 0x17, 0x1f, 0xfd, 0xcb, 0xb8, 0xd3, 0x5d, 0x36, 0xc8, 0x6a,
	0x7b, 0xf8, 0xb3, 0x06, 0x76, 0x1d, 0x7e, 0x3d, 0x99, 0x7a, 0xdf, 0x51, 0xe7, 0x8a, 0x7a, 0x6c,
	0xe8, 0xa2, 0x9a, 0x6c, 0x29, 0x4c, 0x62, 0xfd, 0x20, 0x7f, 0x71, 0xef, 0x11, 0xfe, 0xbb, 0xc6,
	0x1e, 0xb6, 0x02, 0x8f, 0xc1, 0x86, 0xbc, 0xcb, 0x23, 0xb4, 0x57, 0xd7, 0xcc, 0x6a, 0xf7, 0xa0,
	0x74, 0xd5, 0xa8, 0xcb, 0xf9, 0x55, 0x0a, 0x97, 0x3f, 0x47, 0xc5, 0x37, 0x48, 0x26, 0x3c, 0xf9,
	0xf6, 0xfc, 0xeb, 0x0f, 0xf0, 0x67, 0x19, 0x6f, 0x48, 0x66, 0xef, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x83, 0x0f, 0x2d, 0x4c, 0x07, 0x00, 0x00,
}