// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_menu.proto

package menu

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

func (this *GetMenuRequest) Validate() error {
	return nil
}

var _regex_GetMenuResponse_Visitors = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)
var _regex_GetMenuResponse_Managers = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)

func (this *GetMenuResponse) Validate() error {
	if !(len(this.CreateTime) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must length be greater than '1'`, this.CreateTime))
	}
	if !(len(this.CreateTime) < 34) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must length be less than '34'`, this.CreateTime))
	}
	if !(len(this.UpdateTime) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must length be greater than '1'`, this.UpdateTime))
	}
	if !(len(this.UpdateTime) < 34) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must length be less than '34'`, this.UpdateTime))
	}
	if !(this.Org > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must be greater than '0'`, this.Org))
	}
	if !(len(this.Category) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Category", fmt.Errorf(`value '%v' must length be greater than '-1'`, this.Category))
	}
	for _, item := range this.Visitors {
		if !_regex_GetMenuResponse_Visitors.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Visitors", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, item))
		}
	}
	for _, item := range this.Managers {
		if !_regex_GetMenuResponse_Managers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Managers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, item))
		}
	}
	return nil
}
func (this *GetMenuResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}