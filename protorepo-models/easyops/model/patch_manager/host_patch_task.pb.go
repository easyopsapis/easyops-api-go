// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: host_patch_task.proto

package patch_manager

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//主机的补丁安装任务
type HostPatchTask struct {
	//
	//任务ID，easy_command的ID
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	//
	//汇总任务ID
	SummaryTaskId string `protobuf:"bytes,2,opt,name=summaryTaskId,proto3" json:"summaryTaskId" form:"summaryTaskId"`
	//
	//主机ID
	HostId string `protobuf:"bytes,3,opt,name=hostId,proto3" json:"hostId" form:"hostId"`
	//
	//主机IP
	HostIp string `protobuf:"bytes,4,opt,name=hostIp,proto3" json:"hostIp" form:"hostIp"`
	//
	//补丁ID列表
	PatchIdList []string `protobuf:"bytes,5,rep,name=patchIdList,proto3" json:"patchIdList" form:"patchIdList"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,6,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//结束时间
	Etime int32 `protobuf:"varint,7,opt,name=etime,proto3" json:"etime" form:"etime"`
	//
	//任务状态
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status" form:"status"`
	//
	//任务结果
	Result               []*HostPatchTask_Result `protobuf:"bytes,9,rep,name=result,proto3" json:"result" form:"result"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HostPatchTask) Reset()         { *m = HostPatchTask{} }
func (m *HostPatchTask) String() string { return proto.CompactTextString(m) }
func (*HostPatchTask) ProtoMessage()    {}
func (*HostPatchTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_924515f10ddc9eef, []int{0}
}
func (m *HostPatchTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostPatchTask.Unmarshal(m, b)
}
func (m *HostPatchTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostPatchTask.Marshal(b, m, deterministic)
}
func (m *HostPatchTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostPatchTask.Merge(m, src)
}
func (m *HostPatchTask) XXX_Size() int {
	return xxx_messageInfo_HostPatchTask.Size(m)
}
func (m *HostPatchTask) XXX_DiscardUnknown() {
	xxx_messageInfo_HostPatchTask.DiscardUnknown(m)
}

var xxx_messageInfo_HostPatchTask proto.InternalMessageInfo

func (m *HostPatchTask) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *HostPatchTask) GetSummaryTaskId() string {
	if m != nil {
		return m.SummaryTaskId
	}
	return ""
}

func (m *HostPatchTask) GetHostId() string {
	if m != nil {
		return m.HostId
	}
	return ""
}

func (m *HostPatchTask) GetHostIp() string {
	if m != nil {
		return m.HostIp
	}
	return ""
}

func (m *HostPatchTask) GetPatchIdList() []string {
	if m != nil {
		return m.PatchIdList
	}
	return nil
}

func (m *HostPatchTask) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *HostPatchTask) GetEtime() int32 {
	if m != nil {
		return m.Etime
	}
	return 0
}

func (m *HostPatchTask) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *HostPatchTask) GetResult() []*HostPatchTask_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type HostPatchTask_Result struct {
	//
	//补丁ID
	PatchId string `protobuf:"bytes,1,opt,name=patchId,proto3" json:"patchId" form:"patchId"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,2,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//结束时间
	Etime int32 `protobuf:"varint,3,opt,name=etime,proto3" json:"etime" form:"etime"`
	//
	//补丁状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" form:"status"`
	//
	//补丁安装返回码Windows Update Exit Code
	Wuec int32 `protobuf:"varint,5,opt,name=wuec,proto3" json:"wuec" form:"wuec"`
	//
	//补丁安装日志
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message" form:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostPatchTask_Result) Reset()         { *m = HostPatchTask_Result{} }
func (m *HostPatchTask_Result) String() string { return proto.CompactTextString(m) }
func (*HostPatchTask_Result) ProtoMessage()    {}
func (*HostPatchTask_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_924515f10ddc9eef, []int{0, 0}
}
func (m *HostPatchTask_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostPatchTask_Result.Unmarshal(m, b)
}
func (m *HostPatchTask_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostPatchTask_Result.Marshal(b, m, deterministic)
}
func (m *HostPatchTask_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostPatchTask_Result.Merge(m, src)
}
func (m *HostPatchTask_Result) XXX_Size() int {
	return xxx_messageInfo_HostPatchTask_Result.Size(m)
}
func (m *HostPatchTask_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_HostPatchTask_Result.DiscardUnknown(m)
}

var xxx_messageInfo_HostPatchTask_Result proto.InternalMessageInfo

func (m *HostPatchTask_Result) GetPatchId() string {
	if m != nil {
		return m.PatchId
	}
	return ""
}

func (m *HostPatchTask_Result) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *HostPatchTask_Result) GetEtime() int32 {
	if m != nil {
		return m.Etime
	}
	return 0
}

func (m *HostPatchTask_Result) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *HostPatchTask_Result) GetWuec() int32 {
	if m != nil {
		return m.Wuec
	}
	return 0
}

func (m *HostPatchTask_Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HostPatchTask)(nil), "patch_manager.HostPatchTask")
	proto.RegisterType((*HostPatchTask_Result)(nil), "patch_manager.HostPatchTask.Result")
}

