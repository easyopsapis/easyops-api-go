// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package service

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateRequest_RgId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CreateRequest_NamespaceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CreateRequest_WorkloadId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateRequest) Validate() error {
	if !_regex_CreateRequest_RgId.MatchString(this.RgId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RgId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.RgId))
	}
	if !_regex_CreateRequest_NamespaceId.MatchString(this.NamespaceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("NamespaceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.NamespaceId))
	}
	if !_regex_CreateRequest_WorkloadId.MatchString(this.WorkloadId) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkloadId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.WorkloadId))
	}
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	return nil
}
func (this *CreateResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
