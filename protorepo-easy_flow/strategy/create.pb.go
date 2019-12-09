// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package strategy

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
	easy_flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_flow"
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
//Create请求
type CreateRequest struct {
	//
	//策略Id，策略Id自动生成
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//策略名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//api版本
	ApiVersion string `protobuf:"bytes,3,opt,name=apiVersion,proto3" json:"apiVersion" form:"apiVersion"`
	//
	//组织
	Org int32 `protobuf:"varint,4,opt,name=org,proto3" json:"org" form:"org"`
	//
	//应用相关信息
	App *CreateRequest_App `protobuf:"bytes,5,opt,name=app,proto3" json:"app" form:"app"`
	//
	//策略类型
	Type string `protobuf:"bytes,6,opt,name=type,proto3" json:"type" form:"type"`
	//
	//批量策略信息
	BatchStrategy *CreateRequest_BatchStrategy `protobuf:"bytes,7,opt,name=batchStrategy,proto3" json:"batchStrategy" form:"batchStrategy"`
	//
	//范围
	Scope string `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope" form:"scope"`
	//
	//集群列表
	Clusters []*cmdb.ClusterInfo `protobuf:"bytes,9,rep,name=clusters,proto3" json:"clusters" form:"clusters"`
	//
	//目标设备列表
	TargetList []*easy_flow.TargetInfo `protobuf:"bytes,10,rep,name=targetList,proto3" json:"targetList" form:"targetList"`
	//
	//集群环境
	ClusterEnvironment string `protobuf:"bytes,11,opt,name=clusterEnvironment,proto3" json:"clusterEnvironment" form:"clusterEnvironment"`
	//
	//集群类型
	ClusterType string `protobuf:"bytes,12,opt,name=clusterType,proto3" json:"clusterType" form:"clusterType"`
	//
	//部署包信息
	PackageList *CreateRequest_PackageList `protobuf:"bytes,13,opt,name=packageList,proto3" json:"packageList" form:"packageList"`
	//
	//部署策略状态
	Status               *CreateRequest_Status `protobuf:"bytes,14,opt,name=status,proto3" json:"status" form:"status"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *CreateRequest) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *CreateRequest) GetApp() *CreateRequest_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *CreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest) GetBatchStrategy() *CreateRequest_BatchStrategy {
	if m != nil {
		return m.BatchStrategy
	}
	return nil
}

func (m *CreateRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *CreateRequest) GetClusters() []*cmdb.ClusterInfo {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *CreateRequest) GetTargetList() []*easy_flow.TargetInfo {
	if m != nil {
		return m.TargetList
	}
	return nil
}

func (m *CreateRequest) GetClusterEnvironment() string {
	if m != nil {
		return m.ClusterEnvironment
	}
	return ""
}

func (m *CreateRequest) GetClusterType() string {
	if m != nil {
		return m.ClusterType
	}
	return ""
}

func (m *CreateRequest) GetPackageList() *CreateRequest_PackageList {
	if m != nil {
		return m.PackageList
	}
	return nil
}

func (m *CreateRequest) GetStatus() *CreateRequest_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type CreateRequest_App struct {
	//
	//应用名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//应用Id
	AppId                string   `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId" form:"appId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_App) Reset()         { *m = CreateRequest_App{} }
func (m *CreateRequest_App) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_App) ProtoMessage()    {}
func (*CreateRequest_App) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 0}
}
func (m *CreateRequest_App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_App.Unmarshal(m, b)
}
func (m *CreateRequest_App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_App.Marshal(b, m, deterministic)
}
func (m *CreateRequest_App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_App.Merge(m, src)
}
func (m *CreateRequest_App) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_App.Size(m)
}
func (m *CreateRequest_App) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_App.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_App proto.InternalMessageInfo

func (m *CreateRequest_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest_App) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

