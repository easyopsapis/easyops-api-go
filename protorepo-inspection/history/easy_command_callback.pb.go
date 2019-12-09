// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: easy_command_callback.proto

package history

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	easy_command "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_command"
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
//EasyCommandCallback请求
type EasyCommandCallbackRequest struct {
	//
	//套件id
	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId" form:"pluginId"`
	//
	//巡检任务ID
	InspectionTaskId string `protobuf:"bytes,2,opt,name=inspectionTaskId,proto3" json:"inspectionTaskId" form:"inspectionTaskId"`
	//
	//巡检作业ID
	JobId string `protobuf:"bytes,3,opt,name=jobId,proto3" json:"jobId" form:"jobId"`
	//
	//任务ID
	TaskId string `protobuf:"bytes,4,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//任务执行状态
	Status string `protobuf:"bytes,5,opt,name=status,proto3" json:"status" form:"status"`
	//
	//整个任务当前的用时，单位为秒
	UsedTime int32 `protobuf:"varint,6,opt,name=usedTime,proto3" json:"usedTime" form:"usedTime"`
	//
	//任务开始时间
	StartTime string `protobuf:"bytes,7,opt,name=startTime,proto3" json:"startTime" form:"startTime"`
	//
	//任务状态更新时间
	UpdateTime string `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//任务结束时间
	EndTime string `protobuf:"bytes,9,opt,name=endTime,proto3" json:"endTime" form:"endTime"`
	//
	//任务的各个目标机器的执行结果详情
	TargetsLog           []*easy_command.TargetLog `protobuf:"bytes,10,rep,name=targetsLog,proto3" json:"targetsLog" form:"targetsLog"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EasyCommandCallbackRequest) Reset()         { *m = EasyCommandCallbackRequest{} }
func (m *EasyCommandCallbackRequest) String() string { return proto.CompactTextString(m) }
func (*EasyCommandCallbackRequest) ProtoMessage()    {}
func (*EasyCommandCallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab6d61bbd45a925, []int{0}
}
func (m *EasyCommandCallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EasyCommandCallbackRequest.Unmarshal(m, b)
}
func (m *EasyCommandCallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EasyCommandCallbackRequest.Marshal(b, m, deterministic)
}
func (m *EasyCommandCallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EasyCommandCallbackRequest.Merge(m, src)
}
func (m *EasyCommandCallbackRequest) XXX_Size() int {
	return xxx_messageInfo_EasyCommandCallbackRequest.Size(m)
}
func (m *EasyCommandCallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EasyCommandCallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EasyCommandCallbackRequest proto.InternalMessageInfo

func (m *EasyCommandCallbackRequest) GetPluginId() string {
	if m != nil {
		return m.PluginId
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetInspectionTaskId() string {
	if m != nil {
		return m.InspectionTaskId
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetUsedTime() int32 {
	if m != nil {
		return m.UsedTime
	}
	return 0
}

func (m *EasyCommandCallbackRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *EasyCommandCallbackRequest) GetTargetsLog() []*easy_command.TargetLog {
	if m != nil {
		return m.TargetsLog
	}
	return nil
}

//
//EasyCommandCallbackApi返回
type EasyCommandCallbackResponseWrapper struct {
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
	Data                 *types.Empty `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EasyCommandCallbackResponseWrapper) Reset()         { *m = EasyCommandCallbackResponseWrapper{} }
func (m *EasyCommandCallbackResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*EasyCommandCallbackResponseWrapper) ProtoMessage()    {}
func (*EasyCommandCallbackResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab6d61bbd45a925, []int{1}
}
func (m *EasyCommandCallbackResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EasyCommandCallbackResponseWrapper.Unmarshal(m, b)
}
func (m *EasyCommandCallbackResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EasyCommandCallbackResponseWrapper.Marshal(b, m, deterministic)
}
func (m *EasyCommandCallbackResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EasyCommandCallbackResponseWrapper.Merge(m, src)
}
func (m *EasyCommandCallbackResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_EasyCommandCallbackResponseWrapper.Size(m)
}
func (m *EasyCommandCallbackResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_EasyCommandCallbackResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_EasyCommandCallbackResponseWrapper proto.InternalMessageInfo

func (m *EasyCommandCallbackResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EasyCommandCallbackResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *EasyCommandCallbackResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *EasyCommandCallbackResponseWrapper) GetData() *types.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*EasyCommandCallbackRequest)(nil), "history.EasyCommandCallbackRequest")
	proto.RegisterType((*EasyCommandCallbackResponseWrapper)(nil), "history.EasyCommandCallbackResponseWrapper")
}

func init() { proto.RegisterFile("easy_command_callback.proto", fileDescriptor_eab6d61bbd45a925) }

