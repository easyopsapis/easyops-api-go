// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validate_cmdb_permission.proto

package permission

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

var _regex_ValidateCmdbPermissionRequest_User = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_ValidateCmdbPermissionRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_ValidateCmdbPermissionRequest_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *ValidateCmdbPermissionRequest) Validate() error {
	if !_regex_ValidateCmdbPermissionRequest_User.MatchString(this.User) {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.User))
	}
	if !_regex_ValidateCmdbPermissionRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if !_regex_ValidateCmdbPermissionRequest_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	return nil
}
func (this *ValidateCmdbPermissionResponse) Validate() error {
	return nil
}
func (this *ValidateCmdbPermissionResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
