// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_instance_relation_snapshot.proto

package instance_relation

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

var _regex_InstanceRelationSnapshotRequest_RelationId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]*$`)

func (this *InstanceRelationSnapshotRequest) Validate() error {
	if !_regex_InstanceRelationSnapshotRequest_RelationId.MatchString(this.RelationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RelationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]*$"`, this.RelationId))
	}
	return nil
}

var _regex_InstanceRelationSnapshotResponse_LeftInstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_InstanceRelationSnapshotResponse_RightInstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *InstanceRelationSnapshotResponse) Validate() error {
	if this.XLeftInstance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.XLeftInstance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("XLeftInstance", err)
		}
	}
	if this.XRightInstance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.XRightInstance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("XRightInstance", err)
		}
	}
	if !_regex_InstanceRelationSnapshotResponse_LeftInstanceId.MatchString(this.LeftInstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("LeftInstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.LeftInstanceId))
	}
	if !_regex_InstanceRelationSnapshotResponse_RightInstanceId.MatchString(this.RightInstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RightInstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.RightInstanceId))
	}
	return nil
}

var _regex_InstanceRelationSnapshotResponse_XLeftInstance_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_InstanceRelationSnapshotResponse_XLeftInstance_XObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *InstanceRelationSnapshotResponse_XLeftInstance) Validate() error {
	if !_regex_InstanceRelationSnapshotResponse_XLeftInstance_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !_regex_InstanceRelationSnapshotResponse_XLeftInstance_XObjectId.MatchString(this.XObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("XObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.XObjectId))
	}
	return nil
}

var _regex_InstanceRelationSnapshotResponse_XRightInstance_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_InstanceRelationSnapshotResponse_XRightInstance_XObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *InstanceRelationSnapshotResponse_XRightInstance) Validate() error {
	if !_regex_InstanceRelationSnapshotResponse_XRightInstance_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !_regex_InstanceRelationSnapshotResponse_XRightInstance_XObjectId.MatchString(this.XObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("XObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.XObjectId))
	}
	return nil
}
func (this *InstanceRelationSnapshotResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}