var fileDescriptor_eab6d61bbd45a925 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xdd, 0x4e, 0xdb, 0x48,
	0x18, 0x95, 0x17, 0x02, 0x61, 0xb2, 0xbb, 0x80, 0x57, 0x0b, 0x11, 0xa8, 0x72, 0xe4, 0xa2, 0xca,
	0x29, 0x1d, 0x3b, 0x3f, 0x2d, 0x2a, 0xe9, 0x45, 0xd4, 0x20, 0x54, 0x21, 0xa1, 0x5e, 0x58, 0x91,
	0x90, 0x70, 0x43, 0x34, 0xb1, 0x07, 0xe3, 0x62, 0x7b, 0x5c, 0xcf, 0xa4, 0x34, 0x25, 0xbc, 0x50,
	0xa5, 0xbe, 0x45, 0xd5, 0xbe, 0x45, 0x90, 0x2a, 0x55, 0xe2, 0xa6, 0x37, 0x79, 0x82, 0xca, 0x33,
	0x4e, 0x62, 0x01, 0x2f, 0xc0, 0x95, 0xe7, 0xfb, 0xce, 0x39, 0xdf, 0xcc, 0xf9, 0x12, 0x1d, 0xb0,
	0x89, 0x11, 0x1d, 0x74, 0x6d, 0x12, 0x04, 0x28, 0x74, 0xba, 0x36, 0xf2, 0xfd, 0x1e, 0xb2, 0xcf,
	0xf5, 0x28, 0x26, 0x8c, 0xc8, 0x8b, 0x67, 0x1e, 0x65, 0x24, 0x1e, 0x6c, 0x40, 0xd7, 0x63, 0x67,
	0xfd, 0x9e, 0x6e, 0x93, 0xc0, 0x70, 0x89, 0x4b, 0x0c, 0x8e, 0xf7, 0xfa, 0xa7, 0xbc, 0xe2, 0x05,
	0x3f, 0x09, 0xdd, 0xc6, 0x91, 0x4b, 0xf4, 0x64, 0x2e, 0x89, 0xa8, 0xee, 0x13, 0x1b, 0xf9, 0x86,
	0x4d, 0x42, 0x16, 0x23, 0x9b, 0x51, 0xa1, 0x8c, 0x71, 0x44, 0x60, 0x40, 0x1c, 0xec, 0x53, 0x23,
	0x25, 0x1a, 0xbc, 0x34, 0xb2, 0xcf, 0x31, 0x18, 0x8a, 0x5d, 0xcc, 0xba, 0x3e, 0x71, 0xd3, 0xc1,
	0x3b, 0x99, 0x77, 0x04, 0x17, 0x1e, 0x3b, 0x27, 0x17, 0x86, 0x4b, 0x20, 0x07, 0xe1, 0x47, 0xe4,
	0x7b, 0x0e, 0x62, 0x24, 0xa6, 0xc6, 0xf4, 0x98, 0xea, 0x36, 0x5d, 0x42, 0x5c, 0x1f, 0xcf, 0x9e,
	0x8d, 0x83, 0x88, 0x0d, 0x04, 0xa8, 0xfe, 0xce, 0x83, 0x8d, 0x7d, 0x44, 0x07, 0x7b, 0xe2, 0xd6,
	0xbd, 0x74, 0x07, 0x26, 0xfe, 0xd0, 0xc7, 0x94, 0xc9, 0x26, 0xc8, 0x47, 0x7e, 0xdf, 0xf5, 0xc2,
	0x03, 0xa7, 0x28, 0x95, 0x24, 0x6d, 0xa9, 0xb5, 0x33, 0x1e, 0x29, 0xcb, 0xa7, 0x24, 0x0e, 0x1a,
	0xea, 0x04, 0x51, 0x7f, 0x5e, 0x2b, 0x0a, 0x78, 0x74, 0x62, 0x21, 0xf8, 0xf9, 0x35, 0x3c, 0xee,
	0x76, 0xac, 0x0a, 0xdc, 0x9d, 0x9c, 0x2f, 0x2b, 0xcf, 0xea, 0xd5, 0xab, 0x2d, 0x73, 0x3a, 0x47,
	0x7e, 0x03, 0x56, 0xbc, 0x90, 0x46, 0xd8, 0x66, 0x1e, 0x09, 0xdb, 0x88, 0x9e, 0x1f, 0x38, 0xc5,
	0xbf, 0xf8, 0xec, 0xcd, 0xf1, 0x48, 0x59, 0x17, 0xb3, 0x6f, 0x33, 0x54, 0xf3, 0x8e, 0x48, 0x7e,
	0x02, 0x72, 0xef, 0x49, 0xef, 0xc0, 0x29, 0xce, 0x71, 0xf5, 0xca, 0x78, 0xa4, 0xfc, 0x2d, 0xd4,
	0xbc, 0xad, 0x9a, 0x02, 0x96, 0xcb, 0x60, 0x81, 0x89, 0x6b, 0xe6, 0x39, 0x71, 0x75, 0x3c, 0x52,
	0xfe, 0x11, 0x44, 0x96, 0x0e, 0x4f, 0x09, 0x09, 0x95, 0x32, 0xc4, 0xfa, 0xb4, 0x98, 0xbb, 0x4d,
	0x15, 0x7d, 0xd5, 0x4c, 0x09, 0xb2, 0x01, 0xf2, 0x7d, 0x8a, 0x9d, 0xb6, 0x17, 0xe0, 0xe2, 0x42,
	0x49, 0xd2, 0x72, 0xad, 0xff, 0x66, 0xab, 0x99, 0x20, 0xaa, 0x39, 0x25, 0xc9, 0xbf, 0x24, 0xb0,
	0x44, 0x19, 0x8a, 0x19, 0x97, 0x2c, 0xf2, 0xf9, 0xdf, 0xa5, 0xf1, 0x48, 0x59, 0x99, 0x5e, 0x20,
	0xb0, 0x64, 0x9f, 0x5f, 0x25, 0xf0, 0x45, 0x3a, 0xd1, 0xb4, 0x66, 0xc3, 0xaa, 0xc2, 0x5d, 0xbe,
	0xd1, 0xce, 0xd3, 0x72, 0x93, 0x7f, 0x2f, 0x9f, 0x5f, 0x95, 0xa1, 0x56, 0xb5, 0x2a, 0xb0, 0xd6,
	0x19, 0x56, 0x38, 0x5e, 0x86, 0x5a, 0xdd, 0xaa, 0xc0, 0xea, 0xa4, 0x1e, 0x5a, 0x55, 0x58, 0x13,
	0xaa, 0xb2, 0xd5, 0x2e, 0x75, 0xb4, 0x9a, 0x55, 0x81, 0xf5, 0xce, 0x90, 0x73, 0x44, 0xbb, 0xa1,
	0x59, 0x15, 0xf8, 0x62, 0x52, 0xcc, 0xce, 0xda, 0x3b, 0x9d, 0x7f, 0xb7, 0xcb, 0x4d, 0xed, 0x78,
	0x68, 0x6d, 0xc3, 0x8e, 0xd6, 0x6c, 0xdc, 0x23, 0xcf, 0xa8, 0x9b, 0x5b, 0xe6, 0xcc, 0x99, 0x7c,
	0x23, 0x01, 0xd0, 0x8f, 0x1c, 0xc4, 0x30, 0x37, 0x9a, 0xe7, 0x46, 0x7f, 0x24, 0x46, 0x57, 0xd3,
	0xe5, 0x4c, 0xc1, 0x87, 0xe8, 0x34, 0xe3, 0x4d, 0xbe, 0x96, 0xc0, 0x22, 0x0e, 0xc5, 0x7f, 0x60,
	0x89, 0xfb, 0xfc, 0x96, 0xf8, 0xfc, 0x57, 0xf8, 0x4c, 0x91, 0x87, 0x68, 0x72, 0xe2, 0x4a, 0x7e,
	0x0b, 0x80, 0x08, 0x22, 0x7a, 0x48, 0xdc, 0x22, 0x28, 0xcd, 0x69, 0x85, 0xda, 0xba, 0x9e, 0x0d,
	0x2a, 0xbd, 0xcd, 0xf1, 0x43, 0xe2, 0xb6, 0xfe, 0x9f, 0xfd, 0xc6, 0x33, 0x91, 0x6a, 0x66, 0x26,
	0xa8, 0x37, 0x12, 0x50, 0xef, 0xcd, 0x1b, 0x1a, 0x91, 0x90, 0xe2, 0xa3, 0x18, 0x45, 0x11, 0x8e,
	0xe5, 0xc7, 0x60, 0xde, 0x26, 0x0e, 0xe6, 0x99, 0x93, 0x6b, 0x2d, 0x8f, 0x47, 0x4a, 0x41, 0xcc,
	0x4d, 0xba, 0xaa, 0xc9, 0x41, 0xf9, 0x25, 0x28, 0x24, 0xdf, 0xfd, 0x4f, 0x91, 0x8f, 0xbc, 0x30,
	0xcd, 0x90, 0xb5, 0xf1, 0x48, 0x91, 0x67, 0xdc, 0x14, 0x54, 0xcd, 0x2c, 0x35, 0x49, 0x0e, 0x1c,
	0xc7, 0x24, 0xbe, 0x9b, 0x1c, 0xbc, 0xad, 0x9a, 0x02, 0x96, 0x5f, 0x81, 0x79, 0x07, 0x31, 0xc4,
	0x73, 0xa3, 0x50, 0x5b, 0xd3, 0x45, 0x92, 0xea, 0x93, 0x24, 0xd5, 0xf7, 0x93, 0x24, 0xcd, 0x3e,
	0x2f, 0x61, 0xab, 0x26, 0x17, 0xf5, 0x16, 0x38, 0xad, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0x00, 0x4b, 0x3e, 0x66, 0x06, 0x00, 0x00,
}