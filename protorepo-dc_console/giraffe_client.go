package dc_console

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyopsapis/easyops-api-go/protorepo-dc_console/org"
)

type Client struct {
	Org org.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}

	client.Org = org.NewClient(c)
	return &client
}
