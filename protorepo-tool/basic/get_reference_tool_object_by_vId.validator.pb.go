// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_reference_tool_object_by_vId.proto

package basic

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

var _regex_GetReferenceToolObjectByVidRequest_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_GetReferenceToolObjectByVidRequest_VId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *GetReferenceToolObjectByVidRequest) Validate() error {
	if !_regex_GetReferenceToolObjectByVidRequest_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	if !_regex_GetReferenceToolObjectByVidRequest_VId.MatchString(this.VId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VId))
	}
	return nil
}
func (this *GetReferenceToolObjectByVidResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}

var _regex_GetReferenceToolObjectByVidResponse_Data_Id = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *GetReferenceToolObjectByVidResponse_Data) Validate() error {
	if !_regex_GetReferenceToolObjectByVidResponse_Data_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.Id))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *GetReferenceToolObjectByVidResponse_Data_Metadata) Validate() error {
	return nil
}
func (this *GetReferenceToolObjectByVidResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}