// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list.proto

package micro_app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/easyopsapis/easyops-api-go/protorepo-models/easyops/model/app_store"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListAppStoreMicroAppRequest) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	if this.Sort != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sort); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sort", err)
		}
	}
	return nil
}

var _regex_ListAppStoreMicroAppRequest_Query_ProductLineId = regexp.MustCompile(`^[0-9a-z]{13}$`)

func (this *ListAppStoreMicroAppRequest_Query) Validate() error {
	if !_regex_ListAppStoreMicroAppRequest_Query_ProductLineId.MatchString(this.ProductLineId) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductLineId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9a-z]{13}$"`, this.ProductLineId))
	}
	return nil
}
func (this *ListAppStoreMicroAppRequest_Sort) Validate() error {
	return nil
}
func (this *ListAppStoreMicroAppResponse) Validate() error {
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

var _regex_ListAppStoreMicroAppResponse_List_Name = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_ListAppStoreMicroAppResponse_List_Id = regexp.MustCompile(`^[a-zA-Z_][0-9a-zA-Z_]{0,31}$`)
var _regex_ListAppStoreMicroAppResponse_List_Intro = regexp.MustCompile(`^.{0,200}$`)
var _regex_ListAppStoreMicroAppResponse_List_Homepage = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *ListAppStoreMicroAppResponse_List) Validate() error {
	if this.CurrentVersion != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CurrentVersion); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CurrentVersion", err)
		}
	}
	for _, item := range this.ProductLines {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ProductLines", err)
			}
		}
	}
	for _, item := range this.Solutions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Solutions", err)
			}
		}
	}
	if !_regex_ListAppStoreMicroAppResponse_List_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.Name))
	}
	if !_regex_ListAppStoreMicroAppResponse_List_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z_][0-9a-zA-Z_]{0,31}$"`, this.Id))
	}
	if this.Icons != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Icons); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Icons", err)
		}
	}
	if !_regex_ListAppStoreMicroAppResponse_List_Intro.MatchString(this.Intro) {
		return github_com_mwitkow_go_proto_validators.FieldError("Intro", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{0,200}$"`, this.Intro))
	}
	if !_regex_ListAppStoreMicroAppResponse_List_Homepage.MatchString(this.Homepage) {
		return github_com_mwitkow_go_proto_validators.FieldError("Homepage", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Homepage))
	}
	return nil
}

var _regex_ListAppStoreMicroAppResponse_List_Icons_Large = regexp.MustCompile(`^((\/[^\/:\*\?""<>\|\r\n]+)+)|(\/)|[a-zA-Z]:(\\[^\\:\*\?""<>\|\r\n]+)*$`)

func (this *ListAppStoreMicroAppResponse_List_Icons) Validate() error {
	if !_regex_ListAppStoreMicroAppResponse_List_Icons_Large.MatchString(this.Large) {
		return github_com_mwitkow_go_proto_validators.FieldError("Large", fmt.Errorf(`value '%v' must be a string conforming to regex "^((\\/[^\\/:\\*\\?\"\"<>\\|\\r\\n]+)+)|(\\/)|[a-zA-Z]:(\\\\[^\\\\:\\*\\?\"\"<>\\|\\r\\n]+)*$"`, this.Large))
	}
	return nil
}
func (this *ListAppStoreMicroAppResponseWrapper) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
