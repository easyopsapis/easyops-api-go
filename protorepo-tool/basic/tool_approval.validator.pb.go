// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool_approval.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ToolApprovalRequest_Origin = regexp.MustCompile(`(\b(https?|ftp|file)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
var _regex_ToolApprovalRequest_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_ToolApprovalRequest_VId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *ToolApprovalRequest) Validate() error {
	if !_regex_ToolApprovalRequest_Origin.MatchString(this.Origin) {
		return github_com_mwitkow_go_proto_validators.FieldError("Origin", fmt.Errorf(`value '%v' must be a string conforming to regex "(\\b(https?|ftp|file)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]"`, this.Origin))
	}
	if !_regex_ToolApprovalRequest_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	if !_regex_ToolApprovalRequest_VId.MatchString(this.VId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VId))
	}
	return nil
}

var _regex_ToolApprovalResponse_CheckSponsor = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_ToolApprovalResponse_CheckUser = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *ToolApprovalResponse) Validate() error {
	if !_regex_ToolApprovalResponse_CheckSponsor.MatchString(this.CheckSponsor) {
		return github_com_mwitkow_go_proto_validators.FieldError("CheckSponsor", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.CheckSponsor))
	}
	if !_regex_ToolApprovalResponse_CheckUser.MatchString(this.CheckUser) {
		return github_com_mwitkow_go_proto_validators.FieldError("CheckUser", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.CheckUser))
	}
	return nil
}
func (this *ToolApprovalResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
