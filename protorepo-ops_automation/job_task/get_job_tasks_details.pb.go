// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_job_tasks_details.proto

package job_task

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	ops_automation "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/ops_automation"
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
//GetJobTasksDetail请求
type GetJobTasksDetailRequest struct {
	//
	//job Task Id
	JobTaskId            string   `protobuf:"bytes,1,opt,name=jobTaskId,proto3" json:"jobTaskId" form:"jobTaskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobTasksDetailRequest) Reset()         { *m = GetJobTasksDetailRequest{} }
func (m *GetJobTasksDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobTasksDetailRequest) ProtoMessage()    {}
func (*GetJobTasksDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c377135226805df8, []int{0}
}
func (m *GetJobTasksDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobTasksDetailRequest.Unmarshal(m, b)
}
func (m *GetJobTasksDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobTasksDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetJobTasksDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobTasksDetailRequest.Merge(m, src)
}
func (m *GetJobTasksDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobTasksDetailRequest.Size(m)
}
func (m *GetJobTasksDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobTasksDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobTasksDetailRequest proto.InternalMessageInfo

func (m *GetJobTasksDetailRequest) GetJobTaskId() string {
	if m != nil {
		return m.JobTaskId
	}
	return ""
}

//
//GetJobTasksDetail返回
type GetJobTasksDetailResponse struct {
	//
	//任务名称
	JobName string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName" form:"jobName"`
	//
	//菜单名称
	MenuName string `protobuf:"bytes,2,opt,name=menuName,proto3" json:"menuName" form:"menuName"`
	//
	//easy command id
	ExecId string `protobuf:"bytes,3,opt,name=execId,proto3" json:"execId" form:"execId"`
	//
	//工具类型
	ResourceType string `protobuf:"bytes,4,opt,name=resourceType,proto3" json:"resourceType" form:"resourceType"`
	//
	//easy command id
	ResourceId string `protobuf:"bytes,5,opt,name=resourceId,proto3" json:"resourceId" form:"resourceId"`
	//
	//工具版本ID
	ResourceVId string `protobuf:"bytes,6,opt,name=resourceVId,proto3" json:"resourceVId" form:"resourceVId"`
	//
	//工具版本名称
	ResourceVName string `protobuf:"bytes,7,opt,name=resourceVName,proto3" json:"resourceVName" form:"resourceVName"`
	//
	//执行主机列表
	Hosts []string `protobuf:"bytes,8,rep,name=hosts,proto3" json:"hosts" form:"hosts"`
	//
	//通知邮件
	Mail *ops_automation.MailInfo `protobuf:"bytes,9,opt,name=mail,proto3" json:"mail" form:"mail"`
	//
	//作业创建时间 2019-07-02 11:38:59.09199400
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime" form:"createTime"`
	//
	//作业更新时间 2019-07-02 11:39:08.68835900
	UpdateTime string `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//作业创建者
	Creator string `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//org
	Org int32 `protobuf:"varint,13,opt,name=org,proto3" json:"org" form:"org"`
	//
	//job task id
	Id string `protobuf:"bytes,14,opt,name=id,proto3" json:"id" form:"id"`
	//
	//job id
	JobId string `protobuf:"bytes,15,opt,name=jobId,proto3" json:"jobId" form:"jobId"`
	//
	//执行触发器
	Trigger string `protobuf:"bytes,16,opt,name=trigger,proto3" json:"trigger" form:"trigger"`
	//
	//工具执行用户
	ExecUser string `protobuf:"bytes,17,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//工具结果状态
	Status string `protobuf:"bytes,18,opt,name=status,proto3" json:"status" form:"status"`
	//
	//作业成功率。success agents / all agents
	SuccessRate float32 `protobuf:"fixed32,19,opt,name=successRate,proto3" json:"successRate" form:"successRate"`
	//
	//工具报错信息
	Error                string   `protobuf:"bytes,20,opt,name=error,proto3" json:"error" form:"error"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobTasksDetailResponse) Reset()         { *m = GetJobTasksDetailResponse{} }
