// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cluster_info.proto

package cmdb

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

var _regex_ClusterInfo_ClusterId = regexp.MustCompile(`^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$`)

func (this *ClusterInfo) Validate() error {
	if !_regex_ClusterInfo_ClusterId.MatchString(this.ClusterId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}$"`, this.ClusterId))
	}
	return nil
}
