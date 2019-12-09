// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_alert_rule.proto

package metadata_center

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

var _regex_StreamAlertRule_Id = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *StreamAlertRule) Validate() error {
	if !_regex_StreamAlertRule_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.Id))
	}
	for _, item := range this.Dimensions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dimensions", err)
			}
		}
	}
	for _, item := range this.Conditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Conditions", err)
			}
		}
	}
	return nil
}
func (this *StreamAlertRule_Dimensions) Validate() error {
	return nil
}
func (this *StreamAlertRule_Conditions) Validate() error {
	if this.Threshold != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Threshold); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Threshold", err)
		}
	}
	return nil
}