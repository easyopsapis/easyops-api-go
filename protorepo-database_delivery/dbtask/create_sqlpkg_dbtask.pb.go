// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_sqlpkg_dbtask.proto

package dbtask

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
//CreateSQLPackageDBTask请求
type CreateSQLPackageDBTaskRequest struct {
	//
	//SQL包ID
	PkgId string `protobuf:"bytes,1,opt,name=pkgId,proto3" json:"pkgId" form:"pkgId"`
	//
	//SQL包版本ID
	VersionId string `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//创建SQL包变更任务请求参数
	TaskInfo             *CreateSQLPackageDBTaskRequest_TaskInfo `protobuf:"bytes,3,opt,name=taskInfo,proto3" json:"taskInfo" form:"taskInfo"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest) Reset()         { *m = CreateSQLPackageDBTaskRequest{} }
func (m *CreateSQLPackageDBTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageDBTaskRequest) ProtoMessage()    {}
func (*CreateSQLPackageDBTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0}
}
func (m *CreateSQLPackageDBTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest) GetPkgId() string {
	if m != nil {
		return m.PkgId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest) GetTaskInfo() *CreateSQLPackageDBTaskRequest_TaskInfo {
	if m != nil {
		return m.TaskInfo
	}
	return nil
}

type CreateSQLPackageDBTaskRequest_TaskInfo struct {
	//
	//数据库服务ID
	DbServiceId string `protobuf:"bytes,1,opt,name=dbServiceId,proto3" json:"dbServiceId" form:"dbServiceId"`
	//
	//备份任务配置，可以为空
	BackupCfg []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg `protobuf:"bytes,2,rep,name=backupCfg,proto3" json:"backupCfg" form:"backupCfg"`
	//
	//变更任务配置
	ChangeCfg []*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg `protobuf:"bytes,3,rep,name=changeCfg,proto3" json:"changeCfg" form:"changeCfg"`
	//
	//检查任务配置，可以为空
	CheckCfg             []*CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg `protobuf:"bytes,4,rep,name=checkCfg,proto3" json:"checkCfg" form:"checkCfg"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageDBTaskRequest_TaskInfo) ProtoMessage()    {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo) GetDbServiceId() string {
	if m != nil {
		return m.DbServiceId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo) GetBackupCfg() []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg {
	if m != nil {
		return m.BackupCfg
	}
	return nil
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo) GetChangeCfg() []*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg {
	if m != nil {
		return m.ChangeCfg
	}
	return nil
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo) GetCheckCfg() []*CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg {
	if m != nil {
		return m.CheckCfg
	}
	return nil
}

type CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg struct {
	//
	//数据库实例ID
	DbInstanceId string `protobuf:"bytes,1,opt,name=dbInstanceId,proto3" json:"dbInstanceId" form:"dbInstanceId"`
	//
	//主机ID
	HostId string `protobuf:"bytes,2,opt,name=hostId,proto3" json:"hostId" form:"hostId"`
	//
	//主机IP
	HostIp string `protobuf:"bytes,3,opt,name=hostIp,proto3" json:"hostIp" form:"hostIp"`
	//
	//备份任务脚本列表
	Scripts              []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts `protobuf:"bytes,4,rep,name=scripts,proto3" json:"scripts" form:"scripts"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 0}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) GetDbInstanceId() string {
	if m != nil {
		return m.DbInstanceId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) GetHostId() string {
	if m != nil {
		return m.HostId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) GetHostIp() string {
	if m != nil {
		return m.HostIp
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) GetScripts() []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts {
	if m != nil {
		return m.Scripts
	}
	return nil
}

type CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts struct {
	//
	//脚本路径
	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script" form:"script"`
	//
	//脚本参数
	Variables            []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables" form:"variables"`
	XXX_NoUnkeyedLiteral struct{}                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                `json:"-"`
	XXX_sizecache        int32                                                                 `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 0, 0}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) GetVariables() []*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables {
	if m != nil {
		return m.Variables
	}
	return nil
}

type CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables struct {
	//
	//参数名称
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" form:"key"`
	//
	//参数值
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value" form:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 0, 0, 0}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg struct {
	//
	//数据库实例ID
	DbInstanceId string `protobuf:"bytes,1,opt,name=dbInstanceId,proto3" json:"dbInstanceId" form:"dbInstanceId"`
	//
	//数据库实例用户名（临时）
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" form:"username"`
	//
	//数据库实例密码（临时）
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" form:"password"`
	//
	//检查任务脚本路径列表
	Scripts              []*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts `protobuf:"bytes,4,rep,name=scripts,proto3" json:"scripts" form:"scripts"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 1}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) GetDbInstanceId() string {
	if m != nil {
		return m.DbInstanceId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) GetScripts() []*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts {
	if m != nil {
		return m.Scripts
	}
	return nil
}

type CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts struct {
	//
	//变更SQL路径
	Update string `protobuf:"bytes,1,opt,name=update,proto3" json:"update" form:"update"`
	//
	//回退SQL路径
	Rollback             string   `protobuf:"bytes,2,opt,name=rollback,proto3" json:"rollback" form:"rollback"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 1, 0}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) GetUpdate() string {
	if m != nil {
		return m.Update
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) GetRollback() string {
	if m != nil {
		return m.Rollback
	}
	return ""
}

type CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg struct {
	//
	//数据库实例ID
	DbInstanceId string `protobuf:"bytes,1,opt,name=dbInstanceId,proto3" json:"dbInstanceId" form:"dbInstanceId"`
	//
	//数据库实例用户名（临时）
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" form:"username"`
	//
	//数据库实例密码（临时）
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" form:"password"`
	//
	//检查任务脚本路径（对象声明文件）
	Scripts              []string `protobuf:"bytes,4,rep,name=scripts,proto3" json:"scripts" form:"scripts"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) Reset() {
	*m = CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg{}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) String() string {
	return proto.CompactTextString(m)
}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) ProtoMessage() {}
func (*CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{0, 0, 2}
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg.Size(m)
}
func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) GetDbInstanceId() string {
	if m != nil {
		return m.DbInstanceId
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) GetScripts() []string {
	if m != nil {
		return m.Scripts
	}
	return nil
}

//
//CreateSQLPackageDBTask返回
type CreateSQLPackageDBTaskResponse struct {
	//
	//任务ID
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId" form:"taskId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageDBTaskResponse) Reset()         { *m = CreateSQLPackageDBTaskResponse{} }
func (m *CreateSQLPackageDBTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageDBTaskResponse) ProtoMessage()    {}
func (*CreateSQLPackageDBTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{1}
}
func (m *CreateSQLPackageDBTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponse.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskResponse.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponse.Size(m)
}
func (m *CreateSQLPackageDBTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskResponse proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

//
//CreateSQLPackageDBTaskApi返回
type CreateSQLPackageDBTaskResponseWrapper struct {
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
	Data                 *CreateSQLPackageDBTaskResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CreateSQLPackageDBTaskResponseWrapper) Reset()         { *m = CreateSQLPackageDBTaskResponseWrapper{} }
func (m *CreateSQLPackageDBTaskResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageDBTaskResponseWrapper) ProtoMessage()    {}
func (*CreateSQLPackageDBTaskResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_6925c7f4d72e1c90, []int{2}
}
func (m *CreateSQLPackageDBTaskResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper.Unmarshal(m, b)
}
func (m *CreateSQLPackageDBTaskResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageDBTaskResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper.Merge(m, src)
}
func (m *CreateSQLPackageDBTaskResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper.Size(m)
}
func (m *CreateSQLPackageDBTaskResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageDBTaskResponseWrapper proto.InternalMessageInfo

func (m *CreateSQLPackageDBTaskResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateSQLPackageDBTaskResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateSQLPackageDBTaskResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateSQLPackageDBTaskResponseWrapper) GetData() *CreateSQLPackageDBTaskResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSQLPackageDBTaskRequest)(nil), "dbtask.CreateSQLPackageDBTaskRequest")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.BackupCfg")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.BackupCfg.Scripts")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.BackupCfg.Scripts.Variables")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.ChangeCfg")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.ChangeCfg.Scripts")
	proto.RegisterType((*CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg)(nil), "dbtask.CreateSQLPackageDBTaskRequest.TaskInfo.CheckCfg")
	proto.RegisterType((*CreateSQLPackageDBTaskResponse)(nil), "dbtask.CreateSQLPackageDBTaskResponse")
	proto.RegisterType((*CreateSQLPackageDBTaskResponseWrapper)(nil), "dbtask.CreateSQLPackageDBTaskResponseWrapper")
}

func init() { proto.RegisterFile("create_sqlpkg_dbtask.proto", fileDescriptor_6925c7f4d72e1c90) }

var fileDescriptor_6925c7f4d72e1c90 = []byte{
	// 1050 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0x7f, 0x37, 0x9e, 0xec, 0x76, 0xcb, 0x2c, 0xb0, 0x51, 0x24, 0x70, 0x35, 0x6c, 0x57,
	0x76, 0xb7, 0x76, 0x9a, 0xbf, 0x76, 0x37, 0x08, 0x15, 0x52, 0x58, 0xa9, 0xec, 0x4a, 0xc0, 0x74,
	0x81, 0x8b, 0x26, 0x5b, 0x4d, 0xec, 0xc9, 0x8f, 0x9c, 0xc6, 0x5e, 0xdb, 0x49, 0x59, 0x12, 0xf3,
	0x02, 0x5c, 0x20, 0x21, 0xf1, 0x0e, 0xbc, 0x06, 0x97, 0xbc, 0x44, 0x90, 0xb8, 0xe5, 0x06, 0xe5,
	0x86, 0x0b, 0x6e, 0x90, 0x3d, 0x63, 0xc7, 0x25, 0x71, 0xf7, 0x97, 0x4a, 0x5c, 0xc5, 0x3e, 0xe7,
	0xfb, 0xce, 0xf9, 0xe6, 0x9c, 0xc9, 0x9c, 0x31, 0xc8, 0x6a, 0x36, 0x25, 0x2e, 0x3d, 0x71, 0x9e,
	0x74, 0x2d, 0xa3, 0x75, 0xa2, 0x37, 0x5c, 0xe2, 0x18, 0xaa, 0x65, 0x9b, 0xae, 0x09, 0x57, 0xd9,
	0x5b, 0x56, 0x69, 0x75, 0xdc, 0x76, 0xbf, 0xa1, 0x6a, 0xe6, 0x69, 0xae, 0x65, 0xb6, 0xcc, 0x5c,
	0xe0, 0x6e, 0xf4, 0x9b, 0xc1, 0x5b, 0xf0, 0x12, 0x3c, 0x31, 0x5a, 0x76, 0x37, 0x06, 0x3f, 0x3d,
	0xeb, 0xb8, 0x86, 0x79, 0x96, 0x6b, 0x99, 0x4a, 0xe0, 0x54, 0x06, 0xa4, 0xdb, 0xd1, 0x89, 0x6b,
	0xda, 0x4e, 0x2e, 0x7a, 0x64, 0x3c, 0xf4, 0xd7, 0x4d, 0xf0, 0xce, 0x41, 0xa0, 0xe6, 0xe8, 0x8b,
	0x87, 0x9f, 0x13, 0xcd, 0x20, 0x2d, 0xfa, 0x71, 0xf5, 0x11, 0x71, 0x0c, 0x4c, 0x9f, 0xf4, 0xa9,
	0xe3, 0xc2, 0x0a, 0x58, 0xb1, 0x8c, 0xd6, 0xa1, 0x9e, 0x59, 0xd8, 0x58, 0x90, 0x84, 0xea, 0xad,
	0xc9, 0x58, 0xbc, 0xda, 0x34, 0xed, 0xd3, 0x0a, 0x0a, 0xcc, 0xe8, 0xf7, 0xdf, 0xc4, 0x75, 0xb0,
	0xf6, 0xf8, 0x78, 0x47, 0xb9, 0x47, 0x94, 0x6f, 0xeb, 0xc3, 0x7c, 0xd1, 0xbb, 0x85, 0x19, 0x05,
	0xde, 0x07, 0xc2, 0x80, 0xda, 0x4e, 0xc7, 0xec, 0x1d, 0xea, 0x99, 0xc5, 0x80, 0x2f, 0x4d, 0xc6,
	0xe2, 0x3a, 0xe3, 0x47, 0xae, 0xf9, 0x31, 0xa6, 0x54, 0x78, 0x02, 0x52, 0x7e, 0x51, 0x0e, 0x7b,
	0x4d, 0x33, 0xb3, 0xb4, 0xb1, 0x20, 0xa5, 0x0b, 0xaa, 0xca, 0xab, 0x76, 0xa1, 0x78, 0xf5, 0x11,
	0x67, 0x55, 0x6f, 0x4c, 0xc6, 0xe2, 0x75, 0x96, 0x36, 0x8c, 0x84, 0x70, 0x14, 0x34, 0xfb, 0xe7,
	0x5b, 0x20, 0x15, 0x62, 0xe1, 0x43, 0x90, 0xd6, 0x1b, 0x47, 0xd4, 0x1e, 0x74, 0x34, 0x1a, 0xad,
	0x7b, 0x6b, 0x32, 0x16, 0x21, 0x0b, 0x10, 0x73, 0xce, 0x57, 0x1e, 0xa7, 0xc3, 0x2e, 0x10, 0x1a,
	0x44, 0x33, 0xfa, 0xd6, 0x41, 0xb3, 0x95, 0x59, 0xdc, 0x58, 0x92, 0xd2, 0x85, 0xbb, 0x2f, 0x26,
	0x5e, 0xad, 0x86, 0xfc, 0xea, 0x9b, 0xd3, 0xea, 0x45, 0x41, 0x11, 0x9e, 0x26, 0xf0, 0xb3, 0x69,
	0x6d, 0xd2, 0x6b, 0x51, 0x3f, 0xdb, 0xd2, 0x4b, 0x65, 0x3b, 0x08, 0xf9, 0xf1, 0x6c, 0x51, 0x50,
	0x84, 0xa7, 0x09, 0x60, 0x1b, 0xa4, 0xb4, 0x36, 0xd5, 0x0c, 0x3f, 0xd9, 0x72, 0x90, 0x6c, 0xef,
	0x85, 0x93, 0x31, 0x7a, 0xbc, 0x41, 0x61, 0x48, 0x84, 0xa3, 0xe8, 0xd9, 0x5f, 0xaf, 0x01, 0x21,
	0x2a, 0x03, 0x7c, 0x1f, 0x5c, 0xd5, 0x1b, 0x87, 0x3d, 0xc7, 0x25, 0xbd, 0x58, 0x8b, 0x6e, 0x4e,
	0xc6, 0xe2, 0x8d, 0xb0, 0x45, 0x53, 0x2f, 0xc2, 0xe7, 0xc0, 0xf0, 0x03, 0xb0, 0xda, 0x36, 0x1d,
	0x37, 0xda, 0x91, 0x9b, 0x93, 0xb1, 0x78, 0x8d, 0xd1, 0x98, 0x7d, 0x7e, 0x53, 0x39, 0x09, 0xfe,
	0x22, 0x70, 0xbe, 0x15, 0x6c, 0x45, 0xa1, 0xfa, 0xb3, 0xf0, 0xaf, 0x00, 0x96, 0x1f, 0xe0, 0x47,
	0x01, 0xfc, 0x20, 0x3c, 0x96, 0xa4, 0x82, 0x54, 0x3e, 0xde, 0x51, 0xca, 0xf5, 0x61, 0xde, 0x1b,
	0x1d, 0xef, 0x28, 0xa5, 0x7a, 0x4d, 0x1f, 0xe6, 0x3d, 0xd9, 0x7f, 0xce, 0xd7, 0xf7, 0xfd, 0x97,
	0xed, 0x82, 0x27, 0x4b, 0x35, 0xf5, 0x39, 0x91, 0xf2, 0xb0, 0xe8, 0xc9, 0xa3, 0x9a, 0xb3, 0x25,
	0x49, 0x92, 0x2f, 0xf0, 0x23, 0xe5, 0x3e, 0x51, 0x9a, 0xf5, 0x61, 0x7e, 0xbb, 0xe4, 0x55, 0xe4,
	0xe1, 0x9e, 0x37, 0x63, 0x1d, 0x55, 0x64, 0x79, 0x34, 0x17, 0xbc, 0xeb, 0x49, 0x95, 0x19, 0xb4,
	0x24, 0x15, 0x98, 0x8e, 0x51, 0x81, 0xab, 0x18, 0xe5, 0x6b, 0x7a, 0x4d, 0x1f, 0x1d, 0xe7, 0x95,
	0x7b, 0xbe, 0x0e, 0x26, 0xf6, 0x19, 0x18, 0x26, 0x33, 0x31, 0x73, 0xd9, 0x93, 0xa4, 0xd9, 0xdc,
	0x32, 0x5b, 0xe2, 0xa8, 0x72, 0x29, 0x1a, 0x4a, 0x89, 0x1a, 0x7c, 0xda, 0x3c, 0xd7, 0xfe, 0xeb,
	0x14, 0x76, 0x81, 0xb2, 0x62, 0xa2, 0xb2, 0x52, 0x82, 0xb2, 0xe1, 0xce, 0x76, 0xc1, 0xbb, 0x24,
	0x75, 0x85, 0x44, 0x75, 0xe5, 0x64, 0x75, 0xc5, 0xcb, 0x52, 0x97, 0x4f, 0x54, 0xb7, 0x9b, 0xac,
	0xae, 0xf4, 0x5f, 0xa8, 0xab, 0x24, 0x09, 0xd9, 0x4b, 0x16, 0x52, 0x7e, 0xfd, 0x42, 0x64, 0x69,
	0x53, 0xbd, 0x23, 0xef, 0xd7, 0x9c, 0xad, 0xf0, 0x0c, 0xb3, 0x60, 0x0f, 0x5c, 0x71, 0x34, 0xbb,
	0x63, 0xb9, 0x0e, 0x3f, 0xb6, 0x3f, 0x7c, 0xd9, 0x89, 0xa4, 0x1e, 0xb1, 0x38, 0x55, 0x38, 0x19,
	0x8b, 0x6b, 0xec, 0x10, 0xe4, 0xa1, 0x11, 0x0e, 0x93, 0x64, 0xbf, 0x5f, 0x04, 0x57, 0x38, 0x10,
	0xca, 0x60, 0x95, 0x99, 0xf9, 0xa9, 0xfd, 0xc6, 0xf4, 0xf4, 0x64, 0x76, 0x84, 0x39, 0x00, 0x7e,
	0x07, 0x84, 0x01, 0xb1, 0x3b, 0xa4, 0xd1, 0xa5, 0x0e, 0x1f, 0x9d, 0x9f, 0xbe, 0xaa, 0x50, 0xf5,
	0xab, 0x30, 0x62, 0x7c, 0xbc, 0x45, 0x69, 0x10, 0x9e, 0xa6, 0xcc, 0x7e, 0x09, 0x84, 0x08, 0x0d,
	0x37, 0xc0, 0x92, 0x41, 0x9f, 0x72, 0xd1, 0x6b, 0x93, 0xb1, 0x08, 0x18, 0xd5, 0xa0, 0x4f, 0x11,
	0xf6, 0x5d, 0xf0, 0x36, 0x58, 0x19, 0x90, 0x6e, 0x9f, 0xf2, 0xb9, 0xb2, 0x3e, 0xbd, 0x29, 0x05,
	0x66, 0x84, 0x99, 0x3b, 0xfb, 0xd3, 0x12, 0x10, 0xa2, 0x21, 0x0b, 0x3f, 0x9b, 0x3b, 0xcb, 0xee,
	0x24, 0xcc, 0xb2, 0xb9, 0xa3, 0xe9, 0xfc, 0x7c, 0xcb, 0x81, 0x54, 0xdf, 0xa1, 0x76, 0x8f, 0x9c,
	0x86, 0x4a, 0x62, 0xb3, 0x35, 0xf4, 0x20, 0x1c, 0x81, 0x7c, 0x82, 0x45, 0x1c, 0xe7, 0xcc, 0xb4,
	0x75, 0x3e, 0xd2, 0x62, 0x84, 0xd0, 0x83, 0x70, 0x04, 0x7a, 0xf5, 0xed, 0x13, 0xad, 0xfe, 0xf9,
	0xb6, 0x0f, 0x3d, 0xb7, 0x7b, 0xfa, 0x96, 0x4e, 0x5c, 0x3a, 0xbb, 0x7b, 0x98, 0x1d, 0x61, 0x0e,
	0xf0, 0x97, 0x65, 0x9b, 0xdd, 0xae, 0x7f, 0x37, 0x9a, 0xad, 0x43, 0xe8, 0x41, 0x38, 0x02, 0x65,
	0xff, 0x58, 0x00, 0xa9, 0xf0, 0x3e, 0xf2, 0x7f, 0x6c, 0xcb, 0xf6, 0xf9, 0xb6, 0x08, 0x17, 0x16,
	0x15, 0x3d, 0x00, 0xef, 0x26, 0x75, 0xcb, 0xb1, 0xcc, 0x9e, 0x43, 0xfd, 0x5a, 0x07, 0x17, 0x64,
	0x7d, 0xb6, 0xd6, 0xcc, 0x8e, 0x30, 0x07, 0xa0, 0xbf, 0x17, 0xc0, 0xe6, 0xc5, 0xd1, 0xbe, 0xb6,
	0x89, 0x65, 0x51, 0x1b, 0xbe, 0x07, 0x96, 0x35, 0x53, 0x67, 0xed, 0x5b, 0xa9, 0x5e, 0x9f, 0x8c,
	0xc5, 0x34, 0xbf, 0xf5, 0x99, 0x3a, 0x45, 0x38, 0x70, 0xc2, 0xbb, 0x20, 0xed, 0xff, 0x7e, 0xf2,
	0x8d, 0xd5, 0x25, 0x9d, 0x1e, 0x2f, 0xd7, 0xdb, 0xd3, 0x1b, 0x78, 0xcc, 0x89, 0x70, 0x1c, 0xea,
	0xff, 0x07, 0xa9, 0x6d, 0x9b, 0x36, 0xaf, 0x58, 0xec, 0x3f, 0x18, 0x98, 0x11, 0x66, 0x6e, 0xf8,
	0x00, 0x2c, 0xeb, 0xc4, 0x25, 0x99, 0xe5, 0xe0, 0x6b, 0xe2, 0xf6, 0xb3, 0xf6, 0x2f, 0x5b, 0x43,
	0x5c, 0xae, 0xcf, 0x46, 0x38, 0x08, 0xd2, 0x58, 0x0d, 0xbe, 0xa5, 0x8a, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x62, 0x0f, 0xc5, 0xd8, 0x0d, 0x00, 0x00,
}
