// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_by_post.proto

package instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	easy_flow "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/easy_flow"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//GetSearch请求
type GetSearchRequest struct {
	//
	//多个包ID，以`,`分隔
	PackageIds string `protobuf:"bytes,1,opt,name=packageIds,proto3" json:"packageIds" form:"packageIds"`
	//
	//多个设备ID，以`,`分隔
	DeviceIds string `protobuf:"bytes,2,opt,name=deviceIds,proto3" json:"deviceIds" form:"deviceIds"`
	//
	//设备IP
	DeviceIp string `protobuf:"bytes,3,opt,name=deviceIp,proto3" json:"deviceIp" form:"deviceIp"`
	//
	//多个应用ID
	AppIds string `protobuf:"bytes,4,opt,name=appIds,proto3" json:"appIds" form:"appIds"`
	//
	//分页
	Page int32 `protobuf:"varint,5,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize int32 `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	//
	//排序字段
	Order string `protobuf:"bytes,7,opt,name=order,proto3" json:"order" form:"order"`
	//
	//包Id
	PackageId string `protobuf:"bytes,8,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId string `protobuf:"bytes,9,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	//
	//设备Id
	DeviceId string `protobuf:"bytes,10,opt,name=deviceId,proto3" json:"deviceId" form:"deviceId"`
	//
	//应用Id
	AppId                string   `protobuf:"bytes,11,opt,name=appId,proto3" json:"appId" form:"appId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSearchRequest) Reset()         { *m = GetSearchRequest{} }
func (m *GetSearchRequest) String() string { return proto.CompactTextString(m) }
func (*GetSearchRequest) ProtoMessage()    {}
func (*GetSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_301727bc325f7fd5, []int{0}
}
func (m *GetSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSearchRequest.Unmarshal(m, b)
}
func (m *GetSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSearchRequest.Marshal(b, m, deterministic)
}
func (m *GetSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSearchRequest.Merge(m, src)
}
func (m *GetSearchRequest) XXX_Size() int {
	return xxx_messageInfo_GetSearchRequest.Size(m)
}
func (m *GetSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSearchRequest proto.InternalMessageInfo

func (m *GetSearchRequest) GetPackageIds() string {
	if m != nil {
		return m.PackageIds
	}
	return ""
}

func (m *GetSearchRequest) GetDeviceIds() string {
	if m != nil {
		return m.DeviceIds
	}
	return ""
}

func (m *GetSearchRequest) GetDeviceIp() string {
	if m != nil {
		return m.DeviceIp
	}
	return ""
}

func (m *GetSearchRequest) GetAppIds() string {
	if m != nil {
		return m.AppIds
	}
	return ""
}

func (m *GetSearchRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetSearchRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetSearchRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetSearchRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetSearchRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *GetSearchRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *GetSearchRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

//
//GetSearch返回
type GetSearchResponse struct {
	//
	//总页数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total"`
	//
	//实例信息列表
	List []*easy_flow.InstanceInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list" form:"list"`
	//
	//分页
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page" form:"page"`
	//
	//分页大小
	PageSize             int32    `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize" form:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSearchResponse) Reset()         { *m = GetSearchResponse{} }
func (m *GetSearchResponse) String() string { return proto.CompactTextString(m) }
func (*GetSearchResponse) ProtoMessage()    {}
func (*GetSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_301727bc325f7fd5, []int{1}
}
func (m *GetSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSearchResponse.Unmarshal(m, b)
}
func (m *GetSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSearchResponse.Marshal(b, m, deterministic)
}
func (m *GetSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSearchResponse.Merge(m, src)
}
func (m *GetSearchResponse) XXX_Size() int {
	return xxx_messageInfo_GetSearchResponse.Size(m)
}
func (m *GetSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSearchResponse proto.InternalMessageInfo

func (m *GetSearchResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetSearchResponse) GetList() []*easy_flow.InstanceInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetSearchResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetSearchResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
//GetSearchApi返回
type GetSearchResponseWrapper struct {
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
	Data                 *GetSearchResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetSearchResponseWrapper) Reset()         { *m = GetSearchResponseWrapper{} }
func (m *GetSearchResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetSearchResponseWrapper) ProtoMessage()    {}
func (*GetSearchResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_301727bc325f7fd5, []int{2}
}
func (m *GetSearchResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSearchResponseWrapper.Unmarshal(m, b)
}
func (m *GetSearchResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSearchResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetSearchResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSearchResponseWrapper.Merge(m, src)
}
func (m *GetSearchResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetSearchResponseWrapper.Size(m)
}
func (m *GetSearchResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSearchResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetSearchResponseWrapper proto.InternalMessageInfo

func (m *GetSearchResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetSearchResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetSearchResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetSearchResponseWrapper) GetData() *GetSearchResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSearchRequest)(nil), "instance.GetSearchRequest")
	proto.RegisterType((*GetSearchResponse)(nil), "instance.GetSearchResponse")
	proto.RegisterType((*GetSearchResponseWrapper)(nil), "instance.GetSearchResponseWrapper")
}

func init() { proto.RegisterFile("search_by_post.proto", fileDescriptor_301727bc325f7fd5) }

var fileDescriptor_301727bc325f7fd5 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xdb, 0x6e, 0xe3, 0x44,
	0x18, 0xc7, 0x49, 0x9b, 0x94, 0x66, 0x02, 0x6c, 0x6a, 0x71, 0x30, 0x45, 0x22, 0x91, 0x09, 0x68,
	0xdc, 0xad, 0x9d, 0xd8, 0x49, 0xba, 0x6d, 0x84, 0x14, 0xb6, 0xe2, 0xa0, 0xde, 0x7a, 0x11, 0x2b,
	0x48, 0xd2, 0x68, 0x6a, 0x4f, 0xb2, 0x56, 0xdd, 0x8c, 0xf1, 0xb8, 0x2d, 0xbb, 0x89, 0xaf, 0x91,
	0xb8, 0x41, 0xe2, 0x31, 0x78, 0x0c, 0x6e, 0x78, 0x08, 0xa4, 0x20, 0xc1, 0x1b, 0xe4, 0x09, 0xd0,
	0xcc, 0x38, 0x76, 0x96, 0xac, 0xa1, 0x2b, 0x6d, 0x7b, 0x15, 0x7f, 0xa7, 0xff, 0xfc, 0xc6, 0xf1,
	0x7c, 0xf3, 0x81, 0xb7, 0x29, 0x46, 0x81, 0xfd, 0x64, 0x78, 0xf6, 0x74, 0xe8, 0x13, 0x1a, 0xea,
	0x7e, 0x40, 0x42, 0x22, 0x6d, 0xbb, 0x13, 0x1a, 0xa2, 0x89, 0x8d, 0x77, 0xb5, 0xb1, 0x1b, 0x3e,
	0xb9, 0x3c, 0xd3, 0x6d, 0x72, 0x51, 0x1f, 0x93, 0x31, 0xa9, 0xf3, 0x84, 0xb3, 0xcb, 0x11, 0xb7,
	0xb8, 0xc1, 0x9f, 0x44, 0xe1, 0xee, 0xe3, 0x31, 0xd1, 0x31, 0xa2, 0x4f, 0x89, 0x4f, 0x75, 0x8f,
	0xd8, 0xc8, 0xab, 0xdb, 0x64, 0x12, 0x06, 0xc8, 0x0e, 0xa9, 0xa8, 0x0c, 0xb0, 0x4f, 0xb4, 0x0b,
	0xe2, 0x60, 0x8f, 0xd6, 0xe3, 0xc4, 0x3a, 0x37, 0xb9, 0x35, 0x1c, 0x79, 0xe4, 0xba, 0xbe, 0x5c,
	0x7c, 0xe8, 0x4e, 0x46, 0x4b, 0xe1, 0x83, 0x15, 0x8e, 0x8b, 0x6b, 0x37, 0x3c, 0x27, 0xd7, 0xf5,
	0x31, 0xd1, 0x78, 0x50, 0xbb, 0x42, 0x9e, 0xeb, 0xa0, 0x90, 0x04, 0xb4, 0x9e, 0x3c, 0x8a, 0x3a,
	0xe5, 0xb7, 0x32, 0x28, 0x7f, 0x85, 0xc3, 0x47, 0x7c, 0x97, 0x16, 0xfe, 0xfe, 0x12, 0xd3, 0x50,
	0xfa, 0x16, 0x00, 0x1f, 0xd9, 0xe7, 0x68, 0x8c, 0x4f, 0x1c, 0x2a, 0xe7, 0xaa, 0x39, 0x58, 0x3c,
	0x3e, 0x5a, 0xcc, 0x2b, 0x3b, 0x23, 0x12, 0x5c, 0x74, 0x94, 0x34, 0xa6, 0xfc, 0xf5, 0x67, 0xa5,
	0x0a, 0x3e, 0x3c, 0xed, 0x35, 0xb4, 0x23, 0xa4, 0x3d, 0x7b, 0xa8, 0x7d, 0x37, 0xb8, 0x0f, 0xf7,
	0x57, 0x2d, 0x75, 0xaf, 0x66, 0xad, 0x88, 0x49, 0xdf, 0x80, 0xa2, 0x83, 0xaf, 0x5c, 0x9b, 0x2b,
	0x6f, 0x70, 0xe5, 0xc3, 0xc5, 0xbc, 0x52, 0x16, 0xca, 0x49, 0xe8, 0x66, 0xc2, 0xa9, 0x94, 0xf4,
	0x7b, 0x11, 0x6c, 0xc7, 0x96, 0x2f, 0x6f, 0x72, 0xdd, 0x5f, 0x8b, 0x8b, 0x79, 0xe5, 0xde, 0x73,
	0xc2, 0x3e, 0xd3, 0xfd, 0xa5, 0x08, 0x7e, 0x2e, 0x9e, 0x42, 0x68, 0xc2, 0x76, 0xaf, 0xa1, 0xb5,
	0x07, 0x53, 0x23, 0x9a, 0xf5, 0x1a, 0x5a, 0x6b, 0xd0, 0x77, 0xa6, 0x46, 0xa4, 0xb2, 0x67, 0x63,
	0xd0, 0x65, 0xc6, 0xbe, 0x19, 0xa9, 0xb0, 0xaf, 0xdf, 0x30, 0x53, 0x9d, 0x36, 0x23, 0x75, 0xd6,
	0xa7, 0x7b, 0x10, 0x42, 0x46, 0xfa, 0x50, 0xfb, 0x12, 0x69, 0xa3, 0xc1, 0xd4, 0xd8, 0x6f, 0x45,
	0x1d, 0x75, 0xfa, 0x20, 0x5a, 0xf3, 0xce, 0x3a, 0xaa, 0x3a, 0x7b, 0x61, 0xf2, 0x41, 0x04, 0x3b,
	0x6b, 0xd9, 0x10, 0x9a, 0x82, 0x63, 0x66, 0xc6, 0x14, 0x33, 0xa3, 0xef, 0xf4, 0x9d, 0x59, 0xcf,
	0xd0, 0x8e, 0x18, 0x87, 0x80, 0xfd, 0x9f, 0x1c, 0x81, 0x99, 0xb9, 0x72, 0x3b, 0x82, 0x70, 0x7d,
	0x6d, 0x55, 0x6c, 0x71, 0xd6, 0xb9, 0x13, 0x86, 0x56, 0x26, 0x03, 0x2b, 0x7b, 0x51, 0xa8, 0xfb,
	0x2a, 0xc1, 0xfe, 0x83, 0xac, 0x99, 0x49, 0xd6, 0xca, 0x20, 0x9b, 0x36, 0xf6, 0xcd, 0xe8, 0x8e,
	0xe8, 0xcc, 0x4c, 0xba, 0x76, 0x36, 0x5d, 0xf3, 0xae, 0xe8, 0x8c, 0x4c, 0xba, 0x83, 0x6c, 0xba,
	0xd6, 0x6d, 0xd0, 0x75, 0xb2, 0x40, 0x1e, 0x64, 0x83, 0xb4, 0x5f, 0x3d, 0x88, 0x0a, 0x3f, 0xd6,
	0xef, 0xab, 0xdd, 0x3e, 0xdd, 0xab, 0x59, 0x49, 0xf3, 0x92, 0x54, 0xb0, 0x85, 0x7c, 0x9f, 0xb5,
	0xc7, 0x3c, 0x6f, 0x63, 0x3b, 0x8b, 0x79, 0xe5, 0x4d, 0xd1, 0xc5, 0x84, 0x5f, 0xb1, 0xe2, 0x04,
	0x49, 0x05, 0x79, 0x1f, 0x8d, 0xb1, 0x5c, 0xa8, 0xe6, 0x60, 0xe1, 0xf8, 0x9d, 0xc5, 0xbc, 0x52,
	0x5a, 0x76, 0xe8, 0x31, 0x66, 0xad, 0x6e, 0xa3, 0xfc, 0x9a, 0xc5, 0x53, 0xa4, 0x36, 0xd8, 0x66,
	0xbf, 0x8f, 0xdc, 0x67, 0x58, 0xde, 0xe2, 0xe9, 0xef, 0xa7, 0xdd, 0x71, 0x19, 0x59, 0x96, 0x24,
	0xa9, 0xd2, 0xe7, 0xa0, 0x40, 0x02, 0x07, 0x07, 0xf2, 0xeb, 0x9c, 0x45, 0x5f, 0xcc, 0x2b, 0x6f,
	0x88, 0x1a, 0xee, 0x66, 0x05, 0xbb, 0x40, 0x3e, 0xed, 0x89, 0xae, 0x3c, 0x64, 0xaf, 0xeb, 0xa0,
	0x15, 0xc1, 0xaa, 0x83, 0xa9, 0xad, 0x76, 0x2d, 0x51, 0x2c, 0xfd, 0x94, 0x03, 0xc5, 0xe4, 0x0e,
	0x90, 0xb7, 0xb9, 0x94, 0x97, 0x76, 0xfd, 0x24, 0xc4, 0xe4, 0xbe, 0x06, 0xd6, 0x69, 0xaf, 0x2b,
	0xf4, 0x1a, 0xda, 0xd1, 0x60, 0x7a, 0x18, 0x69, 0xcf, 0xd9, 0xad, 0x97, 0xb4, 0x0d, 0x33, 0xaa,
	0x59, 0xe9, 0xf2, 0x1c, 0xe6, 0x0a, 0x07, 0xd4, 0x25, 0x93, 0x13, 0x47, 0x2e, 0xfe, 0x1b, 0x26,
	0x09, 0xdd, 0x22, 0x4c, 0xb2, 0x86, 0xf4, 0x63, 0x2e, 0xb9, 0xb6, 0x1c, 0x19, 0x70, 0x96, 0xf3,
	0xb5, 0x5b, 0xeb, 0x16, 0x51, 0x92, 0xc5, 0xa5, 0x0e, 0x28, 0xf0, 0xaf, 0x4a, 0x2e, 0x71, 0x8a,
	0x5a, 0xfa, 0x4f, 0x73, 0x37, 0x43, 0x28, 0x83, 0xb7, 0x96, 0x17, 0xf2, 0x60, 0x6a, 0x34, 0xa3,
	0x9a, 0x25, 0x4a, 0x94, 0x3f, 0x72, 0x60, 0x67, 0x65, 0x88, 0xa0, 0x3e, 0x99, 0x50, 0x2c, 0x7d,
	0x02, 0x0a, 0x21, 0x09, 0x91, 0xc7, 0x07, 0x88, 0xc2, 0x71, 0x39, 0x55, 0xe4, 0x6e, 0xc5, 0x12,
	0x61, 0xe9, 0x53, 0x90, 0xf7, 0x5c, 0x1a, 0xca, 0x1b, 0xd5, 0x4d, 0x58, 0x32, 0xdf, 0xd3, 0x93,
	0x41, 0x47, 0x3f, 0x89, 0x07, 0x9d, 0x93, 0xc9, 0x88, 0x1c, 0xdf, 0x4b, 0x3f, 0x6f, 0x96, 0xae,
	0x58, 0xbc, 0x2a, 0x39, 0x03, 0x9b, 0x2f, 0x77, 0x06, 0xf2, 0x37, 0x3e, 0x03, 0xca, 0xdf, 0x39,
	0x20, 0xaf, 0xed, 0xee, 0x71, 0x80, 0x7c, 0x1f, 0x07, 0xd2, 0x47, 0x20, 0x6f, 0x13, 0x07, 0xc7,
	0x7b, 0x5c, 0x61, 0x64, 0x5e, 0xc5, 0xe2, 0x41, 0xe9, 0x10, 0x94, 0xd8, 0xef, 0x17, 0x3f, 0xf8,
	0x1e, 0x72, 0x27, 0xf1, 0xd8, 0xf3, 0xee, 0x62, 0x5e, 0x91, 0xd2, 0xdc, 0x38, 0xa8, 0x58, 0xab,
	0xa9, 0xec, 0x1d, 0xe2, 0x20, 0x20, 0x41, 0x3c, 0xd2, 0xac, 0xbc, 0x43, 0xee, 0x56, 0x2c, 0x11,
	0x96, 0x3e, 0x03, 0x79, 0x07, 0x85, 0x88, 0x6f, 0xab, 0x64, 0x7e, 0xa0, 0x2f, 0x47, 0x44, 0x7d,
	0x0d, 0x7c, 0x95, 0x91, 0x95, 0x28, 0x16, 0xaf, 0x3c, 0xdb, 0xe2, 0xf3, 0x60, 0xf3, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0x72, 0x69, 0xdb, 0xf1, 0x0a, 0x00, 0x00,
}