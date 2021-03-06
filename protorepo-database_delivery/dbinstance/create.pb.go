// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package dbinstance

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
//CreateDBInstance请求
type CreateDBInstanceRequest struct {
	//
	//数据库服务ID
	DbServiceId string `protobuf:"bytes,1,opt,name=dbServiceId,proto3" json:"dbServiceId" form:"dbServiceId"`
	//
	//创建数据库实例参数
	CreateDbinstance     *CreateDBInstanceRequest_CreateDbinstance `protobuf:"bytes,2,opt,name=createDbinstance,proto3" json:"createDbinstance" form:"createDbinstance"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *CreateDBInstanceRequest) Reset()         { *m = CreateDBInstanceRequest{} }
func (m *CreateDBInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceRequest) ProtoMessage()    {}
func (*CreateDBInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}
func (m *CreateDBInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceRequest.Unmarshal(m, b)
}
func (m *CreateDBInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceRequest.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceRequest.Merge(m, src)
}
func (m *CreateDBInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceRequest.Size(m)
}
func (m *CreateDBInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceRequest proto.InternalMessageInfo

func (m *CreateDBInstanceRequest) GetDbServiceId() string {
	if m != nil {
		return m.DbServiceId
	}
	return ""
}

func (m *CreateDBInstanceRequest) GetCreateDbinstance() *CreateDBInstanceRequest_CreateDbinstance {
	if m != nil {
		return m.CreateDbinstance
	}
	return nil
}

type CreateDBInstanceRequest_CreateDbinstance struct {
	//
	//连接客户端
	DbClient *CreateDBInstanceRequest_CreateDbinstance_DbClient `protobuf:"bytes,1,opt,name=dbClient,proto3" json:"dbClient" form:"dbClient"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//ip
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//端口
	Port int32 `protobuf:"varint,4,opt,name=port,proto3" json:"port" form:"port"`
	//
	//数据库名/服务名
	DbName string `protobuf:"bytes,5,opt,name=dbName,proto3" json:"dbName" form:"dbName"`
	//
	//连接串
	ConnURL string `protobuf:"bytes,6,opt,name=connURL,proto3" json:"connURL" form:"connURL"`
	//
	//默认用户
	UserName string `protobuf:"bytes,7,opt,name=userName,proto3" json:"userName" form:"userName"`
	//
	//默认密码
	Password string `protobuf:"bytes,8,opt,name=password,proto3" json:"password" form:"password"`
	//
	//是否有效
	IsValid              string   `protobuf:"bytes,9,opt,name=isValid,proto3" json:"isValid" form:"isValid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBInstanceRequest_CreateDbinstance) Reset() {
	*m = CreateDBInstanceRequest_CreateDbinstance{}
}
func (m *CreateDBInstanceRequest_CreateDbinstance) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceRequest_CreateDbinstance) ProtoMessage()    {}
func (*CreateDBInstanceRequest_CreateDbinstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 0}
}
func (m *CreateDBInstanceRequest_CreateDbinstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance.Unmarshal(m, b)
}
func (m *CreateDBInstanceRequest_CreateDbinstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceRequest_CreateDbinstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance.Merge(m, src)
}
func (m *CreateDBInstanceRequest_CreateDbinstance) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance.Size(m)
}
func (m *CreateDBInstanceRequest_CreateDbinstance) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance proto.InternalMessageInfo

func (m *CreateDBInstanceRequest_CreateDbinstance) GetDbClient() *CreateDBInstanceRequest_CreateDbinstance_DbClient {
	if m != nil {
		return m.DbClient
	}
	return nil
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetConnURL() string {
	if m != nil {
		return m.ConnURL
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance) GetIsValid() string {
	if m != nil {
		return m.IsValid
	}
	return ""
}

type CreateDBInstanceRequest_CreateDbinstance_DbClient struct {
	//
	//客户端IP
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//客户端ID
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) Reset() {
	*m = CreateDBInstanceRequest_CreateDbinstance_DbClient{}
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) String() string {
	return proto.CompactTextString(m)
}
func (*CreateDBInstanceRequest_CreateDbinstance_DbClient) ProtoMessage() {}
func (*CreateDBInstanceRequest_CreateDbinstance_DbClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0, 0, 0}
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient.Unmarshal(m, b)
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient.Merge(m, src)
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient.Size(m)
}
func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceRequest_CreateDbinstance_DbClient proto.InternalMessageInfo

func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateDBInstanceRequest_CreateDbinstance_DbClient) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
//CreateDBInstance返回
type CreateDBInstanceResponse struct {
	//
	//数据库服务关系
	DbService []*CreateDBInstanceResponse_DbService `protobuf:"bytes,1,rep,name=dbService,proto3" json:"dbService" form:"dbService"`
	//
	//客户端关系
	Client *CreateDBInstanceResponse_Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client" form:"client"`
	//
	//数据库实例ID,自动生成
	InstanceId string `protobuf:"bytes,3,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" form:"name"`
	//
	//数据库名/服务名
	DbName string `protobuf:"bytes,5,opt,name=dbName,proto3" json:"dbName" form:"dbName"`
	//
	//ip
	Ip string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip" form:"ip"`
	//
	//端口
	Port int32 `protobuf:"varint,7,opt,name=port,proto3" json:"port" form:"port"`
	//
	//连接串
	ConnURL string `protobuf:"bytes,8,opt,name=connURL,proto3" json:"connURL" form:"connURL"`
	//
	//默认用户
	UserName string `protobuf:"bytes,9,opt,name=userName,proto3" json:"userName" form:"userName"`
	//
	//默认密码
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password" form:"password"`
	//
	//是否有效
	IsValid              string   `protobuf:"bytes,11,opt,name=isValid,proto3" json:"isValid" form:"isValid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBInstanceResponse) Reset()         { *m = CreateDBInstanceResponse{} }
func (m *CreateDBInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceResponse) ProtoMessage()    {}
func (*CreateDBInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}
func (m *CreateDBInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceResponse.Unmarshal(m, b)
}
func (m *CreateDBInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceResponse.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceResponse.Merge(m, src)
}
func (m *CreateDBInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceResponse.Size(m)
}
func (m *CreateDBInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceResponse proto.InternalMessageInfo

func (m *CreateDBInstanceResponse) GetDbService() []*CreateDBInstanceResponse_DbService {
	if m != nil {
		return m.DbService
	}
	return nil
}

func (m *CreateDBInstanceResponse) GetClient() *CreateDBInstanceResponse_Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *CreateDBInstanceResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *CreateDBInstanceResponse) GetConnURL() string {
	if m != nil {
		return m.ConnURL
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateDBInstanceResponse) GetIsValid() string {
	if m != nil {
		return m.IsValid
	}
	return ""
}

type CreateDBInstanceResponse_DbService struct {
	//
	//数据库服务ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//数据库类型
	DbType string `protobuf:"bytes,3,opt,name=dbType,proto3" json:"dbType" form:"dbType"`
	//
	//描述
	Desc                 string   `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc" form:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBInstanceResponse_DbService) Reset()         { *m = CreateDBInstanceResponse_DbService{} }
func (m *CreateDBInstanceResponse_DbService) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceResponse_DbService) ProtoMessage()    {}
func (*CreateDBInstanceResponse_DbService) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1, 0}
}
func (m *CreateDBInstanceResponse_DbService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceResponse_DbService.Unmarshal(m, b)
}
func (m *CreateDBInstanceResponse_DbService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceResponse_DbService.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceResponse_DbService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceResponse_DbService.Merge(m, src)
}
func (m *CreateDBInstanceResponse_DbService) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceResponse_DbService.Size(m)
}
func (m *CreateDBInstanceResponse_DbService) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceResponse_DbService.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceResponse_DbService proto.InternalMessageInfo

