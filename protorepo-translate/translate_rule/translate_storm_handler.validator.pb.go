// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: translate_storm_handler.proto

package translate_rule

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

func (this *TranslateStormHandlerRequest) Validate() error {
	return nil
}
func (this *TranslateStormHandlerResponse) Validate() error {
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.Page < 3001) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be less than '3001'`, this.Page))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TranslateStormHandlerResponse_Data) Validate() error {
	if this.MatchRules != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MatchRules); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MatchRules", err)
		}
	}
	return nil
}
func (this *TranslateStormHandlerResponse_Data_MatchRules) Validate() error {
	for _, item := range this.MatchFields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MatchFields", err)
			}
		}
	}
	return nil
}

var _regex_TranslateStormHandlerResponse_Data_MatchRules_MatchFields_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *TranslateStormHandlerResponse_Data_MatchRules_MatchFields) Validate() error {
	if !_regex_TranslateStormHandlerResponse_Data_MatchRules_MatchFields_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	return nil
}
func (this *TranslateStormHandlerResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}