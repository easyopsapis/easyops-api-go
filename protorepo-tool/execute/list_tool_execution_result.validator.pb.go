// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_tool_execution_result.proto

package execute

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

var _regex_ListToolExecutionResultRequest_StartTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_ListToolExecutionResultRequest_EndTime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_ListToolExecutionResultRequest_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_ListToolExecutionResultRequest_VId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *ListToolExecutionResultRequest) Validate() error {
	if !_regex_ListToolExecutionResultRequest_StartTime.MatchString(this.StartTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("StartTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.StartTime))
	}
	if !_regex_ListToolExecutionResultRequest_EndTime.MatchString(this.EndTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("EndTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.EndTime))
	}
	if !_regex_ListToolExecutionResultRequest_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	if !_regex_ListToolExecutionResultRequest_VId.MatchString(this.VId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VId))
	}
	return nil
}
func (this *ListToolExecutionResultResponse) Validate() error {
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
func (this *ListToolExecutionResultResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
