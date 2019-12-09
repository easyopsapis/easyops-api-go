// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: messsage_data.proto

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

var _regex_MessageData_CcAddr = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)
var _regex_MessageData_FromAddr = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)

func (this *MessageData) Validate() error {
	for _, item := range this.CcAddr {
		if !_regex_MessageData_CcAddr.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("CcAddr", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, item))
		}
	}
	if !_regex_MessageData_FromAddr.MatchString(this.FromAddr) {
		return github_com_mwitkow_go_proto_validators.FieldError("FromAddr", fmt.Errorf(`value '%v' must be a string conforming to regex "^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$"`, this.FromAddr))
	}
	return nil
}
