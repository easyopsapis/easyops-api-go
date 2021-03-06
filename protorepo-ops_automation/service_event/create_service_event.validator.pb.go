// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_service_event.proto

package service_event

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

func (this *CreateServiceEventRequest) Validate() error {
	if !(len(this.System) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("System", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.System))
	}
	if !(len(this.System) < 256) {
		return github_com_mwitkow_go_proto_validators.FieldError("System", fmt.Errorf(`value '%v' must have a length smaller than '256'`, this.System))
	}
	if !(len(this.Topic) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Topic", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Topic))
	}
	if !(len(this.Topic) < 256) {
		return github_com_mwitkow_go_proto_validators.FieldError("Topic", fmt.Errorf(`value '%v' must have a length smaller than '256'`, this.Topic))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateServiceEventRequest_Data) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *CreateServiceEventResponse) Validate() error {
	return nil
}
func (this *CreateServiceEventResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
