// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commit_worksapce_v2.proto

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

var _regex_CommitWorkspaceV2Request_VersionId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_CommitWorkspaceV2Request_PackageId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *CommitWorkspaceV2Request) Validate() error {
	if !_regex_CommitWorkspaceV2Request_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.VersionId))
	}
	if !_regex_CommitWorkspaceV2Request_PackageId.MatchString(this.PackageId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageId))
	}
	return nil
}

var _regex_CommitWorkspaceV2Response_Data = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *CommitWorkspaceV2Response) Validate() error {
	if !_regex_CommitWorkspaceV2Response_Data.MatchString(this.Data) {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.Data))
	}
	return nil
}
func (this *CommitWorkspaceV2ResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}