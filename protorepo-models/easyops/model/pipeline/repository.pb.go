// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: repository.proto

package pipeline

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//仓库 https://developer.github.com/v3/repos/ https://docs.gitlab.com/ee/api/projects.html
type Repository struct {
	//
	//git项目名称
	RepoName string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name" form:"repo_name"`
	//
	//git项目id
	RepoId int32 `protobuf:"varint,2,opt,name=repo_id,json=repoId,proto3" json:"repo_id" form:"repo_id"`
	//
	//带namespace的显示名称
	NameWithNamespace string `protobuf:"bytes,3,opt,name=name_with_namespace,json=nameWithNamespace,proto3" json:"name_with_namespace" form:"name_with_namespace"`
	//
	//带namespace的项目路径
	PathWithNamespace string `protobuf:"bytes,4,opt,name=path_with_namespace,json=pathWithNamespace,proto3" json:"path_with_namespace" form:"path_with_namespace"`
	//
	//git ssh 克隆地址
	GitSshUrl string `protobuf:"bytes,5,opt,name=git_ssh_url,json=gitSshUrl,proto3" json:"git_ssh_url" form:"git_ssh_url"`
	//
	//git http克隆地址
	GitHttpUrl string `protobuf:"bytes,6,opt,name=git_http_url,json=gitHttpUrl,proto3" json:"git_http_url" form:"git_http_url"`
	//
	//repo web url
	Link string `protobuf:"bytes,7,opt,name=link,proto3" json:"link" form:"link"`
	//
	//默认分支名称
	DefaultBranch        string   `protobuf:"bytes,8,opt,name=default_branch,json=defaultBranch,proto3" json:"default_branch" form:"default_branch"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{0}
}
func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *Repository) GetRepoId() int32 {
	if m != nil {
		return m.RepoId
	}
	return 0
}

func (m *Repository) GetNameWithNamespace() string {
	if m != nil {
		return m.NameWithNamespace
	}
	return ""
}

func (m *Repository) GetPathWithNamespace() string {
	if m != nil {
		return m.PathWithNamespace
	}
	return ""
}

func (m *Repository) GetGitSshUrl() string {
	if m != nil {
		return m.GitSshUrl
	}
	return ""
}

func (m *Repository) GetGitHttpUrl() string {
	if m != nil {
		return m.GitHttpUrl
	}
	return ""
}

func (m *Repository) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Repository) GetDefaultBranch() string {
	if m != nil {
		return m.DefaultBranch
	}
	return ""
}

func init() {
	proto.RegisterType((*Repository)(nil), "pipeline.Repository")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor_10d86afa5a89ec9d) }

var fileDescriptor_10d86afa5a89ec9d = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xb5, 0xb4, 0xdd, 0xee, 0xba, 0x50, 0x5a, 0x97, 0x3f, 0xa1, 0x07, 0x52, 0x99, 0x4b,
	0x25, 0xd4, 0x8d, 0x10, 0x12, 0x12, 0x9c, 0x50, 0xc4, 0x01, 0x2e, 0x7b, 0x30, 0x42, 0x48, 0x5c,
	0x22, 0x27, 0xf1, 0xc6, 0x16, 0x4e, 0x6c, 0xd9, 0x13, 0xa1, 0xbe, 0x6c, 0x8e, 0x3c, 0x40, 0x9e,
	0x00, 0x79, 0xd2, 0x2d, 0x6d, 0xb5, 0xb7, 0x99, 0xf9, 0xbe, 0xdf, 0x2f, 0x87, 0x98, 0x9c, 0x78,
	0xe9, 0x6c, 0xd0, 0x60, 0xfd, 0xf5, 0xca, 0x79, 0x0b, 0x96, 0x2e, 0x9c, 0x76, 0xd2, 0xe8, 0x4e,
	0x9e, 0x5f, 0x35, 0x1a, 0x54, 0x5f, 0xae, 0x2a, 0xdb, 0x66, 0x8d, 0x6d, 0x6c, 0x86, 0x85, 0xb2,
	0xdf, 0xe0, 0x86, 0x0b, 0x4e, 0x13, 0xc8, 0xfe, 0xee, 0x11, 0xc2, 0x6f, 0x6d, 0xf4, 0x1d, 0x59,
	0x46, 0x77, 0xd1, 0x89, 0x56, 0x26, 0xb3, 0x8b, 0xd9, 0xe5, 0x32, 0x7f, 0x36, 0x0e, 0xe9, 0xc9,
	0xc6, 0xfa, 0xf6, 0x13, 0xbb, 0x8d, 0x18, 0x5f, 0xc4, 0x79, 0x2d, 0x5a, 0x49, 0xdf, 0x92, 0x43,
	0xbc, 0xeb, 0x3a, 0x79, 0x74, 0x31, 0xbb, 0x3c, 0xc8, 0xe9, 0x38, 0xa4, 0xc7, 0x77, 0x00, 0x5d,
	0x33, 0x3e, 0x8f, 0xd3, 0xb7, 0x9a, 0xae, 0xc9, 0x59, 0xe4, 0x8b, 0x3f, 0x1a, 0x14, 0x9a, 0x82,
	0x13, 0x95, 0x4c, 0xf6, 0xf0, 0x4b, 0xaf, 0xc7, 0x21, 0x3d, 0x9f, 0xc0, 0x1d, 0x25, 0xc6, 0x4f,
	0xe3, 0xfc, 0x53, 0x83, 0x5a, 0x6f, 0x6f, 0xd1, 0xe7, 0x04, 0xa8, 0x87, 0xbe, 0xfd, 0x87, 0xbe,
	0x1d, 0x25, 0xc6, 0x4f, 0xe3, 0xf5, 0xbe, 0xef, 0x03, 0x39, 0x6a, 0x34, 0x14, 0x21, 0xa8, 0xa2,
	0xf7, 0x26, 0x39, 0x40, 0xcf, 0x8b, 0x71, 0x48, 0xe9, 0xe4, 0xb9, 0x13, 0x32, 0xbe, 0x6c, 0x34,
	0x7c, 0x0f, 0xea, 0x87, 0x37, 0xf4, 0x23, 0x79, 0x1c, 0x23, 0x05, 0xe0, 0x10, 0x9c, 0x23, 0xf8,
	0x72, 0x1c, 0xd2, 0xb3, 0xff, 0xe0, 0x36, 0x65, 0x9c, 0x34, 0x1a, 0xbe, 0x02, 0xb8, 0x88, 0xbe,
	0x21, 0xfb, 0x46, 0x77, 0xbf, 0x93, 0x43, 0x44, 0x9e, 0x8e, 0x43, 0x7a, 0x34, 0x21, 0xf1, 0xca,
	0x38, 0x86, 0xf4, 0x33, 0x39, 0xae, 0xe5, 0x46, 0xf4, 0x06, 0x8a, 0xd2, 0x8b, 0xae, 0x52, 0xc9,
	0x02, 0xeb, 0xaf, 0xc6, 0x21, 0x7d, 0x3e, 0xd5, 0xef, 0xe7, 0x8c, 0x3f, 0xb9, 0x39, 0xe4, 0xb8,
	0xe7, 0x5f, 0x7e, 0xe5, 0x8d, 0x5d, 0x49, 0x11, 0xae, 0xad, 0x0b, 0x2b, 0x63, 0x2b, 0x61, 0xb2,
	0xca, 0x76, 0xe0, 0x45, 0x05, 0x61, 0x7a, 0x24, 0xf1, 0x1f, 0x5d, 0xb5, 0xb6, 0x96, 0x26, 0x64,
	0x37, 0xc5, 0x0c, 0xd7, 0x6c, 0xfb, 0xba, 0xca, 0x39, 0x16, 0xdf, 0xff, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xc5, 0xeb, 0xa3, 0x82, 0x02, 0x00, 0x00,
}