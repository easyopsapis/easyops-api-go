// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package solution

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	app_store "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/app_store"
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
//GetSolution请求
type GetSolutionRequest struct {
	//
	//solution的instanceId
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSolutionRequest) Reset()         { *m = GetSolutionRequest{} }
func (m *GetSolutionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSolutionRequest) ProtoMessage()    {}
func (*GetSolutionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetSolutionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSolutionRequest.Unmarshal(m, b)
}
func (m *GetSolutionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSolutionRequest.Marshal(b, m, deterministic)
}
func (m *GetSolutionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSolutionRequest.Merge(m, src)
}
func (m *GetSolutionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSolutionRequest.Size(m)
}
func (m *GetSolutionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSolutionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSolutionRequest proto.InternalMessageInfo

func (m *GetSolutionRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

//
//GetSolution返回
type GetSolutionResponse struct {
	//
	//关联的小产品
	MicroApps *app_store.AppStoreMicroApp `protobuf:"bytes,1,opt,name=microApps,proto3" json:"microApps" form:"microApps"`
	//
	//实例id
	InstanceId string `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	//
	//名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//markdown格式的内容
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content" form:"content"`
	//
	//轮播图片
	Thumbnail string `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail" form:"thumbnail"`
	//
	//简介
	Brief string `protobuf:"bytes,6,opt,name=brief,proto3" json:"brief" form:"brief"`
	//
	//图标
	Icon                 *GetSolutionResponse_Icon `protobuf:"bytes,7,opt,name=icon,proto3" json:"icon" form:"icon"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetSolutionResponse) Reset()         { *m = GetSolutionResponse{} }
func (m *GetSolutionResponse) String() string { return proto.CompactTextString(m) }
func (*GetSolutionResponse) ProtoMessage()    {}
func (*GetSolutionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetSolutionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSolutionResponse.Unmarshal(m, b)
}
func (m *GetSolutionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSolutionResponse.Marshal(b, m, deterministic)
}
func (m *GetSolutionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSolutionResponse.Merge(m, src)
}
func (m *GetSolutionResponse) XXX_Size() int {
	return xxx_messageInfo_GetSolutionResponse.Size(m)
}
func (m *GetSolutionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSolutionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSolutionResponse proto.InternalMessageInfo

func (m *GetSolutionResponse) GetMicroApps() *app_store.AppStoreMicroApp {
	if m != nil {
		return m.MicroApps
	}
	return nil
}

func (m *GetSolutionResponse) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *GetSolutionResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSolutionResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *GetSolutionResponse) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *GetSolutionResponse) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

func (m *GetSolutionResponse) GetIcon() *GetSolutionResponse_Icon {
	if m != nil {
		return m.Icon
	}
	return nil
}

type GetSolutionResponse_Icon struct {
	//
	//图标库
	Lib string `protobuf:"bytes,1,opt,name=lib,proto3" json:"lib" form:"lib"`
	//
	//图标url
	Icon                 string   `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon" form:"icon"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSolutionResponse_Icon) Reset()         { *m = GetSolutionResponse_Icon{} }
func (m *GetSolutionResponse_Icon) String() string { return proto.CompactTextString(m) }
func (*GetSolutionResponse_Icon) ProtoMessage()    {}
func (*GetSolutionResponse_Icon) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1, 0}
}
func (m *GetSolutionResponse_Icon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSolutionResponse_Icon.Unmarshal(m, b)
}
func (m *GetSolutionResponse_Icon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSolutionResponse_Icon.Marshal(b, m, deterministic)
}
func (m *GetSolutionResponse_Icon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSolutionResponse_Icon.Merge(m, src)
}
func (m *GetSolutionResponse_Icon) XXX_Size() int {
	return xxx_messageInfo_GetSolutionResponse_Icon.Size(m)
}
func (m *GetSolutionResponse_Icon) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSolutionResponse_Icon.DiscardUnknown(m)
}

var xxx_messageInfo_GetSolutionResponse_Icon proto.InternalMessageInfo

func (m *GetSolutionResponse_Icon) GetLib() string {
	if m != nil {
		return m.Lib
	}
	return ""
}

