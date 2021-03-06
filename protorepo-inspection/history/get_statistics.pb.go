// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_statistics.proto

package history

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inspection "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/inspection"
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
//GetStatistics请求
type GetStatisticsRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//巡检作业Id
	JobId                string   `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId" form:"jobId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatisticsRequest) Reset()         { *m = GetStatisticsRequest{} }
func (m *GetStatisticsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatisticsRequest) ProtoMessage()    {}
func (*GetStatisticsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81af0e9ed6c23, []int{0}
}
func (m *GetStatisticsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatisticsRequest.Unmarshal(m, b)
}
func (m *GetStatisticsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatisticsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatisticsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatisticsRequest.Merge(m, src)
}
func (m *GetStatisticsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatisticsRequest.Size(m)
}
func (m *GetStatisticsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatisticsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatisticsRequest proto.InternalMessageInfo

func (m *GetStatisticsRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *GetStatisticsRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

//
//GetStatistics返回
type GetStatisticsResponse struct {
	//
	//关联的主机
	Host *inspection.InspectionHost `protobuf:"bytes,1,opt,name=host,proto3" json:"host" form:"host"`
	//
	//日志
	Log string `protobuf:"bytes,2,opt,name=log,proto3" json:"log" form:"log"`
	//
	//巡检总项数
	TotalNum int32 `protobuf:"varint,3,opt,name=totalNum,proto3" json:"totalNum" form:"totalNum"`
	//
	//通知项数
	NoticeNum int32 `protobuf:"varint,4,opt,name=noticeNum,proto3" json:"noticeNum" form:"noticeNum"`
	//
	//警告项数
	WarningNum int32 `protobuf:"varint,5,opt,name=warningNum,proto3" json:"warningNum" form:"warningNum"`
	//
	//紧急项数
	EmergencyNum int32 `protobuf:"varint,6,opt,name=emergencyNum,proto3" json:"emergencyNum" form:"emergencyNum"`
	//
	//合格项数
	PassedNum int32 `protobuf:"varint,7,opt,name=passedNum,proto3" json:"passedNum" form:"passedNum"`
	//
	//巡检对象实例ID
	InstanceId string `protobuf:"bytes,8,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//评分(精确到小数点后2位)
	Score float32 `protobuf:"fixed32,9,opt,name=score,proto3" json:"score" form:"score"`
	//
	//unexecuted:脚本未执行; failed:返回码大于0; normal:返回码等于0,没有异常指标; abnormal:返回码等于0,有异常指标
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatisticsResponse) Reset()         { *m = GetStatisticsResponse{} }
func (m *GetStatisticsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatisticsResponse) ProtoMessage()    {}
func (*GetStatisticsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81af0e9ed6c23, []int{1}
}
func (m *GetStatisticsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatisticsResponse.Unmarshal(m, b)
}
func (m *GetStatisticsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatisticsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatisticsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatisticsResponse.Merge(m, src)
}
func (m *GetStatisticsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatisticsResponse.Size(m)
}
func (m *GetStatisticsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatisticsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatisticsResponse proto.InternalMessageInfo

func (m *GetStatisticsResponse) GetHost() *inspection.InspectionHost {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *GetStatisticsResponse) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *GetStatisticsResponse) GetTotalNum() int32 {
	if m != nil {
		return m.TotalNum
	}
	return 0
}

func (m *GetStatisticsResponse) GetNoticeNum() int32 {
	if m != nil {
		return m.NoticeNum
	}
	return 0
}

func (m *GetStatisticsResponse) GetWarningNum() int32 {
	if m != nil {
		return m.WarningNum
	}
	return 0
}

func (m *GetStatisticsResponse) GetEmergencyNum() int32 {
	if m != nil {
		return m.EmergencyNum
	}
	return 0
}

func (m *GetStatisticsResponse) GetPassedNum() int32 {
	if m != nil {
		return m.PassedNum
	}
	return 0
}

func (m *GetStatisticsResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetStatisticsResponse) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GetStatisticsResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//GetStatisticsApi返回
type GetStatisticsResponseWrapper struct {
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
	Data                 *GetStatisticsResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetStatisticsResponseWrapper) Reset()         { *m = GetStatisticsResponseWrapper{} }
func (m *GetStatisticsResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetStatisticsResponseWrapper) ProtoMessage()    {}
func (*GetStatisticsResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81af0e9ed6c23, []int{2}
}
func (m *GetStatisticsResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatisticsResponseWrapper.Unmarshal(m, b)
}
func (m *GetStatisticsResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatisticsResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetStatisticsResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatisticsResponseWrapper.Merge(m, src)
}
func (m *GetStatisticsResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetStatisticsResponseWrapper.Size(m)
}
func (m *GetStatisticsResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatisticsResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatisticsResponseWrapper proto.InternalMessageInfo

func (m *GetStatisticsResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetStatisticsResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetStatisticsResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetStatisticsResponseWrapper) GetData() *GetStatisticsResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStatisticsRequest)(nil), "history.GetStatisticsRequest")
	proto.RegisterType((*GetStatisticsResponse)(nil), "history.GetStatisticsResponse")
	proto.RegisterType((*GetStatisticsResponseWrapper)(nil), "history.GetStatisticsResponseWrapper")
}