type CreateRequest_BatchStrategy struct {
	//
	//自动批量策略信息
	AutoBatch *CreateRequest_BatchStrategy_AutoBatch `protobuf:"bytes,1,opt,name=autoBatch,proto3" json:"autoBatch" form:"autoBatch"`
	//
	//手动批量
	ManualBatch *CreateRequest_BatchStrategy_ManualBatch `protobuf:"bytes,2,opt,name=manualBatch,proto3" json:"manualBatch" form:"manualBatch"`
	//
	//分批类型
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_BatchStrategy) Reset()         { *m = CreateRequest_BatchStrategy{} }
func (m *CreateRequest_BatchStrategy) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_BatchStrategy) ProtoMessage()    {}
func (*CreateRequest_BatchStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 1}
}
func (m *CreateRequest_BatchStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_BatchStrategy.Unmarshal(m, b)
}
func (m *CreateRequest_BatchStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_BatchStrategy.Marshal(b, m, deterministic)
}
func (m *CreateRequest_BatchStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_BatchStrategy.Merge(m, src)
}
func (m *CreateRequest_BatchStrategy) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_BatchStrategy.Size(m)
}
func (m *CreateRequest_BatchStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_BatchStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_BatchStrategy proto.InternalMessageInfo

func (m *CreateRequest_BatchStrategy) GetAutoBatch() *CreateRequest_BatchStrategy_AutoBatch {
	if m != nil {
		return m.AutoBatch
	}
	return nil
}

func (m *CreateRequest_BatchStrategy) GetManualBatch() *CreateRequest_BatchStrategy_ManualBatch {
	if m != nil {
		return m.ManualBatch
	}
	return nil
}

func (m *CreateRequest_BatchStrategy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type CreateRequest_BatchStrategy_AutoBatch struct {
	//
	//自动分批:每批次部署机器数;手动分批:部署总共批次
	BatchNum int32 `protobuf:"varint,1,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//部署间隔时长(s)
	BatchInterval int32 `protobuf:"varint,2,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	//
	//是否失败暂停
	FailedStop           string   `protobuf:"bytes,3,opt,name=failedStop,proto3" json:"failedStop" form:"failedStop"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_BatchStrategy_AutoBatch) Reset()         { *m = CreateRequest_BatchStrategy_AutoBatch{} }
func (m *CreateRequest_BatchStrategy_AutoBatch) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_BatchStrategy_AutoBatch) ProtoMessage()    {}
func (*CreateRequest_BatchStrategy_AutoBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 1, 0}
}
func (m *CreateRequest_BatchStrategy_AutoBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch.Unmarshal(m, b)
}
func (m *CreateRequest_BatchStrategy_AutoBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch.Marshal(b, m, deterministic)
}
func (m *CreateRequest_BatchStrategy_AutoBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch.Merge(m, src)
}
func (m *CreateRequest_BatchStrategy_AutoBatch) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch.Size(m)
}
func (m *CreateRequest_BatchStrategy_AutoBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_BatchStrategy_AutoBatch proto.InternalMessageInfo

func (m *CreateRequest_BatchStrategy_AutoBatch) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *CreateRequest_BatchStrategy_AutoBatch) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

func (m *CreateRequest_BatchStrategy_AutoBatch) GetFailedStop() string {
	if m != nil {
		return m.FailedStop
	}
	return ""
}

type CreateRequest_BatchStrategy_ManualBatch struct {
	//
	//分批详情
	Batches []*CreateRequest_BatchStrategy_ManualBatch_Batches `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches" form:"batches"`
	//
	//自动分批:每批次部署机器数;手动分批:部署总共批次
	BatchNum int32 `protobuf:"varint,2,opt,name=batchNum,proto3" json:"batchNum" form:"batchNum"`
	//
	//部署间隔时长(s)
	BatchInterval int32 `protobuf:"varint,3,opt,name=batchInterval,proto3" json:"batchInterval" form:"batchInterval"`
	//
	//是否失败暂停
	FailedStop           string   `protobuf:"bytes,4,opt,name=failedStop,proto3" json:"failedStop" form:"failedStop"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_BatchStrategy_ManualBatch) Reset() {
	*m = CreateRequest_BatchStrategy_ManualBatch{}
}
func (m *CreateRequest_BatchStrategy_ManualBatch) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_BatchStrategy_ManualBatch) ProtoMessage()    {}
func (*CreateRequest_BatchStrategy_ManualBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 1, 1}
}
func (m *CreateRequest_BatchStrategy_ManualBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch.Unmarshal(m, b)
}
func (m *CreateRequest_BatchStrategy_ManualBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch.Marshal(b, m, deterministic)
}
func (m *CreateRequest_BatchStrategy_ManualBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch.Merge(m, src)
}
func (m *CreateRequest_BatchStrategy_ManualBatch) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch.Size(m)
}
func (m *CreateRequest_BatchStrategy_ManualBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch proto.InternalMessageInfo

func (m *CreateRequest_BatchStrategy_ManualBatch) GetBatches() []*CreateRequest_BatchStrategy_ManualBatch_Batches {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *CreateRequest_BatchStrategy_ManualBatch) GetBatchNum() int32 {
	if m != nil {
		return m.BatchNum
	}
	return 0
}

func (m *CreateRequest_BatchStrategy_ManualBatch) GetBatchInterval() int32 {
	if m != nil {
		return m.BatchInterval
	}
	return 0
}

func (m *CreateRequest_BatchStrategy_ManualBatch) GetFailedStop() string {
	if m != nil {
		return m.FailedStop
	}
	return ""
}

type CreateRequest_BatchStrategy_ManualBatch_Batches struct {
	//
	//批次主机列表
	Targets              []*easy_flow.DeployBatchTarget `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets" form:"targets"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) Reset() {
	*m = CreateRequest_BatchStrategy_ManualBatch_Batches{}
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) String() string {
	return proto.CompactTextString(m)
}
func (*CreateRequest_BatchStrategy_ManualBatch_Batches) ProtoMessage() {}
func (*CreateRequest_BatchStrategy_ManualBatch_Batches) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 1, 1, 0}
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches.Unmarshal(m, b)
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches.Marshal(b, m, deterministic)
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches.Merge(m, src)
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches.Size(m)
}
func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_BatchStrategy_ManualBatch_Batches proto.InternalMessageInfo