func (m *GetJobTasksDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetJobTasksDetailResponse) ProtoMessage()    {}
func (*GetJobTasksDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c377135226805df8, []int{1}
}
func (m *GetJobTasksDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobTasksDetailResponse.Unmarshal(m, b)
}
func (m *GetJobTasksDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobTasksDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetJobTasksDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobTasksDetailResponse.Merge(m, src)
}
func (m *GetJobTasksDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetJobTasksDetailResponse.Size(m)
}
func (m *GetJobTasksDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobTasksDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobTasksDetailResponse proto.InternalMessageInfo

func (m *GetJobTasksDetailResponse) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetMenuName() string {
	if m != nil {
		return m.MenuName
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetExecId() string {
	if m != nil {
		return m.ExecId
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetResourceVId() string {
	if m != nil {
		return m.ResourceVId
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetResourceVName() string {
	if m != nil {
		return m.ResourceVName
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *GetJobTasksDetailResponse) GetMail() *ops_automation.MailInfo {
	if m != nil {
		return m.Mail
	}
	return nil
}

func (m *GetJobTasksDetailResponse) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *GetJobTasksDetailResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetJobTasksDetailResponse) GetSuccessRate() float32 {
	if m != nil {
		return m.SuccessRate
	}
	return 0
}

func (m *GetJobTasksDetailResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

//
//GetJobTasksDetailApi返回
type GetJobTasksDetailResponseWrapper struct {
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
	Data                 *GetJobTasksDetailResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetJobTasksDetailResponseWrapper) Reset()         { *m = GetJobTasksDetailResponseWrapper{} }
func (m *GetJobTasksDetailResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetJobTasksDetailResponseWrapper) ProtoMessage()    {}
func (*GetJobTasksDetailResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_c377135226805df8, []int{2}
}
func (m *GetJobTasksDetailResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobTasksDetailResponseWrapper.Unmarshal(m, b)
}
func (m *GetJobTasksDetailResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobTasksDetailResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetJobTasksDetailResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobTasksDetailResponseWrapper.Merge(m, src)
}
func (m *GetJobTasksDetailResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetJobTasksDetailResponseWrapper.Size(m)
}
func (m *GetJobTasksDetailResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobTasksDetailResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobTasksDetailResponseWrapper proto.InternalMessageInfo

func (m *GetJobTasksDetailResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetJobTasksDetailResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetJobTasksDetailResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetJobTasksDetailResponseWrapper) GetData() *GetJobTasksDetailResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetJobTasksDetailRequest)(nil), "job_task.GetJobTasksDetailRequest")
	proto.RegisterType((*GetJobTasksDetailResponse)(nil), "job_task.GetJobTasksDetailResponse")
	proto.RegisterType((*GetJobTasksDetailResponseWrapper)(nil), "job_task.GetJobTasksDetailResponseWrapper")
}

func init() { proto.RegisterFile("get_job_tasks_details.proto", fileDescriptor_c377135226805df8) }

var fileDescriptor_c377135226805df8 = []byte{
	// 1049 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x6f, 0x6e, 0x23, 0x35,
	0x18, 0xc6, 0x9b, 0xb4, 0x69, 0x1b, 0xb7, 0xdd, 0xb6, 0xde, 0xb2, 0x0c, 0x0b, 0x68, 0xc2, 0xec,
	0x52, 0xcd, 0x2c, 0x99, 0xa4, 0x49, 0x93, 0xee, 0x6e, 0x10, 0x54, 0x5b, 0x41, 0x21, 0x48, 0xf0,
	0xc1, 0x2a, 0x05, 0xd1, 0x74, 0x23, 0x27, 0xe3, 0x66, 0xa7, 0x9b, 0xc4, 0xc3, 0xd8, 0xa1, 0x5d,
	0x92, 0x48, 0xdc, 0x00, 0x89, 0x4b, 0x70, 0x0b, 0x8e, 0x12, 0x24, 0x2e, 0x80, 0x14, 0x0e, 0x00,
	0xb2, 0x9d, 0xc9, 0x4c, 0xff, 0x4c, 0xe1, 0xc3, 0xd2, 0x7e, 0xe9, 0xd8, 0xcf, 0xf3, 0xbc, 0xfe,
	0xd9, 0x8a, 0xfc, 0x1a, 0xbc, 0xdd, 0x22, 0xbc, 0x7e, 0x4a, 0x1b, 0x75, 0x8e, 0xd9, 0x4b, 0x56,
	0x77, 0x08, 0xc7, 0x6e, 0x9b, 0xe5, 0x3c, 0x9f, 0x72, 0x0a, 0x17, 0x03, 0xe1, 0xbe, 0xdd, 0x72,
	0xf9, 0x8b, 0x5e, 0x23, 0xd7, 0xa4, 0x9d, 0x7c, 0x8b, 0xb6, 0x68, 0x5e, 0x1a, 0x1a, 0xbd, 0x13,
	0x39, 0x92, 0x03, 0xf9, 0xa5, 0x82, 0xf7, 0xbf, 0x6d, 0xd1, 0x1c, 0xc1, 0xec, 0x15, 0xf5, 0x58,
	0xae, 0x4d, 0x9b, 0xb8, 0x9d, 0x6f, 0xd2, 0x2e, 0xf7, 0x71, 0x93, 0x33, 0x95, 0xf4, 0x89, 0x47,
	0xed, 0x0e, 0x75, 0x48, 0x9b, 0xe5, 0x27, 0xc6, 0xbc, 0x1c, 0xe6, 0xa9, 0xc7, 0xea, 0xb8, 0xc7,
	0x69, 0x07, 0x73, 0x97, 0x76, 0xf3, 0x1d, 0xec, 0xb6, 0xeb, 0x6e, 0xf7, 0x24, 0xa8, 0xbc, 0x13,
	0x01, 0xe9, 0x9c, 0xb9, 0xfc, 0x25, 0x3d, 0xcb, 0xb7, 0xa8, 0x2d, 0x45, 0xfb, 0x07, 0xdc, 0x76,
	0x1d, 0xcc, 0xa9, 0xcf, 0xf2, 0xd3, 0x4f, 0x95, 0x33, 0x0e, 0x81, 0xf6, 0x19, 0xe1, 0x5f, 0xd0,
	0xc6, 0x81, 0xd8, 0xe7, 0x27, 0x72, 0x9b, 0x88, 0x7c, 0xdf, 0x23, 0x8c, 0xc3, 0x0a, 0x48, 0x9f,
	0x2a, 0xa1, 0xea, 0x68, 0x89, 0x4c, 0xc2, 0x4c, 0xef, 0xbd, 0x33, 0x1e, 0xe9, 0x6b, 0x27, 0xd4,
	0xef, 0x54, 0x8c, 0xa9, 0x64, 0xfc, 0xf1, 0xbb, 0x3e, 0xe7, 0x25, 0xce, 0x0d, 0x14, 0xda, 0x8d,
	0x3f, 0xd7, 0xc1, 0x5b, 0xd7, 0x14, 0x66, 0x1e, 0xed, 0x32, 0x02, 0x8b, 0x60, 0xe1, 0x94, 0x36,
	0xbe, 0xc2, 0x1d, 0x32, 0xa9, 0xab, 0x8d, 0x47, 0xfa, 0x9d, 0x69, 0x5d, 0x21, 0x88, 0xaa, 0x49,
	0x6f, 0x06, 0x05, 0x46, 0xf8, 0x04, 0x2c, 0x76, 0x48, 0xb7, 0x27, 0x43, 0xc9, 0x29, 0xcc, 0xaa,
	0x0a, 0x05, 0x8a, 0x48, 0xa5, 0xbc, 0x99, 0xf3, 0x9f, 0x92, 0x68, 0xea, 0x86, 0x16, 0x98, 0x27,
	0xe7, 0xa4, 0x59, 0x75, 0xb4, 0x59, 0x99, 0x5b, 0x1f, 0x8f, 0xf4, 0x15, 0x95, 0x53, 0xf3, 0x06,
	0x9a, 0x18, 0xe0, 0x87, 0x60, 0xd9, 0x27, 0x8c, 0xf6, 0xfc, 0x26, 0x39, 0x78, 0xe5, 0x11, 0x6d,
	0x4e, 0x06, 0xde, 0x1c, 0x8f, 0xf4, 0xbb, 0x2a, 0x10, 0x55, 0x0d, 0x74, 0xc1, 0x0c, 0xcb, 0x00,
	0x04, 0xe3, 0xaa, 0xa3, 0xa5, 0x64, 0xf4, 0x8d, 0xf1, 0x48, 0x5f, 0xbf, 0x18, 0x15, 0xeb, 0x45,
	0x8c, 0xf0, 0x09, 0x58, 0x0a, 0x46, 0x87, 0x55, 0x47, 0x9b, 0x97, 0xb9, 0x7b, 0xe3, 0x91, 0x0e,
	0x2f, 0xe6, 0x0e, 0x45, 0x30, 0x6a, 0x85, 0x1f, 0x83, 0x95, 0xe9, 0x50, 0x9e, 0xcb, 0xc2, 0xf4,
	0x30, 0x37, 0x2e, 0x65, 0xe5, 0xe1, 0xa0, 0x8b, 0x76, 0xf8, 0x5b, 0x1a, 0xa4, 0x5e, 0x50, 0xc6,
	0x99, 0xb6, 0x98, 0x99, 0x35, 0xd3, 0x7b, 0xbf, 0xa6, 0xc7, 0x23, 0x7d, 0x59, 0x25, 0xe5, 0xbc,
	0x38, 0xce, 0x5f, 0xd2, 0xe0, 0xe7, 0xf4, 0x73, 0xd3, 0x2c, 0x9a, 0xe5, 0xa3, 0x2d, 0xbb, 0x7c,
	0xdc, 0x2f, 0x0c, 0x07, 0x47, 0x5b, 0x76, 0xe9, 0xb8, 0xe6, 0xf4, 0x0b, 0x43, 0x4b, 0x7c, 0x17,
	0x8e, 0x77, 0xc5, 0x20, 0x5b, 0x1c, 0x5a, 0x66, 0x2d, 0xf7, 0x1f, 0x9d, 0x56, 0x7f, 0x7b, 0x68,
	0x0d, 0x6a, 0xec, 0x91, 0x69, 0x9a, 0x47, 0x5b, 0xf6, 0xd3, 0x67, 0xf6, 0x3e, 0xb6, 0x4f, 0x8e,
	0xfb, 0x85, 0x6c, 0x69, 0x58, 0xb1, 0xfa, 0x8f, 0x87, 0x57, 0x66, 0x07, 0x15, 0xcb, 0x1a, 0x5c,
	0x6b, 0xde, 0x19, 0x9a, 0x95, 0x2b, 0x6e, 0xd3, 0x2c, 0x2a, 0x8e, 0x41, 0x71, 0x42, 0x31, 0x28,
	0xd4, 0x9c, 0x9a, 0x33, 0x38, 0x2a, 0xd8, 0x4f, 0x05, 0x87, 0x82, 0xfd, 0x17, 0x8f, 0xc2, 0x8c,
	0x5d, 0xb9, 0x3c, 0x34, 0xcd, 0xab, 0x6b, 0x5b, 0x6a, 0x8b, 0x83, 0xca, 0xad, 0x30, 0x94, 0x62,
	0x19, 0x44, 0xec, 0x3a, 0x69, 0xf7, 0x75, 0x82, 0xdd, 0x40, 0xb6, 0x1d, 0x4b, 0x56, 0x8a, 0x21,
	0xeb, 0x6f, 0x65, 0x8b, 0xc3, 0x5b, 0xa2, 0x2b, 0xc6, 0xd2, 0x95, 0xe3, 0xe9, 0xb6, 0x6f, 0x8b,
	0xae, 0x10, 0x4b, 0xb7, 0x13, 0x4f, 0x57, 0xfa, 0x3f, 0xe8, 0x2a, 0x71, 0x20, 0x8f, 0xe3, 0x41,
	0xca, 0xaf, 0x1f, 0xc4, 0x32, 0xdf, 0xcf, 0x7d, 0x60, 0xed, 0xd6, 0xd8, 0xa3, 0x87, 0x48, 0x5d,
	0x5b, 0xf0, 0x23, 0x30, 0x27, 0x1a, 0xa1, 0x96, 0xce, 0x24, 0xcc, 0xa5, 0xa2, 0x96, 0xbb, 0xd8,
	0x24, 0x73, 0x5f, 0x62, 0xb7, 0x5d, 0xed, 0x9e, 0xd0, 0xbd, 0xd5, 0xf1, 0x48, 0x5f, 0x9a, 0x74,
	0x0a, 0xec, 0xb6, 0x0d, 0x24, 0x63, 0x70, 0x1f, 0x80, 0xa6, 0x4f, 0x30, 0x27, 0x07, 0x6e, 0x87,
	0x68, 0x40, 0x5e, 0x9e, 0x9b, 0xe1, 0x85, 0x1d, 0x6a, 0xe2, 0x1e, 0x5c, 0xf1, 0xfe, 0x0e, 0xfe,
	0x12, 0xe7, 0xef, 0xa1, 0x48, 0x52, 0xd4, 0xe9, 0x79, 0x4e, 0x50, 0x67, 0xe9, 0x72, 0x9d, 0x50,
	0xbb, 0xae, 0x4e, 0xa8, 0xc2, 0x7d, 0xb0, 0x20, 0xab, 0x52, 0x5f, 0x5b, 0x96, 0x45, 0xb2, 0x61,
	0x5b, 0x9c, 0x08, 0xa2, 0xc2, 0x3d, 0xb0, 0xf1, 0xfc, 0xe8, 0x99, 0xfd, 0x1d, 0xb6, 0x7f, 0xac,
	0xdb, 0xc7, 0xb5, 0xb3, 0xfe, 0x76, 0x76, 0xa7, 0x34, 0x7c, 0x88, 0x82, 0x30, 0xdc, 0x04, 0xb3,
	0xd4, 0x6f, 0x69, 0x2b, 0x99, 0x84, 0x99, 0xda, 0xdb, 0x18, 0x8f, 0x74, 0xa0, 0x6a, 0x50, 0xbf,
	0x25, 0xdb, 0xea, 0xda, 0x0c, 0x12, 0x06, 0xf8, 0x2e, 0x48, 0xba, 0x8e, 0x76, 0x47, 0x2e, 0xb5,
	0x32, 0x1e, 0xe9, 0x69, 0x65, 0x73, 0x1d, 0x03, 0x25, 0x5d, 0x07, 0x6e, 0x82, 0xd4, 0x29, 0x6d,
	0x54, 0x1d, 0x6d, 0x55, 0x3a, 0xd6, 0xc2, 0xe6, 0x20, 0xa7, 0x0d, 0xa4, 0x64, 0x98, 0x05, 0x0b,
	0xdc, 0x77, 0x5b, 0x2d, 0xe2, 0x6b, 0x6b, 0xd2, 0x09, 0x43, 0xec, 0x89, 0x60, 0xa0, 0xc0, 0x02,
	0xab, 0x60, 0x51, 0x34, 0xdb, 0xaf, 0x19, 0xf1, 0xb5, 0x75, 0x69, 0xb7, 0xc3, 0x3e, 0x1e, 0x28,
	0x37, 0x6d, 0x73, 0x1a, 0x17, 0x8d, 0x9d, 0x71, 0xcc, 0x7b, 0x4c, 0x83, 0x97, 0x1b, 0xbb, 0x9a,
	0x37, 0xd0, 0xc4, 0x20, 0x9a, 0x2c, 0xeb, 0x35, 0x9b, 0x84, 0x31, 0x84, 0x39, 0xd1, 0xee, 0x66,
	0x12, 0x66, 0x32, 0xda, 0x64, 0x23, 0xa2, 0x81, 0xa2, 0x56, 0x71, 0x0a, 0xc4, 0xf7, 0xa9, 0xaf,
	0x6d, 0x5c, 0x3e, 0x05, 0x39, 0x6d, 0x20, 0x25, 0x1b, 0x7f, 0x25, 0x40, 0x26, 0xf6, 0xc5, 0xf3,
	0x8d, 0x8f, 0x3d, 0x8f, 0xf8, 0xf0, 0x01, 0x98, 0x6b, 0x52, 0x47, 0xbd, 0x7a, 0x52, 0xd1, 0x9f,
	0xa5, 0x98, 0x35, 0x90, 0x14, 0x05, 0xab, 0xf8, 0xff, 0xe9, 0xb9, 0xd7, 0xc6, 0x6e, 0x77, 0xf2,
	0xd8, 0x89, 0xb0, 0x46, 0x44, 0x03, 0x45, 0xad, 0x21, 0xeb, 0xec, 0x8d, 0xac, 0xf0, 0x73, 0x30,
	0xe7, 0x60, 0x8e, 0xe5, 0xf3, 0x66, 0xa9, 0xf8, 0x20, 0x17, 0xbc, 0x67, 0x73, 0xb1, 0x1b, 0x88,
	0xb2, 0x8a, 0xa8, 0x81, 0x64, 0x85, 0xc6, 0xbc, 0x7c, 0x46, 0x6e, 0xff, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x27, 0x46, 0xf3, 0x30, 0x0b, 0x00, 0x00,
}
