// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: retry_sqlpkg_dbtask.proto

package dbtask

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

var _regex_RetrySQLPackageDBTaskRequest_DbTaskId = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func (this *RetrySQLPackageDBTaskRequest) Validate() error {
	if !_regex_RetrySQLPackageDBTaskRequest_DbTaskId.MatchString(this.DbTaskId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DbTaskId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-fA-F]{24}$"`, this.DbTaskId))
	}
	return nil
}
func (this *RetrySQLPackageDBTaskResponse) Validate() error {
	return nil
}
func (this *RetrySQLPackageDBTaskResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}