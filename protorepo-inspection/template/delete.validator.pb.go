// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package template

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

var _regex_DeleteTemplateRequest_PluginId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_DeleteTemplateRequest_TemplateId = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *DeleteTemplateRequest) Validate() error {
	if !_regex_DeleteTemplateRequest_PluginId.MatchString(this.PluginId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PluginId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.PluginId))
	}
	if !_regex_DeleteTemplateRequest_TemplateId.MatchString(this.TemplateId) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.TemplateId))
	}
	return nil
}
func (this *DeleteTemplateResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}