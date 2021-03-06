// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_custom_dbtask_rollbackinfo.proto

package dbtask

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

var _regex_GetCustomDBTaskRollbackInfoRequest_DbTaskId = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *GetCustomDBTaskRollbackInfoRequest) Validate() error {
	if !_regex_GetCustomDBTaskRollbackInfoRequest_DbTaskId.MatchString(this.DbTaskId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DbTaskId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.DbTaskId))
	}
	return nil
}
func (this *GetCustomDBTaskRollbackInfoResponse) Validate() error {
	for _, item := range this.ChangesetList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ChangesetList", err)
			}
		}
	}
	return nil
}

var _regex_GetCustomDBTaskRollbackInfoResponse_ChangesetList_Id = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *GetCustomDBTaskRollbackInfoResponse_ChangesetList) Validate() error {
	if !_regex_GetCustomDBTaskRollbackInfoResponse_ChangesetList_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.Id))
	}
	return nil
}
func (this *GetCustomDBTaskRollbackInfoResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
