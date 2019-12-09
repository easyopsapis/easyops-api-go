// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package sqlpkgs

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	database_delivery "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/database_delivery"
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
//CreateSQLPackage请求
type CreateSQLPackageRequest struct {
	//
	//创建SQL包body
	CreateSqlpkg         *CreateSQLPackageRequest_CreateSqlpkg `protobuf:"bytes,1,opt,name=createSqlpkg,proto3" json:"createSqlpkg" form:"createSqlpkg"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CreateSQLPackageRequest) Reset()         { *m = CreateSQLPackageRequest{} }
func (m *CreateSQLPackageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageRequest) ProtoMessage()    {}
func (*CreateSQLPackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateSQLPackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageRequest.Unmarshal(m, b)
}
func (m *CreateSQLPackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageRequest.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageRequest.Merge(m, src)
}
func (m *CreateSQLPackageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageRequest.Size(m)
}
func (m *CreateSQLPackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageRequest proto.InternalMessageInfo

func (m *CreateSQLPackageRequest) GetCreateSqlpkg() *CreateSQLPackageRequest_CreateSqlpkg {
	if m != nil {
		return m.CreateSqlpkg
	}
	return nil
}

type CreateSQLPackageRequest_CreateSqlpkg struct {
	//
	//名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
	//
	//SQL包的数据库类型
	DbType string `protobuf:"bytes,2,opt,name=dbType,proto3" json:"dbType" form:"dbType"`
	//
	//描述
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//程序包Id
	AppPackageId         string   `protobuf:"bytes,4,opt,name=appPackageId,proto3" json:"appPackageId" form:"appPackageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageRequest_CreateSqlpkg) Reset()         { *m = CreateSQLPackageRequest_CreateSqlpkg{} }
func (m *CreateSQLPackageRequest_CreateSqlpkg) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageRequest_CreateSqlpkg) ProtoMessage()    {}
func (*CreateSQLPackageRequest_CreateSqlpkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 0}
}
func (m *CreateSQLPackageRequest_CreateSqlpkg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg.Unmarshal(m, b)
}
func (m *CreateSQLPackageRequest_CreateSqlpkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageRequest_CreateSqlpkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg.Merge(m, src)
}
func (m *CreateSQLPackageRequest_CreateSqlpkg) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg.Size(m)
}
func (m *CreateSQLPackageRequest_CreateSqlpkg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageRequest_CreateSqlpkg proto.InternalMessageInfo

func (m *CreateSQLPackageRequest_CreateSqlpkg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSQLPackageRequest_CreateSqlpkg) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *CreateSQLPackageRequest_CreateSqlpkg) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *CreateSQLPackageRequest_CreateSqlpkg) GetAppPackageId() string {
	if m != nil {
		return m.AppPackageId
	}
	return ""
}

