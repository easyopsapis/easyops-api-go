// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: view.proto

package topology

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

var _regex_View_Creator = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)
var _regex_View_Modifier = regexp.MustCompile(`^[A-Za-z_-]\w{3,64}$`)

func (this *View) Validate() error {
	if !_regex_View_Creator.MatchString(this.Creator) {
		return github_com_mwitkow_go_proto_validators.FieldError("Creator", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, this.Creator))
	}
	if !_regex_View_Modifier.MatchString(this.Modifier) {
		return github_com_mwitkow_go_proto_validators.FieldError("Modifier", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z_-]\\w{3,64}$"`, this.Modifier))
	}
	if this.RootNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RootNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RootNode", err)
		}
	}
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	for _, item := range this.Links {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
			}
		}
	}
	for _, item := range this.Areas {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Areas", err)
			}
		}
	}
	for _, item := range this.Notes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notes", err)
			}
		}
	}
	if this.Diff != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Diff); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Diff", err)
		}
	}
	return nil
}
func (this *View_Diff) Validate() error {
	for _, item := range this.AddNodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AddNodes", err)
			}
		}
	}
	for _, item := range this.RemoveNodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RemoveNodes", err)
			}
		}
	}
	for _, item := range this.AddLinks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AddLinks", err)
			}
		}
	}
	for _, item := range this.RemoveLinks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RemoveLinks", err)
			}
		}
	}
	return nil
}
