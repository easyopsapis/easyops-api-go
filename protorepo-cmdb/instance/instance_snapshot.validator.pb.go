// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance_snapshot.proto

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

var _regex_InstanceSnapshotRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_InstanceSnapshotRequest_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *InstanceSnapshotRequest) Validate() error {
	if !_regex_InstanceSnapshotRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if !_regex_InstanceSnapshotRequest_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !(this.XVersion > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("XVersion", fmt.Errorf(`value '%v' must be greater than '-1'`, this.XVersion))
	}
	if !(this.XTs > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("XTs", fmt.Errorf(`value '%v' must be greater than '-1'`, this.XTs))
	}
	return nil
}
func (this *InstanceSnapshotResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
