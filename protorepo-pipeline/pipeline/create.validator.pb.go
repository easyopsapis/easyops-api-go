// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package pipeline

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/pipeline"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateRequest_ProjectId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateRequest) Validate() error {
	if !_regex_CreateRequest_ProjectId.MatchString(this.ProjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ProjectId))
	}
	if this.Pipeline != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pipeline); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pipeline", err)
		}
	}
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *CreateResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
