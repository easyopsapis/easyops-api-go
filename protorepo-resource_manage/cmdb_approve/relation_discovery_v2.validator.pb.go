// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relation_discovery_v2.proto

package cmdb_approve

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/cmdb"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_RelationDiscoveryV2Request_RelationId = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]*$`)

func (this *RelationDiscoveryV2Request) Validate() error {
	if !_regex_RelationDiscoveryV2Request_RelationId.MatchString(this.RelationId) {
		return github_com_mwitkow_go_proto_validators.FieldError("RelationId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]*$"`, this.RelationId))
	}
	if this.Match != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Match); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Match", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *RelationDiscoveryV2Request_Match) Validate() error {
	return nil
}
func (this *RelationDiscoveryV2Request_Data) Validate() error {
	if this.LeftInstance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LeftInstance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LeftInstance", err)
		}
	}
	if this.RightInstance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RightInstance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RightInstance", err)
		}
	}
	return nil
}
func (this *RelationDiscoveryV2Response) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *RelationDiscoveryV2Response_Data) Validate() error {
	for _, item := range this.ConnectSuccess {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConnectSuccess", err)
			}
		}
	}
	for _, item := range this.ConnectFail {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConnectFail", err)
			}
		}
	}
	for _, item := range this.DisconnectSuccess {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DisconnectSuccess", err)
			}
		}
	}
	for _, item := range this.DisconnectFail {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DisconnectFail", err)
			}
		}
	}
	for _, item := range this.ApproveList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ApproveList", err)
			}
		}
	}
	return nil
}
func (this *RelationDiscoveryV2ResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
