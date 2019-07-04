package cmdb

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/business_instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/cmdb_object"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_graph"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_relation"

	"github.com/easyopsapis/easyops-api-go/protorepo-cmdb/instance_tree"
)

type Client struct {
	BusinessInstance business_instance.Client

	CmdbObject cmdb_object.Client

	Instance instance.Client

	InstanceGraph instance_graph.Client

	InstanceRelation instance_relation.Client

	InstanceTree instance_tree.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.BusinessInstance = business_instance.NewClient(c)

	client.CmdbObject = cmdb_object.NewClient(c)

	client.Instance = instance.NewClient(c)

	client.InstanceGraph = instance_graph.NewClient(c)

	client.InstanceRelation = instance_relation.NewClient(c)

	client.InstanceTree = instance_tree.NewClient(c)
	return &client
}
