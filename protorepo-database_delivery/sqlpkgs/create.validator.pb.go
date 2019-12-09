// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package sqlpkgs

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/database_delivery"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateSQLPackageRequest) Validate() error {
	if this.CreateSqlpkg != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreateSqlpkg); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreateSqlpkg", err)
		}
	}
	return nil
}
func (this *CreateSQLPackageRequest_CreateSqlpkg) Validate() error {
	return nil
}

var _regex_CreateSQLPackageResponse_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CreateSQLPackageResponse_Creator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *CreateSQLPackageResponse) Validate() error {
	for _, item := range this.VersionList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("VersionList", err)
			}
		}
	}
	if !_regex_CreateSQLPackageResponse_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	if !_regex_CreateSQLPackageResponse_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Creator))
	}
	return nil
}
func (this *CreateSQLPackageResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
