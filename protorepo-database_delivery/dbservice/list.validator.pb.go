// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package dbservice

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/database_delivery"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListDBServiceRequest) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *ListDBServiceResponse) Validate() error {
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

var _regex_ListDBServiceResponse_List_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *ListDBServiceResponse_List) Validate() error {
	for _, item := range this.DbInstance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DbInstance", err)
			}
		}
	}
	for _, item := range this.Owner {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Owner", err)
			}
		}
	}
	if !_regex_ListDBServiceResponse_List_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	return nil
}

var _regex_ListDBServiceResponse_List_Owner_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_ListDBServiceResponse_List_Owner_UserEmail = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)
var _regex_ListDBServiceResponse_List_Owner_Nickname = regexp.MustCompile(`.*`)

func (this *ListDBServiceResponse_List_Owner) Validate() error {
	if !_regex_ListDBServiceResponse_List_Owner_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !_regex_ListDBServiceResponse_List_Owner_UserEmail.MatchString(this.UserEmail) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserEmail", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.UserEmail))
	}
	if !_regex_ListDBServiceResponse_List_Owner_Nickname.MatchString(this.Nickname) {
		return github_com_mwitkow_go_proto_validators.FieldError("Nickname", fmt.Errorf(`value '%v' must be a string conforming to regex ".*"`, this.Nickname))
	}
	return nil
}
func (this *ListDBServiceResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
