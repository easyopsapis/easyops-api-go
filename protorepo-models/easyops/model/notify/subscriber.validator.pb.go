// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscriber.proto

package notify

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

var _regex_Subscriber_Callback = regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)

func (this *Subscriber) Validate() error {
	if !_regex_Subscriber_Callback.MatchString(this.Callback) {
		return github_com_mwitkow_go_proto_validators.FieldError("Callback", fmt.Errorf(`value '%v' must be a string conforming to regex "^https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)$"`, this.Callback))
	}
	for _, item := range this.SubscribeInfo {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SubscribeInfo", err)
			}
		}
	}
	return nil
}
