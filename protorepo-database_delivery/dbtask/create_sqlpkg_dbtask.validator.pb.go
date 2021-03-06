// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_sqlpkg_dbtask.proto

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

var _regex_CreateSQLPackageDBTaskRequest_PkgId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CreateSQLPackageDBTaskRequest_VersionId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateSQLPackageDBTaskRequest) Validate() error {
	if !_regex_CreateSQLPackageDBTaskRequest_PkgId.MatchString(this.PkgId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PkgId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.PkgId))
	}
	if !_regex_CreateSQLPackageDBTaskRequest_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.VersionId))
	}
	if this.TaskInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TaskInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TaskInfo", err)
		}
	}
	return nil
}

var _regex_CreateSQLPackageDBTaskRequest_TaskInfo_DbServiceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateSQLPackageDBTaskRequest_TaskInfo) Validate() error {
	if !_regex_CreateSQLPackageDBTaskRequest_TaskInfo_DbServiceId.MatchString(this.DbServiceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DbServiceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.DbServiceId))
	}
	for _, item := range this.BackupCfg {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BackupCfg", err)
			}
		}
	}
	for _, item := range this.ChangeCfg {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ChangeCfg", err)
			}
		}
	}
	for _, item := range this.CheckCfg {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CheckCfg", err)
			}
		}
	}
	return nil
}

var _regex_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_HostId = regexp.MustCompile(`^[0-9a-z]{13}$`)
var _regex_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_HostIp = regexp.MustCompile(`^((2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3})|\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$`)

func (this *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg) Validate() error {
	if !_regex_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_HostId.MatchString(this.HostId) {
		return github_com_mwitkow_go_proto_validators.FieldError("HostId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.HostId))
	}
	if !_regex_CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_HostIp.MatchString(this.HostIp) {
		return github_com_mwitkow_go_proto_validators.FieldError("HostIp", fmt.Errorf(`value '%v' must be a string conforming to regex "^((2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3})|\\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:)))(%.+)?\\s*$"`, this.HostIp))
	}
	for _, item := range this.Scripts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Scripts", err)
			}
		}
	}
	return nil
}
func (this *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts) Validate() error {
	for _, item := range this.Variables {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Variables", err)
			}
		}
	}
	return nil
}
func (this *CreateSQLPackageDBTaskRequest_TaskInfo_BackupCfg_Scripts_Variables) Validate() error {
	return nil
}

var _regex_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_DbInstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg) Validate() error {
	if !_regex_CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_DbInstanceId.MatchString(this.DbInstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DbInstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.DbInstanceId))
	}
	for _, item := range this.Scripts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Scripts", err)
			}
		}
	}
	return nil
}
func (this *CreateSQLPackageDBTaskRequest_TaskInfo_ChangeCfg_Scripts) Validate() error {
	return nil
}

var _regex_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg_DbInstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg) Validate() error {
	if !_regex_CreateSQLPackageDBTaskRequest_TaskInfo_CheckCfg_DbInstanceId.MatchString(this.DbInstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("DbInstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.DbInstanceId))
	}
	return nil
}
func (this *CreateSQLPackageDBTaskResponse) Validate() error {
	return nil
}
func (this *CreateSQLPackageDBTaskResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
