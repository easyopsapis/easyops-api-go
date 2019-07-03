// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_permission_role_list.proto

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

var _regex_GetPermissionRoleListRequest_User = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)

func (this *GetPermissionRoleListRequest) Validate() error {
	if !_regex_GetPermissionRoleListRequest_User.MatchString(this.User) {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, this.User))
	}
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *GetPermissionRoleListResponse) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}

var _regex_GetPermissionRoleListResponse_Data_Id = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
var _regex_GetPermissionRoleListResponse_Data_User = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)
var _regex_GetPermissionRoleListResponse_Data_Permission = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *GetPermissionRoleListResponse_Data) Validate() error {
	if !_regex_GetPermissionRoleListResponse_Data_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.Id))
	}
	if !(this.Org > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must be greater than '0'`, this.Org))
	}
	for _, item := range this.User {
		if !_regex_GetPermissionRoleListResponse_Data_User.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, item))
		}
	}
	for _, item := range this.Permission {
		if !_regex_GetPermissionRoleListResponse_Data_Permission.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Permission", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, item))
		}
	}
	return nil
}
func (this *GetPermissionRoleListResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
