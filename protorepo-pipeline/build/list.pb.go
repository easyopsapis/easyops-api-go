// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//List请求
type ListRequest struct {
	//
	//页码
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	//
	//每页大小
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//条件查询；需要构造这个query格式；具体格式如下：\ 1.按照日期筛选：\ {"page":"1","page_size":"100", "query":{"created":{"$gt":156024989 }}\ 2.按照创建人筛选：\ {"page":"1","page_size":"100","query":{"sender":"em"} \ 3.按照状态筛选：\ {"page":"1","page_size":"100","query":{"status.state":"failed" }} \ 4.按照pipeline_id筛选 \ {"page": 1,"page_size": 10,"query": {"PIPELINE.instanceId": "58af5330e04d"}} \ 5.按照project_id筛选 \ {"page": 1,"page_size": 10,"query": {"PIPELINE.PROJECT.instanceId": "58a8dd6a23d6e"}}
	Query *types.Struct `protobuf:"bytes,3,opt,name=query,proto3" json:"query" form:"query"`
	//
	//条件查询：按照字段正/倒序 :1表示升序, -1表示降序\ {"page": 1,"page_size": 10,"sort": {"created": -1}} //created可替换当前模型任意字段
	Sort                 *types.Struct `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort" form:"sort"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRequest) GetQuery() *types.Struct {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ListRequest) GetSort() *types.Struct {
	if m != nil {
		return m.Sort
	}
	return nil
}

//
//List返回
type ListResponse struct {
	//
	//返回总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total" form:"total"`
	//
	//页数
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page" form:"page"`
	//
	//该页大小
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	//
	//列表
	List                 []*ListResponse_List `protobuf:"bytes,4,rep,name=list,proto3" json:"list" form:"list"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResponse) GetList() []*ListResponse_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListResponse_List struct {
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
	Artifact *ListResponse_List_Artifact `protobuf:"bytes,7,opt,name=artifact,proto3" json:"artifact" form:"artifact"`
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

func (m *ListResponse_List) Reset()         { *m = ListResponse_List{} }
func (m *ListResponse_List) String() string { return proto.CompactTextString(m) }
func (*ListResponse_List) ProtoMessage()    {}
func (*ListResponse_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0}
}
func (m *ListResponse_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse_List.Unmarshal(m, b)
}
func (m *ListResponse_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse_List.Marshal(b, m, deterministic)
}
func (m *ListResponse_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_List.Merge(m, src)
}
func (m *ListResponse_List) XXX_Size() int {
	return xxx_messageInfo_ListResponse_List.Size(m)
}
func (m *ListResponse_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse_List proto.InternalMessageInfo

func (m *ListResponse_List) GetProject() *pipeline.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ListResponse_List) GetPipeline() *pipeline.Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ListResponse_List) GetStages() []*pipeline.StageStatus {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *ListResponse_List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListResponse_List) GetGitMeta() *pipeline.GitMeta {
	if m != nil {
		return m.GitMeta
	}
	return nil
}

func (m *ListResponse_List) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ListResponse_List) GetArtifact() *ListResponse_List_Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func (m *ListResponse_List) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ListResponse_List) GetYamlString() string {
	if m != nil {
		return m.YamlString
	}
	return ""
}

func (m *ListResponse_List) GetStatus() *pipeline.BuildStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse_List) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *ListResponse_List) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

type ListResponse_List_Artifact struct {
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

func (m *ListResponse_List_Artifact) Reset()         { *m = ListResponse_List_Artifact{} }
func (m *ListResponse_List_Artifact) String() string { return proto.CompactTextString(m) }
func (*ListResponse_List_Artifact) ProtoMessage()    {}
func (*ListResponse_List_Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{1, 0, 0}
}
func (m *ListResponse_List_Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse_List_Artifact.Unmarshal(m, b)
}
func (m *ListResponse_List_Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse_List_Artifact.Marshal(b, m, deterministic)
}
func (m *ListResponse_List_Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_List_Artifact.Merge(m, src)
}
func (m *ListResponse_List_Artifact) XXX_Size() int {
	return xxx_messageInfo_ListResponse_List_Artifact.Size(m)
}
func (m *ListResponse_List_Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse_List_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse_List_Artifact proto.InternalMessageInfo

func (m *ListResponse_List_Artifact) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ListResponse_List_Artifact) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *ListResponse_List_Artifact) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *ListResponse_List_Artifact) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *ListResponse_List_Artifact) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

//
//ListApi返回
type ListResponseWrapper struct {
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
	Data                 *ListResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data" form:"data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponseWrapper) Reset()         { *m = ListResponseWrapper{} }
func (m *ListResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ListResponseWrapper) ProtoMessage()    {}
func (*ListResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_af793ce248ee1bf0, []int{2}
}
func (m *ListResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponseWrapper.Unmarshal(m, b)
}
func (m *ListResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponseWrapper.Marshal(b, m, deterministic)
}
func (m *ListResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponseWrapper.Merge(m, src)
}
func (m *ListResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ListResponseWrapper.Size(m)
}
func (m *ListResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponseWrapper proto.InternalMessageInfo

func (m *ListResponseWrapper) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListResponseWrapper) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *ListResponseWrapper) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ListResponseWrapper) GetData() *ListResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "build.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "build.ListResponse")
	proto.RegisterType((*ListResponse_List)(nil), "build.ListResponse.List")
	proto.RegisterType((*ListResponse_List_Artifact)(nil), "build.ListResponse.List.Artifact")
	proto.RegisterType((*ListResponseWrapper)(nil), "build.ListResponseWrapper")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_af793ce248ee1bf0) }

var fileDescriptor_af793ce248ee1bf0 = []byte{
	// 1080 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xed, 0x6e, 0x23, 0x35,
	0x17, 0x7e, 0x93, 0x26, 0x6d, 0xea, 0xf4, 0xdd, 0xed, 0xba, 0x5a, 0x18, 0x45, 0x48, 0x69, 0xbd,
	0x15, 0x9a, 0xb0, 0x3b, 0x93, 0x8f, 0x2e, 0x6c, 0xb7, 0x02, 0x85, 0x06, 0x01, 0x5a, 0x09, 0x50,
	0xd7, 0x01, 0x21, 0x6d, 0x48, 0x2b, 0x37, 0x71, 0x87, 0xa1, 0x49, 0x3c, 0x3b, 0x76, 0xba, 0xf4,
	0xeb, 0x02, 0xb8, 0x15, 0x24, 0x6e, 0x80, 0x9f, 0x5c, 0x04, 0x12, 0x7f, 0x82, 0xc4, 0x15, 0xa0,
	0x5c, 0x01, 0xf2, 0x19, 0x4f, 0xe2, 0x2e, 0x05, 0xa9, 0x52, 0xf9, 0x35, 0x3e, 0xe7, 0x3c, 0x8f,
	0xfd, 0x1c, 0x9f, 0x73, 0x3c, 0x08, 0x0d, 0x42, 0xa9, 0xfc, 0x28, 0x16, 0x4a, 0xe0, 0xfc, 0xe1,
	0x38, 0x1c, 0xf4, 0x4b, 0x5e, 0x10, 0xaa, 0x6f, 0xc7, 0x87, 0x7e, 0x4f, 0x0c, 0xab, 0x81, 0x08,
	0x44, 0x15, 0xa2, 0x87, 0xe3, 0x23, 0xb0, 0xc0, 0x80, 0x55, 0xc2, 0x2a, 0xed, 0x05, 0xc2, 0xe7,
	0x4c, 0x9e, 0x8a, 0x48, 0xfa, 0x03, 0xd1, 0x63, 0x83, 0x6a, 0x4f, 0x8c, 0x54, 0xcc, 0x7a, 0x4a,
	0x26, 0xcc, 0x98, 0x47, 0xc2, 0x1b, 0x8a, 0x3e, 0x1f, 0xc8, 0xaa, 0x01, 0x56, 0xc1, 0xac, 0x46,
	0x61, 0xc4, 0x07, 0xe1, 0x88, 0x6b, 0xdc, 0x77, 0xbc, 0x67, 0x74, 0x94, 0x9e, 0xdf, 0xc6, 0x8e,
	0x66, 0x61, 0xb6, 0xfc, 0xea, 0x16, 0xb6, 0x94, 0x8a, 0x05, 0xfc, 0x40, 0x2a, 0xa6, 0xc6, 0xf2,
	0x16, 0x95, 0x06, 0xa1, 0x3a, 0x18, 0x72, 0xc5, 0x6e, 0x51, 0x29, 0xd4, 0xf1, 0xaa, 0xd2, 0xf7,
	0xac, 0xa2, 0x0e, 0x5f, 0x85, 0xea, 0x58, 0xbc, 0xaa, 0x06, 0xc2, 0x83, 0xa0, 0x77, 0xc2, 0x06,
	0x61, 0x9f, 0x29, 0x11, 0xcb, 0xea, 0x6c, 0x69, 0x78, 0x6f, 0x05, 0x42, 0x04, 0x03, 0x3e, 0xef,
	0x01, 0xa9, 0xe2, 0x71, 0x5a, 0x29, 0xf2, 0x67, 0x06, 0x15, 0x3f, 0x0b, 0xa5, 0xa2, 0xfc, 0xe5,
	0x98, 0x4b, 0x85, 0x2b, 0x28, 0x17, 0xb1, 0x80, 0x3b, 0x99, 0xf5, 0x8c, 0x9b, 0x6f, 0xdd, 0x9f,
	0x4e, 0xca, 0xc5, 0x23, 0x11, 0x0f, 0x77, 0x88, 0xf6, 0x92, 0x3f, 0x7e, 0x2f, 0x67, 0x57, 0xff,
	0x47, 0x01, 0x82, 0x9f, 0xa0, 0xe5, 0x08, 0xee, 0x33, 0x3c, 0xe3, 0x4e, 0x16, 0xf0, 0xa5, 0xe9,
	0xa4, 0xbc, 0x3a, 0xc7, 0x43, 0x28, 0x25, 0x15, 0xb4, 0xa7, 0x1d, 0x9e, 0x71, 0xdc, 0x44, 0xf9,
	0x97, 0x63, 0x1e, 0x9f, 0x3a, 0x0b, 0xeb, 0x19, 0xb7, 0xd8, 0x78, 0xd3, 0x4f, 0x14, 0xfa, 0xa9,
	0x42, 0xbf, 0x0d, 0x0a, 0x5b, 0xab, 0xd3, 0x49, 0x79, 0x25, 0xd9, 0x0d, 0xf0, 0x84, 0x26, 0x3c,
	0xfc, 0x3e, 0xca, 0x49, 0x11, 0x2b, 0x27, 0xf7, 0xef, 0xfc, 0xbb, 0x73, 0xf5, 0x1a, 0x4e, 0x28,
	0xb0, 0xc8, 0xcf, 0x2b, 0x68, 0x25, 0x49, 0x59, 0x46, 0x62, 0x24, 0x39, 0x7e, 0x1b, 0xe5, 0x95,
	0x50, 0x6c, 0x60, 0x92, 0xb6, 0x8e, 0x05, 0x37, 0xa1, 0x49, 0x78, 0x76, 0x37, 0xd9, 0x1b, 0xde,
	0xcd, 0xc2, 0x0d, 0xee, 0xe6, 0x03, 0x94, 0xd3, 0xf3, 0xec, 0xe4, 0xd6, 0x17, 0xdc, 0x62, 0xc3,
	0xf1, 0xa1, 0x11, 0x7c, 0x5b, 0x2e, 0x18, 0x76, 0x6e, 0x1a, 0x4f, 0x28, 0xd0, 0x4a, 0x13, 0x84,
	0x72, 0x3a, 0x8e, 0x9b, 0x68, 0xc9, 0x8c, 0x24, 0x64, 0x55, 0x6c, 0xdc, 0xf3, 0x67, 0x03, 0xb5,
	0x97, 0x04, 0x5a, 0x78, 0x3a, 0x29, 0xdf, 0x31, 0x8a, 0x12, 0x17, 0xa1, 0x29, 0x0b, 0x7f, 0x84,
	0x0a, 0x29, 0x01, 0x12, 0x2e, 0x36, 0xb0, 0xb5, 0x83, 0x59, 0xb4, 0xd6, 0xa6, 0x93, 0xf2, 0x5d,
	0xb3, 0x85, 0xf1, 0x11, 0x3a, 0x23, 0xe2, 0x0f, 0xd1, 0x22, 0xcc, 0x9c, 0x74, 0x16, 0x20, 0x9f,
	0xfb, 0xf3, 0x2d, 0xda, 0xda, 0xdf, 0x86, 0x06, 0x6f, 0xdd, 0x9b, 0x4e, 0xca, 0xff, 0x37, 0x85,
	0x02, 0x38, 0xa1, 0x86, 0x87, 0xeb, 0x28, 0x1b, 0xf6, 0xa1, 0xd0, 0xcb, 0xad, 0x8d, 0xe9, 0xa4,
	0xbc, 0x9c, 0xc0, 0xc2, 0xbe, 0xbe, 0xba, 0x55, 0x74, 0x67, 0xbf, 0x53, 0xf3, 0x9e, 0x32, 0xef,
	0xac, 0x7b, 0x5e, 0xdf, 0xba, 0xdc, 0xa4, 0xd9, 0xb0, 0x8f, 0x77, 0x51, 0x21, 0x9d, 0x48, 0x27,
	0xff, 0x7a, 0xee, 0x9f, 0x86, 0xea, 0x73, 0xae, 0x98, 0x2d, 0x3c, 0x05, 0x13, 0xba, 0x14, 0x24,
	0x51, 0xbc, 0x87, 0x16, 0x25, 0x1f, 0xf5, 0x79, 0xec, 0x2c, 0xc2, 0xc9, 0xdb, 0x96, 0x40, 0xf0,
	0xeb, 0xd3, 0x1f, 0xa0, 0x8d, 0xfd, 0x0e, 0xf3, 0xce, 0x76, 0xbd, 0x17, 0x35, 0xef, 0x69, 0xb7,
	0xe3, 0xcf, 0xd6, 0x07, 0x5e, 0xf7, 0xbc, 0xf1, 0x68, 0xab, 0x7e, 0xb9, 0x49, 0xcd, 0x3e, 0x98,
	0xa2, 0x02, 0x8b, 0x55, 0x78, 0xc4, 0x7a, 0xca, 0x59, 0x02, 0x51, 0x1b, 0xff, 0x54, 0x5b, 0x7f,
	0xd7, 0x00, 0x6d, 0x91, 0x29, 0x99, 0xd0, 0xd9, 0x3e, 0xf8, 0x11, 0x5a, 0xea, 0xc5, 0x9c, 0x29,
	0xde, 0x77, 0x0a, 0xd0, 0x62, 0x56, 0x41, 0x4d, 0x80, 0xd0, 0x14, 0x82, 0x9f, 0xa0, 0xe2, 0x29,
	0x1b, 0x0e, 0x0e, 0xa4, 0x8a, 0xc3, 0x51, 0xe0, 0x2c, 0x43, 0x62, 0x6f, 0x4c, 0x27, 0x65, 0x9c,
	0x30, 0xac, 0x20, 0xa1, 0x48, 0x5b, 0x6d, 0x30, 0x4c, 0x11, 0xd5, 0x58, 0x3a, 0x08, 0x84, 0x5b,
	0x45, 0x6c, 0xe9, 0x0c, 0xae, 0x2d, 0xa2, 0x1a, 0x9b, 0x22, 0xaa, 0xb1, 0xc4, 0x15, 0xb4, 0x38,
	0x1a, 0x0f, 0x0f, 0x79, 0xec, 0x14, 0xe1, 0x54, 0x0b, 0x9a, 0xf8, 0x09, 0x35, 0x00, 0x0d, 0xe5,
	0x27, 0x7c, 0xa4, 0xa4, 0xb3, 0xb2, 0xbe, 0x70, 0x15, 0x9a, 0xf8, 0x09, 0x35, 0x80, 0xd2, 0x34,
	0x87, 0x0a, 0xe9, 0x55, 0xe1, 0x6d, 0x54, 0x8c, 0x58, 0xef, 0x98, 0x05, 0xfc, 0x0b, 0x36, 0x4c,
	0x9e, 0xaf, 0x2b, 0xd9, 0x59, 0x41, 0x42, 0x6d, 0xa8, 0x66, 0x9e, 0xf0, 0x58, 0x86, 0x62, 0x04,
	0xcc, 0xec, 0xeb, 0x4c, 0x2b, 0x48, 0xa8, 0x0d, 0xc5, 0xbf, 0x65, 0x50, 0xbe, 0xa7, 0xc2, 0x61,
	0x32, 0xe1, 0xcb, 0xad, 0x5f, 0x32, 0xf3, 0x97, 0x03, 0xfc, 0xba, 0x4b, 0x7e, 0xca, 0xa0, 0x1f,
	0x33, 0xfb, 0xae, 0xdb, 0xdc, 0xe9, 0xd4, 0x75, 0x97, 0xe8, 0x56, 0x79, 0xa7, 0xd2, 0x84, 0xef,
	0xf9, 0xe3, 0xcb, 0x8a, 0xe7, 0xd6, 0x3b, 0x35, 0xaf, 0xd1, 0xbd, 0xa8, 0x41, 0xbc, 0xe2, 0xb9,
	0x5b, 0x9d, 0x9a, 0x57, 0x4f, 0xed, 0x8b, 0x4e, 0xdd, 0x6b, 0x24, 0xac, 0x4a, 0xe7, 0xcb, 0xf5,
	0xae, 0xdb, 0xe8, 0xd4, 0xbc, 0xad, 0xee, 0x05, 0x60, 0x12, 0xf7, 0x8e, 0xdb, 0xa9, 0x79, 0xef,
	0xa6, 0xc6, 0x7c, 0xed, 0x7e, 0xe3, 0xc3, 0xf7, 0x61, 0xa5, 0xe9, 0xbe, 0xb8, 0xe8, 0x3c, 0xf4,
	0xba, 0x6e, 0x73, 0xe7, 0x1a, 0xba, 0xc5, 0x6e, 0x6e, 0xd2, 0x24, 0x23, 0xfc, 0x43, 0x46, 0xbf,
	0x60, 0x70, 0x4b, 0xcf, 0xd2, 0xf9, 0x3b, 0xb6, 0x5f, 0x30, 0x13, 0xd2, 0x29, 0xb6, 0xd1, 0x73,
	0x3d, 0x08, 0x47, 0xbb, 0xde, 0x27, 0x90, 0xd5, 0xf6, 0xa5, 0xd7, 0xb4, 0xed, 0xc7, 0x37, 0xb4,
	0xeb, 0x8d, 0xcb, 0x4d, 0x3a, 0x3f, 0x1d, 0xb4, 0x98, 0x7b, 0x7f, 0xd6, 0x87, 0x91, 0xbe, 0xa2,
	0x65, 0x16, 0xfa, 0xef, 0xb4, 0xcc, 0x8f, 0xf8, 0x35, 0x83, 0xd6, 0xec, 0x89, 0xfd, 0x3a, 0x66,
	0x51, 0xc4, 0x63, 0xfc, 0x00, 0xe5, 0x7a, 0xa2, 0x9f, 0xfe, 0x37, 0xad, 0xd7, 0x59, 0x7b, 0x09,
	0x85, 0xa0, 0x6e, 0x35, 0xfd, 0xfd, 0xf8, 0xfb, 0x68, 0xc0, 0xc2, 0xd1, 0xdf, 0x5b, 0xcd, 0x0a,
	0x12, 0x6a, 0x43, 0xf5, 0x2f, 0x8a, 0xc7, 0xb1, 0x88, 0x4d, 0xa7, 0x59, 0xbf, 0x28, 0x70, 0x13,
	0x9a, 0x84, 0xf1, 0x36, 0xca, 0xf5, 0x99, 0x62, 0xe6, 0xcf, 0xb8, 0x76, 0xcd, 0x13, 0x63, 0x6b,
	0xd3, 0x50, 0x42, 0x81, 0x71, 0xb8, 0x08, 0x7f, 0xcf, 0xad, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0xa7, 0x0a, 0x33, 0x4f, 0x0a, 0x00, 0x00,
}
