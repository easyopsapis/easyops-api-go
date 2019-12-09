// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: batch_delete_versions.proto

package version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_BatchDeleteVersionsRequest_VersionIds = regexp.MustCompile(`^[0-9a-zA-Z]+(,[0-9a-zA-Z]+)*$`)
var _regex_BatchDeleteVersionsRequest_PackageId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *BatchDeleteVersionsRequest) Validate() error {
	if !_regex_BatchDeleteVersionsRequest_VersionIds.MatchString(this.VersionIds) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionIds", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-zA-Z]+(,[0-9a-zA-Z]+)*$"`, this.VersionIds))
	}
	if !_regex_BatchDeleteVersionsRequest_PackageId.MatchString(this.PackageId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageId))
	}
	return nil
}
func (this *BatchDeleteVersionsResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}