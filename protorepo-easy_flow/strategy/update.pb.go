// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Update请求
type UpdateRequest struct {
	//
	//策略ID
	StrategyID string `protobuf:"bytes,1,opt,name=strategyID,proto3" json:"strategyID" form:"strategyID"`
	//
	//策略名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//api版本
	ApiVersion string `protobuf:"bytes,3,opt,name=apiVersion,proto3" json:"apiVersion" form:"apiVersion"`
	//
	//集群列表
	Clusters []*cmdb.ClusterInfo `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters" form:"clusters"`
	//
	//部署包信息
	PackageList          *UpdateRequest_PackageList `protobuf:"bytes,5,opt,name=packageList,proto3" json:"packageList" form:"packageList"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{0}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetStrategyID() string {
	if m != nil {
		return m.StrategyID
	}
	return ""
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *UpdateRequest) GetClusters() []*cmdb.ClusterInfo {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *UpdateRequest) GetPackageList() *UpdateRequest_PackageList {
	if m != nil {
		return m.PackageList
	}
	return nil
}

type UpdateRequest_PackageList struct {
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

func (m *UpdateRequest_PackageList) Reset()         { *m = UpdateRequest_PackageList{} }
func (m *UpdateRequest_PackageList) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest_PackageList) ProtoMessage()    {}
func (*UpdateRequest_PackageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{0, 0}
}
func (m *UpdateRequest_PackageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest_PackageList.Unmarshal(m, b)
}
func (m *UpdateRequest_PackageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest_PackageList.Marshal(b, m, deterministic)
}
func (m *UpdateRequest_PackageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest_PackageList.Merge(m, src)
}
func (m *UpdateRequest_PackageList) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest_PackageList.Size(m)
}
func (m *UpdateRequest_PackageList) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest_PackageList.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest_PackageList proto.InternalMessageInfo

func (m *UpdateRequest_PackageList) GetCluster() *cmdb.ClusterInfo {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *UpdateRequest_PackageList) GetTargetVersion() string {
	if m != nil {
		return m.TargetVersion
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetPreStop() string {
	if m != nil {
		return m.PreStop
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetPostRestart() string {
	if m != nil {
		return m.PostRestart
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetAutoStart() string {
	if m != nil {
		return m.AutoStart
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetUserCheck() string {
	if m != nil {
		return m.UserCheck
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetFullUpdate() string {
	if m != nil {
		return m.FullUpdate
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetForce() string {
	if m != nil {
		return m.Force
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetForceInstall() string {
	if m != nil {
		return m.ForceInstall
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetInstallPath() string {
	if m != nil {
		return m.InstallPath
	}
	return ""
}

func (m *UpdateRequest_PackageList) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UpdateRequest_PackageList) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

//
//UpdateApi返回
type UpdateResponseWrapper struct {
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

func (m *UpdateResponseWrapper) Reset()         { *m = UpdateResponseWrapper{} }
func (m *UpdateResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseWrapper) ProtoMessage()    {}
func (*UpdateResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f0fa214029f1c21, []int{1}
}
func (m *UpdateResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseWrapper.Unmarshal(m, b)
}
func (m *UpdateResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseWrapper.Marshal(b, m, deterministic)
}
func (m *UpdateResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseWrapper.Merge(m, src)
}
func (m *UpdateResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseWrapper.Size(m)
}
func (m *UpdateResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseWrapper proto.InternalMessageInfo

func (m *UpdateResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *UpdateResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UpdateResponseWrapper) GetData() *easy_flow.DeployStrategy {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "strategy.UpdateRequest")
	proto.RegisterType((*UpdateRequest_PackageList)(nil), "strategy.UpdateRequest.PackageList")
	proto.RegisterType((*UpdateResponseWrapper)(nil), "strategy.UpdateResponseWrapper")
}

func init() { proto.RegisterFile("update.proto", fileDescriptor_3f0fa214029f1c21) }

var fileDescriptor_3f0fa214029f1c21 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0x66, 0xe9, 0xa6, 0x49, 0x3c, 0x49, 0x7f, 0xdc, 0x16, 0x86, 0xdc, 0x4c, 0xe4, 0x56, 0x68,
	0xd3, 0xb0, 0x33, 0xb0, 0xfc, 0xa8, 0x04, 0x94, 0x88, 0xb4, 0x08, 0xad, 0xc4, 0x45, 0x71, 0x28,
	0x08, 0x32, 0x9b, 0xc8, 0x3b, 0xeb, 0xdd, 0x8c, 0x32, 0xbb, 0x36, 0xb6, 0x97, 0x74, 0x21, 0xbd,
	0xe4, 0x06, 0x89, 0x1b, 0x5e, 0x85, 0xf7, 0x99, 0x4a, 0x7d, 0x84, 0x79, 0x02, 0x64, 0x7b, 0x7e,
	0xbc, 0x85, 0x1b, 0x50, 0x7b, 0x35, 0x73, 0xfc, 0x7d, 0xdf, 0xf1, 0xb1, 0xcf, 0xf9, 0x66, 0xc0,
	0xc6, 0x9c, 0x8f, 0x88, 0xa2, 0x21, 0x17, 0x4c, 0x31, 0xb8, 0x26, 0x95, 0x20, 0x8a, 0x4e, 0x16,
	0x5b, 0xdd, 0x49, 0xaa, 0xce, 0xe6, 0xc3, 0x30, 0x61, 0xd3, 0x68, 0xc2, 0x26, 0x2c, 0x32, 0x84,
	0xe1, 0x7c, 0x6c, 0x22, 0x13, 0x98, 0x37, 0x2b, 0xdc, 0xfa, 0x66, 0xc2, 0x42, 0x4a, 0xe4, 0x82,
	0x71, 0x19, 0x66, 0x2c, 0x21, 0x59, 0x94, 0xb0, 0x99, 0x12, 0x24, 0x51, 0xd2, 0x2a, 0x05, 0xe5,
	0xac, 0x3b, 0x65, 0x23, 0x9a, 0xc9, 0xa8, 0x24, 0x46, 0x26, 0x8c, 0x92, 0xe9, 0x68, 0x18, 0x25,
	0xd9, 0x5c, 0x2a, 0x2a, 0x4e, 0xd3, 0xd9, 0xb8, 0x4a, 0x39, 0xf8, 0xff, 0x29, 0x75, 0x74, 0x3a,
	0xce, 0xd8, 0x45, 0x34, 0xa2, 0x3c, 0x63, 0x8b, 0xd3, 0x21, 0x51, 0xc9, 0xd9, 0xa9, 0x22, 0x62,
	0x42, 0x55, 0x99, 0xfe, 0xc9, 0xab, 0x48, 0x6f, 0x33, 0xba, 0x55, 0xff, 0xf0, 0x0a, 0xab, 0xae,
	0x9a, 0x51, 0xa6, 0xfe, 0xc4, 0x69, 0xc9, 0xf4, 0x22, 0x55, 0xe7, 0xec, 0x22, 0x9a, 0xb0, 0xae,
	0x01, 0xbb, 0x3f, 0x93, 0x2c, 0x1d, 0x11, 0xc5, 0x84, 0x8c, 0xea, 0x57, 0xab, 0x43, 0x7f, 0x01,
	0xb0, 0xf9, 0xc4, 0x74, 0x19, 0xd3, 0x9f, 0xe6, 0x54, 0x2a, 0xf8, 0x47, 0x0b, 0x80, 0x2a, 0x79,
	0xff, 0x91, 0xdf, 0xda, 0x6e, 0x75, 0xd6, 0x0f, 0xa7, 0x45, 0x1e, 0xdc, 0x1c, 0x33, 0x31, 0xdd,
	0x43, 0x0d, 0x86, 0x5e, 0x3c, 0x0f, 0xbe, 0x05, 0xf8, 0xe4, 0xf8, 0x80, 0x74, 0x7f, 0xf9, 0xa2,
	0xfb, 0xe3, 0xfb, 0xdd, 0x4f, 0x07, 0xbf, 0x3e, 0x78, 0xd6, 0x5d, 0x8a, 0x3f, 0xfa, 0x8f, 0xf1,
	0x07, 0xbd, 0x67, 0xf7, 0xb0, 0x53, 0x00, 0xbc, 0x0b, 0xda, 0x33, 0x32, 0xa5, 0xfe, 0x9b, 0xa6,
	0x90, 0xeb, 0x45, 0x1e, 0x78, 0xb6, 0x10, 0xbd, 0x8a, 0xb0, 0x01, 0xe1, 0xc7, 0x00, 0x10, 0x9e,
	0x7e, 0x47, 0x85, 0x4c, 0xd9, 0xcc, 0xbf, 0x62, 0xa8, 0x77, 0x9a, 0x9a, 0x1b, 0x0c, 0x61, 0x87,
	0x08, 0x0f, 0xc1, 0x5a, 0x39, 0x5c, 0xd2, 0x6f, 0x6f, 0x5f, 0xe9, 0x78, 0xbd, 0x9b, 0xa1, 0x1e,
	0xb9, 0xf0, 0xa1, 0x5d, 0xed, 0xcf, 0xc6, 0xec, 0xf0, 0x56, 0x91, 0x07, 0xd7, 0x6d, 0x9e, 0x8a,
	0x8c, 0x70, 0xad, 0x83, 0x03, 0xe0, 0x71, 0x92, 0x9c, 0x93, 0x09, 0xfd, 0x3a, 0x95, 0xca, 0x5f,
	0xd9, 0x6e, 0x75, 0xbc, 0xde, 0xdd, 0xb0, 0xee, 0xcf, 0xd2, 0xed, 0x86, 0x8f, 0x1b, 0xea, 0xe1,
	0x5b, 0x45, 0x1e, 0x40, 0x9b, 0xd8, 0xc9, 0x80, 0xb0, 0x9b, 0x6f, 0xeb, 0xcf, 0x55, 0xe0, 0x39,
	0x22, 0x78, 0x00, 0x56, 0xcb, 0xad, 0x4d, 0x6b, 0xfe, 0xb5, 0x62, 0x58, 0xe4, 0xc1, 0xb5, 0xa5,
	0x8a, 0x11, 0xae, 0x54, 0x70, 0x1f, 0x6c, 0xda, 0xc9, 0xac, 0x6e, 0xcb, 0x5e, 0xac, 0x5f, 0xe4,
	0xc1, 0x6d, 0xab, 0x59, 0x82, 0x11, 0x5e, 0xa6, 0xc3, 0xf7, 0xc0, 0x2a, 0x17, 0xf4, 0x48, 0x31,
	0x5e, 0xde, 0xb3, 0xb3, 0x5b, 0x09, 0x20, 0x5c, 0x51, 0xe0, 0x03, 0xe0, 0x71, 0x26, 0x15, 0xa6,
	0x52, 0x11, 0xa1, 0xfc, 0xb6, 0x51, 0xb8, 0x07, 0x6f, 0x40, 0x7d, 0xf0, 0x26, 0x82, 0x3d, 0xb0,
	0x4e, 0xe6, 0x8a, 0x1d, 0x19, 0xdd, 0x8a, 0xd1, 0xdd, 0x2e, 0xf2, 0xe0, 0x46, 0xd9, 0xd1, 0x0a,
	0x42, 0xb8, 0xa1, 0x69, 0xcd, 0x5c, 0x52, 0xf1, 0xf0, 0x8c, 0x26, 0xe7, 0xfe, 0xd5, 0x97, 0x35,
	0x35, 0x84, 0x70, 0x43, 0xd3, 0xa3, 0x33, 0x9e, 0x67, 0x99, 0x6d, 0x93, 0xbf, 0xfa, 0xf2, 0xe8,
	0x34, 0x18, 0xc2, 0x0e, 0x11, 0xbe, 0x0b, 0x56, 0xc6, 0x4c, 0x24, 0xd4, 0x5f, 0x33, 0x8a, 0x1b,
	0x45, 0x1e, 0x6c, 0x94, 0x0a, 0xbd, 0x8c, 0xb0, 0x85, 0xe1, 0x67, 0x60, 0xc3, 0xbc, 0xf4, 0x67,
	0x52, 0x91, 0x2c, 0xf3, 0xd7, 0x0d, 0xfd, 0xed, 0x22, 0x0f, 0x6e, 0x39, 0xf4, 0x12, 0x45, 0x78,
	0x89, 0x0c, 0x7f, 0x6f, 0x81, 0xf5, 0x72, 0x18, 0xfa, 0x23, 0x1f, 0x18, 0x69, 0xd6, 0x1c, 0xa8,
	0x86, 0x5e, 0x9f, 0x13, 0x9b, 0xed, 0xe1, 0x6e, 0x69, 0x44, 0xaf, 0x3e, 0x81, 0x6b, 0xc4, 0x17,
	0xcf, 0x83, 0x36, 0x7f, 0xe3, 0x69, 0x58, 0x1a, 0xf2, 0xb7, 0x16, 0xf0, 0x52, 0x7b, 0x8a, 0xc7,
	0x44, 0x9d, 0xf9, 0x1b, 0x46, 0x94, 0x34, 0x8d, 0x77, 0x40, 0xad, 0xed, 0x83, 0xaf, 0x4e, 0x3a,
	0x9d, 0x38, 0x3a, 0x3e, 0x89, 0xa3, 0xbd, 0xf8, 0x7e, 0x7c, 0x80, 0xd0, 0xe7, 0xfb, 0xf1, 0x65,
	0x2c, 0xe2, 0xd9, 0x60, 0x77, 0x67, 0x77, 0xe7, 0xb2, 0x13, 0x47, 0x3b, 0x97, 0xc7, 0xb6, 0xbc,
	0xc1, 0x5e, 0x27, 0x8e, 0x8f, 0x4f, 0xe2, 0xf8, 0x9f, 0xcc, 0xfb, 0xf7, 0xb0, 0xbb, 0xaf, 0xfe,
	0x7a, 0xa8, 0x05, 0xa7, 0xfe, 0xe6, 0x76, 0xab, 0xb3, 0xe2, 0x7e, 0x3d, 0xf4, 0x2a, 0xc2, 0x06,
	0x84, 0x11, 0x58, 0xe3, 0x19, 0x51, 0x1a, 0xf0, 0xaf, 0x99, 0x42, 0x1d, 0xcf, 0x57, 0x08, 0xc2,
	0x35, 0x09, 0xe5, 0x2d, 0x70, 0xa7, 0xf2, 0xb5, 0xe4, 0x6c, 0x26, 0xe9, 0xf7, 0x82, 0x70, 0x4e,
	0x85, 0xde, 0x2f, 0x61, 0x23, 0x6a, 0xbc, 0xb9, 0xb4, 0x9f, 0x5e, 0x45, 0xd8, 0x80, 0xda, 0x14,
	0xfa, 0xf9, 0xe5, 0x53, 0x9e, 0x91, 0xb4, 0x32, 0xa0, 0x63, 0x0a, 0x07, 0x44, 0xd8, 0xa5, 0xea,
	0xa9, 0xa3, 0x42, 0x30, 0x51, 0x5a, 0xcf, 0x99, 0x3a, 0xb3, 0x8c, 0xb0, 0x85, 0xe1, 0x3e, 0x68,
	0x8f, 0x88, 0x22, 0xc6, 0x6f, 0x5e, 0xef, 0x9d, 0xb0, 0xfe, 0x7d, 0x84, 0x8f, 0xcc, 0xef, 0xe3,
	0xa8, 0xfc, 0x3a, 0xb9, 0x15, 0x6a, 0x01, 0xc2, 0x46, 0x37, 0xbc, 0x6a, 0xfe, 0x0e, 0x1f, 0xfe,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x6c, 0xbc, 0xd1, 0x02, 0x08, 0x00, 0x00,
}
