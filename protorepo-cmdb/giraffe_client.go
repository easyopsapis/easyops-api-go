package cmdb

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/business_instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/cmdb_object"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/history"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/initialization"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_archive"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_graph"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_path_search"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_relation"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_tree"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/object_attribute"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/object_relation"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/terraform"
)

type Client struct {
	BusinessInstance business_instance.Client

	CmdbObject cmdb_object.Client

	History history.Client

	Initialization initialization.Client

	Instance instance.Client

	InstanceArchive instance_archive.Client

	InstanceGraph instance_graph.Client

	InstancePathSearch instance_path_search.Client

	InstanceRelation instance_relation.Client

	InstanceTree instance_tree.Client

	ObjectAttribute object_attribute.Client

	ObjectRelation object_relation.Client

	Terraform terraform.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.BusinessInstance = business_instance.NewClient(c)

	client.CmdbObject = cmdb_object.NewClient(c)

	client.History = history.NewClient(c)

	client.Initialization = initialization.NewClient(c)

	client.Instance = instance.NewClient(c)

	client.InstanceArchive = instance_archive.NewClient(c)

	client.InstanceGraph = instance_graph.NewClient(c)

	client.InstancePathSearch = instance_path_search.NewClient(c)

	client.InstanceRelation = instance_relation.NewClient(c)

	client.InstanceTree = instance_tree.NewClient(c)

	client.ObjectAttribute = object_attribute.NewClient(c)

	client.ObjectRelation = object_relation.NewClient(c)

	client.Terraform = terraform.NewClient(c)
	return &client
}
