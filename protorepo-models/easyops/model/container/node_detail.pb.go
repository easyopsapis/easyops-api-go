// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node_detail.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//node 实例
type NodeDetail struct {
	//
	//元数据
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata" form:"metadata"`
	//
	//规格
	Spec *NodeDetail_Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" form:"spec"`
	//
	//节点状态
	Status *NodeDetail_Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status" form:"status"`
	//
	//是否为worker节点
	Worker               bool     `protobuf:"varint,4,opt,name=worker,proto3" json:"worker" form:"worker"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDetail) Reset()         { *m = NodeDetail{} }
func (m *NodeDetail) String() string { return proto.CompactTextString(m) }
func (*NodeDetail) ProtoMessage()    {}
func (*NodeDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0}
}
func (m *NodeDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail.Unmarshal(m, b)
}
func (m *NodeDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail.Marshal(b, m, deterministic)
}
func (m *NodeDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail.Merge(m, src)
}
func (m *NodeDetail) XXX_Size() int {
	return xxx_messageInfo_NodeDetail.Size(m)
}
func (m *NodeDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail proto.InternalMessageInfo

func (m *NodeDetail) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NodeDetail) GetSpec() *NodeDetail_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *NodeDetail) GetStatus() *NodeDetail_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *NodeDetail) GetWorker() bool {
	if m != nil {
		return m.Worker
	}
	return false
}

type NodeDetail_Spec struct {
	//
	//pod ID
	PodCIDR string `protobuf:"bytes,1,opt,name=podCIDR,proto3" json:"podCIDR" form:"podCIDR"`
	//
	//provider ID
	ProviderID string `protobuf:"bytes,2,opt,name=providerID,proto3" json:"providerID" form:"providerID"`
	//
	//是否可调度
	Unschedulable        bool     `protobuf:"varint,3,opt,name=unschedulable,proto3" json:"unschedulable" form:"unschedulable"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDetail_Spec) Reset()         { *m = NodeDetail_Spec{} }
