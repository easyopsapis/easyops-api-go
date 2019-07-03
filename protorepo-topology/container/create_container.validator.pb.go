// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_container.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/topology"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateContainerRequest) Validate() error {
	if this.Property != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Property); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Property", err)
		}
	}
	if this.Style != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Style); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Style", err)
		}
	}
	return nil
}
func (this *CreateContainerRequest_Style) Validate() error {
	return nil
}
func (this *CreateContainerResponse) Validate() error {
	return nil
}
func (this *CreateContainerResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
