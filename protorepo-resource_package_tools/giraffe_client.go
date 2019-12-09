package resource_package_tools

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-resource_package_tools/org"
)

type Client struct {
	Org org.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Org = org.NewClient(c)
	return &client
}
