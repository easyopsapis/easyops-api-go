// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: debug_collection_config_callback.proto

package collection_config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DebugCollectionConfigCallbackRequest) Validate() error {
	if !(this.Code > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Code))
	}
	if !(len(this.TaskId) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.TaskId))
	}
	for _, item := range this.TargetsLog {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TargetsLog", err)
			}
		}
	}
	return nil
}
func (this *DebugCollectionConfigCallbackRequest_TargetsLog) Validate() error {
	return nil
}
func (this *DebugCollectionConfigCallbackResponse) Validate() error {
	return nil
}
func (this *DebugCollectionConfigCallbackResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}