// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_archive_list.proto

package archive

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

var _regex_ListdirRequest_Path = regexp.MustCompile(`^\/(([^\/]+)(\/[^\/]+)*)?$`)
var _regex_ListdirRequest_PackageId = regexp.MustCompile(`^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$`)
var _regex_ListdirRequest_VersionId = regexp.MustCompile(`^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$`)

func (this *ListdirRequest) Validate() error {
	if !_regex_ListdirRequest_Path.MatchString(this.Path) {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\/(([^\\/]+)(\\/[^\\/]+)*)?$"`, this.Path))
	}
	if !_regex_ListdirRequest_PackageId.MatchString(this.PackageId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PackageId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$"`, this.PackageId))
	}
	if !_regex_ListdirRequest_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[?a-zA-Z0-9]{8}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{4}-[?a-zA-Z0-9]{12}$"`, this.VersionId))
	}
	return nil
}
func (this *ListdirResponse) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}

var _regex_ListdirResponse_List_Perm = regexp.MustCompile(`^(0[0-7]{3})`)
var _regex_ListdirResponse_List_Mtime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_ListdirResponse_List_Ctime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_ListdirResponse_List_Md5 = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *ListdirResponse_List) Validate() error {
	if !_regex_ListdirResponse_List_Perm.MatchString(this.Perm) {
		return github_com_mwitkow_go_proto_validators.FieldError("Perm", fmt.Errorf(`value '%v' must be a string conforming to regex "^(0[0-7]{3})"`, this.Perm))
	}
	if !_regex_ListdirResponse_List_Mtime.MatchString(this.Mtime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Mtime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Mtime))
	}
	if !_regex_ListdirResponse_List_Ctime.MatchString(this.Ctime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ctime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Ctime))
	}
	if !_regex_ListdirResponse_List_Md5.MatchString(this.Md5) {
		return github_com_mwitkow_go_proto_validators.FieldError("Md5", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.Md5))
	}
	return nil
}
func (this *ListdirResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
