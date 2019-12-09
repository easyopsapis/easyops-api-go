// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_quota.proto

package container

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

var _regex_ResourceQuota_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_ResourceQuota_ResourceName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
var _regex_ResourceQuota_Creator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *ResourceQuota) Validate() error {
	if !_regex_ResourceQuota_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !_regex_ResourceQuota_ResourceName.MatchString(this.ResourceName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ResourceName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"`, this.ResourceName))
	}
	if this.Hard != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hard); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hard", err)
		}
	}
	if this.Used != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Used); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Used", err)
		}
	}
	if !_regex_ResourceQuota_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Creator))
	}
	return nil
}
