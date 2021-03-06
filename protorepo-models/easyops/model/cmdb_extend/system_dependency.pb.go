// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: system_dependency.proto

package cmdb_extend

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
//系统依赖
type SystemDependency struct {
	//
	//系统简称
	Abbreviation string `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation" form:"abbreviation"`
	//
	//系统名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//模型ID
	ObjectId string `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id" form:"object_id"`
	//
	//实例ID
	InstanceId string `protobuf:"bytes,4,opt,name=instance_id,json=instanceId,proto3" json:"instance_id" form:"instance_id"`
	//
	//子系统依赖
	Subsystems           []*SubsystemDependency `protobuf:"bytes,5,rep,name=subsystems,proto3" json:"subsystems" form:"subsystems"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SystemDependency) Reset()         { *m = SystemDependency{} }
func (m *SystemDependency) String() string { return proto.CompactTextString(m) }
func (*SystemDependency) ProtoMessage()    {}
func (*SystemDependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_bddce165cd2a2809, []int{0}
}
func (m *SystemDependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemDependency.Unmarshal(m, b)
}
func (m *SystemDependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemDependency.Marshal(b, m, deterministic)
}
func (m *SystemDependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemDependency.Merge(m, src)
}
func (m *SystemDependency) XXX_Size() int {
	return xxx_messageInfo_SystemDependency.Size(m)
}
func (m *SystemDependency) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemDependency.DiscardUnknown(m)
}

var xxx_messageInfo_SystemDependency proto.InternalMessageInfo

func (m *SystemDependency) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *SystemDependency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemDependency) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *SystemDependency) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *SystemDependency) GetSubsystems() []*SubsystemDependency {
	if m != nil {
		return m.Subsystems
	}
	return nil
}

func init() {
	proto.RegisterType((*SystemDependency)(nil), "cmdb_extend.SystemDependency")
}

func init() { proto.RegisterFile("system_dependency.proto", fileDescriptor_bddce165cd2a2809) }

var fileDescriptor_bddce165cd2a2809 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xd9, 0x3f, 0x8a, 0x3b, 0x11, 0xad, 0x11, 0xd9, 0x50, 0x90, 0x94, 0xb8, 0x17, 0x8b,
	0x98, 0xcc, 0xae, 0x05, 0x45, 0xbd, 0xb2, 0x28, 0xb2, 0x78, 0x97, 0x15, 0x84, 0x5d, 0x76, 0xe3,
	0xfc, 0x6b, 0x8c, 0x26, 0x73, 0x42, 0x66, 0xda, 0x5a, 0x4b, 0x5f, 0x35, 0x82, 0x4f, 0x20, 0x79,
	0x02, 0xe9, 0x4c, 0x9b, 0xc6, 0x0b, 0xef, 0xf6, 0xee, 0x9c, 0x7c, 0xdf, 0xf7, 0x3b, 0x87, 0x93,
	0x41, 0x87, 0x6a, 0xae, 0xb4, 0x28, 0x12, 0x2e, 0x4a, 0x21, 0xb9, 0x90, 0x6c, 0x1e, 0x95, 0x15,
	0x68, 0x70, 0x1d, 0x56, 0x70, 0x9a, 0x88, 0x1f, 0x5a, 0x48, 0xde, 0x0f, 0xd3, 0x4c, 0x7f, 0x9d,
	0xd0, 0x88, 0x41, 0x81, 0x53, 0x48, 0x01, 0x1b, 0x0f, 0x9d, 0x8c, 0x4d, 0x67, 0x1a, 0x53, 0xd9,
	0x6c, 0xff, 0x4b, 0x0a, 0x91, 0x20, 0x6a, 0x0e, 0xa5, 0x8a, 0x72, 0x60, 0x24, 0xc7, 0x0c, 0xa4,
	0xae, 0x08, 0xd3, 0xca, 0x26, 0x2b, 0x51, 0x42, 0x58, 0x00, 0x17, 0xb9, 0xc2, 0x6b, 0x23, 0x36,
	0x2d, 0xee, 0x4c, 0xc5, 0x6a, 0x42, 0xff, 0xb3, 0x5d, 0xff, 0x45, 0x67, 0xa1, 0x62, 0x96, 0xe9,
	0xef, 0x30, 0xc3, 0x29, 0x84, 0x46, 0x0c, 0xa7, 0x24, 0xcf, 0x38, 0xd1, 0x50, 0x29, 0xdc, 0x96,
	0x36, 0x17, 0xfc, 0xd9, 0x45, 0xbd, 0x73, 0xc3, 0x7c, 0xd7, 0x22, 0xdd, 0x37, 0xe8, 0x2e, 0xa1,
	0xb4, 0x12, 0xd3, 0x8c, 0xe8, 0x0c, 0xa4, 0xb7, 0x33, 0xd8, 0x39, 0x3e, 0x18, 0x1d, 0x36, 0xb5,
	0xff, 0x70, 0x0c, 0x55, 0xf1, 0x3a, 0xe8, 0xaa, 0x41, 0xfc, 0x8f, 0xd9, 0x7d, 0x82, 0xf6, 0x25,
	0x29, 0x84, 0xb7, 0x6b, 0x42, 0xf7, 0x9b, 0xda, 0x77, 0x6c, 0x68, 0xf5, 0x35, 0x88, 0x8d, 0xe8,
	0x7e, 0x42, 0x07, 0x40, 0xbf, 0x09, 0xa6, 0x93, 0x8c, 0x7b, 0x7b, 0xc6, 0xf9, 0xb2, 0xa9, 0xfd,
	0x9e, 0x75, 0xb6, 0x52, 0xf0, 0xfb, 0x97, 0xef, 0xa3, 0xc7, 0xd7, 0x97, 0x24, 0xfc, 0xf9, 0x36,
	0xbc, 0x48, 0xae, 0x2e, 0x4f, 0xc2, 0x57, 0x9b, 0x7a, 0x71, 0xf2, 0x6c, 0x78, 0xba, 0x3c, 0x8a,
	0xef, 0x58, 0xfb, 0x19, 0x77, 0x3f, 0x22, 0x27, 0x93, 0x4a, 0x13, 0xc9, 0xc4, 0x8a, 0xbb, 0x6f,
	0xb8, 0x4f, 0x9b, 0xda, 0x77, 0x2d, 0xb7, 0x23, 0xae, 0xc8, 0x3d, 0x74, 0xef, 0x7a, 0x0d, 0xbc,
	0x5a, 0x9c, 0x0e, 0x97, 0x47, 0x31, 0xda, 0x38, 0xce, 0xb8, 0xfb, 0x19, 0xa1, 0xf6, 0xde, 0xca,
	0xbb, 0x35, 0xd8, 0x3b, 0x76, 0x9e, 0x0f, 0xa2, 0xce, 0xef, 0x88, 0xce, 0x37, 0xf2, 0xf6, 0x74,
	0xa3, 0x47, 0x4d, 0xed, 0x3f, 0xb0, 0xd3, 0xb6, 0xe9, 0x20, 0xee, 0xa0, 0x46, 0x1f, 0x2e, 0xde,
	0xdf, 0xc8, 0x73, 0xa0, 0xb7, 0x8d, 0x77, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x7a, 0xb0,
	0xc5, 0xb3, 0x02, 0x00, 0x00,
}
