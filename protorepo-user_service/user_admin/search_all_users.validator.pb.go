// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_all_users.proto

package user_admin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SearchAllUsersInfoRequest) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Fields != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fields); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
		}
	}
	if this.Sort != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sort); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sort", err)
		}
	}
	return nil
}
func (this *SearchAllUsersInfoResponse) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *SearchAllUsersInfoResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