func init() { proto.RegisterFile("get_statistics.proto", fileDescriptor_e4c81af0e9ed6c23) }

var fileDescriptor_e4c81af0e9ed6c23 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdd, 0x6e, 0xd3, 0x3c,
	0x1c, 0xc6, 0xd5, 0xad, 0xdd, 0x87, 0xbb, 0x4f, 0x6f, 0x7b, 0xdf, 0x6a, 0x02, 0x52, 0x79, 0x13,
	0xea, 0x34, 0x9a, 0xec, 0x43, 0x4c, 0x0c, 0x0e, 0x26, 0x8a, 0x10, 0xf4, 0x04, 0x24, 0x73, 0x80,
	0xc4, 0xbe, 0xe4, 0x26, 0x5e, 0x16, 0x48, 0xf3, 0x0f, 0xb6, 0xcb, 0x18, 0xd3, 0x6e, 0x82, 0xab,
	0xe1, 0x6a, 0x82, 0xc4, 0x21, 0x87, 0xb9, 0x02, 0x64, 0x27, 0x4b, 0xb2, 0x69, 0x47, 0xb5, 0xfd,
	0x3c, 0xbf, 0x7f, 0x1f, 0xa7, 0x4f, 0x8a, 0x96, 0x7d, 0xae, 0x4e, 0xa5, 0x62, 0x2a, 0x90, 0x2a,
	0x70, 0xa5, 0x1d, 0x0b, 0x50, 0x80, 0x27, 0xcf, 0x03, 0xa9, 0x40, 0x5c, 0xae, 0x76, 0xfd, 0x40,
	0x9d, 0x8f, 0x06, 0xb6, 0x0b, 0x43, 0xc7, 0x07, 0x1f, 0x1c, 0xa3, 0x0f, 0x46, 0x67, 0x66, 0x67,
	0x36, 0x66, 0x95, 0x71, 0xab, 0xef, 0x7d, 0xb0, 0x39, 0x93, 0x97, 0x10, 0x4b, 0x3b, 0x04, 0x97,
	0x85, 0x8e, 0x0b, 0x91, 0x12, 0xcc, 0x55, 0x32, 0x23, 0x05, 0x8f, 0xa1, 0x3b, 0x04, 0x8f, 0x87,
	0xd2, 0xc9, 0x8d, 0x8e, 0xd9, 0x3a, 0x41, 0x24, 0x63, 0xee, 0xaa, 0x00, 0x22, 0xe7, 0x1c, 0xa4,
	0xca, 0x07, 0xee, 0x55, 0xbe, 0x7f, 0x78, 0x11, 0xa8, 0x2f, 0x70, 0xe1, 0xf8, 0xd0, 0x35, 0x62,
	0xf7, 0x1b, 0x0b, 0x03, 0x8f, 0x29, 0x10, 0xd2, 0x29, 0x96, 0x19, 0x47, 0x7e, 0xd6, 0xd0, 0xf2,
	0x1b, 0xae, 0x3e, 0x14, 0x17, 0xa3, 0xfc, 0xeb, 0x88, 0x4b, 0x85, 0x29, 0x9a, 0x8a, 0xc3, 0x91,
	0x1f, 0x44, 0x7d, 0xaf, 0x55, 0x6b, 0xd7, 0x3a, 0xd3, 0xbd, 0xbd, 0x34, 0xb1, 0xe6, 0xcf, 0x40,
	0x0c, 0x9f, 0x93, 0x1b, 0x85, 0xfc, 0xf9, 0x6d, 0x59, 0xe8, 0xe1, 0xc9, 0x21, 0xeb, 0xfe, 0x78,
	0xd9, 0xfd, 0x74, 0x7a, 0x7c, 0xb8, 0xd5, 0xdd, 0xbf, 0x59, 0x5f, 0x6d, 0x3d, 0xd9, 0xdd, 0xbe,
	0x5e, 0xa7, 0xc5, 0x1c, 0xfc, 0x18, 0x35, 0x3e, 0xc3, 0xa0, 0xef, 0xb5, 0xc6, 0xcc, 0xc0, 0x85,
	0x34, 0xb1, 0x66, 0xb2, 0x81, 0xe6, 0x98, 0xd0, 0x4c, 0x26, 0xbf, 0xea, 0x68, 0xe5, 0x4e, 0x28,
	0x19, 0x43, 0x24, 0x39, 0x3e, 0x40, 0x75, 0x7d, 0x69, 0x93, 0xa8, 0xb9, 0xb3, 0x6a, 0x97, 0x0f,
	0xc3, 0xee, 0x17, 0xcb, 0xb7, 0x20, 0x55, 0x6f, 0x3e, 0x4d, 0xac, 0x66, 0x36, 0x5c, 0x13, 0x84,
	0x1a, 0x10, 0xb7, 0xd1, 0x78, 0x08, 0x7e, 0x1e, 0x60, 0x2e, 0x4d, 0x2c, 0x94, 0x79, 0x42, 0xf0,
	0x09, 0xd5, 0x12, 0x76, 0xd0, 0x94, 0x02, 0xc5, 0xc2, 0x77, 0xa3, 0x61, 0x6b, 0xbc, 0x5d, 0xeb,
	0x34, 0x7a, 0x4b, 0xe5, 0xc5, 0x6f, 0x14, 0x42, 0x0b, 0x13, 0xde, 0x41, 0xd3, 0x11, 0xa8, 0xc0,
	0xe5, 0x9a, 0xa8, 0x1b, 0x62, 0x39, 0x4d, 0xac, 0x85, 0x8c, 0x28, 0x24, 0x42, 0x4b, 0x1b, 0x7e,
	0x8a, 0xd0, 0x05, 0x13, 0x51, 0x10, 0xf9, 0x1a, 0x6a, 0x18, 0x68, 0x25, 0x4d, 0xac, 0xc5, 0x0c,
	0x2a, 0x35, 0x42, 0x2b, 0x46, 0xfc, 0x02, 0xcd, 0xf0, 0x21, 0x17, 0x3e, 0x8f, 0xdc, 0x4b, 0x0d,
	0x4e, 0x18, 0xf0, 0xff, 0x34, 0xb1, 0x96, 0x32, 0xb0, 0xaa, 0x12, 0x7a, 0xcb, 0xac, 0x73, 0xc6,
	0x4c, 0x4a, 0xee, 0x69, 0x72, 0xf2, 0x6e, 0xce, 0x42, 0x22, 0xb4, 0xb4, 0xe1, 0x3e, 0x42, 0x41,
	0x24, 0x15, 0x8b, 0x5c, 0xde, 0xf7, 0x5a, 0x53, 0xe6, 0xa9, 0x6d, 0x94, 0x39, 0x4b, 0x4d, 0x37,
	0x61, 0x01, 0xcd, 0x9d, 0xe4, 0x05, 0x38, 0xbe, 0xda, 0xde, 0xbd, 0x5e, 0xa7, 0x15, 0x18, 0xef,
	0xa3, 0x86, 0x74, 0x41, 0xf0, 0xd6, 0x74, 0xbb, 0xd6, 0x19, 0xeb, 0xad, 0x95, 0x3f, 0xbe, 0x39,
	0xd6, 0x03, 0xe6, 0xd1, 0xec, 0xc9, 0x91, 0xb7, 0xd9, 0x39, 0xb2, 0x8f, 0xbc, 0xcd, 0x8d, 0x83,
	0x75, 0x9a, 0x11, 0x78, 0x03, 0x4d, 0xe8, 0x37, 0x6f, 0x24, 0x5b, 0xc8, 0x24, 0x58, 0x4c, 0x13,
	0x6b, 0x36, 0x67, 0xcd, 0x39, 0xa1, 0xb9, 0x81, 0xfc, 0xad, 0xa1, 0x07, 0xf7, 0x56, 0xe7, 0xa3,
	0x60, 0x71, 0xcc, 0x05, 0x5e, 0x43, 0x75, 0x17, 0x3c, 0x6e, 0x1a, 0xd4, 0xa8, 0xb6, 0x44, 0x9f,
	0x12, 0x6a, 0x44, 0xfc, 0x0c, 0x35, 0xf5, 0xe7, 0xeb, 0xef, 0x71, 0xc8, 0x82, 0x28, 0x6f, 0xcb,
	0x7f, 0x69, 0x62, 0xe1, 0xd2, 0x9b, 0x8b, 0x84, 0x56, 0xad, 0xba, 0xe2, 0x5c, 0x08, 0x10, 0xa6,
	0x3a, 0xb7, 0x2a, 0x6e, 0x8e, 0x09, 0xcd, 0x64, 0xfc, 0x0a, 0xd5, 0x3d, 0xa6, 0x98, 0xe9, 0x4b,
	0x73, 0xe7, 0x91, 0x9d, 0xff, 0x8f, 0xd8, 0xf7, 0x66, 0xaf, 0xc6, 0xd4, 0x14, 0xa1, 0x06, 0x1e,
	0x4c, 0x98, 0x77, 0x78, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xf4, 0xae, 0x20, 0x9c,
	0x04, 0x00, 0x00,
}
