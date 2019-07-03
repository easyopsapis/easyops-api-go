// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stage_status.proto

package pipeline

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

var _regex_StageStatus_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *StageStatus) Validate() error {
	if !_regex_StageStatus_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	for _, item := range this.Steps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Steps", err)
			}
		}
	}
	return nil
}

var _regex_StageStatus_Steps_Id = regexp.MustCompile(`^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$`)
var _regex_StageStatus_Steps_LogId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *StageStatus_Steps) Validate() error {
	if !_regex_StageStatus_Steps_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$"`, this.Id))
	}
	if !_regex_StageStatus_Steps_LogId.MatchString(this.LogId) {
		return github_com_mwitkow_go_proto_validators.FieldError("LogId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.LogId))
	}
	return nil
}