func (m *CreateRequest_BatchStrategy_ManualBatch_Batches) GetTargets() []*easy_flow.DeployBatchTarget {
	if m != nil {
		return m.Targets
	}
	return nil
}

type CreateRequest_PackageList struct {
	//
	//集群
	Cluster *cmdb.ClusterInfo `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster" form:"cluster"`
	//
	//目标版本
	TargetVersion string `protobuf:"bytes,2,opt,name=targetVersion,proto3" json:"targetVersion" form:"targetVersion"`
	//
	//是否提前暂停
	PreStop string `protobuf:"bytes,3,opt,name=preStop,proto3" json:"preStop" form:"preStop"`
	//
	//是否发布重启
	PostRestart string `protobuf:"bytes,4,opt,name=postRestart,proto3" json:"postRestart" form:"postRestart"`
	//
	//是否自动运行
	AutoStart string `protobuf:"bytes,5,opt,name=autoStart,proto3" json:"autoStart" form:"autoStart"`
	//
	//是否需要用户检查
	UserCheck string `protobuf:"bytes,6,opt,name=userCheck,proto3" json:"userCheck" form:"userCheck"`
	//
	//是否完全更新
	FullUpdate string `protobuf:"bytes,7,opt,name=fullUpdate,proto3" json:"fullUpdate" form:"fullUpdate"`
	//
	//强制更新
	Force string `protobuf:"bytes,8,opt,name=force,proto3" json:"force" form:"force"`
	//
	//强制安装
	ForceInstall string `protobuf:"bytes,9,opt,name=forceInstall,proto3" json:"forceInstall" form:"forceInstall"`
	//
	//包Id
	PackageId string `protobuf:"bytes,10,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//包名称
	Name string `protobuf:"bytes,11,opt,name=name,proto3" json:"name" form:"name"`
	//
	//安装路径
	InstallPath string `protobuf:"bytes,12,opt,name=installPath,proto3" json:"installPath" form:"installPath"`
	//
	//版本类型 1 程序包,  2 配置包,  4 文件包
	Type int32 `protobuf:"varint,13,opt,name=type,proto3" json:"type" form:"type"`
	//
	//平台
	Platform             string   `protobuf:"bytes,14,opt,name=platform,proto3" json:"platform" form:"platform"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_PackageList) Reset()         { *m = CreateRequest_PackageList{} }
func (m *CreateRequest_PackageList) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_PackageList) ProtoMessage()    {}
func (*CreateRequest_PackageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 2}
}
func (m *CreateRequest_PackageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_PackageList.Unmarshal(m, b)
}
func (m *CreateRequest_PackageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_PackageList.Marshal(b, m, deterministic)
}
func (m *CreateRequest_PackageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_PackageList.Merge(m, src)
}
func (m *CreateRequest_PackageList) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_PackageList.Size(m)
}
func (m *CreateRequest_PackageList) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_PackageList.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_PackageList proto.InternalMessageInfo

func (m *CreateRequest_PackageList) GetCluster() *cmdb.ClusterInfo {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *CreateRequest_PackageList) GetTargetVersion() string {
	if m != nil {
		return m.TargetVersion
	}
	return ""
}

func (m *CreateRequest_PackageList) GetPreStop() string {
	if m != nil {
		return m.PreStop
	}
	return ""
}

