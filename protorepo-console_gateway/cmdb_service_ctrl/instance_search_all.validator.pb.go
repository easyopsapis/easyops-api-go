// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance_search_all.proto

package cmdb_service_ctrl

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

var _regex_InstanceSearchAllRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *InstanceSearchAllRequest) Validate() error {
	if !_regex_InstanceSearchAllRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Fields != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fields); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
		}
	}
	if this.Sort != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sort); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sort", err)
		}
	}
	return nil
}
func (this *InstanceSearchAllResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *InstanceSearchAllResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
