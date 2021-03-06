// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package cmdb

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

var _regex_User_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_User_Name = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_User_Nickname = regexp.MustCompile(`.*`)
var _regex_User_UserEmail = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)

func (this *User) Validate() error {
	if !_regex_User_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if !_regex_User_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Name))
	}
	if !_regex_User_Nickname.MatchString(this.Nickname) {
		return github_com_mwitkow_go_proto_validators.FieldError("Nickname", fmt.Errorf(`value '%v' must be a string conforming to regex ".*"`, this.Nickname))
	}
	if !_regex_User_UserEmail.MatchString(this.UserEmail) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserEmail", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.UserEmail))
	}
	return nil
}
