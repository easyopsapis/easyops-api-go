// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package dbservice

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
//GetDBService请求
type GetDBServiceRequest struct {
	//
	//数据库服务ID
	DbServiceId          string   `protobuf:"bytes,1,opt,name=dbServiceId,proto3" json:"dbServiceId" form:"dbServiceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDBServiceRequest) Reset()         { *m = GetDBServiceRequest{} }
func (m *GetDBServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetDBServiceRequest) ProtoMessage()    {}
func (*GetDBServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetDBServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDBServiceRequest.Unmarshal(m, b)
}
func (m *GetDBServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDBServiceRequest.Marshal(b, m, deterministic)
}
func (m *GetDBServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDBServiceRequest.Merge(m, src)
}
func (m *GetDBServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetDBServiceRequest.Size(m)
}
func (m *GetDBServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDBServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDBServiceRequest proto.InternalMessageInfo

func (m *GetDBServiceRequest) GetDbServiceId() string {
	if m != nil {
		return m.DbServiceId
	}
	return ""
}

//
//GetDBService返回
type GetDBServiceResponse struct {
	//
	//数据库实例列表
	DbInstance []*database_delivery.DBInstance `protobuf:"bytes,1,rep,name=dbInstance,proto3" json:"dbInstance" form:"dbInstance"`
	//
	//运维负责人列表
	Owner []*GetDBServiceResponse_Owner `protobuf:"bytes,2,rep,name=owner,proto3" json:"owner" form:"owner"`
	//
	//数据库服务ID
	InstanceId string `protobuf:"bytes,3,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" form:"name"`
	//
	//数据库类型
	DbType string `protobuf:"bytes,5,opt,name=dbType,proto3" json:"dbType" form:"dbType"`
	//
	//描述
	Desc string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc" form:"desc"`
	//
	//变更时间
	UpdatedTime          int64    `protobuf:"varint,7,opt,name=updatedTime,proto3" json:"updatedTime" form:"updatedTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDBServiceResponse) Reset()         { *m = GetDBServiceResponse{} }
func (m *GetDBServiceResponse) String() string { return proto.CompactTextString(m) }
func (*GetDBServiceResponse) ProtoMessage()    {}
func (*GetDBServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetDBServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDBServiceResponse.Unmarshal(m, b)
}
func (m *GetDBServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDBServiceResponse.Marshal(b, m, deterministic)
}
func (m *GetDBServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDBServiceResponse.Merge(m, src)
}
func (m *GetDBServiceResponse) XXX_Size() int {
	return xxx_messageInfo_GetDBServiceResponse.Size(m)
}
func (m *GetDBServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDBServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDBServiceResponse proto.InternalMessageInfo

func (m *GetDBServiceResponse) GetDbInstance() []*database_delivery.DBInstance {
	if m != nil {
		return m.DbInstance
	}
	return nil
}

func (m *GetDBServiceResponse) GetOwner() []*GetDBServiceResponse_Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GetDBServiceResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetDBServiceResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetDBServiceResponse) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *GetDBServiceResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *GetDBServiceResponse) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type GetDBServiceResponse_Owner struct {
	//
	//用户ID
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//邮箱
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email" form:"user_email"`
	//
	//电话号码
	UserTel string `protobuf:"bytes,3,opt,name=user_tel,json=userTel,proto3" json:"user_tel" form:"user_tel"`
	//
	//用户类型
	UserType string `protobuf:"bytes,4,opt,name=user_type,json=userType,proto3" json:"user_type" form:"user_type"`
	//
	//名称
	Nickname             string   `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname" form:"nickname"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDBServiceResponse_Owner) Reset()         { *m = GetDBServiceResponse_Owner{} }
func (m *GetDBServiceResponse_Owner) String() string { return proto.CompactTextString(m) }
func (*GetDBServiceResponse_Owner) ProtoMessage()    {}
func (*GetDBServiceResponse_Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1, 0}
}
func (m *GetDBServiceResponse_Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDBServiceResponse_Owner.Unmarshal(m, b)
}
func (m *GetDBServiceResponse_Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDBServiceResponse_Owner.Marshal(b, m, deterministic)
}
func (m *GetDBServiceResponse_Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDBServiceResponse_Owner.Merge(m, src)
}
func (m *GetDBServiceResponse_Owner) XXX_Size() int {
	return xxx_messageInfo_GetDBServiceResponse_Owner.Size(m)
}
func (m *GetDBServiceResponse_Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDBServiceResponse_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_GetDBServiceResponse_Owner proto.InternalMessageInfo

func (m *GetDBServiceResponse_Owner) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetDBServiceResponse_Owner) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *GetDBServiceResponse_Owner) GetUserTel() string {
	if m != nil {
		return m.UserTel
	}
	return ""
}

func (m *GetDBServiceResponse_Owner) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *GetDBServiceResponse_Owner) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

//
//GetDBServiceApi返回
type GetDBServiceResponseWrapper struct {
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
	Data                 *GetDBServiceResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDBServiceResponseWrapper) Reset()         { *m = GetDBServiceResponseWrapper{} }
func (m *GetDBServiceResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetDBServiceResponseWrapper) ProtoMessage()    {}
func (*GetDBServiceResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetDBServiceResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDBServiceResponseWrapper.Unmarshal(m, b)
}
func (m *GetDBServiceResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDBServiceResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetDBServiceResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDBServiceResponseWrapper.Merge(m, src)
}
func (m *GetDBServiceResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetDBServiceResponseWrapper.Size(m)
}
func (m *GetDBServiceResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDBServiceResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetDBServiceResponseWrapper proto.InternalMessageInfo

func (m *GetDBServiceResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetDBServiceResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetDBServiceResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetDBServiceResponseWrapper) GetData() *GetDBServiceResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDBServiceRequest)(nil), "dbservice.GetDBServiceRequest")
	proto.RegisterType((*GetDBServiceResponse)(nil), "dbservice.GetDBServiceResponse")
	proto.RegisterType((*GetDBServiceResponse_Owner)(nil), "dbservice.GetDBServiceResponse.Owner")
	proto.RegisterType((*GetDBServiceResponseWrapper)(nil), "dbservice.GetDBServiceResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x4f, 0x13, 0x4b,
	0x18, 0x4d, 0x4b, 0x5b, 0xe8, 0xf4, 0xde, 0x0b, 0x2c, 0x5c, 0xb3, 0x81, 0x98, 0x25, 0x23, 0x9a,
	0x16, 0xdd, 0x5d, 0x7e, 0x24, 0x0a, 0x1a, 0x13, 0x6d, 0x20, 0xa6, 0x89, 0x86, 0x64, 0x6d, 0x34,
	0x91, 0x02, 0x99, 0xdd, 0xf9, 0xa8, 0x1b, 0xb6, 0x3b, 0xeb, 0xec, 0x14, 0x04, 0xc2, 0x9f, 0xe2,
	0xbf, 0x56, 0x13, 0x13, 0x5f, 0x7c, 0xec, 0x5f, 0x60, 0x66, 0x76, 0x4b, 0x07, 0x6d, 0xe2, 0x83,
	0x4f, 0xdd, 0x39, 0xe7, 0xfb, 0xce, 0xcc, 0x77, 0xce, 0x4c, 0x51, 0xb5, 0x0b, 0xc2, 0x49, 0x38,
	0x13, 0xcc, 0xa8, 0x52, 0x3f, 0x05, 0x7e, 0x16, 0x06, 0xb0, 0x64, 0x77, 0x43, 0xf1, 0xb1, 0xef,
	0x3b, 0x01, 0xeb, 0xb9, 0x5d, 0xd6, 0x65, 0xae, 0xaa, 0xf0, 0xfb, 0x27, 0x6a, 0xa5, 0x16, 0xea,
	0x2b, 0xeb, 0x5c, 0xea, 0x74, 0x99, 0x03, 0x24, 0xbd, 0x60, 0x49, 0xea, 0x44, 0x2c, 0x20, 0x91,
	0x1b, 0xb0, 0x58, 0x70, 0x12, 0x88, 0x34, 0xeb, 0xe4, 0x90, 0x30, 0xbb, 0xc7, 0x28, 0x44, 0xa9,
	0x9b, 0x17, 0xba, 0x6a, 0xe9, 0x52, 0x22, 0x88, 0x4f, 0x52, 0x38, 0xa6, 0x10, 0x85, 0x67, 0xc0,
	0x2f, 0x5c, 0xea, 0x87, 0x71, 0x2a, 0x48, 0x1c, 0x40, 0xae, 0xfe, 0x58, 0x3b, 0x4c, 0xef, 0x3c,
	0x14, 0xa7, 0xec, 0xdc, 0xed, 0x32, 0x5b, 0x91, 0xf6, 0x19, 0x89, 0x42, 0x4a, 0x04, 0xe3, 0xa9,
	0x7b, 0xf3, 0x99, 0xf5, 0xe1, 0x00, 0x2d, 0xbc, 0x02, 0xb1, 0xdb, 0x7c, 0x9b, 0x0d, 0xe5, 0xc1,
	0xa7, 0x3e, 0xa4, 0xc2, 0x78, 0x8d, 0x6a, 0xd4, 0xcf, 0xb1, 0x16, 0x35, 0x0b, 0x2b, 0x85, 0x7a,
	0xb5, 0xb9, 0x36, 0x1c, 0x58, 0xc6, 0x09, 0xe3, 0xbd, 0xa7, 0x58, 0x23, 0xf1, 0xb7, 0xaf, 0xd6,
	0x1c, 0xfa, 0xef, 0xe8, 0x60, 0xdd, 0xde, 0x21, 0xf6, 0xe5, 0xe1, 0xd5, 0xc6, 0xd6, 0xf5, 0xaa,
	0xa7, 0xb7, 0xe3, 0x2f, 0x15, 0xb4, 0x78, 0x7b, 0x97, 0x34, 0x61, 0x71, 0x0a, 0x46, 0x1b, 0x21,
	0xea, 0xb7, 0xf2, 0x49, 0xcc, 0xc2, 0xca, 0x54, 0xbd, 0xb6, 0x79, 0xd7, 0xf9, 0x6d, 0x5c, 0x67,
	0xb7, 0x39, 0x2a, 0x6a, 0xfe, 0x3f, 0x1c, 0x58, 0xf3, 0xa3, 0x43, 0x8c, 0x50, 0xec, 0x69, 0x3a,
	0xc6, 0x1b, 0x54, 0x66, 0xe7, 0x31, 0x70, 0xb3, 0xa8, 0x04, 0xef, 0x3b, 0x37, 0x99, 0x39, 0x93,
	0x4e, 0xe1, 0xec, 0xcb, 0xe2, 0xe6, 0xdc, 0x70, 0x60, 0xfd, 0x93, 0x09, 0xab, 0x6e, 0xec, 0x65,
	0x2a, 0x46, 0x0b, 0xa1, 0x91, 0xd9, 0x2d, 0x6a, 0x4e, 0x29, 0x2b, 0x1a, 0xe3, 0x53, 0x8c, 0xb9,
	0xc9, 0x4e, 0x68, 0xcd, 0xc6, 0x3d, 0x54, 0x8a, 0x49, 0x0f, 0xcc, 0x92, 0x12, 0x99, 0x1d, 0x0e,
	0xac, 0x5a, 0x26, 0x22, 0x51, 0xec, 0x29, 0xd2, 0x68, 0xa0, 0x0a, 0xf5, 0xdb, 0x17, 0x09, 0x98,
	0x65, 0x55, 0x36, 0x3f, 0x1c, 0x58, 0xff, 0x8e, 0x26, 0x96, 0x38, 0xf6, 0xf2, 0x02, 0xa9, 0x47,
	0x21, 0x0d, 0xcc, 0xca, 0xaf, 0x7a, 0x12, 0xc5, 0x9e, 0x22, 0x8d, 0x6d, 0x54, 0xeb, 0x27, 0x94,
	0x08, 0xa0, 0xed, 0xb0, 0x07, 0xe6, 0xf4, 0x4a, 0xa1, 0x3e, 0xd5, 0xbc, 0x33, 0xce, 0x52, 0x23,
	0xb1, 0xa7, 0x97, 0x2e, 0x7d, 0x2f, 0xa2, 0xf2, 0xfe, 0x04, 0x0f, 0x0a, 0x7f, 0xe3, 0x41, 0x1f,
	0xa1, 0x7e, 0x0a, 0xfc, 0x18, 0x7a, 0x24, 0x8c, 0xcc, 0xa2, 0x92, 0x7a, 0x37, 0x96, 0x1a, 0x73,
	0x52, 0xea, 0x39, 0x7a, 0x76, 0x54, 0x3f, 0x78, 0x69, 0x7f, 0x20, 0xf6, 0xe5, 0xba, 0xbd, 0x73,
	0xdc, 0xb1, 0x3b, 0xce, 0x61, 0xe3, 0x61, 0xe7, 0xc5, 0x24, 0xd0, 0x19, 0x81, 0x87, 0x57, 0x9b,
	0x8f, 0xb6, 0xaf, 0x1b, 0xab, 0x5e, 0x55, 0xaa, 0xed, 0x49, 0x31, 0xc3, 0x41, 0x33, 0x4a, 0x5a,
	0x40, 0x94, 0x67, 0xb8, 0x30, 0x1c, 0x58, 0xb3, 0xda, 0xa6, 0x02, 0x22, 0xec, 0x4d, 0xcb, 0xcf,
	0x36, 0x44, 0xc6, 0x06, 0xaa, 0x66, 0xa8, 0x0c, 0x22, 0xcb, 0x6b, 0x71, 0x38, 0xb0, 0xe6, 0xf4,
	0x06, 0x95, 0x85, 0x92, 0x55, 0x69, 0x3c, 0x41, 0x33, 0x71, 0x18, 0x9c, 0xaa, 0x84, 0xb3, 0xe8,
	0x96, 0xc7, 0x5b, 0x8c, 0x18, 0x39, 0x55, 0x09, 0x15, 0x9d, 0x35, 0xef, 0xa6, 0x18, 0xff, 0x28,
	0xa0, 0xe5, 0x49, 0x37, 0xf3, 0x3d, 0x27, 0x49, 0x02, 0x5c, 0xc6, 0x1c, 0x30, 0x0a, 0xca, 0xf7,
	0xb2, 0x1e, 0xb3, 0x44, 0xb1, 0xa7, 0x48, 0x19, 0xb3, 0xfc, 0xdd, 0xfb, 0x9c, 0x44, 0x24, 0x8c,
	0x73, 0x63, 0xb5, 0x98, 0x35, 0x12, 0x7b, 0x7a, 0xa9, 0xf1, 0x00, 0x95, 0x81, 0x73, 0xc6, 0x73,
	0x5f, 0xb4, 0x87, 0xa0, 0x60, 0xec, 0x65, 0xb4, 0xb1, 0x8b, 0x4a, 0xf2, 0x69, 0x2a, 0x37, 0x6a,
	0x9b, 0xd6, 0x1f, 0x9e, 0xd5, 0xad, 0xeb, 0x48, 0x04, 0x91, 0xd7, 0x51, 0x3e, 0xec, 0x8a, 0xfa,
	0xe3, 0xd9, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xb0, 0x28, 0x50, 0x55, 0x05, 0x00, 0x00,
}
