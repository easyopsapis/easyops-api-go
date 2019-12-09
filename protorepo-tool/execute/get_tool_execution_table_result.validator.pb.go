// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_tool_execution_table_result.proto

package execute

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

func (this *GetToolExecutionTableResultRequest) Validate() error {
	return nil
}

var _regex_GetToolExecutionTableResultResponse_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *GetToolExecutionTableResultResponse) Validate() error {
	if !_regex_GetToolExecutionTableResultResponse_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	if this.TableData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TableData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TableData", err)
		}
	}
	for _, item := range this.TableDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TableDefs", err)
			}
		}
	}
	return nil
}
func (this *GetToolExecutionTableResultResponse_TableDefs) Validate() error {
	for _, item := range this.Dimensions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dimensions", err)
			}
		}
	}
	for _, item := range this.Columns {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Columns", err)
			}
		}
	}
	return nil
}
func (this *GetToolExecutionTableResultResponse_TableDefs_Dimensions) Validate() error {
	return nil
}
func (this *GetToolExecutionTableResultResponse_TableDefs_Columns) Validate() error {
	return nil
}
func (this *GetToolExecutionTableResultResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}