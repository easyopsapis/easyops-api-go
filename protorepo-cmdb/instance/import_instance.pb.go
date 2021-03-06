// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_instance.proto

package instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//ImportInstance请求
type ImportInstanceRequest struct {
	//
	//实例所属的模型ID
	ObjectId string `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	//
	//导入实例的字段
	Keys []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys" form:"keys"`
	//
	//导入实例数据列表
	Datas                []*types.Struct `protobuf:"bytes,3,rep,name=datas,proto3" json:"datas" form:"datas"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ImportInstanceRequest) Reset()         { *m = ImportInstanceRequest{} }
func (m *ImportInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceRequest) ProtoMessage()    {}
func (*ImportInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a4e0cc7aafbb298, []int{0}
}
func (m *ImportInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceRequest.Unmarshal(m, b)
}
func (m *ImportInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceRequest.Marshal(b, m, deterministic)
}
func (m *ImportInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceRequest.Merge(m, src)
}
func (m *ImportInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceRequest.Size(m)
}
func (m *ImportInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceRequest proto.InternalMessageInfo

func (m *ImportInstanceRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *ImportInstanceRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ImportInstanceRequest) GetDatas() []*types.Struct {
	if m != nil {
		return m.Datas
	}
	return nil
}

//
//ImportInstance返回
type ImportInstanceResponse struct {
	//
	//成功插入数量
	InsertCount int32 `protobuf:"varint,1,opt,name=insert_count,json=insertCount,proto3" json:"insert_count" form:"insert_count"`
	//
	//成功更新数量
	UpdateCount int32 `protobuf:"varint,2,opt,name=update_count,json=updateCount,proto3" json:"update_count" form:"update_count"`
	//
	//失败数量
	FailedCount int32 `protobuf:"varint,3,opt,name=failed_count,json=failedCount,proto3" json:"failed_count" form:"failed_count"`
	//
	//失败数据
	Data                 []*ImportInstanceResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ImportInstanceResponse) Reset()         { *m = ImportInstanceResponse{} }
func (m *ImportInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceResponse) ProtoMessage()    {}
func (*ImportInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a4e0cc7aafbb298, []int{1}
}
func (m *ImportInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceResponse.Unmarshal(m, b)
}
func (m *ImportInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceResponse.Marshal(b, m, deterministic)
}
func (m *ImportInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceResponse.Merge(m, src)
}
func (m *ImportInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceResponse.Size(m)
}
func (m *ImportInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceResponse proto.InternalMessageInfo

func (m *ImportInstanceResponse) GetInsertCount() int32 {
	if m != nil {
		return m.InsertCount
	}
	return 0
}

func (m *ImportInstanceResponse) GetUpdateCount() int32 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func (m *ImportInstanceResponse) GetFailedCount() int32 {
	if m != nil {
		return m.FailedCount
	}
	return 0
}

func (m *ImportInstanceResponse) GetData() []*ImportInstanceResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImportInstanceResponse_Data struct {
	//
	//错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	//
	//错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error" form:"error"`
	//
	//错误数据
	Data                 []*types.Struct `protobuf:"bytes,3,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ImportInstanceResponse_Data) Reset()         { *m = ImportInstanceResponse_Data{} }
func (m *ImportInstanceResponse_Data) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceResponse_Data) ProtoMessage()    {}
func (*ImportInstanceResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a4e0cc7aafbb298, []int{1, 0}
}
func (m *ImportInstanceResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceResponse_Data.Unmarshal(m, b)
}
func (m *ImportInstanceResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceResponse_Data.Marshal(b, m, deterministic)
}
func (m *ImportInstanceResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceResponse_Data.Merge(m, src)
}
func (m *ImportInstanceResponse_Data) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceResponse_Data.Size(m)
}
func (m *ImportInstanceResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceResponse_Data proto.InternalMessageInfo

func (m *ImportInstanceResponse_Data) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportInstanceResponse_Data) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportInstanceResponse_Data) GetData() []*types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//ImportInstanceApi返回
type ImportInstanceResponseWrapper struct {
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
	Data                 *ImportInstanceResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ImportInstanceResponseWrapper) Reset()         { *m = ImportInstanceResponseWrapper{} }
func (m *ImportInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceResponseWrapper) ProtoMessage()    {}
func (*ImportInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a4e0cc7aafbb298, []int{2}
}
func (m *ImportInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *ImportInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ImportInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceResponseWrapper.Merge(m, src)
}
func (m *ImportInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceResponseWrapper.Size(m)
}
func (m *ImportInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceResponseWrapper proto.InternalMessageInfo

func (m *ImportInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ImportInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportInstanceResponseWrapper) GetData() *ImportInstanceResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportInstanceRequest)(nil), "instance.ImportInstanceRequest")
	proto.RegisterType((*ImportInstanceResponse)(nil), "instance.ImportInstanceResponse")
	proto.RegisterType((*ImportInstanceResponse_Data)(nil), "instance.ImportInstanceResponse.Data")
	proto.RegisterType((*ImportInstanceResponseWrapper)(nil), "instance.ImportInstanceResponseWrapper")
}

func init() { proto.RegisterFile("import_instance.proto", fileDescriptor_3a4e0cc7aafbb298) }

var fileDescriptor_3a4e0cc7aafbb298 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0xeb, 0x14, 0x35, 0xe3, 0x4a, 0x45, 0x46, 0x6d, 0xa3, 0x88, 0xca, 0xd1, 0xf0, 0x50,
	0x16, 0xd8, 0x29, 0xad, 0x54, 0x41, 0x85, 0x84, 0x08, 0x74, 0x11, 0x96, 0xc3, 0x02, 0x89, 0x0a,
	0xa2, 0x89, 0x3d, 0x31, 0xa6, 0x89, 0xaf, 0x99, 0x19, 0x53, 0x1e, 0xe2, 0x23, 0xf8, 0x0f, 0x7e,
	0x86, 0x1f, 0x30, 0x12, 0x7b, 0x36, 0xfe, 0x02, 0x34, 0x33, 0x8e, 0x63, 0x45, 0x81, 0xb2, 0xf2,
	0x9c, 0x7b, 0xee, 0xb9, 0x3a, 0xe7, 0x5e, 0xa3, 0xdd, 0x64, 0x9e, 0x01, 0x97, 0xe3, 0x24, 0x15,
	0x92, 0xa6, 0x21, 0x0b, 0x32, 0x0e, 0x12, 0xdc, 0xad, 0x05, 0xee, 0xfa, 0x71, 0x22, 0xdf, 0xe6,
	0x93, 0x20, 0x84, 0xf9, 0x20, 0x86, 0x18, 0x06, 0xba, 0x61, 0x92, 0x4f, 0x35, 0xd2, 0x40, 0xbf,
	0x8c, 0xb0, 0x7b, 0xd2, 0x68, 0x9f, 0x5f, 0x26, 0xf2, 0x02, 0x2e, 0x07, 0x31, 0xf8, 0x9a, 0xf4,
	0x3f, 0xd0, 0x59, 0x12, 0x51, 0x09, 0x5c, 0x0c, 0xea, 0x67, 0xa5, 0xbb, 0x19, 0x03, 0xc4, 0x33,
	0xb6, 0x9c, 0x2e, 0x24, 0xcf, 0x43, 0x69, 0x58, 0xfc, 0xc3, 0x42, 0xbb, 0x23, 0x6d, 0x74, 0x54,
	0xf9, 0x22, 0xec, 0x7d, 0xce, 0x84, 0x74, 0x09, 0xda, 0x82, 0xc9, 0x3b, 0x16, 0xca, 0x51, 0xd4,
	0xb1, 0x7a, 0x56, 0xbf, 0x3d, 0x3c, 0x29, 0x0b, 0x6f, 0x67, 0x0a, 0x7c, 0x7e, 0x8a, 0x17, 0x0c,
	0xfe, 0xf5, 0xd3, 0xf3, 0xd0, 0xc1, 0x9b, 0x73, 0xea, 0x7f, 0x7e, 0xe2, 0xbf, 0x1a, 0xbf, 0x3e,
	0x3f, 0xf4, 0x1f, 0x2e, 0xde, 0x5f, 0x0e, 0xef, 0x1d, 0xdf, 0xff, 0x7a, 0x9b, 0xd4, 0x73, 0xdc,
	0x5b, 0xa8, 0x75, 0xc1, 0x3e, 0x89, 0xce, 0x46, 0xcf, 0xee, 0xb7, 0x87, 0x3b, 0x65, 0xe1, 0x39,
	0x66, 0x9e, 0xaa, 0x62, 0xa2, 0x49, 0xf7, 0x31, 0xda, 0x8c, 0xa8, 0xa4, 0xa2, 0x63, 0xf7, 0xec,
	0xbe, 0x73, 0xb4, 0x1f, 0x98, 0x00, 0xc1, 0x22, 0x40, 0xf0, 0x42, 0x07, 0x18, 0x5e, 0x2f, 0x0b,
	0x6f, 0xdb, 0xc8, 0x75, 0x3f, 0x26, 0x46, 0x87, 0xbf, 0xdb, 0x68, 0x6f, 0x35, 0x93, 0xc8, 0x20,
	0x15, 0xcc, 0x3d, 0x45, 0xdb, 0x49, 0x2a, 0x18, 0x97, 0xe3, 0x10, 0xf2, 0x54, 0xea, 0x60, 0x9b,
	0xc3, 0xfd, 0xb2, 0xf0, 0x6e, 0x98, 0x49, 0x4d, 0x16, 0x13, 0xc7, 0xc0, 0xa7, 0x0a, 0x29, 0x6d,
	0x9e, 0x45, 0x54, 0xb2, 0x4a, 0xbb, 0xb1, 0xaa, 0x6d, 0xb2, 0x98, 0x38, 0x06, 0xd6, 0xda, 0x29,
	0x4d, 0x66, 0x2c, 0xaa, 0xb4, 0xf6, 0xaa, 0xb6, 0xc9, 0x62, 0xe2, 0x18, 0x68, 0xb4, 0xcf, 0x51,
	0x4b, 0xe5, 0xea, 0xb4, 0xf4, 0x3a, 0xee, 0x04, 0xf5, 0x0f, 0xb5, 0x3e, 0x63, 0xf0, 0x8c, 0x4a,
	0xda, 0xdc, 0xad, 0x12, 0x63, 0xa2, 0x67, 0x74, 0xbf, 0x59, 0xa8, 0xa5, 0x78, 0x75, 0x89, 0x10,
	0x22, 0x56, 0x2d, 0xa0, 0xd1, 0xad, 0xaa, 0x98, 0x68, 0xd2, 0xbd, 0x8b, 0x36, 0x19, 0xe7, 0xc0,
	0x75, 0xd4, 0x76, 0x73, 0xe1, 0xba, 0x8c, 0x89, 0xa1, 0xdd, 0x47, 0x95, 0xc3, 0x2b, 0x0e, 0xb6,
	0xde, 0x13, 0xfe, 0x6d, 0xa1, 0x83, 0xf5, 0x51, 0x5e, 0x72, 0x9a, 0x65, 0x8c, 0xff, 0x9f, 0xd9,
	0x07, 0xc8, 0x51, 0xdf, 0xb3, 0x8f, 0xd9, 0x8c, 0x26, 0x69, 0x65, 0x79, 0xaf, 0x2c, 0x3c, 0x77,
	0xd9, 0x5b, 0x91, 0x98, 0x34, 0x5b, 0x97, 0x31, 0xed, 0x7f, 0xc7, 0x3c, 0xab, 0x0f, 0x61, 0xf5,
	0x9d, 0xa3, 0xde, 0x55, 0x87, 0xf8, 0x4b, 0xde, 0xc9, 0x35, 0xbd, 0x97, 0xe3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0xa6, 0xe3, 0x4c, 0x21, 0x04, 0x00, 0x00,
}
