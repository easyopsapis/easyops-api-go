// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task_detail.proto

package easy_command

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

var _regex_TaskDetail_AppId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_TaskDetail_ClusterId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_TaskDetail_Operator = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)
var _regex_TaskDetail_StartTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_TaskDetail_UpdateTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_TaskDetail_EndTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)

func (this *TaskDetail) Validate() error {
	if !_regex_TaskDetail_AppId.MatchString(this.AppId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.AppId))
	}
	if !_regex_TaskDetail_ClusterId.MatchString(this.ClusterId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ClusterId))
	}
	if this.Callback != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Callback); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Callback", err)
		}
	}
	if !_regex_TaskDetail_Operator.MatchString(this.Operator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Operator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, this.Operator))
	}
	if !_regex_TaskDetail_StartTime.MatchString(this.StartTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("StartTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.StartTime))
	}
	if !_regex_TaskDetail_UpdateTime.MatchString(this.UpdateTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.UpdateTime))
	}
	if !_regex_TaskDetail_EndTime.MatchString(this.EndTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("EndTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.EndTime))
	}
	for _, item := range this.TargetsLog {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TargetsLog", err)
			}
		}
	}
	return nil
}