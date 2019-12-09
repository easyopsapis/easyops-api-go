// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: search_with_uniq_key.proto

package installed_micro_app

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

func (this *SearchInstalledMicroAppWithUniqKeyRequest) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *SearchInstalledMicroAppWithUniqKeyRequest_Query) Validate() error {
	return nil
}
func (this *SearchInstalledMicroAppWithUniqKeyResponse) Validate() error {
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

var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_CurrentVersion = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_AppVersion = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Homepage = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Owner = regexp.MustCompile(`^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Ctime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)
var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Mtime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)

func (this *SearchInstalledMicroAppWithUniqKeyResponse_List) Validate() error {
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	if this.Icons != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Icons); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Icons", err)
		}
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_CurrentVersion.MatchString(this.CurrentVersion) {
		return github_com_mwitkow_go_proto_validators.FieldError("CurrentVersion", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d+\\.\\d+\\.\\d+$"`, this.CurrentVersion))
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_AppVersion.MatchString(this.AppVersion) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppVersion", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d+\\.\\d+\\.\\d+$"`, this.AppVersion))
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Homepage.MatchString(this.Homepage) {
		return github_com_mwitkow_go_proto_validators.FieldError("Homepage", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Homepage))
	}
	if this.ClonedFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ClonedFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ClonedFrom", err)
		}
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Owner.MatchString(this.Owner) {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9][.a-zA-Z0-9_-]{2,31}$"`, this.Owner))
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Ctime.MatchString(this.Ctime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ctime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Ctime))
	}
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Mtime.MatchString(this.Mtime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Mtime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Mtime))
	}
	if this.MenuIcon != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MenuIcon); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MenuIcon", err)
		}
	}
	return nil
}

var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Icons_Large = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *SearchInstalledMicroAppWithUniqKeyResponse_List_Icons) Validate() error {
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_Icons_Large.MatchString(this.Large) {
		return github_com_mwitkow_go_proto_validators.FieldError("Large", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Large))
	}
	return nil
}

var _regex_SearchInstalledMicroAppWithUniqKeyResponse_List_ClonedFrom_Version = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func (this *SearchInstalledMicroAppWithUniqKeyResponse_List_ClonedFrom) Validate() error {
	if !_regex_SearchInstalledMicroAppWithUniqKeyResponse_List_ClonedFrom_Version.MatchString(this.Version) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d+\\.\\d+\\.\\d+$"`, this.Version))
	}
	return nil
}
func (this *SearchInstalledMicroAppWithUniqKeyResponse_List_MenuIcon) Validate() error {
	return nil
}
func (this *SearchInstalledMicroAppWithUniqKeyResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
