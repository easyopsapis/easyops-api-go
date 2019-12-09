// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relation_max_check.proto

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

var _regex_RelationMaxCheckRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_RelationMaxCheckRequest_RelationId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]*$`)
var _regex_RelationMaxCheckRequest_RelationSideId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *RelationMaxCheckRequest) Validate() error {
	if !_regex_RelationMaxCheckRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if !_regex_RelationMaxCheckRequest_RelationId.MatchString(this.RelationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RelationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]*$"`, this.RelationId))
	}
	if !_regex_RelationMaxCheckRequest_RelationSideId.MatchString(this.RelationSideId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RelationSideId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.RelationSideId))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Sort != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sort); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sort", err)
		}
	}
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *RelationMaxCheckResponse) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}

var _regex_RelationMaxCheckResponse_List_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *RelationMaxCheckResponse_List) Validate() error {
	if !_regex_RelationMaxCheckResponse_List_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	return nil
}
func (this *RelationMaxCheckResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
