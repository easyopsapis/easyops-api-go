// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_approve_object_list.proto

package cmdb_approve

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

func (this *GetApproveObjectListRequest) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}

var _regex_GetApproveObjectListResponse_ObjectIdList = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)

func (this *GetApproveObjectListResponse) Validate() error {
	for _, item := range this.ObjectIdList {
		if !_regex_GetApproveObjectListResponse_ObjectIdList.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("ObjectIdList", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, item))
		}
	}
	return nil
}
func (this *GetApproveObjectListResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}