func init() { proto.RegisterFile("host_patch_task.proto", fileDescriptor_924515f10ddc9eef) }

var fileDescriptor_924515f10ddc9eef = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x05, 0x21, 0xe1, 0xc6, 0xdc, 0xf0, 0xc7, 0xf7, 0x56, 0xb2, 0xd8, 0x24, 0x32, 0x50,
	0x8d, 0x69, 0x6c, 0xc7, 0x4e, 0x02, 0x22, 0x52, 0x85, 0xc8, 0x82, 0x36, 0x12, 0xaa, 0x2a, 0x8b,
	0x6e, 0x9a, 0x04, 0x34, 0xc4, 0x4e, 0x88, 0x88, 0x99, 0xc8, 0x33, 0x29, 0xa2, 0xb6, 0x77, 0xdd,
	0x57, 0xea, 0x53, 0xf4, 0x01, 0xfa, 0x02, 0x7d, 0x91, 0x54, 0xea, 0x23, 0xe4, 0x09, 0xaa, 0x99,
	0x31, 0x25, 0x06, 0x5c, 0x58, 0x50, 0x56, 0x99, 0x39, 0xe7, 0xfb, 0xce, 0xfc, 0xe6, 0x58, 0x9a,
	0x13, 0xe1, 0xd9, 0x29, 0xc2, 0xe4, 0x78, 0x08, 0x49, 0xe7, 0xf4, 0x98, 0x40, 0x7c, 0xa6, 0x0d,
	0x3d, 0x44, 0x90, 0x98, 0xe3, 0x11, 0x17, 0x9e, 0xc3, 0x9e, 0xe3, 0xad, 0xaa, 0xbd, 0x3e, 0x39,
	0x1d, 0x9d, 0x68, 0x1d, 0xe4, 0xea, 0x3d, 0xd4, 0x43, 0x3a, 0x53, 0x9d, 0x8c, 0xba, 0x6c, 0xc7,
	0x36, 0x6c, 0xc5, 0xdd, 0xab, 0x5b, 0x53, 0x72, 0xf7, 0xa2, 0x4f, 0xce, 0xd0, 0x85, 0xde, 0x43,
	0x2a, 0x4b, 0xaa, 0x1f, 0xe0, 0xa0, 0x6f, 0x43, 0x82, 0x3c, 0xac, 0xff, 0x5e, 0x72, 0x9f, 0xfc,
	0x6d, 0x49, 0xc8, 0xbd, 0x46, 0x98, 0xbc, 0xa5, 0x87, 0x1f, 0x42, 0x7c, 0x26, 0x2a, 0x42, 0x86,
	0x52, 0x35, 0x6c, 0x69, 0xa6, 0x30, 0x03, 0xb2, 0xf5, 0x95, 0xc9, 0x38, 0x9f, 0xeb, 0x22, 0xcf,
	0xad, 0xc9, 0x3c, 0x2e, 0x5b, 0x91, 0x40, 0x7c, 0x27, 0xe4, 0xf0, 0xc8, 0x75, 0xa1, 0x77, 0x79,
	0xc8, 0x1d, 0xb3, 0xcc, 0xa1, 0x4f, 0xc6, 0xf9, 0xff, 0xb9, 0x23, 0x96, 0x96, 0x7f, 0xfe, 0xc8,
	0xff, 0x27, 0xac, 0x1c, 0x35, 0x4b, 0xea, 0x0e, 0x54, 0xbb, 0x7b, 0xea, 0x7e, 0xdb, 0x37, 0x2b,
	0xe1, 0xba, 0x15, 0xaf, 0x22, 0xbe, 0x14, 0x32, 0xb4, 0x45, 0x0d, 0x5b, 0x4a, 0xb1, 0x7a, 0x1b,
	0xd7, 0x04, 0x3c, 0x4e, 0x0b, 0x2d, 0x0b, 0x8b, 0x51, 0xa1, 0x8f, 0x6d, 0xdf, 0x28, 0x87, 0xeb,
	0x56, 0x64, 0x12, 0xbf, 0x67, 0x23, 0xff, 0x50, 0x9a, 0x63, 0xfe, 0xaf, 0xd9, 0x1b, 0x05, 0x86,
	0xb4, 0xc0, 0x97, 0xac, 0xf0, 0x39, 0x7b, 0x04, 0x80, 0x09, 0xaa, 0xcd, 0x92, 0x5a, 0x6d, 0xfb,
	0x46, 0x18, 0x34, 0x4b, 0x6a, 0xa5, 0xdd, 0xb2, 0x7d, 0x23, 0x54, 0xe8, 0xda, 0x68, 0xef, 0xd2,
	0x4d, 0xd1, 0x0c, 0x15, 0xd0, 0xd2, 0x1e, 0xa8, 0x54, 0xfc, 0x72, 0xa8, 0x04, 0x2d, 0xbc, 0x09,
	0x00, 0xa0, 0x80, 0x7b, 0xea, 0x3e, 0x54, 0xbb, 0x6d, 0xdf, 0x28, 0x56, 0xc2, 0x9a, 0xe2, 0x6f,
	0x87, 0xb7, 0xa2, 0x41, 0x4d, 0x51, 0x82, 0x3b, 0xc5, 0x5b, 0x21, 0xa8, 0xdd, 0x52, 0x03, 0x60,
	0x72, 0x8e, 0xc0, 0x8c, 0x28, 0x02, 0xa3, 0x65, 0xb7, 0xec, 0xa0, 0x69, 0xa8, 0x3b, 0x94, 0x83,
	0xc3, 0xde, 0xa3, 0xe1, 0x98, 0x89, 0x27, 0x57, 0x43, 0x00, 0x6e, 0x9f, 0xad, 0xf0, 0x2b, 0x06,
	0xb5, 0x27, 0x61, 0xa8, 0x24, 0x32, 0x50, 0xdb, 0x5d, 0xa9, 0xdd, 0xc7, 0x04, 0xfb, 0x03, 0x59,
	0x39, 0x91, 0xac, 0x92, 0x40, 0xe6, 0x97, 0x8a, 0x66, 0xf8, 0x44, 0x74, 0x66, 0x22, 0x5d, 0x35,
	0x99, 0xae, 0xfc, 0x54, 0x74, 0x46, 0x22, 0xdd, 0x56, 0x32, 0x5d, 0xe5, 0x6f, 0xd0, 0xd5, 0x92,
	0x40, 0xb6, 0x93, 0x41, 0xaa, 0x8f, 0x0f, 0xa2, 0x80, 0x0d, 0xed, 0x85, 0xb2, 0xdb, 0xc2, 0x9b,
	0x57, 0x6f, 0xd8, 0x50, 0x3c, 0x10, 0x16, 0xd8, 0x38, 0x68, 0xd8, 0x07, 0x7d, 0x4c, 0xa4, 0x74,
	0x21, 0x05, 0xb2, 0xf5, 0xcd, 0xc9, 0x38, 0x2f, 0xf2, 0x67, 0x6c, 0x2a, 0x79, 0xf7, 0x63, 0x38,
	0x6d, 0x17, 0x9f, 0x0b, 0xe9, 0x0e, 0xe9, 0xbb, 0x8e, 0x94, 0x29, 0xcc, 0x80, 0x74, 0x7d, 0x79,
	0x32, 0xce, 0xff, 0xcb, 0xeb, 0xb0, 0xb0, 0x6c, 0xf1, 0x34, 0xd5, 0x39, 0x4c, 0x37, 0x7f, 0x53,
	0xe7, 0x44, 0x3a, 0xf6, 0x4b, 0x47, 0x04, 0x26, 0x90, 0x8c, 0xb0, 0xf4, 0xcf, 0xcd, 0x11, 0xc1,
	0xe3, 0xb2, 0x15, 0x09, 0xc4, 0x37, 0x42, 0xc6, 0x73, 0xf0, 0x68, 0x40, 0xa4, 0x6c, 0x21, 0x05,
	0x16, 0xcc, 0x35, 0x2d, 0x36, 0xe6, 0xb4, 0xd8, 0xec, 0xd1, 0x2c, 0x26, 0x9d, 0xae, 0xc7, 0xcd,
	0xb2, 0x15, 0x55, 0x59, 0xfd, 0x34, 0x2b, 0x64, 0xb8, 0x4a, 0x2c, 0x0a, 0xf3, 0xd1, 0x25, 0xa3,
	0x49, 0x25, 0x4e, 0xc6, 0xf9, 0xc5, 0x58, 0x7f, 0x64, 0xeb, 0x4a, 0x72, 0xdd, 0x83, 0xd9, 0x07,
	0xf6, 0x20, 0xf5, 0xd0, 0x1e, 0xcc, 0xdd, 0xd7, 0x83, 0x35, 0x61, 0xee, 0x62, 0xe4, 0x74, 0xa4,
	0x34, 0xab, 0xb8, 0x34, 0x19, 0xe7, 0x17, 0xb8, 0x90, 0x46, 0x65, 0x8b, 0x25, 0xe9, 0x6d, 0x5c,
	0x07, 0x63, 0xd8, 0xe3, 0x5f, 0x29, 0x76, 0x9b, 0x28, 0x21, 0x5b, 0x57, 0x92, 0x7a, 0xe3, 0xfd,
	0xab, 0x1e, 0xd2, 0x1c, 0x88, 0x2f, 0xd1, 0x10, 0x6b, 0x03, 0xd4, 0x81, 0x03, 0xbd, 0x83, 0xce,
	0x89, 0x07, 0x3b, 0x04, 0xf3, 0xbf, 0x0a, 0x9e, 0x33, 0x44, 0xaa, 0x8b, 0x6c, 0x67, 0x80, 0xf5,
	0x48, 0xa8, 0xb3, 0xad, 0x1e, 0xfb, 0x02, 0x27, 0x19, 0xa6, 0x2e, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0x3b, 0x9d, 0xf9, 0x97, 0x08, 0x00, 0x00,
}
