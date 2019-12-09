// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_check.proto

package cmdb_object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	cmdb "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
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
//ImportCheck请求
type ImportCheckRequest struct {
	//
	//模型列表
	Body                 []*cmdb.ObjectImport `protobuf:"bytes,1,rep,name=body,proto3" json:"body" form:"body"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ImportCheckRequest) Reset()         { *m = ImportCheckRequest{} }
func (m *ImportCheckRequest) String() string { return proto.CompactTextString(m) }
func (*ImportCheckRequest) ProtoMessage()    {}
func (*ImportCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13add06972238eca, []int{0}
}
func (m *ImportCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportCheckRequest.Unmarshal(m, b)
}
func (m *ImportCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportCheckRequest.Marshal(b, m, deterministic)
}
func (m *ImportCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportCheckRequest.Merge(m, src)
}
func (m *ImportCheckRequest) XXX_Size() int {
	return xxx_messageInfo_ImportCheckRequest.Size(m)
}
func (m *ImportCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportCheckRequest proto.InternalMessageInfo

func (m *ImportCheckRequest) GetBody() []*cmdb.ObjectImport {
	if m != nil {
		return m.Body
	}
	return nil
}

//
//ImportCheck返回
type ImportCheckResponse struct {
	//
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//返回消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" form:"message"`
	//
	//'{ "APP": { "duplicated_relations": [ {"relation_id": "APP_cluster_CLUSTER", "left_id": "clusters", "right_id": "app"} ], "cannot_created_relations": [ {"relation_id": "APP_cluster_CLUSTER", "left_id": "clusters", "right_id": "app"} ], "cannot_created_references": [ {"id":"cpu","name":"CPU信息","value":{"type":"FK","external":[{"org_attr":"name","name":"名称"}],"rule":{"obj":"HOST_CPU","mode":"inner"}}} ], "id_is_duplicated": true, "name_is_duplicated": true } }'
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ImportCheckResponse) Reset()         { *m = ImportCheckResponse{} }
func (m *ImportCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ImportCheckResponse) ProtoMessage()    {}
func (*ImportCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13add06972238eca, []int{1}
}
func (m *ImportCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportCheckResponse.Unmarshal(m, b)
}
func (m *ImportCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportCheckResponse.Marshal(b, m, deterministic)
}
func (m *ImportCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportCheckResponse.Merge(m, src)
}
func (m *ImportCheckResponse) XXX_Size() int {
	return xxx_messageInfo_ImportCheckResponse.Size(m)
}
func (m *ImportCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportCheckResponse proto.InternalMessageInfo

func (m *ImportCheckResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportCheckResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportCheckResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ImportCheckResponse) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//ImportCheckApi返回
type ImportCheckResponseWrapper struct {
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
	Data                 *ImportCheckResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ImportCheckResponseWrapper) Reset()         { *m = ImportCheckResponseWrapper{} }
func (m *ImportCheckResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ImportCheckResponseWrapper) ProtoMessage()    {}
func (*ImportCheckResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_13add06972238eca, []int{2}
}
func (m *ImportCheckResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportCheckResponseWrapper.Unmarshal(m, b)
}
func (m *ImportCheckResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportCheckResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ImportCheckResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportCheckResponseWrapper.Merge(m, src)
}
func (m *ImportCheckResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ImportCheckResponseWrapper.Size(m)
}
func (m *ImportCheckResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportCheckResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ImportCheckResponseWrapper proto.InternalMessageInfo

func (m *ImportCheckResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportCheckResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ImportCheckResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportCheckResponseWrapper) GetData() *ImportCheckResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportCheckRequest)(nil), "cmdb_object.ImportCheckRequest")
	proto.RegisterType((*ImportCheckResponse)(nil), "cmdb_object.ImportCheckResponse")
	proto.RegisterType((*ImportCheckResponseWrapper)(nil), "cmdb_object.ImportCheckResponseWrapper")
}

func init() { proto.RegisterFile("import_check.proto", fileDescriptor_13add06972238eca) }

