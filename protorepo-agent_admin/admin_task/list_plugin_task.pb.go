// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_plugin_task.proto

package admin_task

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
//ListPluginTask请求
type ListPluginTaskRequest struct {
	//
	//agent插件ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//页码
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页数据量
	PageSize             int32    `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginTaskRequest) Reset()         { *m = ListPluginTaskRequest{} }
func (m *ListPluginTaskRequest) String() string { return proto.CompactTextString(m) }
func (*ListPluginTaskRequest) ProtoMessage()    {}
func (*ListPluginTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ef1fd70c77d328, []int{0}
}
func (m *ListPluginTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginTaskRequest.Unmarshal(m, b)
}
func (m *ListPluginTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginTaskRequest.Marshal(b, m, deterministic)
}
func (m *ListPluginTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginTaskRequest.Merge(m, src)
}
func (m *ListPluginTaskRequest) XXX_Size() int {
	return xxx_messageInfo_ListPluginTaskRequest.Size(m)
}
func (m *ListPluginTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginTaskRequest proto.InternalMessageInfo

func (m *ListPluginTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListPluginTaskRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginTaskRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//ListPluginTask返回
type ListPluginTaskResponse struct {
	//
	//页数
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//总数
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"total"`
	//
	//数据列表
	List                 []*ListPluginTaskResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListPluginTaskResponse) Reset()         { *m = ListPluginTaskResponse{} }
func (m *ListPluginTaskResponse) String() string { return proto.CompactTextString(m) }
func (*ListPluginTaskResponse) ProtoMessage()    {}
func (*ListPluginTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ef1fd70c77d328, []int{1}
}
func (m *ListPluginTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginTaskResponse.Unmarshal(m, b)
}
func (m *ListPluginTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginTaskResponse.Marshal(b, m, deterministic)
}
func (m *ListPluginTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginTaskResponse.Merge(m, src)
}
func (m *ListPluginTaskResponse) XXX_Size() int {
	return xxx_messageInfo_ListPluginTaskResponse.Size(m)
}
func (m *ListPluginTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginTaskResponse proto.InternalMessageInfo

func (m *ListPluginTaskResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPluginTaskResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPluginTaskResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListPluginTaskResponse) GetList() []*ListPluginTaskResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListPluginTaskResponse_List struct {
	//
	//任务对象
	TargetId string `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId" form:"targetId"`
	//
	//任务对象名称
	TargetName string `protobuf:"bytes,2,opt,name=targetName,proto3" json:"targetName" form:"targetName"`
	//
	//执行用户
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//任务状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status" form:"status"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,5,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//结束时间
	Etime int32 `protobuf:"varint,6,opt,name=etime,proto3" json:"etime" form:"etime"`
	//
	//部署信息
	DeployList           []*ListPluginTaskResponse_List_DeployList `protobuf:"bytes,7,rep,name=deployList,proto3" json:"deployList" form:"deployList"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ListPluginTaskResponse_List) Reset()         { *m = ListPluginTaskResponse_List{} }
func (m *ListPluginTaskResponse_List) String() string { return proto.CompactTextString(m) }
func (*ListPluginTaskResponse_List) ProtoMessage()    {}
func (*ListPluginTaskResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ef1fd70c77d328, []int{1, 0}
}
func (m *ListPluginTaskResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginTaskResponse_List.Unmarshal(m, b)
}
func (m *ListPluginTaskResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginTaskResponse_List.Marshal(b, m, deterministic)
}
func (m *ListPluginTaskResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginTaskResponse_List.Merge(m, src)
}
func (m *ListPluginTaskResponse_List) XXX_Size() int {
	return xxx_messageInfo_ListPluginTaskResponse_List.Size(m)
}
func (m *ListPluginTaskResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginTaskResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginTaskResponse_List proto.InternalMessageInfo

func (m *ListPluginTaskResponse_List) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *ListPluginTaskResponse_List) GetTargetName() string {
	if m != nil {
		return m.TargetName
	}
	return ""
}

func (m *ListPluginTaskResponse_List) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ListPluginTaskResponse_List) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ListPluginTaskResponse_List) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *ListPluginTaskResponse_List) GetEtime() int32 {
	if m != nil {
		return m.Etime
	}
	return 0
}

func (m *ListPluginTaskResponse_List) GetDeployList() []*ListPluginTaskResponse_List_DeployList {
	if m != nil {
		return m.DeployList
	}
	return nil
}

type ListPluginTaskResponse_List_DeployList struct {
	//
	//AgentID
	AgentId string `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId" form:"agentId"`
	//
	//ip
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//部署路径
	DeployPath []string `protobuf:"bytes,3,rep,name=deployPath,proto3" json:"deployPath" form:"deployPath"`
	//
	//插件包ID
	AgentPluginId string `protobuf:"bytes,4,opt,name=agentPluginId,proto3" json:"agentPluginId" form:"agentPluginId"`
	//
	//插件包名称
	AgentPluginName string `protobuf:"bytes,5,opt,name=agentPluginName,proto3" json:"agentPluginName" form:"agentPluginName"`
	//
	//插件包部署前名称
	PreVersionName string `protobuf:"bytes,6,opt,name=preVersionName,proto3" json:"preVersionName" form:"preVersionName"`
	//
	//插件包版本名称
	PluginVersionName string `protobuf:"bytes,7,opt,name=pluginVersionName,proto3" json:"pluginVersionName" form:"pluginVersionName"`
	//
	//插件包部署状态
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPluginTaskResponse_List_DeployList) Reset() {
	*m = ListPluginTaskResponse_List_DeployList{}
}
func (m *ListPluginTaskResponse_List_DeployList) String() string { return proto.CompactTextString(m) }
func (*ListPluginTaskResponse_List_DeployList) ProtoMessage()    {}
func (*ListPluginTaskResponse_List_DeployList) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ef1fd70c77d328, []int{1, 0, 0}
}
func (m *ListPluginTaskResponse_List_DeployList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginTaskResponse_List_DeployList.Unmarshal(m, b)
}
func (m *ListPluginTaskResponse_List_DeployList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginTaskResponse_List_DeployList.Marshal(b, m, deterministic)
}
func (m *ListPluginTaskResponse_List_DeployList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginTaskResponse_List_DeployList.Merge(m, src)
}
func (m *ListPluginTaskResponse_List_DeployList) XXX_Size() int {
	return xxx_messageInfo_ListPluginTaskResponse_List_DeployList.Size(m)
}
func (m *ListPluginTaskResponse_List_DeployList) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginTaskResponse_List_DeployList.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginTaskResponse_List_DeployList proto.InternalMessageInfo

