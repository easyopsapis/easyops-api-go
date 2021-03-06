// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_subsystem.proto

package business_instance

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

var _regex_SearchSubSystemRequest_SystemInstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *SearchSubSystemRequest) Validate() error {
	if !_regex_SearchSubSystemRequest_SystemInstanceId.MatchString(this.SystemInstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("SystemInstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.SystemInstanceId))
	}
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Fields != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fields); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
		}
	}
	if this.Sort != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sort); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sort", err)
		}
	}
	return nil
}
func (this *SearchSubSystemResponse) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *SearchSubSystemResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
