// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: val.proto

package inspection

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

func (this *InspectionVal) Validate() error {
	if !(this.Weight > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Weight", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Weight))
	}
	if !(this.Weight < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("Weight", fmt.Errorf(`value '%v' must be less than '101'`, this.Weight))
	}
	for _, item := range this.Conditions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Conditions", err)
			}
		}
	}
	return nil
}