//
//CreateSQLPackage返回
type CreateSQLPackageResponse struct {
	//
	//SQL包下的版本列表
	VersionList []*database_delivery.SQLPackageVersion `protobuf:"bytes,1,rep,name=versionList,proto3" json:"versionList" form:"versionList"`
	//
	//SQL包ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//SQL包的数据库类型
	DbType string `protobuf:"bytes,4,opt,name=dbType,proto3" json:"dbType" form:"dbType"`
	//
	//描述
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//创建者
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//ctime
	Ctime int64 `protobuf:"varint,7,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//mtime
	Mtime int64 `protobuf:"varint,8,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	//
	//程序包Id
	AppPackageId string `protobuf:"bytes,9,opt,name=appPackageId,proto3" json:"appPackageId" form:"appPackageId"`
	//
	//文件仓库包ID
	RepoPackageId        string   `protobuf:"bytes,10,opt,name=repoPackageId,proto3" json:"repoPackageId" form:"repoPackageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSQLPackageResponse) Reset()         { *m = CreateSQLPackageResponse{} }
func (m *CreateSQLPackageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageResponse) ProtoMessage()    {}
func (*CreateSQLPackageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateSQLPackageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageResponse.Unmarshal(m, b)
}
func (m *CreateSQLPackageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageResponse.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageResponse.Merge(m, src)
}
func (m *CreateSQLPackageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageResponse.Size(m)
}
func (m *CreateSQLPackageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageResponse proto.InternalMessageInfo

func (m *CreateSQLPackageResponse) GetVersionList() []*database_delivery.SQLPackageVersion {
	if m != nil {
		return m.VersionList
	}
	return nil
}

func (m *CreateSQLPackageResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *CreateSQLPackageResponse) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *CreateSQLPackageResponse) GetAppPackageId() string {
	if m != nil {
		return m.AppPackageId
	}
	return ""
}

func (m *CreateSQLPackageResponse) GetRepoPackageId() string {
	if m != nil {
		return m.RepoPackageId
	}
	return ""
}

//
//CreateSQLPackageApi返回
type CreateSQLPackageResponseWrapper struct {
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
	Data                 *CreateSQLPackageResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateSQLPackageResponseWrapper) Reset()         { *m = CreateSQLPackageResponseWrapper{} }
func (m *CreateSQLPackageResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateSQLPackageResponseWrapper) ProtoMessage()    {}
func (*CreateSQLPackageResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{2}
}
func (m *CreateSQLPackageResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSQLPackageResponseWrapper.Unmarshal(m, b)
}
func (m *CreateSQLPackageResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSQLPackageResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateSQLPackageResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSQLPackageResponseWrapper.Merge(m, src)
}
func (m *CreateSQLPackageResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateSQLPackageResponseWrapper.Size(m)
}
func (m *CreateSQLPackageResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSQLPackageResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSQLPackageResponseWrapper proto.InternalMessageInfo

func (m *CreateSQLPackageResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateSQLPackageResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateSQLPackageResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateSQLPackageResponseWrapper) GetData() *CreateSQLPackageResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSQLPackageRequest)(nil), "sqlpkgs.CreateSQLPackageRequest")
	proto.RegisterType((*CreateSQLPackageRequest_CreateSqlpkg)(nil), "sqlpkgs.CreateSQLPackageRequest.CreateSqlpkg")
	proto.RegisterType((*CreateSQLPackageResponse)(nil), "sqlpkgs.CreateSQLPackageResponse")
	proto.RegisterType((*CreateSQLPackageResponseWrapper)(nil), "sqlpkgs.CreateSQLPackageResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6f, 0x4f, 0xd4, 0x30,
	0x18, 0xcf, 0xfd, 0x01, 0xa4, 0x77, 0x20, 0x4e, 0x23, 0x0b, 0x6f, 0x06, 0x85, 0x18, 0x4c, 0xdc,
	0x06, 0x47, 0x62, 0x44, 0x13, 0x13, 0xcf, 0x68, 0x62, 0xc2, 0x0b, 0x2d, 0x46, 0x13, 0x09, 0x5c,
	0x7a, 0x5b, 0x99, 0x93, 0xed, 0x3a, 0xda, 0x01, 0x02, 0xe1, 0xbb, 0xf9, 0xc2, 0x17, 0x7e, 0x8a,
	0x99, 0xf8, 0x09, 0xcc, 0x3e, 0x81, 0xe9, 0xd3, 0x01, 0x3d, 0x2e, 0x10, 0x7c, 0xb5, 0xf6, 0xf9,
	0xfd, 0x69, 0xfb, 0xeb, 0xb3, 0xa2, 0x76, 0x20, 0x18, 0xcd, 0x99, 0x97, 0x09, 0x9e, 0x73, 0x6b,
	0x42, 0xee, 0x27, 0xd9, 0x5e, 0x24, 0xe7, 0xdc, 0x28, 0xce, 0xbf, 0x1e, 0xf4, 0xbd, 0x80, 0xa7,
	0x7e, 0xc4, 0x23, 0xee, 0x03, 0xde, 0x3f, 0xd8, 0x85, 0x19, 0x4c, 0x60, 0xa4, 0x75, 0x73, 0x2c,
	0xe2, 0x1e, 0xa3, 0xf2, 0x98, 0x67, 0xd2, 0x4b, 0x78, 0x40, 0x13, 0x3f, 0xe0, 0x83, 0x5c, 0xd0,
	0x20, 0x97, 0x5a, 0x29, 0x58, 0xc6, 0xdd, 0x94, 0x87, 0x2c, 0x91, 0x7e, 0x45, 0xf4, 0x61, 0xea,
	0x87, 0x34, 0xa7, 0x7d, 0x2a, 0x59, 0x2f, 0x64, 0x49, 0x7c, 0xc8, 0xc4, 0xb1, 0x2f, 0xf7, 0x93,
	0x5e, 0x46, 0x83, 0x3d, 0x1a, 0xb1, 0xde, 0x21, 0x13, 0x32, 0xe6, 0x83, 0x6a, 0x99, 0xa7, 0xc6,
	0xae, 0xd2, 0xa3, 0x38, 0xdf, 0xe3, 0x47, 0x7e, 0xc4, 0x5d, 0x00, 0xdd, 0x43, 0x9a, 0xc4, 0x21,
	0xcd, 0xb9, 0x90, 0xfe, 0xc5, 0x50, 0xeb, 0xf0, 0xaf, 0x3a, 0x9a, 0x7d, 0x0d, 0xe7, 0xdc, 0xfc,
	0xb0, 0xf1, 0x5e, 0x5b, 0x13, 0xb6, 0x7f, 0xc0, 0x64, 0x6e, 0x7d, 0x3b, 0x8f, 0x60, 0x13, 0x8e,
	0x6e, 0xd7, 0xe6, 0x6b, 0xcb, 0xad, 0x8e, 0xeb, 0x55, 0x49, 0x78, 0xd7, 0xe8, 0xce, 0xeb, 0xc0,
	0xea, 0xce, 0x96, 0x85, 0x73, 0x7f, 0x97, 0x8b, 0xf4, 0x39, 0x36, 0xcd, 0x30, 0x19, 0xf2, 0x9e,
	0xfb, 0x51, 0x43, 0x6d, 0x53, 0x67, 0x2d, 0xa2, 0xe6, 0x80, 0xa6, 0x0c, 0x16, 0x9d, 0xec, 0xde,
	0x2d, 0x0b, 0xa7, 0xa5, 0x5d, 0x54, 0x15, 0x13, 0x00, 0xad, 0xc7, 0x68, 0x3c, 0xec, 0x7f, 0x3c,
	0xce, 0x98, 0x5d, 0x07, 0xda, 0xbd, 0xb2, 0x70, 0xa6, 0x34, 0x4d, 0xd7, 0x31, 0xa9, 0x08, 0xca,
	0x2f, 0x65, 0x29, 0xb7, 0x1b, 0x57, 0xfd, 0x54, 0x15, 0x13, 0x00, 0xad, 0x17, 0xa8, 0x4d, 0xb3,
	0xac, 0x3a, 0xce, 0xbb, 0xd0, 0x6e, 0x02, 0xd9, 0x38, 0x82, 0x89, 0x62, 0x32, 0x44, 0xc6, 0x3f,
	0x9b, 0xc8, 0x1e, 0x8d, 0x44, 0x66, 0x7c, 0x20, 0x99, 0xb5, 0x83, 0x5a, 0xd5, 0x85, 0x6d, 0xc4,
	0x32, 0xb7, 0x6b, 0xf3, 0x8d, 0xe5, 0x56, 0x67, 0xc9, 0x1b, 0xb9, 0x62, 0xef, 0x52, 0xfb, 0x49,
	0xf3, 0xbb, 0x0f, 0xcb, 0xc2, 0xb1, 0xf4, 0xf2, 0x86, 0x05, 0x26, 0xa6, 0xa1, 0xb5, 0x8a, 0xea,
	0x71, 0x58, 0xa5, 0xb0, 0x50, 0x16, 0xce, 0xa4, 0x16, 0xc4, 0x21, 0xfe, 0xf3, 0xdb, 0x99, 0x41,
	0xd3, 0x3b, 0x5b, 0x2b, 0xee, 0x3a, 0x75, 0x4f, 0xb6, 0x4f, 0x57, 0xd7, 0xce, 0x96, 0x48, 0x3d,
	0x0e, 0x2f, 0x12, 0x6e, 0xdc, 0x2e, 0xe1, 0xe6, 0x6d, 0x13, 0x1e, 0xbb, 0x29, 0xe1, 0x4d, 0x34,
	0x01, 0xf7, 0xce, 0x85, 0x3d, 0x0e, 0xbc, 0xf5, 0xb2, 0x70, 0xa6, 0x8d, 0xfe, 0xe0, 0x42, 0xed,
	0x78, 0x11, 0x2d, 0xec, 0x6c, 0x51, 0xf7, 0xe4, 0x95, 0xfb, 0x65, 0xc5, 0x5d, 0xdf, 0xde, 0xf2,
	0x2e, 0xc6, 0x3d, 0x77, 0xfb, 0xb4, 0xf3, 0x64, 0x6d, 0xf5, 0x6c, 0x89, 0x9c, 0x3b, 0x59, 0x8f,
	0xd0, 0x58, 0x90, 0xc7, 0x29, 0xb3, 0x27, 0xe6, 0x6b, 0xcb, 0x8d, 0xee, 0x4c, 0x59, 0x38, 0xed,
	0xca, 0x52, 0x95, 0x31, 0xd1, 0xb0, 0xe2, 0xa5, 0xc0, 0xbb, 0x73, 0x95, 0x97, 0x56, 0x3c, 0xf8,
	0x8e, 0xb4, 0xc1, 0xe4, 0x7f, 0xb4, 0x81, 0xf5, 0x12, 0x4d, 0xa9, 0x5f, 0xfa, 0x52, 0x8d, 0x40,
	0x6d, 0x97, 0x85, 0xf3, 0x40, 0xab, 0x87, 0x60, 0x4c, 0x86, 0xe9, 0xf8, 0x6f, 0x0d, 0x39, 0xd7,
	0xb5, 0xd1, 0x67, 0x41, 0xb3, 0x8c, 0x09, 0x15, 0x75, 0xc0, 0x43, 0xfd, 0x73, 0x8c, 0x99, 0x51,
	0xab, 0x2a, 0x26, 0x00, 0x5a, 0xcf, 0x50, 0x4b, 0x7d, 0xdf, 0x7c, 0xcf, 0x12, 0x1a, 0x0f, 0xaa,
	0xde, 0x30, 0x9a, 0xc9, 0x00, 0x31, 0x31, 0xa9, 0x2a, 0x27, 0x26, 0x04, 0x17, 0x55, 0x6b, 0x18,
	0x39, 0x41, 0x19, 0x13, 0x0d, 0x5b, 0x6f, 0x51, 0x53, 0x35, 0x30, 0xb4, 0x46, 0xab, 0xb3, 0x70,
	0xc3, 0xc3, 0xa0, 0xb7, 0x6f, 0xee, 0x54, 0x09, 0x31, 0x01, 0x7d, 0x7f, 0x1c, 0xde, 0xa2, 0xb5,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x43, 0x5d, 0x00, 0x72, 0x05, 0x00, 0x00,
}