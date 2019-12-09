// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool_executon_callback.proto

package execute

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/tool"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ToolExecutionCallbackRequest) Validate() error {
	if this.ExtraInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExtraInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExtraInfo", err)
		}
	}
	return nil
}
func (this *ToolExecutionCallbackResponse) Validate() error {
	if this.ExtraInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExtraInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExtraInfo", err)
		}
	}
	return nil
}
func (this *ToolExecutionCallbackResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}