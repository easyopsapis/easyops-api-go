// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool_task.proto

package tool

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

var _regex_ToolTask_Username = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_ToolTask_Agents = regexp.MustCompile(`^((2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3})|\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$`)
var _regex_ToolTask_ExecUser = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_ToolTask_VId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_ToolTask_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *ToolTask) Validate() error {
	if !_regex_ToolTask_Username.MatchString(this.Username) {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Username))
	}
	if this.Inputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Inputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Inputs", err)
		}
	}
	if this.ExtraDetail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExtraDetail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExtraDetail", err)
		}
	}
	if this.ToolEnv != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToolEnv); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToolEnv", err)
		}
	}
	if this.Outputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Outputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Outputs", err)
		}
	}
	if this.OutViewData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OutViewData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OutViewData", err)
		}
	}
	if this.AgentData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AgentData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AgentData", err)
		}
	}
	for _, item := range this.Agents {
		if !_regex_ToolTask_Agents.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Agents", fmt.Errorf(`value '%v' must be a string conforming to regex "^((2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3})|\\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:)))(%.+)?\\s*$"`, item))
		}
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	if this.Msg != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Msg); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Msg", err)
		}
	}
	if this.EndTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndTime", err)
		}
	}
	if this.ExitStatus != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExitStatus); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExitStatus", err)
		}
	}
	if this.SysStatus != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SysStatus); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SysStatus", err)
		}
	}
	for _, item := range this.ToolOutputsData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ToolOutputsData", err)
			}
		}
	}
	if this.TableData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TableData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TableData", err)
		}
	}
	if this.ToolData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToolData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToolData", err)
		}
	}
	if this.BatchStrategy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BatchStrategy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BatchStrategy", err)
		}
	}
	if !_regex_ToolTask_ExecUser.MatchString(this.ExecUser) {
		return github_com_mwitkow_go_proto_validators.FieldError("ExecUser", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.ExecUser))
	}
	if !_regex_ToolTask_VId.MatchString(this.VId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VId))
	}
	if !_regex_ToolTask_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	for _, item := range this.OutputDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OutputDefs", err)
			}
		}
	}
	for _, item := range this.TableDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TableDefs", err)
			}
		}
	}
	if this.ToolOutputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToolOutputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToolOutputs", err)
		}
	}
	return nil
}
func (this *ToolTask_ExtraDetail) Validate() error {
	if this.Callback != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Callback); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Callback", err)
		}
	}
	if this.ToolEnv != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToolEnv); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToolEnv", err)
		}
	}
	if this.ToolOutputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ToolOutputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ToolOutputs", err)
		}
	}
	for _, item := range this.TableDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TableDefs", err)
			}
		}
	}
	for _, item := range this.OutputDefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OutputDefs", err)
			}
		}
	}
	return nil
}
func (this *ToolTask_ExtraDetail_TableDefs) Validate() error {
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
func (this *ToolTask_ExtraDetail_TableDefs_Dimensions) Validate() error {
	return nil
}
func (this *ToolTask_ExtraDetail_TableDefs_Columns) Validate() error {
	return nil
}
func (this *ToolTask_ExtraDetail_OutputDefs) Validate() error {
	return nil
}
func (this *ToolTask_BatchStrategy) Validate() error {
	return nil
}
func (this *ToolTask_OutputDefs) Validate() error {
	return nil
}
func (this *ToolTask_TableDefs) Validate() error {
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
func (this *ToolTask_TableDefs_Dimensions) Validate() error {
	return nil
}
func (this *ToolTask_TableDefs_Columns) Validate() error {
	return nil
}
