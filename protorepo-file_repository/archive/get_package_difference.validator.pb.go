// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_package_difference.proto

package archive

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/file_repository"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_GetPackageDifferenceRequest_PackageFrom = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_GetPackageDifferenceRequest_PackageTo = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_GetPackageDifferenceRequest_VerFrom = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_GetPackageDifferenceRequest_VerTo = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)
var _regex_GetPackageDifferenceRequest_Path = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *GetPackageDifferenceRequest) Validate() error {
	if !_regex_GetPackageDifferenceRequest_PackageFrom.MatchString(this.PackageFrom) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageFrom", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageFrom))
	}
	if !_regex_GetPackageDifferenceRequest_PackageTo.MatchString(this.PackageTo) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageTo", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.PackageTo))
	}
	if !_regex_GetPackageDifferenceRequest_VerFrom.MatchString(this.VerFrom) {
		return github_com_mwitkow_go_proto_validators.FieldError("VerFrom", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.VerFrom))
	}
	if !_regex_GetPackageDifferenceRequest_VerTo.MatchString(this.VerTo) {
		return github_com_mwitkow_go_proto_validators.FieldError("VerTo", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.VerTo))
	}
	if !_regex_GetPackageDifferenceRequest_Path.MatchString(this.Path) {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Path))
	}
	return nil
}
func (this *GetPackageDifferenceResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
