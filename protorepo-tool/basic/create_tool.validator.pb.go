// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_tool.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateToolResponse_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_CreateToolResponse_VId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *CreateToolResponse) Validate() error {
	if !_regex_CreateToolResponse_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	if !_regex_CreateToolResponse_VId.MatchString(this.VId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VId))
	}
	return nil
}
func (this *CreateToolResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
