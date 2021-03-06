// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: info.proto

package inspection

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

var _regex_InspectionInfo_Id = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_InspectionInfo_ObjectId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_InspectionInfo_Keys = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_InspectionInfo_RelationIdWithHost = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]*$|^$`)
var _regex_InspectionInfo_Ctime = regexp.MustCompile(`^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$`)

func (this *InspectionInfo) Validate() error {
	if !_regex_InspectionInfo_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.Id))
	}
	if !(this.Index > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Index", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Index))
	}
	if !_regex_InspectionInfo_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.ObjectId))
	}
	for _, item := range this.Keys {
		if !_regex_InspectionInfo_Keys.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Keys", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, item))
		}
	}
	if !_regex_InspectionInfo_RelationIdWithHost.MatchString(this.RelationIdWithHost) {
		return github_com_mwitkow_go_proto_validators.FieldError("RelationIdWithHost", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]*$|^$"`, this.RelationIdWithHost))
	}
	if !_regex_InspectionInfo_Ctime.MatchString(this.Ctime) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ctime", fmt.Errorf(`value '%v' must be a string conforming to regex "^((?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[0-1]|0[1-9]|[1-2][0-9])[T ](2[0-3]|[0-1][0-9]):([0-5][0-9]):[0-5][0-9](\\.[0-9]+)?(Z|[+-](?:2[0-3]|[0-1][0-9]):[0-5][0-9])?$"`, this.Ctime))
	}
	return nil
}
