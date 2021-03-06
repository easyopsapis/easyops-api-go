// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_instance_with_csv.proto

package instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//ImportInstanceWithCsv返回
type ImportInstanceWithCsvResponse struct {
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
	Data                 []*ImportInstanceWithCsvResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ImportInstanceWithCsvResponse) Reset()         { *m = ImportInstanceWithCsvResponse{} }
func (m *ImportInstanceWithCsvResponse) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceWithCsvResponse) ProtoMessage()    {}
func (*ImportInstanceWithCsvResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4776d0039a9814d4, []int{0}
}
func (m *ImportInstanceWithCsvResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceWithCsvResponse.Unmarshal(m, b)
}
func (m *ImportInstanceWithCsvResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceWithCsvResponse.Marshal(b, m, deterministic)
}
func (m *ImportInstanceWithCsvResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceWithCsvResponse.Merge(m, src)
}
func (m *ImportInstanceWithCsvResponse) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceWithCsvResponse.Size(m)
}
func (m *ImportInstanceWithCsvResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceWithCsvResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceWithCsvResponse proto.InternalMessageInfo

func (m *ImportInstanceWithCsvResponse) GetInsertCount() int32 {
	if m != nil {
		return m.InsertCount
	}
	return 0
}

func (m *ImportInstanceWithCsvResponse) GetUpdateCount() int32 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func (m *ImportInstanceWithCsvResponse) GetFailedCount() int32 {
	if m != nil {
		return m.FailedCount
	}
	return 0
}

func (m *ImportInstanceWithCsvResponse) GetData() []*ImportInstanceWithCsvResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImportInstanceWithCsvResponse_Data struct {
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

func (m *ImportInstanceWithCsvResponse_Data) Reset()         { *m = ImportInstanceWithCsvResponse_Data{} }
func (m *ImportInstanceWithCsvResponse_Data) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceWithCsvResponse_Data) ProtoMessage()    {}
func (*ImportInstanceWithCsvResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_4776d0039a9814d4, []int{0, 0}
}
func (m *ImportInstanceWithCsvResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceWithCsvResponse_Data.Unmarshal(m, b)
}
func (m *ImportInstanceWithCsvResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceWithCsvResponse_Data.Marshal(b, m, deterministic)
}
func (m *ImportInstanceWithCsvResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceWithCsvResponse_Data.Merge(m, src)
}
func (m *ImportInstanceWithCsvResponse_Data) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceWithCsvResponse_Data.Size(m)
}
func (m *ImportInstanceWithCsvResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceWithCsvResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceWithCsvResponse_Data proto.InternalMessageInfo

func (m *ImportInstanceWithCsvResponse_Data) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportInstanceWithCsvResponse_Data) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportInstanceWithCsvResponse_Data) GetData() []*types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

//
//ImportInstanceWithCsvApi返回
type ImportInstanceWithCsvResponseWrapper struct {
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
	Data                 *ImportInstanceWithCsvResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ImportInstanceWithCsvResponseWrapper) Reset()         { *m = ImportInstanceWithCsvResponseWrapper{} }
func (m *ImportInstanceWithCsvResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ImportInstanceWithCsvResponseWrapper) ProtoMessage()    {}
func (*ImportInstanceWithCsvResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_4776d0039a9814d4, []int{1}
}
func (m *ImportInstanceWithCsvResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportInstanceWithCsvResponseWrapper.Unmarshal(m, b)
}
func (m *ImportInstanceWithCsvResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportInstanceWithCsvResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ImportInstanceWithCsvResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportInstanceWithCsvResponseWrapper.Merge(m, src)
}
func (m *ImportInstanceWithCsvResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ImportInstanceWithCsvResponseWrapper.Size(m)
}
func (m *ImportInstanceWithCsvResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportInstanceWithCsvResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ImportInstanceWithCsvResponseWrapper proto.InternalMessageInfo

func (m *ImportInstanceWithCsvResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ImportInstanceWithCsvResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ImportInstanceWithCsvResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ImportInstanceWithCsvResponseWrapper) GetData() *ImportInstanceWithCsvResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportInstanceWithCsvResponse)(nil), "instance.ImportInstanceWithCsvResponse")
	proto.RegisterType((*ImportInstanceWithCsvResponse_Data)(nil), "instance.ImportInstanceWithCsvResponse.Data")
	proto.RegisterType((*ImportInstanceWithCsvResponseWrapper)(nil), "instance.ImportInstanceWithCsvResponseWrapper")
}

func init() { proto.RegisterFile("import_instance_with_csv.proto", fileDescriptor_4776d0039a9814d4) }

