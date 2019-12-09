// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: detail_collection_config.proto

package collection_config

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

var _regex_DetailCollectionConfigRequest_Id = regexp.MustCompile(`[0-9a-f]{24}`)

func (this *DetailCollectionConfigRequest) Validate() error {
	if !_regex_DetailCollectionConfigRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "[0-9a-f]{24}"`, this.Id))
	}
	return nil
}

var _regex_DetailCollectionConfigResponse_Id = regexp.MustCompile(`[0-9a-f]{24}`)
var _regex_DetailCollectionConfigResponse_TimeRange = regexp.MustCompile(`[0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}`)
var _regex_DetailCollectionConfigResponse_ClazzId = regexp.MustCompile(`[0-9a-f]{24}`)

func (this *DetailCollectionConfigResponse) Validate() error {
	if !_regex_DetailCollectionConfigResponse_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "[0-9a-f]{24}"`, this.Id))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	for _, item := range this.Kwargs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Kwargs", err)
			}
		}
	}
	for _, item := range this.Env {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Env", err)
			}
		}
	}
	if this.Script != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Script); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Script", err)
		}
	}
	if this.HostRange != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HostRange); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HostRange", err)
		}
	}
	if !(this.Interval > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Interval", fmt.Errorf(`value '%v' must be greater than '0'`, this.Interval))
	}
	if !_regex_DetailCollectionConfigResponse_TimeRange.MatchString(this.TimeRange) {
		return github_com_mwitkow_go_proto_validators.FieldError("TimeRange", fmt.Errorf(`value '%v' must be a string conforming to regex "[0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}"`, this.TimeRange))
	}
	if !_regex_DetailCollectionConfigResponse_ClazzId.MatchString(this.ClazzId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClazzId", fmt.Errorf(`value '%v' must be a string conforming to regex "[0-9a-f]{24}"`, this.ClazzId))
	}
	return nil
}
func (this *DetailCollectionConfigResponse_Kwargs) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *DetailCollectionConfigResponse_Env) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *DetailCollectionConfigResponse_Script) Validate() error {
	return nil
}
func (this *DetailCollectionConfigResponse_HostRange) Validate() error {
	for _, item := range this.Query {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
			}
		}
	}
	return nil
}
func (this *DetailCollectionConfigResponse_HostRange_Query) Validate() error {
	if this.Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Value", err)
		}
	}
	return nil
}
func (this *DetailCollectionConfigResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}