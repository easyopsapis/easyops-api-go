// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_service_topology_task.proto

package service_manage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateServiceTopologyTaskRequest) Validate() error {
	if this.Switch != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Switch); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Switch", err)
		}
	}
	return nil
}
func (this *CreateServiceTopologyTaskRequest_Switch) Validate() error {
	return nil
}
func (this *CreateServiceTopologyTaskResponse) Validate() error {
	return nil
}
func (this *CreateServiceTopologyTaskResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
