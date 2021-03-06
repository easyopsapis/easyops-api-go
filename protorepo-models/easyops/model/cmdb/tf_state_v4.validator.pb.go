// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tf_state_v4.proto

package cmdb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TFStateV4) Validate() error {
	if this.Outputs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Outputs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Outputs", err)
		}
	}
	for _, item := range this.Resources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
			}
		}
	}
	return nil
}
func (this *TFStateV4_Resources) Validate() error {
	for _, item := range this.Instances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Instances", err)
			}
		}
	}
	return nil
}
