// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mail_info.proto

package ops_automation

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

var _regex_MailInfo_Notifiers = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)

func (this *MailInfo) Validate() error {
	for _, item := range this.Notifiers {
		if !_regex_MailInfo_Notifiers.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Notifiers", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, item))
		}
	}
	return nil
}
