// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

package mongo

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

func (this *UpdateRequest) Validate() error {
	if this.Selector != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Selector); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Selector", err)
		}
	}
	if this.Update != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Update); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Update", err)
		}
	}
	return nil
}
func (this *UpdateResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
