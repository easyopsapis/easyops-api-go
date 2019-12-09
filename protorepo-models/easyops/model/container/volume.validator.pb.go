// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: volume.proto

package container

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

var _regex_Volume_Name = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)

func (this *Volume) Validate() error {
	if !_regex_Volume_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"`, this.Name))
	}
	if this.HostPath != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HostPath); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HostPath", err)
		}
	}
	if this.EmptyDir != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EmptyDir); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EmptyDir", err)
		}
	}
	if this.Secret != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.ConfigMap != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConfigMap); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConfigMap", err)
		}
	}
	if this.PersistentVolumeClaim != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PersistentVolumeClaim); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PersistentVolumeClaim", err)
		}
	}
	return nil
}

var _regex_Volume_HostPath_Path = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *Volume_HostPath) Validate() error {
	if !_regex_Volume_HostPath_Path.MatchString(this.Path) {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Path))
	}
	return nil
}
func (this *Volume_EmptyDir) Validate() error {
	return nil
}
func (this *Volume_Secret) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *Volume_ConfigMap) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *Volume_PersistentVolumeClaim) Validate() error {
	return nil
}
