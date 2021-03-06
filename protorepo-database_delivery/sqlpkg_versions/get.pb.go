// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package sqlpkg_versions

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
//GetSQLPackageVersion请求
type GetSQLPackageVersionRequest struct {
	//
	//SQL包ID
	PkgId string `protobuf:"bytes,1,opt,name=pkgId,proto3" json:"pkgId" form:"pkgId"`
	//
	//SQL包版本ID
	VersionId            string   `protobuf:"bytes,2,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSQLPackageVersionRequest) Reset()         { *m = GetSQLPackageVersionRequest{} }
func (m *GetSQLPackageVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSQLPackageVersionRequest) ProtoMessage()    {}
func (*GetSQLPackageVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetSQLPackageVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSQLPackageVersionRequest.Unmarshal(m, b)
}
func (m *GetSQLPackageVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSQLPackageVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetSQLPackageVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSQLPackageVersionRequest.Merge(m, src)
}
func (m *GetSQLPackageVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSQLPackageVersionRequest.Size(m)
}
func (m *GetSQLPackageVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSQLPackageVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSQLPackageVersionRequest proto.InternalMessageInfo

func (m *GetSQLPackageVersionRequest) GetPkgId() string {
	if m != nil {
		return m.PkgId
	}
	return ""
}

func (m *GetSQLPackageVersionRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//GetSQLPackageVersion返回
type GetSQLPackageVersionResponse struct {
	//
	//所属的SQL包
	SqlPackage *GetSQLPackageVersionResponse_SqlPackage `protobuf:"bytes,1,opt,name=sqlPackage,proto3" json:"sqlPackage" form:"sqlPackage"`
	//
	//版本ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//仓库版本ID
	RepoVersionId string `protobuf:"bytes,4,opt,name=repoVersionId,proto3" json:"repoVersionId" form:"repoVersionId"`
	//
	//程序包版本ID
	AppVersionId string `protobuf:"bytes,5,opt,name=appVersionId,proto3" json:"appVersionId" form:"appVersionId"`
	//
	//版本名称
	VersionName string `protobuf:"bytes,6,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//备注
	Memo string `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//创建者
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//创建时间
	Ctime int64 `protobuf:"varint,9,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//修改时间
	Mtime                int64    `protobuf:"varint,10,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSQLPackageVersionResponse) Reset()         { *m = GetSQLPackageVersionResponse{} }
func (m *GetSQLPackageVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetSQLPackageVersionResponse) ProtoMessage()    {}
func (*GetSQLPackageVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetSQLPackageVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSQLPackageVersionResponse.Unmarshal(m, b)
}
func (m *GetSQLPackageVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSQLPackageVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetSQLPackageVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSQLPackageVersionResponse.Merge(m, src)
}
func (m *GetSQLPackageVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetSQLPackageVersionResponse.Size(m)
}
func (m *GetSQLPackageVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSQLPackageVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSQLPackageVersionResponse proto.InternalMessageInfo

func (m *GetSQLPackageVersionResponse) GetSqlPackage() *GetSQLPackageVersionResponse_SqlPackage {
	if m != nil {
		return m.SqlPackage
	}
	return nil
}

func (m *GetSQLPackageVersionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetRepoVersionId() string {
	if m != nil {
		return m.RepoVersionId
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetAppVersionId() string {
	if m != nil {
		return m.AppVersionId
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetSQLPackageVersionResponse) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *GetSQLPackageVersionResponse) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

type GetSQLPackageVersionResponse_SqlPackage struct {
	//
	//SQL包ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//描述
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//文件仓库包ID
	RepoPackageId string `protobuf:"bytes,4,opt,name=repoPackageId,proto3" json:"repoPackageId" form:"repoPackageId"`
	//
	//创建者
	Creator string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//mtime
	Mtime                int64    `protobuf:"varint,6,opt,name=mtime,proto3" json:"mtime" form:"mtime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSQLPackageVersionResponse_SqlPackage) Reset() {
	*m = GetSQLPackageVersionResponse_SqlPackage{}
}
func (m *GetSQLPackageVersionResponse_SqlPackage) String() string { return proto.CompactTextString(m) }
func (*GetSQLPackageVersionResponse_SqlPackage) ProtoMessage()    {}
func (*GetSQLPackageVersionResponse_SqlPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1, 0}
}
func (m *GetSQLPackageVersionResponse_SqlPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage.Unmarshal(m, b)
}
func (m *GetSQLPackageVersionResponse_SqlPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage.Marshal(b, m, deterministic)
}
func (m *GetSQLPackageVersionResponse_SqlPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage.Merge(m, src)
}
func (m *GetSQLPackageVersionResponse_SqlPackage) XXX_Size() int {
	return xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage.Size(m)
}
func (m *GetSQLPackageVersionResponse_SqlPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage.DiscardUnknown(m)
}