func (m *CreateRequest_PackageList) GetPostRestart() string {
	if m != nil {
		return m.PostRestart
	}
	return ""
}

func (m *CreateRequest_PackageList) GetAutoStart() string {
	if m != nil {
		return m.AutoStart
	}
	return ""
}

func (m *CreateRequest_PackageList) GetUserCheck() string {
	if m != nil {
		return m.UserCheck
	}
	return ""
}

func (m *CreateRequest_PackageList) GetFullUpdate() string {
	if m != nil {
		return m.FullUpdate
	}
	return ""
}

func (m *CreateRequest_PackageList) GetForce() string {
	if m != nil {
		return m.Force
	}
	return ""
}

func (m *CreateRequest_PackageList) GetForceInstall() string {
	if m != nil {
		return m.ForceInstall
	}
	return ""
}

func (m *CreateRequest_PackageList) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *CreateRequest_PackageList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest_PackageList) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *CreateRequest_PackageList) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CreateRequest_PackageList) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type CreateRequest_Status struct {
	//
	//是否过时
	OutOfDate            string   `protobuf:"bytes,1,opt,name=outOfDate,proto3" json:"outOfDate" form:"outOfDate"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest_Status) Reset()         { *m = CreateRequest_Status{} }
func (m *CreateRequest_Status) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Status) ProtoMessage()    {}
func (*CreateRequest_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 3}
}
func (m *CreateRequest_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest_Status.Unmarshal(m, b)
}
func (m *CreateRequest_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest_Status.Marshal(b, m, deterministic)
}
func (m *CreateRequest_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest_Status.Merge(m, src)
}
func (m *CreateRequest_Status) XXX_Size() int {
	return xxx_messageInfo_CreateRequest_Status.Size(m)
}
func (m *CreateRequest_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest_Status.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest_Status proto.InternalMessageInfo

func (m *CreateRequest_Status) GetOutOfDate() string {
	if m != nil {
		return m.OutOfDate
	}
	return ""
}

//
//CreateApi返回
type CreateResponseWrapper struct {
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
	Data                 *easy_flow.DeployStrategy `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateResponseWrapper) Reset()         { *m = CreateResponseWrapper{} }
