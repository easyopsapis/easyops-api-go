package topology

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/container"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/topo_view"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/view"
)

type Client struct {
	Container container.Client

	Org org.Client

	TopoView topo_view.Client

	View view.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Container = container.NewClient(c)

	client.Org = org.NewClient(c)

	client.TopoView = topo_view.NewClient(c)

	client.View = view.NewClient(c)
	return &client
}
