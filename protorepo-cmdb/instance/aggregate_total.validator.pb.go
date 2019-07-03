// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregate_total.proto

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

var _regex_AggregateCountRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)
var _regex_AggregateCountRequest_AttrId = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)

func (this *AggregateCountRequest) Validate() error {
	if !_regex_AggregateCountRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, this.ObjectId))
	}
	if !_regex_AggregateCountRequest_AttrId.MatchString(this.AttrId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AttrId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, this.AttrId))
	}
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *AggregateCountResponse) Validate() error {
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
func (this *AggregateCountResponse_List) Validate() error {
	if this.Attr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attr", err)
		}
	}
	return nil
}
func (this *AggregateCountResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}