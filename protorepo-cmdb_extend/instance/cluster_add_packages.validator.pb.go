// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cluster_add_packages.proto

package instance

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb_extend"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ClusterAddPackagesRequest_InstanceId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *ClusterAddPackagesRequest) Validate() error {
	if !_regex_ClusterAddPackagesRequest_InstanceId.MatchString(this.InstanceId) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.InstanceId))
	}
	for _, item := range this.Packages {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Packages", err)
			}
		}
	}
	return nil
}
func (this *ClusterAddPackagesResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
