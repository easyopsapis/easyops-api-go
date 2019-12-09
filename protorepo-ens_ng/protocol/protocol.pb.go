// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.easyops.local/contracts/protorepo-ens_ng/protocol/protocol.proto

package protocol

import (
	fmt "fmt"
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

//Framework 框架类型
type Type int32

const (
	//GRPC 表明契约实现方用 GRPC 框架实现
	Type_GRPC Type = 0
	//REST 表明契约实现方用RESTFul框架实现, 此时 URL 中可以带参数
	Type_REST Type = 1
	//REST2 表明契约实现方用特定的RESTFul框架实现, 此时URL为契约全名, 消息体均序列化到body中
	Type_REST2 Type = 2
)

var Type_name = map[int32]string{
	0: "GRPC",
	1: "REST",
	2: "REST2",
}

var Type_value = map[string]int32{
	"GRPC":  0,
	"REST":  1,
	"REST2": 2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1fd9a835bd78b84, []int{0}
}

func init() {
	proto.RegisterEnum("protocol.Type", Type_name, Type_value)
}

func init() {
	proto.RegisterFile("github.com/easyopsapis/easyops-api-go/protorepo-ens_ng/protocol/protocol.proto", fileDescriptor_f1fd9a835bd78b84)
}

var fileDescriptor_f1fd9a835bd78b84 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcf, 0xd7, 0x4b,
	0x4d, 0x2c, 0xae, 0xcc, 0x2f, 0x28, 0xd6, 0xcb, 0xc9, 0x4f, 0x4e, 0xcc, 0xd1, 0x4f, 0xce, 0xcf,
	0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0x4a, 0x2d, 0xc8,
	0xd7, 0x4d, 0xcd, 0x2b, 0x8e, 0xcf, 0x4b, 0x87, 0x08, 0x24, 0xe7, 0xe7, 0xc0, 0x19, 0x7a, 0x60,
	0x86, 0x10, 0x07, 0x8c, 0xaf, 0xa5, 0xca, 0xc5, 0x12, 0x52, 0x59, 0x90, 0x2a, 0xc4, 0xc1, 0xc5,
	0xe2, 0x1e, 0x14, 0xe0, 0x2c, 0xc0, 0x00, 0x62, 0x05, 0xb9, 0x06, 0x87, 0x08, 0x30, 0x0a, 0x71,
	0x72, 0xb1, 0x82, 0x58, 0x46, 0x02, 0x4c, 0x4e, 0xf6, 0x51, 0xb6, 0xe4, 0xd8, 0x6b, 0x0d, 0x63,
	0x24, 0xb1, 0x81, 0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xa2, 0xa5, 0x65, 0xbf,
	0x00, 0x00, 0x00,
}