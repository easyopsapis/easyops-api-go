// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: build.proto

package pipeline

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

var _regex_Build_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_Build_Sender = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *Build) Validate() error {
	if !_regex_Build_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	if this.GitMeta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GitMeta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GitMeta", err)
		}
	}
	if !_regex_Build_Sender.MatchString(this.Sender) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sender", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Sender))
	}
	if this.Artifact != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Artifact); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Artifact", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}

var _regex_Build_Artifact_Ctime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_Build_Artifact_PackageId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_Build_Artifact_VersionId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *Build_Artifact) Validate() error {
	if !_regex_Build_Artifact_Ctime.MatchString(this.Ctime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ctime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Ctime))
	}
	if !_regex_Build_Artifact_PackageId.MatchString(this.PackageId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageId))
	}
	if !_regex_Build_Artifact_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.VersionId))
	}
	return nil
}
