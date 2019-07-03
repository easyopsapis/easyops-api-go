// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user_register.proto

package user_admin

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

var _regex_UserRegisterRequest_Name = regexp.MustCompile(`^[a-zA-Z]\w{0,}$`)
var _regex_UserRegisterRequest_Email = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)

func (this *UserRegisterRequest) Validate() error {
	if !_regex_UserRegisterRequest_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z]\\w{0,}$"`, this.Name))
	}
	if !(len(this.Password) > 5) {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must length be greater than '5'`, this.Password))
	}
	if !_regex_UserRegisterRequest_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.Email))
	}
	return nil
}

var _regex_UserRegisterResponse_Name = regexp.MustCompile(`^[a-zA-Z]\w{0,}$`)
var _regex_UserRegisterResponse_Email = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)

func (this *UserRegisterResponse) Validate() error {
	if !_regex_UserRegisterResponse_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z]\\w{0,}$"`, this.Name))
	}
	if !_regex_UserRegisterResponse_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.Email))
	}
	return nil
}
func (this *UserRegisterResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
