// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: idcrack_unit_device.proto

package cmdb_extend

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

var _regex_IdcrackUnitDevice_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *IdcrackUnitDevice) Validate() error {
	if !_regex_IdcrackUnitDevice_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if this.Device != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Device); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Device", err)
		}
	}
	return nil
}

var _regex_IdcrackUnitDevice_Device_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *IdcrackUnitDevice_Device) Validate() error {
	if !_regex_IdcrackUnitDevice_Device_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	return nil
}
