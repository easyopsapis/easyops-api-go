// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message_receiver.proto

package msgsender

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

var _regex_MessageReceiver_EmailAddr = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)

func (this *MessageReceiver) Validate() error {
	if !_regex_MessageReceiver_EmailAddr.MatchString(this.EmailAddr) {
		return github_com_mwitkow_go_proto_validators.FieldError("EmailAddr", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.EmailAddr))
	}
	return nil
}
