// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_layout.proto

package idcrack

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

var _regex_UpdateIDCRackLayoutRequest_IdcrackId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *UpdateIDCRackLayoutRequest) Validate() error {
	if !_regex_UpdateIDCRackLayoutRequest_IdcrackId.MatchString(this.IdcrackId) {
		return github_com_mwitkow_go_proto_validators.FieldError("IdcrackId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.IdcrackId))
	}
	for _, item := range this.Layout {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Layout", err)
			}
		}
	}
	return nil
}
func (this *UpdateIDCRackLayoutRequest_Layout) Validate() error {
	return nil
}
func (this *UpdateIDCRackLayoutResponse) Validate() error {
	return nil
}
func (this *UpdateIDCRackLayoutResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
