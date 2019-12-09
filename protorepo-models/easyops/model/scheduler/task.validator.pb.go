// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.proto

package scheduler

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

var _regex_SchedulerTask_Id = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
var _regex_SchedulerTask_User = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_SchedulerTask_OperateAuthorizers = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_SchedulerTask_DeleteAuthorizers = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_SchedulerTask_UpdateAuthorizers = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_SchedulerTask_CreateTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_SchedulerTask_UpdateTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)

func (this *SchedulerTask) Validate() error {
	if !_regex_SchedulerTask_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.Id))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	if !(len(this.TaskScheduler) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskScheduler", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.TaskScheduler))
	}
	if this.Annotations != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Annotations); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Annotations", err)
		}
	}
	if !(len(this.SrcId) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("SrcId", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.SrcId))
	}
	if this.CmdConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdConfig", err)
		}
	}
	if !(this.Org > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must be greater than '0'`, this.Org))
	}
	if !_regex_SchedulerTask_User.MatchString(this.User) {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.User))
	}
	for _, item := range this.OperateAuthorizers {
		if !_regex_SchedulerTask_OperateAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("OperateAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, item))
		}
	}
	for _, item := range this.DeleteAuthorizers {
		if !_regex_SchedulerTask_DeleteAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("DeleteAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, item))
		}
	}
	for _, item := range this.UpdateAuthorizers {
		if !_regex_SchedulerTask_UpdateAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, item))
		}
	}
	if !_regex_SchedulerTask_CreateTime.MatchString(this.CreateTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.CreateTime))
	}
	if !_regex_SchedulerTask_UpdateTime.MatchString(this.UpdateTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.UpdateTime))
	}
	if !(len(this.JobId) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("JobId", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.JobId))
	}
	if this.Callback != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Callback); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Callback", err)
		}
	}
	return nil
}
func (this *SchedulerTask_Annotations) Validate() error {
	return nil
}
func (this *SchedulerTask_CmdConfig) Validate() error {
	if !(this.Port > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Port))
	}
	if !(this.Port < 65536) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be less than '65536'`, this.Port))
	}
	if this.Headers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Headers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Headers", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

var _regex_SchedulerTask_Callback_Url = regexp.MustCompile(`(\b(https?|ftp|file)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)

func (this *SchedulerTask_Callback) Validate() error {
	if !_regex_SchedulerTask_Callback_Url.MatchString(this.Url) {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`value '%v' must be a string conforming to regex "(\\b(https?|ftp|file)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]"`, this.Url))
	}
	return nil
}
