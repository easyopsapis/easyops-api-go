// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alert_event.proto

package monitor

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

var _regex_AlertEvent_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_AlertEvent_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *AlertEvent) Validate() error {
	if !(this.Org > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Org", fmt.Errorf(`value '%v' must be greater than '0'`, this.Org))
	}
	for _, item := range this.AlertReceivers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertReceivers", err)
			}
		}
	}
	for _, item := range this.AlertDims {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AlertDims", err)
			}
		}
	}
	for _, item := range this.Actions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Actions", err)
			}
		}
	}
	if this.AlertConditions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AlertConditions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AlertConditions", err)
		}
	}
	if !_regex_AlertEvent_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if !_regex_AlertEvent_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	return nil
}
func (this *AlertEvent_AlertReceivers) Validate() error {
	return nil
}
func (this *AlertEvent_AlertDims) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *AlertEvent_Actions) Validate() error {
	if this.Condition != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Condition); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Condition", err)
		}
	}
	for _, item := range this.ReceiverOwners {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ReceiverOwners", err)
			}
		}
	}
	return nil
}
func (this *AlertEvent_Actions_Condition) Validate() error {
	return nil
}
func (this *AlertEvent_Actions_ReceiverOwners) Validate() error {
	return nil
}
