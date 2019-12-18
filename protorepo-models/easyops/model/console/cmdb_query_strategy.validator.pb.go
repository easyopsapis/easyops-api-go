// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmdb_query_strategy.proto

package console

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CmdbQueryStrategy) Validate() error {
	if this.Instances != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instances); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instances", err)
		}
	}
	if this.Fields != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fields); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
		}
	}
	return nil
}
func (this *CmdbQueryStrategy_Instances) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
