// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: retry.proto

package build

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_RetryRequest_BuildId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *RetryRequest) Validate() error {
	if !_regex_RetryRequest_BuildId.MatchString(this.BuildId) {
		return github_com_mwitkow_go_proto_validators.FieldError("BuildId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.BuildId))
	}
	return nil
}

var _regex_RetryResponse_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *RetryResponse) Validate() error {
	if !_regex_RetryResponse_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	return nil
}
func (this *RetryResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}