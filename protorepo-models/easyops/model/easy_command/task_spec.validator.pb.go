// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task_spec.proto

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

var _regex_TaskSpec_TaskId = regexp.MustCompile(`^([0-9]{4})(1[0-2]|0[1-9])_task[0-9a-f]{16,20}$`)
var _regex_TaskSpec_AppId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_TaskSpec_ClusterId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *TaskSpec) Validate() error {
	if !_regex_TaskSpec_TaskId.MatchString(this.TaskId) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([0-9]{4})(1[0-2]|0[1-9])_task[0-9a-f]{16,20}$"`, this.TaskId))
	}
	for _, item := range this.Actions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Actions", err)
			}
		}
	}
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	if !_regex_TaskSpec_AppId.MatchString(this.AppId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.AppId))
	}
	if !_regex_TaskSpec_ClusterId.MatchString(this.ClusterId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ClusterId))
	}
	if this.Callback != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Callback); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Callback", err)
		}
	}
	return nil
}
