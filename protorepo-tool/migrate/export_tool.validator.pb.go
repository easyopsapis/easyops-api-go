// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: export_tool.proto

package migrate

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

var _regex_ExportToolRequest_VersionId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_ExportToolRequest_ToolId = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *ExportToolRequest) Validate() error {
	if !_regex_ExportToolRequest_VersionId.MatchString(this.VersionId) {
		return github_com_mwitkow_go_proto_validators.FieldError("VersionId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.VersionId))
	}
	if !_regex_ExportToolRequest_ToolId.MatchString(this.ToolId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ToolId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.ToolId))
	}
	return nil
}
