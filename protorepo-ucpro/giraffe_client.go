package ucpro

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-ucpro/desktop"
)

type Client struct {
	Desktop desktop.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Desktop = desktop.NewClient(c)
	return &client
}
