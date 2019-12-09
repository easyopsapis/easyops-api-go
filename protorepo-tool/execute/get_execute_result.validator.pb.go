// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_execute_result.proto

package execute

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

var _regex_GetExecuteResultRequest_TaskId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *GetExecuteResultRequest) Validate() error {
	if !_regex_GetExecuteResultRequest_TaskId.MatchString(this.TaskId) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.TaskId))
	}
	return nil
}
func (this *GetExecuteResultResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}