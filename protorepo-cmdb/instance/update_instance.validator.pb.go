// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_instance.proto

package instance

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

var _regex_UpdateInstanceRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_UpdateInstanceRequest_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *UpdateInstanceRequest) Validate() error {
	if !_regex_UpdateInstanceRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if !_regex_UpdateInstanceRequest_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if this.Instance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instance", err)
		}
	}
	return nil
}
func (this *UpdateInstanceResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
