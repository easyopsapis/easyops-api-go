// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flow.proto

package flow

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

var _regex_Flow_FlowId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_Flow_CreateTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_Flow_Creator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_Flow_VCreator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_Flow_UpdateTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_Flow_ReadAuthorizers = regexp.MustCompile(`(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$`)
var _regex_Flow_UpdateAuthorizers = regexp.MustCompile(`(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$`)
var _regex_Flow_DeleteAuthorizers = regexp.MustCompile(`(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$`)
var _regex_Flow_ExecuteAuthorizers = regexp.MustCompile(`(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$`)
var _regex_Flow_Subscribers = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *Flow) Validate() error {
	if !_regex_Flow_FlowId.MatchString(this.FlowId) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.FlowId))
	}
	if !(this.Org > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must be greater than '0'`, this.Org))
	}
	if !_regex_Flow_CreateTime.MatchString(this.CreateTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.CreateTime))
	}
	if !_regex_Flow_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Creator))
	}
	if !_regex_Flow_VCreator.MatchString(this.VCreator) {
		return github_com_mwitkow_go_proto_validators.FieldError("VCreator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.VCreator))
	}
	if !_regex_Flow_UpdateTime.MatchString(this.UpdateTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.UpdateTime))
	}
	for _, item := range this.ReadAuthorizers {
		if !_regex_Flow_ReadAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("ReadAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$"`, item))
		}
	}
	for _, item := range this.UpdateAuthorizers {
		if !_regex_Flow_UpdateAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$"`, item))
		}
	}
	for _, item := range this.DeleteAuthorizers {
		if !_regex_Flow_DeleteAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("DeleteAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$"`, item))
		}
	}
	for _, item := range this.ExecuteAuthorizers {
		if !_regex_Flow_ExecuteAuthorizers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("ExecuteAuthorizers", fmt.Errorf(`value '%v' must be a string conforming to regex "(^[a-zA-Z0-9][a-zA-Z0-9_-]{2,31})$|(^:[0-9a-z]{13})$"`, item))
		}
	}
	for _, item := range this.Subscribers {
		if !_regex_Flow_Subscribers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Subscribers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, item))
		}
	}
	for _, item := range this.StepList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StepList", err)
			}
		}
	}
	for _, item := range this.TableDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TableDefs", err)
			}
		}
	}
	if this.FlowEnv != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FlowEnv); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FlowEnv", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if this.FlowInputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FlowInputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FlowInputs", err)
		}
	}
	for _, item := range this.FlowOutputs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FlowOutputs", err)
			}
		}
	}
	for _, item := range this.OutputDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OutputDefs", err)
			}
		}
	}
	for _, item := range this.Histories {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Histories", err)
			}
		}
	}
	return nil
}
func (this *Flow_TableDefs) Validate() error {
	for _, item := range this.Dimensions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dimensions", err)
			}
		}
	}
	for _, item := range this.Columns {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Columns", err)
			}
		}
	}
	return nil
}
func (this *Flow_TableDefs_Dimensions) Validate() error {
	return nil
}
func (this *Flow_TableDefs_Columns) Validate() error {
	return nil
}
func (this *Flow_Metadata) Validate() error {
	return nil
}
func (this *Flow_FlowOutputs) Validate() error {
	for _, item := range this.Columns {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Columns", err)
			}
		}
	}
	return nil
}
func (this *Flow_FlowOutputs_Columns) Validate() error {
	return nil
}
func (this *Flow_OutputDefs) Validate() error {
	return nil
}