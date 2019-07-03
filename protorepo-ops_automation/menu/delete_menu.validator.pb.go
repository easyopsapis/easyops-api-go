// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_menu.proto

package menu

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeleteMenuRequest) Validate() error {
	if !(len(this.MenusId) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("MenusId", fmt.Errorf(`value '%v' must length be greater than '1'`, this.MenusId))
	}
	if !(len(this.MenusId) < 34) {
		return github_com_mwitkow_go_proto_validators.FieldError("MenusId", fmt.Errorf(`value '%v' must length be less than '34'`, this.MenusId))
	}
	return nil
}
func (this *DeleteMenuResponse) Validate() error {
	return nil
}
func (this *DeleteMenuResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
