// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_role.proto

package role

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

var _regex_CreateRoleRequest_User = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_CreateRoleRequest_Permission = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *CreateRoleRequest) Validate() error {
	for _, item := range this.User {
		if !_regex_CreateRoleRequest_User.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, item))
		}
	}
	for _, item := range this.Permission {
		if !_regex_CreateRoleRequest_Permission.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Permission", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, item))
		}
	}
	return nil
}

var _regex_CreateRoleResponse_Id = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *CreateRoleResponse) Validate() error {
	if !_regex_CreateRoleResponse_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.Id))
	}
	return nil
}
func (this *CreateRoleResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
