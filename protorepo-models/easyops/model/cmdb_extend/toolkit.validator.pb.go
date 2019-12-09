// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: toolkit.proto

package cmdb_extend

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

func (this *Toolkit) Validate() error {
	if this.XId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.XId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("XId", err)
		}
	}
	if this.Style != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Style); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Style", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *Toolkit_Style) Validate() error {
	return nil
}
