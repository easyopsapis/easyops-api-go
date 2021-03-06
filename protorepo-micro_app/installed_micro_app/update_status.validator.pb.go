// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update_status.proto

package installed_micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UpdateInstalledMicroAppStatusRequest) Validate() error {
	if this.MicroApp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MicroApp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MicroApp", err)
		}
	}
	return nil
}
func (this *UpdateInstalledMicroAppStatusRequest_MicroApp) Validate() error {
	return nil
}
func (this *UpdateInstalledMicroAppStatusResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
