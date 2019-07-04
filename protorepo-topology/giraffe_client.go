package topology

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/container"

	"github.com/easyopsapis/easyops-api-go/protorepo-topology/view"
)

type Client struct {
	Container container.Client

	View view.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Container = container.NewClient(c)

	client.View = view.NewClient(c)
	return &client
}
