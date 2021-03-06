// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: send_mail_by_instanceId.proto

package email

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

var _regex_SendMailByInstanceIdRequest_SendToUser = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_SendMailByInstanceIdRequest_SendToGroup = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_SendMailByInstanceIdRequest_CcToUser = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_SendMailByInstanceIdRequest_CcToGroup = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *SendMailByInstanceIdRequest) Validate() error {
	for _, item := range this.SendToUser {
		if !_regex_SendMailByInstanceIdRequest_SendToUser.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("SendToUser", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, item))
		}
	}
	for _, item := range this.SendToGroup {
		if !_regex_SendMailByInstanceIdRequest_SendToGroup.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("SendToGroup", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, item))
		}
	}
	for _, item := range this.CcToUser {
		if !_regex_SendMailByInstanceIdRequest_CcToUser.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("CcToUser", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, item))
		}
	}
	for _, item := range this.CcToGroup {
		if !_regex_SendMailByInstanceIdRequest_CcToGroup.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("CcToGroup", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, item))
		}
	}
	return nil
}
func (this *SendMailByInstanceIdResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
