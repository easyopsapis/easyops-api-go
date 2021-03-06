// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool.proto

package tool

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
//工具
type Tool struct {
	//
	//工具Id，服务端自动生成
	ToolId string `protobuf:"bytes,1,opt,name=toolId,proto3" json:"toolId" form:"toolId"`
	//
	//工具版本
	VId string `protobuf:"bytes,2,opt,name=vId,proto3" json:"vId" form:"vId"`
	//
	//工具名字
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" form:"name"`
	//
	//工具执行的超时时间(0表示不超时，默认为86400秒)
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout" form:"timeout"`
	//
	//工具的分类
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category" form:"category"`
	//
	//工具的页面显示图标
	Icon string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon" form:"icon"`
	//
	//图标风格
	Style string `protobuf:"bytes,7,opt,name=style,proto3" json:"style" form:"style"`
	//
	//工具说明
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo" form:"memo"`
	//
	//列表中是否显示
	ListVisible bool `protobuf:"varint,9,opt,name=listVisible,proto3" json:"listVisible" form:"listVisible"`
	//
	//标签
	Tags []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags" form:"tags"`
	//
	//禁用工具开关
	Disable bool `protobuf:"varint,11,opt,name=disable,proto3" json:"disable" form:"disable"`
	//
	//工具创建时间
	CreateTime string `protobuf:"bytes,12,opt,name=createTime,proto3" json:"createTime" form:"createTime"`
	//
	//工具更新时间
	UpdateTime string `protobuf:"bytes,13,opt,name=updateTime,proto3" json:"updateTime" form:"updateTime"`
	//
	//工具创建用户名
	Creator string `protobuf:"bytes,14,opt,name=creator,proto3" json:"creator" form:"creator"`
	//
	//客户id
	Org int32 `protobuf:"varint,15,opt,name=org,proto3" json:"org" form:"org"`
	//
	//工具输入参数信息
	Inputs []*ToolInput `protobuf:"bytes,16,rep,name=inputs,proto3" json:"inputs" form:"inputs"`
	//
	//工具输出参数
	Outputs string `protobuf:"bytes,17,opt,name=outputs,proto3" json:"outputs" form:"outputs"`
	//
	//工具脚本的类型(python, shell..)
	Type string `protobuf:"bytes,18,opt,name=type,proto3" json:"type" form:"type"`
	//
	//工具脚本的内容
	Content string `protobuf:"bytes,19,opt,name=content,proto3" json:"content" form:"content"`
	//
	//默认执行用户
	DefaultExecUser string `protobuf:"bytes,20,opt,name=defaultExecUser,proto3" json:"defaultExecUser" form:"defaultExecUser"`
	//
	//执行用户
	ExecUser string `protobuf:"bytes,21,opt,name=execUser,proto3" json:"execUser" form:"execUser"`
	//
	//授权执行用户列表
	AuthUsers []string `protobuf:"bytes,22,rep,name=authUsers,proto3" json:"authUsers" form:"authUsers"`
	//
	//默认执行的agents
	DefaultAgents []string `protobuf:"bytes,23,rep,name=defaultAgents,proto3" json:"defaultAgents" form:"defaultAgents"`
	//
	//系统插件
	SystemPlugin bool `protobuf:"varint,24,opt,name=system_plugin,json=systemPlugin,proto3" json:"system_plugin" form:"system_plugin"`
	//
	//系统工具
	SystemTool bool `protobuf:"varint,25,opt,name=system_tool,json=systemTool,proto3" json:"system_tool" form:"system_tool"`
	//
	//绑定的agents
	LockAgents string `protobuf:"bytes,26,opt,name=lockAgents,proto3" json:"lockAgents" form:"lockAgents"`
	//
	//沙箱执行；默认执行的agents
	SandboxRun bool `protobuf:"varint,27,opt,name=sandboxRun,proto3" json:"sandboxRun" form:"sandboxRun"`
	//
	//白名单
	WhiteList []string `protobuf:"bytes,28,rep,name=whiteList,proto3" json:"whiteList" form:"whiteList"`
	//
	//是否使用windows会话用户
	WindowsSession bool `protobuf:"varint,29,opt,name=windowsSession,proto3" json:"windowsSession" form:"windowsSession"`
	//
	//黑名单
	BlackList []string `protobuf:"bytes,30,rep,name=blackList,proto3" json:"blackList" form:"blackList"`
	//
	//工具的输出, 输出为空为[],非空时为object, 结构如下:
	//fields:
	//- name: dimensions
	//type: object[]
	//description: 维度列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//- name: columns
	//type: object[]
	//description: 输出列
	//fields:
	//- name: id
	//type: string
	//description: 列id
	//- name: name
	//type: string
	//description: 列标题
	//- name: type
	//type: string
	//description: 列类型
	//
	ToolOutputs *types.Value `protobuf:"bytes,31,opt,name=toolOutputs,proto3" json:"toolOutputs" form:"toolOutputs"`
	//
	//工具的输出定义
	OutputDefs []*Tool_OutputDefs `protobuf:"bytes,32,rep,name=outputDefs,proto3" json:"outputDefs" form:"outputDefs"`
	//
	//工具的表格输出定义
	TableDefs []*Tool_TableDefs `protobuf:"bytes,33,rep,name=tableDefs,proto3" json:"tableDefs" form:"tableDefs"`
	//
	//版本创建时间
	VCreateTime string `protobuf:"bytes,34,opt,name=vCreateTime,proto3" json:"vCreateTime" form:"vCreateTime"`
	//
	//版本更新时间
	VUpdateTime string `protobuf:"bytes,35,opt,name=vUpdateTime,proto3" json:"vUpdateTime" form:"vUpdateTime"`
	//
	//版本名称
	VName string `protobuf:"bytes,36,opt,name=vName,proto3" json:"vName" form:"vName"`
	//
	//版本作者
	VCreator string `protobuf:"bytes,37,opt,name=vCreator,proto3" json:"vCreator" form:"vCreator"`
	//
	//查看详情授权用户
	ReadAuthorizers []string `protobuf:"bytes,38,rep,name=readAuthorizers,proto3" json:"readAuthorizers" form:"readAuthorizers"`
	//
	//更新授权用户
	UpdateAuthorizers []string `protobuf:"bytes,39,rep,name=updateAuthorizers,proto3" json:"updateAuthorizers" form:"updateAuthorizers"`
	//
	//删除授权用户
	DeleteAuthorizers []string `protobuf:"bytes,40,rep,name=deleteAuthorizers,proto3" json:"deleteAuthorizers" form:"deleteAuthorizers"`
	//
	//执行授权用户
	ExecuteAuthorizers []string `protobuf:"bytes,41,rep,name=executeAuthorizers,proto3" json:"executeAuthorizers" form:"executeAuthorizers"`
	//
	//执行结果查看授权用户
	ReadExecutionResultAuthorizers []string `protobuf:"bytes,42,rep,name=readExecutionResultAuthorizers,proto3" json:"readExecutionResultAuthorizers" form:"readExecutionResultAuthorizers"`
	//
	//root执行授权用户
	RootExecuteAuthorizers []string `protobuf:"bytes,43,rep,name=rootExecuteAuthorizers,proto3" json:"rootExecuteAuthorizers" form:"rootExecuteAuthorizers"`
	//
	//版本说明
	VDesc string `protobuf:"bytes,44,opt,name=vDesc,proto3" json:"vDesc" form:"vDesc"`
	//
	//工具来源
	SourceFrom string `protobuf:"bytes,45,opt,name=sourceFrom,proto3" json:"sourceFrom" form:"sourceFrom"`
	//
	//版本类型
	EnvType string `protobuf:"bytes,46,opt,name=envType,proto3" json:"envType" form:"envType"`
	//
	//工具审批结果(approving=待审批;approved=已审批;rejected=审批未通过)
	CheckType string `protobuf:"bytes,47,opt,name=checkType,proto3" json:"checkType" form:"checkType"`
	//
	//审批人
	CheckUser string `protobuf:"bytes,48,opt,name=checkUser,proto3" json:"checkUser" form:"checkUser"`
	//
	//审批发起人
	CheckSponsor string `protobuf:"bytes,49,opt,name=checkSponsor,proto3" json:"checkSponsor" form:"checkSponsor"`
	//
	//工具审批结论意见
	CheckComment string `protobuf:"bytes,50,opt,name=checkComment,proto3" json:"checkComment" form:"checkComment"`
	//
	//当为空时是[],
	//非空时为object如,分批策略{"batchNum":10,"batchInterval":3,"failedStop":true,"enabled":true}
	//
	BatchStrategy *types.Value `protobuf:"bytes,51,opt,name=batchStrategy,proto3" json:"batchStrategy" form:"batchStrategy"`
	//
	//初始化后的工具为 SQL  HTTP  JAVA   normal(null)
	FunctionType string `protobuf:"bytes,52,opt,name=functionType,proto3" json:"functionType" form:"functionType"`
	//
	//系统ORG
	IsSystemOrg bool `protobuf:"varint,53,opt,name=is_system_org,json=isSystemOrg,proto3" json:"is_system_org" form:"is_system_org"`
	//
	//邮件订阅用户
	Subscribers []string `protobuf:"bytes,54,rep,name=subscribers,proto3" json:"subscribers" form:"subscribers"`
	//
	//邮件通知方式 ALL | WARNING
	SubscribedChannel string `protobuf:"bytes,55,opt,name=subscribedChannel,proto3" json:"subscribedChannel" form:"subscribedChannel"`
	//
	//是否只读
	ReadOnly bool `protobuf:"varint,56,opt,name=readOnly,proto3" json:"readOnly" form:"readOnly"`
	//
	//模版类型
	TemplateType string `protobuf:"bytes,57,opt,name=templateType,proto3" json:"templateType" form:"templateType"`
	//
	//数据库删除标记
	DeleteMe bool `protobuf:"varint,58,opt,name=delete_me,json=deleteMe,proto3" json:"delete_me" form:"delete_me"`
	//
	//审批用户
	Approvers []string `protobuf:"bytes,59,rep,name=approvers,proto3" json:"approvers" form:"approvers"`
	//
	//Ansible Node机器执行用户
	AnsibleNodeExecUser string `protobuf:"bytes,60,opt,name=ansibleNodeExecUser,proto3" json:"ansibleNodeExecUser" form:"ansibleNodeExecUser"`
	//
	//ansible-playbook verbose mode
	Log                  *Tool_Log `protobuf:"bytes,61,opt,name=log,proto3" json:"log" form:"log"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Tool) Reset()         { *m = Tool{} }
func (m *Tool) String() string { return proto.CompactTextString(m) }
func (*Tool) ProtoMessage()    {}
func (*Tool) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0}
}
func (m *Tool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool.Unmarshal(m, b)
}
func (m *Tool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool.Marshal(b, m, deterministic)
}
func (m *Tool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool.Merge(m, src)
}
func (m *Tool) XXX_Size() int {
	return xxx_messageInfo_Tool.Size(m)
}
func (m *Tool) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool.DiscardUnknown(m)
}

var xxx_messageInfo_Tool proto.InternalMessageInfo

func (m *Tool) GetToolId() string {
	if m != nil {
		return m.ToolId
	}
	return ""
}

func (m *Tool) GetVId() string {
	if m != nil {
		return m.VId
	}
	return ""
}

func (m *Tool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tool) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Tool) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Tool) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Tool) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *Tool) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Tool) GetListVisible() bool {
	if m != nil {
		return m.ListVisible
	}
	return false
}

func (m *Tool) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Tool) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *Tool) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Tool) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Tool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Tool) GetOrg() int32 {
	if m != nil {
		return m.Org
	}
	return 0
}

func (m *Tool) GetInputs() []*ToolInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Tool) GetOutputs() string {
	if m != nil {
		return m.Outputs
	}
	return ""
}

func (m *Tool) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Tool) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Tool) GetDefaultExecUser() string {
	if m != nil {
		return m.DefaultExecUser
	}
	return ""
}

func (m *Tool) GetExecUser() string {
	if m != nil {
		return m.ExecUser
	}
	return ""
}

func (m *Tool) GetAuthUsers() []string {
	if m != nil {
		return m.AuthUsers
	}
	return nil
}

func (m *Tool) GetDefaultAgents() []string {
	if m != nil {
		return m.DefaultAgents
	}
	return nil
}

func (m *Tool) GetSystemPlugin() bool {
	if m != nil {
		return m.SystemPlugin
	}
	return false
}

func (m *Tool) GetSystemTool() bool {
	if m != nil {
		return m.SystemTool
	}
	return false
}

func (m *Tool) GetLockAgents() string {
	if m != nil {
		return m.LockAgents
	}
	return ""
}

func (m *Tool) GetSandboxRun() bool {
	if m != nil {
		return m.SandboxRun
	}
	return false
}

func (m *Tool) GetWhiteList() []string {
	if m != nil {
		return m.WhiteList
	}
	return nil
}

func (m *Tool) GetWindowsSession() bool {
	if m != nil {
		return m.WindowsSession
	}
	return false
}

func (m *Tool) GetBlackList() []string {
	if m != nil {
		return m.BlackList
	}
	return nil
}

func (m *Tool) GetToolOutputs() *types.Value {
	if m != nil {
		return m.ToolOutputs
	}
	return nil
}

func (m *Tool) GetOutputDefs() []*Tool_OutputDefs {
	if m != nil {
		return m.OutputDefs
	}
	return nil
}

func (m *Tool) GetTableDefs() []*Tool_TableDefs {
	if m != nil {
		return m.TableDefs
	}
	return nil
}

func (m *Tool) GetVCreateTime() string {
	if m != nil {
		return m.VCreateTime
	}
	return ""
}

func (m *Tool) GetVUpdateTime() string {
	if m != nil {
		return m.VUpdateTime
	}
	return ""
}

func (m *Tool) GetVName() string {
	if m != nil {
		return m.VName
	}
	return ""
}

func (m *Tool) GetVCreator() string {
	if m != nil {
		return m.VCreator
	}
	return ""
}

func (m *Tool) GetReadAuthorizers() []string {
	if m != nil {
		return m.ReadAuthorizers
	}
	return nil
}

func (m *Tool) GetUpdateAuthorizers() []string {
	if m != nil {
		return m.UpdateAuthorizers
	}
	return nil
}

func (m *Tool) GetDeleteAuthorizers() []string {
	if m != nil {
		return m.DeleteAuthorizers
	}
	return nil
}

func (m *Tool) GetExecuteAuthorizers() []string {
	if m != nil {
		return m.ExecuteAuthorizers
	}
	return nil
}

func (m *Tool) GetReadExecutionResultAuthorizers() []string {
	if m != nil {
		return m.ReadExecutionResultAuthorizers
	}
	return nil
}

func (m *Tool) GetRootExecuteAuthorizers() []string {
	if m != nil {
		return m.RootExecuteAuthorizers
	}
	return nil
}

func (m *Tool) GetVDesc() string {
	if m != nil {
		return m.VDesc
	}
	return ""
}

func (m *Tool) GetSourceFrom() string {
	if m != nil {
		return m.SourceFrom
	}
	return ""
}

func (m *Tool) GetEnvType() string {
	if m != nil {
		return m.EnvType
	}
	return ""
}

func (m *Tool) GetCheckType() string {
	if m != nil {
		return m.CheckType
	}
	return ""
}

func (m *Tool) GetCheckUser() string {
	if m != nil {
		return m.CheckUser
	}
	return ""
}

func (m *Tool) GetCheckSponsor() string {
	if m != nil {
		return m.CheckSponsor
	}
	return ""
}

func (m *Tool) GetCheckComment() string {
	if m != nil {
		return m.CheckComment
	}
	return ""
}

func (m *Tool) GetBatchStrategy() *types.Value {
	if m != nil {
		return m.BatchStrategy
	}
	return nil
}

func (m *Tool) GetFunctionType() string {
	if m != nil {
		return m.FunctionType
	}
	return ""
}

func (m *Tool) GetIsSystemOrg() bool {
	if m != nil {
		return m.IsSystemOrg
	}
	return false
}

func (m *Tool) GetSubscribers() []string {
	if m != nil {
		return m.Subscribers
	}
	return nil
}

func (m *Tool) GetSubscribedChannel() string {
	if m != nil {
		return m.SubscribedChannel
	}
	return ""
}

func (m *Tool) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *Tool) GetTemplateType() string {
	if m != nil {
		return m.TemplateType
	}
	return ""
}

func (m *Tool) GetDeleteMe() bool {
	if m != nil {
		return m.DeleteMe
	}
	return false
}

func (m *Tool) GetApprovers() []string {
	if m != nil {
		return m.Approvers
	}
	return nil
}

func (m *Tool) GetAnsibleNodeExecUser() string {
	if m != nil {
		return m.AnsibleNodeExecUser
	}
	return ""
}

func (m *Tool) GetLog() *Tool_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type Tool_OutputDefs struct {
	//
	//输出参数ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出参数标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tool_OutputDefs) Reset()         { *m = Tool_OutputDefs{} }
func (m *Tool_OutputDefs) String() string { return proto.CompactTextString(m) }
func (*Tool_OutputDefs) ProtoMessage()    {}
func (*Tool_OutputDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0, 0}
}
func (m *Tool_OutputDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool_OutputDefs.Unmarshal(m, b)
}
func (m *Tool_OutputDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool_OutputDefs.Marshal(b, m, deterministic)
}
func (m *Tool_OutputDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool_OutputDefs.Merge(m, src)
}
func (m *Tool_OutputDefs) XXX_Size() int {
	return xxx_messageInfo_Tool_OutputDefs.Size(m)
}
func (m *Tool_OutputDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool_OutputDefs.DiscardUnknown(m)
}

var xxx_messageInfo_Tool_OutputDefs proto.InternalMessageInfo

func (m *Tool_OutputDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tool_OutputDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tool_TableDefs struct {
	//
	//输出表格id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//输出表格标题
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	//
	//维度列
	Dimensions []*Tool_TableDefs_Dimensions `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions" form:"dimensions"`
	//
	//输出列
	Columns              []*Tool_TableDefs_Columns `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns" form:"columns"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tool_TableDefs) Reset()         { *m = Tool_TableDefs{} }
func (m *Tool_TableDefs) String() string { return proto.CompactTextString(m) }
func (*Tool_TableDefs) ProtoMessage()    {}
func (*Tool_TableDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0, 1}
}
func (m *Tool_TableDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool_TableDefs.Unmarshal(m, b)
}
func (m *Tool_TableDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool_TableDefs.Marshal(b, m, deterministic)
}
func (m *Tool_TableDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool_TableDefs.Merge(m, src)
}
func (m *Tool_TableDefs) XXX_Size() int {
	return xxx_messageInfo_Tool_TableDefs.Size(m)
}
func (m *Tool_TableDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool_TableDefs.DiscardUnknown(m)
}

var xxx_messageInfo_Tool_TableDefs proto.InternalMessageInfo

func (m *Tool_TableDefs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tool_TableDefs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tool_TableDefs) GetDimensions() []*Tool_TableDefs_Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Tool_TableDefs) GetColumns() []*Tool_TableDefs_Columns {
	if m != nil {
		return m.Columns
	}
	return nil
}

type Tool_TableDefs_Dimensions struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tool_TableDefs_Dimensions) Reset()         { *m = Tool_TableDefs_Dimensions{} }
func (m *Tool_TableDefs_Dimensions) String() string { return proto.CompactTextString(m) }
func (*Tool_TableDefs_Dimensions) ProtoMessage()    {}
func (*Tool_TableDefs_Dimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0, 1, 0}
}
func (m *Tool_TableDefs_Dimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool_TableDefs_Dimensions.Unmarshal(m, b)
}
func (m *Tool_TableDefs_Dimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool_TableDefs_Dimensions.Marshal(b, m, deterministic)
}
func (m *Tool_TableDefs_Dimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool_TableDefs_Dimensions.Merge(m, src)
}
func (m *Tool_TableDefs_Dimensions) XXX_Size() int {
	return xxx_messageInfo_Tool_TableDefs_Dimensions.Size(m)
}
func (m *Tool_TableDefs_Dimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool_TableDefs_Dimensions.DiscardUnknown(m)
}

var xxx_messageInfo_Tool_TableDefs_Dimensions proto.InternalMessageInfo

func (m *Tool_TableDefs_Dimensions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tool_TableDefs_Dimensions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tool_TableDefs_Columns struct {
	//
	//列id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id"`
	//
	//列标题
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tool_TableDefs_Columns) Reset()         { *m = Tool_TableDefs_Columns{} }
func (m *Tool_TableDefs_Columns) String() string { return proto.CompactTextString(m) }
func (*Tool_TableDefs_Columns) ProtoMessage()    {}
func (*Tool_TableDefs_Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0, 1, 1}
}
func (m *Tool_TableDefs_Columns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool_TableDefs_Columns.Unmarshal(m, b)
}
func (m *Tool_TableDefs_Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool_TableDefs_Columns.Marshal(b, m, deterministic)
}
func (m *Tool_TableDefs_Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool_TableDefs_Columns.Merge(m, src)
}
func (m *Tool_TableDefs_Columns) XXX_Size() int {
	return xxx_messageInfo_Tool_TableDefs_Columns.Size(m)
}
func (m *Tool_TableDefs_Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool_TableDefs_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_Tool_TableDefs_Columns proto.InternalMessageInfo

func (m *Tool_TableDefs_Columns) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tool_TableDefs_Columns) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tool_Log struct {
	//
	//是否启用日志
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled" form:"enabled"`
	//
	//v,vv,vvv for more, vvvv to enable connection debugging
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level" form:"level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tool_Log) Reset()         { *m = Tool_Log{} }
func (m *Tool_Log) String() string { return proto.CompactTextString(m) }
func (*Tool_Log) ProtoMessage()    {}
func (*Tool_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_34355de42c73af60, []int{0, 2}
}
func (m *Tool_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool_Log.Unmarshal(m, b)
}
func (m *Tool_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool_Log.Marshal(b, m, deterministic)
}
func (m *Tool_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool_Log.Merge(m, src)
}
func (m *Tool_Log) XXX_Size() int {
	return xxx_messageInfo_Tool_Log.Size(m)
}
func (m *Tool_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Tool_Log proto.InternalMessageInfo

func (m *Tool_Log) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Tool_Log) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*Tool)(nil), "tool.Tool")
	proto.RegisterType((*Tool_OutputDefs)(nil), "tool.Tool.OutputDefs")
	proto.RegisterType((*Tool_TableDefs)(nil), "tool.Tool.TableDefs")
	proto.RegisterType((*Tool_TableDefs_Dimensions)(nil), "tool.Tool.TableDefs.Dimensions")
	proto.RegisterType((*Tool_TableDefs_Columns)(nil), "tool.Tool.TableDefs.Columns")
	proto.RegisterType((*Tool_Log)(nil), "tool.Tool.Log")
}

