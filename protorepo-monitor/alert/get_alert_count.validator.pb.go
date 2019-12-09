// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_alert_count.proto

package alert

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

func (this *GetAlertCountRequest) Validate() error {
	if !(this.St > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("St", fmt.Errorf(`value '%v' must be greater than '0'`, this.St))
	}
	if !(this.Et > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Et", fmt.Errorf(`value '%v' must be greater than '0'`, this.Et))
	}
	return nil
}
func (this *GetAlertCountResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetAlertCountResponse_Data) Validate() error {
	return nil
}
func (this *GetAlertCountResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
