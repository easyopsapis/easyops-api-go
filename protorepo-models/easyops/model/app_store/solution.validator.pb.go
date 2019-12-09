// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: solution.proto

package app_store

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

var _regex_Solution_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *Solution) Validate() error {
	if !_regex_Solution_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if this.Icon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Icon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Icon", err)
		}
	}
	return nil
}

var _regex_Solution_Icon_Icon = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *Solution_Icon) Validate() error {
	if !_regex_Solution_Icon_Icon.MatchString(this.Icon) {
		return github_com_mwitkow_go_proto_validators.FieldError("Icon", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Icon))
	}
	return nil
}
