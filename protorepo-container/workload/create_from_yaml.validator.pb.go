// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_from_yaml.proto

package workload

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/container"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateFromYamlRequest_RgId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateFromYamlRequest) Validate() error {
	if !_regex_CreateFromYamlRequest_RgId.MatchString(this.RgId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RgId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.RgId))
	}
	return nil
}
func (this *CreateFromYamlResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
