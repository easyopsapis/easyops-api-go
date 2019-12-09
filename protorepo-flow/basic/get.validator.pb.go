// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get.proto

package basic

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/flow"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_GetFlowInfoRequest_FlowId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *GetFlowInfoRequest) Validate() error {
	if !_regex_GetFlowInfoRequest_FlowId.MatchString(this.FlowId) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.FlowId))
	}
	return nil
}
func (this *GetFlowInfoResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}