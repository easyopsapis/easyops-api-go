// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app_health_config.proto

package monitor_config

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
//应用评分配置
type AppHealthConfig struct {
	//
	//应用id
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id" form:"app_id"`
	//
	//配置id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//版本id
	VersionId string `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id" form:"version_id"`
	//
	//创建者
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime int32 `protobuf:"varint,5,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改者
	Modifier string `protobuf:"bytes,6,opt,name=modifier,proto3" json:"modifier" form:"modifier"`
	//
	//修改时间
	Mtime int32 `protobuf:"varint,7,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//org
	Org int32 `protobuf:"varint,8,opt,name=org,proto3" json:"org" form:"org"`
	//
	//应用变量
	Variables []*AppHealthConfigVariable `protobuf:"bytes,9,rep,name=variables,proto3" json:"variables" form:"variables"`
	//
	//应用分层
	Layers               []*AppHealthConfigLayer `protobuf:"bytes,10,rep,name=layers,proto3" json:"layers" form:"layers"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AppHealthConfig) Reset()         { *m = AppHealthConfig{} }
func (m *AppHealthConfig) String() string { return proto.CompactTextString(m) }
func (*AppHealthConfig) ProtoMessage()    {}
func (*AppHealthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a2eac695befb197, []int{0}
}
func (m *AppHealthConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppHealthConfig.Unmarshal(m, b)
}
func (m *AppHealthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppHealthConfig.Marshal(b, m, deterministic)
}
func (m *AppHealthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppHealthConfig.Merge(m, src)
}
func (m *AppHealthConfig) XXX_Size() int {
	return xxx_messageInfo_AppHealthConfig.Size(m)
}
func (m *AppHealthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AppHealthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AppHealthConfig proto.InternalMessageInfo

func (m *AppHealthConfig) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AppHealthConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AppHealthConfig) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *AppHealthConfig) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AppHealthConfig) GetCtime() int32 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *AppHealthConfig) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

func (m *AppHealthConfig) GetMtime() int32 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *AppHealthConfig) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *AppHealthConfig) GetVariables() []*AppHealthConfigVariable {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *AppHealthConfig) GetLayers() []*AppHealthConfigLayer {
	if m != nil {
		return m.Layers
	}
	return nil
}

func init() {
	proto.RegisterType((*AppHealthConfig)(nil), "monitor_config.AppHealthConfig")
}

func init() { proto.RegisterFile("app_health_config.proto", fileDescriptor_9a2eac695befb197) }

var fileDescriptor_9a2eac695befb197 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xb5, 0xbb, 0xdd, 0x6d, 0x63, 0xe8, 0xb6, 0x35, 0x20, 0xa2, 0x4a, 0x28, 0x2b, 0x0b,
	0xc1, 0x1e, 0x68, 0x22, 0x01, 0x27, 0x6e, 0x2c, 0x97, 0x16, 0x21, 0x21, 0xe5, 0xc0, 0x01, 0x0e,
	0x2b, 0x6f, 0xe2, 0x64, 0x2d, 0x25, 0x19, 0xcb, 0x49, 0x2b, 0xf5, 0x29, 0x78, 0x43, 0x3f, 0x84,
	0x9f, 0x00, 0x79, 0xec, 0x00, 0x2d, 0x07, 0x2e, 0x70, 0x8a, 0x67, 0xfe, 0xef, 0xff, 0x27, 0xf2,
	0x98, 0x3c, 0xe5, 0x4a, 0x6d, 0xf7, 0x82, 0x37, 0xc3, 0x7e, 0x5b, 0x40, 0x57, 0xc9, 0x3a, 0x55,
	0x1a, 0x06, 0xa0, 0xcb, 0x16, 0x3a, 0x39, 0x80, 0x0e, 0xdd, 0xf3, 0x8b, 0x5a, 0x0e, 0xfb, 0xeb,
	0x5d, 0x5a, 0x40, 0x9b, 0xd5, 0x50, 0x43, 0x86, 0xd8, 0xee, 0xba, 0xc2, 0x0a, 0x0b, 0x3c, 0x79,
	0xfb, 0xb9, 0xac, 0x21, 0x15, 0xbc, 0xbf, 0x05, 0xd5, 0xa7, 0x0d, 0x14, 0xbc, 0xc9, 0x0a, 0xe8,
	0x06, 0xcd, 0x8b, 0xa1, 0xf7, 0x4e, 0x2d, 0x14, 0x5c, 0xb4, 0x50, 0x8a, 0xa6, 0xcf, 0x02, 0x98,
	0x61, 0x99, 0xdd, 0x1d, 0x9c, 0xfd, 0xf1, 0x83, 0xdb, 0x1b, 0xae, 0x25, 0xdf, 0x35, 0x22, 0x8c,
	0xaa, 0xfe, 0xe3, 0xa8, 0x86, 0xdf, 0x0a, 0xed, 0xe7, 0xb0, 0xef, 0x07, 0xe4, 0xe4, 0xbd, 0x52,
	0x97, 0x08, 0x7c, 0x40, 0x9d, 0xae, 0xc9, 0xc2, 0x99, 0x64, 0x19, 0x4f, 0x56, 0x93, 0x75, 0xb4,
	0x39, 0xb3, 0x26, 0x39, 0xae, 0x40, 0xb7, 0xef, 0x98, 0xef, 0xb3, 0x7c, 0xce, 0x95, 0xba, 0x2a,
	0xe9, 0x33, 0x32, 0x95, 0x65, 0x3c, 0x45, 0xea, 0xd8, 0x9a, 0x24, 0xf2, 0x94, 0x23, 0xa6, 0xb2,
	0xa4, 0x6f, 0x09, 0xb9, 0x11, 0xba, 0x97, 0xd0, 0xb9, 0xb0, 0x19, 0x62, 0x4f, 0xac, 0x49, 0xce,
	0x3c, 0xf6, 0x4b, 0x63, 0x79, 0x14, 0x8a, 0xab, 0x92, 0xbe, 0x22, 0x87, 0x85, 0x16, 0x7c, 0x00,
	0x1d, 0x1f, 0xa0, 0x85, 0x5a, 0x93, 0x2c, 0xbd, 0x25, 0x08, 0x2c, 0x1f, 0x11, 0xfa, 0x82, 0xcc,
	0x8b, 0x41, 0xb6, 0x22, 0x9e, 0xaf, 0x26, 0xeb, 0xf9, 0xe6, 0xd4, 0x9a, 0xe4, 0x61, 0x60, 0x5d,
	0x9b, 0xe5, 0x5e, 0xa6, 0x19, 0x39, 0x6a, 0xa1, 0x94, 0x95, 0x14, 0x3a, 0x5e, 0x60, 0xec, 0x23,
	0x6b, 0x92, 0x13, 0x8f, 0x8e, 0x0a, 0xcb, 0x7f, 0x42, 0x2e, 0xb8, 0xc5, 0xe0, 0xc3, 0xfb, 0xc1,
	0x6d, 0x08, 0xc6, 0x2f, 0x5d, 0x91, 0x19, 0xe8, 0x3a, 0x3e, 0x42, 0x6a, 0x69, 0x4d, 0x42, 0x3c,
	0x05, 0xba, 0x66, 0xb9, 0x93, 0xe8, 0x37, 0x12, 0x8d, 0xdb, 0xed, 0xe3, 0x68, 0x35, 0x5b, 0x3f,
	0x78, 0xfd, 0x32, 0xbd, 0xbb, 0xa5, 0xf4, 0xde, 0x0e, 0xbe, 0x04, 0x7e, 0xf3, 0xd8, 0x9a, 0xe4,
	0x34, 0x5c, 0xd7, 0x98, 0xe1, 0x6e, 0x6b, 0x3c, 0xd3, 0xcf, 0x64, 0x81, 0xfb, 0xec, 0x63, 0x82,
	0xc9, 0xcf, 0xff, 0x92, 0xfc, 0xc9, 0xc1, 0xbf, 0xaf, 0xd4, 0xbb, 0x59, 0x1e, 0x62, 0x36, 0x1f,
	0xbf, 0x5e, 0xfe, 0xab, 0xb7, 0xb7, 0x5b, 0x20, 0xfe, 0xe6, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x24, 0x35, 0x13, 0xf2, 0x91, 0x03, 0x00, 0x00,
}