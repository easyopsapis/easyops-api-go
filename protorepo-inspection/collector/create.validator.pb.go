// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package collector

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/inspection"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateCollectorRequest_PluginId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *CreateCollectorRequest) Validate() error {
	if !_regex_CreateCollectorRequest_PluginId.MatchString(this.PluginId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PluginId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.PluginId))
	}
	for _, item := range this.Args {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Args", err)
			}
		}
	}
	return nil
}

var _regex_CreateCollectorResponse_CollectorId = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *CreateCollectorResponse) Validate() error {
	if !_regex_CreateCollectorResponse_CollectorId.MatchString(this.CollectorId) {
		return github_com_mwitkow_go_proto_validators.FieldError("CollectorId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.CollectorId))
	}
	return nil
}
func (this *CreateCollectorResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
