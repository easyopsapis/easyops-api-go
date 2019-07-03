// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: diff.proto

package file_repository

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

var _regex_Diff_File = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)
var _regex_Diff_Perm = regexp.MustCompile(`^(0[0-7]{3})`)
var _regex_Diff_NewMd5 = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var _regex_Diff_OldMd5 = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func (this *Diff) Validate() error {
	if !_regex_Diff_File.MatchString(this.File) {
		return github_com_mwitkow_go_proto_validators.FieldError("File", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.File))
	}
	if !_regex_Diff_Perm.MatchString(this.Perm) {
		return github_com_mwitkow_go_proto_validators.FieldError("Perm", fmt.Errorf(`value '%v' must be a string conforming to regex "^(0[0-7]{3})"`, this.Perm))
	}
	if !_regex_Diff_NewMd5.MatchString(this.NewMd5) {
		return github_com_mwitkow_go_proto_validators.FieldError("NewMd5", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.NewMd5))
	}
	if !_regex_Diff_OldMd5.MatchString(this.OldMd5) {
		return github_com_mwitkow_go_proto_validators.FieldError("OldMd5", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{32}$"`, this.OldMd5))
	}
	return nil
}