// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metadata.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Metadata_CreationTimestamp = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)

func (this *Metadata) Validate() error {
	if !_regex_Metadata_CreationTimestamp.MatchString(this.CreationTimestamp) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreationTimestamp", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.CreationTimestamp))
	}
	if this.Annotations != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Annotations); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Annotations", err)
		}
	}
	if this.Labels != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Labels); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Labels", err)
		}
	}
	return nil
}