func init() { proto.RegisterFile("tool.proto", fileDescriptor_34355de42c73af60) }

var fileDescriptor_34355de42c73af60 = []byte{
	// 1845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xdd, 0x52, 0x1c, 0xc7,
	0x15, 0xce, 0x0a, 0x24, 0x41, 0xaf, 0x00, 0xd1, 0x20, 0xdc, 0x22, 0x92, 0x07, 0xb5, 0x7e, 0xb2,
	0x58, 0x9a, 0x5d, 0x58, 0x2c, 0xcb, 0x42, 0xb6, 0x30, 0x48, 0x72, 0x45, 0x2e, 0xc9, 0xa8, 0x1a,
	0x44, 0x95, 0xb5, 0xc1, 0xd4, 0xec, 0x6c, 0x33, 0x4c, 0x69, 0x76, 0x7a, 0x33, 0x3f, 0x60, 0x2c,
	0xe9, 0x2d, 0xf2, 0x04, 0xb9, 0x4c, 0x55, 0x9e, 0x23, 0x79, 0x8a, 0x49, 0x55, 0xae, 0x52, 0x95,
	0xbb, 0x79, 0x82, 0x54, 0x9f, 0xf9, 0xeb, 0x99, 0xc5, 0x0a, 0x2e, 0xeb, 0x46, 0x37, 0x30, 0xdd,
	0xdf, 0xf9, 0xbe, 0xd3, 0xa7, 0xa7, 0xfb, 0x9c, 0x33, 0x8b, 0x50, 0x20, 0x84, 0xd3, 0x1c, 0x78,
	0x22, 0x10, 0x78, 0x54, 0x3e, 0xcf, 0xeb, 0x96, 0x1d, 0x1c, 0x84, 0xdd, 0xa6, 0x29, 0xfa, 0x2d,
	0x4b, 0x58, 0xa2, 0x05, 0x60, 0x37, 0xdc, 0x87, 0x11, 0x0c, 0xe0, 0x29, 0x21, 0xcd, 0x6f, 0x5a,
	0xa2, 0xc9, 0x0d, 0xff, 0x58, 0x0c, 0xfc, 0xa6, 0x23, 0x4c, 0xc3, 0x69, 0x99, 0xc2, 0x0d, 0x3c,
	0xc3, 0x0c, 0xfc, 0x84, 0xe9, 0xf1, 0x81, 0xd0, 0xfb, 0xa2, 0xc7, 0x1d, 0xbf, 0x95, 0x1a, 0xb6,
	0x60, 0xd8, 0x92, 0xee, 0xe0, 0xcf, 0x9e, 0xed, 0x0e, 0xc2, 0x20, 0x15, 0xfc, 0x42, 0xf1, 0xdf,
	0x3f, 0xb2, 0x83, 0xd7, 0xe2, 0xa8, 0x65, 0x09, 0x1d, 0x40, 0xfd, 0xd0, 0x70, 0xec, 0x9e, 0x11,
	0x08, 0xcf, 0x6f, 0xe5, 0x8f, 0x29, 0xef, 0x8a, 0x25, 0x84, 0xe5, 0xf0, 0x62, 0xb9, 0x7e, 0xe0,
	0x85, 0x66, 0xaa, 0x4a, 0xff, 0xda, 0x40, 0xa3, 0xdb, 0x42, 0x38, 0xf8, 0x1b, 0x74, 0x4e, 0xba,
	0x7c, 0xda, 0x23, 0xb5, 0x85, 0x5a, 0x63, 0x7c, 0xa3, 0x11, 0x47, 0xda, 0xc4, 0xbe, 0xf0, 0xfa,
	0xab, 0x34, 0x99, 0xa7, 0xff, 0xfe, 0x97, 0x36, 0x83, 0xa6, 0x7f, 0xec, 0x18, 0xfa, 0xfe, 0xba,
	0xfe, 0xed, 0x92, 0x7e, 0x7f, 0xf7, 0xcd, 0x4a, 0xfb, 0xdd, 0x0d, 0x96, 0xf2, 0xf0, 0x3d, 0x34,
	0x72, 0xf8, 0xb4, 0x47, 0xce, 0x00, 0xfd, 0x66, 0x1c, 0x69, 0x28, 0xa1, 0x1f, 0xbe, 0x87, 0x2b,
	0x19, 0xf8, 0x3a, 0x1a, 0x75, 0x8d, 0x3e, 0x27, 0x23, 0xc0, 0x9c, 0x8a, 0x23, 0xad, 0x9e, 0x30,
	0xe5, 0x2c, 0x65, 0x00, 0xe2, 0x3b, 0xe8, 0x7c, 0x60, 0xf7, 0xb9, 0x08, 0x03, 0x32, 0xba, 0x50,
	0x6b, 0x9c, 0xdd, 0xc0, 0x71, 0xa4, 0x4d, 0xa6, 0x0b, 0x4c, 0x00, 0xca, 0x32, 0x13, 0xdc, 0x42,
	0x63, 0xa6, 0x11, 0x70, 0x4b, 0x78, 0xc7, 0xe4, 0x2c, 0xc8, 0xce, 0xc4, 0x91, 0x36, 0x95, 0x98,
	0x67, 0x08, 0x65, 0xb9, 0x91, 0x5c, 0x83, 0x6d, 0x0a, 0x97, 0x9c, 0xab, 0xae, 0x41, 0xce, 0x52,
	0x06, 0x20, 0xbe, 0x85, 0xce, 0xfa, 0xc1, 0xb1, 0xc3, 0xc9, 0x79, 0xb0, 0xba, 0x18, 0x47, 0xda,
	0x85, 0xc4, 0x0a, 0xa6, 0x29, 0x4b, 0x60, 0x29, 0xd6, 0xe7, 0x7d, 0x41, 0xc6, 0xaa, 0x62, 0x72,
	0x96, 0x32, 0x00, 0xf1, 0x97, 0xa8, 0xee, 0xd8, 0x7e, 0xb0, 0x63, 0xfb, 0x76, 0xd7, 0xe1, 0x64,
	0x7c, 0xa1, 0xd6, 0x18, 0xdb, 0x98, 0x8b, 0x23, 0x0d, 0x27, 0xb6, 0x0a, 0x48, 0x99, 0x6a, 0x2a,
	0xe5, 0x03, 0xc3, 0xf2, 0x09, 0x5a, 0x18, 0x29, 0xcb, 0xcb, 0x59, 0xca, 0x00, 0x94, 0xfb, 0xd5,
	0xb3, 0x7d, 0x43, 0x4a, 0xd7, 0x41, 0x5a, 0xd9, 0xaf, 0x14, 0xa0, 0x2c, 0x33, 0xc1, 0xff, 0xa9,
	0x21, 0x64, 0x7a, 0xdc, 0x08, 0xf8, 0xb6, 0xdd, 0xe7, 0xe4, 0x02, 0x2c, 0xfc, 0x1f, 0xb5, 0x38,
	0xd2, 0xa6, 0xd3, 0x3d, 0xcb, 0x41, 0xf9, 0x2e, 0xff, 0x5e, 0x43, 0x7f, 0xab, 0xfd, 0xd8, 0x68,
	0xac, 0xad, 0x76, 0x96, 0xf5, 0xfb, 0xbb, 0x1d, 0xf9, 0x42, 0x3f, 0x5b, 0x5c, 0x83, 0xff, 0x6f,
	0x3e, 0x7f, 0xb7, 0xa8, 0x37, 0x96, 0x3b, 0x4b, 0x7a, 0x7b, 0xf7, 0xed, 0x12, 0xe0, 0x8b, 0x7a,
	0x63, 0xa5, 0xb3, 0xa4, 0x2f, 0x67, 0xe3, 0xb7, 0x9d, 0x65, 0xbd, 0x9d, 0xb0, 0x16, 0x3b, 0xdb,
	0x0b, 0xbb, 0x8d, 0x76, 0x67, 0x49, 0x5f, 0xd9, 0x7d, 0x0b, 0x36, 0xc9, 0xf4, 0x6a, 0xa3, 0xb3,
	0xa4, 0xdf, 0xcd, 0x06, 0xc5, 0x73, 0xe3, 0x4f, 0x4d, 0xf8, 0x7f, 0x7b, 0x71, 0xad, 0xf1, 0xea,
	0x6d, 0xe7, 0xb6, 0xbe, 0xdb, 0x58, 0x5b, 0x3d, 0x81, 0xae, 0xb0, 0xd7, 0x6e, 0x30, 0x25, 0x36,
	0x08, 0x35, 0x1c, 0xf4, 0xb2, 0x50, 0x27, 0x86, 0x42, 0x2d, 0xc0, 0x8f, 0x32, 0xd4, 0x62, 0xf9,
	0x78, 0x0b, 0x9d, 0x87, 0xc0, 0x85, 0x47, 0x26, 0x21, 0xcc, 0xfb, 0xc5, 0x19, 0x48, 0x01, 0x19,
	0xe2, 0x75, 0x74, 0x4d, 0xde, 0xcc, 0x9f, 0xd7, 0xf5, 0x57, 0x52, 0xa5, 0xd3, 0xcc, 0x9f, 0xf7,
	0xf4, 0xdd, 0x37, 0xed, 0x3b, 0x2b, 0xcb, 0xef, 0x6e, 0xb0, 0x4c, 0x09, 0xdf, 0x42, 0x23, 0xc2,
	0xb3, 0xc8, 0x14, 0x5c, 0xc2, 0xd9, 0xe2, 0x9a, 0x0b, 0xcf, 0x92, 0x62, 0x67, 0x2e, 0xfe, 0x8e,
	0x49, 0x03, 0xbc, 0x8a, 0xce, 0x41, 0xfa, 0xf2, 0xc9, 0xc5, 0x85, 0x91, 0x46, 0xbd, 0x3d, 0xd5,
	0x84, 0x94, 0x2a, 0x93, 0xcd, 0x53, 0x39, 0xbf, 0x31, 0x5d, 0x64, 0x98, 0xc4, 0x90, 0xb2, 0x94,
	0x21, 0x0f, 0xaf, 0x08, 0x03, 0x20, 0x4f, 0xc3, 0xc2, 0x95, 0xc3, 0x9b, 0x02, 0x94, 0x65, 0x26,
	0x70, 0x1f, 0x8e, 0x07, 0x9c, 0xe0, 0xea, 0x75, 0x93, 0xb3, 0xf2, 0x3e, 0x1c, 0x0f, 0x20, 0x7f,
	0xc8, 0x04, 0xcc, 0xdd, 0x80, 0xcc, 0x54, 0x25, 0x53, 0x80, 0xb2, 0xcc, 0x04, 0x3f, 0x46, 0x53,
	0x3d, 0xbe, 0x6f, 0x84, 0x4e, 0xf0, 0xe4, 0x27, 0x6e, 0xbe, 0xf4, 0xb9, 0x47, 0x66, 0x81, 0x35,
	0x1f, 0x47, 0xda, 0x5c, 0x7a, 0x8b, 0xca, 0x06, 0x94, 0x55, 0x29, 0x78, 0x07, 0x8d, 0xf1, 0x8c,
	0x7e, 0x09, 0xe8, 0xab, 0x45, 0x16, 0xca, 0x90, 0x53, 0xbf, 0x81, 0x5c, 0x0b, 0xff, 0x80, 0xc6,
	0x8d, 0x30, 0x38, 0x90, 0xcf, 0x3e, 0x99, 0x83, 0x2c, 0xf0, 0x20, 0x8e, 0xb4, 0x8b, 0x89, 0x70,
	0x0e, 0x9d, 0x5a, 0xb9, 0x50, 0xc3, 0x0f, 0xd1, 0x44, 0x1a, 0xc5, 0xba, 0xc5, 0xdd, 0xc0, 0x27,
	0x9f, 0x80, 0x3c, 0x89, 0x23, 0x6d, 0xb6, 0x14, 0x76, 0x02, 0x53, 0x56, 0x36, 0xc7, 0x5f, 0xa3,
	0x09, 0xff, 0xd8, 0x0f, 0x78, 0x7f, 0x6f, 0xe0, 0x84, 0x96, 0xed, 0x12, 0x02, 0xc9, 0x47, 0xe1,
	0x97, 0x60, 0xca, 0x2e, 0x24, 0xe3, 0x17, 0x30, 0xc4, 0xf7, 0x50, 0x3d, 0xc5, 0xe5, 0x61, 0x21,
	0x97, 0xab, 0x49, 0x51, 0x01, 0x29, 0x43, 0xc9, 0x08, 0xca, 0xd7, 0x5d, 0x84, 0x1c, 0x61, 0xbe,
	0x4e, 0x17, 0x3d, 0x0f, 0x9b, 0x7d, 0xa9, 0xb8, 0xd3, 0x05, 0x46, 0x99, 0x62, 0x28, 0x69, 0xbe,
	0xe1, 0xf6, 0xba, 0xe2, 0x27, 0x16, 0xba, 0xe4, 0xf7, 0xe0, 0x4e, 0xa1, 0x15, 0x98, 0xf4, 0x96,
	0x0f, 0x70, 0x1b, 0x8d, 0x1f, 0x1d, 0xd8, 0x01, 0x7f, 0x66, 0xfb, 0x01, 0xb9, 0x02, 0x3b, 0x34,
	0x5b, 0xbc, 0x80, 0x1c, 0xa2, 0xac, 0x30, 0xc3, 0xeb, 0x68, 0xf2, 0xc8, 0x76, 0x7b, 0xe2, 0xc8,
	0xdf, 0xe2, 0xbe, 0x6f, 0x0b, 0x97, 0x5c, 0x05, 0x77, 0x97, 0xe3, 0x48, 0xbb, 0x94, 0x12, 0x4b,
	0x38, 0x65, 0x15, 0x82, 0x74, 0xdb, 0x75, 0x0c, 0xf3, 0x35, 0xb8, 0xfd, 0xb4, 0xea, 0x36, 0x87,
	0x28, 0x2b, 0xcc, 0xf0, 0x0b, 0x54, 0x97, 0xbb, 0xb5, 0x99, 0x5e, 0x27, 0x6d, 0xa1, 0xd6, 0xa8,
	0xb7, 0xe7, 0x9a, 0x49, 0x53, 0xd0, 0xcc, 0x9a, 0x82, 0xe6, 0x8e, 0xe1, 0x84, 0x5c, 0xdd, 0x69,
	0x85, 0x44, 0x99, 0x2a, 0x81, 0x9f, 0x21, 0x94, 0xdc, 0xbc, 0xc7, 0x7c, 0xdf, 0x27, 0x0b, 0x70,
	0xb9, 0x2f, 0x15, 0x97, 0xbb, 0xb9, 0x99, 0x83, 0xea, 0x56, 0x16, 0x14, 0xca, 0x14, 0x3e, 0xfe,
	0x23, 0x1a, 0x0f, 0x64, 0x09, 0x02, 0xb1, 0x6b, 0x20, 0x36, 0xab, 0x88, 0x6d, 0x67, 0x98, 0x1a,
	0x69, 0x4e, 0xa0, 0xac, 0x20, 0xe3, 0xff, 0xd6, 0x50, 0xfd, 0xf0, 0x51, 0x51, 0xc4, 0x28, 0x1c,
	0x82, 0x7f, 0xd6, 0x8a, 0x98, 0x14, 0xf4, 0x63, 0x4c, 0xed, 0x6a, 0x74, 0x49, 0xb4, 0x2f, 0x8b,
	0x3a, 0x76, 0x7d, 0x38, 0xda, 0x97, 0x1f, 0x75, 0x21, 0x53, 0xa3, 0x93, 0x9d, 0xd7, 0xe1, 0xf7,
	0xb2, 0x47, 0xbc, 0x51, 0xed, 0xbc, 0x60, 0x9a, 0xb2, 0x04, 0x96, 0x19, 0x37, 0xd9, 0x24, 0xe1,
	0x91, 0x9b, 0xd5, 0x8c, 0x9b, 0x21, 0xa7, 0xcf, 0xb8, 0x19, 0x43, 0xd6, 0x03, 0x8f, 0x1b, 0xbd,
	0xf5, 0x30, 0x38, 0x10, 0x9e, 0xfd, 0xb3, 0xcc, 0xbb, 0xb7, 0xe0, 0xfe, 0x29, 0xf5, 0xa0, 0x62,
	0x40, 0x59, 0x95, 0x82, 0xbf, 0x43, 0xd3, 0x49, 0x75, 0x56, 0x75, 0xfe, 0x00, 0x3a, 0x57, 0xe2,
	0x48, 0x23, 0x6a, 0xff, 0x51, 0x52, 0x1a, 0xa6, 0x49, 0xad, 0x1e, 0x77, 0x78, 0x59, 0xab, 0x51,
	0xd5, 0x1a, 0x32, 0xa1, 0x6c, 0x98, 0x86, 0x9f, 0x23, 0x2c, 0x6b, 0x4b, 0x58, 0x16, 0x5b, 0x04,
	0xb1, 0xab, 0x71, 0xa4, 0x5d, 0x2e, 0x2a, 0x56, 0x58, 0x51, 0x3b, 0x81, 0x88, 0xff, 0x8c, 0x3e,
	0x95, 0x91, 0x3f, 0x01, 0xc4, 0x16, 0x2e, 0xe3, 0xbe, 0x2c, 0x10, 0x8a, 0xf4, 0x67, 0x20, 0xbd,
	0x18, 0x47, 0xda, 0xcd, 0x62, 0xef, 0x7e, 0xd9, 0x9e, 0xb2, 0xff, 0x23, 0x88, 0x7f, 0x40, 0x73,
	0x9e, 0x10, 0xc1, 0x93, 0xe1, 0x28, 0x6e, 0x83, 0xab, 0x6b, 0x71, 0xa4, 0x5d, 0x4d, 0x5d, 0x9d,
	0x68, 0x47, 0xd9, 0x2f, 0x08, 0xc0, 0xd1, 0x7b, 0xcc, 0x7d, 0x93, 0xdc, 0x19, 0x3a, 0x7a, 0x72,
	0x5a, 0x1e, 0x3d, 0xf9, 0x1f, 0x4a, 0x89, 0x08, 0x3d, 0x93, 0x7f, 0xeb, 0x89, 0x3e, 0xd1, 0xab,
	0x15, 0xa8, 0xc0, 0x64, 0x29, 0xc9, 0x07, 0xb2, 0x2f, 0xe1, 0xee, 0xe1, 0xb6, 0xec, 0x5f, 0x9a,
	0xd5, 0xbe, 0x24, 0x05, 0x28, 0xcb, 0x4c, 0x64, 0x05, 0x30, 0x0f, 0xb8, 0xf9, 0x1a, 0xec, 0x5b,
	0x60, 0xaf, 0xe4, 0xc5, 0x1c, 0xa2, 0xac, 0x30, 0x93, 0xdd, 0x02, 0x0c, 0xa0, 0x0d, 0x59, 0x02,
	0xce, 0x83, 0x0a, 0xe7, 0x57, 0xf5, 0x21, 0x85, 0x1a, 0xee, 0xa2, 0x0b, 0x30, 0xd8, 0x1a, 0x08,
	0xd7, 0x17, 0x1e, 0x59, 0x06, 0xf5, 0x87, 0x71, 0xa4, 0xcd, 0x28, 0xea, 0x29, 0x7a, 0x6a, 0x07,
	0x25, 0x4d, 0xfc, 0x20, 0xf5, 0xf1, 0x48, 0xf4, 0xfb, 0xb2, 0x7b, 0x6b, 0x83, 0x8f, 0x4f, 0x2a,
	0x3e, 0x52, 0x94, 0xb2, 0x92, 0x31, 0xde, 0x41, 0x13, 0x5d, 0x23, 0x30, 0x0f, 0xb6, 0x02, 0x4f,
	0x7e, 0xe9, 0x1d, 0x93, 0x95, 0xf7, 0xd6, 0x3f, 0xa5, 0x4d, 0x29, 0xd1, 0x28, 0x2b, 0xcb, 0xc8,
	0x45, 0xed, 0x87, 0xae, 0x29, 0x0f, 0x23, 0xbc, 0x8a, 0xcf, 0xab, 0x8b, 0x52, 0x51, 0xca, 0x4a,
	0xc6, 0xf8, 0x2b, 0x34, 0x61, 0xfb, 0x7b, 0x69, 0x2b, 0x23, 0x7b, 0xe9, 0xbb, 0xd5, 0x1e, 0xa9,
	0x04, 0x53, 0x56, 0xb7, 0xfd, 0x2d, 0x18, 0x6e, 0x7a, 0x96, 0xfc, 0x6e, 0xf4, 0xc3, 0xae, 0x6f,
	0x7a, 0x76, 0x57, 0x9e, 0xef, 0x2f, 0xe0, 0x7c, 0xab, 0x2d, 0x52, 0x01, 0x52, 0xa6, 0x9a, 0xca,
	0x94, 0x91, 0x0f, 0x7b, 0x8f, 0x0e, 0x0c, 0xd7, 0xe5, 0x0e, 0xb9, 0x07, 0x2b, 0x57, 0x52, 0xc6,
	0x90, 0x09, 0x65, 0xc3, 0x34, 0xf9, 0x81, 0x2d, 0xaf, 0xe4, 0xa6, 0xeb, 0x1c, 0x93, 0x2f, 0x61,
	0xf9, 0xca, 0x07, 0x76, 0x86, 0x50, 0x96, 0x1b, 0xc9, 0x1d, 0x0b, 0x78, 0x7f, 0xe0, 0xc8, 0x8c,
	0x2e, 0x77, 0xec, 0x7e, 0x75, 0xc7, 0x54, 0x94, 0xb2, 0x92, 0x31, 0x5e, 0x46, 0xe3, 0x49, 0xd6,
	0xda, 0xeb, 0x73, 0xb2, 0x0a, 0xee, 0x94, 0x63, 0x9f, 0x43, 0x94, 0x8d, 0x25, 0xcf, 0xcf, 0xe1,
	0xd4, 0x1b, 0x83, 0x81, 0x27, 0x0e, 0xe5, 0x26, 0x3d, 0x18, 0xea, 0x91, 0x33, 0xe8, 0x57, 0xf4,
	0xc8, 0x19, 0x05, 0x1f, 0xa1, 0x19, 0xc3, 0x85, 0x4f, 0xf1, 0xef, 0x45, 0x8f, 0xe7, 0x1f, 0x08,
	0x5f, 0x41, 0x44, 0x4f, 0xe2, 0x48, 0x9b, 0x4f, 0x9d, 0x0c, 0x1b, 0x9d, 0xda, 0xdd, 0x49, 0x1e,
	0xf0, 0x12, 0x1a, 0x71, 0x84, 0x45, 0xbe, 0x86, 0x33, 0x3c, 0xa9, 0x74, 0x49, 0xcf, 0x84, 0xb5,
	0x31, 0x59, 0x7c, 0x8a, 0x39, 0xc2, 0xa2, 0x4c, 0x9a, 0xce, 0xbf, 0x40, 0xa8, 0x68, 0xc7, 0xf0,
	0x55, 0x74, 0xc6, 0xce, 0x7e, 0xdf, 0x99, 0x88, 0x23, 0x6d, 0x3c, 0x3d, 0x6d, 0x3d, 0xca, 0xce,
	0xd8, 0xc5, 0xef, 0x30, 0x67, 0xde, 0xf3, 0x3b, 0xcc, 0xfc, 0x5f, 0x46, 0xd0, 0x78, 0xde, 0x94,
	0x7d, 0x08, 0x45, 0xbc, 0x83, 0x50, 0xcf, 0xee, 0x73, 0x57, 0xb6, 0xb8, 0x3e, 0x19, 0x81, 0x16,
	0x50, 0x3b, 0xa9, 0x05, 0x6c, 0x3e, 0xce, 0xcd, 0xd4, 0xcc, 0x5a, 0x90, 0x29, 0x53, 0x94, 0xf0,
	0x77, 0xf2, 0x8b, 0xcf, 0x09, 0xfb, 0xae, 0x4f, 0x46, 0x41, 0xf4, 0xca, 0x89, 0xa2, 0x8f, 0x12,
	0x9b, 0xf2, 0xf7, 0x20, 0x4c, 0xc1, 0xf7, 0x20, 0x3c, 0xc9, 0x7d, 0x2c, 0x9c, 0x7f, 0x90, 0x7d,
	0x7c, 0x8e, 0xce, 0xa7, 0x9e, 0x3f, 0x88, 0x5c, 0x07, 0x8d, 0x3c, 0x13, 0x56, 0x52, 0x4d, 0x64,
	0x68, 0x89, 0xde, 0x58, 0xb9, 0x9a, 0x00, 0x00, 0xd5, 0x04, 0x9e, 0x64, 0x69, 0x73, 0xf8, 0x21,
	0x77, 0x52, 0x69, 0xa5, 0xb4, 0xc1, 0x34, 0x65, 0x09, 0xbc, 0xf1, 0xcd, 0xab, 0x87, 0xbf, 0xed,
	0xd7, 0xcc, 0xee, 0x39, 0x30, 0x5a, 0xf9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5f, 0xf8,
	0xa9, 0x57, 0x15, 0x00, 0x00,
}