func (m *CreateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateResponseWrapper) ProtoMessage()    {}
func (*CreateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseWrapper.Unmarshal(m, b)
}
func (m *CreateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseWrapper.Merge(m, src)
}
func (m *CreateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateResponseWrapper.Size(m)
}
func (m *CreateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseWrapper proto.InternalMessageInfo

func (m *CreateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponseWrapper) GetData() *easy_flow.DeployStrategy {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "strategy.CreateRequest")
	proto.RegisterType((*CreateRequest_App)(nil), "strategy.CreateRequest.App")
	proto.RegisterType((*CreateRequest_BatchStrategy)(nil), "strategy.CreateRequest.BatchStrategy")
	proto.RegisterType((*CreateRequest_BatchStrategy_AutoBatch)(nil), "strategy.CreateRequest.BatchStrategy.AutoBatch")
	proto.RegisterType((*CreateRequest_BatchStrategy_ManualBatch)(nil), "strategy.CreateRequest.BatchStrategy.ManualBatch")
	proto.RegisterType((*CreateRequest_BatchStrategy_ManualBatch_Batches)(nil), "strategy.CreateRequest.BatchStrategy.ManualBatch.Batches")
	proto.RegisterType((*CreateRequest_PackageList)(nil), "strategy.CreateRequest.PackageList")
	proto.RegisterType((*CreateRequest_Status)(nil), "strategy.CreateRequest.Status")
	proto.RegisterType((*CreateResponseWrapper)(nil), "strategy.CreateResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 1348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdb, 0x72, 0x13, 0x37,
	0x18, 0xc6, 0x49, 0x1c, 0xc7, 0x72, 0x12, 0x82, 0x38, 0x74, 0x71, 0x0f, 0x9b, 0x11, 0x29, 0x63,
	0xa0, 0xf6, 0x42, 0x68, 0x3b, 0x90, 0x52, 0x98, 0x18, 0x68, 0xc7, 0x1d, 0x68, 0x41, 0x81, 0x76,
	0xda, 0x38, 0xc9, 0x28, 0xbb, 0xb2, 0xb3, 0x93, 0xb5, 0xb5, 0xdd, 0x95, 0x81, 0x70, 0xb8, 0xec,
	0x45, 0x2f, 0xfb, 0x2a, 0x6d, 0x9f, 0xc7, 0xcc, 0xf0, 0x08, 0x9e, 0xe9, 0x7d, 0x67, 0x7f, 0x69,
	0x77, 0x65, 0x42, 0x18, 0xa0, 0xf4, 0xce, 0xd2, 0xf7, 0xfd, 0x9f, 0xa4, 0xff, 0xb8, 0x46, 0xb3,
	0x6e, 0xc4, 0x99, 0xe4, 0x8d, 0x30, 0x12, 0x52, 0xe0, 0x99, 0x58, 0x46, 0x4c, 0xf2, 0xee, 0x5e,
	0xb5, 0xde, 0xf5, 0xe5, 0xce, 0x60, 0xbb, 0xe1, 0x8a, 0x9e, 0xd3, 0x15, 0x5d, 0xe1, 0x00, 0x61,
	0x7b, 0xd0, 0x81, 0x15, 0x2c, 0xe0, 0x97, 0x32, 0xac, 0x6e, 0x74, 0x45, 0x83, 0xb3, 0x78, 0x4f,
	0x84, 0x71, 0x23, 0x10, 0x2e, 0x0b, 0x1c, 0x57, 0xf4, 0x65, 0xc4, 0x5c, 0x19, 0x2b, 0xcb, 0x88,
	0x87, 0xa2, 0xde, 0x13, 0x1e, 0x0f, 0x62, 0x47, 0x13, 0x1d, 0x58, 0xc2, 0x6a, 0xab, 0x13, 0x88,
	0x87, 0x8e, 0xc7, 0xc3, 0x40, 0xec, 0x6d, 0x6d, 0x33, 0xe9, 0xee, 0x6c, 0x49, 0x16, 0x75, 0xb9,
	0xd4, 0xf2, 0x77, 0xdf, 0x5d, 0xde, 0xed, 0x79, 0xdb, 0x8e, 0x1b, 0x0c, 0x62, 0xc9, 0xa3, 0x2d,
	0xbf, 0xdf, 0x49, 0x6f, 0x7c, 0xff, 0x7d, 0xdc, 0x58, 0x5d, 0xd2, 0x94, 0xfd, 0xf9, 0x3d, 0x3a,
	0x22, 0x0d, 0x86, 0x96, 0xfe, 0xd2, 0x08, 0x49, 0xef, 0xa1, 0x2f, 0x77, 0xc5, 0x43, 0xa7, 0x2b,
	0xea, 0x00, 0xd6, 0x1f, 0xb0, 0xc0, 0xf7, 0x98, 0x14, 0x51, 0xec, 0x64, 0x3f, 0x95, 0x1d, 0xf9,
	0xf3, 0x04, 0x9a, 0xbb, 0x0e, 0x51, 0xa6, 0xfc, 0xd7, 0x01, 0x8f, 0x25, 0x96, 0x68, 0xc2, 0xf7,
	0xac, 0xc2, 0x62, 0xa1, 0x56, 0x6e, 0x7a, 0xa3, 0xa1, 0x5d, 0xee, 0x88, 0xa8, 0xb7, 0x42, 0x7c,
	0x8f, 0xbc, 0x78, 0x6e, 0xaf, 0xa1, 0xbb, 0x9b, 0xeb, 0xac, 0xde, 0x59, 0xad, 0x7f, 0x73, 0xbe,
	0x7e, 0x79, 0xe3, 0xc9, 0xa5, 0x67, 0xf5, 0x6b, 0xe6, 0xfa, 0xf3, 0xb7, 0x5c, 0x5f, 0x58, 0x7e,
	0xb6, 0x44, 0x27, 0x7c, 0x0f, 0x9f, 0x42, 0x53, 0x7d, 0xd6, 0xe3, 0xd6, 0x04, 0x9c, 0x7b, 0x78,
	0x34, 0xb4, 0x2b, 0xea, 0xdc, 0x64, 0x97, 0x50, 0x00, 0xf1, 0x17, 0x08, 0xb1, 0xd0, 0xff, 0x91,
	0x47, 0xb1, 0x2f, 0xfa, 0xd6, 0x24, 0x50, 0x8f, 0x8f, 0x86, 0xf6, 0x11, 0x45, 0xcd, 0x31, 0x42,
	0x0d, 0x22, 0x3e, 0x8d, 0x26, 0x45, 0xd4, 0xb5, 0xa6, 0x16, 0x0b, 0xb5, 0x62, 0xf3, 0xd8, 0x68,
	0x68, 0x23, 0xc5, 0x17, 0x51, 0x37, 0x79, 0xd3, 0xc4, 0xc2, 0x21, 0x9a, 0x10, 0xf0, 0xd7, 0x68,
	0x92, 0x85, 0xa1, 0x55, 0x5c, 0x2c, 0xd4, 0x2a, 0xcb, 0x1f, 0x36, 0x32, 0x0f, 0x8f, 0xf9, 0xa7,
	0xb1, 0x1a, 0x86, 0xcd, 0xf9, 0x5c, 0x84, 0x85, 0x21, 0xa1, 0x89, 0x5d, 0xf2, 0x04, 0xb9, 0x17,
	0x72, 0x6b, 0xfa, 0xe5, 0x27, 0x24, 0xbb, 0x84, 0x02, 0x88, 0x39, 0x9a, 0x83, 0x14, 0x5e, 0xd3,
	0xe2, 0x56, 0x09, 0x4e, 0xfb, 0xf4, 0xa0, 0xd3, 0x9a, 0x26, 0xb9, 0x69, 0x8d, 0x86, 0xf6, 0x31,
	0x25, 0x3a, 0xa6, 0x42, 0xe8, 0xb8, 0x2a, 0x3e, 0x8d, 0x8a, 0xb1, 0x2b, 0x42, 0x6e, 0xcd, 0xc0,
	0x65, 0x16, 0x46, 0x43, 0x7b, 0x56, 0xd9, 0xc1, 0x36, 0xa1, 0x0a, 0xc6, 0x4d, 0x34, 0xa3, 0xd3,
	0x3f, 0xb6, 0xca, 0x8b, 0x93, 0xb5, 0xca, 0xf2, 0x91, 0x46, 0x52, 0x14, 0x8d, 0xeb, 0x6a, 0xb7,
	0xd5, 0xef, 0x88, 0xe6, 0xd1, 0xd1, 0xd0, 0x3e, 0xac, 0xac, 0x53, 0x32, 0xa1, 0x99, 0x1d, 0xbe,
	0x85, 0x90, 0x4a, 0xf5, 0x5b, 0x7e, 0x2c, 0x2d, 0x04, 0x2a, 0xc7, 0x1b, 0x59, 0xc2, 0x36, 0xee,
	0x01, 0x08, 0x4a, 0x46, 0xb0, 0x72, 0x13, 0x42, 0x0d, 0x7b, 0x7c, 0x1b, 0x61, 0xad, 0x7c, 0xb3,
	0xff, 0xc0, 0x8f, 0x44, 0xbf, 0xc7, 0xfb, 0xd2, 0xaa, 0xc0, 0x33, 0x3e, 0x1e, 0x0d, 0xed, 0x93,
	0x63, 0x17, 0x31, 0x38, 0x84, 0xbe, 0xc2, 0x10, 0x5f, 0x42, 0x15, 0xbd, 0x7b, 0x2f, 0x89, 0xcd,
	0x2c, 0xe8, 0x9c, 0x18, 0x0d, 0x6d, 0x3c, 0xa6, 0x73, 0x0f, 0x42, 0x64, 0x52, 0xf1, 0x06, 0xaa,
	0x84, 0xcc, 0xdd, 0x65, 0x5d, 0x0e, 0xef, 0x9a, 0x83, 0x38, 0x9d, 0x3a, 0x28, 0x4e, 0x77, 0x72,
	0xaa, 0x29, 0x6f, 0x28, 0x10, 0x6a, 0xea, 0xe1, 0x16, 0x9a, 0x8e, 0x25, 0x93, 0x83, 0xd8, 0x9a,
	0x07, 0xe5, 0x4f, 0x0e, 0x52, 0x5e, 0x03, 0x56, 0xf3, 0xc8, 0x68, 0x68, 0xcf, 0xe9, 0x10, 0xc2,
	0x0e, 0xa1, 0x5a, 0xa0, 0xda, 0x41, 0x93, 0xab, 0x2a, 0xff, 0xa0, 0x84, 0x0a, 0xaf, 0x2b, 0xa1,
	0x15, 0x54, 0x64, 0x61, 0xd8, 0xf2, 0x74, 0xa1, 0x2d, 0xe5, 0x89, 0x01, 0xdb, 0x49, 0x3d, 0x2c,
	0xa0, 0xf9, 0xcd, 0xf5, 0xf3, 0xf5, 0xcb, 0xac, 0xfe, 0x78, 0xe3, 0xc9, 0x85, 0x8b, 0xcf, 0x96,
	0xa8, 0x32, 0xa9, 0xfe, 0x3d, 0x8d, 0xe6, 0xc6, 0xf2, 0x11, 0xbb, 0xa8, 0xcc, 0x06, 0x52, 0xc0,
	0x26, 0x9c, 0x5b, 0x59, 0x76, 0xde, 0x28, 0x93, 0x1b, 0xab, 0xa9, 0x19, 0x14, 0xe4, 0x82, 0xbe,
	0x42, 0xba, 0x49, 0x68, 0xae, 0x8b, 0x77, 0x51, 0xa5, 0xc7, 0xfa, 0x03, 0x16, 0xa8, 0x63, 0x26,
	0xe0, 0x98, 0x0b, 0x6f, 0x76, 0xcc, 0xed, 0xdc, 0xd0, 0x0c, 0x8b, 0xa1, 0x47, 0xa8, 0xa9, 0x9e,
	0x15, 0xf1, 0xe4, 0x6b, 0x8a, 0xb8, 0xfa, 0x57, 0x01, 0x95, 0xb3, 0x07, 0x60, 0x07, 0xcd, 0x40,
	0xf1, 0x7d, 0x3f, 0xe8, 0x81, 0x0f, 0x8a, 0x66, 0xc1, 0xa4, 0x08, 0xa1, 0x19, 0x09, 0x5f, 0xd5,
	0x3d, 0xa0, 0xd5, 0x97, 0x3c, 0x7a, 0xc0, 0x02, 0x78, 0x52, 0x71, 0x5f, 0x71, 0xa7, 0x70, 0x5a,
	0xdc, 0xe9, 0x3a, 0x69, 0x83, 0x1d, 0xe6, 0x07, 0xdc, 0x5b, 0x93, 0x22, 0xdc, 0xdf, 0x06, 0x73,
	0x8c, 0x50, 0x83, 0x58, 0xfd, 0x67, 0x02, 0x55, 0x0c, 0x7f, 0xe0, 0x2e, 0x2a, 0x81, 0x2e, 0x8f,
	0xad, 0x02, 0x14, 0xed, 0xe5, 0xb7, 0xf6, 0xa9, 0x42, 0x78, 0xdc, 0xc4, 0xa3, 0xa1, 0x3d, 0x6f,
	0xdc, 0x9d, 0xc7, 0x84, 0xa6, 0xea, 0x63, 0x0e, 0x9a, 0x78, 0x27, 0x07, 0x4d, 0xfe, 0x17, 0x07,
	0x4d, 0xbd, 0xa9, 0x83, 0xee, 0xa3, 0x92, 0x7e, 0x0f, 0xfe, 0x0e, 0x95, 0x54, 0x4f, 0x4a, 0x7d,
	0xf3, 0x91, 0xd1, 0xd0, 0x6e, 0xc0, 0x04, 0x06, 0xaa, 0xea, 0x6d, 0xe6, 0xf3, 0xb5, 0x19, 0xa1,
	0xa9, 0x40, 0xf5, 0x8f, 0x12, 0xaa, 0x18, 0xed, 0x01, 0x5f, 0x43, 0x25, 0xdd, 0x67, 0x74, 0xc9,
	0xbc, 0xa2, 0xe5, 0x1a, 0x82, 0x9a, 0x4b, 0x68, 0x6a, 0x95, 0xb8, 0x47, 0x69, 0xa7, 0x93, 0x50,
	0xd5, 0xb2, 0xe1, 0x9e, 0x31, 0x98, 0xd0, 0x71, 0x3a, 0xfe, 0x0c, 0x95, 0xc2, 0x88, 0x1b, 0xc9,
	0x63, 0x9c, 0xa6, 0x01, 0x42, 0x53, 0x4a, 0xd2, 0x41, 0x43, 0x11, 0x4b, 0xca, 0x63, 0xc9, 0x22,
	0xa9, 0xbd, 0x69, 0xb6, 0xb8, 0x1c, 0x4c, 0x5a, 0x5c, 0xbe, 0xc2, 0xcb, 0xaa, 0x3b, 0xac, 0x81,
	0x5d, 0x11, 0xec, 0x5e, 0x2a, 0xf6, 0x35, 0x65, 0x95, 0xd3, 0x12, 0x9b, 0x41, 0xcc, 0xa3, 0xeb,
	0x3b, 0xdc, 0xdd, 0xd5, 0x93, 0xd4, 0xb0, 0xc9, 0x20, 0x42, 0x73, 0x1a, 0x84, 0x7b, 0x10, 0x04,
	0xf7, 0x43, 0x8f, 0x49, 0x0e, 0x03, 0x75, 0x3c, 0xdc, 0x19, 0x96, 0x84, 0x3b, 0x5b, 0x24, 0x33,
	0xb2, 0x23, 0x22, 0xf7, 0x15, 0x33, 0x12, 0xb6, 0x09, 0x55, 0x30, 0xfe, 0x0a, 0xcd, 0xc2, 0x8f,
	0x56, 0x3f, 0x96, 0x2c, 0x08, 0xac, 0x32, 0xd0, 0x3f, 0x18, 0x0d, 0xed, 0xa3, 0x06, 0x5d, 0xa3,
	0x84, 0x8e, 0x91, 0xf1, 0xef, 0x05, 0x54, 0xd6, 0x6d, 0xbf, 0xe5, 0x59, 0x08, 0x4c, 0x77, 0xf3,
	0x07, 0x65, 0xd0, 0xff, 0xf6, 0x71, 0x95, 0x9f, 0x8e, 0xcf, 0xe9, 0x01, 0x51, 0xc9, 0x1e, 0x60,
	0x0e, 0x88, 0x17, 0xcf, 0xed, 0xa9, 0xf0, 0xd0, 0xa3, 0x86, 0x1e, 0x14, 0xbf, 0x15, 0x50, 0xc5,
	0x57, 0x8f, 0xb8, 0xc3, 0xe4, 0x8e, 0x9e, 0x9c, 0x6e, 0x1e, 0x77, 0x03, 0x4c, 0x6c, 0x5b, 0xe8,
	0xdb, 0xcd, 0x5a, 0xad, 0xed, 0xac, 0x6f, 0xb6, 0x9d, 0x95, 0xf6, 0xd9, 0xf6, 0x35, 0x42, 0xae,
	0x5c, 0x6d, 0x3f, 0x6d, 0x47, 0xed, 0xfe, 0xc6, 0xb9, 0x33, 0xe7, 0xce, 0x3c, 0xad, 0xb5, 0x9d,
	0x33, 0x4f, 0xd7, 0x59, 0xfd, 0xf1, 0x6a, 0xfd, 0x97, 0x8d, 0x95, 0x5a, 0xbb, 0xbd, 0xbe, 0xd9,
	0x6e, 0xef, 0x67, 0x9e, 0x5d, 0xa2, 0xe6, 0xb9, 0x59, 0x43, 0x9e, 0x83, 0x16, 0x70, 0xc0, 0x57,
	0x95, 0x83, 0x66, 0xc2, 0x80, 0xc9, 0x04, 0x80, 0x71, 0x5a, 0x36, 0x3b, 0x4c, 0x8a, 0x10, 0x9a,
	0x91, 0xaa, 0x57, 0xd0, 0xb4, 0x9a, 0xab, 0x49, 0xc2, 0x89, 0x81, 0xfc, 0xa1, 0x73, 0x23, 0xc9,
	0x9d, 0xc2, 0xcb, 0x09, 0x97, 0x41, 0x84, 0xe6, 0x34, 0x32, 0x2c, 0xa0, 0xe3, 0x69, 0x87, 0x8c,
	0x43, 0xd1, 0x8f, 0xf9, 0x4f, 0x11, 0x0b, 0x43, 0x1e, 0x25, 0xb7, 0x75, 0x85, 0xc7, 0xf5, 0x1c,
	0x30, 0x6e, 0x9b, 0xec, 0x12, 0x0a, 0x20, 0x7c, 0x93, 0x08, 0x8f, 0xdf, 0x7c, 0x14, 0x06, 0xcc,
	0x4f, 0xab, 0xd7, 0xfc, 0x26, 0xc9, 0xc1, 0xe4, 0x9b, 0x24, 0x5f, 0x25, 0x29, 0xcb, 0xa3, 0x48,
	0x44, 0xba, 0x6e, 0x8d, 0x94, 0x85, 0x6d, 0x42, 0x15, 0x8c, 0xaf, 0xa2, 0x29, 0x8f, 0x49, 0x06,
	0xc5, 0x5a, 0x59, 0x3e, 0xb9, 0xaf, 0x77, 0x65, 0x1f, 0x94, 0xc6, 0x0d, 0x13, 0x03, 0x42, 0xc1,
	0x6e, 0x7b, 0x1a, 0xfe, 0x1c, 0x5c, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x60, 0x6d, 0x5d, 0xc1,
	0x01, 0x0e, 0x00, 0x00,
}
