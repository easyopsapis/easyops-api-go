// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: export.proto

package info

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ExportSuiteRequest_PluginId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)

func (this *ExportSuiteRequest) Validate() error {
	if !_regex_ExportSuiteRequest_PluginId.MatchString(this.PluginId) {
		return github_com_mwitkow_go_proto_validators.FieldError("PluginId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.PluginId))
	}
	return nil
}