package ucpro

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-ucpro/desktop"

	"github.com/easyopsapis/easyops-api-go/protorepo-ucpro/org"
)

type Client struct {
	Desktop desktop.Client

	Org org.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Desktop = desktop.NewClient(c)

	client.Org = org.NewClient(c)
	return &client
}
