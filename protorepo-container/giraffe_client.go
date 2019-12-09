package container

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/cluster"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/configmap"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/event"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/image"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/namespace"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/node"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/org"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/persistentvolumeclaim"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/proxy"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/resourcegroup"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/secret"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/service"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/storageclass"

	"github.com/easyopsapis/easyops-api-go/protorepo-container/workload"
)

type Client struct {
	Cluster cluster.Client

	Configmap configmap.Client

	Event event.Client

	Image image.Client

	Namespace namespace.Client

	Node node.Client

	Org org.Client

	Persistentvolumeclaim persistentvolumeclaim.Client

	Proxy proxy.Client

	Resourcegroup resourcegroup.Client

	Secret secret.Client

	Service service.Client

	Storageclass storageclass.Client

	Workload workload.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Cluster = cluster.NewClient(c)

	client.Configmap = configmap.NewClient(c)

	client.Event = event.NewClient(c)

	client.Image = image.NewClient(c)

	client.Namespace = namespace.NewClient(c)

	client.Node = node.NewClient(c)

	client.Org = org.NewClient(c)

	client.Persistentvolumeclaim = persistentvolumeclaim.NewClient(c)

	client.Proxy = proxy.NewClient(c)

	client.Resourcegroup = resourcegroup.NewClient(c)

	client.Secret = secret.NewClient(c)

	client.Service = service.NewClient(c)

	client.Storageclass = storageclass.NewClient(c)

	client.Workload = workload.NewClient(c)
	return &client
}
