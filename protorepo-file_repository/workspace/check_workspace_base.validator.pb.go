// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: check_workspace_base.proto

package workspace

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

var _regex_CheckWorkspaceBaseRequest_PackageId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_CheckWorkspaceBaseRequest_WorkspaceBaseId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *CheckWorkspaceBaseRequest) Validate() error {
	if !_regex_CheckWorkspaceBaseRequest_PackageId.MatchString(this.PackageId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageId))
	}
	if !_regex_CheckWorkspaceBaseRequest_WorkspaceBaseId.MatchString(this.WorkspaceBaseId) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkspaceBaseId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.WorkspaceBaseId))
	}
	return nil
}

var _regex_CheckWorkspaceBaseResponse_WorkspaceBaseId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *CheckWorkspaceBaseResponse) Validate() error {
	if !_regex_CheckWorkspaceBaseResponse_WorkspaceBaseId.MatchString(this.WorkspaceBaseId) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkspaceBaseId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.WorkspaceBaseId))
	}
	return nil
}
func (this *CheckWorkspaceBaseResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