var xxx_messageInfo_GetSQLPackageVersionResponse_SqlPackage proto.InternalMessageInfo

func (m *GetSQLPackageVersionResponse_SqlPackage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetSQLPackageVersionResponse_SqlPackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSQLPackageVersionResponse_SqlPackage) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *GetSQLPackageVersionResponse_SqlPackage) GetRepoPackageId() string {
	if m != nil {
		return m.RepoPackageId
	}
	return ""
}

func (m *GetSQLPackageVersionResponse_SqlPackage) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GetSQLPackageVersionResponse_SqlPackage) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

//
//GetSQLPackageVersionApi返回
type GetSQLPackageVersionResponseWrapper struct {
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
	Data                 *GetSQLPackageVersionResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetSQLPackageVersionResponseWrapper) Reset()         { *m = GetSQLPackageVersionResponseWrapper{} }
func (m *GetSQLPackageVersionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetSQLPackageVersionResponseWrapper) ProtoMessage()    {}
func (*GetSQLPackageVersionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetSQLPackageVersionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSQLPackageVersionResponseWrapper.Unmarshal(m, b)
}
func (m *GetSQLPackageVersionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSQLPackageVersionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetSQLPackageVersionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSQLPackageVersionResponseWrapper.Merge(m, src)
}
func (m *GetSQLPackageVersionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetSQLPackageVersionResponseWrapper.Size(m)
}
func (m *GetSQLPackageVersionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSQLPackageVersionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetSQLPackageVersionResponseWrapper proto.InternalMessageInfo

func (m *GetSQLPackageVersionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetSQLPackageVersionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetSQLPackageVersionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetSQLPackageVersionResponseWrapper) GetData() *GetSQLPackageVersionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSQLPackageVersionRequest)(nil), "sqlpkg_versions.GetSQLPackageVersionRequest")
	proto.RegisterType((*GetSQLPackageVersionResponse)(nil), "sqlpkg_versions.GetSQLPackageVersionResponse")
	proto.RegisterType((*GetSQLPackageVersionResponse_SqlPackage)(nil), "sqlpkg_versions.GetSQLPackageVersionResponse.SqlPackage")
	proto.RegisterType((*GetSQLPackageVersionResponseWrapper)(nil), "sqlpkg_versions.GetSQLPackageVersionResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x7f, 0x6b, 0xd3, 0x40,
	0x18, 0xa6, 0xd9, 0xba, 0xd9, 0xeb, 0xdc, 0xe6, 0xf9, 0x2b, 0x4c, 0x21, 0xdb, 0x6d, 0xc8, 0xfe,
	0x30, 0xe9, 0xb6, 0x82, 0xac, 0x13, 0x04, 0x0b, 0x2a, 0x03, 0x19, 0x7a, 0x85, 0x09, 0x96, 0x6e,
	0x5c, 0x9b, 0x5b, 0x0c, 0x6d, 0x7a, 0x69, 0x92, 0x6e, 0xb2, 0x31, 0xf0, 0xab, 0xf8, 0x11, 0xfc,
	0x12, 0x7e, 0x8c, 0x08, 0x7e, 0x84, 0x7c, 0x02, 0xb9, 0xf7, 0xd2, 0x34, 0x2d, 0x5b, 0x41, 0xf1,
	0xaf, 0xdc, 0xbd, 0xcf, 0xf3, 0xbc, 0x79, 0xee, 0xb9, 0x37, 0x41, 0x25, 0x87, 0x47, 0x96, 0x1f,
	0x88, 0x48, 0xe0, 0x95, 0x70, 0xd0, 0xf3, 0xbb, 0xce, 0xe9, 0x39, 0x0f, 0x42, 0x57, 0xf4, 0xc3,
	0x35, 0xd3, 0x71, 0xa3, 0x2f, 0xc3, 0xb6, 0xd5, 0x11, 0x5e, 0xc5, 0x11, 0x8e, 0xa8, 0x00, 0xaf,
	0x3d, 0x3c, 0x83, 0x1d, 0x6c, 0x60, 0xa5, 0xf4, 0x6b, 0x2f, 0x72, 0x74, 0xef, 0xc2, 0x8d, 0xba,
	0xe2, 0xa2, 0xe2, 0x08, 0x13, 0x40, 0xf3, 0x9c, 0xf5, 0x5c, 0x9b, 0x45, 0x22, 0x08, 0x2b, 0xd9,
	0x52, 0xe9, 0xc8, 0xf7, 0x02, 0x7a, 0xf2, 0x8e, 0x47, 0x8d, 0x8f, 0xef, 0x3f, 0xb0, 0x4e, 0x97,
	0x39, 0xfc, 0x58, 0x19, 0xa0, 0x7c, 0x30, 0xe4, 0x61, 0x84, 0x0f, 0x50, 0xd1, 0xef, 0x3a, 0x87,
	0xb6, 0x5e, 0x58, 0x2f, 0x6c, 0x97, 0xea, 0x5b, 0x49, 0x6c, 0x2c, 0x9d, 0x89, 0xc0, 0x3b, 0x20,
	0x50, 0x26, 0xbf, 0x7f, 0x19, 0xab, 0x68, 0xf9, 0xa4, 0xb9, 0x63, 0xd6, 0x98, 0x79, 0xd9, 0xba,
	0xda, 0xad, 0x5e, 0x6f, 0x51, 0x25, 0xc1, 0x6f, 0x51, 0x29, 0x3d, 0xce, 0xa1, 0xad, 0x6b, 0xa0,
	0xdf, 0x4e, 0x62, 0x63, 0x55, 0xe9, 0x33, 0xe8, 0xe6, 0x1e, 0x63, 0x29, 0xf9, 0xb1, 0x88, 0x9e,
	0xde, 0xec, 0x31, 0xf4, 0x45, 0x3f, 0xe4, 0x58, 0x20, 0x14, 0x0e, 0x7a, 0x29, 0x08, 0x4e, 0xcb,
	0x7b, 0xfb, 0xd6, 0x54, 0xa2, 0xd6, 0xac, 0x16, 0x56, 0x23, 0xd3, 0xd7, 0x1f, 0x26, 0xb1, 0x71,
	0x4f, 0x79, 0x1c, 0x77, 0x25, 0x34, 0xf7, 0x0a, 0xbc, 0x8b, 0x34, 0x77, 0x74, 0xa4, 0x8d, 0x24,
	0x36, 0x4a, 0x8a, 0xee, 0xde, 0x72, 0x16, 0xcd, 0xb5, 0xf1, 0x26, 0x9a, 0xef, 0x33, 0x8f, 0xeb,
	0x73, 0x20, 0x5a, 0x49, 0x62, 0xa3, 0xac, 0x44, 0xb2, 0x4a, 0x28, 0x80, 0xf8, 0x15, 0xba, 0x1b,
	0x70, 0x5f, 0x1c, 0x67, 0xa9, 0xcd, 0x03, 0x5b, 0x4f, 0x62, 0xe3, 0x81, 0x62, 0x4f, 0xc0, 0x84,
	0x4e, 0xd2, 0xf1, 0x4b, 0xb4, 0xc4, 0x7c, 0x7f, 0x2c, 0x2f, 0x82, 0xfc, 0x71, 0x12, 0x1b, 0xf7,
	0x95, 0x3c, 0x8f, 0x12, 0x3a, 0x41, 0xc6, 0xfb, 0xa8, 0x9c, 0x66, 0x75, 0x24, 0x8d, 0x2e, 0x80,
	0xf6, 0x51, 0x12, 0x1b, 0x78, 0xe2, 0xc2, 0x8e, 0xc0, 0x6f, 0x9e, 0x2a, 0xcf, 0xe6, 0x71, 0x4f,
	0xe8, 0x8b, 0xd3, 0x67, 0x93, 0x55, 0x42, 0x01, 0xc4, 0x0d, 0xb4, 0xd8, 0x09, 0xb8, 0x1c, 0x3d,
	0xfd, 0x0e, 0xf0, 0x6a, 0x49, 0x6c, 0x2c, 0x2b, 0x5e, 0x0a, 0xc8, 0xf4, 0x36, 0xd1, 0xc6, 0x49,
	0x93, 0x99, 0x97, 0xaf, 0xcd, 0xcf, 0x3b, 0x66, 0xad, 0xd5, 0xb4, 0xb2, 0xf5, 0xa9, 0xd9, 0xba,
	0xda, 0x7b, 0x5e, 0xdd, 0xbd, 0xde, 0xa2, 0xa3, 0x4e, 0xf8, 0x19, 0x2a, 0x76, 0x22, 0xd7, 0xe3,
	0x7a, 0x69, 0xbd, 0xb0, 0x3d, 0x57, 0x5f, 0x1d, 0x8f, 0x27, 0x94, 0x09, 0x55, 0xb0, 0xe4, 0x79,
	0xc0, 0x43, 0xd3, 0x3c, 0x2f, 0xe5, 0xc1, 0x73, 0xed, 0xa7, 0x86, 0x50, 0x63, 0xfa, 0x9e, 0x0b,
	0xff, 0x72, 0xcf, 0xda, 0xac, 0x7b, 0x1e, 0x05, 0x36, 0x37, 0x2b, 0xb0, 0x74, 0x18, 0x52, 0x2f,
	0xb7, 0x0d, 0x43, 0x06, 0xa7, 0xc3, 0x90, 0xed, 0xf3, 0x81, 0x17, 0xff, 0x67, 0xe0, 0x2a, 0xc8,
	0x85, 0x99, 0x41, 0x92, 0x6f, 0x1a, 0xda, 0x9c, 0xf5, 0xc1, 0x7d, 0x0a, 0x98, 0xef, 0xf3, 0x40,
	0x26, 0xd1, 0x11, 0xb6, 0xfa, 0x68, 0x8b, 0xf9, 0x24, 0x64, 0x95, 0x50, 0x00, 0xe5, 0x64, 0xca,
	0xe7, 0x9b, 0xaf, 0x7e, 0x8f, 0xb9, 0xfd, 0x34, 0xda, 0xdc, 0x64, 0xe6, 0x40, 0x42, 0xf3, 0x54,
	0x69, 0x97, 0x07, 0x81, 0x08, 0xd2, 0xa4, 0x73, 0x76, 0xa1, 0x4c, 0xa8, 0x82, 0x31, 0x45, 0xf3,
	0x36, 0x8b, 0x18, 0x44, 0x5c, 0xde, 0x33, 0xff, 0xea, 0xdf, 0x91, 0x77, 0x2d, 0x9b, 0x10, 0x0a,
	0xbd, 0xda, 0x0b, 0xf0, 0x87, 0xad, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xed, 0xb2, 0x4f,
	0xe6, 0x05, 0x00, 0x00,
}
