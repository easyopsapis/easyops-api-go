// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storyboard.proto

package api_gateway

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/micro_app"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StoryBoard) Validate() error {
	if this.App != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.App); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("App", err)
		}
	}
	for _, item := range this.Routes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Routes", err)
			}
		}
	}
	return nil
}
