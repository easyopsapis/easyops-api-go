// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete.proto

package pipeline

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

var _regex_DeletePipelineRequest_ProjectId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_DeletePipelineRequest_PipelineId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *DeletePipelineRequest) Validate() error {
	if !_regex_DeletePipelineRequest_ProjectId.MatchString(this.ProjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ProjectId))
	}
	if !_regex_DeletePipelineRequest_PipelineId.MatchString(this.PipelineId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PipelineId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.PipelineId))
	}
	return nil
}
func (this *DeletePipelineResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
