package terraform

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-terraform/backend"
)

type Client struct {
	Backend backend.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Backend = backend.NewClient(c)
	return &client
}