func (m *CreateDBInstanceResponse_DbService) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateDBInstanceResponse_DbService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDBInstanceResponse_DbService) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *CreateDBInstanceResponse_DbService) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type CreateDBInstanceResponse_Client struct {
	//
	//客户端ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//客户端IP
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip" form:"ip"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBInstanceResponse_Client) Reset()         { *m = CreateDBInstanceResponse_Client{} }
func (m *CreateDBInstanceResponse_Client) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceResponse_Client) ProtoMessage()    {}
func (*CreateDBInstanceResponse_Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1, 1}
}
func (m *CreateDBInstanceResponse_Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceResponse_Client.Unmarshal(m, b)
}
func (m *CreateDBInstanceResponse_Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceResponse_Client.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceResponse_Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceResponse_Client.Merge(m, src)
}
func (m *CreateDBInstanceResponse_Client) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceResponse_Client.Size(m)
}
func (m *CreateDBInstanceResponse_Client) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceResponse_Client.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceResponse_Client proto.InternalMessageInfo

func (m *CreateDBInstanceResponse_Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateDBInstanceResponse_Client) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

//
//CreateDBInstanceApi返回
type CreateDBInstanceResponseWrapper struct {
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
	Data                 *CreateDBInstanceResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateDBInstanceResponseWrapper) Reset()         { *m = CreateDBInstanceResponseWrapper{} }