func (m *ListPluginTaskResponse_List_DeployList) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetDeployPath() []string {
	if m != nil {
		return m.DeployPath
	}
	return nil
}

func (m *ListPluginTaskResponse_List_DeployList) GetAgentPluginId() string {
	if m != nil {
		return m.AgentPluginId
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetAgentPluginName() string {
	if m != nil {
		return m.AgentPluginName
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetPreVersionName() string {
	if m != nil {
		return m.PreVersionName
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetPluginVersionName() string {
	if m != nil {
		return m.PluginVersionName
	}
	return ""
}

func (m *ListPluginTaskResponse_List_DeployList) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//ListPluginTaskApi返回
type ListPluginTaskResponseWrapper struct {
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
	Data                 *ListPluginTaskResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListPluginTaskResponseWrapper) Reset()         { *m = ListPluginTaskResponseWrapper{} }
func (m *ListPluginTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListPluginTaskResponseWrapper) ProtoMessage()    {}
func (*ListPluginTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ef1fd70c77d328, []int{2}
}
func (m *ListPluginTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPluginTaskResponseWrapper.Unmarshal(m, b)
}
func (m *ListPluginTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPluginTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListPluginTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPluginTaskResponseWrapper.Merge(m, src)
}
func (m *ListPluginTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListPluginTaskResponseWrapper.Size(m)
}
func (m *ListPluginTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPluginTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListPluginTaskResponseWrapper proto.InternalMessageInfo

func (m *ListPluginTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListPluginTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListPluginTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListPluginTaskResponseWrapper) GetData() *ListPluginTaskResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPluginTaskRequest)(nil), "admin_task.ListPluginTaskRequest")
	proto.RegisterType((*ListPluginTaskResponse)(nil), "admin_task.ListPluginTaskResponse")
	proto.RegisterType((*ListPluginTaskResponse_List)(nil), "admin_task.ListPluginTaskResponse.List")
	proto.RegisterType((*ListPluginTaskResponse_List_DeployList)(nil), "admin_task.ListPluginTaskResponse.List.DeployList")
	proto.RegisterType((*ListPluginTaskResponseWrapper)(nil), "admin_task.ListPluginTaskResponseWrapper")
}

func init() { proto.RegisterFile("list_plugin_task.proto", fileDescriptor_33ef1fd70c77d328) }

var fileDescriptor_33ef1fd70c77d328 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0xfe, 0xe5, 0x6f, 0x93, 0x93, 0x5f, 0xff, 0x8d, 0x36, 0xac, 0xc1, 0xb2, 0x61, 0x04, 0x6d,
	0xc1, 0x26, 0x50, 0xa9, 0x8a, 0x17, 0x82, 0xa1, 0x22, 0x95, 0x22, 0x65, 0x14, 0xbd, 0x2c, 0xd3,
	0xec, 0x74, 0x3b, 0x34, 0xc9, 0xac, 0xbb, 0x13, 0xab, 0x3e, 0x89, 0xcf, 0xe4, 0x95, 0x37, 0xde,
	0xae, 0xe0, 0x0b, 0x08, 0xfb, 0x04, 0x32, 0x67, 0x36, 0xd9, 0x69, 0x5a, 0xa8, 0x57, 0xb3, 0x7b,
	0xbe, 0xef, 0x3b, 0x73, 0xf6, 0x3b, 0xe7, 0x2c, 0xb4, 0x47, 0x32, 0xd1, 0xc7, 0xd1, 0x68, 0x1a,
	0xca, 0xc9, 0xb1, 0xe6, 0xc9, 0x79, 0x2f, 0x8a, 0x95, 0x56, 0x04, 0x78, 0x30, 0xce, 0x23, 0x9d,
	0x9d, 0x50, 0xea, 0xb3, 0xe9, 0x49, 0x6f, 0xa8, 0xc6, 0xfd, 0x50, 0x85, 0xaa, 0x8f, 0x94, 0x93,
	0xe9, 0x29, 0xbe, 0xe1, 0x0b, 0x3e, 0x59, 0x69, 0xe7, 0xb1, 0x43, 0x1f, 0x5f, 0x48, 0x7d, 0xae,
	0x2e, 0xfa, 0xa1, 0xda, 0x41, 0x70, 0xe7, 0x13, 0x1f, 0xc9, 0x80, 0x6b, 0x15, 0x27, 0xfd, 0xf9,
	0xa3, 0xd5, 0xd1, 0x6f, 0x25, 0xd8, 0x38, 0x94, 0x89, 0x3e, 0xc2, 0x62, 0xde, 0xf1, 0xe4, 0x9c,
	0x89, 0x8f, 0x53, 0x91, 0x68, 0xb2, 0x09, 0x65, 0x19, 0x78, 0xa5, 0x6e, 0x69, 0xab, 0x39, 0x58,
	0xce, 0x52, 0xbf, 0x79, 0xaa, 0xe2, 0xf1, 0x33, 0x2a, 0x03, 0xca, 0xca, 0x32, 0x20, 0xdb, 0x50,
	0x8d, 0x78, 0x28, 0xbc, 0x72, 0xb7, 0xb4, 0x55, 0x1b, 0x6c, 0x64, 0xa9, 0xdf, 0xb2, 0x04, 0x13,
	0xa5, 0xbf, 0x7f, 0xf9, 0xe5, 0xb5, 0xff, 0x18, 0x52, 0xc8, 0x1e, 0x34, 0xcc, 0xf9, 0x56, 0x7e,
	0x15, 0x5e, 0x05, 0xe9, 0x77, 0xb2, 0xd4, 0x5f, 0x2d, 0xe8, 0x06, 0x99, 0x49, 0xe6, 0x54, 0xfa,
	0xbd, 0x01, 0xed, 0xc5, 0xd2, 0x92, 0x48, 0x4d, 0x12, 0x31, 0xbf, 0xbc, 0x74, 0xf3, 0xe5, 0x4f,
	0xa0, 0x69, 0xce, 0xe3, 0xc4, 0xdc, 0x6e, 0x8b, 0xed, 0x64, 0xa9, 0xbf, 0x56, 0xf0, 0x11, 0xba,
	0x72, 0x3d, 0xb9, 0x0f, 0x35, 0xad, 0x34, 0x1f, 0xe5, 0x25, 0xaf, 0x65, 0xa9, 0xff, 0xbf, 0x15,
	0x61, 0x98, 0x32, 0x0b, 0x93, 0x43, 0xa8, 0x9a, 0x76, 0x7a, 0xd5, 0x6e, 0x65, 0xab, 0xb5, 0xfb,
	0xa0, 0x57, 0xf4, 0xb0, 0x77, 0x7d, 0xf5, 0x18, 0x1e, 0xac, 0x16, 0x45, 0x1b, 0x39, 0x65, 0x98,
	0xa5, 0xf3, 0xb3, 0x0e, 0x55, 0x83, 0x93, 0x3e, 0x34, 0x34, 0x8f, 0x43, 0xa1, 0x0f, 0x66, 0x4d,
	0xb8, 0x55, 0x98, 0x36, 0x43, 0x28, 0x9b, 0x93, 0xc8, 0x1e, 0x80, 0x7d, 0x7e, 0xc3, 0xc7, 0xf6,
	0x4b, 0x9b, 0xe8, 0xcc, 0xba, 0x2b, 0x31, 0x18, 0x65, 0x0e, 0x91, 0x3c, 0x84, 0xa5, 0x61, 0x2c,
	0xcc, 0x44, 0xe0, 0x87, 0x36, 0x07, 0x24, 0x4b, 0xfd, 0x15, 0xab, 0xc9, 0x01, 0xca, 0x66, 0x14,
	0xb2, 0x0d, 0xf5, 0x44, 0x73, 0x3d, 0x4d, 0xbc, 0x2a, 0x92, 0xd7, 0xb3, 0xd4, 0x5f, 0xb6, 0x64,
	0x1b, 0xa7, 0x2c, 0x27, 0x18, 0xff, 0x86, 0x5a, 0x8e, 0x85, 0x57, 0x5b, 0xf4, 0x0f, 0xc3, 0x94,
	0x59, 0xd8, 0xf0, 0x04, 0xf2, 0xea, 0x8b, 0x3c, 0x91, 0xf3, 0xf0, 0x24, 0x12, 0x20, 0x10, 0xd1,
	0x48, 0x7d, 0x31, 0xf6, 0x78, 0x4b, 0xe8, 0xf6, 0xee, 0x3f, 0xba, 0xdd, 0xdb, 0x9f, 0x2b, 0x5d,
	0x4f, 0x8a, 0x7c, 0x94, 0x39, 0xc9, 0x3b, 0x3f, 0x2a, 0x00, 0x85, 0xc2, 0x58, 0xc4, 0x43, 0x31,
	0x29, 0x3a, 0xe1, 0x58, 0x94, 0x03, 0x94, 0xcd, 0x28, 0xb8, 0x37, 0x51, 0xee, 0xbf, 0xbb, 0x37,
	0x91, 0xd9, 0x9b, 0xc8, 0xb4, 0xc9, 0xde, 0x74, 0xc4, 0xf5, 0x99, 0x57, 0xe9, 0x56, 0x2e, 0xb7,
	0xa9, 0xc0, 0xe6, 0x25, 0x99, 0x17, 0xf2, 0x1c, 0x96, 0xf1, 0x02, 0xfb, 0x81, 0x07, 0x41, 0xee,
	0xbf, 0x97, 0xa5, 0xfe, 0x6d, 0xa7, 0x92, 0x19, 0x4c, 0xd9, 0x65, 0x3a, 0xd9, 0x87, 0x55, 0x27,
	0x80, 0x23, 0x52, 0xc3, 0x0c, 0x66, 0x19, 0xda, 0x57, 0x32, 0xd8, 0x39, 0x59, 0x94, 0x90, 0x17,
	0xb0, 0x12, 0xc5, 0xe2, 0xbd, 0x88, 0x13, 0xa9, 0x6c, 0x92, 0x3a, 0x26, 0x31, 0xfb, 0xbc, 0x91,
	0x6f, 0xd4, 0x25, 0x9c, 0xb2, 0x05, 0x01, 0x79, 0x0d, 0xeb, 0xf6, 0xc7, 0xe7, 0x66, 0x59, 0xc2,
	0x2c, 0x77, 0xb3, 0xd4, 0xf7, 0xf2, 0x2c, 0x8b, 0x14, 0xca, 0xae, 0xca, 0x9c, 0x69, 0x6c, 0xdc,
	0x30, 0x8d, 0xf4, 0x4f, 0x09, 0x36, 0xaf, 0x1f, 0x90, 0x0f, 0x31, 0x8f, 0x22, 0x11, 0x93, 0x7b,
	0x50, 0x1d, 0xaa, 0x60, 0xf6, 0x4f, 0x71, 0xd6, 0xd3, 0x44, 0x29, 0x43, 0x90, 0x3c, 0x85, 0x96,
	0x39, 0x5f, 0x7e, 0x8e, 0x46, 0x5c, 0x4e, 0xf2, 0x2e, 0xb7, 0xb3, 0xd4, 0x27, 0x05, 0x37, 0x07,
	0x29, 0x73, 0xa9, 0x38, 0xe6, 0x71, 0x3c, 0xdf, 0x32, 0x77, 0xcc, 0x4d, 0xd8, 0x8c, 0xb9, 0x39,
	0xc9, 0x2b, 0xa8, 0x06, 0x5c, 0x73, 0xec, 0x6f, 0x6b, 0x97, 0xde, 0x3c, 0xe0, 0x6e, 0xa9, 0x46,
	0x49, 0x19, 0x26, 0x38, 0xa9, 0xe3, 0x0f, 0xfe, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x02, 0xda, 0xd8, 0x6d, 0x06, 0x00, 0x00,
}
