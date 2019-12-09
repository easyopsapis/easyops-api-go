// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_toolkit.proto

package toolkit

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateToolkitRequest) Validate() error {
	if this.Toolkit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Toolkit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Toolkit", err)
		}
	}
	return nil
}
func (this *CreateToolkitResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}