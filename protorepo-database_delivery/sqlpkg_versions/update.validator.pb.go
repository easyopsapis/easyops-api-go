// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update.proto

package sqlpkg_versions

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

var _regex_UpdateSQLPackageVersionRequest_PkgId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_UpdateSQLPackageVersionRequest_VersionId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *UpdateSQLPackageVersionRequest) Validate() error {
	if !_regex_UpdateSQLPackageVersionRequest_PkgId.MatchString(this.PkgId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PkgId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.PkgId))
	}
	if !_regex_UpdateSQLPackageVersionRequest_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.VersionId))
	}
	if this.UpdateSqlPkgVersion != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateSqlPkgVersion); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateSqlPkgVersion", err)
		}
	}
	return nil
}
func (this *UpdateSQLPackageVersionRequest_UpdateSqlPkgVersion) Validate() error {
	return nil
}

var _regex_UpdateSQLPackageVersionResponse_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_UpdateSQLPackageVersionResponse_Creator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *UpdateSQLPackageVersionResponse) Validate() error {
	if this.SqlPackage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SqlPackage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SqlPackage", err)
		}
	}
	if !_regex_UpdateSQLPackageVersionResponse_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	if !_regex_UpdateSQLPackageVersionResponse_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Creator))
	}
	return nil
}

var _regex_UpdateSQLPackageVersionResponse_SqlPackage_Id = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_UpdateSQLPackageVersionResponse_SqlPackage_Creator = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)

func (this *UpdateSQLPackageVersionResponse_SqlPackage) Validate() error {
	if !_regex_UpdateSQLPackageVersionResponse_SqlPackage_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.Id))
	}
	if !_regex_UpdateSQLPackageVersionResponse_SqlPackage_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Creator))
	}
	return nil
}
func (this *UpdateSQLPackageVersionResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
