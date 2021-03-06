// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: translate_rule_list_handler.proto

package translate_rule_list_handler

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TranslateRuleListHandlerRequest) Validate() error {
	return nil
}
func (this *TranslateRuleListHandlerResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *TranslateRuleListHandlerResponse_Data) Validate() error {
	for _, item := range this.DataNames {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DataNames", err)
			}
		}
	}
	for _, item := range this.TranslateFields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TranslateFields", err)
			}
		}
	}
	for _, item := range this.MatchFields {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MatchFields", err)
			}
		}
	}
	return nil
}
func (this *TranslateRuleListHandlerResponse_Data_DataNames) Validate() error {
	return nil
}
func (this *TranslateRuleListHandlerResponse_Data_TranslateFields) Validate() error {
	return nil
}
func (this *TranslateRuleListHandlerResponse_Data_MatchFields) Validate() error {
	return nil
}
func (this *TranslateRuleListHandlerResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