var fileDescriptor_4776d0039a9814d4 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x4f, 0xfa, 0x30,
	0x1c, 0xcd, 0xfe, 0x1b, 0xff, 0x68, 0x47, 0xa2, 0x99, 0x89, 0x10, 0xa2, 0x8e, 0x54, 0xa3, 0x1c,
	0xb4, 0x24, 0x78, 0x31, 0xc4, 0x13, 0xe8, 0x81, 0xc4, 0x8b, 0xf5, 0xc0, 0x71, 0x29, 0x5b, 0x19,
	0x4d, 0x60, 0x5d, 0xba, 0x0e, 0xfd, 0x1a, 0x7e, 0x1d, 0x3f, 0xcc, 0x3e, 0xc4, 0xae, 0x5e, 0xcc,
	0xda, 0x01, 0x0b, 0x09, 0x84, 0xd3, 0xf6, 0xfa, 0xde, 0x6b, 0x7f, 0xef, 0xf7, 0xc0, 0x15, 0x5b,
	0xc4, 0x5c, 0x48, 0x8f, 0x45, 0x89, 0x24, 0x91, 0x4f, 0xbd, 0x4f, 0x26, 0x67, 0x9e, 0x9f, 0x2c,
	0x51, 0x2c, 0xb8, 0xe4, 0xce, 0xd1, 0x8a, 0x68, 0x3d, 0x84, 0x4c, 0xce, 0xd2, 0x09, 0xf2, 0xf9,
	0xa2, 0x1b, 0xf2, 0x90, 0x77, 0x95, 0x60, 0x92, 0x4e, 0x15, 0x52, 0x40, 0xfd, 0x69, 0x63, 0xeb,
	0x22, 0xe4, 0x3c, 0x9c, 0xd3, 0x8d, 0x2a, 0x91, 0x22, 0xf5, 0xa5, 0x66, 0xe1, 0x8f, 0x09, 0x2e,
	0x47, 0xea, 0xe5, 0x51, 0x79, 0xff, 0x98, 0xc9, 0xd9, 0x30, 0x59, 0x62, 0x9a, 0xc4, 0x3c, 0x4a,
	0xa8, 0xd3, 0x07, 0x75, 0x16, 0x25, 0x54, 0x48, 0xcf, 0xe7, 0x69, 0x24, 0x9b, 0x46, 0xdb, 0xe8,
	0xd4, 0x06, 0x8d, 0x3c, 0x73, 0xcf, 0xa6, 0x5c, 0x2c, 0xfa, 0xb0, 0xca, 0x42, 0x6c, 0x6b, 0x38,
	0x2c, 0x50, 0xe1, 0x4d, 0xe3, 0x80, 0x48, 0x5a, 0x7a, 0xff, 0x6d, 0x7b, 0xab, 0x2c, 0xc4, 0xb6,
	0x86, 0x6b, 0xef, 0x94, 0xb0, 0x39, 0x0d, 0x4a, 0xaf, 0xb9, 0xed, 0xad, 0xb2, 0x10, 0xdb, 0x1a,
	0x6a, 0xef, 0x3b, 0xb0, 0x02, 0x22, 0x49, 0xd3, 0x6a, 0x9b, 0x1d, 0xbb, 0x77, 0x8f, 0x56, 0xbb,
	0x43, 0x7b, 0xa3, 0xa2, 0x17, 0x22, 0xc9, 0xe0, 0x24, 0xcf, 0x5c, 0x5b, 0xbf, 0x50, 0xdc, 0x01,
	0xb1, 0xba, 0xaa, 0xf5, 0x6d, 0x00, 0xab, 0xe0, 0x9d, 0x6b, 0x60, 0xf9, 0x3c, 0xa0, 0xe5, 0x1e,
	0x2a, 0xea, 0xe2, 0x14, 0x62, 0x45, 0x3a, 0xb7, 0xa0, 0x46, 0x85, 0xe0, 0x42, 0x25, 0x3e, 0x1e,
	0x9c, 0xe6, 0x99, 0x5b, 0xd7, 0x2a, 0x75, 0x0c, 0xb1, 0xa6, 0x9d, 0xe7, 0x72, 0x50, 0x53, 0x0d,
	0xda, 0x40, 0xba, 0x2b, 0xb4, 0xea, 0x0a, 0x7d, 0xa8, 0xae, 0x76, 0xcc, 0x04, 0x7f, 0x0d, 0x70,
	0xb3, 0x37, 0xd1, 0x58, 0x90, 0x38, 0xa6, 0xe2, 0xb0, 0x99, 0x9f, 0x80, 0x5d, 0x7c, 0x5f, 0xbf,
	0xe2, 0x39, 0x61, 0x51, 0x39, 0xf9, 0x79, 0x9e, 0xb9, 0xce, 0x46, 0x5b, 0x92, 0x10, 0x57, 0xa5,
	0x9b, 0xb4, 0xe6, 0xfe, 0xb4, 0x6f, 0xeb, 0x5a, 0x8c, 0x8e, 0xdd, 0xbb, 0x3b, 0xb0, 0x96, 0x1d,
	0xe9, 0x27, 0xff, 0xd5, 0x96, 0x1e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x19, 0x7d, 0xb4,
	0x3a, 0x03, 0x00, 0x00,
}
