// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: traverse_graph.proto

package instance_graph

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

var _regex_TraverseGraphRequest_ObjectId = regexp.MustCompile(`^[a-zA-Z_]{1,32}$`)

func (this *TraverseGraphRequest) Validate() error {
	if !_regex_TraverseGraphRequest_ObjectId.MatchString(this.ObjectId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ObjectId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_]{1,32}$"`, this.ObjectId))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Fields != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fields); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fields", err)
		}
	}
	for _, item := range this.Child {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Child", err)
			}
		}
	}
	return nil
}
func (this *TraverseGraphResponse) Validate() error {
	for _, item := range this.TopicVertices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TopicVertices", err)
			}
		}
	}
	for _, item := range this.Vertices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Vertices", err)
			}
		}
	}
	for _, item := range this.Edges {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
			}
		}
	}
	return nil
}
func (this *TraverseGraphResponse_Edges) Validate() error {
	return nil
}
func (this *TraverseGraphResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}