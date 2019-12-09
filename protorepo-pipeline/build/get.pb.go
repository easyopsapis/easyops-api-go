// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	pipeline "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
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
//Get请求
type GetRequest struct {
	//
	//构建任务id
	BuildId              string   `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id" form:"build_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

//
//Get返回
type GetResponse struct {
	//
	//项目
	Project *pipeline.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project" form:"project"`
	//
	//流水线
	Pipeline *pipeline.Pipeline `protobuf:"bytes,2,opt,name=pipeline,proto3" json:"pipeline" form:"pipeline"`
	//
	//stages执行状态
	Stages []*pipeline.StageStatus `protobuf:"bytes,3,rep,name=stages,proto3" json:"stages" form:"stages"`
	//
	//任务id, 服务端自动生成
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id" form:"id"`
	//
	//git信息, 创建的时候传入，不能修改
	GitMeta *pipeline.GitMeta `protobuf:"bytes,5,opt,name=git_meta,json=gitMeta,proto3" json:"git_meta" form:"git_meta"`
	//
	//触发者，创建的时候传入，不能修改
	Sender string `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender" form:"sender"`
	//
	//制品包信息
	Artifact *GetResponse_Artifact `protobuf:"bytes,7,opt,name=artifact,proto3" json:"artifact" form:"artifact"`
	//
	//创建时间，创建的时候传入，不能修改
	Created int32 `protobuf:"varint,8,opt,name=created,proto3" json:"created" form:"created"`
	//
	//workflow定义
	YamlString string `protobuf:"bytes,9,opt,name=yaml_string,json=yamlString,proto3" json:"yaml_string" form:"yaml_string"`
	//
	//执行状态
	Status *pipeline.BuildStatus `protobuf:"bytes,10,opt,name=status,proto3" json:"status" form:"status"`
	//
	//编号
	Number string `protobuf:"bytes,11,opt,name=number,proto3" json:"number" form:"number"`
	//
	//事件列表
	Events               []string `protobuf:"bytes,12,rep,name=events,proto3" json:"events" form:"events"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetProject() *pipeline.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *GetResponse) GetPipeline() *pipeline.Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *GetResponse) GetStages() []*pipeline.StageStatus {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *GetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetResponse) GetGitMeta() *pipeline.GitMeta {
	if m != nil {
		return m.GitMeta
	}
	return nil
}

func (m *GetResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *GetResponse) GetArtifact() *GetResponse_Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func (m *GetResponse) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *GetResponse) GetYamlString() string {
	if m != nil {
		return m.YamlString
	}
	return ""
}

func (m *GetResponse) GetStatus() *pipeline.BuildStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetResponse) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *GetResponse) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetResponse_Artifact struct {
	//
	//制品包名称
	PackageName string `protobuf:"bytes,1,opt,name=packageName,proto3" json:"packageName" form:"packageName"`
	//
	//版本号
	VersionName string `protobuf:"bytes,2,opt,name=versionName,proto3" json:"versionName" form:"versionName"`
	//
	//制品包上传时间
	Ctime string `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime" form:"ctime"`
	//
	//程序包Id
	PackageId string `protobuf:"bytes,4,opt,name=packageId,proto3" json:"packageId" form:"packageId"`
	//
	//版本Id
	VersionId            string   `protobuf:"bytes,5,opt,name=versionId,proto3" json:"versionId" form:"versionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse_Artifact) Reset()         { *m = GetResponse_Artifact{} }
func (m *GetResponse_Artifact) String() string { return proto.CompactTextString(m) }
func (*GetResponse_Artifact) ProtoMessage()    {}
func (*GetResponse_Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{1, 0}
}
func (m *GetResponse_Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse_Artifact.Unmarshal(m, b)
}
func (m *GetResponse_Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse_Artifact.Marshal(b, m, deterministic)
}
func (m *GetResponse_Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse_Artifact.Merge(m, src)
}
func (m *GetResponse_Artifact) XXX_Size() int {
	return xxx_messageInfo_GetResponse_Artifact.Size(m)
}
func (m *GetResponse_Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse_Artifact proto.InternalMessageInfo

func (m *GetResponse_Artifact) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *GetResponse_Artifact) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *GetResponse_Artifact) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *GetResponse_Artifact) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetResponse_Artifact) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//GetApi返回
type GetResponseWrapper struct {
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
	Data                 *GetResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponseWrapper) Reset()         { *m = GetResponseWrapper{} }
func (m *GetResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*GetResponseWrapper) ProtoMessage()    {}
func (*GetResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b2a6be0e6d8388, []int{2}
}
func (m *GetResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseWrapper.Unmarshal(m, b)
}
func (m *GetResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseWrapper.Marshal(b, m, deterministic)
}
func (m *GetResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseWrapper.Merge(m, src)
}
func (m *GetResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_GetResponseWrapper.Size(m)
}
func (m *GetResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseWrapper proto.InternalMessageInfo

func (m *GetResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *GetResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetResponseWrapper) GetData() *GetResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "build.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "build.GetResponse")
	proto.RegisterType((*GetResponse_Artifact)(nil), "build.GetResponse.Artifact")
	proto.RegisterType((*GetResponseWrapper)(nil), "build.GetResponseWrapper")
}

func init() { proto.RegisterFile("get.proto", fileDescriptor_21b2a6be0e6d8388) }

var fileDescriptor_21b2a6be0e6d8388 = []byte{
	// 939 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x93, 0xf8, 0x6f, 0x5c, 0xfa, 0x33, 0x08, 0xb4, 0x0a, 0x17, 0x76, 0xa7, 0x11, 0xac,
	0x69, 0x67, 0xfd, 0x57, 0x68, 0x9a, 0x1b, 0x13, 0x23, 0xa8, 0x22, 0x01, 0x4a, 0x27, 0x20, 0xa4,
	0x2e, 0x4e, 0x34, 0xde, 0x9d, 0x2c, 0x4b, 0x6c, 0xcf, 0xb2, 0x3b, 0x4e, 0x69, 0x13, 0x3f, 0x00,
	0xaf, 0x82, 0xc4, 0x4b, 0xf0, 0x0c, 0xdc, 0x70, 0xe3, 0x4a, 0x3c, 0x82, 0x9f, 0x00, 0xed, 0xd9,
	0x59, 0x7b, 0xd2, 0x96, 0x0b, 0xa4, 0x70, 0xb5, 0x73, 0xe6, 0xfb, 0xbe, 0x33, 0xdf, 0x9c, 0x73,
	0xc6, 0x46, 0xd5, 0x40, 0x28, 0x27, 0x8a, 0xa5, 0x92, 0xb8, 0x38, 0x9a, 0x85, 0x63, 0x7f, 0x9b,
	0x06, 0xa1, 0xfa, 0x71, 0x36, 0x72, 0x3c, 0x39, 0x69, 0x05, 0x32, 0x90, 0x2d, 0x40, 0x47, 0xb3,
	0x53, 0x88, 0x20, 0x80, 0x55, 0xa6, 0xda, 0x3e, 0x0c, 0xa4, 0x23, 0x78, 0xf2, 0x42, 0x46, 0x89,
	0x33, 0x96, 0x1e, 0x1f, 0xb7, 0x3c, 0x39, 0x55, 0x31, 0xf7, 0x54, 0x92, 0x29, 0x63, 0x11, 0x49,
	0x3a, 0x91, 0xbe, 0x18, 0x27, 0x2d, 0x4d, 0x6c, 0x41, 0xd8, 0x8a, 0xc2, 0x48, 0x8c, 0xc3, 0xa9,
	0x48, 0x79, 0x3f, 0x09, 0x4f, 0xfb, 0xd8, 0x7e, 0x7a, 0x1d, 0x19, 0xf5, 0x42, 0xa7, 0xfc, 0xee,
	0x1a, 0x52, 0x26, 0x8a, 0x07, 0xe2, 0x24, 0x51, 0x5c, 0xcd, 0x92, 0x6b, 0x74, 0x1a, 0x84, 0xea,
	0x64, 0x22, 0x14, 0xbf, 0x46, 0xa7, 0xd0, 0xc7, 0xab, 0x4e, 0x3f, 0x35, 0x9a, 0x3a, 0x79, 0x1e,
	0xaa, 0x33, 0xf9, 0xbc, 0x15, 0x48, 0x0a, 0x20, 0x3d, 0xe7, 0xe3, 0xd0, 0xe7, 0x4a, 0xc6, 0x49,
	0x6b, 0xb5, 0xcc, 0x74, 0xe4, 0x10, 0xa1, 0x27, 0x42, 0x31, 0xf1, 0xf3, 0x4c, 0x24, 0x0a, 0x0f,
	0x50, 0x25, 0xcb, 0x1d, 0xfa, 0x56, 0xa1, 0x51, 0xb0, 0xab, 0x83, 0x8f, 0x96, 0x8b, 0xfa, 0xad,
	0x53, 0x19, 0x4f, 0xf6, 0x48, 0x8e, 0x90, 0xbf, 0x5f, 0xd5, 0x6f, 0xa3, 0x9b, 0xc7, 0x6e, 0x9b,
	0x3e, 0xe6, 0xf4, 0xe5, 0xf0, 0xa2, 0xd3, 0x9b, 0xef, 0xb0, 0x32, 0xc0, 0x07, 0x3e, 0x79, 0x85,
	0x50, 0x0d, 0x52, 0x26, 0x91, 0x9c, 0x26, 0x02, 0xf7, 0x51, 0x59, 0xb7, 0x1f, 0x52, 0xd6, 0xba,
	0x77, 0x9c, 0x55, 0xf3, 0x0e, 0x33, 0x60, 0x80, 0x97, 0x8b, 0xfa, 0xcd, 0xec, 0x14, 0xcd, 0x25,
	0x2c, 0x57, 0xe1, 0xcf, 0x51, 0x25, 0x17, 0x58, 0x1b, 0x90, 0x01, 0x1b, 0x19, 0xf4, 0x62, 0xf0,
	0xee, 0xda, 0x68, 0x0e, 0x12, 0xb6, 0x12, 0xe2, 0xcf, 0x50, 0x09, 0xfa, 0x9b, 0x58, 0x9b, 0x8d,
	0x4d, 0xbb, 0xd6, 0x7d, 0x6f, 0x9d, 0xe2, 0x28, 0xdd, 0x3f, 0x82, 0x62, 0x0e, 0xee, 0x2c, 0x17,
	0xf5, 0x77, 0xb2, 0x2c, 0x19, 0x9d, 0x30, 0xad, 0xc3, 0x1d, 0xb4, 0x11, 0xfa, 0xd6, 0x16, 0x54,
	0xe5, 0xee, 0x72, 0x51, 0xaf, 0x66, 0xb4, 0x7f, 0xab, 0xc7, 0x46, 0xe8, 0xe3, 0x7d, 0x54, 0xc9,
	0xbb, 0x6f, 0x15, 0x5f, 0xbf, 0xfb, 0x93, 0x50, 0x7d, 0x2d, 0x14, 0x37, 0x8d, 0xe7, 0x64, 0xc2,
	0xca, 0x41, 0x86, 0xe2, 0x43, 0x54, 0x4a, 0xc4, 0xd4, 0x17, 0xb1, 0x55, 0x82, 0x93, 0x77, 0x0d,
	0x83, 0xb0, 0x9f, 0x9e, 0x7e, 0x0f, 0xdd, 0x3d, 0x76, 0x39, 0x7d, 0xb9, 0x4f, 0x9f, 0xb5, 0xe9,
	0xe3, 0xa1, 0xeb, 0xac, 0xd6, 0x27, 0x74, 0x78, 0xd1, 0x7d, 0xd0, 0xeb, 0xcc, 0x77, 0x98, 0xce,
	0x83, 0xbf, 0x42, 0x15, 0x1e, 0xab, 0xf0, 0x94, 0x7b, 0xca, 0x2a, 0x83, 0xa9, 0x0f, 0x1c, 0xe8,
	0x9d, 0x63, 0x74, 0xcd, 0xd9, 0xd7, 0x14, 0xd3, 0x5e, 0x2e, 0x23, 0x6c, 0x95, 0x01, 0x3f, 0x40,
	0x65, 0x2f, 0x16, 0x5c, 0x09, 0xdf, 0xaa, 0x34, 0x0a, 0x76, 0xd1, 0x6c, 0xa5, 0x06, 0x08, 0xcb,
	0x29, 0xf8, 0x11, 0xaa, 0xbd, 0xe0, 0x93, 0xf1, 0x49, 0xa2, 0xe2, 0x70, 0x1a, 0x58, 0x55, 0xb8,
	0xd2, 0xfb, 0xcb, 0x45, 0x1d, 0x67, 0x0a, 0x03, 0x24, 0x0c, 0xa5, 0xd1, 0x11, 0x04, 0xba, 0x7d,
	0x6a, 0x96, 0x58, 0x08, 0x2c, 0x1b, 0xed, 0x1b, 0xa4, 0xde, 0xdf, 0xda, 0x3e, 0x35, 0xd3, 0xed,
	0x53, 0xb3, 0x04, 0x37, 0x51, 0x69, 0x3a, 0x9b, 0x8c, 0x44, 0x6c, 0xd5, 0xe0, 0x54, 0x83, 0x9a,
	0xed, 0x13, 0xa6, 0x09, 0x29, 0x55, 0x9c, 0x8b, 0xa9, 0x4a, 0xac, 0x1b, 0x8d, 0xcd, 0xab, 0xd4,
	0x6c, 0x9f, 0x30, 0x4d, 0xd8, 0x5e, 0x6e, 0xa1, 0x4a, 0x5e, 0x2a, 0xbc, 0x8b, 0x6a, 0x11, 0xf7,
	0xce, 0x78, 0x20, 0xbe, 0xe1, 0x13, 0xa1, 0x1f, 0x90, 0x71, 0x3b, 0x03, 0x24, 0xcc, 0xa4, 0xa6,
	0xca, 0x73, 0x11, 0x27, 0xa1, 0x9c, 0x82, 0x72, 0xe3, 0x75, 0xa5, 0x01, 0x12, 0x66, 0x52, 0xf1,
	0x5f, 0x05, 0x54, 0xf4, 0x54, 0x38, 0x11, 0xd6, 0x26, 0x88, 0xfe, 0x28, 0x2c, 0x17, 0xf5, 0x1b,
	0xba, 0xfe, 0xe9, 0x7e, 0x3a, 0x1f, 0xbf, 0x17, 0xd0, 0x6f, 0x85, 0x63, 0xdb, 0xee, 0xef, 0xb9,
	0x9d, 0x74, 0x3e, 0xd2, 0x21, 0xf9, 0xb8, 0xd9, 0x87, 0xef, 0xc5, 0xc3, 0x79, 0x93, 0xda, 0x1d,
	0xb7, 0x4d, 0xbb, 0xc3, 0xcb, 0x36, 0xe0, 0x4d, 0x6a, 0xf7, 0xdc, 0x36, 0xed, 0xe4, 0xf1, 0xa5,
	0xdb, 0xa1, 0xdd, 0x4c, 0xd5, 0x74, 0xbf, 0x6d, 0x0c, 0xed, 0xae, 0xdb, 0xa6, 0xbd, 0xe1, 0x25,
	0x70, 0xb2, 0xed, 0x3d, 0xdb, 0x6d, 0xd3, 0x4f, 0xf2, 0x60, 0xbd, 0xb6, 0x7f, 0x70, 0xe0, 0x7b,
	0xbf, 0xd9, 0xb7, 0x9f, 0x5d, 0xba, 0xf7, 0xe9, 0xd0, 0xee, 0xef, 0xbd, 0x45, 0x6e, 0xa8, 0xfb,
	0x3b, 0x2c, 0xbb, 0x11, 0xfe, 0xb5, 0x80, 0xaa, 0xba, 0x4a, 0x07, 0xf9, 0xcb, 0x3b, 0x5b, 0x2e,
	0xea, 0xb7, 0xaf, 0x94, 0xf3, 0x00, 0x1e, 0xe0, 0x11, 0x7a, 0x9a, 0x3e, 0x81, 0xd3, 0x7d, 0xfa,
	0x25, 0xdc, 0x6a, 0x77, 0x4e, 0xfb, 0x66, 0xfc, 0xf0, 0x3f, 0xc6, 0x9d, 0xee, 0x7c, 0x87, 0xad,
	0x4f, 0x07, 0x2f, 0xba, 0xee, 0x07, 0x3e, 0x3c, 0xe6, 0x2b, 0x5e, 0x56, 0xd0, 0xff, 0xe7, 0x65,
	0x7d, 0xc4, 0x9f, 0x05, 0x84, 0x8d, 0xb7, 0xfa, 0x7d, 0xcc, 0xa3, 0x48, 0xc4, 0xf8, 0x1e, 0xda,
	0xf2, 0xa4, 0x9f, 0xcd, 0x5d, 0x71, 0x70, 0x6b, 0xb9, 0xa8, 0xd7, 0xf4, 0x1c, 0x48, 0x5f, 0x10,
	0x06, 0x60, 0x3a, 0x69, 0xe9, 0xf7, 0x8b, 0x5f, 0xa2, 0x31, 0x0f, 0xa7, 0x6f, 0x4e, 0x9a, 0x01,
	0x12, 0x66, 0x52, 0xf1, 0x87, 0xa8, 0x28, 0xe2, 0x58, 0xc6, 0x7a, 0xd0, 0x6e, 0xaf, 0xe7, 0x0c,
	0xb6, 0x09, 0xcb, 0x60, 0xfc, 0x08, 0x6d, 0xf9, 0x5c, 0x71, 0xe8, 0x57, 0xfa, 0x53, 0xfd, 0xc6,
	0x6f, 0x8b, 0x69, 0x2d, 0x65, 0x12, 0x06, 0x82, 0x51, 0x09, 0xfe, 0x91, 0x7a, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x34, 0x9c, 0xe8, 0xbb, 0xb2, 0x08, 0x00, 0x00,
}