var fileDescriptor_13add06972238eca = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x0b, 0xd3, 0x30,
	0x14, 0xc7, 0xa9, 0xeb, 0x14, 0x53, 0x51, 0x89, 0xa0, 0x65, 0x08, 0x2d, 0x11, 0xa4, 0x07, 0x97,
	0xc2, 0x3c, 0x28, 0xe2, 0x69, 0xb2, 0x83, 0x07, 0x11, 0xe2, 0xc1, 0xe3, 0x48, 0xd3, 0xac, 0x9b,
	0xb6, 0x7b, 0x31, 0x49, 0xc1, 0x7d, 0x49, 0x3f, 0x42, 0x3f, 0x80, 0xc7, 0x7e, 0x02, 0x69, 0xd2,
	0xb1, 0x0e, 0xa6, 0x78, 0xda, 0x5e, 0xde, 0xef, 0xff, 0xde, 0xfb, 0xff, 0x29, 0xc2, 0x87, 0x46,
	0x81, 0xb6, 0x5b, 0xb1, 0x97, 0xe2, 0x3b, 0x55, 0x1a, 0x2c, 0xe0, 0x48, 0x34, 0x65, 0xb1, 0x85,
	0xe2, 0x9b, 0x14, 0x76, 0xb1, 0xac, 0x0e, 0x76, 0xdf, 0x16, 0x54, 0x40, 0x93, 0x57, 0x50, 0x41,
	0xee, 0x98, 0xa2, 0xdd, 0xb9, 0xca, 0x15, 0xee, 0x9f, 0xd7, 0x2e, 0x58, 0x05, 0x54, 0x72, 0x73,
	0x02, 0x65, 0x68, 0x0d, 0x82, 0xd7, 0xb9, 0x80, 0xa3, 0xd5, 0x5c, 0x58, 0xe3, 0x95, 0x5a, 0x2a,
	0x58, 0x36, 0x50, 0xca, 0xda, 0xe4, 0x23, 0x98, 0xbb, 0x32, 0x1f, 0xb6, 0xe6, 0x7e, 0xeb, 0xd6,
	0x5f, 0x35, 0xce, 0x7c, 0x5e, 0x01, 0x54, 0xb5, 0xbc, 0x6c, 0x36, 0x56, 0xb7, 0x62, 0xec, 0x92,
	0x4f, 0x08, 0x7f, 0x74, 0xf4, 0x87, 0xc1, 0x02, 0x93, 0x3f, 0x5a, 0x69, 0x2c, 0x7e, 0x83, 0xc2,
	0x02, 0xca, 0x53, 0x1c, 0xa4, 0xb3, 0x2c, 0x5a, 0x61, 0x3a, 0x0c, 0xa7, 0x9f, 0xdd, 0x70, 0x4f,
	0xaf, 0x1f, 0xf5, 0x5d, 0x12, 0xed, 0x40, 0x37, 0xef, 0xc8, 0x40, 0x12, 0xe6, 0x04, 0xe4, 0x57,
	0x80, 0x9e, 0x5c, 0xcd, 0x33, 0x0a, 0x8e, 0x46, 0xe2, 0x17, 0x28, 0x14, 0x50, 0xca, 0x38, 0x48,
	0x83, 0x6c, 0x3e, 0x15, 0x0f, 0xaf, 0x84, 0xb9, 0x26, 0x7e, 0x89, 0xe6, 0x52, 0x6b, 0xd0, 0xf1,
	0x9d, 0x34, 0xc8, 0xee, 0xaf, 0x1f, 0xf7, 0x5d, 0xf2, 0xc0, 0x53, 0xee, 0x99, 0x30, 0xdf, 0xc6,
	0xaf, 0xd0, 0xbd, 0x46, 0x1a, 0xc3, 0x2b, 0x19, 0xcf, 0x1c, 0x89, 0xfb, 0x2e, 0x79, 0xe8, 0xc9,
	0xb1, 0x41, 0xd8, 0x19, 0xc1, 0xef, 0x51, 0x58, 0x72, 0xcb, 0xe3, 0x30, 0x0d, 0xb2, 0x68, 0xf5,
	0x8c, 0xfa, 0x38, 0xe8, 0x39, 0x0e, 0xfa, 0xc5, 0xc5, 0x31, 0xbd, 0x69, 0xc0, 0x09, 0x73, 0x2a,
	0xf2, 0x3b, 0x40, 0x8b, 0x1b, 0x86, 0xbe, 0x6a, 0xae, 0x94, 0xd4, 0xff, 0xe7, 0xeb, 0x2d, 0x8a,
	0x86, 0xdf, 0xcd, 0x4f, 0x55, 0xf3, 0xc3, 0x71, 0x74, 0xf7, 0xb4, 0xef, 0x12, 0x7c, 0x61, 0xc7,
	0x26, 0x61, 0x53, 0xf4, 0x92, 0xc8, 0xec, 0xdf, 0x89, 0x6c, 0xae, 0x3c, 0xa6, 0x74, 0xf2, 0x09,
	0xd2, 0x1b, 0xd7, 0xff, 0xc5, 0x6c, 0x71, 0xd7, 0x85, 0xf2, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0xe4, 0x17, 0xb2, 0xd7, 0x02, 0x00, 0x00,
}