func (m *CreateDBInstanceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*CreateDBInstanceResponseWrapper) ProtoMessage()    {}
func (*CreateDBInstanceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{2}
}
func (m *CreateDBInstanceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBInstanceResponseWrapper.Unmarshal(m, b)
}
func (m *CreateDBInstanceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBInstanceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *CreateDBInstanceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBInstanceResponseWrapper.Merge(m, src)
}
func (m *CreateDBInstanceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_CreateDBInstanceResponseWrapper.Size(m)
}
func (m *CreateDBInstanceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBInstanceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBInstanceResponseWrapper proto.InternalMessageInfo

func (m *CreateDBInstanceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateDBInstanceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *CreateDBInstanceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateDBInstanceResponseWrapper) GetData() *CreateDBInstanceResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDBInstanceRequest)(nil), "dbinstance.CreateDBInstanceRequest")
	proto.RegisterType((*CreateDBInstanceRequest_CreateDbinstance)(nil), "dbinstance.CreateDBInstanceRequest.CreateDbinstance")
	proto.RegisterType((*CreateDBInstanceRequest_CreateDbinstance_DbClient)(nil), "dbinstance.CreateDBInstanceRequest.CreateDbinstance.DbClient")
	proto.RegisterType((*CreateDBInstanceResponse)(nil), "dbinstance.CreateDBInstanceResponse")
	proto.RegisterType((*CreateDBInstanceResponse_DbService)(nil), "dbinstance.CreateDBInstanceResponse.DbService")
	proto.RegisterType((*CreateDBInstanceResponse_Client)(nil), "dbinstance.CreateDBInstanceResponse.Client")
	proto.RegisterType((*CreateDBInstanceResponseWrapper)(nil), "dbinstance.CreateDBInstanceResponseWrapper")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78) }

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 1074 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x59, 0x5f, 0x6f, 0x9b, 0xd6,
	0x1b, 0x16, 0xae, 0xe3, 0x98, 0xe3, 0xfe, 0x1a, 0x97, 0xdf, 0xb4, 0x30, 0xef, 0x82, 0x84, 0xa6,
	0x15, 0xa4, 0x06, 0xdb, 0xf8, 0x4f, 0x1b, 0x4b, 0x53, 0x56, 0x37, 0x9b, 0x64, 0xa9, 0xea, 0x05,
	0xeb, 0x3a, 0x69, 0xb6, 0x23, 0x61, 0x43, 0x32, 0xb4, 0xc4, 0x30, 0x20, 0xcd, 0x56, 0x1b, 0xa9,
	0x57, 0xd3, 0xee, 0x26, 0xed, 0x7e, 0x77, 0xfb, 0x0a, 0xfb, 0x14, 0xfb, 0x0e, 0x4c, 0xda, 0x47,
	0xa0, 0xda, 0xf5, 0xa6, 0x73, 0x0e, 0x60, 0x6a, 0x87, 0x24, 0x9e, 0xd2, 0x48, 0x91, 0xc8, 0x4d,
	0xe0, 0x7d, 0x9f, 0xf7, 0x3d, 0xcf, 0x39, 0x96, 0xde, 0xe7, 0x3c, 0x02, 0xdc, 0x1e, 0x59, 0x9a,
	0xe2, 0x68, 0xa2, 0x69, 0x19, 0x8e, 0x41, 0x01, 0x75, 0xa8, 0x8f, 0x6d, 0x47, 0x19, 0x8f, 0xb4,
	0x92, 0x70, 0xa8, 0x3b, 0xdf, 0x9c, 0x0c, 0xc5, 0x91, 0x71, 0x5c, 0x39, 0x34, 0x0e, 0x8d, 0x0a,
	0x82, 0x0c, 0x4f, 0x0e, 0xd0, 0x1b, 0x7a, 0x41, 0x4f, 0xb8, 0xb4, 0xd4, 0x8a, 0xc1, 0x8f, 0x4f,
	0x75, 0xe7, 0x5b, 0xe3, 0xb4, 0x72, 0x68, 0x08, 0x28, 0x29, 0xbc, 0x52, 0x8e, 0x74, 0x55, 0x71,
	0x0c, 0xcb, 0xae, 0x44, 0x8f, 0xb8, 0x8e, 0xfd, 0xf1, 0x23, 0xb0, 0xfe, 0x14, 0x71, 0xd8, 0xeb,
	0x74, 0x83, 0xb5, 0x65, 0xed, 0xbb, 0x13, 0xcd, 0x76, 0xa8, 0x67, 0xa0, 0xa0, 0x0e, 0xbf, 0xd0,
	0xac, 0x57, 0xfa, 0x48, 0xeb, 0xaa, 0x34, 0xb1, 0x41, 0x70, 0x64, 0x67, 0xdb, 0xf7, 0x18, 0xea,
	0xc0, 0xb0, 0x8e, 0xdb, 0x6c, 0x2c, 0xc9, 0xfe, 0xf5, 0x27, 0x53, 0x04, 0x77, 0xf6, 0x7b, 0x55,
	0x61, 0x47, 0x11, 0x5e, 0x0f, 0x26, 0xb5, 0xba, 0xbb, 0x25, 0xc7, 0xcb, 0x29, 0x17, 0x14, 0xf1,
	0x66, 0xf7, 0xa2, 0x4d, 0xd2, 0x99, 0x0d, 0x82, 0x2b, 0x48, 0x0d, 0x71, 0xb6, 0x6f, 0x31, 0x81,
	0x4c, 0x18, 0x8f, 0x80, 0x9d, 0x8f, 0x7d, 0x8f, 0x59, 0xc7, 0x44, 0xe6, 0xfb, 0xb2, 0xf2, 0xc2,
	0x52, 0xa5, 0x5f, 0xd7, 0x41, 0x71, 0xbe, 0x07, 0x35, 0x06, 0x79, 0x75, 0xf8, 0xf4, 0x48, 0xd7,
	0xc6, 0x0e, 0xda, 0x5e, 0x41, 0xfa, 0xe4, 0xbf, 0x70, 0x11, 0xf7, 0x82, 0x26, 0x9d, 0xff, 0xfb,
	0x1e, 0xb3, 0x16, 0x9e, 0x0e, 0x8e, 0xb1, 0x72, 0xb4, 0x06, 0x75, 0x0f, 0x64, 0xc7, 0xca, 0x31,
	0xde, 0x37, 0xd9, 0x59, 0xf3, 0x3d, 0xa6, 0x80, 0xc1, 0x30, 0xca, 0xca, 0x28, 0x49, 0xfd, 0x4e,
	0x82, 0x8c, 0x6e, 0xd2, 0xb7, 0x10, 0xe6, 0x37, 0xd2, 0xf7, 0x18, 0x12, 0x83, 0x74, 0x13, 0x1e,
	0xf3, 0x2f, 0x24, 0xf8, 0x99, 0xdc, 0xe7, 0x38, 0x89, 0x6b, 0xf6, 0xaa, 0x42, 0x73, 0x30, 0xa9,
	0xb9, 0xd3, 0x5e, 0x55, 0x68, 0x0c, 0xfa, 0xea, 0xa4, 0xe6, 0xf2, 0xf0, 0xb9, 0x36, 0xd8, 0x85,
	0x2f, 0x65, 0xc9, 0xe5, 0xb9, 0xbe, 0x78, 0x49, 0x24, 0x3f, 0xa9, 0xbb, 0xfc, 0xb4, 0x6f, 0x6f,
	0x73, 0x1c, 0x07, 0x7f, 0xc6, 0x27, 0xc2, 0xe7, 0x8a, 0x70, 0x30, 0x98, 0xd4, 0xca, 0x0d, 0xb7,
	0xcd, 0x4f, 0x1e, 0xb9, 0x0b, 0xd1, 0x69, 0x9b, 0xe7, 0xa7, 0x67, 0x82, 0x5b, 0x2e, 0xd7, 0x5e,
	0x40, 0x73, 0x9c, 0x84, 0x79, 0x4c, 0xa5, 0x80, 0xc5, 0xb4, 0xd6, 0x57, 0xfb, 0xea, 0xb4, 0x57,
	0x13, 0x76, 0x20, 0x0f, 0x4c, 0xf6, 0x02, 0x0c, 0xa6, 0x99, 0xb8, 0x72, 0xd3, 0xe5, 0xb8, 0xc5,
	0xb5, 0x79, 0xbc, 0xc5, 0x69, 0xfb, 0x5a, 0x38, 0x34, 0x12, 0x39, 0xc0, 0xb2, 0xb3, 0x52, 0xbb,
	0x57, 0x49, 0xec, 0x1c, 0x66, 0xf5, 0x44, 0x66, 0x8d, 0x04, 0x66, 0x93, 0x6a, 0x59, 0x72, 0xaf,
	0x89, 0x9d, 0x94, 0xc8, 0xae, 0x99, 0xcc, 0xae, 0x7e, 0x5d, 0xec, 0x6a, 0x89, 0xec, 0x5a, 0xc9,
	0xec, 0x1a, 0xef, 0x83, 0x5d, 0x3b, 0x89, 0xc8, 0xa3, 0x64, 0x22, 0xcd, 0xab, 0x27, 0xc2, 0x73,
	0xf7, 0xc5, 0x87, 0xfc, 0x6e, 0xdf, 0xde, 0xde, 0x92, 0x33, 0xba, 0x49, 0xb5, 0x40, 0xd6, 0x34,
	0x2c, 0x87, 0xce, 0x6e, 0x10, 0xdc, 0x4a, 0x87, 0x9d, 0x0d, 0x37, 0x18, 0x85, 0x93, 0x6b, 0xad,
	0xf8, 0x4f, 0xf8, 0x47, 0xd0, 0x6f, 0xde, 0x64, 0x65, 0x84, 0xa7, 0x78, 0x90, 0x53, 0x87, 0xcf,
	0xe1, 0x58, 0x5c, 0x41, 0x23, 0xef, 0xae, 0xef, 0x31, 0xff, 0x0b, 0x67, 0xe8, 0x73, 0x34, 0x18,
	0x03, 0x00, 0x55, 0x06, 0xab, 0x23, 0x63, 0x3c, 0xfe, 0x52, 0x7e, 0x46, 0xe7, 0x10, 0x96, 0xf2,
	0x3d, 0xe6, 0x4e, 0x20, 0x02, 0x38, 0xc1, 0xca, 0x21, 0x84, 0xaa, 0x80, 0xfc, 0x89, 0xad, 0x59,
	0xa8, 0xf5, 0x2a, 0x82, 0xc7, 0xc6, 0x73, 0x98, 0x61, 0xe5, 0x08, 0x44, 0xbd, 0x00, 0x79, 0x53,
	0xb1, 0xed, 0x53, 0xc3, 0x52, 0xe9, 0x3c, 0x2a, 0x78, 0x3c, 0x2b, 0x08, 0x33, 0x70, 0x27, 0x9b,
	0x80, 0xd9, 0xe7, 0x7a, 0x8a, 0xf0, 0xfa, 0x89, 0xf0, 0x75, 0x55, 0xd8, 0xd9, 0xfc, 0xf4, 0xde,
	0xd6, 0xfd, 0xfd, 0x07, 0xe2, 0x6e, 0x65, 0xc0, 0x4f, 0x5a, 0xe5, 0x66, 0xd5, 0xdd, 0x92, 0xa3,
	0x4e, 0x90, 0xb4, 0x6e, 0xbf, 0x84, 0xba, 0x4b, 0x93, 0xf3, 0xa4, 0x83, 0x04, 0x2b, 0x87, 0x90,
	0xd2, 0xdf, 0x24, 0xc8, 0x87, 0x72, 0x12, 0x4a, 0x01, 0x91, 0x4a, 0x41, 0x2a, 0x05, 0xa9, 0x14,
	0xa4, 0x52, 0x70, 0x53, 0xa4, 0xa0, 0x06, 0x32, 0xba, 0x1a, 0xdc, 0x72, 0x37, 0x63, 0x53, 0x2b,
	0xc1, 0x27, 0x64, 0x74, 0x95, 0xfd, 0xa9, 0x04, 0xe8, 0xc5, 0xfb, 0xb6, 0x6d, 0x1a, 0x63, 0x5b,
	0xa3, 0x86, 0x80, 0x8c, 0xac, 0x04, 0x4d, 0x6c, 0xdc, 0xe2, 0x0a, 0x92, 0x78, 0xfe, 0x45, 0x1d,
	0x17, 0x8a, 0x7b, 0x61, 0x55, 0xe7, 0x03, 0xdf, 0x63, 0x8a, 0x73, 0xbe, 0x85, 0x95, 0x67, 0x6d,
	0xa9, 0x97, 0x20, 0x37, 0xc2, 0x4e, 0x00, 0xbb, 0x92, 0x87, 0x97, 0x5a, 0x20, 0xb8, 0xf7, 0xc7,
	0x34, 0x6b, 0x14, 0xdc, 0xfa, 0x83, 0x6e, 0x54, 0x17, 0x80, 0xb0, 0x4d, 0x57, 0x0d, 0x6e, 0xf5,
	0xbc, 0xef, 0x31, 0x77, 0x83, 0x33, 0x89, 0x72, 0x67, 0x9f, 0x4d, 0xac, 0x38, 0xb2, 0x0f, 0xd9,
	0xf3, 0xec, 0xc3, 0x12, 0x72, 0x1a, 0xc8, 0x4b, 0x2e, 0x95, 0x97, 0x54, 0x5e, 0x52, 0x79, 0x49,
	0xe5, 0xe5, 0xa6, 0x39, 0x8d, 0xd5, 0x25, 0x9d, 0x46, 0xcc, 0x3e, 0xe4, 0x97, 0xb3, 0x0f, 0xe4,
	0xb2, 0xf6, 0x01, 0xbc, 0x0f, 0xfb, 0x50, 0xb8, 0xd8, 0x3e, 0xfc, 0x41, 0x00, 0x32, 0x12, 0xbd,
	0x39, 0xed, 0x21, 0xae, 0x42, 0x7b, 0x32, 0x17, 0x6a, 0xcf, 0x8b, 0x1f, 0x4c, 0x2d, 0xd0, 0xb9,
	0x77, 0xb4, 0x07, 0xc6, 0x91, 0xf6, 0xc0, 0x07, 0xd8, 0x4f, 0xd5, 0xec, 0xd1, 0xa2, 0x96, 0xc1,
	0x28, 0x2b, 0xa3, 0x64, 0xe9, 0x2d, 0x09, 0x72, 0x81, 0x15, 0xc2, 0x57, 0x0a, 0x62, 0x89, 0x2b,
	0x45, 0x28, 0x6f, 0x99, 0x54, 0xde, 0x52, 0x79, 0x4b, 0xe5, 0x2d, 0x95, 0xb7, 0x1b, 0x22, 0x6f,
	0xec, 0x5b, 0x02, 0x30, 0x49, 0x86, 0xe3, 0x2b, 0x4b, 0x31, 0x4d, 0xcd, 0x82, 0xe3, 0x73, 0x64,
	0xa8, 0x1a, 0x1a, 0x88, 0x2b, 0xf1, 0xf1, 0x09, 0xa3, 0xac, 0x8c, 0x92, 0xd4, 0x63, 0x50, 0x80,
	0xff, 0x3f, 0xfb, 0xde, 0x3c, 0x52, 0xf4, 0x71, 0x30, 0x08, 0x3f, 0x9c, 0x7d, 0xc0, 0x89, 0x25,
	0x59, 0x39, 0x0e, 0xa5, 0x1e, 0x80, 0x15, 0xcd, 0xb2, 0x0c, 0x2b, 0x98, 0xe3, 0x45, 0xdf, 0x63,
	0x6e, 0xe3, 0x1a, 0x14, 0x66, 0x65, 0x9c, 0xa6, 0xba, 0x20, 0xab, 0x2a, 0x8e, 0x82, 0xa6, 0x78,
	0x41, 0xda, 0xba, 0x8c, 0x65, 0x7a, 0x67, 0xd6, 0x2b, 0x8e, 0x02, 0x67, 0xbd, 0xe2, 0x28, 0xc3,
	0x1c, 0xfa, 0x20, 0x55, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb8, 0xbc, 0x4c, 0x13, 0x1b,
	0x00, 0x00,
}
