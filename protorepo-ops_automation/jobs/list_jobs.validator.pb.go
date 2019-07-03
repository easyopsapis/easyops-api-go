// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_jobs.proto

package jobs

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/ops_automation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListJobsRequest) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *ListJobsResponse) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *ListJobsResponse_List) Validate() error {
	if this.Scheduler != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Scheduler); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Scheduler", err)
		}
	}
	if !(len(this.Id) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must length be greater than '0'`, this.Id))
	}
	if !(len(this.Id) < 34) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must length be less than '34'`, this.Id))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must length be greater than '0'`, this.Name))
	}
	if !(len(this.Category) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Category", fmt.Errorf(`value '%v' must length be greater than '0'`, this.Category))
	}
	return nil
}
func (this *ListJobsResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