func (m *NodeDetail_Spec) String() string { return proto.CompactTextString(m) }
func (*NodeDetail_Spec) ProtoMessage()    {}
func (*NodeDetail_Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0, 0}
}
func (m *NodeDetail_Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail_Spec.Unmarshal(m, b)
}
func (m *NodeDetail_Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail_Spec.Marshal(b, m, deterministic)
}
func (m *NodeDetail_Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail_Spec.Merge(m, src)
}
func (m *NodeDetail_Spec) XXX_Size() int {
	return xxx_messageInfo_NodeDetail_Spec.Size(m)
}
func (m *NodeDetail_Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail_Spec proto.InternalMessageInfo

func (m *NodeDetail_Spec) GetPodCIDR() string {
	if m != nil {
		return m.PodCIDR
	}
	return ""
}

func (m *NodeDetail_Spec) GetProviderID() string {
	if m != nil {
		return m.ProviderID
	}
	return ""
}

func (m *NodeDetail_Spec) GetUnschedulable() bool {
	if m != nil {
		return m.Unschedulable
	}
	return false
}

type NodeDetail_Status struct {
	//
	//容量
	Capacity *NodeDetail_Status_Capacity `protobuf:"bytes,1,opt,name=capacity,proto3" json:"capacity" form:"capacity"`
	//
	//可分配容量
	Allocatable *NodeDetail_Status_Allocatable `protobuf:"bytes,2,opt,name=allocatable,proto3" json:"allocatable" form:"allocatable"`
	//
	//阶段
	Phase string `protobuf:"bytes,3,opt,name=phase,proto3" json:"phase" form:"phase"`
	//
	//节点信息
	NodeInfo             *NodeDetail_Status_NodeInfo `protobuf:"bytes,4,opt,name=nodeInfo,proto3" json:"nodeInfo" form:"nodeInfo"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NodeDetail_Status) Reset()         { *m = NodeDetail_Status{} }
func (m *NodeDetail_Status) String() string { return proto.CompactTextString(m) }
func (*NodeDetail_Status) ProtoMessage()    {}
func (*NodeDetail_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0, 1}
}
func (m *NodeDetail_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail_Status.Unmarshal(m, b)
}
func (m *NodeDetail_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail_Status.Marshal(b, m, deterministic)
}
func (m *NodeDetail_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail_Status.Merge(m, src)
}
func (m *NodeDetail_Status) XXX_Size() int {
	return xxx_messageInfo_NodeDetail_Status.Size(m)
}
func (m *NodeDetail_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail_Status.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail_Status proto.InternalMessageInfo

func (m *NodeDetail_Status) GetCapacity() *NodeDetail_Status_Capacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *NodeDetail_Status) GetAllocatable() *NodeDetail_Status_Allocatable {
	if m != nil {
		return m.Allocatable
	}
	return nil
}

func (m *NodeDetail_Status) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func (m *NodeDetail_Status) GetNodeInfo() *NodeDetail_Status_NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

type NodeDetail_Status_Capacity struct {
	//
	//cpu
	Cpu string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu" form:"cpu"`
	//
	//memory
	Memory string `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory" form:"memory"`
	//
	//pod数
	Pods string `protobuf:"bytes,3,opt,name=pods,proto3" json:"pods" form:"pods"`
	//
	//临时存储
	EphemeralStorage     string   `protobuf:"bytes,4,opt,name=ephemeralStorage,proto3" json:"ephemeralStorage" form:"ephemeralStorage"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDetail_Status_Capacity) Reset()         { *m = NodeDetail_Status_Capacity{} }
func (m *NodeDetail_Status_Capacity) String() string { return proto.CompactTextString(m) }
func (*NodeDetail_Status_Capacity) ProtoMessage()    {}
func (*NodeDetail_Status_Capacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0, 1, 0}
}
func (m *NodeDetail_Status_Capacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail_Status_Capacity.Unmarshal(m, b)
}
func (m *NodeDetail_Status_Capacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail_Status_Capacity.Marshal(b, m, deterministic)
}
func (m *NodeDetail_Status_Capacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail_Status_Capacity.Merge(m, src)
}
func (m *NodeDetail_Status_Capacity) XXX_Size() int {
	return xxx_messageInfo_NodeDetail_Status_Capacity.Size(m)
}
func (m *NodeDetail_Status_Capacity) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail_Status_Capacity.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail_Status_Capacity proto.InternalMessageInfo

func (m *NodeDetail_Status_Capacity) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *NodeDetail_Status_Capacity) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *NodeDetail_Status_Capacity) GetPods() string {
	if m != nil {
		return m.Pods
	}
	return ""
}

func (m *NodeDetail_Status_Capacity) GetEphemeralStorage() string {
	if m != nil {
		return m.EphemeralStorage
	}
	return ""
}

type NodeDetail_Status_Allocatable struct {
	//
	//cpu
	Cpu string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu" form:"cpu"`
	//
	//memory
	Memory string `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory" form:"memory"`
	//
	//pod数
	Pods                 string   `protobuf:"bytes,3,opt,name=pods,proto3" json:"pods" form:"pods"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDetail_Status_Allocatable) Reset()         { *m = NodeDetail_Status_Allocatable{} }
func (m *NodeDetail_Status_Allocatable) String() string { return proto.CompactTextString(m) }
func (*NodeDetail_Status_Allocatable) ProtoMessage()    {}
func (*NodeDetail_Status_Allocatable) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0, 1, 1}
}
func (m *NodeDetail_Status_Allocatable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail_Status_Allocatable.Unmarshal(m, b)
}
func (m *NodeDetail_Status_Allocatable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail_Status_Allocatable.Marshal(b, m, deterministic)
}
func (m *NodeDetail_Status_Allocatable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail_Status_Allocatable.Merge(m, src)
}
func (m *NodeDetail_Status_Allocatable) XXX_Size() int {
	return xxx_messageInfo_NodeDetail_Status_Allocatable.Size(m)
}
func (m *NodeDetail_Status_Allocatable) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail_Status_Allocatable.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail_Status_Allocatable proto.InternalMessageInfo

func (m *NodeDetail_Status_Allocatable) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *NodeDetail_Status_Allocatable) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *NodeDetail_Status_Allocatable) GetPods() string {
	if m != nil {
		return m.Pods
	}
	return ""
}

type NodeDetail_Status_NodeInfo struct {
	//
	//机器ID
	MachineID string `protobuf:"bytes,1,opt,name=machineID,proto3" json:"machineID" form:"machineID"`
	//
	//systemUUID, 专用于Red Hat
	SystemUUID string `protobuf:"bytes,2,opt,name=systemUUID,proto3" json:"systemUUID" form:"systemUUID"`
	//
	//Boot ID reported by the node.
	BootID string `protobuf:"bytes,3,opt,name=bootID,proto3" json:"bootID" form:"bootID"`
	//
	//内核版本
	KernelVersion string `protobuf:"bytes,4,opt,name=kernelVersion,proto3" json:"kernelVersion" form:"kernelVersion"`
	//
	//镜像
	OsImage string `protobuf:"bytes,5,opt,name=osImage,proto3" json:"osImage" form:"osImage"`
	//
	//容器版本,如:docker://1.5.0
	ContainerRuntimeVersion string `protobuf:"bytes,6,opt,name=containerRuntimeVersion,proto3" json:"containerRuntimeVersion" form:"containerRuntimeVersion"`
	//
	//kubeletVersion
	KubeletVersion string `protobuf:"bytes,7,opt,name=kubeletVersion,proto3" json:"kubeletVersion" form:"kubeletVersion"`
	//
	//kubeProxyVersion
	KubeProxyVersion string `protobuf:"bytes,8,opt,name=kubeProxyVersion,proto3" json:"kubeProxyVersion" form:"kubeProxyVersion"`
	//
	//操作系统
	OperatingSystem string `protobuf:"bytes,9,opt,name=operatingSystem,proto3" json:"operatingSystem" form:"operatingSystem"`
	//
	//系统架构
	Architecture         string   `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture" form:"architecture"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeDetail_Status_NodeInfo) Reset()         { *m = NodeDetail_Status_NodeInfo{} }
func (m *NodeDetail_Status_NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeDetail_Status_NodeInfo) ProtoMessage()    {}
func (*NodeDetail_Status_NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7beee7b24470bccc, []int{0, 1, 2}
}
func (m *NodeDetail_Status_NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeDetail_Status_NodeInfo.Unmarshal(m, b)
}
func (m *NodeDetail_Status_NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeDetail_Status_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeDetail_Status_NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDetail_Status_NodeInfo.Merge(m, src)
}
func (m *NodeDetail_Status_NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeDetail_Status_NodeInfo.Size(m)
}
func (m *NodeDetail_Status_NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDetail_Status_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDetail_Status_NodeInfo proto.InternalMessageInfo

func (m *NodeDetail_Status_NodeInfo) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetSystemUUID() string {
	if m != nil {
		return m.SystemUUID
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetBootID() string {
	if m != nil {
		return m.BootID
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetKernelVersion() string {
	if m != nil {
		return m.KernelVersion
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetOsImage() string {
	if m != nil {
		return m.OsImage
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetContainerRuntimeVersion() string {
	if m != nil {
		return m.ContainerRuntimeVersion
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetKubeletVersion() string {
	if m != nil {
		return m.KubeletVersion
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetKubeProxyVersion() string {
	if m != nil {
		return m.KubeProxyVersion
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

func (m *NodeDetail_Status_NodeInfo) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeDetail)(nil), "container.NodeDetail")
	proto.RegisterType((*NodeDetail_Spec)(nil), "container.NodeDetail.Spec")
	proto.RegisterType((*NodeDetail_Status)(nil), "container.NodeDetail.Status")
	proto.RegisterType((*NodeDetail_Status_Capacity)(nil), "container.NodeDetail.Status.Capacity")
	proto.RegisterType((*NodeDetail_Status_Allocatable)(nil), "container.NodeDetail.Status.Allocatable")
	proto.RegisterType((*NodeDetail_Status_NodeInfo)(nil), "container.NodeDetail.Status.NodeInfo")
}

func init() { proto.RegisterFile("node_detail.proto", fileDescriptor_7beee7b24470bccc) }

var fileDescriptor_7beee7b24470bccc = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x56, 0xfe, 0xa6, 0x69, 0x32, 0xe9, 0x75, 0x7a, 0xcb, 0x6f, 0x10, 0xae, 0x06, 0x81, 0x8a,
	0x44, 0x1d, 0xa9, 0x88, 0x0d, 0x48, 0xa0, 0xba, 0x41, 0x55, 0x16, 0x54, 0x68, 0xaa, 0x76, 0x81,
	0x90, 0xd0, 0xc4, 0x9e, 0x26, 0x56, 0x6d, 0x8f, 0x35, 0x1e, 0x03, 0x59, 0xf1, 0x00, 0x3c, 0x0a,
	0x8f, 0x81, 0x78, 0x0d, 0x3f, 0x84, 0xb7, 0x6c, 0xd0, 0x5c, 0x9c, 0x38, 0xe9, 0x85, 0x0d, 0x62,
	0x17, 0x9f, 0xef, 0xe2, 0xf9, 0xce, 0x39, 0xe3, 0x80, 0x8d, 0x98, 0xf9, 0xf4, 0xa3, 0x4f, 0x05,
	0x09, 0x42, 0x27, 0xe1, 0x4c, 0x30, 0xd8, 0xf2, 0x58, 0x2c, 0x48, 0x10, 0x53, 0x6e, 0x1d, 0x0c,
	0x03, 0x31, 0xca, 0x06, 0x8e, 0xc7, 0xa2, 0xee, 0x90, 0x0d, 0x59, 0x57, 0x31, 0x06, 0xd9, 0xa5,
	0x7a, 0x52, 0x0f, 0xea, 0x97, 0x56, 0x5a, 0x78, 0xc8, 0x1c, 0x4a, 0xd2, 0x31, 0x4b, 0x52, 0x27,
	0x64, 0x1e, 0x09, 0xbb, 0xd2, 0x8a, 0x13, 0x4f, 0xa4, 0x5a, 0xc9, 0x69, 0xc2, 0x0e, 0x22, 0xe6,
	0xd3, 0x30, 0xed, 0x1a, 0x62, 0x57, 0x3d, 0x76, 0x27, 0xef, 0xec, 0x46, 0x54, 0x10, 0x9f, 0x08,
	0xa2, 0x3d, 0xd1, 0x8f, 0x15, 0x00, 0x4e, 0x99, 0x4f, 0x7b, 0xea, 0x88, 0xb0, 0x07, 0x9a, 0x25,
	0xa1, 0x53, 0xdb, 0xab, 0xed, 0xb7, 0x0f, 0x37, 0x9d, 0x89, 0xd6, 0x79, 0x6b, 0x20, 0x77, 0xb3,
	0xc8, 0xed, 0xb5, 0x4b, 0xc6, 0xa3, 0x17, 0xa8, 0xa4, 0x23, 0x3c, 0x51, 0xc2, 0xd7, 0xa0, 0x9e,
	0x26, 0xd4, 0xeb, 0xfc, 0xa7, 0x1c, 0xac, 0x8a, 0xc3, 0xf4, 0x55, 0xce, 0x59, 0x42, 0x3d, 0x77,
	0xad, 0xc8, 0xed, 0xb6, 0x36, 0x92, 0x0a, 0x84, 0x95, 0x10, 0x9e, 0x80, 0x46, 0x2a, 0x88, 0xc8,
	0xd2, 0xce, 0x82, 0xb2, 0xb8, 0x7f, 0x8b, 0x85, 0xe2, 0xb8, 0x1b, 0x45, 0x6e, 0xaf, 0x18, 0x13,
	0x55, 0x41, 0xd8, 0xc8, 0xe1, 0x13, 0xd0, 0xf8, 0xcc, 0xf8, 0x15, 0xe5, 0x9d, 0xfa, 0x5e, 0x6d,
	0xbf, 0x59, 0xa5, 0xea, 0x3a, 0xc2, 0x86, 0x60, 0x7d, 0xaf, 0x81, 0xba, 0x3c, 0x13, 0x7c, 0x0a,
	0x96, 0x12, 0xe6, 0x1f, 0xf7, 0x7b, 0x58, 0xb5, 0xa0, 0xe5, 0xc2, 0x22, 0xb7, 0x57, 0xb5, 0xc8,
	0x00, 0x08, 0x97, 0x14, 0xf8, 0x1c, 0x80, 0x84, 0xb3, 0x4f, 0x81, 0x4f, 0x79, 0xbf, 0xa7, 0x12,
	0xb7, 0xdc, 0xed, 0x22, 0xb7, 0x37, 0x8c, 0x60, 0x82, 0x21, 0x5c, 0x21, 0xc2, 0x57, 0x60, 0x25,
	0x8b, 0x53, 0x6f, 0x44, 0xfd, 0x2c, 0x24, 0x83, 0x90, 0xaa, 0xa0, 0x4d, 0xb7, 0x53, 0xe4, 0xf6,
	0x96, 0x56, 0xce, 0xc0, 0x08, 0xcf, 0xd2, 0xad, 0x6f, 0x00, 0x34, 0x74, 0x7c, 0x78, 0x01, 0x9a,
	0x1e, 0x49, 0x88, 0x17, 0x88, 0xb1, 0x99, 0xd9, 0xa3, 0xbb, 0xda, 0xe5, 0x1c, 0x1b, 0x72, 0x75,
	0x8a, 0xa5, 0x01, 0xc2, 0x13, 0x2f, 0x38, 0x00, 0x6d, 0x12, 0xca, 0x3d, 0x13, 0xea, 0x80, 0x7a,
	0x98, 0xfb, 0x77, 0x5a, 0x1f, 0x4d, 0xf9, 0xee, 0x4e, 0x91, 0xdb, 0x50, 0xbb, 0x57, 0x6c, 0x10,
	0xae, 0x9a, 0xc2, 0xc7, 0x60, 0x31, 0x19, 0x91, 0x54, 0xc7, 0x6f, 0xb9, 0xeb, 0x45, 0x6e, 0x2f,
	0x9b, 0xc6, 0xc9, 0x32, 0xc2, 0x1a, 0x96, 0x19, 0xe5, 0x4d, 0xea, 0xc7, 0x97, 0x4c, 0x4d, 0xf2,
	0x4f, 0x19, 0x4f, 0x0d, 0xb9, 0x9a, 0xb1, 0x34, 0x40, 0x78, 0xe2, 0x65, 0xfd, 0xac, 0x81, 0x66,
	0xd9, 0x0f, 0xb8, 0x07, 0x16, 0xbc, 0x24, 0x33, 0x43, 0x5f, 0x2d, 0x72, 0x1b, 0x98, 0xe6, 0x24,
	0x19, 0xc2, 0x12, 0x92, 0xeb, 0x14, 0xd1, 0x88, 0xf1, 0xb1, 0x19, 0x74, 0x65, 0x9d, 0x74, 0x1d,
	0x61, 0x43, 0x80, 0x0f, 0x41, 0x3d, 0x61, 0x7e, 0x6a, 0x82, 0x55, 0xf6, 0x5c, 0x56, 0x11, 0x56,
	0x20, 0x3c, 0x01, 0xeb, 0x34, 0x19, 0xd1, 0x88, 0x72, 0x12, 0x9e, 0x09, 0xc6, 0xc9, 0x90, 0xaa,
	0x78, 0x2d, 0xf7, 0x5e, 0x91, 0xdb, 0xbb, 0x5a, 0x30, 0xcf, 0x40, 0xf8, 0x9a, 0xc8, 0xfa, 0x0a,
	0xda, 0x95, 0xde, 0xff, 0xfb, 0x24, 0xd6, 0xaf, 0x3a, 0x68, 0x96, 0x4d, 0x87, 0x87, 0xa0, 0x15,
	0x11, 0x6f, 0x14, 0xc4, 0xb4, 0xdf, 0x33, 0x87, 0xd8, 0x2a, 0x72, 0x7b, 0xdd, 0xf8, 0x97, 0x10,
	0xc2, 0x53, 0x9a, 0xbc, 0x47, 0xe9, 0x38, 0x15, 0x34, 0x3a, 0x3f, 0xbf, 0xe9, 0x1e, 0x4d, 0x31,
	0x84, 0x2b, 0x44, 0x99, 0x63, 0xc0, 0x98, 0xe8, 0xf7, 0xcc, 0xf1, 0x2a, 0x39, 0x74, 0x1d, 0x61,
	0x43, 0x90, 0x57, 0xee, 0x8a, 0xf2, 0x98, 0x86, 0x17, 0x94, 0xa7, 0x01, 0x8b, 0x4d, 0xa7, 0x2b,
	0x57, 0x6e, 0x06, 0x46, 0x78, 0x96, 0x2e, 0xbf, 0x0b, 0x2c, 0xed, 0x47, 0x72, 0x46, 0x8b, 0xf3,
	0xdf, 0x05, 0x03, 0x20, 0x5c, 0x52, 0xe0, 0x07, 0xb0, 0x3b, 0x59, 0x50, 0x9c, 0xc5, 0x22, 0x88,
	0x68, 0xf9, 0xde, 0x86, 0x52, 0xa3, 0x22, 0xb7, 0x1f, 0x98, 0xb1, 0xdc, 0x4c, 0x44, 0xf8, 0x36,
	0x0b, 0x78, 0x04, 0x56, 0xaf, 0xb2, 0x01, 0x0d, 0xa9, 0x28, 0x4d, 0x97, 0x94, 0xe9, 0xff, 0x45,
	0x6e, 0x6f, 0x9b, 0x30, 0x33, 0x38, 0xc2, 0x73, 0x02, 0xb9, 0x7b, 0xb2, 0xf2, 0x8e, 0xb3, 0x2f,
	0xe3, 0xd2, 0xa4, 0x39, 0xbf, 0x7b, 0xf3, 0x0c, 0x84, 0xaf, 0x89, 0x60, 0x0f, 0xac, 0xb1, 0x84,
	0x72, 0x22, 0x82, 0x78, 0x78, 0xa6, 0x26, 0xd3, 0x69, 0x29, 0x1f, 0xab, 0xc8, 0xed, 0x1d, 0xd3,
	0x9f, 0x59, 0x02, 0xc2, 0xf3, 0x12, 0xf8, 0x12, 0x2c, 0x13, 0xee, 0x8d, 0x02, 0x41, 0x3d, 0x91,
	0x71, 0xda, 0x01, 0xca, 0x62, 0xb7, 0xc8, 0xed, 0x4d, 0xf3, 0x11, 0xa9, 0xa0, 0x08, 0xcf, 0x90,
	0xdd, 0x37, 0xef, 0x8f, 0xff, 0xc2, 0x7f, 0xe3, 0xa0, 0xa1, 0x98, 0xcf, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x8a, 0xfe, 0x80, 0xb6, 0x07, 0x00, 0x00,
}
