// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: installed_micro_app.proto

package micro_app

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

var _regex_InstalledMicroApp_Name = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)
var _regex_InstalledMicroApp_AppId = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)
var _regex_InstalledMicroApp_CurrentVersion = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func (this *InstalledMicroApp) Validate() error {
	if !_regex_InstalledMicroApp_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, this.Name))
	}
	if !_regex_InstalledMicroApp_AppId.MatchString(this.AppId) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, this.AppId))
	}
	if this.Icons != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Icons); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Icons", err)
		}
	}
	if !_regex_InstalledMicroApp_CurrentVersion.MatchString(this.CurrentVersion) {
		return github_com_mwitkow_go_proto_validators.FieldError("CurrentVersion", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d+\\.\\d+\\.\\d+$"`, this.CurrentVersion))
	}
	return nil
}

var _regex_InstalledMicroApp_Icons_Large = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *InstalledMicroApp_Icons) Validate() error {
	if !_regex_InstalledMicroApp_Icons_Large.MatchString(this.Large) {
		return github_com_mwitkow_go_proto_validators.FieldError("Large", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Large))
	}
	return nil
}