func (m *GetSolutionResponse_Icon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

//
//GetSolutionApi返回
type GetSolutionResponseWrapper struct {
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
	Data                 *GetSolutionResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetSolutionResponseWrapper) Reset()         { *m = GetSolutionResponseWrapper{} }
func (m *GetSolutionResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetSolutionResponseWrapper) ProtoMessage()    {}
func (*GetSolutionResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetSolutionResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSolutionResponseWrapper.Unmarshal(m, b)
}
func (m *GetSolutionResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSolutionResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetSolutionResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSolutionResponseWrapper.Merge(m, src)
}
func (m *GetSolutionResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetSolutionResponseWrapper.Size(m)
}
func (m *GetSolutionResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSolutionResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetSolutionResponseWrapper proto.InternalMessageInfo

func (m *GetSolutionResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetSolutionResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetSolutionResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetSolutionResponseWrapper) GetData() *GetSolutionResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSolutionRequest)(nil), "solution.GetSolutionRequest")
	proto.RegisterType((*GetSolutionResponse)(nil), "solution.GetSolutionResponse")
	proto.RegisterType((*GetSolutionResponse_Icon)(nil), "solution.GetSolutionResponse.Icon")
	proto.RegisterType((*GetSolutionResponseWrapper)(nil), "solution.GetSolutionResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6e, 0xd3, 0x3c,
	0x18, 0xc6, 0xd5, 0x6f, 0xdd, 0xf6, 0xd5, 0x45, 0x63, 0x18, 0x84, 0xaa, 0x22, 0x94, 0xc9, 0x4c,
	0xa8, 0xdd, 0x48, 0x02, 0x9b, 0x84, 0x60, 0x42, 0xa0, 0x4d, 0x42, 0x53, 0x0f, 0x10, 0x92, 0x77,
	0x80, 0x44, 0xbd, 0x4e, 0x4e, 0xea, 0x65, 0x11, 0x49, 0x6c, 0x1c, 0x97, 0xc1, 0x18, 0x37, 0xc2,
	0x05, 0x71, 0x19, 0x99, 0xb4, 0x13, 0xce, 0x73, 0x05, 0xc8, 0x76, 0xda, 0x86, 0x31, 0x71, 0xc2,
	0x51, 0x6c, 0x3f, 0xbf, 0xe7, 0xfd, 0xe3, 0xd7, 0x01, 0xad, 0x88, 0x29, 0x4f, 0x48, 0xae, 0x38,
	0xfc, 0x3f, 0xe7, 0xc9, 0x44, 0xc5, 0x3c, 0xeb, 0xba, 0x51, 0xac, 0x4e, 0x26, 0x81, 0x17, 0xf2,
	0xd4, 0x8f, 0x78, 0xc4, 0x7d, 0x03, 0x04, 0x93, 0x63, 0xb3, 0x33, 0x1b, 0xb3, 0xb2, 0xc6, 0xee,
	0x41, 0xc4, 0x3d, 0x46, 0xf3, 0x2f, 0x5c, 0xe4, 0x5e, 0xc2, 0x43, 0x9a, 0xf8, 0x21, 0xcf, 0x94,
	0xa4, 0xa1, 0xca, 0xad, 0x53, 0x32, 0xc1, 0xdd, 0x94, 0x8f, 0x59, 0x92, 0xfb, 0x15, 0xe8, 0x9b,
	0xad, 0x4f, 0x85, 0x38, 0xca, 0x35, 0xe0, 0xa7, 0x71, 0x28, 0xf9, 0x11, 0x15, 0xa2, 0x0a, 0xfa,
	0xb4, 0x56, 0x43, 0x7a, 0x1a, 0xab, 0x0f, 0xfc, 0xd4, 0x8f, 0xb8, 0x6b, 0x44, 0xf7, 0x13, 0x4d,
	0xe2, 0x31, 0x55, 0x5c, 0xe6, 0xfe, 0x6c, 0x69, 0x7d, 0xe8, 0x08, 0xc0, 0x7d, 0xa6, 0x0e, 0xaa,
	0x56, 0x30, 0xfb, 0x38, 0x61, 0xb9, 0x82, 0x03, 0x00, 0xe2, 0x2c, 0x57, 0x34, 0x0b, 0xd9, 0x60,
	0xdc, 0x69, 0xac, 0x35, 0x7a, 0xad, 0xbd, 0x7e, 0x59, 0x38, 0xb7, 0x8e, 0xb9, 0x4c, 0x77, 0xd0,
	0x5c, 0x43, 0x97, 0x17, 0xce, 0x2a, 0x58, 0x19, 0x0d, 0x1f, 0xbb, 0xcf, 0xa9, 0x7b, 0x76, 0xf8,
	0xf5, 0xc9, 0xf6, 0xb7, 0x75, 0x5c, 0x33, 0xa3, 0x1f, 0x4d, 0x70, 0xfb, 0xb7, 0x0c, 0xb9, 0xe0,
	0x59, 0xce, 0xe0, 0x5b, 0xd0, 0x32, 0x3d, 0xec, 0x0a, 0x91, 0x9b, 0x0c, 0xed, 0xad, 0x7b, 0xde,
	0xac, 0x3f, 0x6f, 0x57, 0x88, 0x03, 0xbd, 0x78, 0x53, 0x31, 0x7b, 0x77, 0xca, 0xc2, 0x59, 0xb5,
	0xe9, 0x67, 0x3e, 0x84, 0xe7, 0x31, 0xae, 0xd4, 0xfc, 0xdf, 0x3f, 0xd4, 0x0c, 0x1f, 0x80, 0x66,
	0x46, 0x53, 0xd6, 0x59, 0x30, 0x41, 0x6e, 0x96, 0x85, 0xd3, 0xb6, 0x41, 0xf4, 0x29, 0xc2, 0x46,
	0x84, 0x8f, 0xc0, 0xb2, 0x9e, 0x1b, 0xcb, 0x54, 0xa7, 0x69, 0x38, 0x58, 0x16, 0xce, 0x8a, 0xe5,
	0x2a, 0x01, 0xe1, 0x29, 0x02, 0xb7, 0x40, 0x4b, 0x9d, 0x4c, 0xd2, 0x20, 0xa3, 0x71, 0xd2, 0x59,
	0x34, 0x7c, 0xad, 0xa3, 0x99, 0x84, 0xf0, 0x1c, 0x83, 0x0f, 0xc1, 0x62, 0x20, 0x63, 0x76, 0xdc,
	0x59, 0x32, 0xfc, 0x6a, 0x59, 0x38, 0x37, 0x2c, 0x6f, 0x8e, 0x11, 0xb6, 0x32, 0xdc, 0x07, 0xcd,
	0x38, 0xe4, 0x59, 0x67, 0xd9, 0xdc, 0x22, 0xf2, 0xa6, 0x0f, 0xd3, 0xbb, 0xe6, 0xde, 0xbd, 0x41,
	0xc8, 0xb3, 0x7a, 0x4b, 0xda, 0x89, 0xb0, 0x09, 0xd0, 0xfd, 0xde, 0x00, 0x4d, 0xad, 0xc3, 0x35,
	0xb0, 0x90, 0xc4, 0x41, 0x35, 0xf8, 0x95, 0xb2, 0x70, 0x80, 0x85, 0x93, 0x38, 0x40, 0x58, 0x4b,
	0x50, 0x54, 0x39, 0xed, 0x3d, 0x93, 0x2b, 0xf1, 0x2e, 0x2f, 0x9c, 0x01, 0xd8, 0x1f, 0xf5, 0x7a,
	0xc4, 0x1f, 0x8e, 0x88, 0xbf, 0x43, 0x36, 0xc8, 0x2b, 0x84, 0x5e, 0xbc, 0x24, 0xe7, 0x44, 0x92,
	0xec, 0x70, 0xb3, 0xbf, 0xd9, 0x3f, 0xef, 0x11, 0xbf, 0x7f, 0x3e, 0xa4, 0xee, 0xd9, 0xae, 0xfb,
	0xfe, 0x70, 0xa7, 0x47, 0xc8, 0x70, 0x44, 0xc8, 0x9f, 0xe4, 0xc6, 0xba, 0x2d, 0x0e, 0xfd, 0x6c,
	0x80, 0xee, 0x35, 0x0d, 0xbd, 0x93, 0x54, 0x08, 0x26, 0xf5, 0xcc, 0x42, 0x3e, 0x66, 0xa6, 0xe6,
	0xc5, 0x7a, 0x83, 0xfa, 0x14, 0x61, 0x23, 0xc2, 0x67, 0xa0, 0xad, 0xbf, 0xaf, 0x3f, 0x8b, 0x84,
	0xc6, 0xd3, 0xe2, 0xef, 0x96, 0x85, 0x03, 0xe7, 0x6c, 0x25, 0x22, 0x5c, 0x47, 0xf5, 0x2c, 0x98,
	0x94, 0x5c, 0x56, 0x6f, 0xa2, 0x36, 0x0b, 0x73, 0x8c, 0xb0, 0x95, 0xe1, 0x1e, 0x68, 0x8e, 0xa9,
	0xa2, 0xe6, 0x49, 0xb4, 0xb7, 0xee, 0xff, 0x75, 0x16, 0xf5, 0x2a, 0xb5, 0x09, 0x61, 0xe3, 0x0d,
	0x96, 0xcc, 0xaf, 0xb9, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x65, 0xde, 0xee, 0xe2, 0x6d, 0x04,
	0x00, 0x00,
}
