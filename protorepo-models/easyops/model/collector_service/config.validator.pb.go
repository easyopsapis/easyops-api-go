// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package collector_service

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

var _regex_CollectorConfig_ConfigId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CollectorConfig_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *CollectorConfig) Validate() error {
	if !_regex_CollectorConfig_ConfigId.MatchString(this.ConfigId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ConfigId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ConfigId))
	}
	if !_regex_CollectorConfig_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	return nil
}
func (this *CollectorConfig_Params) Validate() error {
	return nil
}