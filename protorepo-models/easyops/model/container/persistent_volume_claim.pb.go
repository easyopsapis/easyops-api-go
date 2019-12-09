// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistent_volume_claim.proto

package container

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
//PersistentVolumeClaim 资源定义, 用于 CMDB 存储
type PersistentVolumeClaim struct {
	//
	//PVC id，服务端自动生成
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//PVC 类型, 只能是 PersistentVolumeClaim
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind" form:"kind"`
	//
	//PVC 全称，命名规则 clusterId:kind:namespace:name, 创建之后不能修改
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//原始名称
	ResourceName string `protobuf:"bytes,4,opt,name=resourceName,proto3" json:"resourceName" form:"resourceName"`
	//
	//PVC 定义
	ResourceSpec string `protobuf:"bytes,5,opt,name=resourceSpec,proto3" json:"resourceSpec" form:"resourceSpec"`
	//
	//存储容量, 如:8Gi, 800Mi
	Storage string `protobuf:"bytes,6,opt,name=storage,proto3" json:"storage" form:"storage"`
	//
	//访问模式
	AccessModes []string `protobuf:"bytes,7,rep,name=accessModes,proto3" json:"accessModes" form:"accessModes"`
	//
	//命名空间, 创建之后不能修改
	Namespace string `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace" form:"namespace"`
	//
	//存储类, 创建之后不能修改
	StorageClass string `protobuf:"bytes,9,opt,name=storageClass,proto3" json:"storageClass" form:"storageClass"`
	//
	//创建时间
	CreationTimestamp string `protobuf:"bytes,10,opt,name=creationTimestamp,proto3" json:"creationTimestamp" form:"creationTimestamp"`
	//
	//状态
	Phase                string   `protobuf:"bytes,11,opt,name=phase,proto3" json:"phase" form:"phase"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersistentVolumeClaim) Reset()         { *m = PersistentVolumeClaim{} }
func (m *PersistentVolumeClaim) String() string { return proto.CompactTextString(m) }
func (*PersistentVolumeClaim) ProtoMessage()    {}
func (*PersistentVolumeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb968e444c7d529d, []int{0}
}
func (m *PersistentVolumeClaim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistentVolumeClaim.Unmarshal(m, b)
}
func (m *PersistentVolumeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistentVolumeClaim.Marshal(b, m, deterministic)
}
func (m *PersistentVolumeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistentVolumeClaim.Merge(m, src)
}
func (m *PersistentVolumeClaim) XXX_Size() int {
	return xxx_messageInfo_PersistentVolumeClaim.Size(m)
}
func (m *PersistentVolumeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistentVolumeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_PersistentVolumeClaim proto.InternalMessageInfo

func (m *PersistentVolumeClaim) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *PersistentVolumeClaim) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PersistentVolumeClaim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersistentVolumeClaim) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PersistentVolumeClaim) GetResourceSpec() string {
	if m != nil {
		return m.ResourceSpec
	}
	return ""
}

func (m *PersistentVolumeClaim) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *PersistentVolumeClaim) GetAccessModes() []string {
	if m != nil {
		return m.AccessModes
	}
	return nil
}

func (m *PersistentVolumeClaim) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PersistentVolumeClaim) GetStorageClass() string {
	if m != nil {
		return m.StorageClass
	}
	return ""
}

func (m *PersistentVolumeClaim) GetCreationTimestamp() string {
	if m != nil {
		return m.CreationTimestamp
	}
	return ""
}

func (m *PersistentVolumeClaim) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

func init() {
	proto.RegisterType((*PersistentVolumeClaim)(nil), "container.PersistentVolumeClaim")
}

func init() { proto.RegisterFile("persistent_volume_claim.proto", fileDescriptor_bb968e444c7d529d) }

var fileDescriptor_bb968e444c7d529d = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdf, 0x4e, 0xd4, 0x4e,
	0x14, 0xc7, 0xd3, 0x1f, 0xff, 0x7e, 0x3b, 0x10, 0x84, 0xfa, 0xaf, 0x21, 0x31, 0x25, 0x95, 0x98,
	0x16, 0x9c, 0x6d, 0x77, 0x57, 0x8d, 0xac, 0x17, 0x1b, 0xd9, 0x78, 0xc1, 0x85, 0xc6, 0x54, 0xe2,
	0x05, 0xb5, 0x4b, 0x86, 0xd9, 0x61, 0x69, 0x68, 0x3b, 0x4d, 0x67, 0x16, 0x22, 0x2c, 0x37, 0x3e,
	0x8e, 0x89, 0xaf, 0xe2, 0x23, 0xd4, 0xc4, 0x17, 0x30, 0xe9, 0xbd, 0x89, 0x99, 0x99, 0x5d, 0xb6,
	0x80, 0xf1, 0xaa, 0xe7, 0x9c, 0xef, 0xe7, 0x7b, 0xe6, 0x9c, 0x4e, 0x5a, 0xf0, 0x28, 0x23, 0x39,
	0x8b, 0x18, 0x27, 0x29, 0x3f, 0x38, 0xa5, 0xf1, 0x30, 0x21, 0x07, 0x38, 0x46, 0x51, 0x52, 0xcf,
	0x72, 0xca, 0xa9, 0x5e, 0xc3, 0x34, 0xe5, 0x28, 0x4a, 0x49, 0xbe, 0x06, 0x07, 0x11, 0x3f, 0x1e,
	0x1e, 0xd6, 0x31, 0x4d, 0xdc, 0x01, 0x1d, 0x50, 0x57, 0x12, 0x87, 0xc3, 0x23, 0x99, 0xc9, 0x44,
	0x46, 0xca, 0xb9, 0xf6, 0xa2, 0x82, 0x27, 0x67, 0x11, 0x3f, 0xa1, 0x67, 0xee, 0x80, 0x42, 0x29,
	0xc2, 0x53, 0x14, 0x47, 0x7d, 0xc4, 0x69, 0xce, 0xdc, 0xab, 0x50, 0xf9, 0xac, 0x5f, 0xf3, 0xe0,
	0xfe, 0xfb, 0xab, 0x99, 0x3e, 0xca, 0x91, 0xba, 0x62, 0x22, 0x7d, 0x17, 0x80, 0x28, 0x65, 0x1c,
	0xa5, 0x98, 0xec, 0xf6, 0x0d, 0x6d, 0x5d, 0xb3, 0x6b, 0x3b, 0x4e, 0x59, 0x98, 0xab, 0x47, 0x34,
	0x4f, 0xda, 0xd6, 0x54, 0xb3, 0x7e, 0xfe, 0x30, 0x57, 0xc0, 0x72, 0x2f, 0xf0, 0xe0, 0x36, 0x82,
	0xe7, 0xe1, 0x45, 0xa3, 0x75, 0xb9, 0xe1, 0x57, 0xcc, 0xfa, 0x63, 0x30, 0x7b, 0x12, 0xa5, 0x7d,
	0xe3, 0x3f, 0xd9, 0xe4, 0x4e, 0x59, 0x98, 0x8b, 0xaa, 0x89, 0xa8, 0x5a, 0xbe, 0x14, 0x05, 0x94,
	0xa2, 0x84, 0x18, 0x33, 0x37, 0x21, 0x51, 0xb5, 0x7c, 0x29, 0xea, 0x5f, 0x34, 0xb0, 0x94, 0x13,
	0x46, 0x87, 0x39, 0x26, 0xef, 0x04, 0x3d, 0x2b, 0xe9, 0x5e, 0x59, 0x98, 0x77, 0x15, 0x5d, 0x55,
	0xc5, 0x64, 0x5d, 0xf0, 0xba, 0x17, 0x20, 0x78, 0xee, 0xc1, 0xed, 0xd0, 0x0e, 0xe0, 0x38, 0xda,
	0x9c, 0x94, 0x9c, 0x8e, 0xfd, 0xa9, 0xfe, 0x4f, 0xdd, 0xd9, 0xdc, 0xf0, 0xaf, 0x9d, 0xa9, 0xbf,
	0x9a, 0xce, 0xf0, 0x21, 0x23, 0xd8, 0x98, 0x93, 0x33, 0x3c, 0xbc, 0x3d, 0x83, 0x50, 0x2d, 0xff,
	0x1a, 0xac, 0x3f, 0x05, 0x0b, 0x8c, 0xd3, 0x1c, 0x0d, 0x88, 0x31, 0x2f, 0x7d, 0x7a, 0x59, 0x98,
	0xcb, 0xca, 0x37, 0x16, 0x2c, 0x7f, 0x82, 0xe8, 0x2f, 0xc1, 0x22, 0xc2, 0x98, 0x30, 0xf6, 0x96,
	0xf6, 0x09, 0x33, 0x16, 0xd6, 0x67, 0xec, 0xda, 0xce, 0x83, 0xb2, 0x30, 0x75, 0xe5, 0xa8, 0x88,
	0x96, 0x5f, 0x45, 0xf5, 0x26, 0xa8, 0x89, 0x37, 0xc6, 0x32, 0x84, 0x89, 0xf1, 0xbf, 0x3c, 0xe9,
	0x5e, 0x59, 0x98, 0x2b, 0xd3, 0x77, 0x2a, 0x25, 0xcb, 0x9f, 0x62, 0x62, 0xb1, 0xf1, 0xc1, 0xdd,
	0x18, 0x31, 0x66, 0xd4, 0x6e, 0x2e, 0x56, 0x55, 0x2d, 0xff, 0x1a, 0xac, 0xff, 0xd6, 0xc0, 0x2a,
	0xce, 0x09, 0xe2, 0x11, 0x4d, 0xf7, 0xa2, 0x84, 0x30, 0x8e, 0x92, 0xcc, 0x00, 0xb2, 0xc5, 0x77,
	0xad, 0x2c, 0x4c, 0x43, 0xf5, 0xb8, 0xc5, 0x88, 0x5b, 0xfa, 0xa6, 0x81, 0xaf, 0x5a, 0xcf, 0xb6,
	0x3b, 0xed, 0xa0, 0x01, 0xb7, 0xc3, 0x40, 0xde, 0x81, 0xd3, 0x91, 0xcf, 0x8b, 0x67, 0x97, 0x0e,
	0xb4, 0x1b, 0x81, 0x07, 0x9b, 0xe1, 0xc8, 0x93, 0xba, 0x03, 0xed, 0x56, 0xe0, 0xc1, 0xc6, 0x24,
	0x1f, 0x05, 0x0d, 0xd8, 0x54, 0x2e, 0x27, 0xd8, 0x5b, 0x0f, 0xed, 0x66, 0xe0, 0xc1, 0x56, 0x38,
	0x92, 0x8c, 0x2a, 0xb7, 0xed, 0xc0, 0x83, 0xcf, 0x27, 0xc9, 0x34, 0x16, 0xd7, 0x2f, 0x9e, 0x5b,
	0x4e, 0xc7, 0xde, 0x1f, 0x05, 0x5b, 0x30, 0xb4, 0x3b, 0xed, 0xbf, 0xd8, 0x2b, 0xee, 0xce, 0x86,
	0x7f, 0x7b, 0x53, 0xfd, 0x09, 0x98, 0xcb, 0x8e, 0x11, 0x23, 0xc6, 0xa2, 0x5c, 0x79, 0xa5, 0x2c,
	0xcc, 0x25, 0xb5, 0xb1, 0x2c, 0x5b, 0xbe, 0x92, 0x77, 0xde, 0xec, 0x77, 0x07, 0xb4, 0x4e, 0x10,
	0xfb, 0x4c, 0x33, 0x56, 0x8f, 0x29, 0x46, 0xb1, 0x2b, 0x3e, 0xfb, 0x1c, 0x61, 0xce, 0xd4, 0x57,
	0x9e, 0x93, 0x8c, 0xc2, 0x84, 0xf6, 0x49, 0xcc, 0xdc, 0x31, 0xe8, 0xca, 0xd4, 0xbd, 0xfa, 0x3f,
	0x1c, 0xce, 0x4b, 0xb2, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xf7, 0xaa, 0xa9, 0x52, 0x04,
	0x00, 0x00,
}
